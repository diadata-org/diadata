package models

import (
	"encoding/json"
	"time"
)

type ForeignQuotation struct {
	Symbol             string    `json:"symbol"`
	Name               string    `json:"name"`
	Price              float64   `json:"price"`
	PriceYesterday     float64   `json:"price_yesterday"`
	VolumeYesterdayUSD float64   `json:"volume_yesterday_usd"`
	Source             string    `json:"source"`
	Time               time.Time `json:"time"`
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
