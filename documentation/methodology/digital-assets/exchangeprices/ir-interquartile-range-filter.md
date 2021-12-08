---
description: >-
  This section describes how we filter for outliers in ourprice determination
  system.
---

# IR: Interquartile Range Filter

## Outliers and Market Manipulation

The price estimation process needs to be robust against trades with prices that are diverting from the current market price of an asset pair. Reasons for diverting trades can be market manipulation, errors, flash crashes on certain exchanges, or other market irregularities.

Detecting and excluding outliers and market manipulation is an important data processing task, especially in small and (somewhat) intransparent markets. Otherwise, a single low-volume trade can offset the price estimation and serve as a base token for other assets, leading to a chain reaction of misaligned price data.

We calculate our cryptocurrency spot-prices using 120-second moving average (SMA) with a weighting of trade volumes (VWAP). Small volume trades have a proportionally smaller influence on the determined price than bigger volume trades. In order to create reliable pricing, we also exclude data points and sets entirely that lie outside of an [acceptable range](https://en.wikipedia.org/wiki/Interquartile\_range#Outliers) relative to the [interquartile range](https://en.wikipedia.org/wiki/Interquartile\_range) of all data points.

### Data Preprocessing

The Interquartile Range filter in DIA examined all trades in a trades block (i.e. all trades in the predefined time range) and sorts them by their recorded price.

After that, this range of prices is divided into four price blocks, the _Quartiles_. The boundaries of the full price range also determine the boundaries of the first and the last quartile.

To clear out outliers, any trades falling into the first or the last quartile are filtered out and subsequently, only trades falling into the "middle" quartiles are returned to the caller.

### Advanced Filtering

After this clearing process is finished, the data is ready to be processed further in other filters like the [Moving Average](mair-moving-average-with-interquartile-range-filter.md) filter or the [Median](medir-median-with-interquartile-range-filter.md) filter.

