package main

import (
	"flag"
	"sync"

	ratescrapers "github.com/diadata-org/diadata/internal/pkg/ratescrapers"
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

	// Parse the option for the type of interest rate
	rateType := flag.String("type", "ESTER", "Type of interest rate")
	flag.Parse()

	wg := sync.WaitGroup{}
	ds, err := models.NewDataStore()

	// -------------------------------------------------------------------------
	// Prefill the database with historic data when the scrapers are initialized
	// -------------------------------------------------------------------------

	ratescrapers.WriteHistoricSOFR(ds)
	ratescrapers.WriteHistoricSOFRAvg(ds)

	// -------------------------------------------------------------------------

	if err != nil {
		log.Errorln("NewDataStore:", err)
	} else {

		// Spawn the corresponding rate scraper
		sRate := ratescrapers.SpawnRateScraper(ds, *rateType)
		defer sRate.Close()

		// Send rates to the database while the scraper scrapes
		wg.Add(1)
		go handleInterestRate(sRate.Channel(), &wg, ds)
		defer wg.Wait()
	}
}
