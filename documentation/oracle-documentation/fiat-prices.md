---
description: How do I access DEX quotations using the DIA oracle system?
---

# Fiat Prices

The oracle contains information about Fiat prices. You can access a quotation of the current fiat price rate as well as the timestamp of the last update.

1.  Access our [oracle smart contract](https://etherscan.io/address/0xD47FDf51D61c100C447E2D4747c7126F19fa23Ef).
2. Call `getCoinInfo(fiat_name)` with `fiat_name` being the short fiat name such as `EUR`. You can use the "Read" section on Etherscan to execute this call.
3. The response of the call contains four values:
   1. The fiat price in USD with a fix-comma notation of five decimals.
   2. Not used in this asset class.
   3. The [UNIX timestamp](https://www.unixtimestamp.com/) of the last oracle update.
   4. The short name of the protocol, e.g., `EUR` for Euro.

