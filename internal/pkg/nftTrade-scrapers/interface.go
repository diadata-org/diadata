package nfttradescrapers

import (
	"sync"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/ethclient"
)

type nothing struct{}

type NFTTradeScraper interface {
	GetTradeChannel() chan dia.NFTTrade
	FetchTrades() ([]dia.NFTTrade, error)
}

type TradeScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing

	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock     *sync.RWMutex
	error         error
	closed        bool
	ethConnection *ethclient.Client
	datastore     models.RelDatastore
	chanTrade     chan dia.NFTTrade
	source        string
}
