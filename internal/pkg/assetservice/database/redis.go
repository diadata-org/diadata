package database

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Assetcache simply is a redis client. This can be moved into our db at some point
type Assetcache struct {
	redisClient *redis.Client
}

// NewAssetcache returns a redis client.
func NewAssetcache() *Assetcache {

	address := "localhost:6379"
	r := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong2, err := r.Ping().Result()
	if err != nil {
		log.Error("NewDataStore redis", err)
	}
	log.Debug("NewDB", pong2)
	return &Assetcache{redisClient: r}

}

func getKeyAsset(symbol, name string) (key string) {
	key = "dia_asset_" + symbol + "_" + name
	return
}

// SetAsset stores marshalled @asset in redis
func (ac *Assetcache) SetAsset(asset dia.Asset) error {
	return ac.redisClient.Set(getKeyAsset(asset.Symbol, asset.Name), &asset, 0).Err()
}

// GetAsset returns an asset by name and symbol
func (ac *Assetcache) GetAsset(name, symbol string) (dia.Asset, error) {
	asset := dia.Asset{}
	err := ac.redisClient.Get(getKeyAsset(name, symbol)).Scan(&asset)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("Error: %v on GetInterestRate %v\n", err, symbol)
		}
		return asset, err
	}
	return asset, nil
}
