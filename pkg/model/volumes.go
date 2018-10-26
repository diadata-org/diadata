package models

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"strconv"
	"time"
)

const (
	WindowVolume = 60 * 60 * 24
)

var (
	volumeKey = "VOL" + strconv.Itoa(dia.BlockSizeSeconds)
)

func (db *DB) SetVolume(symbol string, exchange string, volume float64, t time.Time) error {
	return db.SetFilter(volumeKey, symbol, exchange, volume, t)
}

func (db *DB) GetVolume(symbol string) (*float64, error) {
	return db.Sum24HoursInflux(symbol, "", volumeKey)
}

func (db *DB) GetVolumeExchange(symbol string, exchange string) (*float64, error) {
	return db.Sum24HoursInflux(symbol, exchange, volumeKey)
}
