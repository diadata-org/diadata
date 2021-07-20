package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/internal/pkg/assetservice/verifiedTokens"
	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

var (
	log *logrus.Logger
	// updateTime = time.Second * 60 * 60
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
	VerifiedAssetsCount   int                    `json:"verifiedAssets"`
	UnverifiedAssetsCount int                    `json:"unverifiedAssets"`
}

func main() {
	var (
		err           error
		relDB         *models.RelDB
		verifiedToken *verifiedTokens.VerifiedTokens
	)
	task := &Task{
		closed: make(chan struct{}),
		/// Retrieve every 4 hours
		ticker: time.NewTicker(time.Minute * 60 * 4),
	}
	relDB, err = models.NewRelDataStore()
	if err != nil {
		log.Error("setting up relDB: ", err)
		panic("Couldn't initialize relDB, error: " + err.Error())
	}

	// verifiedToken come from a tokenlist: https://uniswap.org/blog/token-lists/
	verifiedToken, err = verifiedTokens.New()
	if err != nil {
		log.Error("Error Getting instance of verified tokens: ", verifiedToken)
	}

	// load gitcoin files
	gitcoinfiles := iterateDirectory("gitcoinverified")
	for _, file := range gitcoinfiles {
		path := "gitcoinverified/" + file
		gitcoinSymbols, err := readFile(path)
		if err != nil {
			log.Errorln("Error while reading  file", path, err)
			continue
		}
		err = setGitcoinSymbols(gitcoinSymbols, relDB)
		if err != nil {
			log.Error("error writing gitcoin submissions: ", err)
		}
	}

	updateExchangePairs(relDB, verifiedToken)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	task.wg.Add(1)
	go func() { defer task.wg.Done(); task.run(relDB, verifiedToken) }()
	for range c {
		log.Info("Received stop signal.")
		task.stop()
	}
}

// TO DO: Refactor toggle==true case.

// toggle == false: fetch all exchange's trading pairs from postgres and write them into redis caching layer
// toggle == true:  connect to all exchange's APIs and check for new pairs
func updateExchangePairs(relDB *models.RelDB, verifiedTokens *verifiedTokens.VerifiedTokens) {
	// TO DO: activate toggle
	// toggle := getTogglePairDiscovery(updateTime)

	toggle, err := strconv.ParseBool(utils.Getenv("PAIR_DISCOV_EXCHANGE_TOGGLE", "true"))
	if err != nil {
		log.Error("Wrong content for PAIR_DISCOV_EXCHANGE_TOGGLE", err)
	}
	if !toggle {

		log.Info("GetConfigTogglePairDiscovery = false, using values from config files")
		for _, exchange := range dia.Exchanges() {

			if exchange == "Unknown" {
				continue
			}
			// Fetch all pairs available for @exchange from exchangepair table in postgres
			simplePairs, err := relDB.GetExchangePairSymbols(exchange)
			if err != nil {
				log.Errorf("getting pairs from postgres for exchange %s: %v", exchange, err)
				continue
			}

			// Get filled version with verification and underlying dia.Pair if existent.
			var pairs []dia.ExchangePair
			for _, pair := range simplePairs {
				var fullPair dia.ExchangePair
				fullPair, err = relDB.GetExchangePair(exchange, pair.ForeignName)
				if err != nil {
					log.Error("error fetching exchangepair: ", err)
					continue
				}
				pairs = append(pairs, fullPair)
			}

			// Optional addition of pairs from config file
			pairs, err = addPairsFromConfig(exchange, pairs)
			if err != nil {
				log.Errorf("adding pairs from config file for exchange %s: %v", exchange, err)
			}
			time.Sleep(20 * time.Second)

			// Set pairs in postgres and redis caching layer. The collector will fetch these
			// from the cache in order to build verified trades.
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
					// TO DO: uncomment below block once all exchange scrapers are deployed.
					// // Fetch pairs using the API scraper.
					// pairs, err := scraper.FetchAvailablePairs()
					// if err != nil {
					// 	log.Errorf("fetching pairs for exchange %s: %v", exchange, err)
					// }
					// log.Infof("fetched %v pairs for exchange %s.", len(pairs), exchange)
					// time.Sleep(60 * time.Second)

					var pairs []dia.ExchangePair
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
						fmt.Printf("verifying symbol %s on exchange %s \n ", symbol, exchange)
						var verified bool
						var assetInfo dia.Asset
						var assetCandidates []dia.Asset
						var assetID string
						var ok bool
						// signature for this part:
						// func matchExchangeSymbol(symbol string, exchange string, relDB *models.RelDB)

						time.Sleep(200 * time.Millisecond)
						// First set all symbols traded on the exchange. These are subsequently
						// matched with assets from the asset table.

						// Continue if symbol is already in DB and verified.
						_, verified, err = relDB.GetExchangeSymbolAssetID(exchange, symbol)
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
						assetInfo, err = scraper.FillSymbolData(symbol)
						if err != nil {
							log.Errorf("error fetching ticker data for %s: %v", symbol, err)
							continue
						}
						// Using the gathered information, find matching assets in asset table.
						assetCandidates, err = relDB.IdentifyAsset(assetInfo)
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
							log.Infof("found asset candidate for %s on %s: %v", symbol, exchange, assetCandidates[0])
							// Verify if this asset is in our verified asset list
							// isVerified := verifiedTokens.IsExists(assetCandidates[0])
							isVerified := true
							if isVerified {
								verificationCount++
								assetID, err = relDB.GetAssetID(assetCandidates[0])
								if err != nil {
									log.Error(err)
								}
								ok, err = relDB.VerifyExchangeSymbol(exchange, symbol, assetID)
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
						var exchangepair dia.ExchangePair
						var pairSymbols []string
						var quotetokenID string
						var basetokenID string
						var quotetokenVerified bool
						var basetokenVerified bool
						log.Info("handle pair ", pair)
						// time.Sleep(1 * time.Second)
						// Continue if pair is already verified
						exchangepair, err = relDB.GetExchangePairCache(exchange, pair.ForeignName)
						if err != nil {
							log.Errorf("error getting pair %s from cache", pair.ForeignName)
						}
						if exchangepair.Verified {
							fmt.Printf("pair %s already verified. Continue.\n", pair.ForeignName)
							continue
						}

						// if not yet verified, try do do so
						pairSymbols, err = dia.GetPairSymbols(pair)
						if err != nil {
							log.Errorf("error getting symbols from pair string for %s", pair.ForeignName)
							continue
						}

						quotetokenID, quotetokenVerified, err = relDB.GetExchangeSymbolAssetID(exchange, pairSymbols[0])
						if err != nil {
							log.Error(err)
						}
						basetokenID, basetokenVerified, err = relDB.GetExchangeSymbolAssetID(exchange, pairSymbols[1])
						if err != nil {
							log.Error(err)
						}

						if quotetokenVerified {
							var quotetoken dia.Asset
							quotetoken, err = relDB.GetAssetByID(quotetokenID)
							if err != nil {
								log.Error(err)
							}
							pair.UnderlyingPair.QuoteToken = quotetoken
						}
						if basetokenVerified {
							var basetoken dia.Asset
							basetoken, err = relDB.GetAssetByID(basetokenID)
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
						Assets:              dataTowrite,
						VerifiedAssetsCount: verificationCount,
					}
					file, _ := json.MarshalIndent(dataforfile, "", " ")
					err = ioutil.WriteFile(exchange+".json", file, 0600)
					if err != nil {
						log.Error(err)
					}
					log.Infof("updated exchange %s", exchange)
					time.Sleep(60 * time.Second)
					go func(s scrapers.APIScraper, exchange string) {
						time.Sleep(5 * time.Second)
						log.Error("Closing scraper: ", exchange)
						err = s.Close()
						if err != nil {
							log.Error(err)
						}
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
						err = s.Close()
						if err != nil {
							log.Error(err)
						}
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

// TO DO: activate toggle
// getTogglePairDiscovery switches to true between midnight and midnight + duration
// func getTogglePairDiscovery(d time.Duration) bool {
// 	t := time.Now()
// 	secondsAfterMidnight := t.Hour()*3600 + t.Minute()*60 + t.Second()
// 	if float64(secondsAfterMidnight) < d.Seconds()+10 {
// 		return true
// 	}
// 	return false
// }

// addNewPairs adds pair from @pairs if it's not in our postgres DB yet.
// Equality refers to the unique identifier (exchange,foreignName).
func addNewPairs(exchange string, pairs []dia.ExchangePair, assetDB *models.RelDB) ([]dia.ExchangePair, error) {
	persistentPairs, err := assetDB.GetExchangePairSymbols(exchange)
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
	configFileAPI := configCollectors.ConfigFileConnectors(exchange, ".json")
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

// ----------------------------------------------------------------------------
// This functionality might be worth to extend and to put into a separate main.
// ----------------------------------------------------------------------------

type GitcoinSubmission struct {
	AllItems []SubmissionItem `json:"Tokens"`
}

type SubmissionItem struct {
	Symbol     string `json:"Symbol"`
	Exchange   string `json:"Exchange"`
	Address    string `json:"Address"`
	Blockchain string `json:"Blockchain"`
}

// MarshalBinary is a custom marshaller for BlockChain type
func (gcs *GitcoinSubmission) MarshalBinary() ([]byte, error) {
	return json.Marshal(gcs)
}

// UnmarshalBinary is a custom unmarshaller for BlockChain type
func (gcs *GitcoinSubmission) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &gcs); err != nil {
		return err
	}
	return nil
}

// readFile reads a gitcoin submission json file and returns the slice of items.
func readFile(filename string) (items GitcoinSubmission, err error) {
	var (
		jsonFile  *os.File
		filebytes []byte
	)
	jsonFile, err = os.Open(configCollectors.ConfigFileConnectors(filename, ""))
	if err != nil {
		return
	}
	defer func() {
		cerr := jsonFile.Close()
		if err == nil {
			err = cerr
		}
	}()

	filebytes, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	err = json.Unmarshal(filebytes, &items)
	if err != nil {
		log.Error(err)
	}
	return
}

// TO DO: At some point we can add a consensus mechanism involving several submissions
// for the same bounty.

// setGitcoinSymbols writes all mappings from submissions into exchangesymbol table.
func setGitcoinSymbols(submissions GitcoinSubmission, relDB *models.RelDB) error {
	for _, submission := range submissions.AllItems {
		// Get ID of underlying asset
		asset, err := relDB.GetAsset(submission.Address, submission.Blockchain)
		if err != nil {
			return err
		}
		assetID, err := relDB.GetAssetID(asset)
		if err != nil {
			return err
		}
		// Write into exchangesymbol table
		success, err := relDB.VerifyExchangeSymbol(submission.Exchange, submission.Symbol, assetID)
		if err != nil || !success {
			return err
		}
	}
	return nil
}

func iterateDirectory(foldername string) (files []string) {
	path := configCollectors.ConfigFileConnectors(foldername, "")
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		fileExtension := filepath.Ext(path)

		if fileExtension == ".json" {
			files = append(files, info.Name())
			fmt.Printf("File Name: %s\n", info.Name())
		}

		return nil
	})
	if err != nil {
		log.Error(err)
	}
	return
}
