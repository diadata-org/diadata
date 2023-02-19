// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package platypusfinance

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

// BaseMasterPlatypusMetaData contains all meta data concerning the BaseMasterPlatypus contract.
var BaseMasterPlatypusMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIRewarder\",\"name\":\"rewarder\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIRewarder\",\"name\":\"rewarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ptpPerSec\",\"type\":\"uint256\"}],\"name\":\"UpdateEmissionRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accPtpPerShare\",\"type\":\"uint256\"}],\"name\":\"UpdatePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseAllocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIAsset\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"contractIRewarder\",\"name\":\"_rewarder\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyPtpWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_ptp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ptpPerSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_pids\",\"type\":\"uint256[]\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_pids\",\"type\":\"uint256[]\"}],\"name\":\"multiClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerCandidate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingPtp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bonusTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bonusTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pendingBonusToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIAsset\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAllocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accPtpPerShare\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewarder\",\"name\":\"rewarder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"proposeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ptp\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ptpPerSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"rewarderBonusTokenInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bonusTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bonusTokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseAllocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewarder\",\"name\":\"_rewarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBaseMasterPlatypus\",\"name\":\"_newMasterPlatypus\",\"type\":\"address\"}],\"name\":\"setNewMasterPlatypus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBaseAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ptpPerSec\",\"type\":\"uint256\"}],\"name\":\"updateEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BaseMasterPlatypusABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseMasterPlatypusMetaData.ABI instead.
var BaseMasterPlatypusABI = BaseMasterPlatypusMetaData.ABI

// BaseMasterPlatypus is an auto generated Go binding around an Ethereum contract.
type BaseMasterPlatypus struct {
	BaseMasterPlatypusCaller     // Read-only binding to the contract
	BaseMasterPlatypusTransactor // Write-only binding to the contract
	BaseMasterPlatypusFilterer   // Log filterer for contract events
}

// BaseMasterPlatypusCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseMasterPlatypusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseMasterPlatypusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseMasterPlatypusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseMasterPlatypusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseMasterPlatypusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseMasterPlatypusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseMasterPlatypusSession struct {
	Contract     *BaseMasterPlatypus // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BaseMasterPlatypusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseMasterPlatypusCallerSession struct {
	Contract *BaseMasterPlatypusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BaseMasterPlatypusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseMasterPlatypusTransactorSession struct {
	Contract     *BaseMasterPlatypusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BaseMasterPlatypusRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseMasterPlatypusRaw struct {
	Contract *BaseMasterPlatypus // Generic contract binding to access the raw methods on
}

// BaseMasterPlatypusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseMasterPlatypusCallerRaw struct {
	Contract *BaseMasterPlatypusCaller // Generic read-only contract binding to access the raw methods on
}

// BaseMasterPlatypusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseMasterPlatypusTransactorRaw struct {
	Contract *BaseMasterPlatypusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseMasterPlatypus creates a new instance of BaseMasterPlatypus, bound to a specific deployed contract.
func NewBaseMasterPlatypus(address common.Address, backend bind.ContractBackend) (*BaseMasterPlatypus, error) {
	contract, err := bindBaseMasterPlatypus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypus{BaseMasterPlatypusCaller: BaseMasterPlatypusCaller{contract: contract}, BaseMasterPlatypusTransactor: BaseMasterPlatypusTransactor{contract: contract}, BaseMasterPlatypusFilterer: BaseMasterPlatypusFilterer{contract: contract}}, nil
}

// NewBaseMasterPlatypusCaller creates a new read-only instance of BaseMasterPlatypus, bound to a specific deployed contract.
func NewBaseMasterPlatypusCaller(address common.Address, caller bind.ContractCaller) (*BaseMasterPlatypusCaller, error) {
	contract, err := bindBaseMasterPlatypus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypusCaller{contract: contract}, nil
}

// NewBaseMasterPlatypusTransactor creates a new write-only instance of BaseMasterPlatypus, bound to a specific deployed contract.
func NewBaseMasterPlatypusTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseMasterPlatypusTransactor, error) {
	contract, err := bindBaseMasterPlatypus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypusTransactor{contract: contract}, nil
}

// NewBaseMasterPlatypusFilterer creates a new log filterer instance of BaseMasterPlatypus, bound to a specific deployed contract.
func NewBaseMasterPlatypusFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseMasterPlatypusFilterer, error) {
	contract, err := bindBaseMasterPlatypus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypusFilterer{contract: contract}, nil
}

// bindBaseMasterPlatypus binds a generic wrapper to an already deployed contract.
func bindBaseMasterPlatypus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseMasterPlatypusABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseMasterPlatypus *BaseMasterPlatypusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseMasterPlatypus.Contract.BaseMasterPlatypusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseMasterPlatypus *BaseMasterPlatypusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.BaseMasterPlatypusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseMasterPlatypus *BaseMasterPlatypusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.BaseMasterPlatypusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseMasterPlatypus *BaseMasterPlatypusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseMasterPlatypus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BaseMasterPlatypus *BaseMasterPlatypusCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMasterPlatypus.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) Owner() (common.Address, error) {
	return _BaseMasterPlatypus.Contract.Owner(&_BaseMasterPlatypus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BaseMasterPlatypus *BaseMasterPlatypusCallerSession) Owner() (common.Address, error) {
	return _BaseMasterPlatypus.Contract.Owner(&_BaseMasterPlatypus.CallOpts)
}

// OwnerCandidate is a free data retrieval call binding the contract method 0x5f504a82.
//
// Solidity: function ownerCandidate() view returns(address)
func (_BaseMasterPlatypus *BaseMasterPlatypusCaller) OwnerCandidate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMasterPlatypus.contract.Call(opts, &out, "ownerCandidate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerCandidate is a free data retrieval call binding the contract method 0x5f504a82.
//
// Solidity: function ownerCandidate() view returns(address)
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) OwnerCandidate() (common.Address, error) {
	return _BaseMasterPlatypus.Contract.OwnerCandidate(&_BaseMasterPlatypus.CallOpts)
}

// OwnerCandidate is a free data retrieval call binding the contract method 0x5f504a82.
//
// Solidity: function ownerCandidate() view returns(address)
func (_BaseMasterPlatypus *BaseMasterPlatypusCallerSession) OwnerCandidate() (common.Address, error) {
	return _BaseMasterPlatypus.Contract.OwnerCandidate(&_BaseMasterPlatypus.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BaseMasterPlatypus *BaseMasterPlatypusCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BaseMasterPlatypus.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) Paused() (bool, error) {
	return _BaseMasterPlatypus.Contract.Paused(&_BaseMasterPlatypus.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BaseMasterPlatypus *BaseMasterPlatypusCallerSession) Paused() (bool, error) {
	return _BaseMasterPlatypus.Contract.Paused(&_BaseMasterPlatypus.CallOpts)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingPtp, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_BaseMasterPlatypus *BaseMasterPlatypusCaller) PendingTokens(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (struct {
	PendingPtp        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	var out []interface{}
	err := _BaseMasterPlatypus.contract.Call(opts, &out, "pendingTokens", _pid, _user)

	outstruct := new(struct {
		PendingPtp        *big.Int
		BonusTokenAddress common.Address
		BonusTokenSymbol  string
		PendingBonusToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PendingPtp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BonusTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BonusTokenSymbol = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.PendingBonusToken = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingPtp, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingPtp        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _BaseMasterPlatypus.Contract.PendingTokens(&_BaseMasterPlatypus.CallOpts, _pid, _user)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingPtp, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_BaseMasterPlatypus *BaseMasterPlatypusCallerSession) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingPtp        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _BaseMasterPlatypus.Contract.PendingTokens(&_BaseMasterPlatypus.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 baseAllocPoint, uint256 lastRewardTimestamp, uint256 accPtpPerShare, address rewarder)
func (_BaseMasterPlatypus *BaseMasterPlatypusCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken             common.Address
	BaseAllocPoint      *big.Int
	LastRewardTimestamp *big.Int
	AccPtpPerShare      *big.Int
	Rewarder            common.Address
}, error) {
	var out []interface{}
	err := _BaseMasterPlatypus.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken             common.Address
		BaseAllocPoint      *big.Int
		LastRewardTimestamp *big.Int
		AccPtpPerShare      *big.Int
		Rewarder            common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BaseAllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccPtpPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Rewarder = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 baseAllocPoint, uint256 lastRewardTimestamp, uint256 accPtpPerShare, address rewarder)
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	BaseAllocPoint      *big.Int
	LastRewardTimestamp *big.Int
	AccPtpPerShare      *big.Int
	Rewarder            common.Address
}, error) {
	return _BaseMasterPlatypus.Contract.PoolInfo(&_BaseMasterPlatypus.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 baseAllocPoint, uint256 lastRewardTimestamp, uint256 accPtpPerShare, address rewarder)
func (_BaseMasterPlatypus *BaseMasterPlatypusCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	BaseAllocPoint      *big.Int
	LastRewardTimestamp *big.Int
	AccPtpPerShare      *big.Int
	Rewarder            common.Address
}, error) {
	return _BaseMasterPlatypus.Contract.PoolInfo(&_BaseMasterPlatypus.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseMasterPlatypus.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) PoolLength() (*big.Int, error) {
	return _BaseMasterPlatypus.Contract.PoolLength(&_BaseMasterPlatypus.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusCallerSession) PoolLength() (*big.Int, error) {
	return _BaseMasterPlatypus.Contract.PoolLength(&_BaseMasterPlatypus.CallOpts)
}

// Ptp is a free data retrieval call binding the contract method 0x6af66772.
//
// Solidity: function ptp() view returns(address)
func (_BaseMasterPlatypus *BaseMasterPlatypusCaller) Ptp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMasterPlatypus.contract.Call(opts, &out, "ptp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ptp is a free data retrieval call binding the contract method 0x6af66772.
//
// Solidity: function ptp() view returns(address)
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) Ptp() (common.Address, error) {
	return _BaseMasterPlatypus.Contract.Ptp(&_BaseMasterPlatypus.CallOpts)
}

// Ptp is a free data retrieval call binding the contract method 0x6af66772.
//
// Solidity: function ptp() view returns(address)
func (_BaseMasterPlatypus *BaseMasterPlatypusCallerSession) Ptp() (common.Address, error) {
	return _BaseMasterPlatypus.Contract.Ptp(&_BaseMasterPlatypus.CallOpts)
}

// PtpPerSec is a free data retrieval call binding the contract method 0x9702d3e2.
//
// Solidity: function ptpPerSec() view returns(uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusCaller) PtpPerSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseMasterPlatypus.contract.Call(opts, &out, "ptpPerSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PtpPerSec is a free data retrieval call binding the contract method 0x9702d3e2.
//
// Solidity: function ptpPerSec() view returns(uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) PtpPerSec() (*big.Int, error) {
	return _BaseMasterPlatypus.Contract.PtpPerSec(&_BaseMasterPlatypus.CallOpts)
}

// PtpPerSec is a free data retrieval call binding the contract method 0x9702d3e2.
//
// Solidity: function ptpPerSec() view returns(uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusCallerSession) PtpPerSec() (*big.Int, error) {
	return _BaseMasterPlatypus.Contract.PtpPerSec(&_BaseMasterPlatypus.CallOpts)
}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address bonusTokenAddress, string bonusTokenSymbol)
func (_BaseMasterPlatypus *BaseMasterPlatypusCaller) RewarderBonusTokenInfo(opts *bind.CallOpts, _pid *big.Int) (struct {
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
}, error) {
	var out []interface{}
	err := _BaseMasterPlatypus.contract.Call(opts, &out, "rewarderBonusTokenInfo", _pid)

	outstruct := new(struct {
		BonusTokenAddress common.Address
		BonusTokenSymbol  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BonusTokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BonusTokenSymbol = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address bonusTokenAddress, string bonusTokenSymbol)
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) RewarderBonusTokenInfo(_pid *big.Int) (struct {
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
}, error) {
	return _BaseMasterPlatypus.Contract.RewarderBonusTokenInfo(&_BaseMasterPlatypus.CallOpts, _pid)
}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address bonusTokenAddress, string bonusTokenSymbol)
func (_BaseMasterPlatypus *BaseMasterPlatypusCallerSession) RewarderBonusTokenInfo(_pid *big.Int) (struct {
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
}, error) {
	return _BaseMasterPlatypus.Contract.RewarderBonusTokenInfo(&_BaseMasterPlatypus.CallOpts, _pid)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusCaller) StartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseMasterPlatypus.contract.Call(opts, &out, "startTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) StartTimestamp() (*big.Int, error) {
	return _BaseMasterPlatypus.Contract.StartTimestamp(&_BaseMasterPlatypus.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusCallerSession) StartTimestamp() (*big.Int, error) {
	return _BaseMasterPlatypus.Contract.StartTimestamp(&_BaseMasterPlatypus.CallOpts)
}

// TotalBaseAllocPoint is a free data retrieval call binding the contract method 0x33e045fc.
//
// Solidity: function totalBaseAllocPoint() view returns(uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusCaller) TotalBaseAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseMasterPlatypus.contract.Call(opts, &out, "totalBaseAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBaseAllocPoint is a free data retrieval call binding the contract method 0x33e045fc.
//
// Solidity: function totalBaseAllocPoint() view returns(uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) TotalBaseAllocPoint() (*big.Int, error) {
	return _BaseMasterPlatypus.Contract.TotalBaseAllocPoint(&_BaseMasterPlatypus.CallOpts)
}

// TotalBaseAllocPoint is a free data retrieval call binding the contract method 0x33e045fc.
//
// Solidity: function totalBaseAllocPoint() view returns(uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusCallerSession) TotalBaseAllocPoint() (*big.Int, error) {
	return _BaseMasterPlatypus.Contract.TotalBaseAllocPoint(&_BaseMasterPlatypus.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_BaseMasterPlatypus *BaseMasterPlatypusCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _BaseMasterPlatypus.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _BaseMasterPlatypus.Contract.UserInfo(&_BaseMasterPlatypus.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_BaseMasterPlatypus *BaseMasterPlatypusCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _BaseMasterPlatypus.Contract.UserInfo(&_BaseMasterPlatypus.CallOpts, arg0, arg1)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) AcceptOwnership() (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.AcceptOwnership(&_BaseMasterPlatypus.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.AcceptOwnership(&_BaseMasterPlatypus.TransactOpts)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 _baseAllocPoint, address _lpToken, address _rewarder) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) Add(opts *bind.TransactOpts, _baseAllocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "add", _baseAllocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 _baseAllocPoint, address _lpToken, address _rewarder) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) Add(_baseAllocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Add(&_BaseMasterPlatypus.TransactOpts, _baseAllocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 _baseAllocPoint, address _lpToken, address _rewarder) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) Add(_baseAllocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Add(&_BaseMasterPlatypus.TransactOpts, _baseAllocPoint, _lpToken, _rewarder)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Deposit(&_BaseMasterPlatypus.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Deposit(&_BaseMasterPlatypus.TransactOpts, _pid, _amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x90210d7e.
//
// Solidity: function depositFor(uint256 _pid, uint256 _amount, address _user) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) DepositFor(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "depositFor", _pid, _amount, _user)
}

// DepositFor is a paid mutator transaction binding the contract method 0x90210d7e.
//
// Solidity: function depositFor(uint256 _pid, uint256 _amount, address _user) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) DepositFor(_pid *big.Int, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.DepositFor(&_BaseMasterPlatypus.TransactOpts, _pid, _amount, _user)
}

// DepositFor is a paid mutator transaction binding the contract method 0x90210d7e.
//
// Solidity: function depositFor(uint256 _pid, uint256 _amount, address _user) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) DepositFor(_pid *big.Int, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.DepositFor(&_BaseMasterPlatypus.TransactOpts, _pid, _amount, _user)
}

// EmergencyPtpWithdraw is a paid mutator transaction binding the contract method 0x7dd38dcc.
//
// Solidity: function emergencyPtpWithdraw() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) EmergencyPtpWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "emergencyPtpWithdraw")
}

// EmergencyPtpWithdraw is a paid mutator transaction binding the contract method 0x7dd38dcc.
//
// Solidity: function emergencyPtpWithdraw() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) EmergencyPtpWithdraw() (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.EmergencyPtpWithdraw(&_BaseMasterPlatypus.TransactOpts)
}

// EmergencyPtpWithdraw is a paid mutator transaction binding the contract method 0x7dd38dcc.
//
// Solidity: function emergencyPtpWithdraw() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) EmergencyPtpWithdraw() (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.EmergencyPtpWithdraw(&_BaseMasterPlatypus.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.EmergencyWithdraw(&_BaseMasterPlatypus.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.EmergencyWithdraw(&_BaseMasterPlatypus.TransactOpts, _pid)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _ptp, uint256 _ptpPerSec, uint256 _startTimestamp) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) Initialize(opts *bind.TransactOpts, _ptp common.Address, _ptpPerSec *big.Int, _startTimestamp *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "initialize", _ptp, _ptpPerSec, _startTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _ptp, uint256 _ptpPerSec, uint256 _startTimestamp) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) Initialize(_ptp common.Address, _ptpPerSec *big.Int, _startTimestamp *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Initialize(&_BaseMasterPlatypus.TransactOpts, _ptp, _ptpPerSec, _startTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _ptp, uint256 _ptpPerSec, uint256 _startTimestamp) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) Initialize(_ptp common.Address, _ptpPerSec *big.Int, _startTimestamp *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Initialize(&_BaseMasterPlatypus.TransactOpts, _ptp, _ptpPerSec, _startTimestamp)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) MassUpdatePools() (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.MassUpdatePools(&_BaseMasterPlatypus.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.MassUpdatePools(&_BaseMasterPlatypus.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0xd93bf4fe.
//
// Solidity: function migrate(uint256[] _pids) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) Migrate(opts *bind.TransactOpts, _pids []*big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "migrate", _pids)
}

// Migrate is a paid mutator transaction binding the contract method 0xd93bf4fe.
//
// Solidity: function migrate(uint256[] _pids) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) Migrate(_pids []*big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Migrate(&_BaseMasterPlatypus.TransactOpts, _pids)
}

// Migrate is a paid mutator transaction binding the contract method 0xd93bf4fe.
//
// Solidity: function migrate(uint256[] _pids) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) Migrate(_pids []*big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Migrate(&_BaseMasterPlatypus.TransactOpts, _pids)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x4ed73d28.
//
// Solidity: function multiClaim(uint256[] _pids) returns(uint256, uint256[], uint256[])
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) MultiClaim(opts *bind.TransactOpts, _pids []*big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "multiClaim", _pids)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x4ed73d28.
//
// Solidity: function multiClaim(uint256[] _pids) returns(uint256, uint256[], uint256[])
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) MultiClaim(_pids []*big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.MultiClaim(&_BaseMasterPlatypus.TransactOpts, _pids)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x4ed73d28.
//
// Solidity: function multiClaim(uint256[] _pids) returns(uint256, uint256[], uint256[])
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) MultiClaim(_pids []*big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.MultiClaim(&_BaseMasterPlatypus.TransactOpts, _pids)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) Pause() (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Pause(&_BaseMasterPlatypus.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) Pause() (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Pause(&_BaseMasterPlatypus.TransactOpts)
}

// ProposeOwner is a paid mutator transaction binding the contract method 0xb5ed298a.
//
// Solidity: function proposeOwner(address newOwner) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) ProposeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "proposeOwner", newOwner)
}

// ProposeOwner is a paid mutator transaction binding the contract method 0xb5ed298a.
//
// Solidity: function proposeOwner(address newOwner) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) ProposeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.ProposeOwner(&_BaseMasterPlatypus.TransactOpts, newOwner)
}

// ProposeOwner is a paid mutator transaction binding the contract method 0xb5ed298a.
//
// Solidity: function proposeOwner(address newOwner) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) ProposeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.ProposeOwner(&_BaseMasterPlatypus.TransactOpts, newOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.RenounceOwnership(&_BaseMasterPlatypus.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.RenounceOwnership(&_BaseMasterPlatypus.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _baseAllocPoint, address _rewarder, bool overwrite) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _baseAllocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "set", _pid, _baseAllocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _baseAllocPoint, address _rewarder, bool overwrite) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) Set(_pid *big.Int, _baseAllocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Set(&_BaseMasterPlatypus.TransactOpts, _pid, _baseAllocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _baseAllocPoint, address _rewarder, bool overwrite) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) Set(_pid *big.Int, _baseAllocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Set(&_BaseMasterPlatypus.TransactOpts, _pid, _baseAllocPoint, _rewarder, overwrite)
}

// SetNewMasterPlatypus is a paid mutator transaction binding the contract method 0x7b261591.
//
// Solidity: function setNewMasterPlatypus(address _newMasterPlatypus) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) SetNewMasterPlatypus(opts *bind.TransactOpts, _newMasterPlatypus common.Address) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "setNewMasterPlatypus", _newMasterPlatypus)
}

// SetNewMasterPlatypus is a paid mutator transaction binding the contract method 0x7b261591.
//
// Solidity: function setNewMasterPlatypus(address _newMasterPlatypus) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) SetNewMasterPlatypus(_newMasterPlatypus common.Address) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.SetNewMasterPlatypus(&_BaseMasterPlatypus.TransactOpts, _newMasterPlatypus)
}

// SetNewMasterPlatypus is a paid mutator transaction binding the contract method 0x7b261591.
//
// Solidity: function setNewMasterPlatypus(address _newMasterPlatypus) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) SetNewMasterPlatypus(_newMasterPlatypus common.Address) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.SetNewMasterPlatypus(&_BaseMasterPlatypus.TransactOpts, _newMasterPlatypus)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) Unpause() (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Unpause(&_BaseMasterPlatypus.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) Unpause() (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Unpause(&_BaseMasterPlatypus.TransactOpts)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _ptpPerSec) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) UpdateEmissionRate(opts *bind.TransactOpts, _ptpPerSec *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "updateEmissionRate", _ptpPerSec)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _ptpPerSec) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) UpdateEmissionRate(_ptpPerSec *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.UpdateEmissionRate(&_BaseMasterPlatypus.TransactOpts, _ptpPerSec)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _ptpPerSec) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) UpdateEmissionRate(_ptpPerSec *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.UpdateEmissionRate(&_BaseMasterPlatypus.TransactOpts, _ptpPerSec)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.UpdatePool(&_BaseMasterPlatypus.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.UpdatePool(&_BaseMasterPlatypus.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Withdraw(&_BaseMasterPlatypus.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_BaseMasterPlatypus *BaseMasterPlatypusTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BaseMasterPlatypus.Contract.Withdraw(&_BaseMasterPlatypus.TransactOpts, _pid, _amount)
}

// BaseMasterPlatypusAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusAddIterator struct {
	Event *BaseMasterPlatypusAdd // Event containing the contract specifics and raw log

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
func (it *BaseMasterPlatypusAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseMasterPlatypusAdd)
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
		it.Event = new(BaseMasterPlatypusAdd)
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
func (it *BaseMasterPlatypusAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseMasterPlatypusAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseMasterPlatypusAdd represents a Add event raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusAdd struct {
	Pid        *big.Int
	AllocPoint *big.Int
	LpToken    common.Address
	Rewarder   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x4b16bd2431ad24dc020ab0e1de7fcb6563dead6a24fb10089d6c23e97a70381f.
//
// Solidity: event Add(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) FilterAdd(opts *bind.FilterOpts, pid []*big.Int, lpToken []common.Address, rewarder []common.Address) (*BaseMasterPlatypusAddIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}
	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.FilterLogs(opts, "Add", pidRule, lpTokenRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypusAddIterator{contract: _BaseMasterPlatypus.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x4b16bd2431ad24dc020ab0e1de7fcb6563dead6a24fb10089d6c23e97a70381f.
//
// Solidity: event Add(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *BaseMasterPlatypusAdd, pid []*big.Int, lpToken []common.Address, rewarder []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}
	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.WatchLogs(opts, "Add", pidRule, lpTokenRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseMasterPlatypusAdd)
				if err := _BaseMasterPlatypus.contract.UnpackLog(event, "Add", log); err != nil {
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

// ParseAdd is a log parse operation binding the contract event 0x4b16bd2431ad24dc020ab0e1de7fcb6563dead6a24fb10089d6c23e97a70381f.
//
// Solidity: event Add(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) ParseAdd(log types.Log) (*BaseMasterPlatypusAdd, error) {
	event := new(BaseMasterPlatypusAdd)
	if err := _BaseMasterPlatypus.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseMasterPlatypusDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusDepositIterator struct {
	Event *BaseMasterPlatypusDeposit // Event containing the contract specifics and raw log

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
func (it *BaseMasterPlatypusDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseMasterPlatypusDeposit)
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
		it.Event = new(BaseMasterPlatypusDeposit)
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
func (it *BaseMasterPlatypusDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseMasterPlatypusDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseMasterPlatypusDeposit represents a Deposit event raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*BaseMasterPlatypusDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypusDepositIterator{contract: _BaseMasterPlatypus.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BaseMasterPlatypusDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseMasterPlatypusDeposit)
				if err := _BaseMasterPlatypus.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) ParseDeposit(log types.Log) (*BaseMasterPlatypusDeposit, error) {
	event := new(BaseMasterPlatypusDeposit)
	if err := _BaseMasterPlatypus.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseMasterPlatypusDepositForIterator is returned from FilterDepositFor and is used to iterate over the raw logs and unpacked data for DepositFor events raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusDepositForIterator struct {
	Event *BaseMasterPlatypusDepositFor // Event containing the contract specifics and raw log

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
func (it *BaseMasterPlatypusDepositForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseMasterPlatypusDepositFor)
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
		it.Event = new(BaseMasterPlatypusDepositFor)
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
func (it *BaseMasterPlatypusDepositForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseMasterPlatypusDepositForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseMasterPlatypusDepositFor represents a DepositFor event raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusDepositFor struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositFor is a free log retrieval operation binding the contract event 0x16f3fbfd4bcc50a5cecb2e53e398a1ad77d89f63288ef540d862b264ed57eb1f.
//
// Solidity: event DepositFor(address indexed user, uint256 indexed pid, uint256 amount)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) FilterDepositFor(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*BaseMasterPlatypusDepositForIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.FilterLogs(opts, "DepositFor", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypusDepositForIterator{contract: _BaseMasterPlatypus.contract, event: "DepositFor", logs: logs, sub: sub}, nil
}

// WatchDepositFor is a free log subscription operation binding the contract event 0x16f3fbfd4bcc50a5cecb2e53e398a1ad77d89f63288ef540d862b264ed57eb1f.
//
// Solidity: event DepositFor(address indexed user, uint256 indexed pid, uint256 amount)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) WatchDepositFor(opts *bind.WatchOpts, sink chan<- *BaseMasterPlatypusDepositFor, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.WatchLogs(opts, "DepositFor", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseMasterPlatypusDepositFor)
				if err := _BaseMasterPlatypus.contract.UnpackLog(event, "DepositFor", log); err != nil {
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

// ParseDepositFor is a log parse operation binding the contract event 0x16f3fbfd4bcc50a5cecb2e53e398a1ad77d89f63288ef540d862b264ed57eb1f.
//
// Solidity: event DepositFor(address indexed user, uint256 indexed pid, uint256 amount)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) ParseDepositFor(log types.Log) (*BaseMasterPlatypusDepositFor, error) {
	event := new(BaseMasterPlatypusDepositFor)
	if err := _BaseMasterPlatypus.contract.UnpackLog(event, "DepositFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseMasterPlatypusEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusEmergencyWithdrawIterator struct {
	Event *BaseMasterPlatypusEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *BaseMasterPlatypusEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseMasterPlatypusEmergencyWithdraw)
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
		it.Event = new(BaseMasterPlatypusEmergencyWithdraw)
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
func (it *BaseMasterPlatypusEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseMasterPlatypusEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseMasterPlatypusEmergencyWithdraw represents a EmergencyWithdraw event raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*BaseMasterPlatypusEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypusEmergencyWithdrawIterator{contract: _BaseMasterPlatypus.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *BaseMasterPlatypusEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseMasterPlatypusEmergencyWithdraw)
				if err := _BaseMasterPlatypus.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) ParseEmergencyWithdraw(log types.Log) (*BaseMasterPlatypusEmergencyWithdraw, error) {
	event := new(BaseMasterPlatypusEmergencyWithdraw)
	if err := _BaseMasterPlatypus.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseMasterPlatypusHarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusHarvestIterator struct {
	Event *BaseMasterPlatypusHarvest // Event containing the contract specifics and raw log

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
func (it *BaseMasterPlatypusHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseMasterPlatypusHarvest)
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
		it.Event = new(BaseMasterPlatypusHarvest)
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
func (it *BaseMasterPlatypusHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseMasterPlatypusHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseMasterPlatypusHarvest represents a Harvest event raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusHarvest struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) FilterHarvest(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*BaseMasterPlatypusHarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.FilterLogs(opts, "Harvest", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypusHarvestIterator{contract: _BaseMasterPlatypus.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *BaseMasterPlatypusHarvest, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.WatchLogs(opts, "Harvest", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseMasterPlatypusHarvest)
				if err := _BaseMasterPlatypus.contract.UnpackLog(event, "Harvest", log); err != nil {
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

// ParseHarvest is a log parse operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) ParseHarvest(log types.Log) (*BaseMasterPlatypusHarvest, error) {
	event := new(BaseMasterPlatypusHarvest)
	if err := _BaseMasterPlatypus.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseMasterPlatypusOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusOwnershipTransferredIterator struct {
	Event *BaseMasterPlatypusOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BaseMasterPlatypusOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseMasterPlatypusOwnershipTransferred)
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
		it.Event = new(BaseMasterPlatypusOwnershipTransferred)
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
func (it *BaseMasterPlatypusOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseMasterPlatypusOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseMasterPlatypusOwnershipTransferred represents a OwnershipTransferred event raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BaseMasterPlatypusOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypusOwnershipTransferredIterator{contract: _BaseMasterPlatypus.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BaseMasterPlatypusOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseMasterPlatypusOwnershipTransferred)
				if err := _BaseMasterPlatypus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) ParseOwnershipTransferred(log types.Log) (*BaseMasterPlatypusOwnershipTransferred, error) {
	event := new(BaseMasterPlatypusOwnershipTransferred)
	if err := _BaseMasterPlatypus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseMasterPlatypusPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusPausedIterator struct {
	Event *BaseMasterPlatypusPaused // Event containing the contract specifics and raw log

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
func (it *BaseMasterPlatypusPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseMasterPlatypusPaused)
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
		it.Event = new(BaseMasterPlatypusPaused)
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
func (it *BaseMasterPlatypusPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseMasterPlatypusPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseMasterPlatypusPaused represents a Paused event raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) FilterPaused(opts *bind.FilterOpts) (*BaseMasterPlatypusPausedIterator, error) {

	logs, sub, err := _BaseMasterPlatypus.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypusPausedIterator{contract: _BaseMasterPlatypus.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BaseMasterPlatypusPaused) (event.Subscription, error) {

	logs, sub, err := _BaseMasterPlatypus.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseMasterPlatypusPaused)
				if err := _BaseMasterPlatypus.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) ParsePaused(log types.Log) (*BaseMasterPlatypusPaused, error) {
	event := new(BaseMasterPlatypusPaused)
	if err := _BaseMasterPlatypus.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseMasterPlatypusSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusSetIterator struct {
	Event *BaseMasterPlatypusSet // Event containing the contract specifics and raw log

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
func (it *BaseMasterPlatypusSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseMasterPlatypusSet)
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
		it.Event = new(BaseMasterPlatypusSet)
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
func (it *BaseMasterPlatypusSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseMasterPlatypusSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseMasterPlatypusSet represents a Set event raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusSet struct {
	Pid        *big.Int
	AllocPoint *big.Int
	Rewarder   common.Address
	Overwrite  bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0xa54644aae5c48c5971516f334e4fe8ecbc7930e23f34877d4203c6551e67ffaa.
//
// Solidity: event Set(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) FilterSet(opts *bind.FilterOpts, pid []*big.Int, rewarder []common.Address) (*BaseMasterPlatypusSetIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.FilterLogs(opts, "Set", pidRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypusSetIterator{contract: _BaseMasterPlatypus.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0xa54644aae5c48c5971516f334e4fe8ecbc7930e23f34877d4203c6551e67ffaa.
//
// Solidity: event Set(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *BaseMasterPlatypusSet, pid []*big.Int, rewarder []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.WatchLogs(opts, "Set", pidRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseMasterPlatypusSet)
				if err := _BaseMasterPlatypus.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0xa54644aae5c48c5971516f334e4fe8ecbc7930e23f34877d4203c6551e67ffaa.
//
// Solidity: event Set(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) ParseSet(log types.Log) (*BaseMasterPlatypusSet, error) {
	event := new(BaseMasterPlatypusSet)
	if err := _BaseMasterPlatypus.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseMasterPlatypusUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusUnpausedIterator struct {
	Event *BaseMasterPlatypusUnpaused // Event containing the contract specifics and raw log

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
func (it *BaseMasterPlatypusUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseMasterPlatypusUnpaused)
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
		it.Event = new(BaseMasterPlatypusUnpaused)
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
func (it *BaseMasterPlatypusUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseMasterPlatypusUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseMasterPlatypusUnpaused represents a Unpaused event raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BaseMasterPlatypusUnpausedIterator, error) {

	logs, sub, err := _BaseMasterPlatypus.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypusUnpausedIterator{contract: _BaseMasterPlatypus.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BaseMasterPlatypusUnpaused) (event.Subscription, error) {

	logs, sub, err := _BaseMasterPlatypus.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseMasterPlatypusUnpaused)
				if err := _BaseMasterPlatypus.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) ParseUnpaused(log types.Log) (*BaseMasterPlatypusUnpaused, error) {
	event := new(BaseMasterPlatypusUnpaused)
	if err := _BaseMasterPlatypus.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseMasterPlatypusUpdateEmissionRateIterator is returned from FilterUpdateEmissionRate and is used to iterate over the raw logs and unpacked data for UpdateEmissionRate events raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusUpdateEmissionRateIterator struct {
	Event *BaseMasterPlatypusUpdateEmissionRate // Event containing the contract specifics and raw log

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
func (it *BaseMasterPlatypusUpdateEmissionRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseMasterPlatypusUpdateEmissionRate)
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
		it.Event = new(BaseMasterPlatypusUpdateEmissionRate)
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
func (it *BaseMasterPlatypusUpdateEmissionRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseMasterPlatypusUpdateEmissionRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseMasterPlatypusUpdateEmissionRate represents a UpdateEmissionRate event raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusUpdateEmissionRate struct {
	User      common.Address
	PtpPerSec *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateEmissionRate is a free log retrieval operation binding the contract event 0xe2492e003bbe8afa53088b406f0c1cb5d9e280370fc72a74cf116ffd343c4053.
//
// Solidity: event UpdateEmissionRate(address indexed user, uint256 ptpPerSec)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) FilterUpdateEmissionRate(opts *bind.FilterOpts, user []common.Address) (*BaseMasterPlatypusUpdateEmissionRateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.FilterLogs(opts, "UpdateEmissionRate", userRule)
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypusUpdateEmissionRateIterator{contract: _BaseMasterPlatypus.contract, event: "UpdateEmissionRate", logs: logs, sub: sub}, nil
}

// WatchUpdateEmissionRate is a free log subscription operation binding the contract event 0xe2492e003bbe8afa53088b406f0c1cb5d9e280370fc72a74cf116ffd343c4053.
//
// Solidity: event UpdateEmissionRate(address indexed user, uint256 ptpPerSec)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) WatchUpdateEmissionRate(opts *bind.WatchOpts, sink chan<- *BaseMasterPlatypusUpdateEmissionRate, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.WatchLogs(opts, "UpdateEmissionRate", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseMasterPlatypusUpdateEmissionRate)
				if err := _BaseMasterPlatypus.contract.UnpackLog(event, "UpdateEmissionRate", log); err != nil {
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

// ParseUpdateEmissionRate is a log parse operation binding the contract event 0xe2492e003bbe8afa53088b406f0c1cb5d9e280370fc72a74cf116ffd343c4053.
//
// Solidity: event UpdateEmissionRate(address indexed user, uint256 ptpPerSec)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) ParseUpdateEmissionRate(log types.Log) (*BaseMasterPlatypusUpdateEmissionRate, error) {
	event := new(BaseMasterPlatypusUpdateEmissionRate)
	if err := _BaseMasterPlatypus.contract.UnpackLog(event, "UpdateEmissionRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseMasterPlatypusUpdatePoolIterator is returned from FilterUpdatePool and is used to iterate over the raw logs and unpacked data for UpdatePool events raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusUpdatePoolIterator struct {
	Event *BaseMasterPlatypusUpdatePool // Event containing the contract specifics and raw log

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
func (it *BaseMasterPlatypusUpdatePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseMasterPlatypusUpdatePool)
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
		it.Event = new(BaseMasterPlatypusUpdatePool)
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
func (it *BaseMasterPlatypusUpdatePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseMasterPlatypusUpdatePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseMasterPlatypusUpdatePool represents a UpdatePool event raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusUpdatePool struct {
	Pid                 *big.Int
	LastRewardTimestamp *big.Int
	LpSupply            *big.Int
	AccPtpPerShare      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdatePool is a free log retrieval operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardTimestamp, uint256 lpSupply, uint256 accPtpPerShare)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) FilterUpdatePool(opts *bind.FilterOpts, pid []*big.Int) (*BaseMasterPlatypusUpdatePoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.FilterLogs(opts, "UpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypusUpdatePoolIterator{contract: _BaseMasterPlatypus.contract, event: "UpdatePool", logs: logs, sub: sub}, nil
}

// WatchUpdatePool is a free log subscription operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardTimestamp, uint256 lpSupply, uint256 accPtpPerShare)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) WatchUpdatePool(opts *bind.WatchOpts, sink chan<- *BaseMasterPlatypusUpdatePool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.WatchLogs(opts, "UpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseMasterPlatypusUpdatePool)
				if err := _BaseMasterPlatypus.contract.UnpackLog(event, "UpdatePool", log); err != nil {
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

// ParseUpdatePool is a log parse operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardTimestamp, uint256 lpSupply, uint256 accPtpPerShare)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) ParseUpdatePool(log types.Log) (*BaseMasterPlatypusUpdatePool, error) {
	event := new(BaseMasterPlatypusUpdatePool)
	if err := _BaseMasterPlatypus.contract.UnpackLog(event, "UpdatePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseMasterPlatypusWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusWithdrawIterator struct {
	Event *BaseMasterPlatypusWithdraw // Event containing the contract specifics and raw log

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
func (it *BaseMasterPlatypusWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseMasterPlatypusWithdraw)
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
		it.Event = new(BaseMasterPlatypusWithdraw)
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
func (it *BaseMasterPlatypusWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseMasterPlatypusWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseMasterPlatypusWithdraw represents a Withdraw event raised by the BaseMasterPlatypus contract.
type BaseMasterPlatypusWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*BaseMasterPlatypusWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &BaseMasterPlatypusWithdrawIterator{contract: _BaseMasterPlatypus.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BaseMasterPlatypusWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BaseMasterPlatypus.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseMasterPlatypusWithdraw)
				if err := _BaseMasterPlatypus.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_BaseMasterPlatypus *BaseMasterPlatypusFilterer) ParseWithdraw(log types.Log) (*BaseMasterPlatypusWithdraw, error) {
	event := new(BaseMasterPlatypusWithdraw)
	if err := _BaseMasterPlatypus.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
