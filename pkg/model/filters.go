package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/go-redis/redis"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
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

func (datastore *DB) GetFilterPointsAsset(filter string, exchange string, address string, blockchain string, starttime time.Time, endtime time.Time) (*Points, error) {

	exchangeQuery := "AND exchange='" + exchange + "' "

	q := fmt.Sprintf("SELECT time,address,blockchain,exchange,filter,symbol,value FROM %s"+
		" WHERE filter='%s' %s AND address='%s' and blockchain='%s' AND time>%d and time<=%d ORDER BY DESC",
		influxDbFiltersTable, filter, exchangeQuery, address, blockchain, starttime.UnixNano(), endtime.UnixNano())

	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		log.Errorln("GetFilterPoints", err)
	}

	log.Info("GetFilterPoints query: ", q, " returned ", len(res))

	return &Points{
		DataPoints: res,
	}, err
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

	q := fmt.Sprintf("SELECT last(*) FROM %s"+
		" WHERE filter='%s' and address='%s' and blockchain='%s' and time>%d and time<%d and allExchanges=true group by time(1d) fill(previous) ORDER BY DESC",
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
			// if res[0].Series[0].Values[i][2] != nil {
			// 	filterpoint.Name = res[0].Series[0].Values[i][2].(string)
			// }
			// if res[0].Series[0].Values[i][3] != nil {
			// 	filterpoint.Asset.Symbol = res[0].Series[0].Values[i][3].(string)
			// }
			if res[0].Series[0].Values[i][2] != nil {
				filterpoint.Value, err = res[0].Series[0].Values[i][2].(json.Number).Float64()
			} else {
				log.Errorln("res[0].Series[0].Values[i][2]", res[0].Series[0].Values[i][2])
			}
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

// GetFilterAllExchanges returns a slice of quotations for each exchange the asset given by
// @address and @blockchain has a filter value in the given time-range.
// It returns the most recent filter value in the given time-range for each exchange resp.
func (datastore *DB) GetFilterAllExchanges(
	filter string,
	address string,
	blockchain string,
	starttime time.Time,
	endtime time.Time,
) (assetQuotations []AssetQuotation, err error) {
	q := fmt.Sprintf(`
	SELECT time,address,blockchain,symbol,value
	FROM %s
	WHERE filter='%s'
	AND allExchanges=false
	AND address='%s'
	AND blockchain='%s' 
	AND time>%d
	AND time<=%d
	GROUP BY "exchange"
	ORDER BY DESC
	LIMIT 1`,
		influxDbFiltersTable, filter, address, blockchain, starttime.UnixNano(), endtime.UnixNano())

	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		log.Errorln("GetFilterPoints", err)
		return
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series {
			var aq AssetQuotation
			aq.Source = row.Tags["exchange"]
			aq.Price, err = row.Values[0][4].(json.Number).Float64()
			if err != nil {
				log.Warn("parse price: ", err)
			}
			aq.Asset = dia.Asset{Address: address, Blockchain: blockchain, Symbol: row.Values[0][3].(string)}
			aq.Time, err = time.Parse(time.RFC3339, row.Values[0][0].(string))
			if err != nil {
				log.Warn("parse time: ", err)
			}
			assetQuotations = append(assetQuotations, aq)
		}
	}

	return
}

func getKey(filter string, asset dia.Asset, exchange string) string {
	key := filter + "_" + asset.Blockchain + "_" + asset.Address
	if exchange != "" {
		key = key + "_" + exchange
	}
	return key
}

func getKeyFilterZSET(key string) string {
	return "dia_" + key + "_ZSET"
}

func getKeyFilterSymbolAndExchangeZSET(filter string, asset dia.Asset, exchange string) string {
	if exchange == "" {
		return "dia_" + filter + "_" + asset.Blockchain + "_" + asset.Address + "_ZSET"
	} else {
		return "dia_" + filter + "_" + asset.Blockchain + "_" + asset.Address + "_ZSET"
	}
}

// SaveFilterInflux stores a filter point in influx.
func (datastore *DB) SaveFilterInflux(filter string, asset dia.Asset, exchange string, value float64, t time.Time) error {
	// Create a point and add to batch
	tags := map[string]string{
		"filter":     filter,
		"symbol":     EscapeReplacer.Replace(asset.Symbol),
		"address":    EscapeReplacer.Replace(asset.Address),
		"blockchain": asset.Blockchain,
		"exchange":   exchange,
	}
	fields := map[string]interface{}{
		"value":        value,
		"allExchanges": exchange == "",
	}
	pt, err := clientInfluxdb.NewPoint(influxDbFiltersTable, tags, fields, t)
	if err != nil {
		log.Errorln("new filter influx:", err)
	} else {
		datastore.addPoint(pt)
	}
	return err
}

func (datastore *DB) setZSETValue(key string, value float64, unixTime int64, maxWindow int64) error {
	if datastore.redisClient == nil {
		return nil
	}
	member := strconv.FormatFloat(value, 'f', -1, 64) + " " + strconv.FormatInt(unixTime, 10)

	err := datastore.redisPipe.ZAdd(key, redis.Z{
		Score:  float64(unixTime),
		Member: member,
	}).Err()
	log.Debug("SetZSETValue ", key, member, unixTime)
	if err != nil {
		log.Errorf("Error: %v on SetZSETValue %v\n", err, key)
	}
	// purging old values
	err = datastore.redisPipe.ZRemRangeByScore(key, "-inf", "("+strconv.FormatInt(unixTime-maxWindow, 10)).Err()
	if err != nil {
		log.Errorf("Error: %v on SetZSETValue %v\n", err, key)
	}
	if err = datastore.redisPipe.Expire(key, TimeOutRedis).Err(); err != nil {
		log.Error(err)
	} //TODO put two commands together ?
	return err
}

func (datastore *DB) getZSETValue(key string, atUnixTime int64) (float64, error) {

	result := 0.0
	max := strconv.FormatInt(atUnixTime, 10)
	vals, err := datastore.redisClient.ZRangeByScoreWithScores(key, redis.ZRangeBy{
		Min: "-inf",
		Max: max,
	}).Result()
	log.Debug(key, "vals: %v on getZSETValue maxScore: %v", vals, max)
	if err == nil {
		if len(vals) > 0 {
			_, err = fmt.Sscanf(vals[len(vals)-1].Member.(string), "%f", &result)
			if err != nil {
				log.Error(err)
			}
			log.Debugf("returned value: %v", result)
		} else {
			err = errors.New("getZSETValue no value found")
		}
	}
	return result, err
}

func (datastore *DB) getZSETLastValue(key string) (float64, int64, error) {
	value := 0.0
	var unixTime int64
	vals, err := datastore.redisClient.ZRange(key, -1, -1).Result()
	log.Debug(key, "on getZSETLastValue:", vals)
	if err == nil {
		if len(vals) == 1 {
			_, err = fmt.Sscanf(vals[0], "%f %d", &value, &unixTime)
			if err != nil {
				log.Error(err)
			}
			log.Debugf("returned value: %v", value)
		} else {
			err = errors.New("getZSETLastValue no value found")
			log.Errorln("Error: on getZSETLastValue", err, key)
		}
	}
	return value, unixTime, err
}
