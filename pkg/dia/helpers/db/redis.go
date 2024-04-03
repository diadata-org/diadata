package db

import (
	"context"
	"strconv"

	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/redis/go-redis/v9"
)

func GetRedisClient() *redis.Client {
	var redisFailoverClient *redis.Client

	// This environment variable is either set in docker-compose or empty
	masterName := utils.Getenv("REDIS_MASTER_NAME", "redis")
	address := utils.Getenv("REDISURL", "localhost:26379")
	password := utils.Getenv("REDISPASSWORD", "")
	defaultDB, err := strconv.Atoi(utils.Getenv("REDISUSEDEFAULTDB", "0"))
	if err != nil {
		log.Error("wrong value for redis default db", err)
	}

	redisFailoverClient = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:       masterName,
		Password:         password,
		SentinelAddrs:    []string{address},
		DB:               defaultDB,
		SentinelPassword: password,
	})

	pong, err := redisFailoverClient.Ping(context.Background()).Result()
	if err != nil {
		log.Error("NewDataStore redis: ", err)
	}

	log.Debug("redisFailoverClient", pong)

	return redisFailoverClient

}
