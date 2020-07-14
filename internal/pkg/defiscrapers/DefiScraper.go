package defiscrapers

import (
	"errors"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

const (
	// Determine frequency of scraping
	refrestRateDelay  = time.Second * 10 * 1
	refreshStateDelay = time.Second * 10 * 1
)

type nothing struct{}

type DefiScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing

	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock     sync.RWMutex
	error         error
	closed        bool
	tickerRate    *time.Ticker
	tickerState   *time.Ticker
	datastore     models.Datastore
	chanDefiRate  chan *dia.DefiRate
	chanDefiState chan *dia.DefiProtocolState
}

// SpawnDefiScraper returns a new DefiScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func SpawnDefiScraper(datastore models.Datastore, rateType string) *DefiScraper {
	s := &DefiScraper{
		shutdown:      make(chan nothing),
		shutdownDone:  make(chan nothing),
		error:         nil,
		tickerRate:    time.NewTicker(refrestRateDelay),
		tickerState:   time.NewTicker(refreshStateDelay),
		datastore:     datastore,
		chanDefiRate:  make(chan *dia.DefiRate),
		chanDefiState: make(chan *dia.DefiProtocolState),
	}

	log.Info("Defi scraper is built and triggered")
	go s.mainLoop(rateType)
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *DefiScraper) mainLoop(rateType string) {
	for {
		select {
		case <-s.tickerRate.C:
			s.UpdateRates(rateType)
		case <-s.tickerState.C:
			s.UpdateState(rateType)

		case <-s.shutdown: // user requested shutdown
			log.Println("DefiScraper shutting down")
			s.cleanup(nil)
			return
		}
	}
}

// closes all connected Scrapers. Must only be called from mainLoop
func (s *DefiScraper) cleanup(err error) {

	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	s.tickerRate.Stop()
	s.tickerState.Stop()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections
func (s *DefiScraper) Close() error {
	if s.closed {
		return errors.New("DefiScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Channel returns a channel that can be used to receive rate information
func (s *DefiScraper) RateChannel() chan *dia.DefiRate {
	return s.chanDefiRate
}

func (s *DefiScraper) StateChannel() chan *dia.DefiProtocolState {
	return s.chanDefiState
}

// UpdateRates calls the appropriate function corresponding to the rate type.
func (s *DefiScraper) UpdateRates(defiType string) error {
	switch defiType {
	case "DYDX":
		{

			protocol := dia.DefiProtocol{
				Name:                 "DYDX",
				Address:              "0x1e0447b19bb6ecfdae1e4ae1694b0c3659614e4e",
				UnderlyingBlockchain: "Ethereum",
				Token:                "",
			}
			s.datastore.SetDefiProtocol(protocol)
			return s.UpdateDYDX(protocol)
		}
	case "AAVE":
		{

			protocol := dia.DefiProtocol{
				Name:                 "AAVE",
				Address:              "0x3dfd23A6c5E8BbcFc9581d2E864a68feb6a076d3",
				UnderlyingBlockchain: "Ethereum",
				Token:                "",
			}
			s.datastore.SetDefiProtocol(protocol)
			return s.UpdateAAVE(protocol)
		}
	case "RAY":
		{
			protocol := dia.DefiProtocol{
				Name:                 "RAY",
				Address:              "0xE215e8160a5e0A03f2D6c7900b050F2f04eA5Cbb",
				UnderlyingBlockchain: "Ethereum",
				Token:                "",
			}
			s.datastore.SetDefiProtocol(protocol)
			return s.updateRAY(protocol)
		}
	case "DHARMA":
		{

			protocol := dia.DefiProtocol{
				Name:                 "DHARMA",
				Address:              "0x3f320a0B08B93D7562c1f2d008d8154c44147620",
				UnderlyingBlockchain: "Ethereum",
				Token:                "",
			}
			s.datastore.SetDefiProtocol(protocol)
			return s.UpdateDHARMA(protocol)
		}

	}
	return errors.New("Error: " + defiType + " does not exist in database")
}

func (s *DefiScraper) UpdateState(defiType string) error {
	switch defiType {
	case "DYDX":
		{
			return s.UpdateDYDXState("DYDX")
		}
	case "AAVE":
		{
			return s.UpdateAAVEState("AAVE")
		}
	case "RAY":
		{
			return s.UpdateRAYState("RAY")
		}

	}
	return errors.New("Error: " + defiType + " does not exist in database")
}
