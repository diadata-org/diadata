package diaApi

import (
	"github.com/diadata-org/api-golang/pkg/dia"
	"github.com/diadata-org/api-golang/pkg/model"
	"time"
)

type Change struct {
	EURUSD          *float64
	EURUSDYesterday *float64
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
	Exchanges map[string]models.SymbolExchangeDetails
}
