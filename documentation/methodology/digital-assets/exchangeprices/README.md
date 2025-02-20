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

The vast majority of recorded trades by DIA are between two crypto assets. Only a minority of trades against fiat currencies is recorded on CEXes. Nonetheless, DIA reports all asset prices in USD, even for numerous assets where only crypto-crypto pairs exist.

To retrieve a USD price for each and every asset, DIA uses a _price estimator._ The price estimator is updated with every recorded trade. Each trade consists of a _base token_ and a _quote token_. The price of a quote token is measured in base tokens. Because the order in a trade of base token and quote token is arbitrary, USD is always considered as base token. Other fiat currencies are also always base tokens, except when measuring the [ECB exchange rates](../../traditional-assets/ecb-foriegn-exchange-data.md). An asset needs to be a quote token in at least one market for us to be able to determine a USD price.

After a trade is recorded, it is fed into the price estimator. The estimator determines the current value of the base token measured in USD (by a chain of trades ending at a fiat pair). From there, the price estimator determines the price of the quote token and stores it, should the quote token become a base token in a later trade.

## Filters

After trades have been recorded, they are collected into blocks of 120 seconds. For each block, a range of filters are calculated. A filter is a function to get a single price point from the collection of trades in a block. For example, the [Moving Average](ma-moving-average.md) filter returns the volume weighted moving average of the trades block.

Depending on the use case, other filters can be useful as well. DIA includes filters for various use cases.

Filters are calculated for each asset on each exchange individually, as well as for each asset on all exchanges combined. This combined filter result represents the result closest to the true "whole market" that can be determined by this system.

| Filter Name                                                                       | Community Approval                                                                                                             |
| --------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| [MA](ma-moving-average.md)                                                        | [Crowd-approved](https://vote.diadata.org/#/proposal/0x8d74d825c46958afb24cc3591c587b6101f3bd0903e595f9bb293a0947a9ab3f)       |
| [MAIR](mair-moving-average-with-interquartile-range-filter.md)                    | [Crowd-approved](https://vote.diadata.org/#/proposal/0xb7f22a1b4fe36ec2973834accc5aa6704ffecb8fc0f08bb6a61e88cce0af907e)       |
| [MEDIR](medir-median-with-interquartile-range-filter.md)                          | [Crowd-approved](https://vote.diadata.org/#/proposal/0xa8f1e6f4173c3358c99d085ccb15053ed4df6bc243f95c0a5ac7b37123b3b439)       |
| [VWAP](vwap-volume-weighted-average-price.md)                                     | [Crowd-approved](https://vote.diadata.org/#/proposal/0x69be5d17d80c87480aff9be9effe3617cc4dcac0ef593ba6baa4651b45228f50)       |
| [VWAPIR](vwapir-volume-weighted-average-price-with-interquartile-range-filter.md) | [Crowd-approved](https://vote.diadata.org/#/proposal/0x4df8660f951780cd128126ecc3cbd1c693dbece7efb5c2143ee700666f0d75be)       |
| [EMA](ema-exponential-moving-average.md)                                          | [Approval Outstanding](https://vote.diadata.org/#/proposal/0xa67dc7135ce32ab0e3b9c2aeb6ba2ff495f37e99969e58934f3b43a2f6461406) |

## Outliers and Market Manipulation

Outliers are cleared using the [Interquartile Range Filter](ir-interquartile-range-filter.md) methodology.
