package main

import (
	"strconv"
	"strings"
	"time"

	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

const (
	influxDbOldTradesTable = "trades"
	influxDbTestTable      = "tradestest"
	querySizeHours         = 12
)

var log *logrus.Logger

func init() {
	log = logrus.New()
}

type VerifiableAsset struct {
	Asset    dia.Asset
	Verified bool
}

func main() {

	ReadURL := utils.Getenv("INFLUX_READ_URL", "")
	WriteURL := utils.Getenv("INFLUX_WRITE_URL", "")
	log.Info("readURL: ", ReadURL)
	log.Info("writeURL: ", WriteURL)
	fromTable := utils.Getenv("INFLUX_TABLE_ORIGIN", influxDbOldTradesTable)
	toTable := utils.Getenv("INFLUX_TABLE_DESTINATION", influxDbTestTable)
	log.Info("fromTable: ", fromTable)
	log.Info("toTable: ", toTable)

	// timeFinalString is the last timestamp for which trades are read from influx in Unix time seconds.
	timeFinalString := utils.Getenv("TIME_FINAL", "1636618800")
	timeFinalInt, err := strconv.ParseInt(timeFinalString, 10, 64)
	if err != nil {
		log.Error("parse timeFinal: ", err)
	}
	timeFinal := time.Unix(timeFinalInt, 0)
	log.Info("timeFinal: ", timeFinal)
	// The oldest record we have in influx is younger than 3000 days old.
	timeInit := time.Now().AddDate(0, 0, -3000)

	testmodestring := utils.Getenv("INFLUX_MIGRATION_TESTMODE", "true")
	testmode, err := strconv.ParseBool(testmodestring)
	if err != nil {
		log.Error("parse testmode: ", err)
	}

	dsRead, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore: ", err)
	}
	dsRead.SetInfluxClient(ReadURL)

	dsWrite, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore: ", err)
	}
	dsWrite.SetInfluxClient(WriteURL)

	rdb, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("relational datastore: ", err)
	}

	log.Info("load and assign exchangesymbols... ")
	allSymbolsMap, err := loadAllExchangeSymbols(rdb)
	if err != nil {
		log.Fatal("load exchange symbols' assets: ", err)
	}
	log.Info("...exchangesymbols loaded and assigned.")
	time.Sleep(2 * time.Second)

	allExchanges := dia.Exchanges()
	for i := 0; i < len(allExchanges); i++ {
		exchange := allExchanges[i]
		if !scrapers.Exchanges[exchange].Centralized {
			continue
		}
		if exchange == "Unknown" || exchange == "FinageForex" {
			continue
		}

		func(
			exch string,
			startTime time.Time,
			timeFinal time.Time,
			querySizeHours time.Duration,
			allSymbolsMap map[string]VerifiableAsset,
			fromTable string,
			toTable string,
			dsRead *models.DB,
			dsWrite *models.DB,
		) {

			log.Infof("migrate trades for %s ...", exch)
			log.Info("write into: ", toTable)

			var endTime time.Time
			for startTime.Before(timeFinal) {
				// If not beyond time range, query @querySizeHours hours at a time...
				if (timeInit.Add(querySizeHours)).Before(timeFinal) {
					endTime = startTime.Add(querySizeHours)
				} else {
					// ...else, query until last time of time range.
					endTime = timeFinal
				}
				err = migrateCEXTrades(exch, allSymbolsMap, startTime, endTime, fromTable, toTable, testmode, dsRead, dsWrite)
				if err != nil {
					log.Errorf("migrate trades for exchange %s: %v", exch, err)
				}
				log.Infof("migrated trades between %v and %v", startTime, endTime)
				startTime = startTime.Add(querySizeHours)
			}
		}(exchange, timeInit, timeFinal, querySizeHours*time.Hour, allSymbolsMap, fromTable, toTable, dsRead, dsWrite)
	}

}

// func main() {

// 	ReadURL := utils.Getenv("INFLUX_READ_URL", "")
// 	WriteURL := utils.Getenv("INFLUX_WRITE_URL", "")
// 	log.Info("readURL: ", ReadURL)
// 	log.Info("writeURL: ", WriteURL)
// 	fromTable := utils.Getenv("INFLUX_TABLE_ORIGIN", influxDbOldTradesTable)
// 	toTable := utils.Getenv("INFLUX_TABLE_DESTINATION", influxDbTestTable)
// 	log.Info("fromTable: ", fromTable)
// 	log.Info("toTable: ", toTable)

// 	// timeFinalString is the last timestamp for which trades are read from influx in Unix time seconds.
// 	timeFinalString := utils.Getenv("TIME_FINAL", "1636618800")
// 	timeFinalInt, err := strconv.ParseInt(timeFinalString, 10, 64)
// 	if err != nil {
// 		log.Error("parse timeFinal: ", err)
// 	}
// 	timeFinal := time.Unix(timeFinalInt, 0)
// 	log.Info("timeFinal: ", timeFinal)
// 	// The oldest record we have in influx is younger than 3000 days old.
// 	timeInit := time.Now().AddDate(0, 0, -3000)

// 	testmodestring := utils.Getenv("INFLUX_MIGRATION_TESTMODE", "true")
// 	testmode, err := strconv.ParseBool(testmodestring)
// 	if err != nil {
// 		log.Error("parse testmode: ", err)
// 	}

// 	dsRead, err := models.NewDataStore()
// 	if err != nil {
// 		log.Fatal("datastore: ", err)
// 	}
// 	dsRead.SetInfluxClient(ReadURL)

// 	dsWrite, err := models.NewDataStore()
// 	if err != nil {
// 		log.Fatal("datastore: ", err)
// 	}
// 	dsWrite.SetInfluxClient(WriteURL)

// 	rdb, err := models.NewRelDataStore()
// 	if err != nil {
// 		log.Fatal("relational datastore: ", err)
// 	}

// 	log.Info("load and assign exchangesymbols... ")
// 	allSymbolsMap, err := loadAllExchangeSymbols(rdb)
// 	if err != nil {
// 		log.Fatal("load exchange symbols' assets: ", err)
// 	}
// 	log.Info("...exchangesymbols loaded and assigned.")
// 	time.Sleep(2 * time.Second)

// 	var wg sync.WaitGroup
// 	allExchanges := dia.Exchanges()
// 	for i := 0; i < len(allExchanges); i++ {
// 		exchange := allExchanges[i]
// 		if !scrapers.Exchanges[exchange].Centralized {
// 			continue
// 		}
// 		if exchange == "Unknown" {
// 			continue
// 		}

// 		wg.Add(1)
// 		go func(
// 			exch string,
// 			startTime time.Time,
// 			timeFinal time.Time,
// 			querySizeHours time.Duration,
// 			allSymbolsMap map[string]VerifiableAsset,
// 			fromTable string,
// 			toTable string,
// 			dsRead *models.DB,
// 			dsWrite *models.DB,
// 			w *sync.WaitGroup,
// 		) {
// 			defer w.Done()
// 			log.Infof("migrate trades for %s ...", exch)
// 			log.Info("write into: ", toTable)

// 			var endTime time.Time
// 			for startTime.Before(timeFinal) {
// 				// If not beyond time range, query @querySizeHours hours at a time...
// 				if (timeInit.Add(querySizeHours)).Before(timeFinal) {
// 					endTime = startTime.Add(querySizeHours)
// 				} else {
// 					// ...else, query until last time of time range.
// 					endTime = timeFinal
// 				}
// 				err = migrateCEXTrades(exch, allSymbolsMap, startTime, endTime, fromTable, toTable, testmode, dsRead, dsWrite)
// 				if err != nil {
// 					log.Errorf("migrate trades for exchange %s: %v", exch, err)
// 				}
// 				log.Infof("migrated trades between %v and %v", startTime, endTime)
// 				startTime = startTime.Add(querySizeHours)
// 			}
// 		}(exchange, timeInit, timeFinal, querySizeHours*time.Hour, allSymbolsMap, fromTable, toTable, dsRead, dsWrite, &wg)
// 	}
// 	wg.Wait()

// }

// loadAllExchangeSymbols returns a map that maps the string exchange-symbol to the
// underlying verifiable asset.
func loadAllExchangeSymbols(rdb *models.RelDB) (map[string]VerifiableAsset, error) {
	allSymbolsMap := make(map[string]VerifiableAsset)
	for _, exchange := range dia.Exchanges() {
		if !scrapers.Exchanges[exchange].Centralized {
			continue
		}
		exchangesymbols, err := loadExchangeSymbols(exchange, rdb)
		if err != nil {
			return make(map[string]VerifiableAsset), err
		}
		for symbol, verifAsset := range exchangesymbols {
			allSymbolsMap[exchange+"-"+symbol] = verifAsset
		}
	}
	return allSymbolsMap, nil
}

// loadExchangeSymbols loads all symbols traded on @exchange and assigns the underlying asset
// together with a boolean that is true if the asset is verified on @exchange.
func loadExchangeSymbols(exchange string, rdb *models.RelDB) (map[string]VerifiableAsset, error) {
	exchangesymbolMap := make(map[string]VerifiableAsset)
	exchangesymbols, err := rdb.GetExchangeSymbols(exchange,"")
	if err != nil {
		return make(map[string]VerifiableAsset), err
	}
	for _, symbol := range exchangesymbols {
		asset, assetVerified, err := getExchangeAsset(exchange, symbol, rdb)
		if err != nil {
			log.Warn("get quote Asset: ", err)
		}
		exchangesymbolMap[symbol] = VerifiableAsset{
			Asset:    asset,
			Verified: assetVerified,
		}

	}
	return exchangesymbolMap, nil
}

// migrateExchangeTrades fetches all trades done on @exchange from the old table with name @fromTable,
// adds underlying asset information if available and stores the extended trade in the table @toTable
// in case store==true. Otherwise, the results are just logged onto the screen.
// It takes individual datastores @dsRead and @dsWrite for reading and writing.
func migrateCEXTrades(exchange string, exchangesymbolMap map[string]VerifiableAsset, timeInit time.Time, timeFinal time.Time, fromTable string, toTable string, testmode bool, dsRead *models.DB, dsWrite *models.DB) error {

	trades, err := dsRead.GetOldTradesFromInflux(fromTable, exchange, false, timeInit, timeFinal)
	if err != nil {
		if strings.Contains(err.Error(), "no trades") {
			log.Warn(err)
		} else {
			log.Error("fetching trades: ", err)
			return err
		}
	}
	log.Info("number of trades: ", len(trades))

	for i := range trades {
		quoteAsset, baseAsset, pairVerified, err := getCEXTradeAssets(trades[i], exchangesymbolMap)
		if err != nil {
			log.Error("get trade assets: ", err)
		}
		trades[i].QuoteToken = quoteAsset
		trades[i].BaseToken = baseAsset
		trades[i].VerifiedPair = pairVerified
		if testmode {
			log.Info("extended trade: ", trades[i])
		} else {
			err = dsWrite.SaveTradeInfluxToTable(&trades[i], toTable)
			if err != nil {
				log.Fatal("save trade: ", err)
			}
		}
	}
	return nil
}

// getTradeAssets returns the underlying assets of a trade and a boolean that tells whether both assets were verified.
func getCEXTradeAssets(trade dia.Trade, exchangesymbolMap map[string]VerifiableAsset) (quoteAsset dia.Asset, baseAsset dia.Asset, pairVerified bool, err error) {
	// Get symbols constituting the trading pair.
	symbols, err := dia.GetPairSymbols(dia.ExchangePair{
		Symbol:      trade.Symbol,
		ForeignName: trade.Pair,
		Exchange:    trade.Source,
	})
	if err != nil {
		return
	}

	// // Get underlying assets.
	// quoteAsset = exchangesymbolMap[symbols[0]].Asset
	// quoteAssetVerified := exchangesymbolMap[symbols[0]].Verified
	// baseAsset = exchangesymbolMap[symbols[1]].Asset
	// baseAssetVerified := exchangesymbolMap[symbols[1]].Verified

	// Get underlying assets.
	quoteAsset = exchangesymbolMap[trade.Source+"-"+symbols[0]].Asset
	quoteAssetVerified := exchangesymbolMap[trade.Source+"-"+symbols[0]].Verified
	baseAsset = exchangesymbolMap[trade.Source+"-"+symbols[1]].Asset
	baseAssetVerified := exchangesymbolMap[trade.Source+"-"+symbols[1]].Verified

	// Verify pair if both assets were verified.
	if quoteAssetVerified && baseAssetVerified {
		pairVerified = true
	}
	err = nil
	return
}

// getExchangeAsset returns the asset corresponding to @symbol on @exchange if it exists, along
// with a boolean that designates whether the symbol was verified.
func getExchangeAsset(exchange string, symbol string, rdb *models.RelDB) (dia.Asset, bool, error) {

	assetID, verified, err := rdb.GetExchangeSymbolAssetID(exchange, symbol)
	if err != nil {
		return dia.Asset{}, verified, err
	}
	if !verified {
		return dia.Asset{}, verified, nil
	}
	asset, err := rdb.GetAssetByID(assetID)
	if err != nil {
		return dia.Asset{}, verified, err
	}
	return asset, verified, nil

}
