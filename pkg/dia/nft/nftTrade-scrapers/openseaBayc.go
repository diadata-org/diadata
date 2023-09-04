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
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/config/nftContracts/erc20"
	"github.com/diadata-org/diadata/config/nftContracts/erc721"
	"github.com/diadata-org/diadata/config/nftContracts/opensea"
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
	OpenSeaBAYC = "OpenSeaBayc"
)

type OpenSeaBAYCScraperConfig struct {
	// opensea's exchange contract address on connected blockchain network
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

type OpenSeaBAYCScraperState struct {
	// last block number has been processed
	LastBlockNum uint64 `json:"last_block_num"`

	// last transaction index in the block(curr) has been processed
	LastTxIndex uint `json:"last_tx_index"`

	// holds the latest error message that occurred while scraping
	LastErr string `json:"last_error"`

	// indicates the number of consecutive error, reset on any successful operation
	ErrCounter int `json:"count_of_error"`
}

type OpenSeaBAYCScraper struct {
	tradeScraper TradeScraper

	mu       sync.Mutex
	conf     *OpenSeaBAYCScraperConfig
	state    *OpenSeaBAYCScraperState
	exchange dia.NFTExchange
}

var (

	// default values are valid for the first run which is it saves
	// these configs to the DB
	defOpenSeaBAYCConf = &OpenSeaBAYCScraperConfig{
		// ContractAddr:    "0x7be8076f4ea4a4ad08075c2508e481d6c946d12b", // Wyvern V1
		ContractAddr:    "0x7f268357A8c2552623316e2562D90e642bB538E5", // Wyvern V2
		BatchSize:       5000,
		WaitPeriod:      5 * time.Second,
		FollowDist:      2,
		UseArchiveNode:  false,
		MaxRetry:        5,
		SkipOnErr:       true,
		MaxMetadataSize: 50 * 1024,
		MetadataTimeout: 30 * time.Second,
	}

	// OpenSea market V2 contract has been deployed on the mainnet at
	// block num 14120913, so scraper starts from this block.
	defOpenSeaBAYCState = &OpenSeaBAYCScraperState{LastBlockNum: 14120913}

	whitelistedCollections = []string{
		"0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D",
		"0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB",
		"0x60E4d786628Fea6478F785A6d7e704777c86a7c6",
		"0x49cF6f5d44E70224e2E23fDcdd2C053F30aDA28B",
		"0x8a90CAb2b38dba80c64b7734e58Ee1dB38B8992e",
		"0x1A92f7381B9F03921564a437210bB9396471050C",
		"0xa3AEe8BcE55BEeA1951EF834b99f3Ac60d1ABeeB",
		"0xba30E5F9Bb24caa003E9f2f0497Ad287FDF95623",
		"0xe785E82358879F061BC3dcAC6f0444462D4b5330",
		"0x059EDD72Cd353dF5106D2B9cC5ab83a52287aC3a",
		"0x7Bd29408f11D2bFC23c34f18275bBf23bB716Bc7",
		"0x6dc6001535e15b9def7b0f6A20a2111dFA9454E2",
		"0xaaDc2D4261199ce24A4B0a57370c4FCf43BB60aa",
	}
	assetCacheOpenseaBayc = make(map[string]dia.Asset)
)

func init() {
	var err error

	openSeaABI, err = abi.JSON(strings.NewReader(opensea.ContractABI))
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
}

func NewOpenSeaBAYCScraper(rdb *models.RelDB, exchange dia.NFTExchange) *OpenSeaBAYCScraper {
	ctx := context.Background()

	eth, err := ethclient.Dial(utils.Getenv("ETH_URI_REST", ""))
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	s := &OpenSeaBAYCScraper{
		conf:     &OpenSeaBAYCScraperConfig{},
		state:    &OpenSeaBAYCScraperState{},
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
		log.Errorf("opensea scraper could not be initialized: %s", err.Error())
		return nil
	}

	go s.mainLoop()

	return s
}

// init scraper
// if there are no values stored previously, use defaults and store them
func (s *OpenSeaBAYCScraper) initScraper(ctx context.Context) error {
	if err := s.loadConfig(ctx); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Errorf("unable to read scraper config from rdb: %s", err.Error())
			return err
		}

		// use & store defaults if there is no record in the scraper table

		defConf := *defOpenSeaBAYCConf // copy
		s.conf = &defConf
		if err := s.tradeScraper.datastore.SetScraperConfig(ctx, OpenSeaBAYC, s.conf); err != nil {
			log.Errorf("unable to store scraper config on rdb: %s", err.Error())
			return err
		}

		defState := *defOpenSeaBAYCState // copy
		s.state = &defState
		if err := s.tradeScraper.datastore.SetScraperState(ctx, OpenSeaBAYC, s.state); err != nil {
			log.Errorf("unable to store scraper state on rdb: %s", err.Error())
			return err
		}

		return nil
	}

	return s.loadState(ctx)
}

func (s *OpenSeaBAYCScraper) loadConfig(ctx context.Context) error {
	return s.tradeScraper.datastore.GetScraperConfig(ctx, OpenSeaBAYC, s.conf)
}

func (s *OpenSeaBAYCScraper) loadState(ctx context.Context) error {
	return s.tradeScraper.datastore.GetScraperState(ctx, OpenSeaBAYC, s.state)
}

func (s *OpenSeaBAYCScraper) storeState(ctx context.Context) error {
	return s.tradeScraper.datastore.SetScraperState(ctx, OpenSeaBAYC, s.state)
}

func (s *OpenSeaBAYCScraper) mainLoop() {
	defer func() {
		s.tradeScraper.closed = true

		close(s.tradeScraper.chanTrade)
		close(s.tradeScraper.shutdownDone)
	}()

	log.Infof("opensea scraper has been started (batch: %d, period: %s)", s.conf.BatchSize, s.conf.WaitPeriod.String())

	for stop := false; !stop; {
		if err := s.FetchTrades(); err != nil {
			if errors.Is(err, errOpenSeaShutdownRequest) {
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
func (s *OpenSeaBAYCScraper) FetchTrades() error {
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

	log.Infof("fetching opensea trade transactions from block %d(+%d)", s.state.LastBlockNum, s.conf.BatchSize)

	// fetch trade transactions
	res, err := utils.EthFilterTXs(ctx, s.tradeScraper.ethConnection, utils.EthTxFilterCriteria{
		StartBlockNum:      s.state.LastBlockNum,
		StartTxIndex:       s.state.LastTxIndex,
		LimitBlocks:        s.conf.BatchSize,
		BehindHighestBlock: s.conf.FollowDist,
		EvAddrs:            []common.Address{common.HexToAddress(s.conf.ContractAddr)},
		Events:             []common.Hash{openSeaABI.Events["OrdersMatched"].ID},
	})

	if err != nil {
		log.Warnf("unable to filter opensea trades: %s", err.Error())
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
				if errState := s.storeState(ctx); err != nil {
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

func (s *OpenSeaBAYCScraper) processTx(ctx context.Context, tx *utils.EthFilteredTx) (bool, error) {
	log.Tracef("process tx -> block: %d, tx index: %d, tx hash: %s", tx.BlockNum, tx.TXIndex, tx.TXHash.Hex())

	// skip if the transaction has multiple OrdersMatched logs
	if len(tx.Logs) != 1 {
		return true, nil
	}

	marketContract, err := opensea.NewContract(tx.Logs[0].Address, s.tradeScraper.ethConnection)
	if err != nil {
		log.Errorf("unable to make new market contract for address: %s", tx.Logs[0].Address.Hex())
		return false, err
	}

	ev, err := marketContract.ParseOrdersMatched(tx.Logs[0])
	if err != nil {
		log.Errorf("unable to decode opensea OrdersMatched event(tx: %s, logIndex: %d) (SKIPPED!): %s", tx.TXHash, tx.Logs[0].Index, err.Error())
		return true, nil // skip
	}

	txData, pending, err := s.tradeScraper.ethConnection.TransactionByHash(ctx, tx.TXHash)
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

	currSymbol := "ETH"
	currAddr := common.Address{}
	currDecimals := 18

	// if an ERC20 token used for the trade
	if new(big.Int).Cmp(txData.Value()) == 0 {
		var tokenTransfers []*erc20Transfer
		tokenTransfers, err = s.findERC20Transfers(ctx, receipt, ev.Price)
		if err != nil {
			log.Errorf("unable to find erc20 transfers for transaction(%s): %s", tx.TXHash, err.Error())
			return false, err
		}

		if len(tokenTransfers) == 1 {
			currAddr = tokenTransfers[0].TokenAddr
			currDecimals = tokenTransfers[0].Decimals

			if v := tokenTransfers[0].TokenSymbol; v != nil {
				currSymbol = *v
			}
		}
	}

	transfers, err := s.findERC721Transfers(ctx, receipt)
	if err != nil {
		log.Errorf("unable to find transfers of the event(block: %d, tx index: %d, tx: %s): %s", ev.Raw.BlockNumber, ev.Raw.TxIndex, ev.Raw.TxHash.Hex(), err.Error())
		return false, err
	}

	// skip if the event has no transfer
	if len(transfers) == 0 {
		log.Tracef("event(block: %d, tx index: %d, tx: %s) skipped due to it has no erc721 transfer log", ev.Raw.BlockNumber, ev.Raw.TxIndex, ev.Raw.TxHash.Hex())
		return true, nil
	}

	// skip if the event has multiple transfers due to we can't calculate the price of trade
	if len(transfers) > 1 {
		log.Tracef("event(block: %d, tx index: %d, tx: %s) skipped due to it has multiple erc721 transfer logs", ev.Raw.BlockNumber, ev.Raw.TxIndex, ev.Raw.TxHash.Hex())
		return true, nil
	}

	if !utils.Contains(&whitelistedCollections, transfers[0].NFTAddress.Hex()) {
		log.Warnf("skip non-whitelisted collection")
		return true, nil
	}

	normPrice := decimal.NewFromBigInt(ev.Price, 0).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(currDecimals))))

	usdPrice, err := s.calcUSDPrice(ev.Raw.BlockNumber, currAddr, currSymbol, normPrice)
	if err != nil {
		log.Errorf("unable to calculate usd price of the event(block: %d, log: %d, tx: %s): %s", ev.Raw.BlockNumber, ev.Raw.TxIndex, ev.Raw.TxHash.Hex(), err.Error())
		return false, err
	}

	if err := s.notifyTrade(ev, transfers[0], ev.Price, normPrice, usdPrice, currSymbol, currAddr); err != nil {
		if !errors.Is(err, errOpenSeaShutdownRequest) {
			log.Warnf("event(block: %d, tx index: %d, tx: %s) couldn't processed: %s", ev.Raw.BlockNumber, ev.Raw.TxIndex, ev.Raw.TxHash.Hex(), err.Error())
		}

		return false, err
	}

	return false, nil
}

func (s *OpenSeaBAYCScraper) notifyTrade(ev *opensea.ContractOrdersMatched, transfer *erc721Transfer, price *big.Int, priceDec decimal.Decimal, usdPrice float64, currSymbol string, currAddr common.Address) error {
	nftClass, err := s.createOrReadNFTClass(transfer)
	if err != nil {
		return err
	}

	nft, err := s.createOrReadNFT(nftClass, transfer)
	if err != nil {
		return err
	}

	// Get block time.
	timestamp, err := ethhelper.GetBlockTimeEth(int64(ev.Raw.BlockNumber), s.tradeScraper.datastore, s.tradeScraper.ethConnection)
	if err != nil {
		log.Errorf("getting block time: %+v", err)
	}

	trade := dia.NFTTrade{
		NFT:         *nft,
		Price:       price,
		PriceUSD:    usdPrice,
		FromAddress: transfer.From.Hex(),
		ToAddress:   transfer.To.Hex(),
		BlockNumber: ev.Raw.BlockNumber,
		Timestamp:   timestamp,
		TxHash:      ev.Raw.TxHash.Hex(),
		Exchange:    s.exchange.Name,
	}

	if asset, ok := assetCacheOpenseaBayc[dia.ETHEREUM+"-"+currAddr.Hex()]; ok {
		trade.Currency = asset
	} else {
		currency, err := s.tradeScraper.datastore.GetAsset(currAddr.Hex(), dia.ETHEREUM)
		if err != nil {
			log.Errorf("cannot fetch asset %s -- %s", dia.ETHEREUM, currAddr.Hex())
		}
		trade.Currency = currency
		assetCacheOpenseaBayc[dia.ETHEREUM+"-"+currAddr.Hex()] = currency
	}

	fmt.Println("found trade: ", trade)

	// handle close request if the chanTrade not consumed immediately
	select {
	case s.tradeScraper.chanTrade <- trade:
	case <-s.tradeScraper.shutdown:
		return errOpenSeaShutdownRequest
	}

	return nil
}

func (s *OpenSeaBAYCScraper) createOrReadNFTClass(transfer *erc721Transfer) (*dia.NFTClass, error) {
	nftClass, err := s.tradeScraper.datastore.GetNFTClass(transfer.NFTAddress.Hex(), dia.ETHEREUM)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Warnf("unable to read nftclass from reldb: %s", err.Error())
			return nil, err
		}

		nftClass = dia.NFTClass{
			Address:      transfer.NFTAddress.Hex(),
			Blockchain:   dia.ETHEREUM,
			ContractType: openSeaNFTContractType,
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

func (s *OpenSeaBAYCScraper) createOrReadNFT(nftClass *dia.NFTClass, transfer *erc721Transfer) (*dia.NFT, error) {
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

func (s *OpenSeaBAYCScraper) calcUSDPrice(blockNum uint64, tokenAddr common.Address, symbol string, price decimal.Decimal) (float64, error) {
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

func (s *OpenSeaBAYCScraper) findPrice(blockNum uint64, tokenAddr common.Address, symbol string) (decimal.Decimal, error) {
	// TODO: find the token price in usd for the given block number
	switch symbol {
	case "ETH", "WETH":
		return decimal.NewFromString("2040.0910")

	case "MANA":
		return decimal.NewFromString("0.5")

	default:
		return decimal.NewFromString("1")
	}
}

// GetDataChannel returns the scrapers data channel.
func (s *OpenSeaBAYCScraper) GetTradeChannel() chan dia.NFTTrade {
	return s.tradeScraper.chanTrade
}

func (s *OpenSeaBAYCScraper) Close() error {
	if s.tradeScraper.closed {
		return errors.New("scraper already closed")
	}

	close(s.tradeScraper.shutdown)

	return nil
}

// it searches the creation transaction for the given contract address using binary search in complexity of o(log n)
func (s *OpenSeaBAYCScraper) findContractCreationInfo(ctx context.Context, contractAddr common.Address) (createdBy common.Address, createdAt time.Time, err error) {
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

func (s *OpenSeaBAYCScraper) findERC20Transfers(ctx context.Context, receipt *types.Receipt, filterAmount *big.Int) ([]*erc20Transfer, error) {
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

		if filterAmount != nil && filterAmount.Cmp(ev.Value) != 0 {
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

		metadata, err := erc20.NewERC20Metadata(txLog.Address, s.tradeScraper.ethConnection)
		if err != nil {
			log.Warnf("unable to bind erc20 metadata contract at address %s: %s", txLog.Address.Hex(), err.Error())
			continue
		}

		callOpts := &bind.CallOpts{Context: ctx}

		if s.conf.UseArchiveNode {
			callOpts.BlockNumber = new(big.Int).SetUint64(txLog.BlockNumber)
		}

		symbol, err := metadata.Symbol(callOpts)
		if err != nil {
			log.Warnf("unable to read token symbol from metadata interface of erc20(addr: %s): %s", txLog.Address.Hex(), err.Error())
			continue
		}

		transfer.TokenSymbol = &symbol

		decimals, err := metadata.Decimals(callOpts)
		if err != nil {
			log.Warnf("unable to read token decimals from metadata interface of erc20(addr: %s): %s", txLog.Address.Hex(), err.Error())
			continue
		}

		transfer.Decimals = int(decimals)
	}

	return transfers, nil
}

// it finds the transfer events of ERC721 in the given transaction
func (s *OpenSeaBAYCScraper) findERC721Transfers(ctx context.Context, receipt *types.Receipt) ([]*erc721Transfer, error) {
	transfers := make([]*erc721Transfer, 0, 1)

	for _, txLog := range receipt.Logs {
		if len(txLog.Topics) < 1 || txLog.Topics[0] != erc721ABI.Events["Transfer"].ID {
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

		transfer := &erc721Transfer{
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

func (s *OpenSeaBAYCScraper) readNFTAttr(ctx context.Context, uri string) (map[string]interface{}, error) {
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
