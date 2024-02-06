// compiled using solidity 0.8.19
pragma solidity 0.8.19;

contract DIAOracleV2 {
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

    function setMultipleValues(string[] memory keys, uint256[] memory compressedValues) public {
        require(msg.sender == oracleUpdater);
        require(keys.length == compressedValues.length);
        
        for (uint128 i = 0; i < keys.length; i++) {
            string memory currentKey = keys[i];
            uint256 currentCvalue = compressedValues[i];
            uint128 value = (uint128)(currentCvalue >> 128);
            uint128 timestamp = (uint128)(currentCvalue % 2**128);

            values[currentKey] = currentCvalue;
            emit OracleUpdate(currentKey, value, timestamp);
        }
    }
    
    function getValue(string memory key) external view returns (uint128, uint128) {
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

