package scrapers

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/alexjorgef/go-bittrex/bittrex"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/zekroTJA/timedmap"
)

const (
	bittrexSpotTradingBuy     = "buy"
	bittrexMarketStatusActive = "ONLINE"

	// bittrexMaxConsecutiveErrCount is a max number of consecutive errors
	bittrexMaxConsecutiveErrCount = 10

	bittrexBackoffSeconds = 20
)

type BittrexScraper struct {
	api                   *bittrex.Bittrex
	chanTradesUnprocessed chan bittrex.Trade
	errCh                 chan error

	// signaling channels for session initialization and finishing
	//run                bool
	shutdown           chan nothing
	shutdownDone       chan nothing
	signalShutdown     sync.Once
	signalShutdownDone sync.Once

	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errMutex            sync.RWMutex
	err                 error
	closedMutex         sync.RWMutex
	closed              bool
	consecutiveErrCount int

	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*BittrexPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB
}

func NewBittrexScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BittrexScraper {
	s := &BittrexScraper{
		shutdown:              make(chan nothing),
		shutdownDone:          make(chan nothing),
		pairScrapers:          make(map[string]*BittrexPairScraper),
		exchangeName:          exchange.Name,
		err:                   nil,
		chanTrades:            make(chan *dia.Trade),
		chanTradesUnprocessed: make(chan bittrex.Trade),
		db:                    relDB,
	}

	client := bittrex.New("", "")

	s.api = client

	if scrape {
		go s.mainLoop()
	}

	return s
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *BittrexScraper) Close() error {
	if s.closed {
		return errors.New("BittrexScraper: Already closed")
	}

	s.signalShutdown.Do(func() {
		close(s.shutdown)
	})

	<-s.shutdownDone

	return s.error()
}

// Channel returns a channel that can be used to receive trades
func (s *BittrexScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *BittrexScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	markets, err := s.api.GetMarkets()
	if err != nil {
		return nil, err
	}

	for _, m := range markets {
		if m.Status == bittrexMarketStatusActive {
			pairs = append(pairs, dia.ExchangePair{
				Symbol:      m.BaseCurrencySymbol,
				ForeignName: m.Symbol,
				Exchange:    s.exchangeName,
			})
		}

	}

	return pairs, nil
}

// FillSymbolData adds the name to the asset underlying @symbol on Bittrex
func (s *BittrexScraper) FillSymbolData(symbol string) (asset dia.Asset, err error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *BittrexScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from the Bittrex scraper
func (s *BittrexScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	if err := s.error(); err != nil {
		return nil, err
	}
	if s.isClosed() {
		return nil, errors.New("BittrexScraper: Call ScrapePair on closed scraper")
	}

	p := &BittrexPairScraper{
		parent: s,
		pair:   pair,
		stopCh: make(chan bool),
	}

	s.pairScrapers[pair.ForeignName] = p
	if err := s.subscribe(pair); err != nil {
		return nil, err
	}
	return p, nil
}

// runs in a goroutine until s is closed
func (s *BittrexScraper) mainLoop() {
	defer s.cleanup()

	go s.ping()

	tmFalseDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
	tmDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)

	for {
		select {
		case <-s.shutdown:
			log.Println("BittrexScraper: Shutting down main loop")
		default:
		}

		select {
		case err := <-s.errCh:
			if err != nil {
				s.consecutiveErrCount += 1
				log.Errorf("BittrexScraper: Main loop error %s", err.Error())
				if s.consecutiveErrCount > bittrexMaxConsecutiveErrCount {
					s.setError(err)
					log.Error("BittrexScraper: Shutting down main loop due to facing non-retryable errors")

					// Signal to other subscribers that the scraper is internally shut down.
					go s.signalShutdown.Do(func() {
						close(s.shutdown)
					})
				}
				time.Sleep(time.Duration(bittrexBackoffSeconds) * time.Second)
			}
		case v := <-s.chanTradesUnprocessed:
			s.consecutiveErrCount = 0

			pair, err := s.db.GetExchangePairCache(s.exchangeName, v.Symbol)
			if err != nil {
				log.Error("get exchange pair from cache: ", err)
			}

			volume, _ := v.Quantity.Float64()
			pairs := strings.Split(v.Symbol, `-`)
			symbol := pairs[0] + "/" + pairs[1]
			if v.TakerSide != bittrexSpotTradingBuy {
				volume = -volume
			}
			price, _ := v.Rate.Float64()

			trade := &dia.Trade{
				Symbol:         pairs[0],
				Pair:           symbol,
				Price:          price,
				Time:           v.ExecutedAt,
				Volume:         volume,
				Source:         s.exchangeName,
				ForeignTradeID: v.ID,
				VerifiedPair:   pair.Verified,
				BaseToken:      pair.UnderlyingPair.BaseToken,
				QuoteToken:     pair.UnderlyingPair.QuoteToken,
			}
			if pair.Verified {
				log.Infoln("Got verified trade", trade)
			}
			// Handle duplicate trades.
			discardTrade := trade.IdentifyDuplicateFull(tmFalseDuplicateTrades, duplicateTradesMemory)
			if !discardTrade {
				trade.IdentifyDuplicateTagset(tmDuplicateTrades, duplicateTradesMemory)
				select {
				case <-s.shutdown:
				case s.chanTrades <- trade:
				}
			}

		}
	}
}

func (s *BittrexScraper) ping() {
	t := time.NewTicker(time.Duration(15) * time.Second)
	for {
		select {
		case <-s.shutdown:
			log.Println("BittrexScraper: Shutting down ping loop")

			return
		case <-t.C:
			if s.closed {
				log.Warningf("BittrexScraper: Sending pings fail")
			}
		}
	}
}

func (s *BittrexScraper) cleanup() {
	close(s.chanTrades)
	s.close()
	s.signalShutdownDone.Do(func() {
		close(s.shutdownDone)
	})
}

func (s *BittrexScraper) error() error {
	s.errMutex.RLock()
	defer s.errMutex.RUnlock()

	return s.err
}

func (s *BittrexScraper) setError(err error) {
	s.errMutex.Lock()
	defer s.errMutex.Unlock()

	s.err = err
}

func (s *BittrexScraper) isClosed() bool {
	s.closedMutex.RLock()
	defer s.closedMutex.RUnlock()

	return s.closed
}

func (s *BittrexScraper) close() {
	s.closedMutex.Lock()
	defer s.closedMutex.Unlock()

	s.closed = true
}

func (s *BittrexScraper) subscribe(pair dia.ExchangePair) error {
	if err := s.error(); err != nil {
		return err
	}
	if s.isClosed() {
		return errors.New("BittrexScraper: Subscribe on closed scraper")
	}

	go func(market string, stopCh chan bool) {
		s.errCh <- s.api.SubscribeTradeUpdates(market, s.chanTradesUnprocessed, stopCh)
	}(pair.ForeignName, s.pairScrapers[pair.ForeignName].stopCh)
	select {
	case err := <-s.errCh:
		if err != nil {
			return err
		}
	default:
	}
	return nil
}

func (s *BittrexScraper) unsubscribe(stopCh chan bool) error {
	if err := s.error(); err != nil {
		return err
	}
	if s.isClosed() {
		return errors.New("BittrexScraper: Unsubscribe on closed scraper")
	}

	select {
	case stopCh <- true:
		return nil
	default:
		return errors.New("BittrexScraper: Unsubscribe error")
	}
}

// BittrexPairScraper implements PairScraper for Bittrex
type BittrexPairScraper struct {
	parent *BittrexScraper
	pair   dia.ExchangePair
	closed bool
	stopCh chan bool
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (p *BittrexPairScraper) Error() error {
	return p.parent.error()
}

// Pair returns the pair this scraper is subscribed to
func (p *BittrexPairScraper) Pair() dia.ExchangePair {
	return p.pair
}

// Close stops listening for trades of the pair associated with the Bittrex scraper
func (p *BittrexPairScraper) Close() error {
	if err := p.parent.error(); err != nil {
		return err
	}
	if p.closed {
		return errors.New("BittrexPairScraper: Already closed")
	}
	if err := p.parent.unsubscribe(p.stopCh); err != nil {
		return err
	}

	p.closed = true

	return nil
}
