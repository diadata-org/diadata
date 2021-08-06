package main

import (
	"flag"
	"github.com/diadata-org/diadata/pkg/dia/defiscrapers"
	"sync"

	"github.com/diadata-org/diadata/pkg/dia"

	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// handleDefiInterestRate delegates rate information to Kafka
func handleDefiInterestRate(c chan *dia.DefiRate, wg *sync.WaitGroup, ds models.Datastore) {
	defer wg.Done()
	// Pull from channel as long as not empty
	for {
		t, ok := <-c
		if !ok {
			log.Error("error")
			return
		}
		err := ds.SetDefiRateInflux(t)
		if err != nil {
			log.Error(err)
		}
	}
}

// handleDefiState delegates rate information to Kafka
func handleDefiState(c chan *dia.DefiProtocolState, wg *sync.WaitGroup, ds models.Datastore) {
	defer wg.Done()
	// Pull from channel as long as not empty
	for {
		t, ok := <-c
		if !ok {
			log.Error("error")
			return
		}
		err := ds.SetDefiStateInflux(t)
		if err != nil {
			log.Error(err)
		}
	}
}

func main() {
	rateType := flag.String("type", "DYDX", "Type of Defi rate")
	flag.Parse()

	wg := sync.WaitGroup{}
	ds, err := models.NewDataStore()

	if err != nil {
		log.Errorln("NewDataStore:", err)
	} else {

		sRate := defiscrapers.SpawnDefiScraper(ds, *rateType)
		defer func() {
			err := sRate.Close()
			if err != nil {
				log.Error(err)
			}
		}()

		// Send rates to the database while the scraper scrapes
		wg.Add(2)
		go handleDefiInterestRate(sRate.RateChannel(), &wg, ds)
		go handleDefiState(sRate.StateChannel(), &wg, ds)

		defer wg.Wait()
	}
}
