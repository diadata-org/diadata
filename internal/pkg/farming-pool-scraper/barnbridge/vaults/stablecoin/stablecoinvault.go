// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stablecoinvault

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

// VaultABI is the input ABI used to generate the binding from.
const VaultABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bondTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"susd\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dai\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakeContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"communityVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochsHarvested\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"}],\"name\":\"MassHarvest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NR_OF_EPOCHS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_DISTRIBUTED_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getEpochStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getPoolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"harvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInitializedEpoch\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massHarvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userLastEpochIdHarvested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_Vault *VaultCaller) NROFEPOCHS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "NR_OF_EPOCHS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_Vault *VaultSession) NROFEPOCHS() (*big.Int, error) {
	return _Vault.Contract.NROFEPOCHS(&_Vault.CallOpts)
}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_Vault *VaultCallerSession) NROFEPOCHS() (*big.Int, error) {
	return _Vault.Contract.NROFEPOCHS(&_Vault.CallOpts)
}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_Vault *VaultCaller) TOTALDISTRIBUTEDAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "TOTAL_DISTRIBUTED_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_Vault *VaultSession) TOTALDISTRIBUTEDAMOUNT() (*big.Int, error) {
	return _Vault.Contract.TOTALDISTRIBUTEDAMOUNT(&_Vault.CallOpts)
}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_Vault *VaultCallerSession) TOTALDISTRIBUTEDAMOUNT() (*big.Int, error) {
	return _Vault.Contract.TOTALDISTRIBUTEDAMOUNT(&_Vault.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Vault *VaultCaller) EpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "epochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Vault *VaultSession) EpochDuration() (*big.Int, error) {
	return _Vault.Contract.EpochDuration(&_Vault.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Vault *VaultCallerSession) EpochDuration() (*big.Int, error) {
	return _Vault.Contract.EpochDuration(&_Vault.CallOpts)
}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_Vault *VaultCaller) EpochStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "epochStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_Vault *VaultSession) EpochStart() (*big.Int, error) {
	return _Vault.Contract.EpochStart(&_Vault.CallOpts)
}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_Vault *VaultCallerSession) EpochStart() (*big.Int, error) {
	return _Vault.Contract.EpochStart(&_Vault.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Vault *VaultCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Vault *VaultSession) GetCurrentEpoch() (*big.Int, error) {
	return _Vault.Contract.GetCurrentEpoch(&_Vault.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Vault *VaultCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _Vault.Contract.GetCurrentEpoch(&_Vault.CallOpts)
}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_Vault *VaultCaller) GetEpochStake(opts *bind.CallOpts, userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getEpochStake", userAddress, epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_Vault *VaultSession) GetEpochStake(userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetEpochStake(&_Vault.CallOpts, userAddress, epochId)
}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_Vault *VaultCallerSession) GetEpochStake(userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetEpochStake(&_Vault.CallOpts, userAddress, epochId)
}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_Vault *VaultCaller) GetPoolSize(opts *bind.CallOpts, epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPoolSize", epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_Vault *VaultSession) GetPoolSize(epochId *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetPoolSize(&_Vault.CallOpts, epochId)
}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_Vault *VaultCallerSession) GetPoolSize(epochId *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetPoolSize(&_Vault.CallOpts, epochId)
}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_Vault *VaultCaller) LastInitializedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "lastInitializedEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_Vault *VaultSession) LastInitializedEpoch() (*big.Int, error) {
	return _Vault.Contract.LastInitializedEpoch(&_Vault.CallOpts)
}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_Vault *VaultCallerSession) LastInitializedEpoch() (*big.Int, error) {
	return _Vault.Contract.LastInitializedEpoch(&_Vault.CallOpts)
}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_Vault *VaultCaller) UserLastEpochIdHarvested(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "userLastEpochIdHarvested")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_Vault *VaultSession) UserLastEpochIdHarvested() (*big.Int, error) {
	return _Vault.Contract.UserLastEpochIdHarvested(&_Vault.CallOpts)
}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_Vault *VaultCallerSession) UserLastEpochIdHarvested() (*big.Int, error) {
	return _Vault.Contract.UserLastEpochIdHarvested(&_Vault.CallOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_Vault *VaultTransactor) Harvest(opts *bind.TransactOpts, epochId *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "harvest", epochId)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_Vault *VaultSession) Harvest(epochId *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Harvest(&_Vault.TransactOpts, epochId)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_Vault *VaultTransactorSession) Harvest(epochId *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Harvest(&_Vault.TransactOpts, epochId)
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_Vault *VaultTransactor) MassHarvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "massHarvest")
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_Vault *VaultSession) MassHarvest() (*types.Transaction, error) {
	return _Vault.Contract.MassHarvest(&_Vault.TransactOpts)
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_Vault *VaultTransactorSession) MassHarvest() (*types.Transaction, error) {
	return _Vault.Contract.MassHarvest(&_Vault.TransactOpts)
}

// VaultHarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the Vault contract.
type VaultHarvestIterator struct {
	Event *VaultHarvest // Event containing the contract specifics and raw log

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
func (it *VaultHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultHarvest)
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
		it.Event = new(VaultHarvest)
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
func (it *VaultHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultHarvest represents a Harvest event raised by the Vault contract.
type VaultHarvest struct {
	User    common.Address
	EpochId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0x04ad45a69eeed9c390c3a678fed2d4b90bde98e742de9936d5e0915bf3d0ea4e.
//
// Solidity: event Harvest(address indexed user, uint128 indexed epochId, uint256 amount)
func (_Vault *VaultFilterer) FilterHarvest(opts *bind.FilterOpts, user []common.Address, epochId []*big.Int) (*VaultHarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Harvest", userRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &VaultHarvestIterator{contract: _Vault.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0x04ad45a69eeed9c390c3a678fed2d4b90bde98e742de9936d5e0915bf3d0ea4e.
//
// Solidity: event Harvest(address indexed user, uint128 indexed epochId, uint256 amount)
func (_Vault *VaultFilterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *VaultHarvest, user []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Harvest", userRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultHarvest)
				if err := _Vault.contract.UnpackLog(event, "Harvest", log); err != nil {
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
func (_Vault *VaultFilterer) ParseHarvest(log types.Log) (*VaultHarvest, error) {
	event := new(VaultHarvest)
	if err := _Vault.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultMassHarvestIterator is returned from FilterMassHarvest and is used to iterate over the raw logs and unpacked data for MassHarvest events raised by the Vault contract.
type VaultMassHarvestIterator struct {
	Event *VaultMassHarvest // Event containing the contract specifics and raw log

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
func (it *VaultMassHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultMassHarvest)
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
		it.Event = new(VaultMassHarvest)
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
func (it *VaultMassHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultMassHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultMassHarvest represents a MassHarvest event raised by the Vault contract.
type VaultMassHarvest struct {
	User            common.Address
	EpochsHarvested *big.Int
	TotalValue      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMassHarvest is a free log retrieval operation binding the contract event 0xb68dafc1da13dc868096d0b87347c831d0bda92d178317eb1dec7f788444485c.
//
// Solidity: event MassHarvest(address indexed user, uint256 epochsHarvested, uint256 totalValue)
func (_Vault *VaultFilterer) FilterMassHarvest(opts *bind.FilterOpts, user []common.Address) (*VaultMassHarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "MassHarvest", userRule)
	if err != nil {
		return nil, err
	}
	return &VaultMassHarvestIterator{contract: _Vault.contract, event: "MassHarvest", logs: logs, sub: sub}, nil
}

// WatchMassHarvest is a free log subscription operation binding the contract event 0xb68dafc1da13dc868096d0b87347c831d0bda92d178317eb1dec7f788444485c.
//
// Solidity: event MassHarvest(address indexed user, uint256 epochsHarvested, uint256 totalValue)
func (_Vault *VaultFilterer) WatchMassHarvest(opts *bind.WatchOpts, sink chan<- *VaultMassHarvest, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "MassHarvest", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultMassHarvest)
				if err := _Vault.contract.UnpackLog(event, "MassHarvest", log); err != nil {
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
func (_Vault *VaultFilterer) ParseMassHarvest(log types.Log) (*VaultMassHarvest, error) {
	event := new(VaultMassHarvest)
	if err := _Vault.contract.UnpackLog(event, "MassHarvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
