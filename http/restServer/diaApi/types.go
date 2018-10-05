package diaApi

import (
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/internal/pkg/model"
	"time"
)

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
	Coins []Coin
}

type Pairs struct {
	Pairs []dia.Pair
}

type SymbolDetails struct {
	Coin      Coin
	Exchanges map[string]models.SymbolExchangeDetails
}
