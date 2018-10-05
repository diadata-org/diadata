package models

import (
	"github.com/diadata-org/api-golang/dia"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

func getKeyLastTradeTimeForExchange(symbol string, exchange string) string {
	if exchange == "" {
		return "dia_TLT_" + symbol

	} else {
		return "dia_TLT_" + symbol + "_" + exchange
	}
}

func (db *DB) GetLastTradeTimeForExchange(symbol string, exchange string) (*time.Time, error) {
	key := getKeyLastTradeTimeForExchange(symbol, exchange)
	t, err := db.redisClient.Get(key).Result()
	if err != nil {
		log.Errorln("Error: on GetLastTradeTimeForExchange", err, key)
		return nil, err
	}
	i64, err := strconv.ParseInt(t, 10, 64)
	if err == nil {
		t2 := time.Unix(i64, 0)
		return &t2, nil
	} else {
		return nil, err
	}
}

func (db *DB) SetLastTradeTimeForExchange(symbol string, exchange string, t time.Time) error {
	key := getKeyLastTradeTimeForExchange(symbol, exchange)
	log.Println("setting ", key, t)
	err := db.redisClient.Set(key, t.Unix(), TimeOutRedis).Err()
	if err != nil {
		log.Printf("Error: %v on SetLastTradeTimeForExchange %v\n", err, symbol)
	}
	return err
}

func (db *DB) GetExchangesForSymbol(symbol string) ([]string, error) {
	var result []string
	var cursor uint64
	key := "dia_" + dia.FilterKing + "_" + symbol
	for {
		var keys []string
		var err error
		keys, cursor, err = db.redisClient.Scan(cursor, key+"*", 10).Result()
		if err != nil {
			log.Error("GetPairs err", err)
			return result, err
		}
		for _, value := range keys {
			filteredKey := strings.Replace(strings.Replace(value, key, "", 1), "_ZSET", "", 1)
			s := strings.Split(strings.Replace(filteredKey, key, "", 1), "_")
			if len(s) == 2 {
				result = append(result, s[1])
			}
		}
		if cursor == 0 {
			log.Info("GetExchangesForSymbol %v returns %v", key, result)
			return result, nil
		}
	}
}
