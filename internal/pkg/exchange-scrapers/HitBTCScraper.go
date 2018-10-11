package scrapers

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	ws "github.com/gorilla/websocket"
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
	//initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*HitBTCPairScraper
}

// NewHitBTCScraper returns a new HitBTCScraper for the given pair
func NewHitBTCScraper() *HitBTCScraper {

	s := &HitBTCScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*HitBTCPairScraper),
		error:        nil,
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

	for true {

		message := &Event{}
		if err := s.wsClient.ReadJSON(&message); err != nil {
			println(err.Error())
			break
		}

		if message.Method == "updateTrades" {

			md := message.Params.(map[string]interface{})
			ps, ok := s.pairScrapers[md["symbol"].(string)]

			if ok {

				md_data := md["data"].([]interface{})

				for _, v := range md_data {

					md_element := v.(map[string]interface{})

					f64Price_string := md_element["price"].(string)
					f64Price, err := strconv.ParseFloat(f64Price_string, 64)

					if err == nil {

						f64Volume_string := md_element["quantity"].(string)
						f64Volume, err := strconv.ParseFloat(f64Volume_string, 64)

						if err == nil {
							timeStamp, _ := time.Parse(time.RFC3339, md_element["timestamp"].(string))
							if md_element["id"] != 0 {

								if md_element["side"] == "sell" {
									f64Volume = -f64Volume
								}
								t := &dia.Trade{
									Symbol:         ps.pair.Symbol,
									Pair:           md["symbol"].(string),
									Price:          f64Price,
									Volume:         f64Volume,
									Time:           timeStamp,
									ForeignTradeID: strconv.FormatInt(int64(md_element["id"].(float64)), 16),
									Source:         dia.HitBTCExchange,
								}
								ps.chanTrades <- t
							}
						} else {
							log.Printf("error parsing volume %v", md_element["quantity"].(string))
						}
					} else {
						log.Printf("error parsing price %v", md_element["price"].(string))
					}
				}
			} else {
				log.Printf("Unknown Pair %v", md["symbol"].(string))
			}
		}
	}
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
		parent:     s,
		pair:       pair,
		chanTrades: make(chan *dia.Trade),
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

// HitBTCPairScraper implements PairScraper for HitBTC
type HitBTCPairScraper struct {
	parent     *HitBTCScraper
	pair       dia.Pair
	chanTrades chan *dia.Trade
	closed     bool
}

// Close stops listening for trades of the pair associated with s
func (ps *HitBTCPairScraper) Close() error {
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *HitBTCPairScraper) Channel() chan *dia.Trade {
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
