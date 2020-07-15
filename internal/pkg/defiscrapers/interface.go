package defiscrapers

import (
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
)

type DeFIHelper interface {
	UpdateRate() error
	UpdateState() error
}

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
