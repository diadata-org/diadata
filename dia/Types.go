package dia

import (
	"time"
)

const (
	Diadata = "diadata.org"
)

type Quotation struct {
	Symbol string
	Name   string
	Price  float64
	Source string
	Time   time.Time
}

type Supply struct {
	Symbol            string
	Name              string
	CirculatingSupply float64
	Source            string
	Time              time.Time
}

type Pair struct {
	Symbol      string
	ForeignName string
}

type Trade struct {
	Symbol         string
	Pair           string
	Price          float64
	Volume         float64 // negative if result of Market order Sell
	Time           time.Time
	ForeignTradeID string
	Source         string
}
