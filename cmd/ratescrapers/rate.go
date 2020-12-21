package main

import (
	"flag"
	"github.com/diadata-org/diadata/internal/pkg/ratescrapers/yearn"
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
	rateType := flag.String("type", "SOFR", "Type of interest rate")
	rpcURL := flag.String("rpc-url", "http://127.0.0.1:8545", "RPC endpoint of the Ethereum client")
	yearnAprOracleAddress := flag.String("apr-oracle", "0x97ff4a1b787ade6b94cca95b61f79417c673331d", "Address of the deployed APR Oracle address")

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

		// Init Yearn Manager
		yearnManager := yearn.NewYearnManager(*rpcURL, *yearnAprOracleAddress)
		// Spawn the corresponding rate scraper
		sRate := ratescrapers.SpawnRateScraper(ds, *rateType, yearnManager)
		defer sRate.Close()

		// Send rates to the database while the scraper scrapes
		wg.Add(1)
		go handleInterestRate(sRate.Channel(), &wg, ds)
		defer wg.Wait()
	}
}
