package pool

import (
	"errors"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
)

const (
	// Determine frequency of scraping
	refreshRateDelay  = time.Second * 1 * 1
)

func SpawnPoolScraper(datastore models.Datastore, poolName string) *PoolScraper {
	s := &PoolScraper{
		shutdown:      make(chan nothing),
		shutdownDone:  make(chan nothing),
		error:         nil,
		tickerRate:    time.NewTicker(refreshRateDelay),
		datastore:     datastore,
		chanPoolInfo:  make(chan *models.PoolRate),
		poolName: poolName,



	}

	log.Info("Pool scraper is built and triggered")

	switch poolName {
	case "CVAULT":
		s.poolHelper = NewCvaultScrapper(s)
	}

	go s.mainLoop(poolName)
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *PoolScraper) mainLoop(rateType string) {
	for {
		select {
		case <-s.tickerRate.C:
			s.UpdateRates(rateType)

		case <-s.shutdown: // user requested shutdown
			log.Println("PoolScrapper shutting down")
			s.cleanup(nil)
			return
		}
	}
}

// closes all connected Scrapers. Must only be called from mainLoop
func (s *PoolScraper) cleanup(err error) {

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
func (s *PoolScraper) Close() error {
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
func (s *PoolScraper) RateChannel() chan *models.PoolRate {
	return s.chanPoolInfo
}

// UpdateRates calls the appropriate function corresponding to the rate type.
func (s *PoolScraper) UpdateRates(poolName string) error {

	//s.datastore.SetDefiProtocol(protocol)
	return s.poolHelper.UpdateRate()
}
