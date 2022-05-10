// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package looksrare

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

// OrderTypesMakerOrder is an auto generated low-level Go binding around an user-defined struct.
type OrderTypesMakerOrder struct {
	IsOrderAsk         bool
	Signer             common.Address
	Collection         common.Address
	Price              *big.Int
	TokenId            *big.Int
	Amount             *big.Int
	Strategy           common.Address
	Currency           common.Address
	Nonce              *big.Int
	StartTime          *big.Int
	EndTime            *big.Int
	MinPercentageToAsk *big.Int
	Params             []byte
	V                  uint8
	R                  [32]byte
	S                  [32]byte
}

// OrderTypesTakerOrder is an auto generated low-level Go binding around an user-defined struct.
type OrderTypesTakerOrder struct {
	IsOrderAsk         bool
	Taker              common.Address
	Price              *big.Int
	TokenId            *big.Int
	MinPercentageToAsk *big.Int
	Params             []byte
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_currencyManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_executionManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_royaltyFeeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_protocolFeeRecipient\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMinNonce\",\"type\":\"uint256\"}],\"name\":\"CancelAllOrders\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"CancelMultipleOrders\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currencyManager\",\"type\":\"address\"}],\"name\":\"NewCurrencyManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executionManager\",\"type\":\"address\"}],\"name\":\"NewExecutionManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolFeeRecipient\",\"type\":\"address\"}],\"name\":\"NewProtocolFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"royaltyFeeManager\",\"type\":\"address\"}],\"name\":\"NewRoyaltyFeeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transferSelectorNFT\",\"type\":\"address\"}],\"name\":\"NewTransferSelectorNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RoyaltyPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"TakerAsk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"TakerBid\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minNonce\",\"type\":\"uint256\"}],\"name\":\"cancelAllOrdersForSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"cancelMultipleMakerOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currencyManager\",\"outputs\":[{\"internalType\":\"contractICurrencyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executionManager\",\"outputs\":[{\"internalType\":\"contractIExecutionManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"}],\"name\":\"isUserOrderNonceExecutedOrCancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"internalType\":\"structOrderTypes.TakerOrder\",\"name\":\"takerBid\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structOrderTypes.MakerOrder\",\"name\":\"makerAsk\",\"type\":\"tuple\"}],\"name\":\"matchAskWithTakerBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"internalType\":\"structOrderTypes.TakerOrder\",\"name\":\"takerBid\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structOrderTypes.MakerOrder\",\"name\":\"makerAsk\",\"type\":\"tuple\"}],\"name\":\"matchAskWithTakerBidUsingETHAndWETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"internalType\":\"structOrderTypes.TakerOrder\",\"name\":\"takerAsk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structOrderTypes.MakerOrder\",\"name\":\"makerBid\",\"type\":\"tuple\"}],\"name\":\"matchBidWithTakerAsk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyFeeManager\",\"outputs\":[{\"internalType\":\"contractIRoyaltyFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferSelectorNFT\",\"outputs\":[{\"internalType\":\"contractITransferSelectorNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_currencyManager\",\"type\":\"address\"}],\"name\":\"updateCurrencyManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executionManager\",\"type\":\"address\"}],\"name\":\"updateExecutionManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_protocolFeeRecipient\",\"type\":\"address\"}],\"name\":\"updateProtocolFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyFeeManager\",\"type\":\"address\"}],\"name\":\"updateRoyaltyFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_transferSelectorNFT\",\"type\":\"address\"}],\"name\":\"updateTransferSelectorNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMinOrderNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Contract *ContractCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Contract *ContractSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Contract.Contract.DOMAINSEPARATOR(&_Contract.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Contract *ContractCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Contract.Contract.DOMAINSEPARATOR(&_Contract.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Contract *ContractCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Contract *ContractSession) WETH() (common.Address, error) {
	return _Contract.Contract.WETH(&_Contract.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Contract *ContractCallerSession) WETH() (common.Address, error) {
	return _Contract.Contract.WETH(&_Contract.CallOpts)
}

// CurrencyManager is a free data retrieval call binding the contract method 0x0f747d74.
//
// Solidity: function currencyManager() view returns(address)
func (_Contract *ContractCaller) CurrencyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currencyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrencyManager is a free data retrieval call binding the contract method 0x0f747d74.
//
// Solidity: function currencyManager() view returns(address)
func (_Contract *ContractSession) CurrencyManager() (common.Address, error) {
	return _Contract.Contract.CurrencyManager(&_Contract.CallOpts)
}

// CurrencyManager is a free data retrieval call binding the contract method 0x0f747d74.
//
// Solidity: function currencyManager() view returns(address)
func (_Contract *ContractCallerSession) CurrencyManager() (common.Address, error) {
	return _Contract.Contract.CurrencyManager(&_Contract.CallOpts)
}

// ExecutionManager is a free data retrieval call binding the contract method 0x483abb9f.
//
// Solidity: function executionManager() view returns(address)
func (_Contract *ContractCaller) ExecutionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "executionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutionManager is a free data retrieval call binding the contract method 0x483abb9f.
//
// Solidity: function executionManager() view returns(address)
func (_Contract *ContractSession) ExecutionManager() (common.Address, error) {
	return _Contract.Contract.ExecutionManager(&_Contract.CallOpts)
}

// ExecutionManager is a free data retrieval call binding the contract method 0x483abb9f.
//
// Solidity: function executionManager() view returns(address)
func (_Contract *ContractCallerSession) ExecutionManager() (common.Address, error) {
	return _Contract.Contract.ExecutionManager(&_Contract.CallOpts)
}

// IsUserOrderNonceExecutedOrCancelled is a free data retrieval call binding the contract method 0x31e27e27.
//
// Solidity: function isUserOrderNonceExecutedOrCancelled(address user, uint256 orderNonce) view returns(bool)
func (_Contract *ContractCaller) IsUserOrderNonceExecutedOrCancelled(opts *bind.CallOpts, user common.Address, orderNonce *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isUserOrderNonceExecutedOrCancelled", user, orderNonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserOrderNonceExecutedOrCancelled is a free data retrieval call binding the contract method 0x31e27e27.
//
// Solidity: function isUserOrderNonceExecutedOrCancelled(address user, uint256 orderNonce) view returns(bool)
func (_Contract *ContractSession) IsUserOrderNonceExecutedOrCancelled(user common.Address, orderNonce *big.Int) (bool, error) {
	return _Contract.Contract.IsUserOrderNonceExecutedOrCancelled(&_Contract.CallOpts, user, orderNonce)
}

// IsUserOrderNonceExecutedOrCancelled is a free data retrieval call binding the contract method 0x31e27e27.
//
// Solidity: function isUserOrderNonceExecutedOrCancelled(address user, uint256 orderNonce) view returns(bool)
func (_Contract *ContractCallerSession) IsUserOrderNonceExecutedOrCancelled(user common.Address, orderNonce *big.Int) (bool, error) {
	return _Contract.Contract.IsUserOrderNonceExecutedOrCancelled(&_Contract.CallOpts, user, orderNonce)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_Contract *ContractCaller) ProtocolFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "protocolFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_Contract *ContractSession) ProtocolFeeRecipient() (common.Address, error) {
	return _Contract.Contract.ProtocolFeeRecipient(&_Contract.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_Contract *ContractCallerSession) ProtocolFeeRecipient() (common.Address, error) {
	return _Contract.Contract.ProtocolFeeRecipient(&_Contract.CallOpts)
}

// RoyaltyFeeManager is a free data retrieval call binding the contract method 0x87e4401f.
//
// Solidity: function royaltyFeeManager() view returns(address)
func (_Contract *ContractCaller) RoyaltyFeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "royaltyFeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyFeeManager is a free data retrieval call binding the contract method 0x87e4401f.
//
// Solidity: function royaltyFeeManager() view returns(address)
func (_Contract *ContractSession) RoyaltyFeeManager() (common.Address, error) {
	return _Contract.Contract.RoyaltyFeeManager(&_Contract.CallOpts)
}

// RoyaltyFeeManager is a free data retrieval call binding the contract method 0x87e4401f.
//
// Solidity: function royaltyFeeManager() view returns(address)
func (_Contract *ContractCallerSession) RoyaltyFeeManager() (common.Address, error) {
	return _Contract.Contract.RoyaltyFeeManager(&_Contract.CallOpts)
}

// TransferSelectorNFT is a free data retrieval call binding the contract method 0x5e14f68e.
//
// Solidity: function transferSelectorNFT() view returns(address)
func (_Contract *ContractCaller) TransferSelectorNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "transferSelectorNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TransferSelectorNFT is a free data retrieval call binding the contract method 0x5e14f68e.
//
// Solidity: function transferSelectorNFT() view returns(address)
func (_Contract *ContractSession) TransferSelectorNFT() (common.Address, error) {
	return _Contract.Contract.TransferSelectorNFT(&_Contract.CallOpts)
}

// TransferSelectorNFT is a free data retrieval call binding the contract method 0x5e14f68e.
//
// Solidity: function transferSelectorNFT() view returns(address)
func (_Contract *ContractCallerSession) TransferSelectorNFT() (common.Address, error) {
	return _Contract.Contract.TransferSelectorNFT(&_Contract.CallOpts)
}

// UserMinOrderNonce is a free data retrieval call binding the contract method 0x4266581e.
//
// Solidity: function userMinOrderNonce(address ) view returns(uint256)
func (_Contract *ContractCaller) UserMinOrderNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "userMinOrderNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMinOrderNonce is a free data retrieval call binding the contract method 0x4266581e.
//
// Solidity: function userMinOrderNonce(address ) view returns(uint256)
func (_Contract *ContractSession) UserMinOrderNonce(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.UserMinOrderNonce(&_Contract.CallOpts, arg0)
}

// UserMinOrderNonce is a free data retrieval call binding the contract method 0x4266581e.
//
// Solidity: function userMinOrderNonce(address ) view returns(uint256)
func (_Contract *ContractCallerSession) UserMinOrderNonce(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.UserMinOrderNonce(&_Contract.CallOpts, arg0)
}

// CancelAllOrdersForSender is a paid mutator transaction binding the contract method 0xcbd2ec65.
//
// Solidity: function cancelAllOrdersForSender(uint256 minNonce) returns()
func (_Contract *ContractTransactor) CancelAllOrdersForSender(opts *bind.TransactOpts, minNonce *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "cancelAllOrdersForSender", minNonce)
}

// CancelAllOrdersForSender is a paid mutator transaction binding the contract method 0xcbd2ec65.
//
// Solidity: function cancelAllOrdersForSender(uint256 minNonce) returns()
func (_Contract *ContractSession) CancelAllOrdersForSender(minNonce *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelAllOrdersForSender(&_Contract.TransactOpts, minNonce)
}

// CancelAllOrdersForSender is a paid mutator transaction binding the contract method 0xcbd2ec65.
//
// Solidity: function cancelAllOrdersForSender(uint256 minNonce) returns()
func (_Contract *ContractTransactorSession) CancelAllOrdersForSender(minNonce *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelAllOrdersForSender(&_Contract.TransactOpts, minNonce)
}

// CancelMultipleMakerOrders is a paid mutator transaction binding the contract method 0x9e53a69a.
//
// Solidity: function cancelMultipleMakerOrders(uint256[] orderNonces) returns()
func (_Contract *ContractTransactor) CancelMultipleMakerOrders(opts *bind.TransactOpts, orderNonces []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "cancelMultipleMakerOrders", orderNonces)
}

// CancelMultipleMakerOrders is a paid mutator transaction binding the contract method 0x9e53a69a.
//
// Solidity: function cancelMultipleMakerOrders(uint256[] orderNonces) returns()
func (_Contract *ContractSession) CancelMultipleMakerOrders(orderNonces []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelMultipleMakerOrders(&_Contract.TransactOpts, orderNonces)
}

// CancelMultipleMakerOrders is a paid mutator transaction binding the contract method 0x9e53a69a.
//
// Solidity: function cancelMultipleMakerOrders(uint256[] orderNonces) returns()
func (_Contract *ContractTransactorSession) CancelMultipleMakerOrders(orderNonces []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelMultipleMakerOrders(&_Contract.TransactOpts, orderNonces)
}

// MatchAskWithTakerBid is a paid mutator transaction binding the contract method 0x38e29209.
//
// Solidity: function matchAskWithTakerBid((bool,address,uint256,uint256,uint256,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerAsk) returns()
func (_Contract *ContractTransactor) MatchAskWithTakerBid(opts *bind.TransactOpts, takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "matchAskWithTakerBid", takerBid, makerAsk)
}

// MatchAskWithTakerBid is a paid mutator transaction binding the contract method 0x38e29209.
//
// Solidity: function matchAskWithTakerBid((bool,address,uint256,uint256,uint256,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerAsk) returns()
func (_Contract *ContractSession) MatchAskWithTakerBid(takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Contract.Contract.MatchAskWithTakerBid(&_Contract.TransactOpts, takerBid, makerAsk)
}

// MatchAskWithTakerBid is a paid mutator transaction binding the contract method 0x38e29209.
//
// Solidity: function matchAskWithTakerBid((bool,address,uint256,uint256,uint256,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerAsk) returns()
func (_Contract *ContractTransactorSession) MatchAskWithTakerBid(takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Contract.Contract.MatchAskWithTakerBid(&_Contract.TransactOpts, takerBid, makerAsk)
}

// MatchAskWithTakerBidUsingETHAndWETH is a paid mutator transaction binding the contract method 0xb4e4b296.
//
// Solidity: function matchAskWithTakerBidUsingETHAndWETH((bool,address,uint256,uint256,uint256,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerAsk) payable returns()
func (_Contract *ContractTransactor) MatchAskWithTakerBidUsingETHAndWETH(opts *bind.TransactOpts, takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "matchAskWithTakerBidUsingETHAndWETH", takerBid, makerAsk)
}

// MatchAskWithTakerBidUsingETHAndWETH is a paid mutator transaction binding the contract method 0xb4e4b296.
//
// Solidity: function matchAskWithTakerBidUsingETHAndWETH((bool,address,uint256,uint256,uint256,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerAsk) payable returns()
func (_Contract *ContractSession) MatchAskWithTakerBidUsingETHAndWETH(takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Contract.Contract.MatchAskWithTakerBidUsingETHAndWETH(&_Contract.TransactOpts, takerBid, makerAsk)
}

// MatchAskWithTakerBidUsingETHAndWETH is a paid mutator transaction binding the contract method 0xb4e4b296.
//
// Solidity: function matchAskWithTakerBidUsingETHAndWETH((bool,address,uint256,uint256,uint256,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerAsk) payable returns()
func (_Contract *ContractTransactorSession) MatchAskWithTakerBidUsingETHAndWETH(takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Contract.Contract.MatchAskWithTakerBidUsingETHAndWETH(&_Contract.TransactOpts, takerBid, makerAsk)
}

// MatchBidWithTakerAsk is a paid mutator transaction binding the contract method 0x3b6d032e.
//
// Solidity: function matchBidWithTakerAsk((bool,address,uint256,uint256,uint256,bytes) takerAsk, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerBid) returns()
func (_Contract *ContractTransactor) MatchBidWithTakerAsk(opts *bind.TransactOpts, takerAsk OrderTypesTakerOrder, makerBid OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "matchBidWithTakerAsk", takerAsk, makerBid)
}

// MatchBidWithTakerAsk is a paid mutator transaction binding the contract method 0x3b6d032e.
//
// Solidity: function matchBidWithTakerAsk((bool,address,uint256,uint256,uint256,bytes) takerAsk, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerBid) returns()
func (_Contract *ContractSession) MatchBidWithTakerAsk(takerAsk OrderTypesTakerOrder, makerBid OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Contract.Contract.MatchBidWithTakerAsk(&_Contract.TransactOpts, takerAsk, makerBid)
}

// MatchBidWithTakerAsk is a paid mutator transaction binding the contract method 0x3b6d032e.
//
// Solidity: function matchBidWithTakerAsk((bool,address,uint256,uint256,uint256,bytes) takerAsk, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerBid) returns()
func (_Contract *ContractTransactorSession) MatchBidWithTakerAsk(takerAsk OrderTypesTakerOrder, makerBid OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Contract.Contract.MatchBidWithTakerAsk(&_Contract.TransactOpts, takerAsk, makerBid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// UpdateCurrencyManager is a paid mutator transaction binding the contract method 0x5ce052d7.
//
// Solidity: function updateCurrencyManager(address _currencyManager) returns()
func (_Contract *ContractTransactor) UpdateCurrencyManager(opts *bind.TransactOpts, _currencyManager common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateCurrencyManager", _currencyManager)
}

// UpdateCurrencyManager is a paid mutator transaction binding the contract method 0x5ce052d7.
//
// Solidity: function updateCurrencyManager(address _currencyManager) returns()
func (_Contract *ContractSession) UpdateCurrencyManager(_currencyManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateCurrencyManager(&_Contract.TransactOpts, _currencyManager)
}

// UpdateCurrencyManager is a paid mutator transaction binding the contract method 0x5ce052d7.
//
// Solidity: function updateCurrencyManager(address _currencyManager) returns()
func (_Contract *ContractTransactorSession) UpdateCurrencyManager(_currencyManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateCurrencyManager(&_Contract.TransactOpts, _currencyManager)
}

// UpdateExecutionManager is a paid mutator transaction binding the contract method 0xd4ff41dc.
//
// Solidity: function updateExecutionManager(address _executionManager) returns()
func (_Contract *ContractTransactor) UpdateExecutionManager(opts *bind.TransactOpts, _executionManager common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateExecutionManager", _executionManager)
}

// UpdateExecutionManager is a paid mutator transaction binding the contract method 0xd4ff41dc.
//
// Solidity: function updateExecutionManager(address _executionManager) returns()
func (_Contract *ContractSession) UpdateExecutionManager(_executionManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateExecutionManager(&_Contract.TransactOpts, _executionManager)
}

// UpdateExecutionManager is a paid mutator transaction binding the contract method 0xd4ff41dc.
//
// Solidity: function updateExecutionManager(address _executionManager) returns()
func (_Contract *ContractTransactorSession) UpdateExecutionManager(_executionManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateExecutionManager(&_Contract.TransactOpts, _executionManager)
}

// UpdateProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x1df47f80.
//
// Solidity: function updateProtocolFeeRecipient(address _protocolFeeRecipient) returns()
func (_Contract *ContractTransactor) UpdateProtocolFeeRecipient(opts *bind.TransactOpts, _protocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateProtocolFeeRecipient", _protocolFeeRecipient)
}

// UpdateProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x1df47f80.
//
// Solidity: function updateProtocolFeeRecipient(address _protocolFeeRecipient) returns()
func (_Contract *ContractSession) UpdateProtocolFeeRecipient(_protocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateProtocolFeeRecipient(&_Contract.TransactOpts, _protocolFeeRecipient)
}

// UpdateProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x1df47f80.
//
// Solidity: function updateProtocolFeeRecipient(address _protocolFeeRecipient) returns()
func (_Contract *ContractTransactorSession) UpdateProtocolFeeRecipient(_protocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateProtocolFeeRecipient(&_Contract.TransactOpts, _protocolFeeRecipient)
}

// UpdateRoyaltyFeeManager is a paid mutator transaction binding the contract method 0xc5498769.
//
// Solidity: function updateRoyaltyFeeManager(address _royaltyFeeManager) returns()
func (_Contract *ContractTransactor) UpdateRoyaltyFeeManager(opts *bind.TransactOpts, _royaltyFeeManager common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateRoyaltyFeeManager", _royaltyFeeManager)
}

// UpdateRoyaltyFeeManager is a paid mutator transaction binding the contract method 0xc5498769.
//
// Solidity: function updateRoyaltyFeeManager(address _royaltyFeeManager) returns()
func (_Contract *ContractSession) UpdateRoyaltyFeeManager(_royaltyFeeManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateRoyaltyFeeManager(&_Contract.TransactOpts, _royaltyFeeManager)
}

// UpdateRoyaltyFeeManager is a paid mutator transaction binding the contract method 0xc5498769.
//
// Solidity: function updateRoyaltyFeeManager(address _royaltyFeeManager) returns()
func (_Contract *ContractTransactorSession) UpdateRoyaltyFeeManager(_royaltyFeeManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateRoyaltyFeeManager(&_Contract.TransactOpts, _royaltyFeeManager)
}

// UpdateTransferSelectorNFT is a paid mutator transaction binding the contract method 0xf75ff53f.
//
// Solidity: function updateTransferSelectorNFT(address _transferSelectorNFT) returns()
func (_Contract *ContractTransactor) UpdateTransferSelectorNFT(opts *bind.TransactOpts, _transferSelectorNFT common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateTransferSelectorNFT", _transferSelectorNFT)
}

// UpdateTransferSelectorNFT is a paid mutator transaction binding the contract method 0xf75ff53f.
//
// Solidity: function updateTransferSelectorNFT(address _transferSelectorNFT) returns()
func (_Contract *ContractSession) UpdateTransferSelectorNFT(_transferSelectorNFT common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateTransferSelectorNFT(&_Contract.TransactOpts, _transferSelectorNFT)
}

// UpdateTransferSelectorNFT is a paid mutator transaction binding the contract method 0xf75ff53f.
//
// Solidity: function updateTransferSelectorNFT(address _transferSelectorNFT) returns()
func (_Contract *ContractTransactorSession) UpdateTransferSelectorNFT(_transferSelectorNFT common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateTransferSelectorNFT(&_Contract.TransactOpts, _transferSelectorNFT)
}

// ContractCancelAllOrdersIterator is returned from FilterCancelAllOrders and is used to iterate over the raw logs and unpacked data for CancelAllOrders events raised by the Contract contract.
type ContractCancelAllOrdersIterator struct {
	Event *ContractCancelAllOrders // Event containing the contract specifics and raw log

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
func (it *ContractCancelAllOrdersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCancelAllOrders)
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
		it.Event = new(ContractCancelAllOrders)
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
func (it *ContractCancelAllOrdersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCancelAllOrdersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCancelAllOrders represents a CancelAllOrders event raised by the Contract contract.
type ContractCancelAllOrders struct {
	User        common.Address
	NewMinNonce *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancelAllOrders is a free log retrieval operation binding the contract event 0x1e7178d84f0b0825c65795cd62e7972809ad3aac6917843aaec596161b2c0a97.
//
// Solidity: event CancelAllOrders(address indexed user, uint256 newMinNonce)
func (_Contract *ContractFilterer) FilterCancelAllOrders(opts *bind.FilterOpts, user []common.Address) (*ContractCancelAllOrdersIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CancelAllOrders", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractCancelAllOrdersIterator{contract: _Contract.contract, event: "CancelAllOrders", logs: logs, sub: sub}, nil
}

// WatchCancelAllOrders is a free log subscription operation binding the contract event 0x1e7178d84f0b0825c65795cd62e7972809ad3aac6917843aaec596161b2c0a97.
//
// Solidity: event CancelAllOrders(address indexed user, uint256 newMinNonce)
func (_Contract *ContractFilterer) WatchCancelAllOrders(opts *bind.WatchOpts, sink chan<- *ContractCancelAllOrders, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CancelAllOrders", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCancelAllOrders)
				if err := _Contract.contract.UnpackLog(event, "CancelAllOrders", log); err != nil {
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

// ParseCancelAllOrders is a log parse operation binding the contract event 0x1e7178d84f0b0825c65795cd62e7972809ad3aac6917843aaec596161b2c0a97.
//
// Solidity: event CancelAllOrders(address indexed user, uint256 newMinNonce)
func (_Contract *ContractFilterer) ParseCancelAllOrders(log types.Log) (*ContractCancelAllOrders, error) {
	event := new(ContractCancelAllOrders)
	if err := _Contract.contract.UnpackLog(event, "CancelAllOrders", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCancelMultipleOrdersIterator is returned from FilterCancelMultipleOrders and is used to iterate over the raw logs and unpacked data for CancelMultipleOrders events raised by the Contract contract.
type ContractCancelMultipleOrdersIterator struct {
	Event *ContractCancelMultipleOrders // Event containing the contract specifics and raw log

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
func (it *ContractCancelMultipleOrdersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCancelMultipleOrders)
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
		it.Event = new(ContractCancelMultipleOrders)
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
func (it *ContractCancelMultipleOrdersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCancelMultipleOrdersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCancelMultipleOrders represents a CancelMultipleOrders event raised by the Contract contract.
type ContractCancelMultipleOrders struct {
	User        common.Address
	OrderNonces []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancelMultipleOrders is a free log retrieval operation binding the contract event 0xfa0ae5d80fe3763c880a3839fab0294171a6f730d1f82c4cd5392c6f67b41732.
//
// Solidity: event CancelMultipleOrders(address indexed user, uint256[] orderNonces)
func (_Contract *ContractFilterer) FilterCancelMultipleOrders(opts *bind.FilterOpts, user []common.Address) (*ContractCancelMultipleOrdersIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CancelMultipleOrders", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractCancelMultipleOrdersIterator{contract: _Contract.contract, event: "CancelMultipleOrders", logs: logs, sub: sub}, nil
}

// WatchCancelMultipleOrders is a free log subscription operation binding the contract event 0xfa0ae5d80fe3763c880a3839fab0294171a6f730d1f82c4cd5392c6f67b41732.
//
// Solidity: event CancelMultipleOrders(address indexed user, uint256[] orderNonces)
func (_Contract *ContractFilterer) WatchCancelMultipleOrders(opts *bind.WatchOpts, sink chan<- *ContractCancelMultipleOrders, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CancelMultipleOrders", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCancelMultipleOrders)
				if err := _Contract.contract.UnpackLog(event, "CancelMultipleOrders", log); err != nil {
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

// ParseCancelMultipleOrders is a log parse operation binding the contract event 0xfa0ae5d80fe3763c880a3839fab0294171a6f730d1f82c4cd5392c6f67b41732.
//
// Solidity: event CancelMultipleOrders(address indexed user, uint256[] orderNonces)
func (_Contract *ContractFilterer) ParseCancelMultipleOrders(log types.Log) (*ContractCancelMultipleOrders, error) {
	event := new(ContractCancelMultipleOrders)
	if err := _Contract.contract.UnpackLog(event, "CancelMultipleOrders", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewCurrencyManagerIterator is returned from FilterNewCurrencyManager and is used to iterate over the raw logs and unpacked data for NewCurrencyManager events raised by the Contract contract.
type ContractNewCurrencyManagerIterator struct {
	Event *ContractNewCurrencyManager // Event containing the contract specifics and raw log

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
func (it *ContractNewCurrencyManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewCurrencyManager)
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
		it.Event = new(ContractNewCurrencyManager)
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
func (it *ContractNewCurrencyManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewCurrencyManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewCurrencyManager represents a NewCurrencyManager event raised by the Contract contract.
type ContractNewCurrencyManager struct {
	CurrencyManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewCurrencyManager is a free log retrieval operation binding the contract event 0xb4f5db40df3aced29e88a4babbc3b46e305e07d34098525d18b1497056e63838.
//
// Solidity: event NewCurrencyManager(address indexed currencyManager)
func (_Contract *ContractFilterer) FilterNewCurrencyManager(opts *bind.FilterOpts, currencyManager []common.Address) (*ContractNewCurrencyManagerIterator, error) {

	var currencyManagerRule []interface{}
	for _, currencyManagerItem := range currencyManager {
		currencyManagerRule = append(currencyManagerRule, currencyManagerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewCurrencyManager", currencyManagerRule)
	if err != nil {
		return nil, err
	}
	return &ContractNewCurrencyManagerIterator{contract: _Contract.contract, event: "NewCurrencyManager", logs: logs, sub: sub}, nil
}

// WatchNewCurrencyManager is a free log subscription operation binding the contract event 0xb4f5db40df3aced29e88a4babbc3b46e305e07d34098525d18b1497056e63838.
//
// Solidity: event NewCurrencyManager(address indexed currencyManager)
func (_Contract *ContractFilterer) WatchNewCurrencyManager(opts *bind.WatchOpts, sink chan<- *ContractNewCurrencyManager, currencyManager []common.Address) (event.Subscription, error) {

	var currencyManagerRule []interface{}
	for _, currencyManagerItem := range currencyManager {
		currencyManagerRule = append(currencyManagerRule, currencyManagerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewCurrencyManager", currencyManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewCurrencyManager)
				if err := _Contract.contract.UnpackLog(event, "NewCurrencyManager", log); err != nil {
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

// ParseNewCurrencyManager is a log parse operation binding the contract event 0xb4f5db40df3aced29e88a4babbc3b46e305e07d34098525d18b1497056e63838.
//
// Solidity: event NewCurrencyManager(address indexed currencyManager)
func (_Contract *ContractFilterer) ParseNewCurrencyManager(log types.Log) (*ContractNewCurrencyManager, error) {
	event := new(ContractNewCurrencyManager)
	if err := _Contract.contract.UnpackLog(event, "NewCurrencyManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewExecutionManagerIterator is returned from FilterNewExecutionManager and is used to iterate over the raw logs and unpacked data for NewExecutionManager events raised by the Contract contract.
type ContractNewExecutionManagerIterator struct {
	Event *ContractNewExecutionManager // Event containing the contract specifics and raw log

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
func (it *ContractNewExecutionManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewExecutionManager)
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
		it.Event = new(ContractNewExecutionManager)
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
func (it *ContractNewExecutionManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewExecutionManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewExecutionManager represents a NewExecutionManager event raised by the Contract contract.
type ContractNewExecutionManager struct {
	ExecutionManager common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewExecutionManager is a free log retrieval operation binding the contract event 0x36e2a376eabc3bc60cb88f29c288f53e36874a95a7f407330ab4f166b0905698.
//
// Solidity: event NewExecutionManager(address indexed executionManager)
func (_Contract *ContractFilterer) FilterNewExecutionManager(opts *bind.FilterOpts, executionManager []common.Address) (*ContractNewExecutionManagerIterator, error) {

	var executionManagerRule []interface{}
	for _, executionManagerItem := range executionManager {
		executionManagerRule = append(executionManagerRule, executionManagerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewExecutionManager", executionManagerRule)
	if err != nil {
		return nil, err
	}
	return &ContractNewExecutionManagerIterator{contract: _Contract.contract, event: "NewExecutionManager", logs: logs, sub: sub}, nil
}

// WatchNewExecutionManager is a free log subscription operation binding the contract event 0x36e2a376eabc3bc60cb88f29c288f53e36874a95a7f407330ab4f166b0905698.
//
// Solidity: event NewExecutionManager(address indexed executionManager)
func (_Contract *ContractFilterer) WatchNewExecutionManager(opts *bind.WatchOpts, sink chan<- *ContractNewExecutionManager, executionManager []common.Address) (event.Subscription, error) {

	var executionManagerRule []interface{}
	for _, executionManagerItem := range executionManager {
		executionManagerRule = append(executionManagerRule, executionManagerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewExecutionManager", executionManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewExecutionManager)
				if err := _Contract.contract.UnpackLog(event, "NewExecutionManager", log); err != nil {
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

// ParseNewExecutionManager is a log parse operation binding the contract event 0x36e2a376eabc3bc60cb88f29c288f53e36874a95a7f407330ab4f166b0905698.
//
// Solidity: event NewExecutionManager(address indexed executionManager)
func (_Contract *ContractFilterer) ParseNewExecutionManager(log types.Log) (*ContractNewExecutionManager, error) {
	event := new(ContractNewExecutionManager)
	if err := _Contract.contract.UnpackLog(event, "NewExecutionManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewProtocolFeeRecipientIterator is returned from FilterNewProtocolFeeRecipient and is used to iterate over the raw logs and unpacked data for NewProtocolFeeRecipient events raised by the Contract contract.
type ContractNewProtocolFeeRecipientIterator struct {
	Event *ContractNewProtocolFeeRecipient // Event containing the contract specifics and raw log

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
func (it *ContractNewProtocolFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewProtocolFeeRecipient)
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
		it.Event = new(ContractNewProtocolFeeRecipient)
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
func (it *ContractNewProtocolFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewProtocolFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewProtocolFeeRecipient represents a NewProtocolFeeRecipient event raised by the Contract contract.
type ContractNewProtocolFeeRecipient struct {
	ProtocolFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewProtocolFeeRecipient is a free log retrieval operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address indexed protocolFeeRecipient)
func (_Contract *ContractFilterer) FilterNewProtocolFeeRecipient(opts *bind.FilterOpts, protocolFeeRecipient []common.Address) (*ContractNewProtocolFeeRecipientIterator, error) {

	var protocolFeeRecipientRule []interface{}
	for _, protocolFeeRecipientItem := range protocolFeeRecipient {
		protocolFeeRecipientRule = append(protocolFeeRecipientRule, protocolFeeRecipientItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewProtocolFeeRecipient", protocolFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &ContractNewProtocolFeeRecipientIterator{contract: _Contract.contract, event: "NewProtocolFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchNewProtocolFeeRecipient is a free log subscription operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address indexed protocolFeeRecipient)
func (_Contract *ContractFilterer) WatchNewProtocolFeeRecipient(opts *bind.WatchOpts, sink chan<- *ContractNewProtocolFeeRecipient, protocolFeeRecipient []common.Address) (event.Subscription, error) {

	var protocolFeeRecipientRule []interface{}
	for _, protocolFeeRecipientItem := range protocolFeeRecipient {
		protocolFeeRecipientRule = append(protocolFeeRecipientRule, protocolFeeRecipientItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewProtocolFeeRecipient", protocolFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewProtocolFeeRecipient)
				if err := _Contract.contract.UnpackLog(event, "NewProtocolFeeRecipient", log); err != nil {
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

// ParseNewProtocolFeeRecipient is a log parse operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address indexed protocolFeeRecipient)
func (_Contract *ContractFilterer) ParseNewProtocolFeeRecipient(log types.Log) (*ContractNewProtocolFeeRecipient, error) {
	event := new(ContractNewProtocolFeeRecipient)
	if err := _Contract.contract.UnpackLog(event, "NewProtocolFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewRoyaltyFeeManagerIterator is returned from FilterNewRoyaltyFeeManager and is used to iterate over the raw logs and unpacked data for NewRoyaltyFeeManager events raised by the Contract contract.
type ContractNewRoyaltyFeeManagerIterator struct {
	Event *ContractNewRoyaltyFeeManager // Event containing the contract specifics and raw log

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
func (it *ContractNewRoyaltyFeeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewRoyaltyFeeManager)
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
		it.Event = new(ContractNewRoyaltyFeeManager)
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
func (it *ContractNewRoyaltyFeeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewRoyaltyFeeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewRoyaltyFeeManager represents a NewRoyaltyFeeManager event raised by the Contract contract.
type ContractNewRoyaltyFeeManager struct {
	RoyaltyFeeManager common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewRoyaltyFeeManager is a free log retrieval operation binding the contract event 0x80e3874461ebbd918ac3e81da0a92e5e51387d70f337237c9123e48d20e5a508.
//
// Solidity: event NewRoyaltyFeeManager(address indexed royaltyFeeManager)
func (_Contract *ContractFilterer) FilterNewRoyaltyFeeManager(opts *bind.FilterOpts, royaltyFeeManager []common.Address) (*ContractNewRoyaltyFeeManagerIterator, error) {

	var royaltyFeeManagerRule []interface{}
	for _, royaltyFeeManagerItem := range royaltyFeeManager {
		royaltyFeeManagerRule = append(royaltyFeeManagerRule, royaltyFeeManagerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewRoyaltyFeeManager", royaltyFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return &ContractNewRoyaltyFeeManagerIterator{contract: _Contract.contract, event: "NewRoyaltyFeeManager", logs: logs, sub: sub}, nil
}

// WatchNewRoyaltyFeeManager is a free log subscription operation binding the contract event 0x80e3874461ebbd918ac3e81da0a92e5e51387d70f337237c9123e48d20e5a508.
//
// Solidity: event NewRoyaltyFeeManager(address indexed royaltyFeeManager)
func (_Contract *ContractFilterer) WatchNewRoyaltyFeeManager(opts *bind.WatchOpts, sink chan<- *ContractNewRoyaltyFeeManager, royaltyFeeManager []common.Address) (event.Subscription, error) {

	var royaltyFeeManagerRule []interface{}
	for _, royaltyFeeManagerItem := range royaltyFeeManager {
		royaltyFeeManagerRule = append(royaltyFeeManagerRule, royaltyFeeManagerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewRoyaltyFeeManager", royaltyFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewRoyaltyFeeManager)
				if err := _Contract.contract.UnpackLog(event, "NewRoyaltyFeeManager", log); err != nil {
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

// ParseNewRoyaltyFeeManager is a log parse operation binding the contract event 0x80e3874461ebbd918ac3e81da0a92e5e51387d70f337237c9123e48d20e5a508.
//
// Solidity: event NewRoyaltyFeeManager(address indexed royaltyFeeManager)
func (_Contract *ContractFilterer) ParseNewRoyaltyFeeManager(log types.Log) (*ContractNewRoyaltyFeeManager, error) {
	event := new(ContractNewRoyaltyFeeManager)
	if err := _Contract.contract.UnpackLog(event, "NewRoyaltyFeeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewTransferSelectorNFTIterator is returned from FilterNewTransferSelectorNFT and is used to iterate over the raw logs and unpacked data for NewTransferSelectorNFT events raised by the Contract contract.
type ContractNewTransferSelectorNFTIterator struct {
	Event *ContractNewTransferSelectorNFT // Event containing the contract specifics and raw log

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
func (it *ContractNewTransferSelectorNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewTransferSelectorNFT)
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
		it.Event = new(ContractNewTransferSelectorNFT)
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
func (it *ContractNewTransferSelectorNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewTransferSelectorNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewTransferSelectorNFT represents a NewTransferSelectorNFT event raised by the Contract contract.
type ContractNewTransferSelectorNFT struct {
	TransferSelectorNFT common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewTransferSelectorNFT is a free log retrieval operation binding the contract event 0x205d78ab41afe80bd6b6aaa5d7599d5300ff8690da3ab1302c1b552f7baf7d8c.
//
// Solidity: event NewTransferSelectorNFT(address indexed transferSelectorNFT)
func (_Contract *ContractFilterer) FilterNewTransferSelectorNFT(opts *bind.FilterOpts, transferSelectorNFT []common.Address) (*ContractNewTransferSelectorNFTIterator, error) {

	var transferSelectorNFTRule []interface{}
	for _, transferSelectorNFTItem := range transferSelectorNFT {
		transferSelectorNFTRule = append(transferSelectorNFTRule, transferSelectorNFTItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewTransferSelectorNFT", transferSelectorNFTRule)
	if err != nil {
		return nil, err
	}
	return &ContractNewTransferSelectorNFTIterator{contract: _Contract.contract, event: "NewTransferSelectorNFT", logs: logs, sub: sub}, nil
}

// WatchNewTransferSelectorNFT is a free log subscription operation binding the contract event 0x205d78ab41afe80bd6b6aaa5d7599d5300ff8690da3ab1302c1b552f7baf7d8c.
//
// Solidity: event NewTransferSelectorNFT(address indexed transferSelectorNFT)
func (_Contract *ContractFilterer) WatchNewTransferSelectorNFT(opts *bind.WatchOpts, sink chan<- *ContractNewTransferSelectorNFT, transferSelectorNFT []common.Address) (event.Subscription, error) {

	var transferSelectorNFTRule []interface{}
	for _, transferSelectorNFTItem := range transferSelectorNFT {
		transferSelectorNFTRule = append(transferSelectorNFTRule, transferSelectorNFTItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewTransferSelectorNFT", transferSelectorNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewTransferSelectorNFT)
				if err := _Contract.contract.UnpackLog(event, "NewTransferSelectorNFT", log); err != nil {
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

// ParseNewTransferSelectorNFT is a log parse operation binding the contract event 0x205d78ab41afe80bd6b6aaa5d7599d5300ff8690da3ab1302c1b552f7baf7d8c.
//
// Solidity: event NewTransferSelectorNFT(address indexed transferSelectorNFT)
func (_Contract *ContractFilterer) ParseNewTransferSelectorNFT(log types.Log) (*ContractNewTransferSelectorNFT, error) {
	event := new(ContractNewTransferSelectorNFT)
	if err := _Contract.contract.UnpackLog(event, "NewTransferSelectorNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRoyaltyPaymentIterator is returned from FilterRoyaltyPayment and is used to iterate over the raw logs and unpacked data for RoyaltyPayment events raised by the Contract contract.
type ContractRoyaltyPaymentIterator struct {
	Event *ContractRoyaltyPayment // Event containing the contract specifics and raw log

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
func (it *ContractRoyaltyPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRoyaltyPayment)
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
		it.Event = new(ContractRoyaltyPayment)
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
func (it *ContractRoyaltyPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRoyaltyPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRoyaltyPayment represents a RoyaltyPayment event raised by the Contract contract.
type ContractRoyaltyPayment struct {
	Collection       common.Address
	TokenId          *big.Int
	RoyaltyRecipient common.Address
	Currency         common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyPayment is a free log retrieval operation binding the contract event 0x27c4f0403323142b599832f26acd21c74a9e5b809f2215726e244a4ac588cd7d.
//
// Solidity: event RoyaltyPayment(address indexed collection, uint256 indexed tokenId, address indexed royaltyRecipient, address currency, uint256 amount)
func (_Contract *ContractFilterer) FilterRoyaltyPayment(opts *bind.FilterOpts, collection []common.Address, tokenId []*big.Int, royaltyRecipient []common.Address) (*ContractRoyaltyPaymentIterator, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RoyaltyPayment", collectionRule, tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &ContractRoyaltyPaymentIterator{contract: _Contract.contract, event: "RoyaltyPayment", logs: logs, sub: sub}, nil
}

// WatchRoyaltyPayment is a free log subscription operation binding the contract event 0x27c4f0403323142b599832f26acd21c74a9e5b809f2215726e244a4ac588cd7d.
//
// Solidity: event RoyaltyPayment(address indexed collection, uint256 indexed tokenId, address indexed royaltyRecipient, address currency, uint256 amount)
func (_Contract *ContractFilterer) WatchRoyaltyPayment(opts *bind.WatchOpts, sink chan<- *ContractRoyaltyPayment, collection []common.Address, tokenId []*big.Int, royaltyRecipient []common.Address) (event.Subscription, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RoyaltyPayment", collectionRule, tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRoyaltyPayment)
				if err := _Contract.contract.UnpackLog(event, "RoyaltyPayment", log); err != nil {
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

// ParseRoyaltyPayment is a log parse operation binding the contract event 0x27c4f0403323142b599832f26acd21c74a9e5b809f2215726e244a4ac588cd7d.
//
// Solidity: event RoyaltyPayment(address indexed collection, uint256 indexed tokenId, address indexed royaltyRecipient, address currency, uint256 amount)
func (_Contract *ContractFilterer) ParseRoyaltyPayment(log types.Log) (*ContractRoyaltyPayment, error) {
	event := new(ContractRoyaltyPayment)
	if err := _Contract.contract.UnpackLog(event, "RoyaltyPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTakerAskIterator is returned from FilterTakerAsk and is used to iterate over the raw logs and unpacked data for TakerAsk events raised by the Contract contract.
type ContractTakerAskIterator struct {
	Event *ContractTakerAsk // Event containing the contract specifics and raw log

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
func (it *ContractTakerAskIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTakerAsk)
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
		it.Event = new(ContractTakerAsk)
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
func (it *ContractTakerAskIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTakerAskIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTakerAsk represents a TakerAsk event raised by the Contract contract.
type ContractTakerAsk struct {
	OrderHash  [32]byte
	OrderNonce *big.Int
	Taker      common.Address
	Maker      common.Address
	Strategy   common.Address
	Currency   common.Address
	Collection common.Address
	TokenId    *big.Int
	Amount     *big.Int
	Price      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTakerAsk is a free log retrieval operation binding the contract event 0x68cd251d4d267c6e2034ff0088b990352b97b2002c0476587d0c4da889c11330.
//
// Solidity: event TakerAsk(bytes32 orderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_Contract *ContractFilterer) FilterTakerAsk(opts *bind.FilterOpts, taker []common.Address, maker []common.Address, strategy []common.Address) (*ContractTakerAskIterator, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TakerAsk", takerRule, makerRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &ContractTakerAskIterator{contract: _Contract.contract, event: "TakerAsk", logs: logs, sub: sub}, nil
}

// WatchTakerAsk is a free log subscription operation binding the contract event 0x68cd251d4d267c6e2034ff0088b990352b97b2002c0476587d0c4da889c11330.
//
// Solidity: event TakerAsk(bytes32 orderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_Contract *ContractFilterer) WatchTakerAsk(opts *bind.WatchOpts, sink chan<- *ContractTakerAsk, taker []common.Address, maker []common.Address, strategy []common.Address) (event.Subscription, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TakerAsk", takerRule, makerRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTakerAsk)
				if err := _Contract.contract.UnpackLog(event, "TakerAsk", log); err != nil {
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

// ParseTakerAsk is a log parse operation binding the contract event 0x68cd251d4d267c6e2034ff0088b990352b97b2002c0476587d0c4da889c11330.
//
// Solidity: event TakerAsk(bytes32 orderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_Contract *ContractFilterer) ParseTakerAsk(log types.Log) (*ContractTakerAsk, error) {
	event := new(ContractTakerAsk)
	if err := _Contract.contract.UnpackLog(event, "TakerAsk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTakerBidIterator is returned from FilterTakerBid and is used to iterate over the raw logs and unpacked data for TakerBid events raised by the Contract contract.
type ContractTakerBidIterator struct {
	Event *ContractTakerBid // Event containing the contract specifics and raw log

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
func (it *ContractTakerBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTakerBid)
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
		it.Event = new(ContractTakerBid)
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
func (it *ContractTakerBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTakerBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTakerBid represents a TakerBid event raised by the Contract contract.
type ContractTakerBid struct {
	OrderHash  [32]byte
	OrderNonce *big.Int
	Taker      common.Address
	Maker      common.Address
	Strategy   common.Address
	Currency   common.Address
	Collection common.Address
	TokenId    *big.Int
	Amount     *big.Int
	Price      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTakerBid is a free log retrieval operation binding the contract event 0x95fb6205e23ff6bda16a2d1dba56b9ad7c783f67c96fa149785052f47696f2be.
//
// Solidity: event TakerBid(bytes32 orderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_Contract *ContractFilterer) FilterTakerBid(opts *bind.FilterOpts, taker []common.Address, maker []common.Address, strategy []common.Address) (*ContractTakerBidIterator, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TakerBid", takerRule, makerRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &ContractTakerBidIterator{contract: _Contract.contract, event: "TakerBid", logs: logs, sub: sub}, nil
}

// WatchTakerBid is a free log subscription operation binding the contract event 0x95fb6205e23ff6bda16a2d1dba56b9ad7c783f67c96fa149785052f47696f2be.
//
// Solidity: event TakerBid(bytes32 orderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_Contract *ContractFilterer) WatchTakerBid(opts *bind.WatchOpts, sink chan<- *ContractTakerBid, taker []common.Address, maker []common.Address, strategy []common.Address) (event.Subscription, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TakerBid", takerRule, makerRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTakerBid)
				if err := _Contract.contract.UnpackLog(event, "TakerBid", log); err != nil {
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

// ParseTakerBid is a log parse operation binding the contract event 0x95fb6205e23ff6bda16a2d1dba56b9ad7c783f67c96fa149785052f47696f2be.
//
// Solidity: event TakerBid(bytes32 orderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_Contract *ContractFilterer) ParseTakerBid(log types.Log) (*ContractTakerBid, error) {
	event := new(ContractTakerBid)
	if err := _Contract.contract.UnpackLog(event, "TakerBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
