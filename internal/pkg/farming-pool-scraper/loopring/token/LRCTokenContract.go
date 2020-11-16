// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lrctokencontract

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

// LRCTokenContractABI is the input ABI used to generate the binding from.
const LRCTokenContractABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SNAPSHOT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LRCTokenContract is an auto generated Go binding around an Ethereum contract.
type LRCTokenContract struct {
	LRCTokenContractCaller     // Read-only binding to the contract
	LRCTokenContractTransactor // Write-only binding to the contract
	LRCTokenContractFilterer   // Log filterer for contract events
}

// LRCTokenContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type LRCTokenContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LRCTokenContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LRCTokenContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LRCTokenContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LRCTokenContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LRCTokenContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LRCTokenContractSession struct {
	Contract     *LRCTokenContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LRCTokenContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LRCTokenContractCallerSession struct {
	Contract *LRCTokenContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LRCTokenContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LRCTokenContractTransactorSession struct {
	Contract     *LRCTokenContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LRCTokenContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type LRCTokenContractRaw struct {
	Contract *LRCTokenContract // Generic contract binding to access the raw methods on
}

// LRCTokenContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LRCTokenContractCallerRaw struct {
	Contract *LRCTokenContractCaller // Generic read-only contract binding to access the raw methods on
}

// LRCTokenContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LRCTokenContractTransactorRaw struct {
	Contract *LRCTokenContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLRCTokenContract creates a new instance of LRCTokenContract, bound to a specific deployed contract.
func NewLRCTokenContract(address common.Address, backend bind.ContractBackend) (*LRCTokenContract, error) {
	contract, err := bindLRCTokenContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LRCTokenContract{LRCTokenContractCaller: LRCTokenContractCaller{contract: contract}, LRCTokenContractTransactor: LRCTokenContractTransactor{contract: contract}, LRCTokenContractFilterer: LRCTokenContractFilterer{contract: contract}}, nil
}

// NewLRCTokenContractCaller creates a new read-only instance of LRCTokenContract, bound to a specific deployed contract.
func NewLRCTokenContractCaller(address common.Address, caller bind.ContractCaller) (*LRCTokenContractCaller, error) {
	contract, err := bindLRCTokenContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LRCTokenContractCaller{contract: contract}, nil
}

// NewLRCTokenContractTransactor creates a new write-only instance of LRCTokenContract, bound to a specific deployed contract.
func NewLRCTokenContractTransactor(address common.Address, transactor bind.ContractTransactor) (*LRCTokenContractTransactor, error) {
	contract, err := bindLRCTokenContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LRCTokenContractTransactor{contract: contract}, nil
}

// NewLRCTokenContractFilterer creates a new log filterer instance of LRCTokenContract, bound to a specific deployed contract.
func NewLRCTokenContractFilterer(address common.Address, filterer bind.ContractFilterer) (*LRCTokenContractFilterer, error) {
	contract, err := bindLRCTokenContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LRCTokenContractFilterer{contract: contract}, nil
}

// bindLRCTokenContract binds a generic wrapper to an already deployed contract.
func bindLRCTokenContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LRCTokenContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LRCTokenContract *LRCTokenContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LRCTokenContract.Contract.LRCTokenContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LRCTokenContract *LRCTokenContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.LRCTokenContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LRCTokenContract *LRCTokenContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.LRCTokenContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LRCTokenContract *LRCTokenContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LRCTokenContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LRCTokenContract *LRCTokenContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LRCTokenContract *LRCTokenContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "DEFAULT_ADMIN_ROLE")
	return *ret0, err
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LRCTokenContract.Contract.DEFAULTADMINROLE(&_LRCTokenContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LRCTokenContract.Contract.DEFAULTADMINROLE(&_LRCTokenContract.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "DOMAIN_SEPARATOR")
	return *ret0, err
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _LRCTokenContract.Contract.DOMAINSEPARATOR(&_LRCTokenContract.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _LRCTokenContract.Contract.DOMAINSEPARATOR(&_LRCTokenContract.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "MINTER_ROLE")
	return *ret0, err
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractSession) MINTERROLE() ([32]byte, error) {
	return _LRCTokenContract.Contract.MINTERROLE(&_LRCTokenContract.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractCallerSession) MINTERROLE() ([32]byte, error) {
	return _LRCTokenContract.Contract.MINTERROLE(&_LRCTokenContract.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "PERMIT_TYPEHASH")
	return *ret0, err
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractSession) PERMITTYPEHASH() ([32]byte, error) {
	return _LRCTokenContract.Contract.PERMITTYPEHASH(&_LRCTokenContract.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _LRCTokenContract.Contract.PERMITTYPEHASH(&_LRCTokenContract.CallOpts)
}

// SNAPSHOTROLE is a free data retrieval call binding the contract method 0x7028e2cd.
//
// Solidity: function SNAPSHOT_ROLE() view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractCaller) SNAPSHOTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "SNAPSHOT_ROLE")
	return *ret0, err
}

// SNAPSHOTROLE is a free data retrieval call binding the contract method 0x7028e2cd.
//
// Solidity: function SNAPSHOT_ROLE() view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractSession) SNAPSHOTROLE() ([32]byte, error) {
	return _LRCTokenContract.Contract.SNAPSHOTROLE(&_LRCTokenContract.CallOpts)
}

// SNAPSHOTROLE is a free data retrieval call binding the contract method 0x7028e2cd.
//
// Solidity: function SNAPSHOT_ROLE() view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractCallerSession) SNAPSHOTROLE() ([32]byte, error) {
	return _LRCTokenContract.Contract.SNAPSHOTROLE(&_LRCTokenContract.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LRCTokenContract.Contract.Allowance(&_LRCTokenContract.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LRCTokenContract.Contract.Allowance(&_LRCTokenContract.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LRCTokenContract.Contract.BalanceOf(&_LRCTokenContract.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LRCTokenContract.Contract.BalanceOf(&_LRCTokenContract.CallOpts, account)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractCaller) BalanceOfAt(opts *bind.CallOpts, account common.Address, snapshotId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "balanceOfAt", account, snapshotId)
	return *ret0, err
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractSession) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _LRCTokenContract.Contract.BalanceOfAt(&_LRCTokenContract.CallOpts, account, snapshotId)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractCallerSession) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _LRCTokenContract.Contract.BalanceOfAt(&_LRCTokenContract.CallOpts, account, snapshotId)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LRCTokenContract *LRCTokenContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LRCTokenContract *LRCTokenContractSession) Decimals() (uint8, error) {
	return _LRCTokenContract.Contract.Decimals(&_LRCTokenContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LRCTokenContract *LRCTokenContractCallerSession) Decimals() (uint8, error) {
	return _LRCTokenContract.Contract.Decimals(&_LRCTokenContract.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "getRoleAdmin", role)
	return *ret0, err
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LRCTokenContract.Contract.GetRoleAdmin(&_LRCTokenContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LRCTokenContract *LRCTokenContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LRCTokenContract.Contract.GetRoleAdmin(&_LRCTokenContract.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_LRCTokenContract *LRCTokenContractCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "getRoleMember", role, index)
	return *ret0, err
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_LRCTokenContract *LRCTokenContractSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _LRCTokenContract.Contract.GetRoleMember(&_LRCTokenContract.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_LRCTokenContract *LRCTokenContractCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _LRCTokenContract.Contract.GetRoleMember(&_LRCTokenContract.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "getRoleMemberCount", role)
	return *ret0, err
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _LRCTokenContract.Contract.GetRoleMemberCount(&_LRCTokenContract.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _LRCTokenContract.Contract.GetRoleMemberCount(&_LRCTokenContract.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LRCTokenContract *LRCTokenContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "hasRole", role, account)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LRCTokenContract *LRCTokenContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LRCTokenContract.Contract.HasRole(&_LRCTokenContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LRCTokenContract *LRCTokenContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LRCTokenContract.Contract.HasRole(&_LRCTokenContract.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LRCTokenContract *LRCTokenContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LRCTokenContract *LRCTokenContractSession) Name() (string, error) {
	return _LRCTokenContract.Contract.Name(&_LRCTokenContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LRCTokenContract *LRCTokenContractCallerSession) Name() (string, error) {
	return _LRCTokenContract.Contract.Name(&_LRCTokenContract.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "nonces", arg0)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _LRCTokenContract.Contract.Nonces(&_LRCTokenContract.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _LRCTokenContract.Contract.Nonces(&_LRCTokenContract.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LRCTokenContract *LRCTokenContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LRCTokenContract *LRCTokenContractSession) Symbol() (string, error) {
	return _LRCTokenContract.Contract.Symbol(&_LRCTokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LRCTokenContract *LRCTokenContractCallerSession) Symbol() (string, error) {
	return _LRCTokenContract.Contract.Symbol(&_LRCTokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LRCTokenContract *LRCTokenContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LRCTokenContract *LRCTokenContractSession) TotalSupply() (*big.Int, error) {
	return _LRCTokenContract.Contract.TotalSupply(&_LRCTokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LRCTokenContract *LRCTokenContractCallerSession) TotalSupply() (*big.Int, error) {
	return _LRCTokenContract.Contract.TotalSupply(&_LRCTokenContract.CallOpts)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractCaller) TotalSupplyAt(opts *bind.CallOpts, snapshotId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "totalSupplyAt", snapshotId)
	return *ret0, err
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractSession) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _LRCTokenContract.Contract.TotalSupplyAt(&_LRCTokenContract.CallOpts, snapshotId)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_LRCTokenContract *LRCTokenContractCallerSession) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _LRCTokenContract.Contract.TotalSupplyAt(&_LRCTokenContract.CallOpts, snapshotId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LRCTokenContract *LRCTokenContractCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LRCTokenContract.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LRCTokenContract *LRCTokenContractSession) Version() (string, error) {
	return _LRCTokenContract.Contract.Version(&_LRCTokenContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LRCTokenContract *LRCTokenContractCallerSession) Version() (string, error) {
	return _LRCTokenContract.Contract.Version(&_LRCTokenContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LRCTokenContract *LRCTokenContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LRCTokenContract *LRCTokenContractSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.Approve(&_LRCTokenContract.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LRCTokenContract *LRCTokenContractTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.Approve(&_LRCTokenContract.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_LRCTokenContract *LRCTokenContractTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_LRCTokenContract *LRCTokenContractSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.Burn(&_LRCTokenContract.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_LRCTokenContract *LRCTokenContractTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.Burn(&_LRCTokenContract.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_LRCTokenContract *LRCTokenContractTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_LRCTokenContract *LRCTokenContractSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.BurnFrom(&_LRCTokenContract.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_LRCTokenContract *LRCTokenContractTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.BurnFrom(&_LRCTokenContract.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LRCTokenContract *LRCTokenContractTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LRCTokenContract *LRCTokenContractSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.DecreaseAllowance(&_LRCTokenContract.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LRCTokenContract *LRCTokenContractTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.DecreaseAllowance(&_LRCTokenContract.TransactOpts, spender, subtractedValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LRCTokenContract *LRCTokenContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LRCTokenContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LRCTokenContract *LRCTokenContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.GrantRole(&_LRCTokenContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LRCTokenContract *LRCTokenContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.GrantRole(&_LRCTokenContract.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LRCTokenContract *LRCTokenContractTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LRCTokenContract *LRCTokenContractSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.IncreaseAllowance(&_LRCTokenContract.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LRCTokenContract *LRCTokenContractTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.IncreaseAllowance(&_LRCTokenContract.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_LRCTokenContract *LRCTokenContractTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_LRCTokenContract *LRCTokenContractSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.Mint(&_LRCTokenContract.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_LRCTokenContract *LRCTokenContractTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.Mint(&_LRCTokenContract.TransactOpts, to, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_LRCTokenContract *LRCTokenContractTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LRCTokenContract.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_LRCTokenContract *LRCTokenContractSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.Permit(&_LRCTokenContract.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_LRCTokenContract *LRCTokenContractTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.Permit(&_LRCTokenContract.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LRCTokenContract *LRCTokenContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LRCTokenContract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LRCTokenContract *LRCTokenContractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.RenounceRole(&_LRCTokenContract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LRCTokenContract *LRCTokenContractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.RenounceRole(&_LRCTokenContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LRCTokenContract *LRCTokenContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LRCTokenContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LRCTokenContract *LRCTokenContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.RevokeRole(&_LRCTokenContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LRCTokenContract *LRCTokenContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.RevokeRole(&_LRCTokenContract.TransactOpts, role, account)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns()
func (_LRCTokenContract *LRCTokenContractTransactor) Snapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LRCTokenContract.contract.Transact(opts, "snapshot")
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns()
func (_LRCTokenContract *LRCTokenContractSession) Snapshot() (*types.Transaction, error) {
	return _LRCTokenContract.Contract.Snapshot(&_LRCTokenContract.TransactOpts)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns()
func (_LRCTokenContract *LRCTokenContractTransactorSession) Snapshot() (*types.Transaction, error) {
	return _LRCTokenContract.Contract.Snapshot(&_LRCTokenContract.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LRCTokenContract *LRCTokenContractTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LRCTokenContract *LRCTokenContractSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.Transfer(&_LRCTokenContract.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LRCTokenContract *LRCTokenContractTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.Transfer(&_LRCTokenContract.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LRCTokenContract *LRCTokenContractTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LRCTokenContract *LRCTokenContractSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.TransferFrom(&_LRCTokenContract.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LRCTokenContract *LRCTokenContractTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LRCTokenContract.Contract.TransferFrom(&_LRCTokenContract.TransactOpts, sender, recipient, amount)
}

// LRCTokenContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LRCTokenContract contract.
type LRCTokenContractApprovalIterator struct {
	Event *LRCTokenContractApproval // Event containing the contract specifics and raw log

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
func (it *LRCTokenContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCTokenContractApproval)
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
		it.Event = new(LRCTokenContractApproval)
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
func (it *LRCTokenContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCTokenContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCTokenContractApproval represents a Approval event raised by the LRCTokenContract contract.
type LRCTokenContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LRCTokenContract *LRCTokenContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LRCTokenContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LRCTokenContract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LRCTokenContractApprovalIterator{contract: _LRCTokenContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LRCTokenContract *LRCTokenContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LRCTokenContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LRCTokenContract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCTokenContractApproval)
				if err := _LRCTokenContract.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LRCTokenContract *LRCTokenContractFilterer) ParseApproval(log types.Log) (*LRCTokenContractApproval, error) {
	event := new(LRCTokenContractApproval)
	if err := _LRCTokenContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LRCTokenContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LRCTokenContract contract.
type LRCTokenContractRoleGrantedIterator struct {
	Event *LRCTokenContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *LRCTokenContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCTokenContractRoleGranted)
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
		it.Event = new(LRCTokenContractRoleGranted)
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
func (it *LRCTokenContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCTokenContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCTokenContractRoleGranted represents a RoleGranted event raised by the LRCTokenContract contract.
type LRCTokenContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LRCTokenContract *LRCTokenContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LRCTokenContractRoleGrantedIterator, error) {

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

	logs, sub, err := _LRCTokenContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LRCTokenContractRoleGrantedIterator{contract: _LRCTokenContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LRCTokenContract *LRCTokenContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LRCTokenContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LRCTokenContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCTokenContractRoleGranted)
				if err := _LRCTokenContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_LRCTokenContract *LRCTokenContractFilterer) ParseRoleGranted(log types.Log) (*LRCTokenContractRoleGranted, error) {
	event := new(LRCTokenContractRoleGranted)
	if err := _LRCTokenContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LRCTokenContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LRCTokenContract contract.
type LRCTokenContractRoleRevokedIterator struct {
	Event *LRCTokenContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LRCTokenContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCTokenContractRoleRevoked)
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
		it.Event = new(LRCTokenContractRoleRevoked)
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
func (it *LRCTokenContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCTokenContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCTokenContractRoleRevoked represents a RoleRevoked event raised by the LRCTokenContract contract.
type LRCTokenContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LRCTokenContract *LRCTokenContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LRCTokenContractRoleRevokedIterator, error) {

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

	logs, sub, err := _LRCTokenContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LRCTokenContractRoleRevokedIterator{contract: _LRCTokenContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LRCTokenContract *LRCTokenContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LRCTokenContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LRCTokenContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCTokenContractRoleRevoked)
				if err := _LRCTokenContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_LRCTokenContract *LRCTokenContractFilterer) ParseRoleRevoked(log types.Log) (*LRCTokenContractRoleRevoked, error) {
	event := new(LRCTokenContractRoleRevoked)
	if err := _LRCTokenContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LRCTokenContractSnapshotIterator is returned from FilterSnapshot and is used to iterate over the raw logs and unpacked data for Snapshot events raised by the LRCTokenContract contract.
type LRCTokenContractSnapshotIterator struct {
	Event *LRCTokenContractSnapshot // Event containing the contract specifics and raw log

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
func (it *LRCTokenContractSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCTokenContractSnapshot)
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
		it.Event = new(LRCTokenContractSnapshot)
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
func (it *LRCTokenContractSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCTokenContractSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCTokenContractSnapshot represents a Snapshot event raised by the LRCTokenContract contract.
type LRCTokenContractSnapshot struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSnapshot is a free log retrieval operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_LRCTokenContract *LRCTokenContractFilterer) FilterSnapshot(opts *bind.FilterOpts) (*LRCTokenContractSnapshotIterator, error) {

	logs, sub, err := _LRCTokenContract.contract.FilterLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return &LRCTokenContractSnapshotIterator{contract: _LRCTokenContract.contract, event: "Snapshot", logs: logs, sub: sub}, nil
}

// WatchSnapshot is a free log subscription operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_LRCTokenContract *LRCTokenContractFilterer) WatchSnapshot(opts *bind.WatchOpts, sink chan<- *LRCTokenContractSnapshot) (event.Subscription, error) {

	logs, sub, err := _LRCTokenContract.contract.WatchLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCTokenContractSnapshot)
				if err := _LRCTokenContract.contract.UnpackLog(event, "Snapshot", log); err != nil {
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

// ParseSnapshot is a log parse operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_LRCTokenContract *LRCTokenContractFilterer) ParseSnapshot(log types.Log) (*LRCTokenContractSnapshot, error) {
	event := new(LRCTokenContractSnapshot)
	if err := _LRCTokenContract.contract.UnpackLog(event, "Snapshot", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LRCTokenContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LRCTokenContract contract.
type LRCTokenContractTransferIterator struct {
	Event *LRCTokenContractTransfer // Event containing the contract specifics and raw log

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
func (it *LRCTokenContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LRCTokenContractTransfer)
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
		it.Event = new(LRCTokenContractTransfer)
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
func (it *LRCTokenContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LRCTokenContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LRCTokenContractTransfer represents a Transfer event raised by the LRCTokenContract contract.
type LRCTokenContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LRCTokenContract *LRCTokenContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LRCTokenContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LRCTokenContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LRCTokenContractTransferIterator{contract: _LRCTokenContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LRCTokenContract *LRCTokenContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LRCTokenContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LRCTokenContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LRCTokenContractTransfer)
				if err := _LRCTokenContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_LRCTokenContract *LRCTokenContractFilterer) ParseTransfer(log types.Log) (*LRCTokenContractTransfer, error) {
	event := new(LRCTokenContractTransfer)
	if err := _LRCTokenContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
