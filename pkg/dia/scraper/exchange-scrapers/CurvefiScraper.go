package scrapers

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"
	"reflect"
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
	"github.com/ethereum/go-ethereum/accounts/abi"
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
func makeCurvefiScraper(exchange dia.Exchange, restDial string, wsDial string, relDB *models.RelDB) *CurveFIScraper {
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
		relDB:          relDB,
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

	// Only include pools with (minimum) liquidity bigger than given env var.
	liquidityThreshold, err := strconv.ParseFloat(utils.Getenv("LIQUIDITY_THRESHOLD", "0"), 64)
	if err != nil {
		liquidityThreshold = float64(0)
		log.Warnf("parse liquidity threshold:  %v. Set to default %v", err, liquidityThreshold)
	}

	// Only include pools with (minimum) liquidity USD value bigger than given env var.
	liquidityThresholdUSD, err := strconv.ParseFloat(utils.Getenv("LIQUIDITY_THRESHOLD_USD", "0"), 64)
	if err != nil {
		liquidityThresholdUSD = float64(0)
		log.Warnf("parse liquidity threshold:  %v. Set to default %v", err, liquidityThresholdUSD)
	}

	scraper.loadPools(liquidityThreshold, liquidityThresholdUSD)

	return scraper
}

func NewCurveFIScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *CurveFIScraper {

	var scraper *CurveFIScraper

	switch exchange.Name {
	case dia.CurveFIExchange:
		scraper = makeCurvefiScraper(exchange, curveRestDialEth, curveWsDialEth, relDB)
		basePoolRegistry := curveRegistry{Type: 1, Address: common.HexToAddress(exchange.Contract)}
		cryptoswapPools := curveRegistry{Type: 1, Address: common.HexToAddress("0x8F942C20D02bEfc377D41445793068908E2250D0")}
		metaPoolRegistry := curveRegistry{Type: 2, Address: common.HexToAddress("0xB9fC157394Af804a3578134A6585C0dc9cc990d4")}
		factoryPools := curveRegistry{Type: 3, Address: common.HexToAddress("0xF18056Bbd320E96A48e3Fbf8bC061322531aac99")}
		factory2Pools := curveRegistry{Type: 3, Address: common.HexToAddress("0x4F8846Ae9380B90d2E71D5e3D042dff3E7ebb40d")}
		scraper.registriesUnderlying = []curveRegistry{factoryPools, factory2Pools, metaPoolRegistry, basePoolRegistry, cryptoswapPools}
		scraper.screenPools = true

	case dia.CurveFIExchangeFantom:
		scraper = makeCurvefiScraper(exchange, curveRestDialFantom, curveWsDialFantom, relDB)
		stableSwapFactory := curveRegistry{Type: 2, Address: common.HexToAddress("0x686d67265703d1f124c45e33d47d794c566889ba")}
		scraper.registriesUnderlying = []curveRegistry{stableSwapFactory}
		scraper.screenPools = false

	case dia.CurveFIExchangeMoonbeam:
		scraper = makeCurvefiScraper(exchange, curveRestDialMoonbeam, curveWsDialMoonbeam, relDB)
		stableSwapFactory := curveRegistry{Type: 2, Address: common.HexToAddress("0x4244eB811D6e0Ef302326675207A95113dB4E1F8")}
		scraper.registriesUnderlying = []curveRegistry{stableSwapFactory}
		scraper.screenPools = false

	case dia.CurveFIExchangePolygon:
		scraper = makeCurvefiScraper(exchange, curveRestDialPolygon, curveWsDialPolygon, relDB)
		stableSwapFactory := curveRegistry{Type: 2, Address: common.HexToAddress("0x722272D36ef0Da72FF51c5A65Db7b870E2e8D4ee")}
		scraper.registriesUnderlying = []curveRegistry{stableSwapFactory}
		scraper.screenPools = false
	case dia.CurveFIExchangeArbitrum:
		scraper = makeCurvefiScraper(exchange, curveRestDialArbitrum, curveWsDialArbitrum, relDB)
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
				log.Warn("got underlying swap: ", swpUnderlying)
				// Only fetch trades from USDR pool until we have parsing TokenExchangeUnderlying resolved.
				// if pool == common.HexToAddress("0xa138341185a9d0429b0021a11fb717b225e13e1f").Hex() && Exchanges[scraper.exchangeName].BlockChain.Name == dia.POLYGON {
				scraper.processSwap(pool, swpUnderlying)
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
			isMetaPool          bool
			underlyingCoins     [8]common.Address
			underlyingCoinCount *big.Int
			metaCoinCount       *big.Int
			errNCoin            error
			errCoins            error
		)
		soldID := int(s.SoldId.Int64())
		boughtId := int(s.BoughtId.Int64())
		tokenSold := s.TokensSold
		tokenBought := s.TokensBought

		for _, registry := range scraper.registriesUnderlying {
			if registry.Type == 2 {
				metaRegistryContract, errMeta := curvefimeta.NewCurvefimetaCaller(registry.Address, scraper.RestClient)
				if errMeta != nil {
					log.Error("NewCurvefiCaller: ", errMeta)
				}
				// Get underlying coins from on-chain. soldID and boughtID are referring to these.
				underlyingCoins, errCoins = metaRegistryContract.GetUnderlyingCoins(&bind.CallOpts{}, common.HexToAddress(pool))
				if errCoins != nil || (underlyingCoins[soldID] == (common.Address{}) && underlyingCoins[boughtId] == (common.Address{})) {
					log.Warnf("Failed to call GetUnderlyingCoins: %v", errCoins)
					continue
				} else {
					log.Warn("TokenExchangeUnderlying from meta type pool")
					log.Warnf("bought id and sold id: %v -- %v", boughtId, soldID)
					log.Warnf("tokens bought and tokens sold: %v -- %v ", tokenBought, tokenSold)
					metaCoinCount, underlyingCoinCount, errNCoin = metaRegistryContract.GetMetaNCoins(&bind.CallOpts{}, common.HexToAddress(pool))
					if errNCoin != nil {
						log.Errorf("calling GetMetaNCoins: %v", errCoins)
						continue
					}
					log.Warnf("metaCoinCount: %v, underlyingCoinCount: %v", metaCoinCount, underlyingCoinCount)
					isMetaPool = true
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
				if errCoins != nil || (underlyingCoins[soldID] == (common.Address{}) && underlyingCoins[boughtId] == (common.Address{})) {
					continue
				} else {
					log.Warn("TokenExchangeUnderlying from base type pool")
					log.Warnf("bought id and sold id: %v -- %v", boughtId, soldID)
					log.Warnf("tokens bought and tokens sold: %v -- %v ", tokenBought, tokenSold)
					break
				}

			}
		}
		log.Warnf("meta pool : %v, soldID: %v", isMetaPool, soldID)
		if isMetaPool && soldID > 0 && boughtId == 0 {
			// This is the only case we need to look into the AddLiquidity event for finding the actual amount of sold token
			// as this event contains the meta pool token amount in this case!

			tokenAmount, errTokenAmounts := scraper.getTokenAmount(s.Raw.TxHash, common.HexToAddress(pool), soldID, metaCoinCount, underlyingCoinCount)
			if err != nil {
				log.Error("getting AddLiquidity event for tx: ", s.Raw.TxHash.Hex())
				err = errTokenAmounts
				return
			}
			log.Warnf("token Amount %v", tokenAmount)
			// Again we need to subtract MAX_COIN index from this value, (which is metaCoinCount minus one)
			tokenSold = tokenAmount
		}
		log.Infof("token sold: %v", tokenSold)

		fromTokenAddress := underlyingCoins[int(soldID)]
		toTokenAddress := underlyingCoins[int(boughtId)]

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

		amountIn, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(tokenSold), new(big.Float).SetFloat64(math.Pow10(int(fromToken.Decimals)))).Float64()
		amountOut, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(tokenBought), new(big.Float).SetFloat64(math.Pow10(int(toToken.Decimals)))).Float64()
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

func (scraper *CurveFIScraper) getTokenAmount(txHash common.Hash, poolAddress common.Address, soldID int, metaCoinCount, underlyingCoinCount *big.Int) (*big.Int, error) {
	// Here we need to get the event log manually as the len of the token_amounts array is unknown!
	// (fixed to 2 in the current curvepool contract, but there are cases when it's 3 like this pool:
	// https://etherscan.io/address/0xbebc44782c7db0a1a60cb6fe97d0b483032ff1c7)
	// We get the value using reflection.
	receipt, err := scraper.RestClient.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return nil, err
	}
	basePoolCoinCount := underlyingCoinCount.Int64() - 1 // We need to subtract that one meta pool coin here!
	abiAddLiquidityJson := `[{"name":"AddLiquidity","inputs":[{"type":"address","name":"provider","indexed":true},{"type":"uint256[%d]","name":"token_amounts","indexed":false},{"type":"uint256[%d]","name":"fees","indexed":false},{"type":"uint256","name":"invariant","indexed":false},{"type":"uint256","name":"token_supply","indexed":false}],"anonymous":false,"type":"event"}]`
	abiAddLiquidity, err := abi.JSON(strings.NewReader(fmt.Sprintf(abiAddLiquidityJson, basePoolCoinCount, basePoolCoinCount)))
	if err != nil {
		return nil, err
	}
	for _, log := range receipt.Logs {
		if len(log.Topics) < 2 || log.Topics[1] != poolAddress.Hash() {
			continue
		}
		valueMap := make(map[string]interface{})
		err := abiAddLiquidity.UnpackIntoMap(valueMap, "AddLiquidity", log.Data)
		if err != nil {
			continue
		}
		// Again we need to subtract MAX_COIN index from this value, (which is metaCoinCount minus one)
		tokenBoughtIdx := soldID - (int(metaCoinCount.Int64()) - 1)
		tokenAmount, ok := reflect.ValueOf(valueMap["token_amounts"]).Index(tokenBoughtIdx).Interface().(*big.Int)
		if !ok {
			return nil, errors.New("couldn't parse the AddLiquidity event")
		}
		return tokenAmount, nil
	}
	return nil, errors.New("AddLiquidity log couldn't be found")

}

func (scraper *CurveFIScraper) loadPools(liquidityThreshold float64, liquidityThresholdUSD float64) {

	pools, err := scraper.relDB.GetAllPoolsExchange(scraper.exchangeName, liquidityThreshold)
	if err != nil {
		log.Fatal("fetch pools: ", err)
	}
	log.Infof("found %v pools to subscribe: ", len(pools))

	blacklistedPools, err := getAddressesFromConfig("curve/blacklist_pools/" + scraper.exchangeName)
	if err != nil {
		log.Fatal("blacklisted pools: ", err)
	}
	var blacklistedPoolsString []string
	for _, bp := range blacklistedPools {
		blacklistedPoolsString = append(blacklistedPoolsString, bp.Hex())
	}
	log.Infof("remove %v blacklisted pools: %s", len(blacklistedPools), blacklistedPoolsString)

	lowerBoundCount := 0
	for _, pool := range pools {
		if utils.Contains(&blacklistedPoolsString, pool.Address) {
			continue
		}

		liquidity, lowerBound := pool.GetPoolLiquidityUSD()
		// Discard pool if complete USD liquidity is below threshold.
		if !lowerBound && liquidity < liquidityThresholdUSD {
			continue
		}
		if lowerBound {
			log.Info("lowerBound pool: ", pool.Address)
			lowerBoundCount++
		}

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
	log.Infof("Total number of %v pools to subscribe.", len(scraper.pools.pools))
	log.Infof("Number of lower bound pools: %v", lowerBoundCount)
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
