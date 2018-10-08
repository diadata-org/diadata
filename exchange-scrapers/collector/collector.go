package main

import (
	"flag"
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/dia/helpers/configCollectors"
	"github.com/diadata-org/api-golang/dia/helpers/kafkaHelper"
	"github.com/diadata-org/api-golang/exchange-scrapers/scrapers"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
	"os/user"
	"strings"
	"sync"
)

// pairs contains all pairs currently supported by the DIA scrapers

// handleTrades delegates trade information to Kafka
func handleTrades(ps scrapers.PairScraper, wg *sync.WaitGroup, w *kafka.Writer) {
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

// main manages all PairScrapers and handles incoming trade information
func main() {

	cc := configCollectors.NewConfigCollectors()

	configApi, err := dia.GetConfig(*exchange)
	if err != nil {
		log.Warning("error loading configApi\n")
	}
	es := scrapers.NewAPIScraper(*exchange, configApi.ApiKey, configApi.SecretKey)

	w := kafkaHelper.NewWriter(kafkaHelper.TopicTrades)
	defer w.Close()

	wg := sync.WaitGroup{}

	for _, configPair := range cc.AllPairsForExchange(*exchange) {

		log.Println("Adding pair:", configPair.Symbol, " on exchange ", *exchange)

		ps, err := es.ScrapePair(dia.Pair{
			Symbol:      configPair.Symbol,
			ForeignName: configPair.ForeignName})
		if err != nil {
			log.Println(err)
		} else {
			go handleTrades(ps, &wg, w)
			wg.Add(1)
		}
		defer wg.Wait()
	}
}
