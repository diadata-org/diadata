// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20contract

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

// Erc20contractABI is the input ABI used to generate the binding from.
const Erc20contractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"batFundDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"batFund\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenExchangeRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenCreationCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fundingEndBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethFundDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createTokens\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenCreationMin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fundingStartBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ethFundDeposit\",\"type\":\"address\"},{\"name\":\"_batFundDeposit\",\"type\":\"address\"},{\"name\":\"_fundingStartBlock\",\"type\":\"uint256\"},{\"name\":\"_fundingEndBlock\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"LogRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"CreateBAT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// Erc20contract is an auto generated Go binding around an Ethereum contract.
type Erc20contract struct {
	Erc20contractCaller     // Read-only binding to the contract
	Erc20contractTransactor // Write-only binding to the contract
	Erc20contractFilterer   // Log filterer for contract events
}

// Erc20contractCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20contractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20contractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20contractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20contractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20contractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20contractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20contractSession struct {
	Contract     *Erc20contract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20contractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20contractCallerSession struct {
	Contract *Erc20contractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Erc20contractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20contractTransactorSession struct {
	Contract     *Erc20contractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Erc20contractRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20contractRaw struct {
	Contract *Erc20contract // Generic contract binding to access the raw methods on
}

// Erc20contractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20contractCallerRaw struct {
	Contract *Erc20contractCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20contractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20contractTransactorRaw struct {
	Contract *Erc20contractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20contract creates a new instance of Erc20contract, bound to a specific deployed contract.
func NewErc20contract(address common.Address, backend bind.ContractBackend) (*Erc20contract, error) {
	contract, err := bindErc20contract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20contract{Erc20contractCaller: Erc20contractCaller{contract: contract}, Erc20contractTransactor: Erc20contractTransactor{contract: contract}, Erc20contractFilterer: Erc20contractFilterer{contract: contract}}, nil
}

// NewErc20contractCaller creates a new read-only instance of Erc20contract, bound to a specific deployed contract.
func NewErc20contractCaller(address common.Address, caller bind.ContractCaller) (*Erc20contractCaller, error) {
	contract, err := bindErc20contract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20contractCaller{contract: contract}, nil
}

// NewErc20contractTransactor creates a new write-only instance of Erc20contract, bound to a specific deployed contract.
func NewErc20contractTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20contractTransactor, error) {
	contract, err := bindErc20contract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20contractTransactor{contract: contract}, nil
}

// NewErc20contractFilterer creates a new log filterer instance of Erc20contract, bound to a specific deployed contract.
func NewErc20contractFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20contractFilterer, error) {
	contract, err := bindErc20contract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20contractFilterer{contract: contract}, nil
}

// bindErc20contract binds a generic wrapper to an already deployed contract.
func bindErc20contract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20contractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20contract *Erc20contractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20contract.Contract.Erc20contractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20contract *Erc20contractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20contract.Contract.Erc20contractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20contract *Erc20contractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20contract.Contract.Erc20contractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20contract *Erc20contractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20contract *Erc20contractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20contract *Erc20contractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20contract.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 remaining)
func (_Erc20contract *Erc20contractCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 remaining)
func (_Erc20contract *Erc20contractSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Erc20contract.Contract.Allowance(&_Erc20contract.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 remaining)
func (_Erc20contract *Erc20contractCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Erc20contract.Contract.Allowance(&_Erc20contract.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_Erc20contract *Erc20contractCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_Erc20contract *Erc20contractSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Erc20contract.Contract.BalanceOf(&_Erc20contract.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_Erc20contract *Erc20contractCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Erc20contract.Contract.BalanceOf(&_Erc20contract.CallOpts, _owner)
}

// BatFund is a free data retrieval call binding the contract method 0x229a4978.
//
// Solidity: function batFund() returns(uint256)
func (_Erc20contract *Erc20contractCaller) BatFund(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "batFund")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatFund is a free data retrieval call binding the contract method 0x229a4978.
//
// Solidity: function batFund() returns(uint256)
func (_Erc20contract *Erc20contractSession) BatFund() (*big.Int, error) {
	return _Erc20contract.Contract.BatFund(&_Erc20contract.CallOpts)
}

// BatFund is a free data retrieval call binding the contract method 0x229a4978.
//
// Solidity: function batFund() returns(uint256)
func (_Erc20contract *Erc20contractCallerSession) BatFund() (*big.Int, error) {
	return _Erc20contract.Contract.BatFund(&_Erc20contract.CallOpts)
}

// BatFundDeposit is a free data retrieval call binding the contract method 0x01a7a8c0.
//
// Solidity: function batFundDeposit() returns(address)
func (_Erc20contract *Erc20contractCaller) BatFundDeposit(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "batFundDeposit")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BatFundDeposit is a free data retrieval call binding the contract method 0x01a7a8c0.
//
// Solidity: function batFundDeposit() returns(address)
func (_Erc20contract *Erc20contractSession) BatFundDeposit() (common.Address, error) {
	return _Erc20contract.Contract.BatFundDeposit(&_Erc20contract.CallOpts)
}

// BatFundDeposit is a free data retrieval call binding the contract method 0x01a7a8c0.
//
// Solidity: function batFundDeposit() returns(address)
func (_Erc20contract *Erc20contractCallerSession) BatFundDeposit() (common.Address, error) {
	return _Erc20contract.Contract.BatFundDeposit(&_Erc20contract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256)
func (_Erc20contract *Erc20contractCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256)
func (_Erc20contract *Erc20contractSession) Decimals() (*big.Int, error) {
	return _Erc20contract.Contract.Decimals(&_Erc20contract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256)
func (_Erc20contract *Erc20contractCallerSession) Decimals() (*big.Int, error) {
	return _Erc20contract.Contract.Decimals(&_Erc20contract.CallOpts)
}

// EthFundDeposit is a free data retrieval call binding the contract method 0xa81c3bdf.
//
// Solidity: function ethFundDeposit() returns(address)
func (_Erc20contract *Erc20contractCaller) EthFundDeposit(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "ethFundDeposit")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthFundDeposit is a free data retrieval call binding the contract method 0xa81c3bdf.
//
// Solidity: function ethFundDeposit() returns(address)
func (_Erc20contract *Erc20contractSession) EthFundDeposit() (common.Address, error) {
	return _Erc20contract.Contract.EthFundDeposit(&_Erc20contract.CallOpts)
}

// EthFundDeposit is a free data retrieval call binding the contract method 0xa81c3bdf.
//
// Solidity: function ethFundDeposit() returns(address)
func (_Erc20contract *Erc20contractCallerSession) EthFundDeposit() (common.Address, error) {
	return _Erc20contract.Contract.EthFundDeposit(&_Erc20contract.CallOpts)
}

// FundingEndBlock is a free data retrieval call binding the contract method 0x91b43d13.
//
// Solidity: function fundingEndBlock() returns(uint256)
func (_Erc20contract *Erc20contractCaller) FundingEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "fundingEndBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingEndBlock is a free data retrieval call binding the contract method 0x91b43d13.
//
// Solidity: function fundingEndBlock() returns(uint256)
func (_Erc20contract *Erc20contractSession) FundingEndBlock() (*big.Int, error) {
	return _Erc20contract.Contract.FundingEndBlock(&_Erc20contract.CallOpts)
}

// FundingEndBlock is a free data retrieval call binding the contract method 0x91b43d13.
//
// Solidity: function fundingEndBlock() returns(uint256)
func (_Erc20contract *Erc20contractCallerSession) FundingEndBlock() (*big.Int, error) {
	return _Erc20contract.Contract.FundingEndBlock(&_Erc20contract.CallOpts)
}

// FundingStartBlock is a free data retrieval call binding the contract method 0xd648a647.
//
// Solidity: function fundingStartBlock() returns(uint256)
func (_Erc20contract *Erc20contractCaller) FundingStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "fundingStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingStartBlock is a free data retrieval call binding the contract method 0xd648a647.
//
// Solidity: function fundingStartBlock() returns(uint256)
func (_Erc20contract *Erc20contractSession) FundingStartBlock() (*big.Int, error) {
	return _Erc20contract.Contract.FundingStartBlock(&_Erc20contract.CallOpts)
}

// FundingStartBlock is a free data retrieval call binding the contract method 0xd648a647.
//
// Solidity: function fundingStartBlock() returns(uint256)
func (_Erc20contract *Erc20contractCallerSession) FundingStartBlock() (*big.Int, error) {
	return _Erc20contract.Contract.FundingStartBlock(&_Erc20contract.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() returns(bool)
func (_Erc20contract *Erc20contractCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "isFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() returns(bool)
func (_Erc20contract *Erc20contractSession) IsFinalized() (bool, error) {
	return _Erc20contract.Contract.IsFinalized(&_Erc20contract.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() returns(bool)
func (_Erc20contract *Erc20contractCallerSession) IsFinalized() (bool, error) {
	return _Erc20contract.Contract.IsFinalized(&_Erc20contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_Erc20contract *Erc20contractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_Erc20contract *Erc20contractSession) Name() (string, error) {
	return _Erc20contract.Contract.Name(&_Erc20contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_Erc20contract *Erc20contractCallerSession) Name() (string, error) {
	return _Erc20contract.Contract.Name(&_Erc20contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_Erc20contract *Erc20contractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_Erc20contract *Erc20contractSession) Symbol() (string, error) {
	return _Erc20contract.Contract.Symbol(&_Erc20contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_Erc20contract *Erc20contractCallerSession) Symbol() (string, error) {
	return _Erc20contract.Contract.Symbol(&_Erc20contract.CallOpts)
}

// TokenCreationCap is a free data retrieval call binding the contract method 0x6f7920fd.
//
// Solidity: function tokenCreationCap() returns(uint256)
func (_Erc20contract *Erc20contractCaller) TokenCreationCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "tokenCreationCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCreationCap is a free data retrieval call binding the contract method 0x6f7920fd.
//
// Solidity: function tokenCreationCap() returns(uint256)
func (_Erc20contract *Erc20contractSession) TokenCreationCap() (*big.Int, error) {
	return _Erc20contract.Contract.TokenCreationCap(&_Erc20contract.CallOpts)
}

// TokenCreationCap is a free data retrieval call binding the contract method 0x6f7920fd.
//
// Solidity: function tokenCreationCap() returns(uint256)
func (_Erc20contract *Erc20contractCallerSession) TokenCreationCap() (*big.Int, error) {
	return _Erc20contract.Contract.TokenCreationCap(&_Erc20contract.CallOpts)
}

// TokenCreationMin is a free data retrieval call binding the contract method 0xc039daf6.
//
// Solidity: function tokenCreationMin() returns(uint256)
func (_Erc20contract *Erc20contractCaller) TokenCreationMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "tokenCreationMin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCreationMin is a free data retrieval call binding the contract method 0xc039daf6.
//
// Solidity: function tokenCreationMin() returns(uint256)
func (_Erc20contract *Erc20contractSession) TokenCreationMin() (*big.Int, error) {
	return _Erc20contract.Contract.TokenCreationMin(&_Erc20contract.CallOpts)
}

// TokenCreationMin is a free data retrieval call binding the contract method 0xc039daf6.
//
// Solidity: function tokenCreationMin() returns(uint256)
func (_Erc20contract *Erc20contractCallerSession) TokenCreationMin() (*big.Int, error) {
	return _Erc20contract.Contract.TokenCreationMin(&_Erc20contract.CallOpts)
}

// TokenExchangeRate is a free data retrieval call binding the contract method 0x4172d080.
//
// Solidity: function tokenExchangeRate() returns(uint256)
func (_Erc20contract *Erc20contractCaller) TokenExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "tokenExchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenExchangeRate is a free data retrieval call binding the contract method 0x4172d080.
//
// Solidity: function tokenExchangeRate() returns(uint256)
func (_Erc20contract *Erc20contractSession) TokenExchangeRate() (*big.Int, error) {
	return _Erc20contract.Contract.TokenExchangeRate(&_Erc20contract.CallOpts)
}

// TokenExchangeRate is a free data retrieval call binding the contract method 0x4172d080.
//
// Solidity: function tokenExchangeRate() returns(uint256)
func (_Erc20contract *Erc20contractCallerSession) TokenExchangeRate() (*big.Int, error) {
	return _Erc20contract.Contract.TokenExchangeRate(&_Erc20contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_Erc20contract *Erc20contractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_Erc20contract *Erc20contractSession) TotalSupply() (*big.Int, error) {
	return _Erc20contract.Contract.TotalSupply(&_Erc20contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_Erc20contract *Erc20contractCallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20contract.Contract.TotalSupply(&_Erc20contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(string)
func (_Erc20contract *Erc20contractCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20contract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(string)
func (_Erc20contract *Erc20contractSession) Version() (string, error) {
	return _Erc20contract.Contract.Version(&_Erc20contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(string)
func (_Erc20contract *Erc20contractCallerSession) Version() (string, error) {
	return _Erc20contract.Contract.Version(&_Erc20contract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Erc20contract *Erc20contractTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20contract.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Erc20contract *Erc20contractSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20contract.Contract.Approve(&_Erc20contract.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Erc20contract *Erc20contractTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20contract.Contract.Approve(&_Erc20contract.TransactOpts, _spender, _value)
}

// CreateTokens is a paid mutator transaction binding the contract method 0xb4427263.
//
// Solidity: function createTokens() returns()
func (_Erc20contract *Erc20contractTransactor) CreateTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20contract.contract.Transact(opts, "createTokens")
}

// CreateTokens is a paid mutator transaction binding the contract method 0xb4427263.
//
// Solidity: function createTokens() returns()
func (_Erc20contract *Erc20contractSession) CreateTokens() (*types.Transaction, error) {
	return _Erc20contract.Contract.CreateTokens(&_Erc20contract.TransactOpts)
}

// CreateTokens is a paid mutator transaction binding the contract method 0xb4427263.
//
// Solidity: function createTokens() returns()
func (_Erc20contract *Erc20contractTransactorSession) CreateTokens() (*types.Transaction, error) {
	return _Erc20contract.Contract.CreateTokens(&_Erc20contract.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_Erc20contract *Erc20contractTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20contract.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_Erc20contract *Erc20contractSession) Finalize() (*types.Transaction, error) {
	return _Erc20contract.Contract.Finalize(&_Erc20contract.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_Erc20contract *Erc20contractTransactorSession) Finalize() (*types.Transaction, error) {
	return _Erc20contract.Contract.Finalize(&_Erc20contract.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Erc20contract *Erc20contractTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20contract.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Erc20contract *Erc20contractSession) Refund() (*types.Transaction, error) {
	return _Erc20contract.Contract.Refund(&_Erc20contract.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Erc20contract *Erc20contractTransactorSession) Refund() (*types.Transaction, error) {
	return _Erc20contract.Contract.Refund(&_Erc20contract.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Erc20contract *Erc20contractTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20contract.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Erc20contract *Erc20contractSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20contract.Contract.Transfer(&_Erc20contract.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Erc20contract *Erc20contractTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20contract.Contract.Transfer(&_Erc20contract.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Erc20contract *Erc20contractTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20contract.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Erc20contract *Erc20contractSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20contract.Contract.TransferFrom(&_Erc20contract.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Erc20contract *Erc20contractTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20contract.Contract.TransferFrom(&_Erc20contract.TransactOpts, _from, _to, _value)
}

// Erc20contractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20contract contract.
type Erc20contractApprovalIterator struct {
	Event *Erc20contractApproval // Event containing the contract specifics and raw log

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
func (it *Erc20contractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20contractApproval)
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
		it.Event = new(Erc20contractApproval)
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
func (it *Erc20contractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20contractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20contractApproval represents a Approval event raised by the Erc20contract contract.
type Erc20contractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Erc20contract *Erc20contractFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*Erc20contractApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Erc20contract.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20contractApprovalIterator{contract: _Erc20contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Erc20contract *Erc20contractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20contractApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Erc20contract.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20contractApproval)
				if err := _Erc20contract.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Erc20contract *Erc20contractFilterer) ParseApproval(log types.Log) (*Erc20contractApproval, error) {
	event := new(Erc20contractApproval)
	if err := _Erc20contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20contractCreateBATIterator is returned from FilterCreateBAT and is used to iterate over the raw logs and unpacked data for CreateBAT events raised by the Erc20contract contract.
type Erc20contractCreateBATIterator struct {
	Event *Erc20contractCreateBAT // Event containing the contract specifics and raw log

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
func (it *Erc20contractCreateBATIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20contractCreateBAT)
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
		it.Event = new(Erc20contractCreateBAT)
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
func (it *Erc20contractCreateBATIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20contractCreateBATIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20contractCreateBAT represents a CreateBAT event raised by the Erc20contract contract.
type Erc20contractCreateBAT struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCreateBAT is a free log retrieval operation binding the contract event 0xb33527d2e0d30b7aece2c5e82927985866c1b75173d671c14f4457bf67aa6910.
//
// Solidity: event CreateBAT(address indexed _to, uint256 _value)
func (_Erc20contract *Erc20contractFilterer) FilterCreateBAT(opts *bind.FilterOpts, _to []common.Address) (*Erc20contractCreateBATIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Erc20contract.contract.FilterLogs(opts, "CreateBAT", _toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20contractCreateBATIterator{contract: _Erc20contract.contract, event: "CreateBAT", logs: logs, sub: sub}, nil
}

// WatchCreateBAT is a free log subscription operation binding the contract event 0xb33527d2e0d30b7aece2c5e82927985866c1b75173d671c14f4457bf67aa6910.
//
// Solidity: event CreateBAT(address indexed _to, uint256 _value)
func (_Erc20contract *Erc20contractFilterer) WatchCreateBAT(opts *bind.WatchOpts, sink chan<- *Erc20contractCreateBAT, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Erc20contract.contract.WatchLogs(opts, "CreateBAT", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20contractCreateBAT)
				if err := _Erc20contract.contract.UnpackLog(event, "CreateBAT", log); err != nil {
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

// ParseCreateBAT is a log parse operation binding the contract event 0xb33527d2e0d30b7aece2c5e82927985866c1b75173d671c14f4457bf67aa6910.
//
// Solidity: event CreateBAT(address indexed _to, uint256 _value)
func (_Erc20contract *Erc20contractFilterer) ParseCreateBAT(log types.Log) (*Erc20contractCreateBAT, error) {
	event := new(Erc20contractCreateBAT)
	if err := _Erc20contract.contract.UnpackLog(event, "CreateBAT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20contractLogRefundIterator is returned from FilterLogRefund and is used to iterate over the raw logs and unpacked data for LogRefund events raised by the Erc20contract contract.
type Erc20contractLogRefundIterator struct {
	Event *Erc20contractLogRefund // Event containing the contract specifics and raw log

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
func (it *Erc20contractLogRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20contractLogRefund)
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
		it.Event = new(Erc20contractLogRefund)
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
func (it *Erc20contractLogRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20contractLogRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20contractLogRefund represents a LogRefund event raised by the Erc20contract contract.
type Erc20contractLogRefund struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogRefund is a free log retrieval operation binding the contract event 0xb6c0eca8138e097d71e2dd31e19a1266487f0553f170b7260ffe68bcbe9ff8a7.
//
// Solidity: event LogRefund(address indexed _to, uint256 _value)
func (_Erc20contract *Erc20contractFilterer) FilterLogRefund(opts *bind.FilterOpts, _to []common.Address) (*Erc20contractLogRefundIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Erc20contract.contract.FilterLogs(opts, "LogRefund", _toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20contractLogRefundIterator{contract: _Erc20contract.contract, event: "LogRefund", logs: logs, sub: sub}, nil
}

// WatchLogRefund is a free log subscription operation binding the contract event 0xb6c0eca8138e097d71e2dd31e19a1266487f0553f170b7260ffe68bcbe9ff8a7.
//
// Solidity: event LogRefund(address indexed _to, uint256 _value)
func (_Erc20contract *Erc20contractFilterer) WatchLogRefund(opts *bind.WatchOpts, sink chan<- *Erc20contractLogRefund, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Erc20contract.contract.WatchLogs(opts, "LogRefund", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20contractLogRefund)
				if err := _Erc20contract.contract.UnpackLog(event, "LogRefund", log); err != nil {
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

// ParseLogRefund is a log parse operation binding the contract event 0xb6c0eca8138e097d71e2dd31e19a1266487f0553f170b7260ffe68bcbe9ff8a7.
//
// Solidity: event LogRefund(address indexed _to, uint256 _value)
func (_Erc20contract *Erc20contractFilterer) ParseLogRefund(log types.Log) (*Erc20contractLogRefund, error) {
	event := new(Erc20contractLogRefund)
	if err := _Erc20contract.contract.UnpackLog(event, "LogRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20contractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20contract contract.
type Erc20contractTransferIterator struct {
	Event *Erc20contractTransfer // Event containing the contract specifics and raw log

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
func (it *Erc20contractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20contractTransfer)
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
		it.Event = new(Erc20contractTransfer)
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
func (it *Erc20contractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20contractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20contractTransfer represents a Transfer event raised by the Erc20contract contract.
type Erc20contractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Erc20contract *Erc20contractFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*Erc20contractTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Erc20contract.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20contractTransferIterator{contract: _Erc20contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Erc20contract *Erc20contractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20contractTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Erc20contract.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20contractTransfer)
				if err := _Erc20contract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Erc20contract *Erc20contractFilterer) ParseTransfer(log types.Log) (*Erc20contractTransfer, error) {
	event := new(Erc20contractTransfer)
	if err := _Erc20contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
