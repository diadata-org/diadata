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

	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/platypusfinance"
	platypusAssetABI "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/platypusfinance/asset"
	platypusPoolABI "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/platypusfinance/pool"
	platypusTokenABI "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/platypusfinance/token"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// 30 days of blocks (avgBlocktimeTime = 2.5)
	platypusLookBackBlocks  = (60 / 2.5) * 60 * 24 * 30
	platypusRestDialEth     = "https://api.avax.network/ext/bc/C/rpc"
	platypusWsDialEth       = "wss://api.avax.network/ext/bc/C/ws"
	platypusMasterRegV3Addr = "0x7125B4211357d7C3a90F796c956c12c681146EbB"
	platypusMasterRegV2Addr = "0x68c5f4374228BEEdFa078e77b5ed93C28a2f713E"
	platypusMasterRegV1Addr = "0xB0523f9F473812FB195Ee49BC7d2ab9873a98044"
)

type platypusRegistry struct {
	Address common.Address
	Version int
}

type PlatypusCoin struct {
	Symbol   string
	Decimals uint8
	Address  string
	Name     string
}

type PlatypusPools struct {
	pools     map[string]map[int]*PlatypusCoin
	poolsLock sync.RWMutex
}

func (p *PlatypusPools) setPool(k string, v map[int]*PlatypusCoin) {
	p.poolsLock.Lock()
	defer p.poolsLock.Unlock()
	p.pools[k] = v
}
func (p *PlatypusPools) getPool(k string) (map[int]*PlatypusCoin, bool) {
	p.poolsLock.RLock()
	defer p.poolsLock.RUnlock()
	r, ok := p.pools[k]
	return r, ok
}

func (p *PlatypusPools) poolsAddressNoLock() []string {
	p.poolsLock.RLock()
	defer p.poolsLock.RUnlock()
	var values []string
	for key := range p.pools {
		values = append(values, key)
	}
	return values
}

// PlatypusScraper The scraper object for Platypus Finance
type PlatypusScraper struct {
	exchangeName string

	// channels to signal events
	run          bool
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing
	resubscribe  chan string

	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers map[string]*PlatypusPairScraper
	chanTrades   chan *dia.Trade

	WsClient         *ethclient.Client
	RestClient       *ethclient.Client
	relDB            *models.RelDB
	platypusCoins    map[string]*PlatypusCoin
	pools            *PlatypusPools
	screenPools      bool
	basePoolRegistry platypusRegistry
}

// NewPlatypusScraper Returns a new exchange scraper
func NewPlatypusScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *PlatypusScraper {

	registries := []platypusRegistry{
		{Version: 3, Address: common.HexToAddress(platypusMasterRegV3Addr)},
		{Version: 2, Address: common.HexToAddress(platypusMasterRegV2Addr)},
		{Version: 1, Address: common.HexToAddress(platypusMasterRegV1Addr)},
	}

	log.Infof("init rest and ws client for %s", exchange.BlockChain.Name)
	restClient, err := ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", platypusRestDialEth))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}
	wsClient, err := ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", platypusWsDialEth))
	if err != nil {
		log.Fatal("init ws client: ", err)
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

	scraper := &PlatypusScraper{
		exchangeName:  exchange.Name,
		RestClient:    restClient,
		WsClient:      wsClient,
		initDone:      make(chan nothing),
		shutdown:      make(chan nothing),
		shutdownDone:  make(chan nothing),
		pairScrapers:  make(map[string]*PlatypusPairScraper),
		chanTrades:    make(chan *dia.Trade),
		platypusCoins: make(map[string]*PlatypusCoin),
		resubscribe:   make(chan string),
		pools: &PlatypusPools{
			pools: make(map[string]map[int]*PlatypusCoin),
		},
	}

	scraper.relDB = relDB
	// Load metadata from master registries
	for _, registry := range registries {
		err := scraper.loadPoolsAndCoins(registry, liquidityThreshold, liquidityThresholdUSD)
		if err != nil {
			log.Errorf("loadPoolsAndCoins error w %s registry (v%d): %s", registry.Address.Hex(), registry.Version, err)
		}
		log.Infof("metadata loaded, now scraper have %d pools data and %d coins", len(scraper.pools.pools), len(scraper.platypusCoins))
	}

	scraper.basePoolRegistry = platypusRegistry{Version: 3, Address: common.HexToAddress(platypusMasterRegV3Addr)}
	scraper.screenPools = true

	if scrape {
		go scraper.mainLoop()
	}
	return scraper
}

// Close Closes any existing API connections, as well as channels of
// pairScrapers from calls to ScrapePair
func (s *PlatypusScraper) Close() error {
	s.run = false
	for _, pairScraper := range s.pairScrapers {
		pairScraper.closed = true
	}
	s.WsClient.Close()
	s.RestClient.Close()

	close(s.shutdown)
	<-s.shutdownDone
	return nil
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from the scraper
func (s *PlatypusScraper) ScrapePair(pair dia.ExchangePair) (ps PairScraper, err error) {
	return
}

// Returns the list of all available trade pairs
func (s *PlatypusScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return pairs, nil
}

// Channel returns a channel that can be used to receive trades
func (s *PlatypusScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// FillSymbolData adds the name to the asset underlying @symbol on scraper
func (s *PlatypusScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

// NormalizePair accounts for the pair
func (s *PlatypusScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

type PlatypusPairScraper struct {
	parent *PlatypusScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with the scraper
func (ps *PlatypusPairScraper) Close() error {
	ps.parent.errorLock.RLock()
	defer ps.parent.errorLock.RUnlock()
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed and nil otherwise
func (ps *PlatypusPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *PlatypusPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

// Load pools and coins metadata from master registry
func (s *PlatypusScraper) loadPoolsAndCoins(registry platypusRegistry, liquidityThreshold float64, liquidityThresholdUSD float64) (err error) {
	log.Infof("loading master contract %s version %d and querying registry", registry.Address.Hex(), registry.Version)
	contractMaster, err := platypusfinance.NewBaseMasterPlatypusCaller(registry.Address, s.RestClient)
	if err != nil {
		log.Error("NewBaseMasterPlatypusCaller: ", err)
		return err
	}

	poolCount, err := contractMaster.PoolLength(&bind.CallOpts{})
	if err != nil {
		log.Error("PoolLength: ", err)
		return err
	}

	lowerBoundCount := 0
	for i := 0; i < int(poolCount.Int64()); i++ {
		asset, errPoolInfo := contractMaster.PoolInfo(&bind.CallOpts{}, big.NewInt(int64(i)))
		if errPoolInfo != nil {
			log.Error("PoolInfo: ", errPoolInfo)
			return errPoolInfo
		}
		pool, errPool := s.relDB.GetPoolByAddress(dia.ETHEREUM, asset.LpToken.Hex())
		if errPool != nil {
			log.Errorf("Get pool %s by address: %v", asset.LpToken.Hex(), errPool)
		}

		lowLiqui := false
		for _, av := range pool.Assetvolumes {
			if av.Volume < liquidityThreshold {
				log.Warnf("low liquidity on %s: %v", pool.Address, av.Volume)
				lowLiqui = true
				break
			}
		}
		if lowLiqui {
			continue
		}

		liquidity, lowerBound := pool.GetPoolLiquidityUSD()
		// Discard pool if complete USD liquidity is below threshold.
		if !lowerBound && liquidity < liquidityThresholdUSD {
			continue
		}
		if lowerBound {
			lowerBoundCount++
		}

		errPoolData := s.loadPoolData(asset.LpToken.Hex())
		if errPoolData != nil {
			log.Errorf("loadPoolData error at asset %s: %s", asset.LpToken.Hex(), errPoolData)
			return errPoolData
		}
	}
	log.Info("lowerBound: ", lowerBoundCount)

	return err
}

// Runs in a goroutine until scraper is closed
func (s *PlatypusScraper) mainLoop() {

	s.run = true
	for _, pool := range s.pools.poolsAddressNoLock() {
		err := s.watchSwaps(pool)
		if err != nil {
			log.Error("watchSwaps: ", err)
		}
	}
	if s.screenPools {
		err := s.watchNewPools()
		if err != nil {
			log.Error("watchNewPools: ", err)
		}
	}

	go func() {
		defer func() {
			log.Printf("Shutting down main work routine...\n")
			if a := recover(); a != nil {
				log.Errorf("work routine end. Recover msg: %+v", a)
			}
		}()

		for s.run {
			p := <-s.resubscribe
			log.Info("resub to p: ", p)
			if s.run {
				if p == "NEW_POOLS" {
					if s.screenPools {
						log.Info("resubscribe to new pools")
						err := s.watchNewPools()
						if err != nil {
							log.Error("watchNewPools in resubscribe: ", err)
						}
					}
				} else {
					log.Info("resubscribe to swaps from Pool: " + p)
					err := s.watchSwaps(p)
					if err != nil {
						log.Error("watchSwaps in resubscribe: ", err)
					}
				}
			}
		}
	}()

	if s.run {
		if len(s.pairScrapers) == 0 {
			s.error = errors.New("no pairs to scrape provided")
			log.Error(s.error.Error())
		}
	}

	if s.error == nil {
		s.error = errors.New("main loop terminated by Close()")
	}
	s.cleanup(nil)
}

func (s *PlatypusScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone)
}

func (s *PlatypusScraper) watchSwaps(pool string) error {
	contractPool, err := platypusPoolABI.NewPoolFilterer(common.HexToAddress(pool), s.WsClient)
	if err != nil {
		log.Fatal(err)
	}

	sink := make(chan *platypusPoolABI.PoolSwap)
	header, err := s.RestClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	startblock := header.Number.Uint64() - uint64(20)
	log.Infof("subscribing for trades at %s pool", pool)
	sub, err := contractPool.WatchSwap(&bind.WatchOpts{Start: &startblock}, sink, nil, nil)
	if err != nil {
		log.Error("contractPool.WatchSwap: ", err)
		return err
	}

	go func() {
		defer log.Printf("Shutting down pool work routine %s ...\n", pool)
		defer sub.Unsubscribe()

		subscribed := true
		for s.run && subscribed {
			select {
			case err = <-sub.Err():
				if err != nil {
					log.Error("sub error: ", err)
				}
				subscribed = false
				if s.run {
					log.Warn("resubscribe pool: ", pool)
					s.resubscribe <- pool
					log.Info("scraper: ", s)
				}
				log.Warn("subscription error: ", err)
			case swp := <-sink:
				s.processSwap(pool, swp)
			}
		}
	}()

	return err
}

func (s *PlatypusScraper) processSwap(pool string, swap *platypusPoolABI.PoolSwap) {
	foreignName, volume, price, baseToken, quoteToken, err := s.getSwapData(pool, swap)
	if err != nil {
		log.Error("getSwapDataPlatypus: ", err)
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
		ForeignTradeID: swap.Raw.TxHash.Hex() + "-" + fmt.Sprint(swap.Raw.Index),
		PoolAddress:    pool,
		Source:         s.exchangeName,
		VerifiedPair:   true,
	}

	log.Infof("got trade in pool %s with tx %s", pool, trade.ForeignTradeID)
	log.Info("trade: ", trade)
	s.chanTrades <- trade
}

// getSwapDataPlatypus returns the foreign name, volume and price of a swap
func (s *PlatypusScraper) getSwapData(pool string, swap *platypusPoolABI.PoolSwap) (foreignName string, volume float64, price float64, baseToken, quoteToken dia.Asset, err error) {
	fromToken, ok := s.platypusCoins[swap.FromToken.Hex()]
	if !ok {
		log.Errorf("token not found: %s-%s", pool, swap.FromToken.Hex())
		return
	}
	baseToken = dia.Asset{
		Name:       fromToken.Name,
		Address:    fromToken.Address,
		Symbol:     fromToken.Symbol,
		Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
	}

	toToken, ok := s.platypusCoins[swap.ToToken.Hex()]
	if !ok {
		log.Errorf("token not found: %s-%s", pool, swap.ToToken.Hex())
		return
	}
	quoteToken = dia.Asset{
		Name:       toToken.Name,
		Address:    toToken.Address,
		Symbol:     toToken.Symbol,
		Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
	}

	// amountIn = AmountSold / math.Pow10( fromToken.Decimals )
	amountIn, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.FromAmount), new(big.Float).SetFloat64(math.Pow10(int(fromToken.Decimals)))).Float64()

	// amountOut = AmountBought / math.Pow10( toToken.Decimals )
	amountOut, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.ToAmount), new(big.Float).SetFloat64(math.Pow10(int(toToken.Decimals)))).Float64()

	volume = amountOut
	price = amountIn / amountOut

	foreignName = toToken.Symbol + "-" + fromToken.Symbol

	return
}

func (s *PlatypusScraper) watchNewPools() error {
	contractPool, err := platypusPoolABI.NewPoolFilterer(s.basePoolRegistry.Address, s.WsClient)
	if err != nil {
		log.Error("NewPoolFilterer: ", err)
	}

	sink := make(chan *platypusPoolABI.PoolAssetAdded)
	header, err := s.RestClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	startblock := header.Number.Uint64() - uint64(platypusLookBackBlocks)
	sub, err := contractPool.WatchAssetAdded(&bind.WatchOpts{Start: &startblock}, sink, nil, nil)
	if err != nil {
		log.Error("WatchPoolAdded: ", err)
		return err
	}

	go func() {
		log.Infof("subscribing to new asset added events at latest registry")
		defer log.Info("Unsubscribed to new pools")
		defer sub.Unsubscribe()
		subscribed := true

		for s.run && subscribed {

			select {
			case err = <-sub.Err():
				if err != nil {
					log.Error("subscription error in new pools: ", err)
				}
				subscribed = false
				if s.run {
					s.resubscribe <- "NEW_POOLS"
				}
			case vLog := <-sink:

				if _, ok := s.pools.getPool(vLog.Asset.Hex()); !ok {
					err = s.loadPoolData(vLog.Asset.Hex())
					if err != nil {
						log.Error("loadPoolData in new pools: ", err)
					}
					err = s.watchSwaps(vLog.Asset.Hex())
					if err != nil {
						log.Error("watchSwaps in new pools: ", err)
					}
				}
			}
		}
	}()

	return nil
}

func (s *PlatypusScraper) loadPoolData(asset string) error {
	contractAsset, err := platypusAssetABI.NewAssetCaller(common.HexToAddress(asset), s.RestClient)
	if err != nil {
		log.Error("NewAssetCaller: ", err)
		return err
	}

	pool, err := contractAsset.Pool(&bind.CallOpts{})
	if err != nil {
		log.Error("Pool: ", err)
		return err
	}

	contractPool, err := platypusPoolABI.NewPoolCaller(pool, s.RestClient)
	if err != nil {
		log.Error("NewPoolCaller: ", err)
		return err
	}

	poolTokenAddresses, errGetTokens := contractPool.GetTokenAddresses(&bind.CallOpts{})
	if errGetTokens != nil {
		symbol, err := contractAsset.Symbol(&bind.CallOpts{})
		if err != nil {
			log.Error("contractAsset.Symbol: ", err)
		}
		log.Warnf("error calling GetTokenAddresses for %s %s asset: %s", symbol, asset, errGetTokens)
	}

	var poolCoinsMap = make(map[int]*PlatypusCoin)
	for cIdx, c := range poolTokenAddresses {
		var symbol string
		var decimals *big.Int
		var name string
		if c == common.HexToAddress("0x0000000000000000000000000000000000000000") {
			continue
		} else {
			contractToken, err := platypusTokenABI.NewTokenCaller(c, s.RestClient)
			if err != nil {
				log.Error("NewTokenCaller: ", err)
				continue
			}

			symbol, err = contractToken.Symbol(&bind.CallOpts{})
			if err != nil {
				log.Error("Symbol: ", err, c.Hex())
				continue
			}

			decimals, err = contractToken.Decimals(&bind.CallOpts{})
			if err != nil {
				log.Error("Decimals: ", err)
				continue
			}

			name, err = contractToken.Name(&bind.CallOpts{})
			if err != nil {
				log.Error("Name: ", err)
				continue
			}
		}

		poolCoinsMap[cIdx] = &PlatypusCoin{
			Symbol:   symbol,
			Decimals: uint8(decimals.Uint64()),
			Name:     name,
			Address:  c.Hex(),
		}
		s.platypusCoins[c.Hex()] = &PlatypusCoin{
			Symbol:   symbol,
			Decimals: uint8(decimals.Uint64()),
			Name:     name,
			Address:  c.Hex(),
		}
		s.pools.setPool(pool.Hex(), poolCoinsMap)
	}

	return nil
}
