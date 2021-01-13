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

// BalfactoryABI is the input ABI used to generate the binding from.
const BalfactoryABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blabs\",\"type\":\"address\"}],\"name\":\"LOG_BLABS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"LOG_NEW_POOL\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractBPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBLabs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"isBPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newBPool\",\"outputs\":[{\"internalType\":\"contractBPool\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"setBLabs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Balfactory is an auto generated Go binding around an Ethereum contract.
type Balfactory struct {
	BalfactoryCaller     // Read-only binding to the contract
	BalfactoryTransactor // Write-only binding to the contract
	BalfactoryFilterer   // Log filterer for contract events
}

// BalfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalfactorySession struct {
	Contract     *Balfactory       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalfactoryCallerSession struct {
	Contract *BalfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BalfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalfactoryTransactorSession struct {
	Contract     *BalfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BalfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalfactoryRaw struct {
	Contract *Balfactory // Generic contract binding to access the raw methods on
}

// BalfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalfactoryCallerRaw struct {
	Contract *BalfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// BalfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalfactoryTransactorRaw struct {
	Contract *BalfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalfactory creates a new instance of Balfactory, bound to a specific deployed contract.
func NewBalfactory(address common.Address, backend bind.ContractBackend) (*Balfactory, error) {
	contract, err := bindBalfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balfactory{BalfactoryCaller: BalfactoryCaller{contract: contract}, BalfactoryTransactor: BalfactoryTransactor{contract: contract}, BalfactoryFilterer: BalfactoryFilterer{contract: contract}}, nil
}

// NewBalfactoryCaller creates a new read-only instance of Balfactory, bound to a specific deployed contract.
func NewBalfactoryCaller(address common.Address, caller bind.ContractCaller) (*BalfactoryCaller, error) {
	contract, err := bindBalfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalfactoryCaller{contract: contract}, nil
}

// NewBalfactoryTransactor creates a new write-only instance of Balfactory, bound to a specific deployed contract.
func NewBalfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*BalfactoryTransactor, error) {
	contract, err := bindBalfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalfactoryTransactor{contract: contract}, nil
}

// NewBalfactoryFilterer creates a new log filterer instance of Balfactory, bound to a specific deployed contract.
func NewBalfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*BalfactoryFilterer, error) {
	contract, err := bindBalfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalfactoryFilterer{contract: contract}, nil
}

// bindBalfactory binds a generic wrapper to an already deployed contract.
func bindBalfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalfactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balfactory *BalfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balfactory.Contract.BalfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balfactory *BalfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balfactory.Contract.BalfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balfactory *BalfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balfactory.Contract.BalfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balfactory *BalfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balfactory *BalfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balfactory *BalfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balfactory.Contract.contract.Transact(opts, method, params...)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_Balfactory *BalfactoryCaller) GetBLabs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balfactory.contract.Call(opts, &out, "getBLabs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_Balfactory *BalfactorySession) GetBLabs() (common.Address, error) {
	return _Balfactory.Contract.GetBLabs(&_Balfactory.CallOpts)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_Balfactory *BalfactoryCallerSession) GetBLabs() (common.Address, error) {
	return _Balfactory.Contract.GetBLabs(&_Balfactory.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_Balfactory *BalfactoryCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Balfactory.contract.Call(opts, &out, "getColor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_Balfactory *BalfactorySession) GetColor() ([32]byte, error) {
	return _Balfactory.Contract.GetColor(&_Balfactory.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_Balfactory *BalfactoryCallerSession) GetColor() ([32]byte, error) {
	return _Balfactory.Contract.GetColor(&_Balfactory.CallOpts)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_Balfactory *BalfactoryCaller) IsBPool(opts *bind.CallOpts, b common.Address) (bool, error) {
	var out []interface{}
	err := _Balfactory.contract.Call(opts, &out, "isBPool", b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_Balfactory *BalfactorySession) IsBPool(b common.Address) (bool, error) {
	return _Balfactory.Contract.IsBPool(&_Balfactory.CallOpts, b)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_Balfactory *BalfactoryCallerSession) IsBPool(b common.Address) (bool, error) {
	return _Balfactory.Contract.IsBPool(&_Balfactory.CallOpts, b)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_Balfactory *BalfactoryTransactor) Collect(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _Balfactory.contract.Transact(opts, "collect", pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_Balfactory *BalfactorySession) Collect(pool common.Address) (*types.Transaction, error) {
	return _Balfactory.Contract.Collect(&_Balfactory.TransactOpts, pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_Balfactory *BalfactoryTransactorSession) Collect(pool common.Address) (*types.Transaction, error) {
	return _Balfactory.Contract.Collect(&_Balfactory.TransactOpts, pool)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_Balfactory *BalfactoryTransactor) NewBPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balfactory.contract.Transact(opts, "newBPool")
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_Balfactory *BalfactorySession) NewBPool() (*types.Transaction, error) {
	return _Balfactory.Contract.NewBPool(&_Balfactory.TransactOpts)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_Balfactory *BalfactoryTransactorSession) NewBPool() (*types.Transaction, error) {
	return _Balfactory.Contract.NewBPool(&_Balfactory.TransactOpts)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_Balfactory *BalfactoryTransactor) SetBLabs(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _Balfactory.contract.Transact(opts, "setBLabs", b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_Balfactory *BalfactorySession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _Balfactory.Contract.SetBLabs(&_Balfactory.TransactOpts, b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_Balfactory *BalfactoryTransactorSession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _Balfactory.Contract.SetBLabs(&_Balfactory.TransactOpts, b)
}

// BalfactoryLOGBLABSIterator is returned from FilterLOGBLABS and is used to iterate over the raw logs and unpacked data for LOGBLABS events raised by the Balfactory contract.
type BalfactoryLOGBLABSIterator struct {
	Event *BalfactoryLOGBLABS // Event containing the contract specifics and raw log

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
func (it *BalfactoryLOGBLABSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalfactoryLOGBLABS)
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
		it.Event = new(BalfactoryLOGBLABS)
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
func (it *BalfactoryLOGBLABSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalfactoryLOGBLABSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalfactoryLOGBLABS represents a LOGBLABS event raised by the Balfactory contract.
type BalfactoryLOGBLABS struct {
	Caller common.Address
	Blabs  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGBLABS is a free log retrieval operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_Balfactory *BalfactoryFilterer) FilterLOGBLABS(opts *bind.FilterOpts, caller []common.Address, blabs []common.Address) (*BalfactoryLOGBLABSIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _Balfactory.contract.FilterLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return &BalfactoryLOGBLABSIterator{contract: _Balfactory.contract, event: "LOG_BLABS", logs: logs, sub: sub}, nil
}

// WatchLOGBLABS is a free log subscription operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_Balfactory *BalfactoryFilterer) WatchLOGBLABS(opts *bind.WatchOpts, sink chan<- *BalfactoryLOGBLABS, caller []common.Address, blabs []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _Balfactory.contract.WatchLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalfactoryLOGBLABS)
				if err := _Balfactory.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
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
func (_Balfactory *BalfactoryFilterer) ParseLOGBLABS(log types.Log) (*BalfactoryLOGBLABS, error) {
	event := new(BalfactoryLOGBLABS)
	if err := _Balfactory.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalfactoryLOGNEWPOOLIterator is returned from FilterLOGNEWPOOL and is used to iterate over the raw logs and unpacked data for LOGNEWPOOL events raised by the Balfactory contract.
type BalfactoryLOGNEWPOOLIterator struct {
	Event *BalfactoryLOGNEWPOOL // Event containing the contract specifics and raw log

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
func (it *BalfactoryLOGNEWPOOLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalfactoryLOGNEWPOOL)
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
		it.Event = new(BalfactoryLOGNEWPOOL)
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
func (it *BalfactoryLOGNEWPOOLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalfactoryLOGNEWPOOLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalfactoryLOGNEWPOOL represents a LOGNEWPOOL event raised by the Balfactory contract.
type BalfactoryLOGNEWPOOL struct {
	Caller common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGNEWPOOL is a free log retrieval operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_Balfactory *BalfactoryFilterer) FilterLOGNEWPOOL(opts *bind.FilterOpts, caller []common.Address, pool []common.Address) (*BalfactoryLOGNEWPOOLIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Balfactory.contract.FilterLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &BalfactoryLOGNEWPOOLIterator{contract: _Balfactory.contract, event: "LOG_NEW_POOL", logs: logs, sub: sub}, nil
}

// WatchLOGNEWPOOL is a free log subscription operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_Balfactory *BalfactoryFilterer) WatchLOGNEWPOOL(opts *bind.WatchOpts, sink chan<- *BalfactoryLOGNEWPOOL, caller []common.Address, pool []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Balfactory.contract.WatchLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalfactoryLOGNEWPOOL)
				if err := _Balfactory.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
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
func (_Balfactory *BalfactoryFilterer) ParseLOGNEWPOOL(log types.Log) (*BalfactoryLOGNEWPOOL, error) {
	event := new(BalfactoryLOGNEWPOOL)
	if err := _Balfactory.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
		return nil, err
	}
	return event, nil
}
