package foreignscrapers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

const (
	sourceTwelvedata = "TwelveData"
)

var (
	twelvedataUpdateSeconds int64
	twelvedataApiBaseString = "https://api.twelvedata.com/"
)

type twelvedataAPIFXResponse struct {
	Symbol    string  `json:"symbol"`
	Rate      float64 `json:"rate"`
	Timestamp int64   `json:"timestamp"`
}

type twelvedataAPIStockResponse struct {
	Price string `json:"price"`
}

type TwelvedataScraper struct {
	ticker                 *time.Ticker
	foreignScrapper        ForeignScraper
	twelvedataStockSymbols []string
	twelvedataFXTickers    []string
	apiKey                 string
}

func init() {
	var err error
	twelvedataUpdateSeconds, err = strconv.ParseInt(utils.Getenv("UPDATE_SECONDS", "3600"), 10, 64)
	if err != nil {
		log.Error("Parse UPDATE_SECONDS: ", err)
	}
}

func NewTwelvedataScraper(datastore models.Datastore) *TwelvedataScraper {

	foreignScrapper := ForeignScraper{
		shutdown:      make(chan nothing),
		error:         nil,
		datastore:     datastore,
		chanQuotation: make(chan *models.ForeignQuotation),
	}
	s := &TwelvedataScraper{
		ticker:                 time.NewTicker(time.Duration(twelvedataUpdateSeconds) * time.Second),
		foreignScrapper:        foreignScrapper,
		twelvedataStockSymbols: strings.Split(utils.Getenv("STOCK_SYMBOLS", ""), ","),
		twelvedataFXTickers:    strings.Split(utils.Getenv("FX_TICKERS", ""), ","),
		apiKey:                 utils.Getenv("TWELVEDATA_API_KEY", ""),
	}
	log.Info("apiKey: ", s.apiKey)

	go s.mainLoop()

	return s

}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *TwelvedataScraper) mainLoop() {

	// Initial run.
	err := scraper.UpdateQuotation()
	if err != nil {
		log.Error(err)
	}

	// Update every @twelvedataUpdateSeconds.
	for {
		select {
		case <-scraper.ticker.C:
			err := scraper.UpdateQuotation()
			if err != nil {
				log.Error(err)
			}
		case <-scraper.foreignScrapper.shutdown: // user requested shutdown
			log.Printf("twelvedatascraper shutting down")
			return
		}
	}

}

// Update retrieves new coin information from the twelvedata API and stores it to influx
func (scraper *TwelvedataScraper) UpdateQuotation() error {

	log.Printf("Executing stock data update")
	for _, symbol := range scraper.twelvedataStockSymbols {
		quotation, err := scraper.getTwelveStockData(symbol)
		if err != nil {
			log.Error("getTwelveStockData: ", err)
		}
		price, err := strconv.ParseFloat(quotation.Price, 64)
		if err != nil {
			return err
		}

		foreignQuotation := models.ForeignQuotation{
			Symbol: symbol,
			Price:  price,
			Source: sourceTwelvedata,
			Time:   time.Now(),
		}
		scraper.foreignScrapper.chanQuotation <- &foreignQuotation
	}

	log.Printf("Executing fx data update")
	for _, ticker := range scraper.twelvedataFXTickers {
		quotation, err := scraper.getTwelveFXData(ticker)
		if err != nil {
			log.Error("getTwelveFXData: ", err)
		}

		foreignQuotation := models.ForeignQuotation{
			Symbol: ticker,
			Price:  quotation.Rate,
			Source: sourceTwelvedata,
			Time:   time.Unix(quotation.Timestamp, 0),
		}
		scraper.foreignScrapper.chanQuotation <- &foreignQuotation
	}

	return nil

}

func (scraper *TwelvedataScraper) GetQuoteChannel() chan *models.ForeignQuotation {
	return scraper.foreignScrapper.chanQuotation
}

func (scraper *TwelvedataScraper) getTwelveFXData(symbol string) (fxRate twelvedataAPIFXResponse, err error) {
	var response []byte

	apiURL := twelvedataApiBaseString + "exchange_rate?symbol=" + symbol + "&apikey=" + scraper.apiKey
	response, _, err = utils.GetRequest(apiURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &fxRate)
	return
}

func (scraper *TwelvedataScraper) getTwelveStockData(symbol string) (stockPrice twelvedataAPIStockResponse, err error) {
	var response []byte

	apiURL := twelvedataApiBaseString + "price?symbol=" + symbol + "&apikey=" + scraper.apiKey
	response, _, err = utils.GetRequest(apiURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &stockPrice)
	return
}

// Close closes any existing API connections
func (scraper *TwelvedataScraper) Close() error {
	if scraper.foreignScrapper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.foreignScrapper.shutdown)
	<-scraper.foreignScrapper.shutdownDone
	scraper.foreignScrapper.errorLock.RLock()
	defer scraper.foreignScrapper.errorLock.RUnlock()
	return scraper.foreignScrapper.error
}
