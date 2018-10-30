package main

import (
	"flag"
	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	"github.com/diadata-org/diadata/pkg/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
	"os/user"
	"strings"
	"sync"
	"time"
)

const (
	watchdogDelay = 60.0 * 2
)

type doggyBag struct {
	lastTradeTime time.Time
	mutex         *sync.Mutex
}

// pairs contains all pairs currently supported by the DIA scrapers

// handleTrades delegates trade information to Kafka
func handleTrades(d *doggyBag, ps scrapers.PairScraper, wg *sync.WaitGroup, w *kafka.Writer) {
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
		d.mutex.Lock()
		d.lastTradeTime = time.Now()
		d.mutex.Unlock()
		kafkaHelper.WriteMessage(w, t)
	}
}
func getConfig(exchange string) (*dia.ConfigApi, error) {
	var configApi dia.ConfigApi
	usr, _ := user.Current()
	dir := usr.HomeDir
	configFileApi := dir + "/config/secrets/api_" + strings.ToLower(exchange)
	err := gonfig.GetConf(configFileApi, &configApi)
	return &configApi, err
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

func watchDog(d *doggyBag) {

	t := time.NewTicker(watchdogDelay * time.Second)

	for {
		select {
		case <-t.C:
			d.mutex.Lock()
			duration := time.Since(d.lastTradeTime)
			d.mutex.Unlock()
			time.Now()
			if duration.Seconds() > watchdogDelay {
				log.Error(duration)
				panic("frozen? ")
			}
		}
	}
}

// main manages all PairScrapers and handles incoming trade information
func main() {

	d := &doggyBag{
		lastTradeTime: time.Now(),
		mutex:         &sync.Mutex{},
	}

	//	cc := configCollectors.NewConfigCollectors(*exchange)
	ds, err := models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore:", err)
	} else {

	}
	pairsExchange, err := ds.GetAvailablePairsForExchange(*exchange)

	if err != nil {
		log.Error("error on GetAvailablePairsForExchange", err)
		cc := configCollectors.NewConfigCollectors(*exchange)
		pairsExchange = cc.AllPairs()
	}

	configApi, err := dia.GetConfig(*exchange)
	if err != nil {
		log.Warning("error loading configApi\n")
	}
	es := scrapers.NewAPIScraper(*exchange, configApi.ApiKey, configApi.SecretKey)

	w := kafkaHelper.NewWriter(kafkaHelper.TopicTrades)
	defer w.Close()

	wg := sync.WaitGroup{}

	for _, configPair := range pairsExchange {

		log.Println("Adding pair:", configPair.Symbol, configPair.ForeignName, "on exchange", *exchange)

		ps, err := es.ScrapePair(dia.Pair{
			Symbol:      configPair.Symbol,
			ForeignName: configPair.ForeignName})
		if err != nil {
			log.Println(err)
		} else {
			go handleTrades(d, ps, &wg, w)
			wg.Add(1)
		}
		defer wg.Wait()
	}
	go watchDog(d)
}
