// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Strategy

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

// StrategyABI is the input ABI used to generate the binding from.
const StrategyABI = "[{\"name\":\"TokenExchange\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"sold_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"bought_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchangeUnderlying\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"sold_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"bought_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[2]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[2]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[2]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[2]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[2]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[2]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true,\"unit\":\"sec\"},{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewParameters\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true,\"unit\":\"sec\"},{\"type\":\"uint256\",\"name\":\"A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewParameters\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address[2]\",\"name\":\"_coins\"},{\"type\":\"address[2]\",\"name\":\"_underlying_coins\"},{\"type\":\"address\",\"name\":\"_pool_token\"},{\"type\":\"uint256\",\"name\":\"_A\"},{\"type\":\"uint256\",\"name\":\"_fee\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"get_virtual_price\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1084167},{\"name\":\"calc_token_amount\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256[2]\",\"name\":\"amounts\"},{\"type\":\"bool\",\"name\":\"deposit\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":4239939},{\"name\":\"add_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[2]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"min_mint_amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":6479997},{\"name\":\"get_dy\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2543681},{\"name\":\"get_dx\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dy\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2543687},{\"name\":\"get_dy_underlying\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2543506},{\"name\":\"get_dx_underlying\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dy\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2543512},{\"name\":\"exchange\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"},{\"type\":\"uint256\",\"name\":\"min_dy\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":5184573},{\"name\":\"exchange_underlying\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"},{\"type\":\"uint256\",\"name\":\"min_dy\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":5200817},{\"name\":\"remove_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"},{\"type\":\"uint256[2]\",\"name\":\"min_amounts\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":153898},{\"name\":\"remove_liquidity_imbalance\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[2]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"max_burn_amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":6479708},{\"name\":\"commit_new_parameters\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"amplification\"},{\"type\":\"uint256\",\"name\":\"new_fee\"},{\"type\":\"uint256\",\"name\":\"new_admin_fee\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":146105},{\"name\":\"apply_new_parameters\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":133512},{\"name\":\"revert_new_parameters\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":21835},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":74512},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":60568},{\"name\":\"revert_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":21925},{\"name\":\"withdraw_admin_fees\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":12831},{\"name\":\"kill_me\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":37878},{\"name\":\"unkill_me\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":22015},{\"name\":\"coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2190},{\"name\":\"underlying_coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2220},{\"name\":\"balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2250},{\"name\":\"A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2081},{\"name\":\"fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2111},{\"name\":\"admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2141},{\"name\":\"owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2171},{\"name\":\"admin_actions_deadline\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"sec\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2201},{\"name\":\"transfer_ownership_deadline\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"sec\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2231},{\"name\":\"future_A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2261},{\"name\":\"future_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2291},{\"name\":\"future_admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2321},{\"name\":\"future_owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2351}]"

// Strategy is an auto generated Go binding around an Ethereum contract.
type Strategy struct {
	StrategyCaller     // Read-only binding to the contract
	StrategyTransactor // Write-only binding to the contract
	StrategyFilterer   // Log filterer for contract events
}

// StrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategySession struct {
	Contract     *Strategy         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyCallerSession struct {
	Contract *StrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyTransactorSession struct {
	Contract     *StrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyRaw struct {
	Contract *Strategy // Generic contract binding to access the raw methods on
}

// StrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyCallerRaw struct {
	Contract *StrategyCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyTransactorRaw struct {
	Contract *StrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategy creates a new instance of Strategy, bound to a specific deployed contract.
func NewStrategy(address common.Address, backend bind.ContractBackend) (*Strategy, error) {
	contract, err := bindStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strategy{StrategyCaller: StrategyCaller{contract: contract}, StrategyTransactor: StrategyTransactor{contract: contract}, StrategyFilterer: StrategyFilterer{contract: contract}}, nil
}

// NewStrategyCaller creates a new read-only instance of Strategy, bound to a specific deployed contract.
func NewStrategyCaller(address common.Address, caller bind.ContractCaller) (*StrategyCaller, error) {
	contract, err := bindStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyCaller{contract: contract}, nil
}

// NewStrategyTransactor creates a new write-only instance of Strategy, bound to a specific deployed contract.
func NewStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyTransactor, error) {
	contract, err := bindStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyTransactor{contract: contract}, nil
}

// NewStrategyFilterer creates a new log filterer instance of Strategy, bound to a specific deployed contract.
func NewStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyFilterer, error) {
	contract, err := bindStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyFilterer{contract: contract}, nil
}

// bindStrategy binds a generic wrapper to an already deployed contract.
func bindStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StrategyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strategy *StrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strategy.Contract.StrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strategy *StrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strategy.Contract.StrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strategy *StrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strategy.Contract.StrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strategy *StrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strategy *StrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strategy *StrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strategy.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() returns(uint256 out)
func (_Strategy *StrategyCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() returns(uint256 out)
func (_Strategy *StrategySession) A() (*big.Int, error) {
	return _Strategy.Contract.A(&_Strategy.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() returns(uint256 out)
func (_Strategy *StrategyCallerSession) A() (*big.Int, error) {
	return _Strategy.Contract.A(&_Strategy.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() returns(uint256 out)
func (_Strategy *StrategyCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "admin_actions_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() returns(uint256 out)
func (_Strategy *StrategySession) AdminActionsDeadline() (*big.Int, error) {
	return _Strategy.Contract.AdminActionsDeadline(&_Strategy.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() returns(uint256 out)
func (_Strategy *StrategyCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _Strategy.Contract.AdminActionsDeadline(&_Strategy.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() returns(uint256 out)
func (_Strategy *StrategyCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() returns(uint256 out)
func (_Strategy *StrategySession) AdminFee() (*big.Int, error) {
	return _Strategy.Contract.AdminFee(&_Strategy.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() returns(uint256 out)
func (_Strategy *StrategyCallerSession) AdminFee() (*big.Int, error) {
	return _Strategy.Contract.AdminFee(&_Strategy.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 arg0) returns(uint256 out)
func (_Strategy *StrategyCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 arg0) returns(uint256 out)
func (_Strategy *StrategySession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Strategy.Contract.Balances(&_Strategy.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 arg0) returns(uint256 out)
func (_Strategy *StrategyCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Strategy.Contract.Balances(&_Strategy.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] amounts, bool deposit) returns(uint256 out)
func (_Strategy *StrategyCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [2]*big.Int, deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "calc_token_amount", amounts, deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] amounts, bool deposit) returns(uint256 out)
func (_Strategy *StrategySession) CalcTokenAmount(amounts [2]*big.Int, deposit bool) (*big.Int, error) {
	return _Strategy.Contract.CalcTokenAmount(&_Strategy.CallOpts, amounts, deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] amounts, bool deposit) returns(uint256 out)
func (_Strategy *StrategyCallerSession) CalcTokenAmount(amounts [2]*big.Int, deposit bool) (*big.Int, error) {
	return _Strategy.Contract.CalcTokenAmount(&_Strategy.CallOpts, amounts, deposit)
}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 arg0) returns(address out)
func (_Strategy *StrategyCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 arg0) returns(address out)
func (_Strategy *StrategySession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Strategy.Contract.Coins(&_Strategy.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 arg0) returns(address out)
func (_Strategy *StrategyCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Strategy.Contract.Coins(&_Strategy.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() returns(uint256 out)
func (_Strategy *StrategyCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() returns(uint256 out)
func (_Strategy *StrategySession) Fee() (*big.Int, error) {
	return _Strategy.Contract.Fee(&_Strategy.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() returns(uint256 out)
func (_Strategy *StrategyCallerSession) Fee() (*big.Int, error) {
	return _Strategy.Contract.Fee(&_Strategy.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() returns(uint256 out)
func (_Strategy *StrategyCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() returns(uint256 out)
func (_Strategy *StrategySession) FutureA() (*big.Int, error) {
	return _Strategy.Contract.FutureA(&_Strategy.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() returns(uint256 out)
func (_Strategy *StrategyCallerSession) FutureA() (*big.Int, error) {
	return _Strategy.Contract.FutureA(&_Strategy.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() returns(uint256 out)
func (_Strategy *StrategyCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "future_admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() returns(uint256 out)
func (_Strategy *StrategySession) FutureAdminFee() (*big.Int, error) {
	return _Strategy.Contract.FutureAdminFee(&_Strategy.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() returns(uint256 out)
func (_Strategy *StrategyCallerSession) FutureAdminFee() (*big.Int, error) {
	return _Strategy.Contract.FutureAdminFee(&_Strategy.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() returns(uint256 out)
func (_Strategy *StrategyCaller) FutureFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "future_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() returns(uint256 out)
func (_Strategy *StrategySession) FutureFee() (*big.Int, error) {
	return _Strategy.Contract.FutureFee(&_Strategy.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() returns(uint256 out)
func (_Strategy *StrategyCallerSession) FutureFee() (*big.Int, error) {
	return _Strategy.Contract.FutureFee(&_Strategy.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() returns(address out)
func (_Strategy *StrategyCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "future_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() returns(address out)
func (_Strategy *StrategySession) FutureOwner() (common.Address, error) {
	return _Strategy.Contract.FutureOwner(&_Strategy.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() returns(address out)
func (_Strategy *StrategyCallerSession) FutureOwner() (common.Address, error) {
	return _Strategy.Contract.FutureOwner(&_Strategy.CallOpts)
}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) returns(uint256 out)
func (_Strategy *StrategyCaller) GetDx(opts *bind.CallOpts, i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "get_dx", i, j, dy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) returns(uint256 out)
func (_Strategy *StrategySession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _Strategy.Contract.GetDx(&_Strategy.CallOpts, i, j, dy)
}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) returns(uint256 out)
func (_Strategy *StrategyCallerSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _Strategy.Contract.GetDx(&_Strategy.CallOpts, i, j, dy)
}

// GetDxUnderlying is a free data retrieval call binding the contract method 0x0e71d1b9.
//
// Solidity: function get_dx_underlying(int128 i, int128 j, uint256 dy) returns(uint256 out)
func (_Strategy *StrategyCaller) GetDxUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "get_dx_underlying", i, j, dy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDxUnderlying is a free data retrieval call binding the contract method 0x0e71d1b9.
//
// Solidity: function get_dx_underlying(int128 i, int128 j, uint256 dy) returns(uint256 out)
func (_Strategy *StrategySession) GetDxUnderlying(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _Strategy.Contract.GetDxUnderlying(&_Strategy.CallOpts, i, j, dy)
}

// GetDxUnderlying is a free data retrieval call binding the contract method 0x0e71d1b9.
//
// Solidity: function get_dx_underlying(int128 i, int128 j, uint256 dy) returns(uint256 out)
func (_Strategy *StrategyCallerSession) GetDxUnderlying(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _Strategy.Contract.GetDxUnderlying(&_Strategy.CallOpts, i, j, dy)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) returns(uint256 out)
func (_Strategy *StrategyCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) returns(uint256 out)
func (_Strategy *StrategySession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Strategy.Contract.GetDy(&_Strategy.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) returns(uint256 out)
func (_Strategy *StrategyCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Strategy.Contract.GetDy(&_Strategy.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) returns(uint256 out)
func (_Strategy *StrategyCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "get_dy_underlying", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) returns(uint256 out)
func (_Strategy *StrategySession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Strategy.Contract.GetDyUnderlying(&_Strategy.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) returns(uint256 out)
func (_Strategy *StrategyCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Strategy.Contract.GetDyUnderlying(&_Strategy.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() returns(uint256 out)
func (_Strategy *StrategyCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() returns(uint256 out)
func (_Strategy *StrategySession) GetVirtualPrice() (*big.Int, error) {
	return _Strategy.Contract.GetVirtualPrice(&_Strategy.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() returns(uint256 out)
func (_Strategy *StrategyCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _Strategy.Contract.GetVirtualPrice(&_Strategy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address out)
func (_Strategy *StrategyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address out)
func (_Strategy *StrategySession) Owner() (common.Address, error) {
	return _Strategy.Contract.Owner(&_Strategy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address out)
func (_Strategy *StrategyCallerSession) Owner() (common.Address, error) {
	return _Strategy.Contract.Owner(&_Strategy.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() returns(uint256 out)
func (_Strategy *StrategyCaller) TransferOwnershipDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "transfer_ownership_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() returns(uint256 out)
func (_Strategy *StrategySession) TransferOwnershipDeadline() (*big.Int, error) {
	return _Strategy.Contract.TransferOwnershipDeadline(&_Strategy.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() returns(uint256 out)
func (_Strategy *StrategyCallerSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _Strategy.Contract.TransferOwnershipDeadline(&_Strategy.CallOpts)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 arg0) returns(address out)
func (_Strategy *StrategyCaller) UnderlyingCoins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "underlying_coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 arg0) returns(address out)
func (_Strategy *StrategySession) UnderlyingCoins(arg0 *big.Int) (common.Address, error) {
	return _Strategy.Contract.UnderlyingCoins(&_Strategy.CallOpts, arg0)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 arg0) returns(address out)
func (_Strategy *StrategyCallerSession) UnderlyingCoins(arg0 *big.Int) (common.Address, error) {
	return _Strategy.Contract.UnderlyingCoins(&_Strategy.CallOpts, arg0)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) returns()
func (_Strategy *StrategyTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) returns()
func (_Strategy *StrategySession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.AddLiquidity(&_Strategy.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) returns()
func (_Strategy *StrategyTransactorSession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.AddLiquidity(&_Strategy.TransactOpts, amounts, min_mint_amount)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Strategy *StrategyTransactor) ApplyNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "apply_new_parameters")
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Strategy *StrategySession) ApplyNewParameters() (*types.Transaction, error) {
	return _Strategy.Contract.ApplyNewParameters(&_Strategy.TransactOpts)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Strategy *StrategyTransactorSession) ApplyNewParameters() (*types.Transaction, error) {
	return _Strategy.Contract.ApplyNewParameters(&_Strategy.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Strategy *StrategyTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Strategy *StrategySession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Strategy.Contract.ApplyTransferOwnership(&_Strategy.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Strategy *StrategyTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Strategy.Contract.ApplyTransferOwnership(&_Strategy.TransactOpts)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xee11f5b6.
//
// Solidity: function commit_new_parameters(uint256 amplification, uint256 new_fee, uint256 new_admin_fee) returns()
func (_Strategy *StrategyTransactor) CommitNewParameters(opts *bind.TransactOpts, amplification *big.Int, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "commit_new_parameters", amplification, new_fee, new_admin_fee)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xee11f5b6.
//
// Solidity: function commit_new_parameters(uint256 amplification, uint256 new_fee, uint256 new_admin_fee) returns()
func (_Strategy *StrategySession) CommitNewParameters(amplification *big.Int, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.CommitNewParameters(&_Strategy.TransactOpts, amplification, new_fee, new_admin_fee)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xee11f5b6.
//
// Solidity: function commit_new_parameters(uint256 amplification, uint256 new_fee, uint256 new_admin_fee) returns()
func (_Strategy *StrategyTransactorSession) CommitNewParameters(amplification *big.Int, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.CommitNewParameters(&_Strategy.TransactOpts, amplification, new_fee, new_admin_fee)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Strategy *StrategyTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Strategy *StrategySession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Strategy.Contract.CommitTransferOwnership(&_Strategy.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Strategy *StrategyTransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Strategy.Contract.CommitTransferOwnership(&_Strategy.TransactOpts, _owner)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Strategy *StrategyTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Strategy *StrategySession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.Exchange(&_Strategy.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Strategy *StrategyTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.Exchange(&_Strategy.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Strategy *StrategyTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "exchange_underlying", i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Strategy *StrategySession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.ExchangeUnderlying(&_Strategy.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_Strategy *StrategyTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.ExchangeUnderlying(&_Strategy.TransactOpts, i, j, dx, min_dy)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Strategy *StrategyTransactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Strategy *StrategySession) KillMe() (*types.Transaction, error) {
	return _Strategy.Contract.KillMe(&_Strategy.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Strategy *StrategyTransactorSession) KillMe() (*types.Transaction, error) {
	return _Strategy.Contract.KillMe(&_Strategy.TransactOpts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Strategy *StrategyTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Strategy *StrategySession) RemoveLiquidity(_amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.RemoveLiquidity(&_Strategy.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Strategy *StrategyTransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.RemoveLiquidity(&_Strategy.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 max_burn_amount) returns()
func (_Strategy *StrategyTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, amounts [2]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "remove_liquidity_imbalance", amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 max_burn_amount) returns()
func (_Strategy *StrategySession) RemoveLiquidityImbalance(amounts [2]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.RemoveLiquidityImbalance(&_Strategy.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 max_burn_amount) returns()
func (_Strategy *StrategyTransactorSession) RemoveLiquidityImbalance(amounts [2]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.RemoveLiquidityImbalance(&_Strategy.TransactOpts, amounts, max_burn_amount)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Strategy *StrategyTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Strategy *StrategySession) RevertNewParameters() (*types.Transaction, error) {
	return _Strategy.Contract.RevertNewParameters(&_Strategy.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Strategy *StrategyTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _Strategy.Contract.RevertNewParameters(&_Strategy.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Strategy *StrategyTransactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Strategy *StrategySession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Strategy.Contract.RevertTransferOwnership(&_Strategy.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Strategy *StrategyTransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Strategy.Contract.RevertTransferOwnership(&_Strategy.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Strategy *StrategyTransactor) UnkillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "unkill_me")
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Strategy *StrategySession) UnkillMe() (*types.Transaction, error) {
	return _Strategy.Contract.UnkillMe(&_Strategy.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Strategy *StrategyTransactorSession) UnkillMe() (*types.Transaction, error) {
	return _Strategy.Contract.UnkillMe(&_Strategy.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Strategy *StrategyTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Strategy *StrategySession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Strategy.Contract.WithdrawAdminFees(&_Strategy.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Strategy *StrategyTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Strategy.Contract.WithdrawAdminFees(&_Strategy.TransactOpts)
}

// StrategyAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Strategy contract.
type StrategyAddLiquidityIterator struct {
	Event *StrategyAddLiquidity // Event containing the contract specifics and raw log

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
func (it *StrategyAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyAddLiquidity)
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
		it.Event = new(StrategyAddLiquidity)
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
func (it *StrategyAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyAddLiquidity represents a AddLiquidity event raised by the Strategy contract.
type StrategyAddLiquidity struct {
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
func (_Strategy *StrategyFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*StrategyAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &StrategyAddLiquidityIterator{contract: _Strategy.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x26f55a85081d24974e85c6c00045d0f0453991e95873f52bff0d21af4079a768.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_Strategy *StrategyFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *StrategyAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyAddLiquidity)
				if err := _Strategy.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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
func (_Strategy *StrategyFilterer) ParseAddLiquidity(log types.Log) (*StrategyAddLiquidity, error) {
	event := new(StrategyAddLiquidity)
	if err := _Strategy.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StrategyCommitNewAdminIterator is returned from FilterCommitNewAdmin and is used to iterate over the raw logs and unpacked data for CommitNewAdmin events raised by the Strategy contract.
type StrategyCommitNewAdminIterator struct {
	Event *StrategyCommitNewAdmin // Event containing the contract specifics and raw log

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
func (it *StrategyCommitNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyCommitNewAdmin)
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
		it.Event = new(StrategyCommitNewAdmin)
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
func (it *StrategyCommitNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyCommitNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyCommitNewAdmin represents a CommitNewAdmin event raised by the Strategy contract.
type StrategyCommitNewAdmin struct {
	Deadline *big.Int
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewAdmin is a free log retrieval operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Strategy *StrategyFilterer) FilterCommitNewAdmin(opts *bind.FilterOpts, deadline []*big.Int, admin []common.Address) (*StrategyCommitNewAdminIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &StrategyCommitNewAdminIterator{contract: _Strategy.contract, event: "CommitNewAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitNewAdmin is a free log subscription operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Strategy *StrategyFilterer) WatchCommitNewAdmin(opts *bind.WatchOpts, sink chan<- *StrategyCommitNewAdmin, deadline []*big.Int, admin []common.Address) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyCommitNewAdmin)
				if err := _Strategy.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
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
func (_Strategy *StrategyFilterer) ParseCommitNewAdmin(log types.Log) (*StrategyCommitNewAdmin, error) {
	event := new(StrategyCommitNewAdmin)
	if err := _Strategy.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StrategyCommitNewParametersIterator is returned from FilterCommitNewParameters and is used to iterate over the raw logs and unpacked data for CommitNewParameters events raised by the Strategy contract.
type StrategyCommitNewParametersIterator struct {
	Event *StrategyCommitNewParameters // Event containing the contract specifics and raw log

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
func (it *StrategyCommitNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyCommitNewParameters)
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
		it.Event = new(StrategyCommitNewParameters)
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
func (it *StrategyCommitNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyCommitNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyCommitNewParameters represents a CommitNewParameters event raised by the Strategy contract.
type StrategyCommitNewParameters struct {
	Deadline *big.Int
	A        *big.Int
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewParameters is a free log retrieval operation binding the contract event 0x6081daa3b61098baf24d9c69bcd53af932e0635c89c6fd0617534b9ba76a7f73.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 A, uint256 fee, uint256 admin_fee)
func (_Strategy *StrategyFilterer) FilterCommitNewParameters(opts *bind.FilterOpts, deadline []*big.Int) (*StrategyCommitNewParametersIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return &StrategyCommitNewParametersIterator{contract: _Strategy.contract, event: "CommitNewParameters", logs: logs, sub: sub}, nil
}

// WatchCommitNewParameters is a free log subscription operation binding the contract event 0x6081daa3b61098baf24d9c69bcd53af932e0635c89c6fd0617534b9ba76a7f73.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 A, uint256 fee, uint256 admin_fee)
func (_Strategy *StrategyFilterer) WatchCommitNewParameters(opts *bind.WatchOpts, sink chan<- *StrategyCommitNewParameters, deadline []*big.Int) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyCommitNewParameters)
				if err := _Strategy.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
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
func (_Strategy *StrategyFilterer) ParseCommitNewParameters(log types.Log) (*StrategyCommitNewParameters, error) {
	event := new(StrategyCommitNewParameters)
	if err := _Strategy.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StrategyNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Strategy contract.
type StrategyNewAdminIterator struct {
	Event *StrategyNewAdmin // Event containing the contract specifics and raw log

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
func (it *StrategyNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyNewAdmin)
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
		it.Event = new(StrategyNewAdmin)
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
func (it *StrategyNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyNewAdmin represents a NewAdmin event raised by the Strategy contract.
type StrategyNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Strategy *StrategyFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address) (*StrategyNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &StrategyNewAdminIterator{contract: _Strategy.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Strategy *StrategyFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *StrategyNewAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyNewAdmin)
				if err := _Strategy.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_Strategy *StrategyFilterer) ParseNewAdmin(log types.Log) (*StrategyNewAdmin, error) {
	event := new(StrategyNewAdmin)
	if err := _Strategy.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StrategyNewParametersIterator is returned from FilterNewParameters and is used to iterate over the raw logs and unpacked data for NewParameters events raised by the Strategy contract.
type StrategyNewParametersIterator struct {
	Event *StrategyNewParameters // Event containing the contract specifics and raw log

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
func (it *StrategyNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyNewParameters)
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
		it.Event = new(StrategyNewParameters)
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
func (it *StrategyNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyNewParameters represents a NewParameters event raised by the Strategy contract.
type StrategyNewParameters struct {
	A        *big.Int
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewParameters is a free log retrieval operation binding the contract event 0x752a27d1853eb7af3ee4ff764f2c4a51619386af721573dd3809e929c39db99e.
//
// Solidity: event NewParameters(uint256 A, uint256 fee, uint256 admin_fee)
func (_Strategy *StrategyFilterer) FilterNewParameters(opts *bind.FilterOpts) (*StrategyNewParametersIterator, error) {

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return &StrategyNewParametersIterator{contract: _Strategy.contract, event: "NewParameters", logs: logs, sub: sub}, nil
}

// WatchNewParameters is a free log subscription operation binding the contract event 0x752a27d1853eb7af3ee4ff764f2c4a51619386af721573dd3809e929c39db99e.
//
// Solidity: event NewParameters(uint256 A, uint256 fee, uint256 admin_fee)
func (_Strategy *StrategyFilterer) WatchNewParameters(opts *bind.WatchOpts, sink chan<- *StrategyNewParameters) (event.Subscription, error) {

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyNewParameters)
				if err := _Strategy.contract.UnpackLog(event, "NewParameters", log); err != nil {
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
func (_Strategy *StrategyFilterer) ParseNewParameters(log types.Log) (*StrategyNewParameters, error) {
	event := new(StrategyNewParameters)
	if err := _Strategy.contract.UnpackLog(event, "NewParameters", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StrategyRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Strategy contract.
type StrategyRemoveLiquidityIterator struct {
	Event *StrategyRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *StrategyRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyRemoveLiquidity)
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
		it.Event = new(StrategyRemoveLiquidity)
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
func (it *StrategyRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyRemoveLiquidity represents a RemoveLiquidity event raised by the Strategy contract.
type StrategyRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fees         [2]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_Strategy *StrategyFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*StrategyRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &StrategyRemoveLiquidityIterator{contract: _Strategy.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_Strategy *StrategyFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *StrategyRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyRemoveLiquidity)
				if err := _Strategy.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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
func (_Strategy *StrategyFilterer) ParseRemoveLiquidity(log types.Log) (*StrategyRemoveLiquidity, error) {
	event := new(StrategyRemoveLiquidity)
	if err := _Strategy.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StrategyRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the Strategy contract.
type StrategyRemoveLiquidityImbalanceIterator struct {
	Event *StrategyRemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *StrategyRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyRemoveLiquidityImbalance)
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
		it.Event = new(StrategyRemoveLiquidityImbalance)
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
func (it *StrategyRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the Strategy contract.
type StrategyRemoveLiquidityImbalance struct {
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
func (_Strategy *StrategyFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*StrategyRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &StrategyRemoveLiquidityImbalanceIterator{contract: _Strategy.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x2b5508378d7e19e0d5fa338419034731416c4f5b219a10379956f764317fd47e.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_Strategy *StrategyFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *StrategyRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyRemoveLiquidityImbalance)
				if err := _Strategy.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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
func (_Strategy *StrategyFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*StrategyRemoveLiquidityImbalance, error) {
	event := new(StrategyRemoveLiquidityImbalance)
	if err := _Strategy.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StrategyTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the Strategy contract.
type StrategyTokenExchangeIterator struct {
	Event *StrategyTokenExchange // Event containing the contract specifics and raw log

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
func (it *StrategyTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyTokenExchange)
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
		it.Event = new(StrategyTokenExchange)
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
func (it *StrategyTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyTokenExchange represents a TokenExchange event raised by the Strategy contract.
type StrategyTokenExchange struct {
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
func (_Strategy *StrategyFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*StrategyTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &StrategyTokenExchangeIterator{contract: _Strategy.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Strategy *StrategyFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *StrategyTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyTokenExchange)
				if err := _Strategy.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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
func (_Strategy *StrategyFilterer) ParseTokenExchange(log types.Log) (*StrategyTokenExchange, error) {
	event := new(StrategyTokenExchange)
	if err := _Strategy.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StrategyTokenExchangeUnderlyingIterator is returned from FilterTokenExchangeUnderlying and is used to iterate over the raw logs and unpacked data for TokenExchangeUnderlying events raised by the Strategy contract.
type StrategyTokenExchangeUnderlyingIterator struct {
	Event *StrategyTokenExchangeUnderlying // Event containing the contract specifics and raw log

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
func (it *StrategyTokenExchangeUnderlyingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyTokenExchangeUnderlying)
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
		it.Event = new(StrategyTokenExchangeUnderlying)
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
func (it *StrategyTokenExchangeUnderlyingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyTokenExchangeUnderlyingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyTokenExchangeUnderlying represents a TokenExchangeUnderlying event raised by the Strategy contract.
type StrategyTokenExchangeUnderlying struct {
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
func (_Strategy *StrategyFilterer) FilterTokenExchangeUnderlying(opts *bind.FilterOpts, buyer []common.Address) (*StrategyTokenExchangeUnderlyingIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return &StrategyTokenExchangeUnderlyingIterator{contract: _Strategy.contract, event: "TokenExchangeUnderlying", logs: logs, sub: sub}, nil
}

// WatchTokenExchangeUnderlying is a free log subscription operation binding the contract event 0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b.
//
// Solidity: event TokenExchangeUnderlying(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Strategy *StrategyFilterer) WatchTokenExchangeUnderlying(opts *bind.WatchOpts, sink chan<- *StrategyTokenExchangeUnderlying, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyTokenExchangeUnderlying)
				if err := _Strategy.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
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
func (_Strategy *StrategyFilterer) ParseTokenExchangeUnderlying(log types.Log) (*StrategyTokenExchangeUnderlying, error) {
	event := new(StrategyTokenExchangeUnderlying)
	if err := _Strategy.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
		return nil, err
	}
	return event, nil
}
