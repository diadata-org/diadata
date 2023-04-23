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

type cgPriceUSD struct {
	PriceUSD float64 `json:"usd"`
}

type cgAPIResponse struct {
	Quote struct {
		CurrentPrice cgPriceUSD `json:"current_price"`
	} `json:"market_data"`
}

type CoingeckoScraper struct {
	datastore        *models.DB
	rdb              *models.RelDB
	quotationChannel chan models.AssetQuotation
	doneChannel      chan bool
	firstDate        string
}

const (
	firstDate       = "08-03-2023"
	cgDateLayout    = "02-01-2006"
	cgErrCountLimit = 5
	cgWaitSeconds   = 1
)

func NewCoingeckoScraper(rdb *models.RelDB, datastore *models.DB) (cgScraper CoingeckoScraper) {
	cgScraper.doneChannel = make(chan bool)
	cgScraper.quotationChannel = make(chan models.AssetQuotation)
	cgScraper.firstDate = firstDate
	cgScraper.datastore = datastore
	cgScraper.rdb = rdb

	go func() {
		cgScraper.FetchQuotations()
	}()
	return cgScraper
}

func (s CoingeckoScraper) FetchQuotations() {

	log.Info("Starting historical quotes scraper for Coingecko.")

	// Outlook: Coingecko id for asset and blockchain,address as input so we can run
	// the scraper for various assets.
	ethAsset := dia.Asset{
		Symbol:     "ETH",
		Name:       "Ethereum",
		Address:    "0x0000000000000000000000000000000000000000",
		Decimals:   18,
		Blockchain: "Ethereum",
	}

	// cgAPIKey := utils.Getenv("CG_API_KEY", "")
	endDate, err := time.Parse(cgDateLayout, time.Now().Format(cgDateLayout))
	if err != nil {
		log.Error("Normalize date: ", err)
	}
	// Fetch CG data up to @endDateCG. Afterwards, use our DB.
	var endDateCG time.Time
	var currentDate time.Time

	oldestQuotation, err := s.datastore.GetOldestQuotation(ethAsset)
	if err != nil {
		log.Error("Fetch oldest quotation: ", err)
		endDateCG = endDate
	} else {
		// We only need CG data up to the oldest quotation in our DB. Normalize date to cg format.
		endDateCG, err = time.Parse(cgDateLayout, oldestQuotation.Time.AddDate(0, 0, 1).Format(cgDateLayout))
		if err != nil {
			log.Error("Normalize date: ", err)
		}
	}

	// Get latest recorded timestamp in historicalquotation table.
	latestTimestamp, err := s.rdb.GetLastHistoricalQuotationTimestamp(ethAsset)
	if err != nil {
		log.Warn("fetch latestTimestamp: ", err)
		var errDate error
		currentDate, errDate = time.Parse(cgDateLayout, s.firstDate)
		if errDate != nil {
			log.Fatal("parse cg first date: ", errDate)
		}
		s.fetchCGPrices(endDateCG, currentDate, ethAsset)
		currentDate = endDateCG
	} else {
		currentDate = latestTimestamp.AddDate(0, 0, 1)
	}

	// In case the oldest quotation from DIA DB is not old enough, fetch remaining gap
	// from Coingecko...
	if currentDate.Before(endDateCG) {
		s.fetchCGPrices(endDateCG, currentDate, ethAsset)
		currentDate = endDateCG
	}
	// ... otherwise fetch quotations from DIA DB.
	for currentDate.Before(endDate) {
		quotation, err := s.datastore.GetAssetQuotation(ethAsset, currentDate)
		if err != nil {
			log.Fatal("Get asset quotation: ", err)
		}
		currentDate = currentDate.AddDate(0, 0, 1)
		s.quotationChannel <- *quotation
		time.Sleep(1 * time.Second)
	}
	s.doneChannel <- true

}

func (s *CoingeckoScraper) fetchCGPrices(endDate time.Time, currentDate time.Time, asset dia.Asset) {
	var errCount int
	for endDate.After(currentDate) {
		price, errQuot := fetchCGQuotation("ethereum", currentDate.Format(cgDateLayout), "")
		if errQuot != nil {
			log.Error("fetch CG quotation: ", errQuot)
			if errCount < cgErrCountLimit {
				time.Sleep(3 * time.Second)
				errCount++
				continue
			} else {
				log.Fatal("Repeated fail on CG rest API.")
			}
		} else {
			quotation := models.AssetQuotation{
				Asset:  asset,
				Price:  price,
				Time:   currentDate,
				Source: "Coingecko",
			}
			currentDate = currentDate.AddDate(0, 0, 1)
			s.quotationChannel <- quotation
			time.Sleep(cgWaitSeconds * time.Second)
		}
	}
}

func fetchCGQuotation(assetSymbol string, date string, apiKey string) (price float64, err error) {

	var url string
	if apiKey != "" {
		url = fmt.Sprintf("https://pro-api.coingecko.com/api/v3/coins/%s/history?date=%s&x_cg_pro_api_key=%s", assetSymbol, date, apiKey)
	} else {
		url = fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/history?date=%s", assetSymbol, date)
	}
	response, statusCode, err := utils.GetRequest(url)
	if err != nil {
		return
	}
	if statusCode != http.StatusOK {
		err = errors.New("non-200 status code from Coinmarketcap API")
		return
	}

	var quote cgAPIResponse
	if err := json.Unmarshal(response, &quote); err != nil {
		log.Error("Unmarshal response: ", err)
	}

	price = quote.Quote.CurrentPrice.PriceUSD
	return
}

func (s CoingeckoScraper) QuoteChannel() chan models.AssetQuotation {
	return s.quotationChannel
}

func (s CoingeckoScraper) Done() chan bool {
	return s.doneChannel
}
