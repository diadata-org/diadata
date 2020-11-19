package models

import (
	"time"

	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
)

func (db *DB) SetFiatPriceUSD(symbol string, price float64) error {
	return db.SetFiatQuotation(&FiatQuotation{
		QuoteCurrency: symbol,
		BaseCurrency:  "USD",
		Price:         price,
		Source:        "ECB",
		Time:          time.Now(),
	})
}

func (db *DB) SetPriorFiatPriceUSD(symbol string, price float64, timestamp time.Time) error {
	return db.SetFiatQuotation(&FiatQuotation{
		QuoteCurrency: symbol,
		BaseCurrency:  "USD",
		Price:         price,
		Source:        "ECB",
		Time:          timestamp,
	})
}

func (db *DB) SetFiatQuotation(quotation *FiatQuotation) error {
	if db.influxClient == nil && db.influxBatchPoints == nil {
		return nil
	}

	log.Debug("setting ", quotation)

	tags := map[string]string{
		"quote_currency": quotation.QuoteCurrency,
		"base_currency":  quotation.BaseCurrency,
		"source":         quotation.Source,
	}
	fields := map[string]interface{}{
		"price": quotation.Price,
	}

	pt, err := clientInfluxdb.NewPoint(influxDbFiatQuotationsTable, tags, fields, quotation.Time)
	if err != nil {
		log.Printf("Error: %v on NewPoint %v\n", err, quotation.BaseCurrency)
	}

	db.addPoint(pt)

	err = db.WriteBatchInflux()
	if err != nil {
		log.Printf("Error: %v on WriteBatchInflux %v\n", err, quotation.BaseCurrency)
	}

	return err
}
