package stockscrapers

import (
	"sync"

	models "github.com/diadata-org/diadata/pkg/model"
)

type nothing struct{}

type StockScraperInterface interface {
	GetStockQuotationChannel() chan models.StockQuotation
	FetchQuotes() error
}

type StockScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing

	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock *sync.RWMutex
	error     error
	closed    bool
	datastore models.Datastore
	chanStock chan models.StockQuotation
	source    string
}
