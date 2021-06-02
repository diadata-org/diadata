// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SignedDiaOracle

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

// SignedDiaOracleABI is the input ABI used to generate the binding from.
const SignedDiaOracleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"newPrice\",\"type\":\"uint256\"},{\"name\":\"newSupply\",\"type\":\"uint256\"},{\"name\":\"newTimestamp\",\"type\":\"uint256\"},{\"name\":\"lastSignedData\",\"type\":\"string\"}],\"name\":\"updateCoinInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastUpdateTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastSignedData\",\"type\":\"string\"}],\"name\":\"newCoinInfo\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getCoinInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SignedDiaOracle is an auto generated Go binding around an Ethereum contract.
type SignedDiaOracle struct {
	SignedDiaOracleCaller     // Read-only binding to the contract
	SignedDiaOracleTransactor // Write-only binding to the contract
	SignedDiaOracleFilterer   // Log filterer for contract events
}

// SignedDiaOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignedDiaOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedDiaOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignedDiaOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedDiaOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignedDiaOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedDiaOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignedDiaOracleSession struct {
	Contract     *SignedDiaOracle  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignedDiaOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignedDiaOracleCallerSession struct {
	Contract *SignedDiaOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SignedDiaOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignedDiaOracleTransactorSession struct {
	Contract     *SignedDiaOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SignedDiaOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignedDiaOracleRaw struct {
	Contract *SignedDiaOracle // Generic contract binding to access the raw methods on
}

// SignedDiaOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignedDiaOracleCallerRaw struct {
	Contract *SignedDiaOracleCaller // Generic read-only contract binding to access the raw methods on
}

// SignedDiaOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignedDiaOracleTransactorRaw struct {
	Contract *SignedDiaOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignedDiaOracle creates a new instance of SignedDiaOracle, bound to a specific deployed contract.
func NewSignedDiaOracle(address common.Address, backend bind.ContractBackend) (*SignedDiaOracle, error) {
	contract, err := bindSignedDiaOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignedDiaOracle{SignedDiaOracleCaller: SignedDiaOracleCaller{contract: contract}, SignedDiaOracleTransactor: SignedDiaOracleTransactor{contract: contract}, SignedDiaOracleFilterer: SignedDiaOracleFilterer{contract: contract}}, nil
}

// NewSignedDiaOracleCaller creates a new read-only instance of SignedDiaOracle, bound to a specific deployed contract.
func NewSignedDiaOracleCaller(address common.Address, caller bind.ContractCaller) (*SignedDiaOracleCaller, error) {
	contract, err := bindSignedDiaOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignedDiaOracleCaller{contract: contract}, nil
}

// NewSignedDiaOracleTransactor creates a new write-only instance of SignedDiaOracle, bound to a specific deployed contract.
func NewSignedDiaOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*SignedDiaOracleTransactor, error) {
	contract, err := bindSignedDiaOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignedDiaOracleTransactor{contract: contract}, nil
}

// NewSignedDiaOracleFilterer creates a new log filterer instance of SignedDiaOracle, bound to a specific deployed contract.
func NewSignedDiaOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*SignedDiaOracleFilterer, error) {
	contract, err := bindSignedDiaOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignedDiaOracleFilterer{contract: contract}, nil
}

// bindSignedDiaOracle binds a generic wrapper to an already deployed contract.
func bindSignedDiaOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignedDiaOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedDiaOracle *SignedDiaOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignedDiaOracle.Contract.SignedDiaOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedDiaOracle *SignedDiaOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedDiaOracle.Contract.SignedDiaOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedDiaOracle *SignedDiaOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedDiaOracle.Contract.SignedDiaOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedDiaOracle *SignedDiaOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignedDiaOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedDiaOracle *SignedDiaOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedDiaOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedDiaOracle *SignedDiaOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedDiaOracle.Contract.contract.Transact(opts, method, params...)
}

// GetCoinInfo is a free data retrieval call binding the contract method 0xe1221400.
//
// Solidity: function getCoinInfo(string name) view returns(uint256, uint256, uint256, string, string)
func (_SignedDiaOracle *SignedDiaOracleCaller) GetCoinInfo(opts *bind.CallOpts, name string) (*big.Int, *big.Int, *big.Int, string, string, error) {
	var out []interface{}
	err := _SignedDiaOracle.contract.Call(opts, &out, "getCoinInfo", name)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)

	return out0, out1, out2, out3, out4, err

}

// GetCoinInfo is a free data retrieval call binding the contract method 0xe1221400.
//
// Solidity: function getCoinInfo(string name) view returns(uint256, uint256, uint256, string, string)
func (_SignedDiaOracle *SignedDiaOracleSession) GetCoinInfo(name string) (*big.Int, *big.Int, *big.Int, string, string, error) {
	return _SignedDiaOracle.Contract.GetCoinInfo(&_SignedDiaOracle.CallOpts, name)
}

// GetCoinInfo is a free data retrieval call binding the contract method 0xe1221400.
//
// Solidity: function getCoinInfo(string name) view returns(uint256, uint256, uint256, string, string)
func (_SignedDiaOracle *SignedDiaOracleCallerSession) GetCoinInfo(name string) (*big.Int, *big.Int, *big.Int, string, string, error) {
	return _SignedDiaOracle.Contract.GetCoinInfo(&_SignedDiaOracle.CallOpts, name)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_SignedDiaOracle *SignedDiaOracleTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SignedDiaOracle.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_SignedDiaOracle *SignedDiaOracleSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _SignedDiaOracle.Contract.ChangeOwner(&_SignedDiaOracle.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_SignedDiaOracle *SignedDiaOracleTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _SignedDiaOracle.Contract.ChangeOwner(&_SignedDiaOracle.TransactOpts, newOwner)
}

// UpdateCoinInfo is a paid mutator transaction binding the contract method 0xcb860a99.
//
// Solidity: function updateCoinInfo(string name, string symbol, uint256 newPrice, uint256 newSupply, uint256 newTimestamp, string lastSignedData) returns()
func (_SignedDiaOracle *SignedDiaOracleTransactor) UpdateCoinInfo(opts *bind.TransactOpts, name string, symbol string, newPrice *big.Int, newSupply *big.Int, newTimestamp *big.Int, lastSignedData string) (*types.Transaction, error) {
	return _SignedDiaOracle.contract.Transact(opts, "updateCoinInfo", name, symbol, newPrice, newSupply, newTimestamp, lastSignedData)
}

// UpdateCoinInfo is a paid mutator transaction binding the contract method 0xcb860a99.
//
// Solidity: function updateCoinInfo(string name, string symbol, uint256 newPrice, uint256 newSupply, uint256 newTimestamp, string lastSignedData) returns()
func (_SignedDiaOracle *SignedDiaOracleSession) UpdateCoinInfo(name string, symbol string, newPrice *big.Int, newSupply *big.Int, newTimestamp *big.Int, lastSignedData string) (*types.Transaction, error) {
	return _SignedDiaOracle.Contract.UpdateCoinInfo(&_SignedDiaOracle.TransactOpts, name, symbol, newPrice, newSupply, newTimestamp, lastSignedData)
}

// UpdateCoinInfo is a paid mutator transaction binding the contract method 0xcb860a99.
//
// Solidity: function updateCoinInfo(string name, string symbol, uint256 newPrice, uint256 newSupply, uint256 newTimestamp, string lastSignedData) returns()
func (_SignedDiaOracle *SignedDiaOracleTransactorSession) UpdateCoinInfo(name string, symbol string, newPrice *big.Int, newSupply *big.Int, newTimestamp *big.Int, lastSignedData string) (*types.Transaction, error) {
	return _SignedDiaOracle.Contract.UpdateCoinInfo(&_SignedDiaOracle.TransactOpts, name, symbol, newPrice, newSupply, newTimestamp, lastSignedData)
}

// SignedDiaOracleNewCoinInfoIterator is returned from FilterNewCoinInfo and is used to iterate over the raw logs and unpacked data for NewCoinInfo events raised by the SignedDiaOracle contract.
type SignedDiaOracleNewCoinInfoIterator struct {
	Event *SignedDiaOracleNewCoinInfo // Event containing the contract specifics and raw log

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
func (it *SignedDiaOracleNewCoinInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignedDiaOracleNewCoinInfo)
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
		it.Event = new(SignedDiaOracleNewCoinInfo)
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
func (it *SignedDiaOracleNewCoinInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignedDiaOracleNewCoinInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignedDiaOracleNewCoinInfo represents a NewCoinInfo event raised by the SignedDiaOracle contract.
type SignedDiaOracleNewCoinInfo struct {
	Name                string
	Symbol              string
	Price               *big.Int
	Supply              *big.Int
	LastUpdateTimestamp *big.Int
	LastSignedData      string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewCoinInfo is a free log retrieval operation binding the contract event 0xac95fcc9c84519975daf567027a248de195148a45c550fa5a4addc70aabe625f.
//
// Solidity: event newCoinInfo(string name, string symbol, uint256 price, uint256 supply, uint256 lastUpdateTimestamp, string lastSignedData)
func (_SignedDiaOracle *SignedDiaOracleFilterer) FilterNewCoinInfo(opts *bind.FilterOpts) (*SignedDiaOracleNewCoinInfoIterator, error) {

	logs, sub, err := _SignedDiaOracle.contract.FilterLogs(opts, "newCoinInfo")
	if err != nil {
		return nil, err
	}
	return &SignedDiaOracleNewCoinInfoIterator{contract: _SignedDiaOracle.contract, event: "newCoinInfo", logs: logs, sub: sub}, nil
}

// WatchNewCoinInfo is a free log subscription operation binding the contract event 0xac95fcc9c84519975daf567027a248de195148a45c550fa5a4addc70aabe625f.
//
// Solidity: event newCoinInfo(string name, string symbol, uint256 price, uint256 supply, uint256 lastUpdateTimestamp, string lastSignedData)
func (_SignedDiaOracle *SignedDiaOracleFilterer) WatchNewCoinInfo(opts *bind.WatchOpts, sink chan<- *SignedDiaOracleNewCoinInfo) (event.Subscription, error) {

	logs, sub, err := _SignedDiaOracle.contract.WatchLogs(opts, "newCoinInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignedDiaOracleNewCoinInfo)
				if err := _SignedDiaOracle.contract.UnpackLog(event, "newCoinInfo", log); err != nil {
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

// ParseNewCoinInfo is a log parse operation binding the contract event 0xac95fcc9c84519975daf567027a248de195148a45c550fa5a4addc70aabe625f.
//
// Solidity: event newCoinInfo(string name, string symbol, uint256 price, uint256 supply, uint256 lastUpdateTimestamp, string lastSignedData)
func (_SignedDiaOracle *SignedDiaOracleFilterer) ParseNewCoinInfo(log types.Log) (*SignedDiaOracleNewCoinInfo, error) {
	event := new(SignedDiaOracleNewCoinInfo)
	if err := _SignedDiaOracle.contract.UnpackLog(event, "newCoinInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
