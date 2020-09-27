package models

import "time"

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
