---
description: How do I access DeFi protocol information using the DIA oracle system?
---

# DeFi Protocol Lending and Borrowing Rates

The oracle exists on multiple blockchains and contains information about DeFi protocols. You can access a quotation of the current lending and borrowing rates and the as well as the timestamp of the last update.

1.  Access the corresponding [oracle smart contract](https://docs.diadata.org/documentation/oracle-documentation/deployed-contracts).
2. Call `getCoinInfo(protocol_name)` with `protocol_name` being the full protocol name in capital letters, such as `COMPOUND`. You can use the "Read" section on Etherscan to execute this call.
3. The response of the call contains four values:
   1. The lending rate with a fix-comma notation of five decimals.
   2. The borrowing rate with a fix-comma notation of five decimals.
   3. The [UNIX timestamp](https://www.unixtimestamp.com/) of the last oracle update.
   4. The short name of the asset/pool under consideration, e.g., `DAI` for Compound Protocol.

As of now, the following list of rates is available in this oracle:

**Ethereum Mainnet:**

* MAKERDAO \(ETH-A\)

**Binance Smart Chain:**

* AAVE \(DAI\)
* BITFINEX \(DAI\)
* BZX \(DAI\)
* COMPOUND \(DAI\)
* CREAM \(UNI\)
* DYDX \(DAI\) 
* FORTUBE \(DAI\)
* MAKERDAO \(DAI\)
* NUO \(DAI\)

**Polygon \(Matic\):**

* AAVE \(DAI\)
* BITFINEX \(DAI\)
* BZX \(DAI\)
* COMPOUND \(DAI\)
* CREAM \(UNI\)
* DYDX \(DAI\) 
* FORTUBE \(DAI\)
* MAKERDAO \(DAI\)
* NUO \(DAI\)

**Moonbeam:**

* AAVE \(DAI\)
* BITFINEX \(DAI\)
* BZX \(DAI\)
* COMPOUND \(DAI\)
* CREAM \(UNI\)
* DYDX \(DAI\) 
* FORTUBE \(DAI\)
* MAKERDAO \(DAI\)
* NUO \(DAI\)

_Remark_: The borrowing rate for DAI on MAKERDAO refers to the collateral ETH.

