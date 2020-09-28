package foreignscrapers

import (
	models "github.com/diadata-org/diadata/pkg/model"
	"sync"
	"time"
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
	chanQuotation  chan *models.ForeignQuotation

}

