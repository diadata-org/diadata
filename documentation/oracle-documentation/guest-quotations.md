---
description: How do I access Coingecko quotations using the DIA oracle system?
---

# Coingecko Quotations

The oracle contains information about price quotations of crypto assets, accumulated by Coingecko. You can access quotations for the top 100 assets ranked by market capitalization. The response consists of the most recent quotation as well as the timestamp of the last update. You can use the "Read" section on Etherscan to execute this call as follows:

1.  Access our [Coingecko oracle smart contract](https://etherscan.io/address/0x07e4120dd7411a49e091a20fa0be33a183c35d60).
2. Call `getValue(key)` where `key` is the short name for the asset in capital letters, for instance `BTC` for Bitcoin. You can use the "Read" section on Etherscan to execute this call.
3. The response of the call contains two values:
   1. The current asset price in USD with a fix-comma notation of five decimals.
   2. The [UNIX timestamp](https://www.unixtimestamp.com/) of the last oracle update.

