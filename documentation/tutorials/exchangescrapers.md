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

The other methods are mainly for Error handling, maintenance and shutdown. Our system is running the data scraper on our premises as an example. Of course, it is also possible to host your own data provider but in this simple example we assume that MyScraper is hosted on DIA servers.

For testing your scraper, you should create a config file for your api inside the `config/secrets` directory. Its name must be MySource.json and its format should be:

```javascript
{
  "ApiKey": "xx",
  "SecretKey": "yy"
}
```

Furthermore, add a reference to your scraper in `Config.go`  in the dia package and to the switch statement in `APIScraper.go`  in the scrapers package:

```go
func NewAPIScraper(exchange string, key string, secret string) APIScraper {
	switch exchange {
	case dia.MySourceExchange:
		return NewMySourceScraper(key, secret, dia.MySourceExchange)
	}
}
```

Before running the scraper execute the `main.go` from `cmd/services/pairDiscoveryServices`. Then, `collector.go`  in the  `cmd` folder will try to create a scraper for each exchange and collect the data pairs present in `config/exchange-scrapers.json` written by the method `fetchAvailablePairs()`.

Finally, run the scraping executable flagged as follows:

```text
go run collector.go -exchange MySource
```

