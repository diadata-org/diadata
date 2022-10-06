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

const watchdogDelay = 3600

func main() {

	wg := sync.WaitGroup{}

	db, err := models.NewDataStore()
	if err != nil {
		log.Fatal(" datastore error: ", err)
	}

	scraperType := flag.String("synthAsset", "cETH", "which synthetic asset")
	flag.Parse()
	var scraper synthscrapers.SynthScraperInterface

	switch *scraperType {
	case "cETH":
		log.Println("Start scraping data from cETH")
		scraper = synthscrapers.NewcETHScraper()
	case "atokenv2ethereum":
		log.Print("Start scraping data from aToken ")
		scraper = synthscrapers.NewaTokenScraper(dia.ETHEREUM, "0x7d2768de32b0b80b7a3454c06bdac94a69ddc7a9", 2)
	case "atokenv3polygon":
		log.Print("Start scraping data from aToken ")
		scraper = synthscrapers.NewaTokenScraper(dia.POLYGON, "0x794a61358D6845594F94dc1DB02A252b5b4814aD", 3)
	case "atokenv3avalanche":
		log.Print("Start scraping data from aToken ")
		scraper = synthscrapers.NewaTokenScraper(dia.AVALANCHE, "0x794a61358d6845594f94dc1db02a252b5b4814ad", 3)
	case "atokenv2avalanche":
		log.Print("Start scraping data from aToken ")
		scraper = synthscrapers.NewaTokenScraper(dia.AVALANCHE, "0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c", 2)

	default:
		for {
			time.Sleep(24 * time.Hour)
		}
	}

	wg.Add(1)

	go handleSynthData(scraper.GetSynthSupplyChannel(), &wg, db)
	defer wg.Wait()

}

func handleSynthData(synthChannel chan dia.SynthAssetSupply, wg *sync.WaitGroup, db *models.DB) {
	defer wg.Done()
	lastTradeTime := time.Now()
	t := time.NewTicker(time.Duration(watchdogDelay) * time.Second)
	for {
		select {
		case <-t.C:
			duration := time.Since(lastTradeTime)
			if duration > time.Duration(watchdogDelay)*time.Second {
				log.Error(duration)
				panic("frozen? ")
			}
		case synthData, ok := <-synthChannel:
			if !ok {
				log.Error("error")
				return
			}
			log.Infoln("synthData", synthData)
			err := db.SaveSynthSupplyInflux(&synthData)
			if err != nil {
				log.Errorf("Error saving synth data for %s: %v", synthData.Asset.Address, err)
			} else {
				log.Infof("successfully set synth data for %s", synthData.Asset.Symbol)
			}
			err = db.Flush()
			if err != nil {
				log.Error(err)
			}
		}
	}
}
