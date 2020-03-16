# Write your own rate scraper

These instructions concern scrapers for a single (floating point) number. For scrapers yielding pairs of numbers see the instructions in exchangescrapers.md.

## Add your own scraper

In order to add your own scraper for a new data source, you must adhere to our format. We use Go modules for our scrapers, so that each data provider is living as an independent module.

### Practical advice
Let's assume you want to scrape a data source that provides floating point information. Create a new file in `exchange-scrapers/` and call it `MySourceScraper.go`. The main difference between scrapers is the Update() method, where the actual scraping is done.

```go
func (s *MyScraper) Update() error {
  // scraper code here
}
```

### Overview of architecture


```go
// PairScraper receives trades for a single pc.Pair from a single exchange.
type PairScraper interface {
  io.Closer

  // Channel returns a channel that can be used to receive trades
  Channel() chan *models.InterestRate

  // Error returns an error when the channel Channel() is closed
  // and nil otherwise
  Error() error

  // Pair returns the pair this scraper is subscribed to
  Pair() dia.Pair
}
```

The other methods are mainly for Error handling, maintenance and shutdown. Our system is running the data scraper on our premises as an example. Of course, it is also possible to host your own data provider but in this simple example we assume that MyScraper is hosted on DIA servers.

For testing your scraper, you should create a config file for your api inside the `config/secrets` directory.

Its name must be MySource.json and its format should be:

```javascript
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

This is the basic structure of these files. Run the scraping binary by calling:

```text
go run scraper.go
```

from the `cmd` directory.

