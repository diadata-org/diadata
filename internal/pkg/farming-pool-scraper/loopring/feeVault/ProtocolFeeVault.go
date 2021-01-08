// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package protocolfeevault

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ProtocolFeeVaultABI is the input ABI used to generate the binding from.
const ProtocolFeeVaultABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lrcAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountDAO\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountBurn\",\"type\":\"uint256\"}],\"name\":\"DAOFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LRCClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"SettingsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenSold\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"DAO_PERDENTAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REWARD_PERCENTAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimStakingReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"daoAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fundDAO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProtocolFeeStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accumulatedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedBurn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedDAOFund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingBurn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingDAOFund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingReward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lrcAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sellTokenForLRC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenSellerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userStakingPoolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenSellerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_daoAddress\",\"type\":\"address\"}],\"name\":\"updateSettings\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"userStakingPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ProtocolFeeVault is an auto generated Go binding around an Ethereum contract.
type ProtocolFeeVault struct {
	ProtocolFeeVaultCaller     // Read-only binding to the contract
	ProtocolFeeVaultTransactor // Write-only binding to the contract
	ProtocolFeeVaultFilterer   // Log filterer for contract events
}

// ProtocolFeeVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtocolFeeVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolFeeVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtocolFeeVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolFeeVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtocolFeeVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolFeeVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtocolFeeVaultSession struct {
	Contract     *ProtocolFeeVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtocolFeeVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtocolFeeVaultCallerSession struct {
	Contract *ProtocolFeeVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ProtocolFeeVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtocolFeeVaultTransactorSession struct {
	Contract     *ProtocolFeeVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ProtocolFeeVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtocolFeeVaultRaw struct {
	Contract *ProtocolFeeVault // Generic contract binding to access the raw methods on
}

// ProtocolFeeVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtocolFeeVaultCallerRaw struct {
	Contract *ProtocolFeeVaultCaller // Generic read-only contract binding to access the raw methods on
}

// ProtocolFeeVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtocolFeeVaultTransactorRaw struct {
	Contract *ProtocolFeeVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocolFeeVault creates a new instance of ProtocolFeeVault, bound to a specific deployed contract.
func NewProtocolFeeVault(address common.Address, backend bind.ContractBackend) (*ProtocolFeeVault, error) {
	contract, err := bindProtocolFeeVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProtocolFeeVault{ProtocolFeeVaultCaller: ProtocolFeeVaultCaller{contract: contract}, ProtocolFeeVaultTransactor: ProtocolFeeVaultTransactor{contract: contract}, ProtocolFeeVaultFilterer: ProtocolFeeVaultFilterer{contract: contract}}, nil
}

// NewProtocolFeeVaultCaller creates a new read-only instance of ProtocolFeeVault, bound to a specific deployed contract.
func NewProtocolFeeVaultCaller(address common.Address, caller bind.ContractCaller) (*ProtocolFeeVaultCaller, error) {
	contract, err := bindProtocolFeeVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolFeeVaultCaller{contract: contract}, nil
}

// NewProtocolFeeVaultTransactor creates a new write-only instance of ProtocolFeeVault, bound to a specific deployed contract.
func NewProtocolFeeVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtocolFeeVaultTransactor, error) {
	contract, err := bindProtocolFeeVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolFeeVaultTransactor{contract: contract}, nil
}

// NewProtocolFeeVaultFilterer creates a new log filterer instance of ProtocolFeeVault, bound to a specific deployed contract.
func NewProtocolFeeVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtocolFeeVaultFilterer, error) {
	contract, err := bindProtocolFeeVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtocolFeeVaultFilterer{contract: contract}, nil
}

// bindProtocolFeeVault binds a generic wrapper to an already deployed contract.
func bindProtocolFeeVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolFeeVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolFeeVault *ProtocolFeeVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtocolFeeVault.Contract.ProtocolFeeVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolFeeVault *ProtocolFeeVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.ProtocolFeeVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolFeeVault *ProtocolFeeVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.ProtocolFeeVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolFeeVault *ProtocolFeeVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtocolFeeVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolFeeVault *ProtocolFeeVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolFeeVault *ProtocolFeeVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.contract.Transact(opts, method, params...)
}

// DAOPERDENTAGE is a free data retrieval call binding the contract method 0x68b79c32.
//
// Solidity: function DAO_PERDENTAGE() view returns(uint256)
func (_ProtocolFeeVault *ProtocolFeeVaultCaller) DAOPERDENTAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolFeeVault.contract.Call(opts, &out, "DAO_PERDENTAGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DAOPERDENTAGE is a free data retrieval call binding the contract method 0x68b79c32.
//
// Solidity: function DAO_PERDENTAGE() view returns(uint256)
func (_ProtocolFeeVault *ProtocolFeeVaultSession) DAOPERDENTAGE() (*big.Int, error) {
	return _ProtocolFeeVault.Contract.DAOPERDENTAGE(&_ProtocolFeeVault.CallOpts)
}

// DAOPERDENTAGE is a free data retrieval call binding the contract method 0x68b79c32.
//
// Solidity: function DAO_PERDENTAGE() view returns(uint256)
func (_ProtocolFeeVault *ProtocolFeeVaultCallerSession) DAOPERDENTAGE() (*big.Int, error) {
	return _ProtocolFeeVault.Contract.DAOPERDENTAGE(&_ProtocolFeeVault.CallOpts)
}

// REWARDPERCENTAGE is a free data retrieval call binding the contract method 0xfc72b1ed.
//
// Solidity: function REWARD_PERCENTAGE() view returns(uint256)
func (_ProtocolFeeVault *ProtocolFeeVaultCaller) REWARDPERCENTAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolFeeVault.contract.Call(opts, &out, "REWARD_PERCENTAGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDPERCENTAGE is a free data retrieval call binding the contract method 0xfc72b1ed.
//
// Solidity: function REWARD_PERCENTAGE() view returns(uint256)
func (_ProtocolFeeVault *ProtocolFeeVaultSession) REWARDPERCENTAGE() (*big.Int, error) {
	return _ProtocolFeeVault.Contract.REWARDPERCENTAGE(&_ProtocolFeeVault.CallOpts)
}

// REWARDPERCENTAGE is a free data retrieval call binding the contract method 0xfc72b1ed.
//
// Solidity: function REWARD_PERCENTAGE() view returns(uint256)
func (_ProtocolFeeVault *ProtocolFeeVaultCallerSession) REWARDPERCENTAGE() (*big.Int, error) {
	return _ProtocolFeeVault.Contract.REWARDPERCENTAGE(&_ProtocolFeeVault.CallOpts)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultCaller) DaoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProtocolFeeVault.contract.Call(opts, &out, "daoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultSession) DaoAddress() (common.Address, error) {
	return _ProtocolFeeVault.Contract.DaoAddress(&_ProtocolFeeVault.CallOpts)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultCallerSession) DaoAddress() (common.Address, error) {
	return _ProtocolFeeVault.Contract.DaoAddress(&_ProtocolFeeVault.CallOpts)
}

// GetProtocolFeeStats is a free data retrieval call binding the contract method 0x365a906c.
//
// Solidity: function getProtocolFeeStats() view returns(uint256 accumulatedFees, uint256 accumulatedBurn, uint256 accumulatedDAOFund, uint256 accumulatedReward, uint256 remainingFees, uint256 remainingBurn, uint256 remainingDAOFund, uint256 remainingReward)
func (_ProtocolFeeVault *ProtocolFeeVaultCaller) GetProtocolFeeStats(opts *bind.CallOpts) (struct {
	AccumulatedFees    *big.Int
	AccumulatedBurn    *big.Int
	AccumulatedDAOFund *big.Int
	AccumulatedReward  *big.Int
	RemainingFees      *big.Int
	RemainingBurn      *big.Int
	RemainingDAOFund   *big.Int
	RemainingReward    *big.Int
}, error) {
	var out []interface{}
	err := _ProtocolFeeVault.contract.Call(opts, &out, "getProtocolFeeStats")

	outstruct := new(struct {
		AccumulatedFees    *big.Int
		AccumulatedBurn    *big.Int
		AccumulatedDAOFund *big.Int
		AccumulatedReward  *big.Int
		RemainingFees      *big.Int
		RemainingBurn      *big.Int
		RemainingDAOFund   *big.Int
		RemainingReward    *big.Int
	})

	outstruct.AccumulatedFees = out[0].(*big.Int)
	outstruct.AccumulatedBurn = out[1].(*big.Int)
	outstruct.AccumulatedDAOFund = out[2].(*big.Int)
	outstruct.AccumulatedReward = out[3].(*big.Int)
	outstruct.RemainingFees = out[4].(*big.Int)
	outstruct.RemainingBurn = out[5].(*big.Int)
	outstruct.RemainingDAOFund = out[6].(*big.Int)
	outstruct.RemainingReward = out[7].(*big.Int)

	return *outstruct, err

}

// GetProtocolFeeStats is a free data retrieval call binding the contract method 0x365a906c.
//
// Solidity: function getProtocolFeeStats() view returns(uint256 accumulatedFees, uint256 accumulatedBurn, uint256 accumulatedDAOFund, uint256 accumulatedReward, uint256 remainingFees, uint256 remainingBurn, uint256 remainingDAOFund, uint256 remainingReward)
func (_ProtocolFeeVault *ProtocolFeeVaultSession) GetProtocolFeeStats() (struct {
	AccumulatedFees    *big.Int
	AccumulatedBurn    *big.Int
	AccumulatedDAOFund *big.Int
	AccumulatedReward  *big.Int
	RemainingFees      *big.Int
	RemainingBurn      *big.Int
	RemainingDAOFund   *big.Int
	RemainingReward    *big.Int
}, error) {
	return _ProtocolFeeVault.Contract.GetProtocolFeeStats(&_ProtocolFeeVault.CallOpts)
}

// GetProtocolFeeStats is a free data retrieval call binding the contract method 0x365a906c.
//
// Solidity: function getProtocolFeeStats() view returns(uint256 accumulatedFees, uint256 accumulatedBurn, uint256 accumulatedDAOFund, uint256 accumulatedReward, uint256 remainingFees, uint256 remainingBurn, uint256 remainingDAOFund, uint256 remainingReward)
func (_ProtocolFeeVault *ProtocolFeeVaultCallerSession) GetProtocolFeeStats() (struct {
	AccumulatedFees    *big.Int
	AccumulatedBurn    *big.Int
	AccumulatedDAOFund *big.Int
	AccumulatedReward  *big.Int
	RemainingFees      *big.Int
	RemainingBurn      *big.Int
	RemainingDAOFund   *big.Int
	RemainingReward    *big.Int
}, error) {
	return _ProtocolFeeVault.Contract.GetProtocolFeeStats(&_ProtocolFeeVault.CallOpts)
}

// LrcAddress is a free data retrieval call binding the contract method 0x3d6cf722.
//
// Solidity: function lrcAddress() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultCaller) LrcAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProtocolFeeVault.contract.Call(opts, &out, "lrcAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LrcAddress is a free data retrieval call binding the contract method 0x3d6cf722.
//
// Solidity: function lrcAddress() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultSession) LrcAddress() (common.Address, error) {
	return _ProtocolFeeVault.Contract.LrcAddress(&_ProtocolFeeVault.CallOpts)
}

// LrcAddress is a free data retrieval call binding the contract method 0x3d6cf722.
//
// Solidity: function lrcAddress() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultCallerSession) LrcAddress() (common.Address, error) {
	return _ProtocolFeeVault.Contract.LrcAddress(&_ProtocolFeeVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProtocolFeeVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultSession) Owner() (common.Address, error) {
	return _ProtocolFeeVault.Contract.Owner(&_ProtocolFeeVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultCallerSession) Owner() (common.Address, error) {
	return _ProtocolFeeVault.Contract.Owner(&_ProtocolFeeVault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProtocolFeeVault.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultSession) PendingOwner() (common.Address, error) {
	return _ProtocolFeeVault.Contract.PendingOwner(&_ProtocolFeeVault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultCallerSession) PendingOwner() (common.Address, error) {
	return _ProtocolFeeVault.Contract.PendingOwner(&_ProtocolFeeVault.CallOpts)
}

// TokenSellerAddress is a free data retrieval call binding the contract method 0xe3fe4764.
//
// Solidity: function tokenSellerAddress() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultCaller) TokenSellerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProtocolFeeVault.contract.Call(opts, &out, "tokenSellerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenSellerAddress is a free data retrieval call binding the contract method 0xe3fe4764.
//
// Solidity: function tokenSellerAddress() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultSession) TokenSellerAddress() (common.Address, error) {
	return _ProtocolFeeVault.Contract.TokenSellerAddress(&_ProtocolFeeVault.CallOpts)
}

// TokenSellerAddress is a free data retrieval call binding the contract method 0xe3fe4764.
//
// Solidity: function tokenSellerAddress() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultCallerSession) TokenSellerAddress() (common.Address, error) {
	return _ProtocolFeeVault.Contract.TokenSellerAddress(&_ProtocolFeeVault.CallOpts)
}

// UserStakingPoolAddress is a free data retrieval call binding the contract method 0xa0cea6c0.
//
// Solidity: function userStakingPoolAddress() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultCaller) UserStakingPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProtocolFeeVault.contract.Call(opts, &out, "userStakingPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserStakingPoolAddress is a free data retrieval call binding the contract method 0xa0cea6c0.
//
// Solidity: function userStakingPoolAddress() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultSession) UserStakingPoolAddress() (common.Address, error) {
	return _ProtocolFeeVault.Contract.UserStakingPoolAddress(&_ProtocolFeeVault.CallOpts)
}

// UserStakingPoolAddress is a free data retrieval call binding the contract method 0xa0cea6c0.
//
// Solidity: function userStakingPoolAddress() view returns(address)
func (_ProtocolFeeVault *ProtocolFeeVaultCallerSession) UserStakingPoolAddress() (common.Address, error) {
	return _ProtocolFeeVault.Contract.UserStakingPoolAddress(&_ProtocolFeeVault.CallOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolFeeVault.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_ProtocolFeeVault *ProtocolFeeVaultSession) ClaimOwnership() (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.ClaimOwnership(&_ProtocolFeeVault.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.ClaimOwnership(&_ProtocolFeeVault.TransactOpts)
}

// ClaimStakingReward is a paid mutator transaction binding the contract method 0x239cd4a4.
//
// Solidity: function claimStakingReward(uint256 amount) returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactor) ClaimStakingReward(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ProtocolFeeVault.contract.Transact(opts, "claimStakingReward", amount)
}

// ClaimStakingReward is a paid mutator transaction binding the contract method 0x239cd4a4.
//
// Solidity: function claimStakingReward(uint256 amount) returns()
func (_ProtocolFeeVault *ProtocolFeeVaultSession) ClaimStakingReward(amount *big.Int) (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.ClaimStakingReward(&_ProtocolFeeVault.TransactOpts, amount)
}

// ClaimStakingReward is a paid mutator transaction binding the contract method 0x239cd4a4.
//
// Solidity: function claimStakingReward(uint256 amount) returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactorSession) ClaimStakingReward(amount *big.Int) (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.ClaimStakingReward(&_ProtocolFeeVault.TransactOpts, amount)
}

// FundDAO is a paid mutator transaction binding the contract method 0x61163379.
//
// Solidity: function fundDAO() returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactor) FundDAO(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolFeeVault.contract.Transact(opts, "fundDAO")
}

// FundDAO is a paid mutator transaction binding the contract method 0x61163379.
//
// Solidity: function fundDAO() returns()
func (_ProtocolFeeVault *ProtocolFeeVaultSession) FundDAO() (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.FundDAO(&_ProtocolFeeVault.TransactOpts)
}

// FundDAO is a paid mutator transaction binding the contract method 0x61163379.
//
// Solidity: function fundDAO() returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactorSession) FundDAO() (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.FundDAO(&_ProtocolFeeVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolFeeVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProtocolFeeVault *ProtocolFeeVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.RenounceOwnership(&_ProtocolFeeVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.RenounceOwnership(&_ProtocolFeeVault.TransactOpts)
}

// SellTokenForLRC is a paid mutator transaction binding the contract method 0x2a7bf6bf.
//
// Solidity: function sellTokenForLRC(address token, uint256 amount) returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactor) SellTokenForLRC(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProtocolFeeVault.contract.Transact(opts, "sellTokenForLRC", token, amount)
}

// SellTokenForLRC is a paid mutator transaction binding the contract method 0x2a7bf6bf.
//
// Solidity: function sellTokenForLRC(address token, uint256 amount) returns()
func (_ProtocolFeeVault *ProtocolFeeVaultSession) SellTokenForLRC(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.SellTokenForLRC(&_ProtocolFeeVault.TransactOpts, token, amount)
}

// SellTokenForLRC is a paid mutator transaction binding the contract method 0x2a7bf6bf.
//
// Solidity: function sellTokenForLRC(address token, uint256 amount) returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactorSession) SellTokenForLRC(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.SellTokenForLRC(&_ProtocolFeeVault.TransactOpts, token, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProtocolFeeVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProtocolFeeVault *ProtocolFeeVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.TransferOwnership(&_ProtocolFeeVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.TransferOwnership(&_ProtocolFeeVault.TransactOpts, newOwner)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x57aabe91.
//
// Solidity: function updateSettings(address _userStakingPoolAddress, address _tokenSellerAddress, address _daoAddress) returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactor) UpdateSettings(opts *bind.TransactOpts, _userStakingPoolAddress common.Address, _tokenSellerAddress common.Address, _daoAddress common.Address) (*types.Transaction, error) {
	return _ProtocolFeeVault.contract.Transact(opts, "updateSettings", _userStakingPoolAddress, _tokenSellerAddress, _daoAddress)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x57aabe91.
//
// Solidity: function updateSettings(address _userStakingPoolAddress, address _tokenSellerAddress, address _daoAddress) returns()
func (_ProtocolFeeVault *ProtocolFeeVaultSession) UpdateSettings(_userStakingPoolAddress common.Address, _tokenSellerAddress common.Address, _daoAddress common.Address) (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.UpdateSettings(&_ProtocolFeeVault.TransactOpts, _userStakingPoolAddress, _tokenSellerAddress, _daoAddress)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x57aabe91.
//
// Solidity: function updateSettings(address _userStakingPoolAddress, address _tokenSellerAddress, address _daoAddress) returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactorSession) UpdateSettings(_userStakingPoolAddress common.Address, _tokenSellerAddress common.Address, _daoAddress common.Address) (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.UpdateSettings(&_ProtocolFeeVault.TransactOpts, _userStakingPoolAddress, _tokenSellerAddress, _daoAddress)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ProtocolFeeVault.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ProtocolFeeVault *ProtocolFeeVaultSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.Fallback(&_ProtocolFeeVault.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ProtocolFeeVault *ProtocolFeeVaultTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ProtocolFeeVault.Contract.Fallback(&_ProtocolFeeVault.TransactOpts, calldata)
}

// ProtocolFeeVaultDAOFundedIterator is returned from FilterDAOFunded and is used to iterate over the raw logs and unpacked data for DAOFunded events raised by the ProtocolFeeVault contract.
type ProtocolFeeVaultDAOFundedIterator struct {
	Event *ProtocolFeeVaultDAOFunded // Event containing the contract specifics and raw log

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
func (it *ProtocolFeeVaultDAOFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolFeeVaultDAOFunded)
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
		it.Event = new(ProtocolFeeVaultDAOFunded)
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
func (it *ProtocolFeeVaultDAOFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolFeeVaultDAOFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolFeeVaultDAOFunded represents a DAOFunded event raised by the ProtocolFeeVault contract.
type ProtocolFeeVaultDAOFunded struct {
	AmountDAO  *big.Int
	AmountBurn *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDAOFunded is a free log retrieval operation binding the contract event 0xf5deee3dde315bb684736ad09d95603ad754fe04704f5a2f0719216a09dac68d.
//
// Solidity: event DAOFunded(uint256 amountDAO, uint256 amountBurn)
func (_ProtocolFeeVault *ProtocolFeeVaultFilterer) FilterDAOFunded(opts *bind.FilterOpts) (*ProtocolFeeVaultDAOFundedIterator, error) {

	logs, sub, err := _ProtocolFeeVault.contract.FilterLogs(opts, "DAOFunded")
	if err != nil {
		return nil, err
	}
	return &ProtocolFeeVaultDAOFundedIterator{contract: _ProtocolFeeVault.contract, event: "DAOFunded", logs: logs, sub: sub}, nil
}

// WatchDAOFunded is a free log subscription operation binding the contract event 0xf5deee3dde315bb684736ad09d95603ad754fe04704f5a2f0719216a09dac68d.
//
// Solidity: event DAOFunded(uint256 amountDAO, uint256 amountBurn)
func (_ProtocolFeeVault *ProtocolFeeVaultFilterer) WatchDAOFunded(opts *bind.WatchOpts, sink chan<- *ProtocolFeeVaultDAOFunded) (event.Subscription, error) {

	logs, sub, err := _ProtocolFeeVault.contract.WatchLogs(opts, "DAOFunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolFeeVaultDAOFunded)
				if err := _ProtocolFeeVault.contract.UnpackLog(event, "DAOFunded", log); err != nil {
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

// ParseDAOFunded is a log parse operation binding the contract event 0xf5deee3dde315bb684736ad09d95603ad754fe04704f5a2f0719216a09dac68d.
//
// Solidity: event DAOFunded(uint256 amountDAO, uint256 amountBurn)
func (_ProtocolFeeVault *ProtocolFeeVaultFilterer) ParseDAOFunded(log types.Log) (*ProtocolFeeVaultDAOFunded, error) {
	event := new(ProtocolFeeVaultDAOFunded)
	if err := _ProtocolFeeVault.contract.UnpackLog(event, "DAOFunded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProtocolFeeVaultLRCClaimedIterator is returned from FilterLRCClaimed and is used to iterate over the raw logs and unpacked data for LRCClaimed events raised by the ProtocolFeeVault contract.
type ProtocolFeeVaultLRCClaimedIterator struct {
	Event *ProtocolFeeVaultLRCClaimed // Event containing the contract specifics and raw log

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
func (it *ProtocolFeeVaultLRCClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolFeeVaultLRCClaimed)
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
		it.Event = new(ProtocolFeeVaultLRCClaimed)
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
func (it *ProtocolFeeVaultLRCClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolFeeVaultLRCClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolFeeVaultLRCClaimed represents a LRCClaimed event raised by the ProtocolFeeVault contract.
type ProtocolFeeVaultLRCClaimed struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLRCClaimed is a free log retrieval operation binding the contract event 0xaf9f8e373f3b94eb4c733ac25dd927271dc24febef9ae82f0868502653fdb65c.
//
// Solidity: event LRCClaimed(uint256 amount)
func (_ProtocolFeeVault *ProtocolFeeVaultFilterer) FilterLRCClaimed(opts *bind.FilterOpts) (*ProtocolFeeVaultLRCClaimedIterator, error) {

	logs, sub, err := _ProtocolFeeVault.contract.FilterLogs(opts, "LRCClaimed")
	if err != nil {
		return nil, err
	}
	return &ProtocolFeeVaultLRCClaimedIterator{contract: _ProtocolFeeVault.contract, event: "LRCClaimed", logs: logs, sub: sub}, nil
}

// WatchLRCClaimed is a free log subscription operation binding the contract event 0xaf9f8e373f3b94eb4c733ac25dd927271dc24febef9ae82f0868502653fdb65c.
//
// Solidity: event LRCClaimed(uint256 amount)
func (_ProtocolFeeVault *ProtocolFeeVaultFilterer) WatchLRCClaimed(opts *bind.WatchOpts, sink chan<- *ProtocolFeeVaultLRCClaimed) (event.Subscription, error) {

	logs, sub, err := _ProtocolFeeVault.contract.WatchLogs(opts, "LRCClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolFeeVaultLRCClaimed)
				if err := _ProtocolFeeVault.contract.UnpackLog(event, "LRCClaimed", log); err != nil {
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

// ParseLRCClaimed is a log parse operation binding the contract event 0xaf9f8e373f3b94eb4c733ac25dd927271dc24febef9ae82f0868502653fdb65c.
//
// Solidity: event LRCClaimed(uint256 amount)
func (_ProtocolFeeVault *ProtocolFeeVaultFilterer) ParseLRCClaimed(log types.Log) (*ProtocolFeeVaultLRCClaimed, error) {
	event := new(ProtocolFeeVaultLRCClaimed)
	if err := _ProtocolFeeVault.contract.UnpackLog(event, "LRCClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProtocolFeeVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProtocolFeeVault contract.
type ProtocolFeeVaultOwnershipTransferredIterator struct {
	Event *ProtocolFeeVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProtocolFeeVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolFeeVaultOwnershipTransferred)
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
		it.Event = new(ProtocolFeeVaultOwnershipTransferred)
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
func (it *ProtocolFeeVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolFeeVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolFeeVaultOwnershipTransferred represents a OwnershipTransferred event raised by the ProtocolFeeVault contract.
type ProtocolFeeVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProtocolFeeVault *ProtocolFeeVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProtocolFeeVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProtocolFeeVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolFeeVaultOwnershipTransferredIterator{contract: _ProtocolFeeVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProtocolFeeVault *ProtocolFeeVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProtocolFeeVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProtocolFeeVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolFeeVaultOwnershipTransferred)
				if err := _ProtocolFeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ProtocolFeeVault *ProtocolFeeVaultFilterer) ParseOwnershipTransferred(log types.Log) (*ProtocolFeeVaultOwnershipTransferred, error) {
	event := new(ProtocolFeeVaultOwnershipTransferred)
	if err := _ProtocolFeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProtocolFeeVaultSettingsUpdatedIterator is returned from FilterSettingsUpdated and is used to iterate over the raw logs and unpacked data for SettingsUpdated events raised by the ProtocolFeeVault contract.
type ProtocolFeeVaultSettingsUpdatedIterator struct {
	Event *ProtocolFeeVaultSettingsUpdated // Event containing the contract specifics and raw log

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
func (it *ProtocolFeeVaultSettingsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolFeeVaultSettingsUpdated)
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
		it.Event = new(ProtocolFeeVaultSettingsUpdated)
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
func (it *ProtocolFeeVaultSettingsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolFeeVaultSettingsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolFeeVaultSettingsUpdated represents a SettingsUpdated event raised by the ProtocolFeeVault contract.
type ProtocolFeeVaultSettingsUpdated struct {
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSettingsUpdated is a free log retrieval operation binding the contract event 0x4b804a0bfbdc2639203b93035be561d86f65f52b7e14984f95ec9d298cafccac.
//
// Solidity: event SettingsUpdated(uint256 time)
func (_ProtocolFeeVault *ProtocolFeeVaultFilterer) FilterSettingsUpdated(opts *bind.FilterOpts) (*ProtocolFeeVaultSettingsUpdatedIterator, error) {

	logs, sub, err := _ProtocolFeeVault.contract.FilterLogs(opts, "SettingsUpdated")
	if err != nil {
		return nil, err
	}
	return &ProtocolFeeVaultSettingsUpdatedIterator{contract: _ProtocolFeeVault.contract, event: "SettingsUpdated", logs: logs, sub: sub}, nil
}

// WatchSettingsUpdated is a free log subscription operation binding the contract event 0x4b804a0bfbdc2639203b93035be561d86f65f52b7e14984f95ec9d298cafccac.
//
// Solidity: event SettingsUpdated(uint256 time)
func (_ProtocolFeeVault *ProtocolFeeVaultFilterer) WatchSettingsUpdated(opts *bind.WatchOpts, sink chan<- *ProtocolFeeVaultSettingsUpdated) (event.Subscription, error) {

	logs, sub, err := _ProtocolFeeVault.contract.WatchLogs(opts, "SettingsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolFeeVaultSettingsUpdated)
				if err := _ProtocolFeeVault.contract.UnpackLog(event, "SettingsUpdated", log); err != nil {
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

// ParseSettingsUpdated is a log parse operation binding the contract event 0x4b804a0bfbdc2639203b93035be561d86f65f52b7e14984f95ec9d298cafccac.
//
// Solidity: event SettingsUpdated(uint256 time)
func (_ProtocolFeeVault *ProtocolFeeVaultFilterer) ParseSettingsUpdated(log types.Log) (*ProtocolFeeVaultSettingsUpdated, error) {
	event := new(ProtocolFeeVaultSettingsUpdated)
	if err := _ProtocolFeeVault.contract.UnpackLog(event, "SettingsUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProtocolFeeVaultTokenSoldIterator is returned from FilterTokenSold and is used to iterate over the raw logs and unpacked data for TokenSold events raised by the ProtocolFeeVault contract.
type ProtocolFeeVaultTokenSoldIterator struct {
	Event *ProtocolFeeVaultTokenSold // Event containing the contract specifics and raw log

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
func (it *ProtocolFeeVaultTokenSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolFeeVaultTokenSold)
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
		it.Event = new(ProtocolFeeVaultTokenSold)
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
func (it *ProtocolFeeVaultTokenSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolFeeVaultTokenSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolFeeVaultTokenSold represents a TokenSold event raised by the ProtocolFeeVault contract.
type ProtocolFeeVaultTokenSold struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenSold is a free log retrieval operation binding the contract event 0xfe2ff4cf36ff7d2c2b06eb960897ee0d76d9c3e58da12feb7b93e86b226dd344.
//
// Solidity: event TokenSold(address token, uint256 amount)
func (_ProtocolFeeVault *ProtocolFeeVaultFilterer) FilterTokenSold(opts *bind.FilterOpts) (*ProtocolFeeVaultTokenSoldIterator, error) {

	logs, sub, err := _ProtocolFeeVault.contract.FilterLogs(opts, "TokenSold")
	if err != nil {
		return nil, err
	}
	return &ProtocolFeeVaultTokenSoldIterator{contract: _ProtocolFeeVault.contract, event: "TokenSold", logs: logs, sub: sub}, nil
}

// WatchTokenSold is a free log subscription operation binding the contract event 0xfe2ff4cf36ff7d2c2b06eb960897ee0d76d9c3e58da12feb7b93e86b226dd344.
//
// Solidity: event TokenSold(address token, uint256 amount)
func (_ProtocolFeeVault *ProtocolFeeVaultFilterer) WatchTokenSold(opts *bind.WatchOpts, sink chan<- *ProtocolFeeVaultTokenSold) (event.Subscription, error) {

	logs, sub, err := _ProtocolFeeVault.contract.WatchLogs(opts, "TokenSold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolFeeVaultTokenSold)
				if err := _ProtocolFeeVault.contract.UnpackLog(event, "TokenSold", log); err != nil {
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

// ParseTokenSold is a log parse operation binding the contract event 0xfe2ff4cf36ff7d2c2b06eb960897ee0d76d9c3e58da12feb7b93e86b226dd344.
//
// Solidity: event TokenSold(address token, uint256 amount)
func (_ProtocolFeeVault *ProtocolFeeVaultFilterer) ParseTokenSold(log types.Log) (*ProtocolFeeVaultTokenSold, error) {
	event := new(ProtocolFeeVaultTokenSold)
	if err := _ProtocolFeeVault.contract.UnpackLog(event, "TokenSold", log); err != nil {
		return nil, err
	}
	return event, nil
}
