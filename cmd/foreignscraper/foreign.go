package main

import (
	"flag"
	"sync"

	models "github.com/diadata-org/diadata/pkg/model"

	scrapers "github.com/diadata-org/diadata/internal/pkg/foreign-scrapers"
	log "github.com/sirupsen/logrus"
)

func main() {

	wg := sync.WaitGroup{}

	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}

	scraperType := flag.String("foreignType", "", "which foreignQuotation")
	flag.Parse()
	var sc scrapers.ForeignScrapperer

	switch *scraperType {
	case "Coingecko":
		log.Println("Foreign Scraper: Start scraping data from Coingecko")
		sc = scrapers.NewCoingeckoScraper(ds)
	case "CoinMarketCap":
		log.Println("Foreign Scraper: Start scraping data from CoinMarketCap")
		sc = scrapers.NewCoinMarketCapScraper(ds)
	}

	wg.Add(1)
	go handleQuotation(sc.GetQuoteChannel(), &wg, ds)
	defer wg.Wait()

}

func handleQuotation(quotation chan *models.ForeignQuotation, wg *sync.WaitGroup, ds models.Datastore) {
	defer wg.Done()

	for {
		fq, ok := <-quotation
		if !ok {
			log.Error("error")
			return
		}

		ds.SaveForeignQuotationInflux(*fq)
	}

}
