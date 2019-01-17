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

var _socketurl string = "wss://api.hitbtc.com/api/2/ws"

const WS_TIMEOUT = 10 * time.Second

type Event struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
	Id     int         `json:"id"`
}

type HitBTCScraper struct {
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
	pairScrapers map[string]*HitBTCPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
}

// NewHitBTCScraper returns a new HitBTCScraper for the given pair
func NewHitBTCScraper(exchangeName string) *HitBTCScraper {

	s := &HitBTCScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*HitBTCPairScraper),
		exchangeName: exchangeName,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
	}

	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(_socketurl, nil)

	if err != nil {
		println(err.Error())
	}

	s.wsClient = SwConn
	go s.mainLoop()
	return s
}

// runs in a goroutine until s is closed
func (s *HitBTCScraper) mainLoop() {
	var err error
	for true {
		message := &Event{}
		if err = s.wsClient.ReadJSON(&message); err != nil {
			log.Error(err.Error())
			break
		}
		if message.Method == "updateTrades" {
			md := message.Params.(map[string]interface{})
			ps, ok := s.pairScrapers[md["symbol"].(string)]
			if ok {
				mdData := md["data"].([]interface{})
				for _, v := range mdData {
					mdElement := v.(map[string]interface{})
					f64PriceString := mdElement["price"].(string)
					f64Price, err := strconv.ParseFloat(f64PriceString, 64)
					if err == nil {
						f64VolumeString := mdElement["quantity"].(string)
						f64Volume, err := strconv.ParseFloat(f64VolumeString, 64)
						if err == nil {
							timeStamp, _ := time.Parse(time.RFC3339, mdElement["timestamp"].(string))
							if mdElement["id"] != 0 {
								if mdElement["side"] == "sell" {
									f64Volume = -f64Volume
								}
								t := &dia.Trade{
									Symbol:         ps.pair.Symbol,
									Pair:           md["symbol"].(string),
									Price:          f64Price,
									Volume:         f64Volume,
									Time:           timeStamp,
									ForeignTradeID: strconv.FormatInt(int64(mdElement["id"].(float64)), 16),
									Source:         s.exchangeName,
								}
								ps.parent.chanTrades <- t
							}
						} else {
							log.Error("error parsing volume " + mdElement["quantity"].(string))
						}
					} else {
						log.Error("error parsing price " + mdElement["price"].(string))
					}
				}
			} else {
				log.Error("Unknown Pair " + md["symbol"].(string))
			}
		}
	}
	s.cleanup(err)
}

func (s *HitBTCScraper) cleanup(err error) {
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
func (s *HitBTCScraper) Close() error {
	if s.closed {
		return errors.New("HitBTCScraper: Already closed")
	}
	close(s.shutdown)
	s.wsClient.Close()
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *HitBTCScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()

	if s.error != nil {
		return nil, s.error
	}

	if s.closed {
		return nil, errors.New("HitBTCScraper: Call ScrapePair on closed scraper")
	}

	ps := &HitBTCPairScraper{
		parent: s,
		pair:   pair,
	}

	s.pairScrapers[pair.ForeignName] = ps

	a := &Event{
		Method: "subscribeTrades",
		Params: map[string]interface{}{
			"symbol": pair.ForeignName,
		},
		Id: int(time.Now().Unix()) * 1000,
	}

	if err := s.wsClient.WriteJSON(a); err != nil {
		fmt.Println(err.Error())
	}

	return ps, nil
}
func (s *HitBTCScraper) normalizeSymbol(foreignName string, baseCurrency string) (symbol string, err error) {
	symbol = strings.ToUpper(baseCurrency)
	if helpers.NameForSymbol(symbol) == symbol {
		if !helpers.SymbolIsName(symbol) {
			return symbol, errors.New("Foreign name can not be normalized:" + foreignName + " symbol:" + symbol)
		}
	}
	if helpers.SymbolIsBlackListed(symbol) {
		return symbol, errors.New("Symbol is black listed:" + symbol)
	}
	return symbol, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *HitBTCScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	type APIResponse struct {
		Id                   string  `json:"id"`
		BaseCurrency         string  `json:"baseCurrency"`
		QuoteCurrency        string  `json:"quoteCurrency"`
		QuantityIncrement    float64 `json:"quantityIncrement,string"`
		TickSize             float64 `json:"tickSize,string"`
		TakeLiquidityRate    float64 `json:"takeLiquidityRate,string"`
		ProvideLiquidityRate float64 `json:"provideLiquidityRate,string"`
		FeeCurrency          string  `json:"feeCurrency"`
	}
	response, err := http.Get("https://api.hitbtc.com/api/2/public/symbol")
	if err != nil {
		log.Error("The HTTP request failed:", err)
		return
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	var ar []APIResponse
	err = json.Unmarshal(data, &ar)
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

// HitBTCPairScraper implements PairScraper for HitBTC
type HitBTCPairScraper struct {
	parent *HitBTCScraper
	pair   dia.Pair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *HitBTCPairScraper) Close() error {
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *HitBTCScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *HitBTCPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *HitBTCPairScraper) Pair() dia.Pair {
	return ps.pair
}
