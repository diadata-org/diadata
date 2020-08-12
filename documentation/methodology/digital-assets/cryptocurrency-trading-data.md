---
description: Description of data retrieval and processing of raw trading data
---

# Cryptocurrency Trading Data

## Centralized Cryptocurrency  Exchanges

The raw trading data for all our Centralized Cryptocurrency Exchanges is the basis of all further derived quantities such as prices and circulating supply numbers. Trading data is retrieved through Rest APIs or websocket APIs  which are developed and maintained by the respective trading platform. See the below table for detailed information.   
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

## Decentralized Cryptocurrency Exchanges \(DEXes\)

In contrast to centralized exchanges, 

