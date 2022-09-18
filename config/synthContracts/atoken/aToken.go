// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package atoken

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

// AtokenMetaData contains all meta data concerning the Atoken contract.
var AtokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BalanceTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"incentivesController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"aTokenDecimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aTokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aTokenSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ATOKEN_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_REVISION\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_TREASURY_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDERLYING_ASSET_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiverOfUnderlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIncentivesController\",\"outputs\":[{\"internalType\":\"contractIAaveIncentivesController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getPreviousIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getScaledUserBalanceAndSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"handleRepayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"initializingPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"contractIAaveIncentivesController\",\"name\":\"incentivesController\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"aTokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"aTokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"aTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"mintToTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"scaledBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scaledTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAaveIncentivesController\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"setIncentivesController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferOnLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferUnderlyingTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AtokenABI is the input ABI used to generate the binding from.
// Deprecated: Use AtokenMetaData.ABI instead.
var AtokenABI = AtokenMetaData.ABI

// Atoken is an auto generated Go binding around an Ethereum contract.
type Atoken struct {
	AtokenCaller     // Read-only binding to the contract
	AtokenTransactor // Write-only binding to the contract
	AtokenFilterer   // Log filterer for contract events
}

// AtokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type AtokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AtokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AtokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AtokenSession struct {
	Contract     *Atoken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AtokenCallerSession struct {
	Contract *AtokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AtokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AtokenTransactorSession struct {
	Contract     *AtokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type AtokenRaw struct {
	Contract *Atoken // Generic contract binding to access the raw methods on
}

// AtokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AtokenCallerRaw struct {
	Contract *AtokenCaller // Generic read-only contract binding to access the raw methods on
}

// AtokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AtokenTransactorRaw struct {
	Contract *AtokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAtoken creates a new instance of Atoken, bound to a specific deployed contract.
func NewAtoken(address common.Address, backend bind.ContractBackend) (*Atoken, error) {
	contract, err := bindAtoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Atoken{AtokenCaller: AtokenCaller{contract: contract}, AtokenTransactor: AtokenTransactor{contract: contract}, AtokenFilterer: AtokenFilterer{contract: contract}}, nil
}

// NewAtokenCaller creates a new read-only instance of Atoken, bound to a specific deployed contract.
func NewAtokenCaller(address common.Address, caller bind.ContractCaller) (*AtokenCaller, error) {
	contract, err := bindAtoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AtokenCaller{contract: contract}, nil
}

// NewAtokenTransactor creates a new write-only instance of Atoken, bound to a specific deployed contract.
func NewAtokenTransactor(address common.Address, transactor bind.ContractTransactor) (*AtokenTransactor, error) {
	contract, err := bindAtoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AtokenTransactor{contract: contract}, nil
}

// NewAtokenFilterer creates a new log filterer instance of Atoken, bound to a specific deployed contract.
func NewAtokenFilterer(address common.Address, filterer bind.ContractFilterer) (*AtokenFilterer, error) {
	contract, err := bindAtoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AtokenFilterer{contract: contract}, nil
}

// bindAtoken binds a generic wrapper to an already deployed contract.
func bindAtoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AtokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Atoken *AtokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Atoken.Contract.AtokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Atoken *AtokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atoken.Contract.AtokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Atoken *AtokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Atoken.Contract.AtokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Atoken *AtokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Atoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Atoken *AtokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Atoken *AtokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Atoken.Contract.contract.Transact(opts, method, params...)
}

// ATOKENREVISION is a free data retrieval call binding the contract method 0x0bd7ad3b.
//
// Solidity: function ATOKEN_REVISION() view returns(uint256)
func (_Atoken *AtokenCaller) ATOKENREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "ATOKEN_REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ATOKENREVISION is a free data retrieval call binding the contract method 0x0bd7ad3b.
//
// Solidity: function ATOKEN_REVISION() view returns(uint256)
func (_Atoken *AtokenSession) ATOKENREVISION() (*big.Int, error) {
	return _Atoken.Contract.ATOKENREVISION(&_Atoken.CallOpts)
}

// ATOKENREVISION is a free data retrieval call binding the contract method 0x0bd7ad3b.
//
// Solidity: function ATOKEN_REVISION() view returns(uint256)
func (_Atoken *AtokenCallerSession) ATOKENREVISION() (*big.Int, error) {
	return _Atoken.Contract.ATOKENREVISION(&_Atoken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Atoken *AtokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Atoken *AtokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Atoken.Contract.DOMAINSEPARATOR(&_Atoken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Atoken *AtokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Atoken.Contract.DOMAINSEPARATOR(&_Atoken.CallOpts)
}

// EIP712REVISION is a free data retrieval call binding the contract method 0x78160376.
//
// Solidity: function EIP712_REVISION() view returns(bytes)
func (_Atoken *AtokenCaller) EIP712REVISION(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "EIP712_REVISION")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EIP712REVISION is a free data retrieval call binding the contract method 0x78160376.
//
// Solidity: function EIP712_REVISION() view returns(bytes)
func (_Atoken *AtokenSession) EIP712REVISION() ([]byte, error) {
	return _Atoken.Contract.EIP712REVISION(&_Atoken.CallOpts)
}

// EIP712REVISION is a free data retrieval call binding the contract method 0x78160376.
//
// Solidity: function EIP712_REVISION() view returns(bytes)
func (_Atoken *AtokenCallerSession) EIP712REVISION() ([]byte, error) {
	return _Atoken.Contract.EIP712REVISION(&_Atoken.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Atoken *AtokenCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Atoken *AtokenSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Atoken.Contract.PERMITTYPEHASH(&_Atoken.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Atoken *AtokenCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Atoken.Contract.PERMITTYPEHASH(&_Atoken.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Atoken *AtokenCaller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Atoken *AtokenSession) POOL() (common.Address, error) {
	return _Atoken.Contract.POOL(&_Atoken.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Atoken *AtokenCallerSession) POOL() (common.Address, error) {
	return _Atoken.Contract.POOL(&_Atoken.CallOpts)
}

// RESERVETREASURYADDRESS is a free data retrieval call binding the contract method 0xae167335.
//
// Solidity: function RESERVE_TREASURY_ADDRESS() view returns(address)
func (_Atoken *AtokenCaller) RESERVETREASURYADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "RESERVE_TREASURY_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RESERVETREASURYADDRESS is a free data retrieval call binding the contract method 0xae167335.
//
// Solidity: function RESERVE_TREASURY_ADDRESS() view returns(address)
func (_Atoken *AtokenSession) RESERVETREASURYADDRESS() (common.Address, error) {
	return _Atoken.Contract.RESERVETREASURYADDRESS(&_Atoken.CallOpts)
}

// RESERVETREASURYADDRESS is a free data retrieval call binding the contract method 0xae167335.
//
// Solidity: function RESERVE_TREASURY_ADDRESS() view returns(address)
func (_Atoken *AtokenCallerSession) RESERVETREASURYADDRESS() (common.Address, error) {
	return _Atoken.Contract.RESERVETREASURYADDRESS(&_Atoken.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_Atoken *AtokenCaller) UNDERLYINGASSETADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "UNDERLYING_ASSET_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_Atoken *AtokenSession) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _Atoken.Contract.UNDERLYINGASSETADDRESS(&_Atoken.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_Atoken *AtokenCallerSession) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _Atoken.Contract.UNDERLYINGASSETADDRESS(&_Atoken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Atoken *AtokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Atoken *AtokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Atoken.Contract.Allowance(&_Atoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Atoken *AtokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Atoken.Contract.Allowance(&_Atoken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_Atoken *AtokenCaller) BalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "balanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_Atoken *AtokenSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _Atoken.Contract.BalanceOf(&_Atoken.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_Atoken *AtokenCallerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _Atoken.Contract.BalanceOf(&_Atoken.CallOpts, user)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Atoken *AtokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Atoken *AtokenSession) Decimals() (uint8, error) {
	return _Atoken.Contract.Decimals(&_Atoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Atoken *AtokenCallerSession) Decimals() (uint8, error) {
	return _Atoken.Contract.Decimals(&_Atoken.CallOpts)
}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_Atoken *AtokenCaller) GetIncentivesController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "getIncentivesController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_Atoken *AtokenSession) GetIncentivesController() (common.Address, error) {
	return _Atoken.Contract.GetIncentivesController(&_Atoken.CallOpts)
}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_Atoken *AtokenCallerSession) GetIncentivesController() (common.Address, error) {
	return _Atoken.Contract.GetIncentivesController(&_Atoken.CallOpts)
}

// GetPreviousIndex is a free data retrieval call binding the contract method 0xe0753986.
//
// Solidity: function getPreviousIndex(address user) view returns(uint256)
func (_Atoken *AtokenCaller) GetPreviousIndex(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "getPreviousIndex", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreviousIndex is a free data retrieval call binding the contract method 0xe0753986.
//
// Solidity: function getPreviousIndex(address user) view returns(uint256)
func (_Atoken *AtokenSession) GetPreviousIndex(user common.Address) (*big.Int, error) {
	return _Atoken.Contract.GetPreviousIndex(&_Atoken.CallOpts, user)
}

// GetPreviousIndex is a free data retrieval call binding the contract method 0xe0753986.
//
// Solidity: function getPreviousIndex(address user) view returns(uint256)
func (_Atoken *AtokenCallerSession) GetPreviousIndex(user common.Address) (*big.Int, error) {
	return _Atoken.Contract.GetPreviousIndex(&_Atoken.CallOpts, user)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_Atoken *AtokenCaller) GetScaledUserBalanceAndSupply(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "getScaledUserBalanceAndSupply", user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_Atoken *AtokenSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _Atoken.Contract.GetScaledUserBalanceAndSupply(&_Atoken.CallOpts, user)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_Atoken *AtokenCallerSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _Atoken.Contract.GetScaledUserBalanceAndSupply(&_Atoken.CallOpts, user)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Atoken *AtokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Atoken *AtokenSession) Name() (string, error) {
	return _Atoken.Contract.Name(&_Atoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Atoken *AtokenCallerSession) Name() (string, error) {
	return _Atoken.Contract.Name(&_Atoken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Atoken *AtokenCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Atoken *AtokenSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Atoken.Contract.Nonces(&_Atoken.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Atoken *AtokenCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Atoken.Contract.Nonces(&_Atoken.CallOpts, owner)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_Atoken *AtokenCaller) ScaledBalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "scaledBalanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_Atoken *AtokenSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _Atoken.Contract.ScaledBalanceOf(&_Atoken.CallOpts, user)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_Atoken *AtokenCallerSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _Atoken.Contract.ScaledBalanceOf(&_Atoken.CallOpts, user)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_Atoken *AtokenCaller) ScaledTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "scaledTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_Atoken *AtokenSession) ScaledTotalSupply() (*big.Int, error) {
	return _Atoken.Contract.ScaledTotalSupply(&_Atoken.CallOpts)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_Atoken *AtokenCallerSession) ScaledTotalSupply() (*big.Int, error) {
	return _Atoken.Contract.ScaledTotalSupply(&_Atoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Atoken *AtokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Atoken *AtokenSession) Symbol() (string, error) {
	return _Atoken.Contract.Symbol(&_Atoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Atoken *AtokenCallerSession) Symbol() (string, error) {
	return _Atoken.Contract.Symbol(&_Atoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Atoken *AtokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Atoken *AtokenSession) TotalSupply() (*big.Int, error) {
	return _Atoken.Contract.TotalSupply(&_Atoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Atoken *AtokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Atoken.Contract.TotalSupply(&_Atoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Atoken *AtokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Atoken *AtokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.Approve(&_Atoken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Atoken *AtokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.Approve(&_Atoken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xd7020d0a.
//
// Solidity: function burn(address from, address receiverOfUnderlying, uint256 amount, uint256 index) returns()
func (_Atoken *AtokenTransactor) Burn(opts *bind.TransactOpts, from common.Address, receiverOfUnderlying common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "burn", from, receiverOfUnderlying, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xd7020d0a.
//
// Solidity: function burn(address from, address receiverOfUnderlying, uint256 amount, uint256 index) returns()
func (_Atoken *AtokenSession) Burn(from common.Address, receiverOfUnderlying common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.Burn(&_Atoken.TransactOpts, from, receiverOfUnderlying, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xd7020d0a.
//
// Solidity: function burn(address from, address receiverOfUnderlying, uint256 amount, uint256 index) returns()
func (_Atoken *AtokenTransactorSession) Burn(from common.Address, receiverOfUnderlying common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.Burn(&_Atoken.TransactOpts, from, receiverOfUnderlying, amount, index)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Atoken *AtokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Atoken *AtokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.DecreaseAllowance(&_Atoken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Atoken *AtokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.DecreaseAllowance(&_Atoken.TransactOpts, spender, subtractedValue)
}

// HandleRepayment is a paid mutator transaction binding the contract method 0x88dd91a1.
//
// Solidity: function handleRepayment(address user, uint256 amount) returns()
func (_Atoken *AtokenTransactor) HandleRepayment(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "handleRepayment", user, amount)
}

// HandleRepayment is a paid mutator transaction binding the contract method 0x88dd91a1.
//
// Solidity: function handleRepayment(address user, uint256 amount) returns()
func (_Atoken *AtokenSession) HandleRepayment(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.HandleRepayment(&_Atoken.TransactOpts, user, amount)
}

// HandleRepayment is a paid mutator transaction binding the contract method 0x88dd91a1.
//
// Solidity: function handleRepayment(address user, uint256 amount) returns()
func (_Atoken *AtokenTransactorSession) HandleRepayment(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.HandleRepayment(&_Atoken.TransactOpts, user, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Atoken *AtokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Atoken *AtokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.IncreaseAllowance(&_Atoken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Atoken *AtokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.IncreaseAllowance(&_Atoken.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x183fb413.
//
// Solidity: function initialize(address initializingPool, address treasury, address underlyingAsset, address incentivesController, uint8 aTokenDecimals, string aTokenName, string aTokenSymbol, bytes params) returns()
func (_Atoken *AtokenTransactor) Initialize(opts *bind.TransactOpts, initializingPool common.Address, treasury common.Address, underlyingAsset common.Address, incentivesController common.Address, aTokenDecimals uint8, aTokenName string, aTokenSymbol string, params []byte) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "initialize", initializingPool, treasury, underlyingAsset, incentivesController, aTokenDecimals, aTokenName, aTokenSymbol, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x183fb413.
//
// Solidity: function initialize(address initializingPool, address treasury, address underlyingAsset, address incentivesController, uint8 aTokenDecimals, string aTokenName, string aTokenSymbol, bytes params) returns()
func (_Atoken *AtokenSession) Initialize(initializingPool common.Address, treasury common.Address, underlyingAsset common.Address, incentivesController common.Address, aTokenDecimals uint8, aTokenName string, aTokenSymbol string, params []byte) (*types.Transaction, error) {
	return _Atoken.Contract.Initialize(&_Atoken.TransactOpts, initializingPool, treasury, underlyingAsset, incentivesController, aTokenDecimals, aTokenName, aTokenSymbol, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x183fb413.
//
// Solidity: function initialize(address initializingPool, address treasury, address underlyingAsset, address incentivesController, uint8 aTokenDecimals, string aTokenName, string aTokenSymbol, bytes params) returns()
func (_Atoken *AtokenTransactorSession) Initialize(initializingPool common.Address, treasury common.Address, underlyingAsset common.Address, incentivesController common.Address, aTokenDecimals uint8, aTokenName string, aTokenSymbol string, params []byte) (*types.Transaction, error) {
	return _Atoken.Contract.Initialize(&_Atoken.TransactOpts, initializingPool, treasury, underlyingAsset, incentivesController, aTokenDecimals, aTokenName, aTokenSymbol, params)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address caller, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_Atoken *AtokenTransactor) Mint(opts *bind.TransactOpts, caller common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "mint", caller, onBehalfOf, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address caller, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_Atoken *AtokenSession) Mint(caller common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.Mint(&_Atoken.TransactOpts, caller, onBehalfOf, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address caller, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_Atoken *AtokenTransactorSession) Mint(caller common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.Mint(&_Atoken.TransactOpts, caller, onBehalfOf, amount, index)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x7df5bd3b.
//
// Solidity: function mintToTreasury(uint256 amount, uint256 index) returns()
func (_Atoken *AtokenTransactor) MintToTreasury(opts *bind.TransactOpts, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "mintToTreasury", amount, index)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x7df5bd3b.
//
// Solidity: function mintToTreasury(uint256 amount, uint256 index) returns()
func (_Atoken *AtokenSession) MintToTreasury(amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.MintToTreasury(&_Atoken.TransactOpts, amount, index)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x7df5bd3b.
//
// Solidity: function mintToTreasury(uint256 amount, uint256 index) returns()
func (_Atoken *AtokenTransactorSession) MintToTreasury(amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.MintToTreasury(&_Atoken.TransactOpts, amount, index)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Atoken *AtokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Atoken *AtokenSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Atoken.Contract.Permit(&_Atoken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Atoken *AtokenTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Atoken.Contract.Permit(&_Atoken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns()
func (_Atoken *AtokenTransactor) RescueTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "rescueTokens", token, to, amount)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns()
func (_Atoken *AtokenSession) RescueTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.RescueTokens(&_Atoken.TransactOpts, token, to, amount)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns()
func (_Atoken *AtokenTransactorSession) RescueTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.RescueTokens(&_Atoken.TransactOpts, token, to, amount)
}

// SetIncentivesController is a paid mutator transaction binding the contract method 0xe655dbd8.
//
// Solidity: function setIncentivesController(address controller) returns()
func (_Atoken *AtokenTransactor) SetIncentivesController(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "setIncentivesController", controller)
}

// SetIncentivesController is a paid mutator transaction binding the contract method 0xe655dbd8.
//
// Solidity: function setIncentivesController(address controller) returns()
func (_Atoken *AtokenSession) SetIncentivesController(controller common.Address) (*types.Transaction, error) {
	return _Atoken.Contract.SetIncentivesController(&_Atoken.TransactOpts, controller)
}

// SetIncentivesController is a paid mutator transaction binding the contract method 0xe655dbd8.
//
// Solidity: function setIncentivesController(address controller) returns()
func (_Atoken *AtokenTransactorSession) SetIncentivesController(controller common.Address) (*types.Transaction, error) {
	return _Atoken.Contract.SetIncentivesController(&_Atoken.TransactOpts, controller)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Atoken *AtokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Atoken *AtokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.Transfer(&_Atoken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Atoken *AtokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.Transfer(&_Atoken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Atoken *AtokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Atoken *AtokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.TransferFrom(&_Atoken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Atoken *AtokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.TransferFrom(&_Atoken.TransactOpts, sender, recipient, amount)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address from, address to, uint256 value) returns()
func (_Atoken *AtokenTransactor) TransferOnLiquidation(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "transferOnLiquidation", from, to, value)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address from, address to, uint256 value) returns()
func (_Atoken *AtokenSession) TransferOnLiquidation(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.TransferOnLiquidation(&_Atoken.TransactOpts, from, to, value)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address from, address to, uint256 value) returns()
func (_Atoken *AtokenTransactorSession) TransferOnLiquidation(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.TransferOnLiquidation(&_Atoken.TransactOpts, from, to, value)
}

// TransferUnderlyingTo is a paid mutator transaction binding the contract method 0x4efecaa5.
//
// Solidity: function transferUnderlyingTo(address target, uint256 amount) returns()
func (_Atoken *AtokenTransactor) TransferUnderlyingTo(opts *bind.TransactOpts, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "transferUnderlyingTo", target, amount)
}

// TransferUnderlyingTo is a paid mutator transaction binding the contract method 0x4efecaa5.
//
// Solidity: function transferUnderlyingTo(address target, uint256 amount) returns()
func (_Atoken *AtokenSession) TransferUnderlyingTo(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.TransferUnderlyingTo(&_Atoken.TransactOpts, target, amount)
}

// TransferUnderlyingTo is a paid mutator transaction binding the contract method 0x4efecaa5.
//
// Solidity: function transferUnderlyingTo(address target, uint256 amount) returns()
func (_Atoken *AtokenTransactorSession) TransferUnderlyingTo(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.TransferUnderlyingTo(&_Atoken.TransactOpts, target, amount)
}

// AtokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Atoken contract.
type AtokenApprovalIterator struct {
	Event *AtokenApproval // Event containing the contract specifics and raw log

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
func (it *AtokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtokenApproval)
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
		it.Event = new(AtokenApproval)
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
func (it *AtokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtokenApproval represents a Approval event raised by the Atoken contract.
type AtokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Atoken *AtokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AtokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Atoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AtokenApprovalIterator{contract: _Atoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Atoken *AtokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AtokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Atoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtokenApproval)
				if err := _Atoken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Atoken *AtokenFilterer) ParseApproval(log types.Log) (*AtokenApproval, error) {
	event := new(AtokenApproval)
	if err := _Atoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtokenBalanceTransferIterator is returned from FilterBalanceTransfer and is used to iterate over the raw logs and unpacked data for BalanceTransfer events raised by the Atoken contract.
type AtokenBalanceTransferIterator struct {
	Event *AtokenBalanceTransfer // Event containing the contract specifics and raw log

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
func (it *AtokenBalanceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtokenBalanceTransfer)
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
		it.Event = new(AtokenBalanceTransfer)
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
func (it *AtokenBalanceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtokenBalanceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtokenBalanceTransfer represents a BalanceTransfer event raised by the Atoken contract.
type AtokenBalanceTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBalanceTransfer is a free log retrieval operation binding the contract event 0x4beccb90f994c31aced7a23b5611020728a23d8ec5cddd1a3e9d97b96fda8666.
//
// Solidity: event BalanceTransfer(address indexed from, address indexed to, uint256 value, uint256 index)
func (_Atoken *AtokenFilterer) FilterBalanceTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AtokenBalanceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Atoken.contract.FilterLogs(opts, "BalanceTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AtokenBalanceTransferIterator{contract: _Atoken.contract, event: "BalanceTransfer", logs: logs, sub: sub}, nil
}

// WatchBalanceTransfer is a free log subscription operation binding the contract event 0x4beccb90f994c31aced7a23b5611020728a23d8ec5cddd1a3e9d97b96fda8666.
//
// Solidity: event BalanceTransfer(address indexed from, address indexed to, uint256 value, uint256 index)
func (_Atoken *AtokenFilterer) WatchBalanceTransfer(opts *bind.WatchOpts, sink chan<- *AtokenBalanceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Atoken.contract.WatchLogs(opts, "BalanceTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtokenBalanceTransfer)
				if err := _Atoken.contract.UnpackLog(event, "BalanceTransfer", log); err != nil {
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

// ParseBalanceTransfer is a log parse operation binding the contract event 0x4beccb90f994c31aced7a23b5611020728a23d8ec5cddd1a3e9d97b96fda8666.
//
// Solidity: event BalanceTransfer(address indexed from, address indexed to, uint256 value, uint256 index)
func (_Atoken *AtokenFilterer) ParseBalanceTransfer(log types.Log) (*AtokenBalanceTransfer, error) {
	event := new(AtokenBalanceTransfer)
	if err := _Atoken.contract.UnpackLog(event, "BalanceTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Atoken contract.
type AtokenBurnIterator struct {
	Event *AtokenBurn // Event containing the contract specifics and raw log

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
func (it *AtokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtokenBurn)
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
		it.Event = new(AtokenBurn)
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
func (it *AtokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtokenBurn represents a Burn event raised by the Atoken contract.
type AtokenBurn struct {
	From            common.Address
	Target          common.Address
	Value           *big.Int
	BalanceIncrease *big.Int
	Index           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x4cf25bc1d991c17529c25213d3cc0cda295eeaad5f13f361969b12ea48015f90.
//
// Solidity: event Burn(address indexed from, address indexed target, uint256 value, uint256 balanceIncrease, uint256 index)
func (_Atoken *AtokenFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address, target []common.Address) (*AtokenBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Atoken.contract.FilterLogs(opts, "Burn", fromRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &AtokenBurnIterator{contract: _Atoken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x4cf25bc1d991c17529c25213d3cc0cda295eeaad5f13f361969b12ea48015f90.
//
// Solidity: event Burn(address indexed from, address indexed target, uint256 value, uint256 balanceIncrease, uint256 index)
func (_Atoken *AtokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *AtokenBurn, from []common.Address, target []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Atoken.contract.WatchLogs(opts, "Burn", fromRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtokenBurn)
				if err := _Atoken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x4cf25bc1d991c17529c25213d3cc0cda295eeaad5f13f361969b12ea48015f90.
//
// Solidity: event Burn(address indexed from, address indexed target, uint256 value, uint256 balanceIncrease, uint256 index)
func (_Atoken *AtokenFilterer) ParseBurn(log types.Log) (*AtokenBurn, error) {
	event := new(AtokenBurn)
	if err := _Atoken.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Atoken contract.
type AtokenInitializedIterator struct {
	Event *AtokenInitialized // Event containing the contract specifics and raw log

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
func (it *AtokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtokenInitialized)
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
		it.Event = new(AtokenInitialized)
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
func (it *AtokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtokenInitialized represents a Initialized event raised by the Atoken contract.
type AtokenInitialized struct {
	UnderlyingAsset      common.Address
	Pool                 common.Address
	Treasury             common.Address
	IncentivesController common.Address
	ATokenDecimals       uint8
	ATokenName           string
	ATokenSymbol         string
	Params               []byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xb19e051f8af41150ccccb3fc2c2d8d15f4a4cf434f32a559ba75fe73d6eea20b.
//
// Solidity: event Initialized(address indexed underlyingAsset, address indexed pool, address treasury, address incentivesController, uint8 aTokenDecimals, string aTokenName, string aTokenSymbol, bytes params)
func (_Atoken *AtokenFilterer) FilterInitialized(opts *bind.FilterOpts, underlyingAsset []common.Address, pool []common.Address) (*AtokenInitializedIterator, error) {

	var underlyingAssetRule []interface{}
	for _, underlyingAssetItem := range underlyingAsset {
		underlyingAssetRule = append(underlyingAssetRule, underlyingAssetItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Atoken.contract.FilterLogs(opts, "Initialized", underlyingAssetRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &AtokenInitializedIterator{contract: _Atoken.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xb19e051f8af41150ccccb3fc2c2d8d15f4a4cf434f32a559ba75fe73d6eea20b.
//
// Solidity: event Initialized(address indexed underlyingAsset, address indexed pool, address treasury, address incentivesController, uint8 aTokenDecimals, string aTokenName, string aTokenSymbol, bytes params)
func (_Atoken *AtokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AtokenInitialized, underlyingAsset []common.Address, pool []common.Address) (event.Subscription, error) {

	var underlyingAssetRule []interface{}
	for _, underlyingAssetItem := range underlyingAsset {
		underlyingAssetRule = append(underlyingAssetRule, underlyingAssetItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Atoken.contract.WatchLogs(opts, "Initialized", underlyingAssetRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtokenInitialized)
				if err := _Atoken.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xb19e051f8af41150ccccb3fc2c2d8d15f4a4cf434f32a559ba75fe73d6eea20b.
//
// Solidity: event Initialized(address indexed underlyingAsset, address indexed pool, address treasury, address incentivesController, uint8 aTokenDecimals, string aTokenName, string aTokenSymbol, bytes params)
func (_Atoken *AtokenFilterer) ParseInitialized(log types.Log) (*AtokenInitialized, error) {
	event := new(AtokenInitialized)
	if err := _Atoken.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Atoken contract.
type AtokenMintIterator struct {
	Event *AtokenMint // Event containing the contract specifics and raw log

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
func (it *AtokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtokenMint)
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
		it.Event = new(AtokenMint)
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
func (it *AtokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtokenMint represents a Mint event raised by the Atoken contract.
type AtokenMint struct {
	Caller          common.Address
	OnBehalfOf      common.Address
	Value           *big.Int
	BalanceIncrease *big.Int
	Index           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x458f5fa412d0f69b08dd84872b0215675cc67bc1d5b6fd93300a1c3878b86196.
//
// Solidity: event Mint(address indexed caller, address indexed onBehalfOf, uint256 value, uint256 balanceIncrease, uint256 index)
func (_Atoken *AtokenFilterer) FilterMint(opts *bind.FilterOpts, caller []common.Address, onBehalfOf []common.Address) (*AtokenMintIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _Atoken.contract.FilterLogs(opts, "Mint", callerRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &AtokenMintIterator{contract: _Atoken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x458f5fa412d0f69b08dd84872b0215675cc67bc1d5b6fd93300a1c3878b86196.
//
// Solidity: event Mint(address indexed caller, address indexed onBehalfOf, uint256 value, uint256 balanceIncrease, uint256 index)
func (_Atoken *AtokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *AtokenMint, caller []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _Atoken.contract.WatchLogs(opts, "Mint", callerRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtokenMint)
				if err := _Atoken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x458f5fa412d0f69b08dd84872b0215675cc67bc1d5b6fd93300a1c3878b86196.
//
// Solidity: event Mint(address indexed caller, address indexed onBehalfOf, uint256 value, uint256 balanceIncrease, uint256 index)
func (_Atoken *AtokenFilterer) ParseMint(log types.Log) (*AtokenMint, error) {
	event := new(AtokenMint)
	if err := _Atoken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Atoken contract.
type AtokenTransferIterator struct {
	Event *AtokenTransfer // Event containing the contract specifics and raw log

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
func (it *AtokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtokenTransfer)
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
		it.Event = new(AtokenTransfer)
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
func (it *AtokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtokenTransfer represents a Transfer event raised by the Atoken contract.
type AtokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Atoken *AtokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AtokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Atoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AtokenTransferIterator{contract: _Atoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Atoken *AtokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AtokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Atoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtokenTransfer)
				if err := _Atoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Atoken *AtokenFilterer) ParseTransfer(log types.Log) (*AtokenTransfer, error) {
	event := new(AtokenTransfer)
	if err := _Atoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
