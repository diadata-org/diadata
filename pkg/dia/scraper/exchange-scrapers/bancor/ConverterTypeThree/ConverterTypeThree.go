// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ConverterTypeThree

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

// ConverterTypeThreeABI is the input ABI used to generate the binding from.
const ConverterTypeThreeABI = "[{\"inputs\":[{\"internalType\":\"contractIConverterAnchor\",\"name\":\"_anchor\",\"type\":\"address\"},{\"internalType\":\"contractIContractRegistry\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_maxConversionFee\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"_type\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"contractIConverterAnchor\",\"name\":\"_anchor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"_activated\",\"type\":\"bool\"}],\"name\":\"Activation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20Token\",\"name\":\"_fromToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20Token\",\"name\":\"_toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_return\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"_conversionFee\",\"type\":\"int256\"}],\"name\":\"Conversion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_prevFee\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_newFee\",\"type\":\"uint32\"}],\"name\":\"ConversionFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20Token\",\"name\":\"_reserveToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newSupply\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20Token\",\"name\":\"_reserveToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newSupply\",\"type\":\"uint256\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20Token\",\"name\":\"_token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20Token\",\"name\":\"_token2\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rateN\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rateD\",\"type\":\"uint256\"}],\"name\":\"TokenRateUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptAnchorOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptTokenOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reserve1Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserve2Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token[]\",\"name\":\"_reserveTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_reserveAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token[]\",\"name\":\"_reserveTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_reserveTokenIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveAmount\",\"type\":\"uint256\"}],\"name\":\"addLiquidityCost\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_reserveToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reserveAmount\",\"type\":\"uint256\"}],\"name\":\"addLiquidityReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_weight\",\"type\":\"uint32\"}],\"name\":\"addReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"anchor\",\"outputs\":[{\"internalType\":\"contractIConverterAnchor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"averageRateInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connectorTokenCount\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"connectorTokens\",\"outputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"connectors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"conversionFee\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_sourceToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Token\",\"name\":\"_targetToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"convert\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"converterType\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_connectorToken\",\"type\":\"address\"}],\"name\":\"getConnectorBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_sourceToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Token\",\"name\":\"_targetToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isV28OrHigher\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxConversionFee\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onlyOwnerCanUpdateRegistry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevRegistry\",\"outputs\":[{\"internalType\":\"contractIContractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"recentAverageRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIContractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserve1MinReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserve2MinReturn\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20Token[]\",\"name\":\"_reserveTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_reserveMinReturnAmounts\",\"type\":\"uint256[]\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20Token[]\",\"name\":\"_reserveTokens\",\"type\":\"address[]\"}],\"name\":\"removeLiquidityReturn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_reserveToken\",\"type\":\"address\"}],\"name\":\"reserveBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveTokenCount\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveTokens\",\"outputs\":[{\"internalType\":\"contractIERC20Token[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_reserveToken\",\"type\":\"address\"}],\"name\":\"reserveWeight\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restoreRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_onlyOwnerCanUpdateRegistry\",\"type\":\"bool\"}],\"name\":\"restrictRegistryUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_conversionFee\",\"type\":\"uint32\"}],\"name\":\"setConversionFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_sourceToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Token\",\"name\":\"_targetToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"targetAmountAndFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIConverterAnchor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferAnchorOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferTokenOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ConverterTypeThree is an auto generated Go binding around an Ethereum contract.
type ConverterTypeThree struct {
	ConverterTypeThreeCaller     // Read-only binding to the contract
	ConverterTypeThreeTransactor // Write-only binding to the contract
	ConverterTypeThreeFilterer   // Log filterer for contract events
}

// ConverterTypeThreeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConverterTypeThreeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConverterTypeThreeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConverterTypeThreeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConverterTypeThreeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConverterTypeThreeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConverterTypeThreeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConverterTypeThreeSession struct {
	Contract     *ConverterTypeThree // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ConverterTypeThreeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConverterTypeThreeCallerSession struct {
	Contract *ConverterTypeThreeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ConverterTypeThreeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConverterTypeThreeTransactorSession struct {
	Contract     *ConverterTypeThreeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ConverterTypeThreeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConverterTypeThreeRaw struct {
	Contract *ConverterTypeThree // Generic contract binding to access the raw methods on
}

// ConverterTypeThreeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConverterTypeThreeCallerRaw struct {
	Contract *ConverterTypeThreeCaller // Generic read-only contract binding to access the raw methods on
}

// ConverterTypeThreeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConverterTypeThreeTransactorRaw struct {
	Contract *ConverterTypeThreeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConverterTypeThree creates a new instance of ConverterTypeThree, bound to a specific deployed contract.
func NewConverterTypeThree(address common.Address, backend bind.ContractBackend) (*ConverterTypeThree, error) {
	contract, err := bindConverterTypeThree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeThree{ConverterTypeThreeCaller: ConverterTypeThreeCaller{contract: contract}, ConverterTypeThreeTransactor: ConverterTypeThreeTransactor{contract: contract}, ConverterTypeThreeFilterer: ConverterTypeThreeFilterer{contract: contract}}, nil
}

// NewConverterTypeThreeCaller creates a new read-only instance of ConverterTypeThree, bound to a specific deployed contract.
func NewConverterTypeThreeCaller(address common.Address, caller bind.ContractCaller) (*ConverterTypeThreeCaller, error) {
	contract, err := bindConverterTypeThree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeThreeCaller{contract: contract}, nil
}

// NewConverterTypeThreeTransactor creates a new write-only instance of ConverterTypeThree, bound to a specific deployed contract.
func NewConverterTypeThreeTransactor(address common.Address, transactor bind.ContractTransactor) (*ConverterTypeThreeTransactor, error) {
	contract, err := bindConverterTypeThree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeThreeTransactor{contract: contract}, nil
}

// NewConverterTypeThreeFilterer creates a new log filterer instance of ConverterTypeThree, bound to a specific deployed contract.
func NewConverterTypeThreeFilterer(address common.Address, filterer bind.ContractFilterer) (*ConverterTypeThreeFilterer, error) {
	contract, err := bindConverterTypeThree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeThreeFilterer{contract: contract}, nil
}

// bindConverterTypeThree binds a generic wrapper to an already deployed contract.
func bindConverterTypeThree(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConverterTypeThreeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConverterTypeThree *ConverterTypeThreeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConverterTypeThree.Contract.ConverterTypeThreeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConverterTypeThree *ConverterTypeThreeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.ConverterTypeThreeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConverterTypeThree *ConverterTypeThreeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.ConverterTypeThreeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConverterTypeThree *ConverterTypeThreeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConverterTypeThree.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConverterTypeThree *ConverterTypeThreeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConverterTypeThree *ConverterTypeThreeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.contract.Transact(opts, method, params...)
}

// AddLiquidityCost is a free data retrieval call binding the contract method 0x80d9416d.
//
// Solidity: function addLiquidityCost(address[] _reserveTokens, uint256 _reserveTokenIndex, uint256 _reserveAmount) view returns(uint256[])
func (_ConverterTypeThree *ConverterTypeThreeCaller) AddLiquidityCost(opts *bind.CallOpts, _reserveTokens []common.Address, _reserveTokenIndex *big.Int, _reserveAmount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "addLiquidityCost", _reserveTokens, _reserveTokenIndex, _reserveAmount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AddLiquidityCost is a free data retrieval call binding the contract method 0x80d9416d.
//
// Solidity: function addLiquidityCost(address[] _reserveTokens, uint256 _reserveTokenIndex, uint256 _reserveAmount) view returns(uint256[])
func (_ConverterTypeThree *ConverterTypeThreeSession) AddLiquidityCost(_reserveTokens []common.Address, _reserveTokenIndex *big.Int, _reserveAmount *big.Int) ([]*big.Int, error) {
	return _ConverterTypeThree.Contract.AddLiquidityCost(&_ConverterTypeThree.CallOpts, _reserveTokens, _reserveTokenIndex, _reserveAmount)
}

// AddLiquidityCost is a free data retrieval call binding the contract method 0x80d9416d.
//
// Solidity: function addLiquidityCost(address[] _reserveTokens, uint256 _reserveTokenIndex, uint256 _reserveAmount) view returns(uint256[])
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) AddLiquidityCost(_reserveTokens []common.Address, _reserveTokenIndex *big.Int, _reserveAmount *big.Int) ([]*big.Int, error) {
	return _ConverterTypeThree.Contract.AddLiquidityCost(&_ConverterTypeThree.CallOpts, _reserveTokens, _reserveTokenIndex, _reserveAmount)
}

// AddLiquidityReturn is a free data retrieval call binding the contract method 0x4e40c260.
//
// Solidity: function addLiquidityReturn(address _reserveToken, uint256 _reserveAmount) view returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeCaller) AddLiquidityReturn(opts *bind.CallOpts, _reserveToken common.Address, _reserveAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "addLiquidityReturn", _reserveToken, _reserveAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddLiquidityReturn is a free data retrieval call binding the contract method 0x4e40c260.
//
// Solidity: function addLiquidityReturn(address _reserveToken, uint256 _reserveAmount) view returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeSession) AddLiquidityReturn(_reserveToken common.Address, _reserveAmount *big.Int) (*big.Int, error) {
	return _ConverterTypeThree.Contract.AddLiquidityReturn(&_ConverterTypeThree.CallOpts, _reserveToken, _reserveAmount)
}

// AddLiquidityReturn is a free data retrieval call binding the contract method 0x4e40c260.
//
// Solidity: function addLiquidityReturn(address _reserveToken, uint256 _reserveAmount) view returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) AddLiquidityReturn(_reserveToken common.Address, _reserveAmount *big.Int) (*big.Int, error) {
	return _ConverterTypeThree.Contract.AddLiquidityReturn(&_ConverterTypeThree.CallOpts, _reserveToken, _reserveAmount)
}

// Anchor is a free data retrieval call binding the contract method 0xd3fb73b4.
//
// Solidity: function anchor() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeCaller) Anchor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "anchor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Anchor is a free data retrieval call binding the contract method 0xd3fb73b4.
//
// Solidity: function anchor() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeSession) Anchor() (common.Address, error) {
	return _ConverterTypeThree.Contract.Anchor(&_ConverterTypeThree.CallOpts)
}

// Anchor is a free data retrieval call binding the contract method 0xd3fb73b4.
//
// Solidity: function anchor() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) Anchor() (common.Address, error) {
	return _ConverterTypeThree.Contract.Anchor(&_ConverterTypeThree.CallOpts)
}

// AverageRateInfo is a free data retrieval call binding the contract method 0xf0413a1f.
//
// Solidity: function averageRateInfo() view returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeCaller) AverageRateInfo(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "averageRateInfo")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AverageRateInfo is a free data retrieval call binding the contract method 0xf0413a1f.
//
// Solidity: function averageRateInfo() view returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeSession) AverageRateInfo() (*big.Int, error) {
	return _ConverterTypeThree.Contract.AverageRateInfo(&_ConverterTypeThree.CallOpts)
}

// AverageRateInfo is a free data retrieval call binding the contract method 0xf0413a1f.
//
// Solidity: function averageRateInfo() view returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) AverageRateInfo() (*big.Int, error) {
	return _ConverterTypeThree.Contract.AverageRateInfo(&_ConverterTypeThree.CallOpts)
}

// ConnectorTokenCount is a free data retrieval call binding the contract method 0x71f52bf3.
//
// Solidity: function connectorTokenCount() view returns(uint16)
func (_ConverterTypeThree *ConverterTypeThreeCaller) ConnectorTokenCount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "connectorTokenCount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ConnectorTokenCount is a free data retrieval call binding the contract method 0x71f52bf3.
//
// Solidity: function connectorTokenCount() view returns(uint16)
func (_ConverterTypeThree *ConverterTypeThreeSession) ConnectorTokenCount() (uint16, error) {
	return _ConverterTypeThree.Contract.ConnectorTokenCount(&_ConverterTypeThree.CallOpts)
}

// ConnectorTokenCount is a free data retrieval call binding the contract method 0x71f52bf3.
//
// Solidity: function connectorTokenCount() view returns(uint16)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) ConnectorTokenCount() (uint16, error) {
	return _ConverterTypeThree.Contract.ConnectorTokenCount(&_ConverterTypeThree.CallOpts)
}

// ConnectorTokens is a free data retrieval call binding the contract method 0x19b64015.
//
// Solidity: function connectorTokens(uint256 _index) view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeCaller) ConnectorTokens(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "connectorTokens", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorTokens is a free data retrieval call binding the contract method 0x19b64015.
//
// Solidity: function connectorTokens(uint256 _index) view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeSession) ConnectorTokens(_index *big.Int) (common.Address, error) {
	return _ConverterTypeThree.Contract.ConnectorTokens(&_ConverterTypeThree.CallOpts, _index)
}

// ConnectorTokens is a free data retrieval call binding the contract method 0x19b64015.
//
// Solidity: function connectorTokens(uint256 _index) view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) ConnectorTokens(_index *big.Int) (common.Address, error) {
	return _ConverterTypeThree.Contract.ConnectorTokens(&_ConverterTypeThree.CallOpts, _index)
}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address _address) view returns(uint256, uint32, bool, bool, bool)
func (_ConverterTypeThree *ConverterTypeThreeCaller) Connectors(opts *bind.CallOpts, _address common.Address) (*big.Int, uint32, bool, bool, bool, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "connectors", _address)

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
func (_ConverterTypeThree *ConverterTypeThreeSession) Connectors(_address common.Address) (*big.Int, uint32, bool, bool, bool, error) {
	return _ConverterTypeThree.Contract.Connectors(&_ConverterTypeThree.CallOpts, _address)
}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address _address) view returns(uint256, uint32, bool, bool, bool)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) Connectors(_address common.Address) (*big.Int, uint32, bool, bool, bool, error) {
	return _ConverterTypeThree.Contract.Connectors(&_ConverterTypeThree.CallOpts, _address)
}

// ConversionFee is a free data retrieval call binding the contract method 0x579cd3ca.
//
// Solidity: function conversionFee() view returns(uint32)
func (_ConverterTypeThree *ConverterTypeThreeCaller) ConversionFee(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "conversionFee")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ConversionFee is a free data retrieval call binding the contract method 0x579cd3ca.
//
// Solidity: function conversionFee() view returns(uint32)
func (_ConverterTypeThree *ConverterTypeThreeSession) ConversionFee() (uint32, error) {
	return _ConverterTypeThree.Contract.ConversionFee(&_ConverterTypeThree.CallOpts)
}

// ConversionFee is a free data retrieval call binding the contract method 0x579cd3ca.
//
// Solidity: function conversionFee() view returns(uint32)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) ConversionFee() (uint32, error) {
	return _ConverterTypeThree.Contract.ConversionFee(&_ConverterTypeThree.CallOpts)
}

// ConverterType is a free data retrieval call binding the contract method 0x3e8ff43f.
//
// Solidity: function converterType() pure returns(uint16)
func (_ConverterTypeThree *ConverterTypeThreeCaller) ConverterType(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "converterType")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ConverterType is a free data retrieval call binding the contract method 0x3e8ff43f.
//
// Solidity: function converterType() pure returns(uint16)
func (_ConverterTypeThree *ConverterTypeThreeSession) ConverterType() (uint16, error) {
	return _ConverterTypeThree.Contract.ConverterType(&_ConverterTypeThree.CallOpts)
}

// ConverterType is a free data retrieval call binding the contract method 0x3e8ff43f.
//
// Solidity: function converterType() pure returns(uint16)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) ConverterType() (uint16, error) {
	return _ConverterTypeThree.Contract.ConverterType(&_ConverterTypeThree.CallOpts)
}

// GetConnectorBalance is a free data retrieval call binding the contract method 0xd8959512.
//
// Solidity: function getConnectorBalance(address _connectorToken) view returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeCaller) GetConnectorBalance(opts *bind.CallOpts, _connectorToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "getConnectorBalance", _connectorToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConnectorBalance is a free data retrieval call binding the contract method 0xd8959512.
//
// Solidity: function getConnectorBalance(address _connectorToken) view returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeSession) GetConnectorBalance(_connectorToken common.Address) (*big.Int, error) {
	return _ConverterTypeThree.Contract.GetConnectorBalance(&_ConverterTypeThree.CallOpts, _connectorToken)
}

// GetConnectorBalance is a free data retrieval call binding the contract method 0xd8959512.
//
// Solidity: function getConnectorBalance(address _connectorToken) view returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) GetConnectorBalance(_connectorToken common.Address) (*big.Int, error) {
	return _ConverterTypeThree.Contract.GetConnectorBalance(&_ConverterTypeThree.CallOpts, _connectorToken)
}

// GetReturn is a free data retrieval call binding the contract method 0x1e1401f8.
//
// Solidity: function getReturn(address _sourceToken, address _targetToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeThree *ConverterTypeThreeCaller) GetReturn(opts *bind.CallOpts, _sourceToken common.Address, _targetToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "getReturn", _sourceToken, _targetToken, _amount)

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
func (_ConverterTypeThree *ConverterTypeThreeSession) GetReturn(_sourceToken common.Address, _targetToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeThree.Contract.GetReturn(&_ConverterTypeThree.CallOpts, _sourceToken, _targetToken, _amount)
}

// GetReturn is a free data retrieval call binding the contract method 0x1e1401f8.
//
// Solidity: function getReturn(address _sourceToken, address _targetToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) GetReturn(_sourceToken common.Address, _targetToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeThree.Contract.GetReturn(&_ConverterTypeThree.CallOpts, _sourceToken, _targetToken, _amount)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_ConverterTypeThree *ConverterTypeThreeCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_ConverterTypeThree *ConverterTypeThreeSession) IsActive() (bool, error) {
	return _ConverterTypeThree.Contract.IsActive(&_ConverterTypeThree.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) IsActive() (bool, error) {
	return _ConverterTypeThree.Contract.IsActive(&_ConverterTypeThree.CallOpts)
}

// IsV28OrHigher is a free data retrieval call binding the contract method 0xd260529c.
//
// Solidity: function isV28OrHigher() pure returns(bool)
func (_ConverterTypeThree *ConverterTypeThreeCaller) IsV28OrHigher(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "isV28OrHigher")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsV28OrHigher is a free data retrieval call binding the contract method 0xd260529c.
//
// Solidity: function isV28OrHigher() pure returns(bool)
func (_ConverterTypeThree *ConverterTypeThreeSession) IsV28OrHigher() (bool, error) {
	return _ConverterTypeThree.Contract.IsV28OrHigher(&_ConverterTypeThree.CallOpts)
}

// IsV28OrHigher is a free data retrieval call binding the contract method 0xd260529c.
//
// Solidity: function isV28OrHigher() pure returns(bool)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) IsV28OrHigher() (bool, error) {
	return _ConverterTypeThree.Contract.IsV28OrHigher(&_ConverterTypeThree.CallOpts)
}

// MaxConversionFee is a free data retrieval call binding the contract method 0x94c275ad.
//
// Solidity: function maxConversionFee() view returns(uint32)
func (_ConverterTypeThree *ConverterTypeThreeCaller) MaxConversionFee(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "maxConversionFee")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxConversionFee is a free data retrieval call binding the contract method 0x94c275ad.
//
// Solidity: function maxConversionFee() view returns(uint32)
func (_ConverterTypeThree *ConverterTypeThreeSession) MaxConversionFee() (uint32, error) {
	return _ConverterTypeThree.Contract.MaxConversionFee(&_ConverterTypeThree.CallOpts)
}

// MaxConversionFee is a free data retrieval call binding the contract method 0x94c275ad.
//
// Solidity: function maxConversionFee() view returns(uint32)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) MaxConversionFee() (uint32, error) {
	return _ConverterTypeThree.Contract.MaxConversionFee(&_ConverterTypeThree.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "newOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeSession) NewOwner() (common.Address, error) {
	return _ConverterTypeThree.Contract.NewOwner(&_ConverterTypeThree.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) NewOwner() (common.Address, error) {
	return _ConverterTypeThree.Contract.NewOwner(&_ConverterTypeThree.CallOpts)
}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_ConverterTypeThree *ConverterTypeThreeCaller) OnlyOwnerCanUpdateRegistry(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "onlyOwnerCanUpdateRegistry")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_ConverterTypeThree *ConverterTypeThreeSession) OnlyOwnerCanUpdateRegistry() (bool, error) {
	return _ConverterTypeThree.Contract.OnlyOwnerCanUpdateRegistry(&_ConverterTypeThree.CallOpts)
}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) OnlyOwnerCanUpdateRegistry() (bool, error) {
	return _ConverterTypeThree.Contract.OnlyOwnerCanUpdateRegistry(&_ConverterTypeThree.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeSession) Owner() (common.Address, error) {
	return _ConverterTypeThree.Contract.Owner(&_ConverterTypeThree.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) Owner() (common.Address, error) {
	return _ConverterTypeThree.Contract.Owner(&_ConverterTypeThree.CallOpts)
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeCaller) PrevRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "prevRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeSession) PrevRegistry() (common.Address, error) {
	return _ConverterTypeThree.Contract.PrevRegistry(&_ConverterTypeThree.CallOpts)
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) PrevRegistry() (common.Address, error) {
	return _ConverterTypeThree.Contract.PrevRegistry(&_ConverterTypeThree.CallOpts)
}

// RecentAverageRate is a free data retrieval call binding the contract method 0x1f0181bc.
//
// Solidity: function recentAverageRate(address _token) view returns(uint256, uint256)
func (_ConverterTypeThree *ConverterTypeThreeCaller) RecentAverageRate(opts *bind.CallOpts, _token common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "recentAverageRate", _token)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RecentAverageRate is a free data retrieval call binding the contract method 0x1f0181bc.
//
// Solidity: function recentAverageRate(address _token) view returns(uint256, uint256)
func (_ConverterTypeThree *ConverterTypeThreeSession) RecentAverageRate(_token common.Address) (*big.Int, *big.Int, error) {
	return _ConverterTypeThree.Contract.RecentAverageRate(&_ConverterTypeThree.CallOpts, _token)
}

// RecentAverageRate is a free data retrieval call binding the contract method 0x1f0181bc.
//
// Solidity: function recentAverageRate(address _token) view returns(uint256, uint256)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) RecentAverageRate(_token common.Address) (*big.Int, *big.Int, error) {
	return _ConverterTypeThree.Contract.RecentAverageRate(&_ConverterTypeThree.CallOpts, _token)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeSession) Registry() (common.Address, error) {
	return _ConverterTypeThree.Contract.Registry(&_ConverterTypeThree.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) Registry() (common.Address, error) {
	return _ConverterTypeThree.Contract.Registry(&_ConverterTypeThree.CallOpts)
}

// RemoveLiquidityReturn is a free data retrieval call binding the contract method 0x15458837.
//
// Solidity: function removeLiquidityReturn(uint256 _amount, address[] _reserveTokens) view returns(uint256[])
func (_ConverterTypeThree *ConverterTypeThreeCaller) RemoveLiquidityReturn(opts *bind.CallOpts, _amount *big.Int, _reserveTokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "removeLiquidityReturn", _amount, _reserveTokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RemoveLiquidityReturn is a free data retrieval call binding the contract method 0x15458837.
//
// Solidity: function removeLiquidityReturn(uint256 _amount, address[] _reserveTokens) view returns(uint256[])
func (_ConverterTypeThree *ConverterTypeThreeSession) RemoveLiquidityReturn(_amount *big.Int, _reserveTokens []common.Address) ([]*big.Int, error) {
	return _ConverterTypeThree.Contract.RemoveLiquidityReturn(&_ConverterTypeThree.CallOpts, _amount, _reserveTokens)
}

// RemoveLiquidityReturn is a free data retrieval call binding the contract method 0x15458837.
//
// Solidity: function removeLiquidityReturn(uint256 _amount, address[] _reserveTokens) view returns(uint256[])
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) RemoveLiquidityReturn(_amount *big.Int, _reserveTokens []common.Address) ([]*big.Int, error) {
	return _ConverterTypeThree.Contract.RemoveLiquidityReturn(&_ConverterTypeThree.CallOpts, _amount, _reserveTokens)
}

// ReserveBalance is a free data retrieval call binding the contract method 0xdc8de379.
//
// Solidity: function reserveBalance(address _reserveToken) view returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeCaller) ReserveBalance(opts *bind.CallOpts, _reserveToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "reserveBalance", _reserveToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveBalance is a free data retrieval call binding the contract method 0xdc8de379.
//
// Solidity: function reserveBalance(address _reserveToken) view returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeSession) ReserveBalance(_reserveToken common.Address) (*big.Int, error) {
	return _ConverterTypeThree.Contract.ReserveBalance(&_ConverterTypeThree.CallOpts, _reserveToken)
}

// ReserveBalance is a free data retrieval call binding the contract method 0xdc8de379.
//
// Solidity: function reserveBalance(address _reserveToken) view returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) ReserveBalance(_reserveToken common.Address) (*big.Int, error) {
	return _ConverterTypeThree.Contract.ReserveBalance(&_ConverterTypeThree.CallOpts, _reserveToken)
}

// ReserveBalances is a free data retrieval call binding the contract method 0x613e53a7.
//
// Solidity: function reserveBalances() view returns(uint256, uint256)
func (_ConverterTypeThree *ConverterTypeThreeCaller) ReserveBalances(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "reserveBalances")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ReserveBalances is a free data retrieval call binding the contract method 0x613e53a7.
//
// Solidity: function reserveBalances() view returns(uint256, uint256)
func (_ConverterTypeThree *ConverterTypeThreeSession) ReserveBalances() (*big.Int, *big.Int, error) {
	return _ConverterTypeThree.Contract.ReserveBalances(&_ConverterTypeThree.CallOpts)
}

// ReserveBalances is a free data retrieval call binding the contract method 0x613e53a7.
//
// Solidity: function reserveBalances() view returns(uint256, uint256)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) ReserveBalances() (*big.Int, *big.Int, error) {
	return _ConverterTypeThree.Contract.ReserveBalances(&_ConverterTypeThree.CallOpts)
}

// ReserveTokenCount is a free data retrieval call binding the contract method 0x9b99a8e2.
//
// Solidity: function reserveTokenCount() view returns(uint16)
func (_ConverterTypeThree *ConverterTypeThreeCaller) ReserveTokenCount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "reserveTokenCount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ReserveTokenCount is a free data retrieval call binding the contract method 0x9b99a8e2.
//
// Solidity: function reserveTokenCount() view returns(uint16)
func (_ConverterTypeThree *ConverterTypeThreeSession) ReserveTokenCount() (uint16, error) {
	return _ConverterTypeThree.Contract.ReserveTokenCount(&_ConverterTypeThree.CallOpts)
}

// ReserveTokenCount is a free data retrieval call binding the contract method 0x9b99a8e2.
//
// Solidity: function reserveTokenCount() view returns(uint16)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) ReserveTokenCount() (uint16, error) {
	return _ConverterTypeThree.Contract.ReserveTokenCount(&_ConverterTypeThree.CallOpts)
}

// ReserveTokens is a free data retrieval call binding the contract method 0x27ac36c4.
//
// Solidity: function reserveTokens() view returns(address[])
func (_ConverterTypeThree *ConverterTypeThreeCaller) ReserveTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "reserveTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ReserveTokens is a free data retrieval call binding the contract method 0x27ac36c4.
//
// Solidity: function reserveTokens() view returns(address[])
func (_ConverterTypeThree *ConverterTypeThreeSession) ReserveTokens() ([]common.Address, error) {
	return _ConverterTypeThree.Contract.ReserveTokens(&_ConverterTypeThree.CallOpts)
}

// ReserveTokens is a free data retrieval call binding the contract method 0x27ac36c4.
//
// Solidity: function reserveTokens() view returns(address[])
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) ReserveTokens() ([]common.Address, error) {
	return _ConverterTypeThree.Contract.ReserveTokens(&_ConverterTypeThree.CallOpts)
}

// ReserveWeight is a free data retrieval call binding the contract method 0x1cfab290.
//
// Solidity: function reserveWeight(address _reserveToken) view returns(uint32)
func (_ConverterTypeThree *ConverterTypeThreeCaller) ReserveWeight(opts *bind.CallOpts, _reserveToken common.Address) (uint32, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "reserveWeight", _reserveToken)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ReserveWeight is a free data retrieval call binding the contract method 0x1cfab290.
//
// Solidity: function reserveWeight(address _reserveToken) view returns(uint32)
func (_ConverterTypeThree *ConverterTypeThreeSession) ReserveWeight(_reserveToken common.Address) (uint32, error) {
	return _ConverterTypeThree.Contract.ReserveWeight(&_ConverterTypeThree.CallOpts, _reserveToken)
}

// ReserveWeight is a free data retrieval call binding the contract method 0x1cfab290.
//
// Solidity: function reserveWeight(address _reserveToken) view returns(uint32)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) ReserveWeight(_reserveToken common.Address) (uint32, error) {
	return _ConverterTypeThree.Contract.ReserveWeight(&_ConverterTypeThree.CallOpts, _reserveToken)
}

// TargetAmountAndFee is a free data retrieval call binding the contract method 0xaf94b8d8.
//
// Solidity: function targetAmountAndFee(address _sourceToken, address _targetToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeThree *ConverterTypeThreeCaller) TargetAmountAndFee(opts *bind.CallOpts, _sourceToken common.Address, _targetToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "targetAmountAndFee", _sourceToken, _targetToken, _amount)

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
func (_ConverterTypeThree *ConverterTypeThreeSession) TargetAmountAndFee(_sourceToken common.Address, _targetToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeThree.Contract.TargetAmountAndFee(&_ConverterTypeThree.CallOpts, _sourceToken, _targetToken, _amount)
}

// TargetAmountAndFee is a free data retrieval call binding the contract method 0xaf94b8d8.
//
// Solidity: function targetAmountAndFee(address _sourceToken, address _targetToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) TargetAmountAndFee(_sourceToken common.Address, _targetToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeThree.Contract.TargetAmountAndFee(&_ConverterTypeThree.CallOpts, _sourceToken, _targetToken, _amount)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeSession) Token() (common.Address, error) {
	return _ConverterTypeThree.Contract.Token(&_ConverterTypeThree.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) Token() (common.Address, error) {
	return _ConverterTypeThree.Contract.Token(&_ConverterTypeThree.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_ConverterTypeThree *ConverterTypeThreeCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConverterTypeThree.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_ConverterTypeThree *ConverterTypeThreeSession) Version() (uint16, error) {
	return _ConverterTypeThree.Contract.Version(&_ConverterTypeThree.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_ConverterTypeThree *ConverterTypeThreeCallerSession) Version() (uint16, error) {
	return _ConverterTypeThree.Contract.Version(&_ConverterTypeThree.CallOpts)
}

// AcceptAnchorOwnership is a paid mutator transaction binding the contract method 0xcdc91c69.
//
// Solidity: function acceptAnchorOwnership() returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactor) AcceptAnchorOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "acceptAnchorOwnership")
}

// AcceptAnchorOwnership is a paid mutator transaction binding the contract method 0xcdc91c69.
//
// Solidity: function acceptAnchorOwnership() returns()
func (_ConverterTypeThree *ConverterTypeThreeSession) AcceptAnchorOwnership() (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.AcceptAnchorOwnership(&_ConverterTypeThree.TransactOpts)
}

// AcceptAnchorOwnership is a paid mutator transaction binding the contract method 0xcdc91c69.
//
// Solidity: function acceptAnchorOwnership() returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) AcceptAnchorOwnership() (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.AcceptAnchorOwnership(&_ConverterTypeThree.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ConverterTypeThree *ConverterTypeThreeSession) AcceptOwnership() (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.AcceptOwnership(&_ConverterTypeThree.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.AcceptOwnership(&_ConverterTypeThree.TransactOpts)
}

// AcceptTokenOwnership is a paid mutator transaction binding the contract method 0x38a5e016.
//
// Solidity: function acceptTokenOwnership() returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactor) AcceptTokenOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "acceptTokenOwnership")
}

// AcceptTokenOwnership is a paid mutator transaction binding the contract method 0x38a5e016.
//
// Solidity: function acceptTokenOwnership() returns()
func (_ConverterTypeThree *ConverterTypeThreeSession) AcceptTokenOwnership() (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.AcceptTokenOwnership(&_ConverterTypeThree.TransactOpts)
}

// AcceptTokenOwnership is a paid mutator transaction binding the contract method 0x38a5e016.
//
// Solidity: function acceptTokenOwnership() returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) AcceptTokenOwnership() (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.AcceptTokenOwnership(&_ConverterTypeThree.TransactOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x422f1043.
//
// Solidity: function addLiquidity(uint256 _reserve1Amount, uint256 _reserve2Amount, uint256 _minReturn) payable returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeTransactor) AddLiquidity(opts *bind.TransactOpts, _reserve1Amount *big.Int, _reserve2Amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "addLiquidity", _reserve1Amount, _reserve2Amount, _minReturn)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x422f1043.
//
// Solidity: function addLiquidity(uint256 _reserve1Amount, uint256 _reserve2Amount, uint256 _minReturn) payable returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeSession) AddLiquidity(_reserve1Amount *big.Int, _reserve2Amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.AddLiquidity(&_ConverterTypeThree.TransactOpts, _reserve1Amount, _reserve2Amount, _minReturn)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x422f1043.
//
// Solidity: function addLiquidity(uint256 _reserve1Amount, uint256 _reserve2Amount, uint256 _minReturn) payable returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) AddLiquidity(_reserve1Amount *big.Int, _reserve2Amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.AddLiquidity(&_ConverterTypeThree.TransactOpts, _reserve1Amount, _reserve2Amount, _minReturn)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x7d8916bd.
//
// Solidity: function addLiquidity(address[] _reserveTokens, uint256[] _reserveAmounts, uint256 _minReturn) payable returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeTransactor) AddLiquidity0(opts *bind.TransactOpts, _reserveTokens []common.Address, _reserveAmounts []*big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "addLiquidity0", _reserveTokens, _reserveAmounts, _minReturn)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x7d8916bd.
//
// Solidity: function addLiquidity(address[] _reserveTokens, uint256[] _reserveAmounts, uint256 _minReturn) payable returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeSession) AddLiquidity0(_reserveTokens []common.Address, _reserveAmounts []*big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.AddLiquidity0(&_ConverterTypeThree.TransactOpts, _reserveTokens, _reserveAmounts, _minReturn)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x7d8916bd.
//
// Solidity: function addLiquidity(address[] _reserveTokens, uint256[] _reserveAmounts, uint256 _minReturn) payable returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) AddLiquidity0(_reserveTokens []common.Address, _reserveAmounts []*big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.AddLiquidity0(&_ConverterTypeThree.TransactOpts, _reserveTokens, _reserveAmounts, _minReturn)
}

// AddReserve is a paid mutator transaction binding the contract method 0x6a49d2c4.
//
// Solidity: function addReserve(address _token, uint32 _weight) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactor) AddReserve(opts *bind.TransactOpts, _token common.Address, _weight uint32) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "addReserve", _token, _weight)
}

// AddReserve is a paid mutator transaction binding the contract method 0x6a49d2c4.
//
// Solidity: function addReserve(address _token, uint32 _weight) returns()
func (_ConverterTypeThree *ConverterTypeThreeSession) AddReserve(_token common.Address, _weight uint32) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.AddReserve(&_ConverterTypeThree.TransactOpts, _token, _weight)
}

// AddReserve is a paid mutator transaction binding the contract method 0x6a49d2c4.
//
// Solidity: function addReserve(address _token, uint32 _weight) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) AddReserve(_token common.Address, _weight uint32) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.AddReserve(&_ConverterTypeThree.TransactOpts, _token, _weight)
}

// Convert is a paid mutator transaction binding the contract method 0xe8dc12ff.
//
// Solidity: function convert(address _sourceToken, address _targetToken, uint256 _amount, address _trader, address _beneficiary) payable returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeTransactor) Convert(opts *bind.TransactOpts, _sourceToken common.Address, _targetToken common.Address, _amount *big.Int, _trader common.Address, _beneficiary common.Address) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "convert", _sourceToken, _targetToken, _amount, _trader, _beneficiary)
}

// Convert is a paid mutator transaction binding the contract method 0xe8dc12ff.
//
// Solidity: function convert(address _sourceToken, address _targetToken, uint256 _amount, address _trader, address _beneficiary) payable returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeSession) Convert(_sourceToken common.Address, _targetToken common.Address, _amount *big.Int, _trader common.Address, _beneficiary common.Address) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.Convert(&_ConverterTypeThree.TransactOpts, _sourceToken, _targetToken, _amount, _trader, _beneficiary)
}

// Convert is a paid mutator transaction binding the contract method 0xe8dc12ff.
//
// Solidity: function convert(address _sourceToken, address _targetToken, uint256 _amount, address _trader, address _beneficiary) payable returns(uint256)
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) Convert(_sourceToken common.Address, _targetToken common.Address, _amount *big.Int, _trader common.Address, _beneficiary common.Address) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.Convert(&_ConverterTypeThree.TransactOpts, _sourceToken, _targetToken, _amount, _trader, _beneficiary)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x857620e1.
//
// Solidity: function removeLiquidity(uint256 _amount, uint256 _reserve1MinReturn, uint256 _reserve2MinReturn) returns(uint256, uint256)
func (_ConverterTypeThree *ConverterTypeThreeTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, _reserve1MinReturn *big.Int, _reserve2MinReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "removeLiquidity", _amount, _reserve1MinReturn, _reserve2MinReturn)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x857620e1.
//
// Solidity: function removeLiquidity(uint256 _amount, uint256 _reserve1MinReturn, uint256 _reserve2MinReturn) returns(uint256, uint256)
func (_ConverterTypeThree *ConverterTypeThreeSession) RemoveLiquidity(_amount *big.Int, _reserve1MinReturn *big.Int, _reserve2MinReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.RemoveLiquidity(&_ConverterTypeThree.TransactOpts, _amount, _reserve1MinReturn, _reserve2MinReturn)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x857620e1.
//
// Solidity: function removeLiquidity(uint256 _amount, uint256 _reserve1MinReturn, uint256 _reserve2MinReturn) returns(uint256, uint256)
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) RemoveLiquidity(_amount *big.Int, _reserve1MinReturn *big.Int, _reserve2MinReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.RemoveLiquidity(&_ConverterTypeThree.TransactOpts, _amount, _reserve1MinReturn, _reserve2MinReturn)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0xb127c0a5.
//
// Solidity: function removeLiquidity(uint256 _amount, address[] _reserveTokens, uint256[] _reserveMinReturnAmounts) returns(uint256[])
func (_ConverterTypeThree *ConverterTypeThreeTransactor) RemoveLiquidity0(opts *bind.TransactOpts, _amount *big.Int, _reserveTokens []common.Address, _reserveMinReturnAmounts []*big.Int) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "removeLiquidity0", _amount, _reserveTokens, _reserveMinReturnAmounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0xb127c0a5.
//
// Solidity: function removeLiquidity(uint256 _amount, address[] _reserveTokens, uint256[] _reserveMinReturnAmounts) returns(uint256[])
func (_ConverterTypeThree *ConverterTypeThreeSession) RemoveLiquidity0(_amount *big.Int, _reserveTokens []common.Address, _reserveMinReturnAmounts []*big.Int) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.RemoveLiquidity0(&_ConverterTypeThree.TransactOpts, _amount, _reserveTokens, _reserveMinReturnAmounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0xb127c0a5.
//
// Solidity: function removeLiquidity(uint256 _amount, address[] _reserveTokens, uint256[] _reserveMinReturnAmounts) returns(uint256[])
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) RemoveLiquidity0(_amount *big.Int, _reserveTokens []common.Address, _reserveMinReturnAmounts []*big.Int) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.RemoveLiquidity0(&_ConverterTypeThree.TransactOpts, _amount, _reserveTokens, _reserveMinReturnAmounts)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactor) RestoreRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "restoreRegistry")
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_ConverterTypeThree *ConverterTypeThreeSession) RestoreRegistry() (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.RestoreRegistry(&_ConverterTypeThree.TransactOpts)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) RestoreRegistry() (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.RestoreRegistry(&_ConverterTypeThree.TransactOpts)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactor) RestrictRegistryUpdate(opts *bind.TransactOpts, _onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "restrictRegistryUpdate", _onlyOwnerCanUpdateRegistry)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_ConverterTypeThree *ConverterTypeThreeSession) RestrictRegistryUpdate(_onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.RestrictRegistryUpdate(&_ConverterTypeThree.TransactOpts, _onlyOwnerCanUpdateRegistry)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) RestrictRegistryUpdate(_onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.RestrictRegistryUpdate(&_ConverterTypeThree.TransactOpts, _onlyOwnerCanUpdateRegistry)
}

// SetConversionFee is a paid mutator transaction binding the contract method 0xecbca55d.
//
// Solidity: function setConversionFee(uint32 _conversionFee) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactor) SetConversionFee(opts *bind.TransactOpts, _conversionFee uint32) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "setConversionFee", _conversionFee)
}

// SetConversionFee is a paid mutator transaction binding the contract method 0xecbca55d.
//
// Solidity: function setConversionFee(uint32 _conversionFee) returns()
func (_ConverterTypeThree *ConverterTypeThreeSession) SetConversionFee(_conversionFee uint32) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.SetConversionFee(&_ConverterTypeThree.TransactOpts, _conversionFee)
}

// SetConversionFee is a paid mutator transaction binding the contract method 0xecbca55d.
//
// Solidity: function setConversionFee(uint32 _conversionFee) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) SetConversionFee(_conversionFee uint32) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.SetConversionFee(&_ConverterTypeThree.TransactOpts, _conversionFee)
}

// TransferAnchorOwnership is a paid mutator transaction binding the contract method 0x67b6d57c.
//
// Solidity: function transferAnchorOwnership(address _newOwner) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactor) TransferAnchorOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "transferAnchorOwnership", _newOwner)
}

// TransferAnchorOwnership is a paid mutator transaction binding the contract method 0x67b6d57c.
//
// Solidity: function transferAnchorOwnership(address _newOwner) returns()
func (_ConverterTypeThree *ConverterTypeThreeSession) TransferAnchorOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.TransferAnchorOwnership(&_ConverterTypeThree.TransactOpts, _newOwner)
}

// TransferAnchorOwnership is a paid mutator transaction binding the contract method 0x67b6d57c.
//
// Solidity: function transferAnchorOwnership(address _newOwner) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) TransferAnchorOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.TransferAnchorOwnership(&_ConverterTypeThree.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ConverterTypeThree *ConverterTypeThreeSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.TransferOwnership(&_ConverterTypeThree.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.TransferOwnership(&_ConverterTypeThree.TransactOpts, _newOwner)
}

// TransferTokenOwnership is a paid mutator transaction binding the contract method 0x21e6b53d.
//
// Solidity: function transferTokenOwnership(address _newOwner) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactor) TransferTokenOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "transferTokenOwnership", _newOwner)
}

// TransferTokenOwnership is a paid mutator transaction binding the contract method 0x21e6b53d.
//
// Solidity: function transferTokenOwnership(address _newOwner) returns()
func (_ConverterTypeThree *ConverterTypeThreeSession) TransferTokenOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.TransferTokenOwnership(&_ConverterTypeThree.TransactOpts, _newOwner)
}

// TransferTokenOwnership is a paid mutator transaction binding the contract method 0x21e6b53d.
//
// Solidity: function transferTokenOwnership(address _newOwner) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) TransferTokenOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.TransferTokenOwnership(&_ConverterTypeThree.TransactOpts, _newOwner)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactor) UpdateRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "updateRegistry")
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_ConverterTypeThree *ConverterTypeThreeSession) UpdateRegistry() (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.UpdateRegistry(&_ConverterTypeThree.TransactOpts)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) UpdateRegistry() (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.UpdateRegistry(&_ConverterTypeThree.TransactOpts)
}

// Upgrade is a paid mutator transaction binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactor) Upgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "upgrade")
}

// Upgrade is a paid mutator transaction binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() returns()
func (_ConverterTypeThree *ConverterTypeThreeSession) Upgrade() (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.Upgrade(&_ConverterTypeThree.TransactOpts)
}

// Upgrade is a paid mutator transaction binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) Upgrade() (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.Upgrade(&_ConverterTypeThree.TransactOpts)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x690d8320.
//
// Solidity: function withdrawETH(address _to) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactor) WithdrawETH(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "withdrawETH", _to)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x690d8320.
//
// Solidity: function withdrawETH(address _to) returns()
func (_ConverterTypeThree *ConverterTypeThreeSession) WithdrawETH(_to common.Address) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.WithdrawETH(&_ConverterTypeThree.TransactOpts, _to)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x690d8320.
//
// Solidity: function withdrawETH(address _to) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) WithdrawETH(_to common.Address) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.WithdrawETH(&_ConverterTypeThree.TransactOpts, _to)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactor) WithdrawTokens(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.Transact(opts, "withdrawTokens", _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_ConverterTypeThree *ConverterTypeThreeSession) WithdrawTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.WithdrawTokens(&_ConverterTypeThree.TransactOpts, _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) WithdrawTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.WithdrawTokens(&_ConverterTypeThree.TransactOpts, _token, _to, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeThree.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ConverterTypeThree *ConverterTypeThreeSession) Receive() (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.Receive(&_ConverterTypeThree.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ConverterTypeThree *ConverterTypeThreeTransactorSession) Receive() (*types.Transaction, error) {
	return _ConverterTypeThree.Contract.Receive(&_ConverterTypeThree.TransactOpts)
}

// ConverterTypeThreeActivationIterator is returned from FilterActivation and is used to iterate over the raw logs and unpacked data for Activation events raised by the ConverterTypeThree contract.
type ConverterTypeThreeActivationIterator struct {
	Event *ConverterTypeThreeActivation // Event containing the contract specifics and raw log

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
func (it *ConverterTypeThreeActivationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeThreeActivation)
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
		it.Event = new(ConverterTypeThreeActivation)
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
func (it *ConverterTypeThreeActivationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeThreeActivationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeThreeActivation represents a Activation event raised by the ConverterTypeThree contract.
type ConverterTypeThreeActivation struct {
	Type      uint16
	Anchor    common.Address
	Activated bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterActivation is a free log retrieval operation binding the contract event 0x6b08c2e2c9969e55a647a764db9b554d64dc42f1a704da11a6d5b129ad163f2c.
//
// Solidity: event Activation(uint16 indexed _type, address indexed _anchor, bool indexed _activated)
func (_ConverterTypeThree *ConverterTypeThreeFilterer) FilterActivation(opts *bind.FilterOpts, _type []uint16, _anchor []common.Address, _activated []bool) (*ConverterTypeThreeActivationIterator, error) {

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

	logs, sub, err := _ConverterTypeThree.contract.FilterLogs(opts, "Activation", _typeRule, _anchorRule, _activatedRule)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeThreeActivationIterator{contract: _ConverterTypeThree.contract, event: "Activation", logs: logs, sub: sub}, nil
}

// WatchActivation is a free log subscription operation binding the contract event 0x6b08c2e2c9969e55a647a764db9b554d64dc42f1a704da11a6d5b129ad163f2c.
//
// Solidity: event Activation(uint16 indexed _type, address indexed _anchor, bool indexed _activated)
func (_ConverterTypeThree *ConverterTypeThreeFilterer) WatchActivation(opts *bind.WatchOpts, sink chan<- *ConverterTypeThreeActivation, _type []uint16, _anchor []common.Address, _activated []bool) (event.Subscription, error) {

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

	logs, sub, err := _ConverterTypeThree.contract.WatchLogs(opts, "Activation", _typeRule, _anchorRule, _activatedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeThreeActivation)
				if err := _ConverterTypeThree.contract.UnpackLog(event, "Activation", log); err != nil {
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
func (_ConverterTypeThree *ConverterTypeThreeFilterer) ParseActivation(log types.Log) (*ConverterTypeThreeActivation, error) {
	event := new(ConverterTypeThreeActivation)
	if err := _ConverterTypeThree.contract.UnpackLog(event, "Activation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterTypeThreeConversionIterator is returned from FilterConversion and is used to iterate over the raw logs and unpacked data for Conversion events raised by the ConverterTypeThree contract.
type ConverterTypeThreeConversionIterator struct {
	Event *ConverterTypeThreeConversion // Event containing the contract specifics and raw log

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
func (it *ConverterTypeThreeConversionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeThreeConversion)
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
		it.Event = new(ConverterTypeThreeConversion)
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
func (it *ConverterTypeThreeConversionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeThreeConversionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeThreeConversion represents a Conversion event raised by the ConverterTypeThree contract.
type ConverterTypeThreeConversion struct {
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
func (_ConverterTypeThree *ConverterTypeThreeFilterer) FilterConversion(opts *bind.FilterOpts, _fromToken []common.Address, _toToken []common.Address, _trader []common.Address) (*ConverterTypeThreeConversionIterator, error) {

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

	logs, sub, err := _ConverterTypeThree.contract.FilterLogs(opts, "Conversion", _fromTokenRule, _toTokenRule, _traderRule)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeThreeConversionIterator{contract: _ConverterTypeThree.contract, event: "Conversion", logs: logs, sub: sub}, nil
}

// WatchConversion is a free log subscription operation binding the contract event 0x276856b36cbc45526a0ba64f44611557a2a8b68662c5388e9fe6d72e86e1c8cb.
//
// Solidity: event Conversion(address indexed _fromToken, address indexed _toToken, address indexed _trader, uint256 _amount, uint256 _return, int256 _conversionFee)
func (_ConverterTypeThree *ConverterTypeThreeFilterer) WatchConversion(opts *bind.WatchOpts, sink chan<- *ConverterTypeThreeConversion, _fromToken []common.Address, _toToken []common.Address, _trader []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ConverterTypeThree.contract.WatchLogs(opts, "Conversion", _fromTokenRule, _toTokenRule, _traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeThreeConversion)
				if err := _ConverterTypeThree.contract.UnpackLog(event, "Conversion", log); err != nil {
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
func (_ConverterTypeThree *ConverterTypeThreeFilterer) ParseConversion(log types.Log) (*ConverterTypeThreeConversion, error) {
	event := new(ConverterTypeThreeConversion)
	if err := _ConverterTypeThree.contract.UnpackLog(event, "Conversion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterTypeThreeConversionFeeUpdateIterator is returned from FilterConversionFeeUpdate and is used to iterate over the raw logs and unpacked data for ConversionFeeUpdate events raised by the ConverterTypeThree contract.
type ConverterTypeThreeConversionFeeUpdateIterator struct {
	Event *ConverterTypeThreeConversionFeeUpdate // Event containing the contract specifics and raw log

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
func (it *ConverterTypeThreeConversionFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeThreeConversionFeeUpdate)
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
		it.Event = new(ConverterTypeThreeConversionFeeUpdate)
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
func (it *ConverterTypeThreeConversionFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeThreeConversionFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeThreeConversionFeeUpdate represents a ConversionFeeUpdate event raised by the ConverterTypeThree contract.
type ConverterTypeThreeConversionFeeUpdate struct {
	PrevFee uint32
	NewFee  uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConversionFeeUpdate is a free log retrieval operation binding the contract event 0x81cd2ffb37dd237c0e4e2a3de5265fcf9deb43d3e7801e80db9f1ccfba7ee600.
//
// Solidity: event ConversionFeeUpdate(uint32 _prevFee, uint32 _newFee)
func (_ConverterTypeThree *ConverterTypeThreeFilterer) FilterConversionFeeUpdate(opts *bind.FilterOpts) (*ConverterTypeThreeConversionFeeUpdateIterator, error) {

	logs, sub, err := _ConverterTypeThree.contract.FilterLogs(opts, "ConversionFeeUpdate")
	if err != nil {
		return nil, err
	}
	return &ConverterTypeThreeConversionFeeUpdateIterator{contract: _ConverterTypeThree.contract, event: "ConversionFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchConversionFeeUpdate is a free log subscription operation binding the contract event 0x81cd2ffb37dd237c0e4e2a3de5265fcf9deb43d3e7801e80db9f1ccfba7ee600.
//
// Solidity: event ConversionFeeUpdate(uint32 _prevFee, uint32 _newFee)
func (_ConverterTypeThree *ConverterTypeThreeFilterer) WatchConversionFeeUpdate(opts *bind.WatchOpts, sink chan<- *ConverterTypeThreeConversionFeeUpdate) (event.Subscription, error) {

	logs, sub, err := _ConverterTypeThree.contract.WatchLogs(opts, "ConversionFeeUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeThreeConversionFeeUpdate)
				if err := _ConverterTypeThree.contract.UnpackLog(event, "ConversionFeeUpdate", log); err != nil {
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
func (_ConverterTypeThree *ConverterTypeThreeFilterer) ParseConversionFeeUpdate(log types.Log) (*ConverterTypeThreeConversionFeeUpdate, error) {
	event := new(ConverterTypeThreeConversionFeeUpdate)
	if err := _ConverterTypeThree.contract.UnpackLog(event, "ConversionFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterTypeThreeLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the ConverterTypeThree contract.
type ConverterTypeThreeLiquidityAddedIterator struct {
	Event *ConverterTypeThreeLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *ConverterTypeThreeLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeThreeLiquidityAdded)
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
		it.Event = new(ConverterTypeThreeLiquidityAdded)
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
func (it *ConverterTypeThreeLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeThreeLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeThreeLiquidityAdded represents a LiquidityAdded event raised by the ConverterTypeThree contract.
type ConverterTypeThreeLiquidityAdded struct {
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
func (_ConverterTypeThree *ConverterTypeThreeFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, _provider []common.Address, _reserveToken []common.Address) (*ConverterTypeThreeLiquidityAddedIterator, error) {

	var _providerRule []interface{}
	for _, _providerItem := range _provider {
		_providerRule = append(_providerRule, _providerItem)
	}
	var _reserveTokenRule []interface{}
	for _, _reserveTokenItem := range _reserveToken {
		_reserveTokenRule = append(_reserveTokenRule, _reserveTokenItem)
	}

	logs, sub, err := _ConverterTypeThree.contract.FilterLogs(opts, "LiquidityAdded", _providerRule, _reserveTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeThreeLiquidityAddedIterator{contract: _ConverterTypeThree.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0x4a1a2a6176e9646d9e3157f7c2ab3c499f18337c0b0828cfb28e0a61de4a11f7.
//
// Solidity: event LiquidityAdded(address indexed _provider, address indexed _reserveToken, uint256 _amount, uint256 _newBalance, uint256 _newSupply)
func (_ConverterTypeThree *ConverterTypeThreeFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *ConverterTypeThreeLiquidityAdded, _provider []common.Address, _reserveToken []common.Address) (event.Subscription, error) {

	var _providerRule []interface{}
	for _, _providerItem := range _provider {
		_providerRule = append(_providerRule, _providerItem)
	}
	var _reserveTokenRule []interface{}
	for _, _reserveTokenItem := range _reserveToken {
		_reserveTokenRule = append(_reserveTokenRule, _reserveTokenItem)
	}

	logs, sub, err := _ConverterTypeThree.contract.WatchLogs(opts, "LiquidityAdded", _providerRule, _reserveTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeThreeLiquidityAdded)
				if err := _ConverterTypeThree.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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
func (_ConverterTypeThree *ConverterTypeThreeFilterer) ParseLiquidityAdded(log types.Log) (*ConverterTypeThreeLiquidityAdded, error) {
	event := new(ConverterTypeThreeLiquidityAdded)
	if err := _ConverterTypeThree.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterTypeThreeLiquidityRemovedIterator is returned from FilterLiquidityRemoved and is used to iterate over the raw logs and unpacked data for LiquidityRemoved events raised by the ConverterTypeThree contract.
type ConverterTypeThreeLiquidityRemovedIterator struct {
	Event *ConverterTypeThreeLiquidityRemoved // Event containing the contract specifics and raw log

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
func (it *ConverterTypeThreeLiquidityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeThreeLiquidityRemoved)
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
		it.Event = new(ConverterTypeThreeLiquidityRemoved)
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
func (it *ConverterTypeThreeLiquidityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeThreeLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeThreeLiquidityRemoved represents a LiquidityRemoved event raised by the ConverterTypeThree contract.
type ConverterTypeThreeLiquidityRemoved struct {
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
func (_ConverterTypeThree *ConverterTypeThreeFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, _provider []common.Address, _reserveToken []common.Address) (*ConverterTypeThreeLiquidityRemovedIterator, error) {

	var _providerRule []interface{}
	for _, _providerItem := range _provider {
		_providerRule = append(_providerRule, _providerItem)
	}
	var _reserveTokenRule []interface{}
	for _, _reserveTokenItem := range _reserveToken {
		_reserveTokenRule = append(_reserveTokenRule, _reserveTokenItem)
	}

	logs, sub, err := _ConverterTypeThree.contract.FilterLogs(opts, "LiquidityRemoved", _providerRule, _reserveTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeThreeLiquidityRemovedIterator{contract: _ConverterTypeThree.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

// WatchLiquidityRemoved is a free log subscription operation binding the contract event 0xbc7d19d505c7ec4db83f3b51f19fb98c4c8a99922e7839d1ee608dfbee29501b.
//
// Solidity: event LiquidityRemoved(address indexed _provider, address indexed _reserveToken, uint256 _amount, uint256 _newBalance, uint256 _newSupply)
func (_ConverterTypeThree *ConverterTypeThreeFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *ConverterTypeThreeLiquidityRemoved, _provider []common.Address, _reserveToken []common.Address) (event.Subscription, error) {

	var _providerRule []interface{}
	for _, _providerItem := range _provider {
		_providerRule = append(_providerRule, _providerItem)
	}
	var _reserveTokenRule []interface{}
	for _, _reserveTokenItem := range _reserveToken {
		_reserveTokenRule = append(_reserveTokenRule, _reserveTokenItem)
	}

	logs, sub, err := _ConverterTypeThree.contract.WatchLogs(opts, "LiquidityRemoved", _providerRule, _reserveTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeThreeLiquidityRemoved)
				if err := _ConverterTypeThree.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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
func (_ConverterTypeThree *ConverterTypeThreeFilterer) ParseLiquidityRemoved(log types.Log) (*ConverterTypeThreeLiquidityRemoved, error) {
	event := new(ConverterTypeThreeLiquidityRemoved)
	if err := _ConverterTypeThree.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterTypeThreeOwnerUpdateIterator is returned from FilterOwnerUpdate and is used to iterate over the raw logs and unpacked data for OwnerUpdate events raised by the ConverterTypeThree contract.
type ConverterTypeThreeOwnerUpdateIterator struct {
	Event *ConverterTypeThreeOwnerUpdate // Event containing the contract specifics and raw log

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
func (it *ConverterTypeThreeOwnerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeThreeOwnerUpdate)
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
		it.Event = new(ConverterTypeThreeOwnerUpdate)
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
func (it *ConverterTypeThreeOwnerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeThreeOwnerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeThreeOwnerUpdate represents a OwnerUpdate event raised by the ConverterTypeThree contract.
type ConverterTypeThreeOwnerUpdate struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdate is a free log retrieval operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_ConverterTypeThree *ConverterTypeThreeFilterer) FilterOwnerUpdate(opts *bind.FilterOpts, _prevOwner []common.Address, _newOwner []common.Address) (*ConverterTypeThreeOwnerUpdateIterator, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ConverterTypeThree.contract.FilterLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeThreeOwnerUpdateIterator{contract: _ConverterTypeThree.contract, event: "OwnerUpdate", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdate is a free log subscription operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_ConverterTypeThree *ConverterTypeThreeFilterer) WatchOwnerUpdate(opts *bind.WatchOpts, sink chan<- *ConverterTypeThreeOwnerUpdate, _prevOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ConverterTypeThree.contract.WatchLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeThreeOwnerUpdate)
				if err := _ConverterTypeThree.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
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
func (_ConverterTypeThree *ConverterTypeThreeFilterer) ParseOwnerUpdate(log types.Log) (*ConverterTypeThreeOwnerUpdate, error) {
	event := new(ConverterTypeThreeOwnerUpdate)
	if err := _ConverterTypeThree.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterTypeThreeTokenRateUpdateIterator is returned from FilterTokenRateUpdate and is used to iterate over the raw logs and unpacked data for TokenRateUpdate events raised by the ConverterTypeThree contract.
type ConverterTypeThreeTokenRateUpdateIterator struct {
	Event *ConverterTypeThreeTokenRateUpdate // Event containing the contract specifics and raw log

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
func (it *ConverterTypeThreeTokenRateUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeThreeTokenRateUpdate)
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
		it.Event = new(ConverterTypeThreeTokenRateUpdate)
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
func (it *ConverterTypeThreeTokenRateUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeThreeTokenRateUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeThreeTokenRateUpdate represents a TokenRateUpdate event raised by the ConverterTypeThree contract.
type ConverterTypeThreeTokenRateUpdate struct {
	Token1 common.Address
	Token2 common.Address
	RateN  *big.Int
	RateD  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenRateUpdate is a free log retrieval operation binding the contract event 0x77f29993cf2c084e726f7e802da0719d6a0ade3e204badc7a3ffd57ecb768c24.
//
// Solidity: event TokenRateUpdate(address indexed _token1, address indexed _token2, uint256 _rateN, uint256 _rateD)
func (_ConverterTypeThree *ConverterTypeThreeFilterer) FilterTokenRateUpdate(opts *bind.FilterOpts, _token1 []common.Address, _token2 []common.Address) (*ConverterTypeThreeTokenRateUpdateIterator, error) {

	var _token1Rule []interface{}
	for _, _token1Item := range _token1 {
		_token1Rule = append(_token1Rule, _token1Item)
	}
	var _token2Rule []interface{}
	for _, _token2Item := range _token2 {
		_token2Rule = append(_token2Rule, _token2Item)
	}

	logs, sub, err := _ConverterTypeThree.contract.FilterLogs(opts, "TokenRateUpdate", _token1Rule, _token2Rule)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeThreeTokenRateUpdateIterator{contract: _ConverterTypeThree.contract, event: "TokenRateUpdate", logs: logs, sub: sub}, nil
}

// WatchTokenRateUpdate is a free log subscription operation binding the contract event 0x77f29993cf2c084e726f7e802da0719d6a0ade3e204badc7a3ffd57ecb768c24.
//
// Solidity: event TokenRateUpdate(address indexed _token1, address indexed _token2, uint256 _rateN, uint256 _rateD)
func (_ConverterTypeThree *ConverterTypeThreeFilterer) WatchTokenRateUpdate(opts *bind.WatchOpts, sink chan<- *ConverterTypeThreeTokenRateUpdate, _token1 []common.Address, _token2 []common.Address) (event.Subscription, error) {

	var _token1Rule []interface{}
	for _, _token1Item := range _token1 {
		_token1Rule = append(_token1Rule, _token1Item)
	}
	var _token2Rule []interface{}
	for _, _token2Item := range _token2 {
		_token2Rule = append(_token2Rule, _token2Item)
	}

	logs, sub, err := _ConverterTypeThree.contract.WatchLogs(opts, "TokenRateUpdate", _token1Rule, _token2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeThreeTokenRateUpdate)
				if err := _ConverterTypeThree.contract.UnpackLog(event, "TokenRateUpdate", log); err != nil {
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
func (_ConverterTypeThree *ConverterTypeThreeFilterer) ParseTokenRateUpdate(log types.Log) (*ConverterTypeThreeTokenRateUpdate, error) {
	event := new(ConverterTypeThreeTokenRateUpdate)
	if err := _ConverterTypeThree.contract.UnpackLog(event, "TokenRateUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
