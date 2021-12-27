package scrapers

import (
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cloudingcity/go-ftx/ftx"
	"github.com/cloudingcity/go-ftx/ftx/stream"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
)

const (
	ftxSpotTradingPair        = "spot"
	ftxSpotTradingBuy         = "buy"
	ftxMaxConsecutiveErrCount = 10
)

// FTXScraper is a scraper for FTX
type FTXScraper struct {
	api *ftx.Client
	ws  *stream.Conn

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
	pairScrapers map[string]*FTXPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB
}

// NewFTXScraper returns a new FTX scraper
func NewFTXScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *FTXScraper {
	s := &FTXScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*FTXPairScraper),
		exchangeName: exchange.Name,
		err:          nil,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}

	client := ftx.New()
	ws, err := client.Connect()
	if err != nil {
		log.Error(err)

		return nil
	}

	s.ws = ws
	s.api = client

	if scrape {
		go s.mainLoop()
	}

	return s
}

// Close unsubscribes data and closes any existing WebSocket connections, as well as channels of FTXPairScraper
func (s *FTXScraper) Close() error {
	if s.isClosed() {
		return errors.New("FTXScraper: Already closed")
	}

	s.signalShutdown.Do(func() {
		close(s.shutdown)
	})

	<-s.shutdownDone

	return s.error()
}

// Channel returns a channel that can be used to receive trades
func (s *FTXScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// FetchAvailablePairs returns all traded pairs on FTX
func (s *FTXScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	markets, err := s.api.Markets.All()
	if err != nil {
		return nil, err
	}

	for _, m := range markets {
		if m.Type != ftxSpotTradingPair {
			continue
		}

		pairs = append(pairs, dia.ExchangePair{
			Symbol:      m.BaseCurrency,
			ForeignName: m.Name,
			Exchange:    s.exchangeName,
		})
	}

	return pairs, nil
}

// FillSymbolData adds the name to the asset underlying @symbol on FTX
func (s *FTXScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *FTXScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from the FTX scraper
func (s *FTXScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	if err := s.error(); err != nil {
		return nil, err
	}
	if s.isClosed() {
		return nil, errors.New("FTXScraper: Call ScrapePair on closed scraper")
	}

	ps := &FTXPairScraper{
		parent: s,
		pair:   pair,
	}

	s.pairScrapers[pair.ForeignName] = ps
	if err := s.subscribe(pair); err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *FTXScraper) mainLoop() {
	defer s.cleanup()

	go s.ping()

	for {
		select {
		case <-s.shutdown:
			log.Println("FTXScraper: Shutting down main loop")
			return
		default:
		}

		resp, err := s.ws.Recv()
		if err != nil {
			s.consecutiveErrCount += 1
			log.Errorf("FTXScraper: Main loop error %s", err.Error())
			if s.consecutiveErrCount > ftxMaxConsecutiveErrCount {
				s.setError(err)
				log.Error("FTXScraper: Shutting down main loop due to facing non-retryable errors")

				// Signal to other subscribers that the scraper is internally shut down.
				go s.signalShutdown.Do(func() {
					close(s.shutdown)
				})

				return
			}

			continue
		}

		s.consecutiveErrCount = 0

		switch v := resp.(type) {
		case stream.Trade:
			baseCurrency := strings.Split(v.Market, `/`)[0]
			pair, err := s.db.GetExchangePairCache(s.exchangeName, v.Market)
			if err != nil {
				log.Error(err)
			}

			for _, trade := range v.Data {
				volume := trade.Size
				if trade.Side != ftxSpotTradingBuy {
					volume = -volume
				}

				trade := &dia.Trade{
					Symbol:         baseCurrency,
					Pair:           v.Market,
					Price:          trade.Price,
					Time:           trade.Time,
					Volume:         volume,
					Source:         s.exchangeName,
					ForeignTradeID: strconv.Itoa(trade.ID),
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

func (s *FTXScraper) ping() {
	t := time.NewTicker(time.Duration(15) * time.Second)
	for {
		select {
		case <-s.shutdown:
			log.Println("FTXScraper: Shutting down ping loop")

			return
		case <-t.C:
			if err := s.ws.Ping(); err != nil {
				log.Warningf("FTXScraper: Sending pings fail %s", err.Error())
			}
		}
	}
}

func (s *FTXScraper) cleanup() {
	if err := s.ws.Close(); err != nil {
		s.setError(err)
	}

	close(s.chanTrades)
	s.close()
	s.signalShutdownDone.Do(func() {
		close(s.shutdownDone)
	})
}

func (s *FTXScraper) error() error {
	s.errMutex.RLock()
	defer s.errMutex.RUnlock()

	return s.err
}

func (s *FTXScraper) setError(err error) {
	s.errMutex.Lock()
	defer s.errMutex.Unlock()

	s.err = err
}

func (s *FTXScraper) isClosed() bool {
	s.closedMutex.RLock()
	defer s.closedMutex.RUnlock()

	return s.closed
}

func (s *FTXScraper) close() {
	s.closedMutex.Lock()
	defer s.closedMutex.Unlock()

	s.closed = true
}

func (s *FTXScraper) subscribe(pair dia.ExchangePair) error {
	if err := s.error(); err != nil {
		return err
	}
	if s.isClosed() {
		return errors.New("FTXScraper: Subscribe on closed scraper")
	}

	return s.ws.Subscribe(stream.ChannelTrades, pair.ForeignName)
}

func (s *FTXScraper) unsubscribe(pair dia.ExchangePair) error {
	if err := s.error(); err != nil {
		return err
	}
	if s.isClosed() {
		return errors.New("FTXScraper: Unsubscribe on closed scraper")
	}

	return s.ws.Unsubscribe(stream.ChannelTrades, pair.ForeignName)
}

// FTXPairScraper implements PairScraper for FTX
type FTXPairScraper struct {
	parent *FTXScraper
	pair   dia.ExchangePair
	closed bool
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (p *FTXPairScraper) Error() error {
	return p.parent.error()
}

// Pair returns the pair this scraper is subscribed to
func (p *FTXPairScraper) Pair() dia.ExchangePair {
	return p.pair
}

// Close stops listening for trades of the pair associated with the FTX scraper
func (p *FTXPairScraper) Close() error {
	if err := p.parent.error(); err != nil {
		return err
	}
	if p.closed {
		return errors.New("FTXPairScraper: Already closed")
	}
	if err := p.parent.unsubscribe(p.pair); err != nil {
		return err
	}

	p.closed = true

	return nil
}
