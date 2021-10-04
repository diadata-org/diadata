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
		"sizeAsk":  sq.SizeAskLot,
		"sizeBid":  sq.SizeBidLot,
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
func (db *DB) GetStockQuotation(source string, symbol string, timeInit time.Time, timeFinal time.Time) ([]StockQuotation, error) {
	stockQuotations := []StockQuotation{}

	unixtimeInit := timeInit.UnixNano()
	unixtimeFinal := timeFinal.UnixNano()

	query := "SELECT priceAsk,priceBid,sizeAsk,sizeBid,source,\"isin\",\"name\" FROM %s WHERE source='%s' and \"symbol\"='%s' and time>%d and time<=%d order by time desc"
	q := fmt.Sprintf(query, influxDbStockQuotationsTable, source, symbol, unixtimeInit, unixtimeFinal)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		fmt.Println("Error querying influx")
		return stockQuotations, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		layout := "2006-01-02T15:04:05Z"
		vals := res[0].Series[0].Values

		for i := 0; i < len(vals); i++ {
			var stockQuotation StockQuotation
			stockQuotation.Time, err = time.Parse(layout, vals[i][0].(string))
			if err != nil {
				log.Error(err)
			}
			stockQuotation.PriceAsk, err = vals[i][1].(json.Number).Float64()
			if err != nil {
				log.Error(err)
			}
			stockQuotation.PriceBid, err = vals[i][2].(json.Number).Float64()
			if err != nil {
				log.Error(err)
			}
			stockQuotation.SizeAskLot, err = vals[i][3].(json.Number).Float64()
			if err != nil {
				log.Error(err)
			}
			stockQuotation.SizeBidLot, err = vals[i][4].(json.Number).Float64()
			if err != nil {
				log.Error(err)
			}
			stockQuotation.Source = vals[i][5].(string)
			stockQuotation.ISIN = vals[i][6].(string)
			stockQuotation.Name = vals[i][7].(string)
			stockQuotation.Symbol = symbol

			stockQuotations = append(stockQuotations, stockQuotation)
		}
		return stockQuotations, nil

	}
	return stockQuotations, err
}

// GetStockSymbols returns all symbols available from @source.
func (db *DB) GetStockSymbols() (map[Stock]string, error) {
	allStocks := make(map[Stock]string)

	q := fmt.Sprintf("SELECT \"symbol\",\"name\",\"isin\",source FROM %s WHERE time>now()-7d", influxDbStockQuotationsTable)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Error("query stock symbols from influx: ", err)
		return allStocks, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		// make unique list of stocks - i.e. pairs (isin,source) should be unique.
		vals := res[0].Series[0].Values
		set := make(map[string]struct{})

		for _, val := range vals {
			if _, ok := set[val[3].(string)+val[4].(string)]; !ok {
				allStocks[Stock{
					Symbol: val[1].(string),
					Name:   val[2].(string),
					ISIN:   val[3].(string),
				}] = val[4].(string)
				set[val[3].(string)+val[4].(string)] = struct{}{}
			}
		}

	}

	return allStocks, nil
}

// func (db *DB) GetStockQuotfation(starttime time.Time, endtime time.Time, asset string, protocol string) ([]StockQuotation, error) {
// 	retval := []dia.DefiRate{}
// 	influxQuery := "SELECT \"asset\",borrowRate,lendingRate,\"protocol\" FROM %s WHERE time > %d and time < %d and asset = '%s' and protocol = '%s'"
// 	q := fmt.Sprintf(influxQuery, influxDbDefiRateTable, starttime.UnixNano(), endtime.UnixNano(), asset, protocol)
// 	fmt.Println("influx query: ", q)
// 	res, err := queryInfluxDB(db.influxClient, q)
// 	fmt.Println("res, err: ", res, err)
// 	if err != nil {
// 		return retval, err
// 	}
// 	if len(res) > 0 && len(res[0].Series) > 0 {
// 		for i := 0; i < len(res[0].Series[0].Values); i++ {
// 			currentRate := dia.DefiRate{}
// 			currentRate.Timestamp, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
// 			if err != nil {
// 				return retval, err
// 			}
// 			currentRate.Asset = res[0].Series[0].Values[i][1].(string)
// 			if err != nil {
// 				return retval, err
// 			}
// 			currentRate.BorrowingRate, err = res[0].Series[0].Values[i][2].(json.Number).Float64()
// 			if err != nil {
// 				return retval, err
// 			}
// 			currentRate.LendingRate, err = res[0].Series[0].Values[i][3].(json.Number).Float64()
// 			if err != nil {
// 				return retval, err
// 			}
// 			currentRate.Protocol = protocol
// 			retval = append(retval, currentRate)
// 		}
// 	} else {
// 		return retval, errors.New("Error parsing Defi Lending Rate from Database")
// 	}
// 	return retval, nil
// }
