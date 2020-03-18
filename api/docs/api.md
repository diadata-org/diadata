# diadata.org API

The world's crowd-driven financial data community has a professional API made for you.  
Decentral and transparent by design.  
With our decentral approach to data verification, you can gain a deep insight into current and past pricing, volume and exchange info so you can make the right decisions to stay ahead of the game.  
  
**Find the right data for your needs**  
Show your users the most transparent data on the market with our API. Whether you're building a financial service, a portfolio management tool, a new media offering, or more, we have the most advanced and updated data on the market for your product.  
For Oracle usage see [github](https://github.com/diadata-org/diadata/blob/master/documentation/methodology/Oracles.md).  
  
**Backtest your strategies**  
Use the most efficient and transparent crypto data to run simulations and backtest your trading or investing strategies. With crowd-aggregated hundreds of exchanges you can be sure that you're getting the right picture every single time.  
  
**Run Experiments**  
Build your own models with our data, to further your interest or just for fun. With our flexible and powerful API, we provide you with a set of data that will help you draw insights and make conclusions.  
  
**Request your data**  
Set a bounty on gitcoin.io or drop us [line](mailto:API@diadata.org).

Version: 1.0

## Base URL

The DIA base url is `https://api.diadata.org/`. All API paths are sub-paths of this base URL.

## Paths

### GET /v1/chartPoints/

Get chart points for an exchange.  
Example: [https://api.diadata.org/v1/chartPoints/MEDIR120/GateIO/EOS](https://api.diadata.org/v1/chartPoints/MEDIR120/GateIO/EOS)

Query Params:

* scale \[string\]: scale 5m 30m 1h 4h 1d 1w.

Path Params:

* filter \[string\]: Some filter. \(for now MEDIR120 or MAIR120\)
* trading place \[string\]: Some trading place.
* symbol \[string\]: Some symbol from GET /v1/coins

Responses:

* 200: success.

  [models.Points](api.md#modelspoints)

* 404: Symbol not found.

  [restApi.APIError](api.md#restapiapierror)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

_Remark:_ Careful! Successful responses can be rather large.

### GET /v1/chartPointsAllExchanges/

Get Symbol Details.  
Example: [https://api.diadata.org/v1/chartPointsAllExchanges/MEDIR120/EOS](https://api.diadata.org/v1/chartPointsAllExchanges/MEDIR120/EOS)

Query Params:

* scale \[string\]: scale 5m 30m 1h 4h 1d 1w.

Path Params:

* filter \[string\]: Some filter. \(for now MEDIR120 or MAIR120\)
* symbol \[string\]: Some symbol.

Responses:

* 200: success.

  [models.Points](api.md#modelspoints)

* 404: Symbol not found.

  [restApi.APIError](api.md#restapiapierror)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

_Remark:_ Careful! Successful responses can be rather large.

### GET /v1/quotation/

Get a quotation.  
Example: [https://api.diadata.org/v1/quotation/ETH](https://api.diadata.org/v1/quotation/ETH)

Path Params:

* symbol \[string\]: Some symbol.

Responses:

* 200: success.

  [models.Quotation](api.md#modelsquotation)

* 404: Symbol not found.

  [restApi.APIError](api.md#restapiapierror)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

### GET /v1/interestrate/

Get value for a certain rate type.  
Example: [https://api.diadata.org/v1/interestrate/ESTER/2020-03-16](https://api.diadata.org/v1/interestrate/ESTER/2020-03-16)

Path Params:

* rateType \[string\]: Short hand notation for a rate
* date \[string\]: In the format yyyy:mm:dd  date is an optional parameter. When omitted, the most recent value is returned.

Responses:

* 200: success.

  [models.InterestRate](api.md#modelsquotation)

* 404: Symbol not found.

  [restApi.APIError](api.md#restapiapierror)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

### GET /v1/supply/

Get the circulating supply corresponding to a symbol.  
Example: [https://api.diadata.org/v1/supply/ETH](https://api.diadata.org/v1/supply/ETH)

Path Params:

* symbol \[string\]: Some symbol.

Responses:

* 200: success.

  [dia.Supply](api.md#diasupply)

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

### GET /v1/coins

Get all available coins.  
Example: [https://api.diadata.org/v1/coins](https://api.diadata.org/v1/coins)

Responses:

* 200: success.

  [models.Coins](api.md#modelscoins)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

### GET /v1/pairs/

Get all available pairs.   
Example: [https://api.diadata.org/v1/pairs](https://api.diadata.org/v1/pairs)

Responses:

* 200: success.

  [models.Pairs](api.md#modelspairs)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

### GET /v1/symbols

Get all available symbols.  
Example: [https://api.diadata.org/v1/symbols](https://api.diadata.org/v1/symbols)

Responses:

* 200: success.

  [dia.Symbols](api.md#diasymbols)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

### GET /v1/interestrates

Get all available interest rates.  
Example: [https://api.diadata.org/v1/interestrates](https://api.diadata.org/v1/coins)

Responses:

* 200: success.

  [dia.Symbols](api.md#diasymbols)

* 500: error.

  [restApi.APIError](api.md#restapiapierror)

## Use cases

### Bash scripting

The API can be accessed through a Linux terminal by using curl.  For example  
`curl https://api.diadata.org/v1/interestrate/ESTER/2020-03-16 >> userPath/myFile.txt`  
writes the return value of the GET request into `myFile.txt` for further processing.

### Usage with Python

The object obtained in an API GET request is a string which complies with Python syntax. It can be cast into a list or dictionary resp. using Python's `eval(string)` function.









