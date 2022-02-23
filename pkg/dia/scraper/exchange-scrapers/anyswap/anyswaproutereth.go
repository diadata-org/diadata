// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package anyswaproutereth

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AnyswapV1ERC20MetaData contains all meta data concerning the AnyswapV1ERC20 contract.
var AnyswapV1ERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"changeVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9dc29fac": "burn(address,uint256)",
		"60e232a9": "changeVault(address)",
		"bebbf4d0": "depositVault(uint256,address)",
		"40c10f19": "mint(address,uint256)",
		"6f307dc3": "underlying()",
		"0039d6ec": "withdrawVault(address,uint256,address)",
	},
}

// AnyswapV1ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use AnyswapV1ERC20MetaData.ABI instead.
var AnyswapV1ERC20ABI = AnyswapV1ERC20MetaData.ABI

// Deprecated: Use AnyswapV1ERC20MetaData.Sigs instead.
// AnyswapV1ERC20FuncSigs maps the 4-byte function signature to its string representation.
var AnyswapV1ERC20FuncSigs = AnyswapV1ERC20MetaData.Sigs

// AnyswapV1ERC20 is an auto generated Go binding around an Ethereum contract.
type AnyswapV1ERC20 struct {
	AnyswapV1ERC20Caller     // Read-only binding to the contract
	AnyswapV1ERC20Transactor // Write-only binding to the contract
	AnyswapV1ERC20Filterer   // Log filterer for contract events
}

// AnyswapV1ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type AnyswapV1ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnyswapV1ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AnyswapV1ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnyswapV1ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AnyswapV1ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnyswapV1ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnyswapV1ERC20Session struct {
	Contract     *AnyswapV1ERC20   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnyswapV1ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnyswapV1ERC20CallerSession struct {
	Contract *AnyswapV1ERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AnyswapV1ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnyswapV1ERC20TransactorSession struct {
	Contract     *AnyswapV1ERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AnyswapV1ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type AnyswapV1ERC20Raw struct {
	Contract *AnyswapV1ERC20 // Generic contract binding to access the raw methods on
}

// AnyswapV1ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnyswapV1ERC20CallerRaw struct {
	Contract *AnyswapV1ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// AnyswapV1ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnyswapV1ERC20TransactorRaw struct {
	Contract *AnyswapV1ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAnyswapV1ERC20 creates a new instance of AnyswapV1ERC20, bound to a specific deployed contract.
func NewAnyswapV1ERC20(address common.Address, backend bind.ContractBackend) (*AnyswapV1ERC20, error) {
	contract, err := bindAnyswapV1ERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AnyswapV1ERC20{AnyswapV1ERC20Caller: AnyswapV1ERC20Caller{contract: contract}, AnyswapV1ERC20Transactor: AnyswapV1ERC20Transactor{contract: contract}, AnyswapV1ERC20Filterer: AnyswapV1ERC20Filterer{contract: contract}}, nil
}

// NewAnyswapV1ERC20Caller creates a new read-only instance of AnyswapV1ERC20, bound to a specific deployed contract.
func NewAnyswapV1ERC20Caller(address common.Address, caller bind.ContractCaller) (*AnyswapV1ERC20Caller, error) {
	contract, err := bindAnyswapV1ERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AnyswapV1ERC20Caller{contract: contract}, nil
}

// NewAnyswapV1ERC20Transactor creates a new write-only instance of AnyswapV1ERC20, bound to a specific deployed contract.
func NewAnyswapV1ERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*AnyswapV1ERC20Transactor, error) {
	contract, err := bindAnyswapV1ERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AnyswapV1ERC20Transactor{contract: contract}, nil
}

// NewAnyswapV1ERC20Filterer creates a new log filterer instance of AnyswapV1ERC20, bound to a specific deployed contract.
func NewAnyswapV1ERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*AnyswapV1ERC20Filterer, error) {
	contract, err := bindAnyswapV1ERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AnyswapV1ERC20Filterer{contract: contract}, nil
}

// bindAnyswapV1ERC20 binds a generic wrapper to an already deployed contract.
func bindAnyswapV1ERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AnyswapV1ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnyswapV1ERC20 *AnyswapV1ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AnyswapV1ERC20.Contract.AnyswapV1ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnyswapV1ERC20 *AnyswapV1ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnyswapV1ERC20.Contract.AnyswapV1ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnyswapV1ERC20 *AnyswapV1ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnyswapV1ERC20.Contract.AnyswapV1ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnyswapV1ERC20 *AnyswapV1ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AnyswapV1ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnyswapV1ERC20 *AnyswapV1ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnyswapV1ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnyswapV1ERC20 *AnyswapV1ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnyswapV1ERC20.Contract.contract.Transact(opts, method, params...)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_AnyswapV1ERC20 *AnyswapV1ERC20Caller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnyswapV1ERC20.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_AnyswapV1ERC20 *AnyswapV1ERC20Session) Underlying() (common.Address, error) {
	return _AnyswapV1ERC20.Contract.Underlying(&_AnyswapV1ERC20.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_AnyswapV1ERC20 *AnyswapV1ERC20CallerSession) Underlying() (common.Address, error) {
	return _AnyswapV1ERC20.Contract.Underlying(&_AnyswapV1ERC20.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns(bool)
func (_AnyswapV1ERC20 *AnyswapV1ERC20Transactor) Burn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AnyswapV1ERC20.contract.Transact(opts, "burn", from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns(bool)
func (_AnyswapV1ERC20 *AnyswapV1ERC20Session) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AnyswapV1ERC20.Contract.Burn(&_AnyswapV1ERC20.TransactOpts, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns(bool)
func (_AnyswapV1ERC20 *AnyswapV1ERC20TransactorSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AnyswapV1ERC20.Contract.Burn(&_AnyswapV1ERC20.TransactOpts, from, amount)
}

// ChangeVault is a paid mutator transaction binding the contract method 0x60e232a9.
//
// Solidity: function changeVault(address newVault) returns(bool)
func (_AnyswapV1ERC20 *AnyswapV1ERC20Transactor) ChangeVault(opts *bind.TransactOpts, newVault common.Address) (*types.Transaction, error) {
	return _AnyswapV1ERC20.contract.Transact(opts, "changeVault", newVault)
}

// ChangeVault is a paid mutator transaction binding the contract method 0x60e232a9.
//
// Solidity: function changeVault(address newVault) returns(bool)
func (_AnyswapV1ERC20 *AnyswapV1ERC20Session) ChangeVault(newVault common.Address) (*types.Transaction, error) {
	return _AnyswapV1ERC20.Contract.ChangeVault(&_AnyswapV1ERC20.TransactOpts, newVault)
}

// ChangeVault is a paid mutator transaction binding the contract method 0x60e232a9.
//
// Solidity: function changeVault(address newVault) returns(bool)
func (_AnyswapV1ERC20 *AnyswapV1ERC20TransactorSession) ChangeVault(newVault common.Address) (*types.Transaction, error) {
	return _AnyswapV1ERC20.Contract.ChangeVault(&_AnyswapV1ERC20.TransactOpts, newVault)
}

// DepositVault is a paid mutator transaction binding the contract method 0xbebbf4d0.
//
// Solidity: function depositVault(uint256 amount, address to) returns(uint256)
func (_AnyswapV1ERC20 *AnyswapV1ERC20Transactor) DepositVault(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AnyswapV1ERC20.contract.Transact(opts, "depositVault", amount, to)
}

// DepositVault is a paid mutator transaction binding the contract method 0xbebbf4d0.
//
// Solidity: function depositVault(uint256 amount, address to) returns(uint256)
func (_AnyswapV1ERC20 *AnyswapV1ERC20Session) DepositVault(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AnyswapV1ERC20.Contract.DepositVault(&_AnyswapV1ERC20.TransactOpts, amount, to)
}

// DepositVault is a paid mutator transaction binding the contract method 0xbebbf4d0.
//
// Solidity: function depositVault(uint256 amount, address to) returns(uint256)
func (_AnyswapV1ERC20 *AnyswapV1ERC20TransactorSession) DepositVault(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AnyswapV1ERC20.Contract.DepositVault(&_AnyswapV1ERC20.TransactOpts, amount, to)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_AnyswapV1ERC20 *AnyswapV1ERC20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AnyswapV1ERC20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_AnyswapV1ERC20 *AnyswapV1ERC20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AnyswapV1ERC20.Contract.Mint(&_AnyswapV1ERC20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_AnyswapV1ERC20 *AnyswapV1ERC20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AnyswapV1ERC20.Contract.Mint(&_AnyswapV1ERC20.TransactOpts, to, amount)
}

// WithdrawVault is a paid mutator transaction binding the contract method 0x0039d6ec.
//
// Solidity: function withdrawVault(address from, uint256 amount, address to) returns(uint256)
func (_AnyswapV1ERC20 *AnyswapV1ERC20Transactor) WithdrawVault(opts *bind.TransactOpts, from common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AnyswapV1ERC20.contract.Transact(opts, "withdrawVault", from, amount, to)
}

// WithdrawVault is a paid mutator transaction binding the contract method 0x0039d6ec.
//
// Solidity: function withdrawVault(address from, uint256 amount, address to) returns(uint256)
func (_AnyswapV1ERC20 *AnyswapV1ERC20Session) WithdrawVault(from common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AnyswapV1ERC20.Contract.WithdrawVault(&_AnyswapV1ERC20.TransactOpts, from, amount, to)
}

// WithdrawVault is a paid mutator transaction binding the contract method 0x0039d6ec.
//
// Solidity: function withdrawVault(address from, uint256 amount, address to) returns(uint256)
func (_AnyswapV1ERC20 *AnyswapV1ERC20TransactorSession) WithdrawVault(from common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AnyswapV1ERC20.Contract.WithdrawVault(&_AnyswapV1ERC20.TransactOpts, from, amount, to)
}

// AnyswapV4RouterMetaData contains all meta data concerning the AnyswapV4Router contract.
var AnyswapV4RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wNATIVE\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mpc\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txhash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"LogAnySwapIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"LogAnySwapOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"LogAnySwapTradeTokensForNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"LogAnySwapTradeTokensForTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMPC\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMPC\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"effectiveTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"LogChangeMPC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRouter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"LogChangeRouter\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"anySwapFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"txs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fromChainIDs\",\"type\":\"uint256[]\"}],\"name\":\"anySwapIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txs\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txs\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapInAuto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txs\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapInExactTokensForNative\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txs\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapInExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txs\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapInUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"toChainIDs\",\"type\":\"uint256[]\"}],\"name\":\"anySwapOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForNativeUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForNativeUnderlyingWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForNativeUnderlyingWithTransferPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForTokensUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForTokensUnderlyingWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForTokensUnderlyingWithTransferPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutUnderlyingWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutUnderlyingWithTransferPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMPC\",\"type\":\"address\"}],\"name\":\"changeMPC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"changeVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mpc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wNATIVE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"87cc6e2f": "anySwapFeeTo(address,uint256)",
		"825bb13c": "anySwapIn(bytes32,address,address,uint256,uint256)",
		"25121b76": "anySwapIn(bytes32[],address[],address[],uint256[],uint256[])",
		"0175b1c4": "anySwapInAuto(bytes32,address,address,uint256,uint256)",
		"52a397d5": "anySwapInExactTokensForNative(bytes32,uint256,uint256,address[],address,uint256,uint256)",
		"2fc1e728": "anySwapInExactTokensForTokens(bytes32,uint256,uint256,address[],address,uint256,uint256)",
		"3f88de89": "anySwapInUnderlying(bytes32,address,address,uint256,uint256)",
		"241dc2df": "anySwapOut(address,address,uint256,uint256)",
		"dcfb77b1": "anySwapOut(address[],address[],uint256[],uint256[])",
		"65782f56": "anySwapOutExactTokensForNative(uint256,uint256,address[],address,uint256,uint256)",
		"6a453972": "anySwapOutExactTokensForNativeUnderlying(uint256,uint256,address[],address,uint256,uint256)",
		"4d93bb94": "anySwapOutExactTokensForNativeUnderlyingWithPermit(address,uint256,uint256,address[],address,uint256,uint8,bytes32,bytes32,uint256)",
		"c8e174f6": "anySwapOutExactTokensForNativeUnderlyingWithTransferPermit(address,uint256,uint256,address[],address,uint256,uint8,bytes32,bytes32,uint256)",
		"0bb57203": "anySwapOutExactTokensForTokens(uint256,uint256,address[],address,uint256,uint256)",
		"d8b9f610": "anySwapOutExactTokensForTokensUnderlying(uint256,uint256,address[],address,uint256,uint256)",
		"99cd84b5": "anySwapOutExactTokensForTokensUnderlyingWithPermit(address,uint256,uint256,address[],address,uint256,uint8,bytes32,bytes32,uint256)",
		"9aa1ac61": "anySwapOutExactTokensForTokensUnderlyingWithTransferPermit(address,uint256,uint256,address[],address,uint256,uint8,bytes32,bytes32,uint256)",
		"edbdf5e2": "anySwapOutUnderlying(address,address,uint256,uint256)",
		"8d7d3eea": "anySwapOutUnderlyingWithPermit(address,address,address,uint256,uint256,uint8,bytes32,bytes32,uint256)",
		"1b91a934": "anySwapOutUnderlyingWithTransferPermit(address,address,address,uint256,uint256,uint8,bytes32,bytes32,uint256)",
		"99a2f2d7": "cID()",
		"5b7b018c": "changeMPC(address)",
		"456862aa": "changeVault(address,address)",
		"c45a0155": "factory()",
		"85f8c259": "getAmountIn(uint256,uint256,uint256)",
		"054d50d4": "getAmountOut(uint256,uint256,uint256)",
		"1f00ca74": "getAmountsIn(uint256,address[])",
		"d06ca61f": "getAmountsOut(uint256,address[])",
		"f75c2664": "mpc()",
		"ad615dec": "quote(uint256,uint256,uint256)",
		"8fd903f5": "wNATIVE()",
	},
	Bin: "0x60c06040523480156200001157600080fd5b506040516200499738038062004997833981016040819052620000349162000095565b600180546001600160a01b039092166001600160a01b0319909216919091179055426002556001600160601b0319606092831b8116608052911b1660a052620000de565b80516001600160a01b03811681146200009057600080fd5b919050565b600080600060608486031215620000aa578283fd5b620000b58462000078565b9250620000c56020850162000078565b9150620000d56040850162000078565b90509250925092565b60805160601c60a05160601c61484762000150600039600081816101ec01528181611499015281816116380152611e3d015260008181610aef01528181610cda01528181610dc8015281816115230152818161254b015281816128a701528181613322015261338001526148476000f3fe6080604052600436106101dc5760003560e01c8063825bb13c11610102578063ad615dec11610095578063d8b9f61011610064578063d8b9f610146105aa578063dcfb77b1146105ca578063edbdf5e2146105ea578063f75c26641461060a57610229565b8063ad615dec14610535578063c45a015514610555578063c8e174f61461056a578063d06ca61f1461058a57610229565b80638fd903f5116100d15780638fd903f5146104be57806399a2f2d7146104e057806399cd84b5146104f55780639aa1ac611461051557610229565b8063825bb13c1461043e57806385f8c2591461045e57806387cc6e2f1461047e5780638d7d3eea1461049e57610229565b80632fc1e7281161017a57806352a397d51161014957806352a397d5146103be5780635b7b018c146103de57806365782f56146103fe5780636a4539721461041e57610229565b80632fc1e728146103315780633f88de8914610351578063456862aa146103715780634d93bb941461039e57610229565b80631b91a934116101b65780631b91a934146102a45780631f00ca74146102c4578063241dc2df146102f157806325121b761461031157610229565b80630175b1c41461022e578063054d50d41461024e5780630bb572031461028457610229565b3661022957336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461022757634e487b7160e01b600052600160045260246000fd5b005b600080fd5b34801561023a57600080fd5b50610227610249366004613d21565b61061f565b34801561025a57600080fd5b5061026e610269366004613fa3565b61080b565b60405161027b919061465b565b60405180910390f35b34801561029057600080fd5b5061022761029f366004613f28565b610820565b3480156102b057600080fd5b506102276102bf3660046139aa565b610954565b3480156102d057600080fd5b506102e46102df366004613e5b565b610ae8565b60405161027b91906141ee565b3480156102fd57600080fd5b5061022761030c366004613a34565b610b1e565b34801561031d57600080fd5b5061022761032c366004613c16565b610b31565b34801561033d57600080fd5b506102e461034c366004613d72565b610c7a565b34801561035d57600080fd5b5061022761036c366004613d21565b610ed3565b34801561037d57600080fd5b5061039161038c366004613972565b610fa0565b60405161027b9190614232565b3480156103aa57600080fd5b506102276103b9366004613aa4565b61107e565b3480156103ca57600080fd5b506102e46103d9366004613d72565b611434565b3480156103ea57600080fd5b506103916103f9366004613933565b611726565b34801561040a57600080fd5b50610227610419366004613f28565b611829565b34801561042a57600080fd5b50610227610439366004613f28565b611938565b34801561044a57600080fd5b50610227610459366004613d21565b611b16565b34801561046a57600080fd5b5061026e610479366004613fa3565b611b62565b34801561048a57600080fd5b50610227610499366004613a79565b611b6f565b3480156104aa57600080fd5b506102276104b93660046139aa565b611cb5565b3480156104ca57600080fd5b506104d3611e3b565b60405161027b9190614067565b3480156104ec57600080fd5b5061026e611e5f565b34801561050157600080fd5b50610227610510366004613aa4565b611e63565b34801561052157600080fd5b50610227610530366004613aa4565b6121e0565b34801561054157600080fd5b5061026e610550366004613fa3565b61253c565b34801561056157600080fd5b506104d3612549565b34801561057657600080fd5b50610227610585366004613aa4565b61256d565b34801561059657600080fd5b506102e46105a5366004613e5b565b6128a0565b3480156105b657600080fd5b506102276105c5366004613f28565b6128cd565b3480156105d657600080fd5b506102276105e5366004613b57565b6129e9565b3480156105f657600080fd5b50610227610605366004613a34565b612ac7565b34801561061657600080fd5b506104d3612bd1565b610627612bd1565b6001600160a01b0316336001600160a01b0316146106605760405162461bcd60e51b81526004016106579061444d565b60405180910390fd5b61066d8585858585612bff565b60008490506000816001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156106ad57600080fd5b505afa1580156106c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e59190613956565b90506001600160a01b0381161580159061077b57506040516370a0823160e01b815284906001600160a01b038316906370a0823190610728908a90600401614067565b60206040518083038186803b15801561074057600080fd5b505afa158015610754573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107789190613e43565b10155b1561080257604051620e75bb60e21b81526001600160a01b038316906239d6ec906107ae908890889082906004016140f9565b602060405180830381600087803b1580156107c857600080fd5b505af11580156107dc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108009190613e43565b505b50505050505050565b6000610818848484612cde565b949350505050565b81428110156108415760405162461bcd60e51b815260040161065790614570565b8585600081811061086257634e487b7160e01b600052603260045260246000fd5b90506020020160208101906108779190613933565b6001600160a01b0316639dc29fac338a6040518363ffffffff1660e01b81526004016108a49291906140e0565b602060405180830381600087803b1580156108be57600080fd5b505af11580156108d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108f69190613d01565b506001600160a01b038416337ffea6abdf4fd32f20966dff7619354cd82cd43dc78a3bee479f04c74dbfc585b388888c8c61092f611e5f565b896040516109429695949392919061411c565b60405180910390a35050505050505050565b876001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b15801561098d57600080fd5b505afa1580156109a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109c59190613956565b6001600160a01b031663605629d68a8a89898989896040518863ffffffff1660e01b81526004016109fc979695949392919061409f565b602060405180830381600087803b158015610a1657600080fd5b505af1158015610a2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a4e9190613d01565b50604051630bebbf4d60e41b81526001600160a01b0389169063bebbf4d090610a7d9089908d90600401614664565b602060405180830381600087803b158015610a9757600080fd5b505af1158015610aab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610acf9190613e43565b50610add8989898985612d78565b505050505050505050565b6060610b157f00000000000000000000000000000000000000000000000000000000000000008484612e51565b90505b92915050565b610b2b3385858585612d78565b50505050565b610b39612bd1565b6001600160a01b0316336001600160a01b031614610b695760405162461bcd60e51b81526004016106579061444d565b60005b87811015610c6d57610c5b8b8b83818110610b9757634e487b7160e01b600052603260045260246000fd5b905060200201358a8a84818110610bbe57634e487b7160e01b600052603260045260246000fd5b9050602002016020810190610bd39190613933565b898985818110610bf357634e487b7160e01b600052603260045260246000fd5b9050602002016020810190610c089190613933565b888886818110610c2857634e487b7160e01b600052603260045260246000fd5b90506020020135878787818110610c4f57634e487b7160e01b600052603260045260246000fd5b90506020020135612bff565b80610c6581614792565b915050610b6c565b5050505050505050505050565b6060610c84612bd1565b6001600160a01b0316336001600160a01b031614610cb45760405162461bcd60e51b81526004016106579061444d565b8242811015610cd55760405162461bcd60e51b815260040161065790614570565b610d337f00000000000000000000000000000000000000000000000000000000000000008a89898080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061300b92505050565b9150878260018451610d459190614738565b81518110610d6357634e487b7160e01b600052603260045260246000fd5b60200260200101511015610d895760405162461bcd60e51b8152600401610657906143ba565b610e878a88886000818110610dae57634e487b7160e01b600052603260045260246000fd5b9050602002016020810190610dc39190613933565b610e587f00000000000000000000000000000000000000000000000000000000000000008b8b6000818110610e0857634e487b7160e01b600052603260045260246000fd5b9050602002016020810190610e1d9190613933565b8c8c6001818110610e3e57634e487b7160e01b600052603260045260246000fd5b9050602002016020810190610e539190613933565b6131a5565b85600081518110610e7957634e487b7160e01b600052603260045260246000fd5b602002602001015187612bff565b610ec6828888808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508a9250613215915050565b5098975050505050505050565b610edb612bd1565b6001600160a01b0316336001600160a01b031614610f0b5760405162461bcd60e51b81526004016106579061444d565b610f188585858585612bff565b604051620e75bb60e21b81526001600160a01b038516906239d6ec90610f46908690869082906004016140f9565b602060405180830381600087803b158015610f6057600080fd5b505af1158015610f74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f989190613e43565b505050505050565b6000610faa612bd1565b6001600160a01b0316336001600160a01b031614610fda5760405162461bcd60e51b81526004016106579061444d565b6001600160a01b0382166110005760405162461bcd60e51b8152600401610657906142eb565b6040516360e232a960e01b81526001600160a01b038416906360e232a99061102c908590600401614067565b602060405180830381600087803b15801561104657600080fd5b505af115801561105a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b159190613d01565b844281101561109f5760405162461bcd60e51b815260040161065790614570565b6000898960008181106110c257634e487b7160e01b600052603260045260246000fd5b90506020020160208101906110d79190613933565b6001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b15801561110f57600080fd5b505afa158015611123573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111479190613956565b9050806001600160a01b031663d505accf8e308f8b8b8b8b6040518863ffffffff1660e01b8152600401611181979695949392919061409f565b600060405180830381600087803b15801561119b57600080fd5b505af11580156111af573d6000803e3d6000fd5b505050506111f4818e8c8c60008181106111d957634e487b7160e01b600052603260045260246000fd5b90506020020160208101906111ee9190613933565b8f61344f565b8989600081811061121557634e487b7160e01b600052603260045260246000fd5b905060200201602081019061122a9190613933565b6001600160a01b031663bebbf4d08d8f6040518363ffffffff1660e01b8152600401611257929190614664565b602060405180830381600087803b15801561127157600080fd5b505af1158015611285573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112a99190613e43565b50898960008181106112cb57634e487b7160e01b600052603260045260246000fd5b90506020020160208101906112e09190613933565b6001600160a01b0316639dc29fac8e8e6040518363ffffffff1660e01b815260040161130d9291906140e0565b602060405180830381600087803b15801561132757600080fd5b505af115801561133b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061135f9190613d01565b5060008a8a80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050905060008e905060008a905060008f905060008f905060006113c4611e5f565b90506000899050846001600160a01b0316866001600160a01b03167f278277e0209c347189add7bd92411973b5f6b8644f7ac62ea1be984ce993f8f48987878787604051611416959493929190614186565b60405180910390a35050505050505050505050505050505050505050565b606061143e612bd1565b6001600160a01b0316336001600160a01b03161461146e5760405162461bcd60e51b81526004016106579061444d565b824281101561148f5760405162461bcd60e51b815260040161065790614570565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001687876114c6600182614738565b8181106114e357634e487b7160e01b600052603260045260246000fd5b90506020020160208101906114f89190613933565b6001600160a01b03161461151e5760405162461bcd60e51b8152600401610657906144b9565b61157c7f00000000000000000000000000000000000000000000000000000000000000008a89898080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061300b92505050565b915087826001845161158e9190614738565b815181106115ac57634e487b7160e01b600052603260045260246000fd5b602002602001015110156115d25760405162461bcd60e51b8152600401610657906144f0565b6115f78a88886000818110610dae57634e487b7160e01b600052603260045260246000fd5b61163682888880806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250309250613215915050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632e1a7d4d83600185516116749190614738565b8151811061169257634e487b7160e01b600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b81526004016116b6919061465b565b600060405180830381600087803b1580156116d057600080fd5b505af11580156116e4573d6000803e3d6000fd5b50505050610ec68583600185516116fb9190614738565b8151811061171957634e487b7160e01b600052603260045260246000fd5b6020026020010151613537565b6000611730612bd1565b6001600160a01b0316336001600160a01b0316146117605760405162461bcd60e51b81526004016106579061444d565b6001600160a01b0382166117865760405162461bcd60e51b8152600401610657906142eb565b61178e612bd1565b600080546001600160a01b03199081166001600160a01b0393841617909155600180549091169184169190911790556117ca426202a3006146e1565b60028190556001546000546001600160a01b0391821691167fcda32bc39904597666dfa9f9c845714756e1ffffad55b52e0d344673a219812161180b611e5f565b604051611818919061465b565b60405180910390a45060015b919050565b814281101561184a5760405162461bcd60e51b815260040161065790614570565b8585600081811061186b57634e487b7160e01b600052603260045260246000fd5b90506020020160208101906118809190613933565b6001600160a01b0316639dc29fac338a6040518363ffffffff1660e01b81526004016118ad9291906140e0565b602060405180830381600087803b1580156118c757600080fd5b505af11580156118db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118ff9190613d01565b506001600160a01b038416337f278277e0209c347189add7bd92411973b5f6b8644f7ac62ea1be984ce993f8f488888c8c61092f611e5f565b81428110156119595760405162461bcd60e51b815260040161065790614570565b611a3f8686600081811061197d57634e487b7160e01b600052603260045260246000fd5b90506020020160208101906119929190613933565b6001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156119ca57600080fd5b505afa1580156119de573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a029190613956565b3388886000818110611a2457634e487b7160e01b600052603260045260246000fd5b9050602002016020810190611a399190613933565b8b61344f565b85856000818110611a6057634e487b7160e01b600052603260045260246000fd5b9050602002016020810190611a759190613933565b6001600160a01b031663bebbf4d089336040518363ffffffff1660e01b8152600401611aa2929190614664565b602060405180830381600087803b158015611abc57600080fd5b505af1158015611ad0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611af49190613e43565b508585600081811061186b57634e487b7160e01b600052603260045260246000fd5b611b1e612bd1565b6001600160a01b0316336001600160a01b031614611b4e5760405162461bcd60e51b81526004016106579061444d565b611b5b8585858585612bff565b5050505050565b60006108188484846135c9565b611b77612bd1565b6001600160a01b0316336001600160a01b031614611ba75760405162461bcd60e51b81526004016106579061444d565b6000611bb1612bd1565b6040516340c10f1960e01b81529091506001600160a01b038416906340c10f1990611be290849086906004016140e0565b602060405180830381600087803b158015611bfc57600080fd5b505af1158015611c10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c349190613d01565b50604051620e75bb60e21b81526001600160a01b038416906239d6ec90611c63908490869082906004016140f9565b602060405180830381600087803b158015611c7d57600080fd5b505af1158015611c91573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b2b9190613e43565b6000886001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015611cf057600080fd5b505afa158015611d04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d289190613956565b60405163d505accf60e01b81529091506001600160a01b0382169063d505accf90611d63908d9030908c908c908c908c908c9060040161409f565b600060405180830381600087803b158015611d7d57600080fd5b505af1158015611d91573d6000803e3d6000fd5b50505050611da1818b8b8a61344f565b604051630bebbf4d60e41b81526001600160a01b038a169063bebbf4d090611dcf908a908e90600401614664565b602060405180830381600087803b158015611de957600080fd5b505af1158015611dfd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e219190613e43565b50611e2f8a8a8a8a86612d78565b50505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b4690565b8442811015611e845760405162461bcd60e51b815260040161065790614570565b600089896000818110611ea757634e487b7160e01b600052603260045260246000fd5b9050602002016020810190611ebc9190613933565b6001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015611ef457600080fd5b505afa158015611f08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f2c9190613956565b9050806001600160a01b031663d505accf8e308f8b8b8b8b6040518863ffffffff1660e01b8152600401611f66979695949392919061409f565b600060405180830381600087803b158015611f8057600080fd5b505af1158015611f94573d6000803e3d6000fd5b50505050611fbe818e8c8c60008181106111d957634e487b7160e01b600052603260045260246000fd5b89896000818110611fdf57634e487b7160e01b600052603260045260246000fd5b9050602002016020810190611ff49190613933565b6001600160a01b031663bebbf4d08d8f6040518363ffffffff1660e01b8152600401612021929190614664565b602060405180830381600087803b15801561203b57600080fd5b505af115801561204f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120739190613e43565b508989600081811061209557634e487b7160e01b600052603260045260246000fd5b90506020020160208101906120aa9190613933565b6001600160a01b0316639dc29fac8e8e6040518363ffffffff1660e01b81526004016120d79291906140e0565b602060405180830381600087803b1580156120f157600080fd5b505af1158015612105573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121299190613d01565b5060008a8a80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050905060008e905060008a905060008f905060008f9050600061218e611e5f565b90506000899050846001600160a01b0316866001600160a01b03167ffea6abdf4fd32f20966dff7619354cd82cd43dc78a3bee479f04c74dbfc585b38987878787604051611416959493929190614186565b84428110156122015760405162461bcd60e51b815260040161065790614570565b8888600081811061222257634e487b7160e01b600052603260045260246000fd5b90506020020160208101906122379190613933565b6001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b15801561226f57600080fd5b505afa158015612283573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122a79190613956565b6001600160a01b031663605629d68d8b8b60008181106122d757634e487b7160e01b600052603260045260246000fd5b90506020020160208101906122ec9190613933565b8e8a8a8a8a6040518863ffffffff1660e01b8152600401612313979695949392919061409f565b602060405180830381600087803b15801561232d57600080fd5b505af1158015612341573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123659190613d01565b508888600081811061238757634e487b7160e01b600052603260045260246000fd5b905060200201602081019061239c9190613933565b6001600160a01b031663bebbf4d08c8e6040518363ffffffff1660e01b81526004016123c9929190614664565b602060405180830381600087803b1580156123e357600080fd5b505af11580156123f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061241b9190613e43565b508888600081811061243d57634e487b7160e01b600052603260045260246000fd5b90506020020160208101906124529190613933565b6001600160a01b0316639dc29fac8d8d6040518363ffffffff1660e01b815260040161247f9291906140e0565b602060405180830381600087803b15801561249957600080fd5b505af11580156124ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124d19190613d01565b50866001600160a01b03168c6001600160a01b03167ffea6abdf4fd32f20966dff7619354cd82cd43dc78a3bee479f04c74dbfc585b38b8b8f8f612513611e5f565b896040516125269695949392919061411c565b60405180910390a3505050505050505050505050565b600061081884848461365d565b7f000000000000000000000000000000000000000000000000000000000000000081565b844281101561258e5760405162461bcd60e51b815260040161065790614570565b888860008181106125af57634e487b7160e01b600052603260045260246000fd5b90506020020160208101906125c49190613933565b6001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156125fc57600080fd5b505afa158015612610573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126349190613956565b6001600160a01b031663605629d68d8b8b600081811061266457634e487b7160e01b600052603260045260246000fd5b90506020020160208101906126799190613933565b8e8a8a8a8a6040518863ffffffff1660e01b81526004016126a0979695949392919061409f565b602060405180830381600087803b1580156126ba57600080fd5b505af11580156126ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126f29190613d01565b508888600081811061271457634e487b7160e01b600052603260045260246000fd5b90506020020160208101906127299190613933565b6001600160a01b031663bebbf4d08c8e6040518363ffffffff1660e01b8152600401612756929190614664565b602060405180830381600087803b15801561277057600080fd5b505af1158015612784573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127a89190613e43565b50888860008181106127ca57634e487b7160e01b600052603260045260246000fd5b90506020020160208101906127df9190613933565b6001600160a01b0316639dc29fac8d8d6040518363ffffffff1660e01b815260040161280c9291906140e0565b602060405180830381600087803b15801561282657600080fd5b505af115801561283a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061285e9190613d01565b50866001600160a01b03168c6001600160a01b03167f278277e0209c347189add7bd92411973b5f6b8644f7ac62ea1be984ce993f8f48b8b8f8f612513611e5f565b6060610b157f0000000000000000000000000000000000000000000000000000000000000000848461300b565b81428110156128ee5760405162461bcd60e51b815260040161065790614570565b6129128686600081811061197d57634e487b7160e01b600052603260045260246000fd5b8585600081811061293357634e487b7160e01b600052603260045260246000fd5b90506020020160208101906129489190613933565b6001600160a01b031663bebbf4d089336040518363ffffffff1660e01b8152600401612975929190614664565b602060405180830381600087803b15801561298f57600080fd5b505af11580156129a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129c79190613e43565b508585600081811061086257634e487b7160e01b600052603260045260246000fd5b60005b87811015610add57612ab5338a8a84818110612a1857634e487b7160e01b600052603260045260246000fd5b9050602002016020810190612a2d9190613933565b898985818110612a4d57634e487b7160e01b600052603260045260246000fd5b9050602002016020810190612a629190613933565b888886818110612a8257634e487b7160e01b600052603260045260246000fd5b90506020020135878787818110612aa957634e487b7160e01b600052603260045260246000fd5b90506020020135612d78565b80612abf81614792565b9150506129ec565b612b43846001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015612b0357600080fd5b505afa158015612b17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b3b9190613956565b33868561344f565b604051630bebbf4d60e41b81526001600160a01b0385169063bebbf4d090612b719085903390600401614664565b602060405180830381600087803b158015612b8b57600080fd5b505af1158015612b9f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bc39190613e43565b50610b2b3385858585612d78565b60006002544210612bee57506001546001600160a01b0316612bfc565b506000546001600160a01b03165b90565b6040516340c10f1960e01b81526001600160a01b038516906340c10f1990612c2d90869086906004016140e0565b602060405180830381600087803b158015612c4757600080fd5b505af1158015612c5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c7f9190613d01565b50826001600160a01b0316846001600160a01b0316867faac9ce45fe3adf5143598c4f18a369591a20a3384aedaf1b525d29127e1fcd558585612cc0611e5f565b604051612ccf939291906146cb565b60405180910390a45050505050565b6000808411612cff5760405162461bcd60e51b8152600401610657906145a7565b600083118015612d0f5750600082115b612d2b5760405162461bcd60e51b81526004016106579061423d565b6000612d39856103e56136bf565b90506000612d4782856136bf565b90506000612d6183612d5b886103e86136bf565b906136ff565b9050612d6d81836146f9565b979650505050505050565b604051632770a7eb60e21b81526001600160a01b03851690639dc29fac90612da690889086906004016140e0565b602060405180830381600087803b158015612dc057600080fd5b505af1158015612dd4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612df89190613d01565b50826001600160a01b0316856001600160a01b0316856001600160a01b03167f97116cf6cd4f6412bb47914d6db18da9e16ab2142f543b86e207c24fbd16b23a85612e41611e5f565b86604051612ccf939291906146cb565b6060600282511015612e755760405162461bcd60e51b815260040161065790614626565b815167ffffffffffffffff811115612e9d57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015612ec6578160200160208202803683370190505b509050828160018351612ed99190614738565b81518110612ef757634e487b7160e01b600052603260045260246000fd5b602002602001018181525050600060018351612f139190614738565b90505b801561300357600080612f828786612f2f600187614738565b81518110612f4d57634e487b7160e01b600052603260045260246000fd5b6020026020010151878681518110612f7557634e487b7160e01b600052603260045260246000fd5b602002602001015161372d565b91509150612fb8848481518110612fa957634e487b7160e01b600052603260045260246000fd5b602002602001015183836135c9565b84612fc4600186614738565b81518110612fe257634e487b7160e01b600052603260045260246000fd5b60200260200101818152505050508080612ffb9061477b565b915050612f16565b509392505050565b606060028251101561302f5760405162461bcd60e51b815260040161065790614626565b815167ffffffffffffffff81111561305757634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015613080578160200160208202803683370190505b50905082816000815181106130a557634e487b7160e01b600052603260045260246000fd5b60200260200101818152505060005b600183516130c29190614738565b81101561300357600080613124878685815181106130f057634e487b7160e01b600052603260045260246000fd5b60200260200101518786600161310691906146e1565b81518110612f7557634e487b7160e01b600052603260045260246000fd5b9150915061315a84848151811061314b57634e487b7160e01b600052603260045260246000fd5b60200260200101518383612cde565b846131668560016146e1565b8151811061318457634e487b7160e01b600052603260045260246000fd5b6020026020010181815250505050808061319d90614792565b9150506130b4565b60008060006131b48585613806565b915091508582826040516020016131cc929190613fce565b604051602081830303815290604052805190602001206040516020016131f3929190614011565b60408051601f1981840301815291905280516020909101209695505050505050565b60005b600183516132269190614738565b811015610b2b5760008084838151811061325057634e487b7160e01b600052603260045260246000fd5b60200260200101518584600161326691906146e1565b8151811061328457634e487b7160e01b600052603260045260246000fd5b602002602001015191509150600061329c8383613806565b5090506000876132ad8660016146e1565b815181106132cb57634e487b7160e01b600052603260045260246000fd5b60200260200101519050600080836001600160a01b0316866001600160a01b0316146132f9578260006132fd565b6000835b91509150600060028a516133119190614738565b881061331d5788613379565b6133797f0000000000000000000000000000000000000000000000000000000000000000878c61334e8c60026146e1565b8151811061336c57634e487b7160e01b600052603260045260246000fd5b60200260200101516131a5565b90506133a67f000000000000000000000000000000000000000000000000000000000000000088886131a5565b6001600160a01b031663022c0d9f84848460006040519080825280601f01601f1916602001820160405280156133e3576020820181803683370190505b506040518563ffffffff1660e01b8152600401613403949392919061467b565b600060405180830381600087803b15801561341d57600080fd5b505af1158015613431573d6000803e3d6000fd5b5050505050505050505050808061344790614792565b915050613218565b600080856001600160a01b03166323b872dd8686866040516024016134769392919061407b565b6040516020818303038152906040529060e01b6020820180516001600160e01b0383818316178352505050506040516134af9190613ff5565b6000604051808303816000865af19150503d80600081146134ec576040519150601f19603f3d011682016040523d82523d6000602084013e6134f1565b606091505b509150915081801561351b57508051158061351b57508080602001905181019061351b9190613d01565b610f985760405162461bcd60e51b8152600401610657906145e2565b604080516000808252602082019092526001600160a01b0384169083906040516135619190613ff5565b60006040518083038185875af1925050503d806000811461359e576040519150601f19603f3d011682016040523d82523d6000602084013e6135a3565b606091505b50509050806135c45760405162461bcd60e51b815260040161065790614407565b505050565b60008084116135ea5760405162461bcd60e51b815260040161065790614350565b6000831180156135fa5750600082115b6136165760405162461bcd60e51b81526004016106579061423d565b600061362e6103e861362886886136bf565b906136bf565b905060006136426103e56136288689613890565b90506136536001612d5b83856146f9565b9695505050505050565b600080841161367e5760405162461bcd60e51b81526004016106579061453b565b60008311801561368e5750600082115b6136aa5760405162461bcd60e51b81526004016106579061423d565b826136b585846136bf565b61081891906146f9565b60008115806136e3575082826136d58183614719565b92506136e190836146f9565b145b610b185760405162461bcd60e51b815260040161065790614322565b60008261370c83826146e1565b9150811015610b185760405162461bcd60e51b81526004016106579061438c565b600080600061373c8585613806565b50905060008061374d8888886131a5565b6001600160a01b0316630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b15801561378557600080fd5b505afa158015613799573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137bd9190613df5565b506001600160701b031691506001600160701b03169150826001600160a01b0316876001600160a01b0316146137f45780826137f7565b81815b90999098509650505050505050565b600080826001600160a01b0316846001600160a01b0316141561383b5760405162461bcd60e51b8152600401610657906142a4565b826001600160a01b0316846001600160a01b03161061385b57828461385e565b83835b90925090506001600160a01b0382166138895760405162461bcd60e51b815260040161065790614484565b9250929050565b60008261389d8382614738565b9150811115610b185760405162461bcd60e51b815260040161065790614275565b8035611824816147d9565b60008083601f8401126138da578081fd5b50813567ffffffffffffffff8111156138f1578182fd5b602083019150836020808302850101111561388957600080fd5b80516001600160701b038116811461182457600080fd5b803560ff8116811461182457600080fd5b600060208284031215613944578081fd5b813561394f816147d9565b9392505050565b600060208284031215613967578081fd5b815161394f816147d9565b60008060408385031215613984578081fd5b823561398f816147d9565b9150602083013561399f816147d9565b809150509250929050565b60008060008060008060008060006101208a8c0312156139c8578485fd5b89356139d3816147d9565b985060208a01356139e3816147d9565b975060408a01356139f3816147d9565b965060608a0135955060808a01359450613a0f60a08b01613922565b935060c08a0135925060e08a013591506101008a013590509295985092959850929598565b60008060008060808587031215613a49578384fd5b8435613a54816147d9565b93506020850135613a64816147d9565b93969395505050506040820135916060013590565b60008060408385031215613a8b578182fd5b8235613a96816147d9565b946020939093013593505050565b60008060008060008060008060008060006101408c8e031215613ac5578182fd5b8b35613ad0816147d9565b9a5060208c0135995060408c0135985060608c013567ffffffffffffffff811115613af9578283fd5b613b058e828f016138c9565b90995097505060808c0135613b19816147d9565b955060a08c01359450613b2e60c08d01613922565b935060e08c013592506101008c013591506101208c013590509295989b509295989b9093969950565b6000806000806000806000806080898b031215613b72578182fd5b883567ffffffffffffffff80821115613b89578384fd5b613b958c838d016138c9565b909a50985060208b0135915080821115613bad578384fd5b613bb98c838d016138c9565b909850965060408b0135915080821115613bd1578384fd5b613bdd8c838d016138c9565b909650945060608b0135915080821115613bf5578384fd5b50613c028b828c016138c9565b999c989b5096995094979396929594505050565b60008060008060008060008060008060a08b8d031215613c34578384fd5b8a3567ffffffffffffffff80821115613c4b578586fd5b613c578e838f016138c9565b909c509a5060208d0135915080821115613c6f578586fd5b613c7b8e838f016138c9565b909a50985060408d0135915080821115613c93578586fd5b613c9f8e838f016138c9565b909850965060608d0135915080821115613cb7578586fd5b613cc38e838f016138c9565b909650945060808d0135915080821115613cdb578384fd5b50613ce88d828e016138c9565b915080935050809150509295989b9194979a5092959850565b600060208284031215613d12578081fd5b8151801515811461394f578182fd5b600080600080600060a08688031215613d38578283fd5b853594506020860135613d4a816147d9565b93506040860135613d5a816147d9565b94979396509394606081013594506080013592915050565b60008060008060008060008060e0898b031215613d8d578182fd5b883597506020890135965060408901359550606089013567ffffffffffffffff811115613db8578283fd5b613dc48b828c016138c9565b9096509450506080890135613dd8816147d9565b979a969950949793969295929450505060a08201359160c0013590565b600080600060608486031215613e09578081fd5b613e128461390b565b9250613e206020850161390b565b9150604084015163ffffffff81168114613e38578182fd5b809150509250925092565b600060208284031215613e54578081fd5b5051919050565b60008060408385031215613e6d578182fd5b8235915060208084013567ffffffffffffffff80821115613e8c578384fd5b818601915086601f830112613e9f578384fd5b813581811115613eb157613eb16147c3565b83810260405185828201018181108582111715613ed057613ed06147c3565b604052828152858101935084860182860187018b1015613eee578788fd5b8795505b83861015613f1757613f03816138be565b855260019590950194938601938601613ef2565b508096505050505050509250929050565b600080600080600080600060c0888a031215613f42578081fd5b8735965060208801359550604088013567ffffffffffffffff811115613f66578182fd5b613f728a828b016138c9565b9096509450506060880135613f86816147d9565b969995985093969295946080840135945060a09093013592915050565b600080600060608486031215613fb7578081fd5b505081359360208301359350604090920135919050565b6bffffffffffffffffffffffff19606093841b811682529190921b16601482015260280190565b6000825161400781846020870161474f565b9190910192915050565b6001600160f81b0319815260609290921b6bffffffffffffffffffffffff1916600183015260158201527fe18a34eb0e04b04f7a0ac29a6e80748dca96319b42c54d679cb821dca90c6303603582015260550190565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b0397881681529590961660208601526040850193909352606084019190915260ff16608083015260a082015260c081019190915260e00190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b0393841681526020810192909252909116604082015260600190565b60a0808252810186905260008760c08301825b8981101561415f578235614142816147d9565b6001600160a01b031682526020928301929091019060010161412f565b50602084019790975250506040810193909352606083019190915260809091015292915050565b60a0808252865190820181905260009060209060c0840190828a01845b828110156141c85781516001600160a01b0316845292840192908401906001016141a3565b505050908301969096525060408101939093526060830191909152608090910152919050565b6020808252825182820181905260009190848201906040850190845b818110156142265783518352928401929184019160010161420a565b50909695505050505050565b901515815260200190565b6020808252602a908201526000805160206147f28339815191526040820152695f4c495155494449545960b01b606082015260800190565b60208082526015908201527464732d6d6174682d7375622d756e646572666c6f7760581b604082015260600190565b60208082526027908201527f53757368697377617056324c6962726172793a204944454e544943414c5f41446040820152664452455353455360c81b606082015260800190565b6020808252601d908201527f416e79737761705633526f757465723a20616464726573732830783029000000604082015260600190565b60208082526014908201527364732d6d6174682d6d756c2d6f766572666c6f7760601b604082015260600190565b6020808252602e908201526000805160206147f283398151915260408201526d17d3d55514155517d05353d5539560921b606082015260800190565b60208082526014908201527364732d6d6174682d6164642d6f766572666c6f7760601b604082015260600190565b6020808252602d908201527f5375736869737761705632526f757465723a20494e53554646494349454e545f60408201526c13d55514155517d05353d55395609a1b606082015260800190565b60208082526026908201527f5472616e7366657248656c7065723a204e41544956455f5452414e534645525f60408201526511905253115160d21b606082015260800190565b6020808252601a908201527f416e79737761705633526f757465723a20464f5242494444454e000000000000604082015260600190565b6020808252818101527f53757368697377617056324c6962726172793a205a45524f5f41444452455353604082015260600190565b6020808252601d908201527f416e79737761705633526f757465723a20494e56414c49445f50415448000000604082015260600190565b6020808252602b908201527f416e79737761705633526f757465723a20494e53554646494349454e545f4f5560408201526a1514155517d05353d5539560aa1b606082015260800190565b60208082526027908201526000805160206147f283398151915260408201526617d05353d5539560ca1b606082015260800190565b60208082526018908201527f416e79737761705633526f757465723a20455850495245440000000000000000604082015260600190565b6020808252602d908201526000805160206147f283398151915260408201526c17d25394155517d05353d55395609a1b606082015260800190565b60208082526024908201527f5472616e7366657248656c7065723a205452414e534645525f46524f4d5f46416040820152631253115160e21b606082015260800190565b6020808252818101527f53757368697377617056324c6962726172793a20494e56414c49445f50415448604082015260600190565b90815260200190565b9182526001600160a01b0316602082015260400190565b600085825284602083015260018060a01b03841660408301526080606083015282518060808401526146b48160a085016020870161474f565b601f01601f19169190910160a00195945050505050565b9283526020830191909152604082015260600190565b600082198211156146f4576146f46147ad565b500190565b60008261471457634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615614733576147336147ad565b500290565b60008282101561474a5761474a6147ad565b500390565b60005b8381101561476a578181015183820152602001614752565b83811115610b2b5750506000910152565b60008161478a5761478a6147ad565b506000190190565b60006000198214156147a6576147a66147ad565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146147ee57600080fd5b5056fe53757368697377617056324c6962726172793a20494e53554646494349454e54a2646970667358221220de1914e149706ef357f25f9b6190d0f0338da3e138dd279073255dca43af219464736f6c63430008000033",
}

// AnyswapV4RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use AnyswapV4RouterMetaData.ABI instead.
var AnyswapV4RouterABI = AnyswapV4RouterMetaData.ABI

// Deprecated: Use AnyswapV4RouterMetaData.Sigs instead.
// AnyswapV4RouterFuncSigs maps the 4-byte function signature to its string representation.
var AnyswapV4RouterFuncSigs = AnyswapV4RouterMetaData.Sigs

// AnyswapV4RouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AnyswapV4RouterMetaData.Bin instead.
var AnyswapV4RouterBin = AnyswapV4RouterMetaData.Bin

// DeployAnyswapV4Router deploys a new Ethereum contract, binding an instance of AnyswapV4Router to it.
func DeployAnyswapV4Router(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address, _wNATIVE common.Address, _mpc common.Address) (common.Address, *types.Transaction, *AnyswapV4Router, error) {
	parsed, err := AnyswapV4RouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AnyswapV4RouterBin), backend, _factory, _wNATIVE, _mpc)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AnyswapV4Router{AnyswapV4RouterCaller: AnyswapV4RouterCaller{contract: contract}, AnyswapV4RouterTransactor: AnyswapV4RouterTransactor{contract: contract}, AnyswapV4RouterFilterer: AnyswapV4RouterFilterer{contract: contract}}, nil
}

// AnyswapV4Router is an auto generated Go binding around an Ethereum contract.
type AnyswapV4Router struct {
	AnyswapV4RouterCaller     // Read-only binding to the contract
	AnyswapV4RouterTransactor // Write-only binding to the contract
	AnyswapV4RouterFilterer   // Log filterer for contract events
}

// AnyswapV4RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AnyswapV4RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnyswapV4RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AnyswapV4RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnyswapV4RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AnyswapV4RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnyswapV4RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnyswapV4RouterSession struct {
	Contract     *AnyswapV4Router  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnyswapV4RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnyswapV4RouterCallerSession struct {
	Contract *AnyswapV4RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AnyswapV4RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnyswapV4RouterTransactorSession struct {
	Contract     *AnyswapV4RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AnyswapV4RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AnyswapV4RouterRaw struct {
	Contract *AnyswapV4Router // Generic contract binding to access the raw methods on
}

// AnyswapV4RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnyswapV4RouterCallerRaw struct {
	Contract *AnyswapV4RouterCaller // Generic read-only contract binding to access the raw methods on
}

// AnyswapV4RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnyswapV4RouterTransactorRaw struct {
	Contract *AnyswapV4RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAnyswapV4Router creates a new instance of AnyswapV4Router, bound to a specific deployed contract.
func NewAnyswapV4Router(address common.Address, backend bind.ContractBackend) (*AnyswapV4Router, error) {
	contract, err := bindAnyswapV4Router(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AnyswapV4Router{AnyswapV4RouterCaller: AnyswapV4RouterCaller{contract: contract}, AnyswapV4RouterTransactor: AnyswapV4RouterTransactor{contract: contract}, AnyswapV4RouterFilterer: AnyswapV4RouterFilterer{contract: contract}}, nil
}

// NewAnyswapV4RouterCaller creates a new read-only instance of AnyswapV4Router, bound to a specific deployed contract.
func NewAnyswapV4RouterCaller(address common.Address, caller bind.ContractCaller) (*AnyswapV4RouterCaller, error) {
	contract, err := bindAnyswapV4Router(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AnyswapV4RouterCaller{contract: contract}, nil
}

// NewAnyswapV4RouterTransactor creates a new write-only instance of AnyswapV4Router, bound to a specific deployed contract.
func NewAnyswapV4RouterTransactor(address common.Address, transactor bind.ContractTransactor) (*AnyswapV4RouterTransactor, error) {
	contract, err := bindAnyswapV4Router(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AnyswapV4RouterTransactor{contract: contract}, nil
}

// NewAnyswapV4RouterFilterer creates a new log filterer instance of AnyswapV4Router, bound to a specific deployed contract.
func NewAnyswapV4RouterFilterer(address common.Address, filterer bind.ContractFilterer) (*AnyswapV4RouterFilterer, error) {
	contract, err := bindAnyswapV4Router(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AnyswapV4RouterFilterer{contract: contract}, nil
}

// bindAnyswapV4Router binds a generic wrapper to an already deployed contract.
func bindAnyswapV4Router(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AnyswapV4RouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnyswapV4Router *AnyswapV4RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AnyswapV4Router.Contract.AnyswapV4RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnyswapV4Router *AnyswapV4RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnyswapV4RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnyswapV4Router *AnyswapV4RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnyswapV4RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnyswapV4Router *AnyswapV4RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AnyswapV4Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnyswapV4Router *AnyswapV4RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnyswapV4Router *AnyswapV4RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.contract.Transact(opts, method, params...)
}

// CID is a free data retrieval call binding the contract method 0x99a2f2d7.
//
// Solidity: function cID() view returns(uint256 id)
func (_AnyswapV4Router *AnyswapV4RouterCaller) CID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnyswapV4Router.contract.Call(opts, &out, "cID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CID is a free data retrieval call binding the contract method 0x99a2f2d7.
//
// Solidity: function cID() view returns(uint256 id)
func (_AnyswapV4Router *AnyswapV4RouterSession) CID() (*big.Int, error) {
	return _AnyswapV4Router.Contract.CID(&_AnyswapV4Router.CallOpts)
}

// CID is a free data retrieval call binding the contract method 0x99a2f2d7.
//
// Solidity: function cID() view returns(uint256 id)
func (_AnyswapV4Router *AnyswapV4RouterCallerSession) CID() (*big.Int, error) {
	return _AnyswapV4Router.Contract.CID(&_AnyswapV4Router.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AnyswapV4Router *AnyswapV4RouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnyswapV4Router.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AnyswapV4Router *AnyswapV4RouterSession) Factory() (common.Address, error) {
	return _AnyswapV4Router.Contract.Factory(&_AnyswapV4Router.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AnyswapV4Router *AnyswapV4RouterCallerSession) Factory() (common.Address, error) {
	return _AnyswapV4Router.Contract.Factory(&_AnyswapV4Router.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_AnyswapV4Router *AnyswapV4RouterCaller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AnyswapV4Router.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_AnyswapV4Router *AnyswapV4RouterSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _AnyswapV4Router.Contract.GetAmountIn(&_AnyswapV4Router.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_AnyswapV4Router *AnyswapV4RouterCallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _AnyswapV4Router.Contract.GetAmountIn(&_AnyswapV4Router.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_AnyswapV4Router *AnyswapV4RouterCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AnyswapV4Router.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_AnyswapV4Router *AnyswapV4RouterSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _AnyswapV4Router.Contract.GetAmountOut(&_AnyswapV4Router.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_AnyswapV4Router *AnyswapV4RouterCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _AnyswapV4Router.Contract.GetAmountOut(&_AnyswapV4Router.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_AnyswapV4Router *AnyswapV4RouterCaller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _AnyswapV4Router.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_AnyswapV4Router *AnyswapV4RouterSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _AnyswapV4Router.Contract.GetAmountsIn(&_AnyswapV4Router.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_AnyswapV4Router *AnyswapV4RouterCallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _AnyswapV4Router.Contract.GetAmountsIn(&_AnyswapV4Router.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_AnyswapV4Router *AnyswapV4RouterCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _AnyswapV4Router.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_AnyswapV4Router *AnyswapV4RouterSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _AnyswapV4Router.Contract.GetAmountsOut(&_AnyswapV4Router.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_AnyswapV4Router *AnyswapV4RouterCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _AnyswapV4Router.Contract.GetAmountsOut(&_AnyswapV4Router.CallOpts, amountIn, path)
}

// Mpc is a free data retrieval call binding the contract method 0xf75c2664.
//
// Solidity: function mpc() view returns(address)
func (_AnyswapV4Router *AnyswapV4RouterCaller) Mpc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnyswapV4Router.contract.Call(opts, &out, "mpc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mpc is a free data retrieval call binding the contract method 0xf75c2664.
//
// Solidity: function mpc() view returns(address)
func (_AnyswapV4Router *AnyswapV4RouterSession) Mpc() (common.Address, error) {
	return _AnyswapV4Router.Contract.Mpc(&_AnyswapV4Router.CallOpts)
}

// Mpc is a free data retrieval call binding the contract method 0xf75c2664.
//
// Solidity: function mpc() view returns(address)
func (_AnyswapV4Router *AnyswapV4RouterCallerSession) Mpc() (common.Address, error) {
	return _AnyswapV4Router.Contract.Mpc(&_AnyswapV4Router.CallOpts)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_AnyswapV4Router *AnyswapV4RouterCaller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AnyswapV4Router.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_AnyswapV4Router *AnyswapV4RouterSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _AnyswapV4Router.Contract.Quote(&_AnyswapV4Router.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_AnyswapV4Router *AnyswapV4RouterCallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _AnyswapV4Router.Contract.Quote(&_AnyswapV4Router.CallOpts, amountA, reserveA, reserveB)
}

// WNATIVE is a free data retrieval call binding the contract method 0x8fd903f5.
//
// Solidity: function wNATIVE() view returns(address)
func (_AnyswapV4Router *AnyswapV4RouterCaller) WNATIVE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnyswapV4Router.contract.Call(opts, &out, "wNATIVE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WNATIVE is a free data retrieval call binding the contract method 0x8fd903f5.
//
// Solidity: function wNATIVE() view returns(address)
func (_AnyswapV4Router *AnyswapV4RouterSession) WNATIVE() (common.Address, error) {
	return _AnyswapV4Router.Contract.WNATIVE(&_AnyswapV4Router.CallOpts)
}

// WNATIVE is a free data retrieval call binding the contract method 0x8fd903f5.
//
// Solidity: function wNATIVE() view returns(address)
func (_AnyswapV4Router *AnyswapV4RouterCallerSession) WNATIVE() (common.Address, error) {
	return _AnyswapV4Router.Contract.WNATIVE(&_AnyswapV4Router.CallOpts)
}

// AnySwapFeeTo is a paid mutator transaction binding the contract method 0x87cc6e2f.
//
// Solidity: function anySwapFeeTo(address token, uint256 amount) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapFeeTo(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapFeeTo", token, amount)
}

// AnySwapFeeTo is a paid mutator transaction binding the contract method 0x87cc6e2f.
//
// Solidity: function anySwapFeeTo(address token, uint256 amount) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapFeeTo(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapFeeTo(&_AnyswapV4Router.TransactOpts, token, amount)
}

// AnySwapFeeTo is a paid mutator transaction binding the contract method 0x87cc6e2f.
//
// Solidity: function anySwapFeeTo(address token, uint256 amount) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapFeeTo(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapFeeTo(&_AnyswapV4Router.TransactOpts, token, amount)
}

// AnySwapIn is a paid mutator transaction binding the contract method 0x25121b76.
//
// Solidity: function anySwapIn(bytes32[] txs, address[] tokens, address[] to, uint256[] amounts, uint256[] fromChainIDs) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapIn(opts *bind.TransactOpts, txs [][32]byte, tokens []common.Address, to []common.Address, amounts []*big.Int, fromChainIDs []*big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapIn", txs, tokens, to, amounts, fromChainIDs)
}

// AnySwapIn is a paid mutator transaction binding the contract method 0x25121b76.
//
// Solidity: function anySwapIn(bytes32[] txs, address[] tokens, address[] to, uint256[] amounts, uint256[] fromChainIDs) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapIn(txs [][32]byte, tokens []common.Address, to []common.Address, amounts []*big.Int, fromChainIDs []*big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapIn(&_AnyswapV4Router.TransactOpts, txs, tokens, to, amounts, fromChainIDs)
}

// AnySwapIn is a paid mutator transaction binding the contract method 0x25121b76.
//
// Solidity: function anySwapIn(bytes32[] txs, address[] tokens, address[] to, uint256[] amounts, uint256[] fromChainIDs) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapIn(txs [][32]byte, tokens []common.Address, to []common.Address, amounts []*big.Int, fromChainIDs []*big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapIn(&_AnyswapV4Router.TransactOpts, txs, tokens, to, amounts, fromChainIDs)
}

// AnySwapIn0 is a paid mutator transaction binding the contract method 0x825bb13c.
//
// Solidity: function anySwapIn(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapIn0(opts *bind.TransactOpts, txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapIn0", txs, token, to, amount, fromChainID)
}

// AnySwapIn0 is a paid mutator transaction binding the contract method 0x825bb13c.
//
// Solidity: function anySwapIn(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapIn0(txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapIn0(&_AnyswapV4Router.TransactOpts, txs, token, to, amount, fromChainID)
}

// AnySwapIn0 is a paid mutator transaction binding the contract method 0x825bb13c.
//
// Solidity: function anySwapIn(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapIn0(txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapIn0(&_AnyswapV4Router.TransactOpts, txs, token, to, amount, fromChainID)
}

// AnySwapInAuto is a paid mutator transaction binding the contract method 0x0175b1c4.
//
// Solidity: function anySwapInAuto(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapInAuto(opts *bind.TransactOpts, txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapInAuto", txs, token, to, amount, fromChainID)
}

// AnySwapInAuto is a paid mutator transaction binding the contract method 0x0175b1c4.
//
// Solidity: function anySwapInAuto(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapInAuto(txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapInAuto(&_AnyswapV4Router.TransactOpts, txs, token, to, amount, fromChainID)
}

// AnySwapInAuto is a paid mutator transaction binding the contract method 0x0175b1c4.
//
// Solidity: function anySwapInAuto(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapInAuto(txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapInAuto(&_AnyswapV4Router.TransactOpts, txs, token, to, amount, fromChainID)
}

// AnySwapInExactTokensForNative is a paid mutator transaction binding the contract method 0x52a397d5.
//
// Solidity: function anySwapInExactTokensForNative(bytes32 txs, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 fromChainID) returns(uint256[] amounts)
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapInExactTokensForNative(opts *bind.TransactOpts, txs [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapInExactTokensForNative", txs, amountIn, amountOutMin, path, to, deadline, fromChainID)
}

// AnySwapInExactTokensForNative is a paid mutator transaction binding the contract method 0x52a397d5.
//
// Solidity: function anySwapInExactTokensForNative(bytes32 txs, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 fromChainID) returns(uint256[] amounts)
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapInExactTokensForNative(txs [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapInExactTokensForNative(&_AnyswapV4Router.TransactOpts, txs, amountIn, amountOutMin, path, to, deadline, fromChainID)
}

// AnySwapInExactTokensForNative is a paid mutator transaction binding the contract method 0x52a397d5.
//
// Solidity: function anySwapInExactTokensForNative(bytes32 txs, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 fromChainID) returns(uint256[] amounts)
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapInExactTokensForNative(txs [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapInExactTokensForNative(&_AnyswapV4Router.TransactOpts, txs, amountIn, amountOutMin, path, to, deadline, fromChainID)
}

// AnySwapInExactTokensForTokens is a paid mutator transaction binding the contract method 0x2fc1e728.
//
// Solidity: function anySwapInExactTokensForTokens(bytes32 txs, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 fromChainID) returns(uint256[] amounts)
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapInExactTokensForTokens(opts *bind.TransactOpts, txs [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapInExactTokensForTokens", txs, amountIn, amountOutMin, path, to, deadline, fromChainID)
}

// AnySwapInExactTokensForTokens is a paid mutator transaction binding the contract method 0x2fc1e728.
//
// Solidity: function anySwapInExactTokensForTokens(bytes32 txs, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 fromChainID) returns(uint256[] amounts)
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapInExactTokensForTokens(txs [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapInExactTokensForTokens(&_AnyswapV4Router.TransactOpts, txs, amountIn, amountOutMin, path, to, deadline, fromChainID)
}

// AnySwapInExactTokensForTokens is a paid mutator transaction binding the contract method 0x2fc1e728.
//
// Solidity: function anySwapInExactTokensForTokens(bytes32 txs, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 fromChainID) returns(uint256[] amounts)
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapInExactTokensForTokens(txs [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapInExactTokensForTokens(&_AnyswapV4Router.TransactOpts, txs, amountIn, amountOutMin, path, to, deadline, fromChainID)
}

// AnySwapInUnderlying is a paid mutator transaction binding the contract method 0x3f88de89.
//
// Solidity: function anySwapInUnderlying(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapInUnderlying(opts *bind.TransactOpts, txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapInUnderlying", txs, token, to, amount, fromChainID)
}

// AnySwapInUnderlying is a paid mutator transaction binding the contract method 0x3f88de89.
//
// Solidity: function anySwapInUnderlying(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapInUnderlying(txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapInUnderlying(&_AnyswapV4Router.TransactOpts, txs, token, to, amount, fromChainID)
}

// AnySwapInUnderlying is a paid mutator transaction binding the contract method 0x3f88de89.
//
// Solidity: function anySwapInUnderlying(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapInUnderlying(txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapInUnderlying(&_AnyswapV4Router.TransactOpts, txs, token, to, amount, fromChainID)
}

// AnySwapOut is a paid mutator transaction binding the contract method 0x241dc2df.
//
// Solidity: function anySwapOut(address token, address to, uint256 amount, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapOut(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapOut", token, to, amount, toChainID)
}

// AnySwapOut is a paid mutator transaction binding the contract method 0x241dc2df.
//
// Solidity: function anySwapOut(address token, address to, uint256 amount, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapOut(token common.Address, to common.Address, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOut(&_AnyswapV4Router.TransactOpts, token, to, amount, toChainID)
}

// AnySwapOut is a paid mutator transaction binding the contract method 0x241dc2df.
//
// Solidity: function anySwapOut(address token, address to, uint256 amount, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapOut(token common.Address, to common.Address, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOut(&_AnyswapV4Router.TransactOpts, token, to, amount, toChainID)
}

// AnySwapOut0 is a paid mutator transaction binding the contract method 0xdcfb77b1.
//
// Solidity: function anySwapOut(address[] tokens, address[] to, uint256[] amounts, uint256[] toChainIDs) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapOut0(opts *bind.TransactOpts, tokens []common.Address, to []common.Address, amounts []*big.Int, toChainIDs []*big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapOut0", tokens, to, amounts, toChainIDs)
}

// AnySwapOut0 is a paid mutator transaction binding the contract method 0xdcfb77b1.
//
// Solidity: function anySwapOut(address[] tokens, address[] to, uint256[] amounts, uint256[] toChainIDs) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapOut0(tokens []common.Address, to []common.Address, amounts []*big.Int, toChainIDs []*big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOut0(&_AnyswapV4Router.TransactOpts, tokens, to, amounts, toChainIDs)
}

// AnySwapOut0 is a paid mutator transaction binding the contract method 0xdcfb77b1.
//
// Solidity: function anySwapOut(address[] tokens, address[] to, uint256[] amounts, uint256[] toChainIDs) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapOut0(tokens []common.Address, to []common.Address, amounts []*big.Int, toChainIDs []*big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOut0(&_AnyswapV4Router.TransactOpts, tokens, to, amounts, toChainIDs)
}

// AnySwapOutExactTokensForNative is a paid mutator transaction binding the contract method 0x65782f56.
//
// Solidity: function anySwapOutExactTokensForNative(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapOutExactTokensForNative(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapOutExactTokensForNative", amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForNative is a paid mutator transaction binding the contract method 0x65782f56.
//
// Solidity: function anySwapOutExactTokensForNative(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapOutExactTokensForNative(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForNative(&_AnyswapV4Router.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForNative is a paid mutator transaction binding the contract method 0x65782f56.
//
// Solidity: function anySwapOutExactTokensForNative(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapOutExactTokensForNative(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForNative(&_AnyswapV4Router.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlying is a paid mutator transaction binding the contract method 0x6a453972.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlying(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapOutExactTokensForNativeUnderlying(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapOutExactTokensForNativeUnderlying", amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlying is a paid mutator transaction binding the contract method 0x6a453972.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlying(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapOutExactTokensForNativeUnderlying(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForNativeUnderlying(&_AnyswapV4Router.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlying is a paid mutator transaction binding the contract method 0x6a453972.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlying(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapOutExactTokensForNativeUnderlying(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForNativeUnderlying(&_AnyswapV4Router.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x4d93bb94.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlyingWithPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapOutExactTokensForNativeUnderlyingWithPermit(opts *bind.TransactOpts, from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapOutExactTokensForNativeUnderlyingWithPermit", from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x4d93bb94.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlyingWithPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapOutExactTokensForNativeUnderlyingWithPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForNativeUnderlyingWithPermit(&_AnyswapV4Router.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x4d93bb94.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlyingWithPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapOutExactTokensForNativeUnderlyingWithPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForNativeUnderlyingWithPermit(&_AnyswapV4Router.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0xc8e174f6.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlyingWithTransferPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit(opts *bind.TransactOpts, from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapOutExactTokensForNativeUnderlyingWithTransferPermit", from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0xc8e174f6.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlyingWithTransferPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit(&_AnyswapV4Router.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0xc8e174f6.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlyingWithTransferPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit(&_AnyswapV4Router.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForTokens is a paid mutator transaction binding the contract method 0x0bb57203.
//
// Solidity: function anySwapOutExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapOutExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapOutExactTokensForTokens", amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForTokens is a paid mutator transaction binding the contract method 0x0bb57203.
//
// Solidity: function anySwapOutExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapOutExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForTokens(&_AnyswapV4Router.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForTokens is a paid mutator transaction binding the contract method 0x0bb57203.
//
// Solidity: function anySwapOutExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapOutExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForTokens(&_AnyswapV4Router.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlying is a paid mutator transaction binding the contract method 0xd8b9f610.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlying(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapOutExactTokensForTokensUnderlying(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapOutExactTokensForTokensUnderlying", amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlying is a paid mutator transaction binding the contract method 0xd8b9f610.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlying(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapOutExactTokensForTokensUnderlying(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForTokensUnderlying(&_AnyswapV4Router.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlying is a paid mutator transaction binding the contract method 0xd8b9f610.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlying(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapOutExactTokensForTokensUnderlying(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForTokensUnderlying(&_AnyswapV4Router.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x99cd84b5.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlyingWithPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapOutExactTokensForTokensUnderlyingWithPermit(opts *bind.TransactOpts, from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapOutExactTokensForTokensUnderlyingWithPermit", from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x99cd84b5.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlyingWithPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapOutExactTokensForTokensUnderlyingWithPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForTokensUnderlyingWithPermit(&_AnyswapV4Router.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x99cd84b5.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlyingWithPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapOutExactTokensForTokensUnderlyingWithPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForTokensUnderlyingWithPermit(&_AnyswapV4Router.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0x9aa1ac61.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlyingWithTransferPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit(opts *bind.TransactOpts, from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapOutExactTokensForTokensUnderlyingWithTransferPermit", from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0x9aa1ac61.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlyingWithTransferPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit(&_AnyswapV4Router.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0x9aa1ac61.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlyingWithTransferPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit(&_AnyswapV4Router.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutUnderlying is a paid mutator transaction binding the contract method 0xedbdf5e2.
//
// Solidity: function anySwapOutUnderlying(address token, address to, uint256 amount, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapOutUnderlying(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapOutUnderlying", token, to, amount, toChainID)
}

// AnySwapOutUnderlying is a paid mutator transaction binding the contract method 0xedbdf5e2.
//
// Solidity: function anySwapOutUnderlying(address token, address to, uint256 amount, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapOutUnderlying(token common.Address, to common.Address, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutUnderlying(&_AnyswapV4Router.TransactOpts, token, to, amount, toChainID)
}

// AnySwapOutUnderlying is a paid mutator transaction binding the contract method 0xedbdf5e2.
//
// Solidity: function anySwapOutUnderlying(address token, address to, uint256 amount, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapOutUnderlying(token common.Address, to common.Address, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutUnderlying(&_AnyswapV4Router.TransactOpts, token, to, amount, toChainID)
}

// AnySwapOutUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x8d7d3eea.
//
// Solidity: function anySwapOutUnderlyingWithPermit(address from, address token, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapOutUnderlyingWithPermit(opts *bind.TransactOpts, from common.Address, token common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapOutUnderlyingWithPermit", from, token, to, amount, deadline, v, r, s, toChainID)
}

// AnySwapOutUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x8d7d3eea.
//
// Solidity: function anySwapOutUnderlyingWithPermit(address from, address token, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapOutUnderlyingWithPermit(from common.Address, token common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutUnderlyingWithPermit(&_AnyswapV4Router.TransactOpts, from, token, to, amount, deadline, v, r, s, toChainID)
}

// AnySwapOutUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x8d7d3eea.
//
// Solidity: function anySwapOutUnderlyingWithPermit(address from, address token, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapOutUnderlyingWithPermit(from common.Address, token common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutUnderlyingWithPermit(&_AnyswapV4Router.TransactOpts, from, token, to, amount, deadline, v, r, s, toChainID)
}

// AnySwapOutUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0x1b91a934.
//
// Solidity: function anySwapOutUnderlyingWithTransferPermit(address from, address token, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) AnySwapOutUnderlyingWithTransferPermit(opts *bind.TransactOpts, from common.Address, token common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "anySwapOutUnderlyingWithTransferPermit", from, token, to, amount, deadline, v, r, s, toChainID)
}

// AnySwapOutUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0x1b91a934.
//
// Solidity: function anySwapOutUnderlyingWithTransferPermit(address from, address token, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) AnySwapOutUnderlyingWithTransferPermit(from common.Address, token common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutUnderlyingWithTransferPermit(&_AnyswapV4Router.TransactOpts, from, token, to, amount, deadline, v, r, s, toChainID)
}

// AnySwapOutUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0x1b91a934.
//
// Solidity: function anySwapOutUnderlyingWithTransferPermit(address from, address token, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) AnySwapOutUnderlyingWithTransferPermit(from common.Address, token common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.AnySwapOutUnderlyingWithTransferPermit(&_AnyswapV4Router.TransactOpts, from, token, to, amount, deadline, v, r, s, toChainID)
}

// ChangeMPC is a paid mutator transaction binding the contract method 0x5b7b018c.
//
// Solidity: function changeMPC(address newMPC) returns(bool)
func (_AnyswapV4Router *AnyswapV4RouterTransactor) ChangeMPC(opts *bind.TransactOpts, newMPC common.Address) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "changeMPC", newMPC)
}

// ChangeMPC is a paid mutator transaction binding the contract method 0x5b7b018c.
//
// Solidity: function changeMPC(address newMPC) returns(bool)
func (_AnyswapV4Router *AnyswapV4RouterSession) ChangeMPC(newMPC common.Address) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.ChangeMPC(&_AnyswapV4Router.TransactOpts, newMPC)
}

// ChangeMPC is a paid mutator transaction binding the contract method 0x5b7b018c.
//
// Solidity: function changeMPC(address newMPC) returns(bool)
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) ChangeMPC(newMPC common.Address) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.ChangeMPC(&_AnyswapV4Router.TransactOpts, newMPC)
}

// ChangeVault is a paid mutator transaction binding the contract method 0x456862aa.
//
// Solidity: function changeVault(address token, address newVault) returns(bool)
func (_AnyswapV4Router *AnyswapV4RouterTransactor) ChangeVault(opts *bind.TransactOpts, token common.Address, newVault common.Address) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.Transact(opts, "changeVault", token, newVault)
}

// ChangeVault is a paid mutator transaction binding the contract method 0x456862aa.
//
// Solidity: function changeVault(address token, address newVault) returns(bool)
func (_AnyswapV4Router *AnyswapV4RouterSession) ChangeVault(token common.Address, newVault common.Address) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.ChangeVault(&_AnyswapV4Router.TransactOpts, token, newVault)
}

// ChangeVault is a paid mutator transaction binding the contract method 0x456862aa.
//
// Solidity: function changeVault(address token, address newVault) returns(bool)
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) ChangeVault(token common.Address, newVault common.Address) (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.ChangeVault(&_AnyswapV4Router.TransactOpts, token, newVault)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnyswapV4Router.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AnyswapV4Router *AnyswapV4RouterSession) Receive() (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.Receive(&_AnyswapV4Router.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AnyswapV4Router *AnyswapV4RouterTransactorSession) Receive() (*types.Transaction, error) {
	return _AnyswapV4Router.Contract.Receive(&_AnyswapV4Router.TransactOpts)
}

// AnyswapV4RouterLogAnySwapInIterator is returned from FilterLogAnySwapIn and is used to iterate over the raw logs and unpacked data for LogAnySwapIn events raised by the AnyswapV4Router contract.
type AnyswapV4RouterLogAnySwapInIterator struct {
	Event *AnyswapV4RouterLogAnySwapIn // Event containing the contract specifics and raw log

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
func (it *AnyswapV4RouterLogAnySwapInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnyswapV4RouterLogAnySwapIn)
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
		it.Event = new(AnyswapV4RouterLogAnySwapIn)
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
func (it *AnyswapV4RouterLogAnySwapInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnyswapV4RouterLogAnySwapInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnyswapV4RouterLogAnySwapIn represents a LogAnySwapIn event raised by the AnyswapV4Router contract.
type AnyswapV4RouterLogAnySwapIn struct {
	Txhash      [32]byte
	Token       common.Address
	To          common.Address
	Amount      *big.Int
	FromChainID *big.Int
	ToChainID   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogAnySwapIn is a free log retrieval operation binding the contract event 0xaac9ce45fe3adf5143598c4f18a369591a20a3384aedaf1b525d29127e1fcd55.
//
// Solidity: event LogAnySwapIn(bytes32 indexed txhash, address indexed token, address indexed to, uint256 amount, uint256 fromChainID, uint256 toChainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) FilterLogAnySwapIn(opts *bind.FilterOpts, txhash [][32]byte, token []common.Address, to []common.Address) (*AnyswapV4RouterLogAnySwapInIterator, error) {

	var txhashRule []interface{}
	for _, txhashItem := range txhash {
		txhashRule = append(txhashRule, txhashItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AnyswapV4Router.contract.FilterLogs(opts, "LogAnySwapIn", txhashRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AnyswapV4RouterLogAnySwapInIterator{contract: _AnyswapV4Router.contract, event: "LogAnySwapIn", logs: logs, sub: sub}, nil
}

// WatchLogAnySwapIn is a free log subscription operation binding the contract event 0xaac9ce45fe3adf5143598c4f18a369591a20a3384aedaf1b525d29127e1fcd55.
//
// Solidity: event LogAnySwapIn(bytes32 indexed txhash, address indexed token, address indexed to, uint256 amount, uint256 fromChainID, uint256 toChainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) WatchLogAnySwapIn(opts *bind.WatchOpts, sink chan<- *AnyswapV4RouterLogAnySwapIn, txhash [][32]byte, token []common.Address, to []common.Address) (event.Subscription, error) {

	var txhashRule []interface{}
	for _, txhashItem := range txhash {
		txhashRule = append(txhashRule, txhashItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AnyswapV4Router.contract.WatchLogs(opts, "LogAnySwapIn", txhashRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnyswapV4RouterLogAnySwapIn)
				if err := _AnyswapV4Router.contract.UnpackLog(event, "LogAnySwapIn", log); err != nil {
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

// ParseLogAnySwapIn is a log parse operation binding the contract event 0xaac9ce45fe3adf5143598c4f18a369591a20a3384aedaf1b525d29127e1fcd55.
//
// Solidity: event LogAnySwapIn(bytes32 indexed txhash, address indexed token, address indexed to, uint256 amount, uint256 fromChainID, uint256 toChainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) ParseLogAnySwapIn(log types.Log) (*AnyswapV4RouterLogAnySwapIn, error) {
	event := new(AnyswapV4RouterLogAnySwapIn)
	if err := _AnyswapV4Router.contract.UnpackLog(event, "LogAnySwapIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnyswapV4RouterLogAnySwapOutIterator is returned from FilterLogAnySwapOut and is used to iterate over the raw logs and unpacked data for LogAnySwapOut events raised by the AnyswapV4Router contract.
type AnyswapV4RouterLogAnySwapOutIterator struct {
	Event *AnyswapV4RouterLogAnySwapOut // Event containing the contract specifics and raw log

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
func (it *AnyswapV4RouterLogAnySwapOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnyswapV4RouterLogAnySwapOut)
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
		it.Event = new(AnyswapV4RouterLogAnySwapOut)
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
func (it *AnyswapV4RouterLogAnySwapOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnyswapV4RouterLogAnySwapOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnyswapV4RouterLogAnySwapOut represents a LogAnySwapOut event raised by the AnyswapV4Router contract.
type AnyswapV4RouterLogAnySwapOut struct {
	Token       common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	FromChainID *big.Int
	ToChainID   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogAnySwapOut is a free log retrieval operation binding the contract event 0x97116cf6cd4f6412bb47914d6db18da9e16ab2142f543b86e207c24fbd16b23a.
//
// Solidity: event LogAnySwapOut(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 fromChainID, uint256 toChainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) FilterLogAnySwapOut(opts *bind.FilterOpts, token []common.Address, from []common.Address, to []common.Address) (*AnyswapV4RouterLogAnySwapOutIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AnyswapV4Router.contract.FilterLogs(opts, "LogAnySwapOut", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AnyswapV4RouterLogAnySwapOutIterator{contract: _AnyswapV4Router.contract, event: "LogAnySwapOut", logs: logs, sub: sub}, nil
}

// WatchLogAnySwapOut is a free log subscription operation binding the contract event 0x97116cf6cd4f6412bb47914d6db18da9e16ab2142f543b86e207c24fbd16b23a.
//
// Solidity: event LogAnySwapOut(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 fromChainID, uint256 toChainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) WatchLogAnySwapOut(opts *bind.WatchOpts, sink chan<- *AnyswapV4RouterLogAnySwapOut, token []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AnyswapV4Router.contract.WatchLogs(opts, "LogAnySwapOut", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnyswapV4RouterLogAnySwapOut)
				if err := _AnyswapV4Router.contract.UnpackLog(event, "LogAnySwapOut", log); err != nil {
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

// ParseLogAnySwapOut is a log parse operation binding the contract event 0x97116cf6cd4f6412bb47914d6db18da9e16ab2142f543b86e207c24fbd16b23a.
//
// Solidity: event LogAnySwapOut(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 fromChainID, uint256 toChainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) ParseLogAnySwapOut(log types.Log) (*AnyswapV4RouterLogAnySwapOut, error) {
	event := new(AnyswapV4RouterLogAnySwapOut)
	if err := _AnyswapV4Router.contract.UnpackLog(event, "LogAnySwapOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnyswapV4RouterLogAnySwapTradeTokensForNativeIterator is returned from FilterLogAnySwapTradeTokensForNative and is used to iterate over the raw logs and unpacked data for LogAnySwapTradeTokensForNative events raised by the AnyswapV4Router contract.
type AnyswapV4RouterLogAnySwapTradeTokensForNativeIterator struct {
	Event *AnyswapV4RouterLogAnySwapTradeTokensForNative // Event containing the contract specifics and raw log

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
func (it *AnyswapV4RouterLogAnySwapTradeTokensForNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnyswapV4RouterLogAnySwapTradeTokensForNative)
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
		it.Event = new(AnyswapV4RouterLogAnySwapTradeTokensForNative)
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
func (it *AnyswapV4RouterLogAnySwapTradeTokensForNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnyswapV4RouterLogAnySwapTradeTokensForNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnyswapV4RouterLogAnySwapTradeTokensForNative represents a LogAnySwapTradeTokensForNative event raised by the AnyswapV4Router contract.
type AnyswapV4RouterLogAnySwapTradeTokensForNative struct {
	Path         []common.Address
	From         common.Address
	To           common.Address
	AmountIn     *big.Int
	AmountOutMin *big.Int
	FromChainID  *big.Int
	ToChainID    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogAnySwapTradeTokensForNative is a free log retrieval operation binding the contract event 0x278277e0209c347189add7bd92411973b5f6b8644f7ac62ea1be984ce993f8f4.
//
// Solidity: event LogAnySwapTradeTokensForNative(address[] path, address indexed from, address indexed to, uint256 amountIn, uint256 amountOutMin, uint256 fromChainID, uint256 toChainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) FilterLogAnySwapTradeTokensForNative(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AnyswapV4RouterLogAnySwapTradeTokensForNativeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AnyswapV4Router.contract.FilterLogs(opts, "LogAnySwapTradeTokensForNative", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AnyswapV4RouterLogAnySwapTradeTokensForNativeIterator{contract: _AnyswapV4Router.contract, event: "LogAnySwapTradeTokensForNative", logs: logs, sub: sub}, nil
}

// WatchLogAnySwapTradeTokensForNative is a free log subscription operation binding the contract event 0x278277e0209c347189add7bd92411973b5f6b8644f7ac62ea1be984ce993f8f4.
//
// Solidity: event LogAnySwapTradeTokensForNative(address[] path, address indexed from, address indexed to, uint256 amountIn, uint256 amountOutMin, uint256 fromChainID, uint256 toChainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) WatchLogAnySwapTradeTokensForNative(opts *bind.WatchOpts, sink chan<- *AnyswapV4RouterLogAnySwapTradeTokensForNative, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AnyswapV4Router.contract.WatchLogs(opts, "LogAnySwapTradeTokensForNative", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnyswapV4RouterLogAnySwapTradeTokensForNative)
				if err := _AnyswapV4Router.contract.UnpackLog(event, "LogAnySwapTradeTokensForNative", log); err != nil {
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

// ParseLogAnySwapTradeTokensForNative is a log parse operation binding the contract event 0x278277e0209c347189add7bd92411973b5f6b8644f7ac62ea1be984ce993f8f4.
//
// Solidity: event LogAnySwapTradeTokensForNative(address[] path, address indexed from, address indexed to, uint256 amountIn, uint256 amountOutMin, uint256 fromChainID, uint256 toChainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) ParseLogAnySwapTradeTokensForNative(log types.Log) (*AnyswapV4RouterLogAnySwapTradeTokensForNative, error) {
	event := new(AnyswapV4RouterLogAnySwapTradeTokensForNative)
	if err := _AnyswapV4Router.contract.UnpackLog(event, "LogAnySwapTradeTokensForNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnyswapV4RouterLogAnySwapTradeTokensForTokensIterator is returned from FilterLogAnySwapTradeTokensForTokens and is used to iterate over the raw logs and unpacked data for LogAnySwapTradeTokensForTokens events raised by the AnyswapV4Router contract.
type AnyswapV4RouterLogAnySwapTradeTokensForTokensIterator struct {
	Event *AnyswapV4RouterLogAnySwapTradeTokensForTokens // Event containing the contract specifics and raw log

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
func (it *AnyswapV4RouterLogAnySwapTradeTokensForTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnyswapV4RouterLogAnySwapTradeTokensForTokens)
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
		it.Event = new(AnyswapV4RouterLogAnySwapTradeTokensForTokens)
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
func (it *AnyswapV4RouterLogAnySwapTradeTokensForTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnyswapV4RouterLogAnySwapTradeTokensForTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnyswapV4RouterLogAnySwapTradeTokensForTokens represents a LogAnySwapTradeTokensForTokens event raised by the AnyswapV4Router contract.
type AnyswapV4RouterLogAnySwapTradeTokensForTokens struct {
	Path         []common.Address
	From         common.Address
	To           common.Address
	AmountIn     *big.Int
	AmountOutMin *big.Int
	FromChainID  *big.Int
	ToChainID    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogAnySwapTradeTokensForTokens is a free log retrieval operation binding the contract event 0xfea6abdf4fd32f20966dff7619354cd82cd43dc78a3bee479f04c74dbfc585b3.
//
// Solidity: event LogAnySwapTradeTokensForTokens(address[] path, address indexed from, address indexed to, uint256 amountIn, uint256 amountOutMin, uint256 fromChainID, uint256 toChainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) FilterLogAnySwapTradeTokensForTokens(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AnyswapV4RouterLogAnySwapTradeTokensForTokensIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AnyswapV4Router.contract.FilterLogs(opts, "LogAnySwapTradeTokensForTokens", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AnyswapV4RouterLogAnySwapTradeTokensForTokensIterator{contract: _AnyswapV4Router.contract, event: "LogAnySwapTradeTokensForTokens", logs: logs, sub: sub}, nil
}

// WatchLogAnySwapTradeTokensForTokens is a free log subscription operation binding the contract event 0xfea6abdf4fd32f20966dff7619354cd82cd43dc78a3bee479f04c74dbfc585b3.
//
// Solidity: event LogAnySwapTradeTokensForTokens(address[] path, address indexed from, address indexed to, uint256 amountIn, uint256 amountOutMin, uint256 fromChainID, uint256 toChainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) WatchLogAnySwapTradeTokensForTokens(opts *bind.WatchOpts, sink chan<- *AnyswapV4RouterLogAnySwapTradeTokensForTokens, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AnyswapV4Router.contract.WatchLogs(opts, "LogAnySwapTradeTokensForTokens", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnyswapV4RouterLogAnySwapTradeTokensForTokens)
				if err := _AnyswapV4Router.contract.UnpackLog(event, "LogAnySwapTradeTokensForTokens", log); err != nil {
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

// ParseLogAnySwapTradeTokensForTokens is a log parse operation binding the contract event 0xfea6abdf4fd32f20966dff7619354cd82cd43dc78a3bee479f04c74dbfc585b3.
//
// Solidity: event LogAnySwapTradeTokensForTokens(address[] path, address indexed from, address indexed to, uint256 amountIn, uint256 amountOutMin, uint256 fromChainID, uint256 toChainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) ParseLogAnySwapTradeTokensForTokens(log types.Log) (*AnyswapV4RouterLogAnySwapTradeTokensForTokens, error) {
	event := new(AnyswapV4RouterLogAnySwapTradeTokensForTokens)
	if err := _AnyswapV4Router.contract.UnpackLog(event, "LogAnySwapTradeTokensForTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnyswapV4RouterLogChangeMPCIterator is returned from FilterLogChangeMPC and is used to iterate over the raw logs and unpacked data for LogChangeMPC events raised by the AnyswapV4Router contract.
type AnyswapV4RouterLogChangeMPCIterator struct {
	Event *AnyswapV4RouterLogChangeMPC // Event containing the contract specifics and raw log

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
func (it *AnyswapV4RouterLogChangeMPCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnyswapV4RouterLogChangeMPC)
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
		it.Event = new(AnyswapV4RouterLogChangeMPC)
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
func (it *AnyswapV4RouterLogChangeMPCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnyswapV4RouterLogChangeMPCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnyswapV4RouterLogChangeMPC represents a LogChangeMPC event raised by the AnyswapV4Router contract.
type AnyswapV4RouterLogChangeMPC struct {
	OldMPC        common.Address
	NewMPC        common.Address
	EffectiveTime *big.Int
	ChainID       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLogChangeMPC is a free log retrieval operation binding the contract event 0xcda32bc39904597666dfa9f9c845714756e1ffffad55b52e0d344673a2198121.
//
// Solidity: event LogChangeMPC(address indexed oldMPC, address indexed newMPC, uint256 indexed effectiveTime, uint256 chainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) FilterLogChangeMPC(opts *bind.FilterOpts, oldMPC []common.Address, newMPC []common.Address, effectiveTime []*big.Int) (*AnyswapV4RouterLogChangeMPCIterator, error) {

	var oldMPCRule []interface{}
	for _, oldMPCItem := range oldMPC {
		oldMPCRule = append(oldMPCRule, oldMPCItem)
	}
	var newMPCRule []interface{}
	for _, newMPCItem := range newMPC {
		newMPCRule = append(newMPCRule, newMPCItem)
	}
	var effectiveTimeRule []interface{}
	for _, effectiveTimeItem := range effectiveTime {
		effectiveTimeRule = append(effectiveTimeRule, effectiveTimeItem)
	}

	logs, sub, err := _AnyswapV4Router.contract.FilterLogs(opts, "LogChangeMPC", oldMPCRule, newMPCRule, effectiveTimeRule)
	if err != nil {
		return nil, err
	}
	return &AnyswapV4RouterLogChangeMPCIterator{contract: _AnyswapV4Router.contract, event: "LogChangeMPC", logs: logs, sub: sub}, nil
}

// WatchLogChangeMPC is a free log subscription operation binding the contract event 0xcda32bc39904597666dfa9f9c845714756e1ffffad55b52e0d344673a2198121.
//
// Solidity: event LogChangeMPC(address indexed oldMPC, address indexed newMPC, uint256 indexed effectiveTime, uint256 chainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) WatchLogChangeMPC(opts *bind.WatchOpts, sink chan<- *AnyswapV4RouterLogChangeMPC, oldMPC []common.Address, newMPC []common.Address, effectiveTime []*big.Int) (event.Subscription, error) {

	var oldMPCRule []interface{}
	for _, oldMPCItem := range oldMPC {
		oldMPCRule = append(oldMPCRule, oldMPCItem)
	}
	var newMPCRule []interface{}
	for _, newMPCItem := range newMPC {
		newMPCRule = append(newMPCRule, newMPCItem)
	}
	var effectiveTimeRule []interface{}
	for _, effectiveTimeItem := range effectiveTime {
		effectiveTimeRule = append(effectiveTimeRule, effectiveTimeItem)
	}

	logs, sub, err := _AnyswapV4Router.contract.WatchLogs(opts, "LogChangeMPC", oldMPCRule, newMPCRule, effectiveTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnyswapV4RouterLogChangeMPC)
				if err := _AnyswapV4Router.contract.UnpackLog(event, "LogChangeMPC", log); err != nil {
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

// ParseLogChangeMPC is a log parse operation binding the contract event 0xcda32bc39904597666dfa9f9c845714756e1ffffad55b52e0d344673a2198121.
//
// Solidity: event LogChangeMPC(address indexed oldMPC, address indexed newMPC, uint256 indexed effectiveTime, uint256 chainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) ParseLogChangeMPC(log types.Log) (*AnyswapV4RouterLogChangeMPC, error) {
	event := new(AnyswapV4RouterLogChangeMPC)
	if err := _AnyswapV4Router.contract.UnpackLog(event, "LogChangeMPC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnyswapV4RouterLogChangeRouterIterator is returned from FilterLogChangeRouter and is used to iterate over the raw logs and unpacked data for LogChangeRouter events raised by the AnyswapV4Router contract.
type AnyswapV4RouterLogChangeRouterIterator struct {
	Event *AnyswapV4RouterLogChangeRouter // Event containing the contract specifics and raw log

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
func (it *AnyswapV4RouterLogChangeRouterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnyswapV4RouterLogChangeRouter)
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
		it.Event = new(AnyswapV4RouterLogChangeRouter)
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
func (it *AnyswapV4RouterLogChangeRouterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnyswapV4RouterLogChangeRouterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnyswapV4RouterLogChangeRouter represents a LogChangeRouter event raised by the AnyswapV4Router contract.
type AnyswapV4RouterLogChangeRouter struct {
	OldRouter common.Address
	NewRouter common.Address
	ChainID   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogChangeRouter is a free log retrieval operation binding the contract event 0x7eefe162042d50d604dca716bef4ff4c5e318a056f712c0195d016f78089955a.
//
// Solidity: event LogChangeRouter(address indexed oldRouter, address indexed newRouter, uint256 chainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) FilterLogChangeRouter(opts *bind.FilterOpts, oldRouter []common.Address, newRouter []common.Address) (*AnyswapV4RouterLogChangeRouterIterator, error) {

	var oldRouterRule []interface{}
	for _, oldRouterItem := range oldRouter {
		oldRouterRule = append(oldRouterRule, oldRouterItem)
	}
	var newRouterRule []interface{}
	for _, newRouterItem := range newRouter {
		newRouterRule = append(newRouterRule, newRouterItem)
	}

	logs, sub, err := _AnyswapV4Router.contract.FilterLogs(opts, "LogChangeRouter", oldRouterRule, newRouterRule)
	if err != nil {
		return nil, err
	}
	return &AnyswapV4RouterLogChangeRouterIterator{contract: _AnyswapV4Router.contract, event: "LogChangeRouter", logs: logs, sub: sub}, nil
}

// WatchLogChangeRouter is a free log subscription operation binding the contract event 0x7eefe162042d50d604dca716bef4ff4c5e318a056f712c0195d016f78089955a.
//
// Solidity: event LogChangeRouter(address indexed oldRouter, address indexed newRouter, uint256 chainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) WatchLogChangeRouter(opts *bind.WatchOpts, sink chan<- *AnyswapV4RouterLogChangeRouter, oldRouter []common.Address, newRouter []common.Address) (event.Subscription, error) {

	var oldRouterRule []interface{}
	for _, oldRouterItem := range oldRouter {
		oldRouterRule = append(oldRouterRule, oldRouterItem)
	}
	var newRouterRule []interface{}
	for _, newRouterItem := range newRouter {
		newRouterRule = append(newRouterRule, newRouterItem)
	}

	logs, sub, err := _AnyswapV4Router.contract.WatchLogs(opts, "LogChangeRouter", oldRouterRule, newRouterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnyswapV4RouterLogChangeRouter)
				if err := _AnyswapV4Router.contract.UnpackLog(event, "LogChangeRouter", log); err != nil {
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

// ParseLogChangeRouter is a log parse operation binding the contract event 0x7eefe162042d50d604dca716bef4ff4c5e318a056f712c0195d016f78089955a.
//
// Solidity: event LogChangeRouter(address indexed oldRouter, address indexed newRouter, uint256 chainID)
func (_AnyswapV4Router *AnyswapV4RouterFilterer) ParseLogChangeRouter(log types.Log) (*AnyswapV4RouterLogChangeRouter, error) {
	event := new(AnyswapV4RouterLogChangeRouter)
	if err := _AnyswapV4Router.contract.UnpackLog(event, "LogChangeRouter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"transferWithPermit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"605629d6": "transferWithPermit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

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
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
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
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address target, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20 *IERC20Transactor) Permit(opts *bind.TransactOpts, target common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "permit", target, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address target, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20 *IERC20Session) Permit(target common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20.Contract.Permit(&_IERC20.TransactOpts, target, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address target, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20 *IERC20TransactorSession) Permit(target common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20.Contract.Permit(&_IERC20.TransactOpts, target, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferWithPermit is a paid mutator transaction binding the contract method 0x605629d6.
//
// Solidity: function transferWithPermit(address target, address to, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_IERC20 *IERC20Transactor) TransferWithPermit(opts *bind.TransactOpts, target common.Address, to common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferWithPermit", target, to, value, deadline, v, r, s)
}

// TransferWithPermit is a paid mutator transaction binding the contract method 0x605629d6.
//
// Solidity: function transferWithPermit(address target, address to, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_IERC20 *IERC20Session) TransferWithPermit(target common.Address, to common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20.Contract.TransferWithPermit(&_IERC20.TransactOpts, target, to, value, deadline, v, r, s)
}

// TransferWithPermit is a paid mutator transaction binding the contract method 0x605629d6.
//
// Solidity: function transferWithPermit(address target, address to, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferWithPermit(target common.Address, to common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20.Contract.TransferWithPermit(&_IERC20.TransactOpts, target, to, value, deadline, v, r, s)
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

// ISushiswapV2PairMetaData contains all meta data concerning the ISushiswapV2Pair contract.
var ISushiswapV2PairMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c45a0155": "factory()",
		"0902f1ac": "getReserves()",
		"022c0d9f": "swap(uint256,uint256,address,bytes)",
		"0dfe1681": "token0()",
		"d21220a7": "token1()",
	},
}

// ISushiswapV2PairABI is the input ABI used to generate the binding from.
// Deprecated: Use ISushiswapV2PairMetaData.ABI instead.
var ISushiswapV2PairABI = ISushiswapV2PairMetaData.ABI

// Deprecated: Use ISushiswapV2PairMetaData.Sigs instead.
// ISushiswapV2PairFuncSigs maps the 4-byte function signature to its string representation.
var ISushiswapV2PairFuncSigs = ISushiswapV2PairMetaData.Sigs

// ISushiswapV2Pair is an auto generated Go binding around an Ethereum contract.
type ISushiswapV2Pair struct {
	ISushiswapV2PairCaller     // Read-only binding to the contract
	ISushiswapV2PairTransactor // Write-only binding to the contract
	ISushiswapV2PairFilterer   // Log filterer for contract events
}

// ISushiswapV2PairCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISushiswapV2PairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISushiswapV2PairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISushiswapV2PairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISushiswapV2PairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISushiswapV2PairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISushiswapV2PairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISushiswapV2PairSession struct {
	Contract     *ISushiswapV2Pair // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISushiswapV2PairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISushiswapV2PairCallerSession struct {
	Contract *ISushiswapV2PairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ISushiswapV2PairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISushiswapV2PairTransactorSession struct {
	Contract     *ISushiswapV2PairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ISushiswapV2PairRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISushiswapV2PairRaw struct {
	Contract *ISushiswapV2Pair // Generic contract binding to access the raw methods on
}

// ISushiswapV2PairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISushiswapV2PairCallerRaw struct {
	Contract *ISushiswapV2PairCaller // Generic read-only contract binding to access the raw methods on
}

// ISushiswapV2PairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISushiswapV2PairTransactorRaw struct {
	Contract *ISushiswapV2PairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISushiswapV2Pair creates a new instance of ISushiswapV2Pair, bound to a specific deployed contract.
func NewISushiswapV2Pair(address common.Address, backend bind.ContractBackend) (*ISushiswapV2Pair, error) {
	contract, err := bindISushiswapV2Pair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISushiswapV2Pair{ISushiswapV2PairCaller: ISushiswapV2PairCaller{contract: contract}, ISushiswapV2PairTransactor: ISushiswapV2PairTransactor{contract: contract}, ISushiswapV2PairFilterer: ISushiswapV2PairFilterer{contract: contract}}, nil
}

// NewISushiswapV2PairCaller creates a new read-only instance of ISushiswapV2Pair, bound to a specific deployed contract.
func NewISushiswapV2PairCaller(address common.Address, caller bind.ContractCaller) (*ISushiswapV2PairCaller, error) {
	contract, err := bindISushiswapV2Pair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISushiswapV2PairCaller{contract: contract}, nil
}

// NewISushiswapV2PairTransactor creates a new write-only instance of ISushiswapV2Pair, bound to a specific deployed contract.
func NewISushiswapV2PairTransactor(address common.Address, transactor bind.ContractTransactor) (*ISushiswapV2PairTransactor, error) {
	contract, err := bindISushiswapV2Pair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISushiswapV2PairTransactor{contract: contract}, nil
}

// NewISushiswapV2PairFilterer creates a new log filterer instance of ISushiswapV2Pair, bound to a specific deployed contract.
func NewISushiswapV2PairFilterer(address common.Address, filterer bind.ContractFilterer) (*ISushiswapV2PairFilterer, error) {
	contract, err := bindISushiswapV2Pair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISushiswapV2PairFilterer{contract: contract}, nil
}

// bindISushiswapV2Pair binds a generic wrapper to an already deployed contract.
func bindISushiswapV2Pair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISushiswapV2PairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISushiswapV2Pair *ISushiswapV2PairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISushiswapV2Pair.Contract.ISushiswapV2PairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISushiswapV2Pair *ISushiswapV2PairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISushiswapV2Pair.Contract.ISushiswapV2PairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISushiswapV2Pair *ISushiswapV2PairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISushiswapV2Pair.Contract.ISushiswapV2PairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISushiswapV2Pair *ISushiswapV2PairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISushiswapV2Pair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISushiswapV2Pair *ISushiswapV2PairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISushiswapV2Pair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISushiswapV2Pair *ISushiswapV2PairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISushiswapV2Pair.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_ISushiswapV2Pair *ISushiswapV2PairCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISushiswapV2Pair.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_ISushiswapV2Pair *ISushiswapV2PairSession) Factory() (common.Address, error) {
	return _ISushiswapV2Pair.Contract.Factory(&_ISushiswapV2Pair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_ISushiswapV2Pair *ISushiswapV2PairCallerSession) Factory() (common.Address, error) {
	return _ISushiswapV2Pair.Contract.Factory(&_ISushiswapV2Pair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_ISushiswapV2Pair *ISushiswapV2PairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _ISushiswapV2Pair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_ISushiswapV2Pair *ISushiswapV2PairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _ISushiswapV2Pair.Contract.GetReserves(&_ISushiswapV2Pair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_ISushiswapV2Pair *ISushiswapV2PairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _ISushiswapV2Pair.Contract.GetReserves(&_ISushiswapV2Pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ISushiswapV2Pair *ISushiswapV2PairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISushiswapV2Pair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ISushiswapV2Pair *ISushiswapV2PairSession) Token0() (common.Address, error) {
	return _ISushiswapV2Pair.Contract.Token0(&_ISushiswapV2Pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ISushiswapV2Pair *ISushiswapV2PairCallerSession) Token0() (common.Address, error) {
	return _ISushiswapV2Pair.Contract.Token0(&_ISushiswapV2Pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ISushiswapV2Pair *ISushiswapV2PairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISushiswapV2Pair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ISushiswapV2Pair *ISushiswapV2PairSession) Token1() (common.Address, error) {
	return _ISushiswapV2Pair.Contract.Token1(&_ISushiswapV2Pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ISushiswapV2Pair *ISushiswapV2PairCallerSession) Token1() (common.Address, error) {
	return _ISushiswapV2Pair.Contract.Token1(&_ISushiswapV2Pair.CallOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_ISushiswapV2Pair *ISushiswapV2PairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _ISushiswapV2Pair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_ISushiswapV2Pair *ISushiswapV2PairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _ISushiswapV2Pair.Contract.Swap(&_ISushiswapV2Pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_ISushiswapV2Pair *ISushiswapV2PairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _ISushiswapV2Pair.Contract.Swap(&_ISushiswapV2Pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// IwNATIVEMetaData contains all meta data concerning the IwNATIVE contract.
var IwNATIVEMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d0e30db0": "deposit()",
		"a9059cbb": "transfer(address,uint256)",
		"2e1a7d4d": "withdraw(uint256)",
	},
}

// IwNATIVEABI is the input ABI used to generate the binding from.
// Deprecated: Use IwNATIVEMetaData.ABI instead.
var IwNATIVEABI = IwNATIVEMetaData.ABI

// Deprecated: Use IwNATIVEMetaData.Sigs instead.
// IwNATIVEFuncSigs maps the 4-byte function signature to its string representation.
var IwNATIVEFuncSigs = IwNATIVEMetaData.Sigs

// IwNATIVE is an auto generated Go binding around an Ethereum contract.
type IwNATIVE struct {
	IwNATIVECaller     // Read-only binding to the contract
	IwNATIVETransactor // Write-only binding to the contract
	IwNATIVEFilterer   // Log filterer for contract events
}

// IwNATIVECaller is an auto generated read-only Go binding around an Ethereum contract.
type IwNATIVECaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IwNATIVETransactor is an auto generated write-only Go binding around an Ethereum contract.
type IwNATIVETransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IwNATIVEFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IwNATIVEFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IwNATIVESession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IwNATIVESession struct {
	Contract     *IwNATIVE         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IwNATIVECallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IwNATIVECallerSession struct {
	Contract *IwNATIVECaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IwNATIVETransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IwNATIVETransactorSession struct {
	Contract     *IwNATIVETransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IwNATIVERaw is an auto generated low-level Go binding around an Ethereum contract.
type IwNATIVERaw struct {
	Contract *IwNATIVE // Generic contract binding to access the raw methods on
}

// IwNATIVECallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IwNATIVECallerRaw struct {
	Contract *IwNATIVECaller // Generic read-only contract binding to access the raw methods on
}

// IwNATIVETransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IwNATIVETransactorRaw struct {
	Contract *IwNATIVETransactor // Generic write-only contract binding to access the raw methods on
}

// NewIwNATIVE creates a new instance of IwNATIVE, bound to a specific deployed contract.
func NewIwNATIVE(address common.Address, backend bind.ContractBackend) (*IwNATIVE, error) {
	contract, err := bindIwNATIVE(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IwNATIVE{IwNATIVECaller: IwNATIVECaller{contract: contract}, IwNATIVETransactor: IwNATIVETransactor{contract: contract}, IwNATIVEFilterer: IwNATIVEFilterer{contract: contract}}, nil
}

// NewIwNATIVECaller creates a new read-only instance of IwNATIVE, bound to a specific deployed contract.
func NewIwNATIVECaller(address common.Address, caller bind.ContractCaller) (*IwNATIVECaller, error) {
	contract, err := bindIwNATIVE(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IwNATIVECaller{contract: contract}, nil
}

// NewIwNATIVETransactor creates a new write-only instance of IwNATIVE, bound to a specific deployed contract.
func NewIwNATIVETransactor(address common.Address, transactor bind.ContractTransactor) (*IwNATIVETransactor, error) {
	contract, err := bindIwNATIVE(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IwNATIVETransactor{contract: contract}, nil
}

// NewIwNATIVEFilterer creates a new log filterer instance of IwNATIVE, bound to a specific deployed contract.
func NewIwNATIVEFilterer(address common.Address, filterer bind.ContractFilterer) (*IwNATIVEFilterer, error) {
	contract, err := bindIwNATIVE(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IwNATIVEFilterer{contract: contract}, nil
}

// bindIwNATIVE binds a generic wrapper to an already deployed contract.
func bindIwNATIVE(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IwNATIVEABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IwNATIVE *IwNATIVERaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IwNATIVE.Contract.IwNATIVECaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IwNATIVE *IwNATIVERaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IwNATIVE.Contract.IwNATIVETransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IwNATIVE *IwNATIVERaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IwNATIVE.Contract.IwNATIVETransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IwNATIVE *IwNATIVECallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IwNATIVE.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IwNATIVE *IwNATIVETransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IwNATIVE.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IwNATIVE *IwNATIVETransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IwNATIVE.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IwNATIVE *IwNATIVETransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IwNATIVE.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IwNATIVE *IwNATIVESession) Deposit() (*types.Transaction, error) {
	return _IwNATIVE.Contract.Deposit(&_IwNATIVE.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IwNATIVE *IwNATIVETransactorSession) Deposit() (*types.Transaction, error) {
	return _IwNATIVE.Contract.Deposit(&_IwNATIVE.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IwNATIVE *IwNATIVETransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IwNATIVE.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IwNATIVE *IwNATIVESession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IwNATIVE.Contract.Transfer(&_IwNATIVE.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IwNATIVE *IwNATIVETransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IwNATIVE.Contract.Transfer(&_IwNATIVE.TransactOpts, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IwNATIVE *IwNATIVETransactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IwNATIVE.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IwNATIVE *IwNATIVESession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IwNATIVE.Contract.Withdraw(&_IwNATIVE.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IwNATIVE *IwNATIVETransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IwNATIVE.Contract.Withdraw(&_IwNATIVE.TransactOpts, arg0)
}

// SafeMathSushiswapMetaData contains all meta data concerning the SafeMathSushiswap contract.
var SafeMathSushiswapMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ed902324e881173170adeb46593dc5ec204bd33ad4a36f8bc03855a5b917d62c64736f6c63430008000033",
}

// SafeMathSushiswapABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathSushiswapMetaData.ABI instead.
var SafeMathSushiswapABI = SafeMathSushiswapMetaData.ABI

// SafeMathSushiswapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathSushiswapMetaData.Bin instead.
var SafeMathSushiswapBin = SafeMathSushiswapMetaData.Bin

// DeploySafeMathSushiswap deploys a new Ethereum contract, binding an instance of SafeMathSushiswap to it.
func DeploySafeMathSushiswap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMathSushiswap, error) {
	parsed, err := SafeMathSushiswapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathSushiswapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMathSushiswap{SafeMathSushiswapCaller: SafeMathSushiswapCaller{contract: contract}, SafeMathSushiswapTransactor: SafeMathSushiswapTransactor{contract: contract}, SafeMathSushiswapFilterer: SafeMathSushiswapFilterer{contract: contract}}, nil
}

// SafeMathSushiswap is an auto generated Go binding around an Ethereum contract.
type SafeMathSushiswap struct {
	SafeMathSushiswapCaller     // Read-only binding to the contract
	SafeMathSushiswapTransactor // Write-only binding to the contract
	SafeMathSushiswapFilterer   // Log filterer for contract events
}

// SafeMathSushiswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathSushiswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSushiswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathSushiswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSushiswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathSushiswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSushiswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSushiswapSession struct {
	Contract     *SafeMathSushiswap // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SafeMathSushiswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathSushiswapCallerSession struct {
	Contract *SafeMathSushiswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SafeMathSushiswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathSushiswapTransactorSession struct {
	Contract     *SafeMathSushiswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SafeMathSushiswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathSushiswapRaw struct {
	Contract *SafeMathSushiswap // Generic contract binding to access the raw methods on
}

// SafeMathSushiswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathSushiswapCallerRaw struct {
	Contract *SafeMathSushiswapCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathSushiswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathSushiswapTransactorRaw struct {
	Contract *SafeMathSushiswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMathSushiswap creates a new instance of SafeMathSushiswap, bound to a specific deployed contract.
func NewSafeMathSushiswap(address common.Address, backend bind.ContractBackend) (*SafeMathSushiswap, error) {
	contract, err := bindSafeMathSushiswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMathSushiswap{SafeMathSushiswapCaller: SafeMathSushiswapCaller{contract: contract}, SafeMathSushiswapTransactor: SafeMathSushiswapTransactor{contract: contract}, SafeMathSushiswapFilterer: SafeMathSushiswapFilterer{contract: contract}}, nil
}

// NewSafeMathSushiswapCaller creates a new read-only instance of SafeMathSushiswap, bound to a specific deployed contract.
func NewSafeMathSushiswapCaller(address common.Address, caller bind.ContractCaller) (*SafeMathSushiswapCaller, error) {
	contract, err := bindSafeMathSushiswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathSushiswapCaller{contract: contract}, nil
}

// NewSafeMathSushiswapTransactor creates a new write-only instance of SafeMathSushiswap, bound to a specific deployed contract.
func NewSafeMathSushiswapTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathSushiswapTransactor, error) {
	contract, err := bindSafeMathSushiswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathSushiswapTransactor{contract: contract}, nil
}

// NewSafeMathSushiswapFilterer creates a new log filterer instance of SafeMathSushiswap, bound to a specific deployed contract.
func NewSafeMathSushiswapFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathSushiswapFilterer, error) {
	contract, err := bindSafeMathSushiswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathSushiswapFilterer{contract: contract}, nil
}

// bindSafeMathSushiswap binds a generic wrapper to an already deployed contract.
func bindSafeMathSushiswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathSushiswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathSushiswap *SafeMathSushiswapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMathSushiswap.Contract.SafeMathSushiswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathSushiswap *SafeMathSushiswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathSushiswap.Contract.SafeMathSushiswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathSushiswap *SafeMathSushiswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathSushiswap.Contract.SafeMathSushiswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathSushiswap *SafeMathSushiswapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMathSushiswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathSushiswap *SafeMathSushiswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathSushiswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathSushiswap *SafeMathSushiswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathSushiswap.Contract.contract.Transact(opts, method, params...)
}

// SushiswapV2LibraryMetaData contains all meta data concerning the SushiswapV2Library contract.
var SushiswapV2LibraryMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220db417f1d5f862447118e7db81ef76aa6de04c36a895079b4b02d27188d2a977764736f6c63430008000033",
}

// SushiswapV2LibraryABI is the input ABI used to generate the binding from.
// Deprecated: Use SushiswapV2LibraryMetaData.ABI instead.
var SushiswapV2LibraryABI = SushiswapV2LibraryMetaData.ABI

// SushiswapV2LibraryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SushiswapV2LibraryMetaData.Bin instead.
var SushiswapV2LibraryBin = SushiswapV2LibraryMetaData.Bin

// DeploySushiswapV2Library deploys a new Ethereum contract, binding an instance of SushiswapV2Library to it.
func DeploySushiswapV2Library(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SushiswapV2Library, error) {
	parsed, err := SushiswapV2LibraryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SushiswapV2LibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SushiswapV2Library{SushiswapV2LibraryCaller: SushiswapV2LibraryCaller{contract: contract}, SushiswapV2LibraryTransactor: SushiswapV2LibraryTransactor{contract: contract}, SushiswapV2LibraryFilterer: SushiswapV2LibraryFilterer{contract: contract}}, nil
}

// SushiswapV2Library is an auto generated Go binding around an Ethereum contract.
type SushiswapV2Library struct {
	SushiswapV2LibraryCaller     // Read-only binding to the contract
	SushiswapV2LibraryTransactor // Write-only binding to the contract
	SushiswapV2LibraryFilterer   // Log filterer for contract events
}

// SushiswapV2LibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SushiswapV2LibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiswapV2LibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SushiswapV2LibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiswapV2LibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SushiswapV2LibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiswapV2LibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SushiswapV2LibrarySession struct {
	Contract     *SushiswapV2Library // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SushiswapV2LibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SushiswapV2LibraryCallerSession struct {
	Contract *SushiswapV2LibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SushiswapV2LibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SushiswapV2LibraryTransactorSession struct {
	Contract     *SushiswapV2LibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SushiswapV2LibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SushiswapV2LibraryRaw struct {
	Contract *SushiswapV2Library // Generic contract binding to access the raw methods on
}

// SushiswapV2LibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SushiswapV2LibraryCallerRaw struct {
	Contract *SushiswapV2LibraryCaller // Generic read-only contract binding to access the raw methods on
}

// SushiswapV2LibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SushiswapV2LibraryTransactorRaw struct {
	Contract *SushiswapV2LibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSushiswapV2Library creates a new instance of SushiswapV2Library, bound to a specific deployed contract.
func NewSushiswapV2Library(address common.Address, backend bind.ContractBackend) (*SushiswapV2Library, error) {
	contract, err := bindSushiswapV2Library(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SushiswapV2Library{SushiswapV2LibraryCaller: SushiswapV2LibraryCaller{contract: contract}, SushiswapV2LibraryTransactor: SushiswapV2LibraryTransactor{contract: contract}, SushiswapV2LibraryFilterer: SushiswapV2LibraryFilterer{contract: contract}}, nil
}

// NewSushiswapV2LibraryCaller creates a new read-only instance of SushiswapV2Library, bound to a specific deployed contract.
func NewSushiswapV2LibraryCaller(address common.Address, caller bind.ContractCaller) (*SushiswapV2LibraryCaller, error) {
	contract, err := bindSushiswapV2Library(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SushiswapV2LibraryCaller{contract: contract}, nil
}

// NewSushiswapV2LibraryTransactor creates a new write-only instance of SushiswapV2Library, bound to a specific deployed contract.
func NewSushiswapV2LibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*SushiswapV2LibraryTransactor, error) {
	contract, err := bindSushiswapV2Library(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SushiswapV2LibraryTransactor{contract: contract}, nil
}

// NewSushiswapV2LibraryFilterer creates a new log filterer instance of SushiswapV2Library, bound to a specific deployed contract.
func NewSushiswapV2LibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*SushiswapV2LibraryFilterer, error) {
	contract, err := bindSushiswapV2Library(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SushiswapV2LibraryFilterer{contract: contract}, nil
}

// bindSushiswapV2Library binds a generic wrapper to an already deployed contract.
func bindSushiswapV2Library(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SushiswapV2LibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SushiswapV2Library *SushiswapV2LibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SushiswapV2Library.Contract.SushiswapV2LibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SushiswapV2Library *SushiswapV2LibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SushiswapV2Library.Contract.SushiswapV2LibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SushiswapV2Library *SushiswapV2LibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SushiswapV2Library.Contract.SushiswapV2LibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SushiswapV2Library *SushiswapV2LibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SushiswapV2Library.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SushiswapV2Library *SushiswapV2LibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SushiswapV2Library.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SushiswapV2Library *SushiswapV2LibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SushiswapV2Library.Contract.contract.Transact(opts, method, params...)
}

// TransferHelperMetaData contains all meta data concerning the TransferHelper contract.
var TransferHelperMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122039925f33bcae3476fee70e43402ab2b0b7c8b046276b71bd09c1efa1a0ac1b8264736f6c63430008000033",
}

// TransferHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferHelperMetaData.ABI instead.
var TransferHelperABI = TransferHelperMetaData.ABI

// TransferHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransferHelperMetaData.Bin instead.
var TransferHelperBin = TransferHelperMetaData.Bin

// DeployTransferHelper deploys a new Ethereum contract, binding an instance of TransferHelper to it.
func DeployTransferHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TransferHelper, error) {
	parsed, err := TransferHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransferHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// TransferHelper is an auto generated Go binding around an Ethereum contract.
type TransferHelper struct {
	TransferHelperCaller     // Read-only binding to the contract
	TransferHelperTransactor // Write-only binding to the contract
	TransferHelperFilterer   // Log filterer for contract events
}

// TransferHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferHelperSession struct {
	Contract     *TransferHelper   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferHelperCallerSession struct {
	Contract *TransferHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TransferHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferHelperTransactorSession struct {
	Contract     *TransferHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TransferHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferHelperRaw struct {
	Contract *TransferHelper // Generic contract binding to access the raw methods on
}

// TransferHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferHelperCallerRaw struct {
	Contract *TransferHelperCaller // Generic read-only contract binding to access the raw methods on
}

// TransferHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferHelperTransactorRaw struct {
	Contract *TransferHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferHelper creates a new instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelper(address common.Address, backend bind.ContractBackend) (*TransferHelper, error) {
	contract, err := bindTransferHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// NewTransferHelperCaller creates a new read-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperCaller(address common.Address, caller bind.ContractCaller) (*TransferHelperCaller, error) {
	contract, err := bindTransferHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperCaller{contract: contract}, nil
}

// NewTransferHelperTransactor creates a new write-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferHelperTransactor, error) {
	contract, err := bindTransferHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperTransactor{contract: contract}, nil
}

// NewTransferHelperFilterer creates a new log filterer instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferHelperFilterer, error) {
	contract, err := bindTransferHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferHelperFilterer{contract: contract}, nil
}

// bindTransferHelper binds a generic wrapper to an already deployed contract.
func bindTransferHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.TransferHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transact(opts, method, params...)
}
