package main

import (
	"bufio"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/assetservice/assetstore"
	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	"github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

var (
	log *logrus.Logger
)

const (
	postgresKey = "postgres_key.txt"
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

	cache, err := assetstore.NewRedisDataStore()
	if err != nil {
		panic("Can not initialize cache, error: " + err.Error())
	}
	postgresDB, err := assetstore.NewPostgresDataStore()
	if err != nil {
		panic("Can not initialize postgres DB, error: " + err.Error())
	}

	updateExchangePairs(cache, postgresDB)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	task.wg.Add(1)
	go func() { defer task.wg.Done(); task.run(cache, postgresDB) }()
	select {
	case <-c:
		log.Info("Received stop signal.")
		task.stop()
	}
}

// updateExchangePairs fetches all exchange's trading pairs from postgres and writes them into redis caching layer (toggle == false)
// Periodically, it connects to the exchange's API and checks for new pairs (toggle == true)
func updateExchangePairs(cache, assetDB *assetstore.RelDB) {
	toggle, err := getConfigTogglePairDiscovery()
	if err != nil {
		log.Errorf("updateExchangePairs GetConfigTogglePairDiscovery: %v", err)
		return
	}
	if toggle == false {

		log.Info("GetConfigTogglePairDiscovery = false, using values from config files")
		for _, exchange := range dia.Exchanges() {
			if exchange == "Unknown" {
				continue
			}
			// Fetch all pairs available for @exchange in our asset database
			pairs, err := assetDB.GetAvailablePairs(exchange)
			if err != nil {
				log.Errorf("getting pairs from config for exchange %s: %v", exchange, err)
				continue
			}
			// Set pairs in redis caching layer
			err = cache.CacheSetAvailablePairs(exchange, pairs)
			if err != nil {
				log.Errorf("setting pairs to redis for exchange %s: %v", exchange, err)
				continue
			}
			log.Infof("exchange %s updated\n", exchange)
		}
		log.Info("Update complete.")

	} else {

		log.Info("GetConfigTogglePairDiscovery = true, fetch new pairs from exchange's API")
		for _, exchange := range dia.Exchanges() {

			// Make API Scraper
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

			// If no error, fetch pairs by method implemented for each scraper resp.
			if scraper != nil {
				// Fetch pairs using the API scraper
				pairs, err := scraper.FetchAvailablePairs()
				if err != nil {
					log.Errorf("fetching pairs for exchange %s: %v", exchange, err)
				}
				// Add pairs from assetDB
				pairs, err = addPairsFromAssetDB(exchange, pairs, assetDB)
				if err != nil {
					log.Errorf("adding pairs from asset DB for exchange %s: %v", exchange, err)
				}
				// Optional addition of pairs from config file
				pairs, err = addPairsFromConfig(exchange, pairs)
				if err != nil {
					log.Errorf("adding pairs from config file for exchange %s: %v", exchange, err)
				}
				// Set pairs to assetDB
				for _, pair := range pairs {
					err = assetDB.SetPair(exchange, pair)
					if err != nil {
						log.Errorf("setting exchangepair table for pair on exchange %s: %v", exchange, err)
					}
				}
				// Set pairs to redis caching layer
				err = cache.CacheSetAvailablePairs(exchange, pairs)
				if err != nil {
					log.Errorf("setting caching layer for pair on exchange %s: %v", exchange, err)
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

func getConfigTogglePairDiscovery() (bool, error) {
	// Activates periodically
	return false, nil //TOFIX
}

// addPairsFromAssetDB adds all pairs available for @exchange in our persistent asset database
func addPairsFromAssetDB(exchange string, pairs []dia.Pair, assetDB *assetstore.RelDB) ([]dia.Pair, error) {
	persistentPairs, err := assetDB.GetAvailablePairs(exchange)
	if err != nil {
		return pairs, err
	}
	return dia.MergePairs(pairs, persistentPairs), nil

}

// addPairsFromConfig adds pairs from the config file to @pairs
func addPairsFromConfig(exchange string, pairs []dia.Pair) ([]dia.Pair, error) {
	pairsFromConfig, err := getPairsFromConfig(exchange)
	if err != nil {
		return pairs, err
	}
	return dia.MergePairs(pairs, pairsFromConfig), nil
}

// getPairsFromConfig returns pairs from exchange's config file
func getPairsFromConfig(exchange string) ([]dia.Pair, error) {
	configFileAPI := configCollectors.ConfigFileConnectors(exchange)
	// configFileAPI := "config/" + exchange + ".json"
	var coins Pairs
	err := gonfig.GetConf(configFileAPI, &coins)
	return coins.Coins, err
}

func (t *Task) run(cache *assetstore.RelDB, assetDB *assetstore.RelDB) {
	for {
		select {
		case <-t.closed:
			return
		case <-t.ticker.C:
			updateExchangePairs(cache, assetDB)
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

func getPostgresKeyFromSecrets() string {
	var lines []string
	executionMode := os.Getenv("EXEC_MODE")
	var file *os.File
	var err error
	if executionMode == "production" {
		file, err = os.Open("/run/secrets/" + postgresKey)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file, err = os.Open("../../../secrets/" + postgresKey)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(lines) != 1 {
		log.Fatal("Secrets file should have exactly one line")
	}
	return lines[0]
}
