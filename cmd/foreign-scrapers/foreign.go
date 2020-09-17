package main

import (
	"flag"
	"fmt"

	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)


var foreignQuotation  = flag.String("foreign", "", "which foreignQuotation")


func main() {
	// TO DO
	ds, err := models.NewInfluxDataStore()
	if err != nil {
		log.Errorln("NewInfluxDataStore:", err)
	} else {
		if *foreignQuotation == "Coingecko" {
			log.Println("Foreign Scraper: Start scrapping data from Coingecko")
			s := scrapers.foreign-scrapers.NewCoingeckoScraper(ds)
			go s.Update()
		}
	}
	log.Println("Done Foreign scrapping")	
}
