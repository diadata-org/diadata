package main

import (
	"flag"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"

	scrapers "github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/block-scrapers"
	log "github.com/sirupsen/logrus"
)

func main() {

	wg := sync.WaitGroup{}

	ds, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}

	scraperType := flag.String("blockchain", "", "which blockchain")
	flag.Parse()
	var blockscraper scrapers.BlockScraperInterface

	switch *scraperType {
	case "Ethereum":
		log.Println("Block-scraper: Start scraping block data from Ethereum")
		blockscraper = scrapers.NewEthereumScraper(ds)
	default:
		for {
			time.Sleep(24 * time.Hour)
		}
	}

	wg.Add(1)
	go handleBlockData(blockscraper.GetDataChannel(), &wg, ds)
	defer wg.Wait()

}

func handleBlockData(blockdatachannel chan dia.BlockData, wg *sync.WaitGroup, ds models.RelDatastore) {
	defer wg.Done()

	for {
		blockdata, ok := <-blockdatachannel
		if !ok {
			log.Error("error")
			return
		}
		log.Infof("got block number %s: %v", blockdata.Number, blockdata.Data)
		err := ds.SetBlockData(blockdata)
		if err != nil {
			log.Error("setting block data: ", err)
			return
		}
	}

}
