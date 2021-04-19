package main

import (
	"flag"
	"sync"

	pool "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

// Fetch Pools
func main() {
	poolName := flag.String("type", "CVAULT", "Name of Pool")
	flag.Parse()

	hashWriter, err := kafkaHelper.NewHashWriter(kafkaHelper.HASH_FARMINGPOOLS, true)
	if err != nil {
		log.Fatal(err)
	}
	defer hashWriter.Close()

	wg := sync.WaitGroup{}

	ds, err := models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore:", err)
	} else {
		sRate := pool.SpawnPoolScraper(ds, *poolName)
		defer sRate.Close()

		// Send rates to the database while the scraper scrapes
		wg.Add(2)
		go handlerate(sRate.RateChannel(), &wg, ds, hashWriter)

		defer wg.Wait()
	}

}
func handlerate(c chan *models.FarmingPool, wg *sync.WaitGroup, ds models.Datastore, hashWriter *kafka.Writer) {
	defer wg.Done()
	// Pull from channel as long as not empty
	for {
		pr, ok := <-c
		if !ok {
			log.Error("error")
			return
		}
		log.Print("Write pool info: ", pr)
		ds.SetFarmingPool(pr, hashWriter)
	}
}
