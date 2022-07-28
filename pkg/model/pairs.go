package models

import (
	"github.com/diadata-org/diadata/pkg/dia"
)

// GetPairs returns the number of exchangepairs/pools on @exchange.
func (rdb *RelDB) GetNumPairs(exchange dia.Exchange) (numPairs int, err error) {

	exchangeType := GetExchangeType(exchange)
	switch exchangeType {
	case "CEX":
		pairs, err := rdb.GetExchangePairSymbols(exchange.Name)
		if err != nil {
			return len(pairs), err
		}
		numPairs = len(pairs)
	case "DEX":
		pools, err := rdb.GetAllPoolAddrsExchange(exchange.Name)
		if err != nil {
			return len(pools), err
		}
		numPairs = len(pools)
	}

	return
}
