# Oracles in DIA

Oracles are Smart Contract interfaces that bring data from an external source into a smart contract.
By generating verifiable oracle data, any user can use financial data in smart contracts.
The first oracle in DIA is the oracle for the biggest cryptocurrencies from Coinhub.

## Coinhub Oracles
The Coinhub oracles are located in a [single smart contract](https://etherscan.io/address/0xD47FDf51D61c100C447E2D4747c7126F19fa23Ef).
This smart contract holds mappings from an index to an asset's name, its price, and supply data.
By using `getCoinInfo(string name)` it is possible to retrieve this data.
Along with the actual data there is a metadata field for the timestamp of the last update.

Updates are supplied by the DIA Oracle service that periodically supplies updates into the smart contract.
In the event view, the latest updates can be seen.
