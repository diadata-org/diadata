// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package platform

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

// PlatformABI is the input ABI used to generate the binding from.
const PlatformABI = "[{\"name\":\"TokenExchange\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"sold_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"bought_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchangeUnderlying\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"sold_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"bought_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[4]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[4]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[4]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[4]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[4]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[4]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true,\"unit\":\"sec\"},{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewParameters\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true,\"unit\":\"sec\"},{\"type\":\"uint256\",\"name\":\"A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewParameters\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address[4]\",\"name\":\"_coins\"},{\"type\":\"address[4]\",\"name\":\"_underlying_coins\"},{\"type\":\"address\",\"name\":\"_pool_token\"},{\"type\":\"uint256\",\"name\":\"_A\"},{\"type\":\"uint256\",\"name\":\"_fee\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"get_virtual_price\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1570535},{\"name\":\"calc_token_amount\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256[4]\",\"name\":\"amounts\"},{\"type\":\"bool\",\"name\":\"deposit\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6103471},{\"name\":\"add_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[4]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"min_mint_amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":9331701},{\"name\":\"get_dy\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":3489637},{\"name\":\"get_dy_underlying\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":3489467},{\"name\":\"exchange\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"},{\"type\":\"uint256\",\"name\":\"min_dy\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":7034253},{\"name\":\"exchange_underlying\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"},{\"type\":\"uint256\",\"name\":\"min_dy\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":7050488},{\"name\":\"remove_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"},{\"type\":\"uint256[4]\",\"name\":\"min_amounts\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":241191},{\"name\":\"remove_liquidity_imbalance\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[4]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"max_burn_amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":9330864},{\"name\":\"commit_new_parameters\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"amplification\"},{\"type\":\"uint256\",\"name\":\"new_fee\"},{\"type\":\"uint256\",\"name\":\"new_admin_fee\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":146045},{\"name\":\"apply_new_parameters\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":133452},{\"name\":\"revert_new_parameters\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":21775},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":74452},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":60508},{\"name\":\"revert_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":21865},{\"name\":\"withdraw_admin_fees\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":23448},{\"name\":\"kill_me\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":37818},{\"name\":\"unkill_me\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":21955},{\"name\":\"coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2130},{\"name\":\"underlying_coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2160},{\"name\":\"balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2190},{\"name\":\"A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2021},{\"name\":\"fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2051},{\"name\":\"admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2081},{\"name\":\"owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2111},{\"name\":\"admin_actions_deadline\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"sec\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2141},{\"name\":\"transfer_ownership_deadline\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"sec\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2171},{\"name\":\"future_A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2201},{\"name\":\"future_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2231},{\"name\":\"future_admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2261},{\"name\":\"future_owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2291}]"

// Platform is an auto generated Go binding around an Ethereum contract.
type Platform struct {
	PlatformCaller     // Read-only binding to the contract
	PlatformTransactor // Write-only binding to the contract
	PlatformFilterer   // Log filterer for contract events
}

// PlatformCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlatformCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatformTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlatformTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatformFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlatformFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatformSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlatformSession struct {
	Contract     *Platform         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlatformCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlatformCallerSession struct {
	Contract *PlatformCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PlatformTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlatformTransactorSession struct {
	Contract     *PlatformTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PlatformRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlatformRaw struct {
	Contract *Platform // Generic contract binding to access the raw methods on
}

// PlatformCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlatformCallerRaw struct {
	Contract *PlatformCaller // Generic read-only contract binding to access the raw methods on
}

// PlatformTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlatformTransactorRaw struct {
	Contract *PlatformTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlatform creates a new instance of Platform, bound to a specific deployed contract.
func NewPlatform(address common.Address, backend bind.ContractBackend) (*Platform, error) {
	contract, err := bindPlatform(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Platform{PlatformCaller: PlatformCaller{contract: contract}, PlatformTransactor: PlatformTransactor{contract: contract}, PlatformFilterer: PlatformFilterer{contract: contract}}, nil
}

// NewPlatformCaller creates a new read-only instance of Platform, bound to a specific deployed contract.
func NewPlatformCaller(address common.Address, caller bind.ContractCaller) (*PlatformCaller, error) {
	contract, err := bindPlatform(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlatformCaller{contract: contract}, nil
}

// NewPlatformTransactor creates a new write-only instance of Platform, bound to a specific deployed contract.
func NewPlatformTransactor(address common.Address, transactor bind.ContractTransactor) (*PlatformTransactor, error) {
	contract, err := bindPlatform(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlatformTransactor{contract: contract}, nil
}

// NewPlatformFilterer creates a new log filterer instance of Platform, bound to a specific deployed contract.
func NewPlatformFilterer(address common.Address, filterer bind.ContractFilterer) (*PlatformFilterer, error) {
	contract, err := bindPlatform(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlatformFilterer{contract: contract}, nil
}

// bindPlatform binds a generic wrapper to an already deployed contract.
func bindPlatform(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlatformABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Platform *PlatformRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Platform.Contract.PlatformCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Platform *PlatformRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Platform.Contract.PlatformTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Platform *PlatformRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Platform.Contract.PlatformTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Platform *PlatformCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Platform.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Platform *PlatformTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Platform.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Platform *PlatformTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Platform.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() returns(uint256)
func (_Platform *PlatformCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "A")
	return *ret0, err
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() returns(uint256)
func (_Platform *PlatformSession) A() (*big.Int, error) {
	return _Platform.Contract.A(&_Platform.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() returns(uint256)
func (_Platform *PlatformCallerSession) A() (*big.Int, error) {
	return _Platform.Contract.A(&_Platform.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() returns(uint256)
func (_Platform *PlatformCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "admin_actions_deadline")
	return *ret0, err
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() returns(uint256)
func (_Platform *PlatformSession) AdminActionsDeadline() (*big.Int, error) {
	return _Platform.Contract.AdminActionsDeadline(&_Platform.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() returns(uint256)
func (_Platform *PlatformCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _Platform.Contract.AdminActionsDeadline(&_Platform.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() returns(uint256)
func (_Platform *PlatformCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "admin_fee")
	return *ret0, err
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() returns(uint256)
func (_Platform *PlatformSession) AdminFee() (*big.Int, error) {
	return _Platform.Contract.AdminFee(&_Platform.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() returns(uint256)
func (_Platform *PlatformCallerSession) AdminFee() (*big.Int, error) {
	return _Platform.Contract.AdminFee(&_Platform.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 arg0) returns(uint256)
func (_Platform *PlatformCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 arg0) returns(uint256)
func (_Platform *PlatformSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Platform.Contract.Balances(&_Platform.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 arg0) returns(uint256)
func (_Platform *PlatformCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Platform.Contract.Balances(&_Platform.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xcf701ff7.
//
// Solidity: function calc_token_amount(uint256[4] amounts, bool deposit) returns(uint256)
func (_Platform *PlatformCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [4]*big.Int, deposit bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "calc_token_amount", amounts, deposit)
	return *ret0, err
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xcf701ff7.
//
// Solidity: function calc_token_amount(uint256[4] amounts, bool deposit) returns(uint256)
func (_Platform *PlatformSession) CalcTokenAmount(amounts [4]*big.Int, deposit bool) (*big.Int, error) {
	return _Platform.Contract.CalcTokenAmount(&_Platform.CallOpts, amounts, deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xcf701ff7.
//
// Solidity: function calc_token_amount(uint256[4] amounts, bool deposit) returns(uint256)
func (_Platform *PlatformCallerSession) CalcTokenAmount(amounts [4]*big.Int, deposit bool) (*big.Int, error) {
	return _Platform.Contract.CalcTokenAmount(&_Platform.CallOpts, amounts, deposit)
}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 arg0) returns(address)
func (_Platform *PlatformCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "coins", arg0)
	return *ret0, err
}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 arg0) returns(address)
func (_Platform *PlatformSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Platform.Contract.Coins(&_Platform.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 arg0) returns(address)
func (_Platform *PlatformCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Platform.Contract.Coins(&_Platform.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() returns(uint256)
func (_Platform *PlatformCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() returns(uint256)
func (_Platform *PlatformSession) Fee() (*big.Int, error) {
	return _Platform.Contract.Fee(&_Platform.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() returns(uint256)
func (_Platform *PlatformCallerSession) Fee() (*big.Int, error) {
	return _Platform.Contract.Fee(&_Platform.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() returns(uint256)
func (_Platform *PlatformCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "future_A")
	return *ret0, err
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() returns(uint256)
func (_Platform *PlatformSession) FutureA() (*big.Int, error) {
	return _Platform.Contract.FutureA(&_Platform.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() returns(uint256)
func (_Platform *PlatformCallerSession) FutureA() (*big.Int, error) {
	return _Platform.Contract.FutureA(&_Platform.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() returns(uint256)
func (_Platform *PlatformCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "future_admin_fee")
	return *ret0, err
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() returns(uint256)
func (_Platform *PlatformSession) FutureAdminFee() (*big.Int, error) {
	return _Platform.Contract.FutureAdminFee(&_Platform.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() returns(uint256)
func (_Platform *PlatformCallerSession) FutureAdminFee() (*big.Int, error) {
	return _Platform.Contract.FutureAdminFee(&_Platform.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() returns(uint256)
func (_Platform *PlatformCaller) FutureFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "future_fee")
	return *ret0, err
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() returns(uint256)
func (_Platform *PlatformSession) FutureFee() (*big.Int, error) {
	return _Platform.Contract.FutureFee(&_Platform.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() returns(uint256)
func (_Platform *PlatformCallerSession) FutureFee() (*big.Int, error) {
	return _Platform.Contract.FutureFee(&_Platform.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() returns(address)
func (_Platform *PlatformCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "future_owner")
	return *ret0, err
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() returns(address)
func (_Platform *PlatformSession) FutureOwner() (common.Address, error) {
	return _Platform.Contract.FutureOwner(&_Platform.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() returns(address)
func (_Platform *PlatformCallerSession) FutureOwner() (common.Address, error) {
	return _Platform.Contract.FutureOwner(&_Platform.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Platform *PlatformCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "get_dy", i, j, dx)
	return *ret0, err
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Platform *PlatformSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Platform.Contract.GetDy(&_Platform.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Platform *PlatformCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Platform.Contract.GetDy(&_Platform.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Platform *PlatformCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "get_dy_underlying", i, j, dx)
	return *ret0, err
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Platform *PlatformSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Platform.Contract.GetDyUnderlying(&_Platform.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Platform *PlatformCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Platform.Contract.GetDyUnderlying(&_Platform.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() returns(uint256)
func (_Platform *PlatformCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "get_virtual_price")
	return *ret0, err
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() returns(uint256)
func (_Platform *PlatformSession) GetVirtualPrice() (*big.Int, error) {
	return _Platform.Contract.GetVirtualPrice(&_Platform.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() returns(uint256)
func (_Platform *PlatformCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _Platform.Contract.GetVirtualPrice(&_Platform.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Platform *PlatformCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Platform *PlatformSession) Owner() (common.Address, error) {
	return _Platform.Contract.Owner(&_Platform.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Platform *PlatformCallerSession) Owner() (common.Address, error) {
	return _Platform.Contract.Owner(&_Platform.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() returns(uint256)
func (_Platform *PlatformCaller) TransferOwnershipDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "transfer_ownership_deadline")
	return *ret0, err
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() returns(uint256)
func (_Platform *PlatformSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _Platform.Contract.TransferOwnershipDeadline(&_Platform.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() returns(uint256)
func (_Platform *PlatformCallerSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _Platform.Contract.TransferOwnershipDeadline(&_Platform.CallOpts)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 arg0) returns(address)
func (_Platform *PlatformCaller) UnderlyingCoins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Platform.contract.Call(opts, out, "underlying_coins", arg0)
	return *ret0, err
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 arg0) returns(address)
func (_Platform *PlatformSession) UnderlyingCoins(arg0 *big.Int) (common.Address, error) {
	return _Platform.Contract.UnderlyingCoins(&_Platform.CallOpts, arg0)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 arg0) returns(address)
func (_Platform *PlatformCallerSession) UnderlyingCoins(arg0 *big.Int) (common.Address, error) {
	return _Platform.Contract.UnderlyingCoins(&_Platform.CallOpts, arg0)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x029b2f34.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 min_mint_amount) returns()
func (_Platform *PlatformTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [4]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x029b2f34.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 min_mint_amount) returns()
func (_Platform *PlatformSession) AddLiquidity(amounts [4]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Platform.Contract.AddLiquidity(&_Platform.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x029b2f34.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 min_mint_amount) returns()
func (_Platform *PlatformTransactorSession) AddLiquidity(amounts [4]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Platform.Contract.AddLiquidity(&_Platform.TransactOpts, amounts, min_mint_amount)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Platform *PlatformTransactor) ApplyNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "apply_new_parameters")
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Platform *PlatformSession) ApplyNewParameters() (*types.Transaction, error) {
	return _Platform.Contract.ApplyNewParameters(&_Platform.TransactOpts)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Platform *PlatformTransactorSession) ApplyNewParameters() (*types.Transaction, error) {
	return _Platform.Contract.ApplyNewParameters(&_Platform.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Platform *PlatformTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Platform *PlatformSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Platform.Contract.ApplyTransferOwnership(&_Platform.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Platform *PlatformTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Platform.Contract.ApplyTransferOwnership(&_Platform.TransactOpts)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xee11f5b6.
//
// Solidity: function commit_new_parameters(uint256 amplification, uint256 new_fee, uint256 new_admin_fee) returns()
func (_Platform *PlatformTransactor) CommitNewParameters(opts *bind.TransactOpts, amplification *big.Int, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "commit_new_parameters", amplification, new_fee, new_admin_fee)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xee11f5b6.
//
// Solidity: function commit_new_parameters(uint256 amplification, uint256 new_fee, uint256 new_admin_fee) returns()
func (_Platform *PlatformSession) CommitNewParameters(amplification *big.Int, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Platform.Contract.CommitNewParameters(&_Platform.TransactOpts, amplification, new_fee, new_admin_fee)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xee11f5b6.
//
// Solidity: function commit_new_parameters(uint256 amplification, uint256 new_fee, uint256 new_admin_fee) returns()
func (_Platform *PlatformTransactorSession) CommitNewParameters(amplification *big.Int, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Platform.Contract.CommitNewParameters(&_Platform.TransactOpts, amplification, new_fee, new_admin_fee)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Platform *PlatformTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Platform *PlatformSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Platform.Contract.CommitTransferOwnership(&_Platform.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Platform *PlatformTransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Platform.Contract.CommitTransferOwnership(&_Platform.TransactOpts, _owner)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Platform *PlatformTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Platform *PlatformSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Platform.Contract.Exchange(&_Platform.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Platform *PlatformTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Platform.Contract.Exchange(&_Platform.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Platform *PlatformTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "exchange_underlying", i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Platform *PlatformSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Platform.Contract.ExchangeUnderlying(&_Platform.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Platform *PlatformTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Platform.Contract.ExchangeUnderlying(&_Platform.TransactOpts, i, j, dx, min_dy)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Platform *PlatformTransactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Platform *PlatformSession) KillMe() (*types.Transaction, error) {
	return _Platform.Contract.KillMe(&_Platform.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Platform *PlatformTransactorSession) KillMe() (*types.Transaction, error) {
	return _Platform.Contract.KillMe(&_Platform.TransactOpts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x7d49d875.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[4] min_amounts) returns()
func (_Platform *PlatformTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [4]*big.Int) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x7d49d875.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[4] min_amounts) returns()
func (_Platform *PlatformSession) RemoveLiquidity(_amount *big.Int, min_amounts [4]*big.Int) (*types.Transaction, error) {
	return _Platform.Contract.RemoveLiquidity(&_Platform.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x7d49d875.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[4] min_amounts) returns()
func (_Platform *PlatformTransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [4]*big.Int) (*types.Transaction, error) {
	return _Platform.Contract.RemoveLiquidity(&_Platform.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x18a7bd76.
//
// Solidity: function remove_liquidity_imbalance(uint256[4] amounts, uint256 max_burn_amount) returns()
func (_Platform *PlatformTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, amounts [4]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "remove_liquidity_imbalance", amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x18a7bd76.
//
// Solidity: function remove_liquidity_imbalance(uint256[4] amounts, uint256 max_burn_amount) returns()
func (_Platform *PlatformSession) RemoveLiquidityImbalance(amounts [4]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Platform.Contract.RemoveLiquidityImbalance(&_Platform.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x18a7bd76.
//
// Solidity: function remove_liquidity_imbalance(uint256[4] amounts, uint256 max_burn_amount) returns()
func (_Platform *PlatformTransactorSession) RemoveLiquidityImbalance(amounts [4]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Platform.Contract.RemoveLiquidityImbalance(&_Platform.TransactOpts, amounts, max_burn_amount)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Platform *PlatformTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Platform *PlatformSession) RevertNewParameters() (*types.Transaction, error) {
	return _Platform.Contract.RevertNewParameters(&_Platform.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Platform *PlatformTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _Platform.Contract.RevertNewParameters(&_Platform.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Platform *PlatformTransactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Platform *PlatformSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Platform.Contract.RevertTransferOwnership(&_Platform.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Platform *PlatformTransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Platform.Contract.RevertTransferOwnership(&_Platform.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Platform *PlatformTransactor) UnkillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "unkill_me")
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Platform *PlatformSession) UnkillMe() (*types.Transaction, error) {
	return _Platform.Contract.UnkillMe(&_Platform.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Platform *PlatformTransactorSession) UnkillMe() (*types.Transaction, error) {
	return _Platform.Contract.UnkillMe(&_Platform.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Platform *PlatformTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Platform *PlatformSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Platform.Contract.WithdrawAdminFees(&_Platform.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Platform *PlatformTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Platform.Contract.WithdrawAdminFees(&_Platform.TransactOpts)
}

// PlatformAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Platform contract.
type PlatformAddLiquidityIterator struct {
	Event *PlatformAddLiquidity // Event containing the contract specifics and raw log

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
func (it *PlatformAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlatformAddLiquidity)
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
		it.Event = new(PlatformAddLiquidity)
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
func (it *PlatformAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlatformAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlatformAddLiquidity represents a AddLiquidity event raised by the Platform contract.
type PlatformAddLiquidity struct {
	Provider     common.Address
	TokenAmounts [4]*big.Int
	Fees         [4]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x3f1915775e0c9a38a57a7bb7f1f9005f486fb904e1f84aa215364d567319a58d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_Platform *PlatformFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*PlatformAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Platform.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &PlatformAddLiquidityIterator{contract: _Platform.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x3f1915775e0c9a38a57a7bb7f1f9005f486fb904e1f84aa215364d567319a58d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_Platform *PlatformFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *PlatformAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Platform.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlatformAddLiquidity)
				if err := _Platform.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x3f1915775e0c9a38a57a7bb7f1f9005f486fb904e1f84aa215364d567319a58d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_Platform *PlatformFilterer) ParseAddLiquidity(log types.Log) (*PlatformAddLiquidity, error) {
	event := new(PlatformAddLiquidity)
	if err := _Platform.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PlatformCommitNewAdminIterator is returned from FilterCommitNewAdmin and is used to iterate over the raw logs and unpacked data for CommitNewAdmin events raised by the Platform contract.
type PlatformCommitNewAdminIterator struct {
	Event *PlatformCommitNewAdmin // Event containing the contract specifics and raw log

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
func (it *PlatformCommitNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlatformCommitNewAdmin)
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
		it.Event = new(PlatformCommitNewAdmin)
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
func (it *PlatformCommitNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlatformCommitNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlatformCommitNewAdmin represents a CommitNewAdmin event raised by the Platform contract.
type PlatformCommitNewAdmin struct {
	Deadline *big.Int
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewAdmin is a free log retrieval operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Platform *PlatformFilterer) FilterCommitNewAdmin(opts *bind.FilterOpts, deadline []*big.Int, admin []common.Address) (*PlatformCommitNewAdminIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Platform.contract.FilterLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &PlatformCommitNewAdminIterator{contract: _Platform.contract, event: "CommitNewAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitNewAdmin is a free log subscription operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Platform *PlatformFilterer) WatchCommitNewAdmin(opts *bind.WatchOpts, sink chan<- *PlatformCommitNewAdmin, deadline []*big.Int, admin []common.Address) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Platform.contract.WatchLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlatformCommitNewAdmin)
				if err := _Platform.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
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

// ParseCommitNewAdmin is a log parse operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Platform *PlatformFilterer) ParseCommitNewAdmin(log types.Log) (*PlatformCommitNewAdmin, error) {
	event := new(PlatformCommitNewAdmin)
	if err := _Platform.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PlatformCommitNewParametersIterator is returned from FilterCommitNewParameters and is used to iterate over the raw logs and unpacked data for CommitNewParameters events raised by the Platform contract.
type PlatformCommitNewParametersIterator struct {
	Event *PlatformCommitNewParameters // Event containing the contract specifics and raw log

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
func (it *PlatformCommitNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlatformCommitNewParameters)
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
		it.Event = new(PlatformCommitNewParameters)
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
func (it *PlatformCommitNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlatformCommitNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlatformCommitNewParameters represents a CommitNewParameters event raised by the Platform contract.
type PlatformCommitNewParameters struct {
	Deadline *big.Int
	A        *big.Int
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewParameters is a free log retrieval operation binding the contract event 0x6081daa3b61098baf24d9c69bcd53af932e0635c89c6fd0617534b9ba76a7f73.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 A, uint256 fee, uint256 admin_fee)
func (_Platform *PlatformFilterer) FilterCommitNewParameters(opts *bind.FilterOpts, deadline []*big.Int) (*PlatformCommitNewParametersIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Platform.contract.FilterLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return &PlatformCommitNewParametersIterator{contract: _Platform.contract, event: "CommitNewParameters", logs: logs, sub: sub}, nil
}

// WatchCommitNewParameters is a free log subscription operation binding the contract event 0x6081daa3b61098baf24d9c69bcd53af932e0635c89c6fd0617534b9ba76a7f73.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 A, uint256 fee, uint256 admin_fee)
func (_Platform *PlatformFilterer) WatchCommitNewParameters(opts *bind.WatchOpts, sink chan<- *PlatformCommitNewParameters, deadline []*big.Int) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Platform.contract.WatchLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlatformCommitNewParameters)
				if err := _Platform.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
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

// ParseCommitNewParameters is a log parse operation binding the contract event 0x6081daa3b61098baf24d9c69bcd53af932e0635c89c6fd0617534b9ba76a7f73.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 A, uint256 fee, uint256 admin_fee)
func (_Platform *PlatformFilterer) ParseCommitNewParameters(log types.Log) (*PlatformCommitNewParameters, error) {
	event := new(PlatformCommitNewParameters)
	if err := _Platform.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PlatformNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Platform contract.
type PlatformNewAdminIterator struct {
	Event *PlatformNewAdmin // Event containing the contract specifics and raw log

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
func (it *PlatformNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlatformNewAdmin)
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
		it.Event = new(PlatformNewAdmin)
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
func (it *PlatformNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlatformNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlatformNewAdmin represents a NewAdmin event raised by the Platform contract.
type PlatformNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Platform *PlatformFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address) (*PlatformNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Platform.contract.FilterLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &PlatformNewAdminIterator{contract: _Platform.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Platform *PlatformFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *PlatformNewAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Platform.contract.WatchLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlatformNewAdmin)
				if err := _Platform.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Platform *PlatformFilterer) ParseNewAdmin(log types.Log) (*PlatformNewAdmin, error) {
	event := new(PlatformNewAdmin)
	if err := _Platform.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PlatformNewParametersIterator is returned from FilterNewParameters and is used to iterate over the raw logs and unpacked data for NewParameters events raised by the Platform contract.
type PlatformNewParametersIterator struct {
	Event *PlatformNewParameters // Event containing the contract specifics and raw log

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
func (it *PlatformNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlatformNewParameters)
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
		it.Event = new(PlatformNewParameters)
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
func (it *PlatformNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlatformNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlatformNewParameters represents a NewParameters event raised by the Platform contract.
type PlatformNewParameters struct {
	A        *big.Int
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewParameters is a free log retrieval operation binding the contract event 0x752a27d1853eb7af3ee4ff764f2c4a51619386af721573dd3809e929c39db99e.
//
// Solidity: event NewParameters(uint256 A, uint256 fee, uint256 admin_fee)
func (_Platform *PlatformFilterer) FilterNewParameters(opts *bind.FilterOpts) (*PlatformNewParametersIterator, error) {

	logs, sub, err := _Platform.contract.FilterLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return &PlatformNewParametersIterator{contract: _Platform.contract, event: "NewParameters", logs: logs, sub: sub}, nil
}

// WatchNewParameters is a free log subscription operation binding the contract event 0x752a27d1853eb7af3ee4ff764f2c4a51619386af721573dd3809e929c39db99e.
//
// Solidity: event NewParameters(uint256 A, uint256 fee, uint256 admin_fee)
func (_Platform *PlatformFilterer) WatchNewParameters(opts *bind.WatchOpts, sink chan<- *PlatformNewParameters) (event.Subscription, error) {

	logs, sub, err := _Platform.contract.WatchLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlatformNewParameters)
				if err := _Platform.contract.UnpackLog(event, "NewParameters", log); err != nil {
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

// ParseNewParameters is a log parse operation binding the contract event 0x752a27d1853eb7af3ee4ff764f2c4a51619386af721573dd3809e929c39db99e.
//
// Solidity: event NewParameters(uint256 A, uint256 fee, uint256 admin_fee)
func (_Platform *PlatformFilterer) ParseNewParameters(log types.Log) (*PlatformNewParameters, error) {
	event := new(PlatformNewParameters)
	if err := _Platform.contract.UnpackLog(event, "NewParameters", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PlatformRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Platform contract.
type PlatformRemoveLiquidityIterator struct {
	Event *PlatformRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *PlatformRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlatformRemoveLiquidity)
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
		it.Event = new(PlatformRemoveLiquidity)
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
func (it *PlatformRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlatformRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlatformRemoveLiquidity represents a RemoveLiquidity event raised by the Platform contract.
type PlatformRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [4]*big.Int
	Fees         [4]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x9878ca375e106f2a43c3b599fc624568131c4c9a4ba66a14563715763be9d59d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 token_supply)
func (_Platform *PlatformFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*PlatformRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Platform.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &PlatformRemoveLiquidityIterator{contract: _Platform.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x9878ca375e106f2a43c3b599fc624568131c4c9a4ba66a14563715763be9d59d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 token_supply)
func (_Platform *PlatformFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *PlatformRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Platform.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlatformRemoveLiquidity)
				if err := _Platform.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x9878ca375e106f2a43c3b599fc624568131c4c9a4ba66a14563715763be9d59d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 token_supply)
func (_Platform *PlatformFilterer) ParseRemoveLiquidity(log types.Log) (*PlatformRemoveLiquidity, error) {
	event := new(PlatformRemoveLiquidity)
	if err := _Platform.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PlatformRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the Platform contract.
type PlatformRemoveLiquidityImbalanceIterator struct {
	Event *PlatformRemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *PlatformRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlatformRemoveLiquidityImbalance)
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
		it.Event = new(PlatformRemoveLiquidityImbalance)
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
func (it *PlatformRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlatformRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlatformRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the Platform contract.
type PlatformRemoveLiquidityImbalance struct {
	Provider     common.Address
	TokenAmounts [4]*big.Int
	Fees         [4]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityImbalance is a free log retrieval operation binding the contract event 0xb964b72f73f5ef5bf0fdc559b2fab9a7b12a39e47817a547f1f0aee47febd602.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_Platform *PlatformFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*PlatformRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Platform.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &PlatformRemoveLiquidityImbalanceIterator{contract: _Platform.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0xb964b72f73f5ef5bf0fdc559b2fab9a7b12a39e47817a547f1f0aee47febd602.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_Platform *PlatformFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *PlatformRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Platform.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlatformRemoveLiquidityImbalance)
				if err := _Platform.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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

// ParseRemoveLiquidityImbalance is a log parse operation binding the contract event 0xb964b72f73f5ef5bf0fdc559b2fab9a7b12a39e47817a547f1f0aee47febd602.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_Platform *PlatformFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*PlatformRemoveLiquidityImbalance, error) {
	event := new(PlatformRemoveLiquidityImbalance)
	if err := _Platform.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PlatformTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the Platform contract.
type PlatformTokenExchangeIterator struct {
	Event *PlatformTokenExchange // Event containing the contract specifics and raw log

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
func (it *PlatformTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlatformTokenExchange)
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
		it.Event = new(PlatformTokenExchange)
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
func (it *PlatformTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlatformTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlatformTokenExchange represents a TokenExchange event raised by the Platform contract.
type PlatformTokenExchange struct {
	Buyer        common.Address
	SoldId       *big.Int
	TokensSold   *big.Int
	BoughtId     *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenExchange is a free log retrieval operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Platform *PlatformFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*PlatformTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Platform.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &PlatformTokenExchangeIterator{contract: _Platform.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Platform *PlatformFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *PlatformTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Platform.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlatformTokenExchange)
				if err := _Platform.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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

// ParseTokenExchange is a log parse operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Platform *PlatformFilterer) ParseTokenExchange(log types.Log) (*PlatformTokenExchange, error) {
	event := new(PlatformTokenExchange)
	if err := _Platform.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PlatformTokenExchangeUnderlyingIterator is returned from FilterTokenExchangeUnderlying and is used to iterate over the raw logs and unpacked data for TokenExchangeUnderlying events raised by the Platform contract.
type PlatformTokenExchangeUnderlyingIterator struct {
	Event *PlatformTokenExchangeUnderlying // Event containing the contract specifics and raw log

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
func (it *PlatformTokenExchangeUnderlyingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlatformTokenExchangeUnderlying)
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
		it.Event = new(PlatformTokenExchangeUnderlying)
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
func (it *PlatformTokenExchangeUnderlyingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlatformTokenExchangeUnderlyingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlatformTokenExchangeUnderlying represents a TokenExchangeUnderlying event raised by the Platform contract.
type PlatformTokenExchangeUnderlying struct {
	Buyer        common.Address
	SoldId       *big.Int
	TokensSold   *big.Int
	BoughtId     *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangeUnderlying is a free log retrieval operation binding the contract event 0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b.
//
// Solidity: event TokenExchangeUnderlying(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Platform *PlatformFilterer) FilterTokenExchangeUnderlying(opts *bind.FilterOpts, buyer []common.Address) (*PlatformTokenExchangeUnderlyingIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Platform.contract.FilterLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return &PlatformTokenExchangeUnderlyingIterator{contract: _Platform.contract, event: "TokenExchangeUnderlying", logs: logs, sub: sub}, nil
}

// WatchTokenExchangeUnderlying is a free log subscription operation binding the contract event 0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b.
//
// Solidity: event TokenExchangeUnderlying(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Platform *PlatformFilterer) WatchTokenExchangeUnderlying(opts *bind.WatchOpts, sink chan<- *PlatformTokenExchangeUnderlying, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Platform.contract.WatchLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlatformTokenExchangeUnderlying)
				if err := _Platform.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
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

// ParseTokenExchangeUnderlying is a log parse operation binding the contract event 0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b.
//
// Solidity: event TokenExchangeUnderlying(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Platform *PlatformFilterer) ParseTokenExchangeUnderlying(log types.Log) (*PlatformTokenExchangeUnderlying, error) {
	event := new(PlatformTokenExchangeUnderlying)
	if err := _Platform.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
		return nil, err
	}
	return event, nil
}
