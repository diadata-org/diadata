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

var mode *string

func main() {

	var (
		w            *kafka.Writer
		NFTExchanges = make(map[string]dia.NFTExchange)
	)

	rdb, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("get rel datastore: ", err)
	}

	exchanges, err := rdb.GetAllNFTExchanges()
	if err != nil {
		log.Fatal("get all exchanges: ", err)
	}
	for _, exchange := range exchanges {
		NFTExchanges[exchange.Name] = exchange
	}

	// if testOpenSeaScraper := false; testOpenSeaScraper {
	// 	w = kafkaHelper.NewWriter(kafkaHelper.TopicNFTTradesTest)

	// 	rdb, err := models.NewRelDataStore()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	scraper := nfttradescrapers.NewOpenSeaScraper(rdb, NFTExchanges[dia.Opensea])
	// 	go func() { time.Sleep(3 * time.Minute); scraper.Close() }()

	// 	wg := sync.WaitGroup{}
	// 	wg.Add(1)
	// 	go handleData(scraper.GetTradeChannel(), &wg, w, rdb)
	// 	wg.Wait()

	// 	return
	// }
	w = kafkaHelper.NewWriter(kafkaHelper.TopicNFTTrades)

	wg := sync.WaitGroup{}


	scraperType := flag.String("nftclass", "Cryptopunk", "which NFT class")
	mode = flag.String("mode", "", "run local without kafka.")
	flag.Parse()
	var scraper nfttradescrapers.NFTTradeScraper

	switch *scraperType {
	case dia.CryptoPunks:
		log.Infoln("NFT Trades Scraper: Start scraping trades from Cryptopunks")
		scraper = nfttradescrapers.NewCryptoPunkScraper(rdb, NFTExchanges[dia.CryptoPunks])
	case dia.CryptoKitties:
		log.Infoln("NFT Trades Scraper: Start scraping trades from CryptoKitties")
		scraper = nfttradescrapers.NewCryptoKittiesScraper(rdb)
	case dia.Topshot:
		log.Infoln("NFT Trades Scraper: Start scraping trades from NBA Topshot")
		scraper = nfttradescrapers.NewNBATopshotScraper(rdb)
	case dia.X2Y2:
		log.Infoln("NFT Trades Scraper: Start scraping trades from X2Y2")
		scraper = nfttradescrapers.NewX2Y2Scraper(rdb, NFTExchanges[dia.X2Y2])
	case dia.Opensea:
		log.Infoln("NFT Trades Scraper: Start scraping trades from Opensea")
		scraper = nfttradescrapers.NewOpenSeaScraper(rdb, NFTExchanges[dia.Opensea])
	case dia.OpenseaBAYC:
		log.Infoln("NFT Trades Scraper: Start scraping trades from Opensea")
		scraper = nfttradescrapers.NewOpenSeaBAYCScraper(rdb, NFTExchanges[dia.Opensea])
	case dia.OpenseaSeaport:
		log.Infoln("NFT Trades Scraper: Start scraping trades from Opensea Seaport contract")
		scraper = nfttradescrapers.NewOpenSeaSeaportScraper(rdb, NFTExchanges[dia.Opensea])
	case dia.BlurV1:
		log.Infoln("NFT Trades Scraper: Start scraping trades from BlurV1 contract")
		scraper = nfttradescrapers.NewBlurV1Scraper(rdb, NFTExchanges[dia.BlurV1])
	case dia.LooksRare:
		log.Infoln("NFT Trades Scraper: Start scraping trades from LooksRare")
		scraper = nfttradescrapers.NewLooksRareScraper(rdb, NFTExchanges[dia.LooksRare])
	case dia.TofuNFTAstar:
		log.Infoln("NFT Trades Scraper: Start scraping trades from TofuNFT on Astar")
		scraper = nfttradescrapers.NewTofuNFTScraper(rdb, NFTExchanges[dia.TofuNFTAstar])
	case dia.TofuNFTBinanceSmartChain:
		log.Infoln("NFT Trades Scraper: Start scraping trades from TofuNFT on BinanceSmartChain")
		scraper = nfttradescrapers.NewTofuNFTScraper(rdb, NFTExchanges[dia.TofuNFTBinanceSmartChain])
	case dia.MagicEden:
		log.Infoln("NFT Trades Scraper: Start scraping trades from MagicEden on Solana")
		scraper = nfttradescrapers.NewMagicEdenScraper(rdb, NFTExchanges[dia.MagicEden])

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
			if *mode != "local" {
				err = writeNFTTradeToKafka(w, &trade)
				if err != nil {
					log.Errorf("Error writing trade to kafka with tx hash %s: %v", trade.TxHash, err)
				}
			}

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
