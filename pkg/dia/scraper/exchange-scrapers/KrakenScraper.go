package scrapers

import (
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
	"github.com/zekroTJA/timedmap"
)

const (
	krakenRefreshDelay  = time.Second * 30 * 1
	krakenMaxSubPerConn = 90
)

var (
	krakenWSBaseString = utils.Getenv("KRAKEN_WS_BASE_STRING", "wss://ws.kraken.com/v2")
)

type KrakenScraper struct {
	connections map[int]KrakenWSConnection
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock    sync.RWMutex
	error        error
	closed       bool
	pairScrapers map[string]*KrakenPairScraper // pc.ExchangePair -> pairScraperSet

	ticker       *time.Ticker
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB
}

type KrakenWSConnection struct {
	wsConn           *ws.Conn
	numSubscriptions int
}

type krakenWSSubscribeMessage struct {
	Method string       `json:"method"`
	Params krakenParams `json:"params"`
}

type krakenParams struct {
	Channel  string   `json:"channel"`
	Symbol   []string `json:"symbol"`
	Snapshot bool     `json:"snapshot"`
}

type krakenWSResponse struct {
	Channel string                 `json:"channel"`
	Type    string                 `json:"type"`
	Data    []krakenWSResponseData `json:"data"`
}

type krakenWSResponseData struct {
	Symbol    string  `json:"symbol"`
	Side      string  `json:"side"`
	Price     float64 `json:"price"`
	Size      float64 `json:"qty"`
	OrderType string  `json:"order_type"`
	TradeID   int     `json:"trade_id"`
	Time      string  `json:"timestamp"`
}

// NewKrakenScraper returns a new KrakenScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func NewKrakenScraper(key string, secret string, exchange dia.Exchange, scrape bool, relDB *models.RelDB) *KrakenScraper {

	s := &KrakenScraper{
		connections:  make(map[int]KrakenWSConnection),
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*KrakenPairScraper),
		ticker:       time.NewTicker(krakenRefreshDelay),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}
	err := s.newConn()
	if err != nil {
		log.Fatal("newConn: ", err)
	}

	if scrape {
		go s.mainLoop()
	}
	return s
}

func (s *KrakenScraper) mainLoop() {
	// Wait for subscription to all pairs.
	time.Sleep(5 * time.Second)
	for _, c := range s.connections {
		go s.subLoop(c.wsConn)
	}
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *KrakenScraper) subLoop(client *ws.Conn) {
	tmFalseDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
	tmDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)

	for {
		var message krakenWSResponse
		err := client.ReadJSON(&message)
		if err != nil {
			log.Error("JSON decode error: ", err)
			continue
		}

		if message.Channel == "trade" && message.Type == "update" {
			for _, data := range message.Data {

				symbols := strings.Split(data.Symbol, "/")
				if len(symbols) != 2 {
					continue
				}

				exchangePair, err := s.db.GetExchangePairCache(s.exchangeName, symbols[0]+symbols[1])
				if err != nil {
					log.Error("get exchange pair from cache: ", err)
				}

				price, volume, timestamp, foreignTradeID, err := parseKrakenTradeMessage(data)
				if err != nil {
					log.Error("parseKrakenTradeMessage: ", err)
					continue
				}
				trade := dia.Trade{
					Symbol:         symbols[0],
					Pair:           data.Symbol,
					QuoteToken:     exchangePair.UnderlyingPair.QuoteToken,
					BaseToken:      exchangePair.UnderlyingPair.BaseToken,
					Price:          price,
					Volume:         volume,
					Time:           timestamp,
					ForeignTradeID: foreignTradeID,
					Source:         s.exchangeName,
					VerifiedPair:   exchangePair.Verified,
				}

				// Handle duplicate trades.
				discardTrade := trade.IdentifyDuplicateFull(tmFalseDuplicateTrades, duplicateTradesMemory)
				if !discardTrade {
					trade.IdentifyDuplicateTagset(tmDuplicateTrades, duplicateTradesMemory)
					log.Info("got trade: ", trade)
					s.chanTrades <- &trade
				}
			}

		}
	}
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *KrakenScraper) cleanup(err error) {

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
func (s *KrakenScraper) Close() error {
	if s.closed {
		return errors.New("KrakenScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// KrakenPairScraper implements PairScraper for Kraken
type KrakenPairScraper struct {
	parent *KrakenScraper
	pair   dia.ExchangePair
	closed bool
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *KrakenScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("KrakenScraper: Call ScrapePair on closed scraper")
	}
	ps := &KrakenPairScraper{
		parent: s,
		pair:   pair,
	}

	s.pairScrapers[pair.Symbol] = ps
	err := s.subscribe(pair, true)
	if err != nil {
		return ps, err
	}

	return ps, nil
}

func (s *KrakenScraper) subscribe(pair dia.ExchangePair, subscribe bool) error {
	id := len(s.connections)

	subscribeType := "unsubscribe"
	if subscribe {
		subscribeType = "subscribe"
	}
	a := &krakenWSSubscribeMessage{
		Method: subscribeType,
		Params: krakenParams{
			Channel:  "trade",
			Symbol:   []string{pair.Symbol + "/" + strings.TrimPrefix(pair.ForeignName, pair.Symbol)},
			Snapshot: false,
		},
	}
	if s.connections[id-1].numSubscriptions < krakenMaxSubPerConn {
		if err := s.connections[id-1].wsConn.WriteJSON(a); err != nil {
			return err
		}
		log.Info("subscribed to: ", a.Params.Symbol)
		conn := s.connections[id-1]
		conn.numSubscriptions++
		s.connections[id-1] = conn
	} else {
		err := s.newConn()
		if err != nil {
			return err
		}
		log.Infof("make new connection. Total number of connections: %v", len(s.connections))
		id++
		if err := s.connections[id-1].wsConn.WriteJSON(a); err != nil {
			return err
		}
		log.Info("subscribed to: ", a.Params.Symbol)
		conn := s.connections[id-1]
		conn.numSubscriptions++
		s.connections[id-1] = conn
	}

	return nil
}

func (s *KrakenScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *KrakenScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return []dia.ExchangePair{}, errors.New("FetchAvailablePairs() not implemented")
}

func (s *KrakenScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// Add a connection to the connection pool.
func (s *KrakenScraper) newConn() error {
	var wsDialer ws.Dialer
	wsConn, _, err := wsDialer.Dial(krakenWSBaseString, nil)
	if err != nil {
		return err
	}
	s.connections[len(s.connections)] = KrakenWSConnection{wsConn: wsConn, numSubscriptions: 0}
	return nil
}

// Channel returns a channel that can be used to receive trades/pricing information
func (ps *KrakenScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

func (ps *KrakenPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *KrakenPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *KrakenPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

func parseKrakenTradeMessage(message krakenWSResponseData) (
	price float64,
	volume float64,
	timestamp time.Time,
	foreignTradeID string,
	err error,
) {
	price = message.Price
	volume = message.Size
	if message.Side == "sell" {
		volume -= 1
	}
	timestamp, err = time.Parse("2006-01-02T15:04:05.000000Z", message.Time)
	if err != nil {
		return
	}

	foreignTradeID = strconv.Itoa(message.TradeID)
	return
}
