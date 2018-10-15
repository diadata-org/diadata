package models

import (
	"time"
)

func (db *DB) SetFilter(filter string, symbol string, exchange string, volume float64) error {
	db.SaveFilterInflux(filter, symbol, exchange, volume)
	err := db.setZSETValue(getKeyFilterZSET(getKey(filter, symbol, exchange)), volume, time.Now().Unix(), BiggestWindow)
	return err
}
