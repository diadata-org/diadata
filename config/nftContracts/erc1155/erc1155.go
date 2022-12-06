// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc1155

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

// Erc1155MetaData contains all meta data concerning the Erc1155 contract.
var Erc1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_youTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_justTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_winTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"JustTokenAddress\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WinTokenAddress\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"YouTokenAddress\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"burnNReceiveNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_Uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_Uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"youTokensToBurn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"justTokensToBurn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winTokensToBurn\",\"type\":\"uint256\"}],\"name\":\"createToFarm\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getTokensToBurnForNFT\",\"outputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"\",\"type\":\"uint256[3]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newBaseMetadataURI\",\"type\":\"string\"}],\"name\":\"setBaseMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Erc1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc1155MetaData.ABI instead.
var Erc1155ABI = Erc1155MetaData.ABI

// Erc1155 is an auto generated Go binding around an Ethereum contract.
type Erc1155 struct {
	Erc1155Caller     // Read-only binding to the contract
	Erc1155Transactor // Write-only binding to the contract
	Erc1155Filterer   // Log filterer for contract events
}

// Erc1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc1155Session struct {
	Contract     *Erc1155          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc1155CallerSession struct {
	Contract *Erc1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Erc1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc1155TransactorSession struct {
	Contract     *Erc1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Erc1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc1155Raw struct {
	Contract *Erc1155 // Generic contract binding to access the raw methods on
}

// Erc1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc1155CallerRaw struct {
	Contract *Erc1155Caller // Generic read-only contract binding to access the raw methods on
}

// Erc1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc1155TransactorRaw struct {
	Contract *Erc1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc1155 creates a new instance of Erc1155, bound to a specific deployed contract.
func NewErc1155(address common.Address, backend bind.ContractBackend) (*Erc1155, error) {
	contract, err := bindErc1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc1155{Erc1155Caller: Erc1155Caller{contract: contract}, Erc1155Transactor: Erc1155Transactor{contract: contract}, Erc1155Filterer: Erc1155Filterer{contract: contract}}, nil
}

// NewErc1155Caller creates a new read-only instance of Erc1155, bound to a specific deployed contract.
func NewErc1155Caller(address common.Address, caller bind.ContractCaller) (*Erc1155Caller, error) {
	contract, err := bindErc1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1155Caller{contract: contract}, nil
}

// NewErc1155Transactor creates a new write-only instance of Erc1155, bound to a specific deployed contract.
func NewErc1155Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc1155Transactor, error) {
	contract, err := bindErc1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1155Transactor{contract: contract}, nil
}

// NewErc1155Filterer creates a new log filterer instance of Erc1155, bound to a specific deployed contract.
func NewErc1155Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc1155Filterer, error) {
	contract, err := bindErc1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc1155Filterer{contract: contract}, nil
}

// bindErc1155 binds a generic wrapper to an already deployed contract.
func bindErc1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1155 *Erc1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc1155.Contract.Erc1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1155 *Erc1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1155.Contract.Erc1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1155 *Erc1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1155.Contract.Erc1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1155 *Erc1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1155 *Erc1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1155 *Erc1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1155.Contract.contract.Transact(opts, method, params...)
}

// JustTokenAddress is a free data retrieval call binding the contract method 0xb9c29221.
//
// Solidity: function JustTokenAddress() view returns(address)
func (_Erc1155 *Erc1155Caller) JustTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "JustTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JustTokenAddress is a free data retrieval call binding the contract method 0xb9c29221.
//
// Solidity: function JustTokenAddress() view returns(address)
func (_Erc1155 *Erc1155Session) JustTokenAddress() (common.Address, error) {
	return _Erc1155.Contract.JustTokenAddress(&_Erc1155.CallOpts)
}

// JustTokenAddress is a free data retrieval call binding the contract method 0xb9c29221.
//
// Solidity: function JustTokenAddress() view returns(address)
func (_Erc1155 *Erc1155CallerSession) JustTokenAddress() (common.Address, error) {
	return _Erc1155.Contract.JustTokenAddress(&_Erc1155.CallOpts)
}

// WinTokenAddress is a free data retrieval call binding the contract method 0x4766ac77.
//
// Solidity: function WinTokenAddress() view returns(address)
func (_Erc1155 *Erc1155Caller) WinTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "WinTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WinTokenAddress is a free data retrieval call binding the contract method 0x4766ac77.
//
// Solidity: function WinTokenAddress() view returns(address)
func (_Erc1155 *Erc1155Session) WinTokenAddress() (common.Address, error) {
	return _Erc1155.Contract.WinTokenAddress(&_Erc1155.CallOpts)
}

// WinTokenAddress is a free data retrieval call binding the contract method 0x4766ac77.
//
// Solidity: function WinTokenAddress() view returns(address)
func (_Erc1155 *Erc1155CallerSession) WinTokenAddress() (common.Address, error) {
	return _Erc1155.Contract.WinTokenAddress(&_Erc1155.CallOpts)
}

// YouTokenAddress is a free data retrieval call binding the contract method 0x89add238.
//
// Solidity: function YouTokenAddress() view returns(address)
func (_Erc1155 *Erc1155Caller) YouTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "YouTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YouTokenAddress is a free data retrieval call binding the contract method 0x89add238.
//
// Solidity: function YouTokenAddress() view returns(address)
func (_Erc1155 *Erc1155Session) YouTokenAddress() (common.Address, error) {
	return _Erc1155.Contract.YouTokenAddress(&_Erc1155.CallOpts)
}

// YouTokenAddress is a free data retrieval call binding the contract method 0x89add238.
//
// Solidity: function YouTokenAddress() view returns(address)
func (_Erc1155 *Erc1155CallerSession) YouTokenAddress() (common.Address, error) {
	return _Erc1155.Contract.YouTokenAddress(&_Erc1155.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Erc1155 *Erc1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Erc1155 *Erc1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Erc1155.Contract.BalanceOf(&_Erc1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Erc1155 *Erc1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Erc1155.Contract.BalanceOf(&_Erc1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Erc1155 *Erc1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Erc1155 *Erc1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Erc1155.Contract.BalanceOfBatch(&_Erc1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Erc1155 *Erc1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Erc1155.Contract.BalanceOfBatch(&_Erc1155.CallOpts, accounts, ids)
}

// GetTokensToBurnForNFT is a free data retrieval call binding the contract method 0xa53d2929.
//
// Solidity: function getTokensToBurnForNFT(uint256 id) view returns(uint256[3])
func (_Erc1155 *Erc1155Caller) GetTokensToBurnForNFT(opts *bind.CallOpts, id *big.Int) ([3]*big.Int, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "getTokensToBurnForNFT", id)

	if err != nil {
		return *new([3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([3]*big.Int)).(*[3]*big.Int)

	return out0, err

}

// GetTokensToBurnForNFT is a free data retrieval call binding the contract method 0xa53d2929.
//
// Solidity: function getTokensToBurnForNFT(uint256 id) view returns(uint256[3])
func (_Erc1155 *Erc1155Session) GetTokensToBurnForNFT(id *big.Int) ([3]*big.Int, error) {
	return _Erc1155.Contract.GetTokensToBurnForNFT(&_Erc1155.CallOpts, id)
}

// GetTokensToBurnForNFT is a free data retrieval call binding the contract method 0xa53d2929.
//
// Solidity: function getTokensToBurnForNFT(uint256 id) view returns(uint256[3])
func (_Erc1155 *Erc1155CallerSession) GetTokensToBurnForNFT(id *big.Int) ([3]*big.Int, error) {
	return _Erc1155.Contract.GetTokensToBurnForNFT(&_Erc1155.CallOpts, id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Erc1155 *Erc1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Erc1155 *Erc1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Erc1155.Contract.IsApprovedForAll(&_Erc1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Erc1155 *Erc1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Erc1155.Contract.IsApprovedForAll(&_Erc1155.CallOpts, account, operator)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Erc1155 *Erc1155Caller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Erc1155 *Erc1155Session) NominatedOwner() (common.Address, error) {
	return _Erc1155.Contract.NominatedOwner(&_Erc1155.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Erc1155 *Erc1155CallerSession) NominatedOwner() (common.Address, error) {
	return _Erc1155.Contract.NominatedOwner(&_Erc1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc1155 *Erc1155Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc1155 *Erc1155Session) Owner() (common.Address, error) {
	return _Erc1155.Contract.Owner(&_Erc1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc1155 *Erc1155CallerSession) Owner() (common.Address, error) {
	return _Erc1155.Contract.Owner(&_Erc1155.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc1155 *Erc1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc1155 *Erc1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc1155.Contract.SupportsInterface(&_Erc1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc1155 *Erc1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc1155.Contract.SupportsInterface(&_Erc1155.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_Erc1155 *Erc1155Caller) TotalSupply(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "totalSupply", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_Erc1155 *Erc1155Session) TotalSupply(_id *big.Int) (*big.Int, error) {
	return _Erc1155.Contract.TotalSupply(&_Erc1155.CallOpts, _id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_Erc1155 *Erc1155CallerSession) TotalSupply(_id *big.Int) (*big.Int, error) {
	return _Erc1155.Contract.TotalSupply(&_Erc1155.CallOpts, _id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _tokenId) view returns(string)
func (_Erc1155 *Erc1155Caller) Uri(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "uri", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _tokenId) view returns(string)
func (_Erc1155 *Erc1155Session) Uri(_tokenId *big.Int) (string, error) {
	return _Erc1155.Contract.Uri(&_Erc1155.CallOpts, _tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _tokenId) view returns(string)
func (_Erc1155 *Erc1155CallerSession) Uri(_tokenId *big.Int) (string, error) {
	return _Erc1155.Contract.Uri(&_Erc1155.CallOpts, _tokenId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Erc1155 *Erc1155Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Erc1155 *Erc1155Session) AcceptOwnership() (*types.Transaction, error) {
	return _Erc1155.Contract.AcceptOwnership(&_Erc1155.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Erc1155 *Erc1155TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Erc1155.Contract.AcceptOwnership(&_Erc1155.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_Erc1155 *Erc1155Transactor) Burn(opts *bind.TransactOpts, account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "burn", account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_Erc1155 *Erc1155Session) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Erc1155.Contract.Burn(&_Erc1155.TransactOpts, account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_Erc1155 *Erc1155TransactorSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Erc1155.Contract.Burn(&_Erc1155.TransactOpts, account, id, value)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_Erc1155 *Erc1155Transactor) BurnBatch(opts *bind.TransactOpts, account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "burnBatch", account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_Erc1155 *Erc1155Session) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _Erc1155.Contract.BurnBatch(&_Erc1155.TransactOpts, account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_Erc1155 *Erc1155TransactorSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _Erc1155.Contract.BurnBatch(&_Erc1155.TransactOpts, account, ids, values)
}

// BurnNReceiveNFT is a paid mutator transaction binding the contract method 0xf2b399ee.
//
// Solidity: function burnNReceiveNFT(uint256 _id) returns(bool)
func (_Erc1155 *Erc1155Transactor) BurnNReceiveNFT(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "burnNReceiveNFT", _id)
}

// BurnNReceiveNFT is a paid mutator transaction binding the contract method 0xf2b399ee.
//
// Solidity: function burnNReceiveNFT(uint256 _id) returns(bool)
func (_Erc1155 *Erc1155Session) BurnNReceiveNFT(_id *big.Int) (*types.Transaction, error) {
	return _Erc1155.Contract.BurnNReceiveNFT(&_Erc1155.TransactOpts, _id)
}

// BurnNReceiveNFT is a paid mutator transaction binding the contract method 0xf2b399ee.
//
// Solidity: function burnNReceiveNFT(uint256 _id) returns(bool)
func (_Erc1155 *Erc1155TransactorSession) BurnNReceiveNFT(_id *big.Int) (*types.Transaction, error) {
	return _Erc1155.Contract.BurnNReceiveNFT(&_Erc1155.TransactOpts, _id)
}

// Create is a paid mutator transaction binding the contract method 0xe38e3b24.
//
// Solidity: function create(address _to, uint256 _initialSupply, string _Uri, bytes _data) returns(uint256)
func (_Erc1155 *Erc1155Transactor) Create(opts *bind.TransactOpts, _to common.Address, _initialSupply *big.Int, _Uri string, _data []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "create", _to, _initialSupply, _Uri, _data)
}

// Create is a paid mutator transaction binding the contract method 0xe38e3b24.
//
// Solidity: function create(address _to, uint256 _initialSupply, string _Uri, bytes _data) returns(uint256)
func (_Erc1155 *Erc1155Session) Create(_to common.Address, _initialSupply *big.Int, _Uri string, _data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.Create(&_Erc1155.TransactOpts, _to, _initialSupply, _Uri, _data)
}

// Create is a paid mutator transaction binding the contract method 0xe38e3b24.
//
// Solidity: function create(address _to, uint256 _initialSupply, string _Uri, bytes _data) returns(uint256)
func (_Erc1155 *Erc1155TransactorSession) Create(_to common.Address, _initialSupply *big.Int, _Uri string, _data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.Create(&_Erc1155.TransactOpts, _to, _initialSupply, _Uri, _data)
}

// CreateToFarm is a paid mutator transaction binding the contract method 0x116ed561.
//
// Solidity: function createToFarm(uint256 _initialSupply, string _Uri, bytes _data, uint256 youTokensToBurn, uint256 justTokensToBurn, uint256 winTokensToBurn) returns(uint256)
func (_Erc1155 *Erc1155Transactor) CreateToFarm(opts *bind.TransactOpts, _initialSupply *big.Int, _Uri string, _data []byte, youTokensToBurn *big.Int, justTokensToBurn *big.Int, winTokensToBurn *big.Int) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "createToFarm", _initialSupply, _Uri, _data, youTokensToBurn, justTokensToBurn, winTokensToBurn)
}

// CreateToFarm is a paid mutator transaction binding the contract method 0x116ed561.
//
// Solidity: function createToFarm(uint256 _initialSupply, string _Uri, bytes _data, uint256 youTokensToBurn, uint256 justTokensToBurn, uint256 winTokensToBurn) returns(uint256)
func (_Erc1155 *Erc1155Session) CreateToFarm(_initialSupply *big.Int, _Uri string, _data []byte, youTokensToBurn *big.Int, justTokensToBurn *big.Int, winTokensToBurn *big.Int) (*types.Transaction, error) {
	return _Erc1155.Contract.CreateToFarm(&_Erc1155.TransactOpts, _initialSupply, _Uri, _data, youTokensToBurn, justTokensToBurn, winTokensToBurn)
}

// CreateToFarm is a paid mutator transaction binding the contract method 0x116ed561.
//
// Solidity: function createToFarm(uint256 _initialSupply, string _Uri, bytes _data, uint256 youTokensToBurn, uint256 justTokensToBurn, uint256 winTokensToBurn) returns(uint256)
func (_Erc1155 *Erc1155TransactorSession) CreateToFarm(_initialSupply *big.Int, _Uri string, _data []byte, youTokensToBurn *big.Int, justTokensToBurn *big.Int, winTokensToBurn *big.Int) (*types.Transaction, error) {
	return _Erc1155.Contract.CreateToFarm(&_Erc1155.TransactOpts, _initialSupply, _Uri, _data, youTokensToBurn, justTokensToBurn, winTokensToBurn)
}

// Mint is a paid mutator transaction binding the contract method 0x08dc9f42.
//
// Solidity: function mint(uint256 _id, uint256 _quantity, bytes _data) returns()
func (_Erc1155 *Erc1155Transactor) Mint(opts *bind.TransactOpts, _id *big.Int, _quantity *big.Int, _data []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "mint", _id, _quantity, _data)
}

// Mint is a paid mutator transaction binding the contract method 0x08dc9f42.
//
// Solidity: function mint(uint256 _id, uint256 _quantity, bytes _data) returns()
func (_Erc1155 *Erc1155Session) Mint(_id *big.Int, _quantity *big.Int, _data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.Mint(&_Erc1155.TransactOpts, _id, _quantity, _data)
}

// Mint is a paid mutator transaction binding the contract method 0x08dc9f42.
//
// Solidity: function mint(uint256 _id, uint256 _quantity, bytes _data) returns()
func (_Erc1155 *Erc1155TransactorSession) Mint(_id *big.Int, _quantity *big.Int, _data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.Mint(&_Erc1155.TransactOpts, _id, _quantity, _data)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Erc1155 *Erc1155Transactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Erc1155 *Erc1155Session) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Erc1155.Contract.NominateNewOwner(&_Erc1155.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Erc1155 *Erc1155TransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Erc1155.Contract.NominateNewOwner(&_Erc1155.TransactOpts, _owner)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Erc1155 *Erc1155Transactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Erc1155 *Erc1155Session) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.OnERC1155BatchReceived(&_Erc1155.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Erc1155 *Erc1155TransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.OnERC1155BatchReceived(&_Erc1155.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Erc1155 *Erc1155Transactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Erc1155 *Erc1155Session) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.OnERC1155Received(&_Erc1155.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Erc1155 *Erc1155TransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.OnERC1155Received(&_Erc1155.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Erc1155 *Erc1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Erc1155 *Erc1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SafeBatchTransferFrom(&_Erc1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Erc1155 *Erc1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SafeBatchTransferFrom(&_Erc1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Erc1155 *Erc1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Erc1155 *Erc1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SafeTransferFrom(&_Erc1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Erc1155 *Erc1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SafeTransferFrom(&_Erc1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc1155 *Erc1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc1155 *Erc1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc1155.Contract.SetApprovalForAll(&_Erc1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc1155 *Erc1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc1155.Contract.SetApprovalForAll(&_Erc1155.TransactOpts, operator, approved)
}

// SetBaseMetadataURI is a paid mutator transaction binding the contract method 0x7e518ec8.
//
// Solidity: function setBaseMetadataURI(string _newBaseMetadataURI) returns()
func (_Erc1155 *Erc1155Transactor) SetBaseMetadataURI(opts *bind.TransactOpts, _newBaseMetadataURI string) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setBaseMetadataURI", _newBaseMetadataURI)
}

// SetBaseMetadataURI is a paid mutator transaction binding the contract method 0x7e518ec8.
//
// Solidity: function setBaseMetadataURI(string _newBaseMetadataURI) returns()
func (_Erc1155 *Erc1155Session) SetBaseMetadataURI(_newBaseMetadataURI string) (*types.Transaction, error) {
	return _Erc1155.Contract.SetBaseMetadataURI(&_Erc1155.TransactOpts, _newBaseMetadataURI)
}

// SetBaseMetadataURI is a paid mutator transaction binding the contract method 0x7e518ec8.
//
// Solidity: function setBaseMetadataURI(string _newBaseMetadataURI) returns()
func (_Erc1155 *Erc1155TransactorSession) SetBaseMetadataURI(_newBaseMetadataURI string) (*types.Transaction, error) {
	return _Erc1155.Contract.SetBaseMetadataURI(&_Erc1155.TransactOpts, _newBaseMetadataURI)
}

// Erc1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Erc1155 contract.
type Erc1155ApprovalForAllIterator struct {
	Event *Erc1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Erc1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155ApprovalForAll)
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
		it.Event = new(Erc1155ApprovalForAll)
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
func (it *Erc1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155ApprovalForAll represents a ApprovalForAll event raised by the Erc1155 contract.
type Erc1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Erc1155 *Erc1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*Erc1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155ApprovalForAllIterator{contract: _Erc1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Erc1155 *Erc1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Erc1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155ApprovalForAll)
				if err := _Erc1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Erc1155 *Erc1155Filterer) ParseApprovalForAll(log types.Log) (*Erc1155ApprovalForAll, error) {
	event := new(Erc1155ApprovalForAll)
	if err := _Erc1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155OwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Erc1155 contract.
type Erc1155OwnerChangedIterator struct {
	Event *Erc1155OwnerChanged // Event containing the contract specifics and raw log

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
func (it *Erc1155OwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155OwnerChanged)
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
		it.Event = new(Erc1155OwnerChanged)
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
func (it *Erc1155OwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155OwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155OwnerChanged represents a OwnerChanged event raised by the Erc1155 contract.
type Erc1155OwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Erc1155 *Erc1155Filterer) FilterOwnerChanged(opts *bind.FilterOpts) (*Erc1155OwnerChangedIterator, error) {

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &Erc1155OwnerChangedIterator{contract: _Erc1155.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Erc1155 *Erc1155Filterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *Erc1155OwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155OwnerChanged)
				if err := _Erc1155.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Erc1155 *Erc1155Filterer) ParseOwnerChanged(log types.Log) (*Erc1155OwnerChanged, error) {
	event := new(Erc1155OwnerChanged)
	if err := _Erc1155.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155OwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the Erc1155 contract.
type Erc1155OwnerNominatedIterator struct {
	Event *Erc1155OwnerNominated // Event containing the contract specifics and raw log

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
func (it *Erc1155OwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155OwnerNominated)
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
		it.Event = new(Erc1155OwnerNominated)
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
func (it *Erc1155OwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155OwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155OwnerNominated represents a OwnerNominated event raised by the Erc1155 contract.
type Erc1155OwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Erc1155 *Erc1155Filterer) FilterOwnerNominated(opts *bind.FilterOpts) (*Erc1155OwnerNominatedIterator, error) {

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &Erc1155OwnerNominatedIterator{contract: _Erc1155.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Erc1155 *Erc1155Filterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *Erc1155OwnerNominated) (event.Subscription, error) {

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155OwnerNominated)
				if err := _Erc1155.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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

// ParseOwnerNominated is a log parse operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Erc1155 *Erc1155Filterer) ParseOwnerNominated(log types.Log) (*Erc1155OwnerNominated, error) {
	event := new(Erc1155OwnerNominated)
	if err := _Erc1155.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Erc1155 contract.
type Erc1155TransferBatchIterator struct {
	Event *Erc1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *Erc1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155TransferBatch)
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
		it.Event = new(Erc1155TransferBatch)
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
func (it *Erc1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155TransferBatch represents a TransferBatch event raised by the Erc1155 contract.
type Erc1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Erc1155 *Erc1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Erc1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155TransferBatchIterator{contract: _Erc1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Erc1155 *Erc1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *Erc1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155TransferBatch)
				if err := _Erc1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Erc1155 *Erc1155Filterer) ParseTransferBatch(log types.Log) (*Erc1155TransferBatch, error) {
	event := new(Erc1155TransferBatch)
	if err := _Erc1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Erc1155 contract.
type Erc1155TransferSingleIterator struct {
	Event *Erc1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *Erc1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155TransferSingle)
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
		it.Event = new(Erc1155TransferSingle)
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
func (it *Erc1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155TransferSingle represents a TransferSingle event raised by the Erc1155 contract.
type Erc1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Erc1155 *Erc1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Erc1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155TransferSingleIterator{contract: _Erc1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Erc1155 *Erc1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *Erc1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155TransferSingle)
				if err := _Erc1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Erc1155 *Erc1155Filterer) ParseTransferSingle(log types.Log) (*Erc1155TransferSingle, error) {
	event := new(Erc1155TransferSingle)
	if err := _Erc1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Erc1155 contract.
type Erc1155URIIterator struct {
	Event *Erc1155URI // Event containing the contract specifics and raw log

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
func (it *Erc1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155URI)
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
		it.Event = new(Erc1155URI)
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
func (it *Erc1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155URI represents a URI event raised by the Erc1155 contract.
type Erc1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Erc1155 *Erc1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*Erc1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155URIIterator{contract: _Erc1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Erc1155 *Erc1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *Erc1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155URI)
				if err := _Erc1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Erc1155 *Erc1155Filterer) ParseURI(log types.Log) (*Erc1155URI, error) {
	event := new(Erc1155URI)
	if err := _Erc1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
