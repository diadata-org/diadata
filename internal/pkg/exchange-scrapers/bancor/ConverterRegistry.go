// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ConverterRegistry

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

// ConverterRegistryABI is the input ABI used to generate the binding from.
const ConverterRegistryABI = "[{\"inputs\":[{\"internalType\":\"contractIContractRegistry\",\"name\":\"_registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIConverterAnchor\",\"name\":\"_anchor\",\"type\":\"address\"}],\"name\":\"ConverterAnchorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIConverterAnchor\",\"name\":\"_anchor\",\"type\":\"address\"}],\"name\":\"ConverterAnchorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20Token\",\"name\":\"_convertibleToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIConverterAnchor\",\"name\":\"_smartToken\",\"type\":\"address\"}],\"name\":\"ConvertibleTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20Token\",\"name\":\"_convertibleToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIConverterAnchor\",\"name\":\"_smartToken\",\"type\":\"address\"}],\"name\":\"ConvertibleTokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIConverterAnchor\",\"name\":\"_liquidityPool\",\"type\":\"address\"}],\"name\":\"LiquidityPoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIConverterAnchor\",\"name\":\"_liquidityPool\",\"type\":\"address\"}],\"name\":\"LiquidityPoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIConverterAnchor\",\"name\":\"_smartToken\",\"type\":\"address\"}],\"name\":\"SmartTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIConverterAnchor\",\"name\":\"_smartToken\",\"type\":\"address\"}],\"name\":\"SmartTokenRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConverter\",\"name\":\"_converter\",\"type\":\"address\"}],\"name\":\"addConverter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getAnchor\",\"outputs\":[{\"internalType\":\"contractIConverterAnchor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAnchorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAnchors\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_anchors\",\"type\":\"address[]\"}],\"name\":\"getConvertersByAnchors\",\"outputs\":[{\"internalType\":\"contractIConverter[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_smartTokens\",\"type\":\"address[]\"}],\"name\":\"getConvertersBySmartTokens\",\"outputs\":[{\"internalType\":\"contractIConverter[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getConvertibleToken\",\"outputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_convertibleToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getConvertibleTokenAnchor\",\"outputs\":[{\"internalType\":\"contractIConverterAnchor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_convertibleToken\",\"type\":\"address\"}],\"name\":\"getConvertibleTokenAnchorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_convertibleToken\",\"type\":\"address\"}],\"name\":\"getConvertibleTokenAnchors\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConvertibleTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_convertibleToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getConvertibleTokenSmartToken\",\"outputs\":[{\"internalType\":\"contractIConverterAnchor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_convertibleToken\",\"type\":\"address\"}],\"name\":\"getConvertibleTokenSmartTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_convertibleToken\",\"type\":\"address\"}],\"name\":\"getConvertibleTokenSmartTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConvertibleTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getLiquidityPool\",\"outputs\":[{\"internalType\":\"contractIConverterAnchor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_type\",\"type\":\"uint16\"},{\"internalType\":\"contractIERC20Token[]\",\"name\":\"_reserveTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"_reserveWeights\",\"type\":\"uint32[]\"}],\"name\":\"getLiquidityPoolByConfig\",\"outputs\":[{\"internalType\":\"contractIConverterAnchor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token[]\",\"name\":\"_reserveTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"_reserveWeights\",\"type\":\"uint32[]\"}],\"name\":\"getLiquidityPoolByReserveConfig\",\"outputs\":[{\"internalType\":\"contractIConverterAnchor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidityPoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidityPools\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getSmartToken\",\"outputs\":[{\"internalType\":\"contractIConverterAnchor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSmartTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSmartTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_value\",\"type\":\"address\"}],\"name\":\"isAnchor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConverter\",\"name\":\"_converter\",\"type\":\"address\"}],\"name\":\"isConverterValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_value\",\"type\":\"address\"}],\"name\":\"isConvertibleToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_convertibleToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_value\",\"type\":\"address\"}],\"name\":\"isConvertibleTokenAnchor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_convertibleToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_value\",\"type\":\"address\"}],\"name\":\"isConvertibleTokenSmartToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_value\",\"type\":\"address\"}],\"name\":\"isLiquidityPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConverter\",\"name\":\"_converter\",\"type\":\"address\"}],\"name\":\"isSimilarLiquidityPoolRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_value\",\"type\":\"address\"}],\"name\":\"isSmartToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_type\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_maxConversionFee\",\"type\":\"uint32\"},{\"internalType\":\"contractIERC20Token[]\",\"name\":\"_reserveTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"_reserveWeights\",\"type\":\"uint32[]\"}],\"name\":\"newConverter\",\"outputs\":[{\"internalType\":\"contractIConverter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onlyOwnerCanUpdateRegistry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevRegistry\",\"outputs\":[{\"internalType\":\"contractIContractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIContractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConverter\",\"name\":\"_converter\",\"type\":\"address\"}],\"name\":\"removeConverter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restoreRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_onlyOwnerCanUpdateRegistry\",\"type\":\"bool\"}],\"name\":\"restrictRegistryUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ConverterRegistry is an auto generated Go binding around an Ethereum contract.
type ConverterRegistry struct {
	ConverterRegistryCaller     // Read-only binding to the contract
	ConverterRegistryTransactor // Write-only binding to the contract
	ConverterRegistryFilterer   // Log filterer for contract events
}

// ConverterRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConverterRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConverterRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConverterRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConverterRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConverterRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConverterRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConverterRegistrySession struct {
	Contract     *ConverterRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConverterRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConverterRegistryCallerSession struct {
	Contract *ConverterRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ConverterRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConverterRegistryTransactorSession struct {
	Contract     *ConverterRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ConverterRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConverterRegistryRaw struct {
	Contract *ConverterRegistry // Generic contract binding to access the raw methods on
}

// ConverterRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConverterRegistryCallerRaw struct {
	Contract *ConverterRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ConverterRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConverterRegistryTransactorRaw struct {
	Contract *ConverterRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConverterRegistry creates a new instance of ConverterRegistry, bound to a specific deployed contract.
func NewConverterRegistry(address common.Address, backend bind.ContractBackend) (*ConverterRegistry, error) {
	contract, err := bindConverterRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConverterRegistry{ConverterRegistryCaller: ConverterRegistryCaller{contract: contract}, ConverterRegistryTransactor: ConverterRegistryTransactor{contract: contract}, ConverterRegistryFilterer: ConverterRegistryFilterer{contract: contract}}, nil
}

// NewConverterRegistryCaller creates a new read-only instance of ConverterRegistry, bound to a specific deployed contract.
func NewConverterRegistryCaller(address common.Address, caller bind.ContractCaller) (*ConverterRegistryCaller, error) {
	contract, err := bindConverterRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConverterRegistryCaller{contract: contract}, nil
}

// NewConverterRegistryTransactor creates a new write-only instance of ConverterRegistry, bound to a specific deployed contract.
func NewConverterRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ConverterRegistryTransactor, error) {
	contract, err := bindConverterRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConverterRegistryTransactor{contract: contract}, nil
}

// NewConverterRegistryFilterer creates a new log filterer instance of ConverterRegistry, bound to a specific deployed contract.
func NewConverterRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ConverterRegistryFilterer, error) {
	contract, err := bindConverterRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConverterRegistryFilterer{contract: contract}, nil
}

// bindConverterRegistry binds a generic wrapper to an already deployed contract.
func bindConverterRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConverterRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConverterRegistry *ConverterRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConverterRegistry.Contract.ConverterRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConverterRegistry *ConverterRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterRegistry.Contract.ConverterRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConverterRegistry *ConverterRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConverterRegistry.Contract.ConverterRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConverterRegistry *ConverterRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConverterRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConverterRegistry *ConverterRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConverterRegistry *ConverterRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConverterRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetAnchor is a free data retrieval call binding the contract method 0x4c7df18f.
//
// Solidity: function getAnchor(uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistryCaller) GetAnchor(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getAnchor", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAnchor is a free data retrieval call binding the contract method 0x4c7df18f.
//
// Solidity: function getAnchor(uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistrySession) GetAnchor(_index *big.Int) (common.Address, error) {
	return _ConverterRegistry.Contract.GetAnchor(&_ConverterRegistry.CallOpts, _index)
}

// GetAnchor is a free data retrieval call binding the contract method 0x4c7df18f.
//
// Solidity: function getAnchor(uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistryCallerSession) GetAnchor(_index *big.Int) (common.Address, error) {
	return _ConverterRegistry.Contract.GetAnchor(&_ConverterRegistry.CallOpts, _index)
}

// GetAnchorCount is a free data retrieval call binding the contract method 0xd3182bed.
//
// Solidity: function getAnchorCount() view returns(uint256)
func (_ConverterRegistry *ConverterRegistryCaller) GetAnchorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getAnchorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAnchorCount is a free data retrieval call binding the contract method 0xd3182bed.
//
// Solidity: function getAnchorCount() view returns(uint256)
func (_ConverterRegistry *ConverterRegistrySession) GetAnchorCount() (*big.Int, error) {
	return _ConverterRegistry.Contract.GetAnchorCount(&_ConverterRegistry.CallOpts)
}

// GetAnchorCount is a free data retrieval call binding the contract method 0xd3182bed.
//
// Solidity: function getAnchorCount() view returns(uint256)
func (_ConverterRegistry *ConverterRegistryCallerSession) GetAnchorCount() (*big.Int, error) {
	return _ConverterRegistry.Contract.GetAnchorCount(&_ConverterRegistry.CallOpts)
}

// GetAnchors is a free data retrieval call binding the contract method 0xeffb3c6e.
//
// Solidity: function getAnchors() view returns(address[])
func (_ConverterRegistry *ConverterRegistryCaller) GetAnchors(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getAnchors")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAnchors is a free data retrieval call binding the contract method 0xeffb3c6e.
//
// Solidity: function getAnchors() view returns(address[])
func (_ConverterRegistry *ConverterRegistrySession) GetAnchors() ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetAnchors(&_ConverterRegistry.CallOpts)
}

// GetAnchors is a free data retrieval call binding the contract method 0xeffb3c6e.
//
// Solidity: function getAnchors() view returns(address[])
func (_ConverterRegistry *ConverterRegistryCallerSession) GetAnchors() ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetAnchors(&_ConverterRegistry.CallOpts)
}

// GetConvertersByAnchors is a free data retrieval call binding the contract method 0x610c0b05.
//
// Solidity: function getConvertersByAnchors(address[] _anchors) view returns(address[])
func (_ConverterRegistry *ConverterRegistryCaller) GetConvertersByAnchors(opts *bind.CallOpts, _anchors []common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getConvertersByAnchors", _anchors)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetConvertersByAnchors is a free data retrieval call binding the contract method 0x610c0b05.
//
// Solidity: function getConvertersByAnchors(address[] _anchors) view returns(address[])
func (_ConverterRegistry *ConverterRegistrySession) GetConvertersByAnchors(_anchors []common.Address) ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertersByAnchors(&_ConverterRegistry.CallOpts, _anchors)
}

// GetConvertersByAnchors is a free data retrieval call binding the contract method 0x610c0b05.
//
// Solidity: function getConvertersByAnchors(address[] _anchors) view returns(address[])
func (_ConverterRegistry *ConverterRegistryCallerSession) GetConvertersByAnchors(_anchors []common.Address) ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertersByAnchors(&_ConverterRegistry.CallOpts, _anchors)
}

// GetConvertersBySmartTokens is a free data retrieval call binding the contract method 0x1f8e2620.
//
// Solidity: function getConvertersBySmartTokens(address[] _smartTokens) view returns(address[])
func (_ConverterRegistry *ConverterRegistryCaller) GetConvertersBySmartTokens(opts *bind.CallOpts, _smartTokens []common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getConvertersBySmartTokens", _smartTokens)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetConvertersBySmartTokens is a free data retrieval call binding the contract method 0x1f8e2620.
//
// Solidity: function getConvertersBySmartTokens(address[] _smartTokens) view returns(address[])
func (_ConverterRegistry *ConverterRegistrySession) GetConvertersBySmartTokens(_smartTokens []common.Address) ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertersBySmartTokens(&_ConverterRegistry.CallOpts, _smartTokens)
}

// GetConvertersBySmartTokens is a free data retrieval call binding the contract method 0x1f8e2620.
//
// Solidity: function getConvertersBySmartTokens(address[] _smartTokens) view returns(address[])
func (_ConverterRegistry *ConverterRegistryCallerSession) GetConvertersBySmartTokens(_smartTokens []common.Address) ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertersBySmartTokens(&_ConverterRegistry.CallOpts, _smartTokens)
}

// GetConvertibleToken is a free data retrieval call binding the contract method 0x865cf194.
//
// Solidity: function getConvertibleToken(uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistryCaller) GetConvertibleToken(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getConvertibleToken", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConvertibleToken is a free data retrieval call binding the contract method 0x865cf194.
//
// Solidity: function getConvertibleToken(uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistrySession) GetConvertibleToken(_index *big.Int) (common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertibleToken(&_ConverterRegistry.CallOpts, _index)
}

// GetConvertibleToken is a free data retrieval call binding the contract method 0x865cf194.
//
// Solidity: function getConvertibleToken(uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistryCallerSession) GetConvertibleToken(_index *big.Int) (common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertibleToken(&_ConverterRegistry.CallOpts, _index)
}

// GetConvertibleTokenAnchor is a free data retrieval call binding the contract method 0x603f51e4.
//
// Solidity: function getConvertibleTokenAnchor(address _convertibleToken, uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistryCaller) GetConvertibleTokenAnchor(opts *bind.CallOpts, _convertibleToken common.Address, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getConvertibleTokenAnchor", _convertibleToken, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConvertibleTokenAnchor is a free data retrieval call binding the contract method 0x603f51e4.
//
// Solidity: function getConvertibleTokenAnchor(address _convertibleToken, uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistrySession) GetConvertibleTokenAnchor(_convertibleToken common.Address, _index *big.Int) (common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokenAnchor(&_ConverterRegistry.CallOpts, _convertibleToken, _index)
}

// GetConvertibleTokenAnchor is a free data retrieval call binding the contract method 0x603f51e4.
//
// Solidity: function getConvertibleTokenAnchor(address _convertibleToken, uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistryCallerSession) GetConvertibleTokenAnchor(_convertibleToken common.Address, _index *big.Int) (common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokenAnchor(&_ConverterRegistry.CallOpts, _convertibleToken, _index)
}

// GetConvertibleTokenAnchorCount is a free data retrieval call binding the contract method 0x49c5f32b.
//
// Solidity: function getConvertibleTokenAnchorCount(address _convertibleToken) view returns(uint256)
func (_ConverterRegistry *ConverterRegistryCaller) GetConvertibleTokenAnchorCount(opts *bind.CallOpts, _convertibleToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getConvertibleTokenAnchorCount", _convertibleToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConvertibleTokenAnchorCount is a free data retrieval call binding the contract method 0x49c5f32b.
//
// Solidity: function getConvertibleTokenAnchorCount(address _convertibleToken) view returns(uint256)
func (_ConverterRegistry *ConverterRegistrySession) GetConvertibleTokenAnchorCount(_convertibleToken common.Address) (*big.Int, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokenAnchorCount(&_ConverterRegistry.CallOpts, _convertibleToken)
}

// GetConvertibleTokenAnchorCount is a free data retrieval call binding the contract method 0x49c5f32b.
//
// Solidity: function getConvertibleTokenAnchorCount(address _convertibleToken) view returns(uint256)
func (_ConverterRegistry *ConverterRegistryCallerSession) GetConvertibleTokenAnchorCount(_convertibleToken common.Address) (*big.Int, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokenAnchorCount(&_ConverterRegistry.CallOpts, _convertibleToken)
}

// GetConvertibleTokenAnchors is a free data retrieval call binding the contract method 0x11839064.
//
// Solidity: function getConvertibleTokenAnchors(address _convertibleToken) view returns(address[])
func (_ConverterRegistry *ConverterRegistryCaller) GetConvertibleTokenAnchors(opts *bind.CallOpts, _convertibleToken common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getConvertibleTokenAnchors", _convertibleToken)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetConvertibleTokenAnchors is a free data retrieval call binding the contract method 0x11839064.
//
// Solidity: function getConvertibleTokenAnchors(address _convertibleToken) view returns(address[])
func (_ConverterRegistry *ConverterRegistrySession) GetConvertibleTokenAnchors(_convertibleToken common.Address) ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokenAnchors(&_ConverterRegistry.CallOpts, _convertibleToken)
}

// GetConvertibleTokenAnchors is a free data retrieval call binding the contract method 0x11839064.
//
// Solidity: function getConvertibleTokenAnchors(address _convertibleToken) view returns(address[])
func (_ConverterRegistry *ConverterRegistryCallerSession) GetConvertibleTokenAnchors(_convertibleToken common.Address) ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokenAnchors(&_ConverterRegistry.CallOpts, _convertibleToken)
}

// GetConvertibleTokenCount is a free data retrieval call binding the contract method 0x69be4784.
//
// Solidity: function getConvertibleTokenCount() view returns(uint256)
func (_ConverterRegistry *ConverterRegistryCaller) GetConvertibleTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getConvertibleTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConvertibleTokenCount is a free data retrieval call binding the contract method 0x69be4784.
//
// Solidity: function getConvertibleTokenCount() view returns(uint256)
func (_ConverterRegistry *ConverterRegistrySession) GetConvertibleTokenCount() (*big.Int, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokenCount(&_ConverterRegistry.CallOpts)
}

// GetConvertibleTokenCount is a free data retrieval call binding the contract method 0x69be4784.
//
// Solidity: function getConvertibleTokenCount() view returns(uint256)
func (_ConverterRegistry *ConverterRegistryCallerSession) GetConvertibleTokenCount() (*big.Int, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokenCount(&_ConverterRegistry.CallOpts)
}

// GetConvertibleTokenSmartToken is a free data retrieval call binding the contract method 0xd6c4b5b2.
//
// Solidity: function getConvertibleTokenSmartToken(address _convertibleToken, uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistryCaller) GetConvertibleTokenSmartToken(opts *bind.CallOpts, _convertibleToken common.Address, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getConvertibleTokenSmartToken", _convertibleToken, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConvertibleTokenSmartToken is a free data retrieval call binding the contract method 0xd6c4b5b2.
//
// Solidity: function getConvertibleTokenSmartToken(address _convertibleToken, uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistrySession) GetConvertibleTokenSmartToken(_convertibleToken common.Address, _index *big.Int) (common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokenSmartToken(&_ConverterRegistry.CallOpts, _convertibleToken, _index)
}

// GetConvertibleTokenSmartToken is a free data retrieval call binding the contract method 0xd6c4b5b2.
//
// Solidity: function getConvertibleTokenSmartToken(address _convertibleToken, uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistryCallerSession) GetConvertibleTokenSmartToken(_convertibleToken common.Address, _index *big.Int) (common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokenSmartToken(&_ConverterRegistry.CallOpts, _convertibleToken, _index)
}

// GetConvertibleTokenSmartTokenCount is a free data retrieval call binding the contract method 0xa43d5e94.
//
// Solidity: function getConvertibleTokenSmartTokenCount(address _convertibleToken) view returns(uint256)
func (_ConverterRegistry *ConverterRegistryCaller) GetConvertibleTokenSmartTokenCount(opts *bind.CallOpts, _convertibleToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getConvertibleTokenSmartTokenCount", _convertibleToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConvertibleTokenSmartTokenCount is a free data retrieval call binding the contract method 0xa43d5e94.
//
// Solidity: function getConvertibleTokenSmartTokenCount(address _convertibleToken) view returns(uint256)
func (_ConverterRegistry *ConverterRegistrySession) GetConvertibleTokenSmartTokenCount(_convertibleToken common.Address) (*big.Int, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokenSmartTokenCount(&_ConverterRegistry.CallOpts, _convertibleToken)
}

// GetConvertibleTokenSmartTokenCount is a free data retrieval call binding the contract method 0xa43d5e94.
//
// Solidity: function getConvertibleTokenSmartTokenCount(address _convertibleToken) view returns(uint256)
func (_ConverterRegistry *ConverterRegistryCallerSession) GetConvertibleTokenSmartTokenCount(_convertibleToken common.Address) (*big.Int, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokenSmartTokenCount(&_ConverterRegistry.CallOpts, _convertibleToken)
}

// GetConvertibleTokenSmartTokens is a free data retrieval call binding the contract method 0xf4fb86c0.
//
// Solidity: function getConvertibleTokenSmartTokens(address _convertibleToken) view returns(address[])
func (_ConverterRegistry *ConverterRegistryCaller) GetConvertibleTokenSmartTokens(opts *bind.CallOpts, _convertibleToken common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getConvertibleTokenSmartTokens", _convertibleToken)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetConvertibleTokenSmartTokens is a free data retrieval call binding the contract method 0xf4fb86c0.
//
// Solidity: function getConvertibleTokenSmartTokens(address _convertibleToken) view returns(address[])
func (_ConverterRegistry *ConverterRegistrySession) GetConvertibleTokenSmartTokens(_convertibleToken common.Address) ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokenSmartTokens(&_ConverterRegistry.CallOpts, _convertibleToken)
}

// GetConvertibleTokenSmartTokens is a free data retrieval call binding the contract method 0xf4fb86c0.
//
// Solidity: function getConvertibleTokenSmartTokens(address _convertibleToken) view returns(address[])
func (_ConverterRegistry *ConverterRegistryCallerSession) GetConvertibleTokenSmartTokens(_convertibleToken common.Address) ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokenSmartTokens(&_ConverterRegistry.CallOpts, _convertibleToken)
}

// GetConvertibleTokens is a free data retrieval call binding the contract method 0x5f1b50fe.
//
// Solidity: function getConvertibleTokens() view returns(address[])
func (_ConverterRegistry *ConverterRegistryCaller) GetConvertibleTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getConvertibleTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetConvertibleTokens is a free data retrieval call binding the contract method 0x5f1b50fe.
//
// Solidity: function getConvertibleTokens() view returns(address[])
func (_ConverterRegistry *ConverterRegistrySession) GetConvertibleTokens() ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokens(&_ConverterRegistry.CallOpts)
}

// GetConvertibleTokens is a free data retrieval call binding the contract method 0x5f1b50fe.
//
// Solidity: function getConvertibleTokens() view returns(address[])
func (_ConverterRegistry *ConverterRegistryCallerSession) GetConvertibleTokens() ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetConvertibleTokens(&_ConverterRegistry.CallOpts)
}

// GetLiquidityPool is a free data retrieval call binding the contract method 0xa74498aa.
//
// Solidity: function getLiquidityPool(uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistryCaller) GetLiquidityPool(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getLiquidityPool", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLiquidityPool is a free data retrieval call binding the contract method 0xa74498aa.
//
// Solidity: function getLiquidityPool(uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistrySession) GetLiquidityPool(_index *big.Int) (common.Address, error) {
	return _ConverterRegistry.Contract.GetLiquidityPool(&_ConverterRegistry.CallOpts, _index)
}

// GetLiquidityPool is a free data retrieval call binding the contract method 0xa74498aa.
//
// Solidity: function getLiquidityPool(uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistryCallerSession) GetLiquidityPool(_index *big.Int) (common.Address, error) {
	return _ConverterRegistry.Contract.GetLiquidityPool(&_ConverterRegistry.CallOpts, _index)
}

// GetLiquidityPoolByConfig is a free data retrieval call binding the contract method 0x1d3fccd5.
//
// Solidity: function getLiquidityPoolByConfig(uint16 _type, address[] _reserveTokens, uint32[] _reserveWeights) view returns(address)
func (_ConverterRegistry *ConverterRegistryCaller) GetLiquidityPoolByConfig(opts *bind.CallOpts, _type uint16, _reserveTokens []common.Address, _reserveWeights []uint32) (common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getLiquidityPoolByConfig", _type, _reserveTokens, _reserveWeights)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLiquidityPoolByConfig is a free data retrieval call binding the contract method 0x1d3fccd5.
//
// Solidity: function getLiquidityPoolByConfig(uint16 _type, address[] _reserveTokens, uint32[] _reserveWeights) view returns(address)
func (_ConverterRegistry *ConverterRegistrySession) GetLiquidityPoolByConfig(_type uint16, _reserveTokens []common.Address, _reserveWeights []uint32) (common.Address, error) {
	return _ConverterRegistry.Contract.GetLiquidityPoolByConfig(&_ConverterRegistry.CallOpts, _type, _reserveTokens, _reserveWeights)
}

// GetLiquidityPoolByConfig is a free data retrieval call binding the contract method 0x1d3fccd5.
//
// Solidity: function getLiquidityPoolByConfig(uint16 _type, address[] _reserveTokens, uint32[] _reserveWeights) view returns(address)
func (_ConverterRegistry *ConverterRegistryCallerSession) GetLiquidityPoolByConfig(_type uint16, _reserveTokens []common.Address, _reserveWeights []uint32) (common.Address, error) {
	return _ConverterRegistry.Contract.GetLiquidityPoolByConfig(&_ConverterRegistry.CallOpts, _type, _reserveTokens, _reserveWeights)
}

// GetLiquidityPoolByReserveConfig is a free data retrieval call binding the contract method 0xc22b82f0.
//
// Solidity: function getLiquidityPoolByReserveConfig(address[] _reserveTokens, uint32[] _reserveWeights) view returns(address)
func (_ConverterRegistry *ConverterRegistryCaller) GetLiquidityPoolByReserveConfig(opts *bind.CallOpts, _reserveTokens []common.Address, _reserveWeights []uint32) (common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getLiquidityPoolByReserveConfig", _reserveTokens, _reserveWeights)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLiquidityPoolByReserveConfig is a free data retrieval call binding the contract method 0xc22b82f0.
//
// Solidity: function getLiquidityPoolByReserveConfig(address[] _reserveTokens, uint32[] _reserveWeights) view returns(address)
func (_ConverterRegistry *ConverterRegistrySession) GetLiquidityPoolByReserveConfig(_reserveTokens []common.Address, _reserveWeights []uint32) (common.Address, error) {
	return _ConverterRegistry.Contract.GetLiquidityPoolByReserveConfig(&_ConverterRegistry.CallOpts, _reserveTokens, _reserveWeights)
}

// GetLiquidityPoolByReserveConfig is a free data retrieval call binding the contract method 0xc22b82f0.
//
// Solidity: function getLiquidityPoolByReserveConfig(address[] _reserveTokens, uint32[] _reserveWeights) view returns(address)
func (_ConverterRegistry *ConverterRegistryCallerSession) GetLiquidityPoolByReserveConfig(_reserveTokens []common.Address, _reserveWeights []uint32) (common.Address, error) {
	return _ConverterRegistry.Contract.GetLiquidityPoolByReserveConfig(&_ConverterRegistry.CallOpts, _reserveTokens, _reserveWeights)
}

// GetLiquidityPoolCount is a free data retrieval call binding the contract method 0x7a5f0ffd.
//
// Solidity: function getLiquidityPoolCount() view returns(uint256)
func (_ConverterRegistry *ConverterRegistryCaller) GetLiquidityPoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getLiquidityPoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidityPoolCount is a free data retrieval call binding the contract method 0x7a5f0ffd.
//
// Solidity: function getLiquidityPoolCount() view returns(uint256)
func (_ConverterRegistry *ConverterRegistrySession) GetLiquidityPoolCount() (*big.Int, error) {
	return _ConverterRegistry.Contract.GetLiquidityPoolCount(&_ConverterRegistry.CallOpts)
}

// GetLiquidityPoolCount is a free data retrieval call binding the contract method 0x7a5f0ffd.
//
// Solidity: function getLiquidityPoolCount() view returns(uint256)
func (_ConverterRegistry *ConverterRegistryCallerSession) GetLiquidityPoolCount() (*big.Int, error) {
	return _ConverterRegistry.Contract.GetLiquidityPoolCount(&_ConverterRegistry.CallOpts)
}

// GetLiquidityPools is a free data retrieval call binding the contract method 0x7f45c4c3.
//
// Solidity: function getLiquidityPools() view returns(address[])
func (_ConverterRegistry *ConverterRegistryCaller) GetLiquidityPools(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getLiquidityPools")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetLiquidityPools is a free data retrieval call binding the contract method 0x7f45c4c3.
//
// Solidity: function getLiquidityPools() view returns(address[])
func (_ConverterRegistry *ConverterRegistrySession) GetLiquidityPools() ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetLiquidityPools(&_ConverterRegistry.CallOpts)
}

// GetLiquidityPools is a free data retrieval call binding the contract method 0x7f45c4c3.
//
// Solidity: function getLiquidityPools() view returns(address[])
func (_ConverterRegistry *ConverterRegistryCallerSession) GetLiquidityPools() ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetLiquidityPools(&_ConverterRegistry.CallOpts)
}

// GetSmartToken is a free data retrieval call binding the contract method 0xa109d214.
//
// Solidity: function getSmartToken(uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistryCaller) GetSmartToken(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getSmartToken", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSmartToken is a free data retrieval call binding the contract method 0xa109d214.
//
// Solidity: function getSmartToken(uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistrySession) GetSmartToken(_index *big.Int) (common.Address, error) {
	return _ConverterRegistry.Contract.GetSmartToken(&_ConverterRegistry.CallOpts, _index)
}

// GetSmartToken is a free data retrieval call binding the contract method 0xa109d214.
//
// Solidity: function getSmartToken(uint256 _index) view returns(address)
func (_ConverterRegistry *ConverterRegistryCallerSession) GetSmartToken(_index *big.Int) (common.Address, error) {
	return _ConverterRegistry.Contract.GetSmartToken(&_ConverterRegistry.CallOpts, _index)
}

// GetSmartTokenCount is a free data retrieval call binding the contract method 0xe571049b.
//
// Solidity: function getSmartTokenCount() view returns(uint256)
func (_ConverterRegistry *ConverterRegistryCaller) GetSmartTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getSmartTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSmartTokenCount is a free data retrieval call binding the contract method 0xe571049b.
//
// Solidity: function getSmartTokenCount() view returns(uint256)
func (_ConverterRegistry *ConverterRegistrySession) GetSmartTokenCount() (*big.Int, error) {
	return _ConverterRegistry.Contract.GetSmartTokenCount(&_ConverterRegistry.CallOpts)
}

// GetSmartTokenCount is a free data retrieval call binding the contract method 0xe571049b.
//
// Solidity: function getSmartTokenCount() view returns(uint256)
func (_ConverterRegistry *ConverterRegistryCallerSession) GetSmartTokenCount() (*big.Int, error) {
	return _ConverterRegistry.Contract.GetSmartTokenCount(&_ConverterRegistry.CallOpts)
}

// GetSmartTokens is a free data retrieval call binding the contract method 0x04ceaf41.
//
// Solidity: function getSmartTokens() view returns(address[])
func (_ConverterRegistry *ConverterRegistryCaller) GetSmartTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "getSmartTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSmartTokens is a free data retrieval call binding the contract method 0x04ceaf41.
//
// Solidity: function getSmartTokens() view returns(address[])
func (_ConverterRegistry *ConverterRegistrySession) GetSmartTokens() ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetSmartTokens(&_ConverterRegistry.CallOpts)
}

// GetSmartTokens is a free data retrieval call binding the contract method 0x04ceaf41.
//
// Solidity: function getSmartTokens() view returns(address[])
func (_ConverterRegistry *ConverterRegistryCallerSession) GetSmartTokens() ([]common.Address, error) {
	return _ConverterRegistry.Contract.GetSmartTokens(&_ConverterRegistry.CallOpts)
}

// IsAnchor is a free data retrieval call binding the contract method 0xd8cced2a.
//
// Solidity: function isAnchor(address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCaller) IsAnchor(opts *bind.CallOpts, _value common.Address) (bool, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "isAnchor", _value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAnchor is a free data retrieval call binding the contract method 0xd8cced2a.
//
// Solidity: function isAnchor(address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistrySession) IsAnchor(_value common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsAnchor(&_ConverterRegistry.CallOpts, _value)
}

// IsAnchor is a free data retrieval call binding the contract method 0xd8cced2a.
//
// Solidity: function isAnchor(address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCallerSession) IsAnchor(_value common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsAnchor(&_ConverterRegistry.CallOpts, _value)
}

// IsConverterValid is a free data retrieval call binding the contract method 0x954254f5.
//
// Solidity: function isConverterValid(address _converter) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCaller) IsConverterValid(opts *bind.CallOpts, _converter common.Address) (bool, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "isConverterValid", _converter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConverterValid is a free data retrieval call binding the contract method 0x954254f5.
//
// Solidity: function isConverterValid(address _converter) view returns(bool)
func (_ConverterRegistry *ConverterRegistrySession) IsConverterValid(_converter common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsConverterValid(&_ConverterRegistry.CallOpts, _converter)
}

// IsConverterValid is a free data retrieval call binding the contract method 0x954254f5.
//
// Solidity: function isConverterValid(address _converter) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCallerSession) IsConverterValid(_converter common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsConverterValid(&_ConverterRegistry.CallOpts, _converter)
}

// IsConvertibleToken is a free data retrieval call binding the contract method 0x3ab8857c.
//
// Solidity: function isConvertibleToken(address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCaller) IsConvertibleToken(opts *bind.CallOpts, _value common.Address) (bool, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "isConvertibleToken", _value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConvertibleToken is a free data retrieval call binding the contract method 0x3ab8857c.
//
// Solidity: function isConvertibleToken(address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistrySession) IsConvertibleToken(_value common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsConvertibleToken(&_ConverterRegistry.CallOpts, _value)
}

// IsConvertibleToken is a free data retrieval call binding the contract method 0x3ab8857c.
//
// Solidity: function isConvertibleToken(address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCallerSession) IsConvertibleToken(_value common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsConvertibleToken(&_ConverterRegistry.CallOpts, _value)
}

// IsConvertibleTokenAnchor is a free data retrieval call binding the contract method 0xb4c4197a.
//
// Solidity: function isConvertibleTokenAnchor(address _convertibleToken, address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCaller) IsConvertibleTokenAnchor(opts *bind.CallOpts, _convertibleToken common.Address, _value common.Address) (bool, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "isConvertibleTokenAnchor", _convertibleToken, _value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConvertibleTokenAnchor is a free data retrieval call binding the contract method 0xb4c4197a.
//
// Solidity: function isConvertibleTokenAnchor(address _convertibleToken, address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistrySession) IsConvertibleTokenAnchor(_convertibleToken common.Address, _value common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsConvertibleTokenAnchor(&_ConverterRegistry.CallOpts, _convertibleToken, _value)
}

// IsConvertibleTokenAnchor is a free data retrieval call binding the contract method 0xb4c4197a.
//
// Solidity: function isConvertibleTokenAnchor(address _convertibleToken, address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCallerSession) IsConvertibleTokenAnchor(_convertibleToken common.Address, _value common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsConvertibleTokenAnchor(&_ConverterRegistry.CallOpts, _convertibleToken, _value)
}

// IsConvertibleTokenSmartToken is a free data retrieval call binding the contract method 0x725b8786.
//
// Solidity: function isConvertibleTokenSmartToken(address _convertibleToken, address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCaller) IsConvertibleTokenSmartToken(opts *bind.CallOpts, _convertibleToken common.Address, _value common.Address) (bool, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "isConvertibleTokenSmartToken", _convertibleToken, _value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConvertibleTokenSmartToken is a free data retrieval call binding the contract method 0x725b8786.
//
// Solidity: function isConvertibleTokenSmartToken(address _convertibleToken, address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistrySession) IsConvertibleTokenSmartToken(_convertibleToken common.Address, _value common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsConvertibleTokenSmartToken(&_ConverterRegistry.CallOpts, _convertibleToken, _value)
}

// IsConvertibleTokenSmartToken is a free data retrieval call binding the contract method 0x725b8786.
//
// Solidity: function isConvertibleTokenSmartToken(address _convertibleToken, address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCallerSession) IsConvertibleTokenSmartToken(_convertibleToken common.Address, _value common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsConvertibleTokenSmartToken(&_ConverterRegistry.CallOpts, _convertibleToken, _value)
}

// IsLiquidityPool is a free data retrieval call binding the contract method 0xe85455d7.
//
// Solidity: function isLiquidityPool(address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCaller) IsLiquidityPool(opts *bind.CallOpts, _value common.Address) (bool, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "isLiquidityPool", _value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidityPool is a free data retrieval call binding the contract method 0xe85455d7.
//
// Solidity: function isLiquidityPool(address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistrySession) IsLiquidityPool(_value common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsLiquidityPool(&_ConverterRegistry.CallOpts, _value)
}

// IsLiquidityPool is a free data retrieval call binding the contract method 0xe85455d7.
//
// Solidity: function isLiquidityPool(address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCallerSession) IsLiquidityPool(_value common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsLiquidityPool(&_ConverterRegistry.CallOpts, _value)
}

// IsSimilarLiquidityPoolRegistered is a free data retrieval call binding the contract method 0x8f1d0e1a.
//
// Solidity: function isSimilarLiquidityPoolRegistered(address _converter) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCaller) IsSimilarLiquidityPoolRegistered(opts *bind.CallOpts, _converter common.Address) (bool, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "isSimilarLiquidityPoolRegistered", _converter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSimilarLiquidityPoolRegistered is a free data retrieval call binding the contract method 0x8f1d0e1a.
//
// Solidity: function isSimilarLiquidityPoolRegistered(address _converter) view returns(bool)
func (_ConverterRegistry *ConverterRegistrySession) IsSimilarLiquidityPoolRegistered(_converter common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsSimilarLiquidityPoolRegistered(&_ConverterRegistry.CallOpts, _converter)
}

// IsSimilarLiquidityPoolRegistered is a free data retrieval call binding the contract method 0x8f1d0e1a.
//
// Solidity: function isSimilarLiquidityPoolRegistered(address _converter) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCallerSession) IsSimilarLiquidityPoolRegistered(_converter common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsSimilarLiquidityPoolRegistered(&_ConverterRegistry.CallOpts, _converter)
}

// IsSmartToken is a free data retrieval call binding the contract method 0x4123ef60.
//
// Solidity: function isSmartToken(address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCaller) IsSmartToken(opts *bind.CallOpts, _value common.Address) (bool, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "isSmartToken", _value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSmartToken is a free data retrieval call binding the contract method 0x4123ef60.
//
// Solidity: function isSmartToken(address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistrySession) IsSmartToken(_value common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsSmartToken(&_ConverterRegistry.CallOpts, _value)
}

// IsSmartToken is a free data retrieval call binding the contract method 0x4123ef60.
//
// Solidity: function isSmartToken(address _value) view returns(bool)
func (_ConverterRegistry *ConverterRegistryCallerSession) IsSmartToken(_value common.Address) (bool, error) {
	return _ConverterRegistry.Contract.IsSmartToken(&_ConverterRegistry.CallOpts, _value)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ConverterRegistry *ConverterRegistryCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "newOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ConverterRegistry *ConverterRegistrySession) NewOwner() (common.Address, error) {
	return _ConverterRegistry.Contract.NewOwner(&_ConverterRegistry.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ConverterRegistry *ConverterRegistryCallerSession) NewOwner() (common.Address, error) {
	return _ConverterRegistry.Contract.NewOwner(&_ConverterRegistry.CallOpts)
}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_ConverterRegistry *ConverterRegistryCaller) OnlyOwnerCanUpdateRegistry(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "onlyOwnerCanUpdateRegistry")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_ConverterRegistry *ConverterRegistrySession) OnlyOwnerCanUpdateRegistry() (bool, error) {
	return _ConverterRegistry.Contract.OnlyOwnerCanUpdateRegistry(&_ConverterRegistry.CallOpts)
}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_ConverterRegistry *ConverterRegistryCallerSession) OnlyOwnerCanUpdateRegistry() (bool, error) {
	return _ConverterRegistry.Contract.OnlyOwnerCanUpdateRegistry(&_ConverterRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConverterRegistry *ConverterRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConverterRegistry *ConverterRegistrySession) Owner() (common.Address, error) {
	return _ConverterRegistry.Contract.Owner(&_ConverterRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConverterRegistry *ConverterRegistryCallerSession) Owner() (common.Address, error) {
	return _ConverterRegistry.Contract.Owner(&_ConverterRegistry.CallOpts)
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_ConverterRegistry *ConverterRegistryCaller) PrevRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "prevRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_ConverterRegistry *ConverterRegistrySession) PrevRegistry() (common.Address, error) {
	return _ConverterRegistry.Contract.PrevRegistry(&_ConverterRegistry.CallOpts)
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_ConverterRegistry *ConverterRegistryCallerSession) PrevRegistry() (common.Address, error) {
	return _ConverterRegistry.Contract.PrevRegistry(&_ConverterRegistry.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConverterRegistry *ConverterRegistryCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterRegistry.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConverterRegistry *ConverterRegistrySession) Registry() (common.Address, error) {
	return _ConverterRegistry.Contract.Registry(&_ConverterRegistry.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConverterRegistry *ConverterRegistryCallerSession) Registry() (common.Address, error) {
	return _ConverterRegistry.Contract.Registry(&_ConverterRegistry.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ConverterRegistry *ConverterRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterRegistry.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ConverterRegistry *ConverterRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _ConverterRegistry.Contract.AcceptOwnership(&_ConverterRegistry.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ConverterRegistry *ConverterRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ConverterRegistry.Contract.AcceptOwnership(&_ConverterRegistry.TransactOpts)
}

// AddConverter is a paid mutator transaction binding the contract method 0x6ce1c4dc.
//
// Solidity: function addConverter(address _converter) returns()
func (_ConverterRegistry *ConverterRegistryTransactor) AddConverter(opts *bind.TransactOpts, _converter common.Address) (*types.Transaction, error) {
	return _ConverterRegistry.contract.Transact(opts, "addConverter", _converter)
}

// AddConverter is a paid mutator transaction binding the contract method 0x6ce1c4dc.
//
// Solidity: function addConverter(address _converter) returns()
func (_ConverterRegistry *ConverterRegistrySession) AddConverter(_converter common.Address) (*types.Transaction, error) {
	return _ConverterRegistry.Contract.AddConverter(&_ConverterRegistry.TransactOpts, _converter)
}

// AddConverter is a paid mutator transaction binding the contract method 0x6ce1c4dc.
//
// Solidity: function addConverter(address _converter) returns()
func (_ConverterRegistry *ConverterRegistryTransactorSession) AddConverter(_converter common.Address) (*types.Transaction, error) {
	return _ConverterRegistry.Contract.AddConverter(&_ConverterRegistry.TransactOpts, _converter)
}

// NewConverter is a paid mutator transaction binding the contract method 0x5a0a6618.
//
// Solidity: function newConverter(uint16 _type, string _name, string _symbol, uint8 _decimals, uint32 _maxConversionFee, address[] _reserveTokens, uint32[] _reserveWeights) returns(address)
func (_ConverterRegistry *ConverterRegistryTransactor) NewConverter(opts *bind.TransactOpts, _type uint16, _name string, _symbol string, _decimals uint8, _maxConversionFee uint32, _reserveTokens []common.Address, _reserveWeights []uint32) (*types.Transaction, error) {
	return _ConverterRegistry.contract.Transact(opts, "newConverter", _type, _name, _symbol, _decimals, _maxConversionFee, _reserveTokens, _reserveWeights)
}

// NewConverter is a paid mutator transaction binding the contract method 0x5a0a6618.
//
// Solidity: function newConverter(uint16 _type, string _name, string _symbol, uint8 _decimals, uint32 _maxConversionFee, address[] _reserveTokens, uint32[] _reserveWeights) returns(address)
func (_ConverterRegistry *ConverterRegistrySession) NewConverter(_type uint16, _name string, _symbol string, _decimals uint8, _maxConversionFee uint32, _reserveTokens []common.Address, _reserveWeights []uint32) (*types.Transaction, error) {
	return _ConverterRegistry.Contract.NewConverter(&_ConverterRegistry.TransactOpts, _type, _name, _symbol, _decimals, _maxConversionFee, _reserveTokens, _reserveWeights)
}

// NewConverter is a paid mutator transaction binding the contract method 0x5a0a6618.
//
// Solidity: function newConverter(uint16 _type, string _name, string _symbol, uint8 _decimals, uint32 _maxConversionFee, address[] _reserveTokens, uint32[] _reserveWeights) returns(address)
func (_ConverterRegistry *ConverterRegistryTransactorSession) NewConverter(_type uint16, _name string, _symbol string, _decimals uint8, _maxConversionFee uint32, _reserveTokens []common.Address, _reserveWeights []uint32) (*types.Transaction, error) {
	return _ConverterRegistry.Contract.NewConverter(&_ConverterRegistry.TransactOpts, _type, _name, _symbol, _decimals, _maxConversionFee, _reserveTokens, _reserveWeights)
}

// RemoveConverter is a paid mutator transaction binding the contract method 0x9e76a007.
//
// Solidity: function removeConverter(address _converter) returns()
func (_ConverterRegistry *ConverterRegistryTransactor) RemoveConverter(opts *bind.TransactOpts, _converter common.Address) (*types.Transaction, error) {
	return _ConverterRegistry.contract.Transact(opts, "removeConverter", _converter)
}

// RemoveConverter is a paid mutator transaction binding the contract method 0x9e76a007.
//
// Solidity: function removeConverter(address _converter) returns()
func (_ConverterRegistry *ConverterRegistrySession) RemoveConverter(_converter common.Address) (*types.Transaction, error) {
	return _ConverterRegistry.Contract.RemoveConverter(&_ConverterRegistry.TransactOpts, _converter)
}

// RemoveConverter is a paid mutator transaction binding the contract method 0x9e76a007.
//
// Solidity: function removeConverter(address _converter) returns()
func (_ConverterRegistry *ConverterRegistryTransactorSession) RemoveConverter(_converter common.Address) (*types.Transaction, error) {
	return _ConverterRegistry.Contract.RemoveConverter(&_ConverterRegistry.TransactOpts, _converter)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_ConverterRegistry *ConverterRegistryTransactor) RestoreRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterRegistry.contract.Transact(opts, "restoreRegistry")
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_ConverterRegistry *ConverterRegistrySession) RestoreRegistry() (*types.Transaction, error) {
	return _ConverterRegistry.Contract.RestoreRegistry(&_ConverterRegistry.TransactOpts)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_ConverterRegistry *ConverterRegistryTransactorSession) RestoreRegistry() (*types.Transaction, error) {
	return _ConverterRegistry.Contract.RestoreRegistry(&_ConverterRegistry.TransactOpts)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_ConverterRegistry *ConverterRegistryTransactor) RestrictRegistryUpdate(opts *bind.TransactOpts, _onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _ConverterRegistry.contract.Transact(opts, "restrictRegistryUpdate", _onlyOwnerCanUpdateRegistry)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_ConverterRegistry *ConverterRegistrySession) RestrictRegistryUpdate(_onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _ConverterRegistry.Contract.RestrictRegistryUpdate(&_ConverterRegistry.TransactOpts, _onlyOwnerCanUpdateRegistry)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_ConverterRegistry *ConverterRegistryTransactorSession) RestrictRegistryUpdate(_onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _ConverterRegistry.Contract.RestrictRegistryUpdate(&_ConverterRegistry.TransactOpts, _onlyOwnerCanUpdateRegistry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ConverterRegistry *ConverterRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ConverterRegistry.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ConverterRegistry *ConverterRegistrySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterRegistry.Contract.TransferOwnership(&_ConverterRegistry.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ConverterRegistry *ConverterRegistryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterRegistry.Contract.TransferOwnership(&_ConverterRegistry.TransactOpts, _newOwner)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_ConverterRegistry *ConverterRegistryTransactor) UpdateRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterRegistry.contract.Transact(opts, "updateRegistry")
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_ConverterRegistry *ConverterRegistrySession) UpdateRegistry() (*types.Transaction, error) {
	return _ConverterRegistry.Contract.UpdateRegistry(&_ConverterRegistry.TransactOpts)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_ConverterRegistry *ConverterRegistryTransactorSession) UpdateRegistry() (*types.Transaction, error) {
	return _ConverterRegistry.Contract.UpdateRegistry(&_ConverterRegistry.TransactOpts)
}

// ConverterRegistryConverterAnchorAddedIterator is returned from FilterConverterAnchorAdded and is used to iterate over the raw logs and unpacked data for ConverterAnchorAdded events raised by the ConverterRegistry contract.
type ConverterRegistryConverterAnchorAddedIterator struct {
	Event *ConverterRegistryConverterAnchorAdded // Event containing the contract specifics and raw log

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
func (it *ConverterRegistryConverterAnchorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterRegistryConverterAnchorAdded)
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
		it.Event = new(ConverterRegistryConverterAnchorAdded)
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
func (it *ConverterRegistryConverterAnchorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterRegistryConverterAnchorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterRegistryConverterAnchorAdded represents a ConverterAnchorAdded event raised by the ConverterRegistry contract.
type ConverterRegistryConverterAnchorAdded struct {
	Anchor common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterConverterAnchorAdded is a free log retrieval operation binding the contract event 0xc0a6d303d67b7ed9fa0abae1c48878df32acc0e7ca4334c7dad2bceeee5956fd.
//
// Solidity: event ConverterAnchorAdded(address indexed _anchor)
func (_ConverterRegistry *ConverterRegistryFilterer) FilterConverterAnchorAdded(opts *bind.FilterOpts, _anchor []common.Address) (*ConverterRegistryConverterAnchorAddedIterator, error) {

	var _anchorRule []interface{}
	for _, _anchorItem := range _anchor {
		_anchorRule = append(_anchorRule, _anchorItem)
	}

	logs, sub, err := _ConverterRegistry.contract.FilterLogs(opts, "ConverterAnchorAdded", _anchorRule)
	if err != nil {
		return nil, err
	}
	return &ConverterRegistryConverterAnchorAddedIterator{contract: _ConverterRegistry.contract, event: "ConverterAnchorAdded", logs: logs, sub: sub}, nil
}

// WatchConverterAnchorAdded is a free log subscription operation binding the contract event 0xc0a6d303d67b7ed9fa0abae1c48878df32acc0e7ca4334c7dad2bceeee5956fd.
//
// Solidity: event ConverterAnchorAdded(address indexed _anchor)
func (_ConverterRegistry *ConverterRegistryFilterer) WatchConverterAnchorAdded(opts *bind.WatchOpts, sink chan<- *ConverterRegistryConverterAnchorAdded, _anchor []common.Address) (event.Subscription, error) {

	var _anchorRule []interface{}
	for _, _anchorItem := range _anchor {
		_anchorRule = append(_anchorRule, _anchorItem)
	}

	logs, sub, err := _ConverterRegistry.contract.WatchLogs(opts, "ConverterAnchorAdded", _anchorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterRegistryConverterAnchorAdded)
				if err := _ConverterRegistry.contract.UnpackLog(event, "ConverterAnchorAdded", log); err != nil {
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

// ParseConverterAnchorAdded is a log parse operation binding the contract event 0xc0a6d303d67b7ed9fa0abae1c48878df32acc0e7ca4334c7dad2bceeee5956fd.
//
// Solidity: event ConverterAnchorAdded(address indexed _anchor)
func (_ConverterRegistry *ConverterRegistryFilterer) ParseConverterAnchorAdded(log types.Log) (*ConverterRegistryConverterAnchorAdded, error) {
	event := new(ConverterRegistryConverterAnchorAdded)
	if err := _ConverterRegistry.contract.UnpackLog(event, "ConverterAnchorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterRegistryConverterAnchorRemovedIterator is returned from FilterConverterAnchorRemoved and is used to iterate over the raw logs and unpacked data for ConverterAnchorRemoved events raised by the ConverterRegistry contract.
type ConverterRegistryConverterAnchorRemovedIterator struct {
	Event *ConverterRegistryConverterAnchorRemoved // Event containing the contract specifics and raw log

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
func (it *ConverterRegistryConverterAnchorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterRegistryConverterAnchorRemoved)
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
		it.Event = new(ConverterRegistryConverterAnchorRemoved)
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
func (it *ConverterRegistryConverterAnchorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterRegistryConverterAnchorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterRegistryConverterAnchorRemoved represents a ConverterAnchorRemoved event raised by the ConverterRegistry contract.
type ConverterRegistryConverterAnchorRemoved struct {
	Anchor common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterConverterAnchorRemoved is a free log retrieval operation binding the contract event 0xbfdf1baaa7e4871111360083540f067050014f651c9e4610a2a4a4bdf8bfab5d.
//
// Solidity: event ConverterAnchorRemoved(address indexed _anchor)
func (_ConverterRegistry *ConverterRegistryFilterer) FilterConverterAnchorRemoved(opts *bind.FilterOpts, _anchor []common.Address) (*ConverterRegistryConverterAnchorRemovedIterator, error) {

	var _anchorRule []interface{}
	for _, _anchorItem := range _anchor {
		_anchorRule = append(_anchorRule, _anchorItem)
	}

	logs, sub, err := _ConverterRegistry.contract.FilterLogs(opts, "ConverterAnchorRemoved", _anchorRule)
	if err != nil {
		return nil, err
	}
	return &ConverterRegistryConverterAnchorRemovedIterator{contract: _ConverterRegistry.contract, event: "ConverterAnchorRemoved", logs: logs, sub: sub}, nil
}

// WatchConverterAnchorRemoved is a free log subscription operation binding the contract event 0xbfdf1baaa7e4871111360083540f067050014f651c9e4610a2a4a4bdf8bfab5d.
//
// Solidity: event ConverterAnchorRemoved(address indexed _anchor)
func (_ConverterRegistry *ConverterRegistryFilterer) WatchConverterAnchorRemoved(opts *bind.WatchOpts, sink chan<- *ConverterRegistryConverterAnchorRemoved, _anchor []common.Address) (event.Subscription, error) {

	var _anchorRule []interface{}
	for _, _anchorItem := range _anchor {
		_anchorRule = append(_anchorRule, _anchorItem)
	}

	logs, sub, err := _ConverterRegistry.contract.WatchLogs(opts, "ConverterAnchorRemoved", _anchorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterRegistryConverterAnchorRemoved)
				if err := _ConverterRegistry.contract.UnpackLog(event, "ConverterAnchorRemoved", log); err != nil {
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

// ParseConverterAnchorRemoved is a log parse operation binding the contract event 0xbfdf1baaa7e4871111360083540f067050014f651c9e4610a2a4a4bdf8bfab5d.
//
// Solidity: event ConverterAnchorRemoved(address indexed _anchor)
func (_ConverterRegistry *ConverterRegistryFilterer) ParseConverterAnchorRemoved(log types.Log) (*ConverterRegistryConverterAnchorRemoved, error) {
	event := new(ConverterRegistryConverterAnchorRemoved)
	if err := _ConverterRegistry.contract.UnpackLog(event, "ConverterAnchorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterRegistryConvertibleTokenAddedIterator is returned from FilterConvertibleTokenAdded and is used to iterate over the raw logs and unpacked data for ConvertibleTokenAdded events raised by the ConverterRegistry contract.
type ConverterRegistryConvertibleTokenAddedIterator struct {
	Event *ConverterRegistryConvertibleTokenAdded // Event containing the contract specifics and raw log

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
func (it *ConverterRegistryConvertibleTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterRegistryConvertibleTokenAdded)
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
		it.Event = new(ConverterRegistryConvertibleTokenAdded)
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
func (it *ConverterRegistryConvertibleTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterRegistryConvertibleTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterRegistryConvertibleTokenAdded represents a ConvertibleTokenAdded event raised by the ConverterRegistry contract.
type ConverterRegistryConvertibleTokenAdded struct {
	ConvertibleToken common.Address
	SmartToken       common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterConvertibleTokenAdded is a free log retrieval operation binding the contract event 0xf2e7cf6d6ed3f77039511409a43d4fa5108f09ab71d72b014380364c910233a5.
//
// Solidity: event ConvertibleTokenAdded(address indexed _convertibleToken, address indexed _smartToken)
func (_ConverterRegistry *ConverterRegistryFilterer) FilterConvertibleTokenAdded(opts *bind.FilterOpts, _convertibleToken []common.Address, _smartToken []common.Address) (*ConverterRegistryConvertibleTokenAddedIterator, error) {

	var _convertibleTokenRule []interface{}
	for _, _convertibleTokenItem := range _convertibleToken {
		_convertibleTokenRule = append(_convertibleTokenRule, _convertibleTokenItem)
	}
	var _smartTokenRule []interface{}
	for _, _smartTokenItem := range _smartToken {
		_smartTokenRule = append(_smartTokenRule, _smartTokenItem)
	}

	logs, sub, err := _ConverterRegistry.contract.FilterLogs(opts, "ConvertibleTokenAdded", _convertibleTokenRule, _smartTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConverterRegistryConvertibleTokenAddedIterator{contract: _ConverterRegistry.contract, event: "ConvertibleTokenAdded", logs: logs, sub: sub}, nil
}

// WatchConvertibleTokenAdded is a free log subscription operation binding the contract event 0xf2e7cf6d6ed3f77039511409a43d4fa5108f09ab71d72b014380364c910233a5.
//
// Solidity: event ConvertibleTokenAdded(address indexed _convertibleToken, address indexed _smartToken)
func (_ConverterRegistry *ConverterRegistryFilterer) WatchConvertibleTokenAdded(opts *bind.WatchOpts, sink chan<- *ConverterRegistryConvertibleTokenAdded, _convertibleToken []common.Address, _smartToken []common.Address) (event.Subscription, error) {

	var _convertibleTokenRule []interface{}
	for _, _convertibleTokenItem := range _convertibleToken {
		_convertibleTokenRule = append(_convertibleTokenRule, _convertibleTokenItem)
	}
	var _smartTokenRule []interface{}
	for _, _smartTokenItem := range _smartToken {
		_smartTokenRule = append(_smartTokenRule, _smartTokenItem)
	}

	logs, sub, err := _ConverterRegistry.contract.WatchLogs(opts, "ConvertibleTokenAdded", _convertibleTokenRule, _smartTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterRegistryConvertibleTokenAdded)
				if err := _ConverterRegistry.contract.UnpackLog(event, "ConvertibleTokenAdded", log); err != nil {
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

// ParseConvertibleTokenAdded is a log parse operation binding the contract event 0xf2e7cf6d6ed3f77039511409a43d4fa5108f09ab71d72b014380364c910233a5.
//
// Solidity: event ConvertibleTokenAdded(address indexed _convertibleToken, address indexed _smartToken)
func (_ConverterRegistry *ConverterRegistryFilterer) ParseConvertibleTokenAdded(log types.Log) (*ConverterRegistryConvertibleTokenAdded, error) {
	event := new(ConverterRegistryConvertibleTokenAdded)
	if err := _ConverterRegistry.contract.UnpackLog(event, "ConvertibleTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterRegistryConvertibleTokenRemovedIterator is returned from FilterConvertibleTokenRemoved and is used to iterate over the raw logs and unpacked data for ConvertibleTokenRemoved events raised by the ConverterRegistry contract.
type ConverterRegistryConvertibleTokenRemovedIterator struct {
	Event *ConverterRegistryConvertibleTokenRemoved // Event containing the contract specifics and raw log

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
func (it *ConverterRegistryConvertibleTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterRegistryConvertibleTokenRemoved)
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
		it.Event = new(ConverterRegistryConvertibleTokenRemoved)
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
func (it *ConverterRegistryConvertibleTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterRegistryConvertibleTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterRegistryConvertibleTokenRemoved represents a ConvertibleTokenRemoved event raised by the ConverterRegistry contract.
type ConverterRegistryConvertibleTokenRemoved struct {
	ConvertibleToken common.Address
	SmartToken       common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterConvertibleTokenRemoved is a free log retrieval operation binding the contract event 0x9430ad6ff45d6c3e126c7711bf0036bd9bc6b202fa19628abd88e59cf43ced43.
//
// Solidity: event ConvertibleTokenRemoved(address indexed _convertibleToken, address indexed _smartToken)
func (_ConverterRegistry *ConverterRegistryFilterer) FilterConvertibleTokenRemoved(opts *bind.FilterOpts, _convertibleToken []common.Address, _smartToken []common.Address) (*ConverterRegistryConvertibleTokenRemovedIterator, error) {

	var _convertibleTokenRule []interface{}
	for _, _convertibleTokenItem := range _convertibleToken {
		_convertibleTokenRule = append(_convertibleTokenRule, _convertibleTokenItem)
	}
	var _smartTokenRule []interface{}
	for _, _smartTokenItem := range _smartToken {
		_smartTokenRule = append(_smartTokenRule, _smartTokenItem)
	}

	logs, sub, err := _ConverterRegistry.contract.FilterLogs(opts, "ConvertibleTokenRemoved", _convertibleTokenRule, _smartTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConverterRegistryConvertibleTokenRemovedIterator{contract: _ConverterRegistry.contract, event: "ConvertibleTokenRemoved", logs: logs, sub: sub}, nil
}

// WatchConvertibleTokenRemoved is a free log subscription operation binding the contract event 0x9430ad6ff45d6c3e126c7711bf0036bd9bc6b202fa19628abd88e59cf43ced43.
//
// Solidity: event ConvertibleTokenRemoved(address indexed _convertibleToken, address indexed _smartToken)
func (_ConverterRegistry *ConverterRegistryFilterer) WatchConvertibleTokenRemoved(opts *bind.WatchOpts, sink chan<- *ConverterRegistryConvertibleTokenRemoved, _convertibleToken []common.Address, _smartToken []common.Address) (event.Subscription, error) {

	var _convertibleTokenRule []interface{}
	for _, _convertibleTokenItem := range _convertibleToken {
		_convertibleTokenRule = append(_convertibleTokenRule, _convertibleTokenItem)
	}
	var _smartTokenRule []interface{}
	for _, _smartTokenItem := range _smartToken {
		_smartTokenRule = append(_smartTokenRule, _smartTokenItem)
	}

	logs, sub, err := _ConverterRegistry.contract.WatchLogs(opts, "ConvertibleTokenRemoved", _convertibleTokenRule, _smartTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterRegistryConvertibleTokenRemoved)
				if err := _ConverterRegistry.contract.UnpackLog(event, "ConvertibleTokenRemoved", log); err != nil {
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

// ParseConvertibleTokenRemoved is a log parse operation binding the contract event 0x9430ad6ff45d6c3e126c7711bf0036bd9bc6b202fa19628abd88e59cf43ced43.
//
// Solidity: event ConvertibleTokenRemoved(address indexed _convertibleToken, address indexed _smartToken)
func (_ConverterRegistry *ConverterRegistryFilterer) ParseConvertibleTokenRemoved(log types.Log) (*ConverterRegistryConvertibleTokenRemoved, error) {
	event := new(ConverterRegistryConvertibleTokenRemoved)
	if err := _ConverterRegistry.contract.UnpackLog(event, "ConvertibleTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterRegistryLiquidityPoolAddedIterator is returned from FilterLiquidityPoolAdded and is used to iterate over the raw logs and unpacked data for LiquidityPoolAdded events raised by the ConverterRegistry contract.
type ConverterRegistryLiquidityPoolAddedIterator struct {
	Event *ConverterRegistryLiquidityPoolAdded // Event containing the contract specifics and raw log

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
func (it *ConverterRegistryLiquidityPoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterRegistryLiquidityPoolAdded)
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
		it.Event = new(ConverterRegistryLiquidityPoolAdded)
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
func (it *ConverterRegistryLiquidityPoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterRegistryLiquidityPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterRegistryLiquidityPoolAdded represents a LiquidityPoolAdded event raised by the ConverterRegistry contract.
type ConverterRegistryLiquidityPoolAdded struct {
	LiquidityPool common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLiquidityPoolAdded is a free log retrieval operation binding the contract event 0xb893f883ef734b712208a877459424ee509832c57e0461fb1ac99ed4d42f2d89.
//
// Solidity: event LiquidityPoolAdded(address indexed _liquidityPool)
func (_ConverterRegistry *ConverterRegistryFilterer) FilterLiquidityPoolAdded(opts *bind.FilterOpts, _liquidityPool []common.Address) (*ConverterRegistryLiquidityPoolAddedIterator, error) {

	var _liquidityPoolRule []interface{}
	for _, _liquidityPoolItem := range _liquidityPool {
		_liquidityPoolRule = append(_liquidityPoolRule, _liquidityPoolItem)
	}

	logs, sub, err := _ConverterRegistry.contract.FilterLogs(opts, "LiquidityPoolAdded", _liquidityPoolRule)
	if err != nil {
		return nil, err
	}
	return &ConverterRegistryLiquidityPoolAddedIterator{contract: _ConverterRegistry.contract, event: "LiquidityPoolAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityPoolAdded is a free log subscription operation binding the contract event 0xb893f883ef734b712208a877459424ee509832c57e0461fb1ac99ed4d42f2d89.
//
// Solidity: event LiquidityPoolAdded(address indexed _liquidityPool)
func (_ConverterRegistry *ConverterRegistryFilterer) WatchLiquidityPoolAdded(opts *bind.WatchOpts, sink chan<- *ConverterRegistryLiquidityPoolAdded, _liquidityPool []common.Address) (event.Subscription, error) {

	var _liquidityPoolRule []interface{}
	for _, _liquidityPoolItem := range _liquidityPool {
		_liquidityPoolRule = append(_liquidityPoolRule, _liquidityPoolItem)
	}

	logs, sub, err := _ConverterRegistry.contract.WatchLogs(opts, "LiquidityPoolAdded", _liquidityPoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterRegistryLiquidityPoolAdded)
				if err := _ConverterRegistry.contract.UnpackLog(event, "LiquidityPoolAdded", log); err != nil {
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

// ParseLiquidityPoolAdded is a log parse operation binding the contract event 0xb893f883ef734b712208a877459424ee509832c57e0461fb1ac99ed4d42f2d89.
//
// Solidity: event LiquidityPoolAdded(address indexed _liquidityPool)
func (_ConverterRegistry *ConverterRegistryFilterer) ParseLiquidityPoolAdded(log types.Log) (*ConverterRegistryLiquidityPoolAdded, error) {
	event := new(ConverterRegistryLiquidityPoolAdded)
	if err := _ConverterRegistry.contract.UnpackLog(event, "LiquidityPoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterRegistryLiquidityPoolRemovedIterator is returned from FilterLiquidityPoolRemoved and is used to iterate over the raw logs and unpacked data for LiquidityPoolRemoved events raised by the ConverterRegistry contract.
type ConverterRegistryLiquidityPoolRemovedIterator struct {
	Event *ConverterRegistryLiquidityPoolRemoved // Event containing the contract specifics and raw log

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
func (it *ConverterRegistryLiquidityPoolRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterRegistryLiquidityPoolRemoved)
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
		it.Event = new(ConverterRegistryLiquidityPoolRemoved)
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
func (it *ConverterRegistryLiquidityPoolRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterRegistryLiquidityPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterRegistryLiquidityPoolRemoved represents a LiquidityPoolRemoved event raised by the ConverterRegistry contract.
type ConverterRegistryLiquidityPoolRemoved struct {
	LiquidityPool common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLiquidityPoolRemoved is a free log retrieval operation binding the contract event 0x59c3fbcae88f30e9b0e35c132a7f68c53231dffa4722f197c7ecb0ee013eee60.
//
// Solidity: event LiquidityPoolRemoved(address indexed _liquidityPool)
func (_ConverterRegistry *ConverterRegistryFilterer) FilterLiquidityPoolRemoved(opts *bind.FilterOpts, _liquidityPool []common.Address) (*ConverterRegistryLiquidityPoolRemovedIterator, error) {

	var _liquidityPoolRule []interface{}
	for _, _liquidityPoolItem := range _liquidityPool {
		_liquidityPoolRule = append(_liquidityPoolRule, _liquidityPoolItem)
	}

	logs, sub, err := _ConverterRegistry.contract.FilterLogs(opts, "LiquidityPoolRemoved", _liquidityPoolRule)
	if err != nil {
		return nil, err
	}
	return &ConverterRegistryLiquidityPoolRemovedIterator{contract: _ConverterRegistry.contract, event: "LiquidityPoolRemoved", logs: logs, sub: sub}, nil
}

// WatchLiquidityPoolRemoved is a free log subscription operation binding the contract event 0x59c3fbcae88f30e9b0e35c132a7f68c53231dffa4722f197c7ecb0ee013eee60.
//
// Solidity: event LiquidityPoolRemoved(address indexed _liquidityPool)
func (_ConverterRegistry *ConverterRegistryFilterer) WatchLiquidityPoolRemoved(opts *bind.WatchOpts, sink chan<- *ConverterRegistryLiquidityPoolRemoved, _liquidityPool []common.Address) (event.Subscription, error) {

	var _liquidityPoolRule []interface{}
	for _, _liquidityPoolItem := range _liquidityPool {
		_liquidityPoolRule = append(_liquidityPoolRule, _liquidityPoolItem)
	}

	logs, sub, err := _ConverterRegistry.contract.WatchLogs(opts, "LiquidityPoolRemoved", _liquidityPoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterRegistryLiquidityPoolRemoved)
				if err := _ConverterRegistry.contract.UnpackLog(event, "LiquidityPoolRemoved", log); err != nil {
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

// ParseLiquidityPoolRemoved is a log parse operation binding the contract event 0x59c3fbcae88f30e9b0e35c132a7f68c53231dffa4722f197c7ecb0ee013eee60.
//
// Solidity: event LiquidityPoolRemoved(address indexed _liquidityPool)
func (_ConverterRegistry *ConverterRegistryFilterer) ParseLiquidityPoolRemoved(log types.Log) (*ConverterRegistryLiquidityPoolRemoved, error) {
	event := new(ConverterRegistryLiquidityPoolRemoved)
	if err := _ConverterRegistry.contract.UnpackLog(event, "LiquidityPoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterRegistryOwnerUpdateIterator is returned from FilterOwnerUpdate and is used to iterate over the raw logs and unpacked data for OwnerUpdate events raised by the ConverterRegistry contract.
type ConverterRegistryOwnerUpdateIterator struct {
	Event *ConverterRegistryOwnerUpdate // Event containing the contract specifics and raw log

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
func (it *ConverterRegistryOwnerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterRegistryOwnerUpdate)
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
		it.Event = new(ConverterRegistryOwnerUpdate)
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
func (it *ConverterRegistryOwnerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterRegistryOwnerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterRegistryOwnerUpdate represents a OwnerUpdate event raised by the ConverterRegistry contract.
type ConverterRegistryOwnerUpdate struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdate is a free log retrieval operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_ConverterRegistry *ConverterRegistryFilterer) FilterOwnerUpdate(opts *bind.FilterOpts, _prevOwner []common.Address, _newOwner []common.Address) (*ConverterRegistryOwnerUpdateIterator, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ConverterRegistry.contract.FilterLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConverterRegistryOwnerUpdateIterator{contract: _ConverterRegistry.contract, event: "OwnerUpdate", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdate is a free log subscription operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_ConverterRegistry *ConverterRegistryFilterer) WatchOwnerUpdate(opts *bind.WatchOpts, sink chan<- *ConverterRegistryOwnerUpdate, _prevOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ConverterRegistry.contract.WatchLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterRegistryOwnerUpdate)
				if err := _ConverterRegistry.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
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

// ParseOwnerUpdate is a log parse operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_ConverterRegistry *ConverterRegistryFilterer) ParseOwnerUpdate(log types.Log) (*ConverterRegistryOwnerUpdate, error) {
	event := new(ConverterRegistryOwnerUpdate)
	if err := _ConverterRegistry.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterRegistrySmartTokenAddedIterator is returned from FilterSmartTokenAdded and is used to iterate over the raw logs and unpacked data for SmartTokenAdded events raised by the ConverterRegistry contract.
type ConverterRegistrySmartTokenAddedIterator struct {
	Event *ConverterRegistrySmartTokenAdded // Event containing the contract specifics and raw log

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
func (it *ConverterRegistrySmartTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterRegistrySmartTokenAdded)
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
		it.Event = new(ConverterRegistrySmartTokenAdded)
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
func (it *ConverterRegistrySmartTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterRegistrySmartTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterRegistrySmartTokenAdded represents a SmartTokenAdded event raised by the ConverterRegistry contract.
type ConverterRegistrySmartTokenAdded struct {
	SmartToken common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSmartTokenAdded is a free log retrieval operation binding the contract event 0x88881feecdf61136ac4bdb1f681f2f3746a82910263d21ffea94750d2a78c0ab.
//
// Solidity: event SmartTokenAdded(address indexed _smartToken)
func (_ConverterRegistry *ConverterRegistryFilterer) FilterSmartTokenAdded(opts *bind.FilterOpts, _smartToken []common.Address) (*ConverterRegistrySmartTokenAddedIterator, error) {

	var _smartTokenRule []interface{}
	for _, _smartTokenItem := range _smartToken {
		_smartTokenRule = append(_smartTokenRule, _smartTokenItem)
	}

	logs, sub, err := _ConverterRegistry.contract.FilterLogs(opts, "SmartTokenAdded", _smartTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConverterRegistrySmartTokenAddedIterator{contract: _ConverterRegistry.contract, event: "SmartTokenAdded", logs: logs, sub: sub}, nil
}

// WatchSmartTokenAdded is a free log subscription operation binding the contract event 0x88881feecdf61136ac4bdb1f681f2f3746a82910263d21ffea94750d2a78c0ab.
//
// Solidity: event SmartTokenAdded(address indexed _smartToken)
func (_ConverterRegistry *ConverterRegistryFilterer) WatchSmartTokenAdded(opts *bind.WatchOpts, sink chan<- *ConverterRegistrySmartTokenAdded, _smartToken []common.Address) (event.Subscription, error) {

	var _smartTokenRule []interface{}
	for _, _smartTokenItem := range _smartToken {
		_smartTokenRule = append(_smartTokenRule, _smartTokenItem)
	}

	logs, sub, err := _ConverterRegistry.contract.WatchLogs(opts, "SmartTokenAdded", _smartTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterRegistrySmartTokenAdded)
				if err := _ConverterRegistry.contract.UnpackLog(event, "SmartTokenAdded", log); err != nil {
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

// ParseSmartTokenAdded is a log parse operation binding the contract event 0x88881feecdf61136ac4bdb1f681f2f3746a82910263d21ffea94750d2a78c0ab.
//
// Solidity: event SmartTokenAdded(address indexed _smartToken)
func (_ConverterRegistry *ConverterRegistryFilterer) ParseSmartTokenAdded(log types.Log) (*ConverterRegistrySmartTokenAdded, error) {
	event := new(ConverterRegistrySmartTokenAdded)
	if err := _ConverterRegistry.contract.UnpackLog(event, "SmartTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterRegistrySmartTokenRemovedIterator is returned from FilterSmartTokenRemoved and is used to iterate over the raw logs and unpacked data for SmartTokenRemoved events raised by the ConverterRegistry contract.
type ConverterRegistrySmartTokenRemovedIterator struct {
	Event *ConverterRegistrySmartTokenRemoved // Event containing the contract specifics and raw log

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
func (it *ConverterRegistrySmartTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterRegistrySmartTokenRemoved)
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
		it.Event = new(ConverterRegistrySmartTokenRemoved)
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
func (it *ConverterRegistrySmartTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterRegistrySmartTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterRegistrySmartTokenRemoved represents a SmartTokenRemoved event raised by the ConverterRegistry contract.
type ConverterRegistrySmartTokenRemoved struct {
	SmartToken common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSmartTokenRemoved is a free log retrieval operation binding the contract event 0x2aff63790c7da80d1c50ede92d23bc841c384837735c92c184331f3d7b91e5bf.
//
// Solidity: event SmartTokenRemoved(address indexed _smartToken)
func (_ConverterRegistry *ConverterRegistryFilterer) FilterSmartTokenRemoved(opts *bind.FilterOpts, _smartToken []common.Address) (*ConverterRegistrySmartTokenRemovedIterator, error) {

	var _smartTokenRule []interface{}
	for _, _smartTokenItem := range _smartToken {
		_smartTokenRule = append(_smartTokenRule, _smartTokenItem)
	}

	logs, sub, err := _ConverterRegistry.contract.FilterLogs(opts, "SmartTokenRemoved", _smartTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConverterRegistrySmartTokenRemovedIterator{contract: _ConverterRegistry.contract, event: "SmartTokenRemoved", logs: logs, sub: sub}, nil
}

// WatchSmartTokenRemoved is a free log subscription operation binding the contract event 0x2aff63790c7da80d1c50ede92d23bc841c384837735c92c184331f3d7b91e5bf.
//
// Solidity: event SmartTokenRemoved(address indexed _smartToken)
func (_ConverterRegistry *ConverterRegistryFilterer) WatchSmartTokenRemoved(opts *bind.WatchOpts, sink chan<- *ConverterRegistrySmartTokenRemoved, _smartToken []common.Address) (event.Subscription, error) {

	var _smartTokenRule []interface{}
	for _, _smartTokenItem := range _smartToken {
		_smartTokenRule = append(_smartTokenRule, _smartTokenItem)
	}

	logs, sub, err := _ConverterRegistry.contract.WatchLogs(opts, "SmartTokenRemoved", _smartTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterRegistrySmartTokenRemoved)
				if err := _ConverterRegistry.contract.UnpackLog(event, "SmartTokenRemoved", log); err != nil {
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

// ParseSmartTokenRemoved is a log parse operation binding the contract event 0x2aff63790c7da80d1c50ede92d23bc841c384837735c92c184331f3d7b91e5bf.
//
// Solidity: event SmartTokenRemoved(address indexed _smartToken)
func (_ConverterRegistry *ConverterRegistryFilterer) ParseSmartTokenRemoved(log types.Log) (*ConverterRegistrySmartTokenRemoved, error) {
	event := new(ConverterRegistrySmartTokenRemoved)
	if err := _ConverterRegistry.contract.UnpackLog(event, "SmartTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
