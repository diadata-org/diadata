// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ammswap

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

// LRCAmmSwapContractABI is the input ABI used to generate the binding from.
const LRCAmmSwapContractABI = [{"inputs":[{"internalType":"address","name":"_exchange","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"address","name":"token","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"Drained","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"user","type":"address"},{"indexed":true,"internalType":"bytes4","name":"selector","type":"bytes4"},{"indexed":false,"internalType":"bool","name":"allowed","type":"bool"}],"name":"PermissionUpdate","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bool","name":"open","type":"bool"}],"name":"SubmitBlocksAccessOpened","type":"event"},{"stateMutability":"payable","type":"fallback"},{"inputs":[{"internalType":"address","name":"drainer","type":"address"}],"name":"canDrain","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"claimOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"address","name":"token","type":"address"}],"name":"drain","outputs":[{"internalType":"uint256","name":"amount","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"user","type":"address"},{"internalType":"bytes4","name":"selector","type":"bytes4"},{"internalType":"bool","name":"granted","type":"bool"}],"name":"grantAccess","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"user","type":"address"},{"internalType":"bytes4","name":"selector","type":"bytes4"}],"name":"hasAccessTo","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"signHash","type":"bytes32"},{"internalType":"bytes","name":"signature","type":"bytes"}],"name":"isValidSignature","outputs":[{"internalType":"bytes4","name":"","type":"bytes4"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"open","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bool","name":"_open","type":"bool"}],"name":"openAccessToSubmitBlocks","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"pendingOwner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"},{"internalType":"bytes4","name":"","type":"bytes4"}],"name":"permissions","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bool","name":"isDataCompressed","type":"bool"},{"internalType":"bytes","name":"data","type":"bytes"},{"components":[{"components":[{"internalType":"uint16","name":"blockIdx","type":"uint16"},{"components":[{"internalType":"uint16","name":"txIdx","type":"uint16"},{"internalType":"uint16","name":"receiverIdx","type":"uint16"},{"internalType":"bytes","name":"data","type":"bytes"}],"internalType":"struct LoopringIOExchangeOwner.TxCallback[]","name":"txCallbacks","type":"tuple[]"}],"internalType":"struct LoopringIOExchangeOwner.BlockCallback[]","name":"blockCallbacks","type":"tuple[]"},{"internalType":"address[]","name":"receivers","type":"address[]"}],"internalType":"struct LoopringIOExchangeOwner.CallbackConfig","name":"callbackConfig","type":"tuple"}],"name":"submitBlocksWithCallbacks","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"target","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes","name":"data","type":"bytes"}],"name":"transact","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"stateMutability":"payable","type":"receive"}]

// LRCAmmSwapContract is an auto generated Go binding around an Ethereum contract.
type LRCAmmSwapContract struct {
	LRCAmmSwapContractCaller     // Read-only binding to the contract
	LRCAmmSwapContractTransactor // Write-only binding to the contract
	LRCAmmSwapContractFilterer   // Log filterer for contract events
}

// LRCAmmSwapContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type LRCAmmSwapContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LRCAmmSwapContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LRCAmmSwapContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LRCAmmSwapContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LRCAmmSwapContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LRCAmmSwapContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LRCAmmSwapContractSession struct {
	Contract     *LRCAmmSwapContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LRCAmmSwapContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LRCAmmSwapContractCallerSession struct {
	Contract *LRCAmmSwapContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// LRCAmmSwapContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LRCAmmSwapContractTransactorSession struct {
	Contract     *LRCAmmSwapContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// LRCAmmSwapContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type LRCAmmSwapContractRaw struct {
	Contract *LRCAmmSwapContract // Generic contract binding to access the raw methods on
}

// LRCAmmSwapContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LRCAmmSwapContractCallerRaw struct {
	Contract *LRCAmmSwapContractCaller // Generic read-only contract binding to access the raw methods on
}

// LRCAmmSwapContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LRCAmmSwapContractTransactorRaw struct {
	Contract *LRCAmmSwapContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLRCAmmSwapContract creates a new instance of LRCAmmSwapContract, bound to a specific deployed contract.
func NewLRCAmmSwapContract(address common.Address, backend bind.ContractBackend) (*LRCAmmSwapContract, error) {
	contract, err := bindLRCAmmSwapContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LRCAmmSwapContract{LRCAmmSwapContractCaller: LRCAmmSwapContractCaller{contract: contract}, LRCAmmSwapContractTransactor: LRCAmmSwapContractTransactor{contract: contract}, LRCAmmSwapContractFilterer: LRCAmmSwapContractFilterer{contract: contract}}, nil
}

// NewLRCAmmSwapContractCaller creates a new read-only instance of LRCAmmSwapContract, bound to a specific deployed contract.
func NewLRCAmmSwapContractCaller(address common.Address, caller bind.ContractCaller) (*LRCAmmSwapContractCaller, error) {
	contract, err := bindLRCAmmSwapContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LRCAmmSwapContractCaller{contract: contract}, nil
}

// NewLRCAmmSwapContractTransactor creates a new write-only instance of LRCAmmSwapContract, bound to a specific deployed contract.
func NewLRCAmmSwapContractTransactor(address common.Address, transactor bind.ContractTransactor) (*LRCAmmSwapContractTransactor, error) {
	contract, err := bindLRCAmmSwapContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LRCAmmSwapContractTransactor{contract: contract}, nil
}

// NewLRCAmmSwapContractFilterer creates a new log filterer instance of LRCAmmSwapContract, bound to a specific deployed contract.
func NewLRCAmmSwapContractFilterer(address common.Address, filterer bind.ContractFilterer) (*LRCAmmSwapContractFilterer, error) {
	contract, err := bindLRCAmmSwapContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LRCAmmSwapContractFilterer{contract: contract}, nil
}

// bindLRCAmmSwapContract binds a generic wrapper to an already deployed contract.
func bindLRCAmmSwapContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LRCAmmSwapContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LRCAmmSwapContract *LRCAmmSwapContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LRCAmmSwapContract.Contract.LRCAmmSwapContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LRCAmmSwapContract *LRCAmmSwapContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LRCAmmSwapContract.Contract.LRCAmmSwapContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LRCAmmSwapContract *LRCAmmSwapContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LRCAmmSwapContract.Contract.LRCAmmSwapContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LRCAmmSwapContract *LRCAmmSwapContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LRCAmmSwapContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LRCAmmSwapContract *LRCAmmSwapContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LRCAmmSwapContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LRCAmmSwapContract *LRCAmmSwapContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LRCAmmSwapContract.Contract.contract.Transact(opts, method, params...)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_LRCAmmSwapContract *LRCAmmSwapContractCaller) GetBLabs(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LRCAmmSwapContract.contract.Call(opts, out, "getBLabs")
	return *ret0, err
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_LRCAmmSwapContract *LRCAmmSwapContractSession) GetBLabs() (common.Address, error) {
	return _LRCAmmSwapContract.Contract.GetBLabs(&_LRCAmmSwapContract.CallOpts)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_LRCAmmSwapContract *LRCAmmSwapContractCallerSession) GetBLabs() (common.Address, error) {
	return _LRCAmmSwapContract.Contract.GetBLabs(&_LRCAmmSwapContract.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_LRCAmmSwapContract *LRCAmmSwapContractCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LRCAmmSwapContract.contract.Call(opts, out, "getColor")
	return *ret0, err
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_LRCAmmSwapContract *LRCAmmSwapContractSession) GetColor() ([32]byte, error) {
	return _LRCAmmSwapContract.Contract.GetColor(&_LRCAmmSwapContract.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_LRCAmmSwapContract *LRCAmmSwapContractCallerSession) GetColor() ([32]byte, error) {
	return _LRCAmmSwapContract.Contract.GetColor(&_LRCAmmSwapContract.CallOpts)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_LRCAmmSwapContract *LRCAmmSwapContractCaller) IsBPool(opts *bind.CallOpts, b common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LRCAmmSwapContract.contract.Call(opts, out, "isBPool", b)
	return *ret0, err
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_LRCAmmSwapContract *LRCAmmSwapContractSession) IsBPool(b common.Address) (bool, error) {
	return _LRCAmmSwapContract.Contract.IsBPool(&_LRCAmmSwapContract.CallOpts, b)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_LRCAmmSwapContract *LRCAmmSwapContractCallerSession) IsBPool(b common.Address) (bool, error) {
	return _LRCAmmSwapContract.Contract.IsBPool(&_LRCAmmSwapContract.CallOpts, b)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_LRCAmmSwapContract *LRCAmmSwapContractTransactor) Collect(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _LRCAmmSwapContract.contract.Transact(opts, "collect", pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_LRCAmmSwapContract *LRCAmmSwapContractSession) Collect(pool common.Address) (*types.Transaction, error) {
	return _LRCAmmSwapContract.Contract.Collect(&_LRCAmmSwapContract.TransactOpts, pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_LRCAmmSwapContract *LRCAmmSwapContractTransactorSession) Collect(pool common.Address) (*types.Transaction, error) {
	return _LRCAmmSwapContract.Contract.Collect(&_LRCAmmSwapContract.TransactOpts, pool)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_LRCAmmSwapContract *LRCAmmSwapContractTransactor) NewBPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LRCAmmSwapContract.contract.Transact(opts, "newBPool")
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_LRCAmmSwapContract *LRCAmmSwapContractSession) NewBPool() (*types.Transaction, error) {
	return _LRCAmmSwapContract.Contract.NewBPool(&_LRCAmmSwapContract.TransactOpts)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_LRCAmmSwapContract *LRCAmmSwapContractTransactorSession) NewBPool() (*types.Transaction, error) {
	return _LRCAmmSwapContract.Contract.NewBPool(&_LRCAmmSwapContract.TransactOpts)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_LRCAmmSwapContract *LRCAmmSwapContractTransactor) SetBLabs(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _LRCAmmSwapContract.contract.Transact(opts, "setBLabs", b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_LRCAmmSwapContract *LRCAmmSwapContractSession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _LRCAmmSwapContract.Contract.SetBLabs(&_LRCAmmSwapContract.TransactOpts, b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_LRCAmmSwapContract *LRCAmmSwapContractTransactorSession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _LRCAmmSwapContract.Contract.SetBLabs(&_LRCAmmSwapContract.TransactOpts, b)
}

// LRCAmmSwapContractLOGBLABSIterator is returned from FilterLOGBLABS and is used to iterate over the raw logs and unpacked data for LOGBLABS events raised by the LRCAmmSwapContract contract.
type LRCAmmSwapContractLOGBLABSIterator struct {
	Event *LRCAmmSwapContractLOGBLABS // Event containing the contract specifics and raw log

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
func (it *LRCAmmSwapContractLOGBLABSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCAmmSwapContractLOGBLABS)
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
		it.Event = new(LRCAmmSwapContractLOGBLABS)
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
func (it *LRCAmmSwapContractLOGBLABSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCAmmSwapContractLOGBLABSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCAmmSwapContractLOGBLABS represents a LOGBLABS event raised by the LRCAmmSwapContract contract.
type LRCAmmSwapContractLOGBLABS struct {
	Caller common.Address
	Blabs  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGBLABS is a free log retrieval operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_LRCAmmSwapContract *LRCAmmSwapContractFilterer) FilterLOGBLABS(opts *bind.FilterOpts, caller []common.Address, blabs []common.Address) (*LRCAmmSwapContractLOGBLABSIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _LRCAmmSwapContract.contract.FilterLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return &LRCAmmSwapContractLOGBLABSIterator{contract: _LRCAmmSwapContract.contract, event: "LOG_BLABS", logs: logs, sub: sub}, nil
}

// WatchLOGBLABS is a free log subscription operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_LRCAmmSwapContract *LRCAmmSwapContractFilterer) WatchLOGBLABS(opts *bind.WatchOpts, sink chan<- *LRCAmmSwapContractLOGBLABS, caller []common.Address, blabs []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _LRCAmmSwapContract.contract.WatchLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCAmmSwapContractLOGBLABS)
				if err := _LRCAmmSwapContract.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
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
func (_LRCAmmSwapContract *LRCAmmSwapContractFilterer) ParseLOGBLABS(log types.Log) (*LRCAmmSwapContractLOGBLABS, error) {
	event := new(LRCAmmSwapContractLOGBLABS)
	if err := _LRCAmmSwapContract.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LRCAmmSwapContractLOGNEWPOOLIterator is returned from FilterLOGNEWPOOL and is used to iterate over the raw logs and unpacked data for LOGNEWPOOL events raised by the LRCAmmSwapContract contract.
type LRCAmmSwapContractLOGNEWPOOLIterator struct {
	Event *LRCAmmSwapContractLOGNEWPOOL // Event containing the contract specifics and raw log

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
func (it *LRCAmmSwapContractLOGNEWPOOLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCAmmSwapContractLOGNEWPOOL)
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
		it.Event = new(LRCAmmSwapContractLOGNEWPOOL)
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
func (it *LRCAmmSwapContractLOGNEWPOOLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCAmmSwapContractLOGNEWPOOLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCAmmSwapContractLOGNEWPOOL represents a LOGNEWPOOL event raised by the LRCAmmSwapContract contract.
type LRCAmmSwapContractLOGNEWPOOL struct {
	Caller common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGNEWPOOL is a free log retrieval operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_LRCAmmSwapContract *LRCAmmSwapContractFilterer) FilterLOGNEWPOOL(opts *bind.FilterOpts, caller []common.Address, pool []common.Address) (*LRCAmmSwapContractLOGNEWPOOLIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _LRCAmmSwapContract.contract.FilterLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &LRCAmmSwapContractLOGNEWPOOLIterator{contract: _LRCAmmSwapContract.contract, event: "LOG_NEW_POOL", logs: logs, sub: sub}, nil
}

// WatchLOGNEWPOOL is a free log subscription operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_LRCAmmSwapContract *LRCAmmSwapContractFilterer) WatchLOGNEWPOOL(opts *bind.WatchOpts, sink chan<- *LRCAmmSwapContractLOGNEWPOOL, caller []common.Address, pool []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _LRCAmmSwapContract.contract.WatchLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCAmmSwapContractLOGNEWPOOL)
				if err := _LRCAmmSwapContract.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
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
func (_LRCAmmSwapContract *LRCAmmSwapContractFilterer) ParseLOGNEWPOOL(log types.Log) (*LRCAmmSwapContractLOGNEWPOOL, error) {
	event := new(LRCAmmSwapContractLOGNEWPOOL)
	if err := _LRCAmmSwapContract.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
		return nil, err
	}
	return event, nil
}
