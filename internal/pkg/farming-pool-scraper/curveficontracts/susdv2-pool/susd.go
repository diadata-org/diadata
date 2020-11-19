// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package susd

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

// SusdABI is the input ABI used to generate the binding from.
const SusdABI = "[{\"name\":\"TokenExchange\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"sold_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"bought_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchangeUnderlying\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"sold_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"bought_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[4]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[4]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[4]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[4]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[4]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[4]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true,\"unit\":\"sec\"},{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewParameters\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true,\"unit\":\"sec\"},{\"type\":\"uint256\",\"name\":\"A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewParameters\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address[4]\",\"name\":\"_coins\"},{\"type\":\"address[4]\",\"name\":\"_underlying_coins\"},{\"type\":\"address\",\"name\":\"_pool_token\"},{\"type\":\"uint256\",\"name\":\"_A\"},{\"type\":\"uint256\",\"name\":\"_fee\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"get_virtual_price\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1570535},{\"name\":\"calc_token_amount\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256[4]\",\"name\":\"amounts\"},{\"type\":\"bool\",\"name\":\"deposit\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6103471},{\"name\":\"add_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[4]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"min_mint_amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":9331701},{\"name\":\"get_dy\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":3489637},{\"name\":\"get_dy_underlying\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":3489467},{\"name\":\"exchange\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"},{\"type\":\"uint256\",\"name\":\"min_dy\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":7034253},{\"name\":\"exchange_underlying\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"},{\"type\":\"uint256\",\"name\":\"min_dy\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":7050488},{\"name\":\"remove_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"},{\"type\":\"uint256[4]\",\"name\":\"min_amounts\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":241191},{\"name\":\"remove_liquidity_imbalance\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[4]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"max_burn_amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":9330864},{\"name\":\"commit_new_parameters\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"amplification\"},{\"type\":\"uint256\",\"name\":\"new_fee\"},{\"type\":\"uint256\",\"name\":\"new_admin_fee\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":146045},{\"name\":\"apply_new_parameters\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":133452},{\"name\":\"revert_new_parameters\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":21775},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":74452},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":60508},{\"name\":\"revert_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":21865},{\"name\":\"withdraw_admin_fees\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":23448},{\"name\":\"kill_me\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":37818},{\"name\":\"unkill_me\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":21955},{\"name\":\"coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2130},{\"name\":\"underlying_coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2160},{\"name\":\"balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2190},{\"name\":\"A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2021},{\"name\":\"fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2051},{\"name\":\"admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2081},{\"name\":\"owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2111},{\"name\":\"admin_actions_deadline\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"sec\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2141},{\"name\":\"transfer_ownership_deadline\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"sec\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2171},{\"name\":\"future_A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2201},{\"name\":\"future_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2231},{\"name\":\"future_admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2261},{\"name\":\"future_owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2291}]"

// Susd is an auto generated Go binding around an Ethereum contract.
type Susd struct {
	SusdCaller     // Read-only binding to the contract
	SusdTransactor // Write-only binding to the contract
	SusdFilterer   // Log filterer for contract events
}

// SusdCaller is an auto generated read-only Go binding around an Ethereum contract.
type SusdCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SusdTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SusdTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SusdFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SusdFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SusdSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SusdSession struct {
	Contract     *Susd             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SusdCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SusdCallerSession struct {
	Contract *SusdCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SusdTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SusdTransactorSession struct {
	Contract     *SusdTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SusdRaw is an auto generated low-level Go binding around an Ethereum contract.
type SusdRaw struct {
	Contract *Susd // Generic contract binding to access the raw methods on
}

// SusdCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SusdCallerRaw struct {
	Contract *SusdCaller // Generic read-only contract binding to access the raw methods on
}

// SusdTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SusdTransactorRaw struct {
	Contract *SusdTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSusd creates a new instance of Susd, bound to a specific deployed contract.
func NewSusd(address common.Address, backend bind.ContractBackend) (*Susd, error) {
	contract, err := bindSusd(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Susd{SusdCaller: SusdCaller{contract: contract}, SusdTransactor: SusdTransactor{contract: contract}, SusdFilterer: SusdFilterer{contract: contract}}, nil
}

// NewSusdCaller creates a new read-only instance of Susd, bound to a specific deployed contract.
func NewSusdCaller(address common.Address, caller bind.ContractCaller) (*SusdCaller, error) {
	contract, err := bindSusd(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SusdCaller{contract: contract}, nil
}

// NewSusdTransactor creates a new write-only instance of Susd, bound to a specific deployed contract.
func NewSusdTransactor(address common.Address, transactor bind.ContractTransactor) (*SusdTransactor, error) {
	contract, err := bindSusd(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SusdTransactor{contract: contract}, nil
}

// NewSusdFilterer creates a new log filterer instance of Susd, bound to a specific deployed contract.
func NewSusdFilterer(address common.Address, filterer bind.ContractFilterer) (*SusdFilterer, error) {
	contract, err := bindSusd(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SusdFilterer{contract: contract}, nil
}

// bindSusd binds a generic wrapper to an already deployed contract.
func bindSusd(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SusdABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Susd *SusdRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Susd.Contract.SusdCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Susd *SusdRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Susd.Contract.SusdTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Susd *SusdRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Susd.Contract.SusdTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Susd *SusdCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Susd.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Susd *SusdTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Susd.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Susd *SusdTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Susd.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() returns(uint256)
func (_Susd *SusdCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() returns(uint256)
func (_Susd *SusdSession) A() (*big.Int, error) {
	return _Susd.Contract.A(&_Susd.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() returns(uint256)
func (_Susd *SusdCallerSession) A() (*big.Int, error) {
	return _Susd.Contract.A(&_Susd.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() returns(uint256)
func (_Susd *SusdCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "admin_actions_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() returns(uint256)
func (_Susd *SusdSession) AdminActionsDeadline() (*big.Int, error) {
	return _Susd.Contract.AdminActionsDeadline(&_Susd.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() returns(uint256)
func (_Susd *SusdCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _Susd.Contract.AdminActionsDeadline(&_Susd.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() returns(uint256)
func (_Susd *SusdCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() returns(uint256)
func (_Susd *SusdSession) AdminFee() (*big.Int, error) {
	return _Susd.Contract.AdminFee(&_Susd.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() returns(uint256)
func (_Susd *SusdCallerSession) AdminFee() (*big.Int, error) {
	return _Susd.Contract.AdminFee(&_Susd.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 arg0) returns(uint256)
func (_Susd *SusdCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 arg0) returns(uint256)
func (_Susd *SusdSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Susd.Contract.Balances(&_Susd.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 arg0) returns(uint256)
func (_Susd *SusdCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Susd.Contract.Balances(&_Susd.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xcf701ff7.
//
// Solidity: function calc_token_amount(uint256[4] amounts, bool deposit) returns(uint256)
func (_Susd *SusdCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [4]*big.Int, deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "calc_token_amount", amounts, deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xcf701ff7.
//
// Solidity: function calc_token_amount(uint256[4] amounts, bool deposit) returns(uint256)
func (_Susd *SusdSession) CalcTokenAmount(amounts [4]*big.Int, deposit bool) (*big.Int, error) {
	return _Susd.Contract.CalcTokenAmount(&_Susd.CallOpts, amounts, deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xcf701ff7.
//
// Solidity: function calc_token_amount(uint256[4] amounts, bool deposit) returns(uint256)
func (_Susd *SusdCallerSession) CalcTokenAmount(amounts [4]*big.Int, deposit bool) (*big.Int, error) {
	return _Susd.Contract.CalcTokenAmount(&_Susd.CallOpts, amounts, deposit)
}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 arg0) returns(address)
func (_Susd *SusdCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 arg0) returns(address)
func (_Susd *SusdSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Susd.Contract.Coins(&_Susd.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 arg0) returns(address)
func (_Susd *SusdCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Susd.Contract.Coins(&_Susd.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() returns(uint256)
func (_Susd *SusdCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() returns(uint256)
func (_Susd *SusdSession) Fee() (*big.Int, error) {
	return _Susd.Contract.Fee(&_Susd.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() returns(uint256)
func (_Susd *SusdCallerSession) Fee() (*big.Int, error) {
	return _Susd.Contract.Fee(&_Susd.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() returns(uint256)
func (_Susd *SusdCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() returns(uint256)
func (_Susd *SusdSession) FutureA() (*big.Int, error) {
	return _Susd.Contract.FutureA(&_Susd.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() returns(uint256)
func (_Susd *SusdCallerSession) FutureA() (*big.Int, error) {
	return _Susd.Contract.FutureA(&_Susd.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() returns(uint256)
func (_Susd *SusdCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "future_admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() returns(uint256)
func (_Susd *SusdSession) FutureAdminFee() (*big.Int, error) {
	return _Susd.Contract.FutureAdminFee(&_Susd.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() returns(uint256)
func (_Susd *SusdCallerSession) FutureAdminFee() (*big.Int, error) {
	return _Susd.Contract.FutureAdminFee(&_Susd.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() returns(uint256)
func (_Susd *SusdCaller) FutureFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "future_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() returns(uint256)
func (_Susd *SusdSession) FutureFee() (*big.Int, error) {
	return _Susd.Contract.FutureFee(&_Susd.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() returns(uint256)
func (_Susd *SusdCallerSession) FutureFee() (*big.Int, error) {
	return _Susd.Contract.FutureFee(&_Susd.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() returns(address)
func (_Susd *SusdCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "future_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() returns(address)
func (_Susd *SusdSession) FutureOwner() (common.Address, error) {
	return _Susd.Contract.FutureOwner(&_Susd.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() returns(address)
func (_Susd *SusdCallerSession) FutureOwner() (common.Address, error) {
	return _Susd.Contract.FutureOwner(&_Susd.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Susd *SusdCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Susd *SusdSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Susd.Contract.GetDy(&_Susd.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Susd *SusdCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Susd.Contract.GetDy(&_Susd.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Susd *SusdCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "get_dy_underlying", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Susd *SusdSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Susd.Contract.GetDyUnderlying(&_Susd.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Susd *SusdCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Susd.Contract.GetDyUnderlying(&_Susd.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() returns(uint256)
func (_Susd *SusdCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() returns(uint256)
func (_Susd *SusdSession) GetVirtualPrice() (*big.Int, error) {
	return _Susd.Contract.GetVirtualPrice(&_Susd.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() returns(uint256)
func (_Susd *SusdCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _Susd.Contract.GetVirtualPrice(&_Susd.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Susd *SusdCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Susd *SusdSession) Owner() (common.Address, error) {
	return _Susd.Contract.Owner(&_Susd.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Susd *SusdCallerSession) Owner() (common.Address, error) {
	return _Susd.Contract.Owner(&_Susd.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() returns(uint256)
func (_Susd *SusdCaller) TransferOwnershipDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "transfer_ownership_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() returns(uint256)
func (_Susd *SusdSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _Susd.Contract.TransferOwnershipDeadline(&_Susd.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() returns(uint256)
func (_Susd *SusdCallerSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _Susd.Contract.TransferOwnershipDeadline(&_Susd.CallOpts)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 arg0) returns(address)
func (_Susd *SusdCaller) UnderlyingCoins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Susd.contract.Call(opts, &out, "underlying_coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 arg0) returns(address)
func (_Susd *SusdSession) UnderlyingCoins(arg0 *big.Int) (common.Address, error) {
	return _Susd.Contract.UnderlyingCoins(&_Susd.CallOpts, arg0)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 arg0) returns(address)
func (_Susd *SusdCallerSession) UnderlyingCoins(arg0 *big.Int) (common.Address, error) {
	return _Susd.Contract.UnderlyingCoins(&_Susd.CallOpts, arg0)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x029b2f34.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 min_mint_amount) returns()
func (_Susd *SusdTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [4]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Susd.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x029b2f34.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 min_mint_amount) returns()
func (_Susd *SusdSession) AddLiquidity(amounts [4]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Susd.Contract.AddLiquidity(&_Susd.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x029b2f34.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 min_mint_amount) returns()
func (_Susd *SusdTransactorSession) AddLiquidity(amounts [4]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Susd.Contract.AddLiquidity(&_Susd.TransactOpts, amounts, min_mint_amount)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Susd *SusdTransactor) ApplyNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Susd.contract.Transact(opts, "apply_new_parameters")
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Susd *SusdSession) ApplyNewParameters() (*types.Transaction, error) {
	return _Susd.Contract.ApplyNewParameters(&_Susd.TransactOpts)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Susd *SusdTransactorSession) ApplyNewParameters() (*types.Transaction, error) {
	return _Susd.Contract.ApplyNewParameters(&_Susd.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Susd *SusdTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Susd.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Susd *SusdSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Susd.Contract.ApplyTransferOwnership(&_Susd.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Susd *SusdTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Susd.Contract.ApplyTransferOwnership(&_Susd.TransactOpts)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xee11f5b6.
//
// Solidity: function commit_new_parameters(uint256 amplification, uint256 new_fee, uint256 new_admin_fee) returns()
func (_Susd *SusdTransactor) CommitNewParameters(opts *bind.TransactOpts, amplification *big.Int, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Susd.contract.Transact(opts, "commit_new_parameters", amplification, new_fee, new_admin_fee)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xee11f5b6.
//
// Solidity: function commit_new_parameters(uint256 amplification, uint256 new_fee, uint256 new_admin_fee) returns()
func (_Susd *SusdSession) CommitNewParameters(amplification *big.Int, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Susd.Contract.CommitNewParameters(&_Susd.TransactOpts, amplification, new_fee, new_admin_fee)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xee11f5b6.
//
// Solidity: function commit_new_parameters(uint256 amplification, uint256 new_fee, uint256 new_admin_fee) returns()
func (_Susd *SusdTransactorSession) CommitNewParameters(amplification *big.Int, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Susd.Contract.CommitNewParameters(&_Susd.TransactOpts, amplification, new_fee, new_admin_fee)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Susd *SusdTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Susd.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Susd *SusdSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Susd.Contract.CommitTransferOwnership(&_Susd.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Susd *SusdTransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Susd.Contract.CommitTransferOwnership(&_Susd.TransactOpts, _owner)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Susd *SusdTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Susd.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Susd *SusdSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Susd.Contract.Exchange(&_Susd.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Susd *SusdTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Susd.Contract.Exchange(&_Susd.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Susd *SusdTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Susd.contract.Transact(opts, "exchange_underlying", i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Susd *SusdSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Susd.Contract.ExchangeUnderlying(&_Susd.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Susd *SusdTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Susd.Contract.ExchangeUnderlying(&_Susd.TransactOpts, i, j, dx, min_dy)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Susd *SusdTransactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Susd.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Susd *SusdSession) KillMe() (*types.Transaction, error) {
	return _Susd.Contract.KillMe(&_Susd.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Susd *SusdTransactorSession) KillMe() (*types.Transaction, error) {
	return _Susd.Contract.KillMe(&_Susd.TransactOpts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x7d49d875.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[4] min_amounts) returns()
func (_Susd *SusdTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [4]*big.Int) (*types.Transaction, error) {
	return _Susd.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x7d49d875.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[4] min_amounts) returns()
func (_Susd *SusdSession) RemoveLiquidity(_amount *big.Int, min_amounts [4]*big.Int) (*types.Transaction, error) {
	return _Susd.Contract.RemoveLiquidity(&_Susd.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x7d49d875.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[4] min_amounts) returns()
func (_Susd *SusdTransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [4]*big.Int) (*types.Transaction, error) {
	return _Susd.Contract.RemoveLiquidity(&_Susd.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x18a7bd76.
//
// Solidity: function remove_liquidity_imbalance(uint256[4] amounts, uint256 max_burn_amount) returns()
func (_Susd *SusdTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, amounts [4]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Susd.contract.Transact(opts, "remove_liquidity_imbalance", amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x18a7bd76.
//
// Solidity: function remove_liquidity_imbalance(uint256[4] amounts, uint256 max_burn_amount) returns()
func (_Susd *SusdSession) RemoveLiquidityImbalance(amounts [4]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Susd.Contract.RemoveLiquidityImbalance(&_Susd.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x18a7bd76.
//
// Solidity: function remove_liquidity_imbalance(uint256[4] amounts, uint256 max_burn_amount) returns()
func (_Susd *SusdTransactorSession) RemoveLiquidityImbalance(amounts [4]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Susd.Contract.RemoveLiquidityImbalance(&_Susd.TransactOpts, amounts, max_burn_amount)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Susd *SusdTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Susd.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Susd *SusdSession) RevertNewParameters() (*types.Transaction, error) {
	return _Susd.Contract.RevertNewParameters(&_Susd.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Susd *SusdTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _Susd.Contract.RevertNewParameters(&_Susd.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Susd *SusdTransactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Susd.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Susd *SusdSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Susd.Contract.RevertTransferOwnership(&_Susd.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Susd *SusdTransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Susd.Contract.RevertTransferOwnership(&_Susd.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Susd *SusdTransactor) UnkillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Susd.contract.Transact(opts, "unkill_me")
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Susd *SusdSession) UnkillMe() (*types.Transaction, error) {
	return _Susd.Contract.UnkillMe(&_Susd.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Susd *SusdTransactorSession) UnkillMe() (*types.Transaction, error) {
	return _Susd.Contract.UnkillMe(&_Susd.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Susd *SusdTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Susd.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Susd *SusdSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Susd.Contract.WithdrawAdminFees(&_Susd.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Susd *SusdTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Susd.Contract.WithdrawAdminFees(&_Susd.TransactOpts)
}

// SusdAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Susd contract.
type SusdAddLiquidityIterator struct {
	Event *SusdAddLiquidity // Event containing the contract specifics and raw log

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
func (it *SusdAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SusdAddLiquidity)
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
		it.Event = new(SusdAddLiquidity)
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
func (it *SusdAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SusdAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SusdAddLiquidity represents a AddLiquidity event raised by the Susd contract.
type SusdAddLiquidity struct {
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
func (_Susd *SusdFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*SusdAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Susd.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &SusdAddLiquidityIterator{contract: _Susd.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x3f1915775e0c9a38a57a7bb7f1f9005f486fb904e1f84aa215364d567319a58d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_Susd *SusdFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *SusdAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Susd.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SusdAddLiquidity)
				if err := _Susd.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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
func (_Susd *SusdFilterer) ParseAddLiquidity(log types.Log) (*SusdAddLiquidity, error) {
	event := new(SusdAddLiquidity)
	if err := _Susd.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SusdCommitNewAdminIterator is returned from FilterCommitNewAdmin and is used to iterate over the raw logs and unpacked data for CommitNewAdmin events raised by the Susd contract.
type SusdCommitNewAdminIterator struct {
	Event *SusdCommitNewAdmin // Event containing the contract specifics and raw log

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
func (it *SusdCommitNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SusdCommitNewAdmin)
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
		it.Event = new(SusdCommitNewAdmin)
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
func (it *SusdCommitNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SusdCommitNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SusdCommitNewAdmin represents a CommitNewAdmin event raised by the Susd contract.
type SusdCommitNewAdmin struct {
	Deadline *big.Int
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewAdmin is a free log retrieval operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Susd *SusdFilterer) FilterCommitNewAdmin(opts *bind.FilterOpts, deadline []*big.Int, admin []common.Address) (*SusdCommitNewAdminIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Susd.contract.FilterLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &SusdCommitNewAdminIterator{contract: _Susd.contract, event: "CommitNewAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitNewAdmin is a free log subscription operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Susd *SusdFilterer) WatchCommitNewAdmin(opts *bind.WatchOpts, sink chan<- *SusdCommitNewAdmin, deadline []*big.Int, admin []common.Address) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Susd.contract.WatchLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SusdCommitNewAdmin)
				if err := _Susd.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
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
func (_Susd *SusdFilterer) ParseCommitNewAdmin(log types.Log) (*SusdCommitNewAdmin, error) {
	event := new(SusdCommitNewAdmin)
	if err := _Susd.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SusdCommitNewParametersIterator is returned from FilterCommitNewParameters and is used to iterate over the raw logs and unpacked data for CommitNewParameters events raised by the Susd contract.
type SusdCommitNewParametersIterator struct {
	Event *SusdCommitNewParameters // Event containing the contract specifics and raw log

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
func (it *SusdCommitNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SusdCommitNewParameters)
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
		it.Event = new(SusdCommitNewParameters)
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
func (it *SusdCommitNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SusdCommitNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SusdCommitNewParameters represents a CommitNewParameters event raised by the Susd contract.
type SusdCommitNewParameters struct {
	Deadline *big.Int
	A        *big.Int
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewParameters is a free log retrieval operation binding the contract event 0x6081daa3b61098baf24d9c69bcd53af932e0635c89c6fd0617534b9ba76a7f73.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 A, uint256 fee, uint256 admin_fee)
func (_Susd *SusdFilterer) FilterCommitNewParameters(opts *bind.FilterOpts, deadline []*big.Int) (*SusdCommitNewParametersIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Susd.contract.FilterLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return &SusdCommitNewParametersIterator{contract: _Susd.contract, event: "CommitNewParameters", logs: logs, sub: sub}, nil
}

// WatchCommitNewParameters is a free log subscription operation binding the contract event 0x6081daa3b61098baf24d9c69bcd53af932e0635c89c6fd0617534b9ba76a7f73.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 A, uint256 fee, uint256 admin_fee)
func (_Susd *SusdFilterer) WatchCommitNewParameters(opts *bind.WatchOpts, sink chan<- *SusdCommitNewParameters, deadline []*big.Int) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Susd.contract.WatchLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SusdCommitNewParameters)
				if err := _Susd.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
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
func (_Susd *SusdFilterer) ParseCommitNewParameters(log types.Log) (*SusdCommitNewParameters, error) {
	event := new(SusdCommitNewParameters)
	if err := _Susd.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SusdNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Susd contract.
type SusdNewAdminIterator struct {
	Event *SusdNewAdmin // Event containing the contract specifics and raw log

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
func (it *SusdNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SusdNewAdmin)
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
		it.Event = new(SusdNewAdmin)
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
func (it *SusdNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SusdNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SusdNewAdmin represents a NewAdmin event raised by the Susd contract.
type SusdNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Susd *SusdFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address) (*SusdNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Susd.contract.FilterLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &SusdNewAdminIterator{contract: _Susd.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Susd *SusdFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *SusdNewAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Susd.contract.WatchLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SusdNewAdmin)
				if err := _Susd.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_Susd *SusdFilterer) ParseNewAdmin(log types.Log) (*SusdNewAdmin, error) {
	event := new(SusdNewAdmin)
	if err := _Susd.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SusdNewParametersIterator is returned from FilterNewParameters and is used to iterate over the raw logs and unpacked data for NewParameters events raised by the Susd contract.
type SusdNewParametersIterator struct {
	Event *SusdNewParameters // Event containing the contract specifics and raw log

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
func (it *SusdNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SusdNewParameters)
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
		it.Event = new(SusdNewParameters)
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
func (it *SusdNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SusdNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SusdNewParameters represents a NewParameters event raised by the Susd contract.
type SusdNewParameters struct {
	A        *big.Int
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewParameters is a free log retrieval operation binding the contract event 0x752a27d1853eb7af3ee4ff764f2c4a51619386af721573dd3809e929c39db99e.
//
// Solidity: event NewParameters(uint256 A, uint256 fee, uint256 admin_fee)
func (_Susd *SusdFilterer) FilterNewParameters(opts *bind.FilterOpts) (*SusdNewParametersIterator, error) {

	logs, sub, err := _Susd.contract.FilterLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return &SusdNewParametersIterator{contract: _Susd.contract, event: "NewParameters", logs: logs, sub: sub}, nil
}

// WatchNewParameters is a free log subscription operation binding the contract event 0x752a27d1853eb7af3ee4ff764f2c4a51619386af721573dd3809e929c39db99e.
//
// Solidity: event NewParameters(uint256 A, uint256 fee, uint256 admin_fee)
func (_Susd *SusdFilterer) WatchNewParameters(opts *bind.WatchOpts, sink chan<- *SusdNewParameters) (event.Subscription, error) {

	logs, sub, err := _Susd.contract.WatchLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SusdNewParameters)
				if err := _Susd.contract.UnpackLog(event, "NewParameters", log); err != nil {
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
func (_Susd *SusdFilterer) ParseNewParameters(log types.Log) (*SusdNewParameters, error) {
	event := new(SusdNewParameters)
	if err := _Susd.contract.UnpackLog(event, "NewParameters", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SusdRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Susd contract.
type SusdRemoveLiquidityIterator struct {
	Event *SusdRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *SusdRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SusdRemoveLiquidity)
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
		it.Event = new(SusdRemoveLiquidity)
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
func (it *SusdRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SusdRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SusdRemoveLiquidity represents a RemoveLiquidity event raised by the Susd contract.
type SusdRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [4]*big.Int
	Fees         [4]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x9878ca375e106f2a43c3b599fc624568131c4c9a4ba66a14563715763be9d59d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 token_supply)
func (_Susd *SusdFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*SusdRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Susd.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &SusdRemoveLiquidityIterator{contract: _Susd.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x9878ca375e106f2a43c3b599fc624568131c4c9a4ba66a14563715763be9d59d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 token_supply)
func (_Susd *SusdFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *SusdRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Susd.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SusdRemoveLiquidity)
				if err := _Susd.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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
func (_Susd *SusdFilterer) ParseRemoveLiquidity(log types.Log) (*SusdRemoveLiquidity, error) {
	event := new(SusdRemoveLiquidity)
	if err := _Susd.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SusdRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the Susd contract.
type SusdRemoveLiquidityImbalanceIterator struct {
	Event *SusdRemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *SusdRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SusdRemoveLiquidityImbalance)
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
		it.Event = new(SusdRemoveLiquidityImbalance)
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
func (it *SusdRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SusdRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SusdRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the Susd contract.
type SusdRemoveLiquidityImbalance struct {
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
func (_Susd *SusdFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*SusdRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Susd.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &SusdRemoveLiquidityImbalanceIterator{contract: _Susd.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0xb964b72f73f5ef5bf0fdc559b2fab9a7b12a39e47817a547f1f0aee47febd602.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_Susd *SusdFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *SusdRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Susd.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SusdRemoveLiquidityImbalance)
				if err := _Susd.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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
func (_Susd *SusdFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*SusdRemoveLiquidityImbalance, error) {
	event := new(SusdRemoveLiquidityImbalance)
	if err := _Susd.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SusdTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the Susd contract.
type SusdTokenExchangeIterator struct {
	Event *SusdTokenExchange // Event containing the contract specifics and raw log

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
func (it *SusdTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SusdTokenExchange)
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
		it.Event = new(SusdTokenExchange)
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
func (it *SusdTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SusdTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SusdTokenExchange represents a TokenExchange event raised by the Susd contract.
type SusdTokenExchange struct {
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
func (_Susd *SusdFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*SusdTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Susd.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &SusdTokenExchangeIterator{contract: _Susd.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Susd *SusdFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *SusdTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Susd.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SusdTokenExchange)
				if err := _Susd.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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
func (_Susd *SusdFilterer) ParseTokenExchange(log types.Log) (*SusdTokenExchange, error) {
	event := new(SusdTokenExchange)
	if err := _Susd.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SusdTokenExchangeUnderlyingIterator is returned from FilterTokenExchangeUnderlying and is used to iterate over the raw logs and unpacked data for TokenExchangeUnderlying events raised by the Susd contract.
type SusdTokenExchangeUnderlyingIterator struct {
	Event *SusdTokenExchangeUnderlying // Event containing the contract specifics and raw log

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
func (it *SusdTokenExchangeUnderlyingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SusdTokenExchangeUnderlying)
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
		it.Event = new(SusdTokenExchangeUnderlying)
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
func (it *SusdTokenExchangeUnderlyingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SusdTokenExchangeUnderlyingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SusdTokenExchangeUnderlying represents a TokenExchangeUnderlying event raised by the Susd contract.
type SusdTokenExchangeUnderlying struct {
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
func (_Susd *SusdFilterer) FilterTokenExchangeUnderlying(opts *bind.FilterOpts, buyer []common.Address) (*SusdTokenExchangeUnderlyingIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Susd.contract.FilterLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return &SusdTokenExchangeUnderlyingIterator{contract: _Susd.contract, event: "TokenExchangeUnderlying", logs: logs, sub: sub}, nil
}

// WatchTokenExchangeUnderlying is a free log subscription operation binding the contract event 0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b.
//
// Solidity: event TokenExchangeUnderlying(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Susd *SusdFilterer) WatchTokenExchangeUnderlying(opts *bind.WatchOpts, sink chan<- *SusdTokenExchangeUnderlying, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Susd.contract.WatchLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SusdTokenExchangeUnderlying)
				if err := _Susd.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
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
func (_Susd *SusdFilterer) ParseTokenExchangeUnderlying(log types.Log) (*SusdTokenExchangeUnderlying, error) {
	event := new(SusdTokenExchangeUnderlying)
	if err := _Susd.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
		return nil, err
	}
	return event, nil
}
