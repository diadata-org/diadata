// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gauge

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

// GaugeABI is the input ABI used to generate the binding from.
const GaugeABI = "[{\"name\":\"Deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateLiquidityLimit\",\"inputs\":[{\"type\":\"address\",\"name\":\"user\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"original_balance\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"original_supply\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"working_balance\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"working_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"lp_addr\"},{\"type\":\"address\",\"name\":\"_minter\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"user_checkpoint\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":2079152},{\"name\":\"claimable_tokens\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":1998318},{\"name\":\"kick\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":2084532},{\"name\":\"set_approve_deposit\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"bool\",\"name\":\"can_deposit\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":35766},{\"name\":\"deposit\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"deposit\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"},{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":2208318},{\"name\":\"integrate_checkpoint\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2297},{\"name\":\"minter\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1421},{\"name\":\"crv_token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1451},{\"name\":\"lp_token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1481},{\"name\":\"controller\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1511},{\"name\":\"voting_escrow\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1541},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1725},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1601},{\"name\":\"future_epoch_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1631},{\"name\":\"approved_to_deposit\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"address\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1969},{\"name\":\"working_balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1845},{\"name\":\"working_supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1721},{\"name\":\"period\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1751},{\"name\":\"period_timestamp\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1890},{\"name\":\"integrate_inv_supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1920},{\"name\":\"integrate_inv_supply_of\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1995},{\"name\":\"integrate_checkpoint_of\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2025},{\"name\":\"integrate_fraction\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2055},{\"name\":\"inflation_rate\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1931}]"

// Gauge is an auto generated Go binding around an Ethereum contract.
type Gauge struct {
	GaugeCaller     // Read-only binding to the contract
	GaugeTransactor // Write-only binding to the contract
	GaugeFilterer   // Log filterer for contract events
}

// GaugeCaller is an auto generated read-only Go binding around an Ethereum contract.
type GaugeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GaugeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GaugeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GaugeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GaugeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GaugeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GaugeSession struct {
	Contract     *Gauge            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GaugeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GaugeCallerSession struct {
	Contract *GaugeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GaugeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GaugeTransactorSession struct {
	Contract     *GaugeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GaugeRaw is an auto generated low-level Go binding around an Ethereum contract.
type GaugeRaw struct {
	Contract *Gauge // Generic contract binding to access the raw methods on
}

// GaugeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GaugeCallerRaw struct {
	Contract *GaugeCaller // Generic read-only contract binding to access the raw methods on
}

// GaugeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GaugeTransactorRaw struct {
	Contract *GaugeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGauge creates a new instance of Gauge, bound to a specific deployed contract.
func NewGauge(address common.Address, backend bind.ContractBackend) (*Gauge, error) {
	contract, err := bindGauge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gauge{GaugeCaller: GaugeCaller{contract: contract}, GaugeTransactor: GaugeTransactor{contract: contract}, GaugeFilterer: GaugeFilterer{contract: contract}}, nil
}

// NewGaugeCaller creates a new read-only instance of Gauge, bound to a specific deployed contract.
func NewGaugeCaller(address common.Address, caller bind.ContractCaller) (*GaugeCaller, error) {
	contract, err := bindGauge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GaugeCaller{contract: contract}, nil
}

// NewGaugeTransactor creates a new write-only instance of Gauge, bound to a specific deployed contract.
func NewGaugeTransactor(address common.Address, transactor bind.ContractTransactor) (*GaugeTransactor, error) {
	contract, err := bindGauge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GaugeTransactor{contract: contract}, nil
}

// NewGaugeFilterer creates a new log filterer instance of Gauge, bound to a specific deployed contract.
func NewGaugeFilterer(address common.Address, filterer bind.ContractFilterer) (*GaugeFilterer, error) {
	contract, err := bindGauge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GaugeFilterer{contract: contract}, nil
}

// bindGauge binds a generic wrapper to an already deployed contract.
func bindGauge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GaugeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gauge *GaugeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gauge.Contract.GaugeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gauge *GaugeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gauge.Contract.GaugeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gauge *GaugeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gauge.Contract.GaugeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gauge *GaugeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gauge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gauge *GaugeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gauge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gauge *GaugeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gauge.Contract.contract.Transact(opts, method, params...)
}

// ApprovedToDeposit is a free data retrieval call binding the contract method 0xe1522536.
//
// Solidity: function approved_to_deposit(address arg0, address arg1) view returns(bool)
func (_Gauge *GaugeCaller) ApprovedToDeposit(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "approved_to_deposit", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedToDeposit is a free data retrieval call binding the contract method 0xe1522536.
//
// Solidity: function approved_to_deposit(address arg0, address arg1) view returns(bool)
func (_Gauge *GaugeSession) ApprovedToDeposit(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Gauge.Contract.ApprovedToDeposit(&_Gauge.CallOpts, arg0, arg1)
}

// ApprovedToDeposit is a free data retrieval call binding the contract method 0xe1522536.
//
// Solidity: function approved_to_deposit(address arg0, address arg1) view returns(bool)
func (_Gauge *GaugeCallerSession) ApprovedToDeposit(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Gauge.Contract.ApprovedToDeposit(&_Gauge.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Gauge *GaugeCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Gauge *GaugeSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.BalanceOf(&_Gauge.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Gauge *GaugeCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.BalanceOf(&_Gauge.CallOpts, arg0)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Gauge *GaugeCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Gauge *GaugeSession) Controller() (common.Address, error) {
	return _Gauge.Contract.Controller(&_Gauge.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Gauge *GaugeCallerSession) Controller() (common.Address, error) {
	return _Gauge.Contract.Controller(&_Gauge.CallOpts)
}

// CrvToken is a free data retrieval call binding the contract method 0x76d8b117.
//
// Solidity: function crv_token() view returns(address)
func (_Gauge *GaugeCaller) CrvToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "crv_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrvToken is a free data retrieval call binding the contract method 0x76d8b117.
//
// Solidity: function crv_token() view returns(address)
func (_Gauge *GaugeSession) CrvToken() (common.Address, error) {
	return _Gauge.Contract.CrvToken(&_Gauge.CallOpts)
}

// CrvToken is a free data retrieval call binding the contract method 0x76d8b117.
//
// Solidity: function crv_token() view returns(address)
func (_Gauge *GaugeCallerSession) CrvToken() (common.Address, error) {
	return _Gauge.Contract.CrvToken(&_Gauge.CallOpts)
}

// FutureEpochTime is a free data retrieval call binding the contract method 0xbe5d1be9.
//
// Solidity: function future_epoch_time() view returns(uint256)
func (_Gauge *GaugeCaller) FutureEpochTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "future_epoch_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureEpochTime is a free data retrieval call binding the contract method 0xbe5d1be9.
//
// Solidity: function future_epoch_time() view returns(uint256)
func (_Gauge *GaugeSession) FutureEpochTime() (*big.Int, error) {
	return _Gauge.Contract.FutureEpochTime(&_Gauge.CallOpts)
}

// FutureEpochTime is a free data retrieval call binding the contract method 0xbe5d1be9.
//
// Solidity: function future_epoch_time() view returns(uint256)
func (_Gauge *GaugeCallerSession) FutureEpochTime() (*big.Int, error) {
	return _Gauge.Contract.FutureEpochTime(&_Gauge.CallOpts)
}

// InflationRate is a free data retrieval call binding the contract method 0x180692d0.
//
// Solidity: function inflation_rate() view returns(uint256)
func (_Gauge *GaugeCaller) InflationRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "inflation_rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InflationRate is a free data retrieval call binding the contract method 0x180692d0.
//
// Solidity: function inflation_rate() view returns(uint256)
func (_Gauge *GaugeSession) InflationRate() (*big.Int, error) {
	return _Gauge.Contract.InflationRate(&_Gauge.CallOpts)
}

// InflationRate is a free data retrieval call binding the contract method 0x180692d0.
//
// Solidity: function inflation_rate() view returns(uint256)
func (_Gauge *GaugeCallerSession) InflationRate() (*big.Int, error) {
	return _Gauge.Contract.InflationRate(&_Gauge.CallOpts)
}

// IntegrateCheckpoint is a free data retrieval call binding the contract method 0xd31f3f6d.
//
// Solidity: function integrate_checkpoint() view returns(uint256)
func (_Gauge *GaugeCaller) IntegrateCheckpoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "integrate_checkpoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateCheckpoint is a free data retrieval call binding the contract method 0xd31f3f6d.
//
// Solidity: function integrate_checkpoint() view returns(uint256)
func (_Gauge *GaugeSession) IntegrateCheckpoint() (*big.Int, error) {
	return _Gauge.Contract.IntegrateCheckpoint(&_Gauge.CallOpts)
}

// IntegrateCheckpoint is a free data retrieval call binding the contract method 0xd31f3f6d.
//
// Solidity: function integrate_checkpoint() view returns(uint256)
func (_Gauge *GaugeCallerSession) IntegrateCheckpoint() (*big.Int, error) {
	return _Gauge.Contract.IntegrateCheckpoint(&_Gauge.CallOpts)
}

// IntegrateCheckpointOf is a free data retrieval call binding the contract method 0x9bd324f2.
//
// Solidity: function integrate_checkpoint_of(address arg0) view returns(uint256)
func (_Gauge *GaugeCaller) IntegrateCheckpointOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "integrate_checkpoint_of", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateCheckpointOf is a free data retrieval call binding the contract method 0x9bd324f2.
//
// Solidity: function integrate_checkpoint_of(address arg0) view returns(uint256)
func (_Gauge *GaugeSession) IntegrateCheckpointOf(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.IntegrateCheckpointOf(&_Gauge.CallOpts, arg0)
}

// IntegrateCheckpointOf is a free data retrieval call binding the contract method 0x9bd324f2.
//
// Solidity: function integrate_checkpoint_of(address arg0) view returns(uint256)
func (_Gauge *GaugeCallerSession) IntegrateCheckpointOf(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.IntegrateCheckpointOf(&_Gauge.CallOpts, arg0)
}

// IntegrateFraction is a free data retrieval call binding the contract method 0x09400707.
//
// Solidity: function integrate_fraction(address arg0) view returns(uint256)
func (_Gauge *GaugeCaller) IntegrateFraction(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "integrate_fraction", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateFraction is a free data retrieval call binding the contract method 0x09400707.
//
// Solidity: function integrate_fraction(address arg0) view returns(uint256)
func (_Gauge *GaugeSession) IntegrateFraction(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.IntegrateFraction(&_Gauge.CallOpts, arg0)
}

// IntegrateFraction is a free data retrieval call binding the contract method 0x09400707.
//
// Solidity: function integrate_fraction(address arg0) view returns(uint256)
func (_Gauge *GaugeCallerSession) IntegrateFraction(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.IntegrateFraction(&_Gauge.CallOpts, arg0)
}

// IntegrateInvSupply is a free data retrieval call binding the contract method 0xfec8ee0c.
//
// Solidity: function integrate_inv_supply(uint256 arg0) view returns(uint256)
func (_Gauge *GaugeCaller) IntegrateInvSupply(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "integrate_inv_supply", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateInvSupply is a free data retrieval call binding the contract method 0xfec8ee0c.
//
// Solidity: function integrate_inv_supply(uint256 arg0) view returns(uint256)
func (_Gauge *GaugeSession) IntegrateInvSupply(arg0 *big.Int) (*big.Int, error) {
	return _Gauge.Contract.IntegrateInvSupply(&_Gauge.CallOpts, arg0)
}

// IntegrateInvSupply is a free data retrieval call binding the contract method 0xfec8ee0c.
//
// Solidity: function integrate_inv_supply(uint256 arg0) view returns(uint256)
func (_Gauge *GaugeCallerSession) IntegrateInvSupply(arg0 *big.Int) (*big.Int, error) {
	return _Gauge.Contract.IntegrateInvSupply(&_Gauge.CallOpts, arg0)
}

// IntegrateInvSupplyOf is a free data retrieval call binding the contract method 0xde263bfa.
//
// Solidity: function integrate_inv_supply_of(address arg0) view returns(uint256)
func (_Gauge *GaugeCaller) IntegrateInvSupplyOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "integrate_inv_supply_of", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateInvSupplyOf is a free data retrieval call binding the contract method 0xde263bfa.
//
// Solidity: function integrate_inv_supply_of(address arg0) view returns(uint256)
func (_Gauge *GaugeSession) IntegrateInvSupplyOf(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.IntegrateInvSupplyOf(&_Gauge.CallOpts, arg0)
}

// IntegrateInvSupplyOf is a free data retrieval call binding the contract method 0xde263bfa.
//
// Solidity: function integrate_inv_supply_of(address arg0) view returns(uint256)
func (_Gauge *GaugeCallerSession) IntegrateInvSupplyOf(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.IntegrateInvSupplyOf(&_Gauge.CallOpts, arg0)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_Gauge *GaugeCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "lp_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_Gauge *GaugeSession) LpToken() (common.Address, error) {
	return _Gauge.Contract.LpToken(&_Gauge.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_Gauge *GaugeCallerSession) LpToken() (common.Address, error) {
	return _Gauge.Contract.LpToken(&_Gauge.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Gauge *GaugeCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Gauge *GaugeSession) Minter() (common.Address, error) {
	return _Gauge.Contract.Minter(&_Gauge.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Gauge *GaugeCallerSession) Minter() (common.Address, error) {
	return _Gauge.Contract.Minter(&_Gauge.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(int128)
func (_Gauge *GaugeCaller) Period(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "period")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(int128)
func (_Gauge *GaugeSession) Period() (*big.Int, error) {
	return _Gauge.Contract.Period(&_Gauge.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(int128)
func (_Gauge *GaugeCallerSession) Period() (*big.Int, error) {
	return _Gauge.Contract.Period(&_Gauge.CallOpts)
}

// PeriodTimestamp is a free data retrieval call binding the contract method 0x7598108c.
//
// Solidity: function period_timestamp(uint256 arg0) view returns(uint256)
func (_Gauge *GaugeCaller) PeriodTimestamp(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "period_timestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodTimestamp is a free data retrieval call binding the contract method 0x7598108c.
//
// Solidity: function period_timestamp(uint256 arg0) view returns(uint256)
func (_Gauge *GaugeSession) PeriodTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _Gauge.Contract.PeriodTimestamp(&_Gauge.CallOpts, arg0)
}

// PeriodTimestamp is a free data retrieval call binding the contract method 0x7598108c.
//
// Solidity: function period_timestamp(uint256 arg0) view returns(uint256)
func (_Gauge *GaugeCallerSession) PeriodTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _Gauge.Contract.PeriodTimestamp(&_Gauge.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gauge *GaugeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gauge *GaugeSession) TotalSupply() (*big.Int, error) {
	return _Gauge.Contract.TotalSupply(&_Gauge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gauge *GaugeCallerSession) TotalSupply() (*big.Int, error) {
	return _Gauge.Contract.TotalSupply(&_Gauge.CallOpts)
}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_Gauge *GaugeCaller) VotingEscrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "voting_escrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_Gauge *GaugeSession) VotingEscrow() (common.Address, error) {
	return _Gauge.Contract.VotingEscrow(&_Gauge.CallOpts)
}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_Gauge *GaugeCallerSession) VotingEscrow() (common.Address, error) {
	return _Gauge.Contract.VotingEscrow(&_Gauge.CallOpts)
}

// WorkingBalances is a free data retrieval call binding the contract method 0x13ecb1ca.
//
// Solidity: function working_balances(address arg0) view returns(uint256)
func (_Gauge *GaugeCaller) WorkingBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "working_balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WorkingBalances is a free data retrieval call binding the contract method 0x13ecb1ca.
//
// Solidity: function working_balances(address arg0) view returns(uint256)
func (_Gauge *GaugeSession) WorkingBalances(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.WorkingBalances(&_Gauge.CallOpts, arg0)
}

// WorkingBalances is a free data retrieval call binding the contract method 0x13ecb1ca.
//
// Solidity: function working_balances(address arg0) view returns(uint256)
func (_Gauge *GaugeCallerSession) WorkingBalances(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.WorkingBalances(&_Gauge.CallOpts, arg0)
}

// WorkingSupply is a free data retrieval call binding the contract method 0x17e28089.
//
// Solidity: function working_supply() view returns(uint256)
func (_Gauge *GaugeCaller) WorkingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "working_supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WorkingSupply is a free data retrieval call binding the contract method 0x17e28089.
//
// Solidity: function working_supply() view returns(uint256)
func (_Gauge *GaugeSession) WorkingSupply() (*big.Int, error) {
	return _Gauge.Contract.WorkingSupply(&_Gauge.CallOpts)
}

// WorkingSupply is a free data retrieval call binding the contract method 0x17e28089.
//
// Solidity: function working_supply() view returns(uint256)
func (_Gauge *GaugeCallerSession) WorkingSupply() (*big.Int, error) {
	return _Gauge.Contract.WorkingSupply(&_Gauge.CallOpts)
}

// ClaimableTokens is a paid mutator transaction binding the contract method 0x33134583.
//
// Solidity: function claimable_tokens(address addr) returns(uint256)
func (_Gauge *GaugeTransactor) ClaimableTokens(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "claimable_tokens", addr)
}

// ClaimableTokens is a paid mutator transaction binding the contract method 0x33134583.
//
// Solidity: function claimable_tokens(address addr) returns(uint256)
func (_Gauge *GaugeSession) ClaimableTokens(addr common.Address) (*types.Transaction, error) {
	return _Gauge.Contract.ClaimableTokens(&_Gauge.TransactOpts, addr)
}

// ClaimableTokens is a paid mutator transaction binding the contract method 0x33134583.
//
// Solidity: function claimable_tokens(address addr) returns(uint256)
func (_Gauge *GaugeTransactorSession) ClaimableTokens(addr common.Address) (*types.Transaction, error) {
	return _Gauge.Contract.ClaimableTokens(&_Gauge.TransactOpts, addr)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _value) returns()
func (_Gauge *GaugeTransactor) Deposit(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "deposit", _value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _value) returns()
func (_Gauge *GaugeSession) Deposit(_value *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.Deposit(&_Gauge.TransactOpts, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _value) returns()
func (_Gauge *GaugeTransactorSession) Deposit(_value *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.Deposit(&_Gauge.TransactOpts, _value)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _value, address addr) returns()
func (_Gauge *GaugeTransactor) Deposit0(opts *bind.TransactOpts, _value *big.Int, addr common.Address) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "deposit0", _value, addr)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _value, address addr) returns()
func (_Gauge *GaugeSession) Deposit0(_value *big.Int, addr common.Address) (*types.Transaction, error) {
	return _Gauge.Contract.Deposit0(&_Gauge.TransactOpts, _value, addr)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _value, address addr) returns()
func (_Gauge *GaugeTransactorSession) Deposit0(_value *big.Int, addr common.Address) (*types.Transaction, error) {
	return _Gauge.Contract.Deposit0(&_Gauge.TransactOpts, _value, addr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address addr) returns()
func (_Gauge *GaugeTransactor) Kick(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "kick", addr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address addr) returns()
func (_Gauge *GaugeSession) Kick(addr common.Address) (*types.Transaction, error) {
	return _Gauge.Contract.Kick(&_Gauge.TransactOpts, addr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address addr) returns()
func (_Gauge *GaugeTransactorSession) Kick(addr common.Address) (*types.Transaction, error) {
	return _Gauge.Contract.Kick(&_Gauge.TransactOpts, addr)
}

// SetApproveDeposit is a paid mutator transaction binding the contract method 0x1d2747d4.
//
// Solidity: function set_approve_deposit(address addr, bool can_deposit) returns()
func (_Gauge *GaugeTransactor) SetApproveDeposit(opts *bind.TransactOpts, addr common.Address, can_deposit bool) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "set_approve_deposit", addr, can_deposit)
}

// SetApproveDeposit is a paid mutator transaction binding the contract method 0x1d2747d4.
//
// Solidity: function set_approve_deposit(address addr, bool can_deposit) returns()
func (_Gauge *GaugeSession) SetApproveDeposit(addr common.Address, can_deposit bool) (*types.Transaction, error) {
	return _Gauge.Contract.SetApproveDeposit(&_Gauge.TransactOpts, addr, can_deposit)
}

// SetApproveDeposit is a paid mutator transaction binding the contract method 0x1d2747d4.
//
// Solidity: function set_approve_deposit(address addr, bool can_deposit) returns()
func (_Gauge *GaugeTransactorSession) SetApproveDeposit(addr common.Address, can_deposit bool) (*types.Transaction, error) {
	return _Gauge.Contract.SetApproveDeposit(&_Gauge.TransactOpts, addr, can_deposit)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0x4b820093.
//
// Solidity: function user_checkpoint(address addr) returns(bool)
func (_Gauge *GaugeTransactor) UserCheckpoint(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "user_checkpoint", addr)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0x4b820093.
//
// Solidity: function user_checkpoint(address addr) returns(bool)
func (_Gauge *GaugeSession) UserCheckpoint(addr common.Address) (*types.Transaction, error) {
	return _Gauge.Contract.UserCheckpoint(&_Gauge.TransactOpts, addr)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0x4b820093.
//
// Solidity: function user_checkpoint(address addr) returns(bool)
func (_Gauge *GaugeTransactorSession) UserCheckpoint(addr common.Address) (*types.Transaction, error) {
	return _Gauge.Contract.UserCheckpoint(&_Gauge.TransactOpts, addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_Gauge *GaugeTransactor) Withdraw(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "withdraw", _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_Gauge *GaugeSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.Withdraw(&_Gauge.TransactOpts, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_Gauge *GaugeTransactorSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.Withdraw(&_Gauge.TransactOpts, _value)
}

// GaugeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Gauge contract.
type GaugeDepositIterator struct {
	Event *GaugeDeposit // Event containing the contract specifics and raw log

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
func (it *GaugeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GaugeDeposit)
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
		it.Event = new(GaugeDeposit)
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
func (it *GaugeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GaugeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GaugeDeposit represents a Deposit event raised by the Gauge contract.
type GaugeDeposit struct {
	Provider common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed provider, uint256 value)
func (_Gauge *GaugeFilterer) FilterDeposit(opts *bind.FilterOpts, provider []common.Address) (*GaugeDepositIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Gauge.contract.FilterLogs(opts, "Deposit", providerRule)
	if err != nil {
		return nil, err
	}
	return &GaugeDepositIterator{contract: _Gauge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed provider, uint256 value)
func (_Gauge *GaugeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *GaugeDeposit, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Gauge.contract.WatchLogs(opts, "Deposit", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GaugeDeposit)
				if err := _Gauge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed provider, uint256 value)
func (_Gauge *GaugeFilterer) ParseDeposit(log types.Log) (*GaugeDeposit, error) {
	event := new(GaugeDeposit)
	if err := _Gauge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GaugeUpdateLiquidityLimitIterator is returned from FilterUpdateLiquidityLimit and is used to iterate over the raw logs and unpacked data for UpdateLiquidityLimit events raised by the Gauge contract.
type GaugeUpdateLiquidityLimitIterator struct {
	Event *GaugeUpdateLiquidityLimit // Event containing the contract specifics and raw log

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
func (it *GaugeUpdateLiquidityLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GaugeUpdateLiquidityLimit)
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
		it.Event = new(GaugeUpdateLiquidityLimit)
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
func (it *GaugeUpdateLiquidityLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GaugeUpdateLiquidityLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GaugeUpdateLiquidityLimit represents a UpdateLiquidityLimit event raised by the Gauge contract.
type GaugeUpdateLiquidityLimit struct {
	User            common.Address
	OriginalBalance *big.Int
	OriginalSupply  *big.Int
	WorkingBalance  *big.Int
	WorkingSupply   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateLiquidityLimit is a free log retrieval operation binding the contract event 0x7ecd84343f76a23d2227290e0288da3251b045541698e575a5515af4f04197a3.
//
// Solidity: event UpdateLiquidityLimit(address user, uint256 original_balance, uint256 original_supply, uint256 working_balance, uint256 working_supply)
func (_Gauge *GaugeFilterer) FilterUpdateLiquidityLimit(opts *bind.FilterOpts) (*GaugeUpdateLiquidityLimitIterator, error) {

	logs, sub, err := _Gauge.contract.FilterLogs(opts, "UpdateLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return &GaugeUpdateLiquidityLimitIterator{contract: _Gauge.contract, event: "UpdateLiquidityLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateLiquidityLimit is a free log subscription operation binding the contract event 0x7ecd84343f76a23d2227290e0288da3251b045541698e575a5515af4f04197a3.
//
// Solidity: event UpdateLiquidityLimit(address user, uint256 original_balance, uint256 original_supply, uint256 working_balance, uint256 working_supply)
func (_Gauge *GaugeFilterer) WatchUpdateLiquidityLimit(opts *bind.WatchOpts, sink chan<- *GaugeUpdateLiquidityLimit) (event.Subscription, error) {

	logs, sub, err := _Gauge.contract.WatchLogs(opts, "UpdateLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GaugeUpdateLiquidityLimit)
				if err := _Gauge.contract.UnpackLog(event, "UpdateLiquidityLimit", log); err != nil {
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

// ParseUpdateLiquidityLimit is a log parse operation binding the contract event 0x7ecd84343f76a23d2227290e0288da3251b045541698e575a5515af4f04197a3.
//
// Solidity: event UpdateLiquidityLimit(address user, uint256 original_balance, uint256 original_supply, uint256 working_balance, uint256 working_supply)
func (_Gauge *GaugeFilterer) ParseUpdateLiquidityLimit(log types.Log) (*GaugeUpdateLiquidityLimit, error) {
	event := new(GaugeUpdateLiquidityLimit)
	if err := _Gauge.contract.UnpackLog(event, "UpdateLiquidityLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GaugeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Gauge contract.
type GaugeWithdrawIterator struct {
	Event *GaugeWithdraw // Event containing the contract specifics and raw log

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
func (it *GaugeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GaugeWithdraw)
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
		it.Event = new(GaugeWithdraw)
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
func (it *GaugeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GaugeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GaugeWithdraw represents a Withdraw event raised by the Gauge contract.
type GaugeWithdraw struct {
	Provider common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed provider, uint256 value)
func (_Gauge *GaugeFilterer) FilterWithdraw(opts *bind.FilterOpts, provider []common.Address) (*GaugeWithdrawIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Gauge.contract.FilterLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return &GaugeWithdrawIterator{contract: _Gauge.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed provider, uint256 value)
func (_Gauge *GaugeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *GaugeWithdraw, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Gauge.contract.WatchLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GaugeWithdraw)
				if err := _Gauge.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed provider, uint256 value)
func (_Gauge *GaugeFilterer) ParseWithdraw(log types.Log) (*GaugeWithdraw, error) {
	event := new(GaugeWithdraw)
	if err := _Gauge.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
