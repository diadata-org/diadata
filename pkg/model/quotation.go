package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/go-redis/redis"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
)

const (
	WindowYesterday       = 24 * 60 * 60
	Window1h              = 60 * 60
	Window7d              = 7 * 24 * 60 * 60
	Window14d             = 7 * 24 * 60 * 60
	Window30d             = 30 * 24 * 60 * 60
	Window2               = 24 * 60 * 60 * 8
	BufferTTL             = 60 * 60
	BiggestWindow         = Window2
	TimeOutRedis          = time.Duration(time.Second*BiggestWindow + time.Second*BufferTTL)
	TimeOutAssetQuotation = time.Duration(time.Second * WindowYesterday)
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
	assetQuotation, err := db.GetAssetQuotation(asset)
	if err != nil {
		return
	}
	price = assetQuotation.Price
	return
}

// AddAssetQuotationsToBatch is a helper function that adds a slice of
// quotations to an influx batch.
func (db *DB) AddAssetQuotationsToBatch(quotations []*AssetQuotation) error {
	for _, quotation := range quotations {
		tags := map[string]string{
			"symbol":     quotation.Asset.Symbol,
			"name":       quotation.Asset.Name,
			"address":    quotation.Asset.Address,
			"blockchain": quotation.Asset.Blockchain,
		}
		fields := map[string]interface{}{
			"price": quotation.Price,
		}
		pt, err := clientInfluxdb.NewPoint(influxDBAssetQuotationsTable, tags, fields, quotation.Time)
		if err != nil {
			log.Errorln("addAssetQuotationsToBatch:", err)
			return err
		}
		db.addPoint(pt)
	}
	return nil
}

// SetAssetQuotation stores the full quotation of @asset into influx and cache.
func (db *DB) SetAssetQuotation(quotation *AssetQuotation) error {
	// Write to influx
	tags := map[string]string{
		"symbol":     quotation.Asset.Symbol,
		"name":       quotation.Asset.Name,
		"address":    quotation.Asset.Address,
		"blockchain": quotation.Asset.Blockchain,
	}
	fields := map[string]interface{}{
		"price": quotation.Price,
	}

	pt, err := clientInfluxdb.NewPoint(influxDBAssetQuotationsTable, tags, fields, quotation.Time)
	if err != nil {
		log.Errorln("SetAssetQuotation:", err)
	} else {
		db.addPoint(pt)
	}

	// Write latest point to redis cache
	log.Printf("write to cache: %s", quotation.Asset.Symbol)
	_, err = db.SetAssetQuotationCache(quotation)
	return err

}

// GetAssetQuotation returns the latest full quotation for @asset.
func (db *DB) GetAssetQuotation(asset dia.Asset) (*AssetQuotation, error) {

	// First attempt to get latest quotation from redis cache
	quotation, err := db.GetAssetQuotationCache(asset)
	if err == nil {
		log.Infof("got asset quotation for %s from cache: %v", asset.Symbol, quotation)
		return quotation, nil
	}

	// if not in cache, get quotation from influx
	log.Infof("asset %s not in cache. Query influx...", asset.Symbol)
	q := fmt.Sprintf("SELECT price FROM %s WHERE address='%s' AND blockchain='%s' ORDER BY DESC LIMIT 1", influxDBAssetQuotationsTable, asset.Address, asset.Blockchain)

	res, err := queryInfluxDB(db.influxClient, q)
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
		log.Infof("queried price for %s: %v", asset.Symbol, quotation.Price)
	} else {
		return quotation, errors.New("no assetQuotation in influx")
	}
	quotation.Asset = asset
	quotation.Source = dia.Diadata
	return quotation, nil
}

// SetAssetQuotationCache stores @quotation in redis cache
func (db *DB) SetAssetQuotationCache(quotation *AssetQuotation) (bool, error) {
	// fetch current state of cache
	cachestate, err := db.GetAssetQuotationCache(quotation.Asset)
	if err != nil && !errors.Is(err, redis.Nil) {
		return false, err
	}
	// Do not write to cache if more recent entry exists
	if (quotation.Time).Before(cachestate.Time) {
		return false, nil
	}
	// Otherwise write to cache
	key := getKeyAssetQuotation(quotation.Asset.Blockchain, quotation.Asset.Address)
	return true, db.redisClient.Set(key, quotation, TimeOutAssetQuotation).Err()
}

// GetAssetQuotationCache returns the latest quotation for @asset from the redis cache.
func (db *DB) GetAssetQuotationCache(asset dia.Asset) (*AssetQuotation, error) {
	log.Infof("get asset quotation from cache for asset %s with address %s ", asset.Symbol, asset.Address)
	key := getKeyAssetQuotation(asset.Blockchain, asset.Address)
	quotation := &AssetQuotation{}
	err := db.redisClient.Get(key).Scan(quotation)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Errorf("GetAssetQuotationCache on %s: %v\n", asset.Name, err)
		}
		return quotation, err
	}
	return quotation, nil
}

// GetAssetPriceUSDCache returns the last price of @asset from the cache.
func (db *DB) GetAssetPriceUSDCache(asset dia.Asset) (price float64, err error) {
	quotation, err := db.GetAssetQuotationCache(asset)
	if err != nil {
		log.Errorf("get asset quotation for %s from cache: %v", asset.Symbol, err)
		return
	}
	price = quotation.Price
	return
}

// GetSortedQuotations returns quotations for all assets in @assets, sorted by 24h volume
// in descending order.
func (db *DB) GetSortedAssetQuotations(assets []dia.Asset) ([]AssetQuotation, error) {
	var quotations []AssetQuotation
	var volumes []float64
	for _, asset := range assets {
		var quotation *AssetQuotation
		var volume *float64
		var err error
		quotation, err = db.GetAssetQuotation(asset)
		if err != nil {
			log.Errorf("get quotation for symbol %s with address %s on blockchain %s: %v", asset.Symbol, asset.Address, asset.Blockchain, err)
			continue
		}
		volume, err = db.GetVolume(asset)
		if err != nil {
			log.Error("get volume for symbol %s with address %s on blockchain %s: %v", asset.Symbol, asset.Address, asset.Blockchain, err)
			continue
		}
		quotations = append(quotations, *quotation)
		volumes = append(volumes, *volume)
	}
	if len(quotations) == 0 {
		return quotations, nil
	}

	var quotationsSorted []AssetQuotation
	volumesSorted := utils.NewFloat64Slice(sort.Float64Slice(volumes))
	sort.Sort(volumesSorted)
	for _, ind := range volumesSorted.Ind() {
		quotationsSorted = append([]AssetQuotation{quotations[ind]}, quotationsSorted...)
	}
	return quotationsSorted, nil
}

// ------------------------------------------------------------------------------
// MARKET MEASURES
// ------------------------------------------------------------------------------

// GetAssetsMarketCap returns the actual market cap of @asset.
func (db *DB) GetAssetsMarketCap(asset dia.Asset) (float64, error) {
	price, err := db.GetAssetPriceUSD(asset)
	if err != nil {
		return 0, err
	}
	supply, err := db.GetSupplyInflux(asset, time.Time{}, time.Time{})
	if err != nil {
		return 0, err
	}
	if len(supply) > 0 {
		return price * supply[0].CirculatingSupply, nil
	}
	return 0, errors.New("no circulating supply available")
}

// GetTopAsset returns the asset with highest market cap among all assets with symbol @symbol.
// This method allows us to use all API endpoints called on a symbol.
func (db *DB) GetTopAsset(symbol string, relDB *RelDB) (topAsset dia.Asset, err error) {
	assets, err := relDB.GetAssets(symbol)
	if err != nil {
		return
	}
	if len(assets) == 0 {
		err = errors.New("no matching asset")
		return
	}
	var mcap float64
	for _, asset := range assets {
		var value float64
		value, err = db.GetAssetsMarketCap(asset)
		if err != nil {
			log.Error(err)
			continue
		}
		if value == 0 {
			continue
		}
		if value > mcap {
			mcap = value
			topAsset = asset
		}
	}
	if mcap == 0 {
		err = errors.New("no quotation for symbol")
	} else {
		err = nil
	}
	return
}

// ------------------------------------------------------------------------------
// GOLD Derivatives
// ------------------------------------------------------------------------------

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
		if !errors.Is(err, redis.Nil) {
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
		if !errors.Is(err, redis.Nil) {
			log.Errorf("Error: %v on GetQuotation %v\n", err, key)
		}
		return nil, err
	}
	value.Name = helpers.NameForSymbol(symbol) // in case we updated the helper functions ;)
	preliminaryAsset := dia.Asset{
		Symbol: symbol,
	}
	v, err2 := db.GetPriceYesterday(preliminaryAsset, "")
	if err2 == nil {
		value.PriceYesterday = &v
	}
	// v2, err2 := db.GetVolume(symbol)
	// value.VolumeYesterdayUSD = v2
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
