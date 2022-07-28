---
description: This page contains information about the EMA pricing methodology.
---

# EMA: Exponential Moving Average

The EMA (Exponential Moving Average) filter is a filter that can be applied to a time series of existing price points. These price points originate from one of the other filter methods (e.g., [VWAPIR](vwapir-volume-weighted-average-price-with-interquartile-range-filter.md), [MAIR](mair-moving-average-with-interquartile-range-filter.md), or [MEDIR](medir-median-with-interquartile-range-filter.md)). The EMA filter then produces a moving average over a number of the latest of these price points.

### Filter Application

The EMA filter is used as a post-processing filter in our graphql frontend. It uses the underlying MA120 filter points and produces a time series using a moving window approach.

For each EMA filter point, the algorithm takes into account a configurable amount of MA120 filter points from the past and calculated a weighted average. This weighting depends on the age of the MA120 filter points and decreases exponentially towards the past. By that the EMA filter ensures that recent data points have a higher weight in the filter end result compared to older ones.

A detailed writeup of EMA functionality can be [found here](https://www.investopedia.com/terms/e/ema.asp).

### Implementation

The filter is implemented as part of the FiltersBlockService in [this file in our Github repository](../../../../internal/pkg/filtersBlockService/FilterEMA.go).
