package main

import (
	"flag"
	"fmt"
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/exchange-scrapers/scrapers"
	"github.com/tkanos/gonfig"
	"log"
	"os"
	"sync"
)

// pairs contains all pairs currently supported by the DIA scrapers

// handleTrades delegates trade information to Kafka
func handleTrades(ps scrapers.PairScraper, wg *sync.WaitGroup) {
	for {
		t, ok := <-ps.Channel()

		if !ok {
			// log error and return
			if ps.Error() != nil {
				log.Printf("Error: %s\n", ps.Error())
			} else {
				log.Printf("PairScraper for %s was shut down by user", ps.Pair())
			}
			wg.Done()
			return
		}
		log.Printf("Trade: %v\n", t)
	}
}

var (
	exchange = flag.String("exchange", "", "which exchange")
)

func init() {
	flag.Parse()
	if *exchange == "" {
		flag.Usage()
		log.Println(dia.Exchanges())
		log.Fatal("exchange is required")
	}
}

// main manages all PairScrapers and handles incoming trade information
func main() {

	s := map[string]scrapers.APIScraper{}

	var configConnector dia.ConfigConnector

	configFile := os.Getenv("GOPATH") + "/src/github.com/diadata-org/api-golang/exchange-scrapers/config/exchange-scrapers.json"
	err := gonfig.GetConf(configFile, &configConnector)

	if err != nil {
		fmt.Println("error loading configFile", configFile)
	}
	wg := sync.WaitGroup{}

	for _, configPair := range configConnector.Coins {

		if *exchange == configPair.Exchange {

			fmt.Println("Adding pair:", configPair.Symbol, "(", configPair.ForeignName, ") on exchange", configPair.Exchange)

			_, ok := s[configPair.Exchange]
			if ok == false {

				configExchangeApi, err := dia.GetConfig(configPair.Exchange)
				if err != nil {
					fmt.Println(err)
				}
				aPIScraper := scrapers.NewAPIScraper(configPair.Exchange, configExchangeApi.ApiKey, configExchangeApi.SecretKey)
				if s != nil {
					s[configPair.Exchange] = aPIScraper
				} else {
					fmt.Println("Couldnt create APIScraper for ", configPair.Exchange)
				}
			}
			es := s[configPair.Exchange]
			if es != nil {
				ps, err := es.ScrapePair(dia.Pair{
					Symbol:      configPair.Symbol,
					ForeignName: configPair.ForeignName})
				if err != nil {
					log.Println(err)
				} else {
					go handleTrades(ps, &wg)
					wg.Add(1)
				}
			}
		}
		defer wg.Wait()
	}
}
