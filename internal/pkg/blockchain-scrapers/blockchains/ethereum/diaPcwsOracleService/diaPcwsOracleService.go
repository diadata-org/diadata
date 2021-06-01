// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package diaPcwsOracleService

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

// DIAPcwsOracleABI is the input ABI used to generate the binding from.
const DIAPcwsOracleABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"name\":\"OracleUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUpdater\",\"type\":\"address\"}],\"name\":\"UpdaterAddressChange\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOracleUpdaterAddress\",\"type\":\"address\"}],\"name\":\"updateOracleUpdaterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DIAPcwsOracleFuncSigs maps the 4-byte function signature to its string representation.
var DIAPcwsOracleFuncSigs = map[string]string{
	"960384a0": "getValue(string)",
	"7898e0c2": "setValue(string,uint128,uint128)",
	"6aa45efc": "updateOracleUpdaterAddress(address)",
	"5a9ade8b": "values(string)",
}

// DIAPcwsOracleBin is the compiled bytecode used for deploying new contracts.
var DIAPcwsOracleBin = "0x608060405234801561001057600080fd5b50600180546001600160a01b0319163317905561054d806100326000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80635a9ade8b146100515780636aa45efc146101095780637898e0c214610131578063960384a0146101ed575b600080fd5b6100f76004803603602081101561006757600080fd5b81019060208101813564010000000081111561008257600080fd5b82018360208201111561009457600080fd5b803590602001918460018302840111640100000000831117156100b657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506102c2945050505050565b60408051918252519081900360200190f35b61012f6004803603602081101561011f57600080fd5b50356001600160a01b03166102df565b005b61012f6004803603606081101561014757600080fd5b81019060208101813564010000000081111561016257600080fd5b82018360208201111561017457600080fd5b8035906020019184600183028401116401000000008311171561019657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550506001600160801b03833581169450602090930135909216915061034a9050565b6102936004803603602081101561020357600080fd5b81019060208101813564010000000081111561021e57600080fd5b82018360208201111561023057600080fd5b8035906020019184600183028401116401000000008311171561025257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061049a945050505050565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b805160208183018101805160008252928201919093012091525481565b6001546001600160a01b031633146102f657600080fd5b600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f9181900360200190a150565b6001546001600160a01b0316331461036157600080fd5b6000816001600160801b03166080846001600160801b0316901b019050806000856040518082805190602001908083835b602083106103b15780518252601f199092019160209182019101610392565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382018520959095556001600160801b03888116858301528716948401949094525050606080825286519082015285517fa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782928792879287928291608083019187019080838360005b83811015610458578181015183820152602001610440565b50505050905090810190601f1680156104855780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a150505050565b600080600080846040518082805190602001908083835b602083106104d05780518252601f1990920191602091820191016104b1565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054608081901c976001600160801b03909116965094505050505056fea2646970667358221220f3525764c37c875437a86a2e25af9f39be8b5903b0339584be3fa42f2f60be7364736f6c63430007040033"

// DeployDIAPcwsOracle deploys a new Ethereum contract, binding an instance of DIAPcwsOracle to it.
func DeployDIAPcwsOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DIAPcwsOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(DIAPcwsOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DIAPcwsOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DIAPcwsOracle{DIAPcwsOracleCaller: DIAPcwsOracleCaller{contract: contract}, DIAPcwsOracleTransactor: DIAPcwsOracleTransactor{contract: contract}, DIAPcwsOracleFilterer: DIAPcwsOracleFilterer{contract: contract}}, nil
}

// DIAPcwsOracle is an auto generated Go binding around an Ethereum contract.
type DIAPcwsOracle struct {
	DIAPcwsOracleCaller     // Read-only binding to the contract
	DIAPcwsOracleTransactor // Write-only binding to the contract
	DIAPcwsOracleFilterer   // Log filterer for contract events
}

// DIAPcwsOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type DIAPcwsOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIAPcwsOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DIAPcwsOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIAPcwsOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DIAPcwsOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIAPcwsOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DIAPcwsOracleSession struct {
	Contract     *DIAPcwsOracle    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DIAPcwsOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DIAPcwsOracleCallerSession struct {
	Contract *DIAPcwsOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DIAPcwsOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DIAPcwsOracleTransactorSession struct {
	Contract     *DIAPcwsOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DIAPcwsOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type DIAPcwsOracleRaw struct {
	Contract *DIAPcwsOracle // Generic contract binding to access the raw methods on
}

// DIAPcwsOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DIAPcwsOracleCallerRaw struct {
	Contract *DIAPcwsOracleCaller // Generic read-only contract binding to access the raw methods on
}

// DIAPcwsOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DIAPcwsOracleTransactorRaw struct {
	Contract *DIAPcwsOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDIAPcwsOracle creates a new instance of DIAPcwsOracle, bound to a specific deployed contract.
func NewDIAPcwsOracle(address common.Address, backend bind.ContractBackend) (*DIAPcwsOracle, error) {
	contract, err := bindDIAPcwsOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DIAPcwsOracle{DIAPcwsOracleCaller: DIAPcwsOracleCaller{contract: contract}, DIAPcwsOracleTransactor: DIAPcwsOracleTransactor{contract: contract}, DIAPcwsOracleFilterer: DIAPcwsOracleFilterer{contract: contract}}, nil
}

// NewDIAPcwsOracleCaller creates a new read-only instance of DIAPcwsOracle, bound to a specific deployed contract.
func NewDIAPcwsOracleCaller(address common.Address, caller bind.ContractCaller) (*DIAPcwsOracleCaller, error) {
	contract, err := bindDIAPcwsOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DIAPcwsOracleCaller{contract: contract}, nil
}

// NewDIAPcwsOracleTransactor creates a new write-only instance of DIAPcwsOracle, bound to a specific deployed contract.
func NewDIAPcwsOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*DIAPcwsOracleTransactor, error) {
	contract, err := bindDIAPcwsOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DIAPcwsOracleTransactor{contract: contract}, nil
}

// NewDIAPcwsOracleFilterer creates a new log filterer instance of DIAPcwsOracle, bound to a specific deployed contract.
func NewDIAPcwsOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*DIAPcwsOracleFilterer, error) {
	contract, err := bindDIAPcwsOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DIAPcwsOracleFilterer{contract: contract}, nil
}

// bindDIAPcwsOracle binds a generic wrapper to an already deployed contract.
func bindDIAPcwsOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DIAPcwsOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIAPcwsOracle *DIAPcwsOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIAPcwsOracle.Contract.DIAPcwsOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIAPcwsOracle *DIAPcwsOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIAPcwsOracle.Contract.DIAPcwsOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIAPcwsOracle *DIAPcwsOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIAPcwsOracle.Contract.DIAPcwsOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIAPcwsOracle *DIAPcwsOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIAPcwsOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIAPcwsOracle *DIAPcwsOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIAPcwsOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIAPcwsOracle *DIAPcwsOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIAPcwsOracle.Contract.contract.Transact(opts, method, params...)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_DIAPcwsOracle *DIAPcwsOracleCaller) GetValue(opts *bind.CallOpts, key string) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _DIAPcwsOracle.contract.Call(opts, &out, "getValue", key)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_DIAPcwsOracle *DIAPcwsOracleSession) GetValue(key string) (*big.Int, *big.Int, error) {
	return _DIAPcwsOracle.Contract.GetValue(&_DIAPcwsOracle.CallOpts, key)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_DIAPcwsOracle *DIAPcwsOracleCallerSession) GetValue(key string) (*big.Int, *big.Int, error) {
	return _DIAPcwsOracle.Contract.GetValue(&_DIAPcwsOracle.CallOpts, key)
}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_DIAPcwsOracle *DIAPcwsOracleCaller) Values(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _DIAPcwsOracle.contract.Call(opts, &out, "values", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_DIAPcwsOracle *DIAPcwsOracleSession) Values(arg0 string) (*big.Int, error) {
	return _DIAPcwsOracle.Contract.Values(&_DIAPcwsOracle.CallOpts, arg0)
}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_DIAPcwsOracle *DIAPcwsOracleCallerSession) Values(arg0 string) (*big.Int, error) {
	return _DIAPcwsOracle.Contract.Values(&_DIAPcwsOracle.CallOpts, arg0)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_DIAPcwsOracle *DIAPcwsOracleTransactor) SetValue(opts *bind.TransactOpts, key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _DIAPcwsOracle.contract.Transact(opts, "setValue", key, value, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_DIAPcwsOracle *DIAPcwsOracleSession) SetValue(key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _DIAPcwsOracle.Contract.SetValue(&_DIAPcwsOracle.TransactOpts, key, value, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_DIAPcwsOracle *DIAPcwsOracleTransactorSession) SetValue(key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _DIAPcwsOracle.Contract.SetValue(&_DIAPcwsOracle.TransactOpts, key, value, timestamp)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DIAPcwsOracle *DIAPcwsOracleTransactor) UpdateOracleUpdaterAddress(opts *bind.TransactOpts, newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DIAPcwsOracle.contract.Transact(opts, "updateOracleUpdaterAddress", newOracleUpdaterAddress)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DIAPcwsOracle *DIAPcwsOracleSession) UpdateOracleUpdaterAddress(newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DIAPcwsOracle.Contract.UpdateOracleUpdaterAddress(&_DIAPcwsOracle.TransactOpts, newOracleUpdaterAddress)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DIAPcwsOracle *DIAPcwsOracleTransactorSession) UpdateOracleUpdaterAddress(newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DIAPcwsOracle.Contract.UpdateOracleUpdaterAddress(&_DIAPcwsOracle.TransactOpts, newOracleUpdaterAddress)
}

// DIAPcwsOracleOracleUpdateIterator is returned from FilterOracleUpdate and is used to iterate over the raw logs and unpacked data for OracleUpdate events raised by the DIAPcwsOracle contract.
type DIAPcwsOracleOracleUpdateIterator struct {
	Event *DIAPcwsOracleOracleUpdate // Event containing the contract specifics and raw log

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
func (it *DIAPcwsOracleOracleUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIAPcwsOracleOracleUpdate)
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
		it.Event = new(DIAPcwsOracleOracleUpdate)
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
func (it *DIAPcwsOracleOracleUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIAPcwsOracleOracleUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIAPcwsOracleOracleUpdate represents a OracleUpdate event raised by the DIAPcwsOracle contract.
type DIAPcwsOracleOracleUpdate struct {
	Key       string
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleUpdate is a free log retrieval operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_DIAPcwsOracle *DIAPcwsOracleFilterer) FilterOracleUpdate(opts *bind.FilterOpts) (*DIAPcwsOracleOracleUpdateIterator, error) {

	logs, sub, err := _DIAPcwsOracle.contract.FilterLogs(opts, "OracleUpdate")
	if err != nil {
		return nil, err
	}
	return &DIAPcwsOracleOracleUpdateIterator{contract: _DIAPcwsOracle.contract, event: "OracleUpdate", logs: logs, sub: sub}, nil
}

// WatchOracleUpdate is a free log subscription operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_DIAPcwsOracle *DIAPcwsOracleFilterer) WatchOracleUpdate(opts *bind.WatchOpts, sink chan<- *DIAPcwsOracleOracleUpdate) (event.Subscription, error) {

	logs, sub, err := _DIAPcwsOracle.contract.WatchLogs(opts, "OracleUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIAPcwsOracleOracleUpdate)
				if err := _DIAPcwsOracle.contract.UnpackLog(event, "OracleUpdate", log); err != nil {
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
func (_DIAPcwsOracle *DIAPcwsOracleFilterer) ParseOracleUpdate(log types.Log) (*DIAPcwsOracleOracleUpdate, error) {
	event := new(DIAPcwsOracleOracleUpdate)
	if err := _DIAPcwsOracle.contract.UnpackLog(event, "OracleUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DIAPcwsOracleUpdaterAddressChangeIterator is returned from FilterUpdaterAddressChange and is used to iterate over the raw logs and unpacked data for UpdaterAddressChange events raised by the DIAPcwsOracle contract.
type DIAPcwsOracleUpdaterAddressChangeIterator struct {
	Event *DIAPcwsOracleUpdaterAddressChange // Event containing the contract specifics and raw log

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
func (it *DIAPcwsOracleUpdaterAddressChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIAPcwsOracleUpdaterAddressChange)
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
		it.Event = new(DIAPcwsOracleUpdaterAddressChange)
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
func (it *DIAPcwsOracleUpdaterAddressChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIAPcwsOracleUpdaterAddressChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIAPcwsOracleUpdaterAddressChange represents a UpdaterAddressChange event raised by the DIAPcwsOracle contract.
type DIAPcwsOracleUpdaterAddressChange struct {
	NewUpdater common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdaterAddressChange is a free log retrieval operation binding the contract event 0x121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f.
//
// Solidity: event UpdaterAddressChange(address newUpdater)
func (_DIAPcwsOracle *DIAPcwsOracleFilterer) FilterUpdaterAddressChange(opts *bind.FilterOpts) (*DIAPcwsOracleUpdaterAddressChangeIterator, error) {

	logs, sub, err := _DIAPcwsOracle.contract.FilterLogs(opts, "UpdaterAddressChange")
	if err != nil {
		return nil, err
	}
	return &DIAPcwsOracleUpdaterAddressChangeIterator{contract: _DIAPcwsOracle.contract, event: "UpdaterAddressChange", logs: logs, sub: sub}, nil
}

// WatchUpdaterAddressChange is a free log subscription operation binding the contract event 0x121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f.
//
// Solidity: event UpdaterAddressChange(address newUpdater)
func (_DIAPcwsOracle *DIAPcwsOracleFilterer) WatchUpdaterAddressChange(opts *bind.WatchOpts, sink chan<- *DIAPcwsOracleUpdaterAddressChange) (event.Subscription, error) {

	logs, sub, err := _DIAPcwsOracle.contract.WatchLogs(opts, "UpdaterAddressChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIAPcwsOracleUpdaterAddressChange)
				if err := _DIAPcwsOracle.contract.UnpackLog(event, "UpdaterAddressChange", log); err != nil {
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
func (_DIAPcwsOracle *DIAPcwsOracleFilterer) ParseUpdaterAddressChange(log types.Log) (*DIAPcwsOracleUpdaterAddressChange, error) {
	event := new(DIAPcwsOracleUpdaterAddressChange)
	if err := _DIAPcwsOracle.contract.UnpackLog(event, "UpdaterAddressChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
