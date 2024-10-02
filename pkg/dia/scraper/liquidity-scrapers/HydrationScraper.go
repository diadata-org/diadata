package liquidityscrapers

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

type HydrationLiquidityScraper struct {
	logger                    *logrus.Entry
	poolChannel               chan dia.Pool
	doneChannel               chan bool
	blockchain                string
	exchangeName              string
	relDB                     *models.RelDB
	datastore                 *models.DB
	targetSwapContract        string
	swapContractsLimit        int
	handlerType               string
	sleepBetweenContractCalls time.Duration
	apiURL                    string
}

// NewHydrationLiquidityScraper returns a new HydrationLiquidityScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
// ENV values:
//
//	 	BIFROST_SLEEP_TIMEOUT - (optional,millisecond), make timeout between API calls, default "Hydrationhelper.DefaultSleepBetweenContractCalls" value
//		BIFROST_TARGET_SWAP_CONTRACT - (optional, string), useful for debug, default = ""
//		BIFROST_DEBUG - (optional, bool), make stdout output with Hydration client http call, default = false
func NewHydrationLiquidityScraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) *HydrationLiquidityScraper {
	targetSwapContract := utils.Getenv(
		strings.ToUpper(exchange.Name)+"_TARGET_SWAP_CONTRACT",
		"",
	)
	sleepBetweenContractCalls := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(
			strings.ToUpper(exchange.Name)+"_SLEEP_TIMEOUT",
			hydrationhelper.DefaultSleepBetweenContractCalls,
		),
	)
	swapContractsLimit := utils.GetenvInt(
		strings.ToUpper(exchange.Name)+"_SWAP_CONTRACTS_LIMIT",
		hydrationhelper.DefaultSwapContractsLimit,
	)

	var (
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		scraper     *HydrationLiquidityScraper
	)

	apiURL := utils.Getenv(strings.ToUpper(exchange.Name)+"_API_URL", "http://localhost:3000/hydration/v1")

	scraper = &HydrationLiquidityScraper{
		poolChannel:               poolChannel,
		doneChannel:               doneChannel,
		exchangeName:              "Hydration",
		blockchain:                "Hydration",
		relDB:                     relDB,
		datastore:                 datastore,
		targetSwapContract:        targetSwapContract,
		swapContractsLimit:        swapContractsLimit,
		handlerType:               "liquidity",
		sleepBetweenContractCalls: sleepBetweenContractCalls,
		apiURL:                    apiURL,
	}
	scraper.logger = logrus.
		New().
		WithContext(context.Background()).
		WithField("handlerType", scraper.handlerType).
		WithField("context", "HydrationLiquidityScraper")

	go scraper.fetchPools()

	return scraper
}

// fetches liquidity pool data from the API and processes it into dia.Pool objects.
//
// This method performs an HTTP GET request to the API endpoint to retrieve pool metadata.
// It retries the request up to 3 times in case of failure. Once the response is successfully
// fetched, it decodes the JSON response into a slice of `HydrationPoolMetadata` objects and
// processes them into `dia.Pool` objects using the `parseAssets` method.
//
//
// Returns:
//   - []*dia.Pool: A slice of pointers to `dia.Pool` objects representing the parsed liquidity pools.
//   - error: An error if the API request or JSON decoding fails, otherwise `nil`.
func (s *HydrationLiquidityScraper) getPools() ([]*dia.Pool, error) {
	endpoint := fmt.Sprintf("%s/pools", s.apiURL)

	resp, err := s.fetchWithRetry(endpoint, "application/json", 3)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch swap events after retries: %w", err)
	}
	defer resp.Body.Close()

	var assets []hydrationhelper.HydrationPoolMetada
	if err := json.NewDecoder(resp.Body).Decode(&assets); err != nil {
		return nil, fmt.Errorf("failed to decode swap events: %w", err)
	}
	return s.parseAssets(assets), nil
}

// converts a slice of HydrationPoolMetadata into dia.Pool objects.
//
// This method processes liquidity pool metadata by retrieving associated assets from the database,
// calculating their volumes, and constructing `dia.Pool` objects. Each pool is assigned a unique
// address based on its tokens, and the asset volumes are populated with the appropriate balance
// and USD equivalent for each asset. The function skips any pools that do not contain exactly two
// asset types.
//
// Parameters:
//   - poolMetadata: A slice of `HydrationPoolMetadata` containing the raw metadata of liquidity pools.
//
// Returns:
//   - []*dia.Pool: A slice of pointers to `dia.Pool` objects representing the processed pools.
func (s *HydrationLiquidityScraper) parseAssets(poolMetadata []hydrationhelper.HydrationPoolMetada) []*dia.Pool {
	pools := make([]*dia.Pool, 0)

	for _, metadataPair := range poolMetadata {
		pair := &dia.Pool{
			Assetvolumes: []dia.AssetVolume{},
			Time:         time.Now(),
		}

		var tokenNames []string

		for _, token := range metadataPair.Tokens {
			assetKey := token.ID
			dbAsset, err := s.relDB.GetAsset(assetKey, s.blockchain)
			if err != nil {
				s.logger.WithError(err).Warn("Failed to GetAsset with key: ", assetKey)
				continue
			}

			balance, _ := utils.StringToFloat64(token.Balance, int64(token.Decimals))
			usdBalance, _ := utils.StringToFloat64(token.UsdBalance, int64(6))

			pair.Assetvolumes = append(pair.Assetvolumes, dia.AssetVolume{
				Index:     uint8(token.Index),
				Asset:     dbAsset,
				Volume:    balance,
				VolumeUSD: usdBalance,
			})

			tokenNames = append(tokenNames, strings.ToLower(token.ID))
		}

		if len(pair.Assetvolumes) < 2 {
			s.logger.Warn("Found less than 2 asset types for the pool")
			continue
		}

		pair.Address = strings.Join(tokenNames, "_")
		pair.Exchange = dia.Exchange{Name: s.exchangeName}
		pair.Blockchain = dia.BlockChain{Name: s.blockchain}

		pools = append(pools, pair)

	}
	return pools
}

// performs an HTTP GET request to the specified endpoint with retries.
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
func (s *HydrationLiquidityScraper) fetchWithRetry(endpoint string, contentType string, retries int) (*http.Response, error) {
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

func (s *HydrationLiquidityScraper) fetchPools() {
	// Fetch all pair tokens pool entries from api
	pools, err := s.getPools()
	if err != nil {
		s.logger.WithError(err).Error("failed to GetAllPoolAssets")
	}
	s.logger.Infof("Found %d pools", len(pools))

	for _, pool := range pools {
		s.poolChannel <- *pool
	}

	s.doneChannel <- true
}

func (s *HydrationLiquidityScraper) Pool() chan dia.Pool {
	return s.poolChannel
}

func (s *HydrationLiquidityScraper) Done() chan bool {
	return s.doneChannel
}
