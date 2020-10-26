package models

import (
	"math/big"
	"time"
)

type PoolRate struct {
 	Rate  *big.Float
 	ProtocolName string
	BlockNumber int64
	PoolID string // hold pool id respective to protocol
 	TimeStamp time.Time
 	OutputAsset string
	InputAsset []string //some pools have more than 2 input assets

}
