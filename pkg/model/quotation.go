package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/go-redis/redis"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
)

const (
	WindowYesterday = 24 * 60 * 60
	Window1h        = 60 * 60
	Window7d        = 7 * 24 * 60 * 60
	Window14d       = 7 * 24 * 60 * 60
	Window30d       = 30 * 24 * 60 * 60
	Window2         = 24 * 60 * 60 * 8
	BufferTTL       = 60 * 60
	BiggestWindow   = Window2
	TimeOutRedis    = time.Duration(time.Second * (BiggestWindow + BufferTTL))
)

func getKeyQuotation(value string) string {
	return "dia_quotation_USD_" + value
}

func getKeyQuotationEUR(value string) string {
	return "dia_quotation_EUR_" + value
}

func getKeyAssetQuotation(blockchain, address string) string {
	return "dia_assetquotation_USD_" + blockchain + "_" + address
}

// ------------------------------------------------------------------------------
// ASSET EXCHANGE RATES (WIP)
// ------------------------------------------------------------------------------

// SetAssetPriceUSD stores the price of @asset in influx and the caching layer.
// The latter only holds the most recent price point.
func (db *DB) SetAssetPriceUSD(asset dia.Asset, price float64, timestamp time.Time) error {
	return db.SetAssetQuotation(&AssetQuotation{
		Asset:  asset,
		Price:  price,
		Source: dia.Diadata,
		Time:   timestamp,
	})
}

// GetAssetPriceUSD returns the last price of @asset.
func (db *DB) GetAssetPriceUSD(asset dia.Asset) (price float64, err error) {
	// TO DO
	return
}

// SetAssetQuotation stores the full quotation of @asset into influx and cache.
func (db *DB) SetAssetQuotation(quotation *AssetQuotation) error {
	// Write to influx
	tags := map[string]string{
		"symbol":     quotation.Asset.Symbol,
		"name":       quotation.Asset.Name,
		"address":    quotation.Asset.Address,
		"blockchain": quotation.Asset.Blockchain.Name,
	}
	fields := map[string]interface{}{
		"price":           quotation.Price,
		"priceYesterday":  quotation.PriceYesterday,
		"volumeYesterday": quotation.VolumeYesterdayUSD,
	}

	pt, err := clientInfluxdb.NewPoint(influxDBAssetQuotationsTable, tags, fields, quotation.Time)
	if err != nil {
		log.Errorln("NewTradeInflux:", err)
	} else {
		db.addPoint(pt)
	}
	err = db.WriteBatchInflux()
	if err != nil {
		log.Error("Write batch for SetAssetQuotation: ", err)
	}

	// Write latest point to redis cache
	// TO DO: make check for timestamp of cache and only write if younger
	return db.SetAssetQuotationCache(quotation)

}

// GetAssetQuotation returns the latest full quotation for @asset.
func (db *DB) GetAssetQuotation(asset dia.Asset) (*AssetQuotation, error) {

	// First attempt to get latest quotation from redis cache
	quotation, err := db.GetAssetQuotationCache(asset)
	if err == nil {
		return quotation, nil
	}

	// if not in cache, get quotation from influx
	var query string
	query = fmt.Sprintf("SELECT price,priceYesterday,volumeYesterday FROM %s WHERE address='%s' AND blockchain='%s' ORDER BY DESC LIMIT 1", influxDBAssetQuotationsTable, asset.Address, asset.Blockchain)

	res, err := queryInfluxDB(db.influxClient, query)
	if err != nil {
		return quotation, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {

		quotation.Time, err = time.Parse(time.RFC3339, res[0].Series[0].Values[0][0].(string))
		if err != nil {
			return quotation, err
		}
		quotation.Price, err = res[0].Series[0].Values[0][1].(json.Number).Float64()
		if err != nil {
			return quotation, err
		}
		quotation.PriceYesterday, err = res[0].Series[0].Values[0][2].(json.Number).Float64()
		if err != nil {
			return quotation, err
		}
		quotation.VolumeYesterdayUSD, err = res[0].Series[0].Values[0][3].(json.Number).Float64()
		if err != nil {
			log.Errorf("no 24h volume available for %s: %v", asset.Symbol, err)
		}
	} else {
		return quotation, errors.New("Error parsing Trade from Database")
	}
	quotation.Asset = asset
	quotation.Source = dia.Diadata
	return quotation, nil
}

// SetAssetQuotationCache stores @quotation in redis cache
func (db *DB) SetAssetQuotationCache(quotation *AssetQuotation) error {
	key := getKeyAssetQuotation(quotation.Asset.Blockchain.Name, quotation.Asset.Address)
	return db.redisClient.Set(key, quotation, 0).Err()
}

// GetAssetQuotationCache returns the latest quotation for @asset from the redis cache.
func (db *DB) GetAssetQuotationCache(asset dia.Asset) (*AssetQuotation, error) {
	key := getKeyAssetQuotation(asset.Blockchain.Name, asset.Address)
	quotation := &AssetQuotation{}
	err := db.redisClient.Get(key).Scan(quotation)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("GetAssetQuotationCache on %s: %v\n", asset.Name, err)
		}
		return quotation, err
	}
	return quotation, nil
}

// ------------------------------------------------------------------------------
// EXCHANGE RATES (Deprecating)
// ------------------------------------------------------------------------------

func (db *DB) SetPriceUSD(symbol string, price float64) error {

	return db.SetQuotation(&Quotation{
		Symbol: symbol,
		Name:   helpers.NameForSymbol(symbol),
		Price:  price,
		Source: dia.Diadata,
		Time:   time.Now(),
	})
}

func (a *DB) SetPriceEUR(symbol string, price float64) error {
	return a.SetQuotationEUR(&Quotation{
		Symbol: symbol,
		Name:   helpers.NameForSymbol(symbol),
		Price:  price,
		Source: dia.Diadata,
		Time:   time.Now(),
	})
}

func (db *DB) GetPriceUSD(symbol string) (float64, error) {
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

func (db *DB) GetQuotation(symbol string) (*Quotation, error) {
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

func (db *DB) SetQuotation(quotation *Quotation) error {
	if db.redisClient == nil {
		return nil
	}
	key := getKeyQuotation(quotation.Symbol)
	log.Debug("setting ", key, quotation)
	err := db.redisClient.Set(key, quotation, TimeOutRedis).Err()
	if err != nil {
		log.Printf("Error: %v on SetQuotation %v\n", err, quotation.Symbol)
	}
	return err
}

func (db *DB) SetQuotationEUR(quotation *Quotation) error {
	if db.redisClient == nil {
		return nil
	}
	key := getKeyQuotationEUR(quotation.Symbol)
	log.Debug("setting ", key, quotation)
	err := db.redisClient.Set(key, quotation, TimeOutRedis).Err()
	if err != nil {
		log.Printf("Error: %v on SetQuotation %v\n", err, quotation.Symbol)
	}
	return err
}

func (db *DB) GetPaxgQuotationOunces() (*Quotation, error) {
	return db.GetQuotation("PAXG")
}

func (db *DB) GetPaxgQuotationGrams() (*Quotation, error) {
	q, err := db.GetQuotation("PAXG")
	if err != nil {
		return nil, err
	}
	q.Symbol = q.Symbol + "-gram"
	q.Name = q.Name + "-gram"
	q.Price = q.Price / 31.1034768
	*q.PriceYesterday = *q.PriceYesterday / 31.1034768
	return q, err
}
