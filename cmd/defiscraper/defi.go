package main

import (
	"flag"
	"sync"

	"github.com/diadata-org/diadata/pkg/dia"

	defiscraper "github.com/diadata-org/diadata/internal/pkg/defiscrapers"
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
		ds.SetDefiRateInflux(t)
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
		ds.SetDefiStateInflux(t)
	}
}

func main() {
	rateType := flag.String("type", "DYDX", "Type of Defi rate")
	rpcURL := flag.String("rpc-url", "http://127.0.0.1:8545", "RPC endpoint of the Ethereum client")
	yearnAprOracleAddress := flag.String("apr-oracle", "0x97ff4a1b787ade6b94cca95b61f79417c673331d", "Address of the deployed APR Oracle address")

	flag.Parse()

	wg := sync.WaitGroup{}
	ds, err := models.NewDataStore()

	if err != nil {
		log.Errorln("NewDataStore:", err)
	} else {
		// Init Yearn Manager
		sRate := defiscraper.SpawnDefiScraper(ds, *rateType, *rpcURL, *yearnAprOracleAddress)
		defer sRate.Close()

		// Send rates to the database while the scraper scrapes
		wg.Add(2)
		go handleDefiInterestRate(sRate.RateChannel(), &wg, ds)
		go handleDefiState(sRate.StateChannel(), &wg, ds)

		defer wg.Wait()
	}
}
