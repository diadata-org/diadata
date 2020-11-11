package scrapers

import (
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

var _GateIOsocketurl string = "wss://ws.gate.io/v3"

type SubscribeGate struct {
	Id     int      `json:"id"`
	Method string   `json:"method"`
	Params []string `json:"params"`
}

type ResponseGate struct {
	Method string        `json:"method,omitempty"`
	Params []interface{} `json:"params,omitempty"`
	Id     interface{}   `json:"id,omitempty"`
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
func NewGateIOScraper(exchangeName string) *GateIOScraper {

	s := &GateIOScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*GateIOPairScraper),
		exchangeName: exchangeName,
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

// runs in a goroutine until s is closed
func (s *GateIOScraper) mainLoop() {

	// wait for all pairs have added into s.PairScrapers
	time.Sleep(4 * time.Second)
	allPairs := make([]string, len(s.pairScrapers))
	var index = 0
	for key, _ := range s.pairScrapers {
		allPairs[index] = key
		index += 1
	}

	// Only one subscribe for all pairs
	a := &SubscribeGate{
		Id:     12312,
		Method: "trades.subscribe",
		Params: allPairs,
	}
	var err error
	if err = s.wsClient.WriteJSON(a); err != nil {
		log.Error(err.Error())
	}

	for true {

		message := &ResponseGate{}
		if err = s.wsClient.ReadJSON(&message); err != nil {
			log.Error(err.Error())
			break
		}
		var pairRetrieved string
		for key, v := range message.Params {
			// key 0 -> pair
			// key 1 -> datas
			if key == 0 {
				pairRetrieved = v.(string)
			}
			if key == 1 {
				ps, ok := s.pairScrapers[pairRetrieved]
				if ok {

					md := v.([]interface{})
					for _, v := range md {

						md_inner := v.(map[string]interface{})
						f64Price_string := md_inner["price"].(string)
						f64Price, err := strconv.ParseFloat(f64Price_string, 64)

						if err == nil {

							f64Volume_string := md_inner["amount"].(string)
							f64Volume, err := strconv.ParseFloat(f64Volume_string, 64)

							if err == nil {
								if md_inner["type"] == "sell" {
									f64Volume = -f64Volume
								}
								timestamp_temp := int64(md_inner["time"].(float64))
								timestamp := time.Unix(timestamp_temp, 0).UTC()
								t := &dia.Trade{
									Symbol:         ps.pair.Symbol,
									Pair:           pairRetrieved,
									Price:          f64Price,
									Volume:         f64Volume,
									Time:           timestamp,
									ForeignTradeID: strconv.FormatInt(int64(md_inner["id"].(float64)), 16),
									Source:         s.exchangeName,
								}
								ps.parent.chanTrades <- t
							} else {
								log.Error("error parsing volume %v " + md_inner["amount"].(string))
							}
						} else {
							log.Error("error parsing price %v " + md_inner["price"].(string))
						}
					}
				}
			}
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
