package models

import (
	"fmt"

	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
)

func (db *DB) SetBatchFiatPriceInflux(fiatQuotations []*FiatQuotation) error {
	err := checkInfluxIsAvailable(db)
	if err != nil {
		return err
	}

	addMultiplePointsToBatch(db, fiatQuotations)

	err = db.WriteBatchInflux()
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
			"quote_currency": fq.QuoteCurrency,
			"base_currency":  fq.BaseCurrency,
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

func (db *DB) SetSingleFiatPriceRedis(fiatQuotation *FiatQuotation) error {
	err := checkRedisIsAvailable(db)
	if err != nil {
		return err
	}

	key := getKeyQuotation(fiatQuotation.QuoteCurrency)
	log.Info("setting ", key, fiatQuotation)

	err = db.redisClient.Set(key, fiatQuotation, TimeOutRedis).Err()
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
