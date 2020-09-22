package models

import (
	"encoding/json"
	"fmt"
	"time"

	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
)

const influxDbForeignQuotationTable = "foreignquotationi"

func (db *DB) SaveForeignQuotationInflux(fq ForeignQuotation) error {
	tags := map[string]string{"name": fq.Name}
	priceYesterday := *fq.PriceYesterday
	volumeYesterdayUSD := *fq.VolumeYesterdayUSD
	fields := map[string]interface{}{
		"symbol":             fq.Symbol,
		"price":              fq.Price,
		"priceYesterday":     priceYesterday,
		"volumeYesterdayUSD": volumeYesterdayUSD,
		"source":             fq.Source,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbForeignQuotationTable, tags, fields, fq.Time)
	if err != nil {
		log.Errorln("NewOptionInflux:", err)
	} else {
		db.addPoint(pt)
	}

	err = db.WriteBatchInflux()
	if err != nil {
		log.Errorln("SaveForeignQuotationInflux", err)
	}
	log.Println("Coingecko: Foreign Quotation saved")

	return err
}

func (db *DB) GetForeignQuotationInflux(symbol string) (ForeignQuotation, error) {
	retval := ForeignQuotation{}
	q := fmt.Sprintf("SELECT LAST(price) ,priceYesterday, volumeYesterdayUSD, symbol, source FROM %s WHERE symbol='%s'", influxDbForeignQuotationTable, symbol)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		fmt.Println("Error")
		return retval, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		retval.Time, err = time.Parse(time.RFC3339, res[0].Series[0].Values[0][0].(string))
		if err != nil {
			return retval, err
		}

		rawPrice := res[0].Series[0].Values[0][1].(json.Number)
		retval.Price, err = rawPrice.Float64()

		priceYesterday, err := res[0].Series[0].Values[0][2].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		retval.PriceYesterday = &priceYesterday

		volumeYesterdayUSD, err := res[0].Series[0].Values[0][3].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		retval.VolumeYesterdayUSD = &volumeYesterdayUSD

		retval.Symbol = res[0].Series[0].Values[0][4].(string)
		if err != nil {
			return retval, err
		}
		retval.Source = res[0].Series[0].Values[0][5].(string)

		retval.Name = ""
	}
	return retval, nil
}
