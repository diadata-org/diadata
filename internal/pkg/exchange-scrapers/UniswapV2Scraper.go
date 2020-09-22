package scrapers

import (
	"errors"
	"math"
	"math/big"
	"sync"
	"time"

	uniswapcontract "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/uniswap"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	exchangeFactoryContractAddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
)

const (
	wsDial   = "ws://159.69.120.42:8546/"
	restDial = "http://159.69.120.42:8545/"
)

type UniswapToken struct {
	Address  common.Address
	Symbol   string
	Decimals uint8
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
}

// NewUniswapScraper returns a new UniswapScraper for the given pair
func NewUniswapScraper(exchangeName string) *UniswapScraper {

	switch exchangeName {
	case dia.UniswapExchange:
		exchangeFactoryContractAddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"

		break
	case dia.SushiSwapExchange:
		exchangeFactoryContractAddress = "0xc0aee478e3658e2610c5f7a4a2e1777ce9e4f2ac"
	}

	s := &UniswapScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*UniswapPairScraper),
		exchangeName: exchangeName,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
	}

	wsClient, err := ethclient.Dial(wsDial)
	if err != nil {
		log.Fatal(err)
	}
	s.WsClient = wsClient
	restClient, err := ethclient.Dial(restDial)
	if err != nil {
		log.Fatal(err)
	}
	s.RestClient = restClient

	go s.mainLoop()
	return s
}

// runs in a goroutine until s is closed
func (s *UniswapScraper) mainLoop() {

	// wait for all pairs have added into s.PairScrapers
	time.Sleep(4 * time.Second)
	s.run = true

	numPairs, err := s.getNumPairs()
	if err != nil {
		log.Fatal(err)
	}

	if len(s.pairScrapers) == 0 {
		s.error = errors.New("bancor: No pairs to scrap provided")
		log.Error(s.error.Error())
	}
	for i := 0; i < numPairs; i++ {

		pair, err := s.GetPairByID(int64(i))
		if err != nil {
			log.Error("error fetching pair: ", err)
		}
		if len(pair.Token0.Symbol) < 2 || len(pair.Token1.Symbol) < 2 {
			log.Info("skip pair: ", pair.ForeignName)
			continue
		}
		pair.normalizeUniPair()
		ps, ok := s.pairScrapers[pair.ForeignName]
		if ok {
			log.Info(i, ": found pair scraper for: ", pair.ForeignName, " with address ", pair.Address.Hex())
			sink, err := s.GetSwapsChannel(pair.Address)
			if err != nil {
				log.Error("error fetching swaps channel: ", err)
			}

			go func() {
				for {
					rawSwap, ok := <-sink
					if ok {
						swap, err := s.normalizeUniswapSwap(*rawSwap)
						if err != nil {
							log.Error("error normalizing swap: ", err)
						}
						price, volume, err := getSwapData(swap)
						if err != nil {
							log.Error("error getting swap data: ", err)
						}

						t := &dia.Trade{
							Symbol:         ps.pair.Symbol,
							Pair:           ps.pair.ForeignName,
							Price:          price,
							Volume:         volume,
							Time:           time.Unix(swap.Timestamp, 0),
							ForeignTradeID: swap.ID,
							Source:         s.exchangeName,
						}
						ps.parent.chanTrades <- t
						log.Info("Got trade: ", t)
					}
				}
			}()
		}

	}

	// s.cleanup(err)
}

// GetSwapsChannel returns a channel for swaps of the pair with address @pairAddress
func (s *UniswapScraper) GetSwapsChannel(pairAddress common.Address) (chan *uniswapcontract.UniswapV2PairSwap, error) {

	sink := make(chan *uniswapcontract.UniswapV2PairSwap)
	var pairFiltererContract *uniswapcontract.UniswapV2PairFilterer
	pairFiltererContract, err := uniswapcontract.NewUniswapV2PairFilterer(pairAddress, s.WsClient)
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
func (s *UniswapScraper) normalizeUniswapSwap(swap uniswapcontract.UniswapV2PairSwap) (normalizedSwap UniswapSwap, err error) {

	pair, err := s.GetPairByAddress(swap.Raw.Address)
	if err != nil {
		log.Error("error getting pair by address: ", err)
		return
	}
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

// FetchAvailablePairs returns a list with all available trade pairs as dia.Pair for the pairDiscorvery service
func (s *UniswapScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {

	uniPairs, err := s.GetAllPairs()
	if err != nil {
		return
	}
	for _, pair := range uniPairs {
		pairs = append(pairs, dia.Pair{
			Symbol:      pair.Token0.Symbol,
			ForeignName: pair.ForeignName,
			Exchange:    "UniswapV2",
			Ignore:      false,
		})
	}

	return
}

// GetAllPairs is similar to FetchAvailablePairs. But instead of dia.Pairs it returns all pairs as UniswapPairs,
// i.e. including the pair's address
func (s *UniswapScraper) GetAllPairs() ([]UniswapPair, error) {
	connection := s.RestClient
	var contract *uniswapcontract.IUniswapV2FactoryCaller
	contract, err := uniswapcontract.NewIUniswapV2FactoryCaller(common.HexToAddress(exchangeFactoryContractAddress), connection)
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
			}
			uniPair.normalizeUniPair()
			pairs[index] = uniPair
		}(i)
	}
	return pairs, nil
}

// Account for WETH is identified with ETH
func (up *UniswapPair) normalizeUniPair() {
	if up.Token0.Symbol == "WETH" {
		up.Token0.Symbol = "ETH"
		up.ForeignName = up.Token0.Symbol + "-" + up.Token1.Symbol
	}
	if up.Token1.Symbol == "WETH" {
		up.Token1.Symbol = "ETH"
		up.ForeignName = up.Token0.Symbol + "-" + up.Token1.Symbol
	}
}

// GetPairByID returns the UniswapPair with the integer id @num
func (s *UniswapScraper) GetPairByID(num int64) (UniswapPair, error) {
	var contract *uniswapcontract.IUniswapV2FactoryCaller
	contract, err := uniswapcontract.NewIUniswapV2FactoryCaller(common.HexToAddress(exchangeFactoryContractAddress), s.RestClient)
	if err != nil {
		log.Error(err)
	}
	numToken := big.NewInt(num)
	pairAddress, err := contract.AllPairs(&bind.CallOpts{}, numToken)

	return s.GetPairByAddress(pairAddress)
}

// GetPairByAddress returns the UniswapPair with pair address @pairAddress
func (s *UniswapScraper) GetPairByAddress(pairAddress common.Address) (pair UniswapPair, err error) {
	connection := s.RestClient
	var pairContract *uniswapcontract.IUniswapV2PairCaller
	pairContract, err = uniswapcontract.NewIUniswapV2PairCaller(pairAddress, connection)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}

	// Getting tokens from pair ---------------------
	address0, _ := pairContract.Token0(&bind.CallOpts{})
	address1, _ := pairContract.Token1(&bind.CallOpts{})
	var token0Contract *uniswapcontract.IERC20Caller
	var token1Contract *uniswapcontract.IERC20Caller
	token0Contract, err = uniswapcontract.NewIERC20Caller(address0, connection)
	if err != nil {
		log.Error(err)
	}
	token1Contract, err = uniswapcontract.NewIERC20Caller(address1, connection)
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
	token0 := UniswapToken{
		Address:  address0,
		Symbol:   symbol0,
		Decimals: decimals0,
	}
	token1 := UniswapToken{
		Address:  address1,
		Symbol:   symbol1,
		Decimals: decimals1,
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

	var contract *uniswapcontract.IERC20Caller
	contract, err = uniswapcontract.NewIERC20Caller(tokenAddress, s.RestClient)
	if err != nil {
		log.Error(err)
		return
	}
	decimals, err = contract.Decimals(&bind.CallOpts{})

	return
}

// getNumPairs returns the number of available pairs on Uniswap
func (s *UniswapScraper) getNumPairs() (int, error) {

	var contract *uniswapcontract.IUniswapV2FactoryCaller
	contract, err := uniswapcontract.NewIUniswapV2FactoryCaller(common.HexToAddress(exchangeFactoryContractAddress), s.RestClient)
	if err != nil {
		log.Error(err)
	}

	// Getting pairs ---------------
	numPairs, err := contract.AllPairsLength(&bind.CallOpts{})
	return int(numPairs.Int64()), err
}

// getSwapData returns price, volume and sell/buy information of @swap
func getSwapData(swap UniswapSwap) (price float64, volume float64, err error) {
	if swap.Amount0In == float64(0) {
		volume = swap.Amount0Out
		price = swap.Amount1In / swap.Amount0Out
		return
	}

	volume = -swap.Amount0In
	price = swap.Amount1Out / swap.Amount0In
	return
}

func (s *UniswapScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone)
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
func (s *UniswapScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {

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
	pair   dia.Pair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *UniswapPairScraper) Close() error {
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
func (ps *UniswapPairScraper) Pair() dia.Pair {
	return ps.pair
}
