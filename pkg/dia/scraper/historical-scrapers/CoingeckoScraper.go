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
	dateLimit        time.Duration
}

func NewCoingeckoScraper(rdb *models.RelDB, datastore *models.DB) (cgScraper CoingeckoScraper) {
	cgScraper.doneChannel = make(chan bool)
	cgScraper.quotationChannel = make(chan models.AssetQuotation)
	cgScraper.dateLimit = cmcDateLimit
	cgScraper.datastore = datastore
	cgScraper.rdb = rdb

	go func() {
		cgScraper.FetchQuotations()
	}()
	return cgScraper
}

func (s CoingeckoScraper) FetchQuotations() {

	log.Info("Starting historical quotes scraper for Coingecko.")

	ethAsset := dia.Asset{
		Symbol:     "ETH",
		Name:       "Ethereum",
		Address:    "0x0000000000000000000000000000000000000000",
		Decimals:   18,
		Blockchain: "Ethereum",
	}

	// cgAPIKey := utils.Getenv("CG_API_KEY", "")
	endtime := time.Now()

	latestTimestamp, err := s.rdb.GetLastHistoricalQuotationTimestamp(ethAsset)
	if err != nil {
		// TO DO: Fetch historical data from CG
		s.doneChannel <- true
	}

	oldestQuotation, err := s.datastore.GetOldestQuotation(ethAsset)
	if err != nil {
		log.Fatal("Fetch oldest quotation: ", err)
	}
	for latestTimestamp.Before(oldestQuotation.Time) {
		// TO DO: Fetch from CG API
		latestTimestamp.AddDate(0, 0, 1)
	}
	for latestTimestamp.Before(endtime) {
		// TO DO: Fetch from DIA DB
		latestTimestamp.AddDate(0, 0, 1)
	}
	s.doneChannel <- true

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
		log.Error("unmarshal response: ", err)
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
