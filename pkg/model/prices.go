package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
)

func (db *DB) SetPriceZSET(symbol string, exchange string, price float64, t time.Time) error {
	db.SaveFilterInflux(dia.FilterKing, symbol, exchange, price, t)
	key := getKeyFilterZSET(getKey(dia.FilterKing, symbol, exchange))
	log.Debug("SetPriceZSET ", key)
	return db.setZSETValue(key, price, time.Now().Unix(), BiggestWindow)
}

func (db *DB) GetPrice(symbol string, exchange string) (float64, error) {
	key := getKeyFilterSymbolAndExchangeZSET(dia.FilterKing, symbol, exchange)
	v, _, err := db.getZSETLastValue(key)
	return v, err
}

func (db *DB) GetPriceYesterday(symbol string, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, symbol, exchange)), time.Now().Unix()-WindowYesterday)
}

// GetAssetPriceInflux returns the last price of asset with @symbol before @timeFinal as retrieved from MA120 filter in influx.
// Data in the time range [@timeInit, @timeFinal] is taken into account (make time window smaller for quicker result).
func (db *DB) GetAssetPriceInflux(exchange, symbol string, timeInit, timeFinal time.Time) (float64, time.Time, error) {

	unixTimeInit := timeInit.UnixNano()
	unixTimeFinal := timeFinal.UnixNano()
	q := fmt.Sprintf("SELECT value FROM filters WHERE filter='MA120' and exchange='%s' and symbol='%s' and time>%d and time<=%d ORDER BY DESC limit 1", exchange, symbol, unixTimeInit, unixTimeFinal)
	log.Info("influx query: ", q)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Errorf("Influx query %s error: %s", q, err)
		return float64(0), time.Time{}, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 && len(res[0].Series[0].Values) > 0 {
		val := res[0].Series[0].Values[0]

		price, err := val[1].(json.Number).Float64()
		if err != nil {
			return float64(0), time.Time{}, err
		}

		timestamp, err := time.Parse(time.RFC3339, val[0].(string))
		if err != nil {
			return float64(0), time.Time{}, err
		}
		return price, timestamp, nil
	}

	return float64(0), time.Time{}, errors.New("influx: empty response")
}
