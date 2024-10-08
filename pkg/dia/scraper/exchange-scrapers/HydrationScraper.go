package scrapers

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	hydrationhelper "github.com/diadata-org/diadata/pkg/dia/helpers/hydration-helper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
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
	wsClient     *websocket.Conn
}

func NewHydrationScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *HydrationScraper {

	wsApi, err := hydrationhelper.NewSubstrateEventHelper(hydrationhelper.DiaPolkadotApi)
	if err != nil {
		logrus.WithError(err).Error("Failed to create Hydration Substrate event helper")
		return nil
	}

	s := &HydrationScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
		wsApi: wsApi,
		exchangeName: "Hydration",
		blockchain:   "Hydration",
	}

	s.logger = logrus.
		New().
		WithContext(context.Background()).
		WithField("context", "HydrationScraper")

	s.logger.Info("Initialized HydrationScraper")

	if scrape {
		go s.mainLoop()
	}
	return s
}

// processNewBlock processes new blocks and filters SellExecuted events.
// https://hydration.subscan.io/event?block=6148977&page=1&time_dimension=date&module=xyk&event_id=sellexecuted
// pool = 7KKXieLDbfJPUaVohYTbbib97LdC1URmZuMNFq9rvTudmDMv
// block = 0xfd38c9dc2c95278fd3015f73b48a01e804320865a1a6153e31471cb782be92f0
// blocknumber = 
func (s *HydrationScraper) mainLoop() {
	//go s.wsApi.ListenForNewBlocks(s.processEvents)
	s.logger.Info("Listening for new blocks")
	go s.wsApi.ListenForSpecificBlock(6149553, s.processEvents)
	defer s.cleanup(nil)

	for {
		select {
		case <-s.shutdown:
			s.logger.Println("shutting down")
			s.cleanup(nil)
			return
		}
	}
}

func (s *HydrationScraper) processEvents(events []hydrationhelper.EventSellExecuted) {
	s.logger.Info("Processing events")
	for _, event := range events {

		assetIn := fmt.Sprint(event.AssetIn)
		assetOut := fmt.Sprint(event.AssetOut)
		pool, err := s.db.GetPoolByAssetPair(assetIn, assetOut, s.exchangeName)
		if err != nil {
			continue
		}

		if len(pool.Assetvolumes) < 2 {
			s.logger.WithField("poolAddress", pool.Address).Error("Pool has fewer than 2 asset volumes")
			continue
		}

		diaTrade := s.handleTrade(pool, event, time.Now())
		s.chanTrades <- diaTrade
	}
	s.logger.Fatal("No more events")
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
		decimalsIn = int64(baseToken.Decimals)
		decimalsOut = int64(quoteToken.Decimals)
	} else {
		baseToken = pool.Assetvolumes[1].Asset
		quoteToken = pool.Assetvolumes[0].Asset
		decimalsIn = int64(baseToken.Decimals)
		decimalsOut = int64(quoteToken.Decimals)
	}

	amountIn, _ := utils.StringToFloat64(fmt.Sprint(event.AmountIn), decimalsIn)
	amountOut, _ := utils.StringToFloat64(fmt.Sprint(event.AmountOut), decimalsOut)

	volume = amountIn
	price = amountOut / amountIn

	symbolPair := fmt.Sprintf("%s-%s", baseToken.Symbol, quoteToken.Symbol)

	return &dia.Trade{
		Time:           time,
		Symbol:         baseToken.Symbol,
		Pair:           symbolPair,
		// ForeignTradeID: event.Who, //TODO:Hydration change it to the actual trade ID
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
