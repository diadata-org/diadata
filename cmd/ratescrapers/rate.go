package main

import (
	"flag"
	"sync"

	ratescrapers "github.com/diadata-org/diadata/internal/pkg/ratescrapers"
	staticscrapers "github.com/diadata-org/diadata/internal/pkg/static-scrapers"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

// handleInterestRate delegates rate information to Kafka
func handleInterestRate(c chan *models.InterestRate, wg *sync.WaitGroup, ds models.Datastore, hashWriter *kafka.Writer) {
	defer wg.Done()
	// Pull from channel as long as not empty
	for {
		t, ok := <-c
		if !ok {
			log.Error("error")
			return
		}
		ds.SetInterestRate(t, hashWriter)
	}
}

// main manages all Scraper and handles incoming trade information
func main() {

	// Parse the option for the type of interest rate. The available values
	// for the flags can be found in the Update() method in RateScraper.go.
	rateType := flag.String("type", "SOFR", "Type of interest rate")
	flag.Parse()

	hashWriter, err := kafkaHelper.NewHashWriter(kafkaHelper.HASH_INTERESTRATES, true)
	if err != nil {
		log.Fatal(err)
	}
	defer hashWriter.Close()

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
		err = staticscrapers.WriteHistoricRate(ds, *rateType, hashWriter)
		if err != nil {
			log.Errorf("Error writing rate %s: %v", *rateType, err)
		}

		// Spawn the corresponding rate scraper
		sRate := ratescrapers.SpawnRateScraper(ds, *rateType)
		defer sRate.Close()

		// Send rates to the database while the scraper scrapes
		wg.Add(1)
		go handleInterestRate(sRate.Channel(), &wg, ds, hashWriter)
		defer wg.Wait()
	}
}
