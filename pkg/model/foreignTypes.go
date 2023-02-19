package models

import (
	"encoding/json"
	"time"
)

type ForeignQuotation struct {
	Symbol             string    `json:"Symbol"`
	Name               string    `json:"Name"`
	Price              float64   `json:"Price"`
	PriceYesterday     float64   `json:"PriceYesterday"`
	VolumeYesterdayUSD float64   `json:"VolumeYesterdayUSD"`
	Source             string    `json:"Source"`
	Time               time.Time `json:"Time"`
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
