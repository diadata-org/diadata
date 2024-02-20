package historicalscrapers

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

const (
	cmcDateLimit = time.Duration(89 * 24 * time.Hour)
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

type CoinmarketcapScraper struct {
	datastore        *models.DB
	rdb              *models.RelDB
	quotationChannel chan models.AssetQuotation
	doneChannel      chan bool
	dateLimit        time.Duration
}

func NewCoinmarketcapScraper(rdb *models.RelDB, datastore *models.DB) (cmcScraper CoinmarketcapScraper) {
	cmcScraper.doneChannel = make(chan bool)
	cmcScraper.quotationChannel = make(chan models.AssetQuotation)
	cmcScraper.dateLimit = cmcDateLimit
	cmcScraper.datastore = datastore
	cmcScraper.rdb = rdb

	go func() {
		cmcScraper.FetchQuotations()
	}()
	return cmcScraper
}

func (s CoinmarketcapScraper) FetchQuotations() {

	log.Info("Starting historical quotes scraper for Coinmarketcap")

	ethAsset := dia.Asset{
		Symbol:     "ETH",
		Name:       "Ethereum",
		Address:    "0x0000000000000000000000000000000000000000",
		Decimals:   18,
		Blockchain: "Ethereum",
	}

	cmcAPIKey := utils.Getenv("CMC_API_KEY", "")

	// Get the timestamp of the last recorded quotation
	endtime := time.Now()
	starttime := endtime.Add(-s.dateLimit)
	lastTimestamp, err := s.rdb.GetLastHistoricalQuotationTimestamp(ethAsset)
	if err != nil {
		log.Infof("No last time found for %s. Initialize with historical data.", ethAsset.Symbol)
		quotes, err := fetchCMCQuotations(ethAsset.Symbol, starttime, endtime, cmcAPIKey)
		if err != nil {
			log.Errorf("Failed to fetch quotation from CMC API: %v", err)
		}

		for _, quote := range quotes.Data.Quotes {
			historicalQuote := models.AssetQuotation{
				Asset:  ethAsset,
				Price:  quote.FullQuote.Quote.Price,
				Source: "Coinmarketcap",
				Time:   quote.FullQuote.Quote.Timestamp,
			}
			s.quotationChannel <- historicalQuote
		}
		s.doneChannel <- true
	} else {
		// TO DO: Otherwise try to fetch quotations from our DB. If too far in the past, fetch again from CMC.
		for timestamp := lastTimestamp.Add(1 * time.Hour); timestamp.Before(endtime); timestamp = timestamp.Add(1 * time.Hour) {
			quote, err := s.datastore.GetAssetQuotation(ethAsset, time.Time{}, timestamp)
			if err != nil {
				log.Error("get asset quotation: ", err)
			}

			historicalQuote := models.AssetQuotation{
				Asset:  ethAsset,
				Price:  quote.Price,
				Source: quote.Source,
				Time:   quote.Time,
			}

			s.quotationChannel <- historicalQuote
		}
	}
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

func (s CoinmarketcapScraper) QuoteChannel() chan models.AssetQuotation {
	return s.quotationChannel
}

func (s CoinmarketcapScraper) Done() chan bool {
	return s.doneChannel
}
