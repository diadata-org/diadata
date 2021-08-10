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

// Quotation is deprecating. Going to be substituted by AssetQuotation
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

// MarshalBinary for quotations
func (e *Quotation) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary for quotations
func (e *Quotation) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

type FiatQuotation struct {
	QuoteCurrency string
	BaseCurrency  string
	Price         float64
	Source        string
	Time          time.Time
}

// MarshalBinary for fiat quotations
func (fq *FiatQuotation) MarshalBinary() ([]byte, error) {
	return json.Marshal(fq)
}

// UnmarshalBinary for fiat quotations
func (fq *FiatQuotation) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &fq); err != nil {
		return err
	}
	return nil
}

// AssetQuotation is the most recent price point information on an asset.
type AssetQuotation struct {
	Asset  dia.Asset
	Price  float64
	Source string
	Time   time.Time
}

// MarshalBinary for quotations
func (aq *AssetQuotation) MarshalBinary() ([]byte, error) {
	return json.Marshal(aq)
}

// UnmarshalBinary for quotations
func (aq *AssetQuotation) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &aq); err != nil {
		return err
	}
	return nil
}

type Price struct {
	Symbol string
	Name   string
	Price  float64
	Time   time.Time
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

type Coins struct {
	CompleteCoinList []CoinSymbolAndName
	Change           *Change
	Coins            []Coin
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

// SymbolDetails is used for API return values
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

// Point is used exclusively for chart points in the API
type Point struct {
	UnixTime int64
	Value    float64
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

// SymbolShort is used in ForeignQuotation.
// TO DO: Switch from ITIN to Address/Identifier
type SymbolShort struct {
	Symbol string
	ITIN   string
}

type CoinSymbolAndName struct {
	Symbol string
	Name   string
}

type Pairs struct {
	Pairs []dia.ExchangePair
}
