// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakingpool

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

// UserStakingPoolABI is the input ABI used to generate the binding from.
const UserStakingPoolABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserStaking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawalWaitTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardWaitTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingReward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"total\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"depositedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"claimedAt\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numAddresses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lrcAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"claimedAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalStaking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_CLAIM_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_protocolFeeVaultAddress\",\"type\":\"address\"}],\"name\":\"setProtocolFeeVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_WITHDRAW_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"protocolFeeVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"depositedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"claimedAt\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lrcAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeVaultAddress\",\"type\":\"address\"}],\"name\":\"ProtocolFeeVaultChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LRCStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LRCWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LRCRewarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// UserStakingPool is an auto generated Go binding around an Ethereum contract.
type UserStakingPool struct {
	UserStakingPoolCaller     // Read-only binding to the contract
	UserStakingPoolTransactor // Write-only binding to the contract
	UserStakingPoolFilterer   // Log filterer for contract events
}

// UserStakingPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserStakingPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserStakingPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserStakingPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserStakingPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserStakingPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserStakingPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserStakingPoolSession struct {
	Contract     *UserStakingPool  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserStakingPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserStakingPoolCallerSession struct {
	Contract *UserStakingPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UserStakingPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserStakingPoolTransactorSession struct {
	Contract     *UserStakingPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UserStakingPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserStakingPoolRaw struct {
	Contract *UserStakingPool // Generic contract binding to access the raw methods on
}

// UserStakingPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserStakingPoolCallerRaw struct {
	Contract *UserStakingPoolCaller // Generic read-only contract binding to access the raw methods on
}

// UserStakingPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserStakingPoolTransactorRaw struct {
	Contract *UserStakingPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserStakingPool creates a new instance of UserStakingPool, bound to a specific deployed contract.
func NewUserStakingPool(address common.Address, backend bind.ContractBackend) (*UserStakingPool, error) {
	contract, err := bindUserStakingPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserStakingPool{UserStakingPoolCaller: UserStakingPoolCaller{contract: contract}, UserStakingPoolTransactor: UserStakingPoolTransactor{contract: contract}, UserStakingPoolFilterer: UserStakingPoolFilterer{contract: contract}}, nil
}

// NewUserStakingPoolCaller creates a new read-only instance of UserStakingPool, bound to a specific deployed contract.
func NewUserStakingPoolCaller(address common.Address, caller bind.ContractCaller) (*UserStakingPoolCaller, error) {
	contract, err := bindUserStakingPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserStakingPoolCaller{contract: contract}, nil
}

// NewUserStakingPoolTransactor creates a new write-only instance of UserStakingPool, bound to a specific deployed contract.
func NewUserStakingPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*UserStakingPoolTransactor, error) {
	contract, err := bindUserStakingPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserStakingPoolTransactor{contract: contract}, nil
}

// NewUserStakingPoolFilterer creates a new log filterer instance of UserStakingPool, bound to a specific deployed contract.
func NewUserStakingPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*UserStakingPoolFilterer, error) {
	contract, err := bindUserStakingPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserStakingPoolFilterer{contract: contract}, nil
}

// bindUserStakingPool binds a generic wrapper to an already deployed contract.
func bindUserStakingPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserStakingPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserStakingPool *UserStakingPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserStakingPool.Contract.UserStakingPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserStakingPool *UserStakingPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserStakingPool.Contract.UserStakingPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserStakingPool *UserStakingPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserStakingPool.Contract.UserStakingPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserStakingPool *UserStakingPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserStakingPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserStakingPool *UserStakingPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserStakingPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserStakingPool *UserStakingPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserStakingPool.Contract.contract.Transact(opts, method, params...)
}

// MINCLAIMDELAY is a free data retrieval call binding the contract method 0x61605b92.
//
// Solidity: function MIN_CLAIM_DELAY() view returns(uint256)
func (_UserStakingPool *UserStakingPoolCaller) MINCLAIMDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserStakingPool.contract.Call(opts, &out, "MIN_CLAIM_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINCLAIMDELAY is a free data retrieval call binding the contract method 0x61605b92.
//
// Solidity: function MIN_CLAIM_DELAY() view returns(uint256)
func (_UserStakingPool *UserStakingPoolSession) MINCLAIMDELAY() (*big.Int, error) {
	return _UserStakingPool.Contract.MINCLAIMDELAY(&_UserStakingPool.CallOpts)
}

// MINCLAIMDELAY is a free data retrieval call binding the contract method 0x61605b92.
//
// Solidity: function MIN_CLAIM_DELAY() view returns(uint256)
func (_UserStakingPool *UserStakingPoolCallerSession) MINCLAIMDELAY() (*big.Int, error) {
	return _UserStakingPool.Contract.MINCLAIMDELAY(&_UserStakingPool.CallOpts)
}

// MINWITHDRAWDELAY is a free data retrieval call binding the contract method 0xae159439.
//
// Solidity: function MIN_WITHDRAW_DELAY() view returns(uint256)
func (_UserStakingPool *UserStakingPoolCaller) MINWITHDRAWDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserStakingPool.contract.Call(opts, &out, "MIN_WITHDRAW_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWITHDRAWDELAY is a free data retrieval call binding the contract method 0xae159439.
//
// Solidity: function MIN_WITHDRAW_DELAY() view returns(uint256)
func (_UserStakingPool *UserStakingPoolSession) MINWITHDRAWDELAY() (*big.Int, error) {
	return _UserStakingPool.Contract.MINWITHDRAWDELAY(&_UserStakingPool.CallOpts)
}

// MINWITHDRAWDELAY is a free data retrieval call binding the contract method 0xae159439.
//
// Solidity: function MIN_WITHDRAW_DELAY() view returns(uint256)
func (_UserStakingPool *UserStakingPoolCallerSession) MINWITHDRAWDELAY() (*big.Int, error) {
	return _UserStakingPool.Contract.MINWITHDRAWDELAY(&_UserStakingPool.CallOpts)
}

// GetTotalStaking is a free data retrieval call binding the contract method 0x58d57545.
//
// Solidity: function getTotalStaking() view returns(uint256)
func (_UserStakingPool *UserStakingPoolCaller) GetTotalStaking(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserStakingPool.contract.Call(opts, &out, "getTotalStaking")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStaking is a free data retrieval call binding the contract method 0x58d57545.
//
// Solidity: function getTotalStaking() view returns(uint256)
func (_UserStakingPool *UserStakingPoolSession) GetTotalStaking() (*big.Int, error) {
	return _UserStakingPool.Contract.GetTotalStaking(&_UserStakingPool.CallOpts)
}

// GetTotalStaking is a free data retrieval call binding the contract method 0x58d57545.
//
// Solidity: function getTotalStaking() view returns(uint256)
func (_UserStakingPool *UserStakingPoolCallerSession) GetTotalStaking() (*big.Int, error) {
	return _UserStakingPool.Contract.GetTotalStaking(&_UserStakingPool.CallOpts)
}

// GetUserStaking is a free data retrieval call binding the contract method 0x0c0fd549.
//
// Solidity: function getUserStaking(address user) view returns(uint256 withdrawalWaitTime, uint256 rewardWaitTime, uint256 balance, uint256 pendingReward)
func (_UserStakingPool *UserStakingPoolCaller) GetUserStaking(opts *bind.CallOpts, user common.Address) (struct {
	WithdrawalWaitTime *big.Int
	RewardWaitTime     *big.Int
	Balance            *big.Int
	PendingReward      *big.Int
}, error) {
	var out []interface{}
	err := _UserStakingPool.contract.Call(opts, &out, "getUserStaking", user)

	outstruct := new(struct {
		WithdrawalWaitTime *big.Int
		RewardWaitTime     *big.Int
		Balance            *big.Int
		PendingReward      *big.Int
	})

	outstruct.WithdrawalWaitTime = out[0].(*big.Int)
	outstruct.RewardWaitTime = out[1].(*big.Int)
	outstruct.Balance = out[2].(*big.Int)
	outstruct.PendingReward = out[3].(*big.Int)

	return *outstruct, err

}

// GetUserStaking is a free data retrieval call binding the contract method 0x0c0fd549.
//
// Solidity: function getUserStaking(address user) view returns(uint256 withdrawalWaitTime, uint256 rewardWaitTime, uint256 balance, uint256 pendingReward)
func (_UserStakingPool *UserStakingPoolSession) GetUserStaking(user common.Address) (struct {
	WithdrawalWaitTime *big.Int
	RewardWaitTime     *big.Int
	Balance            *big.Int
	PendingReward      *big.Int
}, error) {
	return _UserStakingPool.Contract.GetUserStaking(&_UserStakingPool.CallOpts, user)
}

// GetUserStaking is a free data retrieval call binding the contract method 0x0c0fd549.
//
// Solidity: function getUserStaking(address user) view returns(uint256 withdrawalWaitTime, uint256 rewardWaitTime, uint256 balance, uint256 pendingReward)
func (_UserStakingPool *UserStakingPoolCallerSession) GetUserStaking(user common.Address) (struct {
	WithdrawalWaitTime *big.Int
	RewardWaitTime     *big.Int
	Balance            *big.Int
	PendingReward      *big.Int
}, error) {
	return _UserStakingPool.Contract.GetUserStaking(&_UserStakingPool.CallOpts, user)
}

// LrcAddress is a free data retrieval call binding the contract method 0x3d6cf722.
//
// Solidity: function lrcAddress() view returns(address)
func (_UserStakingPool *UserStakingPoolCaller) LrcAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserStakingPool.contract.Call(opts, &out, "lrcAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LrcAddress is a free data retrieval call binding the contract method 0x3d6cf722.
//
// Solidity: function lrcAddress() view returns(address)
func (_UserStakingPool *UserStakingPoolSession) LrcAddress() (common.Address, error) {
	return _UserStakingPool.Contract.LrcAddress(&_UserStakingPool.CallOpts)
}

// LrcAddress is a free data retrieval call binding the contract method 0x3d6cf722.
//
// Solidity: function lrcAddress() view returns(address)
func (_UserStakingPool *UserStakingPoolCallerSession) LrcAddress() (common.Address, error) {
	return _UserStakingPool.Contract.LrcAddress(&_UserStakingPool.CallOpts)
}

// NumAddresses is a free data retrieval call binding the contract method 0x2f8646d6.
//
// Solidity: function numAddresses() view returns(uint256)
func (_UserStakingPool *UserStakingPoolCaller) NumAddresses(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserStakingPool.contract.Call(opts, &out, "numAddresses")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumAddresses is a free data retrieval call binding the contract method 0x2f8646d6.
//
// Solidity: function numAddresses() view returns(uint256)
func (_UserStakingPool *UserStakingPoolSession) NumAddresses() (*big.Int, error) {
	return _UserStakingPool.Contract.NumAddresses(&_UserStakingPool.CallOpts)
}

// NumAddresses is a free data retrieval call binding the contract method 0x2f8646d6.
//
// Solidity: function numAddresses() view returns(uint256)
func (_UserStakingPool *UserStakingPoolCallerSession) NumAddresses() (*big.Int, error) {
	return _UserStakingPool.Contract.NumAddresses(&_UserStakingPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UserStakingPool *UserStakingPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserStakingPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UserStakingPool *UserStakingPoolSession) Owner() (common.Address, error) {
	return _UserStakingPool.Contract.Owner(&_UserStakingPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UserStakingPool *UserStakingPoolCallerSession) Owner() (common.Address, error) {
	return _UserStakingPool.Contract.Owner(&_UserStakingPool.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UserStakingPool *UserStakingPoolCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserStakingPool.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UserStakingPool *UserStakingPoolSession) PendingOwner() (common.Address, error) {
	return _UserStakingPool.Contract.PendingOwner(&_UserStakingPool.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UserStakingPool *UserStakingPoolCallerSession) PendingOwner() (common.Address, error) {
	return _UserStakingPool.Contract.PendingOwner(&_UserStakingPool.CallOpts)
}

// ProtocolFeeVaultAddress is a free data retrieval call binding the contract method 0xc730eca4.
//
// Solidity: function protocolFeeVaultAddress() view returns(address)
func (_UserStakingPool *UserStakingPoolCaller) ProtocolFeeVaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserStakingPool.contract.Call(opts, &out, "protocolFeeVaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeVaultAddress is a free data retrieval call binding the contract method 0xc730eca4.
//
// Solidity: function protocolFeeVaultAddress() view returns(address)
func (_UserStakingPool *UserStakingPoolSession) ProtocolFeeVaultAddress() (common.Address, error) {
	return _UserStakingPool.Contract.ProtocolFeeVaultAddress(&_UserStakingPool.CallOpts)
}

// ProtocolFeeVaultAddress is a free data retrieval call binding the contract method 0xc730eca4.
//
// Solidity: function protocolFeeVaultAddress() view returns(address)
func (_UserStakingPool *UserStakingPoolCallerSession) ProtocolFeeVaultAddress() (common.Address, error) {
	return _UserStakingPool.Contract.ProtocolFeeVaultAddress(&_UserStakingPool.CallOpts)
}

// Stakings is a free data retrieval call binding the contract method 0xdc6e13e1.
//
// Solidity: function stakings(address ) view returns(uint256 balance, uint64 depositedAt, uint64 claimedAt)
func (_UserStakingPool *UserStakingPoolCaller) Stakings(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance     *big.Int
	DepositedAt uint64
	ClaimedAt   uint64
}, error) {
	var out []interface{}
	err := _UserStakingPool.contract.Call(opts, &out, "stakings", arg0)

	outstruct := new(struct {
		Balance     *big.Int
		DepositedAt uint64
		ClaimedAt   uint64
	})

	outstruct.Balance = out[0].(*big.Int)
	outstruct.DepositedAt = out[1].(uint64)
	outstruct.ClaimedAt = out[2].(uint64)

	return *outstruct, err

}

// Stakings is a free data retrieval call binding the contract method 0xdc6e13e1.
//
// Solidity: function stakings(address ) view returns(uint256 balance, uint64 depositedAt, uint64 claimedAt)
func (_UserStakingPool *UserStakingPoolSession) Stakings(arg0 common.Address) (struct {
	Balance     *big.Int
	DepositedAt uint64
	ClaimedAt   uint64
}, error) {
	return _UserStakingPool.Contract.Stakings(&_UserStakingPool.CallOpts, arg0)
}

// Stakings is a free data retrieval call binding the contract method 0xdc6e13e1.
//
// Solidity: function stakings(address ) view returns(uint256 balance, uint64 depositedAt, uint64 claimedAt)
func (_UserStakingPool *UserStakingPoolCallerSession) Stakings(arg0 common.Address) (struct {
	Balance     *big.Int
	DepositedAt uint64
	ClaimedAt   uint64
}, error) {
	return _UserStakingPool.Contract.Stakings(&_UserStakingPool.CallOpts, arg0)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256 balance, uint64 depositedAt, uint64 claimedAt)
func (_UserStakingPool *UserStakingPoolCaller) Total(opts *bind.CallOpts) (struct {
	Balance     *big.Int
	DepositedAt uint64
	ClaimedAt   uint64
}, error) {
	var out []interface{}
	err := _UserStakingPool.contract.Call(opts, &out, "total")

	outstruct := new(struct {
		Balance     *big.Int
		DepositedAt uint64
		ClaimedAt   uint64
	})

	outstruct.Balance = out[0].(*big.Int)
	outstruct.DepositedAt = out[1].(uint64)
	outstruct.ClaimedAt = out[2].(uint64)

	return *outstruct, err

}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256 balance, uint64 depositedAt, uint64 claimedAt)
func (_UserStakingPool *UserStakingPoolSession) Total() (struct {
	Balance     *big.Int
	DepositedAt uint64
	ClaimedAt   uint64
}, error) {
	return _UserStakingPool.Contract.Total(&_UserStakingPool.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256 balance, uint64 depositedAt, uint64 claimedAt)
func (_UserStakingPool *UserStakingPoolCallerSession) Total() (struct {
	Balance     *big.Int
	DepositedAt uint64
	ClaimedAt   uint64
}, error) {
	return _UserStakingPool.Contract.Total(&_UserStakingPool.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256 claimedAmount)
func (_UserStakingPool *UserStakingPoolTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserStakingPool.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256 claimedAmount)
func (_UserStakingPool *UserStakingPoolSession) Claim() (*types.Transaction, error) {
	return _UserStakingPool.Contract.Claim(&_UserStakingPool.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256 claimedAmount)
func (_UserStakingPool *UserStakingPoolTransactorSession) Claim() (*types.Transaction, error) {
	return _UserStakingPool.Contract.Claim(&_UserStakingPool.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_UserStakingPool *UserStakingPoolTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserStakingPool.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_UserStakingPool *UserStakingPoolSession) ClaimOwnership() (*types.Transaction, error) {
	return _UserStakingPool.Contract.ClaimOwnership(&_UserStakingPool.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_UserStakingPool *UserStakingPoolTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _UserStakingPool.Contract.ClaimOwnership(&_UserStakingPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UserStakingPool *UserStakingPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserStakingPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UserStakingPool *UserStakingPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _UserStakingPool.Contract.RenounceOwnership(&_UserStakingPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UserStakingPool *UserStakingPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UserStakingPool.Contract.RenounceOwnership(&_UserStakingPool.TransactOpts)
}

// SetProtocolFeeVault is a paid mutator transaction binding the contract method 0x88e95c47.
//
// Solidity: function setProtocolFeeVault(address _protocolFeeVaultAddress) returns()
func (_UserStakingPool *UserStakingPoolTransactor) SetProtocolFeeVault(opts *bind.TransactOpts, _protocolFeeVaultAddress common.Address) (*types.Transaction, error) {
	return _UserStakingPool.contract.Transact(opts, "setProtocolFeeVault", _protocolFeeVaultAddress)
}

// SetProtocolFeeVault is a paid mutator transaction binding the contract method 0x88e95c47.
//
// Solidity: function setProtocolFeeVault(address _protocolFeeVaultAddress) returns()
func (_UserStakingPool *UserStakingPoolSession) SetProtocolFeeVault(_protocolFeeVaultAddress common.Address) (*types.Transaction, error) {
	return _UserStakingPool.Contract.SetProtocolFeeVault(&_UserStakingPool.TransactOpts, _protocolFeeVaultAddress)
}

// SetProtocolFeeVault is a paid mutator transaction binding the contract method 0x88e95c47.
//
// Solidity: function setProtocolFeeVault(address _protocolFeeVaultAddress) returns()
func (_UserStakingPool *UserStakingPoolTransactorSession) SetProtocolFeeVault(_protocolFeeVaultAddress common.Address) (*types.Transaction, error) {
	return _UserStakingPool.Contract.SetProtocolFeeVault(&_UserStakingPool.TransactOpts, _protocolFeeVaultAddress)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_UserStakingPool *UserStakingPoolTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _UserStakingPool.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_UserStakingPool *UserStakingPoolSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _UserStakingPool.Contract.Stake(&_UserStakingPool.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_UserStakingPool *UserStakingPoolTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _UserStakingPool.Contract.Stake(&_UserStakingPool.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UserStakingPool *UserStakingPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UserStakingPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UserStakingPool *UserStakingPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UserStakingPool.Contract.TransferOwnership(&_UserStakingPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UserStakingPool *UserStakingPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UserStakingPool.Contract.TransferOwnership(&_UserStakingPool.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_UserStakingPool *UserStakingPoolTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _UserStakingPool.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_UserStakingPool *UserStakingPoolSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _UserStakingPool.Contract.Withdraw(&_UserStakingPool.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_UserStakingPool *UserStakingPoolTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _UserStakingPool.Contract.Withdraw(&_UserStakingPool.TransactOpts, amount)
}

// UserStakingPoolLRCRewardedIterator is returned from FilterLRCRewarded and is used to iterate over the raw logs and unpacked data for LRCRewarded events raised by the UserStakingPool contract.
type UserStakingPoolLRCRewardedIterator struct {
	Event *UserStakingPoolLRCRewarded // Event containing the contract specifics and raw log

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
func (it *UserStakingPoolLRCRewardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserStakingPoolLRCRewarded)
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
		it.Event = new(UserStakingPoolLRCRewarded)
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
func (it *UserStakingPoolLRCRewardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserStakingPoolLRCRewardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserStakingPoolLRCRewarded represents a LRCRewarded event raised by the UserStakingPool contract.
type UserStakingPoolLRCRewarded struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLRCRewarded is a free log retrieval operation binding the contract event 0xfaa4a7e6a2ecb2217533f3081ad082d42261efe260c8842dd32a5ebb72011b78.
//
// Solidity: event LRCRewarded(address indexed user, uint256 amount)
func (_UserStakingPool *UserStakingPoolFilterer) FilterLRCRewarded(opts *bind.FilterOpts, user []common.Address) (*UserStakingPoolLRCRewardedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UserStakingPool.contract.FilterLogs(opts, "LRCRewarded", userRule)
	if err != nil {
		return nil, err
	}
	return &UserStakingPoolLRCRewardedIterator{contract: _UserStakingPool.contract, event: "LRCRewarded", logs: logs, sub: sub}, nil
}

// WatchLRCRewarded is a free log subscription operation binding the contract event 0xfaa4a7e6a2ecb2217533f3081ad082d42261efe260c8842dd32a5ebb72011b78.
//
// Solidity: event LRCRewarded(address indexed user, uint256 amount)
func (_UserStakingPool *UserStakingPoolFilterer) WatchLRCRewarded(opts *bind.WatchOpts, sink chan<- *UserStakingPoolLRCRewarded, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UserStakingPool.contract.WatchLogs(opts, "LRCRewarded", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserStakingPoolLRCRewarded)
				if err := _UserStakingPool.contract.UnpackLog(event, "LRCRewarded", log); err != nil {
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

// ParseLRCRewarded is a log parse operation binding the contract event 0xfaa4a7e6a2ecb2217533f3081ad082d42261efe260c8842dd32a5ebb72011b78.
//
// Solidity: event LRCRewarded(address indexed user, uint256 amount)
func (_UserStakingPool *UserStakingPoolFilterer) ParseLRCRewarded(log types.Log) (*UserStakingPoolLRCRewarded, error) {
	event := new(UserStakingPoolLRCRewarded)
	if err := _UserStakingPool.contract.UnpackLog(event, "LRCRewarded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UserStakingPoolLRCStakedIterator is returned from FilterLRCStaked and is used to iterate over the raw logs and unpacked data for LRCStaked events raised by the UserStakingPool contract.
type UserStakingPoolLRCStakedIterator struct {
	Event *UserStakingPoolLRCStaked // Event containing the contract specifics and raw log

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
func (it *UserStakingPoolLRCStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserStakingPoolLRCStaked)
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
		it.Event = new(UserStakingPoolLRCStaked)
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
func (it *UserStakingPoolLRCStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserStakingPoolLRCStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserStakingPoolLRCStaked represents a LRCStaked event raised by the UserStakingPool contract.
type UserStakingPoolLRCStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLRCStaked is a free log retrieval operation binding the contract event 0x7bcf38427a7f7fc25e271d30f03522e51794ec2ba24df0703481c2b7222ced4e.
//
// Solidity: event LRCStaked(address indexed user, uint256 amount)
func (_UserStakingPool *UserStakingPoolFilterer) FilterLRCStaked(opts *bind.FilterOpts, user []common.Address) (*UserStakingPoolLRCStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UserStakingPool.contract.FilterLogs(opts, "LRCStaked", userRule)
	if err != nil {
		return nil, err
	}
	return &UserStakingPoolLRCStakedIterator{contract: _UserStakingPool.contract, event: "LRCStaked", logs: logs, sub: sub}, nil
}

// WatchLRCStaked is a free log subscription operation binding the contract event 0x7bcf38427a7f7fc25e271d30f03522e51794ec2ba24df0703481c2b7222ced4e.
//
// Solidity: event LRCStaked(address indexed user, uint256 amount)
func (_UserStakingPool *UserStakingPoolFilterer) WatchLRCStaked(opts *bind.WatchOpts, sink chan<- *UserStakingPoolLRCStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UserStakingPool.contract.WatchLogs(opts, "LRCStaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserStakingPoolLRCStaked)
				if err := _UserStakingPool.contract.UnpackLog(event, "LRCStaked", log); err != nil {
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

// ParseLRCStaked is a log parse operation binding the contract event 0x7bcf38427a7f7fc25e271d30f03522e51794ec2ba24df0703481c2b7222ced4e.
//
// Solidity: event LRCStaked(address indexed user, uint256 amount)
func (_UserStakingPool *UserStakingPoolFilterer) ParseLRCStaked(log types.Log) (*UserStakingPoolLRCStaked, error) {
	event := new(UserStakingPoolLRCStaked)
	if err := _UserStakingPool.contract.UnpackLog(event, "LRCStaked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UserStakingPoolLRCWithdrawnIterator is returned from FilterLRCWithdrawn and is used to iterate over the raw logs and unpacked data for LRCWithdrawn events raised by the UserStakingPool contract.
type UserStakingPoolLRCWithdrawnIterator struct {
	Event *UserStakingPoolLRCWithdrawn // Event containing the contract specifics and raw log

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
func (it *UserStakingPoolLRCWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserStakingPoolLRCWithdrawn)
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
		it.Event = new(UserStakingPoolLRCWithdrawn)
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
func (it *UserStakingPoolLRCWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserStakingPoolLRCWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserStakingPoolLRCWithdrawn represents a LRCWithdrawn event raised by the UserStakingPool contract.
type UserStakingPoolLRCWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLRCWithdrawn is a free log retrieval operation binding the contract event 0x2958029af76592f61d851f6d7c7456b3d9ae7536e7dc715c408fe02cd06c4651.
//
// Solidity: event LRCWithdrawn(address indexed user, uint256 amount)
func (_UserStakingPool *UserStakingPoolFilterer) FilterLRCWithdrawn(opts *bind.FilterOpts, user []common.Address) (*UserStakingPoolLRCWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UserStakingPool.contract.FilterLogs(opts, "LRCWithdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &UserStakingPoolLRCWithdrawnIterator{contract: _UserStakingPool.contract, event: "LRCWithdrawn", logs: logs, sub: sub}, nil
}

// WatchLRCWithdrawn is a free log subscription operation binding the contract event 0x2958029af76592f61d851f6d7c7456b3d9ae7536e7dc715c408fe02cd06c4651.
//
// Solidity: event LRCWithdrawn(address indexed user, uint256 amount)
func (_UserStakingPool *UserStakingPoolFilterer) WatchLRCWithdrawn(opts *bind.WatchOpts, sink chan<- *UserStakingPoolLRCWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UserStakingPool.contract.WatchLogs(opts, "LRCWithdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserStakingPoolLRCWithdrawn)
				if err := _UserStakingPool.contract.UnpackLog(event, "LRCWithdrawn", log); err != nil {
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

// ParseLRCWithdrawn is a log parse operation binding the contract event 0x2958029af76592f61d851f6d7c7456b3d9ae7536e7dc715c408fe02cd06c4651.
//
// Solidity: event LRCWithdrawn(address indexed user, uint256 amount)
func (_UserStakingPool *UserStakingPoolFilterer) ParseLRCWithdrawn(log types.Log) (*UserStakingPoolLRCWithdrawn, error) {
	event := new(UserStakingPoolLRCWithdrawn)
	if err := _UserStakingPool.contract.UnpackLog(event, "LRCWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UserStakingPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UserStakingPool contract.
type UserStakingPoolOwnershipTransferredIterator struct {
	Event *UserStakingPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UserStakingPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserStakingPoolOwnershipTransferred)
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
		it.Event = new(UserStakingPoolOwnershipTransferred)
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
func (it *UserStakingPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserStakingPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserStakingPoolOwnershipTransferred represents a OwnershipTransferred event raised by the UserStakingPool contract.
type UserStakingPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UserStakingPool *UserStakingPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UserStakingPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UserStakingPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UserStakingPoolOwnershipTransferredIterator{contract: _UserStakingPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UserStakingPool *UserStakingPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UserStakingPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UserStakingPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserStakingPoolOwnershipTransferred)
				if err := _UserStakingPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UserStakingPool *UserStakingPoolFilterer) ParseOwnershipTransferred(log types.Log) (*UserStakingPoolOwnershipTransferred, error) {
	event := new(UserStakingPoolOwnershipTransferred)
	if err := _UserStakingPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UserStakingPoolProtocolFeeVaultChangedIterator is returned from FilterProtocolFeeVaultChanged and is used to iterate over the raw logs and unpacked data for ProtocolFeeVaultChanged events raised by the UserStakingPool contract.
type UserStakingPoolProtocolFeeVaultChangedIterator struct {
	Event *UserStakingPoolProtocolFeeVaultChanged // Event containing the contract specifics and raw log

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
func (it *UserStakingPoolProtocolFeeVaultChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserStakingPoolProtocolFeeVaultChanged)
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
		it.Event = new(UserStakingPoolProtocolFeeVaultChanged)
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
func (it *UserStakingPoolProtocolFeeVaultChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserStakingPoolProtocolFeeVaultChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserStakingPoolProtocolFeeVaultChanged represents a ProtocolFeeVaultChanged event raised by the UserStakingPool contract.
type UserStakingPoolProtocolFeeVaultChanged struct {
	FeeVaultAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeVaultChanged is a free log retrieval operation binding the contract event 0xb3b08642421ba31da9df792b3dbe9f52ff071894103ba3b1c4317b2daf33736c.
//
// Solidity: event ProtocolFeeVaultChanged(address feeVaultAddress)
func (_UserStakingPool *UserStakingPoolFilterer) FilterProtocolFeeVaultChanged(opts *bind.FilterOpts) (*UserStakingPoolProtocolFeeVaultChangedIterator, error) {

	logs, sub, err := _UserStakingPool.contract.FilterLogs(opts, "ProtocolFeeVaultChanged")
	if err != nil {
		return nil, err
	}
	return &UserStakingPoolProtocolFeeVaultChangedIterator{contract: _UserStakingPool.contract, event: "ProtocolFeeVaultChanged", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeVaultChanged is a free log subscription operation binding the contract event 0xb3b08642421ba31da9df792b3dbe9f52ff071894103ba3b1c4317b2daf33736c.
//
// Solidity: event ProtocolFeeVaultChanged(address feeVaultAddress)
func (_UserStakingPool *UserStakingPoolFilterer) WatchProtocolFeeVaultChanged(opts *bind.WatchOpts, sink chan<- *UserStakingPoolProtocolFeeVaultChanged) (event.Subscription, error) {

	logs, sub, err := _UserStakingPool.contract.WatchLogs(opts, "ProtocolFeeVaultChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserStakingPoolProtocolFeeVaultChanged)
				if err := _UserStakingPool.contract.UnpackLog(event, "ProtocolFeeVaultChanged", log); err != nil {
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

// ParseProtocolFeeVaultChanged is a log parse operation binding the contract event 0xb3b08642421ba31da9df792b3dbe9f52ff071894103ba3b1c4317b2daf33736c.
//
// Solidity: event ProtocolFeeVaultChanged(address feeVaultAddress)
func (_UserStakingPool *UserStakingPoolFilterer) ParseProtocolFeeVaultChanged(log types.Log) (*UserStakingPoolProtocolFeeVaultChanged, error) {
	event := new(UserStakingPoolProtocolFeeVaultChanged)
	if err := _UserStakingPool.contract.UnpackLog(event, "ProtocolFeeVaultChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}
