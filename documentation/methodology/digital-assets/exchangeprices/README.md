---
description: >-
  This section covers the methodology of how DIA gets and calculates prices for
  cryptocurrency assets.
---

# Pricing Methodology of Crypto Exchange Quotations

Cryptocurrency assets are mainly traded on centralised and decentralised exchange markets (CEX and DEX). Centralized markets generally require traders to open accounts, transfer their assets to a centralised wallet belonging to the exchange and then trading on the off-chain trading engine belonging to the exchange. Only when an asset is transferred from the exchange to another wallet, a change of ownership can be recorded on-chain.

In contrast, DEXes operate entirely on-chain and assets are exchanged directly from wallet to wallet. While this gets rid of the requirement for accounts and trust in the exchange itself, it also restratins the tradable assets to ones that are accessible to the on-chain component representing the DEX (usually a smart contract). Also, on-chain transaction costs can be higher than in an off-chain system due to gas usage.

Thus, conversion between assets from incompatible blockchains is restricted to centralized exchanges, as well as conversion between crypto and fiat assets.

## Price Determination

The vast majority of recorded trades by DIA are between two crypto assets. Only a minority of trades against fiat currencies is recorded on CEXes. Nontheless, DIA reports all asset prices in USD, even for numerous assets where only crypto-crypto pairs exist.

To retrieve a USD price for each and every asset, DIA uses a _price estimator._ The price estimator is updated with every recorded trade. Each trade consists of a _base token_ and a _quote token_. The price of a quote token is measured in base tokens. Because the order in a trade of base token and quote token is arbitrary, USD is always considered as base token. Other fiat currencies are also always base tokens, except when measuring the [ECB exchange rates](../../traditional-assets/ecb-foriegn-exchange-data.md). An asset needs to be a quote token in at least one market for us to be able to determine a USD price.

After a trade is recorded, it is fed into the price estimator. The estimator determines the current value of the base token measured in USD (by a chain of trades ending at a fiat pair). From there, the price estimator determines the price of the quote token and stores it, should the quote token become a base token in a later trade.

## Filters

After trades have been recorded, they are collected into blocks of 120 seconds. For each block, a range of filters are calculated. A filter is a function to get a single price point from the collection of trades in a block. For example, the [Moving Average](ma-moving-average.md) filter returns the volume weighted moving average of the trades block.

Depending on the use case, other filters can be useful as well. DIA includes filters for various use cases.

Filters are calculated for each asset on each exchange individually, as well as for each asset on all exchanges combined. This combined filter result represents the result closest to the true "whole market" that can be determined by this system.

| Filter Name                                                    | Community Approval   |
| -------------------------------------------------------------- | -------------------- |
| [MA](ma-moving-average.md)                                     | Approval in Progress |
| [MAIR](mair-moving-average-with-interquartile-range-filter.md) | Approval Outstanding |
| [MEDIR](medir-median-with-interquartile-range-filter.md)       | Approval Outstanding |
| [VWAP](vwap-volume-weighted-average-price.md)                  | Approval Outstanding |

## Outliers and Market Manipulation

Outliers are cleared using the [Interquartile Range Filter](ir-interquartile-range-filter.md) methodology.
