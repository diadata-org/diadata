---
description: Description of data retrieval and processing of cryptocurrency data
---

# Cryptocurrency Data Collection

## Trading Data

The raw trading data is the basis of all further derived quantities such as prices and circulating supply numbers for all cryptocurrency exchanges. We retrieve and store trading data from as close to the data source as possible and with the highest precision possible.

### Centralized Cryptocurrency  Exchanges

All trading data is retrieved through Rest APIs or websocket APIs  which are developed and maintained by the respective trading platform. See the below table for detailed information.   
We remark that by the very nature of a websocket API, there is no retrieval frequency but instead, trading data comes in continuously as it is produced.

| Exchange | API Type | API Link | Retrieval Period |
| :--- | :--- | :--- | :--- |
| Binance | Rest | [Binance Rest API Documentation](https://binance-docs.github.io/apidocs/spot/en/#introduction) | 1 sec. |
| Bitfinex | Websocket | [Bitfinex Websocket Documentation](https://docs.bitfinex.com/docs/ws-general) | - |
| Bittrex   | Rest | [Bittrex Rest API Documentation](https://bittrex.github.io/api/v1-1) | 7 sec. |
| Coinbase | Websocket | [Coinbase Websocket Documentation](https://docs.pro.coinbase.com/) | - |
| GateIO | Websocket | [GateIO Websocket Documentation](https://www.gate.io/docs/websocket/index.html) | - |
| HitBTC | Websocket | [HitBTC Websocket Documentation](https://api.hitbtc.com/#socket-api-reference) | - |
| Huobi | Websocket | [Huobi Websocket Documentation](https://huobiapi.github.io/docs/spot/v1/en/#websocket-market-data) | - |
| Kraken | Rest | [Kraken Rest API Documentation](https://www.kraken.com/features/api#public-market-data) | 3 min. |
| LBank | Websocket | [LBank Websocket Documentation](https://docs.lbkex.net/en/#websocket-api-market-data) | - |
| OKEx | Websocket | [OKEx Websocket Documentation](https://www.okex.com/docs/en/#ws_swap-README) | - |
| Quoine | Rest | [Quoine Rest API Documentation](https://developers.liquid.com/) | 1.5 sec.  |
| ZB | Websocket | [ZB Websocket Documentation](https://www.zb.com/api#WebSocket%20API) | - |

### Decentralized Cryptocurrency Exchanges \(DEXes\)

In contrast to centralized exchanges, in decentralized exchanges it is possible to retrieve trading data directly from the respective blockchain.

{% hint style="info" %}
 In order to supply data to our community as quickly as possible, we retrieved the trading data through different types of APIs for now \(see table below\). However, we remark that we are currently implementing data retrieval through the blockchain directly in order to reduce unnecessary dependencies. This concerns _all_ decentralized exchanges.
{% endhint %}

| Exchange | Data Retrieval | API Link | Blockchain | Update Period |
| :--- | :--- | :--- | :--- | :--- |
| 0x | Smart Contract |  | Ethereum | 2 min. |
| Balancer | Smart Contract | - | Ethereum  | 2 min. |
| Bancor | Rest API | [Bancor API Documentation](https://support.bancor.network/hc/en-us/articles/360002246912-Price-Discovery-value-API) | Ethereum | 2 min. |
| CurveFi | Smart Contract |  | Ethereum | 2 min. |
| DForce | Smart Contract |  | Ethereum | 2 min. |
| Gnosis | Smart Contract |  | Ethereum | 2 min. |
| KuCoin | Websocket API | [Kucoin API Documentation](https://docs.kucoin.com/#websocket-feed) | Ethereum | 2 min. |
| Kyber | Smart Contract |  | Ethereum  | 2 min. |
| Loopring | Websocket API | [Loopring Websocket Documentation](https://docs.loopring.io/en/websocket/overview.html) | Ethereum | 2 min. |
| SushiSwap | Smart Contract |  | Ethereum | 2 min. |
| Uniswap V2 | Smart Contract |  | Ethereum | 2 min. |
| \_\_ |  | _Coming soon:_ |  |  |
| Maker | Rest API | [Maker API Documentation](https://developer.makerdao.com/oasis/api/2/) | Ethereum | 1 min. |

## Lending/Borrowing Data

We retrieve and store lending and borrowing rates such as locked volumes from as close to the data source as possible and with the highest precision possible. All lending/borrowing protocols in our database are decentralized and hence accessible directly through the respective blockchain.

{% hint style="info" %}
 In order to supply data to our community as quickly as possible, we retrieved lending and borrowing data through different types of APIs for now \(see table below\). However, we remark that we are currently implementing data retrieval through the blockchain directly in order to reduce unnecessary dependencies. This concerns _all_ decentralized lending/borrowing protocols.
{% endhint %}

| Protocol | Data Retrieval | API Link | Blockchain | Retrieval Period |
| :--- | :--- | :--- | :--- | :--- |
| AAVE | Thegraph API | [AAVE API Documentation](https://github.com/aave/aave-protocol/tree/master/thegraph) | Ethereum | 1 min. |
| bZx | Blockchain | [bZx Base Protocol](https://github.com/bZxNetwork/bZx-monorepo/tree/development/packages/contracts/contracts) | Ethereum | 1 min. |
| Compound | Blockchain | [Compound Base Protocol](https://github.com/diadata-org/diadata/tree/master/internal/pkg/defiscrapers/compound) | Ethereum | 1 min. |
| DDEX | Rest API | [DDEX API Documentation](https://margin-docs.ddex.io/#get-lending-pool-stats) | Ethereum | 1 min. |
| DY/DX | Rest API | [DY/DX API Documentation](https://docs.dydx.exchange/#solo-get-v1-markets) | Ethereum | 1 min. |
| NUO | Rest API | - | Ethereum | 1 min. |

