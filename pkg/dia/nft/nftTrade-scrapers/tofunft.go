package nfttradescrapers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/config/nftContracts/erc721"
	"github.com/diadata-org/diadata/config/nftContracts/tofunft"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

const (
	// we assume all of the NFTs traded on TofuNFT are ERC721(1155 is an extension of it)
	tofuNFTNFTContractType = "ERC721"
)

type TofuNFTScraperConfig struct {
	// tofunft's exchange contract address on connected blockchain network
	ContractAddr string `json:"contract_addr"`

	// indicates the batch size during read the filtered events
	BatchSize int `json:"batch_size"`

	// wait for a while between batch retrieval of filtered events
	WaitPeriod time.Duration `json:"wait_per_batch"`

	// it enables read contract data from the event's block
	// height instead of the last state
	FollowDist int `json:"following_distance_blocks"`

	// if set it will read erc721 attributes at the currently
	// processing block
	UseArchiveNode bool `json:"use_archive_node_fetaures"`

	// indicates the number of retries to scrape the target
	// in case of an unexpected error
	MaxRetry int `json:"max_retry"`

	// if true the scraper will skip the currently scraping
	// block when retries reach to the value MaxRetry
	SkipOnErr bool `json:"skip_on_error"`

	// it limits read bytes for NFT's metadata from external url
	MaxMetadataSize int `json:"max_metadata_size"`

	// it limits duration of read for NFT's metadata from external url
	MetadataTimeout time.Duration `json:"metadata_timeout"`
}

type TofuNFTScraperState struct {
	// last block number has been processed
	LastBlockNum uint64 `json:"last_block_num"`

	// last transaction index in the block(curr) has been processed
	LastTxIndex uint `json:"last_tx_index"`

	// holds the latest error message that occurred while scraping
	LastErr string `json:"last_error"`

	// indicates the number of consecutive error, reset on any successful operation
	ErrCounter int `json:"count_of_error"`
}

type TofuNFTScraper struct {
	tradeScraper TradeScraper
	blockchain   string

	mu           sync.Mutex
	conf         *TofuNFTScraperConfig
	state        *TofuNFTScraperState
	exchangeName string
}

type astarERC721Transfer struct {
	NFTAddress  common.Address
	Name        *string
	Symbol      *string
	TotalSupply *big.Int
	From        common.Address
	To          common.Address
	TokenID     *big.Int
	TokenURI    *string
	TokenAttrs  map[string]interface{}
}

var (
	errTofuNFTShutdownRequest = errors.New("shutdown requested")

	// default values are valid for the first run which is it saves
	// these configs to the DB
	defTofuNFTConf = &TofuNFTScraperConfig{
		BatchSize:       5000,
		WaitPeriod:      20 * time.Second,
		FollowDist:      10,
		UseArchiveNode:  false,
		MaxRetry:        5,
		SkipOnErr:       true,
		MaxMetadataSize: 50 * 1024,
		MetadataTimeout: 30 * time.Second,
	}

	// // TofuNFT MarketNG contract has been deployed on the Astar mainnet at
	// // block num 225714, so scraper starts from this block
	// defTofuNFTState = &TofuNFTScraperState{LastBlockNum: 225714}
	defTofuNFTState = &TofuNFTScraperState{LastBlockNum: 11391732}

	// This string is the identifier of the scraper in conf and state fields in postgres.
	TofuNFT = ""

	tofuNFTABI abi.ABI
	_erc721ABI abi.ABI

	assetCacheTofuAstar = make(map[string]dia.Asset)
)

func init() {
	var err error

	tofuNFTABI, err = abi.JSON(strings.NewReader(tofunft.TofunftABI))
	if err != nil {
		panic(err)
	}

	_erc721ABI, err = abi.JSON(strings.NewReader(erc721.ERC721ABI))
	if err != nil {
		panic(err)
	}

	TofuNFT = utils.Getenv("SCRAPER_NAME_STATE", "")

	// If scraper state is not set yet, start from this block
	initBlockNumString := utils.Getenv("LAST_BLOCK_NUM", "")
	initBlockNum, err := strconv.ParseInt(initBlockNumString, 10, 64)
	if err != nil {
		log.Error("parse timeFinal: ", err)
	}
	defTofuNFTState.LastBlockNum = uint64(initBlockNum)

}

func NewTofuNFTScraper(rdb *models.RelDB, exchange dia.NFTExchange) (scraper *TofuNFTScraper) {
	switch exchange.BlockChain.Name {
	case dia.ASTAR:
		defTofuNFTConf.ContractAddr = "0x7Cae7FeB55349FeADB8f84468F692450D92597bc"
		scraper = makeNewTofuNFTScraper(exchange, rdb)
	case dia.BINANCESMARTCHAIN:
		defTofuNFTConf.ContractAddr = "0x449D05C544601631785a7C062DCDFF530330317e"
		scraper = makeNewTofuNFTScraper(exchange, rdb)
	}
	return
}

func makeNewTofuNFTScraper(exchange dia.NFTExchange, rdb *models.RelDB) *TofuNFTScraper {
	ctx := context.Background()

	eth, err := ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", ""))
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	s := &TofuNFTScraper{
		conf:  defTofuNFTConf,
		state: defTofuNFTState,
		tradeScraper: TradeScraper{
			shutdown:      make(chan nothing),
			shutdownDone:  make(chan nothing),
			datastore:     rdb,
			chanTrade:     make(chan dia.NFTTrade),
			source:        TofuNFT,
			ethConnection: eth,
		},
		exchangeName: exchange.Name,
		blockchain:   exchange.BlockChain.Name,
	}

	if err := s.initScraper(ctx); err != nil {
		log.Errorf("tofunft scraper could not be initialized: %s", err.Error())
		return nil
	}

	log.Infof("scraper %s starts at block: %v", TofuNFT, s.state.LastBlockNum)
	go s.mainLoop()

	return s
}

// init scraper
// if there are no values stored previously, use defaults and store them
func (s *TofuNFTScraper) initScraper(ctx context.Context) error {
	if err := s.loadConfig(ctx); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Errorf("unable to read scraper config from rdb: %s", err.Error())
			return err
		}

		// use & store defaults if there is no record in the scraper table

		defConf := *defTofuNFTConf
		s.conf = &defConf
		if err := s.tradeScraper.datastore.SetScraperConfig(ctx, TofuNFT, s.conf); err != nil {
			log.Errorf("unable to store scraper config on rdb: %s", err.Error())
			return err
		}

		defState := *defTofuNFTState
		s.state = &defState
		if err := s.tradeScraper.datastore.SetScraperState(ctx, TofuNFT, s.state); err != nil {
			log.Errorf("unable to store scraper state on rdb: %s", err.Error())
			return err
		}

		return nil
	}

	return s.loadState(ctx)
}

func (s *TofuNFTScraper) loadConfig(ctx context.Context) error {
	return s.tradeScraper.datastore.GetScraperConfig(ctx, TofuNFT, s.conf)
}

func (s *TofuNFTScraper) loadState(ctx context.Context) error {
	return s.tradeScraper.datastore.GetScraperState(ctx, TofuNFT, s.state)
}

func (s *TofuNFTScraper) storeState(ctx context.Context) error {
	return s.tradeScraper.datastore.SetScraperState(ctx, TofuNFT, s.state)
}

func (s *TofuNFTScraper) mainLoop() {
	defer func() {
		s.tradeScraper.closed = true

		close(s.tradeScraper.chanTrade)
		close(s.tradeScraper.shutdownDone)
	}()

	log.Infof("tofunft scraper has been started (batch: %d, period: %s)", s.conf.BatchSize, s.conf.WaitPeriod.String())

	for stop := false; !stop; {
		if err := s.FetchTrades(); err != nil {
			if errors.Is(err, errTofuNFTShutdownRequest) {
				stop = true
				continue
			}
		}

		log.Debugf("wait for %s", s.conf.WaitPeriod)

		select {
		case <-time.After(s.conf.WaitPeriod):
		case <-s.tradeScraper.shutdown:
			stop = true
		}
	}
}

// FetchTrades searches for trades on-chain by the next block range
func (s *TofuNFTScraper) FetchTrades() error {
	var err error

	// TODO: make FetchTrades context-aware
	ctx := context.Background()

	// it must be run once at a time
	s.mu.Lock()
	defer s.mu.Unlock()

	// read config
	if err = s.loadConfig(ctx); err != nil {
		log.Warnf("unable to load scraper config: %s", err.Error())
		return err
	}

	// read state
	if err = s.loadState(ctx); err != nil {
		log.Warnf("unable to load scraper state: %s", err.Error())
		return err
	}

	log.Infof("fetching tofunft trade transactions from block %d(+%d)", s.state.LastBlockNum, s.conf.BatchSize)

	// fetch trade transactions
	res, err := utils.EthFilterTXs(ctx, s.tradeScraper.ethConnection, utils.EthTxFilterCriteria{
		StartBlockNum:      s.state.LastBlockNum,
		StartTxIndex:       s.state.LastTxIndex,
		LimitBlocks:        s.conf.BatchSize,
		BehindHighestBlock: s.conf.FollowDist,
		EvAddrs:            []common.Address{common.HexToAddress(s.conf.ContractAddr)},
		Events:             []common.Hash{tofuNFTABI.Events["EvInventoryUpdate"].ID},
	})

	if err != nil {
		log.Warnf("unable to filter tofunft trades: %s", err.Error())
		return err
	}

	log.Infof("found %d trade(logs: %d) transactions in %d blocks(from %d [tx index offset: %d] to %d, sync: %t[stay behind: -%d]), exploring details...", res.NumTXs, res.NumLogs, res.NumBlocks, s.state.LastBlockNum, s.state.LastTxIndex, res.LastBlockNum, res.Synced, s.conf.FollowDist)

	numTrades := 0

	// process trade transactions
	for _, tx := range res.TXs {
		s.state.LastBlockNum = tx.BlockNum
		s.state.LastTxIndex = tx.TXIndex
		s.state.LastErr = ""
		log.Info("current state.ErrCounter: ", s.state.ErrCounter)

		skipped, err := s.processTx(ctx, tx)
		if err != nil {
			s.state.ErrCounter++

			if s.state.ErrCounter <= s.conf.MaxRetry {
				s.state.LastErr = fmt.Sprintf("unable to process trade transaction(%s): %s", tx.TXHash.Hex(), err.Error())
				log.Error(s.state.LastErr)
				// store state
				if errState := s.storeState(ctx); errState != nil {
					log.Warnf("unable to store scraper state: %s", errState.Error())
					return errState
				}
				return err
			}

			log.Warnf("SKIPPING PERMANENTLY! block: %d, tx index: %d - error: %s", s.state.LastBlockNum, s.state.LastTxIndex, err.Error())
		}

		if !skipped {
			numTrades++
		}

		// reset consecutive error counter
		s.state.ErrCounter = 0

		// move next
		s.state.LastTxIndex = tx.TXIndex + 1

		// store state
		if err := s.storeState(ctx); err != nil {
			log.Warnf("unable to store scraper state: %s", err.Error())
			return err
		}
	}

	s.state.LastBlockNum = res.LastBlockNum + 1
	s.state.LastTxIndex = 0

	if err := s.storeState(ctx); err != nil {
		log.Warnf("unable to store scraper state: %s", err.Error())
		return err
	}

	log.Infof("processed %d trades", numTrades)

	return nil
}

func (s *TofuNFTScraper) processTx(ctx context.Context, tx *utils.EthFilteredTx) (bool, error) {
	log.Tracef("process tx -> block: %d, tx index: %d, tx hash: %s", tx.BlockNum, tx.TXIndex, tx.TXHash.Hex())

	var ev tofunft.TofunftEvInventoryUpdate
	err := tofuNFTABI.UnpackIntoInterface(&ev, "EvInventoryUpdate", tx.Logs[0].Data)
	if err != nil {
		log.Errorf("unable to read EvInventoryUpdate log from transaction(%s)", tx.TXHash)
		return false, err
	}

	receipt, err := s.tradeScraper.ethConnection.TransactionReceipt(ctx, tx.TXHash)
	if err != nil {
		log.Errorf("unable to read transaction(%s) receipt: %s", tx.TXHash, err.Error())
		return false, err
	}

	currSymbol := "ASTR"
	currAddr := ev.Inventory.Currency
	currDecimals := 18

	transfers, err := s.findERC721Transfers(ctx, receipt)
	if err != nil {
		log.Errorf("unable to find transfers of the event(block: %d, tx index: %d, tx: %s): %s", tx.BlockNum, tx.TXIndex, tx.TXHash.Hex(), err.Error())
		return false, err
	}

	// skip if the event has no transfer
	if len(transfers) == 0 {
		log.Tracef("event(block: %d, tx index: %d, tx: %s) skipped due to it has no erc721 transfer log", tx.BlockNum, tx.TXIndex, tx.TXHash.Hex())
		return true, nil
	}

	// skip if the event has multiple transfers due to we can't calculate the price of trade
	if len(transfers) > 1 {
		log.Tracef("event(block: %d, tx index: %d, tx: %s) skipped due to it has multiple erc721 transfer logs", tx.BlockNum, tx.TXIndex, tx.TXHash.Hex())
		return true, nil
	}

	normPrice := decimal.NewFromBigInt(ev.Inventory.Price, 0).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(currDecimals))))

	usdPrice, err := s.calcUSDPrice(tx.BlockNum, currAddr, currSymbol, normPrice)
	if err != nil {
		log.Errorf("unable to calculate usd price of the event(block: %d, log: %d, tx: %s): %s", tx.BlockNum, tx.TXIndex, tx.TXHash.Hex(), err.Error())
		return false, err
	}

	if err := s.notifyTrade(tx, &ev, transfers[0], ev.Inventory.Price, normPrice, usdPrice, currSymbol, currAddr); err != nil {
		if !errors.Is(err, errTofuNFTShutdownRequest) {
			log.Warnf("event(block: %d, tx index: %d, tx: %s) couldn't processed: %s", tx.BlockNum, tx.TXIndex, tx.TXHash.Hex(), err.Error())
		}

		return false, err
	}

	return false, nil
}

func (s *TofuNFTScraper) notifyTrade(tx *utils.EthFilteredTx, ev *tofunft.TofunftEvInventoryUpdate, transfer *astarERC721Transfer, price *big.Int, priceDec decimal.Decimal, usdPrice float64, currSymbol string, currAddr common.Address) error {
	nftClass, err := s.createOrReadNFTClass(transfer)
	if err != nil {
		return err
	}

	nft, err := s.createOrReadNFT(nftClass, transfer)
	if err != nil {
		return err
	}

	// Get block time.
	block, err := s.tradeScraper.ethConnection.BlockByNumber(context.Background(), big.NewInt(int64(tx.BlockNum)))
	if err != nil {
		log.Errorf("getting block: %+v", err)
	}

	trade := dia.NFTTrade{
		NFT:         *nft,
		Price:       price,
		PriceUSD:    usdPrice,
		FromAddress: transfer.From.Hex(),
		ToAddress:   transfer.To.Hex(),
		BlockNumber: tx.BlockNum,
		Timestamp:   time.Unix(int64(block.Time()), 0),
		TxHash:      tx.TXHash.Hex(),
		Exchange:    s.exchangeName,
	}

	if asset, ok := assetCacheTofuAstar[s.blockchain+"-"+currAddr.Hex()]; ok {
		trade.Currency = asset
	} else {
		currency, err := s.tradeScraper.datastore.GetAsset(currAddr.Hex(), s.blockchain)
		if err != nil {
			log.Errorf("cannot fetch asset %s -- %s", s.blockchain, currAddr.Hex())
		}
		trade.Currency = currency
		assetCacheTofuAstar[s.blockchain+"-"+currAddr.Hex()] = currency
	}

	fmt.Println("found trade: ", trade)

	// handle close request if the chanTrade not consumed immediately
	select {
	case s.tradeScraper.chanTrade <- trade:
	case <-s.tradeScraper.shutdown:
		return errTofuNFTShutdownRequest
	}

	return nil
}

func (s *TofuNFTScraper) createOrReadNFTClass(transfer *astarERC721Transfer) (*dia.NFTClass, error) {
	nftClass, err := s.tradeScraper.datastore.GetNFTClass(transfer.NFTAddress.Hex(), s.blockchain)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Warnf("unable to read nftclass from reldb: %s", err.Error())
			return nil, err
		}

		nftClass = dia.NFTClass{
			Address:      transfer.NFTAddress.Hex(),
			Blockchain:   s.blockchain,
			ContractType: tofuNFTNFTContractType,
		}

		if transfer.Name != nil {
			nftClass.Name = *transfer.Name
		}

		if transfer.Symbol != nil {
			nftClass.Symbol = *transfer.Symbol
		}

		if err = s.tradeScraper.datastore.SetNFTClass(nftClass); err != nil {
			log.Warnf("unable to create nftclass on reldb: %s", err.Error())
			return nil, err
		}
	}

	return &nftClass, nil
}

func (s *TofuNFTScraper) createOrReadNFT(nftClass *dia.NFTClass, transfer *astarERC721Transfer) (*dia.NFT, error) {
	nft, err := s.tradeScraper.datastore.GetNFT(nftClass.Address, s.blockchain, transfer.TokenID.String())
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Warnf("unable to read nft from reldb: %s", err.Error())
			return nil, err
		}

		createdBy, createdAt, err := s.findContractCreationInfo(context.Background(), common.HexToAddress(nftClass.Address))
		if err != nil {
			log.Warnf("unable to find the creation info for the nft contract(%s): %s", nftClass.Address, err.Error())
			return nil, err
		}

		nft = dia.NFT{
			NFTClass:       *nftClass,
			TokenID:        transfer.TokenID.String(),
			CreationTime:   createdAt,
			CreatorAddress: createdBy.Hex(),
			Attributes:     transfer.TokenAttrs,
		}

		if transfer.TokenURI != nil {
			nft.URI = *transfer.TokenURI
		}

		if err = s.tradeScraper.datastore.SetNFT(nft); err != nil {
			log.Warnf("unable to create nft on reldb: %s", err.Error())
			return nil, err
		}
	}

	return &nft, nil
}

func (s *TofuNFTScraper) calcUSDPrice(blockNum uint64, tokenAddr common.Address, symbol string, price decimal.Decimal) (float64, error) {
	tokenPrice, err := s.findPrice(blockNum, tokenAddr, symbol)
	if err != nil {
		return 0, err
	}

	usdPrice := price.Mul(tokenPrice)

	// using float type is not a good idea to handle prices
	// we ignore if the price cannot be presentable as float64
	f, _ := usdPrice.Float64()

	return f, nil
}

func (s *TofuNFTScraper) findPrice(blockNum uint64, tokenAddr common.Address, symbol string) (decimal.Decimal, error) {
	// TODO: find the token price in usd for the given block number
	switch symbol {
	case "ASTR", "WETH", "WASTR":
		return decimal.NewFromString("2.59")

	default:
		return decimal.NewFromString("1")
	}
}

// GetDataChannel returns the scrapers data channel.
func (s *TofuNFTScraper) GetTradeChannel() chan dia.NFTTrade {
	return s.tradeScraper.chanTrade
}

func (s *TofuNFTScraper) Close() error {
	if s.tradeScraper.closed {
		return errors.New("scraper already closed")
	}

	close(s.tradeScraper.shutdown)

	return nil
}

// it searches the creation transaction for the given contract address using binary search in complexity of o(log n)
func (s *TofuNFTScraper) findContractCreationInfo(ctx context.Context, contractAddr common.Address) (createdBy common.Address, createdAt time.Time, err error) {
	if !s.conf.UseArchiveNode {
		log.Trace("nft contract creation info could not found because UseArchiveNode flag is not set, using zero values")
		return common.Address{}, time.Time{}, nil
	}

	var (
		lo, hi, blockNum uint64
		code             []byte
		receipt          *types.Receipt
		chainID          *big.Int
		block            *types.Block
	)

	hi, err = s.tradeScraper.ethConnection.BlockNumber(ctx)
	if err != nil {
		return
	}

	for lo <= hi {
		blockNum = (lo + hi) / 2

		code, err = s.tradeScraper.ethConnection.CodeAt(ctx, contractAddr, new(big.Int).SetUint64(blockNum))
		if err != nil {
			return
		}

		if len(code) == 0 {
			lo = blockNum
		} else {
			hi = blockNum
		}

		if hi == lo+1 {
			blockNum = hi
			break
		}
	}

	block, err = s.tradeScraper.ethConnection.BlockByNumber(ctx, new(big.Int).SetUint64(blockNum))
	if err != nil {
		return
	}

	chainID, err = s.tradeScraper.ethConnection.NetworkID(ctx)
	if err != nil {
		return
	}

	signer := types.NewEIP155Signer(chainID)

	for _, trx := range block.Transactions() {
		// recipient must be nill for contract creation transactions
		if trx.To() != nil {
			continue
		}

		receipt, err = s.tradeScraper.ethConnection.TransactionReceipt(ctx, trx.Hash())
		if err != nil {
			return
		}

		// note that if the nft was created by another smart contract
		// we can't find its creation info with this method
		if bytes.Equal(receipt.ContractAddress.Bytes(), contractAddr.Bytes()) {
			createdAt = time.Unix(int64(block.Time()), 0).UTC()
			createdBy, err = types.Sender(signer, trx)
			if err != nil {
				return
			}

			return
		}
	}

	return
}

// it finds the transfer events of ERC721 in the given transaction
func (s *TofuNFTScraper) findERC721Transfers(ctx context.Context, receipt *types.Receipt) ([]*astarERC721Transfer, error) {
	transfers := make([]*astarERC721Transfer, 0, 1)

	for _, txLog := range receipt.Logs {
		// Erc721 Transfers have 4 indexed topics.
		if len(txLog.Topics) != 4 || txLog.Topics[0] != _erc721ABI.Events["Transfer"].ID {
			continue
		}

		nft, err := erc721.NewERC721(txLog.Address, s.tradeScraper.ethConnection)
		if err != nil {
			log.Warnf("unable to bind erc721 contract at address %s: %s", txLog.Address.Hex(), err.Error())
			continue
		}

		transferLog, err := nft.ParseTransfer(*txLog)
		if err != nil {
			// it means this log data not comply to erc721's transfer event
			//
			// some old erc721 contracts have unindexed tokenid parameter
			// so it is not compliant with the eip-721.

			// best effort...
			compat, err := erc721.NewERC721Compat(txLog.Address, s.tradeScraper.ethConnection)
			if err != nil {
				log.Warnf("unable to bind erc721compat contract at address %s: %s", txLog.Address.Hex(), err.Error())
				continue
			}

			compatLog, err := compat.ParseTransfer(*txLog)
			if err != nil {
				log.Tracef("the event cannot comply to erc721's transfer: %s", err)
				continue
			}

			transferLog = &erc721.ERC721Transfer{
				From:    compatLog.From,
				To:      compatLog.To,
				TokenId: compatLog.TokenId,
				Raw:     compatLog.Raw,
			}
		}

		transfer := &astarERC721Transfer{
			NFTAddress: txLog.Address,
			From:       transferLog.From,
			To:         transferLog.To,
			TokenID:    transferLog.TokenId,
			TokenAttrs: make(map[string]interface{}),
		}

		callOpts := &bind.CallOpts{Context: ctx}

		if s.conf.UseArchiveNode {
			callOpts.BlockNumber = new(big.Int).SetUint64(txLog.BlockNumber)
		}

		if md, err := erc721.NewERC721Metadata(txLog.Address, s.tradeScraper.ethConnection); err != nil {
			log.Warnf("unable to bind erc721 metadata contract at address %s: %s", txLog.Address.Hex(), err.Error())
		} else {
			if nftName, err := md.Name(callOpts); err != nil {
				log.Warnf("unable to read nft name from metadata interface of erc721(addr: %s): %s", txLog.Address.Hex(), err.Error())
			} else {
				transfer.Name = &nftName
			}

			if nftSymbol, err := md.Symbol(callOpts); err != nil {
				log.Warnf("unable to read nft symbol from metadata interface of nft(addr: %s): %s", txLog.Address.Hex(), err.Error())
			} else {
				transfer.Symbol = &nftSymbol
			}

			if tokenURI, err := md.TokenURI(callOpts, transfer.TokenID); err != nil {
				log.Warnf("unable to find token(%s) uri: %s", transfer.TokenID.String(), err.Error())
			} else if attrs, err := s.readNFTAttr(ctx, tokenURI); err != nil {
				log.Warnf("unable to read token(%s) attributes: %s", transfer.TokenID.String(), err.Error())
			} else {
				transfer.TokenURI = &tokenURI
				transfer.TokenAttrs = attrs
			}
		}

		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

func (s *TofuNFTScraper) readNFTAttr(ctx context.Context, uri string) (map[string]interface{}, error) {
	if uri == "" {
		return nil, nil
	}

	if strings.HasPrefix(uri, "ipfs://") {
		// TODO: add IPFS support
		return nil, nil
	}

	attrs := make(map[string]interface{})

	ctx, cancel := context.WithTimeout(ctx, s.conf.MetadataTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			// Handle the error from closing the response body
			log.Println("Error closing response body:", cerr)
		}
	}()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New("unable to read token attributes: " + resp.Status)
	}

	if err := json.NewDecoder(io.LimitReader(resp.Body, int64(s.conf.MaxMetadataSize))).Decode(&attrs); err != nil {
		return nil, err
	}

	return attrs, nil
}
