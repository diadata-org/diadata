// SPDX-License-Identifier: GPL-3.0-or-later

pragma solidity ^0.8.0;

import "./IDEX.sol";

contract DEX is IDEX {
    address constant private PRECOMPILE = address(0x0000000000000000000000000000000000000405);

    /**
     * @dev Get liquidity pool of the currency_id_a and currency_id_b.
     * Returns (liquidity_a, liquidity_b).
     */
    function getLiquidityPool(address tokenA, address tokenB)
    public
    view
    override
    returns (uint256, uint256)
    {
        require(tokenA != address(0), "DEX: tokenA is zero address");
        require(tokenB != address(0), "DEX: tokenB is zero address");

        (bool success, bytes memory returnData) = PRECOMPILE.staticcall(
            abi.encodeWithSignature("getLiquidityPool(address,address)", tokenA, tokenB)
        );
        assembly {
            if eq(success, 0) {
                revert(add(returnData, 0x20), returndatasize())
            }
        }

        return abi.decode(returnData, (uint256, uint256));
    }

    /**
     * @dev Get Liquidity token address.
     * Returns (liquidity_token_address). Return address(0x0) if the liquidity token address is not mapped.
     */
    function getLiquidityTokenAddress(address tokenA, address tokenB)
    public
    view
    override
    returns (address) {
        require(tokenA != address(0), "DEX: tokenA is zero address");
        require(tokenB != address(0), "DEX: tokenB is zero address");

        (bool success, bytes memory returnData) = PRECOMPILE.staticcall(
            abi.encodeWithSignature("getLiquidityTokenAddress(address,address)", tokenA, tokenB)
        );
        assembly {
            if eq(success, 0) {
                revert(add(returnData, 0x20), returndatasize())
            }
        }

        return abi.decode(returnData, (address));
    }


    /**
     * @dev Get swap target amount.
     * Returns (target_amount). Returns 0 if getting the target amount fails.
     */
    function getSwapTargetAmount(address[] memory path, uint256 supplyAmount)
    public
    view
    override
    returns (uint256) {
        for (uint i = 0; i < path.length; i++) {
            require(path[i] != address(0), "DEX: token is zero address");
        }
        require(supplyAmount != 0, "DEX: supplyAmount is zero");

        (bool success, bytes memory returnData) = PRECOMPILE.staticcall(
            abi.encodeWithSignature("getSwapTargetAmount(address[],uint256)", path, supplyAmount)
        );
        assembly {
            if eq(success, 0) {
                revert(add(returnData, 0x20), returndatasize())
            }
        }

        return abi.decode(returnData, (uint256));
    }

    /**
     * @dev Get swap supply amount.
     * Returns (supply_amount). Returns 0 if getting the supply amount fails.
     */
    function getSwapSupplyAmount(address[] memory path, uint256 targetAmount)
    public
    view
    override
    returns (uint256) {
        for (uint i = 0; i < path.length; i++) {
            require(path[i] != address(0), "DEX: token is zero address");
        }
        require(targetAmount != 0, "DEX: targetAmount is zero");

        (bool success, bytes memory returnData) = PRECOMPILE.staticcall(
            abi.encodeWithSignature("getSwapSupplyAmount(address[],uint256)", path, targetAmount)
        );
        assembly {
            if eq(success, 0) {
                revert(add(returnData, 0x20), returndatasize())
            }
        }

        return abi.decode(returnData, (uint256));
    }

    /**
     * @dev Swap with exact supply.
     * Returns a boolean value indicating whether the operation succeeded.
     */
    function swapWithExactSupply(address[] memory path, uint256 supplyAmount, uint256 minTargetAmount)
    public
    override
    returns (bool) {
        for (uint i = 0; i < path.length; i++) {
            require(path[i] != address(0), "DEX: token is zero address");
        }
        require(supplyAmount != 0, "DEX: supplyAmount is zero");

        (bool success, bytes memory returnData) = PRECOMPILE.call(
            abi.encodeWithSignature(
                "swapWithExactSupply(address,address[],uint256,uint256)",
                msg.sender, path, supplyAmount, minTargetAmount
            )
        );
        assembly {
            if eq(success, 0) {
                revert(add(returnData, 0x20), returndatasize())
            }
        }

        emit Swaped(msg.sender, path, supplyAmount, abi.decode(returnData, (uint256)));
        return true;
    }

    /**
     * @dev Swap with exact target.
     * Returns a boolean value indicating whether the operation succeeded.
     */
    function swapWithExactTarget(address[] memory path, uint256 targetAmount, uint256 maxSupplyAmount)
    public
    override
    returns (bool) {
        for (uint i = 0; i < path.length; i++) {
            require(path[i] != address(0), "DEX: token is zero address");
        }
        require(targetAmount != 0, "DEX: targetAmount is zero");

        (bool success, bytes memory returnData) = PRECOMPILE.call(
            abi.encodeWithSignature(
                "swapWithExactTarget(address,address[],uint256,uint256)",
                msg.sender, path, targetAmount, maxSupplyAmount
            )
        );
        assembly {
            if eq(success, 0) {
                revert(add(returnData, 0x20), returndatasize())
            }
        }

        emit Swaped(msg.sender, path, abi.decode(returnData, (uint256)), targetAmount);
        return true;
    }

    /**
     * @dev Add liquidity to the trading pair.
     * Returns a boolean value indicating whether the operation succeeded.
     */
    function addLiquidity(
        address tokenA,
        address tokenB,
        uint256 maxAmountA,
        uint256 maxAmountB,
        uint256 minShareIncrement
    )
    public
    override
    returns (bool) {
        require(tokenA != address(0), "DEX: tokenA is zero address");
        require(tokenB != address(0), "DEX: tokenB is zero address");
        require(maxAmountA != 0, "DEX: maxAmountA is zero");
        require(maxAmountB != 0, "DEX: maxAmountB is zero");

        (bool success, bytes memory returnData) = PRECOMPILE.call(
            abi.encodeWithSignature(
                "addLiquidity(address,address,address,uint256,uint256,uint256)",
                msg.sender, tokenA, tokenB, maxAmountA, maxAmountB, minShareIncrement
            )
        );
        assembly {
            if eq(success, 0) {
                revert(add(returnData, 0x20), returndatasize())
            }
        }

        emit AddedLiquidity(msg.sender, tokenA, tokenB, maxAmountA, maxAmountB);
        return true;
    }

    /**
     * @dev Remove liquidity from the trading pair.
     * Returns a boolean value indicating whether the operation succeeded.
     */
    function removeLiquidity(
        address tokenA,
        address tokenB,
        uint256 removeShare,
        uint256 minWithdrawnA,
        uint256 minWithdrawnB
    )
    public
    override
    returns (bool) {
        require(tokenA != address(0), "DEX: tokenA is zero address");
        require(tokenB != address(0), "DEX: tokenB is zero address");
        require(removeShare != 0, "DEX: removeShare is zero");

        (bool success, bytes memory returnData) = PRECOMPILE.call(
            abi.encodeWithSignature(
                "removeLiquidity(address,address,address,uint256,uint256,uint256)",
                msg.sender, tokenA, tokenB, removeShare, minWithdrawnA, minWithdrawnB
            )
        );
        assembly {
            if eq(success, 0) {
                revert(add(returnData, 0x20), returndatasize())
            }
        }

        emit RemovedLiquidity(msg.sender, tokenA, tokenB, removeShare);
        return true;
    }
}