package scrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	ws "github.com/gorilla/websocket"
	"go.uber.org/ratelimit"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
)

const (
	cryptoDotComAPIEndpoint    = "https://api.crypto.com/v2"
	cryptoDotComWSEndpoint     = "wss://stream.crypto.com/v2/market"
	cryptoDotComSpotTradingBuy = "BUY"

	// cryptoDotComWSRateLimitPerSec is a max request per second for sending websocket requests
	cryptoDotComWSRateLimitPerSec = 10

	// cryptoDotComTaskMaxRetry is a max retry value used when retrying subscribe/unsubscribe trades
	cryptoDotComTaskMaxRetry = 20

	// cryptoDotComConnMaxRetry is a max retry value used when retrying to create a new connection
	cryptoDotComConnMaxRetry = 50

	// cryptoDotComRateLimitError is a rate limit error code
	cryptoDotComRateLimitError = 10006
)

// cryptoDotComWSTask is a websocket task tracking subscription/unsubscription
type cryptoDotComWSTask struct {
	Method     string
	Params     cryptoDotComWSRequestParams
	RetryCount int
}

func (c *cryptoDotComWSTask) toString() string {
	return fmt.Sprintf("method=%s, param=%s, retry=%d", c.Method, c.Params.toString(), c.RetryCount)
}

// cryptoDotComWSRequest is a websocket request
type cryptoDotComWSRequest struct {
	ID     int                         `json:"id"`
	Method string                      `json:"method"`
	Params cryptoDotComWSRequestParams `json:"params,omitempty"`
	Nonce  int64                       `json:"nonce,omitempty"`
}

// cryptoDotComWSRequestParams is a websocket request param
type cryptoDotComWSRequestParams struct {
	Channels []string `json:"channels"`
}

func (c *cryptoDotComWSRequestParams) toString() string {
	length := len(c.Channels)
	if length == 1 {
		return c.Channels[0]
	}
	if length > 1 {
		return fmt.Sprintf("%s +%d more", c.Channels[0], length-1)
	}

	return ""
}

// cryptoDotComWSResponse is a websocket response
type cryptoDotComWSResponse struct {
	ID     int             `json:"id"`
	Code   int             `json:"code"`
	Method string          `json:"method"`
	Result json.RawMessage `json:"result"`
}

// cryptoDotComWSSubscriptionResult is a trade result coming from websocket
type cryptoDotComWSSubscriptionResult struct {
	InstrumentName string            `json:"instrument_name"`
	Subscription   string            `json:"subscription"`
	Channel        string            `json:"channel"`
	Data           []json.RawMessage `json:"data"`
}

// cryptoDotComWSInstrument represents a trade
type cryptoDotComWSInstrument struct {
	Price     float64 `json:"p"`
	Quantity  float64 `json:"q"`
	Side      string  `json:"s"`
	TradeID   int     `json:"d"`
	TradeTime int64   `json:"t"`
}

// cryptoDotComInstrument represents a trading pair
type cryptoDotComInstrument struct {
	InstrumentName          string `json:"instrument_name"`
	QuoteCurrency           string `json:"quote_currency"`
	BaseCurrency            string `json:"base_currency"`
	PriceDecimals           int    `json:"price_decimals"`
	QuantityDecimals        int    `json:"quantity_decimals"`
	MarginTradingEnabled    bool   `json:"margin_trading_enabled"`
	MarginTradingEnabled5x  bool   `json:"margin_trading_enabled_5x"`
	MarginTradingEnabled10x bool   `json:"margin_trading_enabled_10x"`
	MaxQuantity             string `json:"max_quantity"`
	MinQuantity             string `json:"min_quantity"`
}

// cryptoDotComInstrumentResponse is an API response for retrieving instruments
type cryptoDotComInstrumentResponse struct {
	Code   int `json:"code"`
	Result struct {
		Instruments []cryptoDotComInstrument `json:"instruments"`
	} `json:"result"`
}

// CryptoDotComScraper is a scraper for Crypto.com
type CryptoDotComScraper struct {
	ws *ws.Conn
	rl ratelimit.Limiter

	// signaling channels for session initialization and finishing
	shutdown           chan nothing
	shutdownDone       chan nothing
	signalShutdown     sync.Once
	signalShutdownDone sync.Once

	// error handling; err should be read from error(), closed should be read from isClosed()
	// those two methods implement RW lock
	errMutex            sync.RWMutex
	err                 error
	closedMutex         sync.RWMutex
	closed              bool
	consecutiveErrCount int

	// used to keep track of trading pairs that we subscribed to
	pairScrapers sync.Map
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB
	taskCount    int32
	tasks        sync.Map

	// used to handle connection retry
	connMutex      sync.RWMutex
	connRetryCount int
}

// NewCryptoDotComScraper returns a new Crypto.com scraper
func NewCryptoDotComScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *CryptoDotComScraper {
	s := &CryptoDotComScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		exchangeName: exchange.Name,
		err:          nil,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}

	if err := s.newConn(); err != nil {
		log.Error(err)

		return nil
	}

	s.rl = ratelimit.New(cryptoDotComWSRateLimitPerSec)

	if scrape {
		go s.mainLoop()
	}

	return s
}

// Close unsubscribes data and closes any existing WebSocket connections, as well as channels of CryptoDotComScraper
func (s *CryptoDotComScraper) Close() error {
	if s.isClosed() {
		return errors.New("CryptoDotComScraper: Already closed")
	}

	s.signalShutdown.Do(func() {
		close(s.shutdown)
	})

	<-s.shutdownDone

	return s.error()
}

// Channel returns a channel that can be used to receive trades
func (s *CryptoDotComScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// FetchAvailablePairs returns all traded pairs on Crypto.com
func (s *CryptoDotComScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	data, _, err := utils.GetRequest(cryptoDotComAPIEndpoint + "/public/get-instruments")
	if err != nil {
		return nil, err
	}

	var res cryptoDotComInstrumentResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	if res.Code != 0 {
		return nil, fmt.Errorf("CryptoDotComScraper: Getting available pairs error with code %d", res.Code)
	}

	for _, i := range res.Result.Instruments {
		pairs = append(pairs, dia.ExchangePair{
			Symbol:      i.BaseCurrency,
			ForeignName: i.InstrumentName,
			Exchange:    s.exchangeName,
		})
	}

	return pairs, nil
}

// FillSymbolData adds the name to the asset underlying @symbol on Crypto.com
func (s *CryptoDotComScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *CryptoDotComScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from the Crypto.com scraper
func (s *CryptoDotComScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	if err := s.error(); err != nil {
		return nil, err
	}
	if s.isClosed() {
		return nil, errors.New("CryptoDotComScraper: Call ScrapePair on closed scraper")
	}

	ps := &CryptoDotComPairScraper{
		parent: s,
		pair:   pair,
	}
	if err := s.subscribe([]dia.ExchangePair{pair}); err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *CryptoDotComScraper) mainLoop() {
	defer s.cleanup()

	for {
		select {
		case <-s.shutdown:
			log.Println("CryptoDotComScraper: Shutting down main loop")
			return
		default:
		}

		var res cryptoDotComWSResponse
		if err := s.wsConn().ReadJSON(&res); err != nil {
			log.Warnf("CryptoDotComScraper: Creating a new connection caused by err=%s", err.Error())

			if retryErr := s.retryConnection(); retryErr != nil {
			   s.setError(retryErr)
			   log.Errorf("CryptoDotComScraper: Shutting down main loop after retrying to create a new connection, err=%s", retryErr.Error())

			   return
			}

			log.Info("CryptoDotComScraper: Successfully created a new connection")

			continue
		}
		if res.Code == cryptoDotComRateLimitError {
			if err := s.retryTask(res.ID); err != nil {
				s.setError(err)
				log.Errorf("CryptoDotComScraper: Shutting down main loop due to failing to retry a task, err=%s", err.Error())

				return
			}

			continue
		}
		if res.Code != 0 {
			log.Errorf("CryptoDotComScraper: Shutting down main loop due to non-retryable response code %d", res.Code)

			return
		}

		switch res.Method {
		case "public/heartbeat":
			if err := s.ping(res.ID); err != nil {
				s.setError(err)
				log.Errorf("CryptoDotComScraper: Shutting down main loop due to heartbeat failure, err=%s", err.Error())

				return
			}
		case "subscribe":
			if len(res.Result) == 0 {
				continue
			}

			var subscription cryptoDotComWSSubscriptionResult
			if err := json.Unmarshal(res.Result, &subscription); err != nil {
				s.setError(err)
				log.Errorf("CryptoDotComScraper: Shutting down main loop due to response unmarshaling failure, err=%s", err.Error())

				return
			}
			if subscription.Channel != "trade" {
				continue
			}

			baseCurrency := strings.Split(subscription.InstrumentName, `_`)[0]
			pair, err := s.db.GetExchangePairCache(s.exchangeName, subscription.InstrumentName)
			if err != nil {
				log.Error(err)
			}

			for _, data := range subscription.Data {
				var i cryptoDotComWSInstrument
				if err := json.Unmarshal(data, &i); err != nil {
					s.setError(err)
					log.Errorf("CryptoDotComScraper: Shutting down main loop due to instrument unmarshaling failure, err=%s", err.Error())

					return
				}

				volume := i.Quantity
				if i.Side != cryptoDotComSpotTradingBuy {
					volume = -volume
				}

				trade := &dia.Trade{
					Symbol:         baseCurrency,
					Pair:           subscription.InstrumentName,
					Price:          i.Price,
					Time:           time.Unix(0, i.TradeTime*int64(time.Millisecond)),
					Volume:         volume,
					Source:         s.exchangeName,
					ForeignTradeID: strconv.Itoa(i.TradeID),
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

func (s *CryptoDotComScraper) newConn() error {
	conn, _, err := ws.DefaultDialer.Dial(cryptoDotComWSEndpoint, nil)
	if err != nil {
		return err
	}

	// Crypto.com recommends adding a 1-second sleep after establishing the websocket connection, and before requests are sent
	// to avoid occurrences of rate-limit (`TOO_MANY_REQUESTS`) errors.
	// https://exchange-docs.crypto.com/spot/index.html?javascript#websocket-subscriptions
	time.Sleep(time.Duration(1) * time.Second)

	defer s.connMutex.Unlock()
	s.connMutex.Lock()
	s.ws = conn

	return nil
}

func (s *CryptoDotComScraper) wsConn() *ws.Conn {
	defer s.connMutex.RUnlock()
	s.connMutex.RLock()

	return s.ws
}

func (s *CryptoDotComScraper) ping(id int) error {
	s.rl.Take()

	return s.wsConn().WriteJSON(&cryptoDotComWSRequest{
		ID:     id,
		Method: "public/respond-heartbeat",
	})
}

func (s *CryptoDotComScraper) cleanup() {
	if err := s.wsConn().Close(); err != nil {
		s.setError(err)
	}

	close(s.chanTrades)
	s.close()
	s.signalShutdownDone.Do(func() {
		close(s.shutdownDone)
	})
}

func (s *CryptoDotComScraper) error() error {
	s.errMutex.RLock()
	defer s.errMutex.RUnlock()

	return s.err
}

func (s *CryptoDotComScraper) setError(err error) {
	s.errMutex.Lock()
	defer s.errMutex.Unlock()

	s.err = err
}

func (s *CryptoDotComScraper) isClosed() bool {
	s.closedMutex.RLock()
	defer s.closedMutex.RUnlock()

	return s.closed
}

func (s *CryptoDotComScraper) close() {
	s.closedMutex.Lock()
	defer s.closedMutex.Unlock()

	s.closed = true
}

func (s *CryptoDotComScraper) subscribe(pairs []dia.ExchangePair) error {
	channels := make([]string, len(pairs))
	for idx, pair := range pairs {
		channels[idx] = "trade." + pair.ForeignName
		s.pairScrapers.Store(pair.ForeignName, pair)
	}

	taskID := int(atomic.AddInt32(&s.taskCount, 1))
	task := cryptoDotComWSTask{
		Method: "subscribe",
		Params: cryptoDotComWSRequestParams{
			Channels: channels,
		},
		RetryCount: 0,
	}
	s.tasks.Store(taskID, task)

	return s.send(taskID, task)
}

func (s *CryptoDotComScraper) unsubscribe(pairs []dia.ExchangePair) error {
	channels := make([]string, len(pairs))
	for idx, pair := range pairs {
		channels[idx] = "trade." + pair.ForeignName
		s.pairScrapers.Delete(pair.ForeignName)
	}

	taskID := int(atomic.AddInt32(&s.taskCount, 1))
	task := cryptoDotComWSTask{
		Method: "unsubscribe",
		Params: cryptoDotComWSRequestParams{
			Channels: channels,
		},
		RetryCount: 0,
	}
	s.tasks.Store(taskID, task)

	return s.send(taskID, task)
}

func (s *CryptoDotComScraper) retryConnection() error {
	s.connRetryCount += 1
	if s.connRetryCount > cryptoDotComConnMaxRetry {
		return errors.New("CryptoDotComPairScraper: Reached max retry connection")
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

func (s *CryptoDotComScraper) retryTask(taskID int) error {
	val, ok := s.tasks.Load(taskID)
	if !ok {
		return fmt.Errorf("CryptoDotComScraper: Facing unknown task id, taskId=%d", taskID)
	}

	task := val.(cryptoDotComWSTask)
	task.RetryCount += 1
	if task.RetryCount > cryptoDotComTaskMaxRetry {
		return fmt.Errorf("CryptoDotComScraper: Exeeding max retry, taskId=%d, %s", taskID, task.toString())
	}

	log.Warnf("CryptoDotComScraper: Retrying a task, taskId=%d, %s", taskID, task.toString())
	s.tasks.Store(taskID, task)

	return s.send(taskID, task)
}

func (s *CryptoDotComScraper) send(taskID int, task cryptoDotComWSTask) error {
	s.rl.Take()

	return s.wsConn().WriteJSON(&cryptoDotComWSRequest{
		ID:     taskID,
		Method: task.Method,
		Params: task.Params,
		Nonce:  time.Now().UnixNano() / 1000,
	})
}

// CryptoDotComPairScraper implements PairScraper for Crypto.com
type CryptoDotComPairScraper struct {
	parent *CryptoDotComScraper
	pair   dia.ExchangePair
	closed bool
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (p *CryptoDotComPairScraper) Error() error {
	return p.parent.error()
}

// Pair returns the pair this scraper is subscribed to
func (p *CryptoDotComPairScraper) Pair() dia.ExchangePair {
	return p.pair
}

// Close stops listening for trades of the pair associated with the Crypto.com scraper
func (p *CryptoDotComPairScraper) Close() error {
	if err := p.parent.error(); err != nil {
		return err
	}
	if p.closed {
		return errors.New("CryptoDotComPairScraper: Already closed")
	}
	if err := p.parent.unsubscribe([]dia.ExchangePair{p.pair}); err != nil {
		return err
	}

	p.closed = true

	return nil
}
