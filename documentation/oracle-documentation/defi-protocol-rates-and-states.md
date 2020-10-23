---
description: How do I access DeFi protocol information using the DIA oracle system?
---

# DeFi Protocol Lending and Borrowing Rates

The oracle contains information about DeFi protocols. You can access a quotation of the current lending and borrowing rates and the  as well as the timestamp of the last update.

1.  Access our [oracle smart contract](https://etherscan.io/address/0xD47FDf51D61c100C447E2D4747c7126F19fa23Ef).
2. Call `getCoinInfo(protocol_name)` with `protocol_name` being the full protocol name in capital letters, such as `COMPOUND`. You can use the "Read" section on Etherscan to execute this call.
3. The response of the call contains four values:
   1. The lending rate with a fix-comma notation of five decimals.
   2. The borrowing rate with a fix-comma notation of five decimals.
   3. The [UNIX timestamp](https://www.unixtimestamp.com/) of the last oracle update.
   4. The short name of the asset under consideration, e.g., `DAI` for Compound Protocol.

