// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancerpool

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

// BalancerpoolABI is the input ABI used to generate the binding from.
const BalancerpoolABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blabs\",\"type\":\"address\"}],\"name\":\"LOG_BLABS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"LOG_NEW_POOL\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractBPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBLabs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"isBPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newBPool\",\"outputs\":[{\"internalType\":\"contractBPool\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"setBLabs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Balancerpool is an auto generated Go binding around an Ethereum contract.
type Balancerpool struct {
	BalancerpoolCaller     // Read-only binding to the contract
	BalancerpoolTransactor // Write-only binding to the contract
	BalancerpoolFilterer   // Log filterer for contract events
}

// BalancerpoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerpoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerpoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerpoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerpoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerpoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerpoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerpoolSession struct {
	Contract     *Balancerpool     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalancerpoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerpoolCallerSession struct {
	Contract *BalancerpoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BalancerpoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerpoolTransactorSession struct {
	Contract     *BalancerpoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BalancerpoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerpoolRaw struct {
	Contract *Balancerpool // Generic contract binding to access the raw methods on
}

// BalancerpoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerpoolCallerRaw struct {
	Contract *BalancerpoolCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerpoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerpoolTransactorRaw struct {
	Contract *BalancerpoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerpool creates a new instance of Balancerpool, bound to a specific deployed contract.
func NewBalancerpool(address common.Address, backend bind.ContractBackend) (*Balancerpool, error) {
	contract, err := bindBalancerpool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balancerpool{BalancerpoolCaller: BalancerpoolCaller{contract: contract}, BalancerpoolTransactor: BalancerpoolTransactor{contract: contract}, BalancerpoolFilterer: BalancerpoolFilterer{contract: contract}}, nil
}

// NewBalancerpoolCaller creates a new read-only instance of Balancerpool, bound to a specific deployed contract.
func NewBalancerpoolCaller(address common.Address, caller bind.ContractCaller) (*BalancerpoolCaller, error) {
	contract, err := bindBalancerpool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerpoolCaller{contract: contract}, nil
}

// NewBalancerpoolTransactor creates a new write-only instance of Balancerpool, bound to a specific deployed contract.
func NewBalancerpoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerpoolTransactor, error) {
	contract, err := bindBalancerpool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerpoolTransactor{contract: contract}, nil
}

// NewBalancerpoolFilterer creates a new log filterer instance of Balancerpool, bound to a specific deployed contract.
func NewBalancerpoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerpoolFilterer, error) {
	contract, err := bindBalancerpool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerpoolFilterer{contract: contract}, nil
}

// bindBalancerpool binds a generic wrapper to an already deployed contract.
func bindBalancerpool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancerpoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancerpool *BalancerpoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancerpool.Contract.BalancerpoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancerpool *BalancerpoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerpool.Contract.BalancerpoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancerpool *BalancerpoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancerpool.Contract.BalancerpoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancerpool *BalancerpoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancerpool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancerpool *BalancerpoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerpool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancerpool *BalancerpoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancerpool.Contract.contract.Transact(opts, method, params...)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_Balancerpool *BalancerpoolCaller) GetBLabs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balancerpool.contract.Call(opts, &out, "getBLabs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_Balancerpool *BalancerpoolSession) GetBLabs() (common.Address, error) {
	return _Balancerpool.Contract.GetBLabs(&_Balancerpool.CallOpts)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_Balancerpool *BalancerpoolCallerSession) GetBLabs() (common.Address, error) {
	return _Balancerpool.Contract.GetBLabs(&_Balancerpool.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_Balancerpool *BalancerpoolCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Balancerpool.contract.Call(opts, &out, "getColor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_Balancerpool *BalancerpoolSession) GetColor() ([32]byte, error) {
	return _Balancerpool.Contract.GetColor(&_Balancerpool.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_Balancerpool *BalancerpoolCallerSession) GetColor() ([32]byte, error) {
	return _Balancerpool.Contract.GetColor(&_Balancerpool.CallOpts)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_Balancerpool *BalancerpoolCaller) IsBPool(opts *bind.CallOpts, b common.Address) (bool, error) {
	var out []interface{}
	err := _Balancerpool.contract.Call(opts, &out, "isBPool", b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_Balancerpool *BalancerpoolSession) IsBPool(b common.Address) (bool, error) {
	return _Balancerpool.Contract.IsBPool(&_Balancerpool.CallOpts, b)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_Balancerpool *BalancerpoolCallerSession) IsBPool(b common.Address) (bool, error) {
	return _Balancerpool.Contract.IsBPool(&_Balancerpool.CallOpts, b)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_Balancerpool *BalancerpoolTransactor) Collect(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _Balancerpool.contract.Transact(opts, "collect", pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_Balancerpool *BalancerpoolSession) Collect(pool common.Address) (*types.Transaction, error) {
	return _Balancerpool.Contract.Collect(&_Balancerpool.TransactOpts, pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_Balancerpool *BalancerpoolTransactorSession) Collect(pool common.Address) (*types.Transaction, error) {
	return _Balancerpool.Contract.Collect(&_Balancerpool.TransactOpts, pool)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_Balancerpool *BalancerpoolTransactor) NewBPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerpool.contract.Transact(opts, "newBPool")
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_Balancerpool *BalancerpoolSession) NewBPool() (*types.Transaction, error) {
	return _Balancerpool.Contract.NewBPool(&_Balancerpool.TransactOpts)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_Balancerpool *BalancerpoolTransactorSession) NewBPool() (*types.Transaction, error) {
	return _Balancerpool.Contract.NewBPool(&_Balancerpool.TransactOpts)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_Balancerpool *BalancerpoolTransactor) SetBLabs(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _Balancerpool.contract.Transact(opts, "setBLabs", b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_Balancerpool *BalancerpoolSession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _Balancerpool.Contract.SetBLabs(&_Balancerpool.TransactOpts, b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_Balancerpool *BalancerpoolTransactorSession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _Balancerpool.Contract.SetBLabs(&_Balancerpool.TransactOpts, b)
}

// BalancerpoolLOGBLABSIterator is returned from FilterLOGBLABS and is used to iterate over the raw logs and unpacked data for LOGBLABS events raised by the Balancerpool contract.
type BalancerpoolLOGBLABSIterator struct {
	Event *BalancerpoolLOGBLABS // Event containing the contract specifics and raw log

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
func (it *BalancerpoolLOGBLABSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerpoolLOGBLABS)
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
		it.Event = new(BalancerpoolLOGBLABS)
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
func (it *BalancerpoolLOGBLABSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerpoolLOGBLABSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerpoolLOGBLABS represents a LOGBLABS event raised by the Balancerpool contract.
type BalancerpoolLOGBLABS struct {
	Caller common.Address
	Blabs  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGBLABS is a free log retrieval operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_Balancerpool *BalancerpoolFilterer) FilterLOGBLABS(opts *bind.FilterOpts, caller []common.Address, blabs []common.Address) (*BalancerpoolLOGBLABSIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _Balancerpool.contract.FilterLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return &BalancerpoolLOGBLABSIterator{contract: _Balancerpool.contract, event: "LOG_BLABS", logs: logs, sub: sub}, nil
}

// WatchLOGBLABS is a free log subscription operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_Balancerpool *BalancerpoolFilterer) WatchLOGBLABS(opts *bind.WatchOpts, sink chan<- *BalancerpoolLOGBLABS, caller []common.Address, blabs []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _Balancerpool.contract.WatchLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerpoolLOGBLABS)
				if err := _Balancerpool.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
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
func (_Balancerpool *BalancerpoolFilterer) ParseLOGBLABS(log types.Log) (*BalancerpoolLOGBLABS, error) {
	event := new(BalancerpoolLOGBLABS)
	if err := _Balancerpool.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerpoolLOGNEWPOOLIterator is returned from FilterLOGNEWPOOL and is used to iterate over the raw logs and unpacked data for LOGNEWPOOL events raised by the Balancerpool contract.
type BalancerpoolLOGNEWPOOLIterator struct {
	Event *BalancerpoolLOGNEWPOOL // Event containing the contract specifics and raw log

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
func (it *BalancerpoolLOGNEWPOOLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerpoolLOGNEWPOOL)
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
		it.Event = new(BalancerpoolLOGNEWPOOL)
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
func (it *BalancerpoolLOGNEWPOOLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerpoolLOGNEWPOOLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerpoolLOGNEWPOOL represents a LOGNEWPOOL event raised by the Balancerpool contract.
type BalancerpoolLOGNEWPOOL struct {
	Caller common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGNEWPOOL is a free log retrieval operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_Balancerpool *BalancerpoolFilterer) FilterLOGNEWPOOL(opts *bind.FilterOpts, caller []common.Address, pool []common.Address) (*BalancerpoolLOGNEWPOOLIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Balancerpool.contract.FilterLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &BalancerpoolLOGNEWPOOLIterator{contract: _Balancerpool.contract, event: "LOG_NEW_POOL", logs: logs, sub: sub}, nil
}

// WatchLOGNEWPOOL is a free log subscription operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_Balancerpool *BalancerpoolFilterer) WatchLOGNEWPOOL(opts *bind.WatchOpts, sink chan<- *BalancerpoolLOGNEWPOOL, caller []common.Address, pool []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Balancerpool.contract.WatchLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerpoolLOGNEWPOOL)
				if err := _Balancerpool.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
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
func (_Balancerpool *BalancerpoolFilterer) ParseLOGNEWPOOL(log types.Log) (*BalancerpoolLOGNEWPOOL, error) {
	event := new(BalancerpoolLOGNEWPOOL)
	if err := _Balancerpool.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}