package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
)

const influxDbForeignQuotationTable = "foreignquotation"

// SaveForeignQuotationInflux stores a quotation which is not from DIA to an influx batch
func (db *DB) SaveForeignQuotationInflux(fq ForeignQuotation) error {
	tags := map[string]string{"source": fq.Source}
	fields := map[string]interface{}{
		"itin":               fq.ITIN,
		"name":               fq.Name,
		"price":              fq.Price,
		"priceYesterday":     fq.PriceYesterday,
		"symbol":             fq.Symbol,
		"volumeYesterdayUSD": fq.VolumeYesterdayUSD,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbForeignQuotationTable, tags, fields, fq.Time)
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

// GetForeignQuotationInflux returns the last quotation of @symbol before @timestamp
func (db *DB) GetForeignQuotationInflux(symbol, source string, timestamp time.Time) (ForeignQuotation, error) {
	retval := ForeignQuotation{}

	unixtime := timestamp.UnixNano()
	q := fmt.Sprintf("SELECT * FROM %s WHERE source='%s' and symbol='%s' and time<%d order by time desc limit 1", influxDbForeignQuotationTable, source, symbol, unixtime)
	fmt.Println("query: ", q)
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
		retval.ITIN = vals[1].(string)
		retval.Name = vals[2].(string)
		retval.Price, err = vals[3].(json.Number).Float64()
		if err != nil {
			log.Error(err)
		}
		retval.PriceYesterday, err = vals[4].(json.Number).Float64()
		if err != nil {
			log.Error(err)
		}
		retval.Source = vals[5].(string)
		retval.Symbol = vals[6].(string)
		retval.VolumeYesterdayUSD, err = vals[7].(json.Number).Float64()
		if err != nil {
			log.Error(err)
		}
		return retval, nil

	}
	return retval, err
}

// GetForeignPriceYesterday returns the average price of @symbol on @source from yesterday
func (db *DB) GetForeignPriceYesterday(symbol, source string) (float64, error) {

	// Get time range for yesterday in order to average the price
	now := time.Now()
	secondsFromYesterday := now.Hour()*60*60 + now.Minute()*60 + now.Second()
	timeFinal := int(now.Unix()) - secondsFromYesterday - 1
	timeInit := timeFinal - 24*60*60
	unixtimeFinal := strconv.Itoa(timeFinal) + "000000000"
	unixtimeInit := strconv.Itoa(timeInit) + "000000000"

	// Make corresponding influx query
	q := fmt.Sprintf("SELECT price FROM %s WHERE source='%s' and symbol='%s' and time>%s and time<%s", influxDbForeignQuotationTable, source, symbol, unixtimeInit, unixtimeFinal)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		fmt.Println("Error querying influx")
		return 0, err
	}

	// Simple average over all yesterday's prices
	var price float64
	errs := 0
	if len(res) > 0 && len(res[0].Series) > 0 && len(res[0].Series[0].Values) > 0 {
		numPrices := len(res[0].Series[0].Values)
		for i := range res[0].Series[0].Values {
			pricepoint, err := res[0].Series[0].Values[i][1].(json.Number).Float64()
			if err != nil {
				log.Error(err)
				errs++
			} else {
				price += pricepoint
			}

		}
		if numPrices > errs {
			return price / float64(numPrices-errs), nil
		}
	}
	return 0, errors.New("No data available from yesterday")
}

// GetForeignSymbolsInflux returns a list with all symbols available for quotation from @source,
// along with their ITIN.
func (db *DB) GetForeignSymbolsInflux(source string) (symbols []SymbolShort, err error) {

	q := fmt.Sprintf("SELECT distinct(symbol) FROM %s where source='%s'", influxDbForeignQuotationTable, source)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		fmt.Println("Error querying influx")
		return
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		vals := res[0].Series[0].Values
		for _, val := range vals {
			itin, err := db.GetItinBySymbol(val[1].(string))
			if err != nil {
				symbol := SymbolShort{
					Symbol: val[1].(string),
					ITIN:   "",
				}
				symbols = append(symbols, symbol)
			} else {
				symbol := SymbolShort{
					Symbol: val[1].(string),
					ITIN:   itin.Itin,
				}
				symbols = append(symbols, symbol)
			}

		}
	}
	return
}
