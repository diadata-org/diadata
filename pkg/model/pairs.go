package models

import (
	"github.com/diadata-org/diadata/pkg/dia"
)

// GetPairs returns all exchangepairs on @exchange.
// For exchange=="" it returns all exchangepairs across exchanges.
func (rdb *RelDB) GetPairs(exchange string) (pairs []dia.ExchangePair, err error) {
	if exchange == "" {
		for _, exch := range dia.Exchanges() {
			var exchangepairs []dia.ExchangePair
			exchangepairs, err = rdb.GetExchangePairSymbols(exch)
			if err != nil {
				log.Error(err)
			}
			pairs = append(pairs, exchangepairs...)
		}
	} else {
		pairs, err = rdb.GetExchangePairSymbols(exchange)
		if err != nil {
			log.Error(err)
		}
	}
	return
}

// // Deprecating
// // exchange = "" for all exchanges
// func (db *DB) GetPairs(exchange string) ([]dia.ExchangePair, error) {
// 	var result []dia.ExchangePair
// 	var cursor uint64
// 	key := "dia_" + dia.FilterKing + "_"
// 	for {
// 		var keys []string
// 		var err error
// 		keys, cursor, err = db.redisClient.Scan(cursor, key+"*", 10).Result()
// 		if err != nil {
// 			log.Error("GetPairs err", err)
// 			return result, err
// 		}
// 		for _, value := range keys {
// 			filteredKey := strings.Replace(strings.Replace(value, key, "", 1), "_ZSET", "", 1)
// 			s := strings.Split(strings.Replace(filteredKey, key, "", 1), "_")
// 			if len(s) == 2 {
// 				result = append(result, dia.ExchangePair{
// 					Symbol:   s[0],
// 					Exchange: s[1],
// 				})
// 			}
// 		}
// 		if cursor == 0 {
// 			log.Debugf("GetPairs %v returns %v", key, result)
// 			return result, nil
// 		}
// 	}
// }
