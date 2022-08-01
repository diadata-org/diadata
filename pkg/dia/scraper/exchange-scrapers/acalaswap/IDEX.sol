// SPDX-License-Identifier: GPL-3.0-or-later

pragma solidity ^0.8.0;

interface IDEX {
    event Swaped(address indexed sender, address[] path, uint256 supplyAmount, uint256 targetAmount);
    event AddedLiquidity(
        address indexed sender,
        address indexed tokenA,
        address indexed tokenB,
        uint256 maxAmountA,
        uint256 maxAmountB
    );
    event RemovedLiquidity(address indexed sender, address indexed tokenA, address indexed tokenB, uint256 removeShare);

    // Get liquidity pool of the currency_id_a and currency_id_b.
    // Returns (liquidity_a, liquidity_b).
    function getLiquidityPool(address tokenA, address tokenB) external view returns (uint256, uint256);

    // Get Liquidity token address.
    // Returns (liquidity_token_address). Return address(0x0) if the liquidity token address is not mapped.
    function getLiquidityTokenAddress(address tokenA, address tokenB) external view returns (address);

    // Get swap target amount.
    // Returns (target_amount). Returns 0 if getting the target amount fails.
    function getSwapTargetAmount(address[] calldata path, uint256 supplyAmount) external view returns (uint256);

    // Get swap supply amount.
    // Returns (supply_amount). Returns 0 if getting the supply amount fails.
    function getSwapSupplyAmount(address[] calldata path, uint256 targetAmount) external view returns (uint256);

    // Swap with exact supply.
    // Returns a boolean value indicating whether the operation succeeded.
    function swapWithExactSupply(
        address[] calldata path,
        uint256 supplyAmount,
        uint256 minTargetAmount
    ) external returns (bool);

    // Swap with exact target.
    // Returns a boolean value indicating whether the operation succeeded.
    function swapWithExactTarget(
        address[] calldata path,
        uint256 targetAmount,
        uint256 maxSupplyAmount
    ) external returns (bool);

    // Add liquidity to the trading pair.
    // Returns a boolean value indicating whether the operation succeeded.
    function addLiquidity(
        address tokenA,
        address tokenB,
        uint256 maxAmountA,
        uint256 maxAmountB,
        uint256 minShareIncrement
    ) external returns (bool);

    // Remove liquidity from the trading pair.
    // Returns a boolean value indicating whether the operation succeeded.
    function removeLiquidity(
        address tokenA,
        address tokenB,
        uint256 removeShare,
        uint256 minWithdrawnA,
        uint256 minWithdrawnB
    ) external returns (bool);
}
