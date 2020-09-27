package foreignscrapers

// Please implement the scraping of coingecko quotations here.

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

const (
	coingeckoRefreshDelay = time.Second * 60 * 1
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
	error    error
	shutdown chan nothing
	ticker   *time.Ticker
}

func NewCoingeckoScraper() {
	s := &CoingeckoScraper{
		ticker:   time.NewTicker(coingeckoRefreshDelay),
		shutdown: make(chan nothing),
		error:    nil,
	}
	s.mainLoop()

}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *CoingeckoScraper) mainLoop() {
	for {
		select {
		case <-scraper.ticker.C:
			scraper.Update()
		case <-scraper.shutdown: // user requested shutdown
			log.Printf("Coingeckoscraper shutting down")
			return
		}
	}

}

// Update retrieves new coin information from the coingecko API and stores it to influx
func (scraper *CoingeckoScraper) Update() error {
	log.Printf("Executing CoingeckoScraper update")
	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}
	coins, err := getCoingeckoData()
	if err != nil {
		return err
	}

	for _, coin := range coins {
		// TO DO: normalize symbol instead of capitalizing
		coin.Symbol = strings.ToUpper(coin.Symbol)

		// Parse last updated timestamp
		layout := "2006-01-02T15:04:05.000Z"
		timestamp, err := time.Parse(layout, coin.LastUpdated)
		if err != nil {
			log.Error("error parsing time")
		}

		// get ITIN if available in redis
		itin, err := ds.GetItinBySymbol(coin.Symbol)
		if err != nil {
			// log.Errorf("error: no ITIN available for %s \n", coin.Symbol)
			itin = dia.ItinToken{}
		}

		// Get yesterday's price from influx if available
		priceYesterday, err := ds.GetForeignPriceYesterday(coin.Symbol, source)
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
		ds.SaveForeignQuotationInflux(foreignQuotation)

	}
	// Flush influx batch at the end of update
	err = ds.Flush()
	if err != nil {
		log.Errorln("SaveForeignQuotationInflux", err)
	}
	time.Sleep(time.Second * 2)

	return nil

}

func getCoingeckoData() (coins []CoingeckoCoin, err error) {
	response, err := utils.GetRequest("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=200&page=1&sparkline=false")
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &coins)
	if err != nil {
		return
	}
	return
}
