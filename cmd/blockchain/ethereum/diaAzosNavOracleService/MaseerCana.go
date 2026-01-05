// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
	_ = abi.ConvertType
)

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gem_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pip_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"act_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adm_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cop_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"flo_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ClaimableAfter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"}],\"name\":\"DustThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedsCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"InvalidChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recovered\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketClosed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPendingClaim\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"PermitDeadlineExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ContractRedemption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Issued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Settled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"smelter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Smelted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"act\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adm\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burncost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"canPass\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"capacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cop\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_out\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_out\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintcost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"navprice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"obligated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pip\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"redemptionAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"redemptionAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redemptionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"redemptionDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"redemptions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"date\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_out\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"smelt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"smelt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"terms\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unsettled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Main *MainCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Main *MainSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Main.Contract.DOMAINSEPARATOR(&_Main.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Main *MainCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Main.Contract.DOMAINSEPARATOR(&_Main.CallOpts)
}

// Act is a free data retrieval call binding the contract method 0xee42b10b.
//
// Solidity: function act() view returns(address)
func (_Main *MainCaller) Act(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "act")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Act is a free data retrieval call binding the contract method 0xee42b10b.
//
// Solidity: function act() view returns(address)
func (_Main *MainSession) Act() (common.Address, error) {
	return _Main.Contract.Act(&_Main.CallOpts)
}

// Act is a free data retrieval call binding the contract method 0xee42b10b.
//
// Solidity: function act() view returns(address)
func (_Main *MainCallerSession) Act() (common.Address, error) {
	return _Main.Contract.Act(&_Main.CallOpts)
}

// Adm is a free data retrieval call binding the contract method 0x04d7aef2.
//
// Solidity: function adm() view returns(address)
func (_Main *MainCaller) Adm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "adm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Adm is a free data retrieval call binding the contract method 0x04d7aef2.
//
// Solidity: function adm() view returns(address)
func (_Main *MainSession) Adm() (common.Address, error) {
	return _Main.Contract.Adm(&_Main.CallOpts)
}

// Adm is a free data retrieval call binding the contract method 0x04d7aef2.
//
// Solidity: function adm() view returns(address)
func (_Main *MainCallerSession) Adm() (common.Address, error) {
	return _Main.Contract.Adm(&_Main.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Main *MainCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Main *MainSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Main.Contract.Allowance(&_Main.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Main *MainCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Main.Contract.Allowance(&_Main.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Main *MainCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Main *MainSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.BalanceOf(&_Main.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Main *MainCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.BalanceOf(&_Main.CallOpts, arg0)
}

// Burnable is a free data retrieval call binding the contract method 0xa07c7ce4.
//
// Solidity: function burnable() view returns(bool)
func (_Main *MainCaller) Burnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "burnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Burnable is a free data retrieval call binding the contract method 0xa07c7ce4.
//
// Solidity: function burnable() view returns(bool)
func (_Main *MainSession) Burnable() (bool, error) {
	return _Main.Contract.Burnable(&_Main.CallOpts)
}

// Burnable is a free data retrieval call binding the contract method 0xa07c7ce4.
//
// Solidity: function burnable() view returns(bool)
func (_Main *MainCallerSession) Burnable() (bool, error) {
	return _Main.Contract.Burnable(&_Main.CallOpts)
}

// Burncost is a free data retrieval call binding the contract method 0x6c96474f.
//
// Solidity: function burncost() view returns(uint256)
func (_Main *MainCaller) Burncost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "burncost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Burncost is a free data retrieval call binding the contract method 0x6c96474f.
//
// Solidity: function burncost() view returns(uint256)
func (_Main *MainSession) Burncost() (*big.Int, error) {
	return _Main.Contract.Burncost(&_Main.CallOpts)
}

// Burncost is a free data retrieval call binding the contract method 0x6c96474f.
//
// Solidity: function burncost() view returns(uint256)
func (_Main *MainCallerSession) Burncost() (*big.Int, error) {
	return _Main.Contract.Burncost(&_Main.CallOpts)
}

// CanPass is a free data retrieval call binding the contract method 0x248c11f5.
//
// Solidity: function canPass(address usr) view returns(bool)
func (_Main *MainCaller) CanPass(opts *bind.CallOpts, usr common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "canPass", usr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPass is a free data retrieval call binding the contract method 0x248c11f5.
//
// Solidity: function canPass(address usr) view returns(bool)
func (_Main *MainSession) CanPass(usr common.Address) (bool, error) {
	return _Main.Contract.CanPass(&_Main.CallOpts, usr)
}

// CanPass is a free data retrieval call binding the contract method 0x248c11f5.
//
// Solidity: function canPass(address usr) view returns(bool)
func (_Main *MainCallerSession) CanPass(usr common.Address) (bool, error) {
	return _Main.Contract.CanPass(&_Main.CallOpts, usr)
}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() view returns(uint256)
func (_Main *MainCaller) Capacity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "capacity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() view returns(uint256)
func (_Main *MainSession) Capacity() (*big.Int, error) {
	return _Main.Contract.Capacity(&_Main.CallOpts)
}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() view returns(uint256)
func (_Main *MainCallerSession) Capacity() (*big.Int, error) {
	return _Main.Contract.Capacity(&_Main.CallOpts)
}

// Cooldown is a free data retrieval call binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() view returns(uint256)
func (_Main *MainCaller) Cooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "cooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cooldown is a free data retrieval call binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() view returns(uint256)
func (_Main *MainSession) Cooldown() (*big.Int, error) {
	return _Main.Contract.Cooldown(&_Main.CallOpts)
}

// Cooldown is a free data retrieval call binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() view returns(uint256)
func (_Main *MainCallerSession) Cooldown() (*big.Int, error) {
	return _Main.Contract.Cooldown(&_Main.CallOpts)
}

// Cop is a free data retrieval call binding the contract method 0x5f7efbea.
//
// Solidity: function cop() view returns(address)
func (_Main *MainCaller) Cop(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "cop")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cop is a free data retrieval call binding the contract method 0x5f7efbea.
//
// Solidity: function cop() view returns(address)
func (_Main *MainSession) Cop() (common.Address, error) {
	return _Main.Contract.Cop(&_Main.CallOpts)
}

// Cop is a free data retrieval call binding the contract method 0x5f7efbea.
//
// Solidity: function cop() view returns(address)
func (_Main *MainCallerSession) Cop() (common.Address, error) {
	return _Main.Contract.Cop(&_Main.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Main *MainCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Main *MainSession) Decimals() (uint8, error) {
	return _Main.Contract.Decimals(&_Main.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Main *MainCallerSession) Decimals() (uint8, error) {
	return _Main.Contract.Decimals(&_Main.CallOpts)
}

// Flo is a free data retrieval call binding the contract method 0x0a77bd49.
//
// Solidity: function flo() view returns(address)
func (_Main *MainCaller) Flo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "flo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Flo is a free data retrieval call binding the contract method 0x0a77bd49.
//
// Solidity: function flo() view returns(address)
func (_Main *MainSession) Flo() (common.Address, error) {
	return _Main.Contract.Flo(&_Main.CallOpts)
}

// Flo is a free data retrieval call binding the contract method 0x0a77bd49.
//
// Solidity: function flo() view returns(address)
func (_Main *MainCallerSession) Flo() (common.Address, error) {
	return _Main.Contract.Flo(&_Main.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_Main *MainCaller) Gem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "gem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_Main *MainSession) Gem() (common.Address, error) {
	return _Main.Contract.Gem(&_Main.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_Main *MainCallerSession) Gem() (common.Address, error) {
	return _Main.Contract.Gem(&_Main.CallOpts)
}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_Main *MainCaller) Mintable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "mintable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_Main *MainSession) Mintable() (bool, error) {
	return _Main.Contract.Mintable(&_Main.CallOpts)
}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_Main *MainCallerSession) Mintable() (bool, error) {
	return _Main.Contract.Mintable(&_Main.CallOpts)
}

// Mintcost is a free data retrieval call binding the contract method 0x94c3cbf4.
//
// Solidity: function mintcost() view returns(uint256)
func (_Main *MainCaller) Mintcost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "mintcost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Mintcost is a free data retrieval call binding the contract method 0x94c3cbf4.
//
// Solidity: function mintcost() view returns(uint256)
func (_Main *MainSession) Mintcost() (*big.Int, error) {
	return _Main.Contract.Mintcost(&_Main.CallOpts)
}

// Mintcost is a free data retrieval call binding the contract method 0x94c3cbf4.
//
// Solidity: function mintcost() view returns(uint256)
func (_Main *MainCallerSession) Mintcost() (*big.Int, error) {
	return _Main.Contract.Mintcost(&_Main.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Main *MainCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Main *MainSession) Name() (string, error) {
	return _Main.Contract.Name(&_Main.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Main *MainCallerSession) Name() (string, error) {
	return _Main.Contract.Name(&_Main.CallOpts)
}

// Navprice is a free data retrieval call binding the contract method 0x973efb5c.
//
// Solidity: function navprice() view returns(uint256)
func (_Main *MainCaller) Navprice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "navprice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Navprice is a free data retrieval call binding the contract method 0x973efb5c.
//
// Solidity: function navprice() view returns(uint256)
func (_Main *MainSession) Navprice() (*big.Int, error) {
	return _Main.Contract.Navprice(&_Main.CallOpts)
}

// Navprice is a free data retrieval call binding the contract method 0x973efb5c.
//
// Solidity: function navprice() view returns(uint256)
func (_Main *MainCallerSession) Navprice() (*big.Int, error) {
	return _Main.Contract.Navprice(&_Main.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Main *MainCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Main *MainSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.Nonces(&_Main.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Main *MainCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.Nonces(&_Main.CallOpts, arg0)
}

// Obligated is a free data retrieval call binding the contract method 0x4d773aa6.
//
// Solidity: function obligated() view returns(uint256)
func (_Main *MainCaller) Obligated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "obligated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Obligated is a free data retrieval call binding the contract method 0x4d773aa6.
//
// Solidity: function obligated() view returns(uint256)
func (_Main *MainSession) Obligated() (*big.Int, error) {
	return _Main.Contract.Obligated(&_Main.CallOpts)
}

// Obligated is a free data retrieval call binding the contract method 0x4d773aa6.
//
// Solidity: function obligated() view returns(uint256)
func (_Main *MainCallerSession) Obligated() (*big.Int, error) {
	return _Main.Contract.Obligated(&_Main.CallOpts)
}

// Pip is a free data retrieval call binding the contract method 0xd741e2f9.
//
// Solidity: function pip() view returns(address)
func (_Main *MainCaller) Pip(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "pip")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pip is a free data retrieval call binding the contract method 0xd741e2f9.
//
// Solidity: function pip() view returns(address)
func (_Main *MainSession) Pip() (common.Address, error) {
	return _Main.Contract.Pip(&_Main.CallOpts)
}

// Pip is a free data retrieval call binding the contract method 0xd741e2f9.
//
// Solidity: function pip() view returns(address)
func (_Main *MainCallerSession) Pip() (common.Address, error) {
	return _Main.Contract.Pip(&_Main.CallOpts)
}

// RedemptionAddr is a free data retrieval call binding the contract method 0x51f5d98d.
//
// Solidity: function redemptionAddr(uint256 id) view returns(address)
func (_Main *MainCaller) RedemptionAddr(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "redemptionAddr", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RedemptionAddr is a free data retrieval call binding the contract method 0x51f5d98d.
//
// Solidity: function redemptionAddr(uint256 id) view returns(address)
func (_Main *MainSession) RedemptionAddr(id *big.Int) (common.Address, error) {
	return _Main.Contract.RedemptionAddr(&_Main.CallOpts, id)
}

// RedemptionAddr is a free data retrieval call binding the contract method 0x51f5d98d.
//
// Solidity: function redemptionAddr(uint256 id) view returns(address)
func (_Main *MainCallerSession) RedemptionAddr(id *big.Int) (common.Address, error) {
	return _Main.Contract.RedemptionAddr(&_Main.CallOpts, id)
}

// RedemptionAmount is a free data retrieval call binding the contract method 0x06ff7a41.
//
// Solidity: function redemptionAmount(uint256 id) view returns(uint256)
func (_Main *MainCaller) RedemptionAmount(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "redemptionAmount", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedemptionAmount is a free data retrieval call binding the contract method 0x06ff7a41.
//
// Solidity: function redemptionAmount(uint256 id) view returns(uint256)
func (_Main *MainSession) RedemptionAmount(id *big.Int) (*big.Int, error) {
	return _Main.Contract.RedemptionAmount(&_Main.CallOpts, id)
}

// RedemptionAmount is a free data retrieval call binding the contract method 0x06ff7a41.
//
// Solidity: function redemptionAmount(uint256 id) view returns(uint256)
func (_Main *MainCallerSession) RedemptionAmount(id *big.Int) (*big.Int, error) {
	return _Main.Contract.RedemptionAmount(&_Main.CallOpts, id)
}

// RedemptionCount is a free data retrieval call binding the contract method 0x26ea3799.
//
// Solidity: function redemptionCount() view returns(uint256)
func (_Main *MainCaller) RedemptionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "redemptionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedemptionCount is a free data retrieval call binding the contract method 0x26ea3799.
//
// Solidity: function redemptionCount() view returns(uint256)
func (_Main *MainSession) RedemptionCount() (*big.Int, error) {
	return _Main.Contract.RedemptionCount(&_Main.CallOpts)
}

// RedemptionCount is a free data retrieval call binding the contract method 0x26ea3799.
//
// Solidity: function redemptionCount() view returns(uint256)
func (_Main *MainCallerSession) RedemptionCount() (*big.Int, error) {
	return _Main.Contract.RedemptionCount(&_Main.CallOpts)
}

// RedemptionDate is a free data retrieval call binding the contract method 0x42c4a1aa.
//
// Solidity: function redemptionDate(uint256 id) view returns(uint256)
func (_Main *MainCaller) RedemptionDate(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "redemptionDate", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedemptionDate is a free data retrieval call binding the contract method 0x42c4a1aa.
//
// Solidity: function redemptionDate(uint256 id) view returns(uint256)
func (_Main *MainSession) RedemptionDate(id *big.Int) (*big.Int, error) {
	return _Main.Contract.RedemptionDate(&_Main.CallOpts, id)
}

// RedemptionDate is a free data retrieval call binding the contract method 0x42c4a1aa.
//
// Solidity: function redemptionDate(uint256 id) view returns(uint256)
func (_Main *MainCallerSession) RedemptionDate(id *big.Int) (*big.Int, error) {
	return _Main.Contract.RedemptionDate(&_Main.CallOpts, id)
}

// Redemptions is a free data retrieval call binding the contract method 0xbeb65893.
//
// Solidity: function redemptions(uint256 ) view returns(uint256 amount, address redeemer, uint96 date)
func (_Main *MainCaller) Redemptions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount   *big.Int
	Redeemer common.Address
	Date     *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "redemptions", arg0)

	outstruct := new(struct {
		Amount   *big.Int
		Redeemer common.Address
		Date     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Redeemer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Date = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Redemptions is a free data retrieval call binding the contract method 0xbeb65893.
//
// Solidity: function redemptions(uint256 ) view returns(uint256 amount, address redeemer, uint96 date)
func (_Main *MainSession) Redemptions(arg0 *big.Int) (struct {
	Amount   *big.Int
	Redeemer common.Address
	Date     *big.Int
}, error) {
	return _Main.Contract.Redemptions(&_Main.CallOpts, arg0)
}

// Redemptions is a free data retrieval call binding the contract method 0xbeb65893.
//
// Solidity: function redemptions(uint256 ) view returns(uint256 amount, address redeemer, uint96 date)
func (_Main *MainCallerSession) Redemptions(arg0 *big.Int) (struct {
	Amount   *big.Int
	Redeemer common.Address
	Date     *big.Int
}, error) {
	return _Main.Contract.Redemptions(&_Main.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Main *MainCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Main *MainSession) Symbol() (string, error) {
	return _Main.Contract.Symbol(&_Main.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Main *MainCallerSession) Symbol() (string, error) {
	return _Main.Contract.Symbol(&_Main.CallOpts)
}

// Terms is a free data retrieval call binding the contract method 0xd5025625.
//
// Solidity: function terms() view returns(string)
func (_Main *MainCaller) Terms(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "terms")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Terms is a free data retrieval call binding the contract method 0xd5025625.
//
// Solidity: function terms() view returns(string)
func (_Main *MainSession) Terms() (string, error) {
	return _Main.Contract.Terms(&_Main.CallOpts)
}

// Terms is a free data retrieval call binding the contract method 0xd5025625.
//
// Solidity: function terms() view returns(string)
func (_Main *MainCallerSession) Terms() (string, error) {
	return _Main.Contract.Terms(&_Main.CallOpts)
}

// TotalPending is a free data retrieval call binding the contract method 0x3f90916a.
//
// Solidity: function totalPending() view returns(uint256)
func (_Main *MainCaller) TotalPending(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "totalPending")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPending is a free data retrieval call binding the contract method 0x3f90916a.
//
// Solidity: function totalPending() view returns(uint256)
func (_Main *MainSession) TotalPending() (*big.Int, error) {
	return _Main.Contract.TotalPending(&_Main.CallOpts)
}

// TotalPending is a free data retrieval call binding the contract method 0x3f90916a.
//
// Solidity: function totalPending() view returns(uint256)
func (_Main *MainCallerSession) TotalPending() (*big.Int, error) {
	return _Main.Contract.TotalPending(&_Main.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Main *MainCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Main *MainSession) TotalSupply() (*big.Int, error) {
	return _Main.Contract.TotalSupply(&_Main.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Main *MainCallerSession) TotalSupply() (*big.Int, error) {
	return _Main.Contract.TotalSupply(&_Main.CallOpts)
}

// Unsettled is a free data retrieval call binding the contract method 0xede75094.
//
// Solidity: function unsettled() view returns(uint256)
func (_Main *MainCaller) Unsettled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "unsettled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Unsettled is a free data retrieval call binding the contract method 0xede75094.
//
// Solidity: function unsettled() view returns(uint256)
func (_Main *MainSession) Unsettled() (*big.Int, error) {
	return _Main.Contract.Unsettled(&_Main.CallOpts)
}

// Unsettled is a free data retrieval call binding the contract method 0xede75094.
//
// Solidity: function unsettled() view returns(uint256)
func (_Main *MainCallerSession) Unsettled() (*big.Int, error) {
	return _Main.Contract.Unsettled(&_Main.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address usr, uint256 wad) returns(bool)
func (_Main *MainTransactor) Approve(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "approve", usr, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address usr, uint256 wad) returns(bool)
func (_Main *MainSession) Approve(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Approve(&_Main.TransactOpts, usr, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address usr, uint256 wad) returns(bool)
func (_Main *MainTransactorSession) Approve(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Approve(&_Main.TransactOpts, usr, wad)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address usr) returns(bool)
func (_Main *MainTransactor) Approve0(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "approve0", usr)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address usr) returns(bool)
func (_Main *MainSession) Approve0(usr common.Address) (*types.Transaction, error) {
	return _Main.Contract.Approve0(&_Main.TransactOpts, usr)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address usr) returns(bool)
func (_Main *MainTransactorSession) Approve0(usr common.Address) (*types.Transaction, error) {
	return _Main.Contract.Approve0(&_Main.TransactOpts, usr)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 id) returns(uint256 _out)
func (_Main *MainTransactor) Exit(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "exit", id)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 id) returns(uint256 _out)
func (_Main *MainSession) Exit(id *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Exit(&_Main.TransactOpts, id)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 id) returns(uint256 _out)
func (_Main *MainTransactorSession) Exit(id *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Exit(&_Main.TransactOpts, id)
}

// Issue is a paid mutator transaction binding the contract method 0xcc872b66.
//
// Solidity: function issue(uint256 amt) returns()
func (_Main *MainTransactor) Issue(opts *bind.TransactOpts, amt *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "issue", amt)
}

// Issue is a paid mutator transaction binding the contract method 0xcc872b66.
//
// Solidity: function issue(uint256 amt) returns()
func (_Main *MainSession) Issue(amt *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Issue(&_Main.TransactOpts, amt)
}

// Issue is a paid mutator transaction binding the contract method 0xcc872b66.
//
// Solidity: function issue(uint256 amt) returns()
func (_Main *MainTransactorSession) Issue(amt *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Issue(&_Main.TransactOpts, amt)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amt) returns(uint256 _out)
func (_Main *MainTransactor) Mint(opts *bind.TransactOpts, amt *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "mint", amt)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amt) returns(uint256 _out)
func (_Main *MainSession) Mint(amt *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Mint(&_Main.TransactOpts, amt)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amt) returns(uint256 _out)
func (_Main *MainTransactorSession) Mint(amt *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Mint(&_Main.TransactOpts, amt)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Main *MainTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Main *MainSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Main.Contract.Permit(&_Main.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Main *MainTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Main.Contract.Permit(&_Main.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amt) returns(uint256 _id)
func (_Main *MainTransactor) Redeem(opts *bind.TransactOpts, amt *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "redeem", amt)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amt) returns(uint256 _id)
func (_Main *MainSession) Redeem(amt *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Redeem(&_Main.TransactOpts, amt)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amt) returns(uint256 _id)
func (_Main *MainTransactorSession) Redeem(amt *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Redeem(&_Main.TransactOpts, amt)
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() returns(uint256 _out)
func (_Main *MainTransactor) Settle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "settle")
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() returns(uint256 _out)
func (_Main *MainSession) Settle() (*types.Transaction, error) {
	return _Main.Contract.Settle(&_Main.TransactOpts)
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() returns(uint256 _out)
func (_Main *MainTransactorSession) Settle() (*types.Transaction, error) {
	return _Main.Contract.Settle(&_Main.TransactOpts)
}

// Smelt is a paid mutator transaction binding the contract method 0x219df1cc.
//
// Solidity: function smelt(uint256 amt) returns()
func (_Main *MainTransactor) Smelt(opts *bind.TransactOpts, amt *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "smelt", amt)
}

// Smelt is a paid mutator transaction binding the contract method 0x219df1cc.
//
// Solidity: function smelt(uint256 amt) returns()
func (_Main *MainSession) Smelt(amt *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Smelt(&_Main.TransactOpts, amt)
}

// Smelt is a paid mutator transaction binding the contract method 0x219df1cc.
//
// Solidity: function smelt(uint256 amt) returns()
func (_Main *MainTransactorSession) Smelt(amt *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Smelt(&_Main.TransactOpts, amt)
}

// Smelt0 is a paid mutator transaction binding the contract method 0xad1a2396.
//
// Solidity: function smelt(address usr, uint256 amt) returns()
func (_Main *MainTransactor) Smelt0(opts *bind.TransactOpts, usr common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "smelt0", usr, amt)
}

// Smelt0 is a paid mutator transaction binding the contract method 0xad1a2396.
//
// Solidity: function smelt(address usr, uint256 amt) returns()
func (_Main *MainSession) Smelt0(usr common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Smelt0(&_Main.TransactOpts, usr, amt)
}

// Smelt0 is a paid mutator transaction binding the contract method 0xad1a2396.
//
// Solidity: function smelt(address usr, uint256 amt) returns()
func (_Main *MainTransactorSession) Smelt0(usr common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Smelt0(&_Main.TransactOpts, usr, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Main *MainTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Main *MainSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Transfer(&_Main.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Main *MainTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Transfer(&_Main.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Main *MainTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Main *MainSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TransferFrom(&_Main.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Main *MainTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TransferFrom(&_Main.TransactOpts, src, dst, wad)
}

// MainApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Main contract.
type MainApprovalIterator struct {
	Event *MainApproval // Event containing the contract specifics and raw log

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
func (it *MainApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainApproval)
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
		it.Event = new(MainApproval)
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
func (it *MainApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainApproval represents a Approval event raised by the Main contract.
type MainApproval struct {
	Src common.Address
	Usr common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed usr, uint256 wad)
func (_Main *MainFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, usr []common.Address) (*MainApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Approval", srcRule, usrRule)
	if err != nil {
		return nil, err
	}
	return &MainApprovalIterator{contract: _Main.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed usr, uint256 wad)
func (_Main *MainFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MainApproval, src []common.Address, usr []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Approval", srcRule, usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainApproval)
				if err := _Main.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed src, address indexed usr, uint256 wad)
func (_Main *MainFilterer) ParseApproval(log types.Log) (*MainApproval, error) {
	event := new(MainApproval)
	if err := _Main.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainClaimProcessedIterator is returned from FilterClaimProcessed and is used to iterate over the raw logs and unpacked data for ClaimProcessed events raised by the Main contract.
type MainClaimProcessedIterator struct {
	Event *MainClaimProcessed // Event containing the contract specifics and raw log

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
func (it *MainClaimProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainClaimProcessed)
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
		it.Event = new(MainClaimProcessed)
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
func (it *MainClaimProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainClaimProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainClaimProcessed represents a ClaimProcessed event raised by the Main contract.
type MainClaimProcessed struct {
	Id      *big.Int
	Claimer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimProcessed is a free log retrieval operation binding the contract event 0x786e090d783caced3baddcb29bd555ffb93faf484529ab87feab4ec4c48a81e1.
//
// Solidity: event ClaimProcessed(uint256 indexed id, address indexed claimer, uint256 indexed amount)
func (_Main *MainFilterer) FilterClaimProcessed(opts *bind.FilterOpts, id []*big.Int, claimer []common.Address, amount []*big.Int) (*MainClaimProcessedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "ClaimProcessed", idRule, claimerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &MainClaimProcessedIterator{contract: _Main.contract, event: "ClaimProcessed", logs: logs, sub: sub}, nil
}

// WatchClaimProcessed is a free log subscription operation binding the contract event 0x786e090d783caced3baddcb29bd555ffb93faf484529ab87feab4ec4c48a81e1.
//
// Solidity: event ClaimProcessed(uint256 indexed id, address indexed claimer, uint256 indexed amount)
func (_Main *MainFilterer) WatchClaimProcessed(opts *bind.WatchOpts, sink chan<- *MainClaimProcessed, id []*big.Int, claimer []common.Address, amount []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "ClaimProcessed", idRule, claimerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainClaimProcessed)
				if err := _Main.contract.UnpackLog(event, "ClaimProcessed", log); err != nil {
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

// ParseClaimProcessed is a log parse operation binding the contract event 0x786e090d783caced3baddcb29bd555ffb93faf484529ab87feab4ec4c48a81e1.
//
// Solidity: event ClaimProcessed(uint256 indexed id, address indexed claimer, uint256 indexed amount)
func (_Main *MainFilterer) ParseClaimProcessed(log types.Log) (*MainClaimProcessed, error) {
	event := new(MainClaimProcessed)
	if err := _Main.contract.UnpackLog(event, "ClaimProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainContractCreatedIterator is returned from FilterContractCreated and is used to iterate over the raw logs and unpacked data for ContractCreated events raised by the Main contract.
type MainContractCreatedIterator struct {
	Event *MainContractCreated // Event containing the contract specifics and raw log

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
func (it *MainContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainContractCreated)
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
		it.Event = new(MainContractCreated)
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
func (it *MainContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainContractCreated represents a ContractCreated event raised by the Main contract.
type MainContractCreated struct {
	Creator common.Address
	Price   *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractCreated is a free log retrieval operation binding the contract event 0x58781ca443f8e3ecffad28c83044b16c6e549d8e9c677aebb3b8ee176f99f53f.
//
// Solidity: event ContractCreated(address indexed creator, uint256 indexed price, uint256 indexed amount)
func (_Main *MainFilterer) FilterContractCreated(opts *bind.FilterOpts, creator []common.Address, price []*big.Int, amount []*big.Int) (*MainContractCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "ContractCreated", creatorRule, priceRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &MainContractCreatedIterator{contract: _Main.contract, event: "ContractCreated", logs: logs, sub: sub}, nil
}

// WatchContractCreated is a free log subscription operation binding the contract event 0x58781ca443f8e3ecffad28c83044b16c6e549d8e9c677aebb3b8ee176f99f53f.
//
// Solidity: event ContractCreated(address indexed creator, uint256 indexed price, uint256 indexed amount)
func (_Main *MainFilterer) WatchContractCreated(opts *bind.WatchOpts, sink chan<- *MainContractCreated, creator []common.Address, price []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "ContractCreated", creatorRule, priceRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainContractCreated)
				if err := _Main.contract.UnpackLog(event, "ContractCreated", log); err != nil {
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

// ParseContractCreated is a log parse operation binding the contract event 0x58781ca443f8e3ecffad28c83044b16c6e549d8e9c677aebb3b8ee176f99f53f.
//
// Solidity: event ContractCreated(address indexed creator, uint256 indexed price, uint256 indexed amount)
func (_Main *MainFilterer) ParseContractCreated(log types.Log) (*MainContractCreated, error) {
	event := new(MainContractCreated)
	if err := _Main.contract.UnpackLog(event, "ContractCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainContractRedemptionIterator is returned from FilterContractRedemption and is used to iterate over the raw logs and unpacked data for ContractRedemption events raised by the Main contract.
type MainContractRedemptionIterator struct {
	Event *MainContractRedemption // Event containing the contract specifics and raw log

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
func (it *MainContractRedemptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainContractRedemption)
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
		it.Event = new(MainContractRedemption)
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
func (it *MainContractRedemptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainContractRedemptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainContractRedemption represents a ContractRedemption event raised by the Main contract.
type MainContractRedemption struct {
	Id       *big.Int
	Date     *big.Int
	Redeemer common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractRedemption is a free log retrieval operation binding the contract event 0x5a084092180e1b482afdddc2a20027f4d3619225151fb1a9934c119cd7307f32.
//
// Solidity: event ContractRedemption(uint256 indexed id, uint256 indexed date, address indexed redeemer, uint256 amount)
func (_Main *MainFilterer) FilterContractRedemption(opts *bind.FilterOpts, id []*big.Int, date []*big.Int, redeemer []common.Address) (*MainContractRedemptionIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var dateRule []interface{}
	for _, dateItem := range date {
		dateRule = append(dateRule, dateItem)
	}
	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "ContractRedemption", idRule, dateRule, redeemerRule)
	if err != nil {
		return nil, err
	}
	return &MainContractRedemptionIterator{contract: _Main.contract, event: "ContractRedemption", logs: logs, sub: sub}, nil
}

// WatchContractRedemption is a free log subscription operation binding the contract event 0x5a084092180e1b482afdddc2a20027f4d3619225151fb1a9934c119cd7307f32.
//
// Solidity: event ContractRedemption(uint256 indexed id, uint256 indexed date, address indexed redeemer, uint256 amount)
func (_Main *MainFilterer) WatchContractRedemption(opts *bind.WatchOpts, sink chan<- *MainContractRedemption, id []*big.Int, date []*big.Int, redeemer []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var dateRule []interface{}
	for _, dateItem := range date {
		dateRule = append(dateRule, dateItem)
	}
	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "ContractRedemption", idRule, dateRule, redeemerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainContractRedemption)
				if err := _Main.contract.UnpackLog(event, "ContractRedemption", log); err != nil {
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

// ParseContractRedemption is a log parse operation binding the contract event 0x5a084092180e1b482afdddc2a20027f4d3619225151fb1a9934c119cd7307f32.
//
// Solidity: event ContractRedemption(uint256 indexed id, uint256 indexed date, address indexed redeemer, uint256 amount)
func (_Main *MainFilterer) ParseContractRedemption(log types.Log) (*MainContractRedemption, error) {
	event := new(MainContractRedemption)
	if err := _Main.contract.UnpackLog(event, "ContractRedemption", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainIssuedIterator is returned from FilterIssued and is used to iterate over the raw logs and unpacked data for Issued events raised by the Main contract.
type MainIssuedIterator struct {
	Event *MainIssued // Event containing the contract specifics and raw log

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
func (it *MainIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainIssued)
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
		it.Event = new(MainIssued)
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
func (it *MainIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainIssued represents a Issued event raised by the Main contract.
type MainIssued struct {
	Issuer common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssued is a free log retrieval operation binding the contract event 0xa59f12e354e8cd10bb74c559844c2dd69a5458e31fe56c7594c62ca57480509a.
//
// Solidity: event Issued(address indexed issuer, uint256 indexed amount)
func (_Main *MainFilterer) FilterIssued(opts *bind.FilterOpts, issuer []common.Address, amount []*big.Int) (*MainIssuedIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Issued", issuerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &MainIssuedIterator{contract: _Main.contract, event: "Issued", logs: logs, sub: sub}, nil
}

// WatchIssued is a free log subscription operation binding the contract event 0xa59f12e354e8cd10bb74c559844c2dd69a5458e31fe56c7594c62ca57480509a.
//
// Solidity: event Issued(address indexed issuer, uint256 indexed amount)
func (_Main *MainFilterer) WatchIssued(opts *bind.WatchOpts, sink chan<- *MainIssued, issuer []common.Address, amount []*big.Int) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Issued", issuerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainIssued)
				if err := _Main.contract.UnpackLog(event, "Issued", log); err != nil {
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

// ParseIssued is a log parse operation binding the contract event 0xa59f12e354e8cd10bb74c559844c2dd69a5458e31fe56c7594c62ca57480509a.
//
// Solidity: event Issued(address indexed issuer, uint256 indexed amount)
func (_Main *MainFilterer) ParseIssued(log types.Log) (*MainIssued, error) {
	event := new(MainIssued)
	if err := _Main.contract.UnpackLog(event, "Issued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainSettledIterator is returned from FilterSettled and is used to iterate over the raw logs and unpacked data for Settled events raised by the Main contract.
type MainSettledIterator struct {
	Event *MainSettled // Event containing the contract specifics and raw log

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
func (it *MainSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainSettled)
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
		it.Event = new(MainSettled)
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
func (it *MainSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainSettled represents a Settled event raised by the Main contract.
type MainSettled struct {
	Conduit common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSettled is a free log retrieval operation binding the contract event 0x7823e479a1a4ebe2418874847436f8a1680c5ee5b17f38bb59dbff28e1b45552.
//
// Solidity: event Settled(address indexed conduit, uint256 indexed amount)
func (_Main *MainFilterer) FilterSettled(opts *bind.FilterOpts, conduit []common.Address, amount []*big.Int) (*MainSettledIterator, error) {

	var conduitRule []interface{}
	for _, conduitItem := range conduit {
		conduitRule = append(conduitRule, conduitItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Settled", conduitRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &MainSettledIterator{contract: _Main.contract, event: "Settled", logs: logs, sub: sub}, nil
}

// WatchSettled is a free log subscription operation binding the contract event 0x7823e479a1a4ebe2418874847436f8a1680c5ee5b17f38bb59dbff28e1b45552.
//
// Solidity: event Settled(address indexed conduit, uint256 indexed amount)
func (_Main *MainFilterer) WatchSettled(opts *bind.WatchOpts, sink chan<- *MainSettled, conduit []common.Address, amount []*big.Int) (event.Subscription, error) {

	var conduitRule []interface{}
	for _, conduitItem := range conduit {
		conduitRule = append(conduitRule, conduitItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Settled", conduitRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainSettled)
				if err := _Main.contract.UnpackLog(event, "Settled", log); err != nil {
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

// ParseSettled is a log parse operation binding the contract event 0x7823e479a1a4ebe2418874847436f8a1680c5ee5b17f38bb59dbff28e1b45552.
//
// Solidity: event Settled(address indexed conduit, uint256 indexed amount)
func (_Main *MainFilterer) ParseSettled(log types.Log) (*MainSettled, error) {
	event := new(MainSettled)
	if err := _Main.contract.UnpackLog(event, "Settled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainSmeltedIterator is returned from FilterSmelted and is used to iterate over the raw logs and unpacked data for Smelted events raised by the Main contract.
type MainSmeltedIterator struct {
	Event *MainSmelted // Event containing the contract specifics and raw log

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
func (it *MainSmeltedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainSmelted)
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
		it.Event = new(MainSmelted)
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
func (it *MainSmeltedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainSmeltedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainSmelted represents a Smelted event raised by the Main contract.
type MainSmelted struct {
	Smelter common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSmelted is a free log retrieval operation binding the contract event 0x266ba2fe22d2b352fe1eb8c45b5d486ff3c0b83ef953959b6d34397d53fb088d.
//
// Solidity: event Smelted(address indexed smelter, uint256 indexed amount)
func (_Main *MainFilterer) FilterSmelted(opts *bind.FilterOpts, smelter []common.Address, amount []*big.Int) (*MainSmeltedIterator, error) {

	var smelterRule []interface{}
	for _, smelterItem := range smelter {
		smelterRule = append(smelterRule, smelterItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Smelted", smelterRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &MainSmeltedIterator{contract: _Main.contract, event: "Smelted", logs: logs, sub: sub}, nil
}

// WatchSmelted is a free log subscription operation binding the contract event 0x266ba2fe22d2b352fe1eb8c45b5d486ff3c0b83ef953959b6d34397d53fb088d.
//
// Solidity: event Smelted(address indexed smelter, uint256 indexed amount)
func (_Main *MainFilterer) WatchSmelted(opts *bind.WatchOpts, sink chan<- *MainSmelted, smelter []common.Address, amount []*big.Int) (event.Subscription, error) {

	var smelterRule []interface{}
	for _, smelterItem := range smelter {
		smelterRule = append(smelterRule, smelterItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Smelted", smelterRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainSmelted)
				if err := _Main.contract.UnpackLog(event, "Smelted", log); err != nil {
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

// ParseSmelted is a log parse operation binding the contract event 0x266ba2fe22d2b352fe1eb8c45b5d486ff3c0b83ef953959b6d34397d53fb088d.
//
// Solidity: event Smelted(address indexed smelter, uint256 indexed amount)
func (_Main *MainFilterer) ParseSmelted(log types.Log) (*MainSmelted, error) {
	event := new(MainSmelted)
	if err := _Main.contract.UnpackLog(event, "Smelted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Main contract.
type MainTransferIterator struct {
	Event *MainTransfer // Event containing the contract specifics and raw log

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
func (it *MainTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTransfer)
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
		it.Event = new(MainTransfer)
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
func (it *MainTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTransfer represents a Transfer event raised by the Main contract.
type MainTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Main *MainFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*MainTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &MainTransferIterator{contract: _Main.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Main *MainFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MainTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTransfer)
				if err := _Main.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Main *MainFilterer) ParseTransfer(log types.Log) (*MainTransfer, error) {
	event := new(MainTransfer)
	if err := _Main.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
