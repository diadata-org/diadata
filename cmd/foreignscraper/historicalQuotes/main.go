package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"net/http"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
)

func fetchCMCQuotation(assetSymbol string, timestamp time.Time, apiKey string) (float64, error) {
	url := fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/historical?symbol=%s&interval=hourly&time_start=%d", assetSymbol, timestamp.Unix())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("X-CMC_PRO_API_KEY", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, errors.New("non-200 status code from Coinmarketcap API")
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	quoteData := result["data"].(map[string]interface{})
	quote := quoteData[assetSymbol].(map[string]interface{})
	priceData := quote["quote"].(map[string]interface{})
	usdData := priceData["USD"].(map[string]interface{})
	price, ok := usdData["price"].(float64)
	if !ok {
		return 0, errors.New("unable to extract price from API response")
	}

	return price, nil
}

func main() {
	log.Info("Starting historical quotes scraper")

	// Create an instance of the DIA datastore
	datastore, err := models.NewDataStoreWithoutRedis()
	if err != nil {
		log.Fatalf("Failed to create datastore: %v", err)
	}

	ethAsset := dia.Asset{
		Symbol:     "ETH",
		Name:       "Ethereum",
		Address:    "",
		Decimals:   18,
		Blockchain: "Ethereum",
	}

	cmcAPIKey := os.Getenv("CMC_API_KEY")

	ticker := time.NewTicker(1 * time.Hour)
	for range ticker.C {
		// Get the timestamp of the last recorded quotation
		lastTimestamp, err := datastore.GetLastHistoricalQuoteTimestamp(ethAsset.Symbol)
		if err != nil {
			log.Infof("No historical quotes found for %s, using current time", ethAsset.Symbol)
			lastTimestamp = time.Now().Add(-1 * time.Hour)
		}

		// Fetch hourly ETH quotations from the last timestamp until now
		for timestamp := lastTimestamp.Add(1 * time.Hour); timestamp.Before(time.Now()); timestamp = timestamp.Add(1 * time.Hour) {
			quote, err := datastore.GetAssetQuotation(ethAsset, timestamp)
			if err != nil {
				log.Infof("Quotation not found in our DB for timestamp %s, fetching from CMC API", timestamp)

				price, err := fetchCMCQuotation(ethAsset.Symbol, timestamp, cmcAPIKey)
				if err != nil {
					log.Errorf("Failed to fetch quotation from CMC API: %v", err)
					continue
				}

				quote = &models.AssetQuotation{
					Asset:  ethAsset,
					Price:  price,
					Time:   timestamp,
					Source: "Coinmarketcap",
				}
			}

			historicalQuote := models.HistoricalQuote{
				Symbol:    quote.Asset.Symbol,
				Price:     quote.Price,
				QuoteTime: quote.Time,
				Source:    quote.Source,
			}

			if err := datastore.SetHistoricalQuote(&historicalQuote); err != nil {
				log.Errorf("Failed to store historical quote: %v", err)
			} else {
				log.Infof("Stored historical quote for %s at %s", ethAsset.Symbol, quote.Time)
			}
		}
	}
}
