package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/helpers"
	log "github.com/sirupsen/logrus"
)

// > SELECT MEAN(value) FROM filters WHERE "symbol"='BTC' and "filter"='MA120' GROUP BY TIME(10m) ORDER by time desc limit 10;
func (db *DB) GetChartPoints7Days(symbol string) (r []Point, err error) {
	r = []Point{}
	table := influxDbFiltersTable
	filter := "MA120"

	q := fmt.Sprintf("SELECT time, value FROM %s WHERE time > now() - 7d and filter='%s' and exchange='' and symbol='%s' ORDER BY DESC", table, filter, symbol)

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
		log.Errorln("Empty response GetChartPoints7Days")
	}
	return
}

// GetFilterPoints returns filter points from either a specific exchange or all exchanges
func (db *DB) GetFilterPoints(filter string, exchange string, symbol string, scale string, starttime time.Time, endtime time.Time) (*Points, error) {
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

	q := fmt.Sprintf("SELECT time,exchange, filter, symbol, value FROM %s WHERE filter='%s' %sand symbol='%s' and time>%d and time<%d ORDER BY DESC",
		table, filter, exchangeQuery, symbol, starttime.UnixNano(), endtime.UnixNano())

	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Errorln("GetFilterPoints", err)
	}

	log.Info("GetFilterPoints query: ", q, " returned ", len(res))

	return &Points{
		DataPoints: res,
	}, err
}

func (db *DB) GetLastPriceBefore(symbol string, filter string, exchange string, timestamp time.Time) (Price, error) {
	exchangeQuery := "exchange='" + exchange + "'"
	table := influxDbFiltersTable
	// q := fmt.Sprintf("SELECT LAST(value) FROM %s WHERE filter='%s' AND symbol='%s' AND %s AND time < %d",
	// 	table, filter, symbol, exchangeQuery, timestamp.UnixNano())

	q := fmt.Sprintf("SELECT value FROM %s WHERE filter='%s' AND symbol='%s' AND %s AND time > %d AND time < now() ORDER BY ASC LIMIT 1",
		table, filter, symbol, exchangeQuery, timestamp.UnixNano())

	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Errorln("GetLastFilterPointBefore", err)
	}

	log.Info("GetLastFilterPointBefore query: ", q, " returned ", len(res))

	var price Price
	price.Symbol = symbol
	price.Name = helpers.NameForSymbol(symbol)
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
