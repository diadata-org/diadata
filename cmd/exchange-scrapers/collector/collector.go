package main

import (
	"flag"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var (
	log                  *logrus.Logger
	swapTradesOnExchange = []string{
		dia.CurveFIExchange,
		dia.OmniDexExchange,
		dia.DiffusionExchange,
		dia.BalancerV2Exchange,
		dia.BeetsExchange,
		dia.PanCakeSwap,
		dia.SolarbeamExchange,
		dia.AnyswapExchange,
		dia.HermesExchange,
		dia.HuckleberryExchange,
		dia.NetswapExchange,
		dia.PangolinExchange,
		dia.ArthswapExchange,
	}
)

func init() {
	log = logrus.New()
}

func handleTrades(c chan *dia.Trade, wg *sync.WaitGroup, w *kafka.Writer, ds *models.DB, exchange string, mode string) {
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
			// Trades are sent to the tradesblockservice through a kafka channel - either through trades topic
			// or historical trades topic.
			if mode == "current" || mode == "historical" || mode == "estimation" {

				// Write trade to Kafka.
				err := kafkaHelper.WriteMessage(w, t)
				if err != nil {
					log.Error(err)
				}

				// Write reversed trade to Kafka as well for some exchanges.
				if utils.Contains(&swapTradesOnExchange, t.Source) {
					tSwapped, err := dia.SwapTrade(*t)
					if err != nil {
						log.Error("swap trade: ", err)
					} else {
						err := kafkaHelper.WriteMessage(w, &tSwapped)
						if err != nil {
							log.Error(err)
						}
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
		}
	}
}

var (
	exchange         = flag.String("exchange", "", "which exchange")
	onePairPerSymbol = flag.Bool("onePairPerSymbol", false, "one Pair max Per Symbol ?")
	// mode==current:		default mode. Trades are forwarded to TBS and FBS.
	// mode==storeTrades:	trades are not forwarded to TBS and FBS and stored as raw trades in influx.
	// mode==estimation:	trades are forwarded to tradesEstimationService, i.e. same as storeTrades mode
	//						but estimatedUSDPrice is filled by tradesEstimationService.
	// mode==historical:	trades are sent through kafka to TBS in tradesHistorical topic.
	mode = flag.String("mode", "current", "either storeTrades, current, historical or estimation")
)

func isValidExchange(estring string) bool {
	for e := range scrapers.Exchanges {
		if e == estring {
			return true
		}
	}
	return false
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
		log.Fatal("exchange is required")
	}
	if !isValidExchange(*exchange) {
		log.Fatal("Invalid exchange string: ", *exchange)
	}

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

	pairsExchange, err := relDB.GetExchangePairSymbols(*exchange)
	log.Info("available exchangePairs:", len(pairsExchange))

	if err != nil || len(pairsExchange) == 0 {
		log.Error("error on GetExchangePairSymbols", err)
		cc := configCollectors.NewConfigCollectors(*exchange, ".json")
		pairsExchange = cc.AllPairs()
	}

	configApi, err := dia.GetConfig(*exchange)
	if err != nil {
		log.Warning("no config for exchange's api ", err)
	}
	es := scrapers.NewAPIScraper(*exchange, true, configApi.ApiKey, configApi.SecretKey, relDB)

	var w *kafka.Writer
	switch *mode {
	case "current":
		w = kafkaHelper.NewWriter(kafkaHelper.TopicTrades)
	case "historical":
		w = kafkaHelper.NewWriter(kafkaHelper.TopicTradesHistorical)
	case "estimation":
		w = kafkaHelper.NewWriter(kafkaHelper.TopicTradesEstimation)
	}

	defer func() {
		err := w.Close()
		if err != nil {
			log.Error(err)
		}
	}()

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
	go handleTrades(es.Channel(), &wg, w, ds, *exchange, *mode)
}
