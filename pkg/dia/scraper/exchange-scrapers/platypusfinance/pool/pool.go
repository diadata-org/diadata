// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pool

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

// PoolMetaData contains all meta data concerning the Pool contract.
var PoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"AssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDev\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDev\",\"type\":\"address\"}],\"name\":\"DevUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousHaircut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newHaircut\",\"type\":\"uint256\"}],\"name\":\"HaircutRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"OracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousPriceDeviation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPriceDeviation\",\"type\":\"uint256\"}],\"name\":\"PriceDeviationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousRetentionRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRetentionRatio\",\"type\":\"uint256\"}],\"name\":\"RetentionRatioUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousK\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newK\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousN\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newN\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousC1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newC1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousXThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newXThreshold\",\"type\":\"uint256\"}],\"name\":\"SlippageParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"addAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"assetOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getC1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEquilibriumCoverageRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHaircutRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxPriceDeviation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRetentionRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSlippageParamK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSlippageParamN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getXThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wantedToken\",\"type\":\"address\"}],\"name\":\"quoteMaxInitialAssetWithdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxInitialAssetAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"}],\"name\":\"quotePotentialSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"potentialOutcome\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"haircut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"quotePotentialWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"enoughCash\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wantedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"quotePotentialWithdrawFromOtherAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"recoverUserFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"}],\"name\":\"removeAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"haircutRate_\",\"type\":\"uint256\"}],\"name\":\"setHaircutRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxPriceDeviation_\",\"type\":\"uint256\"}],\"name\":\"setMaxPriceDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceOracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"retentionRatio_\",\"type\":\"uint256\"}],\"name\":\"setRetentionRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"k_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"c1_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"xThreshold_\",\"type\":\"uint256\"}],\"name\":\"setSlippageParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumToAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualToAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"haircut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wantedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"withdrawFromOtherAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMetaData.ABI instead.
var PoolABI = PoolMetaData.ABI

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// AssetOf is a free data retrieval call binding the contract method 0x71f96211.
//
// Solidity: function assetOf(address token) view returns(address)
func (_Pool *PoolCaller) AssetOf(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "assetOf", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetOf is a free data retrieval call binding the contract method 0x71f96211.
//
// Solidity: function assetOf(address token) view returns(address)
func (_Pool *PoolSession) AssetOf(token common.Address) (common.Address, error) {
	return _Pool.Contract.AssetOf(&_Pool.CallOpts, token)
}

// AssetOf is a free data retrieval call binding the contract method 0x71f96211.
//
// Solidity: function assetOf(address token) view returns(address)
func (_Pool *PoolCallerSession) AssetOf(token common.Address) (common.Address, error) {
	return _Pool.Contract.AssetOf(&_Pool.CallOpts, token)
}

// GetC1 is a free data retrieval call binding the contract method 0xa76f54d2.
//
// Solidity: function getC1() view returns(uint256)
func (_Pool *PoolCaller) GetC1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getC1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetC1 is a free data retrieval call binding the contract method 0xa76f54d2.
//
// Solidity: function getC1() view returns(uint256)
func (_Pool *PoolSession) GetC1() (*big.Int, error) {
	return _Pool.Contract.GetC1(&_Pool.CallOpts)
}

// GetC1 is a free data retrieval call binding the contract method 0xa76f54d2.
//
// Solidity: function getC1() view returns(uint256)
func (_Pool *PoolCallerSession) GetC1() (*big.Int, error) {
	return _Pool.Contract.GetC1(&_Pool.CallOpts)
}

// GetDev is a free data retrieval call binding the contract method 0x09bb9267.
//
// Solidity: function getDev() view returns(address)
func (_Pool *PoolCaller) GetDev(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getDev")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDev is a free data retrieval call binding the contract method 0x09bb9267.
//
// Solidity: function getDev() view returns(address)
func (_Pool *PoolSession) GetDev() (common.Address, error) {
	return _Pool.Contract.GetDev(&_Pool.CallOpts)
}

// GetDev is a free data retrieval call binding the contract method 0x09bb9267.
//
// Solidity: function getDev() view returns(address)
func (_Pool *PoolCallerSession) GetDev() (common.Address, error) {
	return _Pool.Contract.GetDev(&_Pool.CallOpts)
}

// GetEquilibriumCoverageRatio is a free data retrieval call binding the contract method 0x05f7bc26.
//
// Solidity: function getEquilibriumCoverageRatio() view returns(uint256)
func (_Pool *PoolCaller) GetEquilibriumCoverageRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getEquilibriumCoverageRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEquilibriumCoverageRatio is a free data retrieval call binding the contract method 0x05f7bc26.
//
// Solidity: function getEquilibriumCoverageRatio() view returns(uint256)
func (_Pool *PoolSession) GetEquilibriumCoverageRatio() (*big.Int, error) {
	return _Pool.Contract.GetEquilibriumCoverageRatio(&_Pool.CallOpts)
}

// GetEquilibriumCoverageRatio is a free data retrieval call binding the contract method 0x05f7bc26.
//
// Solidity: function getEquilibriumCoverageRatio() view returns(uint256)
func (_Pool *PoolCallerSession) GetEquilibriumCoverageRatio() (*big.Int, error) {
	return _Pool.Contract.GetEquilibriumCoverageRatio(&_Pool.CallOpts)
}

// GetHaircutRate is a free data retrieval call binding the contract method 0x7fdd5a8e.
//
// Solidity: function getHaircutRate() view returns(uint256)
func (_Pool *PoolCaller) GetHaircutRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getHaircutRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHaircutRate is a free data retrieval call binding the contract method 0x7fdd5a8e.
//
// Solidity: function getHaircutRate() view returns(uint256)
func (_Pool *PoolSession) GetHaircutRate() (*big.Int, error) {
	return _Pool.Contract.GetHaircutRate(&_Pool.CallOpts)
}

// GetHaircutRate is a free data retrieval call binding the contract method 0x7fdd5a8e.
//
// Solidity: function getHaircutRate() view returns(uint256)
func (_Pool *PoolCallerSession) GetHaircutRate() (*big.Int, error) {
	return _Pool.Contract.GetHaircutRate(&_Pool.CallOpts)
}

// GetMaxPriceDeviation is a free data retrieval call binding the contract method 0xddcbc516.
//
// Solidity: function getMaxPriceDeviation() view returns(uint256)
func (_Pool *PoolCaller) GetMaxPriceDeviation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getMaxPriceDeviation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPriceDeviation is a free data retrieval call binding the contract method 0xddcbc516.
//
// Solidity: function getMaxPriceDeviation() view returns(uint256)
func (_Pool *PoolSession) GetMaxPriceDeviation() (*big.Int, error) {
	return _Pool.Contract.GetMaxPriceDeviation(&_Pool.CallOpts)
}

// GetMaxPriceDeviation is a free data retrieval call binding the contract method 0xddcbc516.
//
// Solidity: function getMaxPriceDeviation() view returns(uint256)
func (_Pool *PoolCallerSession) GetMaxPriceDeviation() (*big.Int, error) {
	return _Pool.Contract.GetMaxPriceDeviation(&_Pool.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_Pool *PoolCaller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_Pool *PoolSession) GetPriceOracle() (common.Address, error) {
	return _Pool.Contract.GetPriceOracle(&_Pool.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_Pool *PoolCallerSession) GetPriceOracle() (common.Address, error) {
	return _Pool.Contract.GetPriceOracle(&_Pool.CallOpts)
}

// GetRetentionRatio is a free data retrieval call binding the contract method 0xcb733d7a.
//
// Solidity: function getRetentionRatio() view returns(uint256)
func (_Pool *PoolCaller) GetRetentionRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getRetentionRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRetentionRatio is a free data retrieval call binding the contract method 0xcb733d7a.
//
// Solidity: function getRetentionRatio() view returns(uint256)
func (_Pool *PoolSession) GetRetentionRatio() (*big.Int, error) {
	return _Pool.Contract.GetRetentionRatio(&_Pool.CallOpts)
}

// GetRetentionRatio is a free data retrieval call binding the contract method 0xcb733d7a.
//
// Solidity: function getRetentionRatio() view returns(uint256)
func (_Pool *PoolCallerSession) GetRetentionRatio() (*big.Int, error) {
	return _Pool.Contract.GetRetentionRatio(&_Pool.CallOpts)
}

// GetSlippageParamK is a free data retrieval call binding the contract method 0x55af008a.
//
// Solidity: function getSlippageParamK() view returns(uint256)
func (_Pool *PoolCaller) GetSlippageParamK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getSlippageParamK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlippageParamK is a free data retrieval call binding the contract method 0x55af008a.
//
// Solidity: function getSlippageParamK() view returns(uint256)
func (_Pool *PoolSession) GetSlippageParamK() (*big.Int, error) {
	return _Pool.Contract.GetSlippageParamK(&_Pool.CallOpts)
}

// GetSlippageParamK is a free data retrieval call binding the contract method 0x55af008a.
//
// Solidity: function getSlippageParamK() view returns(uint256)
func (_Pool *PoolCallerSession) GetSlippageParamK() (*big.Int, error) {
	return _Pool.Contract.GetSlippageParamK(&_Pool.CallOpts)
}

// GetSlippageParamN is a free data retrieval call binding the contract method 0x7727c655.
//
// Solidity: function getSlippageParamN() view returns(uint256)
func (_Pool *PoolCaller) GetSlippageParamN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getSlippageParamN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlippageParamN is a free data retrieval call binding the contract method 0x7727c655.
//
// Solidity: function getSlippageParamN() view returns(uint256)
func (_Pool *PoolSession) GetSlippageParamN() (*big.Int, error) {
	return _Pool.Contract.GetSlippageParamN(&_Pool.CallOpts)
}

// GetSlippageParamN is a free data retrieval call binding the contract method 0x7727c655.
//
// Solidity: function getSlippageParamN() view returns(uint256)
func (_Pool *PoolCallerSession) GetSlippageParamN() (*big.Int, error) {
	return _Pool.Contract.GetSlippageParamN(&_Pool.CallOpts)
}

// GetTokenAddresses is a free data retrieval call binding the contract method 0xee8c24b8.
//
// Solidity: function getTokenAddresses() view returns(address[])
func (_Pool *PoolCaller) GetTokenAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getTokenAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenAddresses is a free data retrieval call binding the contract method 0xee8c24b8.
//
// Solidity: function getTokenAddresses() view returns(address[])
func (_Pool *PoolSession) GetTokenAddresses() ([]common.Address, error) {
	return _Pool.Contract.GetTokenAddresses(&_Pool.CallOpts)
}

// GetTokenAddresses is a free data retrieval call binding the contract method 0xee8c24b8.
//
// Solidity: function getTokenAddresses() view returns(address[])
func (_Pool *PoolCallerSession) GetTokenAddresses() ([]common.Address, error) {
	return _Pool.Contract.GetTokenAddresses(&_Pool.CallOpts)
}

// GetXThreshold is a free data retrieval call binding the contract method 0x7a1c36d1.
//
// Solidity: function getXThreshold() view returns(uint256)
func (_Pool *PoolCaller) GetXThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getXThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetXThreshold is a free data retrieval call binding the contract method 0x7a1c36d1.
//
// Solidity: function getXThreshold() view returns(uint256)
func (_Pool *PoolSession) GetXThreshold() (*big.Int, error) {
	return _Pool.Contract.GetXThreshold(&_Pool.CallOpts)
}

// GetXThreshold is a free data retrieval call binding the contract method 0x7a1c36d1.
//
// Solidity: function getXThreshold() view returns(uint256)
func (_Pool *PoolCallerSession) GetXThreshold() (*big.Int, error) {
	return _Pool.Contract.GetXThreshold(&_Pool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolCallerSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pool *PoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pool *PoolSession) Paused() (bool, error) {
	return _Pool.Contract.Paused(&_Pool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pool *PoolCallerSession) Paused() (bool, error) {
	return _Pool.Contract.Paused(&_Pool.CallOpts)
}

// QuoteMaxInitialAssetWithdrawable is a free data retrieval call binding the contract method 0x2a803b5e.
//
// Solidity: function quoteMaxInitialAssetWithdrawable(address initialToken, address wantedToken) view returns(uint256 maxInitialAssetAmount)
func (_Pool *PoolCaller) QuoteMaxInitialAssetWithdrawable(opts *bind.CallOpts, initialToken common.Address, wantedToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "quoteMaxInitialAssetWithdrawable", initialToken, wantedToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteMaxInitialAssetWithdrawable is a free data retrieval call binding the contract method 0x2a803b5e.
//
// Solidity: function quoteMaxInitialAssetWithdrawable(address initialToken, address wantedToken) view returns(uint256 maxInitialAssetAmount)
func (_Pool *PoolSession) QuoteMaxInitialAssetWithdrawable(initialToken common.Address, wantedToken common.Address) (*big.Int, error) {
	return _Pool.Contract.QuoteMaxInitialAssetWithdrawable(&_Pool.CallOpts, initialToken, wantedToken)
}

// QuoteMaxInitialAssetWithdrawable is a free data retrieval call binding the contract method 0x2a803b5e.
//
// Solidity: function quoteMaxInitialAssetWithdrawable(address initialToken, address wantedToken) view returns(uint256 maxInitialAssetAmount)
func (_Pool *PoolCallerSession) QuoteMaxInitialAssetWithdrawable(initialToken common.Address, wantedToken common.Address) (*big.Int, error) {
	return _Pool.Contract.QuoteMaxInitialAssetWithdrawable(&_Pool.CallOpts, initialToken, wantedToken)
}

// QuotePotentialSwap is a free data retrieval call binding the contract method 0x43c2e2f5.
//
// Solidity: function quotePotentialSwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 potentialOutcome, uint256 haircut)
func (_Pool *PoolCaller) QuotePotentialSwap(opts *bind.CallOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int) (struct {
	PotentialOutcome *big.Int
	Haircut          *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "quotePotentialSwap", fromToken, toToken, fromAmount)

	outstruct := new(struct {
		PotentialOutcome *big.Int
		Haircut          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PotentialOutcome = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Haircut = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuotePotentialSwap is a free data retrieval call binding the contract method 0x43c2e2f5.
//
// Solidity: function quotePotentialSwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 potentialOutcome, uint256 haircut)
func (_Pool *PoolSession) QuotePotentialSwap(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (struct {
	PotentialOutcome *big.Int
	Haircut          *big.Int
}, error) {
	return _Pool.Contract.QuotePotentialSwap(&_Pool.CallOpts, fromToken, toToken, fromAmount)
}

// QuotePotentialSwap is a free data retrieval call binding the contract method 0x43c2e2f5.
//
// Solidity: function quotePotentialSwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 potentialOutcome, uint256 haircut)
func (_Pool *PoolCallerSession) QuotePotentialSwap(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (struct {
	PotentialOutcome *big.Int
	Haircut          *big.Int
}, error) {
	return _Pool.Contract.QuotePotentialSwap(&_Pool.CallOpts, fromToken, toToken, fromAmount)
}

// QuotePotentialWithdraw is a free data retrieval call binding the contract method 0x907448ed.
//
// Solidity: function quotePotentialWithdraw(address token, uint256 liquidity) view returns(uint256 amount, uint256 fee, bool enoughCash)
func (_Pool *PoolCaller) QuotePotentialWithdraw(opts *bind.CallOpts, token common.Address, liquidity *big.Int) (struct {
	Amount     *big.Int
	Fee        *big.Int
	EnoughCash bool
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "quotePotentialWithdraw", token, liquidity)

	outstruct := new(struct {
		Amount     *big.Int
		Fee        *big.Int
		EnoughCash bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EnoughCash = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// QuotePotentialWithdraw is a free data retrieval call binding the contract method 0x907448ed.
//
// Solidity: function quotePotentialWithdraw(address token, uint256 liquidity) view returns(uint256 amount, uint256 fee, bool enoughCash)
func (_Pool *PoolSession) QuotePotentialWithdraw(token common.Address, liquidity *big.Int) (struct {
	Amount     *big.Int
	Fee        *big.Int
	EnoughCash bool
}, error) {
	return _Pool.Contract.QuotePotentialWithdraw(&_Pool.CallOpts, token, liquidity)
}

// QuotePotentialWithdraw is a free data retrieval call binding the contract method 0x907448ed.
//
// Solidity: function quotePotentialWithdraw(address token, uint256 liquidity) view returns(uint256 amount, uint256 fee, bool enoughCash)
func (_Pool *PoolCallerSession) QuotePotentialWithdraw(token common.Address, liquidity *big.Int) (struct {
	Amount     *big.Int
	Fee        *big.Int
	EnoughCash bool
}, error) {
	return _Pool.Contract.QuotePotentialWithdraw(&_Pool.CallOpts, token, liquidity)
}

// QuotePotentialWithdrawFromOtherAsset is a free data retrieval call binding the contract method 0xa4275ceb.
//
// Solidity: function quotePotentialWithdrawFromOtherAsset(address initialToken, address wantedToken, uint256 liquidity) view returns(uint256 amount, uint256 fee)
func (_Pool *PoolCaller) QuotePotentialWithdrawFromOtherAsset(opts *bind.CallOpts, initialToken common.Address, wantedToken common.Address, liquidity *big.Int) (struct {
	Amount *big.Int
	Fee    *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "quotePotentialWithdrawFromOtherAsset", initialToken, wantedToken, liquidity)

	outstruct := new(struct {
		Amount *big.Int
		Fee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuotePotentialWithdrawFromOtherAsset is a free data retrieval call binding the contract method 0xa4275ceb.
//
// Solidity: function quotePotentialWithdrawFromOtherAsset(address initialToken, address wantedToken, uint256 liquidity) view returns(uint256 amount, uint256 fee)
func (_Pool *PoolSession) QuotePotentialWithdrawFromOtherAsset(initialToken common.Address, wantedToken common.Address, liquidity *big.Int) (struct {
	Amount *big.Int
	Fee    *big.Int
}, error) {
	return _Pool.Contract.QuotePotentialWithdrawFromOtherAsset(&_Pool.CallOpts, initialToken, wantedToken, liquidity)
}

// QuotePotentialWithdrawFromOtherAsset is a free data retrieval call binding the contract method 0xa4275ceb.
//
// Solidity: function quotePotentialWithdrawFromOtherAsset(address initialToken, address wantedToken, uint256 liquidity) view returns(uint256 amount, uint256 fee)
func (_Pool *PoolCallerSession) QuotePotentialWithdrawFromOtherAsset(initialToken common.Address, wantedToken common.Address, liquidity *big.Int) (struct {
	Amount *big.Int
	Fee    *big.Int
}, error) {
	return _Pool.Contract.QuotePotentialWithdrawFromOtherAsset(&_Pool.CallOpts, initialToken, wantedToken, liquidity)
}

// AddAsset is a paid mutator transaction binding the contract method 0xda489997.
//
// Solidity: function addAsset(address token, address asset) returns()
func (_Pool *PoolTransactor) AddAsset(opts *bind.TransactOpts, token common.Address, asset common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "addAsset", token, asset)
}

// AddAsset is a paid mutator transaction binding the contract method 0xda489997.
//
// Solidity: function addAsset(address token, address asset) returns()
func (_Pool *PoolSession) AddAsset(token common.Address, asset common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AddAsset(&_Pool.TransactOpts, token, asset)
}

// AddAsset is a paid mutator transaction binding the contract method 0xda489997.
//
// Solidity: function addAsset(address token, address asset) returns()
func (_Pool *PoolTransactorSession) AddAsset(token common.Address, asset common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AddAsset(&_Pool.TransactOpts, token, asset)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address token, uint256 amount, address to, uint256 deadline) returns(uint256 liquidity)
func (_Pool *PoolTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "deposit", token, amount, to, deadline)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address token, uint256 amount, address to, uint256 deadline) returns(uint256 liquidity)
func (_Pool *PoolSession) Deposit(token common.Address, amount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Deposit(&_Pool.TransactOpts, token, amount, to, deadline)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address token, uint256 amount, address to, uint256 deadline) returns(uint256 liquidity)
func (_Pool *PoolTransactorSession) Deposit(token common.Address, amount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Deposit(&_Pool.TransactOpts, token, amount, to, deadline)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Pool *PoolTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Pool *PoolSession) Initialize() (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Pool *PoolTransactorSession) Initialize() (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pool *PoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pool *PoolSession) Pause() (*types.Transaction, error) {
	return _Pool.Contract.Pause(&_Pool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pool *PoolTransactorSession) Pause() (*types.Transaction, error) {
	return _Pool.Contract.Pause(&_Pool.TransactOpts)
}

// RecoverUserFunds is a paid mutator transaction binding the contract method 0x0d4636b5.
//
// Solidity: function recoverUserFunds(address token) returns()
func (_Pool *PoolTransactor) RecoverUserFunds(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "recoverUserFunds", token)
}

// RecoverUserFunds is a paid mutator transaction binding the contract method 0x0d4636b5.
//
// Solidity: function recoverUserFunds(address token) returns()
func (_Pool *PoolSession) RecoverUserFunds(token common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RecoverUserFunds(&_Pool.TransactOpts, token)
}

// RecoverUserFunds is a paid mutator transaction binding the contract method 0x0d4636b5.
//
// Solidity: function recoverUserFunds(address token) returns()
func (_Pool *PoolTransactorSession) RecoverUserFunds(token common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RecoverUserFunds(&_Pool.TransactOpts, token)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x4a5e42b1.
//
// Solidity: function removeAsset(address key) returns()
func (_Pool *PoolTransactor) RemoveAsset(opts *bind.TransactOpts, key common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "removeAsset", key)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x4a5e42b1.
//
// Solidity: function removeAsset(address key) returns()
func (_Pool *PoolSession) RemoveAsset(key common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RemoveAsset(&_Pool.TransactOpts, key)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x4a5e42b1.
//
// Solidity: function removeAsset(address key) returns()
func (_Pool *PoolTransactorSession) RemoveAsset(key common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RemoveAsset(&_Pool.TransactOpts, key)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pool.Contract.RenounceOwnership(&_Pool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pool.Contract.RenounceOwnership(&_Pool.TransactOpts)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev) returns()
func (_Pool *PoolTransactor) SetDev(opts *bind.TransactOpts, dev common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setDev", dev)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev) returns()
func (_Pool *PoolSession) SetDev(dev common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetDev(&_Pool.TransactOpts, dev)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev) returns()
func (_Pool *PoolTransactorSession) SetDev(dev common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetDev(&_Pool.TransactOpts, dev)
}

// SetHaircutRate is a paid mutator transaction binding the contract method 0xf57e84d5.
//
// Solidity: function setHaircutRate(uint256 haircutRate_) returns()
func (_Pool *PoolTransactor) SetHaircutRate(opts *bind.TransactOpts, haircutRate_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setHaircutRate", haircutRate_)
}

// SetHaircutRate is a paid mutator transaction binding the contract method 0xf57e84d5.
//
// Solidity: function setHaircutRate(uint256 haircutRate_) returns()
func (_Pool *PoolSession) SetHaircutRate(haircutRate_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetHaircutRate(&_Pool.TransactOpts, haircutRate_)
}

// SetHaircutRate is a paid mutator transaction binding the contract method 0xf57e84d5.
//
// Solidity: function setHaircutRate(uint256 haircutRate_) returns()
func (_Pool *PoolTransactorSession) SetHaircutRate(haircutRate_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetHaircutRate(&_Pool.TransactOpts, haircutRate_)
}

// SetMaxPriceDeviation is a paid mutator transaction binding the contract method 0x9ee4c057.
//
// Solidity: function setMaxPriceDeviation(uint256 maxPriceDeviation_) returns()
func (_Pool *PoolTransactor) SetMaxPriceDeviation(opts *bind.TransactOpts, maxPriceDeviation_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setMaxPriceDeviation", maxPriceDeviation_)
}

// SetMaxPriceDeviation is a paid mutator transaction binding the contract method 0x9ee4c057.
//
// Solidity: function setMaxPriceDeviation(uint256 maxPriceDeviation_) returns()
func (_Pool *PoolSession) SetMaxPriceDeviation(maxPriceDeviation_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetMaxPriceDeviation(&_Pool.TransactOpts, maxPriceDeviation_)
}

// SetMaxPriceDeviation is a paid mutator transaction binding the contract method 0x9ee4c057.
//
// Solidity: function setMaxPriceDeviation(uint256 maxPriceDeviation_) returns()
func (_Pool *PoolTransactorSession) SetMaxPriceDeviation(maxPriceDeviation_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetMaxPriceDeviation(&_Pool.TransactOpts, maxPriceDeviation_)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_Pool *PoolTransactor) SetPriceOracle(opts *bind.TransactOpts, priceOracle common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setPriceOracle", priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_Pool *PoolSession) SetPriceOracle(priceOracle common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetPriceOracle(&_Pool.TransactOpts, priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_Pool *PoolTransactorSession) SetPriceOracle(priceOracle common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetPriceOracle(&_Pool.TransactOpts, priceOracle)
}

// SetRetentionRatio is a paid mutator transaction binding the contract method 0x44db964a.
//
// Solidity: function setRetentionRatio(uint256 retentionRatio_) returns()
func (_Pool *PoolTransactor) SetRetentionRatio(opts *bind.TransactOpts, retentionRatio_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setRetentionRatio", retentionRatio_)
}

// SetRetentionRatio is a paid mutator transaction binding the contract method 0x44db964a.
//
// Solidity: function setRetentionRatio(uint256 retentionRatio_) returns()
func (_Pool *PoolSession) SetRetentionRatio(retentionRatio_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetRetentionRatio(&_Pool.TransactOpts, retentionRatio_)
}

// SetRetentionRatio is a paid mutator transaction binding the contract method 0x44db964a.
//
// Solidity: function setRetentionRatio(uint256 retentionRatio_) returns()
func (_Pool *PoolTransactorSession) SetRetentionRatio(retentionRatio_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetRetentionRatio(&_Pool.TransactOpts, retentionRatio_)
}

// SetSlippageParams is a paid mutator transaction binding the contract method 0xdf220181.
//
// Solidity: function setSlippageParams(uint256 k_, uint256 n_, uint256 c1_, uint256 xThreshold_) returns()
func (_Pool *PoolTransactor) SetSlippageParams(opts *bind.TransactOpts, k_ *big.Int, n_ *big.Int, c1_ *big.Int, xThreshold_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setSlippageParams", k_, n_, c1_, xThreshold_)
}

// SetSlippageParams is a paid mutator transaction binding the contract method 0xdf220181.
//
// Solidity: function setSlippageParams(uint256 k_, uint256 n_, uint256 c1_, uint256 xThreshold_) returns()
func (_Pool *PoolSession) SetSlippageParams(k_ *big.Int, n_ *big.Int, c1_ *big.Int, xThreshold_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetSlippageParams(&_Pool.TransactOpts, k_, n_, c1_, xThreshold_)
}

// SetSlippageParams is a paid mutator transaction binding the contract method 0xdf220181.
//
// Solidity: function setSlippageParams(uint256 k_, uint256 n_, uint256 c1_, uint256 xThreshold_) returns()
func (_Pool *PoolTransactorSession) SetSlippageParams(k_ *big.Int, n_ *big.Int, c1_ *big.Int, xThreshold_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetSlippageParams(&_Pool.TransactOpts, k_, n_, c1_, xThreshold_)
}

// Swap is a paid mutator transaction binding the contract method 0x9908fc8b.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minimumToAmount, address to, uint256 deadline) returns(uint256 actualToAmount, uint256 haircut)
func (_Pool *PoolTransactor) Swap(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int, minimumToAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "swap", fromToken, toToken, fromAmount, minimumToAmount, to, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x9908fc8b.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minimumToAmount, address to, uint256 deadline) returns(uint256 actualToAmount, uint256 haircut)
func (_Pool *PoolSession) Swap(fromToken common.Address, toToken common.Address, fromAmount *big.Int, minimumToAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Swap(&_Pool.TransactOpts, fromToken, toToken, fromAmount, minimumToAmount, to, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x9908fc8b.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minimumToAmount, address to, uint256 deadline) returns(uint256 actualToAmount, uint256 haircut)
func (_Pool *PoolTransactorSession) Swap(fromToken common.Address, toToken common.Address, fromAmount *big.Int, minimumToAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Swap(&_Pool.TransactOpts, fromToken, toToken, fromAmount, minimumToAmount, to, deadline)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.TransferOwnership(&_Pool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.TransferOwnership(&_Pool.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pool *PoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pool *PoolSession) Unpause() (*types.Transaction, error) {
	return _Pool.Contract.Unpause(&_Pool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pool *PoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pool.Contract.Unpause(&_Pool.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x09a5fca3.
//
// Solidity: function withdraw(address token, uint256 liquidity, uint256 minimumAmount, address to, uint256 deadline) returns(uint256 amount)
func (_Pool *PoolTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, minimumAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "withdraw", token, liquidity, minimumAmount, to, deadline)
}

// Withdraw is a paid mutator transaction binding the contract method 0x09a5fca3.
//
// Solidity: function withdraw(address token, uint256 liquidity, uint256 minimumAmount, address to, uint256 deadline) returns(uint256 amount)
func (_Pool *PoolSession) Withdraw(token common.Address, liquidity *big.Int, minimumAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Withdraw(&_Pool.TransactOpts, token, liquidity, minimumAmount, to, deadline)
}

// Withdraw is a paid mutator transaction binding the contract method 0x09a5fca3.
//
// Solidity: function withdraw(address token, uint256 liquidity, uint256 minimumAmount, address to, uint256 deadline) returns(uint256 amount)
func (_Pool *PoolTransactorSession) Withdraw(token common.Address, liquidity *big.Int, minimumAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Withdraw(&_Pool.TransactOpts, token, liquidity, minimumAmount, to, deadline)
}

// WithdrawFromOtherAsset is a paid mutator transaction binding the contract method 0x0f91f06f.
//
// Solidity: function withdrawFromOtherAsset(address initialToken, address wantedToken, uint256 liquidity, uint256 minimumAmount, address to, uint256 deadline) returns(uint256 amount)
func (_Pool *PoolTransactor) WithdrawFromOtherAsset(opts *bind.TransactOpts, initialToken common.Address, wantedToken common.Address, liquidity *big.Int, minimumAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "withdrawFromOtherAsset", initialToken, wantedToken, liquidity, minimumAmount, to, deadline)
}

// WithdrawFromOtherAsset is a paid mutator transaction binding the contract method 0x0f91f06f.
//
// Solidity: function withdrawFromOtherAsset(address initialToken, address wantedToken, uint256 liquidity, uint256 minimumAmount, address to, uint256 deadline) returns(uint256 amount)
func (_Pool *PoolSession) WithdrawFromOtherAsset(initialToken common.Address, wantedToken common.Address, liquidity *big.Int, minimumAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.WithdrawFromOtherAsset(&_Pool.TransactOpts, initialToken, wantedToken, liquidity, minimumAmount, to, deadline)
}

// WithdrawFromOtherAsset is a paid mutator transaction binding the contract method 0x0f91f06f.
//
// Solidity: function withdrawFromOtherAsset(address initialToken, address wantedToken, uint256 liquidity, uint256 minimumAmount, address to, uint256 deadline) returns(uint256 amount)
func (_Pool *PoolTransactorSession) WithdrawFromOtherAsset(initialToken common.Address, wantedToken common.Address, liquidity *big.Int, minimumAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.WithdrawFromOtherAsset(&_Pool.TransactOpts, initialToken, wantedToken, liquidity, minimumAmount, to, deadline)
}

// PoolAssetAddedIterator is returned from FilterAssetAdded and is used to iterate over the raw logs and unpacked data for AssetAdded events raised by the Pool contract.
type PoolAssetAddedIterator struct {
	Event *PoolAssetAdded // Event containing the contract specifics and raw log

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
func (it *PoolAssetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAssetAdded)
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
		it.Event = new(PoolAssetAdded)
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
func (it *PoolAssetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAssetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAssetAdded represents a AssetAdded event raised by the Pool contract.
type PoolAssetAdded struct {
	Token common.Address
	Asset common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetAdded is a free log retrieval operation binding the contract event 0x0bb5715f0f217c2fe9a0c877ea87d474380c641102f3440ee2a4c8b9d9790918.
//
// Solidity: event AssetAdded(address indexed token, address indexed asset)
func (_Pool *PoolFilterer) FilterAssetAdded(opts *bind.FilterOpts, token []common.Address, asset []common.Address) (*PoolAssetAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "AssetAdded", tokenRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &PoolAssetAddedIterator{contract: _Pool.contract, event: "AssetAdded", logs: logs, sub: sub}, nil
}

// WatchAssetAdded is a free log subscription operation binding the contract event 0x0bb5715f0f217c2fe9a0c877ea87d474380c641102f3440ee2a4c8b9d9790918.
//
// Solidity: event AssetAdded(address indexed token, address indexed asset)
func (_Pool *PoolFilterer) WatchAssetAdded(opts *bind.WatchOpts, sink chan<- *PoolAssetAdded, token []common.Address, asset []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "AssetAdded", tokenRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAssetAdded)
				if err := _Pool.contract.UnpackLog(event, "AssetAdded", log); err != nil {
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

// ParseAssetAdded is a log parse operation binding the contract event 0x0bb5715f0f217c2fe9a0c877ea87d474380c641102f3440ee2a4c8b9d9790918.
//
// Solidity: event AssetAdded(address indexed token, address indexed asset)
func (_Pool *PoolFilterer) ParseAssetAdded(log types.Log) (*PoolAssetAdded, error) {
	event := new(PoolAssetAdded)
	if err := _Pool.contract.UnpackLog(event, "AssetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Pool contract.
type PoolDepositIterator struct {
	Event *PoolDeposit // Event containing the contract specifics and raw log

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
func (it *PoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolDeposit)
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
		it.Event = new(PoolDeposit)
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
func (it *PoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolDeposit represents a Deposit event raised by the Pool contract.
type PoolDeposit struct {
	Sender    common.Address
	Token     common.Address
	Amount    *big.Int
	Liquidity *big.Int
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xf5dd9317b9e63ac316ce44acc85f670b54b339cfa3e9076e1dd55065b922314b.
//
// Solidity: event Deposit(address indexed sender, address token, uint256 amount, uint256 liquidity, address indexed to)
func (_Pool *PoolFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PoolDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolDepositIterator{contract: _Pool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xf5dd9317b9e63ac316ce44acc85f670b54b339cfa3e9076e1dd55065b922314b.
//
// Solidity: event Deposit(address indexed sender, address token, uint256 amount, uint256 liquidity, address indexed to)
func (_Pool *PoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PoolDeposit, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolDeposit)
				if err := _Pool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xf5dd9317b9e63ac316ce44acc85f670b54b339cfa3e9076e1dd55065b922314b.
//
// Solidity: event Deposit(address indexed sender, address token, uint256 amount, uint256 liquidity, address indexed to)
func (_Pool *PoolFilterer) ParseDeposit(log types.Log) (*PoolDeposit, error) {
	event := new(PoolDeposit)
	if err := _Pool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolDevUpdatedIterator is returned from FilterDevUpdated and is used to iterate over the raw logs and unpacked data for DevUpdated events raised by the Pool contract.
type PoolDevUpdatedIterator struct {
	Event *PoolDevUpdated // Event containing the contract specifics and raw log

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
func (it *PoolDevUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolDevUpdated)
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
		it.Event = new(PoolDevUpdated)
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
func (it *PoolDevUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolDevUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolDevUpdated represents a DevUpdated event raised by the Pool contract.
type PoolDevUpdated struct {
	PreviousDev common.Address
	NewDev      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDevUpdated is a free log retrieval operation binding the contract event 0xaa37e6ba252409ccb1daf704a52d3f2fb61265b2f480371d6041193f1cf16ac2.
//
// Solidity: event DevUpdated(address indexed previousDev, address indexed newDev)
func (_Pool *PoolFilterer) FilterDevUpdated(opts *bind.FilterOpts, previousDev []common.Address, newDev []common.Address) (*PoolDevUpdatedIterator, error) {

	var previousDevRule []interface{}
	for _, previousDevItem := range previousDev {
		previousDevRule = append(previousDevRule, previousDevItem)
	}
	var newDevRule []interface{}
	for _, newDevItem := range newDev {
		newDevRule = append(newDevRule, newDevItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "DevUpdated", previousDevRule, newDevRule)
	if err != nil {
		return nil, err
	}
	return &PoolDevUpdatedIterator{contract: _Pool.contract, event: "DevUpdated", logs: logs, sub: sub}, nil
}

// WatchDevUpdated is a free log subscription operation binding the contract event 0xaa37e6ba252409ccb1daf704a52d3f2fb61265b2f480371d6041193f1cf16ac2.
//
// Solidity: event DevUpdated(address indexed previousDev, address indexed newDev)
func (_Pool *PoolFilterer) WatchDevUpdated(opts *bind.WatchOpts, sink chan<- *PoolDevUpdated, previousDev []common.Address, newDev []common.Address) (event.Subscription, error) {

	var previousDevRule []interface{}
	for _, previousDevItem := range previousDev {
		previousDevRule = append(previousDevRule, previousDevItem)
	}
	var newDevRule []interface{}
	for _, newDevItem := range newDev {
		newDevRule = append(newDevRule, newDevItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "DevUpdated", previousDevRule, newDevRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolDevUpdated)
				if err := _Pool.contract.UnpackLog(event, "DevUpdated", log); err != nil {
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

// ParseDevUpdated is a log parse operation binding the contract event 0xaa37e6ba252409ccb1daf704a52d3f2fb61265b2f480371d6041193f1cf16ac2.
//
// Solidity: event DevUpdated(address indexed previousDev, address indexed newDev)
func (_Pool *PoolFilterer) ParseDevUpdated(log types.Log) (*PoolDevUpdated, error) {
	event := new(PoolDevUpdated)
	if err := _Pool.contract.UnpackLog(event, "DevUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolHaircutRateUpdatedIterator is returned from FilterHaircutRateUpdated and is used to iterate over the raw logs and unpacked data for HaircutRateUpdated events raised by the Pool contract.
type PoolHaircutRateUpdatedIterator struct {
	Event *PoolHaircutRateUpdated // Event containing the contract specifics and raw log

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
func (it *PoolHaircutRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolHaircutRateUpdated)
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
		it.Event = new(PoolHaircutRateUpdated)
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
func (it *PoolHaircutRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolHaircutRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolHaircutRateUpdated represents a HaircutRateUpdated event raised by the Pool contract.
type PoolHaircutRateUpdated struct {
	PreviousHaircut *big.Int
	NewHaircut      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterHaircutRateUpdated is a free log retrieval operation binding the contract event 0xbb59363309f944f517c724d17a8f3c89d62efc951e2d7768f7a9d7f307bcad7f.
//
// Solidity: event HaircutRateUpdated(uint256 previousHaircut, uint256 newHaircut)
func (_Pool *PoolFilterer) FilterHaircutRateUpdated(opts *bind.FilterOpts) (*PoolHaircutRateUpdatedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "HaircutRateUpdated")
	if err != nil {
		return nil, err
	}
	return &PoolHaircutRateUpdatedIterator{contract: _Pool.contract, event: "HaircutRateUpdated", logs: logs, sub: sub}, nil
}

// WatchHaircutRateUpdated is a free log subscription operation binding the contract event 0xbb59363309f944f517c724d17a8f3c89d62efc951e2d7768f7a9d7f307bcad7f.
//
// Solidity: event HaircutRateUpdated(uint256 previousHaircut, uint256 newHaircut)
func (_Pool *PoolFilterer) WatchHaircutRateUpdated(opts *bind.WatchOpts, sink chan<- *PoolHaircutRateUpdated) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "HaircutRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolHaircutRateUpdated)
				if err := _Pool.contract.UnpackLog(event, "HaircutRateUpdated", log); err != nil {
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

// ParseHaircutRateUpdated is a log parse operation binding the contract event 0xbb59363309f944f517c724d17a8f3c89d62efc951e2d7768f7a9d7f307bcad7f.
//
// Solidity: event HaircutRateUpdated(uint256 previousHaircut, uint256 newHaircut)
func (_Pool *PoolFilterer) ParseHaircutRateUpdated(log types.Log) (*PoolHaircutRateUpdated, error) {
	event := new(PoolHaircutRateUpdated)
	if err := _Pool.contract.UnpackLog(event, "HaircutRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolOracleUpdatedIterator is returned from FilterOracleUpdated and is used to iterate over the raw logs and unpacked data for OracleUpdated events raised by the Pool contract.
type PoolOracleUpdatedIterator struct {
	Event *PoolOracleUpdated // Event containing the contract specifics and raw log

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
func (it *PoolOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolOracleUpdated)
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
		it.Event = new(PoolOracleUpdated)
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
func (it *PoolOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolOracleUpdated represents a OracleUpdated event raised by the Pool contract.
type PoolOracleUpdated struct {
	PreviousOracle common.Address
	NewOracle      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOracleUpdated is a free log retrieval operation binding the contract event 0x078c3b417dadf69374a59793b829c52001247130433427049317bde56607b1b7.
//
// Solidity: event OracleUpdated(address indexed previousOracle, address indexed newOracle)
func (_Pool *PoolFilterer) FilterOracleUpdated(opts *bind.FilterOpts, previousOracle []common.Address, newOracle []common.Address) (*PoolOracleUpdatedIterator, error) {

	var previousOracleRule []interface{}
	for _, previousOracleItem := range previousOracle {
		previousOracleRule = append(previousOracleRule, previousOracleItem)
	}
	var newOracleRule []interface{}
	for _, newOracleItem := range newOracle {
		newOracleRule = append(newOracleRule, newOracleItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "OracleUpdated", previousOracleRule, newOracleRule)
	if err != nil {
		return nil, err
	}
	return &PoolOracleUpdatedIterator{contract: _Pool.contract, event: "OracleUpdated", logs: logs, sub: sub}, nil
}

// WatchOracleUpdated is a free log subscription operation binding the contract event 0x078c3b417dadf69374a59793b829c52001247130433427049317bde56607b1b7.
//
// Solidity: event OracleUpdated(address indexed previousOracle, address indexed newOracle)
func (_Pool *PoolFilterer) WatchOracleUpdated(opts *bind.WatchOpts, sink chan<- *PoolOracleUpdated, previousOracle []common.Address, newOracle []common.Address) (event.Subscription, error) {

	var previousOracleRule []interface{}
	for _, previousOracleItem := range previousOracle {
		previousOracleRule = append(previousOracleRule, previousOracleItem)
	}
	var newOracleRule []interface{}
	for _, newOracleItem := range newOracle {
		newOracleRule = append(newOracleRule, newOracleItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "OracleUpdated", previousOracleRule, newOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolOracleUpdated)
				if err := _Pool.contract.UnpackLog(event, "OracleUpdated", log); err != nil {
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

// ParseOracleUpdated is a log parse operation binding the contract event 0x078c3b417dadf69374a59793b829c52001247130433427049317bde56607b1b7.
//
// Solidity: event OracleUpdated(address indexed previousOracle, address indexed newOracle)
func (_Pool *PoolFilterer) ParseOracleUpdated(log types.Log) (*PoolOracleUpdated, error) {
	event := new(PoolOracleUpdated)
	if err := _Pool.contract.UnpackLog(event, "OracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pool contract.
type PoolOwnershipTransferredIterator struct {
	Event *PoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolOwnershipTransferred)
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
		it.Event = new(PoolOwnershipTransferred)
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
func (it *PoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolOwnershipTransferred represents a OwnershipTransferred event raised by the Pool contract.
type PoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pool *PoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PoolOwnershipTransferredIterator{contract: _Pool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pool *PoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolOwnershipTransferred)
				if err := _Pool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pool *PoolFilterer) ParseOwnershipTransferred(log types.Log) (*PoolOwnershipTransferred, error) {
	event := new(PoolOwnershipTransferred)
	if err := _Pool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pool contract.
type PoolPausedIterator struct {
	Event *PoolPaused // Event containing the contract specifics and raw log

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
func (it *PoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolPaused)
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
		it.Event = new(PoolPaused)
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
func (it *PoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolPaused represents a Paused event raised by the Pool contract.
type PoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pool *PoolFilterer) FilterPaused(opts *bind.FilterOpts) (*PoolPausedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PoolPausedIterator{contract: _Pool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pool *PoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PoolPaused) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolPaused)
				if err := _Pool.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Pool *PoolFilterer) ParsePaused(log types.Log) (*PoolPaused, error) {
	event := new(PoolPaused)
	if err := _Pool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolPriceDeviationUpdatedIterator is returned from FilterPriceDeviationUpdated and is used to iterate over the raw logs and unpacked data for PriceDeviationUpdated events raised by the Pool contract.
type PoolPriceDeviationUpdatedIterator struct {
	Event *PoolPriceDeviationUpdated // Event containing the contract specifics and raw log

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
func (it *PoolPriceDeviationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolPriceDeviationUpdated)
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
		it.Event = new(PoolPriceDeviationUpdated)
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
func (it *PoolPriceDeviationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolPriceDeviationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolPriceDeviationUpdated represents a PriceDeviationUpdated event raised by the Pool contract.
type PoolPriceDeviationUpdated struct {
	PreviousPriceDeviation *big.Int
	NewPriceDeviation      *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterPriceDeviationUpdated is a free log retrieval operation binding the contract event 0x193a6231bcc3d7de036c920f05884fe0d8e1dad662793abd049d33b4d83a45cd.
//
// Solidity: event PriceDeviationUpdated(uint256 previousPriceDeviation, uint256 newPriceDeviation)
func (_Pool *PoolFilterer) FilterPriceDeviationUpdated(opts *bind.FilterOpts) (*PoolPriceDeviationUpdatedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "PriceDeviationUpdated")
	if err != nil {
		return nil, err
	}
	return &PoolPriceDeviationUpdatedIterator{contract: _Pool.contract, event: "PriceDeviationUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceDeviationUpdated is a free log subscription operation binding the contract event 0x193a6231bcc3d7de036c920f05884fe0d8e1dad662793abd049d33b4d83a45cd.
//
// Solidity: event PriceDeviationUpdated(uint256 previousPriceDeviation, uint256 newPriceDeviation)
func (_Pool *PoolFilterer) WatchPriceDeviationUpdated(opts *bind.WatchOpts, sink chan<- *PoolPriceDeviationUpdated) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "PriceDeviationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolPriceDeviationUpdated)
				if err := _Pool.contract.UnpackLog(event, "PriceDeviationUpdated", log); err != nil {
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

// ParsePriceDeviationUpdated is a log parse operation binding the contract event 0x193a6231bcc3d7de036c920f05884fe0d8e1dad662793abd049d33b4d83a45cd.
//
// Solidity: event PriceDeviationUpdated(uint256 previousPriceDeviation, uint256 newPriceDeviation)
func (_Pool *PoolFilterer) ParsePriceDeviationUpdated(log types.Log) (*PoolPriceDeviationUpdated, error) {
	event := new(PoolPriceDeviationUpdated)
	if err := _Pool.contract.UnpackLog(event, "PriceDeviationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRetentionRatioUpdatedIterator is returned from FilterRetentionRatioUpdated and is used to iterate over the raw logs and unpacked data for RetentionRatioUpdated events raised by the Pool contract.
type PoolRetentionRatioUpdatedIterator struct {
	Event *PoolRetentionRatioUpdated // Event containing the contract specifics and raw log

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
func (it *PoolRetentionRatioUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRetentionRatioUpdated)
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
		it.Event = new(PoolRetentionRatioUpdated)
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
func (it *PoolRetentionRatioUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRetentionRatioUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRetentionRatioUpdated represents a RetentionRatioUpdated event raised by the Pool contract.
type PoolRetentionRatioUpdated struct {
	PreviousRetentionRatio *big.Int
	NewRetentionRatio      *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRetentionRatioUpdated is a free log retrieval operation binding the contract event 0xbc341359e3658046c114365fc4bf0a6bfe11e53d72faf8ac4fe19b53aec6b3f9.
//
// Solidity: event RetentionRatioUpdated(uint256 previousRetentionRatio, uint256 newRetentionRatio)
func (_Pool *PoolFilterer) FilterRetentionRatioUpdated(opts *bind.FilterOpts) (*PoolRetentionRatioUpdatedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "RetentionRatioUpdated")
	if err != nil {
		return nil, err
	}
	return &PoolRetentionRatioUpdatedIterator{contract: _Pool.contract, event: "RetentionRatioUpdated", logs: logs, sub: sub}, nil
}

// WatchRetentionRatioUpdated is a free log subscription operation binding the contract event 0xbc341359e3658046c114365fc4bf0a6bfe11e53d72faf8ac4fe19b53aec6b3f9.
//
// Solidity: event RetentionRatioUpdated(uint256 previousRetentionRatio, uint256 newRetentionRatio)
func (_Pool *PoolFilterer) WatchRetentionRatioUpdated(opts *bind.WatchOpts, sink chan<- *PoolRetentionRatioUpdated) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "RetentionRatioUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRetentionRatioUpdated)
				if err := _Pool.contract.UnpackLog(event, "RetentionRatioUpdated", log); err != nil {
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

// ParseRetentionRatioUpdated is a log parse operation binding the contract event 0xbc341359e3658046c114365fc4bf0a6bfe11e53d72faf8ac4fe19b53aec6b3f9.
//
// Solidity: event RetentionRatioUpdated(uint256 previousRetentionRatio, uint256 newRetentionRatio)
func (_Pool *PoolFilterer) ParseRetentionRatioUpdated(log types.Log) (*PoolRetentionRatioUpdated, error) {
	event := new(PoolRetentionRatioUpdated)
	if err := _Pool.contract.UnpackLog(event, "RetentionRatioUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolSlippageParamsUpdatedIterator is returned from FilterSlippageParamsUpdated and is used to iterate over the raw logs and unpacked data for SlippageParamsUpdated events raised by the Pool contract.
type PoolSlippageParamsUpdatedIterator struct {
	Event *PoolSlippageParamsUpdated // Event containing the contract specifics and raw log

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
func (it *PoolSlippageParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolSlippageParamsUpdated)
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
		it.Event = new(PoolSlippageParamsUpdated)
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
func (it *PoolSlippageParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolSlippageParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolSlippageParamsUpdated represents a SlippageParamsUpdated event raised by the Pool contract.
type PoolSlippageParamsUpdated struct {
	PreviousK          *big.Int
	NewK               *big.Int
	PreviousN          *big.Int
	NewN               *big.Int
	PreviousC1         *big.Int
	NewC1              *big.Int
	PreviousXThreshold *big.Int
	NewXThreshold      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSlippageParamsUpdated is a free log retrieval operation binding the contract event 0x4f79edc127ae2274ce8320073956fad0b9e727d4ed275b2d3665a75edaffb7a3.
//
// Solidity: event SlippageParamsUpdated(uint256 previousK, uint256 newK, uint256 previousN, uint256 newN, uint256 previousC1, uint256 newC1, uint256 previousXThreshold, uint256 newXThreshold)
func (_Pool *PoolFilterer) FilterSlippageParamsUpdated(opts *bind.FilterOpts) (*PoolSlippageParamsUpdatedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "SlippageParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &PoolSlippageParamsUpdatedIterator{contract: _Pool.contract, event: "SlippageParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchSlippageParamsUpdated is a free log subscription operation binding the contract event 0x4f79edc127ae2274ce8320073956fad0b9e727d4ed275b2d3665a75edaffb7a3.
//
// Solidity: event SlippageParamsUpdated(uint256 previousK, uint256 newK, uint256 previousN, uint256 newN, uint256 previousC1, uint256 newC1, uint256 previousXThreshold, uint256 newXThreshold)
func (_Pool *PoolFilterer) WatchSlippageParamsUpdated(opts *bind.WatchOpts, sink chan<- *PoolSlippageParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "SlippageParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolSlippageParamsUpdated)
				if err := _Pool.contract.UnpackLog(event, "SlippageParamsUpdated", log); err != nil {
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

// ParseSlippageParamsUpdated is a log parse operation binding the contract event 0x4f79edc127ae2274ce8320073956fad0b9e727d4ed275b2d3665a75edaffb7a3.
//
// Solidity: event SlippageParamsUpdated(uint256 previousK, uint256 newK, uint256 previousN, uint256 newN, uint256 previousC1, uint256 newC1, uint256 previousXThreshold, uint256 newXThreshold)
func (_Pool *PoolFilterer) ParseSlippageParamsUpdated(log types.Log) (*PoolSlippageParamsUpdated, error) {
	event := new(PoolSlippageParamsUpdated)
	if err := _Pool.contract.UnpackLog(event, "SlippageParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Pool contract.
type PoolSwapIterator struct {
	Event *PoolSwap // Event containing the contract specifics and raw log

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
func (it *PoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolSwap)
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
		it.Event = new(PoolSwap)
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
func (it *PoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolSwap represents a Swap event raised by the Pool contract.
type PoolSwap struct {
	Sender     common.Address
	FromToken  common.Address
	ToToken    common.Address
	FromAmount *big.Int
	ToAmount   *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x54787c404bb33c88e86f4baf88183a3b0141d0a848e6a9f7a13b66ae3a9b73d1.
//
// Solidity: event Swap(address indexed sender, address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address indexed to)
func (_Pool *PoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolSwapIterator{contract: _Pool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x54787c404bb33c88e86f4baf88183a3b0141d0a848e6a9f7a13b66ae3a9b73d1.
//
// Solidity: event Swap(address indexed sender, address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address indexed to)
func (_Pool *PoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PoolSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolSwap)
				if err := _Pool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x54787c404bb33c88e86f4baf88183a3b0141d0a848e6a9f7a13b66ae3a9b73d1.
//
// Solidity: event Swap(address indexed sender, address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address indexed to)
func (_Pool *PoolFilterer) ParseSwap(log types.Log) (*PoolSwap, error) {
	event := new(PoolSwap)
	if err := _Pool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pool contract.
type PoolUnpausedIterator struct {
	Event *PoolUnpaused // Event containing the contract specifics and raw log

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
func (it *PoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUnpaused)
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
		it.Event = new(PoolUnpaused)
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
func (it *PoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUnpaused represents a Unpaused event raised by the Pool contract.
type PoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pool *PoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PoolUnpausedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PoolUnpausedIterator{contract: _Pool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pool *PoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUnpaused)
				if err := _Pool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Pool *PoolFilterer) ParseUnpaused(log types.Log) (*PoolUnpaused, error) {
	event := new(PoolUnpaused)
	if err := _Pool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Pool contract.
type PoolWithdrawIterator struct {
	Event *PoolWithdraw // Event containing the contract specifics and raw log

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
func (it *PoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolWithdraw)
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
		it.Event = new(PoolWithdraw)
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
func (it *PoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolWithdraw represents a Withdraw event raised by the Pool contract.
type PoolWithdraw struct {
	Sender    common.Address
	Token     common.Address
	Amount    *big.Int
	Liquidity *big.Int
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfb80d861da582b723be2d19507ce3e03851820c464abea89156ec77e089b1ad9.
//
// Solidity: event Withdraw(address indexed sender, address token, uint256 amount, uint256 liquidity, address indexed to)
func (_Pool *PoolFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PoolWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolWithdrawIterator{contract: _Pool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfb80d861da582b723be2d19507ce3e03851820c464abea89156ec77e089b1ad9.
//
// Solidity: event Withdraw(address indexed sender, address token, uint256 amount, uint256 liquidity, address indexed to)
func (_Pool *PoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PoolWithdraw, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolWithdraw)
				if err := _Pool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfb80d861da582b723be2d19507ce3e03851820c464abea89156ec77e089b1ad9.
//
// Solidity: event Withdraw(address indexed sender, address token, uint256 amount, uint256 liquidity, address indexed to)
func (_Pool *PoolFilterer) ParseWithdraw(log types.Log) (*PoolWithdraw, error) {
	event := new(PoolWithdraw)
	if err := _Pool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
