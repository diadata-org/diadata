package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"time"

	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

var (
	log *logrus.Logger
	db  models.Datastore
)

type Pairs struct {
	Coins []dia.Pair
}

func getPairsFromConfig(exchange string) ([]dia.Pair, error) {
	configFileAPI := configCollectors.ConfigFileConnectors(exchange, ".json")
	// configFileAPI := "config/" + exchange + ".json"
	var coins Pairs
	err := gonfig.GetConf(configFileAPI, &coins)
	return coins.Coins, err
}

type Task struct {
	closed chan struct{}
	wg     sync.WaitGroup
	ticker *time.Ticker
}

func (t *Task) run() {
	for {
		select {
		case <-t.closed:
			return
		case <-t.ticker.C:
			updateExchangePairs()
		}
	}
}

func (t *Task) stop() {
	log.Println("Stoping exchange pair update thread...")
	close(t.closed)
	t.wg.Wait()
	log.Println("Thread stopped, cleaning...")
	// Clean if required
	log.Println("Done")
}

func savePairsToFile(exchange string, pairs []dia.Pair) {
	log.Info("savePairsToFile: ", exchange)
	b, err := json.Marshal(&Pairs{pairs})
	if err != nil {
		log.Error("error while saving pairs to file", err)
	}
	err = ioutil.WriteFile("/tmp/" + exchange + ".json", b, 0644)
}

func updateExchangePairs() {
	toggle, err := db.GetConfigTogglePairDiscovery()
	if err != nil {
		log.Error("updateExchangePairs GetConfigTogglePairDiscovery: ", err.Error())
		return
	}
	if toggle == false {
		log.Info("GetConfigTogglePairDiscovery = false, using default values")
		getInitialExchangePairs()
	} else {
		for _, exchange := range dia.Exchanges() {
			if exchange == "CoinBase" || exchange == "Huobi" || exchange == "Unknown" {
				continue
			}
			log.Info("Updating exchange ", exchange)
			var scraper scrapers.APIScraper
			config, err := dia.GetConfig(exchange)
			if err == nil { //TODO: APIs with no need for a key
				scraper = scrapers.NewAPIScraper(exchange, config.ApiKey, config.SecretKey)
			} else {
				log.Info("No valid API config for exchange: ", exchange, " Error: ", err.Error())
				log.Info("Proceeding with no API secrets")
				scraper = scrapers.NewAPIScraper(exchange, "", "")
			}
			if scraper != nil {
				pairs, err := scraper.FetchAvailablePairs()
				if err == nil {
					addLocalPairs(exchange, pairs)
					err := db.SetAvailablePairsForExchange(exchange, pairs)
					if err == nil {
						log.Info("Exchange: ", exchange, " updated")
					} else {
						log.Error("Error adding pairs to redis for exchange: ", exchange, " error: ", err.Error())
					}
				} else {
					//	log.Info("locale ", err.Error())
					log.Error("Error fetching pairs for exchange: ", exchange, " error: ", err.Error())
				}
				go func(s scrapers.APIScraper, exchange string) {
					time.Sleep(5 * time.Second)
					log.Error("Closing scraper: ", exchange)
					scraper.Close()
				}(scraper, exchange)
			} else {
				log.Error("Error creating APIScraper for exchange: ", exchange)
			}
		}
		log.Info("Update complete.")
	}
}

func addLocalPairs(exchange string, remotePairs []dia.Pair) []dia.Pair {
	localPairs, _ := getPairsFromConfig(exchange)
	log.Info(exchange, " num remote: ", len(remotePairs), ", num pLocales: ", len(localPairs))
	for i := range remotePairs {
		remotePairs[i].Ignore = true
	}
	for i, remotePair := range remotePairs {
		for _, localPair := range localPairs {
			if localPair.Exchange == remotePair.Exchange && localPair.Symbol == remotePair.Symbol && remotePair.ForeignName == localPair.ForeignName {
				remotePairs[i].Ignore = false
			}
		}
	}
	savePairsToFile(exchange, remotePairs)
	return remotePairs
}

func getInitialExchangePairs() {
	log.Info("Loading pairs from config...")
	for _, e := range dia.Exchanges() {
		if e == "Unknown" {
			continue
		}
		p, err := getPairsFromConfig(e)
		if err == nil {
			pairsToSave := []dia.Pair{}
			for _, pp := range p {
				if !pp.Ignore {
					pairsToSave = append(pairsToSave, pp)
				} else {
					log.Debug("ignoring", pp)
				}
			}
			// savePairsToFile(e, p)
			err := db.SetAvailablePairsForExchange(e, pairsToSave)
			if err == nil {
				log.Info("Exchange: ", e, " set")
			} else {
				log.Error("Error setting pairs for exchange:", e, " error:", err.Error())
			}
		} else {
			log.Error("Error processing config for exchange:", e, " error:", err.Error())
		}
	}
	log.Info("Update complete.")
}

func main() {
	task := &Task{
		closed: make(chan struct{}),
		/// Retrieve every hour
		ticker: time.NewTicker(time.Second * 60 * 60),
	}
	var err error
  db , err = models.NewDataStore()
	if err != nil {
		panic("Can not initialize db, error: " + err.Error())
	}
	updateExchangePairs()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	task.wg.Add(1)
	go func() { defer task.wg.Done(); task.run() }()
	select {
	case <-c:
		log.Info("Received stop signal.")
		task.stop()
	}
}
func init() {
	log = logrus.New()
}
