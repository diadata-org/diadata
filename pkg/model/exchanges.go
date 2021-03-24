package models

import (
	// "encoding/json"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
)

// GetExchanges returns all available trading places.
// Comment: Think about getting the exchanges from redis.
func (db *DB) GetExchanges() (allExchanges []string) {
	listExch := dia.Exchanges()
	for _, exchange := range listExch {
		if exchange != "Unknown" {
			allExchanges = append(allExchanges, exchange)
		}
	}
	return
}

func getKeyLastTradeTimeForExchange(asset dia.Asset, exchange string) string {
	if exchange == "" {
		return "dia_TLT_" + asset.Blockchain + "_" + asset.Address

	} else {
		return "dia_TLT_" + asset.Blockchain + "_" + asset.Address + "_" + exchange
	}
}

func (db *DB) GetLastTradeTimeForExchange(asset dia.Asset, exchange string) (*time.Time, error) {
	key := getKeyLastTradeTimeForExchange(asset, exchange)
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

func (db *DB) SetLastTradeTimeForExchange(asset dia.Asset, exchange string, t time.Time) error {
	if db.redisClient == nil {
		return nil
	}
	key := getKeyLastTradeTimeForExchange(asset, exchange)
	log.Debug("setting ", key, t)
	err := db.redisClient.Set(key, t.Unix(), TimeOutRedis).Err()
	if err != nil {
		log.Printf("Error: %v on SetLastTradeTimeForExchange %v\n", err, asset.Symbol)
	}
	return err
}

func (rdb *RelDB) GetExchangesForSymbol(symbol string) (exchanges []string, err error) {

	query := fmt.Sprintf("select distinct(exchange) from %s where symbol=$1", exchangesymbolTable)
	rows, err := rdb.postgresClient.Query(context.Background(), query, symbol)
	if err != nil {
		return
	}
	for rows.Next() {
		exchange := ""
		rows.Scan(&exchange)
		exchanges = append(exchanges, exchange)
	}
	return
}

// Deprecating
// func (db *DB) GetExchangesForSymbol(symbol string) ([]string, error) { // TOFIX. use influx db trades on 24 hours
// 	var result []string
// 	var cursor uint64
// 	key := "dia_" + dia.FilterKing + "_" + symbol
// 	for {
// 		var keys []string
// 		var err error
// 		keys, cursor, err = db.redisClient.Scan(cursor, key+"*", 15).Result()
// 		log.Debug("GetExchangesForSymbol ", key+"*", cursor)
// 		if err != nil {
// 			log.Error("GetPairs err", err)
// 			return result, err
// 		}
// 		for _, value := range keys {
// 			log.Debug("GetExchangesForSymbol ", value)
// 			filteredKey := strings.Replace(strings.Replace(value, key, "", 1), "_ZSET", "", 1)
// 			s := strings.Split(strings.Replace(filteredKey, key, "", 1), "_")
// 			if len(s) == 2 {
// 				result = append(result, s[1])
// 			}
// 		}
// 		if cursor == 0 {
// 			log.Debugf("GetExchangesForSymbol %v returns %v", key, result)
// 			return result, nil
// 		}
// 	}
// }

// SetAvailablePairs stores @pairs in redis
// TO DO: Setter and getter should act on RelDB
func (db *DB) SetAvailablePairs(exchange string, pairs []dia.ExchangePair) error {
	key := "dia_available_pairs_" + exchange
	var p dia.Pairs = pairs
	return db.redisClient.Set(key, &p, 0).Err()
}

// GetAvailablePairs a slice of all pairs available in the exchange in the internal redis db
func (db *DB) GetAvailablePairs(exchange string) ([]dia.ExchangePair, error) {
	key := "dia_available_pairs_" + exchange
	p := dia.Pairs{}
	err := db.redisClient.Get(key).Scan(&p)
	if err != nil {
		log.Errorf("Error: %v on GetAvailablePairs %v\n", err, exchange)
		return nil, err
	}
	return p, nil
}
