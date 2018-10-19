package models

import (
	"fmt"
	"github.com/diadata-org/diadata/pkg/dia"
	clientInfluxdb "github.com/influxdata/influxdb/client/v2"
	log "github.com/sirupsen/logrus"
)

func (db *DB) GetChartPoints(symbol string) ([]Point, error) {
	result := []Point{}
	vals, err := db.redisClient.ZRange(getKeyFilterZSET(dia.FilterKing+"_"+symbol), 0, -1).Result()
	if err == nil {
		var p Point
		for _, v := range vals {
			fmt.Sscanf(v, "%f %d", &p.Value, &p.UnixTime)
			result = append(result, p)
		}
	}
	log.Println(symbol, "GetChartPoints:%v", result)
	return result, err
}

func (db *DB) GetChartPoints7Days(symbol string) ([]float64, error) {
	result := []float64{}
	vals, err := db.redisClient.ZRange(getKeyFilterZSET(dia.FilterKing+"_"+symbol), 0, -1).Result()
	if err == nil {
		if len(vals) != 0 {
			var lastUnixTime int64
			var lastValue float64
			var secondsInOnePoint int64 = 3600 * 24
			var numberOfPoints int64 = 8
			fmt.Sscanf(vals[len(vals)-1], "%f %d", &lastValue, &lastUnixTime)
			lastTime := lastUnixTime - secondsInOnePoint*numberOfPoints
			var value float64
			var unixTime int64
			for _, v := range vals {
				fmt.Sscanf(v, "%f %d", &value, &unixTime)
				if unixTime-lastTime > secondsInOnePoint {
					lastTime = unixTime
					result = append(result, value)
				}
			}
			result[len(result)-1] = lastValue
		}
	}
	log.Println(symbol, "GetChartPoints:%v", result)
	return result, err
}

func (db *DB) GetFilterPoints(filter string, exchange string, symbol string, scale string) ([]clientInfluxdb.Result, error) {
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

	q := fmt.Sprintf("SELECT * FROM %s WHERE filter='%s' %sand symbol='%s' ORDER BY DESC", table, filter, exchangeQuery, symbol)

	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Errorln("GetFilterPoints", err)
	}

	log.Info("GetFilterPoints query:", q, "returned ", len(res))

	return res, err
}
