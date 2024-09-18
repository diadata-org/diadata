package scrapers

import (
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	ws "github.com/gorilla/websocket"
	"github.com/zekroTJA/timedmap"
)

const (
	bitgetWsAPI        = "wss://ws.bitget.com/v2/ws/public"
	bitgetPingInterval = 30
)

type bitgetSubscribeMessage struct {
	Operation string                     `json:"op"`
	Arguments []bitgetSubscribeArguments `json:"args"`
}

type bitgetSubscribeArguments struct {
	InstrumentType string `json:"instType"`
	Channel        string `json:"channel"`
	InstrumentID   string `json:"instId"`
}

type bitgetWsResponse struct {
	Action    string         `json:"action"`
	Argument  bitgetArgument `json:"arg"`
	Data      []bitgetData   `json:"data"`
	Timestamp int64          `json:"ts"`
}

type bitgetArgument struct {
	InstrumentType string `json:"instType"`
	Channel        string `json:"channel"`
	InstrumentID   string `json:"instId"`
}

type bitgetData struct {
	Timestamp      string `json:"ts"`
	Price          string `json:"price"`
	Volume         string `json:"size"`
	Side           string `json:"side"`
	ForeignTradeID string `json:"tradeId"`
}

type BitgetScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock    sync.RWMutex
	error        error
	closed       bool
	pairScrapers map[string]*BitgetPairScraper // pc.ExchangePair -> pairScraperSet
	wsConn       *ws.Conn
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB
}

// NewBitgetScraper returns a new BitgetScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func NewBitgetScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BitgetScraper {
	s := &BitgetScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*BitgetPairScraper),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}
	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(bitgetWsAPI, nil)
	if err != nil {
		log.Errorf("Dial websocket api: %s", err.Error())
	}

	go s.pingRoutine(time.Duration(bitgetPingInterval * time.Second))

	s.wsConn = SwConn
	if scrape {
		go s.mainLoop()
	}
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *BitgetScraper) mainLoop() {

	tmFalseDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
	tmDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
	time.Sleep(5 * time.Second)

	for {

		// Check if we get a pong message back.
		_, p, err := s.wsConn.ReadMessage()
		if err != nil {
			log.Error("ReadMessage: ", err)
		} else {
			if strings.Contains(string(p), "pong") || strings.Contains(string(p), "ng") {
				log.Infof("got  %s", string(p))
			}
		}

		var message bitgetWsResponse
		if err = s.wsConn.ReadJSON(&message); err != nil {
			log.Errorf("ReadJSON: %s", err.Error())
			log.Info("instead of pong got ", string(p))
			if strings.Contains(err.Error(), "invalid character") {
				continue
			}
			return
		}

		ps, ok := s.pairScrapers[message.Argument.InstrumentID]
		if ok && message.Action != "snapshot" {
			for _, data := range message.Data {
				var f64Price float64
				var f64Volume float64
				var exchangepair dia.ExchangePair
				f64Price, err = strconv.ParseFloat(data.Price, 64)
				if err != nil {
					log.Error("error parsing price " + data.Price)
				}
				f64Volume, err = strconv.ParseFloat(data.Volume, 64)
				if err != nil {
					log.Error("error parsing volume " + data.Volume)
				}

				if data.Side != "buy" {
					f64Volume = -f64Volume
				}

				timestamp, err := strconv.ParseInt(data.Timestamp, 10, 64)
				if err != nil {
					log.Error("Parse timestamp: ", err)
				}

				exchangepair, err = s.db.GetExchangePairCache(s.exchangeName, message.Argument.InstrumentID)
				if err != nil {
					// log.Error("get exchangepair from cache: ", err)
				}
				t := dia.Trade{
					Symbol:         ps.pair.Symbol,
					Pair:           message.Argument.InstrumentID,
					Price:          f64Price,
					Volume:         f64Volume,
					Time:           time.Unix(0, timestamp*1e6),
					ForeignTradeID: data.ForeignTradeID,
					Source:         s.exchangeName,
					VerifiedPair:   exchangepair.Verified,
					BaseToken:      exchangepair.UnderlyingPair.BaseToken,
					QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
				}
				if t.VerifiedPair {
					log.Info("got verified trade: ", t)
				} else {
					log.Infof("got trade at %v : %s -- %v -- %v", t.Time, t.Pair, t.Price, t.Volume)
				}
				// Handle duplicate trades.
				discardTrade := t.IdentifyDuplicateFull(tmFalseDuplicateTrades, duplicateTradesMemory)
				if !discardTrade {
					t.IdentifyDuplicateTagset(tmDuplicateTrades, duplicateTradesMemory)
					ps.parent.chanTrades <- &t
				}
			}

		}
	}

}

func (s *BitgetScraper) cleanup(err error) {
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
func (s *BitgetScraper) Close() error {
	if s.closed {
		return errors.New("BitgetScraper: Already closed")
	}
	err := s.wsConn.Close()
	if err != nil {
		log.Error(err)
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (s *BitgetScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

func (s *BitgetScraper) pingRoutine(d time.Duration) {
	ticker := time.NewTicker(d)
	for range ticker.C {
		if err := s.wsConn.WriteMessage(ws.TextMessage, []byte("ping")); err != nil {
			log.Errorf("send ping: %s.", err.Error())
		} else {
			log.Info("sent ping.")
		}
	}
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *BitgetScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {

	// data, _, err := utils.GetRequest("https://api.pro.coinbase.com/products")
	// if err != nil {
	// 	return
	// }

	// err = json.Unmarshal(data, &ar)
	// if err == nil {
	// 	for _, p := range ar {
	// 		pairToNormalise := dia.ExchangePair{
	// 			Symbol:      p.BaseCurrency,
	// 			ForeignName: p.ID,
	// 			Exchange:    s.exchangeName,
	// 		}
	// 		pair, serr := s.NormalizePair(pairToNormalise)
	// 		if serr == nil {
	// 			pairs = append(pairs, pair)
	// 		} else {
	// 			log.Error(serr)
	// 		}
	// 	}
	// }
	return
}

// FillSymbolData collects all available information on an asset traded on Bitget
func (s *BitgetScraper) FillSymbolData(symbol string) (asset dia.Asset, err error) {
	asset.Symbol = symbol
	return asset, nil
}

// BitgetPairScraper implements PairScraper
type BitgetPairScraper struct {
	parent     *BitgetScraper
	pair       dia.ExchangePair
	closed     bool
	lastRecord int64
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *BitgetScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("BitgetScraper: Call ScrapePair on closed scraper")
	}
	ps := &BitgetPairScraper{
		parent:     s,
		pair:       pair,
		lastRecord: 0,
	}

	s.pairScrapers[pair.ForeignName] = ps

	subscribeMessage := bitgetSubscribeMessage{
		Operation: "subscribe",
		Arguments: []bitgetSubscribeArguments{
			{
				InstrumentType: "SPOT",
				Channel:        "trade",
				InstrumentID:   pair.ForeignName,
			},
		},
	}
	if err := s.wsConn.WriteJSON(subscribeMessage); err != nil {
		println(err.Error())
	}
	log.Info("subscribed to: ", pair.ForeignName)

	return ps, nil
}

// Channel returns a channel that can be used to receive trades/pricing information
func (ps *BitgetScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

func (ps *BitgetPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *BitgetPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *BitgetPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
