package main

import (
	"errors"
	"flag"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/jackc/pgconn"

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
		log.Println("NFT Bids Scraper: Start scraping bids from Cryptopunk")
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
		bid, ok := <-bidChannel
		if !ok {
			log.Error("error")
			return
		}
		err := rdb.SetNFTBid(bid)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == "23505" {
					log.Infof("bid with tx hash %s already in db. continue.", bid.TxHash)
					continue
				} else {
					log.Errorf("postgres error saving bid with tx hash %s: %v", bid.TxHash, err)
				}
			} else {
				log.Errorf("Error saving bid with tx hash %s: %v", bid.TxHash, err)
			}
		} else {
			log.Infof("successfully set bid with tx hash %s", bid.TxHash)
		}
	}
}
