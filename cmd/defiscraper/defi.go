package main

import (
	"flag"
	"sync"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	"github.com/segmentio/kafka-go"

	defiscraper "github.com/diadata-org/diadata/internal/pkg/defiscrapers"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// handleDefiInterestRate delegates rate information to Kafka
func handleDefiInterestRate(c chan *dia.DefiRate, wg *sync.WaitGroup, ds models.Datastore, hashWriter *kafka.Writer) {
	defer wg.Done()
	// Pull from channel as long as not empty
	for {
		t, ok := <-c
		if !ok {
			log.Error("error")
			return
		}
		ds.SetDefiRateInflux(t, hashWriter)
	}
}

// handleDefiState delegates rate information to Kafka
func handleDefiState(c chan *dia.DefiProtocolState, wg *sync.WaitGroup, ds models.Datastore, hashWriter *kafka.Writer) {
	defer wg.Done()
	// Pull from channel as long as not empty
	for {
		t, ok := <-c
		if !ok {
			log.Error("error")
			return
		}
		ds.SetDefiStateInflux(t, hashWriter)
	}
}

func main() {
	rateType := flag.String("type", "DYDX", "Type of Defi rate")
	flag.Parse()

	wg := sync.WaitGroup{}
	hashWriterRates, err := kafkaHelper.NewHashWriter("hash-lendingrates", true)
	if err != nil {
		log.Fatal(err)
	}
	hashWriterStates, err := kafkaHelper.NewHashWriter("hash-lendingstates", true)
	if err != nil {
		log.Fatal(err)
	}

	ds, err := models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore:", err)
	} else {

		sRate := defiscraper.SpawnDefiScraper(ds, *rateType)
		defer sRate.Close()

		// Send rates to the database while the scraper scrapes
		wg.Add(2)
		go handleDefiInterestRate(sRate.RateChannel(), &wg, ds, hashWriterRates)
		go handleDefiState(sRate.StateChannel(), &wg, ds, hashWriterStates)

		defer wg.Wait()
	}
}
