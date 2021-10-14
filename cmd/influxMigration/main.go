package main

import (
	"strings"
	"time"

	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

const (
	influxDbOldTradesTable = "trades"
	influxDbTestTable      = "tradestest"
)

func main() {

	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore: ", err)
	}

	rdb, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("relational datastore: ", err)
	}

	// The oldest record we have in influx is younger than 3000 days old.
	timeInit := time.Now().AddDate(0, 0, -3000)
	timeFinal := time.Now()
	testmodestring := utils.Getenv("INFLUX_MIGRATION_TESTMODE", "true")
	var testmode bool
	if testmodestring == "true" {
		testmode = true
	}
	if testmodestring == "false" {
		testmode = false
	}
	fromTable := utils.Getenv("INFLUX_MIGRATION_ORIGIN", influxDbOldTradesTable)
	toTable := utils.Getenv("INFLUX_MIGRATION_DESTINATION", influxDbTestTable)

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
			err = migrateCEXTrades(exchange, startTime, endTime, fromTable, toTable, testmode, ds, rdb)
			if err != nil {
				log.Fatalf("migrate trades for exchange %s: %v", exchange, err)
			}
			startTime = startTime.AddDate(0, 0, 1)
		}

	}
}

// migrateExchangeTrades fetches all trades done on @exchange from the old table with name @fromTable,
// adds underlying asset information if available and stores the extended trade in the table @toTable
// in case store==true. Otherwise, the results are just logged onto the screen.
func migrateCEXTrades(exchange string, timeInit time.Time, timeFinal time.Time, fromTable string, toTable string, testmode bool, ds *models.DB, rdb *models.RelDB) error {
	trades, err := ds.GetOldTradesFromInflux(fromTable, exchange, timeInit, timeFinal)
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
			err = ds.SaveTradeInfluxToTable(&trades[i], toTable)
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
		log.Warn("get quote Asset: ", err)
	}
	baseAsset, baseAssetVerified, err := getExchangeAsset(trade.Source, symbols[1], rdb)
	if err != nil {
		log.Warn("get base Asset: ", err)
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
