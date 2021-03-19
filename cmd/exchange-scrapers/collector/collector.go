package main

import (
	"flag"
	"sync"
	"time"

	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
}

func handleTrades(c chan *dia.Trade, wg *sync.WaitGroup, w *kafka.Writer, exchange string) {
	lastTradeTime := time.Now()
	watchdogDelay := scrapers.Exchanges[exchange].WatchdogDelay
	t := time.NewTicker(time.Duration(watchdogDelay) * time.Second)
	for {
		select {
		case <-t.C:
			duration := time.Since(lastTradeTime)
			if duration > time.Duration(watchdogDelay)*time.Second {
				log.Error(duration)
				panic("frozen? ")
			}
		case t, ok := <-c:
			if !ok {
				wg.Done()
				log.Error("handleTrades")
				return
			}
			lastTradeTime = time.Now()
			kafkaHelper.WriteMessage(w, t)
		}
	}
}

var (
	exchange         = flag.String("exchange", "", "which exchange")
	onePairPerSymbol = flag.Bool("onePairPerSymbol", false, "one Pair max Per Symbol ?")
)

func init() {
	flag.Parse()
	if *exchange == "" {
		flag.Usage()
		log.Println(dia.Exchanges())
		for true {
			time.Sleep(24 * time.Hour)
		}
		// log.Fatal("exchange is required")
	}
}

// main manages all PairScrapers and handles incoming trade information
func main() {

	relDB, err := models.NewRelDataStore()
	if err != nil {
		log.Errorln("NewDataStore:", err)
	}

	pairsExchange, err := relDB.GetExchangePairSymbols(*exchange)
	log.Info("available exchangePairs:", len(pairsExchange))

	if err != nil || len(pairsExchange) == 0 {
		log.Error("error on GetExchangePairSymbols", err)
		cc := configCollectors.NewConfigCollectors(*exchange)
		pairsExchange = cc.AllPairs()
	}

	configApi, err := dia.GetConfig(*exchange)
	if err != nil {
		log.Warning("no config for exchange's api ", err)
	}
	es := scrapers.NewAPIScraper(*exchange, true, configApi.ApiKey, configApi.SecretKey, relDB)

	w := kafkaHelper.NewWriter(kafkaHelper.TopicTrades)
	defer w.Close()

	wg := sync.WaitGroup{}

	exchangePairs := make(map[string]string)

	// TO DO: Add check for new pairs. i.e. put a ticker around the following loop
	// and add a control whether new pairs are there.
	for _, configPair := range pairsExchange {
		dontAddPair := false
		if *onePairPerSymbol {
			_, dontAddPair = exchangePairs[configPair.Symbol]
			exchangePairs[configPair.Symbol] = configPair.Symbol
		}
		if dontAddPair {
			log.Println("Skipping pair:", configPair.Symbol, configPair.ForeignName, "on exchange", *exchange)
		} else {
			log.Println("Adding pair:", configPair.Symbol, configPair.ForeignName, "on exchange", *exchange)
			_, err := es.ScrapePair(dia.ExchangePair{
				Symbol:      configPair.Symbol,
				ForeignName: configPair.ForeignName})
			if err != nil {
				log.Println(err)
			} else {
				wg.Add(1)
			}
		}
		defer wg.Wait()
	}
	go handleTrades(es.Channel(), &wg, w, *exchange)
}
