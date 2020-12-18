// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dforce

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

// DforceABI is the input ABI used to generate the binding from.
const DforceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_implementation\",\"type\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"}]"

// Dforce is an auto generated Go binding around an Ethereum contract.
type Dforce struct {
	DforceCaller     // Read-only binding to the contract
	DforceTransactor // Write-only binding to the contract
	DforceFilterer   // Log filterer for contract events
}

// DforceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DforceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DforceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DforceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DforceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DforceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DforceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DforceSession struct {
	Contract     *Dforce           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DforceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DforceCallerSession struct {
	Contract *DforceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DforceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DforceTransactorSession struct {
	Contract     *DforceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DforceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DforceRaw struct {
	Contract *Dforce // Generic contract binding to access the raw methods on
}

// DforceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DforceCallerRaw struct {
	Contract *DforceCaller // Generic read-only contract binding to access the raw methods on
}

// DforceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DforceTransactorRaw struct {
	Contract *DforceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDforce creates a new instance of Dforce, bound to a specific deployed contract.
func NewDforce(address common.Address, backend bind.ContractBackend) (*Dforce, error) {
	contract, err := bindDforce(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dforce{DforceCaller: DforceCaller{contract: contract}, DforceTransactor: DforceTransactor{contract: contract}, DforceFilterer: DforceFilterer{contract: contract}}, nil
}

// NewDforceCaller creates a new read-only instance of Dforce, bound to a specific deployed contract.
func NewDforceCaller(address common.Address, caller bind.ContractCaller) (*DforceCaller, error) {
	contract, err := bindDforce(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DforceCaller{contract: contract}, nil
}

// NewDforceTransactor creates a new write-only instance of Dforce, bound to a specific deployed contract.
func NewDforceTransactor(address common.Address, transactor bind.ContractTransactor) (*DforceTransactor, error) {
	contract, err := bindDforce(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DforceTransactor{contract: contract}, nil
}

// NewDforceFilterer creates a new log filterer instance of Dforce, bound to a specific deployed contract.
func NewDforceFilterer(address common.Address, filterer bind.ContractFilterer) (*DforceFilterer, error) {
	contract, err := bindDforce(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DforceFilterer{contract: contract}, nil
}

// bindDforce binds a generic wrapper to an already deployed contract.
func bindDforce(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DforceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dforce *DforceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dforce.Contract.DforceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dforce *DforceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dforce.Contract.DforceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dforce *DforceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dforce.Contract.DforceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dforce *DforceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dforce.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dforce *DforceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dforce.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dforce *DforceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dforce.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Dforce *DforceCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Dforce *DforceSession) Admin() (common.Address, error) {
	return _Dforce.Contract.Admin(&_Dforce.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Dforce *DforceCallerSession) Admin() (common.Address, error) {
	return _Dforce.Contract.Admin(&_Dforce.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Dforce *DforceCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Dforce *DforceSession) Implementation() (common.Address, error) {
	return _Dforce.Contract.Implementation(&_Dforce.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Dforce *DforceCallerSession) Implementation() (common.Address, error) {
	return _Dforce.Contract.Implementation(&_Dforce.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Dforce *DforceCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Dforce *DforceSession) PendingAdmin() (common.Address, error) {
	return _Dforce.Contract.PendingAdmin(&_Dforce.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Dforce *DforceCallerSession) PendingAdmin() (common.Address, error) {
	return _Dforce.Contract.PendingAdmin(&_Dforce.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Dforce *DforceTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Dforce *DforceSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.ChangeAdmin(&_Dforce.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Dforce *DforceTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.ChangeAdmin(&_Dforce.TransactOpts, newAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xd3b2f598.
//
// Solidity: function updateAdmin() returns()
func (_Dforce *DforceTransactor) UpdateAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "updateAdmin")
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xd3b2f598.
//
// Solidity: function updateAdmin() returns()
func (_Dforce *DforceSession) UpdateAdmin() (*types.Transaction, error) {
	return _Dforce.Contract.UpdateAdmin(&_Dforce.TransactOpts)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xd3b2f598.
//
// Solidity: function updateAdmin() returns()
func (_Dforce *DforceTransactorSession) UpdateAdmin() (*types.Transaction, error) {
	return _Dforce.Contract.UpdateAdmin(&_Dforce.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Dforce *DforceTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Dforce *DforceSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.UpgradeTo(&_Dforce.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Dforce *DforceTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.UpgradeTo(&_Dforce.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Dforce *DforceTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Dforce *DforceSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Dforce.Contract.UpgradeToAndCall(&_Dforce.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Dforce *DforceTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Dforce.Contract.UpgradeToAndCall(&_Dforce.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Dforce *DforceTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Dforce.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Dforce *DforceSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Dforce.Contract.Fallback(&_Dforce.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Dforce *DforceTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Dforce.Contract.Fallback(&_Dforce.TransactOpts, calldata)
}

// DforceAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Dforce contract.
type DforceAdminChangedIterator struct {
	Event *DforceAdminChanged // Event containing the contract specifics and raw log

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
func (it *DforceAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceAdminChanged)
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
		it.Event = new(DforceAdminChanged)
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
func (it *DforceAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceAdminChanged represents a AdminChanged event raised by the Dforce contract.
type DforceAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Dforce *DforceFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*DforceAdminChangedIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &DforceAdminChangedIterator{contract: _Dforce.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Dforce *DforceFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *DforceAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceAdminChanged)
				if err := _Dforce.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Dforce *DforceFilterer) ParseAdminChanged(log types.Log) (*DforceAdminChanged, error) {
	event := new(DforceAdminChanged)
	if err := _Dforce.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceAdminUpdatedIterator is returned from FilterAdminUpdated and is used to iterate over the raw logs and unpacked data for AdminUpdated events raised by the Dforce contract.
type DforceAdminUpdatedIterator struct {
	Event *DforceAdminUpdated // Event containing the contract specifics and raw log

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
func (it *DforceAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceAdminUpdated)
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
		it.Event = new(DforceAdminUpdated)
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
func (it *DforceAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceAdminUpdated represents a AdminUpdated event raised by the Dforce contract.
type DforceAdminUpdated struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminUpdated is a free log retrieval operation binding the contract event 0x54e4612788f90384e6843298d7854436f3a585b2c3831ab66abf1de63bfa6c2d.
//
// Solidity: event AdminUpdated(address newAdmin)
func (_Dforce *DforceFilterer) FilterAdminUpdated(opts *bind.FilterOpts) (*DforceAdminUpdatedIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "AdminUpdated")
	if err != nil {
		return nil, err
	}
	return &DforceAdminUpdatedIterator{contract: _Dforce.contract, event: "AdminUpdated", logs: logs, sub: sub}, nil
}

// WatchAdminUpdated is a free log subscription operation binding the contract event 0x54e4612788f90384e6843298d7854436f3a585b2c3831ab66abf1de63bfa6c2d.
//
// Solidity: event AdminUpdated(address newAdmin)
func (_Dforce *DforceFilterer) WatchAdminUpdated(opts *bind.WatchOpts, sink chan<- *DforceAdminUpdated) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "AdminUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceAdminUpdated)
				if err := _Dforce.contract.UnpackLog(event, "AdminUpdated", log); err != nil {
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

// ParseAdminUpdated is a log parse operation binding the contract event 0x54e4612788f90384e6843298d7854436f3a585b2c3831ab66abf1de63bfa6c2d.
//
// Solidity: event AdminUpdated(address newAdmin)
func (_Dforce *DforceFilterer) ParseAdminUpdated(log types.Log) (*DforceAdminUpdated, error) {
	event := new(DforceAdminUpdated)
	if err := _Dforce.contract.UnpackLog(event, "AdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Dforce contract.
type DforceUpgradedIterator struct {
	Event *DforceUpgraded // Event containing the contract specifics and raw log

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
func (it *DforceUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceUpgraded)
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
		it.Event = new(DforceUpgraded)
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
func (it *DforceUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceUpgraded represents a Upgraded event raised by the Dforce contract.
type DforceUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Dforce *DforceFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*DforceUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &DforceUpgradedIterator{contract: _Dforce.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Dforce *DforceFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *DforceUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceUpgraded)
				if err := _Dforce.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Dforce *DforceFilterer) ParseUpgraded(log types.Log) (*DforceUpgraded, error) {
	event := new(DforceUpgraded)
	if err := _Dforce.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
