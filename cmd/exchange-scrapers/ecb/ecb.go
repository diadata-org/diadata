package main

import (
	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	"sync"
)

// pairs contains all pairs currently supported by the DIA scrapers
var (
	pairs = []string{
		"EURUSD",
		"EURJPY",
		//"EURBGN",
		//"EURCZK",
		//"EURDKK",
		"EURGBP",
		//	"EURHUF",
		//	"EURPLN",
		//	"EURRON",
		"EURSEK",
		"EURCHF",
		"EURNOK",
		//	"EURHRK",
		//	"EURRUB",
		//	"EURTRY",
		"EURAUD",
		//	"EURBRL",
		"EURCAD",
		"EURCNY",
		//	"EURHKD",
		//	"EURIDR",
		//	"EURILS",
		//	"EURINR",
		"EURKRW",
		//	"EURMXN",
		//	"EURMYR",
		//	"EURNZD",
		//	"EURPHP",
		//	"EURSGD",
		//	"EURTHB",
		//	"EURZAR",
	}
)

var usdFor1Euro = -1.0

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
			ds.SetPriceUSD(symbol, 1)
			ds.SetPriceEUR(symbol, 1/usdFor1Euro)
			ds.SetPriceUSD("EUR", usdFor1Euro)
			ds.SetPriceEUR("EUR", 1)
		} else {
			if usdFor1Euro > 0 {
				log.Info("setting ", symbol, usdFor1Euro/t.Price)
				ds.SetPriceUSD(symbol, usdFor1Euro/t.Price)
				ds.SetPriceEUR(symbol, 1/t.Price)
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
		sECB := scrapers.NewECBScraper(ds)
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
		sECB.Update()
		go handleTrades(sECB.Channel(), &wg, ds)
		defer wg.Wait()
	}
}
