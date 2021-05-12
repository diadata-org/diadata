package scrapers

import (
	"bytes"
	"compress/flate"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	utils "github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
)

var _OKExSocketURL = "wss://ws.okex.com:8443/ws/v5/public"

//var _OKExSocketURL = url.URL{Scheme: "wss", Host: "real.okex.com:10441", Path: "/ws/v1", RawQuery: "compress=true"}

type Response struct {
	Channel string     `json:"channel"`
	Data    [][]string `json:"data"`
	Binary  int        `json:"binary"`
}

type Responses []Response

type Subscribe struct {
	OP   string     `json:"op"`
	Args []OKEXArgs `json:"args"`
}

type OKEXArgs struct {
	Channel string `json:"channel"`
	InstID  string `json:"instId"`
}

type OKExScraper struct {
	wsClient *ws.Conn
	// signaling channels for session initialization and finishing
	run          bool
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*OKExPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
}

// NewOKExScraper returns a new OKExScraper for the given pair
func NewOKExScraper(exchange dia.Exchange) *OKExScraper {

	s := &OKExScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*OKExPairScraper),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
	}

	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(_OKExSocketURL, nil)
	if err != nil {
		log.Error("dial:", err)
	}

	s.wsClient = SwConn
	go s.mainLoop()
	return s
}

// Useful to reconnect to ws when the connection is down
func (s *OKExScraper) reconnectToWS() {

	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(_OKExSocketURL, nil)

	if err != nil {
		log.Error("dial:", err)
	}

	s.wsClient = SwConn
}

type OKEXMarket struct {
	Alias     string `json:"alias"`
	BaseCcy   string `json:"baseCcy"`
	Category  string `json:"category"`
	CtMult    string `json:"ctMult"`
	CtType    string `json:"ctType"`
	CtVal     string `json:"ctVal"`
	CtValCcy  string `json:"ctValCcy"`
	ExpTime   string `json:"expTime"`
	InstID    string `json:"instId"`
	InstType  string `json:"instType"`
	Lever     string `json:"lever"`
	ListTime  string `json:"listTime"`
	LotSz     string `json:"lotSz"`
	MinSz     string `json:"minSz"`
	OptType   string `json:"optType"`
	QuoteCcy  string `json:"quoteCcy"`
	SettleCcy string `json:"settleCcy"`
	State     string `json:"state"`
	Stk       string `json:"stk"`
	TickSz    string `json:"tickSz"`
	Uly       string `json:"uly"`
}

type AllOKEXMarketResponse struct {
	Code string       `json:"code"`
	Data []OKEXMarket `json:"data"`
	Msg  string       `json:"msg"`
}

// Subscribe again to all channels
func (s *OKExScraper) subscribeToALL() {

	var (
		resp     AllOKEXMarketResponse
		allPairs []OKEXArgs
	)

	b, err := utils.GetRequest("https://aws.okex.com/api/v5/public/instruments?instType=SPOT")
	if err != nil {
		log.Errorln("Error getting OKex market", err)
	}

	err = json.Unmarshal(b, &resp)
	if err != nil {
		log.Errorln("Error Unmarshalling OKex market json", err)
	}

	for _, v := range resp.Data {
		allPairs = append(allPairs, OKEXArgs{Channel: "trades", InstID: v.InstID})
	}

	a := &Subscribe{
		OP:   "subscribe",
		Args: allPairs,
	}

	if err := s.wsClient.WriteJSON(a); err != nil {
		log.Errorln(err.Error())
	}

}

type OKEXWSResponse struct {
	Arg struct {
		Channel string `json:"channel"`
		InstID  string `json:"instId"`
	} `json:"arg"`
	Data []struct {
		InstID  string `json:"instId"`
		TradeID string `json:"tradeId"`
		Px      string `json:"px"`
		Sz      string `json:"sz"`
		Side    string `json:"side"`
		Ts      string `json:"ts"`
	} `json:"data"`
}

// runs in a goroutine until s is closed
func (s *OKExScraper) mainLoop() {

	s.run = true
	for s.run {
		var message OKEXWSResponse
		messageType, messageTemp, err := s.wsClient.ReadMessage()
		if err != nil {
			log.Warning("reconnect the scraping to ws, ", err, ":", message)
			s.reconnectToWS()
			s.subscribeToALL()
		} else {
			switch messageType {
			case ws.TextMessage:
				// no need uncompressed
				err := json.Unmarshal(messageTemp, &message)
				if err != nil {
					log.Errorln("Error parsing response")
				}
				ps, ok := s.pairScrapers[message.Arg.InstID]

				if ok && len(message.Data) > 0 {

					f64PriceString := message.Data[0].Px
					f64Price, err := strconv.ParseFloat(f64PriceString, 64)

					if err == nil {

						f64VolumeString := message.Data[0].Sz
						f64Volume, err := strconv.ParseFloat(f64VolumeString, 64)

						if err == nil {

							ts, _ := strconv.ParseInt(message.Data[0].Ts, 10, 64)
							timeStamp := time.Unix(int64(ts)/1e3, 0)
							if message.Data[0].Side == "sell" {
								f64Volume = -f64Volume
							}

							t := &dia.Trade{
								Symbol:         ps.pair.Symbol,
								Pair:           message.Arg.InstID,
								Price:          f64Price,
								Volume:         f64Volume,
								Time:           timeStamp,
								ForeignTradeID: message.Data[0].TradeID,
								Source:         s.exchangeName,
							}
							ps.parent.chanTrades <- t
							log.Infoln("Got trade", t)

						} else {
							log.Errorf("parsing volume %v", f64VolumeString)
						}

					} else {
						log.Errorf("parsing price %v", f64PriceString)
					}
				}

			}
		}
	}
	s.cleanup(errors.New("main loop terminated by Close()"))
}

func GzipDecode(in []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(in))
	defer reader.Close()

	return ioutil.ReadAll(reader)
}

func (s *OKExScraper) cleanup(err error) {
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
func (s *OKExScraper) Close() error {

	if s.closed {
		return errors.New("OKExScraper: Already closed")
	}

	close(s.shutdown)
	// Set false first to prevent reconnect
	s.run = false
	err := s.wsClient.Close()
	if err != nil {
		return err
	}
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *OKExScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()

	if s.error != nil {
		return nil, s.error
	}

	if s.closed {
		return nil, errors.New("OKExScraper: Call ScrapePair on closed scraper")
	}

	ps := &OKExPairScraper{
		parent: s,
		pair:   pair,
	}

	s.pairScrapers[pair.ForeignName] = ps

	//a := &Subscribe{
	//	OP:      "addChannel",
	//	Args: "ok_sub_spot_" + strings.ToLower(pair.ForeignName) + "_deals",
	//}
	//
	//subByteString := `{"channel":` + `"` + a.Channel + `"` + `,"event":` + `"` + a.Event + `"}`
	//if err := s.wsClient.WriteMessage(ws.TextMessage, []byte(subByteString)); err != nil {
	//	fmt.Println(err.Error())
	//}

	return ps, nil
}
func (s *OKExScraper) normalizeSymbol(foreignName string, baseCurrency string) (symbol string, err error) {
	symbol = strings.ToUpper(baseCurrency)
	if helpers.NameForSymbol(symbol) == symbol {
		if !helpers.SymbolIsName(symbol) {
			if symbol == "IOTA" {
				return "MIOTA", nil
			}
			if symbol == "YOYO" {
				return "YOYOW", nil
			}
			return symbol, errors.New("Foreign name can not be normalized:" + foreignName + " symbol:" + symbol)
		}
	}
	if helpers.SymbolIsBlackListed(symbol) {
		return symbol, errors.New("Symbol is black listed:" + symbol)
	}
	return symbol, nil
}

func (s *OKExScraper) NormalizePair(pair dia.Pair) (dia.Pair, error) {
	symbol := strings.ToUpper(pair.Symbol)
	pair.Symbol = symbol

	if helpers.NameForSymbol(symbol) == symbol {
		if !helpers.SymbolIsName(symbol) {
			if pair.Symbol == "IOTA" {
				pair.Symbol = "MIOTA"
			}
			if pair.Symbol == "YOYO" {
				pair.Symbol = "YOYOW"
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
func (s *OKExScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	type APIResponse struct {
		Id           string `json:"instrument_id"`
		BaseCurrency string `json:"base_currency"`
	}

	data, err := utils.GetRequest("https://www.okex.com/api/spot/v3/products")

	if err != nil {
		return
	}

	var ar []APIResponse
	err = json.Unmarshal(data, &ar)
	if err == nil {
		for _, p := range ar {
			pairToNormalize := dia.Pair{
				Symbol:      p.BaseCurrency,
				ForeignName: p.Id,
				Exchange:    s.exchangeName,
			}
			pair, serr := s.NormalizePair(pairToNormalize)
			if serr == nil {
				pairs = append(pairs, pair)
			} else {
				log.Error(serr)
			}
		}
	}
	return
}

// OKExPairScraper implements PairScraper for OKEx exchange
type OKExPairScraper struct {
	parent *OKExScraper
	pair   dia.Pair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *OKExPairScraper) Close() error {
	return nil
}

// Channel returns a channel that can be used to receive trades
func (s *OKExScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *OKExPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *OKExPairScraper) Pair() dia.Pair {
	return ps.pair
}
