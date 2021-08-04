---
title: DIA Data
description: How to use request data from a DIA Oracle in your Celo Dapp using smart contracts
---

# DIA Data Oracles on Celo

## Introduction

DIA is an ecosystem for open financial data in a financial smart contract ecosystem.
The target of DIA is to bring together data analysts, data providers and data users.
In general, DIA provides a reliable and verifiable bridge between off-chain data from various sources and on-chain smart contracts that can be used to build a variety of financial dApps. 
DApp developers who want to leverage DIA oracles can access the published data on Celo.
DIA offers data about traditional financial assets and cryptocurrencies.
[Read our documentation](https://docs.diadata.org) to learn about our methodologies, API, oracles, and how to contribute.

## Supported Assets

DIA supports assets from various categories to be included into the oracle. A selection of assets is listed here:

Data Feed Name  |   Data Feed Query           | Data Type                           |
| :-----------: | :-------------------------: | :---------------------------------: |
| DOT Token     |   DOT                       |  Crypto Price                       |
| Bitcoin       |   Bitcoin                   |  Crypto Price                       |
| Ethereum      |   Ethereum                  |  Crypto Price                       |
| USDT          |   Tether                    |  Crypto Price                       |
| XRP           |   Ripple                    |  Crypto price                       |
| Barnbridge    |   BARNBRIDGE                |  Farming Pool Rate: Stablecoin Pool |
| yearn.finance |   YFI                       |  Farming Pool Rate: WETH Pool       |

To retrieve data, query the oracle for an asset as listed in column "Data Field Query".
The query string is case-sensitive.

## Data Access

All asset prices are determined in USD according to our [methodology](https://docs.diadata.org/documentation/methodology).
They are denominated in a fix comma format with 5 decimals, so you need to divide them by 100000 to retrieve the actual value in USD.
Where applicable, the oracle also provides information on circulating supply and the timestamp of data collection.
The query in the smart contract is realized with the symbol of the asset.

### Smart Contract

DIA data is published in the `DIAOracle` smart contract.
By querying the `getCoinInfo()` function you can retrieve the requested data.
You can interact with our contract in the [blockchain explorer](https://explorer.celo.org/address/0xCd8E18890E416Aa7ab09aa793b406C187747C687), where you can query for any supported asset.

The contract takes the full name of the asset as input, e.g., `Bitcoin` and returns this struct of data:

```
struct CoinInfo {
	uint256 price;
	uint256 supply;
	uint256 lastUpdateTimestamp;
	string symbol;
}
```

### Write your own DApp

To access oracle data, you can either use the explorer above or write your own contract and reference the oracle.
The following snippet shows how to retrieve the price of an asset (e.g. `DOT`) measured in another asset (BTC) using a wrapper smart contract.
On Celo, you can initialize the oracle address in your wrapper contract by calling `setOracleAddress()` with the address of our deployed demo contract `0xCd8E18890E416Aa7ab09aa793b406C187747C687`.

```
pragma solidity ^0.4.24;

contract DiaOracle {
	address owner;

	struct CoinInfo {
		uint256 price;
		uint256 supply;
		uint256 lastUpdateTimestamp;
		string symbol;
	}

	mapping(string => CoinInfo) diaOracles;

	event newCoinInfo(
		string name,
		string symbol,
		uint256 price,
		uint256 supply,
		uint256 lastUpdateTimestamp
	);

	constructor() public {
		owner = msg.sender;
	}

	function changeOwner(address newOwner) public {
		require(msg.sender == owner);
		owner = newOwner;
	}

	function updateCoinInfo(string name, string symbol, uint256 newPrice, uint256 newSupply, uint256 newTimestamp) public {
		require(msg.sender == owner);
		diaOracles[name] = (CoinInfo(newPrice, newSupply, newTimestamp, symbol));
		emit newCoinInfo(name, symbol, newPrice, newSupply, newTimestamp);
	}

	function getCoinInfo(string name) public view returns (uint256, uint256, uint256, string) {
		return (
			diaOracles[name].price,
			diaOracles[name].supply,
			diaOracles[name].lastUpdateTimestamp,
			diaOracles[name].symbol
		);
	}
}

contract DiaAssetBtcOracle {
	DiaOracle oracle;
	address owner;
    
	constructor() public {
		owner = msg.sender;
	}
    
	function setOracleAddress(address _address) public {
		require(msg.sender == owner);
		oracle = DiaOracle(_address);
	}
    
	function getAssetBtcRate(string asset) constant public returns (uint256) {
		(uint assetPrice,,,) = oracle.getCoinInfo(asset);
		(uint btcPrice,,,) = oracle.getCoinInfo("Bitcoin");
		return (assetPrice / btcPrice);
	}
    
}
```

#### Deployed Addresses

| Chain name    |        Oracle Contract Address          | Update Frequency |
| :------------ | :------------------------------------------: | :----------: |
| Celo | `0xCd8E18890E416Aa7ab09aa793b406C187747C687` |    1/day    |

### Github and Contact

DIA provides a broad range of assets. You can find an overview in the DIA documentation [here](https://docs.diadata.org/documentation/oracle-documentation).
All our code is open-source and can be found on our [Github repositoy](https://github.com/diadata-org/diadata).
For the deployment of specific oracles (source/methodology/frequency) please [contact the DIA team](mailto:bd@diadata.org).

You can follow us on [Telegram](https://t.me/DIAdata_org), [Github](https://github.com/diadata-org), and [Medium](https://medium.com/dia-insights).
