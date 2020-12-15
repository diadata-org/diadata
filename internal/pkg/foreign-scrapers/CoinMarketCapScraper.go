package foreignscrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

const (
	coinMarketCapRefreshDelay = time.Second * 60 * 2
	coinMarketCapsource       = "CoinMarketCap"
)

type Quote struct {
	Price            float64 `json:"price"`
	Volume24H        float64 `json:"volume_24h"`
	PercentChange1H  float64 `json:"percent_change_1h"`
	PercentChange24H float64 `json:"percent_change_24h"`
	PercentChange7D  float64 `json:"percent_change_7d"`
	MarketCap        float64 `json:"market_cap"`
	LastUpdated      string  `json:"last_updated"`
}
type CoinMarketCapListing struct {
	Data []struct {
		ID                int              `json:"id"`
		Name              string           `json:"name"`
		Symbol            string           `json:"symbol"`
		Slug              string           `json:"slug"`
		CmcRank           int              `json:"cmc_rank"`
		NumMarketPairs    int              `json:"num_market_pairs"`
		CirculatingSupply float64          `json:"circulating_supply"`
		TotalSupply       float64          `json:"total_supply"`
		MaxSupply         float64          `json:"max_supply"`
		LastUpdated       string           `json:"last_updated"`
		DateAdded         string           `json:"date_added"`
		Tags              []string         `json:"tags"`
		Platform          interface{}      `json:"platform"`
		Quote             map[string]Quote `json:"quote"`
	} `json:"data"`
	Status struct {
		Timestamp    time.Time `json:"timestamp"`
		ErrorCode    int       `json:"error_code"`
		ErrorMessage string    `json:"error_message"`
		Elapsed      int       `json:"elapsed"`
		CreditCount  int       `json:"credit_count"`
	} `json:"status"`
}

type CoinMarketCapScraper struct {
	ticker          *time.Ticker
	foreignScrapper ForeignScraper
}

func NewCoinMarketCapScraper(datastore models.Datastore) *CoinMarketCapScraper {

	foreignScrapper := ForeignScraper{
		shutdown:      make(chan nothing),
		error:         nil,
		datastore:     datastore,
		chanQuotation: make(chan *models.ForeignQuotation),
	}
	s := &CoinMarketCapScraper{
		ticker:          time.NewTicker(coinMarketCapRefreshDelay),
		foreignScrapper: foreignScrapper,
	}
	go s.mainLoop()

	return s

}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *CoinMarketCapScraper) mainLoop() {
	for true {
		select {
		case <-scraper.ticker.C:
			scraper.UpdateQuotation()
		case <-scraper.foreignScrapper.shutdown: // user requested shutdown
			log.Printf("CoinMarketCapscraper shutting down")
			return
		}
	}

}

// Update retrieves new coin information from the CoinMarketCap API and stores it to influx
func (scraper *CoinMarketCapScraper) UpdateQuotation() error {
	log.Printf("Executing CoinMarketCapScraper update")

	listing, err := getCoinMarketCapData()

	if err != nil {
		log.Errorln("Error getting data from CoinMarketCap", err)
		return err
	}

	for _, coin := range listing.Data {

		// TO DO: normalize symbol instead of all upper
		coin.Symbol = strings.ToUpper(coin.Symbol)

		// Parse last updated timestamp
		layout := "2006-01-02T15:04:05.000Z"
		timestamp, err := time.Parse(layout, coin.LastUpdated)
		if err != nil {
			log.Errorln("error parsing time")
		}

		// get ITIN if available in redis
		itin, err := scraper.foreignScrapper.datastore.GetItinBySymbol(coin.Symbol)
		if err != nil {
			// log.Errorf("error: no ITIN available for %s \n", coin.Symbol)
			itin = dia.ItinToken{}
		}

		// Get yesterday's price from influx if available
		priceYesterday, err := scraper.foreignScrapper.datastore.GetForeignPriceYesterday(coin.Symbol, source)
		if err != nil {
			priceYesterday = 0
		}
		usdQuote, ok := coin.Quote["USD"]
		if !ok {
			log.Warnf("Couldn't find usd price for coin: %v", coin.Symbol)
			continue
		}
		foreignQuotation := models.ForeignQuotation{
			Symbol:             coin.Symbol,
			Name:               coin.Name,
			Price:              usdQuote.Price,
			PriceYesterday:     priceYesterday,
			VolumeYesterdayUSD: usdQuote.Volume24H,
			Source:             coinMarketCapsource,
			Time:               timestamp,
			ITIN:               itin.Itin,
		}
		scraper.foreignScrapper.chanQuotation <- &foreignQuotation
	}
	return nil

}

func (scraper *CoinMarketCapScraper) GetQuoteChannel() chan *models.ForeignQuotation {
	return scraper.foreignScrapper.chanQuotation
}

func getCoinMarketCapData() (listing CoinMarketCapListing, err error) {
	// There must be a pro coinmarketcap api key for this to work properly
	var apiKey string
	var ok bool
	if apiKey, ok = os.LookupEnv("CMC_API_KEY"); !ok {
		err = fmt.Errorf("please provide a CoinMarketCap api key env (CMC_API_KEY)")
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Print(err)
		return
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "200")
	q.Add("convert", "USD")
	q.Add("sort", "market_cap")
	q.Add("sort_dir", "desc")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	response, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(response, &listing)
	if err != nil {
		return
	}
	return
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *CoinMarketCapScraper) cleanup(err error) {

	scraper.foreignScrapper.errorLock.Lock()
	defer scraper.foreignScrapper.errorLock.Unlock()

	scraper.foreignScrapper.tickerRate.Stop()
	scraper.foreignScrapper.tickerState.Stop()

	if err != nil {
		scraper.foreignScrapper.error = err
	}
	scraper.foreignScrapper.closed = true

	close(scraper.foreignScrapper.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections
func (scraper *CoinMarketCapScraper) Close() error {
	if scraper.foreignScrapper.closed {
		return errors.New("Scraper: Already closed")
	}
	close(scraper.foreignScrapper.shutdown)
	<-scraper.foreignScrapper.shutdownDone
	scraper.foreignScrapper.errorLock.RLock()
	defer scraper.foreignScrapper.errorLock.RUnlock()
	return scraper.foreignScrapper.error
}
