package scrapers

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"go.uber.org/ratelimit"

	vault "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/balancerv3/vault"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
)

const (
	balancerV3RateLimitPerSec = 50
	balancerV3FilterPageSize  = 5000
	balancerV3RestDial        = ""
	balancerV3WSDial          = ""
)

var (
	balancerV3VaultContract      = ""
	reverseBasetokensBalancerV3  *[]string
	reverseQuotetokensBalancerV3 *[]string
)

// BalancerV3Swap is a swap information
type BalancerV3Swap struct {
	SellToken  string
	BuyToken   string
	SellVolume float64
	BuyVolume  float64
	ID         string
	Timestamp  int64
}

// BalancerV3Scraper is a scraper for Balancer V3
type BalancerV3Scraper struct {
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
	pairScrapers map[string]*BalancerV3PairScraper
	exchangeName string
	chanTrades   chan *dia.Trade

	tokensMap       map[string]dia.Asset
	admissiblePools map[common.Address]struct{}
	cachedAssets    sync.Map // map[string]dia.Asset
}

// NewBalancerV3Scraper returns a Balancer V3 scraper
func NewBalancerV3Scraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BalancerV3Scraper {
	balancerV3VaultContract = exchange.Contract
	scraper := &BalancerV3Scraper{
		exchangeName:    exchange.Name,
		err:             nil,
		shutdown:        make(chan nothing),
		shutdownDone:    make(chan nothing),
		pairScrapers:    make(map[string]*BalancerV3PairScraper),
		chanTrades:      make(chan *dia.Trade),
		tokensMap:       make(map[string]dia.Asset),
		admissiblePools: make(map[common.Address]struct{}),
	}

	var err error

	ws, err := ethclient.Dial(utils.Getenv("ETH_URI_WS", balancerV3WSDial))
	if err != nil {
		log.Fatal("init ws client: ", err)
	}

	rest, err := ethclient.Dial(utils.Getenv("ETH_URI_REST", balancerV3RestDial))
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
	scraper.rl = ratelimit.New(balancerV3RateLimitPerSec)

	if scrape {
		go scraper.mainLoop()
	}

	return scraper
}

func (s *BalancerV3Scraper) mainLoop() {

	// Import tokens which appear as base token and we need a quotation for
	var err error
	reverseBasetokensBalancerV3, err = getReverseTokensFromConfig("balancer/reverse_tokens/" + s.exchangeName + "Basetoken")
	if err != nil {
		log.Error("error getting tokens for which pairs should be reversed: ", err)
	}
	log.Info("reverse basetokens: ", reverseBasetokensBalancerV3)
	reverseQuotetokensBalancerV3, err = getReverseTokensFromConfig("balancer/reverse_tokens/" + s.exchangeName + "Quotetoken")
	if err != nil {
		log.Error("error getting tokens for which pairs should be reversed: ", err)
	}
	log.Info("reverse quotetokens: ", reverseQuotetokensBalancerV3)

	defer s.cleanup()

	filterer, err := vault.NewVaultFilterer(common.HexToAddress(balancerV3VaultContract), s.ws)
	if err != nil {
		s.setError(err)
		log.Fatalf("%s: Cannot create vault filter, err=%s", s.exchangeName, err.Error())
	}

	currBlock, err := s.rest.BlockNumber(context.Background())
	if err != nil {
		s.setError(err)
		log.Fatalf("%s: Cannot get a current block number, err=%s", s.exchangeName, err.Error())
	}

	sink := make(chan *vault.VaultSwap)
	sub, err := filterer.WatchSwap(&bind.WatchOpts{Start: &currBlock}, sink, nil, nil, nil)
	if err != nil {
		s.setError(err)
		log.Fatalf("%s: Cannot watch swap events, err=%s", s.exchangeName, err.Error())
	}

	defer sub.Unsubscribe()

	for {
		select {
		case <-s.shutdown:
			log.Println("BalancerV3Scraper: Shutting down main loop")
		case err := <-sub.Err():
			s.setError(err)
			log.Errorf("BalancerV3Scraper: Subscription error, err=%s", err.Error())
		case event := <-sink:

			if _, ok := s.admissiblePools[event.Pool]; !ok {
				log.Warnf("pool %s not admissible, skip trade.", event.Pool)
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
			swap := BalancerV3Swap{
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
				PoolAddress:    event.Pool.Hex(),
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

// Close unsubscribes data and closes any existing WebSocket connections, as well as channels of BalancerV3Scraper
func (s *BalancerV3Scraper) Close() error {
	if s.isClosed() {
		return errors.New("BalancerV3Scraper: Already closed")
	}

	s.signalShutdown.Do(func() {
		close(s.shutdown)
	})

	<-s.shutdownDone

	return s.error()
}

// Channel returns a channel that can be used to receive trades
func (s *BalancerV3Scraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from the BalancerV3 scraper
func (s *BalancerV3Scraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	if err := s.error(); err != nil {
		return nil, err
	}
	if s.isClosed() {
		return nil, errors.New("BalancerV3Scraper: Call ScrapePair on closed scraper")
	}

	ps := &BalancerV3PairScraper{
		parent: s,
		pair:   pair,
	}

	s.pairScrapers[pair.ForeignName] = ps

	return ps, nil
}

// fetchAdmissiblePools fetches all pools from postgres with native liquidity > liquidityThreshold and
// (if available) liquidity in USD > liquidityThresholdUSD.
func (s *BalancerV3Scraper) fetchAdmissiblePools(liquidityThreshold float64, liquidityThresholdUSD float64) {
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

func (s *BalancerV3Scraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return
}

func (s *BalancerV3Scraper) assetFromToken(token common.Address) (dia.Asset, error) {
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

func (s *BalancerV3Scraper) makePair(token0, token1 common.Address) (dia.ExchangePair, error) {
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

// FillSymbolData adds the name to the asset underlying @symbol on BalancerV3
func (s *BalancerV3Scraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *BalancerV3Scraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

func (s *BalancerV3Scraper) cleanup() {
	close(s.chanTrades)
	s.ws.Close()
	s.rest.Close()
	s.close()
	s.signalShutdownDone.Do(func() {
		close(s.shutdownDone)
	})
}

func (s *BalancerV3Scraper) error() error {
	s.errMutex.RLock()
	defer s.errMutex.RUnlock()

	return s.err
}

func (s *BalancerV3Scraper) setError(err error) {
	s.errMutex.Lock()
	defer s.errMutex.Unlock()

	s.err = err
}

func (s *BalancerV3Scraper) isClosed() bool {
	s.closedMutex.RLock()
	defer s.closedMutex.RUnlock()

	return s.closed
}

func (s *BalancerV3Scraper) close() {
	s.closedMutex.Lock()
	defer s.closedMutex.Unlock()

	s.closed = true
}

// BalancerV3PairScraper implements PairScraper for BalancerV3
type BalancerV3PairScraper struct {
	parent *BalancerV3Scraper
	pair   dia.ExchangePair
	closed bool
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (p *BalancerV3PairScraper) Error() error {
	return p.parent.error()
}

// Pair returns the pair this scraper is subscribed to
func (p *BalancerV3PairScraper) Pair() dia.ExchangePair {
	return p.pair
}

// Close stops listening for trades of the pair associated with the BalancerV3Scraper
func (p *BalancerV3PairScraper) Close() error {
	if err := p.parent.error(); err != nil {
		return err
	}
	if p.closed {
		return errors.New("BalancerV3Scraper: Already closed")
	}

	p.closed = true

	return nil
}
