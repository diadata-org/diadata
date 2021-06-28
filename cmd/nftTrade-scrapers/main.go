package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	nfttradescrapers "github.com/diadata-org/diadata/internal/pkg/nftTrade-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"

	log "github.com/sirupsen/logrus"
)

func main() {
	if testOpenSeaScraper := true; testOpenSeaScraper {
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
		fq, ok := <-tradeChannel
		if !ok {
			log.Error("error")
			return
		}
		if 1 > 0 {
			fmt.Printf("got trade: %s -> (%s) -> %s for %s %s (%.4f USD) \n", fq.FromAddress.Hex(), fq.NFT.NFTClass.Name, fq.ToAddress.Hex(), fq.PriceDec.String(), fq.CurrencySymbol, fq.PriceUSD)
		}
		// rdb.SetNFT(fq)
	}

}
