package models

import (
	"fmt"
	"time"

	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
)

type BenchmarkedIndex struct {
	Name   string
	Values []BenchmarkedIndexValue
}

type BenchmarkedIndexValue struct {
	CalculationTime time.Time
	Value           string
}

func (datastore *DB) SaveIndexEngineTimeInflux(tags map[string]string, fields map[string]interface{}, timestamp time.Time) error {
	pt, err := clientInfluxdb.NewPoint(influxDbBenchmarkedIndexTableName, tags, fields, timestamp)
	if err != nil {
		log.Errorln("newPoint:", err)
	} else {
		datastore.addPoint(pt)
		errWriteBatchInflux := datastore.WriteBatchInflux()
		if errWriteBatchInflux != nil {
			log.Errorln("newPoint:", errWriteBatchInflux)
		}
	}
	return err
}

func (datastore *DB) GetBenchmarkedIndexValuesInflux(symbol string, starttime time.Time, endtime time.Time) (BenchmarkedIndex, error) {
	var retval BenchmarkedIndex
	q := fmt.Sprintf("SELECT time,\"name\",value from %s WHERE time > %d and time < %d and \"name\" = '%s' ORDER BY time DESC", influxDbBenchmarkedIndexTableName, starttime.UnixNano(), endtime.UnixNano(), symbol)
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		return retval, err
	}

	currentIndex := BenchmarkedIndex{}
	var indexValues []BenchmarkedIndexValue
	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			// Get benchmarked index symbol/name
			currentIndex.Name = res[0].Series[0].Values[i][1].(string)
			// Value
			curValue := res[0].Series[0].Values[i][2].(string)
			// Calculation time
			curCalculationTime, err := time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return currentIndex, err
			}
			indexValue := BenchmarkedIndexValue{
				Value:           curValue,
				CalculationTime: curCalculationTime,
			}
			indexValues = append(indexValues, indexValue)
		}
	}
	currentIndex.Values = indexValues
	return currentIndex, nil
}
