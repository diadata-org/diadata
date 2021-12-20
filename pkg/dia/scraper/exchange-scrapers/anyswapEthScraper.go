package scrapers

import (
	"errors"
	"math"
	"math/big"
	"strconv"
	"sync"
	"time"

	anyswap "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/anyswap"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswap"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	anyswapEthereumContractAddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
)

const (
	anyswapWaitMilliseconds = "200"
)

type AnyswapToken struct {
	Address  common.Address
	Symbol   string
	Decimals uint8
	Name     string
}

type AnyswapPair struct {
	Token0      UniswapToken
	Token1      UniswapToken
	ForeignName string
	Address     common.Address
}

type AnyswapSwap struct {
	ID         string
	Timestamp  int64
	Pair       UniswapPair
	Amount0In  float64
	Amount0Out float64
	Amount1In  float64
	Amount1Out float64
}

type AnyswapScraper struct {
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
	pairScrapers map[string]*AnyswapPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
	waitTime     int
	// If true, only pairs given in config file are scraped. Default is false.
	listenByAddress bool
}

// NewUniswapScraper returns a new UniswapScraper for the given pair
func NewAnyswapScraper(exchange dia.Exchange, scrape bool) *AnyswapScraper {
	log.Info("NewUniswapScraper: ", exchange.Name)
	var wsClient, restClient *ethclient.Client
	var waitTime int
	var listenByAddress bool
	var err error

	switch exchange.Name {
	case dia.AnyswapExchange:
		exchangeFactoryContractAddress = exchange.Contract.Hex()
		restClient, err = ethclient.Dial(utils.Getenv("ETH_URI_REST", restDial))
		if err != nil {
			log.Fatal(err)
		}

		wsClient, err = ethclient.Dial(utils.Getenv("ETH_URI_WS", wsDial))
		if err != nil {
			log.Fatal(err)
		}
		waitTimeString := utils.Getenv("UNISWAP_WAIT_TIME", anyswapWaitMilliseconds)
		waitTime, err = strconv.Atoi(waitTimeString)
		if err != nil {
			log.Error("could not parse wait time: ", err)
			waitTime = 100
		}

	}

	s := &AnyswapScraper{
		shutdown:        make(chan nothing),
		shutdownDone:    make(chan nothing),
		pairScrapers:    make(map[string]*AnyswapPairScraper),
		exchangeName:    exchange.Name,
		error:           nil,
		chanTrades:      make(chan *dia.Trade),
		waitTime:        waitTime,
		listenByAddress: listenByAddress,
	}

	s.WsClient = wsClient
	s.RestClient = restClient
	if scrape {
		go s.mainLoop()
	}
	return s
}

// runs in a goroutine until s is closed
func (s *AnyswapScraper) mainLoop() {

	// wait for all pairs have added into s.PairScrapers
	time.Sleep(4 * time.Second)
	s.run = true

	if s.listenByAddress {

		// Collect all pair addresses from json file.
		pairAddresses, err := getAddressesFromConfig("uniswap/subscribe_pools/" + s.exchangeName)
		if err != nil {
			log.Error("fetch pool addresses from config file: ", err)
		}
		numPairs := len(pairAddresses)
		log.Infof("listening to %d pools: %v", numPairs, pairAddresses)

		var wg sync.WaitGroup
		for i := 0; i < numPairs; i++ {
			time.Sleep(time.Duration(s.waitTime) * time.Millisecond)
			wg.Add(1)
			go func(index int, address common.Address, w *sync.WaitGroup) {
				defer w.Done()
				s.ListenToPair(index, address, s.listenByAddress)
			}(i, pairAddresses[i], &wg)
		}
		wg.Wait()

	} else {

		if len(s.pairScrapers) == 0 {
			s.error = errors.New("uniswap: No pairs to scrap provided")
			log.Error(s.error.Error())
		}

		var wg sync.WaitGroup
		for i := 0; i < 1; i++ {
			time.Sleep(time.Duration(s.waitTime) * time.Millisecond)
			wg.Add(1)
			go func(index int, address common.Address, w *sync.WaitGroup) {
				defer w.Done()
				s.ListenToPair(index, address, s.listenByAddress)
			}(i, common.Address{}, &wg)
		}
		wg.Wait()

	}
}

// ListenToPair subscribes to a uniswap pool.
// If @byAddress is true, it listens by pool address, otherwise by index.
func (s *AnyswapScraper) ListenToPair(i int, address common.Address, byAddress bool) {
	var pair UniswapPair
	var err error

	if !byAddress {
		pair, err = s.GetPairByID(int64(i))
		if err != nil {
			log.Error("error fetching pair: ", err)
		}
	} else {
		pair, err = s.GetPairByAddress(address)
		if err != nil {
			log.Error("error fetching pair: ", err)
		}
	}

	// Normalize WETH on Ethereum
	if Exchanges[s.exchangeName].BlockChain.Name == dia.ETHEREUM {
		pair.normalizeUniPair()
	}
	// ps := s.pairScrapers[pair.ForeignName]
	// if ok {
	log.Info(i, ": add pair scraper for: ", pair.ForeignName, " with address ", pair.Address.Hex())
	sink, err := s.GetSwapOutChannel([]common.Address{
		common.HexToAddress("0x7ea2be2df7ba6e54b1a9c70676f668455e329d29"),
		common.HexToAddress("0xbbc4a8d076f4b1888fec42581b6fc58d242cf2d5"),
		common.HexToAddress("0x739ca6d71365a08f584c8fc4e1029045fa8abc4b"),
		common.HexToAddress("0xb153fb3d196a8eb25522705560ac152eeec57901"),
		common.HexToAddress("0x22648c12acd87912ea1710357b1302c6a4154ebc"),
		common.HexToAddress("0xdb25f211ab05b1c97d595516f45794528a807ad8"),
		common.HexToAddress("0x0bd19f6447cd676255c7c7b00428462b3da67e3a"),
	})
	if err != nil {
		log.Error("error fetching swaps channel: ", err)
	}

	go func() {
		for {
			rawSwap, ok := <-sink
			if ok {
				log.Info("token: ", rawSwap.Token.Hex())
				log.Info("amount: ", rawSwap.Amount)
				log.Info("to chain ID: ", rawSwap.ToChainID)
			}
		}
	}()
}

func (s *AnyswapScraper) GetSwapOutChannel(tokens []common.Address) (chan *anyswap.AnyswapV4RouterLogAnySwapOut, error) {
	sink := make(chan *anyswap.AnyswapV4RouterLogAnySwapOut)
	outFiltererContract, err := anyswap.NewAnyswapV4RouterFilterer(Exchanges[s.exchangeName].Contract, s.WsClient)
	if err != nil {
		log.Fatal(err)
	}
	_, err = outFiltererContract.WatchLogAnySwapOut(&bind.WatchOpts{}, sink, tokens, []common.Address{}, []common.Address{})
	if err != nil {
		return sink, err
	}
	return sink, nil
}

// GetSwapsChannel returns a channel for swaps of the pair with address @pairAddress
func (s *AnyswapScraper) GetSwapsChannel(pairAddress common.Address) (chan *uniswap.UniswapV2PairSwap, error) {

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

// normalizeUniswapSwap takes a swap as returned by the swap contract's channel and converts it to a UniswapSwap type
func (s *AnyswapScraper) normalizeUniswapSwap(swap uniswap.UniswapV2PairSwap, pair UniswapPair) (normalizedSwap UniswapSwap, err error) {

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

// FetchAvailablePairs returns a list with all available trade pairs as dia.ExchangePair for the pairDiscorvery service
func (s *AnyswapScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
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
func (s *AnyswapScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{}, nil
}

// GetAllPairs is similar to FetchAvailablePairs. But instead of dia.ExchangePairs it returns all pairs as UniswapPairs,
// i.e. including the pair's address
func (s *AnyswapScraper) GetAllPairs() ([]UniswapPair, error) {
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
		if s.exchangeName == dia.PanCakeSwap {
			time.Sleep(time.Duration(s.waitTime) * time.Millisecond)
		}
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

func (up *AnyswapScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// GetPairByID returns the UniswapPair with the integer id @num
func (s *AnyswapScraper) GetPairByID(num int64) (UniswapPair, error) {
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
func (s *AnyswapScraper) GetPairByAddress(pairAddress common.Address) (pair UniswapPair, err error) {
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
func (s *AnyswapScraper) GetDecimals(tokenAddress common.Address) (decimals uint8, err error) {

	var contract *uniswap.IERC20Caller
	contract, err = uniswap.NewIERC20Caller(tokenAddress, s.RestClient)
	if err != nil {
		log.Error(err)
		return
	}
	decimals, err = contract.Decimals(&bind.CallOpts{})

	return
}

func (s *AnyswapScraper) GetName(tokenAddress common.Address) (name string, err error) {

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
func (s *AnyswapScraper) getNumPairs() (int, error) {

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

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *AnyswapScraper) Close() error {
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
func (s *AnyswapScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("UniswapScraper: Call ScrapePair on closed scraper")
	}
	ps := &AnyswapPairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

// UniswapPairScraper implements PairScraper for Uniswap
type AnyswapPairScraper struct {
	parent *AnyswapScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *AnyswapPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *AnyswapScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *AnyswapPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *AnyswapPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
