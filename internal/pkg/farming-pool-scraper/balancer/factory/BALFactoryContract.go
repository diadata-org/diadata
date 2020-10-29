// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balfactory

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

// BALFactoryContractABI is the input ABI used to generate the binding from.
const BALFactoryContractABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blabs\",\"type\":\"address\"}],\"name\":\"LOG_BLABS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"LOG_NEW_POOL\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractBPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBLabs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"isBPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newBPool\",\"outputs\":[{\"internalType\":\"contractBPool\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"setBLabs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BALFactoryContract is an auto generated Go binding around an Ethereum contract.
type BALFactoryContract struct {
	BALFactoryContractCaller     // Read-only binding to the contract
	BALFactoryContractTransactor // Write-only binding to the contract
	BALFactoryContractFilterer   // Log filterer for contract events
}

// BALFactoryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BALFactoryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BALFactoryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BALFactoryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BALFactoryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BALFactoryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BALFactoryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BALFactoryContractSession struct {
	Contract     *BALFactoryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BALFactoryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BALFactoryContractCallerSession struct {
	Contract *BALFactoryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BALFactoryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BALFactoryContractTransactorSession struct {
	Contract     *BALFactoryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BALFactoryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BALFactoryContractRaw struct {
	Contract *BALFactoryContract // Generic contract binding to access the raw methods on
}

// BALFactoryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BALFactoryContractCallerRaw struct {
	Contract *BALFactoryContractCaller // Generic read-only contract binding to access the raw methods on
}

// BALFactoryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BALFactoryContractTransactorRaw struct {
	Contract *BALFactoryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBALFactoryContract creates a new instance of BALFactoryContract, bound to a specific deployed contract.
func NewBALFactoryContract(address common.Address, backend bind.ContractBackend) (*BALFactoryContract, error) {
	contract, err := bindBALFactoryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BALFactoryContract{BALFactoryContractCaller: BALFactoryContractCaller{contract: contract}, BALFactoryContractTransactor: BALFactoryContractTransactor{contract: contract}, BALFactoryContractFilterer: BALFactoryContractFilterer{contract: contract}}, nil
}

// NewBALFactoryContractCaller creates a new read-only instance of BALFactoryContract, bound to a specific deployed contract.
func NewBALFactoryContractCaller(address common.Address, caller bind.ContractCaller) (*BALFactoryContractCaller, error) {
	contract, err := bindBALFactoryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BALFactoryContractCaller{contract: contract}, nil
}

// NewBALFactoryContractTransactor creates a new write-only instance of BALFactoryContract, bound to a specific deployed contract.
func NewBALFactoryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*BALFactoryContractTransactor, error) {
	contract, err := bindBALFactoryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BALFactoryContractTransactor{contract: contract}, nil
}

// NewBALFactoryContractFilterer creates a new log filterer instance of BALFactoryContract, bound to a specific deployed contract.
func NewBALFactoryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*BALFactoryContractFilterer, error) {
	contract, err := bindBALFactoryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BALFactoryContractFilterer{contract: contract}, nil
}

// bindBALFactoryContract binds a generic wrapper to an already deployed contract.
func bindBALFactoryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BALFactoryContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BALFactoryContract *BALFactoryContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BALFactoryContract.Contract.BALFactoryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BALFactoryContract *BALFactoryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BALFactoryContract.Contract.BALFactoryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BALFactoryContract *BALFactoryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BALFactoryContract.Contract.BALFactoryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BALFactoryContract *BALFactoryContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BALFactoryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BALFactoryContract *BALFactoryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BALFactoryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BALFactoryContract *BALFactoryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BALFactoryContract.Contract.contract.Transact(opts, method, params...)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_BALFactoryContract *BALFactoryContractCaller) GetBLabs(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BALFactoryContract.contract.Call(opts, out, "getBLabs")
	return *ret0, err
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_BALFactoryContract *BALFactoryContractSession) GetBLabs() (common.Address, error) {
	return _BALFactoryContract.Contract.GetBLabs(&_BALFactoryContract.CallOpts)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_BALFactoryContract *BALFactoryContractCallerSession) GetBLabs() (common.Address, error) {
	return _BALFactoryContract.Contract.GetBLabs(&_BALFactoryContract.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BALFactoryContract *BALFactoryContractCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BALFactoryContract.contract.Call(opts, out, "getColor")
	return *ret0, err
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BALFactoryContract *BALFactoryContractSession) GetColor() ([32]byte, error) {
	return _BALFactoryContract.Contract.GetColor(&_BALFactoryContract.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BALFactoryContract *BALFactoryContractCallerSession) GetColor() ([32]byte, error) {
	return _BALFactoryContract.Contract.GetColor(&_BALFactoryContract.CallOpts)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_BALFactoryContract *BALFactoryContractCaller) IsBPool(opts *bind.CallOpts, b common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BALFactoryContract.contract.Call(opts, out, "isBPool", b)
	return *ret0, err
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_BALFactoryContract *BALFactoryContractSession) IsBPool(b common.Address) (bool, error) {
	return _BALFactoryContract.Contract.IsBPool(&_BALFactoryContract.CallOpts, b)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_BALFactoryContract *BALFactoryContractCallerSession) IsBPool(b common.Address) (bool, error) {
	return _BALFactoryContract.Contract.IsBPool(&_BALFactoryContract.CallOpts, b)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_BALFactoryContract *BALFactoryContractTransactor) Collect(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _BALFactoryContract.contract.Transact(opts, "collect", pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_BALFactoryContract *BALFactoryContractSession) Collect(pool common.Address) (*types.Transaction, error) {
	return _BALFactoryContract.Contract.Collect(&_BALFactoryContract.TransactOpts, pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_BALFactoryContract *BALFactoryContractTransactorSession) Collect(pool common.Address) (*types.Transaction, error) {
	return _BALFactoryContract.Contract.Collect(&_BALFactoryContract.TransactOpts, pool)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_BALFactoryContract *BALFactoryContractTransactor) NewBPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BALFactoryContract.contract.Transact(opts, "newBPool")
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_BALFactoryContract *BALFactoryContractSession) NewBPool() (*types.Transaction, error) {
	return _BALFactoryContract.Contract.NewBPool(&_BALFactoryContract.TransactOpts)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_BALFactoryContract *BALFactoryContractTransactorSession) NewBPool() (*types.Transaction, error) {
	return _BALFactoryContract.Contract.NewBPool(&_BALFactoryContract.TransactOpts)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_BALFactoryContract *BALFactoryContractTransactor) SetBLabs(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _BALFactoryContract.contract.Transact(opts, "setBLabs", b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_BALFactoryContract *BALFactoryContractSession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _BALFactoryContract.Contract.SetBLabs(&_BALFactoryContract.TransactOpts, b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_BALFactoryContract *BALFactoryContractTransactorSession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _BALFactoryContract.Contract.SetBLabs(&_BALFactoryContract.TransactOpts, b)
}

// BALFactoryContractLOGBLABSIterator is returned from FilterLOGBLABS and is used to iterate over the raw logs and unpacked data for LOGBLABS events raised by the BALFactoryContract contract.
type BALFactoryContractLOGBLABSIterator struct {
	Event *BALFactoryContractLOGBLABS // Event containing the contract specifics and raw log

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
func (it *BALFactoryContractLOGBLABSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BALFactoryContractLOGBLABS)
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
		it.Event = new(BALFactoryContractLOGBLABS)
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
func (it *BALFactoryContractLOGBLABSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BALFactoryContractLOGBLABSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BALFactoryContractLOGBLABS represents a LOGBLABS event raised by the BALFactoryContract contract.
type BALFactoryContractLOGBLABS struct {
	Caller common.Address
	Blabs  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGBLABS is a free log retrieval operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_BALFactoryContract *BALFactoryContractFilterer) FilterLOGBLABS(opts *bind.FilterOpts, caller []common.Address, blabs []common.Address) (*BALFactoryContractLOGBLABSIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _BALFactoryContract.contract.FilterLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return &BALFactoryContractLOGBLABSIterator{contract: _BALFactoryContract.contract, event: "LOG_BLABS", logs: logs, sub: sub}, nil
}

// WatchLOGBLABS is a free log subscription operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_BALFactoryContract *BALFactoryContractFilterer) WatchLOGBLABS(opts *bind.WatchOpts, sink chan<- *BALFactoryContractLOGBLABS, caller []common.Address, blabs []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _BALFactoryContract.contract.WatchLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BALFactoryContractLOGBLABS)
				if err := _BALFactoryContract.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
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

// ParseLOGBLABS is a log parse operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_BALFactoryContract *BALFactoryContractFilterer) ParseLOGBLABS(log types.Log) (*BALFactoryContractLOGBLABS, error) {
	event := new(BALFactoryContractLOGBLABS)
	if err := _BALFactoryContract.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BALFactoryContractLOGNEWPOOLIterator is returned from FilterLOGNEWPOOL and is used to iterate over the raw logs and unpacked data for LOGNEWPOOL events raised by the BALFactoryContract contract.
type BALFactoryContractLOGNEWPOOLIterator struct {
	Event *BALFactoryContractLOGNEWPOOL // Event containing the contract specifics and raw log

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
func (it *BALFactoryContractLOGNEWPOOLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BALFactoryContractLOGNEWPOOL)
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
		it.Event = new(BALFactoryContractLOGNEWPOOL)
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
func (it *BALFactoryContractLOGNEWPOOLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BALFactoryContractLOGNEWPOOLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BALFactoryContractLOGNEWPOOL represents a LOGNEWPOOL event raised by the BALFactoryContract contract.
type BALFactoryContractLOGNEWPOOL struct {
	Caller common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGNEWPOOL is a free log retrieval operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_BALFactoryContract *BALFactoryContractFilterer) FilterLOGNEWPOOL(opts *bind.FilterOpts, caller []common.Address, pool []common.Address) (*BALFactoryContractLOGNEWPOOLIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BALFactoryContract.contract.FilterLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &BALFactoryContractLOGNEWPOOLIterator{contract: _BALFactoryContract.contract, event: "LOG_NEW_POOL", logs: logs, sub: sub}, nil
}

// WatchLOGNEWPOOL is a free log subscription operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_BALFactoryContract *BALFactoryContractFilterer) WatchLOGNEWPOOL(opts *bind.WatchOpts, sink chan<- *BALFactoryContractLOGNEWPOOL, caller []common.Address, pool []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BALFactoryContract.contract.WatchLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BALFactoryContractLOGNEWPOOL)
				if err := _BALFactoryContract.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
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

// ParseLOGNEWPOOL is a log parse operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_BALFactoryContract *BALFactoryContractFilterer) ParseLOGNEWPOOL(log types.Log) (*BALFactoryContractLOGNEWPOOL, error) {
	event := new(BALFactoryContractLOGNEWPOOL)
	if err := _BALFactoryContract.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
		return nil, err
	}
	return event, nil
}
