package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"net/http"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
)

type cmcQuote struct {
	Quote struct {
		Price     float64   `json:"price"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"USD"`
}

type cmcFullQuote struct {
	FullQuote cmcQuote  `json:"quote"`
	Timestamp time.Time `json:"timestamp"`
}

type cmcQuotes struct {
	Quotes []cmcFullQuote `json:"quotes"`
	Symbol string         `json:"symbol"`
}

type cmcResponse struct {
	Data cmcQuotes `json:"data"`
}

func fetchCMCQuotations(assetSymbol string, starttime time.Time, endtime time.Time, apiKey string) (response cmcResponse, err error) {
	log.Infof("starttime -- endtime: %v -- %v ", starttime, endtime)

	url := fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/historical?symbol=%s&interval=hourly&time_start=%d&time_end=%d", assetSymbol, starttime.Unix(), endtime.Unix())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("X-CMC_PRO_API_KEY", apiKey)

	resp, statusCode, err := utils.HTTPRequest(req)
	if err != nil {
		return
	}
	if statusCode != http.StatusOK {
		err = errors.New("non-200 status code from Coinmarketcap API")
		return
	}

	if err := json.Unmarshal(resp, &response); err != nil {
		log.Error("")
	}

	return
}

func main() {

	// TO DO: Switch historical quotes to full assets.
	log.Info("Starting historical quotes scraper")

	// Create an instance of the DIA datastore
	datastore, err := models.NewDataStoreWithoutRedis()
	if err != nil {
		log.Fatalf("Failed to create datastore: %v", err)
	}

	ethAsset := dia.Asset{
		Symbol:     "ETH",
		Name:       "Ethereum",
		Address:    "0x0000000000000000000000000000000000000000",
		Decimals:   18,
		Blockchain: "Ethereum",
	}

	cmcAPIKey := utils.Getenv("CMC_API_KEY", "676d0014-2760-485c-8e6d-3de41396cb48")

	// Get the timestamp of the last recorded quotation
	lastTimestamp, err := datastore.GetLastHistoricalQuoteTimestamp(ethAsset.Symbol)
	if err != nil {
		log.Infof("No historical quotes found for %s, using current time", ethAsset.Symbol)
		endtime := time.Now()
		starttime := endtime.AddDate(0, 0, -89)
		quotes, err := fetchCMCQuotations(ethAsset.Symbol, starttime, endtime, cmcAPIKey)
		if err != nil {
			log.Errorf("Failed to fetch quotation from CMC API: %v", err)
		}
		for _, quote := range quotes.Data.Quotes {
			historicalQuote := models.HistoricalQuote{
				Symbol:    ethAsset.Symbol,
				Price:     quote.FullQuote.Quote.Price,
				QuoteTime: quote.FullQuote.Quote.Timestamp,
				Source:    "Coinmarketcap",
			}
			if err := datastore.SetHistoricalQuote(&historicalQuote); err != nil {
				log.Errorf("Failed to store historical quote: %v", err)
			} else {
				log.Infof("Stored historical quote for %s at %s", ethAsset.Symbol, quote.FullQuote.Quote.Timestamp)
			}
		}
	}

	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {

		// Fetch hourly ETH quotations from the last timestamp until now
		for timestamp := lastTimestamp.Add(1 * time.Hour); timestamp.Before(time.Now()); timestamp = timestamp.Add(1 * time.Hour) {
			quote, err := datastore.GetAssetQuotation(ethAsset, timestamp)
			if err != nil {
				log.Error("get asset quotation: ", err)
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
