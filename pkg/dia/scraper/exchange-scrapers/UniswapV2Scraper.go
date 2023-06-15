package scrapers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswap"
	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	exchangeFactoryContractAddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
	reverseBasetokens              *[]string
	reverseQuotetokens             *[]string
	mainBaseAssets                 = []string{
		"0xdAC17F958D2ee523a2206206994597C13D831ec7",
	}
	poolMap = make(map[string]UniswapPair)
)

const (
	restDialEth = ""
	wsDialEth   = ""

	restDialBSC = ""
	wsDialBSC   = ""

	restDialPolygon = ""
	wsDialPolygon   = ""

	restDialCelo = ""
	wsDialCelo   = ""

	restDialFantom = ""
	wsDialFantom   = ""

	restDialMoonriver = ""
	wsDialMoonriver   = ""

	restDialAurora = ""
	wsDialAurora   = ""

	restDialArbitrum = ""
	wsDialArbitrum   = ""

	restDialMetis = ""
	wsDialMetis   = ""

	restDialAvalanche = ""
	wsDialAvalanche   = ""

	restDialTelos = ""
	wsDialTelos   = ""

	restDialEvmos = ""
	wsDialEvmos   = ""

	restDialAstar = ""
	wsDialAstar   = ""

	restDialMoonbeam = ""
	wsDialMoonbeam   = ""

	restDialWanchain = ""
	wsDialWanchain   = ""

	uniswapWaitMilliseconds     = "25"
	sushiswapWaitMilliseconds   = "100"
	pancakeswapWaitMilliseconds = "200"
	dfynWaitMilliseconds        = "100"
	quickswapWaitMilliseconds   = "200"
	ubeswapWaitMilliseconds     = "200"
	spookyswapWaitMilliseconds  = "200"
	solarbeamWaitMilliseconds   = "400"
	trisolarisWaitMilliseconds  = "200"
	metisWaitMilliseconds       = "200"
	moonriverWaitMilliseconds   = "500"
	avalancheWaitMilliseconds   = "200"
	telosWaitMilliseconds       = "400"
	evmosWaitMilliseconds       = "400"
	astarWaitMilliseconds       = "1000"
	moonbeamWaitMilliseconds    = "1000"
	wanchainWaitMilliseconds    = "1000"
)

type UniswapToken struct {
	Address  common.Address
	Symbol   string
	Decimals uint8
	Name     string
}

type UniswapPair struct {
	Token0      UniswapToken
	Token1      UniswapToken
	ForeignName string
	Address     common.Address
}

type UniswapSwap struct {
	ID         string
	Timestamp  int64
	Pair       UniswapPair
	Amount0In  float64
	Amount0Out float64
	Amount1In  float64
	Amount1Out float64
}

type UniswapScraper struct {
	WsClient   *ethclient.Client
	RestClient *ethclient.Client
	relDB      *models.RelDB
	datastore  *models.DB
	// signaling channels for session initialization and finishing
	//initDone     chan nothing
	run          bool
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*UniswapPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
	waitTime     int
	// If true, only pairs given in config file are scraped. Default is false.
	listenByAddress  bool
	fetchPoolsFromDB bool
}

// NewUniswapScraper returns a new UniswapScraper for the given pair
func NewUniswapScraper(exchange dia.Exchange, scrape bool) *UniswapScraper {
	log.Info("NewUniswapScraper: ", exchange.Name)
	var (
		s                *UniswapScraper
		listenByAddress  bool
		fetchPoolsFromDB bool
		err              error
	)
	exchangeFactoryContractAddress = exchange.Contract

	listenByAddress, err = strconv.ParseBool(utils.Getenv("LISTEN_BY_ADDRESS", ""))
	if err != nil {
		log.Fatal("parse LISTEN_BY_ADDRESS: ", err)
	}

	fetchPoolsFromDB, err = strconv.ParseBool(utils.Getenv("FETCH_POOLS_FROM_DB", ""))
	if err != nil {
		log.Fatal("parse FETCH_POOLS_FROM_DB: ", err)
	}

	switch exchange.Name {
	case dia.UniswapExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialEth, wsDialEth, uniswapWaitMilliseconds)
	case dia.SushiSwapExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialEth, wsDialEth, sushiswapWaitMilliseconds)
	case dia.SushiSwapExchangePolygon:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialPolygon, wsDialPolygon, metisWaitMilliseconds)
	case dia.SushiSwapExchangeFantom:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialFantom, wsDialFantom, metisWaitMilliseconds)
	case dia.SushiSwapExchangeArbitrum:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialArbitrum, wsDialArbitrum, metisWaitMilliseconds)
	case dia.CamelotExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialArbitrum, wsDialArbitrum, metisWaitMilliseconds)
	case dia.PanCakeSwap:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialBSC, wsDialBSC, pancakeswapWaitMilliseconds)
	case dia.DfynNetwork:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialPolygon, wsDialPolygon, dfynWaitMilliseconds)
	case dia.QuickswapExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialPolygon, wsDialPolygon, quickswapWaitMilliseconds)
	case dia.UbeswapExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialCelo, wsDialCelo, ubeswapWaitMilliseconds)
	case dia.SpookyswapExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialFantom, wsDialFantom, spookyswapWaitMilliseconds)
	case dia.SpiritswapExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialFantom, wsDialFantom, spookyswapWaitMilliseconds)
	case dia.SolarbeamExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialMoonriver, wsDialMoonriver, solarbeamWaitMilliseconds)
	case dia.TrisolarisExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialAurora, wsDialAurora, trisolarisWaitMilliseconds)
	case dia.NetswapExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialMetis, wsDialMetis, metisWaitMilliseconds)
	case dia.HuckleberryExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialMoonriver, wsDialMoonriver, moonriverWaitMilliseconds)
	case dia.TraderJoeExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialAvalanche, wsDialAvalanche, avalancheWaitMilliseconds)
	case dia.PangolinExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialAvalanche, wsDialAvalanche, avalancheWaitMilliseconds)
	case dia.TethysExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialMetis, wsDialMetis, metisWaitMilliseconds)
	case dia.HermesExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialMetis, wsDialMetis, metisWaitMilliseconds)
	case dia.OmniDexExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialTelos, wsDialTelos, telosWaitMilliseconds)
	case dia.DiffusionExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialEvmos, wsDialEvmos, evmosWaitMilliseconds)
	case dia.ApeswapExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialBSC, wsDialBSC, pancakeswapWaitMilliseconds)
	case dia.BiswapExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialBSC, wsDialBSC, pancakeswapWaitMilliseconds)
	case dia.ArthswapExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialAstar, wsDialAstar, astarWaitMilliseconds)
	case dia.StellaswapExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialMoonbeam, wsDialMoonbeam, moonbeamWaitMilliseconds)
	case dia.WanswapExchange:
		s = makeUniswapScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialWanchain, wsDialWanchain, wanchainWaitMilliseconds)
	}

	s.relDB, err = models.NewPostgresDataStore()
	if err != nil {
		log.Fatal("new postgres datastore: ", err)
	}

	s.datastore, err = models.NewDataStore()
	if err != nil {
		log.Fatal("new datastore: ", err)
	}

	// Only include pools with (minimum) liquidity bigger than given env var.
	liquidityThreshold, err := strconv.ParseFloat(utils.Getenv("LIQUIDITY_THRESHOLD", "0"), 64)
	if err != nil {
		liquidityThreshold = float64(0)
		log.Warnf("parse liquidity threshold:  %v. Set to default %v", err, liquidityThreshold)
	}

	// Only include pools with (minimum) liquidity USD value bigger than given env var.
	liquidityThresholdUSD, err := strconv.ParseFloat(utils.Getenv("LIQUIDITY_THRESHOLD_USD", "0"), 64)
	if err != nil {
		liquidityThreshold = float64(0)
		log.Warnf("parse liquidity threshold:  %v. Set to default %v", err, liquidityThreshold)
	}

	// Fetch all pool with given liquidity threshold from database.
	poolMap, err = s.makeUniPoolMap(liquidityThreshold, liquidityThresholdUSD)
	if err != nil {
		log.Fatal("build poolMap: ", err)
	}

	if scrape {
		go s.mainLoop()
	}
	return s
}

// makeUniswapScraper returns a uniswap scraper as used in NewUniswapScraper.
func makeUniswapScraper(exchange dia.Exchange, listenByAddress bool, fetchPoolsFromDB bool, restDial string, wsDial string, waitMilliseconds string) *UniswapScraper {
	var (
		restClient, wsClient *ethclient.Client
		err                  error
		s                    *UniswapScraper
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
		waitTime = 500
	}

	s = &UniswapScraper{
		WsClient:         wsClient,
		RestClient:       restClient,
		shutdown:         make(chan nothing),
		shutdownDone:     make(chan nothing),
		pairScrapers:     make(map[string]*UniswapPairScraper),
		exchangeName:     exchange.Name,
		error:            nil,
		chanTrades:       make(chan *dia.Trade),
		waitTime:         waitTime,
		listenByAddress:  listenByAddress,
		fetchPoolsFromDB: fetchPoolsFromDB,
	}
	return s
}

// runs in a goroutine until s is closed
func (s *UniswapScraper) mainLoop() {

	// Import tokens which appear as base token and we need a quotation for
	var err error
	reverseBasetokens, err = getReverseTokensFromConfig("uniswap/reverse_tokens/" + s.exchangeName + "Basetoken")
	if err != nil {
		log.Error("error getting tokens for which pairs should be reversed: ", err)
	}
	log.Info("reverse basetokens: ", reverseBasetokens)
	reverseQuotetokens, err = getReverseTokensFromConfig("uniswap/reverse_tokens/" + s.exchangeName + "Quotetoken")
	if err != nil {
		log.Error("error getting tokens for which pairs should be reversed: ", err)
	}
	log.Info("reverse quotetokens: ", reverseQuotetokens)

	// wait for all pairs have added into s.PairScrapers
	time.Sleep(4 * time.Second)
	s.run = true

	if s.listenByAddress || s.fetchPoolsFromDB {

		var wg sync.WaitGroup
		count := 0
		for address := range poolMap {
			time.Sleep(time.Duration(s.waitTime) * time.Millisecond)
			wg.Add(1)
			go func(index int, address common.Address, w *sync.WaitGroup) {
				defer w.Done()
				s.ListenToPair(index, address)
			}(count, common.HexToAddress(address), &wg)
			count++
		}
		wg.Wait()

	} else {

		numPairs, err := s.getNumPairs()
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Found ", numPairs, " pairs")
		log.Info("Found ", len(s.pairScrapers), " pairScrapers")

		if len(s.pairScrapers) == 0 {
			s.error = errors.New("uniswap: No pairs to scrap provided")
			log.Error(s.error.Error())
		}

		var wg sync.WaitGroup
		for i := 0; i < numPairs; i++ {
			time.Sleep(time.Duration(s.waitTime) * time.Millisecond)
			wg.Add(1)
			go func(index int, address common.Address, w *sync.WaitGroup) {
				defer w.Done()
				s.ListenToPair(index, address)
			}(i, common.Address{}, &wg)
		}
		wg.Wait()

	}
}

// ListenToPair subscribes to a uniswap pool.
// If @byAddress is true, it listens by pool address, otherwise by index.
func (s *UniswapScraper) ListenToPair(i int, address common.Address) {
	var (
		pair UniswapPair
		err  error
	)

	if !s.listenByAddress && !s.fetchPoolsFromDB {
		// Get pool info from on-chain. @poolMap is empty.
		pair, err = s.GetPairByID(int64(i))
		if err != nil {
			log.Error("error fetching pair: ", err)
		}
	} else {
		// Relevant pool info is retrieved from @poolMap.
		pair = poolMap[address.Hex()]
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
	sink, err := s.GetSwapsChannel(pair.Address)
	if err != nil {
		log.Error("error fetching swaps channel: ", err)
	}

	go func() {
		for {
			rawSwap, ok := <-sink
			if ok {
				swap, err := s.normalizeUniswapSwap(*rawSwap, pair)
				if err != nil {
					log.Error("error normalizing swap: ", err)
				}
				price, volume := getSwapData(swap)
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
					Time:           time.Unix(swap.Timestamp, 0),
					PoolAddress:    rawSwap.Raw.Address.Hex(),
					ForeignTradeID: swap.ID,
					Source:         s.exchangeName,
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
				// case token0.Address == "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" && !utils.Contains(&mainBaseAssets, token1.Address):
				// 	// Reverse almost all pairs WETH-XXX ...
				// 	if s.exchangeName == dia.UniswapExchange || s.exchangeName == dia.SushiSwapExchange {
				// 		tSwapped, err := dia.SwapTrade(*t)
				// 		if err == nil {
				// 			t = &tSwapped
				// 		}
				// 	}
				// // ...and USDT-XXX on Ethereum, i.e. Uniswap and Sushiswap
				// case token0.Address == mainBaseAssets[0] && token0.Blockchain == dia.ETHEREUM:
				// 	tSwapped, err := dia.SwapTrade(*t)
				// 	if err == nil {
				// 		t = &tSwapped
				// 	}
				// // Reverse USDC-XXX pairs on Fantom
				// case token0.Address == "0x04068DA6C83AFCFA0e13ba15A6696662335D5B75" && token0.Blockchain == dia.FANTOM:
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
					s.chanTrades <- t
				}
			}
		}
	}()
}

// GetSwapsChannel returns a channel for swaps of the pair with address @pairAddress
func (s *UniswapScraper) GetSwapsChannel(pairAddress common.Address) (chan *uniswap.UniswapV2PairSwap, error) {

	sink := make(chan *uniswap.UniswapV2PairSwap)
	var pairFiltererContract *uniswap.UniswapV2PairFilterer
	pairFiltererContract, err := uniswap.NewUniswapV2PairFilterer(pairAddress, s.WsClient)
	if err != nil {
		log.Fatal(err)
	}

	_, err = pairFiltererContract.WatchSwap(&bind.WatchOpts{}, sink, []common.Address{}, []common.Address{})
	if err != nil {
		log.Error("error in get swaps channel: ", err)
	}

	return sink, nil

}

// getReverseTokensFromConfig returns a list of addresses from config file.
func getReverseTokensFromConfig(filename string) (*[]string, error) {

	var reverseTokens []string

	// Load file and read data
	filehandle := configCollectors.ConfigFileConnectors(filename, ".json")
	jsonFile, err := os.Open(filehandle)
	if err != nil {
		return &[]string{}, err
	}
	defer func() {
		err = jsonFile.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	byteData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return &[]string{}, err
	}

	// Unmarshal read data
	type lockedAsset struct {
		Address string `json:"Address"`
		Symbol  string `json:"Symbol"`
	}
	type lockedAssetList struct {
		AllAssets []lockedAsset `json:"Tokens"`
	}
	var allAssets lockedAssetList
	err = json.Unmarshal(byteData, &allAssets)
	if err != nil {
		return &[]string{}, err
	}

	// Extract addresses
	for _, token := range allAssets.AllAssets {
		reverseTokens = append(reverseTokens, token.Address)
	}

	return &reverseTokens, nil
}

// getAddressesFromConfig returns a list of Uniswap pool addresses taken from a config file.
func getAddressesFromConfig(filename string) (pairAddresses []common.Address, err error) {

	// Load file and read data
	filehandle := configCollectors.ConfigFileConnectors(filename, ".json")
	jsonFile, err := os.Open(filehandle)
	if err != nil {
		return
	}
	defer func() {
		err = jsonFile.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	byteData, err := ioutil.ReadAll(jsonFile)
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

// normalizeUniswapSwap takes a swap as returned by the swap contract's channel and converts it to a UniswapSwap type
func (s *UniswapScraper) normalizeUniswapSwap(swap uniswap.UniswapV2PairSwap, pair UniswapPair) (normalizedSwap UniswapSwap, err error) {

	decimals0 := int(pair.Token0.Decimals)
	decimals1 := int(pair.Token1.Decimals)
	amount0In, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount0In), new(big.Float).SetFloat64(math.Pow10(decimals0))).Float64()
	amount0Out, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount0Out), new(big.Float).SetFloat64(math.Pow10(decimals0))).Float64()
	amount1In, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount1In), new(big.Float).SetFloat64(math.Pow10(decimals1))).Float64()
	amount1Out, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount1Out), new(big.Float).SetFloat64(math.Pow10(decimals1))).Float64()

	normalizedSwap = UniswapSwap{
		ID:         swap.Raw.TxHash.Hex(),
		Timestamp:  time.Now().Unix(),
		Pair:       pair,
		Amount0In:  amount0In,
		Amount0Out: amount0Out,
		Amount1In:  amount1In,
		Amount1Out: amount1Out,
	}
	return
}

// pairHealthCheck returns true if the involved tokens are not blacklisted and do not have zero entries
func (up *UniswapPair) pairHealthCheck() bool {
	if up.Token0.Symbol == "" || up.Token1.Symbol == "" || up.Token0.Address.Hex() == "" || up.Token1.Address.Hex() == "" {
		return false
	}
	if helpers.SymbolIsBlackListed(up.Token0.Symbol) || helpers.SymbolIsBlackListed(up.Token1.Symbol) {
		if helpers.SymbolIsBlackListed(up.Token0.Symbol) {
			log.Infof("skip pair %s. symbol %s is blacklisted", up.ForeignName, up.Token0.Symbol)
		} else {
			log.Infof("skip pair %s. symbol %s is blacklisted", up.ForeignName, up.Token1.Symbol)
		}
		return false
	}
	if helpers.AddressIsBlacklisted(up.Token0.Address) || helpers.AddressIsBlacklisted(up.Token1.Address) {
		log.Info("skip pair ", up.ForeignName, ", address is blacklisted")
		return false
	}
	return true
}

// FetchAvailablePairs returns a list with all available trade pairs as dia.ExchangePair for the pairDiscorvery service
func (s *UniswapScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	time.Sleep(100 * time.Millisecond)
	uniPairs, err := s.GetAllPairs()
	if err != nil {
		return
	}
	for _, pair := range uniPairs {
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
			Exchange:       "UniswapV2",
			Verified:       true,
			UnderlyingPair: dia.Pair{BaseToken: basetoken, QuoteToken: quotetoken},
		}
		normalizedPair, _ := s.NormalizePair(pairToNormalise)
		pairs = append(pairs, normalizedPair)
	}

	return
}

// FillSymbolData is not used by DEX scrapers.
func (s *UniswapScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{}, nil
}

// GetAllPairs is similar to FetchAvailablePairs. But instead of dia.ExchangePairs it returns all pairs as UniswapPairs,
// i.e. including the pair's address
func (s *UniswapScraper) GetAllPairs() ([]UniswapPair, error) {
	time.Sleep(20 * time.Millisecond)
	connection := s.RestClient
	var contract *uniswap.IUniswapV2FactoryCaller
	contract, err := uniswap.NewIUniswapV2FactoryCaller(common.HexToAddress(exchangeFactoryContractAddress), connection)
	if err != nil {
		log.Error(err)
	}

	numPairs, err := contract.AllPairsLength(&bind.CallOpts{})
	if err != nil {
		return []UniswapPair{}, err
	}
	wg := sync.WaitGroup{}
	defer wg.Wait()
	pairs := make([]UniswapPair, int(numPairs.Int64()))
	for i := 0; i < int(numPairs.Int64()); i++ {
		// Sleep in order not to run into rate limits.
		time.Sleep(time.Duration(s.waitTime) * time.Millisecond)
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			uniPair, err := s.GetPairByID(int64(index))
			if err != nil {
				log.Error("error retrieving pair by ID: ", err)
				return
			}
			pairs[index] = uniPair
		}(i)
	}
	return pairs, nil
}

func (up *UniswapScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// GetPairByID returns the UniswapPair with the integer id @num
func (s *UniswapScraper) GetPairByID(num int64) (UniswapPair, error) {
	var contract *uniswap.IUniswapV2FactoryCaller
	contract, err := uniswap.NewIUniswapV2FactoryCaller(common.HexToAddress(exchangeFactoryContractAddress), s.RestClient)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}
	numToken := big.NewInt(num)
	pairAddress, err := contract.AllPairs(&bind.CallOpts{}, numToken)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}

	pair, err := s.GetPairByAddress(pairAddress)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}
	return pair, err
}

// GetPairByAddress returns the UniswapPair with pair address @pairAddress
func (s *UniswapScraper) GetPairByAddress(pairAddress common.Address) (pair UniswapPair, err error) {
	connection := s.RestClient
	var pairContract *uniswap.IUniswapV2PairCaller
	pairContract, err = uniswap.NewIUniswapV2PairCaller(pairAddress, connection)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}

	// Getting tokens from pair ---------------------
	address0, _ := pairContract.Token0(&bind.CallOpts{})
	address1, _ := pairContract.Token1(&bind.CallOpts{})
	var token0Contract *uniswap.IERC20Caller
	var token1Contract *uniswap.IERC20Caller
	token0Contract, err = uniswap.NewIERC20Caller(address0, connection)
	if err != nil {
		log.Error(err)
	}
	token1Contract, err = uniswap.NewIERC20Caller(address1, connection)
	if err != nil {
		log.Error(err)
	}
	symbol0, err := token0Contract.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}
	symbol1, err := token1Contract.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}
	decimals0, err := s.GetDecimals(address0)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}
	decimals1, err := s.GetDecimals(address1)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}

	name0, err := s.GetName(address0)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}
	name1, err := s.GetName(address1)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}
	token0 := UniswapToken{
		Address:  address0,
		Symbol:   symbol0,
		Decimals: decimals0,
		Name:     name0,
	}
	token1 := UniswapToken{
		Address:  address1,
		Symbol:   symbol1,
		Decimals: decimals1,
		Name:     name1,
	}
	foreignName := symbol0 + "-" + symbol1
	pair = UniswapPair{
		ForeignName: foreignName,
		Address:     pairAddress,
		Token0:      token0,
		Token1:      token1,
	}
	return pair, nil
}

// GetDecimals returns the decimals of the token with address @tokenAddress
func (s *UniswapScraper) GetDecimals(tokenAddress common.Address) (decimals uint8, err error) {

	var contract *uniswap.IERC20Caller
	contract, err = uniswap.NewIERC20Caller(tokenAddress, s.RestClient)
	if err != nil {
		log.Error(err)
		return
	}
	decimals, err = contract.Decimals(&bind.CallOpts{})

	return
}

func (s *UniswapScraper) GetName(tokenAddress common.Address) (name string, err error) {

	var contract *uniswap.IERC20Caller
	contract, err = uniswap.NewIERC20Caller(tokenAddress, s.RestClient)
	if err != nil {
		log.Error(err)
		return
	}
	name, err = contract.Name(&bind.CallOpts{})

	return
}

// getNumPairs returns the number of available pairs on Uniswap
func (s *UniswapScraper) getNumPairs() (int, error) {

	var contract *uniswap.IUniswapV2FactoryCaller
	contract, err := uniswap.NewIUniswapV2FactoryCaller(common.HexToAddress(exchangeFactoryContractAddress), s.RestClient)
	if err != nil {
		log.Error(err)
	}

	// Getting pairs ---------------
	numPairs, err := contract.AllPairsLength(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}
	return int(numPairs.Int64()), err
}

// getSwapData returns price, volume and sell/buy information of @swap
func getSwapData(swap UniswapSwap) (price float64, volume float64) {
	if swap.Amount0In == float64(0) {
		volume = swap.Amount0Out
		price = swap.Amount1In / swap.Amount0Out
		return
	}
	volume = -swap.Amount0In
	price = swap.Amount1Out / swap.Amount0In
	return
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *UniswapScraper) Close() error {
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
func (s *UniswapScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("UniswapScraper: Call ScrapePair on closed scraper")
	}
	ps := &UniswapPairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

// UniswapPairScraper implements PairScraper for Uniswap
type UniswapPairScraper struct {
	parent *UniswapScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *UniswapPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *UniswapScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *UniswapPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *UniswapPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

// makeUniPoolMap returns a map with pool addresses as keys and the underlying UniswapPair as values.
// If s.listenByAddress is true, it only loads the corresponding assets from the list.
func (s *UniswapScraper) makeUniPoolMap(liquiThreshold float64, liquiThresholdUSD float64) (map[string]UniswapPair, error) {
	pm := make(map[string]UniswapPair)
	var (
		pools []dia.Pool
		err   error
	)

	if s.listenByAddress {
		// Only load pool info for addresses from json file.
		poolAddresses, errAddr := getAddressesFromConfig("uniswap/subscribe_pools/" + s.exchangeName)
		if err != nil {
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
		poolsPreselection, errPreselect := s.relDB.GetAllPoolsExchange(s.exchangeName, liquiThreshold)
		if errPreselect != nil {
			return pm, errPreselect
		}
		log.Infof("Found %v pools after preselection.", len(poolsPreselection))

		for _, pool := range poolsPreselection {
			liquidity, lowerBound, errLiqui := s.datastore.GetPoolLiquidityUSD(pool)
			// Discard pool if complete USD liquidity is below threshold.
			if errLiqui == nil && !lowerBound && liquidity < liquiThresholdUSD {
				continue
			} else {
				pools = append(pools, pool)
			}
		}
		log.Infof("Found %v pools after USD liquidity filtering.", len(pools))

	} else {
		// Pool info will be fetched from on-chain and poolMap is not needed.
		return pm, nil
	}

	log.Info("Found ", len(pools), " pools.")
	log.Info("make pool map...")
	for _, pool := range pools {
		if len(pool.Assetvolumes) != 2 {
			log.Warn("not enough assets in pool with address: ", pool.Address)
			continue
		}
		up := UniswapPair{
			Address: common.HexToAddress(pool.Address),
		}
		if pool.Assetvolumes[0].Index == 0 {
			up.Token0 = asset2UniAsset(pool.Assetvolumes[0].Asset)
			up.Token1 = asset2UniAsset(pool.Assetvolumes[1].Asset)
		} else {
			up.Token0 = asset2UniAsset(pool.Assetvolumes[1].Asset)
			up.Token1 = asset2UniAsset(pool.Assetvolumes[0].Asset)
		}
		up.ForeignName = up.Token0.Symbol + "-" + up.Token1.Symbol
		pm[pool.Address] = up
	}

	log.Infof("found %v subscribable pools.", len(pm))
	return pm, err
}
