// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ConverterTypeZero

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

// ConverterTypeZeroABI is the input ABI used to generate the binding from.
const ConverterTypeZeroABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_onlyOwnerCanUpdateRegistry\",\"type\":\"bool\"}],\"name\":\"restrictRegistryUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_connectorToken\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"_virtualBalance\",\"type\":\"uint256\"}],\"name\":\"updateConnector\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"connectors\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bancorX\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserveToken\",\"type\":\"address\"}],\"name\":\"getReserveBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"connectorTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_fromToken\",\"type\":\"address\"},{\"name\":\"_toToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getReturn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferTokenOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"quickConvertPrioritized\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fromToken\",\"type\":\"address\"},{\"name\":\"_toToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"convertInternal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserveToken\",\"type\":\"address\"}],\"name\":\"getReserveRatio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"onlyOwnerCanUpdateRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptTokenOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_magnitude\",\"type\":\"uint8\"}],\"name\":\"getFinalAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_weight\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"bool\"}],\"name\":\"addConnector\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_whitelist\",\"type\":\"address\"}],\"name\":\"setConversionWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_conversionId\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"completeXConversion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"conversionFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fromToken\",\"type\":\"address\"},{\"name\":\"_toToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"change\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prevRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalReserveRatio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_ratio\",\"type\":\"uint32\"}],\"name\":\"addReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fromToken\",\"type\":\"address\"},{\"name\":\"_toToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"convert2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"connectorTokenCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserveToken\",\"type\":\"address\"},{\"name\":\"_sellAmount\",\"type\":\"uint256\"}],\"name\":\"getSaleReturn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fromToken\",\"type\":\"address\"},{\"name\":\"_toToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"convert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_fromConnectorToken\",\"type\":\"address\"},{\"name\":\"_toConnectorToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getCrossConnectorReturn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONVERTER_CONVERSION_WHITELIST\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserveToken\",\"type\":\"address\"},{\"name\":\"_virtualBalance\",\"type\":\"uint256\"}],\"name\":\"updateReserveVirtualBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxConversionFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveTokenCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserveToken\",\"type\":\"address\"},{\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"getPurchaseReturn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"quickConvertPrioritized2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restoreRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"conversionsEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"conversionWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_fromReserveToken\",\"type\":\"address\"},{\"name\":\"_toReserveToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getCrossReserveReturn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reserveTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"reserves\",\"outputs\":[{\"name\":\"virtualBalance\",\"type\":\"uint256\"},{\"name\":\"ratio\",\"type\":\"uint32\"},{\"name\":\"isVirtualBalanceEnabled\",\"type\":\"bool\"},{\"name\":\"isSaleEnabled\",\"type\":\"bool\"},{\"name\":\"isSet\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_connectorToken\",\"type\":\"address\"}],\"name\":\"getConnectorBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bancorX\",\"type\":\"address\"}],\"name\":\"setBancorX\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"quickConvert2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_conversionId\",\"type\":\"uint256\"}],\"name\":\"completeXConversion2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_conversionFee\",\"type\":\"uint32\"}],\"name\":\"setConversionFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"quickConvert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"claimTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_registry\",\"type\":\"address\"},{\"name\":\"_maxConversionFee\",\"type\":\"uint32\"},{\"name\":\"_reserveToken\",\"type\":\"address\"},{\"name\":\"_reserveRatio\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_fromToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_toToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_trader\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_return\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_conversionFee\",\"type\":\"int256\"}],\"name\":\"Conversion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_connectorToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_connectorBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_connectorWeight\",\"type\":\"uint32\"}],\"name\":\"PriceDataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prevFee\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"_newFee\",\"type\":\"uint32\"}],\"name\":\"ConversionFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// ConverterTypeZero is an auto generated Go binding around an Ethereum contract.
type ConverterTypeZero struct {
	ConverterTypeZeroCaller     // Read-only binding to the contract
	ConverterTypeZeroTransactor // Write-only binding to the contract
	ConverterTypeZeroFilterer   // Log filterer for contract events
}

// ConverterTypeZeroCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConverterTypeZeroCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConverterTypeZeroTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConverterTypeZeroTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConverterTypeZeroFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConverterTypeZeroFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConverterTypeZeroSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConverterTypeZeroSession struct {
	Contract     *ConverterTypeZero // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConverterTypeZeroCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConverterTypeZeroCallerSession struct {
	Contract *ConverterTypeZeroCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ConverterTypeZeroTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConverterTypeZeroTransactorSession struct {
	Contract     *ConverterTypeZeroTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ConverterTypeZeroRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConverterTypeZeroRaw struct {
	Contract *ConverterTypeZero // Generic contract binding to access the raw methods on
}

// ConverterTypeZeroCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConverterTypeZeroCallerRaw struct {
	Contract *ConverterTypeZeroCaller // Generic read-only contract binding to access the raw methods on
}

// ConverterTypeZeroTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConverterTypeZeroTransactorRaw struct {
	Contract *ConverterTypeZeroTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConverterTypeZero creates a new instance of ConverterTypeZero, bound to a specific deployed contract.
func NewConverterTypeZero(address common.Address, backend bind.ContractBackend) (*ConverterTypeZero, error) {
	contract, err := bindConverterTypeZero(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeZero{ConverterTypeZeroCaller: ConverterTypeZeroCaller{contract: contract}, ConverterTypeZeroTransactor: ConverterTypeZeroTransactor{contract: contract}, ConverterTypeZeroFilterer: ConverterTypeZeroFilterer{contract: contract}}, nil
}

// NewConverterTypeZeroCaller creates a new read-only instance of ConverterTypeZero, bound to a specific deployed contract.
func NewConverterTypeZeroCaller(address common.Address, caller bind.ContractCaller) (*ConverterTypeZeroCaller, error) {
	contract, err := bindConverterTypeZero(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeZeroCaller{contract: contract}, nil
}

// NewConverterTypeZeroTransactor creates a new write-only instance of ConverterTypeZero, bound to a specific deployed contract.
func NewConverterTypeZeroTransactor(address common.Address, transactor bind.ContractTransactor) (*ConverterTypeZeroTransactor, error) {
	contract, err := bindConverterTypeZero(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeZeroTransactor{contract: contract}, nil
}

// NewConverterTypeZeroFilterer creates a new log filterer instance of ConverterTypeZero, bound to a specific deployed contract.
func NewConverterTypeZeroFilterer(address common.Address, filterer bind.ContractFilterer) (*ConverterTypeZeroFilterer, error) {
	contract, err := bindConverterTypeZero(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeZeroFilterer{contract: contract}, nil
}

// bindConverterTypeZero binds a generic wrapper to an already deployed contract.
func bindConverterTypeZero(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConverterTypeZeroABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConverterTypeZero *ConverterTypeZeroRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConverterTypeZero.Contract.ConverterTypeZeroCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConverterTypeZero *ConverterTypeZeroRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.ConverterTypeZeroTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConverterTypeZero *ConverterTypeZeroRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.ConverterTypeZeroTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConverterTypeZero *ConverterTypeZeroCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConverterTypeZero.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConverterTypeZero *ConverterTypeZeroTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConverterTypeZero *ConverterTypeZeroTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.contract.Transact(opts, method, params...)
}

// CONVERTERCONVERSIONWHITELIST is a free data retrieval call binding the contract method 0x92d1abb7.
//
// Solidity: function CONVERTER_CONVERSION_WHITELIST() view returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroCaller) CONVERTERCONVERSIONWHITELIST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "CONVERTER_CONVERSION_WHITELIST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONVERTERCONVERSIONWHITELIST is a free data retrieval call binding the contract method 0x92d1abb7.
//
// Solidity: function CONVERTER_CONVERSION_WHITELIST() view returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) CONVERTERCONVERSIONWHITELIST() (*big.Int, error) {
	return _ConverterTypeZero.Contract.CONVERTERCONVERSIONWHITELIST(&_ConverterTypeZero.CallOpts)
}

// CONVERTERCONVERSIONWHITELIST is a free data retrieval call binding the contract method 0x92d1abb7.
//
// Solidity: function CONVERTER_CONVERSION_WHITELIST() view returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) CONVERTERCONVERSIONWHITELIST() (*big.Int, error) {
	return _ConverterTypeZero.Contract.CONVERTERCONVERSIONWHITELIST(&_ConverterTypeZero.CallOpts)
}

// BancorX is a free data retrieval call binding the contract method 0x1120a776.
//
// Solidity: function bancorX() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCaller) BancorX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "bancorX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BancorX is a free data retrieval call binding the contract method 0x1120a776.
//
// Solidity: function bancorX() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroSession) BancorX() (common.Address, error) {
	return _ConverterTypeZero.Contract.BancorX(&_ConverterTypeZero.CallOpts)
}

// BancorX is a free data retrieval call binding the contract method 0x1120a776.
//
// Solidity: function bancorX() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) BancorX() (common.Address, error) {
	return _ConverterTypeZero.Contract.BancorX(&_ConverterTypeZero.CallOpts)
}

// ConnectorTokenCount is a free data retrieval call binding the contract method 0x71f52bf3.
//
// Solidity: function connectorTokenCount() view returns(uint16)
func (_ConverterTypeZero *ConverterTypeZeroCaller) ConnectorTokenCount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "connectorTokenCount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ConnectorTokenCount is a free data retrieval call binding the contract method 0x71f52bf3.
//
// Solidity: function connectorTokenCount() view returns(uint16)
func (_ConverterTypeZero *ConverterTypeZeroSession) ConnectorTokenCount() (uint16, error) {
	return _ConverterTypeZero.Contract.ConnectorTokenCount(&_ConverterTypeZero.CallOpts)
}

// ConnectorTokenCount is a free data retrieval call binding the contract method 0x71f52bf3.
//
// Solidity: function connectorTokenCount() view returns(uint16)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) ConnectorTokenCount() (uint16, error) {
	return _ConverterTypeZero.Contract.ConnectorTokenCount(&_ConverterTypeZero.CallOpts)
}

// ConnectorTokens is a free data retrieval call binding the contract method 0x19b64015.
//
// Solidity: function connectorTokens(uint256 _index) view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCaller) ConnectorTokens(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "connectorTokens", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorTokens is a free data retrieval call binding the contract method 0x19b64015.
//
// Solidity: function connectorTokens(uint256 _index) view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroSession) ConnectorTokens(_index *big.Int) (common.Address, error) {
	return _ConverterTypeZero.Contract.ConnectorTokens(&_ConverterTypeZero.CallOpts, _index)
}

// ConnectorTokens is a free data retrieval call binding the contract method 0x19b64015.
//
// Solidity: function connectorTokens(uint256 _index) view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) ConnectorTokens(_index *big.Int) (common.Address, error) {
	return _ConverterTypeZero.Contract.ConnectorTokens(&_ConverterTypeZero.CallOpts, _index)
}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address _address) view returns(uint256, uint32, bool, bool, bool)
func (_ConverterTypeZero *ConverterTypeZeroCaller) Connectors(opts *bind.CallOpts, _address common.Address) (*big.Int, uint32, bool, bool, bool, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "connectors", _address)

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
func (_ConverterTypeZero *ConverterTypeZeroSession) Connectors(_address common.Address) (*big.Int, uint32, bool, bool, bool, error) {
	return _ConverterTypeZero.Contract.Connectors(&_ConverterTypeZero.CallOpts, _address)
}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address _address) view returns(uint256, uint32, bool, bool, bool)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) Connectors(_address common.Address) (*big.Int, uint32, bool, bool, bool, error) {
	return _ConverterTypeZero.Contract.Connectors(&_ConverterTypeZero.CallOpts, _address)
}

// ConversionFee is a free data retrieval call binding the contract method 0x579cd3ca.
//
// Solidity: function conversionFee() view returns(uint32)
func (_ConverterTypeZero *ConverterTypeZeroCaller) ConversionFee(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "conversionFee")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ConversionFee is a free data retrieval call binding the contract method 0x579cd3ca.
//
// Solidity: function conversionFee() view returns(uint32)
func (_ConverterTypeZero *ConverterTypeZeroSession) ConversionFee() (uint32, error) {
	return _ConverterTypeZero.Contract.ConversionFee(&_ConverterTypeZero.CallOpts)
}

// ConversionFee is a free data retrieval call binding the contract method 0x579cd3ca.
//
// Solidity: function conversionFee() view returns(uint32)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) ConversionFee() (uint32, error) {
	return _ConverterTypeZero.Contract.ConversionFee(&_ConverterTypeZero.CallOpts)
}

// ConversionWhitelist is a free data retrieval call binding the contract method 0xc45d3d92.
//
// Solidity: function conversionWhitelist() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCaller) ConversionWhitelist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "conversionWhitelist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConversionWhitelist is a free data retrieval call binding the contract method 0xc45d3d92.
//
// Solidity: function conversionWhitelist() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroSession) ConversionWhitelist() (common.Address, error) {
	return _ConverterTypeZero.Contract.ConversionWhitelist(&_ConverterTypeZero.CallOpts)
}

// ConversionWhitelist is a free data retrieval call binding the contract method 0xc45d3d92.
//
// Solidity: function conversionWhitelist() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) ConversionWhitelist() (common.Address, error) {
	return _ConverterTypeZero.Contract.ConversionWhitelist(&_ConverterTypeZero.CallOpts)
}

// ConversionsEnabled is a free data retrieval call binding the contract method 0xbf754558.
//
// Solidity: function conversionsEnabled() view returns(bool)
func (_ConverterTypeZero *ConverterTypeZeroCaller) ConversionsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "conversionsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ConversionsEnabled is a free data retrieval call binding the contract method 0xbf754558.
//
// Solidity: function conversionsEnabled() view returns(bool)
func (_ConverterTypeZero *ConverterTypeZeroSession) ConversionsEnabled() (bool, error) {
	return _ConverterTypeZero.Contract.ConversionsEnabled(&_ConverterTypeZero.CallOpts)
}

// ConversionsEnabled is a free data retrieval call binding the contract method 0xbf754558.
//
// Solidity: function conversionsEnabled() view returns(bool)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) ConversionsEnabled() (bool, error) {
	return _ConverterTypeZero.Contract.ConversionsEnabled(&_ConverterTypeZero.CallOpts)
}

// GetConnectorBalance is a free data retrieval call binding the contract method 0xd8959512.
//
// Solidity: function getConnectorBalance(address _connectorToken) view returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroCaller) GetConnectorBalance(opts *bind.CallOpts, _connectorToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "getConnectorBalance", _connectorToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConnectorBalance is a free data retrieval call binding the contract method 0xd8959512.
//
// Solidity: function getConnectorBalance(address _connectorToken) view returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) GetConnectorBalance(_connectorToken common.Address) (*big.Int, error) {
	return _ConverterTypeZero.Contract.GetConnectorBalance(&_ConverterTypeZero.CallOpts, _connectorToken)
}

// GetConnectorBalance is a free data retrieval call binding the contract method 0xd8959512.
//
// Solidity: function getConnectorBalance(address _connectorToken) view returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) GetConnectorBalance(_connectorToken common.Address) (*big.Int, error) {
	return _ConverterTypeZero.Contract.GetConnectorBalance(&_ConverterTypeZero.CallOpts, _connectorToken)
}

// GetCrossConnectorReturn is a free data retrieval call binding the contract method 0x8e3047e0.
//
// Solidity: function getCrossConnectorReturn(address _fromConnectorToken, address _toConnectorToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeZero *ConverterTypeZeroCaller) GetCrossConnectorReturn(opts *bind.CallOpts, _fromConnectorToken common.Address, _toConnectorToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "getCrossConnectorReturn", _fromConnectorToken, _toConnectorToken, _amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCrossConnectorReturn is a free data retrieval call binding the contract method 0x8e3047e0.
//
// Solidity: function getCrossConnectorReturn(address _fromConnectorToken, address _toConnectorToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) GetCrossConnectorReturn(_fromConnectorToken common.Address, _toConnectorToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeZero.Contract.GetCrossConnectorReturn(&_ConverterTypeZero.CallOpts, _fromConnectorToken, _toConnectorToken, _amount)
}

// GetCrossConnectorReturn is a free data retrieval call binding the contract method 0x8e3047e0.
//
// Solidity: function getCrossConnectorReturn(address _fromConnectorToken, address _toConnectorToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) GetCrossConnectorReturn(_fromConnectorToken common.Address, _toConnectorToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeZero.Contract.GetCrossConnectorReturn(&_ConverterTypeZero.CallOpts, _fromConnectorToken, _toConnectorToken, _amount)
}

// GetCrossReserveReturn is a free data retrieval call binding the contract method 0xcf73266a.
//
// Solidity: function getCrossReserveReturn(address _fromReserveToken, address _toReserveToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeZero *ConverterTypeZeroCaller) GetCrossReserveReturn(opts *bind.CallOpts, _fromReserveToken common.Address, _toReserveToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "getCrossReserveReturn", _fromReserveToken, _toReserveToken, _amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCrossReserveReturn is a free data retrieval call binding the contract method 0xcf73266a.
//
// Solidity: function getCrossReserveReturn(address _fromReserveToken, address _toReserveToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) GetCrossReserveReturn(_fromReserveToken common.Address, _toReserveToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeZero.Contract.GetCrossReserveReturn(&_ConverterTypeZero.CallOpts, _fromReserveToken, _toReserveToken, _amount)
}

// GetCrossReserveReturn is a free data retrieval call binding the contract method 0xcf73266a.
//
// Solidity: function getCrossReserveReturn(address _fromReserveToken, address _toReserveToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) GetCrossReserveReturn(_fromReserveToken common.Address, _toReserveToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeZero.Contract.GetCrossReserveReturn(&_ConverterTypeZero.CallOpts, _fromReserveToken, _toReserveToken, _amount)
}

// GetFinalAmount is a free data retrieval call binding the contract method 0x3aa0145a.
//
// Solidity: function getFinalAmount(uint256 _amount, uint8 _magnitude) view returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroCaller) GetFinalAmount(opts *bind.CallOpts, _amount *big.Int, _magnitude uint8) (*big.Int, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "getFinalAmount", _amount, _magnitude)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFinalAmount is a free data retrieval call binding the contract method 0x3aa0145a.
//
// Solidity: function getFinalAmount(uint256 _amount, uint8 _magnitude) view returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) GetFinalAmount(_amount *big.Int, _magnitude uint8) (*big.Int, error) {
	return _ConverterTypeZero.Contract.GetFinalAmount(&_ConverterTypeZero.CallOpts, _amount, _magnitude)
}

// GetFinalAmount is a free data retrieval call binding the contract method 0x3aa0145a.
//
// Solidity: function getFinalAmount(uint256 _amount, uint8 _magnitude) view returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) GetFinalAmount(_amount *big.Int, _magnitude uint8) (*big.Int, error) {
	return _ConverterTypeZero.Contract.GetFinalAmount(&_ConverterTypeZero.CallOpts, _amount, _magnitude)
}

// GetPurchaseReturn is a free data retrieval call binding the contract method 0xa2c4c336.
//
// Solidity: function getPurchaseReturn(address _reserveToken, uint256 _depositAmount) view returns(uint256, uint256)
func (_ConverterTypeZero *ConverterTypeZeroCaller) GetPurchaseReturn(opts *bind.CallOpts, _reserveToken common.Address, _depositAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "getPurchaseReturn", _reserveToken, _depositAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPurchaseReturn is a free data retrieval call binding the contract method 0xa2c4c336.
//
// Solidity: function getPurchaseReturn(address _reserveToken, uint256 _depositAmount) view returns(uint256, uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) GetPurchaseReturn(_reserveToken common.Address, _depositAmount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeZero.Contract.GetPurchaseReturn(&_ConverterTypeZero.CallOpts, _reserveToken, _depositAmount)
}

// GetPurchaseReturn is a free data retrieval call binding the contract method 0xa2c4c336.
//
// Solidity: function getPurchaseReturn(address _reserveToken, uint256 _depositAmount) view returns(uint256, uint256)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) GetPurchaseReturn(_reserveToken common.Address, _depositAmount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeZero.Contract.GetPurchaseReturn(&_ConverterTypeZero.CallOpts, _reserveToken, _depositAmount)
}

// GetReserveBalance is a free data retrieval call binding the contract method 0x15226b54.
//
// Solidity: function getReserveBalance(address _reserveToken) view returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroCaller) GetReserveBalance(opts *bind.CallOpts, _reserveToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "getReserveBalance", _reserveToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveBalance is a free data retrieval call binding the contract method 0x15226b54.
//
// Solidity: function getReserveBalance(address _reserveToken) view returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) GetReserveBalance(_reserveToken common.Address) (*big.Int, error) {
	return _ConverterTypeZero.Contract.GetReserveBalance(&_ConverterTypeZero.CallOpts, _reserveToken)
}

// GetReserveBalance is a free data retrieval call binding the contract method 0x15226b54.
//
// Solidity: function getReserveBalance(address _reserveToken) view returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) GetReserveBalance(_reserveToken common.Address) (*big.Int, error) {
	return _ConverterTypeZero.Contract.GetReserveBalance(&_ConverterTypeZero.CallOpts, _reserveToken)
}

// GetReserveRatio is a free data retrieval call binding the contract method 0x2c12b446.
//
// Solidity: function getReserveRatio(address _reserveToken) view returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroCaller) GetReserveRatio(opts *bind.CallOpts, _reserveToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "getReserveRatio", _reserveToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveRatio is a free data retrieval call binding the contract method 0x2c12b446.
//
// Solidity: function getReserveRatio(address _reserveToken) view returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) GetReserveRatio(_reserveToken common.Address) (*big.Int, error) {
	return _ConverterTypeZero.Contract.GetReserveRatio(&_ConverterTypeZero.CallOpts, _reserveToken)
}

// GetReserveRatio is a free data retrieval call binding the contract method 0x2c12b446.
//
// Solidity: function getReserveRatio(address _reserveToken) view returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) GetReserveRatio(_reserveToken common.Address) (*big.Int, error) {
	return _ConverterTypeZero.Contract.GetReserveRatio(&_ConverterTypeZero.CallOpts, _reserveToken)
}

// GetReturn is a free data retrieval call binding the contract method 0x1e1401f8.
//
// Solidity: function getReturn(address _fromToken, address _toToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeZero *ConverterTypeZeroCaller) GetReturn(opts *bind.CallOpts, _fromToken common.Address, _toToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "getReturn", _fromToken, _toToken, _amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetReturn is a free data retrieval call binding the contract method 0x1e1401f8.
//
// Solidity: function getReturn(address _fromToken, address _toToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) GetReturn(_fromToken common.Address, _toToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeZero.Contract.GetReturn(&_ConverterTypeZero.CallOpts, _fromToken, _toToken, _amount)
}

// GetReturn is a free data retrieval call binding the contract method 0x1e1401f8.
//
// Solidity: function getReturn(address _fromToken, address _toToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) GetReturn(_fromToken common.Address, _toToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeZero.Contract.GetReturn(&_ConverterTypeZero.CallOpts, _fromToken, _toToken, _amount)
}

// GetSaleReturn is a free data retrieval call binding the contract method 0x72b44b2c.
//
// Solidity: function getSaleReturn(address _reserveToken, uint256 _sellAmount) view returns(uint256, uint256)
func (_ConverterTypeZero *ConverterTypeZeroCaller) GetSaleReturn(opts *bind.CallOpts, _reserveToken common.Address, _sellAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "getSaleReturn", _reserveToken, _sellAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetSaleReturn is a free data retrieval call binding the contract method 0x72b44b2c.
//
// Solidity: function getSaleReturn(address _reserveToken, uint256 _sellAmount) view returns(uint256, uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) GetSaleReturn(_reserveToken common.Address, _sellAmount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeZero.Contract.GetSaleReturn(&_ConverterTypeZero.CallOpts, _reserveToken, _sellAmount)
}

// GetSaleReturn is a free data retrieval call binding the contract method 0x72b44b2c.
//
// Solidity: function getSaleReturn(address _reserveToken, uint256 _sellAmount) view returns(uint256, uint256)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) GetSaleReturn(_reserveToken common.Address, _sellAmount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeZero.Contract.GetSaleReturn(&_ConverterTypeZero.CallOpts, _reserveToken, _sellAmount)
}

// MaxConversionFee is a free data retrieval call binding the contract method 0x94c275ad.
//
// Solidity: function maxConversionFee() view returns(uint32)
func (_ConverterTypeZero *ConverterTypeZeroCaller) MaxConversionFee(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "maxConversionFee")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxConversionFee is a free data retrieval call binding the contract method 0x94c275ad.
//
// Solidity: function maxConversionFee() view returns(uint32)
func (_ConverterTypeZero *ConverterTypeZeroSession) MaxConversionFee() (uint32, error) {
	return _ConverterTypeZero.Contract.MaxConversionFee(&_ConverterTypeZero.CallOpts)
}

// MaxConversionFee is a free data retrieval call binding the contract method 0x94c275ad.
//
// Solidity: function maxConversionFee() view returns(uint32)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) MaxConversionFee() (uint32, error) {
	return _ConverterTypeZero.Contract.MaxConversionFee(&_ConverterTypeZero.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "newOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroSession) NewOwner() (common.Address, error) {
	return _ConverterTypeZero.Contract.NewOwner(&_ConverterTypeZero.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) NewOwner() (common.Address, error) {
	return _ConverterTypeZero.Contract.NewOwner(&_ConverterTypeZero.CallOpts)
}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_ConverterTypeZero *ConverterTypeZeroCaller) OnlyOwnerCanUpdateRegistry(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "onlyOwnerCanUpdateRegistry")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_ConverterTypeZero *ConverterTypeZeroSession) OnlyOwnerCanUpdateRegistry() (bool, error) {
	return _ConverterTypeZero.Contract.OnlyOwnerCanUpdateRegistry(&_ConverterTypeZero.CallOpts)
}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) OnlyOwnerCanUpdateRegistry() (bool, error) {
	return _ConverterTypeZero.Contract.OnlyOwnerCanUpdateRegistry(&_ConverterTypeZero.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroSession) Owner() (common.Address, error) {
	return _ConverterTypeZero.Contract.Owner(&_ConverterTypeZero.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) Owner() (common.Address, error) {
	return _ConverterTypeZero.Contract.Owner(&_ConverterTypeZero.CallOpts)
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCaller) PrevRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "prevRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroSession) PrevRegistry() (common.Address, error) {
	return _ConverterTypeZero.Contract.PrevRegistry(&_ConverterTypeZero.CallOpts)
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) PrevRegistry() (common.Address, error) {
	return _ConverterTypeZero.Contract.PrevRegistry(&_ConverterTypeZero.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroSession) Registry() (common.Address, error) {
	return _ConverterTypeZero.Contract.Registry(&_ConverterTypeZero.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) Registry() (common.Address, error) {
	return _ConverterTypeZero.Contract.Registry(&_ConverterTypeZero.CallOpts)
}

// ReserveTokenCount is a free data retrieval call binding the contract method 0x9b99a8e2.
//
// Solidity: function reserveTokenCount() view returns(uint16)
func (_ConverterTypeZero *ConverterTypeZeroCaller) ReserveTokenCount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "reserveTokenCount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ReserveTokenCount is a free data retrieval call binding the contract method 0x9b99a8e2.
//
// Solidity: function reserveTokenCount() view returns(uint16)
func (_ConverterTypeZero *ConverterTypeZeroSession) ReserveTokenCount() (uint16, error) {
	return _ConverterTypeZero.Contract.ReserveTokenCount(&_ConverterTypeZero.CallOpts)
}

// ReserveTokenCount is a free data retrieval call binding the contract method 0x9b99a8e2.
//
// Solidity: function reserveTokenCount() view returns(uint16)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) ReserveTokenCount() (uint16, error) {
	return _ConverterTypeZero.Contract.ReserveTokenCount(&_ConverterTypeZero.CallOpts)
}

// ReserveTokens is a free data retrieval call binding the contract method 0xd031370b.
//
// Solidity: function reserveTokens(uint256 ) view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCaller) ReserveTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "reserveTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReserveTokens is a free data retrieval call binding the contract method 0xd031370b.
//
// Solidity: function reserveTokens(uint256 ) view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroSession) ReserveTokens(arg0 *big.Int) (common.Address, error) {
	return _ConverterTypeZero.Contract.ReserveTokens(&_ConverterTypeZero.CallOpts, arg0)
}

// ReserveTokens is a free data retrieval call binding the contract method 0xd031370b.
//
// Solidity: function reserveTokens(uint256 ) view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) ReserveTokens(arg0 *big.Int) (common.Address, error) {
	return _ConverterTypeZero.Contract.ReserveTokens(&_ConverterTypeZero.CallOpts, arg0)
}

// Reserves is a free data retrieval call binding the contract method 0xd66bd524.
//
// Solidity: function reserves(address ) view returns(uint256 virtualBalance, uint32 ratio, bool isVirtualBalanceEnabled, bool isSaleEnabled, bool isSet)
func (_ConverterTypeZero *ConverterTypeZeroCaller) Reserves(opts *bind.CallOpts, arg0 common.Address) (struct {
	VirtualBalance          *big.Int
	Ratio                   uint32
	IsVirtualBalanceEnabled bool
	IsSaleEnabled           bool
	IsSet                   bool
}, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "reserves", arg0)

	outstruct := new(struct {
		VirtualBalance          *big.Int
		Ratio                   uint32
		IsVirtualBalanceEnabled bool
		IsSaleEnabled           bool
		IsSet                   bool
	})

	outstruct.VirtualBalance = out[0].(*big.Int)
	outstruct.Ratio = out[1].(uint32)
	outstruct.IsVirtualBalanceEnabled = out[2].(bool)
	outstruct.IsSaleEnabled = out[3].(bool)
	outstruct.IsSet = out[4].(bool)

	return *outstruct, err

}

// Reserves is a free data retrieval call binding the contract method 0xd66bd524.
//
// Solidity: function reserves(address ) view returns(uint256 virtualBalance, uint32 ratio, bool isVirtualBalanceEnabled, bool isSaleEnabled, bool isSet)
func (_ConverterTypeZero *ConverterTypeZeroSession) Reserves(arg0 common.Address) (struct {
	VirtualBalance          *big.Int
	Ratio                   uint32
	IsVirtualBalanceEnabled bool
	IsSaleEnabled           bool
	IsSet                   bool
}, error) {
	return _ConverterTypeZero.Contract.Reserves(&_ConverterTypeZero.CallOpts, arg0)
}

// Reserves is a free data retrieval call binding the contract method 0xd66bd524.
//
// Solidity: function reserves(address ) view returns(uint256 virtualBalance, uint32 ratio, bool isVirtualBalanceEnabled, bool isSaleEnabled, bool isSet)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) Reserves(arg0 common.Address) (struct {
	VirtualBalance          *big.Int
	Ratio                   uint32
	IsVirtualBalanceEnabled bool
	IsSaleEnabled           bool
	IsSet                   bool
}, error) {
	return _ConverterTypeZero.Contract.Reserves(&_ConverterTypeZero.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroSession) Token() (common.Address, error) {
	return _ConverterTypeZero.Contract.Token(&_ConverterTypeZero.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) Token() (common.Address, error) {
	return _ConverterTypeZero.Contract.Token(&_ConverterTypeZero.CallOpts)
}

// TotalReserveRatio is a free data retrieval call binding the contract method 0x6520d6fb.
//
// Solidity: function totalReserveRatio() view returns(uint32)
func (_ConverterTypeZero *ConverterTypeZeroCaller) TotalReserveRatio(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "totalReserveRatio")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalReserveRatio is a free data retrieval call binding the contract method 0x6520d6fb.
//
// Solidity: function totalReserveRatio() view returns(uint32)
func (_ConverterTypeZero *ConverterTypeZeroSession) TotalReserveRatio() (uint32, error) {
	return _ConverterTypeZero.Contract.TotalReserveRatio(&_ConverterTypeZero.CallOpts)
}

// TotalReserveRatio is a free data retrieval call binding the contract method 0x6520d6fb.
//
// Solidity: function totalReserveRatio() view returns(uint32)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) TotalReserveRatio() (uint32, error) {
	return _ConverterTypeZero.Contract.TotalReserveRatio(&_ConverterTypeZero.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_ConverterTypeZero *ConverterTypeZeroCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConverterTypeZero.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_ConverterTypeZero *ConverterTypeZeroSession) Version() (uint16, error) {
	return _ConverterTypeZero.Contract.Version(&_ConverterTypeZero.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_ConverterTypeZero *ConverterTypeZeroCallerSession) Version() (uint16, error) {
	return _ConverterTypeZero.Contract.Version(&_ConverterTypeZero.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) AcceptOwnership() (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.AcceptOwnership(&_ConverterTypeZero.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.AcceptOwnership(&_ConverterTypeZero.TransactOpts)
}

// AcceptTokenOwnership is a paid mutator transaction binding the contract method 0x38a5e016.
//
// Solidity: function acceptTokenOwnership() returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) AcceptTokenOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "acceptTokenOwnership")
}

// AcceptTokenOwnership is a paid mutator transaction binding the contract method 0x38a5e016.
//
// Solidity: function acceptTokenOwnership() returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) AcceptTokenOwnership() (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.AcceptTokenOwnership(&_ConverterTypeZero.TransactOpts)
}

// AcceptTokenOwnership is a paid mutator transaction binding the contract method 0x38a5e016.
//
// Solidity: function acceptTokenOwnership() returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) AcceptTokenOwnership() (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.AcceptTokenOwnership(&_ConverterTypeZero.TransactOpts)
}

// AddConnector is a paid mutator transaction binding the contract method 0x3f4d2fc2.
//
// Solidity: function addConnector(address _token, uint32 _weight, bool ) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) AddConnector(opts *bind.TransactOpts, _token common.Address, _weight uint32, arg2 bool) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "addConnector", _token, _weight, arg2)
}

// AddConnector is a paid mutator transaction binding the contract method 0x3f4d2fc2.
//
// Solidity: function addConnector(address _token, uint32 _weight, bool ) returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) AddConnector(_token common.Address, _weight uint32, arg2 bool) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.AddConnector(&_ConverterTypeZero.TransactOpts, _token, _weight, arg2)
}

// AddConnector is a paid mutator transaction binding the contract method 0x3f4d2fc2.
//
// Solidity: function addConnector(address _token, uint32 _weight, bool ) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) AddConnector(_token common.Address, _weight uint32, arg2 bool) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.AddConnector(&_ConverterTypeZero.TransactOpts, _token, _weight, arg2)
}

// AddReserve is a paid mutator transaction binding the contract method 0x6a49d2c4.
//
// Solidity: function addReserve(address _token, uint32 _ratio) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) AddReserve(opts *bind.TransactOpts, _token common.Address, _ratio uint32) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "addReserve", _token, _ratio)
}

// AddReserve is a paid mutator transaction binding the contract method 0x6a49d2c4.
//
// Solidity: function addReserve(address _token, uint32 _ratio) returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) AddReserve(_token common.Address, _ratio uint32) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.AddReserve(&_ConverterTypeZero.TransactOpts, _token, _ratio)
}

// AddReserve is a paid mutator transaction binding the contract method 0x6a49d2c4.
//
// Solidity: function addReserve(address _token, uint32 _ratio) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) AddReserve(_token common.Address, _ratio uint32) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.AddReserve(&_ConverterTypeZero.TransactOpts, _token, _ratio)
}

// Change is a paid mutator transaction binding the contract method 0x5e5144eb.
//
// Solidity: function change(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactor) Change(opts *bind.TransactOpts, _fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "change", _fromToken, _toToken, _amount, _minReturn)
}

// Change is a paid mutator transaction binding the contract method 0x5e5144eb.
//
// Solidity: function change(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) Change(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.Change(&_ConverterTypeZero.TransactOpts, _fromToken, _toToken, _amount, _minReturn)
}

// Change is a paid mutator transaction binding the contract method 0x5e5144eb.
//
// Solidity: function change(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) Change(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.Change(&_ConverterTypeZero.TransactOpts, _fromToken, _toToken, _amount, _minReturn)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xfe417fa5.
//
// Solidity: function claimTokens(address _from, uint256 _amount) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) ClaimTokens(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "claimTokens", _from, _amount)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xfe417fa5.
//
// Solidity: function claimTokens(address _from, uint256 _amount) returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) ClaimTokens(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.ClaimTokens(&_ConverterTypeZero.TransactOpts, _from, _amount)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xfe417fa5.
//
// Solidity: function claimTokens(address _from, uint256 _amount) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) ClaimTokens(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.ClaimTokens(&_ConverterTypeZero.TransactOpts, _from, _amount)
}

// CompleteXConversion is a paid mutator transaction binding the contract method 0x50057351.
//
// Solidity: function completeXConversion(address[] _path, uint256 _minReturn, uint256 _conversionId, uint256 , uint8 , bytes32 , bytes32 ) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactor) CompleteXConversion(opts *bind.TransactOpts, _path []common.Address, _minReturn *big.Int, _conversionId *big.Int, arg3 *big.Int, arg4 uint8, arg5 [32]byte, arg6 [32]byte) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "completeXConversion", _path, _minReturn, _conversionId, arg3, arg4, arg5, arg6)
}

// CompleteXConversion is a paid mutator transaction binding the contract method 0x50057351.
//
// Solidity: function completeXConversion(address[] _path, uint256 _minReturn, uint256 _conversionId, uint256 , uint8 , bytes32 , bytes32 ) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) CompleteXConversion(_path []common.Address, _minReturn *big.Int, _conversionId *big.Int, arg3 *big.Int, arg4 uint8, arg5 [32]byte, arg6 [32]byte) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.CompleteXConversion(&_ConverterTypeZero.TransactOpts, _path, _minReturn, _conversionId, arg3, arg4, arg5, arg6)
}

// CompleteXConversion is a paid mutator transaction binding the contract method 0x50057351.
//
// Solidity: function completeXConversion(address[] _path, uint256 _minReturn, uint256 _conversionId, uint256 , uint8 , bytes32 , bytes32 ) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) CompleteXConversion(_path []common.Address, _minReturn *big.Int, _conversionId *big.Int, arg3 *big.Int, arg4 uint8, arg5 [32]byte, arg6 [32]byte) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.CompleteXConversion(&_ConverterTypeZero.TransactOpts, _path, _minReturn, _conversionId, arg3, arg4, arg5, arg6)
}

// CompleteXConversion2 is a paid mutator transaction binding the contract method 0xebf94700.
//
// Solidity: function completeXConversion2(address[] _path, uint256 _minReturn, uint256 _conversionId) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactor) CompleteXConversion2(opts *bind.TransactOpts, _path []common.Address, _minReturn *big.Int, _conversionId *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "completeXConversion2", _path, _minReturn, _conversionId)
}

// CompleteXConversion2 is a paid mutator transaction binding the contract method 0xebf94700.
//
// Solidity: function completeXConversion2(address[] _path, uint256 _minReturn, uint256 _conversionId) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) CompleteXConversion2(_path []common.Address, _minReturn *big.Int, _conversionId *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.CompleteXConversion2(&_ConverterTypeZero.TransactOpts, _path, _minReturn, _conversionId)
}

// CompleteXConversion2 is a paid mutator transaction binding the contract method 0xebf94700.
//
// Solidity: function completeXConversion2(address[] _path, uint256 _minReturn, uint256 _conversionId) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) CompleteXConversion2(_path []common.Address, _minReturn *big.Int, _conversionId *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.CompleteXConversion2(&_ConverterTypeZero.TransactOpts, _path, _minReturn, _conversionId)
}

// Convert is a paid mutator transaction binding the contract method 0x75892cf1.
//
// Solidity: function convert(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactor) Convert(opts *bind.TransactOpts, _fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "convert", _fromToken, _toToken, _amount, _minReturn)
}

// Convert is a paid mutator transaction binding the contract method 0x75892cf1.
//
// Solidity: function convert(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) Convert(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.Convert(&_ConverterTypeZero.TransactOpts, _fromToken, _toToken, _amount, _minReturn)
}

// Convert is a paid mutator transaction binding the contract method 0x75892cf1.
//
// Solidity: function convert(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) Convert(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.Convert(&_ConverterTypeZero.TransactOpts, _fromToken, _toToken, _amount, _minReturn)
}

// Convert2 is a paid mutator transaction binding the contract method 0x6ebf36c0.
//
// Solidity: function convert2(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactor) Convert2(opts *bind.TransactOpts, _fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "convert2", _fromToken, _toToken, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// Convert2 is a paid mutator transaction binding the contract method 0x6ebf36c0.
//
// Solidity: function convert2(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) Convert2(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.Convert2(&_ConverterTypeZero.TransactOpts, _fromToken, _toToken, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// Convert2 is a paid mutator transaction binding the contract method 0x6ebf36c0.
//
// Solidity: function convert2(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) Convert2(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.Convert2(&_ConverterTypeZero.TransactOpts, _fromToken, _toToken, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// ConvertInternal is a paid mutator transaction binding the contract method 0x2a2e2f0c.
//
// Solidity: function convertInternal(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactor) ConvertInternal(opts *bind.TransactOpts, _fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "convertInternal", _fromToken, _toToken, _amount, _minReturn)
}

// ConvertInternal is a paid mutator transaction binding the contract method 0x2a2e2f0c.
//
// Solidity: function convertInternal(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) ConvertInternal(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.ConvertInternal(&_ConverterTypeZero.TransactOpts, _fromToken, _toToken, _amount, _minReturn)
}

// ConvertInternal is a paid mutator transaction binding the contract method 0x2a2e2f0c.
//
// Solidity: function convertInternal(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) ConvertInternal(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.ConvertInternal(&_ConverterTypeZero.TransactOpts, _fromToken, _toToken, _amount, _minReturn)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 _amount) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) Fund(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "fund", _amount)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 _amount) returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) Fund(_amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.Fund(&_ConverterTypeZero.TransactOpts, _amount)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 _amount) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) Fund(_amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.Fund(&_ConverterTypeZero.TransactOpts, _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 _amount) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) Liquidate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "liquidate", _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 _amount) returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) Liquidate(_amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.Liquidate(&_ConverterTypeZero.TransactOpts, _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 _amount) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) Liquidate(_amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.Liquidate(&_ConverterTypeZero.TransactOpts, _amount)
}

// QuickConvert is a paid mutator transaction binding the contract method 0xf0843ba9.
//
// Solidity: function quickConvert(address[] _path, uint256 _amount, uint256 _minReturn) payable returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactor) QuickConvert(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "quickConvert", _path, _amount, _minReturn)
}

// QuickConvert is a paid mutator transaction binding the contract method 0xf0843ba9.
//
// Solidity: function quickConvert(address[] _path, uint256 _amount, uint256 _minReturn) payable returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) QuickConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.QuickConvert(&_ConverterTypeZero.TransactOpts, _path, _amount, _minReturn)
}

// QuickConvert is a paid mutator transaction binding the contract method 0xf0843ba9.
//
// Solidity: function quickConvert(address[] _path, uint256 _amount, uint256 _minReturn) payable returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) QuickConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.QuickConvert(&_ConverterTypeZero.TransactOpts, _path, _amount, _minReturn)
}

// QuickConvert2 is a paid mutator transaction binding the contract method 0xe4dd22f6.
//
// Solidity: function quickConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactor) QuickConvert2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "quickConvert2", _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// QuickConvert2 is a paid mutator transaction binding the contract method 0xe4dd22f6.
//
// Solidity: function quickConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) QuickConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.QuickConvert2(&_ConverterTypeZero.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// QuickConvert2 is a paid mutator transaction binding the contract method 0xe4dd22f6.
//
// Solidity: function quickConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) QuickConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.QuickConvert2(&_ConverterTypeZero.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// QuickConvertPrioritized is a paid mutator transaction binding the contract method 0x22742564.
//
// Solidity: function quickConvertPrioritized(address[] _path, uint256 _amount, uint256 _minReturn, uint256 , uint8 , bytes32 , bytes32 ) payable returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactor) QuickConvertPrioritized(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, arg3 *big.Int, arg4 uint8, arg5 [32]byte, arg6 [32]byte) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "quickConvertPrioritized", _path, _amount, _minReturn, arg3, arg4, arg5, arg6)
}

// QuickConvertPrioritized is a paid mutator transaction binding the contract method 0x22742564.
//
// Solidity: function quickConvertPrioritized(address[] _path, uint256 _amount, uint256 _minReturn, uint256 , uint8 , bytes32 , bytes32 ) payable returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) QuickConvertPrioritized(_path []common.Address, _amount *big.Int, _minReturn *big.Int, arg3 *big.Int, arg4 uint8, arg5 [32]byte, arg6 [32]byte) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.QuickConvertPrioritized(&_ConverterTypeZero.TransactOpts, _path, _amount, _minReturn, arg3, arg4, arg5, arg6)
}

// QuickConvertPrioritized is a paid mutator transaction binding the contract method 0x22742564.
//
// Solidity: function quickConvertPrioritized(address[] _path, uint256 _amount, uint256 _minReturn, uint256 , uint8 , bytes32 , bytes32 ) payable returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) QuickConvertPrioritized(_path []common.Address, _amount *big.Int, _minReturn *big.Int, arg3 *big.Int, arg4 uint8, arg5 [32]byte, arg6 [32]byte) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.QuickConvertPrioritized(&_ConverterTypeZero.TransactOpts, _path, _amount, _minReturn, arg3, arg4, arg5, arg6)
}

// QuickConvertPrioritized2 is a paid mutator transaction binding the contract method 0xb3a426d5.
//
// Solidity: function quickConvertPrioritized2(address[] _path, uint256 _amount, uint256 _minReturn, uint256[] , address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactor) QuickConvertPrioritized2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, arg3 []*big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "quickConvertPrioritized2", _path, _amount, _minReturn, arg3, _affiliateAccount, _affiliateFee)
}

// QuickConvertPrioritized2 is a paid mutator transaction binding the contract method 0xb3a426d5.
//
// Solidity: function quickConvertPrioritized2(address[] _path, uint256 _amount, uint256 _minReturn, uint256[] , address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroSession) QuickConvertPrioritized2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, arg3 []*big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.QuickConvertPrioritized2(&_ConverterTypeZero.TransactOpts, _path, _amount, _minReturn, arg3, _affiliateAccount, _affiliateFee)
}

// QuickConvertPrioritized2 is a paid mutator transaction binding the contract method 0xb3a426d5.
//
// Solidity: function quickConvertPrioritized2(address[] _path, uint256 _amount, uint256 _minReturn, uint256[] , address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) QuickConvertPrioritized2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, arg3 []*big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.QuickConvertPrioritized2(&_ConverterTypeZero.TransactOpts, _path, _amount, _minReturn, arg3, _affiliateAccount, _affiliateFee)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) RestoreRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "restoreRegistry")
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) RestoreRegistry() (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.RestoreRegistry(&_ConverterTypeZero.TransactOpts)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) RestoreRegistry() (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.RestoreRegistry(&_ConverterTypeZero.TransactOpts)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) RestrictRegistryUpdate(opts *bind.TransactOpts, _onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "restrictRegistryUpdate", _onlyOwnerCanUpdateRegistry)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) RestrictRegistryUpdate(_onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.RestrictRegistryUpdate(&_ConverterTypeZero.TransactOpts, _onlyOwnerCanUpdateRegistry)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) RestrictRegistryUpdate(_onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.RestrictRegistryUpdate(&_ConverterTypeZero.TransactOpts, _onlyOwnerCanUpdateRegistry)
}

// SetBancorX is a paid mutator transaction binding the contract method 0xd924f0c3.
//
// Solidity: function setBancorX(address _bancorX) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) SetBancorX(opts *bind.TransactOpts, _bancorX common.Address) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "setBancorX", _bancorX)
}

// SetBancorX is a paid mutator transaction binding the contract method 0xd924f0c3.
//
// Solidity: function setBancorX(address _bancorX) returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) SetBancorX(_bancorX common.Address) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.SetBancorX(&_ConverterTypeZero.TransactOpts, _bancorX)
}

// SetBancorX is a paid mutator transaction binding the contract method 0xd924f0c3.
//
// Solidity: function setBancorX(address _bancorX) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) SetBancorX(_bancorX common.Address) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.SetBancorX(&_ConverterTypeZero.TransactOpts, _bancorX)
}

// SetConversionFee is a paid mutator transaction binding the contract method 0xecbca55d.
//
// Solidity: function setConversionFee(uint32 _conversionFee) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) SetConversionFee(opts *bind.TransactOpts, _conversionFee uint32) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "setConversionFee", _conversionFee)
}

// SetConversionFee is a paid mutator transaction binding the contract method 0xecbca55d.
//
// Solidity: function setConversionFee(uint32 _conversionFee) returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) SetConversionFee(_conversionFee uint32) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.SetConversionFee(&_ConverterTypeZero.TransactOpts, _conversionFee)
}

// SetConversionFee is a paid mutator transaction binding the contract method 0xecbca55d.
//
// Solidity: function setConversionFee(uint32 _conversionFee) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) SetConversionFee(_conversionFee uint32) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.SetConversionFee(&_ConverterTypeZero.TransactOpts, _conversionFee)
}

// SetConversionWhitelist is a paid mutator transaction binding the contract method 0x4af80f0e.
//
// Solidity: function setConversionWhitelist(address _whitelist) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) SetConversionWhitelist(opts *bind.TransactOpts, _whitelist common.Address) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "setConversionWhitelist", _whitelist)
}

// SetConversionWhitelist is a paid mutator transaction binding the contract method 0x4af80f0e.
//
// Solidity: function setConversionWhitelist(address _whitelist) returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) SetConversionWhitelist(_whitelist common.Address) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.SetConversionWhitelist(&_ConverterTypeZero.TransactOpts, _whitelist)
}

// SetConversionWhitelist is a paid mutator transaction binding the contract method 0x4af80f0e.
//
// Solidity: function setConversionWhitelist(address _whitelist) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) SetConversionWhitelist(_whitelist common.Address) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.SetConversionWhitelist(&_ConverterTypeZero.TransactOpts, _whitelist)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.TransferOwnership(&_ConverterTypeZero.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.TransferOwnership(&_ConverterTypeZero.TransactOpts, _newOwner)
}

// TransferTokenOwnership is a paid mutator transaction binding the contract method 0x21e6b53d.
//
// Solidity: function transferTokenOwnership(address _newOwner) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) TransferTokenOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "transferTokenOwnership", _newOwner)
}

// TransferTokenOwnership is a paid mutator transaction binding the contract method 0x21e6b53d.
//
// Solidity: function transferTokenOwnership(address _newOwner) returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) TransferTokenOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.TransferTokenOwnership(&_ConverterTypeZero.TransactOpts, _newOwner)
}

// TransferTokenOwnership is a paid mutator transaction binding the contract method 0x21e6b53d.
//
// Solidity: function transferTokenOwnership(address _newOwner) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) TransferTokenOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.TransferTokenOwnership(&_ConverterTypeZero.TransactOpts, _newOwner)
}

// UpdateConnector is a paid mutator transaction binding the contract method 0x0ca78923.
//
// Solidity: function updateConnector(address _connectorToken, uint32 , bool , uint256 _virtualBalance) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) UpdateConnector(opts *bind.TransactOpts, _connectorToken common.Address, arg1 uint32, arg2 bool, _virtualBalance *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "updateConnector", _connectorToken, arg1, arg2, _virtualBalance)
}

// UpdateConnector is a paid mutator transaction binding the contract method 0x0ca78923.
//
// Solidity: function updateConnector(address _connectorToken, uint32 , bool , uint256 _virtualBalance) returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) UpdateConnector(_connectorToken common.Address, arg1 uint32, arg2 bool, _virtualBalance *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.UpdateConnector(&_ConverterTypeZero.TransactOpts, _connectorToken, arg1, arg2, _virtualBalance)
}

// UpdateConnector is a paid mutator transaction binding the contract method 0x0ca78923.
//
// Solidity: function updateConnector(address _connectorToken, uint32 , bool , uint256 _virtualBalance) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) UpdateConnector(_connectorToken common.Address, arg1 uint32, arg2 bool, _virtualBalance *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.UpdateConnector(&_ConverterTypeZero.TransactOpts, _connectorToken, arg1, arg2, _virtualBalance)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) UpdateRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "updateRegistry")
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) UpdateRegistry() (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.UpdateRegistry(&_ConverterTypeZero.TransactOpts)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) UpdateRegistry() (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.UpdateRegistry(&_ConverterTypeZero.TransactOpts)
}

// UpdateReserveVirtualBalance is a paid mutator transaction binding the contract method 0x935e2ae1.
//
// Solidity: function updateReserveVirtualBalance(address _reserveToken, uint256 _virtualBalance) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) UpdateReserveVirtualBalance(opts *bind.TransactOpts, _reserveToken common.Address, _virtualBalance *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "updateReserveVirtualBalance", _reserveToken, _virtualBalance)
}

// UpdateReserveVirtualBalance is a paid mutator transaction binding the contract method 0x935e2ae1.
//
// Solidity: function updateReserveVirtualBalance(address _reserveToken, uint256 _virtualBalance) returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) UpdateReserveVirtualBalance(_reserveToken common.Address, _virtualBalance *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.UpdateReserveVirtualBalance(&_ConverterTypeZero.TransactOpts, _reserveToken, _virtualBalance)
}

// UpdateReserveVirtualBalance is a paid mutator transaction binding the contract method 0x935e2ae1.
//
// Solidity: function updateReserveVirtualBalance(address _reserveToken, uint256 _virtualBalance) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) UpdateReserveVirtualBalance(_reserveToken common.Address, _virtualBalance *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.UpdateReserveVirtualBalance(&_ConverterTypeZero.TransactOpts, _reserveToken, _virtualBalance)
}

// Upgrade is a paid mutator transaction binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) Upgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "upgrade")
}

// Upgrade is a paid mutator transaction binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) Upgrade() (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.Upgrade(&_ConverterTypeZero.TransactOpts)
}

// Upgrade is a paid mutator transaction binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) Upgrade() (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.Upgrade(&_ConverterTypeZero.TransactOpts)
}

// WithdrawFromToken is a paid mutator transaction binding the contract method 0x41a5b33d.
//
// Solidity: function withdrawFromToken(address _token, address _to, uint256 _amount) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) WithdrawFromToken(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "withdrawFromToken", _token, _to, _amount)
}

// WithdrawFromToken is a paid mutator transaction binding the contract method 0x41a5b33d.
//
// Solidity: function withdrawFromToken(address _token, address _to, uint256 _amount) returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) WithdrawFromToken(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.WithdrawFromToken(&_ConverterTypeZero.TransactOpts, _token, _to, _amount)
}

// WithdrawFromToken is a paid mutator transaction binding the contract method 0x41a5b33d.
//
// Solidity: function withdrawFromToken(address _token, address _to, uint256 _amount) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) WithdrawFromToken(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.WithdrawFromToken(&_ConverterTypeZero.TransactOpts, _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactor) WithdrawTokens(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.contract.Transact(opts, "withdrawTokens", _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_ConverterTypeZero *ConverterTypeZeroSession) WithdrawTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.WithdrawTokens(&_ConverterTypeZero.TransactOpts, _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_ConverterTypeZero *ConverterTypeZeroTransactorSession) WithdrawTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeZero.Contract.WithdrawTokens(&_ConverterTypeZero.TransactOpts, _token, _to, _amount)
}

// ConverterTypeZeroConversionIterator is returned from FilterConversion and is used to iterate over the raw logs and unpacked data for Conversion events raised by the ConverterTypeZero contract.
type ConverterTypeZeroConversionIterator struct {
	Event *ConverterTypeZeroConversion // Event containing the contract specifics and raw log

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
func (it *ConverterTypeZeroConversionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeZeroConversion)
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
		it.Event = new(ConverterTypeZeroConversion)
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
func (it *ConverterTypeZeroConversionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeZeroConversionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeZeroConversion represents a Conversion event raised by the ConverterTypeZero contract.
type ConverterTypeZeroConversion struct {
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
func (_ConverterTypeZero *ConverterTypeZeroFilterer) FilterConversion(opts *bind.FilterOpts, _fromToken []common.Address, _toToken []common.Address, _trader []common.Address) (*ConverterTypeZeroConversionIterator, error) {

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

	logs, sub, err := _ConverterTypeZero.contract.FilterLogs(opts, "Conversion", _fromTokenRule, _toTokenRule, _traderRule)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeZeroConversionIterator{contract: _ConverterTypeZero.contract, event: "Conversion", logs: logs, sub: sub}, nil
}

// WatchConversion is a free log subscription operation binding the contract event 0x276856b36cbc45526a0ba64f44611557a2a8b68662c5388e9fe6d72e86e1c8cb.
//
// Solidity: event Conversion(address indexed _fromToken, address indexed _toToken, address indexed _trader, uint256 _amount, uint256 _return, int256 _conversionFee)
func (_ConverterTypeZero *ConverterTypeZeroFilterer) WatchConversion(opts *bind.WatchOpts, sink chan<- *ConverterTypeZeroConversion, _fromToken []common.Address, _toToken []common.Address, _trader []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ConverterTypeZero.contract.WatchLogs(opts, "Conversion", _fromTokenRule, _toTokenRule, _traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeZeroConversion)
				if err := _ConverterTypeZero.contract.UnpackLog(event, "Conversion", log); err != nil {
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
func (_ConverterTypeZero *ConverterTypeZeroFilterer) ParseConversion(log types.Log) (*ConverterTypeZeroConversion, error) {
	event := new(ConverterTypeZeroConversion)
	if err := _ConverterTypeZero.contract.UnpackLog(event, "Conversion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterTypeZeroConversionFeeUpdateIterator is returned from FilterConversionFeeUpdate and is used to iterate over the raw logs and unpacked data for ConversionFeeUpdate events raised by the ConverterTypeZero contract.
type ConverterTypeZeroConversionFeeUpdateIterator struct {
	Event *ConverterTypeZeroConversionFeeUpdate // Event containing the contract specifics and raw log

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
func (it *ConverterTypeZeroConversionFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeZeroConversionFeeUpdate)
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
		it.Event = new(ConverterTypeZeroConversionFeeUpdate)
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
func (it *ConverterTypeZeroConversionFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeZeroConversionFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeZeroConversionFeeUpdate represents a ConversionFeeUpdate event raised by the ConverterTypeZero contract.
type ConverterTypeZeroConversionFeeUpdate struct {
	PrevFee uint32
	NewFee  uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConversionFeeUpdate is a free log retrieval operation binding the contract event 0x81cd2ffb37dd237c0e4e2a3de5265fcf9deb43d3e7801e80db9f1ccfba7ee600.
//
// Solidity: event ConversionFeeUpdate(uint32 _prevFee, uint32 _newFee)
func (_ConverterTypeZero *ConverterTypeZeroFilterer) FilterConversionFeeUpdate(opts *bind.FilterOpts) (*ConverterTypeZeroConversionFeeUpdateIterator, error) {

	logs, sub, err := _ConverterTypeZero.contract.FilterLogs(opts, "ConversionFeeUpdate")
	if err != nil {
		return nil, err
	}
	return &ConverterTypeZeroConversionFeeUpdateIterator{contract: _ConverterTypeZero.contract, event: "ConversionFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchConversionFeeUpdate is a free log subscription operation binding the contract event 0x81cd2ffb37dd237c0e4e2a3de5265fcf9deb43d3e7801e80db9f1ccfba7ee600.
//
// Solidity: event ConversionFeeUpdate(uint32 _prevFee, uint32 _newFee)
func (_ConverterTypeZero *ConverterTypeZeroFilterer) WatchConversionFeeUpdate(opts *bind.WatchOpts, sink chan<- *ConverterTypeZeroConversionFeeUpdate) (event.Subscription, error) {

	logs, sub, err := _ConverterTypeZero.contract.WatchLogs(opts, "ConversionFeeUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeZeroConversionFeeUpdate)
				if err := _ConverterTypeZero.contract.UnpackLog(event, "ConversionFeeUpdate", log); err != nil {
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
func (_ConverterTypeZero *ConverterTypeZeroFilterer) ParseConversionFeeUpdate(log types.Log) (*ConverterTypeZeroConversionFeeUpdate, error) {
	event := new(ConverterTypeZeroConversionFeeUpdate)
	if err := _ConverterTypeZero.contract.UnpackLog(event, "ConversionFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterTypeZeroOwnerUpdateIterator is returned from FilterOwnerUpdate and is used to iterate over the raw logs and unpacked data for OwnerUpdate events raised by the ConverterTypeZero contract.
type ConverterTypeZeroOwnerUpdateIterator struct {
	Event *ConverterTypeZeroOwnerUpdate // Event containing the contract specifics and raw log

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
func (it *ConverterTypeZeroOwnerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeZeroOwnerUpdate)
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
		it.Event = new(ConverterTypeZeroOwnerUpdate)
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
func (it *ConverterTypeZeroOwnerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeZeroOwnerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeZeroOwnerUpdate represents a OwnerUpdate event raised by the ConverterTypeZero contract.
type ConverterTypeZeroOwnerUpdate struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdate is a free log retrieval operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_ConverterTypeZero *ConverterTypeZeroFilterer) FilterOwnerUpdate(opts *bind.FilterOpts, _prevOwner []common.Address, _newOwner []common.Address) (*ConverterTypeZeroOwnerUpdateIterator, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ConverterTypeZero.contract.FilterLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeZeroOwnerUpdateIterator{contract: _ConverterTypeZero.contract, event: "OwnerUpdate", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdate is a free log subscription operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_ConverterTypeZero *ConverterTypeZeroFilterer) WatchOwnerUpdate(opts *bind.WatchOpts, sink chan<- *ConverterTypeZeroOwnerUpdate, _prevOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ConverterTypeZero.contract.WatchLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeZeroOwnerUpdate)
				if err := _ConverterTypeZero.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
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
func (_ConverterTypeZero *ConverterTypeZeroFilterer) ParseOwnerUpdate(log types.Log) (*ConverterTypeZeroOwnerUpdate, error) {
	event := new(ConverterTypeZeroOwnerUpdate)
	if err := _ConverterTypeZero.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterTypeZeroPriceDataUpdateIterator is returned from FilterPriceDataUpdate and is used to iterate over the raw logs and unpacked data for PriceDataUpdate events raised by the ConverterTypeZero contract.
type ConverterTypeZeroPriceDataUpdateIterator struct {
	Event *ConverterTypeZeroPriceDataUpdate // Event containing the contract specifics and raw log

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
func (it *ConverterTypeZeroPriceDataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeZeroPriceDataUpdate)
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
		it.Event = new(ConverterTypeZeroPriceDataUpdate)
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
func (it *ConverterTypeZeroPriceDataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeZeroPriceDataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeZeroPriceDataUpdate represents a PriceDataUpdate event raised by the ConverterTypeZero contract.
type ConverterTypeZeroPriceDataUpdate struct {
	ConnectorToken   common.Address
	TokenSupply      *big.Int
	ConnectorBalance *big.Int
	ConnectorWeight  uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPriceDataUpdate is a free log retrieval operation binding the contract event 0x8a6a7f53b3c8fa1dc4b83e3f1be668c1b251ff8d44cdcb83eb3acec3fec6a788.
//
// Solidity: event PriceDataUpdate(address indexed _connectorToken, uint256 _tokenSupply, uint256 _connectorBalance, uint32 _connectorWeight)
func (_ConverterTypeZero *ConverterTypeZeroFilterer) FilterPriceDataUpdate(opts *bind.FilterOpts, _connectorToken []common.Address) (*ConverterTypeZeroPriceDataUpdateIterator, error) {

	var _connectorTokenRule []interface{}
	for _, _connectorTokenItem := range _connectorToken {
		_connectorTokenRule = append(_connectorTokenRule, _connectorTokenItem)
	}

	logs, sub, err := _ConverterTypeZero.contract.FilterLogs(opts, "PriceDataUpdate", _connectorTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeZeroPriceDataUpdateIterator{contract: _ConverterTypeZero.contract, event: "PriceDataUpdate", logs: logs, sub: sub}, nil
}

// WatchPriceDataUpdate is a free log subscription operation binding the contract event 0x8a6a7f53b3c8fa1dc4b83e3f1be668c1b251ff8d44cdcb83eb3acec3fec6a788.
//
// Solidity: event PriceDataUpdate(address indexed _connectorToken, uint256 _tokenSupply, uint256 _connectorBalance, uint32 _connectorWeight)
func (_ConverterTypeZero *ConverterTypeZeroFilterer) WatchPriceDataUpdate(opts *bind.WatchOpts, sink chan<- *ConverterTypeZeroPriceDataUpdate, _connectorToken []common.Address) (event.Subscription, error) {

	var _connectorTokenRule []interface{}
	for _, _connectorTokenItem := range _connectorToken {
		_connectorTokenRule = append(_connectorTokenRule, _connectorTokenItem)
	}

	logs, sub, err := _ConverterTypeZero.contract.WatchLogs(opts, "PriceDataUpdate", _connectorTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeZeroPriceDataUpdate)
				if err := _ConverterTypeZero.contract.UnpackLog(event, "PriceDataUpdate", log); err != nil {
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
func (_ConverterTypeZero *ConverterTypeZeroFilterer) ParsePriceDataUpdate(log types.Log) (*ConverterTypeZeroPriceDataUpdate, error) {
	event := new(ConverterTypeZeroPriceDataUpdate)
	if err := _ConverterTypeZero.contract.UnpackLog(event, "PriceDataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
