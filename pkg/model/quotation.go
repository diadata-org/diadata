package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
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

func getKeyAssetQuotation(blockchain, address string) string {
	return "dia_assetquotation_USD_" + blockchain + "_" + address
}

// ------------------------------------------------------------------------------
// ASSET EXCHANGE RATES (WIP)
// ------------------------------------------------------------------------------

// SetAssetPriceUSD stores the price of @asset in influx and the caching layer.
// The latter only holds the most recent price point.
func (datastore *DB) SetAssetPriceUSD(asset dia.Asset, price float64, timestamp time.Time) error {
	return datastore.SetAssetQuotation(&AssetQuotation{
		Asset:  asset,
		Price:  price,
		Source: dia.Diadata,
		Time:   timestamp,
	})
}

// GetAssetPriceUSDLatest returns the latest price of @asset.
func (datastore *DB) GetAssetPriceUSDLatest(asset dia.Asset) (price float64, err error) {
	assetQuotation, err := datastore.GetAssetQuotationLatest(asset)
	if err != nil {
		return
	}
	price = assetQuotation.Price
	return
}

// GetAssetPriceUSD returns the latest USD price of @asset before @timestamp.
func (datastore *DB) GetAssetPriceUSD(asset dia.Asset, timestamp time.Time) (price float64, err error) {
	assetQuotation, err := datastore.GetAssetQuotation(asset, timestamp)
	if err != nil {
		return
	}
	price = assetQuotation.Price
	return
}

// AddAssetQuotationsToBatch is a helper function that adds a slice of
// quotations to an influx batch.
func (datastore *DB) AddAssetQuotationsToBatch(quotations []*AssetQuotation) error {
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
		datastore.addPoint(pt)
	}
	return nil
}

// SetAssetQuotation stores the full quotation of @asset into influx and cache.
func (datastore *DB) SetAssetQuotation(quotation *AssetQuotation) error {
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
		datastore.addPoint(pt)
	}

	// Write latest point to redis cache
	// log.Printf("write to cache: %s", quotation.Asset.Symbol)
	_, err = datastore.SetAssetQuotationCache(quotation, false)
	return err

}

// GetAssetQuotation returns the latest full quotation for @asset.
func (datastore *DB) GetAssetQuotationLatest(asset dia.Asset) (*AssetQuotation, error) {

	// First attempt to get latest quotation from redis cache
	quotation, err := datastore.GetAssetQuotationCache(asset)
	if err == nil {
		log.Infof("got asset quotation for %s from cache: %v", asset.Symbol, quotation)
		return quotation, nil
	}

	// if not in cache, get quotation from influx
	log.Infof("asset %s not in cache. Query influx...", asset.Symbol)
	return datastore.GetAssetQuotation(asset, time.Now())

}

// GetAssetQuotation returns the latest full quotation for @asset before @timestamp.
func (datastore *DB) GetAssetQuotation(asset dia.Asset, timestamp time.Time) (*AssetQuotation, error) {

	quotation := AssetQuotation{}
	q := fmt.Sprintf("SELECT price FROM %s WHERE address='%s' AND blockchain='%s' AND time<=%d ORDER BY DESC LIMIT 1", influxDBAssetQuotationsTable, asset.Address, asset.Blockchain, timestamp.UnixNano())
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		return &quotation, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		if len(res[0].Series[0].Values) > 0 {
			quotation.Time, err = time.Parse(time.RFC3339, res[0].Series[0].Values[0][0].(string))
			if err != nil {
				return &quotation, err
			}
			quotation.Price, err = res[0].Series[0].Values[0][1].(json.Number).Float64()
			if err != nil {
				return &quotation, err
			}
			log.Infof("queried price for %s: %v", asset.Symbol, quotation.Price)
		} else {
			return &quotation, errors.New("no assetQuotation in DB")
		}
	} else {
		return &quotation, errors.New("no assetQuotation in DB")
	}
	quotation.Asset = asset
	quotation.Source = dia.Diadata
	return &quotation, nil
}

// GetAssetQuotations returns all assetQuotations for @asset in the given time-range.
func (datastore *DB) GetAssetQuotations(asset dia.Asset, starttime time.Time, endtime time.Time) ([]AssetQuotation, error) {

	quotations := []AssetQuotation{}
	q := fmt.Sprintf(
		"SELECT price FROM %s WHERE address='%s' AND blockchain='%s' AND time>%d AND time<=%d ORDER BY DESC",
		influxDBAssetQuotationsTable,
		asset.Address,
		asset.Blockchain,
		starttime.UnixNano(),
		endtime.UnixNano(),
	)

	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		return quotations, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := range res[0].Series[0].Values {
			var quotation AssetQuotation
			quotation.Time, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return quotations, err
			}
			quotation.Price, err = res[0].Series[0].Values[i][1].(json.Number).Float64()
			if err != nil {
				return quotations, err
			}
			quotation.Asset = asset
			quotation.Source = dia.Diadata
			quotations = append(quotations, quotation)
		}
	} else {
		return quotations, errors.New("no assetQuotation in DB")
	}

	return quotations, nil
}

// SetAssetQuotationCache stores @quotation in redis cache.
// If @check is true, it checks for a more recent quotation first.
func (datastore *DB) SetAssetQuotationCache(quotation *AssetQuotation, check bool) (bool, error) {
	if check {
		// fetch current state of cache
		cachestate, err := datastore.GetAssetQuotationCache(quotation.Asset)
		if err != nil && !errors.Is(err, redis.Nil) {
			return false, err
		}
		// Do not write to cache if more recent entry exists
		if (quotation.Time).Before(cachestate.Time) {
			return false, nil
		}
	}
	// Otherwise write to cache
	key := getKeyAssetQuotation(quotation.Asset.Blockchain, quotation.Asset.Address)
	return true, datastore.redisPipe.Set(key, quotation, TimeOutAssetQuotation).Err()
}

// GetAssetQuotationCache returns the latest quotation for @asset from the redis cache.
func (datastore *DB) GetAssetQuotationCache(asset dia.Asset) (*AssetQuotation, error) {
	key := getKeyAssetQuotation(asset.Blockchain, asset.Address)
	// log.Infof("get asset quotation from cache for asset %s with address %s using key as %s ", asset.Symbol, asset.Address, key)

	quotation := &AssetQuotation{}

	err := datastore.redisClient.Get(key).Scan(quotation)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Errorf("GetAssetQuotationCache on %s: %v\n", asset.Name, err)
		}
		return quotation, err
	}
	return quotation, nil
}

// GetAssetPriceUSDCache returns the latest price of @asset from the cache.
func (datastore *DB) GetAssetPriceUSDCache(asset dia.Asset) (price float64, err error) {
	quotation, err := datastore.GetAssetQuotationCache(asset)
	if err != nil {
		log.Errorf("get asset quotation for %s from cache: %v", asset.Symbol, err)
		return
	}
	price = quotation.Price
	return
}

// GetSortedQuotations returns quotations for all assets in @assets, sorted by 24h volume
// in descending order.
func (datastore *DB) GetSortedAssetQuotations(assets []dia.Asset) ([]AssetQuotation, error) {
	var quotations []AssetQuotation
	var volumes []float64
	for _, asset := range assets {
		var quotation *AssetQuotation
		var volume *float64
		var err error
		quotation, err = datastore.GetAssetQuotationLatest(asset)
		if err != nil {
			log.Errorf("get quotation for symbol %s with address %s on blockchain %s: %v", asset.Symbol, asset.Address, asset.Blockchain, err)
			continue
		}
		volume, err = datastore.Get24HoursAssetVolume(asset)
		if err != nil {
			log.Errorf("get volume for symbol %s with address %s on blockchain %s: %v", asset.Symbol, asset.Address, asset.Blockchain, err)
			continue
		}
		quotations = append(quotations, *quotation)
		volumes = append(volumes, *volume)
	}
	if len(quotations) == 0 {
		return quotations, errors.New("no quotations available")
	}

	var quotationsSorted []AssetQuotation
	volumesSorted := utils.NewFloat64Slice(sort.Float64Slice(volumes))
	sort.Sort(volumesSorted)
	for _, ind := range volumesSorted.Ind() {
		quotationsSorted = append([]AssetQuotation{quotations[ind]}, quotationsSorted...)
	}
	return quotationsSorted, nil
}

func (datastore *DB) GetOldestQuotation(asset dia.Asset) (AssetQuotation, error) {
	// TO DO: Return assetquotation with oldest timestamp.
	return AssetQuotation{}, nil
}

// ------------------------------------------------------------------------------
// HISTORICAL QUOTES
// ------------------------------------------------------------------------------

// SetHistoricalQuote stores a historical quote for an asset symbol at a specific time into postgres.
func (rdb *RelDB) SetHistoricalQuotation(quotation AssetQuotation) error {
	// TO DO
	return nil
}

func (rdb *RelDB) GetHistoricalQuotation(asset dia.Asset, timestamp time.Time) (quotation AssetQuotation, err error) {
	// TO DO
	return
}

// GetLastHistoricalQuoteTimestamp returns the timestamp of the last historical quote for asset symbol.
func (rdb *RelDB) GetLastHistoricalQuotationTimestamp(asset dia.Asset) (time.Time, error) {
	// TO DO
	return time.Now(), nil
}

// ------------------------------------------------------------------------------
// MARKET MEASURES
// ------------------------------------------------------------------------------

// GetAssetsMarketCap returns the actual market cap of @asset.
func (datastore *DB) GetAssetsMarketCap(asset dia.Asset) (float64, error) {
	price, err := datastore.GetAssetPriceUSDLatest(asset)
	if err != nil {
		return 0, err
	}
	supply, err := datastore.GetSupplyCache(asset)
	if err != nil {
		return 0, err
	}
	return price * supply.CirculatingSupply, nil
}

// GetTopAssetByVolume returns the asset with highest volume among all assets with symbol @symbol.
// This method allows us to use all API endpoints called on a symbol.
func (datastore *DB) GetTopAssetByVolume(symbol string, relDB *RelDB) (topAsset dia.Asset, err error) {
	assets, err := relDB.GetAssets(symbol)
	if err != nil {
		return
	}
	if len(assets) == 0 {
		err = errors.New("no matching asset")
		return
	}
	var volume float64
	for _, asset := range assets {
		var value *float64
		value, err = datastore.Get24HoursAssetVolume(asset)
		if err != nil {
			log.Error(err)
			continue
		}
		if value == nil {
			continue
		}
		if *value > volume {
			volume = *value
			topAsset = asset
		}
	}
	if volume == 0 {
		err = errors.New("no quotation for symbol")
	} else {
		err = nil
	}
	return
}

// GetTopAssetByMcap returns the asset with highest market cap among all assets with symbol @symbol.
func (datastore *DB) GetTopAssetByMcap(symbol string, relDB *RelDB) (topAsset dia.Asset, err error) {
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
		value, err = datastore.GetAssetsMarketCap(asset)
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
