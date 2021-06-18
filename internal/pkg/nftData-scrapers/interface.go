package nftdatascrapers

import (
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	refreshDelay = time.Hour * 24
)

type nothing struct{}
type NFTDataScraper interface {
	GetDataChannel() chan dia.NFT
	FetchData() ([]dia.NFT, error)
}

type NFTScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing

	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock     *sync.RWMutex
	error         error
	closed        bool
	ethConnection *ethclient.Client
	relDB         models.RelDatastore
	chanData      chan dia.NFT
}
