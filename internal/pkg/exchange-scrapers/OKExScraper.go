package scrapers

import (
	"bytes"
	"compress/flate"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	utils "github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var _OKExSocketURL = url.URL{Scheme: "wss", Host: "real.okex.com:10441", Path: "/ws/v1", RawQuery: "compress=true"}

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
	chanTrades   chan *dia.Trade
}

// NewOKExScraper returns a new OKExScraper for the given pair
func NewOKExScraper(exchangeName string) *OKExScraper {

	s := &OKExScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*OKExPairScraper),
		exchangeName: exchangeName,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
	}

	SwConn, _, err := ws.DefaultDialer.Dial(_OKExSocketURL.String(), nil)

	if err != nil {
		log.Error("dial:", err)
	}

	s.wsClient = SwConn
	go s.mainLoop()
	return s
}

// Useful to reconnect to ws when the connection is down
func (s *OKExScraper) reconnectToWS() {

	SwConnNew, _, err := ws.DefaultDialer.Dial(_OKExSocketURL.String(), nil)

	if err != nil {
		log.Error("dial:", err)
	}

	s.wsClient = SwConnNew
}

// Subscribe again to all channels
func (s *OKExScraper) subscribeToALL() {

	for key := range s.pairScrapers {
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
		messageType, messageTemp, err := s.wsClient.ReadMessage()
		if err != nil {
			log.Warning("reconnect the scraping to ws, ", err, ":", message)
			s.reconnectToWS()
			s.subscribeToALL()
		} else {
			switch messageType {
			case ws.TextMessage:
				// no need uncompressed
				fmt.Printf("recv text: %s", messageTemp)
			case ws.BinaryMessage:
				// uncompressed
				text, err := GzipDecode(messageTemp)
				if err != nil {
					log.Error("err", err)
					return
				}
				err = json.Unmarshal(text, &message)
				if err != nil {
					log.Error("err", err)
					return
				}

				//skip subscribe channel confirm data
				if message[0].Binary == 0 {

					var splitString = strings.Split(message[0].Channel, "_")
					var forName = strings.ToUpper(splitString[3] + "_" + splitString[4])
					ps, ok := s.pairScrapers[forName]

					if ok {

						f64PriceString := message[0].Data[0][1]
						f64Price, err := strconv.ParseFloat(f64PriceString, 64)

						if err == nil {

							f64VolumeString := message[0].Data[0][2]
							f64Volume, err := strconv.ParseFloat(f64VolumeString, 64)

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
								ps.parent.chanTrades <- t

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

	a := &Subscribe{
		Event:   "addChannel",
		Channel: "ok_sub_spot_" + strings.ToLower(pair.ForeignName) + "_deals",
	}

	subByteString := `{"channel":` + `"` + a.Channel + `"` + `,"event":` + `"` + a.Event + `"}`
	if err := s.wsClient.WriteMessage(ws.TextMessage, []byte(subByteString)); err != nil {
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

	data, err := utils.GetRequest("https://www.okex.com/api/spot/v3/products")

	if err != nil {
		return
	}

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
