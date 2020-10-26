package models

import (
	"time"
)

type FarmingPool struct {
	Rate         float64
	Balance      float64
	ProtocolName string
	BlockNumber  int64
	PoolID       string // hold pool id respective to protocol
	TimeStamp    time.Time
	OutputAsset  string
	InputAsset   []string //some pools have more than 2 input assets

}
