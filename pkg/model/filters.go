package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
)

// SetFilter stores a filter point
func (datastore *DB) SetFilter(filter string, asset dia.Asset, exchange string, value float64, t time.Time) error {
	// return datastore.SaveFilterInflux(filter, asset, exchange, value, t)
	err := datastore.SaveFilterInflux(filter, asset, exchange, value, t)
	if err != nil {
		return err
	}
	err = datastore.setZSETValue(getKeyFilterZSET(getKey(filter, asset, exchange)), value, t.Unix(), BiggestWindow)
	if err != nil {
		return err
	}
	return err
}

// GetFilterPoints returns filter points from either a specific exchange or all exchanges.
// symbol is mapped to the underlying asset with biggest market cap.
func (datastore *DB) GetFilterPoints(filter string, exchange string, symbol string, scale string, starttime time.Time, endtime time.Time) (*Points, error) {
	relDB, err := NewRelDataStore()
	if err != nil {
		log.Errorln("NewDataStore:", err)
	}
	// First get asset with @symbol with largest volume.
	// topAsset, err := db.GetTopAssetByVolume(symbol, relDB)
	sortedAssets, err := relDB.GetTopAssetByVolume(symbol)
	if err != nil {
		log.Error(err)
		return &Points{}, err
	}
	if len(sortedAssets) == 0 {
		return nil, errors.New("no traded assets found")
	}
	topAsset := sortedAssets[0]

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

	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		log.Errorln("GetFilterPoints", err)
	}

	log.Info("GetFilterPoints query: ", q, " returned ", len(res))

	return &Points{
		DataPoints: res,
	}, err
}

type FilterPoint struct {
	Time     time.Time
	Exchange string
	Filter   string
	Symbol   string
	Value    float64
}

func (datastore *DB) GetFilter(filter string, topAsset dia.Asset, scale string, starttime time.Time, endtime time.Time) ([]dia.FilterPoint, error) {
	var allFilters []dia.FilterPoint
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
		" WHERE filter='%s' and address='%s' and blockchain='%s' and time>%d and time<%d ORDER BY DESC",
		table, filter, topAsset.Address, topAsset.Blockchain, starttime.UnixNano(), endtime.UnixNano())

	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		log.Errorln("GetFilterPoints", err)
	}
	log.Info("GetFilterPoints query: ", q, " returned ", len(res))

	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			var filterpoint dia.FilterPoint
			filterpoint.Time, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return allFilters, err
			}
			// if res[0].Series[0].Values[i][1] != nil {
			// 	filterpoint.Exchange = res[0].Series[0].Values[i][1].(string)
			// }
			if res[0].Series[0].Values[i][2] != nil {
				filterpoint.Name = res[0].Series[0].Values[i][2].(string)
			}
			if res[0].Series[0].Values[i][3] != nil {
				filterpoint.Asset.Symbol = res[0].Series[0].Values[i][3].(string)
			}
			filterpoint.Value, err = res[0].Series[0].Values[i][1].(json.Number).Float64()
			if err != nil {
				return allFilters, err
			}
			allFilters = append(allFilters, filterpoint)
		}
	} else {
		return allFilters, errors.New("no filter points in time range")
	}

	return allFilters, err
}
