package tradescraper

type TradeScraper struct {
	ChainName   string
	PoolAddress string
	TradeType   string  // sell or buy
	Amount0In   float64 // Token0 amount in
	Amount1In   float64 // Token1 amount in
	Amount0Out  float64 // Token0 amount out
	Amount1Out  float64 // Token1 amount out
	Price       float64
	Timestamp   int64 // trade timestamp in Unix format
}

type TradesData struct {
	TradeEntries []TradeScraper
}
