// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancerfactory

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

// BalancerfactoryABI is the input ABI used to generate the binding from.
const BalancerfactoryABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blabs\",\"type\":\"address\"}],\"name\":\"LOG_BLABS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"LOG_NEW_POOL\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractBPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBLabs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"isBPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newBPool\",\"outputs\":[{\"internalType\":\"contractBPool\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"setBLabs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Balancerfactory is an auto generated Go binding around an Ethereum contract.
type Balancerfactory struct {
	BalancerfactoryCaller     // Read-only binding to the contract
	BalancerfactoryTransactor // Write-only binding to the contract
	BalancerfactoryFilterer   // Log filterer for contract events
}

// BalancerfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerfactorySession struct {
	Contract     *Balancerfactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalancerfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerfactoryCallerSession struct {
	Contract *BalancerfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BalancerfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerfactoryTransactorSession struct {
	Contract     *BalancerfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BalancerfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerfactoryRaw struct {
	Contract *Balancerfactory // Generic contract binding to access the raw methods on
}

// BalancerfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerfactoryCallerRaw struct {
	Contract *BalancerfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerfactoryTransactorRaw struct {
	Contract *BalancerfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerfactory creates a new instance of Balancerfactory, bound to a specific deployed contract.
func NewBalancerfactory(address common.Address, backend bind.ContractBackend) (*Balancerfactory, error) {
	contract, err := bindBalancerfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balancerfactory{BalancerfactoryCaller: BalancerfactoryCaller{contract: contract}, BalancerfactoryTransactor: BalancerfactoryTransactor{contract: contract}, BalancerfactoryFilterer: BalancerfactoryFilterer{contract: contract}}, nil
}

// NewBalancerfactoryCaller creates a new read-only instance of Balancerfactory, bound to a specific deployed contract.
func NewBalancerfactoryCaller(address common.Address, caller bind.ContractCaller) (*BalancerfactoryCaller, error) {
	contract, err := bindBalancerfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerfactoryCaller{contract: contract}, nil
}

// NewBalancerfactoryTransactor creates a new write-only instance of Balancerfactory, bound to a specific deployed contract.
func NewBalancerfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerfactoryTransactor, error) {
	contract, err := bindBalancerfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerfactoryTransactor{contract: contract}, nil
}

// NewBalancerfactoryFilterer creates a new log filterer instance of Balancerfactory, bound to a specific deployed contract.
func NewBalancerfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerfactoryFilterer, error) {
	contract, err := bindBalancerfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerfactoryFilterer{contract: contract}, nil
}

// bindBalancerfactory binds a generic wrapper to an already deployed contract.
func bindBalancerfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancerfactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancerfactory *BalancerfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancerfactory.Contract.BalancerfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancerfactory *BalancerfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerfactory.Contract.BalancerfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancerfactory *BalancerfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancerfactory.Contract.BalancerfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancerfactory *BalancerfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancerfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancerfactory *BalancerfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancerfactory *BalancerfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancerfactory.Contract.contract.Transact(opts, method, params...)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_Balancerfactory *BalancerfactoryCaller) GetBLabs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balancerfactory.contract.Call(opts, &out, "getBLabs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_Balancerfactory *BalancerfactorySession) GetBLabs() (common.Address, error) {
	return _Balancerfactory.Contract.GetBLabs(&_Balancerfactory.CallOpts)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_Balancerfactory *BalancerfactoryCallerSession) GetBLabs() (common.Address, error) {
	return _Balancerfactory.Contract.GetBLabs(&_Balancerfactory.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_Balancerfactory *BalancerfactoryCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Balancerfactory.contract.Call(opts, &out, "getColor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_Balancerfactory *BalancerfactorySession) GetColor() ([32]byte, error) {
	return _Balancerfactory.Contract.GetColor(&_Balancerfactory.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_Balancerfactory *BalancerfactoryCallerSession) GetColor() ([32]byte, error) {
	return _Balancerfactory.Contract.GetColor(&_Balancerfactory.CallOpts)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_Balancerfactory *BalancerfactoryCaller) IsBPool(opts *bind.CallOpts, b common.Address) (bool, error) {
	var out []interface{}
	err := _Balancerfactory.contract.Call(opts, &out, "isBPool", b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_Balancerfactory *BalancerfactorySession) IsBPool(b common.Address) (bool, error) {
	return _Balancerfactory.Contract.IsBPool(&_Balancerfactory.CallOpts, b)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_Balancerfactory *BalancerfactoryCallerSession) IsBPool(b common.Address) (bool, error) {
	return _Balancerfactory.Contract.IsBPool(&_Balancerfactory.CallOpts, b)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_Balancerfactory *BalancerfactoryTransactor) Collect(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _Balancerfactory.contract.Transact(opts, "collect", pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_Balancerfactory *BalancerfactorySession) Collect(pool common.Address) (*types.Transaction, error) {
	return _Balancerfactory.Contract.Collect(&_Balancerfactory.TransactOpts, pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_Balancerfactory *BalancerfactoryTransactorSession) Collect(pool common.Address) (*types.Transaction, error) {
	return _Balancerfactory.Contract.Collect(&_Balancerfactory.TransactOpts, pool)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_Balancerfactory *BalancerfactoryTransactor) NewBPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerfactory.contract.Transact(opts, "newBPool")
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_Balancerfactory *BalancerfactorySession) NewBPool() (*types.Transaction, error) {
	return _Balancerfactory.Contract.NewBPool(&_Balancerfactory.TransactOpts)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_Balancerfactory *BalancerfactoryTransactorSession) NewBPool() (*types.Transaction, error) {
	return _Balancerfactory.Contract.NewBPool(&_Balancerfactory.TransactOpts)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_Balancerfactory *BalancerfactoryTransactor) SetBLabs(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _Balancerfactory.contract.Transact(opts, "setBLabs", b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_Balancerfactory *BalancerfactorySession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _Balancerfactory.Contract.SetBLabs(&_Balancerfactory.TransactOpts, b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_Balancerfactory *BalancerfactoryTransactorSession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _Balancerfactory.Contract.SetBLabs(&_Balancerfactory.TransactOpts, b)
}

// BalancerfactoryLOGBLABSIterator is returned from FilterLOGBLABS and is used to iterate over the raw logs and unpacked data for LOGBLABS events raised by the Balancerfactory contract.
type BalancerfactoryLOGBLABSIterator struct {
	Event *BalancerfactoryLOGBLABS // Event containing the contract specifics and raw log

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
func (it *BalancerfactoryLOGBLABSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerfactoryLOGBLABS)
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
		it.Event = new(BalancerfactoryLOGBLABS)
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
func (it *BalancerfactoryLOGBLABSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerfactoryLOGBLABSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerfactoryLOGBLABS represents a LOGBLABS event raised by the Balancerfactory contract.
type BalancerfactoryLOGBLABS struct {
	Caller common.Address
	Blabs  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGBLABS is a free log retrieval operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_Balancerfactory *BalancerfactoryFilterer) FilterLOGBLABS(opts *bind.FilterOpts, caller []common.Address, blabs []common.Address) (*BalancerfactoryLOGBLABSIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _Balancerfactory.contract.FilterLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return &BalancerfactoryLOGBLABSIterator{contract: _Balancerfactory.contract, event: "LOG_BLABS", logs: logs, sub: sub}, nil
}

// WatchLOGBLABS is a free log subscription operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_Balancerfactory *BalancerfactoryFilterer) WatchLOGBLABS(opts *bind.WatchOpts, sink chan<- *BalancerfactoryLOGBLABS, caller []common.Address, blabs []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _Balancerfactory.contract.WatchLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerfactoryLOGBLABS)
				if err := _Balancerfactory.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
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
func (_Balancerfactory *BalancerfactoryFilterer) ParseLOGBLABS(log types.Log) (*BalancerfactoryLOGBLABS, error) {
	event := new(BalancerfactoryLOGBLABS)
	if err := _Balancerfactory.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerfactoryLOGNEWPOOLIterator is returned from FilterLOGNEWPOOL and is used to iterate over the raw logs and unpacked data for LOGNEWPOOL events raised by the Balancerfactory contract.
type BalancerfactoryLOGNEWPOOLIterator struct {
	Event *BalancerfactoryLOGNEWPOOL // Event containing the contract specifics and raw log

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
func (it *BalancerfactoryLOGNEWPOOLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerfactoryLOGNEWPOOL)
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
		it.Event = new(BalancerfactoryLOGNEWPOOL)
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
func (it *BalancerfactoryLOGNEWPOOLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerfactoryLOGNEWPOOLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerfactoryLOGNEWPOOL represents a LOGNEWPOOL event raised by the Balancerfactory contract.
type BalancerfactoryLOGNEWPOOL struct {
	Caller common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGNEWPOOL is a free log retrieval operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_Balancerfactory *BalancerfactoryFilterer) FilterLOGNEWPOOL(opts *bind.FilterOpts, caller []common.Address, pool []common.Address) (*BalancerfactoryLOGNEWPOOLIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Balancerfactory.contract.FilterLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &BalancerfactoryLOGNEWPOOLIterator{contract: _Balancerfactory.contract, event: "LOG_NEW_POOL", logs: logs, sub: sub}, nil
}

// WatchLOGNEWPOOL is a free log subscription operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_Balancerfactory *BalancerfactoryFilterer) WatchLOGNEWPOOL(opts *bind.WatchOpts, sink chan<- *BalancerfactoryLOGNEWPOOL, caller []common.Address, pool []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Balancerfactory.contract.WatchLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerfactoryLOGNEWPOOL)
				if err := _Balancerfactory.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
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
func (_Balancerfactory *BalancerfactoryFilterer) ParseLOGNEWPOOL(log types.Log) (*BalancerfactoryLOGNEWPOOL, error) {
	event := new(BalancerfactoryLOGNEWPOOL)
	if err := _Balancerfactory.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
