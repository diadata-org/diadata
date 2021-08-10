package main

import (
	"flag"
	"github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper"
	"sync"

	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// Fetch Pools
func main() {
	poolName := flag.String("type", "CVAULT", "Name of Pool")
	flag.Parse()
	ds, err := models.NewDataStore()
	wg := sync.WaitGroup{}

	if err != nil {
		log.Errorln("NewDataStore:", err)
	} else {
		sRate := pool.SpawnPoolScraper(ds, *poolName)
		defer func() {
			err := sRate.Close()
			if err != nil {
				log.Error(err)
			}
		}()

		// Send rates to the database while the scraper scrapes
		wg.Add(2)
		go handlerate(sRate.RateChannel(), &wg, ds)

		defer wg.Wait()
	}

}
func handlerate(c chan *models.FarmingPool, wg *sync.WaitGroup, ds models.Datastore) {
	defer wg.Done()
	// Pull from channel as long as not empty
	for {
		pr, ok := <-c
		if !ok {
			log.Error("error")
			return
		}
		log.Print("Write pool info: ", pr)
		err := ds.SetFarmingPool(pr)
		if err != nil {
			log.Error(err)
		}
	}
}
