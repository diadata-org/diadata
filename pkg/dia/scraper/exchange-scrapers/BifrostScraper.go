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
	bifrosthelper "github.com/diadata-org/diadata/pkg/dia/helpers/bifrost-helper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gorilla/websocket"
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
	//todo: move common code to a common helper
	wsApi        *bifrosthelper.SubstrateEventHelper
	exchangeName string
	blockchain   string
	wsClient     *websocket.Conn
}

// NewBifrostScraper initializes and returns a new BifrostScraper instance.
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
//   - *BifrostScraper: A pointer to the newly created BifrostScraper instance.
//
// Behavior:
//   - The function calculates the refresh delay for polling trade data by checking an environment variable
//     named "<EXCHANGE_NAME>_REFRESH_DELAY". If the variable is not set, it defaults to 10000 milliseconds (10 seconds).
//   - It initializes the BifrostScraper with a shutdown mechanism, a ticker for periodic updates, a channel for trade data,
//     and a logging mechanism.
//   - If the 'scrape' flag is true, the scraper starts its main scraping loop in a separate goroutine.
func NewBifrostScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BifrostScraper {
	//todo remove this
	refreshDelay := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(strings.ToUpper(exchange.Name)+"_REFRESH_DELAY", 10000),
	)
	logger := logrus.
		New().
		WithContext(context.Background()).
		WithField("context", "BifrostScraper")

	wsApi, err := bifrosthelper.NewSubstrateEventHelper("wss://bifrost-polkadot.api.onfinality.io/public-ws", logger)
	if err != nil {
		logrus.WithError(err).Error("Failed to create Bifrost Substrate event helper")
		return nil
	}

	s := &BifrostScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		//todo: remove this
		ticker:       time.NewTicker(refreshDelay),
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

// func (s *BifrostScraper) connectWebSocket() error {
// 	var err error
// 	s.wsClient, _, err = websocket.DefaultDialer.Dial(s.wsApi, nil)
// 	if err != nil {
// 		return fmt.Errorf("failed to connect to websocket: %w", err)
// 	}
// 	s.logger.Info("Connected to Bifrost WebSocket")
// 	return nil
// }

// func (s *BifrostScraper) connectWebSocket() error {
// 	// Parse the base WebSocket URL
// 	u, err := url.Parse(s.wsApi)
// 	if err != nil {
// 		return fmt.Errorf("failed to parse WebSocket URL: %w", err)
// 	}

// 	// Add query parameters
// 	query := u.Query()
// 	query.Set("method", "TokenSwapped")
// 	query.Set("section", "stableAsset")
// 	u.RawQuery = query.Encode()

// 	// Connect to the WebSocket with the updated URL
// 	s.wsClient, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
// 	if err != nil {
// 		return fmt.Errorf("failed to connect to websocket: %w", err)
// 	}
// 	s.logger.Info("Connected to Bifrost WebSocket")
// 	return nil
// }

// func (s *BifrostScraper) mainLoop() {
// 	if err := s.connectWebSocket(); err != nil {
// 		s.logger.Error(err)
// 		return
// 	}
// 	defer s.wsClient.Close()

// 	for {
// 		select {
// 		case <-s.shutdown:
// 			s.logger.Println("shutting down")
// 			s.cleanup(nil)
// 			return
// 		default:
// 			s.listenForTrades()
// 		}
// 	}
// }

func (s *BifrostScraper) mainLoop() {
	//go s.wsApi.ListenForNewBlocks(s.processEvents)
	s.logger.Info("Listening for new blocks")
	go s.wsApi.ListenForSpecificBlock(5697382, s.processEvents)
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

func (s *BifrostScraper) processEvents(events *[]bifrosthelper.EventSellExecuted) {
	s.logger.Info("Processing events")
	for _, event := range *events {

		assetIn := fmt.Sprint(event.InputAsset)
		assetOut := fmt.Sprint(event.OutputAsset)
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

type TestJson struct {
	Method  string `json:"method"`
	Section string `json:"section"`
}

// new update
func (s *BifrostScraper) listenForTrades() {
	// jsontest := TestJson{Method: "TokenSwapped", Section: "stableAsset"}
	// err := s.wsClient.WriteJSON(jsontest)
	// err := s.wsClient.WriteMessage(websocket.TextMessage, []byte(`{"method":"TokenSwapped","section":"stableAsset"}`))
	// s.logger.Info(err)
	// _, message, err := s.wsClient.ReadMessage()
	// if err != nil {
	// 	s.logger.Error("Error reading message from websocket: ", err)
	// 	return
	// }

	// var events []bifrosthelper.StableSwapEvent
	// // message := &StableSwapEvent{}
	// // err := s.wsClient.ReadJSON(&message)
	// // if err != nil {
	// // 	s.logger.Error("Error reading message from websocket: ", err)
	// // 	// break
	// // }
	// if err := json.Unmarshal(message, &events); err != nil {
	// 	s.logger.Error("Error unmarshalling websocket message: ", err)
	// 	return
	// }

	// // Check if the event is a Bifrost-specific event
	// // if event.Type != "SellExecuted" || event.PoolID != 101 { // Example checks
	// // 	return
	// // }
	// //pools, err := s.db.GetPoolsByAssetPair(event.AssetIn, event.AssetOut, s.exchangeName)
	// // pools, err := s.getDBPools()
	// // if err != nil {
	// // 	s.logger.WithError(err).Error("Failed to get database pools")
	// // 	return
	// // }

	// if err != nil {
	// 	s.logger.WithField("events", events).Debug("Fetched swap events")
	// 	return
	// }

	// for _, event := range events {
	// 	diaTrade := s.handleTrade(pool, event, time.Now())
	// 	s.chanTrades <- diaTrade
	// }
	// // for _, pool := range pools {
	// // 	if len(pool.Assetvolumes) < 2 {
	// // 		s.logger.WithField("poolAddress", pool.Address).Error("Pool has fewer than 2 asset volumes")
	// // 		continue
	// // 	}

	// // 	diaTrade := s.handleTrade(pool, event, time.Now())
	// // 	s.chanTrades <- diaTrade
	// // }

	// for _, pool := range pools {
	// 	if len(pool.Assetvolumes) != 2 {
	// 		s.logger.WithField("poolAddress", pool.Address).Error("Pool is missing required asset volumes")
	// 		continue
	// 	}

	// 	poolEvents := s.filterPoolEvents(pool.Address, events)
	// 	if len(poolEvents) == 0 {
	// 		continue
	// 	}

	// 	for _, event := range poolEvents {
	// 		s.logger.WithField("poolAddress", pool.Address).Info("Processing event for pool")
	// 		diaTrade := s.handleTrade(pool, event, time.Now())
	// 		s.chanTrades <- diaTrade
	// 	}
	// }
}

// Update fetches and processes the latest stable swap events from the Bifrost API.
//
// This method retrieves swap events from the Bifrost API, decodes them into
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
func (s *BifrostScraper) Update() error {
	// To test the scraper with a specific block
	endpoint := fmt.Sprintf("%s/events/swap/0x5d184786631f977395636a98b231358b9a82b63c235c801c170fb7e532cb839b")
	//endpoint := fmt.Sprintf("%s/events/swap/0x5d184786631f977395636a98b231358b9a82b63c235c801c170fb7e532cb839b", s.api)
	//endpoint := fmt.Sprintf("%s/events/swap", s.api)

	resp, err := s.fetchWithRetry(endpoint, "application/json", 3)
	if err != nil {
		return fmt.Errorf("failed to fetch swap events after retries: %w", err)
	}
	defer resp.Body.Close()
	fmt.Println("Resp.body", resp.Body)
	s.logger.Debug("Resp.body", resp.Body)
	var events []bifrosthelper.StableSwapEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return fmt.Errorf("failed to decode swap events: %w", err)
	}
	fmt.Println("Events", events)
	s.logger.WithField("events", events).Debug("Fetched swap events")
	pools, err := s.getDBPools()
	if err != nil {
		s.logger.WithError(err).Error("Failed to get database pools")
		return err
	}

	if len(pools) == 0 {
		s.logger.Warn("No pools found in the database")
		return nil
	}

	// // Process each pool looking for matching events and processing
	// for _, pool := range pools {
	// 	if len(pool.Assetvolumes) != 2 {
	// 		s.logger.WithField("poolAddress", pool.Address).Error("Pool is missing required asset volumes")
	// 		continue
	// 	}

	// 	poolEvents := s.filterPoolEvents(pool.Address, events)
	// 	if len(poolEvents) == 0 {
	// 		continue
	// 	}

	// 	for _, event := range poolEvents {
	// 		s.logger.WithField("poolAddress", pool.Address).Info("Processing event for pool")
	// 		diaTrade := s.handleTrade(pool, event, time.Now())
	// 		s.chanTrades <- diaTrade
	// 	}
	// }

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
//   - events: A slice of `StableSwapEvent` objects representing the swap events to be filtered.
//
// Returns:
//   - A slice of `StableSwapEvent` objects that match the provided pool address.
func (s *BifrostScraper) filterPoolEvents(poolAddress string, events []bifrosthelper.StableSwapEvent) []bifrosthelper.StableSwapEvent {
	var matchedEvents []bifrosthelper.StableSwapEvent

	for _, event := range events {
		if event.PoolID != poolAddress {
			continue
		}

		matchedEvents = append(matchedEvents, event)
	}

	return matchedEvents
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
func (s *BifrostScraper) fetchWithRetry(endpoint string, contentType string, retries int) (*http.Response, error) {
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

func (s *BifrostScraper) handleTrade(pool dia.Pool, event bifrosthelper.EventSellExecuted, time time.Time) *dia.Trade {
	var volume, price float64
	var baseToken, quoteToken dia.Asset
	var decimalsIn, decimalsOut int64

	//if fmt.Sprint(event.InputAsset) == pool.Assetvolumes[0].Asset.Symbol {
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

	amountIn, _ := utils.StringToFloat64(fmt.Sprint(event.InputAmount), decimalsIn)
	amountOut, _ := utils.StringToFloat64(fmt.Sprint(event.OutputAmount), decimalsOut)

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
