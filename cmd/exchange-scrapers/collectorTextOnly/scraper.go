package main

import (
	"flag"
	"fmt"
	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	log "github.com/sirupsen/logrus"
	"sync"
)

// pairs contains all pairs currently supported by the DIA scrapers

// handleTrades delegates trade information to Kafka
func handleTrades(c chan *dia.Trade, wg *sync.WaitGroup) {
	for {
		t, ok := <-c
		if !ok {
			log.Error("error")
			return
		}
		log.Printf("handleTrades: %v\n", t)
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

	cc := configCollectors.NewConfigCollectors(*exchange)

	wg := sync.WaitGroup{}

	for _, configPair := range cc.AllPairs() {

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
				go handleTrades(aPIScraper.Channel(), &wg)
			} else {
				fmt.Println("Couldn't create APIScraper for ", configPair.Exchange)
			}
		}
		es := s[configPair.Exchange]
		if es != nil {
			_, err := es.ScrapePair(dia.Pair{
				Symbol:      configPair.Symbol,
				ForeignName: configPair.ForeignName})
			if err != nil {
				log.Println(err)
			} else {
				wg.Add(1)
			}
		}
	}
	defer wg.Wait()
}
