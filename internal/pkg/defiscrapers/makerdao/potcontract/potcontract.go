// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package potcontract

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

// PotcontractABI is the input ABI used to generate the binding from.
const PotcontractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"Pie\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"drip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tmp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dsr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"join\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rho\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Potcontract is an auto generated Go binding around an Ethereum contract.
type Potcontract struct {
	PotcontractCaller     // Read-only binding to the contract
	PotcontractTransactor // Write-only binding to the contract
	PotcontractFilterer   // Log filterer for contract events
}

// PotcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PotcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PotcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PotcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PotcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PotcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PotcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PotcontractSession struct {
	Contract     *Potcontract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PotcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PotcontractCallerSession struct {
	Contract *PotcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PotcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PotcontractTransactorSession struct {
	Contract     *PotcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PotcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PotcontractRaw struct {
	Contract *Potcontract // Generic contract binding to access the raw methods on
}

// PotcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PotcontractCallerRaw struct {
	Contract *PotcontractCaller // Generic read-only contract binding to access the raw methods on
}

// PotcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PotcontractTransactorRaw struct {
	Contract *PotcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPotcontract creates a new instance of Potcontract, bound to a specific deployed contract.
func NewPotcontract(address common.Address, backend bind.ContractBackend) (*Potcontract, error) {
	contract, err := bindPotcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Potcontract{PotcontractCaller: PotcontractCaller{contract: contract}, PotcontractTransactor: PotcontractTransactor{contract: contract}, PotcontractFilterer: PotcontractFilterer{contract: contract}}, nil
}

// NewPotcontractCaller creates a new read-only instance of Potcontract, bound to a specific deployed contract.
func NewPotcontractCaller(address common.Address, caller bind.ContractCaller) (*PotcontractCaller, error) {
	contract, err := bindPotcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PotcontractCaller{contract: contract}, nil
}

// NewPotcontractTransactor creates a new write-only instance of Potcontract, bound to a specific deployed contract.
func NewPotcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*PotcontractTransactor, error) {
	contract, err := bindPotcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PotcontractTransactor{contract: contract}, nil
}

// NewPotcontractFilterer creates a new log filterer instance of Potcontract, bound to a specific deployed contract.
func NewPotcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*PotcontractFilterer, error) {
	contract, err := bindPotcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PotcontractFilterer{contract: contract}, nil
}

// bindPotcontract binds a generic wrapper to an already deployed contract.
func bindPotcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PotcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Potcontract *PotcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Potcontract.Contract.PotcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Potcontract *PotcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Potcontract.Contract.PotcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Potcontract *PotcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Potcontract.Contract.PotcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Potcontract *PotcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Potcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Potcontract *PotcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Potcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Potcontract *PotcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Potcontract.Contract.contract.Transact(opts, method, params...)
}

// Pie is a free data retrieval call binding the contract method 0x2c69ed58.
//
// Solidity: function Pie() view returns(uint256)
func (_Potcontract *PotcontractCaller) Pie(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Potcontract.contract.Call(opts, &out, "Pie")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pie is a free data retrieval call binding the contract method 0x2c69ed58.
//
// Solidity: function Pie() view returns(uint256)
func (_Potcontract *PotcontractSession) Pie() (*big.Int, error) {
	return _Potcontract.Contract.Pie(&_Potcontract.CallOpts)
}

// Pie is a free data retrieval call binding the contract method 0x2c69ed58.
//
// Solidity: function Pie() view returns(uint256)
func (_Potcontract *PotcontractCallerSession) Pie() (*big.Int, error) {
	return _Potcontract.Contract.Pie(&_Potcontract.CallOpts)
}

// Chi is a free data retrieval call binding the contract method 0xc92aecc4.
//
// Solidity: function chi() view returns(uint256)
func (_Potcontract *PotcontractCaller) Chi(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Potcontract.contract.Call(opts, &out, "chi")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Chi is a free data retrieval call binding the contract method 0xc92aecc4.
//
// Solidity: function chi() view returns(uint256)
func (_Potcontract *PotcontractSession) Chi() (*big.Int, error) {
	return _Potcontract.Contract.Chi(&_Potcontract.CallOpts)
}

// Chi is a free data retrieval call binding the contract method 0xc92aecc4.
//
// Solidity: function chi() view returns(uint256)
func (_Potcontract *PotcontractCallerSession) Chi() (*big.Int, error) {
	return _Potcontract.Contract.Chi(&_Potcontract.CallOpts)
}

// Dsr is a free data retrieval call binding the contract method 0x487bf082.
//
// Solidity: function dsr() view returns(uint256)
func (_Potcontract *PotcontractCaller) Dsr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Potcontract.contract.Call(opts, &out, "dsr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dsr is a free data retrieval call binding the contract method 0x487bf082.
//
// Solidity: function dsr() view returns(uint256)
func (_Potcontract *PotcontractSession) Dsr() (*big.Int, error) {
	return _Potcontract.Contract.Dsr(&_Potcontract.CallOpts)
}

// Dsr is a free data retrieval call binding the contract method 0x487bf082.
//
// Solidity: function dsr() view returns(uint256)
func (_Potcontract *PotcontractCallerSession) Dsr() (*big.Int, error) {
	return _Potcontract.Contract.Dsr(&_Potcontract.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Potcontract *PotcontractCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Potcontract.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Potcontract *PotcontractSession) Live() (*big.Int, error) {
	return _Potcontract.Contract.Live(&_Potcontract.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Potcontract *PotcontractCallerSession) Live() (*big.Int, error) {
	return _Potcontract.Contract.Live(&_Potcontract.CallOpts)
}

// Rho is a free data retrieval call binding the contract method 0x20aba08b.
//
// Solidity: function rho() view returns(uint256)
func (_Potcontract *PotcontractCaller) Rho(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Potcontract.contract.Call(opts, &out, "rho")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rho is a free data retrieval call binding the contract method 0x20aba08b.
//
// Solidity: function rho() view returns(uint256)
func (_Potcontract *PotcontractSession) Rho() (*big.Int, error) {
	return _Potcontract.Contract.Rho(&_Potcontract.CallOpts)
}

// Rho is a free data retrieval call binding the contract method 0x20aba08b.
//
// Solidity: function rho() view returns(uint256)
func (_Potcontract *PotcontractCallerSession) Rho() (*big.Int, error) {
	return _Potcontract.Contract.Rho(&_Potcontract.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Potcontract *PotcontractCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Potcontract.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Potcontract *PotcontractSession) Vat() (common.Address, error) {
	return _Potcontract.Contract.Vat(&_Potcontract.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Potcontract *PotcontractCallerSession) Vat() (common.Address, error) {
	return _Potcontract.Contract.Vat(&_Potcontract.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Potcontract *PotcontractCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Potcontract.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Potcontract *PotcontractSession) Vow() (common.Address, error) {
	return _Potcontract.Contract.Vow(&_Potcontract.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Potcontract *PotcontractCallerSession) Vow() (common.Address, error) {
	return _Potcontract.Contract.Vow(&_Potcontract.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Potcontract *PotcontractCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Potcontract.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Potcontract *PotcontractSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Potcontract.Contract.Wards(&_Potcontract.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Potcontract *PotcontractCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Potcontract.Contract.Wards(&_Potcontract.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Potcontract *PotcontractTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Potcontract.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Potcontract *PotcontractSession) Cage() (*types.Transaction, error) {
	return _Potcontract.Contract.Cage(&_Potcontract.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Potcontract *PotcontractTransactorSession) Cage() (*types.Transaction, error) {
	return _Potcontract.Contract.Cage(&_Potcontract.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Potcontract *PotcontractTransactor) Deny(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Potcontract.contract.Transact(opts, "deny", guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Potcontract *PotcontractSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _Potcontract.Contract.Deny(&_Potcontract.TransactOpts, guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Potcontract *PotcontractTransactorSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _Potcontract.Contract.Deny(&_Potcontract.TransactOpts, guy)
}

// Drip is a paid mutator transaction binding the contract method 0x9f678cca.
//
// Solidity: function drip() returns(uint256 tmp)
func (_Potcontract *PotcontractTransactor) Drip(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Potcontract.contract.Transact(opts, "drip")
}

// Drip is a paid mutator transaction binding the contract method 0x9f678cca.
//
// Solidity: function drip() returns(uint256 tmp)
func (_Potcontract *PotcontractSession) Drip() (*types.Transaction, error) {
	return _Potcontract.Contract.Drip(&_Potcontract.TransactOpts)
}

// Drip is a paid mutator transaction binding the contract method 0x9f678cca.
//
// Solidity: function drip() returns(uint256 tmp)
func (_Potcontract *PotcontractTransactorSession) Drip() (*types.Transaction, error) {
	return _Potcontract.Contract.Drip(&_Potcontract.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 wad) returns()
func (_Potcontract *PotcontractTransactor) Exit(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Potcontract.contract.Transact(opts, "exit", wad)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 wad) returns()
func (_Potcontract *PotcontractSession) Exit(wad *big.Int) (*types.Transaction, error) {
	return _Potcontract.Contract.Exit(&_Potcontract.TransactOpts, wad)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 wad) returns()
func (_Potcontract *PotcontractTransactorSession) Exit(wad *big.Int) (*types.Transaction, error) {
	return _Potcontract.Contract.Exit(&_Potcontract.TransactOpts, wad)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Potcontract *PotcontractTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Potcontract.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Potcontract *PotcontractSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Potcontract.Contract.File(&_Potcontract.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Potcontract *PotcontractTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Potcontract.Contract.File(&_Potcontract.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address addr) returns()
func (_Potcontract *PotcontractTransactor) File0(opts *bind.TransactOpts, what [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Potcontract.contract.Transact(opts, "file0", what, addr)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address addr) returns()
func (_Potcontract *PotcontractSession) File0(what [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Potcontract.Contract.File0(&_Potcontract.TransactOpts, what, addr)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address addr) returns()
func (_Potcontract *PotcontractTransactorSession) File0(what [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Potcontract.Contract.File0(&_Potcontract.TransactOpts, what, addr)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(uint256 wad) returns()
func (_Potcontract *PotcontractTransactor) Join(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Potcontract.contract.Transact(opts, "join", wad)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(uint256 wad) returns()
func (_Potcontract *PotcontractSession) Join(wad *big.Int) (*types.Transaction, error) {
	return _Potcontract.Contract.Join(&_Potcontract.TransactOpts, wad)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(uint256 wad) returns()
func (_Potcontract *PotcontractTransactorSession) Join(wad *big.Int) (*types.Transaction, error) {
	return _Potcontract.Contract.Join(&_Potcontract.TransactOpts, wad)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Potcontract *PotcontractTransactor) Rely(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Potcontract.contract.Transact(opts, "rely", guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Potcontract *PotcontractSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _Potcontract.Contract.Rely(&_Potcontract.TransactOpts, guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Potcontract *PotcontractTransactorSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _Potcontract.Contract.Rely(&_Potcontract.TransactOpts, guy)
}
