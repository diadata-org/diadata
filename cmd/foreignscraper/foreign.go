package main

import (
	"flag"

	scrapers "github.com/diadata-org/diadata/internal/pkg/foreign-scrapers"
	log "github.com/sirupsen/logrus"
)

// TO DO: write handlequotation method, handling the foreignQuotation channel from
// Coingecko Scraper. Save to influx.

func main() {

	scraperType := flag.String("foreignType", "", "which foreignQuotation")
	flag.Parse()

	switch *scraperType {
	case "Coingecko":
		log.Println("Foreign Scraper: Start scraping data from Coingecko")
		scrapers.NewCoingeckoScraper()
	}
}
