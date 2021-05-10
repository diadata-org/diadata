package main

import (
	"flag"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"

	nfttradescrapers "github.com/diadata-org/diadata/internal/pkg/nftTrade-scrapers"
	log "github.com/sirupsen/logrus"
)

func main() {

	wg := sync.WaitGroup{}

	rdb, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("relational datastore error: ", err)
	}

	scraperType := flag.String("exchange", "Opensea", "which NFT marketplace")
	flag.Parse()
	var scraper nfttradescrapers.NFTTradeScraper

	switch *scraperType {
	case "Opensea":
		log.Println("NFT Trade Scraper: Start scraping data from Opensea")
		scraper = nfttradescrapers.NewOpenseaScraper(rdb)
	default:
		for {
			time.Sleep(24 * time.Hour)
		}
	}

	wg.Add(1)
	go handleData(scraper.GetTradeChannel(), &wg, rdb)
	defer wg.Wait()

}

func handleData(dataChannel chan *dia.NFTTrade, wg *sync.WaitGroup, rdb *models.RelDB) {
	defer wg.Done()

	for {
		fq, ok := <-dataChannel
		if !ok {
			log.Error("error")
			return
		}
		rdb.SetNFTTrade(*fq)
	}

}
