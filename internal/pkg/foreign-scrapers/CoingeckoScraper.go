package foreignscrapers

// Please implement the scraping of coingecko quotations here.

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

const (
	coingeckoRefreshDelay = time.Second * 10 * 1
	source                = "Coingecko"
)

type CoingeckoCoin struct {
	ID                 string  `json:"id"`
	Symbol             string  `json:"symbol"`
	Name               string  `json:"name"`
	CurrentPrice       float64 `json:"current_price"`
	Yesterday24hVolume float64 `json:"total_volume"`
	LastUpdated        string  `json:"last_updated"`
}

type nothing struct{}

type CoingeckoScraper struct {
	ticker   *time.Ticker
	foreignScrapper ForeignScraper
}

func NewCoingeckoScraper(datastore models.Datastore)*CoingeckoScraper {

	foreignScrapper := ForeignScraper{
		shutdown: make(chan nothing),
		error:    nil,
		datastore: datastore,
		chanQuotation: make(chan *models.ForeignQuotation),

	}
	s := &CoingeckoScraper{
		ticker:   time.NewTicker(coingeckoRefreshDelay),
		foreignScrapper:foreignScrapper,
 	}
	  go s.mainLoop()

	return s

}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *CoingeckoScraper) mainLoop() {
	for true {
 		select {
		case <-scraper.ticker.C:
			 scraper.UpdateQuotation()
		case <-scraper.foreignScrapper.shutdown: // user requested shutdown
			log.Printf("Coingeckoscraper shutting down")
			return
		}
	}

}



// Update retrieves new coin information from the coingecko API and stores it to influx
func (scraper *CoingeckoScraper) UpdateQuotation() error {
	log.Printf("Executing CoingeckoScraper update")

	coins, err := getCoingeckoData()

	if err != nil {
		log.Errorln("Error getting data from coingecko",err)
		return err
	}

	for _, coin := range coins {

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

		foreignQuotation := models.ForeignQuotation{
			Symbol:             coin.Symbol,
			Name:               coin.Name,
			Price:              coin.CurrentPrice,
			PriceYesterday:     priceYesterday,
			VolumeYesterdayUSD: coin.Yesterday24hVolume,
			Source:             source,
			Time:               timestamp,
			ITIN:               itin.Itin,
		}
		scraper.foreignScrapper.chanQuotation <- &foreignQuotation
	}
	return nil

}

func (scraper *CoingeckoScraper) GetQuoteChannel() chan *models.ForeignQuotation {
	return scraper.foreignScrapper.chanQuotation
}

func getCoingeckoData() (coins []CoingeckoCoin, err error) {
	response, err := utils.GetRequest("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false")
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &coins)
	if err != nil {
		return
	}
	return
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *CoingeckoScraper) cleanup(err error) {

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
func (scraper *CoingeckoScraper) Close() error {
	if scraper.foreignScrapper.closed {
		return errors.New("Scraper: Already closed")
	}
	close(scraper.foreignScrapper.shutdown)
	<-scraper.foreignScrapper.shutdownDone
	scraper.foreignScrapper.errorLock.RLock()
	defer scraper.foreignScrapper.errorLock.RUnlock()
	return scraper.foreignScrapper.error
}
