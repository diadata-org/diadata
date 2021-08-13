package main

import (
	"flag"
	"sync"
	"time"

	stockscrapers "github.com/diadata-org/diadata/internal/pkg/stock-scrapers"
	models "github.com/diadata-org/diadata/pkg/model"

	log "github.com/sirupsen/logrus"
)

func main() {

	wg := sync.WaitGroup{}

	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}

	scraperType := flag.String("source", "Fineage", "which source for stock quotes")
	flag.Parse()
	var scraper stockscrapers.StockScraperInterface

	switch *scraperType {
	case "Fineage":
		log.Println("Stock Quote Scraper: Start scraping trades from Fineage")
		scraper = stockscrapers.NewFineageScraper(ds)
	default:
		for {
			time.Sleep(24 * time.Hour)
		}
	}

	wg.Add(1)
	go handleData(scraper.GetStockQuotationChannel(), &wg, ds)
	defer wg.Wait()

}

func handleData(stockQuotationChannel chan models.StockQuotation, wg *sync.WaitGroup, ds *models.DB) {
	defer wg.Done()

	for {
		quote, ok := <-stockQuotationChannel
		if !ok {
			log.Error("error")
			return
		}
		err := ds.SetStockQuotation(quote)
		if err != nil {
			log.Error("setting stock quotation: ", err)
			return
		} else {
			log.Infof("successfully set stock quotation for %s", quote.Symbol)
		}
	}

}
