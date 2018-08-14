package dia

import (
	"time"
)

type Supply struct {
	Symbol            string  `json:"symbol"`
	CirculatingSupply float64 `json:"circulatingSupply"`
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
