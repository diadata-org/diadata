package models

import (
	"time"
)

type FarmingPool struct {
	// Rate is the pool rate. More information: https://docs.diadata.org/documentation/methodology/digital-assets/return-rates-in-crypto-farming
	Rate float64
	// Balance is the total supply of pool token
	Balance      float64
	ProtocolName string
	BlockNumber  int64
	PoolID       string // hold pool id respective to protocol
	TimeStamp    time.Time
	// OutputAsset is the list of tokens that you get back for staking.
	OutputAsset []string
	InputAsset  []string // some pools have more than 2 input assets
}

type FarmingPoolType struct {
	ProtocolName string
	InputAsset   []string
	PoolID       string
}
