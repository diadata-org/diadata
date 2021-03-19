package main

import (
	"encoding/json"
	"fmt"
	"github.com/diadata-org/diadata/cmd/services/assetCollectionService/verifiedTokens"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"time"

	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

var (
	log        *logrus.Logger
	updateTime = time.Second * 60 * 60
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

type ExchangeTicker struct {
	Assets                map[string][]dia.Asset `json:"assets"`
	VerifiedAssetsCount   int                  `json:"verifiedAssets"`
	UnverifiedAssetsCount int                  `json:"unverifiedAssets"`
}

func main() {
	var (
		err           error
		relDB         *models.RelDB
		verifiedToken *verifiedTokens.VerifiedTokens
	)
	task := &Task{
		closed: make(chan struct{}),
		/// Retrieve every hour
		ticker: time.NewTicker(time.Second * 30 * 60),
	}
	relDB, err = models.NewRelDataStore()
	if err != nil {
		log.Error("setting up relDB: ", err)
		panic("Couldn't initialize relDB, error: " + err.Error())
	}

	verifiedToken, err = verifiedTokens.New()
	if err != nil {
		log.Error("Error Getting instance of verified tokens: ", verifiedToken)
	}

	updateExchangePairs(relDB, verifiedToken)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	task.wg.Add(1)
	go func() { defer task.wg.Done(); task.run(relDB, verifiedToken) }()
	select {
	case <-c:
		log.Info("Received stop signal.")
		task.stop()
	}
}

// TO DO: Refactor toggle==true case.

// toggle == false: fetch all exchange's trading pairs from postgres and write them into redis caching layer
// toggle == true:  connect to all exchange's APIs and check for new pairs
func updateExchangePairs(relDB *models.RelDB, verifiedTokens *verifiedTokens.VerifiedTokens) {
	toggle := getTogglePairDiscovery(updateTime)

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
			// Set pairs in postgres and redis caching layer. The collector will fetch these
			// in order to build the corresponding pair scrapers.
			for _, pair := range pairs {
				err = relDB.SetExchangePair(exchange, pair, true)
				if err != nil {
					log.Errorf("setting exchangepair table for pair on exchange %s: %v", exchange, err)
				}
			}
			log.Infof("exchange %s updated\n", exchange)
		}
		log.Info("Update complete.")

	} else {

		log.Info("GetConfigTogglePairDiscovery = true, fetch new pairs from exchange's API")
		exchangeMap := scrapers.Exchanges
		for _, exchange := range dia.Exchanges() {
			dataTowrite := make(map[string][]dia.Asset)
		 

			// TO DO: the next cond is only for testing. Remove when deploying.
			//if exchange == "Uniswap" {
			//	continue
			//}

			// Make exchange API Scraper in order to fetch pairs
			log.Info("Updating exchange ", exchange)
			var scraper scrapers.APIScraper
			config, err := dia.GetConfig(exchange)
			if err == nil {
				scraper = scrapers.NewAPIScraper(exchange, false, config.ApiKey, config.SecretKey, relDB)
			} else {
				log.Info("No valid API config for exchange: ", exchange, " Error: ", err.Error())
				log.Info("Proceeding with no API secrets")
				scraper = scrapers.NewAPIScraper(exchange, false, "", "", relDB)
			}

			// If no error, fetch pairs by method implemented for each scraper resp.
			if scraper != nil {
				if exchangeMap[exchange].Centralized {

					// --------- 1. Step: collect pairs from Exchange API, DB and config file ---------
					// Fetch pairs using the API scraper.
					pairs, err := scraper.FetchAvailablePairs()
					if err != nil {
						log.Errorf("fetching pairs for exchange %s: %v", exchange, err)
					}
					// If not in postgres yet, add fetched pairs to postgres pairs
					pairs, err = addNewPairs(exchange, pairs, relDB)
					if err != nil {
						log.Errorf("adding pairs from asset DB for exchange %s: %v", exchange, err)
					}
					// Optional addition of pairs from config file.
					pairs, err = addPairsFromConfig(exchange, pairs)
					if err != nil {
						log.Errorf("adding pairs from config file for exchange %s: %v", exchange, err)
					}

					// --------- 2. Step: Try to verify all pairs collected above ---------

					// 2.a Get list of symbols available on exchange and try to match to assets.

					symbols, err := dia.GetAllSymbolsFromPairs(pairs)
					if err != nil {
						log.Error(err)
					}
					verificationCount := 0
					for _, symbol := range symbols {
						// signature for this part:
						// func matchExchangeSymbol(symbol string, exchange string, relDB *models.RelDB)

						time.Sleep(200 * time.Millisecond)
						// First set all symbols traded on the exchange. These are subsequently
						// matched with assets from the asset table.

						// Continue if symbol is already in DB and verified.
						_, verified, err := relDB.GetExchangeSymbolAssetID(exchange, symbol)
						if err != nil {
							if err.Error() != pgx.ErrNoRows.Error() {
								log.Errorf("error getting exchange symbol %s: %v", symbol, err)
							}
						}
						if verified {
							verificationCount++
							continue
						}
						// Set exchange symbol if not in table yet.
						err = relDB.SetExchangeSymbol(exchange, symbol)
						if err != nil {
							log.Errorf("error setting exchange symbol %s: %v", symbol, err)
						}
						// Gather as much information on @symbol as available on the exchange's API.
						assetInfo, err := scraper.FillSymbolData(symbol)
						if err != nil {
							log.Errorf("error fetching ticker data for %s: %v", symbol, err)
							continue
						}
						// Using the gathered information, find matching assets in asset table.
						assetCandidates, err := relDB.IdentifyAsset(assetInfo)
						if err != nil {

							log.Errorf("error getting asset candidates for %s: %v", symbol, err)
							continue
						}
						if len(assetCandidates) != 1 {
							if dataTowrite[symbol] == nil {
								dataTowrite[symbol] = []dia.Asset{}

							}
							dataTowrite[symbol] = append(dataTowrite[symbol], assetCandidates...)
							log.Errorf("could not uniquely identify token ticker %s on exchange %s. Please identify manually.", symbol, exchange)
							continue
						}
						// In case of a unique match, verify symbol in postgres and
						// assign it the corresponding foreign key from the asset table.
						if len(assetCandidates) == 1 {
							// Verify if this asset is in our verified asset list
							isVerified := verifiedTokens.IsExists(assetCandidates[0])
							if isVerified {
								verificationCount++
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

						} else {
							dataTowrite[symbol] = append(dataTowrite[symbol], assetCandidates...)
							log.Errorf("could not verify identify token ticker %s on verified List %s. Please identify manually.", symbol, exchange)
						}
					}
					log.Infof("verification of symbols on exchange %s done. Verified %d out of %d symbols.\n", exchange, verificationCount, len(symbols))

					// 2.b Verify/falsify exchange pairs using the exchangesymbol table in postgres.
					for _, pair := range pairs {
						log.Info("handle pair ", pair)
						// time.Sleep(1 * time.Second)
						// Continue if pair is already verified
						exchangepair, err := relDB.GetExchangePairCache(exchange, pair.ForeignName)
						if err != nil {
							log.Errorf("error getting pair %s from cache", pair.ForeignName)
						}
						if exchangepair.Verified {
							fmt.Println("continue for ", pair.ForeignName)
							continue
						}
						// if not yet verified, try do do so
						pairSymbols, err := dia.GetPairSymbols(pair)
						if err != nil {
							log.Errorf("error getting symbols from pair string for %s", pair.ForeignName)
							continue
						}
						quotetokenID, quotetokenVerified, err := relDB.GetExchangeSymbolAssetID(exchange, pairSymbols[0])
						basetokenID, basetokenVerified, err := relDB.GetExchangeSymbolAssetID(exchange, pairSymbols[1])

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
						// Set pair to postgres and redis cache.
						err = relDB.SetExchangePair(exchange, pair, true)
						if err != nil {
							log.Errorf("setting exchangepair table for pair on exchange %s: %v", exchange, err)
						}
					}

					dataforfile := &ExchangeTicker{
						Assets: dataTowrite,
						VerifiedAssetsCount: verificationCount,
					}
					file, _ := json.MarshalIndent(dataforfile, "", " ")
					ioutil.WriteFile(exchange+".json", file, 0644)
					log.Infof("updated exchange %s", exchange)
					time.Sleep(60 * time.Second)
					go func(s scrapers.APIScraper, exchange string) {
						time.Sleep(5 * time.Second)
						log.Error("Closing scraper: ", exchange)
						scraper.Close()
					}(scraper, exchange)
				} else {
					// For DEXes, FetchAvailablePairs can retrieve unique information.
					// Pairs must contain base- and quotetoken addresses and blockchains.
					pairs, err := scraper.FetchAvailablePairs()
					if err != nil {
						log.Errorf("fetching pairs for exchange %s: %v", exchange, err)
					}
					for _, pair := range pairs {
						// Set pair to postgres and redis cache
						err = relDB.SetExchangePair(exchange, pair, true)
						if err != nil {
							log.Errorf("setting exchangepair table for pair on exchange %s: %v", exchange, err)
						}
					}
					// For the sake of completeness/statistics we could also write the symbols into exchangesymbol table.
					// symbols, err := dia.GetAllSymbolsFromPairs(pairs)
					// if err != nil {
					// 	log.Error(err)
					// }

					go func(s scrapers.APIScraper, exchange string) {
						time.Sleep(5 * time.Second)
						log.Error("Closing scraper: ", exchange)
						scraper.Close()
					}(scraper, exchange)
				}
			} else {
				log.Error("Error creating APIScraper for exchange: ", exchange)
			}
			// Write exchange file

		}
		log.Info("Update complete.")

	}
}

// getTogglePairDiscovery switches to true between midnight and midnight + duration
func getTogglePairDiscovery(d time.Duration) bool {
	t := time.Now()
	secondsAfterMidnight := t.Hour()*3600 + t.Minute()*60 + t.Second()
	if float64(secondsAfterMidnight) < d.Seconds()+10 {
		return true
	}
	return false
}

// addNewPairsToPG adds pair from @pairs if it's not in our postgres DB yet.
// Equality refers to the unique identifier (exchange,foreignName).
func addNewPairs(exchange string, pairs []dia.ExchangePair, assetDB *models.RelDB) ([]dia.ExchangePair, error) {
	persistentPairs, err := assetDB.GetExchangePairs(exchange)
	if err != nil {
		return pairs, err
	}
	// The order counts here. persistentPairs have priority.
	return dia.MergeExchangePairs(persistentPairs, pairs), nil
}

// addPairsFromConfig adds pairs from the config file to @pairs, if not in there yet.
// Equality refers to the unique identifier (exchange,foreignName).
func addPairsFromConfig(exchange string, pairs []dia.ExchangePair) ([]dia.ExchangePair, error) {
	pairsFromConfig, err := getPairsFromConfig(exchange)
	if err != nil {
		return pairs, err
	}
	return dia.MergeExchangePairs(pairs, pairsFromConfig), nil
}

// getPairsFromConfig returns pairs from exchange's config file.
func getPairsFromConfig(exchange string) ([]dia.ExchangePair, error) {
	configFileAPI := configCollectors.ConfigFileConnectors(exchange)
	type Pairs struct {
		Coins []dia.ExchangePair
	}
	var coins Pairs
	err := gonfig.GetConf(configFileAPI, &coins)
	return coins.Coins, err
}

func (t *Task) run(relDB *models.RelDB, verifiedTokens *verifiedTokens.VerifiedTokens) {
	for {
		select {
		case <-t.closed:
			return
		case <-t.ticker.C:
			updateExchangePairs(relDB, verifiedTokens)
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
