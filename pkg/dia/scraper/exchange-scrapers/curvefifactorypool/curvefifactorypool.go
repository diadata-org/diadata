// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curvefifactorypool

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

// CurvefifactorypoolMetaData contains all meta data concerning the Curvefifactorypool contract.
var CurvefifactorypoolMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_index\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewParameters\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"admin_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"mid_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"out_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"adjustment_step\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ma_half_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewParameters\",\"inputs\":[{\"name\":\"admin_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"mid_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"out_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"adjustment_step\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ma_half_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampAgamma\",\"inputs\":[{\"name\":\"initial_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"name\":\"current_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"current_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ClaimAdminFee\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true},{\"name\":\"tokens\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_weth\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange_underlying\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange_underlying\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange_extended\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"cb\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"name\":\"min_mint_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"name\":\"min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"name\":\"min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[2]\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claim_admin_fees\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"ramp_A_gamma\",\"inputs\":[{\"name\":\"future_A\",\"type\":\"uint256\"},{\"name\":\"future_gamma\",\"type\":\"uint256\"},{\"name\":\"future_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"stop_ramp_A_gamma\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_new_parameters\",\"inputs\":[{\"name\":\"_new_mid_fee\",\"type\":\"uint256\"},{\"name\":\"_new_out_fee\",\"type\":\"uint256\"},{\"name\":\"_new_admin_fee\",\"type\":\"uint256\"},{\"name\":\"_new_fee_gamma\",\"type\":\"uint256\"},{\"name\":\"_new_allowed_extra_profit\",\"type\":\"uint256\"},{\"name\":\"_new_adjustment_step\",\"type\":\"uint256\"},{\"name\":\"_new_ma_half_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_new_parameters\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revert_new_parameters\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_amount\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_withdraw_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lp_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"gamma\",\"type\":\"uint256\"},{\"name\":\"mid_fee\",\"type\":\"uint256\"},{\"name\":\"out_fee\",\"type\":\"uint256\"},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\"},{\"name\":\"fee_gamma\",\"type\":\"uint256\"},{\"name\":\"adjustment_step\",\"type\":\"uint256\"},{\"name\":\"admin_fee\",\"type\":\"uint256\"},{\"name\":\"ma_half_time\",\"type\":\"uint256\"},{\"name\":\"initial_price\",\"type\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_coins\",\"type\":\"address[2]\"},{\"name\":\"_precisions\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_scale\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_prices\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_prices_timestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_gamma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_gamma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowed_extra_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_allowed_extra_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_fee_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"adjustment_step\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_adjustment_step\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_half_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_ma_half_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"mid_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"out_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_mid_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_out_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"D\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xcp_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xcp_profit_a\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_actions_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// CurvefifactorypoolABI is the input ABI used to generate the binding from.
// Deprecated: Use CurvefifactorypoolMetaData.ABI instead.
var CurvefifactorypoolABI = CurvefifactorypoolMetaData.ABI

// Curvefifactorypool is an auto generated Go binding around an Ethereum contract.
type Curvefifactorypool struct {
	CurvefifactorypoolCaller     // Read-only binding to the contract
	CurvefifactorypoolTransactor // Write-only binding to the contract
	CurvefifactorypoolFilterer   // Log filterer for contract events
}

// CurvefifactorypoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurvefifactorypoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvefifactorypoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurvefifactorypoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvefifactorypoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurvefifactorypoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvefifactorypoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurvefifactorypoolSession struct {
	Contract     *Curvefifactorypool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CurvefifactorypoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurvefifactorypoolCallerSession struct {
	Contract *CurvefifactorypoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CurvefifactorypoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurvefifactorypoolTransactorSession struct {
	Contract     *CurvefifactorypoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CurvefifactorypoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurvefifactorypoolRaw struct {
	Contract *Curvefifactorypool // Generic contract binding to access the raw methods on
}

// CurvefifactorypoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurvefifactorypoolCallerRaw struct {
	Contract *CurvefifactorypoolCaller // Generic read-only contract binding to access the raw methods on
}

// CurvefifactorypoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurvefifactorypoolTransactorRaw struct {
	Contract *CurvefifactorypoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurvefifactorypool creates a new instance of Curvefifactorypool, bound to a specific deployed contract.
func NewCurvefifactorypool(address common.Address, backend bind.ContractBackend) (*Curvefifactorypool, error) {
	contract, err := bindCurvefifactorypool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Curvefifactorypool{CurvefifactorypoolCaller: CurvefifactorypoolCaller{contract: contract}, CurvefifactorypoolTransactor: CurvefifactorypoolTransactor{contract: contract}, CurvefifactorypoolFilterer: CurvefifactorypoolFilterer{contract: contract}}, nil
}

// NewCurvefifactorypoolCaller creates a new read-only instance of Curvefifactorypool, bound to a specific deployed contract.
func NewCurvefifactorypoolCaller(address common.Address, caller bind.ContractCaller) (*CurvefifactorypoolCaller, error) {
	contract, err := bindCurvefifactorypool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurvefifactorypoolCaller{contract: contract}, nil
}

// NewCurvefifactorypoolTransactor creates a new write-only instance of Curvefifactorypool, bound to a specific deployed contract.
func NewCurvefifactorypoolTransactor(address common.Address, transactor bind.ContractTransactor) (*CurvefifactorypoolTransactor, error) {
	contract, err := bindCurvefifactorypool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurvefifactorypoolTransactor{contract: contract}, nil
}

// NewCurvefifactorypoolFilterer creates a new log filterer instance of Curvefifactorypool, bound to a specific deployed contract.
func NewCurvefifactorypoolFilterer(address common.Address, filterer bind.ContractFilterer) (*CurvefifactorypoolFilterer, error) {
	contract, err := bindCurvefifactorypool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurvefifactorypoolFilterer{contract: contract}, nil
}

// bindCurvefifactorypool binds a generic wrapper to an already deployed contract.
func bindCurvefifactorypool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurvefifactorypoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curvefifactorypool *CurvefifactorypoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Curvefifactorypool.Contract.CurvefifactorypoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curvefifactorypool *CurvefifactorypoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.CurvefifactorypoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curvefifactorypool *CurvefifactorypoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.CurvefifactorypoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curvefifactorypool *CurvefifactorypoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Curvefifactorypool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curvefifactorypool *CurvefifactorypoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curvefifactorypool *CurvefifactorypoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) A() (*big.Int, error) {
	return _Curvefifactorypool.Contract.A(&_Curvefifactorypool.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) A() (*big.Int, error) {
	return _Curvefifactorypool.Contract.A(&_Curvefifactorypool.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) D(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "D")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) D() (*big.Int, error) {
	return _Curvefifactorypool.Contract.D(&_Curvefifactorypool.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) D() (*big.Int, error) {
	return _Curvefifactorypool.Contract.D(&_Curvefifactorypool.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) AdjustmentStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "adjustment_step")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) AdjustmentStep() (*big.Int, error) {
	return _Curvefifactorypool.Contract.AdjustmentStep(&_Curvefifactorypool.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) AdjustmentStep() (*big.Int, error) {
	return _Curvefifactorypool.Contract.AdjustmentStep(&_Curvefifactorypool.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "admin_actions_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) AdminActionsDeadline() (*big.Int, error) {
	return _Curvefifactorypool.Contract.AdminActionsDeadline(&_Curvefifactorypool.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _Curvefifactorypool.Contract.AdminActionsDeadline(&_Curvefifactorypool.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) AdminFee() (*big.Int, error) {
	return _Curvefifactorypool.Contract.AdminFee(&_Curvefifactorypool.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) AdminFee() (*big.Int, error) {
	return _Curvefifactorypool.Contract.AdminFee(&_Curvefifactorypool.CallOpts)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) AllowedExtraProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "allowed_extra_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) AllowedExtraProfit() (*big.Int, error) {
	return _Curvefifactorypool.Contract.AllowedExtraProfit(&_Curvefifactorypool.CallOpts)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) AllowedExtraProfit() (*big.Int, error) {
	return _Curvefifactorypool.Contract.AllowedExtraProfit(&_Curvefifactorypool.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Curvefifactorypool.Contract.Balances(&_Curvefifactorypool.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Curvefifactorypool.Contract.Balances(&_Curvefifactorypool.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x8d8ea727.
//
// Solidity: function calc_token_amount(uint256[2] amounts) view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [2]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "calc_token_amount", amounts)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x8d8ea727.
//
// Solidity: function calc_token_amount(uint256[2] amounts) view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) CalcTokenAmount(amounts [2]*big.Int) (*big.Int, error) {
	return _Curvefifactorypool.Contract.CalcTokenAmount(&_Curvefifactorypool.CallOpts, amounts)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x8d8ea727.
//
// Solidity: function calc_token_amount(uint256[2] amounts) view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) CalcTokenAmount(amounts [2]*big.Int) (*big.Int, error) {
	return _Curvefifactorypool.Contract.CalcTokenAmount(&_Curvefifactorypool.CallOpts, amounts)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "calc_withdraw_one_coin", token_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Curvefifactorypool.Contract.CalcWithdrawOneCoin(&_Curvefifactorypool.CallOpts, token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Curvefifactorypool.Contract.CalcWithdrawOneCoin(&_Curvefifactorypool.CallOpts, token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Curvefifactorypool *CurvefifactorypoolCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Curvefifactorypool *CurvefifactorypoolSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Curvefifactorypool.Contract.Coins(&_Curvefifactorypool.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Curvefifactorypool.Contract.Coins(&_Curvefifactorypool.CallOpts, arg0)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Curvefifactorypool *CurvefifactorypoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Curvefifactorypool *CurvefifactorypoolSession) Factory() (common.Address, error) {
	return _Curvefifactorypool.Contract.Factory(&_Curvefifactorypool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) Factory() (common.Address, error) {
	return _Curvefifactorypool.Contract.Factory(&_Curvefifactorypool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) Fee() (*big.Int, error) {
	return _Curvefifactorypool.Contract.Fee(&_Curvefifactorypool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) Fee() (*big.Int, error) {
	return _Curvefifactorypool.Contract.Fee(&_Curvefifactorypool.CallOpts)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) FeeGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "fee_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) FeeGamma() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FeeGamma(&_Curvefifactorypool.CallOpts)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) FeeGamma() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FeeGamma(&_Curvefifactorypool.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) FutureAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "future_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) FutureAGamma() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureAGamma(&_Curvefifactorypool.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) FutureAGamma() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureAGamma(&_Curvefifactorypool.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) FutureAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "future_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) FutureAGammaTime() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureAGammaTime(&_Curvefifactorypool.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) FutureAGammaTime() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureAGammaTime(&_Curvefifactorypool.CallOpts)
}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) FutureAdjustmentStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "future_adjustment_step")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) FutureAdjustmentStep() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureAdjustmentStep(&_Curvefifactorypool.CallOpts)
}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) FutureAdjustmentStep() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureAdjustmentStep(&_Curvefifactorypool.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "future_admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) FutureAdminFee() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureAdminFee(&_Curvefifactorypool.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) FutureAdminFee() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureAdminFee(&_Curvefifactorypool.CallOpts)
}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) FutureAllowedExtraProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "future_allowed_extra_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) FutureAllowedExtraProfit() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureAllowedExtraProfit(&_Curvefifactorypool.CallOpts)
}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) FutureAllowedExtraProfit() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureAllowedExtraProfit(&_Curvefifactorypool.CallOpts)
}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) FutureFeeGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "future_fee_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) FutureFeeGamma() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureFeeGamma(&_Curvefifactorypool.CallOpts)
}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) FutureFeeGamma() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureFeeGamma(&_Curvefifactorypool.CallOpts)
}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) FutureMaHalfTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "future_ma_half_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) FutureMaHalfTime() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureMaHalfTime(&_Curvefifactorypool.CallOpts)
}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) FutureMaHalfTime() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureMaHalfTime(&_Curvefifactorypool.CallOpts)
}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) FutureMidFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "future_mid_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) FutureMidFee() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureMidFee(&_Curvefifactorypool.CallOpts)
}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) FutureMidFee() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureMidFee(&_Curvefifactorypool.CallOpts)
}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) FutureOutFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "future_out_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) FutureOutFee() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureOutFee(&_Curvefifactorypool.CallOpts)
}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) FutureOutFee() (*big.Int, error) {
	return _Curvefifactorypool.Contract.FutureOutFee(&_Curvefifactorypool.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) Gamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) Gamma() (*big.Int, error) {
	return _Curvefifactorypool.Contract.Gamma(&_Curvefifactorypool.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) Gamma() (*big.Int, error) {
	return _Curvefifactorypool.Contract.Gamma(&_Curvefifactorypool.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Curvefifactorypool.Contract.GetDy(&_Curvefifactorypool.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Curvefifactorypool.Contract.GetDy(&_Curvefifactorypool.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) GetVirtualPrice() (*big.Int, error) {
	return _Curvefifactorypool.Contract.GetVirtualPrice(&_Curvefifactorypool.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _Curvefifactorypool.Contract.GetVirtualPrice(&_Curvefifactorypool.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) InitialAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "initial_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) InitialAGamma() (*big.Int, error) {
	return _Curvefifactorypool.Contract.InitialAGamma(&_Curvefifactorypool.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) InitialAGamma() (*big.Int, error) {
	return _Curvefifactorypool.Contract.InitialAGamma(&_Curvefifactorypool.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) InitialAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "initial_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) InitialAGammaTime() (*big.Int, error) {
	return _Curvefifactorypool.Contract.InitialAGammaTime(&_Curvefifactorypool.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) InitialAGammaTime() (*big.Int, error) {
	return _Curvefifactorypool.Contract.InitialAGammaTime(&_Curvefifactorypool.CallOpts)
}

// LastPrices is a free data retrieval call binding the contract method 0xc146bf94.
//
// Solidity: function last_prices() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) LastPrices(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "last_prices")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrices is a free data retrieval call binding the contract method 0xc146bf94.
//
// Solidity: function last_prices() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) LastPrices() (*big.Int, error) {
	return _Curvefifactorypool.Contract.LastPrices(&_Curvefifactorypool.CallOpts)
}

// LastPrices is a free data retrieval call binding the contract method 0xc146bf94.
//
// Solidity: function last_prices() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) LastPrices() (*big.Int, error) {
	return _Curvefifactorypool.Contract.LastPrices(&_Curvefifactorypool.CallOpts)
}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) LastPricesTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "last_prices_timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) LastPricesTimestamp() (*big.Int, error) {
	return _Curvefifactorypool.Contract.LastPricesTimestamp(&_Curvefifactorypool.CallOpts)
}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) LastPricesTimestamp() (*big.Int, error) {
	return _Curvefifactorypool.Contract.LastPricesTimestamp(&_Curvefifactorypool.CallOpts)
}

// LpPrice is a free data retrieval call binding the contract method 0x54f0f7d5.
//
// Solidity: function lp_price() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) LpPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "lp_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpPrice is a free data retrieval call binding the contract method 0x54f0f7d5.
//
// Solidity: function lp_price() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) LpPrice() (*big.Int, error) {
	return _Curvefifactorypool.Contract.LpPrice(&_Curvefifactorypool.CallOpts)
}

// LpPrice is a free data retrieval call binding the contract method 0x54f0f7d5.
//
// Solidity: function lp_price() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) LpPrice() (*big.Int, error) {
	return _Curvefifactorypool.Contract.LpPrice(&_Curvefifactorypool.CallOpts)
}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) MaHalfTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "ma_half_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) MaHalfTime() (*big.Int, error) {
	return _Curvefifactorypool.Contract.MaHalfTime(&_Curvefifactorypool.CallOpts)
}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) MaHalfTime() (*big.Int, error) {
	return _Curvefifactorypool.Contract.MaHalfTime(&_Curvefifactorypool.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) MidFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "mid_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) MidFee() (*big.Int, error) {
	return _Curvefifactorypool.Contract.MidFee(&_Curvefifactorypool.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) MidFee() (*big.Int, error) {
	return _Curvefifactorypool.Contract.MidFee(&_Curvefifactorypool.CallOpts)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) OutFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "out_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) OutFee() (*big.Int, error) {
	return _Curvefifactorypool.Contract.OutFee(&_Curvefifactorypool.CallOpts)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) OutFee() (*big.Int, error) {
	return _Curvefifactorypool.Contract.OutFee(&_Curvefifactorypool.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) PriceOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "price_oracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) PriceOracle() (*big.Int, error) {
	return _Curvefifactorypool.Contract.PriceOracle(&_Curvefifactorypool.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) PriceOracle() (*big.Int, error) {
	return _Curvefifactorypool.Contract.PriceOracle(&_Curvefifactorypool.CallOpts)
}

// PriceScale is a free data retrieval call binding the contract method 0xb9e8c9fd.
//
// Solidity: function price_scale() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) PriceScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "price_scale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceScale is a free data retrieval call binding the contract method 0xb9e8c9fd.
//
// Solidity: function price_scale() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) PriceScale() (*big.Int, error) {
	return _Curvefifactorypool.Contract.PriceScale(&_Curvefifactorypool.CallOpts)
}

// PriceScale is a free data retrieval call binding the contract method 0xb9e8c9fd.
//
// Solidity: function price_scale() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) PriceScale() (*big.Int, error) {
	return _Curvefifactorypool.Contract.PriceScale(&_Curvefifactorypool.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Curvefifactorypool *CurvefifactorypoolCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Curvefifactorypool *CurvefifactorypoolSession) Token() (common.Address, error) {
	return _Curvefifactorypool.Contract.Token(&_Curvefifactorypool.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) Token() (common.Address, error) {
	return _Curvefifactorypool.Contract.Token(&_Curvefifactorypool.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) VirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) VirtualPrice() (*big.Int, error) {
	return _Curvefifactorypool.Contract.VirtualPrice(&_Curvefifactorypool.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) VirtualPrice() (*big.Int, error) {
	return _Curvefifactorypool.Contract.VirtualPrice(&_Curvefifactorypool.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) XcpProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "xcp_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) XcpProfit() (*big.Int, error) {
	return _Curvefifactorypool.Contract.XcpProfit(&_Curvefifactorypool.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) XcpProfit() (*big.Int, error) {
	return _Curvefifactorypool.Contract.XcpProfit(&_Curvefifactorypool.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCaller) XcpProfitA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactorypool.contract.Call(opts, &out, "xcp_profit_a")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) XcpProfitA() (*big.Int, error) {
	return _Curvefifactorypool.Contract.XcpProfitA(&_Curvefifactorypool.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolCallerSession) XcpProfitA() (*big.Int, error) {
	return _Curvefifactorypool.Contract.XcpProfitA(&_Curvefifactorypool.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.AddLiquidity(&_Curvefifactorypool.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.AddLiquidity(&_Curvefifactorypool.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xee22be23.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool use_eth) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactor) AddLiquidity0(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "add_liquidity0", amounts, min_mint_amount, use_eth)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xee22be23.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool use_eth) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) AddLiquidity0(amounts [2]*big.Int, min_mint_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.AddLiquidity0(&_Curvefifactorypool.TransactOpts, amounts, min_mint_amount, use_eth)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xee22be23.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool use_eth) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) AddLiquidity0(amounts [2]*big.Int, min_mint_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.AddLiquidity0(&_Curvefifactorypool.TransactOpts, amounts, min_mint_amount, use_eth)
}

// AddLiquidity1 is a paid mutator transaction binding the contract method 0x7328333b.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool use_eth, address receiver) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactor) AddLiquidity1(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "add_liquidity1", amounts, min_mint_amount, use_eth, receiver)
}

// AddLiquidity1 is a paid mutator transaction binding the contract method 0x7328333b.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool use_eth, address receiver) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) AddLiquidity1(amounts [2]*big.Int, min_mint_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.AddLiquidity1(&_Curvefifactorypool.TransactOpts, amounts, min_mint_amount, use_eth, receiver)
}

// AddLiquidity1 is a paid mutator transaction binding the contract method 0x7328333b.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool use_eth, address receiver) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) AddLiquidity1(amounts [2]*big.Int, min_mint_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.AddLiquidity1(&_Curvefifactorypool.TransactOpts, amounts, min_mint_amount, use_eth, receiver)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactor) ApplyNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "apply_new_parameters")
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Curvefifactorypool *CurvefifactorypoolSession) ApplyNewParameters() (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.ApplyNewParameters(&_Curvefifactorypool.TransactOpts)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) ApplyNewParameters() (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.ApplyNewParameters(&_Curvefifactorypool.TransactOpts)
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactor) ClaimAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "claim_admin_fees")
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_Curvefifactorypool *CurvefifactorypoolSession) ClaimAdminFees() (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.ClaimAdminFees(&_Curvefifactorypool.TransactOpts)
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) ClaimAdminFees() (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.ClaimAdminFees(&_Curvefifactorypool.TransactOpts)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactor) CommitNewParameters(opts *bind.TransactOpts, _new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "commit_new_parameters", _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_Curvefifactorypool *CurvefifactorypoolSession) CommitNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.CommitNewParameters(&_Curvefifactorypool.TransactOpts, _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) CommitNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.CommitNewParameters(&_Curvefifactorypool.TransactOpts, _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.Exchange(&_Curvefifactorypool.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.Exchange(&_Curvefifactorypool.TransactOpts, i, j, dx, min_dy)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "exchange0", i, j, dx, min_dy, use_eth)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) Exchange0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.Exchange0(&_Curvefifactorypool.TransactOpts, i, j, dx, min_dy, use_eth)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) Exchange0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.Exchange0(&_Curvefifactorypool.TransactOpts, i, j, dx, min_dy, use_eth)
}

// Exchange1 is a paid mutator transaction binding the contract method 0xce7d6503.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address receiver) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactor) Exchange1(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "exchange1", i, j, dx, min_dy, use_eth, receiver)
}

// Exchange1 is a paid mutator transaction binding the contract method 0xce7d6503.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address receiver) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) Exchange1(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.Exchange1(&_Curvefifactorypool.TransactOpts, i, j, dx, min_dy, use_eth, receiver)
}

// Exchange1 is a paid mutator transaction binding the contract method 0xce7d6503.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address receiver) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) Exchange1(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.Exchange1(&_Curvefifactorypool.TransactOpts, i, j, dx, min_dy, use_eth, receiver)
}

// ExchangeExtended is a paid mutator transaction binding the contract method 0xdd96994f.
//
// Solidity: function exchange_extended(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address sender, address receiver, bytes32 cb) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactor) ExchangeExtended(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, sender common.Address, receiver common.Address, cb [32]byte) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "exchange_extended", i, j, dx, min_dy, use_eth, sender, receiver, cb)
}

// ExchangeExtended is a paid mutator transaction binding the contract method 0xdd96994f.
//
// Solidity: function exchange_extended(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address sender, address receiver, bytes32 cb) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) ExchangeExtended(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, sender common.Address, receiver common.Address, cb [32]byte) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.ExchangeExtended(&_Curvefifactorypool.TransactOpts, i, j, dx, min_dy, use_eth, sender, receiver, cb)
}

// ExchangeExtended is a paid mutator transaction binding the contract method 0xdd96994f.
//
// Solidity: function exchange_extended(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address sender, address receiver, bytes32 cb) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) ExchangeExtended(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, sender common.Address, receiver common.Address, cb [32]byte) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.ExchangeExtended(&_Curvefifactorypool.TransactOpts, i, j, dx, min_dy, use_eth, sender, receiver, cb)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "exchange_underlying", i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.ExchangeUnderlying(&_Curvefifactorypool.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.ExchangeUnderlying(&_Curvefifactorypool.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying0 is a paid mutator transaction binding the contract method 0xe2ad025a.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy, address receiver) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactor) ExchangeUnderlying0(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "exchange_underlying0", i, j, dx, min_dy, receiver)
}

// ExchangeUnderlying0 is a paid mutator transaction binding the contract method 0xe2ad025a.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy, address receiver) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) ExchangeUnderlying0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.ExchangeUnderlying0(&_Curvefifactorypool.TransactOpts, i, j, dx, min_dy, receiver)
}

// ExchangeUnderlying0 is a paid mutator transaction binding the contract method 0xe2ad025a.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy, address receiver) payable returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) ExchangeUnderlying0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.ExchangeUnderlying0(&_Curvefifactorypool.TransactOpts, i, j, dx, min_dy, receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0xa39e95c5.
//
// Solidity: function initialize(uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price, address _token, address[2] _coins, uint256 _precisions) returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactor) Initialize(opts *bind.TransactOpts, A *big.Int, gamma *big.Int, mid_fee *big.Int, out_fee *big.Int, allowed_extra_profit *big.Int, fee_gamma *big.Int, adjustment_step *big.Int, admin_fee *big.Int, ma_half_time *big.Int, initial_price *big.Int, _token common.Address, _coins [2]common.Address, _precisions *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "initialize", A, gamma, mid_fee, out_fee, allowed_extra_profit, fee_gamma, adjustment_step, admin_fee, ma_half_time, initial_price, _token, _coins, _precisions)
}

// Initialize is a paid mutator transaction binding the contract method 0xa39e95c5.
//
// Solidity: function initialize(uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price, address _token, address[2] _coins, uint256 _precisions) returns()
func (_Curvefifactorypool *CurvefifactorypoolSession) Initialize(A *big.Int, gamma *big.Int, mid_fee *big.Int, out_fee *big.Int, allowed_extra_profit *big.Int, fee_gamma *big.Int, adjustment_step *big.Int, admin_fee *big.Int, ma_half_time *big.Int, initial_price *big.Int, _token common.Address, _coins [2]common.Address, _precisions *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.Initialize(&_Curvefifactorypool.TransactOpts, A, gamma, mid_fee, out_fee, allowed_extra_profit, fee_gamma, adjustment_step, admin_fee, ma_half_time, initial_price, _token, _coins, _precisions)
}

// Initialize is a paid mutator transaction binding the contract method 0xa39e95c5.
//
// Solidity: function initialize(uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price, address _token, address[2] _coins, uint256 _precisions) returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) Initialize(A *big.Int, gamma *big.Int, mid_fee *big.Int, out_fee *big.Int, allowed_extra_profit *big.Int, fee_gamma *big.Int, adjustment_step *big.Int, admin_fee *big.Int, ma_half_time *big.Int, initial_price *big.Int, _token common.Address, _coins [2]common.Address, _precisions *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.Initialize(&_Curvefifactorypool.TransactOpts, A, gamma, mid_fee, out_fee, allowed_extra_profit, fee_gamma, adjustment_step, admin_fee, ma_half_time, initial_price, _token, _coins, _precisions)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactor) RampAGamma(opts *bind.TransactOpts, future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "ramp_A_gamma", future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_Curvefifactorypool *CurvefifactorypoolSession) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RampAGamma(&_Curvefifactorypool.TransactOpts, future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RampAGamma(&_Curvefifactorypool.TransactOpts, future_A, future_gamma, future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Curvefifactorypool *CurvefifactorypoolSession) RemoveLiquidity(_amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RemoveLiquidity(&_Curvefifactorypool.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RemoveLiquidity(&_Curvefifactorypool.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x269b5581.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, bool use_eth) returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactor) RemoveLiquidity0(opts *bind.TransactOpts, _amount *big.Int, min_amounts [2]*big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "remove_liquidity0", _amount, min_amounts, use_eth)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x269b5581.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, bool use_eth) returns()
func (_Curvefifactorypool *CurvefifactorypoolSession) RemoveLiquidity0(_amount *big.Int, min_amounts [2]*big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RemoveLiquidity0(&_Curvefifactorypool.TransactOpts, _amount, min_amounts, use_eth)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x269b5581.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, bool use_eth) returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) RemoveLiquidity0(_amount *big.Int, min_amounts [2]*big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RemoveLiquidity0(&_Curvefifactorypool.TransactOpts, _amount, min_amounts, use_eth)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x1808e84a.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, bool use_eth, address receiver) returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactor) RemoveLiquidity1(opts *bind.TransactOpts, _amount *big.Int, min_amounts [2]*big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "remove_liquidity1", _amount, min_amounts, use_eth, receiver)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x1808e84a.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, bool use_eth, address receiver) returns()
func (_Curvefifactorypool *CurvefifactorypoolSession) RemoveLiquidity1(_amount *big.Int, min_amounts [2]*big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RemoveLiquidity1(&_Curvefifactorypool.TransactOpts, _amount, min_amounts, use_eth, receiver)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x1808e84a.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, bool use_eth, address receiver) returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) RemoveLiquidity1(_amount *big.Int, min_amounts [2]*big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RemoveLiquidity1(&_Curvefifactorypool.TransactOpts, _amount, min_amounts, use_eth, receiver)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "remove_liquidity_one_coin", token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RemoveLiquidityOneCoin(&_Curvefifactorypool.TransactOpts, token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RemoveLiquidityOneCoin(&_Curvefifactorypool.TransactOpts, token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x8f15b6b5.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth) returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "remove_liquidity_one_coin0", token_amount, i, min_amount, use_eth)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x8f15b6b5.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth) returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) RemoveLiquidityOneCoin0(token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RemoveLiquidityOneCoin0(&_Curvefifactorypool.TransactOpts, token_amount, i, min_amount, use_eth)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x8f15b6b5.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth) returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) RemoveLiquidityOneCoin0(token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RemoveLiquidityOneCoin0(&_Curvefifactorypool.TransactOpts, token_amount, i, min_amount, use_eth)
}

// RemoveLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0x07329bcd.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth, address receiver) returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactor) RemoveLiquidityOneCoin1(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "remove_liquidity_one_coin1", token_amount, i, min_amount, use_eth, receiver)
}

// RemoveLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0x07329bcd.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth, address receiver) returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolSession) RemoveLiquidityOneCoin1(token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RemoveLiquidityOneCoin1(&_Curvefifactorypool.TransactOpts, token_amount, i, min_amount, use_eth, receiver)
}

// RemoveLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0x07329bcd.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth, address receiver) returns(uint256)
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) RemoveLiquidityOneCoin1(token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RemoveLiquidityOneCoin1(&_Curvefifactorypool.TransactOpts, token_amount, i, min_amount, use_eth, receiver)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Curvefifactorypool *CurvefifactorypoolSession) RevertNewParameters() (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RevertNewParameters(&_Curvefifactorypool.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.RevertNewParameters(&_Curvefifactorypool.TransactOpts)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactor) StopRampAGamma(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.Transact(opts, "stop_ramp_A_gamma")
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_Curvefifactorypool *CurvefifactorypoolSession) StopRampAGamma() (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.StopRampAGamma(&_Curvefifactorypool.TransactOpts)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) StopRampAGamma() (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.StopRampAGamma(&_Curvefifactorypool.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Curvefifactorypool.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Curvefifactorypool *CurvefifactorypoolSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.Fallback(&_Curvefifactorypool.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Curvefifactorypool *CurvefifactorypoolTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Curvefifactorypool.Contract.Fallback(&_Curvefifactorypool.TransactOpts, calldata)
}

// CurvefifactorypoolAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Curvefifactorypool contract.
type CurvefifactorypoolAddLiquidityIterator struct {
	Event *CurvefifactorypoolAddLiquidity // Event containing the contract specifics and raw log

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
func (it *CurvefifactorypoolAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactorypoolAddLiquidity)
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
		it.Event = new(CurvefifactorypoolAddLiquidity)
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
func (it *CurvefifactorypoolAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactorypoolAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactorypoolAddLiquidity represents a AddLiquidity event raised by the Curvefifactorypool contract.
type CurvefifactorypoolAddLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fee          *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x540ab385f9b5d450a27404172caade516b3ba3f4be88239ac56a2ad1de2a1f5a.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256 fee, uint256 token_supply)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurvefifactorypoolAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curvefifactorypool.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvefifactorypoolAddLiquidityIterator{contract: _Curvefifactorypool.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x540ab385f9b5d450a27404172caade516b3ba3f4be88239ac56a2ad1de2a1f5a.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256 fee, uint256 token_supply)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *CurvefifactorypoolAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curvefifactorypool.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactorypoolAddLiquidity)
				if err := _Curvefifactorypool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x540ab385f9b5d450a27404172caade516b3ba3f4be88239ac56a2ad1de2a1f5a.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256 fee, uint256 token_supply)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) ParseAddLiquidity(log types.Log) (*CurvefifactorypoolAddLiquidity, error) {
	event := new(CurvefifactorypoolAddLiquidity)
	if err := _Curvefifactorypool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactorypoolClaimAdminFeeIterator is returned from FilterClaimAdminFee and is used to iterate over the raw logs and unpacked data for ClaimAdminFee events raised by the Curvefifactorypool contract.
type CurvefifactorypoolClaimAdminFeeIterator struct {
	Event *CurvefifactorypoolClaimAdminFee // Event containing the contract specifics and raw log

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
func (it *CurvefifactorypoolClaimAdminFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactorypoolClaimAdminFee)
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
		it.Event = new(CurvefifactorypoolClaimAdminFee)
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
func (it *CurvefifactorypoolClaimAdminFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactorypoolClaimAdminFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactorypoolClaimAdminFee represents a ClaimAdminFee event raised by the Curvefifactorypool contract.
type CurvefifactorypoolClaimAdminFee struct {
	Admin  common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimAdminFee is a free log retrieval operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) FilterClaimAdminFee(opts *bind.FilterOpts, admin []common.Address) (*CurvefifactorypoolClaimAdminFeeIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Curvefifactorypool.contract.FilterLogs(opts, "ClaimAdminFee", adminRule)
	if err != nil {
		return nil, err
	}
	return &CurvefifactorypoolClaimAdminFeeIterator{contract: _Curvefifactorypool.contract, event: "ClaimAdminFee", logs: logs, sub: sub}, nil
}

// WatchClaimAdminFee is a free log subscription operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) WatchClaimAdminFee(opts *bind.WatchOpts, sink chan<- *CurvefifactorypoolClaimAdminFee, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Curvefifactorypool.contract.WatchLogs(opts, "ClaimAdminFee", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactorypoolClaimAdminFee)
				if err := _Curvefifactorypool.contract.UnpackLog(event, "ClaimAdminFee", log); err != nil {
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

// ParseClaimAdminFee is a log parse operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) ParseClaimAdminFee(log types.Log) (*CurvefifactorypoolClaimAdminFee, error) {
	event := new(CurvefifactorypoolClaimAdminFee)
	if err := _Curvefifactorypool.contract.UnpackLog(event, "ClaimAdminFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactorypoolCommitNewParametersIterator is returned from FilterCommitNewParameters and is used to iterate over the raw logs and unpacked data for CommitNewParameters events raised by the Curvefifactorypool contract.
type CurvefifactorypoolCommitNewParametersIterator struct {
	Event *CurvefifactorypoolCommitNewParameters // Event containing the contract specifics and raw log

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
func (it *CurvefifactorypoolCommitNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactorypoolCommitNewParameters)
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
		it.Event = new(CurvefifactorypoolCommitNewParameters)
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
func (it *CurvefifactorypoolCommitNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactorypoolCommitNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactorypoolCommitNewParameters represents a CommitNewParameters event raised by the Curvefifactorypool contract.
type CurvefifactorypoolCommitNewParameters struct {
	Deadline           *big.Int
	AdminFee           *big.Int
	MidFee             *big.Int
	OutFee             *big.Int
	FeeGamma           *big.Int
	AllowedExtraProfit *big.Int
	AdjustmentStep     *big.Int
	MaHalfTime         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCommitNewParameters is a free log retrieval operation binding the contract event 0x913fde9a37e1f8ab67876a4d0ce80790d764fcfc5692f4529526df9c6bdde553.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) FilterCommitNewParameters(opts *bind.FilterOpts, deadline []*big.Int) (*CurvefifactorypoolCommitNewParametersIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Curvefifactorypool.contract.FilterLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return &CurvefifactorypoolCommitNewParametersIterator{contract: _Curvefifactorypool.contract, event: "CommitNewParameters", logs: logs, sub: sub}, nil
}

// WatchCommitNewParameters is a free log subscription operation binding the contract event 0x913fde9a37e1f8ab67876a4d0ce80790d764fcfc5692f4529526df9c6bdde553.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) WatchCommitNewParameters(opts *bind.WatchOpts, sink chan<- *CurvefifactorypoolCommitNewParameters, deadline []*big.Int) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Curvefifactorypool.contract.WatchLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactorypoolCommitNewParameters)
				if err := _Curvefifactorypool.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
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

// ParseCommitNewParameters is a log parse operation binding the contract event 0x913fde9a37e1f8ab67876a4d0ce80790d764fcfc5692f4529526df9c6bdde553.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) ParseCommitNewParameters(log types.Log) (*CurvefifactorypoolCommitNewParameters, error) {
	event := new(CurvefifactorypoolCommitNewParameters)
	if err := _Curvefifactorypool.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactorypoolNewParametersIterator is returned from FilterNewParameters and is used to iterate over the raw logs and unpacked data for NewParameters events raised by the Curvefifactorypool contract.
type CurvefifactorypoolNewParametersIterator struct {
	Event *CurvefifactorypoolNewParameters // Event containing the contract specifics and raw log

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
func (it *CurvefifactorypoolNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactorypoolNewParameters)
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
		it.Event = new(CurvefifactorypoolNewParameters)
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
func (it *CurvefifactorypoolNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactorypoolNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactorypoolNewParameters represents a NewParameters event raised by the Curvefifactorypool contract.
type CurvefifactorypoolNewParameters struct {
	AdminFee           *big.Int
	MidFee             *big.Int
	OutFee             *big.Int
	FeeGamma           *big.Int
	AllowedExtraProfit *big.Int
	AdjustmentStep     *big.Int
	MaHalfTime         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewParameters is a free log retrieval operation binding the contract event 0x1c65bbdc939f346e5d6f0bde1f072819947438d4fc7b182cc59c2f6dc5504087.
//
// Solidity: event NewParameters(uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) FilterNewParameters(opts *bind.FilterOpts) (*CurvefifactorypoolNewParametersIterator, error) {

	logs, sub, err := _Curvefifactorypool.contract.FilterLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return &CurvefifactorypoolNewParametersIterator{contract: _Curvefifactorypool.contract, event: "NewParameters", logs: logs, sub: sub}, nil
}

// WatchNewParameters is a free log subscription operation binding the contract event 0x1c65bbdc939f346e5d6f0bde1f072819947438d4fc7b182cc59c2f6dc5504087.
//
// Solidity: event NewParameters(uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) WatchNewParameters(opts *bind.WatchOpts, sink chan<- *CurvefifactorypoolNewParameters) (event.Subscription, error) {

	logs, sub, err := _Curvefifactorypool.contract.WatchLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactorypoolNewParameters)
				if err := _Curvefifactorypool.contract.UnpackLog(event, "NewParameters", log); err != nil {
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

// ParseNewParameters is a log parse operation binding the contract event 0x1c65bbdc939f346e5d6f0bde1f072819947438d4fc7b182cc59c2f6dc5504087.
//
// Solidity: event NewParameters(uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) ParseNewParameters(log types.Log) (*CurvefifactorypoolNewParameters, error) {
	event := new(CurvefifactorypoolNewParameters)
	if err := _Curvefifactorypool.contract.UnpackLog(event, "NewParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactorypoolRampAgammaIterator is returned from FilterRampAgamma and is used to iterate over the raw logs and unpacked data for RampAgamma events raised by the Curvefifactorypool contract.
type CurvefifactorypoolRampAgammaIterator struct {
	Event *CurvefifactorypoolRampAgamma // Event containing the contract specifics and raw log

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
func (it *CurvefifactorypoolRampAgammaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactorypoolRampAgamma)
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
		it.Event = new(CurvefifactorypoolRampAgamma)
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
func (it *CurvefifactorypoolRampAgammaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactorypoolRampAgammaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactorypoolRampAgamma represents a RampAgamma event raised by the Curvefifactorypool contract.
type CurvefifactorypoolRampAgamma struct {
	InitialA     *big.Int
	FutureA      *big.Int
	InitialGamma *big.Int
	FutureGamma  *big.Int
	InitialTime  *big.Int
	FutureTime   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRampAgamma is a free log retrieval operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) FilterRampAgamma(opts *bind.FilterOpts) (*CurvefifactorypoolRampAgammaIterator, error) {

	logs, sub, err := _Curvefifactorypool.contract.FilterLogs(opts, "RampAgamma")
	if err != nil {
		return nil, err
	}
	return &CurvefifactorypoolRampAgammaIterator{contract: _Curvefifactorypool.contract, event: "RampAgamma", logs: logs, sub: sub}, nil
}

// WatchRampAgamma is a free log subscription operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) WatchRampAgamma(opts *bind.WatchOpts, sink chan<- *CurvefifactorypoolRampAgamma) (event.Subscription, error) {

	logs, sub, err := _Curvefifactorypool.contract.WatchLogs(opts, "RampAgamma")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactorypoolRampAgamma)
				if err := _Curvefifactorypool.contract.UnpackLog(event, "RampAgamma", log); err != nil {
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

// ParseRampAgamma is a log parse operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) ParseRampAgamma(log types.Log) (*CurvefifactorypoolRampAgamma, error) {
	event := new(CurvefifactorypoolRampAgamma)
	if err := _Curvefifactorypool.contract.UnpackLog(event, "RampAgamma", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactorypoolRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Curvefifactorypool contract.
type CurvefifactorypoolRemoveLiquidityIterator struct {
	Event *CurvefifactorypoolRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *CurvefifactorypoolRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactorypoolRemoveLiquidity)
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
		it.Event = new(CurvefifactorypoolRemoveLiquidity)
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
func (it *CurvefifactorypoolRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactorypoolRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactorypoolRemoveLiquidity represents a RemoveLiquidity event raised by the Curvefifactorypool contract.
type CurvefifactorypoolRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xdd3c0336a16f1b64f172b7bb0dad5b2b3c7c76f91e8c4aafd6aae60dce800153.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256 token_supply)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurvefifactorypoolRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curvefifactorypool.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvefifactorypoolRemoveLiquidityIterator{contract: _Curvefifactorypool.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xdd3c0336a16f1b64f172b7bb0dad5b2b3c7c76f91e8c4aafd6aae60dce800153.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256 token_supply)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *CurvefifactorypoolRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curvefifactorypool.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactorypoolRemoveLiquidity)
				if err := _Curvefifactorypool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0xdd3c0336a16f1b64f172b7bb0dad5b2b3c7c76f91e8c4aafd6aae60dce800153.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256 token_supply)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) ParseRemoveLiquidity(log types.Log) (*CurvefifactorypoolRemoveLiquidity, error) {
	event := new(CurvefifactorypoolRemoveLiquidity)
	if err := _Curvefifactorypool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactorypoolRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the Curvefifactorypool contract.
type CurvefifactorypoolRemoveLiquidityOneIterator struct {
	Event *CurvefifactorypoolRemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *CurvefifactorypoolRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactorypoolRemoveLiquidityOne)
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
		it.Event = new(CurvefifactorypoolRemoveLiquidityOne)
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
func (it *CurvefifactorypoolRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactorypoolRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactorypoolRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the Curvefifactorypool contract.
type CurvefifactorypoolRemoveLiquidityOne struct {
	Provider    common.Address
	TokenAmount *big.Int
	CoinIndex   *big.Int
	CoinAmount  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*CurvefifactorypoolRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curvefifactorypool.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvefifactorypoolRemoveLiquidityOneIterator{contract: _Curvefifactorypool.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *CurvefifactorypoolRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curvefifactorypool.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactorypoolRemoveLiquidityOne)
				if err := _Curvefifactorypool.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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

// ParseRemoveLiquidityOne is a log parse operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) ParseRemoveLiquidityOne(log types.Log) (*CurvefifactorypoolRemoveLiquidityOne, error) {
	event := new(CurvefifactorypoolRemoveLiquidityOne)
	if err := _Curvefifactorypool.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactorypoolStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the Curvefifactorypool contract.
type CurvefifactorypoolStopRampAIterator struct {
	Event *CurvefifactorypoolStopRampA // Event containing the contract specifics and raw log

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
func (it *CurvefifactorypoolStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactorypoolStopRampA)
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
		it.Event = new(CurvefifactorypoolStopRampA)
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
func (it *CurvefifactorypoolStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactorypoolStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactorypoolStopRampA represents a StopRampA event raised by the Curvefifactorypool contract.
type CurvefifactorypoolStopRampA struct {
	CurrentA     *big.Int
	CurrentGamma *big.Int
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) FilterStopRampA(opts *bind.FilterOpts) (*CurvefifactorypoolStopRampAIterator, error) {

	logs, sub, err := _Curvefifactorypool.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &CurvefifactorypoolStopRampAIterator{contract: _Curvefifactorypool.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *CurvefifactorypoolStopRampA) (event.Subscription, error) {

	logs, sub, err := _Curvefifactorypool.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactorypoolStopRampA)
				if err := _Curvefifactorypool.contract.UnpackLog(event, "StopRampA", log); err != nil {
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

// ParseStopRampA is a log parse operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) ParseStopRampA(log types.Log) (*CurvefifactorypoolStopRampA, error) {
	event := new(CurvefifactorypoolStopRampA)
	if err := _Curvefifactorypool.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactorypoolTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the Curvefifactorypool contract.
type CurvefifactorypoolTokenExchangeIterator struct {
	Event *CurvefifactorypoolTokenExchange // Event containing the contract specifics and raw log

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
func (it *CurvefifactorypoolTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactorypoolTokenExchange)
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
		it.Event = new(CurvefifactorypoolTokenExchange)
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
func (it *CurvefifactorypoolTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactorypoolTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactorypoolTokenExchange represents a TokenExchange event raised by the Curvefifactorypool contract.
type CurvefifactorypoolTokenExchange struct {
	Buyer        common.Address
	SoldId       *big.Int
	TokensSold   *big.Int
	BoughtId     *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenExchange is a free log retrieval operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*CurvefifactorypoolTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Curvefifactorypool.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &CurvefifactorypoolTokenExchangeIterator{contract: _Curvefifactorypool.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *CurvefifactorypoolTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Curvefifactorypool.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactorypoolTokenExchange)
				if err := _Curvefifactorypool.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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

// ParseTokenExchange is a log parse operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_Curvefifactorypool *CurvefifactorypoolFilterer) ParseTokenExchange(log types.Log) (*CurvefifactorypoolTokenExchange, error) {
	event := new(CurvefifactorypoolTokenExchange)
	if err := _Curvefifactorypool.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
