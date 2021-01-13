// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc

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

// CvaultcontractABI is the input ABI used to generate the binding from.
const CvaultcontractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LPTokenClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LiquidityAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"log\",\"type\":\"string\"}],\"name\":\"Log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPGenerationCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPperETHUnit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"agreesToTermsOutlinedInLiquidityGenerationParticipationAgreement\",\"type\":\"bool\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addLiquidityToUniswapCORExWETHPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fromBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimLPTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractStartTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createUniswapPairMainnet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyDrain24hAfterLiquidityGenerationEventIsDone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ethContributed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeDistributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCurrentVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPriorVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSecondsLeftInLiquidityGenerationEvent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityGenerationOngoing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityGenerationParticipationAgreement\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeDistributor\",\"type\":\"address\"}],\"name\":\"setFeeDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_transferCheckerAddress\",\"type\":\"address\"}],\"name\":\"setShouldTransferChecker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenUniswapPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalETHContributed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLPTokensMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferCheckerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapFactory\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapRouterV2\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Cvaultcontract is an auto generated Go binding around an Ethereum contract.
type Cvaultcontract struct {
	CvaultcontractCaller     // Read-only binding to the contract
	CvaultcontractTransactor // Write-only binding to the contract
	CvaultcontractFilterer   // Log filterer for contract events
}

// CvaultcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CvaultcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CvaultcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CvaultcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CvaultcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CvaultcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CvaultcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CvaultcontractSession struct {
	Contract     *Cvaultcontract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CvaultcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CvaultcontractCallerSession struct {
	Contract *CvaultcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CvaultcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CvaultcontractTransactorSession struct {
	Contract     *CvaultcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CvaultcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CvaultcontractRaw struct {
	Contract *Cvaultcontract // Generic contract binding to access the raw methods on
}

// CvaultcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CvaultcontractCallerRaw struct {
	Contract *CvaultcontractCaller // Generic read-only contract binding to access the raw methods on
}

// CvaultcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CvaultcontractTransactorRaw struct {
	Contract *CvaultcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCvaultcontract creates a new instance of Cvaultcontract, bound to a specific deployed contract.
func NewCvaultcontract(address common.Address, backend bind.ContractBackend) (*Cvaultcontract, error) {
	contract, err := bindCvaultcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cvaultcontract{CvaultcontractCaller: CvaultcontractCaller{contract: contract}, CvaultcontractTransactor: CvaultcontractTransactor{contract: contract}, CvaultcontractFilterer: CvaultcontractFilterer{contract: contract}}, nil
}

// NewCvaultcontractCaller creates a new read-only instance of Cvaultcontract, bound to a specific deployed contract.
func NewCvaultcontractCaller(address common.Address, caller bind.ContractCaller) (*CvaultcontractCaller, error) {
	contract, err := bindCvaultcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CvaultcontractCaller{contract: contract}, nil
}

// NewCvaultcontractTransactor creates a new write-only instance of Cvaultcontract, bound to a specific deployed contract.
func NewCvaultcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*CvaultcontractTransactor, error) {
	contract, err := bindCvaultcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CvaultcontractTransactor{contract: contract}, nil
}

// NewCvaultcontractFilterer creates a new log filterer instance of Cvaultcontract, bound to a specific deployed contract.
func NewCvaultcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*CvaultcontractFilterer, error) {
	contract, err := bindCvaultcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CvaultcontractFilterer{contract: contract}, nil
}

// bindCvaultcontract binds a generic wrapper to an already deployed contract.
func bindCvaultcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CvaultcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cvaultcontract *CvaultcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cvaultcontract.Contract.CvaultcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cvaultcontract *CvaultcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.CvaultcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cvaultcontract *CvaultcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.CvaultcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cvaultcontract *CvaultcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cvaultcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cvaultcontract *CvaultcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cvaultcontract *CvaultcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Cvaultcontract *CvaultcontractCaller) DELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "DELEGATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Cvaultcontract *CvaultcontractSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _Cvaultcontract.Contract.DELEGATIONTYPEHASH(&_Cvaultcontract.CallOpts)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Cvaultcontract *CvaultcontractCallerSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _Cvaultcontract.Contract.DELEGATIONTYPEHASH(&_Cvaultcontract.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Cvaultcontract *CvaultcontractCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Cvaultcontract *CvaultcontractSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Cvaultcontract.Contract.DOMAINTYPEHASH(&_Cvaultcontract.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Cvaultcontract *CvaultcontractCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Cvaultcontract.Contract.DOMAINTYPEHASH(&_Cvaultcontract.CallOpts)
}

// LPGenerationCompleted is a free data retrieval call binding the contract method 0x14b8fecc.
//
// Solidity: function LPGenerationCompleted() view returns(bool)
func (_Cvaultcontract *CvaultcontractCaller) LPGenerationCompleted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "LPGenerationCompleted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LPGenerationCompleted is a free data retrieval call binding the contract method 0x14b8fecc.
//
// Solidity: function LPGenerationCompleted() view returns(bool)
func (_Cvaultcontract *CvaultcontractSession) LPGenerationCompleted() (bool, error) {
	return _Cvaultcontract.Contract.LPGenerationCompleted(&_Cvaultcontract.CallOpts)
}

// LPGenerationCompleted is a free data retrieval call binding the contract method 0x14b8fecc.
//
// Solidity: function LPGenerationCompleted() view returns(bool)
func (_Cvaultcontract *CvaultcontractCallerSession) LPGenerationCompleted() (bool, error) {
	return _Cvaultcontract.Contract.LPGenerationCompleted(&_Cvaultcontract.CallOpts)
}

// LPperETHUnit is a free data retrieval call binding the contract method 0x31a22a20.
//
// Solidity: function LPperETHUnit() view returns(uint256)
func (_Cvaultcontract *CvaultcontractCaller) LPperETHUnit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "LPperETHUnit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LPperETHUnit is a free data retrieval call binding the contract method 0x31a22a20.
//
// Solidity: function LPperETHUnit() view returns(uint256)
func (_Cvaultcontract *CvaultcontractSession) LPperETHUnit() (*big.Int, error) {
	return _Cvaultcontract.Contract.LPperETHUnit(&_Cvaultcontract.CallOpts)
}

// LPperETHUnit is a free data retrieval call binding the contract method 0x31a22a20.
//
// Solidity: function LPperETHUnit() view returns(uint256)
func (_Cvaultcontract *CvaultcontractCallerSession) LPperETHUnit() (*big.Int, error) {
	return _Cvaultcontract.Contract.LPperETHUnit(&_Cvaultcontract.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cvaultcontract *CvaultcontractCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cvaultcontract *CvaultcontractSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Cvaultcontract.Contract.Allowance(&_Cvaultcontract.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cvaultcontract *CvaultcontractCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Cvaultcontract.Contract.Allowance(&_Cvaultcontract.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Cvaultcontract *CvaultcontractCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Cvaultcontract *CvaultcontractSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Cvaultcontract.Contract.BalanceOf(&_Cvaultcontract.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Cvaultcontract *CvaultcontractCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Cvaultcontract.Contract.BalanceOf(&_Cvaultcontract.CallOpts, _owner)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint256 votes)
func (_Cvaultcontract *CvaultcontractCaller) Checkpoints(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "checkpoints", arg0, arg1)

	outstruct := new(struct {
		FromBlock uint32
		Votes     *big.Int
	})

	outstruct.FromBlock = out[0].(uint32)
	outstruct.Votes = out[1].(*big.Int)

	return *outstruct, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint256 votes)
func (_Cvaultcontract *CvaultcontractSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _Cvaultcontract.Contract.Checkpoints(&_Cvaultcontract.CallOpts, arg0, arg1)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint256 votes)
func (_Cvaultcontract *CvaultcontractCallerSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _Cvaultcontract.Contract.Checkpoints(&_Cvaultcontract.CallOpts, arg0, arg1)
}

// ContractStartTimestamp is a free data retrieval call binding the contract method 0x6a2f796c.
//
// Solidity: function contractStartTimestamp() view returns(uint256)
func (_Cvaultcontract *CvaultcontractCaller) ContractStartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "contractStartTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractStartTimestamp is a free data retrieval call binding the contract method 0x6a2f796c.
//
// Solidity: function contractStartTimestamp() view returns(uint256)
func (_Cvaultcontract *CvaultcontractSession) ContractStartTimestamp() (*big.Int, error) {
	return _Cvaultcontract.Contract.ContractStartTimestamp(&_Cvaultcontract.CallOpts)
}

// ContractStartTimestamp is a free data retrieval call binding the contract method 0x6a2f796c.
//
// Solidity: function contractStartTimestamp() view returns(uint256)
func (_Cvaultcontract *CvaultcontractCallerSession) ContractStartTimestamp() (*big.Int, error) {
	return _Cvaultcontract.Contract.ContractStartTimestamp(&_Cvaultcontract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cvaultcontract *CvaultcontractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cvaultcontract *CvaultcontractSession) Decimals() (uint8, error) {
	return _Cvaultcontract.Contract.Decimals(&_Cvaultcontract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cvaultcontract *CvaultcontractCallerSession) Decimals() (uint8, error) {
	return _Cvaultcontract.Contract.Decimals(&_Cvaultcontract.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_Cvaultcontract *CvaultcontractCaller) Delegates(opts *bind.CallOpts, delegator common.Address) (common.Address, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "delegates", delegator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_Cvaultcontract *CvaultcontractSession) Delegates(delegator common.Address) (common.Address, error) {
	return _Cvaultcontract.Contract.Delegates(&_Cvaultcontract.CallOpts, delegator)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_Cvaultcontract *CvaultcontractCallerSession) Delegates(delegator common.Address) (common.Address, error) {
	return _Cvaultcontract.Contract.Delegates(&_Cvaultcontract.CallOpts, delegator)
}

// EthContributed is a free data retrieval call binding the contract method 0xf96f5b35.
//
// Solidity: function ethContributed(address ) view returns(uint256)
func (_Cvaultcontract *CvaultcontractCaller) EthContributed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "ethContributed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthContributed is a free data retrieval call binding the contract method 0xf96f5b35.
//
// Solidity: function ethContributed(address ) view returns(uint256)
func (_Cvaultcontract *CvaultcontractSession) EthContributed(arg0 common.Address) (*big.Int, error) {
	return _Cvaultcontract.Contract.EthContributed(&_Cvaultcontract.CallOpts, arg0)
}

// EthContributed is a free data retrieval call binding the contract method 0xf96f5b35.
//
// Solidity: function ethContributed(address ) view returns(uint256)
func (_Cvaultcontract *CvaultcontractCallerSession) EthContributed(arg0 common.Address) (*big.Int, error) {
	return _Cvaultcontract.Contract.EthContributed(&_Cvaultcontract.CallOpts, arg0)
}

// FeeDistributor is a free data retrieval call binding the contract method 0x0d43e8ad.
//
// Solidity: function feeDistributor() view returns(address)
func (_Cvaultcontract *CvaultcontractCaller) FeeDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "feeDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeDistributor is a free data retrieval call binding the contract method 0x0d43e8ad.
//
// Solidity: function feeDistributor() view returns(address)
func (_Cvaultcontract *CvaultcontractSession) FeeDistributor() (common.Address, error) {
	return _Cvaultcontract.Contract.FeeDistributor(&_Cvaultcontract.CallOpts)
}

// FeeDistributor is a free data retrieval call binding the contract method 0x0d43e8ad.
//
// Solidity: function feeDistributor() view returns(address)
func (_Cvaultcontract *CvaultcontractCallerSession) FeeDistributor() (common.Address, error) {
	return _Cvaultcontract.Contract.FeeDistributor(&_Cvaultcontract.CallOpts)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint256)
func (_Cvaultcontract *CvaultcontractCaller) GetCurrentVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "getCurrentVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint256)
func (_Cvaultcontract *CvaultcontractSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _Cvaultcontract.Contract.GetCurrentVotes(&_Cvaultcontract.CallOpts, account)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint256)
func (_Cvaultcontract *CvaultcontractCallerSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _Cvaultcontract.Contract.GetCurrentVotes(&_Cvaultcontract.CallOpts, account)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint256)
func (_Cvaultcontract *CvaultcontractCaller) GetPriorVotes(opts *bind.CallOpts, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "getPriorVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint256)
func (_Cvaultcontract *CvaultcontractSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _Cvaultcontract.Contract.GetPriorVotes(&_Cvaultcontract.CallOpts, account, blockNumber)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint256)
func (_Cvaultcontract *CvaultcontractCallerSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _Cvaultcontract.Contract.GetPriorVotes(&_Cvaultcontract.CallOpts, account, blockNumber)
}

// GetSecondsLeftInLiquidityGenerationEvent is a free data retrieval call binding the contract method 0x5b5f3e87.
//
// Solidity: function getSecondsLeftInLiquidityGenerationEvent() view returns(uint256)
func (_Cvaultcontract *CvaultcontractCaller) GetSecondsLeftInLiquidityGenerationEvent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "getSecondsLeftInLiquidityGenerationEvent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSecondsLeftInLiquidityGenerationEvent is a free data retrieval call binding the contract method 0x5b5f3e87.
//
// Solidity: function getSecondsLeftInLiquidityGenerationEvent() view returns(uint256)
func (_Cvaultcontract *CvaultcontractSession) GetSecondsLeftInLiquidityGenerationEvent() (*big.Int, error) {
	return _Cvaultcontract.Contract.GetSecondsLeftInLiquidityGenerationEvent(&_Cvaultcontract.CallOpts)
}

// GetSecondsLeftInLiquidityGenerationEvent is a free data retrieval call binding the contract method 0x5b5f3e87.
//
// Solidity: function getSecondsLeftInLiquidityGenerationEvent() view returns(uint256)
func (_Cvaultcontract *CvaultcontractCallerSession) GetSecondsLeftInLiquidityGenerationEvent() (*big.Int, error) {
	return _Cvaultcontract.Contract.GetSecondsLeftInLiquidityGenerationEvent(&_Cvaultcontract.CallOpts)
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() view returns(uint256)
func (_Cvaultcontract *CvaultcontractCaller) InitialSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "initialSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() view returns(uint256)
func (_Cvaultcontract *CvaultcontractSession) InitialSupply() (*big.Int, error) {
	return _Cvaultcontract.Contract.InitialSupply(&_Cvaultcontract.CallOpts)
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() view returns(uint256)
func (_Cvaultcontract *CvaultcontractCallerSession) InitialSupply() (*big.Int, error) {
	return _Cvaultcontract.Contract.InitialSupply(&_Cvaultcontract.CallOpts)
}

// LiquidityGenerationOngoing is a free data retrieval call binding the contract method 0x60a02590.
//
// Solidity: function liquidityGenerationOngoing() view returns(bool)
func (_Cvaultcontract *CvaultcontractCaller) LiquidityGenerationOngoing(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "liquidityGenerationOngoing")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LiquidityGenerationOngoing is a free data retrieval call binding the contract method 0x60a02590.
//
// Solidity: function liquidityGenerationOngoing() view returns(bool)
func (_Cvaultcontract *CvaultcontractSession) LiquidityGenerationOngoing() (bool, error) {
	return _Cvaultcontract.Contract.LiquidityGenerationOngoing(&_Cvaultcontract.CallOpts)
}

// LiquidityGenerationOngoing is a free data retrieval call binding the contract method 0x60a02590.
//
// Solidity: function liquidityGenerationOngoing() view returns(bool)
func (_Cvaultcontract *CvaultcontractCallerSession) LiquidityGenerationOngoing() (bool, error) {
	return _Cvaultcontract.Contract.LiquidityGenerationOngoing(&_Cvaultcontract.CallOpts)
}

// LiquidityGenerationParticipationAgreement is a free data retrieval call binding the contract method 0x8e8e2925.
//
// Solidity: function liquidityGenerationParticipationAgreement() view returns(string)
func (_Cvaultcontract *CvaultcontractCaller) LiquidityGenerationParticipationAgreement(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "liquidityGenerationParticipationAgreement")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LiquidityGenerationParticipationAgreement is a free data retrieval call binding the contract method 0x8e8e2925.
//
// Solidity: function liquidityGenerationParticipationAgreement() view returns(string)
func (_Cvaultcontract *CvaultcontractSession) LiquidityGenerationParticipationAgreement() (string, error) {
	return _Cvaultcontract.Contract.LiquidityGenerationParticipationAgreement(&_Cvaultcontract.CallOpts)
}

// LiquidityGenerationParticipationAgreement is a free data retrieval call binding the contract method 0x8e8e2925.
//
// Solidity: function liquidityGenerationParticipationAgreement() view returns(string)
func (_Cvaultcontract *CvaultcontractCallerSession) LiquidityGenerationParticipationAgreement() (string, error) {
	return _Cvaultcontract.Contract.LiquidityGenerationParticipationAgreement(&_Cvaultcontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cvaultcontract *CvaultcontractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cvaultcontract *CvaultcontractSession) Name() (string, error) {
	return _Cvaultcontract.Contract.Name(&_Cvaultcontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cvaultcontract *CvaultcontractCallerSession) Name() (string, error) {
	return _Cvaultcontract.Contract.Name(&_Cvaultcontract.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Cvaultcontract *CvaultcontractCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Cvaultcontract *CvaultcontractSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Cvaultcontract.Contract.Nonces(&_Cvaultcontract.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Cvaultcontract *CvaultcontractCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Cvaultcontract.Contract.Nonces(&_Cvaultcontract.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Cvaultcontract *CvaultcontractCaller) NumCheckpoints(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "numCheckpoints", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Cvaultcontract *CvaultcontractSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _Cvaultcontract.Contract.NumCheckpoints(&_Cvaultcontract.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Cvaultcontract *CvaultcontractCallerSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _Cvaultcontract.Contract.NumCheckpoints(&_Cvaultcontract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cvaultcontract *CvaultcontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cvaultcontract *CvaultcontractSession) Owner() (common.Address, error) {
	return _Cvaultcontract.Contract.Owner(&_Cvaultcontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cvaultcontract *CvaultcontractCallerSession) Owner() (common.Address, error) {
	return _Cvaultcontract.Contract.Owner(&_Cvaultcontract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cvaultcontract *CvaultcontractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cvaultcontract *CvaultcontractSession) Symbol() (string, error) {
	return _Cvaultcontract.Contract.Symbol(&_Cvaultcontract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cvaultcontract *CvaultcontractCallerSession) Symbol() (string, error) {
	return _Cvaultcontract.Contract.Symbol(&_Cvaultcontract.CallOpts)
}

// TokenUniswapPair is a free data retrieval call binding the contract method 0x4d332457.
//
// Solidity: function tokenUniswapPair() view returns(address)
func (_Cvaultcontract *CvaultcontractCaller) TokenUniswapPair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "tokenUniswapPair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenUniswapPair is a free data retrieval call binding the contract method 0x4d332457.
//
// Solidity: function tokenUniswapPair() view returns(address)
func (_Cvaultcontract *CvaultcontractSession) TokenUniswapPair() (common.Address, error) {
	return _Cvaultcontract.Contract.TokenUniswapPair(&_Cvaultcontract.CallOpts)
}

// TokenUniswapPair is a free data retrieval call binding the contract method 0x4d332457.
//
// Solidity: function tokenUniswapPair() view returns(address)
func (_Cvaultcontract *CvaultcontractCallerSession) TokenUniswapPair() (common.Address, error) {
	return _Cvaultcontract.Contract.TokenUniswapPair(&_Cvaultcontract.CallOpts)
}

// TotalETHContributed is a free data retrieval call binding the contract method 0x23399434.
//
// Solidity: function totalETHContributed() view returns(uint256)
func (_Cvaultcontract *CvaultcontractCaller) TotalETHContributed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "totalETHContributed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalETHContributed is a free data retrieval call binding the contract method 0x23399434.
//
// Solidity: function totalETHContributed() view returns(uint256)
func (_Cvaultcontract *CvaultcontractSession) TotalETHContributed() (*big.Int, error) {
	return _Cvaultcontract.Contract.TotalETHContributed(&_Cvaultcontract.CallOpts)
}

// TotalETHContributed is a free data retrieval call binding the contract method 0x23399434.
//
// Solidity: function totalETHContributed() view returns(uint256)
func (_Cvaultcontract *CvaultcontractCallerSession) TotalETHContributed() (*big.Int, error) {
	return _Cvaultcontract.Contract.TotalETHContributed(&_Cvaultcontract.CallOpts)
}

// TotalLPTokensMinted is a free data retrieval call binding the contract method 0x002b1329.
//
// Solidity: function totalLPTokensMinted() view returns(uint256)
func (_Cvaultcontract *CvaultcontractCaller) TotalLPTokensMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "totalLPTokensMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLPTokensMinted is a free data retrieval call binding the contract method 0x002b1329.
//
// Solidity: function totalLPTokensMinted() view returns(uint256)
func (_Cvaultcontract *CvaultcontractSession) TotalLPTokensMinted() (*big.Int, error) {
	return _Cvaultcontract.Contract.TotalLPTokensMinted(&_Cvaultcontract.CallOpts)
}

// TotalLPTokensMinted is a free data retrieval call binding the contract method 0x002b1329.
//
// Solidity: function totalLPTokensMinted() view returns(uint256)
func (_Cvaultcontract *CvaultcontractCallerSession) TotalLPTokensMinted() (*big.Int, error) {
	return _Cvaultcontract.Contract.TotalLPTokensMinted(&_Cvaultcontract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cvaultcontract *CvaultcontractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cvaultcontract *CvaultcontractSession) TotalSupply() (*big.Int, error) {
	return _Cvaultcontract.Contract.TotalSupply(&_Cvaultcontract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cvaultcontract *CvaultcontractCallerSession) TotalSupply() (*big.Int, error) {
	return _Cvaultcontract.Contract.TotalSupply(&_Cvaultcontract.CallOpts)
}

// TransferCheckerAddress is a free data retrieval call binding the contract method 0xb2aef26b.
//
// Solidity: function transferCheckerAddress() view returns(address)
func (_Cvaultcontract *CvaultcontractCaller) TransferCheckerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "transferCheckerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TransferCheckerAddress is a free data retrieval call binding the contract method 0xb2aef26b.
//
// Solidity: function transferCheckerAddress() view returns(address)
func (_Cvaultcontract *CvaultcontractSession) TransferCheckerAddress() (common.Address, error) {
	return _Cvaultcontract.Contract.TransferCheckerAddress(&_Cvaultcontract.CallOpts)
}

// TransferCheckerAddress is a free data retrieval call binding the contract method 0xb2aef26b.
//
// Solidity: function transferCheckerAddress() view returns(address)
func (_Cvaultcontract *CvaultcontractCallerSession) TransferCheckerAddress() (common.Address, error) {
	return _Cvaultcontract.Contract.TransferCheckerAddress(&_Cvaultcontract.CallOpts)
}

// UniswapFactory is a free data retrieval call binding the contract method 0x8bdb2afa.
//
// Solidity: function uniswapFactory() view returns(address)
func (_Cvaultcontract *CvaultcontractCaller) UniswapFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "uniswapFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapFactory is a free data retrieval call binding the contract method 0x8bdb2afa.
//
// Solidity: function uniswapFactory() view returns(address)
func (_Cvaultcontract *CvaultcontractSession) UniswapFactory() (common.Address, error) {
	return _Cvaultcontract.Contract.UniswapFactory(&_Cvaultcontract.CallOpts)
}

// UniswapFactory is a free data retrieval call binding the contract method 0x8bdb2afa.
//
// Solidity: function uniswapFactory() view returns(address)
func (_Cvaultcontract *CvaultcontractCallerSession) UniswapFactory() (common.Address, error) {
	return _Cvaultcontract.Contract.UniswapFactory(&_Cvaultcontract.CallOpts)
}

// UniswapRouterV2 is a free data retrieval call binding the contract method 0x596fa9e3.
//
// Solidity: function uniswapRouterV2() view returns(address)
func (_Cvaultcontract *CvaultcontractCaller) UniswapRouterV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cvaultcontract.contract.Call(opts, &out, "uniswapRouterV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapRouterV2 is a free data retrieval call binding the contract method 0x596fa9e3.
//
// Solidity: function uniswapRouterV2() view returns(address)
func (_Cvaultcontract *CvaultcontractSession) UniswapRouterV2() (common.Address, error) {
	return _Cvaultcontract.Contract.UniswapRouterV2(&_Cvaultcontract.CallOpts)
}

// UniswapRouterV2 is a free data retrieval call binding the contract method 0x596fa9e3.
//
// Solidity: function uniswapRouterV2() view returns(address)
func (_Cvaultcontract *CvaultcontractCallerSession) UniswapRouterV2() (common.Address, error) {
	return _Cvaultcontract.Contract.UniswapRouterV2(&_Cvaultcontract.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xda620cd7.
//
// Solidity: function addLiquidity(bool agreesToTermsOutlinedInLiquidityGenerationParticipationAgreement) payable returns()
func (_Cvaultcontract *CvaultcontractTransactor) AddLiquidity(opts *bind.TransactOpts, agreesToTermsOutlinedInLiquidityGenerationParticipationAgreement bool) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "addLiquidity", agreesToTermsOutlinedInLiquidityGenerationParticipationAgreement)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xda620cd7.
//
// Solidity: function addLiquidity(bool agreesToTermsOutlinedInLiquidityGenerationParticipationAgreement) payable returns()
func (_Cvaultcontract *CvaultcontractSession) AddLiquidity(agreesToTermsOutlinedInLiquidityGenerationParticipationAgreement bool) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.AddLiquidity(&_Cvaultcontract.TransactOpts, agreesToTermsOutlinedInLiquidityGenerationParticipationAgreement)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xda620cd7.
//
// Solidity: function addLiquidity(bool agreesToTermsOutlinedInLiquidityGenerationParticipationAgreement) payable returns()
func (_Cvaultcontract *CvaultcontractTransactorSession) AddLiquidity(agreesToTermsOutlinedInLiquidityGenerationParticipationAgreement bool) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.AddLiquidity(&_Cvaultcontract.TransactOpts, agreesToTermsOutlinedInLiquidityGenerationParticipationAgreement)
}

// AddLiquidityToUniswapCORExWETHPair is a paid mutator transaction binding the contract method 0x09a2ba10.
//
// Solidity: function addLiquidityToUniswapCORExWETHPair() returns()
func (_Cvaultcontract *CvaultcontractTransactor) AddLiquidityToUniswapCORExWETHPair(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "addLiquidityToUniswapCORExWETHPair")
}

// AddLiquidityToUniswapCORExWETHPair is a paid mutator transaction binding the contract method 0x09a2ba10.
//
// Solidity: function addLiquidityToUniswapCORExWETHPair() returns()
func (_Cvaultcontract *CvaultcontractSession) AddLiquidityToUniswapCORExWETHPair() (*types.Transaction, error) {
	return _Cvaultcontract.Contract.AddLiquidityToUniswapCORExWETHPair(&_Cvaultcontract.TransactOpts)
}

// AddLiquidityToUniswapCORExWETHPair is a paid mutator transaction binding the contract method 0x09a2ba10.
//
// Solidity: function addLiquidityToUniswapCORExWETHPair() returns()
func (_Cvaultcontract *CvaultcontractTransactorSession) AddLiquidityToUniswapCORExWETHPair() (*types.Transaction, error) {
	return _Cvaultcontract.Contract.AddLiquidityToUniswapCORExWETHPair(&_Cvaultcontract.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cvaultcontract *CvaultcontractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cvaultcontract *CvaultcontractSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.Approve(&_Cvaultcontract.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cvaultcontract *CvaultcontractTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.Approve(&_Cvaultcontract.TransactOpts, spender, amount)
}

// ClaimLPTokens is a paid mutator transaction binding the contract method 0x38af6632.
//
// Solidity: function claimLPTokens() returns()
func (_Cvaultcontract *CvaultcontractTransactor) ClaimLPTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "claimLPTokens")
}

// ClaimLPTokens is a paid mutator transaction binding the contract method 0x38af6632.
//
// Solidity: function claimLPTokens() returns()
func (_Cvaultcontract *CvaultcontractSession) ClaimLPTokens() (*types.Transaction, error) {
	return _Cvaultcontract.Contract.ClaimLPTokens(&_Cvaultcontract.TransactOpts)
}

// ClaimLPTokens is a paid mutator transaction binding the contract method 0x38af6632.
//
// Solidity: function claimLPTokens() returns()
func (_Cvaultcontract *CvaultcontractTransactorSession) ClaimLPTokens() (*types.Transaction, error) {
	return _Cvaultcontract.Contract.ClaimLPTokens(&_Cvaultcontract.TransactOpts)
}

// CreateUniswapPairMainnet is a paid mutator transaction binding the contract method 0x75b208bc.
//
// Solidity: function createUniswapPairMainnet() returns(address)
func (_Cvaultcontract *CvaultcontractTransactor) CreateUniswapPairMainnet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "createUniswapPairMainnet")
}

// CreateUniswapPairMainnet is a paid mutator transaction binding the contract method 0x75b208bc.
//
// Solidity: function createUniswapPairMainnet() returns(address)
func (_Cvaultcontract *CvaultcontractSession) CreateUniswapPairMainnet() (*types.Transaction, error) {
	return _Cvaultcontract.Contract.CreateUniswapPairMainnet(&_Cvaultcontract.TransactOpts)
}

// CreateUniswapPairMainnet is a paid mutator transaction binding the contract method 0x75b208bc.
//
// Solidity: function createUniswapPairMainnet() returns(address)
func (_Cvaultcontract *CvaultcontractTransactorSession) CreateUniswapPairMainnet() (*types.Transaction, error) {
	return _Cvaultcontract.Contract.CreateUniswapPairMainnet(&_Cvaultcontract.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Cvaultcontract *CvaultcontractTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Cvaultcontract *CvaultcontractSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.DecreaseAllowance(&_Cvaultcontract.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Cvaultcontract *CvaultcontractTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.DecreaseAllowance(&_Cvaultcontract.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Cvaultcontract *CvaultcontractTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Cvaultcontract *CvaultcontractSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.Delegate(&_Cvaultcontract.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Cvaultcontract *CvaultcontractTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.Delegate(&_Cvaultcontract.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Cvaultcontract *CvaultcontractTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Cvaultcontract *CvaultcontractSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.DelegateBySig(&_Cvaultcontract.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Cvaultcontract *CvaultcontractTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.DelegateBySig(&_Cvaultcontract.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// EmergencyDrain24hAfterLiquidityGenerationEventIsDone is a paid mutator transaction binding the contract method 0x4f1a0f7d.
//
// Solidity: function emergencyDrain24hAfterLiquidityGenerationEventIsDone() returns()
func (_Cvaultcontract *CvaultcontractTransactor) EmergencyDrain24hAfterLiquidityGenerationEventIsDone(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "emergencyDrain24hAfterLiquidityGenerationEventIsDone")
}

// EmergencyDrain24hAfterLiquidityGenerationEventIsDone is a paid mutator transaction binding the contract method 0x4f1a0f7d.
//
// Solidity: function emergencyDrain24hAfterLiquidityGenerationEventIsDone() returns()
func (_Cvaultcontract *CvaultcontractSession) EmergencyDrain24hAfterLiquidityGenerationEventIsDone() (*types.Transaction, error) {
	return _Cvaultcontract.Contract.EmergencyDrain24hAfterLiquidityGenerationEventIsDone(&_Cvaultcontract.TransactOpts)
}

// EmergencyDrain24hAfterLiquidityGenerationEventIsDone is a paid mutator transaction binding the contract method 0x4f1a0f7d.
//
// Solidity: function emergencyDrain24hAfterLiquidityGenerationEventIsDone() returns()
func (_Cvaultcontract *CvaultcontractTransactorSession) EmergencyDrain24hAfterLiquidityGenerationEventIsDone() (*types.Transaction, error) {
	return _Cvaultcontract.Contract.EmergencyDrain24hAfterLiquidityGenerationEventIsDone(&_Cvaultcontract.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Cvaultcontract *CvaultcontractTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Cvaultcontract *CvaultcontractSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.IncreaseAllowance(&_Cvaultcontract.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Cvaultcontract *CvaultcontractTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.IncreaseAllowance(&_Cvaultcontract.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cvaultcontract *CvaultcontractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cvaultcontract *CvaultcontractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cvaultcontract.Contract.RenounceOwnership(&_Cvaultcontract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cvaultcontract *CvaultcontractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cvaultcontract.Contract.RenounceOwnership(&_Cvaultcontract.TransactOpts)
}

// SetFeeDistributor is a paid mutator transaction binding the contract method 0xccfc2e8d.
//
// Solidity: function setFeeDistributor(address _feeDistributor) returns()
func (_Cvaultcontract *CvaultcontractTransactor) SetFeeDistributor(opts *bind.TransactOpts, _feeDistributor common.Address) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "setFeeDistributor", _feeDistributor)
}

// SetFeeDistributor is a paid mutator transaction binding the contract method 0xccfc2e8d.
//
// Solidity: function setFeeDistributor(address _feeDistributor) returns()
func (_Cvaultcontract *CvaultcontractSession) SetFeeDistributor(_feeDistributor common.Address) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.SetFeeDistributor(&_Cvaultcontract.TransactOpts, _feeDistributor)
}

// SetFeeDistributor is a paid mutator transaction binding the contract method 0xccfc2e8d.
//
// Solidity: function setFeeDistributor(address _feeDistributor) returns()
func (_Cvaultcontract *CvaultcontractTransactorSession) SetFeeDistributor(_feeDistributor common.Address) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.SetFeeDistributor(&_Cvaultcontract.TransactOpts, _feeDistributor)
}

// SetShouldTransferChecker is a paid mutator transaction binding the contract method 0x10a7a659.
//
// Solidity: function setShouldTransferChecker(address _transferCheckerAddress) returns()
func (_Cvaultcontract *CvaultcontractTransactor) SetShouldTransferChecker(opts *bind.TransactOpts, _transferCheckerAddress common.Address) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "setShouldTransferChecker", _transferCheckerAddress)
}

// SetShouldTransferChecker is a paid mutator transaction binding the contract method 0x10a7a659.
//
// Solidity: function setShouldTransferChecker(address _transferCheckerAddress) returns()
func (_Cvaultcontract *CvaultcontractSession) SetShouldTransferChecker(_transferCheckerAddress common.Address) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.SetShouldTransferChecker(&_Cvaultcontract.TransactOpts, _transferCheckerAddress)
}

// SetShouldTransferChecker is a paid mutator transaction binding the contract method 0x10a7a659.
//
// Solidity: function setShouldTransferChecker(address _transferCheckerAddress) returns()
func (_Cvaultcontract *CvaultcontractTransactorSession) SetShouldTransferChecker(_transferCheckerAddress common.Address) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.SetShouldTransferChecker(&_Cvaultcontract.TransactOpts, _transferCheckerAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cvaultcontract *CvaultcontractTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cvaultcontract *CvaultcontractSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.Transfer(&_Cvaultcontract.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cvaultcontract *CvaultcontractTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.Transfer(&_Cvaultcontract.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cvaultcontract *CvaultcontractTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cvaultcontract *CvaultcontractSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.TransferFrom(&_Cvaultcontract.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cvaultcontract *CvaultcontractTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.TransferFrom(&_Cvaultcontract.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cvaultcontract *CvaultcontractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cvaultcontract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cvaultcontract *CvaultcontractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.TransferOwnership(&_Cvaultcontract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cvaultcontract *CvaultcontractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cvaultcontract.Contract.TransferOwnership(&_Cvaultcontract.TransactOpts, newOwner)
}

// CvaultcontractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Cvaultcontract contract.
type CvaultcontractApprovalIterator struct {
	Event *CvaultcontractApproval // Event containing the contract specifics and raw log

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
func (it *CvaultcontractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CvaultcontractApproval)
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
		it.Event = new(CvaultcontractApproval)
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
func (it *CvaultcontractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CvaultcontractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CvaultcontractApproval represents a Approval event raised by the Cvaultcontract contract.
type CvaultcontractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Cvaultcontract *CvaultcontractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CvaultcontractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cvaultcontract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CvaultcontractApprovalIterator{contract: _Cvaultcontract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Cvaultcontract *CvaultcontractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CvaultcontractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cvaultcontract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CvaultcontractApproval)
				if err := _Cvaultcontract.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Cvaultcontract *CvaultcontractFilterer) ParseApproval(log types.Log) (*CvaultcontractApproval, error) {
	event := new(CvaultcontractApproval)
	if err := _Cvaultcontract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CvaultcontractDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the Cvaultcontract contract.
type CvaultcontractDelegateChangedIterator struct {
	Event *CvaultcontractDelegateChanged // Event containing the contract specifics and raw log

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
func (it *CvaultcontractDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CvaultcontractDelegateChanged)
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
		it.Event = new(CvaultcontractDelegateChanged)
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
func (it *CvaultcontractDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CvaultcontractDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CvaultcontractDelegateChanged represents a DelegateChanged event raised by the Cvaultcontract contract.
type CvaultcontractDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_Cvaultcontract *CvaultcontractFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*CvaultcontractDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _Cvaultcontract.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &CvaultcontractDelegateChangedIterator{contract: _Cvaultcontract.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_Cvaultcontract *CvaultcontractFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *CvaultcontractDelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _Cvaultcontract.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CvaultcontractDelegateChanged)
				if err := _Cvaultcontract.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_Cvaultcontract *CvaultcontractFilterer) ParseDelegateChanged(log types.Log) (*CvaultcontractDelegateChanged, error) {
	event := new(CvaultcontractDelegateChanged)
	if err := _Cvaultcontract.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CvaultcontractDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the Cvaultcontract contract.
type CvaultcontractDelegateVotesChangedIterator struct {
	Event *CvaultcontractDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *CvaultcontractDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CvaultcontractDelegateVotesChanged)
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
		it.Event = new(CvaultcontractDelegateVotesChanged)
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
func (it *CvaultcontractDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CvaultcontractDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CvaultcontractDelegateVotesChanged represents a DelegateVotesChanged event raised by the Cvaultcontract contract.
type CvaultcontractDelegateVotesChanged struct {
	Delegate        common.Address
	PreviousBalance *big.Int
	NewBalance      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_Cvaultcontract *CvaultcontractFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*CvaultcontractDelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _Cvaultcontract.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &CvaultcontractDelegateVotesChangedIterator{contract: _Cvaultcontract.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_Cvaultcontract *CvaultcontractFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *CvaultcontractDelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _Cvaultcontract.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CvaultcontractDelegateVotesChanged)
				if err := _Cvaultcontract.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_Cvaultcontract *CvaultcontractFilterer) ParseDelegateVotesChanged(log types.Log) (*CvaultcontractDelegateVotesChanged, error) {
	event := new(CvaultcontractDelegateVotesChanged)
	if err := _Cvaultcontract.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CvaultcontractLPTokenClaimedIterator is returned from FilterLPTokenClaimed and is used to iterate over the raw logs and unpacked data for LPTokenClaimed events raised by the Cvaultcontract contract.
type CvaultcontractLPTokenClaimedIterator struct {
	Event *CvaultcontractLPTokenClaimed // Event containing the contract specifics and raw log

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
func (it *CvaultcontractLPTokenClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CvaultcontractLPTokenClaimed)
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
		it.Event = new(CvaultcontractLPTokenClaimed)
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
func (it *CvaultcontractLPTokenClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CvaultcontractLPTokenClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CvaultcontractLPTokenClaimed represents a LPTokenClaimed event raised by the Cvaultcontract contract.
type CvaultcontractLPTokenClaimed struct {
	Dst   common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLPTokenClaimed is a free log retrieval operation binding the contract event 0x586e28f4f60b4d906fc69694ea6d7fe5c5668730ce3286d7af8eca868f3c2760.
//
// Solidity: event LPTokenClaimed(address dst, uint256 value)
func (_Cvaultcontract *CvaultcontractFilterer) FilterLPTokenClaimed(opts *bind.FilterOpts) (*CvaultcontractLPTokenClaimedIterator, error) {

	logs, sub, err := _Cvaultcontract.contract.FilterLogs(opts, "LPTokenClaimed")
	if err != nil {
		return nil, err
	}
	return &CvaultcontractLPTokenClaimedIterator{contract: _Cvaultcontract.contract, event: "LPTokenClaimed", logs: logs, sub: sub}, nil
}

// WatchLPTokenClaimed is a free log subscription operation binding the contract event 0x586e28f4f60b4d906fc69694ea6d7fe5c5668730ce3286d7af8eca868f3c2760.
//
// Solidity: event LPTokenClaimed(address dst, uint256 value)
func (_Cvaultcontract *CvaultcontractFilterer) WatchLPTokenClaimed(opts *bind.WatchOpts, sink chan<- *CvaultcontractLPTokenClaimed) (event.Subscription, error) {

	logs, sub, err := _Cvaultcontract.contract.WatchLogs(opts, "LPTokenClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CvaultcontractLPTokenClaimed)
				if err := _Cvaultcontract.contract.UnpackLog(event, "LPTokenClaimed", log); err != nil {
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

// ParseLPTokenClaimed is a log parse operation binding the contract event 0x586e28f4f60b4d906fc69694ea6d7fe5c5668730ce3286d7af8eca868f3c2760.
//
// Solidity: event LPTokenClaimed(address dst, uint256 value)
func (_Cvaultcontract *CvaultcontractFilterer) ParseLPTokenClaimed(log types.Log) (*CvaultcontractLPTokenClaimed, error) {
	event := new(CvaultcontractLPTokenClaimed)
	if err := _Cvaultcontract.contract.UnpackLog(event, "LPTokenClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CvaultcontractLiquidityAdditionIterator is returned from FilterLiquidityAddition and is used to iterate over the raw logs and unpacked data for LiquidityAddition events raised by the Cvaultcontract contract.
type CvaultcontractLiquidityAdditionIterator struct {
	Event *CvaultcontractLiquidityAddition // Event containing the contract specifics and raw log

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
func (it *CvaultcontractLiquidityAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CvaultcontractLiquidityAddition)
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
		it.Event = new(CvaultcontractLiquidityAddition)
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
func (it *CvaultcontractLiquidityAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CvaultcontractLiquidityAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CvaultcontractLiquidityAddition represents a LiquidityAddition event raised by the Cvaultcontract contract.
type CvaultcontractLiquidityAddition struct {
	Dst   common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAddition is a free log retrieval operation binding the contract event 0x20b711375edba008429d2f91787c68aa13aab7f267c346bf91be1a104d8b7b8b.
//
// Solidity: event LiquidityAddition(address indexed dst, uint256 value)
func (_Cvaultcontract *CvaultcontractFilterer) FilterLiquidityAddition(opts *bind.FilterOpts, dst []common.Address) (*CvaultcontractLiquidityAdditionIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Cvaultcontract.contract.FilterLogs(opts, "LiquidityAddition", dstRule)
	if err != nil {
		return nil, err
	}
	return &CvaultcontractLiquidityAdditionIterator{contract: _Cvaultcontract.contract, event: "LiquidityAddition", logs: logs, sub: sub}, nil
}

// WatchLiquidityAddition is a free log subscription operation binding the contract event 0x20b711375edba008429d2f91787c68aa13aab7f267c346bf91be1a104d8b7b8b.
//
// Solidity: event LiquidityAddition(address indexed dst, uint256 value)
func (_Cvaultcontract *CvaultcontractFilterer) WatchLiquidityAddition(opts *bind.WatchOpts, sink chan<- *CvaultcontractLiquidityAddition, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Cvaultcontract.contract.WatchLogs(opts, "LiquidityAddition", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CvaultcontractLiquidityAddition)
				if err := _Cvaultcontract.contract.UnpackLog(event, "LiquidityAddition", log); err != nil {
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

// ParseLiquidityAddition is a log parse operation binding the contract event 0x20b711375edba008429d2f91787c68aa13aab7f267c346bf91be1a104d8b7b8b.
//
// Solidity: event LiquidityAddition(address indexed dst, uint256 value)
func (_Cvaultcontract *CvaultcontractFilterer) ParseLiquidityAddition(log types.Log) (*CvaultcontractLiquidityAddition, error) {
	event := new(CvaultcontractLiquidityAddition)
	if err := _Cvaultcontract.contract.UnpackLog(event, "LiquidityAddition", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CvaultcontractLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the Cvaultcontract contract.
type CvaultcontractLogIterator struct {
	Event *CvaultcontractLog // Event containing the contract specifics and raw log

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
func (it *CvaultcontractLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CvaultcontractLog)
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
		it.Event = new(CvaultcontractLog)
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
func (it *CvaultcontractLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CvaultcontractLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CvaultcontractLog represents a Log event raised by the Cvaultcontract contract.
type CvaultcontractLog struct {
	Log string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0xcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab.
//
// Solidity: event Log(string log)
func (_Cvaultcontract *CvaultcontractFilterer) FilterLog(opts *bind.FilterOpts) (*CvaultcontractLogIterator, error) {

	logs, sub, err := _Cvaultcontract.contract.FilterLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return &CvaultcontractLogIterator{contract: _Cvaultcontract.contract, event: "Log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0xcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab.
//
// Solidity: event Log(string log)
func (_Cvaultcontract *CvaultcontractFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *CvaultcontractLog) (event.Subscription, error) {

	logs, sub, err := _Cvaultcontract.contract.WatchLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CvaultcontractLog)
				if err := _Cvaultcontract.contract.UnpackLog(event, "Log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0xcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab.
//
// Solidity: event Log(string log)
func (_Cvaultcontract *CvaultcontractFilterer) ParseLog(log types.Log) (*CvaultcontractLog, error) {
	event := new(CvaultcontractLog)
	if err := _Cvaultcontract.contract.UnpackLog(event, "Log", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CvaultcontractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Cvaultcontract contract.
type CvaultcontractOwnershipTransferredIterator struct {
	Event *CvaultcontractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CvaultcontractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CvaultcontractOwnershipTransferred)
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
		it.Event = new(CvaultcontractOwnershipTransferred)
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
func (it *CvaultcontractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CvaultcontractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CvaultcontractOwnershipTransferred represents a OwnershipTransferred event raised by the Cvaultcontract contract.
type CvaultcontractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cvaultcontract *CvaultcontractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CvaultcontractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cvaultcontract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CvaultcontractOwnershipTransferredIterator{contract: _Cvaultcontract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cvaultcontract *CvaultcontractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CvaultcontractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cvaultcontract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CvaultcontractOwnershipTransferred)
				if err := _Cvaultcontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Cvaultcontract *CvaultcontractFilterer) ParseOwnershipTransferred(log types.Log) (*CvaultcontractOwnershipTransferred, error) {
	event := new(CvaultcontractOwnershipTransferred)
	if err := _Cvaultcontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CvaultcontractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Cvaultcontract contract.
type CvaultcontractTransferIterator struct {
	Event *CvaultcontractTransfer // Event containing the contract specifics and raw log

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
func (it *CvaultcontractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CvaultcontractTransfer)
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
		it.Event = new(CvaultcontractTransfer)
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
func (it *CvaultcontractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CvaultcontractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CvaultcontractTransfer represents a Transfer event raised by the Cvaultcontract contract.
type CvaultcontractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Cvaultcontract *CvaultcontractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CvaultcontractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cvaultcontract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CvaultcontractTransferIterator{contract: _Cvaultcontract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Cvaultcontract *CvaultcontractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CvaultcontractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cvaultcontract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CvaultcontractTransfer)
				if err := _Cvaultcontract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Cvaultcontract *CvaultcontractFilterer) ParseTransfer(log types.Log) (*CvaultcontractTransfer, error) {
	event := new(CvaultcontractTransfer)
	if err := _Cvaultcontract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
