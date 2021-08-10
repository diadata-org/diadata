package db

import (
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func GetRedisClient() *redis.Client {
	var redisClient *redis.Client

	// This environment variable is either set in docker-compose or empty
	address := utils.Getenv("REDISURL", "localhost:6379")
	password := utils.Getenv("REDISPASSWORD", "")
	defaultDB, err := strconv.Atoi(utils.Getenv("REDISUSEDEFAULTDB", "0"))
	if err != nil {
		log.Error("wrong value for redis default db", err)
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,  // no password set
		DB:       defaultDB, // use default DB
	})

	pong2, err := redisClient.Ping().Result()
	if err != nil {
		log.Error("NewDataStore redis: ", err)
	}

	log.Debug("NewDB", pong2)

	return redisClient
}
