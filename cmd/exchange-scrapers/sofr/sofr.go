package main

import (
	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// handleInterestRate delegates rate information to Kafka
func handleInterestRate(c chan *dia.InterestRate, ds models.Datastore) {

	// Pull from channel as long as not empty
	for {
		t, ok := <-c
		if !ok {
			log.Error("error")
			return
		}
		log.Println(t.Symbol)
		ds.SetInterestRate(*t)
	}
}

// main manages all Scraper and handles incoming trade information
func main() {

	ds, err := models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore:", err)
	} else {

		// sSOFR points to a SOFRScraper struct
		sSOFR := scrapers.NewSOFRScraper(ds)
		defer sSOFR.Close()

		sSOFR.Update()
		handleInterestRate(sSOFR.Channel(), ds)

	}
}
