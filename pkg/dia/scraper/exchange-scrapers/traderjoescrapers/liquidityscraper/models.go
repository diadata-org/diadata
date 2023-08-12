package liquidityscraper

type LiquidityScraper struct {
	ChainName string 
	PoolAddress string 
	Token0      string  // address of first token
	Token1      string  // address of second token
	Reserves0   float64 // Token0 reserves
	Reserves1   float64 // Token1 reserves
	Timestamp   int64   // timestamp of when trade data was scraped.
}

type LiquidityData struct {
	LiquidityEntries []LiquidityScraper
}