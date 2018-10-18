package diaApi

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
	"time"
)

type CoinSymbolAndName struct {
	Symbol string
	Name   string
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
	CompleteCoinList []CoinSymbolAndName
	Change           *models.Change
	Coins            []Coin
}

type Pairs struct {
	Pairs []dia.Pair
}

type SymbolDetails struct {
	Change    *models.Change
	Coin      Coin
	Exchanges []models.SymbolExchangeDetails
}
