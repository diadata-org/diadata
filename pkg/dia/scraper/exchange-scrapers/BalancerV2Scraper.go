package scrapers

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"go.uber.org/ratelimit"

	balancervault "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/balancerv2/vault"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
)

const (
	balancerV2RateLimitPerSec = 50
	balancerV2FilterPageSize  = 5000
	balancerV2RestDial        = ""
	balancerV2WSDial          = ""
)

var (
	balancerV2VaultContract          = ""
	balancerV2StartBlockPoolRegister = 16896080
	reverseBasetokensBalancer        *[]string
	reverseQuotetokensBalancer       *[]string
)

// BalancerV2Swap is a swap information
type BalancerV2Swap struct {
	SellToken  string
	BuyToken   string
	SellVolume float64
	BuyVolume  float64
	ID         string
	Timestamp  int64
}

// BalancerV2Scraper is a scraper for Balancer V2
type BalancerV2Scraper struct {
	rest  *ethclient.Client
	ws    *ethclient.Client
	rl    ratelimit.Limiter
	relDB *models.RelDB

	// signaling channels for session initialization and finishing
	shutdown           chan nothing
	shutdownDone       chan nothing
	signalShutdown     sync.Once
	signalShutdownDone sync.Once

	// error handling; err should be read from error(), closed should be read from isClosed()
	// those two methods implement RW lock
	errMutex    sync.RWMutex
	err         error
	closedMutex sync.RWMutex
	closed      bool

	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*BalancerV2PairScraper
	exchangeName string
	chanTrades   chan *dia.Trade

	tokensMap       map[string]dia.Asset
	poolsMap        map[[32]byte]common.Address
	admissiblePools map[common.Address]struct{}
	cachedAssets    sync.Map // map[string]dia.Asset
}

// NewBalancerV2Scraper returns a Balancer V2 scraper
func NewBalancerV2Scraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BalancerV2Scraper {
	balancerV2VaultContract = exchange.Contract
	scraper := &BalancerV2Scraper{
		exchangeName:    exchange.Name,
		err:             nil,
		shutdown:        make(chan nothing),
		shutdownDone:    make(chan nothing),
		pairScrapers:    make(map[string]*BalancerV2PairScraper),
		chanTrades:      make(chan *dia.Trade),
		tokensMap:       make(map[string]dia.Asset),
		poolsMap:        make(map[[32]byte]common.Address),
		admissiblePools: make(map[common.Address]struct{}),
	}

	switch exchange.Name {
	case dia.BalancerV2Exchange:
		balancerV2StartBlockPoolRegister = 12272146
	case dia.BalancerV2ExchangeArbitrum:
		balancerV2StartBlockPoolRegister = 222832
	case dia.BeetsExchange:
		balancerV2StartBlockPoolRegister = 16896080
	case dia.BeetsV2ExchangeSonic:
		balancerV2StartBlockPoolRegister = 368312
	case dia.BalancerV2ExchangePolygon:
		balancerV2StartBlockPoolRegister = 15832990
	}

	var err error

	ws, err := ethclient.Dial(utils.Getenv("ETH_URI_WS", balancerV2WSDial))
	if err != nil {
		log.Fatal("init ws client: ", err)
	}

	rest, err := ethclient.Dial(utils.Getenv("ETH_URI_REST", balancerV2RestDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}

	scraper.relDB = relDB

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

	scraper.fetchAdmissiblePools(liquidityThreshold, liquidityThresholdUSD)

	scraper.ws = ws
	scraper.rest = rest
	scraper.rl = ratelimit.New(balancerV2RateLimitPerSec)

	if scrape {
		go scraper.mainLoop()
	}

	return scraper
}

func (s *BalancerV2Scraper) mainLoop() {

	// Import tokens which appear as base token and we need a quotation for
	var err error
	reverseBasetokensBalancer, err = getReverseTokensFromConfig("balancer/reverse_tokens/" + s.exchangeName + "Basetoken")
	if err != nil {
		log.Error("error getting tokens for which pairs should be reversed: ", err)
	}
	log.Info("reverse basetokens: ", reverseBasetokensBalancer)
	reverseQuotetokensBalancer, err = getReverseTokensFromConfig("balancer/reverse_tokens/" + s.exchangeName + "Quotetoken")
	if err != nil {
		log.Error("error getting tokens for which pairs should be reversed: ", err)
	}
	log.Info("reverse quotetokens: ", reverseQuotetokensBalancer)

	defer s.cleanup()

	filterer, err := balancervault.NewBalancerVaultFilterer(common.HexToAddress(balancerV2VaultContract), s.ws)
	if err != nil {
		s.setError(err)
		log.Fatalf("%s: Cannot create vault filter, err=%s", s.exchangeName, err.Error())
	}

	balancerVaultCaller, err := balancervault.NewBalancerVaultCaller(common.HexToAddress(balancerV2VaultContract), s.rest)
	if err != nil {
		log.Error("balancer vault caller: ", err)
	}

	currBlock, err := s.rest.BlockNumber(context.Background())
	if err != nil {
		s.setError(err)
		log.Fatalf("%s: Cannot get a current block number, err=%s", s.exchangeName, err.Error())
	}

	sink := make(chan *balancervault.BalancerVaultSwap)
	sub, err := filterer.WatchSwap(&bind.WatchOpts{Start: &currBlock}, sink, nil, nil, nil)
	if err != nil {
		s.setError(err)
		log.Fatalf("%s: Cannot watch swap events, err=%s", s.exchangeName, err.Error())
	}

	defer sub.Unsubscribe()

	for {
		select {
		case <-s.shutdown:
			log.Println("BalancerV2Scraper: Shutting down main loop")
		case err := <-sub.Err():
			s.setError(err)
			log.Errorf("BalancerV2Scraper: Subscription error, err=%s", err.Error())
		case event := <-sink:

			// Fetch pool address in order to check admissibility.
			poolAddress, ok := s.poolsMap[event.PoolId]
			if !ok {
				poolAddress, _, err = balancerVaultCaller.GetPool(&bind.CallOpts{}, event.PoolId)
				if err != nil {
					log.Error("get pool: ", err)
				}
			}
			if _, ok = s.admissiblePools[poolAddress]; !ok {
				log.Warnf("pool %s not admissible, skip trade.", poolAddress)
				continue
			}

			assetIn, ok := s.tokensMap[event.TokenIn.Hex()]
			if !ok {
				asset, err := s.assetFromToken(event.TokenIn)
				if err != nil {
					log.Warnf("%s: Retrieving asset-in %s, err=%s", s.exchangeName, event.TokenIn.Hex(), err.Error())
					continue
				}
				s.tokensMap[asset.Address] = asset
				assetIn = asset
			}

			assetOut, ok := s.tokensMap[event.TokenOut.Hex()]
			if !ok {
				asset, err := s.assetFromToken(event.TokenOut)
				if err != nil {
					log.Warnf("%s: Retrieving asset-out %s, err=%s", s.exchangeName, event.TokenOut.Hex(), err.Error())
					continue
				}
				s.tokensMap[asset.Address] = asset
				assetOut = asset
			}

			decimalsIn := int(assetIn.Decimals)
			decimalsOut := int(assetOut.Decimals)
			amountIn, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(event.AmountIn), new(big.Float).SetFloat64(math.Pow10(decimalsIn))).Float64()
			amountOut, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(event.AmountOut), new(big.Float).SetFloat64(math.Pow10(decimalsOut))).Float64()
			swap := BalancerV2Swap{
				SellToken:  assetIn.Symbol,
				BuyToken:   assetOut.Symbol,
				SellVolume: amountIn,
				BuyVolume:  amountOut,
				ID:         event.Raw.TxHash.String() + "-" + fmt.Sprint(event.Raw.Index),
				Timestamp:  time.Now().Unix(),
			}

			foreignName := swap.BuyToken + "-" + swap.SellToken
			volume := swap.BuyVolume
			trade := &dia.Trade{
				Symbol:         swap.BuyToken,
				Pair:           foreignName,
				Price:          swap.SellVolume / swap.BuyVolume,
				Volume:         volume,
				Time:           time.Unix(swap.Timestamp, 0),
				PoolAddress:    poolAddress.Hex(),
				ForeignTradeID: swap.ID,
				Source:         s.exchangeName,
				BaseToken:      assetIn,
				QuoteToken:     assetOut,
				VerifiedPair:   true,
			}
			switch {
			case utils.Contains(reverseBasetokensBalancer, trade.BaseToken.Address):
				// If we need quotation of a base token, reverse pair
				tSwapped, err := dia.SwapTrade(*trade)
				if err == nil {
					trade = &tSwapped
				}
			case utils.Contains(reverseQuotetokensBalancer, trade.QuoteToken.Address):
				// If we don't need quotation of quote token, reverse pair.
				tSwapped, err := dia.SwapTrade(*trade)
				if err == nil {
					trade = &tSwapped
				}
			}

			select {
			case <-s.shutdown:
			case s.chanTrades <- trade:
				// Take into account reversed trade as well in either of both cases
				// 1. Base asset is not bluechip
				// 2. Both assets are bluechip
				if !utils.Contains(reverseQuotetokensBalancer, trade.BaseToken.Address) ||
					(utils.Contains(reverseQuotetokensBalancer, trade.BaseToken.Address) && utils.Contains(reverseQuotetokensBalancer, trade.QuoteToken.Address)) {
					tSwapped, err := dia.SwapTrade(*trade)
					if err == nil {
						s.chanTrades <- &tSwapped
					}
				}
				log.Info("got trade: ", trade)
			}
		}
	}
}

// Close unsubscribes data and closes any existing WebSocket connections, as well as channels of BalancerV2Scraper
func (s *BalancerV2Scraper) Close() error {
	if s.isClosed() {
		return errors.New("BalancerV2Scraper: Already closed")
	}

	s.signalShutdown.Do(func() {
		close(s.shutdown)
	})

	<-s.shutdownDone

	return s.error()
}

// Channel returns a channel that can be used to receive trades
func (s *BalancerV2Scraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from the BalancerV2 scraper
func (s *BalancerV2Scraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	if err := s.error(); err != nil {
		return nil, err
	}
	if s.isClosed() {
		return nil, errors.New("BalancerV2Scraper: Call ScrapePair on closed scraper")
	}

	ps := &BalancerV2PairScraper{
		parent: s,
		pair:   pair,
	}

	s.pairScrapers[pair.ForeignName] = ps

	return ps, nil
}

// fetchAdmissiblePools fetches all pools from postgres with native liquidity > liquidityThreshold and
// (if available) liquidity in USD > liquidityThresholdUSD.
func (s *BalancerV2Scraper) fetchAdmissiblePools(liquidityThreshold float64, liquidityThresholdUSD float64) {
	whitelistedPools := utils.Getenv("POOLS_WHITELIST", "")
	if whitelistedPools != "" {
		log.Warn("only taking into account pools from POOLS_WHITELIST")
		pools := strings.Split(whitelistedPools, ",")
		for _, p := range pools {
			log.Info(p)
			s.admissiblePools[common.HexToAddress(p)] = struct{}{}
		}
		return
	}

	poolsPreselection, err := s.relDB.GetAllPoolsExchange(s.exchangeName, liquidityThreshold)
	if err != nil {
		log.Error("fetch all admissible pools: ", err)
	}
	log.Infof("Found %v pools after preselection.", len(poolsPreselection))

	for _, pool := range poolsPreselection {
		liquidity, lowerBound := pool.GetPoolLiquidityUSD()
		// Discard pool if complete USD liquidity is below threshold.
		if !lowerBound && liquidity < liquidityThresholdUSD {
			continue
		} else {
			s.admissiblePools[common.HexToAddress(pool.Address)] = struct{}{}
		}
	}
	log.Infof("Found %v pools after USD liquidity filtering.", len(s.admissiblePools))
}

func (s *BalancerV2Scraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	pools, err := s.listPools()
	if err != nil {
		log.Warn("list pools: ", err)
		// return nil, err
	}

	log.Infof("%s: Total pools are %v", s.exchangeName, len(pools))

	pp, err := s.listPairs(pools)
	if err != nil {
		return nil, err
	}

	existingPair := make(map[string]struct{})
	for _, p := range pp {
		quoteAddr := p.UnderlyingPair.QuoteToken.Address
		baseAddr := p.UnderlyingPair.BaseToken.Address
		if _, ok := existingPair[baseAddr+":"+quoteAddr]; !ok {
			pairs = append(pairs, p)
			existingPair[baseAddr+":"+quoteAddr] = struct{}{}
		}
	}

	log.Infof("%s: Total pairs are %v", s.exchangeName, len(pairs))

	return
}

func (s *BalancerV2Scraper) assetFromToken(token common.Address) (dia.Asset, error) {
	cached, ok := s.cachedAssets.Load(token.Hex())
	if !ok {
		asset, err := ethhelper.ETHAddressToAsset(token, s.rest, Exchanges[s.exchangeName].BlockChain.Name)
		if err != nil {
			return dia.Asset{}, err
		}

		s.cachedAssets.Store(token.Hex(), asset)

		return asset, nil
	}

	asset := cached.(dia.Asset)

	return asset, nil
}

func (s *BalancerV2Scraper) makePair(token0, token1 common.Address) (dia.ExchangePair, error) {
	asset0, err := s.assetFromToken(token0)
	if err != nil {
		return dia.ExchangePair{}, err
	}
	asset1, err := s.assetFromToken(token1)
	if err != nil {
		return dia.ExchangePair{}, err
	}

	var pair dia.ExchangePair
	pair.UnderlyingPair.QuoteToken = asset0
	pair.UnderlyingPair.BaseToken = asset1
	pair.ForeignName = asset0.Symbol + "-" + asset1.Symbol
	pair.Verified = true
	pair.Exchange = s.exchangeName
	pair.Symbol = asset0.Symbol

	return pair, nil
}

func (s *BalancerV2Scraper) listPairs(pools [][]common.Address) (pairs []dia.ExchangePair, err error) {
	pairCount := 0
	pairMap := make(map[int]dia.ExchangePair)
	var g errgroup.Group
	var mu sync.Mutex
	for _, tokens := range pools {
		if len(tokens) < 2 {
			continue
		}

		for i := 0; i < len(tokens); i++ {
			for j := i + 1; j < len(tokens); j++ {
				pairCount++
				i := i
				j := j
				pairCount := pairCount
				tokens := tokens
				g.Go(func() error {
					s.rl.Take()
					pair, err := s.makePair(tokens[i], tokens[j])
					if err != nil {
						log.Warn(err)

						return nil
					}

					mu.Lock()
					defer mu.Unlock()

					pairMap[pairCount] = pair

					return nil
				})
			}
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	keys := make([]int, 0, len(pairMap))
	for k := range pairMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		pairs = append(pairs, pairMap[k])
	}

	return
}

func (s *BalancerV2Scraper) listPools() ([][]common.Address, error) {
	events, err := s.allRegisteredPools()
	if err != nil {
		return nil, err
	}

	caller, err := balancervault.NewBalancerVaultCaller(common.HexToAddress(balancerV2VaultContract), s.rest)
	if err != nil {
		return nil, err
	}

	var g errgroup.Group
	var mu sync.Mutex
	pools := make([][]common.Address, len(events))
	for idx, evt := range events {
		idx := idx
		evt := evt
		g.Go(func() error {
			s.rl.Take()
			pool, err := caller.GetPoolTokens(&bind.CallOpts{}, evt.PoolId)
			if err != nil {
				log.Warn("get pool tokens: ", err)
				return err
			}

			mu.Lock()
			defer mu.Unlock()

			pools[idx] = pool.Tokens

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return pools, nil
}

func (s *BalancerV2Scraper) allRegisteredPools() ([]*balancervault.BalancerVaultPoolRegistered, error) {
	filterer, err := balancervault.NewBalancerVaultFilterer(common.HexToAddress(balancerV2VaultContract), s.rest)
	if err != nil {
		return nil, err
	}

	currBlock, err := s.rest.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	var offset uint64 = balancerV2FilterPageSize
	var startBlock uint64 = uint64(balancerV2StartBlockPoolRegister)
	var endBlock = startBlock + offset
	var events []*balancervault.BalancerVaultPoolRegistered
	for {
		if endBlock > currBlock {
			endBlock = currBlock
		}
		log.Infof("startblock - endblock: %v --- %v ", startBlock, endBlock)

		it, err := filterer.FilterPoolRegistered(&bind.FilterOpts{
			Start: startBlock,
			End:   &endBlock,
		}, nil, nil)
		if err != nil {
			log.Warn("filterpoolregistered: ", err)
			continue
		}

		for it.Next() {
			events = append(events, it.Event)
		}
		if err := it.Close(); err != nil {
			log.Warn("closing iterator: ", it)
		}

		if endBlock == currBlock {
			break
		}

		startBlock = endBlock + 1
		endBlock = endBlock + offset
	}

	return events, nil
}

// FillSymbolData adds the name to the asset underlying @symbol on BalancerV2
func (s *BalancerV2Scraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *BalancerV2Scraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

func (s *BalancerV2Scraper) cleanup() {
	close(s.chanTrades)
	s.ws.Close()
	s.rest.Close()
	s.close()
	s.signalShutdownDone.Do(func() {
		close(s.shutdownDone)
	})
}

func (s *BalancerV2Scraper) error() error {
	s.errMutex.RLock()
	defer s.errMutex.RUnlock()

	return s.err
}

func (s *BalancerV2Scraper) setError(err error) {
	s.errMutex.Lock()
	defer s.errMutex.Unlock()

	s.err = err
}

func (s *BalancerV2Scraper) isClosed() bool {
	s.closedMutex.RLock()
	defer s.closedMutex.RUnlock()

	return s.closed
}

func (s *BalancerV2Scraper) close() {
	s.closedMutex.Lock()
	defer s.closedMutex.Unlock()

	s.closed = true
}

// BalancerV2PairScraper implements PairScraper for BalancerV2
type BalancerV2PairScraper struct {
	parent *BalancerV2Scraper
	pair   dia.ExchangePair
	closed bool
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (p *BalancerV2PairScraper) Error() error {
	return p.parent.error()
}

// Pair returns the pair this scraper is subscribed to
func (p *BalancerV2PairScraper) Pair() dia.ExchangePair {
	return p.pair
}

// Close stops listening for trades of the pair associated with the BalancerV2Scraper
func (p *BalancerV2PairScraper) Close() error {
	if err := p.parent.error(); err != nil {
		return err
	}
	if p.closed {
		return errors.New("BalancerV2Scraper: Already closed")
	}

	p.closed = true

	return nil
}
