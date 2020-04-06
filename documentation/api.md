# DIAdata.org API

The world's crowd-driven financial data community has a professional API made for you.  


## Decentral and transparent by design

  
With our decentral approach to data verification, you can gain a deep insight into current and past pricing, volume and exchange info so you can make the right decisions to stay ahead of the game.  
  
**Find the right data for your needs**  
Show your users the most transparent data on the market with our API. Whether you're building a financial service, a portfolio management tool, a new media offering, or more, we have the most advanced and updated data on the market for your product.  
For Oracle usage see [github](https://github.com/diadata-org/diadata/tree/master/documentation/methodology/oracles.md).  
  
**Backtest your strategies**  
Use the most efficient and transparent crypto data to run simulations and backtest your trading or investing strategies. With crowd-aggregated hundreds of exchanges you can be sure that you're getting the right picture every single time.  
  
**Run Experiments**  
Build your own models with our data, to further your interest or just for fun. With our flexible and powerful API, we provide you with a set of data that will help you draw insights and make conclusions.  
  
**Request your data**  
Set a bounty on gitcoin.io or drop us [line](mailto:API@diadata.org).

Version: 1.0

## Base URL

The DIA base url is `https://api.diadata.org/`. All API paths are sub-paths of this base URL.

## Fair usage limits

We employ a fair-usage policy. The number of requests client applications make is logged, and these logs are subject to review. Excessive \(and unfair\) usage may result in action being taken to limit an application’s access to the Web API. Currently this is at **5 requests per minute.**

## Paths

### GET /v1/chartPoints/

Get chart points for.

Query Params:

* scale \[string\]: scale 5m 30m 1h 4h 1d 1w.

Path Params:

* symbol \[string\]: Some symbol.
* exchange \[string\]: Some exchange.
* filter \[string\]: Some filter.

Responses:

* 200: success.

  [models.Points](api.md#modelspoints)

* 404: Symbol not found.

  [restApi.APIError](api.md#restapiapierror)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

### GET /v1/chartPointsAllExchanges/

Get Symbol Details.

Query Params:

* scale \[string\]: scale 5m 30m 1h 4h 1d 1w.

Path Params:

* symbol \[string\]: Some symbol.
* filter \[string\]: Some filter.

Responses:

* 200: success.

  [models.Points](api.md#modelspoints)

* 404: Symbol not found.

  [restApi.APIError](api.md#restapiapierror)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

### GET /v1/coins

Get coins.

Responses:

* 200: success.

  [models.Coins](api.md#modelscoins)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

### GET /v1/pairs/

Get pairs.

Responses:

* 200: success.

  [models.Pairs](api.md#modelspairs)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

### GET /v1/quotation/

Get quotation.

Path Params:

* symbol \[string\]: Some symbol.

Responses:

* 200: success.

  [models.Quotation](api.md#modelsquotation)

* 404: Symbol not found.

  [restApi.APIError](api.md#restapiapierror)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

### POST /v1/supply

Post the circulating supply.

Query Params:

* Symbol \[string\]: Coin symbol.
* CirculatingSupply \[number\]: number of coins in circulating supply.

Responses:

* 200: success.

  [dia.Supply](api.md#diasupply)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

### GET /v1/supply/

Get supply.

Path Params:

* symbol \[string\]: Some symbol.

Responses:

* 200: success.

  [dia.Supply](api.md#diasupply)

* 404: Symbol not found.

  [restApi.APIError](api.md#restapiapierror)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

### GET /v1/symbol/

Get Symbol Details.

Path Params:

* symbol \[string\]: Some symbol.

Responses:

* 200: success.

  [models.SymbolDetails](api.md#modelssymboldetails)

* 404: Symbol not found.

  [restApi.APIError](api.md#restapiapierror)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

### GET /v1/symbols

Get all symbols list.

Query Params:

* scale \[string\]: scale 5m 30m 1h 4h 1d 1w.

Path Params:

* symbol \[string\]: Some symbol.
* filter \[string\]: Some filter.

Responses:

* 200: success.

  [dia.Symbols](api.md#diasymbols)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

## Definitions

### dia.Pair

* exchange _\(string\)_ - foreignName _\(string\)_ - ignore _\(boolean\)_ - symbol _\(string\)_

  **dia.Supply**

* block _\(integer\)_ - circulatingSupply _\(number\)_ - name _\(string\)_ - source _\(string\)_ - symbol _\(string\)_ - time _\(string\)_

  **dia.Symbols**

* symbols _\(array\)_ - type: string

  **dia.Trade**

* estimatedUSDPrice _\(number\)_ - foreignTradeID _\(string\)_ - pair _\(string\)_ - price _\(number\)_ - source _\(string\)_ - symbol _\(string\)_ - time _\(string\)_ - volume _\(number\)_

  **models.Change**

* usd _\(array\)_ - [models.CurrencyChange](api.md#modelscurrencychange)

  **models.Coin**

* circulatingSupply _\(number\)_ - name _\(string\)_ - price _\(number\)_ - priceYesterday _\(number\)_ - symbol _\(string\)_ - time _\(string\)_ - volumeYesterdayUSD _\(number\)_

  **models.Coins**

* change _\(object\)_ - coins _\(array\)_ - [models.Coin](api.md#modelscoin) - completeCoinList _\(array\)_ - [models.CoinSymbolAndName](api.md#modelscoinsymbolandname)

  **models.CoinSymbolAndName**

* name _\(string\)_ - symbol _\(string\)_

  **models.CurrencyChange**

* rate _\(number\)_ - rateYesterday _\(number\)_ - symbol _\(string\)_

  **models.Pairs**

* pairs _\(array\)_ - [dia.Pair](api.md#diapair)

  **models.Points**

* dataPoints _\(string\)_

  **models.Quotation**

* name _\(string\)_ - price _\(number\)_ - priceYesterday _\(number\)_ - source _\(string\)_ - symbol _\(string\)_ - time _\(string\)_ - volumeYesterdayUSD _\(number\)_

  **models.SymbolDetails**

* change _\(object\)_ - coin _\(object\)_ - exchanges _\(array\)_ - [models.SymbolExchangeDetails](api.md#modelssymbolexchangedetails) - gfx1 _\(object\)_ - rank _\(integer\)_

  **models.SymbolExchangeDetails**

* lastTrades _\(array\)_ - [dia.Trade](api.md#diatrade) - name _\(string\)_ - price _\(number\)_ - priceYesterday _\(number\)_ - time _\(string\)_ - volumeYesterdayUSD _\(number\)_

  **restApi.APIError**

* errorcode _\(integer\)_ - errormessage _\(string\)_

