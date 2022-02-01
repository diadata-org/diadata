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

	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswap"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UniswapHistoryScraper struct {
	WsClient   *ethclient.Client
	RestClient *ethclient.Client
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
	pairScrapers  map[string]*UniswapHistoryPairScraper
	exchangeName  string
	chanTrades    chan *dia.Trade
	waitTime      int
	genesisBlock  uint64
	pairmap       map[common.Address]UniswapPair
	pairAddresses []common.Address
	db            *models.RelDB
	// If true, only pairs given in config file are scraped. Default is false.
	listenByAddress bool
}

const (
	// genesisBlockUniswap            = uint64(10019990)
	// genesisBlockUniswap            = uint64(10520000)
	genesisBlockUniswap            = uint64(12120000)
	filterQueryBlockNums           = 60
	uniswapHistoryWaitMilliseconds = "500"
)

// NewUniswapScraper returns a new UniswapScraper for the given pair
func NewUniswapHistoryScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *UniswapHistoryScraper {
	log.Info("NewUniswapHistoryScraper: ", exchange.Name)
	var s *UniswapHistoryScraper
	var listenByAddress bool
	exchangeFactoryContractAddress = exchange.Contract.Hex()

	switch exchange.Name {
	case dia.UniswapExchange:
		listenByAddress = false
		s = makeUniswapHistoryScraper(exchange, listenByAddress, restDialEth, wsDialEth, uniswapHistoryWaitMilliseconds)
	case dia.SushiSwapExchange:
		listenByAddress = false
		s = makeUniswapHistoryScraper(exchange, listenByAddress, restDialEth, wsDialEth, sushiswapWaitMilliseconds)
	case dia.PanCakeSwap:
		listenByAddress = true
		s = makeUniswapHistoryScraper(exchange, listenByAddress, restDialBSC, wsDialBSC, pancakeswapWaitMilliseconds)
	case dia.DfynNetwork:
		listenByAddress = false
		s = makeUniswapHistoryScraper(exchange, listenByAddress, restDialPolygon, wsDialPolygon, dfynWaitMilliseconds)
	}

	if scrape {
		go s.mainLoop()
	}
	return s
}

// makeUniswapScraper returns a uniswap scraper as used in NewUniswapScraper.
func makeUniswapHistoryScraper(exchange dia.Exchange, listenByAddress bool, restDial string, wsDial string, waitMilliseconds string) *UniswapHistoryScraper {
	var restClient, wsClient *ethclient.Client
	var err error
	var s *UniswapHistoryScraper

	log.Infof("Init rest and ws client for %s.", exchange.BlockChain.Name)
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

	s = &UniswapHistoryScraper{
		WsClient:        wsClient,
		RestClient:      restClient,
		shutdown:        make(chan nothing),
		shutdownDone:    make(chan nothing),
		pairScrapers:    make(map[string]*UniswapHistoryPairScraper),
		exchangeName:    exchange.Name,
		error:           nil,
		chanTrades:      make(chan *dia.Trade),
		waitTime:        waitTime,
		listenByAddress: listenByAddress,
		genesisBlock:    genesisBlockUniswap,
	}
	return s
}

func (s *UniswapHistoryScraper) loadPairMap() {
	var maps []map[common.Address]UniswapPair

	if s.listenByAddress {

		// Collect all pair addresses from json file.
		pairAddresses, err := getAddressesFromConfig("uniswap/subscribe_pools/" + s.exchangeName)
		if err != nil {
			log.Error("fetch pool addresses from config file: ", err)
		}
		numPairs := len(pairAddresses)
		log.Infof("listening to %d pools: %v", numPairs, pairAddresses)
		for _, pairAddress := range pairAddresses {
			uniPair, err := s.GetPairByAddress(pairAddress)
			if err != nil {
				log.Errorf("get pair with address %s: %v", pairAddress.Hex(), err)
			}
			auxmap := make(map[common.Address]UniswapPair)
			auxmap[pairAddress] = uniPair
			maps = append(maps, auxmap)
		}

	} else {
		numPairs, err := s.getNumPairs()
		if err != nil {
			log.Fatal(err)
		}
		// numPairs := 1
		batchSize := 1000

		log.Infof("load all %d pairs: ", numPairs)

		var wg sync.WaitGroup
		for k := 0; k < numPairs/batchSize; k++ {
			time.Sleep(8 * time.Second)
			for i := batchSize * k; i < batchSize*(k+1); i++ {
				wg.Add(1)
				auxmap := make(map[common.Address]UniswapPair)
				go func(index int, w *sync.WaitGroup) {
					defer w.Done()
					pair, err := s.GetPairByID(int64(index))
					if err != nil {
						log.Error(err)
					}
					auxmap[pair.Address] = pair
				}(i, &wg)
				maps = append(maps, auxmap)
			}
			wg.Wait()
		}
		for i := numPairs - numPairs%batchSize; i < numPairs; i++ {
			wg.Add(1)
			auxmap := make(map[common.Address]UniswapPair)
			go func(index int, w *sync.WaitGroup) {
				defer w.Done()
				pair, err := s.GetPairByID(int64(index))
				if err != nil {
					log.Error(err)
				}
				auxmap[pair.Address] = pair
			}(i, &wg)
			maps = append(maps, auxmap)
		}
		wg.Wait()

		log.Info("len: ", len(maps))
	}

	pairmap := make(map[common.Address]UniswapPair)
	for _, m := range maps {
		for i, j := range m {
			j.normalizeUniPair()
			pairmap[i] = j
		}
	}

	log.Info("len: ", len(pairmap))
	s.pairmap = pairmap
}

// runs in a goroutine until s is closed
func (s *UniswapHistoryScraper) mainLoop() {

	// Import tokens which appear as base token and we need a quotation for
	var err error
	reversePairs, err = getReverseTokensFromConfig("uniswap/reverse_tokens")
	if err != nil {
		log.Error("error getting tokens for which pairs should be reversed: ", err)
	}

	// wait for all pairs have added into s.PairScrapers
	time.Sleep(4 * time.Second)
	s.run = true

	// load all pairs into and from pair map.
	s.loadPairMap()
	var addresses []common.Address
	for k := range s.pairmap {
		addresses = append(addresses, k)
	}
	s.pairAddresses = addresses

	if len(addresses) == 0 {
		s.error = errors.New("uniswap: No pairs to scrape provided")
		log.Error(s.error.Error())
	} else {
		log.Infof("%d pairs loaded.", len(addresses))
	}

	latestBlock, err := s.RestClient.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Error("get current block number: ", err)
	}

	startblock := s.genesisBlock
	endblock := startblock + uint64(filterQueryBlockNums)

	for startblock < latestBlock.NumberU64() {
		err := s.fetchSwaps(startblock, endblock)
		if err != nil {
			if strings.Contains(err.Error(), "EOF") {
				endblock = startblock + (endblock-startblock)/2
				time.Sleep(2 * time.Second)
				continue
			}
			log.Error("get filter logs: ", err)
		}
		startblock = endblock
		endblock = startblock + filterQueryBlockNums
	}

	// ---------------------------------------------------------------------------
	// Concurrent block scraping
	// ---------------------------------------------------------------------------

	// a := s.genesisBlock
	// b := latestBlock.NumberU64()
	// numSubblocks := 1
	// startblocks := []uint64{}
	// finalBlocks := []uint64{}
	// for k := 0; k < numSubblocks; k++ {
	// 	startblocks = append(startblocks, a+uint64(k)*(b-a)/uint64(numSubblocks))
	// 	finalBlocks = append(finalBlocks, a+uint64(k+1)*(b-a)/uint64(numSubblocks))
	// }

	// var wg sync.WaitGroup
	// for i := 0; i < len(startblocks); i++ {
	// 	startblock := startblocks[i]
	// 	endblock := startblocks[i] + uint64(filterQueryBlockNums)

	// 	wg.Add(1)
	// 	go func(startblock, endblock uint64, index int, w *sync.WaitGroup) {
	// 		defer w.Done()
	// 		for startblock < finalBlocks[index] {
	// 			err := s.fetchSwaps(startblock, endblock)
	// 			if err != nil {
	// 				if strings.Contains(err.Error(), "EOF") {
	// 					endblock = startblock + (endblock-startblock)/2
	// 					time.Sleep(2 * time.Second)
	// 					continue
	// 				}
	// 				log.Error("get filter logs: ", err)
	// 			}
	// 			startblock = endblock
	// 			endblock = startblock + filterQueryBlockNums
	// 		}
	// 	}(startblock, endblock, i, &wg)
	// }
	// wg.Wait()
}

func (s *UniswapHistoryScraper) fetchSwaps(startblock uint64, endblock uint64) error {
	log.Infof("get swaps from block %d to block %d.", startblock, endblock)
	hashSwap := common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822")
	topics := make([][]common.Hash, 1)
	topics[0] = append(topics[0], hashSwap)

	config := ethereum.FilterQuery{
		Addresses: s.pairAddresses,
		Topics:    topics,
		FromBlock: new(big.Int).SetUint64(startblock),
		ToBlock:   new(big.Int).SetUint64(endblock),
	}

	t := time.Now()
	logs, err := s.RestClient.FilterLogs(context.Background(), config)
	if err != nil {
		return err
	}
	log.Info("time passed for filter logs: ", time.Since(t))

	for _, logg := range logs {

		pairFilterer, err := uniswap.NewUniswapV2PairFilterer(common.Address{}, s.RestClient)
		if err != nil {
			log.Error(err)
		}

		// blockdata, err := s.RestClient.BlockByNumber(context.Background(), big.NewInt(int64(logg.BlockNumber)))
		// if err != nil {
		// 	log.Info("get block by number: ", err)
		// }

		blockdata, err := ethhelper.GetBlockData(int64(logg.BlockNumber), s.db, s.RestClient)
		if err != nil {
			return err
		}

		swap, err := pairFilterer.ParseSwap(logg)
		if err != nil {
			log.Error(err)
		}
		swp := s.normalizeUniswapSwapHistory(*swap, logg.Address)

		price, volume := getSwapData(swp)
		token0 := dia.Asset{
			Address:    swp.Pair.Token0.Address.Hex(),
			Symbol:     swp.Pair.Token0.Symbol,
			Name:       swp.Pair.Token0.Name,
			Decimals:   swp.Pair.Token0.Decimals,
			Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
		}
		token1 := dia.Asset{
			Address:    swp.Pair.Token1.Address.Hex(),
			Symbol:     swp.Pair.Token1.Symbol,
			Name:       swp.Pair.Token1.Name,
			Decimals:   swp.Pair.Token1.Decimals,
			Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
		}
		var timestamp time.Time
		switch blockdata.Data["Time"].(type) {
		case float64:
			timestamp = time.Unix(int64(blockdata.Data["Time"].(float64)), 0)
		case uint64:
			timestamp = time.Unix(int64(blockdata.Data["Time"].(uint64)), 0)
		}
		t := &dia.Trade{
			Symbol:     swp.Pair.Token0.Symbol,
			Pair:       swp.Pair.ForeignName,
			Price:      price,
			Volume:     volume,
			BaseToken:  token1,
			QuoteToken: token0,
			Time:       timestamp,
			// Time: time.Now(),
			// Time:           time.Unix(int64(blockdata.Time()), 0),
			ForeignTradeID: logg.TxHash.Hex(),
			Source:         s.exchangeName,
			VerifiedPair:   true,
		}
		// If we need quotation of a base token, reverse pair
		if utils.Contains(reversePairs, swp.Pair.Token1.Address.Hex()) {
			tSwapped, err := dia.SwapTrade(*t)
			if err == nil {
				t = &tSwapped
			}
		}
		if price > 0 {
			log.Infof("Got trade at time %v - symbol: %s, pair: %s, price: %v, volume:%v", t.Time, t.Symbol, t.Pair, t.Price, t.Volume)
			s.chanTrades <- t
		}
		if price == 0 {
			log.Info("Got zero trade: ", t)
		}

	}

	log.Info("number of swaps: ", len(logs))
	return nil

}

// normalizeUniswapSwap takes a swap as returned by the swap contract's channel and converts it to a UniswapSwap type
func (s *UniswapHistoryScraper) normalizeUniswapSwapHistory(swap uniswap.UniswapV2PairSwap, pairAddress common.Address) (normalizedSwap UniswapSwap) {

	pair := s.pairmap[pairAddress]

	decimals0 := int(pair.Token0.Decimals)
	decimals1 := int(pair.Token1.Decimals)
	amount0In, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount0In), new(big.Float).SetFloat64(math.Pow10(decimals0))).Float64()
	amount0Out, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount0Out), new(big.Float).SetFloat64(math.Pow10(decimals0))).Float64()
	amount1In, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount1In), new(big.Float).SetFloat64(math.Pow10(decimals1))).Float64()
	amount1Out, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount1Out), new(big.Float).SetFloat64(math.Pow10(decimals1))).Float64()

	normalizedSwap = UniswapSwap{
		ID:         swap.Raw.TxHash.Hex(),
		Pair:       pair,
		Amount0In:  amount0In,
		Amount0Out: amount0Out,
		Amount1In:  amount1In,
		Amount1Out: amount1Out,
	}
	return
}

// FetchAvailablePairs returns a list with all available trade pairs as dia.ExchangePair for the pairDiscorvery service
func (s *UniswapHistoryScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
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
			Blockchain: dia.ETHEREUM,
		}
		basetoken := dia.Asset{
			Symbol:     pair.Token1.Symbol,
			Name:       pair.Token1.Name,
			Address:    pair.Token1.Address.Hex(),
			Decimals:   pair.Token1.Decimals,
			Blockchain: dia.ETHEREUM,
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
func (s *UniswapHistoryScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{}, nil
}

// GetAllPairs is similar to FetchAvailablePairs. But instead of dia.ExchangePairs it returns all pairs as UniswapPairs,
// i.e. including the pair's address
func (s *UniswapHistoryScraper) GetAllPairs() ([]UniswapPair, error) {
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
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			uniPair, err := s.GetPairByID(int64(index))
			if err != nil {
				log.Error("error retrieving pair by ID: ", err)
				return
			}
			uniPair.normalizeUniPair()
			pairs[index] = uniPair
		}(i)
	}
	return pairs, nil
}

func (up *UniswapHistoryScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// GetPairByID returns the UniswapPair with the integer id @num
func (s *UniswapHistoryScraper) GetPairByID(num int64) (UniswapPair, error) {
	log.Info("Get pair ID: ", num)
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
func (s *UniswapHistoryScraper) GetPairByAddress(pairAddress common.Address) (pair UniswapPair, err error) {
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
func (s *UniswapHistoryScraper) GetDecimals(tokenAddress common.Address) (decimals uint8, err error) {

	var contract *uniswap.IERC20Caller
	contract, err = uniswap.NewIERC20Caller(tokenAddress, s.RestClient)
	if err != nil {
		log.Error(err)
		return
	}
	decimals, err = contract.Decimals(&bind.CallOpts{})

	return
}

func (s *UniswapHistoryScraper) GetName(tokenAddress common.Address) (name string, err error) {

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
func (s *UniswapHistoryScraper) getNumPairs() (int, error) {

	var contract *uniswap.IUniswapV2FactoryCaller
	contract, err := uniswap.NewIUniswapV2FactoryCaller(common.HexToAddress(exchangeFactoryContractAddress), s.RestClient)
	if err != nil {
		log.Error(err)
	}

	// Getting pairs ---------------
	numPairs, err := contract.AllPairsLength(&bind.CallOpts{})
	return int(numPairs.Int64()), err
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *UniswapHistoryScraper) Close() error {
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
func (s *UniswapHistoryScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("UniswapScraper: Call ScrapePair on closed scraper")
	}
	ps := &UniswapHistoryPairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

// UniswapPairScraper implements PairScraper for Uniswap
type UniswapHistoryPairScraper struct {
	parent *UniswapHistoryScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *UniswapHistoryPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *UniswapHistoryScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *UniswapHistoryPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *UniswapHistoryPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
