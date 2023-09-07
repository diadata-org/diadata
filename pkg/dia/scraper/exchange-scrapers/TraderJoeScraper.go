package scrapers

import (
	"encoding/json"
	"errors"
	"io"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	traderjoe "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/traderjoe2.1"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	TraderJoeExchangeFactoryContractAddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"

	MapOfPools = make(map[string]TraderJoePair)
)

type TraderJoeTokens struct {
	Address  common.Address
	Symbol   string
	Decimals uint8
	Name     string
}

type TraderJoePair struct {
	Token0      TraderJoeTokens
	Token1      TraderJoeTokens
	ForeignName string
	Address     common.Address
}

type TraderJoeSwap struct {
    ID        string
    Timestamp int64
    Pair      TraderJoePair
    Amount0   float64
    Amount1   float64
}


type TraderJoeScraper struct {
	// Ethereum WebSocket client for real-time data.
	WsClient *ethclient.Client
	// Ethereum REST client for querying historical data.
	RestClient *ethclient.Client
	// Relational database connection.
	relDB *models.RelDB
	// Signaling channels for managing session start and shutdown.
	run          bool
	shutdown     chan nothing
	shutdownDone chan nothing
	// Error handling; read lock for error or closed status.
	errorLock sync.RWMutex
	error     error
	closed    bool
	// Map of active TraderJoeTradeScraper instances for trading pairs.
	pairScrapers map[string]*TraderJoeTradeScraper
	// Channel to receive new trading pairs for scraping.
	pairReceived chan *TraderJoePair
	// Name of the exchange.
	exchangeName string
	// Ethereum block number to start scraping from.
	startBlock uint64
	// Time interval for waiting between actions.
	waitTime int
	// Option to listen for trading pairs by address.
	listenByAddress bool
	fetchPoolsFromDB bool
	// Channel for receiving trade data.
	chanTrades chan *dia.Trade
	// Address of the factory contract for the exchange.
	factoryContractAddress common.Address
}

// NewTraderJoeScraper initializes a Trader Joe scraper instance with the provided exchange information,
// scraping flag, and relational database connection. It configures parameters, sets up pool maps,
// and starts the scraping process if requested.
func NewTraderJoeScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *TraderJoeScraper {
	log.Info("NewTraderJoeScraper ", exchange.Name)
	log.Info("NewTraderJoeScraper Address ", exchange.Contract)

	var (
		tjs             *TraderJoeScraper
		listenByAddress bool
		err             error
	)

	listenByAddress, err = strconv.ParseBool(utils.Getenv("LISTEN_BY_ADDRESS", ""))
	if err != nil {
		log.Fatal("parse LISTEN_BY_ADDRESS: ", err)
	}


	// TODO: revisit block numbers
	switch exchange.Name {
	case dia.TraderJoeExchangeV2_1:
		tjs = makeTraderJoeScraper(exchange, listenByAddress, "", "", "200", uint64(22757913))
	case dia.TraderJoeExchangeV2_1Arbitrum:
		tjs = makeTraderJoeScraper(exchange, listenByAddress, "", "", "200", uint64(107808631)) // startBlock pulled from: https://arbiscan.io/tx/0x7e1d7980bdbfb09b0a64d35e660cee300accaf017674db7a3ac9f655ca72ac35
	case dia.TraderJoeExchangeV2_1Avalanche:
		tjs = makeTraderJoeScraper(exchange, listenByAddress, "", "", "200", uint64(165))
	case dia.TraderJoeExchangeV2_1BNB:
		tjs = makeTraderJoeScraper(exchange, listenByAddress, "", "", "200", uint64(27099340)) // startBlock pulled from: https://bscscan.com/tx/0x4c8a1f2f3d1a92281a3067aaea799536118e02a072380ccd3df642e8adecba6d
	}

	tjs.relDB = relDB

	// Only include pools with (minimum) liquidity bigger than given env var.
	liquidityThreshold, err := strconv.ParseFloat(utils.Getenv("LIQUIDITY_THRESHOLD", "0"), 64)
	if err != nil {
		liquidityThreshold = float64(0)
		log.Warnf("parse liquidity threshold:  %v. Set to default %v", err, liquidityThreshold)
	}
	liquidityThresholdUSD, err := strconv.ParseFloat(utils.Getenv("LIQUIDITY_THRESHOLD_USD", "0"), 64)
	if err != nil {
		liquidityThresholdUSD = float64(0)
		log.Warnf("parse liquidity threshold:  %v. Set to default %v", err, liquidityThresholdUSD)
	}

	MapOfPools, err = tjs.makeTraderJoePoolMap(liquidityThreshold, liquidityThresholdUSD)
	if err != nil {
		log.Fatal("build poolMap: ", err)
	}

	if scrape {
		go tjs.mainLoop()
	}
	return tjs
}

// makeTraderJoeScraper creates and initializes a Trader Joe scraper instance with the given exchange information,
// connection details, and configuration parameters. It establishes REST and WebSocket clients for the blockchain,
// determines wait time, and sets up various channels and data structures for scraping tasks.
func makeTraderJoeScraper(exchange dia.Exchange, listenByAddress bool, restDial string, wsDial string, waitMilliseconds string, startBlock uint64) *TraderJoeScraper {
	var (
		restClient, wsClient *ethclient.Client
		s *TraderJoeScraper
		err error
	)

	log.Infof("Initialize rest and ws client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", restDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}
	wsClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", wsDial))
	if err != nil {
		log.Fatal("init ws client: ", err)
	}

	var waitTime int
	waitTimeString := utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_WAIT_TIME", waitMilliseconds)
	waitTime, err = strconv.Atoi(waitTimeString)
	if err != nil {
		log.Error("could not parse wait time: ", err)
		waitTime = 500
	}

	s = &TraderJoeScraper{
		WsClient:               wsClient,
		RestClient:             restClient,
		shutdown:               make(chan nothing),
		shutdownDone:           make(chan nothing),
		pairScrapers:           make(map[string]*TraderJoeTradeScraper),
		exchangeName:           exchange.Name,
		pairReceived:           make(chan *TraderJoePair),
		error:                  nil,
		chanTrades:             make(chan *dia.Trade),
		waitTime:               waitTime,
		listenByAddress:        listenByAddress,
		startBlock:             startBlock,
		factoryContractAddress: common.HexToAddress(exchange.Contract),
	}

	return s
}

// GetSwapsChannel returns a channel for swaps of the pair with pair address.
func (tjs *TraderJoeScraper) GetSwapsChannel(pairAddress common.Address) (chan *traderjoe.TraderjoeLBPairCreated, error) {
	sink := make(chan *traderjoe.TraderjoeLBPairCreated)
	var pairFiltererContract *traderjoe.TraderjoeFilterer

	pairFiltererContract, err := traderjoe.NewTraderjoeFilterer(pairAddress, tjs.WsClient)
	if err != nil {
		log.Fatal(err)
	}
	_, err = pairFiltererContract.WatchOwnershipTransferred(&bind.WatchOpts{}, sink, []common.Address{}, []common.Address{})
	if err != nil {
		log.Error("error in get swaps channel: ", err)
	}

	return sink, nil
}

// TODO: Return here.
func (tjs *TraderJoeScraper) normalizeTraderJoeSwap(swapI interface{}) (normalizedSwap TraderJoeSwap, err error) {
	switch swap := swapI.(type) {
	case traderjoe.TraderjoeLBPairCreated:
		pair := MapOfPools[swap.Raw.Address.Hex()]
		decimals0 := int(pair.Token0.Decimals)
		decimals1 := int(pair.Token1.Decimals)
		amount0, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(), new(big.Float).SetFloat64(math.Pow10(decimals0))).Float64()
		amount1, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(), new(big.Float).SetFloat64(math.Pow10(decimals1))).Float64()

		normalizedSwap = TraderJoeSwap{
			ID:        swap.Raw.TxHash.Hex(),
			Timestamp: time.Now().Unix(),
			Pair:      pair,
			Amount0:   amount0,
			Amount1:   amount1,
		}
	}
	return
}

// mainLoop is the central loop of the Trader Joe scraper that manages the subscription and scraping of pairs.
// It initializes the process by retrieving reverse base tokens and quote tokens from configuration. After a brief
// initial delay, it sets the `run` flag to true and kicks off a goroutine to feed pools to subscriptions.
// The function then listens for incoming pairs from the `pairReceived` channel and subscribes to and scrapes data
// for each pair. It performs various checks to skip pairs that don't meet certain criteria, such as blacklisted
// tokens or pools. It also logs relevant information about the progress of the loop.
func (tjs *TraderJoeScraper) mainLoop() {
	var err error
	reverseBasetokens, err = getReverseTokensFromConfig("traderjoe/reverse_tokens/" + tjs.exchangeName + "Basetoken")
	if err != nil {
		log.Error("error getting base tokens for which pairs should be reversed: ", err)
	}
	log.Infof("reverse the following basetokens on %s: %v", tjs.exchangeName, reverseBasetokens)
	reverseQuotetokens, err = getReverseTokensFromConfig("traderjoe/reverse_tokens/" + tjs.exchangeName + "Quotetoken")
	if err != nil {
		log.Error("error getting quote tokens for which pairs should be reversed: ", err)
	}
	log.Infof("reverse the following quotetokens on %s: %v", tjs.exchangeName, reverseQuotetokens)

	time.Sleep(4 * time.Second)
	tjs.run = true

	go func() {
		pools := tjs.feedPoolsToSubscriptions()
		log.Info("Found ", len(pools), " pairs")
		log.Info("Found ", len(tjs.pairScrapers), " pairScrapers")
	}()

	if len(tjs.pairScrapers) == 0 {
		tjs.error = errors.New("traderjoe scraper: No pairs to scrape provided")
		log.Error(tjs.error.Error())
	}

	count := 0
	for {
		pool := <-tjs.pairReceived
		log.Infoln("Subscribing for pair: ", pool)

		if len(pool.Token0.Symbol) < 2 || len(pool.Token1.Symbol) < 2 {
			log.Info("skip pair: ", pool.ForeignName)
			continue
		}
		if helpers.AddressIsBlacklisted(pool.Token0.Address) || helpers.AddressIsBlacklisted(pool.Token1.Address) {
			log.Info("skip pair ", pool.ForeignName, ", address is blacklisted")
			continue
		}
		if helpers.PoolIsBlacklisted(pool.Address) {
			log.Info("skip blacklisted pool ", pool.Address)
			continue
		}
		log.Infof("%v found pair scraper for: %s with address %s", count, pool.ForeignName, pool.Address.Hex())
		count++
		
		sink, err := tjs.GetSwapsChannel(pool.Address)
		if err != nil {
			log.Error("error fetching swaps channel: ", err)
		}
		go func() {
			for {
				rawSwap, ok := <-sink
				if ok {
					swap := tjs.normalizeTraderJoeSwap(*rawSwap)
					tjs.sendTrade(swap, pool)
				}
			}
		}()
		
	}
}

// makeTraderJoePoolMap generates a map of Trader Joe pool pairs based on the provided liquidity thresholds and configuration.
// It retrieves pool information either by specific addresses from a JSON file or by querying the database for all pools above
// the liquidity threshold. The resulting pool map includes pairs with sufficient liquidity and handles lower-bound checks.
// It returns the generated pool map and any error encountered during the process.
func (tjs *TraderJoeScraper) makeTraderJoePoolMap(liquidityThreshold, liquidityThresholdUSD float64) (map[string]TraderJoePair, error) {
	poolMap := make(map[string]TraderJoePair)
	var (
		pools []dia.Pool
		err   error
	)

	if tjs.listenByAddress {
		// Only load pool info for addresses from json file.
		poolAddresses, errAddr := getTradeAddressesFromConfig("traderjoe/subscribe_pools/" + tjs.exchangeName)
		// TODO: is this address correct?
		if errAddr != nil {
			log.Error("fetch pool addresses from config file: ", errAddr)
		}
		for _, address := range poolAddresses {
			pool, errPool := tjs.relDB.GetPoolByAddress(Exchanges[tjs.exchangeName].BlockChain.Name, address.Hex())
			if errPool != nil {
				log.Fatalf("Get pool with address %s: %v", address.Hex(), errPool)
			}
			pools = append(pools, pool)
		}
	} else {
		// Load all pools above liquidity threshold.
		pools, err = tjs.relDB.GetAllPoolsExchange(tjs.exchangeName, liquidityThreshold)
		if err != nil {
			return poolMap, err
		}
	}

	log.Info("Found ", len(pools), " pools.")
	log.Info("make pool map...")
	lowerBoundCount := 0
	for _, pool := range pools {
		if len(pool.Assetvolumes) != 2 {
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

		up := TraderJoePair{
			Address: common.HexToAddress(pool.Address),
		}
		if pool.Assetvolumes[0].Index == 0 {
			up.Token0 = asset2TraderJoeAsset(pool.Assetvolumes[0].Asset)
			up.Token1 = asset2TraderJoeAsset(pool.Assetvolumes[1].Asset)
		} else {
			up.Token0 = asset2TraderJoeAsset(pool.Assetvolumes[1].Asset)
			up.Token1 = asset2TraderJoeAsset(pool.Assetvolumes[0].Asset)
		}
		up.ForeignName = up.Token0.Symbol + "-" + up.Token1.Symbol
		poolMap[pool.Address] = up
	}

	log.Infof("found %v subscribable pools.", len(poolMap))
	log.Infof("%v pools with lowerBound=true.", lowerBoundCount)

	return poolMap, err
}

func (tjs *TraderJoeScraper) subscribeToSwapEvents(pair TraderJoePair) {}

// sendTrade receives Trader Joe trade data and transforms it into a standardized dia.Trade
// structure for further analysis and publication.
func (tjs *TraderJoeScraper) sendTrade(traderjoeswap TraderJoeSwap, pool *TraderJoePair) {
	price, volume := tjs.getTradeData(traderjoeswap)
	token0 := dia.Asset{
		Address:    pool.Token0.Address.Hex(),
		Symbol:     pool.Token0.Symbol,
		Name:       pool.Token0.Name,
		Decimals:   pool.Token0.Decimals,
		Blockchain: Exchanges[tjs.exchangeName].BlockChain.Name,
	}
	token1 := dia.Asset{
		Address:    pool.Token1.Address.Hex(),
		Symbol:     pool.Token1.Symbol,
		Name:       pool.Token1.Name,
		Decimals:   pool.Token1.Decimals,
		Blockchain: Exchanges[tjs.exchangeName].BlockChain.Name,
	}

	t := &dia.Trade{
		Symbol:         pool.Token0.Symbol,
		Pair:           pool.ForeignName,
		QuoteToken:     token1,
		BaseToken:      token0,
		Price:          price,
		Volume:         volume,
		Time:           time.Unix(traderjoeswap.Timestamp, 0),
		PoolAddress:    pool.Address.Hex(),
		ForeignTradeID: traderjoeswap.ID,
		//EstimatedUSDPrice: 0,
		Source:       tjs.exchangeName,
		VerifiedPair: true,
	}

	switch {
	case utils.Contains(reverseBasetokens, pool.Token1.Address.Hex()):
		// If we need quotation of a base token, reverse pair
		tSwapped, err := dia.SwapTrade(*t)
		if err == nil {
			t = &tSwapped
		}
	case utils.Contains(reverseQuotetokens, pool.Token0.Address.Hex()):
		// If we need quotation of a base token, reverse pair
		tSwapped, err := dia.SwapTrade(*t)
		if err == nil {
			t = &tSwapped
		}
	}
	if price > 0 {
		log.Infof("Got trade on pool %s: %v", pool.Address.Hex(), t)
		tjs.chanTrades <- t
	}
}

func (tjs *TraderJoeScraper) getSwapData(swap TraderJoeSwap) (price, volume float64) {
	volume = swap.Amount0
	price = math.Abs(swap.Amount1 / swap.Amount0)
	return
}

// pairHealthCheck returns true if the involved tokens are not blacklisted and do not have zero entries
func (tp *TraderJoePair) pairHealthCheck() bool {
	if tp.Token0.Symbol == "" || tp.Token1.Symbol == "" || tp.Token0.Address.Hex() == "" || tp.Token1.Address.Hex() == "" {
		return false
	}
	if helpers.SymbolIsBlackListed(tp.Token0.Symbol) || helpers.SymbolIsBlackListed(tp.Token1.Symbol) {
		if helpers.SymbolIsBlackListed(tp.Token0.Symbol) {
			log.Infof("skip pair %s. symbol %s is blacklisted", tp.ForeignName, tp.Token0.Symbol)
		} else {
			log.Infof("skip pair %s. symbol %s is blacklisted", tp.ForeignName, tp.Token1.Symbol)
		}
		return false
	}
	if helpers.AddressIsBlacklisted(tp.Token0.Address) || helpers.AddressIsBlacklisted(tp.Token1.Address) {
		log.Info("skip pair ", tp.ForeignName, ", address is blacklisted")
		return false
	}
	return true
}

// FetchAvailablePairs retrieves the list of available dia.ExchangePairs from the Trader Joe exchange.
func (tjs *TraderJoeScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	time.Sleep(100 * time.Millisecond)
	
	traderjoePair, err := tjs.GetAllPairs()
	if err != nil {
		return
	}
	for _, pair := range traderjoePair {
		if !pair.pairHealthCheck() {
			continue
		}
		quotetoken := dia.Asset{
			Symbol:     pair.Token0.Symbol,
			Name:       pair.Token0.Name,
			Address:    pair.Token0.Address.Hex(),
			Decimals:   pair.Token0.Decimals,
			Blockchain: Exchanges[tjs.exchangeName].BlockChain.Name,
		}
		basetoken := dia.Asset{
			Symbol:     pair.Token1.Symbol,
			Name:       pair.Token1.Name,
			Address:    pair.Token1.Address.Hex(),
			Decimals:   pair.Token1.Decimals,
			Blockchain: Exchanges[tjs.exchangeName].BlockChain.Name,
		}
		pairToNormalise := dia.ExchangePair{
			Symbol:         pair.Token0.Symbol,
			ForeignName:    pair.ForeignName,
			Exchange:       "TraderJoeExchange",
			Verified:       true,
			UnderlyingPair: dia.Pair{BaseToken: basetoken, QuoteToken: quotetoken},
		}
		normalizedPair, _ := tjs.NormalizePair(pairToNormalise)
		pairs = append(pairs, normalizedPair)
	}

	return
}

// GetAllPairs is similar to FetchAvailablePairs. But instead of dia.ExchangePairs it returns all pairs as TraderJoePairs,
// i.e. including the pair's address
func (tjs *TraderJoeScraper) GetAllPairs() ([]TraderJoePair, error) {
	time.Sleep(20 * time.Millisecond)
	connection := tjs.RestClient
	var contract *traderjoe.TraderjoeCaller
	contract, err := traderjoe.NewTraderjoeCaller(common.HexToAddress(exchangeFactoryContractAddress), connection)
	if err != nil {
		log.Error(err)
	}

	numPairs, err := contract.GetAllLBPairs(&bind.CallOpts{}, common.Address{}, common.Address{})
	if err != nil {
		return []TraderJoePair{}, err
	}
	wg := sync.WaitGroup{}
	defer wg.Wait()
	pairs := make([]TraderJoePair, int(numPairs.Int64()))
	for i := 0; i < int(numPairs.Int64()); i++ {
		// Sleep in order not to run into rate limits.
		time.Sleep(time.Duration(tjs.waitTime) * time.Millisecond)
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			traderJoePair, err := tjs.GetPairByID(int64(index))
			if err != nil {
				log.Error("error retrieving pair by ID: ", err)
				return
			}
			pairs[index] = traderJoePair
		}(i)
	}
	return pairs, nil
}

func (tjs *TraderJoeScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (tjs *TraderJoeScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// TODO
func (tjs *TraderJoeScraper) GetPairData() {}

// TraderJoeTradeScraper represents a scraper for collecting trade data associated with a specific dia.ExchangePair
// within the Trader Joe exchange.
type TraderJoeTradeScraper struct {
	parent *TraderJoeScraper
	pair   dia.ExchangePair
}

// Close closes any existing API connections, as well as channels of PairScrapers from calls to ScrapePair
func (tjs *TraderJoeScraper) Close() error {
	if tjs.closed {
		return errors.New("TraderJoeScraper: Already closed")
	}
	tjs.WsClient.Close()
	tjs.RestClient.Close()
	close(tjs.shutdown)
	<-tjs.shutdownDone
	tjs.errorLock.RLock()
	defer tjs.errorLock.RUnlock()
	return tjs.error
}

// TODO
func (tjs *TraderJoeScraper) ListenToPair(i int, address common.Address) {
	var (
		pair TraderJoePair
		err error
	)
	if !tjs.listenByAddress && !tjs.fetchPoolsFromDB {
		// Get pool info from on-chain, when MapOfPools is empty.
		pair, err = tjs.GetPairByID(int64(i))
		if err != nil {
			log.Error("error fetching pair: ", err)
		} else {
			// relevant pool is retrieved from MapOfPools.
		}
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

	log.Info(i, ": add pair scraper for: ", pair.ForeignName, " with address ", pair.Address.Hex())
	sink, err := tjs.GetSwapsChannel(pair.Address)
	if err != nil {
		log.Error("error fetching swaps channel: ", err)
	}

	go func() {
		for {
			rawSwap, ok := <-sink
			if ok {
				swap, err := tjs.normalizeTraderJoeSwap(*rawSwap, pair)
				if err != nil {
					log.Error("error normalizing swap: ", err)
				}
				price, volume := getSwapData(swap)
				token0 := dia.Asset{
					Address:    pair.Token0.Address.Hex(),
					Symbol:     pair.Token0.Symbol,
					Name:       pair.Token0.Name,
					Decimals:   pair.Token0.Decimals,
					Blockchain: Exchanges[tjs.exchangeName].BlockChain.Name,
				}
				token1 := dia.Asset{
					Address:    pair.Token1.Address.Hex(),
					Symbol:     pair.Token1.Symbol,
					Name:       pair.Token1.Name,
					Decimals:   pair.Token1.Decimals,
					Blockchain: Exchanges[tjs.exchangeName].BlockChain.Name,
				}
				t := &dia.Trade{
					Symbol:         pair.Token0.Symbol,
					Pair:           pair.ForeignName,
					Price:          price,
					Volume:         volume,
					BaseToken:      token1,
					QuoteToken:     token0,
					Time:           time.Unix(swap.Timestamp, 0),
					PoolAddress:    rawSwap.Raw.Address.Hex(),
					ForeignTradeID: swap.ID,
					Source:         tjs.exchangeName,
					VerifiedPair:   true,
				}

				// // TO DO: Refactor approach for reversing pairs.
				// switch {
				// case utils.Contains(reverseBasetokens, pair.Token1.Address.Hex()):
				// 	// If we need quotation of a base token, reverse pair
				// 	tSwapped, err := dia.SwapTrade(*t)
				// 	if err == nil {
				// 		t = &tSwapped
				// 	}
				// case utils.Contains(reverseQuotetokens, pair.Token0.Address.Hex()):
				// 	// If we don't need quotation of quote token, reverse pair.
				// 	tSwapped, err := dia.SwapTrade(*t)
				// 	if err == nil {
				// 		t = &tSwapped
				// 	}
				// case token0.Address == "0x...REPLACE" && !utils.Contains(&mainBaseAssets, token1.Address):
				// 	// Reverse almost all pairs WETH-XXX ...
				// 	if tjs.exchangeName == dia.TraderJoeExchangeV2_1 {
				// 		tSwapped, err := dia.SwapTrade(*t)
				// 		if err == nil {
				// 			t = &tSwapped
				// 		}
				// 	}
				// // ...and USDT-XXX on Ethereum,
				// case token0.Address == mainBaseAssets[0] && token0.Blockchain == dia.ETHEREUM:
				// 	tSwapped, err := dia.SwapTrade(*t)
				// 	if err == nil {
				// 		t = &tSwapped
				// 	}
				// // Reverse USDC-XXX pairs on Fantom
				// case token0.Address == "0x..REPLACE" && token0.Blockchain == dia.FANTOM:
				// 	tSwapped, err := dia.SwapTrade(*t)
				// 	if err == nil {
				// 		t = &tSwapped
				// 	}
				// }
				if price > 0 {
					log.Info("tx hash: ", swap.ID)
					log.Infof("Got trade at time %v - symbol: %s, pair: %s, price: %v, volume:%v", t.Time, t.Symbol, t.Pair, t.Price, t.Volume)
					// log.Infof("Base token info --- Symbol: %s - Address: %s - Blockchain: %s ", t.BaseToken.Symbol, t.BaseToken.Address, t.BaseToken.Blockchain)
					// log.Info("----------------")
					tjs.chanTrades <- t
				}
			}
		}
	}()
}

// TODO
// GetPairByID returns the UniswapPair with the integer id @num
func (tjs *TraderJoeScraper) GetPairByID(num int64) (TraderJoePair, error) {
	var contract *traderjoe.TraderjoeCaller
	contract, err := traderjoe.NewTraderjoeCaller(common.HexToAddress(exchangeFactoryContractAddress), tjs.RestClient)
	if err != nil {
		log.Error(err)
		return TraderJoePair{}, err
	}

	pairAddress, err := contract.GetAllLBPairs(&bind.CallOpts{}, common.Address{}, common.Address{})
	if err != nil {
		log.Error(err)
		return TraderJoePair{}, err
	}

	pair, err := tjs.GetPairByAddress(pairAddress)
	if err != nil {
		log.Error(err)
		return TraderJoePair{}, err
	}
	return pair, err
}

// getTradeData extracts price and volume data from TraderJoe trade information.
func (tjs *TraderJoeScraper) getTradeData(swap TraderJoeSwap) (price, volume float64) {
	volume = swap.Amount0
	price = math.Abs(swap.Amount1 / swap.Amount0)
	return
}

// ScrapePair initiates a new scraping process for the specified dia.ExchangePair within the Trader Joe scraper.
// It checks for any previously encountered errors using a read lock on the error lock. If an error is present,
// it returns that error. Additionally, if the Trader Joe scraper has been closed, it returns an error indicating
// that ScrapePair cannot be called on a closed pair. Otherwise, it creates a new TraderJoeTradeScraper instance
// associated with the provided ExchangePair, adds it to the list of active pair scrapers, and returns it along
// with a nil error to indicate successful initiation.
func (tjs *TraderJoeScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	tjs.errorLock.RLock()
	defer tjs.errorLock.RUnlock()
	if tjs.error != nil {
		return nil, tjs.error
	}

	if tjs.closed {
		return nil, errors.New("TraderJoeScraper: Call Scrape Pair on closed pair")
	}

	pairScraper := &TraderJoeTradeScraper{
		parent: tjs,
		pair:   pair,
	}

	tjs.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}

// Close closes the TraderJoeTradeScraper instance.
func (ps TraderJoeTradeScraper) Close() error {
	return nil
}

// Error returns the error associated with the parent Trader Joe scraper. It retrieves the error from the parent scraper's state
// using a read lock on the error lock. This function is useful for obtaining any error that occurred during scraping tasks.
func (ps TraderJoeTradeScraper) Error() error {
	tjs := ps.parent
	tjs.errorLock.RLock()
	defer tjs.errorLock.RUnlock()
	return tjs.error
}

// Pair returns the dia.ExchangePair associated with the current Trader Joe trade scraper.
// It simply retrieves and returns the ExchangePair stored within the scraper's state.
func (ps TraderJoeTradeScraper) Pair() dia.ExchangePair {
	return ps.pair
}

// Channel returns a channel that can be used to receive trades
func (tjs *TraderJoeScraper) Channel() chan *dia.Trade {
	return tjs.chanTrades
}

// asset2TraderJoeAsset converts a dia.Asset into a TraderJoeTokens structure.
// It takes the provided asset's address, decimals, symbol, and name,
// and returns a TraderJoeTokens representation containing the same information.
func asset2TraderJoeAsset(asset dia.Asset) TraderJoeTokens {
	return TraderJoeTokens{
		Address:  common.HexToAddress(asset.Address),
		Decimals: asset.Decimals,
		Symbol:   asset.Symbol,
		Name:     asset.Name,
	}
}

// getTradeAddressesFromConfig reads a JSON configuration file specified by the provided filename and retrieves
// trading pair addresses. The function opens and reads the file, unmarshals the data to extract pairs' addresses
// and foreign names, and returns a slice of common.Address containing the extracted addresses. In case of any
func getTradeAddressesFromConfig(filename string) (pairAddresses []common.Address, err error) {

	// Load file and read data
	fileHandle := configCollectors.ConfigFileConnectors(filename, ".json")
	jsonFile, err := os.Open(fileHandle)
	if err != nil {
		return
	}
	defer func() {
		err = jsonFile.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	byteData, err := io.ReadAll(jsonFile)
	if err != nil {
		return
	}

	// Unmarshal read data
	type scrapedPair struct {
		Address     string `json:"Address"`
		ForeignName string `json:"ForeignName"`
	}
	type scrapedPairList struct {
		AllPairs []scrapedPair `json:"Pools"`
	}
	var allPairs scrapedPairList
	err = json.Unmarshal(byteData, &allPairs)
	if err != nil {
		return
	}

	// Extract addresses
	for _, token := range allPairs.AllPairs {
		pairAddresses = append(pairAddresses, common.HexToAddress(token.Address))
	}

	return
}

// feedPoolsToSubscriptions sends a list of TraderJoePairs to subscription channels.
func (tjs *TraderJoeScraper) feedPoolsToSubscriptions() (pairs []TraderJoePair) {
	for i := range MapOfPools {
		up := MapOfPools[i]
		pairs = append(pairs, up)
		tjs.pairReceived <- &up
	}
	return
}