// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package camelotv3pair

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

// Camelotv3pairMetaData contains all meta data concerning the Camelotv3pair contract.
var Camelotv3pairMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"DrainWrongToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"token0FeePercent\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"token1FeePercent\",\"type\":\"uint16\"}],\"name\":\"FeePercentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SetPairTypeImmutable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prevStableSwap\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"stableSwap\",\"type\":\"bool\"}],\"name\":\"SetStableSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Skim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_FEE_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"drainWrongToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint16\",\"name\":\"_token0FeePercent\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_token1FeePercent\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pairTypeImmutable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"precisionMultiplier0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"precisionMultiplier1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newToken0FeePercent\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"newToken1FeePercent\",\"type\":\"uint16\"}],\"name\":\"setFeePercent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setPairTypeImmutable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"expectedReserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"expectedReserve1\",\"type\":\"uint112\"}],\"name\":\"setStableSwap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stableSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token0FeePercent\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token1FeePercent\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Camelotv3pairABI is the input ABI used to generate the binding from.
// Deprecated: Use Camelotv3pairMetaData.ABI instead.
var Camelotv3pairABI = Camelotv3pairMetaData.ABI

// Camelotv3pair is an auto generated Go binding around an Ethereum contract.
type Camelotv3pair struct {
	Camelotv3pairCaller     // Read-only binding to the contract
	Camelotv3pairTransactor // Write-only binding to the contract
	Camelotv3pairFilterer   // Log filterer for contract events
}

// Camelotv3pairCaller is an auto generated read-only Go binding around an Ethereum contract.
type Camelotv3pairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Camelotv3pairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Camelotv3pairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Camelotv3pairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Camelotv3pairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Camelotv3pairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Camelotv3pairSession struct {
	Contract     *Camelotv3pair    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Camelotv3pairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Camelotv3pairCallerSession struct {
	Contract *Camelotv3pairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Camelotv3pairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Camelotv3pairTransactorSession struct {
	Contract     *Camelotv3pairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Camelotv3pairRaw is an auto generated low-level Go binding around an Ethereum contract.
type Camelotv3pairRaw struct {
	Contract *Camelotv3pair // Generic contract binding to access the raw methods on
}

// Camelotv3pairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Camelotv3pairCallerRaw struct {
	Contract *Camelotv3pairCaller // Generic read-only contract binding to access the raw methods on
}

// Camelotv3pairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Camelotv3pairTransactorRaw struct {
	Contract *Camelotv3pairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCamelotv3pair creates a new instance of Camelotv3pair, bound to a specific deployed contract.
func NewCamelotv3pair(address common.Address, backend bind.ContractBackend) (*Camelotv3pair, error) {
	contract, err := bindCamelotv3pair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Camelotv3pair{Camelotv3pairCaller: Camelotv3pairCaller{contract: contract}, Camelotv3pairTransactor: Camelotv3pairTransactor{contract: contract}, Camelotv3pairFilterer: Camelotv3pairFilterer{contract: contract}}, nil
}

// NewCamelotv3pairCaller creates a new read-only instance of Camelotv3pair, bound to a specific deployed contract.
func NewCamelotv3pairCaller(address common.Address, caller bind.ContractCaller) (*Camelotv3pairCaller, error) {
	contract, err := bindCamelotv3pair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Camelotv3pairCaller{contract: contract}, nil
}

// NewCamelotv3pairTransactor creates a new write-only instance of Camelotv3pair, bound to a specific deployed contract.
func NewCamelotv3pairTransactor(address common.Address, transactor bind.ContractTransactor) (*Camelotv3pairTransactor, error) {
	contract, err := bindCamelotv3pair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Camelotv3pairTransactor{contract: contract}, nil
}

// NewCamelotv3pairFilterer creates a new log filterer instance of Camelotv3pair, bound to a specific deployed contract.
func NewCamelotv3pairFilterer(address common.Address, filterer bind.ContractFilterer) (*Camelotv3pairFilterer, error) {
	contract, err := bindCamelotv3pair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Camelotv3pairFilterer{contract: contract}, nil
}

// bindCamelotv3pair binds a generic wrapper to an already deployed contract.
func bindCamelotv3pair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Camelotv3pairMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Camelotv3pair *Camelotv3pairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Camelotv3pair.Contract.Camelotv3pairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Camelotv3pair *Camelotv3pairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Camelotv3pairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Camelotv3pair *Camelotv3pairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Camelotv3pairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Camelotv3pair *Camelotv3pairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Camelotv3pair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Camelotv3pair *Camelotv3pairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Camelotv3pair *Camelotv3pairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Camelotv3pair *Camelotv3pairCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Camelotv3pair *Camelotv3pairSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Camelotv3pair.Contract.DOMAINSEPARATOR(&_Camelotv3pair.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Camelotv3pair *Camelotv3pairCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Camelotv3pair.Contract.DOMAINSEPARATOR(&_Camelotv3pair.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCaller) FEEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "FEE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairSession) FEEDENOMINATOR() (*big.Int, error) {
	return _Camelotv3pair.Contract.FEEDENOMINATOR(&_Camelotv3pair.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCallerSession) FEEDENOMINATOR() (*big.Int, error) {
	return _Camelotv3pair.Contract.FEEDENOMINATOR(&_Camelotv3pair.CallOpts)
}

// MAXFEEPERCENT is a free data retrieval call binding the contract method 0x67d81740.
//
// Solidity: function MAX_FEE_PERCENT() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCaller) MAXFEEPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "MAX_FEE_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEEPERCENT is a free data retrieval call binding the contract method 0x67d81740.
//
// Solidity: function MAX_FEE_PERCENT() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairSession) MAXFEEPERCENT() (*big.Int, error) {
	return _Camelotv3pair.Contract.MAXFEEPERCENT(&_Camelotv3pair.CallOpts)
}

// MAXFEEPERCENT is a free data retrieval call binding the contract method 0x67d81740.
//
// Solidity: function MAX_FEE_PERCENT() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCallerSession) MAXFEEPERCENT() (*big.Int, error) {
	return _Camelotv3pair.Contract.MAXFEEPERCENT(&_Camelotv3pair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Camelotv3pair.Contract.MINIMUMLIQUIDITY(&_Camelotv3pair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Camelotv3pair.Contract.MINIMUMLIQUIDITY(&_Camelotv3pair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Camelotv3pair *Camelotv3pairCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Camelotv3pair *Camelotv3pairSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Camelotv3pair.Contract.PERMITTYPEHASH(&_Camelotv3pair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Camelotv3pair *Camelotv3pairCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Camelotv3pair.Contract.PERMITTYPEHASH(&_Camelotv3pair.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Camelotv3pair *Camelotv3pairSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Camelotv3pair.Contract.Allowance(&_Camelotv3pair.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Camelotv3pair.Contract.Allowance(&_Camelotv3pair.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Camelotv3pair *Camelotv3pairSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Camelotv3pair.Contract.BalanceOf(&_Camelotv3pair.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Camelotv3pair.Contract.BalanceOf(&_Camelotv3pair.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Camelotv3pair *Camelotv3pairCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Camelotv3pair *Camelotv3pairSession) Decimals() (uint8, error) {
	return _Camelotv3pair.Contract.Decimals(&_Camelotv3pair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Camelotv3pair *Camelotv3pairCallerSession) Decimals() (uint8, error) {
	return _Camelotv3pair.Contract.Decimals(&_Camelotv3pair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Camelotv3pair *Camelotv3pairCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Camelotv3pair *Camelotv3pairSession) Factory() (common.Address, error) {
	return _Camelotv3pair.Contract.Factory(&_Camelotv3pair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Camelotv3pair *Camelotv3pairCallerSession) Factory() (common.Address, error) {
	return _Camelotv3pair.Contract.Factory(&_Camelotv3pair.CallOpts)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn) view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, tokenIn common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "getAmountOut", amountIn, tokenIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn) view returns(uint256)
func (_Camelotv3pair *Camelotv3pairSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address) (*big.Int, error) {
	return _Camelotv3pair.Contract.GetAmountOut(&_Camelotv3pair.CallOpts, amountIn, tokenIn)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn) view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCallerSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address) (*big.Int, error) {
	return _Camelotv3pair.Contract.GetAmountOut(&_Camelotv3pair.CallOpts, amountIn, tokenIn)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint16 _token0FeePercent, uint16 _token1FeePercent)
func (_Camelotv3pair *Camelotv3pairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0         *big.Int
	Reserve1         *big.Int
	Token0FeePercent uint16
	Token1FeePercent uint16
}, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0         *big.Int
		Reserve1         *big.Int
		Token0FeePercent uint16
		Token1FeePercent uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Token0FeePercent = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.Token1FeePercent = *abi.ConvertType(out[3], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint16 _token0FeePercent, uint16 _token1FeePercent)
func (_Camelotv3pair *Camelotv3pairSession) GetReserves() (struct {
	Reserve0         *big.Int
	Reserve1         *big.Int
	Token0FeePercent uint16
	Token1FeePercent uint16
}, error) {
	return _Camelotv3pair.Contract.GetReserves(&_Camelotv3pair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint16 _token0FeePercent, uint16 _token1FeePercent)
func (_Camelotv3pair *Camelotv3pairCallerSession) GetReserves() (struct {
	Reserve0         *big.Int
	Reserve1         *big.Int
	Token0FeePercent uint16
	Token1FeePercent uint16
}, error) {
	return _Camelotv3pair.Contract.GetReserves(&_Camelotv3pair.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Camelotv3pair *Camelotv3pairCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Camelotv3pair *Camelotv3pairSession) Initialized() (bool, error) {
	return _Camelotv3pair.Contract.Initialized(&_Camelotv3pair.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Camelotv3pair *Camelotv3pairCallerSession) Initialized() (bool, error) {
	return _Camelotv3pair.Contract.Initialized(&_Camelotv3pair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairSession) KLast() (*big.Int, error) {
	return _Camelotv3pair.Contract.KLast(&_Camelotv3pair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCallerSession) KLast() (*big.Int, error) {
	return _Camelotv3pair.Contract.KLast(&_Camelotv3pair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Camelotv3pair *Camelotv3pairCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Camelotv3pair *Camelotv3pairSession) Name() (string, error) {
	return _Camelotv3pair.Contract.Name(&_Camelotv3pair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Camelotv3pair *Camelotv3pairCallerSession) Name() (string, error) {
	return _Camelotv3pair.Contract.Name(&_Camelotv3pair.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Camelotv3pair *Camelotv3pairSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Camelotv3pair.Contract.Nonces(&_Camelotv3pair.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Camelotv3pair.Contract.Nonces(&_Camelotv3pair.CallOpts, arg0)
}

// PairTypeImmutable is a free data retrieval call binding the contract method 0xb6200b07.
//
// Solidity: function pairTypeImmutable() view returns(bool)
func (_Camelotv3pair *Camelotv3pairCaller) PairTypeImmutable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "pairTypeImmutable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PairTypeImmutable is a free data retrieval call binding the contract method 0xb6200b07.
//
// Solidity: function pairTypeImmutable() view returns(bool)
func (_Camelotv3pair *Camelotv3pairSession) PairTypeImmutable() (bool, error) {
	return _Camelotv3pair.Contract.PairTypeImmutable(&_Camelotv3pair.CallOpts)
}

// PairTypeImmutable is a free data retrieval call binding the contract method 0xb6200b07.
//
// Solidity: function pairTypeImmutable() view returns(bool)
func (_Camelotv3pair *Camelotv3pairCallerSession) PairTypeImmutable() (bool, error) {
	return _Camelotv3pair.Contract.PairTypeImmutable(&_Camelotv3pair.CallOpts)
}

// PrecisionMultiplier0 is a free data retrieval call binding the contract method 0x3b9f1dc0.
//
// Solidity: function precisionMultiplier0() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCaller) PrecisionMultiplier0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "precisionMultiplier0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrecisionMultiplier0 is a free data retrieval call binding the contract method 0x3b9f1dc0.
//
// Solidity: function precisionMultiplier0() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairSession) PrecisionMultiplier0() (*big.Int, error) {
	return _Camelotv3pair.Contract.PrecisionMultiplier0(&_Camelotv3pair.CallOpts)
}

// PrecisionMultiplier0 is a free data retrieval call binding the contract method 0x3b9f1dc0.
//
// Solidity: function precisionMultiplier0() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCallerSession) PrecisionMultiplier0() (*big.Int, error) {
	return _Camelotv3pair.Contract.PrecisionMultiplier0(&_Camelotv3pair.CallOpts)
}

// PrecisionMultiplier1 is a free data retrieval call binding the contract method 0x288e5d02.
//
// Solidity: function precisionMultiplier1() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCaller) PrecisionMultiplier1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "precisionMultiplier1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrecisionMultiplier1 is a free data retrieval call binding the contract method 0x288e5d02.
//
// Solidity: function precisionMultiplier1() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairSession) PrecisionMultiplier1() (*big.Int, error) {
	return _Camelotv3pair.Contract.PrecisionMultiplier1(&_Camelotv3pair.CallOpts)
}

// PrecisionMultiplier1 is a free data retrieval call binding the contract method 0x288e5d02.
//
// Solidity: function precisionMultiplier1() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCallerSession) PrecisionMultiplier1() (*big.Int, error) {
	return _Camelotv3pair.Contract.PrecisionMultiplier1(&_Camelotv3pair.CallOpts)
}

// StableSwap is a free data retrieval call binding the contract method 0x9e548b7f.
//
// Solidity: function stableSwap() view returns(bool)
func (_Camelotv3pair *Camelotv3pairCaller) StableSwap(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "stableSwap")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StableSwap is a free data retrieval call binding the contract method 0x9e548b7f.
//
// Solidity: function stableSwap() view returns(bool)
func (_Camelotv3pair *Camelotv3pairSession) StableSwap() (bool, error) {
	return _Camelotv3pair.Contract.StableSwap(&_Camelotv3pair.CallOpts)
}

// StableSwap is a free data retrieval call binding the contract method 0x9e548b7f.
//
// Solidity: function stableSwap() view returns(bool)
func (_Camelotv3pair *Camelotv3pairCallerSession) StableSwap() (bool, error) {
	return _Camelotv3pair.Contract.StableSwap(&_Camelotv3pair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Camelotv3pair *Camelotv3pairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Camelotv3pair *Camelotv3pairSession) Symbol() (string, error) {
	return _Camelotv3pair.Contract.Symbol(&_Camelotv3pair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Camelotv3pair *Camelotv3pairCallerSession) Symbol() (string, error) {
	return _Camelotv3pair.Contract.Symbol(&_Camelotv3pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Camelotv3pair *Camelotv3pairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Camelotv3pair *Camelotv3pairSession) Token0() (common.Address, error) {
	return _Camelotv3pair.Contract.Token0(&_Camelotv3pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Camelotv3pair *Camelotv3pairCallerSession) Token0() (common.Address, error) {
	return _Camelotv3pair.Contract.Token0(&_Camelotv3pair.CallOpts)
}

// Token0FeePercent is a free data retrieval call binding the contract method 0x62ecec03.
//
// Solidity: function token0FeePercent() view returns(uint16)
func (_Camelotv3pair *Camelotv3pairCaller) Token0FeePercent(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "token0FeePercent")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Token0FeePercent is a free data retrieval call binding the contract method 0x62ecec03.
//
// Solidity: function token0FeePercent() view returns(uint16)
func (_Camelotv3pair *Camelotv3pairSession) Token0FeePercent() (uint16, error) {
	return _Camelotv3pair.Contract.Token0FeePercent(&_Camelotv3pair.CallOpts)
}

// Token0FeePercent is a free data retrieval call binding the contract method 0x62ecec03.
//
// Solidity: function token0FeePercent() view returns(uint16)
func (_Camelotv3pair *Camelotv3pairCallerSession) Token0FeePercent() (uint16, error) {
	return _Camelotv3pair.Contract.Token0FeePercent(&_Camelotv3pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Camelotv3pair *Camelotv3pairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Camelotv3pair *Camelotv3pairSession) Token1() (common.Address, error) {
	return _Camelotv3pair.Contract.Token1(&_Camelotv3pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Camelotv3pair *Camelotv3pairCallerSession) Token1() (common.Address, error) {
	return _Camelotv3pair.Contract.Token1(&_Camelotv3pair.CallOpts)
}

// Token1FeePercent is a free data retrieval call binding the contract method 0x2fcd1692.
//
// Solidity: function token1FeePercent() view returns(uint16)
func (_Camelotv3pair *Camelotv3pairCaller) Token1FeePercent(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "token1FeePercent")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Token1FeePercent is a free data retrieval call binding the contract method 0x2fcd1692.
//
// Solidity: function token1FeePercent() view returns(uint16)
func (_Camelotv3pair *Camelotv3pairSession) Token1FeePercent() (uint16, error) {
	return _Camelotv3pair.Contract.Token1FeePercent(&_Camelotv3pair.CallOpts)
}

// Token1FeePercent is a free data retrieval call binding the contract method 0x2fcd1692.
//
// Solidity: function token1FeePercent() view returns(uint16)
func (_Camelotv3pair *Camelotv3pairCallerSession) Token1FeePercent() (uint16, error) {
	return _Camelotv3pair.Contract.Token1FeePercent(&_Camelotv3pair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Camelotv3pair.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairSession) TotalSupply() (*big.Int, error) {
	return _Camelotv3pair.Contract.TotalSupply(&_Camelotv3pair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Camelotv3pair *Camelotv3pairCallerSession) TotalSupply() (*big.Int, error) {
	return _Camelotv3pair.Contract.TotalSupply(&_Camelotv3pair.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Camelotv3pair *Camelotv3pairTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Camelotv3pair.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Camelotv3pair *Camelotv3pairSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Approve(&_Camelotv3pair.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Camelotv3pair *Camelotv3pairTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Approve(&_Camelotv3pair.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Camelotv3pair *Camelotv3pairTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Camelotv3pair *Camelotv3pairSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Burn(&_Camelotv3pair.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Camelotv3pair *Camelotv3pairTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Burn(&_Camelotv3pair.TransactOpts, to)
}

// DrainWrongToken is a paid mutator transaction binding the contract method 0xf39ac11f.
//
// Solidity: function drainWrongToken(address token, address to) returns()
func (_Camelotv3pair *Camelotv3pairTransactor) DrainWrongToken(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.contract.Transact(opts, "drainWrongToken", token, to)
}

// DrainWrongToken is a paid mutator transaction binding the contract method 0xf39ac11f.
//
// Solidity: function drainWrongToken(address token, address to) returns()
func (_Camelotv3pair *Camelotv3pairSession) DrainWrongToken(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.DrainWrongToken(&_Camelotv3pair.TransactOpts, token, to)
}

// DrainWrongToken is a paid mutator transaction binding the contract method 0xf39ac11f.
//
// Solidity: function drainWrongToken(address token, address to) returns()
func (_Camelotv3pair *Camelotv3pairTransactorSession) DrainWrongToken(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.DrainWrongToken(&_Camelotv3pair.TransactOpts, token, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_Camelotv3pair *Camelotv3pairTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.contract.Transact(opts, "initialize", _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_Camelotv3pair *Camelotv3pairSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Initialize(&_Camelotv3pair.TransactOpts, _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_Camelotv3pair *Camelotv3pairTransactorSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Initialize(&_Camelotv3pair.TransactOpts, _token0, _token1)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Camelotv3pair *Camelotv3pairTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Camelotv3pair *Camelotv3pairSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Mint(&_Camelotv3pair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Camelotv3pair *Camelotv3pairTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Mint(&_Camelotv3pair.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Camelotv3pair *Camelotv3pairTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Camelotv3pair.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Camelotv3pair *Camelotv3pairSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Permit(&_Camelotv3pair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Camelotv3pair *Camelotv3pairTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Permit(&_Camelotv3pair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0x48e5d260.
//
// Solidity: function setFeePercent(uint16 newToken0FeePercent, uint16 newToken1FeePercent) returns()
func (_Camelotv3pair *Camelotv3pairTransactor) SetFeePercent(opts *bind.TransactOpts, newToken0FeePercent uint16, newToken1FeePercent uint16) (*types.Transaction, error) {
	return _Camelotv3pair.contract.Transact(opts, "setFeePercent", newToken0FeePercent, newToken1FeePercent)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0x48e5d260.
//
// Solidity: function setFeePercent(uint16 newToken0FeePercent, uint16 newToken1FeePercent) returns()
func (_Camelotv3pair *Camelotv3pairSession) SetFeePercent(newToken0FeePercent uint16, newToken1FeePercent uint16) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.SetFeePercent(&_Camelotv3pair.TransactOpts, newToken0FeePercent, newToken1FeePercent)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0x48e5d260.
//
// Solidity: function setFeePercent(uint16 newToken0FeePercent, uint16 newToken1FeePercent) returns()
func (_Camelotv3pair *Camelotv3pairTransactorSession) SetFeePercent(newToken0FeePercent uint16, newToken1FeePercent uint16) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.SetFeePercent(&_Camelotv3pair.TransactOpts, newToken0FeePercent, newToken1FeePercent)
}

// SetPairTypeImmutable is a paid mutator transaction binding the contract method 0x3ba17077.
//
// Solidity: function setPairTypeImmutable() returns()
func (_Camelotv3pair *Camelotv3pairTransactor) SetPairTypeImmutable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Camelotv3pair.contract.Transact(opts, "setPairTypeImmutable")
}

// SetPairTypeImmutable is a paid mutator transaction binding the contract method 0x3ba17077.
//
// Solidity: function setPairTypeImmutable() returns()
func (_Camelotv3pair *Camelotv3pairSession) SetPairTypeImmutable() (*types.Transaction, error) {
	return _Camelotv3pair.Contract.SetPairTypeImmutable(&_Camelotv3pair.TransactOpts)
}

// SetPairTypeImmutable is a paid mutator transaction binding the contract method 0x3ba17077.
//
// Solidity: function setPairTypeImmutable() returns()
func (_Camelotv3pair *Camelotv3pairTransactorSession) SetPairTypeImmutable() (*types.Transaction, error) {
	return _Camelotv3pair.Contract.SetPairTypeImmutable(&_Camelotv3pair.TransactOpts)
}

// SetStableSwap is a paid mutator transaction binding the contract method 0x3029e5d4.
//
// Solidity: function setStableSwap(bool stable, uint112 expectedReserve0, uint112 expectedReserve1) returns()
func (_Camelotv3pair *Camelotv3pairTransactor) SetStableSwap(opts *bind.TransactOpts, stable bool, expectedReserve0 *big.Int, expectedReserve1 *big.Int) (*types.Transaction, error) {
	return _Camelotv3pair.contract.Transact(opts, "setStableSwap", stable, expectedReserve0, expectedReserve1)
}

// SetStableSwap is a paid mutator transaction binding the contract method 0x3029e5d4.
//
// Solidity: function setStableSwap(bool stable, uint112 expectedReserve0, uint112 expectedReserve1) returns()
func (_Camelotv3pair *Camelotv3pairSession) SetStableSwap(stable bool, expectedReserve0 *big.Int, expectedReserve1 *big.Int) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.SetStableSwap(&_Camelotv3pair.TransactOpts, stable, expectedReserve0, expectedReserve1)
}

// SetStableSwap is a paid mutator transaction binding the contract method 0x3029e5d4.
//
// Solidity: function setStableSwap(bool stable, uint112 expectedReserve0, uint112 expectedReserve1) returns()
func (_Camelotv3pair *Camelotv3pairTransactorSession) SetStableSwap(stable bool, expectedReserve0 *big.Int, expectedReserve1 *big.Int) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.SetStableSwap(&_Camelotv3pair.TransactOpts, stable, expectedReserve0, expectedReserve1)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Camelotv3pair *Camelotv3pairTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Camelotv3pair *Camelotv3pairSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Skim(&_Camelotv3pair.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Camelotv3pair *Camelotv3pairTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Skim(&_Camelotv3pair.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Camelotv3pair *Camelotv3pairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Camelotv3pair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Camelotv3pair *Camelotv3pairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Swap(&_Camelotv3pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Camelotv3pair *Camelotv3pairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Swap(&_Camelotv3pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap0 is a paid mutator transaction binding the contract method 0x6e1fdd7f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data, address referrer) returns()
func (_Camelotv3pair *Camelotv3pairTransactor) Swap0(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte, referrer common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.contract.Transact(opts, "swap0", amount0Out, amount1Out, to, data, referrer)
}

// Swap0 is a paid mutator transaction binding the contract method 0x6e1fdd7f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data, address referrer) returns()
func (_Camelotv3pair *Camelotv3pairSession) Swap0(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte, referrer common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Swap0(&_Camelotv3pair.TransactOpts, amount0Out, amount1Out, to, data, referrer)
}

// Swap0 is a paid mutator transaction binding the contract method 0x6e1fdd7f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data, address referrer) returns()
func (_Camelotv3pair *Camelotv3pairTransactorSession) Swap0(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte, referrer common.Address) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Swap0(&_Camelotv3pair.TransactOpts, amount0Out, amount1Out, to, data, referrer)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Camelotv3pair *Camelotv3pairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Camelotv3pair.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Camelotv3pair *Camelotv3pairSession) Sync() (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Sync(&_Camelotv3pair.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Camelotv3pair *Camelotv3pairTransactorSession) Sync() (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Sync(&_Camelotv3pair.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Camelotv3pair *Camelotv3pairTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Camelotv3pair.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Camelotv3pair *Camelotv3pairSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Transfer(&_Camelotv3pair.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Camelotv3pair *Camelotv3pairTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.Transfer(&_Camelotv3pair.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Camelotv3pair *Camelotv3pairTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Camelotv3pair.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Camelotv3pair *Camelotv3pairSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.TransferFrom(&_Camelotv3pair.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Camelotv3pair *Camelotv3pairTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Camelotv3pair.Contract.TransferFrom(&_Camelotv3pair.TransactOpts, from, to, value)
}

// Camelotv3pairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Camelotv3pair contract.
type Camelotv3pairApprovalIterator struct {
	Event *Camelotv3pairApproval // Event containing the contract specifics and raw log

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
func (it *Camelotv3pairApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3pairApproval)
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
		it.Event = new(Camelotv3pairApproval)
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
func (it *Camelotv3pairApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3pairApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3pairApproval represents a Approval event raised by the Camelotv3pair contract.
type Camelotv3pairApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Camelotv3pair *Camelotv3pairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Camelotv3pairApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Camelotv3pair.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Camelotv3pairApprovalIterator{contract: _Camelotv3pair.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Camelotv3pair *Camelotv3pairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Camelotv3pairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Camelotv3pair.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3pairApproval)
				if err := _Camelotv3pair.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Camelotv3pair *Camelotv3pairFilterer) ParseApproval(log types.Log) (*Camelotv3pairApproval, error) {
	event := new(Camelotv3pairApproval)
	if err := _Camelotv3pair.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Camelotv3pairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Camelotv3pair contract.
type Camelotv3pairBurnIterator struct {
	Event *Camelotv3pairBurn // Event containing the contract specifics and raw log

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
func (it *Camelotv3pairBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3pairBurn)
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
		it.Event = new(Camelotv3pairBurn)
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
func (it *Camelotv3pairBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3pairBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3pairBurn represents a Burn event raised by the Camelotv3pair contract.
type Camelotv3pairBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Camelotv3pair *Camelotv3pairFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*Camelotv3pairBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Camelotv3pair.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Camelotv3pairBurnIterator{contract: _Camelotv3pair.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Camelotv3pair *Camelotv3pairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *Camelotv3pairBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Camelotv3pair.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3pairBurn)
				if err := _Camelotv3pair.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Camelotv3pair *Camelotv3pairFilterer) ParseBurn(log types.Log) (*Camelotv3pairBurn, error) {
	event := new(Camelotv3pairBurn)
	if err := _Camelotv3pair.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Camelotv3pairDrainWrongTokenIterator is returned from FilterDrainWrongToken and is used to iterate over the raw logs and unpacked data for DrainWrongToken events raised by the Camelotv3pair contract.
type Camelotv3pairDrainWrongTokenIterator struct {
	Event *Camelotv3pairDrainWrongToken // Event containing the contract specifics and raw log

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
func (it *Camelotv3pairDrainWrongTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3pairDrainWrongToken)
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
		it.Event = new(Camelotv3pairDrainWrongToken)
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
func (it *Camelotv3pairDrainWrongTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3pairDrainWrongTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3pairDrainWrongToken represents a DrainWrongToken event raised by the Camelotv3pair contract.
type Camelotv3pairDrainWrongToken struct {
	Token common.Address
	To    common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDrainWrongToken is a free log retrieval operation binding the contract event 0x368a9dc863ecb94b5ba32a682e26295b10d9c2666fad7d785ebdf262c3c52413.
//
// Solidity: event DrainWrongToken(address indexed token, address to)
func (_Camelotv3pair *Camelotv3pairFilterer) FilterDrainWrongToken(opts *bind.FilterOpts, token []common.Address) (*Camelotv3pairDrainWrongTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Camelotv3pair.contract.FilterLogs(opts, "DrainWrongToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &Camelotv3pairDrainWrongTokenIterator{contract: _Camelotv3pair.contract, event: "DrainWrongToken", logs: logs, sub: sub}, nil
}

// WatchDrainWrongToken is a free log subscription operation binding the contract event 0x368a9dc863ecb94b5ba32a682e26295b10d9c2666fad7d785ebdf262c3c52413.
//
// Solidity: event DrainWrongToken(address indexed token, address to)
func (_Camelotv3pair *Camelotv3pairFilterer) WatchDrainWrongToken(opts *bind.WatchOpts, sink chan<- *Camelotv3pairDrainWrongToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Camelotv3pair.contract.WatchLogs(opts, "DrainWrongToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3pairDrainWrongToken)
				if err := _Camelotv3pair.contract.UnpackLog(event, "DrainWrongToken", log); err != nil {
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

// ParseDrainWrongToken is a log parse operation binding the contract event 0x368a9dc863ecb94b5ba32a682e26295b10d9c2666fad7d785ebdf262c3c52413.
//
// Solidity: event DrainWrongToken(address indexed token, address to)
func (_Camelotv3pair *Camelotv3pairFilterer) ParseDrainWrongToken(log types.Log) (*Camelotv3pairDrainWrongToken, error) {
	event := new(Camelotv3pairDrainWrongToken)
	if err := _Camelotv3pair.contract.UnpackLog(event, "DrainWrongToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Camelotv3pairFeePercentUpdatedIterator is returned from FilterFeePercentUpdated and is used to iterate over the raw logs and unpacked data for FeePercentUpdated events raised by the Camelotv3pair contract.
type Camelotv3pairFeePercentUpdatedIterator struct {
	Event *Camelotv3pairFeePercentUpdated // Event containing the contract specifics and raw log

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
func (it *Camelotv3pairFeePercentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3pairFeePercentUpdated)
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
		it.Event = new(Camelotv3pairFeePercentUpdated)
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
func (it *Camelotv3pairFeePercentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3pairFeePercentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3pairFeePercentUpdated represents a FeePercentUpdated event raised by the Camelotv3pair contract.
type Camelotv3pairFeePercentUpdated struct {
	Token0FeePercent uint16
	Token1FeePercent uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFeePercentUpdated is a free log retrieval operation binding the contract event 0xa4877b8ecb5a00ba277e4bceeeb187a669e7113649774dfbea05c259ce27f17b.
//
// Solidity: event FeePercentUpdated(uint16 token0FeePercent, uint16 token1FeePercent)
func (_Camelotv3pair *Camelotv3pairFilterer) FilterFeePercentUpdated(opts *bind.FilterOpts) (*Camelotv3pairFeePercentUpdatedIterator, error) {

	logs, sub, err := _Camelotv3pair.contract.FilterLogs(opts, "FeePercentUpdated")
	if err != nil {
		return nil, err
	}
	return &Camelotv3pairFeePercentUpdatedIterator{contract: _Camelotv3pair.contract, event: "FeePercentUpdated", logs: logs, sub: sub}, nil
}

// WatchFeePercentUpdated is a free log subscription operation binding the contract event 0xa4877b8ecb5a00ba277e4bceeeb187a669e7113649774dfbea05c259ce27f17b.
//
// Solidity: event FeePercentUpdated(uint16 token0FeePercent, uint16 token1FeePercent)
func (_Camelotv3pair *Camelotv3pairFilterer) WatchFeePercentUpdated(opts *bind.WatchOpts, sink chan<- *Camelotv3pairFeePercentUpdated) (event.Subscription, error) {

	logs, sub, err := _Camelotv3pair.contract.WatchLogs(opts, "FeePercentUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3pairFeePercentUpdated)
				if err := _Camelotv3pair.contract.UnpackLog(event, "FeePercentUpdated", log); err != nil {
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

// ParseFeePercentUpdated is a log parse operation binding the contract event 0xa4877b8ecb5a00ba277e4bceeeb187a669e7113649774dfbea05c259ce27f17b.
//
// Solidity: event FeePercentUpdated(uint16 token0FeePercent, uint16 token1FeePercent)
func (_Camelotv3pair *Camelotv3pairFilterer) ParseFeePercentUpdated(log types.Log) (*Camelotv3pairFeePercentUpdated, error) {
	event := new(Camelotv3pairFeePercentUpdated)
	if err := _Camelotv3pair.contract.UnpackLog(event, "FeePercentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Camelotv3pairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Camelotv3pair contract.
type Camelotv3pairMintIterator struct {
	Event *Camelotv3pairMint // Event containing the contract specifics and raw log

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
func (it *Camelotv3pairMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3pairMint)
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
		it.Event = new(Camelotv3pairMint)
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
func (it *Camelotv3pairMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3pairMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3pairMint represents a Mint event raised by the Camelotv3pair contract.
type Camelotv3pairMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Camelotv3pair *Camelotv3pairFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*Camelotv3pairMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Camelotv3pair.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &Camelotv3pairMintIterator{contract: _Camelotv3pair.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Camelotv3pair *Camelotv3pairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *Camelotv3pairMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Camelotv3pair.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3pairMint)
				if err := _Camelotv3pair.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Camelotv3pair *Camelotv3pairFilterer) ParseMint(log types.Log) (*Camelotv3pairMint, error) {
	event := new(Camelotv3pairMint)
	if err := _Camelotv3pair.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Camelotv3pairSetPairTypeImmutableIterator is returned from FilterSetPairTypeImmutable and is used to iterate over the raw logs and unpacked data for SetPairTypeImmutable events raised by the Camelotv3pair contract.
type Camelotv3pairSetPairTypeImmutableIterator struct {
	Event *Camelotv3pairSetPairTypeImmutable // Event containing the contract specifics and raw log

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
func (it *Camelotv3pairSetPairTypeImmutableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3pairSetPairTypeImmutable)
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
		it.Event = new(Camelotv3pairSetPairTypeImmutable)
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
func (it *Camelotv3pairSetPairTypeImmutableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3pairSetPairTypeImmutableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3pairSetPairTypeImmutable represents a SetPairTypeImmutable event raised by the Camelotv3pair contract.
type Camelotv3pairSetPairTypeImmutable struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetPairTypeImmutable is a free log retrieval operation binding the contract event 0x09122c41ae733a4d7740324d50e35fbd6ee85be3c1312a45596d8045150ab2f2.
//
// Solidity: event SetPairTypeImmutable()
func (_Camelotv3pair *Camelotv3pairFilterer) FilterSetPairTypeImmutable(opts *bind.FilterOpts) (*Camelotv3pairSetPairTypeImmutableIterator, error) {

	logs, sub, err := _Camelotv3pair.contract.FilterLogs(opts, "SetPairTypeImmutable")
	if err != nil {
		return nil, err
	}
	return &Camelotv3pairSetPairTypeImmutableIterator{contract: _Camelotv3pair.contract, event: "SetPairTypeImmutable", logs: logs, sub: sub}, nil
}

// WatchSetPairTypeImmutable is a free log subscription operation binding the contract event 0x09122c41ae733a4d7740324d50e35fbd6ee85be3c1312a45596d8045150ab2f2.
//
// Solidity: event SetPairTypeImmutable()
func (_Camelotv3pair *Camelotv3pairFilterer) WatchSetPairTypeImmutable(opts *bind.WatchOpts, sink chan<- *Camelotv3pairSetPairTypeImmutable) (event.Subscription, error) {

	logs, sub, err := _Camelotv3pair.contract.WatchLogs(opts, "SetPairTypeImmutable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3pairSetPairTypeImmutable)
				if err := _Camelotv3pair.contract.UnpackLog(event, "SetPairTypeImmutable", log); err != nil {
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

// ParseSetPairTypeImmutable is a log parse operation binding the contract event 0x09122c41ae733a4d7740324d50e35fbd6ee85be3c1312a45596d8045150ab2f2.
//
// Solidity: event SetPairTypeImmutable()
func (_Camelotv3pair *Camelotv3pairFilterer) ParseSetPairTypeImmutable(log types.Log) (*Camelotv3pairSetPairTypeImmutable, error) {
	event := new(Camelotv3pairSetPairTypeImmutable)
	if err := _Camelotv3pair.contract.UnpackLog(event, "SetPairTypeImmutable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Camelotv3pairSetStableSwapIterator is returned from FilterSetStableSwap and is used to iterate over the raw logs and unpacked data for SetStableSwap events raised by the Camelotv3pair contract.
type Camelotv3pairSetStableSwapIterator struct {
	Event *Camelotv3pairSetStableSwap // Event containing the contract specifics and raw log

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
func (it *Camelotv3pairSetStableSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3pairSetStableSwap)
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
		it.Event = new(Camelotv3pairSetStableSwap)
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
func (it *Camelotv3pairSetStableSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3pairSetStableSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3pairSetStableSwap represents a SetStableSwap event raised by the Camelotv3pair contract.
type Camelotv3pairSetStableSwap struct {
	PrevStableSwap bool
	StableSwap     bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetStableSwap is a free log retrieval operation binding the contract event 0xb6a86710bde53aa7fb1b3856279e2af5b476d53e2dd0902cf17a0911b5a43a8b.
//
// Solidity: event SetStableSwap(bool prevStableSwap, bool stableSwap)
func (_Camelotv3pair *Camelotv3pairFilterer) FilterSetStableSwap(opts *bind.FilterOpts) (*Camelotv3pairSetStableSwapIterator, error) {

	logs, sub, err := _Camelotv3pair.contract.FilterLogs(opts, "SetStableSwap")
	if err != nil {
		return nil, err
	}
	return &Camelotv3pairSetStableSwapIterator{contract: _Camelotv3pair.contract, event: "SetStableSwap", logs: logs, sub: sub}, nil
}

// WatchSetStableSwap is a free log subscription operation binding the contract event 0xb6a86710bde53aa7fb1b3856279e2af5b476d53e2dd0902cf17a0911b5a43a8b.
//
// Solidity: event SetStableSwap(bool prevStableSwap, bool stableSwap)
func (_Camelotv3pair *Camelotv3pairFilterer) WatchSetStableSwap(opts *bind.WatchOpts, sink chan<- *Camelotv3pairSetStableSwap) (event.Subscription, error) {

	logs, sub, err := _Camelotv3pair.contract.WatchLogs(opts, "SetStableSwap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3pairSetStableSwap)
				if err := _Camelotv3pair.contract.UnpackLog(event, "SetStableSwap", log); err != nil {
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

// ParseSetStableSwap is a log parse operation binding the contract event 0xb6a86710bde53aa7fb1b3856279e2af5b476d53e2dd0902cf17a0911b5a43a8b.
//
// Solidity: event SetStableSwap(bool prevStableSwap, bool stableSwap)
func (_Camelotv3pair *Camelotv3pairFilterer) ParseSetStableSwap(log types.Log) (*Camelotv3pairSetStableSwap, error) {
	event := new(Camelotv3pairSetStableSwap)
	if err := _Camelotv3pair.contract.UnpackLog(event, "SetStableSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Camelotv3pairSkimIterator is returned from FilterSkim and is used to iterate over the raw logs and unpacked data for Skim events raised by the Camelotv3pair contract.
type Camelotv3pairSkimIterator struct {
	Event *Camelotv3pairSkim // Event containing the contract specifics and raw log

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
func (it *Camelotv3pairSkimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3pairSkim)
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
		it.Event = new(Camelotv3pairSkim)
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
func (it *Camelotv3pairSkimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3pairSkimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3pairSkim represents a Skim event raised by the Camelotv3pair contract.
type Camelotv3pairSkim struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSkim is a free log retrieval operation binding the contract event 0x21ad22495c9c75cd1c94756f91824e779c0c8a8e168b267c790df464fe056b79.
//
// Solidity: event Skim()
func (_Camelotv3pair *Camelotv3pairFilterer) FilterSkim(opts *bind.FilterOpts) (*Camelotv3pairSkimIterator, error) {

	logs, sub, err := _Camelotv3pair.contract.FilterLogs(opts, "Skim")
	if err != nil {
		return nil, err
	}
	return &Camelotv3pairSkimIterator{contract: _Camelotv3pair.contract, event: "Skim", logs: logs, sub: sub}, nil
}

// WatchSkim is a free log subscription operation binding the contract event 0x21ad22495c9c75cd1c94756f91824e779c0c8a8e168b267c790df464fe056b79.
//
// Solidity: event Skim()
func (_Camelotv3pair *Camelotv3pairFilterer) WatchSkim(opts *bind.WatchOpts, sink chan<- *Camelotv3pairSkim) (event.Subscription, error) {

	logs, sub, err := _Camelotv3pair.contract.WatchLogs(opts, "Skim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3pairSkim)
				if err := _Camelotv3pair.contract.UnpackLog(event, "Skim", log); err != nil {
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

// ParseSkim is a log parse operation binding the contract event 0x21ad22495c9c75cd1c94756f91824e779c0c8a8e168b267c790df464fe056b79.
//
// Solidity: event Skim()
func (_Camelotv3pair *Camelotv3pairFilterer) ParseSkim(log types.Log) (*Camelotv3pairSkim, error) {
	event := new(Camelotv3pairSkim)
	if err := _Camelotv3pair.contract.UnpackLog(event, "Skim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Camelotv3pairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Camelotv3pair contract.
type Camelotv3pairSwapIterator struct {
	Event *Camelotv3pairSwap // Event containing the contract specifics and raw log

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
func (it *Camelotv3pairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3pairSwap)
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
		it.Event = new(Camelotv3pairSwap)
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
func (it *Camelotv3pairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3pairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3pairSwap represents a Swap event raised by the Camelotv3pair contract.
type Camelotv3pairSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Camelotv3pair *Camelotv3pairFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*Camelotv3pairSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Camelotv3pair.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Camelotv3pairSwapIterator{contract: _Camelotv3pair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Camelotv3pair *Camelotv3pairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *Camelotv3pairSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Camelotv3pair.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3pairSwap)
				if err := _Camelotv3pair.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Camelotv3pair *Camelotv3pairFilterer) ParseSwap(log types.Log) (*Camelotv3pairSwap, error) {
	event := new(Camelotv3pairSwap)
	if err := _Camelotv3pair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Camelotv3pairSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the Camelotv3pair contract.
type Camelotv3pairSyncIterator struct {
	Event *Camelotv3pairSync // Event containing the contract specifics and raw log

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
func (it *Camelotv3pairSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3pairSync)
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
		it.Event = new(Camelotv3pairSync)
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
func (it *Camelotv3pairSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3pairSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3pairSync represents a Sync event raised by the Camelotv3pair contract.
type Camelotv3pairSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Camelotv3pair *Camelotv3pairFilterer) FilterSync(opts *bind.FilterOpts) (*Camelotv3pairSyncIterator, error) {

	logs, sub, err := _Camelotv3pair.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &Camelotv3pairSyncIterator{contract: _Camelotv3pair.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Camelotv3pair *Camelotv3pairFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *Camelotv3pairSync) (event.Subscription, error) {

	logs, sub, err := _Camelotv3pair.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3pairSync)
				if err := _Camelotv3pair.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Camelotv3pair *Camelotv3pairFilterer) ParseSync(log types.Log) (*Camelotv3pairSync, error) {
	event := new(Camelotv3pairSync)
	if err := _Camelotv3pair.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Camelotv3pairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Camelotv3pair contract.
type Camelotv3pairTransferIterator struct {
	Event *Camelotv3pairTransfer // Event containing the contract specifics and raw log

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
func (it *Camelotv3pairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3pairTransfer)
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
		it.Event = new(Camelotv3pairTransfer)
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
func (it *Camelotv3pairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3pairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3pairTransfer represents a Transfer event raised by the Camelotv3pair contract.
type Camelotv3pairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Camelotv3pair *Camelotv3pairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Camelotv3pairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Camelotv3pair.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Camelotv3pairTransferIterator{contract: _Camelotv3pair.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Camelotv3pair *Camelotv3pairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Camelotv3pairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Camelotv3pair.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3pairTransfer)
				if err := _Camelotv3pair.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Camelotv3pair *Camelotv3pairFilterer) ParseTransfer(log types.Log) (*Camelotv3pairTransfer, error) {
	event := new(Camelotv3pairTransfer)
	if err := _Camelotv3pair.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
