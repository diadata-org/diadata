package scrapers

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefi"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefi/curvepool"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefi/token"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefimeta"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	curveFiMetaPoolsFactory = "0xB9fC157394Af804a3578134A6585C0dc9cc990d4"
	curveFiLookBackBlocks   = 6 * 60 * 24 * 20
	curveWsDial             = ""
	curveRestDial           = ""
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

	log.Infof("Init rest and ws client for %s.", exchange.BlockChain.Name)
	restClient, err := ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", curveRestDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}
	wsClient, err := ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", curveWsDial))
	if err != nil {
		log.Fatal("init ws client: ", err)
	}

	scraper.RestClient = restClient
	scraper.WsClient = wsClient

	// Load meta pools.
	err = scraper.loadPoolsAndCoins(common.HexToAddress(curveFiMetaPoolsFactory))
	if err != nil {
		log.Error("loadPoolsAndCoins: ", err)
	}
	log.Infof("loaded meta pools. Now %v pools.", len(scraper.pools.pools))

	// Load standard pools.
	err = scraper.loadPoolsAndCoins(scraper.contract)
	if err != nil {
		log.Error("loadPoolsAndCoins: ", err)
	}
	log.Infof("loaded standard pools. Now %v pools.", len(scraper.pools.pools))

	// Load crypto pools.
	// err = scraper.loadPoolsAndCoins(common.HexToAddress("0xF18056Bbd320E96A48e3Fbf8bC061322531aac99"))
	// if err != nil {
	// 	log.Error("loadPoolsAndCoins: ", err)
	// }
	// log.Infof("loaded crypto pools. Now %v pools.", len(scraper.pools.pools))

	// Load cryptoswap pools.
	err = scraper.loadPoolsAndCoins(common.HexToAddress("0x8F942C20D02bEfc377D41445793068908E2250D0"))
	if err != nil {
		log.Error("loadPoolsAndCoins: ", err)
	}
	log.Infof("loaded cryptoswap pools. Now %v pools.", len(scraper.pools.pools))

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
			log.Error("watchSwaps: ", err)
		}
	}
	scraper.watchNewPools()

	go func() {
		for scraper.run {
			p := <-scraper.resubscribe
			log.Info("resub to p: ", p)

			if scraper.run {
				if p == "NEW_POOLS" {
					log.Info("resubscribe to new pools")
					scraper.watchNewPools()
				} else {
					log.Info("resubscribe to swaps from Pool: " + p)
					err := scraper.watchSwaps(p)
					if err != nil {
						log.Error("watchSwaps in resubscribe: ", err)
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
	startblock := header.Number.Uint64() - uint64(20)

	sub, err := filterer.WatchTokenExchange(&bind.WatchOpts{Start: &startblock}, sink, nil)
	if err != nil {
		log.Error("WatchTokenExchange: ", err)
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
					log.Error("sub error: ", err)
				}
				subscribed = false
				if scraper.run {
					log.Warn("resubscribe pool: ", pool)
					scraper.resubscribe <- pool
					log.Info("scraper: ", scraper)
				}
				log.Warn("subscription error: ", err)
			case swp := <-sink:
				scraper.processSwap(pool, swp)
			}
		}
	}()

	return err

}

func (scraper *CurveFIScraper) processSwap(pool string, swp *curvepool.CurvepoolTokenExchange) {

	foreignName, volume, price, baseToken, quoteToken, err := scraper.getSwapDataCurve(pool, swp)
	if err != nil {
		log.Error("getSwapDataCurve: ", err)
	}
	timestamp := time.Now().Unix()

	trade := &dia.Trade{
		Symbol:         quoteToken.Symbol,
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
	// log.Infof("Got Trade in pool %s:\n %v", pool, trade)

	scraper.chanTrades <- trade

}

// getSwapDataCurve returns the foreign name, volume and price of a swap
func (scraper *CurveFIScraper) getSwapDataCurve(pool string, s *curvepool.CurvepoolTokenExchange) (foreignName string, volume float64, price float64, baseToken, quoteToken dia.Asset, err error) {

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

func (scraper *CurveFIScraper) watchNewPools() {
	contract, err := curvefi.NewCurvefiFilterer(scraper.contract, scraper.WsClient)
	if err != nil {
		log.Error("NewCurvefiFilterer: ", err)
	}
	sink := make(chan *curvefi.CurvefiPoolAdded)

	header, err := scraper.RestClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	startblock := header.Number.Uint64() - uint64(curveFiLookBackBlocks)

	sub, err := contract.WatchPoolAdded(&bind.WatchOpts{Start: &startblock}, sink, nil)
	if err != nil {
		log.Error("WatchPoolAdded: ", err)
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
					log.Error("subscription error in new pools: ", err)
				}
				subscribed = false
				if scraper.run {
					scraper.resubscribe <- "NEW_POOLS"
				}
			case vLog := <-sink:

				if _, ok := scraper.pools.getPool(vLog.Pool.Hex()); !ok {
					err = scraper.loadPoolData(vLog.Pool.Hex(), scraper.contract)
					if err != nil {
						log.Error("loadPoolData in new pools: ", err)
					}
					err = scraper.watchSwaps(vLog.Pool.Hex())
					if err != nil {
						log.Error("watchSwaps in new pools: ", err)
					}
				}
			}
		}
	}()

}

// contract.poolList.map(contract.GetPoolCoins(pool).)
func (scraper *CurveFIScraper) loadPoolsAndCoins(factoryAddress common.Address) (err error) {

	if factoryAddress == common.HexToAddress(curveFiMetaPoolsFactory) || factoryAddress == common.HexToAddress("0xF18056Bbd320E96A48e3Fbf8bC061322531aac99") {
		log.Info("load meta pools...")
		var contract *curvefimeta.CurvefimetaCaller
		var poolCount *big.Int
		contract, err = curvefimeta.NewCurvefimetaCaller(factoryAddress, scraper.RestClient)
		if err != nil {
			log.Error("NewCurvefiCaller: ", err)
		}
		poolCount, err = contract.PoolCount(&bind.CallOpts{})
		if err != nil {
			log.Error("PoolCount: ", err)
		}
		log.Info("poolCount: ", int(poolCount.Int64()))
		for i := 0; i < int(poolCount.Int64()); i++ {
			var pool common.Address
			pool, err = contract.PoolList(&bind.CallOpts{}, big.NewInt(int64(i)))
			if err != nil {
				log.Error("PoolList: ", err)
			}
			log.Info("pool: ", pool.Hex())

			err = scraper.loadPoolData(pool.Hex(), factoryAddress)
			if err != nil {
				log.Info("load pool data: ", err)
				// return err
			}
		}
	} else {
		var contract *curvefi.CurvefiCaller
		var poolCount *big.Int
		contract, err = curvefi.NewCurvefiCaller(factoryAddress, scraper.RestClient)
		if err != nil {
			log.Error("NewCurvefiCaller: ", err)
		}
		poolCount, err = contract.PoolCount(&bind.CallOpts{})
		if err != nil {
			log.Error("PoolCount: ", err)
		}
		log.Info("poolCount: ", int(poolCount.Int64()))
		for i := 0; i < int(poolCount.Int64()); i++ {
			var pool common.Address
			pool, err = contract.PoolList(&bind.CallOpts{}, big.NewInt(int64(i)))
			if err != nil {
				log.Error("PoolList: ", err)
			}
			log.Info("pool: ", pool.Hex())

			err = scraper.loadPoolData(pool.Hex(), factoryAddress)
			if err != nil {
				log.Info("load pool data: ", err)
				// return err
			}
		}
	}
	return err
}

func (scraper *CurveFIScraper) loadPoolData(pool string, factoryContract common.Address) error {
	var poolCoins [8]common.Address
	var poolCoinsMap = make(map[int]*CurveCoin)

	if factoryContract == common.HexToAddress(curveFiMetaPoolsFactory) || factoryContract == common.HexToAddress("0xF18056Bbd320E96A48e3Fbf8bC061322531aac99") {
		contract, err := curvefimeta.NewCurvefimetaCaller(factoryContract, scraper.RestClient)

		if err != nil {
			log.Error("loadPoolData - NewCurvefiCaller: ", err)
		}

		aux, err := contract.GetCoins(&bind.CallOpts{}, common.HexToAddress(pool))
		if err != nil {
			log.Error("loadPoolData - GetCoins: ", err)
		}
		// GetCoins on meta pools returns [4]common.Address instead of [8]common.Address for standard pools.
		for i, item := range aux {
			poolCoins[i] = item
		}
		log.Info("pool coins: ", poolCoins)
	} else {
		contract, err := curvefi.NewCurvefiCaller(factoryContract, scraper.RestClient)
		if err != nil {
			log.Error("loadPoolData - NewCurvefiCaller: ", err)
		}

		poolCoins, err = contract.GetCoins(&bind.CallOpts{}, common.HexToAddress(pool))
		if err != nil {
			log.Error("loadPoolData - GetCoins: ", err)
		}
		log.Info("pool coins: ", poolCoins)
		poolName, err := contract.GetPoolName(&bind.CallOpts{}, common.HexToAddress(pool))
		if err != nil {
			log.Error("loadPoolData - GetPoolName: ", err)
		}
		log.Info("pool name: ", poolName)
	}

	var err error
	for cIdx, c := range poolCoins {
		var coinCaller *token.TokenCaller
		var symbol string
		var decimals *big.Int
		var name string
		if c == common.HexToAddress("0x0000000000000000000000000000000000000000") {
			continue
		} else if c == common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {
			symbol = "ETH"
			decimals = big.NewInt(int64(18))
			name = "Ether"
			c = common.HexToAddress("0x0000000000000000000000000000000000000000")
		} else {
			coinCaller, err = token.NewTokenCaller(c, scraper.RestClient)
			if err != nil {
				log.Error("NewTokenCaller: ", err)
				continue
			}
			symbol, err = coinCaller.Symbol(&bind.CallOpts{})
			if err != nil {
				log.Error("Symbol: ", err, c.Hex())
				continue
			}
			decimals, err = coinCaller.Decimals(&bind.CallOpts{})
			if err != nil {
				log.Error("Decimals: ", err)
				continue
			}
			name, err = coinCaller.Name(&bind.CallOpts{})
			if err != nil {
				log.Error("Name: ", err)
				continue
			}
		}
		log.Info(symbol, " ", decimals, " ", "'", name, "'", " ", c)

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

func (scraper *CurveFIScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {

	err = scraper.loadPoolsAndCoins(scraper.contract)
	pools := scraper.pools.pools
	for _, value := range pools {
		// In case a pool has more than 2 assets, consider all possible combinations.

		pairSet := make(map[string]struct{})
		for _, p1 := range value {
			for _, p2 := range value {
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
	}

	return
}

func (scraper *CurveFIScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
}

func (scraper *CurveFIScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{}, nil
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
