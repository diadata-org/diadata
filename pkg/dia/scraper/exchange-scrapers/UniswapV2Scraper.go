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
)

const (
	restDialEth = "http://159.69.120.42:8545/"
	wsDialEth   = "ws://159.69.120.42:8546/"

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
	listenByAddress bool
}

// NewUniswapScraper returns a new UniswapScraper for the given pair
func NewUniswapScraper(exchange dia.Exchange, scrape bool) *UniswapScraper {
	log.Info("NewUniswapScraper: ", exchange.Name)
	var s *UniswapScraper
	var listenByAddress bool
	exchangeFactoryContractAddress = exchange.Contract.Hex()

	switch exchange.Name {
	case dia.UniswapExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialEth, wsDialEth, uniswapWaitMilliseconds)
	case dia.SushiSwapExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialEth, wsDialEth, sushiswapWaitMilliseconds)
	case dia.PanCakeSwap:
		listenByAddress = true
		s = makeUniswapScraper(exchange, listenByAddress, restDialBSC, wsDialBSC, pancakeswapWaitMilliseconds)
	case dia.DfynNetwork:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialPolygon, wsDialPolygon, dfynWaitMilliseconds)
	case dia.QuickswapExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialPolygon, wsDialPolygon, quickswapWaitMilliseconds)
	case dia.UbeswapExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialCelo, wsDialCelo, ubeswapWaitMilliseconds)
	case dia.SpookyswapExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialFantom, wsDialFantom, spookyswapWaitMilliseconds)
	case dia.SpiritswapExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialFantom, wsDialFantom, spookyswapWaitMilliseconds)
	case dia.SolarbeamExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialMoonriver, wsDialMoonriver, solarbeamWaitMilliseconds)
	case dia.TrisolarisExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialAurora, wsDialAurora, trisolarisWaitMilliseconds)
	case dia.NetswapExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialMetis, wsDialMetis, metisWaitMilliseconds)
	case dia.SushiSwapExchangePolygon:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialPolygon, wsDialPolygon, metisWaitMilliseconds)
	case dia.SushiSwapExchangeFantom:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialFantom, wsDialFantom, metisWaitMilliseconds)
	case dia.HuckleberryExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialMoonriver, wsDialMoonriver, moonriverWaitMilliseconds)
	case dia.TraderJoeExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialAvalanche, wsDialAvalanche, avalancheWaitMilliseconds)
	case dia.PangolinExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialAvalanche, wsDialAvalanche, avalancheWaitMilliseconds)
	case dia.TethysExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialMetis, wsDialMetis, metisWaitMilliseconds)
	case dia.HermesExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialMetis, wsDialMetis, metisWaitMilliseconds)
	case dia.OmniDexExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialTelos, wsDialTelos, telosWaitMilliseconds)
	case dia.DiffusionExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialEvmos, wsDialEvmos, evmosWaitMilliseconds)
	case dia.ApeswapExchange:
		listenByAddress = true
		s = makeUniswapScraper(exchange, listenByAddress, restDialBSC, wsDialBSC, pancakeswapWaitMilliseconds)
	case dia.BiswapExchange:
		listenByAddress = true
		s = makeUniswapScraper(exchange, listenByAddress, restDialBSC, wsDialBSC, pancakeswapWaitMilliseconds)
	case dia.ArthswapExchange:
		listenByAddress = false
		s = makeUniswapScraper(exchange, listenByAddress, restDialAstar, wsDialAstar, astarWaitMilliseconds)
	}

	if scrape {
		go s.mainLoop()
	}
	return s
}

// makeUniswapScraper returns a uniswap scraper as used in NewUniswapScraper.
func makeUniswapScraper(exchange dia.Exchange, listenByAddress bool, restDial string, wsDial string, waitMilliseconds string) *UniswapScraper {
	var restClient, wsClient *ethclient.Client
	var err error
	var s *UniswapScraper

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

	s = &UniswapScraper{
		WsClient:        wsClient,
		RestClient:      restClient,
		shutdown:        make(chan nothing),
		shutdownDone:    make(chan nothing),
		pairScrapers:    make(map[string]*UniswapPairScraper),
		exchangeName:    exchange.Name,
		error:           nil,
		chanTrades:      make(chan *dia.Trade),
		waitTime:        waitTime,
		listenByAddress: listenByAddress,
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
				s.ListenToPair(index, address, s.listenByAddress)
			}(i, common.Address{}, &wg)
		}
		wg.Wait()

	}
}

// ListenToPair subscribes to a uniswap pool.
// If @byAddress is true, it listens by pool address, otherwise by index.
func (s *UniswapScraper) ListenToPair(i int, address common.Address, byAddress bool) {
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

	if len(pair.Token0.Symbol) < 2 || len(pair.Token1.Symbol) < 2 {
		log.Info("skip pair: ", pair.ForeignName)
		return
	}

	if helpers.AddressIsBlacklisted(pair.Token0.Address) || helpers.AddressIsBlacklisted(pair.Token1.Address) {
		log.Info("skip pair ", pair.ForeignName, ", address is blacklisted")
		return
	}

	// Normalize WETH on Ethereum
	if Exchanges[s.exchangeName].BlockChain.Name == dia.ETHEREUM {
		pair.normalizeUniPair()
	}
	// ps := s.pairScrapers[pair.ForeignName]
	// if ok {
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
					ForeignTradeID: swap.ID,
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
				case token0.Address == "0x0000000000000000000000000000000000000000" && !utils.Contains(&mainBaseAssets, token1.Address):
					// Reverse almost all pairs ETH-XXX ...
					if s.exchangeName == dia.UniswapExchange || s.exchangeName == dia.SushiSwapExchange {
						tSwapped, err := dia.SwapTrade(*t)
						if err == nil {
							t = &tSwapped
						}
					}
				// ...and USDT-XXX on Ethereum, i.e. Uniswap and Sushiswap
				case token0.Address == mainBaseAssets[0] && token0.Blockchain == dia.ETHEREUM:
					tSwapped, err := dia.SwapTrade(*t)
					if err == nil {
						t = &tSwapped
					}
				// Reverse USDC-XXX pairs on Fantom
				case token0.Address == "0x04068DA6C83AFCFA0e13ba15A6696662335D5B75" && token0.Blockchain == dia.FANTOM:
					tSwapped, err := dia.SwapTrade(*t)
					if err == nil {
						t = &tSwapped
					}
				}
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
			uniPair.normalizeUniPair()
			pairs[index] = uniPair
		}(i)
	}
	return pairs, nil
}

func (up *UniswapScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// Account for WETH is identified with ETH
func (up *UniswapPair) normalizeUniPair() {
	if up.Token0.Address == common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2") {
		up.Token0.Symbol = "ETH"
		up.Token0.Address = common.HexToAddress("0x0000000000000000000000000000000000000000")
		up.ForeignName = up.Token0.Symbol + "-" + up.Token1.Symbol
	}
	if up.Token1.Address == common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2") {
		up.Token1.Symbol = "ETH"
		up.Token1.Address = common.HexToAddress("0x0000000000000000000000000000000000000000")
		up.ForeignName = up.Token0.Symbol + "-" + up.Token1.Symbol
	}
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
