package models

import (
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
)

func (db *DB) SetPriceZSET(symbol string, exchange string, price float64, t time.Time) error {
	db.SaveFilterInflux(dia.FilterKing, symbol, exchange, price, t)
	key := getKeyFilterZSET(getKey(dia.FilterKing, symbol, exchange))
	log.Debug("SetPriceZSET ", key)
	return db.setZSETValue(key, price, time.Now().Unix(), Window30d)
}

func (db *DB) GetPrice(symbol string, exchange string) (float64, error) {
	key := getKeyFilterSymbolAndExchangeZSET(dia.FilterKing, symbol, exchange)
	v, _, err := db.getZSETLastValue(key)
	return v, err
}

func (db *DB) GetPriceYesterday(symbol string, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, symbol, exchange)), time.Now().Unix()-WindowYesterday)
}

func (db *DB) GetPrice1h(symbol string, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, symbol, exchange)), time.Now().Unix()-Window1h)
}

func (db *DB) GetPrice7d(symbol string, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, symbol, exchange)), time.Now().Unix()-Window7d)
}

func (db *DB) GetPrice14d(symbol string, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, symbol, exchange)), time.Now().Unix()-Window14d)
}

func (db *DB) GetPrice30d(symbol string, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, symbol, exchange)), time.Now().Unix()-Window30d)
}
