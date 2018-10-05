package models

import (
	"github.com/diadata-org/api-golang/dia"
	log "github.com/sirupsen/logrus"
	"time"
)

func (db *DB) SetPriceZSET(symbol string, exchange string, price float64) error {

	key := getKeyFilterZSET(getKey(dia.FilterKing, symbol, exchange))
	log.Println("SetPriceZSET ", key)
	return db.setZSETValue(key, price, time.Now().Unix(), BiggestWindow)
}
