package main

import (
	"flag"

	scrapers "github.com/diadata-org/diadata/internal/pkg/foreign-scrapers"
	log "github.com/sirupsen/logrus"
)

func main() {

	scraperType := flag.String("type", "", "which foreignQuotation")
	flag.Parse()

	switch *scraperType {
	case "Coingecko":
		log.Println("Foreign Scraper: Start scraping data from Coingecko")
		scrapers.NewCoingeckoScraper()
	}
}
