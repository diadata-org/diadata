package scrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
)

const (
	bitForexAPIEndpoint         = "https://api.bitforex.com/api/v1"
	bitForexWSEndpoint          = "wss://www.bitforex.com/mkapi/coinGroup1/ws"
	bitForexSpotTradingBuy      = 1
	bitForexPingMessage         = "ping_p"
	bitForexPongMessage         = "pong_p"
	bitForexInitialTradeReqSize = 10
	bitForexWSBucketSize        = 50

	// bitForexConnMaxRetry is a max retry value used when retrying to create a new connection
	bitForexConnMaxRetry = 50
)

// bitForexWSRequest is a websocket request
type bitForexWSRequest struct {
	Type  string                 `json:"type"`
	Event string                 `json:"event"`
	Param bitForexWSRequestParam `json:"param"`
}

// bitForexWSRequestParams is a websocket request param
type bitForexWSRequestParam struct {
	BusinessType string `json:"businessType"`
	Size         int    `json:"size"`
}

// bitForexWSResponse is a websocket response
type bitForexWSResponse struct {
	Success *bool                  `json:"success"`
	Event   string                 `json:"event"`
	Code    int                    `json:"code"`
	Param   bitForexWSRequestParam `json:"param"`
	Data    json.RawMessage        `json:"data"`
}

// bitForexTradeResult is a trade result coming from websocket
type bitForexTradeResult struct {
	Price     float64 `json:"price"`
	Amount    float64 `json:"amount"`
	Direction int     `json:"direction"`
	TradeID   string  `json:"tid"`
	Timestamp int     `json:"time"`
}

// bitForexSymbolResponse is an API response for retrieving instruments
type bitForexSymbolResponse struct {
	Success bool             `json:"success"`
	Code    int              `json:"code"`
	Data    []bitForexSymbol `json:"data"`
	Message string           `json:"message"`
}

// bitForexSymbol is a symbol data from api response
type bitForexSymbol struct {
	Symbol          string  `json:"symbol"`
	PricePrecision  int     `json:"pricePrecision"`
	AmountPrecision int     `json:"amountPrecision"`
	MinOrderAmount  float64 `json:"minOrderAmount"`
}

// BitforexScraper is a scraper for Crypto.com
type BitforexScraper struct {
	ws *ws.Conn

	// signaling channels for session initialization and finishing
	shutdown           chan nothing
	shutdownDone       chan nothing
	signalShutdown     sync.Once
	signalShutdownDone sync.Once

	// error handling; err should be read from error(), closed should be read from isClosed()
	// those two methods implement RW lock
	errMutex    sync.RWMutex
	err         error
	closedMutex sync.RWMutex
	closed      bool

	// used to keep track of trading pairs that we subscribed to
	pairScrapers sync.Map
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB

	// used to handle connection retry
	connMutex      sync.Mutex
	connRetryCount int
}

// NewBitforexScraper returns a new Crypto.com scraper
func NewBitforexScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BitforexScraper {
	s := &BitforexScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		exchangeName: exchange.Name,
		err:          nil,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}

	if err := s.newConn(); err != nil {
		log.Error("s.newConn()", err)

		return nil
	}

	if scrape {
		go s.mainLoop()
	}

	return s
}

// Close unsubscribes data and closes any existing WebSocket connections, as well as channels of BitforexScraper
func (s *BitforexScraper) Close() error {
	if s.isClosed() {
		return errors.New("BitforexScraper: Already closed")
	}

	s.signalShutdown.Do(func() {
		close(s.shutdown)
	})

	<-s.shutdownDone

	return s.error()
}

// Channel returns a channel that can be used to receive trades
func (s *BitforexScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// FetchAvailablePairs returns all traded pairs on Crypto.com
func (s *BitforexScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	data, _, err := utils.GetRequest(bitForexAPIEndpoint + "/market/symbols")
	if err != nil {
		return nil, err
	}

	var res bitForexSymbolResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	if !res.Success {
		log.Errorf("BitforexScraper: Fetching pairs error=%s with code=%d", res.Message, res.Code)

		return nil, fmt.Errorf("BitforexScraper: Getting available pairs error with code %d", res.Code)
	}

	for _, symbol := range res.Data {
		baseCurrency, foreignName := s.extractSymbol(symbol.Symbol)
		pairs = append(pairs, dia.ExchangePair{
			Symbol:      baseCurrency,
			ForeignName: foreignName,
			Exchange:    s.exchangeName,
		})
	}

	return pairs, nil
}

// FillSymbolData adds the name to the asset underlying @symbol on Crypto.com
func (s *BitforexScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *BitforexScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from the Crypto.com scraper
func (s *BitforexScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	if err := s.error(); err != nil {
		return nil, err
	}
	if s.isClosed() {
		return nil, errors.New("BitforexScraper: Call ScrapePair on closed scraper")
	}

	ps := &BitforexPairScraper{
		parent: s,
		pair:   pair,
	}
	if err := s.subscribe([]dia.ExchangePair{pair}); err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *BitforexScraper) mainLoop() {
	defer s.cleanup()
	defer func() {
		// Signal to other subscribers that the scraper is internally shut down.
		go s.signalShutdown.Do(func() {
			close(s.shutdown)
		})
	}()

	go s.ping()

	for {
		select {
		case <-s.shutdown:
			log.Info("BitforexScraper: Shutting down main loop")
			return
		default:
		}

		_, msg, err := s.wsConn().ReadMessage()
		if err != nil {
			log.Warnf("BitforexScraper: Creating a new connection caused by err=%s", err.Error())

			if retryErr := s.retryConnection(); retryErr != nil {
				s.setError(retryErr)
				log.Errorf("BitforexScraper: Shutting down main loop after retrying to create a new connection, err=%s", retryErr.Error())

				return
			}

			log.Info("BitforexScraper: Successfully created a new connection")

			continue
		}

		if string(msg) == bitForexPongMessage {
			continue
		}

		var res bitForexWSResponse
		if err := json.Unmarshal(msg, &res); err != nil {
			s.setError(err)
			log.Errorf("BitforexScraper: Shutting down main loop due to unmarshalling failure, err=%s", err.Error())

			return
		}
		if res.Success != nil && !*res.Success {
			err := fmt.Errorf("BitforexScraper: Websocket error, code=%d, event=%s", res.Code, res.Event)
			msg := fmt.Sprintf("BitforexScraper: Shutting down main loop due to non-retryable, err=%s", err.Error())
			s.setError(err)
			log.Error(msg)

			return
		}

		switch res.Event {
		case "trade":
			baseCurrency, foreignName := s.extractSymbol(res.Param.BusinessType)
			pair, err := s.db.GetExchangePairCache(s.exchangeName, foreignName)
			if err != nil {
				log.Error("GetExchangePairCache", err)
			}

			var trades []bitForexTradeResult
			if err := json.Unmarshal(res.Data, &trades); err != nil {
				s.setError(err)
				log.Errorf("BitforexScraper: Shutting down main loop due to response unmarshaling failure, err=%s", err.Error())

			}

			for _, trade := range trades {
				volume := trade.Amount
				if trade.Direction != bitForexSpotTradingBuy {
					volume = -volume
				}

				trade := &dia.Trade{
					Symbol:         baseCurrency,
					Pair:           foreignName,
					Price:          trade.Price,
					Time:           time.Unix(int64(trade.Timestamp/1000), 0),
					Volume:         volume,
					Source:         s.exchangeName,
					ForeignTradeID: trade.TradeID,
					VerifiedPair:   pair.Verified,
					BaseToken:      pair.UnderlyingPair.BaseToken,
					QuoteToken:     pair.UnderlyingPair.QuoteToken,
				}

				if pair.Verified {
					log.Infoln("Got verified trade", trade)
				}

				select {
				case <-s.shutdown:
				case s.chanTrades <- trade:
				}
			}
		}
	}
}

func (s *BitforexScraper) toBitforexSymbol(foreignName string) string {
	return strings.ToLower("coin-" + foreignName)
}

func (s *BitforexScraper) extractSymbol(symbol string) (baseCurrency, foreignName string) {
	ss := strings.SplitN(symbol, "-", 3)
	baseCurrency = strings.ToUpper(ss[2])
	foreignName = strings.ToUpper(ss[2] + "-" + ss[1])

	return baseCurrency, foreignName
}

func (s *BitforexScraper) newConn() error {
	conn, _, err := ws.DefaultDialer.Dial(bitForexWSEndpoint, nil)
	if err != nil {
		return err
	}

	defer s.connMutex.Unlock()
	s.connMutex.Lock()
	s.ws = conn

	return nil
}

func (s *BitforexScraper) wsConn() *ws.Conn {
	s.connMutex.Lock()
	defer s.connMutex.Unlock()

	return s.ws
}

func (s *BitforexScraper) ping() {
	t := time.NewTicker(time.Duration(15) * time.Second)
	for {
		select {
		case <-s.shutdown:
			log.Info("BitforexScraper: Shutting down ping loop")

			return
		case <-t.C:
			if err := s.sendMessage([]byte(bitForexPingMessage)); err != nil {
				log.Warningf("BitforexScraper: Sending pings fail %s", err.Error())
			}
		}
	}
}

func (s *BitforexScraper) cleanup() {
	if err := s.wsConn().Close(); err != nil {
		s.setError(err)
	}

	close(s.chanTrades)
	s.close()
	s.signalShutdownDone.Do(func() {
		close(s.shutdownDone)
	})
}

func (s *BitforexScraper) error() error {
	s.errMutex.RLock()
	defer s.errMutex.RUnlock()

	return s.err
}

func (s *BitforexScraper) setError(err error) {
	s.errMutex.Lock()
	defer s.errMutex.Unlock()

	s.err = err
}

func (s *BitforexScraper) isClosed() bool {
	s.closedMutex.RLock()
	defer s.closedMutex.RUnlock()

	return s.closed
}

func (s *BitforexScraper) close() {
	s.closedMutex.Lock()
	defer s.closedMutex.Unlock()

	s.closed = true
}

func (s *BitforexScraper) subscribe(pairs []dia.ExchangePair) error {
	requests := make([]bitForexWSRequest, 0, len(pairs))
	for _, pair := range pairs {
		requests = append(requests, bitForexWSRequest{
			Type:  "subHq",
			Event: "trade",
			Param: bitForexWSRequestParam{
				BusinessType: s.toBitforexSymbol(pair.ForeignName),
				Size:         bitForexInitialTradeReqSize,
			},
		})
		s.pairScrapers.Store(pair.ForeignName, pair)
	}

	return s.send(requests)
}

func (s *BitforexScraper) unsubscribe(pairs []dia.ExchangePair) error {
	requests := make([]bitForexWSRequest, 0, len(pairs))
	for _, pair := range pairs {
		requests = append(requests, bitForexWSRequest{
			Type:  "subHq_cancel",
			Event: "trade",
			Param: bitForexWSRequestParam{
				BusinessType: s.toBitforexSymbol(pair.ForeignName),
				Size:         bitForexInitialTradeReqSize,
			},
		})
		s.pairScrapers.Delete(pair.ForeignName)
	}

	return s.send(requests)
}

func (s *BitforexScraper) retryConnection() error {
	s.connRetryCount += 1
	if s.connRetryCount > bitForexConnMaxRetry {
		return errors.New("BitforexPairScraper: Reached max retry connection")
	}
	if err := s.wsConn().Close(); err != nil {
		return err
	}
	if err := s.newConn(); err != nil {
		return err
	}

	var pairs []dia.ExchangePair
	s.pairScrapers.Range(func(key, value interface{}) bool {
		pair := value.(dia.ExchangePair)
		pairs = append(pairs, pair)
		return true
	})
	if err := s.subscribe(pairs); err != nil {
		return err
	}

	return nil
}

// divideIntoBuckets divides a []bitForexWSRequest slice into multiple buckets
func (s *BitforexScraper) divideIntoBuckets(requests []bitForexWSRequest, bucketSize int) [][]bitForexWSRequest {
	var bucket []bitForexWSRequest
	var buckets [][]bitForexWSRequest
	count := 1
	for _, request := range requests {
		bucket = append(bucket, request)
		if count%bucketSize == 0 {
			buckets = append(buckets, bucket)
			bucket = []bitForexWSRequest{}
		}

		count++
	}

	if len(bucket) > 0 {
		buckets = append(buckets, bucket)
	}

	return buckets
}

func (s *BitforexScraper) send(requests []bitForexWSRequest) error {
	buckets := s.divideIntoBuckets(requests, bitForexWSBucketSize)
	for _, bucket := range buckets {
		err := s.wsConn().WriteJSON(bucket)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *BitforexScraper) sendMessage(msg []byte) error {
	return s.wsConn().WriteMessage(ws.TextMessage, msg)
}

// BitforexPairScraper implements PairScraper for Crypto.com
type BitforexPairScraper struct {
	parent *BitforexScraper
	pair   dia.ExchangePair
	closed bool
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (p *BitforexPairScraper) Error() error {
	return p.parent.error()
}

// Pair returns the pair this scraper is subscribed to
func (p *BitforexPairScraper) Pair() dia.ExchangePair {
	return p.pair
}

// Close stops listening for trades of the pair associated with the Crypto.com scraper
func (p *BitforexPairScraper) Close() error {
	if err := p.parent.error(); err != nil {
		return err
	}
	if p.closed {
		return errors.New("BitforexPairScraper: Already closed")
	}
	if err := p.parent.unsubscribe([]dia.ExchangePair{p.pair}); err != nil {
		return err
	}

	p.closed = true

	return nil
}
