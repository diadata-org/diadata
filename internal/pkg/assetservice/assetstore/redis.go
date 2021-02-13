package assetstore

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/go-redis/redis"
)

func getKeyAsset(symbol, name string) (key string) {
	key = "dia_asset_" + symbol + "_" + name
	return
}

// CacheSetAsset stores @asset in redis
func (rdb *RelDB) CacheSetAsset(asset dia.Asset) error {
	return rdb.redisClient.Set(getKeyAsset(asset.Symbol, asset.Name), &asset, 0).Err()
}

// CacheGetAsset returns an asset by name and symbol
func (rdb *RelDB) CacheGetAsset(symbol, name string) (dia.Asset, error) {
	asset := dia.Asset{}
	err := rdb.redisClient.Get(getKeyAsset(symbol, name)).Scan(&asset)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("Error: %v on GetAsset %v\n", err, symbol)
		}
		return asset, err
	}
	return asset, nil
}

// CacheCount returns the number of assets in the cache
func (rdb *RelDB) CacheCount() (uint32, error) {
	keysPattern := "dia_asset_*"
	allAssets := rdb.redisClient.Keys(keysPattern).Val()
	return uint32(len(allAssets)), nil
}

// CacheSetAvailablePairs stores @pairs in redis
func (rdb *RelDB) CacheSetAvailablePairs(exchange string, pairs []dia.Pair) error {
	key := "dia_available_pairs_" + exchange
	var p dia.Pairs = pairs
	return rdb.redisClient.Set(key, &p, 0).Err()
}
