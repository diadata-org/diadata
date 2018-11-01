package scrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	ws "github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var _OKExSocketurl string = "wss://real.okex.com:10441/websocket"

type Response struct {
	Channel string     `json:"channel"`
	Data    [][]string `json:"data"`
	Binary  int        `json:"binary"`
}

type Responses []Response

type Subscribe struct {
	Event   string `json:"event"`
	Channel string `json:"channel"`
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
}

// NewOKExScraper returns a new OKExScraper for the given pair
func NewOKExScraper(exchangeName string) *OKExScraper {

	s := &OKExScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*OKExPairScraper),
		exchangeName: exchangeName,
		error:        nil,
	}

	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(_OKExSocketurl, nil)

	if err != nil {
		println(err.Error())
	}

	s.wsClient = SwConn
	go s.mainLoop()
	return s
}

// Useful to reconnect to ws when the connection is down
func (s *OKExScraper) reconnectToWS() {
	var wsDialer ws.Dialer
	SwConn_new, _, err := wsDialer.Dial(_OKExSocketurl, nil)

	if err != nil {
		println(err.Error())
	}

	s.wsClient = SwConn_new
}

// Subscribe again to all channels
func (s *OKExScraper) subscribeToALL() {

	for key, _ := range s.pairScrapers {
		a := &Subscribe{
			Event:   "addChannel",
			Channel: "ok_sub_spot_" + strings.ToLower(key) + "_deals",
		}
		if err := s.wsClient.WriteJSON(a); err != nil {
			fmt.Println(err.Error())
		}
	}
}

// runs in a goroutine until s is closed
func (s *OKExScraper) mainLoop() {

	s.run = true
	for s.run {
		var message Responses
		if err := s.wsClient.ReadJSON(&message); err != nil {
			jsonError := strings.HasPrefix(err.Error(), "json:")
			if jsonError == false {
				log.Warning("reconnect the scraping to ws, ", err, ":", message)
				s.reconnectToWS()
				s.subscribeToALL()
			}
		} else {

			var splitString = strings.Split(message[0].Channel, "_")
			var forName = strings.ToUpper(splitString[3] + "_" + splitString[4])
			ps, ok := s.pairScrapers[forName]

			if ok {

				f64Price_string := message[0].Data[0][1]
				f64Price, err := strconv.ParseFloat(f64Price_string, 64)

				if err == nil {

					f64Volume_string := message[0].Data[0][2]
					f64Volume, err := strconv.ParseFloat(f64Volume_string, 64)

					if err == nil {

						timeStamp := time.Now().UTC()
						if message[0].Data[0][4] == "ask" {
							f64Volume = -f64Volume
						}

						t := &dia.Trade{
							Symbol:         ps.pair.Symbol,
							Pair:           forName,
							Price:          f64Price,
							Volume:         f64Volume,
							Time:           timeStamp,
							ForeignTradeID: message[0].Data[0][0],
							Source:         s.exchangeName,
						}
						ps.chanTrades <- t

					} else {
						log.Error("parsing volume %v", f64Volume_string)
					}

				} else {
					log.Error("parsing price %v", f64Price_string)
				}
			}
		}
	}
	s.cleanup(errors.New("Main loop terminated by Close()"))
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
	s.wsClient.Close()
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
		parent:     s,
		pair:       pair,
		chanTrades: make(chan *dia.Trade),
	}

	s.pairScrapers[pair.ForeignName] = ps

	a := &Subscribe{
		Event:   "addChannel",
		Channel: "ok_sub_spot_" + strings.ToLower(pair.ForeignName) + "_deals",
	}

	if err := s.wsClient.WriteJSON(a); err != nil {
		fmt.Println(err.Error())
	}

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

// FetchAvailablePairs returns a list with all available trade pairs
func (s *OKExScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	type APIResponse struct {
		Id           string `json:"instrument_id"`
		BaseCurrency string `json:"base_currency"`
	}
	response, err := http.Get("https://www.okex.com/api/spot/v3/products")
	if err != nil {
		log.Error("The HTTP request failed:", err)
		return
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	var ar []APIResponse
	err = json.Unmarshal(data, &ar)
	if err == nil {
		for _, p := range ar {
			symbol, serr := s.normalizeSymbol(p.Id, p.BaseCurrency)
			if serr == nil {
				pairs = append(pairs, dia.Pair{
					Symbol:      symbol,
					ForeignName: p.Id,
					Exchange:    s.exchangeName,
				})
			} else {
				log.Error(serr)
			}
		}
	}
	return
}

// OKExPairScraper implements PairScraper for OKEx exchange
type OKExPairScraper struct {
	parent     *OKExScraper
	pair       dia.Pair
	chanTrades chan *dia.Trade
	closed     bool
}

// Close stops listening for trades of the pair associated with s
func (ps *OKExPairScraper) Close() error {
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *OKExPairScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
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
