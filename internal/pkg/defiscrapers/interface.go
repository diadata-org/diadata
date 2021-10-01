package defiscrapers

import (
	"context"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/sirupsen/logrus"
)

type DeFIHelper interface {
	UpdateRate() error
	UpdateState() error
}

type DefiScraper struct {
	ctx context.Context
	log *logrus.Entry

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
	relDB         *models.RelDB
	datastore     models.Datastore
	chanDefiRate  chan *dia.DefiRate
	chanDefiState chan *dia.DefiProtocolState
}
