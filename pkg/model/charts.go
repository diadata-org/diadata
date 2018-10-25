package models

import (
	"encoding/json"
	"fmt"
	//"github.com/diadata-org/diadata/pkg/dia"
	clientInfluxdb "github.com/influxdata/influxdb/client/v2"
	log "github.com/sirupsen/logrus"
	"time"
)

/*
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
*/

func (db *DB) GetChartPoints7Days(symbol string) (r []Point, err error) {
	r = []Point{}
	table := influxDbFiltersTable
	filter := "MA120"

	q := fmt.Sprintf("SELECT time, value FROM %s WHERE filter='%s' and exchange='' and symbol='%s' ORDER BY DESC", table, filter, symbol)

	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Errorln("GetFilterPoints", err)
	}

	log.Info("GetFilterPoints query:", q, "returned ", len(res))

	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {
			t, err := time.Parse(time.RFC3339, row[0].(string))
			if err == nil {
				var price float64
				v, o := row[1].(json.Number)
				if o {
					price, _ = v.Float64()
					r = append(r, Point{t.Unix(), price})
				} else {
					log.Errorln("GetChartPoints7Days: error on parsing row price", row)
				}
			} else {
				log.Errorln("GetChartPoints7Days: error on parsing row time", err, row)
			}
		}
	} else {
		log.Errorln("Empty res", len(res), res)
	}
	return
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
