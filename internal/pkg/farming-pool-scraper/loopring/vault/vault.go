// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feevault

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

// LRCFeeVaultContractABI is the input ABI used to generate the binding from.
const LRCFeeVaultContractABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blabs\",\"type\":\"address\"}],\"name\":\"LOG_BLABS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"LOG_NEW_POOL\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractBPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBLabs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"isBPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newBPool\",\"outputs\":[{\"internalType\":\"contractBPool\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"setBLabs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LRCFeeVaultContract is an auto generated Go binding around an Ethereum contract.
type LRCFeeVaultContract struct {
	LRCFeeVaultContractCaller     // Read-only binding to the contract
	LRCFeeVaultContractTransactor // Write-only binding to the contract
	LRCFeeVaultContractFilterer   // Log filterer for contract events
}

// LRCFeeVaultContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type LRCFeeVaultContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LRCFeeVaultContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LRCFeeVaultContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LRCFeeVaultContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LRCFeeVaultContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LRCFeeVaultContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LRCFeeVaultContractSession struct {
	Contract     *LRCFeeVaultContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LRCFeeVaultContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LRCFeeVaultContractCallerSession struct {
	Contract *LRCFeeVaultContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// LRCFeeVaultContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LRCFeeVaultContractTransactorSession struct {
	Contract     *LRCFeeVaultContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// LRCFeeVaultContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type LRCFeeVaultContractRaw struct {
	Contract *LRCFeeVaultContract // Generic contract binding to access the raw methods on
}

// LRCFeeVaultContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LRCFeeVaultContractCallerRaw struct {
	Contract *LRCFeeVaultContractCaller // Generic read-only contract binding to access the raw methods on
}

// LRCFeeVaultContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LRCFeeVaultContractTransactorRaw struct {
	Contract *LRCFeeVaultContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLRCFeeVaultContract creates a new instance of LRCFeeVaultContract, bound to a specific deployed contract.
func NewLRCFeeVaultContract(address common.Address, backend bind.ContractBackend) (*LRCFeeVaultContract, error) {
	contract, err := bindLRCFeeVaultContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LRCFeeVaultContract{LRCFeeVaultContractCaller: LRCFeeVaultContractCaller{contract: contract}, LRCFeeVaultContractTransactor: LRCFeeVaultContractTransactor{contract: contract}, LRCFeeVaultContractFilterer: LRCFeeVaultContractFilterer{contract: contract}}, nil
}

// NewLRCFeeVaultContractCaller creates a new read-only instance of LRCFeeVaultContract, bound to a specific deployed contract.
func NewLRCFeeVaultContractCaller(address common.Address, caller bind.ContractCaller) (*LRCFeeVaultContractCaller, error) {
	contract, err := bindLRCFeeVaultContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LRCFeeVaultContractCaller{contract: contract}, nil
}

// NewLRCFeeVaultContractTransactor creates a new write-only instance of LRCFeeVaultContract, bound to a specific deployed contract.
func NewLRCFeeVaultContractTransactor(address common.Address, transactor bind.ContractTransactor) (*LRCFeeVaultContractTransactor, error) {
	contract, err := bindLRCFeeVaultContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LRCFeeVaultContractTransactor{contract: contract}, nil
}

// NewLRCFeeVaultContractFilterer creates a new log filterer instance of LRCFeeVaultContract, bound to a specific deployed contract.
func NewLRCFeeVaultContractFilterer(address common.Address, filterer bind.ContractFilterer) (*LRCFeeVaultContractFilterer, error) {
	contract, err := bindLRCFeeVaultContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LRCFeeVaultContractFilterer{contract: contract}, nil
}

// bindLRCFeeVaultContract binds a generic wrapper to an already deployed contract.
func bindLRCFeeVaultContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LRCFeeVaultContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LRCFeeVaultContract *LRCFeeVaultContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LRCFeeVaultContract.Contract.LRCFeeVaultContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LRCFeeVaultContract *LRCFeeVaultContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LRCFeeVaultContract.Contract.LRCFeeVaultContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LRCFeeVaultContract *LRCFeeVaultContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LRCFeeVaultContract.Contract.LRCFeeVaultContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LRCFeeVaultContract *LRCFeeVaultContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LRCFeeVaultContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LRCFeeVaultContract *LRCFeeVaultContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LRCFeeVaultContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LRCFeeVaultContract *LRCFeeVaultContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LRCFeeVaultContract.Contract.contract.Transact(opts, method, params...)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_LRCFeeVaultContract *LRCFeeVaultContractCaller) GetBLabs(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LRCFeeVaultContract.contract.Call(opts, out, "getBLabs")
	return *ret0, err
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_LRCFeeVaultContract *LRCFeeVaultContractSession) GetBLabs() (common.Address, error) {
	return _LRCFeeVaultContract.Contract.GetBLabs(&_LRCFeeVaultContract.CallOpts)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_LRCFeeVaultContract *LRCFeeVaultContractCallerSession) GetBLabs() (common.Address, error) {
	return _LRCFeeVaultContract.Contract.GetBLabs(&_LRCFeeVaultContract.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_LRCFeeVaultContract *LRCFeeVaultContractCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LRCFeeVaultContract.contract.Call(opts, out, "getColor")
	return *ret0, err
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_LRCFeeVaultContract *LRCFeeVaultContractSession) GetColor() ([32]byte, error) {
	return _LRCFeeVaultContract.Contract.GetColor(&_LRCFeeVaultContract.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_LRCFeeVaultContract *LRCFeeVaultContractCallerSession) GetColor() ([32]byte, error) {
	return _LRCFeeVaultContract.Contract.GetColor(&_LRCFeeVaultContract.CallOpts)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_LRCFeeVaultContract *LRCFeeVaultContractCaller) IsBPool(opts *bind.CallOpts, b common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LRCFeeVaultContract.contract.Call(opts, out, "isBPool", b)
	return *ret0, err
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_LRCFeeVaultContract *LRCFeeVaultContractSession) IsBPool(b common.Address) (bool, error) {
	return _LRCFeeVaultContract.Contract.IsBPool(&_LRCFeeVaultContract.CallOpts, b)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_LRCFeeVaultContract *LRCFeeVaultContractCallerSession) IsBPool(b common.Address) (bool, error) {
	return _LRCFeeVaultContract.Contract.IsBPool(&_LRCFeeVaultContract.CallOpts, b)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_LRCFeeVaultContract *LRCFeeVaultContractTransactor) Collect(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _LRCFeeVaultContract.contract.Transact(opts, "collect", pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_LRCFeeVaultContract *LRCFeeVaultContractSession) Collect(pool common.Address) (*types.Transaction, error) {
	return _LRCFeeVaultContract.Contract.Collect(&_LRCFeeVaultContract.TransactOpts, pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_LRCFeeVaultContract *LRCFeeVaultContractTransactorSession) Collect(pool common.Address) (*types.Transaction, error) {
	return _LRCFeeVaultContract.Contract.Collect(&_LRCFeeVaultContract.TransactOpts, pool)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_LRCFeeVaultContract *LRCFeeVaultContractTransactor) NewBPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LRCFeeVaultContract.contract.Transact(opts, "newBPool")
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_LRCFeeVaultContract *LRCFeeVaultContractSession) NewBPool() (*types.Transaction, error) {
	return _LRCFeeVaultContract.Contract.NewBPool(&_LRCFeeVaultContract.TransactOpts)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_LRCFeeVaultContract *LRCFeeVaultContractTransactorSession) NewBPool() (*types.Transaction, error) {
	return _LRCFeeVaultContract.Contract.NewBPool(&_LRCFeeVaultContract.TransactOpts)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_LRCFeeVaultContract *LRCFeeVaultContractTransactor) SetBLabs(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _LRCFeeVaultContract.contract.Transact(opts, "setBLabs", b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_LRCFeeVaultContract *LRCFeeVaultContractSession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _LRCFeeVaultContract.Contract.SetBLabs(&_LRCFeeVaultContract.TransactOpts, b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_LRCFeeVaultContract *LRCFeeVaultContractTransactorSession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _LRCFeeVaultContract.Contract.SetBLabs(&_LRCFeeVaultContract.TransactOpts, b)
}

// LRCFeeVaultContractLOGBLABSIterator is returned from FilterLOGBLABS and is used to iterate over the raw logs and unpacked data for LOGBLABS events raised by the LRCFeeVaultContract contract.
type LRCFeeVaultContractLOGBLABSIterator struct {
	Event *LRCFeeVaultContractLOGBLABS // Event containing the contract specifics and raw log

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
func (it *LRCFeeVaultContractLOGBLABSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCFeeVaultContractLOGBLABS)
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
		it.Event = new(LRCFeeVaultContractLOGBLABS)
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
func (it *LRCFeeVaultContractLOGBLABSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCFeeVaultContractLOGBLABSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCFeeVaultContractLOGBLABS represents a LOGBLABS event raised by the LRCFeeVaultContract contract.
type LRCFeeVaultContractLOGBLABS struct {
	Caller common.Address
	Blabs  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGBLABS is a free log retrieval operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_LRCFeeVaultContract *LRCFeeVaultContractFilterer) FilterLOGBLABS(opts *bind.FilterOpts, caller []common.Address, blabs []common.Address) (*LRCFeeVaultContractLOGBLABSIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _LRCFeeVaultContract.contract.FilterLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return &LRCFeeVaultContractLOGBLABSIterator{contract: _LRCFeeVaultContract.contract, event: "LOG_BLABS", logs: logs, sub: sub}, nil
}

// WatchLOGBLABS is a free log subscription operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_LRCFeeVaultContract *LRCFeeVaultContractFilterer) WatchLOGBLABS(opts *bind.WatchOpts, sink chan<- *LRCFeeVaultContractLOGBLABS, caller []common.Address, blabs []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _LRCFeeVaultContract.contract.WatchLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCFeeVaultContractLOGBLABS)
				if err := _LRCFeeVaultContract.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
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
func (_LRCFeeVaultContract *LRCFeeVaultContractFilterer) ParseLOGBLABS(log types.Log) (*LRCFeeVaultContractLOGBLABS, error) {
	event := new(LRCFeeVaultContractLOGBLABS)
	if err := _LRCFeeVaultContract.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LRCFeeVaultContractLOGNEWPOOLIterator is returned from FilterLOGNEWPOOL and is used to iterate over the raw logs and unpacked data for LOGNEWPOOL events raised by the LRCFeeVaultContract contract.
type LRCFeeVaultContractLOGNEWPOOLIterator struct {
	Event *LRCFeeVaultContractLOGNEWPOOL // Event containing the contract specifics and raw log

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
func (it *LRCFeeVaultContractLOGNEWPOOLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCFeeVaultContractLOGNEWPOOL)
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
		it.Event = new(LRCFeeVaultContractLOGNEWPOOL)
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
func (it *LRCFeeVaultContractLOGNEWPOOLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCFeeVaultContractLOGNEWPOOLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCFeeVaultContractLOGNEWPOOL represents a LOGNEWPOOL event raised by the LRCFeeVaultContract contract.
type LRCFeeVaultContractLOGNEWPOOL struct {
	Caller common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGNEWPOOL is a free log retrieval operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_LRCFeeVaultContract *LRCFeeVaultContractFilterer) FilterLOGNEWPOOL(opts *bind.FilterOpts, caller []common.Address, pool []common.Address) (*LRCFeeVaultContractLOGNEWPOOLIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _LRCFeeVaultContract.contract.FilterLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &LRCFeeVaultContractLOGNEWPOOLIterator{contract: _LRCFeeVaultContract.contract, event: "LOG_NEW_POOL", logs: logs, sub: sub}, nil
}

// WatchLOGNEWPOOL is a free log subscription operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_LRCFeeVaultContract *LRCFeeVaultContractFilterer) WatchLOGNEWPOOL(opts *bind.WatchOpts, sink chan<- *LRCFeeVaultContractLOGNEWPOOL, caller []common.Address, pool []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _LRCFeeVaultContract.contract.WatchLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCFeeVaultContractLOGNEWPOOL)
				if err := _LRCFeeVaultContract.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
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
func (_LRCFeeVaultContract *LRCFeeVaultContractFilterer) ParseLOGNEWPOOL(log types.Log) (*LRCFeeVaultContractLOGNEWPOOL, error) {
	event := new(LRCFeeVaultContractLOGNEWPOOL)
	if err := _LRCFeeVaultContract.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
		return nil, err
	}
	return event, nil
}
