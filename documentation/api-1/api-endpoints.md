---
description: >-
  The DIA base url is https://api.diadata.org. All API paths are sub-paths of
  this base URL.
---

# API Endpoints

{% api-method method="get" host="https://api.diadata.org" path="/v1/chartPoints" %}
{% api-method-summary %}
Chart Points
{% endapi-method-summary %}

{% api-method-description %}
Get chart points for an exchange.  
Example: https://api.diadata.org/v1/chartPoints/MEDIR120/GateIO/EOS  
  
For a list of available trading places see:  
https://docs.diadata.org/documentation/api-1\#api-access  
  
Remark: Successful responses can be rather large.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="filter" type="string" required=true %}
Which filter should be applied \(Available options: MEDIAR120 and MAIR120\).
{% endapi-method-parameter %}

{% api-method-parameter name="exchange" type="string" required=true %}
Which exchange to use.
{% endapi-method-parameter %}

{% api-method-parameter name="symbol" type="string" required=true %}
A valid symbol from GET /v1/coins, e.g., BTC
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
Cake successfully retrieved.
{% endapi-method-response-example-description %}

```
{"DataPoints":[{"Series":[{"name":"filters","columns":["time","exchange","filter","symbol","value"],"values":[["2020-05-19T08:02:09Z","GateIO","MEDIR120","EOS",2.6218717017500084]]}],"Messages":null}]}
```
{% endapi-method-response-example %}

{% api-method-response-example httpCode=404 %}
{% api-method-response-example-description %}
Could not find a cake matching this query.
{% endapi-method-response-example-description %}

```
{    "message": "Error 404"}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}



