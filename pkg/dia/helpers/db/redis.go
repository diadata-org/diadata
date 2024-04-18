package db

import (
	"context"

	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/redis/go-redis/v9"
)

func GetRedisClient() *redis.ClusterClient {
	var redisClusterClient *redis.ClusterClient

	// This environment variable is either set in docker-compose or empty
	password := utils.Getenv("REDISPASSWORD", "")

	redisClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Password: password,
		// Addrs:    []string{},
	})

	pong, err := redisClusterClient.Ping(context.Background()).Result()
	if err != nil {
		log.Error("NewDataStore redis: ", err)
	}

	log.Debug("redisFailoverClient", pong)

	return redisClusterClient

}
