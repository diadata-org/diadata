// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package diaOracleV2MultiupdateService

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

// DiaOracleV2MultiupdateServiceMetaData contains all meta data concerning the DiaOracleV2MultiupdateService contract.
var DiaOracleV2MultiupdateServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"name\":\"OracleUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUpdater\",\"type\":\"address\"}],\"name\":\"UpdaterAddressChange\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"compressedValues\",\"type\":\"uint256[]\"}],\"name\":\"setMultipleValues\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOracleUpdaterAddress\",\"type\":\"address\"}],\"name\":\"updateOracleUpdaterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DiaOracleV2MultiupdateServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use DiaOracleV2MultiupdateServiceMetaData.ABI instead.
var DiaOracleV2MultiupdateServiceABI = DiaOracleV2MultiupdateServiceMetaData.ABI

// DiaOracleV2MultiupdateService is an auto generated Go binding around an Ethereum contract.
type DiaOracleV2MultiupdateService struct {
	DiaOracleV2MultiupdateServiceCaller     // Read-only binding to the contract
	DiaOracleV2MultiupdateServiceTransactor // Write-only binding to the contract
	DiaOracleV2MultiupdateServiceFilterer   // Log filterer for contract events
}

// DiaOracleV2MultiupdateServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiaOracleV2MultiupdateServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiaOracleV2MultiupdateServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiaOracleV2MultiupdateServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiaOracleV2MultiupdateServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiaOracleV2MultiupdateServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiaOracleV2MultiupdateServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiaOracleV2MultiupdateServiceSession struct {
	Contract     *DiaOracleV2MultiupdateService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// DiaOracleV2MultiupdateServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiaOracleV2MultiupdateServiceCallerSession struct {
	Contract *DiaOracleV2MultiupdateServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// DiaOracleV2MultiupdateServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiaOracleV2MultiupdateServiceTransactorSession struct {
	Contract     *DiaOracleV2MultiupdateServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// DiaOracleV2MultiupdateServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiaOracleV2MultiupdateServiceRaw struct {
	Contract *DiaOracleV2MultiupdateService // Generic contract binding to access the raw methods on
}

// DiaOracleV2MultiupdateServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiaOracleV2MultiupdateServiceCallerRaw struct {
	Contract *DiaOracleV2MultiupdateServiceCaller // Generic read-only contract binding to access the raw methods on
}

// DiaOracleV2MultiupdateServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiaOracleV2MultiupdateServiceTransactorRaw struct {
	Contract *DiaOracleV2MultiupdateServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiaOracleV2MultiupdateService creates a new instance of DiaOracleV2MultiupdateService, bound to a specific deployed contract.
func NewDiaOracleV2MultiupdateService(address common.Address, backend bind.ContractBackend) (*DiaOracleV2MultiupdateService, error) {
	contract, err := bindDiaOracleV2MultiupdateService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiaOracleV2MultiupdateService{DiaOracleV2MultiupdateServiceCaller: DiaOracleV2MultiupdateServiceCaller{contract: contract}, DiaOracleV2MultiupdateServiceTransactor: DiaOracleV2MultiupdateServiceTransactor{contract: contract}, DiaOracleV2MultiupdateServiceFilterer: DiaOracleV2MultiupdateServiceFilterer{contract: contract}}, nil
}

// NewDiaOracleV2MultiupdateServiceCaller creates a new read-only instance of DiaOracleV2MultiupdateService, bound to a specific deployed contract.
func NewDiaOracleV2MultiupdateServiceCaller(address common.Address, caller bind.ContractCaller) (*DiaOracleV2MultiupdateServiceCaller, error) {
	contract, err := bindDiaOracleV2MultiupdateService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiaOracleV2MultiupdateServiceCaller{contract: contract}, nil
}

// NewDiaOracleV2MultiupdateServiceTransactor creates a new write-only instance of DiaOracleV2MultiupdateService, bound to a specific deployed contract.
func NewDiaOracleV2MultiupdateServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*DiaOracleV2MultiupdateServiceTransactor, error) {
	contract, err := bindDiaOracleV2MultiupdateService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiaOracleV2MultiupdateServiceTransactor{contract: contract}, nil
}

// NewDiaOracleV2MultiupdateServiceFilterer creates a new log filterer instance of DiaOracleV2MultiupdateService, bound to a specific deployed contract.
func NewDiaOracleV2MultiupdateServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*DiaOracleV2MultiupdateServiceFilterer, error) {
	contract, err := bindDiaOracleV2MultiupdateService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiaOracleV2MultiupdateServiceFilterer{contract: contract}, nil
}

// bindDiaOracleV2MultiupdateService binds a generic wrapper to an already deployed contract.
func bindDiaOracleV2MultiupdateService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DiaOracleV2MultiupdateServiceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiaOracleV2MultiupdateService.Contract.DiaOracleV2MultiupdateServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.DiaOracleV2MultiupdateServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.DiaOracleV2MultiupdateServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiaOracleV2MultiupdateService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.contract.Transact(opts, method, params...)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceCaller) GetValue(opts *bind.CallOpts, key string) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _DiaOracleV2MultiupdateService.contract.Call(opts, &out, "getValue", key)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceSession) GetValue(key string) (*big.Int, *big.Int, error) {
	return _DiaOracleV2MultiupdateService.Contract.GetValue(&_DiaOracleV2MultiupdateService.CallOpts, key)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceCallerSession) GetValue(key string) (*big.Int, *big.Int, error) {
	return _DiaOracleV2MultiupdateService.Contract.GetValue(&_DiaOracleV2MultiupdateService.CallOpts, key)
}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceCaller) Values(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _DiaOracleV2MultiupdateService.contract.Call(opts, &out, "values", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceSession) Values(arg0 string) (*big.Int, error) {
	return _DiaOracleV2MultiupdateService.Contract.Values(&_DiaOracleV2MultiupdateService.CallOpts, arg0)
}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceCallerSession) Values(arg0 string) (*big.Int, error) {
	return _DiaOracleV2MultiupdateService.Contract.Values(&_DiaOracleV2MultiupdateService.CallOpts, arg0)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x8d241526.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] compressedValues) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactor) SetMultipleValues(opts *bind.TransactOpts, keys []string, compressedValues []*big.Int) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.contract.Transact(opts, "setMultipleValues", keys, compressedValues)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x8d241526.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] compressedValues) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceSession) SetMultipleValues(keys []string, compressedValues []*big.Int) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.SetMultipleValues(&_DiaOracleV2MultiupdateService.TransactOpts, keys, compressedValues)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x8d241526.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] compressedValues) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactorSession) SetMultipleValues(keys []string, compressedValues []*big.Int) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.SetMultipleValues(&_DiaOracleV2MultiupdateService.TransactOpts, keys, compressedValues)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactor) SetValue(opts *bind.TransactOpts, key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.contract.Transact(opts, "setValue", key, value, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceSession) SetValue(key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.SetValue(&_DiaOracleV2MultiupdateService.TransactOpts, key, value, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactorSession) SetValue(key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.SetValue(&_DiaOracleV2MultiupdateService.TransactOpts, key, value, timestamp)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactor) UpdateOracleUpdaterAddress(opts *bind.TransactOpts, newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.contract.Transact(opts, "updateOracleUpdaterAddress", newOracleUpdaterAddress)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceSession) UpdateOracleUpdaterAddress(newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.UpdateOracleUpdaterAddress(&_DiaOracleV2MultiupdateService.TransactOpts, newOracleUpdaterAddress)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactorSession) UpdateOracleUpdaterAddress(newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.UpdateOracleUpdaterAddress(&_DiaOracleV2MultiupdateService.TransactOpts, newOracleUpdaterAddress)
}

// DiaOracleV2MultiupdateServiceOracleUpdateIterator is returned from FilterOracleUpdate and is used to iterate over the raw logs and unpacked data for OracleUpdate events raised by the DiaOracleV2MultiupdateService contract.
type DiaOracleV2MultiupdateServiceOracleUpdateIterator struct {
	Event *DiaOracleV2MultiupdateServiceOracleUpdate // Event containing the contract specifics and raw log

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
func (it *DiaOracleV2MultiupdateServiceOracleUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiaOracleV2MultiupdateServiceOracleUpdate)
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
		it.Event = new(DiaOracleV2MultiupdateServiceOracleUpdate)
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
func (it *DiaOracleV2MultiupdateServiceOracleUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiaOracleV2MultiupdateServiceOracleUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiaOracleV2MultiupdateServiceOracleUpdate represents a OracleUpdate event raised by the DiaOracleV2MultiupdateService contract.
type DiaOracleV2MultiupdateServiceOracleUpdate struct {
	Key       string
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleUpdate is a free log retrieval operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceFilterer) FilterOracleUpdate(opts *bind.FilterOpts) (*DiaOracleV2MultiupdateServiceOracleUpdateIterator, error) {

	logs, sub, err := _DiaOracleV2MultiupdateService.contract.FilterLogs(opts, "OracleUpdate")
	if err != nil {
		return nil, err
	}
	return &DiaOracleV2MultiupdateServiceOracleUpdateIterator{contract: _DiaOracleV2MultiupdateService.contract, event: "OracleUpdate", logs: logs, sub: sub}, nil
}

// WatchOracleUpdate is a free log subscription operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceFilterer) WatchOracleUpdate(opts *bind.WatchOpts, sink chan<- *DiaOracleV2MultiupdateServiceOracleUpdate) (event.Subscription, error) {

	logs, sub, err := _DiaOracleV2MultiupdateService.contract.WatchLogs(opts, "OracleUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiaOracleV2MultiupdateServiceOracleUpdate)
				if err := _DiaOracleV2MultiupdateService.contract.UnpackLog(event, "OracleUpdate", log); err != nil {
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

// ParseOracleUpdate is a log parse operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceFilterer) ParseOracleUpdate(log types.Log) (*DiaOracleV2MultiupdateServiceOracleUpdate, error) {
	event := new(DiaOracleV2MultiupdateServiceOracleUpdate)
	if err := _DiaOracleV2MultiupdateService.contract.UnpackLog(event, "OracleUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiaOracleV2MultiupdateServiceUpdaterAddressChangeIterator is returned from FilterUpdaterAddressChange and is used to iterate over the raw logs and unpacked data for UpdaterAddressChange events raised by the DiaOracleV2MultiupdateService contract.
type DiaOracleV2MultiupdateServiceUpdaterAddressChangeIterator struct {
	Event *DiaOracleV2MultiupdateServiceUpdaterAddressChange // Event containing the contract specifics and raw log

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
func (it *DiaOracleV2MultiupdateServiceUpdaterAddressChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiaOracleV2MultiupdateServiceUpdaterAddressChange)
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
		it.Event = new(DiaOracleV2MultiupdateServiceUpdaterAddressChange)
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
func (it *DiaOracleV2MultiupdateServiceUpdaterAddressChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiaOracleV2MultiupdateServiceUpdaterAddressChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiaOracleV2MultiupdateServiceUpdaterAddressChange represents a UpdaterAddressChange event raised by the DiaOracleV2MultiupdateService contract.
type DiaOracleV2MultiupdateServiceUpdaterAddressChange struct {
	NewUpdater common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdaterAddressChange is a free log retrieval operation binding the contract event 0x121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f.
//
// Solidity: event UpdaterAddressChange(address newUpdater)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceFilterer) FilterUpdaterAddressChange(opts *bind.FilterOpts) (*DiaOracleV2MultiupdateServiceUpdaterAddressChangeIterator, error) {

	logs, sub, err := _DiaOracleV2MultiupdateService.contract.FilterLogs(opts, "UpdaterAddressChange")
	if err != nil {
		return nil, err
	}
	return &DiaOracleV2MultiupdateServiceUpdaterAddressChangeIterator{contract: _DiaOracleV2MultiupdateService.contract, event: "UpdaterAddressChange", logs: logs, sub: sub}, nil
}

// WatchUpdaterAddressChange is a free log subscription operation binding the contract event 0x121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f.
//
// Solidity: event UpdaterAddressChange(address newUpdater)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceFilterer) WatchUpdaterAddressChange(opts *bind.WatchOpts, sink chan<- *DiaOracleV2MultiupdateServiceUpdaterAddressChange) (event.Subscription, error) {

	logs, sub, err := _DiaOracleV2MultiupdateService.contract.WatchLogs(opts, "UpdaterAddressChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiaOracleV2MultiupdateServiceUpdaterAddressChange)
				if err := _DiaOracleV2MultiupdateService.contract.UnpackLog(event, "UpdaterAddressChange", log); err != nil {
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

// ParseUpdaterAddressChange is a log parse operation binding the contract event 0x121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f.
//
// Solidity: event UpdaterAddressChange(address newUpdater)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceFilterer) ParseUpdaterAddressChange(log types.Log) (*DiaOracleV2MultiupdateServiceUpdaterAddressChange, error) {
	event := new(DiaOracleV2MultiupdateServiceUpdaterAddressChange)
	if err := _DiaOracleV2MultiupdateService.contract.UnpackLog(event, "UpdaterAddressChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
