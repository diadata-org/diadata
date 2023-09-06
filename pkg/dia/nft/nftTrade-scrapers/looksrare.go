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
	"github.com/diadata-org/diadata/config/nftContracts/looksrare"
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
	// we assume all of the NFTs traded on LooksRare are ERC721(1155 is an extension of it)
	looksRareNFTContractType = "ERC721"

	LooksRare = "LooksRare"
)

type LooksRareScraperConfig struct {
	// looksrare's exchange contract address on connected blockchain network
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

type LooksRareScraperState struct {
	// last block number has been processed
	LastBlockNum uint64 `json:"last_block_num"`

	// last transaction index in the block(curr) has been processed
	LastTxIndex uint `json:"last_tx_index"`

	// holds the latest error message that occurred while scraping
	LastErr string `json:"last_error"`

	// indicates the number of consecutive error, reset on any successful operation
	ErrCounter int `json:"count_of_error"`
}

type LooksRareScraper struct {
	tradeScraper TradeScraper

	mu       sync.Mutex
	conf     *LooksRareScraperConfig
	state    *LooksRareScraperState
	exchange dia.NFTExchange
}

type erc721Metadata struct {
	NFTAddress common.Address
	Name       *string
	Symbol     *string
	TokenID    *big.Int
	TokenURI   *string
	TokenAttrs map[string]interface{}
}

type looksRareTakerBidAskEvent struct {
	OrderHash  [32]byte
	OrderNonce *big.Int
	From       common.Address
	To         common.Address
	Strategy   common.Address
	Currency   common.Address
	Collection common.Address
	TokenId    *big.Int
	Amount     *big.Int
	Price      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

var (
	errLooksRareShutdownRequest = errors.New("shutdown requested")

	// default values are valid for the first run which is it saves
	// these configs to the DB
	defLooksRareConf = &LooksRareScraperConfig{
		ContractAddr:    "0x59728544B08AB483533076417FbBB2fD0B17CE3a",
		BatchSize:       5000,
		WaitPeriod:      60 * time.Second,
		FollowDist:      10,
		UseArchiveNode:  false,
		MaxRetry:        5,
		SkipOnErr:       true,
		MaxMetadataSize: 50 * 1024,
		MetadataTimeout: 30 * time.Second,
	}

	// LooksRare market contract has been deployed on the mainnet at
	// block num 13885625, so scraper starts from this block
	defLooksRareState = &LooksRareScraperState{LastBlockNum: 13885625}

	looksRareABI abi.ABI
	//looksRareErc20ABI  abi.ABI
	//looksRareErc721ABI abi.ABI

	assetCacheLooksrare = make(map[string]dia.Asset)
)

func init() {
	var err error

	looksRareABI, err = abi.JSON(strings.NewReader(looksrare.ContractABI))
	if err != nil {
		panic(err)
	}

	_, err = abi.JSON(strings.NewReader(erc20.ERC20ABI))
	if err != nil {
		panic(err)
	}

	_, err = abi.JSON(strings.NewReader(erc721.ERC721ABI))
	if err != nil {
		panic(err)
	}
}

func NewLooksRareScraper(rdb *models.RelDB, exchange dia.NFTExchange) *LooksRareScraper {
	ctx := context.Background()

	eth, err := ethclient.Dial(utils.Getenv("ETH_URI_REST", ""))
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	s := &LooksRareScraper{
		conf:     &LooksRareScraperConfig{},
		state:    &LooksRareScraperState{},
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
		log.Errorf("looksrare scraper could not be initialized: %s", err.Error())
		return nil
	}

	go s.mainLoop()

	return s
}

// init scraper
// if there are no values stored previously, use defaults and store them
func (s *LooksRareScraper) initScraper(ctx context.Context) error {
	if err := s.loadConfig(ctx); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Errorf("unable to read scraper config from rdb: %s", err.Error())
			return err
		}

		// use & store defaults if there is no record in the scraper table

		defConf := *defLooksRareConf // copy
		s.conf = &defConf
		if err := s.tradeScraper.datastore.SetScraperConfig(ctx, LooksRare, s.conf); err != nil {
			log.Errorf("unable to store scraper config on rdb: %s", err.Error())
			return err
		}

		defState := *defLooksRareState // copy
		s.state = &defState
		if err := s.tradeScraper.datastore.SetScraperState(ctx, LooksRare, s.state); err != nil {
			log.Errorf("unable to store scraper state on rdb: %s", err.Error())
			return err
		}

		return nil
	}

	return s.loadState(ctx)
}

func (s *LooksRareScraper) loadConfig(ctx context.Context) error {
	return s.tradeScraper.datastore.GetScraperConfig(ctx, LooksRare, s.conf)
}

func (s *LooksRareScraper) loadState(ctx context.Context) error {
	return s.tradeScraper.datastore.GetScraperState(ctx, LooksRare, s.state)
}

func (s *LooksRareScraper) storeState(ctx context.Context) error {
	return s.tradeScraper.datastore.SetScraperState(ctx, LooksRare, s.state)
}

func (s *LooksRareScraper) mainLoop() {
	defer func() {
		s.tradeScraper.closed = true

		close(s.tradeScraper.chanTrade)
		close(s.tradeScraper.shutdownDone)
	}()

	log.Infof("looksrare scraper has been started (batch: %d, period: %s)", s.conf.BatchSize, s.conf.WaitPeriod.String())

	for stop := false; !stop; {
		if err := s.FetchTrades(); err != nil {
			if errors.Is(err, errLooksRareShutdownRequest) {
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
func (s *LooksRareScraper) FetchTrades() error {
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

	log.Infof("fetching looksrare trade transactions from block %d(+%d)", s.state.LastBlockNum, s.conf.BatchSize)

	// fetch trade transactions
	res, err := utils.EthFilterTXs(ctx, s.tradeScraper.ethConnection, utils.EthTxFilterCriteria{
		StartBlockNum:      s.state.LastBlockNum,
		StartTxIndex:       s.state.LastTxIndex,
		LimitBlocks:        s.conf.BatchSize,
		BehindHighestBlock: s.conf.FollowDist,
		EvAddrs:            []common.Address{common.HexToAddress(s.conf.ContractAddr), common.HexToAddress(s.conf.ContractAddr)},
		Events:             []common.Hash{looksRareABI.Events["TakerAsk"].ID, looksRareABI.Events["TakerBid"].ID},
	})

	if err != nil {
		log.Warnf("unable to filter looksrare trades: %s", err.Error())
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
				if err = s.storeState(ctx); err != nil {
					log.Warnf("unable to store scraper state: %s", err.Error())
					return err
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

func (s *LooksRareScraper) processTx(ctx context.Context, tx *utils.EthFilteredTx) (bool, error) {
	log.Tracef("process tx -> block: %d, tx index: %d, tx hash: %s", tx.BlockNum, tx.TXIndex, tx.TXHash.Hex())

	// skip if the transaction has multiple TakerBid/TakerAsk logs
	if len(tx.Logs) != 1 {
		return true, nil
	}

	marketContract, err := looksrare.NewContract(tx.Logs[0].Address, s.tradeScraper.ethConnection)
	if err != nil {
		log.Errorf("unable to make new market contract for address: %s", tx.Logs[0].Address.Hex())
		return false, err
	}

	ev := &looksRareTakerBidAskEvent{}
	if looksRareABI.Events["TakerAsk"].ID == tx.Logs[0].Topics[0] {

		evTakerAsk, errParseTakerAsk := marketContract.ParseTakerAsk(tx.Logs[0])
		if errParseTakerAsk != nil {
			log.Errorf("unable to decode looksrare TakerAsk event(tx: %s, logIndex: %d) (SKIPPED!): %s", tx.TXHash, tx.Logs[0].Index, errParseTakerAsk.Error())
			return true, nil // skip
		}
		ev.OrderHash = evTakerAsk.OrderHash
		ev.OrderNonce = evTakerAsk.OrderNonce
		ev.From = evTakerAsk.Taker
		ev.To = evTakerAsk.Maker
		ev.Strategy = evTakerAsk.Strategy
		ev.Currency = evTakerAsk.Currency
		ev.Collection = evTakerAsk.Collection
		ev.TokenId = evTakerAsk.TokenId
		ev.Amount = evTakerAsk.Amount
		ev.Price = evTakerAsk.Price
		ev.Raw = evTakerAsk.Raw

	}
	if looksRareABI.Events["TakerBid"].ID == tx.Logs[0].Topics[0] {

		evTakerBid, errParseTakerBid := marketContract.ParseTakerBid(tx.Logs[0])
		if errParseTakerBid != nil {
			log.Errorf("unable to decode looksrare TakerBid event(tx: %s, logIndex: %d) (SKIPPED!): %s", tx.TXHash, tx.Logs[0].Index, errParseTakerBid.Error())
			return true, nil // skip
		}
		ev.OrderHash = evTakerBid.OrderHash
		ev.OrderNonce = evTakerBid.OrderNonce
		ev.To = evTakerBid.Taker
		ev.From = evTakerBid.Maker
		ev.Strategy = evTakerBid.Strategy
		ev.Currency = evTakerBid.Currency
		ev.Collection = evTakerBid.Collection
		ev.TokenId = evTakerBid.TokenId
		ev.Amount = evTakerBid.Amount
		ev.Price = evTakerBid.Price
		ev.Raw = evTakerBid.Raw

	}

	currSymbol := "ETH"
	currAddr := common.Address{}
	currDecimals := 18

	if ev.Currency.Hex() != "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" { // WETH
		symbol, decimals, errERC20Metadata := s.getERC20Metadata(ctx, ev.Currency, ev.Raw.BlockNumber)
		if errERC20Metadata != nil {
			log.Errorf("unable to find erc20 metadata for address (%s) in transaction(%s): %s", ev.Currency, tx.TXHash, errERC20Metadata.Error())
		} else {
			currAddr = ev.Currency
			currSymbol = *symbol
			currDecimals = decimals
		}
	}

	erc721Meta, errMetadata := s.getERC721Metadata(ctx, ev.Collection, ev.TokenId, ev.Raw.BlockNumber)
	if errMetadata != nil {
		log.Errorf("unable to find transfers of the event(block: %d, tx index: %d, tx: %s): %s", ev.Raw.BlockNumber, ev.Raw.TxIndex, ev.Raw.TxHash.Hex(), errMetadata.Error())
		return false, errMetadata
	}

	normPrice := decimal.NewFromBigInt(ev.Price, 0).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(currDecimals))))

	usdPrice, errUSDPrice := s.calcUSDPrice(ev.Raw.BlockNumber, currAddr, currSymbol, normPrice)
	if errUSDPrice != nil {
		log.Errorf("unable to calculate usd price of the event(block: %d, log: %d, tx: %s): %s", ev.Raw.BlockNumber, ev.Raw.TxIndex, ev.Raw.TxHash.Hex(), errUSDPrice.Error())
		return false, errUSDPrice
	}

	if errTrade := s.notifyTrade(ev, erc721Meta, ev.Price, normPrice, usdPrice, currSymbol, currAddr); errTrade != nil {
		if !errors.Is(errTrade, errLooksRareShutdownRequest) {
			log.Warnf("event(block: %d, tx index: %d, tx: %s) couldn't processed: %s", ev.Raw.BlockNumber, ev.Raw.TxIndex, ev.Raw.TxHash.Hex(), errTrade.Error())
		}

		return false, errTrade
	}

	return false, nil
}

func (s *LooksRareScraper) getERC20Metadata(ctx context.Context, address common.Address, blockNumber uint64) (*string, int, error) {
	metadata, err := erc20.NewERC20Metadata(address, s.tradeScraper.ethConnection)
	if err != nil {
		return nil, 0, err
	}

	callOpts := &bind.CallOpts{Context: ctx}

	if s.conf.UseArchiveNode {
		callOpts.BlockNumber = new(big.Int).SetUint64(blockNumber)
	}

	symbol, err := metadata.Symbol(callOpts)
	if err != nil {
		log.Warnf("unable to read token symbol from metadata interface of erc20(addr: %s): %s", address.Hex(), err.Error())
		return nil, 0, err
	}

	decimals, err := metadata.Decimals(callOpts)
	if err != nil {
		log.Warnf("unable to read token decimals from metadata interface of erc20(addr: %s): %s", address.Hex(), err.Error())
		return nil, 0, err
	}

	return &symbol, int(decimals), nil
}

func (s *LooksRareScraper) notifyTrade(ev *looksRareTakerBidAskEvent, erc721 *erc721Metadata, price *big.Int, priceDec decimal.Decimal, usdPrice float64, currSymbol string, currAddr common.Address) error {

	nftClass, err := s.createOrReadNFTClass(erc721)
	if err != nil {
		return err
	}

	nft, err := s.createOrReadNFT(nftClass, erc721)
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
		FromAddress: ev.From.Hex(),
		ToAddress:   ev.To.Hex(),
		BlockNumber: ev.Raw.BlockNumber,
		Timestamp:   timestamp,
		TxHash:      ev.Raw.TxHash.Hex(),
		Exchange:    s.exchange.Name,
	}

	if asset, ok := assetCacheLooksrare[dia.ETHEREUM+"-"+currAddr.Hex()]; ok {
		trade.Currency = asset
	} else {
		currency, err := s.tradeScraper.datastore.GetAsset(currAddr.Hex(), dia.ETHEREUM)
		if err != nil {
			log.Errorf("cannot fetch asset %s -- %s", dia.ETHEREUM, currAddr.Hex())
		}
		trade.Currency = currency
		assetCacheLooksrare[dia.ETHEREUM+"-"+currAddr.Hex()] = currency
	}

	fmt.Println("found trade: ", trade)

	// handle close request if the chanTrade not consumed immediately
	select {
	case s.tradeScraper.chanTrade <- trade:
	case <-s.tradeScraper.shutdown:
		return errLooksRareShutdownRequest
	}

	return nil
}

func (s *LooksRareScraper) createOrReadNFTClass(erc721 *erc721Metadata) (*dia.NFTClass, error) {
	nftClass, err := s.tradeScraper.datastore.GetNFTClass(erc721.NFTAddress.Hex(), dia.ETHEREUM)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Warnf("unable to read nftclass from reldb: %s", err.Error())
			return nil, err
		}

		nftClass = dia.NFTClass{
			Address:      erc721.NFTAddress.Hex(),
			Blockchain:   dia.ETHEREUM,
			ContractType: looksRareNFTContractType,
		}

		if erc721.Name != nil {
			nftClass.Name = *erc721.Name
		}

		if erc721.Symbol != nil {
			nftClass.Symbol = *erc721.Symbol
		}

		if err = s.tradeScraper.datastore.SetNFTClass(nftClass); err != nil {
			log.Warnf("unable to create nftclass on reldb: %s", err.Error())
			return nil, err
		}
	}

	return &nftClass, nil
}

func (s *LooksRareScraper) createOrReadNFT(nftClass *dia.NFTClass, erc721 *erc721Metadata) (*dia.NFT, error) {
	nft, err := s.tradeScraper.datastore.GetNFT(nftClass.Address, dia.ETHEREUM, erc721.TokenID.String())
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
			TokenID:        erc721.TokenID.String(),
			CreationTime:   createdAt,
			CreatorAddress: createdBy.Hex(),
			Attributes:     erc721.TokenAttrs,
		}

		if erc721.TokenURI != nil {
			nft.URI = *erc721.TokenURI
		}

		if err = s.tradeScraper.datastore.SetNFT(nft); err != nil {
			log.Warnf("unable to create nft on reldb: %s", err.Error())
			return nil, err
		}
	}

	return &nft, nil
}

func (s *LooksRareScraper) calcUSDPrice(blockNum uint64, tokenAddr common.Address, symbol string, price decimal.Decimal) (float64, error) {
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

func (s *LooksRareScraper) findPrice(blockNum uint64, tokenAddr common.Address, symbol string) (decimal.Decimal, error) {
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
func (s *LooksRareScraper) GetTradeChannel() chan dia.NFTTrade {
	return s.tradeScraper.chanTrade
}

func (s *LooksRareScraper) Close() error {
	if s.tradeScraper.closed {
		return errors.New("scraper already closed")
	}

	close(s.tradeScraper.shutdown)

	return nil
}

// it searches the creation transaction for the given contract address using binary search in complexity of o(log n)
func (s *LooksRareScraper) findContractCreationInfo(ctx context.Context, contractAddr common.Address) (createdBy common.Address, createdAt time.Time, err error) {
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

// it gets the ERC721 metadata
func (s *LooksRareScraper) getERC721Metadata(ctx context.Context, nftAddress common.Address, tokenID *big.Int, blockNumber uint64) (*erc721Metadata, error) {

	callOpts := &bind.CallOpts{Context: ctx}

	if s.conf.UseArchiveNode {
		callOpts.BlockNumber = new(big.Int).SetUint64(blockNumber)
	}

	if md, err := erc721.NewERC721Metadata(nftAddress, s.tradeScraper.ethConnection); err != nil {
		log.Warnf("unable to bind erc721 metadata contract at address %s: %s", nftAddress.Hex(), err.Error())
		return nil, err
	} else {
		erc721 := &erc721Metadata{}
		erc721.NFTAddress = nftAddress
		erc721.TokenID = tokenID
		if nftName, err := md.Name(callOpts); err != nil {
			log.Warnf("unable to read nft name from metadata interface of erc721(addr: %s): %s", nftAddress.Hex(), err.Error())
		} else {
			erc721.Name = &nftName
		}

		if nftSymbol, err := md.Symbol(callOpts); err != nil {
			log.Warnf("unable to read nft symbol from metadata interface of nft(addr: %s): %s", nftAddress.Hex(), err.Error())
		} else {
			erc721.Symbol = &nftSymbol
		}

		if tokenURI, err := md.TokenURI(callOpts, tokenID); err != nil {
			log.Warnf("unable to find token(%s) uri: %s", tokenID.String(), err.Error())
		} else if attrs, err := s.readNFTAttr(ctx, tokenURI); err != nil {
			log.Warnf("unable to read token(%s) attributes: %s", tokenID.String(), err.Error())
		} else {
			erc721.TokenURI = &tokenURI
			erc721.TokenAttrs = attrs
		}
		return erc721, nil
	}

}

func (s *LooksRareScraper) readNFTAttr(ctx context.Context, uri string) (map[string]interface{}, error) {
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
