package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"strings"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

var (
	log      *logrus.Logger
	exchange string
	exch     = flag.String("exchange", "", "which exchange")
	mode     = flag.String("mode", "verification", "verification or remoteFetch: fetching pairs from exchange's API.")
)

func init() {
	log = logrus.New()
	flag.Parse()

	exchange = *exch
	exchangeStruct, ok := scrapers.Exchanges[exchange]
	if (!exchangeStruct.Centralized || !ok) && *mode == "verification" {
		log.Warnf("%s cannot be found in the list of centralized exchanges.", exchange)
		exchangeStruct, ok := scrapers.ExchangeDuplicates[exchange]
		if (!exchangeStruct.Centralized || !ok) && *mode == "verification" {
			log.Fatalf("%s cannot be found in exchange scraper duplicates either.", exchange)
		}
	}
}

func main() {
	relDB, err := models.NewRelDataStore()

	if err != nil {
		log.Fatal("Unable to initialize relDB: " + err.Error())
	}

	switch *mode {
	case "verification":
		err = updateExchangePairs(relDB)
		if err != nil {
			log.Fatalf("update exchange pairs for %s: %v", exchange, err)
		}
	case "remoteFetch":
		err = fetchFromExchangeAndStore(relDB)

		if err != nil {
			log.Fatalf("update exchange pairs for %s: %v", exchange, err)
		}
	default:
		log.Fatal("unknown mode.")
	}

}

func updateExchangePairs(relDB *models.RelDB) (err error) {

	// --------- Collect pairs from DB and config file ---------
	var pairs []dia.ExchangePair

	// Fetch pairs from postgres.
	pairs, err = relDB.GetExchangePairSymbols(exchange)
	if err != nil {
		return err
	}

	// Add pairs from config file.
	pairs, err = addPairsFromConfig(exchange, pairs)
	if err != nil {
		log.Errorf("adding pairs from config file for exchange %s: %v", exchange, err)
	}

	// --------- Verify all pairs collected above ---------

	// Extract list of symbols from pairs.
	symbols, err := dia.GetAllSymbolsFromPairs(pairs)
	if err != nil {
		log.Error("get symbols from pairs: ", err)
	}

	// Set exchange symbol if not in table yet.
	for _, symbol := range symbols {
		err = relDB.SetExchangeSymbol(exchange, symbol)
		if err != nil {
			log.Errorf("error setting exchange symbol %s: %v", symbol, err)
		}
	}

	// Verify symbols using the file containing verification mappings of CEX symbols.
	gitcoinSymbols, err := readFile("gitcoinverified/" + exchange + ".json")
	if err != nil {
		log.Errorf("Error while reading file %s: %v", "gitcoinverified/"+exchange+".json", err)
	}
	setGitcoinSymbols(gitcoinSymbols, relDB)

	// Verify/falsify exchange pairs using the exchangesymbol table in postgres.
	for _, pair := range pairs {
		var exchangepair dia.ExchangePair
		var pairSymbols []string
		var quotetokenID string
		var basetokenID string
		var quotetokenVerified bool
		var basetokenVerified bool
		log.Info("handle pair ", pair)

		// Continue if pair is already verified.
		exchangepair, err = relDB.GetExchangePairCache(exchange, pair.ForeignName)
		if err != nil {
			log.Errorf("error getting pair %s from cache", pair.ForeignName)
		}
		if exchangepair.Verified {
			log.Infof("pair %s already verified. Continue.", pair.ForeignName)
			continue
		}

		// If not yet verified, try do do so.
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
			return err
		}
	}

	log.Infof("updated exchange %s, closing scraper.", exchange)
	return nil
}

func fetchFromExchangeAndStore(relDB *models.RelDB) error {
	var scraper scrapers.APIScraper
	var pairs []dia.ExchangePair

	// Set up scraper.
	config, err := dia.GetConfig(exchange)
	if err == nil {
		scraper = scrapers.NewAPIScraper(exchange, false, config.ApiKey, config.SecretKey, relDB)
	} else {
		log.Info("No valid API config for exchange: ", exchange, " Error: ", err.Error())
		log.Info("Proceeding with no API secrets")
		scraper = scrapers.NewAPIScraper(exchange, false, "", "", relDB)
	}

	// Fetch pairs from exchange's API.
	pairs, err = scraper.FetchAvailablePairs()
	if err != nil {
		log.Errorf("fetching pairs from API for exchange %s: %v", exchange, err)
		return err
	}

	// Save pairs in json file.
	err = savePairsToFile(exchange, pairs)
	if err != nil {
		log.Error("write pairs to json file: ", err)
		return err
	}
	log.Infof("fetched %v pairs from API of exchange %s and wrote to %s.json in /tmp folder.", len(pairs), exchange, exchange)
	return nil
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

// setGitcoinSymbols writes all mappings from submissions into exchangesymbol table.
func setGitcoinSymbols(submissions GitcoinSubmission, relDB *models.RelDB) {

	errorCount := 0
	for _, submission := range submissions.AllItems {
		// Get ID of underlying asset
		asset, err := relDB.GetAsset(submission.Address, submission.Blockchain)
		if err != nil {
			if strings.Contains(err.Error(), "no rows") {
				errorCount++
				log.Warnf("asset with address %s on %s not in asset table", submission.Address, submission.Blockchain)
				continue
			} else {
				errorCount++
				log.Error("get asset: ", err)
				continue
			}
		}
		assetID, err := relDB.GetAssetID(asset)
		if err != nil {
			errorCount++
			log.Error("get asset ID: ", err)
			continue
		}
		// Write into exchangesymbol table
		success, err := relDB.VerifyExchangeSymbol(submission.Exchange, submission.Symbol, assetID)
		if err != nil || !success {
			errorCount++
			log.Errorf("verify symbol %s on %s: %v", submission.Symbol, submission.Exchange, err)
			continue
		}
	}
	log.Warnf("could not integrate %v out of %v gitcoin verified symbols.", errorCount, len(submissions.AllItems))
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

type LocalPair struct {
	Symbol      string `json:"Symbol"`
	ForeignName string `json:"ForeignName"`
	Exchange    string `json:"Exchange"`
	Ignore      bool   `json:"Ignore"`
}

type LocalExchangePairs struct {
	Coins []LocalPair `json:"Coins"`
}

func savePairsToFile(exchange string, pairs []dia.ExchangePair) error {
	log.Info("savePairsToFile: ", exchange)

	var localPairs []LocalPair

	for _, field := range pairs {
		localPairs = append(localPairs, LocalPair{Exchange: field.Exchange, Symbol: field.Symbol, ForeignName: field.ForeignName, Ignore: false})
	}
	localExchangePairs := LocalExchangePairs{Coins: localPairs}
	b, err := json.Marshal(localExchangePairs)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("/tmp/"+exchange+".json", b, 0644)
}
