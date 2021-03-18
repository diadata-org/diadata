package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
)

func (db *DB) GetPrice(asset dia.Asset, exchange string) (float64, error) {
	key := getKeyFilterSymbolAndExchangeZSET(dia.FilterKing, asset, exchange)
	v, _, err := db.getZSETLastValue(key)
	return v, err
}

func (db *DB) GetPriceYesterday(asset dia.Asset, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, asset, exchange)), time.Now().Unix()-WindowYesterday)
}

func (db *DB) GetPrice1h(asset dia.Asset, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, asset, exchange)), time.Now().Unix()-Window1h)
}

func (db *DB) GetPrice7d(asset dia.Asset, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, asset, exchange)), time.Now().Unix()-Window7d)
}

func (db *DB) GetPrice14d(asset dia.Asset, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, asset, exchange)), time.Now().Unix()-Window14d)
}

func (db *DB) GetPrice30d(asset dia.Asset, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, asset, exchange)), time.Now().Unix()-Window30d)
}

func (db *DB) GetTradePriceBefore(asset dia.Asset, exchange string, timestamp time.Time) (*dia.Trade, error) {
	return db.GetTradeInflux(asset, exchange, timestamp)
}

func (db *DB) GetTradePrice1h(asset dia.Asset, exchange string) (*dia.Trade, error) {
	return db.GetTradePriceBefore(asset, exchange, time.Now().Add(-1*time.Hour))
}

func (db *DB) GetTradePrice24h(asset dia.Asset, exchange string) (*dia.Trade, error) {
	return db.GetTradePriceBefore(asset, exchange, time.Now().Add(-24*time.Hour))
}

func (db *DB) GetTradePrice7d(asset dia.Asset, exchange string) (*dia.Trade, error) {
	return db.GetTradePriceBefore(asset, exchange, time.Now().Add(-7*24*time.Hour))
}

func (db *DB) GetTradePrice14d(asset dia.Asset, exchange string) (*dia.Trade, error) {
	return db.GetTradePriceBefore(asset, exchange, time.Now().Add(-14*24*time.Hour))
}

func (db *DB) GetTradePrice30d(asset dia.Asset, exchange string) (*dia.Trade, error) {
	return db.GetTradePriceBefore(asset, exchange, time.Now().Add(-30*24*time.Hour))
}

func (db *DB) GetLastPriceBefore(asset dia.Asset, filter string, exchange string, timestamp time.Time) (Price, error) {
	exchangeQuery := "exchange='" + exchange + "'"
	table := influxDbFiltersTable
	// q := fmt.Sprintf("SELECT LAST(value) FROM %s WHERE filter='%s' AND symbol='%s' AND %s AND time < %d",
	// 	table, filter, symbol, exchangeQuery, timestamp.UnixNano())

	q := fmt.Sprintf("SELECT value FROM %s WHERE filter='%s' AND address='%s' AND blockchain='%s' AND %s AND time > %d ORDER BY ASC LIMIT 1",
		table, filter, asset.Address, asset.Blockchain.Name, exchangeQuery, timestamp.UnixNano())

	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Errorln("GetLastFilterPointBefore", err)
	}

	log.Info("GetLastFilterPointBefore query: ", q, " returned ", len(res))

	var price Price
	price.Symbol = asset.Symbol
	price.Name = helpers.NameForSymbol(asset.Symbol)
	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {
			price.Time, err = time.Parse(time.RFC3339, row[0].(string))
			if err == nil {
				value, ok := row[1].(json.Number)
				if ok {
					price.Price, _ = value.Float64()
				} else {
					log.Errorln("GetLastFilterPointBefore: error on parsing row price", row)
				}
			} else {
				log.Errorln("GetLastFilterPointBefore: error on parsing row time", err, row)
			}
		}
	} else {
		log.Errorln("Empty response GetLastFilterPointBefore")
	}

	return price, err
}
