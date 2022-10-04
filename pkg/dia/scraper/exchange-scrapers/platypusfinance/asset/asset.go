// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package asset

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
)

// AssetMetaData contains all meta data concerning the Asset contract.
var AssetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousCashPosition\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashBeingAdded\",\"type\":\"uint256\"}],\"name\":\"CashAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousCashPosition\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashBeingRemoved\",\"type\":\"uint256\"}],\"name\":\"CashRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousLiabilityPosition\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liabilityBeingAdded\",\"type\":\"uint256\"}],\"name\":\"LiabilityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousLiabilityPosition\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liabilityBeingRemoved\",\"type\":\"uint256\"}],\"name\":\"LiabilityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousMaxSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousPool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"PoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addCash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addLiability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregateAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlyingToken_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"aggregateAccount_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liability\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeCash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeLiability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregateAccount_\",\"type\":\"address\"}],\"name\":\"setAggregateAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply_\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"}],\"name\":\"setPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferUnderlyingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AssetABI is the input ABI used to generate the binding from.
// Deprecated: Use AssetMetaData.ABI instead.
var AssetABI = AssetMetaData.ABI

// Asset is an auto generated Go binding around an Ethereum contract.
type Asset struct {
	AssetCaller     // Read-only binding to the contract
	AssetTransactor // Write-only binding to the contract
	AssetFilterer   // Log filterer for contract events
}

// AssetCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetSession struct {
	Contract     *Asset            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetCallerSession struct {
	Contract *AssetCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AssetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetTransactorSession struct {
	Contract     *AssetTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetRaw struct {
	Contract *Asset // Generic contract binding to access the raw methods on
}

// AssetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetCallerRaw struct {
	Contract *AssetCaller // Generic read-only contract binding to access the raw methods on
}

// AssetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetTransactorRaw struct {
	Contract *AssetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAsset creates a new instance of Asset, bound to a specific deployed contract.
func NewAsset(address common.Address, backend bind.ContractBackend) (*Asset, error) {
	contract, err := bindAsset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Asset{AssetCaller: AssetCaller{contract: contract}, AssetTransactor: AssetTransactor{contract: contract}, AssetFilterer: AssetFilterer{contract: contract}}, nil
}

// NewAssetCaller creates a new read-only instance of Asset, bound to a specific deployed contract.
func NewAssetCaller(address common.Address, caller bind.ContractCaller) (*AssetCaller, error) {
	contract, err := bindAsset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetCaller{contract: contract}, nil
}

// NewAssetTransactor creates a new write-only instance of Asset, bound to a specific deployed contract.
func NewAssetTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetTransactor, error) {
	contract, err := bindAsset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetTransactor{contract: contract}, nil
}

// NewAssetFilterer creates a new log filterer instance of Asset, bound to a specific deployed contract.
func NewAssetFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetFilterer, error) {
	contract, err := bindAsset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetFilterer{contract: contract}, nil
}

// bindAsset binds a generic wrapper to an already deployed contract.
func bindAsset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Asset *AssetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Asset.Contract.AssetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Asset *AssetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Asset.Contract.AssetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Asset *AssetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Asset.Contract.AssetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Asset *AssetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Asset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Asset *AssetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Asset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Asset *AssetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Asset.Contract.contract.Transact(opts, method, params...)
}

// AggregateAccount is a free data retrieval call binding the contract method 0x7e1317fa.
//
// Solidity: function aggregateAccount() view returns(address)
func (_Asset *AssetCaller) AggregateAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "aggregateAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggregateAccount is a free data retrieval call binding the contract method 0x7e1317fa.
//
// Solidity: function aggregateAccount() view returns(address)
func (_Asset *AssetSession) AggregateAccount() (common.Address, error) {
	return _Asset.Contract.AggregateAccount(&_Asset.CallOpts)
}

// AggregateAccount is a free data retrieval call binding the contract method 0x7e1317fa.
//
// Solidity: function aggregateAccount() view returns(address)
func (_Asset *AssetCallerSession) AggregateAccount() (common.Address, error) {
	return _Asset.Contract.AggregateAccount(&_Asset.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Asset *AssetCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Asset *AssetSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Asset.Contract.Allowance(&_Asset.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Asset *AssetCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Asset.Contract.Allowance(&_Asset.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Asset *AssetCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Asset *AssetSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Asset.Contract.BalanceOf(&_Asset.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Asset *AssetCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Asset.Contract.BalanceOf(&_Asset.CallOpts, account)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_Asset *AssetCaller) Cash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "cash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_Asset *AssetSession) Cash() (*big.Int, error) {
	return _Asset.Contract.Cash(&_Asset.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_Asset *AssetCallerSession) Cash() (*big.Int, error) {
	return _Asset.Contract.Cash(&_Asset.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Asset *AssetCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Asset *AssetSession) Decimals() (uint8, error) {
	return _Asset.Contract.Decimals(&_Asset.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Asset *AssetCallerSession) Decimals() (uint8, error) {
	return _Asset.Contract.Decimals(&_Asset.CallOpts)
}

// Liability is a free data retrieval call binding the contract method 0x705727b5.
//
// Solidity: function liability() view returns(uint256)
func (_Asset *AssetCaller) Liability(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "liability")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liability is a free data retrieval call binding the contract method 0x705727b5.
//
// Solidity: function liability() view returns(uint256)
func (_Asset *AssetSession) Liability() (*big.Int, error) {
	return _Asset.Contract.Liability(&_Asset.CallOpts)
}

// Liability is a free data retrieval call binding the contract method 0x705727b5.
//
// Solidity: function liability() view returns(uint256)
func (_Asset *AssetCallerSession) Liability() (*big.Int, error) {
	return _Asset.Contract.Liability(&_Asset.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Asset *AssetCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Asset *AssetSession) MaxSupply() (*big.Int, error) {
	return _Asset.Contract.MaxSupply(&_Asset.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Asset *AssetCallerSession) MaxSupply() (*big.Int, error) {
	return _Asset.Contract.MaxSupply(&_Asset.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Asset *AssetCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Asset *AssetSession) Name() (string, error) {
	return _Asset.Contract.Name(&_Asset.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Asset *AssetCallerSession) Name() (string, error) {
	return _Asset.Contract.Name(&_Asset.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Asset *AssetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Asset *AssetSession) Owner() (common.Address, error) {
	return _Asset.Contract.Owner(&_Asset.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Asset *AssetCallerSession) Owner() (common.Address, error) {
	return _Asset.Contract.Owner(&_Asset.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Asset *AssetCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Asset *AssetSession) Pool() (common.Address, error) {
	return _Asset.Contract.Pool(&_Asset.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Asset *AssetCallerSession) Pool() (common.Address, error) {
	return _Asset.Contract.Pool(&_Asset.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Asset *AssetCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Asset *AssetSession) Symbol() (string, error) {
	return _Asset.Contract.Symbol(&_Asset.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Asset *AssetCallerSession) Symbol() (string, error) {
	return _Asset.Contract.Symbol(&_Asset.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Asset *AssetCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Asset *AssetSession) TotalSupply() (*big.Int, error) {
	return _Asset.Contract.TotalSupply(&_Asset.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Asset *AssetCallerSession) TotalSupply() (*big.Int, error) {
	return _Asset.Contract.TotalSupply(&_Asset.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_Asset *AssetCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_Asset *AssetSession) UnderlyingToken() (common.Address, error) {
	return _Asset.Contract.UnderlyingToken(&_Asset.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_Asset *AssetCallerSession) UnderlyingToken() (common.Address, error) {
	return _Asset.Contract.UnderlyingToken(&_Asset.CallOpts)
}

// UnderlyingTokenBalance is a free data retrieval call binding the contract method 0x99c91a64.
//
// Solidity: function underlyingTokenBalance() view returns(uint256)
func (_Asset *AssetCaller) UnderlyingTokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Asset.contract.Call(opts, &out, "underlyingTokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingTokenBalance is a free data retrieval call binding the contract method 0x99c91a64.
//
// Solidity: function underlyingTokenBalance() view returns(uint256)
func (_Asset *AssetSession) UnderlyingTokenBalance() (*big.Int, error) {
	return _Asset.Contract.UnderlyingTokenBalance(&_Asset.CallOpts)
}

// UnderlyingTokenBalance is a free data retrieval call binding the contract method 0x99c91a64.
//
// Solidity: function underlyingTokenBalance() view returns(uint256)
func (_Asset *AssetCallerSession) UnderlyingTokenBalance() (*big.Int, error) {
	return _Asset.Contract.UnderlyingTokenBalance(&_Asset.CallOpts)
}

// AddCash is a paid mutator transaction binding the contract method 0x16c9e7a0.
//
// Solidity: function addCash(uint256 amount) returns()
func (_Asset *AssetTransactor) AddCash(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "addCash", amount)
}

// AddCash is a paid mutator transaction binding the contract method 0x16c9e7a0.
//
// Solidity: function addCash(uint256 amount) returns()
func (_Asset *AssetSession) AddCash(amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.AddCash(&_Asset.TransactOpts, amount)
}

// AddCash is a paid mutator transaction binding the contract method 0x16c9e7a0.
//
// Solidity: function addCash(uint256 amount) returns()
func (_Asset *AssetTransactorSession) AddCash(amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.AddCash(&_Asset.TransactOpts, amount)
}

// AddLiability is a paid mutator transaction binding the contract method 0xa0f0f604.
//
// Solidity: function addLiability(uint256 amount) returns()
func (_Asset *AssetTransactor) AddLiability(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "addLiability", amount)
}

// AddLiability is a paid mutator transaction binding the contract method 0xa0f0f604.
//
// Solidity: function addLiability(uint256 amount) returns()
func (_Asset *AssetSession) AddLiability(amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.AddLiability(&_Asset.TransactOpts, amount)
}

// AddLiability is a paid mutator transaction binding the contract method 0xa0f0f604.
//
// Solidity: function addLiability(uint256 amount) returns()
func (_Asset *AssetTransactorSession) AddLiability(amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.AddLiability(&_Asset.TransactOpts, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Asset *AssetTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Asset *AssetSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Approve(&_Asset.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Asset *AssetTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Approve(&_Asset.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address to, uint256 amount) returns()
func (_Asset *AssetTransactor) Burn(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "burn", to, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address to, uint256 amount) returns()
func (_Asset *AssetSession) Burn(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Burn(&_Asset.TransactOpts, to, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address to, uint256 amount) returns()
func (_Asset *AssetTransactorSession) Burn(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Burn(&_Asset.TransactOpts, to, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Asset *AssetTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Asset *AssetSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.DecreaseAllowance(&_Asset.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Asset *AssetTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.DecreaseAllowance(&_Asset.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Asset *AssetTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Asset *AssetSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.IncreaseAllowance(&_Asset.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Asset *AssetTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.IncreaseAllowance(&_Asset.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x613d25bb.
//
// Solidity: function initialize(address underlyingToken_, string name_, string symbol_, address aggregateAccount_) returns()
func (_Asset *AssetTransactor) Initialize(opts *bind.TransactOpts, underlyingToken_ common.Address, name_ string, symbol_ string, aggregateAccount_ common.Address) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "initialize", underlyingToken_, name_, symbol_, aggregateAccount_)
}

// Initialize is a paid mutator transaction binding the contract method 0x613d25bb.
//
// Solidity: function initialize(address underlyingToken_, string name_, string symbol_, address aggregateAccount_) returns()
func (_Asset *AssetSession) Initialize(underlyingToken_ common.Address, name_ string, symbol_ string, aggregateAccount_ common.Address) (*types.Transaction, error) {
	return _Asset.Contract.Initialize(&_Asset.TransactOpts, underlyingToken_, name_, symbol_, aggregateAccount_)
}

// Initialize is a paid mutator transaction binding the contract method 0x613d25bb.
//
// Solidity: function initialize(address underlyingToken_, string name_, string symbol_, address aggregateAccount_) returns()
func (_Asset *AssetTransactorSession) Initialize(underlyingToken_ common.Address, name_ string, symbol_ string, aggregateAccount_ common.Address) (*types.Transaction, error) {
	return _Asset.Contract.Initialize(&_Asset.TransactOpts, underlyingToken_, name_, symbol_, aggregateAccount_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Asset *AssetTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Asset *AssetSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Mint(&_Asset.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Asset *AssetTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Mint(&_Asset.TransactOpts, to, amount)
}

// RemoveCash is a paid mutator transaction binding the contract method 0x9f9ef988.
//
// Solidity: function removeCash(uint256 amount) returns()
func (_Asset *AssetTransactor) RemoveCash(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "removeCash", amount)
}

// RemoveCash is a paid mutator transaction binding the contract method 0x9f9ef988.
//
// Solidity: function removeCash(uint256 amount) returns()
func (_Asset *AssetSession) RemoveCash(amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.RemoveCash(&_Asset.TransactOpts, amount)
}

// RemoveCash is a paid mutator transaction binding the contract method 0x9f9ef988.
//
// Solidity: function removeCash(uint256 amount) returns()
func (_Asset *AssetTransactorSession) RemoveCash(amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.RemoveCash(&_Asset.TransactOpts, amount)
}

// RemoveLiability is a paid mutator transaction binding the contract method 0xd8b87853.
//
// Solidity: function removeLiability(uint256 amount) returns()
func (_Asset *AssetTransactor) RemoveLiability(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "removeLiability", amount)
}

// RemoveLiability is a paid mutator transaction binding the contract method 0xd8b87853.
//
// Solidity: function removeLiability(uint256 amount) returns()
func (_Asset *AssetSession) RemoveLiability(amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.RemoveLiability(&_Asset.TransactOpts, amount)
}

// RemoveLiability is a paid mutator transaction binding the contract method 0xd8b87853.
//
// Solidity: function removeLiability(uint256 amount) returns()
func (_Asset *AssetTransactorSession) RemoveLiability(amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.RemoveLiability(&_Asset.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Asset *AssetTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Asset *AssetSession) RenounceOwnership() (*types.Transaction, error) {
	return _Asset.Contract.RenounceOwnership(&_Asset.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Asset *AssetTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Asset.Contract.RenounceOwnership(&_Asset.TransactOpts)
}

// SetAggregateAccount is a paid mutator transaction binding the contract method 0x95d6f7b9.
//
// Solidity: function setAggregateAccount(address aggregateAccount_) returns()
func (_Asset *AssetTransactor) SetAggregateAccount(opts *bind.TransactOpts, aggregateAccount_ common.Address) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "setAggregateAccount", aggregateAccount_)
}

// SetAggregateAccount is a paid mutator transaction binding the contract method 0x95d6f7b9.
//
// Solidity: function setAggregateAccount(address aggregateAccount_) returns()
func (_Asset *AssetSession) SetAggregateAccount(aggregateAccount_ common.Address) (*types.Transaction, error) {
	return _Asset.Contract.SetAggregateAccount(&_Asset.TransactOpts, aggregateAccount_)
}

// SetAggregateAccount is a paid mutator transaction binding the contract method 0x95d6f7b9.
//
// Solidity: function setAggregateAccount(address aggregateAccount_) returns()
func (_Asset *AssetTransactorSession) SetAggregateAccount(aggregateAccount_ common.Address) (*types.Transaction, error) {
	return _Asset.Contract.SetAggregateAccount(&_Asset.TransactOpts, aggregateAccount_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_Asset *AssetTransactor) SetMaxSupply(opts *bind.TransactOpts, maxSupply_ *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "setMaxSupply", maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_Asset *AssetSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.SetMaxSupply(&_Asset.TransactOpts, maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_Asset *AssetTransactorSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.SetMaxSupply(&_Asset.TransactOpts, maxSupply_)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address pool_) returns()
func (_Asset *AssetTransactor) SetPool(opts *bind.TransactOpts, pool_ common.Address) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "setPool", pool_)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address pool_) returns()
func (_Asset *AssetSession) SetPool(pool_ common.Address) (*types.Transaction, error) {
	return _Asset.Contract.SetPool(&_Asset.TransactOpts, pool_)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address pool_) returns()
func (_Asset *AssetTransactorSession) SetPool(pool_ common.Address) (*types.Transaction, error) {
	return _Asset.Contract.SetPool(&_Asset.TransactOpts, pool_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Asset *AssetTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Asset *AssetSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Transfer(&_Asset.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Asset *AssetTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Transfer(&_Asset.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Asset *AssetTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Asset *AssetSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.TransferFrom(&_Asset.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Asset *AssetTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.TransferFrom(&_Asset.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Asset *AssetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Asset *AssetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Asset.Contract.TransferOwnership(&_Asset.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Asset *AssetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Asset.Contract.TransferOwnership(&_Asset.TransactOpts, newOwner)
}

// TransferUnderlyingToken is a paid mutator transaction binding the contract method 0x9e79eaa5.
//
// Solidity: function transferUnderlyingToken(address to, uint256 amount) returns()
func (_Asset *AssetTransactor) TransferUnderlyingToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "transferUnderlyingToken", to, amount)
}

// TransferUnderlyingToken is a paid mutator transaction binding the contract method 0x9e79eaa5.
//
// Solidity: function transferUnderlyingToken(address to, uint256 amount) returns()
func (_Asset *AssetSession) TransferUnderlyingToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.TransferUnderlyingToken(&_Asset.TransactOpts, to, amount)
}

// TransferUnderlyingToken is a paid mutator transaction binding the contract method 0x9e79eaa5.
//
// Solidity: function transferUnderlyingToken(address to, uint256 amount) returns()
func (_Asset *AssetTransactorSession) TransferUnderlyingToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.TransferUnderlyingToken(&_Asset.TransactOpts, to, amount)
}

// AssetApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Asset contract.
type AssetApprovalIterator struct {
	Event *AssetApproval // Event containing the contract specifics and raw log

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
func (it *AssetApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetApproval)
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
		it.Event = new(AssetApproval)
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
func (it *AssetApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetApproval represents a Approval event raised by the Asset contract.
type AssetApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Asset *AssetFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AssetApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Asset.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AssetApprovalIterator{contract: _Asset.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Asset *AssetFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AssetApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Asset.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetApproval)
				if err := _Asset.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Asset *AssetFilterer) ParseApproval(log types.Log) (*AssetApproval, error) {
	event := new(AssetApproval)
	if err := _Asset.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetCashAddedIterator is returned from FilterCashAdded and is used to iterate over the raw logs and unpacked data for CashAdded events raised by the Asset contract.
type AssetCashAddedIterator struct {
	Event *AssetCashAdded // Event containing the contract specifics and raw log

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
func (it *AssetCashAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetCashAdded)
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
		it.Event = new(AssetCashAdded)
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
func (it *AssetCashAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetCashAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetCashAdded represents a CashAdded event raised by the Asset contract.
type AssetCashAdded struct {
	PreviousCashPosition *big.Int
	CashBeingAdded       *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCashAdded is a free log retrieval operation binding the contract event 0x04da412052b8d39d78da489e294630fcb3874f03dcb0ead4481c0a6d70df1e15.
//
// Solidity: event CashAdded(uint256 previousCashPosition, uint256 cashBeingAdded)
func (_Asset *AssetFilterer) FilterCashAdded(opts *bind.FilterOpts) (*AssetCashAddedIterator, error) {

	logs, sub, err := _Asset.contract.FilterLogs(opts, "CashAdded")
	if err != nil {
		return nil, err
	}
	return &AssetCashAddedIterator{contract: _Asset.contract, event: "CashAdded", logs: logs, sub: sub}, nil
}

// WatchCashAdded is a free log subscription operation binding the contract event 0x04da412052b8d39d78da489e294630fcb3874f03dcb0ead4481c0a6d70df1e15.
//
// Solidity: event CashAdded(uint256 previousCashPosition, uint256 cashBeingAdded)
func (_Asset *AssetFilterer) WatchCashAdded(opts *bind.WatchOpts, sink chan<- *AssetCashAdded) (event.Subscription, error) {

	logs, sub, err := _Asset.contract.WatchLogs(opts, "CashAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetCashAdded)
				if err := _Asset.contract.UnpackLog(event, "CashAdded", log); err != nil {
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

// ParseCashAdded is a log parse operation binding the contract event 0x04da412052b8d39d78da489e294630fcb3874f03dcb0ead4481c0a6d70df1e15.
//
// Solidity: event CashAdded(uint256 previousCashPosition, uint256 cashBeingAdded)
func (_Asset *AssetFilterer) ParseCashAdded(log types.Log) (*AssetCashAdded, error) {
	event := new(AssetCashAdded)
	if err := _Asset.contract.UnpackLog(event, "CashAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetCashRemovedIterator is returned from FilterCashRemoved and is used to iterate over the raw logs and unpacked data for CashRemoved events raised by the Asset contract.
type AssetCashRemovedIterator struct {
	Event *AssetCashRemoved // Event containing the contract specifics and raw log

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
func (it *AssetCashRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetCashRemoved)
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
		it.Event = new(AssetCashRemoved)
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
func (it *AssetCashRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetCashRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetCashRemoved represents a CashRemoved event raised by the Asset contract.
type AssetCashRemoved struct {
	PreviousCashPosition *big.Int
	CashBeingRemoved     *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCashRemoved is a free log retrieval operation binding the contract event 0xf15a954400c2f966714cd09162f79a6682b77351200ad1d595000057fc4ee999.
//
// Solidity: event CashRemoved(uint256 previousCashPosition, uint256 cashBeingRemoved)
func (_Asset *AssetFilterer) FilterCashRemoved(opts *bind.FilterOpts) (*AssetCashRemovedIterator, error) {

	logs, sub, err := _Asset.contract.FilterLogs(opts, "CashRemoved")
	if err != nil {
		return nil, err
	}
	return &AssetCashRemovedIterator{contract: _Asset.contract, event: "CashRemoved", logs: logs, sub: sub}, nil
}

// WatchCashRemoved is a free log subscription operation binding the contract event 0xf15a954400c2f966714cd09162f79a6682b77351200ad1d595000057fc4ee999.
//
// Solidity: event CashRemoved(uint256 previousCashPosition, uint256 cashBeingRemoved)
func (_Asset *AssetFilterer) WatchCashRemoved(opts *bind.WatchOpts, sink chan<- *AssetCashRemoved) (event.Subscription, error) {

	logs, sub, err := _Asset.contract.WatchLogs(opts, "CashRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetCashRemoved)
				if err := _Asset.contract.UnpackLog(event, "CashRemoved", log); err != nil {
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

// ParseCashRemoved is a log parse operation binding the contract event 0xf15a954400c2f966714cd09162f79a6682b77351200ad1d595000057fc4ee999.
//
// Solidity: event CashRemoved(uint256 previousCashPosition, uint256 cashBeingRemoved)
func (_Asset *AssetFilterer) ParseCashRemoved(log types.Log) (*AssetCashRemoved, error) {
	event := new(AssetCashRemoved)
	if err := _Asset.contract.UnpackLog(event, "CashRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetLiabilityAddedIterator is returned from FilterLiabilityAdded and is used to iterate over the raw logs and unpacked data for LiabilityAdded events raised by the Asset contract.
type AssetLiabilityAddedIterator struct {
	Event *AssetLiabilityAdded // Event containing the contract specifics and raw log

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
func (it *AssetLiabilityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetLiabilityAdded)
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
		it.Event = new(AssetLiabilityAdded)
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
func (it *AssetLiabilityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetLiabilityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetLiabilityAdded represents a LiabilityAdded event raised by the Asset contract.
type AssetLiabilityAdded struct {
	PreviousLiabilityPosition *big.Int
	LiabilityBeingAdded       *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterLiabilityAdded is a free log retrieval operation binding the contract event 0x2b74a49d287a99ef6b8a9f27aaef936372e282e0e95a6352f07c9fd12596655c.
//
// Solidity: event LiabilityAdded(uint256 previousLiabilityPosition, uint256 liabilityBeingAdded)
func (_Asset *AssetFilterer) FilterLiabilityAdded(opts *bind.FilterOpts) (*AssetLiabilityAddedIterator, error) {

	logs, sub, err := _Asset.contract.FilterLogs(opts, "LiabilityAdded")
	if err != nil {
		return nil, err
	}
	return &AssetLiabilityAddedIterator{contract: _Asset.contract, event: "LiabilityAdded", logs: logs, sub: sub}, nil
}

// WatchLiabilityAdded is a free log subscription operation binding the contract event 0x2b74a49d287a99ef6b8a9f27aaef936372e282e0e95a6352f07c9fd12596655c.
//
// Solidity: event LiabilityAdded(uint256 previousLiabilityPosition, uint256 liabilityBeingAdded)
func (_Asset *AssetFilterer) WatchLiabilityAdded(opts *bind.WatchOpts, sink chan<- *AssetLiabilityAdded) (event.Subscription, error) {

	logs, sub, err := _Asset.contract.WatchLogs(opts, "LiabilityAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetLiabilityAdded)
				if err := _Asset.contract.UnpackLog(event, "LiabilityAdded", log); err != nil {
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

// ParseLiabilityAdded is a log parse operation binding the contract event 0x2b74a49d287a99ef6b8a9f27aaef936372e282e0e95a6352f07c9fd12596655c.
//
// Solidity: event LiabilityAdded(uint256 previousLiabilityPosition, uint256 liabilityBeingAdded)
func (_Asset *AssetFilterer) ParseLiabilityAdded(log types.Log) (*AssetLiabilityAdded, error) {
	event := new(AssetLiabilityAdded)
	if err := _Asset.contract.UnpackLog(event, "LiabilityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetLiabilityRemovedIterator is returned from FilterLiabilityRemoved and is used to iterate over the raw logs and unpacked data for LiabilityRemoved events raised by the Asset contract.
type AssetLiabilityRemovedIterator struct {
	Event *AssetLiabilityRemoved // Event containing the contract specifics and raw log

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
func (it *AssetLiabilityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetLiabilityRemoved)
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
		it.Event = new(AssetLiabilityRemoved)
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
func (it *AssetLiabilityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetLiabilityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetLiabilityRemoved represents a LiabilityRemoved event raised by the Asset contract.
type AssetLiabilityRemoved struct {
	PreviousLiabilityPosition *big.Int
	LiabilityBeingRemoved     *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterLiabilityRemoved is a free log retrieval operation binding the contract event 0xdf20ac3c7d97136ceef3f041d542947447276d67c158dced2e33d1ee7984f530.
//
// Solidity: event LiabilityRemoved(uint256 previousLiabilityPosition, uint256 liabilityBeingRemoved)
func (_Asset *AssetFilterer) FilterLiabilityRemoved(opts *bind.FilterOpts) (*AssetLiabilityRemovedIterator, error) {

	logs, sub, err := _Asset.contract.FilterLogs(opts, "LiabilityRemoved")
	if err != nil {
		return nil, err
	}
	return &AssetLiabilityRemovedIterator{contract: _Asset.contract, event: "LiabilityRemoved", logs: logs, sub: sub}, nil
}

// WatchLiabilityRemoved is a free log subscription operation binding the contract event 0xdf20ac3c7d97136ceef3f041d542947447276d67c158dced2e33d1ee7984f530.
//
// Solidity: event LiabilityRemoved(uint256 previousLiabilityPosition, uint256 liabilityBeingRemoved)
func (_Asset *AssetFilterer) WatchLiabilityRemoved(opts *bind.WatchOpts, sink chan<- *AssetLiabilityRemoved) (event.Subscription, error) {

	logs, sub, err := _Asset.contract.WatchLogs(opts, "LiabilityRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetLiabilityRemoved)
				if err := _Asset.contract.UnpackLog(event, "LiabilityRemoved", log); err != nil {
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

// ParseLiabilityRemoved is a log parse operation binding the contract event 0xdf20ac3c7d97136ceef3f041d542947447276d67c158dced2e33d1ee7984f530.
//
// Solidity: event LiabilityRemoved(uint256 previousLiabilityPosition, uint256 liabilityBeingRemoved)
func (_Asset *AssetFilterer) ParseLiabilityRemoved(log types.Log) (*AssetLiabilityRemoved, error) {
	event := new(AssetLiabilityRemoved)
	if err := _Asset.contract.UnpackLog(event, "LiabilityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetMaxSupplyUpdatedIterator is returned from FilterMaxSupplyUpdated and is used to iterate over the raw logs and unpacked data for MaxSupplyUpdated events raised by the Asset contract.
type AssetMaxSupplyUpdatedIterator struct {
	Event *AssetMaxSupplyUpdated // Event containing the contract specifics and raw log

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
func (it *AssetMaxSupplyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetMaxSupplyUpdated)
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
		it.Event = new(AssetMaxSupplyUpdated)
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
func (it *AssetMaxSupplyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetMaxSupplyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetMaxSupplyUpdated represents a MaxSupplyUpdated event raised by the Asset contract.
type AssetMaxSupplyUpdated struct {
	PreviousMaxSupply *big.Int
	NewMaxSupply      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMaxSupplyUpdated is a free log retrieval operation binding the contract event 0x44ecfc706d63e347851cfd40acfa6cf2e3a41faa3e8b460210c03938e84a91ad.
//
// Solidity: event MaxSupplyUpdated(uint256 previousMaxSupply, uint256 newMaxSupply)
func (_Asset *AssetFilterer) FilterMaxSupplyUpdated(opts *bind.FilterOpts) (*AssetMaxSupplyUpdatedIterator, error) {

	logs, sub, err := _Asset.contract.FilterLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return &AssetMaxSupplyUpdatedIterator{contract: _Asset.contract, event: "MaxSupplyUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSupplyUpdated is a free log subscription operation binding the contract event 0x44ecfc706d63e347851cfd40acfa6cf2e3a41faa3e8b460210c03938e84a91ad.
//
// Solidity: event MaxSupplyUpdated(uint256 previousMaxSupply, uint256 newMaxSupply)
func (_Asset *AssetFilterer) WatchMaxSupplyUpdated(opts *bind.WatchOpts, sink chan<- *AssetMaxSupplyUpdated) (event.Subscription, error) {

	logs, sub, err := _Asset.contract.WatchLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetMaxSupplyUpdated)
				if err := _Asset.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
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

// ParseMaxSupplyUpdated is a log parse operation binding the contract event 0x44ecfc706d63e347851cfd40acfa6cf2e3a41faa3e8b460210c03938e84a91ad.
//
// Solidity: event MaxSupplyUpdated(uint256 previousMaxSupply, uint256 newMaxSupply)
func (_Asset *AssetFilterer) ParseMaxSupplyUpdated(log types.Log) (*AssetMaxSupplyUpdated, error) {
	event := new(AssetMaxSupplyUpdated)
	if err := _Asset.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Asset contract.
type AssetOwnershipTransferredIterator struct {
	Event *AssetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AssetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetOwnershipTransferred)
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
		it.Event = new(AssetOwnershipTransferred)
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
func (it *AssetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetOwnershipTransferred represents a OwnershipTransferred event raised by the Asset contract.
type AssetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Asset *AssetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AssetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Asset.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AssetOwnershipTransferredIterator{contract: _Asset.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Asset *AssetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AssetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Asset.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetOwnershipTransferred)
				if err := _Asset.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Asset *AssetFilterer) ParseOwnershipTransferred(log types.Log) (*AssetOwnershipTransferred, error) {
	event := new(AssetOwnershipTransferred)
	if err := _Asset.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetPoolUpdatedIterator is returned from FilterPoolUpdated and is used to iterate over the raw logs and unpacked data for PoolUpdated events raised by the Asset contract.
type AssetPoolUpdatedIterator struct {
	Event *AssetPoolUpdated // Event containing the contract specifics and raw log

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
func (it *AssetPoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetPoolUpdated)
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
		it.Event = new(AssetPoolUpdated)
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
func (it *AssetPoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetPoolUpdated represents a PoolUpdated event raised by the Asset contract.
type AssetPoolUpdated struct {
	PreviousPool common.Address
	NewPool      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPoolUpdated is a free log retrieval operation binding the contract event 0x90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a66627.
//
// Solidity: event PoolUpdated(address indexed previousPool, address indexed newPool)
func (_Asset *AssetFilterer) FilterPoolUpdated(opts *bind.FilterOpts, previousPool []common.Address, newPool []common.Address) (*AssetPoolUpdatedIterator, error) {

	var previousPoolRule []interface{}
	for _, previousPoolItem := range previousPool {
		previousPoolRule = append(previousPoolRule, previousPoolItem)
	}
	var newPoolRule []interface{}
	for _, newPoolItem := range newPool {
		newPoolRule = append(newPoolRule, newPoolItem)
	}

	logs, sub, err := _Asset.contract.FilterLogs(opts, "PoolUpdated", previousPoolRule, newPoolRule)
	if err != nil {
		return nil, err
	}
	return &AssetPoolUpdatedIterator{contract: _Asset.contract, event: "PoolUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolUpdated is a free log subscription operation binding the contract event 0x90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a66627.
//
// Solidity: event PoolUpdated(address indexed previousPool, address indexed newPool)
func (_Asset *AssetFilterer) WatchPoolUpdated(opts *bind.WatchOpts, sink chan<- *AssetPoolUpdated, previousPool []common.Address, newPool []common.Address) (event.Subscription, error) {

	var previousPoolRule []interface{}
	for _, previousPoolItem := range previousPool {
		previousPoolRule = append(previousPoolRule, previousPoolItem)
	}
	var newPoolRule []interface{}
	for _, newPoolItem := range newPool {
		newPoolRule = append(newPoolRule, newPoolItem)
	}

	logs, sub, err := _Asset.contract.WatchLogs(opts, "PoolUpdated", previousPoolRule, newPoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetPoolUpdated)
				if err := _Asset.contract.UnpackLog(event, "PoolUpdated", log); err != nil {
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

// ParsePoolUpdated is a log parse operation binding the contract event 0x90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a66627.
//
// Solidity: event PoolUpdated(address indexed previousPool, address indexed newPool)
func (_Asset *AssetFilterer) ParsePoolUpdated(log types.Log) (*AssetPoolUpdated, error) {
	event := new(AssetPoolUpdated)
	if err := _Asset.contract.UnpackLog(event, "PoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Asset contract.
type AssetTransferIterator struct {
	Event *AssetTransfer // Event containing the contract specifics and raw log

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
func (it *AssetTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetTransfer)
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
		it.Event = new(AssetTransfer)
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
func (it *AssetTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetTransfer represents a Transfer event raised by the Asset contract.
type AssetTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Asset *AssetFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AssetTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Asset.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AssetTransferIterator{contract: _Asset.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Asset *AssetFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AssetTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Asset.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetTransfer)
				if err := _Asset.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Asset *AssetFilterer) ParseTransfer(log types.Log) (*AssetTransfer, error) {
	event := new(AssetTransfer)
	if err := _Asset.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
