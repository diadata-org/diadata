// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package diaOracleV2MultiupdateService

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

// DiaOracleV2MultiupdateServiceMetaData contains all meta data concerning the DiaOracleV2MultiupdateService contract.
var DiaOracleV2MultiupdateServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"name\":\"OracleUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUpdater\",\"type\":\"address\"}],\"name\":\"UpdaterAddressChange\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"compressedValues\",\"type\":\"uint256[]\"}],\"name\":\"setMultipleValues\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOracleUpdaterAddress\",\"type\":\"address\"}],\"name\":\"updateOracleUpdaterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610d2b806100616000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80635a9ade8b1461005c5780636aa45efc1461008c5780637898e0c2146100a85780638d241526146100c4578063960384a0146100e0575b600080fd5b6100766004803603810190610071919061064d565b610111565b60405161008391906106af565b60405180910390f35b6100a660048036038101906100a19190610728565b61013f565b005b6100c260048036038101906100bd919061079d565b610214565b005b6100de60048036038101906100d991906109e1565b61030a565b005b6100fa60048036038101906100f5919061064d565b610496565b604051610108929190610a68565b60405180910390f35b6000818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461019957600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f816040516102099190610aa0565b60405180910390a150565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461026e57600080fd5b6000816fffffffffffffffffffffffffffffffff166080846fffffffffffffffffffffffffffffffff16901b6102a49190610aea565b9050806000856040516102b79190610b8f565b9081526020016040518091039020819055507fa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee22917828484846040516102fc93929190610bf0565b60405180910390a150505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461036457600080fd5b805182511461037257600080fd5b60005b8251816fffffffffffffffffffffffffffffffff16101561049157600083826fffffffffffffffffffffffffffffffff16815181106103b7576103b6610c2e565b5b60200260200101519050600083836fffffffffffffffffffffffffffffffff16815181106103e8576103e7610c2e565b5b602002602001015190506000608082901c905060007001000000000000000000000000000000008361041a9190610c8c565b90508260008560405161042d9190610b8f565b9081526020016040518091039020819055507fa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee229178284838360405161047293929190610bf0565b60405180910390a150505050808061048990610cbd565b915050610375565b505050565b600080600080846040516104aa9190610b8f565b90815260200160405180910390205490506000700100000000000000000000000000000000826104da9190610c8c565b90506000608083901c9050808294509450505050915091565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61055a82610511565b810181811067ffffffffffffffff8211171561057957610578610522565b5b80604052505050565b600061058c6104f3565b90506105988282610551565b919050565b600067ffffffffffffffff8211156105b8576105b7610522565b5b6105c182610511565b9050602081019050919050565b82818337600083830152505050565b60006105f06105eb8461059d565b610582565b90508281526020810184848401111561060c5761060b61050c565b5b6106178482856105ce565b509392505050565b600082601f83011261063457610633610507565b5b81356106448482602086016105dd565b91505092915050565b600060208284031215610663576106626104fd565b5b600082013567ffffffffffffffff81111561068157610680610502565b5b61068d8482850161061f565b91505092915050565b6000819050919050565b6106a981610696565b82525050565b60006020820190506106c460008301846106a0565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006106f5826106ca565b9050919050565b610705816106ea565b811461071057600080fd5b50565b600081359050610722816106fc565b92915050565b60006020828403121561073e5761073d6104fd565b5b600061074c84828501610713565b91505092915050565b60006fffffffffffffffffffffffffffffffff82169050919050565b61077a81610755565b811461078557600080fd5b50565b60008135905061079781610771565b92915050565b6000806000606084860312156107b6576107b56104fd565b5b600084013567ffffffffffffffff8111156107d4576107d3610502565b5b6107e08682870161061f565b93505060206107f186828701610788565b925050604061080286828701610788565b9150509250925092565b600067ffffffffffffffff82111561082757610826610522565b5b602082029050602081019050919050565b600080fd5b600061085061084b8461080c565b610582565b9050808382526020820190506020840283018581111561087357610872610838565b5b835b818110156108ba57803567ffffffffffffffff81111561089857610897610507565b5b8086016108a5898261061f565b85526020850194505050602081019050610875565b5050509392505050565b600082601f8301126108d9576108d8610507565b5b81356108e984826020860161083d565b91505092915050565b600067ffffffffffffffff82111561090d5761090c610522565b5b602082029050602081019050919050565b61092781610696565b811461093257600080fd5b50565b6000813590506109448161091e565b92915050565b600061095d610958846108f2565b610582565b905080838252602082019050602084028301858111156109805761097f610838565b5b835b818110156109a957806109958882610935565b845260208401935050602081019050610982565b5050509392505050565b600082601f8301126109c8576109c7610507565b5b81356109d884826020860161094a565b91505092915050565b600080604083850312156109f8576109f76104fd565b5b600083013567ffffffffffffffff811115610a1657610a15610502565b5b610a22858286016108c4565b925050602083013567ffffffffffffffff811115610a4357610a42610502565b5b610a4f858286016109b3565b9150509250929050565b610a6281610755565b82525050565b6000604082019050610a7d6000830185610a59565b610a8a6020830184610a59565b9392505050565b610a9a816106ea565b82525050565b6000602082019050610ab56000830184610a91565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610af582610696565b9150610b0083610696565b9250828201905080821115610b1857610b17610abb565b5b92915050565b600081519050919050565b600081905092915050565b60005b83811015610b52578082015181840152602081019050610b37565b60008484015250505050565b6000610b6982610b1e565b610b738185610b29565b9350610b83818560208601610b34565b80840191505092915050565b6000610b9b8284610b5e565b915081905092915050565b600082825260208201905092915050565b6000610bc282610b1e565b610bcc8185610ba6565b9350610bdc818560208601610b34565b610be581610511565b840191505092915050565b60006060820190508181036000830152610c0a8186610bb7565b9050610c196020830185610a59565b610c266040830184610a59565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610c9782610696565b9150610ca283610696565b925082610cb257610cb1610c5d565b5b828206905092915050565b6000610cc882610755565b91506fffffffffffffffffffffffffffffffff8203610cea57610ce9610abb565b5b60018201905091905056fea264697066735822122096e5eb13f0172bb20b069f03fe9ffde7076e5202691b7e7d22386add3a6144ba64736f6c63430008130033",
}

// DiaOracleV2MultiupdateServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use DiaOracleV2MultiupdateServiceMetaData.ABI instead.
var DiaOracleV2MultiupdateServiceABI = DiaOracleV2MultiupdateServiceMetaData.ABI

// DiaOracleV2MultiupdateServiceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DiaOracleV2MultiupdateServiceMetaData.Bin instead.
var DiaOracleV2MultiupdateServiceBin = DiaOracleV2MultiupdateServiceMetaData.Bin

// DeployDiaOracleV2MultiupdateService deploys a new Ethereum contract, binding an instance of DiaOracleV2MultiupdateService to it.
func DeployDiaOracleV2MultiupdateService(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DiaOracleV2MultiupdateService, error) {
	parsed, err := DiaOracleV2MultiupdateServiceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DiaOracleV2MultiupdateServiceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DiaOracleV2MultiupdateService{DiaOracleV2MultiupdateServiceCaller: DiaOracleV2MultiupdateServiceCaller{contract: contract}, DiaOracleV2MultiupdateServiceTransactor: DiaOracleV2MultiupdateServiceTransactor{contract: contract}, DiaOracleV2MultiupdateServiceFilterer: DiaOracleV2MultiupdateServiceFilterer{contract: contract}}, nil
}

// DiaOracleV2MultiupdateService is an auto generated Go binding around an Ethereum contract.
type DiaOracleV2MultiupdateService struct {
	DiaOracleV2MultiupdateServiceCaller     // Read-only binding to the contract
	DiaOracleV2MultiupdateServiceTransactor // Write-only binding to the contract
	DiaOracleV2MultiupdateServiceFilterer   // Log filterer for contract events
}

// DiaOracleV2MultiupdateServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiaOracleV2MultiupdateServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiaOracleV2MultiupdateServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiaOracleV2MultiupdateServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiaOracleV2MultiupdateServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiaOracleV2MultiupdateServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiaOracleV2MultiupdateServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiaOracleV2MultiupdateServiceSession struct {
	Contract     *DiaOracleV2MultiupdateService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// DiaOracleV2MultiupdateServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiaOracleV2MultiupdateServiceCallerSession struct {
	Contract *DiaOracleV2MultiupdateServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// DiaOracleV2MultiupdateServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiaOracleV2MultiupdateServiceTransactorSession struct {
	Contract     *DiaOracleV2MultiupdateServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// DiaOracleV2MultiupdateServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiaOracleV2MultiupdateServiceRaw struct {
	Contract *DiaOracleV2MultiupdateService // Generic contract binding to access the raw methods on
}

// DiaOracleV2MultiupdateServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiaOracleV2MultiupdateServiceCallerRaw struct {
	Contract *DiaOracleV2MultiupdateServiceCaller // Generic read-only contract binding to access the raw methods on
}

// DiaOracleV2MultiupdateServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiaOracleV2MultiupdateServiceTransactorRaw struct {
	Contract *DiaOracleV2MultiupdateServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiaOracleV2MultiupdateService creates a new instance of DiaOracleV2MultiupdateService, bound to a specific deployed contract.
func NewDiaOracleV2MultiupdateService(address common.Address, backend bind.ContractBackend) (*DiaOracleV2MultiupdateService, error) {
	contract, err := bindDiaOracleV2MultiupdateService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiaOracleV2MultiupdateService{DiaOracleV2MultiupdateServiceCaller: DiaOracleV2MultiupdateServiceCaller{contract: contract}, DiaOracleV2MultiupdateServiceTransactor: DiaOracleV2MultiupdateServiceTransactor{contract: contract}, DiaOracleV2MultiupdateServiceFilterer: DiaOracleV2MultiupdateServiceFilterer{contract: contract}}, nil
}

// NewDiaOracleV2MultiupdateServiceCaller creates a new read-only instance of DiaOracleV2MultiupdateService, bound to a specific deployed contract.
func NewDiaOracleV2MultiupdateServiceCaller(address common.Address, caller bind.ContractCaller) (*DiaOracleV2MultiupdateServiceCaller, error) {
	contract, err := bindDiaOracleV2MultiupdateService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiaOracleV2MultiupdateServiceCaller{contract: contract}, nil
}

// NewDiaOracleV2MultiupdateServiceTransactor creates a new write-only instance of DiaOracleV2MultiupdateService, bound to a specific deployed contract.
func NewDiaOracleV2MultiupdateServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*DiaOracleV2MultiupdateServiceTransactor, error) {
	contract, err := bindDiaOracleV2MultiupdateService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiaOracleV2MultiupdateServiceTransactor{contract: contract}, nil
}

// NewDiaOracleV2MultiupdateServiceFilterer creates a new log filterer instance of DiaOracleV2MultiupdateService, bound to a specific deployed contract.
func NewDiaOracleV2MultiupdateServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*DiaOracleV2MultiupdateServiceFilterer, error) {
	contract, err := bindDiaOracleV2MultiupdateService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiaOracleV2MultiupdateServiceFilterer{contract: contract}, nil
}

// bindDiaOracleV2MultiupdateService binds a generic wrapper to an already deployed contract.
func bindDiaOracleV2MultiupdateService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DiaOracleV2MultiupdateServiceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiaOracleV2MultiupdateService.Contract.DiaOracleV2MultiupdateServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.DiaOracleV2MultiupdateServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.DiaOracleV2MultiupdateServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiaOracleV2MultiupdateService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.contract.Transact(opts, method, params...)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceCaller) GetValue(opts *bind.CallOpts, key string) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _DiaOracleV2MultiupdateService.contract.Call(opts, &out, "getValue", key)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceSession) GetValue(key string) (*big.Int, *big.Int, error) {
	return _DiaOracleV2MultiupdateService.Contract.GetValue(&_DiaOracleV2MultiupdateService.CallOpts, key)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceCallerSession) GetValue(key string) (*big.Int, *big.Int, error) {
	return _DiaOracleV2MultiupdateService.Contract.GetValue(&_DiaOracleV2MultiupdateService.CallOpts, key)
}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceCaller) Values(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _DiaOracleV2MultiupdateService.contract.Call(opts, &out, "values", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceSession) Values(arg0 string) (*big.Int, error) {
	return _DiaOracleV2MultiupdateService.Contract.Values(&_DiaOracleV2MultiupdateService.CallOpts, arg0)
}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceCallerSession) Values(arg0 string) (*big.Int, error) {
	return _DiaOracleV2MultiupdateService.Contract.Values(&_DiaOracleV2MultiupdateService.CallOpts, arg0)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x8d241526.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] compressedValues) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactor) SetMultipleValues(opts *bind.TransactOpts, keys []string, compressedValues []*big.Int) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.contract.Transact(opts, "setMultipleValues", keys, compressedValues)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x8d241526.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] compressedValues) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceSession) SetMultipleValues(keys []string, compressedValues []*big.Int) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.SetMultipleValues(&_DiaOracleV2MultiupdateService.TransactOpts, keys, compressedValues)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x8d241526.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] compressedValues) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactorSession) SetMultipleValues(keys []string, compressedValues []*big.Int) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.SetMultipleValues(&_DiaOracleV2MultiupdateService.TransactOpts, keys, compressedValues)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactor) SetValue(opts *bind.TransactOpts, key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.contract.Transact(opts, "setValue", key, value, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceSession) SetValue(key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.SetValue(&_DiaOracleV2MultiupdateService.TransactOpts, key, value, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactorSession) SetValue(key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.SetValue(&_DiaOracleV2MultiupdateService.TransactOpts, key, value, timestamp)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactor) UpdateOracleUpdaterAddress(opts *bind.TransactOpts, newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.contract.Transact(opts, "updateOracleUpdaterAddress", newOracleUpdaterAddress)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceSession) UpdateOracleUpdaterAddress(newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.UpdateOracleUpdaterAddress(&_DiaOracleV2MultiupdateService.TransactOpts, newOracleUpdaterAddress)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceTransactorSession) UpdateOracleUpdaterAddress(newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DiaOracleV2MultiupdateService.Contract.UpdateOracleUpdaterAddress(&_DiaOracleV2MultiupdateService.TransactOpts, newOracleUpdaterAddress)
}

// DiaOracleV2MultiupdateServiceOracleUpdateIterator is returned from FilterOracleUpdate and is used to iterate over the raw logs and unpacked data for OracleUpdate events raised by the DiaOracleV2MultiupdateService contract.
type DiaOracleV2MultiupdateServiceOracleUpdateIterator struct {
	Event *DiaOracleV2MultiupdateServiceOracleUpdate // Event containing the contract specifics and raw log

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
func (it *DiaOracleV2MultiupdateServiceOracleUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiaOracleV2MultiupdateServiceOracleUpdate)
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
		it.Event = new(DiaOracleV2MultiupdateServiceOracleUpdate)
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
func (it *DiaOracleV2MultiupdateServiceOracleUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiaOracleV2MultiupdateServiceOracleUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiaOracleV2MultiupdateServiceOracleUpdate represents a OracleUpdate event raised by the DiaOracleV2MultiupdateService contract.
type DiaOracleV2MultiupdateServiceOracleUpdate struct {
	Key       string
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleUpdate is a free log retrieval operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceFilterer) FilterOracleUpdate(opts *bind.FilterOpts) (*DiaOracleV2MultiupdateServiceOracleUpdateIterator, error) {

	logs, sub, err := _DiaOracleV2MultiupdateService.contract.FilterLogs(opts, "OracleUpdate")
	if err != nil {
		return nil, err
	}
	return &DiaOracleV2MultiupdateServiceOracleUpdateIterator{contract: _DiaOracleV2MultiupdateService.contract, event: "OracleUpdate", logs: logs, sub: sub}, nil
}

// WatchOracleUpdate is a free log subscription operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceFilterer) WatchOracleUpdate(opts *bind.WatchOpts, sink chan<- *DiaOracleV2MultiupdateServiceOracleUpdate) (event.Subscription, error) {

	logs, sub, err := _DiaOracleV2MultiupdateService.contract.WatchLogs(opts, "OracleUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiaOracleV2MultiupdateServiceOracleUpdate)
				if err := _DiaOracleV2MultiupdateService.contract.UnpackLog(event, "OracleUpdate", log); err != nil {
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

// ParseOracleUpdate is a log parse operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceFilterer) ParseOracleUpdate(log types.Log) (*DiaOracleV2MultiupdateServiceOracleUpdate, error) {
	event := new(DiaOracleV2MultiupdateServiceOracleUpdate)
	if err := _DiaOracleV2MultiupdateService.contract.UnpackLog(event, "OracleUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiaOracleV2MultiupdateServiceUpdaterAddressChangeIterator is returned from FilterUpdaterAddressChange and is used to iterate over the raw logs and unpacked data for UpdaterAddressChange events raised by the DiaOracleV2MultiupdateService contract.
type DiaOracleV2MultiupdateServiceUpdaterAddressChangeIterator struct {
	Event *DiaOracleV2MultiupdateServiceUpdaterAddressChange // Event containing the contract specifics and raw log

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
func (it *DiaOracleV2MultiupdateServiceUpdaterAddressChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiaOracleV2MultiupdateServiceUpdaterAddressChange)
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
		it.Event = new(DiaOracleV2MultiupdateServiceUpdaterAddressChange)
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
func (it *DiaOracleV2MultiupdateServiceUpdaterAddressChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiaOracleV2MultiupdateServiceUpdaterAddressChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiaOracleV2MultiupdateServiceUpdaterAddressChange represents a UpdaterAddressChange event raised by the DiaOracleV2MultiupdateService contract.
type DiaOracleV2MultiupdateServiceUpdaterAddressChange struct {
	NewUpdater common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdaterAddressChange is a free log retrieval operation binding the contract event 0x121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f.
//
// Solidity: event UpdaterAddressChange(address newUpdater)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceFilterer) FilterUpdaterAddressChange(opts *bind.FilterOpts) (*DiaOracleV2MultiupdateServiceUpdaterAddressChangeIterator, error) {

	logs, sub, err := _DiaOracleV2MultiupdateService.contract.FilterLogs(opts, "UpdaterAddressChange")
	if err != nil {
		return nil, err
	}
	return &DiaOracleV2MultiupdateServiceUpdaterAddressChangeIterator{contract: _DiaOracleV2MultiupdateService.contract, event: "UpdaterAddressChange", logs: logs, sub: sub}, nil
}

// WatchUpdaterAddressChange is a free log subscription operation binding the contract event 0x121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f.
//
// Solidity: event UpdaterAddressChange(address newUpdater)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceFilterer) WatchUpdaterAddressChange(opts *bind.WatchOpts, sink chan<- *DiaOracleV2MultiupdateServiceUpdaterAddressChange) (event.Subscription, error) {

	logs, sub, err := _DiaOracleV2MultiupdateService.contract.WatchLogs(opts, "UpdaterAddressChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiaOracleV2MultiupdateServiceUpdaterAddressChange)
				if err := _DiaOracleV2MultiupdateService.contract.UnpackLog(event, "UpdaterAddressChange", log); err != nil {
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

// ParseUpdaterAddressChange is a log parse operation binding the contract event 0x121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f.
//
// Solidity: event UpdaterAddressChange(address newUpdater)
func (_DiaOracleV2MultiupdateService *DiaOracleV2MultiupdateServiceFilterer) ParseUpdaterAddressChange(log types.Log) (*DiaOracleV2MultiupdateServiceUpdaterAddressChange, error) {
	event := new(DiaOracleV2MultiupdateServiceUpdaterAddressChange)
	if err := _DiaOracleV2MultiupdateService.contract.UnpackLog(event, "UpdaterAddressChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
