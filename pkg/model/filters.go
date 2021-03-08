package models

import (
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
)

// SetFilter stores a filter point
func (db *DB) SetFilter(filter string, asset dia.Asset, exchange string, value float64, t time.Time) error {
	db.SaveFilterInflux(filter, asset, exchange, value, t)
	err := db.setZSETValue(getKeyFilterZSET(getKey(filter, asset, exchange)), value, t.Unix(), BiggestWindow)
	return err
}
