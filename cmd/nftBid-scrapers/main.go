package main

import (
	"flag"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"

	nftbidscrapers "github.com/diadata-org/diadata/internal/pkg/nftBid-scrapers"
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
	var scraper nftbidscrapers.NFTBidScraper

	switch *scraperType {
	case "Cryptopunk":
		log.Println("NFT Data Scraper: Start scraping trades from Cryptopunk")
		scraper = nftbidscrapers.NewCryptoPunkScraper(rdb)
	default:
		for {
			time.Sleep(24 * time.Hour)
		}
	}

	wg.Add(1)
	go handleBids(scraper.GetBidChannel(), &wg, rdb)
	defer wg.Wait()

}

func handleBids(bidChannel chan dia.NFTBid, wg *sync.WaitGroup, rdb *models.RelDB) {
	defer wg.Done()
	for {
		nb, ok := <-bidChannel
		if !ok {
			log.Error("error")
			return
		}
		rdb.SetNFTBid(nb)
	}
}
