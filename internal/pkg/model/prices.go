package models

import (
	"time"
)

func (db *DB) SetPriceZSET(symbol string, exchange string, price float64) error {
	return db.setZSETValue(getKeyFilterZSET(getKey(symbol, exchange)), price, time.Now().Unix(), BiggestWindow)
}
