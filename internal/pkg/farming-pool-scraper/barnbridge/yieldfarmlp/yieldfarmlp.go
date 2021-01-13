// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yieldfarmlp

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

// YieldfarmlpABI is the input ABI used to generate the binding from.
const YieldfarmlpABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bondTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniLP\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakeContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"communityVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochsHarvested\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"}],\"name\":\"MassHarvest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NR_OF_EPOCHS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_DISTRIBUTED_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getEpochStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getPoolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"harvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInitializedEpoch\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massHarvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userLastEpochIdHarvested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Yieldfarmlp is an auto generated Go binding around an Ethereum contract.
type Yieldfarmlp struct {
	YieldfarmlpCaller     // Read-only binding to the contract
	YieldfarmlpTransactor // Write-only binding to the contract
	YieldfarmlpFilterer   // Log filterer for contract events
}

// YieldfarmlpCaller is an auto generated read-only Go binding around an Ethereum contract.
type YieldfarmlpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldfarmlpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YieldfarmlpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldfarmlpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YieldfarmlpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldfarmlpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YieldfarmlpSession struct {
	Contract     *Yieldfarmlp      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YieldfarmlpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YieldfarmlpCallerSession struct {
	Contract *YieldfarmlpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// YieldfarmlpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YieldfarmlpTransactorSession struct {
	Contract     *YieldfarmlpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// YieldfarmlpRaw is an auto generated low-level Go binding around an Ethereum contract.
type YieldfarmlpRaw struct {
	Contract *Yieldfarmlp // Generic contract binding to access the raw methods on
}

// YieldfarmlpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YieldfarmlpCallerRaw struct {
	Contract *YieldfarmlpCaller // Generic read-only contract binding to access the raw methods on
}

// YieldfarmlpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YieldfarmlpTransactorRaw struct {
	Contract *YieldfarmlpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYieldfarmlp creates a new instance of Yieldfarmlp, bound to a specific deployed contract.
func NewYieldfarmlp(address common.Address, backend bind.ContractBackend) (*Yieldfarmlp, error) {
	contract, err := bindYieldfarmlp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yieldfarmlp{YieldfarmlpCaller: YieldfarmlpCaller{contract: contract}, YieldfarmlpTransactor: YieldfarmlpTransactor{contract: contract}, YieldfarmlpFilterer: YieldfarmlpFilterer{contract: contract}}, nil
}

// NewYieldfarmlpCaller creates a new read-only instance of Yieldfarmlp, bound to a specific deployed contract.
func NewYieldfarmlpCaller(address common.Address, caller bind.ContractCaller) (*YieldfarmlpCaller, error) {
	contract, err := bindYieldfarmlp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YieldfarmlpCaller{contract: contract}, nil
}

// NewYieldfarmlpTransactor creates a new write-only instance of Yieldfarmlp, bound to a specific deployed contract.
func NewYieldfarmlpTransactor(address common.Address, transactor bind.ContractTransactor) (*YieldfarmlpTransactor, error) {
	contract, err := bindYieldfarmlp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YieldfarmlpTransactor{contract: contract}, nil
}

// NewYieldfarmlpFilterer creates a new log filterer instance of Yieldfarmlp, bound to a specific deployed contract.
func NewYieldfarmlpFilterer(address common.Address, filterer bind.ContractFilterer) (*YieldfarmlpFilterer, error) {
	contract, err := bindYieldfarmlp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YieldfarmlpFilterer{contract: contract}, nil
}

// bindYieldfarmlp binds a generic wrapper to an already deployed contract.
func bindYieldfarmlp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YieldfarmlpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yieldfarmlp *YieldfarmlpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yieldfarmlp.Contract.YieldfarmlpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yieldfarmlp *YieldfarmlpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yieldfarmlp.Contract.YieldfarmlpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yieldfarmlp *YieldfarmlpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yieldfarmlp.Contract.YieldfarmlpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yieldfarmlp *YieldfarmlpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yieldfarmlp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yieldfarmlp *YieldfarmlpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yieldfarmlp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yieldfarmlp *YieldfarmlpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yieldfarmlp.Contract.contract.Transact(opts, method, params...)
}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCaller) NROFEPOCHS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmlp.contract.Call(opts, &out, "NR_OF_EPOCHS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpSession) NROFEPOCHS() (*big.Int, error) {
	return _Yieldfarmlp.Contract.NROFEPOCHS(&_Yieldfarmlp.CallOpts)
}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCallerSession) NROFEPOCHS() (*big.Int, error) {
	return _Yieldfarmlp.Contract.NROFEPOCHS(&_Yieldfarmlp.CallOpts)
}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCaller) TOTALDISTRIBUTEDAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmlp.contract.Call(opts, &out, "TOTAL_DISTRIBUTED_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpSession) TOTALDISTRIBUTEDAMOUNT() (*big.Int, error) {
	return _Yieldfarmlp.Contract.TOTALDISTRIBUTEDAMOUNT(&_Yieldfarmlp.CallOpts)
}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCallerSession) TOTALDISTRIBUTEDAMOUNT() (*big.Int, error) {
	return _Yieldfarmlp.Contract.TOTALDISTRIBUTEDAMOUNT(&_Yieldfarmlp.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCaller) EpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmlp.contract.Call(opts, &out, "epochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpSession) EpochDuration() (*big.Int, error) {
	return _Yieldfarmlp.Contract.EpochDuration(&_Yieldfarmlp.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCallerSession) EpochDuration() (*big.Int, error) {
	return _Yieldfarmlp.Contract.EpochDuration(&_Yieldfarmlp.CallOpts)
}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCaller) EpochStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmlp.contract.Call(opts, &out, "epochStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpSession) EpochStart() (*big.Int, error) {
	return _Yieldfarmlp.Contract.EpochStart(&_Yieldfarmlp.CallOpts)
}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCallerSession) EpochStart() (*big.Int, error) {
	return _Yieldfarmlp.Contract.EpochStart(&_Yieldfarmlp.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmlp.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpSession) GetCurrentEpoch() (*big.Int, error) {
	return _Yieldfarmlp.Contract.GetCurrentEpoch(&_Yieldfarmlp.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _Yieldfarmlp.Contract.GetCurrentEpoch(&_Yieldfarmlp.CallOpts)
}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCaller) GetEpochStake(opts *bind.CallOpts, userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmlp.contract.Call(opts, &out, "getEpochStake", userAddress, epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpSession) GetEpochStake(userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _Yieldfarmlp.Contract.GetEpochStake(&_Yieldfarmlp.CallOpts, userAddress, epochId)
}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCallerSession) GetEpochStake(userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _Yieldfarmlp.Contract.GetEpochStake(&_Yieldfarmlp.CallOpts, userAddress, epochId)
}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCaller) GetPoolSize(opts *bind.CallOpts, epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmlp.contract.Call(opts, &out, "getPoolSize", epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpSession) GetPoolSize(epochId *big.Int) (*big.Int, error) {
	return _Yieldfarmlp.Contract.GetPoolSize(&_Yieldfarmlp.CallOpts, epochId)
}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCallerSession) GetPoolSize(epochId *big.Int) (*big.Int, error) {
	return _Yieldfarmlp.Contract.GetPoolSize(&_Yieldfarmlp.CallOpts, epochId)
}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_Yieldfarmlp *YieldfarmlpCaller) LastInitializedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmlp.contract.Call(opts, &out, "lastInitializedEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_Yieldfarmlp *YieldfarmlpSession) LastInitializedEpoch() (*big.Int, error) {
	return _Yieldfarmlp.Contract.LastInitializedEpoch(&_Yieldfarmlp.CallOpts)
}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_Yieldfarmlp *YieldfarmlpCallerSession) LastInitializedEpoch() (*big.Int, error) {
	return _Yieldfarmlp.Contract.LastInitializedEpoch(&_Yieldfarmlp.CallOpts)
}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCaller) UserLastEpochIdHarvested(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmlp.contract.Call(opts, &out, "userLastEpochIdHarvested")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpSession) UserLastEpochIdHarvested() (*big.Int, error) {
	return _Yieldfarmlp.Contract.UserLastEpochIdHarvested(&_Yieldfarmlp.CallOpts)
}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_Yieldfarmlp *YieldfarmlpCallerSession) UserLastEpochIdHarvested() (*big.Int, error) {
	return _Yieldfarmlp.Contract.UserLastEpochIdHarvested(&_Yieldfarmlp.CallOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_Yieldfarmlp *YieldfarmlpTransactor) Harvest(opts *bind.TransactOpts, epochId *big.Int) (*types.Transaction, error) {
	return _Yieldfarmlp.contract.Transact(opts, "harvest", epochId)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_Yieldfarmlp *YieldfarmlpSession) Harvest(epochId *big.Int) (*types.Transaction, error) {
	return _Yieldfarmlp.Contract.Harvest(&_Yieldfarmlp.TransactOpts, epochId)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_Yieldfarmlp *YieldfarmlpTransactorSession) Harvest(epochId *big.Int) (*types.Transaction, error) {
	return _Yieldfarmlp.Contract.Harvest(&_Yieldfarmlp.TransactOpts, epochId)
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_Yieldfarmlp *YieldfarmlpTransactor) MassHarvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yieldfarmlp.contract.Transact(opts, "massHarvest")
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_Yieldfarmlp *YieldfarmlpSession) MassHarvest() (*types.Transaction, error) {
	return _Yieldfarmlp.Contract.MassHarvest(&_Yieldfarmlp.TransactOpts)
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_Yieldfarmlp *YieldfarmlpTransactorSession) MassHarvest() (*types.Transaction, error) {
	return _Yieldfarmlp.Contract.MassHarvest(&_Yieldfarmlp.TransactOpts)
}

// YieldfarmlpHarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the Yieldfarmlp contract.
type YieldfarmlpHarvestIterator struct {
	Event *YieldfarmlpHarvest // Event containing the contract specifics and raw log

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
func (it *YieldfarmlpHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YieldfarmlpHarvest)
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
		it.Event = new(YieldfarmlpHarvest)
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
func (it *YieldfarmlpHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YieldfarmlpHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YieldfarmlpHarvest represents a Harvest event raised by the Yieldfarmlp contract.
type YieldfarmlpHarvest struct {
	User    common.Address
	EpochId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0x04ad45a69eeed9c390c3a678fed2d4b90bde98e742de9936d5e0915bf3d0ea4e.
//
// Solidity: event Harvest(address indexed user, uint128 indexed epochId, uint256 amount)
func (_Yieldfarmlp *YieldfarmlpFilterer) FilterHarvest(opts *bind.FilterOpts, user []common.Address, epochId []*big.Int) (*YieldfarmlpHarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Yieldfarmlp.contract.FilterLogs(opts, "Harvest", userRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &YieldfarmlpHarvestIterator{contract: _Yieldfarmlp.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0x04ad45a69eeed9c390c3a678fed2d4b90bde98e742de9936d5e0915bf3d0ea4e.
//
// Solidity: event Harvest(address indexed user, uint128 indexed epochId, uint256 amount)
func (_Yieldfarmlp *YieldfarmlpFilterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *YieldfarmlpHarvest, user []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Yieldfarmlp.contract.WatchLogs(opts, "Harvest", userRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YieldfarmlpHarvest)
				if err := _Yieldfarmlp.contract.UnpackLog(event, "Harvest", log); err != nil {
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

// ParseHarvest is a log parse operation binding the contract event 0x04ad45a69eeed9c390c3a678fed2d4b90bde98e742de9936d5e0915bf3d0ea4e.
//
// Solidity: event Harvest(address indexed user, uint128 indexed epochId, uint256 amount)
func (_Yieldfarmlp *YieldfarmlpFilterer) ParseHarvest(log types.Log) (*YieldfarmlpHarvest, error) {
	event := new(YieldfarmlpHarvest)
	if err := _Yieldfarmlp.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// YieldfarmlpMassHarvestIterator is returned from FilterMassHarvest and is used to iterate over the raw logs and unpacked data for MassHarvest events raised by the Yieldfarmlp contract.
type YieldfarmlpMassHarvestIterator struct {
	Event *YieldfarmlpMassHarvest // Event containing the contract specifics and raw log

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
func (it *YieldfarmlpMassHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YieldfarmlpMassHarvest)
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
		it.Event = new(YieldfarmlpMassHarvest)
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
func (it *YieldfarmlpMassHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YieldfarmlpMassHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YieldfarmlpMassHarvest represents a MassHarvest event raised by the Yieldfarmlp contract.
type YieldfarmlpMassHarvest struct {
	User            common.Address
	EpochsHarvested *big.Int
	TotalValue      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMassHarvest is a free log retrieval operation binding the contract event 0xb68dafc1da13dc868096d0b87347c831d0bda92d178317eb1dec7f788444485c.
//
// Solidity: event MassHarvest(address indexed user, uint256 epochsHarvested, uint256 totalValue)
func (_Yieldfarmlp *YieldfarmlpFilterer) FilterMassHarvest(opts *bind.FilterOpts, user []common.Address) (*YieldfarmlpMassHarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Yieldfarmlp.contract.FilterLogs(opts, "MassHarvest", userRule)
	if err != nil {
		return nil, err
	}
	return &YieldfarmlpMassHarvestIterator{contract: _Yieldfarmlp.contract, event: "MassHarvest", logs: logs, sub: sub}, nil
}

// WatchMassHarvest is a free log subscription operation binding the contract event 0xb68dafc1da13dc868096d0b87347c831d0bda92d178317eb1dec7f788444485c.
//
// Solidity: event MassHarvest(address indexed user, uint256 epochsHarvested, uint256 totalValue)
func (_Yieldfarmlp *YieldfarmlpFilterer) WatchMassHarvest(opts *bind.WatchOpts, sink chan<- *YieldfarmlpMassHarvest, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Yieldfarmlp.contract.WatchLogs(opts, "MassHarvest", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YieldfarmlpMassHarvest)
				if err := _Yieldfarmlp.contract.UnpackLog(event, "MassHarvest", log); err != nil {
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

// ParseMassHarvest is a log parse operation binding the contract event 0xb68dafc1da13dc868096d0b87347c831d0bda92d178317eb1dec7f788444485c.
//
// Solidity: event MassHarvest(address indexed user, uint256 epochsHarvested, uint256 totalValue)
func (_Yieldfarmlp *YieldfarmlpFilterer) ParseMassHarvest(log types.Log) (*YieldfarmlpMassHarvest, error) {
	event := new(YieldfarmlpMassHarvest)
	if err := _Yieldfarmlp.contract.UnpackLog(event, "MassHarvest", log); err != nil {
		return nil, err
	}
	return event, nil
}
