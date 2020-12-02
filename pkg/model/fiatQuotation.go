package models

import (
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
)

// Format fiatQuotation structs as influxdb points and write them
func (db *DB) SetFiatPriceUSD(fqs []*FiatQuotation) error {
	for _, fq := range fqs {
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

	err := db.WriteBatchInflux()
	if err != nil {
		log.Printf("Error on WriteBatchInflux: %v\n", err)
	}

	return err
}
