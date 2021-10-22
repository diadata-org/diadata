---
description: >-
  This page contains an overview over the different smart contract interfaces
  deployed by DIA.
---

# Access the Oracle

DIA uses two different smart contracts for its oracles, depending on the use cases for each deployment. Its definitions are listed here. To learn how to access a foreign smart contract (such as our oracles) from your own smart contract, have a look at [this tutorial](https://ethereum.org/en/developers/tutorials/interact-with-other-contracts-from-solidity/).

### DIA Key-Value Oracle Contract

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
```

To access the key-value oracle, you first need to know the key of the data you are interested in. On querying `getValue(key)`, the oracle returns the latest value for that key and the timestamp of its last update. This key is a string and each update emits an event containing the key of the updated value. Typical keys are listed in the details section of each deployed contract.

The response value is an integer in a fix comma format and the timestamp associated with each value is a Unix timestamp for the UTC timezone.

### DIA CoinInfo Oracle Contract

```
pragma solidity ^0.4.21;

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
```

The CoinInfo oracle contract provides several information about an asset. Apart from the asset price, a response also contains the symbol name, its full name and (if applicable) its circulating supply. A timestamp of the last update is returned as well.

A call of `getCoinInfo(name)` triggers the response and it can then be used in other smart contracts. The timestamp associated with each response is a Unix timestamp for the UTC timezone.
