package models

import (
	"encoding/json"
	"time"
)

type ForeignQuotation struct {
	Symbol             string
	Name               string
	Price              float64
	PriceYesterday     float64
	VolumeYesterdayUSD float64
	Source             string
	Time               time.Time
	ITIN               string
}

type DefiScore struct {
	Protocol        string
	Symbol          string
	Score           float64
	LiquidityIndex  float64
	CollateralIndex float64
}

// MarshalBinary -
func (fq *ForeignQuotation) MarshalBinary() ([]byte, error) {
	return json.Marshal(fq)
}

// UnmarshalBinary -
func (fq *ForeignQuotation) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &fq); err != nil {
		return err
	}
	return nil
}
