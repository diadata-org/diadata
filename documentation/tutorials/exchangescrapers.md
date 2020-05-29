# Write your own exchange scraper

## Add your own scraper

Before you begin writing a scraper, please check whether the exchange under consideration offers integration through websockets. In this case please implement your scraper using the websocket instead of a RESTful API.

Now, let's assume you want to scrape a data source that provides trade information. Create a new file in `exchange-scrapers/` and call it `MySourceScraper.go`. In order for us to be able to call your scraper from our system, you should introduce a `type MySourceScraper struct` that implements the `APIScraper` interface from the scrapers package:

```go
type APIScraper interface {
	io.Closer
	// ScrapePair returns a PairScraper that continuously scrapes trades for a
	// single pair from this APIScraper
	ScrapePair(pair dia.Pair) (PairScraper, error)
	// FetchAvailablePairs returns a list with all available trade pairs (usually
	// fetched from an exchange's API)
	FetchAvailablePairs() (pairs []dia.Pair, err error)
	// Channel returns a channel that can be used to receive trades
	Channel() chan *dia.Trade
}
```

From the `MySourceScraper` type you derive a `MySourcePairScraper` type which restricts the scraper to a specific pair and is returned by the `ScrapePair` method. Next, you should write a function with signature  `NewMySourceScraper(exchangeName string) *MySourceScraper`. We recommend that this function calls a method `mainLoop()`  in a go routine, constantly receiving trade information through the trade channel of `MySourceScraper`  as long as the channel is open.

Finally, please take care of proper error handling and cleanup. More precisely, you should include a method `Error()` which returns an error as soon as the scraper's channel closes, and methods `Close()` and `cleanup()` handling the closing/shutting down of channels.

Furthermore, in order for our system to see your scraper, add a reference to it in `Config.go`  in the dia package, and to the switch statement in `APIScraper.go`  in the scrapers package:

```go
func NewAPIScraper(exchange string, key string, secret string) APIScraper {
	switch exchange {
	case dia.MySourceExchange:
		return NewMySourceScraper(key, secret, dia.MySourceExchange)
	}
}
```

Before running the scraper execute the `main.go` from `cmd/services/pairDiscoveryServices`. Then, `collector.go`  in the folder  `cmd/exchange-scrapers/collector` will try to create a scraper for each exchange and collect the data pairs present in `config/exchange-scrapers.json` written by the method `fetchAvailablePairs()`.

Finally, run the scraping executable flagged as follows:

```text
go run collector.go -exchange MySource
```

