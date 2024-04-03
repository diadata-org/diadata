package models

import (
	"context"
	"strings"

	"github.com/diadata-org/diadata/pkg/dia"
)

func (datastore *DB) GetSymbols(exchange string) ([]string, error) {
	var result []string
	var cursor uint64
	key := "dia_" + dia.FilterKing + "_"
	for {
		var keys []string
		var err error
		keys, cursor, err = datastore.redisClient.Scan(context.Background(), cursor, key+"*", 10).Result()
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
			log.Debugf("GetSymbols %v returns %v", key, result)
			return result, nil
		}
	}
}
