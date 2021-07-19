// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxy

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

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Session) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20CallerSession) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Session) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20CallerSession) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Session) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20CallerSession) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"}]"

// OwnedFuncSigs maps the 4-byte function signature to its string representation.
var OwnedFuncSigs = map[string]string{
	"79ba5097": "acceptOwnership()",
	"1627540c": "nominateNewOwner(address)",
	"53a47bb7": "nominatedOwner()",
	"8da5cb5b": "owner()",
}

// OwnedBin is the compiled bytecode used for deploying new contracts.
var OwnedBin = "0x608060405234801561001057600080fd5b5060405160208061044e8339810160405251600160a060020a038116151561009957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e657220616464726573732063616e6e6f74206265203000000000000000604482015290519081900360640190fd5b60008054600160a060020a031916600160a060020a038316908117825560408051928352602083019190915280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a15061034d806101016000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631627540c811461006657806353a47bb71461008957806379ba5097146100ba5780638da5cb5b146100cf575b600080fd5b34801561007257600080fd5b50610087600160a060020a03600435166100e4565b005b34801561009557600080fd5b5061009e6101e4565b60408051600160a060020a039092168252519081900360200190f35b3480156100c657600080fd5b506100876101f3565b3480156100db57600080fd5b5061009e610312565b600054600160a060020a0316331461018357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f4f6e6c792074686520636f6e7472616374206f776e6572206d6179207065726660448201527f6f726d207468697320616374696f6e0000000000000000000000000000000000606482015290519081900360840190fd5b60018054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517f906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce229181900360200190a150565b600154600160a060020a031681565b600154600160a060020a0316331461029257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f596f75206d757374206265206e6f6d696e61746564206265666f726520796f7560448201527f2063616e20616363657074206f776e6572736869700000000000000000000000606482015290519081900360840190fd5b60005460015460408051600160a060020a03938416815292909116602083015280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a1600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a0316815600a165627a7a723058207ef60240d0b7fa023d07179fa4a14609dc89c256fabd51d397602fb1ab64aff60029"

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
	OwnedFilterer   // Log filterer for contract events
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// NewOwnedFilterer creates a new log filterer instance of Owned, bound to a specific deployed contract.
func NewOwnedFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnedFilterer, error) {
	contract, err := bindOwned(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnedFilterer{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Owned *OwnedCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Owned.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Owned *OwnedSession) NominatedOwner() (common.Address, error) {
	return _Owned.Contract.NominatedOwner(&_Owned.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Owned *OwnedCallerSession) NominatedOwner() (common.Address, error) {
	return _Owned.Contract.NominatedOwner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Owned *OwnedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Owned.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Owned *OwnedSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Owned *OwnedCallerSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedSession) AcceptOwnership() (*types.Transaction, error) {
	return _Owned.Contract.AcceptOwnership(&_Owned.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Owned.Contract.AcceptOwnership(&_Owned.TransactOpts)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Owned *OwnedTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Owned *OwnedSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.NominateNewOwner(&_Owned.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Owned *OwnedTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.NominateNewOwner(&_Owned.TransactOpts, _owner)
}

// OwnedOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Owned contract.
type OwnedOwnerChangedIterator struct {
	Event *OwnedOwnerChanged // Event containing the contract specifics and raw log

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
func (it *OwnedOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnedOwnerChanged)
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
		it.Event = new(OwnedOwnerChanged)
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
func (it *OwnedOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnedOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnedOwnerChanged represents a OwnerChanged event raised by the Owned contract.
type OwnedOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Owned *OwnedFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*OwnedOwnerChangedIterator, error) {

	logs, sub, err := _Owned.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &OwnedOwnerChangedIterator{contract: _Owned.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Owned *OwnedFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *OwnedOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Owned.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnedOwnerChanged)
				if err := _Owned.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Owned *OwnedFilterer) ParseOwnerChanged(log types.Log) (*OwnedOwnerChanged, error) {
	event := new(OwnedOwnerChanged)
	if err := _Owned.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnedOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the Owned contract.
type OwnedOwnerNominatedIterator struct {
	Event *OwnedOwnerNominated // Event containing the contract specifics and raw log

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
func (it *OwnedOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnedOwnerNominated)
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
		it.Event = new(OwnedOwnerNominated)
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
func (it *OwnedOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnedOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnedOwnerNominated represents a OwnerNominated event raised by the Owned contract.
type OwnedOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Owned *OwnedFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*OwnedOwnerNominatedIterator, error) {

	logs, sub, err := _Owned.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &OwnedOwnerNominatedIterator{contract: _Owned.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Owned *OwnedFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *OwnedOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _Owned.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnedOwnerNominated)
				if err := _Owned.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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

// ParseOwnerNominated is a log parse operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Owned *OwnedFilterer) ParseOwnerNominated(log types.Log) (*OwnedOwnerNominated, error) {
	event := new(OwnedOwnerNominated)
	if err := _Owned.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyABI is the input ABI used to generate the binding from.
const ProxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"setTarget\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"callData\",\"type\":\"bytes\"},{\"name\":\"numTopics\",\"type\":\"uint256\"},{\"name\":\"topic1\",\"type\":\"bytes32\"},{\"name\":\"topic2\",\"type\":\"bytes32\"},{\"name\":\"topic3\",\"type\":\"bytes32\"},{\"name\":\"topic4\",\"type\":\"bytes32\"}],\"name\":\"_emit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"useDELEGATECALL\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setUseDELEGATECALL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"target\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newTarget\",\"type\":\"address\"}],\"name\":\"TargetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"}]"

// ProxyFuncSigs maps the 4-byte function signature to its string representation.
var ProxyFuncSigs = map[string]string{
	"907dff97": "_emit(bytes,uint256,bytes32,bytes32,bytes32,bytes32)",
	"79ba5097": "acceptOwnership()",
	"1627540c": "nominateNewOwner(address)",
	"53a47bb7": "nominatedOwner()",
	"8da5cb5b": "owner()",
	"776d1a01": "setTarget(address)",
	"befff6af": "setUseDELEGATECALL(bool)",
	"d4b83992": "target()",
	"95578ebd": "useDELEGATECALL()",
}

// ProxyBin is the compiled bytecode used for deploying new contracts.
var ProxyBin = "0x608060405234801561001057600080fd5b506040516020806108dc833981016040525180600160a060020a038116151561009a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e657220616464726573732063616e6e6f74206265203000000000000000604482015290519081900360640190fd5b60008054600160a060020a031916600160a060020a038316908117825560408051928352602083019190915280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a150506107d9806101036000396000f3006080604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631627540c811461018357806353a47bb7146101a4578063776d1a01146101d557806379ba5097146101f65780638da5cb5b1461020b578063907dff971461022057806395578ebd14610250578063befff6af14610279578063d4b8399214610293575b60025474010000000000000000000000000000000000000000900460ff16156100e157604051366000823760008036836002545af43d6000833e8015156100dd573d82fd5b3d82f35b600254604080517fbc67f8320000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163bc67f8329160248082019260009290919082900301818387803b15801561014657600080fd5b505af115801561015a573d6000803e3d6000fd5b5050505060405136600082376000803683346002545af13d6000833e8015156100dd573d82fd5b005b34801561018f57600080fd5b50610181600160a060020a03600435166102a8565b3480156101b057600080fd5b506101b9610391565b60408051600160a060020a039092168252519081900360200190f35b3480156101e157600080fd5b50610181600160a060020a03600435166103a0565b34801561020257600080fd5b50610181610489565b34801561021757600080fd5b506101b9610591565b34801561022c57600080fd5b5061018160246004803582810192910135903560443560643560843560a4356105a0565b34801561025c57600080fd5b506102656106b5565b604080519115158252519081900360200190f35b34801561028557600080fd5b5061018160043515156106d6565b34801561029f57600080fd5b506101b961079e565b600054600160a060020a03163314610330576040805160e560020a62461bcd02815260206004820152602f60248201527f4f6e6c792074686520636f6e7472616374206f776e6572206d6179207065726660448201527f6f726d207468697320616374696f6e0000000000000000000000000000000000606482015290519081900360840190fd5b60018054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517f906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce229181900360200190a150565b600154600160a060020a031681565b600054600160a060020a03163314610428576040805160e560020a62461bcd02815260206004820152602f60248201527f4f6e6c792074686520636f6e7472616374206f776e6572206d6179207065726660448201527f6f726d207468697320616374696f6e0000000000000000000000000000000000606482015290519081900360840190fd5b60028054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517f814250a3b8c79fcbe2ead2c131c952a278491c8f4322a79fe84b5040a810373e9181900360200190a150565b600154600160a060020a03163314610511576040805160e560020a62461bcd02815260206004820152603560248201527f596f75206d757374206265206e6f6d696e61746564206265666f726520796f7560448201527f2063616e20616363657074206f776e6572736869700000000000000000000000606482015290519081900360840190fd5b60005460015460408051600160a060020a03938416815292909116602083015280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a1600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a031681565b600254600090606090600160a060020a03163314610608576040805160e560020a62461bcd02815260206004820152601460248201527f4d7573742062652070726f787920746172676574000000000000000000000000604482015290519081900360640190fd5b604080516020601f8b01819004810282018101909252898152899350908a908490819084018382808284378201915050505050509050866000811461066c576001811461067757600281146106835760038114610690576004811461069e576106a9565b8260208301a06106a9565b868360208401a16106a9565b85878460208501a26106a9565b8486888560208601a36106a9565b838587898660208701a45b50505050505050505050565b60025474010000000000000000000000000000000000000000900460ff1681565b600054600160a060020a0316331461075e576040805160e560020a62461bcd02815260206004820152602f60248201527f4f6e6c792074686520636f6e7472616374206f776e6572206d6179207065726660448201527f6f726d207468697320616374696f6e0000000000000000000000000000000000606482015290519081900360840190fd5b60028054911515740100000000000000000000000000000000000000000274ff000000000000000000000000000000000000000019909216919091179055565b600254600160a060020a0316815600a165627a7a72305820ab00d813af2a2ce0a914522b14872e57cc7c514c9c7c98379a79dccf387351710029"

// DeployProxy deploys a new Ethereum contract, binding an instance of Proxy to it.
func DeployProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *Proxy, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Proxy *ProxyCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Proxy *ProxySession) NominatedOwner() (common.Address, error) {
	return _Proxy.Contract.NominatedOwner(&_Proxy.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Proxy *ProxyCallerSession) NominatedOwner() (common.Address, error) {
	return _Proxy.Contract.NominatedOwner(&_Proxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxy *ProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxy *ProxySession) Owner() (common.Address, error) {
	return _Proxy.Contract.Owner(&_Proxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxy *ProxyCallerSession) Owner() (common.Address, error) {
	return _Proxy.Contract.Owner(&_Proxy.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_Proxy *ProxyCaller) Target(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "target")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_Proxy *ProxySession) Target() (common.Address, error) {
	return _Proxy.Contract.Target(&_Proxy.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_Proxy *ProxyCallerSession) Target() (common.Address, error) {
	return _Proxy.Contract.Target(&_Proxy.CallOpts)
}

// UseDELEGATECALL is a free data retrieval call binding the contract method 0x95578ebd.
//
// Solidity: function useDELEGATECALL() view returns(bool)
func (_Proxy *ProxyCaller) UseDELEGATECALL(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "useDELEGATECALL")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDELEGATECALL is a free data retrieval call binding the contract method 0x95578ebd.
//
// Solidity: function useDELEGATECALL() view returns(bool)
func (_Proxy *ProxySession) UseDELEGATECALL() (bool, error) {
	return _Proxy.Contract.UseDELEGATECALL(&_Proxy.CallOpts)
}

// UseDELEGATECALL is a free data retrieval call binding the contract method 0x95578ebd.
//
// Solidity: function useDELEGATECALL() view returns(bool)
func (_Proxy *ProxyCallerSession) UseDELEGATECALL() (bool, error) {
	return _Proxy.Contract.UseDELEGATECALL(&_Proxy.CallOpts)
}

// Emit is a paid mutator transaction binding the contract method 0x907dff97.
//
// Solidity: function _emit(bytes callData, uint256 numTopics, bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4) returns()
func (_Proxy *ProxyTransactor) Emit(opts *bind.TransactOpts, callData []byte, numTopics *big.Int, topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "_emit", callData, numTopics, topic1, topic2, topic3, topic4)
}

// Emit is a paid mutator transaction binding the contract method 0x907dff97.
//
// Solidity: function _emit(bytes callData, uint256 numTopics, bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4) returns()
func (_Proxy *ProxySession) Emit(callData []byte, numTopics *big.Int, topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte) (*types.Transaction, error) {
	return _Proxy.Contract.Emit(&_Proxy.TransactOpts, callData, numTopics, topic1, topic2, topic3, topic4)
}

// Emit is a paid mutator transaction binding the contract method 0x907dff97.
//
// Solidity: function _emit(bytes callData, uint256 numTopics, bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4) returns()
func (_Proxy *ProxyTransactorSession) Emit(callData []byte, numTopics *big.Int, topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte) (*types.Transaction, error) {
	return _Proxy.Contract.Emit(&_Proxy.TransactOpts, callData, numTopics, topic1, topic2, topic3, topic4)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Proxy *ProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Proxy *ProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _Proxy.Contract.AcceptOwnership(&_Proxy.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Proxy *ProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Proxy.Contract.AcceptOwnership(&_Proxy.TransactOpts)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Proxy *ProxyTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Proxy *ProxySession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.NominateNewOwner(&_Proxy.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Proxy *ProxyTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.NominateNewOwner(&_Proxy.TransactOpts, _owner)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address _target) returns()
func (_Proxy *ProxyTransactor) SetTarget(opts *bind.TransactOpts, _target common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setTarget", _target)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address _target) returns()
func (_Proxy *ProxySession) SetTarget(_target common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetTarget(&_Proxy.TransactOpts, _target)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address _target) returns()
func (_Proxy *ProxyTransactorSession) SetTarget(_target common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetTarget(&_Proxy.TransactOpts, _target)
}

// SetUseDELEGATECALL is a paid mutator transaction binding the contract method 0xbefff6af.
//
// Solidity: function setUseDELEGATECALL(bool value) returns()
func (_Proxy *ProxyTransactor) SetUseDELEGATECALL(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setUseDELEGATECALL", value)
}

// SetUseDELEGATECALL is a paid mutator transaction binding the contract method 0xbefff6af.
//
// Solidity: function setUseDELEGATECALL(bool value) returns()
func (_Proxy *ProxySession) SetUseDELEGATECALL(value bool) (*types.Transaction, error) {
	return _Proxy.Contract.SetUseDELEGATECALL(&_Proxy.TransactOpts, value)
}

// SetUseDELEGATECALL is a paid mutator transaction binding the contract method 0xbefff6af.
//
// Solidity: function setUseDELEGATECALL(bool value) returns()
func (_Proxy *ProxyTransactorSession) SetUseDELEGATECALL(value bool) (*types.Transaction, error) {
	return _Proxy.Contract.SetUseDELEGATECALL(&_Proxy.TransactOpts, value)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Proxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

// ProxyOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Proxy contract.
type ProxyOwnerChangedIterator struct {
	Event *ProxyOwnerChanged // Event containing the contract specifics and raw log

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
func (it *ProxyOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyOwnerChanged)
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
		it.Event = new(ProxyOwnerChanged)
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
func (it *ProxyOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyOwnerChanged represents a OwnerChanged event raised by the Proxy contract.
type ProxyOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Proxy *ProxyFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*ProxyOwnerChangedIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &ProxyOwnerChangedIterator{contract: _Proxy.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Proxy *ProxyFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *ProxyOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyOwnerChanged)
				if err := _Proxy.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Proxy *ProxyFilterer) ParseOwnerChanged(log types.Log) (*ProxyOwnerChanged, error) {
	event := new(ProxyOwnerChanged)
	if err := _Proxy.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the Proxy contract.
type ProxyOwnerNominatedIterator struct {
	Event *ProxyOwnerNominated // Event containing the contract specifics and raw log

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
func (it *ProxyOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyOwnerNominated)
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
		it.Event = new(ProxyOwnerNominated)
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
func (it *ProxyOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyOwnerNominated represents a OwnerNominated event raised by the Proxy contract.
type ProxyOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Proxy *ProxyFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*ProxyOwnerNominatedIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &ProxyOwnerNominatedIterator{contract: _Proxy.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Proxy *ProxyFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *ProxyOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyOwnerNominated)
				if err := _Proxy.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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

// ParseOwnerNominated is a log parse operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Proxy *ProxyFilterer) ParseOwnerNominated(log types.Log) (*ProxyOwnerNominated, error) {
	event := new(ProxyOwnerNominated)
	if err := _Proxy.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyTargetUpdatedIterator is returned from FilterTargetUpdated and is used to iterate over the raw logs and unpacked data for TargetUpdated events raised by the Proxy contract.
type ProxyTargetUpdatedIterator struct {
	Event *ProxyTargetUpdated // Event containing the contract specifics and raw log

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
func (it *ProxyTargetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyTargetUpdated)
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
		it.Event = new(ProxyTargetUpdated)
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
func (it *ProxyTargetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyTargetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyTargetUpdated represents a TargetUpdated event raised by the Proxy contract.
type ProxyTargetUpdated struct {
	NewTarget common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTargetUpdated is a free log retrieval operation binding the contract event 0x814250a3b8c79fcbe2ead2c131c952a278491c8f4322a79fe84b5040a810373e.
//
// Solidity: event TargetUpdated(address newTarget)
func (_Proxy *ProxyFilterer) FilterTargetUpdated(opts *bind.FilterOpts) (*ProxyTargetUpdatedIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "TargetUpdated")
	if err != nil {
		return nil, err
	}
	return &ProxyTargetUpdatedIterator{contract: _Proxy.contract, event: "TargetUpdated", logs: logs, sub: sub}, nil
}

// WatchTargetUpdated is a free log subscription operation binding the contract event 0x814250a3b8c79fcbe2ead2c131c952a278491c8f4322a79fe84b5040a810373e.
//
// Solidity: event TargetUpdated(address newTarget)
func (_Proxy *ProxyFilterer) WatchTargetUpdated(opts *bind.WatchOpts, sink chan<- *ProxyTargetUpdated) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "TargetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyTargetUpdated)
				if err := _Proxy.contract.UnpackLog(event, "TargetUpdated", log); err != nil {
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

// ParseTargetUpdated is a log parse operation binding the contract event 0x814250a3b8c79fcbe2ead2c131c952a278491c8f4322a79fe84b5040a810373e.
//
// Solidity: event TargetUpdated(address newTarget)
func (_Proxy *ProxyFilterer) ParseTargetUpdated(log types.Log) (*ProxyTargetUpdated, error) {
	event := new(ProxyTargetUpdated)
	if err := _Proxy.contract.UnpackLog(event, "TargetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyERC20ABI is the input ABI used to generate the binding from.
const ProxyERC20ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"setTarget\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"callData\",\"type\":\"bytes\"},{\"name\":\"numTopics\",\"type\":\"uint256\"},{\"name\":\"topic1\",\"type\":\"bytes32\"},{\"name\":\"topic2\",\"type\":\"bytes32\"},{\"name\":\"topic3\",\"type\":\"bytes32\"},{\"name\":\"topic4\",\"type\":\"bytes32\"}],\"name\":\"_emit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"useDELEGATECALL\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setUseDELEGATECALL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"target\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newTarget\",\"type\":\"address\"}],\"name\":\"TargetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"}]"

// ProxyERC20FuncSigs maps the 4-byte function signature to its string representation.
var ProxyERC20FuncSigs = map[string]string{
	"907dff97": "_emit(bytes,uint256,bytes32,bytes32,bytes32,bytes32)",
	"79ba5097": "acceptOwnership()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"1627540c": "nominateNewOwner(address)",
	"53a47bb7": "nominatedOwner()",
	"8da5cb5b": "owner()",
	"776d1a01": "setTarget(address)",
	"befff6af": "setUseDELEGATECALL(bool)",
	"95d89b41": "symbol()",
	"d4b83992": "target()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"95578ebd": "useDELEGATECALL()",
}

// ProxyERC20Bin is the compiled bytecode used for deploying new contracts.
var ProxyERC20Bin = "0x608060405234801561001057600080fd5b5060405160208061114583398101604052518080600160a060020a038116151561009b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e657220616464726573732063616e6e6f74206265203000000000000000604482015290519081900360640190fd5b60008054600160a060020a031916600160a060020a038316908117825560408051928352602083019190915280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a1505050611040806101056000396000f3006080604052600436106100fb5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146101d0578063095ea7b31461025a5780631627540c1461029257806318160ddd146102b357806323b872dd146102da578063313ce5671461030457806353a47bb71461032f57806370a0823114610360578063776d1a011461038157806379ba5097146103a25780638da5cb5b146103b7578063907dff97146103cc57806395578ebd146103fc57806395d89b4114610411578063a9059cbb14610426578063befff6af1461044a578063d4b8399214610464578063dd62ed3e14610479575b60025474010000000000000000000000000000000000000000900460ff161561014457604051366000823760008036836002545af43d6000833e801515610140573d82fd5b3d82f35b6002546040805160e160020a635e33fc190281523360048201529051600160a060020a039092169163bc67f8329160248082019260009290919082900301818387803b15801561019357600080fd5b505af11580156101a7573d6000803e3d6000fd5b5050505060405136600082376000803683346002545af13d6000833e801515610140573d82fd5b005b3480156101dc57600080fd5b506101e56104a0565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561021f578181015183820152602001610207565b50505050905090810190601f16801561024c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561026657600080fd5b5061027e600160a060020a036004351660243561058c565b604080519115158252519081900360200190f35b34801561029e57600080fd5b506101ce600160a060020a0360043516610696565b3480156102bf57600080fd5b506102c861077f565b60408051918252519081900360200190f35b3480156102e657600080fd5b5061027e600160a060020a036004358116906024351660443561080f565b34801561031057600080fd5b50610319610922565b6040805160ff9092168252519081900360200190f35b34801561033b57600080fd5b50610344610981565b60408051600160a060020a039092168252519081900360200190f35b34801561036c57600080fd5b506102c8600160a060020a0360043516610990565b34801561038d57600080fd5b506101ce600160a060020a0360043516610a2d565b3480156103ae57600080fd5b506101ce610b16565b3480156103c357600080fd5b50610344610c1e565b3480156103d857600080fd5b506101ce60246004803582810192910135903560443560643560843560a435610c2d565b34801561040857600080fd5b5061027e610d42565b34801561041d57600080fd5b506101e5610d63565b34801561043257600080fd5b5061027e600160a060020a0360043516602435610dc2565b34801561045657600080fd5b506101ce6004351515610e97565b34801561047057600080fd5b50610344610f5f565b34801561048557600080fd5b506102c8600160a060020a0360043581169060243516610f6e565b600254604080517f06fdde030000000000000000000000000000000000000000000000000000000081529051606092600160a060020a0316916306fdde0391600480830192600092919082900301818387803b1580156104ff57600080fd5b505af1158015610513573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561053c57600080fd5b81019080805164010000000081111561055457600080fd5b8201602081018481111561056757600080fd5b815164010000000081118282018710171561058157600080fd5b509094505050505090565b6002546040805160e160020a635e33fc190281523360048201529051600092600160a060020a03169163bc67f832916024808301928692919082900301818387803b1580156105da57600080fd5b505af11580156105ee573d6000803e3d6000fd5b5050600254604080517f095ea7b3000000000000000000000000000000000000000000000000000000008152600160a060020a03888116600483015260248201889052915191909216935063095ea7b3925060448083019260209291908290030181600087803b15801561066157600080fd5b505af1158015610675573d6000803e3d6000fd5b505050506040513d602081101561068b57600080fd5b506001949350505050565b600054600160a060020a0316331461071e576040805160e560020a62461bcd02815260206004820152602f60248201527f4f6e6c792074686520636f6e7472616374206f776e6572206d6179207065726660448201527f6f726d207468697320616374696f6e0000000000000000000000000000000000606482015290519081900360840190fd5b60018054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517f906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce229181900360200190a150565b600254604080517f18160ddd0000000000000000000000000000000000000000000000000000000081529051600092600160a060020a0316916318160ddd91600480830192602092919082900301818787803b1580156107de57600080fd5b505af11580156107f2573d6000803e3d6000fd5b505050506040513d602081101561080857600080fd5b5051905090565b6002546040805160e160020a635e33fc190281523360048201529051600092600160a060020a03169163bc67f832916024808301928692919082900301818387803b15801561085d57600080fd5b505af1158015610871573d6000803e3d6000fd5b5050600254604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a03898116600483015288811660248301526044820188905291519190921693506323b872dd925060648083019260209291908290030181600087803b1580156108ec57600080fd5b505af1158015610900573d6000803e3d6000fd5b505050506040513d602081101561091657600080fd5b50600195945050505050565b600254604080517f313ce5670000000000000000000000000000000000000000000000000000000081529051600092600160a060020a03169163313ce56791600480830192602092919082900301818787803b1580156107de57600080fd5b600154600160a060020a031681565b600254604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915160009392909216916370a082319160248082019260209290919082900301818787803b1580156109fb57600080fd5b505af1158015610a0f573d6000803e3d6000fd5b505050506040513d6020811015610a2557600080fd5b505192915050565b600054600160a060020a03163314610ab5576040805160e560020a62461bcd02815260206004820152602f60248201527f4f6e6c792074686520636f6e7472616374206f776e6572206d6179207065726660448201527f6f726d207468697320616374696f6e0000000000000000000000000000000000606482015290519081900360840190fd5b60028054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517f814250a3b8c79fcbe2ead2c131c952a278491c8f4322a79fe84b5040a810373e9181900360200190a150565b600154600160a060020a03163314610b9e576040805160e560020a62461bcd02815260206004820152603560248201527f596f75206d757374206265206e6f6d696e61746564206265666f726520796f7560448201527f2063616e20616363657074206f776e6572736869700000000000000000000000606482015290519081900360840190fd5b60005460015460408051600160a060020a03938416815292909116602083015280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a1600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a031681565b600254600090606090600160a060020a03163314610c95576040805160e560020a62461bcd02815260206004820152601460248201527f4d7573742062652070726f787920746172676574000000000000000000000000604482015290519081900360640190fd5b604080516020601f8b01819004810282018101909252898152899350908a9084908190840183828082843782019150505050505090508660008114610cf95760018114610d045760028114610d105760038114610d1d5760048114610d2b57610d36565b8260208301a0610d36565b868360208401a1610d36565b85878460208501a2610d36565b8486888560208601a3610d36565b838587898660208701a45b50505050505050505050565b60025474010000000000000000000000000000000000000000900460ff1681565b600254604080517f95d89b410000000000000000000000000000000000000000000000000000000081529051606092600160a060020a0316916395d89b4191600480830192600092919082900301818387803b1580156104ff57600080fd5b6002546040805160e160020a635e33fc190281523360048201529051600092600160a060020a03169163bc67f832916024808301928692919082900301818387803b158015610e1057600080fd5b505af1158015610e24573d6000803e3d6000fd5b5050600254604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a03888116600483015260248201889052915191909216935063a9059cbb925060448083019260209291908290030181600087803b15801561066157600080fd5b600054600160a060020a03163314610f1f576040805160e560020a62461bcd02815260206004820152602f60248201527f4f6e6c792074686520636f6e7472616374206f776e6572206d6179207065726660448201527f6f726d207468697320616374696f6e0000000000000000000000000000000000606482015290519081900360840190fd5b60028054911515740100000000000000000000000000000000000000000274ff000000000000000000000000000000000000000019909216919091179055565b600254600160a060020a031681565b600254604080517fdd62ed3e000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015284811660248301529151600093929092169163dd62ed3e9160448082019260209290919082900301818787803b158015610fe157600080fd5b505af1158015610ff5573d6000803e3d6000fd5b505050506040513d602081101561100b57600080fd5b505193925050505600a165627a7a72305820a75667782854283dd051460d9f3f87d9de3a1d40f72d0585ec97b62c9b0dd1d30029"

// DeployProxyERC20 deploys a new Ethereum contract, binding an instance of ProxyERC20 to it.
func DeployProxyERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *ProxyERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyERC20Bin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProxyERC20{ProxyERC20Caller: ProxyERC20Caller{contract: contract}, ProxyERC20Transactor: ProxyERC20Transactor{contract: contract}, ProxyERC20Filterer: ProxyERC20Filterer{contract: contract}}, nil
}

// ProxyERC20 is an auto generated Go binding around an Ethereum contract.
type ProxyERC20 struct {
	ProxyERC20Caller     // Read-only binding to the contract
	ProxyERC20Transactor // Write-only binding to the contract
	ProxyERC20Filterer   // Log filterer for contract events
}

// ProxyERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyERC20Session struct {
	Contract     *ProxyERC20       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyERC20CallerSession struct {
	Contract *ProxyERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ProxyERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyERC20TransactorSession struct {
	Contract     *ProxyERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProxyERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyERC20Raw struct {
	Contract *ProxyERC20 // Generic contract binding to access the raw methods on
}

// ProxyERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyERC20CallerRaw struct {
	Contract *ProxyERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ProxyERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyERC20TransactorRaw struct {
	Contract *ProxyERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyERC20 creates a new instance of ProxyERC20, bound to a specific deployed contract.
func NewProxyERC20(address common.Address, backend bind.ContractBackend) (*ProxyERC20, error) {
	contract, err := bindProxyERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyERC20{ProxyERC20Caller: ProxyERC20Caller{contract: contract}, ProxyERC20Transactor: ProxyERC20Transactor{contract: contract}, ProxyERC20Filterer: ProxyERC20Filterer{contract: contract}}, nil
}

// NewProxyERC20Caller creates a new read-only instance of ProxyERC20, bound to a specific deployed contract.
func NewProxyERC20Caller(address common.Address, caller bind.ContractCaller) (*ProxyERC20Caller, error) {
	contract, err := bindProxyERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyERC20Caller{contract: contract}, nil
}

// NewProxyERC20Transactor creates a new write-only instance of ProxyERC20, bound to a specific deployed contract.
func NewProxyERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ProxyERC20Transactor, error) {
	contract, err := bindProxyERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyERC20Transactor{contract: contract}, nil
}

// NewProxyERC20Filterer creates a new log filterer instance of ProxyERC20, bound to a specific deployed contract.
func NewProxyERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ProxyERC20Filterer, error) {
	contract, err := bindProxyERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyERC20Filterer{contract: contract}, nil
}

// bindProxyERC20 binds a generic wrapper to an already deployed contract.
func bindProxyERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyERC20 *ProxyERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyERC20.Contract.ProxyERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyERC20 *ProxyERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyERC20.Contract.ProxyERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyERC20 *ProxyERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyERC20.Contract.ProxyERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyERC20 *ProxyERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyERC20 *ProxyERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyERC20 *ProxyERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ProxyERC20 *ProxyERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ProxyERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ProxyERC20 *ProxyERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ProxyERC20.Contract.Allowance(&_ProxyERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ProxyERC20 *ProxyERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ProxyERC20.Contract.Allowance(&_ProxyERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ProxyERC20 *ProxyERC20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ProxyERC20.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ProxyERC20 *ProxyERC20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ProxyERC20.Contract.BalanceOf(&_ProxyERC20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ProxyERC20 *ProxyERC20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ProxyERC20.Contract.BalanceOf(&_ProxyERC20.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ProxyERC20 *ProxyERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ProxyERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ProxyERC20 *ProxyERC20Session) Decimals() (uint8, error) {
	return _ProxyERC20.Contract.Decimals(&_ProxyERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ProxyERC20 *ProxyERC20CallerSession) Decimals() (uint8, error) {
	return _ProxyERC20.Contract.Decimals(&_ProxyERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ProxyERC20 *ProxyERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProxyERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ProxyERC20 *ProxyERC20Session) Name() (string, error) {
	return _ProxyERC20.Contract.Name(&_ProxyERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ProxyERC20 *ProxyERC20CallerSession) Name() (string, error) {
	return _ProxyERC20.Contract.Name(&_ProxyERC20.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_ProxyERC20 *ProxyERC20Caller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProxyERC20.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_ProxyERC20 *ProxyERC20Session) NominatedOwner() (common.Address, error) {
	return _ProxyERC20.Contract.NominatedOwner(&_ProxyERC20.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_ProxyERC20 *ProxyERC20CallerSession) NominatedOwner() (common.Address, error) {
	return _ProxyERC20.Contract.NominatedOwner(&_ProxyERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyERC20 *ProxyERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProxyERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyERC20 *ProxyERC20Session) Owner() (common.Address, error) {
	return _ProxyERC20.Contract.Owner(&_ProxyERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyERC20 *ProxyERC20CallerSession) Owner() (common.Address, error) {
	return _ProxyERC20.Contract.Owner(&_ProxyERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ProxyERC20 *ProxyERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProxyERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ProxyERC20 *ProxyERC20Session) Symbol() (string, error) {
	return _ProxyERC20.Contract.Symbol(&_ProxyERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ProxyERC20 *ProxyERC20CallerSession) Symbol() (string, error) {
	return _ProxyERC20.Contract.Symbol(&_ProxyERC20.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_ProxyERC20 *ProxyERC20Caller) Target(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProxyERC20.contract.Call(opts, &out, "target")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_ProxyERC20 *ProxyERC20Session) Target() (common.Address, error) {
	return _ProxyERC20.Contract.Target(&_ProxyERC20.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_ProxyERC20 *ProxyERC20CallerSession) Target() (common.Address, error) {
	return _ProxyERC20.Contract.Target(&_ProxyERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ProxyERC20 *ProxyERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProxyERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ProxyERC20 *ProxyERC20Session) TotalSupply() (*big.Int, error) {
	return _ProxyERC20.Contract.TotalSupply(&_ProxyERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ProxyERC20 *ProxyERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ProxyERC20.Contract.TotalSupply(&_ProxyERC20.CallOpts)
}

// UseDELEGATECALL is a free data retrieval call binding the contract method 0x95578ebd.
//
// Solidity: function useDELEGATECALL() view returns(bool)
func (_ProxyERC20 *ProxyERC20Caller) UseDELEGATECALL(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ProxyERC20.contract.Call(opts, &out, "useDELEGATECALL")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDELEGATECALL is a free data retrieval call binding the contract method 0x95578ebd.
//
// Solidity: function useDELEGATECALL() view returns(bool)
func (_ProxyERC20 *ProxyERC20Session) UseDELEGATECALL() (bool, error) {
	return _ProxyERC20.Contract.UseDELEGATECALL(&_ProxyERC20.CallOpts)
}

// UseDELEGATECALL is a free data retrieval call binding the contract method 0x95578ebd.
//
// Solidity: function useDELEGATECALL() view returns(bool)
func (_ProxyERC20 *ProxyERC20CallerSession) UseDELEGATECALL() (bool, error) {
	return _ProxyERC20.Contract.UseDELEGATECALL(&_ProxyERC20.CallOpts)
}

// Emit is a paid mutator transaction binding the contract method 0x907dff97.
//
// Solidity: function _emit(bytes callData, uint256 numTopics, bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4) returns()
func (_ProxyERC20 *ProxyERC20Transactor) Emit(opts *bind.TransactOpts, callData []byte, numTopics *big.Int, topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte) (*types.Transaction, error) {
	return _ProxyERC20.contract.Transact(opts, "_emit", callData, numTopics, topic1, topic2, topic3, topic4)
}

// Emit is a paid mutator transaction binding the contract method 0x907dff97.
//
// Solidity: function _emit(bytes callData, uint256 numTopics, bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4) returns()
func (_ProxyERC20 *ProxyERC20Session) Emit(callData []byte, numTopics *big.Int, topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte) (*types.Transaction, error) {
	return _ProxyERC20.Contract.Emit(&_ProxyERC20.TransactOpts, callData, numTopics, topic1, topic2, topic3, topic4)
}

// Emit is a paid mutator transaction binding the contract method 0x907dff97.
//
// Solidity: function _emit(bytes callData, uint256 numTopics, bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4) returns()
func (_ProxyERC20 *ProxyERC20TransactorSession) Emit(callData []byte, numTopics *big.Int, topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte) (*types.Transaction, error) {
	return _ProxyERC20.Contract.Emit(&_ProxyERC20.TransactOpts, callData, numTopics, topic1, topic2, topic3, topic4)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ProxyERC20 *ProxyERC20Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyERC20.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ProxyERC20 *ProxyERC20Session) AcceptOwnership() (*types.Transaction, error) {
	return _ProxyERC20.Contract.AcceptOwnership(&_ProxyERC20.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ProxyERC20 *ProxyERC20TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ProxyERC20.Contract.AcceptOwnership(&_ProxyERC20.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ProxyERC20 *ProxyERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ProxyERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ProxyERC20 *ProxyERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ProxyERC20.Contract.Approve(&_ProxyERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ProxyERC20 *ProxyERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ProxyERC20.Contract.Approve(&_ProxyERC20.TransactOpts, spender, value)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_ProxyERC20 *ProxyERC20Transactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _ProxyERC20.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_ProxyERC20 *ProxyERC20Session) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _ProxyERC20.Contract.NominateNewOwner(&_ProxyERC20.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_ProxyERC20 *ProxyERC20TransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _ProxyERC20.Contract.NominateNewOwner(&_ProxyERC20.TransactOpts, _owner)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address _target) returns()
func (_ProxyERC20 *ProxyERC20Transactor) SetTarget(opts *bind.TransactOpts, _target common.Address) (*types.Transaction, error) {
	return _ProxyERC20.contract.Transact(opts, "setTarget", _target)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address _target) returns()
func (_ProxyERC20 *ProxyERC20Session) SetTarget(_target common.Address) (*types.Transaction, error) {
	return _ProxyERC20.Contract.SetTarget(&_ProxyERC20.TransactOpts, _target)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address _target) returns()
func (_ProxyERC20 *ProxyERC20TransactorSession) SetTarget(_target common.Address) (*types.Transaction, error) {
	return _ProxyERC20.Contract.SetTarget(&_ProxyERC20.TransactOpts, _target)
}

// SetUseDELEGATECALL is a paid mutator transaction binding the contract method 0xbefff6af.
//
// Solidity: function setUseDELEGATECALL(bool value) returns()
func (_ProxyERC20 *ProxyERC20Transactor) SetUseDELEGATECALL(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ProxyERC20.contract.Transact(opts, "setUseDELEGATECALL", value)
}

// SetUseDELEGATECALL is a paid mutator transaction binding the contract method 0xbefff6af.
//
// Solidity: function setUseDELEGATECALL(bool value) returns()
func (_ProxyERC20 *ProxyERC20Session) SetUseDELEGATECALL(value bool) (*types.Transaction, error) {
	return _ProxyERC20.Contract.SetUseDELEGATECALL(&_ProxyERC20.TransactOpts, value)
}

// SetUseDELEGATECALL is a paid mutator transaction binding the contract method 0xbefff6af.
//
// Solidity: function setUseDELEGATECALL(bool value) returns()
func (_ProxyERC20 *ProxyERC20TransactorSession) SetUseDELEGATECALL(value bool) (*types.Transaction, error) {
	return _ProxyERC20.Contract.SetUseDELEGATECALL(&_ProxyERC20.TransactOpts, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ProxyERC20 *ProxyERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ProxyERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ProxyERC20 *ProxyERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ProxyERC20.Contract.Transfer(&_ProxyERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ProxyERC20 *ProxyERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ProxyERC20.Contract.Transfer(&_ProxyERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ProxyERC20 *ProxyERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ProxyERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ProxyERC20 *ProxyERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ProxyERC20.Contract.TransferFrom(&_ProxyERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ProxyERC20 *ProxyERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ProxyERC20.Contract.TransferFrom(&_ProxyERC20.TransactOpts, from, to, value)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ProxyERC20 *ProxyERC20Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ProxyERC20.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ProxyERC20 *ProxyERC20Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ProxyERC20.Contract.Fallback(&_ProxyERC20.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ProxyERC20 *ProxyERC20TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ProxyERC20.Contract.Fallback(&_ProxyERC20.TransactOpts, calldata)
}

// ProxyERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ProxyERC20 contract.
type ProxyERC20ApprovalIterator struct {
	Event *ProxyERC20Approval // Event containing the contract specifics and raw log

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
func (it *ProxyERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyERC20Approval)
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
		it.Event = new(ProxyERC20Approval)
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
func (it *ProxyERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyERC20Approval represents a Approval event raised by the ProxyERC20 contract.
type ProxyERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ProxyERC20 *ProxyERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ProxyERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ProxyERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ProxyERC20ApprovalIterator{contract: _ProxyERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ProxyERC20 *ProxyERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ProxyERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ProxyERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyERC20Approval)
				if err := _ProxyERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ProxyERC20 *ProxyERC20Filterer) ParseApproval(log types.Log) (*ProxyERC20Approval, error) {
	event := new(ProxyERC20Approval)
	if err := _ProxyERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyERC20OwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the ProxyERC20 contract.
type ProxyERC20OwnerChangedIterator struct {
	Event *ProxyERC20OwnerChanged // Event containing the contract specifics and raw log

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
func (it *ProxyERC20OwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyERC20OwnerChanged)
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
		it.Event = new(ProxyERC20OwnerChanged)
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
func (it *ProxyERC20OwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyERC20OwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyERC20OwnerChanged represents a OwnerChanged event raised by the ProxyERC20 contract.
type ProxyERC20OwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_ProxyERC20 *ProxyERC20Filterer) FilterOwnerChanged(opts *bind.FilterOpts) (*ProxyERC20OwnerChangedIterator, error) {

	logs, sub, err := _ProxyERC20.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &ProxyERC20OwnerChangedIterator{contract: _ProxyERC20.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_ProxyERC20 *ProxyERC20Filterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *ProxyERC20OwnerChanged) (event.Subscription, error) {

	logs, sub, err := _ProxyERC20.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyERC20OwnerChanged)
				if err := _ProxyERC20.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_ProxyERC20 *ProxyERC20Filterer) ParseOwnerChanged(log types.Log) (*ProxyERC20OwnerChanged, error) {
	event := new(ProxyERC20OwnerChanged)
	if err := _ProxyERC20.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyERC20OwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the ProxyERC20 contract.
type ProxyERC20OwnerNominatedIterator struct {
	Event *ProxyERC20OwnerNominated // Event containing the contract specifics and raw log

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
func (it *ProxyERC20OwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyERC20OwnerNominated)
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
		it.Event = new(ProxyERC20OwnerNominated)
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
func (it *ProxyERC20OwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyERC20OwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyERC20OwnerNominated represents a OwnerNominated event raised by the ProxyERC20 contract.
type ProxyERC20OwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_ProxyERC20 *ProxyERC20Filterer) FilterOwnerNominated(opts *bind.FilterOpts) (*ProxyERC20OwnerNominatedIterator, error) {

	logs, sub, err := _ProxyERC20.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &ProxyERC20OwnerNominatedIterator{contract: _ProxyERC20.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_ProxyERC20 *ProxyERC20Filterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *ProxyERC20OwnerNominated) (event.Subscription, error) {

	logs, sub, err := _ProxyERC20.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyERC20OwnerNominated)
				if err := _ProxyERC20.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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

// ParseOwnerNominated is a log parse operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_ProxyERC20 *ProxyERC20Filterer) ParseOwnerNominated(log types.Log) (*ProxyERC20OwnerNominated, error) {
	event := new(ProxyERC20OwnerNominated)
	if err := _ProxyERC20.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyERC20TargetUpdatedIterator is returned from FilterTargetUpdated and is used to iterate over the raw logs and unpacked data for TargetUpdated events raised by the ProxyERC20 contract.
type ProxyERC20TargetUpdatedIterator struct {
	Event *ProxyERC20TargetUpdated // Event containing the contract specifics and raw log

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
func (it *ProxyERC20TargetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyERC20TargetUpdated)
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
		it.Event = new(ProxyERC20TargetUpdated)
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
func (it *ProxyERC20TargetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyERC20TargetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyERC20TargetUpdated represents a TargetUpdated event raised by the ProxyERC20 contract.
type ProxyERC20TargetUpdated struct {
	NewTarget common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTargetUpdated is a free log retrieval operation binding the contract event 0x814250a3b8c79fcbe2ead2c131c952a278491c8f4322a79fe84b5040a810373e.
//
// Solidity: event TargetUpdated(address newTarget)
func (_ProxyERC20 *ProxyERC20Filterer) FilterTargetUpdated(opts *bind.FilterOpts) (*ProxyERC20TargetUpdatedIterator, error) {

	logs, sub, err := _ProxyERC20.contract.FilterLogs(opts, "TargetUpdated")
	if err != nil {
		return nil, err
	}
	return &ProxyERC20TargetUpdatedIterator{contract: _ProxyERC20.contract, event: "TargetUpdated", logs: logs, sub: sub}, nil
}

// WatchTargetUpdated is a free log subscription operation binding the contract event 0x814250a3b8c79fcbe2ead2c131c952a278491c8f4322a79fe84b5040a810373e.
//
// Solidity: event TargetUpdated(address newTarget)
func (_ProxyERC20 *ProxyERC20Filterer) WatchTargetUpdated(opts *bind.WatchOpts, sink chan<- *ProxyERC20TargetUpdated) (event.Subscription, error) {

	logs, sub, err := _ProxyERC20.contract.WatchLogs(opts, "TargetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyERC20TargetUpdated)
				if err := _ProxyERC20.contract.UnpackLog(event, "TargetUpdated", log); err != nil {
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

// ParseTargetUpdated is a log parse operation binding the contract event 0x814250a3b8c79fcbe2ead2c131c952a278491c8f4322a79fe84b5040a810373e.
//
// Solidity: event TargetUpdated(address newTarget)
func (_ProxyERC20 *ProxyERC20Filterer) ParseTargetUpdated(log types.Log) (*ProxyERC20TargetUpdated, error) {
	event := new(ProxyERC20TargetUpdated)
	if err := _ProxyERC20.contract.UnpackLog(event, "TargetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ProxyERC20 contract.
type ProxyERC20TransferIterator struct {
	Event *ProxyERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ProxyERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyERC20Transfer)
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
		it.Event = new(ProxyERC20Transfer)
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
func (it *ProxyERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyERC20Transfer represents a Transfer event raised by the ProxyERC20 contract.
type ProxyERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ProxyERC20 *ProxyERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ProxyERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ProxyERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ProxyERC20TransferIterator{contract: _ProxyERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ProxyERC20 *ProxyERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ProxyERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ProxyERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyERC20Transfer)
				if err := _ProxyERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ProxyERC20 *ProxyERC20Filterer) ParseTransfer(log types.Log) (*ProxyERC20Transfer, error) {
	event := new(ProxyERC20Transfer)
	if err := _ProxyERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyableABI is the input ABI used to generate the binding from.
const ProxyableABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_integrationProxy\",\"type\":\"address\"}],\"name\":\"setIntegrationProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"setProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"integrationProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"setMessageSender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"proxyAddress\",\"type\":\"address\"}],\"name\":\"ProxyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"}]"

// ProxyableFuncSigs maps the 4-byte function signature to its string representation.
var ProxyableFuncSigs = map[string]string{
	"79ba5097": "acceptOwnership()",
	"9cbdaeb6": "integrationProxy()",
	"1627540c": "nominateNewOwner(address)",
	"53a47bb7": "nominatedOwner()",
	"8da5cb5b": "owner()",
	"ec556889": "proxy()",
	"131b0ae7": "setIntegrationProxy(address)",
	"bc67f832": "setMessageSender(address)",
	"97107d6d": "setProxy(address)",
}

// ProxyableBin is the compiled bytecode used for deploying new contracts.
var ProxyableBin = "0x608060405234801561001057600080fd5b506040516040806107a583398101604052805160209091015180600160a060020a03811615156100a157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e657220616464726573732063616e6e6f74206265203000000000000000604482015290519081900360640190fd5b60008054600160a060020a031916600160a060020a038316908117825560408051928352602083019190915280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a15060028054600160a060020a038416600160a060020a0319909116811790915560408051918252517ffc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e9181900360200190a150506106498061015c6000396000f3006080604052600436106100985763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663131b0ae7811461009d5780631627540c146100c057806353a47bb7146100e157806379ba5097146101125780638da5cb5b1461012757806397107d6d1461013c5780639cbdaeb61461015d578063bc67f83214610172578063ec55688914610193575b600080fd5b3480156100a957600080fd5b506100be600160a060020a03600435166101a8565b005b3480156100cc57600080fd5b506100be600160a060020a036004351661025f565b3480156100ed57600080fd5b506100f6610348565b60408051600160a060020a039092168252519081900360200190f35b34801561011e57600080fd5b506100be610357565b34801561013357600080fd5b506100f661045f565b34801561014857600080fd5b506100be600160a060020a036004351661046e565b34801561016957600080fd5b506100f6610557565b34801561017e57600080fd5b506100be600160a060020a0360043516610566565b34801561019f57600080fd5b506100f661060e565b600054600160a060020a03163314610230576040805160e560020a62461bcd02815260206004820152602f60248201527f4f6e6c792074686520636f6e7472616374206f776e6572206d6179207065726660448201527f6f726d207468697320616374696f6e0000000000000000000000000000000000606482015290519081900360840190fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031633146102e7576040805160e560020a62461bcd02815260206004820152602f60248201527f4f6e6c792074686520636f6e7472616374206f776e6572206d6179207065726660448201527f6f726d207468697320616374696f6e0000000000000000000000000000000000606482015290519081900360840190fd5b60018054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517f906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce229181900360200190a150565b600154600160a060020a031681565b600154600160a060020a031633146103df576040805160e560020a62461bcd02815260206004820152603560248201527f596f75206d757374206265206e6f6d696e61746564206265666f726520796f7560448201527f2063616e20616363657074206f776e6572736869700000000000000000000000606482015290519081900360840190fd5b60005460015460408051600160a060020a03938416815292909116602083015280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a1600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a031681565b600054600160a060020a031633146104f6576040805160e560020a62461bcd02815260206004820152602f60248201527f4f6e6c792074686520636f6e7472616374206f776e6572206d6179207065726660448201527f6f726d207468697320616374696f6e0000000000000000000000000000000000606482015290519081900360840190fd5b60028054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517ffc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e9181900360200190a150565b600354600160a060020a031681565b600254600160a060020a03163314806105895750600354600160a060020a031633145b15156105df576040805160e560020a62461bcd02815260206004820152601760248201527f4f6e6c79207468652070726f78792063616e2063616c6c000000000000000000604482015290519081900360640190fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a0316815600a165627a7a72305820845c5c971c249e9bc06e91d595527a13650fc12fda3caaafff0690aaa1b7a7df0029"

// DeployProxyable deploys a new Ethereum contract, binding an instance of Proxyable to it.
func DeployProxyable(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address, _owner common.Address) (common.Address, *types.Transaction, *Proxyable, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyableBin), backend, _proxy, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proxyable{ProxyableCaller: ProxyableCaller{contract: contract}, ProxyableTransactor: ProxyableTransactor{contract: contract}, ProxyableFilterer: ProxyableFilterer{contract: contract}}, nil
}

// Proxyable is an auto generated Go binding around an Ethereum contract.
type Proxyable struct {
	ProxyableCaller     // Read-only binding to the contract
	ProxyableTransactor // Write-only binding to the contract
	ProxyableFilterer   // Log filterer for contract events
}

// ProxyableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyableSession struct {
	Contract     *Proxyable        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyableCallerSession struct {
	Contract *ProxyableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ProxyableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyableTransactorSession struct {
	Contract     *ProxyableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ProxyableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyableRaw struct {
	Contract *Proxyable // Generic contract binding to access the raw methods on
}

// ProxyableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyableCallerRaw struct {
	Contract *ProxyableCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyableTransactorRaw struct {
	Contract *ProxyableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyable creates a new instance of Proxyable, bound to a specific deployed contract.
func NewProxyable(address common.Address, backend bind.ContractBackend) (*Proxyable, error) {
	contract, err := bindProxyable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxyable{ProxyableCaller: ProxyableCaller{contract: contract}, ProxyableTransactor: ProxyableTransactor{contract: contract}, ProxyableFilterer: ProxyableFilterer{contract: contract}}, nil
}

// NewProxyableCaller creates a new read-only instance of Proxyable, bound to a specific deployed contract.
func NewProxyableCaller(address common.Address, caller bind.ContractCaller) (*ProxyableCaller, error) {
	contract, err := bindProxyable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyableCaller{contract: contract}, nil
}

// NewProxyableTransactor creates a new write-only instance of Proxyable, bound to a specific deployed contract.
func NewProxyableTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyableTransactor, error) {
	contract, err := bindProxyable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyableTransactor{contract: contract}, nil
}

// NewProxyableFilterer creates a new log filterer instance of Proxyable, bound to a specific deployed contract.
func NewProxyableFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyableFilterer, error) {
	contract, err := bindProxyable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyableFilterer{contract: contract}, nil
}

// bindProxyable binds a generic wrapper to an already deployed contract.
func bindProxyable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxyable *ProxyableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxyable.Contract.ProxyableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxyable *ProxyableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxyable.Contract.ProxyableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxyable *ProxyableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxyable.Contract.ProxyableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxyable *ProxyableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxyable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxyable *ProxyableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxyable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxyable *ProxyableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxyable.Contract.contract.Transact(opts, method, params...)
}

// IntegrationProxy is a free data retrieval call binding the contract method 0x9cbdaeb6.
//
// Solidity: function integrationProxy() view returns(address)
func (_Proxyable *ProxyableCaller) IntegrationProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxyable.contract.Call(opts, &out, "integrationProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IntegrationProxy is a free data retrieval call binding the contract method 0x9cbdaeb6.
//
// Solidity: function integrationProxy() view returns(address)
func (_Proxyable *ProxyableSession) IntegrationProxy() (common.Address, error) {
	return _Proxyable.Contract.IntegrationProxy(&_Proxyable.CallOpts)
}

// IntegrationProxy is a free data retrieval call binding the contract method 0x9cbdaeb6.
//
// Solidity: function integrationProxy() view returns(address)
func (_Proxyable *ProxyableCallerSession) IntegrationProxy() (common.Address, error) {
	return _Proxyable.Contract.IntegrationProxy(&_Proxyable.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Proxyable *ProxyableCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxyable.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Proxyable *ProxyableSession) NominatedOwner() (common.Address, error) {
	return _Proxyable.Contract.NominatedOwner(&_Proxyable.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Proxyable *ProxyableCallerSession) NominatedOwner() (common.Address, error) {
	return _Proxyable.Contract.NominatedOwner(&_Proxyable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxyable *ProxyableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxyable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxyable *ProxyableSession) Owner() (common.Address, error) {
	return _Proxyable.Contract.Owner(&_Proxyable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxyable *ProxyableCallerSession) Owner() (common.Address, error) {
	return _Proxyable.Contract.Owner(&_Proxyable.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_Proxyable *ProxyableCaller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxyable.contract.Call(opts, &out, "proxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_Proxyable *ProxyableSession) Proxy() (common.Address, error) {
	return _Proxyable.Contract.Proxy(&_Proxyable.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_Proxyable *ProxyableCallerSession) Proxy() (common.Address, error) {
	return _Proxyable.Contract.Proxy(&_Proxyable.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Proxyable *ProxyableTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxyable.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Proxyable *ProxyableSession) AcceptOwnership() (*types.Transaction, error) {
	return _Proxyable.Contract.AcceptOwnership(&_Proxyable.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Proxyable *ProxyableTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Proxyable.Contract.AcceptOwnership(&_Proxyable.TransactOpts)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Proxyable *ProxyableTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Proxyable.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Proxyable *ProxyableSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.NominateNewOwner(&_Proxyable.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Proxyable *ProxyableTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.NominateNewOwner(&_Proxyable.TransactOpts, _owner)
}

// SetIntegrationProxy is a paid mutator transaction binding the contract method 0x131b0ae7.
//
// Solidity: function setIntegrationProxy(address _integrationProxy) returns()
func (_Proxyable *ProxyableTransactor) SetIntegrationProxy(opts *bind.TransactOpts, _integrationProxy common.Address) (*types.Transaction, error) {
	return _Proxyable.contract.Transact(opts, "setIntegrationProxy", _integrationProxy)
}

// SetIntegrationProxy is a paid mutator transaction binding the contract method 0x131b0ae7.
//
// Solidity: function setIntegrationProxy(address _integrationProxy) returns()
func (_Proxyable *ProxyableSession) SetIntegrationProxy(_integrationProxy common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.SetIntegrationProxy(&_Proxyable.TransactOpts, _integrationProxy)
}

// SetIntegrationProxy is a paid mutator transaction binding the contract method 0x131b0ae7.
//
// Solidity: function setIntegrationProxy(address _integrationProxy) returns()
func (_Proxyable *ProxyableTransactorSession) SetIntegrationProxy(_integrationProxy common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.SetIntegrationProxy(&_Proxyable.TransactOpts, _integrationProxy)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_Proxyable *ProxyableTransactor) SetMessageSender(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _Proxyable.contract.Transact(opts, "setMessageSender", sender)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_Proxyable *ProxyableSession) SetMessageSender(sender common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.SetMessageSender(&_Proxyable.TransactOpts, sender)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_Proxyable *ProxyableTransactorSession) SetMessageSender(sender common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.SetMessageSender(&_Proxyable.TransactOpts, sender)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_Proxyable *ProxyableTransactor) SetProxy(opts *bind.TransactOpts, _proxy common.Address) (*types.Transaction, error) {
	return _Proxyable.contract.Transact(opts, "setProxy", _proxy)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_Proxyable *ProxyableSession) SetProxy(_proxy common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.SetProxy(&_Proxyable.TransactOpts, _proxy)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_Proxyable *ProxyableTransactorSession) SetProxy(_proxy common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.SetProxy(&_Proxyable.TransactOpts, _proxy)
}

// ProxyableOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Proxyable contract.
type ProxyableOwnerChangedIterator struct {
	Event *ProxyableOwnerChanged // Event containing the contract specifics and raw log

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
func (it *ProxyableOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyableOwnerChanged)
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
		it.Event = new(ProxyableOwnerChanged)
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
func (it *ProxyableOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyableOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyableOwnerChanged represents a OwnerChanged event raised by the Proxyable contract.
type ProxyableOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Proxyable *ProxyableFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*ProxyableOwnerChangedIterator, error) {

	logs, sub, err := _Proxyable.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &ProxyableOwnerChangedIterator{contract: _Proxyable.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Proxyable *ProxyableFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *ProxyableOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Proxyable.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyableOwnerChanged)
				if err := _Proxyable.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Proxyable *ProxyableFilterer) ParseOwnerChanged(log types.Log) (*ProxyableOwnerChanged, error) {
	event := new(ProxyableOwnerChanged)
	if err := _Proxyable.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyableOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the Proxyable contract.
type ProxyableOwnerNominatedIterator struct {
	Event *ProxyableOwnerNominated // Event containing the contract specifics and raw log

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
func (it *ProxyableOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyableOwnerNominated)
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
		it.Event = new(ProxyableOwnerNominated)
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
func (it *ProxyableOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyableOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyableOwnerNominated represents a OwnerNominated event raised by the Proxyable contract.
type ProxyableOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Proxyable *ProxyableFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*ProxyableOwnerNominatedIterator, error) {

	logs, sub, err := _Proxyable.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &ProxyableOwnerNominatedIterator{contract: _Proxyable.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Proxyable *ProxyableFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *ProxyableOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _Proxyable.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyableOwnerNominated)
				if err := _Proxyable.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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

// ParseOwnerNominated is a log parse operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Proxyable *ProxyableFilterer) ParseOwnerNominated(log types.Log) (*ProxyableOwnerNominated, error) {
	event := new(ProxyableOwnerNominated)
	if err := _Proxyable.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyableProxyUpdatedIterator is returned from FilterProxyUpdated and is used to iterate over the raw logs and unpacked data for ProxyUpdated events raised by the Proxyable contract.
type ProxyableProxyUpdatedIterator struct {
	Event *ProxyableProxyUpdated // Event containing the contract specifics and raw log

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
func (it *ProxyableProxyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyableProxyUpdated)
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
		it.Event = new(ProxyableProxyUpdated)
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
func (it *ProxyableProxyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyableProxyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyableProxyUpdated represents a ProxyUpdated event raised by the Proxyable contract.
type ProxyableProxyUpdated struct {
	ProxyAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProxyUpdated is a free log retrieval operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_Proxyable *ProxyableFilterer) FilterProxyUpdated(opts *bind.FilterOpts) (*ProxyableProxyUpdatedIterator, error) {

	logs, sub, err := _Proxyable.contract.FilterLogs(opts, "ProxyUpdated")
	if err != nil {
		return nil, err
	}
	return &ProxyableProxyUpdatedIterator{contract: _Proxyable.contract, event: "ProxyUpdated", logs: logs, sub: sub}, nil
}

// WatchProxyUpdated is a free log subscription operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_Proxyable *ProxyableFilterer) WatchProxyUpdated(opts *bind.WatchOpts, sink chan<- *ProxyableProxyUpdated) (event.Subscription, error) {

	logs, sub, err := _Proxyable.contract.WatchLogs(opts, "ProxyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyableProxyUpdated)
				if err := _Proxyable.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
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

// ParseProxyUpdated is a log parse operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_Proxyable *ProxyableFilterer) ParseProxyUpdated(log types.Log) (*ProxyableProxyUpdated, error) {
	event := new(ProxyableProxyUpdated)
	if err := _Proxyable.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
