// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package diaCoingeckoOracleService

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

// DIACoingecoOracleABI is the input ABI used to generate the binding from.
const DIACoingecoOracleABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"name\":\"OracleUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUpdater\",\"type\":\"address\"}],\"name\":\"UpdaterAddressChange\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOracleUpdaterAddress\",\"type\":\"address\"}],\"name\":\"updateOracleUpdaterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DIACoingecoOracleFuncSigs maps the 4-byte function signature to its string representation.
var DIACoingecoOracleFuncSigs = map[string]string{
	"960384a0": "getValue(string)",
	"7898e0c2": "setValue(string,uint128,uint128)",
	"6aa45efc": "updateOracleUpdaterAddress(address)",
	"5a9ade8b": "values(string)",
}

// DIACoingecoOracleBin is the compiled bytecode used for deploying new contracts.
var DIACoingecoOracleBin = "0x608060405234801561001057600080fd5b50600180546001600160a01b0319163317905561054d806100326000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80635a9ade8b146100515780636aa45efc146101095780637898e0c214610131578063960384a0146101ed575b600080fd5b6100f76004803603602081101561006757600080fd5b81019060208101813564010000000081111561008257600080fd5b82018360208201111561009457600080fd5b803590602001918460018302840111640100000000831117156100b657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506102c2945050505050565b60408051918252519081900360200190f35b61012f6004803603602081101561011f57600080fd5b50356001600160a01b03166102df565b005b61012f6004803603606081101561014757600080fd5b81019060208101813564010000000081111561016257600080fd5b82018360208201111561017457600080fd5b8035906020019184600183028401116401000000008311171561019657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550506001600160801b03833581169450602090930135909216915061034a9050565b6102936004803603602081101561020357600080fd5b81019060208101813564010000000081111561021e57600080fd5b82018360208201111561023057600080fd5b8035906020019184600183028401116401000000008311171561025257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061049a945050505050565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b805160208183018101805160008252928201919093012091525481565b6001546001600160a01b031633146102f657600080fd5b600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f9181900360200190a150565b6001546001600160a01b0316331461036157600080fd5b6000816001600160801b03166080846001600160801b0316901b019050806000856040518082805190602001908083835b602083106103b15780518252601f199092019160209182019101610392565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382018520959095556001600160801b03888116858301528716948401949094525050606080825286519082015285517fa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782928792879287928291608083019187019080838360005b83811015610458578181015183820152602001610440565b50505050905090810190601f1680156104855780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a150505050565b600080600080846040518082805190602001908083835b602083106104d05780518252601f1990920191602091820191016104b1565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054608081901c976001600160801b03909116965094505050505056fea2646970667358221220ee50177bc78c356d3eacf55cc6cc7ca7eb242b260453245fb2c2eb78340bd52864736f6c63430007000033"

// DeployDIACoingecoOracle deploys a new Ethereum contract, binding an instance of DIACoingecoOracle to it.
func DeployDIACoingecoOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DIACoingecoOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(DIACoingecoOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DIACoingecoOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DIACoingecoOracle{DIACoingecoOracleCaller: DIACoingecoOracleCaller{contract: contract}, DIACoingecoOracleTransactor: DIACoingecoOracleTransactor{contract: contract}, DIACoingecoOracleFilterer: DIACoingecoOracleFilterer{contract: contract}}, nil
}

// DIACoingecoOracle is an auto generated Go binding around an Ethereum contract.
type DIACoingecoOracle struct {
	DIACoingecoOracleCaller     // Read-only binding to the contract
	DIACoingecoOracleTransactor // Write-only binding to the contract
	DIACoingecoOracleFilterer   // Log filterer for contract events
}

// DIACoingecoOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type DIACoingecoOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIACoingecoOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DIACoingecoOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIACoingecoOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DIACoingecoOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIACoingecoOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DIACoingecoOracleSession struct {
	Contract     *DIACoingecoOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DIACoingecoOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DIACoingecoOracleCallerSession struct {
	Contract *DIACoingecoOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DIACoingecoOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DIACoingecoOracleTransactorSession struct {
	Contract     *DIACoingecoOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DIACoingecoOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type DIACoingecoOracleRaw struct {
	Contract *DIACoingecoOracle // Generic contract binding to access the raw methods on
}

// DIACoingecoOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DIACoingecoOracleCallerRaw struct {
	Contract *DIACoingecoOracleCaller // Generic read-only contract binding to access the raw methods on
}

// DIACoingecoOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DIACoingecoOracleTransactorRaw struct {
	Contract *DIACoingecoOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDIACoingecoOracle creates a new instance of DIACoingecoOracle, bound to a specific deployed contract.
func NewDIACoingecoOracle(address common.Address, backend bind.ContractBackend) (*DIACoingecoOracle, error) {
	contract, err := bindDIACoingecoOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DIACoingecoOracle{DIACoingecoOracleCaller: DIACoingecoOracleCaller{contract: contract}, DIACoingecoOracleTransactor: DIACoingecoOracleTransactor{contract: contract}, DIACoingecoOracleFilterer: DIACoingecoOracleFilterer{contract: contract}}, nil
}

// NewDIACoingecoOracleCaller creates a new read-only instance of DIACoingecoOracle, bound to a specific deployed contract.
func NewDIACoingecoOracleCaller(address common.Address, caller bind.ContractCaller) (*DIACoingecoOracleCaller, error) {
	contract, err := bindDIACoingecoOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DIACoingecoOracleCaller{contract: contract}, nil
}

// NewDIACoingecoOracleTransactor creates a new write-only instance of DIACoingecoOracle, bound to a specific deployed contract.
func NewDIACoingecoOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*DIACoingecoOracleTransactor, error) {
	contract, err := bindDIACoingecoOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DIACoingecoOracleTransactor{contract: contract}, nil
}

// NewDIACoingecoOracleFilterer creates a new log filterer instance of DIACoingecoOracle, bound to a specific deployed contract.
func NewDIACoingecoOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*DIACoingecoOracleFilterer, error) {
	contract, err := bindDIACoingecoOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DIACoingecoOracleFilterer{contract: contract}, nil
}

// bindDIACoingecoOracle binds a generic wrapper to an already deployed contract.
func bindDIACoingecoOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DIACoingecoOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIACoingecoOracle *DIACoingecoOracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DIACoingecoOracle.Contract.DIACoingecoOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIACoingecoOracle *DIACoingecoOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIACoingecoOracle.Contract.DIACoingecoOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIACoingecoOracle *DIACoingecoOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIACoingecoOracle.Contract.DIACoingecoOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIACoingecoOracle *DIACoingecoOracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DIACoingecoOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIACoingecoOracle *DIACoingecoOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIACoingecoOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIACoingecoOracle *DIACoingecoOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIACoingecoOracle.Contract.contract.Transact(opts, method, params...)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_DIACoingecoOracle *DIACoingecoOracleCaller) GetValue(opts *bind.CallOpts, key string) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DIACoingecoOracle.contract.Call(opts, out, "getValue", key)
	return *ret0, *ret1, err
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_DIACoingecoOracle *DIACoingecoOracleSession) GetValue(key string) (*big.Int, *big.Int, error) {
	return _DIACoingecoOracle.Contract.GetValue(&_DIACoingecoOracle.CallOpts, key)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_DIACoingecoOracle *DIACoingecoOracleCallerSession) GetValue(key string) (*big.Int, *big.Int, error) {
	return _DIACoingecoOracle.Contract.GetValue(&_DIACoingecoOracle.CallOpts, key)
}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_DIACoingecoOracle *DIACoingecoOracleCaller) Values(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DIACoingecoOracle.contract.Call(opts, out, "values", arg0)
	return *ret0, err
}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_DIACoingecoOracle *DIACoingecoOracleSession) Values(arg0 string) (*big.Int, error) {
	return _DIACoingecoOracle.Contract.Values(&_DIACoingecoOracle.CallOpts, arg0)
}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_DIACoingecoOracle *DIACoingecoOracleCallerSession) Values(arg0 string) (*big.Int, error) {
	return _DIACoingecoOracle.Contract.Values(&_DIACoingecoOracle.CallOpts, arg0)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_DIACoingecoOracle *DIACoingecoOracleTransactor) SetValue(opts *bind.TransactOpts, key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _DIACoingecoOracle.contract.Transact(opts, "setValue", key, value, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_DIACoingecoOracle *DIACoingecoOracleSession) SetValue(key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _DIACoingecoOracle.Contract.SetValue(&_DIACoingecoOracle.TransactOpts, key, value, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_DIACoingecoOracle *DIACoingecoOracleTransactorSession) SetValue(key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _DIACoingecoOracle.Contract.SetValue(&_DIACoingecoOracle.TransactOpts, key, value, timestamp)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DIACoingecoOracle *DIACoingecoOracleTransactor) UpdateOracleUpdaterAddress(opts *bind.TransactOpts, newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DIACoingecoOracle.contract.Transact(opts, "updateOracleUpdaterAddress", newOracleUpdaterAddress)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DIACoingecoOracle *DIACoingecoOracleSession) UpdateOracleUpdaterAddress(newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DIACoingecoOracle.Contract.UpdateOracleUpdaterAddress(&_DIACoingecoOracle.TransactOpts, newOracleUpdaterAddress)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DIACoingecoOracle *DIACoingecoOracleTransactorSession) UpdateOracleUpdaterAddress(newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DIACoingecoOracle.Contract.UpdateOracleUpdaterAddress(&_DIACoingecoOracle.TransactOpts, newOracleUpdaterAddress)
}

// DIACoingecoOracleOracleUpdateIterator is returned from FilterOracleUpdate and is used to iterate over the raw logs and unpacked data for OracleUpdate events raised by the DIACoingecoOracle contract.
type DIACoingecoOracleOracleUpdateIterator struct {
	Event *DIACoingecoOracleOracleUpdate // Event containing the contract specifics and raw log

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
func (it *DIACoingecoOracleOracleUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIACoingecoOracleOracleUpdate)
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
		it.Event = new(DIACoingecoOracleOracleUpdate)
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
func (it *DIACoingecoOracleOracleUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIACoingecoOracleOracleUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIACoingecoOracleOracleUpdate represents a OracleUpdate event raised by the DIACoingecoOracle contract.
type DIACoingecoOracleOracleUpdate struct {
	Key       string
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleUpdate is a free log retrieval operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_DIACoingecoOracle *DIACoingecoOracleFilterer) FilterOracleUpdate(opts *bind.FilterOpts) (*DIACoingecoOracleOracleUpdateIterator, error) {

	logs, sub, err := _DIACoingecoOracle.contract.FilterLogs(opts, "OracleUpdate")
	if err != nil {
		return nil, err
	}
	return &DIACoingecoOracleOracleUpdateIterator{contract: _DIACoingecoOracle.contract, event: "OracleUpdate", logs: logs, sub: sub}, nil
}

// WatchOracleUpdate is a free log subscription operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_DIACoingecoOracle *DIACoingecoOracleFilterer) WatchOracleUpdate(opts *bind.WatchOpts, sink chan<- *DIACoingecoOracleOracleUpdate) (event.Subscription, error) {

	logs, sub, err := _DIACoingecoOracle.contract.WatchLogs(opts, "OracleUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIACoingecoOracleOracleUpdate)
				if err := _DIACoingecoOracle.contract.UnpackLog(event, "OracleUpdate", log); err != nil {
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
func (_DIACoingecoOracle *DIACoingecoOracleFilterer) ParseOracleUpdate(log types.Log) (*DIACoingecoOracleOracleUpdate, error) {
	event := new(DIACoingecoOracleOracleUpdate)
	if err := _DIACoingecoOracle.contract.UnpackLog(event, "OracleUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DIACoingecoOracleUpdaterAddressChangeIterator is returned from FilterUpdaterAddressChange and is used to iterate over the raw logs and unpacked data for UpdaterAddressChange events raised by the DIACoingecoOracle contract.
type DIACoingecoOracleUpdaterAddressChangeIterator struct {
	Event *DIACoingecoOracleUpdaterAddressChange // Event containing the contract specifics and raw log

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
func (it *DIACoingecoOracleUpdaterAddressChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIACoingecoOracleUpdaterAddressChange)
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
		it.Event = new(DIACoingecoOracleUpdaterAddressChange)
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
func (it *DIACoingecoOracleUpdaterAddressChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIACoingecoOracleUpdaterAddressChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIACoingecoOracleUpdaterAddressChange represents a UpdaterAddressChange event raised by the DIACoingecoOracle contract.
type DIACoingecoOracleUpdaterAddressChange struct {
	NewUpdater common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdaterAddressChange is a free log retrieval operation binding the contract event 0x121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f.
//
// Solidity: event UpdaterAddressChange(address newUpdater)
func (_DIACoingecoOracle *DIACoingecoOracleFilterer) FilterUpdaterAddressChange(opts *bind.FilterOpts) (*DIACoingecoOracleUpdaterAddressChangeIterator, error) {

	logs, sub, err := _DIACoingecoOracle.contract.FilterLogs(opts, "UpdaterAddressChange")
	if err != nil {
		return nil, err
	}
	return &DIACoingecoOracleUpdaterAddressChangeIterator{contract: _DIACoingecoOracle.contract, event: "UpdaterAddressChange", logs: logs, sub: sub}, nil
}

// WatchUpdaterAddressChange is a free log subscription operation binding the contract event 0x121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f.
//
// Solidity: event UpdaterAddressChange(address newUpdater)
func (_DIACoingecoOracle *DIACoingecoOracleFilterer) WatchUpdaterAddressChange(opts *bind.WatchOpts, sink chan<- *DIACoingecoOracleUpdaterAddressChange) (event.Subscription, error) {

	logs, sub, err := _DIACoingecoOracle.contract.WatchLogs(opts, "UpdaterAddressChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIACoingecoOracleUpdaterAddressChange)
				if err := _DIACoingecoOracle.contract.UnpackLog(event, "UpdaterAddressChange", log); err != nil {
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
func (_DIACoingecoOracle *DIACoingecoOracleFilterer) ParseUpdaterAddressChange(log types.Log) (*DIACoingecoOracleUpdaterAddressChange, error) {
	event := new(DIACoingecoOracleUpdaterAddressChange)
	if err := _DIACoingecoOracle.contract.UnpackLog(event, "UpdaterAddressChange", log); err != nil {
		return nil, err
	}
	return event, nil
}
