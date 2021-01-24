package database

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type Assetcache struct {
	redisClient *redis.Client
}

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

func (ac *Assetcache) SetAsset(asset dia.Asset) error {
	key := "dia_asset_" + asset.Symbol + "_" + asset.Name
	return ac.redisClient.Set(key, &asset, 0).Err()
}

// GetAvailablePairsForExchange a slice of all pairs available in the exchange in the internal redis db
