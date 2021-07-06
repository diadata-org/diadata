package main

import (
	"errors"
	"flag"
	"fmt"
	"sync"
	"time"

	nfttradescrapers "github.com/diadata-org/diadata/internal/pkg/nftTrade-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/jackc/pgconn"

	log "github.com/sirupsen/logrus"
)

func main() {
	if testOpenSeaScraper := false; testOpenSeaScraper {
		rdb, err := models.NewRelDataStore()
		if err != nil {
			panic(err)
		}

		scraper := nfttradescrapers.NewOpenSeaScraper(rdb)
		go func() { time.Sleep(3 * time.Minute); scraper.Close() }()

		wg := sync.WaitGroup{}
		wg.Add(1)
		go handleData(scraper.GetTradeChannel(), &wg, rdb)
		wg.Wait()

		return
	}

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
		trade, ok := <-tradeChannel
		if !ok {
			log.Error("error")
			return
		}
		if 1 < 0 {
			fmt.Printf("got trade: %s -> (%s) -> %s for %s (%.4f USD) \n", trade.FromAddress, trade.NFT.NFTClass.Name, trade.ToAddress, trade.CurrencySymbol, trade.PriceUSD)
		}
		err := rdb.SetNFTTrade(trade)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == "23505" {
					log.Infof("trade with tx hash %s already in db. continue.", trade.TxHash)
					continue
				} else {
					log.Errorf("postgres error saving trade with tx hash %s: %v", trade.TxHash, err)
				}
			} else {
				log.Errorf("Error saving trade with tx hash %s: %v", trade.TxHash, err)
			}
		} else {
			log.Infof("successfully set trade with tx hash %s", trade.TxHash)
		}
	}

}
