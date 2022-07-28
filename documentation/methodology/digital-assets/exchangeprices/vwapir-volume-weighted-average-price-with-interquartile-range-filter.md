---
description: This page contains information about the VWAPIR pricing methodology.
---

# VWAPIR: Volume Weighted Average Price with Interquartile Range Filter

The VWAPIR (Volume Weighted Average Price with Interquartile Range Clearing) filter is the application of the [Volume Weighted Average Price](vwap-volume-weighted-average-price.md) filter to data that was previously cleared by the [Interquartile Range](ir-interquartile-range-filter.md) filtering system.

### Filter Application

The VWAPIR filter is used in DIA's price determination. Our API can display the latest MEDIR-120 filter values, i.e., the filter results from a 120 second interval of all recorded trades for an asset.

### Implementation

The filter is implemented as part of the FiltersBlockService in [this file in our Github repository](https://github.com/diadata-org/diadata/blob/master-ibm-tmp/internal/pkg/filtersBlockService/FilterVWAPIR.go).
