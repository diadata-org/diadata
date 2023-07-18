package main

import (
	"flag"

	liquidityscraper "github.com/diadata-org/diadata/pkg/dia/scraper/liquidity-scrapers"
	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/sirupsen/logrus"
)

var (
	exchangeName *string
	log          *logrus.Logger
)

func init() {
	exchangeName = flag.String("exchange", "Uniswap", "name of DEX.")
	flag.Parse()
	log = logrus.New()
}

func main() {

	log.Println("Liquidity Scraper: Start scraping liquidity")

	relDB, err := models.NewRelDataStore()
	if err != nil {
		log.Errorln("Error connecting to postgres: ", err)
		return
	}

	datastore, err := models.NewDataStore()
	if err != nil {
		log.Errorln("Error connecting to postgres: ", err)
		return
	}

	runLiquiditySource(relDB, datastore, *exchangeName)
	log.Infof("Successfully ran pool collector for %s", *exchangeName)

}

func runLiquiditySource(relDB *models.RelDB, datastore *models.DB, source string) {
	log.Info("Fetching pools from ", source)
	scraper := liquidityscraper.NewLiquidityScraper(source, relDB, datastore)

	for {
		select {
		case receivedPool := <-scraper.Pool():

			// Set to persistent DB.
			err := relDB.SetPool(receivedPool)
			if err != nil {
				log.Errorf("Error saving pool %v: %v", receivedPool, err)
			} else {
				log.Info("successfully set pool ", receivedPool)
			}

		case <-scraper.Done():
			return
		}
	}

}
