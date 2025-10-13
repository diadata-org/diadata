// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_getterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_redeemerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_susdpVaultAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outputDecimals\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"name\":\"OracleUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGetterAddress\",\"type\":\"address\"}],\"name\":\"setGetterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"compressedValues\",\"type\":\"uint256[]\"}],\"name\":\"setMultipleValues\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDecimals\",\"type\":\"uint256\"}],\"name\":\"setOutputDecimals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRedeemerAddress\",\"type\":\"address\"}],\"name\":\"setRedeemerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newSusdpVaultAddress\",\"type\":\"address\"}],\"name\":\"setSusdpVaultAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newUsdpQueryString\",\"type\":\"string\"}],\"name\":\"setUsdpQueryString\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainSpecificUsdpIssued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainSpecificUsdpPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSusdpPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getterAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outputDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUERY_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"susdpVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDP_QUERY_STRING\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// QUERYDECIMALS is a free data retrieval call binding the contract method 0xf904dbcc.
//
// Solidity: function QUERY_DECIMALS() view returns(uint256)
func (_Main *MainCaller) QUERYDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "QUERY_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUERYDECIMALS is a free data retrieval call binding the contract method 0xf904dbcc.
//
// Solidity: function QUERY_DECIMALS() view returns(uint256)
func (_Main *MainSession) QUERYDECIMALS() (*big.Int, error) {
	return _Main.Contract.QUERYDECIMALS(&_Main.CallOpts)
}

// QUERYDECIMALS is a free data retrieval call binding the contract method 0xf904dbcc.
//
// Solidity: function QUERY_DECIMALS() view returns(uint256)
func (_Main *MainCallerSession) QUERYDECIMALS() (*big.Int, error) {
	return _Main.Contract.QUERYDECIMALS(&_Main.CallOpts)
}

// USDPQUERYSTRING is a free data retrieval call binding the contract method 0x2b3eb84e.
//
// Solidity: function USDP_QUERY_STRING() view returns(string)
func (_Main *MainCaller) USDPQUERYSTRING(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "USDP_QUERY_STRING")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// USDPQUERYSTRING is a free data retrieval call binding the contract method 0x2b3eb84e.
//
// Solidity: function USDP_QUERY_STRING() view returns(string)
func (_Main *MainSession) USDPQUERYSTRING() (string, error) {
	return _Main.Contract.USDPQUERYSTRING(&_Main.CallOpts)
}

// USDPQUERYSTRING is a free data retrieval call binding the contract method 0x2b3eb84e.
//
// Solidity: function USDP_QUERY_STRING() view returns(string)
func (_Main *MainCallerSession) USDPQUERYSTRING() (string, error) {
	return _Main.Contract.USDPQUERYSTRING(&_Main.CallOpts)
}

// GetChainSpecificUsdpIssued is a free data retrieval call binding the contract method 0xa44cee4f.
//
// Solidity: function getChainSpecificUsdpIssued() view returns(uint256)
func (_Main *MainCaller) GetChainSpecificUsdpIssued(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getChainSpecificUsdpIssued")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainSpecificUsdpIssued is a free data retrieval call binding the contract method 0xa44cee4f.
//
// Solidity: function getChainSpecificUsdpIssued() view returns(uint256)
func (_Main *MainSession) GetChainSpecificUsdpIssued() (*big.Int, error) {
	return _Main.Contract.GetChainSpecificUsdpIssued(&_Main.CallOpts)
}

// GetChainSpecificUsdpIssued is a free data retrieval call binding the contract method 0xa44cee4f.
//
// Solidity: function getChainSpecificUsdpIssued() view returns(uint256)
func (_Main *MainCallerSession) GetChainSpecificUsdpIssued() (*big.Int, error) {
	return _Main.Contract.GetChainSpecificUsdpIssued(&_Main.CallOpts)
}

// GetChainSpecificUsdpPrice is a free data retrieval call binding the contract method 0xf190fb89.
//
// Solidity: function getChainSpecificUsdpPrice() view returns(uint256)
func (_Main *MainCaller) GetChainSpecificUsdpPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getChainSpecificUsdpPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainSpecificUsdpPrice is a free data retrieval call binding the contract method 0xf190fb89.
//
// Solidity: function getChainSpecificUsdpPrice() view returns(uint256)
func (_Main *MainSession) GetChainSpecificUsdpPrice() (*big.Int, error) {
	return _Main.Contract.GetChainSpecificUsdpPrice(&_Main.CallOpts)
}

// GetChainSpecificUsdpPrice is a free data retrieval call binding the contract method 0xf190fb89.
//
// Solidity: function getChainSpecificUsdpPrice() view returns(uint256)
func (_Main *MainCallerSession) GetChainSpecificUsdpPrice() (*big.Int, error) {
	return _Main.Contract.GetChainSpecificUsdpPrice(&_Main.CallOpts)
}

// GetSusdpPrice is a free data retrieval call binding the contract method 0xa1b6012f.
//
// Solidity: function getSusdpPrice() view returns(uint256)
func (_Main *MainCaller) GetSusdpPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getSusdpPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSusdpPrice is a free data retrieval call binding the contract method 0xa1b6012f.
//
// Solidity: function getSusdpPrice() view returns(uint256)
func (_Main *MainSession) GetSusdpPrice() (*big.Int, error) {
	return _Main.Contract.GetSusdpPrice(&_Main.CallOpts)
}

// GetSusdpPrice is a free data retrieval call binding the contract method 0xa1b6012f.
//
// Solidity: function getSusdpPrice() view returns(uint256)
func (_Main *MainCallerSession) GetSusdpPrice() (*big.Int, error) {
	return _Main.Contract.GetSusdpPrice(&_Main.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_Main *MainCaller) GetValue(opts *bind.CallOpts, key string) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getValue", key)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_Main *MainSession) GetValue(key string) (*big.Int, *big.Int, error) {
	return _Main.Contract.GetValue(&_Main.CallOpts, key)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_Main *MainCallerSession) GetValue(key string) (*big.Int, *big.Int, error) {
	return _Main.Contract.GetValue(&_Main.CallOpts, key)
}

// GetterAddress is a free data retrieval call binding the contract method 0x5daf9a97.
//
// Solidity: function getterAddress() view returns(address)
func (_Main *MainCaller) GetterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getterAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetterAddress is a free data retrieval call binding the contract method 0x5daf9a97.
//
// Solidity: function getterAddress() view returns(address)
func (_Main *MainSession) GetterAddress() (common.Address, error) {
	return _Main.Contract.GetterAddress(&_Main.CallOpts)
}

// GetterAddress is a free data retrieval call binding the contract method 0x5daf9a97.
//
// Solidity: function getterAddress() view returns(address)
func (_Main *MainCallerSession) GetterAddress() (common.Address, error) {
	return _Main.Contract.GetterAddress(&_Main.CallOpts)
}

// OutputDecimals is a free data retrieval call binding the contract method 0xbf560ad8.
//
// Solidity: function outputDecimals() view returns(uint256)
func (_Main *MainCaller) OutputDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "outputDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutputDecimals is a free data retrieval call binding the contract method 0xbf560ad8.
//
// Solidity: function outputDecimals() view returns(uint256)
func (_Main *MainSession) OutputDecimals() (*big.Int, error) {
	return _Main.Contract.OutputDecimals(&_Main.CallOpts)
}

// OutputDecimals is a free data retrieval call binding the contract method 0xbf560ad8.
//
// Solidity: function outputDecimals() view returns(uint256)
func (_Main *MainCallerSession) OutputDecimals() (*big.Int, error) {
	return _Main.Contract.OutputDecimals(&_Main.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCallerSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// RedeemerAddress is a free data retrieval call binding the contract method 0x1157b834.
//
// Solidity: function redeemerAddress() view returns(address)
func (_Main *MainCaller) RedeemerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "redeemerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RedeemerAddress is a free data retrieval call binding the contract method 0x1157b834.
//
// Solidity: function redeemerAddress() view returns(address)
func (_Main *MainSession) RedeemerAddress() (common.Address, error) {
	return _Main.Contract.RedeemerAddress(&_Main.CallOpts)
}

// RedeemerAddress is a free data retrieval call binding the contract method 0x1157b834.
//
// Solidity: function redeemerAddress() view returns(address)
func (_Main *MainCallerSession) RedeemerAddress() (common.Address, error) {
	return _Main.Contract.RedeemerAddress(&_Main.CallOpts)
}

// SusdpVaultAddress is a free data retrieval call binding the contract method 0xe65f4a07.
//
// Solidity: function susdpVaultAddress() view returns(address)
func (_Main *MainCaller) SusdpVaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "susdpVaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SusdpVaultAddress is a free data retrieval call binding the contract method 0xe65f4a07.
//
// Solidity: function susdpVaultAddress() view returns(address)
func (_Main *MainSession) SusdpVaultAddress() (common.Address, error) {
	return _Main.Contract.SusdpVaultAddress(&_Main.CallOpts)
}

// SusdpVaultAddress is a free data retrieval call binding the contract method 0xe65f4a07.
//
// Solidity: function susdpVaultAddress() view returns(address)
func (_Main *MainCallerSession) SusdpVaultAddress() (common.Address, error) {
	return _Main.Contract.SusdpVaultAddress(&_Main.CallOpts)
}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_Main *MainCaller) Values(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "values", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_Main *MainSession) Values(arg0 string) (*big.Int, error) {
	return _Main.Contract.Values(&_Main.CallOpts, arg0)
}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_Main *MainCallerSession) Values(arg0 string) (*big.Int, error) {
	return _Main.Contract.Values(&_Main.CallOpts, arg0)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Main *MainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Main *MainSession) RenounceOwnership() (*types.Transaction, error) {
	return _Main.Contract.RenounceOwnership(&_Main.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Main *MainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Main.Contract.RenounceOwnership(&_Main.TransactOpts)
}

// SetGetterAddress is a paid mutator transaction binding the contract method 0x45eb8b44.
//
// Solidity: function setGetterAddress(address _newGetterAddress) returns()
func (_Main *MainTransactor) SetGetterAddress(opts *bind.TransactOpts, _newGetterAddress common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setGetterAddress", _newGetterAddress)
}

// SetGetterAddress is a paid mutator transaction binding the contract method 0x45eb8b44.
//
// Solidity: function setGetterAddress(address _newGetterAddress) returns()
func (_Main *MainSession) SetGetterAddress(_newGetterAddress common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetGetterAddress(&_Main.TransactOpts, _newGetterAddress)
}

// SetGetterAddress is a paid mutator transaction binding the contract method 0x45eb8b44.
//
// Solidity: function setGetterAddress(address _newGetterAddress) returns()
func (_Main *MainTransactorSession) SetGetterAddress(_newGetterAddress common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetGetterAddress(&_Main.TransactOpts, _newGetterAddress)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x8d241526.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] compressedValues) returns()
func (_Main *MainTransactor) SetMultipleValues(opts *bind.TransactOpts, keys []string, compressedValues []*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setMultipleValues", keys, compressedValues)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x8d241526.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] compressedValues) returns()
func (_Main *MainSession) SetMultipleValues(keys []string, compressedValues []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetMultipleValues(&_Main.TransactOpts, keys, compressedValues)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x8d241526.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] compressedValues) returns()
func (_Main *MainTransactorSession) SetMultipleValues(keys []string, compressedValues []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetMultipleValues(&_Main.TransactOpts, keys, compressedValues)
}

// SetOutputDecimals is a paid mutator transaction binding the contract method 0x6b59e8c3.
//
// Solidity: function setOutputDecimals(uint256 _newDecimals) returns()
func (_Main *MainTransactor) SetOutputDecimals(opts *bind.TransactOpts, _newDecimals *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setOutputDecimals", _newDecimals)
}

// SetOutputDecimals is a paid mutator transaction binding the contract method 0x6b59e8c3.
//
// Solidity: function setOutputDecimals(uint256 _newDecimals) returns()
func (_Main *MainSession) SetOutputDecimals(_newDecimals *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetOutputDecimals(&_Main.TransactOpts, _newDecimals)
}

// SetOutputDecimals is a paid mutator transaction binding the contract method 0x6b59e8c3.
//
// Solidity: function setOutputDecimals(uint256 _newDecimals) returns()
func (_Main *MainTransactorSession) SetOutputDecimals(_newDecimals *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetOutputDecimals(&_Main.TransactOpts, _newDecimals)
}

// SetRedeemerAddress is a paid mutator transaction binding the contract method 0xa2af0419.
//
// Solidity: function setRedeemerAddress(address _newRedeemerAddress) returns()
func (_Main *MainTransactor) SetRedeemerAddress(opts *bind.TransactOpts, _newRedeemerAddress common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setRedeemerAddress", _newRedeemerAddress)
}

// SetRedeemerAddress is a paid mutator transaction binding the contract method 0xa2af0419.
//
// Solidity: function setRedeemerAddress(address _newRedeemerAddress) returns()
func (_Main *MainSession) SetRedeemerAddress(_newRedeemerAddress common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetRedeemerAddress(&_Main.TransactOpts, _newRedeemerAddress)
}

// SetRedeemerAddress is a paid mutator transaction binding the contract method 0xa2af0419.
//
// Solidity: function setRedeemerAddress(address _newRedeemerAddress) returns()
func (_Main *MainTransactorSession) SetRedeemerAddress(_newRedeemerAddress common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetRedeemerAddress(&_Main.TransactOpts, _newRedeemerAddress)
}

// SetSusdpVaultAddress is a paid mutator transaction binding the contract method 0x73893608.
//
// Solidity: function setSusdpVaultAddress(address _newSusdpVaultAddress) returns()
func (_Main *MainTransactor) SetSusdpVaultAddress(opts *bind.TransactOpts, _newSusdpVaultAddress common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setSusdpVaultAddress", _newSusdpVaultAddress)
}

// SetSusdpVaultAddress is a paid mutator transaction binding the contract method 0x73893608.
//
// Solidity: function setSusdpVaultAddress(address _newSusdpVaultAddress) returns()
func (_Main *MainSession) SetSusdpVaultAddress(_newSusdpVaultAddress common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetSusdpVaultAddress(&_Main.TransactOpts, _newSusdpVaultAddress)
}

// SetSusdpVaultAddress is a paid mutator transaction binding the contract method 0x73893608.
//
// Solidity: function setSusdpVaultAddress(address _newSusdpVaultAddress) returns()
func (_Main *MainTransactorSession) SetSusdpVaultAddress(_newSusdpVaultAddress common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetSusdpVaultAddress(&_Main.TransactOpts, _newSusdpVaultAddress)
}

// SetUsdpQueryString is a paid mutator transaction binding the contract method 0x495ae3cf.
//
// Solidity: function setUsdpQueryString(string newUsdpQueryString) returns()
func (_Main *MainTransactor) SetUsdpQueryString(opts *bind.TransactOpts, newUsdpQueryString string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setUsdpQueryString", newUsdpQueryString)
}

// SetUsdpQueryString is a paid mutator transaction binding the contract method 0x495ae3cf.
//
// Solidity: function setUsdpQueryString(string newUsdpQueryString) returns()
func (_Main *MainSession) SetUsdpQueryString(newUsdpQueryString string) (*types.Transaction, error) {
	return _Main.Contract.SetUsdpQueryString(&_Main.TransactOpts, newUsdpQueryString)
}

// SetUsdpQueryString is a paid mutator transaction binding the contract method 0x495ae3cf.
//
// Solidity: function setUsdpQueryString(string newUsdpQueryString) returns()
func (_Main *MainTransactorSession) SetUsdpQueryString(newUsdpQueryString string) (*types.Transaction, error) {
	return _Main.Contract.SetUsdpQueryString(&_Main.TransactOpts, newUsdpQueryString)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_Main *MainTransactor) SetValue(opts *bind.TransactOpts, key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setValue", key, value, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_Main *MainSession) SetValue(key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetValue(&_Main.TransactOpts, key, value, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_Main *MainTransactorSession) SetValue(key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetValue(&_Main.TransactOpts, key, value, timestamp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.TransferOwnership(&_Main.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.TransferOwnership(&_Main.TransactOpts, newOwner)
}

// MainOracleUpdateIterator is returned from FilterOracleUpdate and is used to iterate over the raw logs and unpacked data for OracleUpdate events raised by the Main contract.
type MainOracleUpdateIterator struct {
	Event *MainOracleUpdate // Event containing the contract specifics and raw log

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
func (it *MainOracleUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainOracleUpdate)
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
		it.Event = new(MainOracleUpdate)
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
func (it *MainOracleUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainOracleUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainOracleUpdate represents a OracleUpdate event raised by the Main contract.
type MainOracleUpdate struct {
	Key       string
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleUpdate is a free log retrieval operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_Main *MainFilterer) FilterOracleUpdate(opts *bind.FilterOpts) (*MainOracleUpdateIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "OracleUpdate")
	if err != nil {
		return nil, err
	}
	return &MainOracleUpdateIterator{contract: _Main.contract, event: "OracleUpdate", logs: logs, sub: sub}, nil
}

// WatchOracleUpdate is a free log subscription operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_Main *MainFilterer) WatchOracleUpdate(opts *bind.WatchOpts, sink chan<- *MainOracleUpdate) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "OracleUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainOracleUpdate)
				if err := _Main.contract.UnpackLog(event, "OracleUpdate", log); err != nil {
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

// ParseOracleUpdate is a log parse operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_Main *MainFilterer) ParseOracleUpdate(log types.Log) (*MainOracleUpdate, error) {
	event := new(MainOracleUpdate)
	if err := _Main.contract.UnpackLog(event, "OracleUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Main contract.
type MainOwnershipTransferredIterator struct {
	Event *MainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainOwnershipTransferred)
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
		it.Event = new(MainOwnershipTransferred)
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
func (it *MainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainOwnershipTransferred represents a OwnershipTransferred event raised by the Main contract.
type MainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Main *MainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MainOwnershipTransferredIterator{contract: _Main.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Main *MainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainOwnershipTransferred)
				if err := _Main.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Main *MainFilterer) ParseOwnershipTransferred(log types.Log) (*MainOwnershipTransferred, error) {
	event := new(MainOwnershipTransferred)
	if err := _Main.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
