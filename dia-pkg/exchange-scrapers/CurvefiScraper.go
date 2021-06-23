package scrapers

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	curvefi2 "github.com/diadata-org/diadata/dia-pkg/exchange-scrapers/curvefi"
	curvepool2 "github.com/diadata-org/diadata/dia-pkg/exchange-scrapers/curvefi/curvepool"
	token2 "github.com/diadata-org/diadata/dia-pkg/exchange-scrapers/curvefi/token"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// curveFiContract       = "0x7002B727Ef8F5571Cb5F9D70D13DBEEb4dFAe9d1"
	curveFiLookBackBlocks = 6 * 60 * 24 * 20
	curveWsDial           = "ws://159.69.120.42:8546/"
	curveRestDial         = "http://159.69.120.42:8545/"
)

type CurveCoin struct {
	Symbol   string
	Decimals uint8
	Address  string
	Name     string
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
	for key := range p.pools {
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

func NewCurveFIScraper(exchange dia.Exchange, scrape bool) *CurveFIScraper {
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

	err = scraper.loadPoolsAndCoins()
	if err != nil {
		log.Error(err)
	}

	if scrape {
		go scraper.mainLoop()
	}
	return scraper
}

func (scraper *CurveFIScraper) mainLoop() {
	scraper.run = true

	for _, pool := range scraper.pools.poolsAddressNoLock() {
		err := scraper.watchSwaps(pool)
		if err != nil {
			log.Error(err)
		}
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
					err := scraper.watchSwaps(p)
					if err != nil {
						log.Error(err)
					}
				}
			}
		}
	}()

	if scraper.run {
		if len(scraper.pairScrapers) == 0 {
			scraper.error = errors.New("no pairs to scrape provided")
			log.Error(scraper.error.Error())
		}
	}

	time.Sleep(10 * time.Second)

	if scraper.error == nil {
		scraper.error = errors.New("main loop terminated by Close()")
	}
	scraper.cleanup(nil)
}

func (scraper *CurveFIScraper) watchNewPools() {
	contract, err := curvefi2.NewCurvefiFilterer(scraper.contract, scraper.WsClient)
	if err != nil {
		log.Error(err)
	}
	sink := make(chan *curvefi2.CurvefiPoolAdded)

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
			case err = <-sub.Err():
				if err != nil {
					log.Error(err)
				}
				subscribed = false
				if scraper.run {
					scraper.resubscribe <- "NEW_POOLS"
				}
			case vLog := <-sink:

				if _, ok := scraper.pools.getPool(vLog.Pool.Hex()); !ok {
					err = scraper.loadPoolData(vLog.Pool.Hex())
					if err != nil {
						log.Error(err)
					}
					err = scraper.watchSwaps(vLog.Pool.Hex())
					if err != nil {
						log.Error(err)
					}
				}
			}
		}
	}()

}

func (scraper *CurveFIScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{}, nil
}

// contract.poolList.map(contract.GetPoolCoins(pool).)
func (scraper *CurveFIScraper) loadPoolsAndCoins() error {
	contract, err := curvefi2.NewCurvefiCaller(scraper.contract, scraper.RestClient)
	if err != nil {
		log.Error(err)
	}
	poolCount, err := contract.PoolCount(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}
	for i := 0; i < int(poolCount.Int64()); i++ {
		var pool common.Address
		pool, err = contract.PoolList(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			log.Error(err)
		}

		err = scraper.loadPoolData(pool.Hex())
		if err != nil {
			return err
		}

	}
	return err
}

func (scraper *CurveFIScraper) loadPoolData(pool string) error {
	contract, err := curvefi2.NewCurvefiCaller(scraper.contract, scraper.RestClient)
	if err != nil {
		log.Error(err)
	}

	poolCoinsMap := make(map[int]*CurveCoin)

	poolCoins, err := contract.GetPoolCoins(&bind.CallOpts{}, common.HexToAddress(pool))
	if err != nil {
		log.Error(err)
	}

	for cIdx, c := range poolCoins.Coins {
		var coinCaller *token2.TokenCaller
		var symbol string
		var decimals *big.Int
		var name string
		coinCaller, err = token2.NewTokenCaller(c, scraper.RestClient)
		if err != nil {
			log.Error(err)
			continue
		}
		symbol, err = coinCaller.Symbol(&bind.CallOpts{})
		if err != nil {
			log.Error(err, c.Hex())
			continue
		}
		decimals, err = coinCaller.Decimals(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
			continue
		}
		name, err = coinCaller.Name(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
			continue
		}

		poolCoinsMap[cIdx] = &CurveCoin{
			Symbol:   symbol,
			Decimals: uint8(decimals.Uint64()),
			Name:     name,
			Address:  c.String(),
		}
		scraper.curveCoins[c.Hex()] = &CurveCoin{
			Symbol:   symbol,
			Decimals: uint8(decimals.Uint64()),
		}

		scraper.pools.setPool(pool, poolCoinsMap)

	}
	return err
}

func (scraper *CurveFIScraper) processSwap(pool string, swp *curvepool2.CurvepoolTokenExchange) {

	foreignName, volume, price, baseToken, quoteToken, err := scraper.getSwapDataCurve(pool, swp)
	if err != nil {
		log.Error(err)
	}
	timestamp := time.Now().Unix()

	if pairScraper, ok := scraper.pairScrapers[foreignName]; ok {

		trade := &dia.Trade{
			Symbol:         pairScraper.pair.Symbol,
			Pair:           foreignName,
			BaseToken:      baseToken,
			QuoteToken:     quoteToken,
			Price:          price,
			Volume:         volume,
			Time:           time.Unix(timestamp, 0),
			ForeignTradeID: swp.Raw.TxHash.Hex() + "-" + fmt.Sprint(swp.Raw.Index),
			Source:         scraper.exchangeName,
			VerifiedPair:   true,
		}
		log.Infoln("Got Trade  ", trade)

		scraper.chanTrades <- trade
	}
}

func (scraper *CurveFIScraper) watchSwaps(pool string) error {

	filterer, err := curvepool2.NewCurvepoolFilterer(common.HexToAddress(pool), scraper.WsClient)
	if err != nil {
		log.Fatal(err)
	}
	sink := make(chan *curvepool2.CurvepoolTokenExchange)

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
			case err = <-sub.Err():
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
func (scraper *CurveFIScraper) getSwapDataCurve(pool string, s *curvepool2.CurvepoolTokenExchange) (foreignName string, volume float64, price float64, baseToken, quoteToken dia.Asset, err error) {

	// fromToken, _ := scraper.curveCoins[s.TokenSold.Hex()]
	// toToken, _ := scraper.curveCoins[s.TokenBought.Hex()]

	fromToken, ok := scraper.pools.getPoolCoin(pool, int(s.SoldId.Int64()))
	if !ok {
		err = fmt.Errorf("token not found: " + pool + "-" + s.SoldId.String())
	}
	baseToken = dia.Asset{
		Name:       fromToken.Name,
		Address:    fromToken.Address,
		Symbol:     fromToken.Symbol,
		Blockchain: dia.ETHEREUM,
	}
	toToken, ok := scraper.pools.getPoolCoin(pool, int(s.BoughtId.Int64()))
	if !ok {
		err = fmt.Errorf("token not found: " + pool + "-" + s.SoldId.String())
	}

	quoteToken = dia.Asset{
		Name:       toToken.Name,
		Address:    toToken.Address,
		Symbol:     toToken.Symbol,
		Blockchain: dia.ETHEREUM,
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

func (scraper *CurveFIScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {

	pairSet := make(map[string]struct{})
	for _, p1 := range scraper.curveCoins {
		for _, p2 := range scraper.curveCoins {
			token1 := p1
			token2 := p2
			if token1 != token2 {

				foreignName := token1.Symbol + "-" + token2.Symbol
				if _, ok := pairSet[foreignName]; !ok {
					pairs = append(pairs, dia.ExchangePair{
						Symbol:      token1.Symbol,
						ForeignName: foreignName,
						Exchange:    scraper.exchangeName,
					})
					pairSet[foreignName] = struct{}{}
				}

				foreignName = token2.Symbol + "-" + token1.Symbol
				if _, ok := pairSet[foreignName]; !ok {
					pairs = append(pairs, dia.ExchangePair{
						Symbol:      token2.Symbol,
						ForeignName: foreignName,
						Exchange:    scraper.exchangeName,
					})
					pairSet[foreignName] = struct{}{}
				}

			}
		}
	}
	return
}

func (scraper *CurveFIScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
}

func (scraper *CurveFIScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
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
	pair   dia.ExchangePair
	closed bool
}

func (pairScraper *CurveFIPairScraper) Pair() dia.ExchangePair {
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
