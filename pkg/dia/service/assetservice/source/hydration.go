package source

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	hydrationhelper "github.com/diadata-org/diadata/pkg/dia/helpers/hydration-helper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

// HydrationAssetSource asset collector object - which serves assetCollector command
type HydrationAssetSource struct {
	// client - interaction with hydration REST API services
	hydrationClient *hydrationhelper.HydrationClient
	// channel to store received asset info
	assetChannel chan dia.Asset
	// channel which informs about work is finished
	doneChannel chan bool
	// blockchain name
	blockchain string
	// exchange name
	exchange string
	// DB connector to interact with databases
	relDB *models.RelDB
	// logs all events here
	logger *logrus.Entry
	// swap contracts count limitation in hydration REST API
	swapContractsLimit int
	sleepTimeout       time.Duration
	targetSwapContract string
	api                string
}

// NewHydrationAssetSource creates object to get hydration assets
// ENV values:
//
//	 	HYDRATION_ASSETS_SLEEP_TIMEOUT - (optional,millisecond), make timeout between API calls, default "hydrationhelper.DefaultSleepBetweenContractCalls" value
//		HYDRATION_SWAP_CONTRACTS_LIMIT - (optional, int), limit to get swap contact addresses, default "hydrationhelper.DefaultSwapContractsLimit" value
//		HYDRATION_TARGET_SWAP_CONTRACT - (optional, string), useful for debug, default = ""
//		HYDRATION_DEBUG - (optional, bool), make stdout output with hydration client http call, default = false
func NewHydrationAssetSource(exchange dia.Exchange, relDB *models.RelDB) *HydrationAssetSource {
	sleepBetweenContractCalls := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(strings.ToUpper(exchange.Name)+"_SLEEP_TIMEOUT", hydrationhelper.DefaultSleepBetweenContractCalls),
	)
	swapContractsLimit := utils.GetenvInt(
		strings.ToUpper(exchange.Name)+"_SWAP_CONTRACTS_LIMIT",
		hydrationhelper.DefaultSwapContractsLimit,
	)
	targetSwapContract := utils.Getenv(strings.ToUpper(exchange.Name)+"_TARGET_SWAP_CONTRACT", "")
	isDebug := utils.GetenvBool(strings.ToUpper(exchange.Name)+"_DEBUG", false)
	var (
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
	)
	hydrationClient := hydrationhelper.NewHydrationClient(
		log.WithContext(context.Background()).WithField("context", "HydrationClient"),
		sleepBetweenContractCalls,
		isDebug,
	)
	logger := log.
		WithContext(context.Background()).
		WithField("service", "assetCollector").
		WithField("network", "Hydration")
	apiURL := utils.Getenv(strings.ToUpper(exchange.Name)+"_API_URL", "http://localhost:3000/hydration/v1")
	scraper := &HydrationAssetSource{
		hydrationClient:    hydrationClient,
		assetChannel:       assetChannel,
		doneChannel:        doneChannel,
		blockchain:         "Hydration",
		exchange:           "Hydration",
		relDB:              relDB,
		logger:             logger,
		swapContractsLimit: swapContractsLimit,
		sleepTimeout:       sleepBetweenContractCalls,
		targetSwapContract: targetSwapContract,
		api:                apiURL,
	}
	go scraper.fetchAssets()
	return scraper
}

// fetchAssets scrapes asset data and sends it through channels for processing.
//
// This method logs the start of the asset scraping process, fetches assets using
// the `scrapAssets` method, and handles any errors that occur during the scraping.
// If successful, it sends each scraped asset into the `assetChannel` for further
// processing. Once all assets have been sent, a signal is sent to the `doneChannel`
// indicating the completion of the process.
//
// No return values.
func (s *HydrationAssetSource) fetchAssets() {
	s.logger.Info("Scraping assets...")
	assets, err := s.scrapAssets()
	if err != nil {
		s.logger.Error("Error when scraping assets: ", err)
		return
	}
	for _, asset := range assets {
		s.assetChannel <- *asset
	}
	s.doneChannel <- true
}

// scrapAssets fetches asset metadata from the API and processes it into dia.Asset objects.
//
// This method performs an HTTP request to retrieve asset data from the Hydration API.
// It retries the request up to 3 times if necessary. After successfully fetching
// the data, it decodes the JSON response into a slice of `HydrationAssetMetadata` objects
// and parses them into `dia.Asset` objects using the `parseAssets` method.
//
// Parameters:
//   - None.
//
// Returns:
//   - []*dia.Asset: A slice of pointers to `dia.Asset` objects representing the processed assets.
//   - error: An error if the API request or JSON decoding fails, otherwise `nil`.
func (s *HydrationAssetSource) scrapAssets() ([]*dia.Asset, error) {
	endpoint := fmt.Sprintf("%s/assets", s.api)
	resp, err := s.fetchWithRetry(endpoint, "application/json", 3)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch swap events after retries: %w", err)
	}
	defer resp.Body.Close()
	var assets []hydrationhelper.HydrationAssetMetadata
	if err := json.NewDecoder(resp.Body).Decode(&assets); err != nil {
		return nil, fmt.Errorf("failed to decode swap events: %w", err)
	}
	return s.parseAssets(assets), nil
}
func (s *HydrationAssetSource) Asset() chan dia.Asset {
	return s.assetChannel
}
func (s *HydrationAssetSource) Done() chan bool {
	return s.doneChannel
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
func (s *HydrationAssetSource) fetchWithRetry(endpoint string, contentType string, retries int) (*http.Response, error) {
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
func (s *HydrationAssetSource) parseAssets(assetsMetadata []hydrationhelper.HydrationAssetMetadata) []*dia.Asset {
	assets := make([]*dia.Asset, 0, len(assetsMetadata))
	for _, assetMetadata := range assetsMetadata {
		asset := s.parseAsset(assetMetadata)
		assets = append(assets, asset)
	}
	return assets
}

// parseAsset converts a HydrationAssetMetadata object into a dia.Asset object.
//
// This method takes metadata about an asset, such as its name, symbol, and decimals,
// and converts it into a `dia.Asset` object. The blockchain name is added from the
// `HydrationAssetSource` instance, and the asset's address is constructed by appending
// the asset ID (in lowercase) to a fixed prefix.
//
// Parameters:
//   - assetMetadata: A `HydrationAssetMetadata` object containing the raw asset data.
//
// Returns:
//   - *dia.Asset: A pointer to the newly created `dia.Asset` object.
func (s *HydrationAssetSource) parseAsset(assetMetadata hydrationhelper.HydrationAssetMetadata) *dia.Asset {
	return &dia.Asset{
		Name:       assetMetadata.Name,
		Symbol:     assetMetadata.Symbol,
		Decimals:   uint8(assetMetadata.Decimals),
		Blockchain: "Hydration",
		Address:    assetMetadata.Id,
	}
}
