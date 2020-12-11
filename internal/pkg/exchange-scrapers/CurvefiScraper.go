package scrapers

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/curvefi"
	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/curvefi/curvepool"
	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/curvefi/token"
)

const (
	curveFiContract       = "0x7002B727Ef8F5571Cb5F9D70D13DBEEb4dFAe9d1"
	curveFiLookBackBlocks = 6 * 60 * 24 * 20
	curveWsDial           = "ws://159.69.120.42:8546/"
	curveRestDial         = "http://159.69.120.42:8545/"
)

type CurveCoin struct {
	Symbol   string
	Decimals uint8
}

type Pools struct {
	pools     map[string]map[int]*CurveCoin
	poolsLock sync.RWMutex
}

func (p *Pools) setPool(k string, v map[int]*CurveCoin) {
	p.poolsLock.Lock()
	defer p.poolsLock.Unlock()
	p.pools[k] = v
}

func (p *Pools) getPool(k string) (map[int]*CurveCoin, bool) {
	p.poolsLock.RLock()
	defer p.poolsLock.RUnlock()
	r, ok := p.pools[k]
	return r, ok
}
func (p *Pools) getPoolCoin(poolk string, coink int) (*CurveCoin, bool) {
	p.poolsLock.RLock()
	defer p.poolsLock.RUnlock()
	r, ok := p.pools[poolk][coink]
	return r, ok
}

func (p *Pools) poolsAddressNoLock() []string {
	p.poolsLock.RLock()
	defer p.poolsLock.RUnlock()
	var values []string
	for key, _ := range p.pools {
		values = append(values, key)
	}
	return values
}

type CurveFIScraper struct {
	exchangeName string

	// channels to signal events
	run          bool
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers   map[string]*CurveFIPairScraper
	productPairIds map[string]int
	chanTrades     chan *dia.Trade

	WsClient    *ethclient.Client
	RestClient  *ethclient.Client
	curveCoins  map[string]*CurveCoin
	resubscribe chan string
	pools       *Pools
	contract    common.Address
}

func NewCurveFIScraper(exchange dia.Exchange) *CurveFIScraper {
	scraper := &CurveFIScraper{
		exchangeName:   exchange.Name,
		contract:       exchange.Contract,
		initDone:       make(chan nothing),
		shutdown:       make(chan nothing),
		shutdownDone:   make(chan nothing),
		productPairIds: make(map[string]int),
		pairScrapers:   make(map[string]*CurveFIPairScraper),
		chanTrades:     make(chan *dia.Trade),
		curveCoins:     make(map[string]*CurveCoin),
		resubscribe:    make(chan string),
		pools: &Pools{
			pools: make(map[string]map[int]*CurveCoin),
		},
	}

	wsClient, err := ethclient.Dial(curveWsDial)
	if err != nil {
		log.Fatal(err)
	}
	scraper.WsClient = wsClient
	restClient, err := ethclient.Dial(curveRestDial)
	if err != nil {
		log.Fatal(err)
	}
	scraper.RestClient = restClient

	scraper.loadPoolsAndCoins()

	go scraper.mainLoop()
	return scraper
}

func (scraper *CurveFIScraper) mainLoop() {
	scraper.run = true

	for _, pool := range scraper.pools.poolsAddressNoLock() {
		scraper.watchSwaps(pool)
	}
	scraper.watchNewPools()

	go func() {
		for scraper.run {
			p := <-scraper.resubscribe

			if scraper.run {
				if p == "NEW_POOLS" {
					log.Info("resubscribe to new pools")
					scraper.watchNewPools()
				} else {
					log.Info("resubscribe to swaps from Pool: " + p)
					scraper.watchSwaps(p)
				}
			}
		}
	}()

	for scraper.run {
		if len(scraper.pairScrapers) == 0 {
			scraper.error = errors.New("Curvefi: No pairs to scrape provided")
			log.Error(scraper.error.Error())
			break
		}
	}
	time.Sleep(time.Duration(10) * time.Second)
	if scraper.error == nil {
		scraper.error = errors.New("Main loop terminated by Close().")
	}
	scraper.cleanup(nil)
}

func (scraper *CurveFIScraper) watchNewPools() {
	contract, err := curvefi.NewCurvefiFilterer(scraper.contract, scraper.WsClient)
	if err != nil {
		log.Error(err)
	}
	sink := make(chan *curvefi.CurvefiPoolAdded)

	header, err := scraper.RestClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	startblock := header.Number.Uint64() - uint64(curveFiLookBackBlocks)

	sub, err := contract.WatchPoolAdded(&bind.WatchOpts{Start: &startblock}, sink, nil)
	if err != nil {
		log.Error(err)
	}

	go func() {
		fmt.Println("subscribed to new pools")
		defer fmt.Println("Unsubscribed to new pools")
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
					scraper.resubscribe <- "NEW_POOLS"
				}
			case vLog := <-sink:

				if _, ok := scraper.pools.getPool(vLog.Pool.Hex()); !ok {
					scraper.loadPoolData(vLog.Pool.Hex())
					scraper.watchSwaps(vLog.Pool.Hex())
				}
			}
		}
	}()

}

// contract.poolList.map(contract.GetPoolCoins(pool).)
func (scraper *CurveFIScraper) loadPoolsAndCoins() error {
	contract, err := curvefi.NewCurvefiCaller(scraper.contract, scraper.RestClient)
	if err != nil {
		log.Error(err)
	}
	poolCount, err := contract.PoolCount(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}
	for i := 0; i < int(poolCount.Int64()); i++ {
		pool, err := contract.PoolList(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			log.Error(err)
		}

		scraper.loadPoolData(pool.Hex())

	}
	return err
}

func (scraper *CurveFIScraper) loadPoolData(pool string) error {
	contract, err := curvefi.NewCurvefiCaller(scraper.contract, scraper.RestClient)
	if err != nil {
		log.Error(err)
	}

	poolCoinsMap := make(map[int]*CurveCoin)

	poolCoins, err := contract.GetPoolCoins(&bind.CallOpts{}, common.HexToAddress(pool))
	if err != nil {
		log.Error(err)
	}

	for cIdx, c := range poolCoins.Coins {

		coinCaller, err := token.NewTokenCaller(c, scraper.RestClient)
		if err != nil {
			log.Error(err)
			continue
		}
		symbol, err := coinCaller.Symbol(&bind.CallOpts{})
		if err != nil {
			log.Error(err, c.Hex())
			continue
		}
		decimals, err := coinCaller.Decimals(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
			continue
		}

		poolCoinsMap[cIdx] = &CurveCoin{
			Symbol:   symbol,
			Decimals: uint8(decimals),
		}
		scraper.curveCoins[c.Hex()] = &CurveCoin{
			Symbol:   symbol,
			Decimals: decimals,
		}

		scraper.pools.setPool(pool, poolCoinsMap)

	}
	return err
}

func (scraper *CurveFIScraper) processSwap(pool string, swp *curvepool.CurvepoolTokenExchange) {

	foreignName, volume, price, err := scraper.getSwapDataCurve(pool, swp)
	if err != nil {
		log.Error(err)
	}
	timestamp := time.Now().Unix()

	if pairScraper, ok := scraper.pairScrapers[foreignName]; ok {

		trade := &dia.Trade{
			Symbol:         pairScraper.pair.Symbol,
			Pair:           foreignName,
			Price:          price,
			Volume:         volume,
			Time:           time.Unix(timestamp, 0),
			ForeignTradeID: swp.Raw.TxHash.Hex() + "-" + fmt.Sprint(swp.Raw.Index),
			Source:         scraper.exchangeName,
		}
		log.Infoln("Got Trade  ", trade)

		scraper.chanTrades <- trade
	}
}

func (scraper *CurveFIScraper) watchSwaps(pool string) error {

	filterer, err := curvepool.NewCurvepoolFilterer(common.HexToAddress(pool), scraper.WsClient)
	if err != nil {
		log.Fatal(err)
	}
	sink := make(chan *curvepool.CurvepoolTokenExchange)

	header, err := scraper.RestClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	startblock := header.Number.Uint64() - uint64(15250)

	sub, err := filterer.WatchTokenExchange(&bind.WatchOpts{Start: &startblock}, sink, nil)

	if err != nil {
		log.Error(err)
	}

	go func() {
		fmt.Println("Curvefi Subscribed to pool: " + pool)
		defer fmt.Println("Curvefi UnSubscribed to pool: " + pool)
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
					scraper.resubscribe <- pool
				}
			case swp := <-sink:

				// fmt.Println(swp)
				scraper.processSwap(pool, swp)
			}
		}
	}()
	return err

}

// getSwapDataCurve returns the foreign name, volume and price of a swap
func (scraper *CurveFIScraper) getSwapDataCurve(pool string, s *curvepool.CurvepoolTokenExchange) (foreignName string, volume float64, price float64, err error) {

	// fromToken, _ := scraper.curveCoins[s.TokenSold.Hex()]
	// toToken, _ := scraper.curveCoins[s.TokenBought.Hex()]

	fromToken, ok := scraper.pools.getPoolCoin(pool, int(s.SoldId.Int64()))
	if !ok {
		err = fmt.Errorf("token not found: " + pool + "-" + s.SoldId.String())
	}
	toToken, ok := scraper.pools.getPoolCoin(pool, int(s.BoughtId.Int64()))
	if !ok {
		err = fmt.Errorf("token not found: " + pool + "-" + s.SoldId.String())
	}

	// amountIn := s.AmountSold. / math.Pow10( fromToken.Decimals )
	amountIn, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(s.TokensSold), new(big.Float).SetFloat64(math.Pow10(int(fromToken.Decimals)))).Float64()

	// amountOut := s.AmountBought / math.Pow10( toToken.Decimals )
	amountOut, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(s.TokensBought), new(big.Float).SetFloat64(math.Pow10(int(toToken.Decimals)))).Float64()

	volume = amountOut
	price = amountIn / amountOut

	foreignName = toToken.Symbol + "-" + fromToken.Symbol

	return
}

func (scraper *CurveFIScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {

	pairSet := make(map[string]struct{})
	for _, p1 := range scraper.curveCoins {
		for _, p2 := range scraper.curveCoins {
			token1 := p1
			token2 := p2
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

func (scraper *CurveFIScraper) NormalizePair(pair dia.Pair) (dia.Pair, error) {
	return dia.Pair{}, nil
}

func (scraper *CurveFIScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()

	if scraper.error != nil {
		return nil, scraper.error
	}

	if scraper.closed {
		return nil, errors.New("CurveFIScraper is closed")
	}

	pairScraper := &CurveFIPairScraper{
		parent: scraper,
		pair:   pair,
	}

	scraper.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}
func (scraper *CurveFIScraper) cleanup(err error) {
	scraper.errorLock.Lock()
	defer scraper.errorLock.Unlock()
	if err != nil {
		scraper.error = err
	}
	scraper.closed = true
	close(scraper.shutdownDone)
}

func (scraper *CurveFIScraper) Close() error {
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

type CurveFIPairScraper struct {
	parent *CurveFIScraper
	pair   dia.Pair
	closed bool
}

func (pairScraper *CurveFIPairScraper) Pair() dia.Pair {
	return pairScraper.pair
}

func (scraper *CurveFIScraper) Channel() chan *dia.Trade {
	return scraper.chanTrades
}

func (pairScraper *CurveFIPairScraper) Error() error {
	s := pairScraper.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (pairScraper *CurveFIPairScraper) Close() error {
	pairScraper.parent.errorLock.RLock()
	defer pairScraper.parent.errorLock.RUnlock()
	pairScraper.closed = true
	return nil
}
