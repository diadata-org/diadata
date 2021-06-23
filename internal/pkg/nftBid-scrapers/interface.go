package nftbidscrapers

import (
	"sync"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/ethclient"
)

type nothing struct{}

type NFTBidScraper interface {
	// NFT bids should be streamed through dia.NFTBid channel.
	GetBidChannel() chan dia.NFTBid
	// Should fetch bids and send them to the channel.
	FetchBids() error
}

type BidScraper struct {
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
	chanBid       chan dia.NFTBid
}
