package scrapers

import (
	"errors"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/velodrome"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	velodromeExchangeFactoryContractAddress = "0xF1046053aa5682b4F9a81b5481394DA16BE5FF5a"
)

type VelodromeScraper struct {
	RestClient *ethclient.Client
	// relDB      *models.RelDB
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
	pairScrapers map[string]*VelodromePairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
	testChan     chan int
	waitTime     int
	// If true, only pairs given in config file are scraped. Default is false.
	listenByAddress  bool
	fetchPoolsFromDB bool
}

func NewVelodromeScraper(exchange dia.Exchange, scrape bool) *VelodromeScraper {
	var (
		restClient *ethclient.Client
		err        error
		s          *VelodromeScraper
		waitTime   int
	)
	restClient, err = ethclient.Dial(exchange.RestAPI)
	waitTime, err = strconv.Atoi(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_WAIT_TIME", "500"))

	if err != nil {
		log.Fatal("waittime error", err)
	}

	s = &VelodromeScraper{
		RestClient:       restClient,
		shutdown:         make(chan nothing),
		shutdownDone:     make(chan nothing),
		pairScrapers:     make(map[string]*VelodromePairScraper),
		exchangeName:     exchange.Name,
		error:            nil,
		chanTrades:       make(chan *dia.Trade),
		testChan:         make(chan int),
		waitTime:         waitTime,
		listenByAddress:  true,
		fetchPoolsFromDB: true,
	}

	if scrape {
		go s.mainLoop()
	}

	return s

}

func (s *VelodromeScraper) mainLoop() {

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
			pair, err := s.GetPairByID(int64(index))
			if err != nil {
				log.Error("error retrieving pair by ID: ", err)
				return
			}
			pairs[index] = pair
		}(i)
	}
	return pairs, nil
}

func (s *VelodromeScraper) GetPairByID(num int64) (dia.ExchangePair, error) {
	connection := s.RestClient

	var contract *velodrome.PoolFactoryCaller
	contract, err := velodrome.NewPoolFactoryCaller(common.HexToAddress(velodromeExchangeFactoryContractAddress), connection)

	if err != nil {
		log.Error(err)
		return dia.ExchangePair{}, err
	}

	numToken := big.NewInt(num)
	pairAddress, err := contract.AllPools(&bind.CallOpts{}, numToken)

	if err != nil {
		log.Error(err)
		return dia.ExchangePair{}, err
	}

	pair, err := s.GetPairByAddress(pairAddress)
	if err != nil {
		log.Error(err)
		return dia.ExchangePair{}, err
	}

	return pair, err
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
		Exchange:       "Velodrome",
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
