// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package anyerc20

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

// Anyerc20MetaData contains all meta data concerning the Anyerc20 contract.
var Anyerc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"LogAddAuth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"effectiveHeight\",\"type\":\"uint256\"}],\"name\":\"LogChangeMPCOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"effectiveTime\",\"type\":\"uint256\"}],\"name\":\"LogChangeVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txhash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogSwapin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bindaddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogSwapout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txhash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Swapin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bindaddr\",\"type\":\"address\"}],\"name\":\"Swapout\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"applyMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"applyVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"changeMPCOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"changeVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayMinter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositWithTransferPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllMinters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"initVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mpc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auth\",\"type\":\"address\"}],\"name\":\"revokeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auth\",\"type\":\"address\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"setVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setVaultOnly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"transferWithPermit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Anyerc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Anyerc20MetaData.ABI instead.
var Anyerc20ABI = Anyerc20MetaData.ABI

// Anyerc20 is an auto generated Go binding around an Ethereum contract.
type Anyerc20 struct {
	Anyerc20Caller     // Read-only binding to the contract
	Anyerc20Transactor // Write-only binding to the contract
	Anyerc20Filterer   // Log filterer for contract events
}

// Anyerc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Anyerc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Anyerc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Anyerc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Anyerc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Anyerc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Anyerc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Anyerc20Session struct {
	Contract     *Anyerc20         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Anyerc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Anyerc20CallerSession struct {
	Contract *Anyerc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Anyerc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Anyerc20TransactorSession struct {
	Contract     *Anyerc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Anyerc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Anyerc20Raw struct {
	Contract *Anyerc20 // Generic contract binding to access the raw methods on
}

// Anyerc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Anyerc20CallerRaw struct {
	Contract *Anyerc20Caller // Generic read-only contract binding to access the raw methods on
}

// Anyerc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Anyerc20TransactorRaw struct {
	Contract *Anyerc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAnyerc20 creates a new instance of Anyerc20, bound to a specific deployed contract.
func NewAnyerc20(address common.Address, backend bind.ContractBackend) (*Anyerc20, error) {
	contract, err := bindAnyerc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Anyerc20{Anyerc20Caller: Anyerc20Caller{contract: contract}, Anyerc20Transactor: Anyerc20Transactor{contract: contract}, Anyerc20Filterer: Anyerc20Filterer{contract: contract}}, nil
}

// NewAnyerc20Caller creates a new read-only instance of Anyerc20, bound to a specific deployed contract.
func NewAnyerc20Caller(address common.Address, caller bind.ContractCaller) (*Anyerc20Caller, error) {
	contract, err := bindAnyerc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Anyerc20Caller{contract: contract}, nil
}

// NewAnyerc20Transactor creates a new write-only instance of Anyerc20, bound to a specific deployed contract.
func NewAnyerc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Anyerc20Transactor, error) {
	contract, err := bindAnyerc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Anyerc20Transactor{contract: contract}, nil
}

// NewAnyerc20Filterer creates a new log filterer instance of Anyerc20, bound to a specific deployed contract.
func NewAnyerc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Anyerc20Filterer, error) {
	contract, err := bindAnyerc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Anyerc20Filterer{contract: contract}, nil
}

// bindAnyerc20 binds a generic wrapper to an already deployed contract.
func bindAnyerc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Anyerc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Anyerc20 *Anyerc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Anyerc20.Contract.Anyerc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Anyerc20 *Anyerc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anyerc20.Contract.Anyerc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Anyerc20 *Anyerc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Anyerc20.Contract.Anyerc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Anyerc20 *Anyerc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Anyerc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Anyerc20 *Anyerc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anyerc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Anyerc20 *Anyerc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Anyerc20.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Anyerc20 *Anyerc20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Anyerc20 *Anyerc20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _Anyerc20.Contract.DOMAINSEPARATOR(&_Anyerc20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Anyerc20 *Anyerc20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Anyerc20.Contract.DOMAINSEPARATOR(&_Anyerc20.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Anyerc20 *Anyerc20Caller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Anyerc20 *Anyerc20Session) PERMITTYPEHASH() ([32]byte, error) {
	return _Anyerc20.Contract.PERMITTYPEHASH(&_Anyerc20.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Anyerc20 *Anyerc20CallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Anyerc20.Contract.PERMITTYPEHASH(&_Anyerc20.CallOpts)
}

// TRANSFERTYPEHASH is a free data retrieval call binding the contract method 0x00bf26f4.
//
// Solidity: function TRANSFER_TYPEHASH() view returns(bytes32)
func (_Anyerc20 *Anyerc20Caller) TRANSFERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "TRANSFER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFERTYPEHASH is a free data retrieval call binding the contract method 0x00bf26f4.
//
// Solidity: function TRANSFER_TYPEHASH() view returns(bytes32)
func (_Anyerc20 *Anyerc20Session) TRANSFERTYPEHASH() ([32]byte, error) {
	return _Anyerc20.Contract.TRANSFERTYPEHASH(&_Anyerc20.CallOpts)
}

// TRANSFERTYPEHASH is a free data retrieval call binding the contract method 0x00bf26f4.
//
// Solidity: function TRANSFER_TYPEHASH() view returns(bytes32)
func (_Anyerc20 *Anyerc20CallerSession) TRANSFERTYPEHASH() ([32]byte, error) {
	return _Anyerc20.Contract.TRANSFERTYPEHASH(&_Anyerc20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Anyerc20 *Anyerc20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Anyerc20 *Anyerc20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Anyerc20.Contract.Allowance(&_Anyerc20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Anyerc20 *Anyerc20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Anyerc20.Contract.Allowance(&_Anyerc20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Anyerc20 *Anyerc20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Anyerc20 *Anyerc20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Anyerc20.Contract.BalanceOf(&_Anyerc20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Anyerc20 *Anyerc20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Anyerc20.Contract.BalanceOf(&_Anyerc20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Anyerc20 *Anyerc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Anyerc20 *Anyerc20Session) Decimals() (uint8, error) {
	return _Anyerc20.Contract.Decimals(&_Anyerc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Anyerc20 *Anyerc20CallerSession) Decimals() (uint8, error) {
	return _Anyerc20.Contract.Decimals(&_Anyerc20.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_Anyerc20 *Anyerc20Caller) Delay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "delay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_Anyerc20 *Anyerc20Session) Delay() (*big.Int, error) {
	return _Anyerc20.Contract.Delay(&_Anyerc20.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_Anyerc20 *Anyerc20CallerSession) Delay() (*big.Int, error) {
	return _Anyerc20.Contract.Delay(&_Anyerc20.CallOpts)
}

// DelayDelay is a free data retrieval call binding the contract method 0xa29dff72.
//
// Solidity: function delayDelay() view returns(uint256)
func (_Anyerc20 *Anyerc20Caller) DelayDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "delayDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayDelay is a free data retrieval call binding the contract method 0xa29dff72.
//
// Solidity: function delayDelay() view returns(uint256)
func (_Anyerc20 *Anyerc20Session) DelayDelay() (*big.Int, error) {
	return _Anyerc20.Contract.DelayDelay(&_Anyerc20.CallOpts)
}

// DelayDelay is a free data retrieval call binding the contract method 0xa29dff72.
//
// Solidity: function delayDelay() view returns(uint256)
func (_Anyerc20 *Anyerc20CallerSession) DelayDelay() (*big.Int, error) {
	return _Anyerc20.Contract.DelayDelay(&_Anyerc20.CallOpts)
}

// DelayMinter is a free data retrieval call binding the contract method 0xc3081240.
//
// Solidity: function delayMinter() view returns(uint256)
func (_Anyerc20 *Anyerc20Caller) DelayMinter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "delayMinter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayMinter is a free data retrieval call binding the contract method 0xc3081240.
//
// Solidity: function delayMinter() view returns(uint256)
func (_Anyerc20 *Anyerc20Session) DelayMinter() (*big.Int, error) {
	return _Anyerc20.Contract.DelayMinter(&_Anyerc20.CallOpts)
}

// DelayMinter is a free data retrieval call binding the contract method 0xc3081240.
//
// Solidity: function delayMinter() view returns(uint256)
func (_Anyerc20 *Anyerc20CallerSession) DelayMinter() (*big.Int, error) {
	return _Anyerc20.Contract.DelayMinter(&_Anyerc20.CallOpts)
}

// DelayVault is a free data retrieval call binding the contract method 0x87689e28.
//
// Solidity: function delayVault() view returns(uint256)
func (_Anyerc20 *Anyerc20Caller) DelayVault(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "delayVault")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayVault is a free data retrieval call binding the contract method 0x87689e28.
//
// Solidity: function delayVault() view returns(uint256)
func (_Anyerc20 *Anyerc20Session) DelayVault() (*big.Int, error) {
	return _Anyerc20.Contract.DelayVault(&_Anyerc20.CallOpts)
}

// DelayVault is a free data retrieval call binding the contract method 0x87689e28.
//
// Solidity: function delayVault() view returns(uint256)
func (_Anyerc20 *Anyerc20CallerSession) DelayVault() (*big.Int, error) {
	return _Anyerc20.Contract.DelayVault(&_Anyerc20.CallOpts)
}

// GetAllMinters is a free data retrieval call binding the contract method 0xa045442c.
//
// Solidity: function getAllMinters() view returns(address[])
func (_Anyerc20 *Anyerc20Caller) GetAllMinters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "getAllMinters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllMinters is a free data retrieval call binding the contract method 0xa045442c.
//
// Solidity: function getAllMinters() view returns(address[])
func (_Anyerc20 *Anyerc20Session) GetAllMinters() ([]common.Address, error) {
	return _Anyerc20.Contract.GetAllMinters(&_Anyerc20.CallOpts)
}

// GetAllMinters is a free data retrieval call binding the contract method 0xa045442c.
//
// Solidity: function getAllMinters() view returns(address[])
func (_Anyerc20 *Anyerc20CallerSession) GetAllMinters() ([]common.Address, error) {
	return _Anyerc20.Contract.GetAllMinters(&_Anyerc20.CallOpts)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_Anyerc20 *Anyerc20Caller) IsMinter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "isMinter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_Anyerc20 *Anyerc20Session) IsMinter(arg0 common.Address) (bool, error) {
	return _Anyerc20.Contract.IsMinter(&_Anyerc20.CallOpts, arg0)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_Anyerc20 *Anyerc20CallerSession) IsMinter(arg0 common.Address) (bool, error) {
	return _Anyerc20.Contract.IsMinter(&_Anyerc20.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0x8623ec7b.
//
// Solidity: function minters(uint256 ) view returns(address)
func (_Anyerc20 *Anyerc20Caller) Minters(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "minters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minters is a free data retrieval call binding the contract method 0x8623ec7b.
//
// Solidity: function minters(uint256 ) view returns(address)
func (_Anyerc20 *Anyerc20Session) Minters(arg0 *big.Int) (common.Address, error) {
	return _Anyerc20.Contract.Minters(&_Anyerc20.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0x8623ec7b.
//
// Solidity: function minters(uint256 ) view returns(address)
func (_Anyerc20 *Anyerc20CallerSession) Minters(arg0 *big.Int) (common.Address, error) {
	return _Anyerc20.Contract.Minters(&_Anyerc20.CallOpts, arg0)
}

// Mpc is a free data retrieval call binding the contract method 0xf75c2664.
//
// Solidity: function mpc() view returns(address)
func (_Anyerc20 *Anyerc20Caller) Mpc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "mpc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mpc is a free data retrieval call binding the contract method 0xf75c2664.
//
// Solidity: function mpc() view returns(address)
func (_Anyerc20 *Anyerc20Session) Mpc() (common.Address, error) {
	return _Anyerc20.Contract.Mpc(&_Anyerc20.CallOpts)
}

// Mpc is a free data retrieval call binding the contract method 0xf75c2664.
//
// Solidity: function mpc() view returns(address)
func (_Anyerc20 *Anyerc20CallerSession) Mpc() (common.Address, error) {
	return _Anyerc20.Contract.Mpc(&_Anyerc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Anyerc20 *Anyerc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Anyerc20 *Anyerc20Session) Name() (string, error) {
	return _Anyerc20.Contract.Name(&_Anyerc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Anyerc20 *Anyerc20CallerSession) Name() (string, error) {
	return _Anyerc20.Contract.Name(&_Anyerc20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Anyerc20 *Anyerc20Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Anyerc20 *Anyerc20Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Anyerc20.Contract.Nonces(&_Anyerc20.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Anyerc20 *Anyerc20CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Anyerc20.Contract.Nonces(&_Anyerc20.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Anyerc20 *Anyerc20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Anyerc20 *Anyerc20Session) Owner() (common.Address, error) {
	return _Anyerc20.Contract.Owner(&_Anyerc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Anyerc20 *Anyerc20CallerSession) Owner() (common.Address, error) {
	return _Anyerc20.Contract.Owner(&_Anyerc20.CallOpts)
}

// PendingDelay is a free data retrieval call binding the contract method 0x4ca8f0ed.
//
// Solidity: function pendingDelay() view returns(uint256)
func (_Anyerc20 *Anyerc20Caller) PendingDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "pendingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingDelay is a free data retrieval call binding the contract method 0x4ca8f0ed.
//
// Solidity: function pendingDelay() view returns(uint256)
func (_Anyerc20 *Anyerc20Session) PendingDelay() (*big.Int, error) {
	return _Anyerc20.Contract.PendingDelay(&_Anyerc20.CallOpts)
}

// PendingDelay is a free data retrieval call binding the contract method 0x4ca8f0ed.
//
// Solidity: function pendingDelay() view returns(uint256)
func (_Anyerc20 *Anyerc20CallerSession) PendingDelay() (*big.Int, error) {
	return _Anyerc20.Contract.PendingDelay(&_Anyerc20.CallOpts)
}

// PendingMinter is a free data retrieval call binding the contract method 0x91c5df49.
//
// Solidity: function pendingMinter() view returns(address)
func (_Anyerc20 *Anyerc20Caller) PendingMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "pendingMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingMinter is a free data retrieval call binding the contract method 0x91c5df49.
//
// Solidity: function pendingMinter() view returns(address)
func (_Anyerc20 *Anyerc20Session) PendingMinter() (common.Address, error) {
	return _Anyerc20.Contract.PendingMinter(&_Anyerc20.CallOpts)
}

// PendingMinter is a free data retrieval call binding the contract method 0x91c5df49.
//
// Solidity: function pendingMinter() view returns(address)
func (_Anyerc20 *Anyerc20CallerSession) PendingMinter() (common.Address, error) {
	return _Anyerc20.Contract.PendingMinter(&_Anyerc20.CallOpts)
}

// PendingVault is a free data retrieval call binding the contract method 0x52113ba7.
//
// Solidity: function pendingVault() view returns(address)
func (_Anyerc20 *Anyerc20Caller) PendingVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "pendingVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingVault is a free data retrieval call binding the contract method 0x52113ba7.
//
// Solidity: function pendingVault() view returns(address)
func (_Anyerc20 *Anyerc20Session) PendingVault() (common.Address, error) {
	return _Anyerc20.Contract.PendingVault(&_Anyerc20.CallOpts)
}

// PendingVault is a free data retrieval call binding the contract method 0x52113ba7.
//
// Solidity: function pendingVault() view returns(address)
func (_Anyerc20 *Anyerc20CallerSession) PendingVault() (common.Address, error) {
	return _Anyerc20.Contract.PendingVault(&_Anyerc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Anyerc20 *Anyerc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Anyerc20 *Anyerc20Session) Symbol() (string, error) {
	return _Anyerc20.Contract.Symbol(&_Anyerc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Anyerc20 *Anyerc20CallerSession) Symbol() (string, error) {
	return _Anyerc20.Contract.Symbol(&_Anyerc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Anyerc20 *Anyerc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Anyerc20 *Anyerc20Session) TotalSupply() (*big.Int, error) {
	return _Anyerc20.Contract.TotalSupply(&_Anyerc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Anyerc20 *Anyerc20CallerSession) TotalSupply() (*big.Int, error) {
	return _Anyerc20.Contract.TotalSupply(&_Anyerc20.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Anyerc20 *Anyerc20Caller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Anyerc20 *Anyerc20Session) Underlying() (common.Address, error) {
	return _Anyerc20.Contract.Underlying(&_Anyerc20.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Anyerc20 *Anyerc20CallerSession) Underlying() (common.Address, error) {
	return _Anyerc20.Contract.Underlying(&_Anyerc20.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Anyerc20 *Anyerc20Caller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Anyerc20.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Anyerc20 *Anyerc20Session) Vault() (common.Address, error) {
	return _Anyerc20.Contract.Vault(&_Anyerc20.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Anyerc20 *Anyerc20CallerSession) Vault() (common.Address, error) {
	return _Anyerc20.Contract.Vault(&_Anyerc20.CallOpts)
}

// Swapin is a paid mutator transaction binding the contract method 0xec126c77.
//
// Solidity: function Swapin(bytes32 txhash, address account, uint256 amount) returns(bool)
func (_Anyerc20 *Anyerc20Transactor) Swapin(opts *bind.TransactOpts, txhash [32]byte, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "Swapin", txhash, account, amount)
}

// Swapin is a paid mutator transaction binding the contract method 0xec126c77.
//
// Solidity: function Swapin(bytes32 txhash, address account, uint256 amount) returns(bool)
func (_Anyerc20 *Anyerc20Session) Swapin(txhash [32]byte, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.Swapin(&_Anyerc20.TransactOpts, txhash, account, amount)
}

// Swapin is a paid mutator transaction binding the contract method 0xec126c77.
//
// Solidity: function Swapin(bytes32 txhash, address account, uint256 amount) returns(bool)
func (_Anyerc20 *Anyerc20TransactorSession) Swapin(txhash [32]byte, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.Swapin(&_Anyerc20.TransactOpts, txhash, account, amount)
}

// Swapout is a paid mutator transaction binding the contract method 0x628d6cba.
//
// Solidity: function Swapout(uint256 amount, address bindaddr) returns(bool)
func (_Anyerc20 *Anyerc20Transactor) Swapout(opts *bind.TransactOpts, amount *big.Int, bindaddr common.Address) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "Swapout", amount, bindaddr)
}

// Swapout is a paid mutator transaction binding the contract method 0x628d6cba.
//
// Solidity: function Swapout(uint256 amount, address bindaddr) returns(bool)
func (_Anyerc20 *Anyerc20Session) Swapout(amount *big.Int, bindaddr common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.Swapout(&_Anyerc20.TransactOpts, amount, bindaddr)
}

// Swapout is a paid mutator transaction binding the contract method 0x628d6cba.
//
// Solidity: function Swapout(uint256 amount, address bindaddr) returns(bool)
func (_Anyerc20 *Anyerc20TransactorSession) Swapout(amount *big.Int, bindaddr common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.Swapout(&_Anyerc20.TransactOpts, amount, bindaddr)
}

// ApplyMinter is a paid mutator transaction binding the contract method 0x0d707df8.
//
// Solidity: function applyMinter() returns()
func (_Anyerc20 *Anyerc20Transactor) ApplyMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "applyMinter")
}

// ApplyMinter is a paid mutator transaction binding the contract method 0x0d707df8.
//
// Solidity: function applyMinter() returns()
func (_Anyerc20 *Anyerc20Session) ApplyMinter() (*types.Transaction, error) {
	return _Anyerc20.Contract.ApplyMinter(&_Anyerc20.TransactOpts)
}

// ApplyMinter is a paid mutator transaction binding the contract method 0x0d707df8.
//
// Solidity: function applyMinter() returns()
func (_Anyerc20 *Anyerc20TransactorSession) ApplyMinter() (*types.Transaction, error) {
	return _Anyerc20.Contract.ApplyMinter(&_Anyerc20.TransactOpts)
}

// ApplyVault is a paid mutator transaction binding the contract method 0xd93f2445.
//
// Solidity: function applyVault() returns()
func (_Anyerc20 *Anyerc20Transactor) ApplyVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "applyVault")
}

// ApplyVault is a paid mutator transaction binding the contract method 0xd93f2445.
//
// Solidity: function applyVault() returns()
func (_Anyerc20 *Anyerc20Session) ApplyVault() (*types.Transaction, error) {
	return _Anyerc20.Contract.ApplyVault(&_Anyerc20.TransactOpts)
}

// ApplyVault is a paid mutator transaction binding the contract method 0xd93f2445.
//
// Solidity: function applyVault() returns()
func (_Anyerc20 *Anyerc20TransactorSession) ApplyVault() (*types.Transaction, error) {
	return _Anyerc20.Contract.ApplyVault(&_Anyerc20.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Anyerc20 *Anyerc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Anyerc20 *Anyerc20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.Approve(&_Anyerc20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Anyerc20 *Anyerc20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.Approve(&_Anyerc20.TransactOpts, spender, value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_Anyerc20 *Anyerc20Transactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "approveAndCall", spender, value, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_Anyerc20 *Anyerc20Session) ApproveAndCall(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Anyerc20.Contract.ApproveAndCall(&_Anyerc20.TransactOpts, spender, value, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_Anyerc20 *Anyerc20TransactorSession) ApproveAndCall(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Anyerc20.Contract.ApproveAndCall(&_Anyerc20.TransactOpts, spender, value, data)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns(bool)
func (_Anyerc20 *Anyerc20Transactor) Burn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "burn", from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns(bool)
func (_Anyerc20 *Anyerc20Session) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.Burn(&_Anyerc20.TransactOpts, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns(bool)
func (_Anyerc20 *Anyerc20TransactorSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.Burn(&_Anyerc20.TransactOpts, from, amount)
}

// ChangeMPCOwner is a paid mutator transaction binding the contract method 0x5f9b105d.
//
// Solidity: function changeMPCOwner(address newVault) returns(bool)
func (_Anyerc20 *Anyerc20Transactor) ChangeMPCOwner(opts *bind.TransactOpts, newVault common.Address) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "changeMPCOwner", newVault)
}

// ChangeMPCOwner is a paid mutator transaction binding the contract method 0x5f9b105d.
//
// Solidity: function changeMPCOwner(address newVault) returns(bool)
func (_Anyerc20 *Anyerc20Session) ChangeMPCOwner(newVault common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.ChangeMPCOwner(&_Anyerc20.TransactOpts, newVault)
}

// ChangeMPCOwner is a paid mutator transaction binding the contract method 0x5f9b105d.
//
// Solidity: function changeMPCOwner(address newVault) returns(bool)
func (_Anyerc20 *Anyerc20TransactorSession) ChangeMPCOwner(newVault common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.ChangeMPCOwner(&_Anyerc20.TransactOpts, newVault)
}

// ChangeVault is a paid mutator transaction binding the contract method 0x60e232a9.
//
// Solidity: function changeVault(address newVault) returns(bool)
func (_Anyerc20 *Anyerc20Transactor) ChangeVault(opts *bind.TransactOpts, newVault common.Address) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "changeVault", newVault)
}

// ChangeVault is a paid mutator transaction binding the contract method 0x60e232a9.
//
// Solidity: function changeVault(address newVault) returns(bool)
func (_Anyerc20 *Anyerc20Session) ChangeVault(newVault common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.ChangeVault(&_Anyerc20.TransactOpts, newVault)
}

// ChangeVault is a paid mutator transaction binding the contract method 0x60e232a9.
//
// Solidity: function changeVault(address newVault) returns(bool)
func (_Anyerc20 *Anyerc20TransactorSession) ChangeVault(newVault common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.ChangeVault(&_Anyerc20.TransactOpts, newVault)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address to) returns(uint256)
func (_Anyerc20 *Anyerc20Transactor) Deposit(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "deposit", amount, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address to) returns(uint256)
func (_Anyerc20 *Anyerc20Session) Deposit(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.Deposit(&_Anyerc20.TransactOpts, amount, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address to) returns(uint256)
func (_Anyerc20 *Anyerc20TransactorSession) Deposit(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.Deposit(&_Anyerc20.TransactOpts, amount, to)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns(uint256)
func (_Anyerc20 *Anyerc20Transactor) Deposit0(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "deposit0", amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns(uint256)
func (_Anyerc20 *Anyerc20Session) Deposit0(amount *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.Deposit0(&_Anyerc20.TransactOpts, amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns(uint256)
func (_Anyerc20 *Anyerc20TransactorSession) Deposit0(amount *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.Deposit0(&_Anyerc20.TransactOpts, amount)
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Anyerc20 *Anyerc20Transactor) Deposit1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "deposit1")
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Anyerc20 *Anyerc20Session) Deposit1() (*types.Transaction, error) {
	return _Anyerc20.Contract.Deposit1(&_Anyerc20.TransactOpts)
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Anyerc20 *Anyerc20TransactorSession) Deposit1() (*types.Transaction, error) {
	return _Anyerc20.Contract.Deposit1(&_Anyerc20.TransactOpts)
}

// DepositVault is a paid mutator transaction binding the contract method 0xbebbf4d0.
//
// Solidity: function depositVault(uint256 amount, address to) returns(uint256)
func (_Anyerc20 *Anyerc20Transactor) DepositVault(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "depositVault", amount, to)
}

// DepositVault is a paid mutator transaction binding the contract method 0xbebbf4d0.
//
// Solidity: function depositVault(uint256 amount, address to) returns(uint256)
func (_Anyerc20 *Anyerc20Session) DepositVault(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.DepositVault(&_Anyerc20.TransactOpts, amount, to)
}

// DepositVault is a paid mutator transaction binding the contract method 0xbebbf4d0.
//
// Solidity: function depositVault(uint256 amount, address to) returns(uint256)
func (_Anyerc20 *Anyerc20TransactorSession) DepositVault(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.DepositVault(&_Anyerc20.TransactOpts, amount, to)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x81a37c18.
//
// Solidity: function depositWithPermit(address target, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s, address to) returns(uint256)
func (_Anyerc20 *Anyerc20Transactor) DepositWithPermit(opts *bind.TransactOpts, target common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "depositWithPermit", target, value, deadline, v, r, s, to)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x81a37c18.
//
// Solidity: function depositWithPermit(address target, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s, address to) returns(uint256)
func (_Anyerc20 *Anyerc20Session) DepositWithPermit(target common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.DepositWithPermit(&_Anyerc20.TransactOpts, target, value, deadline, v, r, s, to)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x81a37c18.
//
// Solidity: function depositWithPermit(address target, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s, address to) returns(uint256)
func (_Anyerc20 *Anyerc20TransactorSession) DepositWithPermit(target common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.DepositWithPermit(&_Anyerc20.TransactOpts, target, value, deadline, v, r, s, to)
}

// DepositWithTransferPermit is a paid mutator transaction binding the contract method 0xf954734e.
//
// Solidity: function depositWithTransferPermit(address target, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s, address to) returns(uint256)
func (_Anyerc20 *Anyerc20Transactor) DepositWithTransferPermit(opts *bind.TransactOpts, target common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "depositWithTransferPermit", target, value, deadline, v, r, s, to)
}

// DepositWithTransferPermit is a paid mutator transaction binding the contract method 0xf954734e.
//
// Solidity: function depositWithTransferPermit(address target, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s, address to) returns(uint256)
func (_Anyerc20 *Anyerc20Session) DepositWithTransferPermit(target common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.DepositWithTransferPermit(&_Anyerc20.TransactOpts, target, value, deadline, v, r, s, to)
}

// DepositWithTransferPermit is a paid mutator transaction binding the contract method 0xf954734e.
//
// Solidity: function depositWithTransferPermit(address target, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s, address to) returns(uint256)
func (_Anyerc20 *Anyerc20TransactorSession) DepositWithTransferPermit(target common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.DepositWithTransferPermit(&_Anyerc20.TransactOpts, target, value, deadline, v, r, s, to)
}

// InitVault is a paid mutator transaction binding the contract method 0x2ebe3fbb.
//
// Solidity: function initVault(address _vault) returns()
func (_Anyerc20 *Anyerc20Transactor) InitVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "initVault", _vault)
}

// InitVault is a paid mutator transaction binding the contract method 0x2ebe3fbb.
//
// Solidity: function initVault(address _vault) returns()
func (_Anyerc20 *Anyerc20Session) InitVault(_vault common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.InitVault(&_Anyerc20.TransactOpts, _vault)
}

// InitVault is a paid mutator transaction binding the contract method 0x2ebe3fbb.
//
// Solidity: function initVault(address _vault) returns()
func (_Anyerc20 *Anyerc20TransactorSession) InitVault(_vault common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.InitVault(&_Anyerc20.TransactOpts, _vault)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_Anyerc20 *Anyerc20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_Anyerc20 *Anyerc20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.Mint(&_Anyerc20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_Anyerc20 *Anyerc20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.Mint(&_Anyerc20.TransactOpts, to, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address target, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Anyerc20 *Anyerc20Transactor) Permit(opts *bind.TransactOpts, target common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "permit", target, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address target, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Anyerc20 *Anyerc20Session) Permit(target common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Anyerc20.Contract.Permit(&_Anyerc20.TransactOpts, target, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address target, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Anyerc20 *Anyerc20TransactorSession) Permit(target common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Anyerc20.Contract.Permit(&_Anyerc20.TransactOpts, target, spender, value, deadline, v, r, s)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address _auth) returns()
func (_Anyerc20 *Anyerc20Transactor) RevokeMinter(opts *bind.TransactOpts, _auth common.Address) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "revokeMinter", _auth)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address _auth) returns()
func (_Anyerc20 *Anyerc20Session) RevokeMinter(_auth common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.RevokeMinter(&_Anyerc20.TransactOpts, _auth)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address _auth) returns()
func (_Anyerc20 *Anyerc20TransactorSession) RevokeMinter(_auth common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.RevokeMinter(&_Anyerc20.TransactOpts, _auth)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _auth) returns()
func (_Anyerc20 *Anyerc20Transactor) SetMinter(opts *bind.TransactOpts, _auth common.Address) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "setMinter", _auth)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _auth) returns()
func (_Anyerc20 *Anyerc20Session) SetMinter(_auth common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.SetMinter(&_Anyerc20.TransactOpts, _auth)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _auth) returns()
func (_Anyerc20 *Anyerc20TransactorSession) SetMinter(_auth common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.SetMinter(&_Anyerc20.TransactOpts, _auth)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_Anyerc20 *Anyerc20Transactor) SetVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "setVault", _vault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_Anyerc20 *Anyerc20Session) SetVault(_vault common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.SetVault(&_Anyerc20.TransactOpts, _vault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_Anyerc20 *Anyerc20TransactorSession) SetVault(_vault common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.SetVault(&_Anyerc20.TransactOpts, _vault)
}

// SetVaultOnly is a paid mutator transaction binding the contract method 0xc4b740f5.
//
// Solidity: function setVaultOnly(bool enabled) returns()
func (_Anyerc20 *Anyerc20Transactor) SetVaultOnly(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "setVaultOnly", enabled)
}

// SetVaultOnly is a paid mutator transaction binding the contract method 0xc4b740f5.
//
// Solidity: function setVaultOnly(bool enabled) returns()
func (_Anyerc20 *Anyerc20Session) SetVaultOnly(enabled bool) (*types.Transaction, error) {
	return _Anyerc20.Contract.SetVaultOnly(&_Anyerc20.TransactOpts, enabled)
}

// SetVaultOnly is a paid mutator transaction binding the contract method 0xc4b740f5.
//
// Solidity: function setVaultOnly(bool enabled) returns()
func (_Anyerc20 *Anyerc20TransactorSession) SetVaultOnly(enabled bool) (*types.Transaction, error) {
	return _Anyerc20.Contract.SetVaultOnly(&_Anyerc20.TransactOpts, enabled)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Anyerc20 *Anyerc20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Anyerc20 *Anyerc20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.Transfer(&_Anyerc20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Anyerc20 *Anyerc20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.Transfer(&_Anyerc20.TransactOpts, to, value)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_Anyerc20 *Anyerc20Transactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "transferAndCall", to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_Anyerc20 *Anyerc20Session) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Anyerc20.Contract.TransferAndCall(&_Anyerc20.TransactOpts, to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_Anyerc20 *Anyerc20TransactorSession) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Anyerc20.Contract.TransferAndCall(&_Anyerc20.TransactOpts, to, value, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Anyerc20 *Anyerc20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Anyerc20 *Anyerc20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.TransferFrom(&_Anyerc20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Anyerc20 *Anyerc20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.TransferFrom(&_Anyerc20.TransactOpts, from, to, value)
}

// TransferWithPermit is a paid mutator transaction binding the contract method 0x605629d6.
//
// Solidity: function transferWithPermit(address target, address to, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Anyerc20 *Anyerc20Transactor) TransferWithPermit(opts *bind.TransactOpts, target common.Address, to common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "transferWithPermit", target, to, value, deadline, v, r, s)
}

// TransferWithPermit is a paid mutator transaction binding the contract method 0x605629d6.
//
// Solidity: function transferWithPermit(address target, address to, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Anyerc20 *Anyerc20Session) TransferWithPermit(target common.Address, to common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Anyerc20.Contract.TransferWithPermit(&_Anyerc20.TransactOpts, target, to, value, deadline, v, r, s)
}

// TransferWithPermit is a paid mutator transaction binding the contract method 0x605629d6.
//
// Solidity: function transferWithPermit(address target, address to, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Anyerc20 *Anyerc20TransactorSession) TransferWithPermit(target common.Address, to common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Anyerc20.Contract.TransferWithPermit(&_Anyerc20.TransactOpts, target, to, value, deadline, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address to) returns(uint256)
func (_Anyerc20 *Anyerc20Transactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "withdraw", amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address to) returns(uint256)
func (_Anyerc20 *Anyerc20Session) Withdraw(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.Withdraw(&_Anyerc20.TransactOpts, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address to) returns(uint256)
func (_Anyerc20 *Anyerc20TransactorSession) Withdraw(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.Withdraw(&_Anyerc20.TransactOpts, amount, to)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(uint256)
func (_Anyerc20 *Anyerc20Transactor) Withdraw0(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "withdraw0", amount)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(uint256)
func (_Anyerc20 *Anyerc20Session) Withdraw0(amount *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.Withdraw0(&_Anyerc20.TransactOpts, amount)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(uint256)
func (_Anyerc20 *Anyerc20TransactorSession) Withdraw0(amount *big.Int) (*types.Transaction, error) {
	return _Anyerc20.Contract.Withdraw0(&_Anyerc20.TransactOpts, amount)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Anyerc20 *Anyerc20Transactor) Withdraw1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "withdraw1")
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Anyerc20 *Anyerc20Session) Withdraw1() (*types.Transaction, error) {
	return _Anyerc20.Contract.Withdraw1(&_Anyerc20.TransactOpts)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Anyerc20 *Anyerc20TransactorSession) Withdraw1() (*types.Transaction, error) {
	return _Anyerc20.Contract.Withdraw1(&_Anyerc20.TransactOpts)
}

// WithdrawVault is a paid mutator transaction binding the contract method 0x0039d6ec.
//
// Solidity: function withdrawVault(address from, uint256 amount, address to) returns(uint256)
func (_Anyerc20 *Anyerc20Transactor) WithdrawVault(opts *bind.TransactOpts, from common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.contract.Transact(opts, "withdrawVault", from, amount, to)
}

// WithdrawVault is a paid mutator transaction binding the contract method 0x0039d6ec.
//
// Solidity: function withdrawVault(address from, uint256 amount, address to) returns(uint256)
func (_Anyerc20 *Anyerc20Session) WithdrawVault(from common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.WithdrawVault(&_Anyerc20.TransactOpts, from, amount, to)
}

// WithdrawVault is a paid mutator transaction binding the contract method 0x0039d6ec.
//
// Solidity: function withdrawVault(address from, uint256 amount, address to) returns(uint256)
func (_Anyerc20 *Anyerc20TransactorSession) WithdrawVault(from common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Anyerc20.Contract.WithdrawVault(&_Anyerc20.TransactOpts, from, amount, to)
}

// Anyerc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Anyerc20 contract.
type Anyerc20ApprovalIterator struct {
	Event *Anyerc20Approval // Event containing the contract specifics and raw log

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
func (it *Anyerc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Anyerc20Approval)
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
		it.Event = new(Anyerc20Approval)
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
func (it *Anyerc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Anyerc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Anyerc20Approval represents a Approval event raised by the Anyerc20 contract.
type Anyerc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Anyerc20 *Anyerc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Anyerc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Anyerc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Anyerc20ApprovalIterator{contract: _Anyerc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Anyerc20 *Anyerc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Anyerc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Anyerc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Anyerc20Approval)
				if err := _Anyerc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Anyerc20 *Anyerc20Filterer) ParseApproval(log types.Log) (*Anyerc20Approval, error) {
	event := new(Anyerc20Approval)
	if err := _Anyerc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Anyerc20LogAddAuthIterator is returned from FilterLogAddAuth and is used to iterate over the raw logs and unpacked data for LogAddAuth events raised by the Anyerc20 contract.
type Anyerc20LogAddAuthIterator struct {
	Event *Anyerc20LogAddAuth // Event containing the contract specifics and raw log

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
func (it *Anyerc20LogAddAuthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Anyerc20LogAddAuth)
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
		it.Event = new(Anyerc20LogAddAuth)
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
func (it *Anyerc20LogAddAuthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Anyerc20LogAddAuthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Anyerc20LogAddAuth represents a LogAddAuth event raised by the Anyerc20 contract.
type Anyerc20LogAddAuth struct {
	Auth      common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogAddAuth is a free log retrieval operation binding the contract event 0xff9be4a5a0b9027fd253167d4c170ef1bbf8403af21bf06a0ed87ac8c8ecb5c6.
//
// Solidity: event LogAddAuth(address indexed auth, uint256 timestamp)
func (_Anyerc20 *Anyerc20Filterer) FilterLogAddAuth(opts *bind.FilterOpts, auth []common.Address) (*Anyerc20LogAddAuthIterator, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}

	logs, sub, err := _Anyerc20.contract.FilterLogs(opts, "LogAddAuth", authRule)
	if err != nil {
		return nil, err
	}
	return &Anyerc20LogAddAuthIterator{contract: _Anyerc20.contract, event: "LogAddAuth", logs: logs, sub: sub}, nil
}

// WatchLogAddAuth is a free log subscription operation binding the contract event 0xff9be4a5a0b9027fd253167d4c170ef1bbf8403af21bf06a0ed87ac8c8ecb5c6.
//
// Solidity: event LogAddAuth(address indexed auth, uint256 timestamp)
func (_Anyerc20 *Anyerc20Filterer) WatchLogAddAuth(opts *bind.WatchOpts, sink chan<- *Anyerc20LogAddAuth, auth []common.Address) (event.Subscription, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}

	logs, sub, err := _Anyerc20.contract.WatchLogs(opts, "LogAddAuth", authRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Anyerc20LogAddAuth)
				if err := _Anyerc20.contract.UnpackLog(event, "LogAddAuth", log); err != nil {
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

// ParseLogAddAuth is a log parse operation binding the contract event 0xff9be4a5a0b9027fd253167d4c170ef1bbf8403af21bf06a0ed87ac8c8ecb5c6.
//
// Solidity: event LogAddAuth(address indexed auth, uint256 timestamp)
func (_Anyerc20 *Anyerc20Filterer) ParseLogAddAuth(log types.Log) (*Anyerc20LogAddAuth, error) {
	event := new(Anyerc20LogAddAuth)
	if err := _Anyerc20.contract.UnpackLog(event, "LogAddAuth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Anyerc20LogChangeMPCOwnerIterator is returned from FilterLogChangeMPCOwner and is used to iterate over the raw logs and unpacked data for LogChangeMPCOwner events raised by the Anyerc20 contract.
type Anyerc20LogChangeMPCOwnerIterator struct {
	Event *Anyerc20LogChangeMPCOwner // Event containing the contract specifics and raw log

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
func (it *Anyerc20LogChangeMPCOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Anyerc20LogChangeMPCOwner)
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
		it.Event = new(Anyerc20LogChangeMPCOwner)
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
func (it *Anyerc20LogChangeMPCOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Anyerc20LogChangeMPCOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Anyerc20LogChangeMPCOwner represents a LogChangeMPCOwner event raised by the Anyerc20 contract.
type Anyerc20LogChangeMPCOwner struct {
	OldOwner        common.Address
	NewOwner        common.Address
	EffectiveHeight *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogChangeMPCOwner is a free log retrieval operation binding the contract event 0x1d065115f314fb9bad9557bd5460b9e3c66f7223b1dd04e73e828f0bb5afe89f.
//
// Solidity: event LogChangeMPCOwner(address indexed oldOwner, address indexed newOwner, uint256 indexed effectiveHeight)
func (_Anyerc20 *Anyerc20Filterer) FilterLogChangeMPCOwner(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address, effectiveHeight []*big.Int) (*Anyerc20LogChangeMPCOwnerIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}
	var effectiveHeightRule []interface{}
	for _, effectiveHeightItem := range effectiveHeight {
		effectiveHeightRule = append(effectiveHeightRule, effectiveHeightItem)
	}

	logs, sub, err := _Anyerc20.contract.FilterLogs(opts, "LogChangeMPCOwner", oldOwnerRule, newOwnerRule, effectiveHeightRule)
	if err != nil {
		return nil, err
	}
	return &Anyerc20LogChangeMPCOwnerIterator{contract: _Anyerc20.contract, event: "LogChangeMPCOwner", logs: logs, sub: sub}, nil
}

// WatchLogChangeMPCOwner is a free log subscription operation binding the contract event 0x1d065115f314fb9bad9557bd5460b9e3c66f7223b1dd04e73e828f0bb5afe89f.
//
// Solidity: event LogChangeMPCOwner(address indexed oldOwner, address indexed newOwner, uint256 indexed effectiveHeight)
func (_Anyerc20 *Anyerc20Filterer) WatchLogChangeMPCOwner(opts *bind.WatchOpts, sink chan<- *Anyerc20LogChangeMPCOwner, oldOwner []common.Address, newOwner []common.Address, effectiveHeight []*big.Int) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}
	var effectiveHeightRule []interface{}
	for _, effectiveHeightItem := range effectiveHeight {
		effectiveHeightRule = append(effectiveHeightRule, effectiveHeightItem)
	}

	logs, sub, err := _Anyerc20.contract.WatchLogs(opts, "LogChangeMPCOwner", oldOwnerRule, newOwnerRule, effectiveHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Anyerc20LogChangeMPCOwner)
				if err := _Anyerc20.contract.UnpackLog(event, "LogChangeMPCOwner", log); err != nil {
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

// ParseLogChangeMPCOwner is a log parse operation binding the contract event 0x1d065115f314fb9bad9557bd5460b9e3c66f7223b1dd04e73e828f0bb5afe89f.
//
// Solidity: event LogChangeMPCOwner(address indexed oldOwner, address indexed newOwner, uint256 indexed effectiveHeight)
func (_Anyerc20 *Anyerc20Filterer) ParseLogChangeMPCOwner(log types.Log) (*Anyerc20LogChangeMPCOwner, error) {
	event := new(Anyerc20LogChangeMPCOwner)
	if err := _Anyerc20.contract.UnpackLog(event, "LogChangeMPCOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Anyerc20LogChangeVaultIterator is returned from FilterLogChangeVault and is used to iterate over the raw logs and unpacked data for LogChangeVault events raised by the Anyerc20 contract.
type Anyerc20LogChangeVaultIterator struct {
	Event *Anyerc20LogChangeVault // Event containing the contract specifics and raw log

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
func (it *Anyerc20LogChangeVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Anyerc20LogChangeVault)
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
		it.Event = new(Anyerc20LogChangeVault)
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
func (it *Anyerc20LogChangeVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Anyerc20LogChangeVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Anyerc20LogChangeVault represents a LogChangeVault event raised by the Anyerc20 contract.
type Anyerc20LogChangeVault struct {
	OldVault      common.Address
	NewVault      common.Address
	EffectiveTime *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLogChangeVault is a free log retrieval operation binding the contract event 0x5c364079e7102c27c608f9b237c735a1b7bfa0b67f27c2ad26bad447bf965cac.
//
// Solidity: event LogChangeVault(address indexed oldVault, address indexed newVault, uint256 indexed effectiveTime)
func (_Anyerc20 *Anyerc20Filterer) FilterLogChangeVault(opts *bind.FilterOpts, oldVault []common.Address, newVault []common.Address, effectiveTime []*big.Int) (*Anyerc20LogChangeVaultIterator, error) {

	var oldVaultRule []interface{}
	for _, oldVaultItem := range oldVault {
		oldVaultRule = append(oldVaultRule, oldVaultItem)
	}
	var newVaultRule []interface{}
	for _, newVaultItem := range newVault {
		newVaultRule = append(newVaultRule, newVaultItem)
	}
	var effectiveTimeRule []interface{}
	for _, effectiveTimeItem := range effectiveTime {
		effectiveTimeRule = append(effectiveTimeRule, effectiveTimeItem)
	}

	logs, sub, err := _Anyerc20.contract.FilterLogs(opts, "LogChangeVault", oldVaultRule, newVaultRule, effectiveTimeRule)
	if err != nil {
		return nil, err
	}
	return &Anyerc20LogChangeVaultIterator{contract: _Anyerc20.contract, event: "LogChangeVault", logs: logs, sub: sub}, nil
}

// WatchLogChangeVault is a free log subscription operation binding the contract event 0x5c364079e7102c27c608f9b237c735a1b7bfa0b67f27c2ad26bad447bf965cac.
//
// Solidity: event LogChangeVault(address indexed oldVault, address indexed newVault, uint256 indexed effectiveTime)
func (_Anyerc20 *Anyerc20Filterer) WatchLogChangeVault(opts *bind.WatchOpts, sink chan<- *Anyerc20LogChangeVault, oldVault []common.Address, newVault []common.Address, effectiveTime []*big.Int) (event.Subscription, error) {

	var oldVaultRule []interface{}
	for _, oldVaultItem := range oldVault {
		oldVaultRule = append(oldVaultRule, oldVaultItem)
	}
	var newVaultRule []interface{}
	for _, newVaultItem := range newVault {
		newVaultRule = append(newVaultRule, newVaultItem)
	}
	var effectiveTimeRule []interface{}
	for _, effectiveTimeItem := range effectiveTime {
		effectiveTimeRule = append(effectiveTimeRule, effectiveTimeItem)
	}

	logs, sub, err := _Anyerc20.contract.WatchLogs(opts, "LogChangeVault", oldVaultRule, newVaultRule, effectiveTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Anyerc20LogChangeVault)
				if err := _Anyerc20.contract.UnpackLog(event, "LogChangeVault", log); err != nil {
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

// ParseLogChangeVault is a log parse operation binding the contract event 0x5c364079e7102c27c608f9b237c735a1b7bfa0b67f27c2ad26bad447bf965cac.
//
// Solidity: event LogChangeVault(address indexed oldVault, address indexed newVault, uint256 indexed effectiveTime)
func (_Anyerc20 *Anyerc20Filterer) ParseLogChangeVault(log types.Log) (*Anyerc20LogChangeVault, error) {
	event := new(Anyerc20LogChangeVault)
	if err := _Anyerc20.contract.UnpackLog(event, "LogChangeVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Anyerc20LogSwapinIterator is returned from FilterLogSwapin and is used to iterate over the raw logs and unpacked data for LogSwapin events raised by the Anyerc20 contract.
type Anyerc20LogSwapinIterator struct {
	Event *Anyerc20LogSwapin // Event containing the contract specifics and raw log

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
func (it *Anyerc20LogSwapinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Anyerc20LogSwapin)
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
		it.Event = new(Anyerc20LogSwapin)
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
func (it *Anyerc20LogSwapinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Anyerc20LogSwapinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Anyerc20LogSwapin represents a LogSwapin event raised by the Anyerc20 contract.
type Anyerc20LogSwapin struct {
	Txhash  [32]byte
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogSwapin is a free log retrieval operation binding the contract event 0x05d0634fe981be85c22e2942a880821b70095d84e152c3ea3c17a4e4250d9d61.
//
// Solidity: event LogSwapin(bytes32 indexed txhash, address indexed account, uint256 amount)
func (_Anyerc20 *Anyerc20Filterer) FilterLogSwapin(opts *bind.FilterOpts, txhash [][32]byte, account []common.Address) (*Anyerc20LogSwapinIterator, error) {

	var txhashRule []interface{}
	for _, txhashItem := range txhash {
		txhashRule = append(txhashRule, txhashItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Anyerc20.contract.FilterLogs(opts, "LogSwapin", txhashRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &Anyerc20LogSwapinIterator{contract: _Anyerc20.contract, event: "LogSwapin", logs: logs, sub: sub}, nil
}

// WatchLogSwapin is a free log subscription operation binding the contract event 0x05d0634fe981be85c22e2942a880821b70095d84e152c3ea3c17a4e4250d9d61.
//
// Solidity: event LogSwapin(bytes32 indexed txhash, address indexed account, uint256 amount)
func (_Anyerc20 *Anyerc20Filterer) WatchLogSwapin(opts *bind.WatchOpts, sink chan<- *Anyerc20LogSwapin, txhash [][32]byte, account []common.Address) (event.Subscription, error) {

	var txhashRule []interface{}
	for _, txhashItem := range txhash {
		txhashRule = append(txhashRule, txhashItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Anyerc20.contract.WatchLogs(opts, "LogSwapin", txhashRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Anyerc20LogSwapin)
				if err := _Anyerc20.contract.UnpackLog(event, "LogSwapin", log); err != nil {
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

// ParseLogSwapin is a log parse operation binding the contract event 0x05d0634fe981be85c22e2942a880821b70095d84e152c3ea3c17a4e4250d9d61.
//
// Solidity: event LogSwapin(bytes32 indexed txhash, address indexed account, uint256 amount)
func (_Anyerc20 *Anyerc20Filterer) ParseLogSwapin(log types.Log) (*Anyerc20LogSwapin, error) {
	event := new(Anyerc20LogSwapin)
	if err := _Anyerc20.contract.UnpackLog(event, "LogSwapin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Anyerc20LogSwapoutIterator is returned from FilterLogSwapout and is used to iterate over the raw logs and unpacked data for LogSwapout events raised by the Anyerc20 contract.
type Anyerc20LogSwapoutIterator struct {
	Event *Anyerc20LogSwapout // Event containing the contract specifics and raw log

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
func (it *Anyerc20LogSwapoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Anyerc20LogSwapout)
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
		it.Event = new(Anyerc20LogSwapout)
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
func (it *Anyerc20LogSwapoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Anyerc20LogSwapoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Anyerc20LogSwapout represents a LogSwapout event raised by the Anyerc20 contract.
type Anyerc20LogSwapout struct {
	Account  common.Address
	Bindaddr common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogSwapout is a free log retrieval operation binding the contract event 0x6b616089d04950dc06c45c6dd787d657980543f89651aec47924752c7d16c888.
//
// Solidity: event LogSwapout(address indexed account, address indexed bindaddr, uint256 amount)
func (_Anyerc20 *Anyerc20Filterer) FilterLogSwapout(opts *bind.FilterOpts, account []common.Address, bindaddr []common.Address) (*Anyerc20LogSwapoutIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var bindaddrRule []interface{}
	for _, bindaddrItem := range bindaddr {
		bindaddrRule = append(bindaddrRule, bindaddrItem)
	}

	logs, sub, err := _Anyerc20.contract.FilterLogs(opts, "LogSwapout", accountRule, bindaddrRule)
	if err != nil {
		return nil, err
	}
	return &Anyerc20LogSwapoutIterator{contract: _Anyerc20.contract, event: "LogSwapout", logs: logs, sub: sub}, nil
}

// WatchLogSwapout is a free log subscription operation binding the contract event 0x6b616089d04950dc06c45c6dd787d657980543f89651aec47924752c7d16c888.
//
// Solidity: event LogSwapout(address indexed account, address indexed bindaddr, uint256 amount)
func (_Anyerc20 *Anyerc20Filterer) WatchLogSwapout(opts *bind.WatchOpts, sink chan<- *Anyerc20LogSwapout, account []common.Address, bindaddr []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var bindaddrRule []interface{}
	for _, bindaddrItem := range bindaddr {
		bindaddrRule = append(bindaddrRule, bindaddrItem)
	}

	logs, sub, err := _Anyerc20.contract.WatchLogs(opts, "LogSwapout", accountRule, bindaddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Anyerc20LogSwapout)
				if err := _Anyerc20.contract.UnpackLog(event, "LogSwapout", log); err != nil {
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

// ParseLogSwapout is a log parse operation binding the contract event 0x6b616089d04950dc06c45c6dd787d657980543f89651aec47924752c7d16c888.
//
// Solidity: event LogSwapout(address indexed account, address indexed bindaddr, uint256 amount)
func (_Anyerc20 *Anyerc20Filterer) ParseLogSwapout(log types.Log) (*Anyerc20LogSwapout, error) {
	event := new(Anyerc20LogSwapout)
	if err := _Anyerc20.contract.UnpackLog(event, "LogSwapout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Anyerc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Anyerc20 contract.
type Anyerc20TransferIterator struct {
	Event *Anyerc20Transfer // Event containing the contract specifics and raw log

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
func (it *Anyerc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Anyerc20Transfer)
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
		it.Event = new(Anyerc20Transfer)
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
func (it *Anyerc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Anyerc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Anyerc20Transfer represents a Transfer event raised by the Anyerc20 contract.
type Anyerc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Anyerc20 *Anyerc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Anyerc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Anyerc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Anyerc20TransferIterator{contract: _Anyerc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Anyerc20 *Anyerc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Anyerc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Anyerc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Anyerc20Transfer)
				if err := _Anyerc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Anyerc20 *Anyerc20Filterer) ParseTransfer(log types.Log) (*Anyerc20Transfer, error) {
	event := new(Anyerc20Transfer)
	if err := _Anyerc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
