package models

import (
	"time"
)

func (db *DB) SetFilter(filter string, symbol string, exchange string, volume float64, t time.Time) error {
	db.SaveFilterInflux(filter, symbol, exchange, volume, t)
	err := db.setZSETValue(getKeyFilterZSET(getKey(filter, symbol, exchange)), volume, t.Unix(), BiggestWindow)
	return err
}
