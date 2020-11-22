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

// ProxyERC20ABI is the input ABI used to generate the binding from.
const ProxyERC20ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"setTarget\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"callData\",\"type\":\"bytes\"},{\"name\":\"numTopics\",\"type\":\"uint256\"},{\"name\":\"topic1\",\"type\":\"bytes32\"},{\"name\":\"topic2\",\"type\":\"bytes32\"},{\"name\":\"topic3\",\"type\":\"bytes32\"},{\"name\":\"topic4\",\"type\":\"bytes32\"}],\"name\":\"_emit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"useDELEGATECALL\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setUseDELEGATECALL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"target\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newTarget\",\"type\":\"address\"}],\"name\":\"TargetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"}]"

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
func (_ProxyERC20 *ProxyERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_ProxyERC20 *ProxyERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProxyERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProxyERC20.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
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
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ProxyERC20.contract.Call(opts, out, "decimals")
	return *ret0, err
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
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ProxyERC20.contract.Call(opts, out, "name")
	return *ret0, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProxyERC20.contract.Call(opts, out, "nominatedOwner")
	return *ret0, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProxyERC20.contract.Call(opts, out, "owner")
	return *ret0, err
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
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ProxyERC20.contract.Call(opts, out, "symbol")
	return *ret0, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProxyERC20.contract.Call(opts, out, "target")
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProxyERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
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
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ProxyERC20.contract.Call(opts, out, "useDELEGATECALL")
	return *ret0, err
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
	return event, nil
}
