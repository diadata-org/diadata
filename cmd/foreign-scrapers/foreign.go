package main

import (
	"flag"

	scrapers "github.com/diadata-org/diadata/internal/pkg/foreign-scrapers"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

const Coingecko = "Coingecko"


func main() {
	// TO DO
	
	foreignQuotation := flag.String("foreign", "", "which foreignQuotation")
	flag.Parse()

	ds, err := models.NewInfluxDataStore()
	if err != nil {
		log.Errorln("NewInfluxDataStore:", err)
	} else {
		if *foreignQuotation == Coingecko {
			log.Println("Foreign Scraper: Start scrapping data from Coingecko")
			scrapers.NewCoingeckoScraper(ds)
		}
		select {} 
	}
	log.Println("Foreign scrapping failed to start")
}
