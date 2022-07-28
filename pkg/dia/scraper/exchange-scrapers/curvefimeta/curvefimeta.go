// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curvefimeta

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

// CurvefimetaMetaData contains all meta data concerning the Curvefimeta contract.
var CurvefimetaMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"BasePoolAdded\",\"inputs\":[{\"name\":\"base_pool\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"PlainPoolDeployed\",\"inputs\":[{\"name\":\"coins\",\"type\":\"address[4]\",\"indexed\":false},{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"MetaPoolDeployed\",\"inputs\":[{\"name\":\"coin\",\"type\":\"address\",\"indexed\":false},{\"name\":\"base_pool\",\"type\":\"address\",\"indexed\":false},{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"LiquidityGaugeDeployed\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false},{\"name\":\"gauge\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_fee_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"metapool_implementations\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[10]\"}],\"gas\":21716},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_base_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2663},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_n_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2699},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_meta_n_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":5201},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[4]\"}],\"gas\":9164},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}],\"gas\":21345},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"gas\":20185},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":19730},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_metapool_rates\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"gas\":5281},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"gas\":20435},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":39733},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_A\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3135},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_fees\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":5821},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_admin_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"gas\":13535},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_indices\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"bool\"}],\"gas\":33407},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3089},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_implementation_address\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3119},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_meta\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":3152},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_asset_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":5450},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_fee_receiver\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":5480},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_plain_pool\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[4]\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_plain_pool\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[4]\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_asset_type\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_plain_pool\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[4]\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_asset_type\",\"type\":\"uint256\"},{\"name\":\"_implementation_idx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_metapool\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coin\",\"type\":\"address\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_metapool\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coin\",\"type\":\"address\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_implementation_idx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":93079},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_base_pool\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_fee_receiver\",\"type\":\"address\"},{\"name\":\"_asset_type\",\"type\":\"uint256\"},{\"name\":\"_implementations\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":1206132},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_metapool_implementations\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_implementations\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":382072},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_plain_implementations\",\"inputs\":[{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_implementations\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":379687},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_gauge_implementation\",\"inputs\":[{\"name\":\"_gauge_implementation\",\"type\":\"address\"}],\"outputs\":[],\"gas\":38355},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"batch_set_pool_asset_type\",\"inputs\":[{\"name\":\"_pools\",\"type\":\"address[32]\"},{\"name\":\"_asset_types\",\"type\":\"uint256[32]\"}],\"outputs\":[],\"gas\":1139545},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_transfer_ownership\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"outputs\":[],\"gas\":38415},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"accept_transfer_ownership\",\"inputs\":[],\"outputs\":[],\"gas\":58366},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_manager\",\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\"}],\"outputs\":[],\"gas\":40996},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_fee_receiver\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_fee_receiver\",\"type\":\"address\"}],\"outputs\":[],\"gas\":38770},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"convert_metapool_fees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":12880},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_existing_metapools\",\"inputs\":[{\"name\":\"_pools\",\"type\":\"address[10]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":8610792},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3438},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3468},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3498},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_list\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3573},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3558},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"base_pool_list\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3633},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"base_pool_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3618},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"base_pool_assets\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":3863},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"plain_implementations\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"},{\"name\":\"arg1\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3838},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gauge_implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3708}]",
}

// CurvefimetaABI is the input ABI used to generate the binding from.
// Deprecated: Use CurvefimetaMetaData.ABI instead.
var CurvefimetaABI = CurvefimetaMetaData.ABI

// Curvefimeta is an auto generated Go binding around an Ethereum contract.
type Curvefimeta struct {
	CurvefimetaCaller     // Read-only binding to the contract
	CurvefimetaTransactor // Write-only binding to the contract
	CurvefimetaFilterer   // Log filterer for contract events
}

// CurvefimetaCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurvefimetaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvefimetaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurvefimetaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvefimetaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurvefimetaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvefimetaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurvefimetaSession struct {
	Contract     *Curvefimeta      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurvefimetaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurvefimetaCallerSession struct {
	Contract *CurvefimetaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CurvefimetaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurvefimetaTransactorSession struct {
	Contract     *CurvefimetaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CurvefimetaRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurvefimetaRaw struct {
	Contract *Curvefimeta // Generic contract binding to access the raw methods on
}

// CurvefimetaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurvefimetaCallerRaw struct {
	Contract *CurvefimetaCaller // Generic read-only contract binding to access the raw methods on
}

// CurvefimetaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurvefimetaTransactorRaw struct {
	Contract *CurvefimetaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurvefimeta creates a new instance of Curvefimeta, bound to a specific deployed contract.
func NewCurvefimeta(address common.Address, backend bind.ContractBackend) (*Curvefimeta, error) {
	contract, err := bindCurvefimeta(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Curvefimeta{CurvefimetaCaller: CurvefimetaCaller{contract: contract}, CurvefimetaTransactor: CurvefimetaTransactor{contract: contract}, CurvefimetaFilterer: CurvefimetaFilterer{contract: contract}}, nil
}

// NewCurvefimetaCaller creates a new read-only instance of Curvefimeta, bound to a specific deployed contract.
func NewCurvefimetaCaller(address common.Address, caller bind.ContractCaller) (*CurvefimetaCaller, error) {
	contract, err := bindCurvefimeta(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurvefimetaCaller{contract: contract}, nil
}

// NewCurvefimetaTransactor creates a new write-only instance of Curvefimeta, bound to a specific deployed contract.
func NewCurvefimetaTransactor(address common.Address, transactor bind.ContractTransactor) (*CurvefimetaTransactor, error) {
	contract, err := bindCurvefimeta(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurvefimetaTransactor{contract: contract}, nil
}

// NewCurvefimetaFilterer creates a new log filterer instance of Curvefimeta, bound to a specific deployed contract.
func NewCurvefimetaFilterer(address common.Address, filterer bind.ContractFilterer) (*CurvefimetaFilterer, error) {
	contract, err := bindCurvefimeta(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurvefimetaFilterer{contract: contract}, nil
}

// bindCurvefimeta binds a generic wrapper to an already deployed contract.
func bindCurvefimeta(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurvefimetaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curvefimeta *CurvefimetaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Curvefimeta.Contract.CurvefimetaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curvefimeta *CurvefimetaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefimeta.Contract.CurvefimetaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curvefimeta *CurvefimetaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curvefimeta.Contract.CurvefimetaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curvefimeta *CurvefimetaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Curvefimeta.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curvefimeta *CurvefimetaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefimeta.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curvefimeta *CurvefimetaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curvefimeta.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Curvefimeta *CurvefimetaCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Curvefimeta *CurvefimetaSession) Admin() (common.Address, error) {
	return _Curvefimeta.Contract.Admin(&_Curvefimeta.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Curvefimeta *CurvefimetaCallerSession) Admin() (common.Address, error) {
	return _Curvefimeta.Contract.Admin(&_Curvefimeta.CallOpts)
}

// BasePoolAssets is a free data retrieval call binding the contract method 0x10a002df.
//
// Solidity: function base_pool_assets(address arg0) view returns(bool)
func (_Curvefimeta *CurvefimetaCaller) BasePoolAssets(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "base_pool_assets", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BasePoolAssets is a free data retrieval call binding the contract method 0x10a002df.
//
// Solidity: function base_pool_assets(address arg0) view returns(bool)
func (_Curvefimeta *CurvefimetaSession) BasePoolAssets(arg0 common.Address) (bool, error) {
	return _Curvefimeta.Contract.BasePoolAssets(&_Curvefimeta.CallOpts, arg0)
}

// BasePoolAssets is a free data retrieval call binding the contract method 0x10a002df.
//
// Solidity: function base_pool_assets(address arg0) view returns(bool)
func (_Curvefimeta *CurvefimetaCallerSession) BasePoolAssets(arg0 common.Address) (bool, error) {
	return _Curvefimeta.Contract.BasePoolAssets(&_Curvefimeta.CallOpts, arg0)
}

// BasePoolCount is a free data retrieval call binding the contract method 0xde5e4a3b.
//
// Solidity: function base_pool_count() view returns(uint256)
func (_Curvefimeta *CurvefimetaCaller) BasePoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "base_pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasePoolCount is a free data retrieval call binding the contract method 0xde5e4a3b.
//
// Solidity: function base_pool_count() view returns(uint256)
func (_Curvefimeta *CurvefimetaSession) BasePoolCount() (*big.Int, error) {
	return _Curvefimeta.Contract.BasePoolCount(&_Curvefimeta.CallOpts)
}

// BasePoolCount is a free data retrieval call binding the contract method 0xde5e4a3b.
//
// Solidity: function base_pool_count() view returns(uint256)
func (_Curvefimeta *CurvefimetaCallerSession) BasePoolCount() (*big.Int, error) {
	return _Curvefimeta.Contract.BasePoolCount(&_Curvefimeta.CallOpts)
}

// BasePoolList is a free data retrieval call binding the contract method 0x22fe5671.
//
// Solidity: function base_pool_list(uint256 arg0) view returns(address)
func (_Curvefimeta *CurvefimetaCaller) BasePoolList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "base_pool_list", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BasePoolList is a free data retrieval call binding the contract method 0x22fe5671.
//
// Solidity: function base_pool_list(uint256 arg0) view returns(address)
func (_Curvefimeta *CurvefimetaSession) BasePoolList(arg0 *big.Int) (common.Address, error) {
	return _Curvefimeta.Contract.BasePoolList(&_Curvefimeta.CallOpts, arg0)
}

// BasePoolList is a free data retrieval call binding the contract method 0x22fe5671.
//
// Solidity: function base_pool_list(uint256 arg0) view returns(address)
func (_Curvefimeta *CurvefimetaCallerSession) BasePoolList(arg0 *big.Int) (common.Address, error) {
	return _Curvefimeta.Contract.BasePoolList(&_Curvefimeta.CallOpts, arg0)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Curvefimeta *CurvefimetaCaller) FindPoolForCoins(opts *bind.CallOpts, _from common.Address, _to common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "find_pool_for_coins", _from, _to)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Curvefimeta *CurvefimetaSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _Curvefimeta.Contract.FindPoolForCoins(&_Curvefimeta.CallOpts, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Curvefimeta *CurvefimetaCallerSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _Curvefimeta.Contract.FindPoolForCoins(&_Curvefimeta.CallOpts, _from, _to)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Curvefimeta *CurvefimetaCaller) FindPoolForCoins0(opts *bind.CallOpts, _from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "find_pool_for_coins0", _from, _to, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Curvefimeta *CurvefimetaSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _Curvefimeta.Contract.FindPoolForCoins0(&_Curvefimeta.CallOpts, _from, _to, i)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Curvefimeta *CurvefimetaCallerSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _Curvefimeta.Contract.FindPoolForCoins0(&_Curvefimeta.CallOpts, _from, _to, i)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Curvefimeta *CurvefimetaCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Curvefimeta *CurvefimetaSession) FutureAdmin() (common.Address, error) {
	return _Curvefimeta.Contract.FutureAdmin(&_Curvefimeta.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Curvefimeta *CurvefimetaCallerSession) FutureAdmin() (common.Address, error) {
	return _Curvefimeta.Contract.FutureAdmin(&_Curvefimeta.CallOpts)
}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_Curvefimeta *CurvefimetaCaller) GaugeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "gauge_implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_Curvefimeta *CurvefimetaSession) GaugeImplementation() (common.Address, error) {
	return _Curvefimeta.Contract.GaugeImplementation(&_Curvefimeta.CallOpts)
}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_Curvefimeta *CurvefimetaCallerSession) GaugeImplementation() (common.Address, error) {
	return _Curvefimeta.Contract.GaugeImplementation(&_Curvefimeta.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_Curvefimeta *CurvefimetaCaller) GetA(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_A", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_Curvefimeta *CurvefimetaSession) GetA(_pool common.Address) (*big.Int, error) {
	return _Curvefimeta.Contract.GetA(&_Curvefimeta.CallOpts, _pool)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_Curvefimeta *CurvefimetaCallerSession) GetA(_pool common.Address) (*big.Int, error) {
	return _Curvefimeta.Contract.GetA(&_Curvefimeta.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[4])
func (_Curvefimeta *CurvefimetaCaller) GetAdminBalances(opts *bind.CallOpts, _pool common.Address) ([4]*big.Int, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_admin_balances", _pool)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[4])
func (_Curvefimeta *CurvefimetaSession) GetAdminBalances(_pool common.Address) ([4]*big.Int, error) {
	return _Curvefimeta.Contract.GetAdminBalances(&_Curvefimeta.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[4])
func (_Curvefimeta *CurvefimetaCallerSession) GetAdminBalances(_pool common.Address) ([4]*big.Int, error) {
	return _Curvefimeta.Contract.GetAdminBalances(&_Curvefimeta.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[4])
func (_Curvefimeta *CurvefimetaCaller) GetBalances(opts *bind.CallOpts, _pool common.Address) ([4]*big.Int, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_balances", _pool)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[4])
func (_Curvefimeta *CurvefimetaSession) GetBalances(_pool common.Address) ([4]*big.Int, error) {
	return _Curvefimeta.Contract.GetBalances(&_Curvefimeta.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[4])
func (_Curvefimeta *CurvefimetaCallerSession) GetBalances(_pool common.Address) ([4]*big.Int, error) {
	return _Curvefimeta.Contract.GetBalances(&_Curvefimeta.CallOpts, _pool)
}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_Curvefimeta *CurvefimetaCaller) GetBasePool(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_base_pool", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_Curvefimeta *CurvefimetaSession) GetBasePool(_pool common.Address) (common.Address, error) {
	return _Curvefimeta.Contract.GetBasePool(&_Curvefimeta.CallOpts, _pool)
}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_Curvefimeta *CurvefimetaCallerSession) GetBasePool(_pool common.Address) (common.Address, error) {
	return _Curvefimeta.Contract.GetBasePool(&_Curvefimeta.CallOpts, _pool)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Curvefimeta *CurvefimetaCaller) GetCoinIndices(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_coin_indices", _pool, _from, _to)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Curvefimeta *CurvefimetaSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _Curvefimeta.Contract.GetCoinIndices(&_Curvefimeta.CallOpts, _pool, _from, _to)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Curvefimeta *CurvefimetaCallerSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _Curvefimeta.Contract.GetCoinIndices(&_Curvefimeta.CallOpts, _pool, _from, _to)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[4])
func (_Curvefimeta *CurvefimetaCaller) GetCoins(opts *bind.CallOpts, _pool common.Address) ([4]common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_coins", _pool)

	if err != nil {
		return *new([4]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([4]common.Address)).(*[4]common.Address)

	return out0, err

}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[4])
func (_Curvefimeta *CurvefimetaSession) GetCoins(_pool common.Address) ([4]common.Address, error) {
	return _Curvefimeta.Contract.GetCoins(&_Curvefimeta.CallOpts, _pool)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[4])
func (_Curvefimeta *CurvefimetaCallerSession) GetCoins(_pool common.Address) ([4]common.Address, error) {
	return _Curvefimeta.Contract.GetCoins(&_Curvefimeta.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[4])
func (_Curvefimeta *CurvefimetaCaller) GetDecimals(opts *bind.CallOpts, _pool common.Address) ([4]*big.Int, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_decimals", _pool)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[4])
func (_Curvefimeta *CurvefimetaSession) GetDecimals(_pool common.Address) ([4]*big.Int, error) {
	return _Curvefimeta.Contract.GetDecimals(&_Curvefimeta.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[4])
func (_Curvefimeta *CurvefimetaCallerSession) GetDecimals(_pool common.Address) ([4]*big.Int, error) {
	return _Curvefimeta.Contract.GetDecimals(&_Curvefimeta.CallOpts, _pool)
}

// GetFeeReceiver is a free data retrieval call binding the contract method 0x154aa8f5.
//
// Solidity: function get_fee_receiver(address _pool) view returns(address)
func (_Curvefimeta *CurvefimetaCaller) GetFeeReceiver(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_fee_receiver", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeReceiver is a free data retrieval call binding the contract method 0x154aa8f5.
//
// Solidity: function get_fee_receiver(address _pool) view returns(address)
func (_Curvefimeta *CurvefimetaSession) GetFeeReceiver(_pool common.Address) (common.Address, error) {
	return _Curvefimeta.Contract.GetFeeReceiver(&_Curvefimeta.CallOpts, _pool)
}

// GetFeeReceiver is a free data retrieval call binding the contract method 0x154aa8f5.
//
// Solidity: function get_fee_receiver(address _pool) view returns(address)
func (_Curvefimeta *CurvefimetaCallerSession) GetFeeReceiver(_pool common.Address) (common.Address, error) {
	return _Curvefimeta.Contract.GetFeeReceiver(&_Curvefimeta.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256, uint256)
func (_Curvefimeta *CurvefimetaCaller) GetFees(opts *bind.CallOpts, _pool common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_fees", _pool)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256, uint256)
func (_Curvefimeta *CurvefimetaSession) GetFees(_pool common.Address) (*big.Int, *big.Int, error) {
	return _Curvefimeta.Contract.GetFees(&_Curvefimeta.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256, uint256)
func (_Curvefimeta *CurvefimetaCallerSession) GetFees(_pool common.Address) (*big.Int, *big.Int, error) {
	return _Curvefimeta.Contract.GetFees(&_Curvefimeta.CallOpts, _pool)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_Curvefimeta *CurvefimetaCaller) GetGauge(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_gauge", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_Curvefimeta *CurvefimetaSession) GetGauge(_pool common.Address) (common.Address, error) {
	return _Curvefimeta.Contract.GetGauge(&_Curvefimeta.CallOpts, _pool)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_Curvefimeta *CurvefimetaCallerSession) GetGauge(_pool common.Address) (common.Address, error) {
	return _Curvefimeta.Contract.GetGauge(&_Curvefimeta.CallOpts, _pool)
}

// GetImplementationAddress is a free data retrieval call binding the contract method 0x510d98a4.
//
// Solidity: function get_implementation_address(address _pool) view returns(address)
func (_Curvefimeta *CurvefimetaCaller) GetImplementationAddress(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_implementation_address", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementationAddress is a free data retrieval call binding the contract method 0x510d98a4.
//
// Solidity: function get_implementation_address(address _pool) view returns(address)
func (_Curvefimeta *CurvefimetaSession) GetImplementationAddress(_pool common.Address) (common.Address, error) {
	return _Curvefimeta.Contract.GetImplementationAddress(&_Curvefimeta.CallOpts, _pool)
}

// GetImplementationAddress is a free data retrieval call binding the contract method 0x510d98a4.
//
// Solidity: function get_implementation_address(address _pool) view returns(address)
func (_Curvefimeta *CurvefimetaCallerSession) GetImplementationAddress(_pool common.Address) (common.Address, error) {
	return _Curvefimeta.Contract.GetImplementationAddress(&_Curvefimeta.CallOpts, _pool)
}

// GetMetaNCoins is a free data retrieval call binding the contract method 0xeb73f37d.
//
// Solidity: function get_meta_n_coins(address _pool) view returns(uint256, uint256)
func (_Curvefimeta *CurvefimetaCaller) GetMetaNCoins(opts *bind.CallOpts, _pool common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_meta_n_coins", _pool)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetMetaNCoins is a free data retrieval call binding the contract method 0xeb73f37d.
//
// Solidity: function get_meta_n_coins(address _pool) view returns(uint256, uint256)
func (_Curvefimeta *CurvefimetaSession) GetMetaNCoins(_pool common.Address) (*big.Int, *big.Int, error) {
	return _Curvefimeta.Contract.GetMetaNCoins(&_Curvefimeta.CallOpts, _pool)
}

// GetMetaNCoins is a free data retrieval call binding the contract method 0xeb73f37d.
//
// Solidity: function get_meta_n_coins(address _pool) view returns(uint256, uint256)
func (_Curvefimeta *CurvefimetaCallerSession) GetMetaNCoins(_pool common.Address) (*big.Int, *big.Int, error) {
	return _Curvefimeta.Contract.GetMetaNCoins(&_Curvefimeta.CallOpts, _pool)
}

// GetMetapoolRates is a free data retrieval call binding the contract method 0x06d8f160.
//
// Solidity: function get_metapool_rates(address _pool) view returns(uint256[2])
func (_Curvefimeta *CurvefimetaCaller) GetMetapoolRates(opts *bind.CallOpts, _pool common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_metapool_rates", _pool)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetMetapoolRates is a free data retrieval call binding the contract method 0x06d8f160.
//
// Solidity: function get_metapool_rates(address _pool) view returns(uint256[2])
func (_Curvefimeta *CurvefimetaSession) GetMetapoolRates(_pool common.Address) ([2]*big.Int, error) {
	return _Curvefimeta.Contract.GetMetapoolRates(&_Curvefimeta.CallOpts, _pool)
}

// GetMetapoolRates is a free data retrieval call binding the contract method 0x06d8f160.
//
// Solidity: function get_metapool_rates(address _pool) view returns(uint256[2])
func (_Curvefimeta *CurvefimetaCallerSession) GetMetapoolRates(_pool common.Address) ([2]*big.Int, error) {
	return _Curvefimeta.Contract.GetMetapoolRates(&_Curvefimeta.CallOpts, _pool)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_Curvefimeta *CurvefimetaCaller) GetNCoins(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_n_coins", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_Curvefimeta *CurvefimetaSession) GetNCoins(_pool common.Address) (*big.Int, error) {
	return _Curvefimeta.Contract.GetNCoins(&_Curvefimeta.CallOpts, _pool)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_Curvefimeta *CurvefimetaCallerSession) GetNCoins(_pool common.Address) (*big.Int, error) {
	return _Curvefimeta.Contract.GetNCoins(&_Curvefimeta.CallOpts, _pool)
}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_Curvefimeta *CurvefimetaCaller) GetPoolAssetType(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_pool_asset_type", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_Curvefimeta *CurvefimetaSession) GetPoolAssetType(_pool common.Address) (*big.Int, error) {
	return _Curvefimeta.Contract.GetPoolAssetType(&_Curvefimeta.CallOpts, _pool)
}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_Curvefimeta *CurvefimetaCallerSession) GetPoolAssetType(_pool common.Address) (*big.Int, error) {
	return _Curvefimeta.Contract.GetPoolAssetType(&_Curvefimeta.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Curvefimeta *CurvefimetaCaller) GetUnderlyingBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_underlying_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Curvefimeta *CurvefimetaSession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefimeta.Contract.GetUnderlyingBalances(&_Curvefimeta.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Curvefimeta *CurvefimetaCallerSession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefimeta.Contract.GetUnderlyingBalances(&_Curvefimeta.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Curvefimeta *CurvefimetaCaller) GetUnderlyingCoins(opts *bind.CallOpts, _pool common.Address) ([8]common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_underlying_coins", _pool)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Curvefimeta *CurvefimetaSession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _Curvefimeta.Contract.GetUnderlyingCoins(&_Curvefimeta.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Curvefimeta *CurvefimetaCallerSession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _Curvefimeta.Contract.GetUnderlyingCoins(&_Curvefimeta.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Curvefimeta *CurvefimetaCaller) GetUnderlyingDecimals(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "get_underlying_decimals", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Curvefimeta *CurvefimetaSession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefimeta.Contract.GetUnderlyingDecimals(&_Curvefimeta.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Curvefimeta *CurvefimetaCallerSession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefimeta.Contract.GetUnderlyingDecimals(&_Curvefimeta.CallOpts, _pool)
}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_Curvefimeta *CurvefimetaCaller) IsMeta(opts *bind.CallOpts, _pool common.Address) (bool, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "is_meta", _pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_Curvefimeta *CurvefimetaSession) IsMeta(_pool common.Address) (bool, error) {
	return _Curvefimeta.Contract.IsMeta(&_Curvefimeta.CallOpts, _pool)
}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_Curvefimeta *CurvefimetaCallerSession) IsMeta(_pool common.Address) (bool, error) {
	return _Curvefimeta.Contract.IsMeta(&_Curvefimeta.CallOpts, _pool)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Curvefimeta *CurvefimetaCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Curvefimeta *CurvefimetaSession) Manager() (common.Address, error) {
	return _Curvefimeta.Contract.Manager(&_Curvefimeta.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Curvefimeta *CurvefimetaCallerSession) Manager() (common.Address, error) {
	return _Curvefimeta.Contract.Manager(&_Curvefimeta.CallOpts)
}

// MetapoolImplementations is a free data retrieval call binding the contract method 0x970fa3f3.
//
// Solidity: function metapool_implementations(address _base_pool) view returns(address[10])
func (_Curvefimeta *CurvefimetaCaller) MetapoolImplementations(opts *bind.CallOpts, _base_pool common.Address) ([10]common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "metapool_implementations", _base_pool)

	if err != nil {
		return *new([10]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([10]common.Address)).(*[10]common.Address)

	return out0, err

}

// MetapoolImplementations is a free data retrieval call binding the contract method 0x970fa3f3.
//
// Solidity: function metapool_implementations(address _base_pool) view returns(address[10])
func (_Curvefimeta *CurvefimetaSession) MetapoolImplementations(_base_pool common.Address) ([10]common.Address, error) {
	return _Curvefimeta.Contract.MetapoolImplementations(&_Curvefimeta.CallOpts, _base_pool)
}

// MetapoolImplementations is a free data retrieval call binding the contract method 0x970fa3f3.
//
// Solidity: function metapool_implementations(address _base_pool) view returns(address[10])
func (_Curvefimeta *CurvefimetaCallerSession) MetapoolImplementations(_base_pool common.Address) ([10]common.Address, error) {
	return _Curvefimeta.Contract.MetapoolImplementations(&_Curvefimeta.CallOpts, _base_pool)
}

// PlainImplementations is a free data retrieval call binding the contract method 0x31a4f865.
//
// Solidity: function plain_implementations(uint256 arg0, uint256 arg1) view returns(address)
func (_Curvefimeta *CurvefimetaCaller) PlainImplementations(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "plain_implementations", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlainImplementations is a free data retrieval call binding the contract method 0x31a4f865.
//
// Solidity: function plain_implementations(uint256 arg0, uint256 arg1) view returns(address)
func (_Curvefimeta *CurvefimetaSession) PlainImplementations(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Curvefimeta.Contract.PlainImplementations(&_Curvefimeta.CallOpts, arg0, arg1)
}

// PlainImplementations is a free data retrieval call binding the contract method 0x31a4f865.
//
// Solidity: function plain_implementations(uint256 arg0, uint256 arg1) view returns(address)
func (_Curvefimeta *CurvefimetaCallerSession) PlainImplementations(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Curvefimeta.Contract.PlainImplementations(&_Curvefimeta.CallOpts, arg0, arg1)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Curvefimeta *CurvefimetaCaller) PoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Curvefimeta *CurvefimetaSession) PoolCount() (*big.Int, error) {
	return _Curvefimeta.Contract.PoolCount(&_Curvefimeta.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Curvefimeta *CurvefimetaCallerSession) PoolCount() (*big.Int, error) {
	return _Curvefimeta.Contract.PoolCount(&_Curvefimeta.CallOpts)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_Curvefimeta *CurvefimetaCaller) PoolList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curvefimeta.contract.Call(opts, &out, "pool_list", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_Curvefimeta *CurvefimetaSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _Curvefimeta.Contract.PoolList(&_Curvefimeta.CallOpts, arg0)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_Curvefimeta *CurvefimetaCallerSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _Curvefimeta.Contract.PoolList(&_Curvefimeta.CallOpts, arg0)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_Curvefimeta *CurvefimetaTransactor) AcceptTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "accept_transfer_ownership")
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_Curvefimeta *CurvefimetaSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _Curvefimeta.Contract.AcceptTransferOwnership(&_Curvefimeta.TransactOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_Curvefimeta *CurvefimetaTransactorSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _Curvefimeta.Contract.AcceptTransferOwnership(&_Curvefimeta.TransactOpts)
}

// AddBasePool is a paid mutator transaction binding the contract method 0x2fc05653.
//
// Solidity: function add_base_pool(address _base_pool, address _fee_receiver, uint256 _asset_type, address[10] _implementations) returns()
func (_Curvefimeta *CurvefimetaTransactor) AddBasePool(opts *bind.TransactOpts, _base_pool common.Address, _fee_receiver common.Address, _asset_type *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "add_base_pool", _base_pool, _fee_receiver, _asset_type, _implementations)
}

// AddBasePool is a paid mutator transaction binding the contract method 0x2fc05653.
//
// Solidity: function add_base_pool(address _base_pool, address _fee_receiver, uint256 _asset_type, address[10] _implementations) returns()
func (_Curvefimeta *CurvefimetaSession) AddBasePool(_base_pool common.Address, _fee_receiver common.Address, _asset_type *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.AddBasePool(&_Curvefimeta.TransactOpts, _base_pool, _fee_receiver, _asset_type, _implementations)
}

// AddBasePool is a paid mutator transaction binding the contract method 0x2fc05653.
//
// Solidity: function add_base_pool(address _base_pool, address _fee_receiver, uint256 _asset_type, address[10] _implementations) returns()
func (_Curvefimeta *CurvefimetaTransactorSession) AddBasePool(_base_pool common.Address, _fee_receiver common.Address, _asset_type *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.AddBasePool(&_Curvefimeta.TransactOpts, _base_pool, _fee_receiver, _asset_type, _implementations)
}

// AddExistingMetapools is a paid mutator transaction binding the contract method 0xdb89fabc.
//
// Solidity: function add_existing_metapools(address[10] _pools) returns(bool)
func (_Curvefimeta *CurvefimetaTransactor) AddExistingMetapools(opts *bind.TransactOpts, _pools [10]common.Address) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "add_existing_metapools", _pools)
}

// AddExistingMetapools is a paid mutator transaction binding the contract method 0xdb89fabc.
//
// Solidity: function add_existing_metapools(address[10] _pools) returns(bool)
func (_Curvefimeta *CurvefimetaSession) AddExistingMetapools(_pools [10]common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.AddExistingMetapools(&_Curvefimeta.TransactOpts, _pools)
}

// AddExistingMetapools is a paid mutator transaction binding the contract method 0xdb89fabc.
//
// Solidity: function add_existing_metapools(address[10] _pools) returns(bool)
func (_Curvefimeta *CurvefimetaTransactorSession) AddExistingMetapools(_pools [10]common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.AddExistingMetapools(&_Curvefimeta.TransactOpts, _pools)
}

// BatchSetPoolAssetType is a paid mutator transaction binding the contract method 0x7542f078.
//
// Solidity: function batch_set_pool_asset_type(address[32] _pools, uint256[32] _asset_types) returns()
func (_Curvefimeta *CurvefimetaTransactor) BatchSetPoolAssetType(opts *bind.TransactOpts, _pools [32]common.Address, _asset_types [32]*big.Int) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "batch_set_pool_asset_type", _pools, _asset_types)
}

// BatchSetPoolAssetType is a paid mutator transaction binding the contract method 0x7542f078.
//
// Solidity: function batch_set_pool_asset_type(address[32] _pools, uint256[32] _asset_types) returns()
func (_Curvefimeta *CurvefimetaSession) BatchSetPoolAssetType(_pools [32]common.Address, _asset_types [32]*big.Int) (*types.Transaction, error) {
	return _Curvefimeta.Contract.BatchSetPoolAssetType(&_Curvefimeta.TransactOpts, _pools, _asset_types)
}

// BatchSetPoolAssetType is a paid mutator transaction binding the contract method 0x7542f078.
//
// Solidity: function batch_set_pool_asset_type(address[32] _pools, uint256[32] _asset_types) returns()
func (_Curvefimeta *CurvefimetaTransactorSession) BatchSetPoolAssetType(_pools [32]common.Address, _asset_types [32]*big.Int) (*types.Transaction, error) {
	return _Curvefimeta.Contract.BatchSetPoolAssetType(&_Curvefimeta.TransactOpts, _pools, _asset_types)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_Curvefimeta *CurvefimetaTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "commit_transfer_ownership", _addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_Curvefimeta *CurvefimetaSession) CommitTransferOwnership(_addr common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.CommitTransferOwnership(&_Curvefimeta.TransactOpts, _addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_Curvefimeta *CurvefimetaTransactorSession) CommitTransferOwnership(_addr common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.CommitTransferOwnership(&_Curvefimeta.TransactOpts, _addr)
}

// ConvertMetapoolFees is a paid mutator transaction binding the contract method 0xbcc981d2.
//
// Solidity: function convert_metapool_fees() returns(bool)
func (_Curvefimeta *CurvefimetaTransactor) ConvertMetapoolFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "convert_metapool_fees")
}

// ConvertMetapoolFees is a paid mutator transaction binding the contract method 0xbcc981d2.
//
// Solidity: function convert_metapool_fees() returns(bool)
func (_Curvefimeta *CurvefimetaSession) ConvertMetapoolFees() (*types.Transaction, error) {
	return _Curvefimeta.Contract.ConvertMetapoolFees(&_Curvefimeta.TransactOpts)
}

// ConvertMetapoolFees is a paid mutator transaction binding the contract method 0xbcc981d2.
//
// Solidity: function convert_metapool_fees() returns(bool)
func (_Curvefimeta *CurvefimetaTransactorSession) ConvertMetapoolFees() (*types.Transaction, error) {
	return _Curvefimeta.Contract.ConvertMetapoolFees(&_Curvefimeta.TransactOpts)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_Curvefimeta *CurvefimetaTransactor) DeployGauge(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "deploy_gauge", _pool)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_Curvefimeta *CurvefimetaSession) DeployGauge(_pool common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.DeployGauge(&_Curvefimeta.TransactOpts, _pool)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_Curvefimeta *CurvefimetaTransactorSession) DeployGauge(_pool common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.DeployGauge(&_Curvefimeta.TransactOpts, _pool)
}

// DeployMetapool is a paid mutator transaction binding the contract method 0xe339eb4f.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee) returns(address)
func (_Curvefimeta *CurvefimetaTransactor) DeployMetapool(opts *bind.TransactOpts, _base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "deploy_metapool", _base_pool, _name, _symbol, _coin, _A, _fee)
}

// DeployMetapool is a paid mutator transaction binding the contract method 0xe339eb4f.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee) returns(address)
func (_Curvefimeta *CurvefimetaSession) DeployMetapool(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _Curvefimeta.Contract.DeployMetapool(&_Curvefimeta.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee)
}

// DeployMetapool is a paid mutator transaction binding the contract method 0xe339eb4f.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee) returns(address)
func (_Curvefimeta *CurvefimetaTransactorSession) DeployMetapool(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _Curvefimeta.Contract.DeployMetapool(&_Curvefimeta.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee)
}

// DeployMetapool0 is a paid mutator transaction binding the contract method 0xde7fe3bf.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee, uint256 _implementation_idx) returns(address)
func (_Curvefimeta *CurvefimetaTransactor) DeployMetapool0(opts *bind.TransactOpts, _base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "deploy_metapool0", _base_pool, _name, _symbol, _coin, _A, _fee, _implementation_idx)
}

// DeployMetapool0 is a paid mutator transaction binding the contract method 0xde7fe3bf.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee, uint256 _implementation_idx) returns(address)
func (_Curvefimeta *CurvefimetaSession) DeployMetapool0(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _Curvefimeta.Contract.DeployMetapool0(&_Curvefimeta.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee, _implementation_idx)
}

// DeployMetapool0 is a paid mutator transaction binding the contract method 0xde7fe3bf.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee, uint256 _implementation_idx) returns(address)
func (_Curvefimeta *CurvefimetaTransactorSession) DeployMetapool0(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _Curvefimeta.Contract.DeployMetapool0(&_Curvefimeta.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee, _implementation_idx)
}

// DeployPlainPool is a paid mutator transaction binding the contract method 0xcd419bb5.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee) returns(address)
func (_Curvefimeta *CurvefimetaTransactor) DeployPlainPool(opts *bind.TransactOpts, _name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "deploy_plain_pool", _name, _symbol, _coins, _A, _fee)
}

// DeployPlainPool is a paid mutator transaction binding the contract method 0xcd419bb5.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee) returns(address)
func (_Curvefimeta *CurvefimetaSession) DeployPlainPool(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _Curvefimeta.Contract.DeployPlainPool(&_Curvefimeta.TransactOpts, _name, _symbol, _coins, _A, _fee)
}

// DeployPlainPool is a paid mutator transaction binding the contract method 0xcd419bb5.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee) returns(address)
func (_Curvefimeta *CurvefimetaTransactorSession) DeployPlainPool(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _Curvefimeta.Contract.DeployPlainPool(&_Curvefimeta.TransactOpts, _name, _symbol, _coins, _A, _fee)
}

// DeployPlainPool0 is a paid mutator transaction binding the contract method 0x5c16487b.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type) returns(address)
func (_Curvefimeta *CurvefimetaTransactor) DeployPlainPool0(opts *bind.TransactOpts, _name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "deploy_plain_pool0", _name, _symbol, _coins, _A, _fee, _asset_type)
}

// DeployPlainPool0 is a paid mutator transaction binding the contract method 0x5c16487b.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type) returns(address)
func (_Curvefimeta *CurvefimetaSession) DeployPlainPool0(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int) (*types.Transaction, error) {
	return _Curvefimeta.Contract.DeployPlainPool0(&_Curvefimeta.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type)
}

// DeployPlainPool0 is a paid mutator transaction binding the contract method 0x5c16487b.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type) returns(address)
func (_Curvefimeta *CurvefimetaTransactorSession) DeployPlainPool0(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int) (*types.Transaction, error) {
	return _Curvefimeta.Contract.DeployPlainPool0(&_Curvefimeta.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type)
}

// DeployPlainPool1 is a paid mutator transaction binding the contract method 0x52f2db69.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type, uint256 _implementation_idx) returns(address)
func (_Curvefimeta *CurvefimetaTransactor) DeployPlainPool1(opts *bind.TransactOpts, _name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "deploy_plain_pool1", _name, _symbol, _coins, _A, _fee, _asset_type, _implementation_idx)
}

// DeployPlainPool1 is a paid mutator transaction binding the contract method 0x52f2db69.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type, uint256 _implementation_idx) returns(address)
func (_Curvefimeta *CurvefimetaSession) DeployPlainPool1(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _Curvefimeta.Contract.DeployPlainPool1(&_Curvefimeta.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type, _implementation_idx)
}

// DeployPlainPool1 is a paid mutator transaction binding the contract method 0x52f2db69.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type, uint256 _implementation_idx) returns(address)
func (_Curvefimeta *CurvefimetaTransactorSession) DeployPlainPool1(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _Curvefimeta.Contract.DeployPlainPool1(&_Curvefimeta.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type, _implementation_idx)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0x36d2b77a.
//
// Solidity: function set_fee_receiver(address _base_pool, address _fee_receiver) returns()
func (_Curvefimeta *CurvefimetaTransactor) SetFeeReceiver(opts *bind.TransactOpts, _base_pool common.Address, _fee_receiver common.Address) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "set_fee_receiver", _base_pool, _fee_receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0x36d2b77a.
//
// Solidity: function set_fee_receiver(address _base_pool, address _fee_receiver) returns()
func (_Curvefimeta *CurvefimetaSession) SetFeeReceiver(_base_pool common.Address, _fee_receiver common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.SetFeeReceiver(&_Curvefimeta.TransactOpts, _base_pool, _fee_receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0x36d2b77a.
//
// Solidity: function set_fee_receiver(address _base_pool, address _fee_receiver) returns()
func (_Curvefimeta *CurvefimetaTransactorSession) SetFeeReceiver(_base_pool common.Address, _fee_receiver common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.SetFeeReceiver(&_Curvefimeta.TransactOpts, _base_pool, _fee_receiver)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_Curvefimeta *CurvefimetaTransactor) SetGaugeImplementation(opts *bind.TransactOpts, _gauge_implementation common.Address) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "set_gauge_implementation", _gauge_implementation)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_Curvefimeta *CurvefimetaSession) SetGaugeImplementation(_gauge_implementation common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.SetGaugeImplementation(&_Curvefimeta.TransactOpts, _gauge_implementation)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_Curvefimeta *CurvefimetaTransactorSession) SetGaugeImplementation(_gauge_implementation common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.SetGaugeImplementation(&_Curvefimeta.TransactOpts, _gauge_implementation)
}

// SetManager is a paid mutator transaction binding the contract method 0x9aece83e.
//
// Solidity: function set_manager(address _manager) returns()
func (_Curvefimeta *CurvefimetaTransactor) SetManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "set_manager", _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0x9aece83e.
//
// Solidity: function set_manager(address _manager) returns()
func (_Curvefimeta *CurvefimetaSession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.SetManager(&_Curvefimeta.TransactOpts, _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0x9aece83e.
//
// Solidity: function set_manager(address _manager) returns()
func (_Curvefimeta *CurvefimetaTransactorSession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.SetManager(&_Curvefimeta.TransactOpts, _manager)
}

// SetMetapoolImplementations is a paid mutator transaction binding the contract method 0xcb956b46.
//
// Solidity: function set_metapool_implementations(address _base_pool, address[10] _implementations) returns()
func (_Curvefimeta *CurvefimetaTransactor) SetMetapoolImplementations(opts *bind.TransactOpts, _base_pool common.Address, _implementations [10]common.Address) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "set_metapool_implementations", _base_pool, _implementations)
}

// SetMetapoolImplementations is a paid mutator transaction binding the contract method 0xcb956b46.
//
// Solidity: function set_metapool_implementations(address _base_pool, address[10] _implementations) returns()
func (_Curvefimeta *CurvefimetaSession) SetMetapoolImplementations(_base_pool common.Address, _implementations [10]common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.SetMetapoolImplementations(&_Curvefimeta.TransactOpts, _base_pool, _implementations)
}

// SetMetapoolImplementations is a paid mutator transaction binding the contract method 0xcb956b46.
//
// Solidity: function set_metapool_implementations(address _base_pool, address[10] _implementations) returns()
func (_Curvefimeta *CurvefimetaTransactorSession) SetMetapoolImplementations(_base_pool common.Address, _implementations [10]common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.SetMetapoolImplementations(&_Curvefimeta.TransactOpts, _base_pool, _implementations)
}

// SetPlainImplementations is a paid mutator transaction binding the contract method 0x9ddbf4b9.
//
// Solidity: function set_plain_implementations(uint256 _n_coins, address[10] _implementations) returns()
func (_Curvefimeta *CurvefimetaTransactor) SetPlainImplementations(opts *bind.TransactOpts, _n_coins *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _Curvefimeta.contract.Transact(opts, "set_plain_implementations", _n_coins, _implementations)
}

// SetPlainImplementations is a paid mutator transaction binding the contract method 0x9ddbf4b9.
//
// Solidity: function set_plain_implementations(uint256 _n_coins, address[10] _implementations) returns()
func (_Curvefimeta *CurvefimetaSession) SetPlainImplementations(_n_coins *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.SetPlainImplementations(&_Curvefimeta.TransactOpts, _n_coins, _implementations)
}

// SetPlainImplementations is a paid mutator transaction binding the contract method 0x9ddbf4b9.
//
// Solidity: function set_plain_implementations(uint256 _n_coins, address[10] _implementations) returns()
func (_Curvefimeta *CurvefimetaTransactorSession) SetPlainImplementations(_n_coins *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _Curvefimeta.Contract.SetPlainImplementations(&_Curvefimeta.TransactOpts, _n_coins, _implementations)
}

// CurvefimetaBasePoolAddedIterator is returned from FilterBasePoolAdded and is used to iterate over the raw logs and unpacked data for BasePoolAdded events raised by the Curvefimeta contract.
type CurvefimetaBasePoolAddedIterator struct {
	Event *CurvefimetaBasePoolAdded // Event containing the contract specifics and raw log

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
func (it *CurvefimetaBasePoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefimetaBasePoolAdded)
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
		it.Event = new(CurvefimetaBasePoolAdded)
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
func (it *CurvefimetaBasePoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefimetaBasePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefimetaBasePoolAdded represents a BasePoolAdded event raised by the Curvefimeta contract.
type CurvefimetaBasePoolAdded struct {
	BasePool common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBasePoolAdded is a free log retrieval operation binding the contract event 0xcc6afdfec79da6be08142ecee25cf14b665961e25d30d8eba45959be9547635f.
//
// Solidity: event BasePoolAdded(address base_pool)
func (_Curvefimeta *CurvefimetaFilterer) FilterBasePoolAdded(opts *bind.FilterOpts) (*CurvefimetaBasePoolAddedIterator, error) {

	logs, sub, err := _Curvefimeta.contract.FilterLogs(opts, "BasePoolAdded")
	if err != nil {
		return nil, err
	}
	return &CurvefimetaBasePoolAddedIterator{contract: _Curvefimeta.contract, event: "BasePoolAdded", logs: logs, sub: sub}, nil
}

// WatchBasePoolAdded is a free log subscription operation binding the contract event 0xcc6afdfec79da6be08142ecee25cf14b665961e25d30d8eba45959be9547635f.
//
// Solidity: event BasePoolAdded(address base_pool)
func (_Curvefimeta *CurvefimetaFilterer) WatchBasePoolAdded(opts *bind.WatchOpts, sink chan<- *CurvefimetaBasePoolAdded) (event.Subscription, error) {

	logs, sub, err := _Curvefimeta.contract.WatchLogs(opts, "BasePoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefimetaBasePoolAdded)
				if err := _Curvefimeta.contract.UnpackLog(event, "BasePoolAdded", log); err != nil {
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

// ParseBasePoolAdded is a log parse operation binding the contract event 0xcc6afdfec79da6be08142ecee25cf14b665961e25d30d8eba45959be9547635f.
//
// Solidity: event BasePoolAdded(address base_pool)
func (_Curvefimeta *CurvefimetaFilterer) ParseBasePoolAdded(log types.Log) (*CurvefimetaBasePoolAdded, error) {
	event := new(CurvefimetaBasePoolAdded)
	if err := _Curvefimeta.contract.UnpackLog(event, "BasePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefimetaLiquidityGaugeDeployedIterator is returned from FilterLiquidityGaugeDeployed and is used to iterate over the raw logs and unpacked data for LiquidityGaugeDeployed events raised by the Curvefimeta contract.
type CurvefimetaLiquidityGaugeDeployedIterator struct {
	Event *CurvefimetaLiquidityGaugeDeployed // Event containing the contract specifics and raw log

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
func (it *CurvefimetaLiquidityGaugeDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefimetaLiquidityGaugeDeployed)
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
		it.Event = new(CurvefimetaLiquidityGaugeDeployed)
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
func (it *CurvefimetaLiquidityGaugeDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefimetaLiquidityGaugeDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefimetaLiquidityGaugeDeployed represents a LiquidityGaugeDeployed event raised by the Curvefimeta contract.
type CurvefimetaLiquidityGaugeDeployed struct {
	Pool  common.Address
	Gauge common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLiquidityGaugeDeployed is a free log retrieval operation binding the contract event 0x656bb34c20491970a8c163f3bd62ead82022b379c3924960ec60f6dbfc5aab3b.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address gauge)
func (_Curvefimeta *CurvefimetaFilterer) FilterLiquidityGaugeDeployed(opts *bind.FilterOpts) (*CurvefimetaLiquidityGaugeDeployedIterator, error) {

	logs, sub, err := _Curvefimeta.contract.FilterLogs(opts, "LiquidityGaugeDeployed")
	if err != nil {
		return nil, err
	}
	return &CurvefimetaLiquidityGaugeDeployedIterator{contract: _Curvefimeta.contract, event: "LiquidityGaugeDeployed", logs: logs, sub: sub}, nil
}

// WatchLiquidityGaugeDeployed is a free log subscription operation binding the contract event 0x656bb34c20491970a8c163f3bd62ead82022b379c3924960ec60f6dbfc5aab3b.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address gauge)
func (_Curvefimeta *CurvefimetaFilterer) WatchLiquidityGaugeDeployed(opts *bind.WatchOpts, sink chan<- *CurvefimetaLiquidityGaugeDeployed) (event.Subscription, error) {

	logs, sub, err := _Curvefimeta.contract.WatchLogs(opts, "LiquidityGaugeDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefimetaLiquidityGaugeDeployed)
				if err := _Curvefimeta.contract.UnpackLog(event, "LiquidityGaugeDeployed", log); err != nil {
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

// ParseLiquidityGaugeDeployed is a log parse operation binding the contract event 0x656bb34c20491970a8c163f3bd62ead82022b379c3924960ec60f6dbfc5aab3b.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address gauge)
func (_Curvefimeta *CurvefimetaFilterer) ParseLiquidityGaugeDeployed(log types.Log) (*CurvefimetaLiquidityGaugeDeployed, error) {
	event := new(CurvefimetaLiquidityGaugeDeployed)
	if err := _Curvefimeta.contract.UnpackLog(event, "LiquidityGaugeDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefimetaMetaPoolDeployedIterator is returned from FilterMetaPoolDeployed and is used to iterate over the raw logs and unpacked data for MetaPoolDeployed events raised by the Curvefimeta contract.
type CurvefimetaMetaPoolDeployedIterator struct {
	Event *CurvefimetaMetaPoolDeployed // Event containing the contract specifics and raw log

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
func (it *CurvefimetaMetaPoolDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefimetaMetaPoolDeployed)
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
		it.Event = new(CurvefimetaMetaPoolDeployed)
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
func (it *CurvefimetaMetaPoolDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefimetaMetaPoolDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefimetaMetaPoolDeployed represents a MetaPoolDeployed event raised by the Curvefimeta contract.
type CurvefimetaMetaPoolDeployed struct {
	Coin     common.Address
	BasePool common.Address
	A        *big.Int
	Fee      *big.Int
	Deployer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMetaPoolDeployed is a free log retrieval operation binding the contract event 0x01f31cd2abdeb4e5e10ba500f2db0f937d9e8c735ab04681925441b4ea37eda5.
//
// Solidity: event MetaPoolDeployed(address coin, address base_pool, uint256 A, uint256 fee, address deployer)
func (_Curvefimeta *CurvefimetaFilterer) FilterMetaPoolDeployed(opts *bind.FilterOpts) (*CurvefimetaMetaPoolDeployedIterator, error) {

	logs, sub, err := _Curvefimeta.contract.FilterLogs(opts, "MetaPoolDeployed")
	if err != nil {
		return nil, err
	}
	return &CurvefimetaMetaPoolDeployedIterator{contract: _Curvefimeta.contract, event: "MetaPoolDeployed", logs: logs, sub: sub}, nil
}

// WatchMetaPoolDeployed is a free log subscription operation binding the contract event 0x01f31cd2abdeb4e5e10ba500f2db0f937d9e8c735ab04681925441b4ea37eda5.
//
// Solidity: event MetaPoolDeployed(address coin, address base_pool, uint256 A, uint256 fee, address deployer)
func (_Curvefimeta *CurvefimetaFilterer) WatchMetaPoolDeployed(opts *bind.WatchOpts, sink chan<- *CurvefimetaMetaPoolDeployed) (event.Subscription, error) {

	logs, sub, err := _Curvefimeta.contract.WatchLogs(opts, "MetaPoolDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefimetaMetaPoolDeployed)
				if err := _Curvefimeta.contract.UnpackLog(event, "MetaPoolDeployed", log); err != nil {
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

// ParseMetaPoolDeployed is a log parse operation binding the contract event 0x01f31cd2abdeb4e5e10ba500f2db0f937d9e8c735ab04681925441b4ea37eda5.
//
// Solidity: event MetaPoolDeployed(address coin, address base_pool, uint256 A, uint256 fee, address deployer)
func (_Curvefimeta *CurvefimetaFilterer) ParseMetaPoolDeployed(log types.Log) (*CurvefimetaMetaPoolDeployed, error) {
	event := new(CurvefimetaMetaPoolDeployed)
	if err := _Curvefimeta.contract.UnpackLog(event, "MetaPoolDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefimetaPlainPoolDeployedIterator is returned from FilterPlainPoolDeployed and is used to iterate over the raw logs and unpacked data for PlainPoolDeployed events raised by the Curvefimeta contract.
type CurvefimetaPlainPoolDeployedIterator struct {
	Event *CurvefimetaPlainPoolDeployed // Event containing the contract specifics and raw log

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
func (it *CurvefimetaPlainPoolDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefimetaPlainPoolDeployed)
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
		it.Event = new(CurvefimetaPlainPoolDeployed)
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
func (it *CurvefimetaPlainPoolDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefimetaPlainPoolDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefimetaPlainPoolDeployed represents a PlainPoolDeployed event raised by the Curvefimeta contract.
type CurvefimetaPlainPoolDeployed struct {
	Coins    [4]common.Address
	A        *big.Int
	Fee      *big.Int
	Deployer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPlainPoolDeployed is a free log retrieval operation binding the contract event 0x5b4a28c940282b5bf183df6a046b8119cf6edeb62859f75e835eb7ba834cce8d.
//
// Solidity: event PlainPoolDeployed(address[4] coins, uint256 A, uint256 fee, address deployer)
func (_Curvefimeta *CurvefimetaFilterer) FilterPlainPoolDeployed(opts *bind.FilterOpts) (*CurvefimetaPlainPoolDeployedIterator, error) {

	logs, sub, err := _Curvefimeta.contract.FilterLogs(opts, "PlainPoolDeployed")
	if err != nil {
		return nil, err
	}
	return &CurvefimetaPlainPoolDeployedIterator{contract: _Curvefimeta.contract, event: "PlainPoolDeployed", logs: logs, sub: sub}, nil
}

// WatchPlainPoolDeployed is a free log subscription operation binding the contract event 0x5b4a28c940282b5bf183df6a046b8119cf6edeb62859f75e835eb7ba834cce8d.
//
// Solidity: event PlainPoolDeployed(address[4] coins, uint256 A, uint256 fee, address deployer)
func (_Curvefimeta *CurvefimetaFilterer) WatchPlainPoolDeployed(opts *bind.WatchOpts, sink chan<- *CurvefimetaPlainPoolDeployed) (event.Subscription, error) {

	logs, sub, err := _Curvefimeta.contract.WatchLogs(opts, "PlainPoolDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefimetaPlainPoolDeployed)
				if err := _Curvefimeta.contract.UnpackLog(event, "PlainPoolDeployed", log); err != nil {
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

// ParsePlainPoolDeployed is a log parse operation binding the contract event 0x5b4a28c940282b5bf183df6a046b8119cf6edeb62859f75e835eb7ba834cce8d.
//
// Solidity: event PlainPoolDeployed(address[4] coins, uint256 A, uint256 fee, address deployer)
func (_Curvefimeta *CurvefimetaFilterer) ParsePlainPoolDeployed(log types.Log) (*CurvefimetaPlainPoolDeployed, error) {
	event := new(CurvefimetaPlainPoolDeployed)
	if err := _Curvefimeta.contract.UnpackLog(event, "PlainPoolDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
