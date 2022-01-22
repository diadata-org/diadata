package scrapers

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"go.uber.org/ratelimit"

	balancervault "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/balancerv2/vault"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
)

const (
	balancerV2RateLimitPerSec        = 50
	balancerV2StartBlockPoolRegister = 12272146
	balancerV2FilterPageSize         = 50000
	balancerV2VaultContract          = "0xBA12222222228d8Ba445958a75a0704d566BF2C8"
	balancerV2RestDial               = "http://159.69.120.42:8545/"
	balancerV2WSDial                 = "ws://159.69.120.42:8546/"
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
	rest *ethclient.Client
	ws   *ethclient.Client
	rl   ratelimit.Limiter

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

	tokensMap    map[string]dia.Asset
	cachedAssets sync.Map // map[string]dia.Asset
}

// NewBalancerV2Scraper returns a Balancer V2 scraper
func NewBalancerV2Scraper(exchange dia.Exchange, scrape bool) *BalancerV2Scraper {
	scraper := &BalancerV2Scraper{
		exchangeName: exchange.Name,
		err:          nil,
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*BalancerV2PairScraper),
		chanTrades:   make(chan *dia.Trade),
		tokensMap:    make(map[string]dia.Asset),
	}

	ws, err := ethclient.Dial(balancerV2WSDial)
	if err != nil {
		log.Error(err)

		return nil
	}

	rest, err := ethclient.Dial(balancerV2RestDial)
	if err != nil {
		log.Error(err)

		return nil
	}

	scraper.ws = ws
	scraper.rest = rest
	scraper.rl = ratelimit.New(balancerV2RateLimitPerSec)

	if scrape {
		go scraper.mainLoop()
	}

	return scraper
}

func (s *BalancerV2Scraper) mainLoop() {
	defer s.cleanup()

	pairs, err := s.FetchAvailablePairs()
	if err != nil {
		s.setError(err)
		log.Errorf("BalancerV2Scraper: Cannot fetch avaiable pairs ,err=%s", err.Error())

		return
	}

	for _, pair := range pairs {
		quoteToken := pair.UnderlyingPair.QuoteToken
		baseToken := pair.UnderlyingPair.BaseToken
		s.tokensMap[quoteToken.Address] = quoteToken
		s.tokensMap[baseToken.Address] = baseToken
	}

	filterer, err := balancervault.NewBalancerVaultFilterer(common.HexToAddress(balancerV2VaultContract), s.ws)
	if err != nil {
		s.setError(err)
		log.Errorf("BalancerV2Scraper: Cannot create vault filter, err=%s", err.Error())

		return
	}

	currBlock, err := s.rest.BlockNumber(context.Background())
	if err != nil {
		s.setError(err)
		log.Errorf("BalancerV2Scraper: Cannot get a current block number, err=%s", err.Error())

		return
	}

	sink := make(chan *balancervault.BalancerVaultSwap)
	sub, err := filterer.WatchSwap(&bind.WatchOpts{Start: &currBlock}, sink, nil, nil, nil)
	if err != nil {
		s.setError(err)
		log.Errorf("BalancerV2Scraper: Cannot watch swap events, err=%s", err.Error())

		return
	}

	defer sub.Unsubscribe()

	for {
		select {
		case <-s.shutdown:
			log.Println("BalancerV2Scraper: Shutting down main loop")

			return
		case err := <-sub.Err():
			s.setError(err)
			log.Errorf("BalancerV2Scraper: Subscription error, err=%s", err.Error())

			return
		case event := <-sink:
			assetIn, ok := s.tokensMap[event.TokenIn.Hex()]
			if !ok {
				asset, err := s.assetFromToken(event.TokenIn)
				if err != nil {
					log.Warnf("BalancerV2Scraper: Retrieving asset-in %s, err=%s", event.TokenIn.Hex(), err.Error())

					continue
				}

				s.tokensMap[asset.Address] = asset
				assetIn = asset
			}

			assetOut, ok := s.tokensMap[event.TokenOut.Hex()]
			if !ok {
				asset, err := s.assetFromToken(event.TokenOut)
				if err != nil {
					log.Warnf("BalancerV2Scraper: Retrieving asset-out %s, err=%s", event.TokenOut.Hex(), err.Error())

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
			revForeignName := swap.SellToken + "-" + swap.BuyToken
			pairScraper, hasPair := s.pairScrapers[foreignName]
			revPairScraper, hasRevPair := s.pairScrapers[revForeignName]
			if !hasPair && !hasRevPair {
				log.Warnf("BalancerV2Scraper: Pair %s does not have a corresponding pair scraper", foreignName)

				continue
			}

			var ps *BalancerV2PairScraper
			volume := swap.BuyVolume
			if hasPair {
				ps = pairScraper
			}
			if hasRevPair {
				ps = revPairScraper
				volume = -volume
			}

			trade := &dia.Trade{
				Symbol:         ps.pair.Symbol,
				Pair:           ps.pair.ForeignName,
				Price:          swap.SellVolume / swap.BuyVolume,
				Volume:         volume,
				Time:           time.Unix(swap.Timestamp, 0),
				ForeignTradeID: swap.ID,
				Source:         s.exchangeName,
				BaseToken:      assetIn,
				QuoteToken:     assetOut,
				VerifiedPair:   true,
			}

			select {
			case <-s.shutdown:
			case s.chanTrades <- trade:
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

func (s *BalancerV2Scraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	pools, err := s.listPools()
	if err != nil {
		return nil, err
	}

	log.Info("BalancerV2Scraper: Total pools are ", len(pools))

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

	log.Info("BalancerV2Scraper: Total pairs are ", len(pairs))

	return
}

func (s *BalancerV2Scraper) assetFromToken(token common.Address) (dia.Asset, error) {
	cached, ok := s.cachedAssets.Load(token.Hex())
	if !ok {
		asset, err := ethhelper.ETHAddressToAsset(token, s.rest, dia.ETHEREUM)
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
	var startBlock uint64 = balancerV2StartBlockPoolRegister
	var endBlock = startBlock + offset
	var events []*balancervault.BalancerVaultPoolRegistered
	for {
		if endBlock > currBlock {
			endBlock = currBlock
		}

		it, err := filterer.FilterPoolRegistered(&bind.FilterOpts{
			Start: startBlock,
			End:   &endBlock,
		}, nil, nil)
		if err != nil {
			return nil, err
		}

		for it.Next() {
			events = append(events, it.Event)
		}
		if err := it.Close(); err != nil {
			return nil, err
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
