package main

import (
	"errors"
	"flag"
	"sync"
	"time"

	blockscrapers "github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/block-scrapers"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/jackc/pgconn"

	log "github.com/sirupsen/logrus"
)

func main() {

	wg := sync.WaitGroup{}

	rdb, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}

	scraperType := flag.String("blockchain", "Ethereum", "which blockchain")
	flag.Parse()
	var blockscraper blockscrapers.BlockScraperInterface

	switch *scraperType {
	case "Ethereum":
		log.Println("Block-scraper: Start scraping block data from Ethereum")
		blockscraper = blockscrapers.NewEthereumScraper(rdb)
	default:
		for {
			time.Sleep(24 * time.Hour)
		}
	}

	wg.Add(1)
	go handleBlockData(blockscraper.GetDataChannel(), &wg, rdb)
	defer wg.Wait()

}

func handleBlockData(blockdatachannel chan dia.BlockData, wg *sync.WaitGroup, rdb *models.RelDB) {
	defer wg.Done()

	for {
		blockdata, ok := <-blockdatachannel
		if !ok {
			log.Error("blockdatachannel error")
			return
		}
		err := rdb.SetBlockData(blockdata)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == "23505" {
					log.Infof("block %v from chain %s already in db. continue.", blockdata.BlockNumber, blockdata.BlockchainName)
					continue
				} else {
					log.Errorf("postgres error saving block %v: %v", blockdata.BlockNumber, err)
					return
				}
			} else {
				log.Errorf("Error saving block %v from chain %s: %v", blockdata.BlockNumber, blockdata.BlockchainName, err)
				return
			}
		} else {
			log.Infof("successfully set block %v from chain %s \n", blockdata.BlockNumber, blockdata.BlockchainName)
		}
	}

}
