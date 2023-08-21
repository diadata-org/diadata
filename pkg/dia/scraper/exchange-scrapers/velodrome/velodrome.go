// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package velodrome

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

// ClonesMetaData contains all meta data concerning the Clones contract.
var ClonesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x607a6032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220aa659e64033eadb68f61388be31a1ea2eecadeeaf7e1bc07b0dc35e8805162e864736f6c637827302e382e32322d646576656c6f702e323032332e382e342b636f6d6d69742e35313137313235370058",
}

// ClonesABI is the input ABI used to generate the binding from.
// Deprecated: Use ClonesMetaData.ABI instead.
var ClonesABI = ClonesMetaData.ABI

// ClonesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ClonesMetaData.Bin instead.
var ClonesBin = ClonesMetaData.Bin

// DeployClones deploys a new Ethereum contract, binding an instance of Clones to it.
func DeployClones(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Clones, error) {
	parsed, err := ClonesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ClonesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Clones{ClonesCaller: ClonesCaller{contract: contract}, ClonesTransactor: ClonesTransactor{contract: contract}, ClonesFilterer: ClonesFilterer{contract: contract}}, nil
}

// Clones is an auto generated Go binding around an Ethereum contract.
type Clones struct {
	ClonesCaller     // Read-only binding to the contract
	ClonesTransactor // Write-only binding to the contract
	ClonesFilterer   // Log filterer for contract events
}

// ClonesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClonesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClonesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClonesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClonesSession struct {
	Contract     *Clones           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClonesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClonesCallerSession struct {
	Contract *ClonesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClonesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClonesTransactorSession struct {
	Contract     *ClonesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClonesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClonesRaw struct {
	Contract *Clones // Generic contract binding to access the raw methods on
}

// ClonesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClonesCallerRaw struct {
	Contract *ClonesCaller // Generic read-only contract binding to access the raw methods on
}

// ClonesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClonesTransactorRaw struct {
	Contract *ClonesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClones creates a new instance of Clones, bound to a specific deployed contract.
func NewClones(address common.Address, backend bind.ContractBackend) (*Clones, error) {
	contract, err := bindClones(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Clones{ClonesCaller: ClonesCaller{contract: contract}, ClonesTransactor: ClonesTransactor{contract: contract}, ClonesFilterer: ClonesFilterer{contract: contract}}, nil
}

// NewClonesCaller creates a new read-only instance of Clones, bound to a specific deployed contract.
func NewClonesCaller(address common.Address, caller bind.ContractCaller) (*ClonesCaller, error) {
	contract, err := bindClones(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClonesCaller{contract: contract}, nil
}

// NewClonesTransactor creates a new write-only instance of Clones, bound to a specific deployed contract.
func NewClonesTransactor(address common.Address, transactor bind.ContractTransactor) (*ClonesTransactor, error) {
	contract, err := bindClones(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClonesTransactor{contract: contract}, nil
}

// NewClonesFilterer creates a new log filterer instance of Clones, bound to a specific deployed contract.
func NewClonesFilterer(address common.Address, filterer bind.ContractFilterer) (*ClonesFilterer, error) {
	contract, err := bindClones(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClonesFilterer{contract: contract}, nil
}

// bindClones binds a generic wrapper to an already deployed contract.
func bindClones(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClonesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clones *ClonesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clones.Contract.ClonesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clones *ClonesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clones.Contract.ClonesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clones *ClonesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clones.Contract.ClonesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clones *ClonesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clones.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clones *ClonesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clones.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clones *ClonesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clones.Contract.contract.Transact(opts, method, params...)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
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

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
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

// IERC20MetadataMetaData contains all meta data concerning the IERC20Metadata contract.
var IERC20MetadataMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20MetadataABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetadataMetaData.ABI instead.
var IERC20MetadataABI = IERC20MetadataMetaData.ABI

// Deprecated: Use IERC20MetadataMetaData.Sigs instead.
// IERC20MetadataFuncSigs maps the 4-byte function signature to its string representation.
var IERC20MetadataFuncSigs = IERC20MetadataMetaData.Sigs

// IERC20Metadata is an auto generated Go binding around an Ethereum contract.
type IERC20Metadata struct {
	IERC20MetadataCaller     // Read-only binding to the contract
	IERC20MetadataTransactor // Write-only binding to the contract
	IERC20MetadataFilterer   // Log filterer for contract events
}

// IERC20MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20MetadataSession struct {
	Contract     *IERC20Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20MetadataCallerSession struct {
	Contract *IERC20MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC20MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20MetadataTransactorSession struct {
	Contract     *IERC20MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC20MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20MetadataRaw struct {
	Contract *IERC20Metadata // Generic contract binding to access the raw methods on
}

// IERC20MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20MetadataCallerRaw struct {
	Contract *IERC20MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactorRaw struct {
	Contract *IERC20MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Metadata creates a new instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20Metadata(address common.Address, backend bind.ContractBackend) (*IERC20Metadata, error) {
	contract, err := bindIERC20Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Metadata{IERC20MetadataCaller: IERC20MetadataCaller{contract: contract}, IERC20MetadataTransactor: IERC20MetadataTransactor{contract: contract}, IERC20MetadataFilterer: IERC20MetadataFilterer{contract: contract}}, nil
}

// NewIERC20MetadataCaller creates a new read-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataCaller(address common.Address, caller bind.ContractCaller) (*IERC20MetadataCaller, error) {
	contract, err := bindIERC20Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataCaller{contract: contract}, nil
}

// NewIERC20MetadataTransactor creates a new write-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MetadataTransactor, error) {
	contract, err := bindIERC20Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransactor{contract: contract}, nil
}

// NewIERC20MetadataFilterer creates a new log filterer instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MetadataFilterer, error) {
	contract, err := bindIERC20Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataFilterer{contract: contract}, nil
}

// bindIERC20Metadata binds a generic wrapper to an already deployed contract.
func bindIERC20Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.IERC20MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCallerSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, from, to, amount)
}

// IERC20MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20Metadata contract.
type IERC20MetadataApprovalIterator struct {
	Event *IERC20MetadataApproval // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataApproval)
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
		it.Event = new(IERC20MetadataApproval)
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
func (it *IERC20MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataApproval represents a Approval event raised by the IERC20Metadata contract.
type IERC20MetadataApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20MetadataApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataApprovalIterator{contract: _IERC20Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20MetadataApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataApproval)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20Metadata *IERC20MetadataFilterer) ParseApproval(log types.Log) (*IERC20MetadataApproval, error) {
	event := new(IERC20MetadataApproval)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20Metadata contract.
type IERC20MetadataTransferIterator struct {
	Event *IERC20MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataTransfer)
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
		it.Event = new(IERC20MetadataTransfer)
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
func (it *IERC20MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataTransfer represents a Transfer event raised by the IERC20Metadata contract.
type IERC20MetadataTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20MetadataTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransferIterator{contract: _IERC20Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20MetadataTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataTransfer)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20Metadata *IERC20MetadataFilterer) ParseTransfer(log types.Log) (*IERC20MetadataTransfer, error) {
	event := new(IERC20MetadataTransfer)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolMetaData contains all meta data concerning the IPool contract.
var IPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BelowMinimumK\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositsNotEqual\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FactoryAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientInputAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidityBurned\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidityMinted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientOutputAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IsPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"K\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEmergencyCouncil\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Fees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_reserve0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserve1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestampLast\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_stable\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadata\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dec0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dec1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r1\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"st\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"t0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t1\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"89afcb44": "burn(address)",
		"d294f093": "claimFees()",
		"f140a35a": "getAmountOut(uint256,address)",
		"0902f1ac": "getReserves()",
		"e4bbb5a8": "initialize(address,address,bool)",
		"392f37e9": "metadata()",
		"6a627842": "mint(address)",
		"bc25cf77": "skim(address)",
		"22be3de1": "stable()",
		"022c0d9f": "swap(uint256,uint256,address,bytes)",
		"0dfe1681": "token0()",
		"d21220a7": "token1()",
		"9d63848a": "tokens()",
	},
}

// IPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use IPoolMetaData.ABI instead.
var IPoolABI = IPoolMetaData.ABI

// Deprecated: Use IPoolMetaData.Sigs instead.
// IPoolFuncSigs maps the 4-byte function signature to its string representation.
var IPoolFuncSigs = IPoolMetaData.Sigs

// IPool is an auto generated Go binding around an Ethereum contract.
type IPool struct {
	IPoolCaller     // Read-only binding to the contract
	IPoolTransactor // Write-only binding to the contract
	IPoolFilterer   // Log filterer for contract events
}

// IPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPoolSession struct {
	Contract     *IPool            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPoolCallerSession struct {
	Contract *IPoolCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPoolTransactorSession struct {
	Contract     *IPoolTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPoolRaw struct {
	Contract *IPool // Generic contract binding to access the raw methods on
}

// IPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPoolCallerRaw struct {
	Contract *IPoolCaller // Generic read-only contract binding to access the raw methods on
}

// IPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPoolTransactorRaw struct {
	Contract *IPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPool creates a new instance of IPool, bound to a specific deployed contract.
func NewIPool(address common.Address, backend bind.ContractBackend) (*IPool, error) {
	contract, err := bindIPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPool{IPoolCaller: IPoolCaller{contract: contract}, IPoolTransactor: IPoolTransactor{contract: contract}, IPoolFilterer: IPoolFilterer{contract: contract}}, nil
}

// NewIPoolCaller creates a new read-only instance of IPool, bound to a specific deployed contract.
func NewIPoolCaller(address common.Address, caller bind.ContractCaller) (*IPoolCaller, error) {
	contract, err := bindIPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolCaller{contract: contract}, nil
}

// NewIPoolTransactor creates a new write-only instance of IPool, bound to a specific deployed contract.
func NewIPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IPoolTransactor, error) {
	contract, err := bindIPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolTransactor{contract: contract}, nil
}

// NewIPoolFilterer creates a new log filterer instance of IPool, bound to a specific deployed contract.
func NewIPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IPoolFilterer, error) {
	contract, err := bindIPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPoolFilterer{contract: contract}, nil
}

// bindIPool binds a generic wrapper to an already deployed contract.
func bindIPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPool *IPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPool.Contract.IPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPool *IPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPool.Contract.IPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPool *IPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPool.Contract.IPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPool *IPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPool *IPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPool *IPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPool.Contract.contract.Transact(opts, method, params...)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 , address ) view returns(uint256)
func (_IPool *IPoolCaller) GetAmountOut(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPool.contract.Call(opts, &out, "getAmountOut", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 , address ) view returns(uint256)
func (_IPool *IPoolSession) GetAmountOut(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _IPool.Contract.GetAmountOut(&_IPool.CallOpts, arg0, arg1)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 , address ) view returns(uint256)
func (_IPool *IPoolCallerSession) GetAmountOut(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _IPool.Contract.GetAmountOut(&_IPool.CallOpts, arg0, arg1)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1, uint256 _blockTimestampLast)
func (_IPool *IPoolCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast *big.Int
}, error) {
	var out []interface{}
	err := _IPool.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1, uint256 _blockTimestampLast)
func (_IPool *IPoolSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast *big.Int
}, error) {
	return _IPool.Contract.GetReserves(&_IPool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1, uint256 _blockTimestampLast)
func (_IPool *IPoolCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast *big.Int
}, error) {
	return _IPool.Contract.GetReserves(&_IPool.CallOpts)
}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(uint256 dec0, uint256 dec1, uint256 r0, uint256 r1, bool st, address t0, address t1)
func (_IPool *IPoolCaller) Metadata(opts *bind.CallOpts) (struct {
	Dec0 *big.Int
	Dec1 *big.Int
	R0   *big.Int
	R1   *big.Int
	St   bool
	T0   common.Address
	T1   common.Address
}, error) {
	var out []interface{}
	err := _IPool.contract.Call(opts, &out, "metadata")

	outstruct := new(struct {
		Dec0 *big.Int
		Dec1 *big.Int
		R0   *big.Int
		R1   *big.Int
		St   bool
		T0   common.Address
		T1   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dec0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Dec1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.R0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.R1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.St = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.T0 = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.T1 = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(uint256 dec0, uint256 dec1, uint256 r0, uint256 r1, bool st, address t0, address t1)
func (_IPool *IPoolSession) Metadata() (struct {
	Dec0 *big.Int
	Dec1 *big.Int
	R0   *big.Int
	R1   *big.Int
	St   bool
	T0   common.Address
	T1   common.Address
}, error) {
	return _IPool.Contract.Metadata(&_IPool.CallOpts)
}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(uint256 dec0, uint256 dec1, uint256 r0, uint256 r1, bool st, address t0, address t1)
func (_IPool *IPoolCallerSession) Metadata() (struct {
	Dec0 *big.Int
	Dec1 *big.Int
	R0   *big.Int
	R1   *big.Int
	St   bool
	T0   common.Address
	T1   common.Address
}, error) {
	return _IPool.Contract.Metadata(&_IPool.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(bool)
func (_IPool *IPoolCaller) Stable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IPool.contract.Call(opts, &out, "stable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(bool)
func (_IPool *IPoolSession) Stable() (bool, error) {
	return _IPool.Contract.Stable(&_IPool.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(bool)
func (_IPool *IPoolCallerSession) Stable() (bool, error) {
	return _IPool.Contract.Stable(&_IPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IPool *IPoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IPool *IPoolSession) Token0() (common.Address, error) {
	return _IPool.Contract.Token0(&_IPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IPool *IPoolCallerSession) Token0() (common.Address, error) {
	return _IPool.Contract.Token0(&_IPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IPool *IPoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IPool *IPoolSession) Token1() (common.Address, error) {
	return _IPool.Contract.Token1(&_IPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IPool *IPoolCallerSession) Token1() (common.Address, error) {
	return _IPool.Contract.Token1(&_IPool.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() view returns(address, address)
func (_IPool *IPoolCaller) Tokens(opts *bind.CallOpts) (common.Address, common.Address, error) {
	var out []interface{}
	err := _IPool.contract.Call(opts, &out, "tokens")

	if err != nil {
		return *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() view returns(address, address)
func (_IPool *IPoolSession) Tokens() (common.Address, common.Address, error) {
	return _IPool.Contract.Tokens(&_IPool.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() view returns(address, address)
func (_IPool *IPoolCallerSession) Tokens() (common.Address, common.Address, error) {
	return _IPool.Contract.Tokens(&_IPool.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_IPool *IPoolTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_IPool *IPoolSession) Burn(to common.Address) (*types.Transaction, error) {
	return _IPool.Contract.Burn(&_IPool.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_IPool *IPoolTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _IPool.Contract.Burn(&_IPool.TransactOpts, to)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns(uint256, uint256)
func (_IPool *IPoolTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns(uint256, uint256)
func (_IPool *IPoolSession) ClaimFees() (*types.Transaction, error) {
	return _IPool.Contract.ClaimFees(&_IPool.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns(uint256, uint256)
func (_IPool *IPoolTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _IPool.Contract.ClaimFees(&_IPool.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4bbb5a8.
//
// Solidity: function initialize(address _token0, address _token1, bool _stable) returns()
func (_IPool *IPoolTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address, _stable bool) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "initialize", _token0, _token1, _stable)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4bbb5a8.
//
// Solidity: function initialize(address _token0, address _token1, bool _stable) returns()
func (_IPool *IPoolSession) Initialize(_token0 common.Address, _token1 common.Address, _stable bool) (*types.Transaction, error) {
	return _IPool.Contract.Initialize(&_IPool.TransactOpts, _token0, _token1, _stable)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4bbb5a8.
//
// Solidity: function initialize(address _token0, address _token1, bool _stable) returns()
func (_IPool *IPoolTransactorSession) Initialize(_token0 common.Address, _token1 common.Address, _stable bool) (*types.Transaction, error) {
	return _IPool.Contract.Initialize(&_IPool.TransactOpts, _token0, _token1, _stable)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_IPool *IPoolTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_IPool *IPoolSession) Mint(to common.Address) (*types.Transaction, error) {
	return _IPool.Contract.Mint(&_IPool.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_IPool *IPoolTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _IPool.Contract.Mint(&_IPool.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_IPool *IPoolTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_IPool *IPoolSession) Skim(to common.Address) (*types.Transaction, error) {
	return _IPool.Contract.Skim(&_IPool.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_IPool *IPoolTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _IPool.Contract.Skim(&_IPool.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_IPool *IPoolTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_IPool *IPoolSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IPool.Contract.Swap(&_IPool.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_IPool *IPoolTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IPool.Contract.Swap(&_IPool.TransactOpts, amount0Out, amount1Out, to, data)
}

// IPoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the IPool contract.
type IPoolBurnIterator struct {
	Event *IPoolBurn // Event containing the contract specifics and raw log

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
func (it *IPoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolBurn)
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
		it.Event = new(IPoolBurn)
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
func (it *IPoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolBurn represents a Burn event raised by the IPool contract.
type IPoolBurn struct {
	Sender  common.Address
	To      common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x5d624aa9c148153ab3446c1b154f660ee7701e549fe9b62dab7171b1c80e6fa2.
//
// Solidity: event Burn(address indexed sender, address indexed to, uint256 amount0, uint256 amount1)
func (_IPool *IPoolFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*IPoolBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPool.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IPoolBurnIterator{contract: _IPool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x5d624aa9c148153ab3446c1b154f660ee7701e549fe9b62dab7171b1c80e6fa2.
//
// Solidity: event Burn(address indexed sender, address indexed to, uint256 amount0, uint256 amount1)
func (_IPool *IPoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *IPoolBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPool.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolBurn)
				if err := _IPool.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x5d624aa9c148153ab3446c1b154f660ee7701e549fe9b62dab7171b1c80e6fa2.
//
// Solidity: event Burn(address indexed sender, address indexed to, uint256 amount0, uint256 amount1)
func (_IPool *IPoolFilterer) ParseBurn(log types.Log) (*IPoolBurn, error) {
	event := new(IPoolBurn)
	if err := _IPool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the IPool contract.
type IPoolClaimIterator struct {
	Event *IPoolClaim // Event containing the contract specifics and raw log

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
func (it *IPoolClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolClaim)
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
		it.Event = new(IPoolClaim)
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
func (it *IPoolClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolClaim represents a Claim event raised by the IPool contract.
type IPoolClaim struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x865ca08d59f5cb456e85cd2f7ef63664ea4f73327414e9d8152c4158b0e94645.
//
// Solidity: event Claim(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1)
func (_IPool *IPoolFilterer) FilterClaim(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*IPoolClaimIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IPool.contract.FilterLogs(opts, "Claim", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &IPoolClaimIterator{contract: _IPool.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x865ca08d59f5cb456e85cd2f7ef63664ea4f73327414e9d8152c4158b0e94645.
//
// Solidity: event Claim(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1)
func (_IPool *IPoolFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *IPoolClaim, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IPool.contract.WatchLogs(opts, "Claim", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolClaim)
				if err := _IPool.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x865ca08d59f5cb456e85cd2f7ef63664ea4f73327414e9d8152c4158b0e94645.
//
// Solidity: event Claim(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1)
func (_IPool *IPoolFilterer) ParseClaim(log types.Log) (*IPoolClaim, error) {
	event := new(IPoolClaim)
	if err := _IPool.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolFeesIterator is returned from FilterFees and is used to iterate over the raw logs and unpacked data for Fees events raised by the IPool contract.
type IPoolFeesIterator struct {
	Event *IPoolFees // Event containing the contract specifics and raw log

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
func (it *IPoolFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolFees)
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
		it.Event = new(IPoolFees)
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
func (it *IPoolFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolFees represents a Fees event raised by the IPool contract.
type IPoolFees struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFees is a free log retrieval operation binding the contract event 0x112c256902bf554b6ed882d2936687aaeb4225e8cd5b51303c90ca6cf43a8602.
//
// Solidity: event Fees(address indexed sender, uint256 amount0, uint256 amount1)
func (_IPool *IPoolFilterer) FilterFees(opts *bind.FilterOpts, sender []common.Address) (*IPoolFeesIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IPool.contract.FilterLogs(opts, "Fees", senderRule)
	if err != nil {
		return nil, err
	}
	return &IPoolFeesIterator{contract: _IPool.contract, event: "Fees", logs: logs, sub: sub}, nil
}

// WatchFees is a free log subscription operation binding the contract event 0x112c256902bf554b6ed882d2936687aaeb4225e8cd5b51303c90ca6cf43a8602.
//
// Solidity: event Fees(address indexed sender, uint256 amount0, uint256 amount1)
func (_IPool *IPoolFilterer) WatchFees(opts *bind.WatchOpts, sink chan<- *IPoolFees, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IPool.contract.WatchLogs(opts, "Fees", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolFees)
				if err := _IPool.contract.UnpackLog(event, "Fees", log); err != nil {
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

// ParseFees is a log parse operation binding the contract event 0x112c256902bf554b6ed882d2936687aaeb4225e8cd5b51303c90ca6cf43a8602.
//
// Solidity: event Fees(address indexed sender, uint256 amount0, uint256 amount1)
func (_IPool *IPoolFilterer) ParseFees(log types.Log) (*IPoolFees, error) {
	event := new(IPoolFees)
	if err := _IPool.contract.UnpackLog(event, "Fees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IPool contract.
type IPoolMintIterator struct {
	Event *IPoolMint // Event containing the contract specifics and raw log

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
func (it *IPoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolMint)
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
		it.Event = new(IPoolMint)
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
func (it *IPoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolMint represents a Mint event raised by the IPool contract.
type IPoolMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_IPool *IPoolFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*IPoolMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IPool.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &IPoolMintIterator{contract: _IPool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_IPool *IPoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IPoolMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IPool.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolMint)
				if err := _IPool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_IPool *IPoolFilterer) ParseMint(log types.Log) (*IPoolMint, error) {
	event := new(IPoolMint)
	if err := _IPool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the IPool contract.
type IPoolSwapIterator struct {
	Event *IPoolSwap // Event containing the contract specifics and raw log

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
func (it *IPoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolSwap)
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
		it.Event = new(IPoolSwap)
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
func (it *IPoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolSwap represents a Swap event raised by the IPool contract.
type IPoolSwap struct {
	Sender     common.Address
	To         common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xb3e2773606abfd36b5bd91394b3a54d1398336c65005baf7bf7a05efeffaf75b.
//
// Solidity: event Swap(address indexed sender, address indexed to, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out)
func (_IPool *IPoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*IPoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPool.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IPoolSwapIterator{contract: _IPool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xb3e2773606abfd36b5bd91394b3a54d1398336c65005baf7bf7a05efeffaf75b.
//
// Solidity: event Swap(address indexed sender, address indexed to, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out)
func (_IPool *IPoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *IPoolSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPool.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolSwap)
				if err := _IPool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xb3e2773606abfd36b5bd91394b3a54d1398336c65005baf7bf7a05efeffaf75b.
//
// Solidity: event Swap(address indexed sender, address indexed to, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out)
func (_IPool *IPoolFilterer) ParseSwap(log types.Log) (*IPoolSwap, error) {
	event := new(IPoolSwap)
	if err := _IPool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the IPool contract.
type IPoolSyncIterator struct {
	Event *IPoolSync // Event containing the contract specifics and raw log

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
func (it *IPoolSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolSync)
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
		it.Event = new(IPoolSync)
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
func (it *IPoolSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolSync represents a Sync event raised by the IPool contract.
type IPoolSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_IPool *IPoolFilterer) FilterSync(opts *bind.FilterOpts) (*IPoolSyncIterator, error) {

	logs, sub, err := _IPool.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &IPoolSyncIterator{contract: _IPool.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_IPool *IPoolFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *IPoolSync) (event.Subscription, error) {

	logs, sub, err := _IPool.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolSync)
				if err := _IPool.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_IPool *IPoolFilterer) ParseSync(log types.Log) (*IPoolSync, error) {
	event := new(IPoolSync)
	if err := _IPool.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolFactoryMetaData contains all meta data concerning the IPoolFactory contract.
var IPoolFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"FeeInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotFeeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotPauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSinkConverter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotVoter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFee\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"SetCustomFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"}],\"name\":\"SetFeeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"SetPauseState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"SetPauser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"SetVoter\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allPoolsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_stable\",\"type\":\"bool\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPair\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setCustomFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeManager\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"setPauseState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"setPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sinkConvert\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_velo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_veloV2\",\"type\":\"address\"}],\"name\":\"setSinkConverter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"setVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sinkConverter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"velo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"veloV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"efde4e64": "allPoolsLength()",
		"82dfdce4": "createPair(address,address,bool)",
		"36bf95a0": "createPool(address,address,bool)",
		"a1671295": "createPool(address,address,uint24)",
		"cc56b2c5": "getFee(address,bool)",
		"6801cc30": "getPair(address,address,bool)",
		"79bc57d5": "getPool(address,address,bool)",
		"1698ee82": "getPool(address,address,uint24)",
		"5c60da1b": "implementation()",
		"e5e31b13": "isPair(address)",
		"b187bd26": "isPaused()",
		"5b16ebb7": "isPool(address)",
		"d49466a8": "setCustomFee(address,uint256)",
		"e1f76b44": "setFee(bool,uint256)",
		"472d35b9": "setFeeManager(address)",
		"cdb88ad1": "setPauseState(bool)",
		"2d88af4a": "setPauser(address)",
		"37068e7b": "setSinkConverter(address,address,address)",
		"4bc2a657": "setVoter(address)",
		"8e39ee16": "sinkConverter()",
		"8c7c53ce": "velo()",
		"c6751c09": "veloV2()",
		"46c96aac": "voter()",
	},
}

// IPoolFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IPoolFactoryMetaData.ABI instead.
var IPoolFactoryABI = IPoolFactoryMetaData.ABI

// Deprecated: Use IPoolFactoryMetaData.Sigs instead.
// IPoolFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IPoolFactoryFuncSigs = IPoolFactoryMetaData.Sigs

// IPoolFactory is an auto generated Go binding around an Ethereum contract.
type IPoolFactory struct {
	IPoolFactoryCaller     // Read-only binding to the contract
	IPoolFactoryTransactor // Write-only binding to the contract
	IPoolFactoryFilterer   // Log filterer for contract events
}

// IPoolFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPoolFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPoolFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPoolFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPoolFactorySession struct {
	Contract     *IPoolFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPoolFactoryCallerSession struct {
	Contract *IPoolFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IPoolFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPoolFactoryTransactorSession struct {
	Contract     *IPoolFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IPoolFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPoolFactoryRaw struct {
	Contract *IPoolFactory // Generic contract binding to access the raw methods on
}

// IPoolFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPoolFactoryCallerRaw struct {
	Contract *IPoolFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IPoolFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPoolFactoryTransactorRaw struct {
	Contract *IPoolFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPoolFactory creates a new instance of IPoolFactory, bound to a specific deployed contract.
func NewIPoolFactory(address common.Address, backend bind.ContractBackend) (*IPoolFactory, error) {
	contract, err := bindIPoolFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPoolFactory{IPoolFactoryCaller: IPoolFactoryCaller{contract: contract}, IPoolFactoryTransactor: IPoolFactoryTransactor{contract: contract}, IPoolFactoryFilterer: IPoolFactoryFilterer{contract: contract}}, nil
}

// NewIPoolFactoryCaller creates a new read-only instance of IPoolFactory, bound to a specific deployed contract.
func NewIPoolFactoryCaller(address common.Address, caller bind.ContractCaller) (*IPoolFactoryCaller, error) {
	contract, err := bindIPoolFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolFactoryCaller{contract: contract}, nil
}

// NewIPoolFactoryTransactor creates a new write-only instance of IPoolFactory, bound to a specific deployed contract.
func NewIPoolFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IPoolFactoryTransactor, error) {
	contract, err := bindIPoolFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolFactoryTransactor{contract: contract}, nil
}

// NewIPoolFactoryFilterer creates a new log filterer instance of IPoolFactory, bound to a specific deployed contract.
func NewIPoolFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IPoolFactoryFilterer, error) {
	contract, err := bindIPoolFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPoolFactoryFilterer{contract: contract}, nil
}

// bindIPoolFactory binds a generic wrapper to an already deployed contract.
func bindIPoolFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPoolFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoolFactory *IPoolFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoolFactory.Contract.IPoolFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoolFactory *IPoolFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoolFactory.Contract.IPoolFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoolFactory *IPoolFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoolFactory.Contract.IPoolFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoolFactory *IPoolFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoolFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoolFactory *IPoolFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoolFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoolFactory *IPoolFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoolFactory.Contract.contract.Transact(opts, method, params...)
}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (_IPoolFactory *IPoolFactoryCaller) AllPoolsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPoolFactory.contract.Call(opts, &out, "allPoolsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (_IPoolFactory *IPoolFactorySession) AllPoolsLength() (*big.Int, error) {
	return _IPoolFactory.Contract.AllPoolsLength(&_IPoolFactory.CallOpts)
}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (_IPoolFactory *IPoolFactoryCallerSession) AllPoolsLength() (*big.Int, error) {
	return _IPoolFactory.Contract.AllPoolsLength(&_IPoolFactory.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xcc56b2c5.
//
// Solidity: function getFee(address _pool, bool _stable) view returns(uint256)
func (_IPoolFactory *IPoolFactoryCaller) GetFee(opts *bind.CallOpts, _pool common.Address, _stable bool) (*big.Int, error) {
	var out []interface{}
	err := _IPoolFactory.contract.Call(opts, &out, "getFee", _pool, _stable)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xcc56b2c5.
//
// Solidity: function getFee(address _pool, bool _stable) view returns(uint256)
func (_IPoolFactory *IPoolFactorySession) GetFee(_pool common.Address, _stable bool) (*big.Int, error) {
	return _IPoolFactory.Contract.GetFee(&_IPoolFactory.CallOpts, _pool, _stable)
}

// GetFee is a free data retrieval call binding the contract method 0xcc56b2c5.
//
// Solidity: function getFee(address _pool, bool _stable) view returns(uint256)
func (_IPoolFactory *IPoolFactoryCallerSession) GetFee(_pool common.Address, _stable bool) (*big.Int, error) {
	return _IPoolFactory.Contract.GetFee(&_IPoolFactory.CallOpts, _pool, _stable)
}

// GetPair is a free data retrieval call binding the contract method 0x6801cc30.
//
// Solidity: function getPair(address tokenA, address tokenB, bool stable) view returns(address)
func (_IPoolFactory *IPoolFactoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	var out []interface{}
	err := _IPoolFactory.contract.Call(opts, &out, "getPair", tokenA, tokenB, stable)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0x6801cc30.
//
// Solidity: function getPair(address tokenA, address tokenB, bool stable) view returns(address)
func (_IPoolFactory *IPoolFactorySession) GetPair(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _IPoolFactory.Contract.GetPair(&_IPoolFactory.CallOpts, tokenA, tokenB, stable)
}

// GetPair is a free data retrieval call binding the contract method 0x6801cc30.
//
// Solidity: function getPair(address tokenA, address tokenB, bool stable) view returns(address)
func (_IPoolFactory *IPoolFactoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _IPoolFactory.Contract.GetPair(&_IPoolFactory.CallOpts, tokenA, tokenB, stable)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address tokenA, address tokenB, uint24 fee) view returns(address)
func (_IPoolFactory *IPoolFactoryCaller) GetPool(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, fee *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IPoolFactory.contract.Call(opts, &out, "getPool", tokenA, tokenB, fee)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address tokenA, address tokenB, uint24 fee) view returns(address)
func (_IPoolFactory *IPoolFactorySession) GetPool(tokenA common.Address, tokenB common.Address, fee *big.Int) (common.Address, error) {
	return _IPoolFactory.Contract.GetPool(&_IPoolFactory.CallOpts, tokenA, tokenB, fee)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address tokenA, address tokenB, uint24 fee) view returns(address)
func (_IPoolFactory *IPoolFactoryCallerSession) GetPool(tokenA common.Address, tokenB common.Address, fee *big.Int) (common.Address, error) {
	return _IPoolFactory.Contract.GetPool(&_IPoolFactory.CallOpts, tokenA, tokenB, fee)
}

// GetPool0 is a free data retrieval call binding the contract method 0x79bc57d5.
//
// Solidity: function getPool(address tokenA, address tokenB, bool stable) view returns(address)
func (_IPoolFactory *IPoolFactoryCaller) GetPool0(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	var out []interface{}
	err := _IPoolFactory.contract.Call(opts, &out, "getPool0", tokenA, tokenB, stable)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool0 is a free data retrieval call binding the contract method 0x79bc57d5.
//
// Solidity: function getPool(address tokenA, address tokenB, bool stable) view returns(address)
func (_IPoolFactory *IPoolFactorySession) GetPool0(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _IPoolFactory.Contract.GetPool0(&_IPoolFactory.CallOpts, tokenA, tokenB, stable)
}

// GetPool0 is a free data retrieval call binding the contract method 0x79bc57d5.
//
// Solidity: function getPool(address tokenA, address tokenB, bool stable) view returns(address)
func (_IPoolFactory *IPoolFactoryCallerSession) GetPool0(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _IPoolFactory.Contract.GetPool0(&_IPoolFactory.CallOpts, tokenA, tokenB, stable)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IPoolFactory *IPoolFactoryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPoolFactory.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IPoolFactory *IPoolFactorySession) Implementation() (common.Address, error) {
	return _IPoolFactory.Contract.Implementation(&_IPoolFactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IPoolFactory *IPoolFactoryCallerSession) Implementation() (common.Address, error) {
	return _IPoolFactory.Contract.Implementation(&_IPoolFactory.CallOpts)
}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address pool) view returns(bool)
func (_IPoolFactory *IPoolFactoryCaller) IsPair(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _IPoolFactory.contract.Call(opts, &out, "isPair", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address pool) view returns(bool)
func (_IPoolFactory *IPoolFactorySession) IsPair(pool common.Address) (bool, error) {
	return _IPoolFactory.Contract.IsPair(&_IPoolFactory.CallOpts, pool)
}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address pool) view returns(bool)
func (_IPoolFactory *IPoolFactoryCallerSession) IsPair(pool common.Address) (bool, error) {
	return _IPoolFactory.Contract.IsPair(&_IPoolFactory.CallOpts, pool)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_IPoolFactory *IPoolFactoryCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IPoolFactory.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_IPoolFactory *IPoolFactorySession) IsPaused() (bool, error) {
	return _IPoolFactory.Contract.IsPaused(&_IPoolFactory.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_IPoolFactory *IPoolFactoryCallerSession) IsPaused() (bool, error) {
	return _IPoolFactory.Contract.IsPaused(&_IPoolFactory.CallOpts)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address pool) view returns(bool)
func (_IPoolFactory *IPoolFactoryCaller) IsPool(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _IPoolFactory.contract.Call(opts, &out, "isPool", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address pool) view returns(bool)
func (_IPoolFactory *IPoolFactorySession) IsPool(pool common.Address) (bool, error) {
	return _IPoolFactory.Contract.IsPool(&_IPoolFactory.CallOpts, pool)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address pool) view returns(bool)
func (_IPoolFactory *IPoolFactoryCallerSession) IsPool(pool common.Address) (bool, error) {
	return _IPoolFactory.Contract.IsPool(&_IPoolFactory.CallOpts, pool)
}

// SinkConverter is a free data retrieval call binding the contract method 0x8e39ee16.
//
// Solidity: function sinkConverter() view returns(address)
func (_IPoolFactory *IPoolFactoryCaller) SinkConverter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPoolFactory.contract.Call(opts, &out, "sinkConverter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SinkConverter is a free data retrieval call binding the contract method 0x8e39ee16.
//
// Solidity: function sinkConverter() view returns(address)
func (_IPoolFactory *IPoolFactorySession) SinkConverter() (common.Address, error) {
	return _IPoolFactory.Contract.SinkConverter(&_IPoolFactory.CallOpts)
}

// SinkConverter is a free data retrieval call binding the contract method 0x8e39ee16.
//
// Solidity: function sinkConverter() view returns(address)
func (_IPoolFactory *IPoolFactoryCallerSession) SinkConverter() (common.Address, error) {
	return _IPoolFactory.Contract.SinkConverter(&_IPoolFactory.CallOpts)
}

// Velo is a free data retrieval call binding the contract method 0x8c7c53ce.
//
// Solidity: function velo() view returns(address)
func (_IPoolFactory *IPoolFactoryCaller) Velo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPoolFactory.contract.Call(opts, &out, "velo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Velo is a free data retrieval call binding the contract method 0x8c7c53ce.
//
// Solidity: function velo() view returns(address)
func (_IPoolFactory *IPoolFactorySession) Velo() (common.Address, error) {
	return _IPoolFactory.Contract.Velo(&_IPoolFactory.CallOpts)
}

// Velo is a free data retrieval call binding the contract method 0x8c7c53ce.
//
// Solidity: function velo() view returns(address)
func (_IPoolFactory *IPoolFactoryCallerSession) Velo() (common.Address, error) {
	return _IPoolFactory.Contract.Velo(&_IPoolFactory.CallOpts)
}

// VeloV2 is a free data retrieval call binding the contract method 0xc6751c09.
//
// Solidity: function veloV2() view returns(address)
func (_IPoolFactory *IPoolFactoryCaller) VeloV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPoolFactory.contract.Call(opts, &out, "veloV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VeloV2 is a free data retrieval call binding the contract method 0xc6751c09.
//
// Solidity: function veloV2() view returns(address)
func (_IPoolFactory *IPoolFactorySession) VeloV2() (common.Address, error) {
	return _IPoolFactory.Contract.VeloV2(&_IPoolFactory.CallOpts)
}

// VeloV2 is a free data retrieval call binding the contract method 0xc6751c09.
//
// Solidity: function veloV2() view returns(address)
func (_IPoolFactory *IPoolFactoryCallerSession) VeloV2() (common.Address, error) {
	return _IPoolFactory.Contract.VeloV2(&_IPoolFactory.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_IPoolFactory *IPoolFactoryCaller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPoolFactory.contract.Call(opts, &out, "voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_IPoolFactory *IPoolFactorySession) Voter() (common.Address, error) {
	return _IPoolFactory.Contract.Voter(&_IPoolFactory.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_IPoolFactory *IPoolFactoryCallerSession) Voter() (common.Address, error) {
	return _IPoolFactory.Contract.Voter(&_IPoolFactory.CallOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0x82dfdce4.
//
// Solidity: function createPair(address tokenA, address tokenB, bool stable) returns(address pool)
func (_IPoolFactory *IPoolFactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _IPoolFactory.contract.Transact(opts, "createPair", tokenA, tokenB, stable)
}

// CreatePair is a paid mutator transaction binding the contract method 0x82dfdce4.
//
// Solidity: function createPair(address tokenA, address tokenB, bool stable) returns(address pool)
func (_IPoolFactory *IPoolFactorySession) CreatePair(tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _IPoolFactory.Contract.CreatePair(&_IPoolFactory.TransactOpts, tokenA, tokenB, stable)
}

// CreatePair is a paid mutator transaction binding the contract method 0x82dfdce4.
//
// Solidity: function createPair(address tokenA, address tokenB, bool stable) returns(address pool)
func (_IPoolFactory *IPoolFactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _IPoolFactory.Contract.CreatePair(&_IPoolFactory.TransactOpts, tokenA, tokenB, stable)
}

// CreatePool is a paid mutator transaction binding the contract method 0x36bf95a0.
//
// Solidity: function createPool(address tokenA, address tokenB, bool stable) returns(address pool)
func (_IPoolFactory *IPoolFactoryTransactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _IPoolFactory.contract.Transact(opts, "createPool", tokenA, tokenB, stable)
}

// CreatePool is a paid mutator transaction binding the contract method 0x36bf95a0.
//
// Solidity: function createPool(address tokenA, address tokenB, bool stable) returns(address pool)
func (_IPoolFactory *IPoolFactorySession) CreatePool(tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _IPoolFactory.Contract.CreatePool(&_IPoolFactory.TransactOpts, tokenA, tokenB, stable)
}

// CreatePool is a paid mutator transaction binding the contract method 0x36bf95a0.
//
// Solidity: function createPool(address tokenA, address tokenB, bool stable) returns(address pool)
func (_IPoolFactory *IPoolFactoryTransactorSession) CreatePool(tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _IPoolFactory.Contract.CreatePool(&_IPoolFactory.TransactOpts, tokenA, tokenB, stable)
}

// CreatePool0 is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_IPoolFactory *IPoolFactoryTransactor) CreatePool0(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _IPoolFactory.contract.Transact(opts, "createPool0", tokenA, tokenB, fee)
}

// CreatePool0 is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_IPoolFactory *IPoolFactorySession) CreatePool0(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _IPoolFactory.Contract.CreatePool0(&_IPoolFactory.TransactOpts, tokenA, tokenB, fee)
}

// CreatePool0 is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_IPoolFactory *IPoolFactoryTransactorSession) CreatePool0(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _IPoolFactory.Contract.CreatePool0(&_IPoolFactory.TransactOpts, tokenA, tokenB, fee)
}

// SetCustomFee is a paid mutator transaction binding the contract method 0xd49466a8.
//
// Solidity: function setCustomFee(address _pool, uint256 _fee) returns()
func (_IPoolFactory *IPoolFactoryTransactor) SetCustomFee(opts *bind.TransactOpts, _pool common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _IPoolFactory.contract.Transact(opts, "setCustomFee", _pool, _fee)
}

// SetCustomFee is a paid mutator transaction binding the contract method 0xd49466a8.
//
// Solidity: function setCustomFee(address _pool, uint256 _fee) returns()
func (_IPoolFactory *IPoolFactorySession) SetCustomFee(_pool common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _IPoolFactory.Contract.SetCustomFee(&_IPoolFactory.TransactOpts, _pool, _fee)
}

// SetCustomFee is a paid mutator transaction binding the contract method 0xd49466a8.
//
// Solidity: function setCustomFee(address _pool, uint256 _fee) returns()
func (_IPoolFactory *IPoolFactoryTransactorSession) SetCustomFee(_pool common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _IPoolFactory.Contract.SetCustomFee(&_IPoolFactory.TransactOpts, _pool, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0xe1f76b44.
//
// Solidity: function setFee(bool _stable, uint256 _fee) returns()
func (_IPoolFactory *IPoolFactoryTransactor) SetFee(opts *bind.TransactOpts, _stable bool, _fee *big.Int) (*types.Transaction, error) {
	return _IPoolFactory.contract.Transact(opts, "setFee", _stable, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0xe1f76b44.
//
// Solidity: function setFee(bool _stable, uint256 _fee) returns()
func (_IPoolFactory *IPoolFactorySession) SetFee(_stable bool, _fee *big.Int) (*types.Transaction, error) {
	return _IPoolFactory.Contract.SetFee(&_IPoolFactory.TransactOpts, _stable, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0xe1f76b44.
//
// Solidity: function setFee(bool _stable, uint256 _fee) returns()
func (_IPoolFactory *IPoolFactoryTransactorSession) SetFee(_stable bool, _fee *big.Int) (*types.Transaction, error) {
	return _IPoolFactory.Contract.SetFee(&_IPoolFactory.TransactOpts, _stable, _fee)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_IPoolFactory *IPoolFactoryTransactor) SetFeeManager(opts *bind.TransactOpts, _feeManager common.Address) (*types.Transaction, error) {
	return _IPoolFactory.contract.Transact(opts, "setFeeManager", _feeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_IPoolFactory *IPoolFactorySession) SetFeeManager(_feeManager common.Address) (*types.Transaction, error) {
	return _IPoolFactory.Contract.SetFeeManager(&_IPoolFactory.TransactOpts, _feeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_IPoolFactory *IPoolFactoryTransactorSession) SetFeeManager(_feeManager common.Address) (*types.Transaction, error) {
	return _IPoolFactory.Contract.SetFeeManager(&_IPoolFactory.TransactOpts, _feeManager)
}

// SetPauseState is a paid mutator transaction binding the contract method 0xcdb88ad1.
//
// Solidity: function setPauseState(bool _state) returns()
func (_IPoolFactory *IPoolFactoryTransactor) SetPauseState(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _IPoolFactory.contract.Transact(opts, "setPauseState", _state)
}

// SetPauseState is a paid mutator transaction binding the contract method 0xcdb88ad1.
//
// Solidity: function setPauseState(bool _state) returns()
func (_IPoolFactory *IPoolFactorySession) SetPauseState(_state bool) (*types.Transaction, error) {
	return _IPoolFactory.Contract.SetPauseState(&_IPoolFactory.TransactOpts, _state)
}

// SetPauseState is a paid mutator transaction binding the contract method 0xcdb88ad1.
//
// Solidity: function setPauseState(bool _state) returns()
func (_IPoolFactory *IPoolFactoryTransactorSession) SetPauseState(_state bool) (*types.Transaction, error) {
	return _IPoolFactory.Contract.SetPauseState(&_IPoolFactory.TransactOpts, _state)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_IPoolFactory *IPoolFactoryTransactor) SetPauser(opts *bind.TransactOpts, _pauser common.Address) (*types.Transaction, error) {
	return _IPoolFactory.contract.Transact(opts, "setPauser", _pauser)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_IPoolFactory *IPoolFactorySession) SetPauser(_pauser common.Address) (*types.Transaction, error) {
	return _IPoolFactory.Contract.SetPauser(&_IPoolFactory.TransactOpts, _pauser)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_IPoolFactory *IPoolFactoryTransactorSession) SetPauser(_pauser common.Address) (*types.Transaction, error) {
	return _IPoolFactory.Contract.SetPauser(&_IPoolFactory.TransactOpts, _pauser)
}

// SetSinkConverter is a paid mutator transaction binding the contract method 0x37068e7b.
//
// Solidity: function setSinkConverter(address _sinkConvert, address _velo, address _veloV2) returns()
func (_IPoolFactory *IPoolFactoryTransactor) SetSinkConverter(opts *bind.TransactOpts, _sinkConvert common.Address, _velo common.Address, _veloV2 common.Address) (*types.Transaction, error) {
	return _IPoolFactory.contract.Transact(opts, "setSinkConverter", _sinkConvert, _velo, _veloV2)
}

// SetSinkConverter is a paid mutator transaction binding the contract method 0x37068e7b.
//
// Solidity: function setSinkConverter(address _sinkConvert, address _velo, address _veloV2) returns()
func (_IPoolFactory *IPoolFactorySession) SetSinkConverter(_sinkConvert common.Address, _velo common.Address, _veloV2 common.Address) (*types.Transaction, error) {
	return _IPoolFactory.Contract.SetSinkConverter(&_IPoolFactory.TransactOpts, _sinkConvert, _velo, _veloV2)
}

// SetSinkConverter is a paid mutator transaction binding the contract method 0x37068e7b.
//
// Solidity: function setSinkConverter(address _sinkConvert, address _velo, address _veloV2) returns()
func (_IPoolFactory *IPoolFactoryTransactorSession) SetSinkConverter(_sinkConvert common.Address, _velo common.Address, _veloV2 common.Address) (*types.Transaction, error) {
	return _IPoolFactory.Contract.SetSinkConverter(&_IPoolFactory.TransactOpts, _sinkConvert, _velo, _veloV2)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _voter) returns()
func (_IPoolFactory *IPoolFactoryTransactor) SetVoter(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _IPoolFactory.contract.Transact(opts, "setVoter", _voter)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _voter) returns()
func (_IPoolFactory *IPoolFactorySession) SetVoter(_voter common.Address) (*types.Transaction, error) {
	return _IPoolFactory.Contract.SetVoter(&_IPoolFactory.TransactOpts, _voter)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _voter) returns()
func (_IPoolFactory *IPoolFactoryTransactorSession) SetVoter(_voter common.Address) (*types.Transaction, error) {
	return _IPoolFactory.Contract.SetVoter(&_IPoolFactory.TransactOpts, _voter)
}

// IPoolFactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the IPoolFactory contract.
type IPoolFactoryPoolCreatedIterator struct {
	Event *IPoolFactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *IPoolFactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolFactoryPoolCreated)
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
		it.Event = new(IPoolFactoryPoolCreated)
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
func (it *IPoolFactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolFactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolFactoryPoolCreated represents a PoolCreated event raised by the IPoolFactory contract.
type IPoolFactoryPoolCreated struct {
	Token0 common.Address
	Token1 common.Address
	Stable bool
	Pool   common.Address
	Arg4   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x2128d88d14c80cb081c1252a5acff7a264671bf199ce226b53788fb26065005e.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, bool indexed stable, address pool, uint256 arg4)
func (_IPoolFactory *IPoolFactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address, stable []bool) (*IPoolFactoryPoolCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var stableRule []interface{}
	for _, stableItem := range stable {
		stableRule = append(stableRule, stableItem)
	}

	logs, sub, err := _IPoolFactory.contract.FilterLogs(opts, "PoolCreated", token0Rule, token1Rule, stableRule)
	if err != nil {
		return nil, err
	}
	return &IPoolFactoryPoolCreatedIterator{contract: _IPoolFactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x2128d88d14c80cb081c1252a5acff7a264671bf199ce226b53788fb26065005e.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, bool indexed stable, address pool, uint256 arg4)
func (_IPoolFactory *IPoolFactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *IPoolFactoryPoolCreated, token0 []common.Address, token1 []common.Address, stable []bool) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var stableRule []interface{}
	for _, stableItem := range stable {
		stableRule = append(stableRule, stableItem)
	}

	logs, sub, err := _IPoolFactory.contract.WatchLogs(opts, "PoolCreated", token0Rule, token1Rule, stableRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolFactoryPoolCreated)
				if err := _IPoolFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x2128d88d14c80cb081c1252a5acff7a264671bf199ce226b53788fb26065005e.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, bool indexed stable, address pool, uint256 arg4)
func (_IPoolFactory *IPoolFactoryFilterer) ParsePoolCreated(log types.Log) (*IPoolFactoryPoolCreated, error) {
	event := new(IPoolFactoryPoolCreated)
	if err := _IPoolFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolFactorySetCustomFeeIterator is returned from FilterSetCustomFee and is used to iterate over the raw logs and unpacked data for SetCustomFee events raised by the IPoolFactory contract.
type IPoolFactorySetCustomFeeIterator struct {
	Event *IPoolFactorySetCustomFee // Event containing the contract specifics and raw log

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
func (it *IPoolFactorySetCustomFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolFactorySetCustomFee)
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
		it.Event = new(IPoolFactorySetCustomFee)
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
func (it *IPoolFactorySetCustomFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolFactorySetCustomFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolFactorySetCustomFee represents a SetCustomFee event raised by the IPoolFactory contract.
type IPoolFactorySetCustomFee struct {
	Pool common.Address
	Fee  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetCustomFee is a free log retrieval operation binding the contract event 0xae468ce586f9a87660fdffc1448cee942042c16ae2f02046b134b5224f31936b.
//
// Solidity: event SetCustomFee(address indexed pool, uint256 fee)
func (_IPoolFactory *IPoolFactoryFilterer) FilterSetCustomFee(opts *bind.FilterOpts, pool []common.Address) (*IPoolFactorySetCustomFeeIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IPoolFactory.contract.FilterLogs(opts, "SetCustomFee", poolRule)
	if err != nil {
		return nil, err
	}
	return &IPoolFactorySetCustomFeeIterator{contract: _IPoolFactory.contract, event: "SetCustomFee", logs: logs, sub: sub}, nil
}

// WatchSetCustomFee is a free log subscription operation binding the contract event 0xae468ce586f9a87660fdffc1448cee942042c16ae2f02046b134b5224f31936b.
//
// Solidity: event SetCustomFee(address indexed pool, uint256 fee)
func (_IPoolFactory *IPoolFactoryFilterer) WatchSetCustomFee(opts *bind.WatchOpts, sink chan<- *IPoolFactorySetCustomFee, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IPoolFactory.contract.WatchLogs(opts, "SetCustomFee", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolFactorySetCustomFee)
				if err := _IPoolFactory.contract.UnpackLog(event, "SetCustomFee", log); err != nil {
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

// ParseSetCustomFee is a log parse operation binding the contract event 0xae468ce586f9a87660fdffc1448cee942042c16ae2f02046b134b5224f31936b.
//
// Solidity: event SetCustomFee(address indexed pool, uint256 fee)
func (_IPoolFactory *IPoolFactoryFilterer) ParseSetCustomFee(log types.Log) (*IPoolFactorySetCustomFee, error) {
	event := new(IPoolFactorySetCustomFee)
	if err := _IPoolFactory.contract.UnpackLog(event, "SetCustomFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolFactorySetFeeManagerIterator is returned from FilterSetFeeManager and is used to iterate over the raw logs and unpacked data for SetFeeManager events raised by the IPoolFactory contract.
type IPoolFactorySetFeeManagerIterator struct {
	Event *IPoolFactorySetFeeManager // Event containing the contract specifics and raw log

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
func (it *IPoolFactorySetFeeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolFactorySetFeeManager)
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
		it.Event = new(IPoolFactorySetFeeManager)
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
func (it *IPoolFactorySetFeeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolFactorySetFeeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolFactorySetFeeManager represents a SetFeeManager event raised by the IPoolFactory contract.
type IPoolFactorySetFeeManager struct {
	FeeManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetFeeManager is a free log retrieval operation binding the contract event 0x5d0517e3a4eabea892d9750138cd21d4a6cf3b935b43d0598df7055f463819b2.
//
// Solidity: event SetFeeManager(address feeManager)
func (_IPoolFactory *IPoolFactoryFilterer) FilterSetFeeManager(opts *bind.FilterOpts) (*IPoolFactorySetFeeManagerIterator, error) {

	logs, sub, err := _IPoolFactory.contract.FilterLogs(opts, "SetFeeManager")
	if err != nil {
		return nil, err
	}
	return &IPoolFactorySetFeeManagerIterator{contract: _IPoolFactory.contract, event: "SetFeeManager", logs: logs, sub: sub}, nil
}

// WatchSetFeeManager is a free log subscription operation binding the contract event 0x5d0517e3a4eabea892d9750138cd21d4a6cf3b935b43d0598df7055f463819b2.
//
// Solidity: event SetFeeManager(address feeManager)
func (_IPoolFactory *IPoolFactoryFilterer) WatchSetFeeManager(opts *bind.WatchOpts, sink chan<- *IPoolFactorySetFeeManager) (event.Subscription, error) {

	logs, sub, err := _IPoolFactory.contract.WatchLogs(opts, "SetFeeManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolFactorySetFeeManager)
				if err := _IPoolFactory.contract.UnpackLog(event, "SetFeeManager", log); err != nil {
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

// ParseSetFeeManager is a log parse operation binding the contract event 0x5d0517e3a4eabea892d9750138cd21d4a6cf3b935b43d0598df7055f463819b2.
//
// Solidity: event SetFeeManager(address feeManager)
func (_IPoolFactory *IPoolFactoryFilterer) ParseSetFeeManager(log types.Log) (*IPoolFactorySetFeeManager, error) {
	event := new(IPoolFactorySetFeeManager)
	if err := _IPoolFactory.contract.UnpackLog(event, "SetFeeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolFactorySetPauseStateIterator is returned from FilterSetPauseState and is used to iterate over the raw logs and unpacked data for SetPauseState events raised by the IPoolFactory contract.
type IPoolFactorySetPauseStateIterator struct {
	Event *IPoolFactorySetPauseState // Event containing the contract specifics and raw log

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
func (it *IPoolFactorySetPauseStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolFactorySetPauseState)
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
		it.Event = new(IPoolFactorySetPauseState)
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
func (it *IPoolFactorySetPauseStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolFactorySetPauseStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolFactorySetPauseState represents a SetPauseState event raised by the IPoolFactory contract.
type IPoolFactorySetPauseState struct {
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetPauseState is a free log retrieval operation binding the contract event 0x0d76538efc408318a051137c2720a9e82902acdbd46b802d488b74ca3a09a116.
//
// Solidity: event SetPauseState(bool state)
func (_IPoolFactory *IPoolFactoryFilterer) FilterSetPauseState(opts *bind.FilterOpts) (*IPoolFactorySetPauseStateIterator, error) {

	logs, sub, err := _IPoolFactory.contract.FilterLogs(opts, "SetPauseState")
	if err != nil {
		return nil, err
	}
	return &IPoolFactorySetPauseStateIterator{contract: _IPoolFactory.contract, event: "SetPauseState", logs: logs, sub: sub}, nil
}

// WatchSetPauseState is a free log subscription operation binding the contract event 0x0d76538efc408318a051137c2720a9e82902acdbd46b802d488b74ca3a09a116.
//
// Solidity: event SetPauseState(bool state)
func (_IPoolFactory *IPoolFactoryFilterer) WatchSetPauseState(opts *bind.WatchOpts, sink chan<- *IPoolFactorySetPauseState) (event.Subscription, error) {

	logs, sub, err := _IPoolFactory.contract.WatchLogs(opts, "SetPauseState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolFactorySetPauseState)
				if err := _IPoolFactory.contract.UnpackLog(event, "SetPauseState", log); err != nil {
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

// ParseSetPauseState is a log parse operation binding the contract event 0x0d76538efc408318a051137c2720a9e82902acdbd46b802d488b74ca3a09a116.
//
// Solidity: event SetPauseState(bool state)
func (_IPoolFactory *IPoolFactoryFilterer) ParseSetPauseState(log types.Log) (*IPoolFactorySetPauseState, error) {
	event := new(IPoolFactorySetPauseState)
	if err := _IPoolFactory.contract.UnpackLog(event, "SetPauseState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolFactorySetPauserIterator is returned from FilterSetPauser and is used to iterate over the raw logs and unpacked data for SetPauser events raised by the IPoolFactory contract.
type IPoolFactorySetPauserIterator struct {
	Event *IPoolFactorySetPauser // Event containing the contract specifics and raw log

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
func (it *IPoolFactorySetPauserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolFactorySetPauser)
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
		it.Event = new(IPoolFactorySetPauser)
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
func (it *IPoolFactorySetPauserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolFactorySetPauserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolFactorySetPauser represents a SetPauser event raised by the IPoolFactory contract.
type IPoolFactorySetPauser struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetPauser is a free log retrieval operation binding the contract event 0xe02efb9e8f0fc21546730ab32d594f62d586e1bbb15bb5045edd0b1878a77b35.
//
// Solidity: event SetPauser(address pauser)
func (_IPoolFactory *IPoolFactoryFilterer) FilterSetPauser(opts *bind.FilterOpts) (*IPoolFactorySetPauserIterator, error) {

	logs, sub, err := _IPoolFactory.contract.FilterLogs(opts, "SetPauser")
	if err != nil {
		return nil, err
	}
	return &IPoolFactorySetPauserIterator{contract: _IPoolFactory.contract, event: "SetPauser", logs: logs, sub: sub}, nil
}

// WatchSetPauser is a free log subscription operation binding the contract event 0xe02efb9e8f0fc21546730ab32d594f62d586e1bbb15bb5045edd0b1878a77b35.
//
// Solidity: event SetPauser(address pauser)
func (_IPoolFactory *IPoolFactoryFilterer) WatchSetPauser(opts *bind.WatchOpts, sink chan<- *IPoolFactorySetPauser) (event.Subscription, error) {

	logs, sub, err := _IPoolFactory.contract.WatchLogs(opts, "SetPauser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolFactorySetPauser)
				if err := _IPoolFactory.contract.UnpackLog(event, "SetPauser", log); err != nil {
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

// ParseSetPauser is a log parse operation binding the contract event 0xe02efb9e8f0fc21546730ab32d594f62d586e1bbb15bb5045edd0b1878a77b35.
//
// Solidity: event SetPauser(address pauser)
func (_IPoolFactory *IPoolFactoryFilterer) ParseSetPauser(log types.Log) (*IPoolFactorySetPauser, error) {
	event := new(IPoolFactorySetPauser)
	if err := _IPoolFactory.contract.UnpackLog(event, "SetPauser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolFactorySetVoterIterator is returned from FilterSetVoter and is used to iterate over the raw logs and unpacked data for SetVoter events raised by the IPoolFactory contract.
type IPoolFactorySetVoterIterator struct {
	Event *IPoolFactorySetVoter // Event containing the contract specifics and raw log

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
func (it *IPoolFactorySetVoterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolFactorySetVoter)
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
		it.Event = new(IPoolFactorySetVoter)
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
func (it *IPoolFactorySetVoterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolFactorySetVoterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolFactorySetVoter represents a SetVoter event raised by the IPoolFactory contract.
type IPoolFactorySetVoter struct {
	Voter common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetVoter is a free log retrieval operation binding the contract event 0xc6ff127433b785c51da9ae4088ee184c909b1a55b9afd82ae6c64224d3bc15d2.
//
// Solidity: event SetVoter(address voter)
func (_IPoolFactory *IPoolFactoryFilterer) FilterSetVoter(opts *bind.FilterOpts) (*IPoolFactorySetVoterIterator, error) {

	logs, sub, err := _IPoolFactory.contract.FilterLogs(opts, "SetVoter")
	if err != nil {
		return nil, err
	}
	return &IPoolFactorySetVoterIterator{contract: _IPoolFactory.contract, event: "SetVoter", logs: logs, sub: sub}, nil
}

// WatchSetVoter is a free log subscription operation binding the contract event 0xc6ff127433b785c51da9ae4088ee184c909b1a55b9afd82ae6c64224d3bc15d2.
//
// Solidity: event SetVoter(address voter)
func (_IPoolFactory *IPoolFactoryFilterer) WatchSetVoter(opts *bind.WatchOpts, sink chan<- *IPoolFactorySetVoter) (event.Subscription, error) {

	logs, sub, err := _IPoolFactory.contract.WatchLogs(opts, "SetVoter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolFactorySetVoter)
				if err := _IPoolFactory.contract.UnpackLog(event, "SetVoter", log); err != nil {
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

// ParseSetVoter is a log parse operation binding the contract event 0xc6ff127433b785c51da9ae4088ee184c909b1a55b9afd82ae6c64224d3bc15d2.
//
// Solidity: event SetVoter(address voter)
func (_IPoolFactory *IPoolFactoryFilterer) ParseSetVoter(log types.Log) (*IPoolFactorySetVoter, error) {
	event := new(IPoolFactorySetVoter)
	if err := _IPoolFactory.contract.UnpackLog(event, "SetVoter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFactoryMetaData contains all meta data concerning the PoolFactory contract.
var PoolFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FeeInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotFeeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotPauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSinkConverter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotVoter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFee\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"SetCustomFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"}],\"name\":\"SetFeeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"SetPauseState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"SetPauser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"SetVoter\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_FEE_INDICATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPoolsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"customFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_stable\",\"type\":\"bool\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPair\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setCustomFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeManager\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"setPauseState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"setPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sinkConverter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_velo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_veloV2\",\"type\":\"address\"}],\"name\":\"setSinkConverter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"setVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sinkConverter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"velo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"veloV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"volatileFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bc063e1a": "MAX_FEE()",
		"38c55d46": "ZERO_FEE_INDICATOR()",
		"41d1de97": "allPools(uint256)",
		"efde4e64": "allPoolsLength()",
		"82dfdce4": "createPair(address,address,bool)",
		"36bf95a0": "createPool(address,address,bool)",
		"a1671295": "createPool(address,address,uint24)",
		"4d419abc": "customFee(address)",
		"d0fb0203": "feeManager()",
		"cc56b2c5": "getFee(address,bool)",
		"6801cc30": "getPair(address,address,bool)",
		"79bc57d5": "getPool(address,address,bool)",
		"1698ee82": "getPool(address,address,uint24)",
		"5c60da1b": "implementation()",
		"e5e31b13": "isPair(address)",
		"b187bd26": "isPaused()",
		"5b16ebb7": "isPool(address)",
		"9fd0506d": "pauser()",
		"d49466a8": "setCustomFee(address,uint256)",
		"e1f76b44": "setFee(bool,uint256)",
		"472d35b9": "setFeeManager(address)",
		"cdb88ad1": "setPauseState(bool)",
		"2d88af4a": "setPauser(address)",
		"37068e7b": "setSinkConverter(address,address,address)",
		"4bc2a657": "setVoter(address)",
		"8e39ee16": "sinkConverter()",
		"40bbd775": "stableFee()",
		"8c7c53ce": "velo()",
		"c6751c09": "veloV2()",
		"5084ed03": "volatileFee()",
		"46c96aac": "voter()",
	},
	Bin: "0x60a060405234801561000f575f80fd5b506040516111c23803806111c283398101604081905261002e91610094565b6001600160a01b031660805260048054336001600160a01b031991821681179092555f8054600380548416851790556007805490931684179092556001600160a81b031990911661010090920260ff19169190911790556005600155601e6002556100c1565b5f602082840312156100a4575f80fd5b81516001600160a01b03811681146100ba575f80fd5b9392505050565b6080516110e26100e05f395f818161031f01526106db01526110e25ff3fe608060405234801561000f575f80fd5b50600436106101dc575f3560e01c806379bc57d511610109578063c6751c091161009e578063d49466a81161006e578063d49466a814610447578063e1f76b441461045a578063e5e31b13146102df578063efde4e641461046d575f80fd5b8063c6751c09146103fb578063cc56b2c51461040e578063cdb88ad114610421578063d0fb020314610434575f80fd5b80639fd0506d116100d95780639fd0506d146103bd578063a1671295146103d4578063b187bd26146103e7578063bc063e1a146103f3575f80fd5b806379bc57d51461034157806382dfdce4146103845780638c7c53ce146103975780638e39ee16146103aa575f80fd5b806346c96aac1161017f5780635084ed031161014f5780635084ed03146102d65780635b16ebb7146102df5780635c60da1b1461031a5780636801cc3014610341575f80fd5b806346c96aac1461027e578063472d35b9146102915780634bc2a657146102a45780634d419abc146102b7575f80fd5b806337068e7b116101ba57806337068e7b1461023857806338c55d461461024b57806340bbd7751461026257806341d1de971461026b575f80fd5b80631698ee82146101e05780632d88af4a1461021057806336bf95a014610225575b5f80fd5b6101f36101ee366004610ef3565b610475565b6040516001600160a01b0390911681526020015b60405180910390f35b61022361021e366004610f3e565b610503565b005b6101f3610233366004610f6d565b6105b5565b610223610246366004610fad565b61086a565b6102546101a481565b604051908152602001610207565b61025460015481565b6101f3610279366004610fe4565b610a81565b6004546101f3906001600160a01b031681565b61022361029f366004610f3e565b610aa9565b6102236102b2366004610f3e565b610b49565b6102546102c5366004610f3e565b600b6020525f908152604090205481565b61025460025481565b61030a6102ed366004610f3e565b6001600160a01b03165f908152600a602052604090205460ff1690565b6040519015158152602001610207565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b6101f361034f366004610f6d565b6001600160a01b039283165f908152600860209081526040808320948616835293815283822092151582529190915220541690565b6101f3610392366004610f6d565b610bc2565b6005546101f3906001600160a01b031681565b6007546101f3906001600160a01b031681565b5f546101f39061010090046001600160a01b031681565b6101f36103e2366004610ef3565b610bce565b5f5461030a9060ff1681565b610254606481565b6006546101f3906001600160a01b031681565b61025461041c366004610ffb565b610c13565b61022361042f36600461102c565b610c62565b6003546101f3906001600160a01b031681565b610223610455366004611045565b610cd1565b61022361046836600461106d565b610dbb565b600954610254565b5f60018262ffffff16116104f9578162ffffff166001146104c4576001600160a01b038085165f90815260086020908152604080832087851684528252808320838052909152902054166104fb565b6001600160a01b038085165f9081526008602090815260408083208785168452825280832060018452909152902054166104fb565b5f5b949350505050565b5f5461010090046001600160a01b031633146105325760405163492f678160e01b815260040160405180910390fd5b6001600160a01b0381166105595760405163d92e233d60e01b815260040160405180910390fd5b5f8054610100600160a81b0319166101006001600160a01b038416908102919091179091556040519081527fe02efb9e8f0fc21546730ab32d594f62d586e1bbb15bb5045edd0b1878a77b35906020015b60405180910390a150565b5f826001600160a01b0316846001600160a01b0316036105e85760405163367558c360e01b815260040160405180910390fd5b5f80846001600160a01b0316866001600160a01b03161061060a57848661060d565b85855b90925090506001600160a01b0382166106395760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038281165f908152600860209081526040808320858516845282528083208815158452909152902054161561068857604051630188c99160e11b815260040160405180910390fd5b6040516bffffffffffffffffffffffff19606084811b8216602084015283901b16603482015284151560f81b60488201525f906049016040516020818303038152906040528051906020012090506107007f000000000000000000000000000000000000000000000000000000000000000082610e3a565b604051631c9776b560e31b81526001600160a01b038581166004830152848116602483015287151560448301529195509085169063e4bbb5a8906064015f604051808303815f87803b158015610754575f80fd5b505af1158015610766573d5f803e3d5ffd5b505050506001600160a01b038381165f8181526008602081815260408084208887168086529083528185208c15158087529084528286208054988d166001600160a01b0319998a16811790915582875294845282862087875284528286208187528452828620805489168617905560098054600181810183557f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af9091018054909a168717909955858752600a855295839020805460ff191690981790975593548151938452918301919091529192917f2128d88d14c80cb081c1252a5acff7a264671bf199ce226b53788fb26065005e910160405180910390a45050509392505050565b6007546001600160a01b03163314610895576040516315517aa960e31b815260040160405180910390fd5b600780546001600160a01b038086166001600160a01b031992831617909255600580548584169083168117909155600680549385169390921683179091555f918291106108e35782846108e6565b83835b600780546001600160a01b038481165f818152600860208181526040808420898716808652908352818520600180875281855283872080546001600160a01b03199081169b8b169b909b1790558a54838852958552838720888852855283872081885280865284882080548c16978b16979097179096558a5487805291855283872080548b16928a1692909217909155895494845282862080548a169589169590951790945588546009805480870182557f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af018054909a1691891691909117909855885487168552600a835293819020805460ff1916841790559654955487519690951686528501939093529597509395509392917f2128d88d14c80cb081c1252a5acff7a264671bf199ce226b53788fb26065005e910160405180910390a4600754600954604080516001600160a01b03938416815260208101929092525f9284811692908616917f2128d88d14c80cb081c1252a5acff7a264671bf199ce226b53788fb26065005e910160405180910390a45050505050565b60098181548110610a90575f80fd5b5f918252602090912001546001600160a01b0316905081565b6003546001600160a01b03163314610ad45760405163f5d267eb60e01b815260040160405180910390fd5b6001600160a01b038116610afb5760405163d92e233d60e01b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0383169081179091556040519081527f5d0517e3a4eabea892d9750138cd21d4a6cf3b935b43d0598df7055f463819b2906020016105aa565b6004546001600160a01b03163314610b745760405163c18384c160e01b815260040160405180910390fd5b600480546001600160a01b0319166001600160a01b0383169081179091556040519081527fc6ff127433b785c51da9ae4088ee184c909b1a55b9afd82ae6c64224d3bc15d2906020016105aa565b5f6104fb8484846105b5565b5f60018262ffffff161115610bf6576040516352dadcf960e01b815260040160405180910390fd5b600162ffffff831614610c0a8585836105b5565b95945050505050565b6001600160a01b0382165f908152600b60205260408120546101a48114610c5657805f03610c505782610c4857600254610c58565b600154610c58565b80610c58565b5f5b9150505b92915050565b5f5461010090046001600160a01b03163314610c915760405163492f678160e01b815260040160405180910390fd5b5f805460ff19168215159081179091556040519081527f0d76538efc408318a051137c2720a9e82902acdbd46b802d488b74ca3a09a116906020016105aa565b6003546001600160a01b03163314610cfc5760405163f5d267eb60e01b815260040160405180910390fd5b606481118015610d0e57506101a48114155b15610d2c5760405163cd4e616760e01b815260040160405180910390fd5b6001600160a01b0382165f908152600a602052604090205460ff16610d635760405162820f3560e61b815260040160405180910390fd5b6001600160a01b0382165f818152600b602052604090819020839055517fae468ce586f9a87660fdffc1448cee942042c16ae2f02046b134b5224f31936b90610daf9084815260200190565b60405180910390a25050565b6003546001600160a01b03163314610de65760405163f5d267eb60e01b815260040160405180910390fd5b6064811115610e085760405163cd4e616760e01b815260040160405180910390fd5b805f03610e285760405163af13986d60e01b815260040160405180910390fd5b8115610e345760015550565b60025550565b5f763d602d80600a3d3981f3363d3d373d3d3d363d730000008360601b60e81c175f526e5af43d82803e903d91602b57fd5bf38360781b1760205281603760095ff590506001600160a01b038116610c5c5760405162461bcd60e51b815260206004820152601760248201527f455243313136373a2063726561746532206661696c6564000000000000000000604482015260640160405180910390fd5b80356001600160a01b0381168114610eee575f80fd5b919050565b5f805f60608486031215610f05575f80fd5b610f0e84610ed8565b9250610f1c60208501610ed8565b9150604084013562ffffff81168114610f33575f80fd5b809150509250925092565b5f60208284031215610f4e575f80fd5b610f5782610ed8565b9392505050565b80358015158114610eee575f80fd5b5f805f60608486031215610f7f575f80fd5b610f8884610ed8565b9250610f9660208501610ed8565b9150610fa460408501610f5e565b90509250925092565b5f805f60608486031215610fbf575f80fd5b610fc884610ed8565b9250610fd660208501610ed8565b9150610fa460408501610ed8565b5f60208284031215610ff4575f80fd5b5035919050565b5f806040838503121561100c575f80fd5b61101583610ed8565b915061102360208401610f5e565b90509250929050565b5f6020828403121561103c575f80fd5b610f5782610f5e565b5f8060408385031215611056575f80fd5b61105f83610ed8565b946020939093013593505050565b5f806040838503121561107e575f80fd5b61105f83610f5e56fea2646970667358221220d487009a6c788088fbd15e7079bc3e2ce6b2d55e06a8b2faed6b368083a389f864736f6c637827302e382e32322d646576656c6f702e323032332e382e342b636f6d6d69742e35313137313235370058",
}

// PoolFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolFactoryMetaData.ABI instead.
var PoolFactoryABI = PoolFactoryMetaData.ABI

// Deprecated: Use PoolFactoryMetaData.Sigs instead.
// PoolFactoryFuncSigs maps the 4-byte function signature to its string representation.
var PoolFactoryFuncSigs = PoolFactoryMetaData.Sigs

// PoolFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoolFactoryMetaData.Bin instead.
var PoolFactoryBin = PoolFactoryMetaData.Bin

// DeployPoolFactory deploys a new Ethereum contract, binding an instance of PoolFactory to it.
func DeployPoolFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _implementation common.Address) (common.Address, *types.Transaction, *PoolFactory, error) {
	parsed, err := PoolFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoolFactoryBin), backend, _implementation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoolFactory{PoolFactoryCaller: PoolFactoryCaller{contract: contract}, PoolFactoryTransactor: PoolFactoryTransactor{contract: contract}, PoolFactoryFilterer: PoolFactoryFilterer{contract: contract}}, nil
}

// PoolFactory is an auto generated Go binding around an Ethereum contract.
type PoolFactory struct {
	PoolFactoryCaller     // Read-only binding to the contract
	PoolFactoryTransactor // Write-only binding to the contract
	PoolFactoryFilterer   // Log filterer for contract events
}

// PoolFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolFactorySession struct {
	Contract     *PoolFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolFactoryCallerSession struct {
	Contract *PoolFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PoolFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolFactoryTransactorSession struct {
	Contract     *PoolFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PoolFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolFactoryRaw struct {
	Contract *PoolFactory // Generic contract binding to access the raw methods on
}

// PoolFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolFactoryCallerRaw struct {
	Contract *PoolFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PoolFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolFactoryTransactorRaw struct {
	Contract *PoolFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolFactory creates a new instance of PoolFactory, bound to a specific deployed contract.
func NewPoolFactory(address common.Address, backend bind.ContractBackend) (*PoolFactory, error) {
	contract, err := bindPoolFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolFactory{PoolFactoryCaller: PoolFactoryCaller{contract: contract}, PoolFactoryTransactor: PoolFactoryTransactor{contract: contract}, PoolFactoryFilterer: PoolFactoryFilterer{contract: contract}}, nil
}

// NewPoolFactoryCaller creates a new read-only instance of PoolFactory, bound to a specific deployed contract.
func NewPoolFactoryCaller(address common.Address, caller bind.ContractCaller) (*PoolFactoryCaller, error) {
	contract, err := bindPoolFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolFactoryCaller{contract: contract}, nil
}

// NewPoolFactoryTransactor creates a new write-only instance of PoolFactory, bound to a specific deployed contract.
func NewPoolFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolFactoryTransactor, error) {
	contract, err := bindPoolFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolFactoryTransactor{contract: contract}, nil
}

// NewPoolFactoryFilterer creates a new log filterer instance of PoolFactory, bound to a specific deployed contract.
func NewPoolFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFactoryFilterer, error) {
	contract, err := bindPoolFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFactoryFilterer{contract: contract}, nil
}

// bindPoolFactory binds a generic wrapper to an already deployed contract.
func bindPoolFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolFactory *PoolFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolFactory.Contract.PoolFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolFactory *PoolFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolFactory.Contract.PoolFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolFactory *PoolFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolFactory.Contract.PoolFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolFactory *PoolFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolFactory *PoolFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolFactory *PoolFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolFactory.Contract.contract.Transact(opts, method, params...)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_PoolFactory *PoolFactoryCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_PoolFactory *PoolFactorySession) MAXFEE() (*big.Int, error) {
	return _PoolFactory.Contract.MAXFEE(&_PoolFactory.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_PoolFactory *PoolFactoryCallerSession) MAXFEE() (*big.Int, error) {
	return _PoolFactory.Contract.MAXFEE(&_PoolFactory.CallOpts)
}

// ZEROFEEINDICATOR is a free data retrieval call binding the contract method 0x38c55d46.
//
// Solidity: function ZERO_FEE_INDICATOR() view returns(uint256)
func (_PoolFactory *PoolFactoryCaller) ZEROFEEINDICATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "ZERO_FEE_INDICATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZEROFEEINDICATOR is a free data retrieval call binding the contract method 0x38c55d46.
//
// Solidity: function ZERO_FEE_INDICATOR() view returns(uint256)
func (_PoolFactory *PoolFactorySession) ZEROFEEINDICATOR() (*big.Int, error) {
	return _PoolFactory.Contract.ZEROFEEINDICATOR(&_PoolFactory.CallOpts)
}

// ZEROFEEINDICATOR is a free data retrieval call binding the contract method 0x38c55d46.
//
// Solidity: function ZERO_FEE_INDICATOR() view returns(uint256)
func (_PoolFactory *PoolFactoryCallerSession) ZEROFEEINDICATOR() (*big.Int, error) {
	return _PoolFactory.Contract.ZEROFEEINDICATOR(&_PoolFactory.CallOpts)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_PoolFactory *PoolFactoryCaller) AllPools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "allPools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_PoolFactory *PoolFactorySession) AllPools(arg0 *big.Int) (common.Address, error) {
	return _PoolFactory.Contract.AllPools(&_PoolFactory.CallOpts, arg0)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) AllPools(arg0 *big.Int) (common.Address, error) {
	return _PoolFactory.Contract.AllPools(&_PoolFactory.CallOpts, arg0)
}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (_PoolFactory *PoolFactoryCaller) AllPoolsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "allPoolsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (_PoolFactory *PoolFactorySession) AllPoolsLength() (*big.Int, error) {
	return _PoolFactory.Contract.AllPoolsLength(&_PoolFactory.CallOpts)
}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (_PoolFactory *PoolFactoryCallerSession) AllPoolsLength() (*big.Int, error) {
	return _PoolFactory.Contract.AllPoolsLength(&_PoolFactory.CallOpts)
}

// CustomFee is a free data retrieval call binding the contract method 0x4d419abc.
//
// Solidity: function customFee(address ) view returns(uint256)
func (_PoolFactory *PoolFactoryCaller) CustomFee(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "customFee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CustomFee is a free data retrieval call binding the contract method 0x4d419abc.
//
// Solidity: function customFee(address ) view returns(uint256)
func (_PoolFactory *PoolFactorySession) CustomFee(arg0 common.Address) (*big.Int, error) {
	return _PoolFactory.Contract.CustomFee(&_PoolFactory.CallOpts, arg0)
}

// CustomFee is a free data retrieval call binding the contract method 0x4d419abc.
//
// Solidity: function customFee(address ) view returns(uint256)
func (_PoolFactory *PoolFactoryCallerSession) CustomFee(arg0 common.Address) (*big.Int, error) {
	return _PoolFactory.Contract.CustomFee(&_PoolFactory.CallOpts, arg0)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_PoolFactory *PoolFactoryCaller) FeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_PoolFactory *PoolFactorySession) FeeManager() (common.Address, error) {
	return _PoolFactory.Contract.FeeManager(&_PoolFactory.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) FeeManager() (common.Address, error) {
	return _PoolFactory.Contract.FeeManager(&_PoolFactory.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xcc56b2c5.
//
// Solidity: function getFee(address pool, bool _stable) view returns(uint256)
func (_PoolFactory *PoolFactoryCaller) GetFee(opts *bind.CallOpts, pool common.Address, _stable bool) (*big.Int, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "getFee", pool, _stable)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xcc56b2c5.
//
// Solidity: function getFee(address pool, bool _stable) view returns(uint256)
func (_PoolFactory *PoolFactorySession) GetFee(pool common.Address, _stable bool) (*big.Int, error) {
	return _PoolFactory.Contract.GetFee(&_PoolFactory.CallOpts, pool, _stable)
}

// GetFee is a free data retrieval call binding the contract method 0xcc56b2c5.
//
// Solidity: function getFee(address pool, bool _stable) view returns(uint256)
func (_PoolFactory *PoolFactoryCallerSession) GetFee(pool common.Address, _stable bool) (*big.Int, error) {
	return _PoolFactory.Contract.GetFee(&_PoolFactory.CallOpts, pool, _stable)
}

// GetPair is a free data retrieval call binding the contract method 0x6801cc30.
//
// Solidity: function getPair(address tokenA, address tokenB, bool stable) view returns(address)
func (_PoolFactory *PoolFactoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "getPair", tokenA, tokenB, stable)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0x6801cc30.
//
// Solidity: function getPair(address tokenA, address tokenB, bool stable) view returns(address)
func (_PoolFactory *PoolFactorySession) GetPair(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _PoolFactory.Contract.GetPair(&_PoolFactory.CallOpts, tokenA, tokenB, stable)
}

// GetPair is a free data retrieval call binding the contract method 0x6801cc30.
//
// Solidity: function getPair(address tokenA, address tokenB, bool stable) view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _PoolFactory.Contract.GetPair(&_PoolFactory.CallOpts, tokenA, tokenB, stable)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address tokenA, address tokenB, uint24 fee) view returns(address)
func (_PoolFactory *PoolFactoryCaller) GetPool(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, fee *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "getPool", tokenA, tokenB, fee)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address tokenA, address tokenB, uint24 fee) view returns(address)
func (_PoolFactory *PoolFactorySession) GetPool(tokenA common.Address, tokenB common.Address, fee *big.Int) (common.Address, error) {
	return _PoolFactory.Contract.GetPool(&_PoolFactory.CallOpts, tokenA, tokenB, fee)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address tokenA, address tokenB, uint24 fee) view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) GetPool(tokenA common.Address, tokenB common.Address, fee *big.Int) (common.Address, error) {
	return _PoolFactory.Contract.GetPool(&_PoolFactory.CallOpts, tokenA, tokenB, fee)
}

// GetPool0 is a free data retrieval call binding the contract method 0x79bc57d5.
//
// Solidity: function getPool(address tokenA, address tokenB, bool stable) view returns(address)
func (_PoolFactory *PoolFactoryCaller) GetPool0(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "getPool0", tokenA, tokenB, stable)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool0 is a free data retrieval call binding the contract method 0x79bc57d5.
//
// Solidity: function getPool(address tokenA, address tokenB, bool stable) view returns(address)
func (_PoolFactory *PoolFactorySession) GetPool0(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _PoolFactory.Contract.GetPool0(&_PoolFactory.CallOpts, tokenA, tokenB, stable)
}

// GetPool0 is a free data retrieval call binding the contract method 0x79bc57d5.
//
// Solidity: function getPool(address tokenA, address tokenB, bool stable) view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) GetPool0(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _PoolFactory.Contract.GetPool0(&_PoolFactory.CallOpts, tokenA, tokenB, stable)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_PoolFactory *PoolFactoryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_PoolFactory *PoolFactorySession) Implementation() (common.Address, error) {
	return _PoolFactory.Contract.Implementation(&_PoolFactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) Implementation() (common.Address, error) {
	return _PoolFactory.Contract.Implementation(&_PoolFactory.CallOpts)
}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address pool) view returns(bool)
func (_PoolFactory *PoolFactoryCaller) IsPair(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "isPair", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address pool) view returns(bool)
func (_PoolFactory *PoolFactorySession) IsPair(pool common.Address) (bool, error) {
	return _PoolFactory.Contract.IsPair(&_PoolFactory.CallOpts, pool)
}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address pool) view returns(bool)
func (_PoolFactory *PoolFactoryCallerSession) IsPair(pool common.Address) (bool, error) {
	return _PoolFactory.Contract.IsPair(&_PoolFactory.CallOpts, pool)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_PoolFactory *PoolFactoryCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_PoolFactory *PoolFactorySession) IsPaused() (bool, error) {
	return _PoolFactory.Contract.IsPaused(&_PoolFactory.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_PoolFactory *PoolFactoryCallerSession) IsPaused() (bool, error) {
	return _PoolFactory.Contract.IsPaused(&_PoolFactory.CallOpts)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address pool) view returns(bool)
func (_PoolFactory *PoolFactoryCaller) IsPool(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "isPool", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address pool) view returns(bool)
func (_PoolFactory *PoolFactorySession) IsPool(pool common.Address) (bool, error) {
	return _PoolFactory.Contract.IsPool(&_PoolFactory.CallOpts, pool)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address pool) view returns(bool)
func (_PoolFactory *PoolFactoryCallerSession) IsPool(pool common.Address) (bool, error) {
	return _PoolFactory.Contract.IsPool(&_PoolFactory.CallOpts, pool)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_PoolFactory *PoolFactoryCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_PoolFactory *PoolFactorySession) Pauser() (common.Address, error) {
	return _PoolFactory.Contract.Pauser(&_PoolFactory.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) Pauser() (common.Address, error) {
	return _PoolFactory.Contract.Pauser(&_PoolFactory.CallOpts)
}

// SinkConverter is a free data retrieval call binding the contract method 0x8e39ee16.
//
// Solidity: function sinkConverter() view returns(address)
func (_PoolFactory *PoolFactoryCaller) SinkConverter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "sinkConverter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SinkConverter is a free data retrieval call binding the contract method 0x8e39ee16.
//
// Solidity: function sinkConverter() view returns(address)
func (_PoolFactory *PoolFactorySession) SinkConverter() (common.Address, error) {
	return _PoolFactory.Contract.SinkConverter(&_PoolFactory.CallOpts)
}

// SinkConverter is a free data retrieval call binding the contract method 0x8e39ee16.
//
// Solidity: function sinkConverter() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) SinkConverter() (common.Address, error) {
	return _PoolFactory.Contract.SinkConverter(&_PoolFactory.CallOpts)
}

// StableFee is a free data retrieval call binding the contract method 0x40bbd775.
//
// Solidity: function stableFee() view returns(uint256)
func (_PoolFactory *PoolFactoryCaller) StableFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "stableFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableFee is a free data retrieval call binding the contract method 0x40bbd775.
//
// Solidity: function stableFee() view returns(uint256)
func (_PoolFactory *PoolFactorySession) StableFee() (*big.Int, error) {
	return _PoolFactory.Contract.StableFee(&_PoolFactory.CallOpts)
}

// StableFee is a free data retrieval call binding the contract method 0x40bbd775.
//
// Solidity: function stableFee() view returns(uint256)
func (_PoolFactory *PoolFactoryCallerSession) StableFee() (*big.Int, error) {
	return _PoolFactory.Contract.StableFee(&_PoolFactory.CallOpts)
}

// Velo is a free data retrieval call binding the contract method 0x8c7c53ce.
//
// Solidity: function velo() view returns(address)
func (_PoolFactory *PoolFactoryCaller) Velo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "velo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Velo is a free data retrieval call binding the contract method 0x8c7c53ce.
//
// Solidity: function velo() view returns(address)
func (_PoolFactory *PoolFactorySession) Velo() (common.Address, error) {
	return _PoolFactory.Contract.Velo(&_PoolFactory.CallOpts)
}

// Velo is a free data retrieval call binding the contract method 0x8c7c53ce.
//
// Solidity: function velo() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) Velo() (common.Address, error) {
	return _PoolFactory.Contract.Velo(&_PoolFactory.CallOpts)
}

// VeloV2 is a free data retrieval call binding the contract method 0xc6751c09.
//
// Solidity: function veloV2() view returns(address)
func (_PoolFactory *PoolFactoryCaller) VeloV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "veloV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VeloV2 is a free data retrieval call binding the contract method 0xc6751c09.
//
// Solidity: function veloV2() view returns(address)
func (_PoolFactory *PoolFactorySession) VeloV2() (common.Address, error) {
	return _PoolFactory.Contract.VeloV2(&_PoolFactory.CallOpts)
}

// VeloV2 is a free data retrieval call binding the contract method 0xc6751c09.
//
// Solidity: function veloV2() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) VeloV2() (common.Address, error) {
	return _PoolFactory.Contract.VeloV2(&_PoolFactory.CallOpts)
}

// VolatileFee is a free data retrieval call binding the contract method 0x5084ed03.
//
// Solidity: function volatileFee() view returns(uint256)
func (_PoolFactory *PoolFactoryCaller) VolatileFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "volatileFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VolatileFee is a free data retrieval call binding the contract method 0x5084ed03.
//
// Solidity: function volatileFee() view returns(uint256)
func (_PoolFactory *PoolFactorySession) VolatileFee() (*big.Int, error) {
	return _PoolFactory.Contract.VolatileFee(&_PoolFactory.CallOpts)
}

// VolatileFee is a free data retrieval call binding the contract method 0x5084ed03.
//
// Solidity: function volatileFee() view returns(uint256)
func (_PoolFactory *PoolFactoryCallerSession) VolatileFee() (*big.Int, error) {
	return _PoolFactory.Contract.VolatileFee(&_PoolFactory.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_PoolFactory *PoolFactoryCaller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_PoolFactory *PoolFactorySession) Voter() (common.Address, error) {
	return _PoolFactory.Contract.Voter(&_PoolFactory.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) Voter() (common.Address, error) {
	return _PoolFactory.Contract.Voter(&_PoolFactory.CallOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0x82dfdce4.
//
// Solidity: function createPair(address tokenA, address tokenB, bool stable) returns(address pool)
func (_PoolFactory *PoolFactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "createPair", tokenA, tokenB, stable)
}

// CreatePair is a paid mutator transaction binding the contract method 0x82dfdce4.
//
// Solidity: function createPair(address tokenA, address tokenB, bool stable) returns(address pool)
func (_PoolFactory *PoolFactorySession) CreatePair(tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _PoolFactory.Contract.CreatePair(&_PoolFactory.TransactOpts, tokenA, tokenB, stable)
}

// CreatePair is a paid mutator transaction binding the contract method 0x82dfdce4.
//
// Solidity: function createPair(address tokenA, address tokenB, bool stable) returns(address pool)
func (_PoolFactory *PoolFactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _PoolFactory.Contract.CreatePair(&_PoolFactory.TransactOpts, tokenA, tokenB, stable)
}

// CreatePool is a paid mutator transaction binding the contract method 0x36bf95a0.
//
// Solidity: function createPool(address tokenA, address tokenB, bool stable) returns(address pool)
func (_PoolFactory *PoolFactoryTransactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "createPool", tokenA, tokenB, stable)
}

// CreatePool is a paid mutator transaction binding the contract method 0x36bf95a0.
//
// Solidity: function createPool(address tokenA, address tokenB, bool stable) returns(address pool)
func (_PoolFactory *PoolFactorySession) CreatePool(tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _PoolFactory.Contract.CreatePool(&_PoolFactory.TransactOpts, tokenA, tokenB, stable)
}

// CreatePool is a paid mutator transaction binding the contract method 0x36bf95a0.
//
// Solidity: function createPool(address tokenA, address tokenB, bool stable) returns(address pool)
func (_PoolFactory *PoolFactoryTransactorSession) CreatePool(tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _PoolFactory.Contract.CreatePool(&_PoolFactory.TransactOpts, tokenA, tokenB, stable)
}

// CreatePool0 is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_PoolFactory *PoolFactoryTransactor) CreatePool0(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "createPool0", tokenA, tokenB, fee)
}

// CreatePool0 is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_PoolFactory *PoolFactorySession) CreatePool0(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _PoolFactory.Contract.CreatePool0(&_PoolFactory.TransactOpts, tokenA, tokenB, fee)
}

// CreatePool0 is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_PoolFactory *PoolFactoryTransactorSession) CreatePool0(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _PoolFactory.Contract.CreatePool0(&_PoolFactory.TransactOpts, tokenA, tokenB, fee)
}

// SetCustomFee is a paid mutator transaction binding the contract method 0xd49466a8.
//
// Solidity: function setCustomFee(address pool, uint256 fee) returns()
func (_PoolFactory *PoolFactoryTransactor) SetCustomFee(opts *bind.TransactOpts, pool common.Address, fee *big.Int) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "setCustomFee", pool, fee)
}

// SetCustomFee is a paid mutator transaction binding the contract method 0xd49466a8.
//
// Solidity: function setCustomFee(address pool, uint256 fee) returns()
func (_PoolFactory *PoolFactorySession) SetCustomFee(pool common.Address, fee *big.Int) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetCustomFee(&_PoolFactory.TransactOpts, pool, fee)
}

// SetCustomFee is a paid mutator transaction binding the contract method 0xd49466a8.
//
// Solidity: function setCustomFee(address pool, uint256 fee) returns()
func (_PoolFactory *PoolFactoryTransactorSession) SetCustomFee(pool common.Address, fee *big.Int) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetCustomFee(&_PoolFactory.TransactOpts, pool, fee)
}

// SetFee is a paid mutator transaction binding the contract method 0xe1f76b44.
//
// Solidity: function setFee(bool _stable, uint256 _fee) returns()
func (_PoolFactory *PoolFactoryTransactor) SetFee(opts *bind.TransactOpts, _stable bool, _fee *big.Int) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "setFee", _stable, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0xe1f76b44.
//
// Solidity: function setFee(bool _stable, uint256 _fee) returns()
func (_PoolFactory *PoolFactorySession) SetFee(_stable bool, _fee *big.Int) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetFee(&_PoolFactory.TransactOpts, _stable, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0xe1f76b44.
//
// Solidity: function setFee(bool _stable, uint256 _fee) returns()
func (_PoolFactory *PoolFactoryTransactorSession) SetFee(_stable bool, _fee *big.Int) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetFee(&_PoolFactory.TransactOpts, _stable, _fee)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_PoolFactory *PoolFactoryTransactor) SetFeeManager(opts *bind.TransactOpts, _feeManager common.Address) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "setFeeManager", _feeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_PoolFactory *PoolFactorySession) SetFeeManager(_feeManager common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetFeeManager(&_PoolFactory.TransactOpts, _feeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_PoolFactory *PoolFactoryTransactorSession) SetFeeManager(_feeManager common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetFeeManager(&_PoolFactory.TransactOpts, _feeManager)
}

// SetPauseState is a paid mutator transaction binding the contract method 0xcdb88ad1.
//
// Solidity: function setPauseState(bool _state) returns()
func (_PoolFactory *PoolFactoryTransactor) SetPauseState(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "setPauseState", _state)
}

// SetPauseState is a paid mutator transaction binding the contract method 0xcdb88ad1.
//
// Solidity: function setPauseState(bool _state) returns()
func (_PoolFactory *PoolFactorySession) SetPauseState(_state bool) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetPauseState(&_PoolFactory.TransactOpts, _state)
}

// SetPauseState is a paid mutator transaction binding the contract method 0xcdb88ad1.
//
// Solidity: function setPauseState(bool _state) returns()
func (_PoolFactory *PoolFactoryTransactorSession) SetPauseState(_state bool) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetPauseState(&_PoolFactory.TransactOpts, _state)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_PoolFactory *PoolFactoryTransactor) SetPauser(opts *bind.TransactOpts, _pauser common.Address) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "setPauser", _pauser)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_PoolFactory *PoolFactorySession) SetPauser(_pauser common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetPauser(&_PoolFactory.TransactOpts, _pauser)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_PoolFactory *PoolFactoryTransactorSession) SetPauser(_pauser common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetPauser(&_PoolFactory.TransactOpts, _pauser)
}

// SetSinkConverter is a paid mutator transaction binding the contract method 0x37068e7b.
//
// Solidity: function setSinkConverter(address _sinkConverter, address _velo, address _veloV2) returns()
func (_PoolFactory *PoolFactoryTransactor) SetSinkConverter(opts *bind.TransactOpts, _sinkConverter common.Address, _velo common.Address, _veloV2 common.Address) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "setSinkConverter", _sinkConverter, _velo, _veloV2)
}

// SetSinkConverter is a paid mutator transaction binding the contract method 0x37068e7b.
//
// Solidity: function setSinkConverter(address _sinkConverter, address _velo, address _veloV2) returns()
func (_PoolFactory *PoolFactorySession) SetSinkConverter(_sinkConverter common.Address, _velo common.Address, _veloV2 common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetSinkConverter(&_PoolFactory.TransactOpts, _sinkConverter, _velo, _veloV2)
}

// SetSinkConverter is a paid mutator transaction binding the contract method 0x37068e7b.
//
// Solidity: function setSinkConverter(address _sinkConverter, address _velo, address _veloV2) returns()
func (_PoolFactory *PoolFactoryTransactorSession) SetSinkConverter(_sinkConverter common.Address, _velo common.Address, _veloV2 common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetSinkConverter(&_PoolFactory.TransactOpts, _sinkConverter, _velo, _veloV2)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _voter) returns()
func (_PoolFactory *PoolFactoryTransactor) SetVoter(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "setVoter", _voter)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _voter) returns()
func (_PoolFactory *PoolFactorySession) SetVoter(_voter common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetVoter(&_PoolFactory.TransactOpts, _voter)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _voter) returns()
func (_PoolFactory *PoolFactoryTransactorSession) SetVoter(_voter common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetVoter(&_PoolFactory.TransactOpts, _voter)
}

// PoolFactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the PoolFactory contract.
type PoolFactoryPoolCreatedIterator struct {
	Event *PoolFactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *PoolFactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFactoryPoolCreated)
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
		it.Event = new(PoolFactoryPoolCreated)
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
func (it *PoolFactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFactoryPoolCreated represents a PoolCreated event raised by the PoolFactory contract.
type PoolFactoryPoolCreated struct {
	Token0 common.Address
	Token1 common.Address
	Stable bool
	Pool   common.Address
	Arg4   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x2128d88d14c80cb081c1252a5acff7a264671bf199ce226b53788fb26065005e.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, bool indexed stable, address pool, uint256 arg4)
func (_PoolFactory *PoolFactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address, stable []bool) (*PoolFactoryPoolCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var stableRule []interface{}
	for _, stableItem := range stable {
		stableRule = append(stableRule, stableItem)
	}

	logs, sub, err := _PoolFactory.contract.FilterLogs(opts, "PoolCreated", token0Rule, token1Rule, stableRule)
	if err != nil {
		return nil, err
	}
	return &PoolFactoryPoolCreatedIterator{contract: _PoolFactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x2128d88d14c80cb081c1252a5acff7a264671bf199ce226b53788fb26065005e.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, bool indexed stable, address pool, uint256 arg4)
func (_PoolFactory *PoolFactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *PoolFactoryPoolCreated, token0 []common.Address, token1 []common.Address, stable []bool) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var stableRule []interface{}
	for _, stableItem := range stable {
		stableRule = append(stableRule, stableItem)
	}

	logs, sub, err := _PoolFactory.contract.WatchLogs(opts, "PoolCreated", token0Rule, token1Rule, stableRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFactoryPoolCreated)
				if err := _PoolFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x2128d88d14c80cb081c1252a5acff7a264671bf199ce226b53788fb26065005e.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, bool indexed stable, address pool, uint256 arg4)
func (_PoolFactory *PoolFactoryFilterer) ParsePoolCreated(log types.Log) (*PoolFactoryPoolCreated, error) {
	event := new(PoolFactoryPoolCreated)
	if err := _PoolFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFactorySetCustomFeeIterator is returned from FilterSetCustomFee and is used to iterate over the raw logs and unpacked data for SetCustomFee events raised by the PoolFactory contract.
type PoolFactorySetCustomFeeIterator struct {
	Event *PoolFactorySetCustomFee // Event containing the contract specifics and raw log

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
func (it *PoolFactorySetCustomFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFactorySetCustomFee)
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
		it.Event = new(PoolFactorySetCustomFee)
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
func (it *PoolFactorySetCustomFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFactorySetCustomFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFactorySetCustomFee represents a SetCustomFee event raised by the PoolFactory contract.
type PoolFactorySetCustomFee struct {
	Pool common.Address
	Fee  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetCustomFee is a free log retrieval operation binding the contract event 0xae468ce586f9a87660fdffc1448cee942042c16ae2f02046b134b5224f31936b.
//
// Solidity: event SetCustomFee(address indexed pool, uint256 fee)
func (_PoolFactory *PoolFactoryFilterer) FilterSetCustomFee(opts *bind.FilterOpts, pool []common.Address) (*PoolFactorySetCustomFeeIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _PoolFactory.contract.FilterLogs(opts, "SetCustomFee", poolRule)
	if err != nil {
		return nil, err
	}
	return &PoolFactorySetCustomFeeIterator{contract: _PoolFactory.contract, event: "SetCustomFee", logs: logs, sub: sub}, nil
}

// WatchSetCustomFee is a free log subscription operation binding the contract event 0xae468ce586f9a87660fdffc1448cee942042c16ae2f02046b134b5224f31936b.
//
// Solidity: event SetCustomFee(address indexed pool, uint256 fee)
func (_PoolFactory *PoolFactoryFilterer) WatchSetCustomFee(opts *bind.WatchOpts, sink chan<- *PoolFactorySetCustomFee, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _PoolFactory.contract.WatchLogs(opts, "SetCustomFee", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFactorySetCustomFee)
				if err := _PoolFactory.contract.UnpackLog(event, "SetCustomFee", log); err != nil {
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

// ParseSetCustomFee is a log parse operation binding the contract event 0xae468ce586f9a87660fdffc1448cee942042c16ae2f02046b134b5224f31936b.
//
// Solidity: event SetCustomFee(address indexed pool, uint256 fee)
func (_PoolFactory *PoolFactoryFilterer) ParseSetCustomFee(log types.Log) (*PoolFactorySetCustomFee, error) {
	event := new(PoolFactorySetCustomFee)
	if err := _PoolFactory.contract.UnpackLog(event, "SetCustomFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFactorySetFeeManagerIterator is returned from FilterSetFeeManager and is used to iterate over the raw logs and unpacked data for SetFeeManager events raised by the PoolFactory contract.
type PoolFactorySetFeeManagerIterator struct {
	Event *PoolFactorySetFeeManager // Event containing the contract specifics and raw log

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
func (it *PoolFactorySetFeeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFactorySetFeeManager)
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
		it.Event = new(PoolFactorySetFeeManager)
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
func (it *PoolFactorySetFeeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFactorySetFeeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFactorySetFeeManager represents a SetFeeManager event raised by the PoolFactory contract.
type PoolFactorySetFeeManager struct {
	FeeManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetFeeManager is a free log retrieval operation binding the contract event 0x5d0517e3a4eabea892d9750138cd21d4a6cf3b935b43d0598df7055f463819b2.
//
// Solidity: event SetFeeManager(address feeManager)
func (_PoolFactory *PoolFactoryFilterer) FilterSetFeeManager(opts *bind.FilterOpts) (*PoolFactorySetFeeManagerIterator, error) {

	logs, sub, err := _PoolFactory.contract.FilterLogs(opts, "SetFeeManager")
	if err != nil {
		return nil, err
	}
	return &PoolFactorySetFeeManagerIterator{contract: _PoolFactory.contract, event: "SetFeeManager", logs: logs, sub: sub}, nil
}

// WatchSetFeeManager is a free log subscription operation binding the contract event 0x5d0517e3a4eabea892d9750138cd21d4a6cf3b935b43d0598df7055f463819b2.
//
// Solidity: event SetFeeManager(address feeManager)
func (_PoolFactory *PoolFactoryFilterer) WatchSetFeeManager(opts *bind.WatchOpts, sink chan<- *PoolFactorySetFeeManager) (event.Subscription, error) {

	logs, sub, err := _PoolFactory.contract.WatchLogs(opts, "SetFeeManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFactorySetFeeManager)
				if err := _PoolFactory.contract.UnpackLog(event, "SetFeeManager", log); err != nil {
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

// ParseSetFeeManager is a log parse operation binding the contract event 0x5d0517e3a4eabea892d9750138cd21d4a6cf3b935b43d0598df7055f463819b2.
//
// Solidity: event SetFeeManager(address feeManager)
func (_PoolFactory *PoolFactoryFilterer) ParseSetFeeManager(log types.Log) (*PoolFactorySetFeeManager, error) {
	event := new(PoolFactorySetFeeManager)
	if err := _PoolFactory.contract.UnpackLog(event, "SetFeeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFactorySetPauseStateIterator is returned from FilterSetPauseState and is used to iterate over the raw logs and unpacked data for SetPauseState events raised by the PoolFactory contract.
type PoolFactorySetPauseStateIterator struct {
	Event *PoolFactorySetPauseState // Event containing the contract specifics and raw log

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
func (it *PoolFactorySetPauseStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFactorySetPauseState)
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
		it.Event = new(PoolFactorySetPauseState)
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
func (it *PoolFactorySetPauseStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFactorySetPauseStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFactorySetPauseState represents a SetPauseState event raised by the PoolFactory contract.
type PoolFactorySetPauseState struct {
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetPauseState is a free log retrieval operation binding the contract event 0x0d76538efc408318a051137c2720a9e82902acdbd46b802d488b74ca3a09a116.
//
// Solidity: event SetPauseState(bool state)
func (_PoolFactory *PoolFactoryFilterer) FilterSetPauseState(opts *bind.FilterOpts) (*PoolFactorySetPauseStateIterator, error) {

	logs, sub, err := _PoolFactory.contract.FilterLogs(opts, "SetPauseState")
	if err != nil {
		return nil, err
	}
	return &PoolFactorySetPauseStateIterator{contract: _PoolFactory.contract, event: "SetPauseState", logs: logs, sub: sub}, nil
}

// WatchSetPauseState is a free log subscription operation binding the contract event 0x0d76538efc408318a051137c2720a9e82902acdbd46b802d488b74ca3a09a116.
//
// Solidity: event SetPauseState(bool state)
func (_PoolFactory *PoolFactoryFilterer) WatchSetPauseState(opts *bind.WatchOpts, sink chan<- *PoolFactorySetPauseState) (event.Subscription, error) {

	logs, sub, err := _PoolFactory.contract.WatchLogs(opts, "SetPauseState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFactorySetPauseState)
				if err := _PoolFactory.contract.UnpackLog(event, "SetPauseState", log); err != nil {
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

// ParseSetPauseState is a log parse operation binding the contract event 0x0d76538efc408318a051137c2720a9e82902acdbd46b802d488b74ca3a09a116.
//
// Solidity: event SetPauseState(bool state)
func (_PoolFactory *PoolFactoryFilterer) ParseSetPauseState(log types.Log) (*PoolFactorySetPauseState, error) {
	event := new(PoolFactorySetPauseState)
	if err := _PoolFactory.contract.UnpackLog(event, "SetPauseState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFactorySetPauserIterator is returned from FilterSetPauser and is used to iterate over the raw logs and unpacked data for SetPauser events raised by the PoolFactory contract.
type PoolFactorySetPauserIterator struct {
	Event *PoolFactorySetPauser // Event containing the contract specifics and raw log

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
func (it *PoolFactorySetPauserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFactorySetPauser)
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
		it.Event = new(PoolFactorySetPauser)
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
func (it *PoolFactorySetPauserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFactorySetPauserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFactorySetPauser represents a SetPauser event raised by the PoolFactory contract.
type PoolFactorySetPauser struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetPauser is a free log retrieval operation binding the contract event 0xe02efb9e8f0fc21546730ab32d594f62d586e1bbb15bb5045edd0b1878a77b35.
//
// Solidity: event SetPauser(address pauser)
func (_PoolFactory *PoolFactoryFilterer) FilterSetPauser(opts *bind.FilterOpts) (*PoolFactorySetPauserIterator, error) {

	logs, sub, err := _PoolFactory.contract.FilterLogs(opts, "SetPauser")
	if err != nil {
		return nil, err
	}
	return &PoolFactorySetPauserIterator{contract: _PoolFactory.contract, event: "SetPauser", logs: logs, sub: sub}, nil
}

// WatchSetPauser is a free log subscription operation binding the contract event 0xe02efb9e8f0fc21546730ab32d594f62d586e1bbb15bb5045edd0b1878a77b35.
//
// Solidity: event SetPauser(address pauser)
func (_PoolFactory *PoolFactoryFilterer) WatchSetPauser(opts *bind.WatchOpts, sink chan<- *PoolFactorySetPauser) (event.Subscription, error) {

	logs, sub, err := _PoolFactory.contract.WatchLogs(opts, "SetPauser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFactorySetPauser)
				if err := _PoolFactory.contract.UnpackLog(event, "SetPauser", log); err != nil {
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

// ParseSetPauser is a log parse operation binding the contract event 0xe02efb9e8f0fc21546730ab32d594f62d586e1bbb15bb5045edd0b1878a77b35.
//
// Solidity: event SetPauser(address pauser)
func (_PoolFactory *PoolFactoryFilterer) ParseSetPauser(log types.Log) (*PoolFactorySetPauser, error) {
	event := new(PoolFactorySetPauser)
	if err := _PoolFactory.contract.UnpackLog(event, "SetPauser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFactorySetVoterIterator is returned from FilterSetVoter and is used to iterate over the raw logs and unpacked data for SetVoter events raised by the PoolFactory contract.
type PoolFactorySetVoterIterator struct {
	Event *PoolFactorySetVoter // Event containing the contract specifics and raw log

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
func (it *PoolFactorySetVoterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFactorySetVoter)
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
		it.Event = new(PoolFactorySetVoter)
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
func (it *PoolFactorySetVoterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFactorySetVoterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFactorySetVoter represents a SetVoter event raised by the PoolFactory contract.
type PoolFactorySetVoter struct {
	Voter common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetVoter is a free log retrieval operation binding the contract event 0xc6ff127433b785c51da9ae4088ee184c909b1a55b9afd82ae6c64224d3bc15d2.
//
// Solidity: event SetVoter(address voter)
func (_PoolFactory *PoolFactoryFilterer) FilterSetVoter(opts *bind.FilterOpts) (*PoolFactorySetVoterIterator, error) {

	logs, sub, err := _PoolFactory.contract.FilterLogs(opts, "SetVoter")
	if err != nil {
		return nil, err
	}
	return &PoolFactorySetVoterIterator{contract: _PoolFactory.contract, event: "SetVoter", logs: logs, sub: sub}, nil
}

// WatchSetVoter is a free log subscription operation binding the contract event 0xc6ff127433b785c51da9ae4088ee184c909b1a55b9afd82ae6c64224d3bc15d2.
//
// Solidity: event SetVoter(address voter)
func (_PoolFactory *PoolFactoryFilterer) WatchSetVoter(opts *bind.WatchOpts, sink chan<- *PoolFactorySetVoter) (event.Subscription, error) {

	logs, sub, err := _PoolFactory.contract.WatchLogs(opts, "SetVoter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFactorySetVoter)
				if err := _PoolFactory.contract.UnpackLog(event, "SetVoter", log); err != nil {
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

// ParseSetVoter is a log parse operation binding the contract event 0xc6ff127433b785c51da9ae4088ee184c909b1a55b9afd82ae6c64224d3bc15d2.
//
// Solidity: event SetVoter(address voter)
func (_PoolFactory *PoolFactoryFilterer) ParseSetVoter(log types.Log) (*PoolFactorySetVoter, error) {
	event := new(PoolFactorySetVoter)
	if err := _PoolFactory.contract.UnpackLog(event, "SetVoter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
