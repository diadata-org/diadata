package models

import (
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/go-redis/redis"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
)

// WindowYesterday, Window2, BufferTTL, BiggestWindow & TimeOutRedis constants are found in quotation.go

// ------------------------------------------------------------------------------
// EXCHANGE RATES
// ------------------------------------------------------------------------------

func (db *DB) SetFiatPriceUSD(symbol string, price float64) error {
	return db.SetFiatQuotation(&Quotation{
		Symbol: symbol,
		Name:   helpers.NameForSymbol(symbol),
		Price:  price,
		Source: dia.Diadata,
		Time:   time.Now(),
	})
}

func (db *DB) SetPriorFiatPriceUSD(symbol string, price float64, t time.Time) error {
	return db.SetFiatQuotation(&Quotation{
		Symbol: symbol,
		Name:   helpers.NameForSymbol(symbol),
		Price:  price,
		Source: dia.Diadata,
		Time:   t,
	})
}

func (db *DB) GetFiatPriceUSD(symbol string) (float64, error) {
	key := getKeyQuotation(symbol)
	value := &Quotation{}
	err := db.redisClient.Get(key).Scan(value)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("Error: %v on GetPriceUSD %v\n", err, symbol)
		}
		return 0.0, err
	}
	return value.Price, nil
}

func (db *DB) GetFiatQuotation(symbol string) (*Quotation, error) {
	key := getKeyQuotation(symbol)
	value := &Quotation{}
	err := db.redisClient.Get(key).Scan(value)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("Error: %v on GetQuotation %v\n", err, key)
		}
		return nil, err
	}
	value.Name = helpers.NameForSymbol(symbol) // in case we updated the helper functions ;)
	v, err2 := db.GetPriceYesterday(symbol, "")
	if err2 == nil {
		value.PriceYesterday = &v
	}
	v2, err2 := db.GetVolume(symbol)
	value.VolumeYesterdayUSD = v2
	itin, err := db.GetItinBySymbol(symbol)
	if err != nil {
		value.ITIN = "undefined"
		log.Error(err)
	} else {
		value.ITIN = itin.Itin
	}
	return value, nil
}

func (db *DB) SetFiatQuotation(quotation *Quotation) error {
	if db.influxClient != nil && db.influxBatchPoints != nil {
		return nil
	}
	key := getKeyQuotation(quotation.Symbol)
	log.Debug("setting ", key, quotation)

	fields := map[string]interface{}{
		"quotation": quotation.Price,
	}
	tags := map[string]string{
		"symbol": quotation.Symbol,
		"name":   quotation.Name,
		"source": quotation.Source,
	}

	pt, err := clientInfluxdb.NewPoint(key, tags, fields, quotation.Time)
	if err != nil {
		log.Printf("Error: %v on SetFiatQuotation %v\n", err, quotation.Symbol)
	}

	db.addPoint(pt)

	err = db.WriteBatchInflux()
	if err != nil {
		log.Printf("Error: %v on SetQuotation %v\n", err, quotation.Symbol)
	}

	return err
}
