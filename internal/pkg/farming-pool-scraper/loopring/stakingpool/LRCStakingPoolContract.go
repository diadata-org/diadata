// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lrcstakingpool

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

// LRCStakingPoolContractABI is the input ABI used to generate the binding from.
const LRCStakingPoolContractABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blabs\",\"type\":\"address\"}],\"name\":\"LOG_BLABS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"LOG_NEW_POOL\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractBPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBLabs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"isBPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newBPool\",\"outputs\":[{\"internalType\":\"contractBPool\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"setBLabs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LRCStakingPoolContract is an auto generated Go binding around an Ethereum contract.
type LRCStakingPoolContract struct {
	LRCStakingPoolContractCaller     // Read-only binding to the contract
	LRCStakingPoolContractTransactor // Write-only binding to the contract
	LRCStakingPoolContractFilterer   // Log filterer for contract events
}

// LRCStakingPoolContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type LRCStakingPoolContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LRCStakingPoolContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LRCStakingPoolContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LRCStakingPoolContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LRCStakingPoolContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LRCStakingPoolContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LRCStakingPoolContractSession struct {
	Contract     *LRCStakingPoolContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LRCStakingPoolContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LRCStakingPoolContractCallerSession struct {
	Contract *LRCStakingPoolContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// LRCStakingPoolContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LRCStakingPoolContractTransactorSession struct {
	Contract     *LRCStakingPoolContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// LRCStakingPoolContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type LRCStakingPoolContractRaw struct {
	Contract *LRCStakingPoolContract // Generic contract binding to access the raw methods on
}

// LRCStakingPoolContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LRCStakingPoolContractCallerRaw struct {
	Contract *LRCStakingPoolContractCaller // Generic read-only contract binding to access the raw methods on
}

// LRCStakingPoolContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LRCStakingPoolContractTransactorRaw struct {
	Contract *LRCStakingPoolContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLRCStakingPoolContract creates a new instance of LRCStakingPoolContract, bound to a specific deployed contract.
func NewLRCStakingPoolContract(address common.Address, backend bind.ContractBackend) (*LRCStakingPoolContract, error) {
	contract, err := bindLRCStakingPoolContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LRCStakingPoolContract{LRCStakingPoolContractCaller: LRCStakingPoolContractCaller{contract: contract}, LRCStakingPoolContractTransactor: LRCStakingPoolContractTransactor{contract: contract}, LRCStakingPoolContractFilterer: LRCStakingPoolContractFilterer{contract: contract}}, nil
}

// NewLRCStakingPoolContractCaller creates a new read-only instance of LRCStakingPoolContract, bound to a specific deployed contract.
func NewLRCStakingPoolContractCaller(address common.Address, caller bind.ContractCaller) (*LRCStakingPoolContractCaller, error) {
	contract, err := bindLRCStakingPoolContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LRCStakingPoolContractCaller{contract: contract}, nil
}

// NewLRCStakingPoolContractTransactor creates a new write-only instance of LRCStakingPoolContract, bound to a specific deployed contract.
func NewLRCStakingPoolContractTransactor(address common.Address, transactor bind.ContractTransactor) (*LRCStakingPoolContractTransactor, error) {
	contract, err := bindLRCStakingPoolContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LRCStakingPoolContractTransactor{contract: contract}, nil
}

// NewLRCStakingPoolContractFilterer creates a new log filterer instance of LRCStakingPoolContract, bound to a specific deployed contract.
func NewLRCStakingPoolContractFilterer(address common.Address, filterer bind.ContractFilterer) (*LRCStakingPoolContractFilterer, error) {
	contract, err := bindLRCStakingPoolContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LRCStakingPoolContractFilterer{contract: contract}, nil
}

// bindLRCStakingPoolContract binds a generic wrapper to an already deployed contract.
func bindLRCStakingPoolContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LRCStakingPoolContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LRCStakingPoolContract *LRCStakingPoolContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LRCStakingPoolContract.Contract.LRCStakingPoolContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LRCStakingPoolContract *LRCStakingPoolContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LRCStakingPoolContract.Contract.LRCStakingPoolContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LRCStakingPoolContract *LRCStakingPoolContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LRCStakingPoolContract.Contract.LRCStakingPoolContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LRCStakingPoolContract *LRCStakingPoolContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LRCStakingPoolContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LRCStakingPoolContract *LRCStakingPoolContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LRCStakingPoolContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LRCStakingPoolContract *LRCStakingPoolContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LRCStakingPoolContract.Contract.contract.Transact(opts, method, params...)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_LRCStakingPoolContract *LRCStakingPoolContractCaller) GetBLabs(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LRCStakingPoolContract.contract.Call(opts, out, "getBLabs")
	return *ret0, err
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_LRCStakingPoolContract *LRCStakingPoolContractSession) GetBLabs() (common.Address, error) {
	return _LRCStakingPoolContract.Contract.GetBLabs(&_LRCStakingPoolContract.CallOpts)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_LRCStakingPoolContract *LRCStakingPoolContractCallerSession) GetBLabs() (common.Address, error) {
	return _LRCStakingPoolContract.Contract.GetBLabs(&_LRCStakingPoolContract.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_LRCStakingPoolContract *LRCStakingPoolContractCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LRCStakingPoolContract.contract.Call(opts, out, "getColor")
	return *ret0, err
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_LRCStakingPoolContract *LRCStakingPoolContractSession) GetColor() ([32]byte, error) {
	return _LRCStakingPoolContract.Contract.GetColor(&_LRCStakingPoolContract.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_LRCStakingPoolContract *LRCStakingPoolContractCallerSession) GetColor() ([32]byte, error) {
	return _LRCStakingPoolContract.Contract.GetColor(&_LRCStakingPoolContract.CallOpts)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_LRCStakingPoolContract *LRCStakingPoolContractCaller) IsBPool(opts *bind.CallOpts, b common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LRCStakingPoolContract.contract.Call(opts, out, "isBPool", b)
	return *ret0, err
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_LRCStakingPoolContract *LRCStakingPoolContractSession) IsBPool(b common.Address) (bool, error) {
	return _LRCStakingPoolContract.Contract.IsBPool(&_LRCStakingPoolContract.CallOpts, b)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_LRCStakingPoolContract *LRCStakingPoolContractCallerSession) IsBPool(b common.Address) (bool, error) {
	return _LRCStakingPoolContract.Contract.IsBPool(&_LRCStakingPoolContract.CallOpts, b)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_LRCStakingPoolContract *LRCStakingPoolContractTransactor) Collect(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _LRCStakingPoolContract.contract.Transact(opts, "collect", pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_LRCStakingPoolContract *LRCStakingPoolContractSession) Collect(pool common.Address) (*types.Transaction, error) {
	return _LRCStakingPoolContract.Contract.Collect(&_LRCStakingPoolContract.TransactOpts, pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_LRCStakingPoolContract *LRCStakingPoolContractTransactorSession) Collect(pool common.Address) (*types.Transaction, error) {
	return _LRCStakingPoolContract.Contract.Collect(&_LRCStakingPoolContract.TransactOpts, pool)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_LRCStakingPoolContract *LRCStakingPoolContractTransactor) NewBPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LRCStakingPoolContract.contract.Transact(opts, "newBPool")
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_LRCStakingPoolContract *LRCStakingPoolContractSession) NewBPool() (*types.Transaction, error) {
	return _LRCStakingPoolContract.Contract.NewBPool(&_LRCStakingPoolContract.TransactOpts)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_LRCStakingPoolContract *LRCStakingPoolContractTransactorSession) NewBPool() (*types.Transaction, error) {
	return _LRCStakingPoolContract.Contract.NewBPool(&_LRCStakingPoolContract.TransactOpts)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_LRCStakingPoolContract *LRCStakingPoolContractTransactor) SetBLabs(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _LRCStakingPoolContract.contract.Transact(opts, "setBLabs", b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_LRCStakingPoolContract *LRCStakingPoolContractSession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _LRCStakingPoolContract.Contract.SetBLabs(&_LRCStakingPoolContract.TransactOpts, b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_LRCStakingPoolContract *LRCStakingPoolContractTransactorSession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _LRCStakingPoolContract.Contract.SetBLabs(&_LRCStakingPoolContract.TransactOpts, b)
}

// LRCStakingPoolContractLOGBLABSIterator is returned from FilterLOGBLABS and is used to iterate over the raw logs and unpacked data for LOGBLABS events raised by the LRCStakingPoolContract contract.
type LRCStakingPoolContractLOGBLABSIterator struct {
	Event *LRCStakingPoolContractLOGBLABS // Event containing the contract specifics and raw log

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
func (it *LRCStakingPoolContractLOGBLABSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCStakingPoolContractLOGBLABS)
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
		it.Event = new(LRCStakingPoolContractLOGBLABS)
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
func (it *LRCStakingPoolContractLOGBLABSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCStakingPoolContractLOGBLABSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCStakingPoolContractLOGBLABS represents a LOGBLABS event raised by the LRCStakingPoolContract contract.
type LRCStakingPoolContractLOGBLABS struct {
	Caller common.Address
	Blabs  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGBLABS is a free log retrieval operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_LRCStakingPoolContract *LRCStakingPoolContractFilterer) FilterLOGBLABS(opts *bind.FilterOpts, caller []common.Address, blabs []common.Address) (*LRCStakingPoolContractLOGBLABSIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _LRCStakingPoolContract.contract.FilterLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return &LRCStakingPoolContractLOGBLABSIterator{contract: _LRCStakingPoolContract.contract, event: "LOG_BLABS", logs: logs, sub: sub}, nil
}

// WatchLOGBLABS is a free log subscription operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_LRCStakingPoolContract *LRCStakingPoolContractFilterer) WatchLOGBLABS(opts *bind.WatchOpts, sink chan<- *LRCStakingPoolContractLOGBLABS, caller []common.Address, blabs []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _LRCStakingPoolContract.contract.WatchLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCStakingPoolContractLOGBLABS)
				if err := _LRCStakingPoolContract.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
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
func (_LRCStakingPoolContract *LRCStakingPoolContractFilterer) ParseLOGBLABS(log types.Log) (*LRCStakingPoolContractLOGBLABS, error) {
	event := new(LRCStakingPoolContractLOGBLABS)
	if err := _LRCStakingPoolContract.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LRCStakingPoolContractLOGNEWPOOLIterator is returned from FilterLOGNEWPOOL and is used to iterate over the raw logs and unpacked data for LOGNEWPOOL events raised by the LRCStakingPoolContract contract.
type LRCStakingPoolContractLOGNEWPOOLIterator struct {
	Event *LRCStakingPoolContractLOGNEWPOOL // Event containing the contract specifics and raw log

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
func (it *LRCStakingPoolContractLOGNEWPOOLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCStakingPoolContractLOGNEWPOOL)
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
		it.Event = new(LRCStakingPoolContractLOGNEWPOOL)
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
func (it *LRCStakingPoolContractLOGNEWPOOLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCStakingPoolContractLOGNEWPOOLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCStakingPoolContractLOGNEWPOOL represents a LOGNEWPOOL event raised by the LRCStakingPoolContract contract.
type LRCStakingPoolContractLOGNEWPOOL struct {
	Caller common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGNEWPOOL is a free log retrieval operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_LRCStakingPoolContract *LRCStakingPoolContractFilterer) FilterLOGNEWPOOL(opts *bind.FilterOpts, caller []common.Address, pool []common.Address) (*LRCStakingPoolContractLOGNEWPOOLIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _LRCStakingPoolContract.contract.FilterLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &LRCStakingPoolContractLOGNEWPOOLIterator{contract: _LRCStakingPoolContract.contract, event: "LOG_NEW_POOL", logs: logs, sub: sub}, nil
}

// WatchLOGNEWPOOL is a free log subscription operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_LRCStakingPoolContract *LRCStakingPoolContractFilterer) WatchLOGNEWPOOL(opts *bind.WatchOpts, sink chan<- *LRCStakingPoolContractLOGNEWPOOL, caller []common.Address, pool []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _LRCStakingPoolContract.contract.WatchLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCStakingPoolContractLOGNEWPOOL)
				if err := _LRCStakingPoolContract.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
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
func (_LRCStakingPoolContract *LRCStakingPoolContractFilterer) ParseLOGNEWPOOL(log types.Log) (*LRCStakingPoolContractLOGNEWPOOL, error) {
	event := new(LRCStakingPoolContractLOGNEWPOOL)
	if err := _LRCStakingPoolContract.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
		return nil, err
	}
	return event, nil
}
