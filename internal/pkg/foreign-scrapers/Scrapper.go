package foreignscrapers

import (
	"sync"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type ForeignScrapperer interface {
	UpdateQuotation() error
	GetQuoteChannel() chan *models.ForeignQuotation
}

type ForeignScraper struct {
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
	chanQuotation chan *models.ForeignQuotation
}

// The below generic scraper is an empty scraper used for the default case of the main function
type genericScraper struct{}

func (scraper *genericScraper) UpdateQuotation() error {
	return nil
}
func (scraper *genericScraper) GetQuoteChannel() chan *models.ForeignQuotation {
	return make(chan *models.ForeignQuotation)
}
func NewGenericForeignScraper() *genericScraper {
	s := &genericScraper{}
	log.Info("started genericforeignscraper")
	return s
}
