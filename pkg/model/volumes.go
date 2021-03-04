package models

import (
	"strconv"

	"github.com/diadata-org/diadata/pkg/dia"
)

const (
	WindowVolume = 60 * 60 * 24
)

var (
	volumeKey = "VOL" + strconv.Itoa(dia.BlockSizeSeconds)
)

// GetVolume returns the 24h trading volume of @asset across exchanges.
func (db *DB) GetVolume(asset dia.Asset) (*float64, error) {
	return db.Sum24HoursInflux(asset, "", volumeKey)
}
