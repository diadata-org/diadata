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
	substratehelper "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/parser"

	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

type BifrostScraper struct {
	logger       *logrus.Entry
	pairScrapers map[string]*BifrostPairScraper // pc.ExchangePair -> pairScraperSet
	shutdown     chan nothing
	shutdownDone chan nothing
	errorLock    sync.RWMutex
	error        error
	closed       bool
	ticker       *time.Ticker
	chanTrades   chan *dia.Trade
	db           *models.RelDB
	wsApi        *substratehelper.SubstrateEventHelper
	exchangeName string
	blockchain   string
	currentBlock uint64
}

func NewBifrostScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BifrostScraper {
	logger := logrus.
		New().
		WithContext(context.Background()).
		WithField("context", "BifrostScraper")

	wsApi, err := substratehelper.NewSubstrateEventHelper(exchange.WsAPI, logger)
	if err != nil {
		logrus.WithError(err).Error("Failed to create Bifrost Substrate event helper")
		return nil
	}

	startBlock := utils.Getenv(strings.ToUpper(exchange.Name)+"_START_BLOCK", "0")
	startBlockUint64, err := strconv.ParseUint(startBlock, 10, 64)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse start block, using default value of 10")
		startBlockUint64 = 10
	}

	s := &BifrostScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
		wsApi:        wsApi,
		exchangeName: exchange.Name,
		blockchain:   "Bifrost",
		currentBlock: startBlockUint64,
	}

	s.logger = logger

	s.logger.Info("Initialized BifrostScraper")

	if scrape {
		go s.mainLoop()
	}
	return s
}

func (s *BifrostScraper) mainLoop() {
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

func (s *BifrostScraper) processEvents(events []*parser.Event, blockNumber uint64) {
	s.logger.Info("Processing events")

	for _, e := range events {
		if e.Name == "StableAsset.TokenSwapped" {
			parsedEvent := parseFields(e)
			parsedEvent.ExtrinsicID = fmt.Sprintf("%d-%d", blockNumber, e.Phase.AsApplyExtrinsic)
			pool, err := s.db.GetPoolByAddress(s.blockchain, parsedEvent.PoolId)

			if len(pool.Assetvolumes) < 2 {
				s.logger.WithField("poolAddress", pool.Address).Error("Pool has fewer than 2 asset volumes")
				continue
			}
			if err != nil {
				continue
			}

			diaTrade := s.handleTrade(pool, parsedEvent, time.Now())

			s.logger.WithFields(logrus.Fields{
				"Pair":   diaTrade.Pair,
				"Price":  diaTrade.Price,
				"Volume": diaTrade.Volume,
			}).Info("Trade processed")

			s.chanTrades <- diaTrade
		}
	}
}

type ParsedEvent struct {
	InputAsset   string
	OutputAsset  string
	InputAmount  string
	OutputAmount string
	PoolId       string
	ExtrinsicID  string
}

type TokenValue struct {
	Token *string `json:"Token,omitempty"` // Token variant, e.g., "KSM"
}
type VTokenValue struct {
	VToken string `json:"VToken,omitempty"` // Token variant, e.g., "KSM"
}

// TokenSymbol represents the token symbols as an enum
type TokenSymbol int

const (
	ASG TokenSymbol = iota
	BNC
	KUSD
	DOT
	KSM
	ETH
	KAR
	ZLK
	PHA
	RMRK
	MOVR
)

// String returns the string representation of the TokenSymbol
func (ts TokenSymbol) String() string {
	return [...]string{"ASG", "BNC", "KUSD", "DOT", "KSM", "ETH", "KAR", "ZLK", "PHA", "RMRK", "MOVR"}[ts]
}

func parseFields(event *parser.Event) ParsedEvent {
	var parsedEvent ParsedEvent
	for _, v := range event.Fields {
		switch v.Name {
		case "bifrost_primitives.currency.CurrencyId.input_asset":
			if result, ok := v.Value.(registry.VariantDecoderResult); ok {
				if decodedFields, ok := result.Value.(registry.DecodedFields); ok {
					if len(decodedFields) > 0 {
						parsedEvent.InputAsset = strings.ToLower(result.FieldName) + "-" + strings.ToLower(fmt.Sprint(decodedFields[0].Value))
					}
				}
			}
		case "bifrost_primitives.currency.CurrencyId.output_asset":
			if result, ok := v.Value.(registry.VariantDecoderResult); ok {
				if decodedFields, ok := result.Value.(registry.DecodedFields); ok {
					if len(decodedFields) > 0 {
						if len(decodedFields) > 0 {
							parsedEvent.OutputAsset = strings.ToLower(result.FieldName) + "-" + strings.ToLower(fmt.Sprint(decodedFields[0].Value))
						}
					}
				}
			}

		case "input_amount":
			parsedEvent.InputAmount = fmt.Sprint(v.Value)
		case "output_amount":
			parsedEvent.OutputAmount = fmt.Sprint(v.Value)
		case "pool_id":
			parsedEvent.PoolId = fmt.Sprint(v.Value)
		}
	}
	return parsedEvent
}

// handleTrade processes a swap event and converts it into a dia.Trade object.
//
// This function takes a pool and a corresponding stable swap event, calculates
// the trade volume and price based on the asset amounts in the event, and returns
// a `dia.Trade` object representing the processed trade. The price is calculated
// as the ratio of the input to output amounts, and the volume is set as the
// negative of the input amount to indicate the amount being swapped.
//
// The `dia.Trade` object includes metadata such as the trade timestamp, the trading pair,
// the pool address, and the exchange source.
//
// Parameters:
//   - pool: A `dia.Pool` object representing the liquidity pool where the swap occurred.
//   - event: A `ParsedEvent` containing the swap details such as asset amounts and event ID.
//   - time: The timestamp for the trade event.
//
// Returns:
//   - *dia.Trade: A pointer to the constructed `dia.Trade` object containing the trade details.

func (s *BifrostScraper) handleTrade(pool dia.Pool, event ParsedEvent, time time.Time) *dia.Trade {
	var volume, price float64
	var baseToken, quoteToken dia.Asset
	var decimalsIn, decimalsOut int64

	if fmt.Sprint(event.InputAsset) == pool.Assetvolumes[0].Asset.Address {
		baseToken = pool.Assetvolumes[0].Asset
		quoteToken = pool.Assetvolumes[1].Asset
	} else {
		baseToken = pool.Assetvolumes[1].Asset
		quoteToken = pool.Assetvolumes[0].Asset
	}

	decimalsIn = int64(baseToken.Decimals)
	decimalsOut = int64(quoteToken.Decimals)
	amountIn, _ := utils.StringToFloat64(event.InputAmount, decimalsIn)
	amountOut, _ := utils.StringToFloat64(event.OutputAmount, decimalsOut)

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

// FetchAvailablePairs returns a list with all trading pairs available on
// the exchange associated to the APIScraper. The format is such that it can
// be used by the corr. pairScraper in order to fetch trades.
func (s *BifrostScraper) FetchAvailablePairs() ([]dia.ExchangePair, error) {
	return []dia.ExchangePair{}, nil
}

func (s *BifrostScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *BifrostScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// ScrapePair initializes and returns a `BifrostPairScraper`.
//
// Parameters:
//   - pair: The `dia.ExchangePair` representing the trading pair (e.g: `BifrostPairScraper`) to be scraped.
//
// Returns:
//   - PairScraper: A `PairScraper` (specifically a `BifrostPairScraper`) for the given exchange pair.
//   - error: An error if the scraper is closed or if an error has occurred, otherwise `nil`.
func (s *BifrostScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("BifrostScraper: Call ScrapePair on closed scraper")
	}
	ps := &BifrostPairScraper{
		parent:     s,
		pair:       pair,
		lastRecord: 0,
	}

	s.pairScrapers[pair.Symbol] = ps

	return ps, nil
}

// cleanup handles the shutdown procedure.
func (s *BifrostScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	s.ticker.Stop()

	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone)
}

// Close gracefully shuts down the BifrostScraper.
func (s *BifrostScraper) Close() error {
	if s.closed {
		return errors.New("BifrostScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Channel returns the channel used to receive trades/pricing information.
func (s *BifrostScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

type BifrostPairScraper struct {
	parent     *BifrostScraper
	pair       dia.ExchangePair
	closed     bool
	lastRecord int64
}

func (ps *BifrostPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

func (ps *BifrostPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *BifrostPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}
