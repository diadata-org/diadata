package foreignscrapers

// Please implement the scraping of coingecko quotations here.

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
	source = "Coingecko"
)

var (
	coingeckoUpdateSeconds int64
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
	ticker          *time.Ticker
	foreignScrapper ForeignScraper
	apiKey          string
	apiSecret       string
}

func init() {
	var err error
	coingeckoUpdateSeconds, err = strconv.ParseInt(utils.Getenv("COINGECKO_UPDATE_SECONDS", "14400"), 10, 64)
	if err != nil {
		log.Error("Parse COINGECKO_UPDATE_SECONDS: ", err)
	}
}

func NewCoingeckoScraper(datastore models.Datastore, apiKey string, apiSecret string) *CoingeckoScraper {

	foreignScrapper := ForeignScraper{
		shutdown:      make(chan nothing),
		error:         nil,
		datastore:     datastore,
		chanQuotation: make(chan *models.ForeignQuotation),
	}
	s := &CoingeckoScraper{
		ticker:          time.NewTicker(time.Duration(coingeckoUpdateSeconds) * time.Second),
		foreignScrapper: foreignScrapper,
		apiKey:          apiKey,
		apiSecret:       apiSecret,
	}
	go s.mainLoop()

	return s

}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *CoingeckoScraper) mainLoop() {

	// Initial run.
	err := scraper.UpdateQuotation()
	if err != nil {
		log.Error(err)
	}

	// Update every @coingeckoUpdateSeconds.
	for {
		select {
		case <-scraper.ticker.C:
			err := scraper.UpdateQuotation()
			if err != nil {
				log.Error(err)
			}
		case <-scraper.foreignScrapper.shutdown: // user requested shutdown
			log.Printf("Coingeckoscraper shutting down")
			return
		}
	}

}

// Update retrieves new coin information from the coingecko API and stores it to influx
func (scraper *CoingeckoScraper) UpdateQuotation() error {
	log.Printf("Executing CoingeckoScraper update")

	coins, err := scraper.getCoingeckoData()

	if err != nil {
		log.Errorln("Error getting data from coingecko", err)
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
		}
		scraper.foreignScrapper.chanQuotation <- &foreignQuotation
	}
	return nil

}

func (scraper *CoingeckoScraper) GetQuoteChannel() chan *models.ForeignQuotation {
	return scraper.foreignScrapper.chanQuotation
}

func (scraper *CoingeckoScraper) getCoingeckoData() (coins []CoingeckoCoin, err error) {
	var response []byte
	if scraper.apiKey != "" {
		baseString := "https://pro-api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1&sparkline=false&x_cg_pro_api_key=" + scraper.apiKey
		response, _, err = utils.GetRequest(baseString)

	} else {
		response, _, err = utils.GetRequest("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1&sparkline=false")
	}
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &coins)
	if err != nil {
		return
	}
	return
}

// Close closes any existing API connections
func (scraper *CoingeckoScraper) Close() error {
	if scraper.foreignScrapper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.foreignScrapper.shutdown)
	<-scraper.foreignScrapper.shutdownDone
	scraper.foreignScrapper.errorLock.RLock()
	defer scraper.foreignScrapper.errorLock.RUnlock()
	return scraper.foreignScrapper.error
}
