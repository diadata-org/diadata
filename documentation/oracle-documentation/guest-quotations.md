---
description: How do I access Coingecko quotations using the DIA oracle system?
---

# Coingecko Quotations

The oracle contains information about price quotations of crypto assets Coingecko. You can access quotations for the top 100 assets ranked by market capitalization. A list of these assets can be retrieved from the [coingecko API](https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false).  
You can execute an oracle call as follows:

1.  Access DIA's [Coingecko oracle smart contract](https://etherscan.io/address/0x07e4120dd7411a49e091a20fa0be33a183c35d60).
2. Call `getValue(key)` where `key` is the symbol for the asset in capital letters, for instance `BTC` for Bitcoin. You can use the "Read" section on Etherscan to execute this call.
3. The response of the call contains two values:
   1. The current asset price in USD with a fix-comma notation of five decimals.
   2. The [UNIX timestamp](https://www.unixtimestamp.com/) of the last oracle update.

