package main

import (
	"strconv"
	"strings"
	"time"

	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

const (
	influxDbOldTradesTable = "trades"
	influxDbTestTable      = "tradestest"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
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
	timeFinalString := utils.Getenv("TIME_FINAL", "1634564891")
	testmodestring := utils.Getenv("INFLUX_MIGRATION_TESTMODE", "true")
	var testmode bool
	if testmodestring == "true" {
		testmode = true
	}
	if testmodestring == "false" {
		testmode = false
	}
	timeFinalInt, err := strconv.ParseInt(timeFinalString, 10, 64)
	if err != nil {
		log.Error("parse timeFinal: ", err)
	}
	timeFinal := time.Unix(timeFinalInt, 0)
	log.Info("timeFinal: ", timeFinal)
	// The oldest record we have in influx is younger than 3000 days old.
	timeInit := time.Now().AddDate(0, 0, -3000)

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

	for _, exchange := range dia.Exchanges() {
		if !scrapers.Exchanges[exchange].Centralized {
			continue
		}

		startTime := timeInit
		var endTime time.Time
		for startTime.Before(timeFinal) {
			// If not beyond time range, query one day at a time...
			if timeInit.AddDate(0, 0, 1).Before(timeFinal) {
				endTime = startTime.AddDate(0, 0, 1)
			} else {
				// ...else, query until last time of time range.
				endTime = timeFinal
			}
			err = migrateCEXTrades(exchange, startTime, endTime, fromTable, toTable, testmode, dsRead, dsWrite, rdb)
			if err != nil {
				log.Fatalf("migrate trades for exchange %s: %v", exchange, err)
			}
			log.Infof("migrated trades from %v to %v", startTime, endTime)
			startTime = startTime.AddDate(0, 0, 1)
		}

	}
}

// migrateExchangeTrades fetches all trades done on @exchange from the old table with name @fromTable,
// adds underlying asset information if available and stores the extended trade in the table @toTable
// in case store==true. Otherwise, the results are just logged onto the screen.
func migrateCEXTrades(exchange string, timeInit time.Time, timeFinal time.Time, fromTable string, toTable string, testmode bool, dsRead *models.DB, dsWrite *models.DB, rdb *models.RelDB) error {
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
	if len(trades) > 0 {
		time.Sleep(5 * time.Second)
	}

	for i := range trades {
		quoteAsset, baseAsset, pairVerified, err := getCEXTradeAssets(trades[i], rdb)
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
func getCEXTradeAssets(trade dia.Trade, rdb *models.RelDB) (quoteAsset dia.Asset, baseAsset dia.Asset, pairVerified bool, err error) {
	// Get symbols constituting the trading pair.
	symbols, err := dia.GetPairSymbols(dia.ExchangePair{
		Symbol:      trade.Symbol,
		ForeignName: trade.Pair,
		Exchange:    trade.Source,
	})
	if err != nil {
		return
	}

	// Get underlying assets.
	quoteAsset, quoteAssetVerified, err := getExchangeAsset(trade.Source, symbols[0], rdb)
	if err != nil {
		log.Warnf("get quote Asset %s on %s: %v", symbols[0], trade.Source, err)
	}
	baseAsset, baseAssetVerified, err := getExchangeAsset(trade.Source, symbols[1], rdb)
	if err != nil {
		log.Warnf("get base Asset %s on %s: %v", symbols[1], trade.Source, err)
	}
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
