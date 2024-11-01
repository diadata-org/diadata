package scrapers

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	hydrationhelper "github.com/diadata-org/diadata/pkg/dia/helpers/hydration-helper"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/parser"
	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/parser"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

const (
	Blockchain = "Hydration"
)

type HydrationScraper struct {
	logger       *logrus.Entry
	pairScrapers map[string]*HydrationPairScraper // pc.ExchangePair -> pairScraperSet
	shutdown     chan nothing
	shutdownDone chan nothing
	errorLock    sync.RWMutex
	error        error
	closed       bool
	chanTrades   chan *dia.Trade
	db           *models.RelDB
	wsApi        *hydrationhelper.SubstrateEventHelper
	exchangeName string
	blockchain   string
	currentBlock uint64
}

func NewHydrationScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *HydrationScraper {
	logger := logrus.
		New().
		WithContext(context.Background()).
		WithField("context", "HydrationScraper")

	wsApi, err := hydrationhelper.NewSubstrateEventHelper(logger, exchange.WsAPI)
	if err != nil {
		logrus.WithError(err).Error("Failed to create Hydration Substrate event helper")
		return nil
	}

	startBlock := utils.Getenv(strings.ToUpper(exchange.Name)+"_START_BLOCK", "0")
	startBlockUint64, err := strconv.ParseUint(startBlock, 10, 64)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse start block, using default value of 0")
		startBlockUint64 = 0
	}

	s := &HydrationScraper{
		logger:       logger, // Ensure logger is initialized
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
		wsApi:        wsApi,
		exchangeName: exchange.Name,
		blockchain:   Blockchain,
		currentBlock: startBlockUint64,
	}

	s.logger.Info("WS API", s.wsApi)
	if scrape {
		go s.mainLoop()
	}
	return s
}

// processNewBlock processes new blocks and filters SellExecuted events.
// https://hydration.subscan.io/event?block=6148977&page=1&time_dimension=date&module=xyk&event_id=sellexecuted
// pool = 7KKXieLDbfJPUaVohYTbbib97LdC1URmZuMNFq9rvTudmDMv
// block = 0xfd38c9dc2c95278fd3015f73b48a01e804320865a1a6153e31471cb782be92f0
// blocknumber = 6148977
func (s *HydrationScraper) mainLoop() {
	s.logger.Info("Listening for new blocks")
	defer s.cleanup(nil)

	for {
		select {
		case <-s.shutdown:
			s.logger.Println("shutting down")
			return
		default:
			s.logger.Info("Processing block:", s.currentBlock)

			if s.currentBlock == 0 {
				s.wsApi.ListenForNewBlocks(s.processEvents)
			} else {
				s.wsApi.ListenForSpecificBlock(s.currentBlock, s.processEvents)
				s.currentBlock++
				time.Sleep(time.Second)
				latestBlock, err := s.wsApi.API.RPC.Chain.GetBlockLatest()
				if err != nil {
					s.logger.WithError(err).Error("Failed to get latest block")
					return
				}

				if s.currentBlock > uint64(latestBlock.Block.Header.Number) {
					s.logger.Info("Reached the latest block")
					s.wsApi.ListenForNewBlocks(s.processEvents)
				}
			}
		}
	}
}

func (s *HydrationScraper) processEvents(events []*parser.Event, blockNumber uint64) {
	for _, e := range events {
		parsedEvent := s.parseRouterEvent(e)
		if parsedEvent == nil {
			continue
		}
		if e.Phase.IsApplyExtrinsic {
			parsedEvent.ExtrinsicID = fmt.Sprintf("%d-%d", blockNumber, e.Phase.AsApplyExtrinsic)
		}

		pools, err := s.db.GetAllPoolsExchange(s.exchangeName, 0)
		if err != nil {
			s.logger.Error("Failed to get pools from database")
			continue
		}

		pool := s.filterPools(pools, parsedEvent)

		if len(pool.Assetvolumes) < 2 {
			// look for pool address in other events
			secundaryEvent := s.parseSecundaryEvent(events, parsedEvent, blockNumber)
			secundaryPool := s.filterPools(pools, secundaryEvent)

			// Manually fill asset from routerEvent
			assetIn, err := s.db.GetAsset(parsedEvent.AssetIn, s.blockchain)
			if err != nil {
				s.logger.Errorf("Failed to get assetIn for asset address %s", parsedEvent.AssetIn)
				continue

			}
			assetOut, err := s.db.GetAsset(parsedEvent.AssetOut, s.blockchain)
			if err != nil {
				s.logger.Errorf("Failed to get assetOut for asset address %s", parsedEvent.AssetOut)
				continue
			}
			pool.Address = secundaryPool.Address
			pool.Assetvolumes = []dia.AssetVolume{
				{Asset: assetIn},
				{Asset: assetOut},
			}
			if len(pool.Assetvolumes) < 2 {
				s.logger.WithField("poolAddress", pool.Address).Error("Pool has fewer than 2 asset volumes")
				continue
			}
		}

		diaTrade := s.handleTrade(pool, *parsedEvent, time.Now())

		s.logger.WithFields(logrus.Fields{
			"Pair":   diaTrade.Pair,
			"Price":  diaTrade.Price,
			"Volume": diaTrade.Volume,
		}).Info("Trade processed")

		s.chanTrades <- diaTrade
	}
}

func (s *HydrationScraper) filterPools(pools []dia.Pool, event *HydrationParsedEvent) dia.Pool {
	for _, pool := range pools {
		assetInFound := false
		assetOutFound := false

		for _, assetVolume := range pool.Assetvolumes {
			if assetVolume.Asset.Address == event.AssetIn {
				assetInFound = true
			}
			if assetVolume.Asset.Address == event.AssetOut {
				assetOutFound = true
			}
			if assetInFound && assetOutFound {
				return pool
			}
		}
	}
	return dia.Pool{}
}

func (s *HydrationScraper) parseSecundaryEvent(events []*parser.Event, routerEvent *HydrationParsedEvent, blockNumber uint64) *HydrationParsedEvent {
	parsedEvent := &HydrationParsedEvent{}
	for _, e := range events {
		if strings.EqualFold(e.Name, "Omnipool.SellExecuted") ||
			strings.EqualFold(e.Name, "Omnipool.BuyExecuted") ||
			strings.EqualFold(e.Name, "Stableswap.BuyExecuted") ||
			strings.EqualFold(e.Name, "Stableswap.SellExecuted") {
			if e.Phase.IsApplyExtrinsic {
				extrinsicId := fmt.Sprintf("%d-%d", blockNumber, e.Phase.AsApplyExtrinsic)
				if extrinsicId == routerEvent.ExtrinsicID {
					parsedEvent := &HydrationParsedEvent{}
					parsedEvent = s.parseFields(e)
					parsedEvent.AmountIn = routerEvent.AmountIn
					parsedEvent.AmountOut = routerEvent.AmountOut
					parsedEvent.ExtrinsicID = extrinsicId
					return parsedEvent
				}
			}
		}
	}

	return parsedEvent
}

func (s *HydrationScraper) parseRouterEvent(event *parser.Event) *HydrationParsedEvent {
	if strings.EqualFold(event.Name, "Router.Executed") {
		return s.parseFields(event)
	}

	return nil
}

func (s *HydrationScraper) parseFields(event *parser.Event) *HydrationParsedEvent {
	parsedEvent := &HydrationParsedEvent{}
	for _, v := range event.Fields {
		switch v.Name {
		case "asset_in":
			parsedEvent.AssetIn = fmt.Sprint(v.Value)
		case "asset_out":
			parsedEvent.AssetOut = fmt.Sprint(v.Value)
		case "amount_in":
			parsedEvent.AmountIn = fmt.Sprint(v.Value)
		case "amount_out":
			parsedEvent.AmountOut = fmt.Sprint(v.Value)
		}
	}
	return parsedEvent
}

func (s *HydrationScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone)
}

func (s *HydrationScraper) Close() error {
	if s.closed {
		return errors.New("HydrationScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (s *HydrationScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

type HydrationPairScraper struct {
	parent     *HydrationScraper
	pair       dia.ExchangePair
	closed     bool
	lastRecord int64
}

func (ps *HydrationPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

func (ps *HydrationPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *HydrationPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Channel returns the channel used to receive trades/pricing information.
func (s *HydrationScraper) handleTrade(pool dia.Pool, event hydrationhelper.EventSellExecuted, time time.Time) *dia.Trade {
	var volume, price float64
	var decimalsIn, decimalsOut int64
	var quoteToken, baseToken dia.Asset

	// Determine which asset is being sold (this is the base asset)
	if fmt.Sprint(event.AssetIn) == pool.Assetvolumes[0].Asset.Address {
		baseToken = pool.Assetvolumes[0].Asset
		quoteToken = pool.Assetvolumes[1].Asset
	} else {
		baseToken = pool.Assetvolumes[1].Asset
		quoteToken = pool.Assetvolumes[0].Asset
	}

	decimalsIn = int64(baseToken.Decimals)
	decimalsOut = int64(quoteToken.Decimals)
	amountIn, _ := utils.StringToFloat64(event.AmountIn, decimalsIn)
	amountOut, _ := utils.StringToFloat64(event.AmountOut, decimalsOut)

	volume = amountOut

	price = amountIn / amountOut

	symbolPair := fmt.Sprintf("%s-%s", quoteToken.Symbol, baseToken.Symbol)

	return &dia.Trade{
		Time:           time,
		Symbol:         quoteToken.Symbol,
		Pair:           symbolPair,
		ForeignTradeID: event.ExtrinsicID,
		Source:         s.exchangeName,
		Price:          price,
		Volume:         volume,
		VerifiedPair:   true,
		QuoteToken:     quoteToken,
		BaseToken:      baseToken,
		PoolAddress:    pool.Address,
	}
}

func (s *HydrationScraper) FetchAvailablePairs() ([]dia.ExchangePair, error) {
	return []dia.ExchangePair{}, nil
}

func (s *HydrationScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *HydrationScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

func (s *HydrationScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("HydrationScraper: Call ScrapePair on closed scraper")
	}
	ps := &HydrationPairScraper{
		parent:     s,
		pair:       pair,
		lastRecord: 0,
	}

	s.pairScrapers[pair.Symbol] = ps

	return ps, nil
}

type HydrationParsedEvent struct {
	Name        string
	ExtrinsicID string
	AssetIn     string
	AssetOut    string
	AmountIn    string
	AmountOut   string
}
