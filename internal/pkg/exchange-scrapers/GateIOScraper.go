package scrapers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	utils "github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
)

var _GateIOsocketurl string = "wss://api.gateio.ws/ws/v4/"

type ResponseGate struct {
	Method string        `json:"method,omitempty"`
	Params []interface{} `json:"params,omitempty"`
	Id     interface{}   `json:"id,omitempty"`
}

type SubscribeGate struct {
	Time    int64    `json:"time"`
	Channel string   `json:"channel"`
	Event   string   `json:"event"`
	Payload []string `json:"payload"`
}

type GateIOScraper struct {
	wsClient *ws.Conn
	// signaling channels for session initialization and finishing
	//initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*GateIOPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
}

// NewGateIOScraper returns a new GateIOScraper for the given pair
func NewGateIOScraper(exchange dia.Exchange) *GateIOScraper {

	s := &GateIOScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*GateIOPairScraper),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
	}

	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(_GateIOsocketurl, nil)

	if err != nil {
		println(err.Error())
	}

	s.wsClient = SwConn
	go s.mainLoop()
	return s
}

type GateIPPairResponse []GateIOPair
type GateIOPair struct {
	ID              string `json:"id"`
	Base            string `json:"base"`
	Quote           string `json:"quote"`
	Fee             string `json:"fee"`
	MinQuoteAmount  string `json:"min_quote_amount,omitempty"`
	AmountPrecision int    `json:"amount_precision"`
	Precision       int    `json:"precision"`
	TradeStatus     string `json:"trade_status"`
	SellStart       int    `json:"sell_start"`
	BuyStart        int    `json:"buy_start"`
	MinBaseAmount   string `json:"min_base_amount,omitempty"`
}

type GateIOResponseTrade struct {
	Time    int    `json:"time"`
	Channel string `json:"channel"`
	Event   string `json:"event"`
	Result  struct {
		ID           int    `json:"id"`
		CreateTime   int    `json:"create_time"`
		CreateTimeMs string `json:"create_time_ms"`
		Side         string `json:"side"`
		CurrencyPair string `json:"currency_pair"`
		Amount       string `json:"amount"`
		Price        string `json:"price"`
	} `json:"result"`
}

// runs in a goroutine until s is closed
func (s *GateIOScraper) mainLoop() {
	var (
		gresponse GateIPPairResponse
		allPairs  []string
	)
	var err error
	var f64Price float64
	var f64Volume float64

	b, _ := utils.GetRequest("https://api.gateio.ws/api/v4/spot/currency_pairs")
	err = json.Unmarshal(b, &gresponse)
	if err != nil {
		log.Error("deconding currency pairs: ", err)
	}

	for _, v := range gresponse {
		allPairs = append(allPairs, v.ID)
	}

	a := &SubscribeGate{
		Event:   "subscribe",
		Time:    time.Now().Unix(),
		Channel: "spot.trades",
		Payload: allPairs,
	}

	log.Infoln("subscribed", allPairs)
	if err = s.wsClient.WriteJSON(a); err != nil {
		log.Error(err.Error())
	}

	for {

		var message GateIOResponseTrade
		if err = s.wsClient.ReadJSON(&message); err != nil {
			log.Error(err.Error())
			break
		}

		ps, ok := s.pairScrapers[message.Result.CurrencyPair]
		if ok {

			f64Price, err = strconv.ParseFloat(message.Result.Price, 64)
			if err != nil {
				log.Errorln("error parsing float Price", err)
				continue
			}

			f64Volume, err = strconv.ParseFloat(message.Result.Amount, 64)
			if err != nil {
				log.Errorln("error parsing float Price", err)
				continue
			}

			if message.Result.Side == "sell" {
				f64Volume = -f64Volume
			}

			t := &dia.Trade{
				Symbol:         ps.pair.Symbol,
				Pair:           message.Result.CurrencyPair,
				Price:          f64Price,
				Volume:         f64Volume,
				Time:           time.Unix(int64(message.Result.CreateTime), 0),
				ForeignTradeID: strconv.FormatInt(int64(message.Result.ID), 16),
				Source:         s.exchangeName,
			}
			ps.parent.chanTrades <- t
			log.Infoln("got trade", t)

		}

	}
	s.cleanup(err)
}

func (s *GateIOScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone)
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *GateIOScraper) Close() error {

	if s.closed {
		return errors.New("GateIOScraper: Already closed")
	}
	s.wsClient.Close()
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *GateIOScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()

	if s.error != nil {
		return nil, s.error
	}

	if s.closed {
		return nil, errors.New("GateIOScraper: Call ScrapePair on closed scraper")
	}

	ps := &GateIOPairScraper{
		parent: s,
		pair:   pair,
	}

	s.pairScrapers[pair.ForeignName] = ps

	return ps, nil
}
func (s *GateIOScraper) normalizeSymbol(foreignName string, params ...interface{}) (symbol string, err error) {
	str := strings.Split(foreignName, "_")
	symbol = strings.ToUpper(str[0])
	if helpers.NameForSymbol(symbol) == symbol {
		if !helpers.SymbolIsName(symbol) {
			if symbol == "IOTA" {
				return "MIOTA", nil
			}
			return symbol, errors.New("Foreign name can not be normalized:" + foreignName + " symbol:" + symbol)
		}
	}
	if helpers.SymbolIsBlackListed(symbol) {
		return symbol, errors.New("Symbol is black listed:" + symbol)
	}
	return symbol, nil
}

func (s *GateIOScraper) NormalizePair(pair dia.Pair) (dia.Pair, error) {
	str := strings.Split(pair.ForeignName, "_")
	symbol := strings.ToUpper(str[0])
	pair.Symbol = symbol
	if helpers.NameForSymbol(symbol) == symbol {
		if !helpers.SymbolIsName(symbol) {
			if symbol == "IOTA" {
				pair.Symbol = "MIOTA"
			}
			return pair, errors.New("Foreign name can not be normalized:" + pair.ForeignName + " symbol:" + symbol)
		}
	}
	if helpers.SymbolIsBlackListed(symbol) {
		return pair, errors.New("Symbol is black listed:" + symbol)
	}
	return pair, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *GateIOScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	data, err := utils.GetRequest("https://data.gate.io/api2/1/pairs")
	if err != nil {
		return
	}
	ls := strings.Split(strings.Replace(string(data)[1:len(data)-1], "\"", "", -1), ",")
	for _, p := range ls {
		pairToNormalize := dia.Pair{
			Symbol:      "",
			ForeignName: p,
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

// GateIOPairScraper implements PairScraper for GateIO
type GateIOPairScraper struct {
	parent *GateIOScraper
	pair   dia.Pair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *GateIOPairScraper) Close() error {
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *GateIOScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *GateIOPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *GateIOPairScraper) Pair() dia.Pair {
	return ps.pair
}
