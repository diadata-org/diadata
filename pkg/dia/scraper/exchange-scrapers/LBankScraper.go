package scrapers

import (
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	models "github.com/diadata-org/diadata/pkg/model"
	utils "github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
)

var _LBankSocketurl string = "wss://www.lbkex.net/ws/V2/"

const (
	timeZoneLBank      = "Asia/Singapore"
	timeFormatResponse = "2006-01-02T15:04:05"
)

type ResponseLBank struct {
	Pair   string      `json:"pair"`
	Trade  interface{} `json:"trade"`
	Type   string      `json:"type"`
	Server string      `json:"SERVER"`
	Ts     string      `json:"TS"`
}

type SubscribeLBank struct {
	Action    string `json:"action"`
	Subscribe string `json:"subscribe"`
	Pair      string `json:"pair"`
}

type SubscribePing struct {
	Action string `json:"action"`
}

type PongMessage struct {
	Action string `json:"action"`
	Value  string `json:"pong"`
}

type LBankScraper struct {
	wsClient *ws.Conn
	// signaling channels for session initialization and finishing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*LBankPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB
}

// NewLBankScraper returns a new LBankScraper for the given pair
func NewLBankScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *LBankScraper {

	s := &LBankScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*LBankPairScraper),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}

	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(_LBankSocketurl, nil)
	if err != nil {
		println(err.Error())
	}
	s.wsClient = SwConn

	if scrape {
		go s.mainLoop()
	}
	return s
}

// runs in a goroutine until s is closed
func (s *LBankScraper) mainLoop() {
	var err error
	defer s.cleanup(err)
	// Wait until all pairs are subscribed in order to prevent concurrent ws write.
	time.Sleep(10 * time.Second)
	err = s.subscribePing()
	if err != nil {
		log.Fatal("subscribe ping: ", err)
	}
	for {
		var message map[string]interface{}
		if err = s.wsClient.ReadJSON(&message); err != nil {
			log.Error("read ping: ", err)
		}
		if messageType, ok := message["type"]; ok {
			if messageType == "trade" {
				pair := strings.ToUpper(message["pair"].(string))
				ps, ok := s.pairScrapers[pair]
				if ok {
					var exchangepair dia.ExchangePair
					var timestamp time.Time

					tradeMap := message["trade"].(map[string]interface{})
					f64Price := tradeMap["price"].(float64)
					f64Volume := tradeMap["volume"].(float64)
					if tradeMap["direction"] == "sell" {
						f64Volume = -f64Volume
					}
					timestamp, err = parseAsianTime(message["TS"].(string))
					if err != nil {
						log.Error("parse time: ", err)
					}

					exchangepair, err = s.db.GetExchangePairCache(s.exchangeName, pair)
					if err != nil {
						log.Error(err)
					}
					t := &dia.Trade{
						Symbol:       ps.pair.Symbol,
						Pair:         pair,
						Price:        f64Price,
						Volume:       f64Volume,
						Time:         timestamp,
						Source:       s.exchangeName,
						VerifiedPair: exchangepair.Verified,
						BaseToken:    exchangepair.UnderlyingPair.BaseToken,
						QuoteToken:   exchangepair.UnderlyingPair.QuoteToken,
					}
					if exchangepair.Verified {
						log.Infoln("Got verified trade", t)
					}
					ps.parent.chanTrades <- t
				}
			}
		} else {
			if pingMessage, ok := message["ping"]; ok {
				var pongMessageMarshalled []byte
				pongMessage := PongMessage{
					Action: "pong",
					Value:  pingMessage.(string),
				}
				pongMessageMarshalled, err = json.Marshal(pongMessage)
				if err != nil {
					log.Error("marshal pong: ", err)
				}
				err = s.pong(pongMessageMarshalled)
				if err != nil {
					log.Error("send pong: ", err)
				}
			}
		}
	}
}

// Pong sends the string "pong" to the server.
func (s *LBankScraper) pong(message []byte) error {
	err := s.wsClient.WriteMessage(1, message)
	if err != nil {
		return err
	}
	return nil
}

func (s *LBankScraper) subscribePing() error {
	a := &SubscribePing{
		Action: "ping",
	}
	return s.wsClient.WriteJSON(a)
}

func parseAsianTime(timestring string) (time.Time, error) {
	IST, err := time.LoadLocation(timeZoneLBank)
	if err != nil {
		return time.Time{}, err
	}
	timestampAsia, err := time.ParseInLocation(timeFormatResponse, timestring, IST)
	if err != nil {
		return time.Time{}, err
	}
	UTC, err := time.LoadLocation("UTC")
	if err != nil {
		return time.Time{}, err
	}
	return timestampAsia.In(UTC), nil
}

func (s *LBankScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone)
}

func (s *LBankScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *LBankScraper) Close() error {

	if s.closed {
		return errors.New("LBankScraper: Already closed")
	}
	err := s.wsClient.Close()
	if err != nil {
		return err
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *LBankScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("LBankScraper: Call ScrapePair on closed scraper")
	}
	ps := &LBankPairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	a := &SubscribeLBank{
		Action:    "subscribe",
		Subscribe: "trade",
		Pair:      strings.ToLower(pair.ForeignName),
	}
	if err := s.wsClient.WriteJSON(a); err != nil {
		log.Error("ScrapePair" + err.Error())
	}
	return ps, nil
}

func (s *LBankScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	str := strings.Split(pair.ForeignName, "_")
	symbol := strings.ToUpper(str[0])
	pair.Symbol = symbol
	if helpers.NameForSymbol(symbol) == symbol {
		if !helpers.SymbolIsName(symbol) {
			return pair, errors.New("Foreign name can not be normalized:" + pair.ForeignName + " symbol:" + symbol)
		}
	}
	if helpers.SymbolIsBlackListed(symbol) {
		return pair, errors.New("Symbol is black listed:" + symbol)
	}
	return pair, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *LBankScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {

	data, _, err := utils.GetRequest("https://api.lbkex.com/v1/currencyPairs.do")
	if err != nil {
		return
	}
	ls := strings.Split(strings.Replace(string(data)[1:len(data)-1], "\"", "", -1), ",")
	for _, p := range ls {
		pairToNormalize := dia.ExchangePair{
			Symbol:      "",
			ForeignName: strings.ToUpper(p),
			Exchange:    s.exchangeName,
		}
		pair, serr := s.NormalizePair(pairToNormalize)
		if serr == nil {
			pairs = append(pairs, pair)
		} else {
			log.Error(serr)
		}
	}
	return

}

// LBankPairScraper implements PairScraper for LBank exchange
type LBankPairScraper struct {
	parent *LBankScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *LBankPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *LBankScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *LBankPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *LBankPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
