# API

## Base URL

The DIA base url is `https://api.diadata.org/`. All API paths are sub-paths of this base URL.

## Paths

### GET /v1/chartPoints/

Get chart points for an exchange.  
Example: [https://api.diadata.org/v1/chartPoints/MEDIR120/GateIO/EOS](https://api.diadata.org/v1/chartPoints/MEDIR120/GateIO/EOS)

For a list of available trading places see:  
[https://docs.diadata.org/documentation/api-1\#api-access](https://docs.diadata.org/documentation/api-1#api-access)

Query Params:

* scale \[string\]: scale 5m 30m 1h 4h 1d 1w.

Path Params:

* filter \[string\]: Some filter. \(for now MEDIR120 or MAIR120\)
* trading place \[string\]: Some trading place.
* symbol \[string\]: Some symbol from GET /v1/coins

_Remark:_ Careful! Successful responses can be rather large.

### GET /v1/chartPointsAllExchanges/

Get Symbol Details.  
Example: [https://api.diadata.org/v1/chartPointsAllExchanges/MEDIR120/EOS](https://api.diadata.org/v1/chartPointsAllExchanges/MEDIR120/EOS)  


Query Params:

* scale \[string\]: scale 5m 30m 1h 4h 1d 1w.

Path Params:

* filter \[string\]: Some filter. \(for now MEDIR120 or MAIR120\)
* symbol \[string\]: Some symbol.

_Remark:_ Careful! Successful responses can be rather large.

### GET /v1/interestrate/

Get value for a certain rate type.  
Example: [https://api.diadata.org/v1/interestrate/ESTER/2020-04-20](https://api.diadata.org/v1/interestrate/ESTER/2020-04-20)

Get rate values for a range of timestamps using optional query parameters.  
Example: [https://api.diadata.org/v1/interestrate/ESTER?dateInit=2020-02-20&dateFinal=2020-04-16](https://api.diadata.org/v1/interestrate/ESTER?dateInit=2020-02-20&dateFinal=2020-04-16)

Path Params:

* rateType \[string\]: Short hand notation/symbol for a rate
* date \[string\]: In the format yyyy:mm:dd  date is an optional parameter. When omitted, the most recent value is returned.

Optional Query Params:

* dateInit, dateFinal \[string\]: In the format yyyy:mm:dd

### GET /v1/quotation/

Get a quotation.  
Example: [https://api.diadata.org/v1/quotation/ETH](https://api.diadata.org/v1/quotation/ETH)

Path Params:

* symbol \[string\]: Some symbol.

### GET /v1/supply/

Get the circulating supply corresponding to a symbol.  
Example: [https://api.diadata.org/v1/supply/ETH](https://api.diadata.org/v1/supply/ETH)

Path Params:

* symbol \[string\]: Some symbol.

### POST /v1/supply

Post the circulating supply.

Query Params:

* Symbol \[string\]: Coin symbol.
* CirculatingSupply \[number\]: number of coins in circulating supply.

### GET /v1/symbol/

Get Symbol Details.  
Example: [https://api.diadata.org/v1/symbol/ETH](https://api.diadata.org/v1/symbol/ETH)

Path Params:

* symbol \[string\]: Some symbol.

### GET /v1/cviIndex

Get all values of the Crypto Volatility Index.  
Example: [https://api.diadata.org/v1/cviIndex](https://api.diadata.org/v1/cviIndex)

* Parameters: starttime \[int\]: Unix timestamp where the array values should begin, endtime \[int\] Unix timestamp where the array should end

### GET /v1/coins

Get a list of all available coins.  
Example: [https://api.diadata.org/v1/coins](https://api.diadata.org/v1/coins)

### GET /v1/exchanges

Get a list of all available trading places.  
Example: [https://api.diadata.org/v1/exchanges](https://api.diadata.org/v1/exchanges)

### GET /v1/interestrates

Get a list of all available interest rates along with some metadata on the rate such as first publication date and Issuing entity.  
Example: [https://api.diadata.org/v1/interestrates](https://api.diadata.org/v1/interestrates)

### GET /v1/pairs/

Get a list of all available pairs.  
Example: [https://api.diadata.org/v1/pairs](https://api.diadata.org/v1/pairs)

### GET /v1/symbols

Get a list of all available symbols.  
Example: [https://api.diadata.org/v1/symbols](https://api.diadata.org/v1/symbols)

### Responses for all GET  requests:

* 200: success.

  Return the respective JSON object

* 404: Symbol not found.

  [restApi.APIError](./#restapiapierror)

* 500: error.

  [restApi.APIError](./#restapiapierror)

## Use cases

### Bash scripting

The API can be accessed through a Linux terminal by using curl. For example  
`curl https://api.diadata.org/v1/interestrate/ESTER/2020-03-16 >> userPath/myFile.txt`  
writes the return value of the GET request into `myFile.txt` for further processing.

### Usage with Python

The JSON object obtained in an API GET request complies with Python syntax. It can be cast into a list or dictionary resp. using Python's `eval(string)` function.

