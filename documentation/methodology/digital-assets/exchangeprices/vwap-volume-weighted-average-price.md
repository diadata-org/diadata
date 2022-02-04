---
description: This page contains information about the VWAP pricing methodology
---

# VWAP: Volume Weighted Average Price

VWAP (Volume Weighted Average Price) is a methodology for trade-based price determination that takes into account the different volumes of trades.

### Trade Collection

All trades from the queried time range are collected and weighted by their volume. Weighting means that the (normalized) volume of the trade is multiplied by its executed price.

### Price Calculation

As soon as the block has been finalized, the VWAP price of these prices is calculated. This is done by accumulating the previously calculated volume-price-products of all trades and dividing them by the sum of all volumes combined.

The result is then returned as the result of the filter operation.

### Filter Application

The VWAP filter is used in DIA's price determination. Our API can display the latest VWAP filter values, i.e., the filter results from a 120 second interval of all recorded trades for an asset.

### Implementation

The filter is implemented as part of the FiltersBlockService in [this file in our Github repository](https://github.com/diadata-org/diadata/blob/master-ibm-tmp/internal/pkg/filtersBlockService/FilterVWAP.go).

