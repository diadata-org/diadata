---
description: How do I access CoinMarketCap quotations using the DIA oracle system?
---

# CoinMarketCap Quotations

The oracle contains information about price quotations of crypto assets from [CoinMarketCap](https://coinmarketcap.com/). You can access quotations for the top 50 assets ranked by market capitalization.  
You can execute an oracle call as follows:

1.  Access DIA's [CoinMarketCap Oracle](../deployed-contracts.md).
2. Call `getValue(key)` where `key` is the symbol for the asset in capital letters, for instance `BTC` for Bitcoin. You can use the "Read" section on Etherscan to execute this call.
3. The response of the call contains two values:
   1. The current asset price in USD with a fix-comma notation of five decimals.
   2. The [UNIX timestamp](https://www.unixtimestamp.com/) of the last oracle update.

