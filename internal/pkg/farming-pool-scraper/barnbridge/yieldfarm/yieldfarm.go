// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yieldfarm

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

// YieldfarmABI is the input ABI used to generate the binding from.
const YieldfarmABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bondTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"susd\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dai\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakeContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"communityVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochsHarvested\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"}],\"name\":\"MassHarvest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NR_OF_EPOCHS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_DISTRIBUTED_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getEpochStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getPoolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"harvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInitializedEpoch\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massHarvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userLastEpochIdHarvested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Yieldfarm is an auto generated Go binding around an Ethereum contract.
type Yieldfarm struct {
	YieldfarmCaller     // Read-only binding to the contract
	YieldfarmTransactor // Write-only binding to the contract
	YieldfarmFilterer   // Log filterer for contract events
}

// YieldfarmCaller is an auto generated read-only Go binding around an Ethereum contract.
type YieldfarmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldfarmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YieldfarmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldfarmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YieldfarmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldfarmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YieldfarmSession struct {
	Contract     *Yieldfarm        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YieldfarmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YieldfarmCallerSession struct {
	Contract *YieldfarmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// YieldfarmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YieldfarmTransactorSession struct {
	Contract     *YieldfarmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// YieldfarmRaw is an auto generated low-level Go binding around an Ethereum contract.
type YieldfarmRaw struct {
	Contract *Yieldfarm // Generic contract binding to access the raw methods on
}

// YieldfarmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YieldfarmCallerRaw struct {
	Contract *YieldfarmCaller // Generic read-only contract binding to access the raw methods on
}

// YieldfarmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YieldfarmTransactorRaw struct {
	Contract *YieldfarmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYieldfarm creates a new instance of Yieldfarm, bound to a specific deployed contract.
func NewYieldfarm(address common.Address, backend bind.ContractBackend) (*Yieldfarm, error) {
	contract, err := bindYieldfarm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yieldfarm{YieldfarmCaller: YieldfarmCaller{contract: contract}, YieldfarmTransactor: YieldfarmTransactor{contract: contract}, YieldfarmFilterer: YieldfarmFilterer{contract: contract}}, nil
}

// NewYieldfarmCaller creates a new read-only instance of Yieldfarm, bound to a specific deployed contract.
func NewYieldfarmCaller(address common.Address, caller bind.ContractCaller) (*YieldfarmCaller, error) {
	contract, err := bindYieldfarm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YieldfarmCaller{contract: contract}, nil
}

// NewYieldfarmTransactor creates a new write-only instance of Yieldfarm, bound to a specific deployed contract.
func NewYieldfarmTransactor(address common.Address, transactor bind.ContractTransactor) (*YieldfarmTransactor, error) {
	contract, err := bindYieldfarm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YieldfarmTransactor{contract: contract}, nil
}

// NewYieldfarmFilterer creates a new log filterer instance of Yieldfarm, bound to a specific deployed contract.
func NewYieldfarmFilterer(address common.Address, filterer bind.ContractFilterer) (*YieldfarmFilterer, error) {
	contract, err := bindYieldfarm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YieldfarmFilterer{contract: contract}, nil
}

// bindYieldfarm binds a generic wrapper to an already deployed contract.
func bindYieldfarm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YieldfarmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yieldfarm *YieldfarmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yieldfarm.Contract.YieldfarmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yieldfarm *YieldfarmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yieldfarm.Contract.YieldfarmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yieldfarm *YieldfarmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yieldfarm.Contract.YieldfarmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yieldfarm *YieldfarmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yieldfarm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yieldfarm *YieldfarmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yieldfarm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yieldfarm *YieldfarmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yieldfarm.Contract.contract.Transact(opts, method, params...)
}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_Yieldfarm *YieldfarmCaller) NROFEPOCHS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarm.contract.Call(opts, &out, "NR_OF_EPOCHS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_Yieldfarm *YieldfarmSession) NROFEPOCHS() (*big.Int, error) {
	return _Yieldfarm.Contract.NROFEPOCHS(&_Yieldfarm.CallOpts)
}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_Yieldfarm *YieldfarmCallerSession) NROFEPOCHS() (*big.Int, error) {
	return _Yieldfarm.Contract.NROFEPOCHS(&_Yieldfarm.CallOpts)
}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_Yieldfarm *YieldfarmCaller) TOTALDISTRIBUTEDAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarm.contract.Call(opts, &out, "TOTAL_DISTRIBUTED_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_Yieldfarm *YieldfarmSession) TOTALDISTRIBUTEDAMOUNT() (*big.Int, error) {
	return _Yieldfarm.Contract.TOTALDISTRIBUTEDAMOUNT(&_Yieldfarm.CallOpts)
}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_Yieldfarm *YieldfarmCallerSession) TOTALDISTRIBUTEDAMOUNT() (*big.Int, error) {
	return _Yieldfarm.Contract.TOTALDISTRIBUTEDAMOUNT(&_Yieldfarm.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Yieldfarm *YieldfarmCaller) EpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarm.contract.Call(opts, &out, "epochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Yieldfarm *YieldfarmSession) EpochDuration() (*big.Int, error) {
	return _Yieldfarm.Contract.EpochDuration(&_Yieldfarm.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Yieldfarm *YieldfarmCallerSession) EpochDuration() (*big.Int, error) {
	return _Yieldfarm.Contract.EpochDuration(&_Yieldfarm.CallOpts)
}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_Yieldfarm *YieldfarmCaller) EpochStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarm.contract.Call(opts, &out, "epochStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_Yieldfarm *YieldfarmSession) EpochStart() (*big.Int, error) {
	return _Yieldfarm.Contract.EpochStart(&_Yieldfarm.CallOpts)
}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_Yieldfarm *YieldfarmCallerSession) EpochStart() (*big.Int, error) {
	return _Yieldfarm.Contract.EpochStart(&_Yieldfarm.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Yieldfarm *YieldfarmCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarm.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Yieldfarm *YieldfarmSession) GetCurrentEpoch() (*big.Int, error) {
	return _Yieldfarm.Contract.GetCurrentEpoch(&_Yieldfarm.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Yieldfarm *YieldfarmCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _Yieldfarm.Contract.GetCurrentEpoch(&_Yieldfarm.CallOpts)
}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_Yieldfarm *YieldfarmCaller) GetEpochStake(opts *bind.CallOpts, userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarm.contract.Call(opts, &out, "getEpochStake", userAddress, epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_Yieldfarm *YieldfarmSession) GetEpochStake(userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _Yieldfarm.Contract.GetEpochStake(&_Yieldfarm.CallOpts, userAddress, epochId)
}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_Yieldfarm *YieldfarmCallerSession) GetEpochStake(userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _Yieldfarm.Contract.GetEpochStake(&_Yieldfarm.CallOpts, userAddress, epochId)
}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_Yieldfarm *YieldfarmCaller) GetPoolSize(opts *bind.CallOpts, epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarm.contract.Call(opts, &out, "getPoolSize", epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_Yieldfarm *YieldfarmSession) GetPoolSize(epochId *big.Int) (*big.Int, error) {
	return _Yieldfarm.Contract.GetPoolSize(&_Yieldfarm.CallOpts, epochId)
}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_Yieldfarm *YieldfarmCallerSession) GetPoolSize(epochId *big.Int) (*big.Int, error) {
	return _Yieldfarm.Contract.GetPoolSize(&_Yieldfarm.CallOpts, epochId)
}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_Yieldfarm *YieldfarmCaller) LastInitializedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarm.contract.Call(opts, &out, "lastInitializedEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_Yieldfarm *YieldfarmSession) LastInitializedEpoch() (*big.Int, error) {
	return _Yieldfarm.Contract.LastInitializedEpoch(&_Yieldfarm.CallOpts)
}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_Yieldfarm *YieldfarmCallerSession) LastInitializedEpoch() (*big.Int, error) {
	return _Yieldfarm.Contract.LastInitializedEpoch(&_Yieldfarm.CallOpts)
}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_Yieldfarm *YieldfarmCaller) UserLastEpochIdHarvested(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yieldfarm.contract.Call(opts, &out, "userLastEpochIdHarvested")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_Yieldfarm *YieldfarmSession) UserLastEpochIdHarvested() (*big.Int, error) {
	return _Yieldfarm.Contract.UserLastEpochIdHarvested(&_Yieldfarm.CallOpts)
}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_Yieldfarm *YieldfarmCallerSession) UserLastEpochIdHarvested() (*big.Int, error) {
	return _Yieldfarm.Contract.UserLastEpochIdHarvested(&_Yieldfarm.CallOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_Yieldfarm *YieldfarmTransactor) Harvest(opts *bind.TransactOpts, epochId *big.Int) (*types.Transaction, error) {
	return _Yieldfarm.contract.Transact(opts, "harvest", epochId)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_Yieldfarm *YieldfarmSession) Harvest(epochId *big.Int) (*types.Transaction, error) {
	return _Yieldfarm.Contract.Harvest(&_Yieldfarm.TransactOpts, epochId)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_Yieldfarm *YieldfarmTransactorSession) Harvest(epochId *big.Int) (*types.Transaction, error) {
	return _Yieldfarm.Contract.Harvest(&_Yieldfarm.TransactOpts, epochId)
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_Yieldfarm *YieldfarmTransactor) MassHarvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yieldfarm.contract.Transact(opts, "massHarvest")
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_Yieldfarm *YieldfarmSession) MassHarvest() (*types.Transaction, error) {
	return _Yieldfarm.Contract.MassHarvest(&_Yieldfarm.TransactOpts)
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_Yieldfarm *YieldfarmTransactorSession) MassHarvest() (*types.Transaction, error) {
	return _Yieldfarm.Contract.MassHarvest(&_Yieldfarm.TransactOpts)
}

// YieldfarmHarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the Yieldfarm contract.
type YieldfarmHarvestIterator struct {
	Event *YieldfarmHarvest // Event containing the contract specifics and raw log

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
func (it *YieldfarmHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YieldfarmHarvest)
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
		it.Event = new(YieldfarmHarvest)
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
func (it *YieldfarmHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YieldfarmHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YieldfarmHarvest represents a Harvest event raised by the Yieldfarm contract.
type YieldfarmHarvest struct {
	User    common.Address
	EpochId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0x04ad45a69eeed9c390c3a678fed2d4b90bde98e742de9936d5e0915bf3d0ea4e.
//
// Solidity: event Harvest(address indexed user, uint128 indexed epochId, uint256 amount)
func (_Yieldfarm *YieldfarmFilterer) FilterHarvest(opts *bind.FilterOpts, user []common.Address, epochId []*big.Int) (*YieldfarmHarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Yieldfarm.contract.FilterLogs(opts, "Harvest", userRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &YieldfarmHarvestIterator{contract: _Yieldfarm.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0x04ad45a69eeed9c390c3a678fed2d4b90bde98e742de9936d5e0915bf3d0ea4e.
//
// Solidity: event Harvest(address indexed user, uint128 indexed epochId, uint256 amount)
func (_Yieldfarm *YieldfarmFilterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *YieldfarmHarvest, user []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Yieldfarm.contract.WatchLogs(opts, "Harvest", userRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YieldfarmHarvest)
				if err := _Yieldfarm.contract.UnpackLog(event, "Harvest", log); err != nil {
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
func (_Yieldfarm *YieldfarmFilterer) ParseHarvest(log types.Log) (*YieldfarmHarvest, error) {
	event := new(YieldfarmHarvest)
	if err := _Yieldfarm.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// YieldfarmMassHarvestIterator is returned from FilterMassHarvest and is used to iterate over the raw logs and unpacked data for MassHarvest events raised by the Yieldfarm contract.
type YieldfarmMassHarvestIterator struct {
	Event *YieldfarmMassHarvest // Event containing the contract specifics and raw log

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
func (it *YieldfarmMassHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YieldfarmMassHarvest)
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
		it.Event = new(YieldfarmMassHarvest)
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
func (it *YieldfarmMassHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YieldfarmMassHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YieldfarmMassHarvest represents a MassHarvest event raised by the Yieldfarm contract.
type YieldfarmMassHarvest struct {
	User            common.Address
	EpochsHarvested *big.Int
	TotalValue      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMassHarvest is a free log retrieval operation binding the contract event 0xb68dafc1da13dc868096d0b87347c831d0bda92d178317eb1dec7f788444485c.
//
// Solidity: event MassHarvest(address indexed user, uint256 epochsHarvested, uint256 totalValue)
func (_Yieldfarm *YieldfarmFilterer) FilterMassHarvest(opts *bind.FilterOpts, user []common.Address) (*YieldfarmMassHarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Yieldfarm.contract.FilterLogs(opts, "MassHarvest", userRule)
	if err != nil {
		return nil, err
	}
	return &YieldfarmMassHarvestIterator{contract: _Yieldfarm.contract, event: "MassHarvest", logs: logs, sub: sub}, nil
}

// WatchMassHarvest is a free log subscription operation binding the contract event 0xb68dafc1da13dc868096d0b87347c831d0bda92d178317eb1dec7f788444485c.
//
// Solidity: event MassHarvest(address indexed user, uint256 epochsHarvested, uint256 totalValue)
func (_Yieldfarm *YieldfarmFilterer) WatchMassHarvest(opts *bind.WatchOpts, sink chan<- *YieldfarmMassHarvest, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Yieldfarm.contract.WatchLogs(opts, "MassHarvest", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YieldfarmMassHarvest)
				if err := _Yieldfarm.contract.UnpackLog(event, "MassHarvest", log); err != nil {
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
func (_Yieldfarm *YieldfarmFilterer) ParseMassHarvest(log types.Log) (*YieldfarmMassHarvest, error) {
	event := new(YieldfarmMassHarvest)
	if err := _Yieldfarm.contract.UnpackLog(event, "MassHarvest", log); err != nil {
		return nil, err
	}
	return event, nil
}
