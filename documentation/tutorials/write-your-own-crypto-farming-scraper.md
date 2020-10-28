# Write your own Crypto Farming scraper

## Getting Started

To add a Farming Scraper you should retrieve all data necessary to fill the following struct

```text
type FarmingPool struct {
	Rate         float64
	Balance      float64
	ProtocolName string
	BlockNumber  int64
	PoolID       string
	TimeStamp    time.Time
	OutputAsset  []string
	InputAsset   []string
}

```

which can be found in pkg/model/farmingpools.go.  
The scraper should be implemented as a `NewYourSourceScraper` in internal/pkg/farming-pool-scraper/poolscraper.go in line 30ff.

The new scraper should call a `mainLoop` in a go routine, scraping new data continuously. If the smart contract allows for, please use events in order to update pool data.

For illustration please have a look at the already existing farming scrapers in internal/pkg/farming-pool-scraper.

