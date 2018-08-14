package main

import (
	"fmt"
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/exchange-scrapers"
	"github.com/tkanos/gonfig"
	"log"
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

// main manages all PairScrapers and handles incoming trade information
func main() {

	s := map[string]scrapers.APIScraper{}

	var configConnector dia.ConfigConnector

	configFile := "../config/exchange-scrapers.json"
	err := gonfig.GetConf(configFile, &configConnector)

	if err != nil {
		fmt.Printf("error loading configFile")
	}
	wg := sync.WaitGroup{}

	for _, configPair := range configConnector.Coins {

		fmt.Println("Adding pair:", configPair.Symbol, " on exchange ", configPair.Exchange)

		_, ok := s[configPair.Exchange]
		if ok == false {
			var configApi dia.ConfigApi
			configFileApi := "../config/secrets/" + configPair.Exchange + ".json"
			err = gonfig.GetConf(configFileApi, &configApi)
			if err != nil {
				fmt.Println(err)
			} else {
				switch e := configPair.Exchange; e {
				case dia.BinanceExchange:
					s[configPair.Exchange] = scrapers.NewBinanceScraper(configApi.ApiKey, configApi.SecretKey)
				case dia.BitfinexExchange:
					s[configPair.Exchange] = scrapers.NewBitfinexScraper(configApi.ApiKey, configApi.SecretKey)
				case dia.CoinBaseExchange:
					s[configPair.Exchange] = scrapers.NewCoinBaseScraper()
				case dia.KrakenExchange:
					s[configPair.Exchange] = scrapers.NewKrakenScraper(configApi.ApiKey, configApi.SecretKey)
				default:
					fmt.Printf("Unknown exchange %s.", e)
					return
				}
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
		defer wg.Wait()
	}
}
