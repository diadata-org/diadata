package main

import (
	"flag"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/ethclient"

	nfttradescrapers "github.com/diadata-org/diadata/internal/pkg/nftTrade-scrapers"
	log "github.com/sirupsen/logrus"
)

func main() {

	wg := sync.WaitGroup{}

	rdb, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("relational datastore error: ", err)
	}

	scraperType := flag.String("exchange", "Opensea", "which NFT marketplace")
	flag.Parse()
	var scraper nfttradescrapers.NFTTradeScraper

	// connection, err := ethhelper.NewETHClient()
	// if err != nil {
	// 	log.Error("Error connecting Eth Client")
	// }
	wsConnection, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/3c7bc809e9174a85ad56ee9abcb68d32")
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	nfttradescrapers.GetData(wsConnection)
	// tx := nfttradescrapers.GetTxByHash(connection, common.HexToHash("0xe2afc7d40283d8c2f8760d4e0053767e681632a859fc5042982cba10a0db1fa7"))
	// fmt.Println("tx: ", tx)

	switch *scraperType {
	case "Opensea":
		log.Println("NFT Trade Scraper: Start scraping data from Opensea")
		scraper = nfttradescrapers.NewOpenseaScraper(rdb)
	default:
		for {
			time.Sleep(24 * time.Hour)
		}
	}

	wg.Add(1)
	go handleData(scraper.GetTradeChannel(), &wg, rdb)
	defer wg.Wait()

}

func handleData(dataChannel chan *dia.NFTTrade, wg *sync.WaitGroup, rdb *models.RelDB) {
	defer wg.Done()

	for {
		fq, ok := <-dataChannel
		if !ok {
			log.Error("error")
			return
		}
		rdb.SetNFTTrade(*fq)
	}

}
