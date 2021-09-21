package models

import (
	"fmt"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
)

// SetFilter stores a filter point
func (db *DB) SetFilter(filter string, asset dia.Asset, exchange string, value float64, t time.Time) error {
	err := db.SaveFilterInflux(filter, asset, exchange, value, t)
	if err != nil {
		return err
	}
	err = db.setZSETValue(getKeyFilterZSET(getKey(filter, asset, exchange)), value, t.Unix(), BiggestWindow)
	if err != nil {
		return err
	}
	return err
}

// GetFilterPoints returns filter points from either a specific exchange or all exchanges.
// symbol is mapped to the underlying asset with biggest market cap.
func (db *DB) GetFilterPoints(filter string, exchange string, symbol string, scale string, starttime time.Time, endtime time.Time) (*Points, error) {
	relDB, err := NewRelDataStore()
	if err != nil {
		log.Errorln("NewDataStore:", err)
	}
	// First get asset with @symbol with largest volume.
	topAsset, err := db.GetTopAssetByVolume(symbol, relDB)
	if err != nil {
		log.Error(err)
		return &Points{}, err
	}

	exchangeQuery := "and exchange='" + exchange + "' "
	table := ""
	//	5m 30m 1h 4h 1d 1w
	if scale != "" {
		if filter == "VOL120" {
			table = "a_year.filters_sum_"
		} else {
			table = "a_year.filters_mean_"
		}
		table = table + scale
	} else {
		table = influxDbFiltersTable
	}

	q := fmt.Sprintf("SELECT time,exchange,filter,symbol,value FROM %s"+
		" WHERE filter='%s' %sand address='%s' and blockchain='%s' and time>%d and time<%d ORDER BY DESC",
		table, filter, exchangeQuery, topAsset.Address, topAsset.Blockchain, starttime.UnixNano(), endtime.UnixNano())

	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Errorln("GetFilterPoints", err)
	}

	log.Info("GetFilterPoints query: ", q, " returned ", len(res))

	return &Points{
		DataPoints: res,
	}, err
}
