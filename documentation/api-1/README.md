# API Documentation

**Find the right data for your needs**  
Show your users the most transparent data on the market with our API. Whether you're building a financial service, a portfolio management tool, a new media offering, or more, we have the most advanced and updated data on the market for your product.  
For Oracle usage see [Oracles in DIA](https://docs.diadata.org/documentation/api-1/oracles).

**Backtest your strategies**  
Use the most efficient and transparent crypto data to run simulations and backtest your trading or investing strategies. With crowd-aggregated hundreds of exchanges you can be sure that you're getting the right picture every single time.

**Run Experiments**  
Build your own models with our data, to further your interest or just for fun. With our flexible and powerful API, we provide you with a set of data that will help you draw insights and make conclusions.

**Request your data**  
Set a bounty on gitcoin.io or drop us a [line](mailto:API@diadata.org).

Version: 1.0

## API Access

The DIA base url is `https://api.diadata.org/v1`. All API paths are sub-paths of this base URL. You can find specific documentation for the endpoints of our API on the [API documentation site](https://docs.diadata.org/documentation/api-1/api-endpoints). 

Currently, DIA collects reference rates from traditional financial markets and data from several cryptocurrency exchanges \(see the table below\).

| Exchange |  |  |  |
| :--- | :--- | :--- | :--- |
| Number of pairs |  |  |  |

| Name | Number of Pairs |
| :--- | :--- |
| Binance | [47](https://github.com/diadata-org/diadata/tree/c982072de2ac488c5f0bdf32b677cbac1965583e/documentation/config/Binance.json) |
| Bitfinex | [25](https://github.com/diadata-org/diadata/tree/c982072de2ac488c5f0bdf32b677cbac1965583e/documentation/config/Bitfinex.json) |
| Coinbase | [3](https://github.com/diadata-org/diadata/tree/c982072de2ac488c5f0bdf32b677cbac1965583e/documentation/config/CoinBase.json) |
| GateIO | [42](https://github.com/diadata-org/diadata/tree/c982072de2ac488c5f0bdf32b677cbac1965583e/documentation/config/GateIO.json) |
| HitBTC | [43](https://github.com/diadata-org/diadata/tree/c982072de2ac488c5f0bdf32b677cbac1965583e/documentation/config/HitBTC.json) |
| Huobi | [40](https://github.com/diadata-org/diadata/tree/c982072de2ac488c5f0bdf32b677cbac1965583e/documentation/config/Huobi.json) |
| Kraken | [8](https://github.com/diadata-org/diadata/tree/c982072de2ac488c5f0bdf32b677cbac1965583e/documentation/config/Kraken.json) |
| LBank | [20](https://github.com/diadata-org/diadata/tree/c982072de2ac488c5f0bdf32b677cbac1965583e/documentation/config/LBank.json) |
| OKEx | [27](https://github.com/diadata-org/diadata/tree/c982072de2ac488c5f0bdf32b677cbac1965583e/documentation/config/OKEx.json) |
| Simex | [14](https://github.com/diadata-org/diadata/tree/c982072de2ac488c5f0bdf32b677cbac1965583e/documentation/config/Simex.json) |
| ZB.com | [127](https://github.com/diadata-org/diadata/tree/c982072de2ac488c5f0bdf32b677cbac1965583e/documentation/config/ZB.json) |

## Use cases

### Bash scripting

The API can be accessed through a Linux terminal by using curl. For example  
`curl https://api.diadata.org/v1/interestrate/ESTER/2020-03-16 >> userPath/myFile.txt`  
writes the return value of the GET request into `myFile.txt` for further processing.

### Usage with Python

The JSON object obtained in an API GET request complies with Python syntax. It can be cast into a list or dictionary resp. using Python's `eval(string)` function.

