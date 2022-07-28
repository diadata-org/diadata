# MEDIR: Median with Interquartile Range Filter

The Median filter with [Interquartile Range Clearing](ir-interquartile-range-filter.md) is another popular filters for price determination.

### Trade Collection

All trades from the queried time range are ordered by timestamp.

For each second in the time range, there exists a "slot" where trades are put into.&#x20;

### Price Calculation

As soon as all going through all trades in the block has been finalized, the median of these prices is calculated. This is done by accumulating prices and ordering them by price. After that, the trade in the middle of the ordered set is selected. Its price represents the median of the set.

The result is then returned as the result of the filter operation.

### Filter Application

The MEDIR filter is used in DIA's price determination. Our API can display the latest MEDIR-120 filter values, i.e., the filter results from a 120 second interval of all recorded trades for an asset.

### Implementation

The filter is implemented as part of the FiltersBlockService in [this file in our Github repository](../../../../internal/pkg/filtersBlockService/FilterMEDIR.go).
