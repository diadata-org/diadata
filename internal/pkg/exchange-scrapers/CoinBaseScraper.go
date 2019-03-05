package scrapers

import (
	"encoding/json"
	"errors"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	ws "github.com/gorilla/websocket"
	gdax "github.com/preichenberger/go-gdax"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type CoinBaseScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock    sync.RWMutex
	error        error
	closed       bool
	pairScrapers map[string]*CoinBasePairScraper // pc.Pair -> pairScraperSet
	wsConn       *ws.Conn
	exchangeName string
	chanTrades   chan *dia.Trade
}

const (
	ChannelHeartbeat = "heartbeat"
	ChannelTicker    = "ticker"
	ChannelLevel2    = "level2"
	ChannelUser      = "user"
	ChannelMatches   = "matches"
	ChannelFull      = "full"
)

// NewCoinBaseScraper returns a new CoinBaseScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func NewCoinBaseScraper(exchangeName string) *CoinBaseScraper {
	s := &CoinBaseScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*CoinBasePairScraper),
		exchangeName: exchangeName,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
	}
	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial("wss://ws-feed.pro.coinbase.com", nil)
	if err != nil {
		println(err.Error())
	}
	s.wsConn = SwConn
	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *CoinBaseScraper) mainLoop() {
	var err error
	for true {
		message := gdax.Message{}
		if err = s.wsConn.ReadJSON(&message); err != nil {
			println(err.Error())
			break
		}
		if message.Type == ChannelTicker {
			ps, ok := s.pairScrapers[message.ProductID]
			if ok {
				f64Price, err := strconv.ParseFloat(message.Price, 64)
				if err == nil {
					f64Volume, err := strconv.ParseFloat(message.LastSize, 64)
					if err == nil {
						if message.TradeID != 0 {
							if message.Side == "sell" {
								f64Volume = -f64Volume
							}
							t := &dia.Trade{
								Symbol:         ps.pair.Symbol,
								Pair:           message.ProductID,
								Price:          f64Price,
								Volume:         f64Volume,
								Time:           message.Time.Time(),
								ForeignTradeID: strconv.FormatInt(int64(message.TradeID), 16),
								Source:         s.exchangeName,
							}
							ps.parent.chanTrades <- t
						}
					} else {
						log.Error("error parsing LastSize " + message.LastSize)
					}
				} else {
					log.Error("error parsing price " + message.Price)
				}
			} else {
				log.Error("unknown productError" + message.ProductID)
			}
		}
	}
	s.cleanup(err)
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *CoinBaseScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *CoinBaseScraper) Close() error {
	if s.closed {
		return errors.New("CoinBaseScraper: Already closed")
	}
	s.wsConn.Close()
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}
func (s *CoinBaseScraper) normalizeSymbol(foreignName string) (symbol string, err error) {
	str := strings.Split(foreignName, "-")
	symbol = str[0]
	if helpers.NameForSymbol(symbol) == symbol {
		return symbol, errors.New("Foreign name can not be normalized:" + foreignName + " symbol:" + symbol)
	}
	if helpers.SymbolIsBlackListed(symbol) {
		return symbol, errors.New("Symbol is black listed:" + symbol)
	}
	return symbol, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *CoinBaseScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	response, err := http.Get("https://api.pro.coinbase.com/products")
	if err != nil {
		log.Error("The HTTP request failed:", err)
		return
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	var ar []gdax.Product
	err = json.Unmarshal(data, &ar)
	if err == nil {
		for _, p := range ar {
			symbol, serr := s.normalizeSymbol(p.ID)
			if serr == nil {
				pairs = append(pairs, dia.Pair{
					Symbol:      symbol,
					ForeignName: p.ID,
					Exchange:    s.exchangeName,
				})
			} else {
				log.Error(serr)
			}
		}
	}
	return
}

// NewCoinBaseScraper implements PairScraper for GDax
type CoinBasePairScraper struct {
	parent     *CoinBaseScraper
	pair       dia.Pair
	closed     bool
	lastRecord int64
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *CoinBaseScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("CoinBaseScraper: Call ScrapePair on closed scraper")
	}
	ps := &CoinBasePairScraper{
		parent:     s,
		pair:       pair,
		lastRecord: 0, //TODO FIX to figure out the last we got...
	}

	s.pairScrapers[pair.ForeignName] = ps

	subscribe := gdax.Message{
		Type: "subscribe",
		Channels: []gdax.MessageChannel{
			gdax.MessageChannel{
				Name: ChannelHeartbeat,
				ProductIds: []string{
					pair.ForeignName,
				},
			},
			gdax.MessageChannel{
				Name: ChannelTicker,
				ProductIds: []string{
					pair.ForeignName,
				},
			},
		},
	}
	if err := s.wsConn.WriteJSON(subscribe); err != nil {
		println(err.Error())
	}

	return ps, nil
}

// Channel returns a channel that can be used to receive trades/pricing information
func (ps *CoinBaseScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

func (ps *CoinBasePairScraper) Close() error {
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *CoinBasePairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *CoinBasePairScraper) Pair() dia.Pair {
	return ps.pair
}
