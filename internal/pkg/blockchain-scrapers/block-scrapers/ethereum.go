package blockscrapers

import (
	"context"
	"errors"
	"math/big"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/structs"
	"github.com/jackc/pgx/v4"

	models "github.com/diadata-org/diadata/pkg/model"
)

const (
	EthereumBlocks = "EthereumBlockScraper"
)

var (
	defEthereumBlockScraperConfig = &EthereumBlockScraperConfig{
		FollowDist: uint64(8),
	}

	// Start from genesis block when no state is stored.
	defEthereumBlockScraperState = &EthereumBlockScraperState{LastBlockNum: 0}
)

type EthereumBlockScraperConfig struct {
	// Stay @FollowDist behind the head
	FollowDist uint64 `json:"following_distance_blocks"`
}

type EthereumBlockScraperState struct {
	// last block number has been processed
	LastBlockNum uint64 `json:"last_block_num"`
}

type EthereumScraper struct {
	blockscraper BlockScraper
	client       *ethclient.Client
	ticker       *time.Ticker
	config       *EthereumBlockScraperConfig
	state        *EthereumBlockScraperState
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
		config:       &EthereumBlockScraperConfig{},
		state:        &EthereumBlockScraperState{},
	}

	err = s.initScraper(context.Background())
	if err != nil {
		log.Error("could not initialize scraper config and state: ", err)
		return nil
	}

	go s.mainLoop()
	return s
}

// init scraper
// if there are no values stored previously, use defaults and store them
func (s *EthereumScraper) initScraper(ctx context.Context) error {
	if err := s.loadConfig(ctx); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Errorf("unable to read scraper config from rdb: %s", err.Error())
			return err
		}

		// use & store defaults if there is no record in the scraper table

		defConf := *defEthereumBlockScraperConfig
		s.config = &defConf
		if err := s.blockscraper.relDB.SetScraperConfig(ctx, EthereumBlocks, s.config); err != nil {
			log.Errorf("unable to store scraper config on rdb: %s", err.Error())
			return err
		}

		defState := *defEthereumBlockScraperState
		s.state = &defState
		if err := s.blockscraper.relDB.SetScraperState(ctx, EthereumBlocks, s.state); err != nil {
			log.Errorf("unable to store scraper state on rdb: %s", err.Error())
			return err
		}

		return nil
	}

	return s.loadState(ctx)
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

	block, err := scraper.client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return err
	}
	currentBlockNumber := block.NumberU64()
	for i := scraper.state.LastBlockNum + 1; i < currentBlockNumber-scraper.config.FollowDist; i++ {
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
		blockdata.Number = strconv.Itoa(int(ethblockdata.Number))
		blockdata.Data = structs.Map(ethblockdata)

		scraper.GetDataChannel() <- blockdata

		err = scraper.storeState(context.Background())
		if err != nil {
			log.Error("store scraper state: ", err)
		}
	}
	return nil
}

func (scraper *EthereumScraper) GetDataChannel() chan dia.BlockData {
	return scraper.blockscraper.chanData
}

func (s *EthereumScraper) loadConfig(ctx context.Context) error {
	return s.blockscraper.relDB.GetScraperConfig(ctx, EthereumBlocks, s.config)
}

func (s *EthereumScraper) loadState(ctx context.Context) error {
	return s.blockscraper.relDB.GetScraperState(ctx, EthereumBlocks, s.state)
}

func (s *EthereumScraper) storeState(ctx context.Context) error {
	return s.blockscraper.relDB.SetScraperState(ctx, EthereumBlocks, s.state)
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
