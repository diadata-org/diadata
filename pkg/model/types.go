package models

import (
	"encoding/json"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
)

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
	ITIN               string
}

type Coin struct {
	Symbol             string
	Name               string
	Price              float64
	PriceYesterday     *float64
	VolumeYesterdayUSD *float64
	Time               time.Time
	CirculatingSupply  *float64
	ITIN               string
}

type Coins struct {
	CompleteCoinList []CoinSymbolAndName
	Change           *Change
	Coins            []Coin
}

// SymbolDetails is used for API return values
type SymbolDetails struct {
	Change    *Change
	Coin      Coin
	Rank      int
	Exchanges []SymbolExchangeDetails
	Gfx1      *Points
}

type CurrencyChange struct {
	Symbol        string
	Rate          float64
	RateYesterday float64
}

type Change struct {
	USD []CurrencyChange
}

// Point is used exclusively for chart points in the API
type Point struct {
	UnixTime int64
	Value    float64
}

type Points struct {
	DataPoints []clientInfluxdb.Result
}

// SymbolShort is used in ForeignQuotation.
// TO DO: Switch from ITIN to Address/Identifier
type SymbolShort struct {
	Symbol string
	ITIN   string
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

// CryptoIndex is the container for API endpoint CryptoIndex
type CryptoIndex struct {
	Name              string
	Address           string
	Value             float64
	Price             float64
	Price1h           float64
	Price24h          float64
	Price7d           float64
	Price14d          float64
	Price30d          float64
	Volume24hUSD      float64
	CirculatingSupply float64
	Time              time.Time
	Constituents      []struct {
		Name              string
		Symbol            string
		Address           string
		Price             float64
		CirculatingSupply float64
	}
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
