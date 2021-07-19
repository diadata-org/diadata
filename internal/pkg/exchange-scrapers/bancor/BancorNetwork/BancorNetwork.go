// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BancorNetwork

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

// BancorNetworkABI is the input ABI used to generate the binding from.
const BancorNetworkABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_onlyOwnerCanUpdateRegistry\",\"type\":\"bool\"}],\"name\":\"restrictRegistryUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_register\",\"type\":\"bool\"}],\"name\":\"registerEtherToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getReturnByPath\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"claimAndConvertFor2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"onlyOwnerCanUpdateRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"convert2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAffiliateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prevRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"rateByPath\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"etherTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_bancorX\",\"type\":\"address\"},{\"name\":\"_conversionId\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"completeXConversion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"convertFor2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"claimAndConvertFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restoreRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"convertByPath\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_targetBlockchain\",\"type\":\"bytes32\"},{\"name\":\"_targetAccount\",\"type\":\"bytes32\"},{\"name\":\"_conversionId\",\"type\":\"uint256\"}],\"name\":\"xConvert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"claimAndConvert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"convertFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_targetBlockchain\",\"type\":\"bytes32\"},{\"name\":\"_targetAccount\",\"type\":\"bytes32\"},{\"name\":\"_conversionId\",\"type\":\"uint256\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"xConvert2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sourceToken\",\"type\":\"address\"},{\"name\":\"_targetToken\",\"type\":\"address\"}],\"name\":\"conversionPath\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"claimAndConvert2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"convert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxAffiliateFee\",\"type\":\"uint256\"}],\"name\":\"setMaxAffiliateFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_smartToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_fromToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_toToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"Conversion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// BancorNetwork is an auto generated Go binding around an Ethereum contract.
type BancorNetwork struct {
	BancorNetworkCaller     // Read-only binding to the contract
	BancorNetworkTransactor // Write-only binding to the contract
	BancorNetworkFilterer   // Log filterer for contract events
}

// BancorNetworkCaller is an auto generated read-only Go binding around an Ethereum contract.
type BancorNetworkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BancorNetworkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BancorNetworkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BancorNetworkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BancorNetworkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BancorNetworkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BancorNetworkSession struct {
	Contract     *BancorNetwork    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BancorNetworkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BancorNetworkCallerSession struct {
	Contract *BancorNetworkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BancorNetworkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BancorNetworkTransactorSession struct {
	Contract     *BancorNetworkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BancorNetworkRaw is an auto generated low-level Go binding around an Ethereum contract.
type BancorNetworkRaw struct {
	Contract *BancorNetwork // Generic contract binding to access the raw methods on
}

// BancorNetworkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BancorNetworkCallerRaw struct {
	Contract *BancorNetworkCaller // Generic read-only contract binding to access the raw methods on
}

// BancorNetworkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BancorNetworkTransactorRaw struct {
	Contract *BancorNetworkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBancorNetwork creates a new instance of BancorNetwork, bound to a specific deployed contract.
func NewBancorNetwork(address common.Address, backend bind.ContractBackend) (*BancorNetwork, error) {
	contract, err := bindBancorNetwork(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BancorNetwork{BancorNetworkCaller: BancorNetworkCaller{contract: contract}, BancorNetworkTransactor: BancorNetworkTransactor{contract: contract}, BancorNetworkFilterer: BancorNetworkFilterer{contract: contract}}, nil
}

// NewBancorNetworkCaller creates a new read-only instance of BancorNetwork, bound to a specific deployed contract.
func NewBancorNetworkCaller(address common.Address, caller bind.ContractCaller) (*BancorNetworkCaller, error) {
	contract, err := bindBancorNetwork(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BancorNetworkCaller{contract: contract}, nil
}

// NewBancorNetworkTransactor creates a new write-only instance of BancorNetwork, bound to a specific deployed contract.
func NewBancorNetworkTransactor(address common.Address, transactor bind.ContractTransactor) (*BancorNetworkTransactor, error) {
	contract, err := bindBancorNetwork(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BancorNetworkTransactor{contract: contract}, nil
}

// NewBancorNetworkFilterer creates a new log filterer instance of BancorNetwork, bound to a specific deployed contract.
func NewBancorNetworkFilterer(address common.Address, filterer bind.ContractFilterer) (*BancorNetworkFilterer, error) {
	contract, err := bindBancorNetwork(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BancorNetworkFilterer{contract: contract}, nil
}

// bindBancorNetwork binds a generic wrapper to an already deployed contract.
func bindBancorNetwork(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BancorNetworkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BancorNetwork *BancorNetworkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BancorNetwork.Contract.BancorNetworkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BancorNetwork *BancorNetworkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BancorNetwork.Contract.BancorNetworkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BancorNetwork *BancorNetworkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BancorNetwork.Contract.BancorNetworkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BancorNetwork *BancorNetworkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BancorNetwork.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BancorNetwork *BancorNetworkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BancorNetwork.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BancorNetwork *BancorNetworkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BancorNetwork.Contract.contract.Transact(opts, method, params...)
}

// ConversionPath is a free data retrieval call binding the contract method 0xd734fa19.
//
// Solidity: function conversionPath(address _sourceToken, address _targetToken) view returns(address[])
func (_BancorNetwork *BancorNetworkCaller) ConversionPath(opts *bind.CallOpts, _sourceToken common.Address, _targetToken common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _BancorNetwork.contract.Call(opts, &out, "conversionPath", _sourceToken, _targetToken)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ConversionPath is a free data retrieval call binding the contract method 0xd734fa19.
//
// Solidity: function conversionPath(address _sourceToken, address _targetToken) view returns(address[])
func (_BancorNetwork *BancorNetworkSession) ConversionPath(_sourceToken common.Address, _targetToken common.Address) ([]common.Address, error) {
	return _BancorNetwork.Contract.ConversionPath(&_BancorNetwork.CallOpts, _sourceToken, _targetToken)
}

// ConversionPath is a free data retrieval call binding the contract method 0xd734fa19.
//
// Solidity: function conversionPath(address _sourceToken, address _targetToken) view returns(address[])
func (_BancorNetwork *BancorNetworkCallerSession) ConversionPath(_sourceToken common.Address, _targetToken common.Address) ([]common.Address, error) {
	return _BancorNetwork.Contract.ConversionPath(&_BancorNetwork.CallOpts, _sourceToken, _targetToken)
}

// EtherTokens is a free data retrieval call binding the contract method 0x8077ccf7.
//
// Solidity: function etherTokens(address ) view returns(bool)
func (_BancorNetwork *BancorNetworkCaller) EtherTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BancorNetwork.contract.Call(opts, &out, "etherTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EtherTokens is a free data retrieval call binding the contract method 0x8077ccf7.
//
// Solidity: function etherTokens(address ) view returns(bool)
func (_BancorNetwork *BancorNetworkSession) EtherTokens(arg0 common.Address) (bool, error) {
	return _BancorNetwork.Contract.EtherTokens(&_BancorNetwork.CallOpts, arg0)
}

// EtherTokens is a free data retrieval call binding the contract method 0x8077ccf7.
//
// Solidity: function etherTokens(address ) view returns(bool)
func (_BancorNetwork *BancorNetworkCallerSession) EtherTokens(arg0 common.Address) (bool, error) {
	return _BancorNetwork.Contract.EtherTokens(&_BancorNetwork.CallOpts, arg0)
}

// GetReturnByPath is a free data retrieval call binding the contract method 0x0c8496cc.
//
// Solidity: function getReturnByPath(address[] _path, uint256 _amount) view returns(uint256, uint256)
func (_BancorNetwork *BancorNetworkCaller) GetReturnByPath(opts *bind.CallOpts, _path []common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _BancorNetwork.contract.Call(opts, &out, "getReturnByPath", _path, _amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetReturnByPath is a free data retrieval call binding the contract method 0x0c8496cc.
//
// Solidity: function getReturnByPath(address[] _path, uint256 _amount) view returns(uint256, uint256)
func (_BancorNetwork *BancorNetworkSession) GetReturnByPath(_path []common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _BancorNetwork.Contract.GetReturnByPath(&_BancorNetwork.CallOpts, _path, _amount)
}

// GetReturnByPath is a free data retrieval call binding the contract method 0x0c8496cc.
//
// Solidity: function getReturnByPath(address[] _path, uint256 _amount) view returns(uint256, uint256)
func (_BancorNetwork *BancorNetworkCallerSession) GetReturnByPath(_path []common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _BancorNetwork.Contract.GetReturnByPath(&_BancorNetwork.CallOpts, _path, _amount)
}

// MaxAffiliateFee is a free data retrieval call binding the contract method 0x5d732ff2.
//
// Solidity: function maxAffiliateFee() view returns(uint256)
func (_BancorNetwork *BancorNetworkCaller) MaxAffiliateFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BancorNetwork.contract.Call(opts, &out, "maxAffiliateFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAffiliateFee is a free data retrieval call binding the contract method 0x5d732ff2.
//
// Solidity: function maxAffiliateFee() view returns(uint256)
func (_BancorNetwork *BancorNetworkSession) MaxAffiliateFee() (*big.Int, error) {
	return _BancorNetwork.Contract.MaxAffiliateFee(&_BancorNetwork.CallOpts)
}

// MaxAffiliateFee is a free data retrieval call binding the contract method 0x5d732ff2.
//
// Solidity: function maxAffiliateFee() view returns(uint256)
func (_BancorNetwork *BancorNetworkCallerSession) MaxAffiliateFee() (*big.Int, error) {
	return _BancorNetwork.Contract.MaxAffiliateFee(&_BancorNetwork.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_BancorNetwork *BancorNetworkCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BancorNetwork.contract.Call(opts, &out, "newOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_BancorNetwork *BancorNetworkSession) NewOwner() (common.Address, error) {
	return _BancorNetwork.Contract.NewOwner(&_BancorNetwork.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_BancorNetwork *BancorNetworkCallerSession) NewOwner() (common.Address, error) {
	return _BancorNetwork.Contract.NewOwner(&_BancorNetwork.CallOpts)
}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_BancorNetwork *BancorNetworkCaller) OnlyOwnerCanUpdateRegistry(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BancorNetwork.contract.Call(opts, &out, "onlyOwnerCanUpdateRegistry")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_BancorNetwork *BancorNetworkSession) OnlyOwnerCanUpdateRegistry() (bool, error) {
	return _BancorNetwork.Contract.OnlyOwnerCanUpdateRegistry(&_BancorNetwork.CallOpts)
}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_BancorNetwork *BancorNetworkCallerSession) OnlyOwnerCanUpdateRegistry() (bool, error) {
	return _BancorNetwork.Contract.OnlyOwnerCanUpdateRegistry(&_BancorNetwork.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BancorNetwork *BancorNetworkCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BancorNetwork.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BancorNetwork *BancorNetworkSession) Owner() (common.Address, error) {
	return _BancorNetwork.Contract.Owner(&_BancorNetwork.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BancorNetwork *BancorNetworkCallerSession) Owner() (common.Address, error) {
	return _BancorNetwork.Contract.Owner(&_BancorNetwork.CallOpts)
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_BancorNetwork *BancorNetworkCaller) PrevRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BancorNetwork.contract.Call(opts, &out, "prevRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_BancorNetwork *BancorNetworkSession) PrevRegistry() (common.Address, error) {
	return _BancorNetwork.Contract.PrevRegistry(&_BancorNetwork.CallOpts)
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_BancorNetwork *BancorNetworkCallerSession) PrevRegistry() (common.Address, error) {
	return _BancorNetwork.Contract.PrevRegistry(&_BancorNetwork.CallOpts)
}

// RateByPath is a free data retrieval call binding the contract method 0x7f9c0ecd.
//
// Solidity: function rateByPath(address[] _path, uint256 _amount) view returns(uint256)
func (_BancorNetwork *BancorNetworkCaller) RateByPath(opts *bind.CallOpts, _path []common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BancorNetwork.contract.Call(opts, &out, "rateByPath", _path, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RateByPath is a free data retrieval call binding the contract method 0x7f9c0ecd.
//
// Solidity: function rateByPath(address[] _path, uint256 _amount) view returns(uint256)
func (_BancorNetwork *BancorNetworkSession) RateByPath(_path []common.Address, _amount *big.Int) (*big.Int, error) {
	return _BancorNetwork.Contract.RateByPath(&_BancorNetwork.CallOpts, _path, _amount)
}

// RateByPath is a free data retrieval call binding the contract method 0x7f9c0ecd.
//
// Solidity: function rateByPath(address[] _path, uint256 _amount) view returns(uint256)
func (_BancorNetwork *BancorNetworkCallerSession) RateByPath(_path []common.Address, _amount *big.Int) (*big.Int, error) {
	return _BancorNetwork.Contract.RateByPath(&_BancorNetwork.CallOpts, _path, _amount)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_BancorNetwork *BancorNetworkCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BancorNetwork.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_BancorNetwork *BancorNetworkSession) Registry() (common.Address, error) {
	return _BancorNetwork.Contract.Registry(&_BancorNetwork.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_BancorNetwork *BancorNetworkCallerSession) Registry() (common.Address, error) {
	return _BancorNetwork.Contract.Registry(&_BancorNetwork.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BancorNetwork *BancorNetworkTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BancorNetwork *BancorNetworkSession) AcceptOwnership() (*types.Transaction, error) {
	return _BancorNetwork.Contract.AcceptOwnership(&_BancorNetwork.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BancorNetwork *BancorNetworkTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BancorNetwork.Contract.AcceptOwnership(&_BancorNetwork.TransactOpts)
}

// ClaimAndConvert is a paid mutator transaction binding the contract method 0xc7ba24bc.
//
// Solidity: function claimAndConvert(address[] _path, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_BancorNetwork *BancorNetworkTransactor) ClaimAndConvert(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "claimAndConvert", _path, _amount, _minReturn)
}

// ClaimAndConvert is a paid mutator transaction binding the contract method 0xc7ba24bc.
//
// Solidity: function claimAndConvert(address[] _path, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_BancorNetwork *BancorNetworkSession) ClaimAndConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.ClaimAndConvert(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn)
}

// ClaimAndConvert is a paid mutator transaction binding the contract method 0xc7ba24bc.
//
// Solidity: function claimAndConvert(address[] _path, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_BancorNetwork *BancorNetworkTransactorSession) ClaimAndConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.ClaimAndConvert(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn)
}

// ClaimAndConvert2 is a paid mutator transaction binding the contract method 0xe57738e5.
//
// Solidity: function claimAndConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_BancorNetwork *BancorNetworkTransactor) ClaimAndConvert2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "claimAndConvert2", _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// ClaimAndConvert2 is a paid mutator transaction binding the contract method 0xe57738e5.
//
// Solidity: function claimAndConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_BancorNetwork *BancorNetworkSession) ClaimAndConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.ClaimAndConvert2(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// ClaimAndConvert2 is a paid mutator transaction binding the contract method 0xe57738e5.
//
// Solidity: function claimAndConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_BancorNetwork *BancorNetworkTransactorSession) ClaimAndConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.ClaimAndConvert2(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// ClaimAndConvertFor is a paid mutator transaction binding the contract method 0xb1e9932b.
//
// Solidity: function claimAndConvertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_BancorNetwork *BancorNetworkTransactor) ClaimAndConvertFor(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "claimAndConvertFor", _path, _amount, _minReturn, _beneficiary)
}

// ClaimAndConvertFor is a paid mutator transaction binding the contract method 0xb1e9932b.
//
// Solidity: function claimAndConvertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_BancorNetwork *BancorNetworkSession) ClaimAndConvertFor(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _BancorNetwork.Contract.ClaimAndConvertFor(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _beneficiary)
}

// ClaimAndConvertFor is a paid mutator transaction binding the contract method 0xb1e9932b.
//
// Solidity: function claimAndConvertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_BancorNetwork *BancorNetworkTransactorSession) ClaimAndConvertFor(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _BancorNetwork.Contract.ClaimAndConvertFor(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _beneficiary)
}

// ClaimAndConvertFor2 is a paid mutator transaction binding the contract method 0x2978c10e.
//
// Solidity: function claimAndConvertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_BancorNetwork *BancorNetworkTransactor) ClaimAndConvertFor2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "claimAndConvertFor2", _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ClaimAndConvertFor2 is a paid mutator transaction binding the contract method 0x2978c10e.
//
// Solidity: function claimAndConvertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_BancorNetwork *BancorNetworkSession) ClaimAndConvertFor2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.ClaimAndConvertFor2(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ClaimAndConvertFor2 is a paid mutator transaction binding the contract method 0x2978c10e.
//
// Solidity: function claimAndConvertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_BancorNetwork *BancorNetworkTransactorSession) ClaimAndConvertFor2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.ClaimAndConvertFor2(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// CompleteXConversion is a paid mutator transaction binding the contract method 0x89f9cc61.
//
// Solidity: function completeXConversion(address[] _path, address _bancorX, uint256 _conversionId, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_BancorNetwork *BancorNetworkTransactor) CompleteXConversion(opts *bind.TransactOpts, _path []common.Address, _bancorX common.Address, _conversionId *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "completeXConversion", _path, _bancorX, _conversionId, _minReturn, _beneficiary)
}

// CompleteXConversion is a paid mutator transaction binding the contract method 0x89f9cc61.
//
// Solidity: function completeXConversion(address[] _path, address _bancorX, uint256 _conversionId, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_BancorNetwork *BancorNetworkSession) CompleteXConversion(_path []common.Address, _bancorX common.Address, _conversionId *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _BancorNetwork.Contract.CompleteXConversion(&_BancorNetwork.TransactOpts, _path, _bancorX, _conversionId, _minReturn, _beneficiary)
}

// CompleteXConversion is a paid mutator transaction binding the contract method 0x89f9cc61.
//
// Solidity: function completeXConversion(address[] _path, address _bancorX, uint256 _conversionId, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_BancorNetwork *BancorNetworkTransactorSession) CompleteXConversion(_path []common.Address, _bancorX common.Address, _conversionId *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _BancorNetwork.Contract.CompleteXConversion(&_BancorNetwork.TransactOpts, _path, _bancorX, _conversionId, _minReturn, _beneficiary)
}

// Convert is a paid mutator transaction binding the contract method 0xf3898a97.
//
// Solidity: function convert(address[] _path, uint256 _amount, uint256 _minReturn) payable returns(uint256)
func (_BancorNetwork *BancorNetworkTransactor) Convert(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "convert", _path, _amount, _minReturn)
}

// Convert is a paid mutator transaction binding the contract method 0xf3898a97.
//
// Solidity: function convert(address[] _path, uint256 _amount, uint256 _minReturn) payable returns(uint256)
func (_BancorNetwork *BancorNetworkSession) Convert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.Convert(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn)
}

// Convert is a paid mutator transaction binding the contract method 0xf3898a97.
//
// Solidity: function convert(address[] _path, uint256 _amount, uint256 _minReturn) payable returns(uint256)
func (_BancorNetwork *BancorNetworkTransactorSession) Convert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.Convert(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn)
}

// Convert2 is a paid mutator transaction binding the contract method 0x569706eb.
//
// Solidity: function convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_BancorNetwork *BancorNetworkTransactor) Convert2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "convert2", _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// Convert2 is a paid mutator transaction binding the contract method 0x569706eb.
//
// Solidity: function convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_BancorNetwork *BancorNetworkSession) Convert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.Convert2(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// Convert2 is a paid mutator transaction binding the contract method 0x569706eb.
//
// Solidity: function convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_BancorNetwork *BancorNetworkTransactorSession) Convert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.Convert2(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// ConvertByPath is a paid mutator transaction binding the contract method 0xb77d239b.
//
// Solidity: function convertByPath(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_BancorNetwork *BancorNetworkTransactor) ConvertByPath(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "convertByPath", _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ConvertByPath is a paid mutator transaction binding the contract method 0xb77d239b.
//
// Solidity: function convertByPath(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_BancorNetwork *BancorNetworkSession) ConvertByPath(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.ConvertByPath(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ConvertByPath is a paid mutator transaction binding the contract method 0xb77d239b.
//
// Solidity: function convertByPath(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_BancorNetwork *BancorNetworkTransactorSession) ConvertByPath(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.ConvertByPath(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ConvertFor is a paid mutator transaction binding the contract method 0xc98fefed.
//
// Solidity: function convertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) payable returns(uint256)
func (_BancorNetwork *BancorNetworkTransactor) ConvertFor(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "convertFor", _path, _amount, _minReturn, _beneficiary)
}

// ConvertFor is a paid mutator transaction binding the contract method 0xc98fefed.
//
// Solidity: function convertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) payable returns(uint256)
func (_BancorNetwork *BancorNetworkSession) ConvertFor(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _BancorNetwork.Contract.ConvertFor(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _beneficiary)
}

// ConvertFor is a paid mutator transaction binding the contract method 0xc98fefed.
//
// Solidity: function convertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) payable returns(uint256)
func (_BancorNetwork *BancorNetworkTransactorSession) ConvertFor(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _BancorNetwork.Contract.ConvertFor(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _beneficiary)
}

// ConvertFor2 is a paid mutator transaction binding the contract method 0xab6214ce.
//
// Solidity: function convertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_BancorNetwork *BancorNetworkTransactor) ConvertFor2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "convertFor2", _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ConvertFor2 is a paid mutator transaction binding the contract method 0xab6214ce.
//
// Solidity: function convertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_BancorNetwork *BancorNetworkSession) ConvertFor2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.ConvertFor2(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ConvertFor2 is a paid mutator transaction binding the contract method 0xab6214ce.
//
// Solidity: function convertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_BancorNetwork *BancorNetworkTransactorSession) ConvertFor2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.ConvertFor2(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// RegisterEtherToken is a paid mutator transaction binding the contract method 0x02ef521e.
//
// Solidity: function registerEtherToken(address _token, bool _register) returns()
func (_BancorNetwork *BancorNetworkTransactor) RegisterEtherToken(opts *bind.TransactOpts, _token common.Address, _register bool) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "registerEtherToken", _token, _register)
}

// RegisterEtherToken is a paid mutator transaction binding the contract method 0x02ef521e.
//
// Solidity: function registerEtherToken(address _token, bool _register) returns()
func (_BancorNetwork *BancorNetworkSession) RegisterEtherToken(_token common.Address, _register bool) (*types.Transaction, error) {
	return _BancorNetwork.Contract.RegisterEtherToken(&_BancorNetwork.TransactOpts, _token, _register)
}

// RegisterEtherToken is a paid mutator transaction binding the contract method 0x02ef521e.
//
// Solidity: function registerEtherToken(address _token, bool _register) returns()
func (_BancorNetwork *BancorNetworkTransactorSession) RegisterEtherToken(_token common.Address, _register bool) (*types.Transaction, error) {
	return _BancorNetwork.Contract.RegisterEtherToken(&_BancorNetwork.TransactOpts, _token, _register)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_BancorNetwork *BancorNetworkTransactor) RestoreRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "restoreRegistry")
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_BancorNetwork *BancorNetworkSession) RestoreRegistry() (*types.Transaction, error) {
	return _BancorNetwork.Contract.RestoreRegistry(&_BancorNetwork.TransactOpts)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_BancorNetwork *BancorNetworkTransactorSession) RestoreRegistry() (*types.Transaction, error) {
	return _BancorNetwork.Contract.RestoreRegistry(&_BancorNetwork.TransactOpts)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_BancorNetwork *BancorNetworkTransactor) RestrictRegistryUpdate(opts *bind.TransactOpts, _onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "restrictRegistryUpdate", _onlyOwnerCanUpdateRegistry)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_BancorNetwork *BancorNetworkSession) RestrictRegistryUpdate(_onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _BancorNetwork.Contract.RestrictRegistryUpdate(&_BancorNetwork.TransactOpts, _onlyOwnerCanUpdateRegistry)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_BancorNetwork *BancorNetworkTransactorSession) RestrictRegistryUpdate(_onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _BancorNetwork.Contract.RestrictRegistryUpdate(&_BancorNetwork.TransactOpts, _onlyOwnerCanUpdateRegistry)
}

// SetMaxAffiliateFee is a paid mutator transaction binding the contract method 0xf3bc7d2a.
//
// Solidity: function setMaxAffiliateFee(uint256 _maxAffiliateFee) returns()
func (_BancorNetwork *BancorNetworkTransactor) SetMaxAffiliateFee(opts *bind.TransactOpts, _maxAffiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "setMaxAffiliateFee", _maxAffiliateFee)
}

// SetMaxAffiliateFee is a paid mutator transaction binding the contract method 0xf3bc7d2a.
//
// Solidity: function setMaxAffiliateFee(uint256 _maxAffiliateFee) returns()
func (_BancorNetwork *BancorNetworkSession) SetMaxAffiliateFee(_maxAffiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.SetMaxAffiliateFee(&_BancorNetwork.TransactOpts, _maxAffiliateFee)
}

// SetMaxAffiliateFee is a paid mutator transaction binding the contract method 0xf3bc7d2a.
//
// Solidity: function setMaxAffiliateFee(uint256 _maxAffiliateFee) returns()
func (_BancorNetwork *BancorNetworkTransactorSession) SetMaxAffiliateFee(_maxAffiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.SetMaxAffiliateFee(&_BancorNetwork.TransactOpts, _maxAffiliateFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_BancorNetwork *BancorNetworkTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_BancorNetwork *BancorNetworkSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BancorNetwork.Contract.TransferOwnership(&_BancorNetwork.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_BancorNetwork *BancorNetworkTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BancorNetwork.Contract.TransferOwnership(&_BancorNetwork.TransactOpts, _newOwner)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_BancorNetwork *BancorNetworkTransactor) UpdateRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "updateRegistry")
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_BancorNetwork *BancorNetworkSession) UpdateRegistry() (*types.Transaction, error) {
	return _BancorNetwork.Contract.UpdateRegistry(&_BancorNetwork.TransactOpts)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_BancorNetwork *BancorNetworkTransactorSession) UpdateRegistry() (*types.Transaction, error) {
	return _BancorNetwork.Contract.UpdateRegistry(&_BancorNetwork.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_BancorNetwork *BancorNetworkTransactor) WithdrawTokens(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "withdrawTokens", _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_BancorNetwork *BancorNetworkSession) WithdrawTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.WithdrawTokens(&_BancorNetwork.TransactOpts, _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_BancorNetwork *BancorNetworkTransactorSession) WithdrawTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.WithdrawTokens(&_BancorNetwork.TransactOpts, _token, _to, _amount)
}

// XConvert is a paid mutator transaction binding the contract method 0xc52173de.
//
// Solidity: function xConvert(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId) payable returns(uint256)
func (_BancorNetwork *BancorNetworkTransactor) XConvert(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "xConvert", _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId)
}

// XConvert is a paid mutator transaction binding the contract method 0xc52173de.
//
// Solidity: function xConvert(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId) payable returns(uint256)
func (_BancorNetwork *BancorNetworkSession) XConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.XConvert(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId)
}

// XConvert is a paid mutator transaction binding the contract method 0xc52173de.
//
// Solidity: function xConvert(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId) payable returns(uint256)
func (_BancorNetwork *BancorNetworkTransactorSession) XConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.XConvert(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId)
}

// XConvert2 is a paid mutator transaction binding the contract method 0xcb32564e.
//
// Solidity: function xConvert2(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_BancorNetwork *BancorNetworkTransactor) XConvert2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.contract.Transact(opts, "xConvert2", _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId, _affiliateAccount, _affiliateFee)
}

// XConvert2 is a paid mutator transaction binding the contract method 0xcb32564e.
//
// Solidity: function xConvert2(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_BancorNetwork *BancorNetworkSession) XConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.XConvert2(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId, _affiliateAccount, _affiliateFee)
}

// XConvert2 is a paid mutator transaction binding the contract method 0xcb32564e.
//
// Solidity: function xConvert2(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_BancorNetwork *BancorNetworkTransactorSession) XConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _BancorNetwork.Contract.XConvert2(&_BancorNetwork.TransactOpts, _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId, _affiliateAccount, _affiliateFee)
}

// BancorNetworkConversionIterator is returned from FilterConversion and is used to iterate over the raw logs and unpacked data for Conversion events raised by the BancorNetwork contract.
type BancorNetworkConversionIterator struct {
	Event *BancorNetworkConversion // Event containing the contract specifics and raw log

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
func (it *BancorNetworkConversionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BancorNetworkConversion)
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
		it.Event = new(BancorNetworkConversion)
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
func (it *BancorNetworkConversionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BancorNetworkConversionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BancorNetworkConversion represents a Conversion event raised by the BancorNetwork contract.
type BancorNetworkConversion struct {
	SmartToken common.Address
	FromToken  common.Address
	ToToken    common.Address
	FromAmount *big.Int
	ToAmount   *big.Int
	Trader     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConversion is a free log retrieval operation binding the contract event 0x7154b38b5dd31bb3122436a96d4e09aba5b323ae1fd580025fab55074334c095.
//
// Solidity: event Conversion(address indexed _smartToken, address indexed _fromToken, address indexed _toToken, uint256 _fromAmount, uint256 _toAmount, address _trader)
func (_BancorNetwork *BancorNetworkFilterer) FilterConversion(opts *bind.FilterOpts, _smartToken []common.Address, _fromToken []common.Address, _toToken []common.Address) (*BancorNetworkConversionIterator, error) {

	var _smartTokenRule []interface{}
	for _, _smartTokenItem := range _smartToken {
		_smartTokenRule = append(_smartTokenRule, _smartTokenItem)
	}
	var _fromTokenRule []interface{}
	for _, _fromTokenItem := range _fromToken {
		_fromTokenRule = append(_fromTokenRule, _fromTokenItem)
	}
	var _toTokenRule []interface{}
	for _, _toTokenItem := range _toToken {
		_toTokenRule = append(_toTokenRule, _toTokenItem)
	}

	logs, sub, err := _BancorNetwork.contract.FilterLogs(opts, "Conversion", _smartTokenRule, _fromTokenRule, _toTokenRule)
	if err != nil {
		return nil, err
	}
	return &BancorNetworkConversionIterator{contract: _BancorNetwork.contract, event: "Conversion", logs: logs, sub: sub}, nil
}

// WatchConversion is a free log subscription operation binding the contract event 0x7154b38b5dd31bb3122436a96d4e09aba5b323ae1fd580025fab55074334c095.
//
// Solidity: event Conversion(address indexed _smartToken, address indexed _fromToken, address indexed _toToken, uint256 _fromAmount, uint256 _toAmount, address _trader)
func (_BancorNetwork *BancorNetworkFilterer) WatchConversion(opts *bind.WatchOpts, sink chan<- *BancorNetworkConversion, _smartToken []common.Address, _fromToken []common.Address, _toToken []common.Address) (event.Subscription, error) {

	var _smartTokenRule []interface{}
	for _, _smartTokenItem := range _smartToken {
		_smartTokenRule = append(_smartTokenRule, _smartTokenItem)
	}
	var _fromTokenRule []interface{}
	for _, _fromTokenItem := range _fromToken {
		_fromTokenRule = append(_fromTokenRule, _fromTokenItem)
	}
	var _toTokenRule []interface{}
	for _, _toTokenItem := range _toToken {
		_toTokenRule = append(_toTokenRule, _toTokenItem)
	}

	logs, sub, err := _BancorNetwork.contract.WatchLogs(opts, "Conversion", _smartTokenRule, _fromTokenRule, _toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BancorNetworkConversion)
				if err := _BancorNetwork.contract.UnpackLog(event, "Conversion", log); err != nil {
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

// ParseConversion is a log parse operation binding the contract event 0x7154b38b5dd31bb3122436a96d4e09aba5b323ae1fd580025fab55074334c095.
//
// Solidity: event Conversion(address indexed _smartToken, address indexed _fromToken, address indexed _toToken, uint256 _fromAmount, uint256 _toAmount, address _trader)
func (_BancorNetwork *BancorNetworkFilterer) ParseConversion(log types.Log) (*BancorNetworkConversion, error) {
	event := new(BancorNetworkConversion)
	if err := _BancorNetwork.contract.UnpackLog(event, "Conversion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BancorNetworkOwnerUpdateIterator is returned from FilterOwnerUpdate and is used to iterate over the raw logs and unpacked data for OwnerUpdate events raised by the BancorNetwork contract.
type BancorNetworkOwnerUpdateIterator struct {
	Event *BancorNetworkOwnerUpdate // Event containing the contract specifics and raw log

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
func (it *BancorNetworkOwnerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BancorNetworkOwnerUpdate)
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
		it.Event = new(BancorNetworkOwnerUpdate)
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
func (it *BancorNetworkOwnerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BancorNetworkOwnerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BancorNetworkOwnerUpdate represents a OwnerUpdate event raised by the BancorNetwork contract.
type BancorNetworkOwnerUpdate struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdate is a free log retrieval operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_BancorNetwork *BancorNetworkFilterer) FilterOwnerUpdate(opts *bind.FilterOpts, _prevOwner []common.Address, _newOwner []common.Address) (*BancorNetworkOwnerUpdateIterator, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _BancorNetwork.contract.FilterLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BancorNetworkOwnerUpdateIterator{contract: _BancorNetwork.contract, event: "OwnerUpdate", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdate is a free log subscription operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_BancorNetwork *BancorNetworkFilterer) WatchOwnerUpdate(opts *bind.WatchOpts, sink chan<- *BancorNetworkOwnerUpdate, _prevOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _BancorNetwork.contract.WatchLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BancorNetworkOwnerUpdate)
				if err := _BancorNetwork.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
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

// ParseOwnerUpdate is a log parse operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_BancorNetwork *BancorNetworkFilterer) ParseOwnerUpdate(log types.Log) (*BancorNetworkOwnerUpdate, error) {
	event := new(BancorNetworkOwnerUpdate)
	if err := _BancorNetwork.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
