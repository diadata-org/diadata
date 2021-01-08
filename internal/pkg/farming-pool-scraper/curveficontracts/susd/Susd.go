// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package suds

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

// SudsABI is the input ABI used to generate the binding from.
const SudsABI = "[{\"name\":\"TokenExchange\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"sold_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"bought_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchangeUnderlying\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"sold_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"bought_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[4]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[4]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[4]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[4]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[4]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[4]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true,\"unit\":\"sec\"},{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewParameters\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true,\"unit\":\"sec\"},{\"type\":\"uint256\",\"name\":\"A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewParameters\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address[4]\",\"name\":\"_coins\"},{\"type\":\"address[4]\",\"name\":\"_underlying_coins\"},{\"type\":\"address\",\"name\":\"_pool_token\"},{\"type\":\"uint256\",\"name\":\"_A\"},{\"type\":\"uint256\",\"name\":\"_fee\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"get_virtual_price\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1570535},{\"name\":\"calc_token_amount\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256[4]\",\"name\":\"amounts\"},{\"type\":\"bool\",\"name\":\"deposit\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6103471},{\"name\":\"add_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[4]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"min_mint_amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":9331701},{\"name\":\"get_dy\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":3489637},{\"name\":\"get_dy_underlying\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":3489467},{\"name\":\"exchange\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"},{\"type\":\"uint256\",\"name\":\"min_dy\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":7034253},{\"name\":\"exchange_underlying\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"},{\"type\":\"uint256\",\"name\":\"min_dy\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":7050488},{\"name\":\"remove_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"},{\"type\":\"uint256[4]\",\"name\":\"min_amounts\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":241191},{\"name\":\"remove_liquidity_imbalance\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[4]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"max_burn_amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":9330864},{\"name\":\"commit_new_parameters\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"amplification\"},{\"type\":\"uint256\",\"name\":\"new_fee\"},{\"type\":\"uint256\",\"name\":\"new_admin_fee\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":146045},{\"name\":\"apply_new_parameters\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":133452},{\"name\":\"revert_new_parameters\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":21775},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":74452},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":60508},{\"name\":\"revert_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":21865},{\"name\":\"withdraw_admin_fees\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":23448},{\"name\":\"kill_me\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":37818},{\"name\":\"unkill_me\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":21955},{\"name\":\"coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2130},{\"name\":\"underlying_coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2160},{\"name\":\"balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2190},{\"name\":\"A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2021},{\"name\":\"fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2051},{\"name\":\"admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2081},{\"name\":\"owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2111},{\"name\":\"admin_actions_deadline\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"sec\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2141},{\"name\":\"transfer_ownership_deadline\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"sec\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2171},{\"name\":\"future_A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2201},{\"name\":\"future_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2231},{\"name\":\"future_admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2261},{\"name\":\"future_owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2291}]"

// Suds is an auto generated Go binding around an Ethereum contract.
type Suds struct {
	SudsCaller     // Read-only binding to the contract
	SudsTransactor // Write-only binding to the contract
	SudsFilterer   // Log filterer for contract events
}

// SudsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SudsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SudsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SudsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SudsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SudsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SudsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SudsSession struct {
	Contract     *Suds             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SudsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SudsCallerSession struct {
	Contract *SudsCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SudsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SudsTransactorSession struct {
	Contract     *SudsTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SudsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SudsRaw struct {
	Contract *Suds // Generic contract binding to access the raw methods on
}

// SudsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SudsCallerRaw struct {
	Contract *SudsCaller // Generic read-only contract binding to access the raw methods on
}

// SudsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SudsTransactorRaw struct {
	Contract *SudsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSuds creates a new instance of Suds, bound to a specific deployed contract.
func NewSuds(address common.Address, backend bind.ContractBackend) (*Suds, error) {
	contract, err := bindSuds(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Suds{SudsCaller: SudsCaller{contract: contract}, SudsTransactor: SudsTransactor{contract: contract}, SudsFilterer: SudsFilterer{contract: contract}}, nil
}

// NewSudsCaller creates a new read-only instance of Suds, bound to a specific deployed contract.
func NewSudsCaller(address common.Address, caller bind.ContractCaller) (*SudsCaller, error) {
	contract, err := bindSuds(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SudsCaller{contract: contract}, nil
}

// NewSudsTransactor creates a new write-only instance of Suds, bound to a specific deployed contract.
func NewSudsTransactor(address common.Address, transactor bind.ContractTransactor) (*SudsTransactor, error) {
	contract, err := bindSuds(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SudsTransactor{contract: contract}, nil
}

// NewSudsFilterer creates a new log filterer instance of Suds, bound to a specific deployed contract.
func NewSudsFilterer(address common.Address, filterer bind.ContractFilterer) (*SudsFilterer, error) {
	contract, err := bindSuds(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SudsFilterer{contract: contract}, nil
}

// bindSuds binds a generic wrapper to an already deployed contract.
func bindSuds(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SudsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Suds *SudsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Suds.Contract.SudsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Suds *SudsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Suds.Contract.SudsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Suds *SudsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Suds.Contract.SudsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Suds *SudsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Suds.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Suds *SudsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Suds.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Suds *SudsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Suds.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() returns(uint256)
func (_Suds *SudsCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "A")
	return *ret0, err
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() returns(uint256)
func (_Suds *SudsSession) A() (*big.Int, error) {
	return _Suds.Contract.A(&_Suds.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() returns(uint256)
func (_Suds *SudsCallerSession) A() (*big.Int, error) {
	return _Suds.Contract.A(&_Suds.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() returns(uint256)
func (_Suds *SudsCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "admin_actions_deadline")
	return *ret0, err
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() returns(uint256)
func (_Suds *SudsSession) AdminActionsDeadline() (*big.Int, error) {
	return _Suds.Contract.AdminActionsDeadline(&_Suds.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() returns(uint256)
func (_Suds *SudsCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _Suds.Contract.AdminActionsDeadline(&_Suds.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() returns(uint256)
func (_Suds *SudsCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "admin_fee")
	return *ret0, err
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() returns(uint256)
func (_Suds *SudsSession) AdminFee() (*big.Int, error) {
	return _Suds.Contract.AdminFee(&_Suds.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() returns(uint256)
func (_Suds *SudsCallerSession) AdminFee() (*big.Int, error) {
	return _Suds.Contract.AdminFee(&_Suds.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 arg0) returns(uint256)
func (_Suds *SudsCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 arg0) returns(uint256)
func (_Suds *SudsSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Suds.Contract.Balances(&_Suds.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 arg0) returns(uint256)
func (_Suds *SudsCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Suds.Contract.Balances(&_Suds.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xcf701ff7.
//
// Solidity: function calc_token_amount(uint256[4] amounts, bool deposit) returns(uint256)
func (_Suds *SudsCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [4]*big.Int, deposit bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "calc_token_amount", amounts, deposit)
	return *ret0, err
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xcf701ff7.
//
// Solidity: function calc_token_amount(uint256[4] amounts, bool deposit) returns(uint256)
func (_Suds *SudsSession) CalcTokenAmount(amounts [4]*big.Int, deposit bool) (*big.Int, error) {
	return _Suds.Contract.CalcTokenAmount(&_Suds.CallOpts, amounts, deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xcf701ff7.
//
// Solidity: function calc_token_amount(uint256[4] amounts, bool deposit) returns(uint256)
func (_Suds *SudsCallerSession) CalcTokenAmount(amounts [4]*big.Int, deposit bool) (*big.Int, error) {
	return _Suds.Contract.CalcTokenAmount(&_Suds.CallOpts, amounts, deposit)
}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 arg0) returns(address)
func (_Suds *SudsCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "coins", arg0)
	return *ret0, err
}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 arg0) returns(address)
func (_Suds *SudsSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Suds.Contract.Coins(&_Suds.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 arg0) returns(address)
func (_Suds *SudsCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Suds.Contract.Coins(&_Suds.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() returns(uint256)
func (_Suds *SudsCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() returns(uint256)
func (_Suds *SudsSession) Fee() (*big.Int, error) {
	return _Suds.Contract.Fee(&_Suds.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() returns(uint256)
func (_Suds *SudsCallerSession) Fee() (*big.Int, error) {
	return _Suds.Contract.Fee(&_Suds.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() returns(uint256)
func (_Suds *SudsCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "future_A")
	return *ret0, err
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() returns(uint256)
func (_Suds *SudsSession) FutureA() (*big.Int, error) {
	return _Suds.Contract.FutureA(&_Suds.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() returns(uint256)
func (_Suds *SudsCallerSession) FutureA() (*big.Int, error) {
	return _Suds.Contract.FutureA(&_Suds.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() returns(uint256)
func (_Suds *SudsCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "future_admin_fee")
	return *ret0, err
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() returns(uint256)
func (_Suds *SudsSession) FutureAdminFee() (*big.Int, error) {
	return _Suds.Contract.FutureAdminFee(&_Suds.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() returns(uint256)
func (_Suds *SudsCallerSession) FutureAdminFee() (*big.Int, error) {
	return _Suds.Contract.FutureAdminFee(&_Suds.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() returns(uint256)
func (_Suds *SudsCaller) FutureFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "future_fee")
	return *ret0, err
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() returns(uint256)
func (_Suds *SudsSession) FutureFee() (*big.Int, error) {
	return _Suds.Contract.FutureFee(&_Suds.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() returns(uint256)
func (_Suds *SudsCallerSession) FutureFee() (*big.Int, error) {
	return _Suds.Contract.FutureFee(&_Suds.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() returns(address)
func (_Suds *SudsCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "future_owner")
	return *ret0, err
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() returns(address)
func (_Suds *SudsSession) FutureOwner() (common.Address, error) {
	return _Suds.Contract.FutureOwner(&_Suds.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() returns(address)
func (_Suds *SudsCallerSession) FutureOwner() (common.Address, error) {
	return _Suds.Contract.FutureOwner(&_Suds.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Suds *SudsCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "get_dy", i, j, dx)
	return *ret0, err
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Suds *SudsSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Suds.Contract.GetDy(&_Suds.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Suds *SudsCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Suds.Contract.GetDy(&_Suds.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Suds *SudsCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "get_dy_underlying", i, j, dx)
	return *ret0, err
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Suds *SudsSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Suds.Contract.GetDyUnderlying(&_Suds.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) returns(uint256)
func (_Suds *SudsCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Suds.Contract.GetDyUnderlying(&_Suds.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() returns(uint256)
func (_Suds *SudsCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "get_virtual_price")
	return *ret0, err
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() returns(uint256)
func (_Suds *SudsSession) GetVirtualPrice() (*big.Int, error) {
	return _Suds.Contract.GetVirtualPrice(&_Suds.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() returns(uint256)
func (_Suds *SudsCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _Suds.Contract.GetVirtualPrice(&_Suds.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Suds *SudsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Suds *SudsSession) Owner() (common.Address, error) {
	return _Suds.Contract.Owner(&_Suds.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Suds *SudsCallerSession) Owner() (common.Address, error) {
	return _Suds.Contract.Owner(&_Suds.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() returns(uint256)
func (_Suds *SudsCaller) TransferOwnershipDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "transfer_ownership_deadline")
	return *ret0, err
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() returns(uint256)
func (_Suds *SudsSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _Suds.Contract.TransferOwnershipDeadline(&_Suds.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() returns(uint256)
func (_Suds *SudsCallerSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _Suds.Contract.TransferOwnershipDeadline(&_Suds.CallOpts)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 arg0) returns(address)
func (_Suds *SudsCaller) UnderlyingCoins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Suds.contract.Call(opts, out, "underlying_coins", arg0)
	return *ret0, err
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 arg0) returns(address)
func (_Suds *SudsSession) UnderlyingCoins(arg0 *big.Int) (common.Address, error) {
	return _Suds.Contract.UnderlyingCoins(&_Suds.CallOpts, arg0)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 arg0) returns(address)
func (_Suds *SudsCallerSession) UnderlyingCoins(arg0 *big.Int) (common.Address, error) {
	return _Suds.Contract.UnderlyingCoins(&_Suds.CallOpts, arg0)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x029b2f34.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 min_mint_amount) returns()
func (_Suds *SudsTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [4]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Suds.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x029b2f34.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 min_mint_amount) returns()
func (_Suds *SudsSession) AddLiquidity(amounts [4]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Suds.Contract.AddLiquidity(&_Suds.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x029b2f34.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 min_mint_amount) returns()
func (_Suds *SudsTransactorSession) AddLiquidity(amounts [4]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Suds.Contract.AddLiquidity(&_Suds.TransactOpts, amounts, min_mint_amount)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Suds *SudsTransactor) ApplyNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Suds.contract.Transact(opts, "apply_new_parameters")
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Suds *SudsSession) ApplyNewParameters() (*types.Transaction, error) {
	return _Suds.Contract.ApplyNewParameters(&_Suds.TransactOpts)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Suds *SudsTransactorSession) ApplyNewParameters() (*types.Transaction, error) {
	return _Suds.Contract.ApplyNewParameters(&_Suds.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Suds *SudsTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Suds.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Suds *SudsSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Suds.Contract.ApplyTransferOwnership(&_Suds.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Suds *SudsTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Suds.Contract.ApplyTransferOwnership(&_Suds.TransactOpts)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xee11f5b6.
//
// Solidity: function commit_new_parameters(uint256 amplification, uint256 new_fee, uint256 new_admin_fee) returns()
func (_Suds *SudsTransactor) CommitNewParameters(opts *bind.TransactOpts, amplification *big.Int, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Suds.contract.Transact(opts, "commit_new_parameters", amplification, new_fee, new_admin_fee)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xee11f5b6.
//
// Solidity: function commit_new_parameters(uint256 amplification, uint256 new_fee, uint256 new_admin_fee) returns()
func (_Suds *SudsSession) CommitNewParameters(amplification *big.Int, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Suds.Contract.CommitNewParameters(&_Suds.TransactOpts, amplification, new_fee, new_admin_fee)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xee11f5b6.
//
// Solidity: function commit_new_parameters(uint256 amplification, uint256 new_fee, uint256 new_admin_fee) returns()
func (_Suds *SudsTransactorSession) CommitNewParameters(amplification *big.Int, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Suds.Contract.CommitNewParameters(&_Suds.TransactOpts, amplification, new_fee, new_admin_fee)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Suds *SudsTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Suds.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Suds *SudsSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Suds.Contract.CommitTransferOwnership(&_Suds.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Suds *SudsTransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Suds.Contract.CommitTransferOwnership(&_Suds.TransactOpts, _owner)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Suds *SudsTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Suds.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Suds *SudsSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Suds.Contract.Exchange(&_Suds.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Suds *SudsTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Suds.Contract.Exchange(&_Suds.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Suds *SudsTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Suds.contract.Transact(opts, "exchange_underlying", i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Suds *SudsSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Suds.Contract.ExchangeUnderlying(&_Suds.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Suds *SudsTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Suds.Contract.ExchangeUnderlying(&_Suds.TransactOpts, i, j, dx, min_dy)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Suds *SudsTransactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Suds.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Suds *SudsSession) KillMe() (*types.Transaction, error) {
	return _Suds.Contract.KillMe(&_Suds.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Suds *SudsTransactorSession) KillMe() (*types.Transaction, error) {
	return _Suds.Contract.KillMe(&_Suds.TransactOpts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x7d49d875.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[4] min_amounts) returns()
func (_Suds *SudsTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [4]*big.Int) (*types.Transaction, error) {
	return _Suds.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x7d49d875.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[4] min_amounts) returns()
func (_Suds *SudsSession) RemoveLiquidity(_amount *big.Int, min_amounts [4]*big.Int) (*types.Transaction, error) {
	return _Suds.Contract.RemoveLiquidity(&_Suds.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x7d49d875.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[4] min_amounts) returns()
func (_Suds *SudsTransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [4]*big.Int) (*types.Transaction, error) {
	return _Suds.Contract.RemoveLiquidity(&_Suds.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x18a7bd76.
//
// Solidity: function remove_liquidity_imbalance(uint256[4] amounts, uint256 max_burn_amount) returns()
func (_Suds *SudsTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, amounts [4]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Suds.contract.Transact(opts, "remove_liquidity_imbalance", amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x18a7bd76.
//
// Solidity: function remove_liquidity_imbalance(uint256[4] amounts, uint256 max_burn_amount) returns()
func (_Suds *SudsSession) RemoveLiquidityImbalance(amounts [4]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Suds.Contract.RemoveLiquidityImbalance(&_Suds.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x18a7bd76.
//
// Solidity: function remove_liquidity_imbalance(uint256[4] amounts, uint256 max_burn_amount) returns()
func (_Suds *SudsTransactorSession) RemoveLiquidityImbalance(amounts [4]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Suds.Contract.RemoveLiquidityImbalance(&_Suds.TransactOpts, amounts, max_burn_amount)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Suds *SudsTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Suds.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Suds *SudsSession) RevertNewParameters() (*types.Transaction, error) {
	return _Suds.Contract.RevertNewParameters(&_Suds.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Suds *SudsTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _Suds.Contract.RevertNewParameters(&_Suds.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Suds *SudsTransactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Suds.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Suds *SudsSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Suds.Contract.RevertTransferOwnership(&_Suds.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Suds *SudsTransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Suds.Contract.RevertTransferOwnership(&_Suds.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Suds *SudsTransactor) UnkillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Suds.contract.Transact(opts, "unkill_me")
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Suds *SudsSession) UnkillMe() (*types.Transaction, error) {
	return _Suds.Contract.UnkillMe(&_Suds.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Suds *SudsTransactorSession) UnkillMe() (*types.Transaction, error) {
	return _Suds.Contract.UnkillMe(&_Suds.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Suds *SudsTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Suds.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Suds *SudsSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Suds.Contract.WithdrawAdminFees(&_Suds.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Suds *SudsTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Suds.Contract.WithdrawAdminFees(&_Suds.TransactOpts)
}

// SudsAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Suds contract.
type SudsAddLiquidityIterator struct {
	Event *SudsAddLiquidity // Event containing the contract specifics and raw log

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
func (it *SudsAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SudsAddLiquidity)
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
		it.Event = new(SudsAddLiquidity)
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
func (it *SudsAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SudsAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SudsAddLiquidity represents a AddLiquidity event raised by the Suds contract.
type SudsAddLiquidity struct {
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
func (_Suds *SudsFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*SudsAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Suds.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &SudsAddLiquidityIterator{contract: _Suds.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x3f1915775e0c9a38a57a7bb7f1f9005f486fb904e1f84aa215364d567319a58d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_Suds *SudsFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *SudsAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Suds.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SudsAddLiquidity)
				if err := _Suds.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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
func (_Suds *SudsFilterer) ParseAddLiquidity(log types.Log) (*SudsAddLiquidity, error) {
	event := new(SudsAddLiquidity)
	if err := _Suds.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SudsCommitNewAdminIterator is returned from FilterCommitNewAdmin and is used to iterate over the raw logs and unpacked data for CommitNewAdmin events raised by the Suds contract.
type SudsCommitNewAdminIterator struct {
	Event *SudsCommitNewAdmin // Event containing the contract specifics and raw log

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
func (it *SudsCommitNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SudsCommitNewAdmin)
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
		it.Event = new(SudsCommitNewAdmin)
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
func (it *SudsCommitNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SudsCommitNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SudsCommitNewAdmin represents a CommitNewAdmin event raised by the Suds contract.
type SudsCommitNewAdmin struct {
	Deadline *big.Int
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewAdmin is a free log retrieval operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Suds *SudsFilterer) FilterCommitNewAdmin(opts *bind.FilterOpts, deadline []*big.Int, admin []common.Address) (*SudsCommitNewAdminIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Suds.contract.FilterLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &SudsCommitNewAdminIterator{contract: _Suds.contract, event: "CommitNewAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitNewAdmin is a free log subscription operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Suds *SudsFilterer) WatchCommitNewAdmin(opts *bind.WatchOpts, sink chan<- *SudsCommitNewAdmin, deadline []*big.Int, admin []common.Address) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Suds.contract.WatchLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SudsCommitNewAdmin)
				if err := _Suds.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
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
func (_Suds *SudsFilterer) ParseCommitNewAdmin(log types.Log) (*SudsCommitNewAdmin, error) {
	event := new(SudsCommitNewAdmin)
	if err := _Suds.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SudsCommitNewParametersIterator is returned from FilterCommitNewParameters and is used to iterate over the raw logs and unpacked data for CommitNewParameters events raised by the Suds contract.
type SudsCommitNewParametersIterator struct {
	Event *SudsCommitNewParameters // Event containing the contract specifics and raw log

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
func (it *SudsCommitNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SudsCommitNewParameters)
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
		it.Event = new(SudsCommitNewParameters)
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
func (it *SudsCommitNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SudsCommitNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SudsCommitNewParameters represents a CommitNewParameters event raised by the Suds contract.
type SudsCommitNewParameters struct {
	Deadline *big.Int
	A        *big.Int
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewParameters is a free log retrieval operation binding the contract event 0x6081daa3b61098baf24d9c69bcd53af932e0635c89c6fd0617534b9ba76a7f73.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 A, uint256 fee, uint256 admin_fee)
func (_Suds *SudsFilterer) FilterCommitNewParameters(opts *bind.FilterOpts, deadline []*big.Int) (*SudsCommitNewParametersIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Suds.contract.FilterLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return &SudsCommitNewParametersIterator{contract: _Suds.contract, event: "CommitNewParameters", logs: logs, sub: sub}, nil
}

// WatchCommitNewParameters is a free log subscription operation binding the contract event 0x6081daa3b61098baf24d9c69bcd53af932e0635c89c6fd0617534b9ba76a7f73.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 A, uint256 fee, uint256 admin_fee)
func (_Suds *SudsFilterer) WatchCommitNewParameters(opts *bind.WatchOpts, sink chan<- *SudsCommitNewParameters, deadline []*big.Int) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Suds.contract.WatchLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SudsCommitNewParameters)
				if err := _Suds.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
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
func (_Suds *SudsFilterer) ParseCommitNewParameters(log types.Log) (*SudsCommitNewParameters, error) {
	event := new(SudsCommitNewParameters)
	if err := _Suds.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SudsNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Suds contract.
type SudsNewAdminIterator struct {
	Event *SudsNewAdmin // Event containing the contract specifics and raw log

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
func (it *SudsNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SudsNewAdmin)
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
		it.Event = new(SudsNewAdmin)
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
func (it *SudsNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SudsNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SudsNewAdmin represents a NewAdmin event raised by the Suds contract.
type SudsNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Suds *SudsFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address) (*SudsNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Suds.contract.FilterLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &SudsNewAdminIterator{contract: _Suds.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Suds *SudsFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *SudsNewAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Suds.contract.WatchLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SudsNewAdmin)
				if err := _Suds.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_Suds *SudsFilterer) ParseNewAdmin(log types.Log) (*SudsNewAdmin, error) {
	event := new(SudsNewAdmin)
	if err := _Suds.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SudsNewParametersIterator is returned from FilterNewParameters and is used to iterate over the raw logs and unpacked data for NewParameters events raised by the Suds contract.
type SudsNewParametersIterator struct {
	Event *SudsNewParameters // Event containing the contract specifics and raw log

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
func (it *SudsNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SudsNewParameters)
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
		it.Event = new(SudsNewParameters)
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
func (it *SudsNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SudsNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SudsNewParameters represents a NewParameters event raised by the Suds contract.
type SudsNewParameters struct {
	A        *big.Int
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewParameters is a free log retrieval operation binding the contract event 0x752a27d1853eb7af3ee4ff764f2c4a51619386af721573dd3809e929c39db99e.
//
// Solidity: event NewParameters(uint256 A, uint256 fee, uint256 admin_fee)
func (_Suds *SudsFilterer) FilterNewParameters(opts *bind.FilterOpts) (*SudsNewParametersIterator, error) {

	logs, sub, err := _Suds.contract.FilterLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return &SudsNewParametersIterator{contract: _Suds.contract, event: "NewParameters", logs: logs, sub: sub}, nil
}

// WatchNewParameters is a free log subscription operation binding the contract event 0x752a27d1853eb7af3ee4ff764f2c4a51619386af721573dd3809e929c39db99e.
//
// Solidity: event NewParameters(uint256 A, uint256 fee, uint256 admin_fee)
func (_Suds *SudsFilterer) WatchNewParameters(opts *bind.WatchOpts, sink chan<- *SudsNewParameters) (event.Subscription, error) {

	logs, sub, err := _Suds.contract.WatchLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SudsNewParameters)
				if err := _Suds.contract.UnpackLog(event, "NewParameters", log); err != nil {
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
func (_Suds *SudsFilterer) ParseNewParameters(log types.Log) (*SudsNewParameters, error) {
	event := new(SudsNewParameters)
	if err := _Suds.contract.UnpackLog(event, "NewParameters", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SudsRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Suds contract.
type SudsRemoveLiquidityIterator struct {
	Event *SudsRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *SudsRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SudsRemoveLiquidity)
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
		it.Event = new(SudsRemoveLiquidity)
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
func (it *SudsRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SudsRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SudsRemoveLiquidity represents a RemoveLiquidity event raised by the Suds contract.
type SudsRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [4]*big.Int
	Fees         [4]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x9878ca375e106f2a43c3b599fc624568131c4c9a4ba66a14563715763be9d59d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 token_supply)
func (_Suds *SudsFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*SudsRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Suds.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &SudsRemoveLiquidityIterator{contract: _Suds.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x9878ca375e106f2a43c3b599fc624568131c4c9a4ba66a14563715763be9d59d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 token_supply)
func (_Suds *SudsFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *SudsRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Suds.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SudsRemoveLiquidity)
				if err := _Suds.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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
func (_Suds *SudsFilterer) ParseRemoveLiquidity(log types.Log) (*SudsRemoveLiquidity, error) {
	event := new(SudsRemoveLiquidity)
	if err := _Suds.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SudsRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the Suds contract.
type SudsRemoveLiquidityImbalanceIterator struct {
	Event *SudsRemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *SudsRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SudsRemoveLiquidityImbalance)
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
		it.Event = new(SudsRemoveLiquidityImbalance)
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
func (it *SudsRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SudsRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SudsRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the Suds contract.
type SudsRemoveLiquidityImbalance struct {
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
func (_Suds *SudsFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*SudsRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Suds.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &SudsRemoveLiquidityImbalanceIterator{contract: _Suds.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0xb964b72f73f5ef5bf0fdc559b2fab9a7b12a39e47817a547f1f0aee47febd602.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_Suds *SudsFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *SudsRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Suds.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SudsRemoveLiquidityImbalance)
				if err := _Suds.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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
func (_Suds *SudsFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*SudsRemoveLiquidityImbalance, error) {
	event := new(SudsRemoveLiquidityImbalance)
	if err := _Suds.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SudsTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the Suds contract.
type SudsTokenExchangeIterator struct {
	Event *SudsTokenExchange // Event containing the contract specifics and raw log

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
func (it *SudsTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SudsTokenExchange)
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
		it.Event = new(SudsTokenExchange)
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
func (it *SudsTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SudsTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SudsTokenExchange represents a TokenExchange event raised by the Suds contract.
type SudsTokenExchange struct {
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
func (_Suds *SudsFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*SudsTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Suds.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &SudsTokenExchangeIterator{contract: _Suds.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Suds *SudsFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *SudsTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Suds.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SudsTokenExchange)
				if err := _Suds.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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
func (_Suds *SudsFilterer) ParseTokenExchange(log types.Log) (*SudsTokenExchange, error) {
	event := new(SudsTokenExchange)
	if err := _Suds.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SudsTokenExchangeUnderlyingIterator is returned from FilterTokenExchangeUnderlying and is used to iterate over the raw logs and unpacked data for TokenExchangeUnderlying events raised by the Suds contract.
type SudsTokenExchangeUnderlyingIterator struct {
	Event *SudsTokenExchangeUnderlying // Event containing the contract specifics and raw log

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
func (it *SudsTokenExchangeUnderlyingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SudsTokenExchangeUnderlying)
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
		it.Event = new(SudsTokenExchangeUnderlying)
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
func (it *SudsTokenExchangeUnderlyingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SudsTokenExchangeUnderlyingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SudsTokenExchangeUnderlying represents a TokenExchangeUnderlying event raised by the Suds contract.
type SudsTokenExchangeUnderlying struct {
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
func (_Suds *SudsFilterer) FilterTokenExchangeUnderlying(opts *bind.FilterOpts, buyer []common.Address) (*SudsTokenExchangeUnderlyingIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Suds.contract.FilterLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return &SudsTokenExchangeUnderlyingIterator{contract: _Suds.contract, event: "TokenExchangeUnderlying", logs: logs, sub: sub}, nil
}

// WatchTokenExchangeUnderlying is a free log subscription operation binding the contract event 0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b.
//
// Solidity: event TokenExchangeUnderlying(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Suds *SudsFilterer) WatchTokenExchangeUnderlying(opts *bind.WatchOpts, sink chan<- *SudsTokenExchangeUnderlying, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Suds.contract.WatchLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SudsTokenExchangeUnderlying)
				if err := _Suds.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
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
func (_Suds *SudsFilterer) ParseTokenExchangeUnderlying(log types.Log) (*SudsTokenExchangeUnderlying, error) {
	event := new(SudsTokenExchangeUnderlying)
	if err := _Suds.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
		return nil, err
	}
	return event, nil
}
