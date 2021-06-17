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

	scraperType := flag.String("nftclass", "Cryptopunk", "which NFT class")
	flag.Parse()
	var scraper nfttradescrapers.NFTTradeScraper

	switch *scraperType {
	case "Cryptopunk":
		log.Println("NFT Data Scraper: Start scraping trades from Cryptopunk")
		scraper = nfttradescrapers.NewCryptoPunkScraper(rdb)
	default:
		for {
			time.Sleep(24 * time.Hour)
		}
	}

	wg.Add(1)
	go handleData(scraper.GetTradeChannel(), &wg, rdb)
	defer wg.Wait()

}

func handleData(tradeChannel chan dia.NFTTrade, wg *sync.WaitGroup, rdb *models.RelDB) {
	defer wg.Done()

	for {
		fq, ok := <-tradeChannel
		if !ok {
			log.Error("error")
			return
		}
		if 1 > 2 {
			log.Info(fq)
		}
		// rdb.SetNFT(fq)
	}

}
