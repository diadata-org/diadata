# Pricing Methodology of Crypto Exchange Quotations

## Outliers and Market Manipulation

Detecting and excluding outliers and market manipulation is an important data processing task, especially in small and \(somewhat\) intransparent markets.

We calculate our cryptocurrency-spot-prices using 120-second moving average \(SMA\). In order to create reliable pricing, we also exclude data points & sets that lie outside of an [acceptable range](https://en.wikipedia.org/wiki/Interquartile_range#Outliers) relative to the [interquartile range](https://en.wikipedia.org/wiki/Interquartile_range) of all data points.

This will be constantly improved over time with crowd-driven algorithmic improvements from the traditional financial space, taking into consideration that high volatility is a part of the cryptocurrency space.

## API Location

Our crypto exchange quotations can be retrieved from our API using [https://api.diadata.org/v1/quotation/\[TLA\]](https://api.diadata.org/v1/quotation/[TLA]), with TLA being the short name of a currency. As an example, the current Bitcoin price is stored at [https://api.diadata.org/v1/quotation/BTC](https://api.diadata.org/v1/quotation/BTC).

