package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
)

func (datastore *DB) GetPrice(asset dia.Asset, exchange string) (float64, error) {
	key := getKeyFilterSymbolAndExchangeZSET(dia.FilterKing, asset, exchange)
	v, _, err := datastore.getZSETLastValue(key)
	return v, err
}

func (datastore *DB) GetPriceYesterday(asset dia.Asset, exchange string) (float64, error) {
	return datastore.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, asset, exchange)), time.Now().Unix()-WindowYesterday)
}

func (datastore *DB) GetPrice1h(asset dia.Asset, exchange string) (float64, error) {
	return datastore.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, asset, exchange)), time.Now().Unix()-Window1h)
}

func (datastore *DB) GetPrice7d(asset dia.Asset, exchange string) (float64, error) {
	return datastore.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, asset, exchange)), time.Now().Unix()-Window7d)
}

func (datastore *DB) GetPrice14d(asset dia.Asset, exchange string) (float64, error) {
	return datastore.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, asset, exchange)), time.Now().Unix()-Window14d)
}

func (datastore *DB) GetPrice30d(asset dia.Asset, exchange string) (float64, error) {
	return datastore.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, asset, exchange)), time.Now().Unix()-Window30d)
}

func (datastore *DB) GetTradePriceBefore(asset dia.Asset, exchange string, timestamp time.Time, window time.Duration) (*dia.Trade, error) {
	return datastore.GetTradeInflux(asset, exchange, timestamp, window)
}

func (datastore *DB) GetTradePrice1h(asset dia.Asset, exchange string) (*dia.Trade, error) {
	return datastore.GetTradePriceBefore(asset, exchange, time.Now().Add(-1*time.Hour), time.Duration(4*time.Hour))
}

func (datastore *DB) GetTradePrice24h(asset dia.Asset, exchange string) (*dia.Trade, error) {
	return datastore.GetTradePriceBefore(asset, exchange, time.Now().Add(-24*time.Hour), time.Duration(7*24*time.Hour))
}

func (datastore *DB) GetTradePrice7d(asset dia.Asset, exchange string) (*dia.Trade, error) {
	return datastore.GetTradePriceBefore(asset, exchange, time.Now().Add(-7*24*time.Hour), time.Duration(14*24*time.Hour))
}

func (datastore *DB) GetTradePrice14d(asset dia.Asset, exchange string) (*dia.Trade, error) {
	return datastore.GetTradePriceBefore(asset, exchange, time.Now().Add(-14*24*time.Hour), time.Duration(28*24*time.Hour))
}

func (datastore *DB) GetTradePrice30d(asset dia.Asset, exchange string) (*dia.Trade, error) {
	return datastore.GetTradePriceBefore(asset, exchange, time.Now().Add(-30*24*time.Hour), time.Duration(60*24*time.Hour))
}

func (datastore *DB) GetLastPriceBefore(asset dia.Asset, filter string, exchange string, timestamp time.Time) (Price, error) {
	exchangeQuery := "exchange='" + exchange + "'"
	table := influxDbFiltersTable
	// q := fmt.Sprintf("SELECT LAST(value) FROM %s WHERE filter='%s' AND symbol='%s' AND %s AND time < %d",
	// 	table, filter, symbol, exchangeQuery, timestamp.UnixNano())

	q := fmt.Sprintf("SELECT value FROM %s WHERE filter='%s' AND address='%s' AND blockchain='%s' AND %s AND time<now() AND time > %d ORDER BY ASC LIMIT 1",
		table, filter, asset.Address, asset.Blockchain, exchangeQuery, timestamp.UnixNano())

	res, err := queryInfluxDB(datastore.influxClient, q)
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
