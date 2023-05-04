// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curvefifactory

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

// CurvefifactoryMetaData contains all meta data concerning the Curvefifactory contract.
var CurvefifactoryMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_index\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewParameters\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"admin_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"mid_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"out_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"adjustment_step\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ma_half_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewParameters\",\"inputs\":[{\"name\":\"admin_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"mid_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"out_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"adjustment_step\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ma_half_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampAgamma\",\"inputs\":[{\"name\":\"initial_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"name\":\"current_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"current_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ClaimAdminFee\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true},{\"name\":\"tokens\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_weth\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange_underlying\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange_underlying\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange_extended\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"cb\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"name\":\"min_mint_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"name\":\"min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"name\":\"min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[2]\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claim_admin_fees\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"ramp_A_gamma\",\"inputs\":[{\"name\":\"future_A\",\"type\":\"uint256\"},{\"name\":\"future_gamma\",\"type\":\"uint256\"},{\"name\":\"future_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"stop_ramp_A_gamma\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_new_parameters\",\"inputs\":[{\"name\":\"_new_mid_fee\",\"type\":\"uint256\"},{\"name\":\"_new_out_fee\",\"type\":\"uint256\"},{\"name\":\"_new_admin_fee\",\"type\":\"uint256\"},{\"name\":\"_new_fee_gamma\",\"type\":\"uint256\"},{\"name\":\"_new_allowed_extra_profit\",\"type\":\"uint256\"},{\"name\":\"_new_adjustment_step\",\"type\":\"uint256\"},{\"name\":\"_new_ma_half_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_new_parameters\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revert_new_parameters\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_amount\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_withdraw_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lp_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"gamma\",\"type\":\"uint256\"},{\"name\":\"mid_fee\",\"type\":\"uint256\"},{\"name\":\"out_fee\",\"type\":\"uint256\"},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\"},{\"name\":\"fee_gamma\",\"type\":\"uint256\"},{\"name\":\"adjustment_step\",\"type\":\"uint256\"},{\"name\":\"admin_fee\",\"type\":\"uint256\"},{\"name\":\"ma_half_time\",\"type\":\"uint256\"},{\"name\":\"initial_price\",\"type\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_coins\",\"type\":\"address[2]\"},{\"name\":\"_precisions\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_scale\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_prices\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_prices_timestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_gamma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_gamma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowed_extra_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_allowed_extra_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_fee_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"adjustment_step\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_adjustment_step\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_half_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_ma_half_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"mid_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"out_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_mid_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_out_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"D\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xcp_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xcp_profit_a\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_actions_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// CurvefifactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CurvefifactoryMetaData.ABI instead.
var CurvefifactoryABI = CurvefifactoryMetaData.ABI

// Curvefifactory is an auto generated Go binding around an Ethereum contract.
type Curvefifactory struct {
	CurvefifactoryCaller     // Read-only binding to the contract
	CurvefifactoryTransactor // Write-only binding to the contract
	CurvefifactoryFilterer   // Log filterer for contract events
}

// CurvefifactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurvefifactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvefifactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurvefifactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvefifactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurvefifactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvefifactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurvefifactorySession struct {
	Contract     *Curvefifactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurvefifactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurvefifactoryCallerSession struct {
	Contract *CurvefifactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CurvefifactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurvefifactoryTransactorSession struct {
	Contract     *CurvefifactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CurvefifactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurvefifactoryRaw struct {
	Contract *Curvefifactory // Generic contract binding to access the raw methods on
}

// CurvefifactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurvefifactoryCallerRaw struct {
	Contract *CurvefifactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CurvefifactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurvefifactoryTransactorRaw struct {
	Contract *CurvefifactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurvefifactory creates a new instance of Curvefifactory, bound to a specific deployed contract.
func NewCurvefifactory(address common.Address, backend bind.ContractBackend) (*Curvefifactory, error) {
	contract, err := bindCurvefifactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Curvefifactory{CurvefifactoryCaller: CurvefifactoryCaller{contract: contract}, CurvefifactoryTransactor: CurvefifactoryTransactor{contract: contract}, CurvefifactoryFilterer: CurvefifactoryFilterer{contract: contract}}, nil
}

// NewCurvefifactoryCaller creates a new read-only instance of Curvefifactory, bound to a specific deployed contract.
func NewCurvefifactoryCaller(address common.Address, caller bind.ContractCaller) (*CurvefifactoryCaller, error) {
	contract, err := bindCurvefifactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurvefifactoryCaller{contract: contract}, nil
}

// NewCurvefifactoryTransactor creates a new write-only instance of Curvefifactory, bound to a specific deployed contract.
func NewCurvefifactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CurvefifactoryTransactor, error) {
	contract, err := bindCurvefifactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurvefifactoryTransactor{contract: contract}, nil
}

// NewCurvefifactoryFilterer creates a new log filterer instance of Curvefifactory, bound to a specific deployed contract.
func NewCurvefifactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CurvefifactoryFilterer, error) {
	contract, err := bindCurvefifactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurvefifactoryFilterer{contract: contract}, nil
}

// bindCurvefifactory binds a generic wrapper to an already deployed contract.
func bindCurvefifactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurvefifactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curvefifactory *CurvefifactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Curvefifactory.Contract.CurvefifactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curvefifactory *CurvefifactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefifactory.Contract.CurvefifactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curvefifactory *CurvefifactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curvefifactory.Contract.CurvefifactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curvefifactory *CurvefifactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Curvefifactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curvefifactory *CurvefifactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefifactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curvefifactory *CurvefifactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curvefifactory.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) A() (*big.Int, error) {
	return _Curvefifactory.Contract.A(&_Curvefifactory.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) A() (*big.Int, error) {
	return _Curvefifactory.Contract.A(&_Curvefifactory.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) D(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "D")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) D() (*big.Int, error) {
	return _Curvefifactory.Contract.D(&_Curvefifactory.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) D() (*big.Int, error) {
	return _Curvefifactory.Contract.D(&_Curvefifactory.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) AdjustmentStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "adjustment_step")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) AdjustmentStep() (*big.Int, error) {
	return _Curvefifactory.Contract.AdjustmentStep(&_Curvefifactory.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) AdjustmentStep() (*big.Int, error) {
	return _Curvefifactory.Contract.AdjustmentStep(&_Curvefifactory.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "admin_actions_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) AdminActionsDeadline() (*big.Int, error) {
	return _Curvefifactory.Contract.AdminActionsDeadline(&_Curvefifactory.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _Curvefifactory.Contract.AdminActionsDeadline(&_Curvefifactory.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) AdminFee() (*big.Int, error) {
	return _Curvefifactory.Contract.AdminFee(&_Curvefifactory.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) AdminFee() (*big.Int, error) {
	return _Curvefifactory.Contract.AdminFee(&_Curvefifactory.CallOpts)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) AllowedExtraProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "allowed_extra_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) AllowedExtraProfit() (*big.Int, error) {
	return _Curvefifactory.Contract.AllowedExtraProfit(&_Curvefifactory.CallOpts)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) AllowedExtraProfit() (*big.Int, error) {
	return _Curvefifactory.Contract.AllowedExtraProfit(&_Curvefifactory.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Curvefifactory.Contract.Balances(&_Curvefifactory.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Curvefifactory.Contract.Balances(&_Curvefifactory.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x8d8ea727.
//
// Solidity: function calc_token_amount(uint256[2] amounts) view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [2]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "calc_token_amount", amounts)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x8d8ea727.
//
// Solidity: function calc_token_amount(uint256[2] amounts) view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) CalcTokenAmount(amounts [2]*big.Int) (*big.Int, error) {
	return _Curvefifactory.Contract.CalcTokenAmount(&_Curvefifactory.CallOpts, amounts)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x8d8ea727.
//
// Solidity: function calc_token_amount(uint256[2] amounts) view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) CalcTokenAmount(amounts [2]*big.Int) (*big.Int, error) {
	return _Curvefifactory.Contract.CalcTokenAmount(&_Curvefifactory.CallOpts, amounts)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "calc_withdraw_one_coin", token_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Curvefifactory.Contract.CalcWithdrawOneCoin(&_Curvefifactory.CallOpts, token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Curvefifactory.Contract.CalcWithdrawOneCoin(&_Curvefifactory.CallOpts, token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Curvefifactory *CurvefifactoryCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Curvefifactory *CurvefifactorySession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Curvefifactory.Contract.Coins(&_Curvefifactory.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Curvefifactory *CurvefifactoryCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Curvefifactory.Contract.Coins(&_Curvefifactory.CallOpts, arg0)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Curvefifactory *CurvefifactoryCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Curvefifactory *CurvefifactorySession) Factory() (common.Address, error) {
	return _Curvefifactory.Contract.Factory(&_Curvefifactory.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Curvefifactory *CurvefifactoryCallerSession) Factory() (common.Address, error) {
	return _Curvefifactory.Contract.Factory(&_Curvefifactory.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) Fee() (*big.Int, error) {
	return _Curvefifactory.Contract.Fee(&_Curvefifactory.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) Fee() (*big.Int, error) {
	return _Curvefifactory.Contract.Fee(&_Curvefifactory.CallOpts)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) FeeGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "fee_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) FeeGamma() (*big.Int, error) {
	return _Curvefifactory.Contract.FeeGamma(&_Curvefifactory.CallOpts)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) FeeGamma() (*big.Int, error) {
	return _Curvefifactory.Contract.FeeGamma(&_Curvefifactory.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) FutureAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "future_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) FutureAGamma() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureAGamma(&_Curvefifactory.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) FutureAGamma() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureAGamma(&_Curvefifactory.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) FutureAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "future_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) FutureAGammaTime() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureAGammaTime(&_Curvefifactory.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) FutureAGammaTime() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureAGammaTime(&_Curvefifactory.CallOpts)
}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) FutureAdjustmentStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "future_adjustment_step")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) FutureAdjustmentStep() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureAdjustmentStep(&_Curvefifactory.CallOpts)
}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) FutureAdjustmentStep() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureAdjustmentStep(&_Curvefifactory.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "future_admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) FutureAdminFee() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureAdminFee(&_Curvefifactory.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) FutureAdminFee() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureAdminFee(&_Curvefifactory.CallOpts)
}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) FutureAllowedExtraProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "future_allowed_extra_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) FutureAllowedExtraProfit() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureAllowedExtraProfit(&_Curvefifactory.CallOpts)
}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) FutureAllowedExtraProfit() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureAllowedExtraProfit(&_Curvefifactory.CallOpts)
}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) FutureFeeGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "future_fee_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) FutureFeeGamma() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureFeeGamma(&_Curvefifactory.CallOpts)
}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) FutureFeeGamma() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureFeeGamma(&_Curvefifactory.CallOpts)
}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) FutureMaHalfTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "future_ma_half_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) FutureMaHalfTime() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureMaHalfTime(&_Curvefifactory.CallOpts)
}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) FutureMaHalfTime() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureMaHalfTime(&_Curvefifactory.CallOpts)
}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) FutureMidFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "future_mid_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) FutureMidFee() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureMidFee(&_Curvefifactory.CallOpts)
}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) FutureMidFee() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureMidFee(&_Curvefifactory.CallOpts)
}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) FutureOutFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "future_out_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) FutureOutFee() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureOutFee(&_Curvefifactory.CallOpts)
}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) FutureOutFee() (*big.Int, error) {
	return _Curvefifactory.Contract.FutureOutFee(&_Curvefifactory.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) Gamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) Gamma() (*big.Int, error) {
	return _Curvefifactory.Contract.Gamma(&_Curvefifactory.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) Gamma() (*big.Int, error) {
	return _Curvefifactory.Contract.Gamma(&_Curvefifactory.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Curvefifactory.Contract.GetDy(&_Curvefifactory.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Curvefifactory.Contract.GetDy(&_Curvefifactory.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) GetVirtualPrice() (*big.Int, error) {
	return _Curvefifactory.Contract.GetVirtualPrice(&_Curvefifactory.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _Curvefifactory.Contract.GetVirtualPrice(&_Curvefifactory.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) InitialAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "initial_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) InitialAGamma() (*big.Int, error) {
	return _Curvefifactory.Contract.InitialAGamma(&_Curvefifactory.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) InitialAGamma() (*big.Int, error) {
	return _Curvefifactory.Contract.InitialAGamma(&_Curvefifactory.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) InitialAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "initial_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) InitialAGammaTime() (*big.Int, error) {
	return _Curvefifactory.Contract.InitialAGammaTime(&_Curvefifactory.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) InitialAGammaTime() (*big.Int, error) {
	return _Curvefifactory.Contract.InitialAGammaTime(&_Curvefifactory.CallOpts)
}

// LastPrices is a free data retrieval call binding the contract method 0xc146bf94.
//
// Solidity: function last_prices() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) LastPrices(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "last_prices")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrices is a free data retrieval call binding the contract method 0xc146bf94.
//
// Solidity: function last_prices() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) LastPrices() (*big.Int, error) {
	return _Curvefifactory.Contract.LastPrices(&_Curvefifactory.CallOpts)
}

// LastPrices is a free data retrieval call binding the contract method 0xc146bf94.
//
// Solidity: function last_prices() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) LastPrices() (*big.Int, error) {
	return _Curvefifactory.Contract.LastPrices(&_Curvefifactory.CallOpts)
}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) LastPricesTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "last_prices_timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) LastPricesTimestamp() (*big.Int, error) {
	return _Curvefifactory.Contract.LastPricesTimestamp(&_Curvefifactory.CallOpts)
}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) LastPricesTimestamp() (*big.Int, error) {
	return _Curvefifactory.Contract.LastPricesTimestamp(&_Curvefifactory.CallOpts)
}

// LpPrice is a free data retrieval call binding the contract method 0x54f0f7d5.
//
// Solidity: function lp_price() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) LpPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "lp_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpPrice is a free data retrieval call binding the contract method 0x54f0f7d5.
//
// Solidity: function lp_price() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) LpPrice() (*big.Int, error) {
	return _Curvefifactory.Contract.LpPrice(&_Curvefifactory.CallOpts)
}

// LpPrice is a free data retrieval call binding the contract method 0x54f0f7d5.
//
// Solidity: function lp_price() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) LpPrice() (*big.Int, error) {
	return _Curvefifactory.Contract.LpPrice(&_Curvefifactory.CallOpts)
}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) MaHalfTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "ma_half_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) MaHalfTime() (*big.Int, error) {
	return _Curvefifactory.Contract.MaHalfTime(&_Curvefifactory.CallOpts)
}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) MaHalfTime() (*big.Int, error) {
	return _Curvefifactory.Contract.MaHalfTime(&_Curvefifactory.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) MidFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "mid_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) MidFee() (*big.Int, error) {
	return _Curvefifactory.Contract.MidFee(&_Curvefifactory.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) MidFee() (*big.Int, error) {
	return _Curvefifactory.Contract.MidFee(&_Curvefifactory.CallOpts)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) OutFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "out_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) OutFee() (*big.Int, error) {
	return _Curvefifactory.Contract.OutFee(&_Curvefifactory.CallOpts)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) OutFee() (*big.Int, error) {
	return _Curvefifactory.Contract.OutFee(&_Curvefifactory.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) PriceOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "price_oracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) PriceOracle() (*big.Int, error) {
	return _Curvefifactory.Contract.PriceOracle(&_Curvefifactory.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) PriceOracle() (*big.Int, error) {
	return _Curvefifactory.Contract.PriceOracle(&_Curvefifactory.CallOpts)
}

// PriceScale is a free data retrieval call binding the contract method 0xb9e8c9fd.
//
// Solidity: function price_scale() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) PriceScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "price_scale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceScale is a free data retrieval call binding the contract method 0xb9e8c9fd.
//
// Solidity: function price_scale() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) PriceScale() (*big.Int, error) {
	return _Curvefifactory.Contract.PriceScale(&_Curvefifactory.CallOpts)
}

// PriceScale is a free data retrieval call binding the contract method 0xb9e8c9fd.
//
// Solidity: function price_scale() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) PriceScale() (*big.Int, error) {
	return _Curvefifactory.Contract.PriceScale(&_Curvefifactory.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Curvefifactory *CurvefifactoryCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Curvefifactory *CurvefifactorySession) Token() (common.Address, error) {
	return _Curvefifactory.Contract.Token(&_Curvefifactory.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Curvefifactory *CurvefifactoryCallerSession) Token() (common.Address, error) {
	return _Curvefifactory.Contract.Token(&_Curvefifactory.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) VirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) VirtualPrice() (*big.Int, error) {
	return _Curvefifactory.Contract.VirtualPrice(&_Curvefifactory.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) VirtualPrice() (*big.Int, error) {
	return _Curvefifactory.Contract.VirtualPrice(&_Curvefifactory.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) XcpProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "xcp_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) XcpProfit() (*big.Int, error) {
	return _Curvefifactory.Contract.XcpProfit(&_Curvefifactory.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) XcpProfit() (*big.Int, error) {
	return _Curvefifactory.Contract.XcpProfit(&_Curvefifactory.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCaller) XcpProfitA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefifactory.contract.Call(opts, &out, "xcp_profit_a")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_Curvefifactory *CurvefifactorySession) XcpProfitA() (*big.Int, error) {
	return _Curvefifactory.Contract.XcpProfitA(&_Curvefifactory.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_Curvefifactory *CurvefifactoryCallerSession) XcpProfitA() (*big.Int, error) {
	return _Curvefifactory.Contract.XcpProfitA(&_Curvefifactory.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) payable returns(uint256)
func (_Curvefifactory *CurvefifactorySession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.AddLiquidity(&_Curvefifactory.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactorSession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.AddLiquidity(&_Curvefifactory.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xee22be23.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool use_eth) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactor) AddLiquidity0(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "add_liquidity0", amounts, min_mint_amount, use_eth)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xee22be23.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool use_eth) payable returns(uint256)
func (_Curvefifactory *CurvefifactorySession) AddLiquidity0(amounts [2]*big.Int, min_mint_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactory.Contract.AddLiquidity0(&_Curvefifactory.TransactOpts, amounts, min_mint_amount, use_eth)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xee22be23.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool use_eth) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactorSession) AddLiquidity0(amounts [2]*big.Int, min_mint_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactory.Contract.AddLiquidity0(&_Curvefifactory.TransactOpts, amounts, min_mint_amount, use_eth)
}

// AddLiquidity1 is a paid mutator transaction binding the contract method 0x7328333b.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool use_eth, address receiver) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactor) AddLiquidity1(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "add_liquidity1", amounts, min_mint_amount, use_eth, receiver)
}

// AddLiquidity1 is a paid mutator transaction binding the contract method 0x7328333b.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool use_eth, address receiver) payable returns(uint256)
func (_Curvefifactory *CurvefifactorySession) AddLiquidity1(amounts [2]*big.Int, min_mint_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactory.Contract.AddLiquidity1(&_Curvefifactory.TransactOpts, amounts, min_mint_amount, use_eth, receiver)
}

// AddLiquidity1 is a paid mutator transaction binding the contract method 0x7328333b.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool use_eth, address receiver) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactorSession) AddLiquidity1(amounts [2]*big.Int, min_mint_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactory.Contract.AddLiquidity1(&_Curvefifactory.TransactOpts, amounts, min_mint_amount, use_eth, receiver)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Curvefifactory *CurvefifactoryTransactor) ApplyNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "apply_new_parameters")
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Curvefifactory *CurvefifactorySession) ApplyNewParameters() (*types.Transaction, error) {
	return _Curvefifactory.Contract.ApplyNewParameters(&_Curvefifactory.TransactOpts)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Curvefifactory *CurvefifactoryTransactorSession) ApplyNewParameters() (*types.Transaction, error) {
	return _Curvefifactory.Contract.ApplyNewParameters(&_Curvefifactory.TransactOpts)
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_Curvefifactory *CurvefifactoryTransactor) ClaimAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "claim_admin_fees")
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_Curvefifactory *CurvefifactorySession) ClaimAdminFees() (*types.Transaction, error) {
	return _Curvefifactory.Contract.ClaimAdminFees(&_Curvefifactory.TransactOpts)
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_Curvefifactory *CurvefifactoryTransactorSession) ClaimAdminFees() (*types.Transaction, error) {
	return _Curvefifactory.Contract.ClaimAdminFees(&_Curvefifactory.TransactOpts)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_Curvefifactory *CurvefifactoryTransactor) CommitNewParameters(opts *bind.TransactOpts, _new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "commit_new_parameters", _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_Curvefifactory *CurvefifactorySession) CommitNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.CommitNewParameters(&_Curvefifactory.TransactOpts, _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_Curvefifactory *CurvefifactoryTransactorSession) CommitNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.CommitNewParameters(&_Curvefifactory.TransactOpts, _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Curvefifactory *CurvefifactorySession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.Exchange(&_Curvefifactory.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.Exchange(&_Curvefifactory.TransactOpts, i, j, dx, min_dy)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "exchange0", i, j, dx, min_dy, use_eth)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns(uint256)
func (_Curvefifactory *CurvefifactorySession) Exchange0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactory.Contract.Exchange0(&_Curvefifactory.TransactOpts, i, j, dx, min_dy, use_eth)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactorSession) Exchange0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactory.Contract.Exchange0(&_Curvefifactory.TransactOpts, i, j, dx, min_dy, use_eth)
}

// Exchange1 is a paid mutator transaction binding the contract method 0xce7d6503.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address receiver) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactor) Exchange1(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "exchange1", i, j, dx, min_dy, use_eth, receiver)
}

// Exchange1 is a paid mutator transaction binding the contract method 0xce7d6503.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address receiver) payable returns(uint256)
func (_Curvefifactory *CurvefifactorySession) Exchange1(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactory.Contract.Exchange1(&_Curvefifactory.TransactOpts, i, j, dx, min_dy, use_eth, receiver)
}

// Exchange1 is a paid mutator transaction binding the contract method 0xce7d6503.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address receiver) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactorSession) Exchange1(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactory.Contract.Exchange1(&_Curvefifactory.TransactOpts, i, j, dx, min_dy, use_eth, receiver)
}

// ExchangeExtended is a paid mutator transaction binding the contract method 0xdd96994f.
//
// Solidity: function exchange_extended(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address sender, address receiver, bytes32 cb) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactor) ExchangeExtended(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, sender common.Address, receiver common.Address, cb [32]byte) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "exchange_extended", i, j, dx, min_dy, use_eth, sender, receiver, cb)
}

// ExchangeExtended is a paid mutator transaction binding the contract method 0xdd96994f.
//
// Solidity: function exchange_extended(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address sender, address receiver, bytes32 cb) payable returns(uint256)
func (_Curvefifactory *CurvefifactorySession) ExchangeExtended(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, sender common.Address, receiver common.Address, cb [32]byte) (*types.Transaction, error) {
	return _Curvefifactory.Contract.ExchangeExtended(&_Curvefifactory.TransactOpts, i, j, dx, min_dy, use_eth, sender, receiver, cb)
}

// ExchangeExtended is a paid mutator transaction binding the contract method 0xdd96994f.
//
// Solidity: function exchange_extended(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address sender, address receiver, bytes32 cb) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactorSession) ExchangeExtended(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, sender common.Address, receiver common.Address, cb [32]byte) (*types.Transaction, error) {
	return _Curvefifactory.Contract.ExchangeExtended(&_Curvefifactory.TransactOpts, i, j, dx, min_dy, use_eth, sender, receiver, cb)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "exchange_underlying", i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Curvefifactory *CurvefifactorySession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.ExchangeUnderlying(&_Curvefifactory.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.ExchangeUnderlying(&_Curvefifactory.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying0 is a paid mutator transaction binding the contract method 0xe2ad025a.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy, address receiver) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactor) ExchangeUnderlying0(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "exchange_underlying0", i, j, dx, min_dy, receiver)
}

// ExchangeUnderlying0 is a paid mutator transaction binding the contract method 0xe2ad025a.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy, address receiver) payable returns(uint256)
func (_Curvefifactory *CurvefifactorySession) ExchangeUnderlying0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactory.Contract.ExchangeUnderlying0(&_Curvefifactory.TransactOpts, i, j, dx, min_dy, receiver)
}

// ExchangeUnderlying0 is a paid mutator transaction binding the contract method 0xe2ad025a.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy, address receiver) payable returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactorSession) ExchangeUnderlying0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactory.Contract.ExchangeUnderlying0(&_Curvefifactory.TransactOpts, i, j, dx, min_dy, receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0xa39e95c5.
//
// Solidity: function initialize(uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price, address _token, address[2] _coins, uint256 _precisions) returns()
func (_Curvefifactory *CurvefifactoryTransactor) Initialize(opts *bind.TransactOpts, A *big.Int, gamma *big.Int, mid_fee *big.Int, out_fee *big.Int, allowed_extra_profit *big.Int, fee_gamma *big.Int, adjustment_step *big.Int, admin_fee *big.Int, ma_half_time *big.Int, initial_price *big.Int, _token common.Address, _coins [2]common.Address, _precisions *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "initialize", A, gamma, mid_fee, out_fee, allowed_extra_profit, fee_gamma, adjustment_step, admin_fee, ma_half_time, initial_price, _token, _coins, _precisions)
}

// Initialize is a paid mutator transaction binding the contract method 0xa39e95c5.
//
// Solidity: function initialize(uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price, address _token, address[2] _coins, uint256 _precisions) returns()
func (_Curvefifactory *CurvefifactorySession) Initialize(A *big.Int, gamma *big.Int, mid_fee *big.Int, out_fee *big.Int, allowed_extra_profit *big.Int, fee_gamma *big.Int, adjustment_step *big.Int, admin_fee *big.Int, ma_half_time *big.Int, initial_price *big.Int, _token common.Address, _coins [2]common.Address, _precisions *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.Initialize(&_Curvefifactory.TransactOpts, A, gamma, mid_fee, out_fee, allowed_extra_profit, fee_gamma, adjustment_step, admin_fee, ma_half_time, initial_price, _token, _coins, _precisions)
}

// Initialize is a paid mutator transaction binding the contract method 0xa39e95c5.
//
// Solidity: function initialize(uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price, address _token, address[2] _coins, uint256 _precisions) returns()
func (_Curvefifactory *CurvefifactoryTransactorSession) Initialize(A *big.Int, gamma *big.Int, mid_fee *big.Int, out_fee *big.Int, allowed_extra_profit *big.Int, fee_gamma *big.Int, adjustment_step *big.Int, admin_fee *big.Int, ma_half_time *big.Int, initial_price *big.Int, _token common.Address, _coins [2]common.Address, _precisions *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.Initialize(&_Curvefifactory.TransactOpts, A, gamma, mid_fee, out_fee, allowed_extra_profit, fee_gamma, adjustment_step, admin_fee, ma_half_time, initial_price, _token, _coins, _precisions)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_Curvefifactory *CurvefifactoryTransactor) RampAGamma(opts *bind.TransactOpts, future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "ramp_A_gamma", future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_Curvefifactory *CurvefifactorySession) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.RampAGamma(&_Curvefifactory.TransactOpts, future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_Curvefifactory *CurvefifactoryTransactorSession) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.RampAGamma(&_Curvefifactory.TransactOpts, future_A, future_gamma, future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Curvefifactory *CurvefifactoryTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Curvefifactory *CurvefifactorySession) RemoveLiquidity(_amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.RemoveLiquidity(&_Curvefifactory.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Curvefifactory *CurvefifactoryTransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.RemoveLiquidity(&_Curvefifactory.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x269b5581.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, bool use_eth) returns()
func (_Curvefifactory *CurvefifactoryTransactor) RemoveLiquidity0(opts *bind.TransactOpts, _amount *big.Int, min_amounts [2]*big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "remove_liquidity0", _amount, min_amounts, use_eth)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x269b5581.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, bool use_eth) returns()
func (_Curvefifactory *CurvefifactorySession) RemoveLiquidity0(_amount *big.Int, min_amounts [2]*big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactory.Contract.RemoveLiquidity0(&_Curvefifactory.TransactOpts, _amount, min_amounts, use_eth)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x269b5581.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, bool use_eth) returns()
func (_Curvefifactory *CurvefifactoryTransactorSession) RemoveLiquidity0(_amount *big.Int, min_amounts [2]*big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactory.Contract.RemoveLiquidity0(&_Curvefifactory.TransactOpts, _amount, min_amounts, use_eth)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x1808e84a.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, bool use_eth, address receiver) returns()
func (_Curvefifactory *CurvefifactoryTransactor) RemoveLiquidity1(opts *bind.TransactOpts, _amount *big.Int, min_amounts [2]*big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "remove_liquidity1", _amount, min_amounts, use_eth, receiver)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x1808e84a.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, bool use_eth, address receiver) returns()
func (_Curvefifactory *CurvefifactorySession) RemoveLiquidity1(_amount *big.Int, min_amounts [2]*big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactory.Contract.RemoveLiquidity1(&_Curvefifactory.TransactOpts, _amount, min_amounts, use_eth, receiver)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x1808e84a.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, bool use_eth, address receiver) returns()
func (_Curvefifactory *CurvefifactoryTransactorSession) RemoveLiquidity1(_amount *big.Int, min_amounts [2]*big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactory.Contract.RemoveLiquidity1(&_Curvefifactory.TransactOpts, _amount, min_amounts, use_eth, receiver)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "remove_liquidity_one_coin", token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns(uint256)
func (_Curvefifactory *CurvefifactorySession) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.RemoveLiquidityOneCoin(&_Curvefifactory.TransactOpts, token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactorSession) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Curvefifactory.Contract.RemoveLiquidityOneCoin(&_Curvefifactory.TransactOpts, token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x8f15b6b5.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth) returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "remove_liquidity_one_coin0", token_amount, i, min_amount, use_eth)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x8f15b6b5.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth) returns(uint256)
func (_Curvefifactory *CurvefifactorySession) RemoveLiquidityOneCoin0(token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactory.Contract.RemoveLiquidityOneCoin0(&_Curvefifactory.TransactOpts, token_amount, i, min_amount, use_eth)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x8f15b6b5.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth) returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactorSession) RemoveLiquidityOneCoin0(token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Curvefifactory.Contract.RemoveLiquidityOneCoin0(&_Curvefifactory.TransactOpts, token_amount, i, min_amount, use_eth)
}

// RemoveLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0x07329bcd.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth, address receiver) returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactor) RemoveLiquidityOneCoin1(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "remove_liquidity_one_coin1", token_amount, i, min_amount, use_eth, receiver)
}

// RemoveLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0x07329bcd.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth, address receiver) returns(uint256)
func (_Curvefifactory *CurvefifactorySession) RemoveLiquidityOneCoin1(token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactory.Contract.RemoveLiquidityOneCoin1(&_Curvefifactory.TransactOpts, token_amount, i, min_amount, use_eth, receiver)
}

// RemoveLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0x07329bcd.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth, address receiver) returns(uint256)
func (_Curvefifactory *CurvefifactoryTransactorSession) RemoveLiquidityOneCoin1(token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _Curvefifactory.Contract.RemoveLiquidityOneCoin1(&_Curvefifactory.TransactOpts, token_amount, i, min_amount, use_eth, receiver)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Curvefifactory *CurvefifactoryTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Curvefifactory *CurvefifactorySession) RevertNewParameters() (*types.Transaction, error) {
	return _Curvefifactory.Contract.RevertNewParameters(&_Curvefifactory.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Curvefifactory *CurvefifactoryTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _Curvefifactory.Contract.RevertNewParameters(&_Curvefifactory.TransactOpts)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_Curvefifactory *CurvefifactoryTransactor) StopRampAGamma(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefifactory.contract.Transact(opts, "stop_ramp_A_gamma")
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_Curvefifactory *CurvefifactorySession) StopRampAGamma() (*types.Transaction, error) {
	return _Curvefifactory.Contract.StopRampAGamma(&_Curvefifactory.TransactOpts)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_Curvefifactory *CurvefifactoryTransactorSession) StopRampAGamma() (*types.Transaction, error) {
	return _Curvefifactory.Contract.StopRampAGamma(&_Curvefifactory.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Curvefifactory *CurvefifactoryTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Curvefifactory.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Curvefifactory *CurvefifactorySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Curvefifactory.Contract.Fallback(&_Curvefifactory.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Curvefifactory *CurvefifactoryTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Curvefifactory.Contract.Fallback(&_Curvefifactory.TransactOpts, calldata)
}

// CurvefifactoryAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Curvefifactory contract.
type CurvefifactoryAddLiquidityIterator struct {
	Event *CurvefifactoryAddLiquidity // Event containing the contract specifics and raw log

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
func (it *CurvefifactoryAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactoryAddLiquidity)
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
		it.Event = new(CurvefifactoryAddLiquidity)
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
func (it *CurvefifactoryAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactoryAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactoryAddLiquidity represents a AddLiquidity event raised by the Curvefifactory contract.
type CurvefifactoryAddLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fee          *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x540ab385f9b5d450a27404172caade516b3ba3f4be88239ac56a2ad1de2a1f5a.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256 fee, uint256 token_supply)
func (_Curvefifactory *CurvefifactoryFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurvefifactoryAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curvefifactory.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvefifactoryAddLiquidityIterator{contract: _Curvefifactory.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x540ab385f9b5d450a27404172caade516b3ba3f4be88239ac56a2ad1de2a1f5a.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256 fee, uint256 token_supply)
func (_Curvefifactory *CurvefifactoryFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *CurvefifactoryAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curvefifactory.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactoryAddLiquidity)
				if err := _Curvefifactory.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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
func (_Curvefifactory *CurvefifactoryFilterer) ParseAddLiquidity(log types.Log) (*CurvefifactoryAddLiquidity, error) {
	event := new(CurvefifactoryAddLiquidity)
	if err := _Curvefifactory.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactoryClaimAdminFeeIterator is returned from FilterClaimAdminFee and is used to iterate over the raw logs and unpacked data for ClaimAdminFee events raised by the Curvefifactory contract.
type CurvefifactoryClaimAdminFeeIterator struct {
	Event *CurvefifactoryClaimAdminFee // Event containing the contract specifics and raw log

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
func (it *CurvefifactoryClaimAdminFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactoryClaimAdminFee)
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
		it.Event = new(CurvefifactoryClaimAdminFee)
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
func (it *CurvefifactoryClaimAdminFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactoryClaimAdminFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactoryClaimAdminFee represents a ClaimAdminFee event raised by the Curvefifactory contract.
type CurvefifactoryClaimAdminFee struct {
	Admin  common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimAdminFee is a free log retrieval operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_Curvefifactory *CurvefifactoryFilterer) FilterClaimAdminFee(opts *bind.FilterOpts, admin []common.Address) (*CurvefifactoryClaimAdminFeeIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Curvefifactory.contract.FilterLogs(opts, "ClaimAdminFee", adminRule)
	if err != nil {
		return nil, err
	}
	return &CurvefifactoryClaimAdminFeeIterator{contract: _Curvefifactory.contract, event: "ClaimAdminFee", logs: logs, sub: sub}, nil
}

// WatchClaimAdminFee is a free log subscription operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_Curvefifactory *CurvefifactoryFilterer) WatchClaimAdminFee(opts *bind.WatchOpts, sink chan<- *CurvefifactoryClaimAdminFee, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Curvefifactory.contract.WatchLogs(opts, "ClaimAdminFee", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactoryClaimAdminFee)
				if err := _Curvefifactory.contract.UnpackLog(event, "ClaimAdminFee", log); err != nil {
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
func (_Curvefifactory *CurvefifactoryFilterer) ParseClaimAdminFee(log types.Log) (*CurvefifactoryClaimAdminFee, error) {
	event := new(CurvefifactoryClaimAdminFee)
	if err := _Curvefifactory.contract.UnpackLog(event, "ClaimAdminFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactoryCommitNewParametersIterator is returned from FilterCommitNewParameters and is used to iterate over the raw logs and unpacked data for CommitNewParameters events raised by the Curvefifactory contract.
type CurvefifactoryCommitNewParametersIterator struct {
	Event *CurvefifactoryCommitNewParameters // Event containing the contract specifics and raw log

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
func (it *CurvefifactoryCommitNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactoryCommitNewParameters)
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
		it.Event = new(CurvefifactoryCommitNewParameters)
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
func (it *CurvefifactoryCommitNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactoryCommitNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactoryCommitNewParameters represents a CommitNewParameters event raised by the Curvefifactory contract.
type CurvefifactoryCommitNewParameters struct {
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
func (_Curvefifactory *CurvefifactoryFilterer) FilterCommitNewParameters(opts *bind.FilterOpts, deadline []*big.Int) (*CurvefifactoryCommitNewParametersIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Curvefifactory.contract.FilterLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return &CurvefifactoryCommitNewParametersIterator{contract: _Curvefifactory.contract, event: "CommitNewParameters", logs: logs, sub: sub}, nil
}

// WatchCommitNewParameters is a free log subscription operation binding the contract event 0x913fde9a37e1f8ab67876a4d0ce80790d764fcfc5692f4529526df9c6bdde553.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Curvefifactory *CurvefifactoryFilterer) WatchCommitNewParameters(opts *bind.WatchOpts, sink chan<- *CurvefifactoryCommitNewParameters, deadline []*big.Int) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Curvefifactory.contract.WatchLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactoryCommitNewParameters)
				if err := _Curvefifactory.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
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
func (_Curvefifactory *CurvefifactoryFilterer) ParseCommitNewParameters(log types.Log) (*CurvefifactoryCommitNewParameters, error) {
	event := new(CurvefifactoryCommitNewParameters)
	if err := _Curvefifactory.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactoryNewParametersIterator is returned from FilterNewParameters and is used to iterate over the raw logs and unpacked data for NewParameters events raised by the Curvefifactory contract.
type CurvefifactoryNewParametersIterator struct {
	Event *CurvefifactoryNewParameters // Event containing the contract specifics and raw log

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
func (it *CurvefifactoryNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactoryNewParameters)
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
		it.Event = new(CurvefifactoryNewParameters)
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
func (it *CurvefifactoryNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactoryNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactoryNewParameters represents a NewParameters event raised by the Curvefifactory contract.
type CurvefifactoryNewParameters struct {
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
func (_Curvefifactory *CurvefifactoryFilterer) FilterNewParameters(opts *bind.FilterOpts) (*CurvefifactoryNewParametersIterator, error) {

	logs, sub, err := _Curvefifactory.contract.FilterLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return &CurvefifactoryNewParametersIterator{contract: _Curvefifactory.contract, event: "NewParameters", logs: logs, sub: sub}, nil
}

// WatchNewParameters is a free log subscription operation binding the contract event 0x1c65bbdc939f346e5d6f0bde1f072819947438d4fc7b182cc59c2f6dc5504087.
//
// Solidity: event NewParameters(uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Curvefifactory *CurvefifactoryFilterer) WatchNewParameters(opts *bind.WatchOpts, sink chan<- *CurvefifactoryNewParameters) (event.Subscription, error) {

	logs, sub, err := _Curvefifactory.contract.WatchLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactoryNewParameters)
				if err := _Curvefifactory.contract.UnpackLog(event, "NewParameters", log); err != nil {
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
func (_Curvefifactory *CurvefifactoryFilterer) ParseNewParameters(log types.Log) (*CurvefifactoryNewParameters, error) {
	event := new(CurvefifactoryNewParameters)
	if err := _Curvefifactory.contract.UnpackLog(event, "NewParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactoryRampAgammaIterator is returned from FilterRampAgamma and is used to iterate over the raw logs and unpacked data for RampAgamma events raised by the Curvefifactory contract.
type CurvefifactoryRampAgammaIterator struct {
	Event *CurvefifactoryRampAgamma // Event containing the contract specifics and raw log

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
func (it *CurvefifactoryRampAgammaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactoryRampAgamma)
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
		it.Event = new(CurvefifactoryRampAgamma)
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
func (it *CurvefifactoryRampAgammaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactoryRampAgammaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactoryRampAgamma represents a RampAgamma event raised by the Curvefifactory contract.
type CurvefifactoryRampAgamma struct {
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
func (_Curvefifactory *CurvefifactoryFilterer) FilterRampAgamma(opts *bind.FilterOpts) (*CurvefifactoryRampAgammaIterator, error) {

	logs, sub, err := _Curvefifactory.contract.FilterLogs(opts, "RampAgamma")
	if err != nil {
		return nil, err
	}
	return &CurvefifactoryRampAgammaIterator{contract: _Curvefifactory.contract, event: "RampAgamma", logs: logs, sub: sub}, nil
}

// WatchRampAgamma is a free log subscription operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_Curvefifactory *CurvefifactoryFilterer) WatchRampAgamma(opts *bind.WatchOpts, sink chan<- *CurvefifactoryRampAgamma) (event.Subscription, error) {

	logs, sub, err := _Curvefifactory.contract.WatchLogs(opts, "RampAgamma")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactoryRampAgamma)
				if err := _Curvefifactory.contract.UnpackLog(event, "RampAgamma", log); err != nil {
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
func (_Curvefifactory *CurvefifactoryFilterer) ParseRampAgamma(log types.Log) (*CurvefifactoryRampAgamma, error) {
	event := new(CurvefifactoryRampAgamma)
	if err := _Curvefifactory.contract.UnpackLog(event, "RampAgamma", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactoryRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Curvefifactory contract.
type CurvefifactoryRemoveLiquidityIterator struct {
	Event *CurvefifactoryRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *CurvefifactoryRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactoryRemoveLiquidity)
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
		it.Event = new(CurvefifactoryRemoveLiquidity)
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
func (it *CurvefifactoryRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactoryRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactoryRemoveLiquidity represents a RemoveLiquidity event raised by the Curvefifactory contract.
type CurvefifactoryRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xdd3c0336a16f1b64f172b7bb0dad5b2b3c7c76f91e8c4aafd6aae60dce800153.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256 token_supply)
func (_Curvefifactory *CurvefifactoryFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurvefifactoryRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curvefifactory.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvefifactoryRemoveLiquidityIterator{contract: _Curvefifactory.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xdd3c0336a16f1b64f172b7bb0dad5b2b3c7c76f91e8c4aafd6aae60dce800153.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256 token_supply)
func (_Curvefifactory *CurvefifactoryFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *CurvefifactoryRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curvefifactory.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactoryRemoveLiquidity)
				if err := _Curvefifactory.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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
func (_Curvefifactory *CurvefifactoryFilterer) ParseRemoveLiquidity(log types.Log) (*CurvefifactoryRemoveLiquidity, error) {
	event := new(CurvefifactoryRemoveLiquidity)
	if err := _Curvefifactory.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactoryRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the Curvefifactory contract.
type CurvefifactoryRemoveLiquidityOneIterator struct {
	Event *CurvefifactoryRemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *CurvefifactoryRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactoryRemoveLiquidityOne)
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
		it.Event = new(CurvefifactoryRemoveLiquidityOne)
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
func (it *CurvefifactoryRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactoryRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactoryRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the Curvefifactory contract.
type CurvefifactoryRemoveLiquidityOne struct {
	Provider    common.Address
	TokenAmount *big.Int
	CoinIndex   *big.Int
	CoinAmount  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount)
func (_Curvefifactory *CurvefifactoryFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*CurvefifactoryRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curvefifactory.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvefifactoryRemoveLiquidityOneIterator{contract: _Curvefifactory.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount)
func (_Curvefifactory *CurvefifactoryFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *CurvefifactoryRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Curvefifactory.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactoryRemoveLiquidityOne)
				if err := _Curvefifactory.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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
func (_Curvefifactory *CurvefifactoryFilterer) ParseRemoveLiquidityOne(log types.Log) (*CurvefifactoryRemoveLiquidityOne, error) {
	event := new(CurvefifactoryRemoveLiquidityOne)
	if err := _Curvefifactory.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactoryStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the Curvefifactory contract.
type CurvefifactoryStopRampAIterator struct {
	Event *CurvefifactoryStopRampA // Event containing the contract specifics and raw log

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
func (it *CurvefifactoryStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactoryStopRampA)
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
		it.Event = new(CurvefifactoryStopRampA)
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
func (it *CurvefifactoryStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactoryStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactoryStopRampA represents a StopRampA event raised by the Curvefifactory contract.
type CurvefifactoryStopRampA struct {
	CurrentA     *big.Int
	CurrentGamma *big.Int
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_Curvefifactory *CurvefifactoryFilterer) FilterStopRampA(opts *bind.FilterOpts) (*CurvefifactoryStopRampAIterator, error) {

	logs, sub, err := _Curvefifactory.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &CurvefifactoryStopRampAIterator{contract: _Curvefifactory.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_Curvefifactory *CurvefifactoryFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *CurvefifactoryStopRampA) (event.Subscription, error) {

	logs, sub, err := _Curvefifactory.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactoryStopRampA)
				if err := _Curvefifactory.contract.UnpackLog(event, "StopRampA", log); err != nil {
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
func (_Curvefifactory *CurvefifactoryFilterer) ParseStopRampA(log types.Log) (*CurvefifactoryStopRampA, error) {
	event := new(CurvefifactoryStopRampA)
	if err := _Curvefifactory.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefifactoryTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the Curvefifactory contract.
type CurvefifactoryTokenExchangeIterator struct {
	Event *CurvefifactoryTokenExchange // Event containing the contract specifics and raw log

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
func (it *CurvefifactoryTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefifactoryTokenExchange)
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
		it.Event = new(CurvefifactoryTokenExchange)
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
func (it *CurvefifactoryTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefifactoryTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefifactoryTokenExchange represents a TokenExchange event raised by the Curvefifactory contract.
type CurvefifactoryTokenExchange struct {
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
func (_Curvefifactory *CurvefifactoryFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*CurvefifactoryTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Curvefifactory.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &CurvefifactoryTokenExchangeIterator{contract: _Curvefifactory.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_Curvefifactory *CurvefifactoryFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *CurvefifactoryTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Curvefifactory.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefifactoryTokenExchange)
				if err := _Curvefifactory.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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
func (_Curvefifactory *CurvefifactoryFilterer) ParseTokenExchange(log types.Log) (*CurvefifactoryTokenExchange, error) {
	event := new(CurvefifactoryTokenExchange)
	if err := _Curvefifactory.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
