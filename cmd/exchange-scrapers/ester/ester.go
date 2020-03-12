package main

import (
	"sync"

	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
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

	wg := sync.WaitGroup{}
	ds, err := models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore:", err)
	} else {

		// sESTER points to a ESTERScraper struct
		sESTER := scrapers.SpawnESTERScraper(ds)
		defer sESTER.Close()

		wg.Add(1)
		go handleInterestRate(sESTER.Channel(), &wg, ds)
		defer wg.Wait()
	}
}
