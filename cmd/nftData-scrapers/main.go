package main

import (
	"errors"
	"flag"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/jackc/pgconn"

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

func handleData(dataChannel chan dia.NFT, wg *sync.WaitGroup, rdb *models.RelDB) {
	defer wg.Done()

	for {
		nft, ok := <-dataChannel
		if !ok {
			log.Error("error")
			return
		}
		log.Info("set nft: ", nft)
		err := rdb.SetNFT(nft)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == "23505" {
					log.Infof("nft %v from class %s already in db. continue.", nft.TokenID, nft.NFTClass.Name)
					continue
				} else {
					log.Errorf("postgres error saving nft %v: %v", nft.NFTClass.Name, nft.TokenID)
				}
			} else {
				log.Errorf("Error saving nft %v: %v", nft.NFTClass.Name, nft.TokenID)
			}
		} else {
			log.Info("successfully set nft ", nft.NFTClass.Name, nft.TokenID)
		}
	}

}
