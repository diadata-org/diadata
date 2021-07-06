package blockscrapers

import (
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
)

const (
	refreshDelay = time.Hour * 24
)

type nothing struct{}
type BlockScraperInterface interface {
	// NFT data should be streamed through dia.NFT channel.
	GetDataChannel() chan dia.BlockData
	// Should fetch nft data and send it to the channel.
	FetchData() error
}

type BlockScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing

	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock *sync.RWMutex
	error     error
	closed    bool
	relDB     models.RelDatastore
	chanData  chan dia.BlockData
}
