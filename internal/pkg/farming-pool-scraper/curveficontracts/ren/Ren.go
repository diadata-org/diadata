// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ren

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

// RenABI is the input ABI used to generate the binding from.
const RenABI = "[{\"name\":\"TokenExchange\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"sold_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"bought_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[2]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[2]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[2]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[2]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"token_amount\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"coin_amount\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[2]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[2]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true,\"unit\":\"sec\"},{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true,\"unit\":\"sec\"},{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampA\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"old_A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"new_A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"initial_time\",\"indexed\":false,\"unit\":\"sec\"},{\"type\":\"uint256\",\"name\":\"future_time\",\"indexed\":false,\"unit\":\"sec\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"t\",\"indexed\":false,\"unit\":\"sec\"}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address[2]\",\"name\":\"_coins\"},{\"type\":\"address\",\"name\":\"_pool_token\"},{\"type\":\"uint256\",\"name\":\"_A\"},{\"type\":\"uint256\",\"name\":\"_fee\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":5227},{\"name\":\"get_virtual_price\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":967716},{\"name\":\"calc_token_amount\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256[2]\",\"name\":\"amounts\"},{\"type\":\"bool\",\"name\":\"deposit\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":3810860},{\"name\":\"add_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[2]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"min_mint_amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":5858197},{\"name\":\"get_dy\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2327087},{\"name\":\"get_dy_underlying\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2326882},{\"name\":\"exchange\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"},{\"type\":\"uint256\",\"name\":\"min_dy\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":4782610},{\"name\":\"remove_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"},{\"type\":\"uint256[2]\",\"name\":\"min_amounts\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":153460},{\"name\":\"remove_liquidity_imbalance\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[2]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"max_burn_amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":5857985},{\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_token_amount\"},{\"type\":\"int128\",\"name\":\"i\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":10920},{\"name\":\"remove_liquidity_one_coin\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_token_amount\"},{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"uint256\",\"name\":\"min_amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":3677379},{\"name\":\"ramp_A\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_future_A\"},{\"type\":\"uint256\",\"unit\":\"sec\",\"name\":\"_future_time\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":151937},{\"name\":\"stop_ramp_A\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":148697},{\"name\":\"commit_new_fee\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"new_fee\"},{\"type\":\"uint256\",\"name\":\"new_admin_fee\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":110521},{\"name\":\"apply_new_fee\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":97220},{\"name\":\"revert_new_parameters\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":21955},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":74632},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":60688},{\"name\":\"revert_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":22045},{\"name\":\"withdraw_admin_fees\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":12424},{\"name\":\"kill_me\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":37998},{\"name\":\"unkill_me\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":22135},{\"name\":\"coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2310},{\"name\":\"balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2340},{\"name\":\"fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2171},{\"name\":\"admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2201},{\"name\":\"owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2231},{\"name\":\"initial_A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2261},{\"name\":\"future_A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2291},{\"name\":\"initial_A_time\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"sec\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2321},{\"name\":\"future_A_time\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"sec\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2351},{\"name\":\"admin_actions_deadline\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"sec\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2381},{\"name\":\"transfer_ownership_deadline\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"sec\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2411},{\"name\":\"future_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2441},{\"name\":\"future_admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2471},{\"name\":\"future_owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2501}]"

// Ren is an auto generated Go binding around an Ethereum contract.
type Ren struct {
	RenCaller     // Read-only binding to the contract
	RenTransactor // Write-only binding to the contract
	RenFilterer   // Log filterer for contract events
}

// RenCaller is an auto generated read-only Go binding around an Ethereum contract.
type RenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RenSession struct {
	Contract     *Ren              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RenCallerSession struct {
	Contract *RenCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RenTransactorSession struct {
	Contract     *RenTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RenRaw is an auto generated low-level Go binding around an Ethereum contract.
type RenRaw struct {
	Contract *Ren // Generic contract binding to access the raw methods on
}

// RenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RenCallerRaw struct {
	Contract *RenCaller // Generic read-only contract binding to access the raw methods on
}

// RenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RenTransactorRaw struct {
	Contract *RenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRen creates a new instance of Ren, bound to a specific deployed contract.
func NewRen(address common.Address, backend bind.ContractBackend) (*Ren, error) {
	contract, err := bindRen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ren{RenCaller: RenCaller{contract: contract}, RenTransactor: RenTransactor{contract: contract}, RenFilterer: RenFilterer{contract: contract}}, nil
}

// NewRenCaller creates a new read-only instance of Ren, bound to a specific deployed contract.
func NewRenCaller(address common.Address, caller bind.ContractCaller) (*RenCaller, error) {
	contract, err := bindRen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RenCaller{contract: contract}, nil
}

// NewRenTransactor creates a new write-only instance of Ren, bound to a specific deployed contract.
func NewRenTransactor(address common.Address, transactor bind.ContractTransactor) (*RenTransactor, error) {
	contract, err := bindRen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RenTransactor{contract: contract}, nil
}

// NewRenFilterer creates a new log filterer instance of Ren, bound to a specific deployed contract.
func NewRenFilterer(address common.Address, filterer bind.ContractFilterer) (*RenFilterer, error) {
	contract, err := bindRen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RenFilterer{contract: contract}, nil
}

// bindRen binds a generic wrapper to an already deployed contract.
func bindRen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ren *RenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ren.Contract.RenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ren *RenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ren.Contract.RenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ren *RenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ren.Contract.RenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ren *RenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ren.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ren *RenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ren.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ren *RenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ren.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() returns(uint256)
func (_Ren *RenCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() returns(uint256)
func (_Ren *RenSession) A() (*big.Int, error) {
	return _Ren.Contract.A(&_Ren.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() returns(uint256)
func (_Ren *RenCallerSession) A() (*big.Int, error) {
	return _Ren.Contract.A(&_Ren.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() returns(uint256)
func (_Ren *RenCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "admin_actions_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() returns(uint256)
func (_Ren *RenSession) AdminActionsDeadline() (*big.Int, error) {
	return _Ren.Contract.AdminActionsDeadline(&_Ren.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() returns(uint256)
func (_Ren *RenCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _Ren.Contract.AdminActionsDeadline(&_Ren.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() returns(uint256)
func (_Ren *RenCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() returns(uint256)
func (_Ren *RenSession) AdminFee() (*big.Int, error) {
	return _Ren.Contract.AdminFee(&_Ren.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() returns(uint256)
func (_Ren *RenCallerSession) AdminFee() (*big.Int, error) {
	return _Ren.Contract.AdminFee(&_Ren.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 arg0) returns(uint256)
func (_Ren *RenCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 arg0) returns(uint256)
func (_Ren *RenSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Ren.Contract.Balances(&_Ren.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 arg0) returns(uint256)
func (_Ren *RenCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Ren.Contract.Balances(&_Ren.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] amounts, bool deposit) returns(uint256)
func (_Ren *RenCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [2]*big.Int, deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "calc_token_amount", amounts, deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] amounts, bool deposit) returns(uint256)
func (_Ren *RenSession) CalcTokenAmount(amounts [2]*big.Int, deposit bool) (*big.Int, error) {
	return _Ren.Contract.CalcTokenAmount(&_Ren.CallOpts, amounts, deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] amounts, bool deposit) returns(uint256)
func (_Ren *RenCallerSession) CalcTokenAmount(amounts [2]*big.Int, deposit bool) (*big.Int, error) {
	return _Ren.Contract.CalcTokenAmount(&_Ren.CallOpts, amounts, deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) returns(uint256)
func (_Ren *RenCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, _token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "calc_withdraw_one_coin", _token_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) returns(uint256)
func (_Ren *RenSession) CalcWithdrawOneCoin(_token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Ren.Contract.CalcWithdrawOneCoin(&_Ren.CallOpts, _token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) returns(uint256)
func (_Ren *RenCallerSession) CalcWithdrawOneCoin(_token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Ren.Contract.CalcWithdrawOneCoin(&_Ren.CallOpts, _token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 arg0) returns(address)
func (_Ren *RenCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 arg0) returns(address)
func (_Ren *RenSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Ren.Contract.Coins(&_Ren.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 arg0) returns(address)
func (_Ren *RenCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Ren.Contract.Coins(&_Ren.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() returns(uint256)
func (_Ren *RenCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() returns(uint256)
func (_Ren *RenSession) Fee() (*big.Int, error) {
	return _Ren.Contract.Fee(&_Ren.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() returns(uint256)
func (_Ren *RenCallerSession) Fee() (*big.Int, error) {
	return _Ren.Contract.Fee(&_Ren.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() returns(uint256)
func (_Ren *RenCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() returns(uint256)
func (_Ren *RenSession) FutureA() (*big.Int, error) {
	return _Ren.Contract.FutureA(&_Ren.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() returns(uint256)
func (_Ren *RenCallerSession) FutureA() (*big.Int, error) {
	return _Ren.Contract.FutureA(&_Ren.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() returns(uint256)
func (_Ren *RenCaller) FutureATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "future_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() returns(uint256)
func (_Ren *RenSession) FutureATime() (*big.Int, error) {
	return _Ren.Contract.FutureATime(&_Ren.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() returns(uint256)
func (_Ren *RenCallerSession) FutureATime() (*big.Int, error) {
	return _Ren.Contract.FutureATime(&_Ren.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() returns(uint256)
func (_Ren *RenCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "future_admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() returns(uint256)
func (_Ren *RenSession) FutureAdminFee() (*big.Int, error) {
	return _Ren.Contract.FutureAdminFee(&_Ren.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() returns(uint256)
func (_Ren *RenCallerSession) FutureAdminFee() (*big.Int, error) {
	return _Ren.Contract.FutureAdminFee(&_Ren.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() returns(uint256)
func (_Ren *RenCaller) FutureFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "future_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() returns(uint256)
func (_Ren *RenSession) FutureFee() (*big.Int, error) {
	return _Ren.Contract.FutureFee(&_Ren.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() returns(uint256)
func (_Ren *RenCallerSession) FutureFee() (*big.Int, error) {
	return _Ren.Contract.FutureFee(&_Ren.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() returns(address)
func (_Ren *RenCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "future_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() returns(address)
func (_Ren *RenSession) FutureOwner() (common.Address, error) {
	return _Ren.Contract.FutureOwner(&_Ren.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() returns(address)
func (_Ren *RenCallerSession) FutureOwner() (common.Address, error) {
	return _Ren.Contract.FutureOwner(&_Ren.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Ren *RenCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Ren *RenSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Ren.Contract.GetDy(&_Ren.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Ren *RenCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Ren.Contract.GetDy(&_Ren.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Ren *RenCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "get_dy_underlying", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Ren *RenSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Ren.Contract.GetDyUnderlying(&_Ren.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Ren *RenCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Ren.Contract.GetDyUnderlying(&_Ren.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() returns(uint256)
func (_Ren *RenCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() returns(uint256)
func (_Ren *RenSession) GetVirtualPrice() (*big.Int, error) {
	return _Ren.Contract.GetVirtualPrice(&_Ren.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() returns(uint256)
func (_Ren *RenCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _Ren.Contract.GetVirtualPrice(&_Ren.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() returns(uint256)
func (_Ren *RenCaller) InitialA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "initial_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() returns(uint256)
func (_Ren *RenSession) InitialA() (*big.Int, error) {
	return _Ren.Contract.InitialA(&_Ren.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() returns(uint256)
func (_Ren *RenCallerSession) InitialA() (*big.Int, error) {
	return _Ren.Contract.InitialA(&_Ren.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() returns(uint256)
func (_Ren *RenCaller) InitialATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "initial_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() returns(uint256)
func (_Ren *RenSession) InitialATime() (*big.Int, error) {
	return _Ren.Contract.InitialATime(&_Ren.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() returns(uint256)
func (_Ren *RenCallerSession) InitialATime() (*big.Int, error) {
	return _Ren.Contract.InitialATime(&_Ren.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Ren *RenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Ren *RenSession) Owner() (common.Address, error) {
	return _Ren.Contract.Owner(&_Ren.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Ren *RenCallerSession) Owner() (common.Address, error) {
	return _Ren.Contract.Owner(&_Ren.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() returns(uint256)
func (_Ren *RenCaller) TransferOwnershipDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ren.contract.Call(opts, &out, "transfer_ownership_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() returns(uint256)
func (_Ren *RenSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _Ren.Contract.TransferOwnershipDeadline(&_Ren.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() returns(uint256)
func (_Ren *RenCallerSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _Ren.Contract.TransferOwnershipDeadline(&_Ren.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) returns()
func (_Ren *RenTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) returns()
func (_Ren *RenSession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Ren.Contract.AddLiquidity(&_Ren.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) returns()
func (_Ren *RenTransactorSession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Ren.Contract.AddLiquidity(&_Ren.TransactOpts, amounts, min_mint_amount)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_Ren *RenTransactor) ApplyNewFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "apply_new_fee")
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_Ren *RenSession) ApplyNewFee() (*types.Transaction, error) {
	return _Ren.Contract.ApplyNewFee(&_Ren.TransactOpts)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_Ren *RenTransactorSession) ApplyNewFee() (*types.Transaction, error) {
	return _Ren.Contract.ApplyNewFee(&_Ren.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Ren *RenTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Ren *RenSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Ren.Contract.ApplyTransferOwnership(&_Ren.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Ren *RenTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Ren.Contract.ApplyTransferOwnership(&_Ren.TransactOpts)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x5b5a1467.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee) returns()
func (_Ren *RenTransactor) CommitNewFee(opts *bind.TransactOpts, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "commit_new_fee", new_fee, new_admin_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x5b5a1467.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee) returns()
func (_Ren *RenSession) CommitNewFee(new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Ren.Contract.CommitNewFee(&_Ren.TransactOpts, new_fee, new_admin_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x5b5a1467.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee) returns()
func (_Ren *RenTransactorSession) CommitNewFee(new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Ren.Contract.CommitNewFee(&_Ren.TransactOpts, new_fee, new_admin_fee)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Ren *RenTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Ren *RenSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Ren.Contract.CommitTransferOwnership(&_Ren.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Ren *RenTransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Ren.Contract.CommitTransferOwnership(&_Ren.TransactOpts, _owner)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Ren *RenTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Ren *RenSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Ren.Contract.Exchange(&_Ren.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Ren *RenTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Ren.Contract.Exchange(&_Ren.TransactOpts, i, j, dx, min_dy)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Ren *RenTransactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Ren *RenSession) KillMe() (*types.Transaction, error) {
	return _Ren.Contract.KillMe(&_Ren.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Ren *RenTransactorSession) KillMe() (*types.Transaction, error) {
	return _Ren.Contract.KillMe(&_Ren.TransactOpts)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_Ren *RenTransactor) RampA(opts *bind.TransactOpts, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "ramp_A", _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_Ren *RenSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _Ren.Contract.RampA(&_Ren.TransactOpts, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_Ren *RenTransactorSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _Ren.Contract.RampA(&_Ren.TransactOpts, _future_A, _future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Ren *RenTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Ren *RenSession) RemoveLiquidity(_amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Ren.Contract.RemoveLiquidity(&_Ren.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Ren *RenTransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Ren.Contract.RemoveLiquidity(&_Ren.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 max_burn_amount) returns()
func (_Ren *RenTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, amounts [2]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "remove_liquidity_imbalance", amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 max_burn_amount) returns()
func (_Ren *RenSession) RemoveLiquidityImbalance(amounts [2]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Ren.Contract.RemoveLiquidityImbalance(&_Ren.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 max_burn_amount) returns()
func (_Ren *RenTransactorSession) RemoveLiquidityImbalance(amounts [2]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Ren.Contract.RemoveLiquidityImbalance(&_Ren.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_Ren *RenTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "remove_liquidity_one_coin", _token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_Ren *RenSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Ren.Contract.RemoveLiquidityOneCoin(&_Ren.TransactOpts, _token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_Ren *RenTransactorSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Ren.Contract.RemoveLiquidityOneCoin(&_Ren.TransactOpts, _token_amount, i, min_amount)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Ren *RenTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Ren *RenSession) RevertNewParameters() (*types.Transaction, error) {
	return _Ren.Contract.RevertNewParameters(&_Ren.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Ren *RenTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _Ren.Contract.RevertNewParameters(&_Ren.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Ren *RenTransactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Ren *RenSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Ren.Contract.RevertTransferOwnership(&_Ren.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Ren *RenTransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Ren.Contract.RevertTransferOwnership(&_Ren.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_Ren *RenTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "stop_ramp_A")
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_Ren *RenSession) StopRampA() (*types.Transaction, error) {
	return _Ren.Contract.StopRampA(&_Ren.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_Ren *RenTransactorSession) StopRampA() (*types.Transaction, error) {
	return _Ren.Contract.StopRampA(&_Ren.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Ren *RenTransactor) UnkillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "unkill_me")
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Ren *RenSession) UnkillMe() (*types.Transaction, error) {
	return _Ren.Contract.UnkillMe(&_Ren.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Ren *RenTransactorSession) UnkillMe() (*types.Transaction, error) {
	return _Ren.Contract.UnkillMe(&_Ren.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Ren *RenTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ren.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Ren *RenSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Ren.Contract.WithdrawAdminFees(&_Ren.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Ren *RenTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Ren.Contract.WithdrawAdminFees(&_Ren.TransactOpts)
}

// RenAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Ren contract.
type RenAddLiquidityIterator struct {
	Event *RenAddLiquidity // Event containing the contract specifics and raw log

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
func (it *RenAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenAddLiquidity)
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
		it.Event = new(RenAddLiquidity)
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
func (it *RenAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenAddLiquidity represents a AddLiquidity event raised by the Ren contract.
type RenAddLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fees         [2]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x26f55a85081d24974e85c6c00045d0f0453991e95873f52bff0d21af4079a768.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_Ren *RenFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*RenAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Ren.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &RenAddLiquidityIterator{contract: _Ren.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x26f55a85081d24974e85c6c00045d0f0453991e95873f52bff0d21af4079a768.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_Ren *RenFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *RenAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Ren.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenAddLiquidity)
				if err := _Ren.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x26f55a85081d24974e85c6c00045d0f0453991e95873f52bff0d21af4079a768.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_Ren *RenFilterer) ParseAddLiquidity(log types.Log) (*RenAddLiquidity, error) {
	event := new(RenAddLiquidity)
	if err := _Ren.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RenCommitNewAdminIterator is returned from FilterCommitNewAdmin and is used to iterate over the raw logs and unpacked data for CommitNewAdmin events raised by the Ren contract.
type RenCommitNewAdminIterator struct {
	Event *RenCommitNewAdmin // Event containing the contract specifics and raw log

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
func (it *RenCommitNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenCommitNewAdmin)
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
		it.Event = new(RenCommitNewAdmin)
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
func (it *RenCommitNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenCommitNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenCommitNewAdmin represents a CommitNewAdmin event raised by the Ren contract.
type RenCommitNewAdmin struct {
	Deadline *big.Int
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewAdmin is a free log retrieval operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Ren *RenFilterer) FilterCommitNewAdmin(opts *bind.FilterOpts, deadline []*big.Int, admin []common.Address) (*RenCommitNewAdminIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Ren.contract.FilterLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &RenCommitNewAdminIterator{contract: _Ren.contract, event: "CommitNewAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitNewAdmin is a free log subscription operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Ren *RenFilterer) WatchCommitNewAdmin(opts *bind.WatchOpts, sink chan<- *RenCommitNewAdmin, deadline []*big.Int, admin []common.Address) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Ren.contract.WatchLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenCommitNewAdmin)
				if err := _Ren.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
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
func (_Ren *RenFilterer) ParseCommitNewAdmin(log types.Log) (*RenCommitNewAdmin, error) {
	event := new(RenCommitNewAdmin)
	if err := _Ren.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RenCommitNewFeeIterator is returned from FilterCommitNewFee and is used to iterate over the raw logs and unpacked data for CommitNewFee events raised by the Ren contract.
type RenCommitNewFeeIterator struct {
	Event *RenCommitNewFee // Event containing the contract specifics and raw log

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
func (it *RenCommitNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenCommitNewFee)
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
		it.Event = new(RenCommitNewFee)
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
func (it *RenCommitNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenCommitNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenCommitNewFee represents a CommitNewFee event raised by the Ren contract.
type RenCommitNewFee struct {
	Deadline *big.Int
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewFee is a free log retrieval operation binding the contract event 0x351fc5da2fbf480f2225debf3664a4bc90fa9923743aad58b4603f648e931fe0.
//
// Solidity: event CommitNewFee(uint256 indexed deadline, uint256 fee, uint256 admin_fee)
func (_Ren *RenFilterer) FilterCommitNewFee(opts *bind.FilterOpts, deadline []*big.Int) (*RenCommitNewFeeIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Ren.contract.FilterLogs(opts, "CommitNewFee", deadlineRule)
	if err != nil {
		return nil, err
	}
	return &RenCommitNewFeeIterator{contract: _Ren.contract, event: "CommitNewFee", logs: logs, sub: sub}, nil
}

// WatchCommitNewFee is a free log subscription operation binding the contract event 0x351fc5da2fbf480f2225debf3664a4bc90fa9923743aad58b4603f648e931fe0.
//
// Solidity: event CommitNewFee(uint256 indexed deadline, uint256 fee, uint256 admin_fee)
func (_Ren *RenFilterer) WatchCommitNewFee(opts *bind.WatchOpts, sink chan<- *RenCommitNewFee, deadline []*big.Int) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Ren.contract.WatchLogs(opts, "CommitNewFee", deadlineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenCommitNewFee)
				if err := _Ren.contract.UnpackLog(event, "CommitNewFee", log); err != nil {
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

// ParseCommitNewFee is a log parse operation binding the contract event 0x351fc5da2fbf480f2225debf3664a4bc90fa9923743aad58b4603f648e931fe0.
//
// Solidity: event CommitNewFee(uint256 indexed deadline, uint256 fee, uint256 admin_fee)
func (_Ren *RenFilterer) ParseCommitNewFee(log types.Log) (*RenCommitNewFee, error) {
	event := new(RenCommitNewFee)
	if err := _Ren.contract.UnpackLog(event, "CommitNewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RenNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Ren contract.
type RenNewAdminIterator struct {
	Event *RenNewAdmin // Event containing the contract specifics and raw log

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
func (it *RenNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenNewAdmin)
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
		it.Event = new(RenNewAdmin)
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
func (it *RenNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenNewAdmin represents a NewAdmin event raised by the Ren contract.
type RenNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Ren *RenFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address) (*RenNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Ren.contract.FilterLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &RenNewAdminIterator{contract: _Ren.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Ren *RenFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *RenNewAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Ren.contract.WatchLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenNewAdmin)
				if err := _Ren.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_Ren *RenFilterer) ParseNewAdmin(log types.Log) (*RenNewAdmin, error) {
	event := new(RenNewAdmin)
	if err := _Ren.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RenNewFeeIterator is returned from FilterNewFee and is used to iterate over the raw logs and unpacked data for NewFee events raised by the Ren contract.
type RenNewFeeIterator struct {
	Event *RenNewFee // Event containing the contract specifics and raw log

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
func (it *RenNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenNewFee)
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
		it.Event = new(RenNewFee)
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
func (it *RenNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenNewFee represents a NewFee event raised by the Ren contract.
type RenNewFee struct {
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewFee is a free log retrieval operation binding the contract event 0xbe12859b636aed607d5230b2cc2711f68d70e51060e6cca1f575ef5d2fcc95d1.
//
// Solidity: event NewFee(uint256 fee, uint256 admin_fee)
func (_Ren *RenFilterer) FilterNewFee(opts *bind.FilterOpts) (*RenNewFeeIterator, error) {

	logs, sub, err := _Ren.contract.FilterLogs(opts, "NewFee")
	if err != nil {
		return nil, err
	}
	return &RenNewFeeIterator{contract: _Ren.contract, event: "NewFee", logs: logs, sub: sub}, nil
}

// WatchNewFee is a free log subscription operation binding the contract event 0xbe12859b636aed607d5230b2cc2711f68d70e51060e6cca1f575ef5d2fcc95d1.
//
// Solidity: event NewFee(uint256 fee, uint256 admin_fee)
func (_Ren *RenFilterer) WatchNewFee(opts *bind.WatchOpts, sink chan<- *RenNewFee) (event.Subscription, error) {

	logs, sub, err := _Ren.contract.WatchLogs(opts, "NewFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenNewFee)
				if err := _Ren.contract.UnpackLog(event, "NewFee", log); err != nil {
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

// ParseNewFee is a log parse operation binding the contract event 0xbe12859b636aed607d5230b2cc2711f68d70e51060e6cca1f575ef5d2fcc95d1.
//
// Solidity: event NewFee(uint256 fee, uint256 admin_fee)
func (_Ren *RenFilterer) ParseNewFee(log types.Log) (*RenNewFee, error) {
	event := new(RenNewFee)
	if err := _Ren.contract.UnpackLog(event, "NewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RenRampAIterator is returned from FilterRampA and is used to iterate over the raw logs and unpacked data for RampA events raised by the Ren contract.
type RenRampAIterator struct {
	Event *RenRampA // Event containing the contract specifics and raw log

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
func (it *RenRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenRampA)
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
		it.Event = new(RenRampA)
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
func (it *RenRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenRampA represents a RampA event raised by the Ren contract.
type RenRampA struct {
	OldA        *big.Int
	NewA        *big.Int
	InitialTime *big.Int
	FutureTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRampA is a free log retrieval operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_Ren *RenFilterer) FilterRampA(opts *bind.FilterOpts) (*RenRampAIterator, error) {

	logs, sub, err := _Ren.contract.FilterLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return &RenRampAIterator{contract: _Ren.contract, event: "RampA", logs: logs, sub: sub}, nil
}

// WatchRampA is a free log subscription operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_Ren *RenFilterer) WatchRampA(opts *bind.WatchOpts, sink chan<- *RenRampA) (event.Subscription, error) {

	logs, sub, err := _Ren.contract.WatchLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenRampA)
				if err := _Ren.contract.UnpackLog(event, "RampA", log); err != nil {
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

// ParseRampA is a log parse operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_Ren *RenFilterer) ParseRampA(log types.Log) (*RenRampA, error) {
	event := new(RenRampA)
	if err := _Ren.contract.UnpackLog(event, "RampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RenRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Ren contract.
type RenRemoveLiquidityIterator struct {
	Event *RenRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *RenRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenRemoveLiquidity)
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
		it.Event = new(RenRemoveLiquidity)
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
func (it *RenRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenRemoveLiquidity represents a RemoveLiquidity event raised by the Ren contract.
type RenRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fees         [2]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_Ren *RenFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*RenRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Ren.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &RenRemoveLiquidityIterator{contract: _Ren.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_Ren *RenFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *RenRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Ren.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenRemoveLiquidity)
				if err := _Ren.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_Ren *RenFilterer) ParseRemoveLiquidity(log types.Log) (*RenRemoveLiquidity, error) {
	event := new(RenRemoveLiquidity)
	if err := _Ren.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RenRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the Ren contract.
type RenRemoveLiquidityImbalanceIterator struct {
	Event *RenRemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *RenRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenRemoveLiquidityImbalance)
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
		it.Event = new(RenRemoveLiquidityImbalance)
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
func (it *RenRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the Ren contract.
type RenRemoveLiquidityImbalance struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fees         [2]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityImbalance is a free log retrieval operation binding the contract event 0x2b5508378d7e19e0d5fa338419034731416c4f5b219a10379956f764317fd47e.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_Ren *RenFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*RenRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Ren.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &RenRemoveLiquidityImbalanceIterator{contract: _Ren.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x2b5508378d7e19e0d5fa338419034731416c4f5b219a10379956f764317fd47e.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_Ren *RenFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *RenRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Ren.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenRemoveLiquidityImbalance)
				if err := _Ren.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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

// ParseRemoveLiquidityImbalance is a log parse operation binding the contract event 0x2b5508378d7e19e0d5fa338419034731416c4f5b219a10379956f764317fd47e.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_Ren *RenFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*RenRemoveLiquidityImbalance, error) {
	event := new(RenRemoveLiquidityImbalance)
	if err := _Ren.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RenRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the Ren contract.
type RenRemoveLiquidityOneIterator struct {
	Event *RenRemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *RenRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenRemoveLiquidityOne)
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
		it.Event = new(RenRemoveLiquidityOne)
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
func (it *RenRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the Ren contract.
type RenRemoveLiquidityOne struct {
	Provider    common.Address
	TokenAmount *big.Int
	CoinAmount  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x9e96dd3b997a2a257eec4df9bb6eaf626e206df5f543bd963682d143300be310.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount)
func (_Ren *RenFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*RenRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Ren.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &RenRemoveLiquidityOneIterator{contract: _Ren.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x9e96dd3b997a2a257eec4df9bb6eaf626e206df5f543bd963682d143300be310.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount)
func (_Ren *RenFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *RenRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Ren.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenRemoveLiquidityOne)
				if err := _Ren.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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

// ParseRemoveLiquidityOne is a log parse operation binding the contract event 0x9e96dd3b997a2a257eec4df9bb6eaf626e206df5f543bd963682d143300be310.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount)
func (_Ren *RenFilterer) ParseRemoveLiquidityOne(log types.Log) (*RenRemoveLiquidityOne, error) {
	event := new(RenRemoveLiquidityOne)
	if err := _Ren.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RenStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the Ren contract.
type RenStopRampAIterator struct {
	Event *RenStopRampA // Event containing the contract specifics and raw log

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
func (it *RenStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenStopRampA)
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
		it.Event = new(RenStopRampA)
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
func (it *RenStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenStopRampA represents a StopRampA event raised by the Ren contract.
type RenStopRampA struct {
	A   *big.Int
	T   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_Ren *RenFilterer) FilterStopRampA(opts *bind.FilterOpts) (*RenStopRampAIterator, error) {

	logs, sub, err := _Ren.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &RenStopRampAIterator{contract: _Ren.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_Ren *RenFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *RenStopRampA) (event.Subscription, error) {

	logs, sub, err := _Ren.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenStopRampA)
				if err := _Ren.contract.UnpackLog(event, "StopRampA", log); err != nil {
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

// ParseStopRampA is a log parse operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_Ren *RenFilterer) ParseStopRampA(log types.Log) (*RenStopRampA, error) {
	event := new(RenStopRampA)
	if err := _Ren.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RenTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the Ren contract.
type RenTokenExchangeIterator struct {
	Event *RenTokenExchange // Event containing the contract specifics and raw log

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
func (it *RenTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenTokenExchange)
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
		it.Event = new(RenTokenExchange)
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
func (it *RenTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenTokenExchange represents a TokenExchange event raised by the Ren contract.
type RenTokenExchange struct {
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
func (_Ren *RenFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*RenTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Ren.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &RenTokenExchangeIterator{contract: _Ren.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Ren *RenFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *RenTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Ren.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenTokenExchange)
				if err := _Ren.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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
func (_Ren *RenFilterer) ParseTokenExchange(log types.Log) (*RenTokenExchange, error) {
	event := new(RenTokenExchange)
	if err := _Ren.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
