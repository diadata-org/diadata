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
# pkg/dia/scraper/exchange-scraper/APIScraper.go
func NewAPIScraper(exchange string, key string, secret string) APIScraper {
 switch exchange {
 case dia.MySourceExchange:
  return NewMySourceScraper(key, secret, dia.MySourceExchange)
 }
}
```

Also, if your want to get data from contract, install `abigen` and generate the code from exchanger provided abi.

```sh
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

abigen --abi myexchange/myexchange.abi --pkg myexchange --type MyExchange --out myexchange/myexchange.go
```

Please put your abi file and generated code into a folder under the `exchange-scrapers/` and name it with the exchange name.

## Steps to run a scraper locally

1. Navigate to the `deployments/local/exchange-scraper` directory of the project.
2. Run the required services using `docker-compose up -d`, this will run and prepare Redis, PostgreSQL, and InfluxDB databases.
3. Set the required environment variables using the following commands:

```sh
export USE_ENV=true
export INFLUXURL=http://localhost:8086
export INFLUXUSER=test
export INFLUXPASSWORD=test
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=password
export POSTGRES_HOST=localhost
export POSTGRES_DB=postgres
export REDISURL=localhost:6379
```

Or simple by sourcing the `local.env` inside the `deployments/local/exchange-scraper` directory.

Also don't forget to set your Ethereum Node API Key. (If you are working on ethereum chain.)

```sh
export ETHEREUM_URL_REST=${YOUR_API_REST_ENDPOINT}
export ETHEREUM_URL_WS=${YOUR_API_WS_ENDPOINT}
```

Tips: Find you favourite ethereum node api provider at [EthereumNodes](https://ethereumnodes.com)

If you are working on another ethereum-compatible chain, replace `ETHEREUM` prefix with your chain name and get them from env.

```go
func NewAPIScraper() *APIScraper {
    // some initial stuff...
    log.Infof("Init rest and ws client for %s.", exchange.BlockChain.Name)
    restClient, err := ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", curveRestDial))
    if err != nil {
       log.Fatal("init rest client: ", err)
    }
    wsClient, err := ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", curveWsDial))
    if err != nil {
      log.Fatal("init ws client: ", err)
    }
    // other stuff here...
}

```

4. Execute `main.go` from `cmd/services/pairDiscoveryServices` for fetching the available pairs and setting them in the Redis database.
5. Finally, run the scraping executable flagged as follows:

```sh
cd cmd/exchange-scrapers/collector
go run collector.go -exchange MySource
```

For an illustration you can have a look at the `KrakenScraper.go`.
