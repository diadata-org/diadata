package dia

import (
	"encoding/json"
	"time"
)

const (
	Diadata = "diadata.org"
)

type Supply struct {
	Symbol            string
	Name              string
	CirculatingSupply float64
	Source            string
	Time              time.Time
	Block             int64
}

type Pair struct {
	Symbol      string
	ForeignName string
	Exchange    string
	Ignore      bool
}

type Pairs []Pair

type Trade struct {
	Symbol            string
	Pair              string
	Price             float64
	Volume            float64 // negative if result of Market order Sell
	Time              time.Time
	ForeignTradeID    string
	EstimatedUSDPrice float64 // will be filled by the TradeBlock Service
	Source            string
}

type TradesBlockData struct {
	BeginTime    time.Time
	EndTime      time.Time
	TradesNumber int
	Trades       []Trade
}

type TradesBlock struct {
	BlockHash       string
	TradesBlockData TradesBlockData
}

type FiltersBlock struct {
	BlockHash        string
	FiltersBlockData FiltersBlockData
}

type FiltersBlockData struct {
	TradesBlockHash string
	BeginTime       time.Time
	EndTime         time.Time
	FilterPoints    []FilterPoint
	FiltersNumber   int
}

type FilterPoint struct {
	Symbol string
	Value  float64
	Name   string
	Time   time.Time
}

// MarshalBinary -
func (e *FiltersBlock) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *FiltersBlock) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *Trade) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *Trade) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *TradesBlock) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *TradesBlock) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *Supply) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *Supply) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *Pairs) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *Pairs) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}
