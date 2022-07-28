# MAIR: Moving Average with Interquartile Range Filter

The MAIR (Moving Average with Interquartile Range Clearing) filter is the application of the [Moving Average](ma-moving-average.md) filter to data that was previously cleared by the [Interquartile Range](ir-interquartile-range-filter.md) filtering system.

### Filter Application

The MAIR filter is used in DIA's price determination. Our price quotations are the latest MAIR-120 filter values, i.e., the filter results from a 120 second interval of all recorded trades for an asset.
