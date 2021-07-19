// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package jugcontract

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

// JugcontractABI is the input ABI used to generate the binding from.
const JugcontractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"drip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"duty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rho\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Jugcontract is an auto generated Go binding around an Ethereum contract.
type Jugcontract struct {
	JugcontractCaller     // Read-only binding to the contract
	JugcontractTransactor // Write-only binding to the contract
	JugcontractFilterer   // Log filterer for contract events
}

// JugcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type JugcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JugcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JugcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JugcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JugcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JugcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JugcontractSession struct {
	Contract     *Jugcontract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JugcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JugcontractCallerSession struct {
	Contract *JugcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// JugcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JugcontractTransactorSession struct {
	Contract     *JugcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// JugcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type JugcontractRaw struct {
	Contract *Jugcontract // Generic contract binding to access the raw methods on
}

// JugcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JugcontractCallerRaw struct {
	Contract *JugcontractCaller // Generic read-only contract binding to access the raw methods on
}

// JugcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JugcontractTransactorRaw struct {
	Contract *JugcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJugcontract creates a new instance of Jugcontract, bound to a specific deployed contract.
func NewJugcontract(address common.Address, backend bind.ContractBackend) (*Jugcontract, error) {
	contract, err := bindJugcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Jugcontract{JugcontractCaller: JugcontractCaller{contract: contract}, JugcontractTransactor: JugcontractTransactor{contract: contract}, JugcontractFilterer: JugcontractFilterer{contract: contract}}, nil
}

// NewJugcontractCaller creates a new read-only instance of Jugcontract, bound to a specific deployed contract.
func NewJugcontractCaller(address common.Address, caller bind.ContractCaller) (*JugcontractCaller, error) {
	contract, err := bindJugcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JugcontractCaller{contract: contract}, nil
}

// NewJugcontractTransactor creates a new write-only instance of Jugcontract, bound to a specific deployed contract.
func NewJugcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*JugcontractTransactor, error) {
	contract, err := bindJugcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JugcontractTransactor{contract: contract}, nil
}

// NewJugcontractFilterer creates a new log filterer instance of Jugcontract, bound to a specific deployed contract.
func NewJugcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*JugcontractFilterer, error) {
	contract, err := bindJugcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JugcontractFilterer{contract: contract}, nil
}

// bindJugcontract binds a generic wrapper to an already deployed contract.
func bindJugcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JugcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jugcontract *JugcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jugcontract.Contract.JugcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jugcontract *JugcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jugcontract.Contract.JugcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jugcontract *JugcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jugcontract.Contract.JugcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jugcontract *JugcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jugcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jugcontract *JugcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jugcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jugcontract *JugcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jugcontract.Contract.contract.Transact(opts, method, params...)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_Jugcontract *JugcontractCaller) Base(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jugcontract.contract.Call(opts, &out, "base")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_Jugcontract *JugcontractSession) Base() (*big.Int, error) {
	return _Jugcontract.Contract.Base(&_Jugcontract.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_Jugcontract *JugcontractCallerSession) Base() (*big.Int, error) {
	return _Jugcontract.Contract.Base(&_Jugcontract.CallOpts)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 duty, uint256 rho)
func (_Jugcontract *JugcontractCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Duty *big.Int
	Rho  *big.Int
}, error) {
	var out []interface{}
	err := _Jugcontract.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Duty *big.Int
		Rho  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Duty = out[0].(*big.Int)
	outstruct.Rho = out[1].(*big.Int)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 duty, uint256 rho)
func (_Jugcontract *JugcontractSession) Ilks(arg0 [32]byte) (struct {
	Duty *big.Int
	Rho  *big.Int
}, error) {
	return _Jugcontract.Contract.Ilks(&_Jugcontract.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 duty, uint256 rho)
func (_Jugcontract *JugcontractCallerSession) Ilks(arg0 [32]byte) (struct {
	Duty *big.Int
	Rho  *big.Int
}, error) {
	return _Jugcontract.Contract.Ilks(&_Jugcontract.CallOpts, arg0)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Jugcontract *JugcontractCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jugcontract.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Jugcontract *JugcontractSession) Vat() (common.Address, error) {
	return _Jugcontract.Contract.Vat(&_Jugcontract.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Jugcontract *JugcontractCallerSession) Vat() (common.Address, error) {
	return _Jugcontract.Contract.Vat(&_Jugcontract.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Jugcontract *JugcontractCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jugcontract.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Jugcontract *JugcontractSession) Vow() (common.Address, error) {
	return _Jugcontract.Contract.Vow(&_Jugcontract.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Jugcontract *JugcontractCallerSession) Vow() (common.Address, error) {
	return _Jugcontract.Contract.Vow(&_Jugcontract.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Jugcontract *JugcontractCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Jugcontract.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Jugcontract *JugcontractSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Jugcontract.Contract.Wards(&_Jugcontract.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Jugcontract *JugcontractCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Jugcontract.Contract.Wards(&_Jugcontract.CallOpts, arg0)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Jugcontract *JugcontractTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Jugcontract.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Jugcontract *JugcontractSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Jugcontract.Contract.Deny(&_Jugcontract.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Jugcontract *JugcontractTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Jugcontract.Contract.Deny(&_Jugcontract.TransactOpts, usr)
}

// Drip is a paid mutator transaction binding the contract method 0x44e2a5a8.
//
// Solidity: function drip(bytes32 ilk) returns(uint256 rate)
func (_Jugcontract *JugcontractTransactor) Drip(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Jugcontract.contract.Transact(opts, "drip", ilk)
}

// Drip is a paid mutator transaction binding the contract method 0x44e2a5a8.
//
// Solidity: function drip(bytes32 ilk) returns(uint256 rate)
func (_Jugcontract *JugcontractSession) Drip(ilk [32]byte) (*types.Transaction, error) {
	return _Jugcontract.Contract.Drip(&_Jugcontract.TransactOpts, ilk)
}

// Drip is a paid mutator transaction binding the contract method 0x44e2a5a8.
//
// Solidity: function drip(bytes32 ilk) returns(uint256 rate)
func (_Jugcontract *JugcontractTransactorSession) Drip(ilk [32]byte) (*types.Transaction, error) {
	return _Jugcontract.Contract.Drip(&_Jugcontract.TransactOpts, ilk)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Jugcontract *JugcontractTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jugcontract.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Jugcontract *JugcontractSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jugcontract.Contract.File(&_Jugcontract.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Jugcontract *JugcontractTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jugcontract.Contract.File(&_Jugcontract.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Jugcontract *JugcontractTransactor) File0(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jugcontract.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Jugcontract *JugcontractSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jugcontract.Contract.File0(&_Jugcontract.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Jugcontract *JugcontractTransactorSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jugcontract.Contract.File0(&_Jugcontract.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Jugcontract *JugcontractTransactor) File1(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Jugcontract.contract.Transact(opts, "file1", what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Jugcontract *JugcontractSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Jugcontract.Contract.File1(&_Jugcontract.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Jugcontract *JugcontractTransactorSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Jugcontract.Contract.File1(&_Jugcontract.TransactOpts, what, data)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Jugcontract *JugcontractTransactor) Init(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Jugcontract.contract.Transact(opts, "init", ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Jugcontract *JugcontractSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _Jugcontract.Contract.Init(&_Jugcontract.TransactOpts, ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Jugcontract *JugcontractTransactorSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _Jugcontract.Contract.Init(&_Jugcontract.TransactOpts, ilk)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Jugcontract *JugcontractTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Jugcontract.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Jugcontract *JugcontractSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Jugcontract.Contract.Rely(&_Jugcontract.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Jugcontract *JugcontractTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Jugcontract.Contract.Rely(&_Jugcontract.TransactOpts, usr)
}
