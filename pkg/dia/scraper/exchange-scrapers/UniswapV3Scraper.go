package scrapers

import (
	"errors"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	uniswapcontract "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswap"
	uniswapcontractv3 "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswapv3"
	UniswapV3Pair "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswapv3/uniswapV3Pair"

	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UniswapV3Swap struct {
	ID        string
	Timestamp int64
	Pair      UniswapPair
	Amount0   float64
	Amount1   float64
}

type UniswapV3Scraper struct {
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
	pairScrapers map[string]*UniswapPairV3Scraper
	pairRecieved chan *UniswapPair

	exchangeName           string
	startBlock             uint64
	waitTime               int
	listenByAddress        bool
	chanTrades             chan *dia.Trade
	factoryContractAddress common.Address
}

// NewUniswapV3Scraper returns a new UniswapV3Scraper
func NewUniswapV3Scraper(exchange dia.Exchange, scrape bool) *UniswapV3Scraper {
	log.Info("NewUniswapScraper ", exchange.Name)
	log.Info("NewUniswapScraper Address ", exchange.Contract)

	var s *UniswapV3Scraper

	switch exchange.Name {
	case dia.UniswapExchangeV3:
		s = makeUniswapV3Scraper(exchange, false, "", "", "200", uint64(12369621))
	case dia.UniswapExchangeV3Polygon:
		s = makeUniswapV3Scraper(exchange, false, "", "", "200", uint64(22757913))
	case dia.UniswapExchangeV3Arbitrum:
		s = makeUniswapV3Scraper(exchange, false, "", "", "200", uint64(165))
	}

	if scrape {
		go s.mainLoop()
	}
	return s
}

// makeUniswapV3Scraper returns a uniswap scraper as used in NewUniswapV3Scraper.
func makeUniswapV3Scraper(exchange dia.Exchange, listenByAddress bool, restDial string, wsDial string, waitMilliseconds string, startBlock uint64) *UniswapV3Scraper {
	var restClient, wsClient *ethclient.Client
	var err error
	var s *UniswapV3Scraper

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

	s = &UniswapV3Scraper{
		WsClient:               wsClient,
		RestClient:             restClient,
		shutdown:               make(chan nothing),
		shutdownDone:           make(chan nothing),
		pairScrapers:           make(map[string]*UniswapPairV3Scraper),
		exchangeName:           exchange.Name,
		pairRecieved:           make(chan *UniswapPair),
		error:                  nil,
		chanTrades:             make(chan *dia.Trade),
		waitTime:               waitTime,
		listenByAddress:        listenByAddress,
		startBlock:             startBlock,
		factoryContractAddress: common.HexToAddress(exchange.Contract),
	}
	return s
}

// runs in a goroutine until s is closed
func (s *UniswapV3Scraper) mainLoop() {

	var err error
	reverseBasetokens, err = getReverseTokensFromConfig("uniswapv3/reverse_tokens/" + s.exchangeName + "Basetoken")
	if err != nil {
		log.Error("error getting basetokens for which pairs should be reversed: ", err)
	}
	log.Infof("reverse the following basetokens on %s: %v", s.exchangeName, reverseBasetokens)
	reverseQuotetokens, err = getReverseTokensFromConfig("uniswapv3/reverse_tokens/" + s.exchangeName + "Quotetoken")
	if err != nil {
		log.Error("error getting quotetokens for which pairs should be reversed: ", err)
	}
	log.Infof("reverse the following quotetokens on %s: %v", s.exchangeName, reverseQuotetokens)

	time.Sleep(4 * time.Second)
	s.run = true

	go func() {
		pairs, err := s.getAllPairs()
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Found ", len(pairs), " pairs")
		log.Info("Found ", len(s.pairScrapers), " pairScrapers")
	}()

	if len(s.pairScrapers) == 0 {
		s.error = errors.New("uniswap: No pairs to scrape provided")
		log.Error(s.error.Error())
	}
	for {
		pair := <-s.pairRecieved
		log.Infoln("Subscribing for pair", pair)

		if len(pair.Token0.Symbol) < 2 || len(pair.Token1.Symbol) < 2 {
			log.Info("skip pair: ", pair.ForeignName)
			continue
		}
		if helpers.AddressIsBlacklisted(pair.Token0.Address) || helpers.AddressIsBlacklisted(pair.Token1.Address) {
			log.Info("skip pair ", pair.ForeignName, ", address is blacklisted")
			continue
		}
		if helpers.PoolIsBlacklisted(pair.Address) {
			log.Info("skip blacklisted pool ", pair.Address)
			continue
		}

		log.Info("found pair scraper for: ", pair.ForeignName, " with address ", pair.Address.Hex())
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
					price, volume := s.getSwapData(swap)
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

					switch {
					case utils.Contains(reverseBasetokens, pair.Token1.Address.Hex()):
						// If we need quotation of a base token, reverse pair
						tSwapped, err := dia.SwapTrade(*t)
						if err == nil {
							t = &tSwapped
						}
					case utils.Contains(reverseQuotetokens, pair.Token0.Address.Hex()):
						// If we need quotation of a base token, reverse pair
						tSwapped, err := dia.SwapTrade(*t)
						if err == nil {
							t = &tSwapped
						}
					}
					if price > 0 {
						log.Info("Got trade: ", t)
						s.chanTrades <- t
					}
				}
			}
		}()

	}
}

// GetSwapsChannel returns a channel for swaps of the pair with address @pairAddress
func (s *UniswapV3Scraper) GetSwapsChannel(pairAddress common.Address) (chan *UniswapV3Pair.UniswapV3PairSwap, error) {
	sink := make(chan *UniswapV3Pair.UniswapV3PairSwap)
	var pairFiltererContract *UniswapV3Pair.UniswapV3PairFilterer

	pairFiltererContract, err := UniswapV3Pair.NewUniswapV3PairFilterer(pairAddress, s.WsClient)
	if err != nil {
		log.Fatal(err)
	}

	_, err = pairFiltererContract.WatchSwap(&bind.WatchOpts{}, sink, []common.Address{}, []common.Address{})
	if err != nil {
		log.Error("error in get swaps channel: ", err)
	}

	return sink, nil

}

func (s *UniswapV3Scraper) getSwapData(swap UniswapV3Swap) (price float64, volume float64) {
	volume = swap.Amount0
	price = math.Abs(swap.Amount1 / swap.Amount0)
	return
}

// normalizeUniswapSwap takes a swap as returned by the swap contract's channel and converts it to a UniswapSwap type
func (s *UniswapV3Scraper) normalizeUniswapSwap(swap UniswapV3Pair.UniswapV3PairSwap) (normalizedSwap UniswapV3Swap, err error) {

	pair, err := s.GetPairByAddress(swap.Raw.Address)
	if err != nil {
		log.Error("error getting pair by address: ", err)
		return
	}
	decimals0 := int(pair.Token0.Decimals)
	decimals1 := int(pair.Token1.Decimals)
	amount0, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount0), new(big.Float).SetFloat64(math.Pow10(decimals0))).Float64()
	amount1, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount1), new(big.Float).SetFloat64(math.Pow10(decimals1))).Float64()

	normalizedSwap = UniswapV3Swap{
		ID:        swap.Raw.TxHash.Hex(),
		Timestamp: time.Now().Unix(),
		Pair:      pair,
		Amount0:   amount0,
		Amount1:   amount1,
	}
	return
}

// GetPairByAddress returns the UniswapPair with pair address @pairAddress
func (s *UniswapV3Scraper) GetPairByAddress(pairAddress common.Address) (pair UniswapPair, err error) {
	connection := s.RestClient
	var pairContract *UniswapV3Pair.UniswapV3PairCaller
	pairContract, err = UniswapV3Pair.NewUniswapV3PairCaller(pairAddress, connection)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}

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

// FetchAvailablePairs returns a list with all available trade pairs as dia.Pair for the pairDiscorvery service
func (s *UniswapV3Scraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return
}

func (s *UniswapV3Scraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *UniswapV3Scraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// GetPairByID returns the UniswapPair with the integer id @num
func (s *UniswapV3Scraper) GetPairData(poolEvent *uniswapcontractv3.UniswapV3PoolCreated) (UniswapPair, error) {
	pair, err := s.GetPairByTokenAddress(poolEvent.Token0, poolEvent.Token1, poolEvent.Pool)
	if err != nil {
		log.Error("GetPairData", err)
		return UniswapPair{}, err
	}
	return pair, err
}

func (s *UniswapV3Scraper) GetPairByTokenAddress(address0 common.Address, address1 common.Address, pairAddress common.Address) (pair UniswapPair, err error) {
	connection := s.RestClient

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
func (s *UniswapV3Scraper) GetDecimals(tokenAddress common.Address) (decimals uint8, err error) {

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
func (s *UniswapV3Scraper) getAllPairs() (pairs []UniswapPair, err error) {

	// filter from contract created https://etherscan.io/tx/0x1e20cd6d47d7021ae7e437792823517eeadd835df09dde17ab45afd7a5df4603

	log.Info("get pool creations from address: ", s.factoryContractAddress.Hex())
	poolsCount := 0
	contract, err := uniswapcontractv3.NewUniswapV3Filterer(s.factoryContractAddress, s.WsClient)
	if err != nil {
		log.Error(err)
	}

	tInit := time.Now()
	poolCreated, err := contract.FilterPoolCreated(
		&bind.FilterOpts{Start: s.startBlock},
		[]common.Address{},
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		return nil, err
	}
	log.Info("time spent for filter pool created: ", time.Since(tInit))
	for poolCreated.Next() {
		poolsCount++
		log.Info("pools count: ", poolsCount)
		pair, _ := s.GetPairData(poolCreated.Event)
		pairs = append(pairs, pair)
		s.pairRecieved <- &pair
	}

	return pairs, nil

}

func (s *UniswapV3Scraper) cleanup(err error) {
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
func (s *UniswapV3Scraper) Close() error {

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
func (s *UniswapV3Scraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("UniswapScraper: Call ScrapePair on closed scraper")
	}
	ps := &UniswapPairV3Scraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

// UniswapPairScraper implements PairScraper for Uniswap
type UniswapPairV3Scraper struct {
	parent *UniswapV3Scraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *UniswapPairV3Scraper) Close() error {
	return nil
}

// Channel returns a channel that can be used to receive trades
func (s *UniswapV3Scraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *UniswapPairV3Scraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *UniswapPairV3Scraper) Pair() dia.ExchangePair {
	return ps.pair
}
