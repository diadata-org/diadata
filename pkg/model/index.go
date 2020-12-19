package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
)

const (
	indexNormalization = float64(9507172.247746756)
)

// CryptoIndex is the container for API endpoint CryptoIndex
type CryptoIndex struct {
	Name              string
	Value             float64
	Price             float64
	Price1h           float64
	Price24h          float64
	Price7d           float64
	Price14d          float64
	Price30d          float64
	Volume24hUSD      float64
	CirculatingSupply float64
	CalculationTime   time.Time
	Constituents      []CryptoIndexConstituent
}

type CryptoIndexConstituent struct {
	Name              string
	Symbol            string
	Address           string
	Price             float64
	PriceYesterday    float64
	CirculatingSupply float64
	Weight            float64
	CappingFactor     float64
}

func (db *DB) GetCryptoIndex(starttime time.Time, endtime time.Time, name string) ([]CryptoIndex, error) {
	var retval []CryptoIndex
	q := fmt.Sprintf("SELECT * from %s WHERE time > %d and time < %d and \"name\" = '%s' ORDER BY time DESC LIMIT 10", influxDbCryptoIndexTable, starttime.UnixNano(), endtime.UnixNano(), name)
	res, err := queryInfluxDB(db.influxClient, q)

	if err != nil {
		return retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			currentIndex := CryptoIndex{}
			currentIndex.CalculationTime, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return retval, err
			}
			constituentsSerial := res[0].Series[0].Values[i][1].(string)
			currentIndex.Name = res[0].Series[0].Values[i][2].(string)
			currentIndex.Price, err = res[0].Series[0].Values[i][3].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			tmp, err := res[0].Series[0].Values[i][4].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			tmp /= indexNormalization
			currentIndex.Value = tmp
			var constituents []CryptoIndexConstituent
			// Get constituents
			for _, constituentSymbol := range strings.Split(constituentsSerial, ",") {
				curr, err := db.GetCryptoIndexConstituents(currentIndex.CalculationTime.Add(-5*time.Hour), endtime, constituentSymbol)
				if err != nil {
					return retval, err
				}
				if len(curr) > 0 {
					constituents = append(constituents, curr[0])
				}
			}
			currentIndex.Constituents = constituents
			retval = append(retval, currentIndex)
		}
	}
	return retval, nil
}

func (db *DB) SetCryptoIndex(index *CryptoIndex) error {
	constituentsSerial := ""
	for _, c := range index.Constituents {
		if constituentsSerial != "" {
			constituentsSerial += ","
		}
		constituentsSerial += c.Symbol
	}
	fields := map[string]interface{}{
		"price":        index.Price,
		"value":        index.Value,
		"constituents": constituentsSerial,
	}
	tags := map[string]string{
		"name": index.Name,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbCryptoIndexTable, tags, fields, index.CalculationTime)
	if err != nil {
		log.Error("Writing Crypto Index to Influx: ", err)
		return err
	} else {
		db.addPoint(pt)
	}

	err = db.WriteBatchInflux()
	if err != nil {
		log.Error("Writing Crypto Index to Influx: ", err)
		return err
	}

	for _, constituent := range index.Constituents {
		err = db.SetCryptoIndexConstituent(&constituent)
		if err != nil {
			return err
		}
	}
	return err
}

func (db *DB) GetCryptoIndexConstituentPrice(symbol string, date time.Time) (float64, error) {
	startdate := date.AddDate(0, 0, -1)
	q := fmt.Sprintf("SELECT price from %s where time > %d and time <= %d and symbol = '%s' ORDER BY time DESC LIMIT 1", influxDbCryptoIndexConstituentsTable, startdate.UnixNano(), date.UnixNano(), symbol)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		return float64(0), err
	}
	var price float64
	if len(res) > 0 && len(res[0].Series) > 0 && len(res[0].Series[0].Values) > 0 {
		price, err = res[0].Series[0].Values[0][1].(json.Number).Float64()
	}
	return price, nil

}

func (db *DB) GetCryptoIndexConstituents(starttime time.Time, endtime time.Time, symbol string) ([]CryptoIndexConstituent, error) {
	var retval []CryptoIndexConstituent
	q := fmt.Sprintf("SELECT * from %s WHERE time > %d and time < %d and symbol = '%s' ORDER BY time DESC LIMIT 10", influxDbCryptoIndexConstituentsTable, starttime.UnixNano(), endtime.UnixNano(), symbol)
	res, err := queryInfluxDB(db.influxClient, q)

	if err != nil {
		return retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			currentConstituent := CryptoIndexConstituent{}
			/*currentConstituent.Time, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return retval, err
			}*/ //TODO: Do we need time?
			currentConstituent.Address = res[0].Series[0].Values[i][1].(string)
			currentConstituent.CappingFactor, err = res[0].Series[0].Values[i][2].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			currentConstituent.CirculatingSupply, err = res[0].Series[0].Values[i][3].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			currentConstituent.Name = res[0].Series[0].Values[i][4].(string)
			currentConstituent.Price, err = res[0].Series[0].Values[i][5].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			currentConstituent.Symbol = res[0].Series[0].Values[i][6].(string)
			currentConstituent.Weight, err = res[0].Series[0].Values[i][7].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			// Get price yesterday
			priceYesterday, err := db.GetCryptoIndexConstituentPrice(currentConstituent.Symbol, endtime.AddDate(0, 0, -1))
			if err != nil {
				currentConstituent.PriceYesterday = float64(0)
			} else {
				currentConstituent.PriceYesterday = priceYesterday
			}

			retval = append(retval, currentConstituent)
		}
	}
	return retval, nil
}

func (db *DB) SetCryptoIndexConstituent(constituent *CryptoIndexConstituent) error {
	fields := map[string]interface{}{
		"price":             constituent.Price,
		"circulatingsupply": constituent.CirculatingSupply,
		"weight":            constituent.Weight,
		"cappingfactor":     constituent.CappingFactor,
	}
	tags := map[string]string{
		"name":    constituent.Name,
		"symbol":  constituent.Symbol,
		"address": constituent.Address,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbCryptoIndexConstituentsTable, tags, fields, time.Now())
	if err != nil {
		log.Error("Writing Crypto Index Constituent to Influx: ", err)
		return err
	} else {
		db.addPoint(pt)
	}

	err = db.WriteBatchInflux()
	if err != nil {
		log.Error("Writing Crypto Index Constituent to Influx: ", err)
	}
	return err
}
