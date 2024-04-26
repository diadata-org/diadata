package db

import (
	"context"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/redis/go-redis/v9"
)

func GetRedisClient() *redis.ClusterClient {
	var redisClusterClient *redis.ClusterClient

	// This environment variable is either set in docker-compose or empty
	password := utils.Getenv("REDISPASSWORD", "")
	addresses := strings.Split(utils.Getenv("REDIS_ADDRS", ""), ",")

	redisClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Password:    password,
		Addrs:       addresses,
		ReadTimeout: time.Duration(120 * time.Second),
	})

	pong, err := redisClusterClient.Ping(context.Background()).Result()
	if err != nil {
		log.Error("NewDataStore redis: ", err)
	}

	log.Debug("redisFailoverClient", pong)

	return redisClusterClient

}
