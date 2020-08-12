---
description: How do I access DEX quotations using the DIA oracle system?
---

# DEX Quotations

The oracle contains information about DEX price quotations. You can access the last recorded value as well as the timestamp of the last update.

1.  Access our [oracle smart contract](https://etherscan.io/address/0xD47FDf51D61c100C447E2D4747c7126F19fa23Ef).
2. Call `getCoinInfo(dex_name)` with `dex_name` being the full coin name such as `Bancor`. You can use the "Read" section on Etherscan to execute this call.
3. The response of the call contains four values:
   1. The current USDT price in USD with a fix-comma notation of five decimals.
   2. Not used in this call.
   3. The [UNIX timestamp](https://www.unixtimestamp.com/) of the last oracle update.
   4. The short name of the asset, e.g., `USDT` for USD Tether.

