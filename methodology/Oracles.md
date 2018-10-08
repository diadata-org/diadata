# Oracles in DIA

Oracles are Smart COntract interfaces that birng data from an external source into a smart contract.
By generating verifiable oracle data, any user can use financial data in smart contracts.
The first oracle in DIA is the oracle for the biggest cryptocurrencies from Coinhub.

## Coinhub Oracles
The Coinhub oracles are located in a [sinngle smart contract](https://etherscan.io/address/0x284a9d8cf1913e70142efb5b3ebd1de71f9a5f76).
This smart contract holds mappings from an index to an asset's name, its price, and supply data.
Using `getParameters(asset_index)` it is possible to retrieve this data.
Along with the actual data there is a metadata field for the timestamp of the last update.

Updates are supplied by the DIA Oracle service that periodically supplies updates into the smart contract.
In the event view, the latest updates can be seen.
The event fields show the values in the following locations:

* Price in Cent is in the second field (display as `number`)
* Supply is in the third field (display as `number`)
* UNIX Timestamp of last update is displayed in the fourth field (display as `number`)
* Name is displayed in the sixth and last field (display as `text`).
