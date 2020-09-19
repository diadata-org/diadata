package models

import (
	"encoding/json"
	//"errors"
	"fmt"
	//"os"
	//"strconv"
	"time"


	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
)

const influxDbForeignQuotationTable = "foreignquotation"

func (db *DB) SaveForeignQuotationInflux(fq *ForeignQuotation) error {
	tags := map[string]string{"Name": fq.Name}
	fields := map[string]interface{}{
		"symbol": fq.Symbol,
		"price": fq.Price,
		"priceYesterday":  fq.PriceYesterday,
		"volumeYesterdayUSD":  fq.VolumeYesterdayUSD,
		"source":  fq.Source,

	}
	pt, err := clientInfluxdb.NewPoint(influxDbForeignQuotationTable , tags, fields, fq.Time)
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


func (db *DB) GetForeignQuotationInflux(name string, t time.Time) (ForeignQuotation, error) {
	retval := ForeignQuotation{}
	q := fmt.Sprintf("SELECT LAST(price), priceYesterday, volumeYesterdayUSD, source, symbol, Time FROM %s WHERE Name ='%s'", influxDbForeignQuotationTable, name)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		return retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		retval.Name = name
		retval.Time, err = time.Parse(time.RFC3339, res[0].Series[0].Values[0][0].(string))
		if err != nil {
			return retval, err
		}
		retval.Price, err = res[0].Series[0].Values[0][1].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		//retval.Price = price

		*retval.PriceYesterday, err = res[0].Series[0].Values[0][2].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		*retval.VolumeYesterdayUSD, err = res[0].Series[0].Values[0][3].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		retval.Source = res[0].Series[0].Values[0][4].(string)
		if err != nil {
			return retval, err
		}
		retval.Symbol = res[0].Series[0].Values[0][5].(string)
		if err != nil {
			return retval, err
		}
		return retval, nil
	}
	return retval, nil
}
