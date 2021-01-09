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

func (db *DB) GetTradePriceBefore(symbol string, exchange string, timestamp time.Time) (*dia.Trade, error) {
	return db.GetTradeInflux(symbol, exchange, timestamp)
}

func (db *DB) GetTradePrice1h(symbol string, exchange string) (*dia.Trade, error) {
	return db.GetTradePriceBefore(symbol, exchange, time.Now().Add(-1 * time.Hour))
}

func (db *DB) GetTradePrice24h(symbol string, exchange string) (*dia.Trade, error) {
	return db.GetTradePriceBefore(symbol, exchange, time.Now().Add(-24 * time.Hour))
}

func (db *DB) GetTradePrice7d(symbol string, exchange string) (*dia.Trade, error) {
	return db.GetTradePriceBefore(symbol, exchange, time.Now().Add(-7 * 24 * time.Hour))
}

func (db *DB) GetTradePrice14d(symbol string, exchange string) (*dia.Trade, error) {
	return db.GetTradePriceBefore(symbol, exchange, time.Now().Add(-14 * 24 * time.Hour))
}

func (db *DB) GetTradePrice30d(symbol string, exchange string) (*dia.Trade, error) {
	return db.GetTradePriceBefore(symbol, exchange, time.Now().Add(-30 * 24 * time.Hour))
}
