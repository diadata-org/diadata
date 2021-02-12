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

type Task struct {
	closed chan struct{}
	wg     sync.WaitGroup
	ticker *time.Ticker
}

func init() {
	log = logrus.New()
}

func main() {
	task := &Task{
		closed: make(chan struct{}),
		/// Retrieve every hour
		ticker: time.NewTicker(time.Second * 60 * 60),
	}
	var err error
	db, err = models.NewDataStore()
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
			log.Info("Updating exchange ", exchange)
			var scraper scrapers.APIScraper
			config, err := dia.GetConfig(exchange)
			if err == nil {
				scraper = scrapers.NewAPIScraper(exchange, config.ApiKey, config.SecretKey)
			} else {
				log.Info("No valid API config for exchange: ", exchange, " Error: ", err.Error())
				log.Info("Proceeding with no API secrets")
				scraper = scrapers.NewAPIScraper(exchange, "", "")
			}
			if scraper != nil {
				// Fetch pairs by method implemented for each scraper resp.
				pairs, err := scraper.FetchAvailablePairs()
				if err == nil {
					pairs = addPairsFromConfig(exchange, pairs)
					err := db.SetAvailablePairsForExchange(exchange, pairs)
					if err == nil {
						log.Info("Exchange: ", exchange, " updated")
					} else {
						log.Error("Error adding pairs to redis for exchange: ", exchange, " error: ", err.Error())
					}
				} else {
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

// getInitialExchangePairs fetches trading pairs from each exchange's config file
// and saves them into redis.
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
			// save pairs in redis
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

func getPairsFromConfig(exchange string) ([]dia.Pair, error) {
	configFileAPI := configCollectors.ConfigFileConnectors(exchange)
	// configFileAPI := "config/" + exchange + ".json"
	var coins Pairs
	err := gonfig.GetConf(configFileAPI, &coins)
	return coins.Coins, err
}

func savePairsToFile(exchange string, pairs []dia.Pair) {
	log.Info("savePairsToFile: ", exchange)
	b, err := json.Marshal(&Pairs{pairs})
	if err != nil {
		log.Error("error while saving pairs to file", err)
	}
	err = ioutil.WriteFile("/tmp/"+exchange+".json", b, 0644)
}

// addPairsFromConfig adds pairs from the config file to @remotePairs
func addPairsFromConfig(exchange string, remotePairs []models.Pair) []dia.Pair {
	pairsFromConfig, _ := getPairsFromConfig(exchange)
	log.Info(exchange, " num remote: ", len(remotePairs), ", num pLocales: ", len(pairsFromConfig))
	for _, pair := range remotePairs {
		if ok := models.ContainsPair(remotePairs, pair); !ok {
			remotePairs = append(remotePairs, pair)
		}
	}
	savePairsToFile(exchange, remotePairs)
	return remotePairs
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
