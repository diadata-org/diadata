---
title: DIA Data
description: How to use request data from a DIA Oracle in your Astar Dapp using smart contracts
---

# DIA Data Oracles on Astar

## Introduction

DIA is an ecosystem for open financial data in a financial smart contract ecosystem.
The target of DIA is to bring together data analysts, data providers and data users.
In general, DIA provides a reliable and verifiable bridge between off-chain data from various sources and on-chain smart contracts that can be used to build a variety of financial dApps. 
DApp developers who want to leverage DIA oracles can access the published data on Astar.
DIA offers data about traditional financial assets and cryptocurrencies.
[Read our documentation](https://docs.diadata.org) to learn about our methodologies, API, oracles, and how to contribute.

## Supported Assets

DIA supports assets from various categories to be included into the oracle. A selection of assets is listed here:

Data Feed Name  |   Data Feed Query           | Data Type                           |
| :-----------: | :-------------------------: | :---------------------------------: |
| Bitcoin       |   BTC/USD                   |  Crypto Price                       |
| Ethereum      |   ETH/USD                   |  Crypto Price                       |
| USDC          |   USDC/USD                  |  Crypto Price                       |
| DIA Token     |   DIA/USD                   |  Crypto Price                       |
| ASTR Token    |   ASTR/USD                  |  Crypto Price                       |

To retrieve data, query the oracle for an asset as listed in column "Data Field Query".
The query string is case-sensitive.

## Data Access

All asset prices are determined in USD according to our [methodology](https://docs.diadata.org/documentation/methodology).
They are denominated in a fix comma format with 5 decimals, so you need to divide them by 100000 to retrieve the actual value in USD.
Where applicable, the oracle also provides information on circulating supply and the timestamp of data collection.
The query in the smart contract is realized with the symbol of the asset.

### Smart Contract

DIA data is published in the `DIAOracle` smart contract.
By querying the `getValue()` function you can retrieve the requested data.
You can interact with our contract in the [blockchain explorer](https://blockscout.com/astar/address/0xd79357ebb0cd724e391f2b49a8De0E31688fEc75/transactions), where you can query for any supported asset.

The contract takes the symbol name of the asset as input, e.g., `BTC` and returns a tuple of the timestamp of the last update and the actual value.

### Write your own DApp

To access oracle data, you can either use the explorer above or write your own contract and reference the oracle.
The following snippet shows how to retrieve the price of an asset (e.g. `DIA`) measured in another asset (BTC) using a wrapper smart contract.
On Astar, you can initialize the oracle address in your wrapper contract by calling `setOracleAddress()` with the address of our deployed demo contract `0xd79357ebb0cd724e391f2b49a8De0E31688fEc75`.

```
pragma solidity 0.7.4;

contract DIAOracle {
    mapping (string => uint256) public values;
    address oracleUpdater;
    
    event OracleUpdate(string key, uint128 value, uint128 timestamp);
    event UpdaterAddressChange(address newUpdater);
    
    constructor() {
        oracleUpdater = msg.sender;
    }
    
    function setValue(string memory key, uint128 value, uint128 timestamp) public {
        require(msg.sender == oracleUpdater);
        uint256 cValue = (((uint256)(value)) << 128) + timestamp;
        values[key] = cValue;
        emit OracleUpdate(key, value, timestamp);
    }
    
    function getValue(string memory key) public view returns (uint128, uint128) {
        uint256 cValue = values[key];
        uint128 timestamp = (uint128)(cValue % 2**128);
        uint128 value = (uint128)(cValue >> 128);
        return (value, timestamp);
    }
    
    function updateOracleUpdaterAddress(address newOracleUpdaterAddress) public {
        require(msg.sender == oracleUpdater);
        oracleUpdater = newOracleUpdaterAddress;
        emit UpdaterAddressChange(newOracleUpdaterAddress);
    }
}

contract DiaAssetBtcOracle {
	DIAOracle oracle;
	address owner;
    
	constructor() public {
		owner = msg.sender;
	}
    
	function setOracleAddress(address _address) public {
		require(msg.sender == owner);
		oracle = DIAOracle(_address);
	}
    
	function getAssetBtcRate(string asset) constant public returns (uint256) {
		(uint assetPrice,) = oracle.getValue(asset);
		(uint btcPrice,) = oracle.getValue("BTC/USD");
		return (assetPrice / btcPrice);
	}
    
}
```

#### Deployed Addresses

| Chain name    |        Oracle Contract Address                | Update Frequency |
| :------------ | :-------------------------------------------: | :--------------: |
| Astar         | `0xd79357ebb0cd724e391f2b49a8De0E31688fEc75i` |      120 seconds |

### Github and Contact

DIA provides a broad range of assets. You can find an overview in the DIA documentation [here](https://docs.diadata.org/documentation/oracle-documentation).
All our code is open-source and can be found on our [Github repositoy](https://github.com/diadata-org/diadata).
For the deployment of specific oracles (source/methodology/frequency) please [contact the DIA team](mailto:bd@diadata.org).

You can follow us on [Telegram](https://t.me/DIAdata_org), [Github](https://github.com/diadata-org), and [Medium](https://medium.com/dia-insights).
