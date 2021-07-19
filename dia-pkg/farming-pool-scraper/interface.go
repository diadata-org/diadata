package pool

import (
	"sync"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
)

type nothing struct{}

type PoolHelper interface {
}

type PoolScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing

	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock    sync.RWMutex
	error        error
	closed       bool
	tickerRate   *time.Ticker
	datastore    models.Datastore
	chanPoolInfo chan *models.FarmingPool
	poolHelper   PoolHelper
	poolName     string
}
