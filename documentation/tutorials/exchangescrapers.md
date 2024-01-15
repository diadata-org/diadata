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

From the `MySourceScraper` type you derive a `MySourcePairScraper` type which restricts the scraper to a specific pair. Next, you should write a function with signature  `NewMySourceScraper(exchangeName string) *MySourceScraper` initializing a scraper. We suggest that this function calls a method `func (s *MySourceScraper) mainLoop()`  in a go routine, constantly receiving trade information through the trade channel of `MySourceScraper`  as long as the channel is open. The collection of new trading information inside the `mainLoop()` should be done by an update method with signature `func (s *MySourceScraper) Update()`.  Finally, in order to implement the interface `APIScraper` you should include `ScrapePair` returning a `MySourcePairScraper`  for a specific pair, so our main collection method can iterate over all possible trading pairs.

Also, please take care of proper error handling and cleanup. More precisely, you should include a method `Error()` which returns an error as soon as the scraper's channel closes, and methods `Close()` and `cleanup()` handling the closing/shutting down of channels.

Furthermore, in order for our system to see your scraper, add a reference to it in `Config.go`  in the dia package, and to the switch statement in `APIScraper.go`  in the scrapers package:

```go
func NewAPIScraper(exchange string, key string, secret string) APIScraper {
	switch exchange {
	case dia.MySourceExchange:
		return NewMySourceScraper(key, secret, dia.MySourceExchange)
	}
}
```

## Steps to run a scraper locally

1. Navigate to the `deployments/local/exchange-scraper` directory of the project.
2. Run the required services using `docker-compose up -d`, this will run and prepare Redis, PostgreSQL, and InfluxDB databases.
3. Set the required environment variables using the following commands:

```sh
export USE_ENV=true
export INFLUXURL=http://localhost:8086
export INFLUXUSER=test
export INFLUXPASSWORD=test
export POSTGRES_USER=rider
export POSTGRES_PASSWORD=123test
export POSTGRES_HOST=localhost:5434
export POSTGRES_DB=dia
export REDISURL=localhost:6379
# for local development
export DIA_CONFIG_DIR=/path/to/diadata-project/config
```

Or simple by sourcing the `local.env` inside the `deployments/local/exchange-scraper` directory.

4. Execute `main.go` from `cmd/services/pairDiscoveryService` for fetching the available pairs and setting them in the Redis database.

```text
cd cmd/services/pairDiscoveryService
go run main.go -exchange MySource -mode remoteFetch
```

5. Finally, run the scraping executable flagged as follows:

```text
cd cmd/exchange-scrapers/collector
go run collector.go -exchange MySource
```

For an illustration you can have a look at the `KrakenScraper.go`.


### kafka tips

Sometimes kafka docker container cant create topics, required for writing
It can be resolved it by ssh login to kafka container and run kafka scripts
```
root@ad1f7deb4762:/# kafka-topics.sh --zookeeper zookeeper --topic trades --create --partitions 2 --replication-factor 1
Created topic trades.
root@ad1f7deb4762:/# kafka-topics.sh --zookeeper zookeeper --topic tradesReplicadia --create --partitions 2 --replication-factor 1
Created topic tradesReplicadia.
root@ad1f7deb4762:/# kafka-topics.sh --zookeeper zookeeper --topic tradestest --create --partitions 2 --replication-factor 1
Created topic tradestest.
```


