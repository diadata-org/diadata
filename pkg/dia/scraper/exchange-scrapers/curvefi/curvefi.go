// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curvefi

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

// CurvefiMetaData contains all meta data concerning the Curvefi contract.
var CurvefiMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"PoolAdded\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":true},{\"name\":\"rate_method_id\",\"type\":\"bytes\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"PoolRemoved\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_address_provider\",\"type\":\"address\"},{\"name\":\"_gauge_controller\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_n_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"gas\":1521},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}],\"gas\":12102},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}],\"gas\":12194},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":7874},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":7966},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_rates\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":36992},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauges\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[10]\"},{\"name\":\"\",\"type\":\"int128[10]\"}],\"gas\":20157},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":16583},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":162842},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price_from_lp_token\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":1927},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_A\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":1045},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_parameters\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"future_A\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"admin_fee\",\"type\":\"uint256\"},{\"name\":\"future_fee\",\"type\":\"uint256\"},{\"name\":\"future_admin_fee\",\"type\":\"uint256\"},{\"name\":\"future_owner\",\"type\":\"address\"},{\"name\":\"initial_A\",\"type\":\"uint256\"},{\"name\":\"initial_A_time\",\"type\":\"uint256\"},{\"name\":\"future_A_time\",\"type\":\"uint256\"}],\"gas\":6305},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_fees\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"gas\":1450},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_admin_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":36454},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_indices\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"bool\"}],\"gas\":27131},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"estimate_gas_used\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":32004},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_meta\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":1900},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_name\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":8323},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_swap_count\",\"inputs\":[{\"name\":\"_coin\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":1951},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_swap_complement\",\"inputs\":[{\"name\":\"_coin\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2090},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_asset_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2011},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_lp_token\",\"type\":\"address\"},{\"name\":\"_rate_info\",\"type\":\"bytes32\"},{\"name\":\"_decimals\",\"type\":\"uint256\"},{\"name\":\"_underlying_decimals\",\"type\":\"uint256\"},{\"name\":\"_has_initial_A\",\"type\":\"bool\"},{\"name\":\"_is_v1\",\"type\":\"bool\"},{\"name\":\"_name\",\"type\":\"string\"}],\"outputs\":[],\"gas\":61485845},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_pool_without_underlying\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_lp_token\",\"type\":\"address\"},{\"name\":\"_rate_info\",\"type\":\"bytes32\"},{\"name\":\"_decimals\",\"type\":\"uint256\"},{\"name\":\"_use_rates\",\"type\":\"uint256\"},{\"name\":\"_has_initial_A\",\"type\":\"bool\"},{\"name\":\"_is_v1\",\"type\":\"bool\"},{\"name\":\"_name\",\"type\":\"string\"}],\"outputs\":[],\"gas\":31306062},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_metapool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_lp_token\",\"type\":\"address\"},{\"name\":\"_decimals\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_metapool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_lp_token\",\"type\":\"address\"},{\"name\":\"_decimals\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_base_pool\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[],\"gas\":779731418758},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_pool_gas_estimates\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address[5]\"},{\"name\":\"_amount\",\"type\":\"uint256[2][5]\"}],\"outputs\":[],\"gas\":390460},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_coin_gas_estimates\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address[10]\"},{\"name\":\"_amount\",\"type\":\"uint256[10]\"}],\"outputs\":[],\"gas\":392047},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_gas_estimate_contract\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_estimator\",\"type\":\"address\"}],\"outputs\":[],\"gas\":72629},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_liquidity_gauges\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_liquidity_gauges\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":400675},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_pool_asset_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_asset_type\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":72667},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"batch_set_pool_asset_type\",\"inputs\":[{\"name\":\"_pools\",\"type\":\"address[32]\"},{\"name\":\"_asset_types\",\"type\":\"uint256[32]\"}],\"outputs\":[],\"gas\":1173447},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"address_provider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2048},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gauge_controller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2078},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_list\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2217},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2138},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coin_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2168},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2307},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_from_lp_token\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2443},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_lp_token\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2473},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_updated\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2288}]",
}

// CurvefiABI is the input ABI used to generate the binding from.
// Deprecated: Use CurvefiMetaData.ABI instead.
var CurvefiABI = CurvefiMetaData.ABI

// Curvefi is an auto generated Go binding around an Ethereum contract.
type Curvefi struct {
	CurvefiCaller     // Read-only binding to the contract
	CurvefiTransactor // Write-only binding to the contract
	CurvefiFilterer   // Log filterer for contract events
}

// CurvefiCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurvefiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvefiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurvefiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvefiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurvefiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvefiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurvefiSession struct {
	Contract     *Curvefi          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurvefiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurvefiCallerSession struct {
	Contract *CurvefiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CurvefiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurvefiTransactorSession struct {
	Contract     *CurvefiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CurvefiRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurvefiRaw struct {
	Contract *Curvefi // Generic contract binding to access the raw methods on
}

// CurvefiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurvefiCallerRaw struct {
	Contract *CurvefiCaller // Generic read-only contract binding to access the raw methods on
}

// CurvefiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurvefiTransactorRaw struct {
	Contract *CurvefiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurvefi creates a new instance of Curvefi, bound to a specific deployed contract.
func NewCurvefi(address common.Address, backend bind.ContractBackend) (*Curvefi, error) {
	contract, err := bindCurvefi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Curvefi{CurvefiCaller: CurvefiCaller{contract: contract}, CurvefiTransactor: CurvefiTransactor{contract: contract}, CurvefiFilterer: CurvefiFilterer{contract: contract}}, nil
}

// NewCurvefiCaller creates a new read-only instance of Curvefi, bound to a specific deployed contract.
func NewCurvefiCaller(address common.Address, caller bind.ContractCaller) (*CurvefiCaller, error) {
	contract, err := bindCurvefi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurvefiCaller{contract: contract}, nil
}

// NewCurvefiTransactor creates a new write-only instance of Curvefi, bound to a specific deployed contract.
func NewCurvefiTransactor(address common.Address, transactor bind.ContractTransactor) (*CurvefiTransactor, error) {
	contract, err := bindCurvefi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurvefiTransactor{contract: contract}, nil
}

// NewCurvefiFilterer creates a new log filterer instance of Curvefi, bound to a specific deployed contract.
func NewCurvefiFilterer(address common.Address, filterer bind.ContractFilterer) (*CurvefiFilterer, error) {
	contract, err := bindCurvefi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurvefiFilterer{contract: contract}, nil
}

// bindCurvefi binds a generic wrapper to an already deployed contract.
func bindCurvefi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurvefiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curvefi *CurvefiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Curvefi.Contract.CurvefiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curvefi *CurvefiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefi.Contract.CurvefiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curvefi *CurvefiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curvefi.Contract.CurvefiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curvefi *CurvefiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Curvefi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curvefi *CurvefiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curvefi *CurvefiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curvefi.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_Curvefi *CurvefiCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "address_provider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_Curvefi *CurvefiSession) AddressProvider() (common.Address, error) {
	return _Curvefi.Contract.AddressProvider(&_Curvefi.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_Curvefi *CurvefiCallerSession) AddressProvider() (common.Address, error) {
	return _Curvefi.Contract.AddressProvider(&_Curvefi.CallOpts)
}

// CoinCount is a free data retrieval call binding the contract method 0x5075770f.
//
// Solidity: function coin_count() view returns(uint256)
func (_Curvefi *CurvefiCaller) CoinCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "coin_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoinCount is a free data retrieval call binding the contract method 0x5075770f.
//
// Solidity: function coin_count() view returns(uint256)
func (_Curvefi *CurvefiSession) CoinCount() (*big.Int, error) {
	return _Curvefi.Contract.CoinCount(&_Curvefi.CallOpts)
}

// CoinCount is a free data retrieval call binding the contract method 0x5075770f.
//
// Solidity: function coin_count() view returns(uint256)
func (_Curvefi *CurvefiCallerSession) CoinCount() (*big.Int, error) {
	return _Curvefi.Contract.CoinCount(&_Curvefi.CallOpts)
}

// EstimateGasUsed is a free data retrieval call binding the contract method 0xb0bb365b.
//
// Solidity: function estimate_gas_used(address _pool, address _from, address _to) view returns(uint256)
func (_Curvefi *CurvefiCaller) EstimateGasUsed(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "estimate_gas_used", _pool, _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateGasUsed is a free data retrieval call binding the contract method 0xb0bb365b.
//
// Solidity: function estimate_gas_used(address _pool, address _from, address _to) view returns(uint256)
func (_Curvefi *CurvefiSession) EstimateGasUsed(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, error) {
	return _Curvefi.Contract.EstimateGasUsed(&_Curvefi.CallOpts, _pool, _from, _to)
}

// EstimateGasUsed is a free data retrieval call binding the contract method 0xb0bb365b.
//
// Solidity: function estimate_gas_used(address _pool, address _from, address _to) view returns(uint256)
func (_Curvefi *CurvefiCallerSession) EstimateGasUsed(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, error) {
	return _Curvefi.Contract.EstimateGasUsed(&_Curvefi.CallOpts, _pool, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Curvefi *CurvefiCaller) FindPoolForCoins(opts *bind.CallOpts, _from common.Address, _to common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "find_pool_for_coins", _from, _to)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Curvefi *CurvefiSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _Curvefi.Contract.FindPoolForCoins(&_Curvefi.CallOpts, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Curvefi *CurvefiCallerSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _Curvefi.Contract.FindPoolForCoins(&_Curvefi.CallOpts, _from, _to)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Curvefi *CurvefiCaller) FindPoolForCoins0(opts *bind.CallOpts, _from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "find_pool_for_coins0", _from, _to, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Curvefi *CurvefiSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _Curvefi.Contract.FindPoolForCoins0(&_Curvefi.CallOpts, _from, _to, i)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Curvefi *CurvefiCallerSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _Curvefi.Contract.FindPoolForCoins0(&_Curvefi.CallOpts, _from, _to, i)
}

// GaugeController is a free data retrieval call binding the contract method 0xd8b9a018.
//
// Solidity: function gauge_controller() view returns(address)
func (_Curvefi *CurvefiCaller) GaugeController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "gauge_controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeController is a free data retrieval call binding the contract method 0xd8b9a018.
//
// Solidity: function gauge_controller() view returns(address)
func (_Curvefi *CurvefiSession) GaugeController() (common.Address, error) {
	return _Curvefi.Contract.GaugeController(&_Curvefi.CallOpts)
}

// GaugeController is a free data retrieval call binding the contract method 0xd8b9a018.
//
// Solidity: function gauge_controller() view returns(address)
func (_Curvefi *CurvefiCallerSession) GaugeController() (common.Address, error) {
	return _Curvefi.Contract.GaugeController(&_Curvefi.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_Curvefi *CurvefiCaller) GetA(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_A", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_Curvefi *CurvefiSession) GetA(_pool common.Address) (*big.Int, error) {
	return _Curvefi.Contract.GetA(&_Curvefi.CallOpts, _pool)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_Curvefi *CurvefiCallerSession) GetA(_pool common.Address) (*big.Int, error) {
	return _Curvefi.Contract.GetA(&_Curvefi.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiCaller) GetAdminBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_admin_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiSession) GetAdminBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefi.Contract.GetAdminBalances(&_Curvefi.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiCallerSession) GetAdminBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefi.Contract.GetAdminBalances(&_Curvefi.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiCaller) GetBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiSession) GetBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefi.Contract.GetBalances(&_Curvefi.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiCallerSession) GetBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefi.Contract.GetBalances(&_Curvefi.CallOpts, _pool)
}

// GetCoin is a free data retrieval call binding the contract method 0x45f0db24.
//
// Solidity: function get_coin(uint256 arg0) view returns(address)
func (_Curvefi *CurvefiCaller) GetCoin(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_coin", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCoin is a free data retrieval call binding the contract method 0x45f0db24.
//
// Solidity: function get_coin(uint256 arg0) view returns(address)
func (_Curvefi *CurvefiSession) GetCoin(arg0 *big.Int) (common.Address, error) {
	return _Curvefi.Contract.GetCoin(&_Curvefi.CallOpts, arg0)
}

// GetCoin is a free data retrieval call binding the contract method 0x45f0db24.
//
// Solidity: function get_coin(uint256 arg0) view returns(address)
func (_Curvefi *CurvefiCallerSession) GetCoin(arg0 *big.Int) (common.Address, error) {
	return _Curvefi.Contract.GetCoin(&_Curvefi.CallOpts, arg0)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Curvefi *CurvefiCaller) GetCoinIndices(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_coin_indices", _pool, _from, _to)

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
func (_Curvefi *CurvefiSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _Curvefi.Contract.GetCoinIndices(&_Curvefi.CallOpts, _pool, _from, _to)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Curvefi *CurvefiCallerSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _Curvefi.Contract.GetCoinIndices(&_Curvefi.CallOpts, _pool, _from, _to)
}

// GetCoinSwapComplement is a free data retrieval call binding the contract method 0x5d211982.
//
// Solidity: function get_coin_swap_complement(address _coin, uint256 _index) view returns(address)
func (_Curvefi *CurvefiCaller) GetCoinSwapComplement(opts *bind.CallOpts, _coin common.Address, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_coin_swap_complement", _coin, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCoinSwapComplement is a free data retrieval call binding the contract method 0x5d211982.
//
// Solidity: function get_coin_swap_complement(address _coin, uint256 _index) view returns(address)
func (_Curvefi *CurvefiSession) GetCoinSwapComplement(_coin common.Address, _index *big.Int) (common.Address, error) {
	return _Curvefi.Contract.GetCoinSwapComplement(&_Curvefi.CallOpts, _coin, _index)
}

// GetCoinSwapComplement is a free data retrieval call binding the contract method 0x5d211982.
//
// Solidity: function get_coin_swap_complement(address _coin, uint256 _index) view returns(address)
func (_Curvefi *CurvefiCallerSession) GetCoinSwapComplement(_coin common.Address, _index *big.Int) (common.Address, error) {
	return _Curvefi.Contract.GetCoinSwapComplement(&_Curvefi.CallOpts, _coin, _index)
}

// GetCoinSwapCount is a free data retrieval call binding the contract method 0x98aede16.
//
// Solidity: function get_coin_swap_count(address _coin) view returns(uint256)
func (_Curvefi *CurvefiCaller) GetCoinSwapCount(opts *bind.CallOpts, _coin common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_coin_swap_count", _coin)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCoinSwapCount is a free data retrieval call binding the contract method 0x98aede16.
//
// Solidity: function get_coin_swap_count(address _coin) view returns(uint256)
func (_Curvefi *CurvefiSession) GetCoinSwapCount(_coin common.Address) (*big.Int, error) {
	return _Curvefi.Contract.GetCoinSwapCount(&_Curvefi.CallOpts, _coin)
}

// GetCoinSwapCount is a free data retrieval call binding the contract method 0x98aede16.
//
// Solidity: function get_coin_swap_count(address _coin) view returns(uint256)
func (_Curvefi *CurvefiCallerSession) GetCoinSwapCount(_coin common.Address) (*big.Int, error) {
	return _Curvefi.Contract.GetCoinSwapCount(&_Curvefi.CallOpts, _coin)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_Curvefi *CurvefiCaller) GetCoins(opts *bind.CallOpts, _pool common.Address) ([8]common.Address, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_coins", _pool)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_Curvefi *CurvefiSession) GetCoins(_pool common.Address) ([8]common.Address, error) {
	return _Curvefi.Contract.GetCoins(&_Curvefi.CallOpts, _pool)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_Curvefi *CurvefiCallerSession) GetCoins(_pool common.Address) ([8]common.Address, error) {
	return _Curvefi.Contract.GetCoins(&_Curvefi.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiCaller) GetDecimals(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_decimals", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiSession) GetDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefi.Contract.GetDecimals(&_Curvefi.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiCallerSession) GetDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefi.Contract.GetDecimals(&_Curvefi.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[2])
func (_Curvefi *CurvefiCaller) GetFees(opts *bind.CallOpts, _pool common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_fees", _pool)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[2])
func (_Curvefi *CurvefiSession) GetFees(_pool common.Address) ([2]*big.Int, error) {
	return _Curvefi.Contract.GetFees(&_Curvefi.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[2])
func (_Curvefi *CurvefiCallerSession) GetFees(_pool common.Address) ([2]*big.Int, error) {
	return _Curvefi.Contract.GetFees(&_Curvefi.CallOpts, _pool)
}

// GetGauges is a free data retrieval call binding the contract method 0x56059ffb.
//
// Solidity: function get_gauges(address _pool) view returns(address[10], int128[10])
func (_Curvefi *CurvefiCaller) GetGauges(opts *bind.CallOpts, _pool common.Address) ([10]common.Address, [10]*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_gauges", _pool)

	if err != nil {
		return *new([10]common.Address), *new([10]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([10]common.Address)).(*[10]common.Address)
	out1 := *abi.ConvertType(out[1], new([10]*big.Int)).(*[10]*big.Int)

	return out0, out1, err

}

// GetGauges is a free data retrieval call binding the contract method 0x56059ffb.
//
// Solidity: function get_gauges(address _pool) view returns(address[10], int128[10])
func (_Curvefi *CurvefiSession) GetGauges(_pool common.Address) ([10]common.Address, [10]*big.Int, error) {
	return _Curvefi.Contract.GetGauges(&_Curvefi.CallOpts, _pool)
}

// GetGauges is a free data retrieval call binding the contract method 0x56059ffb.
//
// Solidity: function get_gauges(address _pool) view returns(address[10], int128[10])
func (_Curvefi *CurvefiCallerSession) GetGauges(_pool common.Address) ([10]common.Address, [10]*big.Int, error) {
	return _Curvefi.Contract.GetGauges(&_Curvefi.CallOpts, _pool)
}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address arg0) view returns(address)
func (_Curvefi *CurvefiCaller) GetLpToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_lp_token", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address arg0) view returns(address)
func (_Curvefi *CurvefiSession) GetLpToken(arg0 common.Address) (common.Address, error) {
	return _Curvefi.Contract.GetLpToken(&_Curvefi.CallOpts, arg0)
}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address arg0) view returns(address)
func (_Curvefi *CurvefiCallerSession) GetLpToken(arg0 common.Address) (common.Address, error) {
	return _Curvefi.Contract.GetLpToken(&_Curvefi.CallOpts, arg0)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256[2])
func (_Curvefi *CurvefiCaller) GetNCoins(opts *bind.CallOpts, _pool common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_n_coins", _pool)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256[2])
func (_Curvefi *CurvefiSession) GetNCoins(_pool common.Address) ([2]*big.Int, error) {
	return _Curvefi.Contract.GetNCoins(&_Curvefi.CallOpts, _pool)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256[2])
func (_Curvefi *CurvefiCallerSession) GetNCoins(_pool common.Address) ([2]*big.Int, error) {
	return _Curvefi.Contract.GetNCoins(&_Curvefi.CallOpts, _pool)
}

// GetParameters is a free data retrieval call binding the contract method 0x1f80a957.
//
// Solidity: function get_parameters(address _pool) view returns(uint256 A, uint256 future_A, uint256 fee, uint256 admin_fee, uint256 future_fee, uint256 future_admin_fee, address future_owner, uint256 initial_A, uint256 initial_A_time, uint256 future_A_time)
func (_Curvefi *CurvefiCaller) GetParameters(opts *bind.CallOpts, _pool common.Address) (struct {
	A              *big.Int
	FutureA        *big.Int
	Fee            *big.Int
	AdminFee       *big.Int
	FutureFee      *big.Int
	FutureAdminFee *big.Int
	FutureOwner    common.Address
	InitialA       *big.Int
	InitialATime   *big.Int
	FutureATime    *big.Int
}, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_parameters", _pool)

	outstruct := new(struct {
		A              *big.Int
		FutureA        *big.Int
		Fee            *big.Int
		AdminFee       *big.Int
		FutureFee      *big.Int
		FutureAdminFee *big.Int
		FutureOwner    common.Address
		InitialA       *big.Int
		InitialATime   *big.Int
		FutureATime    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.A = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FutureA = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AdminFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FutureFee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.FutureAdminFee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.FutureOwner = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.InitialA = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.InitialATime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.FutureATime = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetParameters is a free data retrieval call binding the contract method 0x1f80a957.
//
// Solidity: function get_parameters(address _pool) view returns(uint256 A, uint256 future_A, uint256 fee, uint256 admin_fee, uint256 future_fee, uint256 future_admin_fee, address future_owner, uint256 initial_A, uint256 initial_A_time, uint256 future_A_time)
func (_Curvefi *CurvefiSession) GetParameters(_pool common.Address) (struct {
	A              *big.Int
	FutureA        *big.Int
	Fee            *big.Int
	AdminFee       *big.Int
	FutureFee      *big.Int
	FutureAdminFee *big.Int
	FutureOwner    common.Address
	InitialA       *big.Int
	InitialATime   *big.Int
	FutureATime    *big.Int
}, error) {
	return _Curvefi.Contract.GetParameters(&_Curvefi.CallOpts, _pool)
}

// GetParameters is a free data retrieval call binding the contract method 0x1f80a957.
//
// Solidity: function get_parameters(address _pool) view returns(uint256 A, uint256 future_A, uint256 fee, uint256 admin_fee, uint256 future_fee, uint256 future_admin_fee, address future_owner, uint256 initial_A, uint256 initial_A_time, uint256 future_A_time)
func (_Curvefi *CurvefiCallerSession) GetParameters(_pool common.Address) (struct {
	A              *big.Int
	FutureA        *big.Int
	Fee            *big.Int
	AdminFee       *big.Int
	FutureFee      *big.Int
	FutureAdminFee *big.Int
	FutureOwner    common.Address
	InitialA       *big.Int
	InitialATime   *big.Int
	FutureATime    *big.Int
}, error) {
	return _Curvefi.Contract.GetParameters(&_Curvefi.CallOpts, _pool)
}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_Curvefi *CurvefiCaller) GetPoolAssetType(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_pool_asset_type", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_Curvefi *CurvefiSession) GetPoolAssetType(_pool common.Address) (*big.Int, error) {
	return _Curvefi.Contract.GetPoolAssetType(&_Curvefi.CallOpts, _pool)
}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_Curvefi *CurvefiCallerSession) GetPoolAssetType(_pool common.Address) (*big.Int, error) {
	return _Curvefi.Contract.GetPoolAssetType(&_Curvefi.CallOpts, _pool)
}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address arg0) view returns(address)
func (_Curvefi *CurvefiCaller) GetPoolFromLpToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_pool_from_lp_token", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address arg0) view returns(address)
func (_Curvefi *CurvefiSession) GetPoolFromLpToken(arg0 common.Address) (common.Address, error) {
	return _Curvefi.Contract.GetPoolFromLpToken(&_Curvefi.CallOpts, arg0)
}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address arg0) view returns(address)
func (_Curvefi *CurvefiCallerSession) GetPoolFromLpToken(arg0 common.Address) (common.Address, error) {
	return _Curvefi.Contract.GetPoolFromLpToken(&_Curvefi.CallOpts, arg0)
}

// GetPoolName is a free data retrieval call binding the contract method 0x5c911741.
//
// Solidity: function get_pool_name(address _pool) view returns(string)
func (_Curvefi *CurvefiCaller) GetPoolName(opts *bind.CallOpts, _pool common.Address) (string, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_pool_name", _pool)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPoolName is a free data retrieval call binding the contract method 0x5c911741.
//
// Solidity: function get_pool_name(address _pool) view returns(string)
func (_Curvefi *CurvefiSession) GetPoolName(_pool common.Address) (string, error) {
	return _Curvefi.Contract.GetPoolName(&_Curvefi.CallOpts, _pool)
}

// GetPoolName is a free data retrieval call binding the contract method 0x5c911741.
//
// Solidity: function get_pool_name(address _pool) view returns(string)
func (_Curvefi *CurvefiCallerSession) GetPoolName(_pool common.Address) (string, error) {
	return _Curvefi.Contract.GetPoolName(&_Curvefi.CallOpts, _pool)
}

// GetRates is a free data retrieval call binding the contract method 0xce99e45a.
//
// Solidity: function get_rates(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiCaller) GetRates(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_rates", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetRates is a free data retrieval call binding the contract method 0xce99e45a.
//
// Solidity: function get_rates(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiSession) GetRates(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefi.Contract.GetRates(&_Curvefi.CallOpts, _pool)
}

// GetRates is a free data retrieval call binding the contract method 0xce99e45a.
//
// Solidity: function get_rates(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiCallerSession) GetRates(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefi.Contract.GetRates(&_Curvefi.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiCaller) GetUnderlyingBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_underlying_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiSession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefi.Contract.GetUnderlyingBalances(&_Curvefi.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiCallerSession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefi.Contract.GetUnderlyingBalances(&_Curvefi.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Curvefi *CurvefiCaller) GetUnderlyingCoins(opts *bind.CallOpts, _pool common.Address) ([8]common.Address, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_underlying_coins", _pool)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Curvefi *CurvefiSession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _Curvefi.Contract.GetUnderlyingCoins(&_Curvefi.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Curvefi *CurvefiCallerSession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _Curvefi.Contract.GetUnderlyingCoins(&_Curvefi.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiCaller) GetUnderlyingDecimals(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_underlying_decimals", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiSession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefi.Contract.GetUnderlyingDecimals(&_Curvefi.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Curvefi *CurvefiCallerSession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Curvefi.Contract.GetUnderlyingDecimals(&_Curvefi.CallOpts, _pool)
}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_Curvefi *CurvefiCaller) GetVirtualPriceFromLpToken(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "get_virtual_price_from_lp_token", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_Curvefi *CurvefiSession) GetVirtualPriceFromLpToken(_token common.Address) (*big.Int, error) {
	return _Curvefi.Contract.GetVirtualPriceFromLpToken(&_Curvefi.CallOpts, _token)
}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_Curvefi *CurvefiCallerSession) GetVirtualPriceFromLpToken(_token common.Address) (*big.Int, error) {
	return _Curvefi.Contract.GetVirtualPriceFromLpToken(&_Curvefi.CallOpts, _token)
}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_Curvefi *CurvefiCaller) IsMeta(opts *bind.CallOpts, _pool common.Address) (bool, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "is_meta", _pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_Curvefi *CurvefiSession) IsMeta(_pool common.Address) (bool, error) {
	return _Curvefi.Contract.IsMeta(&_Curvefi.CallOpts, _pool)
}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_Curvefi *CurvefiCallerSession) IsMeta(_pool common.Address) (bool, error) {
	return _Curvefi.Contract.IsMeta(&_Curvefi.CallOpts, _pool)
}

// LastUpdated is a free data retrieval call binding the contract method 0x68900961.
//
// Solidity: function last_updated() view returns(uint256)
func (_Curvefi *CurvefiCaller) LastUpdated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "last_updated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdated is a free data retrieval call binding the contract method 0x68900961.
//
// Solidity: function last_updated() view returns(uint256)
func (_Curvefi *CurvefiSession) LastUpdated() (*big.Int, error) {
	return _Curvefi.Contract.LastUpdated(&_Curvefi.CallOpts)
}

// LastUpdated is a free data retrieval call binding the contract method 0x68900961.
//
// Solidity: function last_updated() view returns(uint256)
func (_Curvefi *CurvefiCallerSession) LastUpdated() (*big.Int, error) {
	return _Curvefi.Contract.LastUpdated(&_Curvefi.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Curvefi *CurvefiCaller) PoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Curvefi *CurvefiSession) PoolCount() (*big.Int, error) {
	return _Curvefi.Contract.PoolCount(&_Curvefi.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Curvefi *CurvefiCallerSession) PoolCount() (*big.Int, error) {
	return _Curvefi.Contract.PoolCount(&_Curvefi.CallOpts)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_Curvefi *CurvefiCaller) PoolList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curvefi.contract.Call(opts, &out, "pool_list", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_Curvefi *CurvefiSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _Curvefi.Contract.PoolList(&_Curvefi.CallOpts, arg0)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_Curvefi *CurvefiCallerSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _Curvefi.Contract.PoolList(&_Curvefi.CallOpts, arg0)
}

// AddMetapool is a paid mutator transaction binding the contract method 0xce6f94e1.
//
// Solidity: function add_metapool(address _pool, uint256 _n_coins, address _lp_token, uint256 _decimals, string _name) returns()
func (_Curvefi *CurvefiTransactor) AddMetapool(opts *bind.TransactOpts, _pool common.Address, _n_coins *big.Int, _lp_token common.Address, _decimals *big.Int, _name string) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "add_metapool", _pool, _n_coins, _lp_token, _decimals, _name)
}

// AddMetapool is a paid mutator transaction binding the contract method 0xce6f94e1.
//
// Solidity: function add_metapool(address _pool, uint256 _n_coins, address _lp_token, uint256 _decimals, string _name) returns()
func (_Curvefi *CurvefiSession) AddMetapool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _decimals *big.Int, _name string) (*types.Transaction, error) {
	return _Curvefi.Contract.AddMetapool(&_Curvefi.TransactOpts, _pool, _n_coins, _lp_token, _decimals, _name)
}

// AddMetapool is a paid mutator transaction binding the contract method 0xce6f94e1.
//
// Solidity: function add_metapool(address _pool, uint256 _n_coins, address _lp_token, uint256 _decimals, string _name) returns()
func (_Curvefi *CurvefiTransactorSession) AddMetapool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _decimals *big.Int, _name string) (*types.Transaction, error) {
	return _Curvefi.Contract.AddMetapool(&_Curvefi.TransactOpts, _pool, _n_coins, _lp_token, _decimals, _name)
}

// AddMetapool0 is a paid mutator transaction binding the contract method 0xd2a06baf.
//
// Solidity: function add_metapool(address _pool, uint256 _n_coins, address _lp_token, uint256 _decimals, string _name, address _base_pool) returns()
func (_Curvefi *CurvefiTransactor) AddMetapool0(opts *bind.TransactOpts, _pool common.Address, _n_coins *big.Int, _lp_token common.Address, _decimals *big.Int, _name string, _base_pool common.Address) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "add_metapool0", _pool, _n_coins, _lp_token, _decimals, _name, _base_pool)
}

// AddMetapool0 is a paid mutator transaction binding the contract method 0xd2a06baf.
//
// Solidity: function add_metapool(address _pool, uint256 _n_coins, address _lp_token, uint256 _decimals, string _name, address _base_pool) returns()
func (_Curvefi *CurvefiSession) AddMetapool0(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _decimals *big.Int, _name string, _base_pool common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.AddMetapool0(&_Curvefi.TransactOpts, _pool, _n_coins, _lp_token, _decimals, _name, _base_pool)
}

// AddMetapool0 is a paid mutator transaction binding the contract method 0xd2a06baf.
//
// Solidity: function add_metapool(address _pool, uint256 _n_coins, address _lp_token, uint256 _decimals, string _name, address _base_pool) returns()
func (_Curvefi *CurvefiTransactorSession) AddMetapool0(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _decimals *big.Int, _name string, _base_pool common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.AddMetapool0(&_Curvefi.TransactOpts, _pool, _n_coins, _lp_token, _decimals, _name, _base_pool)
}

// AddPool is a paid mutator transaction binding the contract method 0x99209aa1.
//
// Solidity: function add_pool(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_info, uint256 _decimals, uint256 _underlying_decimals, bool _has_initial_A, bool _is_v1, string _name) returns()
func (_Curvefi *CurvefiTransactor) AddPool(opts *bind.TransactOpts, _pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_info [32]byte, _decimals *big.Int, _underlying_decimals *big.Int, _has_initial_A bool, _is_v1 bool, _name string) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "add_pool", _pool, _n_coins, _lp_token, _rate_info, _decimals, _underlying_decimals, _has_initial_A, _is_v1, _name)
}

// AddPool is a paid mutator transaction binding the contract method 0x99209aa1.
//
// Solidity: function add_pool(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_info, uint256 _decimals, uint256 _underlying_decimals, bool _has_initial_A, bool _is_v1, string _name) returns()
func (_Curvefi *CurvefiSession) AddPool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_info [32]byte, _decimals *big.Int, _underlying_decimals *big.Int, _has_initial_A bool, _is_v1 bool, _name string) (*types.Transaction, error) {
	return _Curvefi.Contract.AddPool(&_Curvefi.TransactOpts, _pool, _n_coins, _lp_token, _rate_info, _decimals, _underlying_decimals, _has_initial_A, _is_v1, _name)
}

// AddPool is a paid mutator transaction binding the contract method 0x99209aa1.
//
// Solidity: function add_pool(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_info, uint256 _decimals, uint256 _underlying_decimals, bool _has_initial_A, bool _is_v1, string _name) returns()
func (_Curvefi *CurvefiTransactorSession) AddPool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_info [32]byte, _decimals *big.Int, _underlying_decimals *big.Int, _has_initial_A bool, _is_v1 bool, _name string) (*types.Transaction, error) {
	return _Curvefi.Contract.AddPool(&_Curvefi.TransactOpts, _pool, _n_coins, _lp_token, _rate_info, _decimals, _underlying_decimals, _has_initial_A, _is_v1, _name)
}

// AddPoolWithoutUnderlying is a paid mutator transaction binding the contract method 0xdcee86a3.
//
// Solidity: function add_pool_without_underlying(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_info, uint256 _decimals, uint256 _use_rates, bool _has_initial_A, bool _is_v1, string _name) returns()
func (_Curvefi *CurvefiTransactor) AddPoolWithoutUnderlying(opts *bind.TransactOpts, _pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_info [32]byte, _decimals *big.Int, _use_rates *big.Int, _has_initial_A bool, _is_v1 bool, _name string) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "add_pool_without_underlying", _pool, _n_coins, _lp_token, _rate_info, _decimals, _use_rates, _has_initial_A, _is_v1, _name)
}

// AddPoolWithoutUnderlying is a paid mutator transaction binding the contract method 0xdcee86a3.
//
// Solidity: function add_pool_without_underlying(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_info, uint256 _decimals, uint256 _use_rates, bool _has_initial_A, bool _is_v1, string _name) returns()
func (_Curvefi *CurvefiSession) AddPoolWithoutUnderlying(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_info [32]byte, _decimals *big.Int, _use_rates *big.Int, _has_initial_A bool, _is_v1 bool, _name string) (*types.Transaction, error) {
	return _Curvefi.Contract.AddPoolWithoutUnderlying(&_Curvefi.TransactOpts, _pool, _n_coins, _lp_token, _rate_info, _decimals, _use_rates, _has_initial_A, _is_v1, _name)
}

// AddPoolWithoutUnderlying is a paid mutator transaction binding the contract method 0xdcee86a3.
//
// Solidity: function add_pool_without_underlying(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_info, uint256 _decimals, uint256 _use_rates, bool _has_initial_A, bool _is_v1, string _name) returns()
func (_Curvefi *CurvefiTransactorSession) AddPoolWithoutUnderlying(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_info [32]byte, _decimals *big.Int, _use_rates *big.Int, _has_initial_A bool, _is_v1 bool, _name string) (*types.Transaction, error) {
	return _Curvefi.Contract.AddPoolWithoutUnderlying(&_Curvefi.TransactOpts, _pool, _n_coins, _lp_token, _rate_info, _decimals, _use_rates, _has_initial_A, _is_v1, _name)
}

// BatchSetPoolAssetType is a paid mutator transaction binding the contract method 0x7542f078.
//
// Solidity: function batch_set_pool_asset_type(address[32] _pools, uint256[32] _asset_types) returns()
func (_Curvefi *CurvefiTransactor) BatchSetPoolAssetType(opts *bind.TransactOpts, _pools [32]common.Address, _asset_types [32]*big.Int) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "batch_set_pool_asset_type", _pools, _asset_types)
}

// BatchSetPoolAssetType is a paid mutator transaction binding the contract method 0x7542f078.
//
// Solidity: function batch_set_pool_asset_type(address[32] _pools, uint256[32] _asset_types) returns()
func (_Curvefi *CurvefiSession) BatchSetPoolAssetType(_pools [32]common.Address, _asset_types [32]*big.Int) (*types.Transaction, error) {
	return _Curvefi.Contract.BatchSetPoolAssetType(&_Curvefi.TransactOpts, _pools, _asset_types)
}

// BatchSetPoolAssetType is a paid mutator transaction binding the contract method 0x7542f078.
//
// Solidity: function batch_set_pool_asset_type(address[32] _pools, uint256[32] _asset_types) returns()
func (_Curvefi *CurvefiTransactorSession) BatchSetPoolAssetType(_pools [32]common.Address, _asset_types [32]*big.Int) (*types.Transaction, error) {
	return _Curvefi.Contract.BatchSetPoolAssetType(&_Curvefi.TransactOpts, _pools, _asset_types)
}

// RemovePool is a paid mutator transaction binding the contract method 0x474932b0.
//
// Solidity: function remove_pool(address _pool) returns()
func (_Curvefi *CurvefiTransactor) RemovePool(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "remove_pool", _pool)
}

// RemovePool is a paid mutator transaction binding the contract method 0x474932b0.
//
// Solidity: function remove_pool(address _pool) returns()
func (_Curvefi *CurvefiSession) RemovePool(_pool common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.RemovePool(&_Curvefi.TransactOpts, _pool)
}

// RemovePool is a paid mutator transaction binding the contract method 0x474932b0.
//
// Solidity: function remove_pool(address _pool) returns()
func (_Curvefi *CurvefiTransactorSession) RemovePool(_pool common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.RemovePool(&_Curvefi.TransactOpts, _pool)
}

// SetCoinGasEstimates is a paid mutator transaction binding the contract method 0x237f89f2.
//
// Solidity: function set_coin_gas_estimates(address[10] _addr, uint256[10] _amount) returns()
func (_Curvefi *CurvefiTransactor) SetCoinGasEstimates(opts *bind.TransactOpts, _addr [10]common.Address, _amount [10]*big.Int) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "set_coin_gas_estimates", _addr, _amount)
}

// SetCoinGasEstimates is a paid mutator transaction binding the contract method 0x237f89f2.
//
// Solidity: function set_coin_gas_estimates(address[10] _addr, uint256[10] _amount) returns()
func (_Curvefi *CurvefiSession) SetCoinGasEstimates(_addr [10]common.Address, _amount [10]*big.Int) (*types.Transaction, error) {
	return _Curvefi.Contract.SetCoinGasEstimates(&_Curvefi.TransactOpts, _addr, _amount)
}

// SetCoinGasEstimates is a paid mutator transaction binding the contract method 0x237f89f2.
//
// Solidity: function set_coin_gas_estimates(address[10] _addr, uint256[10] _amount) returns()
func (_Curvefi *CurvefiTransactorSession) SetCoinGasEstimates(_addr [10]common.Address, _amount [10]*big.Int) (*types.Transaction, error) {
	return _Curvefi.Contract.SetCoinGasEstimates(&_Curvefi.TransactOpts, _addr, _amount)
}

// SetGasEstimateContract is a paid mutator transaction binding the contract method 0xca991b14.
//
// Solidity: function set_gas_estimate_contract(address _pool, address _estimator) returns()
func (_Curvefi *CurvefiTransactor) SetGasEstimateContract(opts *bind.TransactOpts, _pool common.Address, _estimator common.Address) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "set_gas_estimate_contract", _pool, _estimator)
}

// SetGasEstimateContract is a paid mutator transaction binding the contract method 0xca991b14.
//
// Solidity: function set_gas_estimate_contract(address _pool, address _estimator) returns()
func (_Curvefi *CurvefiSession) SetGasEstimateContract(_pool common.Address, _estimator common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.SetGasEstimateContract(&_Curvefi.TransactOpts, _pool, _estimator)
}

// SetGasEstimateContract is a paid mutator transaction binding the contract method 0xca991b14.
//
// Solidity: function set_gas_estimate_contract(address _pool, address _estimator) returns()
func (_Curvefi *CurvefiTransactorSession) SetGasEstimateContract(_pool common.Address, _estimator common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.SetGasEstimateContract(&_Curvefi.TransactOpts, _pool, _estimator)
}

// SetLiquidityGauges is a paid mutator transaction binding the contract method 0xef6b9788.
//
// Solidity: function set_liquidity_gauges(address _pool, address[10] _liquidity_gauges) returns()
func (_Curvefi *CurvefiTransactor) SetLiquidityGauges(opts *bind.TransactOpts, _pool common.Address, _liquidity_gauges [10]common.Address) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "set_liquidity_gauges", _pool, _liquidity_gauges)
}

// SetLiquidityGauges is a paid mutator transaction binding the contract method 0xef6b9788.
//
// Solidity: function set_liquidity_gauges(address _pool, address[10] _liquidity_gauges) returns()
func (_Curvefi *CurvefiSession) SetLiquidityGauges(_pool common.Address, _liquidity_gauges [10]common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.SetLiquidityGauges(&_Curvefi.TransactOpts, _pool, _liquidity_gauges)
}

// SetLiquidityGauges is a paid mutator transaction binding the contract method 0xef6b9788.
//
// Solidity: function set_liquidity_gauges(address _pool, address[10] _liquidity_gauges) returns()
func (_Curvefi *CurvefiTransactorSession) SetLiquidityGauges(_pool common.Address, _liquidity_gauges [10]common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.SetLiquidityGauges(&_Curvefi.TransactOpts, _pool, _liquidity_gauges)
}

// SetPoolAssetType is a paid mutator transaction binding the contract method 0x09e76774.
//
// Solidity: function set_pool_asset_type(address _pool, uint256 _asset_type) returns()
func (_Curvefi *CurvefiTransactor) SetPoolAssetType(opts *bind.TransactOpts, _pool common.Address, _asset_type *big.Int) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "set_pool_asset_type", _pool, _asset_type)
}

// SetPoolAssetType is a paid mutator transaction binding the contract method 0x09e76774.
//
// Solidity: function set_pool_asset_type(address _pool, uint256 _asset_type) returns()
func (_Curvefi *CurvefiSession) SetPoolAssetType(_pool common.Address, _asset_type *big.Int) (*types.Transaction, error) {
	return _Curvefi.Contract.SetPoolAssetType(&_Curvefi.TransactOpts, _pool, _asset_type)
}

// SetPoolAssetType is a paid mutator transaction binding the contract method 0x09e76774.
//
// Solidity: function set_pool_asset_type(address _pool, uint256 _asset_type) returns()
func (_Curvefi *CurvefiTransactorSession) SetPoolAssetType(_pool common.Address, _asset_type *big.Int) (*types.Transaction, error) {
	return _Curvefi.Contract.SetPoolAssetType(&_Curvefi.TransactOpts, _pool, _asset_type)
}

// SetPoolGasEstimates is a paid mutator transaction binding the contract method 0x0733b67a.
//
// Solidity: function set_pool_gas_estimates(address[5] _addr, uint256[2][5] _amount) returns()
func (_Curvefi *CurvefiTransactor) SetPoolGasEstimates(opts *bind.TransactOpts, _addr [5]common.Address, _amount [5][2]*big.Int) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "set_pool_gas_estimates", _addr, _amount)
}

// SetPoolGasEstimates is a paid mutator transaction binding the contract method 0x0733b67a.
//
// Solidity: function set_pool_gas_estimates(address[5] _addr, uint256[2][5] _amount) returns()
func (_Curvefi *CurvefiSession) SetPoolGasEstimates(_addr [5]common.Address, _amount [5][2]*big.Int) (*types.Transaction, error) {
	return _Curvefi.Contract.SetPoolGasEstimates(&_Curvefi.TransactOpts, _addr, _amount)
}

// SetPoolGasEstimates is a paid mutator transaction binding the contract method 0x0733b67a.
//
// Solidity: function set_pool_gas_estimates(address[5] _addr, uint256[2][5] _amount) returns()
func (_Curvefi *CurvefiTransactorSession) SetPoolGasEstimates(_addr [5]common.Address, _amount [5][2]*big.Int) (*types.Transaction, error) {
	return _Curvefi.Contract.SetPoolGasEstimates(&_Curvefi.TransactOpts, _addr, _amount)
}

// CurvefiPoolAddedIterator is returned from FilterPoolAdded and is used to iterate over the raw logs and unpacked data for PoolAdded events raised by the Curvefi contract.
type CurvefiPoolAddedIterator struct {
	Event *CurvefiPoolAdded // Event containing the contract specifics and raw log

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
func (it *CurvefiPoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefiPoolAdded)
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
		it.Event = new(CurvefiPoolAdded)
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
func (it *CurvefiPoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefiPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefiPoolAdded represents a PoolAdded event raised by the Curvefi contract.
type CurvefiPoolAdded struct {
	Pool         common.Address
	RateMethodId []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPoolAdded is a free log retrieval operation binding the contract event 0xe485c16479ab7092c0b3fc4649843c06be7f072194675261590c84473ab0aea9.
//
// Solidity: event PoolAdded(address indexed pool, bytes rate_method_id)
func (_Curvefi *CurvefiFilterer) FilterPoolAdded(opts *bind.FilterOpts, pool []common.Address) (*CurvefiPoolAddedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Curvefi.contract.FilterLogs(opts, "PoolAdded", poolRule)
	if err != nil {
		return nil, err
	}
	return &CurvefiPoolAddedIterator{contract: _Curvefi.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

// WatchPoolAdded is a free log subscription operation binding the contract event 0xe485c16479ab7092c0b3fc4649843c06be7f072194675261590c84473ab0aea9.
//
// Solidity: event PoolAdded(address indexed pool, bytes rate_method_id)
func (_Curvefi *CurvefiFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *CurvefiPoolAdded, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Curvefi.contract.WatchLogs(opts, "PoolAdded", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefiPoolAdded)
				if err := _Curvefi.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

// ParsePoolAdded is a log parse operation binding the contract event 0xe485c16479ab7092c0b3fc4649843c06be7f072194675261590c84473ab0aea9.
//
// Solidity: event PoolAdded(address indexed pool, bytes rate_method_id)
func (_Curvefi *CurvefiFilterer) ParsePoolAdded(log types.Log) (*CurvefiPoolAdded, error) {
	event := new(CurvefiPoolAdded)
	if err := _Curvefi.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvefiPoolRemovedIterator is returned from FilterPoolRemoved and is used to iterate over the raw logs and unpacked data for PoolRemoved events raised by the Curvefi contract.
type CurvefiPoolRemovedIterator struct {
	Event *CurvefiPoolRemoved // Event containing the contract specifics and raw log

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
func (it *CurvefiPoolRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefiPoolRemoved)
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
		it.Event = new(CurvefiPoolRemoved)
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
func (it *CurvefiPoolRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefiPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefiPoolRemoved represents a PoolRemoved event raised by the Curvefi contract.
type CurvefiPoolRemoved struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolRemoved is a free log retrieval operation binding the contract event 0x4106dfdaa577573db51c0ca93f766dbedfa0758faa2e7f5bcdb7c142be803c3f.
//
// Solidity: event PoolRemoved(address indexed pool)
func (_Curvefi *CurvefiFilterer) FilterPoolRemoved(opts *bind.FilterOpts, pool []common.Address) (*CurvefiPoolRemovedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Curvefi.contract.FilterLogs(opts, "PoolRemoved", poolRule)
	if err != nil {
		return nil, err
	}
	return &CurvefiPoolRemovedIterator{contract: _Curvefi.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

// WatchPoolRemoved is a free log subscription operation binding the contract event 0x4106dfdaa577573db51c0ca93f766dbedfa0758faa2e7f5bcdb7c142be803c3f.
//
// Solidity: event PoolRemoved(address indexed pool)
func (_Curvefi *CurvefiFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *CurvefiPoolRemoved, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Curvefi.contract.WatchLogs(opts, "PoolRemoved", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefiPoolRemoved)
				if err := _Curvefi.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

// ParsePoolRemoved is a log parse operation binding the contract event 0x4106dfdaa577573db51c0ca93f766dbedfa0758faa2e7f5bcdb7c142be803c3f.
//
// Solidity: event PoolRemoved(address indexed pool)
func (_Curvefi *CurvefiFilterer) ParsePoolRemoved(log types.Log) (*CurvefiPoolRemoved, error) {
	event := new(CurvefiPoolRemoved)
	if err := _Curvefi.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
