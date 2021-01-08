// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package special

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

// SpecialABI is the input ABI used to generate the binding from.
const SpecialABI = "[{\"name\":\"TokenExchange\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"sold_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"bought_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[2]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[2]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[2]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[2]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"token_amount\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"coin_amount\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[2]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[2]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true},{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampA\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"old_A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"new_A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"initial_time\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"future_time\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"t\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"},{\"type\":\"address[2]\",\"name\":\"_coins\"},{\"type\":\"address\",\"name\":\"_pool_token\"},{\"type\":\"uint256\",\"name\":\"_A\"},{\"type\":\"uint256\",\"name\":\"_fee\"},{\"type\":\"uint256\",\"name\":\"_admin_fee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":5227},{\"name\":\"get_virtual_price\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":955150},{\"name\":\"calc_token_amount\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256[2]\",\"name\":\"amounts\"},{\"type\":\"bool\",\"name\":\"deposit\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3797461},{\"name\":\"add_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[2]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"min_mint_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":5836477},{\"name\":\"get_dy\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2317363},{\"name\":\"get_dy_underlying\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2317058},{\"name\":\"exchange\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"},{\"type\":\"uint256\",\"name\":\"min_dy\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":2454671},{\"name\":\"remove_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"},{\"type\":\"uint256[2]\",\"name\":\"min_amounts\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":143246},{\"name\":\"remove_liquidity_imbalance\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[2]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"max_burn_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":5836308},{\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_token_amount\"},{\"type\":\"int128\",\"name\":\"i\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1102},{\"name\":\"remove_liquidity_one_coin\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_token_amount\"},{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"uint256\",\"name\":\"min_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":3660494},{\"name\":\"ramp_A\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_future_A\"},{\"type\":\"uint256\",\"name\":\"_future_time\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":151919},{\"name\":\"stop_ramp_A\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":148637},{\"name\":\"commit_new_fee\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"new_fee\"},{\"type\":\"uint256\",\"name\":\"new_admin_fee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":110461},{\"name\":\"apply_new_fee\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":97242},{\"name\":\"revert_new_parameters\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":21895},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74572},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":60710},{\"name\":\"revert_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":21985},{\"name\":\"admin_balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"i\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3481},{\"name\":\"withdraw_admin_fees\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":9218},{\"name\":\"donate_admin_fees\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74965},{\"name\":\"kill_me\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37998},{\"name\":\"unkill_me\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":22135},{\"name\":\"coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2220},{\"name\":\"balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2250},{\"name\":\"fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2171},{\"name\":\"admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2201},{\"name\":\"owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2231},{\"name\":\"initial_A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2261},{\"name\":\"future_A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2291},{\"name\":\"initial_A_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2321},{\"name\":\"future_A_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2351},{\"name\":\"admin_actions_deadline\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2381},{\"name\":\"transfer_ownership_deadline\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2411},{\"name\":\"future_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2441},{\"name\":\"future_admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2471},{\"name\":\"future_owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2501}]"

// Special is an auto generated Go binding around an Ethereum contract.
type Special struct {
	SpecialCaller     // Read-only binding to the contract
	SpecialTransactor // Write-only binding to the contract
	SpecialFilterer   // Log filterer for contract events
}

// SpecialCaller is an auto generated read-only Go binding around an Ethereum contract.
type SpecialCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpecialTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SpecialTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpecialFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SpecialFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpecialSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SpecialSession struct {
	Contract     *Special          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpecialCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SpecialCallerSession struct {
	Contract *SpecialCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SpecialTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SpecialTransactorSession struct {
	Contract     *SpecialTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SpecialRaw is an auto generated low-level Go binding around an Ethereum contract.
type SpecialRaw struct {
	Contract *Special // Generic contract binding to access the raw methods on
}

// SpecialCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SpecialCallerRaw struct {
	Contract *SpecialCaller // Generic read-only contract binding to access the raw methods on
}

// SpecialTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SpecialTransactorRaw struct {
	Contract *SpecialTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpecial creates a new instance of Special, bound to a specific deployed contract.
func NewSpecial(address common.Address, backend bind.ContractBackend) (*Special, error) {
	contract, err := bindSpecial(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Special{SpecialCaller: SpecialCaller{contract: contract}, SpecialTransactor: SpecialTransactor{contract: contract}, SpecialFilterer: SpecialFilterer{contract: contract}}, nil
}

// NewSpecialCaller creates a new read-only instance of Special, bound to a specific deployed contract.
func NewSpecialCaller(address common.Address, caller bind.ContractCaller) (*SpecialCaller, error) {
	contract, err := bindSpecial(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SpecialCaller{contract: contract}, nil
}

// NewSpecialTransactor creates a new write-only instance of Special, bound to a specific deployed contract.
func NewSpecialTransactor(address common.Address, transactor bind.ContractTransactor) (*SpecialTransactor, error) {
	contract, err := bindSpecial(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SpecialTransactor{contract: contract}, nil
}

// NewSpecialFilterer creates a new log filterer instance of Special, bound to a specific deployed contract.
func NewSpecialFilterer(address common.Address, filterer bind.ContractFilterer) (*SpecialFilterer, error) {
	contract, err := bindSpecial(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SpecialFilterer{contract: contract}, nil
}

// bindSpecial binds a generic wrapper to an already deployed contract.
func bindSpecial(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SpecialABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Special *SpecialRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Special.Contract.SpecialCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Special *SpecialRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Special.Contract.SpecialTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Special *SpecialRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Special.Contract.SpecialTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Special *SpecialCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Special.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Special *SpecialTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Special.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Special *SpecialTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Special.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Special *SpecialCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "A")
	return *ret0, err
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Special *SpecialSession) A() (*big.Int, error) {
	return _Special.Contract.A(&_Special.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Special *SpecialCallerSession) A() (*big.Int, error) {
	return _Special.Contract.A(&_Special.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_Special *SpecialCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "admin_actions_deadline")
	return *ret0, err
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_Special *SpecialSession) AdminActionsDeadline() (*big.Int, error) {
	return _Special.Contract.AdminActionsDeadline(&_Special.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_Special *SpecialCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _Special.Contract.AdminActionsDeadline(&_Special.CallOpts)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_Special *SpecialCaller) AdminBalances(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "admin_balances", i)
	return *ret0, err
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_Special *SpecialSession) AdminBalances(i *big.Int) (*big.Int, error) {
	return _Special.Contract.AdminBalances(&_Special.CallOpts, i)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_Special *SpecialCallerSession) AdminBalances(i *big.Int) (*big.Int, error) {
	return _Special.Contract.AdminBalances(&_Special.CallOpts, i)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Special *SpecialCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "admin_fee")
	return *ret0, err
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Special *SpecialSession) AdminFee() (*big.Int, error) {
	return _Special.Contract.AdminFee(&_Special.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Special *SpecialCallerSession) AdminFee() (*big.Int, error) {
	return _Special.Contract.AdminFee(&_Special.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Special *SpecialCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Special *SpecialSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Special.Contract.Balances(&_Special.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Special *SpecialCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Special.Contract.Balances(&_Special.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] amounts, bool deposit) view returns(uint256)
func (_Special *SpecialCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [2]*big.Int, deposit bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "calc_token_amount", amounts, deposit)
	return *ret0, err
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] amounts, bool deposit) view returns(uint256)
func (_Special *SpecialSession) CalcTokenAmount(amounts [2]*big.Int, deposit bool) (*big.Int, error) {
	return _Special.Contract.CalcTokenAmount(&_Special.CallOpts, amounts, deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] amounts, bool deposit) view returns(uint256)
func (_Special *SpecialCallerSession) CalcTokenAmount(amounts [2]*big.Int, deposit bool) (*big.Int, error) {
	return _Special.Contract.CalcTokenAmount(&_Special.CallOpts, amounts, deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_Special *SpecialCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, _token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "calc_withdraw_one_coin", _token_amount, i)
	return *ret0, err
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_Special *SpecialSession) CalcWithdrawOneCoin(_token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Special.Contract.CalcWithdrawOneCoin(&_Special.CallOpts, _token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_Special *SpecialCallerSession) CalcWithdrawOneCoin(_token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Special.Contract.CalcWithdrawOneCoin(&_Special.CallOpts, _token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Special *SpecialCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "coins", arg0)
	return *ret0, err
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Special *SpecialSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Special.Contract.Coins(&_Special.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Special *SpecialCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Special.Contract.Coins(&_Special.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Special *SpecialCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Special *SpecialSession) Fee() (*big.Int, error) {
	return _Special.Contract.Fee(&_Special.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Special *SpecialCallerSession) Fee() (*big.Int, error) {
	return _Special.Contract.Fee(&_Special.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_Special *SpecialCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "future_A")
	return *ret0, err
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_Special *SpecialSession) FutureA() (*big.Int, error) {
	return _Special.Contract.FutureA(&_Special.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_Special *SpecialCallerSession) FutureA() (*big.Int, error) {
	return _Special.Contract.FutureA(&_Special.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_Special *SpecialCaller) FutureATime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "future_A_time")
	return *ret0, err
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_Special *SpecialSession) FutureATime() (*big.Int, error) {
	return _Special.Contract.FutureATime(&_Special.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_Special *SpecialCallerSession) FutureATime() (*big.Int, error) {
	return _Special.Contract.FutureATime(&_Special.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_Special *SpecialCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "future_admin_fee")
	return *ret0, err
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_Special *SpecialSession) FutureAdminFee() (*big.Int, error) {
	return _Special.Contract.FutureAdminFee(&_Special.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_Special *SpecialCallerSession) FutureAdminFee() (*big.Int, error) {
	return _Special.Contract.FutureAdminFee(&_Special.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_Special *SpecialCaller) FutureFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "future_fee")
	return *ret0, err
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_Special *SpecialSession) FutureFee() (*big.Int, error) {
	return _Special.Contract.FutureFee(&_Special.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_Special *SpecialCallerSession) FutureFee() (*big.Int, error) {
	return _Special.Contract.FutureFee(&_Special.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_Special *SpecialCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "future_owner")
	return *ret0, err
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_Special *SpecialSession) FutureOwner() (common.Address, error) {
	return _Special.Contract.FutureOwner(&_Special.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_Special *SpecialCallerSession) FutureOwner() (common.Address, error) {
	return _Special.Contract.FutureOwner(&_Special.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_Special *SpecialCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "get_dy", i, j, dx)
	return *ret0, err
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_Special *SpecialSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Special.Contract.GetDy(&_Special.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_Special *SpecialCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Special.Contract.GetDy(&_Special.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_Special *SpecialCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "get_dy_underlying", i, j, dx)
	return *ret0, err
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_Special *SpecialSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Special.Contract.GetDyUnderlying(&_Special.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_Special *SpecialCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Special.Contract.GetDyUnderlying(&_Special.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Special *SpecialCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "get_virtual_price")
	return *ret0, err
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Special *SpecialSession) GetVirtualPrice() (*big.Int, error) {
	return _Special.Contract.GetVirtualPrice(&_Special.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Special *SpecialCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _Special.Contract.GetVirtualPrice(&_Special.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_Special *SpecialCaller) InitialA(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "initial_A")
	return *ret0, err
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_Special *SpecialSession) InitialA() (*big.Int, error) {
	return _Special.Contract.InitialA(&_Special.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_Special *SpecialCallerSession) InitialA() (*big.Int, error) {
	return _Special.Contract.InitialA(&_Special.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_Special *SpecialCaller) InitialATime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "initial_A_time")
	return *ret0, err
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_Special *SpecialSession) InitialATime() (*big.Int, error) {
	return _Special.Contract.InitialATime(&_Special.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_Special *SpecialCallerSession) InitialATime() (*big.Int, error) {
	return _Special.Contract.InitialATime(&_Special.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Special *SpecialCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Special *SpecialSession) Owner() (common.Address, error) {
	return _Special.Contract.Owner(&_Special.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Special *SpecialCallerSession) Owner() (common.Address, error) {
	return _Special.Contract.Owner(&_Special.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_Special *SpecialCaller) TransferOwnershipDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Special.contract.Call(opts, out, "transfer_ownership_deadline")
	return *ret0, err
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_Special *SpecialSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _Special.Contract.TransferOwnershipDeadline(&_Special.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_Special *SpecialCallerSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _Special.Contract.TransferOwnershipDeadline(&_Special.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) returns()
func (_Special *SpecialTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) returns()
func (_Special *SpecialSession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Special.Contract.AddLiquidity(&_Special.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) returns()
func (_Special *SpecialTransactorSession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Special.Contract.AddLiquidity(&_Special.TransactOpts, amounts, min_mint_amount)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_Special *SpecialTransactor) ApplyNewFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "apply_new_fee")
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_Special *SpecialSession) ApplyNewFee() (*types.Transaction, error) {
	return _Special.Contract.ApplyNewFee(&_Special.TransactOpts)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_Special *SpecialTransactorSession) ApplyNewFee() (*types.Transaction, error) {
	return _Special.Contract.ApplyNewFee(&_Special.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Special *SpecialTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Special *SpecialSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Special.Contract.ApplyTransferOwnership(&_Special.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Special *SpecialTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Special.Contract.ApplyTransferOwnership(&_Special.TransactOpts)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x5b5a1467.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee) returns()
func (_Special *SpecialTransactor) CommitNewFee(opts *bind.TransactOpts, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "commit_new_fee", new_fee, new_admin_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x5b5a1467.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee) returns()
func (_Special *SpecialSession) CommitNewFee(new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Special.Contract.CommitNewFee(&_Special.TransactOpts, new_fee, new_admin_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x5b5a1467.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee) returns()
func (_Special *SpecialTransactorSession) CommitNewFee(new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Special.Contract.CommitNewFee(&_Special.TransactOpts, new_fee, new_admin_fee)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Special *SpecialTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Special *SpecialSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Special.Contract.CommitTransferOwnership(&_Special.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Special *SpecialTransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Special.Contract.CommitTransferOwnership(&_Special.TransactOpts, _owner)
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_Special *SpecialTransactor) DonateAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "donate_admin_fees")
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_Special *SpecialSession) DonateAdminFees() (*types.Transaction, error) {
	return _Special.Contract.DonateAdminFees(&_Special.TransactOpts)
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_Special *SpecialTransactorSession) DonateAdminFees() (*types.Transaction, error) {
	return _Special.Contract.DonateAdminFees(&_Special.TransactOpts)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Special *SpecialTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Special *SpecialSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Special.Contract.Exchange(&_Special.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Special *SpecialTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Special.Contract.Exchange(&_Special.TransactOpts, i, j, dx, min_dy)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Special *SpecialTransactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Special *SpecialSession) KillMe() (*types.Transaction, error) {
	return _Special.Contract.KillMe(&_Special.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Special *SpecialTransactorSession) KillMe() (*types.Transaction, error) {
	return _Special.Contract.KillMe(&_Special.TransactOpts)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_Special *SpecialTransactor) RampA(opts *bind.TransactOpts, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "ramp_A", _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_Special *SpecialSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _Special.Contract.RampA(&_Special.TransactOpts, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_Special *SpecialTransactorSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _Special.Contract.RampA(&_Special.TransactOpts, _future_A, _future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Special *SpecialTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Special *SpecialSession) RemoveLiquidity(_amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Special.Contract.RemoveLiquidity(&_Special.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Special *SpecialTransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Special.Contract.RemoveLiquidity(&_Special.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 max_burn_amount) returns()
func (_Special *SpecialTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, amounts [2]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "remove_liquidity_imbalance", amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 max_burn_amount) returns()
func (_Special *SpecialSession) RemoveLiquidityImbalance(amounts [2]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Special.Contract.RemoveLiquidityImbalance(&_Special.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 max_burn_amount) returns()
func (_Special *SpecialTransactorSession) RemoveLiquidityImbalance(amounts [2]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Special.Contract.RemoveLiquidityImbalance(&_Special.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_Special *SpecialTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "remove_liquidity_one_coin", _token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_Special *SpecialSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Special.Contract.RemoveLiquidityOneCoin(&_Special.TransactOpts, _token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_Special *SpecialTransactorSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Special.Contract.RemoveLiquidityOneCoin(&_Special.TransactOpts, _token_amount, i, min_amount)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Special *SpecialTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Special *SpecialSession) RevertNewParameters() (*types.Transaction, error) {
	return _Special.Contract.RevertNewParameters(&_Special.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Special *SpecialTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _Special.Contract.RevertNewParameters(&_Special.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Special *SpecialTransactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Special *SpecialSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Special.Contract.RevertTransferOwnership(&_Special.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Special *SpecialTransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Special.Contract.RevertTransferOwnership(&_Special.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_Special *SpecialTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "stop_ramp_A")
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_Special *SpecialSession) StopRampA() (*types.Transaction, error) {
	return _Special.Contract.StopRampA(&_Special.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_Special *SpecialTransactorSession) StopRampA() (*types.Transaction, error) {
	return _Special.Contract.StopRampA(&_Special.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Special *SpecialTransactor) UnkillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "unkill_me")
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Special *SpecialSession) UnkillMe() (*types.Transaction, error) {
	return _Special.Contract.UnkillMe(&_Special.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Special *SpecialTransactorSession) UnkillMe() (*types.Transaction, error) {
	return _Special.Contract.UnkillMe(&_Special.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Special *SpecialTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Special.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Special *SpecialSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Special.Contract.WithdrawAdminFees(&_Special.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Special *SpecialTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Special.Contract.WithdrawAdminFees(&_Special.TransactOpts)
}

// SpecialAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Special contract.
type SpecialAddLiquidityIterator struct {
	Event *SpecialAddLiquidity // Event containing the contract specifics and raw log

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
func (it *SpecialAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpecialAddLiquidity)
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
		it.Event = new(SpecialAddLiquidity)
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
func (it *SpecialAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpecialAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpecialAddLiquidity represents a AddLiquidity event raised by the Special contract.
type SpecialAddLiquidity struct {
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
func (_Special *SpecialFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*SpecialAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Special.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &SpecialAddLiquidityIterator{contract: _Special.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x26f55a85081d24974e85c6c00045d0f0453991e95873f52bff0d21af4079a768.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_Special *SpecialFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *SpecialAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Special.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpecialAddLiquidity)
				if err := _Special.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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
func (_Special *SpecialFilterer) ParseAddLiquidity(log types.Log) (*SpecialAddLiquidity, error) {
	event := new(SpecialAddLiquidity)
	if err := _Special.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SpecialCommitNewAdminIterator is returned from FilterCommitNewAdmin and is used to iterate over the raw logs and unpacked data for CommitNewAdmin events raised by the Special contract.
type SpecialCommitNewAdminIterator struct {
	Event *SpecialCommitNewAdmin // Event containing the contract specifics and raw log

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
func (it *SpecialCommitNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpecialCommitNewAdmin)
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
		it.Event = new(SpecialCommitNewAdmin)
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
func (it *SpecialCommitNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpecialCommitNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpecialCommitNewAdmin represents a CommitNewAdmin event raised by the Special contract.
type SpecialCommitNewAdmin struct {
	Deadline *big.Int
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewAdmin is a free log retrieval operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Special *SpecialFilterer) FilterCommitNewAdmin(opts *bind.FilterOpts, deadline []*big.Int, admin []common.Address) (*SpecialCommitNewAdminIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Special.contract.FilterLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &SpecialCommitNewAdminIterator{contract: _Special.contract, event: "CommitNewAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitNewAdmin is a free log subscription operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Special *SpecialFilterer) WatchCommitNewAdmin(opts *bind.WatchOpts, sink chan<- *SpecialCommitNewAdmin, deadline []*big.Int, admin []common.Address) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Special.contract.WatchLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpecialCommitNewAdmin)
				if err := _Special.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
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
func (_Special *SpecialFilterer) ParseCommitNewAdmin(log types.Log) (*SpecialCommitNewAdmin, error) {
	event := new(SpecialCommitNewAdmin)
	if err := _Special.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SpecialCommitNewFeeIterator is returned from FilterCommitNewFee and is used to iterate over the raw logs and unpacked data for CommitNewFee events raised by the Special contract.
type SpecialCommitNewFeeIterator struct {
	Event *SpecialCommitNewFee // Event containing the contract specifics and raw log

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
func (it *SpecialCommitNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpecialCommitNewFee)
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
		it.Event = new(SpecialCommitNewFee)
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
func (it *SpecialCommitNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpecialCommitNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpecialCommitNewFee represents a CommitNewFee event raised by the Special contract.
type SpecialCommitNewFee struct {
	Deadline *big.Int
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewFee is a free log retrieval operation binding the contract event 0x351fc5da2fbf480f2225debf3664a4bc90fa9923743aad58b4603f648e931fe0.
//
// Solidity: event CommitNewFee(uint256 indexed deadline, uint256 fee, uint256 admin_fee)
func (_Special *SpecialFilterer) FilterCommitNewFee(opts *bind.FilterOpts, deadline []*big.Int) (*SpecialCommitNewFeeIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Special.contract.FilterLogs(opts, "CommitNewFee", deadlineRule)
	if err != nil {
		return nil, err
	}
	return &SpecialCommitNewFeeIterator{contract: _Special.contract, event: "CommitNewFee", logs: logs, sub: sub}, nil
}

// WatchCommitNewFee is a free log subscription operation binding the contract event 0x351fc5da2fbf480f2225debf3664a4bc90fa9923743aad58b4603f648e931fe0.
//
// Solidity: event CommitNewFee(uint256 indexed deadline, uint256 fee, uint256 admin_fee)
func (_Special *SpecialFilterer) WatchCommitNewFee(opts *bind.WatchOpts, sink chan<- *SpecialCommitNewFee, deadline []*big.Int) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Special.contract.WatchLogs(opts, "CommitNewFee", deadlineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpecialCommitNewFee)
				if err := _Special.contract.UnpackLog(event, "CommitNewFee", log); err != nil {
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
func (_Special *SpecialFilterer) ParseCommitNewFee(log types.Log) (*SpecialCommitNewFee, error) {
	event := new(SpecialCommitNewFee)
	if err := _Special.contract.UnpackLog(event, "CommitNewFee", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SpecialNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Special contract.
type SpecialNewAdminIterator struct {
	Event *SpecialNewAdmin // Event containing the contract specifics and raw log

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
func (it *SpecialNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpecialNewAdmin)
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
		it.Event = new(SpecialNewAdmin)
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
func (it *SpecialNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpecialNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpecialNewAdmin represents a NewAdmin event raised by the Special contract.
type SpecialNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Special *SpecialFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address) (*SpecialNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Special.contract.FilterLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &SpecialNewAdminIterator{contract: _Special.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Special *SpecialFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *SpecialNewAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Special.contract.WatchLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpecialNewAdmin)
				if err := _Special.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_Special *SpecialFilterer) ParseNewAdmin(log types.Log) (*SpecialNewAdmin, error) {
	event := new(SpecialNewAdmin)
	if err := _Special.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SpecialNewFeeIterator is returned from FilterNewFee and is used to iterate over the raw logs and unpacked data for NewFee events raised by the Special contract.
type SpecialNewFeeIterator struct {
	Event *SpecialNewFee // Event containing the contract specifics and raw log

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
func (it *SpecialNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpecialNewFee)
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
		it.Event = new(SpecialNewFee)
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
func (it *SpecialNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpecialNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpecialNewFee represents a NewFee event raised by the Special contract.
type SpecialNewFee struct {
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewFee is a free log retrieval operation binding the contract event 0xbe12859b636aed607d5230b2cc2711f68d70e51060e6cca1f575ef5d2fcc95d1.
//
// Solidity: event NewFee(uint256 fee, uint256 admin_fee)
func (_Special *SpecialFilterer) FilterNewFee(opts *bind.FilterOpts) (*SpecialNewFeeIterator, error) {

	logs, sub, err := _Special.contract.FilterLogs(opts, "NewFee")
	if err != nil {
		return nil, err
	}
	return &SpecialNewFeeIterator{contract: _Special.contract, event: "NewFee", logs: logs, sub: sub}, nil
}

// WatchNewFee is a free log subscription operation binding the contract event 0xbe12859b636aed607d5230b2cc2711f68d70e51060e6cca1f575ef5d2fcc95d1.
//
// Solidity: event NewFee(uint256 fee, uint256 admin_fee)
func (_Special *SpecialFilterer) WatchNewFee(opts *bind.WatchOpts, sink chan<- *SpecialNewFee) (event.Subscription, error) {

	logs, sub, err := _Special.contract.WatchLogs(opts, "NewFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpecialNewFee)
				if err := _Special.contract.UnpackLog(event, "NewFee", log); err != nil {
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
func (_Special *SpecialFilterer) ParseNewFee(log types.Log) (*SpecialNewFee, error) {
	event := new(SpecialNewFee)
	if err := _Special.contract.UnpackLog(event, "NewFee", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SpecialRampAIterator is returned from FilterRampA and is used to iterate over the raw logs and unpacked data for RampA events raised by the Special contract.
type SpecialRampAIterator struct {
	Event *SpecialRampA // Event containing the contract specifics and raw log

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
func (it *SpecialRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpecialRampA)
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
		it.Event = new(SpecialRampA)
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
func (it *SpecialRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpecialRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpecialRampA represents a RampA event raised by the Special contract.
type SpecialRampA struct {
	OldA        *big.Int
	NewA        *big.Int
	InitialTime *big.Int
	FutureTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRampA is a free log retrieval operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_Special *SpecialFilterer) FilterRampA(opts *bind.FilterOpts) (*SpecialRampAIterator, error) {

	logs, sub, err := _Special.contract.FilterLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return &SpecialRampAIterator{contract: _Special.contract, event: "RampA", logs: logs, sub: sub}, nil
}

// WatchRampA is a free log subscription operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_Special *SpecialFilterer) WatchRampA(opts *bind.WatchOpts, sink chan<- *SpecialRampA) (event.Subscription, error) {

	logs, sub, err := _Special.contract.WatchLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpecialRampA)
				if err := _Special.contract.UnpackLog(event, "RampA", log); err != nil {
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
func (_Special *SpecialFilterer) ParseRampA(log types.Log) (*SpecialRampA, error) {
	event := new(SpecialRampA)
	if err := _Special.contract.UnpackLog(event, "RampA", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SpecialRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Special contract.
type SpecialRemoveLiquidityIterator struct {
	Event *SpecialRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *SpecialRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpecialRemoveLiquidity)
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
		it.Event = new(SpecialRemoveLiquidity)
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
func (it *SpecialRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpecialRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpecialRemoveLiquidity represents a RemoveLiquidity event raised by the Special contract.
type SpecialRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fees         [2]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_Special *SpecialFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*SpecialRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Special.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &SpecialRemoveLiquidityIterator{contract: _Special.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_Special *SpecialFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *SpecialRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Special.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpecialRemoveLiquidity)
				if err := _Special.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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
func (_Special *SpecialFilterer) ParseRemoveLiquidity(log types.Log) (*SpecialRemoveLiquidity, error) {
	event := new(SpecialRemoveLiquidity)
	if err := _Special.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SpecialRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the Special contract.
type SpecialRemoveLiquidityImbalanceIterator struct {
	Event *SpecialRemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *SpecialRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpecialRemoveLiquidityImbalance)
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
		it.Event = new(SpecialRemoveLiquidityImbalance)
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
func (it *SpecialRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpecialRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpecialRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the Special contract.
type SpecialRemoveLiquidityImbalance struct {
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
func (_Special *SpecialFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*SpecialRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Special.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &SpecialRemoveLiquidityImbalanceIterator{contract: _Special.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x2b5508378d7e19e0d5fa338419034731416c4f5b219a10379956f764317fd47e.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_Special *SpecialFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *SpecialRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Special.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpecialRemoveLiquidityImbalance)
				if err := _Special.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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
func (_Special *SpecialFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*SpecialRemoveLiquidityImbalance, error) {
	event := new(SpecialRemoveLiquidityImbalance)
	if err := _Special.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SpecialRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the Special contract.
type SpecialRemoveLiquidityOneIterator struct {
	Event *SpecialRemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *SpecialRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpecialRemoveLiquidityOne)
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
		it.Event = new(SpecialRemoveLiquidityOne)
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
func (it *SpecialRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpecialRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpecialRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the Special contract.
type SpecialRemoveLiquidityOne struct {
	Provider    common.Address
	TokenAmount *big.Int
	CoinAmount  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x9e96dd3b997a2a257eec4df9bb6eaf626e206df5f543bd963682d143300be310.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount)
func (_Special *SpecialFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*SpecialRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Special.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &SpecialRemoveLiquidityOneIterator{contract: _Special.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x9e96dd3b997a2a257eec4df9bb6eaf626e206df5f543bd963682d143300be310.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount)
func (_Special *SpecialFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *SpecialRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Special.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpecialRemoveLiquidityOne)
				if err := _Special.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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
func (_Special *SpecialFilterer) ParseRemoveLiquidityOne(log types.Log) (*SpecialRemoveLiquidityOne, error) {
	event := new(SpecialRemoveLiquidityOne)
	if err := _Special.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SpecialStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the Special contract.
type SpecialStopRampAIterator struct {
	Event *SpecialStopRampA // Event containing the contract specifics and raw log

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
func (it *SpecialStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpecialStopRampA)
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
		it.Event = new(SpecialStopRampA)
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
func (it *SpecialStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpecialStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpecialStopRampA represents a StopRampA event raised by the Special contract.
type SpecialStopRampA struct {
	A   *big.Int
	T   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_Special *SpecialFilterer) FilterStopRampA(opts *bind.FilterOpts) (*SpecialStopRampAIterator, error) {

	logs, sub, err := _Special.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &SpecialStopRampAIterator{contract: _Special.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_Special *SpecialFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *SpecialStopRampA) (event.Subscription, error) {

	logs, sub, err := _Special.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpecialStopRampA)
				if err := _Special.contract.UnpackLog(event, "StopRampA", log); err != nil {
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
func (_Special *SpecialFilterer) ParseStopRampA(log types.Log) (*SpecialStopRampA, error) {
	event := new(SpecialStopRampA)
	if err := _Special.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SpecialTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the Special contract.
type SpecialTokenExchangeIterator struct {
	Event *SpecialTokenExchange // Event containing the contract specifics and raw log

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
func (it *SpecialTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpecialTokenExchange)
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
		it.Event = new(SpecialTokenExchange)
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
func (it *SpecialTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpecialTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpecialTokenExchange represents a TokenExchange event raised by the Special contract.
type SpecialTokenExchange struct {
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
func (_Special *SpecialFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*SpecialTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Special.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &SpecialTokenExchangeIterator{contract: _Special.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Special *SpecialFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *SpecialTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Special.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpecialTokenExchange)
				if err := _Special.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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
func (_Special *SpecialFilterer) ParseTokenExchange(log types.Log) (*SpecialTokenExchange, error) {
	event := new(SpecialTokenExchange)
	if err := _Special.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	return event, nil
}
