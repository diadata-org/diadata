package models

import (
	"encoding/json"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
)

type Point struct {
	UnixTime int64
	Value    float64
}

type SymbolExchangeDetails struct {
	Name               string
	Price              float64
	PriceYesterday     *float64
	VolumeYesterdayUSD *float64
	Time               *time.Time
	LastTrades         []dia.Trade
}

type Quotation struct {
	Symbol             string
	Name               string
	Price              float64
	PriceYesterday     *float64
	VolumeYesterdayUSD *float64
	Source             string
	Time               time.Time
}

type InterestRate struct {
	Symbol string
	Value  float64
	Time   time.Time
	Source string
}

type CurrencyChange struct {
	Symbol        string
	Rate          float64
	RateYesterday float64
}

type Change struct {
	USD []CurrencyChange
}

// MarshalBinary -
func (e *Change) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *Change) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

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
	Change           *Change
	Coins            []Coin
}

type Pairs struct {
	Pairs []dia.Pair
}

// MarshalBinary -
func (e *Coins) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *Coins) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

type SymbolDetails struct {
	Change    *Change
	Coin      Coin
	Rank      int
	Exchanges []SymbolExchangeDetails
	Gfx1      *Points
}

// MarshalBinary -
func (e *SymbolDetails) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *SymbolDetails) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

func (e *Coin) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *Coin) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

type Points struct {
	DataPoints []clientInfluxdb.Result
}

func (e *Points) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *Points) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}
