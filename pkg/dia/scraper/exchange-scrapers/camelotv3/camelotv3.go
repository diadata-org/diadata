// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package camelotv3

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

// Camelotv3MetaData contains all meta data concerning the Camelotv3 contract.
var Camelotv3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolDeployer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newDefaultCommunityFee\",\"type\":\"uint8\"}],\"name\":\"DefaultCommunityFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFarmingAddress\",\"type\":\"address\"}],\"name\":\"FarmingAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"alpha1\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"alpha2\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"beta1\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"beta2\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gamma1\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gamma2\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"volumeBeta\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"volumeGamma\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"baseFee\",\"type\":\"uint16\"}],\"name\":\"FeeConfiguration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"Owner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"Pool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVaultAddress\",\"type\":\"address\"}],\"name\":\"VaultAddress\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseFeeConfiguration\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"alpha1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"alpha2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"beta1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"beta2\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"gamma1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"gamma2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"volumeBeta\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"volumeGamma\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"baseFee\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultCommunityFee\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"farmingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolByPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"alpha1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"alpha2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"beta1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"beta2\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"gamma1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"gamma2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"volumeBeta\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"volumeGamma\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"baseFee\",\"type\":\"uint16\"}],\"name\":\"setBaseFeeConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newDefaultCommunityFee\",\"type\":\"uint8\"}],\"name\":\"setDefaultCommunityFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_farmingAddress\",\"type\":\"address\"}],\"name\":\"setFarmingAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultAddress\",\"type\":\"address\"}],\"name\":\"setVaultAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Camelotv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Camelotv3MetaData.ABI instead.
var Camelotv3ABI = Camelotv3MetaData.ABI

// Camelotv3 is an auto generated Go binding around an Ethereum contract.
type Camelotv3 struct {
	Camelotv3Caller     // Read-only binding to the contract
	Camelotv3Transactor // Write-only binding to the contract
	Camelotv3Filterer   // Log filterer for contract events
}

// Camelotv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Camelotv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Camelotv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Camelotv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Camelotv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Camelotv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Camelotv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Camelotv3Session struct {
	Contract     *Camelotv3        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Camelotv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Camelotv3CallerSession struct {
	Contract *Camelotv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Camelotv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Camelotv3TransactorSession struct {
	Contract     *Camelotv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Camelotv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Camelotv3Raw struct {
	Contract *Camelotv3 // Generic contract binding to access the raw methods on
}

// Camelotv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Camelotv3CallerRaw struct {
	Contract *Camelotv3Caller // Generic read-only contract binding to access the raw methods on
}

// Camelotv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Camelotv3TransactorRaw struct {
	Contract *Camelotv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCamelotv3 creates a new instance of Camelotv3, bound to a specific deployed contract.
func NewCamelotv3(address common.Address, backend bind.ContractBackend) (*Camelotv3, error) {
	contract, err := bindCamelotv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Camelotv3{Camelotv3Caller: Camelotv3Caller{contract: contract}, Camelotv3Transactor: Camelotv3Transactor{contract: contract}, Camelotv3Filterer: Camelotv3Filterer{contract: contract}}, nil
}

// NewCamelotv3Caller creates a new read-only instance of Camelotv3, bound to a specific deployed contract.
func NewCamelotv3Caller(address common.Address, caller bind.ContractCaller) (*Camelotv3Caller, error) {
	contract, err := bindCamelotv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Camelotv3Caller{contract: contract}, nil
}

// NewCamelotv3Transactor creates a new write-only instance of Camelotv3, bound to a specific deployed contract.
func NewCamelotv3Transactor(address common.Address, transactor bind.ContractTransactor) (*Camelotv3Transactor, error) {
	contract, err := bindCamelotv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Camelotv3Transactor{contract: contract}, nil
}

// NewCamelotv3Filterer creates a new log filterer instance of Camelotv3, bound to a specific deployed contract.
func NewCamelotv3Filterer(address common.Address, filterer bind.ContractFilterer) (*Camelotv3Filterer, error) {
	contract, err := bindCamelotv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Camelotv3Filterer{contract: contract}, nil
}

// bindCamelotv3 binds a generic wrapper to an already deployed contract.
func bindCamelotv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Camelotv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Camelotv3 *Camelotv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Camelotv3.Contract.Camelotv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Camelotv3 *Camelotv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Camelotv3.Contract.Camelotv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Camelotv3 *Camelotv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Camelotv3.Contract.Camelotv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Camelotv3 *Camelotv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Camelotv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Camelotv3 *Camelotv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Camelotv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Camelotv3 *Camelotv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Camelotv3.Contract.contract.Transact(opts, method, params...)
}

// BaseFeeConfiguration is a free data retrieval call binding the contract method 0x9832853a.
//
// Solidity: function baseFeeConfiguration() view returns(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_Camelotv3 *Camelotv3Caller) BaseFeeConfiguration(opts *bind.CallOpts) (struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
}, error) {
	var out []interface{}
	err := _Camelotv3.contract.Call(opts, &out, "baseFeeConfiguration")

	outstruct := new(struct {
		Alpha1      uint16
		Alpha2      uint16
		Beta1       uint32
		Beta2       uint32
		Gamma1      uint16
		Gamma2      uint16
		VolumeBeta  uint32
		VolumeGamma uint16
		BaseFee     uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Alpha1 = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.Alpha2 = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.Beta1 = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Beta2 = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Gamma1 = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.Gamma2 = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.VolumeBeta = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.VolumeGamma = *abi.ConvertType(out[7], new(uint16)).(*uint16)
	outstruct.BaseFee = *abi.ConvertType(out[8], new(uint16)).(*uint16)

	return *outstruct, err

}

// BaseFeeConfiguration is a free data retrieval call binding the contract method 0x9832853a.
//
// Solidity: function baseFeeConfiguration() view returns(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_Camelotv3 *Camelotv3Session) BaseFeeConfiguration() (struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
}, error) {
	return _Camelotv3.Contract.BaseFeeConfiguration(&_Camelotv3.CallOpts)
}

// BaseFeeConfiguration is a free data retrieval call binding the contract method 0x9832853a.
//
// Solidity: function baseFeeConfiguration() view returns(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_Camelotv3 *Camelotv3CallerSession) BaseFeeConfiguration() (struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
}, error) {
	return _Camelotv3.Contract.BaseFeeConfiguration(&_Camelotv3.CallOpts)
}

// DefaultCommunityFee is a free data retrieval call binding the contract method 0x2f8a39dd.
//
// Solidity: function defaultCommunityFee() view returns(uint8)
func (_Camelotv3 *Camelotv3Caller) DefaultCommunityFee(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Camelotv3.contract.Call(opts, &out, "defaultCommunityFee")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultCommunityFee is a free data retrieval call binding the contract method 0x2f8a39dd.
//
// Solidity: function defaultCommunityFee() view returns(uint8)
func (_Camelotv3 *Camelotv3Session) DefaultCommunityFee() (uint8, error) {
	return _Camelotv3.Contract.DefaultCommunityFee(&_Camelotv3.CallOpts)
}

// DefaultCommunityFee is a free data retrieval call binding the contract method 0x2f8a39dd.
//
// Solidity: function defaultCommunityFee() view returns(uint8)
func (_Camelotv3 *Camelotv3CallerSession) DefaultCommunityFee() (uint8, error) {
	return _Camelotv3.Contract.DefaultCommunityFee(&_Camelotv3.CallOpts)
}

// FarmingAddress is a free data retrieval call binding the contract method 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (_Camelotv3 *Camelotv3Caller) FarmingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Camelotv3.contract.Call(opts, &out, "farmingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FarmingAddress is a free data retrieval call binding the contract method 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (_Camelotv3 *Camelotv3Session) FarmingAddress() (common.Address, error) {
	return _Camelotv3.Contract.FarmingAddress(&_Camelotv3.CallOpts)
}

// FarmingAddress is a free data retrieval call binding the contract method 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (_Camelotv3 *Camelotv3CallerSession) FarmingAddress() (common.Address, error) {
	return _Camelotv3.Contract.FarmingAddress(&_Camelotv3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Camelotv3 *Camelotv3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Camelotv3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Camelotv3 *Camelotv3Session) Owner() (common.Address, error) {
	return _Camelotv3.Contract.Owner(&_Camelotv3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Camelotv3 *Camelotv3CallerSession) Owner() (common.Address, error) {
	return _Camelotv3.Contract.Owner(&_Camelotv3.CallOpts)
}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address , address ) view returns(address)
func (_Camelotv3 *Camelotv3Caller) PoolByPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Camelotv3.contract.Call(opts, &out, "poolByPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address , address ) view returns(address)
func (_Camelotv3 *Camelotv3Session) PoolByPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Camelotv3.Contract.PoolByPair(&_Camelotv3.CallOpts, arg0, arg1)
}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address , address ) view returns(address)
func (_Camelotv3 *Camelotv3CallerSession) PoolByPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Camelotv3.Contract.PoolByPair(&_Camelotv3.CallOpts, arg0, arg1)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_Camelotv3 *Camelotv3Caller) PoolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Camelotv3.contract.Call(opts, &out, "poolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_Camelotv3 *Camelotv3Session) PoolDeployer() (common.Address, error) {
	return _Camelotv3.Contract.PoolDeployer(&_Camelotv3.CallOpts)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_Camelotv3 *Camelotv3CallerSession) PoolDeployer() (common.Address, error) {
	return _Camelotv3.Contract.PoolDeployer(&_Camelotv3.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_Camelotv3 *Camelotv3Caller) VaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Camelotv3.contract.Call(opts, &out, "vaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_Camelotv3 *Camelotv3Session) VaultAddress() (common.Address, error) {
	return _Camelotv3.Contract.VaultAddress(&_Camelotv3.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_Camelotv3 *Camelotv3CallerSession) VaultAddress() (common.Address, error) {
	return _Camelotv3.Contract.VaultAddress(&_Camelotv3.CallOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (_Camelotv3 *Camelotv3Transactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Camelotv3.contract.Transact(opts, "createPool", tokenA, tokenB)
}

// CreatePool is a paid mutator transaction binding the contract method 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (_Camelotv3 *Camelotv3Session) CreatePool(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Camelotv3.Contract.CreatePool(&_Camelotv3.TransactOpts, tokenA, tokenB)
}

// CreatePool is a paid mutator transaction binding the contract method 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (_Camelotv3 *Camelotv3TransactorSession) CreatePool(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Camelotv3.Contract.CreatePool(&_Camelotv3.TransactOpts, tokenA, tokenB)
}

// SetBaseFeeConfiguration is a paid mutator transaction binding the contract method 0x5d6d7e93.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (_Camelotv3 *Camelotv3Transactor) SetBaseFeeConfiguration(opts *bind.TransactOpts, alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) (*types.Transaction, error) {
	return _Camelotv3.contract.Transact(opts, "setBaseFeeConfiguration", alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// SetBaseFeeConfiguration is a paid mutator transaction binding the contract method 0x5d6d7e93.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (_Camelotv3 *Camelotv3Session) SetBaseFeeConfiguration(alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) (*types.Transaction, error) {
	return _Camelotv3.Contract.SetBaseFeeConfiguration(&_Camelotv3.TransactOpts, alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// SetBaseFeeConfiguration is a paid mutator transaction binding the contract method 0x5d6d7e93.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (_Camelotv3 *Camelotv3TransactorSession) SetBaseFeeConfiguration(alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) (*types.Transaction, error) {
	return _Camelotv3.Contract.SetBaseFeeConfiguration(&_Camelotv3.TransactOpts, alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// SetDefaultCommunityFee is a paid mutator transaction binding the contract method 0x371e3521.
//
// Solidity: function setDefaultCommunityFee(uint8 newDefaultCommunityFee) returns()
func (_Camelotv3 *Camelotv3Transactor) SetDefaultCommunityFee(opts *bind.TransactOpts, newDefaultCommunityFee uint8) (*types.Transaction, error) {
	return _Camelotv3.contract.Transact(opts, "setDefaultCommunityFee", newDefaultCommunityFee)
}

// SetDefaultCommunityFee is a paid mutator transaction binding the contract method 0x371e3521.
//
// Solidity: function setDefaultCommunityFee(uint8 newDefaultCommunityFee) returns()
func (_Camelotv3 *Camelotv3Session) SetDefaultCommunityFee(newDefaultCommunityFee uint8) (*types.Transaction, error) {
	return _Camelotv3.Contract.SetDefaultCommunityFee(&_Camelotv3.TransactOpts, newDefaultCommunityFee)
}

// SetDefaultCommunityFee is a paid mutator transaction binding the contract method 0x371e3521.
//
// Solidity: function setDefaultCommunityFee(uint8 newDefaultCommunityFee) returns()
func (_Camelotv3 *Camelotv3TransactorSession) SetDefaultCommunityFee(newDefaultCommunityFee uint8) (*types.Transaction, error) {
	return _Camelotv3.Contract.SetDefaultCommunityFee(&_Camelotv3.TransactOpts, newDefaultCommunityFee)
}

// SetFarmingAddress is a paid mutator transaction binding the contract method 0xb001f618.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (_Camelotv3 *Camelotv3Transactor) SetFarmingAddress(opts *bind.TransactOpts, _farmingAddress common.Address) (*types.Transaction, error) {
	return _Camelotv3.contract.Transact(opts, "setFarmingAddress", _farmingAddress)
}

// SetFarmingAddress is a paid mutator transaction binding the contract method 0xb001f618.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (_Camelotv3 *Camelotv3Session) SetFarmingAddress(_farmingAddress common.Address) (*types.Transaction, error) {
	return _Camelotv3.Contract.SetFarmingAddress(&_Camelotv3.TransactOpts, _farmingAddress)
}

// SetFarmingAddress is a paid mutator transaction binding the contract method 0xb001f618.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (_Camelotv3 *Camelotv3TransactorSession) SetFarmingAddress(_farmingAddress common.Address) (*types.Transaction, error) {
	return _Camelotv3.Contract.SetFarmingAddress(&_Camelotv3.TransactOpts, _farmingAddress)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Camelotv3 *Camelotv3Transactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Camelotv3.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Camelotv3 *Camelotv3Session) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Camelotv3.Contract.SetOwner(&_Camelotv3.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Camelotv3 *Camelotv3TransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Camelotv3.Contract.SetOwner(&_Camelotv3.TransactOpts, _owner)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_Camelotv3 *Camelotv3Transactor) SetVaultAddress(opts *bind.TransactOpts, _vaultAddress common.Address) (*types.Transaction, error) {
	return _Camelotv3.contract.Transact(opts, "setVaultAddress", _vaultAddress)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_Camelotv3 *Camelotv3Session) SetVaultAddress(_vaultAddress common.Address) (*types.Transaction, error) {
	return _Camelotv3.Contract.SetVaultAddress(&_Camelotv3.TransactOpts, _vaultAddress)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_Camelotv3 *Camelotv3TransactorSession) SetVaultAddress(_vaultAddress common.Address) (*types.Transaction, error) {
	return _Camelotv3.Contract.SetVaultAddress(&_Camelotv3.TransactOpts, _vaultAddress)
}

// Camelotv3DefaultCommunityFeeIterator is returned from FilterDefaultCommunityFee and is used to iterate over the raw logs and unpacked data for DefaultCommunityFee events raised by the Camelotv3 contract.
type Camelotv3DefaultCommunityFeeIterator struct {
	Event *Camelotv3DefaultCommunityFee // Event containing the contract specifics and raw log

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
func (it *Camelotv3DefaultCommunityFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3DefaultCommunityFee)
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
		it.Event = new(Camelotv3DefaultCommunityFee)
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
func (it *Camelotv3DefaultCommunityFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3DefaultCommunityFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3DefaultCommunityFee represents a DefaultCommunityFee event raised by the Camelotv3 contract.
type Camelotv3DefaultCommunityFee struct {
	NewDefaultCommunityFee uint8
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterDefaultCommunityFee is a free log retrieval operation binding the contract event 0x88cb5103fd9d88d417e72dc496030c71c65d1500548a9e9530e7d812b6a35558.
//
// Solidity: event DefaultCommunityFee(uint8 newDefaultCommunityFee)
func (_Camelotv3 *Camelotv3Filterer) FilterDefaultCommunityFee(opts *bind.FilterOpts) (*Camelotv3DefaultCommunityFeeIterator, error) {

	logs, sub, err := _Camelotv3.contract.FilterLogs(opts, "DefaultCommunityFee")
	if err != nil {
		return nil, err
	}
	return &Camelotv3DefaultCommunityFeeIterator{contract: _Camelotv3.contract, event: "DefaultCommunityFee", logs: logs, sub: sub}, nil
}

// WatchDefaultCommunityFee is a free log subscription operation binding the contract event 0x88cb5103fd9d88d417e72dc496030c71c65d1500548a9e9530e7d812b6a35558.
//
// Solidity: event DefaultCommunityFee(uint8 newDefaultCommunityFee)
func (_Camelotv3 *Camelotv3Filterer) WatchDefaultCommunityFee(opts *bind.WatchOpts, sink chan<- *Camelotv3DefaultCommunityFee) (event.Subscription, error) {

	logs, sub, err := _Camelotv3.contract.WatchLogs(opts, "DefaultCommunityFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3DefaultCommunityFee)
				if err := _Camelotv3.contract.UnpackLog(event, "DefaultCommunityFee", log); err != nil {
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

// ParseDefaultCommunityFee is a log parse operation binding the contract event 0x88cb5103fd9d88d417e72dc496030c71c65d1500548a9e9530e7d812b6a35558.
//
// Solidity: event DefaultCommunityFee(uint8 newDefaultCommunityFee)
func (_Camelotv3 *Camelotv3Filterer) ParseDefaultCommunityFee(log types.Log) (*Camelotv3DefaultCommunityFee, error) {
	event := new(Camelotv3DefaultCommunityFee)
	if err := _Camelotv3.contract.UnpackLog(event, "DefaultCommunityFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Camelotv3FarmingAddressIterator is returned from FilterFarmingAddress and is used to iterate over the raw logs and unpacked data for FarmingAddress events raised by the Camelotv3 contract.
type Camelotv3FarmingAddressIterator struct {
	Event *Camelotv3FarmingAddress // Event containing the contract specifics and raw log

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
func (it *Camelotv3FarmingAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3FarmingAddress)
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
		it.Event = new(Camelotv3FarmingAddress)
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
func (it *Camelotv3FarmingAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3FarmingAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3FarmingAddress represents a FarmingAddress event raised by the Camelotv3 contract.
type Camelotv3FarmingAddress struct {
	NewFarmingAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFarmingAddress is a free log retrieval operation binding the contract event 0x56b9e8342f530796ceed0d5529abdcdeae6e4f2ac1dc456ceb73bbda898e0cd3.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (_Camelotv3 *Camelotv3Filterer) FilterFarmingAddress(opts *bind.FilterOpts, newFarmingAddress []common.Address) (*Camelotv3FarmingAddressIterator, error) {

	var newFarmingAddressRule []interface{}
	for _, newFarmingAddressItem := range newFarmingAddress {
		newFarmingAddressRule = append(newFarmingAddressRule, newFarmingAddressItem)
	}

	logs, sub, err := _Camelotv3.contract.FilterLogs(opts, "FarmingAddress", newFarmingAddressRule)
	if err != nil {
		return nil, err
	}
	return &Camelotv3FarmingAddressIterator{contract: _Camelotv3.contract, event: "FarmingAddress", logs: logs, sub: sub}, nil
}

// WatchFarmingAddress is a free log subscription operation binding the contract event 0x56b9e8342f530796ceed0d5529abdcdeae6e4f2ac1dc456ceb73bbda898e0cd3.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (_Camelotv3 *Camelotv3Filterer) WatchFarmingAddress(opts *bind.WatchOpts, sink chan<- *Camelotv3FarmingAddress, newFarmingAddress []common.Address) (event.Subscription, error) {

	var newFarmingAddressRule []interface{}
	for _, newFarmingAddressItem := range newFarmingAddress {
		newFarmingAddressRule = append(newFarmingAddressRule, newFarmingAddressItem)
	}

	logs, sub, err := _Camelotv3.contract.WatchLogs(opts, "FarmingAddress", newFarmingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3FarmingAddress)
				if err := _Camelotv3.contract.UnpackLog(event, "FarmingAddress", log); err != nil {
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

// ParseFarmingAddress is a log parse operation binding the contract event 0x56b9e8342f530796ceed0d5529abdcdeae6e4f2ac1dc456ceb73bbda898e0cd3.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (_Camelotv3 *Camelotv3Filterer) ParseFarmingAddress(log types.Log) (*Camelotv3FarmingAddress, error) {
	event := new(Camelotv3FarmingAddress)
	if err := _Camelotv3.contract.UnpackLog(event, "FarmingAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Camelotv3FeeConfigurationIterator is returned from FilterFeeConfiguration and is used to iterate over the raw logs and unpacked data for FeeConfiguration events raised by the Camelotv3 contract.
type Camelotv3FeeConfigurationIterator struct {
	Event *Camelotv3FeeConfiguration // Event containing the contract specifics and raw log

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
func (it *Camelotv3FeeConfigurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3FeeConfiguration)
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
		it.Event = new(Camelotv3FeeConfiguration)
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
func (it *Camelotv3FeeConfigurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3FeeConfigurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3FeeConfiguration represents a FeeConfiguration event raised by the Camelotv3 contract.
type Camelotv3FeeConfiguration struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeConfiguration is a free log retrieval operation binding the contract event 0x4035ab409f15e202f9f114632e1fb14a0552325955722be18503403e7f98730c.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_Camelotv3 *Camelotv3Filterer) FilterFeeConfiguration(opts *bind.FilterOpts) (*Camelotv3FeeConfigurationIterator, error) {

	logs, sub, err := _Camelotv3.contract.FilterLogs(opts, "FeeConfiguration")
	if err != nil {
		return nil, err
	}
	return &Camelotv3FeeConfigurationIterator{contract: _Camelotv3.contract, event: "FeeConfiguration", logs: logs, sub: sub}, nil
}

// WatchFeeConfiguration is a free log subscription operation binding the contract event 0x4035ab409f15e202f9f114632e1fb14a0552325955722be18503403e7f98730c.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_Camelotv3 *Camelotv3Filterer) WatchFeeConfiguration(opts *bind.WatchOpts, sink chan<- *Camelotv3FeeConfiguration) (event.Subscription, error) {

	logs, sub, err := _Camelotv3.contract.WatchLogs(opts, "FeeConfiguration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3FeeConfiguration)
				if err := _Camelotv3.contract.UnpackLog(event, "FeeConfiguration", log); err != nil {
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

// ParseFeeConfiguration is a log parse operation binding the contract event 0x4035ab409f15e202f9f114632e1fb14a0552325955722be18503403e7f98730c.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_Camelotv3 *Camelotv3Filterer) ParseFeeConfiguration(log types.Log) (*Camelotv3FeeConfiguration, error) {
	event := new(Camelotv3FeeConfiguration)
	if err := _Camelotv3.contract.UnpackLog(event, "FeeConfiguration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Camelotv3OwnerIterator is returned from FilterOwner and is used to iterate over the raw logs and unpacked data for Owner events raised by the Camelotv3 contract.
type Camelotv3OwnerIterator struct {
	Event *Camelotv3Owner // Event containing the contract specifics and raw log

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
func (it *Camelotv3OwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3Owner)
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
		it.Event = new(Camelotv3Owner)
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
func (it *Camelotv3OwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3OwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3Owner represents a Owner event raised by the Camelotv3 contract.
type Camelotv3Owner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwner is a free log retrieval operation binding the contract event 0xa5e220c2c27d986cc8efeafa8f34ba6ea6bf96a34e146b29b6bdd8587771b130.
//
// Solidity: event Owner(address indexed newOwner)
func (_Camelotv3 *Camelotv3Filterer) FilterOwner(opts *bind.FilterOpts, newOwner []common.Address) (*Camelotv3OwnerIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Camelotv3.contract.FilterLogs(opts, "Owner", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Camelotv3OwnerIterator{contract: _Camelotv3.contract, event: "Owner", logs: logs, sub: sub}, nil
}

// WatchOwner is a free log subscription operation binding the contract event 0xa5e220c2c27d986cc8efeafa8f34ba6ea6bf96a34e146b29b6bdd8587771b130.
//
// Solidity: event Owner(address indexed newOwner)
func (_Camelotv3 *Camelotv3Filterer) WatchOwner(opts *bind.WatchOpts, sink chan<- *Camelotv3Owner, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Camelotv3.contract.WatchLogs(opts, "Owner", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3Owner)
				if err := _Camelotv3.contract.UnpackLog(event, "Owner", log); err != nil {
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

// ParseOwner is a log parse operation binding the contract event 0xa5e220c2c27d986cc8efeafa8f34ba6ea6bf96a34e146b29b6bdd8587771b130.
//
// Solidity: event Owner(address indexed newOwner)
func (_Camelotv3 *Camelotv3Filterer) ParseOwner(log types.Log) (*Camelotv3Owner, error) {
	event := new(Camelotv3Owner)
	if err := _Camelotv3.contract.UnpackLog(event, "Owner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Camelotv3PoolIterator is returned from FilterPool and is used to iterate over the raw logs and unpacked data for Pool events raised by the Camelotv3 contract.
type Camelotv3PoolIterator struct {
	Event *Camelotv3Pool // Event containing the contract specifics and raw log

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
func (it *Camelotv3PoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3Pool)
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
		it.Event = new(Camelotv3Pool)
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
func (it *Camelotv3PoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3PoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3Pool represents a Pool event raised by the Camelotv3 contract.
type Camelotv3Pool struct {
	Token0 common.Address
	Token1 common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPool is a free log retrieval operation binding the contract event 0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (_Camelotv3 *Camelotv3Filterer) FilterPool(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*Camelotv3PoolIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Camelotv3.contract.FilterLogs(opts, "Pool", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &Camelotv3PoolIterator{contract: _Camelotv3.contract, event: "Pool", logs: logs, sub: sub}, nil
}

// WatchPool is a free log subscription operation binding the contract event 0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (_Camelotv3 *Camelotv3Filterer) WatchPool(opts *bind.WatchOpts, sink chan<- *Camelotv3Pool, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Camelotv3.contract.WatchLogs(opts, "Pool", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3Pool)
				if err := _Camelotv3.contract.UnpackLog(event, "Pool", log); err != nil {
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

// ParsePool is a log parse operation binding the contract event 0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (_Camelotv3 *Camelotv3Filterer) ParsePool(log types.Log) (*Camelotv3Pool, error) {
	event := new(Camelotv3Pool)
	if err := _Camelotv3.contract.UnpackLog(event, "Pool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Camelotv3VaultAddressIterator is returned from FilterVaultAddress and is used to iterate over the raw logs and unpacked data for VaultAddress events raised by the Camelotv3 contract.
type Camelotv3VaultAddressIterator struct {
	Event *Camelotv3VaultAddress // Event containing the contract specifics and raw log

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
func (it *Camelotv3VaultAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Camelotv3VaultAddress)
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
		it.Event = new(Camelotv3VaultAddress)
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
func (it *Camelotv3VaultAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Camelotv3VaultAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Camelotv3VaultAddress represents a VaultAddress event raised by the Camelotv3 contract.
type Camelotv3VaultAddress struct {
	NewVaultAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVaultAddress is a free log retrieval operation binding the contract event 0xb9c265ae4414f501736ec5d4961edc3309e4385eb2ff3feeecb30fb36621dd83.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (_Camelotv3 *Camelotv3Filterer) FilterVaultAddress(opts *bind.FilterOpts, newVaultAddress []common.Address) (*Camelotv3VaultAddressIterator, error) {

	var newVaultAddressRule []interface{}
	for _, newVaultAddressItem := range newVaultAddress {
		newVaultAddressRule = append(newVaultAddressRule, newVaultAddressItem)
	}

	logs, sub, err := _Camelotv3.contract.FilterLogs(opts, "VaultAddress", newVaultAddressRule)
	if err != nil {
		return nil, err
	}
	return &Camelotv3VaultAddressIterator{contract: _Camelotv3.contract, event: "VaultAddress", logs: logs, sub: sub}, nil
}

// WatchVaultAddress is a free log subscription operation binding the contract event 0xb9c265ae4414f501736ec5d4961edc3309e4385eb2ff3feeecb30fb36621dd83.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (_Camelotv3 *Camelotv3Filterer) WatchVaultAddress(opts *bind.WatchOpts, sink chan<- *Camelotv3VaultAddress, newVaultAddress []common.Address) (event.Subscription, error) {

	var newVaultAddressRule []interface{}
	for _, newVaultAddressItem := range newVaultAddress {
		newVaultAddressRule = append(newVaultAddressRule, newVaultAddressItem)
	}

	logs, sub, err := _Camelotv3.contract.WatchLogs(opts, "VaultAddress", newVaultAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Camelotv3VaultAddress)
				if err := _Camelotv3.contract.UnpackLog(event, "VaultAddress", log); err != nil {
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

// ParseVaultAddress is a log parse operation binding the contract event 0xb9c265ae4414f501736ec5d4961edc3309e4385eb2ff3feeecb30fb36621dd83.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (_Camelotv3 *Camelotv3Filterer) ParseVaultAddress(log types.Log) (*Camelotv3VaultAddress, error) {
	event := new(Camelotv3VaultAddress)
	if err := _Camelotv3.contract.UnpackLog(event, "VaultAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
