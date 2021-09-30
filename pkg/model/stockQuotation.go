package models

import (
	"encoding/json"
	"fmt"
	"time"

	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
)

// SetStockQuotationInflux stores a stock quotation to an influx batch.
func (db *DB) SetStockQuotation(sq StockQuotation) error {
	fields := map[string]interface{}{
		"priceAsk": sq.PriceAsk,
		"priceBid": sq.PriceBid,
		"sizeAsk":  sq.SizeAsk,
		"sizeBid":  sq.SizeBid,
		"source":   sq.Source,
	}
	tags := map[string]string{
		"symbol": sq.Symbol,
		"name":   sq.Name,
		"isin":   sq.ISIN,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbStockQuotationsTable, tags, fields, sq.Time)
	if err != nil {
		log.Errorln("NewOptionInflux:", err)
	} else {
		db.addPoint(pt)
	}
	err = db.WriteBatchInflux()
	if err != nil {
		log.Errorln("Write influx batch: ", err)
	}

	return err
}

// GetStockQuotationInflux returns the last quotation of @symbol before @timestamp.
func (db *DB) GetStockQuotation(symbol string, timestamp time.Time) (StockQuotation, error) {
	retval := StockQuotation{}

	unixtime := timestamp.UnixNano()
	q := fmt.Sprintf("SELECT priceAsk,priceBid,sizeAsk,sizeBid,source,\"isin\",\"name\" FROM %s WHERE \"symbol\"='%s' and time<%d order by time desc limit 1", influxDbStockQuotationsTable, symbol, unixtime)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		fmt.Println("Error querying influx")
		return retval, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		layout := "2006-01-02T15:04:05Z"
		vals := res[0].Series[0].Values[0]

		retval.Time, err = time.Parse(layout, vals[0].(string))
		if err != nil {
			log.Error(err)
		}
		retval.PriceAsk, err = vals[1].(json.Number).Float64()
		if err != nil {
			log.Error(err)
		}
		retval.PriceBid, err = vals[2].(json.Number).Float64()
		if err != nil {
			log.Error(err)
		}
		retval.SizeAsk, err = vals[3].(json.Number).Float64()
		if err != nil {
			log.Error(err)
		}
		retval.SizeBid, err = vals[4].(json.Number).Float64()
		if err != nil {
			log.Error(err)
		}
		retval.Source = vals[5].(string)
		if vals[4] != nil {
			retval.ISIN = vals[6].(string)
		}
		retval.Symbol = symbol
		return retval, nil

	}
	return retval, err
}
