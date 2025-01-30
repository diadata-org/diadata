package scrapers

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	stacks "github.com/diadata-org/diadata/pkg/dia/helpers/stackshelper"
	velar "github.com/diadata-org/diadata/pkg/dia/helpers/velarhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

type VelarScraper struct {
	logger             *logrus.Entry
	pairScrapers       map[string]*VelarPairScraper // pc.ExchangePair -> pairScraperSet
	shutdown           chan nothing
	shutdownDone       chan nothing
	errorLock          sync.RWMutex
	error              error
	closed             bool
	ticker             *time.Ticker
	exchangeName       string
	blockchain         string
	chanTrades         chan *dia.Trade
	api                *stacks.StacksClient
	db                 *models.RelDB
	currentHeight      int
	initialBlockHeight int
}

// NewVelarScraper returns a new VelarScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
// ENV values:
//
//	 	VELAR_SLEEP_TIMEOUT - (optional, millisecond), make timeout between API calls, default "stackshelper.DefaultSleepBetweenCalls" value
//		VELAR_REFRESH_DELAY - (optional, millisecond) refresh data after each poll, default "stackshelper.DefaultRefreshDelay" value
//		VELAR_HIRO_API_KEY - (optional, string), Hiro Stacks API key, improves scraping performance, default = ""
//		VELAR_INITIAL_BLOCK_HEIGHT (optional, int), useful for debug, default = 0
//		VELAR_DEBUG - (optional, bool), make stdout output with alephium client http call, default = false
func NewVelarScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *VelarScraper {
	envPrefix := strings.ToUpper(exchange.Name)

	sleepBetweenCalls := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(
			envPrefix+"_SLEEP_TIMEOUT",
			stacks.DefaultSleepBetweenCalls,
		),
	)
	refreshDelay := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(envPrefix+"_REFRESH_DELAY", stacks.DefaultRefreshDelay),
	)
	hiroAPIKey := utils.Getenv(envPrefix+"_HIRO_API_KEY", "")
	initialBlockHeight := utils.GetenvInt(envPrefix+"_INITIAL_BLOCK_HEIGHT", 0)
	isDebug := utils.GetenvBool(envPrefix+"_DEBUG", false)

	stacksClient := stacks.NewStacksClient(
		log.WithContext(context.Background()).WithField("context", "StacksClient"),
		sleepBetweenCalls,
		hiroAPIKey,
		isDebug,
	)

	s := &VelarScraper{
		shutdown:           make(chan nothing),
		shutdownDone:       make(chan nothing),
		pairScrapers:       make(map[string]*VelarPairScraper),
		ticker:             time.NewTicker(refreshDelay),
		chanTrades:         make(chan *dia.Trade),
		api:                stacksClient,
		db:                 relDB,
		exchangeName:       exchange.Name,
		blockchain:         exchange.BlockChain.Name,
		initialBlockHeight: initialBlockHeight,
	}

	s.logger = logrus.
		New().
		WithContext(context.Background()).
		WithField("context", "VelarDEXScraper")

	if scrape {
		go s.mainLoop()
	}
	return s
}

func (s *VelarScraper) mainLoop() {
	if s.initialBlockHeight <= 0 {
		latestBlock, err := s.api.GetLatestBlock()
		if err != nil {
			s.logger.WithError(err).Error("failed to GetLatestBlock")
			s.cleanup(err)
			return
		}
		s.currentHeight = latestBlock.Height
	} else {
		s.currentHeight = s.initialBlockHeight
	}

	for {
		select {
		case <-s.ticker.C:
			err := s.Update()
			if err != nil {
				s.logger.WithError(err).Error("failed to run Update")
			}
		case <-s.shutdown:
			s.logger.Println("shutting down")
			s.cleanup(nil)
			return
		}
	}
}

func (s *VelarScraper) Update() error {
	txs, err := s.api.GetAllBlockTransactions(s.currentHeight)
	if err != nil {
		return err
	}

	if len(txs) == 0 {
		return nil
	}
	s.currentHeight += 1

	swapEvents, err := s.getSwapEvents(txs)
	if err != nil {
		return err
	}
	if len(swapEvents) == 0 {
		return nil
	}

	pools, err := s.getPools()
	if err != nil {
		s.logger.WithError(err).Error("failed to GetAllPoolsExchange")
		return err
	}

	for _, pool := range pools {
		if len(pool.Assetvolumes) != 2 {
			s.logger.WithField("poolAddress", pool.Address).Error("pool is missing required asset volumes")
			continue
		}

		for _, e := range swapEvents {
			if pool.Address != e.TickerID {
				continue
			}

			diaTrade := s.handleTrade(&pool, &e)
			s.logger.
				WithField("parentAddress", pool.Address).
				WithField("height", s.currentHeight-1).
				WithField("diaTrade", diaTrade).
				Info("trade")
			s.chanTrades <- diaTrade
		}
	}

	return nil
}

func (s *VelarScraper) getPools() ([]dia.Pool, error) {
	return s.db.GetAllPoolsExchange(s.exchangeName, 0)
}

func (s *VelarScraper) getSwapEvents(txs []stacks.Transaction) ([]velar.SwapEvent, error) {
	result := make([]velar.SwapEvent, 0)

	for _, tx := range txs {
		if tx.TxStatus != "success" || tx.TxType != "contract_call" || !s.isSwapTransaction(&tx) {
			continue
		}

		// This is a temporary workaround introduced due to a bug in hiro stacks API.
		// Results returned from /blocks/{block_height}/transactions route have empty
		// `name` field in `contract_call.function_args` list.
		// TODO: remove this as soon as the issue is fixed.
		normalizedTx, err := s.api.GetTransactionAt(tx.TxID)
		if err != nil {
			return nil, err
		}

		events, err := velar.DecodeSwapEvents(normalizedTx)
		if err != nil {
			return nil, err
		}
		result = append(result, events...)
	}

	return result, nil
}

func (s *VelarScraper) isSwapTransaction(tx *stacks.Transaction) bool {
	return strings.HasPrefix(tx.ContractCall.FunctionName, "swap") ||
		tx.ContractCall.FunctionName == "apply" ||
		tx.ContractCall.FunctionName == "r4"
}

func (s *VelarScraper) handleTrade(pool *dia.Pool, event *velar.SwapEvent) *dia.Trade {
	var volume, price float64

	token0 := pool.Assetvolumes[0].Asset
	token1 := pool.Assetvolumes[1].Asset

	amountIn := event.AmountIn.String()
	amountOut := event.AmountOut.String()

	var trade dia.Trade

	if event.TokenIn == token0.Address {
		trade.Pair = fmt.Sprintf("%s-%s", token1.Symbol, token0.Symbol)
		trade.Symbol = token1.Symbol
		trade.BaseToken = token0
		trade.QuoteToken = token1

		amount0In, _ := utils.StringToFloat64(amountIn, int64(token0.Decimals))
		amount1Out, _ := utils.StringToFloat64(amountOut, int64(token1.Decimals))
		volume = amount1Out
		price = amount0In / amount1Out
	} else {
		trade.Pair = fmt.Sprintf("%s-%s", token0.Symbol, token1.Symbol)
		trade.Symbol = token0.Symbol
		trade.BaseToken = token1
		trade.QuoteToken = token0

		amount1In, _ := utils.StringToFloat64(amountIn, int64(token1.Decimals))
		amount0Out, _ := utils.StringToFloat64(amountOut, int64(token0.Decimals))
		volume = amount0Out
		price = amount1In / amount0Out
	}

	trade.Time = time.Unix(int64(event.Timestamp), 0)
	trade.ForeignTradeID = event.TxID
	trade.Source = s.exchangeName
	trade.Price = price
	trade.Volume = volume
	trade.VerifiedPair = true

	trade.PoolAddress = pool.Address
	return &trade
}

func (s *VelarScraper) FetchAvailablePairs() ([]dia.ExchangePair, error) {
	return []dia.ExchangePair{}, nil
}

func (s *VelarScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *VelarScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

func (s *VelarScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("VelarScraper: Call ScrapePair on closed scraper")
	}
	ps := &VelarPairScraper{
		parent:     s,
		pair:       pair,
		lastRecord: 0,
	}

	s.pairScrapers[pair.Symbol] = ps

	return ps, nil
}

func (s *VelarScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	s.ticker.Stop()

	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone)
}

// Close gracefully shuts down the VelarScraper.
func (s *VelarScraper) Close() error {
	if s.closed {
		return errors.New("VelarScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Channel returns the channel used to receive trades/pricing information.
func (s *VelarScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

type VelarPairScraper struct {
	parent     *VelarScraper
	pair       dia.ExchangePair
	closed     bool
	lastRecord int64
}

func (ps *VelarPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

func (ps *VelarPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *VelarPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}
