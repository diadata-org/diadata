---
description: How do I access crypto asset information using the DIA oracle system?
---

# Crypto Assets

Note: _This oracle is paused until further notice_

The oracle contains information about crypto assets. You can access a price quotation and the current circulating supply as well as the timestamp of the last update.

1.  Access our [oracle smart contract](https://etherscan.io/address/0xD47FDf51D61c100C447E2D4747c7126F19fa23Ef).
2. Call `getCoinInfo(coin_name)` with `coin_name` being the full coin name such as `Bitcoin`. You can use the "Read" section on Etherscan to execute this call.
3. The response of the call contains four values:
   1. The current asset price in USD with a fix-comma notation of five decimals.
   2. The current circulating supply.
   3. The [UNIX timestamp](https://www.unixtimestamp.com/) of the last oracle update.
   4. The short name of the asset, e.g., `BTC` for Bitcoin.



