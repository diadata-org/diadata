package diaApi

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
	"time"
)

type ChangeCurrency struct {
	Symbol        string
	Rate          float64
	RateYesterday float64
}

type Change struct {
	USD []ChangeCurrency
}

type Coin struct {
	Symbol             string
	Name               string
	Price              float64
	PriceYesterday     *float64
	VolumeYesterdayUSD *float64
	Time               time.Time
	CirculatingSupply  *float64
}

type Coins struct {
	Change Change
	Coins  []Coin
}

type Pairs struct {
	Pairs []dia.Pair
}

type SymbolDetails struct {
	Change    Change
	Coin      Coin
	Exchanges []models.SymbolExchangeDetails
}
