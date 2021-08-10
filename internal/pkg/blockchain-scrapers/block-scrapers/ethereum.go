package blockscrapers

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/structs"

	models "github.com/diadata-org/diadata/pkg/model"
)

const (
	followDist = 8
)

type EthereumScraper struct {
	blockscraper BlockScraper
	client       *ethclient.Client
	ticker       *time.Ticker
}

type EthereumBlockData struct {
	GasLimit    uint64             `json:"gas_limit"`
	GasUsed     uint64             `json:"gas_used"`
	Difficulty  *big.Int           `json:"difficulty"`
	Time        uint64             `json:"time"`
	Size        common.StorageSize `json:"size"`
	Number      uint64             `json:"number"`
	MixDigest   common.Hash        `json:"mix_digest"`
	Nonce       uint64             `json:"nonce"`
	Coinbase    common.Address     `json:"coinbase"`
	Root        common.Hash        `json:"root"`
	ParentHash  common.Hash        `json:"parent_hash"`
	TxHash      common.Hash        `json:"tx_hash"`
	ReceiptHash common.Hash        `json:"receipt_hash"`
	UncleHash   common.Hash        `json:"uncle_hash"`
	Extra       []byte             `json:"extra"`
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

	// Fetch last scraped block number from db.
	var lastBlockNumber int
	blockNumber, err := scraper.blockscraper.relDB.GetLastBlockBlockscraper(dia.ETHEREUM)
	if err != nil {
		log.Errorf("could not find last scraped block: %v. Start from block 0.", err)
	} else {
		lastBlockNumber = int(blockNumber)
		if err != nil {
			log.Error("parse last block number: ", err)
		}
	}

	block, err := scraper.client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return err
	}
	currentBlockNumber := block.NumberU64()
	for i := lastBlockNumber; i < int(currentBlockNumber)-followDist; i++ {
		var ethblockdata EthereumBlockData
		var blockdata dia.BlockData
		block, err := scraper.client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		if err != nil {
			return err
		}

		ethblockdata.Coinbase = block.Coinbase()
		ethblockdata.Difficulty = block.Difficulty()
		ethblockdata.Extra = block.Extra()
		ethblockdata.GasLimit = block.GasLimit()
		ethblockdata.GasUsed = block.GasUsed()
		ethblockdata.MixDigest = block.MixDigest()
		ethblockdata.Nonce = block.Nonce()
		ethblockdata.Number = block.NumberU64()
		ethblockdata.ParentHash = block.ParentHash()
		ethblockdata.ReceiptHash = block.ReceiptHash()
		ethblockdata.Root = block.Root()
		ethblockdata.Size = block.Size()
		ethblockdata.Time = block.Time()
		ethblockdata.TxHash = block.TxHash()
		ethblockdata.UncleHash = block.UncleHash()

		blockdata.BlockchainName = dia.ETHEREUM
		blockdata.BlockNumber = int64(ethblockdata.Number)
		blockdata.Data = structs.Map(ethblockdata)

		scraper.GetDataChannel() <- blockdata

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
