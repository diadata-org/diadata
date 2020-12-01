package main

import (
	"sync"
	"time"

	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

var (
	pairs = []string{
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

var usdFor1Euro = -1.0 // EUR/USD

// handleTrades delegates trade information to Kafka
func handleTrades(c chan *dia.Trade, wg *sync.WaitGroup, ds models.Datastore) {
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

			fq := []*models.FiatQuotation{{
				QuoteCurrency: "EUR",
				BaseCurrency:  "USD",
				Price:         t.Price,
				Source:        "ECB",
				Time:          time.Now(),
			}}

			err := ds.SetFiatPriceUSD(fq)
			if err != nil {
				log.Printf("Error on SetFiatPriceUSD: %v\n", err)
			}
		} else {
			if usdFor1Euro > 0 {
				log.Info("setting ", symbol, usdFor1Euro/t.Price) // compute Symbol/USD

				fq := []*models.FiatQuotation{{
					QuoteCurrency: symbol,
					BaseCurrency:  "USD",
					Price:         usdFor1Euro / t.Price,
					Source:        "ECB",
					Time:          time.Now(),
				}}

				err := ds.SetFiatPriceUSD(fq)
				if err != nil {
					log.Printf("Error on SetFiatPriceUSD: %v\n", err)
				}
			}
		}
	}
}

// main manages all PairScrapers and handles incoming trade information
func main() {
	wg := sync.WaitGroup{}
	ds, err := models.NewDataStore()

	if err != nil {
		log.Errorln("NewDataStore:", err)
	} else {
		// Populate historical prices
		go scrapers.Populate(ds, pairs)

		sECB := scrapers.SpawnECBScraper(ds)
		defer sECB.Close()

		for _, pair := range pairs {
			_, err := sECB.ScrapePair(dia.Pair{
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

		go handleTrades(sECB.Channel(), &wg, ds)
		defer wg.Wait()
	}
}
