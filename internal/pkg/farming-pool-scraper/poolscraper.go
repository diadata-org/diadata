package pool

import (
	"errors"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
)

const (
	// Determine frequency of scraping
	refreshRateDelay = 1 * 20 * time.Second
	restDial         = "http://159.69.120.42:8545/"
	wsDial           = "ws://159.69.120.42:8546/"
)

func SpawnPoolScraper(datastore models.Datastore, poolName string) *PoolScraper {
	s := &PoolScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		error:        nil,
		tickerRate:   time.NewTicker(refreshRateDelay),
		datastore:    datastore,
		chanPoolInfo: make(chan *models.FarmingPool),
		poolName:     poolName,
	}

	log.Info("Pool scraper is built and triggered")

	switch poolName {
	case "CVAULT":
		s.poolHelper = NewCvaultScraper(s)
	case "YFI":
		s.poolHelper = NewYFIPool(s)
	case "BALANCER":
		s.poolHelper = NewBalancerPoolScrapper(s)
	case "BOND":
		s.poolHelper = NewBONDScraper(s)
	case "CURVEFI":
		s.poolHelper = NewCFIScraper(s)
	case "LOOPRING":
		s.poolHelper = NewLRCPoolScraper(s)
	case "SYNTHETIX":
		s.poolHelper = NewSynthetixScraper(s)
	}

	go s.mainLoop(poolName)
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *PoolScraper) mainLoop(rateType string) {
	for {
		select {

		case <-s.shutdown: // user requested shutdown
			log.Println("PoolScraper shutting down")
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

// RateChannel returns a channel that can be used to receive rate information
func (s *PoolScraper) RateChannel() chan *models.FarmingPool {
	return s.chanPoolInfo
}
