---
description: How do I access crypto index values using the DIA oracle system?
---

# Crypto Index Value

The oracle contains information about index values of crypto indices. You can access index values for all current indices offered by DIA.\
You can execute an oracle call as follows:

1. Access DIA's [Crypto Index Oracle](deployed-contracts.md)
2. Call `getValue(key)` where `key` is the symbol for the index in capital letters, for instance `SCIFI` for the SCIFI Index. You can use the "Read" section on Etherscan to execute this call.
3. The response of the call contains two values:
   1. The current index value with a fix-comma notation of six decimals.
   2. The [UNIX timestamp](https://www.unixtimestamp.com) of the last oracle update.

| Query String | Description                                         | Response Format                                                                                       |
| ------------ | --------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| SCIFI        | Returns the current value of the SCIFI crypto index | <ul><li>Index value (fix comma format with 6 digits)</li><li>UNIX timestamp of last update </li></ul> |
