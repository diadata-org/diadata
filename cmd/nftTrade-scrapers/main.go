package main

import (
	"flag"
	nftdatascrapers2 "github.com/diadata-org/diadata/internal/nftData-scrapers"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"

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
	var scraper nftdatascrapers2.NFTDataScraper

	switch *scraperType {
	case "Sorare":
		log.Println("NFT Data Scraper: Start scraping data from Sorare")
		scraper = nftdatascrapers2.NewSorareScraper(rdb)
	default:
		for {
			time.Sleep(24 * time.Hour)
		}
	}

	wg.Add(1)
	go handleData(scraper.GetDataChannel(), &wg, rdb)
	defer wg.Wait()

}

func handleData(dataChannel chan dia.NFT, wg *sync.WaitGroup, rdb *models.RelDB) {
	defer wg.Done()

	for {
		fq, ok := <-dataChannel
		if !ok {
			log.Error("error")
			return
		}
		err := rdb.SetNFT(fq)
		if err != nil {
			log.Error(err)
		}
	}

}
