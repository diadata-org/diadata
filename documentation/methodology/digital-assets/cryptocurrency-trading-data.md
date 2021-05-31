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
| KuCoin | Websocket API | [Kucoin API Documentation](https://docs.kucoin.com/#websocket-feed) | - |
| LBank | Websocket API | [LBank Websocket Documentation](https://docs.lbkex.net/en/#websocket-api-market-data) | - |
| OKEx | Websocket API | [OKEx Websocket Documentation](https://www.okex.com/docs/en/#ws_swap-README) | - |
| Quoine | Rest API | [Quoine Rest API Documentation](https://developers.liquid.com/) | 1.5 sec.  |
| ZB | Websocket API | [ZB Websocket Documentation](https://www.zb.com/api#WebSocket%20API) | - |

### Decentralized Cryptocurrency Exchanges \(DEXes\)

In contrast to centralized exchanges, in decentralized exchanges it is possible to retrieve trading data directly from the respective blockchain.

{% hint style="info" %}
 In order to supply data to our community as quickly as possible, we retrieved the trading data through different types of APIs for now \(see table below\). However, we remark that we are currently implementing data retrieval through the blockchain directly in order to reduce unnecessary dependencies. This concerns _all_ decentralized exchanges.
{% endhint %}

| Exchange | Data Retrieval | API Link | Blockchain | Retrieval Period |
| :--- | :--- | :--- | :--- | :--- |
| 0x | [Smart Contract](https://etherscan.io/address/0x61935CbDd02287B511119DDb11Aeb42F1593b7Ef) | [Github Repository](https://github.com/0xProject) | Ethereum | - |
| Balancer | [Smart Contract](https://etherscan.io/address/0x9424B1412450D0f8Fc2255FAf6046b98213B76Bd) | [Github Repository](https://github.com/balancer-labs) | Ethereum  | - |
| Bancor | Rest API | [Bancor API Documentation](https://support.bancor.network/hc/en-us/articles/360002246912-Price-Discovery-value-API) | Ethereum | 2 min. |
| CurveFi | [Smart Contract](https://etherscan.io/address/0x7002B727Ef8F5571Cb5F9D70D13DBEEb4dFAe9d1) | [Github Repository](https://github.com/curvefi/curve-contract) | Ethereum | - |
| DForce | [Smart Contract](https://etherscan.io/address/0x03eF3f37856bD08eb47E2dE7ABc4Ddd2c19B60F2) | [Github Repository](https://github.com/dforce-network) | Ethereum | - |
| Gnosis | [Smart Contract](https://etherscan.io/address/0x6F400810b62df8E13fded51bE75fF5393eaa841F) | [Github Repository](https://github.com/gnosis/dex-contracts) | Ethereum | - |
| Kyber | [Smart Contract](https://etherscan.io/address/0x9AAb3f75489902f3a48495025729a0AF77d4b11e) | [Github Repository](https://github.com/KyberNetwork/smart-contracts) | Ethereum  | - |
| Loopring | Websocket API | [Loopring Websocket Documentation](https://docs.loopring.io/en/websocket/overview.html) | Ethereum | 2 min. |
| SushiSwap | [Smart Contract](https://etherscan.io/address/0xc0aee478e3658e2610c5f7a4a2e1777ce9e4f2ac) | [Github Repository](https://github.com/sushiswap/sushiswap/tree/master/contracts) | Ethereum | - |
| Uniswap V2 | [Smart Contract](https://etherscan.io/address/0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f) | [Github Repository](https://github.com/Uniswap/uniswap-v2-core) | Ethereum | - |
| Uniswap V3 | [Smart Contract](https://etherscan.io/address/0x1F98431c8aD98523631AE4a59f267346ea31F984#code) | [Github Repository](https://github.com/Uniswap/uniswap-v3-core) | Ethereum |  |
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
| bZx | Smart Contract | [bZx Base Protocol](https://github.com/bZxNetwork/bZx-monorepo/tree/development/packages/contracts/contracts) | Ethereum | 1 min. |
| Compound | Smart Contract | [Compound Base Protocol](https://github.com/diadata-org/diadata/tree/master/internal/pkg/defiscrapers/compound) | Ethereum | 1 min. |
| CREAM | Smart Contract | [Compound Base Protocol](https://github.com/diadata-org/diadata/tree/master/internal/pkg/defiscrapers/compound) | Ethereum | 1 min. |
| DDEX | Rest API | [DDEX API Documentation](https://margin-docs.ddex.io/#get-lending-pool-stats) | Ethereum | 1 min. |
| DY/DX | Rest API | [DY/DX API Documentation](https://docs.dydx.exchange/#solo-get-v1-markets) | Ethereum | 1 min. |
| Fortube | Smart Contract |  | Ethereum | 1 min. |
| NUO | Rest API | - | Ethereum | 1 min. |

