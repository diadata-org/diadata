package main

import (
	"flag"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"

	synthscrapers "github.com/diadata-org/diadata/pkg/dia/scraper/synthetic-scrapers"
	log "github.com/sirupsen/logrus"
)

func main() {

	wg := sync.WaitGroup{}

	rdb, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("relational datastore error: ", err)
	}

	scraperType := flag.String("synthAsset", "cETH", "which synthetic asset")
	flag.Parse()
	var scraper synthscrapers.SynthScraperInterface

	switch *scraperType {
	case "cETH":
		log.Println("Start scraping data from cETH")
		scraper = synthscrapers.NewcETHScraper(rdb)
	default:
		for {
			time.Sleep(24 * time.Hour)
		}
	}

	wg.Add(1)
	go handleSynthData(scraper.GetSynthSupplyChannel(), &wg, rdb)
	defer wg.Wait()

}

func handleSynthData(synthChannel chan dia.SynthAssetSupply, wg *sync.WaitGroup, rdb *models.RelDB) {
	defer wg.Done()
	for {
		synthData, ok := <-synthChannel
		if !ok {
			log.Error("error")
			return
		}
		err := rdb.SetSynthAssetSupply(synthData)
		if err != nil {
			log.Errorf("Error saving synth data for %s: %v", synthData.Asset.Address, err)
		} else {
			log.Infof("successfully set synth data for %s", synthData.Asset.Symbol)
		}
	}
}
