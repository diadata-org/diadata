package scrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	uniswapcontract "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/uniswap"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	exchangeFactoryContractAddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
	reversePairs                   *[]string
)

const (
	wsDial   = "ws://159.69.120.42:8546/"
	restDial = "http://159.69.120.42:8545/"

	wsDialBSC   = "wss://bsc-ws-node.nariox.org:443"
	restDialBSC = "https://bsc-dataseed2.defibit.io/"
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
func NewUniswapScraper(exchange dia.Exchange) *UniswapScraper {
	log.Info("NewUniswapScraper ", exchange.Name)
	var wsClient, restClient *ethclient.Client
	var err error

	switch exchange.Name {
	case dia.UniswapExchange:
		exchangeFactoryContractAddress = exchange.Contract.String()
		wsClient, err = ethclient.Dial(wsDial)
		if err != nil {
			log.Fatal(err)
		}

		restClient, err = ethclient.Dial(restDial)
		if err != nil {
			log.Fatal(err)
		}

		break
	case dia.SushiSwapExchange:
		exchangeFactoryContractAddress = exchange.Contract.String()
		wsClient, err = ethclient.Dial(wsDial)
		if err != nil {
			log.Fatal(err)
		}

		restClient, err = ethclient.Dial(restDial)
		if err != nil {
			log.Fatal(err)
		}

		break
	case dia.PanCakeSwap:
		log.Infoln("Init ws and rest client for BSC chain")
		wsClient, err = ethclient.Dial(wsDialBSC)
		if err != nil {
			log.Fatal(err)
		}
		restClient, err = ethclient.Dial(restDialBSC)
		if err != nil {
			log.Fatal(err)
		}
		exchangeFactoryContractAddress = exchange.Contract.String()
	}

	s := &UniswapScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*UniswapPairScraper),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
	}

	s.WsClient = wsClient
	s.RestClient = restClient

	go s.mainLoop()
	return s
}

// runs in a goroutine until s is closed
func (s *UniswapScraper) mainLoop() {

	// Import tokens which appear as base token and we need a quotation for
	var err error
	reversePairs, err = getReverseTokensFromConfig("reverse_tokens")
	if err != nil {
		log.Error("error getting tokens for which pairs should be reversed: ", err)
	}

	// wait for all pairs have added into s.PairScrapers
	time.Sleep(4 * time.Second)
	s.run = true

	numPairs, err := s.getNumPairs()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Found ", numPairs, " pairs")
	log.Info("Found ", len(s.pairScrapers), " pairScrapers")

	if len(s.pairScrapers) == 0 {
		s.error = errors.New("Uniswap: No pairs to scrap provided")
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
		if helpers.SymbolIsBlackListed(pair.Token0.Symbol) || helpers.SymbolIsBlackListed(pair.Token1.Symbol) {
			if helpers.SymbolIsBlackListed(pair.Token0.Symbol) {
				log.Infof("skip pair %s. symbol %s is blacklisted", pair.ForeignName, pair.Token0.Symbol)
			} else {
				log.Infof("skip pair %s. symbol %s is blacklisted", pair.ForeignName, pair.Token1.Symbol)
			}
			continue
		}
		if helpers.AddressIsBlacklisted(pair.Token0.Address) || helpers.AddressIsBlacklisted(pair.Token1.Address) {
			log.Info("skip pair ", pair.ForeignName, ", address is blacklisted")
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
						// If we need quotation of a base token, reverse pair
						if utils.Contains(reversePairs, strings.ToLower(pair.Token1.Address.Hex())) {
							tSwapped, err := dia.SwapTrade(*t)
							if err == nil {
								t = &tSwapped
							}
						}
						log.Info("Got trade: ", t)
						ps.parent.chanTrades <- t
					}
				}
			}()
		} else {
			log.Info("Skipping pair due to no pairScraper being available")
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

// getReverseTokensFromConfig returns a list of addresses from config file.
func getReverseTokensFromConfig(filename string) (*[]string, error) {

	var reverseTokens []string

	// Load file and read data
	fileName := fmt.Sprintf("../config/uniswap/%s.json", filename)
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return &[]string{}, err
	}
	defer jsonFile.Close()
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
	json.Unmarshal(byteData, &allAssets)

	// Extract addresses
	for _, token := range allAssets.AllAssets {
		reverseTokens = append(reverseTokens, token.Address)
	}

	return &reverseTokens, nil
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
		if pair.Token0.Symbol == "" || pair.Token1.Symbol == "" {
			continue
		}
		pairToNormalise := dia.Pair{
			Symbol:      pair.Token0.Symbol,
			ForeignName: pair.ForeignName,
			Exchange:    "UniswapV2",
			Ignore:      false,
		}
		normalizedPair, _ := s.NormalizePair(pairToNormalise)
		pairs = append(pairs, normalizedPair)
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
				return
			}
			uniPair.normalizeUniPair()
			pairs[index] = uniPair
		}(i)
	}
	return pairs, nil
}

func (up *UniswapScraper) NormalizePair(pair dia.Pair) (dia.Pair, error) {
	if pair.ForeignName == "WETH" {
		pair.Symbol = "ETH"
	}
	return pair, nil
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
	log.Info("Get pair ID: ", num)
	var contract *uniswapcontract.IUniswapV2FactoryCaller
	contract, err := uniswapcontract.NewIUniswapV2FactoryCaller(common.HexToAddress(exchangeFactoryContractAddress), s.RestClient)
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
