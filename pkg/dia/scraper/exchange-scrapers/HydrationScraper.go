package scrapers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
	wsApi        string
	exchangeName string
	blockchain   string
	wsClient     *websocket.Conn
}

// NewHydrationScraper initializes and returns a new HydrationScraper instance.
// The scraper handles periodic polling of trade data from the specified exchange and stores it in a relational database.
//
// Parameters:
//   - exchange: An instance of dia.Exchange representing the exchange to be scraped.
//     The exchange name is used to fetch environment-specific configuration like refresh delay.
//   - scrape: A boolean flag indicating whether to automatically start scraping upon initialization.
//     If true, the scraping process is started in a separate goroutine.
//   - relDB: A pointer to a RelDB instance (models.RelDB), representing the relational database where
//     trade data will be stored.
//
// Returns:
//   - *HydrationScraper: A pointer to the newly created HydrationScraper instance.
//
// Behavior:
//   - The function calculates the refresh delay for polling trade data by checking an environment variable
//     named "<EXCHANGE_NAME>_REFRESH_DELAY". If the variable is not set, it defaults to 10000 milliseconds (10 seconds).
//   - It initializes the HydrationScraper with a shutdown mechanism, a ticker for periodic updates, a channel for trade data,
//     and a logging mechanism.
//   - If the 'scrape' flag is true, the scraper starts its main scraping loop in a separate goroutine.
func NewHydrationScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *HydrationScraper {
	s := &HydrationScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
		wsApi:        exchange.WsAPI,
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

func (s *HydrationScraper) connectWebSocket() error {
	var err error
	s.wsClient, _, err = websocket.DefaultDialer.Dial(s.wsApi, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to websocket: %w", err)
	}
	s.logger.Info("Connected to Hydration WebSocket")
	return nil
}

func (s *HydrationScraper) mainLoop() {
	if err := s.connectWebSocket(); err != nil {
		s.logger.Error(err)
		return
	}
	defer s.wsClient.Close()

	for {
		select {
		case <-s.shutdown:
			s.logger.Println("shutting down")
			s.cleanup(nil)
			return
		default:
			s.listenForTrades()
		}
	}
}

func (s *HydrationScraper) listenForTrades() {
	_, message, err := s.wsClient.ReadMessage()
	if err != nil {
		s.logger.Error("Error reading message from websocket: ", err)
		return
	}

	var event hydrationhelper.HydrationSwapEvent
	if err := json.Unmarshal(message, &event); err != nil {
		s.logger.Error("Error unmarshalling websocket message: ", err)
		return
	}

	// Check if the event is a Hydration-specific event
	if event.Type != "SellExecuted" || event.PoolID != 101 { // Example checks
		return
	}

	pools, err := s.db.GetPoolsByAssetPair(event.AssetIn, event.AssetOut, s.exchangeName)
	if err != nil {
		s.logger.WithError(err).WithFields(logrus.Fields{
			"assetIn":  event.AssetIn,
			"assetOut": event.AssetOut,
		}).Error("Failed to get pools for asset pair")
		return
	}

	for _, pool := range pools {
		if len(pool.Assetvolumes) < 2 {
			s.logger.WithField("poolAddress", pool.Address).Error("Pool has fewer than 2 asset volumes")
			continue
		}

		diaTrade := s.handleTrade(pool, event, time.Now())
		s.chanTrades <- diaTrade
	}
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

// Channel returns the channel used to receive trades/pricing information.
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

// fetchWithRetry performs an HTTP GET request to the specified endpoint with retries.
//
// This function attempts to fetch data from the provided API endpoint, retrying
// the request up to the specified number of times (`retries`) in case of failure.
// The retry logic handles both network errors and server-side errors (HTTP 5xx).
//
// For each attempt, the function logs the attempt number and the endpoint being requested.
// If a request fails due to a server-side error (HTTP 5xx), the function retries
// after a short delay. Client-side errors (HTTP 4xx) and other non-retryable errors
// are returned immediately.
//
// Parameters:
//   - endpoint: The URL of the API endpoint to request data from.
//   - contentType: The content type to set in the request header (e.g., "application/json").
//   - retries: The number of retry attempts in case of failure.
//
// Returns:
//   - *http.Response: The response object if the request is successful.
//   - error: An error if all retry attempts fail, or if a non-retryable error occurs.
func (s *HydrationScraper) fetchWithRetry(endpoint string, contentType string, retries int) (*http.Response, error) {
	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	var resp *http.Response
	var err error
	for i := 0; i < retries; i++ {
		s.logger.WithField("attempt", i+1).Info("Fetching data from API")
		s.logger.WithField("endpoint", endpoint).Info("Requesting data")

		req, err := http.NewRequest("GET", endpoint, bytes.NewBuffer([]byte("")))
		if err != nil {
			s.logger.WithError(err).Error("Error creating request")
			return nil, err
		}
		req.Header.Set("Content-Type", contentType)

		resp, err = client.Do(req)
		if err == nil && resp != nil {
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				return resp, nil
			}
			resp.Body.Close()

			if resp.StatusCode >= 500 && resp.StatusCode < 600 {
				s.logger.WithField("status", resp.StatusCode).Warn("Server error. Retrying...")
			} else {
				// Client error or other non-retryable response
				return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
			}

		} else if err != nil {
			s.logger.WithError(err).Warn("Failed to fetch data from API. Retrying...")
		}

		time.Sleep(time.Second * 2)
	}
	return nil, err
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
func (s *HydrationScraper) handleTrade(pool dia.Pool, event hydrationhelper.HydrationSwapEvent, time time.Time) *dia.Trade {
	var volume, price float64
	var decimalsIn, decimalsOut int64

	var quoteToken, baseToken dia.Asset

	// Determine which asset is being sold (this is the base asset)
	if event.AssetIn == pool.Assetvolumes[0].Asset.Address {
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

	amountIn, _ := utils.StringToFloat64(event.AmountIn, decimalsIn)
	amountOut, _ := utils.StringToFloat64(event.AmountOut, decimalsOut)

	volume = amountIn
	price = amountOut / amountIn

	symbolPair := fmt.Sprintf("%s-%s", baseToken.Symbol, quoteToken.Symbol)

	return &dia.Trade{
		Time:           time,
		Symbol:         baseToken.Symbol,
		Pair:           symbolPair,
		ForeignTradeID: event.TxID,
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
func (s *HydrationScraper) FetchAvailablePairs() ([]dia.ExchangePair, error) {
	return []dia.ExchangePair{}, nil
}

func (s *HydrationScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *HydrationScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// ScrapePair initializes and returns a `HydrationPairScraper`.
//
// Parameters:
//   - pair: The `dia.ExchangePair` representing the trading pair (e.g: `HydrationPairScraper`) to be scraped.
//
// Returns:
//   - PairScraper: A `PairScraper` (specifically a `HydrationPairScraper`) for the given exchange pair.
//   - error: An error if the scraper is closed or if an error has occurred, otherwise `nil`.
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
