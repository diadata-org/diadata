package scrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	ws "github.com/gorilla/websocket"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
)

const (
	cryptoDotComAPIEndpoint            = "https://api.crypto.com/v2"
	cryptoDotComWSEndpoint             = "wss://stream.crypto.com/v2/market"
	cryptoDotComMaxConsecutiveErrCount = 10
)

// Instrument represents a trading pair
type Instrument struct {
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

// InstrumentResponse is an API response for retrieving instruments
type InstrumentResponse struct {
	Code   int `json:"code"`
	Result struct {
		Instruments []Instrument `json:"instruments"`
	} `json:"result"`
}

// CryptoDotComScraper is a scraper for Crypto.com
type CryptoDotComScraper struct {
	ws *ws.Conn

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
	pairScrapers map[string]*CryptoDotComPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB
}

// NewCryptoDotComScraper returns a new Crypto.com scraper
func NewCryptoDotComScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *CryptoDotComScraper {
	s := &CryptoDotComScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*CryptoDotComPairScraper),
		exchangeName: exchange.Name,
		err:          nil,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}

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

	var res InstrumentResponse
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
	return nil, nil
}

func (s *CryptoDotComScraper) mainLoop() {
}

func (s *CryptoDotComScraper) ping() {
}

func (s *CryptoDotComScraper) cleanup() {
	if err := s.ws.Close(); err != nil {
		s.setError(err)
	}

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

func (s *CryptoDotComScraper) subscribe(pair dia.ExchangePair) error {
	// TODO: enforce rate limit

	return nil
}

func (s *CryptoDotComScraper) unsubscribe(pair dia.ExchangePair) error {
	// TODO: enforce rate limit

	return nil
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
	if err := p.parent.unsubscribe(p.pair); err != nil {
		return err
	}

	p.closed = true

	return nil
}
