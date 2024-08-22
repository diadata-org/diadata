package scrapers

import (
	"bytes"
	"compress/flate"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"time"

	ws "github.com/gorilla/websocket"
	"github.com/zekroTJA/timedmap"
	"go.uber.org/ratelimit"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
)

const (
	bitMartAPIEndpoint          = "https://api-cloud.bitmart.com/spot/v1"
	bitMartWSEndpoint           = "wss://ws-manager-compress.bitmart.com/api?protocol=1.1"
	bitMartSpotTradingSell      = "sell"
	bitMartSymbolsStatusActive  = "trading"
	bitMartPingMessage          = "ping"
	bitMartPongMessage          = "pong"
	bitMartWSSpotTradingTopic   = "spot/trade"
	bitMartWSOpSubscribe        = "subscribe"
	bitMartWSOpUnsubscribe      = "unsubscribe"
	bitMartRetryAttempts        = 15  // Max consecutive retry attempts until connection fail.
	bitMartPingInterval         = 15  // Number of seconds between ping messages.
	bitMartMaxConnections       = 10  // Numbers of connections per IP.
	bitMartMaxSubsPerConnection = 100 // Subscription limit for each connection.
)

type BitmartWsRequest struct {
	Op   string   `json:"op"`
	Args []string `json:"args"`
}

type BitmartHttpSymbolsDetailsResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Trace   string `json:"trace"`
	Data    struct {
		Symbols []struct {
			Symbol            string `json:"symbol"`
			SymbolId          int    `json:"symbol_id"`
			BaseCurrency      string `json:"base_currency"`
			QuoteCurrency     string `json:"quote_currency"`
			QuoteIncrement    string `json:"quote_increment"`
			BaseMinSize       string `json:"base_min_size"`
			PriceMinPrecision int    `json:"price_min_precision"`
			PriceMaxPrecision int    `json:"price_max_precision"`
			Expiration        string `json:"expiration"`
			MinBuyAmount      string `json:"min_buy_amount"`
			MinSellAmount     string `json:"min_sell_amount"`
			TradeStatus       string `json:"trade_status"`
		} `json:"symbols"`
	} `json:"data"`
}
type BitmartWsTradeResponse struct {
	Table string `json:"table"`
	Data  []struct {
		Symbol       string `json:"symbol"`
		Price        string `json:"price"`
		Side         string `json:"side"`
		Size         string `json:"size"`
		TimestampSec int    `json:"s_t"`
	} `json:"data"`
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    string `json:"errorCode"`
	Event        string `json:"event"`
}

type bitMartPairScraperSet map[*BitMartPairScraper]nothing

// BitMartScraper is a scraper for BitMart
type BitMartScraper struct {
	// the websocket connection to the BitMart API
	wsClient           []*ws.Conn
	errCount           []int
	countTopic         []int
	lastUsedConnection int
	rl                 ratelimit.Limiter
	listener           chan *BitmartWsTradeResponse
	// signaling channels for session initialization and finishing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; err should be read from error(), closed should be read from isClosed()
	// those two methods implement RW lock
	// only cleanup method should hold write lock
	closedMutex sync.RWMutex // don't need to lock on read
	closed      bool
	errMutex    sync.RWMutex
	err         error
	// used to keep track of trading pairs that we subscribed to
	// use sync.Maps to concurrently handle multiple pairs
	pairScrapers      sync.Map // dia.ExchangePair -> bitMartPairScraperSet
	pairSubscriptions sync.Map // dia.ExchangePair -> int (connection ID)
	exchangeName      string
	chanTrades        chan *dia.Trade
	db                *models.RelDB
}

// NewBitMartScraper returns a new BitMart scraper
func NewBitMartScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BitMartScraper {
	s := &BitMartScraper{
		wsClient:     make([]*ws.Conn, bitMartMaxConnections),
		errCount:     make([]int, bitMartMaxConnections),
		countTopic:   make([]int, bitMartMaxConnections),
		rl:           ratelimit.New(100, ratelimit.Per(10*time.Second)),
		listener:     make(chan *BitmartWsTradeResponse),
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		closed:       false,
		err:          nil,
		exchangeName: exchange.Name,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}
	for i := 0; i < bitMartMaxConnections; i++ {
		var wsDialer ws.Dialer
		wsConn, _, err := wsDialer.Dial(bitMartWSEndpoint, nil)
		if err != nil {
			println(err.Error())
		}
		s.wsClient[i] = wsConn
	}
	if scrape {
		go s.mainLoop()
	}
	return s
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *BitMartScraper) Close() error {
	if s.isClosed() {
		return errors.New("scraper already closed")
	}
	s.close()
	close(s.shutdown)
	for i := 0; i < bitMartMaxConnections; i++ {
		err := s.wsClient[i].Close()
		if err != nil {
			return fmt.Errorf("Close Error: %s", err.Error())
		}
	}
	<-s.shutdownDone
	return s.error()
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from the BitMart scraper
func (s *BitMartScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	if err := s.error(); err != nil {
		return nil, err
	}
	if s.isClosed() {
		return nil, errors.New("scraper is closed")
	}
	ps := &BitMartPairScraper{
		parent: s,
		pair:   pair,
	}
	pairScrapers, _ := s.pairScrapers.LoadOrStore(pair.ForeignName, bitMartPairScraperSet{})
	pairScrapers.(bitMartPairScraperSet)[ps] = nothing{}
	if _, ok := s.pairSubscriptions.Load(pair.ForeignName); !ok {
		wsIdx := s.lastUsedConnection
		if cIdx := s.lastUsedConnection; s.countTopic[cIdx] < bitMartMaxSubsPerConnection {
			wsIdx = cIdx
		} else {
			for i := 0; i < bitMartMaxConnections; i++ {
				if s.countTopic[i] < bitMartMaxSubsPerConnection {
					wsIdx = i
				}
			}
		}
		if err := s.subscribe(pair.ForeignName, wsIdx); err != nil {
			delete(pairScrapers.(bitMartPairScraperSet), ps)
			return nil, err
		}
		s.pairSubscriptions.Store(pair.ForeignName, wsIdx)
	} else {
		return nil, fmt.Errorf("pair %s already subscribed", pair.ForeignName)
	}
	return ps, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *BitMartScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	data, _, err := utils.GetRequest(bitMartAPIEndpoint + "/symbols/details")
	if err != nil {
		return
	}
	var res BitmartHttpSymbolsDetailsResponse
	err = json.Unmarshal(data, &res)
	if err == nil {
		for _, p := range res.Data.Symbols {
			if p.TradeStatus != bitMartSymbolsStatusActive {
				continue
			}
			pair, err := s.NormalizePair(dia.ExchangePair{
				Symbol:      p.BaseCurrency,
				ForeignName: p.Symbol,
				Exchange:    s.exchangeName,
			})
			if err != nil {
				return nil, err
			}
			pairs = append(pairs, pair)
		}
	}
	return pairs, nil
}

// Channel returns a channel that can be used to receive trades
func (s *BitMartScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// TODO: FillSymbolData adds the name to the asset underlying @symbol on BitMart
func (s *BitMartScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

// TODO: NormalizePair accounts for the par
func (s *BitMartScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// BitMartPairScraper implements PairScraper for BitMart
type BitMartPairScraper struct {
	parent *BitMartScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with the BitMart scraper
func (ps *BitMartPairScraper) Close() error {
	if ps.closed {
		return fmt.Errorf("pair %s already unsubscribed", ps.pair.ForeignName)
	}
	log.Infof("Closing %s pair scraper...", ps.pair.ForeignName)
	if err := ps.parent.error(); err != nil {
		return err
	}
	if err := ps.parent.unsubscribe(ps.pair.ForeignName); err != nil {
		return err
	}
	ps.parent.pairSubscriptions.Delete(ps.pair.ForeignName)
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed and nil otherwise
func (ps *BitMartPairScraper) Error() error {
	return ps.parent.error()
}

// Pair returns the pair this scraper is subscribed to
func (ps *BitMartPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

// runs in a goroutine until s is closed
func (s *BitMartScraper) mainLoop() {
	defer s.cleanup(nil)
	defer func() {
		log.Printf("Shutting down main loop...\n")
		return
	}()
	for i := 0; i < bitMartMaxConnections; i++ {
		go func(idx int) {
			defer func() {
				if a := recover(); a != nil {
					log.Errorf("Work routine end. Recover msg: %+v", a)
				}
			}()
			ticker := time.NewTicker(bitMartPingInterval * time.Second)
			defer ticker.Stop()
			for {
				<-ticker.C
				if s.isClosed() {
					return
				}
				s.rl.Take()
				if err := s.wsClient[idx].WriteMessage(ws.TextMessage, []byte(bitMartPingMessage)); err != nil {
					log.Errorf("Error sending ping: %s", err)
					return
				}
			}
		}(i)
		go func(idx int) {
			defer func() {
				if a := recover(); a != nil {
					log.Errorf("Receive routine end. Recover msg: %+v", a)
				}
			}()
			for {
				if s.wsClient[idx] != nil {
					msgType, msg, err := s.wsClient[idx].ReadMessage()
					if err != nil {
						if s.isClosed() || s.errCount[idx] > bitMartRetryAttempts {
							return
						}
						if err := s.retryConnection(idx); err != nil || ws.IsCloseError(err, ws.CloseAbnormalClosure) {
							return
						}
					}

					var output []byte
					switch msgType {
					case ws.BinaryMessage:
						reader := bytes.NewReader(msg)
						gzreader := flate.NewReader(reader)
						if err != nil {
							log.Error("flate reader: ", err)
							return
						}
						output, err = ioutil.ReadAll(gzreader)
						if err != nil {
							log.Error("read all: ", err)
							return
						}
					case ws.TextMessage:
						if string(msg) == bitMartPongMessage {
							s.errCount[idx] = 0
							return
						}
						output = msg
					}

					// Unmarshal output and forward to listener.
					var subResults BitmartWsTradeResponse
					if err := json.Unmarshal(output, &subResults); err != nil {
						log.Errorf("Response error at connection #%d, err=%s\n", idx, err.Error())
						s.setError(err)
						s.errCount[idx]++
						if err := s.retryConnection(idx); err != nil {
							return
						}
					}
					if subResults.ErrorCode != "" {
						log.Errorf("Error code %s at %s event: %s", subResults.ErrorCode, subResults.Event, subResults.ErrorMessage)
						s.errCount[idx]++
						continue
					}
					if subResults.Table == bitMartWSSpotTradingTopic {
						s.errCount[idx] = 0
						s.listener <- &subResults
					}
				}
			}
		}(i)
	}
	for {
		select {
		case response := <-s.listener:
			tmFalseDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
			tmDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)

			for _, data := range response.Data {
				var exchangepair dia.ExchangePair
				volume, _ := strconv.ParseFloat(data.Size, 64)
				if data.Side == bitMartSpotTradingSell {
					volume = -volume
				}
				price, _ := strconv.ParseFloat(data.Price, 64)
				timestamp := time.Unix(int64(data.TimestampSec), 0)
				symbol := strings.Split(data.Symbol, `_`)
				exchangepair, err := s.db.GetExchangePairCache(s.exchangeName, data.Symbol)
				if err != nil {
					log.Error(err)
				}
				t := &dia.Trade{
					Symbol:         symbol[0],
					Pair:           data.Symbol,
					Price:          price,
					Time:           timestamp,
					Volume:         volume,
					Source:         s.exchangeName,
					ForeignTradeID: fmt.Sprintf("%s_%d", data.Symbol, data.TimestampSec),
					VerifiedPair:   exchangepair.Verified,
					BaseToken:      exchangepair.UnderlyingPair.BaseToken,
					QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
				}

				discardTrade := t.IdentifyDuplicateFull(tmFalseDuplicateTrades, duplicateTradesMemory)
				if discardTrade {
					log.Warn("Identical trade already scraped: ", t)
					continue
				} else {
					t.IdentifyDuplicateTagset(tmDuplicateTrades, duplicateTradesMemory)
					s.chanTrades <- t
				}
			}
		case <-s.shutdown:
			return
		}
	}
}

func (s *BitMartScraper) retryConnection(idx int) error {
	s.countTopic[idx] = 0
	log.Errorf("Reconnecting connection #%d...\n", idx)
	wsConn, _, err := ws.DefaultDialer.Dial(bitMartWSEndpoint, nil)
	if err != nil {
		s.errCount[idx]++
		return err
	}
	s.wsClient[idx] = wsConn
	subs := make([]string, 0)
	s.pairSubscriptions.Range(func(k, v interface{}) bool {
		if v.(int) == idx {
			if foreignName := k.(string); foreignName != "" {
				subs = append(subs, foreignName)
			}
		}
		return true
	})
	for _, foreignName := range subs {
		if err = s.subscribe(foreignName, idx); err != nil {
			break
		}
	}
	if err != nil {
		log.Errorf("Recovering %d error., err=%s\n", idx, err.Error())
		s.setError(err)
		s.errCount[idx]++
		return err
	}
	log.Infof("Successfully reconnected connection %d., errCount=%d", idx, s.errCount[idx])
	return nil
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *BitMartScraper) cleanup(err error) {
	s.pairScrapers.Range(func(k, v interface{}) bool {
		for ps := range v.(bitMartPairScraperSet) {
			ps.closed = true
		}
		s.pairScrapers.Delete(k)
		return true
	})
	if err != nil {
		s.setError(err)
	}
	s.close()
	close(s.shutdownDone)
}

func (s *BitMartScraper) isClosed() bool {
	s.closedMutex.RLock()
	defer s.closedMutex.RUnlock()
	return s.closed
}

func (s *BitMartScraper) close() {
	s.closedMutex.Lock()
	defer s.closedMutex.Unlock()
	s.closed = true
}

func (s *BitMartScraper) error() error {
	s.errMutex.RLock()
	defer s.errMutex.RUnlock()
	return s.err
}

func (s *BitMartScraper) setError(err error) {
	s.errMutex.Lock()
	defer s.errMutex.Unlock()
	s.err = err
}

func (s *BitMartScraper) subscribe(foreignName string, id int) error {
	topic := fmt.Sprintf("%s:%s", bitMartWSSpotTradingTopic, foreignName)
	s.rl.Take()
	if err := s.wsClient[id].WriteJSON(BitmartWsRequest{
		Op:   bitMartWSOpSubscribe,
		Args: []string{topic},
	}); err != nil {
		return err
	}
	s.countTopic[id]++
	if s.lastUsedConnection == bitMartMaxConnections-1 {
		s.lastUsedConnection = 0
	} else {
		s.lastUsedConnection++
	}
	return nil
}

func (s *BitMartScraper) unsubscribe(foreignName string) error {
	if id, ok := s.pairSubscriptions.Load(foreignName); ok {
		if err := s.wsClient[id.(int)].WriteJSON(BitmartWsRequest{
			Op:   bitMartWSOpUnsubscribe,
			Args: []string{fmt.Sprintf("%s:%s", bitMartWSSpotTradingTopic, foreignName)},
		}); err != nil {
			return err
		}
		s.pairSubscriptions.Delete(foreignName)
		s.countTopic[id.(int)]--
		return nil
	} else {
		return errors.New("subscription id not found")
	}
}
