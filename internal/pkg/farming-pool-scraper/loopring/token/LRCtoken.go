// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LRCtoken

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

// LRCtokenABI is the input ABI used to generate the binding from.
const LRCtokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bonusPercentages\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DECIMALS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BLOCKS_PER_PHASE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_UNSOLD_RATIO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"HARD_CAP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BASE_RATE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleStarted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issueIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"issueToken\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_firstblock\",\"type\":\"uint256\"}],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hardCapReached\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleEnded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unsoldTokenIssued\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"tokens\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GOAL\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalEthReceived\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleDue\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"target\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NUM_OF_PHASE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"firstblock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SYMBOL\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SaleStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SaleEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InvalidCaller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"msg\",\"type\":\"bytes\"}],\"name\":\"InvalidState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"issueIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SaleSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SaleFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// LRCtoken is an auto generated Go binding around an Ethereum contract.
type LRCtoken struct {
	LRCtokenCaller     // Read-only binding to the contract
	LRCtokenTransactor // Write-only binding to the contract
	LRCtokenFilterer   // Log filterer for contract events
}

// LRCtokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type LRCtokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LRCtokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LRCtokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LRCtokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LRCtokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LRCtokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LRCtokenSession struct {
	Contract     *LRCtoken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LRCtokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LRCtokenCallerSession struct {
	Contract *LRCtokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LRCtokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LRCtokenTransactorSession struct {
	Contract     *LRCtokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LRCtokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type LRCtokenRaw struct {
	Contract *LRCtoken // Generic contract binding to access the raw methods on
}

// LRCtokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LRCtokenCallerRaw struct {
	Contract *LRCtokenCaller // Generic read-only contract binding to access the raw methods on
}

// LRCtokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LRCtokenTransactorRaw struct {
	Contract *LRCtokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLRCtoken creates a new instance of LRCtoken, bound to a specific deployed contract.
func NewLRCtoken(address common.Address, backend bind.ContractBackend) (*LRCtoken, error) {
	contract, err := bindLRCtoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LRCtoken{LRCtokenCaller: LRCtokenCaller{contract: contract}, LRCtokenTransactor: LRCtokenTransactor{contract: contract}, LRCtokenFilterer: LRCtokenFilterer{contract: contract}}, nil
}

// NewLRCtokenCaller creates a new read-only instance of LRCtoken, bound to a specific deployed contract.
func NewLRCtokenCaller(address common.Address, caller bind.ContractCaller) (*LRCtokenCaller, error) {
	contract, err := bindLRCtoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LRCtokenCaller{contract: contract}, nil
}

// NewLRCtokenTransactor creates a new write-only instance of LRCtoken, bound to a specific deployed contract.
func NewLRCtokenTransactor(address common.Address, transactor bind.ContractTransactor) (*LRCtokenTransactor, error) {
	contract, err := bindLRCtoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LRCtokenTransactor{contract: contract}, nil
}

// NewLRCtokenFilterer creates a new log filterer instance of LRCtoken, bound to a specific deployed contract.
func NewLRCtokenFilterer(address common.Address, filterer bind.ContractFilterer) (*LRCtokenFilterer, error) {
	contract, err := bindLRCtoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LRCtokenFilterer{contract: contract}, nil
}

// bindLRCtoken binds a generic wrapper to an already deployed contract.
func bindLRCtoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LRCtokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LRCtoken *LRCtokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LRCtoken.Contract.LRCtokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LRCtoken *LRCtokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LRCtoken.Contract.LRCtokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LRCtoken *LRCtokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LRCtoken.Contract.LRCtokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LRCtoken *LRCtokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LRCtoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LRCtoken *LRCtokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LRCtoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LRCtoken *LRCtokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LRCtoken.Contract.contract.Transact(opts, method, params...)
}

// BASERATE is a free data retrieval call binding the contract method 0x41910f90.
//
// Solidity: function BASE_RATE() returns(uint256)
func (_LRCtoken *LRCtokenCaller) BASERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "BASE_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASERATE is a free data retrieval call binding the contract method 0x41910f90.
//
// Solidity: function BASE_RATE() returns(uint256)
func (_LRCtoken *LRCtokenSession) BASERATE() (*big.Int, error) {
	return _LRCtoken.Contract.BASERATE(&_LRCtoken.CallOpts)
}

// BASERATE is a free data retrieval call binding the contract method 0x41910f90.
//
// Solidity: function BASE_RATE() returns(uint256)
func (_LRCtoken *LRCtokenCallerSession) BASERATE() (*big.Int, error) {
	return _LRCtoken.Contract.BASERATE(&_LRCtoken.CallOpts)
}

// BLOCKSPERPHASE is a free data retrieval call binding the contract method 0x2e5ab94f.
//
// Solidity: function BLOCKS_PER_PHASE() returns(uint16)
func (_LRCtoken *LRCtokenCaller) BLOCKSPERPHASE(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "BLOCKS_PER_PHASE")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BLOCKSPERPHASE is a free data retrieval call binding the contract method 0x2e5ab94f.
//
// Solidity: function BLOCKS_PER_PHASE() returns(uint16)
func (_LRCtoken *LRCtokenSession) BLOCKSPERPHASE() (uint16, error) {
	return _LRCtoken.Contract.BLOCKSPERPHASE(&_LRCtoken.CallOpts)
}

// BLOCKSPERPHASE is a free data retrieval call binding the contract method 0x2e5ab94f.
//
// Solidity: function BLOCKS_PER_PHASE() returns(uint16)
func (_LRCtoken *LRCtokenCallerSession) BLOCKSPERPHASE() (uint16, error) {
	return _LRCtoken.Contract.BLOCKSPERPHASE(&_LRCtoken.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() returns(uint256)
func (_LRCtoken *LRCtokenCaller) DECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() returns(uint256)
func (_LRCtoken *LRCtokenSession) DECIMALS() (*big.Int, error) {
	return _LRCtoken.Contract.DECIMALS(&_LRCtoken.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() returns(uint256)
func (_LRCtoken *LRCtokenCallerSession) DECIMALS() (*big.Int, error) {
	return _LRCtoken.Contract.DECIMALS(&_LRCtoken.CallOpts)
}

// GOAL is a free data retrieval call binding the contract method 0xa1bed0be.
//
// Solidity: function GOAL() returns(uint256)
func (_LRCtoken *LRCtokenCaller) GOAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "GOAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GOAL is a free data retrieval call binding the contract method 0xa1bed0be.
//
// Solidity: function GOAL() returns(uint256)
func (_LRCtoken *LRCtokenSession) GOAL() (*big.Int, error) {
	return _LRCtoken.Contract.GOAL(&_LRCtoken.CallOpts)
}

// GOAL is a free data retrieval call binding the contract method 0xa1bed0be.
//
// Solidity: function GOAL() returns(uint256)
func (_LRCtoken *LRCtokenCallerSession) GOAL() (*big.Int, error) {
	return _LRCtoken.Contract.GOAL(&_LRCtoken.CallOpts)
}

// HARDCAP is a free data retrieval call binding the contract method 0x3a03171c.
//
// Solidity: function HARD_CAP() returns(uint256)
func (_LRCtoken *LRCtokenCaller) HARDCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "HARD_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HARDCAP is a free data retrieval call binding the contract method 0x3a03171c.
//
// Solidity: function HARD_CAP() returns(uint256)
func (_LRCtoken *LRCtokenSession) HARDCAP() (*big.Int, error) {
	return _LRCtoken.Contract.HARDCAP(&_LRCtoken.CallOpts)
}

// HARDCAP is a free data retrieval call binding the contract method 0x3a03171c.
//
// Solidity: function HARD_CAP() returns(uint256)
func (_LRCtoken *LRCtokenCallerSession) HARDCAP() (*big.Int, error) {
	return _LRCtoken.Contract.HARDCAP(&_LRCtoken.CallOpts)
}

// MAXUNSOLDRATIO is a free data retrieval call binding the contract method 0x2f969d43.
//
// Solidity: function MAX_UNSOLD_RATIO() returns(uint256)
func (_LRCtoken *LRCtokenCaller) MAXUNSOLDRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "MAX_UNSOLD_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUNSOLDRATIO is a free data retrieval call binding the contract method 0x2f969d43.
//
// Solidity: function MAX_UNSOLD_RATIO() returns(uint256)
func (_LRCtoken *LRCtokenSession) MAXUNSOLDRATIO() (*big.Int, error) {
	return _LRCtoken.Contract.MAXUNSOLDRATIO(&_LRCtoken.CallOpts)
}

// MAXUNSOLDRATIO is a free data retrieval call binding the contract method 0x2f969d43.
//
// Solidity: function MAX_UNSOLD_RATIO() returns(uint256)
func (_LRCtoken *LRCtokenCallerSession) MAXUNSOLDRATIO() (*big.Int, error) {
	return _LRCtoken.Contract.MAXUNSOLDRATIO(&_LRCtoken.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() returns(string)
func (_LRCtoken *LRCtokenCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() returns(string)
func (_LRCtoken *LRCtokenSession) NAME() (string, error) {
	return _LRCtoken.Contract.NAME(&_LRCtoken.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() returns(string)
func (_LRCtoken *LRCtokenCallerSession) NAME() (string, error) {
	return _LRCtoken.Contract.NAME(&_LRCtoken.CallOpts)
}

// NUMOFPHASE is a free data retrieval call binding the contract method 0xdbefe789.
//
// Solidity: function NUM_OF_PHASE() returns(uint256)
func (_LRCtoken *LRCtokenCaller) NUMOFPHASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "NUM_OF_PHASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NUMOFPHASE is a free data retrieval call binding the contract method 0xdbefe789.
//
// Solidity: function NUM_OF_PHASE() returns(uint256)
func (_LRCtoken *LRCtokenSession) NUMOFPHASE() (*big.Int, error) {
	return _LRCtoken.Contract.NUMOFPHASE(&_LRCtoken.CallOpts)
}

// NUMOFPHASE is a free data retrieval call binding the contract method 0xdbefe789.
//
// Solidity: function NUM_OF_PHASE() returns(uint256)
func (_LRCtoken *LRCtokenCallerSession) NUMOFPHASE() (*big.Int, error) {
	return _LRCtoken.Contract.NUMOFPHASE(&_LRCtoken.CallOpts)
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() returns(string)
func (_LRCtoken *LRCtokenCaller) SYMBOL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "SYMBOL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() returns(string)
func (_LRCtoken *LRCtokenSession) SYMBOL() (string, error) {
	return _LRCtoken.Contract.SYMBOL(&_LRCtoken.CallOpts)
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() returns(string)
func (_LRCtoken *LRCtokenCallerSession) SYMBOL() (string, error) {
	return _LRCtoken.Contract.SYMBOL(&_LRCtoken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 remaining)
func (_LRCtoken *LRCtokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 remaining)
func (_LRCtoken *LRCtokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _LRCtoken.Contract.Allowance(&_LRCtoken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 remaining)
func (_LRCtoken *LRCtokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _LRCtoken.Contract.Allowance(&_LRCtoken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_LRCtoken *LRCtokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_LRCtoken *LRCtokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _LRCtoken.Contract.BalanceOf(&_LRCtoken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_LRCtoken *LRCtokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _LRCtoken.Contract.BalanceOf(&_LRCtoken.CallOpts, _owner)
}

// BonusPercentages is a free data retrieval call binding the contract method 0x1e85107c.
//
// Solidity: function bonusPercentages(uint256 ) returns(uint8)
func (_LRCtoken *LRCtokenCaller) BonusPercentages(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "bonusPercentages", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BonusPercentages is a free data retrieval call binding the contract method 0x1e85107c.
//
// Solidity: function bonusPercentages(uint256 ) returns(uint8)
func (_LRCtoken *LRCtokenSession) BonusPercentages(arg0 *big.Int) (uint8, error) {
	return _LRCtoken.Contract.BonusPercentages(&_LRCtoken.CallOpts, arg0)
}

// BonusPercentages is a free data retrieval call binding the contract method 0x1e85107c.
//
// Solidity: function bonusPercentages(uint256 ) returns(uint8)
func (_LRCtoken *LRCtokenCallerSession) BonusPercentages(arg0 *big.Int) (uint8, error) {
	return _LRCtoken.Contract.BonusPercentages(&_LRCtoken.CallOpts, arg0)
}

// Firstblock is a free data retrieval call binding the contract method 0xe85365d5.
//
// Solidity: function firstblock() returns(uint256)
func (_LRCtoken *LRCtokenCaller) Firstblock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "firstblock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Firstblock is a free data retrieval call binding the contract method 0xe85365d5.
//
// Solidity: function firstblock() returns(uint256)
func (_LRCtoken *LRCtokenSession) Firstblock() (*big.Int, error) {
	return _LRCtoken.Contract.Firstblock(&_LRCtoken.CallOpts)
}

// Firstblock is a free data retrieval call binding the contract method 0xe85365d5.
//
// Solidity: function firstblock() returns(uint256)
func (_LRCtoken *LRCtokenCallerSession) Firstblock() (*big.Int, error) {
	return _LRCtoken.Contract.Firstblock(&_LRCtoken.CallOpts)
}

// HardCapReached is a free data retrieval call binding the contract method 0x9762f802.
//
// Solidity: function hardCapReached() returns(bool)
func (_LRCtoken *LRCtokenCaller) HardCapReached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "hardCapReached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HardCapReached is a free data retrieval call binding the contract method 0x9762f802.
//
// Solidity: function hardCapReached() returns(bool)
func (_LRCtoken *LRCtokenSession) HardCapReached() (bool, error) {
	return _LRCtoken.Contract.HardCapReached(&_LRCtoken.CallOpts)
}

// HardCapReached is a free data retrieval call binding the contract method 0x9762f802.
//
// Solidity: function hardCapReached() returns(bool)
func (_LRCtoken *LRCtokenCallerSession) HardCapReached() (bool, error) {
	return _LRCtoken.Contract.HardCapReached(&_LRCtoken.CallOpts)
}

// IssueIndex is a free data retrieval call binding the contract method 0x6a28f828.
//
// Solidity: function issueIndex() returns(uint256)
func (_LRCtoken *LRCtokenCaller) IssueIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "issueIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IssueIndex is a free data retrieval call binding the contract method 0x6a28f828.
//
// Solidity: function issueIndex() returns(uint256)
func (_LRCtoken *LRCtokenSession) IssueIndex() (*big.Int, error) {
	return _LRCtoken.Contract.IssueIndex(&_LRCtoken.CallOpts)
}

// IssueIndex is a free data retrieval call binding the contract method 0x6a28f828.
//
// Solidity: function issueIndex() returns(uint256)
func (_LRCtoken *LRCtokenCallerSession) IssueIndex() (*big.Int, error) {
	return _LRCtoken.Contract.IssueIndex(&_LRCtoken.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() returns(uint256 tokens)
func (_LRCtoken *LRCtokenCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() returns(uint256 tokens)
func (_LRCtoken *LRCtokenSession) Price() (*big.Int, error) {
	return _LRCtoken.Contract.Price(&_LRCtoken.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() returns(uint256 tokens)
func (_LRCtoken *LRCtokenCallerSession) Price() (*big.Int, error) {
	return _LRCtoken.Contract.Price(&_LRCtoken.CallOpts)
}

// SaleDue is a free data retrieval call binding the contract method 0xbea4ae88.
//
// Solidity: function saleDue() returns(bool)
func (_LRCtoken *LRCtokenCaller) SaleDue(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "saleDue")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SaleDue is a free data retrieval call binding the contract method 0xbea4ae88.
//
// Solidity: function saleDue() returns(bool)
func (_LRCtoken *LRCtokenSession) SaleDue() (bool, error) {
	return _LRCtoken.Contract.SaleDue(&_LRCtoken.CallOpts)
}

// SaleDue is a free data retrieval call binding the contract method 0xbea4ae88.
//
// Solidity: function saleDue() returns(bool)
func (_LRCtoken *LRCtokenCallerSession) SaleDue() (bool, error) {
	return _LRCtoken.Contract.SaleDue(&_LRCtoken.CallOpts)
}

// SaleEnded is a free data retrieval call binding the contract method 0x9b8906ae.
//
// Solidity: function saleEnded() returns(bool)
func (_LRCtoken *LRCtokenCaller) SaleEnded(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "saleEnded")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SaleEnded is a free data retrieval call binding the contract method 0x9b8906ae.
//
// Solidity: function saleEnded() returns(bool)
func (_LRCtoken *LRCtokenSession) SaleEnded() (bool, error) {
	return _LRCtoken.Contract.SaleEnded(&_LRCtoken.CallOpts)
}

// SaleEnded is a free data retrieval call binding the contract method 0x9b8906ae.
//
// Solidity: function saleEnded() returns(bool)
func (_LRCtoken *LRCtokenCallerSession) SaleEnded() (bool, error) {
	return _LRCtoken.Contract.SaleEnded(&_LRCtoken.CallOpts)
}

// SaleStarted is a free data retrieval call binding the contract method 0x5c474f9e.
//
// Solidity: function saleStarted() returns(bool)
func (_LRCtoken *LRCtokenCaller) SaleStarted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "saleStarted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SaleStarted is a free data retrieval call binding the contract method 0x5c474f9e.
//
// Solidity: function saleStarted() returns(bool)
func (_LRCtoken *LRCtokenSession) SaleStarted() (bool, error) {
	return _LRCtoken.Contract.SaleStarted(&_LRCtoken.CallOpts)
}

// SaleStarted is a free data retrieval call binding the contract method 0x5c474f9e.
//
// Solidity: function saleStarted() returns(bool)
func (_LRCtoken *LRCtokenCallerSession) SaleStarted() (bool, error) {
	return _LRCtoken.Contract.SaleStarted(&_LRCtoken.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() returns(address)
func (_LRCtoken *LRCtokenCaller) Target(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "target")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() returns(address)
func (_LRCtoken *LRCtokenSession) Target() (common.Address, error) {
	return _LRCtoken.Contract.Target(&_LRCtoken.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() returns(address)
func (_LRCtoken *LRCtokenCallerSession) Target() (common.Address, error) {
	return _LRCtoken.Contract.Target(&_LRCtoken.CallOpts)
}

// TotalEthReceived is a free data retrieval call binding the contract method 0xa9a18dda.
//
// Solidity: function totalEthReceived() returns(uint256)
func (_LRCtoken *LRCtokenCaller) TotalEthReceived(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "totalEthReceived")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEthReceived is a free data retrieval call binding the contract method 0xa9a18dda.
//
// Solidity: function totalEthReceived() returns(uint256)
func (_LRCtoken *LRCtokenSession) TotalEthReceived() (*big.Int, error) {
	return _LRCtoken.Contract.TotalEthReceived(&_LRCtoken.CallOpts)
}

// TotalEthReceived is a free data retrieval call binding the contract method 0xa9a18dda.
//
// Solidity: function totalEthReceived() returns(uint256)
func (_LRCtoken *LRCtokenCallerSession) TotalEthReceived() (*big.Int, error) {
	return _LRCtoken.Contract.TotalEthReceived(&_LRCtoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_LRCtoken *LRCtokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_LRCtoken *LRCtokenSession) TotalSupply() (*big.Int, error) {
	return _LRCtoken.Contract.TotalSupply(&_LRCtoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_LRCtoken *LRCtokenCallerSession) TotalSupply() (*big.Int, error) {
	return _LRCtoken.Contract.TotalSupply(&_LRCtoken.CallOpts)
}

// UnsoldTokenIssued is a free data retrieval call binding the contract method 0x9d0f17c8.
//
// Solidity: function unsoldTokenIssued() returns(bool)
func (_LRCtoken *LRCtokenCaller) UnsoldTokenIssued(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LRCtoken.contract.Call(opts, &out, "unsoldTokenIssued")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UnsoldTokenIssued is a free data retrieval call binding the contract method 0x9d0f17c8.
//
// Solidity: function unsoldTokenIssued() returns(bool)
func (_LRCtoken *LRCtokenSession) UnsoldTokenIssued() (bool, error) {
	return _LRCtoken.Contract.UnsoldTokenIssued(&_LRCtoken.CallOpts)
}

// UnsoldTokenIssued is a free data retrieval call binding the contract method 0x9d0f17c8.
//
// Solidity: function unsoldTokenIssued() returns(bool)
func (_LRCtoken *LRCtokenCallerSession) UnsoldTokenIssued() (bool, error) {
	return _LRCtoken.Contract.UnsoldTokenIssued(&_LRCtoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_LRCtoken *LRCtokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LRCtoken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_LRCtoken *LRCtokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LRCtoken.Contract.Approve(&_LRCtoken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_LRCtoken *LRCtokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LRCtoken.Contract.Approve(&_LRCtoken.TransactOpts, _spender, _value)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_LRCtoken *LRCtokenTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LRCtoken.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_LRCtoken *LRCtokenSession) Close() (*types.Transaction, error) {
	return _LRCtoken.Contract.Close(&_LRCtoken.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_LRCtoken *LRCtokenTransactorSession) Close() (*types.Transaction, error) {
	return _LRCtoken.Contract.Close(&_LRCtoken.TransactOpts)
}

// IssueToken is a paid mutator transaction binding the contract method 0x89034082.
//
// Solidity: function issueToken(address recipient) returns()
func (_LRCtoken *LRCtokenTransactor) IssueToken(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _LRCtoken.contract.Transact(opts, "issueToken", recipient)
}

// IssueToken is a paid mutator transaction binding the contract method 0x89034082.
//
// Solidity: function issueToken(address recipient) returns()
func (_LRCtoken *LRCtokenSession) IssueToken(recipient common.Address) (*types.Transaction, error) {
	return _LRCtoken.Contract.IssueToken(&_LRCtoken.TransactOpts, recipient)
}

// IssueToken is a paid mutator transaction binding the contract method 0x89034082.
//
// Solidity: function issueToken(address recipient) returns()
func (_LRCtoken *LRCtokenTransactorSession) IssueToken(recipient common.Address) (*types.Transaction, error) {
	return _LRCtoken.Contract.IssueToken(&_LRCtoken.TransactOpts, recipient)
}

// Start is a paid mutator transaction binding the contract method 0x95805dad.
//
// Solidity: function start(uint256 _firstblock) returns()
func (_LRCtoken *LRCtokenTransactor) Start(opts *bind.TransactOpts, _firstblock *big.Int) (*types.Transaction, error) {
	return _LRCtoken.contract.Transact(opts, "start", _firstblock)
}

// Start is a paid mutator transaction binding the contract method 0x95805dad.
//
// Solidity: function start(uint256 _firstblock) returns()
func (_LRCtoken *LRCtokenSession) Start(_firstblock *big.Int) (*types.Transaction, error) {
	return _LRCtoken.Contract.Start(&_LRCtoken.TransactOpts, _firstblock)
}

// Start is a paid mutator transaction binding the contract method 0x95805dad.
//
// Solidity: function start(uint256 _firstblock) returns()
func (_LRCtoken *LRCtokenTransactorSession) Start(_firstblock *big.Int) (*types.Transaction, error) {
	return _LRCtoken.Contract.Start(&_LRCtoken.TransactOpts, _firstblock)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_LRCtoken *LRCtokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LRCtoken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_LRCtoken *LRCtokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LRCtoken.Contract.Transfer(&_LRCtoken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_LRCtoken *LRCtokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LRCtoken.Contract.Transfer(&_LRCtoken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_LRCtoken *LRCtokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LRCtoken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_LRCtoken *LRCtokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LRCtoken.Contract.TransferFrom(&_LRCtoken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_LRCtoken *LRCtokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LRCtoken.Contract.TransferFrom(&_LRCtoken.TransactOpts, _from, _to, _value)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_LRCtoken *LRCtokenTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _LRCtoken.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_LRCtoken *LRCtokenSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _LRCtoken.Contract.Fallback(&_LRCtoken.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_LRCtoken *LRCtokenTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _LRCtoken.Contract.Fallback(&_LRCtoken.TransactOpts, calldata)
}

// LRCtokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LRCtoken contract.
type LRCtokenApprovalIterator struct {
	Event *LRCtokenApproval // Event containing the contract specifics and raw log

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
func (it *LRCtokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCtokenApproval)
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
		it.Event = new(LRCtokenApproval)
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
func (it *LRCtokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCtokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCtokenApproval represents a Approval event raised by the LRCtoken contract.
type LRCtokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LRCtoken *LRCtokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LRCtokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LRCtoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LRCtokenApprovalIterator{contract: _LRCtoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LRCtoken *LRCtokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LRCtokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LRCtoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCtokenApproval)
				if err := _LRCtoken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LRCtoken *LRCtokenFilterer) ParseApproval(log types.Log) (*LRCtokenApproval, error) {
	event := new(LRCtokenApproval)
	if err := _LRCtoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LRCtokenInvalidCallerIterator is returned from FilterInvalidCaller and is used to iterate over the raw logs and unpacked data for InvalidCaller events raised by the LRCtoken contract.
type LRCtokenInvalidCallerIterator struct {
	Event *LRCtokenInvalidCaller // Event containing the contract specifics and raw log

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
func (it *LRCtokenInvalidCallerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCtokenInvalidCaller)
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
		it.Event = new(LRCtokenInvalidCaller)
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
func (it *LRCtokenInvalidCallerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCtokenInvalidCallerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCtokenInvalidCaller represents a InvalidCaller event raised by the LRCtoken contract.
type LRCtokenInvalidCaller struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInvalidCaller is a free log retrieval operation binding the contract event 0xcbd9d2e0b97a08f1b662bf4d639e76b32edd97a5d890cafbd2b3cf1b803243a4.
//
// Solidity: event InvalidCaller(address caller)
func (_LRCtoken *LRCtokenFilterer) FilterInvalidCaller(opts *bind.FilterOpts) (*LRCtokenInvalidCallerIterator, error) {

	logs, sub, err := _LRCtoken.contract.FilterLogs(opts, "InvalidCaller")
	if err != nil {
		return nil, err
	}
	return &LRCtokenInvalidCallerIterator{contract: _LRCtoken.contract, event: "InvalidCaller", logs: logs, sub: sub}, nil
}

// WatchInvalidCaller is a free log subscription operation binding the contract event 0xcbd9d2e0b97a08f1b662bf4d639e76b32edd97a5d890cafbd2b3cf1b803243a4.
//
// Solidity: event InvalidCaller(address caller)
func (_LRCtoken *LRCtokenFilterer) WatchInvalidCaller(opts *bind.WatchOpts, sink chan<- *LRCtokenInvalidCaller) (event.Subscription, error) {

	logs, sub, err := _LRCtoken.contract.WatchLogs(opts, "InvalidCaller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCtokenInvalidCaller)
				if err := _LRCtoken.contract.UnpackLog(event, "InvalidCaller", log); err != nil {
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

// ParseInvalidCaller is a log parse operation binding the contract event 0xcbd9d2e0b97a08f1b662bf4d639e76b32edd97a5d890cafbd2b3cf1b803243a4.
//
// Solidity: event InvalidCaller(address caller)
func (_LRCtoken *LRCtokenFilterer) ParseInvalidCaller(log types.Log) (*LRCtokenInvalidCaller, error) {
	event := new(LRCtokenInvalidCaller)
	if err := _LRCtoken.contract.UnpackLog(event, "InvalidCaller", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LRCtokenInvalidStateIterator is returned from FilterInvalidState and is used to iterate over the raw logs and unpacked data for InvalidState events raised by the LRCtoken contract.
type LRCtokenInvalidStateIterator struct {
	Event *LRCtokenInvalidState // Event containing the contract specifics and raw log

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
func (it *LRCtokenInvalidStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCtokenInvalidState)
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
		it.Event = new(LRCtokenInvalidState)
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
func (it *LRCtokenInvalidStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCtokenInvalidStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCtokenInvalidState represents a InvalidState event raised by the LRCtoken contract.
type LRCtokenInvalidState struct {
	Msg []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInvalidState is a free log retrieval operation binding the contract event 0xa24636c18e73457917a92dad223d797b84c2f7a4bdd82892f15a8c4cd9aec1b7.
//
// Solidity: event InvalidState(bytes msg)
func (_LRCtoken *LRCtokenFilterer) FilterInvalidState(opts *bind.FilterOpts) (*LRCtokenInvalidStateIterator, error) {

	logs, sub, err := _LRCtoken.contract.FilterLogs(opts, "InvalidState")
	if err != nil {
		return nil, err
	}
	return &LRCtokenInvalidStateIterator{contract: _LRCtoken.contract, event: "InvalidState", logs: logs, sub: sub}, nil
}

// WatchInvalidState is a free log subscription operation binding the contract event 0xa24636c18e73457917a92dad223d797b84c2f7a4bdd82892f15a8c4cd9aec1b7.
//
// Solidity: event InvalidState(bytes msg)
func (_LRCtoken *LRCtokenFilterer) WatchInvalidState(opts *bind.WatchOpts, sink chan<- *LRCtokenInvalidState) (event.Subscription, error) {

	logs, sub, err := _LRCtoken.contract.WatchLogs(opts, "InvalidState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCtokenInvalidState)
				if err := _LRCtoken.contract.UnpackLog(event, "InvalidState", log); err != nil {
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

// ParseInvalidState is a log parse operation binding the contract event 0xa24636c18e73457917a92dad223d797b84c2f7a4bdd82892f15a8c4cd9aec1b7.
//
// Solidity: event InvalidState(bytes msg)
func (_LRCtoken *LRCtokenFilterer) ParseInvalidState(log types.Log) (*LRCtokenInvalidState, error) {
	event := new(LRCtokenInvalidState)
	if err := _LRCtoken.contract.UnpackLog(event, "InvalidState", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LRCtokenIssueIterator is returned from FilterIssue and is used to iterate over the raw logs and unpacked data for Issue events raised by the LRCtoken contract.
type LRCtokenIssueIterator struct {
	Event *LRCtokenIssue // Event containing the contract specifics and raw log

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
func (it *LRCtokenIssueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCtokenIssue)
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
		it.Event = new(LRCtokenIssue)
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
func (it *LRCtokenIssueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCtokenIssueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCtokenIssue represents a Issue event raised by the LRCtoken contract.
type LRCtokenIssue struct {
	IssueIndex  *big.Int
	Addr        common.Address
	EthAmount   *big.Int
	TokenAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterIssue is a free log retrieval operation binding the contract event 0xe316e9c07bf6ee91102f762d73f95b6cab9dcc157278814c7408906855c6a3a6.
//
// Solidity: event Issue(uint256 issueIndex, address addr, uint256 ethAmount, uint256 tokenAmount)
func (_LRCtoken *LRCtokenFilterer) FilterIssue(opts *bind.FilterOpts) (*LRCtokenIssueIterator, error) {

	logs, sub, err := _LRCtoken.contract.FilterLogs(opts, "Issue")
	if err != nil {
		return nil, err
	}
	return &LRCtokenIssueIterator{contract: _LRCtoken.contract, event: "Issue", logs: logs, sub: sub}, nil
}

// WatchIssue is a free log subscription operation binding the contract event 0xe316e9c07bf6ee91102f762d73f95b6cab9dcc157278814c7408906855c6a3a6.
//
// Solidity: event Issue(uint256 issueIndex, address addr, uint256 ethAmount, uint256 tokenAmount)
func (_LRCtoken *LRCtokenFilterer) WatchIssue(opts *bind.WatchOpts, sink chan<- *LRCtokenIssue) (event.Subscription, error) {

	logs, sub, err := _LRCtoken.contract.WatchLogs(opts, "Issue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCtokenIssue)
				if err := _LRCtoken.contract.UnpackLog(event, "Issue", log); err != nil {
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

// ParseIssue is a log parse operation binding the contract event 0xe316e9c07bf6ee91102f762d73f95b6cab9dcc157278814c7408906855c6a3a6.
//
// Solidity: event Issue(uint256 issueIndex, address addr, uint256 ethAmount, uint256 tokenAmount)
func (_LRCtoken *LRCtokenFilterer) ParseIssue(log types.Log) (*LRCtokenIssue, error) {
	event := new(LRCtokenIssue)
	if err := _LRCtoken.contract.UnpackLog(event, "Issue", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LRCtokenSaleEndedIterator is returned from FilterSaleEnded and is used to iterate over the raw logs and unpacked data for SaleEnded events raised by the LRCtoken contract.
type LRCtokenSaleEndedIterator struct {
	Event *LRCtokenSaleEnded // Event containing the contract specifics and raw log

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
func (it *LRCtokenSaleEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCtokenSaleEnded)
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
		it.Event = new(LRCtokenSaleEnded)
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
func (it *LRCtokenSaleEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCtokenSaleEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCtokenSaleEnded represents a SaleEnded event raised by the LRCtoken contract.
type LRCtokenSaleEnded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSaleEnded is a free log retrieval operation binding the contract event 0x0bd8a3eb532e5fbcd3f5b00335f0fb42fdc11969e9af0fab7c9e71a36ae0d31a.
//
// Solidity: event SaleEnded()
func (_LRCtoken *LRCtokenFilterer) FilterSaleEnded(opts *bind.FilterOpts) (*LRCtokenSaleEndedIterator, error) {

	logs, sub, err := _LRCtoken.contract.FilterLogs(opts, "SaleEnded")
	if err != nil {
		return nil, err
	}
	return &LRCtokenSaleEndedIterator{contract: _LRCtoken.contract, event: "SaleEnded", logs: logs, sub: sub}, nil
}

// WatchSaleEnded is a free log subscription operation binding the contract event 0x0bd8a3eb532e5fbcd3f5b00335f0fb42fdc11969e9af0fab7c9e71a36ae0d31a.
//
// Solidity: event SaleEnded()
func (_LRCtoken *LRCtokenFilterer) WatchSaleEnded(opts *bind.WatchOpts, sink chan<- *LRCtokenSaleEnded) (event.Subscription, error) {

	logs, sub, err := _LRCtoken.contract.WatchLogs(opts, "SaleEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCtokenSaleEnded)
				if err := _LRCtoken.contract.UnpackLog(event, "SaleEnded", log); err != nil {
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

// ParseSaleEnded is a log parse operation binding the contract event 0x0bd8a3eb532e5fbcd3f5b00335f0fb42fdc11969e9af0fab7c9e71a36ae0d31a.
//
// Solidity: event SaleEnded()
func (_LRCtoken *LRCtokenFilterer) ParseSaleEnded(log types.Log) (*LRCtokenSaleEnded, error) {
	event := new(LRCtokenSaleEnded)
	if err := _LRCtoken.contract.UnpackLog(event, "SaleEnded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LRCtokenSaleFailedIterator is returned from FilterSaleFailed and is used to iterate over the raw logs and unpacked data for SaleFailed events raised by the LRCtoken contract.
type LRCtokenSaleFailedIterator struct {
	Event *LRCtokenSaleFailed // Event containing the contract specifics and raw log

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
func (it *LRCtokenSaleFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCtokenSaleFailed)
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
		it.Event = new(LRCtokenSaleFailed)
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
func (it *LRCtokenSaleFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCtokenSaleFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCtokenSaleFailed represents a SaleFailed event raised by the LRCtoken contract.
type LRCtokenSaleFailed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSaleFailed is a free log retrieval operation binding the contract event 0x100c98fe43bef515268f854995b9fe23e0d72dbe3f2726094d9f1864ab4afde2.
//
// Solidity: event SaleFailed()
func (_LRCtoken *LRCtokenFilterer) FilterSaleFailed(opts *bind.FilterOpts) (*LRCtokenSaleFailedIterator, error) {

	logs, sub, err := _LRCtoken.contract.FilterLogs(opts, "SaleFailed")
	if err != nil {
		return nil, err
	}
	return &LRCtokenSaleFailedIterator{contract: _LRCtoken.contract, event: "SaleFailed", logs: logs, sub: sub}, nil
}

// WatchSaleFailed is a free log subscription operation binding the contract event 0x100c98fe43bef515268f854995b9fe23e0d72dbe3f2726094d9f1864ab4afde2.
//
// Solidity: event SaleFailed()
func (_LRCtoken *LRCtokenFilterer) WatchSaleFailed(opts *bind.WatchOpts, sink chan<- *LRCtokenSaleFailed) (event.Subscription, error) {

	logs, sub, err := _LRCtoken.contract.WatchLogs(opts, "SaleFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCtokenSaleFailed)
				if err := _LRCtoken.contract.UnpackLog(event, "SaleFailed", log); err != nil {
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

// ParseSaleFailed is a log parse operation binding the contract event 0x100c98fe43bef515268f854995b9fe23e0d72dbe3f2726094d9f1864ab4afde2.
//
// Solidity: event SaleFailed()
func (_LRCtoken *LRCtokenFilterer) ParseSaleFailed(log types.Log) (*LRCtokenSaleFailed, error) {
	event := new(LRCtokenSaleFailed)
	if err := _LRCtoken.contract.UnpackLog(event, "SaleFailed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LRCtokenSaleStartedIterator is returned from FilterSaleStarted and is used to iterate over the raw logs and unpacked data for SaleStarted events raised by the LRCtoken contract.
type LRCtokenSaleStartedIterator struct {
	Event *LRCtokenSaleStarted // Event containing the contract specifics and raw log

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
func (it *LRCtokenSaleStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCtokenSaleStarted)
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
		it.Event = new(LRCtokenSaleStarted)
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
func (it *LRCtokenSaleStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCtokenSaleStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCtokenSaleStarted represents a SaleStarted event raised by the LRCtoken contract.
type LRCtokenSaleStarted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSaleStarted is a free log retrieval operation binding the contract event 0x912ee23dde46ec889d6748212cce445d667f7041597691dc89e8549ad8bc0acb.
//
// Solidity: event SaleStarted()
func (_LRCtoken *LRCtokenFilterer) FilterSaleStarted(opts *bind.FilterOpts) (*LRCtokenSaleStartedIterator, error) {

	logs, sub, err := _LRCtoken.contract.FilterLogs(opts, "SaleStarted")
	if err != nil {
		return nil, err
	}
	return &LRCtokenSaleStartedIterator{contract: _LRCtoken.contract, event: "SaleStarted", logs: logs, sub: sub}, nil
}

// WatchSaleStarted is a free log subscription operation binding the contract event 0x912ee23dde46ec889d6748212cce445d667f7041597691dc89e8549ad8bc0acb.
//
// Solidity: event SaleStarted()
func (_LRCtoken *LRCtokenFilterer) WatchSaleStarted(opts *bind.WatchOpts, sink chan<- *LRCtokenSaleStarted) (event.Subscription, error) {

	logs, sub, err := _LRCtoken.contract.WatchLogs(opts, "SaleStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCtokenSaleStarted)
				if err := _LRCtoken.contract.UnpackLog(event, "SaleStarted", log); err != nil {
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

// ParseSaleStarted is a log parse operation binding the contract event 0x912ee23dde46ec889d6748212cce445d667f7041597691dc89e8549ad8bc0acb.
//
// Solidity: event SaleStarted()
func (_LRCtoken *LRCtokenFilterer) ParseSaleStarted(log types.Log) (*LRCtokenSaleStarted, error) {
	event := new(LRCtokenSaleStarted)
	if err := _LRCtoken.contract.UnpackLog(event, "SaleStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LRCtokenSaleSucceededIterator is returned from FilterSaleSucceeded and is used to iterate over the raw logs and unpacked data for SaleSucceeded events raised by the LRCtoken contract.
type LRCtokenSaleSucceededIterator struct {
	Event *LRCtokenSaleSucceeded // Event containing the contract specifics and raw log

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
func (it *LRCtokenSaleSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCtokenSaleSucceeded)
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
		it.Event = new(LRCtokenSaleSucceeded)
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
func (it *LRCtokenSaleSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCtokenSaleSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCtokenSaleSucceeded represents a SaleSucceeded event raised by the LRCtoken contract.
type LRCtokenSaleSucceeded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSaleSucceeded is a free log retrieval operation binding the contract event 0x0151fbf6d2def4666ab0f87412daa4ac6a67e9fa86a50cfbd7b36b16d72705d9.
//
// Solidity: event SaleSucceeded()
func (_LRCtoken *LRCtokenFilterer) FilterSaleSucceeded(opts *bind.FilterOpts) (*LRCtokenSaleSucceededIterator, error) {

	logs, sub, err := _LRCtoken.contract.FilterLogs(opts, "SaleSucceeded")
	if err != nil {
		return nil, err
	}
	return &LRCtokenSaleSucceededIterator{contract: _LRCtoken.contract, event: "SaleSucceeded", logs: logs, sub: sub}, nil
}

// WatchSaleSucceeded is a free log subscription operation binding the contract event 0x0151fbf6d2def4666ab0f87412daa4ac6a67e9fa86a50cfbd7b36b16d72705d9.
//
// Solidity: event SaleSucceeded()
func (_LRCtoken *LRCtokenFilterer) WatchSaleSucceeded(opts *bind.WatchOpts, sink chan<- *LRCtokenSaleSucceeded) (event.Subscription, error) {

	logs, sub, err := _LRCtoken.contract.WatchLogs(opts, "SaleSucceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCtokenSaleSucceeded)
				if err := _LRCtoken.contract.UnpackLog(event, "SaleSucceeded", log); err != nil {
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

// ParseSaleSucceeded is a log parse operation binding the contract event 0x0151fbf6d2def4666ab0f87412daa4ac6a67e9fa86a50cfbd7b36b16d72705d9.
//
// Solidity: event SaleSucceeded()
func (_LRCtoken *LRCtokenFilterer) ParseSaleSucceeded(log types.Log) (*LRCtokenSaleSucceeded, error) {
	event := new(LRCtokenSaleSucceeded)
	if err := _LRCtoken.contract.UnpackLog(event, "SaleSucceeded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LRCtokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LRCtoken contract.
type LRCtokenTransferIterator struct {
	Event *LRCtokenTransfer // Event containing the contract specifics and raw log

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
func (it *LRCtokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCtokenTransfer)
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
		it.Event = new(LRCtokenTransfer)
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
func (it *LRCtokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCtokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCtokenTransfer represents a Transfer event raised by the LRCtoken contract.
type LRCtokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LRCtoken *LRCtokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LRCtokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LRCtoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LRCtokenTransferIterator{contract: _LRCtoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LRCtoken *LRCtokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LRCtokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LRCtoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCtokenTransfer)
				if err := _LRCtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_LRCtoken *LRCtokenFilterer) ParseTransfer(log types.Log) (*LRCtokenTransfer, error) {
	event := new(LRCtokenTransfer)
	if err := _LRCtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
