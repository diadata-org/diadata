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

func (db *DB) GetPrice(symbol string, exchange string) (float64, error) {
	key := getKeyFilterSymbolAndExchangeZSET(dia.FilterKing, symbol, exchange)
	v, _, err := db.getZSETLastValue(key)
	return v, err
}

func (db *DB) GetPriceYesterday(symbol string, exchange string) (float64, error) {
	return db.getZSETValue(getKeyFilterZSET(getKey(dia.FilterKing, symbol, exchange)), time.Now().Unix()-WindowYesterday)
}
