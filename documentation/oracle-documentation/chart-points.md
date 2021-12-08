---
description: >-
  How do I access chart points from several exchanges using the DIA oracle
  system?
---

# Chart Points

The oracle exists on various blockchains and contains information about chart points. You can access a price quotation and the timestamp of the last update.

1. &#x20;Access the corresponding [oracle smart contract](https://docs.diadata.org/documentation/oracle-documentation/deployed-contracts).
2. Call `getCoinInfo(exchange)` with `exchange` being the full name of the exchange such as `Uniswap`. You can use the "Read" section on Etherscan to execute this call.
3. The response of the call contains three values:
   1. The current asset price in USD with a fix-comma notation of five decimals.
   2. The [UNIX timestamp](https://www.unixtimestamp.com) of the last oracle update.
   3. The short name of the quoted asset, e.g., `ETH` for Ethereum.

As of now, the following list of price points is available in this oracle:

**Ethereum Mainnet:**

* PanCakeSwap (WBNB)

**Binance Smart Chain:**

* 0x (ETH)
* Bancor (ETH)
* Bitmax (ETH)
* CREX24 (CREX)
* Curvefi (DAI)
* Gnosis (ETH)
* Kyber (ETH)
* STEX (PLEX)
* SushiSwap (ETH)
* Uniswap (ETH)
* PanCakeSwap (WBNB)

**Polygon (Matic):**

* 0x (ETH)
* Bancor (ETH)
* Bitmax (ETH)
* CREX24 (CREX)
* Curvefi (DAI)
* Gnosis (ETH)
* Kyber (ETH)
* STEX (PLEX)
* SushiSwap (ETH)
* Uniswap (ETH)
* PanCakeSwap (WBNB)&#x20;

**Moonbeam:**

* 0x (ETH)
* Bancor (ETH)
* Bitmax (ETH)
* CREX24 (CREX)
* Curvefi (DAI)
* Gnosis (ETH)
* Kyber (ETH)
* STEX (PLEX)
* SushiSwap (ETH)
* Uniswap (ETH)
* PanCakeSwap (WBNB)&#x20;
