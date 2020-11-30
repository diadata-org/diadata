package models

import (
	"encoding/json"
	"time"
)

type FarmingPool struct {
	Rate         float64
	Balance      float64
	ProtocolName string
	BlockNumber  int64
	PoolID       string // hold pool id respective to protocol
	TimeStamp    time.Time
	OutputAsset  []string
	InputAsset   []string //some pools have more than 2 input assets
}

type FarmingPoolType struct {
	ProtocolName string
	InputAsset   []string
	PoolID       string
}

// MarshalBinary for interest rates
func (fp *FarmingPool) MarshalBinary() ([]byte, error) {
	return json.Marshal(fp)
}

// UnmarshalBinary for interest rates
func (fp *FarmingPool) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, fp); err != nil {
		return err
	}
	return nil
}
