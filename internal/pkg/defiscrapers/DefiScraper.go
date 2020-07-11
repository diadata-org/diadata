package defiscrapers

import (
"errors"
"sync"
"time"

models "github.com/diadata-org/diadata/pkg/model"
log "github.com/sirupsen/logrus"
"github.com/diadata-org/diadata/pkg/dia"

)

const (
	// Determine frequency of scraping
	refreshDelay = time.Second * 10 * 1
)

type nothing struct{}

type DefiScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing

	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock        sync.RWMutex
	error            error
	closed           bool
	ticker           *time.Ticker
	datastore        models.Datastore
	chanDefiRate chan *dia.DefiLendingRate
}

// SpawnDefiScraper returns a new DefiScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func SpawnDefiScraper(datastore models.Datastore, rateType string) *DefiScraper {
	s := &DefiScraper{
		shutdown:         make(chan nothing),
		shutdownDone:     make(chan nothing),
		error:            nil,
		ticker:           time.NewTicker(refreshDelay),
		datastore:        datastore,
		chanDefiRate: make(chan *dia.DefiLendingRate),
	}

	log.Info("Defi scraper is built and triggered")
	go s.mainLoop(rateType)
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *DefiScraper) mainLoop(rateType string) {
	for {
		select {
		case <-s.ticker.C:
			s.Update(rateType)
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

	s.ticker.Stop()

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
func (s *DefiScraper) Channel() chan *dia.DefiLendingRate {
	return s.chanDefiRate
}

// Update calls the appropriate function corresponding to the rate type.
func (s *DefiScraper) Update(defiType string) error {
	switch defiType {
	case "DYDX":
		{
			protocol := dia.DefiProtocol{
				Name: "dydx",
				Address: "",
				UnderlyingBlockchain: "",
				Token:"",
			}
			return s.UpdateDYDX(protocol)
		}

	}
	return errors.New("Error: " + defiType + " does not exist in database")
}
