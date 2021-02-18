package main

import (
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
)

const (
	postgresKey = "postgres_key.txt"
)

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

	relDB, err := models.NewRelDataStore()
	if err != nil {
		panic("Couldn't initialize relDB, error: " + err.Error())
	}

	updateExchangePairs(relDB)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	task.wg.Add(1)
	go func() { defer task.wg.Done(); task.run(relDB) }()
	select {
	case <-c:
		log.Info("Received stop signal.")
		task.stop()
	}
}

// TO DO: Refactor toggle==true case.
// Amend such that we don't overwrite if data is already in DB/cache

// updateExchangePairs fetches all exchange's trading pairs from postgres and writes them into redis caching layer (toggle == false)
// Periodically, it connects to the exchange's API and checks for new pairs (toggle == true)
func updateExchangePairs(relDB *models.RelDB) {
	toggle, err := getConfigTogglePairDiscovery()
	if err != nil {
		log.Errorf("updateExchangePairs GetConfigTogglePairDiscovery: %v", err)
		return
	}
	toggle = true
	if toggle == false {

		log.Info("GetConfigTogglePairDiscovery = false, using values from config files")
		for _, exchange := range dia.Exchanges() {
			if exchange == "Unknown" {
				continue
			}
			// Fetch all pairs available for @exchange from exchangepair table in postgres
			pairs, err := relDB.GetExchangePairs(exchange)
			if err != nil {
				log.Errorf("getting pairs from postgres for exchange %s: %v", exchange, err)
				continue
			}
			// Optional addition of pairs from config file
			pairs, err = addPairsFromConfig(exchange, pairs)
			if err != nil {
				log.Errorf("adding pairs from config file for exchange %s: %v", exchange, err)
			}
			// Set pairs in redis caching layer. The collector will fetch these in order to build
			// the corresponding pair scrapers.
			for _, pair := range pairs {
				err = relDB.SetExchangePairCache(exchange, pair)
				if err != nil {
					log.Errorf("setting pair %s to redis for exchange %s: %v", pair.ForeignName, exchange, err)
					continue
				}
			}
			log.Infof("exchange %s updated\n", exchange)
		}
		log.Info("Update complete.")

	} else {

		log.Info("GetConfigTogglePairDiscovery = true, fetch new pairs from exchange's API")
		for _, exchange := range dia.Exchanges() {

			// Make API Scraper in order to fetch pairs
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

				// --------- 1. Step: collect pairs from Exchange API, DB and config file ---------
				// Fetch pairs using the API scraper
				pairs, err := scraper.FetchAvailablePairs()
				if err != nil {
					log.Errorf("fetching pairs for exchange %s: %v", exchange, err)
				}
				// Add pairs from relDB
				pairs, err = addPairsFromAssetDB(exchange, pairs, relDB)
				if err != nil {
					log.Errorf("adding pairs from asset DB for exchange %s: %v", exchange, err)
				}
				// Optional addition of pairs from config file
				pairs, err = addPairsFromConfig(exchange, pairs)
				if err != nil {
					log.Errorf("adding pairs from config file for exchange %s: %v", exchange, err)
				}

				// --------- 2. Step: Try to verify all pairs collected above ---------

				// First get list of symbols available on exchange and verify/falsify those
				symbols, err := dia.GetAllSymbolsFromPairs(pairs)
				if err != nil {
					log.Error(err)
				}
				verifCount := 0
				for _, symbol := range symbols {
					time.Sleep(1 * time.Second)
					err := relDB.SetExchangeSymbol(exchange, symbol)
					if err != nil {
						log.Errorf("error setting exchange symbol %s", symbol)
						continue
					}
					assetInfo, err := scraper.FetchTickerData(symbol)
					if err != nil {
						log.Errorf("error fetching ticker data for %s", symbol)
						continue
					}
					assetCandidates, err := relDB.IdentifyAsset(assetInfo)
					if err != nil {
						log.Errorf("error getting asset candidates for %s", symbol)
						continue
					}
					if len(assetCandidates) != 1 {
						log.Errorf("could not uniquely identify token ticker %s on exchange %s. Please identify manually.", symbol, exchange)
						continue
					}
					if len(assetCandidates) == 1 {

						verifCount++
						assetID, err := relDB.GetAssetID(assetCandidates[0])
						if err != nil {
							log.Error(err)
						}
						ok, err := relDB.VerifyExchangeSymbol(exchange, symbol, assetID)
						if err != nil {
							log.Error(err)
						}
						if ok {
							log.Infof("verified token ticker %s ", symbol)
						}
					}
				}
				log.Infof("verification of symbols on exchange %s done. Verified %d out of %d symbols.\n", exchange, verifCount, len(symbols))

				// Verify/falsify exchange pairs using the exchangesymbol table in postgres
				for _, pair := range pairs {
					log.Info("handle pair ", pair)
					time.Sleep(1 * time.Second)
					pairSymbols, err := dia.GetPairSymbols(pair)
					if err != nil {
						log.Errorf("error getting symbols from pair string for %s", pair.ForeignName)
						continue
					}
					quotetokenID, quotetokenVerified, err := relDB.GetExchangeSymbolID(exchange, pairSymbols[0])
					basetokenID, basetokenVerified, err := relDB.GetExchangeSymbolID(exchange, pairSymbols[1])

					if quotetokenVerified {
						quotetoken, err := relDB.GetAssetByID(quotetokenID)
						if err != nil {
							log.Error(err)
						}
						pair.UnderlyingPair.QuoteToken = quotetoken
					}
					if basetokenVerified {
						basetoken, err := relDB.GetAssetByID(basetokenID)
						if err != nil {
							log.Error(err)
						}
						pair.UnderlyingPair.BaseToken = basetoken
					}
					if quotetokenVerified && basetokenVerified {
						pair.Verified = true
					}
					// Set pair to postgres
					err = relDB.SetExchangePair(exchange, pair)
					if err != nil {
						log.Errorf("setting exchangepair table for pair on exchange %s: %v", exchange, err)
					}
					// Set pair to redis caching layer
					err = relDB.SetExchangePairCache(exchange, pair)
					if err != nil {
						log.Errorf("setting caching layer for pair on exchange %s: %v", exchange, err)
					}
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

// identifyExchangePair maps an exchange pair to the unique postgres ids of its constituents
func identifyExchangePair(dia.ExchangePair) (assetIDs [2]string, err error) {
	// TO DO
	return
}

func getConfigTogglePairDiscovery() (bool, error) {
	// Activates periodically
	return false, nil //TOFIX
}

// addPairsFromAssetDB adds all pairs available for @exchange in postgres
func addPairsFromAssetDB(exchange string, pairs []dia.ExchangePair, assetDB *models.RelDB) ([]dia.ExchangePair, error) {
	persistentPairs, err := assetDB.GetExchangePairs(exchange)
	if err != nil {
		return pairs, err
	}
	return dia.MergePairs(pairs, persistentPairs), nil

}

// addPairsFromConfig adds pairs from the config file to @pairs
func addPairsFromConfig(exchange string, pairs []dia.ExchangePair) ([]dia.ExchangePair, error) {
	pairsFromConfig, err := getPairsFromConfig(exchange)
	if err != nil {
		return pairs, err
	}
	return dia.MergePairs(pairs, pairsFromConfig), nil
}

// getPairsFromConfig returns pairs from exchange's config file
func getPairsFromConfig(exchange string) ([]dia.ExchangePair, error) {
	configFileAPI := configCollectors.ConfigFileConnectors(exchange)
	type Pairs struct {
		Coins []dia.ExchangePair
	}
	var coins Pairs
	err := gonfig.GetConf(configFileAPI, &coins)
	return coins.Coins, err
}

func (t *Task) run(relDB *models.RelDB) {
	for {
		select {
		case <-t.closed:
			return
		case <-t.ticker.C:
			updateExchangePairs(relDB)
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
