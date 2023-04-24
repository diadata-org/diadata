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

	"github.com/diadata-org/diadata/config/nftContracts/blurv1"
	"github.com/diadata-org/diadata/config/nftContracts/erc1155"
	"github.com/diadata-org/diadata/config/nftContracts/erc20"
	"github.com/diadata-org/diadata/config/nftContracts/erc721"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
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
	// we assume all of the NFTs traded on BlurV1 are ERC721(1155 is an extension of it)
	BlurV1NFTContractType = "ERC721"
)

type BlurV1ScraperConfig struct {
	// BlurV1's exchange contract address on connected blockchain network
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

type BlurV1ScraperState struct {
	// last block number has been processed
	LastBlockNum uint64 `json:"last_block_num"`

	// last transaction index in the block(curr) has been processed
	LastTxIndex uint `json:"last_tx_index"`

	// holds the latest error message that occurred while scraping
	LastErr string `json:"last_error"`

	// indicates the number of consecutive error, reset on any successful operation
	ErrCounter int `json:"count_of_error"`
}

type BlurV1Scraper struct {
	tradeScraper TradeScraper

	mu       sync.Mutex
	conf     *BlurV1ScraperConfig
	state    *BlurV1ScraperState
	exchange dia.NFTExchange
}

type BlurV1Erc721Transfer struct {
	NFTAddress  common.Address
	Name        *string
	Symbol      *string
	TotalSupply *big.Int
	From        common.Address
	To          common.Address
	TokenID     *big.Int
	Price       *big.Int
	PriceUSD    float64
	TokenURI    *string
	TokenAttrs  map[string]interface{}
}

var (
	errBlurV1ShutdownRequest = errors.New("shutdown requested")

	// default values are valid for the first run which is it saves
	// these configs to the DB
	defBlurV1Conf = &BlurV1ScraperConfig{
		ContractAddr:    "0x000000000000Ad05Ccc4F10045630fb830B95127", // Blur V1
		BatchSize:       5000,
		WaitPeriod:      2 * time.Second,
		FollowDist:      10,
		UseArchiveNode:  false,
		MaxRetry:        1,
		SkipOnErr:       true,
		MaxMetadataSize: 50 * 1024,
		MetadataTimeout: 30 * time.Second,
	}

	// BlurV1 market contract has been deployed on the mainnet at
	// block num 15779579, so scraper starts from this block.
	// https://etherscan.io/tx/0x157bdba2b4758a0fae53060ad4c68daf2af43571b1e7002161cbca1ffc5bef15
	defBlurV1State = &BlurV1ScraperState{LastBlockNum: 15779579}

	// This string is the identifier of the scraper in conf and state fields in postgres.
	BlurV1 = "BlurV1"

	BlurV1ABI abi.ABI

	assetCacheBlurV1 = make(map[string]dia.Asset)
)

func init() {
	var err error

	BlurV1ABI, err = abi.JSON(strings.NewReader(blurv1.Blurv1ABI))
	if err != nil {
		panic(err)
	}

	erc20ABI, err = abi.JSON(strings.NewReader(erc20.ERC20ABI))
	if err != nil {
		panic(err)
	}

	erc721ABI, err = abi.JSON(strings.NewReader(erc721.ERC721ABI))
	if err != nil {
		panic(err)
	}

	erc1155ABI, err = abi.JSON(strings.NewReader(erc1155.Erc1155ABI))
	if err != nil {
		panic(err)
	}

	BlurV1 = utils.Getenv("SCRAPER_NAME_STATE", BlurV1)

	// If scraper state is not set yet, start from this block
	initBlockNumString := utils.Getenv("LAST_BLOCK_NUM", "15779579")
	initBlockNum, err := strconv.ParseInt(initBlockNumString, 10, 64)
	if err != nil {
		log.Error("parse timeFinal: ", err)
	}
	defBlurV1State.LastBlockNum = uint64(initBlockNum)

}

func NewBlurV1Scraper(rdb *models.RelDB, exchange dia.NFTExchange) *BlurV1Scraper {
	ctx := context.Background()

	eth, err := ethclient.Dial(utils.Getenv("ETH_URI_REST", ""))
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	s := &BlurV1Scraper{
		conf:     &BlurV1ScraperConfig{},
		state:    &BlurV1ScraperState{},
		exchange: exchange,
		tradeScraper: TradeScraper{
			shutdown:      make(chan nothing),
			shutdownDone:  make(chan nothing),
			datastore:     rdb,
			chanTrade:     make(chan dia.NFTTrade),
			source:        exchange.Name,
			ethConnection: eth,
		},
	}

	if err := s.initScraper(ctx); err != nil {
		log.Errorf("BlurV1 scraper could not be initialized: %s", err.Error())
		return nil
	}

	log.Info("scraper %s starts at block: %s", BlurV1, s.state.LastBlockNum)
	go s.mainLoop()

	return s
}

// init scraper
// if there are no values stored previously, use defaults and store them
func (s *BlurV1Scraper) initScraper(ctx context.Context) error {
	if err := s.loadConfig(ctx); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Errorf("unable to read scraper config from rdb: %s", err.Error())
			return err
		}

		// use & store defaults if there is no record in the scraper table

		defConf := *defBlurV1Conf // copy
		s.conf = &defConf
		if err := s.tradeScraper.datastore.SetScraperConfig(ctx, BlurV1, s.conf); err != nil {
			log.Errorf("unable to store scraper config on rdb: %s", err.Error())
			return err
		}

		defState := *defBlurV1State // copy
		s.state = &defState
		if err := s.tradeScraper.datastore.SetScraperState(ctx, BlurV1, s.state); err != nil {
			log.Errorf("unable to store scraper state on rdb: %s", err.Error())
			return err
		}

		return nil
	}

	return s.loadState(ctx)
}

func (s *BlurV1Scraper) loadConfig(ctx context.Context) error {
	return s.tradeScraper.datastore.GetScraperConfig(ctx, BlurV1, s.conf)
}

func (s *BlurV1Scraper) loadState(ctx context.Context) error {
	return s.tradeScraper.datastore.GetScraperState(ctx, BlurV1, s.state)
}

func (s *BlurV1Scraper) storeState(ctx context.Context) error {
	return s.tradeScraper.datastore.SetScraperState(ctx, BlurV1, s.state)
}

func (s *BlurV1Scraper) mainLoop() {
	defer func() {
		s.tradeScraper.closed = true

		close(s.tradeScraper.chanTrade)
		close(s.tradeScraper.shutdownDone)
	}()

	log.Infof("BlurV1 scraper has been started (batch: %d, period: %s)", s.conf.BatchSize, s.conf.WaitPeriod.String())

	for stop := false; !stop; {
		if err := s.FetchTrades(); err != nil {
			if errors.Is(err, errBlurV1ShutdownRequest) {
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
func (s *BlurV1Scraper) FetchTrades() error {
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

	log.Infof("fetching BlurV1 trade transactions from block %d(+%d)", s.state.LastBlockNum, s.conf.BatchSize)

	// fetch trade transactions
	res, err := utils.EthFilterTXs(ctx, s.tradeScraper.ethConnection, utils.EthTxFilterCriteria{
		StartBlockNum:      s.state.LastBlockNum,
		StartTxIndex:       s.state.LastTxIndex,
		LimitBlocks:        s.conf.BatchSize,
		BehindHighestBlock: s.conf.FollowDist,
		EvAddrs:            []common.Address{common.HexToAddress(s.conf.ContractAddr)},
		Events:             []common.Hash{BlurV1ABI.Events["OrdersMatched"].ID},
	})

	if err != nil {
		log.Warnf("unable to filter BlurV1 trades: %s", err.Error())
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

func (s *BlurV1Scraper) processTx(ctx context.Context, tx *utils.EthFilteredTx) (bool, error) {
	log.Tracef("process tx -> block: %d, tx index: %d, tx hash: %s", tx.BlockNum, tx.TXIndex, tx.TXHash.Hex())

	marketContract, err := blurv1.NewBlurv1(tx.Logs[0].Address, s.tradeScraper.ethConnection)
	if err != nil {
		log.Errorf("unable to make new market contract for address: %s", tx.Logs[0].Address.Hex())
		return false, err
	}

	ev, err := marketContract.ParseOrdersMatched(tx.Logs[0])
	if err != nil {
		log.Errorf("unable to decode BlurV1 OrdersMatched event(tx: %s, logIndex: %d) (SKIPPED!): %s", tx.TXHash, tx.Logs[0].Index, err.Error())
		return true, nil // skip
	}

	_, pending, err := s.tradeScraper.ethConnection.TransactionByHash(ctx, tx.TXHash)
	if err != nil {
		log.Errorf("unable to read transaction(%s): %s", tx.TXHash, err.Error())
		return false, err

	} else if pending {
		err = fmt.Errorf("transaction(%s) status error: pending=true", tx.TXHash)
		log.Error(err.Error())
		return false, err
	}

	receipt, err := s.tradeScraper.ethConnection.TransactionReceipt(ctx, tx.TXHash)
	if err != nil {
		log.Errorf("unable to read transaction(%s) receipt: %s", tx.TXHash, err.Error())
		return false, err
	}
	// Get block time.
	timestamp, err := ethhelper.GetBlockTimeEth(int64(ev.Raw.BlockNumber), s.tradeScraper.datastore, s.tradeScraper.ethConnection)
	if err != nil {
		log.Errorf("getting block time: %+v", err)
	}

	transfers, err := s.findERC721Transfers(ctx, receipt, timestamp)
	if err != nil {
		log.Errorf("unable to find transfers of the event(block: %d, tx index: %d, tx: %s): %s", ev.Raw.BlockNumber, ev.Raw.TxIndex, ev.Raw.TxHash.Hex(), err.Error())
		return false, err
	}

	currAddr := common.Address{}
	for _, transfer := range transfers {
		if err := s.notifyTrade(ev, transfer, currAddr, len(transfers) > 1, timestamp); err != nil {
			if !errors.Is(err, errBlurV1ShutdownRequest) {
				log.Warnf("event(block: %d, tx index: %d, tx: %s) couldn't processed: %s", ev.Raw.BlockNumber, ev.Raw.TxIndex, ev.Raw.TxHash.Hex(), err.Error())
			} else {
				log.Error("notifyTradeErr: ", err)
			}
		}
	}

	return false, nil
}

func (s *BlurV1Scraper) notifyTrade(
	ev *blurv1.Blurv1OrdersMatched,
	erc721Transfer *BlurV1Erc721Transfer,
	currAddr common.Address,
	isBulk bool,
	timestamp time.Time,
) error {

	nftClass, err := s.createOrReadNFTClass(erc721Transfer)
	if err != nil {
		return err
	}

	nft, err := s.createOrReadNFT(nftClass, erc721Transfer)
	if err != nil {
		return err
	}

	trade := dia.NFTTrade{
		NFT:         *nft,
		Price:       erc721Transfer.Price,
		PriceUSD:    erc721Transfer.PriceUSD,
		FromAddress: erc721Transfer.From.Hex(),
		ToAddress:   erc721Transfer.To.Hex(),
		BlockNumber: ev.Raw.BlockNumber,
		Timestamp:   timestamp,
		TxHash:      ev.Raw.TxHash.Hex(),
		Exchange:    s.exchange.Name,
	}

	if isBulk {
		trade.BundleSale = true
	}

	if asset, ok := assetCacheBlurV1[dia.ETHEREUM+"-"+currAddr.Hex()]; ok {
		trade.Currency = asset
	} else {
		currency, err := s.tradeScraper.datastore.GetAsset(currAddr.Hex(), dia.ETHEREUM)
		if err != nil {
			log.Errorf("cannot fetch asset %s -- %s", dia.ETHEREUM, currAddr.Hex())
		}
		trade.Currency = currency
		assetCacheBlurV1[dia.ETHEREUM+"-"+currAddr.Hex()] = currency
	}
	fmt.Println("found trade: ", trade)

	// handle close request if the chanTrade not consumed immediately
	select {
	case s.tradeScraper.chanTrade <- trade:
	case <-s.tradeScraper.shutdown:
		return errBlurV1ShutdownRequest
	}

	return nil
}

func (s *BlurV1Scraper) createOrReadNFTClass(transfer *BlurV1Erc721Transfer) (*dia.NFTClass, error) {
	nftClass, err := s.tradeScraper.datastore.GetNFTClass(transfer.NFTAddress.Hex(), dia.ETHEREUM)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Warnf("unable to read nftclass from reldb: %s", err.Error())
			return nil, err
		}

		nftClass = dia.NFTClass{
			Address:      transfer.NFTAddress.Hex(),
			Blockchain:   dia.ETHEREUM,
			ContractType: BlurV1NFTContractType,
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

func (s *BlurV1Scraper) createOrReadNFT(nftClass *dia.NFTClass, transfer *BlurV1Erc721Transfer) (*dia.NFT, error) {
	nft, err := s.tradeScraper.datastore.GetNFT(nftClass.Address, dia.ETHEREUM, transfer.TokenID.String())
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

func (s *BlurV1Scraper) calcUSDPrice(symbol string, price decimal.Decimal) (float64, error) {
	tokenPrice, err := s.findPrice(symbol)
	if err != nil {
		return 0, err
	}

	usdPrice := price.Mul(tokenPrice)

	// using float type is not a good idea to handle prices
	// we ignore if the price cannot be presentable as float64
	f, _ := usdPrice.Float64()

	return f, nil
}

func (s *BlurV1Scraper) findPrice(symbol string) (decimal.Decimal, error) {
	// TODO: find the token price in usd for the given block number
	switch symbol {
	case "ETH", "WETH":
		return decimal.NewFromString("2040.0910")

	default:
		return decimal.NewFromString("1")
	}
}

// GetDataChannel returns the scrapers data channel.
func (s *BlurV1Scraper) GetTradeChannel() chan dia.NFTTrade {
	return s.tradeScraper.chanTrade
}

func (s *BlurV1Scraper) Close() error {
	if s.tradeScraper.closed {
		return errors.New("scraper already closed")
	}

	close(s.tradeScraper.shutdown)

	return nil
}

// it searches the creation transaction for the given contract address using binary search in complexity of o(log n)
func (s *BlurV1Scraper) findContractCreationInfo(ctx context.Context, contractAddr common.Address) (createdBy common.Address, createdAt time.Time, err error) {
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

func (s *BlurV1Scraper) findERC20Transfers(ctx context.Context, receipt *types.Receipt, filterAmount *big.Int) ([]*erc20Transfer, error) {
	transfers := make([]*erc20Transfer, 0, 2)

	for _, txLog := range receipt.Logs {
		if len(txLog.Topics) < 1 || txLog.Topics[0] != erc20ABI.Events["Transfer"].ID {
			continue
		}

		token, err := erc20.NewERC20(txLog.Address, s.tradeScraper.ethConnection)
		if err != nil {
			log.Warnf("unable to bind erc720 contract at address %s: %s", txLog.Address.Hex(), err.Error())
			continue
		}

		ev, err := token.ParseTransfer(*txLog)
		if err != nil {
			log.Tracef("the event cannot comply to erc20's transfer: %s", err)
			continue
		}

		if filterAmount == nil {
			continue
		}

		transfer := &erc20Transfer{
			TokenAddr: txLog.Address,
			From:      ev.From,
			To:        ev.To,
			Amount:    ev.Value,
			Decimals:  1,
		}

		transfers = append(transfers, transfer)
		callOpts := &bind.CallOpts{Context: ctx}

		if s.conf.UseArchiveNode {
			callOpts.BlockNumber = new(big.Int).SetUint64(txLog.BlockNumber)
		}

		var ethSymbol = "ETH"
		transfer.TokenSymbol = &ethSymbol
		transfer.Decimals = 18
	}

	return transfers, nil
}

// it finds the transfer events of ERC721 in the given transaction
func (s *BlurV1Scraper) findERC721Transfers(ctx context.Context, receipt *types.Receipt, timestamp time.Time) ([]*BlurV1Erc721Transfer, error) {
	transfers := make([]*BlurV1Erc721Transfer, 0, 1)
	transfers_uniqe_id_timestamp := make(map[string]string)
	nft_prices := make(map[string]*big.Int)
	// save the price for each nft
	for _, txLog := range receipt.Logs {
		if len(txLog.Topics) > 0 && txLog.Topics[0].Hex() == "0x61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f64" {
			marketContract, err := blurv1.NewBlurv1(txLog.Address, s.tradeScraper.ethConnection)
			if err != nil {
				log.Errorf("unable to make new market contract for address: %s", txLog.Address.Hex())
				continue
			}
			ev, err := marketContract.ParseOrdersMatched(*txLog)
			if err != nil {
				log.Errorf(
					"unable to decode BlurV1 OrdersMatched event(tx: %s, logIndex: %d) (SKIPPED!): %s",
					txLog.TxHash, txLog.Index, err.Error(),
				)
				continue
			}
			nft_prices[ev.Sell.TokenId.String()+"-"+ev.Sell.Collection.Hex()] = ev.Sell.Price
		}
	}
	for _, txLog := range receipt.Logs {
		if len(txLog.Topics) < 1 || txLog.Topics[0] != erc721ABI.Events["Transfer"].ID {
			continue
		}
		// skip internal transfers of blur
		if txLog.Address.Hex() == "0x0000000000A39bb272e79075ade125fd351887Ac" {
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
			compat, errCompat := erc721.NewERC721Compat(txLog.Address, s.tradeScraper.ethConnection)
			if errCompat != nil {
				log.Warnf("unable to bind erc721compat contract at address %s: %s", txLog.Address.Hex(), errCompat.Error())
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
		key_in_nft_prices := transferLog.TokenId.String() + "-" + txLog.Address.Hex()
		price, price_ok := nft_prices[key_in_nft_prices]
		if !price_ok {
			// skip over nfts which transferred over other marketplaces
			continue
		}
		currSymbol := "ETH"
		currDecimals := 18
		normPrice := decimal.NewFromBigInt(price, 0).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(currDecimals))))
		usdPrice, err := s.calcUSDPrice(currSymbol, normPrice)
		if err != nil {
			log.Errorf("unable to calculate usd price of the event(block: %d, log: %d, tx: %s): %s", transferLog.Raw.BlockNumber, transferLog.Raw.TxIndex, transferLog.Raw.TxHash.Hex(), err.Error())
			continue
		}

		transfer := &BlurV1Erc721Transfer{
			NFTAddress: txLog.Address,
			From:       transferLog.From,
			To:         transferLog.To,
			TokenID:    transferLog.TokenId,
			Price:      price,
			PriceUSD:   usdPrice,
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
		// sometimes there are more than one transfer log per erc271transfer
		key := transfer.TokenID.String() + timestamp.String()
		_, ok := transfers_uniqe_id_timestamp[key]
		if !ok {
			transfers = append(transfers, transfer)
			transfers_uniqe_id_timestamp[key] = key
		}
	}

	return transfers, nil
}

func (s *BlurV1Scraper) readNFTAttr(ctx context.Context, uri string) (map[string]interface{}, error) {
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

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New("unable to read token attributes: " + resp.Status)
	}

	if err := json.NewDecoder(io.LimitReader(resp.Body, int64(s.conf.MaxMetadataSize))).Decode(&attrs); err != nil {
		return nil, err
	}

	return attrs, nil
}
