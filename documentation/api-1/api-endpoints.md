---
description: >-
  The DIA base url is https://api.diadata.org. All API paths are sub-paths of
  this base URL.
---

# API Endpoints

{% api-method method="get" host="https://api.diadata.org" path="/v1/chartPoints/:filter/:exchange/:symbol" %}
{% api-method-summary %}
Chart Points
{% endapi-method-summary %}

{% api-method-description %}
Get chart points for an exchange.  
Example: https://api.diadata.org/v1/chartPointsAllExchanges/MEDIR120/EOS  
  
For a list of available trading places see:  
https://docs.diadata.org/documentation/api-1\#api-access  
  
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
Use this for range queries to determine at which date the query should begin. Forma: yyyy-mm-dd
{% endapi-method-parameter %}

{% api-method-parameter name="dateFinal" type="string" required=false %}
Use this for range queries to determine at which date the query should end. Format: yyyy-mm-dd
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

{% api-method method="get" host="https://api.diadata.org" path="/v1/quotation/:symbol" %}
{% api-method-summary %}
Quotation
{% endapi-method-summary %}

{% api-method-description %}

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
{"Symbol":"BTC","Name":"Bitcoin","Price":9777.19339776667,"PriceYesterday":9574.416265039981,"VolumeYesterdayUSD":298134760.8811487,"Source":"diadata.org","Time":"2020-05-19T08:41:12.499645584Z"}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}



