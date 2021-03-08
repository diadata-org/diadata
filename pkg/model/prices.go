package models

import (
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
)

func (db *DB) GetPrice(asset dia.Asset, exchange string) (float64, error) {
	key := getKeyFilterSymbolAndExchangeZSET(dia.FilterKing, asset, exchange)
	v, _, err := db.getZSETLastValue(key)
	return v, err
}

func (db *DB) GetPriceYesterday(asset dia.Asset, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, asset, exchange)), time.Now().Unix()-WindowYesterday)
}

func (db *DB) GetPrice1h(asset dia.Asset, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, asset, exchange)), time.Now().Unix()-Window1h)
}

func (db *DB) GetPrice7d(asset dia.Asset, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, asset, exchange)), time.Now().Unix()-Window7d)
}

func (db *DB) GetPrice14d(asset dia.Asset, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, asset, exchange)), time.Now().Unix()-Window14d)
}

func (db *DB) GetPrice30d(asset dia.Asset, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, asset, exchange)), time.Now().Unix()-Window30d)
}

func (db *DB) GetTradePriceBefore(asset dia.Asset, exchange string, timestamp time.Time) (*dia.Trade, error) {
	return db.GetTradeInflux(asset, exchange, timestamp)
}

func (db *DB) GetTradePrice1h(asset dia.Asset, exchange string) (*dia.Trade, error) {
	return db.GetTradePriceBefore(asset, exchange, time.Now().Add(-1*time.Hour))
}

func (db *DB) GetTradePrice24h(asset dia.Asset, exchange string) (*dia.Trade, error) {
	return db.GetTradePriceBefore(asset, exchange, time.Now().Add(-24*time.Hour))
}

func (db *DB) GetTradePrice7d(asset dia.Asset, exchange string) (*dia.Trade, error) {
	return db.GetTradePriceBefore(asset, exchange, time.Now().Add(-7*24*time.Hour))
}

func (db *DB) GetTradePrice14d(asset dia.Asset, exchange string) (*dia.Trade, error) {
	return db.GetTradePriceBefore(asset, exchange, time.Now().Add(-14*24*time.Hour))
}

func (db *DB) GetTradePrice30d(asset dia.Asset, exchange string) (*dia.Trade, error) {
	return db.GetTradePriceBefore(asset, exchange, time.Now().Add(-30*24*time.Hour))
}
