package main

import (
	"flag"
	"sync"

	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/segmentio/kafka-go"

	scrapers "github.com/diadata-org/diadata/internal/pkg/foreign-scrapers"
	log "github.com/sirupsen/logrus"
)

func main() {

	wg := sync.WaitGroup{}

	hashWriter, err := kafkaHelper.NewHashWriter("hash-foreignscraper", true)
	if err != nil {
		log.Fatal(err)
	}
	defer hashWriter.Close()

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
	default:
		sc = scrapers.NewGenericForeignScraper()
	}

	wg.Add(1)
	go handleQuotation(sc.GetQuoteChannel(), &wg, ds, hashWriter)
	defer wg.Wait()

}

func handleQuotation(quotation chan *models.ForeignQuotation, wg *sync.WaitGroup, ds models.Datastore, hashWriter *kafka.Writer) {
	defer wg.Done()

	for {
		fq, ok := <-quotation
		if !ok {
			log.Error("error")
			return
		}

		ds.SaveForeignQuotationInflux(*fq, hashWriter)
	}

}
