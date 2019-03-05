// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MKRABI is the input ABI used to generate the binding from.
const MKRABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// MKRBin is the compiled bytecode used for deploying new contracts.
const MKRBin = `0x608060405234801561001057600080fd5b50604051610ba1380380610ba18339810160409081528151602080840151838501516060860151336000908152600585529586208590556003859055918601805194969095919492019261006792908601906100ab565b50805161007b9060019060208401906100ab565b50506002805460ff90921660ff19909216919091179055505060048054600160a060020a03191633179055610146565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ec57805160ff1916838001178555610119565b82800160010185558215610119579182015b828111156101195782518255916020019190600101906100fe565b50610125929150610129565b5090565b61014391905b80821115610125576000815560010161012f565b90565b610a4c806101556000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100dc578063095ea7b31461016657806318160ddd1461019e57806323b872dd146101c5578063313ce567146101ef5780633bed33ce1461021a57806342966c68146102325780636623fc461461024a57806370a08231146102625780638da5cb5b1461028357806395d89b41146102b4578063a9059cbb146102c9578063cd4217c1146102ed578063d7a78db81461030e578063dd62ed3e14610326575b005b3480156100e857600080fd5b506100f161034d565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561012b578181015183820152602001610113565b50505050905090810190601f1680156101585780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561017257600080fd5b5061018a600160a060020a03600435166024356103db565b604080519115158252519081900360200190f35b3480156101aa57600080fd5b506101b3610417565b60408051918252519081900360200190f35b3480156101d157600080fd5b5061018a600160a060020a036004358116906024351660443561041d565b3480156101fb57600080fd5b506102046105b8565b6040805160ff9092168252519081900360200190f35b34801561022657600080fd5b506100da6004356105c1565b34801561023e57600080fd5b5061018a600435610616565b34801561025657600080fd5b5061018a6004356106b7565b34801561026e57600080fd5b506101b3600160a060020a0360043516610771565b34801561028f57600080fd5b50610298610783565b60408051600160a060020a039092168252519081900360200190f35b3480156102c057600080fd5b506100f1610792565b3480156102d557600080fd5b506100da600160a060020a03600435166024356107ec565b3480156102f957600080fd5b506101b3600160a060020a03600435166108f0565b34801561031a57600080fd5b5061018a600435610902565b34801561033257600080fd5b506101b3600160a060020a03600435811690602435166109bc565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103d35780601f106103a8576101008083540402835291602001916103d3565b820191906000526020600020905b8154815290600101906020018083116103b657829003601f168201915b505050505081565b60008082116103e957600080fd5b50336000908152600760209081526040808320600160a060020a039590951683529390529190912055600190565b60035481565b6000600160a060020a038316151561043457600080fd5b6000821161044157600080fd5b600160a060020a03841660009081526005602052604090205482111561046657600080fd5b600160a060020a038316600090815260056020526040902054828101101561048d57600080fd5b600160a060020a03841660009081526007602090815260408083203384529091529020548211156104bd57600080fd5b600160a060020a0384166000908152600560205260409020546104e090836109d9565b600160a060020a03808616600090815260056020526040808220939093559085168152205461050f90836109ed565b600160a060020a03808516600090815260056020908152604080832094909455918716815260078252828120338252909152205461054d90836109d9565b600160a060020a03808616600081815260076020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b60025460ff1681565b600454600160a060020a031633146105d857600080fd5b600454604051600160a060020a039091169082156108fc029083906000818181858888f19350505050158015610612573d6000803e3d6000fd5b5050565b3360009081526005602052604081205482111561063257600080fd5b6000821161063f57600080fd5b3360009081526005602052604090205461065990836109d9565b3360009081526005602052604090205560035461067690836109d9565b60035560408051838152905133917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2506001919050565b336000908152600660205260408120548211156106d357600080fd5b600082116106e057600080fd5b336000908152600660205260409020546106fa90836109d9565b3360009081526006602090815260408083209390935560059052205461072090836109ed565b33600081815260056020908152604091829020939093558051858152905191927f2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f92918290030190a2506001919050565b60056020526000908152604090205481565b600454600160a060020a031681565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103d35780601f106103a8576101008083540402835291602001916103d3565b600160a060020a038216151561080157600080fd5b6000811161080e57600080fd5b3360009081526005602052604090205481111561082a57600080fd5b600160a060020a038216600090815260056020526040902054818101101561085157600080fd5b3360009081526005602052604090205461086b90826109d9565b3360009081526005602052604080822092909255600160a060020a0384168152205461089790826109ed565b600160a060020a0383166000818152600560209081526040918290209390935580518481529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b60066020526000908152604090205481565b3360009081526005602052604081205482111561091e57600080fd5b6000821161092b57600080fd5b3360009081526005602052604090205461094590836109d9565b3360009081526005602090815260408083209390935560069052205461096b90836109ed565b33600081815260066020908152604091829020939093558051858152905191927ff97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e092918290030190a2506001919050565b600760209081526000928352604080842090915290825290205481565b60006109e783831115610a11565b50900390565b6000828201610a0a848210801590610a055750838210155b610a11565b9392505050565b801515610a1d57600080fd5b505600a165627a7a7230582058b972f96f1c0ad44b24007fd63c16acd38c3e961d2e960b84a63868a603eea90029`

// DeployMKR deploys a new Ethereum contract, binding an instance of MKR to it.
func DeployMKR(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, decimalUnits uint8, tokenSymbol string) (common.Address, *types.Transaction, *MKR, error) {
	parsed, err := abi.JSON(strings.NewReader(MKRABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MKRBin), backend, initialSupply, tokenName, decimalUnits, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MKR{MKRCaller: MKRCaller{contract: contract}, MKRTransactor: MKRTransactor{contract: contract}, MKRFilterer: MKRFilterer{contract: contract}}, nil
}

// MKR is an auto generated Go binding around an Ethereum contract.
type MKR struct {
	MKRCaller     // Read-only binding to the contract
	MKRTransactor // Write-only binding to the contract
	MKRFilterer   // Log filterer for contract events
}

// MKRCaller is an auto generated read-only Go binding around an Ethereum contract.
type MKRCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MKRTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MKRTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MKRFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MKRFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MKRSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MKRSession struct {
	Contract     *MKR            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MKRCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MKRCallerSession struct {
	Contract *MKRCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MKRTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MKRTransactorSession struct {
	Contract     *MKRTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MKRRaw is an auto generated low-level Go binding around an Ethereum contract.
type MKRRaw struct {
	Contract *MKR // Generic contract binding to access the raw methods on
}

// MKRCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MKRCallerRaw struct {
	Contract *MKRCaller // Generic read-only contract binding to access the raw methods on
}

// MKRTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MKRTransactorRaw struct {
	Contract *MKRTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMKR creates a new instance of MKR, bound to a specific deployed contract.
func NewMKR(address common.Address, backend bind.ContractBackend) (*MKR, error) {
	contract, err := bindMKR(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MKR{MKRCaller: MKRCaller{contract: contract}, MKRTransactor: MKRTransactor{contract: contract}, MKRFilterer: MKRFilterer{contract: contract}}, nil
}

// NewMKRCaller creates a new read-only instance of MKR, bound to a specific deployed contract.
func NewMKRCaller(address common.Address, caller bind.ContractCaller) (*MKRCaller, error) {
	contract, err := bindMKR(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MKRCaller{contract: contract}, nil
}

// NewMKRTransactor creates a new write-only instance of MKR, bound to a specific deployed contract.
func NewMKRTransactor(address common.Address, transactor bind.ContractTransactor) (*MKRTransactor, error) {
	contract, err := bindMKR(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MKRTransactor{contract: contract}, nil
}

// NewMKRFilterer creates a new log filterer instance of MKR, bound to a specific deployed contract.
func NewMKRFilterer(address common.Address, filterer bind.ContractFilterer) (*MKRFilterer, error) {
	contract, err := bindMKR(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MKRFilterer{contract: contract}, nil
}

// bindMKR binds a generic wrapper to an already deployed contract.
func bindMKR(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MKRABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_MKR *MKRRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MKR.Contract.MKRTransactor.contract.Transact(opts, method, params...)
}

func (_MKR *MKRRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MKR.Contract.MKRCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MKR *MKRRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MKR.Contract.MKRTransactor.contract.Transfer(opts)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MKR *MKRCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MKR.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MKR *MKRTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MKR.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MKR *MKRTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MKR.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_MKR *MKRCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MKR.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_MKR *MKRSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MKR.Contract.Allowance(&_MKR.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_MKR *MKRCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MKR.Contract.Allowance(&_MKR.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_MKR *MKRCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MKR.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_MKR *MKRSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MKR.Contract.BalanceOf(&_MKR.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_MKR *MKRCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MKR.Contract.BalanceOf(&_MKR.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MKR *MKRCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MKR.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MKR *MKRSession) Decimals() (uint8, error) {
	return _MKR.Contract.Decimals(&_MKR.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MKR *MKRCallerSession) Decimals() (uint8, error) {
	return _MKR.Contract.Decimals(&_MKR.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MKR *MKRCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MKR.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MKR *MKRSession) Name() (string, error) {
	return _MKR.Contract.Name(&_MKR.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MKR *MKRCallerSession) Name() (string, error) {
	return _MKR.Contract.Name(&_MKR.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MKR *MKRCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MKR.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MKR *MKRSession) Symbol() (string, error) {
	return _MKR.Contract.Symbol(&_MKR.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MKR *MKRCallerSession) Symbol() (string, error) {
	return _MKR.Contract.Symbol(&_MKR.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MKR *MKRCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MKR.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MKR *MKRSession) TotalSupply() (*big.Int, error) {
	return _MKR.Contract.TotalSupply(&_MKR.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MKR *MKRCallerSession) TotalSupply() (*big.Int, error) {
	return _MKR.Contract.TotalSupply(&_MKR.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_MKR *MKRTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MKR.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_MKR *MKRSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MKR.Contract.Approve(&_MKR.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_MKR *MKRTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MKR.Contract.Approve(&_MKR.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_MKR *MKRTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MKR.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_MKR *MKRSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MKR.Contract.Transfer(&_MKR.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_MKR *MKRTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MKR.Contract.Transfer(&_MKR.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_MKR *MKRTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MKR.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_MKR *MKRSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MKR.Contract.TransferFrom(&_MKR.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_MKR *MKRTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MKR.Contract.TransferFrom(&_MKR.TransactOpts, _from, _to, _value)
}

// MKRApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MKR contract.
type MKRApprovalIterator struct {
	Event *MKRApproval // Event containing the contract specifics and raw log

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
func (it *MKRApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MKRApproval)
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
		it.Event = new(MKRApproval)
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
func (it *MKRApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MKRApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MKRApproval represents a Approval event raised by the MKR contract.
type MKRApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MKR *MKRFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MKRApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MKR.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MKRApprovalIterator{contract: _MKR.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MKR *MKRFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MKRApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MKR.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MKRApproval)
				if err := _MKR.contract.UnpackLog(event, "Approval", log); err != nil {
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

// MKRTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MKR contract.
type MKRTransferIterator struct {
	Event *MKRTransfer // Event containing the contract specifics and raw log

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
func (it *MKRTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MKRTransfer)
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
		it.Event = new(MKRTransfer)
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
func (it *MKRTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MKRTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MKRTransfer represents a Transfer event raised by the MKR contract.
type MKRTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MKR *MKRFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MKRTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MKR.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MKRTransferIterator{contract: _MKR.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MKR *MKRFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MKRTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MKR.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MKRTransfer)
				if err := _MKR.contract.UnpackLog(event, "Transfer", log); err != nil {
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
