package main

import (
	"flag"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var (
	log       *logrus.Logger
	exchange  = flag.String("exchange", "", "which exchange")
	pairsfile = flag.Bool("pairsfile", false, "read pairs from json file in config folder.")
)

func init() {
	log = logrus.New()
	flag.Parse()
	if *exchange == "" {
		flag.Usage()
		for e := range scrapers.Exchanges {
			log.Info("exchange: ", e)
		}
		for {
			time.Sleep(24 * time.Hour)
		}
	}
	if !isValidExchange(*exchange) {
		log.Fatal("Invalid exchange string: ", *exchange)
	}

}

// main manages all PairScrapers and handles incoming trade information
func main() {

	log.Infof("start collector for %s in test-space...", *exchange)

	relDB, err := models.NewRelDataStore()
	if err != nil {
		log.Errorln("NewDataStore:", err)
	}

	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore: ", err)
	}

	// Fetch exchange pairs from database or json file in config folder.
	var pairsExchange []dia.ExchangePair
	if !*pairsfile {
		pairsExchange, err = relDB.GetExchangePairSymbols(*exchange)
		if err != nil {
			log.Fatal("fetch pairs from database: ", err)
		}
	} else {
		log.Error("error on GetExchangePairSymbols", err)
		cc := configCollectors.NewConfigCollectors(*exchange, ".json")
		pairsExchange = cc.AllPairs()
	}
	log.Info("available exchangePairs:", len(pairsExchange))

	// Make API scraper.
	configApi, err := dia.GetConfig(*exchange)
	if err != nil {
		log.Warning("no config for exchange's api ", err)
	}
	es := scrapers.NewAPIScraper(*exchange, true, configApi.ApiKey, configApi.SecretKey, relDB)

	// Set up kafka writer.
	w := kafkaHelper.NewWriter(kafkaHelper.TopicTradesTest)
	log.Info("writer topic: ", w.Topic)

	defer func() {
		err := w.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	wg := sync.WaitGroup{}

	if scrapers.Exchanges[*exchange].Centralized || scrapers.ExchangeDuplicates[*exchange].Centralized {

		// Scrape pairs for CEX scrapers.
		for _, configPair := range pairsExchange {
			log.Println("Adding pair:", configPair.Symbol, configPair.ForeignName, "on exchange", *exchange)
			_, err := es.ScrapePair(dia.ExchangePair{
				Symbol:      configPair.Symbol,
				ForeignName: configPair.ForeignName})
			if err != nil {
				log.Println(err)
			} else {
				wg.Add(1)
			}
			defer wg.Wait()
		}

	} else {

		// Subscription to pool events managed inside scraper for DEX and Bridge scrapers.
		wg.Add(1)
		defer wg.Wait()

	}
	go handleTrades(es.Channel(), &wg, w, ds, *exchange)
}

func handleTrades(c chan *dia.Trade, wg *sync.WaitGroup, w *kafka.Writer, ds *models.DB, exchange string) {
	lastTradeTime := time.Now()
	watchdogDelay := scrapers.Exchanges[exchange].WatchdogDelay
	if watchdogDelay == 0 {
		watchdogDelay = scrapers.ExchangeDuplicates[exchange].WatchdogDelay
	}
	tk := time.NewTicker(time.Duration(watchdogDelay) * time.Second)

	for {
		select {
		case <-tk.C:
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

			// Trades are sent to the tradesblockservice through a kafka channel.
			err := writeTradeToKafka(w, t)
			if err != nil {
				log.Error(err)
			}

		}

	}
}

func writeTradeToKafka(w *kafka.Writer, t *dia.Trade) error {
	err := kafkaHelper.WriteMessage(w, t)
	return err
}

func isValidExchange(estring string) bool {
	for e := range scrapers.Exchanges {
		if e == estring {
			return true
		}
	}
	for e := range scrapers.ExchangeDuplicates {
		if e == estring {
			return true
		}
	}
	return false
}
