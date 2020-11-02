// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curvefi

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Coins              [8]common.Address
	UnderlyingCoins    [8]common.Address
	Decimals           [8]*big.Int
	UnderlyingDecimals [8]*big.Int
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	Balances           [8]*big.Int
	UnderlyingBalances [8]*big.Int
	Decimals           [8]*big.Int
	UnderlyingDecimals [8]*big.Int
	LpToken            common.Address
	A                  *big.Int
	Fee                *big.Int
}

// CurvefiABI is the input ABI used to generate the binding from.
const CurvefiABI = "[{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true},{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchange\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"address\",\"name\":\"pool\",\"indexed\":true},{\"type\":\"address\",\"name\":\"token_sold\",\"indexed\":false},{\"type\":\"address\",\"name\":\"token_bought\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount_sold\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount_bought\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"PoolAdded\",\"inputs\":[{\"type\":\"address\",\"name\":\"pool\",\"indexed\":true},{\"type\":\"bytes\",\"name\":\"rate_method_id\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"PoolRemoved\",\"inputs\":[{\"type\":\"address\",\"name\":\"pool\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address[4]\",\"name\":\"_returns_none\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"constant\":false,\"payable\":true,\"type\":\"fallback\"},{\"name\":\"find_pool_for_coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"}],\"constant\":true,\"payable\":false,\"type\":\"function\"},{\"name\":\"find_pool_for_coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"i\"}],\"constant\":true,\"payable\":false,\"type\":\"function\"},{\"name\":\"get_pool_coins\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"components\":[{\"type\":\"address[8]\",\"name\":\"coins\"},{\"type\":\"address[8]\",\"name\":\"underlying_coins\"},{\"type\":\"uint256[8]\",\"name\":\"decimals\"},{\"type\":\"uint256[8]\",\"name\":\"underlying_decimals\"}]}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":45681},{\"name\":\"get_pool_info\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"components\":[{\"type\":\"uint256[8]\",\"name\":\"balances\"},{\"type\":\"uint256[8]\",\"name\":\"underlying_balances\"},{\"type\":\"uint256[8]\",\"name\":\"decimals\"},{\"type\":\"uint256[8]\",\"name\":\"underlying_decimals\"},{\"type\":\"address\",\"name\":\"lp_token\"},{\"type\":\"uint256\",\"name\":\"A\"},{\"type\":\"uint256\",\"name\":\"fee\"}]}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":87733},{\"name\":\"get_pool_rates\",\"outputs\":[{\"type\":\"uint256[8]\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":44706},{\"name\":\"estimate_gas_used\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":50163},{\"name\":\"get_exchange_amount\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_amount\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":45192},{\"name\":\"exchange\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_amount\"},{\"type\":\"uint256\",\"name\":\"_expected\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":198813},{\"name\":\"get_input_amount\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":135410},{\"name\":\"get_exchange_amounts\",\"outputs\":[{\"type\":\"uint256[100]\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256[100]\",\"name\":\"_amounts\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":138709},{\"name\":\"add_pool\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"int128\",\"name\":\"_n_coins\"},{\"type\":\"address\",\"name\":\"_lp_token\"},{\"type\":\"address\",\"name\":\"_calculator\"},{\"type\":\"bytes32\",\"name\":\"_rate_method_id\"},{\"type\":\"bytes32\",\"name\":\"_decimals\"},{\"type\":\"bytes32\",\"name\":\"_underlying_decimals\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":10290194},{\"name\":\"add_pool_without_underlying\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"int128\",\"name\":\"_n_coins\"},{\"type\":\"address\",\"name\":\"_lp_token\"},{\"type\":\"address\",\"name\":\"_calculator\"},{\"type\":\"bytes32\",\"name\":\"_rate_method_id\"},{\"type\":\"bytes32\",\"name\":\"_decimals\"},{\"type\":\"bytes32\",\"name\":\"_use_rates\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":10273690},{\"name\":\"remove_pool\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":317792223500},{\"name\":\"set_returns_none\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"},{\"type\":\"bool\",\"name\":\"_is_returns_none\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":36818},{\"name\":\"set_pool_gas_estimates\",\"outputs\":[],\"inputs\":[{\"type\":\"address[5]\",\"name\":\"_addr\"},{\"type\":\"uint256[2][5]\",\"name\":\"_amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":354727},{\"name\":\"set_coin_gas_estimates\",\"outputs\":[],\"inputs\":[{\"type\":\"address[10]\",\"name\":\"_addr\"},{\"type\":\"uint256[10]\",\"name\":\"_amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":356796},{\"name\":\"set_gas_estimate_contract\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"address\",\"name\":\"_estimator\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":36911},{\"name\":\"set_calculator\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"address\",\"name\":\"_calculator\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":37019},{\"name\":\"get_calculator\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1953},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_new_admin\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":74452},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":60590},{\"name\":\"revert_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":21865},{\"name\":\"claim_token_balance\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":4854},{\"name\":\"claim_eth_balance\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":36726},{\"name\":\"admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1901},{\"name\":\"pool_list\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2130},{\"name\":\"pool_count\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1961}]"

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
func (_Curvefi *CurvefiRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_Curvefi *CurvefiCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Curvefi *CurvefiCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Curvefi.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Curvefi *CurvefiSession) Admin() (common.Address, error) {
	return _Curvefi.Contract.Admin(&_Curvefi.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Curvefi *CurvefiCallerSession) Admin() (common.Address, error) {
	return _Curvefi.Contract.Admin(&_Curvefi.CallOpts)
}

// EstimateGasUsed is a free data retrieval call binding the contract method 0xb0bb365b.
//
// Solidity: function estimate_gas_used(address _pool, address _from, address _to) returns(uint256)
func (_Curvefi *CurvefiCaller) EstimateGasUsed(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Curvefi.contract.Call(opts, out, "estimate_gas_used", _pool, _from, _to)
	return *ret0, err
}

// EstimateGasUsed is a free data retrieval call binding the contract method 0xb0bb365b.
//
// Solidity: function estimate_gas_used(address _pool, address _from, address _to) returns(uint256)
func (_Curvefi *CurvefiSession) EstimateGasUsed(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, error) {
	return _Curvefi.Contract.EstimateGasUsed(&_Curvefi.CallOpts, _pool, _from, _to)
}

// EstimateGasUsed is a free data retrieval call binding the contract method 0xb0bb365b.
//
// Solidity: function estimate_gas_used(address _pool, address _from, address _to) returns(uint256)
func (_Curvefi *CurvefiCallerSession) EstimateGasUsed(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, error) {
	return _Curvefi.Contract.EstimateGasUsed(&_Curvefi.CallOpts, _pool, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) returns(address)
func (_Curvefi *CurvefiCaller) FindPoolForCoins(opts *bind.CallOpts, _from common.Address, _to common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Curvefi.contract.Call(opts, out, "find_pool_for_coins", _from, _to)
	return *ret0, err
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) returns(address)
func (_Curvefi *CurvefiSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _Curvefi.Contract.FindPoolForCoins(&_Curvefi.CallOpts, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) returns(address)
func (_Curvefi *CurvefiCallerSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _Curvefi.Contract.FindPoolForCoins(&_Curvefi.CallOpts, _from, _to)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) returns(address)
func (_Curvefi *CurvefiCaller) FindPoolForCoins0(opts *bind.CallOpts, _from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Curvefi.contract.Call(opts, out, "find_pool_for_coins0", _from, _to, i)
	return *ret0, err
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) returns(address)
func (_Curvefi *CurvefiSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _Curvefi.Contract.FindPoolForCoins0(&_Curvefi.CallOpts, _from, _to, i)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) returns(address)
func (_Curvefi *CurvefiCallerSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _Curvefi.Contract.FindPoolForCoins0(&_Curvefi.CallOpts, _from, _to, i)
}

// GetCalculator is a free data retrieval call binding the contract method 0x5d7dc825.
//
// Solidity: function get_calculator(address _pool) returns(address)
func (_Curvefi *CurvefiCaller) GetCalculator(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Curvefi.contract.Call(opts, out, "get_calculator", _pool)
	return *ret0, err
}

// GetCalculator is a free data retrieval call binding the contract method 0x5d7dc825.
//
// Solidity: function get_calculator(address _pool) returns(address)
func (_Curvefi *CurvefiSession) GetCalculator(_pool common.Address) (common.Address, error) {
	return _Curvefi.Contract.GetCalculator(&_Curvefi.CallOpts, _pool)
}

// GetCalculator is a free data retrieval call binding the contract method 0x5d7dc825.
//
// Solidity: function get_calculator(address _pool) returns(address)
func (_Curvefi *CurvefiCallerSession) GetCalculator(_pool common.Address) (common.Address, error) {
	return _Curvefi.Contract.GetCalculator(&_Curvefi.CallOpts, _pool)
}

// GetExchangeAmount is a free data retrieval call binding the contract method 0x3973e834.
//
// Solidity: function get_exchange_amount(address _pool, address _from, address _to, uint256 _amount) returns(uint256)
func (_Curvefi *CurvefiCaller) GetExchangeAmount(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Curvefi.contract.Call(opts, out, "get_exchange_amount", _pool, _from, _to, _amount)
	return *ret0, err
}

// GetExchangeAmount is a free data retrieval call binding the contract method 0x3973e834.
//
// Solidity: function get_exchange_amount(address _pool, address _from, address _to, uint256 _amount) returns(uint256)
func (_Curvefi *CurvefiSession) GetExchangeAmount(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*big.Int, error) {
	return _Curvefi.Contract.GetExchangeAmount(&_Curvefi.CallOpts, _pool, _from, _to, _amount)
}

// GetExchangeAmount is a free data retrieval call binding the contract method 0x3973e834.
//
// Solidity: function get_exchange_amount(address _pool, address _from, address _to, uint256 _amount) returns(uint256)
func (_Curvefi *CurvefiCallerSession) GetExchangeAmount(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*big.Int, error) {
	return _Curvefi.Contract.GetExchangeAmount(&_Curvefi.CallOpts, _pool, _from, _to, _amount)
}

// GetPoolCoins is a free data retrieval call binding the contract method 0xe030afb8.
//
// Solidity: function get_pool_coins(address _pool) returns((address[8],address[8],uint256[8],uint256[8]))
func (_Curvefi *CurvefiCaller) GetPoolCoins(opts *bind.CallOpts, _pool common.Address) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _Curvefi.contract.Call(opts, out, "get_pool_coins", _pool)
	return *ret0, err
}

// GetPoolCoins is a free data retrieval call binding the contract method 0xe030afb8.
//
// Solidity: function get_pool_coins(address _pool) returns((address[8],address[8],uint256[8],uint256[8]))
func (_Curvefi *CurvefiSession) GetPoolCoins(_pool common.Address) (Struct0, error) {
	return _Curvefi.Contract.GetPoolCoins(&_Curvefi.CallOpts, _pool)
}

// GetPoolCoins is a free data retrieval call binding the contract method 0xe030afb8.
//
// Solidity: function get_pool_coins(address _pool) returns((address[8],address[8],uint256[8],uint256[8]))
func (_Curvefi *CurvefiCallerSession) GetPoolCoins(_pool common.Address) (Struct0, error) {
	return _Curvefi.Contract.GetPoolCoins(&_Curvefi.CallOpts, _pool)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() returns(uint256)
func (_Curvefi *CurvefiCaller) PoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Curvefi.contract.Call(opts, out, "pool_count")
	return *ret0, err
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() returns(uint256)
func (_Curvefi *CurvefiSession) PoolCount() (*big.Int, error) {
	return _Curvefi.Contract.PoolCount(&_Curvefi.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() returns(uint256)
func (_Curvefi *CurvefiCallerSession) PoolCount() (*big.Int, error) {
	return _Curvefi.Contract.PoolCount(&_Curvefi.CallOpts)
}

// PoolList is a free data retrieval call binding the contract method 0xe140b271.
//
// Solidity: function pool_list(int128 arg0) returns(address)
func (_Curvefi *CurvefiCaller) PoolList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Curvefi.contract.Call(opts, out, "pool_list", arg0)
	return *ret0, err
}

// PoolList is a free data retrieval call binding the contract method 0xe140b271.
//
// Solidity: function pool_list(int128 arg0) returns(address)
func (_Curvefi *CurvefiSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _Curvefi.Contract.PoolList(&_Curvefi.CallOpts, arg0)
}

// PoolList is a free data retrieval call binding the contract method 0xe140b271.
//
// Solidity: function pool_list(int128 arg0) returns(address)
func (_Curvefi *CurvefiCallerSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _Curvefi.Contract.PoolList(&_Curvefi.CallOpts, arg0)
}

// AddPool is a paid mutator transaction binding the contract method 0x74e6a310.
//
// Solidity: function add_pool(address _pool, int128 _n_coins, address _lp_token, address _calculator, bytes32 _rate_method_id, bytes32 _decimals, bytes32 _underlying_decimals) returns()
func (_Curvefi *CurvefiTransactor) AddPool(opts *bind.TransactOpts, _pool common.Address, _n_coins *big.Int, _lp_token common.Address, _calculator common.Address, _rate_method_id [32]byte, _decimals [32]byte, _underlying_decimals [32]byte) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "add_pool", _pool, _n_coins, _lp_token, _calculator, _rate_method_id, _decimals, _underlying_decimals)
}

// AddPool is a paid mutator transaction binding the contract method 0x74e6a310.
//
// Solidity: function add_pool(address _pool, int128 _n_coins, address _lp_token, address _calculator, bytes32 _rate_method_id, bytes32 _decimals, bytes32 _underlying_decimals) returns()
func (_Curvefi *CurvefiSession) AddPool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _calculator common.Address, _rate_method_id [32]byte, _decimals [32]byte, _underlying_decimals [32]byte) (*types.Transaction, error) {
	return _Curvefi.Contract.AddPool(&_Curvefi.TransactOpts, _pool, _n_coins, _lp_token, _calculator, _rate_method_id, _decimals, _underlying_decimals)
}

// AddPool is a paid mutator transaction binding the contract method 0x74e6a310.
//
// Solidity: function add_pool(address _pool, int128 _n_coins, address _lp_token, address _calculator, bytes32 _rate_method_id, bytes32 _decimals, bytes32 _underlying_decimals) returns()
func (_Curvefi *CurvefiTransactorSession) AddPool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _calculator common.Address, _rate_method_id [32]byte, _decimals [32]byte, _underlying_decimals [32]byte) (*types.Transaction, error) {
	return _Curvefi.Contract.AddPool(&_Curvefi.TransactOpts, _pool, _n_coins, _lp_token, _calculator, _rate_method_id, _decimals, _underlying_decimals)
}

// AddPoolWithoutUnderlying is a paid mutator transaction binding the contract method 0x356547c1.
//
// Solidity: function add_pool_without_underlying(address _pool, int128 _n_coins, address _lp_token, address _calculator, bytes32 _rate_method_id, bytes32 _decimals, bytes32 _use_rates) returns()
func (_Curvefi *CurvefiTransactor) AddPoolWithoutUnderlying(opts *bind.TransactOpts, _pool common.Address, _n_coins *big.Int, _lp_token common.Address, _calculator common.Address, _rate_method_id [32]byte, _decimals [32]byte, _use_rates [32]byte) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "add_pool_without_underlying", _pool, _n_coins, _lp_token, _calculator, _rate_method_id, _decimals, _use_rates)
}

// AddPoolWithoutUnderlying is a paid mutator transaction binding the contract method 0x356547c1.
//
// Solidity: function add_pool_without_underlying(address _pool, int128 _n_coins, address _lp_token, address _calculator, bytes32 _rate_method_id, bytes32 _decimals, bytes32 _use_rates) returns()
func (_Curvefi *CurvefiSession) AddPoolWithoutUnderlying(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _calculator common.Address, _rate_method_id [32]byte, _decimals [32]byte, _use_rates [32]byte) (*types.Transaction, error) {
	return _Curvefi.Contract.AddPoolWithoutUnderlying(&_Curvefi.TransactOpts, _pool, _n_coins, _lp_token, _calculator, _rate_method_id, _decimals, _use_rates)
}

// AddPoolWithoutUnderlying is a paid mutator transaction binding the contract method 0x356547c1.
//
// Solidity: function add_pool_without_underlying(address _pool, int128 _n_coins, address _lp_token, address _calculator, bytes32 _rate_method_id, bytes32 _decimals, bytes32 _use_rates) returns()
func (_Curvefi *CurvefiTransactorSession) AddPoolWithoutUnderlying(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _calculator common.Address, _rate_method_id [32]byte, _decimals [32]byte, _use_rates [32]byte) (*types.Transaction, error) {
	return _Curvefi.Contract.AddPoolWithoutUnderlying(&_Curvefi.TransactOpts, _pool, _n_coins, _lp_token, _calculator, _rate_method_id, _decimals, _use_rates)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Curvefi *CurvefiTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Curvefi *CurvefiSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Curvefi.Contract.ApplyTransferOwnership(&_Curvefi.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Curvefi *CurvefiTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Curvefi.Contract.ApplyTransferOwnership(&_Curvefi.TransactOpts)
}

// ClaimEthBalance is a paid mutator transaction binding the contract method 0x2295f41f.
//
// Solidity: function claim_eth_balance() returns()
func (_Curvefi *CurvefiTransactor) ClaimEthBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "claim_eth_balance")
}

// ClaimEthBalance is a paid mutator transaction binding the contract method 0x2295f41f.
//
// Solidity: function claim_eth_balance() returns()
func (_Curvefi *CurvefiSession) ClaimEthBalance() (*types.Transaction, error) {
	return _Curvefi.Contract.ClaimEthBalance(&_Curvefi.TransactOpts)
}

// ClaimEthBalance is a paid mutator transaction binding the contract method 0x2295f41f.
//
// Solidity: function claim_eth_balance() returns()
func (_Curvefi *CurvefiTransactorSession) ClaimEthBalance() (*types.Transaction, error) {
	return _Curvefi.Contract.ClaimEthBalance(&_Curvefi.TransactOpts)
}

// ClaimTokenBalance is a paid mutator transaction binding the contract method 0x93cfe5f5.
//
// Solidity: function claim_token_balance(address _token) returns()
func (_Curvefi *CurvefiTransactor) ClaimTokenBalance(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "claim_token_balance", _token)
}

// ClaimTokenBalance is a paid mutator transaction binding the contract method 0x93cfe5f5.
//
// Solidity: function claim_token_balance(address _token) returns()
func (_Curvefi *CurvefiSession) ClaimTokenBalance(_token common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.ClaimTokenBalance(&_Curvefi.TransactOpts, _token)
}

// ClaimTokenBalance is a paid mutator transaction binding the contract method 0x93cfe5f5.
//
// Solidity: function claim_token_balance(address _token) returns()
func (_Curvefi *CurvefiTransactorSession) ClaimTokenBalance(_token common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.ClaimTokenBalance(&_Curvefi.TransactOpts, _token)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _new_admin) returns()
func (_Curvefi *CurvefiTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _new_admin common.Address) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "commit_transfer_ownership", _new_admin)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _new_admin) returns()
func (_Curvefi *CurvefiSession) CommitTransferOwnership(_new_admin common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.CommitTransferOwnership(&_Curvefi.TransactOpts, _new_admin)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _new_admin) returns()
func (_Curvefi *CurvefiTransactorSession) CommitTransferOwnership(_new_admin common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.CommitTransferOwnership(&_Curvefi.TransactOpts, _new_admin)
}

// Exchange is a paid mutator transaction binding the contract method 0x4798ce5b.
//
// Solidity: function exchange(address _pool, address _from, address _to, uint256 _amount, uint256 _expected) returns(bool)
func (_Curvefi *CurvefiTransactor) Exchange(opts *bind.TransactOpts, _pool common.Address, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "exchange", _pool, _from, _to, _amount, _expected)
}

// Exchange is a paid mutator transaction binding the contract method 0x4798ce5b.
//
// Solidity: function exchange(address _pool, address _from, address _to, uint256 _amount, uint256 _expected) returns(bool)
func (_Curvefi *CurvefiSession) Exchange(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Curvefi.Contract.Exchange(&_Curvefi.TransactOpts, _pool, _from, _to, _amount, _expected)
}

// Exchange is a paid mutator transaction binding the contract method 0x4798ce5b.
//
// Solidity: function exchange(address _pool, address _from, address _to, uint256 _amount, uint256 _expected) returns(bool)
func (_Curvefi *CurvefiTransactorSession) Exchange(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Curvefi.Contract.Exchange(&_Curvefi.TransactOpts, _pool, _from, _to, _amount, _expected)
}

// GetExchangeAmounts is a paid mutator transaction binding the contract method 0x4be9ae42.
//
// Solidity: function get_exchange_amounts(address _pool, address _from, address _to, uint256[100] _amounts) returns(uint256[100])
func (_Curvefi *CurvefiTransactor) GetExchangeAmounts(opts *bind.TransactOpts, _pool common.Address, _from common.Address, _to common.Address, _amounts [100]*big.Int) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "get_exchange_amounts", _pool, _from, _to, _amounts)
}

// GetExchangeAmounts is a paid mutator transaction binding the contract method 0x4be9ae42.
//
// Solidity: function get_exchange_amounts(address _pool, address _from, address _to, uint256[100] _amounts) returns(uint256[100])
func (_Curvefi *CurvefiSession) GetExchangeAmounts(_pool common.Address, _from common.Address, _to common.Address, _amounts [100]*big.Int) (*types.Transaction, error) {
	return _Curvefi.Contract.GetExchangeAmounts(&_Curvefi.TransactOpts, _pool, _from, _to, _amounts)
}

// GetExchangeAmounts is a paid mutator transaction binding the contract method 0x4be9ae42.
//
// Solidity: function get_exchange_amounts(address _pool, address _from, address _to, uint256[100] _amounts) returns(uint256[100])
func (_Curvefi *CurvefiTransactorSession) GetExchangeAmounts(_pool common.Address, _from common.Address, _to common.Address, _amounts [100]*big.Int) (*types.Transaction, error) {
	return _Curvefi.Contract.GetExchangeAmounts(&_Curvefi.TransactOpts, _pool, _from, _to, _amounts)
}

// GetInputAmount is a paid mutator transaction binding the contract method 0x7fa5a654.
//
// Solidity: function get_input_amount(address _pool, address _from, address _to, uint256 _amount) returns(uint256)
func (_Curvefi *CurvefiTransactor) GetInputAmount(opts *bind.TransactOpts, _pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "get_input_amount", _pool, _from, _to, _amount)
}

// GetInputAmount is a paid mutator transaction binding the contract method 0x7fa5a654.
//
// Solidity: function get_input_amount(address _pool, address _from, address _to, uint256 _amount) returns(uint256)
func (_Curvefi *CurvefiSession) GetInputAmount(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Curvefi.Contract.GetInputAmount(&_Curvefi.TransactOpts, _pool, _from, _to, _amount)
}

// GetInputAmount is a paid mutator transaction binding the contract method 0x7fa5a654.
//
// Solidity: function get_input_amount(address _pool, address _from, address _to, uint256 _amount) returns(uint256)
func (_Curvefi *CurvefiTransactorSession) GetInputAmount(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Curvefi.Contract.GetInputAmount(&_Curvefi.TransactOpts, _pool, _from, _to, _amount)
}

// GetPoolInfo is a paid mutator transaction binding the contract method 0x100f2c00.
//
// Solidity: function get_pool_info(address _pool) returns((uint256[8],uint256[8],uint256[8],uint256[8],address,uint256,uint256))
func (_Curvefi *CurvefiTransactor) GetPoolInfo(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "get_pool_info", _pool)
}

// GetPoolInfo is a paid mutator transaction binding the contract method 0x100f2c00.
//
// Solidity: function get_pool_info(address _pool) returns((uint256[8],uint256[8],uint256[8],uint256[8],address,uint256,uint256))
func (_Curvefi *CurvefiSession) GetPoolInfo(_pool common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.GetPoolInfo(&_Curvefi.TransactOpts, _pool)
}

// GetPoolInfo is a paid mutator transaction binding the contract method 0x100f2c00.
//
// Solidity: function get_pool_info(address _pool) returns((uint256[8],uint256[8],uint256[8],uint256[8],address,uint256,uint256))
func (_Curvefi *CurvefiTransactorSession) GetPoolInfo(_pool common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.GetPoolInfo(&_Curvefi.TransactOpts, _pool)
}

// GetPoolRates is a paid mutator transaction binding the contract method 0x5af6a160.
//
// Solidity: function get_pool_rates(address _pool) returns(uint256[8])
func (_Curvefi *CurvefiTransactor) GetPoolRates(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "get_pool_rates", _pool)
}

// GetPoolRates is a paid mutator transaction binding the contract method 0x5af6a160.
//
// Solidity: function get_pool_rates(address _pool) returns(uint256[8])
func (_Curvefi *CurvefiSession) GetPoolRates(_pool common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.GetPoolRates(&_Curvefi.TransactOpts, _pool)
}

// GetPoolRates is a paid mutator transaction binding the contract method 0x5af6a160.
//
// Solidity: function get_pool_rates(address _pool) returns(uint256[8])
func (_Curvefi *CurvefiTransactorSession) GetPoolRates(_pool common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.GetPoolRates(&_Curvefi.TransactOpts, _pool)
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

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Curvefi *CurvefiTransactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Curvefi *CurvefiSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Curvefi.Contract.RevertTransferOwnership(&_Curvefi.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Curvefi *CurvefiTransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Curvefi.Contract.RevertTransferOwnership(&_Curvefi.TransactOpts)
}

// SetCalculator is a paid mutator transaction binding the contract method 0x188c7ee5.
//
// Solidity: function set_calculator(address _pool, address _calculator) returns()
func (_Curvefi *CurvefiTransactor) SetCalculator(opts *bind.TransactOpts, _pool common.Address, _calculator common.Address) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "set_calculator", _pool, _calculator)
}

// SetCalculator is a paid mutator transaction binding the contract method 0x188c7ee5.
//
// Solidity: function set_calculator(address _pool, address _calculator) returns()
func (_Curvefi *CurvefiSession) SetCalculator(_pool common.Address, _calculator common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.SetCalculator(&_Curvefi.TransactOpts, _pool, _calculator)
}

// SetCalculator is a paid mutator transaction binding the contract method 0x188c7ee5.
//
// Solidity: function set_calculator(address _pool, address _calculator) returns()
func (_Curvefi *CurvefiTransactorSession) SetCalculator(_pool common.Address, _calculator common.Address) (*types.Transaction, error) {
	return _Curvefi.Contract.SetCalculator(&_Curvefi.TransactOpts, _pool, _calculator)
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

// SetReturnsNone is a paid mutator transaction binding the contract method 0xb55e47ab.
//
// Solidity: function set_returns_none(address _addr, bool _is_returns_none) returns()
func (_Curvefi *CurvefiTransactor) SetReturnsNone(opts *bind.TransactOpts, _addr common.Address, _is_returns_none bool) (*types.Transaction, error) {
	return _Curvefi.contract.Transact(opts, "set_returns_none", _addr, _is_returns_none)
}

// SetReturnsNone is a paid mutator transaction binding the contract method 0xb55e47ab.
//
// Solidity: function set_returns_none(address _addr, bool _is_returns_none) returns()
func (_Curvefi *CurvefiSession) SetReturnsNone(_addr common.Address, _is_returns_none bool) (*types.Transaction, error) {
	return _Curvefi.Contract.SetReturnsNone(&_Curvefi.TransactOpts, _addr, _is_returns_none)
}

// SetReturnsNone is a paid mutator transaction binding the contract method 0xb55e47ab.
//
// Solidity: function set_returns_none(address _addr, bool _is_returns_none) returns()
func (_Curvefi *CurvefiTransactorSession) SetReturnsNone(_addr common.Address, _is_returns_none bool) (*types.Transaction, error) {
	return _Curvefi.Contract.SetReturnsNone(&_Curvefi.TransactOpts, _addr, _is_returns_none)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Curvefi *CurvefiTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Curvefi.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Curvefi *CurvefiSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Curvefi.Contract.Fallback(&_Curvefi.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Curvefi *CurvefiTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Curvefi.Contract.Fallback(&_Curvefi.TransactOpts, calldata)
}

// CurvefiCommitNewAdminIterator is returned from FilterCommitNewAdmin and is used to iterate over the raw logs and unpacked data for CommitNewAdmin events raised by the Curvefi contract.
type CurvefiCommitNewAdminIterator struct {
	Event *CurvefiCommitNewAdmin // Event containing the contract specifics and raw log

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
func (it *CurvefiCommitNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefiCommitNewAdmin)
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
		it.Event = new(CurvefiCommitNewAdmin)
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
func (it *CurvefiCommitNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefiCommitNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefiCommitNewAdmin represents a CommitNewAdmin event raised by the Curvefi contract.
type CurvefiCommitNewAdmin struct {
	Deadline *big.Int
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewAdmin is a free log retrieval operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Curvefi *CurvefiFilterer) FilterCommitNewAdmin(opts *bind.FilterOpts, deadline []*big.Int, admin []common.Address) (*CurvefiCommitNewAdminIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Curvefi.contract.FilterLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &CurvefiCommitNewAdminIterator{contract: _Curvefi.contract, event: "CommitNewAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitNewAdmin is a free log subscription operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Curvefi *CurvefiFilterer) WatchCommitNewAdmin(opts *bind.WatchOpts, sink chan<- *CurvefiCommitNewAdmin, deadline []*big.Int, admin []common.Address) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Curvefi.contract.WatchLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefiCommitNewAdmin)
				if err := _Curvefi.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
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
func (_Curvefi *CurvefiFilterer) ParseCommitNewAdmin(log types.Log) (*CurvefiCommitNewAdmin, error) {
	event := new(CurvefiCommitNewAdmin)
	if err := _Curvefi.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CurvefiNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Curvefi contract.
type CurvefiNewAdminIterator struct {
	Event *CurvefiNewAdmin // Event containing the contract specifics and raw log

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
func (it *CurvefiNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefiNewAdmin)
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
		it.Event = new(CurvefiNewAdmin)
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
func (it *CurvefiNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefiNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefiNewAdmin represents a NewAdmin event raised by the Curvefi contract.
type CurvefiNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Curvefi *CurvefiFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address) (*CurvefiNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Curvefi.contract.FilterLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &CurvefiNewAdminIterator{contract: _Curvefi.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Curvefi *CurvefiFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *CurvefiNewAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Curvefi.contract.WatchLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefiNewAdmin)
				if err := _Curvefi.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_Curvefi *CurvefiFilterer) ParseNewAdmin(log types.Log) (*CurvefiNewAdmin, error) {
	event := new(CurvefiNewAdmin)
	if err := _Curvefi.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
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
	return event, nil
}

// CurvefiTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the Curvefi contract.
type CurvefiTokenExchangeIterator struct {
	Event *CurvefiTokenExchange // Event containing the contract specifics and raw log

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
func (it *CurvefiTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvefiTokenExchange)
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
		it.Event = new(CurvefiTokenExchange)
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
func (it *CurvefiTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvefiTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvefiTokenExchange represents a TokenExchange event raised by the Curvefi contract.
type CurvefiTokenExchange struct {
	Buyer        common.Address
	Pool         common.Address
	TokenSold    common.Address
	TokenBought  common.Address
	AmountSold   *big.Int
	AmountBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenExchange is a free log retrieval operation binding the contract event 0xc74278adc52a6f43a80f27768caac90c9609be033cf5189c150b0c6d173800b7.
//
// Solidity: event TokenExchange(address indexed buyer, address indexed pool, address token_sold, address token_bought, uint256 amount_sold, uint256 amount_bought)
func (_Curvefi *CurvefiFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address, pool []common.Address) (*CurvefiTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Curvefi.contract.FilterLogs(opts, "TokenExchange", buyerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &CurvefiTokenExchangeIterator{contract: _Curvefi.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0xc74278adc52a6f43a80f27768caac90c9609be033cf5189c150b0c6d173800b7.
//
// Solidity: event TokenExchange(address indexed buyer, address indexed pool, address token_sold, address token_bought, uint256 amount_sold, uint256 amount_bought)
func (_Curvefi *CurvefiFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *CurvefiTokenExchange, buyer []common.Address, pool []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Curvefi.contract.WatchLogs(opts, "TokenExchange", buyerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvefiTokenExchange)
				if err := _Curvefi.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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

// ParseTokenExchange is a log parse operation binding the contract event 0xc74278adc52a6f43a80f27768caac90c9609be033cf5189c150b0c6d173800b7.
//
// Solidity: event TokenExchange(address indexed buyer, address indexed pool, address token_sold, address token_bought, uint256 amount_sold, uint256 amount_bought)
func (_Curvefi *CurvefiFilterer) ParseTokenExchange(log types.Log) (*CurvefiTokenExchange, error) {
	event := new(CurvefiTokenExchange)
	if err := _Curvefi.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	return event, nil
}
