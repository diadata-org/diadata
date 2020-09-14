package scrapers

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	factory "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/balancer/balancerfactory"
	pool "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/balancer/balancerpool"
	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/balancer/balancertoken"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"

	"github.com/diadata-org/diadata/pkg/dia"
)

const (
	BalancerApiDelay   = 20
	BalancerBatchDelay = 60 * 1

	factoryContract = "0x9424B1412450D0f8Fc2255FAf6046b98213B76Bd"

	infuraKey  = "9020e59e34ca4cf59cb243ecefb4e39e"
	startBlock = uint64(10780772 - 5250)

	balancerWsDial   = "wss://mainnet.infura.io/ws/v3/" + infuraKey
	balancerRestDial = "http://159.69.120.42:8545/"
)

type BalancerSwap struct {
	SellToken  string
	BuyToken   string
	SellVolume float64
	BuyVolume  float64
	ID         string
	Timestamp  int64
}

type BalancerToken struct {
	Symbol   string
	Decimals uint8
}

type BalancerScraper struct {
	exchangeName string

	// channels to signal events
	run          bool
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error
	closed    bool

	balancerTokensMap map[string]*BalancerToken
	pairScrapers      map[string]*BalancerPairScraper
	productPairIds    map[string]int
	chanTrades        chan *dia.Trade

	WsClient    *ethclient.Client
	RestClient  *ethclient.Client
	resubscribe chan string
	pools       map[string]struct{}
}

func NewBalancerScraper(exchangeName string) *BalancerScraper {
	scraper := &BalancerScraper{
		exchangeName:      exchangeName,
		initDone:          make(chan nothing),
		shutdown:          make(chan nothing),
		shutdownDone:      make(chan nothing),
		productPairIds:    make(map[string]int),
		pairScrapers:      make(map[string]*BalancerPairScraper),
		chanTrades:        make(chan *dia.Trade),
		balancerTokensMap: make(map[string]*BalancerToken),
		resubscribe:       make(chan string),
		pools:             make(map[string]struct{}),
	}

	wsClient, err := ethclient.Dial(balancerWsDial)
	if err != nil {
		log.Fatal(err)
	}
	scraper.WsClient = wsClient
	restClient, err := ethclient.Dial(balancerRestDial)
	if err != nil {
		log.Fatal(err)
	}
	scraper.RestClient = restClient

	go scraper.mainLoop()
	return scraper
}

func (scraper *BalancerScraper) mainLoop() {

	time.Sleep(5 * time.Second)

	scraper.run = true

	scraper.balancerTokensMap, _ = scraper.getAllTokensMap()

	pools, err := scraper.getAllLogNewPool()
	if err != nil {
		log.Error(err)
	}
	for pools.Next() {
		scraper.pools[pools.Event.Pool.Hex()] = struct{}{}
	}
	scraper.performSubscriptions()

	go func() {
		for scraper.run {
			pool := <-scraper.resubscribe

			if scraper.run {
				if pool == "NEW_POOLS" {
					log.Info("resubscribe to new pools")
					scraper.subscribeToNewPools()
				} else {
					log.Info("resubscribe to pool: " + pool)
					scraper.subscribeToNewSwaps(pool)
				}
			}
		}
	}()

	for scraper.run {
		if len(scraper.pairScrapers) == 0 {
			scraper.error = errors.New("Balancer: No pairs to scrape provided")
			log.Error(scraper.error.Error())
		}
	}

	time.Sleep(time.Duration(10) * time.Second)
	if scraper.error == nil {
		scraper.error = errors.New("main loop terminated by Close()")
	}
	scraper.cleanup(nil)

}

func (scraper *BalancerScraper) performSubscriptions() {
	for pool, _ := range scraper.pools {
		scraper.subscribeToNewSwaps(pool)
	}

	scraper.subscribeToNewPools()
}

func (scraper *BalancerScraper) subscribeToNewPools() error {
	sinkPool, subPool, err := scraper.getNewPoolLogChannel()
	if err != nil {
		log.Error(err)
	}
	go func() {
		fmt.Println("subscribed to NewPools")
		defer fmt.Println("Unsubscribed to NewPools")
		defer subPool.Unsubscribe()
		subscribed := true

		for scraper.run && subscribed {

			select {
			case err := <-subPool.Err():
				if err != nil {
					log.Error(err)
				}
				subscribed = false
				if scraper.run {
					scraper.resubscribe <- "NEW_POOLS"
				}
			case vLog := <-sinkPool:
				if _, ok := scraper.pools[vLog.Pool.Hex()]; !ok {
					scraper.pools[vLog.Pool.Hex()] = struct{}{}
					scraper.subscribeToNewSwaps(vLog.Pool.Hex())
				}
			}
		}
	}()

	return err
}

func (scraper *BalancerScraper) subscribeToNewSwaps(poolToSub string) error {
	sink, sub, err := scraper.getLogSwapsChannel(common.HexToAddress(poolToSub))
	if err != nil {
		log.Error(err)
		return err
	}

	go func() {
		fmt.Println("subscribed to pool: " + poolToSub)
		defer fmt.Println("Unsubscribed to pool: " + poolToSub)
		defer sub.Unsubscribe()
		subscribed := true
		for scraper.run && subscribed {

			select {
			case err := <-sub.Err():
				if err != nil {
					log.Error(err)
				}
				subscribed = false
				if scraper.run {
					scraper.resubscribe <- poolToSub
				}
			case vLog := <-sink:

				decimalsIn := int(scraper.balancerTokensMap[vLog.TokenIn.Hex()].Decimals)
				decimalsOut := int(scraper.balancerTokensMap[vLog.TokenOut.Hex()].Decimals)
				amountIn, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(vLog.TokenAmountIn), new(big.Float).SetFloat64(math.Pow10(decimalsIn))).Float64()
				amountOut, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(vLog.TokenAmountOut), new(big.Float).SetFloat64(math.Pow10(decimalsOut))).Float64()
				swap := BalancerSwap{
					SellToken:  scraper.balancerTokensMap[vLog.TokenIn.Hex()].Symbol,
					BuyToken:   scraper.balancerTokensMap[vLog.TokenOut.Hex()].Symbol,
					SellVolume: amountIn,
					BuyVolume:  amountOut,
					ID:         vLog.Raw.TxHash.String() + "-" + fmt.Sprint(vLog.Raw.Index),
					Timestamp:  time.Now().Unix(),
				}
				swap.normalizeETH()
				pair := swap.BuyToken + "-" + swap.SellToken
				pairScraper, ok := scraper.pairScrapers[pair]
				if !ok {
					return
				}

				// Get trading data from swap in "classic" format
				_, volume, price, err := getSwapDataBalancer(swap)

				if err != nil {
					log.Error("error parsing time: ", err)
				}

				trade := &dia.Trade{
					Symbol:         pairScraper.pair.Symbol,
					Pair:           pair,
					Price:          price,
					Volume:         volume,
					Time:           time.Unix(swap.Timestamp, 0),
					ForeignTradeID: swap.ID,
					Source:         scraper.exchangeName,
				}
				pairScraper.parent.chanTrades <- trade
				fmt.Println("got trade: ", trade)

			}
		}
	}()

	return err

}

// getSwapData returns the foreign name, volume and price of a swap
func getSwapDataBalancer(s BalancerSwap) (foreignName string, volume float64, price float64, err error) {
	volume = s.BuyVolume
	price = s.SellVolume / s.BuyVolume
	foreignName = s.BuyToken + "-" + s.SellToken
	return
}

func (scraper *BalancerScraper) getAllLogNewPool() (*factory.BalancerfactoryLOGNEWPOOLIterator, error) {

	var pairFiltererContract *factory.BalancerfactoryFilterer
	pairFiltererContract, err := factory.NewBalancerfactoryFilterer(common.HexToAddress(factoryContract), scraper.RestClient)
	if err != nil {
		log.Fatal(err)
	}
	start := startBlock
	itr, _ := pairFiltererContract.FilterLOGNEWPOOL(&bind.FilterOpts{Start: start}, []common.Address{}, []common.Address{})
	if err != nil {
		log.Error("error in getAllLogNewPool ", err)
	}
	return itr, err
}

func (scraper *BalancerScraper) getAllTokenAddress() (map[string]struct{}, error) {
	it, err := scraper.getAllLogNewPool()
	if err != nil {
		log.Error(err)
	}

	tokenSet := make(map[string]struct{})
	for it.Next() {

		poolCaller, err := pool.NewBalancerpoolCaller(it.Event.Pool, scraper.RestClient)
		if err != nil {
			log.Error(err)
		}
		tokens, err := poolCaller.GetCurrentTokens(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
		}

		for _, token := range tokens {

			if _, ok := tokenSet[token.Hex()]; !ok {
				tokenSet[token.Hex()] = struct{}{}
			}
		}
	}

	return tokenSet, err
}

func (scraper *BalancerScraper) getAllTokensMap() (map[string]*BalancerToken, error) {
	tokenAddressSet, err := scraper.getAllTokenAddress()
	if err != nil {
		log.Error(err)
	}

	tokenMap := make(map[string]*BalancerToken)

	for token, _ := range tokenAddressSet {

		tokenCaller, err := balancertoken.NewBalancertokenCaller(common.HexToAddress(token), scraper.RestClient)
		if err != nil {
			log.Error(err)
		}
		symbol, err := tokenCaller.Symbol(&bind.CallOpts{})
		if err != nil {
			log.Error("Error: %v", err)
		}
		decimals, err := tokenCaller.Decimals(&bind.CallOpts{})
		if err != nil {
			log.Error("Error: %v", token, err)
		}
		if symbol != "" {
			tokenMap[token] = &BalancerToken{
				Symbol:   symbol,
				Decimals: decimals,
			}
		}
	}
	return tokenMap, err
}

// FetchAvailablePairs get pairs by geting all the LOGNEWPOOL contract events, and
// calling the method getCurrentTokens from each pool contract
func (scraper *BalancerScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {

	tokenMap, err := scraper.getAllTokensMap()
	if err != nil {
		log.Error(err)
	}

	pairSet := make(map[string]struct{})
	for _, token1 := range tokenMap {
		for _, token2 := range tokenMap {

			if token1 != token2 {

				foreignName := token1.Symbol + "-" + token2.Symbol
				if _, ok := pairSet[foreignName]; !ok {
					pairs = append(pairs, dia.Pair{
						Symbol:      token1.Symbol,
						ForeignName: foreignName,
						Exchange:    scraper.exchangeName,
						Ignore:      false,
					})
					pairSet[foreignName] = struct{}{}
				}

				foreignName = token2.Symbol + "-" + token1.Symbol
				if _, ok := pairSet[foreignName]; !ok {
					pairs = append(pairs, dia.Pair{
						Symbol:      token2.Symbol,
						ForeignName: foreignName,
						Exchange:    scraper.exchangeName,
						Ignore:      false,
					})
					pairSet[foreignName] = struct{}{}
				}

			}
		}
	}

	return

}

func (scraper *BalancerScraper) getLogSwapsChannelFilter(address string) (chan *pool.BalancerpoolLOGSWAP, error) {
	sink := make(chan *pool.BalancerpoolLOGSWAP)
	var pairFiltererContract *pool.BalancerpoolFilterer
	pairFiltererContract, _ = pool.NewBalancerpoolFilterer(common.HexToAddress(address), scraper.RestClient)

	start := startBlock
	itr, _ := pairFiltererContract.FilterLOGSWAP(&bind.FilterOpts{Start: start}, []common.Address{}, []common.Address{}, []common.Address{})
	scraper.balancerTokensMap, _ = scraper.getAllTokensMap()
	for itr.Next() {
		// vLog := itr.Event
	}
	return sink, nil
}

func (scraper *BalancerScraper) getLogSwapsChannel(poolAddress common.Address) (chan *pool.BalancerpoolLOGSWAP, event.Subscription, error) {
	sink := make(chan *pool.BalancerpoolLOGSWAP)
	var pairFiltererContract *pool.BalancerpoolFilterer
	pairFiltererContract, err := pool.NewBalancerpoolFilterer(poolAddress, scraper.WsClient)
	if err != nil {
		log.Fatal(err)
	}
	start := startBlock

	sub, _ := pairFiltererContract.WatchLOGSWAP(&bind.WatchOpts{Start: &start}, sink, nil, nil, nil)
	if err != nil {
		log.Error("error in get swaps channel: ", err)
	}

	return sink, sub, nil
}

func (scraper *BalancerScraper) getNewPoolLogChannel() (chan *factory.BalancerfactoryLOGNEWPOOL, event.Subscription, error) {
	sink := make(chan *factory.BalancerfactoryLOGNEWPOOL)
	var factoryFiltererContract *factory.BalancerfactoryFilterer
	factoryFiltererContract, err := factory.NewBalancerfactoryFilterer(common.HexToAddress(factoryContract), scraper.WsClient)
	if err != nil {
		log.Fatal(err)
	}
	start := startBlock

	sub, _ := factoryFiltererContract.WatchLOGNEWPOOL(&bind.WatchOpts{Start: &start}, sink, nil, nil)
	if err != nil {
		log.Error("error in get pools channel: ", err)
	}

	return sink, sub, nil
}

func (bs BalancerSwap) normalizeETH() {
	if bs.SellToken == "WETH" {
		bs.SellToken = "ETH"
	}
	if bs.BuyToken == "WETH" {
		bs.BuyToken = "ETH"
	}

}

func (scraper *BalancerScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()

	if scraper.error != nil {
		return nil, scraper.error
	}

	if scraper.closed {
		return nil, errors.New("BalancerScraper is closed")
	}

	pairScraper := &BalancerPairScraper{
		parent: scraper,
		pair:   pair,
	}

	scraper.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}
func (scraper *BalancerScraper) cleanup(err error) {
	scraper.errorLock.Lock()
	defer scraper.errorLock.Unlock()
	if err != nil {
		scraper.error = err
	}
	scraper.closed = true
	close(scraper.shutdownDone)
}

func (scraper *BalancerScraper) Close() error {
	// close the pair scraper channels
	scraper.run = false
	for _, pairScraper := range scraper.pairScrapers {
		pairScraper.closed = true
	}
	scraper.WsClient.Close()
	scraper.RestClient.Close()

	close(scraper.shutdown)
	<-scraper.shutdownDone
	return nil
}

type BalancerPairScraper struct {
	parent *BalancerScraper
	pair   dia.Pair
	closed bool
}

func (pairScraper *BalancerPairScraper) Pair() dia.Pair {
	return pairScraper.pair
}

func (scraper *BalancerScraper) Channel() chan *dia.Trade {
	return scraper.chanTrades
}

func (pairScraper *BalancerPairScraper) Error() error {
	s := pairScraper.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (pairScraper *BalancerPairScraper) Close() error {
	pairScraper.parent.errorLock.RLock()
	defer pairScraper.parent.errorLock.RUnlock()
	pairScraper.closed = true
	return nil
}
