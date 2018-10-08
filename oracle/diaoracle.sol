pragma solidity ^0.4.21;

contract DIAOracle {
    mapping (uint => string) public names;
    mapping (uint => uint) public prices;
    mapping (uint => uint) supplies;
    uint256 last_update_timestamp;
    address owner;
    
    event newParameters(
        string name,
        uint256 price,
        uint256 supply,
        uint256 timestamp
    );
    
    constructor() public {
        last_update_timestamp = 0;
        owner = msg.sender;
    }
    
    function updateParameters(uint asset_index, string name, uint256 new_price, uint256 new_supply, uint256 new_timestamp) public {
        require(msg.sender == owner);
        names[asset_index] = name;
        prices[asset_index] = new_price;
        supplies[asset_index] = new_supply;
        last_update_timestamp = new_timestamp;
        emit newParameters(name,new_price,new_supply,new_timestamp);
    }
    
    function getParameters(uint asset_index) public view returns (uint256, uint256, uint256) {
        return (prices[asset_index],supplies[asset_index],last_update_timestamp);
    }
}
