package main

import (
	"errors"
	"flag"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	nfttradescrapers "github.com/diadata-org/diadata/pkg/dia/nft/nftTrade-scrapers"
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
	case "CryptoPunks":
		log.Println("NFT Trades Scraper: Start scraping trades from Cryptopunks")
		scraper = nfttradescrapers.NewCryptoPunkScraper(rdb)
	case "CryptoKitties":
		log.Println("NFT Trades Scraper: Start scraping trades from CryptoKitties")
		scraper = nfttradescrapers.NewCryptoKittiesScraper(rdb)
	case "Topshot":
		log.Println("NFT Trades Scraper: Start scraping trades from NBA Topshot")
		scraper = nfttradescrapers.NewNBATopshotScraper(rdb)
	case "Opensea":
		log.Println("NFT Trades Scraper: Start scraping trades from Opensea")
		scraper = nfttradescrapers.NewOpenSeaScraper(rdb)
	case "OpenseaBAYC":
		log.Println("NFT Trades Scraper: Start scraping trades from Opensea")
		scraper = nfttradescrapers.NewOpenSeaBAYCScraper(rdb)
	case "OpenseaSeaport":
		log.Println("NFT Trades Scraper: Start scraping trades from Opensea Seaport contract")
		scraper = nfttradescrapers.NewOpenSeaSeaportScraper(rdb)
	case "LooksRare":
		log.Println("NFT Trades Scraper: Start scraping trades from LooksRare")
		scraper = nfttradescrapers.NewLooksRareScraper(rdb)
	case "TofuNFT-Astar":
		log.Println("NFT Trades Scraper: Start scraping trades from TofuNFT on Astar")
		scraper = nfttradescrapers.NewTofuNFTScraper(dia.ASTAR, rdb)
	case "TofuNFT-BinanceSmartChain":
		log.Println("NFT Trades Scraper: Start scraping trades from TofuNFT on BinanceSmartChain")
		scraper = nfttradescrapers.NewTofuNFTScraper(dia.BINANCESMARTCHAIN, rdb)

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
			log.Infof("got trade: %s -> (%s) -> %s for %s (%.4f USD) \n", trade.FromAddress, trade.NFT.NFTClass.Name, trade.ToAddress, trade.Currency.Symbol, trade.PriceUSD)
		}

		err := rdb.SetNFTTradeToTable(trade, models.NfttradeCurrTable)
		// err := rdb.SetNFTTradeToTable(trade, models.NfttradeSumeriaTable)
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
