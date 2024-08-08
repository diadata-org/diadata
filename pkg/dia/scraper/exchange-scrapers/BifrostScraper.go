package scrapers

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	// "github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/diadata-org/diadata/pkg/dia"
	model "github.com/diadata-org/diadata/pkg/model"
	"github.com/sirupsen/logrus"
)

const BifrostRPCURL = "wss://bifrost-rpc-url"

type BifrostScraper struct {
	logger                    *logrus.Entry
	shutdown                  chan struct{}
	shutdownDone              chan struct{}
	errorLock                 sync.RWMutex
	error                     error
	closed                    bool
	pairScrapers              map[string]*BifrostPairScraper
	client                    *gsrpc.SubstrateAPI
	ticker                    *time.Ticker
	exchangeName              string
	blockchain                string
	currentHeight             int
	chanTrades                chan *dia.Trade
	chanAssets                chan *dia.Asset
	db                        *model.RelDB
	refreshDelay              time.Duration
	sleepBetweenContractCalls time.Duration
}

func NewBifrostScraper(exchange dia.Exchange, scrape bool, relDB *model.RelDB) *BifrostScraper {
	refreshDelay := time.Duration(30) * time.Second
	sleepBetweenContractCalls := time.Duration(1) * time.Second

	client, err := gsrpc.NewSubstrateAPI(BifrostRPCURL)
	if err != nil {
		log.Fatalf("Failed to connect to Bifrost RPC: %v", err)
	}

	s := &BifrostScraper{
		shutdown:                  make(chan struct{}),
		shutdownDone:              make(chan struct{}),
		pairScrapers:              make(map[string]*BifrostPairScraper),
		client:                    client,
		ticker:                    time.NewTicker(refreshDelay),
		exchangeName:              exchange.Name,
		blockchain:                exchange.BlockChain.Name,
		currentHeight:             0,
		chanTrades:                make(chan *dia.Trade),
		chanAssets:                make(chan *dia.Asset),
		db:                        relDB,
		refreshDelay:              refreshDelay,
		sleepBetweenContractCalls: sleepBetweenContractCalls,
	}

	s.logger = logrus.
		New().
		WithContext(context.Background()).
		WithField("context", "BifrostDEXScraper")

	if scrape {
		go s.mainLoop()
	}
	return s
}

func (s *BifrostScraper) mainLoop() {
	for {
		select {
		case <-s.ticker.C:
			err := s.Update()
			if err != nil {
				s.logger.Error(err)
			}
		case <-s.shutdown:
			s.logger.Println("shutting down")
			s.cleanup(nil)
			return
		}
	}
}

func (s *BifrostScraper) Update() error {
	logger := s.logger.WithFields(logrus.Fields{
		"function": "Update",
	})

	meta, err := s.client.RPC.State.GetMetadataLatest()
	if err != nil {
		logger.WithError(err).Error("failed to get metadata")
		return err
	}

	err = s.collectAssets(meta)
	if err != nil {
		logger.WithError(err).Error("failed to collect assets")
		return err
	}

	return nil
}

func (s *BifrostScraper) collectAssets(meta *types.Metadata) error {
    // Create a storage key to access the assets in the DEX
    key, err := types.CreateStorageKey(meta, "Dex", "Assets", nil)
    if err != nil {
        return err
    }
    fmt.Printf("Storage Key: %x\n", key) // Print the storage key

    // Get the hash of the latest block
    latestBlockHash, err := s.client.RPC.Chain.GetBlockHashLatest()
    if err != nil {
        return err
    }
    fmt.Printf("Latest Block Hash: %x\n", latestBlockHash) // Print the latest block hash

    // Use the storage key and latest block hash to fetch the raw storage data
    raw, err := s.client.RPC.State.GetStorageRaw(key, latestBlockHash)
    if err != nil {
        return err
    }
    if raw == nil {
        fmt.Println("No raw data found.")
        return nil
    }
    fmt.Printf("Raw Data: %x\n", raw) // Print the raw data in hex format
    fmt.Printf("Raw Data (string): %s\n", string(*raw)) // Print the raw data as a string

    // // Assuming the raw data is a scale-encoded list of assets
    // var assets []dia.Asset
    // err = types.DecodeFromBytes(*raw, &assets)
    // if err != nil {
    //     fmt.Printf("Error decoding raw data: %v\n", err) // Print the error if decoding fails
    //     return err
    // }
    // fmt.Printf("Decoded Assets: %+v\n", assets) // Print the decoded assets

    // // Send each decoded asset to the chanAssets channel
    // for _, asset := range assets {
    //     s.chanAssets <- &asset
    // }

    return nil
}


func (s *BifrostScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	s.ticker.Stop()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone)
}

func (s *BifrostScraper) Close() error {
	if s.closed {
		return errors.New("BifrostScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (s *BifrostScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

func (s *BifrostScraper) AssetChannel() chan *dia.Asset {
	return s.chanAssets
}

type BifrostPairScraper struct {
	parent     *BifrostScraper
	pair       dia.ExchangePair
	closed     bool
	lastRecord int64
}

func (ps *BifrostPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

func (ps *BifrostPairScraper) Close() error {
	ps.closed = true
	return nil
}

func (ps *BifrostPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}
