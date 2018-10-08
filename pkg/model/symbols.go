package models

import (
	"github.com/diadata-org/api-golang/dia"
	log "github.com/sirupsen/logrus"
	"strings"
)

func (db *DB) GetSymbols(exchange string) ([]string, error) {
	var result []string
	var cursor uint64
	key := "dia_" + dia.FilterKing + "_"
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
			if exchange == "" {
				if len(s) == 1 {
					result = append(result, s[0])
				}
			} else {
				if s[1] == exchange {
					result = append(result, s[0])
				}
			}
		}
		if cursor == 0 {
			log.Info("GetSymbols %v returns %v", key, result)
			return result, nil
		}
	}
}

func (db *DB) GetSymbolExchangeDetails(symbol string, exchange string) (*SymbolExchangeDetails, error) {
	result := &SymbolExchangeDetails{
		Name: exchange,
	}
	v, err := db.GetPrice(symbol, exchange)
	if err == nil {
		result.Price = v
	}

	py, err2 := db.GetPriceYesterday(symbol, exchange)
	if err2 == nil {
		result.PriceYesterday = &py
	}

	v2, _ := db.GetVolumeExchange(symbol, exchange)
	result.VolumeYesterdayUSD = v2

	d, _ := db.GetLastTradeTimeForExchange(symbol, exchange)
	result.Time = d

	return result, err
}
