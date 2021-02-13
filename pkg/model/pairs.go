package models

import (
	"strings"

	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
)

// Pair substitues the old dia.Pair. It includes the new asset type.
type Pair struct {
	BaseToken  dia.Asset
	QuoteToken dia.Asset
	Exchange   dia.Exchange
}

// ForeignName returns the foreign name of the pair @p, i.e. the string Quotetoken-Basetoken
func (p *Pair) ForeignName() string {
	return p.QuoteToken.Symbol + "-" + p.BaseToken.Symbol
}

// exchange = "" for all exchanges
func (db *DB) GetPairs(exchange string) ([]dia.Pair, error) {
	var result []dia.Pair
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
			if len(s) == 2 {
				result = append(result, dia.Pair{
					Symbol:   s[0],
					Exchange: s[1],
				})
			}
		}
		if cursor == 0 {
			log.Debugf("GetPairs %v returns %v", key, result)
			return result, nil
		}
	}
}
