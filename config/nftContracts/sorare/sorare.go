// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sorare

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

// AccessControlABI is the input ABI used to generate the binding from.
const AccessControlABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AccessControlFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlFuncSigs = map[string]string{
	"a217fddf": "DEFAULT_ADMIN_ROLE()",
	"248a9ca3": "getRoleAdmin(bytes32)",
	"9010d07c": "getRoleMember(bytes32,uint256)",
	"ca15c873": "getRoleMemberCount(bytes32)",
	"2f2ff15d": "grantRole(bytes32,address)",
	"91d14854": "hasRole(bytes32,address)",
	"36568abe": "renounceRole(bytes32,address)",
	"d547741f": "revokeRole(bytes32,address)",
}

// AccessControl is an auto generated Go binding around an Ethereum contract.
type AccessControl struct {
	AccessControlCaller     // Read-only binding to the contract
	AccessControlTransactor // Write-only binding to the contract
	AccessControlFilterer   // Log filterer for contract events
}

// AccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlSession struct {
	Contract     *AccessControl    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlCallerSession struct {
	Contract *AccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlTransactorSession struct {
	Contract     *AccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlRaw struct {
	Contract *AccessControl // Generic contract binding to access the raw methods on
}

// AccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlCallerRaw struct {
	Contract *AccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlTransactorRaw struct {
	Contract *AccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControl creates a new instance of AccessControl, bound to a specific deployed contract.
func NewAccessControl(address common.Address, backend bind.ContractBackend) (*AccessControl, error) {
	contract, err := bindAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// NewAccessControlCaller creates a new read-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlCaller(address common.Address, caller bind.ContractCaller) (*AccessControlCaller, error) {
	contract, err := bindAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlCaller{contract: contract}, nil
}

// NewAccessControlTransactor creates a new write-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlTransactor, error) {
	contract, err := bindAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlTransactor{contract: contract}, nil
}

// NewAccessControlFilterer creates a new log filterer instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlFilterer, error) {
	contract, err := bindAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlFilterer{contract: contract}, nil
}

// bindAccessControl binds a generic wrapper to an already deployed contract.
func bindAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.AccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControl.Contract.GetRoleAdmin(&_AccessControl.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControl.Contract.GetRoleAdmin(&_AccessControl.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControl *AccessControlCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControl *AccessControlSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessControl.Contract.GetRoleMember(&_AccessControl.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControl *AccessControlCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessControl.Contract.GetRoleMember(&_AccessControl.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControl *AccessControlCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControl *AccessControlSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControl.Contract.GetRoleMemberCount(&_AccessControl.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControl *AccessControlCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControl.Contract.GetRoleMemberCount(&_AccessControl.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// AccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControl contract.
type AccessControlRoleGrantedIterator struct {
	Event *AccessControlRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleGranted)
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
		it.Event = new(AccessControlRoleGranted)
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
func (it *AccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleGranted represents a RoleGranted event raised by the AccessControl contract.
type AccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleGrantedIterator{contract: _AccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleGranted)
				if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) ParseRoleGranted(log types.Log) (*AccessControlRoleGranted, error) {
	event := new(AccessControlRoleGranted)
	if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControl contract.
type AccessControlRoleRevokedIterator struct {
	Event *AccessControlRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleRevoked)
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
		it.Event = new(AccessControlRoleRevoked)
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
func (it *AccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleRevoked represents a RoleRevoked event raised by the AccessControl contract.
type AccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleRevokedIterator{contract: _AccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleRevoked)
				if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) ParseRoleRevoked(log types.Log) (*AccessControlRoleRevoked, error) {
	event := new(AccessControlRoleRevoked)
	if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202b0eb88c4c4349fe220768698c4f9fc125177f16e50e87b66b8c7d7ec78951ed64736f6c63430006060033"

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ERC165ABI is the input ABI used to generate the binding from.
const ERC165ABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC165FuncSigs maps the 4-byte function signature to its string representation.
var ERC165FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// ERC721ABI is the input ABI used to generate the binding from.
const ERC721ABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721FuncSigs maps the 4-byte function signature to its string representation.
var ERC721FuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"6c0360eb": "baseURI()",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"06fdde03": "name()",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"4f6ccce7": "tokenByIndex(uint256)",
	"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
	"c87b56dd": "tokenURI(uint256)",
	"18160ddd": "totalSupply()",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC721Bin is the compiled bytecode used for deploying new contracts.
var ERC721Bin = "0x60806040523480156200001157600080fd5b5060405162001cd638038062001cd6833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b5060405250620001bc91506301ffc9a760e01b90506001600160e01b036200024116565b8151620001d1906006906020850190620002c6565b508051620001e7906007906020840190620002c6565b50620002036380ac58cd60e01b6001600160e01b036200024116565b6200021e635b5e139f60e01b6001600160e01b036200024116565b6200023963780e9d6360e01b6001600160e01b036200024116565b50506200036b565b6001600160e01b03198082161415620002a1576040805162461bcd60e51b815260206004820152601c60248201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604482015290519081900360640190fd5b6001600160e01b0319166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200030957805160ff191683800117855562000339565b8280016001018555821562000339579182015b82811115620003395782518255916020019190600101906200031c565b50620003479291506200034b565b5090565b6200036891905b8082111562000347576000815560010162000352565b90565b61195b806200037b6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80634f6ccce7116100a257806395d89b411161007157806395d89b4114610349578063a22cb46514610351578063b88d4fde1461037f578063c87b56dd14610445578063e985e9c5146104625761010b565b80634f6ccce7146102e15780636352211e146102fe5780636c0360eb1461031b57806370a08231146103235761010b565b806318160ddd116100de57806318160ddd1461022f57806323b872dd146102495780632f745c591461027f57806342842e0e146102ab5761010b565b806301ffc9a71461011057806306fdde031461014b578063081812fc146101c8578063095ea7b314610201575b600080fd5b6101376004803603602081101561012657600080fd5b50356001600160e01b031916610490565b604080519115158252519081900360200190f35b6101536104b3565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561018d578181015183820152602001610175565b50505050905090810190601f1680156101ba5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101e5600480360360208110156101de57600080fd5b5035610549565b604080516001600160a01b039092168252519081900360200190f35b61022d6004803603604081101561021757600080fd5b506001600160a01b0381351690602001356105ab565b005b610237610686565b60408051918252519081900360200190f35b61022d6004803603606081101561025f57600080fd5b506001600160a01b03813581169160208101359091169060400135610697565b6102376004803603604081101561029557600080fd5b506001600160a01b0381351690602001356106ee565b61022d600480360360608110156102c157600080fd5b506001600160a01b0381358116916020810135909116906040013561071f565b610237600480360360208110156102f757600080fd5b503561073a565b6101e56004803603602081101561031457600080fd5b5035610756565b610153610784565b6102376004803603602081101561033957600080fd5b50356001600160a01b03166107e5565b61015361084d565b61022d6004803603604081101561036757600080fd5b506001600160a01b03813516906020013515156108ae565b61022d6004803603608081101561039557600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156103d057600080fd5b8201836020820111156103e257600080fd5b8035906020019184600183028401116401000000008311171561040457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506109b3945050505050565b6101536004803603602081101561045b57600080fd5b5035610a11565b6101376004803603604081101561047857600080fd5b506001600160a01b0381358116916020013516610cb8565b6001600160e01b0319811660009081526020819052604090205460ff165b919050565b60068054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561053f5780601f106105145761010080835404028352916020019161053f565b820191906000526020600020905b81548152906001019060200180831161052257829003601f168201915b5050505050905090565b600061055482610ce6565b61058f5760405162461bcd60e51b815260040180806020018281038252602c815260200180611850602c913960400191505060405180910390fd5b506000908152600460205260409020546001600160a01b031690565b60006105b682610756565b9050806001600160a01b0316836001600160a01b031614156106095760405162461bcd60e51b81526004018080602001828103825260218152602001806118d46021913960400191505060405180910390fd5b806001600160a01b031661061b610cf9565b6001600160a01b0316148061063c575061063c81610637610cf9565b610cb8565b6106775760405162461bcd60e51b81526004018080602001828103825260388152602001806117a36038913960400191505060405180910390fd5b6106818383610cfd565b505050565b60006106926002610d6b565b905090565b6106a86106a2610cf9565b82610d76565b6106e35760405162461bcd60e51b81526004018080602001828103825260318152602001806118f56031913960400191505060405180910390fd5b610681838383610e1a565b6001600160a01b0382166000908152600160205260408120610716908363ffffffff610f7816565b90505b92915050565b610681838383604051806020016040528060008152506109b3565b60008061074e60028463ffffffff610f8416565b509392505050565b600061071982604051806060016040528060298152602001611805602991396002919063ffffffff610fa016565b60098054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561053f5780601f106105145761010080835404028352916020019161053f565b60006001600160a01b03821661082c5760405162461bcd60e51b815260040180806020018281038252602a8152602001806117db602a913960400191505060405180910390fd5b6001600160a01b038216600090815260016020526040902061071990610d6b565b60078054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561053f5780601f106105145761010080835404028352916020019161053f565b6108b6610cf9565b6001600160a01b0316826001600160a01b0316141561091c576040805162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015290519081900360640190fd5b8060056000610929610cf9565b6001600160a01b03908116825260208083019390935260409182016000908120918716808252919093529120805460ff19169215159290921790915561096d610cf9565b60408051841515815290516001600160a01b0392909216917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c319181900360200190a35050565b6109c46109be610cf9565b83610d76565b6109ff5760405162461bcd60e51b81526004018080602001828103825260318152602001806118f56031913960400191505060405180910390fd5b610a0b84848484610fb7565b50505050565b6060610a1c82610ce6565b610a575760405162461bcd60e51b815260040180806020018281038252602f8152602001806118a5602f913960400191505060405180910390fd5b60008281526008602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845260609392830182828015610aec5780601f10610ac157610100808354040283529160200191610aec565b820191906000526020600020905b815481529060010190602001808311610acf57829003601f168201915b505060095493945050505060026000196101006001841615020190911604610b155790506104ae565b805115610be6576009816040516020018083805460018160011615610100020316600290048015610b7d5780601f10610b5b576101008083540402835291820191610b7d565b820191906000526020600020905b815481529060010190602001808311610b69575b5050825160208401908083835b60208310610ba95780518252601f199092019160209182019101610b8a565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529150506104ae565b6009610bf184611009565b6040516020018083805460018160011615610100020316600290048015610c4f5780601f10610c2d576101008083540402835291820191610c4f565b820191906000526020600020905b815481529060010190602001808311610c3b575b5050825160208401908083835b60208310610c7b5780518252601f199092019160209182019101610c5c565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052915050919050565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b600061071960028363ffffffff6110e416565b3390565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610d3282610756565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000610719826110f0565b6000610d8182610ce6565b610dbc5760405162461bcd60e51b815260040180806020018281038252602c815260200180611777602c913960400191505060405180910390fd5b6000610dc783610756565b9050806001600160a01b0316846001600160a01b03161480610e025750836001600160a01b0316610df784610549565b6001600160a01b0316145b80610e125750610e128185610cb8565b949350505050565b826001600160a01b0316610e2d82610756565b6001600160a01b031614610e725760405162461bcd60e51b815260040180806020018281038252602981526020018061187c6029913960400191505060405180910390fd5b6001600160a01b038216610eb75760405162461bcd60e51b81526004018080602001828103825260248152602001806117536024913960400191505060405180910390fd5b610ec2838383610681565b610ecd600082610cfd565b6001600160a01b0383166000908152600160205260409020610ef5908263ffffffff6110f416565b506001600160a01b0382166000908152600160205260409020610f1e908263ffffffff61110016565b50610f316002828463ffffffff61110c16565b5080826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b60006107168383611122565b6000808080610f938686611186565b9097909650945050505050565b6000610fad848484611201565b90505b9392505050565b610fc2848484610e1a565b610fce848484846112cb565b610a0b5760405162461bcd60e51b81526004018080602001828103825260328152602001806117216032913960400191505060405180910390fd5b60608161102e57506040805180820190915260018152600360fc1b60208201526104ae565b8160005b811561104657600101600a82049150611032565b60608167ffffffffffffffff8111801561105f57600080fd5b506040519080825280601f01601f19166020018201604052801561108a576020820181803683370190505b50859350905060001982015b83156110db57600a840660300160f81b828280600190039350815181106110b957fe5b60200101906001600160f81b031916908160001a905350600a84049350611096565b50949350505050565b60006107168383611506565b5490565b6000610716838361151e565b600061071683836115e4565b6000610fad84846001600160a01b03851661162e565b815460009082106111645760405162461bcd60e51b81526004018080602001828103825260228152602001806116ff6022913960400191505060405180910390fd5b82600001828154811061117357fe5b9060005260206000200154905092915050565b8154600090819083106111ca5760405162461bcd60e51b815260040180806020018281038252602281526020018061182e6022913960400191505060405180910390fd5b60008460000184815481106111db57fe5b906000526020600020906002020190508060000154816001015492509250509250929050565b6000828152600184016020526040812054828161129c5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611261578181015183820152602001611249565b50505050905090810190601f16801561128e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b508460000160018203815481106112af57fe5b9060005260206000209060020201600101549150509392505050565b60006112df846001600160a01b03166116c5565b6112eb57506001610e12565b600060606001600160a01b038616630a85bd0160e11b611309610cf9565b89888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561138257818101518382015260200161136a565b50505050905090810190601f1680156113af5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b602083106114175780518252601f1990920191602091820191016113f8565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611479576040519150601f19603f3d011682016040523d82523d6000602084013e61147e565b606091505b5091509150816114cf578051156114985780518082602001fd5b60405162461bcd60e51b81526004018080602001828103825260328152602001806117216032913960400191505060405180910390fd5b60008180602001905160208110156114e657600080fd5b50516001600160e01b031916630a85bd0160e11b149350610e1292505050565b60009081526001919091016020526040902054151590565b600081815260018301602052604081205480156115da578354600019808301919081019060009087908390811061155157fe5b906000526020600020015490508087600001848154811061156e57fe5b60009182526020808320909101929092558281526001898101909252604090209084019055865487908061159e57fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050610719565b6000915050610719565b60006115f08383611506565b61162657508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610719565b506000610719565b600082815260018401602052604081205480611693575050604080518082018252838152602080820184815286546001818101895560008981528481209551600290930290950191825591519082015586548684528188019092529290912055610fb0565b828560000160018303815481106116a657fe5b9060005260206000209060020201600101819055506000915050610fb0565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590610e1257505015159291505056fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e64734552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724552433732313a207472616e7366657220746f20746865207a65726f20616464726573734552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e456e756d657261626c654d61703a20696e646578206f7574206f6620626f756e64734552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732314d657461646174613a2055524920717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65724552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564a2646970667358221220f6f422b16aa279e945d803b8a56bafd616bb3493ab0fb5195077df294a76a03c64736f6c63430006060033"

// DeployERC721 deploys a new Ethereum contract, binding an instance of ERC721 to it.
func DeployERC721(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *ERC721, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721Bin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// ERC721 is an auto generated Go binding around an Ethereum contract.
type ERC721 struct {
	ERC721Caller     // Read-only binding to the contract
	ERC721Transactor // Write-only binding to the contract
	ERC721Filterer   // Log filterer for contract events
}

// ERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721Session struct {
	Contract     *ERC721           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721CallerSession struct {
	Contract *ERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TransactorSession struct {
	Contract     *ERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721Raw struct {
	Contract *ERC721 // Generic contract binding to access the raw methods on
}

// ERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721CallerRaw struct {
	Contract *ERC721Caller // Generic read-only contract binding to access the raw methods on
}

// ERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TransactorRaw struct {
	Contract *ERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721 creates a new instance of ERC721, bound to a specific deployed contract.
func NewERC721(address common.Address, backend bind.ContractBackend) (*ERC721, error) {
	contract, err := bindERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// NewERC721Caller creates a new read-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Caller(address common.Address, caller bind.ContractCaller) (*ERC721Caller, error) {
	contract, err := bindERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Caller{contract: contract}, nil
}

// NewERC721Transactor creates a new write-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC721Transactor, error) {
	contract, err := bindERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Transactor{contract: contract}, nil
}

// NewERC721Filterer creates a new log filterer instance of ERC721, bound to a specific deployed contract.
func NewERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC721Filterer, error) {
	contract, err := bindERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721Filterer{contract: contract}, nil
}

// bindERC721 binds a generic wrapper to an already deployed contract.
func bindERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.ERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_ERC721 *ERC721Caller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_ERC721 *ERC721Session) BaseURI() (string, error) {
	return _ERC721.Contract.BaseURI(&_ERC721.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_ERC721 *ERC721CallerSession) BaseURI() (string, error) {
	return _ERC721.Contract.BaseURI(&_ERC721.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721Session) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721CallerSession) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721Session) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721CallerSession) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721 *ERC721Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721 *ERC721Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenByIndex(&_ERC721.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721 *ERC721CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenByIndex(&_ERC721.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721 *ERC721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721 *ERC721Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenOfOwnerByIndex(&_ERC721.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721 *ERC721CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenOfOwnerByIndex(&_ERC721.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721 *ERC721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721 *ERC721Session) TotalSupply() (*big.Int, error) {
	return _ERC721.Contract.TotalSupply(&_ERC721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721 *ERC721CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721.Contract.TotalSupply(&_ERC721.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// ERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721 contract.
type ERC721ApprovalIterator struct {
	Event *ERC721Approval // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Approval)
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
		it.Event = new(ERC721Approval)
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
func (it *ERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Approval represents a Approval event raised by the ERC721 contract.
type ERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalIterator{contract: _ERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Approval)
				if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) ParseApproval(log types.Log) (*ERC721Approval, error) {
	event := new(ERC721Approval)
	if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721 contract.
type ERC721ApprovalForAllIterator struct {
	Event *ERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ApprovalForAll)
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
		it.Event = new(ERC721ApprovalForAll)
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
func (it *ERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ApprovalForAll represents a ApprovalForAll event raised by the ERC721 contract.
type ERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalForAllIterator{contract: _ERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ApprovalForAll)
				if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) ParseApprovalForAll(log types.Log) (*ERC721ApprovalForAll, error) {
	event := new(ERC721ApprovalForAll)
	if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721 contract.
type ERC721TransferIterator struct {
	Event *ERC721Transfer // Event containing the contract specifics and raw log

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
func (it *ERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Transfer)
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
		it.Event = new(ERC721Transfer)
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
func (it *ERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Transfer represents a Transfer event raised by the ERC721 contract.
type ERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TransferIterator{contract: _ERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Transfer)
				if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) ParseTransfer(log types.Log) (*ERC721Transfer, error) {
	event := new(ERC721Transfer)
	if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnumerableMapABI is the input ABI used to generate the binding from.
const EnumerableMapABI = "[]"

// EnumerableMapBin is the compiled bytecode used for deploying new contracts.
var EnumerableMapBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ae2d82cdf7f3594805398713255350301a08dbee7b0181a00aa4ca544b0ef88b64736f6c63430006060033"

// DeployEnumerableMap deploys a new Ethereum contract, binding an instance of EnumerableMap to it.
func DeployEnumerableMap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableMap, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableMapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EnumerableMapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableMap{EnumerableMapCaller: EnumerableMapCaller{contract: contract}, EnumerableMapTransactor: EnumerableMapTransactor{contract: contract}, EnumerableMapFilterer: EnumerableMapFilterer{contract: contract}}, nil
}

// EnumerableMap is an auto generated Go binding around an Ethereum contract.
type EnumerableMap struct {
	EnumerableMapCaller     // Read-only binding to the contract
	EnumerableMapTransactor // Write-only binding to the contract
	EnumerableMapFilterer   // Log filterer for contract events
}

// EnumerableMapCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableMapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableMapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableMapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableMapSession struct {
	Contract     *EnumerableMap    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableMapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableMapCallerSession struct {
	Contract *EnumerableMapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableMapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableMapTransactorSession struct {
	Contract     *EnumerableMapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableMapRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableMapRaw struct {
	Contract *EnumerableMap // Generic contract binding to access the raw methods on
}

// EnumerableMapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableMapCallerRaw struct {
	Contract *EnumerableMapCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableMapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableMapTransactorRaw struct {
	Contract *EnumerableMapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableMap creates a new instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMap(address common.Address, backend bind.ContractBackend) (*EnumerableMap, error) {
	contract, err := bindEnumerableMap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableMap{EnumerableMapCaller: EnumerableMapCaller{contract: contract}, EnumerableMapTransactor: EnumerableMapTransactor{contract: contract}, EnumerableMapFilterer: EnumerableMapFilterer{contract: contract}}, nil
}

// NewEnumerableMapCaller creates a new read-only instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapCaller(address common.Address, caller bind.ContractCaller) (*EnumerableMapCaller, error) {
	contract, err := bindEnumerableMap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapCaller{contract: contract}, nil
}

// NewEnumerableMapTransactor creates a new write-only instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableMapTransactor, error) {
	contract, err := bindEnumerableMap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapTransactor{contract: contract}, nil
}

// NewEnumerableMapFilterer creates a new log filterer instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableMapFilterer, error) {
	contract, err := bindEnumerableMap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapFilterer{contract: contract}, nil
}

// bindEnumerableMap binds a generic wrapper to an already deployed contract.
func bindEnumerableMap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableMapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableMap *EnumerableMapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableMap.Contract.EnumerableMapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableMap *EnumerableMapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableMap.Contract.EnumerableMapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableMap *EnumerableMapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableMap.Contract.EnumerableMapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableMap *EnumerableMapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableMap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableMap *EnumerableMapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableMap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableMap *EnumerableMapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableMap.Contract.contract.Transact(opts, method, params...)
}

// EnumerableSetABI is the input ABI used to generate the binding from.
const EnumerableSetABI = "[]"

// EnumerableSetBin is the compiled bytecode used for deploying new contracts.
var EnumerableSetBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c88debb9494829d50bc3d47c25913a9854deb61fadd18a19b59c7669365b211664736f6c63430006060033"

// DeployEnumerableSet deploys a new Ethereum contract, binding an instance of EnumerableSet to it.
func DeployEnumerableSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableSet, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableSetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EnumerableSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// EnumerableSet is an auto generated Go binding around an Ethereum contract.
type EnumerableSet struct {
	EnumerableSetCaller     // Read-only binding to the contract
	EnumerableSetTransactor // Write-only binding to the contract
	EnumerableSetFilterer   // Log filterer for contract events
}

// EnumerableSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableSetSession struct {
	Contract     *EnumerableSet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableSetCallerSession struct {
	Contract *EnumerableSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableSetTransactorSession struct {
	Contract     *EnumerableSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableSetRaw struct {
	Contract *EnumerableSet // Generic contract binding to access the raw methods on
}

// EnumerableSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableSetCallerRaw struct {
	Contract *EnumerableSetCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableSetTransactorRaw struct {
	Contract *EnumerableSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableSet creates a new instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSet(address common.Address, backend bind.ContractBackend) (*EnumerableSet, error) {
	contract, err := bindEnumerableSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// NewEnumerableSetCaller creates a new read-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetCaller(address common.Address, caller bind.ContractCaller) (*EnumerableSetCaller, error) {
	contract, err := bindEnumerableSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetCaller{contract: contract}, nil
}

// NewEnumerableSetTransactor creates a new write-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableSetTransactor, error) {
	contract, err := bindEnumerableSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetTransactor{contract: contract}, nil
}

// NewEnumerableSetFilterer creates a new log filterer instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableSetFilterer, error) {
	contract, err := bindEnumerableSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetFilterer{contract: contract}, nil
}

// bindEnumerableSet binds a generic wrapper to an already deployed contract.
func bindEnumerableSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.EnumerableSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transact(opts, method, params...)
}

// IERC165ABI is the input ABI used to generate the binding from.
const IERC165ABI = "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IERC721ABI is the input ABI used to generate the binding from.
const IERC721ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC721FuncSigs maps the 4-byte function signature to its string representation.
var IERC721FuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC721 is an auto generated Go binding around an Ethereum contract.
type IERC721 struct {
	IERC721Caller     // Read-only binding to the contract
	IERC721Transactor // Write-only binding to the contract
	IERC721Filterer   // Log filterer for contract events
}

// IERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721Session struct {
	Contract     *IERC721          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721CallerSession struct {
	Contract *IERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721TransactorSession struct {
	Contract     *IERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721Raw struct {
	Contract *IERC721 // Generic contract binding to access the raw methods on
}

// IERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721CallerRaw struct {
	Contract *IERC721Caller // Generic read-only contract binding to access the raw methods on
}

// IERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721TransactorRaw struct {
	Contract *IERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721 creates a new instance of IERC721, bound to a specific deployed contract.
func NewIERC721(address common.Address, backend bind.ContractBackend) (*IERC721, error) {
	contract, err := bindIERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721{IERC721Caller: IERC721Caller{contract: contract}, IERC721Transactor: IERC721Transactor{contract: contract}, IERC721Filterer: IERC721Filterer{contract: contract}}, nil
}

// NewIERC721Caller creates a new read-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Caller(address common.Address, caller bind.ContractCaller) (*IERC721Caller, error) {
	contract, err := bindIERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Caller{contract: contract}, nil
}

// NewIERC721Transactor creates a new write-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC721Transactor, error) {
	contract, err := bindIERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Transactor{contract: contract}, nil
}

// NewIERC721Filterer creates a new log filterer instance of IERC721, bound to a specific deployed contract.
func NewIERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC721Filterer, error) {
	contract, err := bindIERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721Filterer{contract: contract}, nil
}

// bindIERC721 binds a generic wrapper to an already deployed contract.
func bindIERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.IERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Session) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721TransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// IERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721 contract.
type IERC721ApprovalIterator struct {
	Event *IERC721Approval // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Approval)
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
		it.Event = new(IERC721Approval)
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
func (it *IERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Approval represents a Approval event raised by the IERC721 contract.
type IERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalIterator{contract: _IERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Approval)
				if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) ParseApproval(log types.Log) (*IERC721Approval, error) {
	event := new(IERC721Approval)
	if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721 contract.
type IERC721ApprovalForAllIterator struct {
	Event *IERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721ApprovalForAll)
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
		it.Event = new(IERC721ApprovalForAll)
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
func (it *IERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721ApprovalForAll represents a ApprovalForAll event raised by the IERC721 contract.
type IERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalForAllIterator{contract: _IERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721ApprovalForAll)
				if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) ParseApprovalForAll(log types.Log) (*IERC721ApprovalForAll, error) {
	event := new(IERC721ApprovalForAll)
	if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721 contract.
type IERC721TransferIterator struct {
	Event *IERC721Transfer // Event containing the contract specifics and raw log

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
func (it *IERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Transfer)
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
		it.Event = new(IERC721Transfer)
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
func (it *IERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Transfer represents a Transfer event raised by the IERC721 contract.
type IERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721TransferIterator{contract: _IERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Transfer)
				if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) ParseTransfer(log types.Log) (*IERC721Transfer, error) {
	event := new(IERC721Transfer)
	if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721EnumerableABI is the input ABI used to generate the binding from.
const IERC721EnumerableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC721EnumerableFuncSigs maps the 4-byte function signature to its string representation.
var IERC721EnumerableFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"4f6ccce7": "tokenByIndex(uint256)",
	"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
	"18160ddd": "totalSupply()",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC721Enumerable is an auto generated Go binding around an Ethereum contract.
type IERC721Enumerable struct {
	IERC721EnumerableCaller     // Read-only binding to the contract
	IERC721EnumerableTransactor // Write-only binding to the contract
	IERC721EnumerableFilterer   // Log filterer for contract events
}

// IERC721EnumerableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721EnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721EnumerableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721EnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721EnumerableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721EnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721EnumerableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721EnumerableSession struct {
	Contract     *IERC721Enumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC721EnumerableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721EnumerableCallerSession struct {
	Contract *IERC721EnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IERC721EnumerableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721EnumerableTransactorSession struct {
	Contract     *IERC721EnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IERC721EnumerableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721EnumerableRaw struct {
	Contract *IERC721Enumerable // Generic contract binding to access the raw methods on
}

// IERC721EnumerableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721EnumerableCallerRaw struct {
	Contract *IERC721EnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721EnumerableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721EnumerableTransactorRaw struct {
	Contract *IERC721EnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Enumerable creates a new instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721Enumerable(address common.Address, backend bind.ContractBackend) (*IERC721Enumerable, error) {
	contract, err := bindIERC721Enumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Enumerable{IERC721EnumerableCaller: IERC721EnumerableCaller{contract: contract}, IERC721EnumerableTransactor: IERC721EnumerableTransactor{contract: contract}, IERC721EnumerableFilterer: IERC721EnumerableFilterer{contract: contract}}, nil
}

// NewIERC721EnumerableCaller creates a new read-only instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721EnumerableCaller(address common.Address, caller bind.ContractCaller) (*IERC721EnumerableCaller, error) {
	contract, err := bindIERC721Enumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableCaller{contract: contract}, nil
}

// NewIERC721EnumerableTransactor creates a new write-only instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721EnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721EnumerableTransactor, error) {
	contract, err := bindIERC721Enumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableTransactor{contract: contract}, nil
}

// NewIERC721EnumerableFilterer creates a new log filterer instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721EnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721EnumerableFilterer, error) {
	contract, err := bindIERC721Enumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableFilterer{contract: contract}, nil
}

// bindIERC721Enumerable binds a generic wrapper to an already deployed contract.
func bindIERC721Enumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721EnumerableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Enumerable *IERC721EnumerableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Enumerable.Contract.IERC721EnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Enumerable *IERC721EnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.IERC721EnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Enumerable *IERC721EnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.IERC721EnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Enumerable *IERC721EnumerableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Enumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Enumerable *IERC721EnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Enumerable *IERC721EnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Enumerable *IERC721EnumerableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Enumerable *IERC721EnumerableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Enumerable.Contract.BalanceOf(&_IERC721Enumerable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Enumerable.Contract.BalanceOf(&_IERC721Enumerable.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Enumerable *IERC721EnumerableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Enumerable *IERC721EnumerableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.GetApproved(&_IERC721Enumerable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.GetApproved(&_IERC721Enumerable.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Enumerable.Contract.IsApprovedForAll(&_IERC721Enumerable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Enumerable.Contract.IsApprovedForAll(&_IERC721Enumerable.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Enumerable *IERC721EnumerableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Enumerable *IERC721EnumerableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.OwnerOf(&_IERC721Enumerable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.OwnerOf(&_IERC721Enumerable.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Enumerable.Contract.SupportsInterface(&_IERC721Enumerable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Enumerable.Contract.SupportsInterface(&_IERC721Enumerable.CallOpts, interfaceId)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenByIndex(&_IERC721Enumerable.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenByIndex(&_IERC721Enumerable.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_IERC721Enumerable *IERC721EnumerableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_IERC721Enumerable *IERC721EnumerableSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenOfOwnerByIndex(&_IERC721Enumerable.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenOfOwnerByIndex(&_IERC721Enumerable.CallOpts, owner, index)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableSession) TotalSupply() (*big.Int, error) {
	return _IERC721Enumerable.Contract.TotalSupply(&_IERC721Enumerable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC721Enumerable.Contract.TotalSupply(&_IERC721Enumerable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.Approve(&_IERC721Enumerable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.Approve(&_IERC721Enumerable.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom0(&_IERC721Enumerable.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom0(&_IERC721Enumerable.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SetApprovalForAll(&_IERC721Enumerable.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SetApprovalForAll(&_IERC721Enumerable.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.TransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.TransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// IERC721EnumerableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721Enumerable contract.
type IERC721EnumerableApprovalIterator struct {
	Event *IERC721EnumerableApproval // Event containing the contract specifics and raw log

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
func (it *IERC721EnumerableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721EnumerableApproval)
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
		it.Event = new(IERC721EnumerableApproval)
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
func (it *IERC721EnumerableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721EnumerableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721EnumerableApproval represents a Approval event raised by the IERC721Enumerable contract.
type IERC721EnumerableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721EnumerableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableApprovalIterator{contract: _IERC721Enumerable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721EnumerableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721EnumerableApproval)
				if err := _IERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) ParseApproval(log types.Log) (*IERC721EnumerableApproval, error) {
	event := new(IERC721EnumerableApproval)
	if err := _IERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721EnumerableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721Enumerable contract.
type IERC721EnumerableApprovalForAllIterator struct {
	Event *IERC721EnumerableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721EnumerableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721EnumerableApprovalForAll)
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
		it.Event = new(IERC721EnumerableApprovalForAll)
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
func (it *IERC721EnumerableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721EnumerableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721EnumerableApprovalForAll represents a ApprovalForAll event raised by the IERC721Enumerable contract.
type IERC721EnumerableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Enumerable *IERC721EnumerableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721EnumerableApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableApprovalForAllIterator{contract: _IERC721Enumerable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Enumerable *IERC721EnumerableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721EnumerableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721EnumerableApprovalForAll)
				if err := _IERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Enumerable *IERC721EnumerableFilterer) ParseApprovalForAll(log types.Log) (*IERC721EnumerableApprovalForAll, error) {
	event := new(IERC721EnumerableApprovalForAll)
	if err := _IERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721EnumerableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721Enumerable contract.
type IERC721EnumerableTransferIterator struct {
	Event *IERC721EnumerableTransfer // Event containing the contract specifics and raw log

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
func (it *IERC721EnumerableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721EnumerableTransfer)
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
		it.Event = new(IERC721EnumerableTransfer)
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
func (it *IERC721EnumerableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721EnumerableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721EnumerableTransfer represents a Transfer event raised by the IERC721Enumerable contract.
type IERC721EnumerableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721EnumerableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableTransferIterator{contract: _IERC721Enumerable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721EnumerableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721EnumerableTransfer)
				if err := _IERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) ParseTransfer(log types.Log) (*IERC721EnumerableTransfer, error) {
	event := new(IERC721EnumerableTransfer)
	if err := _IERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721MetadataABI is the input ABI used to generate the binding from.
const IERC721MetadataABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC721MetadataFuncSigs maps the 4-byte function signature to its string representation.
var IERC721MetadataFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"06fdde03": "name()",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"c87b56dd": "tokenURI(uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC721Metadata is an auto generated Go binding around an Ethereum contract.
type IERC721Metadata struct {
	IERC721MetadataCaller     // Read-only binding to the contract
	IERC721MetadataTransactor // Write-only binding to the contract
	IERC721MetadataFilterer   // Log filterer for contract events
}

// IERC721MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721MetadataSession struct {
	Contract     *IERC721Metadata  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721MetadataCallerSession struct {
	Contract *IERC721MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC721MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721MetadataTransactorSession struct {
	Contract     *IERC721MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC721MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721MetadataRaw struct {
	Contract *IERC721Metadata // Generic contract binding to access the raw methods on
}

// IERC721MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721MetadataCallerRaw struct {
	Contract *IERC721MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721MetadataTransactorRaw struct {
	Contract *IERC721MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Metadata creates a new instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721Metadata(address common.Address, backend bind.ContractBackend) (*IERC721Metadata, error) {
	contract, err := bindIERC721Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Metadata{IERC721MetadataCaller: IERC721MetadataCaller{contract: contract}, IERC721MetadataTransactor: IERC721MetadataTransactor{contract: contract}, IERC721MetadataFilterer: IERC721MetadataFilterer{contract: contract}}, nil
}

// NewIERC721MetadataCaller creates a new read-only instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721MetadataCaller(address common.Address, caller bind.ContractCaller) (*IERC721MetadataCaller, error) {
	contract, err := bindIERC721Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataCaller{contract: contract}, nil
}

// NewIERC721MetadataTransactor creates a new write-only instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721MetadataTransactor, error) {
	contract, err := bindIERC721Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataTransactor{contract: contract}, nil
}

// NewIERC721MetadataFilterer creates a new log filterer instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721MetadataFilterer, error) {
	contract, err := bindIERC721Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataFilterer{contract: contract}, nil
}

// bindIERC721Metadata binds a generic wrapper to an already deployed contract.
func bindIERC721Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Metadata *IERC721MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Metadata.Contract.IERC721MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Metadata *IERC721MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.IERC721MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Metadata *IERC721MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.IERC721MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Metadata *IERC721MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Metadata *IERC721MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Metadata *IERC721MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Metadata *IERC721MetadataCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Metadata *IERC721MetadataSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Metadata.Contract.BalanceOf(&_IERC721Metadata.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Metadata *IERC721MetadataCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Metadata.Contract.BalanceOf(&_IERC721Metadata.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Metadata *IERC721MetadataCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Metadata *IERC721MetadataSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.GetApproved(&_IERC721Metadata.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Metadata *IERC721MetadataCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.GetApproved(&_IERC721Metadata.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Metadata *IERC721MetadataSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Metadata.Contract.IsApprovedForAll(&_IERC721Metadata.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Metadata.Contract.IsApprovedForAll(&_IERC721Metadata.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC721Metadata *IERC721MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC721Metadata *IERC721MetadataSession) Name() (string, error) {
	return _IERC721Metadata.Contract.Name(&_IERC721Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC721Metadata *IERC721MetadataCallerSession) Name() (string, error) {
	return _IERC721Metadata.Contract.Name(&_IERC721Metadata.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Metadata *IERC721MetadataCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Metadata *IERC721MetadataSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.OwnerOf(&_IERC721Metadata.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Metadata *IERC721MetadataCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.OwnerOf(&_IERC721Metadata.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Metadata *IERC721MetadataSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Metadata.Contract.SupportsInterface(&_IERC721Metadata.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Metadata.Contract.SupportsInterface(&_IERC721Metadata.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC721Metadata *IERC721MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC721Metadata *IERC721MetadataSession) Symbol() (string, error) {
	return _IERC721Metadata.Contract.Symbol(&_IERC721Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC721Metadata *IERC721MetadataCallerSession) Symbol() (string, error) {
	return _IERC721Metadata.Contract.Symbol(&_IERC721Metadata.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IERC721Metadata *IERC721MetadataCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IERC721Metadata *IERC721MetadataSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IERC721Metadata.Contract.TokenURI(&_IERC721Metadata.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IERC721Metadata *IERC721MetadataCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IERC721Metadata.Contract.TokenURI(&_IERC721Metadata.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.Approve(&_IERC721Metadata.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.Approve(&_IERC721Metadata.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Metadata *IERC721MetadataSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom0(&_IERC721Metadata.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom0(&_IERC721Metadata.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Metadata *IERC721MetadataSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SetApprovalForAll(&_IERC721Metadata.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SetApprovalForAll(&_IERC721Metadata.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.TransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.TransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// IERC721MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721Metadata contract.
type IERC721MetadataApprovalIterator struct {
	Event *IERC721MetadataApproval // Event containing the contract specifics and raw log

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
func (it *IERC721MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721MetadataApproval)
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
		it.Event = new(IERC721MetadataApproval)
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
func (it *IERC721MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721MetadataApproval represents a Approval event raised by the IERC721Metadata contract.
type IERC721MetadataApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721MetadataApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataApprovalIterator{contract: _IERC721Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721MetadataApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721MetadataApproval)
				if err := _IERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) ParseApproval(log types.Log) (*IERC721MetadataApproval, error) {
	event := new(IERC721MetadataApproval)
	if err := _IERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721MetadataApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721Metadata contract.
type IERC721MetadataApprovalForAllIterator struct {
	Event *IERC721MetadataApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721MetadataApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721MetadataApprovalForAll)
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
		it.Event = new(IERC721MetadataApprovalForAll)
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
func (it *IERC721MetadataApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721MetadataApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721MetadataApprovalForAll represents a ApprovalForAll event raised by the IERC721Metadata contract.
type IERC721MetadataApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Metadata *IERC721MetadataFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721MetadataApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataApprovalForAllIterator{contract: _IERC721Metadata.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Metadata *IERC721MetadataFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721MetadataApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721MetadataApprovalForAll)
				if err := _IERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Metadata *IERC721MetadataFilterer) ParseApprovalForAll(log types.Log) (*IERC721MetadataApprovalForAll, error) {
	event := new(IERC721MetadataApprovalForAll)
	if err := _IERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721Metadata contract.
type IERC721MetadataTransferIterator struct {
	Event *IERC721MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *IERC721MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721MetadataTransfer)
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
		it.Event = new(IERC721MetadataTransfer)
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
func (it *IERC721MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721MetadataTransfer represents a Transfer event raised by the IERC721Metadata contract.
type IERC721MetadataTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721MetadataTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataTransferIterator{contract: _IERC721Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721MetadataTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721MetadataTransfer)
				if err := _IERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) ParseTransfer(log types.Log) (*IERC721MetadataTransfer, error) {
	event := new(IERC721MetadataTransfer)
	if err := _IERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721ReceiverABI is the input ABI used to generate the binding from.
const IERC721ReceiverABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC721ReceiverFuncSigs maps the 4-byte function signature to its string representation.
var IERC721ReceiverFuncSigs = map[string]string{
	"150b7a02": "onERC721Received(address,address,uint256,bytes)",
}

// IERC721Receiver is an auto generated Go binding around an Ethereum contract.
type IERC721Receiver struct {
	IERC721ReceiverCaller     // Read-only binding to the contract
	IERC721ReceiverTransactor // Write-only binding to the contract
	IERC721ReceiverFilterer   // Log filterer for contract events
}

// IERC721ReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721ReceiverSession struct {
	Contract     *IERC721Receiver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721ReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721ReceiverCallerSession struct {
	Contract *IERC721ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC721ReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721ReceiverTransactorSession struct {
	Contract     *IERC721ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC721ReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721ReceiverRaw struct {
	Contract *IERC721Receiver // Generic contract binding to access the raw methods on
}

// IERC721ReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721ReceiverCallerRaw struct {
	Contract *IERC721ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721ReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721ReceiverTransactorRaw struct {
	Contract *IERC721ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Receiver creates a new instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721Receiver(address common.Address, backend bind.ContractBackend) (*IERC721Receiver, error) {
	contract, err := bindIERC721Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Receiver{IERC721ReceiverCaller: IERC721ReceiverCaller{contract: contract}, IERC721ReceiverTransactor: IERC721ReceiverTransactor{contract: contract}, IERC721ReceiverFilterer: IERC721ReceiverFilterer{contract: contract}}, nil
}

// NewIERC721ReceiverCaller creates a new read-only instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverCaller(address common.Address, caller bind.ContractCaller) (*IERC721ReceiverCaller, error) {
	contract, err := bindIERC721Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverCaller{contract: contract}, nil
}

// NewIERC721ReceiverTransactor creates a new write-only instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721ReceiverTransactor, error) {
	contract, err := bindIERC721Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverTransactor{contract: contract}, nil
}

// NewIERC721ReceiverFilterer creates a new log filterer instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721ReceiverFilterer, error) {
	contract, err := bindIERC721Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverFilterer{contract: contract}, nil
}

// bindIERC721Receiver binds a generic wrapper to an already deployed contract.
func bindIERC721Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Receiver *IERC721ReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Receiver.Contract.IERC721ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Receiver *IERC721ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Receiver *IERC721ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Receiver *IERC721ReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.contract.Transact(opts, method, params...)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
}

// INextContractABI is the input ABI used to generate the binding from.
const INextContractABI = "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"migrateTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// INextContractFuncSigs maps the 4-byte function signature to its string representation.
var INextContractFuncSigs = map[string]string{
	"e884703f": "migrateTokens(uint256[],address)",
}

// INextContract is an auto generated Go binding around an Ethereum contract.
type INextContract struct {
	INextContractCaller     // Read-only binding to the contract
	INextContractTransactor // Write-only binding to the contract
	INextContractFilterer   // Log filterer for contract events
}

// INextContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type INextContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INextContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INextContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INextContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INextContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INextContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INextContractSession struct {
	Contract     *INextContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INextContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INextContractCallerSession struct {
	Contract *INextContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// INextContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INextContractTransactorSession struct {
	Contract     *INextContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// INextContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type INextContractRaw struct {
	Contract *INextContract // Generic contract binding to access the raw methods on
}

// INextContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INextContractCallerRaw struct {
	Contract *INextContractCaller // Generic read-only contract binding to access the raw methods on
}

// INextContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INextContractTransactorRaw struct {
	Contract *INextContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINextContract creates a new instance of INextContract, bound to a specific deployed contract.
func NewINextContract(address common.Address, backend bind.ContractBackend) (*INextContract, error) {
	contract, err := bindINextContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INextContract{INextContractCaller: INextContractCaller{contract: contract}, INextContractTransactor: INextContractTransactor{contract: contract}, INextContractFilterer: INextContractFilterer{contract: contract}}, nil
}

// NewINextContractCaller creates a new read-only instance of INextContract, bound to a specific deployed contract.
func NewINextContractCaller(address common.Address, caller bind.ContractCaller) (*INextContractCaller, error) {
	contract, err := bindINextContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INextContractCaller{contract: contract}, nil
}

// NewINextContractTransactor creates a new write-only instance of INextContract, bound to a specific deployed contract.
func NewINextContractTransactor(address common.Address, transactor bind.ContractTransactor) (*INextContractTransactor, error) {
	contract, err := bindINextContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INextContractTransactor{contract: contract}, nil
}

// NewINextContractFilterer creates a new log filterer instance of INextContract, bound to a specific deployed contract.
func NewINextContractFilterer(address common.Address, filterer bind.ContractFilterer) (*INextContractFilterer, error) {
	contract, err := bindINextContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INextContractFilterer{contract: contract}, nil
}

// bindINextContract binds a generic wrapper to an already deployed contract.
func bindINextContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(INextContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INextContract *INextContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INextContract.Contract.INextContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INextContract *INextContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INextContract.Contract.INextContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INextContract *INextContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INextContract.Contract.INextContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INextContract *INextContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INextContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INextContract *INextContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INextContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INextContract *INextContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INextContract.Contract.contract.Transact(opts, method, params...)
}

// MigrateTokens is a paid mutator transaction binding the contract method 0xe884703f.
//
// Solidity: function migrateTokens(uint256[] tokenIds, address to) returns()
func (_INextContract *INextContractTransactor) MigrateTokens(opts *bind.TransactOpts, tokenIds []*big.Int, to common.Address) (*types.Transaction, error) {
	return _INextContract.contract.Transact(opts, "migrateTokens", tokenIds, to)
}

// MigrateTokens is a paid mutator transaction binding the contract method 0xe884703f.
//
// Solidity: function migrateTokens(uint256[] tokenIds, address to) returns()
func (_INextContract *INextContractSession) MigrateTokens(tokenIds []*big.Int, to common.Address) (*types.Transaction, error) {
	return _INextContract.Contract.MigrateTokens(&_INextContract.TransactOpts, tokenIds, to)
}

// MigrateTokens is a paid mutator transaction binding the contract method 0xe884703f.
//
// Solidity: function migrateTokens(uint256[] tokenIds, address to) returns()
func (_INextContract *INextContractTransactorSession) MigrateTokens(tokenIds []*big.Int, to common.Address) (*types.Transaction, error) {
	return _INextContract.Contract.MigrateTokens(&_INextContract.TransactOpts, tokenIds, to)
}

// ISorareCardsABI is the input ABI used to generate the binding from.
const ISorareCardsABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cardId\",\"type\":\"uint256\"}],\"name\":\"cardExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"playerId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"season\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"scarcity\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"serialNumber\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"clubId\",\"type\":\"uint16\"}],\"name\":\"createCard\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cardId\",\"type\":\"uint256\"}],\"name\":\"getCard\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"playerId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"season\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"scarcity\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"serialNumber\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"clubId\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"clubId\",\"type\":\"uint16\"}],\"name\":\"getClub\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"country\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"city\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"yearFounded\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"playerId\",\"type\":\"uint256\"}],\"name\":\"getPlayer\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"yearOfBirth\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"monthOfBirth\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"dayOfBirth\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ISorareCardsFuncSigs maps the 4-byte function signature to its string representation.
var ISorareCardsFuncSigs = map[string]string{
	"4ae3540e": "cardExists(uint256)",
	"01152383": "createCard(uint256,uint16,uint8,uint16,bytes32,uint16)",
	"9188d312": "getCard(uint256)",
	"4a39d2b8": "getClub(uint16)",
	"e55ae4e8": "getPlayer(uint256)",
}

// ISorareCards is an auto generated Go binding around an Ethereum contract.
type ISorareCards struct {
	ISorareCardsCaller     // Read-only binding to the contract
	ISorareCardsTransactor // Write-only binding to the contract
	ISorareCardsFilterer   // Log filterer for contract events
}

// ISorareCardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISorareCardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISorareCardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISorareCardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISorareCardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISorareCardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISorareCardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISorareCardsSession struct {
	Contract     *ISorareCards     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISorareCardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISorareCardsCallerSession struct {
	Contract *ISorareCardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ISorareCardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISorareCardsTransactorSession struct {
	Contract     *ISorareCardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ISorareCardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISorareCardsRaw struct {
	Contract *ISorareCards // Generic contract binding to access the raw methods on
}

// ISorareCardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISorareCardsCallerRaw struct {
	Contract *ISorareCardsCaller // Generic read-only contract binding to access the raw methods on
}

// ISorareCardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISorareCardsTransactorRaw struct {
	Contract *ISorareCardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISorareCards creates a new instance of ISorareCards, bound to a specific deployed contract.
func NewISorareCards(address common.Address, backend bind.ContractBackend) (*ISorareCards, error) {
	contract, err := bindISorareCards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISorareCards{ISorareCardsCaller: ISorareCardsCaller{contract: contract}, ISorareCardsTransactor: ISorareCardsTransactor{contract: contract}, ISorareCardsFilterer: ISorareCardsFilterer{contract: contract}}, nil
}

// NewISorareCardsCaller creates a new read-only instance of ISorareCards, bound to a specific deployed contract.
func NewISorareCardsCaller(address common.Address, caller bind.ContractCaller) (*ISorareCardsCaller, error) {
	contract, err := bindISorareCards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISorareCardsCaller{contract: contract}, nil
}

// NewISorareCardsTransactor creates a new write-only instance of ISorareCards, bound to a specific deployed contract.
func NewISorareCardsTransactor(address common.Address, transactor bind.ContractTransactor) (*ISorareCardsTransactor, error) {
	contract, err := bindISorareCards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISorareCardsTransactor{contract: contract}, nil
}

// NewISorareCardsFilterer creates a new log filterer instance of ISorareCards, bound to a specific deployed contract.
func NewISorareCardsFilterer(address common.Address, filterer bind.ContractFilterer) (*ISorareCardsFilterer, error) {
	contract, err := bindISorareCards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISorareCardsFilterer{contract: contract}, nil
}

// bindISorareCards binds a generic wrapper to an already deployed contract.
func bindISorareCards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISorareCardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISorareCards *ISorareCardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISorareCards.Contract.ISorareCardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISorareCards *ISorareCardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISorareCards.Contract.ISorareCardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISorareCards *ISorareCardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISorareCards.Contract.ISorareCardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISorareCards *ISorareCardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISorareCards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISorareCards *ISorareCardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISorareCards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISorareCards *ISorareCardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISorareCards.Contract.contract.Transact(opts, method, params...)
}

// CardExists is a free data retrieval call binding the contract method 0x4ae3540e.
//
// Solidity: function cardExists(uint256 cardId) view returns(bool)
func (_ISorareCards *ISorareCardsCaller) CardExists(opts *bind.CallOpts, cardId *big.Int) (bool, error) {
	var out []interface{}
	err := _ISorareCards.contract.Call(opts, &out, "cardExists", cardId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CardExists is a free data retrieval call binding the contract method 0x4ae3540e.
//
// Solidity: function cardExists(uint256 cardId) view returns(bool)
func (_ISorareCards *ISorareCardsSession) CardExists(cardId *big.Int) (bool, error) {
	return _ISorareCards.Contract.CardExists(&_ISorareCards.CallOpts, cardId)
}

// CardExists is a free data retrieval call binding the contract method 0x4ae3540e.
//
// Solidity: function cardExists(uint256 cardId) view returns(bool)
func (_ISorareCards *ISorareCardsCallerSession) CardExists(cardId *big.Int) (bool, error) {
	return _ISorareCards.Contract.CardExists(&_ISorareCards.CallOpts, cardId)
}

// GetCard is a free data retrieval call binding the contract method 0x9188d312.
//
// Solidity: function getCard(uint256 _cardId) view returns(uint256 playerId, uint16 season, uint256 scarcity, uint16 serialNumber, bytes metadata, uint16 clubId)
func (_ISorareCards *ISorareCardsCaller) GetCard(opts *bind.CallOpts, _cardId *big.Int) (struct {
	PlayerId     *big.Int
	Season       uint16
	Scarcity     *big.Int
	SerialNumber uint16
	Metadata     []byte
	ClubId       uint16
}, error) {
	var out []interface{}
	err := _ISorareCards.contract.Call(opts, &out, "getCard", _cardId)

	outstruct := new(struct {
		PlayerId     *big.Int
		Season       uint16
		Scarcity     *big.Int
		SerialNumber uint16
		Metadata     []byte
		ClubId       uint16
	})

	outstruct.PlayerId = out[0].(*big.Int)
	outstruct.Season = out[1].(uint16)
	outstruct.Scarcity = out[2].(*big.Int)
	outstruct.SerialNumber = out[3].(uint16)
	outstruct.Metadata = out[4].([]byte)
	outstruct.ClubId = out[5].(uint16)

	return *outstruct, err

}

// GetCard is a free data retrieval call binding the contract method 0x9188d312.
//
// Solidity: function getCard(uint256 _cardId) view returns(uint256 playerId, uint16 season, uint256 scarcity, uint16 serialNumber, bytes metadata, uint16 clubId)
func (_ISorareCards *ISorareCardsSession) GetCard(_cardId *big.Int) (struct {
	PlayerId     *big.Int
	Season       uint16
	Scarcity     *big.Int
	SerialNumber uint16
	Metadata     []byte
	ClubId       uint16
}, error) {
	return _ISorareCards.Contract.GetCard(&_ISorareCards.CallOpts, _cardId)
}

// GetCard is a free data retrieval call binding the contract method 0x9188d312.
//
// Solidity: function getCard(uint256 _cardId) view returns(uint256 playerId, uint16 season, uint256 scarcity, uint16 serialNumber, bytes metadata, uint16 clubId)
func (_ISorareCards *ISorareCardsCallerSession) GetCard(_cardId *big.Int) (struct {
	PlayerId     *big.Int
	Season       uint16
	Scarcity     *big.Int
	SerialNumber uint16
	Metadata     []byte
	ClubId       uint16
}, error) {
	return _ISorareCards.Contract.GetCard(&_ISorareCards.CallOpts, _cardId)
}

// GetClub is a free data retrieval call binding the contract method 0x4a39d2b8.
//
// Solidity: function getClub(uint16 clubId) view returns(string name, string country, string city, uint16 yearFounded)
func (_ISorareCards *ISorareCardsCaller) GetClub(opts *bind.CallOpts, clubId uint16) (struct {
	Name        string
	Country     string
	City        string
	YearFounded uint16
}, error) {
	var out []interface{}
	err := _ISorareCards.contract.Call(opts, &out, "getClub", clubId)

	outstruct := new(struct {
		Name        string
		Country     string
		City        string
		YearFounded uint16
	})

	outstruct.Name = out[0].(string)
	outstruct.Country = out[1].(string)
	outstruct.City = out[2].(string)
	outstruct.YearFounded = out[3].(uint16)

	return *outstruct, err

}

// GetClub is a free data retrieval call binding the contract method 0x4a39d2b8.
//
// Solidity: function getClub(uint16 clubId) view returns(string name, string country, string city, uint16 yearFounded)
func (_ISorareCards *ISorareCardsSession) GetClub(clubId uint16) (struct {
	Name        string
	Country     string
	City        string
	YearFounded uint16
}, error) {
	return _ISorareCards.Contract.GetClub(&_ISorareCards.CallOpts, clubId)
}

// GetClub is a free data retrieval call binding the contract method 0x4a39d2b8.
//
// Solidity: function getClub(uint16 clubId) view returns(string name, string country, string city, uint16 yearFounded)
func (_ISorareCards *ISorareCardsCallerSession) GetClub(clubId uint16) (struct {
	Name        string
	Country     string
	City        string
	YearFounded uint16
}, error) {
	return _ISorareCards.Contract.GetClub(&_ISorareCards.CallOpts, clubId)
}

// GetPlayer is a free data retrieval call binding the contract method 0xe55ae4e8.
//
// Solidity: function getPlayer(uint256 playerId) view returns(string name, uint16 yearOfBirth, uint8 monthOfBirth, uint8 dayOfBirth)
func (_ISorareCards *ISorareCardsCaller) GetPlayer(opts *bind.CallOpts, playerId *big.Int) (struct {
	Name         string
	YearOfBirth  uint16
	MonthOfBirth uint8
	DayOfBirth   uint8
}, error) {
	var out []interface{}
	err := _ISorareCards.contract.Call(opts, &out, "getPlayer", playerId)

	outstruct := new(struct {
		Name         string
		YearOfBirth  uint16
		MonthOfBirth uint8
		DayOfBirth   uint8
	})

	outstruct.Name = out[0].(string)
	outstruct.YearOfBirth = out[1].(uint16)
	outstruct.MonthOfBirth = out[2].(uint8)
	outstruct.DayOfBirth = out[3].(uint8)

	return *outstruct, err

}

// GetPlayer is a free data retrieval call binding the contract method 0xe55ae4e8.
//
// Solidity: function getPlayer(uint256 playerId) view returns(string name, uint16 yearOfBirth, uint8 monthOfBirth, uint8 dayOfBirth)
func (_ISorareCards *ISorareCardsSession) GetPlayer(playerId *big.Int) (struct {
	Name         string
	YearOfBirth  uint16
	MonthOfBirth uint8
	DayOfBirth   uint8
}, error) {
	return _ISorareCards.Contract.GetPlayer(&_ISorareCards.CallOpts, playerId)
}

// GetPlayer is a free data retrieval call binding the contract method 0xe55ae4e8.
//
// Solidity: function getPlayer(uint256 playerId) view returns(string name, uint16 yearOfBirth, uint8 monthOfBirth, uint8 dayOfBirth)
func (_ISorareCards *ISorareCardsCallerSession) GetPlayer(playerId *big.Int) (struct {
	Name         string
	YearOfBirth  uint16
	MonthOfBirth uint8
	DayOfBirth   uint8
}, error) {
	return _ISorareCards.Contract.GetPlayer(&_ISorareCards.CallOpts, playerId)
}

// CreateCard is a paid mutator transaction binding the contract method 0x01152383.
//
// Solidity: function createCard(uint256 playerId, uint16 season, uint8 scarcity, uint16 serialNumber, bytes32 metadata, uint16 clubId) returns(uint256)
func (_ISorareCards *ISorareCardsTransactor) CreateCard(opts *bind.TransactOpts, playerId *big.Int, season uint16, scarcity uint8, serialNumber uint16, metadata [32]byte, clubId uint16) (*types.Transaction, error) {
	return _ISorareCards.contract.Transact(opts, "createCard", playerId, season, scarcity, serialNumber, metadata, clubId)
}

// CreateCard is a paid mutator transaction binding the contract method 0x01152383.
//
// Solidity: function createCard(uint256 playerId, uint16 season, uint8 scarcity, uint16 serialNumber, bytes32 metadata, uint16 clubId) returns(uint256)
func (_ISorareCards *ISorareCardsSession) CreateCard(playerId *big.Int, season uint16, scarcity uint8, serialNumber uint16, metadata [32]byte, clubId uint16) (*types.Transaction, error) {
	return _ISorareCards.Contract.CreateCard(&_ISorareCards.TransactOpts, playerId, season, scarcity, serialNumber, metadata, clubId)
}

// CreateCard is a paid mutator transaction binding the contract method 0x01152383.
//
// Solidity: function createCard(uint256 playerId, uint16 season, uint8 scarcity, uint16 serialNumber, bytes32 metadata, uint16 clubId) returns(uint256)
func (_ISorareCards *ISorareCardsTransactorSession) CreateCard(playerId *big.Int, season uint16, scarcity uint8, serialNumber uint16, metadata [32]byte, clubId uint16) (*types.Transaction, error) {
	return _ISorareCards.Contract.CreateCard(&_ISorareCards.TransactOpts, playerId, season, scarcity, serialNumber, metadata, clubId)
}

// ISorareTokensABI is the input ABI used to generate the binding from.
const ISorareTokensABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"playerId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"season\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"scarcity\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"serialNumber\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"clubId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"createCardAndMintToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cardId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ISorareTokensFuncSigs maps the 4-byte function signature to its string representation.
var ISorareTokensFuncSigs = map[string]string{
	"2850d20e": "createCardAndMintToken(uint256,uint16,uint8,uint16,bytes32,uint16,address)",
	"a140ae23": "mintToken(uint256,address)",
}

// ISorareTokens is an auto generated Go binding around an Ethereum contract.
type ISorareTokens struct {
	ISorareTokensCaller     // Read-only binding to the contract
	ISorareTokensTransactor // Write-only binding to the contract
	ISorareTokensFilterer   // Log filterer for contract events
}

// ISorareTokensCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISorareTokensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISorareTokensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISorareTokensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISorareTokensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISorareTokensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISorareTokensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISorareTokensSession struct {
	Contract     *ISorareTokens    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISorareTokensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISorareTokensCallerSession struct {
	Contract *ISorareTokensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ISorareTokensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISorareTokensTransactorSession struct {
	Contract     *ISorareTokensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ISorareTokensRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISorareTokensRaw struct {
	Contract *ISorareTokens // Generic contract binding to access the raw methods on
}

// ISorareTokensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISorareTokensCallerRaw struct {
	Contract *ISorareTokensCaller // Generic read-only contract binding to access the raw methods on
}

// ISorareTokensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISorareTokensTransactorRaw struct {
	Contract *ISorareTokensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISorareTokens creates a new instance of ISorareTokens, bound to a specific deployed contract.
func NewISorareTokens(address common.Address, backend bind.ContractBackend) (*ISorareTokens, error) {
	contract, err := bindISorareTokens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISorareTokens{ISorareTokensCaller: ISorareTokensCaller{contract: contract}, ISorareTokensTransactor: ISorareTokensTransactor{contract: contract}, ISorareTokensFilterer: ISorareTokensFilterer{contract: contract}}, nil
}

// NewISorareTokensCaller creates a new read-only instance of ISorareTokens, bound to a specific deployed contract.
func NewISorareTokensCaller(address common.Address, caller bind.ContractCaller) (*ISorareTokensCaller, error) {
	contract, err := bindISorareTokens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISorareTokensCaller{contract: contract}, nil
}

// NewISorareTokensTransactor creates a new write-only instance of ISorareTokens, bound to a specific deployed contract.
func NewISorareTokensTransactor(address common.Address, transactor bind.ContractTransactor) (*ISorareTokensTransactor, error) {
	contract, err := bindISorareTokens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISorareTokensTransactor{contract: contract}, nil
}

// NewISorareTokensFilterer creates a new log filterer instance of ISorareTokens, bound to a specific deployed contract.
func NewISorareTokensFilterer(address common.Address, filterer bind.ContractFilterer) (*ISorareTokensFilterer, error) {
	contract, err := bindISorareTokens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISorareTokensFilterer{contract: contract}, nil
}

// bindISorareTokens binds a generic wrapper to an already deployed contract.
func bindISorareTokens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISorareTokensABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISorareTokens *ISorareTokensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISorareTokens.Contract.ISorareTokensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISorareTokens *ISorareTokensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISorareTokens.Contract.ISorareTokensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISorareTokens *ISorareTokensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISorareTokens.Contract.ISorareTokensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISorareTokens *ISorareTokensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISorareTokens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISorareTokens *ISorareTokensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISorareTokens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISorareTokens *ISorareTokensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISorareTokens.Contract.contract.Transact(opts, method, params...)
}

// CreateCardAndMintToken is a paid mutator transaction binding the contract method 0x2850d20e.
//
// Solidity: function createCardAndMintToken(uint256 playerId, uint16 season, uint8 scarcity, uint16 serialNumber, bytes32 metadata, uint16 clubId, address to) returns(uint256)
func (_ISorareTokens *ISorareTokensTransactor) CreateCardAndMintToken(opts *bind.TransactOpts, playerId *big.Int, season uint16, scarcity uint8, serialNumber uint16, metadata [32]byte, clubId uint16, to common.Address) (*types.Transaction, error) {
	return _ISorareTokens.contract.Transact(opts, "createCardAndMintToken", playerId, season, scarcity, serialNumber, metadata, clubId, to)
}

// CreateCardAndMintToken is a paid mutator transaction binding the contract method 0x2850d20e.
//
// Solidity: function createCardAndMintToken(uint256 playerId, uint16 season, uint8 scarcity, uint16 serialNumber, bytes32 metadata, uint16 clubId, address to) returns(uint256)
func (_ISorareTokens *ISorareTokensSession) CreateCardAndMintToken(playerId *big.Int, season uint16, scarcity uint8, serialNumber uint16, metadata [32]byte, clubId uint16, to common.Address) (*types.Transaction, error) {
	return _ISorareTokens.Contract.CreateCardAndMintToken(&_ISorareTokens.TransactOpts, playerId, season, scarcity, serialNumber, metadata, clubId, to)
}

// CreateCardAndMintToken is a paid mutator transaction binding the contract method 0x2850d20e.
//
// Solidity: function createCardAndMintToken(uint256 playerId, uint16 season, uint8 scarcity, uint16 serialNumber, bytes32 metadata, uint16 clubId, address to) returns(uint256)
func (_ISorareTokens *ISorareTokensTransactorSession) CreateCardAndMintToken(playerId *big.Int, season uint16, scarcity uint8, serialNumber uint16, metadata [32]byte, clubId uint16, to common.Address) (*types.Transaction, error) {
	return _ISorareTokens.Contract.CreateCardAndMintToken(&_ISorareTokens.TransactOpts, playerId, season, scarcity, serialNumber, metadata, clubId, to)
}

// MintToken is a paid mutator transaction binding the contract method 0xa140ae23.
//
// Solidity: function mintToken(uint256 cardId, address to) returns(uint256)
func (_ISorareTokens *ISorareTokensTransactor) MintToken(opts *bind.TransactOpts, cardId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ISorareTokens.contract.Transact(opts, "mintToken", cardId, to)
}

// MintToken is a paid mutator transaction binding the contract method 0xa140ae23.
//
// Solidity: function mintToken(uint256 cardId, address to) returns(uint256)
func (_ISorareTokens *ISorareTokensSession) MintToken(cardId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ISorareTokens.Contract.MintToken(&_ISorareTokens.TransactOpts, cardId, to)
}

// MintToken is a paid mutator transaction binding the contract method 0xa140ae23.
//
// Solidity: function mintToken(uint256 cardId, address to) returns(uint256)
func (_ISorareTokens *ISorareTokensTransactorSession) MintToken(cardId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ISorareTokens.Contract.MintToken(&_ISorareTokens.TransactOpts, cardId, to)
}

// MinterAccessABI is the input ABI used to generate the binding from.
const MinterAccessABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MinterAccessFuncSigs maps the 4-byte function signature to its string representation.
var MinterAccessFuncSigs = map[string]string{
	"a217fddf": "DEFAULT_ADMIN_ROLE()",
	"d5391393": "MINTER_ROLE()",
	"983b2d56": "addMinter(address)",
	"248a9ca3": "getRoleAdmin(bytes32)",
	"9010d07c": "getRoleMember(bytes32,uint256)",
	"ca15c873": "getRoleMemberCount(bytes32)",
	"2f2ff15d": "grantRole(bytes32,address)",
	"91d14854": "hasRole(bytes32,address)",
	"8da5cb5b": "owner()",
	"5f112c68": "renounceMinter(address)",
	"715018a6": "renounceOwnership()",
	"36568abe": "renounceRole(bytes32,address)",
	"cfbd4885": "revokeMinter(address)",
	"d547741f": "revokeRole(bytes32,address)",
	"f2fde38b": "transferOwnership(address)",
}

// MinterAccessBin is the compiled bytecode used for deploying new contracts.
var MinterAccessBin = "0x608060405234801561001057600080fd5b5060006100246001600160e01b036100ba16565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506100826000336001600160e01b036100be16565b604080516a4d494e5445525f524f4c4560a81b8152905190819003600b0190206100b590336001600160e01b036100be16565b6101e5565b3390565b6100d182826001600160e01b036100d516565b5050565b60008281526001602090815260409091206100f991839061081e610153821b17901c565b156100d15761010f6001600160e01b036100ba16565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000610171836001600160a01b0384166001600160e01b0361017a16565b90505b92915050565b600061018f83836001600160e01b036101cd16565b6101c557508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610174565b506000610174565b60009081526001919091016020526040902054151590565b610ae4806101f46000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806391d1485411610097578063cfbd488511610066578063cfbd488514610283578063d5391393146102a9578063d547741f146102b1578063f2fde38b146102dd576100f5565b806391d14854146101f8578063983b2d5614610238578063a217fddf1461025e578063ca15c87314610266576100f5565b80635f112c68116100d35780635f112c6814610183578063715018a6146101a95780638da5cb5b146101b15780639010d07c146101d5576100f5565b8063248a9ca3146100fa5780632f2ff15d1461012957806336568abe14610157575b600080fd5b6101176004803603602081101561011057600080fd5b5035610303565b60408051918252519081900360200190f35b6101556004803603604081101561013f57600080fd5b50803590602001356001600160a01b0316610318565b005b6101556004803603604081101561016d57600080fd5b50803590602001356001600160a01b0316610384565b6101556004803603602081101561019957600080fd5b50356001600160a01b03166103e5565b610155610412565b6101b96104c6565b604080516001600160a01b039092168252519081900360200190f35b6101b9600480360360408110156101eb57600080fd5b50803590602001356104d5565b6102246004803603604081101561020e57600080fd5b50803590602001356001600160a01b03166104fc565b604080519115158252519081900360200190f35b6101556004803603602081101561024e57600080fd5b50356001600160a01b031661051a565b610117610544565b6101176004803603602081101561027c57600080fd5b5035610549565b6101556004803603602081101561029957600080fd5b50356001600160a01b0316610560565b61011761058a565b610155600480360360408110156102c757600080fd5b50803590602001356001600160a01b03166105ad565b610155600480360360208110156102f357600080fd5b50356001600160a01b0316610606565b60009081526001602052604090206002015490565b60008281526001602052604090206002015461033b90610336610710565b6104fc565b6103765760405162461bcd60e51b815260040180806020018281038252602f8152602001806109fb602f913960400191505060405180910390fd5b6103808282610714565b5050565b61038c610710565b6001600160a01b0316816001600160a01b0316146103db5760405162461bcd60e51b815260040180806020018281038252602f815260200180610a80602f913960400191505060405180910390fd5b6103808282610783565b604080516a4d494e5445525f524f4c4560a81b8152905190819003600b01902061040f9082610384565b50565b61041a610710565b6000546001600160a01b0390811691161461047c576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b60008281526001602052604081206104f3908363ffffffff6107f216565b90505b92915050565b60008281526001602052604081206104f3908363ffffffff6107fe16565b604080516a4d494e5445525f524f4c4560a81b8152905190819003600b01902061040f9082610318565b600081565b60008181526001602052604081206104f690610813565b604080516a4d494e5445525f524f4c4560a81b8152905190819003600b01902061040f90826105ad565b604080516a4d494e5445525f524f4c4560a81b8152905190819003600b01902081565b6000828152600160205260409020600201546105cb90610336610710565b6103db5760405162461bcd60e51b8152600401808060200182810382526030815260200180610a506030913960400191505060405180910390fd5b61060e610710565b6000546001600160a01b03908116911614610670576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b0381166106b55760405162461bcd60e51b8152600401808060200182810382526026815260200180610a2a6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b6000828152600160205260409020610732908263ffffffff61081e16565b156103805761073f610710565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60008281526001602052604090206107a1908263ffffffff61083316565b15610380576107ae610710565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b60006104f38383610848565b60006104f3836001600160a01b0384166108ac565b60006104f6826108c4565b60006104f3836001600160a01b0384166108c8565b60006104f3836001600160a01b038416610912565b8154600090821061088a5760405162461bcd60e51b81526004018080602001828103825260228152602001806109d96022913960400191505060405180910390fd5b82600001828154811061089957fe5b9060005260206000200154905092915050565b60009081526001919091016020526040902054151590565b5490565b60006108d483836108ac565b61090a575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556104f6565b5060006104f6565b600081815260018301602052604081205480156109ce578354600019808301919081019060009087908390811061094557fe5b906000526020600020015490508087600001848154811061096257fe5b60009182526020808320909101929092558281526001898101909252604090209084019055865487908061099257fe5b600190038181906000526020600020016000905590558660010160008781526020019081526020016000206000905560019450505050506104f6565b60009150506104f656fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e6473416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f206772616e744f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f207265766f6b65416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636520726f6c657320666f722073656c66a264697066735822122078fe2824d83f5f2fa05f30dfd7de0645463e423e9a9bbc5f318c39565b4fb0b964736f6c63430006060033"

// DeployMinterAccess deploys a new Ethereum contract, binding an instance of MinterAccess to it.
func DeployMinterAccess(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MinterAccess, error) {
	parsed, err := abi.JSON(strings.NewReader(MinterAccessABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MinterAccessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MinterAccess{MinterAccessCaller: MinterAccessCaller{contract: contract}, MinterAccessTransactor: MinterAccessTransactor{contract: contract}, MinterAccessFilterer: MinterAccessFilterer{contract: contract}}, nil
}

// MinterAccess is an auto generated Go binding around an Ethereum contract.
type MinterAccess struct {
	MinterAccessCaller     // Read-only binding to the contract
	MinterAccessTransactor // Write-only binding to the contract
	MinterAccessFilterer   // Log filterer for contract events
}

// MinterAccessCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinterAccessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterAccessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinterAccessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterAccessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinterAccessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterAccessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinterAccessSession struct {
	Contract     *MinterAccess     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinterAccessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinterAccessCallerSession struct {
	Contract *MinterAccessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MinterAccessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinterAccessTransactorSession struct {
	Contract     *MinterAccessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MinterAccessRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinterAccessRaw struct {
	Contract *MinterAccess // Generic contract binding to access the raw methods on
}

// MinterAccessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinterAccessCallerRaw struct {
	Contract *MinterAccessCaller // Generic read-only contract binding to access the raw methods on
}

// MinterAccessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinterAccessTransactorRaw struct {
	Contract *MinterAccessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinterAccess creates a new instance of MinterAccess, bound to a specific deployed contract.
func NewMinterAccess(address common.Address, backend bind.ContractBackend) (*MinterAccess, error) {
	contract, err := bindMinterAccess(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinterAccess{MinterAccessCaller: MinterAccessCaller{contract: contract}, MinterAccessTransactor: MinterAccessTransactor{contract: contract}, MinterAccessFilterer: MinterAccessFilterer{contract: contract}}, nil
}

// NewMinterAccessCaller creates a new read-only instance of MinterAccess, bound to a specific deployed contract.
func NewMinterAccessCaller(address common.Address, caller bind.ContractCaller) (*MinterAccessCaller, error) {
	contract, err := bindMinterAccess(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinterAccessCaller{contract: contract}, nil
}

// NewMinterAccessTransactor creates a new write-only instance of MinterAccess, bound to a specific deployed contract.
func NewMinterAccessTransactor(address common.Address, transactor bind.ContractTransactor) (*MinterAccessTransactor, error) {
	contract, err := bindMinterAccess(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinterAccessTransactor{contract: contract}, nil
}

// NewMinterAccessFilterer creates a new log filterer instance of MinterAccess, bound to a specific deployed contract.
func NewMinterAccessFilterer(address common.Address, filterer bind.ContractFilterer) (*MinterAccessFilterer, error) {
	contract, err := bindMinterAccess(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinterAccessFilterer{contract: contract}, nil
}

// bindMinterAccess binds a generic wrapper to an already deployed contract.
func bindMinterAccess(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinterAccessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterAccess *MinterAccessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinterAccess.Contract.MinterAccessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterAccess *MinterAccessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterAccess.Contract.MinterAccessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterAccess *MinterAccessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterAccess.Contract.MinterAccessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterAccess *MinterAccessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinterAccess.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterAccess *MinterAccessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterAccess.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterAccess *MinterAccessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterAccess.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MinterAccess *MinterAccessCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MinterAccess.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MinterAccess *MinterAccessSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MinterAccess.Contract.DEFAULTADMINROLE(&_MinterAccess.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MinterAccess *MinterAccessCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MinterAccess.Contract.DEFAULTADMINROLE(&_MinterAccess.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_MinterAccess *MinterAccessCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MinterAccess.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_MinterAccess *MinterAccessSession) MINTERROLE() ([32]byte, error) {
	return _MinterAccess.Contract.MINTERROLE(&_MinterAccess.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_MinterAccess *MinterAccessCallerSession) MINTERROLE() ([32]byte, error) {
	return _MinterAccess.Contract.MINTERROLE(&_MinterAccess.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MinterAccess *MinterAccessCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MinterAccess.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MinterAccess *MinterAccessSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MinterAccess.Contract.GetRoleAdmin(&_MinterAccess.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MinterAccess *MinterAccessCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MinterAccess.Contract.GetRoleAdmin(&_MinterAccess.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MinterAccess *MinterAccessCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MinterAccess.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MinterAccess *MinterAccessSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MinterAccess.Contract.GetRoleMember(&_MinterAccess.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MinterAccess *MinterAccessCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MinterAccess.Contract.GetRoleMember(&_MinterAccess.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MinterAccess *MinterAccessCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MinterAccess.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MinterAccess *MinterAccessSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MinterAccess.Contract.GetRoleMemberCount(&_MinterAccess.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MinterAccess *MinterAccessCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MinterAccess.Contract.GetRoleMemberCount(&_MinterAccess.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MinterAccess *MinterAccessCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MinterAccess.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MinterAccess *MinterAccessSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MinterAccess.Contract.HasRole(&_MinterAccess.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MinterAccess *MinterAccessCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MinterAccess.Contract.HasRole(&_MinterAccess.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MinterAccess *MinterAccessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinterAccess.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MinterAccess *MinterAccessSession) Owner() (common.Address, error) {
	return _MinterAccess.Contract.Owner(&_MinterAccess.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MinterAccess *MinterAccessCallerSession) Owner() (common.Address, error) {
	return _MinterAccess.Contract.Owner(&_MinterAccess.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_MinterAccess *MinterAccessTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MinterAccess.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_MinterAccess *MinterAccessSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _MinterAccess.Contract.AddMinter(&_MinterAccess.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_MinterAccess *MinterAccessTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _MinterAccess.Contract.AddMinter(&_MinterAccess.TransactOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MinterAccess *MinterAccessTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MinterAccess.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MinterAccess *MinterAccessSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MinterAccess.Contract.GrantRole(&_MinterAccess.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MinterAccess *MinterAccessTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MinterAccess.Contract.GrantRole(&_MinterAccess.TransactOpts, role, account)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x5f112c68.
//
// Solidity: function renounceMinter(address account) returns()
func (_MinterAccess *MinterAccessTransactor) RenounceMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MinterAccess.contract.Transact(opts, "renounceMinter", account)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x5f112c68.
//
// Solidity: function renounceMinter(address account) returns()
func (_MinterAccess *MinterAccessSession) RenounceMinter(account common.Address) (*types.Transaction, error) {
	return _MinterAccess.Contract.RenounceMinter(&_MinterAccess.TransactOpts, account)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x5f112c68.
//
// Solidity: function renounceMinter(address account) returns()
func (_MinterAccess *MinterAccessTransactorSession) RenounceMinter(account common.Address) (*types.Transaction, error) {
	return _MinterAccess.Contract.RenounceMinter(&_MinterAccess.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MinterAccess *MinterAccessTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterAccess.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MinterAccess *MinterAccessSession) RenounceOwnership() (*types.Transaction, error) {
	return _MinterAccess.Contract.RenounceOwnership(&_MinterAccess.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MinterAccess *MinterAccessTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MinterAccess.Contract.RenounceOwnership(&_MinterAccess.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MinterAccess *MinterAccessTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MinterAccess.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MinterAccess *MinterAccessSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MinterAccess.Contract.RenounceRole(&_MinterAccess.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MinterAccess *MinterAccessTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MinterAccess.Contract.RenounceRole(&_MinterAccess.TransactOpts, role, account)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address account) returns()
func (_MinterAccess *MinterAccessTransactor) RevokeMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MinterAccess.contract.Transact(opts, "revokeMinter", account)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address account) returns()
func (_MinterAccess *MinterAccessSession) RevokeMinter(account common.Address) (*types.Transaction, error) {
	return _MinterAccess.Contract.RevokeMinter(&_MinterAccess.TransactOpts, account)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address account) returns()
func (_MinterAccess *MinterAccessTransactorSession) RevokeMinter(account common.Address) (*types.Transaction, error) {
	return _MinterAccess.Contract.RevokeMinter(&_MinterAccess.TransactOpts, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MinterAccess *MinterAccessTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MinterAccess.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MinterAccess *MinterAccessSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MinterAccess.Contract.RevokeRole(&_MinterAccess.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MinterAccess *MinterAccessTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MinterAccess.Contract.RevokeRole(&_MinterAccess.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MinterAccess *MinterAccessTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MinterAccess.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MinterAccess *MinterAccessSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MinterAccess.Contract.TransferOwnership(&_MinterAccess.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MinterAccess *MinterAccessTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MinterAccess.Contract.TransferOwnership(&_MinterAccess.TransactOpts, newOwner)
}

// MinterAccessOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MinterAccess contract.
type MinterAccessOwnershipTransferredIterator struct {
	Event *MinterAccessOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MinterAccessOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterAccessOwnershipTransferred)
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
		it.Event = new(MinterAccessOwnershipTransferred)
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
func (it *MinterAccessOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterAccessOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterAccessOwnershipTransferred represents a OwnershipTransferred event raised by the MinterAccess contract.
type MinterAccessOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MinterAccess *MinterAccessFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MinterAccessOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MinterAccess.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MinterAccessOwnershipTransferredIterator{contract: _MinterAccess.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MinterAccess *MinterAccessFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MinterAccessOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MinterAccess.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterAccessOwnershipTransferred)
				if err := _MinterAccess.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MinterAccess *MinterAccessFilterer) ParseOwnershipTransferred(log types.Log) (*MinterAccessOwnershipTransferred, error) {
	event := new(MinterAccessOwnershipTransferred)
	if err := _MinterAccess.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterAccessRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MinterAccess contract.
type MinterAccessRoleGrantedIterator struct {
	Event *MinterAccessRoleGranted // Event containing the contract specifics and raw log

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
func (it *MinterAccessRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterAccessRoleGranted)
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
		it.Event = new(MinterAccessRoleGranted)
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
func (it *MinterAccessRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterAccessRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterAccessRoleGranted represents a RoleGranted event raised by the MinterAccess contract.
type MinterAccessRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MinterAccess *MinterAccessFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MinterAccessRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MinterAccess.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MinterAccessRoleGrantedIterator{contract: _MinterAccess.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MinterAccess *MinterAccessFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MinterAccessRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MinterAccess.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterAccessRoleGranted)
				if err := _MinterAccess.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MinterAccess *MinterAccessFilterer) ParseRoleGranted(log types.Log) (*MinterAccessRoleGranted, error) {
	event := new(MinterAccessRoleGranted)
	if err := _MinterAccess.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterAccessRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MinterAccess contract.
type MinterAccessRoleRevokedIterator struct {
	Event *MinterAccessRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MinterAccessRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterAccessRoleRevoked)
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
		it.Event = new(MinterAccessRoleRevoked)
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
func (it *MinterAccessRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterAccessRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterAccessRoleRevoked represents a RoleRevoked event raised by the MinterAccess contract.
type MinterAccessRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MinterAccess *MinterAccessFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MinterAccessRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MinterAccess.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MinterAccessRoleRevokedIterator{contract: _MinterAccess.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MinterAccess *MinterAccessFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MinterAccessRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MinterAccess.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterAccessRoleRevoked)
				if err := _MinterAccess.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MinterAccess *MinterAccessFilterer) ParseRoleRevoked(log types.Log) (*MinterAccessRoleRevoked, error) {
	event := new(MinterAccessRoleRevoked)
	if err := _MinterAccess.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTClientABI is the input ABI used to generate the binding from.
const NFTClientABI = "[{\"inputs\":[],\"name\":\"interfaceIdERC721\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"requireERC721\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"page\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rows\",\"type\":\"uint8\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NFTClientFuncSigs maps the 4-byte function signature to its string representation.
var NFTClientFuncSigs = map[string]string{
	"40edd555": "interfaceIdERC721()",
	"0b1a6d6a": "requireERC721(address)",
	"36b7a69c": "tokensOfOwner(address,address,uint8,uint8)",
	"dc800b0d": "transferAll(IERC721Enumerable,address,address)",
	"40125765": "transferTokens(IERC721Enumerable,address,address,uint256[])",
}

// NFTClientBin is the compiled bytecode used for deploying new contracts.
var NFTClientBin = "0x6107bf610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100615760003560e01c80630b1a6d6a1461006657806336b7a69c1461008e578063401257651461011f57806340edd555146101f0578063dc800b0d14610215575b600080fd5b61008c6004803603602081101561007c57600080fd5b50356001600160a01b031661025a565b005b6100cf600480360360808110156100a457600080fd5b506001600160a01b03813581169160208101359091169060ff6040820135811691606001351661031a565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561010b5781810151838201526020016100f3565b505050509050019250505060405180910390f35b81801561012b57600080fd5b5061008c6004803603608081101561014257600080fd5b6001600160a01b038235811692602081013582169260408201359092169181019060808101606082013564010000000081111561017e57600080fd5b82018360208201111561019057600080fd5b803590602001918460208302840111640100000000831117156101b257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506104d6945050505050565b6101f86105bb565b604080516001600160e01b03199092168252519081900360200190f35b81801561022157600080fd5b5061008c6004803603606081101561023857600080fd5b506001600160a01b0381358116916020810135821691604090910135166105c6565b604080516301ffc9a760e01b81526380ac58cd60e01b600482015290516001600160a01b038316916301ffc9a7916024808301926020929190829003018186803b1580156102a757600080fd5b505afa1580156102bb573d6000803e3d6000fd5b505050506040513d60208110156102d157600080fd5b5051610317576040805162461bcd60e51b815260206004820152601060248201526f24a9afa727aa2f9b9918afaa27a5a2a760811b604482015290519081900360640190fd5b50565b60606103258561025a565b60008590506000816001600160a01b03166370a08231876040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561038257600080fd5b505afa158015610396573d6000803e3d6000fd5b505050506040513d60208110156103ac57600080fd5b5051905060ff8585021660008282116103d3576103ce8284038760ff16610771565b6103d6565b60005b905060608167ffffffffffffffff811180156103f157600080fd5b5060405190808252806020026020018201604052801561041b578160200160208202803683370190505b50905060005b828110156104c85760408051632f745c5960e01b81526001600160a01b038c811660048301528684016024830152915191881691632f745c5991604480820192602092909190829003018186803b15801561047b57600080fd5b505afa15801561048f573d6000803e3d6000fd5b505050506040513d60208110156104a557600080fd5b505182518390839081106104b557fe5b6020908102919091010152600101610421565b509998505050505050505050565b60005b81518110156105b4578181815181106104ee57fe5b602002602001015160001415610503576105b4565b846001600160a01b03166342842e0e858585858151811061052057fe5b60200260200101516040518463ffffffff1660e01b815260040180846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b031681526020018281526020019350505050600060405180830381600087803b15801561059057600080fd5b505af11580156105a4573d6000803e3d6000fd5b5050600190920191506104d99050565b5050505050565b6380ac58cd60e01b81565b6000836001600160a01b03166370a08231846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561061e57600080fd5b505afa158015610632573d6000803e3d6000fd5b505050506040513d602081101561064857600080fd5b505190505b801561076b57836001600160a01b03166342842e0e8484876001600160a01b0316632f745c5988600188036040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b031681526020018281526020019250505060206040518083038186803b1580156106c557600080fd5b505afa1580156106d9573d6000803e3d6000fd5b505050506040513d60208110156106ef57600080fd5b5051604080516001600160e01b031960e087901b1681526001600160a01b03948516600482015292909316602483015260448201529051606480830192600092919082900301818387803b15801561074657600080fd5b505af115801561075a573d6000803e3d6000fd5b5050600019909201915061064d9050565b50505050565b60008183116107805782610782565b815b939250505056fea2646970667358221220defc145e0d95c16cbb428654ff8a62e070053292ddea66244b14d28eb036ab7864736f6c63430006060033"

// DeployNFTClient deploys a new Ethereum contract, binding an instance of NFTClient to it.
func DeployNFTClient(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NFTClient, error) {
	parsed, err := abi.JSON(strings.NewReader(NFTClientABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NFTClientBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NFTClient{NFTClientCaller: NFTClientCaller{contract: contract}, NFTClientTransactor: NFTClientTransactor{contract: contract}, NFTClientFilterer: NFTClientFilterer{contract: contract}}, nil
}

// NFTClient is an auto generated Go binding around an Ethereum contract.
type NFTClient struct {
	NFTClientCaller     // Read-only binding to the contract
	NFTClientTransactor // Write-only binding to the contract
	NFTClientFilterer   // Log filterer for contract events
}

// NFTClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTClientSession struct {
	Contract     *NFTClient        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTClientCallerSession struct {
	Contract *NFTClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// NFTClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTClientTransactorSession struct {
	Contract     *NFTClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NFTClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTClientRaw struct {
	Contract *NFTClient // Generic contract binding to access the raw methods on
}

// NFTClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTClientCallerRaw struct {
	Contract *NFTClientCaller // Generic read-only contract binding to access the raw methods on
}

// NFTClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTClientTransactorRaw struct {
	Contract *NFTClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFTClient creates a new instance of NFTClient, bound to a specific deployed contract.
func NewNFTClient(address common.Address, backend bind.ContractBackend) (*NFTClient, error) {
	contract, err := bindNFTClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFTClient{NFTClientCaller: NFTClientCaller{contract: contract}, NFTClientTransactor: NFTClientTransactor{contract: contract}, NFTClientFilterer: NFTClientFilterer{contract: contract}}, nil
}

// NewNFTClientCaller creates a new read-only instance of NFTClient, bound to a specific deployed contract.
func NewNFTClientCaller(address common.Address, caller bind.ContractCaller) (*NFTClientCaller, error) {
	contract, err := bindNFTClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTClientCaller{contract: contract}, nil
}

// NewNFTClientTransactor creates a new write-only instance of NFTClient, bound to a specific deployed contract.
func NewNFTClientTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTClientTransactor, error) {
	contract, err := bindNFTClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTClientTransactor{contract: contract}, nil
}

// NewNFTClientFilterer creates a new log filterer instance of NFTClient, bound to a specific deployed contract.
func NewNFTClientFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTClientFilterer, error) {
	contract, err := bindNFTClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTClientFilterer{contract: contract}, nil
}

// bindNFTClient binds a generic wrapper to an already deployed contract.
func bindNFTClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NFTClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTClient *NFTClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTClient.Contract.NFTClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTClient *NFTClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTClient.Contract.NFTClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTClient *NFTClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTClient.Contract.NFTClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTClient *NFTClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTClient *NFTClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTClient *NFTClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTClient.Contract.contract.Transact(opts, method, params...)
}

// InterfaceIdERC721 is a free data retrieval call binding the contract method 0x40edd555.
//
// Solidity: function interfaceIdERC721() view returns(bytes4)
func (_NFTClient *NFTClientCaller) InterfaceIdERC721(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _NFTClient.contract.Call(opts, &out, "interfaceIdERC721")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// InterfaceIdERC721 is a free data retrieval call binding the contract method 0x40edd555.
//
// Solidity: function interfaceIdERC721() view returns(bytes4)
func (_NFTClient *NFTClientSession) InterfaceIdERC721() ([4]byte, error) {
	return _NFTClient.Contract.InterfaceIdERC721(&_NFTClient.CallOpts)
}

// InterfaceIdERC721 is a free data retrieval call binding the contract method 0x40edd555.
//
// Solidity: function interfaceIdERC721() view returns(bytes4)
func (_NFTClient *NFTClientCallerSession) InterfaceIdERC721() ([4]byte, error) {
	return _NFTClient.Contract.InterfaceIdERC721(&_NFTClient.CallOpts)
}

// RequireERC721 is a free data retrieval call binding the contract method 0x0b1a6d6a.
//
// Solidity: function requireERC721(address _candidate) view returns()
func (_NFTClient *NFTClientCaller) RequireERC721(opts *bind.CallOpts, _candidate common.Address) error {
	var out []interface{}
	err := _NFTClient.contract.Call(opts, &out, "requireERC721", _candidate)

	if err != nil {
		return err
	}

	return err

}

// RequireERC721 is a free data retrieval call binding the contract method 0x0b1a6d6a.
//
// Solidity: function requireERC721(address _candidate) view returns()
func (_NFTClient *NFTClientSession) RequireERC721(_candidate common.Address) error {
	return _NFTClient.Contract.RequireERC721(&_NFTClient.CallOpts, _candidate)
}

// RequireERC721 is a free data retrieval call binding the contract method 0x0b1a6d6a.
//
// Solidity: function requireERC721(address _candidate) view returns()
func (_NFTClient *NFTClientCallerSession) RequireERC721(_candidate common.Address) error {
	return _NFTClient.Contract.RequireERC721(&_NFTClient.CallOpts, _candidate)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x36b7a69c.
//
// Solidity: function tokensOfOwner(address _nftContract, address owner, uint8 page, uint8 rows) view returns(uint256[])
func (_NFTClient *NFTClientCaller) TokensOfOwner(opts *bind.CallOpts, _nftContract common.Address, owner common.Address, page uint8, rows uint8) ([]*big.Int, error) {
	var out []interface{}
	err := _NFTClient.contract.Call(opts, &out, "tokensOfOwner", _nftContract, owner, page, rows)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x36b7a69c.
//
// Solidity: function tokensOfOwner(address _nftContract, address owner, uint8 page, uint8 rows) view returns(uint256[])
func (_NFTClient *NFTClientSession) TokensOfOwner(_nftContract common.Address, owner common.Address, page uint8, rows uint8) ([]*big.Int, error) {
	return _NFTClient.Contract.TokensOfOwner(&_NFTClient.CallOpts, _nftContract, owner, page, rows)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x36b7a69c.
//
// Solidity: function tokensOfOwner(address _nftContract, address owner, uint8 page, uint8 rows) view returns(uint256[])
func (_NFTClient *NFTClientCallerSession) TokensOfOwner(_nftContract common.Address, owner common.Address, page uint8, rows uint8) ([]*big.Int, error) {
	return _NFTClient.Contract.TokensOfOwner(&_NFTClient.CallOpts, _nftContract, owner, page, rows)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayRecipientABI is the input ABI used to generate the binding from.
const RelayRecipientABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"blockRelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelayAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RelayRecipientFuncSigs maps the 4-byte function signature to its string representation.
var RelayRecipientFuncSigs = map[string]string{
	"309e11fc": "blockRelay()",
	"32deedc6": "getRelayAddress()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// RelayRecipientBin is the compiled bytecode used for deploying new contracts.
var RelayRecipientBin = "0x608060405234801561001057600080fd5b506040516105903803806105908339818101604052602081101561003357600080fd5b505160006100486001600160e01b0361011216565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506001600160a01b0381166100ed576040805162461bcd60e51b815260206004820181905260248201527f437573746f6d2072656c61792061646472657373206973207265717569726564604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055610192565b6001546000906001600160a01b0316331461012e575033610142565b61013f6001600160e01b0361014516565b90505b90565b600060606000368080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050503601516001600160a01b031692915050565b6103ef806101a16000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063309e11fc1461005c57806332deedc614610066578063715018a61461008a5780638da5cb5b14610092578063f2fde38b1461009a575b600080fd5b6100646100c0565b005b61006e61013e565b604080516001600160a01b039092168252519081900360200190f35b61006461014e565b61006e610202565b610064600480360360208110156100b057600080fd5b50356001600160a01b0316610211565b6100c861031b565b6000546001600160a01b0390811691161461012a576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600180546001600160a01b03191630179055565b6001546001600160a01b03165b90565b61015661031b565b6000546001600160a01b039081169116146101b8576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b61021961031b565b6000546001600160a01b0390811691161461027b576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b0381166102c05760405162461bcd60e51b81526004018080602001828103825260268152602001806103946026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001546000906001600160a01b0316331461033757503361014b565b61033f610346565b905061014b565b600060606000368080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050503601516001600160a01b03169291505056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373a26469706673582212200ffa9ae5da0047e9678bc345f517b001aa003dfc6c0db4a3360d5d60a60a540d64736f6c63430006060033"

// DeployRelayRecipient deploys a new Ethereum contract, binding an instance of RelayRecipient to it.
func DeployRelayRecipient(auth *bind.TransactOpts, backend bind.ContractBackend, relayAddress common.Address) (common.Address, *types.Transaction, *RelayRecipient, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayRecipientABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RelayRecipientBin), backend, relayAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RelayRecipient{RelayRecipientCaller: RelayRecipientCaller{contract: contract}, RelayRecipientTransactor: RelayRecipientTransactor{contract: contract}, RelayRecipientFilterer: RelayRecipientFilterer{contract: contract}}, nil
}

// RelayRecipient is an auto generated Go binding around an Ethereum contract.
type RelayRecipient struct {
	RelayRecipientCaller     // Read-only binding to the contract
	RelayRecipientTransactor // Write-only binding to the contract
	RelayRecipientFilterer   // Log filterer for contract events
}

// RelayRecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayRecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayRecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayRecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayRecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayRecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayRecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayRecipientSession struct {
	Contract     *RelayRecipient   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayRecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayRecipientCallerSession struct {
	Contract *RelayRecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RelayRecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayRecipientTransactorSession struct {
	Contract     *RelayRecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RelayRecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayRecipientRaw struct {
	Contract *RelayRecipient // Generic contract binding to access the raw methods on
}

// RelayRecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayRecipientCallerRaw struct {
	Contract *RelayRecipientCaller // Generic read-only contract binding to access the raw methods on
}

// RelayRecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayRecipientTransactorRaw struct {
	Contract *RelayRecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayRecipient creates a new instance of RelayRecipient, bound to a specific deployed contract.
func NewRelayRecipient(address common.Address, backend bind.ContractBackend) (*RelayRecipient, error) {
	contract, err := bindRelayRecipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RelayRecipient{RelayRecipientCaller: RelayRecipientCaller{contract: contract}, RelayRecipientTransactor: RelayRecipientTransactor{contract: contract}, RelayRecipientFilterer: RelayRecipientFilterer{contract: contract}}, nil
}

// NewRelayRecipientCaller creates a new read-only instance of RelayRecipient, bound to a specific deployed contract.
func NewRelayRecipientCaller(address common.Address, caller bind.ContractCaller) (*RelayRecipientCaller, error) {
	contract, err := bindRelayRecipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayRecipientCaller{contract: contract}, nil
}

// NewRelayRecipientTransactor creates a new write-only instance of RelayRecipient, bound to a specific deployed contract.
func NewRelayRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayRecipientTransactor, error) {
	contract, err := bindRelayRecipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayRecipientTransactor{contract: contract}, nil
}

// NewRelayRecipientFilterer creates a new log filterer instance of RelayRecipient, bound to a specific deployed contract.
func NewRelayRecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayRecipientFilterer, error) {
	contract, err := bindRelayRecipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayRecipientFilterer{contract: contract}, nil
}

// bindRelayRecipient binds a generic wrapper to an already deployed contract.
func bindRelayRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayRecipientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayRecipient *RelayRecipientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayRecipient.Contract.RelayRecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayRecipient *RelayRecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayRecipient.Contract.RelayRecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayRecipient *RelayRecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayRecipient.Contract.RelayRecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayRecipient *RelayRecipientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayRecipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayRecipient *RelayRecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayRecipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayRecipient *RelayRecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayRecipient.Contract.contract.Transact(opts, method, params...)
}

// GetRelayAddress is a free data retrieval call binding the contract method 0x32deedc6.
//
// Solidity: function getRelayAddress() view returns(address)
func (_RelayRecipient *RelayRecipientCaller) GetRelayAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayRecipient.contract.Call(opts, &out, "getRelayAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRelayAddress is a free data retrieval call binding the contract method 0x32deedc6.
//
// Solidity: function getRelayAddress() view returns(address)
func (_RelayRecipient *RelayRecipientSession) GetRelayAddress() (common.Address, error) {
	return _RelayRecipient.Contract.GetRelayAddress(&_RelayRecipient.CallOpts)
}

// GetRelayAddress is a free data retrieval call binding the contract method 0x32deedc6.
//
// Solidity: function getRelayAddress() view returns(address)
func (_RelayRecipient *RelayRecipientCallerSession) GetRelayAddress() (common.Address, error) {
	return _RelayRecipient.Contract.GetRelayAddress(&_RelayRecipient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RelayRecipient *RelayRecipientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayRecipient.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RelayRecipient *RelayRecipientSession) Owner() (common.Address, error) {
	return _RelayRecipient.Contract.Owner(&_RelayRecipient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RelayRecipient *RelayRecipientCallerSession) Owner() (common.Address, error) {
	return _RelayRecipient.Contract.Owner(&_RelayRecipient.CallOpts)
}

// BlockRelay is a paid mutator transaction binding the contract method 0x309e11fc.
//
// Solidity: function blockRelay() returns()
func (_RelayRecipient *RelayRecipientTransactor) BlockRelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayRecipient.contract.Transact(opts, "blockRelay")
}

// BlockRelay is a paid mutator transaction binding the contract method 0x309e11fc.
//
// Solidity: function blockRelay() returns()
func (_RelayRecipient *RelayRecipientSession) BlockRelay() (*types.Transaction, error) {
	return _RelayRecipient.Contract.BlockRelay(&_RelayRecipient.TransactOpts)
}

// BlockRelay is a paid mutator transaction binding the contract method 0x309e11fc.
//
// Solidity: function blockRelay() returns()
func (_RelayRecipient *RelayRecipientTransactorSession) BlockRelay() (*types.Transaction, error) {
	return _RelayRecipient.Contract.BlockRelay(&_RelayRecipient.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RelayRecipient *RelayRecipientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayRecipient.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RelayRecipient *RelayRecipientSession) RenounceOwnership() (*types.Transaction, error) {
	return _RelayRecipient.Contract.RenounceOwnership(&_RelayRecipient.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RelayRecipient *RelayRecipientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RelayRecipient.Contract.RenounceOwnership(&_RelayRecipient.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RelayRecipient *RelayRecipientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RelayRecipient.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RelayRecipient *RelayRecipientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RelayRecipient.Contract.TransferOwnership(&_RelayRecipient.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RelayRecipient *RelayRecipientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RelayRecipient.Contract.TransferOwnership(&_RelayRecipient.TransactOpts, newOwner)
}

// RelayRecipientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RelayRecipient contract.
type RelayRecipientOwnershipTransferredIterator struct {
	Event *RelayRecipientOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RelayRecipientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayRecipientOwnershipTransferred)
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
		it.Event = new(RelayRecipientOwnershipTransferred)
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
func (it *RelayRecipientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayRecipientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayRecipientOwnershipTransferred represents a OwnershipTransferred event raised by the RelayRecipient contract.
type RelayRecipientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RelayRecipient *RelayRecipientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RelayRecipientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RelayRecipient.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RelayRecipientOwnershipTransferredIterator{contract: _RelayRecipient.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RelayRecipient *RelayRecipientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RelayRecipientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RelayRecipient.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayRecipientOwnershipTransferred)
				if err := _RelayRecipient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RelayRecipient *RelayRecipientFilterer) ParseOwnershipTransferred(log types.Log) (*RelayRecipientOwnershipTransferred, error) {
	event := new(RelayRecipientOwnershipTransferred)
	if err := _RelayRecipient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220198782bac737d17dca7f599e0b6e936a9acdee55337b153e7f0e9c9210e75b9064736f6c63430006060033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SorareTokensABI is the input ABI used to generate the binding from.
const SorareTokensABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sorareCardsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockRelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"playerId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"season\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"scarcity\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"serialNumber\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"clubId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"createCardAndMintToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getCard\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"playerId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"season\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"scarcity\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"serialNumber\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"clubId\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"clubId\",\"type\":\"uint16\"}],\"name\":\"getClub\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"country\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"city\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"yearFounded\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"playerId\",\"type\":\"uint256\"}],\"name\":\"getPlayer\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"yearOfBirth\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"monthOfBirth\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"dayOfBirth\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelayAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"migrateTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cardId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextContract\",\"outputs\":[{\"internalType\":\"contractINextContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nextContractAddress\",\"type\":\"address\"}],\"name\":\"setNextContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prefix\",\"type\":\"string\"}],\"name\":\"setTokenURIPrefix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sorareCards\",\"outputs\":[{\"internalType\":\"contractISorareCards\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"page\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rows\",\"type\":\"uint8\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SorareTokensFuncSigs maps the 4-byte function signature to its string representation.
var SorareTokensFuncSigs = map[string]string{
	"a217fddf": "DEFAULT_ADMIN_ROLE()",
	"d5391393": "MINTER_ROLE()",
	"983b2d56": "addMinter(address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"6c0360eb": "baseURI()",
	"309e11fc": "blockRelay()",
	"2850d20e": "createCardAndMintToken(uint256,uint16,uint8,uint16,bytes32,uint16,address)",
	"081812fc": "getApproved(uint256)",
	"9188d312": "getCard(uint256)",
	"4a39d2b8": "getClub(uint16)",
	"e55ae4e8": "getPlayer(uint256)",
	"32deedc6": "getRelayAddress()",
	"248a9ca3": "getRoleAdmin(bytes32)",
	"9010d07c": "getRoleMember(bytes32,uint256)",
	"ca15c873": "getRoleMemberCount(bytes32)",
	"2f2ff15d": "grantRole(bytes32,address)",
	"91d14854": "hasRole(bytes32,address)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"e5951e60": "migrateTokens(uint256[])",
	"a140ae23": "mintToken(uint256,address)",
	"06fdde03": "name()",
	"d8787fde": "nextContract()",
	"8da5cb5b": "owner()",
	"6352211e": "ownerOf(uint256)",
	"5f112c68": "renounceMinter(address)",
	"715018a6": "renounceOwnership()",
	"36568abe": "renounceRole(bytes32,address)",
	"cfbd4885": "revokeMinter(address)",
	"d547741f": "revokeRole(bytes32,address)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"53390100": "setNextContract(address)",
	"99e0dd7c": "setTokenURIPrefix(string)",
	"16b1b0c8": "sorareCards()",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"4f6ccce7": "tokenByIndex(uint256)",
	"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
	"c87b56dd": "tokenURI(uint256)",
	"4d9c5e58": "tokensOfOwner(address,uint8,uint8)",
	"18160ddd": "totalSupply()",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// SorareTokensBin is the compiled bytecode used for deploying new contracts.
var SorareTokensBin = "0x60806040523480156200001157600080fd5b5060405162003bf638038062003bf6833981810160405260408110156200003757600080fd5b5080516020918201516040805180820182526006815265536f7261726560d01b818601528151808301909252600382526229a7a960e91b94820194909452919290918260006200008f6001600160e01b03620002b416565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350620000ef6000336001600160e01b03620002d116565b604080516a4d494e5445525f524f4c4560a81b8152905190819003600b0190206200012490336001600160e01b03620002d116565b6001600160a01b03811662000180576040805162461bcd60e51b815260206004820181905260248201527f437573746f6d2072656c61792061646472657373206973207265717569726564604482015290519081900360640190fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055620001b26301ffc9a760e01b620002ea565b8151620001c790600990602085019062000512565b508051620001dd90600a90602084019062000512565b50620001f96380ac58cd60e01b6001600160e01b03620002ea16565b62000214635b5e139f60e01b6001600160e01b03620002ea16565b6200022f63780e9d6360e01b6001600160e01b03620002ea16565b50506001600160a01b0382166200028d576040805162461bcd60e51b815260206004820152601f60248201527f536f726172654361726473206164647265737320697320726571756972656400604482015290519081900360640190fd5b50600d80546001600160a01b0319166001600160a01b0392909216919091179055620005b4565b6000620002cb6200036f60201b62002bb91760201c565b90505b90565b620002e682826001600160e01b03620003a816565b5050565b6001600160e01b031980821614156200034a576040805162461bcd60e51b815260206004820152601c60248201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604482015290519081900360640190fd5b6001600160e01b0319166000908152600360205260409020805460ff19166001179055565b6002546000906001600160a01b031633146200038d575033620002ce565b620003a06001600160e01b036200042c16565b9050620002ce565b6000828152600160209081526040909120620003cf91839062002c1662000479821b17901c565b15620002e657620003e86001600160e01b03620002b416565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b600060606000368080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050503601516001600160a01b031692915050565b600062000499836001600160a01b0384166001600160e01b03620004a216565b90505b92915050565b6000620004b983836001600160e01b03620004fa16565b620004f1575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556200049c565b5060006200049c565b60009081526001919091016020526040902054151590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200055557805160ff191683800117855562000585565b8280016001018555821562000585579182015b828111156200058557825182559160200191906001019062000568565b506200059392915062000597565b5090565b620002ce91905b808211156200059357600081556001016200059e565b61363280620005c46000396000f3fe608060405234801561001057600080fd5b506004361061027f5760003560e01c806370a082311161015c578063a22cb465116100ce578063d547741f11610087578063d547741f14610b30578063d8787fde14610b5c578063e55ae4e814610b64578063e5951e6014610c22578063e985e9c514610c90578063f2fde38b14610cbe5761027f565b8063a22cb465146109d6578063b88d4fde14610a04578063c87b56dd14610ac8578063ca15c87314610ae5578063cfbd488514610b02578063d539139314610b285761027f565b806391d148541161012057806391d14854146108a457806395d89b41146108d0578063983b2d56146108d857806399e0dd7c146108fe578063a140ae23146109a2578063a217fddf146109ce5761027f565b806370a08231146107bc578063715018a6146107e25780638da5cb5b146107ea5780639010d07c146107f25780639188d312146108155761027f565b8063309e11fc116101f55780634d9c5e58116101b95780634d9c5e58146106a75780634f6ccce71461072e578063533901001461074b5780635f112c68146107715780636352211e146107975780636c0360eb146107b45761027f565b8063309e11fc146104c657806332deedc6146104ce57806336568abe146104d657806342842e0e146105025780634a39d2b8146105385761027f565b806318160ddd1161024757806318160ddd146103ab57806323b872dd146103c5578063248a9ca3146103fb5780632850d20e146104185780632f2ff15d1461046e5780632f745c591461049a5761027f565b806301ffc9a71461028457806306fdde03146102bf578063081812fc1461033c578063095ea7b31461037557806316b1b0c8146103a3575b600080fd5b6102ab6004803603602081101561029a57600080fd5b50356001600160e01b031916610ce4565b604080519115158252519081900360200190f35b6102c7610d07565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103015781810151838201526020016102e9565b50505050905090810190601f16801561032e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103596004803603602081101561035257600080fd5b5035610d9e565b604080516001600160a01b039092168252519081900360200190f35b6103a16004803603604081101561038b57600080fd5b506001600160a01b038135169060200135610e00565b005b610359610edb565b6103b3610eea565b60408051918252519081900360200190f35b6103a1600480360360608110156103db57600080fd5b506001600160a01b03813581169160208101359091169060400135610efb565b6103b36004803603602081101561041157600080fd5b5035610f52565b6103b3600480360360e081101561042e57600080fd5b50803590602081013561ffff90811691604081013560ff1691606082013581169160808101359160a0820135169060c001356001600160a01b0316610f67565b6103a16004803603604081101561048457600080fd5b50803590602001356001600160a01b03166110a0565b6103b3600480360360408110156104b057600080fd5b506001600160a01b038135169060200135611107565b6103a1611138565b6103596111a4565b6103a1600480360360408110156104ec57600080fd5b50803590602001356001600160a01b03166111b3565b6103a16004803603606081101561051857600080fd5b506001600160a01b03813581169160208101359091169060400135611214565b6105596004803603602081101561054e57600080fd5b503561ffff1661122f565b6040805161ffff83166060820152608080825286519082015285519091829160208084019284019160a08501918a019080838360005b838110156105a757818101518382015260200161058f565b50505050905090810190601f1680156105d45780820380516001836020036101000a031916815260200191505b50848103835287518152875160209182019189019080838360005b838110156106075781810151838201526020016105ef565b50505050905090810190601f1680156106345780820380516001836020036101000a031916815260200191505b50848103825286518152865160209182019188019080838360005b8381101561066757818101518382015260200161064f565b50505050905090810190601f1680156106945780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b6106de600480360360608110156106bd57600080fd5b506001600160a01b038135169060ff602082013581169160400135166114d7565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561071a578181015183820152602001610702565b505050509050019250505060405180910390f35b6103b36004803603602081101561074457600080fd5b503561161d565b6103a16004803603602081101561076157600080fd5b50356001600160a01b0316611639565b6103a16004803603602081101561078757600080fd5b50356001600160a01b0316611711565b610359600480360360208110156107ad57600080fd5b503561173e565b6102c761176c565b6103b3600480360360208110156107d257600080fd5b50356001600160a01b03166117cd565b6103a1611835565b6103596118d7565b6103596004803603604081101561080857600080fd5b50803590602001356118e6565b6108326004803603602081101561082b57600080fd5b5035611904565b604051808781526020018661ffff1661ffff1681526020018581526020018461ffff1661ffff168152602001806020018361ffff1661ffff168152602001828103825284818151815260200191508051906020019080838360008381101561066757818101518382015260200161064f565b6102ab600480360360408110156108ba57600080fd5b50803590602001356001600160a01b0316611a86565b6102c7611aa4565b6103a1600480360360208110156108ee57600080fd5b50356001600160a01b0316611b05565b6103a16004803603602081101561091457600080fd5b810190602081018135600160201b81111561092e57600080fd5b82018360208201111561094057600080fd5b803590602001918460018302840111600160201b8311171561096157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611b2f945050505050565b6103b3600480360360408110156109b857600080fd5b50803590602001356001600160a01b0316611b90565b6103b3611cd9565b6103a1600480360360408110156109ec57600080fd5b506001600160a01b0381351690602001351515611cde565b6103a160048036036080811015610a1a57600080fd5b6001600160a01b03823581169260208101359091169160408201359190810190608081016060820135600160201b811115610a5457600080fd5b820183602082011115610a6657600080fd5b803590602001918460018302840111600160201b83111715610a8757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611de3945050505050565b6102c760048036036020811015610ade57600080fd5b5035611e41565b6103b360048036036020811015610afb57600080fd5b50356120e8565b6103a160048036036020811015610b1857600080fd5b50356001600160a01b03166120ff565b6103b3612129565b6103a160048036036040811015610b4657600080fd5b50803590602001356001600160a01b031661214c565b6103596121a5565b610b8160048036036020811015610b7a57600080fd5b50356121b4565b60405180806020018561ffff1661ffff1681526020018460ff1660ff1681526020018360ff1660ff168152602001828103825286818151815260200191508051906020019080838360005b83811015610be4578181015183820152602001610bcc565b50505050905090810190601f168015610c115780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b6103a160048036036020811015610c3857600080fd5b810190602081018135600160201b811115610c5257600080fd5b820183602082011115610c6457600080fd5b803590602001918460208302840111600160201b83111715610c8557600080fd5b509092509050612312565b6102ab60048036036040811015610ca657600080fd5b506001600160a01b038135811691602001351661244d565b6103a160048036036020811015610cd457600080fd5b50356001600160a01b031661247b565b6001600160e01b0319811660009081526003602052604090205460ff165b919050565b60098054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610d935780601f10610d6857610100808354040283529160200191610d93565b820191906000526020600020905b815481529060010190602001808311610d7657829003601f168201915b505050505090505b90565b6000610da982612573565b610de45760405162461bcd60e51b815260040180806020018281038252602c8152602001806134d8602c913960400191505060405180910390fd5b506000908152600760205260409020546001600160a01b031690565b6000610e0b8261173e565b9050806001600160a01b0316836001600160a01b03161415610e5e5760405162461bcd60e51b815260040180806020018281038252602181526020018061357c6021913960400191505060405180910390fd5b806001600160a01b0316610e70612586565b6001600160a01b03161480610e915750610e9181610e8c612586565b61244d565b610ecc5760405162461bcd60e51b815260040180806020018281038252603881526020018061342b6038913960400191505060405180910390fd5b610ed68383612590565b505050565b600d546001600160a01b031681565b6000610ef660056125fe565b905090565b610f0c610f06612586565b82612609565b610f475760405162461bcd60e51b815260040180806020018281038252603181526020018061359d6031913960400191505060405180910390fd5b610ed68383836126ad565b60009081526001602052604090206002015490565b604080516a4d494e5445525f524f4c4560a81b8152905190819003600b019020600090610f9b90610f96612586565b611a86565b610fe5576040805162461bcd60e51b815260206004820152601660248201527529b2b73232b91034b9903737ba10309036b4b73a32b960511b604482015290519081900360640190fd5b600d5460408051630115238360e01b8152600481018b905261ffff808b16602483015260ff8a166044830152808916606483015260848201889052861660a482015290516000926001600160a01b03169163011523839160c480830192602092919082900301818787803b15801561105c57600080fd5b505af1158015611070573d6000803e3d6000fd5b505050506040513d602081101561108657600080fd5b50519050611094838261280b565b98975050505050505050565b6000828152600160205260409020600201546110be90610f96612586565b6110f95760405162461bcd60e51b815260040180806020018281038252602f815260200180613324602f913960400191505060405180910390fd5b6111038282612945565b5050565b6001600160a01b038216600090815260046020526040812061112f908363ffffffff6129b416565b90505b92915050565b611140612586565b6000546001600160a01b03908116911614611190576040805162461bcd60e51b81526020600482018190526024820152600080516020613504833981519152604482015290519081900360640190fd5b600280546001600160a01b03191630179055565b6002546001600160a01b031690565b6111bb612586565b6001600160a01b0316816001600160a01b03161461120a5760405162461bcd60e51b815260040180806020018281038252602f8152602001806135ce602f913960400191505060405180910390fd5b61110382826129c0565b610ed683838360405180602001604052806000815250611de3565b600d54604080516309473a5760e31b815261ffff841660048201529051606092839283926000926001600160a01b031691634a39d2b89160248083019286929190829003018186803b15801561128457600080fd5b505afa158015611298573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260808110156112c157600080fd5b8101908080516040519392919084600160201b8211156112e057600080fd5b9083019060208201858111156112f557600080fd5b8251600160201b81118282018810171561130e57600080fd5b82525081516020918201929091019080838360005b8381101561133b578181015183820152602001611323565b50505050905090810190601f1680156113685780820380516001836020036101000a031916815260200191505b5060405260200180516040519392919084600160201b82111561138a57600080fd5b90830190602082018581111561139f57600080fd5b8251600160201b8111828201881017156113b857600080fd5b82525081516020918201929091019080838360005b838110156113e55781810151838201526020016113cd565b50505050905090810190601f1680156114125780820380516001836020036101000a031916815260200191505b5060405260200180516040519392919084600160201b82111561143457600080fd5b90830190602082018581111561144957600080fd5b8251600160201b81118282018810171561146257600080fd5b82525081516020918201929091019080838360005b8381101561148f578181015183820152602001611477565b50505050905090810190601f1680156114bc5780820380516001836020036101000a031916815260200191505b5060405260200151949a939950919750929550909350505050565b60408051630dade9a760e21b81523060048201526001600160a01b038516602482015260ff808516604483015283166064820152905160609173__$6dc487fb79da1fe9b9c445d2980a791dce$__916336b7a69c91608480820192600092909190829003018186803b15801561154c57600080fd5b505af4158015611560573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561158957600080fd5b8101908080516040519392919084600160201b8211156115a857600080fd5b9083019060208201858111156115bd57600080fd5b82518660208202830111600160201b821117156115d957600080fd5b82525081516020918201928201910280838360005b838110156116065781810151838201526020016115ee565b5050505090500160405250505090505b9392505050565b60008061163160058463ffffffff612a2f16565b509392505050565b611641612586565b6000546001600160a01b03908116911614611691576040805162461bcd60e51b81526020600482018190526024820152600080516020613504833981519152604482015290519081900360640190fd5b600e546001600160a01b0316156116ef576040805162461bcd60e51b815260206004820152601860248201527f4e657874436f6e747261637420616c7265616479207365740000000000000000604482015290519081900360640190fd5b600e80546001600160a01b0319166001600160a01b0392909216919091179055565b604080516a4d494e5445525f524f4c4560a81b8152905190819003600b01902061173b90826111b3565b50565b60006111328260405180606001604052806029815260200161348d602991396005919063ffffffff612a4b16565b600c8054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610d935780601f10610d6857610100808354040283529160200191610d93565b60006001600160a01b0382166118145760405162461bcd60e51b815260040180806020018281038252602a815260200180613463602a913960400191505060405180910390fd5b6001600160a01b0382166000908152600460205260409020611132906125fe565b61183d612586565b6000546001600160a01b0390811691161461188d576040805162461bcd60e51b81526020600482018190526024820152600080516020613504833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b600082815260016020526040812061112f908363ffffffff6129b416565b60008060008060606000600d60009054906101000a90046001600160a01b03166001600160a01b0316639188d312886040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561196757600080fd5b505afa15801561197b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260c08110156119a457600080fd5b8151602083015160408085015160608601516080870180519351959794969295919491939282019284600160201b8211156119de57600080fd5b9083019060208201858111156119f357600080fd5b8251600160201b811182820188101715611a0c57600080fd5b82525081516020918201929091019080838360005b83811015611a39578181015183820152602001611a21565b50505050905090810190601f168015611a665780820380516001836020036101000a031916815260200191505b5060405260200151969e959d50939b509199509750929550909350505050565b600082815260016020526040812061112f908363ffffffff612a5816565b600a8054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610d935780601f10610d6857610100808354040283529160200191610d93565b604080516a4d494e5445525f524f4c4560a81b8152905190819003600b01902061173b90826110a0565b611b37612586565b6000546001600160a01b03908116911614611b87576040805162461bcd60e51b81526020600482018190526024820152600080516020613504833981519152604482015290519081900360640190fd5b61173b81612a6d565b604080516a4d494e5445525f524f4c4560a81b8152905190819003600b019020600090611bbf90610f96612586565b611c09576040805162461bcd60e51b815260206004820152601660248201527529b2b73232b91034b9903737ba10309036b4b73a32b960511b604482015290519081900360640190fd5b600d5460408051632571aa0760e11b81526004810186905290516001600160a01b0390921691634ae3540e91602480820192602092909190829003018186803b158015611c5557600080fd5b505afa158015611c69573d6000803e3d6000fd5b505050506040513d6020811015611c7f57600080fd5b5051611cc8576040805162461bcd60e51b815260206004820152601360248201527210d85c9908191bd95cc81b9bdd08195e1a5cdd606a1b604482015290519081900360640190fd5b611cd2828461280b565b5090919050565b600081565b611ce6612586565b6001600160a01b0316826001600160a01b03161415611d4c576040805162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015290519081900360640190fd5b8060086000611d59612586565b6001600160a01b03908116825260208083019390935260409182016000908120918716808252919093529120805460ff191692151592909217909155611d9d612586565b60408051841515815290516001600160a01b0392909216917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c319181900360200190a35050565b611df4611dee612586565b83612609565b611e2f5760405162461bcd60e51b815260040180806020018281038252603181526020018061359d6031913960400191505060405180910390fd5b611e3b84848484612a80565b50505050565b6060611e4c82612573565b611e875760405162461bcd60e51b815260040180806020018281038252602f81526020018061354d602f913960400191505060405180910390fd5b6000828152600b602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845260609392830182828015611f1c5780601f10611ef157610100808354040283529160200191611f1c565b820191906000526020600020905b815481529060010190602001808311611eff57829003601f168201915b5050600c5493945050505060026000196101006001841615020190911604611f45579050610d02565b80511561201657600c816040516020018083805460018160011615610100020316600290048015611fad5780601f10611f8b576101008083540402835291820191611fad565b820191906000526020600020905b815481529060010190602001808311611f99575b5050825160208401908083835b60208310611fd95780518252601f199092019160209182019101611fba565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052915050610d02565b600c61202184612ad2565b604051602001808380546001816001161561010002031660029004801561207f5780601f1061205d57610100808354040283529182019161207f565b820191906000526020600020905b81548152906001019060200180831161206b575b5050825160208401908083835b602083106120ab5780518252601f19909201916020918201910161208c565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052915050919050565b6000818152600160205260408120611132906125fe565b604080516a4d494e5445525f524f4c4560a81b8152905190819003600b01902061173b908261214c565b604080516a4d494e5445525f524f4c4560a81b8152905190819003600b01902081565b60008281526001602052604090206002015461216a90610f96612586565b61120a5760405162461bcd60e51b81526004018080602001828103825260308152602001806133fb6030913960400191505060405180910390fd5b600e546001600160a01b031681565b600d5460408051631cab5c9d60e31b8152600481018490529051606092600092839283926001600160a01b03169163e55ae4e89160248083019286929190829003018186803b15801561220657600080fd5b505afa15801561221a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052608081101561224357600080fd5b8101908080516040519392919084600160201b82111561226257600080fd5b90830190602082018581111561227757600080fd5b8251600160201b81118282018810171561229057600080fd5b82525081516020918201929091019080838360005b838110156122bd5781810151838201526020016122a5565b50505050905090810190601f1680156122ea5780820380516001836020036101000a031916815260200191505b506040908152602082015190820151606090920151949a909950909750929550919350505050565b600e546001600160a01b0316612367576040805162461bcd60e51b815260206004820152601560248201527413995e1d0818dbdb9d1c9858dd081b9bdd081cd95d605a1b604482015290519081900360640190fd5b60005b8181101561239e5761239661237d612586565b3085858581811061238a57fe5b90506020020135610efb565b60010161236a565b50600e546001600160a01b031663e884703f83836123ba612586565b604080516001600160e01b031960e087901b1681526001600160a01b038316602482015260048101918252604481018490528190606401856020860280828437600081840152601f19601f820116905080830192505050945050505050600060405180830381600087803b15801561243157600080fd5b505af1158015612445573d6000803e3d6000fd5b505050505050565b6001600160a01b03918216600090815260086020908152604080832093909416825291909152205460ff1690565b612483612586565b6000546001600160a01b039081169116146124d3576040805162461bcd60e51b81526020600482018190526024820152600080516020613504833981519152604482015290519081900360640190fd5b6001600160a01b0381166125185760405162461bcd60e51b81526004018080602001828103825260268152602001806133856026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b600061113260058363ffffffff612bad16565b6000610ef6612bb9565b600081815260076020526040902080546001600160a01b0319166001600160a01b03841690811790915581906125c58261173e565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600061113282612be4565b600061261482612573565b61264f5760405162461bcd60e51b815260040180806020018281038252602c8152602001806133cf602c913960400191505060405180910390fd5b600061265a8361173e565b9050806001600160a01b0316846001600160a01b031614806126955750836001600160a01b031661268a84610d9e565b6001600160a01b0316145b806126a557506126a5818561244d565b949350505050565b826001600160a01b03166126c08261173e565b6001600160a01b0316146127055760405162461bcd60e51b81526004018080602001828103825260298152602001806135246029913960400191505060405180910390fd5b6001600160a01b03821661274a5760405162461bcd60e51b81526004018080602001828103825260248152602001806133ab6024913960400191505060405180910390fd5b612755838383610ed6565b612760600082612590565b6001600160a01b0383166000908152600460205260409020612788908263ffffffff612be816565b506001600160a01b03821660009081526004602052604090206127b1908263ffffffff612bf416565b506127c46005828463ffffffff612c0016565b5080826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b6001600160a01b038216612866576040805162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f2061646472657373604482015290519081900360640190fd5b61286f81612573565b156128c1576040805162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015290519081900360640190fd5b6128cd60008383610ed6565b6001600160a01b03821660009081526004602052604090206128f5908263ffffffff612bf416565b506129086005828463ffffffff612c0016565b5060405181906001600160a01b038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6000828152600160205260409020612963908263ffffffff612c1616565b1561110357612970612586565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b600061112f8383612c2b565b60008281526001602052604090206129de908263ffffffff612c8f16565b15611103576129eb612586565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b6000808080612a3e8686612ca4565b9097909650945050505050565b60006126a5848484612d1f565b600061112f836001600160a01b038416612de9565b805161110390600c906020840190613269565b612a8b8484846126ad565b612a9784848484612e01565b611e3b5760405162461bcd60e51b81526004018080602001828103825260328152602001806133536032913960400191505060405180910390fd5b606081612af757506040805180820190915260018152600360fc1b6020820152610d02565b8160005b8115612b0f57600101600a82049150612afb565b60608167ffffffffffffffff81118015612b2857600080fd5b506040519080825280601f01601f191660200182016040528015612b53576020820181803683370190505b50859350905060001982015b8315612ba457600a840660300160f81b82828060019003935081518110612b8257fe5b60200101906001600160f81b031916908160001a905350600a84049350612b5f565b50949350505050565b600061112f8383612de9565b6002546000906001600160a01b03163314612bd5575033610d9b565b612bdd61303c565b9050610d9b565b5490565b600061112f8383613089565b600061112f838361314f565b60006126a584846001600160a01b038516613199565b600061112f836001600160a01b03841661314f565b81546000908210612c6d5760405162461bcd60e51b81526004018080602001828103825260228152602001806133026022913960400191505060405180910390fd5b826000018281548110612c7c57fe5b9060005260206000200154905092915050565b600061112f836001600160a01b038416613089565b815460009081908310612ce85760405162461bcd60e51b81526004018080602001828103825260228152602001806134b66022913960400191505060405180910390fd5b6000846000018481548110612cf957fe5b906000526020600020906002020190508060000154816001015492509250509250929050565b60008281526001840160205260408120548281612dba5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612d7f578181015183820152602001612d67565b50505050905090810190601f168015612dac5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50846000016001820381548110612dcd57fe5b9060005260206000209060020201600101549150509392505050565b60009081526001919091016020526040902054151590565b6000612e15846001600160a01b0316613230565b612e21575060016126a5565b600060606001600160a01b038616630a85bd0160e11b612e3f612586565b89888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015612eb8578181015183820152602001612ea0565b50505050905090810190601f168015612ee55780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310612f4d5780518252601f199092019160209182019101612f2e565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114612faf576040519150601f19603f3d011682016040523d82523d6000602084013e612fb4565b606091505b50915091508161300557805115612fce5780518082602001fd5b60405162461bcd60e51b81526004018080602001828103825260328152602001806133536032913960400191505060405180910390fd5b600081806020019051602081101561301c57600080fd5b50516001600160e01b031916630a85bd0160e11b1493506126a592505050565b600060606000368080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050503601516001600160a01b031692915050565b6000818152600183016020526040812054801561314557835460001980830191908101906000908790839081106130bc57fe5b90600052602060002001549050808760000184815481106130d957fe5b60009182526020808320909101929092558281526001898101909252604090209084019055865487908061310957fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050611132565b6000915050611132565b600061315b8383612de9565b61319157508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155611132565b506000611132565b6000828152600184016020526040812054806131fe575050604080518082018252838152602080820184815286546001818101895560008981528481209551600290930290950191825591519082015586548684528188019092529290912055611616565b8285600001600183038154811061321157fe5b9060005260206000209060020201600101819055506000915050611616565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708181148015906126a5575050151592915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106132aa57805160ff19168380011785556132d7565b828001600101855582156132d7579182015b828111156132d75782518255916020019190600101906132bc565b506132e39291506132e7565b5090565b610d9b91905b808211156132e357600081556001016132ed56fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e6473416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f206772616e744552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734552433732313a207472616e7366657220746f20746865207a65726f20616464726573734552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f207265766f6b654552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e456e756d657261626c654d61703a20696e646578206f7574206f6620626f756e64734552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65724552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732314d657461646174613a2055524920717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65724552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636520726f6c657320666f722073656c66a26469706673582212207912b05e0717370978fff8a1fba4936e646b8bf8392e1984868ec68f705e6d1f64736f6c63430006060033"

// DeploySorareTokens deploys a new Ethereum contract, binding an instance of SorareTokens to it.
func DeploySorareTokens(auth *bind.TransactOpts, backend bind.ContractBackend, sorareCardsAddress common.Address, relayAddress common.Address) (common.Address, *types.Transaction, *SorareTokens, error) {
	parsed, err := abi.JSON(strings.NewReader(SorareTokensABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	nFTClientAddr, _, _, _ := DeployNFTClient(auth, backend)
	SorareTokensBin = strings.Replace(SorareTokensBin, "__$6dc487fb79da1fe9b9c445d2980a791dce$__", nFTClientAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SorareTokensBin), backend, sorareCardsAddress, relayAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SorareTokens{SorareTokensCaller: SorareTokensCaller{contract: contract}, SorareTokensTransactor: SorareTokensTransactor{contract: contract}, SorareTokensFilterer: SorareTokensFilterer{contract: contract}}, nil
}

// SorareTokens is an auto generated Go binding around an Ethereum contract.
type SorareTokens struct {
	SorareTokensCaller     // Read-only binding to the contract
	SorareTokensTransactor // Write-only binding to the contract
	SorareTokensFilterer   // Log filterer for contract events
}

// SorareTokensCaller is an auto generated read-only Go binding around an Ethereum contract.
type SorareTokensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SorareTokensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SorareTokensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SorareTokensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SorareTokensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SorareTokensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SorareTokensSession struct {
	Contract     *SorareTokens     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SorareTokensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SorareTokensCallerSession struct {
	Contract *SorareTokensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SorareTokensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SorareTokensTransactorSession struct {
	Contract     *SorareTokensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SorareTokensRaw is an auto generated low-level Go binding around an Ethereum contract.
type SorareTokensRaw struct {
	Contract *SorareTokens // Generic contract binding to access the raw methods on
}

// SorareTokensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SorareTokensCallerRaw struct {
	Contract *SorareTokensCaller // Generic read-only contract binding to access the raw methods on
}

// SorareTokensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SorareTokensTransactorRaw struct {
	Contract *SorareTokensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSorareTokens creates a new instance of SorareTokens, bound to a specific deployed contract.
func NewSorareTokens(address common.Address, backend bind.ContractBackend) (*SorareTokens, error) {
	contract, err := bindSorareTokens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SorareTokens{SorareTokensCaller: SorareTokensCaller{contract: contract}, SorareTokensTransactor: SorareTokensTransactor{contract: contract}, SorareTokensFilterer: SorareTokensFilterer{contract: contract}}, nil
}

// NewSorareTokensCaller creates a new read-only instance of SorareTokens, bound to a specific deployed contract.
func NewSorareTokensCaller(address common.Address, caller bind.ContractCaller) (*SorareTokensCaller, error) {
	contract, err := bindSorareTokens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SorareTokensCaller{contract: contract}, nil
}

// NewSorareTokensTransactor creates a new write-only instance of SorareTokens, bound to a specific deployed contract.
func NewSorareTokensTransactor(address common.Address, transactor bind.ContractTransactor) (*SorareTokensTransactor, error) {
	contract, err := bindSorareTokens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SorareTokensTransactor{contract: contract}, nil
}

// NewSorareTokensFilterer creates a new log filterer instance of SorareTokens, bound to a specific deployed contract.
func NewSorareTokensFilterer(address common.Address, filterer bind.ContractFilterer) (*SorareTokensFilterer, error) {
	contract, err := bindSorareTokens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SorareTokensFilterer{contract: contract}, nil
}

// bindSorareTokens binds a generic wrapper to an already deployed contract.
func bindSorareTokens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SorareTokensABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SorareTokens *SorareTokensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SorareTokens.Contract.SorareTokensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SorareTokens *SorareTokensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SorareTokens.Contract.SorareTokensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SorareTokens *SorareTokensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SorareTokens.Contract.SorareTokensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SorareTokens *SorareTokensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SorareTokens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SorareTokens *SorareTokensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SorareTokens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SorareTokens *SorareTokensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SorareTokens.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SorareTokens *SorareTokensCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SorareTokens *SorareTokensSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SorareTokens.Contract.DEFAULTADMINROLE(&_SorareTokens.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SorareTokens *SorareTokensCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SorareTokens.Contract.DEFAULTADMINROLE(&_SorareTokens.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_SorareTokens *SorareTokensCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_SorareTokens *SorareTokensSession) MINTERROLE() ([32]byte, error) {
	return _SorareTokens.Contract.MINTERROLE(&_SorareTokens.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_SorareTokens *SorareTokensCallerSession) MINTERROLE() ([32]byte, error) {
	return _SorareTokens.Contract.MINTERROLE(&_SorareTokens.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SorareTokens *SorareTokensCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SorareTokens *SorareTokensSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SorareTokens.Contract.BalanceOf(&_SorareTokens.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SorareTokens *SorareTokensCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SorareTokens.Contract.BalanceOf(&_SorareTokens.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_SorareTokens *SorareTokensCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_SorareTokens *SorareTokensSession) BaseURI() (string, error) {
	return _SorareTokens.Contract.BaseURI(&_SorareTokens.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_SorareTokens *SorareTokensCallerSession) BaseURI() (string, error) {
	return _SorareTokens.Contract.BaseURI(&_SorareTokens.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SorareTokens *SorareTokensCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SorareTokens *SorareTokensSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SorareTokens.Contract.GetApproved(&_SorareTokens.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SorareTokens *SorareTokensCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SorareTokens.Contract.GetApproved(&_SorareTokens.CallOpts, tokenId)
}

// GetCard is a free data retrieval call binding the contract method 0x9188d312.
//
// Solidity: function getCard(uint256 tokenId) view returns(uint256 playerId, uint16 season, uint256 scarcity, uint16 serialNumber, bytes metadata, uint16 clubId)
func (_SorareTokens *SorareTokensCaller) GetCard(opts *bind.CallOpts, tokenId *big.Int) (struct {
	PlayerId     *big.Int
	Season       uint16
	Scarcity     *big.Int
	SerialNumber uint16
	Metadata     []byte
	ClubId       uint16
}, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "getCard", tokenId)

	outstruct := new(struct {
		PlayerId     *big.Int
		Season       uint16
		Scarcity     *big.Int
		SerialNumber uint16
		Metadata     []byte
		ClubId       uint16
	})

	outstruct.PlayerId = out[0].(*big.Int)
	outstruct.Season = out[1].(uint16)
	outstruct.Scarcity = out[2].(*big.Int)
	outstruct.SerialNumber = out[3].(uint16)
	outstruct.Metadata = out[4].([]byte)
	outstruct.ClubId = out[5].(uint16)

	return *outstruct, err

}

// GetCard is a free data retrieval call binding the contract method 0x9188d312.
//
// Solidity: function getCard(uint256 tokenId) view returns(uint256 playerId, uint16 season, uint256 scarcity, uint16 serialNumber, bytes metadata, uint16 clubId)
func (_SorareTokens *SorareTokensSession) GetCard(tokenId *big.Int) (struct {
	PlayerId     *big.Int
	Season       uint16
	Scarcity     *big.Int
	SerialNumber uint16
	Metadata     []byte
	ClubId       uint16
}, error) {
	return _SorareTokens.Contract.GetCard(&_SorareTokens.CallOpts, tokenId)
}

// GetCard is a free data retrieval call binding the contract method 0x9188d312.
//
// Solidity: function getCard(uint256 tokenId) view returns(uint256 playerId, uint16 season, uint256 scarcity, uint16 serialNumber, bytes metadata, uint16 clubId)
func (_SorareTokens *SorareTokensCallerSession) GetCard(tokenId *big.Int) (struct {
	PlayerId     *big.Int
	Season       uint16
	Scarcity     *big.Int
	SerialNumber uint16
	Metadata     []byte
	ClubId       uint16
}, error) {
	return _SorareTokens.Contract.GetCard(&_SorareTokens.CallOpts, tokenId)
}

// GetClub is a free data retrieval call binding the contract method 0x4a39d2b8.
//
// Solidity: function getClub(uint16 clubId) view returns(string name, string country, string city, uint16 yearFounded)
func (_SorareTokens *SorareTokensCaller) GetClub(opts *bind.CallOpts, clubId uint16) (struct {
	Name        string
	Country     string
	City        string
	YearFounded uint16
}, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "getClub", clubId)

	outstruct := new(struct {
		Name        string
		Country     string
		City        string
		YearFounded uint16
	})

	outstruct.Name = out[0].(string)
	outstruct.Country = out[1].(string)
	outstruct.City = out[2].(string)
	outstruct.YearFounded = out[3].(uint16)

	return *outstruct, err

}

// GetClub is a free data retrieval call binding the contract method 0x4a39d2b8.
//
// Solidity: function getClub(uint16 clubId) view returns(string name, string country, string city, uint16 yearFounded)
func (_SorareTokens *SorareTokensSession) GetClub(clubId uint16) (struct {
	Name        string
	Country     string
	City        string
	YearFounded uint16
}, error) {
	return _SorareTokens.Contract.GetClub(&_SorareTokens.CallOpts, clubId)
}

// GetClub is a free data retrieval call binding the contract method 0x4a39d2b8.
//
// Solidity: function getClub(uint16 clubId) view returns(string name, string country, string city, uint16 yearFounded)
func (_SorareTokens *SorareTokensCallerSession) GetClub(clubId uint16) (struct {
	Name        string
	Country     string
	City        string
	YearFounded uint16
}, error) {
	return _SorareTokens.Contract.GetClub(&_SorareTokens.CallOpts, clubId)
}

// GetPlayer is a free data retrieval call binding the contract method 0xe55ae4e8.
//
// Solidity: function getPlayer(uint256 playerId) view returns(string name, uint16 yearOfBirth, uint8 monthOfBirth, uint8 dayOfBirth)
func (_SorareTokens *SorareTokensCaller) GetPlayer(opts *bind.CallOpts, playerId *big.Int) (struct {
	Name         string
	YearOfBirth  uint16
	MonthOfBirth uint8
	DayOfBirth   uint8
}, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "getPlayer", playerId)

	outstruct := new(struct {
		Name         string
		YearOfBirth  uint16
		MonthOfBirth uint8
		DayOfBirth   uint8
	})

	outstruct.Name = out[0].(string)
	outstruct.YearOfBirth = out[1].(uint16)
	outstruct.MonthOfBirth = out[2].(uint8)
	outstruct.DayOfBirth = out[3].(uint8)

	return *outstruct, err

}

// GetPlayer is a free data retrieval call binding the contract method 0xe55ae4e8.
//
// Solidity: function getPlayer(uint256 playerId) view returns(string name, uint16 yearOfBirth, uint8 monthOfBirth, uint8 dayOfBirth)
func (_SorareTokens *SorareTokensSession) GetPlayer(playerId *big.Int) (struct {
	Name         string
	YearOfBirth  uint16
	MonthOfBirth uint8
	DayOfBirth   uint8
}, error) {
	return _SorareTokens.Contract.GetPlayer(&_SorareTokens.CallOpts, playerId)
}

// GetPlayer is a free data retrieval call binding the contract method 0xe55ae4e8.
//
// Solidity: function getPlayer(uint256 playerId) view returns(string name, uint16 yearOfBirth, uint8 monthOfBirth, uint8 dayOfBirth)
func (_SorareTokens *SorareTokensCallerSession) GetPlayer(playerId *big.Int) (struct {
	Name         string
	YearOfBirth  uint16
	MonthOfBirth uint8
	DayOfBirth   uint8
}, error) {
	return _SorareTokens.Contract.GetPlayer(&_SorareTokens.CallOpts, playerId)
}

// GetRelayAddress is a free data retrieval call binding the contract method 0x32deedc6.
//
// Solidity: function getRelayAddress() view returns(address)
func (_SorareTokens *SorareTokensCaller) GetRelayAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "getRelayAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRelayAddress is a free data retrieval call binding the contract method 0x32deedc6.
//
// Solidity: function getRelayAddress() view returns(address)
func (_SorareTokens *SorareTokensSession) GetRelayAddress() (common.Address, error) {
	return _SorareTokens.Contract.GetRelayAddress(&_SorareTokens.CallOpts)
}

// GetRelayAddress is a free data retrieval call binding the contract method 0x32deedc6.
//
// Solidity: function getRelayAddress() view returns(address)
func (_SorareTokens *SorareTokensCallerSession) GetRelayAddress() (common.Address, error) {
	return _SorareTokens.Contract.GetRelayAddress(&_SorareTokens.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SorareTokens *SorareTokensCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SorareTokens *SorareTokensSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SorareTokens.Contract.GetRoleAdmin(&_SorareTokens.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SorareTokens *SorareTokensCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SorareTokens.Contract.GetRoleAdmin(&_SorareTokens.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_SorareTokens *SorareTokensCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_SorareTokens *SorareTokensSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _SorareTokens.Contract.GetRoleMember(&_SorareTokens.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_SorareTokens *SorareTokensCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _SorareTokens.Contract.GetRoleMember(&_SorareTokens.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_SorareTokens *SorareTokensCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_SorareTokens *SorareTokensSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _SorareTokens.Contract.GetRoleMemberCount(&_SorareTokens.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_SorareTokens *SorareTokensCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _SorareTokens.Contract.GetRoleMemberCount(&_SorareTokens.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SorareTokens *SorareTokensCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SorareTokens *SorareTokensSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SorareTokens.Contract.HasRole(&_SorareTokens.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SorareTokens *SorareTokensCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SorareTokens.Contract.HasRole(&_SorareTokens.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SorareTokens *SorareTokensCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SorareTokens *SorareTokensSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SorareTokens.Contract.IsApprovedForAll(&_SorareTokens.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SorareTokens *SorareTokensCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SorareTokens.Contract.IsApprovedForAll(&_SorareTokens.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SorareTokens *SorareTokensCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SorareTokens *SorareTokensSession) Name() (string, error) {
	return _SorareTokens.Contract.Name(&_SorareTokens.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SorareTokens *SorareTokensCallerSession) Name() (string, error) {
	return _SorareTokens.Contract.Name(&_SorareTokens.CallOpts)
}

// NextContract is a free data retrieval call binding the contract method 0xd8787fde.
//
// Solidity: function nextContract() view returns(address)
func (_SorareTokens *SorareTokensCaller) NextContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "nextContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NextContract is a free data retrieval call binding the contract method 0xd8787fde.
//
// Solidity: function nextContract() view returns(address)
func (_SorareTokens *SorareTokensSession) NextContract() (common.Address, error) {
	return _SorareTokens.Contract.NextContract(&_SorareTokens.CallOpts)
}

// NextContract is a free data retrieval call binding the contract method 0xd8787fde.
//
// Solidity: function nextContract() view returns(address)
func (_SorareTokens *SorareTokensCallerSession) NextContract() (common.Address, error) {
	return _SorareTokens.Contract.NextContract(&_SorareTokens.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SorareTokens *SorareTokensCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SorareTokens *SorareTokensSession) Owner() (common.Address, error) {
	return _SorareTokens.Contract.Owner(&_SorareTokens.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SorareTokens *SorareTokensCallerSession) Owner() (common.Address, error) {
	return _SorareTokens.Contract.Owner(&_SorareTokens.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SorareTokens *SorareTokensCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SorareTokens *SorareTokensSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SorareTokens.Contract.OwnerOf(&_SorareTokens.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SorareTokens *SorareTokensCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SorareTokens.Contract.OwnerOf(&_SorareTokens.CallOpts, tokenId)
}

// SorareCards is a free data retrieval call binding the contract method 0x16b1b0c8.
//
// Solidity: function sorareCards() view returns(address)
func (_SorareTokens *SorareTokensCaller) SorareCards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "sorareCards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SorareCards is a free data retrieval call binding the contract method 0x16b1b0c8.
//
// Solidity: function sorareCards() view returns(address)
func (_SorareTokens *SorareTokensSession) SorareCards() (common.Address, error) {
	return _SorareTokens.Contract.SorareCards(&_SorareTokens.CallOpts)
}

// SorareCards is a free data retrieval call binding the contract method 0x16b1b0c8.
//
// Solidity: function sorareCards() view returns(address)
func (_SorareTokens *SorareTokensCallerSession) SorareCards() (common.Address, error) {
	return _SorareTokens.Contract.SorareCards(&_SorareTokens.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SorareTokens *SorareTokensCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SorareTokens *SorareTokensSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SorareTokens.Contract.SupportsInterface(&_SorareTokens.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SorareTokens *SorareTokensCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SorareTokens.Contract.SupportsInterface(&_SorareTokens.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SorareTokens *SorareTokensCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SorareTokens *SorareTokensSession) Symbol() (string, error) {
	return _SorareTokens.Contract.Symbol(&_SorareTokens.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SorareTokens *SorareTokensCallerSession) Symbol() (string, error) {
	return _SorareTokens.Contract.Symbol(&_SorareTokens.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_SorareTokens *SorareTokensCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_SorareTokens *SorareTokensSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _SorareTokens.Contract.TokenByIndex(&_SorareTokens.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_SorareTokens *SorareTokensCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _SorareTokens.Contract.TokenByIndex(&_SorareTokens.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_SorareTokens *SorareTokensCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_SorareTokens *SorareTokensSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _SorareTokens.Contract.TokenOfOwnerByIndex(&_SorareTokens.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_SorareTokens *SorareTokensCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _SorareTokens.Contract.TokenOfOwnerByIndex(&_SorareTokens.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SorareTokens *SorareTokensCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SorareTokens *SorareTokensSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SorareTokens.Contract.TokenURI(&_SorareTokens.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SorareTokens *SorareTokensCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SorareTokens.Contract.TokenURI(&_SorareTokens.CallOpts, tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x4d9c5e58.
//
// Solidity: function tokensOfOwner(address owner, uint8 page, uint8 rows) view returns(uint256[])
func (_SorareTokens *SorareTokensCaller) TokensOfOwner(opts *bind.CallOpts, owner common.Address, page uint8, rows uint8) ([]*big.Int, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "tokensOfOwner", owner, page, rows)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x4d9c5e58.
//
// Solidity: function tokensOfOwner(address owner, uint8 page, uint8 rows) view returns(uint256[])
func (_SorareTokens *SorareTokensSession) TokensOfOwner(owner common.Address, page uint8, rows uint8) ([]*big.Int, error) {
	return _SorareTokens.Contract.TokensOfOwner(&_SorareTokens.CallOpts, owner, page, rows)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x4d9c5e58.
//
// Solidity: function tokensOfOwner(address owner, uint8 page, uint8 rows) view returns(uint256[])
func (_SorareTokens *SorareTokensCallerSession) TokensOfOwner(owner common.Address, page uint8, rows uint8) ([]*big.Int, error) {
	return _SorareTokens.Contract.TokensOfOwner(&_SorareTokens.CallOpts, owner, page, rows)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SorareTokens *SorareTokensCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SorareTokens.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SorareTokens *SorareTokensSession) TotalSupply() (*big.Int, error) {
	return _SorareTokens.Contract.TotalSupply(&_SorareTokens.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SorareTokens *SorareTokensCallerSession) TotalSupply() (*big.Int, error) {
	return _SorareTokens.Contract.TotalSupply(&_SorareTokens.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_SorareTokens *SorareTokensTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_SorareTokens *SorareTokensSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.AddMinter(&_SorareTokens.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_SorareTokens *SorareTokensTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.AddMinter(&_SorareTokens.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SorareTokens *SorareTokensTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SorareTokens *SorareTokensSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SorareTokens.Contract.Approve(&_SorareTokens.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SorareTokens *SorareTokensTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SorareTokens.Contract.Approve(&_SorareTokens.TransactOpts, to, tokenId)
}

// BlockRelay is a paid mutator transaction binding the contract method 0x309e11fc.
//
// Solidity: function blockRelay() returns()
func (_SorareTokens *SorareTokensTransactor) BlockRelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "blockRelay")
}

// BlockRelay is a paid mutator transaction binding the contract method 0x309e11fc.
//
// Solidity: function blockRelay() returns()
func (_SorareTokens *SorareTokensSession) BlockRelay() (*types.Transaction, error) {
	return _SorareTokens.Contract.BlockRelay(&_SorareTokens.TransactOpts)
}

// BlockRelay is a paid mutator transaction binding the contract method 0x309e11fc.
//
// Solidity: function blockRelay() returns()
func (_SorareTokens *SorareTokensTransactorSession) BlockRelay() (*types.Transaction, error) {
	return _SorareTokens.Contract.BlockRelay(&_SorareTokens.TransactOpts)
}

// CreateCardAndMintToken is a paid mutator transaction binding the contract method 0x2850d20e.
//
// Solidity: function createCardAndMintToken(uint256 playerId, uint16 season, uint8 scarcity, uint16 serialNumber, bytes32 metadata, uint16 clubId, address to) returns(uint256)
func (_SorareTokens *SorareTokensTransactor) CreateCardAndMintToken(opts *bind.TransactOpts, playerId *big.Int, season uint16, scarcity uint8, serialNumber uint16, metadata [32]byte, clubId uint16, to common.Address) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "createCardAndMintToken", playerId, season, scarcity, serialNumber, metadata, clubId, to)
}

// CreateCardAndMintToken is a paid mutator transaction binding the contract method 0x2850d20e.
//
// Solidity: function createCardAndMintToken(uint256 playerId, uint16 season, uint8 scarcity, uint16 serialNumber, bytes32 metadata, uint16 clubId, address to) returns(uint256)
func (_SorareTokens *SorareTokensSession) CreateCardAndMintToken(playerId *big.Int, season uint16, scarcity uint8, serialNumber uint16, metadata [32]byte, clubId uint16, to common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.CreateCardAndMintToken(&_SorareTokens.TransactOpts, playerId, season, scarcity, serialNumber, metadata, clubId, to)
}

// CreateCardAndMintToken is a paid mutator transaction binding the contract method 0x2850d20e.
//
// Solidity: function createCardAndMintToken(uint256 playerId, uint16 season, uint8 scarcity, uint16 serialNumber, bytes32 metadata, uint16 clubId, address to) returns(uint256)
func (_SorareTokens *SorareTokensTransactorSession) CreateCardAndMintToken(playerId *big.Int, season uint16, scarcity uint8, serialNumber uint16, metadata [32]byte, clubId uint16, to common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.CreateCardAndMintToken(&_SorareTokens.TransactOpts, playerId, season, scarcity, serialNumber, metadata, clubId, to)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SorareTokens *SorareTokensTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SorareTokens *SorareTokensSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.GrantRole(&_SorareTokens.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SorareTokens *SorareTokensTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.GrantRole(&_SorareTokens.TransactOpts, role, account)
}

// MigrateTokens is a paid mutator transaction binding the contract method 0xe5951e60.
//
// Solidity: function migrateTokens(uint256[] tokenIds) returns()
func (_SorareTokens *SorareTokensTransactor) MigrateTokens(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "migrateTokens", tokenIds)
}

// MigrateTokens is a paid mutator transaction binding the contract method 0xe5951e60.
//
// Solidity: function migrateTokens(uint256[] tokenIds) returns()
func (_SorareTokens *SorareTokensSession) MigrateTokens(tokenIds []*big.Int) (*types.Transaction, error) {
	return _SorareTokens.Contract.MigrateTokens(&_SorareTokens.TransactOpts, tokenIds)
}

// MigrateTokens is a paid mutator transaction binding the contract method 0xe5951e60.
//
// Solidity: function migrateTokens(uint256[] tokenIds) returns()
func (_SorareTokens *SorareTokensTransactorSession) MigrateTokens(tokenIds []*big.Int) (*types.Transaction, error) {
	return _SorareTokens.Contract.MigrateTokens(&_SorareTokens.TransactOpts, tokenIds)
}

// MintToken is a paid mutator transaction binding the contract method 0xa140ae23.
//
// Solidity: function mintToken(uint256 cardId, address to) returns(uint256)
func (_SorareTokens *SorareTokensTransactor) MintToken(opts *bind.TransactOpts, cardId *big.Int, to common.Address) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "mintToken", cardId, to)
}

// MintToken is a paid mutator transaction binding the contract method 0xa140ae23.
//
// Solidity: function mintToken(uint256 cardId, address to) returns(uint256)
func (_SorareTokens *SorareTokensSession) MintToken(cardId *big.Int, to common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.MintToken(&_SorareTokens.TransactOpts, cardId, to)
}

// MintToken is a paid mutator transaction binding the contract method 0xa140ae23.
//
// Solidity: function mintToken(uint256 cardId, address to) returns(uint256)
func (_SorareTokens *SorareTokensTransactorSession) MintToken(cardId *big.Int, to common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.MintToken(&_SorareTokens.TransactOpts, cardId, to)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x5f112c68.
//
// Solidity: function renounceMinter(address account) returns()
func (_SorareTokens *SorareTokensTransactor) RenounceMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "renounceMinter", account)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x5f112c68.
//
// Solidity: function renounceMinter(address account) returns()
func (_SorareTokens *SorareTokensSession) RenounceMinter(account common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.RenounceMinter(&_SorareTokens.TransactOpts, account)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x5f112c68.
//
// Solidity: function renounceMinter(address account) returns()
func (_SorareTokens *SorareTokensTransactorSession) RenounceMinter(account common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.RenounceMinter(&_SorareTokens.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SorareTokens *SorareTokensTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SorareTokens *SorareTokensSession) RenounceOwnership() (*types.Transaction, error) {
	return _SorareTokens.Contract.RenounceOwnership(&_SorareTokens.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SorareTokens *SorareTokensTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SorareTokens.Contract.RenounceOwnership(&_SorareTokens.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SorareTokens *SorareTokensTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SorareTokens *SorareTokensSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.RenounceRole(&_SorareTokens.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SorareTokens *SorareTokensTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.RenounceRole(&_SorareTokens.TransactOpts, role, account)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address account) returns()
func (_SorareTokens *SorareTokensTransactor) RevokeMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "revokeMinter", account)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address account) returns()
func (_SorareTokens *SorareTokensSession) RevokeMinter(account common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.RevokeMinter(&_SorareTokens.TransactOpts, account)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address account) returns()
func (_SorareTokens *SorareTokensTransactorSession) RevokeMinter(account common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.RevokeMinter(&_SorareTokens.TransactOpts, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SorareTokens *SorareTokensTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SorareTokens *SorareTokensSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.RevokeRole(&_SorareTokens.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SorareTokens *SorareTokensTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.RevokeRole(&_SorareTokens.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SorareTokens *SorareTokensTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SorareTokens *SorareTokensSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SorareTokens.Contract.SafeTransferFrom(&_SorareTokens.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SorareTokens *SorareTokensTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SorareTokens.Contract.SafeTransferFrom(&_SorareTokens.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_SorareTokens *SorareTokensTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_SorareTokens *SorareTokensSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SorareTokens.Contract.SafeTransferFrom0(&_SorareTokens.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_SorareTokens *SorareTokensTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SorareTokens.Contract.SafeTransferFrom0(&_SorareTokens.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SorareTokens *SorareTokensTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SorareTokens *SorareTokensSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SorareTokens.Contract.SetApprovalForAll(&_SorareTokens.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SorareTokens *SorareTokensTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SorareTokens.Contract.SetApprovalForAll(&_SorareTokens.TransactOpts, operator, approved)
}

// SetNextContract is a paid mutator transaction binding the contract method 0x53390100.
//
// Solidity: function setNextContract(address nextContractAddress) returns()
func (_SorareTokens *SorareTokensTransactor) SetNextContract(opts *bind.TransactOpts, nextContractAddress common.Address) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "setNextContract", nextContractAddress)
}

// SetNextContract is a paid mutator transaction binding the contract method 0x53390100.
//
// Solidity: function setNextContract(address nextContractAddress) returns()
func (_SorareTokens *SorareTokensSession) SetNextContract(nextContractAddress common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.SetNextContract(&_SorareTokens.TransactOpts, nextContractAddress)
}

// SetNextContract is a paid mutator transaction binding the contract method 0x53390100.
//
// Solidity: function setNextContract(address nextContractAddress) returns()
func (_SorareTokens *SorareTokensTransactorSession) SetNextContract(nextContractAddress common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.SetNextContract(&_SorareTokens.TransactOpts, nextContractAddress)
}

// SetTokenURIPrefix is a paid mutator transaction binding the contract method 0x99e0dd7c.
//
// Solidity: function setTokenURIPrefix(string prefix) returns()
func (_SorareTokens *SorareTokensTransactor) SetTokenURIPrefix(opts *bind.TransactOpts, prefix string) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "setTokenURIPrefix", prefix)
}

// SetTokenURIPrefix is a paid mutator transaction binding the contract method 0x99e0dd7c.
//
// Solidity: function setTokenURIPrefix(string prefix) returns()
func (_SorareTokens *SorareTokensSession) SetTokenURIPrefix(prefix string) (*types.Transaction, error) {
	return _SorareTokens.Contract.SetTokenURIPrefix(&_SorareTokens.TransactOpts, prefix)
}

// SetTokenURIPrefix is a paid mutator transaction binding the contract method 0x99e0dd7c.
//
// Solidity: function setTokenURIPrefix(string prefix) returns()
func (_SorareTokens *SorareTokensTransactorSession) SetTokenURIPrefix(prefix string) (*types.Transaction, error) {
	return _SorareTokens.Contract.SetTokenURIPrefix(&_SorareTokens.TransactOpts, prefix)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SorareTokens *SorareTokensTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SorareTokens *SorareTokensSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SorareTokens.Contract.TransferFrom(&_SorareTokens.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SorareTokens *SorareTokensTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SorareTokens.Contract.TransferFrom(&_SorareTokens.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SorareTokens *SorareTokensTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SorareTokens.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SorareTokens *SorareTokensSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.TransferOwnership(&_SorareTokens.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SorareTokens *SorareTokensTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SorareTokens.Contract.TransferOwnership(&_SorareTokens.TransactOpts, newOwner)
}

// SorareTokensApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SorareTokens contract.
type SorareTokensApprovalIterator struct {
	Event *SorareTokensApproval // Event containing the contract specifics and raw log

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
func (it *SorareTokensApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SorareTokensApproval)
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
		it.Event = new(SorareTokensApproval)
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
func (it *SorareTokensApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SorareTokensApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SorareTokensApproval represents a Approval event raised by the SorareTokens contract.
type SorareTokensApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SorareTokens *SorareTokensFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SorareTokensApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SorareTokens.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SorareTokensApprovalIterator{contract: _SorareTokens.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SorareTokens *SorareTokensFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SorareTokensApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SorareTokens.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SorareTokensApproval)
				if err := _SorareTokens.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SorareTokens *SorareTokensFilterer) ParseApproval(log types.Log) (*SorareTokensApproval, error) {
	event := new(SorareTokensApproval)
	if err := _SorareTokens.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SorareTokensApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SorareTokens contract.
type SorareTokensApprovalForAllIterator struct {
	Event *SorareTokensApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SorareTokensApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SorareTokensApprovalForAll)
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
		it.Event = new(SorareTokensApprovalForAll)
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
func (it *SorareTokensApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SorareTokensApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SorareTokensApprovalForAll represents a ApprovalForAll event raised by the SorareTokens contract.
type SorareTokensApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SorareTokens *SorareTokensFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SorareTokensApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SorareTokens.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SorareTokensApprovalForAllIterator{contract: _SorareTokens.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SorareTokens *SorareTokensFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SorareTokensApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SorareTokens.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SorareTokensApprovalForAll)
				if err := _SorareTokens.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SorareTokens *SorareTokensFilterer) ParseApprovalForAll(log types.Log) (*SorareTokensApprovalForAll, error) {
	event := new(SorareTokensApprovalForAll)
	if err := _SorareTokens.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SorareTokensOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SorareTokens contract.
type SorareTokensOwnershipTransferredIterator struct {
	Event *SorareTokensOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SorareTokensOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SorareTokensOwnershipTransferred)
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
		it.Event = new(SorareTokensOwnershipTransferred)
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
func (it *SorareTokensOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SorareTokensOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SorareTokensOwnershipTransferred represents a OwnershipTransferred event raised by the SorareTokens contract.
type SorareTokensOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SorareTokens *SorareTokensFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SorareTokensOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SorareTokens.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SorareTokensOwnershipTransferredIterator{contract: _SorareTokens.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SorareTokens *SorareTokensFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SorareTokensOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SorareTokens.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SorareTokensOwnershipTransferred)
				if err := _SorareTokens.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SorareTokens *SorareTokensFilterer) ParseOwnershipTransferred(log types.Log) (*SorareTokensOwnershipTransferred, error) {
	event := new(SorareTokensOwnershipTransferred)
	if err := _SorareTokens.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SorareTokensRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SorareTokens contract.
type SorareTokensRoleGrantedIterator struct {
	Event *SorareTokensRoleGranted // Event containing the contract specifics and raw log

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
func (it *SorareTokensRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SorareTokensRoleGranted)
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
		it.Event = new(SorareTokensRoleGranted)
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
func (it *SorareTokensRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SorareTokensRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SorareTokensRoleGranted represents a RoleGranted event raised by the SorareTokens contract.
type SorareTokensRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SorareTokens *SorareTokensFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SorareTokensRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SorareTokens.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SorareTokensRoleGrantedIterator{contract: _SorareTokens.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SorareTokens *SorareTokensFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SorareTokensRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SorareTokens.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SorareTokensRoleGranted)
				if err := _SorareTokens.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SorareTokens *SorareTokensFilterer) ParseRoleGranted(log types.Log) (*SorareTokensRoleGranted, error) {
	event := new(SorareTokensRoleGranted)
	if err := _SorareTokens.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SorareTokensRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SorareTokens contract.
type SorareTokensRoleRevokedIterator struct {
	Event *SorareTokensRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SorareTokensRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SorareTokensRoleRevoked)
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
		it.Event = new(SorareTokensRoleRevoked)
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
func (it *SorareTokensRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SorareTokensRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SorareTokensRoleRevoked represents a RoleRevoked event raised by the SorareTokens contract.
type SorareTokensRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SorareTokens *SorareTokensFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SorareTokensRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SorareTokens.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SorareTokensRoleRevokedIterator{contract: _SorareTokens.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SorareTokens *SorareTokensFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SorareTokensRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SorareTokens.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SorareTokensRoleRevoked)
				if err := _SorareTokens.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SorareTokens *SorareTokensFilterer) ParseRoleRevoked(log types.Log) (*SorareTokensRoleRevoked, error) {
	event := new(SorareTokensRoleRevoked)
	if err := _SorareTokens.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SorareTokensTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SorareTokens contract.
type SorareTokensTransferIterator struct {
	Event *SorareTokensTransfer // Event containing the contract specifics and raw log

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
func (it *SorareTokensTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SorareTokensTransfer)
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
		it.Event = new(SorareTokensTransfer)
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
func (it *SorareTokensTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SorareTokensTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SorareTokensTransfer represents a Transfer event raised by the SorareTokens contract.
type SorareTokensTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SorareTokens *SorareTokensFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SorareTokensTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SorareTokens.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SorareTokensTransferIterator{contract: _SorareTokens.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SorareTokens *SorareTokensFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SorareTokensTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SorareTokens.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SorareTokensTransfer)
				if err := _SorareTokens.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SorareTokens *SorareTokensFilterer) ParseTransfer(log types.Log) (*SorareTokensTransfer, error) {
	event := new(SorareTokensTransfer)
	if err := _SorareTokens.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StringsABI is the input ABI used to generate the binding from.
const StringsABI = "[]"

// StringsBin is the compiled bytecode used for deploying new contracts.
var StringsBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220da4b33dead2364cb912c6aedb944bdac18432735a78ac109dc4a7f2f7bc8d0b864736f6c63430006060033"

// DeployStrings deploys a new Ethereum contract, binding an instance of Strings to it.
func DeployStrings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Strings, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StringsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	StringsCaller     // Read-only binding to the contract
	StringsTransactor // Write-only binding to the contract
	StringsFilterer   // Log filterer for contract events
}

// StringsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringsSession struct {
	Contract     *Strings          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringsCallerSession struct {
	Contract *StringsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StringsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringsTransactorSession struct {
	Contract     *StringsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StringsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringsRaw struct {
	Contract *Strings // Generic contract binding to access the raw methods on
}

// StringsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringsCallerRaw struct {
	Contract *StringsCaller // Generic read-only contract binding to access the raw methods on
}

// StringsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringsTransactorRaw struct {
	Contract *StringsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrings creates a new instance of Strings, bound to a specific deployed contract.
func NewStrings(address common.Address, backend bind.ContractBackend) (*Strings, error) {
	contract, err := bindStrings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// NewStringsCaller creates a new read-only instance of Strings, bound to a specific deployed contract.
func NewStringsCaller(address common.Address, caller bind.ContractCaller) (*StringsCaller, error) {
	contract, err := bindStrings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringsCaller{contract: contract}, nil
}

// NewStringsTransactor creates a new write-only instance of Strings, bound to a specific deployed contract.
func NewStringsTransactor(address common.Address, transactor bind.ContractTransactor) (*StringsTransactor, error) {
	contract, err := bindStrings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringsTransactor{contract: contract}, nil
}

// NewStringsFilterer creates a new log filterer instance of Strings, bound to a specific deployed contract.
func NewStringsFilterer(address common.Address, filterer bind.ContractFilterer) (*StringsFilterer, error) {
	contract, err := bindStrings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringsFilterer{contract: contract}, nil
}

// bindStrings binds a generic wrapper to an already deployed contract.
func bindStrings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.StringsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transact(opts, method, params...)
}
