package models

import (
	"fmt"

	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	"github.com/influxdata/influxql"
)

func (datastore *DB) SetBatchFiatPriceInflux(fiatQuotations []*FiatQuotation) error {
	err := checkInfluxIsAvailable(datastore)
	if err != nil {
		return err
	}

	addMultiplePointsToBatch(datastore, fiatQuotations)

	err = datastore.WriteBatchInflux()
	if err != nil {
		log.Printf("Error on WriteBatchInflux: %v\n", err)
	}

	return err
}

func checkInfluxIsAvailable(db *DB) error {
	if db.influxClient == nil {
		return fmt.Errorf("SetInfluxFiatPrice error: no influx database available")
	}
	return nil
}

func addMultiplePointsToBatch(db *DB, fiatQuotations []*FiatQuotation) {
	for _, fq := range fiatQuotations {
		tags := map[string]string{
			"quote_currency": influxql.QuoteString(fq.QuoteCurrency),
			"base_currency":  influxql.QuoteString(fq.BaseCurrency),
			"source":         fq.Source,
		}
		fields := map[string]interface{}{
			"price": fq.Price,
		}

		pt, err := clientInfluxdb.NewPoint(influxDbFiatQuotationsTable, tags, fields, fq.Time)
		if err != nil {
			log.Printf("Error: %v on NewPoint %v\n", err, fq.BaseCurrency)
		}

		db.addPoint(pt)
	}
}

func (datastore *DB) SetSingleFiatPriceRedis(fiatQuotation *FiatQuotation) error {
	err := checkRedisIsAvailable(datastore)
	if err != nil {
		return err
	}

	key := getKeyQuotation(fiatQuotation.QuoteCurrency)
	log.Info("setting ", key, fiatQuotation)

	err = datastore.redisClient.Set(key, fiatQuotation, TimeOutRedis).Err()
	if err != nil {
		log.Printf("Error: %v on SetQuotation %v\n", err, fiatQuotation.QuoteCurrency)
	}

	return err
}

func checkRedisIsAvailable(db *DB) error {
	if db.redisClient == nil {
		return fmt.Errorf("SetRedisFiatPrice error: no redis database available")
	}
	return nil
}
