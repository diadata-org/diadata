package main

import (
	"flag"
	"fmt"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/diadata-org/diadata/pkg/utils/probes"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

var (
	exchange = flag.String("exchange", "", "which exchange")
	// mode==current:		default mode. Trades are forwarded to TBS and FBS.
	// mode==storeTrades:	trades are not forwarded to TBS and FBS and stored as raw trades in influx.
	// mode==estimation:	trades are forwarded to tradesEstimationService, i.e. same as storeTrades mode
	//						but estimatedUSDPrice is filled by tradesEstimationService.
	// mode==historical:	trades are sent through kafka to TBS in tradesHistorical topic.
	mode                 = flag.String("mode", "current", "either storeTrades, current, historical or estimation")
	swapTradesOnExchange = []string{
		dia.CurveFIExchange,
		dia.CurveFIExchangeFantom,
		dia.CurveFIExchangeMoonbeam,
		dia.CurveFIExchangePolygon,
		dia.OmniDexExchange,
		dia.DiffusionExchange,
		dia.SolarbeamExchange,
		dia.AnyswapExchange,
		dia.HermesExchange,
		dia.HuckleberryExchange,
		dia.NetswapExchange,
	}
	pairsfile   = flag.Bool("pairsfile", false, "read pairs from json file in config folder.")
	StartupDone = false
	// delay until the service will be restarted
	restartDelayMinutes, _ = strconv.Atoi(utils.Getenv("RESTART_DELAY_MINUTES", "720"))
	startTime              = time.Now()
)

// main manages all PairScrapers and handles incoming trade information
func ready() bool {
	return StartupDone
}

func live() bool {
	if !StartupDone {
		return false
	}
	return startTime.Before(startTime.Add(time.Minute * time.Duration(restartDelayMinutes)))
}

func init() {
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

	log.Infof("start collector for %s in %s mode...", *exchange, *mode)

	log.Infof("%s starttime, %s duration", startTime, startTime.Add(time.Minute*time.Duration(restartDelayMinutes)))

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

	exchangeConfig, err := relDB.GetExchange(dia.KuCoinExchange)
	if err != nil {
		log.Warning("no config for exchange ", err)
	}

	diaExchange := dia.Exchange{Name: dia.KuCoinExchange, Centralized: true, WatchdogDelay: exchangeConfig.WatchdogDelay}
	es := scrapers.NewKuCoinScraper(configApi.ApiKey, configApi.SecretKey, diaExchange, true, relDB)

	// Set up kafka writers for various modes.
	var (
		w     *kafka.Writer
		wTest *kafka.Writer
	)

	switch *mode {
	case "current":
		w = kafkaHelper.NewWriter(kafkaHelper.TopicTrades)
		wTest = kafkaHelper.NewWriter(kafkaHelper.TopicTradesTest)
	case "historical":
		w = kafkaHelper.NewWriter(kafkaHelper.TopicTradesHistorical)
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
	go handleTrades(es.Channel(), &wg, w, wTest, ds, *exchange, *mode)

	probes.Start(live, ready)
	StartupDone = true
}

func handleTrades(c chan *dia.Trade, wg *sync.WaitGroup, w *kafka.Writer, wTest *kafka.Writer, ds *models.DB, exchange string, mode string) {
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
			// Trades are sent to the tradesblockservice through a kafka channel - either
			// through trades topic or historical trades topic.
			if mode == "current" || mode == "historical" || mode == "estimation" {

				// Write trade to productive Kafka.
				err := writeTradeToKafka(w, t)
				if err != nil {
					log.Error(err)
				}

				// Write trade to test Kafka.
				if mode == "current" {
					err = writeTradeToKafka(wTest, t)
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
				} else {
					log.Info("saved trade")
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

	// Write reversed trade to Kafka as well for some exchanges.
	if utils.Contains(&swapTradesOnExchange, t.Source) {
		tSwapped, err := dia.SwapTrade(*t)
		if err != nil {
			log.Error("swap trade: ", err)
		} else {
			err = kafkaHelper.WriteMessage(w, &tSwapped)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func isValidExchange(estring string) bool {
	for e := range scrapers.Exchanges {
		if e == estring {
			return true
		}
	}
	return false
}
