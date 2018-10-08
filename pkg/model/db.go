package models

import (
	"errors"
	"fmt"
	"github.com/diadata-org/api-golang/dia"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Datastore interface {
	SetVolume(symbol string, exchange string, volume float64) error
	GetVolume(symbol string) (*float64, error)
	SymbolsWithASupply() ([]string, error)
	SetPriceUSD(symbol string, price float64) error
	SetPriceEUR(symbol string, price float64) error
	GetPriceUSD(symbol string) (float64, error)
	GetQuotation(symbol string) (*Quotation, error)
	SetQuotation(quotation *Quotation) error
	SetQuotationEUR(quotation *Quotation) error
	GetSupply(symbol string) (*dia.Supply, error)
	SetSupply(supply *dia.Supply) error
	SetPriceZSET(symbol string, exchange string, price float64) error
	GetChartPoints(symbol string) ([]Point, error)
	GetPairs(exchange string) ([]dia.Pair, error)
	GetSymbols(exchange string) ([]string, error)
	GetExchangesForSymbol(symbol string) ([]string, error)
	GetSymbolExchangeDetails(symbol string, exchange string) (*SymbolExchangeDetails, error)
	GetLastTradeTimeForExchange(symbol string, exchange string) (*time.Time, error)
	SetLastTradeTimeForExchange(symbol string, exchange string, t time.Time) error
}

type DB struct {
	redisClient *redis.Client
}

func NewDataStore() (*DB, error) {

	db := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong2, err := db.Ping().Result()
	if err != nil {
		log.Error("NewDataStore", err)
	}
	log.Debug("NewDB", pong2)

	return &DB{db}, nil
}

func getKey(filter string, symbol string, exchange string) string {
	key := filter + "_" + symbol
	if exchange != "" {
		key = key + "_" + exchange
	}
	return key
}

func getKeyFilterZSET(key string) string {
	return "dia_" + key + "_ZSET"
}

func getKeyFilterSymbolAndExchangeZSET(filter string, symbol string, exchange string) string {
	if exchange == "" {
		return "dia_" + filter + "_" + symbol + "_ZSET"
	} else {
		return "dia_" + filter + "_" + symbol + "_" + exchange + "_ZSET"
	}
}

func (db *DB) setZSETValue(key string, value float64, unixTime int64, maxWindow int64) error {
	member := strconv.FormatFloat(value, 'f', -1, 64) + " " + strconv.FormatInt(unixTime, 10)
	err := db.redisClient.ZAdd(key, redis.Z{
		Score:  float64(unixTime),
		Member: member,
	}).Err()
	log.Debug("SetZSETValue ", key, member, unixTime)
	if err != nil {
		log.Error("Error: %v on SetZSETValue %v\n", err, key)
	}
	// purging old values
	err = db.redisClient.ZRemRangeByScore(key, "-inf", "("+strconv.FormatInt(unixTime-maxWindow, 10)).Err()
	if err != nil {
		log.Error("Error: %v on SetZSETValue %v\n", err, key)
	}

	if err := db.redisClient.Expire(key, TimeOutRedis).Err(); err != nil {
		log.Error(err)
	} //TODO put two commands together ?
	return err
}

func (db *DB) getZSETValue(key string, atUnixTime int64) (float64, error) {
	result := 0.0
	max := strconv.FormatInt(atUnixTime, 10)
	vals, err := db.redisClient.ZRangeByScoreWithScores(key, redis.ZRangeBy{
		Min: "-inf",
		Max: max,
	}).Result()
	log.Debug(key, "vals: %v on getZSETValue maxScore: %v", vals, max)
	if err == nil {
		if len(vals) > 0 {
			fmt.Sscanf(vals[len(vals)-1].Member.(string), "%f", &result)
			log.Debug("returned value: %v", result)
		} else {
			err = errors.New("getZSETValue no value found")
		}
	}
	return result, err
}

func (db *DB) getZSETSum(key string) (*float64, error) {

	log.Printf("getZSETSum: %v \n", key)

	vals, err := db.redisClient.ZRange(key, 0, -1).Result()
	if err != nil {
		log.Error("Error: %v on getZSETSum %v\n", err, key)
		return nil, err
	} else {
		result := 0.0
		for _, v := range vals {
			f := 0.0
			fmt.Sscanf(v, "%f", &f)
			result += f
		}
		return &result, err
	}
}

func (db *DB) getZSETLastValue(key string) (float64, int64, error) {
	value := 0.0
	var unixTime int64
	vals, err := db.redisClient.ZRange(key, -1, -1).Result()
	log.Println(key, "on getZSETLastValue:", vals)
	if err == nil {
		if len(vals) == 1 {
			fmt.Sscanf(vals[0], "%f %d", &value, &unixTime)
			log.Debug("returned value: %v", value)
		} else {
			err = errors.New("getZSETLastValue no value found")
			log.Errorln("Error: on getZSETLastValue", err, key)
		}
	}
	return value, unixTime, err
}
