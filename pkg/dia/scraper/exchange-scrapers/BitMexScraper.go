package scrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	ws "github.com/gorilla/websocket"
	"go.uber.org/ratelimit"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
)

const (
	bitMexAPIEndpoint = "https://www.bitMex.com/api/v1"
	bitMexWSEndpoint  = "wss://ws.bitMex.com/realtime"

	// bitMexWSRateLimitPerSec is a max request per second for sending websocket requests
	bitMexWSRateLimitPerSec = 10

	// bitMexTaskMaxRetry is a max retry value used when retrying subscribe/unsubscribe trades
	bitMexTaskMaxRetry = 20

	// bitMexConnMaxRetry is a max retry value used when retrying to create a new connection
	bitMexConnMaxRetry = 50

	// bitMexRateLimitError is a rate limit error code
	bitMexRateLimitError = 429

	// bitMexPingInterval is the number of seconds between ping messages
	bitMexPingInterval = 30
)

// bitMexWSTask is a websocket task tracking subscription/unsubscription
type bitMexWSTask struct {
	Op         string
	Args       []string
	RetryCount int
}

func (c *bitMexWSTask) toString() string {
	return fmt.Sprintf("op=%s, param=%s, retry=%d", c.Op, c.Args, c.RetryCount)
}

// bitMexWSRequest is a websocket request
type bitMexWSRequest struct {
	Op   string   `json:"op"`
	Args []string `json:"args,omitempty"`
}

// bitMexSubscriptionResult is a subscription result coming from websocket
type bitMexSubscriptionResult struct {
	Success   bool            `json:"success"`
	Subscribe string          `json:"subscribe"`
	Error     string          `json:"error"`
	Status    int             `json:"status"`
	Request   json.RawMessage `json:"request"`
	Table     string          `json:"table"`
	Trades    []bitMexWSTrade `json:"data"`
}

// bitMexSubscriptionTrade is a trade result coming from websocket
type bitMexSubscriptionTrade struct {
	Table  string          `json:"table"`
	Trades []bitMexWSTrade `json:"data"`
}

// bitMexWSTrade is a trade result coming from websocket
type bitMexWSTrade struct {
	Timestamp       time.Time `json:"timestamp"`
	Symbol          string    `json:"symbol"`
	Side            string    `json:"side"`
	Size            float64   `json:"size"`
	Price           float64   `json:"price"`
	TickDirection   string    `json:"tickDirection"`
	TrdMatchID      string    `json:"trdMatchID"`
	GrossValue      float64   `json:"grossValue"`
	HomeNotional    float64   `json:"homeNotional"`
	ForeignNotional float64   `json:"foreignNotional"`
}

// bitMexInstrument represents a trading pair
type bitMexInstrument struct {
	Symbol     string    `json:"symbol"`
	RootSymbol string    `json:"rootSymbol"`
	Expiry     time.Time `json:"expiry"`
}

// BitMexScraper is a scraper for bitmex.com
type BitMexScraper struct {
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
	pairScrapers    sync.Map
	exchangeName    string
	chanTrades      chan *dia.Trade
	db              *models.RelDB
	tasks           sync.Map
	pingTicker      *time.Ticker
	stopPingRoutine chan bool

	// used to handle connection retry
	connMutex      sync.RWMutex
	connRetryCount int
}

// NewBitMexScraper returns a new BitMex scraper
func NewBitMexScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BitMexScraper {
	s := &BitMexScraper{
		shutdown:        make(chan nothing),
		shutdownDone:    make(chan nothing),
		exchangeName:    exchange.Name,
		err:             nil,
		chanTrades:      make(chan *dia.Trade),
		db:              relDB,
		pingTicker:      time.NewTicker(bitMexPingInterval * time.Second),
		stopPingRoutine: make(chan bool),
	}

	if err := s.newConn(); err != nil {
		log.Error(err)

		return nil
	}

	s.rl = ratelimit.New(bitMexWSRateLimitPerSec)

	if scrape {
		go s.mainLoop()
	}

	return s
}

// Close unsubscribes data and closes any existing WebSocket connections, as well as channels of BitMexScraper
func (s *BitMexScraper) Close() error {
	if s.isClosed() {
		return errors.New("BitMexScraper: Already closed")
	}

	s.signalShutdown.Do(func() {
		close(s.shutdown)
	})

	<-s.shutdownDone

	return s.error()
}

// Channel returns a channel that can be used to receive trades
func (s *BitMexScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// FetchAvailablePairs returns all traded pairs on BitMex
func (s *BitMexScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {

	data, _, err := utils.GetRequest(bitMexAPIEndpoint + "/instrument")
	if err != nil {
		return nil, err
	}

	var res []bitMexInstrument
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	for _, i := range res {

		// fmt.Printf(`
		// {
		//     "Symbol": "%s",
		//     "ForeignName": "%s",
		//     "Exchange": "BitMex",
		//     "Ignore": false
		// },
		// `, i.RootSymbol, i.RootSymbol+"_"+strings.TrimPrefix(i.Symbol, i.RootSymbol))

		pairs = append(pairs, dia.ExchangePair{
			Symbol:      i.RootSymbol,
			ForeignName: i.RootSymbol + "_" + strings.TrimPrefix(i.Symbol, i.RootSymbol),
			Exchange:    s.exchangeName,
		})
	}

	return pairs, nil
}

// FillSymbolData adds the name to the asset underlying @symbol on BitMex
func (s *BitMexScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *BitMexScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from the BitMex scraper
func (s *BitMexScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {

	if err := s.error(); err != nil {
		return nil, err
	}

	if s.isClosed() {
		return nil, errors.New("BitMexScraper: Call ScrapePair on closed scraper")
	}

	ps := &BitMexPairScraper{
		parent: s,
		pair:   pair,
	}

	if err := s.subscribe([]dia.ExchangePair{pair}); err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *BitMexScraper) startPing() {
	for {
		select {
		case <-s.stopPingRoutine:
			return
		case <-s.pingTicker.C:
			s.ping()

		default:
		}
	}
}

func (s *BitMexScraper) mainLoop() {
	defer s.cleanup()

	go s.startPing()

	for {
		select {
		case <-s.shutdown:
			log.Println("BitMexScraper: Shutting down main loop")
			return
		default:
		}

		_, msg, err := s.wsConn().ReadMessage()
		if err != nil {

			log.Warnf("BitMexScraper: Creating a new connection caused by err=%s", err.Error())

			if retryErr := s.retryConnection(); retryErr != nil {
				s.setError(retryErr)
				log.Errorf("BitMexScraper: Shutting down main loop after retrying to create a new connection, err=%s", retryErr.Error())
				return
			}

			log.Info("BitMexScraper: Successfully created a new connection")

			continue

		}

		if string(msg) == "pong" {
			continue
		}

		var subResult bitMexSubscriptionResult
		if err := json.Unmarshal(msg, &subResult); err == nil {

			if subResult.Table == "trade" {
				s.handleTrades(subResult)
				continue
			}

			if subResult.Success {
				// subscription Success
				continue
			}
			if subResult.Status == bitMexRateLimitError {

				var failedRequest bitMexWSRequest
				if err := json.Unmarshal(subResult.Request, &failedRequest); err == nil {

					task := bitMexWSTask{
						Op:   failedRequest.Op,
						Args: failedRequest.Args,
					}
					if err := s.retryTask(s.getTaskID(task)); err != nil {
						s.setError(err)
						log.Errorf("BitMexScraper: Shutting down main loop due to failing to retry a task, err=%s", err.Error())

						return
					}

				}

				continue
			}

		} else {
			log.Println(err)
		}

		// log.Printf("Unhandled WS message: %s", bytes.NewBuffer(msg))

	}

}

func (s *BitMexScraper) handleTrades(tradesWsResponse bitMexSubscriptionResult) {
	tru := true
	if tru {
		if tru {

			var pair dia.ExchangePair
			for _, data := range tradesWsResponse.Trades {

				if pair == (dia.ExchangePair{}) {
					val, ok := s.pairScrapers.Load(data.Symbol)
					if !ok {
						log.Error("Pair not found %s", data.Symbol)
						continue
					} else {
						pair = val.(dia.ExchangePair)
					}
				}

				volume := data.ForeignNotional

				trade := &dia.Trade{
					Symbol:         pair.Symbol,
					Pair:           pair.ForeignName,
					Price:          data.Price,
					Time:           data.Timestamp,
					Volume:         volume,
					Source:         s.exchangeName,
					ForeignTradeID: data.TrdMatchID,
					VerifiedPair:   pair.Verified,
					BaseToken:      pair.UnderlyingPair.BaseToken,
					QuoteToken:     pair.UnderlyingPair.QuoteToken,
				}
				// fmt.Println("trade:", trade)
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

func (s *BitMexScraper) newConn() error {
	conn, _, err := ws.DefaultDialer.Dial(bitMexWSEndpoint, nil)
	if err != nil {
		return err
	}

	defer s.connMutex.Unlock()
	s.connMutex.Lock()
	s.ws = conn

	return nil
}

func (s *BitMexScraper) wsConn() *ws.Conn {
	defer s.connMutex.RUnlock()
	s.connMutex.RLock()

	return s.ws
}

func (s *BitMexScraper) ping() error {
	s.rl.Take()

	return s.wsConn().WriteMessage(ws.TextMessage, []byte("ping"))
}

func (s *BitMexScraper) cleanup() {
	s.pingTicker.Stop()
	s.stopPingRoutine <- true
	if err := s.wsConn().Close(); err != nil {
		s.setError(err)
	}

	close(s.chanTrades)
	s.close()
	s.signalShutdownDone.Do(func() {
		close(s.shutdownDone)
	})
}

func (s *BitMexScraper) error() error {
	s.errMutex.RLock()
	defer s.errMutex.RUnlock()

	return s.err
}

func (s *BitMexScraper) setError(err error) {
	s.errMutex.Lock()
	defer s.errMutex.Unlock()

	s.err = err
}

func (s *BitMexScraper) isClosed() bool {
	s.closedMutex.RLock()
	defer s.closedMutex.RUnlock()

	return s.closed
}

func (s *BitMexScraper) close() {
	s.closedMutex.Lock()
	defer s.closedMutex.Unlock()

	s.closed = true
}

func (s *BitMexScraper) subscribe(pairs []dia.ExchangePair) error {
	channels := make([]string, len(pairs))
	for idx, pair := range pairs {
		bitMexInstrumentSymbol := strings.Replace(pair.ForeignName, "_", "", 1)
		channels[idx] = "trade:" + bitMexInstrumentSymbol
		s.pairScrapers.Store(bitMexInstrumentSymbol, pair)
	}

	task := bitMexWSTask{
		Op:   "subscribe",
		Args: channels,

		RetryCount: 0,
	}
	taskID := s.getTaskID(task)
	s.tasks.Store(taskID, task)

	return s.send(taskID, task)
}

func (s *BitMexScraper) getTaskID(task bitMexWSTask) string {
	return fmt.Sprintf("%s-%s", task.Op, strings.Join(task.Args, ","))
}

func (s *BitMexScraper) unsubscribe(pairs []dia.ExchangePair) error {
	channels := make([]string, len(pairs))
	for idx, pair := range pairs {
		channels[idx] = "trade." + pair.ForeignName
		s.pairScrapers.Delete(pair.ForeignName)
	}

	task := bitMexWSTask{
		Op:         "unsubscribe",
		Args:       channels,
		RetryCount: 0,
	}
	taskID := s.getTaskID(task)
	s.tasks.Store(taskID, task)

	return s.send(taskID, task)
}

func (s *BitMexScraper) retryConnection() error {
	s.connRetryCount += 1
	if s.connRetryCount > bitMexConnMaxRetry {
		return errors.New("BitMexPairScraper: Reached max retry connection")
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

func (s *BitMexScraper) retryTask(taskID string) error {
	val, ok := s.tasks.Load(taskID)
	if !ok {
		return fmt.Errorf("BitMexScraper: Facing unknown task id, taskId=%d", taskID)
	}

	task := val.(bitMexWSTask)
	task.RetryCount += 1
	if task.RetryCount > bitMexTaskMaxRetry {
		return fmt.Errorf("BitMexScraper: Exeeding max retry, taskId=%d, %s", taskID, task.toString())
	}

	log.Warnf("BitMexScraper: Retrying a task, taskId=%d, %s", taskID, task.toString())
	s.tasks.Store(taskID, task)

	return s.send(taskID, task)
}

func (s *BitMexScraper) send(taskID string, task bitMexWSTask) error {
	s.rl.Take()

	return s.wsConn().WriteJSON(&bitMexWSRequest{
		Op:   task.Op,
		Args: task.Args,
	})
}

// BitMexPairScraper implements PairScraper for BitMex
type BitMexPairScraper struct {
	parent *BitMexScraper
	pair   dia.ExchangePair
	closed bool
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (p *BitMexPairScraper) Error() error {
	return p.parent.error()
}

// Pair returns the pair this scraper is subscribed to
func (p *BitMexPairScraper) Pair() dia.ExchangePair {
	return p.pair
}

// Close stops listening for trades of the pair associated with the BitMex scraper
func (p *BitMexPairScraper) Close() error {
	if err := p.parent.error(); err != nil {
		return err
	}
	if p.closed {
		return errors.New("BitMexPairScraper: Already closed")
	}
	if err := p.parent.unsubscribe([]dia.ExchangePair{p.pair}); err != nil {
		return err
	}

	p.closed = true

	return nil
}
