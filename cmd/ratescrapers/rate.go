package main

import (
	"flag"
	"sync"

	ratescrapers "github.com/diadata-org/diadata/internal/pkg/ratescrapers"
	staticscrapers "github.com/diadata-org/diadata/internal/pkg/static-scrapers"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// handleInterestRate delegates rate information to Kafka
func handleInterestRate(c chan *models.InterestRate, wg *sync.WaitGroup, ds models.Datastore) {
	defer wg.Done()
	// Pull from channel as long as not empty
	for {
		t, ok := <-c
		if !ok {
			log.Error("error")
			return
		}
		ds.SetInterestRate(t)
	}
}

// main manages all Scraper and handles incoming trade information
func main() {

	// Parse the option for the type of interest rate. The available values
	// for the flags can be found in the Update() method in RateScraper.go.
	rateType := flag.String("type", "ESTER", "Type of interest rate")
	flag.Parse()

	wg := sync.WaitGroup{}
	ds, err := models.NewDataStore()

	if err != nil {
		log.Errorln("NewDataStore:", err)
	} else {

		// Download historic data (in case there is)
		err = staticscrapers.LoadHistoricRate(*rateType)
		if err != nil {
			log.Errorf("Error downloading resources for rate %s: %v", *rateType, err)
		}

		// Writing historic data into database
		err = staticscrapers.WriteHistoricRate(ds, *rateType)
		if err != nil {
			log.Errorf("Error writing rate %s: %v", *rateType, err)
		}

		// Spawn the corresponding rate scraper
		sRate := ratescrapers.SpawnRateScraper(ds, *rateType)
		defer sRate.Close()

		// Send rates to the database while the scraper scrapes
		wg.Add(1)
		go handleInterestRate(sRate.Channel(), &wg, ds)
		defer wg.Wait()
	}
}
