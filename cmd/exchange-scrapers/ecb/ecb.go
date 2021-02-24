package main

import (
	"sync"

	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/sirupsen/logrus"
)

var (
	usdFor1Euro = -1.0
	log         *logrus.Logger
	pairs       = []string{
		"EURUSD",
		"EURJPY",
		"EURBGN",
		"EURCZK",
		"EURDKK",
		"EURGBP",
		"EURHUF",
		"EURPLN",
		"EURRON",
		"EURSEK",
		"EURCHF",
		"EURNOK",
		"EURHRK",
		"EURRUB",
		"EURTRY",
		"EURAUD",
		"EURBRL",
		"EURCAD",
		"EURCNY",
		"EURHKD",
		"EURIDR",
		"EURILS",
		"EURINR",
		"EURKRW",
		"EURMXN",
		"EURMYR",
		"EURNZD",
		"EURPHP",
		"EURSGD",
		"EURTHB",
		"EURZAR",
	}
)

func init() {
	log = logrus.New()
}

// handleTrades delegates trade information to Kafka
func handleTrades(c chan *dia.Trade, wg *sync.WaitGroup, ds models.Datastore, rdb models.RelDB) {
	for {
		t, ok := <-c
		if !ok {
			log.Error("error")
			return
		}

		symbol := t.Symbol[len(t.Pair)-3:]

		if symbol == "USD" {
			log.Println(symbol, t.Symbol)
			usdFor1Euro = t.Price

			asset, err := rdb.GetFiatAssetBySymbol("EUR")
			if err != nil {
				log.Error("fetching fiat asset EUR: ", err)
			}
			err = ds.SetAssetPriceUSD(asset, t.Price)
			if err != nil {
				log.Errorf("setting asset quotation for asset %s: %v", asset.Symbol, err)
			}

		} else {
			if usdFor1Euro > 0 {
				asset, err := rdb.GetFiatAssetBySymbol(symbol)
				if err != nil {
					log.Errorf("fetching fiat asset %s: %v", symbol, err)
				}
				err = ds.SetAssetPriceUSD(asset, usdFor1Euro/t.Price)
				if err != nil {
					log.Errorf("setting asset quotation for asset %s: %v", asset.Symbol, err)
				}
			}
		}
	}
}

// main manages all PairScrapers and handles incoming trade information
func main() {
	wg := sync.WaitGroup{}
	ds, err := models.NewDataStore()
	rdb, err := models.NewRelDataStore()

	if err != nil {
		log.Errorln("NewDataStore:", err)
	} else {
		// Populate historical prices
		// go scrapers.Populate(ds, pairs)

		sECB := scrapers.SpawnECBScraper(ds)
		defer sECB.Close()

		for _, pair := range pairs {
			_, err := sECB.ScrapePair(dia.ExchangePair{
				Symbol:      pair,
				ForeignName: pair,
			})
			if err != nil {
				log.Fatal(err)
			}
			wg.Add(1)
		}

		// This should be uncommented in case "go mainLoop.go" is deleted from SpawnECBScraper
		// go sECB.MainLoop()

		go handleTrades(sECB.Channel(), &wg, ds, *rdb)
		defer wg.Wait()
	}
}
