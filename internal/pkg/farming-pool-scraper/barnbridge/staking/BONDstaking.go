// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// BONDStakingABI is the input ABI used to generate the binding from.
const BONDStakingABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch1Start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epochDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"ManualEpochInit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prevBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"prevMultiplier\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"currentMultiplier\",\"type\":\"uint128\"}],\"name\":\"computeNewMultiplier\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpochMultiplier\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch1Start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"epochIsInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getEpochPoolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getEpochUserBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"manualEpochInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BONDStaking is an auto generated Go binding around an Ethereum contract.
type BONDStaking struct {
	BONDStakingCaller     // Read-only binding to the contract
	BONDStakingTransactor // Write-only binding to the contract
	BONDStakingFilterer   // Log filterer for contract events
}

// BONDStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BONDStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BONDStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BONDStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BONDStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BONDStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BONDStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BONDStakingSession struct {
	Contract     *BONDStaking      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BONDStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BONDStakingCallerSession struct {
	Contract *BONDStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BONDStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BONDStakingTransactorSession struct {
	Contract     *BONDStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BONDStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BONDStakingRaw struct {
	Contract *BONDStaking // Generic contract binding to access the raw methods on
}

// BONDStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BONDStakingCallerRaw struct {
	Contract *BONDStakingCaller // Generic read-only contract binding to access the raw methods on
}

// BONDStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BONDStakingTransactorRaw struct {
	Contract *BONDStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBONDStaking creates a new instance of BONDStaking, bound to a specific deployed contract.
func NewBONDStaking(address common.Address, backend bind.ContractBackend) (*BONDStaking, error) {
	contract, err := bindBONDStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BONDStaking{BONDStakingCaller: BONDStakingCaller{contract: contract}, BONDStakingTransactor: BONDStakingTransactor{contract: contract}, BONDStakingFilterer: BONDStakingFilterer{contract: contract}}, nil
}

// NewBONDStakingCaller creates a new read-only instance of BONDStaking, bound to a specific deployed contract.
func NewBONDStakingCaller(address common.Address, caller bind.ContractCaller) (*BONDStakingCaller, error) {
	contract, err := bindBONDStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BONDStakingCaller{contract: contract}, nil
}

// NewBONDStakingTransactor creates a new write-only instance of BONDStaking, bound to a specific deployed contract.
func NewBONDStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*BONDStakingTransactor, error) {
	contract, err := bindBONDStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BONDStakingTransactor{contract: contract}, nil
}

// NewBONDStakingFilterer creates a new log filterer instance of BONDStaking, bound to a specific deployed contract.
func NewBONDStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*BONDStakingFilterer, error) {
	contract, err := bindBONDStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BONDStakingFilterer{contract: contract}, nil
}

// bindBONDStaking binds a generic wrapper to an already deployed contract.
func bindBONDStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BONDStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BONDStaking *BONDStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BONDStaking.Contract.BONDStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BONDStaking *BONDStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BONDStaking.Contract.BONDStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BONDStaking *BONDStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BONDStaking.Contract.BONDStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BONDStaking *BONDStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BONDStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BONDStaking *BONDStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BONDStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BONDStaking *BONDStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BONDStaking.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address user, address token) view returns(uint256)
func (_BONDStaking *BONDStakingCaller) BalanceOf(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BONDStaking.contract.Call(opts, &out, "balanceOf", user, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address user, address token) view returns(uint256)
func (_BONDStaking *BONDStakingSession) BalanceOf(user common.Address, token common.Address) (*big.Int, error) {
	return _BONDStaking.Contract.BalanceOf(&_BONDStaking.CallOpts, user, token)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address user, address token) view returns(uint256)
func (_BONDStaking *BONDStakingCallerSession) BalanceOf(user common.Address, token common.Address) (*big.Int, error) {
	return _BONDStaking.Contract.BalanceOf(&_BONDStaking.CallOpts, user, token)
}

// ComputeNewMultiplier is a free data retrieval call binding the contract method 0x4be41dba.
//
// Solidity: function computeNewMultiplier(uint256 prevBalance, uint128 prevMultiplier, uint256 amount, uint128 currentMultiplier) pure returns(uint128)
func (_BONDStaking *BONDStakingCaller) ComputeNewMultiplier(opts *bind.CallOpts, prevBalance *big.Int, prevMultiplier *big.Int, amount *big.Int, currentMultiplier *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BONDStaking.contract.Call(opts, &out, "computeNewMultiplier", prevBalance, prevMultiplier, amount, currentMultiplier)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeNewMultiplier is a free data retrieval call binding the contract method 0x4be41dba.
//
// Solidity: function computeNewMultiplier(uint256 prevBalance, uint128 prevMultiplier, uint256 amount, uint128 currentMultiplier) pure returns(uint128)
func (_BONDStaking *BONDStakingSession) ComputeNewMultiplier(prevBalance *big.Int, prevMultiplier *big.Int, amount *big.Int, currentMultiplier *big.Int) (*big.Int, error) {
	return _BONDStaking.Contract.ComputeNewMultiplier(&_BONDStaking.CallOpts, prevBalance, prevMultiplier, amount, currentMultiplier)
}

// ComputeNewMultiplier is a free data retrieval call binding the contract method 0x4be41dba.
//
// Solidity: function computeNewMultiplier(uint256 prevBalance, uint128 prevMultiplier, uint256 amount, uint128 currentMultiplier) pure returns(uint128)
func (_BONDStaking *BONDStakingCallerSession) ComputeNewMultiplier(prevBalance *big.Int, prevMultiplier *big.Int, amount *big.Int, currentMultiplier *big.Int) (*big.Int, error) {
	return _BONDStaking.Contract.ComputeNewMultiplier(&_BONDStaking.CallOpts, prevBalance, prevMultiplier, amount, currentMultiplier)
}

// CurrentEpochMultiplier is a free data retrieval call binding the contract method 0xce58a2a8.
//
// Solidity: function currentEpochMultiplier() view returns(uint128)
func (_BONDStaking *BONDStakingCaller) CurrentEpochMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BONDStaking.contract.Call(opts, &out, "currentEpochMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpochMultiplier is a free data retrieval call binding the contract method 0xce58a2a8.
//
// Solidity: function currentEpochMultiplier() view returns(uint128)
func (_BONDStaking *BONDStakingSession) CurrentEpochMultiplier() (*big.Int, error) {
	return _BONDStaking.Contract.CurrentEpochMultiplier(&_BONDStaking.CallOpts)
}

// CurrentEpochMultiplier is a free data retrieval call binding the contract method 0xce58a2a8.
//
// Solidity: function currentEpochMultiplier() view returns(uint128)
func (_BONDStaking *BONDStakingCallerSession) CurrentEpochMultiplier() (*big.Int, error) {
	return _BONDStaking.Contract.CurrentEpochMultiplier(&_BONDStaking.CallOpts)
}

// Epoch1Start is a free data retrieval call binding the contract method 0xf4a4341d.
//
// Solidity: function epoch1Start() view returns(uint256)
func (_BONDStaking *BONDStakingCaller) Epoch1Start(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BONDStaking.contract.Call(opts, &out, "epoch1Start")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch1Start is a free data retrieval call binding the contract method 0xf4a4341d.
//
// Solidity: function epoch1Start() view returns(uint256)
func (_BONDStaking *BONDStakingSession) Epoch1Start() (*big.Int, error) {
	return _BONDStaking.Contract.Epoch1Start(&_BONDStaking.CallOpts)
}

// Epoch1Start is a free data retrieval call binding the contract method 0xf4a4341d.
//
// Solidity: function epoch1Start() view returns(uint256)
func (_BONDStaking *BONDStakingCallerSession) Epoch1Start() (*big.Int, error) {
	return _BONDStaking.Contract.Epoch1Start(&_BONDStaking.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_BONDStaking *BONDStakingCaller) EpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BONDStaking.contract.Call(opts, &out, "epochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_BONDStaking *BONDStakingSession) EpochDuration() (*big.Int, error) {
	return _BONDStaking.Contract.EpochDuration(&_BONDStaking.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_BONDStaking *BONDStakingCallerSession) EpochDuration() (*big.Int, error) {
	return _BONDStaking.Contract.EpochDuration(&_BONDStaking.CallOpts)
}

// EpochIsInitialized is a free data retrieval call binding the contract method 0xaa579154.
//
// Solidity: function epochIsInitialized(address token, uint128 epochId) view returns(bool)
func (_BONDStaking *BONDStakingCaller) EpochIsInitialized(opts *bind.CallOpts, token common.Address, epochId *big.Int) (bool, error) {
	var out []interface{}
	err := _BONDStaking.contract.Call(opts, &out, "epochIsInitialized", token, epochId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EpochIsInitialized is a free data retrieval call binding the contract method 0xaa579154.
//
// Solidity: function epochIsInitialized(address token, uint128 epochId) view returns(bool)
func (_BONDStaking *BONDStakingSession) EpochIsInitialized(token common.Address, epochId *big.Int) (bool, error) {
	return _BONDStaking.Contract.EpochIsInitialized(&_BONDStaking.CallOpts, token, epochId)
}

// EpochIsInitialized is a free data retrieval call binding the contract method 0xaa579154.
//
// Solidity: function epochIsInitialized(address token, uint128 epochId) view returns(bool)
func (_BONDStaking *BONDStakingCallerSession) EpochIsInitialized(token common.Address, epochId *big.Int) (bool, error) {
	return _BONDStaking.Contract.EpochIsInitialized(&_BONDStaking.CallOpts, token, epochId)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint128)
func (_BONDStaking *BONDStakingCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BONDStaking.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint128)
func (_BONDStaking *BONDStakingSession) GetCurrentEpoch() (*big.Int, error) {
	return _BONDStaking.Contract.GetCurrentEpoch(&_BONDStaking.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint128)
func (_BONDStaking *BONDStakingCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _BONDStaking.Contract.GetCurrentEpoch(&_BONDStaking.CallOpts)
}

// GetEpochPoolSize is a free data retrieval call binding the contract method 0x2ca32d7e.
//
// Solidity: function getEpochPoolSize(address tokenAddress, uint128 epochId) view returns(uint256)
func (_BONDStaking *BONDStakingCaller) GetEpochPoolSize(opts *bind.CallOpts, tokenAddress common.Address, epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BONDStaking.contract.Call(opts, &out, "getEpochPoolSize", tokenAddress, epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochPoolSize is a free data retrieval call binding the contract method 0x2ca32d7e.
//
// Solidity: function getEpochPoolSize(address tokenAddress, uint128 epochId) view returns(uint256)
func (_BONDStaking *BONDStakingSession) GetEpochPoolSize(tokenAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _BONDStaking.Contract.GetEpochPoolSize(&_BONDStaking.CallOpts, tokenAddress, epochId)
}

// GetEpochPoolSize is a free data retrieval call binding the contract method 0x2ca32d7e.
//
// Solidity: function getEpochPoolSize(address tokenAddress, uint128 epochId) view returns(uint256)
func (_BONDStaking *BONDStakingCallerSession) GetEpochPoolSize(tokenAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _BONDStaking.Contract.GetEpochPoolSize(&_BONDStaking.CallOpts, tokenAddress, epochId)
}

// GetEpochUserBalance is a free data retrieval call binding the contract method 0x8c028dd0.
//
// Solidity: function getEpochUserBalance(address user, address token, uint128 epochId) view returns(uint256)
func (_BONDStaking *BONDStakingCaller) GetEpochUserBalance(opts *bind.CallOpts, user common.Address, token common.Address, epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BONDStaking.contract.Call(opts, &out, "getEpochUserBalance", user, token, epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochUserBalance is a free data retrieval call binding the contract method 0x8c028dd0.
//
// Solidity: function getEpochUserBalance(address user, address token, uint128 epochId) view returns(uint256)
func (_BONDStaking *BONDStakingSession) GetEpochUserBalance(user common.Address, token common.Address, epochId *big.Int) (*big.Int, error) {
	return _BONDStaking.Contract.GetEpochUserBalance(&_BONDStaking.CallOpts, user, token, epochId)
}

// GetEpochUserBalance is a free data retrieval call binding the contract method 0x8c028dd0.
//
// Solidity: function getEpochUserBalance(address user, address token, uint128 epochId) view returns(uint256)
func (_BONDStaking *BONDStakingCallerSession) GetEpochUserBalance(user common.Address, token common.Address, epochId *big.Int) (*big.Int, error) {
	return _BONDStaking.Contract.GetEpochUserBalance(&_BONDStaking.CallOpts, user, token, epochId)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address tokenAddress, uint256 amount) returns()
func (_BONDStaking *BONDStakingTransactor) Deposit(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BONDStaking.contract.Transact(opts, "deposit", tokenAddress, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address tokenAddress, uint256 amount) returns()
func (_BONDStaking *BONDStakingSession) Deposit(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BONDStaking.Contract.Deposit(&_BONDStaking.TransactOpts, tokenAddress, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address tokenAddress, uint256 amount) returns()
func (_BONDStaking *BONDStakingTransactorSession) Deposit(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BONDStaking.Contract.Deposit(&_BONDStaking.TransactOpts, tokenAddress, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address tokenAddress) returns()
func (_BONDStaking *BONDStakingTransactor) EmergencyWithdraw(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _BONDStaking.contract.Transact(opts, "emergencyWithdraw", tokenAddress)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address tokenAddress) returns()
func (_BONDStaking *BONDStakingSession) EmergencyWithdraw(tokenAddress common.Address) (*types.Transaction, error) {
	return _BONDStaking.Contract.EmergencyWithdraw(&_BONDStaking.TransactOpts, tokenAddress)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address tokenAddress) returns()
func (_BONDStaking *BONDStakingTransactorSession) EmergencyWithdraw(tokenAddress common.Address) (*types.Transaction, error) {
	return _BONDStaking.Contract.EmergencyWithdraw(&_BONDStaking.TransactOpts, tokenAddress)
}

// ManualEpochInit is a paid mutator transaction binding the contract method 0xea2c38ae.
//
// Solidity: function manualEpochInit(address[] tokens, uint128 epochId) returns()
func (_BONDStaking *BONDStakingTransactor) ManualEpochInit(opts *bind.TransactOpts, tokens []common.Address, epochId *big.Int) (*types.Transaction, error) {
	return _BONDStaking.contract.Transact(opts, "manualEpochInit", tokens, epochId)
}

// ManualEpochInit is a paid mutator transaction binding the contract method 0xea2c38ae.
//
// Solidity: function manualEpochInit(address[] tokens, uint128 epochId) returns()
func (_BONDStaking *BONDStakingSession) ManualEpochInit(tokens []common.Address, epochId *big.Int) (*types.Transaction, error) {
	return _BONDStaking.Contract.ManualEpochInit(&_BONDStaking.TransactOpts, tokens, epochId)
}

// ManualEpochInit is a paid mutator transaction binding the contract method 0xea2c38ae.
//
// Solidity: function manualEpochInit(address[] tokens, uint128 epochId) returns()
func (_BONDStaking *BONDStakingTransactorSession) ManualEpochInit(tokens []common.Address, epochId *big.Int) (*types.Transaction, error) {
	return _BONDStaking.Contract.ManualEpochInit(&_BONDStaking.TransactOpts, tokens, epochId)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address tokenAddress, uint256 amount) returns()
func (_BONDStaking *BONDStakingTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BONDStaking.contract.Transact(opts, "withdraw", tokenAddress, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address tokenAddress, uint256 amount) returns()
func (_BONDStaking *BONDStakingSession) Withdraw(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BONDStaking.Contract.Withdraw(&_BONDStaking.TransactOpts, tokenAddress, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address tokenAddress, uint256 amount) returns()
func (_BONDStaking *BONDStakingTransactorSession) Withdraw(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BONDStaking.Contract.Withdraw(&_BONDStaking.TransactOpts, tokenAddress, amount)
}

// BONDStakingDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BONDStaking contract.
type BONDStakingDepositIterator struct {
	Event *BONDStakingDeposit // Event containing the contract specifics and raw log

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
func (it *BONDStakingDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BONDStakingDeposit)
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
		it.Event = new(BONDStakingDeposit)
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
func (it *BONDStakingDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BONDStakingDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BONDStakingDeposit represents a Deposit event raised by the BONDStaking contract.
type BONDStakingDeposit struct {
	User         common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed user, address indexed tokenAddress, uint256 amount)
func (_BONDStaking *BONDStakingFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, tokenAddress []common.Address) (*BONDStakingDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _BONDStaking.contract.FilterLogs(opts, "Deposit", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BONDStakingDepositIterator{contract: _BONDStaking.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed user, address indexed tokenAddress, uint256 amount)
func (_BONDStaking *BONDStakingFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BONDStakingDeposit, user []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _BONDStaking.contract.WatchLogs(opts, "Deposit", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BONDStakingDeposit)
				if err := _BONDStaking.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed user, address indexed tokenAddress, uint256 amount)
func (_BONDStaking *BONDStakingFilterer) ParseDeposit(log types.Log) (*BONDStakingDeposit, error) {
	event := new(BONDStakingDeposit)
	if err := _BONDStaking.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BONDStakingEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the BONDStaking contract.
type BONDStakingEmergencyWithdrawIterator struct {
	Event *BONDStakingEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *BONDStakingEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BONDStakingEmergencyWithdraw)
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
		it.Event = new(BONDStakingEmergencyWithdraw)
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
func (it *BONDStakingEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BONDStakingEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BONDStakingEmergencyWithdraw represents a EmergencyWithdraw event raised by the BONDStaking contract.
type BONDStakingEmergencyWithdraw struct {
	User         common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xf24ef89f38eadc1bde50701ad6e4d6d11a2dc24f7cf834a486991f3883328504.
//
// Solidity: event EmergencyWithdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_BONDStaking *BONDStakingFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, tokenAddress []common.Address) (*BONDStakingEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _BONDStaking.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BONDStakingEmergencyWithdrawIterator{contract: _BONDStaking.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xf24ef89f38eadc1bde50701ad6e4d6d11a2dc24f7cf834a486991f3883328504.
//
// Solidity: event EmergencyWithdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_BONDStaking *BONDStakingFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *BONDStakingEmergencyWithdraw, user []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _BONDStaking.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BONDStakingEmergencyWithdraw)
				if err := _BONDStaking.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xf24ef89f38eadc1bde50701ad6e4d6d11a2dc24f7cf834a486991f3883328504.
//
// Solidity: event EmergencyWithdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_BONDStaking *BONDStakingFilterer) ParseEmergencyWithdraw(log types.Log) (*BONDStakingEmergencyWithdraw, error) {
	event := new(BONDStakingEmergencyWithdraw)
	if err := _BONDStaking.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BONDStakingManualEpochInitIterator is returned from FilterManualEpochInit and is used to iterate over the raw logs and unpacked data for ManualEpochInit events raised by the BONDStaking contract.
type BONDStakingManualEpochInitIterator struct {
	Event *BONDStakingManualEpochInit // Event containing the contract specifics and raw log

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
func (it *BONDStakingManualEpochInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BONDStakingManualEpochInit)
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
		it.Event = new(BONDStakingManualEpochInit)
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
func (it *BONDStakingManualEpochInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BONDStakingManualEpochInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BONDStakingManualEpochInit represents a ManualEpochInit event raised by the BONDStaking contract.
type BONDStakingManualEpochInit struct {
	Caller  common.Address
	EpochId *big.Int
	Tokens  []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManualEpochInit is a free log retrieval operation binding the contract event 0xb85c32b8d9cecc81feba78646289584a693e9a8afea40ab2fd31efae4408429f.
//
// Solidity: event ManualEpochInit(address indexed caller, uint128 indexed epochId, address[] tokens)
func (_BONDStaking *BONDStakingFilterer) FilterManualEpochInit(opts *bind.FilterOpts, caller []common.Address, epochId []*big.Int) (*BONDStakingManualEpochInitIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _BONDStaking.contract.FilterLogs(opts, "ManualEpochInit", callerRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &BONDStakingManualEpochInitIterator{contract: _BONDStaking.contract, event: "ManualEpochInit", logs: logs, sub: sub}, nil
}

// WatchManualEpochInit is a free log subscription operation binding the contract event 0xb85c32b8d9cecc81feba78646289584a693e9a8afea40ab2fd31efae4408429f.
//
// Solidity: event ManualEpochInit(address indexed caller, uint128 indexed epochId, address[] tokens)
func (_BONDStaking *BONDStakingFilterer) WatchManualEpochInit(opts *bind.WatchOpts, sink chan<- *BONDStakingManualEpochInit, caller []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _BONDStaking.contract.WatchLogs(opts, "ManualEpochInit", callerRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BONDStakingManualEpochInit)
				if err := _BONDStaking.contract.UnpackLog(event, "ManualEpochInit", log); err != nil {
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

// ParseManualEpochInit is a log parse operation binding the contract event 0xb85c32b8d9cecc81feba78646289584a693e9a8afea40ab2fd31efae4408429f.
//
// Solidity: event ManualEpochInit(address indexed caller, uint128 indexed epochId, address[] tokens)
func (_BONDStaking *BONDStakingFilterer) ParseManualEpochInit(log types.Log) (*BONDStakingManualEpochInit, error) {
	event := new(BONDStakingManualEpochInit)
	if err := _BONDStaking.contract.UnpackLog(event, "ManualEpochInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BONDStakingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BONDStaking contract.
type BONDStakingWithdrawIterator struct {
	Event *BONDStakingWithdraw // Event containing the contract specifics and raw log

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
func (it *BONDStakingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BONDStakingWithdraw)
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
		it.Event = new(BONDStakingWithdraw)
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
func (it *BONDStakingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BONDStakingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BONDStakingWithdraw represents a Withdraw event raised by the BONDStaking contract.
type BONDStakingWithdraw struct {
	User         common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_BONDStaking *BONDStakingFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, tokenAddress []common.Address) (*BONDStakingWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _BONDStaking.contract.FilterLogs(opts, "Withdraw", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BONDStakingWithdrawIterator{contract: _BONDStaking.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_BONDStaking *BONDStakingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BONDStakingWithdraw, user []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _BONDStaking.contract.WatchLogs(opts, "Withdraw", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BONDStakingWithdraw)
				if err := _BONDStaking.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_BONDStaking *BONDStakingFilterer) ParseWithdraw(log types.Log) (*BONDStakingWithdraw, error) {
	event := new(BONDStakingWithdraw)
	if err := _BONDStaking.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
