package blockscrapers

import (
	"context"
	"errors"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	"github.com/ethereum/go-ethereum/ethclient"

	models "github.com/diadata-org/diadata/pkg/model"
)

const (
	followDist = 2
)

type EthereumScraper struct {
	blockscraper    BlockScraper
	client          *ethclient.Client
	ticker          *time.Ticker
	lastBlockNumber int64
}

func NewEthereumScraper(rdb *models.RelDB) *EthereumScraper {
	connection, err := ethhelper.NewETHClient()
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	blockScraper := BlockScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		error:        nil,
		relDB:        rdb,
		chanData:     make(chan dia.BlockData),
	}
	s := &EthereumScraper{
		blockscraper: blockScraper,
		client:       connection,
		ticker:       time.NewTicker(refreshDelay),
	}

	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *EthereumScraper) mainLoop() {
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

func (scraper *EthereumScraper) FetchData() error {

	// Fetch last scraped block number from db upon initialization.
	if scraper.lastBlockNumber == 0 {
		var err error
		scraper.lastBlockNumber, err = scraper.blockscraper.relDB.GetLastBlockBlockscraper(dia.ETHEREUM)
		if err != nil {
			log.Errorf("could not find last scraped block: %v. Start from block 0.", err)
		}
	}

	block, err := scraper.client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return err
	}
	currentBlockNumber := block.NumberU64()
	for i := int(scraper.lastBlockNumber); i < int(currentBlockNumber)-followDist; i++ {
		blockdata, err := ethhelper.GetBlockDataOnChain(int64(i), scraper.client)
		if err != nil {
			return err
		}
		scraper.GetDataChannel() <- blockdata
		scraper.lastBlockNumber = int64(i)
	}

	return nil
}

func (scraper *EthereumScraper) GetDataChannel() chan dia.BlockData {
	return scraper.blockscraper.chanData
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *EthereumScraper) cleanup(err error) {
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
func (scraper *EthereumScraper) Close() error {
	if scraper.blockscraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.blockscraper.shutdown)
	<-scraper.blockscraper.shutdownDone
	scraper.blockscraper.errorLock.RLock()
	defer scraper.blockscraper.errorLock.RUnlock()
	return scraper.blockscraper.error
}
