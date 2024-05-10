// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reEthFairValueService

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

// ReEthFairValueServiceMetaData contains all meta data concerning the ReEthFairValueService contract.
var ReEthFairValueServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_real\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Minter__NotVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Minter__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"real\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"setNewVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161092f38038061092f833981810160405281019061003191906101ec565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16148061009657505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b156100cd576040517f708a3da500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505061022a565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61018082610157565b9050919050565b61019081610176565b811461019a575f80fd5b50565b5f815190506101ab81610187565b92915050565b5f6101bb82610157565b9050919050565b6101cb816101b1565b81146101d5575f80fd5b50565b5f815190506101e6816101c2565b92915050565b5f806040838503121561020257610201610153565b5b5f61020f8582860161019d565b9250506020610220858286016101d8565b9150509250929050565b6106f8806102375f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c806340c10f19146100645780634b94f50e146100805780639dc29fac1461009e578063c954f076146100ba578063e7b77f70146100d8578063fbfa77cf146100f4575b5f80fd5b61007e6004803603810190610079919061056a565b610112565b005b610088610223565b60405161009591906105b7565b60405180910390f35b6100b860048036038101906100b3919061056a565b6102b7565b005b6100c26103c8565b6040516100cf91906105df565b60405180910390f35b6100f260048036038101906100ed91906105f8565b6103eb565b005b6100fc6104b4565b6040516101099190610643565b60405180910390f35b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610198576040517f2f1bb37a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1983836040518363ffffffff1660e01b81526004016101f292919061065c565b5f604051808303815f87803b158015610209575f80fd5b505af115801561021b573d5f803e3d5ffd5b505050505050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166328a795766040518163ffffffff1660e01b8152600401602060405180830381865afa15801561028e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102b29190610697565b905090565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461033d576040517f2f1bb37a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639dc29fac83836040518363ffffffff1660e01b815260040161039792919061065c565b5f604051808303815f87803b1580156103ae575f80fd5b505af11580156103c0573d5f803e3d5ffd5b505050505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610471576040517f2f1bb37a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610506826104dd565b9050919050565b610516816104fc565b8114610520575f80fd5b50565b5f813590506105318161050d565b92915050565b5f819050919050565b61054981610537565b8114610553575f80fd5b50565b5f8135905061056481610540565b92915050565b5f80604083850312156105805761057f6104d9565b5b5f61058d85828601610523565b925050602061059e85828601610556565b9150509250929050565b6105b181610537565b82525050565b5f6020820190506105ca5f8301846105a8565b92915050565b6105d9816104fc565b82525050565b5f6020820190506105f25f8301846105d0565b92915050565b5f6020828403121561060d5761060c6104d9565b5b5f61061a84828501610523565b91505092915050565b5f61062d826104dd565b9050919050565b61063d81610623565b82525050565b5f6020820190506106565f830184610634565b92915050565b5f60408201905061066f5f8301856105d0565b61067c60208301846105a8565b9392505050565b5f8151905061069181610540565b92915050565b5f602082840312156106ac576106ab6104d9565b5b5f6106b984828501610683565b9150509291505056fea264697066735822122048e29b0ad36c9f37b84efaa88fd6e41cf93a120edce60673b56ed37304fb533a64736f6c63430008150033",
}

// ReEthFairValueServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use ReEthFairValueServiceMetaData.ABI instead.
var ReEthFairValueServiceABI = ReEthFairValueServiceMetaData.ABI

// ReEthFairValueServiceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReEthFairValueServiceMetaData.Bin instead.
var ReEthFairValueServiceBin = ReEthFairValueServiceMetaData.Bin

// DeployReEthFairValueService deploys a new Ethereum contract, binding an instance of ReEthFairValueService to it.
func DeployReEthFairValueService(auth *bind.TransactOpts, backend bind.ContractBackend, _real common.Address, _vault common.Address) (common.Address, *types.Transaction, *ReEthFairValueService, error) {
	parsed, err := ReEthFairValueServiceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReEthFairValueServiceBin), backend, _real, _vault)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReEthFairValueService{ReEthFairValueServiceCaller: ReEthFairValueServiceCaller{contract: contract}, ReEthFairValueServiceTransactor: ReEthFairValueServiceTransactor{contract: contract}, ReEthFairValueServiceFilterer: ReEthFairValueServiceFilterer{contract: contract}}, nil
}

// ReEthFairValueService is an auto generated Go binding around an Ethereum contract.
type ReEthFairValueService struct {
	ReEthFairValueServiceCaller     // Read-only binding to the contract
	ReEthFairValueServiceTransactor // Write-only binding to the contract
	ReEthFairValueServiceFilterer   // Log filterer for contract events
}

// ReEthFairValueServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReEthFairValueServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReEthFairValueServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReEthFairValueServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReEthFairValueServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReEthFairValueServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReEthFairValueServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReEthFairValueServiceSession struct {
	Contract     *ReEthFairValueService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ReEthFairValueServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReEthFairValueServiceCallerSession struct {
	Contract *ReEthFairValueServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ReEthFairValueServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReEthFairValueServiceTransactorSession struct {
	Contract     *ReEthFairValueServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ReEthFairValueServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReEthFairValueServiceRaw struct {
	Contract *ReEthFairValueService // Generic contract binding to access the raw methods on
}

// ReEthFairValueServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReEthFairValueServiceCallerRaw struct {
	Contract *ReEthFairValueServiceCaller // Generic read-only contract binding to access the raw methods on
}

// ReEthFairValueServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReEthFairValueServiceTransactorRaw struct {
	Contract *ReEthFairValueServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReEthFairValueService creates a new instance of ReEthFairValueService, bound to a specific deployed contract.
func NewReEthFairValueService(address common.Address, backend bind.ContractBackend) (*ReEthFairValueService, error) {
	contract, err := bindReEthFairValueService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReEthFairValueService{ReEthFairValueServiceCaller: ReEthFairValueServiceCaller{contract: contract}, ReEthFairValueServiceTransactor: ReEthFairValueServiceTransactor{contract: contract}, ReEthFairValueServiceFilterer: ReEthFairValueServiceFilterer{contract: contract}}, nil
}

// NewReEthFairValueServiceCaller creates a new read-only instance of ReEthFairValueService, bound to a specific deployed contract.
func NewReEthFairValueServiceCaller(address common.Address, caller bind.ContractCaller) (*ReEthFairValueServiceCaller, error) {
	contract, err := bindReEthFairValueService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReEthFairValueServiceCaller{contract: contract}, nil
}

// NewReEthFairValueServiceTransactor creates a new write-only instance of ReEthFairValueService, bound to a specific deployed contract.
func NewReEthFairValueServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*ReEthFairValueServiceTransactor, error) {
	contract, err := bindReEthFairValueService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReEthFairValueServiceTransactor{contract: contract}, nil
}

// NewReEthFairValueServiceFilterer creates a new log filterer instance of ReEthFairValueService, bound to a specific deployed contract.
func NewReEthFairValueServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*ReEthFairValueServiceFilterer, error) {
	contract, err := bindReEthFairValueService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReEthFairValueServiceFilterer{contract: contract}, nil
}

// bindReEthFairValueService binds a generic wrapper to an already deployed contract.
func bindReEthFairValueService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReEthFairValueServiceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReEthFairValueService *ReEthFairValueServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReEthFairValueService.Contract.ReEthFairValueServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReEthFairValueService *ReEthFairValueServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReEthFairValueService.Contract.ReEthFairValueServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReEthFairValueService *ReEthFairValueServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReEthFairValueService.Contract.ReEthFairValueServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReEthFairValueService *ReEthFairValueServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReEthFairValueService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReEthFairValueService *ReEthFairValueServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReEthFairValueService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReEthFairValueService *ReEthFairValueServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReEthFairValueService.Contract.contract.Transact(opts, method, params...)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0x4b94f50e.
//
// Solidity: function getTokenPrice() view returns(uint256 price)
func (_ReEthFairValueService *ReEthFairValueServiceCaller) GetTokenPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReEthFairValueService.contract.Call(opts, &out, "getTokenPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenPrice is a free data retrieval call binding the contract method 0x4b94f50e.
//
// Solidity: function getTokenPrice() view returns(uint256 price)
func (_ReEthFairValueService *ReEthFairValueServiceSession) GetTokenPrice() (*big.Int, error) {
	return _ReEthFairValueService.Contract.GetTokenPrice(&_ReEthFairValueService.CallOpts)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0x4b94f50e.
//
// Solidity: function getTokenPrice() view returns(uint256 price)
func (_ReEthFairValueService *ReEthFairValueServiceCallerSession) GetTokenPrice() (*big.Int, error) {
	return _ReEthFairValueService.Contract.GetTokenPrice(&_ReEthFairValueService.CallOpts)
}

// Real is a free data retrieval call binding the contract method 0xc954f076.
//
// Solidity: function real() view returns(address)
func (_ReEthFairValueService *ReEthFairValueServiceCaller) Real(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReEthFairValueService.contract.Call(opts, &out, "real")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Real is a free data retrieval call binding the contract method 0xc954f076.
//
// Solidity: function real() view returns(address)
func (_ReEthFairValueService *ReEthFairValueServiceSession) Real() (common.Address, error) {
	return _ReEthFairValueService.Contract.Real(&_ReEthFairValueService.CallOpts)
}

// Real is a free data retrieval call binding the contract method 0xc954f076.
//
// Solidity: function real() view returns(address)
func (_ReEthFairValueService *ReEthFairValueServiceCallerSession) Real() (common.Address, error) {
	return _ReEthFairValueService.Contract.Real(&_ReEthFairValueService.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ReEthFairValueService *ReEthFairValueServiceCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReEthFairValueService.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ReEthFairValueService *ReEthFairValueServiceSession) Vault() (common.Address, error) {
	return _ReEthFairValueService.Contract.Vault(&_ReEthFairValueService.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ReEthFairValueService *ReEthFairValueServiceCallerSession) Vault() (common.Address, error) {
	return _ReEthFairValueService.Contract.Vault(&_ReEthFairValueService.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_ReEthFairValueService *ReEthFairValueServiceTransactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ReEthFairValueService.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_ReEthFairValueService *ReEthFairValueServiceSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ReEthFairValueService.Contract.Burn(&_ReEthFairValueService.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_ReEthFairValueService *ReEthFairValueServiceTransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ReEthFairValueService.Contract.Burn(&_ReEthFairValueService.TransactOpts, _from, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_ReEthFairValueService *ReEthFairValueServiceTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ReEthFairValueService.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_ReEthFairValueService *ReEthFairValueServiceSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ReEthFairValueService.Contract.Mint(&_ReEthFairValueService.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_ReEthFairValueService *ReEthFairValueServiceTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ReEthFairValueService.Contract.Mint(&_ReEthFairValueService.TransactOpts, _to, _amount)
}

// SetNewVault is a paid mutator transaction binding the contract method 0xe7b77f70.
//
// Solidity: function setNewVault(address _vault) returns()
func (_ReEthFairValueService *ReEthFairValueServiceTransactor) SetNewVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _ReEthFairValueService.contract.Transact(opts, "setNewVault", _vault)
}

// SetNewVault is a paid mutator transaction binding the contract method 0xe7b77f70.
//
// Solidity: function setNewVault(address _vault) returns()
func (_ReEthFairValueService *ReEthFairValueServiceSession) SetNewVault(_vault common.Address) (*types.Transaction, error) {
	return _ReEthFairValueService.Contract.SetNewVault(&_ReEthFairValueService.TransactOpts, _vault)
}

// SetNewVault is a paid mutator transaction binding the contract method 0xe7b77f70.
//
// Solidity: function setNewVault(address _vault) returns()
func (_ReEthFairValueService *ReEthFairValueServiceTransactorSession) SetNewVault(_vault common.Address) (*types.Transaction, error) {
	return _ReEthFairValueService.Contract.SetNewVault(&_ReEthFairValueService.TransactOpts, _vault)
}
