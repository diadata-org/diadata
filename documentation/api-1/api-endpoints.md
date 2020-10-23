---
description: >-
  The DIA base url is https://api.diadata.org. All API paths are sub-paths of
  this base URL.
---

# API Endpoints

## Digital Assets

{% api-method method="get" host="https://api.diadata.org" path="/v1/symbols" %}
{% api-method-summary %}
Symbols
{% endapi-method-summary %}

{% api-method-description %}
Get a list of all available symbols for cryptocurrencies.  
Example:  
https://api.diadata.org/v1/symbols  
  
Get symbols restricted to an exchange using the query parameter. \(For the moment only for centralized exchanges\).  
Example:  
https://api.diadata.org/v1/symbols?exchange=Kraken  
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-query-parameters %}
{% api-method-parameter name="exchange" type="string" required=false %}
Name of the crypto exchange.
{% endapi-method-parameter %}
{% endapi-method-query-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful retrieval of available symbols for cryptocurrencies. Shown below is an exerpt of the full response.
{% endapi-method-response-example-description %}

```
{"Symbols":["EOS","QTUM","BCH","BFT","FLDC","NXS","BLOCK","GAM","GLD","LOOM",...
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/quotation/:symbol" %}
{% api-method-summary %}
Quotation
{% endapi-method-summary %}

{% api-method-description %}
Get most recent information on the currency corresponding to symbol.  
Example: https://api.diadata.org/v1/quotation/BTC
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="symbol" type="string" required=true %}
Which symbol to get a quotation for, e.g., BTC.
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful retrieval of the BTC symbol.
{% endapi-method-response-example-description %}

```
{"Symbol":"BTC","Name":"Bitcoin","Price":9777.19339776667,"PriceYesterday":9574.416265039981,"VolumeYesterdayUSD":298134760.8811487,"Source":"diadata.org","Time":"2020-05-19T08:41:12.499645584Z","ITIN":"DXVPYDQC3"}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/exchanges" %}
{% api-method-summary %}
Exchanges
{% endapi-method-summary %}

{% api-method-description %}
Get a list of all available crypto exchanges.  
https://api.diadata.org/v1/exchanges
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful retrieval of available exchanges.
{% endapi-method-response-example-description %}

```
["Binance","Bitfinex","Bittrex","CoinBase","GateIO","HitBTC","Huobi","Kraken","LBank","OKEx","Quoine","Simex","ZB"]
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/chartPoints/:filter/:exchange/:symbol" %}
{% api-method-summary %}
Chart Points
{% endapi-method-summary %}

{% api-method-description %}
Get chart points for an exchange.  
https://api.diadata.org/v1/chartPoints/MEDIR120/Binance/BTC  
  
For a list of available exchanges see:  
https://docs.diadata.org/documentation/api-1\#api-access  
or:  
https://docs.diadata.org/documentation/api-1/api-endpoints\#exchanges  
  
  
  
_Remark_: Successful responses can be rather large.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="filter" type="string" required=true %}
Which filter should be applied \(Available options: MEDIR120 and MAIR120\).
{% endapi-method-parameter %}

{% api-method-parameter name="exchange" type="string" required=true %}
Which exchange to use.
{% endapi-method-parameter %}

{% api-method-parameter name="symbol" type="string" required=true %}
A valid symbol from GET /v1/coins, e.g., BTC.
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-query-parameters %}
{% api-method-parameter name="scale" type="string" %}
Which scale the graph points distance should have. Available options: 5m 30m 1h 4h 1d 1w.
{% endapi-method-parameter %}
{% endapi-method-query-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful retrieval of a chart point.
{% endapi-method-response-example-description %}

```
{"DataPoints":[{"Series":[{"name":"filters","columns":["time","exchange","filter","symbol","value"],"values":[["2020-05-19T08:02:09Z","GateIO","MEDIR120","EOS",2.6218717017500084]]}],"Messages":null}]}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/chartPointsAllExchanges/:filter/:symbol" %}
{% api-method-summary %}
Chart Points for all Exchanges
{% endapi-method-summary %}

{% api-method-description %}
Get symbol details for all exchanges.  
Example: https://api.diadata.org/v1/chartPointsAllExchanges/MEDIR120/EOS  
  
_Remark:_ Careful! Successful responses can be rather large.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="filter" type="string" required=true %}
Which filter should be applied \(Available options: MEDIR120 and MAIR120\).
{% endapi-method-parameter %}

{% api-method-parameter name="symbol" type="string" required=true %}
A valid symbol from GET /v1/coins, e.g., BTC.
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-query-parameters %}
{% api-method-parameter name="scale" type="string" required=false %}
Which scale the graph points distance should have. Available options: 5m 30m 1h 4h 1d 1w
{% endapi-method-parameter %}
{% endapi-method-query-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful retrieval of a chart point for all exchanges.
{% endapi-method-response-example-description %}

```
{"DataPoints":[{"Series":[{"name":"filters","columns":["time","exchange","filter","symbol","value"],"values":[["2020-05-19T08:17:59Z",null,"MEDIR120","EOS",2.6236194301032314]]}],"Messages":null}]}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/supply/:symbol" %}
{% api-method-summary %}
Supply
{% endapi-method-summary %}

{% api-method-description %}
Get the current circulating supply for the token corresponding to symbol.  
Example: https://api.diadata.org/v1/supply/BTC
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="symbol" type="string" required=true %}
Which symbol to get the supply for, e.g., BTC
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful retrieval of BTC supply.
{% endapi-method-response-example-description %}

```
{"Symbol":"BTC","Name":"Bitcoin","CirculatingSupply":17655550,"Source":"diadata.org","Time":"2019-04-20T08:44:25.748170404Z","Block":0}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/supplies/:symbol" %}
{% api-method-summary %}
Supplies
{% endapi-method-summary %}

{% api-method-description %}
Get all recorded supply values for the token corresponding to symbol.  
Example:  
https://api.diadata.org/v1/supplies/BTC  
  
Get supply values for a time range using the query parameters.  
Example:  
https://api.diadata.org/v1/supplies/BTC?starttime=1602232273&endtime=1602318673  
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="symbol" type="string" required=true %}
Which symbol to get the supply fot, e.g., BTC
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-query-parameters %}
{% api-method-parameter name="starttime" type="integer" required=false %}
Unix timestamp setting the start of the return array
{% endapi-method-parameter %}

{% api-method-parameter name="endtime" type="integer" required=false %}
Unix timestamp setting the end of the return array
{% endapi-method-parameter %}
{% endapi-method-query-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful retrieval of two supply values for Bitcoin \(BTC\) between timestamps 1591700000 and 1591883936.
{% endapi-method-response-example-description %}

```
[{"Symbol":"BTC","Name":"Bitcoin","CirculatingSupply":18399687,"Source":"diadata.org","Time":"2020-06-09T23:59:59Z","Block":0},{"Symbol":"BTC","Name":"Bitcoin","CirculatingSupply":18400712,"Source":"diadata.org","Time":"2020-06-10T23:59:59Z","Block":0}]
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/symbol/:symbol" %}
{% api-method-summary %}
Symbol
{% endapi-method-summary %}

{% api-method-description %}
Get extensive information on the cryptocurrency corresponding to symbol on various exchanges.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="symbol" type="string" required=true %}
Which symbol to get the details on, e.g., BTC
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Information on the cryptocurrency organized by "Change", "Coin", "Rank", "Exchanges" and "Gfx1"  \(filtered data\). Shown below is an excerpt of a successful response of symbol = BTC.
{% endapi-method-response-example-description %}

```
"Change":{"USD":[{"Symbol":"EUR","Rate":0.8995232526760818,"RateYesterday":0.8995232526760818},...

"Coin":{"Symbol":"BTC","Name":"Bitcoin","Price":9780.807149999986,"PriceYesterday":9574.416265039981,"VolumeYesterdayUSD":354341949.0902907,"Time":"2020-05-19T10:13:22.895692183Z","CirculatingSupply":17655550},...

"Rank":1

"Exchanges":[{"Name":"Huobi","Price":9776.344026379707,"PriceYesterday":9566.082031390646,"VolumeYesterdayUSD":182131794.24870485,"Time":"2020-05-19T10:07:59Z","LastTrades":...

"Gfx1":{"DataPoints":[{"Series":[{"name":"filters","columns":["time","exchange","filter","symbol","value"],"values":[["2020-05-19T10:08:00Z",null,"MA120","BTC",9780.807149999986],...
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/coins" %}
{% api-method-summary %}
Coins
{% endapi-method-summary %}

{% api-method-description %}
Get a list of all available coins.  
https://api.diadata.org/v1/coins
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful retrieval of available coins along with actual information on prices. Shown below is an exerpt of the full response.
{% endapi-method-response-example-description %}

```
"CompleteCoinList":[{"Symbol":"BTC","Name":"Bitcoin"},{"Symbol":"ETH","Name":"Ethereum"},...

"Change":{"USD":[{"Symbol":"EUR","Rate":0.8995232526760818,"RateYesterday":0.8995232526760818},...

"Coins":[{"Symbol":"BTC","Name":"Bitcoin","Price":9773.78345474998,"PriceYesterday":9574.416265039981,"VolumeYesterdayUSD":352085287.0431704,"Time":"2020-05-19T10:05:53.191886175Z","CirculatingSupply":17655550},...
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/volume/:symbol" %}
{% api-method-summary %}
Trade Volume
{% endapi-method-summary %}

{% api-method-description %}
Get the trading volume of the specified symbol in a defined time span.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="symbol" type="string" required=true %}
Which symbol to retrieve the volume of \(e.g. BTC\)
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-query-parameters %}
{% api-method-parameter name="starttime" type="integer" required=false %}
Start of the timespan \(Unix time in seconds\)
{% endapi-method-parameter %}

{% api-method-parameter name="endtime" type="integer" required=false %}
End of the timespan \(Unix time in seconds\)
{% endapi-method-parameter %}
{% endapi-method-query-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
An example response when querying a BTC volume for a typical day.
{% endapi-method-response-example-description %}

```
1431527525.7309263
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/kafka/tradesBlock" %}
{% api-method-summary %}
Raw crypto trades
{% endapi-method-summary %}

{% api-method-description %}
Get a list of all trades that comprised the last block that was used to calculate the latest information on crypto asset pricing.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="offset" type="integer" required=false %}
Get historical blocks \(use the current offset returned in a response to calculate the offset you want to get\)
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
A list of trades wrapped into a block with additional meta information like the time span of this specific block.
{% endapi-method-response-example-description %}

```
{"Result":{"offset":433850,"messages":[[{"BlockHash":"v1_4d7b1e936e7e0808d9ab17a43ec5ef8a","TradesBlockData":{"BeginTime":"2020-05-20T12:24:00Z","EndTime":"2020-05-20T12:26:00Z","TradesNumber":5674,"Trades":[{"Symbol":"EOS","Pair":"EOS_ETH","Price":0.01243882,"Volume":0.0325,"Time":"2020-05-20T12:24:00.050719107Z","ForeignTradeID":"c0d40b32","EstimatedUSDPrice":2.649370741608955,"Source":"LBank"}]}}]]}}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/cviIndex" %}
{% api-method-summary %}
CVI Index 
{% endapi-method-summary %}

{% api-method-description %}
Get all values of the Crypto Volatility Index.  
Example: https://api.diadata.org/v1/cviIndex  
  
Example with query parameters:  
https://api.diadata.org/v1/cviIndex?starttime=1589829000&endtime=1589830000
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-query-parameters %}
{% api-method-parameter name="starttime" type="integer" required=false %}
Unix timestamp setting the start of the return array
{% endapi-method-parameter %}

{% api-method-parameter name="endtime" type="integer" required=false %}
Unix timestamp setting the end of the return array
{% endapi-method-parameter %}
{% endapi-method-query-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful retrieval of CVI Index value for starttime=1589829000 and endtime=1589830000
{% endapi-method-response-example-description %}

```
[{"Timestamp":"2020-05-18T19:12:43Z","Value":142.28101897342574},{"Timestamp":"2020-05-18T19:17:48Z","Value":142.29282246717017},{"Timestamp":"2020-05-18T19:22:51Z","Value":142.3025697159107}]
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/defiLendingRate/:protocol/:asset" %}
{% api-method-summary %}
Defi Interest Rate
{% endapi-method-summary %}

{% api-method-description %}
Get information about a Defi protocol's lending and borrowing rates.  
Time parameter is optional. If omitted, the most recent rate is returned.  
  
Example: https://api.diadata.org/v1/defiLendingRate/COMPOUND/USDC  
  
Get rates for a range of timestamps using optional query parameters.  
https://api.diadata.org/v1/defiLendingRate/COMPOUND/USDC?dateInit=1591646100&dateFinal=1595246100  
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="protocol" type="string" required=true %}
Name of the protocol, in uppercase
{% endapi-method-parameter %}

{% api-method-parameter name="asset" type="string" required=true %}
Asset short name, e.g. ETH for Ether
{% endapi-method-parameter %}

{% api-method-parameter name="time" type="integer" required=false %}
Unix timestamp. Default is the latest available rate
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-query-parameters %}
{% api-method-parameter name="dateInit" type="integer" required=false %}
Initial Unix timestamp for range queries
{% endapi-method-parameter %}

{% api-method-parameter name="dateFinal" type="integer" required=false %}
Final Unix timestamp for range queries
{% endapi-method-parameter %}
{% endapi-method-query-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful retrieval of a Defi interest rate.
{% endapi-method-response-example-description %}

```
{"Timestamp":"2020-07-20T11:54:56Z","LendingRate":1.250020254710238,"BorrowingRate":4.856778356760549,"Asset":"USDC","Protocol":{"Name":"COMPOUND","Address":"0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b","UnderlyingBlockchain":"Ethereum","Token":""}}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/defiLendingState/:protocol" %}
{% api-method-summary %}
Defi Lending Protocol
{% endapi-method-summary %}

{% api-method-description %}
Get meta information about a defi lending protocol such as the underlying blockchain, its name and its currently locked value in USD and ETH.  
  
An example request can look like this: https://api.diadata.org/v1/defiLendingState/COMPOUND
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="protocol" type="string" required=true %}
Name of the protocol, e.g. COMPOUND
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful response containing locking volume, the timestamp of data recording and protocol meta information such as name and the underlying blockchain.
{% endapi-method-response-example-description %}

```
{"TotalUSD":13048619504.89947,"TotalETH":52570793.80784482,"Timestamp":"2020-07-22T16:27:31Z","Protocol":{"Name":"COMPOUND","Address":"0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b","UnderlyingBlockchain":"Ethereum","Token":""}}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org/v1/" path="foreignSymbols/:source" %}
{% api-method-summary %}
Guest Symbols
{% endapi-method-summary %}

{% api-method-description %}
Get the list of available symbols along with their ITIN for guest quotations.  
Example:  
https://api.diadata.org/v1/foreignSymbols/Coingecko
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="source" type="string" required=true %}
source of the quotation
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```

```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org/v1/foreignQuotation/:source/:symbol" path="" %}
{% api-method-summary %}
Guest Quotation
{% endapi-method-summary %}

{% api-method-description %}
Get the latest quotation for a token from a guest source.  
Example:  
https://api.diadata.org/v1/foreignQuotation/Coingecko/BTC  
  
Use the query parameter time in order to get the latest quotation before the specified timestamp.  
Example:  
https://api.diadata.org/v1/foreignQuotation/Coingecko/BTC?time=1601351679  
  
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="source" type="string" required=true %}
source of the quotation
{% endapi-method-parameter %}

{% api-method-parameter name="symbol" type="string" required=true %}
Which symbol to get a quotation for, e.g. BTC
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-query-parameters %}
{% api-method-parameter name="time" type="number" required=false %}
Unix timestamp.
{% endapi-method-parameter %}
{% endapi-method-query-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```

```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

## Traditional Assets

{% api-method method="get" host="https://api.diadata.org" path="/v1/interestrates" %}
{% api-method-summary %}
Interest Rates
{% endapi-method-summary %}

{% api-method-description %}
Get a list of all available interest rates along with metadata on the rates such as first publication date and issuing entity.  
https://api.diadata.org/v1/interestrates
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful retrieval of meta information on available interest rates.
{% endapi-method-response-example-description %}

```
[{"Symbol":"ESTER","FirstDate":"2019-10-01T00:00:00Z","Issuer":"ECB"},{"Symbol":"SOFR90","FirstDate":"2020-03-02T00:00:00Z","Issuer":"FED"},{"Symbol":"SONIA","FirstDate":"1997-01-02T00:00:00Z","Issuer":"BOE"},{"Symbol":"SAFR","FirstDate":"2020-03-02T00:00:00Z","Issuer":"FED"},{"Symbol":"SOFR","FirstDate":"2018-04-02T00:00:00Z","Issuer":"FED"},{"Symbol":"SOFR180","FirstDate":"2020-03-02T00:00:00Z","Issuer":"FED"},{"Symbol":"SOFR30","FirstDate":"2020-03-02T00:00:00Z","Issuer":"FED"}]
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/interestrate/:rateType" %}
{% api-method-summary %}
Interest Rate
{% endapi-method-summary %}

{% api-method-description %}
Get value for a certain rate type.  
Example: https://api.diadata.org/v1/interestrate/ESTER/2020-04-20​  
  
Get rate values for a range of timestamps using optional query parameters.  
Example: https://api.diadata.org/v1/interestrate/ESTER?dateInit=2020-02-20&dateFinal=2020-04-16
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="rateType" type="string" required=true %}
Symbol name for a rate.
{% endapi-method-parameter %}

{% api-method-parameter name="date" type="string" required=false %}
Return the rate for the specified date. Default date is the latest available date. Format: yyyy-mm-dd
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-query-parameters %}
{% api-method-parameter name="dateInit" type="string" required=false %}
Initial date for range queries. Format yyyy-mm-dd
{% endapi-method-parameter %}

{% api-method-parameter name="dateFinal" type="string" required=false %}
Final date for range queries. Format: yyyy-mm-dd
{% endapi-method-parameter %}
{% endapi-method-query-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful retrieval of an interest rate.
{% endapi-method-response-example-description %}

```
{"Symbol":"ESTER","Value":-0.542,"PublicationTime":"2020-05-19T07:15:07Z","EffectiveDate":"2020-05-18T00:00:00Z","Source":"ECB"}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/compoundedRate/:rateType/:dpy/:date" %}
{% api-method-summary %}
Compounded Index
{% endapi-method-summary %}

{% api-method-description %}
Get the value of an index compounded since its first publication date.  
  
Example:  
https://api.diadata.org/v1/compoundedRate/SOFR/360/2020-05-14  
  
Get the compounded index for a range of dates using the query parameters.  
Example:  
https://api.diadata.org/v1/compoundedRate/SOFR/360?dateInit=2020-04-24&dateFinal=2020-05-14  
  
For the methodology of compounded rates see:  
https://docs.diadata.org/documentation/methodology/traditional-assets/compounded-rates\#standard-methodology
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="rateType" type="string" required=true %}
Symbol for a rate name
{% endapi-method-parameter %}

{% api-method-parameter name="dpy" type="integer" required=true %}
Business day convention for the number of days per year
{% endapi-method-parameter %}

{% api-method-parameter name="date" type="string" required=true %}
Return the compounded index for the date specified in the format yyyy-mm-dd
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-query-parameters %}
{% api-method-parameter name="dateInit" type="string" required=false %}
Initial date for range queries. Format: yyyy-mm-dd
{% endapi-method-parameter %}

{% api-method-parameter name="dateFinal" type="string" required=false %}
Final date for range queries. Format: yyyy-mm-dd
{% endapi-method-parameter %}
{% endapi-method-query-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful retrieval of the SOFR Index.
{% endapi-method-response-example-description %}

```
{"Symbol":"SOFR_compounded_by_DIA","Value":1.0414032009923273,"PublicationTime":"0001-01-01T00:00:00Z","EffectiveDate":"2020-05-14T00:00:00Z","Source":"FED"}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/compoundedAvg/:rateType/:period/:dpy/:date" %}
{% api-method-summary %}
Compounded Average
{% endapi-method-summary %}

{% api-method-description %}
Get the average value of a given interest rate compounded over a period of time.  
Example:  
https://api.diadata.org/v1/compoundedAvg/SOFR/30/360/2020-05-14  
  
Get the compounded averages for a range of dates using the query parameters.  
Example:  
https://api.diadata.org/v1/compoundedAvg/SOFR/30/360?dateInit=2020-04-24&dateFinal=2020-05-14  
  
For the methodology see:  
https://docs.diadata.org/documentation/methodology/traditional-assets/compounded-rates\#standard-methodology  
  
Remark: This Get method requires an API key. Please contact us for more information:  
https://docs.google.com/forms/d/e/1FAIpQLSePxDwbEURjes4nw8GUzaT-XfYttRw\_6F2xAR607FKACsn7ew/viewform  
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="rateType" type="string" required=true %}
Symbol for a rate name
{% endapi-method-parameter %}

{% api-method-parameter name="period" type="integer" required=true %}
Rate is compounded over period days
{% endapi-method-parameter %}

{% api-method-parameter name="dpy" type="integer" required=true %}
Business day convention for the number of days per year
{% endapi-method-parameter %}

{% api-method-parameter name="date" type="string" required=true %}
Return the compounded rate for the date specified in the format yyyy-mm-dd
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-query-parameters %}
{% api-method-parameter name="dateInit" type="string" required=false %}
Initial date for range queries. Format: yyyy-mm-dd
{% endapi-method-parameter %}

{% api-method-parameter name="dateFinal" type="string" required=false %}
Final date for range queries. Format: yyyy-mm-dd
{% endapi-method-parameter %}
{% endapi-method-query-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful retrieval of a compounded average of the SOFR over an interest period of 30 calendar days.
{% endapi-method-response-example-description %}

```
{"Symbol":"SOFR30_compounded_by_DIA","Value":0.035667157687857554,"PublicationTime":"0001-01-01T00:00:00Z","EffectiveDate":"2020-05-14T00:00:00Z","Source":"FED"}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org" path="/v1/compoundedAvgDIA/:rateType/:period/:dpy/:date" %}
{% api-method-summary %}
Compounded Average using DIA Method
{% endapi-method-summary %}

{% api-method-description %}
Get the average value of an interest rate compounded over a period of time. Here, we use the DIA methodology for compounding the rate, i.e. interest is compounded for non-business days as well. For details see:  
https://docs.diadata.org/documentation/methodology/traditional-assets/compounded-rates\#dia-methodology  
  
Example:  
https://api.diadata.org/v1/compoundedAvgDIA/SOFR/30/360/2020-05-14  
  
Get the compounded average for a range of dates using the query parameters.   
Example:  
https://api.diadata.org/v1/compoundedAvgDIA/SOFR/30/360?dateInit=2020-04-24&dateFinal=2020-05-14  
  
_Remark_: This Get method requires an API key. Please contact us for more information:  
https://docs.google.com/forms/d/e/1FAIpQLSePxDwbEURjes4nw8GUzaT-XfYttRw\_6F2xAR607FKACsn7ew/viewform  
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="rateType" type="string" required=true %}
Symbol for a rate name
{% endapi-method-parameter %}

{% api-method-parameter name="period" type="integer" required=true %}
Rate is compounded over period days
{% endapi-method-parameter %}

{% api-method-parameter name="dpy" type="integer" required=true %}
Business convention for the number of days per year
{% endapi-method-parameter %}

{% api-method-parameter name="date" type="string" required=true %}
Return the compounded rate for the date specified in the format yyyy-mm-dd
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-query-parameters %}
{% api-method-parameter name="dateInit" type="string" required=false %}
Initial date for range queries. Format: yyyy-mm-dd
{% endapi-method-parameter %}

{% api-method-parameter name="dateFinal" type="string" required=false %}
Final date for range queries. Format: yyyy-mm-dd
{% endapi-method-parameter %}
{% endapi-method-query-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Successful retrieval of the compounded average of SOFR over an interest period of 30 calendar days.
{% endapi-method-response-example-description %}

```
[{"Symbol":"SOFR30_compounded_by_DIA","Value":0.035667175187725775,"PublicationTime":"0001-01-01T00:00:00Z","EffectiveDate":"2020-05-14T00:00:00Z","Source":"FED"}]
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.diadata.org/v1/" path="fiatQuotations" %}
{% api-method-summary %}
Fiat Currency Exchange Rates
{% endapi-method-summary %}

{% api-method-description %}
Get a list of exchange rates for several fiat currencies vs US Dollar.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="" type="string" required=false %}

{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```

```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}



