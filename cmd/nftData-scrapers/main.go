package main

import (
	"flag"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"

	nftdatascrapers "github.com/diadata-org/diadata/internal/pkg/nftData-scrapers"
	log "github.com/sirupsen/logrus"
)

func main() {

	wg := sync.WaitGroup{}

	rdb, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("relational datastore error: ", err)
	}

	scraperType := flag.String("nftclass", "Sorare", "which NFT class")
	flag.Parse()
	var scraper nftdatascrapers.NFTDataScraper

	switch *scraperType {
	case "Sorare":
		log.Println("NFT Data Scraper: Start scraping data from Sorare")
		scraper = nftdatascrapers.NewSorareScraper(rdb)
	default:
		for {
			time.Sleep(24 * time.Hour)
		}
	}

	wg.Add(1)
	go handleData(scraper.GetDataChannel(), &wg, rdb)
	defer wg.Wait()

}

func handleData(dataChannel chan *dia.NFT, wg *sync.WaitGroup, rdb *models.RelDB) {
	defer wg.Done()

	for {
		fq, ok := <-dataChannel
		if !ok {
			log.Error("error")
			return
		}
		rdb.SetNFTData(*fq)
	}

}
