---
description: >-
  The DIA base url is https://api.diadata.org. All API paths are sub-paths of
  this base URL.
---

# API Endpoints

## Digital Assets

{% swagger method="get" path="v1/assetQuotation/:blockchain/:asset" baseUrl="https://api.diadata.org/" summary="Asset Quotation" %}
{% swagger-description %}
Returns the quotation for a fully qualified asset (i.e. distinguished by blockchain and address).

_Example:_ [_https://api.diadata.org/v1/assetQuotation/Bitcoin/0x0000000000000000000000000000000000000000_](https://api.diadata.org/v1/assetQuotation/Bitcoin/0x0000000000000000000000000000000000000000)__
{% endswagger-description %}

{% swagger-parameter in="path" name="blockchain" required="true" %}
Name of the blockchain for requested asset
{% endswagger-parameter %}

{% swagger-parameter in="path" name="asset" required="true" %}
Address of the requested asset
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Return of asset price action information" %}
```javascript
{
    // Response
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="get" path="/v1/assetMarkets/:blockchain/:address" baseUrl="https://api.diadata.org" summary="Asset Markets" %}
{% swagger-description %}
Returns all observed volumes in the last 24h for a specified asset.

_Example:_ [_https://api.diadata.org/v1/assetMarkets/Bitcoin/0x000000000000000000000000000000000000000_](https://api.diadata.org/v1/assetMarkets/Bitcoin/0x000000000000000000000000000000000000000)__
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

{% swagger baseUrl="https://api.diadata.org" path="/v1/assetChartPoints/:filter/:blockchain/:address" method="get" summary="Asset Chart Points" %}
{% swagger-description %}
Get asset details for all exchanges.

_Example_:\
[https://api.diadata.org/v1/assetChartPoints/MA120/Bitcoin/0x0000000000000000000000000000000000000000](https://api.diadata.org/v1/assetChartPoints/MA120/Bitcoin/0x0000000000000000000000000000000000000000)

\
_Remark:_ Careful! Successful responses can be rather large.
{% endswagger-description %}

{% swagger-parameter in="path" name="filter" type="string" required="true" %}
Which filter should be applied (Available options: MA120, MEDIR120, VOL120 and MAIR120).
{% endswagger-parameter %}

{% swagger-parameter in="path" name="blockchain" type="string" required="true" %}
A valid blockchain from GET /v1/blockchains, e.g., Bitcoin.
{% endswagger-parameter %}

{% swagger-parameter in="path" name="address" required="true" %}
A valid asset address from GET /v1/token/:symbol, e.g., 0x000000000000000000000000000000000000000 for BTC.
{% endswagger-parameter %}

{% swagger-parameter in="path" name="starttime" type="integer" %}
Unix timestamp setting the start of the return array
{% endswagger-parameter %}

{% swagger-parameter in="path" name="endtime" type="integer" %}
Unix timestamp setting the end of the return array
{% endswagger-parameter %}

{% swagger-parameter in="query" name="scale" type="string" %}
Which scale the graph points distance should have. Available options: 5m 30m 1h 4h 1d 1w
{% endswagger-parameter %}

{% swagger-response status="200" description="Successful retrieval of a chart points for an asset" %}
```
{"DataPoints":[{"Series":[{"name":"filters","columns":["time","exchange","filter","symbol","value"],"values":[["2020-05-19T08:17:59Z",null,"MEDIR120","EOS",2.6236194301032314]]}],"Messages":null}]}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="get" path="/v1/lastTradesAsset/:blockchain/:address" baseUrl="https://api.diadata.org" summary="Asset Last Trades" %}
{% swagger-description %}
Get last trades for an asset.

_Example:_ [https://api.diadata.org/v1/lastTradesAsset/Bitcoin/0x0000000000000000000000000000000000000000](https://api.diadata.org/v1/lastTradesAsset/Bitcoin/0x0000000000000000000000000000000000000000)
{% endswagger-description %}

{% swagger-parameter in="path" name="blockchain" required="true" %}
A valid blockchain from GET /v1/blockchains, e.g., Bitcoin.
{% endswagger-parameter %}

{% swagger-parameter in="path" name="address" required="true" %}
A valid asset address from GET /v1/token/:symbol, e.g., 0x000000000000000000000000000000000000000 for BTC.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Succesful retrieval of last trades for an asset" %}
```javascript
{
    // Response
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="get" path="/v1/assetSupply/Ethereum/:address" baseUrl="https://api.diadata.org" summary="Asset Supply" %}
{% swagger-description %}
Get circulating and total supply for an asset.

_Example:_ [https://api.diadata.org/v1/assetSupply/Ethereum/0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9](https://api.diadata.org/v1/assetSupply/Ethereum/0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9)

_Note: Currently supports assets only from Ethereum blockchain_
{% endswagger-description %}

{% swagger-parameter in="path" name="address" %}
A valid asset address from GET /v1/token/:symbol, e.g., 0x000000000000000000000000000000000000000 for BTC.
{% endswagger-parameter %}
{% endswagger %}

{% swagger method="get" path="/v1/blockchains" baseUrl="https://api.diadata.org" summary="Blockchains" %}
{% swagger-description %}
Get a list of all available blockchains.
{% endswagger-description %}

{% swagger-response status="200: OK" description="" %}
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
{% endswagger-description %}

{% swagger-response status="200" description="Successful retrieval of available exchanges." %}
```
["KuCoin","Uniswap","Balancer","Maker","Gnosis","Curvefi","Binance","BitBay","Bitfinex","Bittrex","CoinBase","GateIO","HitBTC","Huobi","Kraken","LBank","OKEx","Quoine","Simex","ZB","Bancor","Loopring","SushiSwap","Dforce","0x","Kyber","Bitmax","PanCakeSwap","CREX24","STEX"]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/symbols" method="get" summary="Symbols" %}
{% swagger-description %}
Get a list of all available symbols for cryptocurrencies.\
\
Get symbols restricted to an exchange using the query parameter. (For the moment only for centralized exchanges).

_Example_: [https://api.diadata.org/v1/symbols?exchange=Kraken](https://api.diadata.org/v1/symbols?exchange=Kraken)\

{% endswagger-description %}

{% swagger-parameter in="query" name="exchange" type="string" %}
Name of the crypto exchange.
{% endswagger-parameter %}

{% swagger-parameter in="path" name="substring" %}
Search for coins that match a string, e.g. 

_BTC_

 will return BTC, BTCB and other assets that start with 

_BTC_

 letters.
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

_Example_: [https://api.diadata.org/v1/quotation/BTC](https://api.diadata.org/v1/quotation/BTC)
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

{% swagger method="get" path="/token/:symbol" baseUrl="https://api.diadata.org" summary="Tokens list" %}
{% swagger-description %}
Get a list of blockchains and addresses for all tokens that match the symbol
{% endswagger-description %}

{% swagger-parameter in="path" name="symbol" required="true" %}
Which symbol to get a quotation for, e.g., BTC.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```javascript
{
    // Response
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/chartPoints/:filter/:exchange/:symbol" method="get" summary="Exchange Chart Points" %}
{% swagger-description %}
Get chart points for an exchange.

_Example_: [https://api.diadata.org/v1/chartPoints/MEDIR120/Binance/BTC](https://api.diadata.org/v1/chartPoints/MEDIR120/Binance/BTC)\
\
_Note_: Successful responses can be rather large.
{% endswagger-description %}

{% swagger-parameter in="path" name="filter" type="string" required="true" %}
Which filter should be applied (Available options: MA120, VOL120, MEDIR120 and MAIR120).
{% endswagger-parameter %}

{% swagger-parameter in="path" name="exchange" type="string" required="true" %}
A valid exchange from GET /v1/exchanges, e.g., Binance
{% endswagger-parameter %}

{% swagger-parameter in="path" name="symbol" type="string" required="true" %}
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

{% swagger baseUrl="https://api.diadata.org" path="/v1/supply/:symbol" method="get" summary="Supply" %}
{% swagger-description %}
Get the current circulating supply for the token corresponding to symbol.

_Example_: [https://api.diadata.org/v1/supply/BTC](https://api.diadata.org/v1/supply/BTC)
{% endswagger-description %}

{% swagger-parameter in="path" name="symbol" type="string" required="true" %}
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

_Example_: [https://api.diadata.org/v1/supplies/BTC](https://api.diadata.org/v1/supplies/BTC)\
\
Get supply values for a time range using the query parameters.

_Example_: [https://api.diadata.org/v1/supplies/BTC?starttime=1647349656\&endtime=1650028056](https://api.diadata.org/v1/supplies/BTC?starttime=1647349656\&endtime=1650028056)\

{% endswagger-description %}

{% swagger-parameter in="path" name="symbol" type="string" required="true" %}
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

{% swagger baseUrl="https://api.diadata.org" path="/v1/index/:symbol" method="get" summary="Crypto Index" %}
{% swagger-description %}
Returns information about the cryptoindex indicated by its symbol. This included price and market data, as well as a list of its constituents.

_Example_: [https://api.diadata.org/v1/index/SCIFI](https://api.diadata.org/v1/index/SCIFI)
{% endswagger-description %}

{% swagger-parameter in="path" name="symbol" type="string" required="true" %}
Symbol of the index. GBI or SCIFI
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

{% swagger baseUrl="https://api.diadata.org" path="/v1/foreignSymbols/:source" method="get" summary="Guest Symbols" %}
{% swagger-description %}
Get the list of available symbols along with their ITIN for guest quotations.

_Example_: [https://api.diadata.org/v1/foreignSymbols/Coingecko](https://api.diadata.org/v1/foreignSymbols/Coingecko)
{% endswagger-description %}

{% swagger-parameter in="path" name="source" type="string" required="true" %}
source of the quotation
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.diadata.org" path="/v1/foreignQuotation/:source/:symbol" method="get" summary="Guest Quotation" %}
{% swagger-description %}
Get the latest quotation for a token from a guest source.

_Example_: [https://api.diadata.org/v1/foreignQuotation/CoinMarketCap/BTC](https://api.diadata.org/v1/foreignQuotation/CoinMarketCap/BTC)\
\
Use the query parameter time in order to get the latest quotation before the specified timestamp

\
_Example_: [https://api.diadata.org/v1/foreignQuotation/Coingecko/BTC?time=1647349656](https://api.diadata.org/v1/foreignQuotation/Coingecko/BTC?time=1647349656)\
\
\

{% endswagger-description %}

{% swagger-parameter in="path" name="source" type="string" required="true" %}
source of the quotation
{% endswagger-parameter %}

{% swagger-parameter in="path" name="symbol" type="string" required="true" %}
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

