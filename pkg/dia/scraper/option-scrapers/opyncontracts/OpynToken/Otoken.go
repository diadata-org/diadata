// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Otoken

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

// OtokenABI is the input ABI used to generate the binding from.
const OtokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnOtoken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expiryTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOtokenDetails\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressBook\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_strikeAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isPut\",\"type\":\"bool\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintOtoken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strikeAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strikePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Otoken is an auto generated Go binding around an Ethereum contract.
type Otoken struct {
	OtokenCaller     // Read-only binding to the contract
	OtokenTransactor // Write-only binding to the contract
	OtokenFilterer   // Log filterer for contract events
}

// OtokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type OtokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OtokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OtokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OtokenSession struct {
	Contract     *Otoken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OtokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OtokenCallerSession struct {
	Contract *OtokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OtokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OtokenTransactorSession struct {
	Contract     *OtokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OtokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type OtokenRaw struct {
	Contract *Otoken // Generic contract binding to access the raw methods on
}

// OtokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OtokenCallerRaw struct {
	Contract *OtokenCaller // Generic read-only contract binding to access the raw methods on
}

// OtokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OtokenTransactorRaw struct {
	Contract *OtokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOtoken creates a new instance of Otoken, bound to a specific deployed contract.
func NewOtoken(address common.Address, backend bind.ContractBackend) (*Otoken, error) {
	contract, err := bindOtoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Otoken{OtokenCaller: OtokenCaller{contract: contract}, OtokenTransactor: OtokenTransactor{contract: contract}, OtokenFilterer: OtokenFilterer{contract: contract}}, nil
}

// NewOtokenCaller creates a new read-only instance of Otoken, bound to a specific deployed contract.
func NewOtokenCaller(address common.Address, caller bind.ContractCaller) (*OtokenCaller, error) {
	contract, err := bindOtoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OtokenCaller{contract: contract}, nil
}

// NewOtokenTransactor creates a new write-only instance of Otoken, bound to a specific deployed contract.
func NewOtokenTransactor(address common.Address, transactor bind.ContractTransactor) (*OtokenTransactor, error) {
	contract, err := bindOtoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OtokenTransactor{contract: contract}, nil
}

// NewOtokenFilterer creates a new log filterer instance of Otoken, bound to a specific deployed contract.
func NewOtokenFilterer(address common.Address, filterer bind.ContractFilterer) (*OtokenFilterer, error) {
	contract, err := bindOtoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OtokenFilterer{contract: contract}, nil
}

// bindOtoken binds a generic wrapper to an already deployed contract.
func bindOtoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OtokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Otoken *OtokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Otoken.Contract.OtokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Otoken *OtokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Otoken.Contract.OtokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Otoken *OtokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Otoken.Contract.OtokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Otoken *OtokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Otoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Otoken *OtokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Otoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Otoken *OtokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Otoken.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Otoken *OtokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Otoken *OtokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Otoken.Contract.DOMAINSEPARATOR(&_Otoken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Otoken *OtokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Otoken.Contract.DOMAINSEPARATOR(&_Otoken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Otoken *OtokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Otoken *OtokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Otoken.Contract.Allowance(&_Otoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Otoken *OtokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Otoken.Contract.Allowance(&_Otoken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Otoken *OtokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Otoken *OtokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Otoken.Contract.BalanceOf(&_Otoken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Otoken *OtokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Otoken.Contract.BalanceOf(&_Otoken.CallOpts, account)
}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_Otoken *OtokenCaller) CollateralAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "collateralAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_Otoken *OtokenSession) CollateralAsset() (common.Address, error) {
	return _Otoken.Contract.CollateralAsset(&_Otoken.CallOpts)
}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_Otoken *OtokenCallerSession) CollateralAsset() (common.Address, error) {
	return _Otoken.Contract.CollateralAsset(&_Otoken.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Otoken *OtokenCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Otoken *OtokenSession) Controller() (common.Address, error) {
	return _Otoken.Contract.Controller(&_Otoken.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Otoken *OtokenCallerSession) Controller() (common.Address, error) {
	return _Otoken.Contract.Controller(&_Otoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Otoken *OtokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Otoken *OtokenSession) Decimals() (uint8, error) {
	return _Otoken.Contract.Decimals(&_Otoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Otoken *OtokenCallerSession) Decimals() (uint8, error) {
	return _Otoken.Contract.Decimals(&_Otoken.CallOpts)
}

// ExpiryTimestamp is a free data retrieval call binding the contract method 0xade6e2aa.
//
// Solidity: function expiryTimestamp() view returns(uint256)
func (_Otoken *OtokenCaller) ExpiryTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "expiryTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpiryTimestamp is a free data retrieval call binding the contract method 0xade6e2aa.
//
// Solidity: function expiryTimestamp() view returns(uint256)
func (_Otoken *OtokenSession) ExpiryTimestamp() (*big.Int, error) {
	return _Otoken.Contract.ExpiryTimestamp(&_Otoken.CallOpts)
}

// ExpiryTimestamp is a free data retrieval call binding the contract method 0xade6e2aa.
//
// Solidity: function expiryTimestamp() view returns(uint256)
func (_Otoken *OtokenCallerSession) ExpiryTimestamp() (*big.Int, error) {
	return _Otoken.Contract.ExpiryTimestamp(&_Otoken.CallOpts)
}

// GetOtokenDetails is a free data retrieval call binding the contract method 0xaf0968fc.
//
// Solidity: function getOtokenDetails() view returns(address, address, address, uint256, uint256, bool)
func (_Otoken *OtokenCaller) GetOtokenDetails(opts *bind.CallOpts) (common.Address, common.Address, common.Address, *big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "getOtokenDetails")

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(common.Address), *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, err

}

// GetOtokenDetails is a free data retrieval call binding the contract method 0xaf0968fc.
//
// Solidity: function getOtokenDetails() view returns(address, address, address, uint256, uint256, bool)
func (_Otoken *OtokenSession) GetOtokenDetails() (common.Address, common.Address, common.Address, *big.Int, *big.Int, bool, error) {
	return _Otoken.Contract.GetOtokenDetails(&_Otoken.CallOpts)
}

// GetOtokenDetails is a free data retrieval call binding the contract method 0xaf0968fc.
//
// Solidity: function getOtokenDetails() view returns(address, address, address, uint256, uint256, bool)
func (_Otoken *OtokenCallerSession) GetOtokenDetails() (common.Address, common.Address, common.Address, *big.Int, *big.Int, bool, error) {
	return _Otoken.Contract.GetOtokenDetails(&_Otoken.CallOpts)
}

// IsPut is a free data retrieval call binding the contract method 0xf3c274a6.
//
// Solidity: function isPut() view returns(bool)
func (_Otoken *OtokenCaller) IsPut(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "isPut")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPut is a free data retrieval call binding the contract method 0xf3c274a6.
//
// Solidity: function isPut() view returns(bool)
func (_Otoken *OtokenSession) IsPut() (bool, error) {
	return _Otoken.Contract.IsPut(&_Otoken.CallOpts)
}

// IsPut is a free data retrieval call binding the contract method 0xf3c274a6.
//
// Solidity: function isPut() view returns(bool)
func (_Otoken *OtokenCallerSession) IsPut() (bool, error) {
	return _Otoken.Contract.IsPut(&_Otoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Otoken *OtokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Otoken *OtokenSession) Name() (string, error) {
	return _Otoken.Contract.Name(&_Otoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Otoken *OtokenCallerSession) Name() (string, error) {
	return _Otoken.Contract.Name(&_Otoken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Otoken *OtokenCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Otoken *OtokenSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Otoken.Contract.Nonces(&_Otoken.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Otoken *OtokenCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Otoken.Contract.Nonces(&_Otoken.CallOpts, owner)
}

// StrikeAsset is a free data retrieval call binding the contract method 0x17d69bc8.
//
// Solidity: function strikeAsset() view returns(address)
func (_Otoken *OtokenCaller) StrikeAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "strikeAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrikeAsset is a free data retrieval call binding the contract method 0x17d69bc8.
//
// Solidity: function strikeAsset() view returns(address)
func (_Otoken *OtokenSession) StrikeAsset() (common.Address, error) {
	return _Otoken.Contract.StrikeAsset(&_Otoken.CallOpts)
}

// StrikeAsset is a free data retrieval call binding the contract method 0x17d69bc8.
//
// Solidity: function strikeAsset() view returns(address)
func (_Otoken *OtokenCallerSession) StrikeAsset() (common.Address, error) {
	return _Otoken.Contract.StrikeAsset(&_Otoken.CallOpts)
}

// StrikePrice is a free data retrieval call binding the contract method 0xc52987cf.
//
// Solidity: function strikePrice() view returns(uint256)
func (_Otoken *OtokenCaller) StrikePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "strikePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StrikePrice is a free data retrieval call binding the contract method 0xc52987cf.
//
// Solidity: function strikePrice() view returns(uint256)
func (_Otoken *OtokenSession) StrikePrice() (*big.Int, error) {
	return _Otoken.Contract.StrikePrice(&_Otoken.CallOpts)
}

// StrikePrice is a free data retrieval call binding the contract method 0xc52987cf.
//
// Solidity: function strikePrice() view returns(uint256)
func (_Otoken *OtokenCallerSession) StrikePrice() (*big.Int, error) {
	return _Otoken.Contract.StrikePrice(&_Otoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Otoken *OtokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Otoken *OtokenSession) Symbol() (string, error) {
	return _Otoken.Contract.Symbol(&_Otoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Otoken *OtokenCallerSession) Symbol() (string, error) {
	return _Otoken.Contract.Symbol(&_Otoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Otoken *OtokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Otoken *OtokenSession) TotalSupply() (*big.Int, error) {
	return _Otoken.Contract.TotalSupply(&_Otoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Otoken *OtokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Otoken.Contract.TotalSupply(&_Otoken.CallOpts)
}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(address)
func (_Otoken *OtokenCaller) UnderlyingAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Otoken.contract.Call(opts, &out, "underlyingAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(address)
func (_Otoken *OtokenSession) UnderlyingAsset() (common.Address, error) {
	return _Otoken.Contract.UnderlyingAsset(&_Otoken.CallOpts)
}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(address)
func (_Otoken *OtokenCallerSession) UnderlyingAsset() (common.Address, error) {
	return _Otoken.Contract.UnderlyingAsset(&_Otoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Otoken *OtokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Otoken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Otoken *OtokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Otoken.Contract.Approve(&_Otoken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Otoken *OtokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Otoken.Contract.Approve(&_Otoken.TransactOpts, spender, amount)
}

// BurnOtoken is a paid mutator transaction binding the contract method 0x56d878f7.
//
// Solidity: function burnOtoken(address account, uint256 amount) returns()
func (_Otoken *OtokenTransactor) BurnOtoken(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Otoken.contract.Transact(opts, "burnOtoken", account, amount)
}

// BurnOtoken is a paid mutator transaction binding the contract method 0x56d878f7.
//
// Solidity: function burnOtoken(address account, uint256 amount) returns()
func (_Otoken *OtokenSession) BurnOtoken(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Otoken.Contract.BurnOtoken(&_Otoken.TransactOpts, account, amount)
}

// BurnOtoken is a paid mutator transaction binding the contract method 0x56d878f7.
//
// Solidity: function burnOtoken(address account, uint256 amount) returns()
func (_Otoken *OtokenTransactorSession) BurnOtoken(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Otoken.Contract.BurnOtoken(&_Otoken.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Otoken *OtokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Otoken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Otoken *OtokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Otoken.Contract.DecreaseAllowance(&_Otoken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Otoken *OtokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Otoken.Contract.DecreaseAllowance(&_Otoken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Otoken *OtokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Otoken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Otoken *OtokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Otoken.Contract.IncreaseAllowance(&_Otoken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Otoken *OtokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Otoken.Contract.IncreaseAllowance(&_Otoken.TransactOpts, spender, addedValue)
}

// Init is a paid mutator transaction binding the contract method 0xf630df34.
//
// Solidity: function init(address _addressBook, address _underlyingAsset, address _strikeAsset, address _collateralAsset, uint256 _strikePrice, uint256 _expiryTimestamp, bool _isPut) returns()
func (_Otoken *OtokenTransactor) Init(opts *bind.TransactOpts, _addressBook common.Address, _underlyingAsset common.Address, _strikeAsset common.Address, _collateralAsset common.Address, _strikePrice *big.Int, _expiryTimestamp *big.Int, _isPut bool) (*types.Transaction, error) {
	return _Otoken.contract.Transact(opts, "init", _addressBook, _underlyingAsset, _strikeAsset, _collateralAsset, _strikePrice, _expiryTimestamp, _isPut)
}

// Init is a paid mutator transaction binding the contract method 0xf630df34.
//
// Solidity: function init(address _addressBook, address _underlyingAsset, address _strikeAsset, address _collateralAsset, uint256 _strikePrice, uint256 _expiryTimestamp, bool _isPut) returns()
func (_Otoken *OtokenSession) Init(_addressBook common.Address, _underlyingAsset common.Address, _strikeAsset common.Address, _collateralAsset common.Address, _strikePrice *big.Int, _expiryTimestamp *big.Int, _isPut bool) (*types.Transaction, error) {
	return _Otoken.Contract.Init(&_Otoken.TransactOpts, _addressBook, _underlyingAsset, _strikeAsset, _collateralAsset, _strikePrice, _expiryTimestamp, _isPut)
}

// Init is a paid mutator transaction binding the contract method 0xf630df34.
//
// Solidity: function init(address _addressBook, address _underlyingAsset, address _strikeAsset, address _collateralAsset, uint256 _strikePrice, uint256 _expiryTimestamp, bool _isPut) returns()
func (_Otoken *OtokenTransactorSession) Init(_addressBook common.Address, _underlyingAsset common.Address, _strikeAsset common.Address, _collateralAsset common.Address, _strikePrice *big.Int, _expiryTimestamp *big.Int, _isPut bool) (*types.Transaction, error) {
	return _Otoken.Contract.Init(&_Otoken.TransactOpts, _addressBook, _underlyingAsset, _strikeAsset, _collateralAsset, _strikePrice, _expiryTimestamp, _isPut)
}

// MintOtoken is a paid mutator transaction binding the contract method 0x51b0a410.
//
// Solidity: function mintOtoken(address account, uint256 amount) returns()
func (_Otoken *OtokenTransactor) MintOtoken(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Otoken.contract.Transact(opts, "mintOtoken", account, amount)
}

// MintOtoken is a paid mutator transaction binding the contract method 0x51b0a410.
//
// Solidity: function mintOtoken(address account, uint256 amount) returns()
func (_Otoken *OtokenSession) MintOtoken(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Otoken.Contract.MintOtoken(&_Otoken.TransactOpts, account, amount)
}

// MintOtoken is a paid mutator transaction binding the contract method 0x51b0a410.
//
// Solidity: function mintOtoken(address account, uint256 amount) returns()
func (_Otoken *OtokenTransactorSession) MintOtoken(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Otoken.Contract.MintOtoken(&_Otoken.TransactOpts, account, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Otoken *OtokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Otoken.contract.Transact(opts, "permit", owner, spender, amount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Otoken *OtokenSession) Permit(owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Otoken.Contract.Permit(&_Otoken.TransactOpts, owner, spender, amount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Otoken *OtokenTransactorSession) Permit(owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Otoken.Contract.Permit(&_Otoken.TransactOpts, owner, spender, amount, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Otoken *OtokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Otoken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Otoken *OtokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Otoken.Contract.Transfer(&_Otoken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Otoken *OtokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Otoken.Contract.Transfer(&_Otoken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Otoken *OtokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Otoken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Otoken *OtokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Otoken.Contract.TransferFrom(&_Otoken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Otoken *OtokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Otoken.Contract.TransferFrom(&_Otoken.TransactOpts, sender, recipient, amount)
}

// OtokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Otoken contract.
type OtokenApprovalIterator struct {
	Event *OtokenApproval // Event containing the contract specifics and raw log

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
func (it *OtokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenApproval)
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
		it.Event = new(OtokenApproval)
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
func (it *OtokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenApproval represents a Approval event raised by the Otoken contract.
type OtokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Otoken *OtokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*OtokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Otoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OtokenApprovalIterator{contract: _Otoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Otoken *OtokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OtokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Otoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenApproval)
				if err := _Otoken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Otoken *OtokenFilterer) ParseApproval(log types.Log) (*OtokenApproval, error) {
	event := new(OtokenApproval)
	if err := _Otoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Otoken contract.
type OtokenTransferIterator struct {
	Event *OtokenTransfer // Event containing the contract specifics and raw log

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
func (it *OtokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenTransfer)
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
		it.Event = new(OtokenTransfer)
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
func (it *OtokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenTransfer represents a Transfer event raised by the Otoken contract.
type OtokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Otoken *OtokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OtokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Otoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OtokenTransferIterator{contract: _Otoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Otoken *OtokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OtokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Otoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenTransfer)
				if err := _Otoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Otoken *OtokenFilterer) ParseTransfer(log types.Log) (*OtokenTransfer, error) {
	event := new(OtokenTransfer)
	if err := _Otoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
