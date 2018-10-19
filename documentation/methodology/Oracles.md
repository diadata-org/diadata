# Oracles in DIA

Oracles are Smart Contract interfaces that bring data from an external source into a smart contract.
By generating verifiable oracle data, any user can use financial data in smart contracts.
The first oracle in DIA is the oracle for the biggest cryptocurrencies from Coinhub.

## Coinhub Oracles
The Coinhub oracles are located in a [single smart contract](https://ropsten.etherscan.io/address/0x9f365d522fa57778ce1a70270b240713d270bfeb).
This smart contract holds mappings from an index to an asset's name, its price, and supply data.
By using `getCoinInfo(string name)` it is possible to retrieve this data.
Along with the actual data there is a metadata field for the timestamp of the last update.

Updates are supplied by the DIA Oracle service that periodically supplies updates into the smart contract.
In the event view, the latest updates can be seen.
