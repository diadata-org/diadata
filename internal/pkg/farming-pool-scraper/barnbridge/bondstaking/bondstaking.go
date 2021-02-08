// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bondstaking

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

// BondstakingABI is the input ABI used to generate the binding from.
const BondstakingABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch1Start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epochDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"ManualEpochInit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prevBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"prevMultiplier\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"currentMultiplier\",\"type\":\"uint128\"}],\"name\":\"computeNewMultiplier\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpochMultiplier\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch1Start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"epochIsInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getEpochPoolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getEpochUserBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"manualEpochInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Bondstaking is an auto generated Go binding around an Ethereum contract.
type Bondstaking struct {
	BondstakingCaller     // Read-only binding to the contract
	BondstakingTransactor // Write-only binding to the contract
	BondstakingFilterer   // Log filterer for contract events
}

// BondstakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BondstakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondstakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BondstakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondstakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BondstakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondstakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BondstakingSession struct {
	Contract     *Bondstaking      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BondstakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BondstakingCallerSession struct {
	Contract *BondstakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BondstakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BondstakingTransactorSession struct {
	Contract     *BondstakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BondstakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BondstakingRaw struct {
	Contract *Bondstaking // Generic contract binding to access the raw methods on
}

// BondstakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BondstakingCallerRaw struct {
	Contract *BondstakingCaller // Generic read-only contract binding to access the raw methods on
}

// BondstakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BondstakingTransactorRaw struct {
	Contract *BondstakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBondstaking creates a new instance of Bondstaking, bound to a specific deployed contract.
func NewBondstaking(address common.Address, backend bind.ContractBackend) (*Bondstaking, error) {
	contract, err := bindBondstaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bondstaking{BondstakingCaller: BondstakingCaller{contract: contract}, BondstakingTransactor: BondstakingTransactor{contract: contract}, BondstakingFilterer: BondstakingFilterer{contract: contract}}, nil
}

// NewBondstakingCaller creates a new read-only instance of Bondstaking, bound to a specific deployed contract.
func NewBondstakingCaller(address common.Address, caller bind.ContractCaller) (*BondstakingCaller, error) {
	contract, err := bindBondstaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BondstakingCaller{contract: contract}, nil
}

// NewBondstakingTransactor creates a new write-only instance of Bondstaking, bound to a specific deployed contract.
func NewBondstakingTransactor(address common.Address, transactor bind.ContractTransactor) (*BondstakingTransactor, error) {
	contract, err := bindBondstaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BondstakingTransactor{contract: contract}, nil
}

// NewBondstakingFilterer creates a new log filterer instance of Bondstaking, bound to a specific deployed contract.
func NewBondstakingFilterer(address common.Address, filterer bind.ContractFilterer) (*BondstakingFilterer, error) {
	contract, err := bindBondstaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BondstakingFilterer{contract: contract}, nil
}

// bindBondstaking binds a generic wrapper to an already deployed contract.
func bindBondstaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BondstakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bondstaking *BondstakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bondstaking.Contract.BondstakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bondstaking *BondstakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bondstaking.Contract.BondstakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bondstaking *BondstakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bondstaking.Contract.BondstakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bondstaking *BondstakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bondstaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bondstaking *BondstakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bondstaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bondstaking *BondstakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bondstaking.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address user, address token) view returns(uint256)
func (_Bondstaking *BondstakingCaller) BalanceOf(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bondstaking.contract.Call(opts, &out, "balanceOf", user, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address user, address token) view returns(uint256)
func (_Bondstaking *BondstakingSession) BalanceOf(user common.Address, token common.Address) (*big.Int, error) {
	return _Bondstaking.Contract.BalanceOf(&_Bondstaking.CallOpts, user, token)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address user, address token) view returns(uint256)
func (_Bondstaking *BondstakingCallerSession) BalanceOf(user common.Address, token common.Address) (*big.Int, error) {
	return _Bondstaking.Contract.BalanceOf(&_Bondstaking.CallOpts, user, token)
}

// ComputeNewMultiplier is a free data retrieval call binding the contract method 0x4be41dba.
//
// Solidity: function computeNewMultiplier(uint256 prevBalance, uint128 prevMultiplier, uint256 amount, uint128 currentMultiplier) pure returns(uint128)
func (_Bondstaking *BondstakingCaller) ComputeNewMultiplier(opts *bind.CallOpts, prevBalance *big.Int, prevMultiplier *big.Int, amount *big.Int, currentMultiplier *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bondstaking.contract.Call(opts, &out, "computeNewMultiplier", prevBalance, prevMultiplier, amount, currentMultiplier)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeNewMultiplier is a free data retrieval call binding the contract method 0x4be41dba.
//
// Solidity: function computeNewMultiplier(uint256 prevBalance, uint128 prevMultiplier, uint256 amount, uint128 currentMultiplier) pure returns(uint128)
func (_Bondstaking *BondstakingSession) ComputeNewMultiplier(prevBalance *big.Int, prevMultiplier *big.Int, amount *big.Int, currentMultiplier *big.Int) (*big.Int, error) {
	return _Bondstaking.Contract.ComputeNewMultiplier(&_Bondstaking.CallOpts, prevBalance, prevMultiplier, amount, currentMultiplier)
}

// ComputeNewMultiplier is a free data retrieval call binding the contract method 0x4be41dba.
//
// Solidity: function computeNewMultiplier(uint256 prevBalance, uint128 prevMultiplier, uint256 amount, uint128 currentMultiplier) pure returns(uint128)
func (_Bondstaking *BondstakingCallerSession) ComputeNewMultiplier(prevBalance *big.Int, prevMultiplier *big.Int, amount *big.Int, currentMultiplier *big.Int) (*big.Int, error) {
	return _Bondstaking.Contract.ComputeNewMultiplier(&_Bondstaking.CallOpts, prevBalance, prevMultiplier, amount, currentMultiplier)
}

// CurrentEpochMultiplier is a free data retrieval call binding the contract method 0xce58a2a8.
//
// Solidity: function currentEpochMultiplier() view returns(uint128)
func (_Bondstaking *BondstakingCaller) CurrentEpochMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bondstaking.contract.Call(opts, &out, "currentEpochMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpochMultiplier is a free data retrieval call binding the contract method 0xce58a2a8.
//
// Solidity: function currentEpochMultiplier() view returns(uint128)
func (_Bondstaking *BondstakingSession) CurrentEpochMultiplier() (*big.Int, error) {
	return _Bondstaking.Contract.CurrentEpochMultiplier(&_Bondstaking.CallOpts)
}

// CurrentEpochMultiplier is a free data retrieval call binding the contract method 0xce58a2a8.
//
// Solidity: function currentEpochMultiplier() view returns(uint128)
func (_Bondstaking *BondstakingCallerSession) CurrentEpochMultiplier() (*big.Int, error) {
	return _Bondstaking.Contract.CurrentEpochMultiplier(&_Bondstaking.CallOpts)
}

// Epoch1Start is a free data retrieval call binding the contract method 0xf4a4341d.
//
// Solidity: function epoch1Start() view returns(uint256)
func (_Bondstaking *BondstakingCaller) Epoch1Start(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bondstaking.contract.Call(opts, &out, "epoch1Start")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch1Start is a free data retrieval call binding the contract method 0xf4a4341d.
//
// Solidity: function epoch1Start() view returns(uint256)
func (_Bondstaking *BondstakingSession) Epoch1Start() (*big.Int, error) {
	return _Bondstaking.Contract.Epoch1Start(&_Bondstaking.CallOpts)
}

// Epoch1Start is a free data retrieval call binding the contract method 0xf4a4341d.
//
// Solidity: function epoch1Start() view returns(uint256)
func (_Bondstaking *BondstakingCallerSession) Epoch1Start() (*big.Int, error) {
	return _Bondstaking.Contract.Epoch1Start(&_Bondstaking.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Bondstaking *BondstakingCaller) EpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bondstaking.contract.Call(opts, &out, "epochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Bondstaking *BondstakingSession) EpochDuration() (*big.Int, error) {
	return _Bondstaking.Contract.EpochDuration(&_Bondstaking.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Bondstaking *BondstakingCallerSession) EpochDuration() (*big.Int, error) {
	return _Bondstaking.Contract.EpochDuration(&_Bondstaking.CallOpts)
}

// EpochIsInitialized is a free data retrieval call binding the contract method 0xaa579154.
//
// Solidity: function epochIsInitialized(address token, uint128 epochId) view returns(bool)
func (_Bondstaking *BondstakingCaller) EpochIsInitialized(opts *bind.CallOpts, token common.Address, epochId *big.Int) (bool, error) {
	var out []interface{}
	err := _Bondstaking.contract.Call(opts, &out, "epochIsInitialized", token, epochId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EpochIsInitialized is a free data retrieval call binding the contract method 0xaa579154.
//
// Solidity: function epochIsInitialized(address token, uint128 epochId) view returns(bool)
func (_Bondstaking *BondstakingSession) EpochIsInitialized(token common.Address, epochId *big.Int) (bool, error) {
	return _Bondstaking.Contract.EpochIsInitialized(&_Bondstaking.CallOpts, token, epochId)
}

// EpochIsInitialized is a free data retrieval call binding the contract method 0xaa579154.
//
// Solidity: function epochIsInitialized(address token, uint128 epochId) view returns(bool)
func (_Bondstaking *BondstakingCallerSession) EpochIsInitialized(token common.Address, epochId *big.Int) (bool, error) {
	return _Bondstaking.Contract.EpochIsInitialized(&_Bondstaking.CallOpts, token, epochId)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint128)
func (_Bondstaking *BondstakingCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bondstaking.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint128)
func (_Bondstaking *BondstakingSession) GetCurrentEpoch() (*big.Int, error) {
	return _Bondstaking.Contract.GetCurrentEpoch(&_Bondstaking.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint128)
func (_Bondstaking *BondstakingCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _Bondstaking.Contract.GetCurrentEpoch(&_Bondstaking.CallOpts)
}

// GetEpochPoolSize is a free data retrieval call binding the contract method 0x2ca32d7e.
//
// Solidity: function getEpochPoolSize(address tokenAddress, uint128 epochId) view returns(uint256)
func (_Bondstaking *BondstakingCaller) GetEpochPoolSize(opts *bind.CallOpts, tokenAddress common.Address, epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bondstaking.contract.Call(opts, &out, "getEpochPoolSize", tokenAddress, epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochPoolSize is a free data retrieval call binding the contract method 0x2ca32d7e.
//
// Solidity: function getEpochPoolSize(address tokenAddress, uint128 epochId) view returns(uint256)
func (_Bondstaking *BondstakingSession) GetEpochPoolSize(tokenAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _Bondstaking.Contract.GetEpochPoolSize(&_Bondstaking.CallOpts, tokenAddress, epochId)
}

// GetEpochPoolSize is a free data retrieval call binding the contract method 0x2ca32d7e.
//
// Solidity: function getEpochPoolSize(address tokenAddress, uint128 epochId) view returns(uint256)
func (_Bondstaking *BondstakingCallerSession) GetEpochPoolSize(tokenAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _Bondstaking.Contract.GetEpochPoolSize(&_Bondstaking.CallOpts, tokenAddress, epochId)
}

// GetEpochUserBalance is a free data retrieval call binding the contract method 0x8c028dd0.
//
// Solidity: function getEpochUserBalance(address user, address token, uint128 epochId) view returns(uint256)
func (_Bondstaking *BondstakingCaller) GetEpochUserBalance(opts *bind.CallOpts, user common.Address, token common.Address, epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bondstaking.contract.Call(opts, &out, "getEpochUserBalance", user, token, epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochUserBalance is a free data retrieval call binding the contract method 0x8c028dd0.
//
// Solidity: function getEpochUserBalance(address user, address token, uint128 epochId) view returns(uint256)
func (_Bondstaking *BondstakingSession) GetEpochUserBalance(user common.Address, token common.Address, epochId *big.Int) (*big.Int, error) {
	return _Bondstaking.Contract.GetEpochUserBalance(&_Bondstaking.CallOpts, user, token, epochId)
}

// GetEpochUserBalance is a free data retrieval call binding the contract method 0x8c028dd0.
//
// Solidity: function getEpochUserBalance(address user, address token, uint128 epochId) view returns(uint256)
func (_Bondstaking *BondstakingCallerSession) GetEpochUserBalance(user common.Address, token common.Address, epochId *big.Int) (*big.Int, error) {
	return _Bondstaking.Contract.GetEpochUserBalance(&_Bondstaking.CallOpts, user, token, epochId)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address tokenAddress, uint256 amount) returns()
func (_Bondstaking *BondstakingTransactor) Deposit(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bondstaking.contract.Transact(opts, "deposit", tokenAddress, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address tokenAddress, uint256 amount) returns()
func (_Bondstaking *BondstakingSession) Deposit(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bondstaking.Contract.Deposit(&_Bondstaking.TransactOpts, tokenAddress, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address tokenAddress, uint256 amount) returns()
func (_Bondstaking *BondstakingTransactorSession) Deposit(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bondstaking.Contract.Deposit(&_Bondstaking.TransactOpts, tokenAddress, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address tokenAddress) returns()
func (_Bondstaking *BondstakingTransactor) EmergencyWithdraw(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bondstaking.contract.Transact(opts, "emergencyWithdraw", tokenAddress)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address tokenAddress) returns()
func (_Bondstaking *BondstakingSession) EmergencyWithdraw(tokenAddress common.Address) (*types.Transaction, error) {
	return _Bondstaking.Contract.EmergencyWithdraw(&_Bondstaking.TransactOpts, tokenAddress)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address tokenAddress) returns()
func (_Bondstaking *BondstakingTransactorSession) EmergencyWithdraw(tokenAddress common.Address) (*types.Transaction, error) {
	return _Bondstaking.Contract.EmergencyWithdraw(&_Bondstaking.TransactOpts, tokenAddress)
}

// ManualEpochInit is a paid mutator transaction binding the contract method 0xea2c38ae.
//
// Solidity: function manualEpochInit(address[] tokens, uint128 epochId) returns()
func (_Bondstaking *BondstakingTransactor) ManualEpochInit(opts *bind.TransactOpts, tokens []common.Address, epochId *big.Int) (*types.Transaction, error) {
	return _Bondstaking.contract.Transact(opts, "manualEpochInit", tokens, epochId)
}

// ManualEpochInit is a paid mutator transaction binding the contract method 0xea2c38ae.
//
// Solidity: function manualEpochInit(address[] tokens, uint128 epochId) returns()
func (_Bondstaking *BondstakingSession) ManualEpochInit(tokens []common.Address, epochId *big.Int) (*types.Transaction, error) {
	return _Bondstaking.Contract.ManualEpochInit(&_Bondstaking.TransactOpts, tokens, epochId)
}

// ManualEpochInit is a paid mutator transaction binding the contract method 0xea2c38ae.
//
// Solidity: function manualEpochInit(address[] tokens, uint128 epochId) returns()
func (_Bondstaking *BondstakingTransactorSession) ManualEpochInit(tokens []common.Address, epochId *big.Int) (*types.Transaction, error) {
	return _Bondstaking.Contract.ManualEpochInit(&_Bondstaking.TransactOpts, tokens, epochId)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address tokenAddress, uint256 amount) returns()
func (_Bondstaking *BondstakingTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bondstaking.contract.Transact(opts, "withdraw", tokenAddress, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address tokenAddress, uint256 amount) returns()
func (_Bondstaking *BondstakingSession) Withdraw(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bondstaking.Contract.Withdraw(&_Bondstaking.TransactOpts, tokenAddress, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address tokenAddress, uint256 amount) returns()
func (_Bondstaking *BondstakingTransactorSession) Withdraw(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bondstaking.Contract.Withdraw(&_Bondstaking.TransactOpts, tokenAddress, amount)
}

// BondstakingDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Bondstaking contract.
type BondstakingDepositIterator struct {
	Event *BondstakingDeposit // Event containing the contract specifics and raw log

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
func (it *BondstakingDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondstakingDeposit)
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
		it.Event = new(BondstakingDeposit)
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
func (it *BondstakingDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondstakingDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondstakingDeposit represents a Deposit event raised by the Bondstaking contract.
type BondstakingDeposit struct {
	User         common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Bondstaking *BondstakingFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, tokenAddress []common.Address) (*BondstakingDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Bondstaking.contract.FilterLogs(opts, "Deposit", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BondstakingDepositIterator{contract: _Bondstaking.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Bondstaking *BondstakingFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BondstakingDeposit, user []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Bondstaking.contract.WatchLogs(opts, "Deposit", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondstakingDeposit)
				if err := _Bondstaking.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Bondstaking *BondstakingFilterer) ParseDeposit(log types.Log) (*BondstakingDeposit, error) {
	event := new(BondstakingDeposit)
	if err := _Bondstaking.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondstakingEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Bondstaking contract.
type BondstakingEmergencyWithdrawIterator struct {
	Event *BondstakingEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *BondstakingEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondstakingEmergencyWithdraw)
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
		it.Event = new(BondstakingEmergencyWithdraw)
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
func (it *BondstakingEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondstakingEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondstakingEmergencyWithdraw represents a EmergencyWithdraw event raised by the Bondstaking contract.
type BondstakingEmergencyWithdraw struct {
	User         common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xf24ef89f38eadc1bde50701ad6e4d6d11a2dc24f7cf834a486991f3883328504.
//
// Solidity: event EmergencyWithdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Bondstaking *BondstakingFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, tokenAddress []common.Address) (*BondstakingEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Bondstaking.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BondstakingEmergencyWithdrawIterator{contract: _Bondstaking.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xf24ef89f38eadc1bde50701ad6e4d6d11a2dc24f7cf834a486991f3883328504.
//
// Solidity: event EmergencyWithdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Bondstaking *BondstakingFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *BondstakingEmergencyWithdraw, user []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Bondstaking.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondstakingEmergencyWithdraw)
				if err := _Bondstaking.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xf24ef89f38eadc1bde50701ad6e4d6d11a2dc24f7cf834a486991f3883328504.
//
// Solidity: event EmergencyWithdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Bondstaking *BondstakingFilterer) ParseEmergencyWithdraw(log types.Log) (*BondstakingEmergencyWithdraw, error) {
	event := new(BondstakingEmergencyWithdraw)
	if err := _Bondstaking.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondstakingManualEpochInitIterator is returned from FilterManualEpochInit and is used to iterate over the raw logs and unpacked data for ManualEpochInit events raised by the Bondstaking contract.
type BondstakingManualEpochInitIterator struct {
	Event *BondstakingManualEpochInit // Event containing the contract specifics and raw log

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
func (it *BondstakingManualEpochInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondstakingManualEpochInit)
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
		it.Event = new(BondstakingManualEpochInit)
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
func (it *BondstakingManualEpochInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondstakingManualEpochInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondstakingManualEpochInit represents a ManualEpochInit event raised by the Bondstaking contract.
type BondstakingManualEpochInit struct {
	Caller  common.Address
	EpochId *big.Int
	Tokens  []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManualEpochInit is a free log retrieval operation binding the contract event 0xb85c32b8d9cecc81feba78646289584a693e9a8afea40ab2fd31efae4408429f.
//
// Solidity: event ManualEpochInit(address indexed caller, uint128 indexed epochId, address[] tokens)
func (_Bondstaking *BondstakingFilterer) FilterManualEpochInit(opts *bind.FilterOpts, caller []common.Address, epochId []*big.Int) (*BondstakingManualEpochInitIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Bondstaking.contract.FilterLogs(opts, "ManualEpochInit", callerRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &BondstakingManualEpochInitIterator{contract: _Bondstaking.contract, event: "ManualEpochInit", logs: logs, sub: sub}, nil
}

// WatchManualEpochInit is a free log subscription operation binding the contract event 0xb85c32b8d9cecc81feba78646289584a693e9a8afea40ab2fd31efae4408429f.
//
// Solidity: event ManualEpochInit(address indexed caller, uint128 indexed epochId, address[] tokens)
func (_Bondstaking *BondstakingFilterer) WatchManualEpochInit(opts *bind.WatchOpts, sink chan<- *BondstakingManualEpochInit, caller []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Bondstaking.contract.WatchLogs(opts, "ManualEpochInit", callerRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondstakingManualEpochInit)
				if err := _Bondstaking.contract.UnpackLog(event, "ManualEpochInit", log); err != nil {
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

// ParseManualEpochInit is a log parse operation binding the contract event 0xb85c32b8d9cecc81feba78646289584a693e9a8afea40ab2fd31efae4408429f.
//
// Solidity: event ManualEpochInit(address indexed caller, uint128 indexed epochId, address[] tokens)
func (_Bondstaking *BondstakingFilterer) ParseManualEpochInit(log types.Log) (*BondstakingManualEpochInit, error) {
	event := new(BondstakingManualEpochInit)
	if err := _Bondstaking.contract.UnpackLog(event, "ManualEpochInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondstakingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Bondstaking contract.
type BondstakingWithdrawIterator struct {
	Event *BondstakingWithdraw // Event containing the contract specifics and raw log

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
func (it *BondstakingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondstakingWithdraw)
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
		it.Event = new(BondstakingWithdraw)
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
func (it *BondstakingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondstakingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondstakingWithdraw represents a Withdraw event raised by the Bondstaking contract.
type BondstakingWithdraw struct {
	User         common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Bondstaking *BondstakingFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, tokenAddress []common.Address) (*BondstakingWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Bondstaking.contract.FilterLogs(opts, "Withdraw", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BondstakingWithdrawIterator{contract: _Bondstaking.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Bondstaking *BondstakingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BondstakingWithdraw, user []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Bondstaking.contract.WatchLogs(opts, "Withdraw", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondstakingWithdraw)
				if err := _Bondstaking.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Bondstaking *BondstakingFilterer) ParseWithdraw(log types.Log) (*BondstakingWithdraw, error) {
	event := new(BondstakingWithdraw)
	if err := _Bondstaking.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
