package blockscrapers

import (
	"errors"
	"time"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/config"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/diadata-org/diadata/pkg/dia"
	subhelper "github.com/diadata-org/diadata/pkg/dia/helpers/substratehelper"

	models "github.com/diadata-org/diadata/pkg/model"
)

type SubstrateEventScraper struct {
	blockscraper            BlockScraper
	client                  *gsrpc.SubstrateAPI
	ticker                  *time.Ticker
	lastBlockNumber         int64
	substrateBlockchainName string
}

func NewSubstrateEventScraper(rdb *models.RelDB, substrateBlockchainName string) *SubstrateEventScraper {
	// Create our API with a default connection to the local node, use env RPC_URL.
	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}

	blockScraper := BlockScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		error:        nil,
		relDB:        rdb,
		chanData:     make(chan dia.BlockData),
	}
	s := &SubstrateEventScraper{
		lastBlockNumber:         1102858,
		blockscraper:            blockScraper,
		client:                  api,
		ticker:                  time.NewTicker(refreshDelay),
		substrateBlockchainName: substrateBlockchainName,
	}

	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *SubstrateEventScraper) mainLoop() {
	err := scraper.FetchData()
	if err != nil {
		log.Error("error updating block: ", err)
	}
	for {
		select {
		case <-scraper.ticker.C:
			err := scraper.FetchData()
			if err != nil {
				log.Error("error updating block: ", err)
			}
		case <-scraper.blockscraper.shutdown: // user requested shutdown
			log.Printf("Block scraper shutting down")
			err := scraper.Close()
			scraper.cleanup(err)
			return
		}
	}
}

func (scraper *SubstrateEventScraper) FetchData() error {

	// Fetch last scraped block number from db upon initialization.
	if scraper.lastBlockNumber == 0 {
		var err error
		scraper.lastBlockNumber, err = scraper.blockscraper.relDB.GetLastBlockBlockscraper(scraper.substrateBlockchainName)
		if err != nil {
			log.Errorf("could not find last scraped block: %v. Start from block 0.", err)
		}
	}

	block, err := scraper.client.RPC.Chain.GetBlockLatest()
	if err != nil {
		return err
	}
	meta, err := scraper.client.RPC.State.GetMetadataLatest()
	if err != nil {
		panic(err)
	}

	// Subscribe to system events via storage
	key, err := types.CreateStorageKey(meta, "System", "Events", nil)
	if err != nil {
		panic(err)
	}
	currentBlockNumber := block.Block.Header.Number
	for i := uint64(scraper.lastBlockNumber); i < uint64(currentBlockNumber-followDist); i++ {
		blockdata, err := subhelper.GetBlockDataOnChain(scraper.substrateBlockchainName, i, scraper.client, key, meta)
		if err != nil {
			return err
		}
		scraper.GetDataChannel() <- blockdata
		scraper.lastBlockNumber = int64(i)
	}

	return nil
}

func (scraper *SubstrateEventScraper) GetDataChannel() chan dia.BlockData {
	return scraper.blockscraper.chanData
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *SubstrateEventScraper) cleanup(err error) {
	scraper.blockscraper.errorLock.Lock()
	defer scraper.blockscraper.errorLock.Unlock()
	scraper.ticker.Stop()
	if err != nil {
		scraper.blockscraper.error = err
	}
	scraper.blockscraper.closed = true
	close(scraper.blockscraper.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections
func (scraper *SubstrateEventScraper) Close() error {
	if scraper.blockscraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.blockscraper.shutdown)
	<-scraper.blockscraper.shutdownDone
	scraper.blockscraper.errorLock.RLock()
	defer scraper.blockscraper.errorLock.RUnlock()
	return scraper.blockscraper.error
}
