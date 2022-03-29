---
description: >-
  The DIA base url is https://api.diadata.org. All API paths are sub-paths of
  this base URL.
---

# API Endpoints

Digital Assets

{% swagger baseUrl="https://api.diadata.org" path="/v1/symbols" method="get" summary="Symbols" %}
{% swagger-description %}
Get a list of all available symbols for cryptocurrencies.

\




_Example_

:

\


https://api.diadata.org/v1/symbols

\




\


Get symbols restricted to an exchange using the query parameter. (For the moment only for centralized exchanges).

\




_Example_

:

\


https://api.diadata.org/v1/symbols?exchange=Kraken

\



{% endswagger-description %}

{% swagger-parameter in="query" name="exchange" type="string" %}
Name of the crypto exchange.
{% endswagger-parameter %}

{% swagger-response status="200" description="Successful retrieval of available symbols for cryptocurrencies. Shown below is an exerpt of the full response." %}
```
{"Symbols":["EOS","QTUM","BCH","BFT","FLDC","NXS","BLOCK","GAM","GLD","LOOM",...
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/quotation/:symbol" method="get" summary="Quotation" %}
{% swagger-description %}
Get most recent information on the currency corresponding to symbol.

\




_Example_

:

\


https://api.diadata.org/v1/quotation/BTC
{% endswagger-description %}

{% swagger-parameter in="path" name="symbol" type="string" %}
Which symbol to get a quotation for, e.g., BTC.
{% endswagger-parameter %}

{% swagger-response status="200" description="Successful retrieval of the BTC symbol." %}
```
{"Symbol":"BTC","Name":"Bitcoin","Price":9777.19339776667,"PriceYesterday":9574.416265039981,"VolumeYesterdayUSD":298134760.8811487,"Source":"diadata.org","Time":"2020-05-19T08:41:12.499645584Z","ITIN":"DXVPYDQC3"}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="get" path="/v1/assetQuotation" baseUrl="https://api.diadata.org" summary="Asset Quotation" %}
{% swagger-description %}
Returns the quotation for a fully qualified asset (i.e. distinguished by blockchain and address).

Example: https://api.diadata.org/v1/assetQuotation/Bitcoin/0x000000000000000000000000000000000000000
{% endswagger-description %}

{% swagger-parameter in="path" name="blockchain" required="true" %}
Name of the blockchain for requested asset
{% endswagger-parameter %}

{% swagger-parameter in="path" name="address" required="true" %}
Address of the requested asset
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Asset quotation with details about price, 24h volume, last update time, address, and blockchain." %}
```javascript
{
    // Response
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/exchanges" method="get" summary="Exchanges" %}
{% swagger-description %}
Get a list of all available crypto exchanges.

\


https://api.diadata.org/v1/exchanges
{% endswagger-description %}

{% swagger-response status="200" description="Successful retrieval of available exchanges." %}
```
["KuCoin","Uniswap","Balancer","Maker","Gnosis","Curvefi","Binance","BitBay","Bitfinex","Bittrex","CoinBase","GateIO","HitBTC","Huobi","Kraken","LBank","OKEx","Quoine","Simex","ZB","Bancor","Loopring","SushiSwap","Dforce","0x","Kyber","Bitmax","PanCakeSwap","CREX24","STEX"]
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="get" path="/v1/assetMarkets" baseUrl="https://api.diadata.org" summary="Asset Markets" %}
{% swagger-description %}
Returns all observed volumes in the last 24h for a specified asset.

Example: https://api.diadata.org/v1/assetMarkets/Bitcoin/0x000000000000000000000000000000000000000
{% endswagger-description %}

{% swagger-parameter in="path" name="blockchain" required="true" %}
Name of the blockchain for requested asset
{% endswagger-parameter %}

{% swagger-parameter in="path" name="address" required="true" %}
Address of the requested asset
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="List of exchanges with volumes in the last 24h for all volmes that are > 0" %}
```javascript
[{"Exchange":"Binance","Volume":1199.1039199999916},{"Exchange":"BitBay","Volume":57.20761510999996},{"Exchange":"Bitfinex","Volume":8625.797360533787},{"Exchange":"Bitmax","Volume":388.3493600000159},{"Exchange":"Bittrex","Volume":194.73298052999988},{"Exchange":"ByBit","Volume":3621.601983133232},{"Exchange":"CREX24","Volume":0.3536900000000002},{"Exchange":"CoinBase","Volume":11682.646369501423},{"Exchange":"Crypto.com","Volume":5574.927184001124},{"Exchange":"FTX","Volume":9606.986100000662},{"Exchange":"GateIO","Volume":772.3802987573356},{"Exchange":"HitBTC","Volume":3.6027500000000017},{"Exchange":"Huobi","Volume":8315.162517776456},{"Exchange":"Kraken","Volume":590.9460117199801},{"Exchange":"LBank","Volume":6632.81796153011},{"Exchange":"OKEx","Volume":650.4299946699018},{"Exchange":"Quoine","Volume":495.0948736999894},{"Exchange":"STEX","Volume":0.6928773800000005},{"Exchange":"ZB","Volume":907.4366999999507}]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/chartPoints/:filter/:exchange/:symbol" method="get" summary="Chart Points" %}
{% swagger-description %}
Get chart points for an exchange.

\


https://api.diadata.org/v1/chartPoints/MEDIR120/Binance/BTC

\




\


For a list of available exchanges see:

\


https://docs.diadata.org/documentation/api-1#api-access

\


or:

\


https://docs.diadata.org/documentation/api-1/api-endpoints#exchanges

\




\




\




\




_Remark_

: Successful responses can be rather large.
{% endswagger-description %}

{% swagger-parameter in="path" name="filter" type="string" %}
Which filter should be applied (Available options: MEDIR120 and MAIR120).
{% endswagger-parameter %}

{% swagger-parameter in="path" name="exchange" type="string" %}
Which exchange to use.
{% endswagger-parameter %}

{% swagger-parameter in="path" name="symbol" type="string" %}
A valid symbol from GET /v1/coins, e.g., BTC.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="scale" type="string" %}
Which scale the graph points distance should have. Available options: 5m 30m 1h 4h 1d 1w.
{% endswagger-parameter %}

{% swagger-response status="200" description="Successful retrieval of a chart point." %}
```
{"DataPoints":[{"Series":[{"name":"filters","columns":["time","exchange","filter","symbol","value"],"values":[["2020-05-19T08:02:09Z","GateIO","MEDIR120","EOS",2.6218717017500084]]}],"Messages":null}]}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/chartPointsAllExchanges/:filter/:symbol" method="get" summary="Chart Points for all Exchanges" %}
{% swagger-description %}
Get symbol details for all exchanges.

\




_Example_

:

\


https://api.diadata.org/v1/chartPointsAllExchanges/MEDIR120/EOS

\




\




_Remark:_

 Careful! Successful responses can be rather large.
{% endswagger-description %}

{% swagger-parameter in="path" name="starttime" type="integer" %}
Unix timestamp setting the start of the return array
{% endswagger-parameter %}

{% swagger-parameter in="path" name="endtime" type="integer" %}
Unix timestamp setting the end of the return array
{% endswagger-parameter %}

{% swagger-parameter in="path" name="filter" type="string" %}
Which filter should be applied (Available options: MEDIR120 and MAIR120).
{% endswagger-parameter %}

{% swagger-parameter in="path" name="symbol" type="string" %}
A valid symbol from GET /v1/coins, e.g., BTC.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="scale" type="string" %}
Which scale the graph points distance should have. Available options: 5m 30m 1h 4h 1d 1w
{% endswagger-parameter %}

{% swagger-response status="200" description="Successful retrieval of a chart point for all exchanges." %}
```
{"DataPoints":[{"Series":[{"name":"filters","columns":["time","exchange","filter","symbol","value"],"values":[["2020-05-19T08:17:59Z",null,"MEDIR120","EOS",2.6236194301032314]]}],"Messages":null}]}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/supply/:symbol" method="get" summary="Supply" %}
{% swagger-description %}
Get the current circulating supply for the token corresponding to symbol.

\




_Example_

:

\


https://api.diadata.org/v1/supply/BTC
{% endswagger-description %}

{% swagger-parameter in="path" name="symbol" type="string" %}
Which symbol to get the supply for, e.g., BTC
{% endswagger-parameter %}

{% swagger-response status="200" description="Successful retrieval of BTC supply." %}
```
{"Symbol":"BTC","Name":"Bitcoin","CirculatingSupply":17655550,"Source":"diadata.org","Time":"2019-04-20T08:44:25.748170404Z","Block":0}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/supplies/:symbol" method="get" summary="Supplies" %}
{% swagger-description %}
Get all recorded supply values for the token corresponding to symbol.

\




_Example_

:

\


https://api.diadata.org/v1/supplies/BTC

\




\


Get supply values for a time range using the query parameters.

\




_Example_

:

\


https://api.diadata.org/v1/supplies/BTC?starttime=1602232273&endtime=1602318673

\



{% endswagger-description %}

{% swagger-parameter in="path" name="symbol" type="string" %}
Which symbol to get the supply fot, e.g., BTC
{% endswagger-parameter %}

{% swagger-parameter in="query" name="starttime" type="integer" %}
Unix timestamp setting the start of the return array
{% endswagger-parameter %}

{% swagger-parameter in="query" name="endtime" type="integer" %}
Unix timestamp setting the end of the return array
{% endswagger-parameter %}

{% swagger-response status="200" description="Successful retrieval of two supply values for Bitcoin (BTC) between timestamps 1591700000 and 1591883936." %}
```
[{"Symbol":"BTC","Name":"Bitcoin","CirculatingSupply":18399687,"Source":"diadata.org","Time":"2020-06-09T23:59:59Z","Block":0},{"Symbol":"BTC","Name":"Bitcoin","CirculatingSupply":18400712,"Source":"diadata.org","Time":"2020-06-10T23:59:59Z","Block":0}]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/symbol/:symbol" method="get" summary="Symbol" %}
{% swagger-description %}
Get extensive information on the cryptocurrency corresponding to symbol on various exchanges.
{% endswagger-description %}

{% swagger-parameter in="path" name="symbol" type="string" %}
Which symbol to get the details on, e.g., BTC
{% endswagger-parameter %}

{% swagger-response status="200" description="Information on the cryptocurrency organized by "Change", "Coin", "Rank", "Exchanges" and "Gfx1"  (filtered data). Shown below is an excerpt of a successful response of symbol = BTC." %}
```
"Change":{"USD":[{"Symbol":"EUR","Rate":0.8995232526760818,"RateYesterday":0.8995232526760818},...

"Coin":{"Symbol":"BTC","Name":"Bitcoin","Price":9780.807149999986,"PriceYesterday":9574.416265039981,"VolumeYesterdayUSD":354341949.0902907,"Time":"2020-05-19T10:13:22.895692183Z","CirculatingSupply":17655550},...

"Rank":1

"Exchanges":[{"Name":"Huobi","Price":9776.344026379707,"PriceYesterday":9566.082031390646,"VolumeYesterdayUSD":182131794.24870485,"Time":"2020-05-19T10:07:59Z","LastTrades":...

"Gfx1":{"DataPoints":[{"Series":[{"name":"filters","columns":["time","exchange","filter","symbol","value"],"values":[["2020-05-19T10:08:00Z",null,"MA120","BTC",9780.807149999986],...
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/coins" method="get" summary="Coins" %}
{% swagger-description %}
Get a list of all available coins.
{% endswagger-description %}

{% swagger-response status="200" description="Successful retrieval of available coins along with actual information on prices. Shown below is an exerpt of the full response." %}
```
"CompleteCoinList":[{"Symbol":"BTC","Name":"Bitcoin"},{"Symbol":"ETH","Name":"Ethereum"},...

"Change":{"USD":[{"Symbol":"EUR","Rate":0.8995232526760818,"RateYesterday":0.8995232526760818},...

"Coins":[{"Symbol":"BTC","Name":"Bitcoin","Price":9773.78345474998,"PriceYesterday":9574.416265039981,"VolumeYesterdayUSD":352085287.0431704,"Time":"2020-05-19T10:05:53.191886175Z","CirculatingSupply":17655550},...
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/volume/:symbol" method="get" summary="Trade Volume" %}
{% swagger-description %}
Get the trading volume of the specified symbol in a defined time span.
{% endswagger-description %}

{% swagger-parameter in="path" name="symbol" type="string" %}
Which symbol to retrieve the volume of (e.g. BTC)
{% endswagger-parameter %}

{% swagger-parameter in="query" name="starttime" type="integer" %}
Start of the timespan (Unix time in seconds)
{% endswagger-parameter %}

{% swagger-parameter in="query" name="endtime" type="integer" %}
End of the timespan (Unix time in seconds)
{% endswagger-parameter %}

{% swagger-response status="200" description="An example response when querying a BTC volume for a typical day." %}
```
1431527525.7309263
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/cviIndex" method="get" summary="CVI Index " %}
{% swagger-description %}
Get all values of the Crypto Volatility Index.

\




_Example_

:

\


https://api.diadata.org/v1/cviIndex

\




\




_Example_

 with query parameters:

\


https://api.diadata.org/v1/cviIndex?starttime=1589829000&endtime=1589830000
{% endswagger-description %}

{% swagger-parameter in="query" name="starttime" type="integer" %}
Unix timestamp setting the start of the return array
{% endswagger-parameter %}

{% swagger-parameter in="query" name="endtime" type="integer" %}
Unix timestamp setting the end of the return array
{% endswagger-parameter %}

{% swagger-response status="200" description="Successful retrieval of CVI Index value for starttime=1589829000 and endtime=1589830000" %}
```
[{"Timestamp":"2020-05-18T19:12:43Z","Value":142.28101897342574},{"Timestamp":"2020-05-18T19:17:48Z","Value":142.29282246717017},{"Timestamp":"2020-05-18T19:22:51Z","Value":142.3025697159107}]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/index/:symbol" method="get" summary="Crypto Index" %}
{% swagger-description %}
Returns information about the cryptoindex indicated by its symbol. This included price and market data, as well as a list of its constituents.
{% endswagger-description %}

{% swagger-parameter in="path" name="symbol" type="string" %}
Symbol of the index
{% endswagger-parameter %}

{% swagger-parameter in="path" name="starttime" type="integer" %}
Unix timestamp setting the start of the return array
{% endswagger-parameter %}

{% swagger-parameter in="path" name="endtime" type="string" %}
Unix timestamp setting the end of the return array
{% endswagger-parameter %}

{% swagger-response status="200" description="Successful retrieval of the index value for the GBI index" %}
```
[{"Name":"GBI","Value":97.50947578961045,"Price":1.00778406966601,"Price1h":1.00778406966601,"Price24h":1.0051799695772836,"Price7d":0.9673670995027654,"Price14d":1.0893058979363988,"Price30d":0,"Volume24hUSD":0,"CirculatingSupply":345195.6232478934,"Divisor":1,"CalculationTime":"2021-03-05T13:14:45Z","Constituents":[{"Name":"WBTC","Symbol":"WBTC","Address":"-","Price":47979.76523333797,"PriceYesterday":49023.163422260404,"PriceYesterweek":46756.5210539339,"CirculatingSupply":58722.018321489995,"Weight":0.1392857142857143,"Percentage":0.15142956928279822,"CappingFactor":0,"NumBaseTokens":3077509414229.555},{"Name":"Ethereum","Symbol":"ETH","Address":"-","Price":1487.6297917225763,"PriceYesterday":1554.7081046080527,"PriceYesterweek":1472.0938093047234,"CirculatingSupply":111297265,"Weight":0.1392857142857143,"Percentage":0.12197452878400614,"CappingFactor":0,"NumBaseTokens":79950485178446.9},{"Name":"YFI","Symbol":"YFI","Address":"-","Price":31607.277099187264,"PriceYesterday":32633.133137109693,"PriceYesterweek":31803.892866435,"CirculatingSupply":36666,"Weight":0.1392857142857143,"Percentage":0.12016254448793848,"CappingFactor":0,"NumBaseTokens":3707053500937.577},{"Name":"Uniswap","Symbol":"UNI","Address":"-","Price":27.55576312916283,"PriceYesterday":29.00379988760186,"PriceYesterweek":22.688371324999938,"CirculatingSupply":591130752.8153942,"Weight":0.1392857142857143,"Percentage":0.19239736694689585,"CappingFactor":0,"NumBaseTokens":6808218776724222},{"Name":"Compound Coin","Symbol":"COMP","Address":"-","Price":455.71303136656303,"PriceYesterday":476.00878227409333,"PriceYesterweek":394.35936478768855,"CirculatingSupply":4347413.630491343,"Weight":0.1392857142857143,"Percentage":0.13561391266087963,"CappingFactor":0,"NumBaseTokens":290174750844543.1},{"Name":"Maker","Symbol":"MKR","Address":"-","Price":2127.71418543,"PriceYesterday":2146.1291527178396,"PriceYesterweek":2014.4420156976037,"CirculatingSupply":902135,"Weight":0.1392857142857143,"Percentage":0.12093671724782111,"CappingFactor":0,"NumBaseTokens":55423214185923.1},{"Name":"Chainlink","Symbol":"LINK","Address":"-","Price":26.807735407371414,"PriceYesterday":28.994353119130064,"PriceYesterweek":25.14386537072962,"CirculatingSupply":410509556.43444455,"Weight":0.1392857142857143,"Percentage":0.1391510991421949,"CappingFactor":0,"NumBaseTokens":5061431160340588},{"Name":"SPICE","Symbol":"SPICE","Address":"-","Price":1.2979738789219422,"PriceYesterday":1.5297308326069865,"PriceYesterweek":0.8919250424836395,"CirculatingSupply":1945426.8,"Weight":0.025,"Percentage":0.018334261447465704,"CappingFactor":0,"NumBaseTokens":13773499234182648}]}]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org/v1/" path="foreignSymbols/:source" method="get" summary="Guest Symbols" %}
{% swagger-description %}
Get the list of available symbols along with their ITIN for guest quotations.

\




_Example_

:

\


https://api.diadata.org/v1/foreignSymbols/Coingecko
{% endswagger-description %}

{% swagger-parameter in="path" name="source" type="string" %}
source of the quotation
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org/v1/foreignQuotation/:source/:symbol" path="" method="get" summary="Guest Quotation" %}
{% swagger-description %}
Get the latest quotation for a token from a guest source.

\




_Example_

:

\


https://api.diadata.org/v1/foreignQuotation/CoinMarketCap/BTC

\




\


Use the query parameter time in order to get the latest quotation before the specified timestamp.

\




_Example_

:

\


https://api.diadata.org/v1/foreignQuotation/Coingecko/BTC?time=1601351679

\




\



{% endswagger-description %}

{% swagger-parameter in="path" name="source" type="string" %}
source of the quotation
{% endswagger-parameter %}

{% swagger-parameter in="path" name="symbol" type="string" %}
Which symbol to get a quotation for, e.g. BTC
{% endswagger-parameter %}

{% swagger-parameter in="query" name="time" type="number" %}
Unix timestamp.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}

## Traditional Assets

{% swagger baseUrl="https://api.diadata.org/v1/stockQuotation/:" path="source/:symbol/:time" method="get" summary="Stock Quotation" %}
{% swagger-description %}
Get a stock quotation.

\




_Example_

:

\


https://api.diadata.org/v1/stockQuotation/Finage/AAPL

\




\


Get stock quotations for a time range using the query parameters.

\




_Example_

:

\


https://api.diadata.org/v1/stockQuotation/Finage/AAPL?dateInit=1633343956&dateFinal=1633345556
{% endswagger-description %}

{% swagger-parameter in="path" name="source" type="string" %}
Data source of the quotation.
{% endswagger-parameter %}

{% swagger-parameter in="path" name="symbol" type="string" %}
Symbol of the stock, see stockSymbols endpoint below.

\



{% endswagger-parameter %}

{% swagger-parameter in="query" name="dateInit" type="integer" %}
Initial timestamp for range queries.

\


Format: Unix timestamp.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="dateFinal" type="integer" %}
Final timestamp for range queries.

\


Format: Unix timestamp.
{% endswagger-parameter %}

{% swagger-response status="200" description="Successful retrieval of the latest stock quotation for Apple stock on Finage." %}
```
{"Symbol":"AAPL","Name":"APPLE","PriceAsk":141.55,"PriceBid":141.52,"SizeAskLot":2,"SizeBidLot":10,"Source":"Finage","Time":"2021-10-04T11:47:42Z","ISIN":"US0378331005"}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org/v1/stockSymbols" path="" method="get" summary="Stocks" %}
{% swagger-description %}
Get a list of stocks available for quotation. The field source shows for which source the stock's quotations are available.
{% endswagger-description %}

{% swagger-parameter in="path" name="" type="string" %}

{% endswagger-parameter %}

{% swagger-response status="200" description="Successful retrieval of stocks available for quotation." %}
```
[{"Stock":{"Symbol":"MSFT","Name":"MICROSOFT CORP","ISIN":"US5949181045"},"Source":"Finage"},{"Stock":{"Symbol":"AAPL","Name":"APPLE","ISIN":"US0378331005"},"Source":"Finage"}]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/interestrates" method="get" summary="Interest Rates" %}
{% swagger-description %}
Get a list of all available interest rates along with metadata on the rates such as first publication date and issuing entity.

\


https://api.diadata.org/v1/interestrates
{% endswagger-description %}

{% swagger-response status="200" description="Successful retrieval of meta information on available interest rates." %}
```
[{"Symbol":"ESTER","FirstDate":"2019-10-01T00:00:00Z","Issuer":"ECB"},{"Symbol":"SOFR90","FirstDate":"2020-03-02T00:00:00Z","Issuer":"FED"},{"Symbol":"SONIA","FirstDate":"1997-01-02T00:00:00Z","Issuer":"BOE"},{"Symbol":"SAFR","FirstDate":"2020-03-02T00:00:00Z","Issuer":"FED"},{"Symbol":"SOFR","FirstDate":"2018-04-02T00:00:00Z","Issuer":"FED"},{"Symbol":"SOFR180","FirstDate":"2020-03-02T00:00:00Z","Issuer":"FED"},{"Symbol":"SOFR30","FirstDate":"2020-03-02T00:00:00Z","Issuer":"FED"}]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/interestrate/:rateType" method="get" summary="Interest Rate" %}
{% swagger-description %}
Get value for a certain rate type.

\




_Example_

:

\


https://api.diadata.org/v1/interestrate/ESTER/2020-04-20​

\




\


Get rate values for a range of timestamps using optional query parameters.

\




_Example_

:

\


https://api.diadata.org/v1/interestrate/ESTER?dateInit=2020-02-20&dateFinal=2020-04-16
{% endswagger-description %}

{% swagger-parameter in="path" name="rateType" type="string" %}
Symbol name for a rate.
{% endswagger-parameter %}

{% swagger-parameter in="path" name="date" type="string" %}
Return the rate for the specified date. Default date is the latest available date. Format: yyyy-mm-dd
{% endswagger-parameter %}

{% swagger-parameter in="query" name="dateInit" type="string" %}
Initial date for range queries. Format yyyy-mm-dd
{% endswagger-parameter %}

{% swagger-parameter in="query" name="dateFinal" type="string" %}
Final date for range queries. Format: yyyy-mm-dd
{% endswagger-parameter %}

{% swagger-response status="200" description="Successful retrieval of an interest rate." %}
```
{"Symbol":"ESTER","Value":-0.542,"PublicationTime":"2020-05-19T07:15:07Z","EffectiveDate":"2020-05-18T00:00:00Z","Source":"ECB"}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/compoundedRate/:rateType/:dpy/:date" method="get" summary="Compounded Index" %}
{% swagger-description %}
Get the value of an index compounded since its first publication date.

\




\




_Example_

:

\


https://api.diadata.org/v1/compoundedRate/SOFR/360/2020-05-14

\




\


Get the compounded index for a range of dates using the query parameters.

\




_Example_

:

\


https://api.diadata.org/v1/compoundedRate/SOFR/360?dateInit=2020-04-24&dateFinal=2020-05-14

\




\


For the methodology of compounded rates see:

\


https://docs.diadata.org/documentation/methodology/traditional-assets/compounded-rates#standard-methodology
{% endswagger-description %}

{% swagger-parameter in="path" name="rateType" type="string" %}
Symbol for a rate name
{% endswagger-parameter %}

{% swagger-parameter in="path" name="dpy" type="integer" %}
Business day convention for the number of days per year
{% endswagger-parameter %}

{% swagger-parameter in="path" name="date" type="string" %}
Return the compounded index for the date specified in the format yyyy-mm-dd
{% endswagger-parameter %}

{% swagger-parameter in="query" name="dateInit" type="string" %}
Initial date for range queries. Format: yyyy-mm-dd
{% endswagger-parameter %}

{% swagger-parameter in="query" name="dateFinal" type="string" %}
Final date for range queries. Format: yyyy-mm-dd
{% endswagger-parameter %}

{% swagger-response status="200" description="Successful retrieval of the SOFR Index." %}
```
{"Symbol":"SOFR_compounded_by_DIA","Value":1.0414032009923273,"PublicationTime":"0001-01-01T00:00:00Z","EffectiveDate":"2020-05-14T00:00:00Z","Source":"FED"}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/compoundedAvg/:rateType/:period/:dpy/:date" method="get" summary="Compounded Average" %}
{% swagger-description %}
Get the average value of a given interest rate compounded over a period of time.

\




_Example_

:

\


https://api.diadata.org/v1/compoundedAvg/SOFR/30/360/2020-05-14

\




\


Get the compounded averages for a range of dates using the query parameters.

\




_Example_

:

\


https://api.diadata.org/v1/compoundedAvg/SOFR/30/360?dateInit=2020-04-24&dateFinal=2020-05-14

\




\


For the methodology see:

\


https://docs.diadata.org/documentation/methodology/traditional-assets/compounded-rates#standard-methodology

\




\




_Remark_

: This Get method requires an API key. Please contact us for more information:

\


https://docs.google.com/forms/d/e/1FAIpQLSePxDwbEURjes4nw8GUzaT-XfYttRw_6F2xAR607FKACsn7ew/viewform

\



{% endswagger-description %}

{% swagger-parameter in="path" name="rateType" type="string" %}
Symbol for a rate name
{% endswagger-parameter %}

{% swagger-parameter in="path" name="period" type="integer" %}
Rate is compounded over period days
{% endswagger-parameter %}

{% swagger-parameter in="path" name="dpy" type="integer" %}
Business day convention for the number of days per year
{% endswagger-parameter %}

{% swagger-parameter in="path" name="date" type="string" %}
Return the compounded rate for the date specified in the format yyyy-mm-dd
{% endswagger-parameter %}

{% swagger-parameter in="query" name="dateInit" type="string" %}
Initial date for range queries. Format: yyyy-mm-dd
{% endswagger-parameter %}

{% swagger-parameter in="query" name="dateFinal" type="string" %}
Final date for range queries. Format: yyyy-mm-dd
{% endswagger-parameter %}

{% swagger-response status="200" description="Successful retrieval of a compounded average of the SOFR over an interest period of 30 calendar days." %}
```
{"Symbol":"SOFR30_compounded_by_DIA","Value":0.035667157687857554,"PublicationTime":"0001-01-01T00:00:00Z","EffectiveDate":"2020-05-14T00:00:00Z","Source":"FED"}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/compoundedAvgDIA/:rateType/:period/:dpy/:date" method="get" summary="Compounded Average using DIA Method" %}
{% swagger-description %}
Get the average value of an interest rate compounded over a period of time. Here, we use the DIA methodology for compounding the rate, i.e. interest is compounded for non-business days as well. For details see:

\


https://docs.diadata.org/documentation/methodology/traditional-assets/compounded-rates#dia-methodology

\




\




_Example_

:

\


https://api.diadata.org/v1/compoundedAvgDIA/SOFR/30/360/2020-05-14

\




\


Get the compounded average for a range of dates using the query parameters. 

\




_Example_

:

\


https://api.diadata.org/v1/compoundedAvgDIA/SOFR/30/360?dateInit=2020-04-24&dateFinal=2020-05-14

\




\




_Remark_

: This Get method requires an API key. Please contact us for more information:

\


https://docs.google.com/forms/d/e/1FAIpQLSePxDwbEURjes4nw8GUzaT-XfYttRw_6F2xAR607FKACsn7ew/viewform

\



{% endswagger-description %}

{% swagger-parameter in="path" name="rateType" type="string" %}
Symbol for a rate name
{% endswagger-parameter %}

{% swagger-parameter in="path" name="period" type="integer" %}
Rate is compounded over period days
{% endswagger-parameter %}

{% swagger-parameter in="path" name="dpy" type="integer" %}
Business convention for the number of days per year
{% endswagger-parameter %}

{% swagger-parameter in="path" name="date" type="string" %}
Return the compounded rate for the date specified in the format yyyy-mm-dd
{% endswagger-parameter %}

{% swagger-parameter in="query" name="dateInit" type="string" %}
Initial date for range queries. Format: yyyy-mm-dd
{% endswagger-parameter %}

{% swagger-parameter in="query" name="dateFinal" type="string" %}
Final date for range queries. Format: yyyy-mm-dd
{% endswagger-parameter %}

{% swagger-response status="200" description="Successful retrieval of the compounded average of SOFR over an interest period of 30 calendar days." %}
```
[{"Symbol":"SOFR30_compounded_by_DIA","Value":0.035667175187725775,"PublicationTime":"0001-01-01T00:00:00Z","EffectiveDate":"2020-05-14T00:00:00Z","Source":"FED"}]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org/v1/" path="fiatQuotations" method="get" summary="Fiat Currency Exchange Rates" %}
{% swagger-description %}
Get a list of exchange rates for several fiat currencies vs US Dollar.
{% endswagger-description %}

{% swagger-parameter in="path" name="" type="string" %}

{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org/v1/" path="goldPaxgGrams" method="get" summary="Gold price in Gram" %}
{% swagger-description %}
Gold price for 1g of Gold measured by the PAXG commodity token.
{% endswagger-description %}

{% swagger-response status="200" description="" %}
```
{"Symbol":"PAXG-gram","Name":"PAXG-gram","Price":59.69023528449715,"PriceYesterday":57.93549261152835,"VolumeYesterdayUSD":0,"Source":"diadata.org","Time":"2020-11-25T11:22:31.146028646Z","ITIN":"undefined"}
```
{% endswagger-response %}
{% endswagger %}

