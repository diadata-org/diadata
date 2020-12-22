// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cvaultpoolcontract

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

// CvaultpoolcontractABI is the input ABI used to generate the binding from.
const CvaultpoolcontractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"SuperAdminTransfered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_withdrawable\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_\",\"type\":\"uint256\"}],\"name\":\"addPendingRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"averageFeesPerBlockEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"averagePerBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"averageFeesPerBlockSinceStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"averagePerBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnSuperAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"core\",\"outputs\":[{\"internalType\":\"contractINBUNIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cumulativeRewardsSinceStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositFor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devaddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochCalculationStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINBUNIERC20\",\"name\":\"_core\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_devaddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"superAdmin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"newSuperAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingCore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accCorePerShare\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"withdrawable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsInThisEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setAllowanceForPoolToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_DEV_FEE\",\"type\":\"uint16\"}],\"name\":\"setDevFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_devaddr\",\"type\":\"address\"}],\"name\":\"setDevFeeReciever\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withdrawable\",\"type\":\"bool\"}],\"name\":\"setPoolWithdrawable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setStrategyContractOrDistributionContractAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startNewEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Cvaultpoolcontract is an auto generated Go binding around an Ethereum contract.
type Cvaultpoolcontract struct {
	CvaultpoolcontractCaller     // Read-only binding to the contract
	CvaultpoolcontractTransactor // Write-only binding to the contract
	CvaultpoolcontractFilterer   // Log filterer for contract events
}

// CvaultpoolcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CvaultpoolcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CvaultpoolcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CvaultpoolcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CvaultpoolcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CvaultpoolcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CvaultpoolcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CvaultpoolcontractSession struct {
	Contract     *Cvaultpoolcontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CvaultpoolcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CvaultpoolcontractCallerSession struct {
	Contract *CvaultpoolcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CvaultpoolcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CvaultpoolcontractTransactorSession struct {
	Contract     *CvaultpoolcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CvaultpoolcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CvaultpoolcontractRaw struct {
	Contract *Cvaultpoolcontract // Generic contract binding to access the raw methods on
}

// CvaultpoolcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CvaultpoolcontractCallerRaw struct {
	Contract *CvaultpoolcontractCaller // Generic read-only contract binding to access the raw methods on
}

// CvaultpoolcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CvaultpoolcontractTransactorRaw struct {
	Contract *CvaultpoolcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCvaultpoolcontract creates a new instance of Cvaultpoolcontract, bound to a specific deployed contract.
func NewCvaultpoolcontract(address common.Address, backend bind.ContractBackend) (*Cvaultpoolcontract, error) {
	contract, err := bindCvaultpoolcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cvaultpoolcontract{CvaultpoolcontractCaller: CvaultpoolcontractCaller{contract: contract}, CvaultpoolcontractTransactor: CvaultpoolcontractTransactor{contract: contract}, CvaultpoolcontractFilterer: CvaultpoolcontractFilterer{contract: contract}}, nil
}

// NewCvaultpoolcontractCaller creates a new read-only instance of Cvaultpoolcontract, bound to a specific deployed contract.
func NewCvaultpoolcontractCaller(address common.Address, caller bind.ContractCaller) (*CvaultpoolcontractCaller, error) {
	contract, err := bindCvaultpoolcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CvaultpoolcontractCaller{contract: contract}, nil
}

// NewCvaultpoolcontractTransactor creates a new write-only instance of Cvaultpoolcontract, bound to a specific deployed contract.
func NewCvaultpoolcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*CvaultpoolcontractTransactor, error) {
	contract, err := bindCvaultpoolcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CvaultpoolcontractTransactor{contract: contract}, nil
}

// NewCvaultpoolcontractFilterer creates a new log filterer instance of Cvaultpoolcontract, bound to a specific deployed contract.
func NewCvaultpoolcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*CvaultpoolcontractFilterer, error) {
	contract, err := bindCvaultpoolcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CvaultpoolcontractFilterer{contract: contract}, nil
}

// bindCvaultpoolcontract binds a generic wrapper to an already deployed contract.
func bindCvaultpoolcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CvaultpoolcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cvaultpoolcontract *CvaultpoolcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cvaultpoolcontract.Contract.CvaultpoolcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cvaultpoolcontract *CvaultpoolcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.CvaultpoolcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cvaultpoolcontract *CvaultpoolcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.CvaultpoolcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cvaultpoolcontract *CvaultpoolcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cvaultpoolcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.contract.Transact(opts, method, params...)
}

// AverageFeesPerBlockEpoch is a free data retrieval call binding the contract method 0x03dec009.
//
// Solidity: function averageFeesPerBlockEpoch() view returns(uint256 averagePerBlock)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) AverageFeesPerBlockEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "averageFeesPerBlockEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AverageFeesPerBlockEpoch is a free data retrieval call binding the contract method 0x03dec009.
//
// Solidity: function averageFeesPerBlockEpoch() view returns(uint256 averagePerBlock)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) AverageFeesPerBlockEpoch() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.AverageFeesPerBlockEpoch(&_Cvaultpoolcontract.CallOpts)
}

// AverageFeesPerBlockEpoch is a free data retrieval call binding the contract method 0x03dec009.
//
// Solidity: function averageFeesPerBlockEpoch() view returns(uint256 averagePerBlock)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) AverageFeesPerBlockEpoch() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.AverageFeesPerBlockEpoch(&_Cvaultpoolcontract.CallOpts)
}

// AverageFeesPerBlockSinceStart is a free data retrieval call binding the contract method 0x9dbc2d90.
//
// Solidity: function averageFeesPerBlockSinceStart() view returns(uint256 averagePerBlock)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) AverageFeesPerBlockSinceStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "averageFeesPerBlockSinceStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AverageFeesPerBlockSinceStart is a free data retrieval call binding the contract method 0x9dbc2d90.
//
// Solidity: function averageFeesPerBlockSinceStart() view returns(uint256 averagePerBlock)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) AverageFeesPerBlockSinceStart() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.AverageFeesPerBlockSinceStart(&_Cvaultpoolcontract.CallOpts)
}

// AverageFeesPerBlockSinceStart is a free data retrieval call binding the contract method 0x9dbc2d90.
//
// Solidity: function averageFeesPerBlockSinceStart() view returns(uint256 averagePerBlock)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) AverageFeesPerBlockSinceStart() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.AverageFeesPerBlockSinceStart(&_Cvaultpoolcontract.CallOpts)
}

// ContractStartBlock is a free data retrieval call binding the contract method 0x49c5468d.
//
// Solidity: function contractStartBlock() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) ContractStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "contractStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractStartBlock is a free data retrieval call binding the contract method 0x49c5468d.
//
// Solidity: function contractStartBlock() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) ContractStartBlock() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.ContractStartBlock(&_Cvaultpoolcontract.CallOpts)
}

// ContractStartBlock is a free data retrieval call binding the contract method 0x49c5468d.
//
// Solidity: function contractStartBlock() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) ContractStartBlock() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.ContractStartBlock(&_Cvaultpoolcontract.CallOpts)
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() view returns(address)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) Core(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "core")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() view returns(address)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) Core() (common.Address, error) {
	return _Cvaultpoolcontract.Contract.Core(&_Cvaultpoolcontract.CallOpts)
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() view returns(address)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) Core() (common.Address, error) {
	return _Cvaultpoolcontract.Contract.Core(&_Cvaultpoolcontract.CallOpts)
}

// CumulativeRewardsSinceStart is a free data retrieval call binding the contract method 0xc8ffb873.
//
// Solidity: function cumulativeRewardsSinceStart() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) CumulativeRewardsSinceStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "cumulativeRewardsSinceStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeRewardsSinceStart is a free data retrieval call binding the contract method 0xc8ffb873.
//
// Solidity: function cumulativeRewardsSinceStart() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) CumulativeRewardsSinceStart() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.CumulativeRewardsSinceStart(&_Cvaultpoolcontract.CallOpts)
}

// CumulativeRewardsSinceStart is a free data retrieval call binding the contract method 0xc8ffb873.
//
// Solidity: function cumulativeRewardsSinceStart() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) CumulativeRewardsSinceStart() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.CumulativeRewardsSinceStart(&_Cvaultpoolcontract.CallOpts)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) Devaddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "devaddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) Devaddr() (common.Address, error) {
	return _Cvaultpoolcontract.Contract.Devaddr(&_Cvaultpoolcontract.CallOpts)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) Devaddr() (common.Address, error) {
	return _Cvaultpoolcontract.Contract.Devaddr(&_Cvaultpoolcontract.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) Epoch() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.Epoch(&_Cvaultpoolcontract.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) Epoch() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.Epoch(&_Cvaultpoolcontract.CallOpts)
}

// EpochCalculationStartBlock is a free data retrieval call binding the contract method 0x5d577c18.
//
// Solidity: function epochCalculationStartBlock() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) EpochCalculationStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "epochCalculationStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochCalculationStartBlock is a free data retrieval call binding the contract method 0x5d577c18.
//
// Solidity: function epochCalculationStartBlock() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) EpochCalculationStartBlock() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.EpochCalculationStartBlock(&_Cvaultpoolcontract.CallOpts)
}

// EpochCalculationStartBlock is a free data retrieval call binding the contract method 0x5d577c18.
//
// Solidity: function epochCalculationStartBlock() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) EpochCalculationStartBlock() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.EpochCalculationStartBlock(&_Cvaultpoolcontract.CallOpts)
}

// EpochRewards is a free data retrieval call binding the contract method 0x4dc47d34.
//
// Solidity: function epochRewards(uint256 ) view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) EpochRewards(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "epochRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochRewards is a free data retrieval call binding the contract method 0x4dc47d34.
//
// Solidity: function epochRewards(uint256 ) view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) EpochRewards(arg0 *big.Int) (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.EpochRewards(&_Cvaultpoolcontract.CallOpts, arg0)
}

// EpochRewards is a free data retrieval call binding the contract method 0x4dc47d34.
//
// Solidity: function epochRewards(uint256 ) view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) EpochRewards(arg0 *big.Int) (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.EpochRewards(&_Cvaultpoolcontract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) Owner() (common.Address, error) {
	return _Cvaultpoolcontract.Contract.Owner(&_Cvaultpoolcontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) Owner() (common.Address, error) {
	return _Cvaultpoolcontract.Contract.Owner(&_Cvaultpoolcontract.CallOpts)
}

// PendingCore is a free data retrieval call binding the contract method 0xa4f00c82.
//
// Solidity: function pendingCore(uint256 _pid, address _user) view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) PendingCore(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "pendingCore", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingCore is a free data retrieval call binding the contract method 0xa4f00c82.
//
// Solidity: function pendingCore(uint256 _pid, address _user) view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) PendingCore(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.PendingCore(&_Cvaultpoolcontract.CallOpts, _pid, _user)
}

// PendingCore is a free data retrieval call binding the contract method 0xa4f00c82.
//
// Solidity: function pendingCore(uint256 _pid, address _user) view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) PendingCore(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.PendingCore(&_Cvaultpoolcontract.CallOpts, _pid, _user)
}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) PendingRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "pendingRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) PendingRewards() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.PendingRewards(&_Cvaultpoolcontract.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) PendingRewards() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.PendingRewards(&_Cvaultpoolcontract.CallOpts)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address token, uint256 allocPoint, uint256 accCorePerShare, bool withdrawable)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Token           common.Address
	AllocPoint      *big.Int
	AccCorePerShare *big.Int
	Withdrawable    bool
}, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		Token           common.Address
		AllocPoint      *big.Int
		AccCorePerShare *big.Int
		Withdrawable    bool
	})

	outstruct.Token = out[0].(common.Address)
	outstruct.AllocPoint = out[1].(*big.Int)
	outstruct.AccCorePerShare = out[2].(*big.Int)
	outstruct.Withdrawable = out[3].(bool)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address token, uint256 allocPoint, uint256 accCorePerShare, bool withdrawable)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) PoolInfo(arg0 *big.Int) (struct {
	Token           common.Address
	AllocPoint      *big.Int
	AccCorePerShare *big.Int
	Withdrawable    bool
}, error) {
	return _Cvaultpoolcontract.Contract.PoolInfo(&_Cvaultpoolcontract.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address token, uint256 allocPoint, uint256 accCorePerShare, bool withdrawable)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) PoolInfo(arg0 *big.Int) (struct {
	Token           common.Address
	AllocPoint      *big.Int
	AccCorePerShare *big.Int
	Withdrawable    bool
}, error) {
	return _Cvaultpoolcontract.Contract.PoolInfo(&_Cvaultpoolcontract.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) PoolLength() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.PoolLength(&_Cvaultpoolcontract.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) PoolLength() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.PoolLength(&_Cvaultpoolcontract.CallOpts)
}

// RewardsInThisEpoch is a free data retrieval call binding the contract method 0x608c8d3a.
//
// Solidity: function rewardsInThisEpoch() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) RewardsInThisEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "rewardsInThisEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsInThisEpoch is a free data retrieval call binding the contract method 0x608c8d3a.
//
// Solidity: function rewardsInThisEpoch() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) RewardsInThisEpoch() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.RewardsInThisEpoch(&_Cvaultpoolcontract.CallOpts)
}

// RewardsInThisEpoch is a free data retrieval call binding the contract method 0x608c8d3a.
//
// Solidity: function rewardsInThisEpoch() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) RewardsInThisEpoch() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.RewardsInThisEpoch(&_Cvaultpoolcontract.CallOpts)
}

// SuperAdmin is a free data retrieval call binding the contract method 0x29575f6a.
//
// Solidity: function superAdmin() view returns(address)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) SuperAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "superAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperAdmin is a free data retrieval call binding the contract method 0x29575f6a.
//
// Solidity: function superAdmin() view returns(address)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) SuperAdmin() (common.Address, error) {
	return _Cvaultpoolcontract.Contract.SuperAdmin(&_Cvaultpoolcontract.CallOpts)
}

// SuperAdmin is a free data retrieval call binding the contract method 0x29575f6a.
//
// Solidity: function superAdmin() view returns(address)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) SuperAdmin() (common.Address, error) {
	return _Cvaultpoolcontract.Contract.SuperAdmin(&_Cvaultpoolcontract.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) TotalAllocPoint() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.TotalAllocPoint(&_Cvaultpoolcontract.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _Cvaultpoolcontract.Contract.TotalAllocPoint(&_Cvaultpoolcontract.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Cvaultpoolcontract *CvaultpoolcontractCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _Cvaultpoolcontract.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})

	outstruct.Amount = out[0].(*big.Int)
	outstruct.RewardDebt = out[1].(*big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Cvaultpoolcontract.Contract.UserInfo(&_Cvaultpoolcontract.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Cvaultpoolcontract *CvaultpoolcontractCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Cvaultpoolcontract.Contract.UserInfo(&_Cvaultpoolcontract.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0xc507aeaa.
//
// Solidity: function add(uint256 _allocPoint, address _token, bool _withUpdate, bool _withdrawable) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _token common.Address, _withUpdate bool, _withdrawable bool) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "add", _allocPoint, _token, _withUpdate, _withdrawable)
}

// Add is a paid mutator transaction binding the contract method 0xc507aeaa.
//
// Solidity: function add(uint256 _allocPoint, address _token, bool _withUpdate, bool _withdrawable) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) Add(_allocPoint *big.Int, _token common.Address, _withUpdate bool, _withdrawable bool) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.Add(&_Cvaultpoolcontract.TransactOpts, _allocPoint, _token, _withUpdate, _withdrawable)
}

// Add is a paid mutator transaction binding the contract method 0xc507aeaa.
//
// Solidity: function add(uint256 _allocPoint, address _token, bool _withUpdate, bool _withdrawable) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) Add(_allocPoint *big.Int, _token common.Address, _withUpdate bool, _withdrawable bool) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.Add(&_Cvaultpoolcontract.TransactOpts, _allocPoint, _token, _withUpdate, _withdrawable)
}

// AddPendingRewards is a paid mutator transaction binding the contract method 0x423d6fa0.
//
// Solidity: function addPendingRewards(uint256 _) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) AddPendingRewards(opts *bind.TransactOpts, _ *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "addPendingRewards", nil)
}

// AddPendingRewards is a paid mutator transaction binding the contract method 0x423d6fa0.
//
// Solidity: function addPendingRewards(uint256 _) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) AddPendingRewards(_ *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.AddPendingRewards(&_Cvaultpoolcontract.TransactOpts, nil)
}

// AddPendingRewards is a paid mutator transaction binding the contract method 0x423d6fa0.
//
// Solidity: function addPendingRewards(uint256 _) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) AddPendingRewards(_ *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.AddPendingRewards(&_Cvaultpoolcontract.TransactOpts, nil)
}

// BurnSuperAdmin is a paid mutator transaction binding the contract method 0x934eaa50.
//
// Solidity: function burnSuperAdmin() returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) BurnSuperAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "burnSuperAdmin")
}

// BurnSuperAdmin is a paid mutator transaction binding the contract method 0x934eaa50.
//
// Solidity: function burnSuperAdmin() returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) BurnSuperAdmin() (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.BurnSuperAdmin(&_Cvaultpoolcontract.TransactOpts)
}

// BurnSuperAdmin is a paid mutator transaction binding the contract method 0x934eaa50.
//
// Solidity: function burnSuperAdmin() returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) BurnSuperAdmin() (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.BurnSuperAdmin(&_Cvaultpoolcontract.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.Deposit(&_Cvaultpoolcontract.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.Deposit(&_Cvaultpoolcontract.TransactOpts, _pid, _amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x4cf5fbf5.
//
// Solidity: function depositFor(address depositFor, uint256 _pid, uint256 _amount) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) DepositFor(opts *bind.TransactOpts, depositFor common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "depositFor", depositFor, _pid, _amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x4cf5fbf5.
//
// Solidity: function depositFor(address depositFor, uint256 _pid, uint256 _amount) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) DepositFor(depositFor common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.DepositFor(&_Cvaultpoolcontract.TransactOpts, depositFor, _pid, _amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x4cf5fbf5.
//
// Solidity: function depositFor(address depositFor, uint256 _pid, uint256 _amount) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) DepositFor(depositFor common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.DepositFor(&_Cvaultpoolcontract.TransactOpts, depositFor, _pid, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.EmergencyWithdraw(&_Cvaultpoolcontract.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.EmergencyWithdraw(&_Cvaultpoolcontract.TransactOpts, _pid)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _core, address _devaddr, address superAdmin) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) Initialize(opts *bind.TransactOpts, _core common.Address, _devaddr common.Address, superAdmin common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "initialize", _core, _devaddr, superAdmin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _core, address _devaddr, address superAdmin) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) Initialize(_core common.Address, _devaddr common.Address, superAdmin common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.Initialize(&_Cvaultpoolcontract.TransactOpts, _core, _devaddr, superAdmin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _core, address _devaddr, address superAdmin) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) Initialize(_core common.Address, _devaddr common.Address, superAdmin common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.Initialize(&_Cvaultpoolcontract.TransactOpts, _core, _devaddr, superAdmin)
}

// IsContract is a paid mutator transaction binding the contract method 0x16279055.
//
// Solidity: function isContract(address addr) returns(bool)
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) IsContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "isContract", addr)
}

// IsContract is a paid mutator transaction binding the contract method 0x16279055.
//
// Solidity: function isContract(address addr) returns(bool)
func (_Cvaultpoolcontract *CvaultpoolcontractSession) IsContract(addr common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.IsContract(&_Cvaultpoolcontract.TransactOpts, addr)
}

// IsContract is a paid mutator transaction binding the contract method 0x16279055.
//
// Solidity: function isContract(address addr) returns(bool)
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) IsContract(addr common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.IsContract(&_Cvaultpoolcontract.TransactOpts, addr)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) MassUpdatePools() (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.MassUpdatePools(&_Cvaultpoolcontract.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.MassUpdatePools(&_Cvaultpoolcontract.TransactOpts)
}

// NewSuperAdmin is a paid mutator transaction binding the contract method 0xa676860a.
//
// Solidity: function newSuperAdmin(address newOwner) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) NewSuperAdmin(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "newSuperAdmin", newOwner)
}

// NewSuperAdmin is a paid mutator transaction binding the contract method 0xa676860a.
//
// Solidity: function newSuperAdmin(address newOwner) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) NewSuperAdmin(newOwner common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.NewSuperAdmin(&_Cvaultpoolcontract.TransactOpts, newOwner)
}

// NewSuperAdmin is a paid mutator transaction binding the contract method 0xa676860a.
//
// Solidity: function newSuperAdmin(address newOwner) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) NewSuperAdmin(newOwner common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.NewSuperAdmin(&_Cvaultpoolcontract.TransactOpts, newOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.RenounceOwnership(&_Cvaultpoolcontract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.RenounceOwnership(&_Cvaultpoolcontract.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "set", _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.Set(&_Cvaultpoolcontract.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.Set(&_Cvaultpoolcontract.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// SetAllowanceForPoolToken is a paid mutator transaction binding the contract method 0xdbe0901f.
//
// Solidity: function setAllowanceForPoolToken(address spender, uint256 _pid, uint256 value) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) SetAllowanceForPoolToken(opts *bind.TransactOpts, spender common.Address, _pid *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "setAllowanceForPoolToken", spender, _pid, value)
}

// SetAllowanceForPoolToken is a paid mutator transaction binding the contract method 0xdbe0901f.
//
// Solidity: function setAllowanceForPoolToken(address spender, uint256 _pid, uint256 value) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) SetAllowanceForPoolToken(spender common.Address, _pid *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.SetAllowanceForPoolToken(&_Cvaultpoolcontract.TransactOpts, spender, _pid, value)
}

// SetAllowanceForPoolToken is a paid mutator transaction binding the contract method 0xdbe0901f.
//
// Solidity: function setAllowanceForPoolToken(address spender, uint256 _pid, uint256 value) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) SetAllowanceForPoolToken(spender common.Address, _pid *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.SetAllowanceForPoolToken(&_Cvaultpoolcontract.TransactOpts, spender, _pid, value)
}

// SetDevFee is a paid mutator transaction binding the contract method 0xe18cb4fe.
//
// Solidity: function setDevFee(uint16 _DEV_FEE) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) SetDevFee(opts *bind.TransactOpts, _DEV_FEE uint16) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "setDevFee", _DEV_FEE)
}

// SetDevFee is a paid mutator transaction binding the contract method 0xe18cb4fe.
//
// Solidity: function setDevFee(uint16 _DEV_FEE) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) SetDevFee(_DEV_FEE uint16) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.SetDevFee(&_Cvaultpoolcontract.TransactOpts, _DEV_FEE)
}

// SetDevFee is a paid mutator transaction binding the contract method 0xe18cb4fe.
//
// Solidity: function setDevFee(uint16 _DEV_FEE) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) SetDevFee(_DEV_FEE uint16) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.SetDevFee(&_Cvaultpoolcontract.TransactOpts, _DEV_FEE)
}

// SetDevFeeReciever is a paid mutator transaction binding the contract method 0x2d6754e5.
//
// Solidity: function setDevFeeReciever(address _devaddr) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) SetDevFeeReciever(opts *bind.TransactOpts, _devaddr common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "setDevFeeReciever", _devaddr)
}

// SetDevFeeReciever is a paid mutator transaction binding the contract method 0x2d6754e5.
//
// Solidity: function setDevFeeReciever(address _devaddr) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) SetDevFeeReciever(_devaddr common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.SetDevFeeReciever(&_Cvaultpoolcontract.TransactOpts, _devaddr)
}

// SetDevFeeReciever is a paid mutator transaction binding the contract method 0x2d6754e5.
//
// Solidity: function setDevFeeReciever(address _devaddr) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) SetDevFeeReciever(_devaddr common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.SetDevFeeReciever(&_Cvaultpoolcontract.TransactOpts, _devaddr)
}

// SetPoolWithdrawable is a paid mutator transaction binding the contract method 0x5207cc0d.
//
// Solidity: function setPoolWithdrawable(uint256 _pid, bool _withdrawable) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) SetPoolWithdrawable(opts *bind.TransactOpts, _pid *big.Int, _withdrawable bool) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "setPoolWithdrawable", _pid, _withdrawable)
}

// SetPoolWithdrawable is a paid mutator transaction binding the contract method 0x5207cc0d.
//
// Solidity: function setPoolWithdrawable(uint256 _pid, bool _withdrawable) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) SetPoolWithdrawable(_pid *big.Int, _withdrawable bool) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.SetPoolWithdrawable(&_Cvaultpoolcontract.TransactOpts, _pid, _withdrawable)
}

// SetPoolWithdrawable is a paid mutator transaction binding the contract method 0x5207cc0d.
//
// Solidity: function setPoolWithdrawable(uint256 _pid, bool _withdrawable) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) SetPoolWithdrawable(_pid *big.Int, _withdrawable bool) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.SetPoolWithdrawable(&_Cvaultpoolcontract.TransactOpts, _pid, _withdrawable)
}

// SetStrategyContractOrDistributionContractAllowance is a paid mutator transaction binding the contract method 0xc4014588.
//
// Solidity: function setStrategyContractOrDistributionContractAllowance(address tokenAddress, uint256 _amount, address contractAddress) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) SetStrategyContractOrDistributionContractAllowance(opts *bind.TransactOpts, tokenAddress common.Address, _amount *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "setStrategyContractOrDistributionContractAllowance", tokenAddress, _amount, contractAddress)
}

// SetStrategyContractOrDistributionContractAllowance is a paid mutator transaction binding the contract method 0xc4014588.
//
// Solidity: function setStrategyContractOrDistributionContractAllowance(address tokenAddress, uint256 _amount, address contractAddress) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) SetStrategyContractOrDistributionContractAllowance(tokenAddress common.Address, _amount *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.SetStrategyContractOrDistributionContractAllowance(&_Cvaultpoolcontract.TransactOpts, tokenAddress, _amount, contractAddress)
}

// SetStrategyContractOrDistributionContractAllowance is a paid mutator transaction binding the contract method 0xc4014588.
//
// Solidity: function setStrategyContractOrDistributionContractAllowance(address tokenAddress, uint256 _amount, address contractAddress) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) SetStrategyContractOrDistributionContractAllowance(tokenAddress common.Address, _amount *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.SetStrategyContractOrDistributionContractAllowance(&_Cvaultpoolcontract.TransactOpts, tokenAddress, _amount, contractAddress)
}

// StartNewEpoch is a paid mutator transaction binding the contract method 0x3aab0a62.
//
// Solidity: function startNewEpoch() returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) StartNewEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "startNewEpoch")
}

// StartNewEpoch is a paid mutator transaction binding the contract method 0x3aab0a62.
//
// Solidity: function startNewEpoch() returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) StartNewEpoch() (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.StartNewEpoch(&_Cvaultpoolcontract.TransactOpts)
}

// StartNewEpoch is a paid mutator transaction binding the contract method 0x3aab0a62.
//
// Solidity: function startNewEpoch() returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) StartNewEpoch() (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.StartNewEpoch(&_Cvaultpoolcontract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.TransferOwnership(&_Cvaultpoolcontract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.TransferOwnership(&_Cvaultpoolcontract.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.Withdraw(&_Cvaultpoolcontract.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.Withdraw(&_Cvaultpoolcontract.TransactOpts, _pid, _amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x3a0967cd.
//
// Solidity: function withdrawFrom(address owner, uint256 _pid, uint256 _amount) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactor) WithdrawFrom(opts *bind.TransactOpts, owner common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.contract.Transact(opts, "withdrawFrom", owner, _pid, _amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x3a0967cd.
//
// Solidity: function withdrawFrom(address owner, uint256 _pid, uint256 _amount) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractSession) WithdrawFrom(owner common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.WithdrawFrom(&_Cvaultpoolcontract.TransactOpts, owner, _pid, _amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x3a0967cd.
//
// Solidity: function withdrawFrom(address owner, uint256 _pid, uint256 _amount) returns()
func (_Cvaultpoolcontract *CvaultpoolcontractTransactorSession) WithdrawFrom(owner common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Cvaultpoolcontract.Contract.WithdrawFrom(&_Cvaultpoolcontract.TransactOpts, owner, _pid, _amount)
}

// CvaultpoolcontractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Cvaultpoolcontract contract.
type CvaultpoolcontractApprovalIterator struct {
	Event *CvaultpoolcontractApproval // Event containing the contract specifics and raw log

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
func (it *CvaultpoolcontractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CvaultpoolcontractApproval)
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
		it.Event = new(CvaultpoolcontractApproval)
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
func (it *CvaultpoolcontractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CvaultpoolcontractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CvaultpoolcontractApproval represents a Approval event raised by the Cvaultpoolcontract contract.
type CvaultpoolcontractApproval struct {
	Owner   common.Address
	Spender common.Address
	Pid     *big.Int
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xb3fd5071835887567a0671151121894ddccc2842f1d10bedad13e0d17cace9a7.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 _pid, uint256 value)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CvaultpoolcontractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cvaultpoolcontract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CvaultpoolcontractApprovalIterator{contract: _Cvaultpoolcontract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xb3fd5071835887567a0671151121894ddccc2842f1d10bedad13e0d17cace9a7.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 _pid, uint256 value)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CvaultpoolcontractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cvaultpoolcontract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CvaultpoolcontractApproval)
				if err := _Cvaultpoolcontract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0xb3fd5071835887567a0671151121894ddccc2842f1d10bedad13e0d17cace9a7.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 _pid, uint256 value)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) ParseApproval(log types.Log) (*CvaultpoolcontractApproval, error) {
	event := new(CvaultpoolcontractApproval)
	if err := _Cvaultpoolcontract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CvaultpoolcontractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Cvaultpoolcontract contract.
type CvaultpoolcontractDepositIterator struct {
	Event *CvaultpoolcontractDeposit // Event containing the contract specifics and raw log

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
func (it *CvaultpoolcontractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CvaultpoolcontractDeposit)
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
		it.Event = new(CvaultpoolcontractDeposit)
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
func (it *CvaultpoolcontractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CvaultpoolcontractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CvaultpoolcontractDeposit represents a Deposit event raised by the Cvaultpoolcontract contract.
type CvaultpoolcontractDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*CvaultpoolcontractDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Cvaultpoolcontract.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &CvaultpoolcontractDepositIterator{contract: _Cvaultpoolcontract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CvaultpoolcontractDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Cvaultpoolcontract.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CvaultpoolcontractDeposit)
				if err := _Cvaultpoolcontract.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) ParseDeposit(log types.Log) (*CvaultpoolcontractDeposit, error) {
	event := new(CvaultpoolcontractDeposit)
	if err := _Cvaultpoolcontract.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CvaultpoolcontractEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Cvaultpoolcontract contract.
type CvaultpoolcontractEmergencyWithdrawIterator struct {
	Event *CvaultpoolcontractEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *CvaultpoolcontractEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CvaultpoolcontractEmergencyWithdraw)
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
		it.Event = new(CvaultpoolcontractEmergencyWithdraw)
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
func (it *CvaultpoolcontractEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CvaultpoolcontractEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CvaultpoolcontractEmergencyWithdraw represents a EmergencyWithdraw event raised by the Cvaultpoolcontract contract.
type CvaultpoolcontractEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*CvaultpoolcontractEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Cvaultpoolcontract.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &CvaultpoolcontractEmergencyWithdrawIterator{contract: _Cvaultpoolcontract.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *CvaultpoolcontractEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Cvaultpoolcontract.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CvaultpoolcontractEmergencyWithdraw)
				if err := _Cvaultpoolcontract.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) ParseEmergencyWithdraw(log types.Log) (*CvaultpoolcontractEmergencyWithdraw, error) {
	event := new(CvaultpoolcontractEmergencyWithdraw)
	if err := _Cvaultpoolcontract.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CvaultpoolcontractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Cvaultpoolcontract contract.
type CvaultpoolcontractOwnershipTransferredIterator struct {
	Event *CvaultpoolcontractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CvaultpoolcontractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CvaultpoolcontractOwnershipTransferred)
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
		it.Event = new(CvaultpoolcontractOwnershipTransferred)
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
func (it *CvaultpoolcontractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CvaultpoolcontractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CvaultpoolcontractOwnershipTransferred represents a OwnershipTransferred event raised by the Cvaultpoolcontract contract.
type CvaultpoolcontractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CvaultpoolcontractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cvaultpoolcontract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CvaultpoolcontractOwnershipTransferredIterator{contract: _Cvaultpoolcontract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CvaultpoolcontractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cvaultpoolcontract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CvaultpoolcontractOwnershipTransferred)
				if err := _Cvaultpoolcontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) ParseOwnershipTransferred(log types.Log) (*CvaultpoolcontractOwnershipTransferred, error) {
	event := new(CvaultpoolcontractOwnershipTransferred)
	if err := _Cvaultpoolcontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CvaultpoolcontractSuperAdminTransferedIterator is returned from FilterSuperAdminTransfered and is used to iterate over the raw logs and unpacked data for SuperAdminTransfered events raised by the Cvaultpoolcontract contract.
type CvaultpoolcontractSuperAdminTransferedIterator struct {
	Event *CvaultpoolcontractSuperAdminTransfered // Event containing the contract specifics and raw log

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
func (it *CvaultpoolcontractSuperAdminTransferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CvaultpoolcontractSuperAdminTransfered)
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
		it.Event = new(CvaultpoolcontractSuperAdminTransfered)
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
func (it *CvaultpoolcontractSuperAdminTransferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CvaultpoolcontractSuperAdminTransferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CvaultpoolcontractSuperAdminTransfered represents a SuperAdminTransfered event raised by the Cvaultpoolcontract contract.
type CvaultpoolcontractSuperAdminTransfered struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSuperAdminTransfered is a free log retrieval operation binding the contract event 0xf564c40f4f45e62a2c1e6c22e8bfb46501f0f71fa1c72e5358903fa1115a4b64.
//
// Solidity: event SuperAdminTransfered(address indexed previousOwner, address indexed newOwner)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) FilterSuperAdminTransfered(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CvaultpoolcontractSuperAdminTransferedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cvaultpoolcontract.contract.FilterLogs(opts, "SuperAdminTransfered", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CvaultpoolcontractSuperAdminTransferedIterator{contract: _Cvaultpoolcontract.contract, event: "SuperAdminTransfered", logs: logs, sub: sub}, nil
}

// WatchSuperAdminTransfered is a free log subscription operation binding the contract event 0xf564c40f4f45e62a2c1e6c22e8bfb46501f0f71fa1c72e5358903fa1115a4b64.
//
// Solidity: event SuperAdminTransfered(address indexed previousOwner, address indexed newOwner)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) WatchSuperAdminTransfered(opts *bind.WatchOpts, sink chan<- *CvaultpoolcontractSuperAdminTransfered, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cvaultpoolcontract.contract.WatchLogs(opts, "SuperAdminTransfered", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CvaultpoolcontractSuperAdminTransfered)
				if err := _Cvaultpoolcontract.contract.UnpackLog(event, "SuperAdminTransfered", log); err != nil {
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

// ParseSuperAdminTransfered is a log parse operation binding the contract event 0xf564c40f4f45e62a2c1e6c22e8bfb46501f0f71fa1c72e5358903fa1115a4b64.
//
// Solidity: event SuperAdminTransfered(address indexed previousOwner, address indexed newOwner)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) ParseSuperAdminTransfered(log types.Log) (*CvaultpoolcontractSuperAdminTransfered, error) {
	event := new(CvaultpoolcontractSuperAdminTransfered)
	if err := _Cvaultpoolcontract.contract.UnpackLog(event, "SuperAdminTransfered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CvaultpoolcontractWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Cvaultpoolcontract contract.
type CvaultpoolcontractWithdrawIterator struct {
	Event *CvaultpoolcontractWithdraw // Event containing the contract specifics and raw log

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
func (it *CvaultpoolcontractWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CvaultpoolcontractWithdraw)
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
		it.Event = new(CvaultpoolcontractWithdraw)
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
func (it *CvaultpoolcontractWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CvaultpoolcontractWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CvaultpoolcontractWithdraw represents a Withdraw event raised by the Cvaultpoolcontract contract.
type CvaultpoolcontractWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*CvaultpoolcontractWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Cvaultpoolcontract.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &CvaultpoolcontractWithdrawIterator{contract: _Cvaultpoolcontract.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CvaultpoolcontractWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Cvaultpoolcontract.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CvaultpoolcontractWithdraw)
				if err := _Cvaultpoolcontract.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Cvaultpoolcontract *CvaultpoolcontractFilterer) ParseWithdraw(log types.Log) (*CvaultpoolcontractWithdraw, error) {
	event := new(CvaultpoolcontractWithdraw)
	if err := _Cvaultpoolcontract.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
