package assetstore

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// AssetCache simply is a redis client. This can be moved into our db at some point
type AssetCache struct {
	redisClient *redis.Client
	pagesize    uint32
}

// NewAssetCache returns a redis client.
func NewAssetCache() (*AssetCache, error) {

	address := "localhost:6379"
	r := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong2, err := r.Ping().Result()
	if err != nil {
		log.Error("NewDataStore redis", err)
		return &AssetCache{}, err
	}
	log.Debug("NewDB", pong2)
	return &AssetCache{redisClient: r}, nil

}

func getKeyAsset(symbol, name string) (key string) {
	key = "dia_asset_" + symbol + "_" + name
	return
}

// SetAsset stores @asset in redis
func (ac *AssetCache) SetAsset(asset dia.Asset) error {
	return ac.redisClient.Set(getKeyAsset(asset.Symbol, asset.Name), &asset, 0).Err()
}

// GetAsset returns an asset by name and symbol
func (ac *AssetCache) GetAsset(symbol, name string) (dia.Asset, error) {
	asset := dia.Asset{}
	err := ac.redisClient.Get(getKeyAsset(symbol, name)).Scan(&asset)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("Error: %v on GetAsset %v\n", err, symbol)
		}
		return asset, err
	}
	return asset, nil
}

// GetPage returns
func (ac *AssetCache) GetPage(pageNumber uint32) (assets []dia.Asset, hasNextPage bool, err error) {
	// TO DO
	return
}

// Count returns the number of assets in the cache
func (ac *AssetCache) Count() (uint32, error) {
	keysPattern := "dia_asset_*"
	allAssets := ac.redisClient.Keys(keysPattern).Val()
	return uint32(len(allAssets)), nil
}
