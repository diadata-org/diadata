---
description: How do I access guest quotations using the DIA oracle system?
---

# Guest Quotations

The oracle contains information about price quotations of crypto assets, accumulated by other entities such as Coingecko. You can access a price quotation, the asset's symbol as well as the timestamp of the last update.

1.  Access our [oracle smart contract](https://etherscan.io/address/0xD47FDf51D61c100C447E2D4747c7126F19fa23Ef).
2. Call `getCoinInfo(source-coin_name)` where `source` is the source of the quotation such as Coingecko and `coin_name` is the full coin name such as Bitcoin, including blank spaces if present. You can use the "Read" section on Etherscan to execute this call.
3. The response of the call contains three values:

   1. The current asset price in USD with a fix-comma notation of five decimals.
   2. The short name of the asset, e.g., `BTC` for Bitcoin.
   3. The [UNIX timestamp](https://www.unixtimestamp.com/) of the last oracle update.

Currently available are the following 20 largest quotations by market cap from [Coingecko](https://www.coingecko.com/en):

Bitcoin, Ethereum, Tether, XRP, Binance Coin, Bitcoin Cash, Polkadot, Chainlink, Bitcoin SV, Cardano, Crypto.com Coin, Litecoin, USD Coin, EOS, Monero, TRON, OKB, Tezos, Stellar, NEO

