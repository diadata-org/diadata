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
	"github.com/didaunesp/no-signature-go-substrate-rpc-client-v4/registry"
	"github.com/didaunesp/no-signature-go-substrate-rpc-client-v4/registry/parser"

	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

const (
	Blockchain = "bifrost"
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

	//wsApi, err := substratehelper.NewSubstrateEventHelper("wss://bifrost-polkadot.api.onfinality.io/public-ws", logger)
	//wsApi, err := substratehelper.NewSubstrateEventHelper("wss://bifrost-rpc.liebi.com/ws", logger)
	wsApi, err := substratehelper.NewSubstrateEventHelper(exchange.WsAPI, logger)
	if err != nil {
		logrus.WithError(err).Error("Failed to create Bifrost Substrate event helper")
		return nil
	}

	startBlock := utils.Getenv(strings.ToUpper(exchange.Name)+"_START_BLOCK", "10")
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
		blockchain:   Blockchain,
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
	// Bifrost Kusama block for testing
	go s.wsApi.ListenForSpecificBlock(7560993, s.processEvents)
	//go s.wsApi.ListenForNewBlocks(s.processEvents)

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

func (s *BifrostScraper) processEvents(events []*parser.Event) {
	s.logger.Info("Processing events")

	for _, e := range events {
		if e.Name == "StableAsset.TokenSwapped" {
			parsedEvent := parseFields(e)

			pool, err := s.db.GetPoolByAssetPair(parsedEvent.InputAsset, parsedEvent.OutputAsset, s.exchangeName)
			//pool, err := s.db.GetPoolById(parsedEvent.PoolId, s.exchangeName)

			if len(pool.Assetvolumes) < 2 {
				s.logger.WithField("poolAddress", pool.Address).Error("Pool has fewer than 2 asset volumes")
				continue
			}
			if err != nil {
				continue
			}

			diaTrade := s.handleTrade(pool, parsedEvent, time.Now())

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
}

// AssetValue represents the value of the asset (could be Token, VToken, etc.)
type AssetValue struct {
	Token  *string `json:"Token,omitempty"`  // Token variant, e.g., "KSM"
	VToken *string `json:"VToken,omitempty"` // Optional VToken variant (if applicable)
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
						if enumValue, ok := decodedFields[0].Value.(uint8); ok {
							tokenSymbol := TokenSymbol(enumValue).String()
							parsedEvent.InputAsset = strings.ToLower(result.FieldName) + "-" + strings.ToLower(tokenSymbol)
						}
					}
				}
			}
		case "bifrost_primitives.currency.CurrencyId.output_asset":
			if result, ok := v.Value.(registry.VariantDecoderResult); ok {
				if decodedFields, ok := result.Value.(registry.DecodedFields); ok {
					if len(decodedFields) > 0 {
						if enumValue, ok := decodedFields[0].Value.(uint8); ok {
							tokenSymbol := TokenSymbol(enumValue).String()
							parsedEvent.OutputAsset = strings.ToLower(result.FieldName) + "-" + strings.ToLower(tokenSymbol)
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
//   - event: A `StableSwapEvent` containing the swap details such as asset amounts and transaction ID.
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
		decimalsIn = int64(baseToken.Decimals)
		decimalsOut = int64(quoteToken.Decimals)
	} else {
		baseToken = pool.Assetvolumes[1].Asset
		quoteToken = pool.Assetvolumes[0].Asset
		decimalsIn = int64(baseToken.Decimals)
		decimalsOut = int64(quoteToken.Decimals)
	}

	amountIn, _ := utils.StringToFloat64(event.InputAmount, decimalsIn)
	amountOut, _ := utils.StringToFloat64(event.OutputAmount, decimalsOut)

	volume = amountIn

	price = amountOut / amountIn

	symbolPair := fmt.Sprintf("%s-%s", baseToken.Symbol, quoteToken.Symbol)

	return &dia.Trade{
		Time:   time,
		Symbol: baseToken.Symbol,
		Pair:   symbolPair,
		//ForeignTradeID: event.TxID,
		Source:       s.exchangeName,
		Price:        price,
		Volume:       volume,
		VerifiedPair: true,
		QuoteToken:   quoteToken,
		BaseToken:    baseToken,
		PoolAddress:  pool.Address,
	}
}

// FetchAvailablePairs returns a list with all trading pairs available on
// the exchange associated to the APIScraper. The format is such that it can
// be used by the corr. pairScraper in order to fetch trades.
func (s *BifrostScraper) FetchAvailablePairs() ([]dia.ExchangePair, error) {
	return []dia.ExchangePair{}, nil
}

// getDBPools retrieves all pools from the database associated with the current exchange.
//
// Returns:
//   - []dia.Pool: A slice of `dia.Pool` objects representing the pools retrieved from the database.
//   - error: An error if the database query fails, otherwise `nil`.
func (s *BifrostScraper) getDBPools() ([]dia.Pool, error) {
	return s.db.GetAllPoolsExchange(s.exchangeName, 0)
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
