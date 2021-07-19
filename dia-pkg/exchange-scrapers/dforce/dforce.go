// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dforce

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

// DforceABI is the input ABI used to generate the binding from.
const DforceABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"DToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"DTokenBalance\",\"type\":\"uint256\"}],\"name\":\"DisableDToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"DisableToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"}],\"name\":\"DisableTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"EmergencyStop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"DToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"DTokenBalance\",\"type\":\"uint256\"}],\"name\":\"EnableDToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"EnableToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"}],\"name\":\"EnableTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"SetFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxSwing\",\"type\":\"uint256\"}],\"name\":\"SetMaxSwing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"SetOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputPrice\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"TransferIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"TransferOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"active\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dToken\",\"type\":\"address\"}],\"name\":\"disableDToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"disableOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"disableToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_output\",\"type\":\"address\"}],\"name\":\"disableTrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_changeStateTo\",\"type\":\"bool\"}],\"name\":\"emergencyStop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dToken\",\"type\":\"address\"}],\"name\":\"enableDToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"enableToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_output\",\"type\":\"address\"}],\"name\":\"enableTrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_output\",\"type\":\"address\"}],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_output\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_inputAmount\",\"type\":\"uint256\"}],\"name\":\"getAmountByInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_output\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outputAmount\",\"type\":\"uint256\"}],\"name\":\"getAmountByOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_output\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_inputAmount\",\"type\":\"uint256\"}],\"name\":\"getBestOutputByInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxSwing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"remainingDToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_output\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_input\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_output\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fee\",\"type\":\"uint256[]\"}],\"name\":\"setFeeBatch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSwing\",\"type\":\"uint256\"}],\"name\":\"setMaxSwing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportDToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_output\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_output\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_inputAmount\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_output\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"swapTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_output\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outputAmount\",\"type\":\"uint256\"}],\"name\":\"swapTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokensEnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tradesDisable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferIn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"transferOutALL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner_\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Dforce is an auto generated Go binding around an Ethereum contract.
type Dforce struct {
	DforceCaller     // Read-only binding to the contract
	DforceTransactor // Write-only binding to the contract
	DforceFilterer   // Log filterer for contract events
}

// DforceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DforceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DforceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DforceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DforceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DforceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DforceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DforceSession struct {
	Contract     *Dforce           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DforceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DforceCallerSession struct {
	Contract *DforceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DforceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DforceTransactorSession struct {
	Contract     *DforceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DforceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DforceRaw struct {
	Contract *Dforce // Generic contract binding to access the raw methods on
}

// DforceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DforceCallerRaw struct {
	Contract *DforceCaller // Generic read-only contract binding to access the raw methods on
}

// DforceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DforceTransactorRaw struct {
	Contract *DforceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDforce creates a new instance of Dforce, bound to a specific deployed contract.
func NewDforce(address common.Address, backend bind.ContractBackend) (*Dforce, error) {
	contract, err := bindDforce(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dforce{DforceCaller: DforceCaller{contract: contract}, DforceTransactor: DforceTransactor{contract: contract}, DforceFilterer: DforceFilterer{contract: contract}}, nil
}

// NewDforceCaller creates a new read-only instance of Dforce, bound to a specific deployed contract.
func NewDforceCaller(address common.Address, caller bind.ContractCaller) (*DforceCaller, error) {
	contract, err := bindDforce(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DforceCaller{contract: contract}, nil
}

// NewDforceTransactor creates a new write-only instance of Dforce, bound to a specific deployed contract.
func NewDforceTransactor(address common.Address, transactor bind.ContractTransactor) (*DforceTransactor, error) {
	contract, err := bindDforce(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DforceTransactor{contract: contract}, nil
}

// NewDforceFilterer creates a new log filterer instance of Dforce, bound to a specific deployed contract.
func NewDforceFilterer(address common.Address, filterer bind.ContractFilterer) (*DforceFilterer, error) {
	contract, err := bindDforce(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DforceFilterer{contract: contract}, nil
}

// bindDforce binds a generic wrapper to an already deployed contract.
func bindDforce(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DforceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dforce *DforceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dforce.Contract.DforceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dforce *DforceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dforce.Contract.DforceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dforce *DforceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dforce.Contract.DforceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dforce *DforceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dforce.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dforce *DforceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dforce.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dforce *DforceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dforce.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_Dforce *DforceCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_Dforce *DforceSession) Authority() (common.Address, error) {
	return _Dforce.Contract.Authority(&_Dforce.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_Dforce *DforceCallerSession) Authority() (common.Address, error) {
	return _Dforce.Contract.Authority(&_Dforce.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0xccb0101b.
//
// Solidity: function exchangeRate(address _input, address _output) view returns(uint256)
func (_Dforce *DforceCaller) ExchangeRate(opts *bind.CallOpts, _input common.Address, _output common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "exchangeRate", _input, _output)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRate is a free data retrieval call binding the contract method 0xccb0101b.
//
// Solidity: function exchangeRate(address _input, address _output) view returns(uint256)
func (_Dforce *DforceSession) ExchangeRate(_input common.Address, _output common.Address) (*big.Int, error) {
	return _Dforce.Contract.ExchangeRate(&_Dforce.CallOpts, _input, _output)
}

// ExchangeRate is a free data retrieval call binding the contract method 0xccb0101b.
//
// Solidity: function exchangeRate(address _input, address _output) view returns(uint256)
func (_Dforce *DforceCallerSession) ExchangeRate(_input common.Address, _output common.Address) (*big.Int, error) {
	return _Dforce.Contract.ExchangeRate(&_Dforce.CallOpts, _input, _output)
}

// Fee is a free data retrieval call binding the contract method 0xbe5011a1.
//
// Solidity: function fee(address , address ) view returns(uint256)
func (_Dforce *DforceCaller) Fee(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "fee", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xbe5011a1.
//
// Solidity: function fee(address , address ) view returns(uint256)
func (_Dforce *DforceSession) Fee(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Dforce.Contract.Fee(&_Dforce.CallOpts, arg0, arg1)
}

// Fee is a free data retrieval call binding the contract method 0xbe5011a1.
//
// Solidity: function fee(address , address ) view returns(uint256)
func (_Dforce *DforceCallerSession) Fee(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Dforce.Contract.Fee(&_Dforce.CallOpts, arg0, arg1)
}

// GetAmountByInput is a free data retrieval call binding the contract method 0xaf77fedb.
//
// Solidity: function getAmountByInput(address _input, address _output, uint256 _inputAmount) view returns(uint256)
func (_Dforce *DforceCaller) GetAmountByInput(opts *bind.CallOpts, _input common.Address, _output common.Address, _inputAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "getAmountByInput", _input, _output, _inputAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountByInput is a free data retrieval call binding the contract method 0xaf77fedb.
//
// Solidity: function getAmountByInput(address _input, address _output, uint256 _inputAmount) view returns(uint256)
func (_Dforce *DforceSession) GetAmountByInput(_input common.Address, _output common.Address, _inputAmount *big.Int) (*big.Int, error) {
	return _Dforce.Contract.GetAmountByInput(&_Dforce.CallOpts, _input, _output, _inputAmount)
}

// GetAmountByInput is a free data retrieval call binding the contract method 0xaf77fedb.
//
// Solidity: function getAmountByInput(address _input, address _output, uint256 _inputAmount) view returns(uint256)
func (_Dforce *DforceCallerSession) GetAmountByInput(_input common.Address, _output common.Address, _inputAmount *big.Int) (*big.Int, error) {
	return _Dforce.Contract.GetAmountByInput(&_Dforce.CallOpts, _input, _output, _inputAmount)
}

// GetAmountByOutput is a free data retrieval call binding the contract method 0xf6cf166f.
//
// Solidity: function getAmountByOutput(address _input, address _output, uint256 _outputAmount) view returns(uint256)
func (_Dforce *DforceCaller) GetAmountByOutput(opts *bind.CallOpts, _input common.Address, _output common.Address, _outputAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "getAmountByOutput", _input, _output, _outputAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountByOutput is a free data retrieval call binding the contract method 0xf6cf166f.
//
// Solidity: function getAmountByOutput(address _input, address _output, uint256 _outputAmount) view returns(uint256)
func (_Dforce *DforceSession) GetAmountByOutput(_input common.Address, _output common.Address, _outputAmount *big.Int) (*big.Int, error) {
	return _Dforce.Contract.GetAmountByOutput(&_Dforce.CallOpts, _input, _output, _outputAmount)
}

// GetAmountByOutput is a free data retrieval call binding the contract method 0xf6cf166f.
//
// Solidity: function getAmountByOutput(address _input, address _output, uint256 _outputAmount) view returns(uint256)
func (_Dforce *DforceCallerSession) GetAmountByOutput(_input common.Address, _output common.Address, _outputAmount *big.Int) (*big.Int, error) {
	return _Dforce.Contract.GetAmountByOutput(&_Dforce.CallOpts, _input, _output, _outputAmount)
}

// GetBestOutputByInput is a free data retrieval call binding the contract method 0x858efd4b.
//
// Solidity: function getBestOutputByInput(address _input, address _output, uint256 _inputAmount) view returns(uint256)
func (_Dforce *DforceCaller) GetBestOutputByInput(opts *bind.CallOpts, _input common.Address, _output common.Address, _inputAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "getBestOutputByInput", _input, _output, _inputAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBestOutputByInput is a free data retrieval call binding the contract method 0x858efd4b.
//
// Solidity: function getBestOutputByInput(address _input, address _output, uint256 _inputAmount) view returns(uint256)
func (_Dforce *DforceSession) GetBestOutputByInput(_input common.Address, _output common.Address, _inputAmount *big.Int) (*big.Int, error) {
	return _Dforce.Contract.GetBestOutputByInput(&_Dforce.CallOpts, _input, _output, _inputAmount)
}

// GetBestOutputByInput is a free data retrieval call binding the contract method 0x858efd4b.
//
// Solidity: function getBestOutputByInput(address _input, address _output, uint256 _inputAmount) view returns(uint256)
func (_Dforce *DforceCallerSession) GetBestOutputByInput(_input common.Address, _output common.Address, _inputAmount *big.Int) (*big.Int, error) {
	return _Dforce.Contract.GetBestOutputByInput(&_Dforce.CallOpts, _input, _output, _inputAmount)
}

// GetLiquidity is a free data retrieval call binding the contract method 0xa747b93b.
//
// Solidity: function getLiquidity(address _token) view returns(uint256)
func (_Dforce *DforceCaller) GetLiquidity(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "getLiquidity", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidity is a free data retrieval call binding the contract method 0xa747b93b.
//
// Solidity: function getLiquidity(address _token) view returns(uint256)
func (_Dforce *DforceSession) GetLiquidity(_token common.Address) (*big.Int, error) {
	return _Dforce.Contract.GetLiquidity(&_Dforce.CallOpts, _token)
}

// GetLiquidity is a free data retrieval call binding the contract method 0xa747b93b.
//
// Solidity: function getLiquidity(address _token) view returns(uint256)
func (_Dforce *DforceCallerSession) GetLiquidity(_token common.Address) (*big.Int, error) {
	return _Dforce.Contract.GetLiquidity(&_Dforce.CallOpts, _token)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_Dforce *DforceCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_Dforce *DforceSession) IsOpen() (bool, error) {
	return _Dforce.Contract.IsOpen(&_Dforce.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_Dforce *DforceCallerSession) IsOpen() (bool, error) {
	return _Dforce.Contract.IsOpen(&_Dforce.CallOpts)
}

// MaxSwing is a free data retrieval call binding the contract method 0xc5faf1d5.
//
// Solidity: function maxSwing() view returns(uint256)
func (_Dforce *DforceCaller) MaxSwing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "maxSwing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSwing is a free data retrieval call binding the contract method 0xc5faf1d5.
//
// Solidity: function maxSwing() view returns(uint256)
func (_Dforce *DforceSession) MaxSwing() (*big.Int, error) {
	return _Dforce.Contract.MaxSwing(&_Dforce.CallOpts)
}

// MaxSwing is a free data retrieval call binding the contract method 0xc5faf1d5.
//
// Solidity: function maxSwing() view returns(uint256)
func (_Dforce *DforceCallerSession) MaxSwing() (*big.Int, error) {
	return _Dforce.Contract.MaxSwing(&_Dforce.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Dforce *DforceCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "newOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Dforce *DforceSession) NewOwner() (common.Address, error) {
	return _Dforce.Contract.NewOwner(&_Dforce.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Dforce *DforceCallerSession) NewOwner() (common.Address, error) {
	return _Dforce.Contract.NewOwner(&_Dforce.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Dforce *DforceCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Dforce *DforceSession) Oracle() (common.Address, error) {
	return _Dforce.Contract.Oracle(&_Dforce.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Dforce *DforceCallerSession) Oracle() (common.Address, error) {
	return _Dforce.Contract.Oracle(&_Dforce.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dforce *DforceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dforce *DforceSession) Owner() (common.Address, error) {
	return _Dforce.Contract.Owner(&_Dforce.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dforce *DforceCallerSession) Owner() (common.Address, error) {
	return _Dforce.Contract.Owner(&_Dforce.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Dforce *DforceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Dforce *DforceSession) Paused() (bool, error) {
	return _Dforce.Contract.Paused(&_Dforce.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Dforce *DforceCallerSession) Paused() (bool, error) {
	return _Dforce.Contract.Paused(&_Dforce.CallOpts)
}

// RemainingDToken is a free data retrieval call binding the contract method 0x3a8921a7.
//
// Solidity: function remainingDToken(address ) view returns(address)
func (_Dforce *DforceCaller) RemainingDToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "remainingDToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemainingDToken is a free data retrieval call binding the contract method 0x3a8921a7.
//
// Solidity: function remainingDToken(address ) view returns(address)
func (_Dforce *DforceSession) RemainingDToken(arg0 common.Address) (common.Address, error) {
	return _Dforce.Contract.RemainingDToken(&_Dforce.CallOpts, arg0)
}

// RemainingDToken is a free data retrieval call binding the contract method 0x3a8921a7.
//
// Solidity: function remainingDToken(address ) view returns(address)
func (_Dforce *DforceCallerSession) RemainingDToken(arg0 common.Address) (common.Address, error) {
	return _Dforce.Contract.RemainingDToken(&_Dforce.CallOpts, arg0)
}

// SupportDToken is a free data retrieval call binding the contract method 0xd6692462.
//
// Solidity: function supportDToken(address ) view returns(address)
func (_Dforce *DforceCaller) SupportDToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "supportDToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportDToken is a free data retrieval call binding the contract method 0xd6692462.
//
// Solidity: function supportDToken(address ) view returns(address)
func (_Dforce *DforceSession) SupportDToken(arg0 common.Address) (common.Address, error) {
	return _Dforce.Contract.SupportDToken(&_Dforce.CallOpts, arg0)
}

// SupportDToken is a free data retrieval call binding the contract method 0xd6692462.
//
// Solidity: function supportDToken(address ) view returns(address)
func (_Dforce *DforceCallerSession) SupportDToken(arg0 common.Address) (common.Address, error) {
	return _Dforce.Contract.SupportDToken(&_Dforce.CallOpts, arg0)
}

// TokensEnable is a free data retrieval call binding the contract method 0x338a6397.
//
// Solidity: function tokensEnable(address ) view returns(bool)
func (_Dforce *DforceCaller) TokensEnable(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "tokensEnable", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokensEnable is a free data retrieval call binding the contract method 0x338a6397.
//
// Solidity: function tokensEnable(address ) view returns(bool)
func (_Dforce *DforceSession) TokensEnable(arg0 common.Address) (bool, error) {
	return _Dforce.Contract.TokensEnable(&_Dforce.CallOpts, arg0)
}

// TokensEnable is a free data retrieval call binding the contract method 0x338a6397.
//
// Solidity: function tokensEnable(address ) view returns(bool)
func (_Dforce *DforceCallerSession) TokensEnable(arg0 common.Address) (bool, error) {
	return _Dforce.Contract.TokensEnable(&_Dforce.CallOpts, arg0)
}

// TradesDisable is a free data retrieval call binding the contract method 0xb1db0a6e.
//
// Solidity: function tradesDisable(address , address ) view returns(bool)
func (_Dforce *DforceCaller) TradesDisable(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Dforce.contract.Call(opts, &out, "tradesDisable", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TradesDisable is a free data retrieval call binding the contract method 0xb1db0a6e.
//
// Solidity: function tradesDisable(address , address ) view returns(bool)
func (_Dforce *DforceSession) TradesDisable(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Dforce.Contract.TradesDisable(&_Dforce.CallOpts, arg0, arg1)
}

// TradesDisable is a free data retrieval call binding the contract method 0xb1db0a6e.
//
// Solidity: function tradesDisable(address , address ) view returns(bool)
func (_Dforce *DforceCallerSession) TradesDisable(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Dforce.Contract.TradesDisable(&_Dforce.CallOpts, arg0, arg1)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Dforce *DforceTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Dforce *DforceSession) AcceptOwnership() (*types.Transaction, error) {
	return _Dforce.Contract.AcceptOwnership(&_Dforce.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Dforce *DforceTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Dforce.Contract.AcceptOwnership(&_Dforce.TransactOpts)
}

// Active is a paid mutator transaction binding the contract method 0x4ea71327.
//
// Solidity: function active(address _oracle) returns()
func (_Dforce *DforceTransactor) Active(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "active", _oracle)
}

// Active is a paid mutator transaction binding the contract method 0x4ea71327.
//
// Solidity: function active(address _oracle) returns()
func (_Dforce *DforceSession) Active(_oracle common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.Active(&_Dforce.TransactOpts, _oracle)
}

// Active is a paid mutator transaction binding the contract method 0x4ea71327.
//
// Solidity: function active(address _oracle) returns()
func (_Dforce *DforceTransactorSession) Active(_oracle common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.Active(&_Dforce.TransactOpts, _oracle)
}

// DisableDToken is a paid mutator transaction binding the contract method 0xea000b80.
//
// Solidity: function disableDToken(address _dToken) returns()
func (_Dforce *DforceTransactor) DisableDToken(opts *bind.TransactOpts, _dToken common.Address) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "disableDToken", _dToken)
}

// DisableDToken is a paid mutator transaction binding the contract method 0xea000b80.
//
// Solidity: function disableDToken(address _dToken) returns()
func (_Dforce *DforceSession) DisableDToken(_dToken common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.DisableDToken(&_Dforce.TransactOpts, _dToken)
}

// DisableDToken is a paid mutator transaction binding the contract method 0xea000b80.
//
// Solidity: function disableDToken(address _dToken) returns()
func (_Dforce *DforceTransactorSession) DisableDToken(_dToken common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.DisableDToken(&_Dforce.TransactOpts, _dToken)
}

// DisableOwnership is a paid mutator transaction binding the contract method 0x93c0b096.
//
// Solidity: function disableOwnership() returns()
func (_Dforce *DforceTransactor) DisableOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "disableOwnership")
}

// DisableOwnership is a paid mutator transaction binding the contract method 0x93c0b096.
//
// Solidity: function disableOwnership() returns()
func (_Dforce *DforceSession) DisableOwnership() (*types.Transaction, error) {
	return _Dforce.Contract.DisableOwnership(&_Dforce.TransactOpts)
}

// DisableOwnership is a paid mutator transaction binding the contract method 0x93c0b096.
//
// Solidity: function disableOwnership() returns()
func (_Dforce *DforceTransactorSession) DisableOwnership() (*types.Transaction, error) {
	return _Dforce.Contract.DisableOwnership(&_Dforce.TransactOpts)
}

// DisableToken is a paid mutator transaction binding the contract method 0x23e27a64.
//
// Solidity: function disableToken(address _token) returns()
func (_Dforce *DforceTransactor) DisableToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "disableToken", _token)
}

// DisableToken is a paid mutator transaction binding the contract method 0x23e27a64.
//
// Solidity: function disableToken(address _token) returns()
func (_Dforce *DforceSession) DisableToken(_token common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.DisableToken(&_Dforce.TransactOpts, _token)
}

// DisableToken is a paid mutator transaction binding the contract method 0x23e27a64.
//
// Solidity: function disableToken(address _token) returns()
func (_Dforce *DforceTransactorSession) DisableToken(_token common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.DisableToken(&_Dforce.TransactOpts, _token)
}

// DisableTrade is a paid mutator transaction binding the contract method 0xb80e706a.
//
// Solidity: function disableTrade(address _input, address _output) returns()
func (_Dforce *DforceTransactor) DisableTrade(opts *bind.TransactOpts, _input common.Address, _output common.Address) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "disableTrade", _input, _output)
}

// DisableTrade is a paid mutator transaction binding the contract method 0xb80e706a.
//
// Solidity: function disableTrade(address _input, address _output) returns()
func (_Dforce *DforceSession) DisableTrade(_input common.Address, _output common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.DisableTrade(&_Dforce.TransactOpts, _input, _output)
}

// DisableTrade is a paid mutator transaction binding the contract method 0xb80e706a.
//
// Solidity: function disableTrade(address _input, address _output) returns()
func (_Dforce *DforceTransactorSession) DisableTrade(_input common.Address, _output common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.DisableTrade(&_Dforce.TransactOpts, _input, _output)
}

// EmergencyStop is a paid mutator transaction binding the contract method 0x0ac62e02.
//
// Solidity: function emergencyStop(bool _changeStateTo) returns()
func (_Dforce *DforceTransactor) EmergencyStop(opts *bind.TransactOpts, _changeStateTo bool) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "emergencyStop", _changeStateTo)
}

// EmergencyStop is a paid mutator transaction binding the contract method 0x0ac62e02.
//
// Solidity: function emergencyStop(bool _changeStateTo) returns()
func (_Dforce *DforceSession) EmergencyStop(_changeStateTo bool) (*types.Transaction, error) {
	return _Dforce.Contract.EmergencyStop(&_Dforce.TransactOpts, _changeStateTo)
}

// EmergencyStop is a paid mutator transaction binding the contract method 0x0ac62e02.
//
// Solidity: function emergencyStop(bool _changeStateTo) returns()
func (_Dforce *DforceTransactorSession) EmergencyStop(_changeStateTo bool) (*types.Transaction, error) {
	return _Dforce.Contract.EmergencyStop(&_Dforce.TransactOpts, _changeStateTo)
}

// EnableDToken is a paid mutator transaction binding the contract method 0x4a67645c.
//
// Solidity: function enableDToken(address _dToken) returns()
func (_Dforce *DforceTransactor) EnableDToken(opts *bind.TransactOpts, _dToken common.Address) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "enableDToken", _dToken)
}

// EnableDToken is a paid mutator transaction binding the contract method 0x4a67645c.
//
// Solidity: function enableDToken(address _dToken) returns()
func (_Dforce *DforceSession) EnableDToken(_dToken common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.EnableDToken(&_Dforce.TransactOpts, _dToken)
}

// EnableDToken is a paid mutator transaction binding the contract method 0x4a67645c.
//
// Solidity: function enableDToken(address _dToken) returns()
func (_Dforce *DforceTransactorSession) EnableDToken(_dToken common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.EnableDToken(&_Dforce.TransactOpts, _dToken)
}

// EnableToken is a paid mutator transaction binding the contract method 0xc690908a.
//
// Solidity: function enableToken(address _token) returns()
func (_Dforce *DforceTransactor) EnableToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "enableToken", _token)
}

// EnableToken is a paid mutator transaction binding the contract method 0xc690908a.
//
// Solidity: function enableToken(address _token) returns()
func (_Dforce *DforceSession) EnableToken(_token common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.EnableToken(&_Dforce.TransactOpts, _token)
}

// EnableToken is a paid mutator transaction binding the contract method 0xc690908a.
//
// Solidity: function enableToken(address _token) returns()
func (_Dforce *DforceTransactorSession) EnableToken(_token common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.EnableToken(&_Dforce.TransactOpts, _token)
}

// EnableTrade is a paid mutator transaction binding the contract method 0x800cf13d.
//
// Solidity: function enableTrade(address _input, address _output) returns()
func (_Dforce *DforceTransactor) EnableTrade(opts *bind.TransactOpts, _input common.Address, _output common.Address) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "enableTrade", _input, _output)
}

// EnableTrade is a paid mutator transaction binding the contract method 0x800cf13d.
//
// Solidity: function enableTrade(address _input, address _output) returns()
func (_Dforce *DforceSession) EnableTrade(_input common.Address, _output common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.EnableTrade(&_Dforce.TransactOpts, _input, _output)
}

// EnableTrade is a paid mutator transaction binding the contract method 0x800cf13d.
//
// Solidity: function enableTrade(address _input, address _output) returns()
func (_Dforce *DforceTransactorSession) EnableTrade(_input common.Address, _output common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.EnableTrade(&_Dforce.TransactOpts, _input, _output)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Dforce *DforceTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Dforce *DforceSession) Pause() (*types.Transaction, error) {
	return _Dforce.Contract.Pause(&_Dforce.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Dforce *DforceTransactorSession) Pause() (*types.Transaction, error) {
	return _Dforce.Contract.Pause(&_Dforce.TransactOpts)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Dforce *DforceTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Dforce *DforceSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.SetAuthority(&_Dforce.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Dforce *DforceTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.SetAuthority(&_Dforce.TransactOpts, authority_)
}

// SetFee is a paid mutator transaction binding the contract method 0xdc7e98df.
//
// Solidity: function setFee(address _input, address _output, uint256 _fee) returns()
func (_Dforce *DforceTransactor) SetFee(opts *bind.TransactOpts, _input common.Address, _output common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "setFee", _input, _output, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0xdc7e98df.
//
// Solidity: function setFee(address _input, address _output, uint256 _fee) returns()
func (_Dforce *DforceSession) SetFee(_input common.Address, _output common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _Dforce.Contract.SetFee(&_Dforce.TransactOpts, _input, _output, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0xdc7e98df.
//
// Solidity: function setFee(address _input, address _output, uint256 _fee) returns()
func (_Dforce *DforceTransactorSession) SetFee(_input common.Address, _output common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _Dforce.Contract.SetFee(&_Dforce.TransactOpts, _input, _output, _fee)
}

// SetFeeBatch is a paid mutator transaction binding the contract method 0x8ff76519.
//
// Solidity: function setFeeBatch(address[] _input, address[] _output, uint256[] _fee) returns()
func (_Dforce *DforceTransactor) SetFeeBatch(opts *bind.TransactOpts, _input []common.Address, _output []common.Address, _fee []*big.Int) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "setFeeBatch", _input, _output, _fee)
}

// SetFeeBatch is a paid mutator transaction binding the contract method 0x8ff76519.
//
// Solidity: function setFeeBatch(address[] _input, address[] _output, uint256[] _fee) returns()
func (_Dforce *DforceSession) SetFeeBatch(_input []common.Address, _output []common.Address, _fee []*big.Int) (*types.Transaction, error) {
	return _Dforce.Contract.SetFeeBatch(&_Dforce.TransactOpts, _input, _output, _fee)
}

// SetFeeBatch is a paid mutator transaction binding the contract method 0x8ff76519.
//
// Solidity: function setFeeBatch(address[] _input, address[] _output, uint256[] _fee) returns()
func (_Dforce *DforceTransactorSession) SetFeeBatch(_input []common.Address, _output []common.Address, _fee []*big.Int) (*types.Transaction, error) {
	return _Dforce.Contract.SetFeeBatch(&_Dforce.TransactOpts, _input, _output, _fee)
}

// SetMaxSwing is a paid mutator transaction binding the contract method 0x3e36b146.
//
// Solidity: function setMaxSwing(uint256 _maxSwing) returns()
func (_Dforce *DforceTransactor) SetMaxSwing(opts *bind.TransactOpts, _maxSwing *big.Int) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "setMaxSwing", _maxSwing)
}

// SetMaxSwing is a paid mutator transaction binding the contract method 0x3e36b146.
//
// Solidity: function setMaxSwing(uint256 _maxSwing) returns()
func (_Dforce *DforceSession) SetMaxSwing(_maxSwing *big.Int) (*types.Transaction, error) {
	return _Dforce.Contract.SetMaxSwing(&_Dforce.TransactOpts, _maxSwing)
}

// SetMaxSwing is a paid mutator transaction binding the contract method 0x3e36b146.
//
// Solidity: function setMaxSwing(uint256 _maxSwing) returns()
func (_Dforce *DforceTransactorSession) SetMaxSwing(_maxSwing *big.Int) (*types.Transaction, error) {
	return _Dforce.Contract.SetMaxSwing(&_Dforce.TransactOpts, _maxSwing)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Dforce *DforceTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Dforce *DforceSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.SetOracle(&_Dforce.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Dforce *DforceTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.SetOracle(&_Dforce.TransactOpts, _oracle)
}

// Swap is a paid mutator transaction binding the contract method 0x6e81221c.
//
// Solidity: function swap(address _input, address _output, uint256 _inputAmount, address _receiver) returns()
func (_Dforce *DforceTransactor) Swap(opts *bind.TransactOpts, _input common.Address, _output common.Address, _inputAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "swap", _input, _output, _inputAmount, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x6e81221c.
//
// Solidity: function swap(address _input, address _output, uint256 _inputAmount, address _receiver) returns()
func (_Dforce *DforceSession) Swap(_input common.Address, _output common.Address, _inputAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.Swap(&_Dforce.TransactOpts, _input, _output, _inputAmount, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x6e81221c.
//
// Solidity: function swap(address _input, address _output, uint256 _inputAmount, address _receiver) returns()
func (_Dforce *DforceTransactorSession) Swap(_input common.Address, _output common.Address, _inputAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.Swap(&_Dforce.TransactOpts, _input, _output, _inputAmount, _receiver)
}

// Swap0 is a paid mutator transaction binding the contract method 0xdf791e50.
//
// Solidity: function swap(address _input, address _output, uint256 _inputAmount) returns()
func (_Dforce *DforceTransactor) Swap0(opts *bind.TransactOpts, _input common.Address, _output common.Address, _inputAmount *big.Int) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "swap0", _input, _output, _inputAmount)
}

// Swap0 is a paid mutator transaction binding the contract method 0xdf791e50.
//
// Solidity: function swap(address _input, address _output, uint256 _inputAmount) returns()
func (_Dforce *DforceSession) Swap0(_input common.Address, _output common.Address, _inputAmount *big.Int) (*types.Transaction, error) {
	return _Dforce.Contract.Swap0(&_Dforce.TransactOpts, _input, _output, _inputAmount)
}

// Swap0 is a paid mutator transaction binding the contract method 0xdf791e50.
//
// Solidity: function swap(address _input, address _output, uint256 _inputAmount) returns()
func (_Dforce *DforceTransactorSession) Swap0(_input common.Address, _output common.Address, _inputAmount *big.Int) (*types.Transaction, error) {
	return _Dforce.Contract.Swap0(&_Dforce.TransactOpts, _input, _output, _inputAmount)
}

// SwapTo is a paid mutator transaction binding the contract method 0x2f7846a0.
//
// Solidity: function swapTo(address _input, address _output, uint256 _outputAmount, address _receiver) returns()
func (_Dforce *DforceTransactor) SwapTo(opts *bind.TransactOpts, _input common.Address, _output common.Address, _outputAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "swapTo", _input, _output, _outputAmount, _receiver)
}

// SwapTo is a paid mutator transaction binding the contract method 0x2f7846a0.
//
// Solidity: function swapTo(address _input, address _output, uint256 _outputAmount, address _receiver) returns()
func (_Dforce *DforceSession) SwapTo(_input common.Address, _output common.Address, _outputAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.SwapTo(&_Dforce.TransactOpts, _input, _output, _outputAmount, _receiver)
}

// SwapTo is a paid mutator transaction binding the contract method 0x2f7846a0.
//
// Solidity: function swapTo(address _input, address _output, uint256 _outputAmount, address _receiver) returns()
func (_Dforce *DforceTransactorSession) SwapTo(_input common.Address, _output common.Address, _outputAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.SwapTo(&_Dforce.TransactOpts, _input, _output, _outputAmount, _receiver)
}

// SwapTo0 is a paid mutator transaction binding the contract method 0x95ff216f.
//
// Solidity: function swapTo(address _input, address _output, uint256 _outputAmount) returns()
func (_Dforce *DforceTransactor) SwapTo0(opts *bind.TransactOpts, _input common.Address, _output common.Address, _outputAmount *big.Int) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "swapTo0", _input, _output, _outputAmount)
}

// SwapTo0 is a paid mutator transaction binding the contract method 0x95ff216f.
//
// Solidity: function swapTo(address _input, address _output, uint256 _outputAmount) returns()
func (_Dforce *DforceSession) SwapTo0(_input common.Address, _output common.Address, _outputAmount *big.Int) (*types.Transaction, error) {
	return _Dforce.Contract.SwapTo0(&_Dforce.TransactOpts, _input, _output, _outputAmount)
}

// SwapTo0 is a paid mutator transaction binding the contract method 0x95ff216f.
//
// Solidity: function swapTo(address _input, address _output, uint256 _outputAmount) returns()
func (_Dforce *DforceTransactorSession) SwapTo0(_input common.Address, _output common.Address, _outputAmount *big.Int) (*types.Transaction, error) {
	return _Dforce.Contract.SwapTo0(&_Dforce.TransactOpts, _input, _output, _outputAmount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe8888915.
//
// Solidity: function transferIn(address _token, uint256 _amount) returns()
func (_Dforce *DforceTransactor) TransferIn(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "transferIn", _token, _amount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe8888915.
//
// Solidity: function transferIn(address _token, uint256 _amount) returns()
func (_Dforce *DforceSession) TransferIn(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Dforce.Contract.TransferIn(&_Dforce.TransactOpts, _token, _amount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe8888915.
//
// Solidity: function transferIn(address _token, uint256 _amount) returns()
func (_Dforce *DforceTransactorSession) TransferIn(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Dforce.Contract.TransferIn(&_Dforce.TransactOpts, _token, _amount)
}

// TransferOut is a paid mutator transaction binding the contract method 0x078d3b79.
//
// Solidity: function transferOut(address _token, address _receiver, uint256 _amount) returns()
func (_Dforce *DforceTransactor) TransferOut(opts *bind.TransactOpts, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "transferOut", _token, _receiver, _amount)
}

// TransferOut is a paid mutator transaction binding the contract method 0x078d3b79.
//
// Solidity: function transferOut(address _token, address _receiver, uint256 _amount) returns()
func (_Dforce *DforceSession) TransferOut(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Dforce.Contract.TransferOut(&_Dforce.TransactOpts, _token, _receiver, _amount)
}

// TransferOut is a paid mutator transaction binding the contract method 0x078d3b79.
//
// Solidity: function transferOut(address _token, address _receiver, uint256 _amount) returns()
func (_Dforce *DforceTransactorSession) TransferOut(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Dforce.Contract.TransferOut(&_Dforce.TransactOpts, _token, _receiver, _amount)
}

// TransferOutALL is a paid mutator transaction binding the contract method 0x381dc1e5.
//
// Solidity: function transferOutALL(address _token, address _receiver) returns()
func (_Dforce *DforceTransactor) TransferOutALL(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "transferOutALL", _token, _receiver)
}

// TransferOutALL is a paid mutator transaction binding the contract method 0x381dc1e5.
//
// Solidity: function transferOutALL(address _token, address _receiver) returns()
func (_Dforce *DforceSession) TransferOutALL(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.TransferOutALL(&_Dforce.TransactOpts, _token, _receiver)
}

// TransferOutALL is a paid mutator transaction binding the contract method 0x381dc1e5.
//
// Solidity: function transferOutALL(address _token, address _receiver) returns()
func (_Dforce *DforceTransactorSession) TransferOutALL(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.TransferOutALL(&_Dforce.TransactOpts, _token, _receiver)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner_) returns()
func (_Dforce *DforceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner_ common.Address) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "transferOwnership", newOwner_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner_) returns()
func (_Dforce *DforceSession) TransferOwnership(newOwner_ common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.TransferOwnership(&_Dforce.TransactOpts, newOwner_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner_) returns()
func (_Dforce *DforceTransactorSession) TransferOwnership(newOwner_ common.Address) (*types.Transaction, error) {
	return _Dforce.Contract.TransferOwnership(&_Dforce.TransactOpts, newOwner_)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Dforce *DforceTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dforce.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Dforce *DforceSession) Unpause() (*types.Transaction, error) {
	return _Dforce.Contract.Unpause(&_Dforce.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Dforce *DforceTransactorSession) Unpause() (*types.Transaction, error) {
	return _Dforce.Contract.Unpause(&_Dforce.TransactOpts)
}

// DforceDisableDTokenIterator is returned from FilterDisableDToken and is used to iterate over the raw logs and unpacked data for DisableDToken events raised by the Dforce contract.
type DforceDisableDTokenIterator struct {
	Event *DforceDisableDToken // Event containing the contract specifics and raw log

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
func (it *DforceDisableDTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceDisableDToken)
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
		it.Event = new(DforceDisableDToken)
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
func (it *DforceDisableDTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceDisableDTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceDisableDToken represents a DisableDToken event raised by the Dforce contract.
type DforceDisableDToken struct {
	DToken        common.Address
	Token         common.Address
	DTokenBalance *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDisableDToken is a free log retrieval operation binding the contract event 0xaf142643cc04854b375c86a81031bf9043111c3757c7e489c0e791dd30fe5252.
//
// Solidity: event DisableDToken(address DToken, address token, uint256 DTokenBalance)
func (_Dforce *DforceFilterer) FilterDisableDToken(opts *bind.FilterOpts) (*DforceDisableDTokenIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "DisableDToken")
	if err != nil {
		return nil, err
	}
	return &DforceDisableDTokenIterator{contract: _Dforce.contract, event: "DisableDToken", logs: logs, sub: sub}, nil
}

// WatchDisableDToken is a free log subscription operation binding the contract event 0xaf142643cc04854b375c86a81031bf9043111c3757c7e489c0e791dd30fe5252.
//
// Solidity: event DisableDToken(address DToken, address token, uint256 DTokenBalance)
func (_Dforce *DforceFilterer) WatchDisableDToken(opts *bind.WatchOpts, sink chan<- *DforceDisableDToken) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "DisableDToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceDisableDToken)
				if err := _Dforce.contract.UnpackLog(event, "DisableDToken", log); err != nil {
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

// ParseDisableDToken is a log parse operation binding the contract event 0xaf142643cc04854b375c86a81031bf9043111c3757c7e489c0e791dd30fe5252.
//
// Solidity: event DisableDToken(address DToken, address token, uint256 DTokenBalance)
func (_Dforce *DforceFilterer) ParseDisableDToken(log types.Log) (*DforceDisableDToken, error) {
	event := new(DforceDisableDToken)
	if err := _Dforce.contract.UnpackLog(event, "DisableDToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceDisableTokenIterator is returned from FilterDisableToken and is used to iterate over the raw logs and unpacked data for DisableToken events raised by the Dforce contract.
type DforceDisableTokenIterator struct {
	Event *DforceDisableToken // Event containing the contract specifics and raw log

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
func (it *DforceDisableTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceDisableToken)
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
		it.Event = new(DforceDisableToken)
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
func (it *DforceDisableTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceDisableTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceDisableToken represents a DisableToken event raised by the Dforce contract.
type DforceDisableToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDisableToken is a free log retrieval operation binding the contract event 0x3907c2f289c906800bbef774e6e4a476623eae8c1abc6cd0a78ef88aba9a6225.
//
// Solidity: event DisableToken(address token)
func (_Dforce *DforceFilterer) FilterDisableToken(opts *bind.FilterOpts) (*DforceDisableTokenIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "DisableToken")
	if err != nil {
		return nil, err
	}
	return &DforceDisableTokenIterator{contract: _Dforce.contract, event: "DisableToken", logs: logs, sub: sub}, nil
}

// WatchDisableToken is a free log subscription operation binding the contract event 0x3907c2f289c906800bbef774e6e4a476623eae8c1abc6cd0a78ef88aba9a6225.
//
// Solidity: event DisableToken(address token)
func (_Dforce *DforceFilterer) WatchDisableToken(opts *bind.WatchOpts, sink chan<- *DforceDisableToken) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "DisableToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceDisableToken)
				if err := _Dforce.contract.UnpackLog(event, "DisableToken", log); err != nil {
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

// ParseDisableToken is a log parse operation binding the contract event 0x3907c2f289c906800bbef774e6e4a476623eae8c1abc6cd0a78ef88aba9a6225.
//
// Solidity: event DisableToken(address token)
func (_Dforce *DforceFilterer) ParseDisableToken(log types.Log) (*DforceDisableToken, error) {
	event := new(DforceDisableToken)
	if err := _Dforce.contract.UnpackLog(event, "DisableToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceDisableTradeIterator is returned from FilterDisableTrade and is used to iterate over the raw logs and unpacked data for DisableTrade events raised by the Dforce contract.
type DforceDisableTradeIterator struct {
	Event *DforceDisableTrade // Event containing the contract specifics and raw log

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
func (it *DforceDisableTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceDisableTrade)
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
		it.Event = new(DforceDisableTrade)
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
func (it *DforceDisableTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceDisableTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceDisableTrade represents a DisableTrade event raised by the Dforce contract.
type DforceDisableTrade struct {
	Input  common.Address
	Output common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisableTrade is a free log retrieval operation binding the contract event 0x48026cd74df281df76fcaf355f758047053aba10579ac0251255ea125d28860b.
//
// Solidity: event DisableTrade(address input, address output)
func (_Dforce *DforceFilterer) FilterDisableTrade(opts *bind.FilterOpts) (*DforceDisableTradeIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "DisableTrade")
	if err != nil {
		return nil, err
	}
	return &DforceDisableTradeIterator{contract: _Dforce.contract, event: "DisableTrade", logs: logs, sub: sub}, nil
}

// WatchDisableTrade is a free log subscription operation binding the contract event 0x48026cd74df281df76fcaf355f758047053aba10579ac0251255ea125d28860b.
//
// Solidity: event DisableTrade(address input, address output)
func (_Dforce *DforceFilterer) WatchDisableTrade(opts *bind.WatchOpts, sink chan<- *DforceDisableTrade) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "DisableTrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceDisableTrade)
				if err := _Dforce.contract.UnpackLog(event, "DisableTrade", log); err != nil {
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

// ParseDisableTrade is a log parse operation binding the contract event 0x48026cd74df281df76fcaf355f758047053aba10579ac0251255ea125d28860b.
//
// Solidity: event DisableTrade(address input, address output)
func (_Dforce *DforceFilterer) ParseDisableTrade(log types.Log) (*DforceDisableTrade, error) {
	event := new(DforceDisableTrade)
	if err := _Dforce.contract.UnpackLog(event, "DisableTrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceEmergencyStopIterator is returned from FilterEmergencyStop and is used to iterate over the raw logs and unpacked data for EmergencyStop events raised by the Dforce contract.
type DforceEmergencyStopIterator struct {
	Event *DforceEmergencyStop // Event containing the contract specifics and raw log

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
func (it *DforceEmergencyStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceEmergencyStop)
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
		it.Event = new(DforceEmergencyStop)
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
func (it *DforceEmergencyStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceEmergencyStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceEmergencyStop represents a EmergencyStop event raised by the Dforce contract.
type DforceEmergencyStop struct {
	IsOpen bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStop is a free log retrieval operation binding the contract event 0xcbe1b789cda412d9e1c647ed03d0c71e2f71484be36f9ca2b2a346b29edf1b70.
//
// Solidity: event EmergencyStop(bool isOpen)
func (_Dforce *DforceFilterer) FilterEmergencyStop(opts *bind.FilterOpts) (*DforceEmergencyStopIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "EmergencyStop")
	if err != nil {
		return nil, err
	}
	return &DforceEmergencyStopIterator{contract: _Dforce.contract, event: "EmergencyStop", logs: logs, sub: sub}, nil
}

// WatchEmergencyStop is a free log subscription operation binding the contract event 0xcbe1b789cda412d9e1c647ed03d0c71e2f71484be36f9ca2b2a346b29edf1b70.
//
// Solidity: event EmergencyStop(bool isOpen)
func (_Dforce *DforceFilterer) WatchEmergencyStop(opts *bind.WatchOpts, sink chan<- *DforceEmergencyStop) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "EmergencyStop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceEmergencyStop)
				if err := _Dforce.contract.UnpackLog(event, "EmergencyStop", log); err != nil {
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

// ParseEmergencyStop is a log parse operation binding the contract event 0xcbe1b789cda412d9e1c647ed03d0c71e2f71484be36f9ca2b2a346b29edf1b70.
//
// Solidity: event EmergencyStop(bool isOpen)
func (_Dforce *DforceFilterer) ParseEmergencyStop(log types.Log) (*DforceEmergencyStop, error) {
	event := new(DforceEmergencyStop)
	if err := _Dforce.contract.UnpackLog(event, "EmergencyStop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceEnableDTokenIterator is returned from FilterEnableDToken and is used to iterate over the raw logs and unpacked data for EnableDToken events raised by the Dforce contract.
type DforceEnableDTokenIterator struct {
	Event *DforceEnableDToken // Event containing the contract specifics and raw log

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
func (it *DforceEnableDTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceEnableDToken)
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
		it.Event = new(DforceEnableDToken)
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
func (it *DforceEnableDTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceEnableDTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceEnableDToken represents a EnableDToken event raised by the Dforce contract.
type DforceEnableDToken struct {
	DToken        common.Address
	Token         common.Address
	DTokenBalance *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEnableDToken is a free log retrieval operation binding the contract event 0xb323e69a09394d23d413ed3452a32bfcb64cd9eaf2d5fc27e6a98b3ead0bd501.
//
// Solidity: event EnableDToken(address DToken, address token, uint256 DTokenBalance)
func (_Dforce *DforceFilterer) FilterEnableDToken(opts *bind.FilterOpts) (*DforceEnableDTokenIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "EnableDToken")
	if err != nil {
		return nil, err
	}
	return &DforceEnableDTokenIterator{contract: _Dforce.contract, event: "EnableDToken", logs: logs, sub: sub}, nil
}

// WatchEnableDToken is a free log subscription operation binding the contract event 0xb323e69a09394d23d413ed3452a32bfcb64cd9eaf2d5fc27e6a98b3ead0bd501.
//
// Solidity: event EnableDToken(address DToken, address token, uint256 DTokenBalance)
func (_Dforce *DforceFilterer) WatchEnableDToken(opts *bind.WatchOpts, sink chan<- *DforceEnableDToken) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "EnableDToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceEnableDToken)
				if err := _Dforce.contract.UnpackLog(event, "EnableDToken", log); err != nil {
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

// ParseEnableDToken is a log parse operation binding the contract event 0xb323e69a09394d23d413ed3452a32bfcb64cd9eaf2d5fc27e6a98b3ead0bd501.
//
// Solidity: event EnableDToken(address DToken, address token, uint256 DTokenBalance)
func (_Dforce *DforceFilterer) ParseEnableDToken(log types.Log) (*DforceEnableDToken, error) {
	event := new(DforceEnableDToken)
	if err := _Dforce.contract.UnpackLog(event, "EnableDToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceEnableTokenIterator is returned from FilterEnableToken and is used to iterate over the raw logs and unpacked data for EnableToken events raised by the Dforce contract.
type DforceEnableTokenIterator struct {
	Event *DforceEnableToken // Event containing the contract specifics and raw log

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
func (it *DforceEnableTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceEnableToken)
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
		it.Event = new(DforceEnableToken)
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
func (it *DforceEnableTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceEnableTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceEnableToken represents a EnableToken event raised by the Dforce contract.
type DforceEnableToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEnableToken is a free log retrieval operation binding the contract event 0x1f4b6598cc34e310cc34bcfcb372afb5ba5af654fad97b26cc10ae8289f1c62f.
//
// Solidity: event EnableToken(address token)
func (_Dforce *DforceFilterer) FilterEnableToken(opts *bind.FilterOpts) (*DforceEnableTokenIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "EnableToken")
	if err != nil {
		return nil, err
	}
	return &DforceEnableTokenIterator{contract: _Dforce.contract, event: "EnableToken", logs: logs, sub: sub}, nil
}

// WatchEnableToken is a free log subscription operation binding the contract event 0x1f4b6598cc34e310cc34bcfcb372afb5ba5af654fad97b26cc10ae8289f1c62f.
//
// Solidity: event EnableToken(address token)
func (_Dforce *DforceFilterer) WatchEnableToken(opts *bind.WatchOpts, sink chan<- *DforceEnableToken) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "EnableToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceEnableToken)
				if err := _Dforce.contract.UnpackLog(event, "EnableToken", log); err != nil {
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

// ParseEnableToken is a log parse operation binding the contract event 0x1f4b6598cc34e310cc34bcfcb372afb5ba5af654fad97b26cc10ae8289f1c62f.
//
// Solidity: event EnableToken(address token)
func (_Dforce *DforceFilterer) ParseEnableToken(log types.Log) (*DforceEnableToken, error) {
	event := new(DforceEnableToken)
	if err := _Dforce.contract.UnpackLog(event, "EnableToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceEnableTradeIterator is returned from FilterEnableTrade and is used to iterate over the raw logs and unpacked data for EnableTrade events raised by the Dforce contract.
type DforceEnableTradeIterator struct {
	Event *DforceEnableTrade // Event containing the contract specifics and raw log

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
func (it *DforceEnableTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceEnableTrade)
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
		it.Event = new(DforceEnableTrade)
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
func (it *DforceEnableTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceEnableTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceEnableTrade represents a EnableTrade event raised by the Dforce contract.
type DforceEnableTrade struct {
	Input  common.Address
	Output common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnableTrade is a free log retrieval operation binding the contract event 0x5a27ad565647c044a93ca5b40639a1cf6bcc69260317b5f05e4572cec28f773a.
//
// Solidity: event EnableTrade(address input, address output)
func (_Dforce *DforceFilterer) FilterEnableTrade(opts *bind.FilterOpts) (*DforceEnableTradeIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "EnableTrade")
	if err != nil {
		return nil, err
	}
	return &DforceEnableTradeIterator{contract: _Dforce.contract, event: "EnableTrade", logs: logs, sub: sub}, nil
}

// WatchEnableTrade is a free log subscription operation binding the contract event 0x5a27ad565647c044a93ca5b40639a1cf6bcc69260317b5f05e4572cec28f773a.
//
// Solidity: event EnableTrade(address input, address output)
func (_Dforce *DforceFilterer) WatchEnableTrade(opts *bind.WatchOpts, sink chan<- *DforceEnableTrade) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "EnableTrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceEnableTrade)
				if err := _Dforce.contract.UnpackLog(event, "EnableTrade", log); err != nil {
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

// ParseEnableTrade is a log parse operation binding the contract event 0x5a27ad565647c044a93ca5b40639a1cf6bcc69260317b5f05e4572cec28f773a.
//
// Solidity: event EnableTrade(address input, address output)
func (_Dforce *DforceFilterer) ParseEnableTrade(log types.Log) (*DforceEnableTrade, error) {
	event := new(DforceEnableTrade)
	if err := _Dforce.contract.UnpackLog(event, "EnableTrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the Dforce contract.
type DforceLogSetAuthorityIterator struct {
	Event *DforceLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *DforceLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceLogSetAuthority)
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
		it.Event = new(DforceLogSetAuthority)
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
func (it *DforceLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceLogSetAuthority represents a LogSetAuthority event raised by the Dforce contract.
type DforceLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Dforce *DforceFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*DforceLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &DforceLogSetAuthorityIterator{contract: _Dforce.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Dforce *DforceFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *DforceLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceLogSetAuthority)
				if err := _Dforce.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// ParseLogSetAuthority is a log parse operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Dforce *DforceFilterer) ParseLogSetAuthority(log types.Log) (*DforceLogSetAuthority, error) {
	event := new(DforceLogSetAuthority)
	if err := _Dforce.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the Dforce contract.
type DforceLogSetOwnerIterator struct {
	Event *DforceLogSetOwner // Event containing the contract specifics and raw log

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
func (it *DforceLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceLogSetOwner)
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
		it.Event = new(DforceLogSetOwner)
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
func (it *DforceLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceLogSetOwner represents a LogSetOwner event raised by the Dforce contract.
type DforceLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Dforce *DforceFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*DforceLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &DforceLogSetOwnerIterator{contract: _Dforce.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Dforce *DforceFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *DforceLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceLogSetOwner)
				if err := _Dforce.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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

// ParseLogSetOwner is a log parse operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Dforce *DforceFilterer) ParseLogSetOwner(log types.Log) (*DforceLogSetOwner, error) {
	event := new(DforceLogSetOwner)
	if err := _Dforce.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceOwnerUpdateIterator is returned from FilterOwnerUpdate and is used to iterate over the raw logs and unpacked data for OwnerUpdate events raised by the Dforce contract.
type DforceOwnerUpdateIterator struct {
	Event *DforceOwnerUpdate // Event containing the contract specifics and raw log

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
func (it *DforceOwnerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceOwnerUpdate)
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
		it.Event = new(DforceOwnerUpdate)
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
func (it *DforceOwnerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceOwnerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceOwnerUpdate represents a OwnerUpdate event raised by the Dforce contract.
type DforceOwnerUpdate struct {
	Owner    common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdate is a free log retrieval operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed owner, address indexed newOwner)
func (_Dforce *DforceFilterer) FilterOwnerUpdate(opts *bind.FilterOpts, owner []common.Address, newOwner []common.Address) (*DforceOwnerUpdateIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "OwnerUpdate", ownerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DforceOwnerUpdateIterator{contract: _Dforce.contract, event: "OwnerUpdate", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdate is a free log subscription operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed owner, address indexed newOwner)
func (_Dforce *DforceFilterer) WatchOwnerUpdate(opts *bind.WatchOpts, sink chan<- *DforceOwnerUpdate, owner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "OwnerUpdate", ownerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceOwnerUpdate)
				if err := _Dforce.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
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
// Solidity: event OwnerUpdate(address indexed owner, address indexed newOwner)
func (_Dforce *DforceFilterer) ParseOwnerUpdate(log types.Log) (*DforceOwnerUpdate, error) {
	event := new(DforceOwnerUpdate)
	if err := _Dforce.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforcePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Dforce contract.
type DforcePausedIterator struct {
	Event *DforcePaused // Event containing the contract specifics and raw log

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
func (it *DforcePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforcePaused)
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
		it.Event = new(DforcePaused)
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
func (it *DforcePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforcePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforcePaused represents a Paused event raised by the Dforce contract.
type DforcePaused struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address admin)
func (_Dforce *DforceFilterer) FilterPaused(opts *bind.FilterOpts) (*DforcePausedIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &DforcePausedIterator{contract: _Dforce.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address admin)
func (_Dforce *DforceFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *DforcePaused) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforcePaused)
				if err := _Dforce.contract.UnpackLog(event, "Paused", log); err != nil {
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
// Solidity: event Paused(address admin)
func (_Dforce *DforceFilterer) ParsePaused(log types.Log) (*DforcePaused, error) {
	event := new(DforcePaused)
	if err := _Dforce.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the Dforce contract.
type DforceSetFeeIterator struct {
	Event *DforceSetFee // Event containing the contract specifics and raw log

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
func (it *DforceSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceSetFee)
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
		it.Event = new(DforceSetFee)
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
func (it *DforceSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceSetFee represents a SetFee event raised by the Dforce contract.
type DforceSetFee struct {
	Input  common.Address
	Output common.Address
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0x44a6d70a601a6f8a85c075467e9d7245897140cbf6dd505c9d9d764459f5fb64.
//
// Solidity: event SetFee(address input, address output, uint256 fee)
func (_Dforce *DforceFilterer) FilterSetFee(opts *bind.FilterOpts) (*DforceSetFeeIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "SetFee")
	if err != nil {
		return nil, err
	}
	return &DforceSetFeeIterator{contract: _Dforce.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0x44a6d70a601a6f8a85c075467e9d7245897140cbf6dd505c9d9d764459f5fb64.
//
// Solidity: event SetFee(address input, address output, uint256 fee)
func (_Dforce *DforceFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *DforceSetFee) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "SetFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceSetFee)
				if err := _Dforce.contract.UnpackLog(event, "SetFee", log); err != nil {
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

// ParseSetFee is a log parse operation binding the contract event 0x44a6d70a601a6f8a85c075467e9d7245897140cbf6dd505c9d9d764459f5fb64.
//
// Solidity: event SetFee(address input, address output, uint256 fee)
func (_Dforce *DforceFilterer) ParseSetFee(log types.Log) (*DforceSetFee, error) {
	event := new(DforceSetFee)
	if err := _Dforce.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceSetMaxSwingIterator is returned from FilterSetMaxSwing and is used to iterate over the raw logs and unpacked data for SetMaxSwing events raised by the Dforce contract.
type DforceSetMaxSwingIterator struct {
	Event *DforceSetMaxSwing // Event containing the contract specifics and raw log

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
func (it *DforceSetMaxSwingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceSetMaxSwing)
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
		it.Event = new(DforceSetMaxSwing)
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
func (it *DforceSetMaxSwingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceSetMaxSwingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceSetMaxSwing represents a SetMaxSwing event raised by the Dforce contract.
type DforceSetMaxSwing struct {
	MaxSwing *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetMaxSwing is a free log retrieval operation binding the contract event 0x2d8b75edba0b3a319274edc8de327c1163f1b657579353fd93e6fc4a5eaafc5e.
//
// Solidity: event SetMaxSwing(uint256 maxSwing)
func (_Dforce *DforceFilterer) FilterSetMaxSwing(opts *bind.FilterOpts) (*DforceSetMaxSwingIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "SetMaxSwing")
	if err != nil {
		return nil, err
	}
	return &DforceSetMaxSwingIterator{contract: _Dforce.contract, event: "SetMaxSwing", logs: logs, sub: sub}, nil
}

// WatchSetMaxSwing is a free log subscription operation binding the contract event 0x2d8b75edba0b3a319274edc8de327c1163f1b657579353fd93e6fc4a5eaafc5e.
//
// Solidity: event SetMaxSwing(uint256 maxSwing)
func (_Dforce *DforceFilterer) WatchSetMaxSwing(opts *bind.WatchOpts, sink chan<- *DforceSetMaxSwing) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "SetMaxSwing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceSetMaxSwing)
				if err := _Dforce.contract.UnpackLog(event, "SetMaxSwing", log); err != nil {
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

// ParseSetMaxSwing is a log parse operation binding the contract event 0x2d8b75edba0b3a319274edc8de327c1163f1b657579353fd93e6fc4a5eaafc5e.
//
// Solidity: event SetMaxSwing(uint256 maxSwing)
func (_Dforce *DforceFilterer) ParseSetMaxSwing(log types.Log) (*DforceSetMaxSwing, error) {
	event := new(DforceSetMaxSwing)
	if err := _Dforce.contract.UnpackLog(event, "SetMaxSwing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceSetOracleIterator is returned from FilterSetOracle and is used to iterate over the raw logs and unpacked data for SetOracle events raised by the Dforce contract.
type DforceSetOracleIterator struct {
	Event *DforceSetOracle // Event containing the contract specifics and raw log

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
func (it *DforceSetOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceSetOracle)
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
		it.Event = new(DforceSetOracle)
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
func (it *DforceSetOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceSetOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceSetOracle represents a SetOracle event raised by the Dforce contract.
type DforceSetOracle struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetOracle is a free log retrieval operation binding the contract event 0xd3b5d1e0ffaeff528910f3663f0adace7694ab8241d58e17a91351ced2e08031.
//
// Solidity: event SetOracle(address oracle)
func (_Dforce *DforceFilterer) FilterSetOracle(opts *bind.FilterOpts) (*DforceSetOracleIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "SetOracle")
	if err != nil {
		return nil, err
	}
	return &DforceSetOracleIterator{contract: _Dforce.contract, event: "SetOracle", logs: logs, sub: sub}, nil
}

// WatchSetOracle is a free log subscription operation binding the contract event 0xd3b5d1e0ffaeff528910f3663f0adace7694ab8241d58e17a91351ced2e08031.
//
// Solidity: event SetOracle(address oracle)
func (_Dforce *DforceFilterer) WatchSetOracle(opts *bind.WatchOpts, sink chan<- *DforceSetOracle) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "SetOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceSetOracle)
				if err := _Dforce.contract.UnpackLog(event, "SetOracle", log); err != nil {
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

// ParseSetOracle is a log parse operation binding the contract event 0xd3b5d1e0ffaeff528910f3663f0adace7694ab8241d58e17a91351ced2e08031.
//
// Solidity: event SetOracle(address oracle)
func (_Dforce *DforceFilterer) ParseSetOracle(log types.Log) (*DforceSetOracle, error) {
	event := new(DforceSetOracle)
	if err := _Dforce.contract.UnpackLog(event, "SetOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Dforce contract.
type DforceSwapIterator struct {
	Event *DforceSwap // Event containing the contract specifics and raw log

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
func (it *DforceSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceSwap)
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
		it.Event = new(DforceSwap)
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
func (it *DforceSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceSwap represents a Swap event raised by the Dforce contract.
type DforceSwap struct {
	From         common.Address
	To           common.Address
	Input        common.Address
	InputAmount  *big.Int
	InputPrice   *big.Int
	Output       common.Address
	OutputAmount *big.Int
	OutputPrice  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x3c659b0d071f4048f56712c66ca2ab11885b2f4d5ad1029adbdd300aaa0c18d6.
//
// Solidity: event Swap(address from, address to, address input, uint256 inputAmount, uint256 inputPrice, address output, uint256 outputAmount, uint256 outputPrice)
func (_Dforce *DforceFilterer) FilterSwap(opts *bind.FilterOpts) (*DforceSwapIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &DforceSwapIterator{contract: _Dforce.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x3c659b0d071f4048f56712c66ca2ab11885b2f4d5ad1029adbdd300aaa0c18d6.
//
// Solidity: event Swap(address from, address to, address input, uint256 inputAmount, uint256 inputPrice, address output, uint256 outputAmount, uint256 outputPrice)
func (_Dforce *DforceFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *DforceSwap) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceSwap)
				if err := _Dforce.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x3c659b0d071f4048f56712c66ca2ab11885b2f4d5ad1029adbdd300aaa0c18d6.
//
// Solidity: event Swap(address from, address to, address input, uint256 inputAmount, uint256 inputPrice, address output, uint256 outputAmount, uint256 outputPrice)
func (_Dforce *DforceFilterer) ParseSwap(log types.Log) (*DforceSwap, error) {
	event := new(DforceSwap)
	if err := _Dforce.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceTransferInIterator is returned from FilterTransferIn and is used to iterate over the raw logs and unpacked data for TransferIn events raised by the Dforce contract.
type DforceTransferInIterator struct {
	Event *DforceTransferIn // Event containing the contract specifics and raw log

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
func (it *DforceTransferInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceTransferIn)
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
		it.Event = new(DforceTransferIn)
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
func (it *DforceTransferInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceTransferInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceTransferIn represents a TransferIn event raised by the Dforce contract.
type DforceTransferIn struct {
	Token   common.Address
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferIn is a free log retrieval operation binding the contract event 0x79f2c52b850a60522f5c49f634599f750fba2a92c9112ea49634c172d3d93ab6.
//
// Solidity: event TransferIn(address token, uint256 amount, uint256 balance)
func (_Dforce *DforceFilterer) FilterTransferIn(opts *bind.FilterOpts) (*DforceTransferInIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "TransferIn")
	if err != nil {
		return nil, err
	}
	return &DforceTransferInIterator{contract: _Dforce.contract, event: "TransferIn", logs: logs, sub: sub}, nil
}

// WatchTransferIn is a free log subscription operation binding the contract event 0x79f2c52b850a60522f5c49f634599f750fba2a92c9112ea49634c172d3d93ab6.
//
// Solidity: event TransferIn(address token, uint256 amount, uint256 balance)
func (_Dforce *DforceFilterer) WatchTransferIn(opts *bind.WatchOpts, sink chan<- *DforceTransferIn) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "TransferIn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceTransferIn)
				if err := _Dforce.contract.UnpackLog(event, "TransferIn", log); err != nil {
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

// ParseTransferIn is a log parse operation binding the contract event 0x79f2c52b850a60522f5c49f634599f750fba2a92c9112ea49634c172d3d93ab6.
//
// Solidity: event TransferIn(address token, uint256 amount, uint256 balance)
func (_Dforce *DforceFilterer) ParseTransferIn(log types.Log) (*DforceTransferIn, error) {
	event := new(DforceTransferIn)
	if err := _Dforce.contract.UnpackLog(event, "TransferIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceTransferOutIterator is returned from FilterTransferOut and is used to iterate over the raw logs and unpacked data for TransferOut events raised by the Dforce contract.
type DforceTransferOutIterator struct {
	Event *DforceTransferOut // Event containing the contract specifics and raw log

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
func (it *DforceTransferOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceTransferOut)
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
		it.Event = new(DforceTransferOut)
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
func (it *DforceTransferOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceTransferOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceTransferOut represents a TransferOut event raised by the Dforce contract.
type DforceTransferOut struct {
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Balance  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferOut is a free log retrieval operation binding the contract event 0x08d474da06d67ac0e0658410133ac9d6a35c74cfb6c9da5f68bbd68e66230534.
//
// Solidity: event TransferOut(address token, address receiver, uint256 amount, uint256 balance)
func (_Dforce *DforceFilterer) FilterTransferOut(opts *bind.FilterOpts) (*DforceTransferOutIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "TransferOut")
	if err != nil {
		return nil, err
	}
	return &DforceTransferOutIterator{contract: _Dforce.contract, event: "TransferOut", logs: logs, sub: sub}, nil
}

// WatchTransferOut is a free log subscription operation binding the contract event 0x08d474da06d67ac0e0658410133ac9d6a35c74cfb6c9da5f68bbd68e66230534.
//
// Solidity: event TransferOut(address token, address receiver, uint256 amount, uint256 balance)
func (_Dforce *DforceFilterer) WatchTransferOut(opts *bind.WatchOpts, sink chan<- *DforceTransferOut) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "TransferOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceTransferOut)
				if err := _Dforce.contract.UnpackLog(event, "TransferOut", log); err != nil {
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

// ParseTransferOut is a log parse operation binding the contract event 0x08d474da06d67ac0e0658410133ac9d6a35c74cfb6c9da5f68bbd68e66230534.
//
// Solidity: event TransferOut(address token, address receiver, uint256 amount, uint256 balance)
func (_Dforce *DforceFilterer) ParseTransferOut(log types.Log) (*DforceTransferOut, error) {
	event := new(DforceTransferOut)
	if err := _Dforce.contract.UnpackLog(event, "TransferOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DforceUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Dforce contract.
type DforceUnpausedIterator struct {
	Event *DforceUnpaused // Event containing the contract specifics and raw log

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
func (it *DforceUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DforceUnpaused)
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
		it.Event = new(DforceUnpaused)
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
func (it *DforceUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DforceUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DforceUnpaused represents a Unpaused event raised by the Dforce contract.
type DforceUnpaused struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address admin)
func (_Dforce *DforceFilterer) FilterUnpaused(opts *bind.FilterOpts) (*DforceUnpausedIterator, error) {

	logs, sub, err := _Dforce.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &DforceUnpausedIterator{contract: _Dforce.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address admin)
func (_Dforce *DforceFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *DforceUnpaused) (event.Subscription, error) {

	logs, sub, err := _Dforce.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DforceUnpaused)
				if err := _Dforce.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
// Solidity: event Unpaused(address admin)
func (_Dforce *DforceFilterer) ParseUnpaused(log types.Log) (*DforceUnpaused, error) {
	event := new(DforceUnpaused)
	if err := _Dforce.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
