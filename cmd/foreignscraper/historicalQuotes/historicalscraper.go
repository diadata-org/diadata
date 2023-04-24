package main

import (
	"flag"

	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/historical-scrapers"

	models "github.com/diadata-org/diadata/pkg/model"

	log "github.com/sirupsen/logrus"
)

var (
	source *string
)

func init() {
	source = flag.String("source", "", "which source for historical quotations")
	flag.Parse()
}

func main() {

	rdb, err := models.NewPostgresDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}

	datastore, err := models.NewDataStoreWithoutRedis()
	if err != nil {
		log.Fatalf("Failed to create datastore: %v", err)
	}

	runHistoricalSource(*source, rdb, datastore)

}

func runHistoricalSource(source string, rdb *models.RelDB, datastore *models.DB) {
	scraper := scrapers.NewHistoricalScraper(source, rdb, datastore)

	for {
		select {
		case hq, ok := <-scraper.QuoteChannel():
			if !ok {
				log.Error("error")
				return
			}
			log.Info("historical Quotation: ", hq)
			err := rdb.SetHistoricalQuotation(hq)
			if err != nil {
				log.Error(err)
			}
		case <-scraper.Done():
			log.Info("successfully updated historical data")
			return
		}

	}
}
