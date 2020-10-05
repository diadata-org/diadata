package models

import (
	"math/big"
	"time"
)

type PoolRate struct {
 	Rate  *big.Int
 	ProtocolName string
 	PoolID string // hold pool id respective to protocol
 	TimeStamp time.Time
 	OutputAsset string
 }
