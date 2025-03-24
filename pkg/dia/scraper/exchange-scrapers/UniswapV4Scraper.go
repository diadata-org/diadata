package scrapers

import (
	"context"
	"encoding/hex"
	"errors"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	uniswapcontractv4 "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswapv4"
	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UniswapV4Swap struct {
	ID        string
	Timestamp int64
	Pair      dia.Pair
	Amount0   float64
	Amount1   float64
}

type UniswapV4Scraper struct {
	WsClient   *ethclient.Client
	RestClient *ethclient.Client
	relDB      *models.RelDB
	// signaling channels for session initialization and finishing
	//initDone     chan nothing
	run          bool
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*UniswapPairV4Scraper
	pairRecieved chan *UniswapPair
	poolMap      map[[32]byte]dia.Pool

	exchangeName           string
	startBlock             uint64
	waitTime               int
	listenByAddress        bool
	chanTrades             chan *dia.Trade
	factoryContractAddress common.Address
}

var (
	fullPoolsUniswapV4 *[]string
)

// NewUniswapV4Scraper returns a new UniswapV4Scraper
func NewUniswapV4Scraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *UniswapV4Scraper {
	log.Info("NewUniswapScraper ", exchange.Name)
	log.Info("NewUniswapScraper Address ", exchange.Contract)

	var (
		s               *UniswapV4Scraper
		listenByAddress bool
		err             error
	)

	listenByAddress, err = strconv.ParseBool(utils.Getenv("LISTEN_BY_ADDRESS", "false"))
	if err != nil {
		log.Fatal("parse LISTEN_BY_ADDRESS: ", err)
	}

	switch exchange.Name {
	case dia.UniswapExchangeV4:
		s = makeUniswapV4Scraper(exchange, listenByAddress, "", "", "200", uint64(12369621))

	}

	s.relDB = relDB
	s.poolMap = make(map[[32]byte]dia.Pool)

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

	pingNodeInterval, err := strconv.ParseInt(utils.Getenv("PING_SERVER", "0"), 10, 64)
	if err != nil {
		log.Error("parse PING_SERVER: ", err)
	}
	if pingNodeInterval > 0 {
		s.pingNode(pingNodeInterval)
	}

	if scrape {
		go s.mainLoop()
	}
	return s
}

// makeUniswapV4Scraper returns a uniswap scraper as used in NewUniswapV4Scraper.
func makeUniswapV4Scraper(exchange dia.Exchange, listenByAddress bool, restDial string, wsDial string, waitMilliseconds string, startBlock uint64) *UniswapV4Scraper {
	var restClient, wsClient *ethclient.Client
	var err error
	var s *UniswapV4Scraper

	log.Infof("Init rest and ws client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", restDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}
	wsClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", wsDial))
	if err != nil {
		log.Fatal("init ws client: ", err)
	}

	var waitTime int
	waitTimeString := utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_WAIT_TIME", waitMilliseconds)
	waitTime, err = strconv.Atoi(waitTimeString)
	if err != nil {
		log.Error("could not parse wait time: ", err)
		waitTime = 500
	}

	s = &UniswapV4Scraper{
		WsClient:               wsClient,
		RestClient:             restClient,
		shutdown:               make(chan nothing),
		shutdownDone:           make(chan nothing),
		pairScrapers:           make(map[string]*UniswapPairV4Scraper),
		exchangeName:           exchange.Name,
		pairRecieved:           make(chan *UniswapPair),
		error:                  nil,
		chanTrades:             make(chan *dia.Trade),
		waitTime:               waitTime,
		listenByAddress:        listenByAddress,
		startBlock:             startBlock,
		factoryContractAddress: common.HexToAddress(exchange.Contract),
	}
	return s
}

// runs in a goroutine until s is closed
func (s *UniswapV4Scraper) mainLoop() {

	var err error
	reverseBasetokens, err = getReverseTokensFromConfig("uniswapv4/reverse_tokens/" + s.exchangeName + "Basetoken")
	if err != nil {
		log.Error("error getting basetokens for which pairs should be reversed: ", err)
	}
	log.Infof("reverse the following basetokens on %s: %v", s.exchangeName, reverseBasetokens)
	reverseQuotetokens, err = getReverseTokensFromConfig("uniswapv4/reverse_tokens/" + s.exchangeName + "Quotetoken")
	if err != nil {
		log.Error("error getting quotetokens for which pairs should be reversed: ", err)
	}
	log.Infof("reverse the following quotetokens on %s: %v", s.exchangeName, reverseQuotetokens)
	fullPoolsUniswapV4, err = getReverseTokensFromConfig("uniswapv4/fullPools/" + s.exchangeName + "FullPools")
	if err != nil {
		log.Error("error getting fullPools for which pairs should be reversed: ", err)
	}
	log.Infof("Take into account both directions of a trade on the following pools: %v", fullPoolsUniswapV4)

	time.Sleep(4 * time.Second)
	s.run = true

	if len(s.pairScrapers) == 0 {
		s.error = errors.New("uniswap: No pairs to scrape provided")
		log.Error(s.error.Error())
	}

	sink, err := s.getSwapsChannel()
	if err != nil {
		log.Error("error fetching swaps channel: ", err)
	}

	go func() {
		for {
			rawSwap, ok := <-sink
			if ok {

				// TO DO: Can we use liquidity field in order to assess admissibility of a trade?
				swap, err := s.normalizeRawSwap(rawSwap)
				if err != nil {
					log.Error("normalizeRawSwap: ", err)
					continue
				}
				s.sendTrade(swap, hex.EncodeToString(rawSwap.Id[:]))

			}
		}
	}()

}

func (s *UniswapV4Scraper) getSwapsChannel() (chan *uniswapcontractv4.PoolmanagerSwap, error) {
	contract, err := uniswapcontractv4.NewPoolmanagerFilterer(s.factoryContractAddress, s.WsClient)
	if err != nil {
		log.Error(err)
	}
	tradesSink := make(chan *uniswapcontractv4.PoolmanagerSwap)
	_, err = contract.WatchSwap(&bind.WatchOpts{}, tradesSink, [][32]byte{}, []common.Address{})
	if err != nil {
		log.Fatal("WatchSwap: ", err)
	}

	return tradesSink, nil
}

func (s *UniswapV4Scraper) sendTrade(swap UniswapV4Swap, poolID string) {
	price, volume := s.getSwapData(swap)

	t := &dia.Trade{
		Symbol:         swap.Pair.QuoteToken.Symbol,
		Pair:           swap.Pair.QuoteToken.Symbol + "-" + swap.Pair.BaseToken.Symbol,
		Price:          price,
		Volume:         volume,
		BaseToken:      swap.Pair.BaseToken,
		QuoteToken:     swap.Pair.QuoteToken,
		Time:           time.Unix(swap.Timestamp, 0),
		ForeignTradeID: swap.ID,
		PoolAddress:    poolID,
		Source:         s.exchangeName,
		VerifiedPair:   true,
	}

	// switch {
	// case utils.Contains(reverseBasetokens, pool.Token1.Address.Hex()):
	// 	// If we need quotation of a base token, reverse pair
	// 	tSwapped, err := dia.SwapTrade(*t)
	// 	if err == nil {
	// 		t = &tSwapped
	// 	}
	// case utils.Contains(reverseQuotetokens, pool.Token0.Address.Hex()):
	// 	// If we need quotation of a base token, reverse pair
	// 	tSwapped, err := dia.SwapTrade(*t)
	// 	if err == nil {
	// 		t = &tSwapped
	// 	}
	// }

	// if utils.Contains(fullPoolsUniswapV4, pool.Address.Hex()) {
	// 	tSwapped, err := dia.SwapTrade(*t)
	// 	if err == nil {
	// 		if tSwapped.Price > 0 {
	// 			s.chanTrades <- &tSwapped
	// 		}
	// 	}
	// }

	if price > 0 {
		log.Infof("Got trade on pair %s: %v", t.Pair, t)
		log.Info("------")
		s.chanTrades <- t
	}
}

func (s *UniswapV4Scraper) getSwapData(swap UniswapV4Swap) (price float64, volume float64) {
	volume = swap.Amount0
	price = math.Abs(swap.Amount1 / swap.Amount0)
	return
}

// normalizeUniswapSwap takes a raw swap as returned by the swap contract's channel and converts it to a UniswapSwap type
func (s *UniswapV4Scraper) normalizeRawSwap(rawSwap *uniswapcontractv4.PoolmanagerSwap) (normalizedSwap UniswapV4Swap, err error) {

	pool, ok := s.poolMap[rawSwap.Id]
	if !ok {
		pool, err = s.relDB.GetPoolByAddress(dia.ETHEREUM, hex.EncodeToString(rawSwap.Id[:]))
		if err != nil || len(pool.Assetvolumes) != 2 {
			return
		}
		if len(pool.Assetvolumes) == 2 {
			s.poolMap[rawSwap.Id] = pool
		}
	}

	asset0 := pool.Assetvolumes[pool.Assetvolumes[0].Index].Asset
	asset1 := pool.Assetvolumes[pool.Assetvolumes[1].Index].Asset
	log.Infof("%s-%s", asset0.Symbol, asset1.Symbol)
	decimals0 := int(asset0.Decimals)
	decimals1 := int(asset1.Decimals)
	amount0Big := new(big.Float).Quo(big.NewFloat(0).SetInt(rawSwap.Amount0), new(big.Float).SetFloat64(math.Pow10(decimals0)))
	amount1Big := new(big.Float).Quo(big.NewFloat(0).SetInt(rawSwap.Amount1), new(big.Float).SetFloat64(math.Pow10(decimals1)))
	amount0, _ := amount0Big.Float64()
	amount1, _ := amount1Big.Float64()

	// slippage0, _ := new(big.Float).Quo(amount0Big, big.NewFloat(0).SetInt(rawSwap.Liquidity)).Float64()
	// slippage1, _ := new(big.Float).Quo(amount1Big, big.NewFloat(0).SetInt(rawSwap.Liquidity)).Float64()
	// normalizedLiquidity, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(rawSwap.Liquidity), new(big.Float).SetFloat64(math.Pow10(18))).Float64()
	// log.Infof("slippage0 -- slippage1 -- liquidity: %v -- %v -- %v", slippage0, slippage1, normalizedLiquidity)

	var pairTicker string
	if amount0 < 0 {
		pairTicker = asset1.Symbol + "-" + asset0.Symbol
	} else {
		pairTicker = asset0.Symbol + "-" + asset1.Symbol
	}
	slippage := computeSlippage(rawSwap.SqrtPriceX96, rawSwap.Amount0, rawSwap.Amount1, rawSwap.Liquidity)
	if slippage != nil {
		log.Infof(
			"%s -- slippage: %s",
			pairTicker,
			slippage.String(),
		)
	}

	normalizedSwap = UniswapV4Swap{
		ID:        rawSwap.Raw.TxHash.Hex(),
		Timestamp: time.Now().Unix(),
		Pair:      dia.Pair{QuoteToken: asset0, BaseToken: asset1},
		Amount0:   amount0,
		Amount1:   amount1,
	}

	return
}

// FetchAvailablePairs returns a list with all available trade pairs as dia.Pair for the pairDiscorvery service
func (s *UniswapV4Scraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return
}

func (s *UniswapV4Scraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *UniswapV4Scraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *UniswapV4Scraper) Close() error {

	if s.closed {
		return errors.New("UniswapScraper: Already closed")
	}
	s.WsClient.Close()
	s.RestClient.Close()
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *UniswapV4Scraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("UniswapScraper: Call ScrapePair on closed scraper")
	}
	ps := &UniswapPairV4Scraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

func (s *UniswapV4Scraper) pingNode(pingNodeInterval int64) {
	ticker := time.NewTicker(time.Duration(pingNodeInterval) * time.Second)
	go func() {
		for range ticker.C {
			blockNumber, err := s.WsClient.BlockNumber(context.Background())
			if err != nil {
				log.Error("pingNode: ", err)
			} else {
				log.Infof("%v -- blockNumber: %d", time.Now(), blockNumber)
			}
		}
	}()
}

// UniswapPairScraper implements PairScraper for Uniswap
type UniswapPairV4Scraper struct {
	parent *UniswapV4Scraper
	pair   dia.ExchangePair
	//closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *UniswapPairV4Scraper) Close() error {
	return nil
}

// Channel returns a channel that can be used to receive trades
func (s *UniswapV4Scraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *UniswapPairV4Scraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *UniswapPairV4Scraper) Pair() dia.ExchangePair {
	return ps.pair
}

func computeSlippage(sqrtPriceX96 *big.Int, amount0 *big.Int, amount1 *big.Int, liquidity *big.Int) (slippage *big.Float) {

	price := new(big.Float).Quo(big.NewFloat(0).SetInt(sqrtPriceX96), new(big.Float).SetFloat64(math.Pow(2, 96)))

	if amount0.Sign() < 0 {
		// token0 -> token1
		amount0Abs := big.NewInt(0).Abs(amount0)
		numerator := big.NewFloat(0).Mul(big.NewFloat(0).SetInt(amount0Abs), price)
		return new(big.Float).Quo(numerator, big.NewFloat(0).SetInt(liquidity))
	} else if amount1.Sign() < 0 {
		// token1 -> token0
		numerator := big.NewFloat(0).SetInt(big.NewInt(0).Abs(amount1))
		denominator := big.NewFloat(0).Mul(big.NewFloat(0).SetInt(liquidity), price)
		return new(big.Float).Quo(numerator, denominator)
	}
	log.Infof("sqrtPrice -- amount0 -- amount1 -- liquidity: %s -- %s -- %s -- %s", sqrtPriceX96.String(), amount0.String(), amount1.String(), liquidity.String())
	return nil

}
