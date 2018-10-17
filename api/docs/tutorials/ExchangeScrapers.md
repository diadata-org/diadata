# Write your own exchange scraper

## Add your own scraper
In order to add your own scraper for a new data source, you must adhere to our format.
We use Go modules for our scrapers, so that each data provider is living as an independent module.
A scraper for cryptocurrencies must have a function `ScrapePair(pair dia.Pair)` that can be called from our system.
It returns a `PairScraper`.

Let's assume you want to scrape a data source that provides trade information.
Create a new file in `exchange-scrapers/` and call it `MySourceScraper.go`.
At first, its content looks like this:

```go
func (s *MyScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
  // scraper code here
}
```

The `PairScraper` interface is defined as a collection of methods.
The most important one is returning a channel which is filled with every trade the scraper witnesses

```go
// PairScraper receives trades for a single pc.Pair from a single exchange.
type PairScraper interface {
  io.Closer

  // Channel returns a channel that can be used to receive trades
  Channel() chan *dia.Trade

  // Error returns an error when the channel Channel() is closed
  // and nil otherwise
  Error() error

  // Pair returns the pair this scraper is subscribed to
  Pair() dia.Pair
}
```

The other methods are mainly for Error handling, maintenance and shutdown.
Our system is running the data scraper on our premises as an example.
Of course, it is also possible to host your own data provider but in this simple example we assume that MyScraper is hosted on DIA servers.

For testing your scraper, you should create a config file for your api inside the `config/secrets` directory.

Its name must be MySource.json and its format should be:

```json
{
  "ApiKey": "xx",
  "SecretKey": "yy"
}
```

Add a reference to your scraper in the switch statement

```go
switch e := configPair.Exchange; e {
    case dia.MySourceExchange:
      s[configPair.Exchange] = scrapers.NewMySourceScraper(configApi.ApiKey, configApi.SecretKey)
}
```

scraper.go will try to create a scraper for each exchange and collect the data pairs present in `config/exchange-scrapers.json`

This is the basic structure of these files.
Run the scrapping binary by calling:

```shell
go run scraper.go
```

from the `cmd` directory.
