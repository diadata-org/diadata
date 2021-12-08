---
description: >-
  This section describes the use of the Moving AVerage methodology in the DIA
  filter system.
---

# MA: Moving average

The Moving Average filter is one of the most commonly used filters for price determination.

### Trade Collection

All trades from the queried time range are ordered by timestamp.

For each second in the time range, there exists a "slot" where trades are put into.&#x20;

### Price Calculation

As soon as all going through all trades in the block has been finalized, the simple average of these prices is calculated. This is done by accumulating prices and then dividing by the number of involved trades in this calculation.

The result is then returned as the result of the filter operation.
