---
description: >-
  How do I access chart points from several exchanges using the DIA oracle
  system?
---

# Chart Points

The oracle contains information about chart points of an ever growing list of exchanges and assets. You can access a price quotation and the timestamp of the last update.

1.  Access our [oracle smart contract](https://etherscan.io/address/0xD47FDf51D61c100C447E2D4747c7126F19fa23Ef).
2. Call `getCoinInfo(exchange)` with `exchange` being the full name of the exchange such as `Uniswap`. You can use the "Read" section on Etherscan to execute this call.
3. The response of the call contains three values:
   1. The current asset price in USD with a fix-comma notation of five decimals.
   2. The [UNIX timestamp](https://www.unixtimestamp.com/) of the last oracle update.
   3. The short name of the quoted asset, e.g., `ETH` for Ethereum.

As of now, the following list of exchanges is available in this oracle:  
- Curvefi \(DAI\)  
- Uniswap \(ETH\)  
- Bancor \(ETH\)  
- 0x \(ETH\)  
- Gnosis \(ETH\)  
- Loopring\(ETH\)  


