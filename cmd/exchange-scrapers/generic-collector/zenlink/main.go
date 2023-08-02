package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

var (
	exchange = flag.String("exchange", "", "which exchange")
	// mode==current:		default mode. Trades are forwarded to TBS and FBS.
	// mode==storeTrades:	trades are not forwarded to TBS and FBS and stored as raw trades in influx.
	// mode==estimation:	trades are forwarded to tradesEstimationService, i.e. same as storeTrades mode
	//						but estimatedUSDPrice is filled by tradesEstimationService.
	// mode==historical:	trades are sent through kafka to TBS in tradesHistorical topic.
	mode              = flag.String("mode", "current", "either storeTrades, current, historical or estimation")
	pairsfile         = flag.Bool("pairsfile", false, "read pairs from json file in config folder.")
	replicaKafkaTopic string
)

func init() {
	flag.Parse()
	replicaKafkaTopic = utils.Getenv("REPLICA_KAFKA_TOPIC", "false")
}

// main manages all PairScrapers and handles incoming trade information
func main() {

	log.Infof("start collector for %s in %s mode...", *exchange, *mode)

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

	diaExchange := scrapers.Exchanges[*exchange]
	es := scrapers.NewZenlinkScraper(diaExchange, true)

	// Set up kafka writers for various modes.
	var (
		w *kafka.Writer
		// This topic can be used to forward trades to services other than the prod. tradesblockservice.
		wReplica *kafka.Writer
		wTest    *kafka.Writer
	)

	switch *mode {
	case "current":
		w = kafkaHelper.NewWriter(kafkaHelper.TopicTrades)
		wReplica = kafkaHelper.NewWriter(kafkaHelper.TopicTradesReplica)
		wTest = kafkaHelper.NewWriter(kafkaHelper.TopicTradesTest)
	case "estimation":
		w = kafkaHelper.NewWriter(kafkaHelper.TopicTradesEstimation)
	case "assetmap":
		w = kafkaHelper.NewWriter(kafkaHelper.TopicTradesEstimation)
	}

	defer func() {
		err := w.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	wg := sync.WaitGroup{}

	if scrapers.Exchanges[*exchange].Centralized {

		// Scrape pairs for CEX scrapers.
		for _, configPair := range pairsExchange {
			log.Println("Adding pair:", configPair.Symbol, configPair.ForeignName, "on exchange", exchange)
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
	go handleTrades(es.Channel(), &wg, w, wTest, wReplica, ds, *exchange, *mode)
}

func handleTrades(c chan *dia.Trade, wg *sync.WaitGroup, w *kafka.Writer, wTest *kafka.Writer, wReplica *kafka.Writer, ds *models.DB, exchange string, mode string) {
	lastTradeTime := time.Now()
	watchdogDelay := scrapers.Exchanges[exchange].WatchdogDelay
	if watchdogDelay == 0 {
		watchdogDelay = scrapers.ExchangeDuplicates[exchange].WatchdogDelay
	}
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
			// Trades are sent to the tradesblockservice through a kafka channel - either
			// through trades topic or historical trades topic.
			if mode == "current" || mode == "historical" || mode == "estimation" {

				// Write trade to productive Kafka.
				err := writeTradeToKafka(w, t)
				if err != nil {
					log.Error(err)
				}

				if scrapers.Exchanges[t.Source].Centralized {
					// Write CEX trades to test Kafka.
					if mode == "current" {
						err = writeTradeToKafka(wTest, t)
						if err != nil {
							log.Error(err)
						}
					}
				}

				if replicaKafkaTopic == "true" {
					err := writeTradeToKafka(wReplica, t)
					if err != nil {
						log.Error(err)
					}
				}

			}
			// Trades are just saved in influx - not sent to the tradesblockservice through a kafka channel.
			if mode == "storeTrades" {
				err := ds.SaveTradeInflux(t)
				if err != nil {
					log.Error(err)
				}
			}

			if mode == "assetmap" {
				fmt.Println("recieved trade", t)
			}
		}
	}
}

func writeTradeToKafka(w *kafka.Writer, t *dia.Trade) error {
	// Write trade to Kafka.
	err := kafkaHelper.WriteMessage(w, t)
	if err != nil {
		return err
	}

	// Write reversed trade to Kafka as well for Zenlink.
	tSwapped, err := dia.SwapTrade(*t)
	if err != nil {
		log.Error("swap trade: ", err)
	} else {
		err = kafkaHelper.WriteMessage(w, &tSwapped)
		if err != nil {
			return err
		}
	}

	return nil
}
