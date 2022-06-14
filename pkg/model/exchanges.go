package models

import (
	// "encoding/json"
	"context"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
)

// GetExchanges returns all available trading places.
// TO DO: Switch to postgres method asap.
func (datastore *DB) GetExchanges() (allExchanges []string) {
	listExch := dia.Exchanges()
	for _, exchange := range listExch {
		if exchange != "Unknown" {
			allExchanges = append(allExchanges, exchange)
		}
	}
	sort.Strings(allExchanges)
	return
}

func getKeyLastTradeTimeForExchange(asset dia.Asset, exchange string) string {
	if exchange == "" {
		return "dia_TLT_" + asset.Blockchain + "_" + asset.Address

	} else {
		return "dia_TLT_" + asset.Blockchain + "_" + asset.Address + "_" + exchange
	}
}

func (datastore *DB) GetLastTradeTimeForExchange(asset dia.Asset, exchange string) (*time.Time, error) {
	key := getKeyLastTradeTimeForExchange(asset, exchange)
	t, err := datastore.redisClient.Get(key).Result()
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

func (datastore *DB) SetLastTradeTimeForExchange(asset dia.Asset, exchange string, t time.Time) error {
	if datastore.redisClient == nil {
		return nil
	}
	key := getKeyLastTradeTimeForExchange(asset, exchange)
	log.Debug("setting ", key, t)
	err := datastore.redisPipe.Set(key, t.Unix(), TimeOutRedis).Err()
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
		err = rows.Scan(&exchange)
		if err != nil {
			return []string{}, err
		}
		exchanges = append(exchanges, exchange)
	}
	return
}

// SetAvailablePairs stores @pairs in redis
// TO DO: Setter and getter should act on RelDB
func (datastore *DB) SetAvailablePairs(exchange string, pairs []dia.ExchangePair) error {
	key := "dia_available_pairs_" + exchange
	var p dia.Pairs = pairs
	return datastore.redisClient.Set(key, &p, 0).Err()
}

// GetAvailablePairs a slice of all pairs available in the exchange in the internal redis db
func (datastore *DB) GetAvailablePairs(exchange string) ([]dia.ExchangePair, error) {
	key := "dia_available_pairs_" + exchange
	p := dia.Pairs{}
	err := datastore.redisClient.Get(key).Scan(&p)
	if err != nil {
		log.Errorf("Error: %v on GetAvailablePairs %v\n", err, exchange)
		return nil, err
	}
	return p, nil
}
