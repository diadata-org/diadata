package models

import (
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const (
	WindowVolume = 60 * 60 * 24
)

func getKeyVolume(value string) string {
	return "dia_volume_" + value + "_ZSET"
}

func (db *DB) SetVolume(symbol string, volume float64) error {
	key := getKeyVolume(symbol)

	timeUnix := time.Now().Unix()

	member := strconv.FormatFloat(volume, 'f', -1, 64) + " " + strconv.FormatInt(timeUnix, 10)

	err := db.redisClient.ZAdd(key, redis.Z{
		Score:  float64(timeUnix),
		Member: member,
	}).Err()

	log.Println("SetVolume ", symbol, member, timeUnix)

	if err != nil {
		log.Printf("Error: %v on SetVolume %v\n", err, symbol)
	}
	// purging old values
	err = db.redisClient.ZRemRangeByScore(key, "-inf", "("+strconv.FormatInt(timeUnix-WindowVolume, 10)).Err()
	if err != nil {
		log.Printf("Error: %v on SetVolume %v\n", err, symbol)
	}
	return err
}

func (db *DB) GetVolume(symbol string) (float64, error) {
	result := 0.0
	key := getKeyVolume(symbol)
	vals, err := db.redisClient.ZRange(key, 0, -1).Result()
	if err != nil {
		log.Printf("Error: %v on GetVolume %v\n", err, symbol)
	} else {
		for _, v := range vals {
			f := 0.0
			fmt.Sscanf(v, "%f", &f)
			result += f
		}
	}
	return result, err
}
