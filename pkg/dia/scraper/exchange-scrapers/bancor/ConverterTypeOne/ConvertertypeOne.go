// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ConvertertypeOne

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

// ConvertertypeOneABI is the input ABI used to generate the binding from.
const ConvertertypeOneABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_onlyOwnerCanUpdateRegistry\",\"type\":\"bool\"}],\"name\":\"restrictRegistryUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveRatio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"connectors\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasETHReserve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"connectorTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserveToken\",\"type\":\"address\"}],\"name\":\"reserveWeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sourceToken\",\"type\":\"address\"},{\"name\":\"_targetToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getReturn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferTokenOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"onlyOwnerCanUpdateRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptTokenOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromAnchor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"converterType\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_whitelist\",\"type\":\"address\"}],\"name\":\"setConversionWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"conversionFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prevRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferAnchorOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_weight\",\"type\":\"uint32\"}],\"name\":\"addReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256\"}],\"name\":\"decimalLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"connectorTokenCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserveTokens\",\"type\":\"address[]\"},{\"name\":\"_reserveAmounts\",\"type\":\"uint256[]\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxConversionFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveTokenCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"geometricMean\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sourceToken\",\"type\":\"address\"},{\"name\":\"_targetToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"targetAmountAndFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_reserveTokens\",\"type\":\"address[]\"},{\"name\":\"_reserveMinReturnAmounts\",\"type\":\"uint256[]\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restoreRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_n\",\"type\":\"uint256\"},{\"name\":\"_d\",\"type\":\"uint256\"}],\"name\":\"roundDiv\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"conversionsEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"conversionWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptAnchorOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reserveTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isV28OrHigher\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"anchor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"reserves\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint32\"},{\"name\":\"deprecated1\",\"type\":\"bool\"},{\"name\":\"deprecated2\",\"type\":\"bool\"},{\"name\":\"isSet\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_connectorToken\",\"type\":\"address\"}],\"name\":\"getConnectorBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserveToken\",\"type\":\"address\"}],\"name\":\"reserveBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sourceToken\",\"type\":\"address\"},{\"name\":\"_targetToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_trader\",\"type\":\"address\"},{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"convert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_conversionFee\",\"type\":\"uint32\"}],\"name\":\"setConversionFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_registry\",\"type\":\"address\"},{\"name\":\"_maxConversionFee\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_connectorToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_connectorBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_connectorWeight\",\"type\":\"uint32\"}],\"name\":\"PriceDataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_provider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_reserveToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_newBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_newSupply\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_provider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_reserveToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_newBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_newSupply\",\"type\":\"uint256\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_type\",\"type\":\"uint16\"},{\"indexed\":true,\"name\":\"_anchor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_activated\",\"type\":\"bool\"}],\"name\":\"Activation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_fromToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_toToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_trader\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_return\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_conversionFee\",\"type\":\"int256\"}],\"name\":\"Conversion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_token1\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_token2\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_rateN\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_rateD\",\"type\":\"uint256\"}],\"name\":\"TokenRateUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prevFee\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"_newFee\",\"type\":\"uint32\"}],\"name\":\"ConversionFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// ConvertertypeOne is an auto generated Go binding around an Ethereum contract.
type ConvertertypeOne struct {
	ConvertertypeOneCaller     // Read-only binding to the contract
	ConvertertypeOneTransactor // Write-only binding to the contract
	ConvertertypeOneFilterer   // Log filterer for contract events
}

// ConvertertypeOneCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConvertertypeOneCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvertertypeOneTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConvertertypeOneTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvertertypeOneFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConvertertypeOneFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvertertypeOneSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConvertertypeOneSession struct {
	Contract     *ConvertertypeOne // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConvertertypeOneCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConvertertypeOneCallerSession struct {
	Contract *ConvertertypeOneCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ConvertertypeOneTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConvertertypeOneTransactorSession struct {
	Contract     *ConvertertypeOneTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ConvertertypeOneRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConvertertypeOneRaw struct {
	Contract *ConvertertypeOne // Generic contract binding to access the raw methods on
}

// ConvertertypeOneCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConvertertypeOneCallerRaw struct {
	Contract *ConvertertypeOneCaller // Generic read-only contract binding to access the raw methods on
}

// ConvertertypeOneTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConvertertypeOneTransactorRaw struct {
	Contract *ConvertertypeOneTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConvertertypeOne creates a new instance of ConvertertypeOne, bound to a specific deployed contract.
func NewConvertertypeOne(address common.Address, backend bind.ContractBackend) (*ConvertertypeOne, error) {
	contract, err := bindConvertertypeOne(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConvertertypeOne{ConvertertypeOneCaller: ConvertertypeOneCaller{contract: contract}, ConvertertypeOneTransactor: ConvertertypeOneTransactor{contract: contract}, ConvertertypeOneFilterer: ConvertertypeOneFilterer{contract: contract}}, nil
}

// NewConvertertypeOneCaller creates a new read-only instance of ConvertertypeOne, bound to a specific deployed contract.
func NewConvertertypeOneCaller(address common.Address, caller bind.ContractCaller) (*ConvertertypeOneCaller, error) {
	contract, err := bindConvertertypeOne(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConvertertypeOneCaller{contract: contract}, nil
}

// NewConvertertypeOneTransactor creates a new write-only instance of ConvertertypeOne, bound to a specific deployed contract.
func NewConvertertypeOneTransactor(address common.Address, transactor bind.ContractTransactor) (*ConvertertypeOneTransactor, error) {
	contract, err := bindConvertertypeOne(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConvertertypeOneTransactor{contract: contract}, nil
}

// NewConvertertypeOneFilterer creates a new log filterer instance of ConvertertypeOne, bound to a specific deployed contract.
func NewConvertertypeOneFilterer(address common.Address, filterer bind.ContractFilterer) (*ConvertertypeOneFilterer, error) {
	contract, err := bindConvertertypeOne(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConvertertypeOneFilterer{contract: contract}, nil
}

// bindConvertertypeOne binds a generic wrapper to an already deployed contract.
func bindConvertertypeOne(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConvertertypeOneABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConvertertypeOne *ConvertertypeOneRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConvertertypeOne.Contract.ConvertertypeOneCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConvertertypeOne *ConvertertypeOneRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.ConvertertypeOneTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConvertertypeOne *ConvertertypeOneRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.ConvertertypeOneTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConvertertypeOne *ConvertertypeOneCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConvertertypeOne.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConvertertypeOne *ConvertertypeOneTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConvertertypeOne *ConvertertypeOneTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.contract.Transact(opts, method, params...)
}

// Anchor is a free data retrieval call binding the contract method 0xd3fb73b4.
//
// Solidity: function anchor() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCaller) Anchor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "anchor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Anchor is a free data retrieval call binding the contract method 0xd3fb73b4.
//
// Solidity: function anchor() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneSession) Anchor() (common.Address, error) {
	return _ConvertertypeOne.Contract.Anchor(&_ConvertertypeOne.CallOpts)
}

// Anchor is a free data retrieval call binding the contract method 0xd3fb73b4.
//
// Solidity: function anchor() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) Anchor() (common.Address, error) {
	return _ConvertertypeOne.Contract.Anchor(&_ConvertertypeOne.CallOpts)
}

// ConnectorTokenCount is a free data retrieval call binding the contract method 0x71f52bf3.
//
// Solidity: function connectorTokenCount() view returns(uint16)
func (_ConvertertypeOne *ConvertertypeOneCaller) ConnectorTokenCount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "connectorTokenCount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ConnectorTokenCount is a free data retrieval call binding the contract method 0x71f52bf3.
//
// Solidity: function connectorTokenCount() view returns(uint16)
func (_ConvertertypeOne *ConvertertypeOneSession) ConnectorTokenCount() (uint16, error) {
	return _ConvertertypeOne.Contract.ConnectorTokenCount(&_ConvertertypeOne.CallOpts)
}

// ConnectorTokenCount is a free data retrieval call binding the contract method 0x71f52bf3.
//
// Solidity: function connectorTokenCount() view returns(uint16)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) ConnectorTokenCount() (uint16, error) {
	return _ConvertertypeOne.Contract.ConnectorTokenCount(&_ConvertertypeOne.CallOpts)
}

// ConnectorTokens is a free data retrieval call binding the contract method 0x19b64015.
//
// Solidity: function connectorTokens(uint256 _index) view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCaller) ConnectorTokens(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "connectorTokens", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorTokens is a free data retrieval call binding the contract method 0x19b64015.
//
// Solidity: function connectorTokens(uint256 _index) view returns(address)
func (_ConvertertypeOne *ConvertertypeOneSession) ConnectorTokens(_index *big.Int) (common.Address, error) {
	return _ConvertertypeOne.Contract.ConnectorTokens(&_ConvertertypeOne.CallOpts, _index)
}

// ConnectorTokens is a free data retrieval call binding the contract method 0x19b64015.
//
// Solidity: function connectorTokens(uint256 _index) view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) ConnectorTokens(_index *big.Int) (common.Address, error) {
	return _ConvertertypeOne.Contract.ConnectorTokens(&_ConvertertypeOne.CallOpts, _index)
}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address _address) view returns(uint256, uint32, bool, bool, bool)
func (_ConvertertypeOne *ConvertertypeOneCaller) Connectors(opts *bind.CallOpts, _address common.Address) (*big.Int, uint32, bool, bool, bool, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "connectors", _address)

	if err != nil {
		return *new(*big.Int), *new(uint32), *new(bool), *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)
	out4 := *abi.ConvertType(out[4], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, err

}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address _address) view returns(uint256, uint32, bool, bool, bool)
func (_ConvertertypeOne *ConvertertypeOneSession) Connectors(_address common.Address) (*big.Int, uint32, bool, bool, bool, error) {
	return _ConvertertypeOne.Contract.Connectors(&_ConvertertypeOne.CallOpts, _address)
}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address _address) view returns(uint256, uint32, bool, bool, bool)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) Connectors(_address common.Address) (*big.Int, uint32, bool, bool, bool, error) {
	return _ConvertertypeOne.Contract.Connectors(&_ConvertertypeOne.CallOpts, _address)
}

// ConversionFee is a free data retrieval call binding the contract method 0x579cd3ca.
//
// Solidity: function conversionFee() view returns(uint32)
func (_ConvertertypeOne *ConvertertypeOneCaller) ConversionFee(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "conversionFee")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ConversionFee is a free data retrieval call binding the contract method 0x579cd3ca.
//
// Solidity: function conversionFee() view returns(uint32)
func (_ConvertertypeOne *ConvertertypeOneSession) ConversionFee() (uint32, error) {
	return _ConvertertypeOne.Contract.ConversionFee(&_ConvertertypeOne.CallOpts)
}

// ConversionFee is a free data retrieval call binding the contract method 0x579cd3ca.
//
// Solidity: function conversionFee() view returns(uint32)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) ConversionFee() (uint32, error) {
	return _ConvertertypeOne.Contract.ConversionFee(&_ConvertertypeOne.CallOpts)
}

// ConversionWhitelist is a free data retrieval call binding the contract method 0xc45d3d92.
//
// Solidity: function conversionWhitelist() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCaller) ConversionWhitelist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "conversionWhitelist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConversionWhitelist is a free data retrieval call binding the contract method 0xc45d3d92.
//
// Solidity: function conversionWhitelist() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneSession) ConversionWhitelist() (common.Address, error) {
	return _ConvertertypeOne.Contract.ConversionWhitelist(&_ConvertertypeOne.CallOpts)
}

// ConversionWhitelist is a free data retrieval call binding the contract method 0xc45d3d92.
//
// Solidity: function conversionWhitelist() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) ConversionWhitelist() (common.Address, error) {
	return _ConvertertypeOne.Contract.ConversionWhitelist(&_ConvertertypeOne.CallOpts)
}

// ConversionsEnabled is a free data retrieval call binding the contract method 0xbf754558.
//
// Solidity: function conversionsEnabled() view returns(bool)
func (_ConvertertypeOne *ConvertertypeOneCaller) ConversionsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "conversionsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ConversionsEnabled is a free data retrieval call binding the contract method 0xbf754558.
//
// Solidity: function conversionsEnabled() view returns(bool)
func (_ConvertertypeOne *ConvertertypeOneSession) ConversionsEnabled() (bool, error) {
	return _ConvertertypeOne.Contract.ConversionsEnabled(&_ConvertertypeOne.CallOpts)
}

// ConversionsEnabled is a free data retrieval call binding the contract method 0xbf754558.
//
// Solidity: function conversionsEnabled() view returns(bool)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) ConversionsEnabled() (bool, error) {
	return _ConvertertypeOne.Contract.ConversionsEnabled(&_ConvertertypeOne.CallOpts)
}

// ConverterType is a free data retrieval call binding the contract method 0x3e8ff43f.
//
// Solidity: function converterType() pure returns(uint16)
func (_ConvertertypeOne *ConvertertypeOneCaller) ConverterType(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "converterType")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ConverterType is a free data retrieval call binding the contract method 0x3e8ff43f.
//
// Solidity: function converterType() pure returns(uint16)
func (_ConvertertypeOne *ConvertertypeOneSession) ConverterType() (uint16, error) {
	return _ConvertertypeOne.Contract.ConverterType(&_ConvertertypeOne.CallOpts)
}

// ConverterType is a free data retrieval call binding the contract method 0x3e8ff43f.
//
// Solidity: function converterType() pure returns(uint16)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) ConverterType() (uint16, error) {
	return _ConvertertypeOne.Contract.ConverterType(&_ConvertertypeOne.CallOpts)
}

// DecimalLength is a free data retrieval call binding the contract method 0x6aa5332c.
//
// Solidity: function decimalLength(uint256 _x) pure returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneCaller) DecimalLength(opts *bind.CallOpts, _x *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "decimalLength", _x)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalLength is a free data retrieval call binding the contract method 0x6aa5332c.
//
// Solidity: function decimalLength(uint256 _x) pure returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneSession) DecimalLength(_x *big.Int) (*big.Int, error) {
	return _ConvertertypeOne.Contract.DecimalLength(&_ConvertertypeOne.CallOpts, _x)
}

// DecimalLength is a free data retrieval call binding the contract method 0x6aa5332c.
//
// Solidity: function decimalLength(uint256 _x) pure returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) DecimalLength(_x *big.Int) (*big.Int, error) {
	return _ConvertertypeOne.Contract.DecimalLength(&_ConvertertypeOne.CallOpts, _x)
}

// GeometricMean is a free data retrieval call binding the contract method 0xa60e7724.
//
// Solidity: function geometricMean(uint256[] _values) pure returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneCaller) GeometricMean(opts *bind.CallOpts, _values []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "geometricMean", _values)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeometricMean is a free data retrieval call binding the contract method 0xa60e7724.
//
// Solidity: function geometricMean(uint256[] _values) pure returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneSession) GeometricMean(_values []*big.Int) (*big.Int, error) {
	return _ConvertertypeOne.Contract.GeometricMean(&_ConvertertypeOne.CallOpts, _values)
}

// GeometricMean is a free data retrieval call binding the contract method 0xa60e7724.
//
// Solidity: function geometricMean(uint256[] _values) pure returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) GeometricMean(_values []*big.Int) (*big.Int, error) {
	return _ConvertertypeOne.Contract.GeometricMean(&_ConvertertypeOne.CallOpts, _values)
}

// GetConnectorBalance is a free data retrieval call binding the contract method 0xd8959512.
//
// Solidity: function getConnectorBalance(address _connectorToken) view returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneCaller) GetConnectorBalance(opts *bind.CallOpts, _connectorToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "getConnectorBalance", _connectorToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConnectorBalance is a free data retrieval call binding the contract method 0xd8959512.
//
// Solidity: function getConnectorBalance(address _connectorToken) view returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneSession) GetConnectorBalance(_connectorToken common.Address) (*big.Int, error) {
	return _ConvertertypeOne.Contract.GetConnectorBalance(&_ConvertertypeOne.CallOpts, _connectorToken)
}

// GetConnectorBalance is a free data retrieval call binding the contract method 0xd8959512.
//
// Solidity: function getConnectorBalance(address _connectorToken) view returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) GetConnectorBalance(_connectorToken common.Address) (*big.Int, error) {
	return _ConvertertypeOne.Contract.GetConnectorBalance(&_ConvertertypeOne.CallOpts, _connectorToken)
}

// GetReturn is a free data retrieval call binding the contract method 0x1e1401f8.
//
// Solidity: function getReturn(address _sourceToken, address _targetToken, uint256 _amount) view returns(uint256, uint256)
func (_ConvertertypeOne *ConvertertypeOneCaller) GetReturn(opts *bind.CallOpts, _sourceToken common.Address, _targetToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "getReturn", _sourceToken, _targetToken, _amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetReturn is a free data retrieval call binding the contract method 0x1e1401f8.
//
// Solidity: function getReturn(address _sourceToken, address _targetToken, uint256 _amount) view returns(uint256, uint256)
func (_ConvertertypeOne *ConvertertypeOneSession) GetReturn(_sourceToken common.Address, _targetToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConvertertypeOne.Contract.GetReturn(&_ConvertertypeOne.CallOpts, _sourceToken, _targetToken, _amount)
}

// GetReturn is a free data retrieval call binding the contract method 0x1e1401f8.
//
// Solidity: function getReturn(address _sourceToken, address _targetToken, uint256 _amount) view returns(uint256, uint256)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) GetReturn(_sourceToken common.Address, _targetToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConvertertypeOne.Contract.GetReturn(&_ConvertertypeOne.CallOpts, _sourceToken, _targetToken, _amount)
}

// HasETHReserve is a free data retrieval call binding the contract method 0x12c2aca4.
//
// Solidity: function hasETHReserve() view returns(bool)
func (_ConvertertypeOne *ConvertertypeOneCaller) HasETHReserve(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "hasETHReserve")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasETHReserve is a free data retrieval call binding the contract method 0x12c2aca4.
//
// Solidity: function hasETHReserve() view returns(bool)
func (_ConvertertypeOne *ConvertertypeOneSession) HasETHReserve() (bool, error) {
	return _ConvertertypeOne.Contract.HasETHReserve(&_ConvertertypeOne.CallOpts)
}

// HasETHReserve is a free data retrieval call binding the contract method 0x12c2aca4.
//
// Solidity: function hasETHReserve() view returns(bool)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) HasETHReserve() (bool, error) {
	return _ConvertertypeOne.Contract.HasETHReserve(&_ConvertertypeOne.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_ConvertertypeOne *ConvertertypeOneCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_ConvertertypeOne *ConvertertypeOneSession) IsActive() (bool, error) {
	return _ConvertertypeOne.Contract.IsActive(&_ConvertertypeOne.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) IsActive() (bool, error) {
	return _ConvertertypeOne.Contract.IsActive(&_ConvertertypeOne.CallOpts)
}

// IsV28OrHigher is a free data retrieval call binding the contract method 0xd260529c.
//
// Solidity: function isV28OrHigher() pure returns(bool)
func (_ConvertertypeOne *ConvertertypeOneCaller) IsV28OrHigher(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "isV28OrHigher")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsV28OrHigher is a free data retrieval call binding the contract method 0xd260529c.
//
// Solidity: function isV28OrHigher() pure returns(bool)
func (_ConvertertypeOne *ConvertertypeOneSession) IsV28OrHigher() (bool, error) {
	return _ConvertertypeOne.Contract.IsV28OrHigher(&_ConvertertypeOne.CallOpts)
}

// IsV28OrHigher is a free data retrieval call binding the contract method 0xd260529c.
//
// Solidity: function isV28OrHigher() pure returns(bool)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) IsV28OrHigher() (bool, error) {
	return _ConvertertypeOne.Contract.IsV28OrHigher(&_ConvertertypeOne.CallOpts)
}

// MaxConversionFee is a free data retrieval call binding the contract method 0x94c275ad.
//
// Solidity: function maxConversionFee() view returns(uint32)
func (_ConvertertypeOne *ConvertertypeOneCaller) MaxConversionFee(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "maxConversionFee")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxConversionFee is a free data retrieval call binding the contract method 0x94c275ad.
//
// Solidity: function maxConversionFee() view returns(uint32)
func (_ConvertertypeOne *ConvertertypeOneSession) MaxConversionFee() (uint32, error) {
	return _ConvertertypeOne.Contract.MaxConversionFee(&_ConvertertypeOne.CallOpts)
}

// MaxConversionFee is a free data retrieval call binding the contract method 0x94c275ad.
//
// Solidity: function maxConversionFee() view returns(uint32)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) MaxConversionFee() (uint32, error) {
	return _ConvertertypeOne.Contract.MaxConversionFee(&_ConvertertypeOne.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "newOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneSession) NewOwner() (common.Address, error) {
	return _ConvertertypeOne.Contract.NewOwner(&_ConvertertypeOne.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) NewOwner() (common.Address, error) {
	return _ConvertertypeOne.Contract.NewOwner(&_ConvertertypeOne.CallOpts)
}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_ConvertertypeOne *ConvertertypeOneCaller) OnlyOwnerCanUpdateRegistry(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "onlyOwnerCanUpdateRegistry")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_ConvertertypeOne *ConvertertypeOneSession) OnlyOwnerCanUpdateRegistry() (bool, error) {
	return _ConvertertypeOne.Contract.OnlyOwnerCanUpdateRegistry(&_ConvertertypeOne.CallOpts)
}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) OnlyOwnerCanUpdateRegistry() (bool, error) {
	return _ConvertertypeOne.Contract.OnlyOwnerCanUpdateRegistry(&_ConvertertypeOne.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneSession) Owner() (common.Address, error) {
	return _ConvertertypeOne.Contract.Owner(&_ConvertertypeOne.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) Owner() (common.Address, error) {
	return _ConvertertypeOne.Contract.Owner(&_ConvertertypeOne.CallOpts)
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCaller) PrevRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "prevRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneSession) PrevRegistry() (common.Address, error) {
	return _ConvertertypeOne.Contract.PrevRegistry(&_ConvertertypeOne.CallOpts)
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) PrevRegistry() (common.Address, error) {
	return _ConvertertypeOne.Contract.PrevRegistry(&_ConvertertypeOne.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneSession) Registry() (common.Address, error) {
	return _ConvertertypeOne.Contract.Registry(&_ConvertertypeOne.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) Registry() (common.Address, error) {
	return _ConvertertypeOne.Contract.Registry(&_ConvertertypeOne.CallOpts)
}

// ReserveBalance is a free data retrieval call binding the contract method 0xdc8de379.
//
// Solidity: function reserveBalance(address _reserveToken) view returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneCaller) ReserveBalance(opts *bind.CallOpts, _reserveToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "reserveBalance", _reserveToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveBalance is a free data retrieval call binding the contract method 0xdc8de379.
//
// Solidity: function reserveBalance(address _reserveToken) view returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneSession) ReserveBalance(_reserveToken common.Address) (*big.Int, error) {
	return _ConvertertypeOne.Contract.ReserveBalance(&_ConvertertypeOne.CallOpts, _reserveToken)
}

// ReserveBalance is a free data retrieval call binding the contract method 0xdc8de379.
//
// Solidity: function reserveBalance(address _reserveToken) view returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) ReserveBalance(_reserveToken common.Address) (*big.Int, error) {
	return _ConvertertypeOne.Contract.ReserveBalance(&_ConvertertypeOne.CallOpts, _reserveToken)
}

// ReserveRatio is a free data retrieval call binding the contract method 0x0c7d5cd8.
//
// Solidity: function reserveRatio() view returns(uint32)
func (_ConvertertypeOne *ConvertertypeOneCaller) ReserveRatio(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "reserveRatio")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ReserveRatio is a free data retrieval call binding the contract method 0x0c7d5cd8.
//
// Solidity: function reserveRatio() view returns(uint32)
func (_ConvertertypeOne *ConvertertypeOneSession) ReserveRatio() (uint32, error) {
	return _ConvertertypeOne.Contract.ReserveRatio(&_ConvertertypeOne.CallOpts)
}

// ReserveRatio is a free data retrieval call binding the contract method 0x0c7d5cd8.
//
// Solidity: function reserveRatio() view returns(uint32)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) ReserveRatio() (uint32, error) {
	return _ConvertertypeOne.Contract.ReserveRatio(&_ConvertertypeOne.CallOpts)
}

// ReserveTokenCount is a free data retrieval call binding the contract method 0x9b99a8e2.
//
// Solidity: function reserveTokenCount() view returns(uint16)
func (_ConvertertypeOne *ConvertertypeOneCaller) ReserveTokenCount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "reserveTokenCount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ReserveTokenCount is a free data retrieval call binding the contract method 0x9b99a8e2.
//
// Solidity: function reserveTokenCount() view returns(uint16)
func (_ConvertertypeOne *ConvertertypeOneSession) ReserveTokenCount() (uint16, error) {
	return _ConvertertypeOne.Contract.ReserveTokenCount(&_ConvertertypeOne.CallOpts)
}

// ReserveTokenCount is a free data retrieval call binding the contract method 0x9b99a8e2.
//
// Solidity: function reserveTokenCount() view returns(uint16)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) ReserveTokenCount() (uint16, error) {
	return _ConvertertypeOne.Contract.ReserveTokenCount(&_ConvertertypeOne.CallOpts)
}

// ReserveTokens is a free data retrieval call binding the contract method 0xd031370b.
//
// Solidity: function reserveTokens(uint256 ) view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCaller) ReserveTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "reserveTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReserveTokens is a free data retrieval call binding the contract method 0xd031370b.
//
// Solidity: function reserveTokens(uint256 ) view returns(address)
func (_ConvertertypeOne *ConvertertypeOneSession) ReserveTokens(arg0 *big.Int) (common.Address, error) {
	return _ConvertertypeOne.Contract.ReserveTokens(&_ConvertertypeOne.CallOpts, arg0)
}

// ReserveTokens is a free data retrieval call binding the contract method 0xd031370b.
//
// Solidity: function reserveTokens(uint256 ) view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) ReserveTokens(arg0 *big.Int) (common.Address, error) {
	return _ConvertertypeOne.Contract.ReserveTokens(&_ConvertertypeOne.CallOpts, arg0)
}

// ReserveWeight is a free data retrieval call binding the contract method 0x1cfab290.
//
// Solidity: function reserveWeight(address _reserveToken) view returns(uint32)
func (_ConvertertypeOne *ConvertertypeOneCaller) ReserveWeight(opts *bind.CallOpts, _reserveToken common.Address) (uint32, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "reserveWeight", _reserveToken)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ReserveWeight is a free data retrieval call binding the contract method 0x1cfab290.
//
// Solidity: function reserveWeight(address _reserveToken) view returns(uint32)
func (_ConvertertypeOne *ConvertertypeOneSession) ReserveWeight(_reserveToken common.Address) (uint32, error) {
	return _ConvertertypeOne.Contract.ReserveWeight(&_ConvertertypeOne.CallOpts, _reserveToken)
}

// ReserveWeight is a free data retrieval call binding the contract method 0x1cfab290.
//
// Solidity: function reserveWeight(address _reserveToken) view returns(uint32)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) ReserveWeight(_reserveToken common.Address) (uint32, error) {
	return _ConvertertypeOne.Contract.ReserveWeight(&_ConvertertypeOne.CallOpts, _reserveToken)
}

// Reserves is a free data retrieval call binding the contract method 0xd66bd524.
//
// Solidity: function reserves(address ) view returns(uint256 balance, uint32 weight, bool deprecated1, bool deprecated2, bool isSet)
func (_ConvertertypeOne *ConvertertypeOneCaller) Reserves(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance     *big.Int
	Weight      uint32
	Deprecated1 bool
	Deprecated2 bool
	IsSet       bool
}, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "reserves", arg0)

	outstruct := new(struct {
		Balance     *big.Int
		Weight      uint32
		Deprecated1 bool
		Deprecated2 bool
		IsSet       bool
	})

	outstruct.Balance = out[0].(*big.Int)
	outstruct.Weight = out[1].(uint32)
	outstruct.Deprecated1 = out[2].(bool)
	outstruct.Deprecated2 = out[3].(bool)
	outstruct.IsSet = out[4].(bool)

	return *outstruct, err

}

// Reserves is a free data retrieval call binding the contract method 0xd66bd524.
//
// Solidity: function reserves(address ) view returns(uint256 balance, uint32 weight, bool deprecated1, bool deprecated2, bool isSet)
func (_ConvertertypeOne *ConvertertypeOneSession) Reserves(arg0 common.Address) (struct {
	Balance     *big.Int
	Weight      uint32
	Deprecated1 bool
	Deprecated2 bool
	IsSet       bool
}, error) {
	return _ConvertertypeOne.Contract.Reserves(&_ConvertertypeOne.CallOpts, arg0)
}

// Reserves is a free data retrieval call binding the contract method 0xd66bd524.
//
// Solidity: function reserves(address ) view returns(uint256 balance, uint32 weight, bool deprecated1, bool deprecated2, bool isSet)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) Reserves(arg0 common.Address) (struct {
	Balance     *big.Int
	Weight      uint32
	Deprecated1 bool
	Deprecated2 bool
	IsSet       bool
}, error) {
	return _ConvertertypeOne.Contract.Reserves(&_ConvertertypeOne.CallOpts, arg0)
}

// RoundDiv is a free data retrieval call binding the contract method 0xbbb7e5d8.
//
// Solidity: function roundDiv(uint256 _n, uint256 _d) pure returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneCaller) RoundDiv(opts *bind.CallOpts, _n *big.Int, _d *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "roundDiv", _n, _d)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundDiv is a free data retrieval call binding the contract method 0xbbb7e5d8.
//
// Solidity: function roundDiv(uint256 _n, uint256 _d) pure returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneSession) RoundDiv(_n *big.Int, _d *big.Int) (*big.Int, error) {
	return _ConvertertypeOne.Contract.RoundDiv(&_ConvertertypeOne.CallOpts, _n, _d)
}

// RoundDiv is a free data retrieval call binding the contract method 0xbbb7e5d8.
//
// Solidity: function roundDiv(uint256 _n, uint256 _d) pure returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) RoundDiv(_n *big.Int, _d *big.Int) (*big.Int, error) {
	return _ConvertertypeOne.Contract.RoundDiv(&_ConvertertypeOne.CallOpts, _n, _d)
}

// TargetAmountAndFee is a free data retrieval call binding the contract method 0xaf94b8d8.
//
// Solidity: function targetAmountAndFee(address _sourceToken, address _targetToken, uint256 _amount) view returns(uint256, uint256)
func (_ConvertertypeOne *ConvertertypeOneCaller) TargetAmountAndFee(opts *bind.CallOpts, _sourceToken common.Address, _targetToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "targetAmountAndFee", _sourceToken, _targetToken, _amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// TargetAmountAndFee is a free data retrieval call binding the contract method 0xaf94b8d8.
//
// Solidity: function targetAmountAndFee(address _sourceToken, address _targetToken, uint256 _amount) view returns(uint256, uint256)
func (_ConvertertypeOne *ConvertertypeOneSession) TargetAmountAndFee(_sourceToken common.Address, _targetToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConvertertypeOne.Contract.TargetAmountAndFee(&_ConvertertypeOne.CallOpts, _sourceToken, _targetToken, _amount)
}

// TargetAmountAndFee is a free data retrieval call binding the contract method 0xaf94b8d8.
//
// Solidity: function targetAmountAndFee(address _sourceToken, address _targetToken, uint256 _amount) view returns(uint256, uint256)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) TargetAmountAndFee(_sourceToken common.Address, _targetToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConvertertypeOne.Contract.TargetAmountAndFee(&_ConvertertypeOne.CallOpts, _sourceToken, _targetToken, _amount)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneSession) Token() (common.Address, error) {
	return _ConvertertypeOne.Contract.Token(&_ConvertertypeOne.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) Token() (common.Address, error) {
	return _ConvertertypeOne.Contract.Token(&_ConvertertypeOne.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_ConvertertypeOne *ConvertertypeOneCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConvertertypeOne.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_ConvertertypeOne *ConvertertypeOneSession) Version() (uint16, error) {
	return _ConvertertypeOne.Contract.Version(&_ConvertertypeOne.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_ConvertertypeOne *ConvertertypeOneCallerSession) Version() (uint16, error) {
	return _ConvertertypeOne.Contract.Version(&_ConvertertypeOne.CallOpts)
}

// AcceptAnchorOwnership is a paid mutator transaction binding the contract method 0xcdc91c69.
//
// Solidity: function acceptAnchorOwnership() returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) AcceptAnchorOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "acceptAnchorOwnership")
}

// AcceptAnchorOwnership is a paid mutator transaction binding the contract method 0xcdc91c69.
//
// Solidity: function acceptAnchorOwnership() returns()
func (_ConvertertypeOne *ConvertertypeOneSession) AcceptAnchorOwnership() (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.AcceptAnchorOwnership(&_ConvertertypeOne.TransactOpts)
}

// AcceptAnchorOwnership is a paid mutator transaction binding the contract method 0xcdc91c69.
//
// Solidity: function acceptAnchorOwnership() returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) AcceptAnchorOwnership() (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.AcceptAnchorOwnership(&_ConvertertypeOne.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ConvertertypeOne *ConvertertypeOneSession) AcceptOwnership() (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.AcceptOwnership(&_ConvertertypeOne.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.AcceptOwnership(&_ConvertertypeOne.TransactOpts)
}

// AcceptTokenOwnership is a paid mutator transaction binding the contract method 0x38a5e016.
//
// Solidity: function acceptTokenOwnership() returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) AcceptTokenOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "acceptTokenOwnership")
}

// AcceptTokenOwnership is a paid mutator transaction binding the contract method 0x38a5e016.
//
// Solidity: function acceptTokenOwnership() returns()
func (_ConvertertypeOne *ConvertertypeOneSession) AcceptTokenOwnership() (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.AcceptTokenOwnership(&_ConvertertypeOne.TransactOpts)
}

// AcceptTokenOwnership is a paid mutator transaction binding the contract method 0x38a5e016.
//
// Solidity: function acceptTokenOwnership() returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) AcceptTokenOwnership() (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.AcceptTokenOwnership(&_ConvertertypeOne.TransactOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x7d8916bd.
//
// Solidity: function addLiquidity(address[] _reserveTokens, uint256[] _reserveAmounts, uint256 _minReturn) payable returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) AddLiquidity(opts *bind.TransactOpts, _reserveTokens []common.Address, _reserveAmounts []*big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "addLiquidity", _reserveTokens, _reserveAmounts, _minReturn)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x7d8916bd.
//
// Solidity: function addLiquidity(address[] _reserveTokens, uint256[] _reserveAmounts, uint256 _minReturn) payable returns()
func (_ConvertertypeOne *ConvertertypeOneSession) AddLiquidity(_reserveTokens []common.Address, _reserveAmounts []*big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.AddLiquidity(&_ConvertertypeOne.TransactOpts, _reserveTokens, _reserveAmounts, _minReturn)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x7d8916bd.
//
// Solidity: function addLiquidity(address[] _reserveTokens, uint256[] _reserveAmounts, uint256 _minReturn) payable returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) AddLiquidity(_reserveTokens []common.Address, _reserveAmounts []*big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.AddLiquidity(&_ConvertertypeOne.TransactOpts, _reserveTokens, _reserveAmounts, _minReturn)
}

// AddReserve is a paid mutator transaction binding the contract method 0x6a49d2c4.
//
// Solidity: function addReserve(address _token, uint32 _weight) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) AddReserve(opts *bind.TransactOpts, _token common.Address, _weight uint32) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "addReserve", _token, _weight)
}

// AddReserve is a paid mutator transaction binding the contract method 0x6a49d2c4.
//
// Solidity: function addReserve(address _token, uint32 _weight) returns()
func (_ConvertertypeOne *ConvertertypeOneSession) AddReserve(_token common.Address, _weight uint32) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.AddReserve(&_ConvertertypeOne.TransactOpts, _token, _weight)
}

// AddReserve is a paid mutator transaction binding the contract method 0x6a49d2c4.
//
// Solidity: function addReserve(address _token, uint32 _weight) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) AddReserve(_token common.Address, _weight uint32) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.AddReserve(&_ConvertertypeOne.TransactOpts, _token, _weight)
}

// Convert is a paid mutator transaction binding the contract method 0xe8dc12ff.
//
// Solidity: function convert(address _sourceToken, address _targetToken, uint256 _amount, address _trader, address _beneficiary) payable returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneTransactor) Convert(opts *bind.TransactOpts, _sourceToken common.Address, _targetToken common.Address, _amount *big.Int, _trader common.Address, _beneficiary common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "convert", _sourceToken, _targetToken, _amount, _trader, _beneficiary)
}

// Convert is a paid mutator transaction binding the contract method 0xe8dc12ff.
//
// Solidity: function convert(address _sourceToken, address _targetToken, uint256 _amount, address _trader, address _beneficiary) payable returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneSession) Convert(_sourceToken common.Address, _targetToken common.Address, _amount *big.Int, _trader common.Address, _beneficiary common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.Convert(&_ConvertertypeOne.TransactOpts, _sourceToken, _targetToken, _amount, _trader, _beneficiary)
}

// Convert is a paid mutator transaction binding the contract method 0xe8dc12ff.
//
// Solidity: function convert(address _sourceToken, address _targetToken, uint256 _amount, address _trader, address _beneficiary) payable returns(uint256)
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) Convert(_sourceToken common.Address, _targetToken common.Address, _amount *big.Int, _trader common.Address, _beneficiary common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.Convert(&_ConvertertypeOne.TransactOpts, _sourceToken, _targetToken, _amount, _trader, _beneficiary)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 _amount) payable returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) Fund(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "fund", _amount)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 _amount) payable returns()
func (_ConvertertypeOne *ConvertertypeOneSession) Fund(_amount *big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.Fund(&_ConvertertypeOne.TransactOpts, _amount)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 _amount) payable returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) Fund(_amount *big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.Fund(&_ConvertertypeOne.TransactOpts, _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 _amount) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) Liquidate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "liquidate", _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 _amount) returns()
func (_ConvertertypeOne *ConvertertypeOneSession) Liquidate(_amount *big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.Liquidate(&_ConvertertypeOne.TransactOpts, _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 _amount) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) Liquidate(_amount *big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.Liquidate(&_ConvertertypeOne.TransactOpts, _amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xb127c0a5.
//
// Solidity: function removeLiquidity(uint256 _amount, address[] _reserveTokens, uint256[] _reserveMinReturnAmounts) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, _reserveTokens []common.Address, _reserveMinReturnAmounts []*big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "removeLiquidity", _amount, _reserveTokens, _reserveMinReturnAmounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xb127c0a5.
//
// Solidity: function removeLiquidity(uint256 _amount, address[] _reserveTokens, uint256[] _reserveMinReturnAmounts) returns()
func (_ConvertertypeOne *ConvertertypeOneSession) RemoveLiquidity(_amount *big.Int, _reserveTokens []common.Address, _reserveMinReturnAmounts []*big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.RemoveLiquidity(&_ConvertertypeOne.TransactOpts, _amount, _reserveTokens, _reserveMinReturnAmounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xb127c0a5.
//
// Solidity: function removeLiquidity(uint256 _amount, address[] _reserveTokens, uint256[] _reserveMinReturnAmounts) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) RemoveLiquidity(_amount *big.Int, _reserveTokens []common.Address, _reserveMinReturnAmounts []*big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.RemoveLiquidity(&_ConvertertypeOne.TransactOpts, _amount, _reserveTokens, _reserveMinReturnAmounts)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) RestoreRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "restoreRegistry")
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_ConvertertypeOne *ConvertertypeOneSession) RestoreRegistry() (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.RestoreRegistry(&_ConvertertypeOne.TransactOpts)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) RestoreRegistry() (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.RestoreRegistry(&_ConvertertypeOne.TransactOpts)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) RestrictRegistryUpdate(opts *bind.TransactOpts, _onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "restrictRegistryUpdate", _onlyOwnerCanUpdateRegistry)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_ConvertertypeOne *ConvertertypeOneSession) RestrictRegistryUpdate(_onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.RestrictRegistryUpdate(&_ConvertertypeOne.TransactOpts, _onlyOwnerCanUpdateRegistry)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) RestrictRegistryUpdate(_onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.RestrictRegistryUpdate(&_ConvertertypeOne.TransactOpts, _onlyOwnerCanUpdateRegistry)
}

// SetConversionFee is a paid mutator transaction binding the contract method 0xecbca55d.
//
// Solidity: function setConversionFee(uint32 _conversionFee) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) SetConversionFee(opts *bind.TransactOpts, _conversionFee uint32) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "setConversionFee", _conversionFee)
}

// SetConversionFee is a paid mutator transaction binding the contract method 0xecbca55d.
//
// Solidity: function setConversionFee(uint32 _conversionFee) returns()
func (_ConvertertypeOne *ConvertertypeOneSession) SetConversionFee(_conversionFee uint32) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.SetConversionFee(&_ConvertertypeOne.TransactOpts, _conversionFee)
}

// SetConversionFee is a paid mutator transaction binding the contract method 0xecbca55d.
//
// Solidity: function setConversionFee(uint32 _conversionFee) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) SetConversionFee(_conversionFee uint32) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.SetConversionFee(&_ConvertertypeOne.TransactOpts, _conversionFee)
}

// SetConversionWhitelist is a paid mutator transaction binding the contract method 0x4af80f0e.
//
// Solidity: function setConversionWhitelist(address _whitelist) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) SetConversionWhitelist(opts *bind.TransactOpts, _whitelist common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "setConversionWhitelist", _whitelist)
}

// SetConversionWhitelist is a paid mutator transaction binding the contract method 0x4af80f0e.
//
// Solidity: function setConversionWhitelist(address _whitelist) returns()
func (_ConvertertypeOne *ConvertertypeOneSession) SetConversionWhitelist(_whitelist common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.SetConversionWhitelist(&_ConvertertypeOne.TransactOpts, _whitelist)
}

// SetConversionWhitelist is a paid mutator transaction binding the contract method 0x4af80f0e.
//
// Solidity: function setConversionWhitelist(address _whitelist) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) SetConversionWhitelist(_whitelist common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.SetConversionWhitelist(&_ConvertertypeOne.TransactOpts, _whitelist)
}

// TransferAnchorOwnership is a paid mutator transaction binding the contract method 0x67b6d57c.
//
// Solidity: function transferAnchorOwnership(address _newOwner) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) TransferAnchorOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "transferAnchorOwnership", _newOwner)
}

// TransferAnchorOwnership is a paid mutator transaction binding the contract method 0x67b6d57c.
//
// Solidity: function transferAnchorOwnership(address _newOwner) returns()
func (_ConvertertypeOne *ConvertertypeOneSession) TransferAnchorOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.TransferAnchorOwnership(&_ConvertertypeOne.TransactOpts, _newOwner)
}

// TransferAnchorOwnership is a paid mutator transaction binding the contract method 0x67b6d57c.
//
// Solidity: function transferAnchorOwnership(address _newOwner) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) TransferAnchorOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.TransferAnchorOwnership(&_ConvertertypeOne.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ConvertertypeOne *ConvertertypeOneSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.TransferOwnership(&_ConvertertypeOne.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.TransferOwnership(&_ConvertertypeOne.TransactOpts, _newOwner)
}

// TransferTokenOwnership is a paid mutator transaction binding the contract method 0x21e6b53d.
//
// Solidity: function transferTokenOwnership(address _newOwner) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) TransferTokenOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "transferTokenOwnership", _newOwner)
}

// TransferTokenOwnership is a paid mutator transaction binding the contract method 0x21e6b53d.
//
// Solidity: function transferTokenOwnership(address _newOwner) returns()
func (_ConvertertypeOne *ConvertertypeOneSession) TransferTokenOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.TransferTokenOwnership(&_ConvertertypeOne.TransactOpts, _newOwner)
}

// TransferTokenOwnership is a paid mutator transaction binding the contract method 0x21e6b53d.
//
// Solidity: function transferTokenOwnership(address _newOwner) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) TransferTokenOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.TransferTokenOwnership(&_ConvertertypeOne.TransactOpts, _newOwner)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) UpdateRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "updateRegistry")
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_ConvertertypeOne *ConvertertypeOneSession) UpdateRegistry() (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.UpdateRegistry(&_ConvertertypeOne.TransactOpts)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) UpdateRegistry() (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.UpdateRegistry(&_ConvertertypeOne.TransactOpts)
}

// Upgrade is a paid mutator transaction binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) Upgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "upgrade")
}

// Upgrade is a paid mutator transaction binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() returns()
func (_ConvertertypeOne *ConvertertypeOneSession) Upgrade() (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.Upgrade(&_ConvertertypeOne.TransactOpts)
}

// Upgrade is a paid mutator transaction binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) Upgrade() (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.Upgrade(&_ConvertertypeOne.TransactOpts)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x690d8320.
//
// Solidity: function withdrawETH(address _to) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) WithdrawETH(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "withdrawETH", _to)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x690d8320.
//
// Solidity: function withdrawETH(address _to) returns()
func (_ConvertertypeOne *ConvertertypeOneSession) WithdrawETH(_to common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.WithdrawETH(&_ConvertertypeOne.TransactOpts, _to)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x690d8320.
//
// Solidity: function withdrawETH(address _to) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) WithdrawETH(_to common.Address) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.WithdrawETH(&_ConvertertypeOne.TransactOpts, _to)
}

// WithdrawFromAnchor is a paid mutator transaction binding the contract method 0x395900d4.
//
// Solidity: function withdrawFromAnchor(address _token, address _to, uint256 _amount) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) WithdrawFromAnchor(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "withdrawFromAnchor", _token, _to, _amount)
}

// WithdrawFromAnchor is a paid mutator transaction binding the contract method 0x395900d4.
//
// Solidity: function withdrawFromAnchor(address _token, address _to, uint256 _amount) returns()
func (_ConvertertypeOne *ConvertertypeOneSession) WithdrawFromAnchor(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.WithdrawFromAnchor(&_ConvertertypeOne.TransactOpts, _token, _to, _amount)
}

// WithdrawFromAnchor is a paid mutator transaction binding the contract method 0x395900d4.
//
// Solidity: function withdrawFromAnchor(address _token, address _to, uint256 _amount) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) WithdrawFromAnchor(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.WithdrawFromAnchor(&_ConvertertypeOne.TransactOpts, _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) WithdrawTokens(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.Transact(opts, "withdrawTokens", _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_ConvertertypeOne *ConvertertypeOneSession) WithdrawTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.WithdrawTokens(&_ConvertertypeOne.TransactOpts, _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) WithdrawTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.WithdrawTokens(&_ConvertertypeOne.TransactOpts, _token, _to, _amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ConvertertypeOne *ConvertertypeOneTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ConvertertypeOne.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ConvertertypeOne *ConvertertypeOneSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.Fallback(&_ConvertertypeOne.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ConvertertypeOne *ConvertertypeOneTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ConvertertypeOne.Contract.Fallback(&_ConvertertypeOne.TransactOpts, calldata)
}

// ConvertertypeOneActivationIterator is returned from FilterActivation and is used to iterate over the raw logs and unpacked data for Activation events raised by the ConvertertypeOne contract.
type ConvertertypeOneActivationIterator struct {
	Event *ConvertertypeOneActivation // Event containing the contract specifics and raw log

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
func (it *ConvertertypeOneActivationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvertertypeOneActivation)
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
		it.Event = new(ConvertertypeOneActivation)
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
func (it *ConvertertypeOneActivationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvertertypeOneActivationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvertertypeOneActivation represents a Activation event raised by the ConvertertypeOne contract.
type ConvertertypeOneActivation struct {
	Type      uint16
	Anchor    common.Address
	Activated bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterActivation is a free log retrieval operation binding the contract event 0x6b08c2e2c9969e55a647a764db9b554d64dc42f1a704da11a6d5b129ad163f2c.
//
// Solidity: event Activation(uint16 indexed _type, address indexed _anchor, bool indexed _activated)
func (_ConvertertypeOne *ConvertertypeOneFilterer) FilterActivation(opts *bind.FilterOpts, _type []uint16, _anchor []common.Address, _activated []bool) (*ConvertertypeOneActivationIterator, error) {

	var _typeRule []interface{}
	for _, _typeItem := range _type {
		_typeRule = append(_typeRule, _typeItem)
	}
	var _anchorRule []interface{}
	for _, _anchorItem := range _anchor {
		_anchorRule = append(_anchorRule, _anchorItem)
	}
	var _activatedRule []interface{}
	for _, _activatedItem := range _activated {
		_activatedRule = append(_activatedRule, _activatedItem)
	}

	logs, sub, err := _ConvertertypeOne.contract.FilterLogs(opts, "Activation", _typeRule, _anchorRule, _activatedRule)
	if err != nil {
		return nil, err
	}
	return &ConvertertypeOneActivationIterator{contract: _ConvertertypeOne.contract, event: "Activation", logs: logs, sub: sub}, nil
}

// WatchActivation is a free log subscription operation binding the contract event 0x6b08c2e2c9969e55a647a764db9b554d64dc42f1a704da11a6d5b129ad163f2c.
//
// Solidity: event Activation(uint16 indexed _type, address indexed _anchor, bool indexed _activated)
func (_ConvertertypeOne *ConvertertypeOneFilterer) WatchActivation(opts *bind.WatchOpts, sink chan<- *ConvertertypeOneActivation, _type []uint16, _anchor []common.Address, _activated []bool) (event.Subscription, error) {

	var _typeRule []interface{}
	for _, _typeItem := range _type {
		_typeRule = append(_typeRule, _typeItem)
	}
	var _anchorRule []interface{}
	for _, _anchorItem := range _anchor {
		_anchorRule = append(_anchorRule, _anchorItem)
	}
	var _activatedRule []interface{}
	for _, _activatedItem := range _activated {
		_activatedRule = append(_activatedRule, _activatedItem)
	}

	logs, sub, err := _ConvertertypeOne.contract.WatchLogs(opts, "Activation", _typeRule, _anchorRule, _activatedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvertertypeOneActivation)
				if err := _ConvertertypeOne.contract.UnpackLog(event, "Activation", log); err != nil {
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

// ParseActivation is a log parse operation binding the contract event 0x6b08c2e2c9969e55a647a764db9b554d64dc42f1a704da11a6d5b129ad163f2c.
//
// Solidity: event Activation(uint16 indexed _type, address indexed _anchor, bool indexed _activated)
func (_ConvertertypeOne *ConvertertypeOneFilterer) ParseActivation(log types.Log) (*ConvertertypeOneActivation, error) {
	event := new(ConvertertypeOneActivation)
	if err := _ConvertertypeOne.contract.UnpackLog(event, "Activation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvertertypeOneConversionIterator is returned from FilterConversion and is used to iterate over the raw logs and unpacked data for Conversion events raised by the ConvertertypeOne contract.
type ConvertertypeOneConversionIterator struct {
	Event *ConvertertypeOneConversion // Event containing the contract specifics and raw log

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
func (it *ConvertertypeOneConversionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvertertypeOneConversion)
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
		it.Event = new(ConvertertypeOneConversion)
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
func (it *ConvertertypeOneConversionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvertertypeOneConversionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvertertypeOneConversion represents a Conversion event raised by the ConvertertypeOne contract.
type ConvertertypeOneConversion struct {
	FromToken     common.Address
	ToToken       common.Address
	Trader        common.Address
	Amount        *big.Int
	Return        *big.Int
	ConversionFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConversion is a free log retrieval operation binding the contract event 0x276856b36cbc45526a0ba64f44611557a2a8b68662c5388e9fe6d72e86e1c8cb.
//
// Solidity: event Conversion(address indexed _fromToken, address indexed _toToken, address indexed _trader, uint256 _amount, uint256 _return, int256 _conversionFee)
func (_ConvertertypeOne *ConvertertypeOneFilterer) FilterConversion(opts *bind.FilterOpts, _fromToken []common.Address, _toToken []common.Address, _trader []common.Address) (*ConvertertypeOneConversionIterator, error) {

	var _fromTokenRule []interface{}
	for _, _fromTokenItem := range _fromToken {
		_fromTokenRule = append(_fromTokenRule, _fromTokenItem)
	}
	var _toTokenRule []interface{}
	for _, _toTokenItem := range _toToken {
		_toTokenRule = append(_toTokenRule, _toTokenItem)
	}
	var _traderRule []interface{}
	for _, _traderItem := range _trader {
		_traderRule = append(_traderRule, _traderItem)
	}

	logs, sub, err := _ConvertertypeOne.contract.FilterLogs(opts, "Conversion", _fromTokenRule, _toTokenRule, _traderRule)
	if err != nil {
		return nil, err
	}
	return &ConvertertypeOneConversionIterator{contract: _ConvertertypeOne.contract, event: "Conversion", logs: logs, sub: sub}, nil
}

// WatchConversion is a free log subscription operation binding the contract event 0x276856b36cbc45526a0ba64f44611557a2a8b68662c5388e9fe6d72e86e1c8cb.
//
// Solidity: event Conversion(address indexed _fromToken, address indexed _toToken, address indexed _trader, uint256 _amount, uint256 _return, int256 _conversionFee)
func (_ConvertertypeOne *ConvertertypeOneFilterer) WatchConversion(opts *bind.WatchOpts, sink chan<- *ConvertertypeOneConversion, _fromToken []common.Address, _toToken []common.Address, _trader []common.Address) (event.Subscription, error) {

	var _fromTokenRule []interface{}
	for _, _fromTokenItem := range _fromToken {
		_fromTokenRule = append(_fromTokenRule, _fromTokenItem)
	}
	var _toTokenRule []interface{}
	for _, _toTokenItem := range _toToken {
		_toTokenRule = append(_toTokenRule, _toTokenItem)
	}
	var _traderRule []interface{}
	for _, _traderItem := range _trader {
		_traderRule = append(_traderRule, _traderItem)
	}

	logs, sub, err := _ConvertertypeOne.contract.WatchLogs(opts, "Conversion", _fromTokenRule, _toTokenRule, _traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvertertypeOneConversion)
				if err := _ConvertertypeOne.contract.UnpackLog(event, "Conversion", log); err != nil {
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

// ParseConversion is a log parse operation binding the contract event 0x276856b36cbc45526a0ba64f44611557a2a8b68662c5388e9fe6d72e86e1c8cb.
//
// Solidity: event Conversion(address indexed _fromToken, address indexed _toToken, address indexed _trader, uint256 _amount, uint256 _return, int256 _conversionFee)
func (_ConvertertypeOne *ConvertertypeOneFilterer) ParseConversion(log types.Log) (*ConvertertypeOneConversion, error) {
	event := new(ConvertertypeOneConversion)
	if err := _ConvertertypeOne.contract.UnpackLog(event, "Conversion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvertertypeOneConversionFeeUpdateIterator is returned from FilterConversionFeeUpdate and is used to iterate over the raw logs and unpacked data for ConversionFeeUpdate events raised by the ConvertertypeOne contract.
type ConvertertypeOneConversionFeeUpdateIterator struct {
	Event *ConvertertypeOneConversionFeeUpdate // Event containing the contract specifics and raw log

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
func (it *ConvertertypeOneConversionFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvertertypeOneConversionFeeUpdate)
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
		it.Event = new(ConvertertypeOneConversionFeeUpdate)
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
func (it *ConvertertypeOneConversionFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvertertypeOneConversionFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvertertypeOneConversionFeeUpdate represents a ConversionFeeUpdate event raised by the ConvertertypeOne contract.
type ConvertertypeOneConversionFeeUpdate struct {
	PrevFee uint32
	NewFee  uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConversionFeeUpdate is a free log retrieval operation binding the contract event 0x81cd2ffb37dd237c0e4e2a3de5265fcf9deb43d3e7801e80db9f1ccfba7ee600.
//
// Solidity: event ConversionFeeUpdate(uint32 _prevFee, uint32 _newFee)
func (_ConvertertypeOne *ConvertertypeOneFilterer) FilterConversionFeeUpdate(opts *bind.FilterOpts) (*ConvertertypeOneConversionFeeUpdateIterator, error) {

	logs, sub, err := _ConvertertypeOne.contract.FilterLogs(opts, "ConversionFeeUpdate")
	if err != nil {
		return nil, err
	}
	return &ConvertertypeOneConversionFeeUpdateIterator{contract: _ConvertertypeOne.contract, event: "ConversionFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchConversionFeeUpdate is a free log subscription operation binding the contract event 0x81cd2ffb37dd237c0e4e2a3de5265fcf9deb43d3e7801e80db9f1ccfba7ee600.
//
// Solidity: event ConversionFeeUpdate(uint32 _prevFee, uint32 _newFee)
func (_ConvertertypeOne *ConvertertypeOneFilterer) WatchConversionFeeUpdate(opts *bind.WatchOpts, sink chan<- *ConvertertypeOneConversionFeeUpdate) (event.Subscription, error) {

	logs, sub, err := _ConvertertypeOne.contract.WatchLogs(opts, "ConversionFeeUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvertertypeOneConversionFeeUpdate)
				if err := _ConvertertypeOne.contract.UnpackLog(event, "ConversionFeeUpdate", log); err != nil {
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

// ParseConversionFeeUpdate is a log parse operation binding the contract event 0x81cd2ffb37dd237c0e4e2a3de5265fcf9deb43d3e7801e80db9f1ccfba7ee600.
//
// Solidity: event ConversionFeeUpdate(uint32 _prevFee, uint32 _newFee)
func (_ConvertertypeOne *ConvertertypeOneFilterer) ParseConversionFeeUpdate(log types.Log) (*ConvertertypeOneConversionFeeUpdate, error) {
	event := new(ConvertertypeOneConversionFeeUpdate)
	if err := _ConvertertypeOne.contract.UnpackLog(event, "ConversionFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvertertypeOneLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the ConvertertypeOne contract.
type ConvertertypeOneLiquidityAddedIterator struct {
	Event *ConvertertypeOneLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *ConvertertypeOneLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvertertypeOneLiquidityAdded)
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
		it.Event = new(ConvertertypeOneLiquidityAdded)
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
func (it *ConvertertypeOneLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvertertypeOneLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvertertypeOneLiquidityAdded represents a LiquidityAdded event raised by the ConvertertypeOne contract.
type ConvertertypeOneLiquidityAdded struct {
	Provider     common.Address
	ReserveToken common.Address
	Amount       *big.Int
	NewBalance   *big.Int
	NewSupply    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0x4a1a2a6176e9646d9e3157f7c2ab3c499f18337c0b0828cfb28e0a61de4a11f7.
//
// Solidity: event LiquidityAdded(address indexed _provider, address indexed _reserveToken, uint256 _amount, uint256 _newBalance, uint256 _newSupply)
func (_ConvertertypeOne *ConvertertypeOneFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, _provider []common.Address, _reserveToken []common.Address) (*ConvertertypeOneLiquidityAddedIterator, error) {

	var _providerRule []interface{}
	for _, _providerItem := range _provider {
		_providerRule = append(_providerRule, _providerItem)
	}
	var _reserveTokenRule []interface{}
	for _, _reserveTokenItem := range _reserveToken {
		_reserveTokenRule = append(_reserveTokenRule, _reserveTokenItem)
	}

	logs, sub, err := _ConvertertypeOne.contract.FilterLogs(opts, "LiquidityAdded", _providerRule, _reserveTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConvertertypeOneLiquidityAddedIterator{contract: _ConvertertypeOne.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0x4a1a2a6176e9646d9e3157f7c2ab3c499f18337c0b0828cfb28e0a61de4a11f7.
//
// Solidity: event LiquidityAdded(address indexed _provider, address indexed _reserveToken, uint256 _amount, uint256 _newBalance, uint256 _newSupply)
func (_ConvertertypeOne *ConvertertypeOneFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *ConvertertypeOneLiquidityAdded, _provider []common.Address, _reserveToken []common.Address) (event.Subscription, error) {

	var _providerRule []interface{}
	for _, _providerItem := range _provider {
		_providerRule = append(_providerRule, _providerItem)
	}
	var _reserveTokenRule []interface{}
	for _, _reserveTokenItem := range _reserveToken {
		_reserveTokenRule = append(_reserveTokenRule, _reserveTokenItem)
	}

	logs, sub, err := _ConvertertypeOne.contract.WatchLogs(opts, "LiquidityAdded", _providerRule, _reserveTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvertertypeOneLiquidityAdded)
				if err := _ConvertertypeOne.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

// ParseLiquidityAdded is a log parse operation binding the contract event 0x4a1a2a6176e9646d9e3157f7c2ab3c499f18337c0b0828cfb28e0a61de4a11f7.
//
// Solidity: event LiquidityAdded(address indexed _provider, address indexed _reserveToken, uint256 _amount, uint256 _newBalance, uint256 _newSupply)
func (_ConvertertypeOne *ConvertertypeOneFilterer) ParseLiquidityAdded(log types.Log) (*ConvertertypeOneLiquidityAdded, error) {
	event := new(ConvertertypeOneLiquidityAdded)
	if err := _ConvertertypeOne.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvertertypeOneLiquidityRemovedIterator is returned from FilterLiquidityRemoved and is used to iterate over the raw logs and unpacked data for LiquidityRemoved events raised by the ConvertertypeOne contract.
type ConvertertypeOneLiquidityRemovedIterator struct {
	Event *ConvertertypeOneLiquidityRemoved // Event containing the contract specifics and raw log

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
func (it *ConvertertypeOneLiquidityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvertertypeOneLiquidityRemoved)
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
		it.Event = new(ConvertertypeOneLiquidityRemoved)
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
func (it *ConvertertypeOneLiquidityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvertertypeOneLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvertertypeOneLiquidityRemoved represents a LiquidityRemoved event raised by the ConvertertypeOne contract.
type ConvertertypeOneLiquidityRemoved struct {
	Provider     common.Address
	ReserveToken common.Address
	Amount       *big.Int
	NewBalance   *big.Int
	NewSupply    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLiquidityRemoved is a free log retrieval operation binding the contract event 0xbc7d19d505c7ec4db83f3b51f19fb98c4c8a99922e7839d1ee608dfbee29501b.
//
// Solidity: event LiquidityRemoved(address indexed _provider, address indexed _reserveToken, uint256 _amount, uint256 _newBalance, uint256 _newSupply)
func (_ConvertertypeOne *ConvertertypeOneFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, _provider []common.Address, _reserveToken []common.Address) (*ConvertertypeOneLiquidityRemovedIterator, error) {

	var _providerRule []interface{}
	for _, _providerItem := range _provider {
		_providerRule = append(_providerRule, _providerItem)
	}
	var _reserveTokenRule []interface{}
	for _, _reserveTokenItem := range _reserveToken {
		_reserveTokenRule = append(_reserveTokenRule, _reserveTokenItem)
	}

	logs, sub, err := _ConvertertypeOne.contract.FilterLogs(opts, "LiquidityRemoved", _providerRule, _reserveTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConvertertypeOneLiquidityRemovedIterator{contract: _ConvertertypeOne.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

// WatchLiquidityRemoved is a free log subscription operation binding the contract event 0xbc7d19d505c7ec4db83f3b51f19fb98c4c8a99922e7839d1ee608dfbee29501b.
//
// Solidity: event LiquidityRemoved(address indexed _provider, address indexed _reserveToken, uint256 _amount, uint256 _newBalance, uint256 _newSupply)
func (_ConvertertypeOne *ConvertertypeOneFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *ConvertertypeOneLiquidityRemoved, _provider []common.Address, _reserveToken []common.Address) (event.Subscription, error) {

	var _providerRule []interface{}
	for _, _providerItem := range _provider {
		_providerRule = append(_providerRule, _providerItem)
	}
	var _reserveTokenRule []interface{}
	for _, _reserveTokenItem := range _reserveToken {
		_reserveTokenRule = append(_reserveTokenRule, _reserveTokenItem)
	}

	logs, sub, err := _ConvertertypeOne.contract.WatchLogs(opts, "LiquidityRemoved", _providerRule, _reserveTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvertertypeOneLiquidityRemoved)
				if err := _ConvertertypeOne.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

// ParseLiquidityRemoved is a log parse operation binding the contract event 0xbc7d19d505c7ec4db83f3b51f19fb98c4c8a99922e7839d1ee608dfbee29501b.
//
// Solidity: event LiquidityRemoved(address indexed _provider, address indexed _reserveToken, uint256 _amount, uint256 _newBalance, uint256 _newSupply)
func (_ConvertertypeOne *ConvertertypeOneFilterer) ParseLiquidityRemoved(log types.Log) (*ConvertertypeOneLiquidityRemoved, error) {
	event := new(ConvertertypeOneLiquidityRemoved)
	if err := _ConvertertypeOne.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvertertypeOneOwnerUpdateIterator is returned from FilterOwnerUpdate and is used to iterate over the raw logs and unpacked data for OwnerUpdate events raised by the ConvertertypeOne contract.
type ConvertertypeOneOwnerUpdateIterator struct {
	Event *ConvertertypeOneOwnerUpdate // Event containing the contract specifics and raw log

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
func (it *ConvertertypeOneOwnerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvertertypeOneOwnerUpdate)
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
		it.Event = new(ConvertertypeOneOwnerUpdate)
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
func (it *ConvertertypeOneOwnerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvertertypeOneOwnerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvertertypeOneOwnerUpdate represents a OwnerUpdate event raised by the ConvertertypeOne contract.
type ConvertertypeOneOwnerUpdate struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdate is a free log retrieval operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_ConvertertypeOne *ConvertertypeOneFilterer) FilterOwnerUpdate(opts *bind.FilterOpts, _prevOwner []common.Address, _newOwner []common.Address) (*ConvertertypeOneOwnerUpdateIterator, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ConvertertypeOne.contract.FilterLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConvertertypeOneOwnerUpdateIterator{contract: _ConvertertypeOne.contract, event: "OwnerUpdate", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdate is a free log subscription operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_ConvertertypeOne *ConvertertypeOneFilterer) WatchOwnerUpdate(opts *bind.WatchOpts, sink chan<- *ConvertertypeOneOwnerUpdate, _prevOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ConvertertypeOne.contract.WatchLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvertertypeOneOwnerUpdate)
				if err := _ConvertertypeOne.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
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
func (_ConvertertypeOne *ConvertertypeOneFilterer) ParseOwnerUpdate(log types.Log) (*ConvertertypeOneOwnerUpdate, error) {
	event := new(ConvertertypeOneOwnerUpdate)
	if err := _ConvertertypeOne.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvertertypeOnePriceDataUpdateIterator is returned from FilterPriceDataUpdate and is used to iterate over the raw logs and unpacked data for PriceDataUpdate events raised by the ConvertertypeOne contract.
type ConvertertypeOnePriceDataUpdateIterator struct {
	Event *ConvertertypeOnePriceDataUpdate // Event containing the contract specifics and raw log

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
func (it *ConvertertypeOnePriceDataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvertertypeOnePriceDataUpdate)
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
		it.Event = new(ConvertertypeOnePriceDataUpdate)
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
func (it *ConvertertypeOnePriceDataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvertertypeOnePriceDataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvertertypeOnePriceDataUpdate represents a PriceDataUpdate event raised by the ConvertertypeOne contract.
type ConvertertypeOnePriceDataUpdate struct {
	ConnectorToken   common.Address
	TokenSupply      *big.Int
	ConnectorBalance *big.Int
	ConnectorWeight  uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPriceDataUpdate is a free log retrieval operation binding the contract event 0x8a6a7f53b3c8fa1dc4b83e3f1be668c1b251ff8d44cdcb83eb3acec3fec6a788.
//
// Solidity: event PriceDataUpdate(address indexed _connectorToken, uint256 _tokenSupply, uint256 _connectorBalance, uint32 _connectorWeight)
func (_ConvertertypeOne *ConvertertypeOneFilterer) FilterPriceDataUpdate(opts *bind.FilterOpts, _connectorToken []common.Address) (*ConvertertypeOnePriceDataUpdateIterator, error) {

	var _connectorTokenRule []interface{}
	for _, _connectorTokenItem := range _connectorToken {
		_connectorTokenRule = append(_connectorTokenRule, _connectorTokenItem)
	}

	logs, sub, err := _ConvertertypeOne.contract.FilterLogs(opts, "PriceDataUpdate", _connectorTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConvertertypeOnePriceDataUpdateIterator{contract: _ConvertertypeOne.contract, event: "PriceDataUpdate", logs: logs, sub: sub}, nil
}

// WatchPriceDataUpdate is a free log subscription operation binding the contract event 0x8a6a7f53b3c8fa1dc4b83e3f1be668c1b251ff8d44cdcb83eb3acec3fec6a788.
//
// Solidity: event PriceDataUpdate(address indexed _connectorToken, uint256 _tokenSupply, uint256 _connectorBalance, uint32 _connectorWeight)
func (_ConvertertypeOne *ConvertertypeOneFilterer) WatchPriceDataUpdate(opts *bind.WatchOpts, sink chan<- *ConvertertypeOnePriceDataUpdate, _connectorToken []common.Address) (event.Subscription, error) {

	var _connectorTokenRule []interface{}
	for _, _connectorTokenItem := range _connectorToken {
		_connectorTokenRule = append(_connectorTokenRule, _connectorTokenItem)
	}

	logs, sub, err := _ConvertertypeOne.contract.WatchLogs(opts, "PriceDataUpdate", _connectorTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvertertypeOnePriceDataUpdate)
				if err := _ConvertertypeOne.contract.UnpackLog(event, "PriceDataUpdate", log); err != nil {
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

// ParsePriceDataUpdate is a log parse operation binding the contract event 0x8a6a7f53b3c8fa1dc4b83e3f1be668c1b251ff8d44cdcb83eb3acec3fec6a788.
//
// Solidity: event PriceDataUpdate(address indexed _connectorToken, uint256 _tokenSupply, uint256 _connectorBalance, uint32 _connectorWeight)
func (_ConvertertypeOne *ConvertertypeOneFilterer) ParsePriceDataUpdate(log types.Log) (*ConvertertypeOnePriceDataUpdate, error) {
	event := new(ConvertertypeOnePriceDataUpdate)
	if err := _ConvertertypeOne.contract.UnpackLog(event, "PriceDataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvertertypeOneTokenRateUpdateIterator is returned from FilterTokenRateUpdate and is used to iterate over the raw logs and unpacked data for TokenRateUpdate events raised by the ConvertertypeOne contract.
type ConvertertypeOneTokenRateUpdateIterator struct {
	Event *ConvertertypeOneTokenRateUpdate // Event containing the contract specifics and raw log

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
func (it *ConvertertypeOneTokenRateUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvertertypeOneTokenRateUpdate)
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
		it.Event = new(ConvertertypeOneTokenRateUpdate)
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
func (it *ConvertertypeOneTokenRateUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvertertypeOneTokenRateUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvertertypeOneTokenRateUpdate represents a TokenRateUpdate event raised by the ConvertertypeOne contract.
type ConvertertypeOneTokenRateUpdate struct {
	Token1 common.Address
	Token2 common.Address
	RateN  *big.Int
	RateD  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenRateUpdate is a free log retrieval operation binding the contract event 0x77f29993cf2c084e726f7e802da0719d6a0ade3e204badc7a3ffd57ecb768c24.
//
// Solidity: event TokenRateUpdate(address indexed _token1, address indexed _token2, uint256 _rateN, uint256 _rateD)
func (_ConvertertypeOne *ConvertertypeOneFilterer) FilterTokenRateUpdate(opts *bind.FilterOpts, _token1 []common.Address, _token2 []common.Address) (*ConvertertypeOneTokenRateUpdateIterator, error) {

	var _token1Rule []interface{}
	for _, _token1Item := range _token1 {
		_token1Rule = append(_token1Rule, _token1Item)
	}
	var _token2Rule []interface{}
	for _, _token2Item := range _token2 {
		_token2Rule = append(_token2Rule, _token2Item)
	}

	logs, sub, err := _ConvertertypeOne.contract.FilterLogs(opts, "TokenRateUpdate", _token1Rule, _token2Rule)
	if err != nil {
		return nil, err
	}
	return &ConvertertypeOneTokenRateUpdateIterator{contract: _ConvertertypeOne.contract, event: "TokenRateUpdate", logs: logs, sub: sub}, nil
}

// WatchTokenRateUpdate is a free log subscription operation binding the contract event 0x77f29993cf2c084e726f7e802da0719d6a0ade3e204badc7a3ffd57ecb768c24.
//
// Solidity: event TokenRateUpdate(address indexed _token1, address indexed _token2, uint256 _rateN, uint256 _rateD)
func (_ConvertertypeOne *ConvertertypeOneFilterer) WatchTokenRateUpdate(opts *bind.WatchOpts, sink chan<- *ConvertertypeOneTokenRateUpdate, _token1 []common.Address, _token2 []common.Address) (event.Subscription, error) {

	var _token1Rule []interface{}
	for _, _token1Item := range _token1 {
		_token1Rule = append(_token1Rule, _token1Item)
	}
	var _token2Rule []interface{}
	for _, _token2Item := range _token2 {
		_token2Rule = append(_token2Rule, _token2Item)
	}

	logs, sub, err := _ConvertertypeOne.contract.WatchLogs(opts, "TokenRateUpdate", _token1Rule, _token2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvertertypeOneTokenRateUpdate)
				if err := _ConvertertypeOne.contract.UnpackLog(event, "TokenRateUpdate", log); err != nil {
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

// ParseTokenRateUpdate is a log parse operation binding the contract event 0x77f29993cf2c084e726f7e802da0719d6a0ade3e204badc7a3ffd57ecb768c24.
//
// Solidity: event TokenRateUpdate(address indexed _token1, address indexed _token2, uint256 _rateN, uint256 _rateD)
func (_ConvertertypeOne *ConvertertypeOneFilterer) ParseTokenRateUpdate(log types.Log) (*ConvertertypeOneTokenRateUpdate, error) {
	event := new(ConvertertypeOneTokenRateUpdate)
	if err := _ConvertertypeOne.contract.UnpackLog(event, "TokenRateUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
