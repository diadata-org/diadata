package models

import (
	"github.com/diadata-org/diadata/pkg/dia"
)

// GetPairs returns all exchangepairs on @exchange.
func (rdb *RelDB) GetPairs(exchange string) (pairs []dia.ExchangePair, err error) {

	pairs, err = rdb.GetExchangePairSymbols(exchange)
	if err != nil {
		log.Error(err)
	}

	return
}
