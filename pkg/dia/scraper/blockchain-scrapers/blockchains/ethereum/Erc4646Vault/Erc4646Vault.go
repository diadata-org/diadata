// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Erc4646Vault

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Erc4646VaultMetaData contains all meta data concerning the Erc4646Vault contract.
var Erc4646VaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"AssetAddedToEmissionList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"AssetRemovedFromEmissionList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pendingReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardIntegral\",\"type\":\"uint256\"}],\"name\":\"RewardIntegralForAccountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardIntegral\",\"type\":\"uint256\"}],\"name\":\"RewardIntegralUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardIntegral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"claimStartTime\",\"type\":\"uint32\"}],\"name\":\"RewardParamsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardVault\",\"type\":\"address\"}],\"name\":\"RewardVaultSetting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INITIAL_DEPOSIT_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHARE_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHARE_SYMBOL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDERLYING\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_\",\"type\":\"address\"}],\"name\":\"addAssetToEmissionList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"claimableReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"emissionList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emissionListLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardVault\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isClaimStart\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdate\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_\",\"type\":\"address\"}],\"name\":\"removeAssetFromEmissionList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"rewardIntegralFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"rewardParamsMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardIntegral\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"claimStartTime\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardVault\",\"outputs\":[{\"internalType\":\"contractIRewardVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"claimStartTime\",\"type\":\"uint32\"}],\"name\":\"setRewardParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardVault\",\"type\":\"address\"}],\"name\":\"setRewardVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"simulateDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"simulateMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"simulateRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"simulateWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"storedPendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Erc4646VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc4646VaultMetaData.ABI instead.
var Erc4646VaultABI = Erc4646VaultMetaData.ABI

// Erc4646Vault is an auto generated Go binding around an Ethereum contract.
type Erc4646Vault struct {
	Erc4646VaultCaller     // Read-only binding to the contract
	Erc4646VaultTransactor // Write-only binding to the contract
	Erc4646VaultFilterer   // Log filterer for contract events
}

// Erc4646VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc4646VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc4646VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc4646VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc4646VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc4646VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc4646VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc4646VaultSession struct {
	Contract     *Erc4646Vault     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc4646VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc4646VaultCallerSession struct {
	Contract *Erc4646VaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Erc4646VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc4646VaultTransactorSession struct {
	Contract     *Erc4646VaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Erc4646VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc4646VaultRaw struct {
	Contract *Erc4646Vault // Generic contract binding to access the raw methods on
}

// Erc4646VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc4646VaultCallerRaw struct {
	Contract *Erc4646VaultCaller // Generic read-only contract binding to access the raw methods on
}

// Erc4646VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc4646VaultTransactorRaw struct {
	Contract *Erc4646VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc4646Vault creates a new instance of Erc4646Vault, bound to a specific deployed contract.
func NewErc4646Vault(address common.Address, backend bind.ContractBackend) (*Erc4646Vault, error) {
	contract, err := bindErc4646Vault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc4646Vault{Erc4646VaultCaller: Erc4646VaultCaller{contract: contract}, Erc4646VaultTransactor: Erc4646VaultTransactor{contract: contract}, Erc4646VaultFilterer: Erc4646VaultFilterer{contract: contract}}, nil
}

// NewErc4646VaultCaller creates a new read-only instance of Erc4646Vault, bound to a specific deployed contract.
func NewErc4646VaultCaller(address common.Address, caller bind.ContractCaller) (*Erc4646VaultCaller, error) {
	contract, err := bindErc4646Vault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultCaller{contract: contract}, nil
}

// NewErc4646VaultTransactor creates a new write-only instance of Erc4646Vault, bound to a specific deployed contract.
func NewErc4646VaultTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc4646VaultTransactor, error) {
	contract, err := bindErc4646Vault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultTransactor{contract: contract}, nil
}

// NewErc4646VaultFilterer creates a new log filterer instance of Erc4646Vault, bound to a specific deployed contract.
func NewErc4646VaultFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc4646VaultFilterer, error) {
	contract, err := bindErc4646Vault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultFilterer{contract: contract}, nil
}

// bindErc4646Vault binds a generic wrapper to an already deployed contract.
func bindErc4646Vault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc4646VaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc4646Vault *Erc4646VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc4646Vault.Contract.Erc4646VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc4646Vault *Erc4646VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Erc4646VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc4646Vault *Erc4646VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Erc4646VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc4646Vault *Erc4646VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc4646Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc4646Vault *Erc4646VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc4646Vault *Erc4646VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.contract.Transact(opts, method, params...)
}

// INITIALDEPOSITLIMIT is a free data retrieval call binding the contract method 0xd16002da.
//
// Solidity: function INITIAL_DEPOSIT_LIMIT() view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) INITIALDEPOSITLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "INITIAL_DEPOSIT_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALDEPOSITLIMIT is a free data retrieval call binding the contract method 0xd16002da.
//
// Solidity: function INITIAL_DEPOSIT_LIMIT() view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) INITIALDEPOSITLIMIT() (*big.Int, error) {
	return _Erc4646Vault.Contract.INITIALDEPOSITLIMIT(&_Erc4646Vault.CallOpts)
}

// INITIALDEPOSITLIMIT is a free data retrieval call binding the contract method 0xd16002da.
//
// Solidity: function INITIAL_DEPOSIT_LIMIT() view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) INITIALDEPOSITLIMIT() (*big.Int, error) {
	return _Erc4646Vault.Contract.INITIALDEPOSITLIMIT(&_Erc4646Vault.CallOpts)
}

// SHARENAME is a free data retrieval call binding the contract method 0xc2aa8164.
//
// Solidity: function SHARE_NAME() view returns(string)
func (_Erc4646Vault *Erc4646VaultCaller) SHARENAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "SHARE_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SHARENAME is a free data retrieval call binding the contract method 0xc2aa8164.
//
// Solidity: function SHARE_NAME() view returns(string)
func (_Erc4646Vault *Erc4646VaultSession) SHARENAME() (string, error) {
	return _Erc4646Vault.Contract.SHARENAME(&_Erc4646Vault.CallOpts)
}

// SHARENAME is a free data retrieval call binding the contract method 0xc2aa8164.
//
// Solidity: function SHARE_NAME() view returns(string)
func (_Erc4646Vault *Erc4646VaultCallerSession) SHARENAME() (string, error) {
	return _Erc4646Vault.Contract.SHARENAME(&_Erc4646Vault.CallOpts)
}

// SHARESYMBOL is a free data retrieval call binding the contract method 0xe6f33f40.
//
// Solidity: function SHARE_SYMBOL() view returns(string)
func (_Erc4646Vault *Erc4646VaultCaller) SHARESYMBOL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "SHARE_SYMBOL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SHARESYMBOL is a free data retrieval call binding the contract method 0xe6f33f40.
//
// Solidity: function SHARE_SYMBOL() view returns(string)
func (_Erc4646Vault *Erc4646VaultSession) SHARESYMBOL() (string, error) {
	return _Erc4646Vault.Contract.SHARESYMBOL(&_Erc4646Vault.CallOpts)
}

// SHARESYMBOL is a free data retrieval call binding the contract method 0xe6f33f40.
//
// Solidity: function SHARE_SYMBOL() view returns(string)
func (_Erc4646Vault *Erc4646VaultCallerSession) SHARESYMBOL() (string, error) {
	return _Erc4646Vault.Contract.SHARESYMBOL(&_Erc4646Vault.CallOpts)
}

// UNDERLYING is a free data retrieval call binding the contract method 0xc5d664c6.
//
// Solidity: function UNDERLYING() view returns(address)
func (_Erc4646Vault *Erc4646VaultCaller) UNDERLYING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "UNDERLYING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNDERLYING is a free data retrieval call binding the contract method 0xc5d664c6.
//
// Solidity: function UNDERLYING() view returns(address)
func (_Erc4646Vault *Erc4646VaultSession) UNDERLYING() (common.Address, error) {
	return _Erc4646Vault.Contract.UNDERLYING(&_Erc4646Vault.CallOpts)
}

// UNDERLYING is a free data retrieval call binding the contract method 0xc5d664c6.
//
// Solidity: function UNDERLYING() view returns(address)
func (_Erc4646Vault *Erc4646VaultCallerSession) UNDERLYING() (common.Address, error) {
	return _Erc4646Vault.Contract.UNDERLYING(&_Erc4646Vault.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Erc4646Vault *Erc4646VaultCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Erc4646Vault *Erc4646VaultSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Erc4646Vault.Contract.UPGRADEINTERFACEVERSION(&_Erc4646Vault.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Erc4646Vault *Erc4646VaultCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Erc4646Vault.Contract.UPGRADEINTERFACEVERSION(&_Erc4646Vault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.Allowance(&_Erc4646Vault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.Allowance(&_Erc4646Vault.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Erc4646Vault *Erc4646VaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Erc4646Vault *Erc4646VaultSession) Asset() (common.Address, error) {
	return _Erc4646Vault.Contract.Asset(&_Erc4646Vault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Erc4646Vault *Erc4646VaultCallerSession) Asset() (common.Address, error) {
	return _Erc4646Vault.Contract.Asset(&_Erc4646Vault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.BalanceOf(&_Erc4646Vault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.BalanceOf(&_Erc4646Vault.CallOpts, account)
}

// ClaimableReward is a free data retrieval call binding the contract method 0xd26abffa.
//
// Solidity: function claimableReward(address token, address account) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) ClaimableReward(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "claimableReward", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableReward is a free data retrieval call binding the contract method 0xd26abffa.
//
// Solidity: function claimableReward(address token, address account) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) ClaimableReward(token common.Address, account common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.ClaimableReward(&_Erc4646Vault.CallOpts, token, account)
}

// ClaimableReward is a free data retrieval call binding the contract method 0xd26abffa.
//
// Solidity: function claimableReward(address token, address account) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) ClaimableReward(token common.Address, account common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.ClaimableReward(&_Erc4646Vault.CallOpts, token, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.ConvertToAssets(&_Erc4646Vault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.ConvertToAssets(&_Erc4646Vault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.ConvertToShares(&_Erc4646Vault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.ConvertToShares(&_Erc4646Vault.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc4646Vault *Erc4646VaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc4646Vault *Erc4646VaultSession) Decimals() (uint8, error) {
	return _Erc4646Vault.Contract.Decimals(&_Erc4646Vault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc4646Vault *Erc4646VaultCallerSession) Decimals() (uint8, error) {
	return _Erc4646Vault.Contract.Decimals(&_Erc4646Vault.CallOpts)
}

// EmissionList is a free data retrieval call binding the contract method 0x7db808e9.
//
// Solidity: function emissionList(uint256 ) view returns(address)
func (_Erc4646Vault *Erc4646VaultCaller) EmissionList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "emissionList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmissionList is a free data retrieval call binding the contract method 0x7db808e9.
//
// Solidity: function emissionList(uint256 ) view returns(address)
func (_Erc4646Vault *Erc4646VaultSession) EmissionList(arg0 *big.Int) (common.Address, error) {
	return _Erc4646Vault.Contract.EmissionList(&_Erc4646Vault.CallOpts, arg0)
}

// EmissionList is a free data retrieval call binding the contract method 0x7db808e9.
//
// Solidity: function emissionList(uint256 ) view returns(address)
func (_Erc4646Vault *Erc4646VaultCallerSession) EmissionList(arg0 *big.Int) (common.Address, error) {
	return _Erc4646Vault.Contract.EmissionList(&_Erc4646Vault.CallOpts, arg0)
}

// EmissionListLength is a free data retrieval call binding the contract method 0x1036f087.
//
// Solidity: function emissionListLength() view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) EmissionListLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "emissionListLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmissionListLength is a free data retrieval call binding the contract method 0x1036f087.
//
// Solidity: function emissionListLength() view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) EmissionListLength() (*big.Int, error) {
	return _Erc4646Vault.Contract.EmissionListLength(&_Erc4646Vault.CallOpts)
}

// EmissionListLength is a free data retrieval call binding the contract method 0x1036f087.
//
// Solidity: function emissionListLength() view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) EmissionListLength() (*big.Int, error) {
	return _Erc4646Vault.Contract.EmissionListLength(&_Erc4646Vault.CallOpts)
}

// IsClaimStart is a free data retrieval call binding the contract method 0xe0427c51.
//
// Solidity: function isClaimStart(address token) view returns(bool)
func (_Erc4646Vault *Erc4646VaultCaller) IsClaimStart(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "isClaimStart", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimStart is a free data retrieval call binding the contract method 0xe0427c51.
//
// Solidity: function isClaimStart(address token) view returns(bool)
func (_Erc4646Vault *Erc4646VaultSession) IsClaimStart(token common.Address) (bool, error) {
	return _Erc4646Vault.Contract.IsClaimStart(&_Erc4646Vault.CallOpts, token)
}

// IsClaimStart is a free data retrieval call binding the contract method 0xe0427c51.
//
// Solidity: function isClaimStart(address token) view returns(bool)
func (_Erc4646Vault *Erc4646VaultCallerSession) IsClaimStart(token common.Address) (bool, error) {
	return _Erc4646Vault.Contract.IsClaimStart(&_Erc4646Vault.CallOpts, token)
}

// LastUpdate is a free data retrieval call binding the contract method 0xc0463711.
//
// Solidity: function lastUpdate() view returns(uint32)
func (_Erc4646Vault *Erc4646VaultCaller) LastUpdate(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "lastUpdate")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastUpdate is a free data retrieval call binding the contract method 0xc0463711.
//
// Solidity: function lastUpdate() view returns(uint32)
func (_Erc4646Vault *Erc4646VaultSession) LastUpdate() (uint32, error) {
	return _Erc4646Vault.Contract.LastUpdate(&_Erc4646Vault.CallOpts)
}

// LastUpdate is a free data retrieval call binding the contract method 0xc0463711.
//
// Solidity: function lastUpdate() view returns(uint32)
func (_Erc4646Vault *Erc4646VaultCallerSession) LastUpdate() (uint32, error) {
	return _Erc4646Vault.Contract.LastUpdate(&_Erc4646Vault.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.MaxDeposit(&_Erc4646Vault.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.MaxDeposit(&_Erc4646Vault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.MaxMint(&_Erc4646Vault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.MaxMint(&_Erc4646Vault.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.MaxRedeem(&_Erc4646Vault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.MaxRedeem(&_Erc4646Vault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.MaxWithdraw(&_Erc4646Vault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.MaxWithdraw(&_Erc4646Vault.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc4646Vault *Erc4646VaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc4646Vault *Erc4646VaultSession) Name() (string, error) {
	return _Erc4646Vault.Contract.Name(&_Erc4646Vault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc4646Vault *Erc4646VaultCallerSession) Name() (string, error) {
	return _Erc4646Vault.Contract.Name(&_Erc4646Vault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc4646Vault *Erc4646VaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc4646Vault *Erc4646VaultSession) Owner() (common.Address, error) {
	return _Erc4646Vault.Contract.Owner(&_Erc4646Vault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc4646Vault *Erc4646VaultCallerSession) Owner() (common.Address, error) {
	return _Erc4646Vault.Contract.Owner(&_Erc4646Vault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Erc4646Vault *Erc4646VaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Erc4646Vault *Erc4646VaultSession) Paused() (bool, error) {
	return _Erc4646Vault.Contract.Paused(&_Erc4646Vault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Erc4646Vault *Erc4646VaultCallerSession) Paused() (bool, error) {
	return _Erc4646Vault.Contract.Paused(&_Erc4646Vault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.PreviewDeposit(&_Erc4646Vault.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.PreviewDeposit(&_Erc4646Vault.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.PreviewMint(&_Erc4646Vault.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.PreviewMint(&_Erc4646Vault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.PreviewRedeem(&_Erc4646Vault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.PreviewRedeem(&_Erc4646Vault.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.PreviewWithdraw(&_Erc4646Vault.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.PreviewWithdraw(&_Erc4646Vault.CallOpts, assets)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Erc4646Vault *Erc4646VaultCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Erc4646Vault *Erc4646VaultSession) ProxiableUUID() ([32]byte, error) {
	return _Erc4646Vault.Contract.ProxiableUUID(&_Erc4646Vault.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Erc4646Vault *Erc4646VaultCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Erc4646Vault.Contract.ProxiableUUID(&_Erc4646Vault.CallOpts)
}

// RewardIntegralFor is a free data retrieval call binding the contract method 0x63f57777.
//
// Solidity: function rewardIntegralFor(address token, address account) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) RewardIntegralFor(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "rewardIntegralFor", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardIntegralFor is a free data retrieval call binding the contract method 0x63f57777.
//
// Solidity: function rewardIntegralFor(address token, address account) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) RewardIntegralFor(token common.Address, account common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.RewardIntegralFor(&_Erc4646Vault.CallOpts, token, account)
}

// RewardIntegralFor is a free data retrieval call binding the contract method 0x63f57777.
//
// Solidity: function rewardIntegralFor(address token, address account) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) RewardIntegralFor(token common.Address, account common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.RewardIntegralFor(&_Erc4646Vault.CallOpts, token, account)
}

// RewardParamsMap is a free data retrieval call binding the contract method 0x227ef81f.
//
// Solidity: function rewardParamsMap(address token) view returns(uint256 rewardRate, uint256 rewardIntegral, uint32 claimStartTime)
func (_Erc4646Vault *Erc4646VaultCaller) RewardParamsMap(opts *bind.CallOpts, token common.Address) (struct {
	RewardRate     *big.Int
	RewardIntegral *big.Int
	ClaimStartTime uint32
}, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "rewardParamsMap", token)

	outstruct := new(struct {
		RewardRate     *big.Int
		RewardIntegral *big.Int
		ClaimStartTime uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardRate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardIntegral = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ClaimStartTime = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// RewardParamsMap is a free data retrieval call binding the contract method 0x227ef81f.
//
// Solidity: function rewardParamsMap(address token) view returns(uint256 rewardRate, uint256 rewardIntegral, uint32 claimStartTime)
func (_Erc4646Vault *Erc4646VaultSession) RewardParamsMap(token common.Address) (struct {
	RewardRate     *big.Int
	RewardIntegral *big.Int
	ClaimStartTime uint32
}, error) {
	return _Erc4646Vault.Contract.RewardParamsMap(&_Erc4646Vault.CallOpts, token)
}

// RewardParamsMap is a free data retrieval call binding the contract method 0x227ef81f.
//
// Solidity: function rewardParamsMap(address token) view returns(uint256 rewardRate, uint256 rewardIntegral, uint32 claimStartTime)
func (_Erc4646Vault *Erc4646VaultCallerSession) RewardParamsMap(token common.Address) (struct {
	RewardRate     *big.Int
	RewardIntegral *big.Int
	ClaimStartTime uint32
}, error) {
	return _Erc4646Vault.Contract.RewardParamsMap(&_Erc4646Vault.CallOpts, token)
}

// RewardVault is a free data retrieval call binding the contract method 0x3a2c6777.
//
// Solidity: function rewardVault() view returns(address)
func (_Erc4646Vault *Erc4646VaultCaller) RewardVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "rewardVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardVault is a free data retrieval call binding the contract method 0x3a2c6777.
//
// Solidity: function rewardVault() view returns(address)
func (_Erc4646Vault *Erc4646VaultSession) RewardVault() (common.Address, error) {
	return _Erc4646Vault.Contract.RewardVault(&_Erc4646Vault.CallOpts)
}

// RewardVault is a free data retrieval call binding the contract method 0x3a2c6777.
//
// Solidity: function rewardVault() view returns(address)
func (_Erc4646Vault *Erc4646VaultCallerSession) RewardVault() (common.Address, error) {
	return _Erc4646Vault.Contract.RewardVault(&_Erc4646Vault.CallOpts)
}

// SimulateDeposit is a free data retrieval call binding the contract method 0x923c86e6.
//
// Solidity: function simulateDeposit(uint256 assets) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) SimulateDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "simulateDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SimulateDeposit is a free data retrieval call binding the contract method 0x923c86e6.
//
// Solidity: function simulateDeposit(uint256 assets) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) SimulateDeposit(assets *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.SimulateDeposit(&_Erc4646Vault.CallOpts, assets)
}

// SimulateDeposit is a free data retrieval call binding the contract method 0x923c86e6.
//
// Solidity: function simulateDeposit(uint256 assets) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) SimulateDeposit(assets *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.SimulateDeposit(&_Erc4646Vault.CallOpts, assets)
}

// SimulateMint is a free data retrieval call binding the contract method 0xf9f18df7.
//
// Solidity: function simulateMint(uint256 shares) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) SimulateMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "simulateMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SimulateMint is a free data retrieval call binding the contract method 0xf9f18df7.
//
// Solidity: function simulateMint(uint256 shares) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) SimulateMint(shares *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.SimulateMint(&_Erc4646Vault.CallOpts, shares)
}

// SimulateMint is a free data retrieval call binding the contract method 0xf9f18df7.
//
// Solidity: function simulateMint(uint256 shares) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) SimulateMint(shares *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.SimulateMint(&_Erc4646Vault.CallOpts, shares)
}

// SimulateRedeem is a free data retrieval call binding the contract method 0xec1ac75b.
//
// Solidity: function simulateRedeem(uint256 shares) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) SimulateRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "simulateRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SimulateRedeem is a free data retrieval call binding the contract method 0xec1ac75b.
//
// Solidity: function simulateRedeem(uint256 shares) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) SimulateRedeem(shares *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.SimulateRedeem(&_Erc4646Vault.CallOpts, shares)
}

// SimulateRedeem is a free data retrieval call binding the contract method 0xec1ac75b.
//
// Solidity: function simulateRedeem(uint256 shares) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) SimulateRedeem(shares *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.SimulateRedeem(&_Erc4646Vault.CallOpts, shares)
}

// SimulateWithdraw is a free data retrieval call binding the contract method 0xbb31203d.
//
// Solidity: function simulateWithdraw(uint256 assets) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) SimulateWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "simulateWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SimulateWithdraw is a free data retrieval call binding the contract method 0xbb31203d.
//
// Solidity: function simulateWithdraw(uint256 assets) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) SimulateWithdraw(assets *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.SimulateWithdraw(&_Erc4646Vault.CallOpts, assets)
}

// SimulateWithdraw is a free data retrieval call binding the contract method 0xbb31203d.
//
// Solidity: function simulateWithdraw(uint256 assets) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) SimulateWithdraw(assets *big.Int) (*big.Int, error) {
	return _Erc4646Vault.Contract.SimulateWithdraw(&_Erc4646Vault.CallOpts, assets)
}

// StoredPendingReward is a free data retrieval call binding the contract method 0x51210430.
//
// Solidity: function storedPendingReward(address token, address account) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) StoredPendingReward(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "storedPendingReward", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StoredPendingReward is a free data retrieval call binding the contract method 0x51210430.
//
// Solidity: function storedPendingReward(address token, address account) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) StoredPendingReward(token common.Address, account common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.StoredPendingReward(&_Erc4646Vault.CallOpts, token, account)
}

// StoredPendingReward is a free data retrieval call binding the contract method 0x51210430.
//
// Solidity: function storedPendingReward(address token, address account) view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) StoredPendingReward(token common.Address, account common.Address) (*big.Int, error) {
	return _Erc4646Vault.Contract.StoredPendingReward(&_Erc4646Vault.CallOpts, token, account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc4646Vault *Erc4646VaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc4646Vault *Erc4646VaultSession) Symbol() (string, error) {
	return _Erc4646Vault.Contract.Symbol(&_Erc4646Vault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc4646Vault *Erc4646VaultCallerSession) Symbol() (string, error) {
	return _Erc4646Vault.Contract.Symbol(&_Erc4646Vault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) TotalAssets() (*big.Int, error) {
	return _Erc4646Vault.Contract.TotalAssets(&_Erc4646Vault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) TotalAssets() (*big.Int, error) {
	return _Erc4646Vault.Contract.TotalAssets(&_Erc4646Vault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc4646Vault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) TotalSupply() (*big.Int, error) {
	return _Erc4646Vault.Contract.TotalSupply(&_Erc4646Vault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc4646Vault *Erc4646VaultCallerSession) TotalSupply() (*big.Int, error) {
	return _Erc4646Vault.Contract.TotalSupply(&_Erc4646Vault.CallOpts)
}

// AddAssetToEmissionList is a paid mutator transaction binding the contract method 0xf23331c6.
//
// Solidity: function addAssetToEmissionList(address asset_) returns()
func (_Erc4646Vault *Erc4646VaultTransactor) AddAssetToEmissionList(opts *bind.TransactOpts, asset_ common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "addAssetToEmissionList", asset_)
}

// AddAssetToEmissionList is a paid mutator transaction binding the contract method 0xf23331c6.
//
// Solidity: function addAssetToEmissionList(address asset_) returns()
func (_Erc4646Vault *Erc4646VaultSession) AddAssetToEmissionList(asset_ common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.AddAssetToEmissionList(&_Erc4646Vault.TransactOpts, asset_)
}

// AddAssetToEmissionList is a paid mutator transaction binding the contract method 0xf23331c6.
//
// Solidity: function addAssetToEmissionList(address asset_) returns()
func (_Erc4646Vault *Erc4646VaultTransactorSession) AddAssetToEmissionList(asset_ common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.AddAssetToEmissionList(&_Erc4646Vault.TransactOpts, asset_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc4646Vault *Erc4646VaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc4646Vault *Erc4646VaultSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Approve(&_Erc4646Vault.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc4646Vault *Erc4646VaultTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Approve(&_Erc4646Vault.TransactOpts, spender, value)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Erc4646Vault *Erc4646VaultTransactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "claimReward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Erc4646Vault *Erc4646VaultSession) ClaimReward() (*types.Transaction, error) {
	return _Erc4646Vault.Contract.ClaimReward(&_Erc4646Vault.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Erc4646Vault *Erc4646VaultTransactorSession) ClaimReward() (*types.Transaction, error) {
	return _Erc4646Vault.Contract.ClaimReward(&_Erc4646Vault.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_Erc4646Vault *Erc4646VaultTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Deposit(&_Erc4646Vault.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_Erc4646Vault *Erc4646VaultTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Deposit(&_Erc4646Vault.TransactOpts, assets, receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlying, address _rewardVault) returns()
func (_Erc4646Vault *Erc4646VaultTransactor) Initialize(opts *bind.TransactOpts, _underlying common.Address, _rewardVault common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "initialize", _underlying, _rewardVault)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlying, address _rewardVault) returns()
func (_Erc4646Vault *Erc4646VaultSession) Initialize(_underlying common.Address, _rewardVault common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Initialize(&_Erc4646Vault.TransactOpts, _underlying, _rewardVault)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlying, address _rewardVault) returns()
func (_Erc4646Vault *Erc4646VaultTransactorSession) Initialize(_underlying common.Address, _rewardVault common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Initialize(&_Erc4646Vault.TransactOpts, _underlying, _rewardVault)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_Erc4646Vault *Erc4646VaultTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Mint(&_Erc4646Vault.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_Erc4646Vault *Erc4646VaultTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Mint(&_Erc4646Vault.TransactOpts, shares, receiver)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Erc4646Vault *Erc4646VaultTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Erc4646Vault *Erc4646VaultSession) Pause() (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Pause(&_Erc4646Vault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Erc4646Vault *Erc4646VaultTransactorSession) Pause() (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Pause(&_Erc4646Vault.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_Erc4646Vault *Erc4646VaultTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Redeem(&_Erc4646Vault.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_Erc4646Vault *Erc4646VaultTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Redeem(&_Erc4646Vault.TransactOpts, shares, receiver, owner)
}

// RemoveAssetFromEmissionList is a paid mutator transaction binding the contract method 0xab571c60.
//
// Solidity: function removeAssetFromEmissionList(address asset_) returns()
func (_Erc4646Vault *Erc4646VaultTransactor) RemoveAssetFromEmissionList(opts *bind.TransactOpts, asset_ common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "removeAssetFromEmissionList", asset_)
}

// RemoveAssetFromEmissionList is a paid mutator transaction binding the contract method 0xab571c60.
//
// Solidity: function removeAssetFromEmissionList(address asset_) returns()
func (_Erc4646Vault *Erc4646VaultSession) RemoveAssetFromEmissionList(asset_ common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.RemoveAssetFromEmissionList(&_Erc4646Vault.TransactOpts, asset_)
}

// RemoveAssetFromEmissionList is a paid mutator transaction binding the contract method 0xab571c60.
//
// Solidity: function removeAssetFromEmissionList(address asset_) returns()
func (_Erc4646Vault *Erc4646VaultTransactorSession) RemoveAssetFromEmissionList(asset_ common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.RemoveAssetFromEmissionList(&_Erc4646Vault.TransactOpts, asset_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc4646Vault *Erc4646VaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc4646Vault *Erc4646VaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _Erc4646Vault.Contract.RenounceOwnership(&_Erc4646Vault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc4646Vault *Erc4646VaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Erc4646Vault.Contract.RenounceOwnership(&_Erc4646Vault.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Erc4646Vault *Erc4646VaultTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Erc4646Vault *Erc4646VaultSession) Resume() (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Resume(&_Erc4646Vault.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Erc4646Vault *Erc4646VaultTransactorSession) Resume() (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Resume(&_Erc4646Vault.TransactOpts)
}

// SetRewardParams is a paid mutator transaction binding the contract method 0xbd3e7d97.
//
// Solidity: function setRewardParams(address token, uint256 rewardRate, uint32 claimStartTime) returns()
func (_Erc4646Vault *Erc4646VaultTransactor) SetRewardParams(opts *bind.TransactOpts, token common.Address, rewardRate *big.Int, claimStartTime uint32) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "setRewardParams", token, rewardRate, claimStartTime)
}

// SetRewardParams is a paid mutator transaction binding the contract method 0xbd3e7d97.
//
// Solidity: function setRewardParams(address token, uint256 rewardRate, uint32 claimStartTime) returns()
func (_Erc4646Vault *Erc4646VaultSession) SetRewardParams(token common.Address, rewardRate *big.Int, claimStartTime uint32) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.SetRewardParams(&_Erc4646Vault.TransactOpts, token, rewardRate, claimStartTime)
}

// SetRewardParams is a paid mutator transaction binding the contract method 0xbd3e7d97.
//
// Solidity: function setRewardParams(address token, uint256 rewardRate, uint32 claimStartTime) returns()
func (_Erc4646Vault *Erc4646VaultTransactorSession) SetRewardParams(token common.Address, rewardRate *big.Int, claimStartTime uint32) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.SetRewardParams(&_Erc4646Vault.TransactOpts, token, rewardRate, claimStartTime)
}

// SetRewardVault is a paid mutator transaction binding the contract method 0x8125dd10.
//
// Solidity: function setRewardVault(address _rewardVault) returns()
func (_Erc4646Vault *Erc4646VaultTransactor) SetRewardVault(opts *bind.TransactOpts, _rewardVault common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "setRewardVault", _rewardVault)
}

// SetRewardVault is a paid mutator transaction binding the contract method 0x8125dd10.
//
// Solidity: function setRewardVault(address _rewardVault) returns()
func (_Erc4646Vault *Erc4646VaultSession) SetRewardVault(_rewardVault common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.SetRewardVault(&_Erc4646Vault.TransactOpts, _rewardVault)
}

// SetRewardVault is a paid mutator transaction binding the contract method 0x8125dd10.
//
// Solidity: function setRewardVault(address _rewardVault) returns()
func (_Erc4646Vault *Erc4646VaultTransactorSession) SetRewardVault(_rewardVault common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.SetRewardVault(&_Erc4646Vault.TransactOpts, _rewardVault)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Erc4646Vault *Erc4646VaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Erc4646Vault *Erc4646VaultSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Transfer(&_Erc4646Vault.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Erc4646Vault *Erc4646VaultTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Transfer(&_Erc4646Vault.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Erc4646Vault *Erc4646VaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Erc4646Vault *Erc4646VaultSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.TransferFrom(&_Erc4646Vault.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Erc4646Vault *Erc4646VaultTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.TransferFrom(&_Erc4646Vault.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc4646Vault *Erc4646VaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc4646Vault *Erc4646VaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.TransferOwnership(&_Erc4646Vault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc4646Vault *Erc4646VaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.TransferOwnership(&_Erc4646Vault.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Erc4646Vault *Erc4646VaultTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Erc4646Vault *Erc4646VaultSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.UpgradeToAndCall(&_Erc4646Vault.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Erc4646Vault *Erc4646VaultTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.UpgradeToAndCall(&_Erc4646Vault.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_Erc4646Vault *Erc4646VaultTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_Erc4646Vault *Erc4646VaultSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Withdraw(&_Erc4646Vault.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_Erc4646Vault *Erc4646VaultTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Erc4646Vault.Contract.Withdraw(&_Erc4646Vault.TransactOpts, assets, receiver, owner)
}

// Erc4646VaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc4646Vault contract.
type Erc4646VaultApprovalIterator struct {
	Event *Erc4646VaultApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultApproval represents a Approval event raised by the Erc4646Vault contract.
type Erc4646VaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc4646VaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultApprovalIterator{contract: _Erc4646Vault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc4646VaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultApproval)
				if err := _Erc4646Vault.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc4646Vault *Erc4646VaultFilterer) ParseApproval(log types.Log) (*Erc4646VaultApproval, error) {
	event := new(Erc4646VaultApproval)
	if err := _Erc4646Vault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4646VaultAssetAddedToEmissionListIterator is returned from FilterAssetAddedToEmissionList and is used to iterate over the raw logs and unpacked data for AssetAddedToEmissionList events raised by the Erc4646Vault contract.
type Erc4646VaultAssetAddedToEmissionListIterator struct {
	Event *Erc4646VaultAssetAddedToEmissionList // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultAssetAddedToEmissionListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultAssetAddedToEmissionList)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultAssetAddedToEmissionList)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultAssetAddedToEmissionListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultAssetAddedToEmissionListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultAssetAddedToEmissionList represents a AssetAddedToEmissionList event raised by the Erc4646Vault contract.
type Erc4646VaultAssetAddedToEmissionList struct {
	Asset common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetAddedToEmissionList is a free log retrieval operation binding the contract event 0xfe55b62d54eea531cb2ab0224ef4aa29eb159c1303f08e0b4e5811cabe810a30.
//
// Solidity: event AssetAddedToEmissionList(address indexed asset)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterAssetAddedToEmissionList(opts *bind.FilterOpts, asset []common.Address) (*Erc4646VaultAssetAddedToEmissionListIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "AssetAddedToEmissionList", assetRule)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultAssetAddedToEmissionListIterator{contract: _Erc4646Vault.contract, event: "AssetAddedToEmissionList", logs: logs, sub: sub}, nil
}

// WatchAssetAddedToEmissionList is a free log subscription operation binding the contract event 0xfe55b62d54eea531cb2ab0224ef4aa29eb159c1303f08e0b4e5811cabe810a30.
//
// Solidity: event AssetAddedToEmissionList(address indexed asset)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchAssetAddedToEmissionList(opts *bind.WatchOpts, sink chan<- *Erc4646VaultAssetAddedToEmissionList, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "AssetAddedToEmissionList", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultAssetAddedToEmissionList)
				if err := _Erc4646Vault.contract.UnpackLog(event, "AssetAddedToEmissionList", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetAddedToEmissionList is a log parse operation binding the contract event 0xfe55b62d54eea531cb2ab0224ef4aa29eb159c1303f08e0b4e5811cabe810a30.
//
// Solidity: event AssetAddedToEmissionList(address indexed asset)
func (_Erc4646Vault *Erc4646VaultFilterer) ParseAssetAddedToEmissionList(log types.Log) (*Erc4646VaultAssetAddedToEmissionList, error) {
	event := new(Erc4646VaultAssetAddedToEmissionList)
	if err := _Erc4646Vault.contract.UnpackLog(event, "AssetAddedToEmissionList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4646VaultAssetRemovedFromEmissionListIterator is returned from FilterAssetRemovedFromEmissionList and is used to iterate over the raw logs and unpacked data for AssetRemovedFromEmissionList events raised by the Erc4646Vault contract.
type Erc4646VaultAssetRemovedFromEmissionListIterator struct {
	Event *Erc4646VaultAssetRemovedFromEmissionList // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultAssetRemovedFromEmissionListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultAssetRemovedFromEmissionList)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultAssetRemovedFromEmissionList)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultAssetRemovedFromEmissionListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultAssetRemovedFromEmissionListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultAssetRemovedFromEmissionList represents a AssetRemovedFromEmissionList event raised by the Erc4646Vault contract.
type Erc4646VaultAssetRemovedFromEmissionList struct {
	Asset common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetRemovedFromEmissionList is a free log retrieval operation binding the contract event 0xf317de7b5d0ba757a221e919e6249a2e16d542c2f90b59d64d50484be4be3bd7.
//
// Solidity: event AssetRemovedFromEmissionList(address indexed asset)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterAssetRemovedFromEmissionList(opts *bind.FilterOpts, asset []common.Address) (*Erc4646VaultAssetRemovedFromEmissionListIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "AssetRemovedFromEmissionList", assetRule)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultAssetRemovedFromEmissionListIterator{contract: _Erc4646Vault.contract, event: "AssetRemovedFromEmissionList", logs: logs, sub: sub}, nil
}

// WatchAssetRemovedFromEmissionList is a free log subscription operation binding the contract event 0xf317de7b5d0ba757a221e919e6249a2e16d542c2f90b59d64d50484be4be3bd7.
//
// Solidity: event AssetRemovedFromEmissionList(address indexed asset)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchAssetRemovedFromEmissionList(opts *bind.WatchOpts, sink chan<- *Erc4646VaultAssetRemovedFromEmissionList, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "AssetRemovedFromEmissionList", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultAssetRemovedFromEmissionList)
				if err := _Erc4646Vault.contract.UnpackLog(event, "AssetRemovedFromEmissionList", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetRemovedFromEmissionList is a log parse operation binding the contract event 0xf317de7b5d0ba757a221e919e6249a2e16d542c2f90b59d64d50484be4be3bd7.
//
// Solidity: event AssetRemovedFromEmissionList(address indexed asset)
func (_Erc4646Vault *Erc4646VaultFilterer) ParseAssetRemovedFromEmissionList(log types.Log) (*Erc4646VaultAssetRemovedFromEmissionList, error) {
	event := new(Erc4646VaultAssetRemovedFromEmissionList)
	if err := _Erc4646Vault.contract.UnpackLog(event, "AssetRemovedFromEmissionList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4646VaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Erc4646Vault contract.
type Erc4646VaultDepositIterator struct {
	Event *Erc4646VaultDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultDeposit represents a Deposit event raised by the Erc4646Vault contract.
type Erc4646VaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*Erc4646VaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultDepositIterator{contract: _Erc4646Vault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *Erc4646VaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultDeposit)
				if err := _Erc4646Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Erc4646Vault *Erc4646VaultFilterer) ParseDeposit(log types.Log) (*Erc4646VaultDeposit, error) {
	event := new(Erc4646VaultDeposit)
	if err := _Erc4646Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4646VaultInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Erc4646Vault contract.
type Erc4646VaultInitializedIterator struct {
	Event *Erc4646VaultInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultInitialized represents a Initialized event raised by the Erc4646Vault contract.
type Erc4646VaultInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*Erc4646VaultInitializedIterator, error) {

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultInitializedIterator{contract: _Erc4646Vault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Erc4646VaultInitialized) (event.Subscription, error) {

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultInitialized)
				if err := _Erc4646Vault.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Erc4646Vault *Erc4646VaultFilterer) ParseInitialized(log types.Log) (*Erc4646VaultInitialized, error) {
	event := new(Erc4646VaultInitialized)
	if err := _Erc4646Vault.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4646VaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Erc4646Vault contract.
type Erc4646VaultOwnershipTransferredIterator struct {
	Event *Erc4646VaultOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultOwnershipTransferred represents a OwnershipTransferred event raised by the Erc4646Vault contract.
type Erc4646VaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Erc4646VaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultOwnershipTransferredIterator{contract: _Erc4646Vault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Erc4646VaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultOwnershipTransferred)
				if err := _Erc4646Vault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc4646Vault *Erc4646VaultFilterer) ParseOwnershipTransferred(log types.Log) (*Erc4646VaultOwnershipTransferred, error) {
	event := new(Erc4646VaultOwnershipTransferred)
	if err := _Erc4646Vault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4646VaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Erc4646Vault contract.
type Erc4646VaultPausedIterator struct {
	Event *Erc4646VaultPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultPaused represents a Paused event raised by the Erc4646Vault contract.
type Erc4646VaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterPaused(opts *bind.FilterOpts) (*Erc4646VaultPausedIterator, error) {

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultPausedIterator{contract: _Erc4646Vault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *Erc4646VaultPaused) (event.Subscription, error) {

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultPaused)
				if err := _Erc4646Vault.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Erc4646Vault *Erc4646VaultFilterer) ParsePaused(log types.Log) (*Erc4646VaultPaused, error) {
	event := new(Erc4646VaultPaused)
	if err := _Erc4646Vault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4646VaultRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the Erc4646Vault contract.
type Erc4646VaultRewardClaimedIterator struct {
	Event *Erc4646VaultRewardClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultRewardClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultRewardClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultRewardClaimed represents a RewardClaimed event raised by the Erc4646Vault contract.
type Erc4646VaultRewardClaimed struct {
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed token, address indexed receiver, uint256 amount)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterRewardClaimed(opts *bind.FilterOpts, token []common.Address, receiver []common.Address) (*Erc4646VaultRewardClaimedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "RewardClaimed", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultRewardClaimedIterator{contract: _Erc4646Vault.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed token, address indexed receiver, uint256 amount)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *Erc4646VaultRewardClaimed, token []common.Address, receiver []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "RewardClaimed", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultRewardClaimed)
				if err := _Erc4646Vault.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardClaimed is a log parse operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed token, address indexed receiver, uint256 amount)
func (_Erc4646Vault *Erc4646VaultFilterer) ParseRewardClaimed(log types.Log) (*Erc4646VaultRewardClaimed, error) {
	event := new(Erc4646VaultRewardClaimed)
	if err := _Erc4646Vault.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4646VaultRewardIntegralForAccountUpdatedIterator is returned from FilterRewardIntegralForAccountUpdated and is used to iterate over the raw logs and unpacked data for RewardIntegralForAccountUpdated events raised by the Erc4646Vault contract.
type Erc4646VaultRewardIntegralForAccountUpdatedIterator struct {
	Event *Erc4646VaultRewardIntegralForAccountUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultRewardIntegralForAccountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultRewardIntegralForAccountUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultRewardIntegralForAccountUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultRewardIntegralForAccountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultRewardIntegralForAccountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultRewardIntegralForAccountUpdated represents a RewardIntegralForAccountUpdated event raised by the Erc4646Vault contract.
type Erc4646VaultRewardIntegralForAccountUpdated struct {
	Account        common.Address
	Token          common.Address
	PendingReward  *big.Int
	RewardIntegral *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardIntegralForAccountUpdated is a free log retrieval operation binding the contract event 0xdb1b81b8ffba907fa76754265baca950292a08fdd6584527e35204dea596dcc1.
//
// Solidity: event RewardIntegralForAccountUpdated(address indexed account, address indexed token, uint256 pendingReward, uint256 rewardIntegral)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterRewardIntegralForAccountUpdated(opts *bind.FilterOpts, account []common.Address, token []common.Address) (*Erc4646VaultRewardIntegralForAccountUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "RewardIntegralForAccountUpdated", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultRewardIntegralForAccountUpdatedIterator{contract: _Erc4646Vault.contract, event: "RewardIntegralForAccountUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardIntegralForAccountUpdated is a free log subscription operation binding the contract event 0xdb1b81b8ffba907fa76754265baca950292a08fdd6584527e35204dea596dcc1.
//
// Solidity: event RewardIntegralForAccountUpdated(address indexed account, address indexed token, uint256 pendingReward, uint256 rewardIntegral)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchRewardIntegralForAccountUpdated(opts *bind.WatchOpts, sink chan<- *Erc4646VaultRewardIntegralForAccountUpdated, account []common.Address, token []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "RewardIntegralForAccountUpdated", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultRewardIntegralForAccountUpdated)
				if err := _Erc4646Vault.contract.UnpackLog(event, "RewardIntegralForAccountUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardIntegralForAccountUpdated is a log parse operation binding the contract event 0xdb1b81b8ffba907fa76754265baca950292a08fdd6584527e35204dea596dcc1.
//
// Solidity: event RewardIntegralForAccountUpdated(address indexed account, address indexed token, uint256 pendingReward, uint256 rewardIntegral)
func (_Erc4646Vault *Erc4646VaultFilterer) ParseRewardIntegralForAccountUpdated(log types.Log) (*Erc4646VaultRewardIntegralForAccountUpdated, error) {
	event := new(Erc4646VaultRewardIntegralForAccountUpdated)
	if err := _Erc4646Vault.contract.UnpackLog(event, "RewardIntegralForAccountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4646VaultRewardIntegralUpdatedIterator is returned from FilterRewardIntegralUpdated and is used to iterate over the raw logs and unpacked data for RewardIntegralUpdated events raised by the Erc4646Vault contract.
type Erc4646VaultRewardIntegralUpdatedIterator struct {
	Event *Erc4646VaultRewardIntegralUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultRewardIntegralUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultRewardIntegralUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultRewardIntegralUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultRewardIntegralUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultRewardIntegralUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultRewardIntegralUpdated represents a RewardIntegralUpdated event raised by the Erc4646Vault contract.
type Erc4646VaultRewardIntegralUpdated struct {
	Token          common.Address
	RewardIntegral *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardIntegralUpdated is a free log retrieval operation binding the contract event 0x61d927f91635642c4c52b78d3f339c15ff959118da79ceb10122a91902182b04.
//
// Solidity: event RewardIntegralUpdated(address indexed token, uint256 rewardIntegral)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterRewardIntegralUpdated(opts *bind.FilterOpts, token []common.Address) (*Erc4646VaultRewardIntegralUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "RewardIntegralUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultRewardIntegralUpdatedIterator{contract: _Erc4646Vault.contract, event: "RewardIntegralUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardIntegralUpdated is a free log subscription operation binding the contract event 0x61d927f91635642c4c52b78d3f339c15ff959118da79ceb10122a91902182b04.
//
// Solidity: event RewardIntegralUpdated(address indexed token, uint256 rewardIntegral)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchRewardIntegralUpdated(opts *bind.WatchOpts, sink chan<- *Erc4646VaultRewardIntegralUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "RewardIntegralUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultRewardIntegralUpdated)
				if err := _Erc4646Vault.contract.UnpackLog(event, "RewardIntegralUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardIntegralUpdated is a log parse operation binding the contract event 0x61d927f91635642c4c52b78d3f339c15ff959118da79ceb10122a91902182b04.
//
// Solidity: event RewardIntegralUpdated(address indexed token, uint256 rewardIntegral)
func (_Erc4646Vault *Erc4646VaultFilterer) ParseRewardIntegralUpdated(log types.Log) (*Erc4646VaultRewardIntegralUpdated, error) {
	event := new(Erc4646VaultRewardIntegralUpdated)
	if err := _Erc4646Vault.contract.UnpackLog(event, "RewardIntegralUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4646VaultRewardParamsSetIterator is returned from FilterRewardParamsSet and is used to iterate over the raw logs and unpacked data for RewardParamsSet events raised by the Erc4646Vault contract.
type Erc4646VaultRewardParamsSetIterator struct {
	Event *Erc4646VaultRewardParamsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultRewardParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultRewardParamsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultRewardParamsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultRewardParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultRewardParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultRewardParamsSet represents a RewardParamsSet event raised by the Erc4646Vault contract.
type Erc4646VaultRewardParamsSet struct {
	Token          common.Address
	RewardRate     *big.Int
	RewardIntegral *big.Int
	ClaimStartTime uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardParamsSet is a free log retrieval operation binding the contract event 0xa82b7ef40535dbf27232e95dedda07f0ccf028e6f5432c9f79b843d5ed11d0f5.
//
// Solidity: event RewardParamsSet(address indexed token, uint256 rewardRate, uint256 rewardIntegral, uint32 claimStartTime)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterRewardParamsSet(opts *bind.FilterOpts, token []common.Address) (*Erc4646VaultRewardParamsSetIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "RewardParamsSet", tokenRule)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultRewardParamsSetIterator{contract: _Erc4646Vault.contract, event: "RewardParamsSet", logs: logs, sub: sub}, nil
}

// WatchRewardParamsSet is a free log subscription operation binding the contract event 0xa82b7ef40535dbf27232e95dedda07f0ccf028e6f5432c9f79b843d5ed11d0f5.
//
// Solidity: event RewardParamsSet(address indexed token, uint256 rewardRate, uint256 rewardIntegral, uint32 claimStartTime)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchRewardParamsSet(opts *bind.WatchOpts, sink chan<- *Erc4646VaultRewardParamsSet, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "RewardParamsSet", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultRewardParamsSet)
				if err := _Erc4646Vault.contract.UnpackLog(event, "RewardParamsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardParamsSet is a log parse operation binding the contract event 0xa82b7ef40535dbf27232e95dedda07f0ccf028e6f5432c9f79b843d5ed11d0f5.
//
// Solidity: event RewardParamsSet(address indexed token, uint256 rewardRate, uint256 rewardIntegral, uint32 claimStartTime)
func (_Erc4646Vault *Erc4646VaultFilterer) ParseRewardParamsSet(log types.Log) (*Erc4646VaultRewardParamsSet, error) {
	event := new(Erc4646VaultRewardParamsSet)
	if err := _Erc4646Vault.contract.UnpackLog(event, "RewardParamsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4646VaultRewardVaultSettingIterator is returned from FilterRewardVaultSetting and is used to iterate over the raw logs and unpacked data for RewardVaultSetting events raised by the Erc4646Vault contract.
type Erc4646VaultRewardVaultSettingIterator struct {
	Event *Erc4646VaultRewardVaultSetting // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultRewardVaultSettingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultRewardVaultSetting)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultRewardVaultSetting)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultRewardVaultSettingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultRewardVaultSettingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultRewardVaultSetting represents a RewardVaultSetting event raised by the Erc4646Vault contract.
type Erc4646VaultRewardVaultSetting struct {
	RewardVault common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardVaultSetting is a free log retrieval operation binding the contract event 0x7ccb2cd577263a3595bff3e18ab7719c1f427c7fd48cdd043eff96fd3c583753.
//
// Solidity: event RewardVaultSetting(address indexed rewardVault)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterRewardVaultSetting(opts *bind.FilterOpts, rewardVault []common.Address) (*Erc4646VaultRewardVaultSettingIterator, error) {

	var rewardVaultRule []interface{}
	for _, rewardVaultItem := range rewardVault {
		rewardVaultRule = append(rewardVaultRule, rewardVaultItem)
	}

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "RewardVaultSetting", rewardVaultRule)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultRewardVaultSettingIterator{contract: _Erc4646Vault.contract, event: "RewardVaultSetting", logs: logs, sub: sub}, nil
}

// WatchRewardVaultSetting is a free log subscription operation binding the contract event 0x7ccb2cd577263a3595bff3e18ab7719c1f427c7fd48cdd043eff96fd3c583753.
//
// Solidity: event RewardVaultSetting(address indexed rewardVault)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchRewardVaultSetting(opts *bind.WatchOpts, sink chan<- *Erc4646VaultRewardVaultSetting, rewardVault []common.Address) (event.Subscription, error) {

	var rewardVaultRule []interface{}
	for _, rewardVaultItem := range rewardVault {
		rewardVaultRule = append(rewardVaultRule, rewardVaultItem)
	}

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "RewardVaultSetting", rewardVaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultRewardVaultSetting)
				if err := _Erc4646Vault.contract.UnpackLog(event, "RewardVaultSetting", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardVaultSetting is a log parse operation binding the contract event 0x7ccb2cd577263a3595bff3e18ab7719c1f427c7fd48cdd043eff96fd3c583753.
//
// Solidity: event RewardVaultSetting(address indexed rewardVault)
func (_Erc4646Vault *Erc4646VaultFilterer) ParseRewardVaultSetting(log types.Log) (*Erc4646VaultRewardVaultSetting, error) {
	event := new(Erc4646VaultRewardVaultSetting)
	if err := _Erc4646Vault.contract.UnpackLog(event, "RewardVaultSetting", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4646VaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc4646Vault contract.
type Erc4646VaultTransferIterator struct {
	Event *Erc4646VaultTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultTransfer represents a Transfer event raised by the Erc4646Vault contract.
type Erc4646VaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc4646VaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultTransferIterator{contract: _Erc4646Vault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc4646VaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultTransfer)
				if err := _Erc4646Vault.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc4646Vault *Erc4646VaultFilterer) ParseTransfer(log types.Log) (*Erc4646VaultTransfer, error) {
	event := new(Erc4646VaultTransfer)
	if err := _Erc4646Vault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4646VaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Erc4646Vault contract.
type Erc4646VaultUnpausedIterator struct {
	Event *Erc4646VaultUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultUnpaused represents a Unpaused event raised by the Erc4646Vault contract.
type Erc4646VaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*Erc4646VaultUnpausedIterator, error) {

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultUnpausedIterator{contract: _Erc4646Vault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Erc4646VaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultUnpaused)
				if err := _Erc4646Vault.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Erc4646Vault *Erc4646VaultFilterer) ParseUnpaused(log types.Log) (*Erc4646VaultUnpaused, error) {
	event := new(Erc4646VaultUnpaused)
	if err := _Erc4646Vault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4646VaultUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Erc4646Vault contract.
type Erc4646VaultUpgradedIterator struct {
	Event *Erc4646VaultUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultUpgraded represents a Upgraded event raised by the Erc4646Vault contract.
type Erc4646VaultUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*Erc4646VaultUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultUpgradedIterator{contract: _Erc4646Vault.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *Erc4646VaultUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultUpgraded)
				if err := _Erc4646Vault.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Erc4646Vault *Erc4646VaultFilterer) ParseUpgraded(log types.Log) (*Erc4646VaultUpgraded, error) {
	event := new(Erc4646VaultUpgraded)
	if err := _Erc4646Vault.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4646VaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Erc4646Vault contract.
type Erc4646VaultWithdrawIterator struct {
	Event *Erc4646VaultWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc4646VaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4646VaultWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc4646VaultWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc4646VaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4646VaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4646VaultWithdraw represents a Withdraw event raised by the Erc4646Vault contract.
type Erc4646VaultWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Erc4646Vault *Erc4646VaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*Erc4646VaultWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Erc4646Vault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &Erc4646VaultWithdrawIterator{contract: _Erc4646Vault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Erc4646Vault *Erc4646VaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Erc4646VaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Erc4646Vault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4646VaultWithdraw)
				if err := _Erc4646Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Erc4646Vault *Erc4646VaultFilterer) ParseWithdraw(log types.Log) (*Erc4646VaultWithdraw, error) {
	event := new(Erc4646VaultWithdraw)
	if err := _Erc4646Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
