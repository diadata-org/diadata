package scrapers

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefi"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefi/curvepool"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefifactory"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefimeta"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	curveRestDialEth      = ""
	curveWsDialEth        = ""
	curveRestDialFantom   = ""
	curveWsDialFantom     = ""
	curveRestDialMoonbeam = ""
	curveWsDialMoonbeam   = ""
	curveRestDialPolygon  = ""
	curveWsDialPolygon    = ""
	curveRestDialArbitrum = ""
	curveWsDialArbitrum   = ""
)

type CurveCoin struct {
	Symbol   string
	Decimals uint8
	Address  string
	Name     string
}

func assetToCoin(a dia.Asset) (c *CurveCoin) {
	c = &CurveCoin{}
	c.Address = a.Address
	c.Decimals = a.Decimals
	c.Name = a.Name
	c.Symbol = a.Symbol
	return
}

type curveRegistry struct {
	Address common.Address
	Type    int
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

// CurveFIScraper is a curve finance scraper on a specific blockchain.
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

	WsClient             *ethclient.Client
	RestClient           *ethclient.Client
	relDB                *models.RelDB
	curveCoins           map[string]*CurveCoin
	resubscribe          chan string
	pools                *Pools
	registriesUnderlying []curveRegistry
	screenPools          bool
}

// makeCurvefiScraper returns a curve finance scraper as used in NewCurvefiScraper.
func makeCurvefiScraper(exchange dia.Exchange, restDial string, wsDial string) *CurveFIScraper {
	var (
		restClient, wsClient *ethclient.Client
		err                  error
		scraper              *CurveFIScraper
	)

	log.Infof("Init rest and ws client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", restDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}
	wsClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", wsDial))
	if err != nil {
		log.Fatal("init ws client: ", err)
	}

	scraper = &CurveFIScraper{
		exchangeName:   exchange.Name,
		RestClient:     restClient,
		WsClient:       wsClient,
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

	scraper.relDB, err = models.NewPostgresDataStore()
	if err != nil {
		log.Fatal("new postgres datastore: ", err)
	}
	// Only include pools with (minimum) liquidity bigger than given env var.
	liquidityThreshold, err := strconv.ParseFloat(utils.Getenv("LIQUIDITY_THRESHOLD", "0"), 64)
	if err != nil {
		liquidityThreshold = float64(0)
		log.Warnf("parse liquidity threshold:  %v. Set to default %v", err, liquidityThreshold)
	}

	scraper.loadPools(liquidityThreshold)
	return scraper
}

func NewCurveFIScraper(exchange dia.Exchange, scrape bool) *CurveFIScraper {

	var scraper *CurveFIScraper

	switch exchange.Name {
	case dia.CurveFIExchange:
		scraper = makeCurvefiScraper(exchange, curveRestDialEth, curveWsDialEth)
		basePoolRegistry := curveRegistry{Type: 1, Address: common.HexToAddress("0x90E00ACe148ca3b23Ac1bC8C240C2a7Dd9c2d7f5")}
		metaPoolRegistry := curveRegistry{Type: 2, Address: common.HexToAddress("0xB9fC157394Af804a3578134A6585C0dc9cc990d4")}
		scraper.registriesUnderlying = []curveRegistry{metaPoolRegistry, basePoolRegistry}
		scraper.screenPools = true

	case dia.CurveFIExchangeFantom:
		scraper = makeCurvefiScraper(exchange, curveRestDialFantom, curveWsDialFantom)
		stableSwapFactory := curveRegistry{Type: 2, Address: common.HexToAddress("0x686d67265703d1f124c45e33d47d794c566889ba")}
		scraper.registriesUnderlying = []curveRegistry{stableSwapFactory}
		scraper.screenPools = false

	case dia.CurveFIExchangeMoonbeam:
		scraper = makeCurvefiScraper(exchange, curveRestDialMoonbeam, curveWsDialMoonbeam)
		stableSwapFactory := curveRegistry{Type: 2, Address: common.HexToAddress("0x4244eB811D6e0Ef302326675207A95113dB4E1F8")}
		scraper.registriesUnderlying = []curveRegistry{stableSwapFactory}
		scraper.screenPools = false

	case dia.CurveFIExchangePolygon:
		scraper = makeCurvefiScraper(exchange, curveRestDialPolygon, curveWsDialPolygon)
		stableSwapFactory := curveRegistry{Type: 2, Address: common.HexToAddress("0x722272D36ef0Da72FF51c5A65Db7b870E2e8D4ee")}
		scraper.registriesUnderlying = []curveRegistry{stableSwapFactory}
		scraper.screenPools = false
	case dia.CurveFIExchangeArbitrum:
		scraper = makeCurvefiScraper(exchange, curveRestDialArbitrum, curveWsDialArbitrum)
		stableSwapFactory := curveRegistry{Type: 2, Address: common.HexToAddress("0xb17b674D9c5CB2e441F8e196a2f048A81355d031")}
		scraper.registriesUnderlying = []curveRegistry{stableSwapFactory}
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
			log.Error("watchSwaps: ", err)
		}
	}

	go func() {
		for scraper.run {
			p := <-scraper.resubscribe
			log.Info("resub to p: ", p)

			if scraper.run {
				log.Info("resubscribe to swaps from Pool: " + p)
				err := scraper.watchSwaps(p)
				if err != nil {
					log.Error("watchSwaps in resubscribe: ", err)
				}

			}
		}
	}()

	if scraper.run {
		if len(scraper.pools.pools) == 0 {
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

	filtererV2, err := curvefifactory.NewCurvefifactoryFilterer(common.HexToAddress(pool), scraper.WsClient)
	if err != nil {
		log.Fatal(err)
	}
	sinkV2 := make(chan *curvefifactory.CurvefifactoryTokenExchange)

	header, err := scraper.RestClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	startblock := header.Number.Uint64() - uint64(20)

	sub, err := filterer.WatchTokenExchange(&bind.WatchOpts{Start: &startblock}, sink, nil)
	if err != nil {
		log.Error("WatchTokenExchange: ", err)
	}

	subV2, err := filtererV2.WatchTokenExchange(&bind.WatchOpts{Start: &startblock}, sinkV2, nil)
	if err != nil {
		log.Error("WatchTokenExchange: ", err)
	}

	sinkUnderlying := make(chan *curvepool.CurvepoolTokenExchangeUnderlying)
	subUnderlying, err := filterer.WatchTokenExchangeUnderlying(&bind.WatchOpts{Start: &startblock}, sinkUnderlying, nil)
	if err != nil {
		log.Error("WatchTokenExchangeUnderlying: ", err)
	}

	go func() {
		fmt.Println("Curvefi Subscribed to pool: " + pool)
		defer fmt.Printf("Curvefi UnSubscribed to pool %s with error: %v", pool, err)
		defer sub.Unsubscribe()
		defer subV2.Unsubscribe()
		defer subUnderlying.Unsubscribe()
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
				}
				log.Warn("subscription error: ", err)
			case swp := <-sink:
				log.Info("got swap V1")
				scraper.processSwap(pool, swp)

			case swpV2 := <-sinkV2:
				log.Info("got swap V2")
				scraper.processSwap(pool, swpV2)
			case swpUnderlying := <-sinkUnderlying:
				log.Warn("got underlying USDR swap: ", swpUnderlying)
				// Only fetch trades from USDR pool until we have parsing TokenExchangeUnderlying resolved.
				// if pool == common.HexToAddress("0xa138341185a9d0429b0021a11fb717b225e13e1f").Hex() && Exchanges[scraper.exchangeName].BlockChain.Name == dia.POLYGON {
				// 	scraper.processSwap(pool, swpUnderlying)
				// }
			}
		}
	}()

	return err

}

func (scraper *CurveFIScraper) processSwap(pool string, swp interface{}) {
	var (
		foreignName    string
		volume         float64
		price          float64
		baseToken      dia.Asset
		quoteToken     dia.Asset
		foreignTradeID string
		err            error
	)

	switch s := swp.(type) {
	case *curvepool.CurvepoolTokenExchange:
		foreignName, volume, price, baseToken, quoteToken, err = scraper.getSwapDataCurve(pool, s)
		if err != nil {
			log.Error("getSwapDataCurve: ", err)
			return
		}
		foreignTradeID = s.Raw.TxHash.Hex() + "-" + fmt.Sprint(s.Raw.Index)
	case *curvefifactory.CurvefifactoryTokenExchange:
		foreignName, volume, price, baseToken, quoteToken, err = scraper.getSwapDataCurve(pool, s)
		if err != nil {
			log.Error("getSwapDataCurve: ", err)
			return
		}
		foreignTradeID = s.Raw.TxHash.Hex() + "-" + fmt.Sprint(s.Raw.Index)
	case *curvepool.CurvepoolTokenExchangeUnderlying:
		foreignName, volume, price, baseToken, quoteToken, err = scraper.getSwapDataCurve(pool, s)
		if err != nil {
			log.Error("getSwapDataCurve: ", err)
			return
		}
		foreignTradeID = s.Raw.TxHash.Hex() + "-" + fmt.Sprint(s.Raw.Index)
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
		ForeignTradeID: foreignTradeID,
		PoolAddress:    pool,
		Source:         scraper.exchangeName,
		VerifiedPair:   true,
	}
	log.Infof("Got Trade in pool %s:\n %v", pool, trade)

	scraper.chanTrades <- trade

}

// getSwapDataCurve returns the foreign name, volume and price of a swap
func (scraper *CurveFIScraper) getSwapDataCurve(pool string, swp interface{}) (
	foreignName string,
	volume float64,
	price float64,
	baseToken,
	quoteToken dia.Asset,
	err error,
) {
	var (
		fromToken *CurveCoin
		toToken   *CurveCoin
		amountIn  float64
		amountOut float64
		ok        bool
	)

	switch s := swp.(type) {
	case *curvepool.CurvepoolTokenExchange:
		fromToken, ok = scraper.pools.getPoolCoin(pool, int(s.SoldId.Int64()))
		if !ok {
			err = fmt.Errorf("token not found: " + pool + "-" + s.SoldId.String())
		}
		toToken, ok = scraper.pools.getPoolCoin(pool, int(s.BoughtId.Int64()))
		if !ok {
			err = fmt.Errorf("token not found: " + pool + "-" + s.SoldId.String())
		}
		amountIn, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(s.TokensSold), new(big.Float).SetFloat64(math.Pow10(int(fromToken.Decimals)))).Float64()
		amountOut, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(s.TokensBought), new(big.Float).SetFloat64(math.Pow10(int(toToken.Decimals)))).Float64()
	case *curvefifactory.CurvefifactoryTokenExchange:
		fromToken, ok = scraper.pools.getPoolCoin(pool, int(s.SoldId.Int64()))
		if !ok {
			err = fmt.Errorf("token not found: " + pool + "-" + s.SoldId.String())
		}
		toToken, ok = scraper.pools.getPoolCoin(pool, int(s.BoughtId.Int64()))
		if !ok {
			err = fmt.Errorf("token not found: " + pool + "-" + s.SoldId.String())
		}
		if toToken == nil || fromToken == nil {
			err = errors.New("token not available, skip trade.")
			return
		}
		amountIn, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(s.TokensSold), new(big.Float).SetFloat64(math.Pow10(int(fromToken.Decimals)))).Float64()
		amountOut, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(s.TokensBought), new(big.Float).SetFloat64(math.Pow10(int(toToken.Decimals)))).Float64()
	case *curvepool.CurvepoolTokenExchangeUnderlying:
		log.Info("Got TokenExchangeUnderlying in pool: ", pool)
		var (
			underlyingCoins [8]common.Address
			errCoins        error
		)

		for _, registry := range scraper.registriesUnderlying {
			if registry.Type == 2 {
				metaRegistryContract, errMeta := curvefimeta.NewCurvefimetaCaller(registry.Address, scraper.RestClient)
				if errMeta != nil {
					log.Error("NewCurvefiCaller: ", errMeta)
				}
				// Get underlying coins from on-chain. soldID and boughtID are referring to these.
				underlyingCoins, errCoins = metaRegistryContract.GetUnderlyingCoins(&bind.CallOpts{}, common.HexToAddress(pool))
				if errCoins != nil || (underlyingCoins[int(s.SoldId.Int64())] == (common.Address{}) && underlyingCoins[int(s.BoughtId.Int64())] == (common.Address{})) {
					continue
				} else {
					log.Warn("TokenExchangeUnderlying from meta type pool")
					log.Warnf("bought id and sold id: %v -- %v", s.BoughtId.Int64(), s.SoldId.Int64())
					log.Warnf("tokens bought and tokens sold: %v -- %v ", s.TokensBought, s.TokensSold)
					break
				}
			}
			if registry.Type == 1 {
				basepoolRegistryContract, errBase := curvefi.NewCurvefiCaller(registry.Address, scraper.RestClient)
				if errBase != nil {
					log.Error("NewCurvefiCaller: ", errBase)
				}
				// Get underlying coins from on-chain. soldID and boughtID are referring to these.
				underlyingCoins, errCoins = basepoolRegistryContract.GetUnderlyingCoins(&bind.CallOpts{}, common.HexToAddress(pool))
				if errCoins != nil || (underlyingCoins[int(s.SoldId.Int64())] == (common.Address{}) && underlyingCoins[int(s.BoughtId.Int64())] == (common.Address{})) {
					continue
				} else {
					log.Warn("TokenExchangeUnderlying from base type pool")
					log.Warnf("bought id and sold id: %v -- %v", s.BoughtId.Int64(), s.SoldId.Int64())
					log.Warnf("tokens bought and tokens sold: %v -- %v ", s.TokensBought, s.TokensSold)
					break
				}
			}
		}

		fromTokenAddress := underlyingCoins[int(s.SoldId.Int64())]
		toTokenAddress := underlyingCoins[int(s.BoughtId.Int64())]

		fromTokenAsset, errToken := scraper.relDB.GetAsset(fromTokenAddress.Hex(), Exchanges[scraper.exchangeName].BlockChain.Name)
		if errToken != nil {
			err = errToken
			return
		}
		toTokenAsset, errToken := scraper.relDB.GetAsset(toTokenAddress.Hex(), Exchanges[scraper.exchangeName].BlockChain.Name)
		if errToken != nil {
			err = errToken
			return
		}
		log.Infof("fromToken address -- token: %s -- %v", fromTokenAddress.Hex(), fromTokenAsset)
		log.Infof("toToken address -- token: %s -- %v", toTokenAddress.Hex(), toTokenAsset)
		fromToken = assetToCoin(fromTokenAsset)
		toToken = assetToCoin(toTokenAsset)

		amountIn, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(s.TokensSold), new(big.Float).SetFloat64(math.Pow10(int(fromToken.Decimals)))).Float64()
		amountOut, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(s.TokensBought), new(big.Float).SetFloat64(math.Pow10(int(toToken.Decimals)))).Float64()
	}

	baseToken = dia.Asset{
		Name:       fromToken.Name,
		Address:    fromToken.Address,
		Symbol:     fromToken.Symbol,
		Blockchain: Exchanges[scraper.exchangeName].BlockChain.Name,
	}
	quoteToken = dia.Asset{
		Name:       toToken.Name,
		Address:    toToken.Address,
		Symbol:     toToken.Symbol,
		Blockchain: Exchanges[scraper.exchangeName].BlockChain.Name,
	}

	volume = amountOut
	price = amountIn / amountOut

	foreignName = toToken.Symbol + "-" + fromToken.Symbol

	return
}

func (scraper *CurveFIScraper) loadPools(liquidityThreshold float64) {

	pools, err := scraper.relDB.GetAllPoolsExchange(scraper.exchangeName, liquidityThreshold)
	if err != nil {
		log.Fatal("fetch pools: ", err)
	}
	log.Infof("found %v pools to subscribe: ", len(pools))

	for _, pool := range pools {
		var poolCoinsMap = make(map[int]*CurveCoin)
		for _, av := range pool.Assetvolumes {
			poolCoinsMap[int(av.Index)] = &CurveCoin{
				Symbol:   av.Asset.Symbol,
				Decimals: av.Asset.Decimals,
				Name:     av.Asset.Name,
				Address:  av.Asset.Address,
			}
			scraper.curveCoins[av.Asset.Address] = &CurveCoin{
				Symbol:   av.Asset.Symbol,
				Decimals: av.Asset.Decimals,
			}
		}
		scraper.pools.setPool(pool.Address, poolCoinsMap)
	}
}

func (scraper *CurveFIScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {

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
