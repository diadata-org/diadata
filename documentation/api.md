# DIAdata.org API

The world's crowd-driven financial data community has a professional API made for you.<br><h2>Decentral and transparent by design</h2><br>With our decentral approach to data verification, you can gain a deep insight into current and past pricing, volume and exchange info so you can make the right decisions to stay ahead of the game.<br><br><h3>Find the right data for your needs</h3><br>Show your users the most transparent data on the market with our API. Whether you're building a financial service, a portfolio management tool, a new media offering, or more, we have the most advanced and updated data on the market for your product.<br>For Oracle usage see [github](https://github.com/diadata-org/diadata/blob/master/documentation/methodology/Oracles.md).<br><br><h3>Backtest your strategies</h3><br>Use the most efficient and transparent crypto data to run simulations and backtest your trading or investing strategies. With crowd-aggregated hundreds of exchanges you can be sure that you're getting the right picture every single time.<br><br><h3>Run Experiments</h3><br>Build your own models with our data, to further your interest or just for fun. With our flexible and powerful API, we provide you with a set of data that will help you draw insights and make conclusions.<br><br><h3>Request your data</h3><br>Set a bounty on gitcoin.io or drop us [line](mailto:API@diadata.org).

Version: 1.0

## Base URL

The DIA base url is `https://api.diadata.org/`. All API paths are sub-paths of this base URL.

## Paths

### GET /v1/chartPoints/ 
Get chart points for.

Query Params:
  - scale [string]: scale 5m 30m 1h 4h 1d 1w.

Path Params:
  - symbol [string]: Some symbol.
  - exchange [string]: Some exchange.
  - filter [string]: Some filter.

Responses:
  - 200: success.

    [models.Points](#modelspoints)
  - 404: Symbol not found.

    [restApi.APIError](#restapiapierror)
  - 500: error.

    [restApi.APIError](#restapiapierror)

### GET /v1/chartPointsAllExchanges/ 
Get Symbol Details.

Query Params:
  - scale [string]: scale 5m 30m 1h 4h 1d 1w.

Path Params:
  - symbol [string]: Some symbol.
  - filter [string]: Some filter.

Responses:
  - 200: success.

    [models.Points](#modelspoints)
  - 404: Symbol not found.

    [restApi.APIError](#restapiapierror)
  - 500: error.

    [restApi.APIError](#restapiapierror)

### GET /v1/coins
Get coins.

Responses:
  - 200: success.

    [models.Coins](#modelscoins)
  - 500: error.

    [restApi.APIError](#restapiapierror)

### GET /v1/pairs/
Get pairs.

Responses:
  - 200: success.

    [models.Pairs](#modelspairs)
  - 500: error.

    [restApi.APIError](#restapiapierror)

### GET /v1/quotation/
Get quotation.

Path Params:
  - symbol [string]: Some symbol.

Responses:
  - 200: success.

    [models.Quotation](#modelsquotation)
  - 404: Symbol not found.

    [restApi.APIError](#restapiapierror)
  - 500: error.

    [restApi.APIError](#restapiapierror)

### POST /v1/supply

Post the circulating supply.

Query Params:
  - Symbol [string]: Coin symbol.
  - CirculatingSupply [number]: number of coins in circulating supply.

Responses:
  - 200: success.

    [dia.Supply](#diasupply)
  - 500: error.

    [restApi.APIError](#restapiapierror)

### GET /v1/supply/

Get supply.

Path Params:
  - symbol [string]: Some symbol.

Responses:
  - 200: success.

    [dia.Supply](#diasupply)
  - 404: Symbol not found.

    [restApi.APIError](#restapiapierror)
  - 500: error.

    [restApi.APIError](#restapiapierror)

### GET /v1/symbol/
Get Symbol Details.

Path Params:
  - symbol [string]: Some symbol.

Responses:
  - 200: success.

    [models.SymbolDetails](#modelssymboldetails)
  - 404: Symbol not found.

    [restApi.APIError](#restapiapierror)
  - 500: error.

    [restApi.APIError](#restapiapierror)

### GET /v1/symbols

Get all symbols list.

Query Params:
  - scale [string]: scale 5m 30m 1h 4h 1d 1w.

Path Params:
  - symbol [string]: Some symbol.
  - filter [string]: Some filter.

Responses:
  - 200: success.

    [dia.Symbols](#diasymbols)
  - 500: error.

    [restApi.APIError](#restapiapierror)


## Definitions


### dia.Pair

  - exchange *(string)*  - foreignName *(string)*  - ignore *(boolean)*  - symbol *(string)*
### dia.Supply

  - block *(integer)*  - circulatingSupply *(number)*  - name *(string)*  - source *(string)*  - symbol *(string)*  - time *(string)*
### dia.Symbols

  - symbols *(array)*    - type: string
### dia.Trade

  - estimatedUSDPrice *(number)*  - foreignTradeID *(string)*  - pair *(string)*  - price *(number)*  - source *(string)*  - symbol *(string)*  - time *(string)*  - volume *(number)*
### models.Change

  - usd *(array)*    - [models.CurrencyChange](#modelscurrencychange)
### models.Coin

  - circulatingSupply *(number)*  - name *(string)*  - price *(number)*  - priceYesterday *(number)*  - symbol *(string)*  - time *(string)*  - volumeYesterdayUSD *(number)*
### models.Coins

  - change *(object)*  - coins *(array)*    - [models.Coin](#modelscoin)  - completeCoinList *(array)*    - [models.CoinSymbolAndName](#modelscoinsymbolandname)
### models.CoinSymbolAndName

  - name *(string)*  - symbol *(string)*
### models.CurrencyChange

  - rate *(number)*  - rateYesterday *(number)*  - symbol *(string)*
### models.Pairs

  - pairs *(array)*    - [dia.Pair](#diapair)
### models.Points

  - dataPoints *(string)*
### models.Quotation

  - name *(string)*  - price *(number)*  - priceYesterday *(number)*  - source *(string)*  - symbol *(string)*  - time *(string)*  - volumeYesterdayUSD *(number)*
### models.SymbolDetails

  - change *(object)*  - coin *(object)*  - exchanges *(array)*    - [models.SymbolExchangeDetails](#modelssymbolexchangedetails)  - gfx1 *(object)*  - rank *(integer)*
### models.SymbolExchangeDetails

  - lastTrades *(array)*    - [dia.Trade](#diatrade)  - name *(string)*  - price *(number)*  - priceYesterday *(number)*  - time *(string)*  - volumeYesterdayUSD *(number)*
### restApi.APIError

  - errorcode *(integer)*  - errormessage *(string)*
