package dia

import (
	"time"
)

const (
	Diadata = "diadata.org"
)

type Supply struct {
	Symbol            string
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
