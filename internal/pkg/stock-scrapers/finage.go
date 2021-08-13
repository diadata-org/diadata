package stockscrapers

import (
	"errors"
	"fmt"

	models "github.com/diadata-org/diadata/pkg/model"
)

type FinageScraper struct {
	stockScraper               StockScraper
	apiRestURL                 string
	apiWsURL                   string
	timeResolutionMilliseconds int
}

func NewFinageScraper(db *models.DB) *FinageScraper {

	stockScraper := StockScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		error:        nil,
		datastore:    db,
		chanStock:    make(chan models.StockQuotation),
		source:       "Finage",
	}
	s := &FinageScraper{
		stockScraper: stockScraper,
		// TO DO: Will be fetched as secret as soon as in production
		apiWsURL:   "",
		apiRestURL: "",
		// TimeResolution
		timeResolutionMilliseconds: 1000,
	}
	fmt.Println("scraper built. Start main loop.")
	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *FinageScraper) mainLoop() {

	// Either call FetchQuotes or implement quote fetching here and
	// leave FetchQuotes blank.

}

// FetchQuotes fetches quotes from an API and feeds them into the channel.
func (scraper *FinageScraper) FetchQuotes() error {
	// ...
	// scraper.GetStockChannel() <- quote
	return nil
}

// GetDataChannel returns the scrapers data channel.
func (scraper *FinageScraper) GetStockQuotationChannel() chan models.StockQuotation {
	return scraper.stockScraper.chanStock
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *FinageScraper) cleanup(err error) {
	scraper.stockScraper.errorLock.Lock()
	defer scraper.stockScraper.errorLock.Unlock()
	if err != nil {
		scraper.stockScraper.error = err
	}
	scraper.stockScraper.closed = true
	close(scraper.stockScraper.shutdownDone)
}

// Close closes any existing API connections
func (scraper *FinageScraper) Close() error {
	if scraper.stockScraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.stockScraper.shutdown)
	<-scraper.stockScraper.shutdownDone
	scraper.stockScraper.errorLock.RLock()
	defer scraper.stockScraper.errorLock.RUnlock()
	return scraper.stockScraper.error
}
