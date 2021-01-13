// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yieldfarmbond

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

// YieldfarmbondABI is the input ABI used to generate the binding from.
const YieldfarmbondABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bondTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakeContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"communityVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochsHarvested\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"}],\"name\":\"MassHarvest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EPOCHS_DELAYED_FROM_STAKING_CONTRACT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NR_OF_EPOCHS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_DISTRIBUTED_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getEpochStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getPoolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"harvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInitializedEpoch\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massHarvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userLastEpochIdHarvested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Yieldfarmbond is an auto generated Go binding around an Ethereum contract.
type Yieldfarmbond struct {
	YieldfarmbondCaller     // Read-only binding to the contract
	YieldfarmbondTransactor // Write-only binding to the contract
	YieldfarmbondFilterer   // Log filterer for contract events
}

// YieldfarmbondCaller is an auto generated read-only Go binding around an Ethereum contract.
type YieldfarmbondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldfarmbondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YieldfarmbondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldfarmbondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YieldfarmbondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldfarmbondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YieldfarmbondSession struct {
	Contract     *Yieldfarmbond    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YieldfarmbondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YieldfarmbondCallerSession struct {
	Contract *YieldfarmbondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// YieldfarmbondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YieldfarmbondTransactorSession struct {
	Contract     *YieldfarmbondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// YieldfarmbondRaw is an auto generated low-level Go binding around an Ethereum contract.
type YieldfarmbondRaw struct {
	Contract *Yieldfarmbond // Generic contract binding to access the raw methods on
}

// YieldfarmbondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YieldfarmbondCallerRaw struct {
	Contract *YieldfarmbondCaller // Generic read-only contract binding to access the raw methods on
}

// YieldfarmbondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YieldfarmbondTransactorRaw struct {
	Contract *YieldfarmbondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYieldfarmbond creates a new instance of Yieldfarmbond, bound to a specific deployed contract.
func NewYieldfarmbond(address common.Address, backend bind.ContractBackend) (*Yieldfarmbond, error) {
	contract, err := bindYieldfarmbond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yieldfarmbond{YieldfarmbondCaller: YieldfarmbondCaller{contract: contract}, YieldfarmbondTransactor: YieldfarmbondTransactor{contract: contract}, YieldfarmbondFilterer: YieldfarmbondFilterer{contract: contract}}, nil
}

// NewYieldfarmbondCaller creates a new read-only instance of Yieldfarmbond, bound to a specific deployed contract.
func NewYieldfarmbondCaller(address common.Address, caller bind.ContractCaller) (*YieldfarmbondCaller, error) {
	contract, err := bindYieldfarmbond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YieldfarmbondCaller{contract: contract}, nil
}

// NewYieldfarmbondTransactor creates a new write-only instance of Yieldfarmbond, bound to a specific deployed contract.
func NewYieldfarmbondTransactor(address common.Address, transactor bind.ContractTransactor) (*YieldfarmbondTransactor, error) {
	contract, err := bindYieldfarmbond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YieldfarmbondTransactor{contract: contract}, nil
}

// NewYieldfarmbondFilterer creates a new log filterer instance of Yieldfarmbond, bound to a specific deployed contract.
func NewYieldfarmbondFilterer(address common.Address, filterer bind.ContractFilterer) (*YieldfarmbondFilterer, error) {
	contract, err := bindYieldfarmbond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YieldfarmbondFilterer{contract: contract}, nil
}

// bindYieldfarmbond binds a generic wrapper to an already deployed contract.
func bindYieldfarmbond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YieldfarmbondABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yieldfarmbond *YieldfarmbondRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yieldfarmbond.Contract.YieldfarmbondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yieldfarmbond *YieldfarmbondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yieldfarmbond.Contract.YieldfarmbondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yieldfarmbond *YieldfarmbondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yieldfarmbond.Contract.YieldfarmbondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yieldfarmbond *YieldfarmbondCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yieldfarmbond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yieldfarmbond *YieldfarmbondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yieldfarmbond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yieldfarmbond *YieldfarmbondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yieldfarmbond.Contract.contract.Transact(opts, method, params...)
}

// EPOCHSDELAYEDFROMSTAKINGCONTRACT is a free data retrieval call binding the contract method 0xa6ace8a3.
//
// Solidity: function EPOCHS_DELAYED_FROM_STAKING_CONTRACT() view returns(uint128)
func (_Yieldfarmbond *YieldfarmbondCaller) EPOCHSDELAYEDFROMSTAKINGCONTRACT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmbond.contract.Call(opts, &out, "EPOCHS_DELAYED_FROM_STAKING_CONTRACT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EPOCHSDELAYEDFROMSTAKINGCONTRACT is a free data retrieval call binding the contract method 0xa6ace8a3.
//
// Solidity: function EPOCHS_DELAYED_FROM_STAKING_CONTRACT() view returns(uint128)
func (_Yieldfarmbond *YieldfarmbondSession) EPOCHSDELAYEDFROMSTAKINGCONTRACT() (*big.Int, error) {
	return _Yieldfarmbond.Contract.EPOCHSDELAYEDFROMSTAKINGCONTRACT(&_Yieldfarmbond.CallOpts)
}

// EPOCHSDELAYEDFROMSTAKINGCONTRACT is a free data retrieval call binding the contract method 0xa6ace8a3.
//
// Solidity: function EPOCHS_DELAYED_FROM_STAKING_CONTRACT() view returns(uint128)
func (_Yieldfarmbond *YieldfarmbondCallerSession) EPOCHSDELAYEDFROMSTAKINGCONTRACT() (*big.Int, error) {
	return _Yieldfarmbond.Contract.EPOCHSDELAYEDFROMSTAKINGCONTRACT(&_Yieldfarmbond.CallOpts)
}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCaller) NROFEPOCHS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmbond.contract.Call(opts, &out, "NR_OF_EPOCHS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondSession) NROFEPOCHS() (*big.Int, error) {
	return _Yieldfarmbond.Contract.NROFEPOCHS(&_Yieldfarmbond.CallOpts)
}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCallerSession) NROFEPOCHS() (*big.Int, error) {
	return _Yieldfarmbond.Contract.NROFEPOCHS(&_Yieldfarmbond.CallOpts)
}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCaller) TOTALDISTRIBUTEDAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmbond.contract.Call(opts, &out, "TOTAL_DISTRIBUTED_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondSession) TOTALDISTRIBUTEDAMOUNT() (*big.Int, error) {
	return _Yieldfarmbond.Contract.TOTALDISTRIBUTEDAMOUNT(&_Yieldfarmbond.CallOpts)
}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCallerSession) TOTALDISTRIBUTEDAMOUNT() (*big.Int, error) {
	return _Yieldfarmbond.Contract.TOTALDISTRIBUTEDAMOUNT(&_Yieldfarmbond.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCaller) EpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmbond.contract.Call(opts, &out, "epochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondSession) EpochDuration() (*big.Int, error) {
	return _Yieldfarmbond.Contract.EpochDuration(&_Yieldfarmbond.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCallerSession) EpochDuration() (*big.Int, error) {
	return _Yieldfarmbond.Contract.EpochDuration(&_Yieldfarmbond.CallOpts)
}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCaller) EpochStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmbond.contract.Call(opts, &out, "epochStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondSession) EpochStart() (*big.Int, error) {
	return _Yieldfarmbond.Contract.EpochStart(&_Yieldfarmbond.CallOpts)
}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCallerSession) EpochStart() (*big.Int, error) {
	return _Yieldfarmbond.Contract.EpochStart(&_Yieldfarmbond.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmbond.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondSession) GetCurrentEpoch() (*big.Int, error) {
	return _Yieldfarmbond.Contract.GetCurrentEpoch(&_Yieldfarmbond.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _Yieldfarmbond.Contract.GetCurrentEpoch(&_Yieldfarmbond.CallOpts)
}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCaller) GetEpochStake(opts *bind.CallOpts, userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmbond.contract.Call(opts, &out, "getEpochStake", userAddress, epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondSession) GetEpochStake(userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _Yieldfarmbond.Contract.GetEpochStake(&_Yieldfarmbond.CallOpts, userAddress, epochId)
}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCallerSession) GetEpochStake(userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _Yieldfarmbond.Contract.GetEpochStake(&_Yieldfarmbond.CallOpts, userAddress, epochId)
}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCaller) GetPoolSize(opts *bind.CallOpts, epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmbond.contract.Call(opts, &out, "getPoolSize", epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondSession) GetPoolSize(epochId *big.Int) (*big.Int, error) {
	return _Yieldfarmbond.Contract.GetPoolSize(&_Yieldfarmbond.CallOpts, epochId)
}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCallerSession) GetPoolSize(epochId *big.Int) (*big.Int, error) {
	return _Yieldfarmbond.Contract.GetPoolSize(&_Yieldfarmbond.CallOpts, epochId)
}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_Yieldfarmbond *YieldfarmbondCaller) LastInitializedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmbond.contract.Call(opts, &out, "lastInitializedEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_Yieldfarmbond *YieldfarmbondSession) LastInitializedEpoch() (*big.Int, error) {
	return _Yieldfarmbond.Contract.LastInitializedEpoch(&_Yieldfarmbond.CallOpts)
}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_Yieldfarmbond *YieldfarmbondCallerSession) LastInitializedEpoch() (*big.Int, error) {
	return _Yieldfarmbond.Contract.LastInitializedEpoch(&_Yieldfarmbond.CallOpts)
}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCaller) UserLastEpochIdHarvested(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarmbond.contract.Call(opts, &out, "userLastEpochIdHarvested")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondSession) UserLastEpochIdHarvested() (*big.Int, error) {
	return _Yieldfarmbond.Contract.UserLastEpochIdHarvested(&_Yieldfarmbond.CallOpts)
}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_Yieldfarmbond *YieldfarmbondCallerSession) UserLastEpochIdHarvested() (*big.Int, error) {
	return _Yieldfarmbond.Contract.UserLastEpochIdHarvested(&_Yieldfarmbond.CallOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_Yieldfarmbond *YieldfarmbondTransactor) Harvest(opts *bind.TransactOpts, epochId *big.Int) (*types.Transaction, error) {
	return _Yieldfarmbond.contract.Transact(opts, "harvest", epochId)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_Yieldfarmbond *YieldfarmbondSession) Harvest(epochId *big.Int) (*types.Transaction, error) {
	return _Yieldfarmbond.Contract.Harvest(&_Yieldfarmbond.TransactOpts, epochId)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_Yieldfarmbond *YieldfarmbondTransactorSession) Harvest(epochId *big.Int) (*types.Transaction, error) {
	return _Yieldfarmbond.Contract.Harvest(&_Yieldfarmbond.TransactOpts, epochId)
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_Yieldfarmbond *YieldfarmbondTransactor) MassHarvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yieldfarmbond.contract.Transact(opts, "massHarvest")
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_Yieldfarmbond *YieldfarmbondSession) MassHarvest() (*types.Transaction, error) {
	return _Yieldfarmbond.Contract.MassHarvest(&_Yieldfarmbond.TransactOpts)
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_Yieldfarmbond *YieldfarmbondTransactorSession) MassHarvest() (*types.Transaction, error) {
	return _Yieldfarmbond.Contract.MassHarvest(&_Yieldfarmbond.TransactOpts)
}

// YieldfarmbondHarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the Yieldfarmbond contract.
type YieldfarmbondHarvestIterator struct {
	Event *YieldfarmbondHarvest // Event containing the contract specifics and raw log

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
func (it *YieldfarmbondHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YieldfarmbondHarvest)
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
		it.Event = new(YieldfarmbondHarvest)
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
func (it *YieldfarmbondHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YieldfarmbondHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YieldfarmbondHarvest represents a Harvest event raised by the Yieldfarmbond contract.
type YieldfarmbondHarvest struct {
	User    common.Address
	EpochId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0x04ad45a69eeed9c390c3a678fed2d4b90bde98e742de9936d5e0915bf3d0ea4e.
//
// Solidity: event Harvest(address indexed user, uint128 indexed epochId, uint256 amount)
func (_Yieldfarmbond *YieldfarmbondFilterer) FilterHarvest(opts *bind.FilterOpts, user []common.Address, epochId []*big.Int) (*YieldfarmbondHarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Yieldfarmbond.contract.FilterLogs(opts, "Harvest", userRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &YieldfarmbondHarvestIterator{contract: _Yieldfarmbond.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0x04ad45a69eeed9c390c3a678fed2d4b90bde98e742de9936d5e0915bf3d0ea4e.
//
// Solidity: event Harvest(address indexed user, uint128 indexed epochId, uint256 amount)
func (_Yieldfarmbond *YieldfarmbondFilterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *YieldfarmbondHarvest, user []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Yieldfarmbond.contract.WatchLogs(opts, "Harvest", userRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YieldfarmbondHarvest)
				if err := _Yieldfarmbond.contract.UnpackLog(event, "Harvest", log); err != nil {
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
func (_Yieldfarmbond *YieldfarmbondFilterer) ParseHarvest(log types.Log) (*YieldfarmbondHarvest, error) {
	event := new(YieldfarmbondHarvest)
	if err := _Yieldfarmbond.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// YieldfarmbondMassHarvestIterator is returned from FilterMassHarvest and is used to iterate over the raw logs and unpacked data for MassHarvest events raised by the Yieldfarmbond contract.
type YieldfarmbondMassHarvestIterator struct {
	Event *YieldfarmbondMassHarvest // Event containing the contract specifics and raw log

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
func (it *YieldfarmbondMassHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YieldfarmbondMassHarvest)
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
		it.Event = new(YieldfarmbondMassHarvest)
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
func (it *YieldfarmbondMassHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YieldfarmbondMassHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YieldfarmbondMassHarvest represents a MassHarvest event raised by the Yieldfarmbond contract.
type YieldfarmbondMassHarvest struct {
	User            common.Address
	EpochsHarvested *big.Int
	TotalValue      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMassHarvest is a free log retrieval operation binding the contract event 0xb68dafc1da13dc868096d0b87347c831d0bda92d178317eb1dec7f788444485c.
//
// Solidity: event MassHarvest(address indexed user, uint256 epochsHarvested, uint256 totalValue)
func (_Yieldfarmbond *YieldfarmbondFilterer) FilterMassHarvest(opts *bind.FilterOpts, user []common.Address) (*YieldfarmbondMassHarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Yieldfarmbond.contract.FilterLogs(opts, "MassHarvest", userRule)
	if err != nil {
		return nil, err
	}
	return &YieldfarmbondMassHarvestIterator{contract: _Yieldfarmbond.contract, event: "MassHarvest", logs: logs, sub: sub}, nil
}

// WatchMassHarvest is a free log subscription operation binding the contract event 0xb68dafc1da13dc868096d0b87347c831d0bda92d178317eb1dec7f788444485c.
//
// Solidity: event MassHarvest(address indexed user, uint256 epochsHarvested, uint256 totalValue)
func (_Yieldfarmbond *YieldfarmbondFilterer) WatchMassHarvest(opts *bind.WatchOpts, sink chan<- *YieldfarmbondMassHarvest, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Yieldfarmbond.contract.WatchLogs(opts, "MassHarvest", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YieldfarmbondMassHarvest)
				if err := _Yieldfarmbond.contract.UnpackLog(event, "MassHarvest", log); err != nil {
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
func (_Yieldfarmbond *YieldfarmbondFilterer) ParseMassHarvest(log types.Log) (*YieldfarmbondMassHarvest, error) {
	event := new(YieldfarmbondMassHarvest)
	if err := _Yieldfarmbond.contract.UnpackLog(event, "MassHarvest", log); err != nil {
		return nil, err
	}
	return event, nil
}
