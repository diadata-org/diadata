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
}

type StockQuotation struct {
	Symbol     string
	Name       string
	PriceAsk   float64
	PriceBid   float64
	SizeAskLot float64
	SizeBidLot float64
	Source     string
	Time       time.Time
	ISIN       string
}

type Stock struct {
	Symbol string
	Name   string
	ISIN   string
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

type AssetQuotationFull struct {
	Symbol             string
	Name               string
	Address            string
	Blockchain         string
	Price              float64
	PriceYesterday     float64
	VolumeYesterdayUSD float64
	Time               time.Time
	Source             string
}

// MarshalBinary for quotations
func (aq *AssetQuotationFull) MarshalBinary() ([]byte, error) {
	return json.Marshal(aq)
}

// UnmarshalBinary for quotations
func (aq *AssetQuotationFull) UnmarshalBinary(data []byte) error {
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

type CoinSymbolAndName struct {
	Symbol string
	Name   string
}

type Pairs struct {
	Pairs []dia.ExchangePair
}
