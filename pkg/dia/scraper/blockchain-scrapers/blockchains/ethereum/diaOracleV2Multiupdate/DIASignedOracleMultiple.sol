// SPDX-License-Identifier: MIT

pragma solidity ^0.8.23;

import "@openzeppelin/contracts/access/Ownable.sol";

contract DIASignedOracleMultiple is Ownable {
    address public signer;
    struct Message {
        string Symbol;
        address Address;
        string Blockchain;
        uint256 Price;
        uint256 Time;
    }

    struct RawMessage {
        string Symbol;
        address Address;
        string Blockchain;
        uint256 Price;
        uint256 Time;
        bytes Signature;
    }

    mapping(string => Message) public values;

    event OracleUpdate(string key, uint256 value, uint256 timestamp);
 
     event RequestUpdate(
        address assetAddress,
        string blockchain,
        uint256 timestamp
    );

    constructor() Ownable(msg.sender) {
        signer = msg.sender;
    }

    function setValues(RawMessage[] memory vs) public {
        for (uint i = 0; i < vs.length; i++) {
            this.setValue(vs[i]);
        }
    }

    function setValue(RawMessage memory v) external {
        require(getSigner(v) == signer, "SIGNER_DONT_MATCH");
        require(v.Time >= values[v.Symbol].Time, "NEWER_TIMESTAMP_REQUIRED");

        Message memory valueWithoutSignature = Message({
            Symbol: v.Symbol,
            Address: v.Address,
            Blockchain: v.Blockchain,
            Price: v.Price,
            Time: v.Time
        });

        values[v.Symbol] = valueWithoutSignature;
        emit OracleUpdate(v.Symbol, v.Price, v.Time);
    }

    function requestForUpdate(
        address assetAddress,
        string memory blockchain,
        uint256 timestamp
    ) external {
        emit RequestUpdate(assetAddress, blockchain, timestamp);
    }

    function getValue(
        string memory key
    ) external view returns (Message memory) {
        return values[key];
    }

    function updateSigner(address _newSigner) external onlyOwner {
        signer = _newSigner;
    }

    function changeOwner(address newOwner) external onlyOwner {
        _transferOwnership(newOwner);
    }

    function getSigner(RawMessage memory message) public pure returns (address) {
        bytes memory _signature;
        _signature = message.Signature;

        string
            memory EIP712_DOMAIN_TYPE = "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)";
        string
            memory MESSAGE_TYPE = "Message(string Symbol,address Address,string Blockchain,uint256 Price,uint256 Time)";

        bytes32 DOMAIN_SEPARATOR = keccak256(
            abi.encode(
                keccak256(abi.encodePacked(EIP712_DOMAIN_TYPE)),
                keccak256(abi.encodePacked("DiaData")),
                keccak256(abi.encodePacked("1.0.0")),
                1,
                0x0000000000000000000000000000000000000000
            )
        );

        bytes32 hash = keccak256(
            abi.encodePacked(
                "\x19\x01", // backslash is needed to escape the character
                DOMAIN_SEPARATOR,
                keccak256(
                    abi.encode(
                        keccak256(abi.encodePacked(MESSAGE_TYPE)),
                        keccak256(bytes(message.Symbol)),
                        message.Address,
                        keccak256(bytes(message.Blockchain)),
                        message.Price,
                        message.Time
                    )
                )
            )
        );

        // split signature
        bytes32 r;
        bytes32 s;
        uint8 v;
        if (_signature.length != 65) {
            return address(0);
        }
        assembly {
            r := mload(add(_signature, 32))
            s := mload(add(_signature, 64))
            v := byte(0, mload(add(_signature, 96)))
        }
        if (v < 27) {
            v += 27;
        }
        if (v != 27 && v != 28) {
            return address(0);
        } else {
            // verify
            return ecrecover(hash, v, r, s);
        }
    }
}