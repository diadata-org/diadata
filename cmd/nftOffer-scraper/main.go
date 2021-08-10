package main

import (
	"errors"
	"flag"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/jackc/pgconn"

	nftofferscrapers "github.com/diadata-org/diadata/internal/pkg/nftOffer-scrapers"
	log "github.com/sirupsen/logrus"
)

func main() {

	wg := sync.WaitGroup{}

	rdb, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("relational datastore error: ", err)
	}

	scraperType := flag.String("nftclass", "CryptoPunks", "which NFT class")
	flag.Parse()
	var scraper nftofferscrapers.NFTOfferScraper

	switch *scraperType {
	case "CryptoPunks":
		log.Println("NFT Offers Scraper: Start scraping bids from CryptoPunks")
		scraper = nftofferscrapers.NewCryptoPunksScraper(rdb)
	case "CryptoKitties":
		log.Println("NFT Offers Scraper: Start scraping bids from CryptoKitties")
		scraper = nftofferscrapers.NewCryptokittiesScraper(rdb)
	default:
		for {
			time.Sleep(24 * time.Hour)
		}
	}

	wg.Add(1)
	go handleOffers(scraper.GetOfferChannel(), &wg, rdb)
	defer wg.Wait()

}

func handleOffers(offerChannel chan dia.NFTOffer, wg *sync.WaitGroup, rdb *models.RelDB) {
	defer wg.Done()
	for {
		offer, ok := <-offerChannel
		if !ok {
			log.Error("error")
			return
		}
		err := rdb.SetNFTOffer(offer)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == "23505" {
					log.Infof("offer with tx hash %s already in db. continue.", offer.TxHash)
					continue
				} else {
					log.Errorf("postgres error saving offer with tx hash %s: %v", offer.TxHash, err)
				}
			} else {
				log.Errorf("Error saving offer with tx hash %s: %v", offer.TxHash, err)
			}
		} else {
			log.Infof("successfully set offer with tx hash %s", offer.TxHash)
		}
	}
}
