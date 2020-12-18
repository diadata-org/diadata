package models

import (
//	"fmt"
	"time"
	//log "github.com/sirupsen/logrus"
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
	CirculatingSupply float64
	Weight            float64
	CappingFactor     float64
}

/*func (db *DB) GetCryptoIndex(name string) (*CryptoIndex, error) {
	return nil, nil
}

func (db *DB) SetCryptoIndex(index *CryptoIndex) error {
	constituentsSerial := ""
	for _, c := range index.Constituents {
		constituentsSerial += ","
		constituentsSerial += c.Symbol
	}
	fields := map[string]interface{}{
		"price": index.Price,
		"value": index.Value,
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
		err = db.SetCryptoIndexConstituent(constituent)
		if err != nil {
			return err
		}
	}
	return err

}

func (db *DB) GetCryptoIndexConstituent(starttime time.Time, endtime time.Time, symbol string) ([]CryptoIndexConstituent, error) {
	var retval CryptoIndexConstituent
	q := fmt.Sprintf("SELECT * from %s WHERE time > %d and time < %d and symbol = '%s'", influxDbCryptoIndexConstituentsTable, starttime.UnixNano(), endtime.UnixNano(), symbol)
	res, err := queryInfluxDB(db.influxClient, q)

	if err != nil {
		return retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			currentConstituent := CryptoIndexConstituent{}
			currentConstituent.Timestamp, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return retval, err
			}
		}
	}
	return nil, nil
}

func (db *DB) SetCryptoIndexConstituent(constituent *CryptoIndexConstituent) error {
	fields := map[string]interface{}{
		"price": constituent.Price,
		"circulatingsupply": constituent.CirculatingSupply,
		"weight": constituent.Weight,
		"cappingfactor": constituent.CappingFactor,
	}
	tags := map[string]string{
		"name": constituent.Name,
		"symbol": constituent.Symbol,
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
}*/
