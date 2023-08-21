package scrapers

import (
	"context"
	"errors"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	pairfactorycontract "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/maverick/pairfactory"

	poolcontract "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/maverick/pool"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var maverickPoolMap = make(map[string]MaverickPair)

const (
	factoryContractAddressEth                = "0xEb6625D65a0553c9dBc64449e56abFe519bd9c9B"
	factoryContractAddressDeploymentBlockEth = uint64(17210221)
	maverickWaitMilliseconds                 = "25"
)

type MaverickScraper struct {
	WsClient   *ethclient.Client
	RestClient *ethclient.Client
	relDB      *models.RelDB
	// signaling channels for session initialization and finishing
	initDone     chan nothing
	run          bool
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers                     map[string]*MaverickPairScraper
	exchangeName                     string
	blockchain                       string
	poolFactoryContractAddress       string
	poolFactoryContractCreationBlock uint64
	chanTrades                       chan *dia.Trade
	waitTime                         int
	// If true, only pairs given in config file are scraped. Default is false.
	listenByAddress  bool
	fetchPoolsFromDB bool
}

type Token struct {
	Address  common.Address
	Symbol   string
	Decimals uint8
	Name     string
}

type MaverickPair struct {
	Token0      Token
	Token1      Token
	ForeignName string
	Address     common.Address
}

type MaverickPool struct {
	pair        MaverickPair
	fee         float64
	tickSpacing int64
	lookback    *big.Int
}

// NewMaverickScraper returns a new MaverickScraper for the given pair
func NewMaverickScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *MaverickScraper {
	log.Info("NewMaverickScraper: ", exchange.Name)
	var (
		s                *MaverickScraper
		listenByAddress  bool
		fetchPoolsFromDB bool
		err              error
	)

	listenByAddress, err = strconv.ParseBool(utils.Getenv("LISTEN_BY_ADDRESS", "false"))
	if err != nil {
		log.Fatal("parse LISTEN_BY_ADDRESS: ", err)
	}

	fetchPoolsFromDB, err = strconv.ParseBool(utils.Getenv("FETCH_POOLS_FROM_DB", "false"))
	if err != nil {
		log.Fatal("parse FETCH_POOLS_FROM_DB: ", err)
	}

	switch exchange.Name {
	case dia.MaverickExchange:
		s = makeMaverickScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialEth, wsDialEth, maverickWaitMilliseconds, factoryContractAddressEth, factoryContractAddressDeploymentBlockEth)
		//case dia.MaverickExchangeBNB:
		//	s = makeMaverickScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialEth, wsDialEth, maverickWaitMilliseconds)
		//case dia.MaverickExchangeZKSync:
		//	s = makeMaverickScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialEth, wsDialEth, maverickWaitMilliseconds)
	}

	s.relDB = relDB

	// Only include pools with (minimum) liquidity bigger than given env var.
	liquidityThreshold, err := strconv.ParseFloat(utils.Getenv("LIQUIDITY_THRESHOLD", "0"), 64)
	if err != nil {
		liquidityThreshold = float64(0)
		log.Warnf("parse liquidity threshold:  %v. Set to default %v", err, liquidityThreshold)
	}
	// Only include pools with (minimum) liquidity USD value bigger than given env var.
	liquidityThresholdUSD, err := strconv.ParseFloat(utils.Getenv("LIQUIDITY_THRESHOLD_USD", "0"), 64)
	if err != nil {
		liquidityThresholdUSD = float64(0)
		log.Warnf("parse liquidity threshold:  %v. Set to default %v", err, liquidityThresholdUSD)
	}

	// Fetch all pool with given liquidity threshold from database.
	maverickPoolMap, err = s.makePoolMap(liquidityThreshold, liquidityThresholdUSD)
	if err != nil {
		log.Fatal("build poolMap: ", err)
	}

	if scrape {
		go s.mainLoop()
	}
	return s
}

// makeMaverickScraper returns a maverick scraper as used in NewUniswapScraper.
func makeMaverickScraper(exchange dia.Exchange, listenByAddress bool, fetchPoolsFromDB bool, restDial string, wsDial string, waitMilliseconds string, factoryContractAddess string, factoryContractDeploymentBlock uint64) *MaverickScraper {
	var (
		restClient, wsClient *ethclient.Client
		err                  error
		s                    *MaverickScraper
		waitTime             int
	)

	log.Infof("Init rest and ws client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", restDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}
	wsClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", wsDial))
	if err != nil {
		log.Fatal("init ws client: ", err)
	}

	waitTime, err = strconv.Atoi(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_WAIT_TIME", waitMilliseconds))
	if err != nil {
		log.Error("could not parse wait time: ", err)
		waitTime = 5000
	}

	s = &MaverickScraper{
		WsClient:                         wsClient,
		RestClient:                       restClient,
		shutdown:                         make(chan nothing),
		shutdownDone:                     make(chan nothing),
		pairScrapers:                     make(map[string]*MaverickPairScraper),
		exchangeName:                     exchange.Name,
		error:                            nil,
		chanTrades:                       make(chan *dia.Trade),
		waitTime:                         waitTime,
		listenByAddress:                  listenByAddress,
		fetchPoolsFromDB:                 fetchPoolsFromDB,
		blockchain:                       exchange.BlockChain.Name,
		poolFactoryContractAddress:       factoryContractAddess,
		poolFactoryContractCreationBlock: factoryContractDeploymentBlock,
	}
	return s
}

func (s *MaverickScraper) mainLoop() {

	// Import tokens which appear as base token and we need a quotation for
	var err error
	reverseBasetokens, err = getReverseTokensFromConfig("maverick/reverse_tokens/" + s.exchangeName + "Basetoken")
	if err != nil {
		log.Error("error getting tokens for which pairs should be reversed: ", err)
	}
	log.Info("reverse basetokens: ", reverseBasetokens)
	reverseQuotetokens, err = getReverseTokensFromConfig("maverick/reverse_tokens/" + s.exchangeName + "Quotetoken")
	if err != nil {
		log.Error("error getting tokens for which pairs should be reversed: ", err)
	}
	log.Info("reverse quotetokens: ", reverseQuotetokens)

	// wait for all pairs have added into s.PairScrapers
	time.Sleep(4 * time.Second)
	s.run = true

	if s.listenByAddress || s.fetchPoolsFromDB {
		var wg sync.WaitGroup
		for address := range maverickPoolMap {
			time.Sleep(time.Duration(s.waitTime) * time.Millisecond)
			wg.Add(1)
			go func(address common.Address, w *sync.WaitGroup) {
				defer w.Done()
				s.ListenToPair(address)
			}(common.HexToAddress(address), &wg)
		}

	} else {
		addresses, err := s.getAllPairsAddress()
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Found ", len(addresses), " pairs")
		log.Info("Found ", len(s.pairScrapers), " pairScrapers")

		if len(s.pairScrapers) == 0 {
			s.error = errors.New("maverick: No pairs to scrap provided")
			log.Error(s.error.Error())
		}

		var wg sync.WaitGroup
		for _, address := range addresses {
			time.Sleep(time.Duration(s.waitTime) * time.Millisecond)
			wg.Add(1)
			go func(address common.Address, w *sync.WaitGroup) {
				defer w.Done()
				s.ListenToPair(address)
			}(address, &wg)
		}
		wg.Wait()

	}
}

// ListenToPair subscribes to a uniswap pool.
// If @byAddress is true, it listens by pool address, otherwise by index.
func (s *MaverickScraper) ListenToPair(address common.Address) {
	var (
		pair MaverickPair
		err  error
	)
	pair = maverickPoolMap[address.Hex()]
	if !s.listenByAddress && !s.fetchPoolsFromDB {
		//Get pool info from on-chain. @poolMap is empty.
		pair, err = s.getPairByAddress(address)
		if err != nil {
			log.Error("error fetching pair: ", err)
			return
		}
	} else {
		// Relevant pool info is retrieved from @poolMap.
		pair = maverickPoolMap[address.Hex()]
	}

	if len(pair.Token0.Symbol) < 2 || len(pair.Token1.Symbol) < 2 {
		log.Info("skip pair: ", pair.ForeignName)
		return
	}

	if helpers.AddressIsBlacklisted(pair.Token0.Address) || helpers.AddressIsBlacklisted(pair.Token1.Address) {
		log.Info("skip pair ", pair.ForeignName, ", address is blacklisted")
		return
	}
	if helpers.PoolIsBlacklisted(pair.Address) {
		log.Info("skip blacklisted pool ", pair.Address)
		return
	}

	log.Info("add pair scraper for: ", pair.ForeignName, " with address ", pair.Address.Hex())
	sink, err := s.GetSwapsChannel(pair.Address)
	if err != nil {
		log.Error("error fetching swaps channel: ", err)
	}

	go func() {
		for {
			rawSwap, ok := <-sink
			if ok {
				price, volume := getPriceAndVolumeFromRawSwapData(rawSwap, pair)
				token0 := dia.Asset{
					Address:    pair.Token0.Address.Hex(),
					Symbol:     pair.Token0.Symbol,
					Name:       pair.Token0.Name,
					Decimals:   pair.Token0.Decimals,
					Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
				}
				token1 := dia.Asset{
					Address:    pair.Token1.Address.Hex(),
					Symbol:     pair.Token1.Symbol,
					Name:       pair.Token1.Name,
					Decimals:   pair.Token1.Decimals,
					Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
				}
				t := &dia.Trade{
					Symbol:         pair.Token0.Symbol,
					Pair:           pair.ForeignName,
					Price:          price,
					Volume:         volume,
					BaseToken:      token1,
					QuoteToken:     token0,
					Time:           time.Unix(time.Now().Unix(), 0),
					PoolAddress:    rawSwap.Raw.Address.Hex(),
					ForeignTradeID: rawSwap.Raw.TxHash.Hex(),
					Source:         s.exchangeName,
					VerifiedPair:   true,
				}

				// TO DO: Refactor approach for reversing pairs.
				switch {
				case utils.Contains(reverseBasetokens, pair.Token1.Address.Hex()):
					// If we need quotation of a base token, reverse pair
					tSwapped, err := dia.SwapTrade(*t)
					if err == nil {
						t = &tSwapped
					}
				case utils.Contains(reverseQuotetokens, pair.Token0.Address.Hex()):
					// If we don't need quotation of quote token, reverse pair.
					tSwapped, err := dia.SwapTrade(*t)
					if err == nil {
						t = &tSwapped
					}
				case token0.Address == "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" && !utils.Contains(&mainBaseAssets, token1.Address):
					// Reverse almost all pairs WETH-XXX ...
					if s.exchangeName == dia.MaverickExchange {
						tSwapped, err := dia.SwapTrade(*t)
						if err == nil {
							t = &tSwapped
						}
					}
				// ...and USDT-XXX on Ethereum
				case token0.Address == mainBaseAssets[0] && token0.Blockchain == dia.ETHEREUM:
					tSwapped, err := dia.SwapTrade(*t)
					if err == nil {
						t = &tSwapped
					}
				}
				if price > 0 {
					log.Info("tx hash: ", rawSwap.Raw.TxHash.Hex())
					log.Infof("Got trade at time %v - symbol: %s, pair: %s, price: %v, volume:%v", t.Time, t.Symbol, t.Pair, t.Price, t.Volume)
					log.Infof("Base token info --- Symbol: %s - Address: %s - Blockchain: %s ", t.BaseToken.Symbol, t.BaseToken.Address, t.BaseToken.Blockchain)
					log.Info("----------------")
					s.chanTrades <- t
				}
			}
		}
	}()
}

func (s *MaverickScraper) getPairByAddress(pairAddress common.Address) (MaverickPair, error) {
	var (
		poolContractInstance *poolcontract.PoolCaller
		token0               dia.Asset
		token1               dia.Asset
	)

	connection := s.RestClient
	poolContractInstance, err := poolcontract.NewPoolCaller(pairAddress, connection)
	if err != nil {
		log.Error(err)
		return MaverickPair{}, err
	}

	// Getting tokens from pair
	address0, _ := poolContractInstance.TokenA(&bind.CallOpts{})
	address1, _ := poolContractInstance.TokenB(&bind.CallOpts{})

	// Only fetch assets from on-chain in case they are not in our DB.
	token0, err = s.relDB.GetAsset(address0.Hex(), s.blockchain)
	if err != nil {
		token0, err = ethhelper.ETHAddressToAsset(address0, s.RestClient, s.blockchain)
		if err != nil {
			return MaverickPair{}, err
		}
	}
	token1, err = s.relDB.GetAsset(address1.Hex(), s.blockchain)
	if err != nil {
		token1, err = ethhelper.ETHAddressToAsset(address1, s.RestClient, s.blockchain)
		if err != nil {
			return MaverickPair{}, err
		}
	}

	pair := MaverickPair{
		Token0: Token{
			Address:  common.HexToAddress(token0.Address),
			Symbol:   token0.Symbol,
			Decimals: token0.Decimals,
			Name:     token0.Name,
		},
		Token1: Token{
			Address:  common.HexToAddress(token1.Address),
			Symbol:   token1.Symbol,
			Decimals: token1.Decimals,
			Name:     token1.Name,
		},
		ForeignName: token0.Symbol + "-" + token1.Symbol,
		Address:     pairAddress,
	}

	return pair, nil
}

// getVolumeAndPriceFromRawSwapData returns price, volume and sell/buy information of @swap
func getPriceAndVolumeFromRawSwapData(swap *poolcontract.PoolSwap, pair MaverickPair) (price, volume float64) {
	decimals0 := int(pair.Token0.Decimals)
	decimals1 := int(pair.Token1.Decimals)

	if swap.TokenAIn {
		amount0In, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.AmountIn), new(big.Float).SetFloat64(math.Pow10(decimals0))).Float64()
		amount1Out, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.AmountOut), new(big.Float).SetFloat64(math.Pow10(decimals1))).Float64()
		volume = -amount0In
		price = amount1Out / amount0In
		return
	}

	amount1In, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.AmountIn), new(big.Float).SetFloat64(math.Pow10(decimals1))).Float64()
	amount0Out, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.AmountOut), new(big.Float).SetFloat64(math.Pow10(decimals0))).Float64()
	volume = amount0Out
	price = amount1In / amount0Out
	return
}

// makePoolMap returns a map with pool addresses as keys and the underlying MaverickPool as values.
// If s.listenByAddress is true, it only loads the corresponding assets from the list.
func (s *MaverickScraper) makePoolMap(liquiThreshold float64, liquidityThresholdUSD float64) (map[string]MaverickPair, error) {
	pm := make(map[string]MaverickPair)
	var (
		pools []dia.Pool
		err   error
	)

	if s.listenByAddress {
		// Only load pool info for addresses from json file.
		poolAddresses, errAddr := getAddressesFromConfig("maverick/subscribe_pools/" + s.exchangeName)
		if errAddr != nil {
			log.Error("fetch pool addresses from config file: ", errAddr)
		}
		for _, address := range poolAddresses {
			pool, errPool := s.relDB.GetPoolByAddress(Exchanges[s.exchangeName].BlockChain.Name, address.Hex())
			if errPool != nil {
				log.Fatalf("Get pool with address %s: %v", address.Hex(), errPool)
			}
			pools = append(pools, pool)
		}
	} else if s.fetchPoolsFromDB {
		// Load all pools above liqui threshold.
		pools, err = s.relDB.GetAllPoolsExchange(s.exchangeName, liquiThreshold)
		if err != nil {
			return pm, err
		}
	} else {
		// Pool info will be fetched from on-chain and poolMap is not needed.
		return pm, nil
	}

	log.Info("Found ", len(pools), " pools.")
	log.Info("make pool map...")
	lowerBoundCount := 0
	for _, pool := range pools {
		if len(pool.Assetvolumes) != 2 {
			log.Warn("not enough assets in pool with address: ", pool.Address)
			continue
		}

		liquidity, lowerBound := pool.GetPoolLiquidityUSD()
		// Discard pool if complete USD liquidity is below threshold.
		if !lowerBound && liquidity < liquidityThresholdUSD {
			continue
		}
		if lowerBound {
			lowerBoundCount++
		}

		maverickPair := MaverickPair{
			Address: common.HexToAddress(pool.Address),
		}
		if pool.Assetvolumes[0].Index == 0 {
			maverickPair.Token0 = asset2MaverickAsset(pool.Assetvolumes[0].Asset)
			maverickPair.Token1 = asset2MaverickAsset(pool.Assetvolumes[1].Asset)
		} else {
			maverickPair.Token0 = asset2MaverickAsset(pool.Assetvolumes[1].Asset)
			maverickPair.Token1 = asset2MaverickAsset(pool.Assetvolumes[0].Asset)
		}
		maverickPair.ForeignName = maverickPair.Token0.Symbol + "-" + maverickPair.Token1.Symbol
		pm[pool.Address] = maverickPair
	}

	log.Infof("found %v subscribable pools.", len(pm))
	log.Infof("%v pools with lowerBound=true.", lowerBoundCount)
	return pm, err
}

func asset2MaverickAsset(asset dia.Asset) Token {
	return Token{
		Address:  common.HexToAddress(asset.Address),
		Decimals: asset.Decimals,
		Symbol:   asset.Symbol,
		Name:     asset.Name,
	}
}

// GetSwapsChannel returns a channel for swaps of the pair with address @pairAddress
func (s *MaverickScraper) GetSwapsChannel(pairAddress common.Address) (chan *poolcontract.PoolSwap, error) {

	sink := make(chan *poolcontract.PoolSwap)
	var poolFiltererContract *poolcontract.PoolFilterer
	poolFiltererContract, err := poolcontract.NewPoolFilterer(pairAddress, s.WsClient)
	if err != nil {
		log.Fatal(err)
	}

	_, err = poolFiltererContract.WatchSwap(&bind.WatchOpts{}, sink)
	if err != nil {
		log.Error("error in get swaps channel: ", err)
	}

	return sink, nil

}

func (s *MaverickScraper) getAllPairsAddress() ([]common.Address, error) {
	pools := make([]common.Address, 0)

	var factoryContractInstance *pairfactorycontract.PairfactoryFilterer
	factoryContractInstance, err := pairfactorycontract.NewPairfactoryFilterer(common.HexToAddress(s.poolFactoryContractAddress), s.RestClient)
	if err != nil {
		log.Error(err)
		return pools, err
	}

	currBlock, err := s.RestClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	var offset uint64 = 2500
	startBlock := s.poolFactoryContractCreationBlock
	var endBlock = startBlock + offset

	for {
		if endBlock > currBlock {
			endBlock = currBlock
		}
		log.Infof("startblock - endblock: %v --- %v ", startBlock, endBlock)

		it, err := factoryContractInstance.FilterPoolCreated(
			&bind.FilterOpts{
				Start: startBlock,
				End:   &endBlock,
			})
		if err != nil {
			log.Error(err)
			//return pools, err
			if endBlock == currBlock {
				break
			}

			startBlock = endBlock + 1
			endBlock = endBlock + offset
			continue
		}

		for it.Next() {
			pools = append(pools, it.Event.PoolAddress)
		}
		if err := it.Close(); err != nil {
			log.Warn("closing iterator: ", it)
		}

		if endBlock == currBlock {
			break
		}

		startBlock = endBlock + 1
		endBlock = endBlock + offset
	}

	return pools, err
}

func (s *MaverickScraper) getAllPairs() ([]MaverickPair, error) {
	pairs := make([]MaverickPair, 0)
	addresses, err := s.getAllPairsAddress()
	if err != nil {
		log.Fatal(err)
	}
	for _, address := range addresses {
		pair, err := s.getPairByAddress(address)
		if err != nil {
			log.Warn(err)
			continue
		}
		pairs = append(pairs, pair)
	}
	return pairs, nil
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *MaverickScraper) Close() error {
	if s.closed {
		return errors.New("UniswapScraper: Already closed")
	}
	s.WsClient.Close()
	s.RestClient.Close()
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *MaverickScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("UniswapScraper: Call ScrapePair on closed scraper")
	}
	ps := &MaverickPairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

// FetchAvailablePairs returns a list with all available trade pairs as dia.ExchangePair for the pairDiscorvery service
func (s *MaverickScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	time.Sleep(100 * time.Millisecond)
	maverickPairs, err := s.getAllPairs()
	if err != nil {
		return
	}
	for _, pair := range maverickPairs {
		if !pair.pairHealthCheck() {
			continue
		}
		quotetoken := dia.Asset{
			Symbol:     pair.Token0.Symbol,
			Name:       pair.Token0.Name,
			Address:    pair.Token0.Address.Hex(),
			Decimals:   pair.Token0.Decimals,
			Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
		}
		basetoken := dia.Asset{
			Symbol:     pair.Token1.Symbol,
			Name:       pair.Token1.Name,
			Address:    pair.Token1.Address.Hex(),
			Decimals:   pair.Token1.Decimals,
			Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
		}
		pairToNormalise := dia.ExchangePair{
			Symbol:         pair.Token0.Symbol,
			ForeignName:    pair.ForeignName,
			Exchange:       s.exchangeName,
			Verified:       true,
			UnderlyingPair: dia.Pair{BaseToken: basetoken, QuoteToken: quotetoken},
		}
		normalizedPair, _ := s.NormalizePair(pairToNormalise)
		pairs = append(pairs, normalizedPair)
	}
	return
}

// pairHealthCheck returns true if the involved tokens are not blacklisted and do not have zero entries
func (mp *MaverickPair) pairHealthCheck() bool {
	if mp.Token0.Symbol == "" || mp.Token1.Symbol == "" || mp.Token0.Address.Hex() == "" || mp.Token1.Address.Hex() == "" {
		return false
	}
	if helpers.SymbolIsBlackListed(mp.Token0.Symbol) || helpers.SymbolIsBlackListed(mp.Token1.Symbol) {
		if helpers.SymbolIsBlackListed(mp.Token0.Symbol) {
			log.Infof("skip pair %s. symbol %s is blacklisted", mp.ForeignName, mp.Token0.Symbol)
		} else {
			log.Infof("skip pair %s. symbol %s is blacklisted", mp.ForeignName, mp.Token1.Symbol)
		}
		return false
	}
	if helpers.AddressIsBlacklisted(mp.Token0.Address) || helpers.AddressIsBlacklisted(mp.Token1.Address) {
		log.Info("skip pair ", mp.ForeignName, ", address is blacklisted")
		return false
	}
	return true
}

func (s *MaverickScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// FillSymbolData is not used by DEX scrapers.
func (s *MaverickScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{}, nil
}

// MaverickPairScraper implements PairScraper for Uniswap
type MaverickPairScraper struct {
	parent *MaverickScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *MaverickPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Channel returns a channel that can be used to receive trades
func (s *MaverickScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *MaverickPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *MaverickPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
