package scrapers

import (
	"errors"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/velodrome"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	velodromeRestDial = ""
	velodromeWsDial   = ""
	baseRestDial      = ""
	baseWsDial        = ""
)

type VelodromeSwap struct {
	ID         string
	Timestamp  int64
	IndexIn    int
	IndexOut   int
	Amount0In  float64
	Amount0Out float64
	Amount1In  float64
	Amount1Out float64
}

type VelodromeScraper struct {
	RestClient *ethclient.Client
	WsClient   *ethclient.Client
	relDB      *models.RelDB
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock          sync.RWMutex
	error              error
	closed             bool
	pools              []dia.Pool
	listenByAddress    bool
	reverseQuotetokens *[]string
	reverseBasetokens  *[]string
	fullPools          *[]string
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*VelodromePairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
}

func NewVelodromeScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *VelodromeScraper {
	log.Info("NewVelodromeScraper: ", exchange.Name)
	var (
		s               *VelodromeScraper
		listenByAddress bool
		err             error
	)

	switch exchange.Name {
	case dia.VelodromeExchange:
		s = makeVelodromeScraper(exchange, velodromeRestDial, velodromeWsDial, relDB)
	case dia.AerodromeV1Exchange:
		s = makeVelodromeScraper(exchange, baseRestDial, baseWsDial, relDB)
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

	listenByAddress, err = strconv.ParseBool(utils.Getenv("LISTEN_BY_ADDRESS", ""))
	if err != nil {
		log.Fatal("parse LISTEN_BY_ADDRESS: ", err)
	}
	s.listenByAddress = listenByAddress

	s.reverseBasetokens, err = getReverseTokensFromConfig("velodrome/reverse_tokens/" + s.exchangeName + "Basetoken")
	if err != nil {
		log.Error("error getting basetokens for which pairs should be reversed: ", err)
	}
	log.Infof("reverse the following basetokens on %s: %v", s.exchangeName, s.reverseBasetokens)

	s.reverseQuotetokens, err = getReverseTokensFromConfig("velodrome/reverse_tokens/" + s.exchangeName + "Quotetoken")
	if err != nil {
		log.Error("error getting quotetokens for which pairs should be reversed: ", err)
	}
	log.Infof("reverse the following quotetokens on %s: %v", s.exchangeName, s.reverseQuotetokens)

	s.fullPools, err = getReverseTokensFromConfig("velodrome/fullPools/" + s.exchangeName + "FullPools")
	if err != nil {
		log.Error("error getting fullPools for which pairs should be reversed: ", err)
	}
	log.Infof("Take into account both directions of a trade on the following pools: %v", s.fullPools)

	err = s.loadPools(liquidityThreshold, liquidityThresholdUSD)
	if err != nil {
		log.Fatal("load pools: ", err)
	}

	if scrape {
		go s.mainLoop()
	}

	return s

}

func makeVelodromeScraper(exchange dia.Exchange, restDial string, wsDial string, relDB *models.RelDB) *VelodromeScraper {
	var (
		restClient, wsClient *ethclient.Client
		err                  error
		s                    *VelodromeScraper
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

	s = &VelodromeScraper{
		RestClient:   restClient,
		WsClient:     wsClient,
		relDB:        relDB,
		pairScrapers: make(map[string]*VelodromePairScraper),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
	}

	return s
}

func (s *VelodromeScraper) mainLoop() {

	for _, pool := range s.pools {
		s.WatchSwaps(pool)
	}

}

func (s *VelodromeScraper) WatchSwaps(pool dia.Pool) {
	sink, err := s.GetSwapsChannel(common.HexToAddress(pool.Address))
	if err != nil {
		log.Error("error fetching swaps channel: ", err)
	}

	go func() {
		for {
			rawSwap, ok := <-sink
			if ok {
				swap, indexmap := s.normalizeSwap(*rawSwap, pool)
				price, volume := s.getSwapData(swap)
				token0 := pool.Assetvolumes[indexmap[0]].Asset
				token1 := pool.Assetvolumes[indexmap[1]].Asset

				t := &dia.Trade{
					Symbol:         token0.Symbol,
					Pair:           token0.Symbol + "-" + token1.Symbol,
					Price:          price,
					Volume:         volume,
					BaseToken:      token1,
					QuoteToken:     token0,
					Time:           time.Unix(swap.Timestamp, 0),
					PoolAddress:    rawSwap.Raw.Address.Hex(),
					ForeignTradeID: swap.ID,
					Source:         s.exchangeName,
					VerifiedPair:   true,
				}

				switch {
				case utils.Contains(s.reverseBasetokens, token1.Address):
					// If we need quotation of a base token, reverse pair
					tSwapped, err := dia.SwapTrade(*t)
					if err == nil {
						t = &tSwapped
					}
				case utils.Contains(s.reverseQuotetokens, token0.Address):
					// If we need quotation of a base token, reverse pair
					tSwapped, err := dia.SwapTrade(*t)
					if err == nil {
						t = &tSwapped
					}
				}

				if utils.Contains(s.fullPools, pool.Address) {
					tSwapped, err := dia.SwapTrade(*t)
					if err == nil {
						if tSwapped.Price > 0 {
							s.chanTrades <- &tSwapped
						}
					}
				}

				if price > 0 {
					log.Infof("Got trade at time %v - symbol: %s, pair: %s, price: %v, volume:%v", t.Time, t.Symbol, t.Pair, t.Price, t.Volume)
					s.chanTrades <- t
				}
			}
		}
	}()
}

// GetSwapsChannel returns a channel for swaps of the pair with address @pairAddress
func (s *VelodromeScraper) GetSwapsChannel(pairAddress common.Address) (chan *velodrome.IPoolSwap, error) {

	sink := make(chan *velodrome.IPoolSwap)
	var pairFiltererContract *velodrome.IPoolFilterer
	pairFiltererContract, err := velodrome.NewIPoolFilterer(pairAddress, s.WsClient)
	if err != nil {
		log.Fatal(err)
	}

	_, err = pairFiltererContract.WatchSwap(&bind.WatchOpts{}, sink, []common.Address{}, []common.Address{})
	if err != nil {
		log.Error("error in get swaps channel: ", err)
	}

	return sink, nil

}

// normalizeSwap takes a swap as returned by the swap contract's channel and converts it to a VelodromeSwap type
func (s *VelodromeScraper) normalizeSwap(swap velodrome.IPoolSwap, pool dia.Pool) (normalizedSwap VelodromeSwap, indexmap map[uint8]int) {

	// map the on-chain index of the pair's tokens onto their position in @Assetvolumes slice.
	indexmap = make(map[uint8]int)
	indexmap[pool.Assetvolumes[0].Index] = 0
	indexmap[pool.Assetvolumes[1].Index] = 1

	decimals0 := int(pool.Assetvolumes[indexmap[0]].Asset.Decimals)
	decimals1 := int(pool.Assetvolumes[indexmap[1]].Asset.Decimals)

	amount0In, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount0In), new(big.Float).SetFloat64(math.Pow10(decimals0))).Float64()
	amount0Out, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount0Out), new(big.Float).SetFloat64(math.Pow10(decimals0))).Float64()
	amount1In, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount1In), new(big.Float).SetFloat64(math.Pow10(decimals1))).Float64()
	amount1Out, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(swap.Amount1Out), new(big.Float).SetFloat64(math.Pow10(decimals1))).Float64()

	normalizedSwap = VelodromeSwap{
		ID:         swap.Raw.TxHash.Hex(),
		Timestamp:  time.Now().Unix(),
		Amount0In:  amount0In,
		Amount0Out: amount0Out,
		Amount1In:  amount1In,
		Amount1Out: amount1Out,
	}

	if amount0In > 0 {
		normalizedSwap.IndexIn = 0
		normalizedSwap.IndexOut = 1
	} else {
		normalizedSwap.IndexIn = 1
		normalizedSwap.IndexOut = 0
	}
	return
}

func (s *VelodromeScraper) getSwapData(swap VelodromeSwap) (price float64, volume float64) {
	if swap.Amount0In == float64(0) {
		volume = swap.Amount0Out
		price = swap.Amount1In / swap.Amount0Out
		return
	}
	volume = -swap.Amount0In
	price = swap.Amount1Out / swap.Amount0In
	return
}

func (s *VelodromeScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

func (s *VelodromeScraper) Close() error {
	s.closed = true
	return nil
}

func (s *VelodromeScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return pairs, nil
}

func (s *VelodromeScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{}, nil
}

func (up *VelodromeScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

type VelodromePairScraper struct {
	parent *VelodromeScraper
	pair   dia.ExchangePair
	closed bool
}

func (s *VelodromeScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("Velodrome: Call ScrapePair on closed scraper")
	}
	ps := &VelodromePairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

func (ps *VelodromePairScraper) Close() error {
	ps.closed = true
	return nil
}

func (ps *VelodromePairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *VelodromePairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

// loadPools loads all pools with sufficient liquidity from postgres.
func (s *VelodromeScraper) loadPools(liquiThreshold float64, liquidityThresholdUSD float64) (err error) {
	var pools []dia.Pool

	if s.listenByAddress {

		// Only load pool info for addresses from json file.
		poolAddresses, errAddr := getAddressesFromConfig("velodrome/subscribe_pools/" + s.exchangeName)
		if errAddr != nil {
			log.Error("fetch pool addresses from config file: ", errAddr)
		}
		for _, address := range poolAddresses {
			pool, errPool := s.relDB.GetPoolByAddress(Exchanges[s.exchangeName].BlockChain.Name, address.Hex())
			if errPool != nil {
				log.Fatalf("Get pool with address %s: %v", address.Hex(), errPool)
			}
			s.pools = append(s.pools, pool)
		}

	} else {

		// Load all pools above liqui threshold.
		pools, err = s.relDB.GetAllPoolsExchange(s.exchangeName, liquiThreshold)
		if err != nil {
			return
		}

		log.Info("Found ", len(pools), " pools.")
		log.Info("make pool map...")
		lowerBoundCount := 0
		for _, pool := range pools {
			if len(pool.Assetvolumes) != 2 {
				log.Warn("not enough assets in pool with address: ", pool.Address)
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
			s.pools = append(s.pools, pool)
		}
		log.Infof("found %v subscribable pools.", len(s.pools))
		log.Infof("%v pools with lowerBound=true.", lowerBoundCount)
	}

	return
}
