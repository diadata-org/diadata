package scrapers

import (
	"errors"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/velodrome"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	velodromeExchangeFactoryContractAddress = "0xF1046053aa5682b4F9a81b5481394DA16BE5FF5a"
	velodromeWaitTimeMilliseconds           = "500"
)

type VelodromeSwap struct {
	ID         string
	Timestamp  int64
	Pair       dia.ExchangePair
	Amount0In  float64
	Amount0Out float64
	Amount1In  float64
	Amount1Out float64
}

type VelodromeScraper struct {
	RestClient *ethclient.Client
	WsClient   *ethclient.Client
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*VelodromePairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
	waitTime     int
}

func NewVelodromeScraper(exchange dia.Exchange, scrape bool) *VelodromeScraper {
	log.Info("NewVelodromeScraper: ", exchange.Name)
	var (
		s                *VelodromeScraper
		listenByAddress  bool
		fetchPoolsFromDB bool
	)

	velodromeExchangeFactoryContractAddress = exchange.Contract

	switch exchange.Name {
	case dia.VelodromeExchange:
		s = makeVelodromeScraper(exchange, listenByAddress, fetchPoolsFromDB, velodromeWaitTimeMilliseconds)
	}

	if scrape {
		go s.mainLoop()
	}

	return s

}

func makeVelodromeScraper(exchange dia.Exchange, listenByAddress bool, fetchPoolsFromDB bool, waitMilliseconds string) *VelodromeScraper {
	var (
		restClient, wsClient *ethclient.Client
		err                  error
		s                    *VelodromeScraper
		waitTime             int
	)

	log.Infof("Init rest and ws client for %s.", exchange.BlockChain.Name)

	restClient, err = ethclient.Dial(exchange.RestAPI)
	if err != nil {
		log.Fatal("init rest client: ", err)
	}

	wsClient, err = ethclient.Dial(exchange.WsAPI)
	if err != nil {
		log.Fatal("init ws client: ", err)
	}

	waitTime, err = strconv.Atoi(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_WAIT_TIME", waitMilliseconds))
	if err != nil {
		log.Fatal("waittime error", err)
	}

	s = &VelodromeScraper{
		RestClient:   restClient,
		WsClient:     wsClient,
		pairScrapers: make(map[string]*VelodromePairScraper),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		waitTime:     waitTime,
	}

	return s
}

func (s *VelodromeScraper) mainLoop() {
	numPairs, err := s.getNumPairs()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Found ", numPairs, " pairs")
	log.Info("Found ", len(s.pairScrapers), " pairScrapers")

	if len(s.pairScrapers) == 0 {
		s.error = errors.New("velodrome: No pairs to scrap provided")
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

func (s *VelodromeScraper) ListenToPair(i int, address common.Address) {
	pairAddress, pair, err := s.GetPairByID(int64(i))

	if len(pair.UnderlyingPair.BaseToken.Symbol) < 2 || len(pair.UnderlyingPair.QuoteToken.Symbol) < 2 {
		log.Info("skip pair: ", pair.ForeignName)
		return
	}

	if helpers.AddressIsBlacklisted(common.HexToAddress(pair.UnderlyingPair.BaseToken.Address)) ||
		helpers.AddressIsBlacklisted(common.HexToAddress(pair.UnderlyingPair.QuoteToken.Address)) {
		log.Info("skip pair ", pair.ForeignName, ", address is blacklisted")
		return
	}
	if helpers.PoolIsBlacklisted(pairAddress) {
		log.Info("skip blacklisted pool ", pairAddress)
		return
	}

	log.Info(i, ": add pair scraper for: ", pair.ForeignName, " with address ", pairAddress.Hex())
	sink, err := s.GetSwapsChannel(pairAddress)
	if err != nil {
		log.Error("error fetching swaps channel: ", err)
	}

	go func() {
		for {
			rawSwap, ok := <-sink
			if ok {
				swap, err := s.normalizeSwap(*rawSwap, pair)
				if err != nil {
					log.Error("error normalizing swap: ", err)
				}
				price, volume := s.getSwapData(swap)
				token0 := dia.Asset{
					Address:    pair.UnderlyingPair.BaseToken.Address,
					Symbol:     pair.UnderlyingPair.BaseToken.Symbol,
					Name:       pair.UnderlyingPair.BaseToken.Name,
					Decimals:   pair.UnderlyingPair.BaseToken.Decimals,
					Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
				}
				token1 := dia.Asset{
					Address:    pair.UnderlyingPair.QuoteToken.Address,
					Symbol:     pair.UnderlyingPair.QuoteToken.Symbol,
					Name:       pair.UnderlyingPair.QuoteToken.Name,
					Decimals:   pair.UnderlyingPair.QuoteToken.Decimals,
					Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
				}
				t := &dia.Trade{
					Symbol:         pair.UnderlyingPair.BaseToken.Symbol,
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
func (s *VelodromeScraper) GetSwapsChannel(pairAddress common.Address) (chan *velodrome.IPoolSwap, error) {

	sink := make(chan *velodrome.IPoolSwap)
	var pairFiltererContract *velodrome.IPoolFilterer
	pairFiltererContract, err := velodrome.NewIPoolFilterer(pairAddress, s.WsClient)
	if err != nil {
		log.Fatal(err)
	}

	_, err = pairFiltererContract.WatchSwap(&bind.WatchOpts{}, sink, []common.Address{}, []common.Address{})
	if err != nil {
		log.Error("error in get swaps channel: ", err)
	}

	return sink, nil

}

// normalizeSwap takes a swap as returned by the swap contract's channel and converts it to a VelodromeSwap type
func (s *VelodromeScraper) normalizeSwap(swap velodrome.IPoolSwap, pair dia.ExchangePair) (normalizedSwap VelodromeSwap, err error) {

	decimals0 := int(pair.UnderlyingPair.BaseToken.Decimals)
	decimals1 := int(pair.UnderlyingPair.QuoteToken.Decimals)
	amount0In, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount0In), new(big.Float).SetFloat64(math.Pow10(decimals0))).Float64()
	amount0Out, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount0Out), new(big.Float).SetFloat64(math.Pow10(decimals0))).Float64()
	amount1In, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount1In), new(big.Float).SetFloat64(math.Pow10(decimals1))).Float64()
	amount1Out, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount1Out), new(big.Float).SetFloat64(math.Pow10(decimals1))).Float64()

	normalizedSwap = VelodromeSwap{
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

func (s *VelodromeScraper) getSwapData(swap VelodromeSwap) (price float64, volume float64) {
	if swap.Amount0In == float64(0) {
		volume = swap.Amount0Out
		price = swap.Amount1In / swap.Amount0Out
		return
	}
	volume = -swap.Amount0In
	price = swap.Amount1Out / swap.Amount0In
	return
}

// getNumPairs returns the number of available pairs on Velodrome
func (s *VelodromeScraper) getNumPairs() (int, error) {

	var contract *velodrome.IPoolFactoryCaller
	contract, err := velodrome.NewIPoolFactoryCaller(common.HexToAddress(velodromeExchangeFactoryContractAddress), s.RestClient)
	if err != nil {
		log.Error(err)
	}

	// Getting pairs ---------------
	numPairs, err := contract.AllPoolsLength(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}
	return int(numPairs.Int64()), err
}

func (s *VelodromeScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

func (s *VelodromeScraper) Close() error {
	s.closed = true
	return nil
}

func (s *VelodromeScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	pairs, err = s.GetAllPairs()
	return pairs, nil
}

func (s *VelodromeScraper) GetAllPairs() ([]dia.ExchangePair, error) {
	time.Sleep(20 * time.Millisecond)
	connection := s.RestClient

	var contract *velodrome.PoolFactoryCaller
	contract, err := velodrome.NewPoolFactoryCaller(common.HexToAddress(velodromeExchangeFactoryContractAddress), connection)

	if err != nil {
		log.Error(err)
	}

	numPairs, err := contract.AllPoolsLength(&bind.CallOpts{})

	if err != nil {
		return []dia.ExchangePair{}, err
	}

	wg := sync.WaitGroup{}
	defer wg.Wait()

	lenPairs := int(numPairs.Int64())
	pairs := make([]dia.ExchangePair, lenPairs)

	for i := 0; i < lenPairs; i++ {
		// Sleep in order not to run into rate limits.
		time.Sleep(time.Duration(s.waitTime) * time.Millisecond)
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			log.Info("fetching pair - ", int64(index))
			_, pair, err := s.GetPairByID(int64(index))
			if err != nil {
				log.Error("error retrieving pair by ID: ", err)
				return
			}
			pairs[index] = pair
		}(i)
	}
	return pairs, nil
}

func (s *VelodromeScraper) GetPairByID(num int64) (common.Address, dia.ExchangePair, error) {
	connection := s.RestClient

	var contract *velodrome.PoolFactoryCaller
	contract, err := velodrome.NewPoolFactoryCaller(common.HexToAddress(velodromeExchangeFactoryContractAddress), connection)

	if err != nil {
		log.Error(err)
		return common.Address{}, dia.ExchangePair{}, err
	}

	numToken := big.NewInt(num)
	pairAddress, err := contract.AllPools(&bind.CallOpts{}, numToken)

	if err != nil {
		log.Error(err)
		return common.Address{}, dia.ExchangePair{}, err
	}

	pair, err := s.GetPairByAddress(pairAddress)
	if err != nil {
		log.Error(err)
		return common.Address{}, dia.ExchangePair{}, err
	}

	return pairAddress, pair, err
}

func (s *VelodromeScraper) GetPairByAddress(pairAddress common.Address) (pair dia.ExchangePair, err error) {
	connection := s.RestClient
	var pairContract *velodrome.IPoolCaller
	pairContract, err = velodrome.NewIPoolCaller(pairAddress, connection)

	if err != nil {
		log.Error(err)
		return dia.ExchangePair{}, err
	}

	// Getting tokens from pair ---------------------
	address0, _ := pairContract.Token0(&bind.CallOpts{})
	address1, _ := pairContract.Token1(&bind.CallOpts{})
	var token0Contract *velodrome.IERC20MetadataCaller
	var token1Contract *velodrome.IERC20MetadataCaller
	token0Contract, err = velodrome.NewIERC20MetadataCaller(address0, connection)
	if err != nil {
		log.Error(err)
	}
	token1Contract, err = velodrome.NewIERC20MetadataCaller(address1, connection)
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
		return dia.ExchangePair{}, err
	}
	decimals1, err := s.GetDecimals(address1)
	if err != nil {
		log.Error(err)
		return dia.ExchangePair{}, err
	}

	name0, err := s.GetName(address0)
	if err != nil {
		log.Error(err)
		return dia.ExchangePair{}, err
	}
	name1, err := s.GetName(address1)
	if err != nil {
		log.Error(err)
		return dia.ExchangePair{}, err
	}
	token0 := dia.Asset{
		Address:    address0.String(),
		Symbol:     symbol0,
		Decimals:   decimals0,
		Name:       name0,
		Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
	}
	token1 := dia.Asset{
		Address:    address1.String(),
		Symbol:     symbol1,
		Decimals:   decimals1,
		Name:       name1,
		Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
	}
	foreignName := symbol0 + "-" + symbol1
	pair = dia.ExchangePair{
		Symbol:         token0.Symbol,
		ForeignName:    foreignName,
		Exchange:       s.exchangeName,
		Verified:       true,
		UnderlyingPair: dia.Pair{BaseToken: token0, QuoteToken: token1},
	}
	return pair, nil
}

// GetDecimals returns the decimals of the token with address @tokenAddress
func (s *VelodromeScraper) GetDecimals(tokenAddress common.Address) (decimals uint8, err error) {
	var contract *velodrome.IERC20MetadataCaller
	contract, err = velodrome.NewIERC20MetadataCaller(tokenAddress, s.RestClient)

	if err != nil {
		log.Error(err)
		return
	}
	decimals, err = contract.Decimals(&bind.CallOpts{})

	return
}

func (s *VelodromeScraper) GetName(tokenAddress common.Address) (name string, err error) {
	var contract *velodrome.IERC20MetadataCaller
	contract, err = velodrome.NewIERC20MetadataCaller(tokenAddress, s.RestClient)

	if err != nil {
		log.Error(err)
		return
	}
	name, err = contract.Name(&bind.CallOpts{})

	return
}

func (s *VelodromeScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{}, nil
}

func (up *VelodromeScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

type VelodromePairScraper struct {
	parent *VelodromeScraper
	pair   dia.ExchangePair
	closed bool
}

func (s *VelodromeScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("Velodrome: Call ScrapePair on closed scraper")
	}
	ps := &VelodromePairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

func (ps *VelodromePairScraper) Close() error {
	ps.closed = true
	return nil
}

func (ps *VelodromePairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *VelodromePairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
