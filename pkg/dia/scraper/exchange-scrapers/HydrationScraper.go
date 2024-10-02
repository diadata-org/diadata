package scrapers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	hydrationhelper "github.com/diadata-org/diadata/pkg/dia/helpers/hydration-helper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
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
	ticker       *time.Ticker
	chanTrades   chan *dia.Trade
	db           *models.RelDB
	api          string
	exchangeName string
	blockchain   string
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
	refreshDelay := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(strings.ToUpper(exchange.Name)+"_REFRESH_DELAY", 10000),
	)

	apiURL := utils.Getenv(strings.ToUpper(exchange.Name)+"_API_URL", "http://localhost:3000/hydration/v1")

	s := &HydrationScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		ticker:       time.NewTicker(refreshDelay),
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
		api:          apiURL,
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

// mainLoop is the core loop of the HydrationScraper that periodically fetches
// trade data and handles graceful shutdown.
//
// This method runs an infinite loop where it performs two key actions:
//   - Periodically triggers the Update function when the ticker ticks (based on
//     the refresh interval).
//   - Listens for a shutdown signal on the shutdown channel, at which point it
//     cleans up resources and exits the loop.
//
// Behavior:
//   - On each tick of the ticker (`s.ticker.C`), the Update method is called
//     to fetch and process new trade data from the exchange.
//   - If an error occurs during the Update call, it is logged using `s.logger`.
//   - When a message is received on the `s.shutdown` channel, the method logs
//     the shutdown event, calls the cleanup function, and exits.
//
// This method is designed to run as a goroutine and will keep looping until
// explicitly stopped by sending a signal to the `shutdown` channel.
func (s *HydrationScraper) mainLoop() {
	s.logger.Info("Starting main loop %v", s.ticker)
	for {
		select {
		case <-s.ticker.C:
			err := s.Update()
			if err != nil {
				s.logger.Error(err)
			}
		case <-s.shutdown:
			s.logger.Println("shutting down")
			s.cleanup(nil)
			return
		}
	}
}

// Update fetches and processes the latest stable swap events from the Hydration API.
//
// This method retrieves swap events from the Hydration API, decodes them into
// `StableSwapEvent` objects, and then processes the events by matching them to
// pools retrieved from the database. For each matching pool, it generates a
// `dia.Trade` object and sends it to the `chanTrades` channel for further handling.
//
// The function also includes retry logic when fetching data from the API and logs
// any errors that occur during the process.
//
// Error Handling:
//   - Returns an error if fetching swap events or decoding the API response fails.
//   - Logs and skips any pools that do not have the required asset volumes.
//   - Logs a warning if no events match a specific pool.
//
// Returns:
//   - An error if the API request or pool retrieval fails, otherwise `nil`.
func (s *HydrationScraper) Update() error {
	s.logger.Info("Fetching swap events...")
	// To test the scraper with a specific block
	// endpoint := fmt.Sprintf("%s/events/swap/0x4b4d1d9db6336fd124b7df7d54962137e70f60693633692b6e0b54d71650e4af", s.api)
	endpoint := fmt.Sprintf("%s/events/swap", s.api)

	resp, err := s.fetchWithRetry(endpoint, "application/json", 3)
	if err != nil {
		return fmt.Errorf("failed to fetch swap events after retries: %w", err)
	}
	defer resp.Body.Close()

	var events []hydrationhelper.HydrationSwapEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return fmt.Errorf("failed to decode swap events: %w", err)
	}

	pools, err := s.getDBPools()
	if err != nil {
		s.logger.WithError(err).Error("Failed to get database pools")
		return err
	}

	for _, pool := range pools {
		if len(pool.Assetvolumes) < 2 {
			s.logger.WithField("poolAddress", pool.Address).Error("Pool has fewer than 2 asset volumes")
			continue
		}

		poolEvents := s.filterPoolEvents(pool.Address, events)
		if len(poolEvents) == 0 {
			continue
		}

		for _, event := range poolEvents {
			diaTrade := s.handleTrade(pool, event, time.Now())
			s.chanTrades <- diaTrade
		}
	}

	return nil
}

// filterPoolEvents filters swap events that match the specified pool address.
//
// This method compares each event's input and output assets to the provided
// pool address and returns a list of events that match. It constructs two
// possible event addresses (one for each direction of the swap) and checks if
// either matches the given pool address. If a match is found, the event is
// added to the result set.
//
// Logging is performed for each matching event, including both the pool address
// and the constructed event addresses.
//
// Parameters:
//   - poolAddress: The address of the pool to filter events for.
//   - events: A slice of `HydrationSwapEvent` objects representing the swap events to be filtered.
//
// Returns:
//   - A slice of `HydrationSwapEvent` objects that match the provided pool address.
func (s *HydrationScraper) filterPoolEvents(poolAddress string, events []hydrationhelper.HydrationSwapEvent) []hydrationhelper.HydrationSwapEvent {
	var matchedEvents []hydrationhelper.HydrationSwapEvent

	// Split the poolAddress into its component asset IDs
	poolAssets := strings.Split(poolAddress, "_")

	for _, event := range events {
		// Check if both assetIn and assetOut are in the poolAssets slice
		if containsAsset(poolAssets, event.AssetIn) && containsAsset(poolAssets, event.AssetOut) {
			matchedEvents = append(matchedEvents, event)
			s.logger.WithFields(logrus.Fields{
				"poolAddress": poolAddress,
				"assetIn":     event.AssetIn,
				"assetOut":    event.AssetOut,
			}).Info("Matched event for pool")
		}
	}

	return matchedEvents
}

// Helper function to check if an asset is in the pool assets slice
func containsAsset(poolAssets []string, asset string) bool {
	for _, a := range poolAssets {
		if a == asset {
			return true
		}
	}
	return false
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

	var quoteToken dia.Asset
	var baseToken dia.Asset

	var indexIn, indexOut int

	// Use the simple token ID instead of concatenated address
	assetInAddress := event.AssetIn
	if pool.Assetvolumes[0].Asset.Address == assetInAddress {
		indexIn = 0
		indexOut = 1
	} else {
		indexIn = 1
		indexOut = 0
	}

	quoteToken = pool.Assetvolumes[indexIn].Asset
	baseToken = pool.Assetvolumes[indexOut].Asset

	decimalsIn = int64(pool.Assetvolumes[indexIn].Asset.Decimals)
	decimalsOut = int64(pool.Assetvolumes[indexOut].Asset.Decimals)

	amountIn, _ := utils.StringToFloat64(event.AmountIn, decimalsIn)
	amountOut, _ := utils.StringToFloat64(event.AmountOut, decimalsOut)

	volume = -amountIn
	price = amountIn / amountOut

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

// getDBPools retrieves all pools from the database associated with the current exchange.
//
// Returns:
//   - []dia.Pool: A slice of `dia.Pool` objects representing the pools retrieved from the database.
//   - error: An error if the database query fails, otherwise `nil`.
func (s *HydrationScraper) getDBPools() ([]dia.Pool, error) {
	return s.db.GetAllPoolsExchange(s.exchangeName, 0)
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

// cleanup handles the shutdown procedure.
func (s *HydrationScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	s.ticker.Stop()

	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone)
}

// Close gracefully shuts down the HydrationScraper.
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
