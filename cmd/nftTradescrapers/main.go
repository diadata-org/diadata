package main

import (
	"errors"
	"flag"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	nfttradescrapers "github.com/diadata-org/diadata/pkg/dia/nft/nftTrade-scrapers"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/jackc/pgconn"
	"github.com/segmentio/kafka-go"

	log "github.com/sirupsen/logrus"
)

func main() {

	var (
		w *kafka.Writer
	)

	if testOpenSeaScraper := false; testOpenSeaScraper {
		w = kafkaHelper.NewWriter(kafkaHelper.TopicNFTTradesTest)

		rdb, err := models.NewRelDataStore()
		if err != nil {
			panic(err)
		}

		scraper := nfttradescrapers.NewOpenSeaScraper(rdb)
		go func() { time.Sleep(3 * time.Minute); scraper.Close() }()

		wg := sync.WaitGroup{}
		wg.Add(1)
		go handleData(scraper.GetTradeChannel(), &wg, w, rdb)
		wg.Wait()

		return
	}
	w = kafkaHelper.NewWriter(kafkaHelper.TopicNFTTrades)

	wg := sync.WaitGroup{}

	rdb, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("relational datastore error: ", err)
	}

	scraperType := flag.String("nftclass", "Cryptopunk", "which NFT class")
	flag.Parse()
	var scraper nfttradescrapers.NFTTradeScraper

	switch *scraperType {
	case dia.CryptoPunks:
		log.Infoln("NFT Trades Scraper: Start scraping trades from Cryptopunks")
		scraper = nfttradescrapers.NewCryptoPunkScraper(rdb)
	case dia.CryptoKitties:
		log.Infoln("NFT Trades Scraper: Start scraping trades from CryptoKitties")
		scraper = nfttradescrapers.NewCryptoKittiesScraper(rdb)
	case dia.Topshot:
		log.Infoln("NFT Trades Scraper: Start scraping trades from NBA Topshot")
		scraper = nfttradescrapers.NewNBATopshotScraper(rdb)
	case dia.X2Y2:
		log.Infoln("NFT Trades Scraper: Start scraping trades from X2Y2")
		scraper = nfttradescrapers.NewX2Y2Scraper(rdb)
	case dia.Opensea:
		log.Infoln("NFT Trades Scraper: Start scraping trades from Opensea")
		scraper = nfttradescrapers.NewOpenSeaScraper(rdb)
	case dia.OpenseaBAYC:
		log.Infoln("NFT Trades Scraper: Start scraping trades from Opensea")
		scraper = nfttradescrapers.NewOpenSeaBAYCScraper(rdb)
	case dia.OpenseaSeaport:
		log.Infoln("NFT Trades Scraper: Start scraping trades from Opensea Seaport contract")
		scraper = nfttradescrapers.NewOpenSeaSeaportScraper(rdb)
	case dia.LooksRare:
		log.Infoln("NFT Trades Scraper: Start scraping trades from LooksRare")
		scraper = nfttradescrapers.NewLooksRareScraper(rdb)
	case dia.TofuNFTAstar:
		log.Infoln("NFT Trades Scraper: Start scraping trades from TofuNFT on Astar")
		scraper = nfttradescrapers.NewTofuNFTScraper(dia.ASTAR, rdb)
	case dia.TofuNFTBinanceSmartChain:
		log.Infoln("NFT Trades Scraper: Start scraping trades from TofuNFT on BinanceSmartChain")
		scraper = nfttradescrapers.NewTofuNFTScraper(dia.BINANCESMARTCHAIN, rdb)

	default:
		for {
			time.Sleep(24 * time.Hour)
		}
	}

	wg.Add(1)
	go handleData(scraper.GetTradeChannel(), &wg, w, rdb)
	defer wg.Wait()
}

func handleData(tradeChannel chan dia.NFTTrade, wg *sync.WaitGroup, w *kafka.Writer, rdb *models.RelDB) {
	defer wg.Done()

	for {
		trade, ok := <-tradeChannel
		if !ok {
			log.Error("error")
			return
		}

		log.Infof("got trade: %s -> (%s) -> %s for %v %s (%.4f USD) \n", trade.FromAddress, trade.NFT.NFTClass.Name, trade.ToAddress, trade.Price, trade.Currency.Symbol, trade.PriceUSD)

		err := rdb.SetNFTTradeToTable(trade, models.NfttradeCurrTable)
		writeNFTTradeToKafka(w, &trade)
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

func writeNFTTradeToKafka(w *kafka.Writer, t *dia.NFTTrade) error {
	// Write trade to Kafka.
	err := kafkaHelper.WriteMessage(w, t)
	if err != nil {
		return err
	}

	return nil
}
