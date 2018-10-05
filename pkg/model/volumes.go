package models

import (
	"github.com/diadata-org/api-golang/dia"
	"strconv"
	"time"
)

const (
	WindowVolume = 60 * 60 * 24
)

var (
	volumeKey = "VOL" + strconv.Itoa(dia.BlockSizeSeconds)
)

func (db *DB) SetVolume(symbol string, exchange string, volume float64) error {
	err := db.setZSETValue(getKeyFilterZSET(getKey(volumeKey, symbol, exchange)), volume, time.Now().Unix(), WindowVolume)
	return err
}

func (db *DB) GetVolume(symbol string) (*float64, error) {
	return db.getZSETSum(getKeyFilterZSET(volumeKey + "_" + symbol))
}

func (db *DB) GetVolumeExchange(symbol string, exchange string) (*float64, error) {
	return db.getZSETSum(getKeyFilterZSET(volumeKey + "_" + symbol + "_" + exchange))
}
