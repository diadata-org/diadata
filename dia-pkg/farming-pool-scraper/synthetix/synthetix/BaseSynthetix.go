// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package synthetixcontract

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

// AddressResolverABI is the input ABI used to generate the binding from.
const AddressResolverABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getSynth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"names\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"destinations\",\"type\":\"address[]\"}],\"name\":\"importAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"repository\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"requireAndGetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AddressResolverFuncSigs maps the 4-byte function signature to its string representation.
var AddressResolverFuncSigs = map[string]string{
	"79ba5097": "acceptOwnership()",
	"21f8a721": "getAddress(bytes32)",
	"51456061": "getSynth(bytes32)",
	"ab0b8f77": "importAddresses(bytes32[],address[])",
	"1627540c": "nominateNewOwner(address)",
	"53a47bb7": "nominatedOwner()",
	"8da5cb5b": "owner()",
	"187f7935": "repository(bytes32)",
	"dacb2d01": "requireAndGetAddress(bytes32,string)",
}

// AddressResolverBin is the compiled bytecode used for deploying new contracts.
var AddressResolverBin = "0x608060405234801561001057600080fd5b506040516108233803806108238339818101604052602081101561003357600080fd5b5051806001600160a01b038116610091576040805162461bcd60e51b815260206004820152601960248201527f4f776e657220616464726573732063616e6e6f74206265203000000000000000604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b038316908117825560408051928352602083019190915280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a15050610729806100fa6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806353a47bb71161006657806353a47bb71461013357806379ba50971461013b5780638da5cb5b14610143578063ab0b8f771461014b578063dacb2d011461020d57610093565b80631627540c14610098578063187f7935146100c057806321f8a721146100f95780635145606114610116575b600080fd5b6100be600480360360208110156100ae57600080fd5b50356001600160a01b0316610284565b005b6100dd600480360360208110156100d657600080fd5b50356102e0565b604080516001600160a01b039092168252519081900360200190f35b6100dd6004803603602081101561010f57600080fd5b50356102fb565b6100dd6004803603602081101561012c57600080fd5b5035610316565b6100dd61041e565b6100be61042d565b6100dd6104e9565b6100be6004803603604081101561016157600080fd5b81019060208101813564010000000081111561017c57600080fd5b82018360208201111561018e57600080fd5b803590602001918460208302840111640100000000831117156101b057600080fd5b9193909290916020810190356401000000008111156101ce57600080fd5b8201836020820111156101e057600080fd5b8035906020019184602083028401116401000000008311171561020257600080fd5b5090925090506104f8565b6100dd6004803603604081101561022357600080fd5b8135919081019060408101602082013564010000000081111561024557600080fd5b82018360208201111561025757600080fd5b8035906020019184600183028401116401000000008311171561027957600080fd5b5090925090506105d1565b61028c610645565b600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce229181900360200190a150565b6002602052600090815260409020546001600160a01b031681565b6000908152600260205260409020546001600160a01b031690565b6524b9b9bab2b960d11b600090815260026020527f0651498423135bdecab48e2d306f14d560a72d49179b71410fd95b5d25ce349a546001600160a01b0316806103a7576040805162461bcd60e51b815260206004820152601a60248201527f43616e6e6f742066696e64204973737565722061646472657373000000000000604482015290519081900360640190fd5b806001600160a01b03166332608039846040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156103eb57600080fd5b505afa1580156103ff573d6000803e3d6000fd5b505050506040513d602081101561041557600080fd5b50519392505050565b6001546001600160a01b031681565b6001546001600160a01b031633146104765760405162461bcd60e51b81526004018080602001828103825260358152602001806106916035913960400191505060405180910390fd5b600054600154604080516001600160a01b03938416815292909116602083015280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a160018054600080546001600160a01b03199081166001600160a01b03841617909155169055565b6000546001600160a01b031681565b610500610645565b828114610554576040805162461bcd60e51b815260206004820152601860248201527f496e707574206c656e67746873206d757374206d617463680000000000000000604482015290519081900360640190fd5b60005b838110156105ca5782828281811061056b57fe5b905060200201356001600160a01b03166002600087878581811061058b57fe5b6020908102929092013583525081019190915260400160002080546001600160a01b0319166001600160a01b0392909216919091179055600101610557565b5050505050565b6000838152600260205260408120546001600160a01b031683838261063a5760405162461bcd60e51b815260206004820190815260248201839052908190604401848480828437600083820152604051601f909101601f19169092018290039550909350505050fd5b509095945050505050565b6000546001600160a01b0316331461068e5760405162461bcd60e51b815260040180806020018281038252602f8152602001806106c6602f913960400191505060405180910390fd5b56fe596f75206d757374206265206e6f6d696e61746564206265666f726520796f752063616e20616363657074206f776e6572736869704f6e6c792074686520636f6e7472616374206f776e6572206d617920706572666f726d207468697320616374696f6ea265627a7a7231582056e1afc9479c1034013ea281d282470486c881b3abfb5de7330ddef1bbe45e9f64736f6c63430005100032"

// DeployAddressResolver deploys a new Ethereum contract, binding an instance of AddressResolver to it.
func DeployAddressResolver(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *AddressResolver, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressResolverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressResolverBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressResolver{AddressResolverCaller: AddressResolverCaller{contract: contract}, AddressResolverTransactor: AddressResolverTransactor{contract: contract}, AddressResolverFilterer: AddressResolverFilterer{contract: contract}}, nil
}

// AddressResolver is an auto generated Go binding around an Ethereum contract.
type AddressResolver struct {
	AddressResolverCaller     // Read-only binding to the contract
	AddressResolverTransactor // Write-only binding to the contract
	AddressResolverFilterer   // Log filterer for contract events
}

// AddressResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressResolverSession struct {
	Contract     *AddressResolver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressResolverCallerSession struct {
	Contract *AddressResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AddressResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressResolverTransactorSession struct {
	Contract     *AddressResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AddressResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressResolverRaw struct {
	Contract *AddressResolver // Generic contract binding to access the raw methods on
}

// AddressResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressResolverCallerRaw struct {
	Contract *AddressResolverCaller // Generic read-only contract binding to access the raw methods on
}

// AddressResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressResolverTransactorRaw struct {
	Contract *AddressResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressResolver creates a new instance of AddressResolver, bound to a specific deployed contract.
func NewAddressResolver(address common.Address, backend bind.ContractBackend) (*AddressResolver, error) {
	contract, err := bindAddressResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressResolver{AddressResolverCaller: AddressResolverCaller{contract: contract}, AddressResolverTransactor: AddressResolverTransactor{contract: contract}, AddressResolverFilterer: AddressResolverFilterer{contract: contract}}, nil
}

// NewAddressResolverCaller creates a new read-only instance of AddressResolver, bound to a specific deployed contract.
func NewAddressResolverCaller(address common.Address, caller bind.ContractCaller) (*AddressResolverCaller, error) {
	contract, err := bindAddressResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressResolverCaller{contract: contract}, nil
}

// NewAddressResolverTransactor creates a new write-only instance of AddressResolver, bound to a specific deployed contract.
func NewAddressResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressResolverTransactor, error) {
	contract, err := bindAddressResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressResolverTransactor{contract: contract}, nil
}

// NewAddressResolverFilterer creates a new log filterer instance of AddressResolver, bound to a specific deployed contract.
func NewAddressResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressResolverFilterer, error) {
	contract, err := bindAddressResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressResolverFilterer{contract: contract}, nil
}

// bindAddressResolver binds a generic wrapper to an already deployed contract.
func bindAddressResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressResolver *AddressResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressResolver.Contract.AddressResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressResolver *AddressResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressResolver.Contract.AddressResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressResolver *AddressResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressResolver.Contract.AddressResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressResolver *AddressResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressResolver *AddressResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressResolver *AddressResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressResolver.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 name) view returns(address)
func (_AddressResolver *AddressResolverCaller) GetAddress(opts *bind.CallOpts, name [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AddressResolver.contract.Call(opts, &out, "getAddress", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 name) view returns(address)
func (_AddressResolver *AddressResolverSession) GetAddress(name [32]byte) (common.Address, error) {
	return _AddressResolver.Contract.GetAddress(&_AddressResolver.CallOpts, name)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 name) view returns(address)
func (_AddressResolver *AddressResolverCallerSession) GetAddress(name [32]byte) (common.Address, error) {
	return _AddressResolver.Contract.GetAddress(&_AddressResolver.CallOpts, name)
}

// GetSynth is a free data retrieval call binding the contract method 0x51456061.
//
// Solidity: function getSynth(bytes32 key) view returns(address)
func (_AddressResolver *AddressResolverCaller) GetSynth(opts *bind.CallOpts, key [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AddressResolver.contract.Call(opts, &out, "getSynth", key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSynth is a free data retrieval call binding the contract method 0x51456061.
//
// Solidity: function getSynth(bytes32 key) view returns(address)
func (_AddressResolver *AddressResolverSession) GetSynth(key [32]byte) (common.Address, error) {
	return _AddressResolver.Contract.GetSynth(&_AddressResolver.CallOpts, key)
}

// GetSynth is a free data retrieval call binding the contract method 0x51456061.
//
// Solidity: function getSynth(bytes32 key) view returns(address)
func (_AddressResolver *AddressResolverCallerSession) GetSynth(key [32]byte) (common.Address, error) {
	return _AddressResolver.Contract.GetSynth(&_AddressResolver.CallOpts, key)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_AddressResolver *AddressResolverCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressResolver.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_AddressResolver *AddressResolverSession) NominatedOwner() (common.Address, error) {
	return _AddressResolver.Contract.NominatedOwner(&_AddressResolver.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_AddressResolver *AddressResolverCallerSession) NominatedOwner() (common.Address, error) {
	return _AddressResolver.Contract.NominatedOwner(&_AddressResolver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressResolver *AddressResolverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressResolver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressResolver *AddressResolverSession) Owner() (common.Address, error) {
	return _AddressResolver.Contract.Owner(&_AddressResolver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressResolver *AddressResolverCallerSession) Owner() (common.Address, error) {
	return _AddressResolver.Contract.Owner(&_AddressResolver.CallOpts)
}

// Repository is a free data retrieval call binding the contract method 0x187f7935.
//
// Solidity: function repository(bytes32 ) view returns(address)
func (_AddressResolver *AddressResolverCaller) Repository(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AddressResolver.contract.Call(opts, &out, "repository", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Repository is a free data retrieval call binding the contract method 0x187f7935.
//
// Solidity: function repository(bytes32 ) view returns(address)
func (_AddressResolver *AddressResolverSession) Repository(arg0 [32]byte) (common.Address, error) {
	return _AddressResolver.Contract.Repository(&_AddressResolver.CallOpts, arg0)
}

// Repository is a free data retrieval call binding the contract method 0x187f7935.
//
// Solidity: function repository(bytes32 ) view returns(address)
func (_AddressResolver *AddressResolverCallerSession) Repository(arg0 [32]byte) (common.Address, error) {
	return _AddressResolver.Contract.Repository(&_AddressResolver.CallOpts, arg0)
}

// RequireAndGetAddress is a free data retrieval call binding the contract method 0xdacb2d01.
//
// Solidity: function requireAndGetAddress(bytes32 name, string reason) view returns(address)
func (_AddressResolver *AddressResolverCaller) RequireAndGetAddress(opts *bind.CallOpts, name [32]byte, reason string) (common.Address, error) {
	var out []interface{}
	err := _AddressResolver.contract.Call(opts, &out, "requireAndGetAddress", name, reason)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RequireAndGetAddress is a free data retrieval call binding the contract method 0xdacb2d01.
//
// Solidity: function requireAndGetAddress(bytes32 name, string reason) view returns(address)
func (_AddressResolver *AddressResolverSession) RequireAndGetAddress(name [32]byte, reason string) (common.Address, error) {
	return _AddressResolver.Contract.RequireAndGetAddress(&_AddressResolver.CallOpts, name, reason)
}

// RequireAndGetAddress is a free data retrieval call binding the contract method 0xdacb2d01.
//
// Solidity: function requireAndGetAddress(bytes32 name, string reason) view returns(address)
func (_AddressResolver *AddressResolverCallerSession) RequireAndGetAddress(name [32]byte, reason string) (common.Address, error) {
	return _AddressResolver.Contract.RequireAndGetAddress(&_AddressResolver.CallOpts, name, reason)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AddressResolver *AddressResolverTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressResolver.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AddressResolver *AddressResolverSession) AcceptOwnership() (*types.Transaction, error) {
	return _AddressResolver.Contract.AcceptOwnership(&_AddressResolver.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AddressResolver *AddressResolverTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AddressResolver.Contract.AcceptOwnership(&_AddressResolver.TransactOpts)
}

// ImportAddresses is a paid mutator transaction binding the contract method 0xab0b8f77.
//
// Solidity: function importAddresses(bytes32[] names, address[] destinations) returns()
func (_AddressResolver *AddressResolverTransactor) ImportAddresses(opts *bind.TransactOpts, names [][32]byte, destinations []common.Address) (*types.Transaction, error) {
	return _AddressResolver.contract.Transact(opts, "importAddresses", names, destinations)
}

// ImportAddresses is a paid mutator transaction binding the contract method 0xab0b8f77.
//
// Solidity: function importAddresses(bytes32[] names, address[] destinations) returns()
func (_AddressResolver *AddressResolverSession) ImportAddresses(names [][32]byte, destinations []common.Address) (*types.Transaction, error) {
	return _AddressResolver.Contract.ImportAddresses(&_AddressResolver.TransactOpts, names, destinations)
}

// ImportAddresses is a paid mutator transaction binding the contract method 0xab0b8f77.
//
// Solidity: function importAddresses(bytes32[] names, address[] destinations) returns()
func (_AddressResolver *AddressResolverTransactorSession) ImportAddresses(names [][32]byte, destinations []common.Address) (*types.Transaction, error) {
	return _AddressResolver.Contract.ImportAddresses(&_AddressResolver.TransactOpts, names, destinations)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_AddressResolver *AddressResolverTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _AddressResolver.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_AddressResolver *AddressResolverSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _AddressResolver.Contract.NominateNewOwner(&_AddressResolver.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_AddressResolver *AddressResolverTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _AddressResolver.Contract.NominateNewOwner(&_AddressResolver.TransactOpts, _owner)
}

// AddressResolverOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the AddressResolver contract.
type AddressResolverOwnerChangedIterator struct {
	Event *AddressResolverOwnerChanged // Event containing the contract specifics and raw log

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
func (it *AddressResolverOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressResolverOwnerChanged)
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
		it.Event = new(AddressResolverOwnerChanged)
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
func (it *AddressResolverOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressResolverOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressResolverOwnerChanged represents a OwnerChanged event raised by the AddressResolver contract.
type AddressResolverOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_AddressResolver *AddressResolverFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*AddressResolverOwnerChangedIterator, error) {

	logs, sub, err := _AddressResolver.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &AddressResolverOwnerChangedIterator{contract: _AddressResolver.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_AddressResolver *AddressResolverFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *AddressResolverOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _AddressResolver.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressResolverOwnerChanged)
				if err := _AddressResolver.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_AddressResolver *AddressResolverFilterer) ParseOwnerChanged(log types.Log) (*AddressResolverOwnerChanged, error) {
	event := new(AddressResolverOwnerChanged)
	if err := _AddressResolver.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressResolverOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the AddressResolver contract.
type AddressResolverOwnerNominatedIterator struct {
	Event *AddressResolverOwnerNominated // Event containing the contract specifics and raw log

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
func (it *AddressResolverOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressResolverOwnerNominated)
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
		it.Event = new(AddressResolverOwnerNominated)
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
func (it *AddressResolverOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressResolverOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressResolverOwnerNominated represents a OwnerNominated event raised by the AddressResolver contract.
type AddressResolverOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_AddressResolver *AddressResolverFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*AddressResolverOwnerNominatedIterator, error) {

	logs, sub, err := _AddressResolver.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &AddressResolverOwnerNominatedIterator{contract: _AddressResolver.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_AddressResolver *AddressResolverFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *AddressResolverOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _AddressResolver.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressResolverOwnerNominated)
				if err := _AddressResolver.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_AddressResolver *AddressResolverFilterer) ParseOwnerNominated(log types.Log) (*AddressResolverOwnerNominated, error) {
	event := new(AddressResolverOwnerNominated)
	if err := _AddressResolver.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternStateTokenABI is the input ABI used to generate the binding from.
const ExternStateTokenABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"contractTokenState\",\"name\":\"_tokenState\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"}],\"name\":\"ProxyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTokenState\",\"type\":\"address\"}],\"name\":\"TokenStateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"integrationProxy\",\"outputs\":[{\"internalType\":\"contractProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"messageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"contractProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_integrationProxy\",\"type\":\"address\"}],\"name\":\"setIntegrationProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"setMessageSender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"setProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractTokenState\",\"name\":\"_tokenState\",\"type\":\"address\"}],\"name\":\"setTokenState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenState\",\"outputs\":[{\"internalType\":\"contractTokenState\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ExternStateTokenFuncSigs maps the 4-byte function signature to its string representation.
var ExternStateTokenFuncSigs = map[string]string{
	"79ba5097": "acceptOwnership()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"9cbdaeb6": "integrationProxy()",
	"d67bdd25": "messageSender()",
	"06fdde03": "name()",
	"1627540c": "nominateNewOwner(address)",
	"53a47bb7": "nominatedOwner()",
	"8da5cb5b": "owner()",
	"ec556889": "proxy()",
	"131b0ae7": "setIntegrationProxy(address)",
	"bc67f832": "setMessageSender(address)",
	"97107d6d": "setProxy(address)",
	"9f769807": "setTokenState(address)",
	"95d89b41": "symbol()",
	"e90dd9e2": "tokenState()",
	"18160ddd": "totalSupply()",
}

// ExternStateTokenBin is the compiled bytecode used for deploying new contracts.
var ExternStateTokenBin = "0x60806040523480156200001157600080fd5b50604051620011bf380380620011bf833981810160405260e08110156200003757600080fd5b815160208301516040808501805191519395929483019291846401000000008211156200006357600080fd5b9083019060208201858111156200007957600080fd5b82516401000000008111828201881017156200009457600080fd5b82525081516020918201929091019080838360005b83811015620000c3578181015183820152602001620000a9565b50505050905090810190601f168015620000f15780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011557600080fd5b9083019060208201858111156200012b57600080fd5b82516401000000008111828201881017156200014657600080fd5b82525081516020918201929091019080838360005b83811015620001755781810151838201526020016200015b565b50505050905090810190601f168015620001a35780820380516001836020036101000a031916815260200191505b50604090815260208201519082015160609092015190935090915086816001600160a01b0381166200021c576040805162461bcd60e51b815260206004820152601960248201527f4f776e657220616464726573732063616e6e6f74206265203000000000000000604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b038316908117825560408051928352602083019190915280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a1506000546001600160a01b0316620002c7576040805162461bcd60e51b815260206004820152601160248201527013dddb995c881b5d5cdd081899481cd95d607a1b604482015290519081900360640190fd5b600280546001600160a01b0383166001600160a01b0319909116811790915560408051918252517ffc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e9181900360200190a150600580546001600160a01b0319166001600160a01b03881617905584516200034990600690602088019062000385565b5083516200035f90600790602087019062000385565b50506008919091556009805460ff191660ff909216919091179055506200042a92505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003c857805160ff1916838001178555620003f8565b82800160010185558215620003f8579182015b82811115620003f8578251825591602001919060010190620003db565b50620004069291506200040a565b5090565b6200042791905b8082111562000406576000815560010162000411565b90565b610d85806200043a6000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80638da5cb5b116100ad578063bc67f83211610071578063bc67f8321461031f578063d67bdd2514610345578063dd62ed3e1461034d578063e90dd9e21461037b578063ec5568891461038357610121565b80638da5cb5b146102bb57806395d89b41146102c357806397107d6d146102cb5780639cbdaeb6146102f15780639f769807146102f957610121565b806318160ddd116100f457806318160ddd14610231578063313ce5671461024b57806353a47bb71461026957806370a082311461028d57806379ba5097146102b357610121565b806306fdde0314610126578063095ea7b3146101a3578063131b0ae7146101e35780631627540c1461020b575b600080fd5b61012e61038b565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610168578181015183820152602001610150565b50505050905090810190601f1680156101955780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101cf600480360360408110156101b957600080fd5b506001600160a01b038135169060200135610419565b604080519115158252519081900360200190f35b610209600480360360208110156101f957600080fd5b50356001600160a01b03166104b2565b005b6102096004803603602081101561022157600080fd5b50356001600160a01b03166104dc565b610239610538565b60408051918252519081900360200190f35b61025361053e565b6040805160ff9092168252519081900360200190f35b610271610547565b604080516001600160a01b039092168252519081900360200190f35b610239600480360360208110156102a357600080fd5b50356001600160a01b0316610556565b6102096105d9565b610271610695565b61012e6106a4565b610209600480360360208110156102e157600080fd5b50356001600160a01b03166106ff565b61027161075b565b6102096004803603602081101561030f57600080fd5b50356001600160a01b031661076a565b6102096004803603602081101561033557600080fd5b50356001600160a01b0316610799565b6102716107c3565b6102396004803603604081101561036357600080fd5b506001600160a01b03813581169160200135166107d2565b61027161085e565b61027161086d565b6006805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104115780601f106103e657610100808354040283529160200191610411565b820191906000526020600020905b8154815290600101906020018083116103f457829003601f168201915b505050505081565b600061042361087c565b6004805460055460408051633691826360e21b81526001600160a01b039384169481018590528784166024820152604481018790529051919092169163da46098c91606480830192600092919082900301818387803b15801561048557600080fd5b505af1158015610499573d6000803e3d6000fd5b505050506104a88185856108d3565b5060019392505050565b6104ba610a0c565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6104e4610a0c565b600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce229181900360200190a150565b60085481565b60095460ff1681565b6001546001600160a01b031681565b600554604080516370a0823160e01b81526001600160a01b038481166004830152915160009392909216916370a0823191602480820192602092909190829003018186803b1580156105a757600080fd5b505afa1580156105bb573d6000803e3d6000fd5b505050506040513d60208110156105d157600080fd5b505192915050565b6001546001600160a01b031633146106225760405162461bcd60e51b8152600401808060200182810382526035815260200180610ccc6035913960400191505060405180910390fd5b600054600154604080516001600160a01b03938416815292909116602083015280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a160018054600080546001600160a01b03199081166001600160a01b03841617909155169055565b6000546001600160a01b031681565b6007805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104115780601f106103e657610100808354040283529160200191610411565b610707610a0c565b600280546001600160a01b0383166001600160a01b0319909116811790915560408051918252517ffc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e9181900360200190a150565b6003546001600160a01b031681565b610772610a55565b600580546001600160a01b0319166001600160a01b03831617905561079681610b05565b50565b6107a1610c4b565b600480546001600160a01b0319166001600160a01b0392909216919091179055565b6004546001600160a01b031681565b60055460408051636eb1769f60e11b81526001600160a01b03858116600483015284811660248301529151600093929092169163dd62ed3e91604480820192602092909190829003018186803b15801561082b57600080fd5b505afa15801561083f573d6000803e3d6000fd5b505050506040513d602081101561085557600080fd5b50519392505050565b6005546001600160a01b031681565b6002546001600160a01b031681565b6002546001600160a01b031633148015906108a257506003546001600160a01b03163314155b80156108b957506004546001600160a01b03163314155b156108d157600480546001600160a01b031916331790555b565b60025460408051602080820185905282518083039091018152908201918290526001600160a01b039092169163907dff9791600390806021610d3082396021019050604051809103902061092688610cbf565b61092f88610cbf565b60006040518763ffffffff1660e01b815260040180806020018781526020018681526020018581526020018481526020018360001b8152602001828103825288818151815260200191508051906020019080838360005b8381101561099e578181015183820152602001610986565b50505050905090810190601f1680156109cb5780820380516001836020036101000a031916815260200191505b50975050505050505050600060405180830381600087803b1580156109ef57600080fd5b505af1158015610a03573d6000803e3d6000fd5b50505050505050565b6000546001600160a01b031633146108d15760405162461bcd60e51b815260040180806020018281038252602f815260200180610d01602f913960400191505060405180910390fd5b6002546001600160a01b03163314801590610a7b57506003546001600160a01b03163314155b8015610a9257506004546001600160a01b03163314155b15610aaa57600480546001600160a01b031916331790555b6000546004546001600160a01b039081169116146108d1576040805162461bcd60e51b815260206004820152601360248201527227bbb732b91037b7363c90333ab731ba34b7b760691b604482015290519081900360640190fd5b600254604080516001600160a01b038481166020808401919091528351808403820181528385018086527f546f6b656e5374617465557064617465642861646472657373290000000000009052935192839003605a01832063907dff9760e01b8452600160248501819052604485018290526000606486018190526084860181905260a4860181905260c060048701908152875160c48801528751959098169763907dff97979692959394919384938493839260e490920191908a0190808383885b83811015610bdf578181015183820152602001610bc7565b50505050905090810190601f168015610c0c5780820380516001836020036101000a031916815260200191505b50975050505050505050600060405180830381600087803b158015610c3057600080fd5b505af1158015610c44573d6000803e3d6000fd5b5050505050565b6002546001600160a01b0316331480610c6e57506003546001600160a01b031633145b6108d1576040805162461bcd60e51b815260206004820152601760248201527f4f6e6c79207468652070726f78792063616e2063616c6c000000000000000000604482015290519081900360640190fd5b6001600160a01b03169056fe596f75206d757374206265206e6f6d696e61746564206265666f726520796f752063616e20616363657074206f776e6572736869704f6e6c792074686520636f6e7472616374206f776e6572206d617920706572666f726d207468697320616374696f6e417070726f76616c28616464726573732c616464726573732c75696e7432353629a265627a7a7231582013e432cb246f030dbacd2f9b098eee40f44e16feb5b3f2b26f6f4b6990aed95564736f6c63430005100032"

// DeployExternStateToken deploys a new Ethereum contract, binding an instance of ExternStateToken to it.
func DeployExternStateToken(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address, _tokenState common.Address, _name string, _symbol string, _totalSupply *big.Int, _decimals uint8, _owner common.Address) (common.Address, *types.Transaction, *ExternStateToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ExternStateTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExternStateTokenBin), backend, _proxy, _tokenState, _name, _symbol, _totalSupply, _decimals, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExternStateToken{ExternStateTokenCaller: ExternStateTokenCaller{contract: contract}, ExternStateTokenTransactor: ExternStateTokenTransactor{contract: contract}, ExternStateTokenFilterer: ExternStateTokenFilterer{contract: contract}}, nil
}

// ExternStateToken is an auto generated Go binding around an Ethereum contract.
type ExternStateToken struct {
	ExternStateTokenCaller     // Read-only binding to the contract
	ExternStateTokenTransactor // Write-only binding to the contract
	ExternStateTokenFilterer   // Log filterer for contract events
}

// ExternStateTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExternStateTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExternStateTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExternStateTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExternStateTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExternStateTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExternStateTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExternStateTokenSession struct {
	Contract     *ExternStateToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExternStateTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExternStateTokenCallerSession struct {
	Contract *ExternStateTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ExternStateTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExternStateTokenTransactorSession struct {
	Contract     *ExternStateTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ExternStateTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExternStateTokenRaw struct {
	Contract *ExternStateToken // Generic contract binding to access the raw methods on
}

// ExternStateTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExternStateTokenCallerRaw struct {
	Contract *ExternStateTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ExternStateTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExternStateTokenTransactorRaw struct {
	Contract *ExternStateTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExternStateToken creates a new instance of ExternStateToken, bound to a specific deployed contract.
func NewExternStateToken(address common.Address, backend bind.ContractBackend) (*ExternStateToken, error) {
	contract, err := bindExternStateToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExternStateToken{ExternStateTokenCaller: ExternStateTokenCaller{contract: contract}, ExternStateTokenTransactor: ExternStateTokenTransactor{contract: contract}, ExternStateTokenFilterer: ExternStateTokenFilterer{contract: contract}}, nil
}

// NewExternStateTokenCaller creates a new read-only instance of ExternStateToken, bound to a specific deployed contract.
func NewExternStateTokenCaller(address common.Address, caller bind.ContractCaller) (*ExternStateTokenCaller, error) {
	contract, err := bindExternStateToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExternStateTokenCaller{contract: contract}, nil
}

// NewExternStateTokenTransactor creates a new write-only instance of ExternStateToken, bound to a specific deployed contract.
func NewExternStateTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ExternStateTokenTransactor, error) {
	contract, err := bindExternStateToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExternStateTokenTransactor{contract: contract}, nil
}

// NewExternStateTokenFilterer creates a new log filterer instance of ExternStateToken, bound to a specific deployed contract.
func NewExternStateTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ExternStateTokenFilterer, error) {
	contract, err := bindExternStateToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExternStateTokenFilterer{contract: contract}, nil
}

// bindExternStateToken binds a generic wrapper to an already deployed contract.
func bindExternStateToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExternStateTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExternStateToken *ExternStateTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExternStateToken.Contract.ExternStateTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExternStateToken *ExternStateTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExternStateToken.Contract.ExternStateTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExternStateToken *ExternStateTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExternStateToken.Contract.ExternStateTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExternStateToken *ExternStateTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExternStateToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExternStateToken *ExternStateTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExternStateToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExternStateToken *ExternStateTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExternStateToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExternStateToken *ExternStateTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExternStateToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExternStateToken *ExternStateTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ExternStateToken.Contract.Allowance(&_ExternStateToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExternStateToken *ExternStateTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ExternStateToken.Contract.Allowance(&_ExternStateToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExternStateToken *ExternStateTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExternStateToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExternStateToken *ExternStateTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExternStateToken.Contract.BalanceOf(&_ExternStateToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExternStateToken *ExternStateTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExternStateToken.Contract.BalanceOf(&_ExternStateToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExternStateToken *ExternStateTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ExternStateToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExternStateToken *ExternStateTokenSession) Decimals() (uint8, error) {
	return _ExternStateToken.Contract.Decimals(&_ExternStateToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExternStateToken *ExternStateTokenCallerSession) Decimals() (uint8, error) {
	return _ExternStateToken.Contract.Decimals(&_ExternStateToken.CallOpts)
}

// IntegrationProxy is a free data retrieval call binding the contract method 0x9cbdaeb6.
//
// Solidity: function integrationProxy() view returns(address)
func (_ExternStateToken *ExternStateTokenCaller) IntegrationProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExternStateToken.contract.Call(opts, &out, "integrationProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IntegrationProxy is a free data retrieval call binding the contract method 0x9cbdaeb6.
//
// Solidity: function integrationProxy() view returns(address)
func (_ExternStateToken *ExternStateTokenSession) IntegrationProxy() (common.Address, error) {
	return _ExternStateToken.Contract.IntegrationProxy(&_ExternStateToken.CallOpts)
}

// IntegrationProxy is a free data retrieval call binding the contract method 0x9cbdaeb6.
//
// Solidity: function integrationProxy() view returns(address)
func (_ExternStateToken *ExternStateTokenCallerSession) IntegrationProxy() (common.Address, error) {
	return _ExternStateToken.Contract.IntegrationProxy(&_ExternStateToken.CallOpts)
}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_ExternStateToken *ExternStateTokenCaller) MessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExternStateToken.contract.Call(opts, &out, "messageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_ExternStateToken *ExternStateTokenSession) MessageSender() (common.Address, error) {
	return _ExternStateToken.Contract.MessageSender(&_ExternStateToken.CallOpts)
}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_ExternStateToken *ExternStateTokenCallerSession) MessageSender() (common.Address, error) {
	return _ExternStateToken.Contract.MessageSender(&_ExternStateToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExternStateToken *ExternStateTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExternStateToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExternStateToken *ExternStateTokenSession) Name() (string, error) {
	return _ExternStateToken.Contract.Name(&_ExternStateToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExternStateToken *ExternStateTokenCallerSession) Name() (string, error) {
	return _ExternStateToken.Contract.Name(&_ExternStateToken.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_ExternStateToken *ExternStateTokenCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExternStateToken.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_ExternStateToken *ExternStateTokenSession) NominatedOwner() (common.Address, error) {
	return _ExternStateToken.Contract.NominatedOwner(&_ExternStateToken.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_ExternStateToken *ExternStateTokenCallerSession) NominatedOwner() (common.Address, error) {
	return _ExternStateToken.Contract.NominatedOwner(&_ExternStateToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExternStateToken *ExternStateTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExternStateToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExternStateToken *ExternStateTokenSession) Owner() (common.Address, error) {
	return _ExternStateToken.Contract.Owner(&_ExternStateToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExternStateToken *ExternStateTokenCallerSession) Owner() (common.Address, error) {
	return _ExternStateToken.Contract.Owner(&_ExternStateToken.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_ExternStateToken *ExternStateTokenCaller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExternStateToken.contract.Call(opts, &out, "proxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_ExternStateToken *ExternStateTokenSession) Proxy() (common.Address, error) {
	return _ExternStateToken.Contract.Proxy(&_ExternStateToken.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_ExternStateToken *ExternStateTokenCallerSession) Proxy() (common.Address, error) {
	return _ExternStateToken.Contract.Proxy(&_ExternStateToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExternStateToken *ExternStateTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExternStateToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExternStateToken *ExternStateTokenSession) Symbol() (string, error) {
	return _ExternStateToken.Contract.Symbol(&_ExternStateToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExternStateToken *ExternStateTokenCallerSession) Symbol() (string, error) {
	return _ExternStateToken.Contract.Symbol(&_ExternStateToken.CallOpts)
}

// TokenState is a free data retrieval call binding the contract method 0xe90dd9e2.
//
// Solidity: function tokenState() view returns(address)
func (_ExternStateToken *ExternStateTokenCaller) TokenState(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExternStateToken.contract.Call(opts, &out, "tokenState")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenState is a free data retrieval call binding the contract method 0xe90dd9e2.
//
// Solidity: function tokenState() view returns(address)
func (_ExternStateToken *ExternStateTokenSession) TokenState() (common.Address, error) {
	return _ExternStateToken.Contract.TokenState(&_ExternStateToken.CallOpts)
}

// TokenState is a free data retrieval call binding the contract method 0xe90dd9e2.
//
// Solidity: function tokenState() view returns(address)
func (_ExternStateToken *ExternStateTokenCallerSession) TokenState() (common.Address, error) {
	return _ExternStateToken.Contract.TokenState(&_ExternStateToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExternStateToken *ExternStateTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExternStateToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExternStateToken *ExternStateTokenSession) TotalSupply() (*big.Int, error) {
	return _ExternStateToken.Contract.TotalSupply(&_ExternStateToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExternStateToken *ExternStateTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ExternStateToken.Contract.TotalSupply(&_ExternStateToken.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ExternStateToken *ExternStateTokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExternStateToken.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ExternStateToken *ExternStateTokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _ExternStateToken.Contract.AcceptOwnership(&_ExternStateToken.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ExternStateToken *ExternStateTokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ExternStateToken.Contract.AcceptOwnership(&_ExternStateToken.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ExternStateToken *ExternStateTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExternStateToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ExternStateToken *ExternStateTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExternStateToken.Contract.Approve(&_ExternStateToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ExternStateToken *ExternStateTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExternStateToken.Contract.Approve(&_ExternStateToken.TransactOpts, spender, value)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_ExternStateToken *ExternStateTokenTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _ExternStateToken.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_ExternStateToken *ExternStateTokenSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _ExternStateToken.Contract.NominateNewOwner(&_ExternStateToken.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_ExternStateToken *ExternStateTokenTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _ExternStateToken.Contract.NominateNewOwner(&_ExternStateToken.TransactOpts, _owner)
}

// SetIntegrationProxy is a paid mutator transaction binding the contract method 0x131b0ae7.
//
// Solidity: function setIntegrationProxy(address _integrationProxy) returns()
func (_ExternStateToken *ExternStateTokenTransactor) SetIntegrationProxy(opts *bind.TransactOpts, _integrationProxy common.Address) (*types.Transaction, error) {
	return _ExternStateToken.contract.Transact(opts, "setIntegrationProxy", _integrationProxy)
}

// SetIntegrationProxy is a paid mutator transaction binding the contract method 0x131b0ae7.
//
// Solidity: function setIntegrationProxy(address _integrationProxy) returns()
func (_ExternStateToken *ExternStateTokenSession) SetIntegrationProxy(_integrationProxy common.Address) (*types.Transaction, error) {
	return _ExternStateToken.Contract.SetIntegrationProxy(&_ExternStateToken.TransactOpts, _integrationProxy)
}

// SetIntegrationProxy is a paid mutator transaction binding the contract method 0x131b0ae7.
//
// Solidity: function setIntegrationProxy(address _integrationProxy) returns()
func (_ExternStateToken *ExternStateTokenTransactorSession) SetIntegrationProxy(_integrationProxy common.Address) (*types.Transaction, error) {
	return _ExternStateToken.Contract.SetIntegrationProxy(&_ExternStateToken.TransactOpts, _integrationProxy)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_ExternStateToken *ExternStateTokenTransactor) SetMessageSender(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _ExternStateToken.contract.Transact(opts, "setMessageSender", sender)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_ExternStateToken *ExternStateTokenSession) SetMessageSender(sender common.Address) (*types.Transaction, error) {
	return _ExternStateToken.Contract.SetMessageSender(&_ExternStateToken.TransactOpts, sender)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_ExternStateToken *ExternStateTokenTransactorSession) SetMessageSender(sender common.Address) (*types.Transaction, error) {
	return _ExternStateToken.Contract.SetMessageSender(&_ExternStateToken.TransactOpts, sender)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_ExternStateToken *ExternStateTokenTransactor) SetProxy(opts *bind.TransactOpts, _proxy common.Address) (*types.Transaction, error) {
	return _ExternStateToken.contract.Transact(opts, "setProxy", _proxy)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_ExternStateToken *ExternStateTokenSession) SetProxy(_proxy common.Address) (*types.Transaction, error) {
	return _ExternStateToken.Contract.SetProxy(&_ExternStateToken.TransactOpts, _proxy)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_ExternStateToken *ExternStateTokenTransactorSession) SetProxy(_proxy common.Address) (*types.Transaction, error) {
	return _ExternStateToken.Contract.SetProxy(&_ExternStateToken.TransactOpts, _proxy)
}

// SetTokenState is a paid mutator transaction binding the contract method 0x9f769807.
//
// Solidity: function setTokenState(address _tokenState) returns()
func (_ExternStateToken *ExternStateTokenTransactor) SetTokenState(opts *bind.TransactOpts, _tokenState common.Address) (*types.Transaction, error) {
	return _ExternStateToken.contract.Transact(opts, "setTokenState", _tokenState)
}

// SetTokenState is a paid mutator transaction binding the contract method 0x9f769807.
//
// Solidity: function setTokenState(address _tokenState) returns()
func (_ExternStateToken *ExternStateTokenSession) SetTokenState(_tokenState common.Address) (*types.Transaction, error) {
	return _ExternStateToken.Contract.SetTokenState(&_ExternStateToken.TransactOpts, _tokenState)
}

// SetTokenState is a paid mutator transaction binding the contract method 0x9f769807.
//
// Solidity: function setTokenState(address _tokenState) returns()
func (_ExternStateToken *ExternStateTokenTransactorSession) SetTokenState(_tokenState common.Address) (*types.Transaction, error) {
	return _ExternStateToken.Contract.SetTokenState(&_ExternStateToken.TransactOpts, _tokenState)
}

// ExternStateTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ExternStateToken contract.
type ExternStateTokenApprovalIterator struct {
	Event *ExternStateTokenApproval // Event containing the contract specifics and raw log

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
func (it *ExternStateTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternStateTokenApproval)
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
		it.Event = new(ExternStateTokenApproval)
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
func (it *ExternStateTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternStateTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternStateTokenApproval represents a Approval event raised by the ExternStateToken contract.
type ExternStateTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ExternStateToken *ExternStateTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ExternStateTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ExternStateToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ExternStateTokenApprovalIterator{contract: _ExternStateToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ExternStateToken *ExternStateTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ExternStateTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ExternStateToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternStateTokenApproval)
				if err := _ExternStateToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ExternStateToken *ExternStateTokenFilterer) ParseApproval(log types.Log) (*ExternStateTokenApproval, error) {
	event := new(ExternStateTokenApproval)
	if err := _ExternStateToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternStateTokenOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the ExternStateToken contract.
type ExternStateTokenOwnerChangedIterator struct {
	Event *ExternStateTokenOwnerChanged // Event containing the contract specifics and raw log

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
func (it *ExternStateTokenOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternStateTokenOwnerChanged)
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
		it.Event = new(ExternStateTokenOwnerChanged)
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
func (it *ExternStateTokenOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternStateTokenOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternStateTokenOwnerChanged represents a OwnerChanged event raised by the ExternStateToken contract.
type ExternStateTokenOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_ExternStateToken *ExternStateTokenFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*ExternStateTokenOwnerChangedIterator, error) {

	logs, sub, err := _ExternStateToken.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &ExternStateTokenOwnerChangedIterator{contract: _ExternStateToken.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_ExternStateToken *ExternStateTokenFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *ExternStateTokenOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _ExternStateToken.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternStateTokenOwnerChanged)
				if err := _ExternStateToken.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_ExternStateToken *ExternStateTokenFilterer) ParseOwnerChanged(log types.Log) (*ExternStateTokenOwnerChanged, error) {
	event := new(ExternStateTokenOwnerChanged)
	if err := _ExternStateToken.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternStateTokenOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the ExternStateToken contract.
type ExternStateTokenOwnerNominatedIterator struct {
	Event *ExternStateTokenOwnerNominated // Event containing the contract specifics and raw log

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
func (it *ExternStateTokenOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternStateTokenOwnerNominated)
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
		it.Event = new(ExternStateTokenOwnerNominated)
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
func (it *ExternStateTokenOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternStateTokenOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternStateTokenOwnerNominated represents a OwnerNominated event raised by the ExternStateToken contract.
type ExternStateTokenOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_ExternStateToken *ExternStateTokenFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*ExternStateTokenOwnerNominatedIterator, error) {

	logs, sub, err := _ExternStateToken.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &ExternStateTokenOwnerNominatedIterator{contract: _ExternStateToken.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_ExternStateToken *ExternStateTokenFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *ExternStateTokenOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _ExternStateToken.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternStateTokenOwnerNominated)
				if err := _ExternStateToken.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_ExternStateToken *ExternStateTokenFilterer) ParseOwnerNominated(log types.Log) (*ExternStateTokenOwnerNominated, error) {
	event := new(ExternStateTokenOwnerNominated)
	if err := _ExternStateToken.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternStateTokenProxyUpdatedIterator is returned from FilterProxyUpdated and is used to iterate over the raw logs and unpacked data for ProxyUpdated events raised by the ExternStateToken contract.
type ExternStateTokenProxyUpdatedIterator struct {
	Event *ExternStateTokenProxyUpdated // Event containing the contract specifics and raw log

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
func (it *ExternStateTokenProxyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternStateTokenProxyUpdated)
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
		it.Event = new(ExternStateTokenProxyUpdated)
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
func (it *ExternStateTokenProxyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternStateTokenProxyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternStateTokenProxyUpdated represents a ProxyUpdated event raised by the ExternStateToken contract.
type ExternStateTokenProxyUpdated struct {
	ProxyAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProxyUpdated is a free log retrieval operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_ExternStateToken *ExternStateTokenFilterer) FilterProxyUpdated(opts *bind.FilterOpts) (*ExternStateTokenProxyUpdatedIterator, error) {

	logs, sub, err := _ExternStateToken.contract.FilterLogs(opts, "ProxyUpdated")
	if err != nil {
		return nil, err
	}
	return &ExternStateTokenProxyUpdatedIterator{contract: _ExternStateToken.contract, event: "ProxyUpdated", logs: logs, sub: sub}, nil
}

// WatchProxyUpdated is a free log subscription operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_ExternStateToken *ExternStateTokenFilterer) WatchProxyUpdated(opts *bind.WatchOpts, sink chan<- *ExternStateTokenProxyUpdated) (event.Subscription, error) {

	logs, sub, err := _ExternStateToken.contract.WatchLogs(opts, "ProxyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternStateTokenProxyUpdated)
				if err := _ExternStateToken.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
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

// ParseProxyUpdated is a log parse operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_ExternStateToken *ExternStateTokenFilterer) ParseProxyUpdated(log types.Log) (*ExternStateTokenProxyUpdated, error) {
	event := new(ExternStateTokenProxyUpdated)
	if err := _ExternStateToken.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternStateTokenTokenStateUpdatedIterator is returned from FilterTokenStateUpdated and is used to iterate over the raw logs and unpacked data for TokenStateUpdated events raised by the ExternStateToken contract.
type ExternStateTokenTokenStateUpdatedIterator struct {
	Event *ExternStateTokenTokenStateUpdated // Event containing the contract specifics and raw log

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
func (it *ExternStateTokenTokenStateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternStateTokenTokenStateUpdated)
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
		it.Event = new(ExternStateTokenTokenStateUpdated)
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
func (it *ExternStateTokenTokenStateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternStateTokenTokenStateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternStateTokenTokenStateUpdated represents a TokenStateUpdated event raised by the ExternStateToken contract.
type ExternStateTokenTokenStateUpdated struct {
	NewTokenState common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenStateUpdated is a free log retrieval operation binding the contract event 0xa538c4dcfe9fb148efee2952bafe34982d2d07d5fbb38ae5b44abf659a46bfd8.
//
// Solidity: event TokenStateUpdated(address newTokenState)
func (_ExternStateToken *ExternStateTokenFilterer) FilterTokenStateUpdated(opts *bind.FilterOpts) (*ExternStateTokenTokenStateUpdatedIterator, error) {

	logs, sub, err := _ExternStateToken.contract.FilterLogs(opts, "TokenStateUpdated")
	if err != nil {
		return nil, err
	}
	return &ExternStateTokenTokenStateUpdatedIterator{contract: _ExternStateToken.contract, event: "TokenStateUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenStateUpdated is a free log subscription operation binding the contract event 0xa538c4dcfe9fb148efee2952bafe34982d2d07d5fbb38ae5b44abf659a46bfd8.
//
// Solidity: event TokenStateUpdated(address newTokenState)
func (_ExternStateToken *ExternStateTokenFilterer) WatchTokenStateUpdated(opts *bind.WatchOpts, sink chan<- *ExternStateTokenTokenStateUpdated) (event.Subscription, error) {

	logs, sub, err := _ExternStateToken.contract.WatchLogs(opts, "TokenStateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternStateTokenTokenStateUpdated)
				if err := _ExternStateToken.contract.UnpackLog(event, "TokenStateUpdated", log); err != nil {
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

// ParseTokenStateUpdated is a log parse operation binding the contract event 0xa538c4dcfe9fb148efee2952bafe34982d2d07d5fbb38ae5b44abf659a46bfd8.
//
// Solidity: event TokenStateUpdated(address newTokenState)
func (_ExternStateToken *ExternStateTokenFilterer) ParseTokenStateUpdated(log types.Log) (*ExternStateTokenTokenStateUpdated, error) {
	event := new(ExternStateTokenTokenStateUpdated)
	if err := _ExternStateToken.contract.UnpackLog(event, "TokenStateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternStateTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ExternStateToken contract.
type ExternStateTokenTransferIterator struct {
	Event *ExternStateTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ExternStateTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternStateTokenTransfer)
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
		it.Event = new(ExternStateTokenTransfer)
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
func (it *ExternStateTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternStateTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternStateTokenTransfer represents a Transfer event raised by the ExternStateToken contract.
type ExternStateTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ExternStateToken *ExternStateTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExternStateTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExternStateToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExternStateTokenTransferIterator{contract: _ExternStateToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ExternStateToken *ExternStateTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ExternStateTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExternStateToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternStateTokenTransfer)
				if err := _ExternStateToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ExternStateToken *ExternStateTokenFilterer) ParseTransfer(log types.Log) (*ExternStateTokenTransfer, error) {
	event := new(ExternStateTokenTransfer)
	if err := _ExternStateToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAddressResolverABI is the input ABI used to generate the binding from.
const IAddressResolverABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getSynth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"requireAndGetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IAddressResolverFuncSigs maps the 4-byte function signature to its string representation.
var IAddressResolverFuncSigs = map[string]string{
	"21f8a721": "getAddress(bytes32)",
	"51456061": "getSynth(bytes32)",
	"dacb2d01": "requireAndGetAddress(bytes32,string)",
}

// IAddressResolver is an auto generated Go binding around an Ethereum contract.
type IAddressResolver struct {
	IAddressResolverCaller     // Read-only binding to the contract
	IAddressResolverTransactor // Write-only binding to the contract
	IAddressResolverFilterer   // Log filterer for contract events
}

// IAddressResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAddressResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAddressResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAddressResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAddressResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAddressResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAddressResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAddressResolverSession struct {
	Contract     *IAddressResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAddressResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAddressResolverCallerSession struct {
	Contract *IAddressResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IAddressResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAddressResolverTransactorSession struct {
	Contract     *IAddressResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IAddressResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAddressResolverRaw struct {
	Contract *IAddressResolver // Generic contract binding to access the raw methods on
}

// IAddressResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAddressResolverCallerRaw struct {
	Contract *IAddressResolverCaller // Generic read-only contract binding to access the raw methods on
}

// IAddressResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAddressResolverTransactorRaw struct {
	Contract *IAddressResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAddressResolver creates a new instance of IAddressResolver, bound to a specific deployed contract.
func NewIAddressResolver(address common.Address, backend bind.ContractBackend) (*IAddressResolver, error) {
	contract, err := bindIAddressResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAddressResolver{IAddressResolverCaller: IAddressResolverCaller{contract: contract}, IAddressResolverTransactor: IAddressResolverTransactor{contract: contract}, IAddressResolverFilterer: IAddressResolverFilterer{contract: contract}}, nil
}

// NewIAddressResolverCaller creates a new read-only instance of IAddressResolver, bound to a specific deployed contract.
func NewIAddressResolverCaller(address common.Address, caller bind.ContractCaller) (*IAddressResolverCaller, error) {
	contract, err := bindIAddressResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAddressResolverCaller{contract: contract}, nil
}

// NewIAddressResolverTransactor creates a new write-only instance of IAddressResolver, bound to a specific deployed contract.
func NewIAddressResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*IAddressResolverTransactor, error) {
	contract, err := bindIAddressResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAddressResolverTransactor{contract: contract}, nil
}

// NewIAddressResolverFilterer creates a new log filterer instance of IAddressResolver, bound to a specific deployed contract.
func NewIAddressResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*IAddressResolverFilterer, error) {
	contract, err := bindIAddressResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAddressResolverFilterer{contract: contract}, nil
}

// bindIAddressResolver binds a generic wrapper to an already deployed contract.
func bindIAddressResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAddressResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAddressResolver *IAddressResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAddressResolver.Contract.IAddressResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAddressResolver *IAddressResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAddressResolver.Contract.IAddressResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAddressResolver *IAddressResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAddressResolver.Contract.IAddressResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAddressResolver *IAddressResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAddressResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAddressResolver *IAddressResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAddressResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAddressResolver *IAddressResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAddressResolver.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 name) view returns(address)
func (_IAddressResolver *IAddressResolverCaller) GetAddress(opts *bind.CallOpts, name [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IAddressResolver.contract.Call(opts, &out, "getAddress", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 name) view returns(address)
func (_IAddressResolver *IAddressResolverSession) GetAddress(name [32]byte) (common.Address, error) {
	return _IAddressResolver.Contract.GetAddress(&_IAddressResolver.CallOpts, name)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 name) view returns(address)
func (_IAddressResolver *IAddressResolverCallerSession) GetAddress(name [32]byte) (common.Address, error) {
	return _IAddressResolver.Contract.GetAddress(&_IAddressResolver.CallOpts, name)
}

// GetSynth is a free data retrieval call binding the contract method 0x51456061.
//
// Solidity: function getSynth(bytes32 key) view returns(address)
func (_IAddressResolver *IAddressResolverCaller) GetSynth(opts *bind.CallOpts, key [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IAddressResolver.contract.Call(opts, &out, "getSynth", key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSynth is a free data retrieval call binding the contract method 0x51456061.
//
// Solidity: function getSynth(bytes32 key) view returns(address)
func (_IAddressResolver *IAddressResolverSession) GetSynth(key [32]byte) (common.Address, error) {
	return _IAddressResolver.Contract.GetSynth(&_IAddressResolver.CallOpts, key)
}

// GetSynth is a free data retrieval call binding the contract method 0x51456061.
//
// Solidity: function getSynth(bytes32 key) view returns(address)
func (_IAddressResolver *IAddressResolverCallerSession) GetSynth(key [32]byte) (common.Address, error) {
	return _IAddressResolver.Contract.GetSynth(&_IAddressResolver.CallOpts, key)
}

// RequireAndGetAddress is a free data retrieval call binding the contract method 0xdacb2d01.
//
// Solidity: function requireAndGetAddress(bytes32 name, string reason) view returns(address)
func (_IAddressResolver *IAddressResolverCaller) RequireAndGetAddress(opts *bind.CallOpts, name [32]byte, reason string) (common.Address, error) {
	var out []interface{}
	err := _IAddressResolver.contract.Call(opts, &out, "requireAndGetAddress", name, reason)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RequireAndGetAddress is a free data retrieval call binding the contract method 0xdacb2d01.
//
// Solidity: function requireAndGetAddress(bytes32 name, string reason) view returns(address)
func (_IAddressResolver *IAddressResolverSession) RequireAndGetAddress(name [32]byte, reason string) (common.Address, error) {
	return _IAddressResolver.Contract.RequireAndGetAddress(&_IAddressResolver.CallOpts, name, reason)
}

// RequireAndGetAddress is a free data retrieval call binding the contract method 0xdacb2d01.
//
// Solidity: function requireAndGetAddress(bytes32 name, string reason) view returns(address)
func (_IAddressResolver *IAddressResolverCallerSession) RequireAndGetAddress(name [32]byte, reason string) (common.Address, error) {
	return _IAddressResolver.Contract.RequireAndGetAddress(&_IAddressResolver.CallOpts, name, reason)
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Session) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20CallerSession) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Session) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20CallerSession) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Session) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20CallerSession) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExchangerABI is the input ABI used to generate the binding from.
const IExchangerABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refunded\",\"type\":\"uint256\"}],\"name\":\"calculateAmountAfterSettlement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountAfterSettlement\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"exchangeForAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"}],\"name\":\"exchangeOnBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"exchangeForAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"}],\"name\":\"exchangeOnBehalfWithTracking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"}],\"name\":\"exchangeWithTracking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"}],\"name\":\"exchangeWithVirtual\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"},{\"internalType\":\"contractIVirtualSynth\",\"name\":\"vSynth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"}],\"name\":\"feeRateForExchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"exchangeFeeRate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"}],\"name\":\"getAmountsForExchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchangeFeeRate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"hasWaitingPeriodOrSettlementOwing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"isSynthRateInvalid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"maxSecsLeftInWaitingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceDeviationThresholdFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"setLastExchangeRateForSynth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"settle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reclaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refunded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEntries\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"settlementOwing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reclaimAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rebateAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEntries\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"suspendSynthWithInvalidRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"waitingPeriodSecs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IExchangerFuncSigs maps the 4-byte function signature to its string representation.
var IExchangerFuncSigs = map[string]string{
	"4c268fc8": "calculateAmountAfterSettlement(address,bytes32,uint256,uint256)",
	"0a1e187d": "exchange(address,bytes32,uint256,bytes32,address)",
	"6a1c4758": "exchangeOnBehalf(address,address,bytes32,uint256,bytes32)",
	"dfffca76": "exchangeOnBehalfWithTracking(address,address,bytes32,uint256,bytes32,address,bytes32)",
	"86baa45c": "exchangeWithTracking(address,bytes32,uint256,bytes32,address,address,bytes32)",
	"f3995224": "exchangeWithVirtual(address,bytes32,uint256,bytes32,address,bytes32)",
	"1a5c6095": "feeRateForExchange(bytes32,bytes32)",
	"f450aa34": "getAmountsForExchange(uint256,bytes32,bytes32)",
	"d6f32e06": "hasWaitingPeriodOrSettlementOwing(address,bytes32)",
	"57af302c": "isSynthRateInvalid(bytes32)",
	"059c29ec": "maxSecsLeftInWaitingPeriod(address,bytes32)",
	"372a395a": "priceDeviationThresholdFactor()",
	"ce096940": "setLastExchangeRateForSynth(bytes32,uint256)",
	"1b16802c": "settle(address,bytes32)",
	"19d5c665": "settlementOwing(address,bytes32)",
	"0b9e31c9": "suspendSynthWithInvalidRate(bytes32)",
	"89257117": "waitingPeriodSecs()",
}

// IExchanger is an auto generated Go binding around an Ethereum contract.
type IExchanger struct {
	IExchangerCaller     // Read-only binding to the contract
	IExchangerTransactor // Write-only binding to the contract
	IExchangerFilterer   // Log filterer for contract events
}

// IExchangerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IExchangerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExchangerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IExchangerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExchangerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IExchangerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExchangerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IExchangerSession struct {
	Contract     *IExchanger       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IExchangerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IExchangerCallerSession struct {
	Contract *IExchangerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IExchangerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IExchangerTransactorSession struct {
	Contract     *IExchangerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IExchangerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IExchangerRaw struct {
	Contract *IExchanger // Generic contract binding to access the raw methods on
}

// IExchangerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IExchangerCallerRaw struct {
	Contract *IExchangerCaller // Generic read-only contract binding to access the raw methods on
}

// IExchangerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IExchangerTransactorRaw struct {
	Contract *IExchangerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIExchanger creates a new instance of IExchanger, bound to a specific deployed contract.
func NewIExchanger(address common.Address, backend bind.ContractBackend) (*IExchanger, error) {
	contract, err := bindIExchanger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IExchanger{IExchangerCaller: IExchangerCaller{contract: contract}, IExchangerTransactor: IExchangerTransactor{contract: contract}, IExchangerFilterer: IExchangerFilterer{contract: contract}}, nil
}

// NewIExchangerCaller creates a new read-only instance of IExchanger, bound to a specific deployed contract.
func NewIExchangerCaller(address common.Address, caller bind.ContractCaller) (*IExchangerCaller, error) {
	contract, err := bindIExchanger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IExchangerCaller{contract: contract}, nil
}

// NewIExchangerTransactor creates a new write-only instance of IExchanger, bound to a specific deployed contract.
func NewIExchangerTransactor(address common.Address, transactor bind.ContractTransactor) (*IExchangerTransactor, error) {
	contract, err := bindIExchanger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IExchangerTransactor{contract: contract}, nil
}

// NewIExchangerFilterer creates a new log filterer instance of IExchanger, bound to a specific deployed contract.
func NewIExchangerFilterer(address common.Address, filterer bind.ContractFilterer) (*IExchangerFilterer, error) {
	contract, err := bindIExchanger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IExchangerFilterer{contract: contract}, nil
}

// bindIExchanger binds a generic wrapper to an already deployed contract.
func bindIExchanger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IExchangerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExchanger *IExchangerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExchanger.Contract.IExchangerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExchanger *IExchangerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExchanger.Contract.IExchangerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExchanger *IExchangerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExchanger.Contract.IExchangerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExchanger *IExchangerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExchanger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExchanger *IExchangerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExchanger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExchanger *IExchangerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExchanger.Contract.contract.Transact(opts, method, params...)
}

// CalculateAmountAfterSettlement is a free data retrieval call binding the contract method 0x4c268fc8.
//
// Solidity: function calculateAmountAfterSettlement(address from, bytes32 currencyKey, uint256 amount, uint256 refunded) view returns(uint256 amountAfterSettlement)
func (_IExchanger *IExchangerCaller) CalculateAmountAfterSettlement(opts *bind.CallOpts, from common.Address, currencyKey [32]byte, amount *big.Int, refunded *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IExchanger.contract.Call(opts, &out, "calculateAmountAfterSettlement", from, currencyKey, amount, refunded)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateAmountAfterSettlement is a free data retrieval call binding the contract method 0x4c268fc8.
//
// Solidity: function calculateAmountAfterSettlement(address from, bytes32 currencyKey, uint256 amount, uint256 refunded) view returns(uint256 amountAfterSettlement)
func (_IExchanger *IExchangerSession) CalculateAmountAfterSettlement(from common.Address, currencyKey [32]byte, amount *big.Int, refunded *big.Int) (*big.Int, error) {
	return _IExchanger.Contract.CalculateAmountAfterSettlement(&_IExchanger.CallOpts, from, currencyKey, amount, refunded)
}

// CalculateAmountAfterSettlement is a free data retrieval call binding the contract method 0x4c268fc8.
//
// Solidity: function calculateAmountAfterSettlement(address from, bytes32 currencyKey, uint256 amount, uint256 refunded) view returns(uint256 amountAfterSettlement)
func (_IExchanger *IExchangerCallerSession) CalculateAmountAfterSettlement(from common.Address, currencyKey [32]byte, amount *big.Int, refunded *big.Int) (*big.Int, error) {
	return _IExchanger.Contract.CalculateAmountAfterSettlement(&_IExchanger.CallOpts, from, currencyKey, amount, refunded)
}

// FeeRateForExchange is a free data retrieval call binding the contract method 0x1a5c6095.
//
// Solidity: function feeRateForExchange(bytes32 sourceCurrencyKey, bytes32 destinationCurrencyKey) view returns(uint256 exchangeFeeRate)
func (_IExchanger *IExchangerCaller) FeeRateForExchange(opts *bind.CallOpts, sourceCurrencyKey [32]byte, destinationCurrencyKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IExchanger.contract.Call(opts, &out, "feeRateForExchange", sourceCurrencyKey, destinationCurrencyKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRateForExchange is a free data retrieval call binding the contract method 0x1a5c6095.
//
// Solidity: function feeRateForExchange(bytes32 sourceCurrencyKey, bytes32 destinationCurrencyKey) view returns(uint256 exchangeFeeRate)
func (_IExchanger *IExchangerSession) FeeRateForExchange(sourceCurrencyKey [32]byte, destinationCurrencyKey [32]byte) (*big.Int, error) {
	return _IExchanger.Contract.FeeRateForExchange(&_IExchanger.CallOpts, sourceCurrencyKey, destinationCurrencyKey)
}

// FeeRateForExchange is a free data retrieval call binding the contract method 0x1a5c6095.
//
// Solidity: function feeRateForExchange(bytes32 sourceCurrencyKey, bytes32 destinationCurrencyKey) view returns(uint256 exchangeFeeRate)
func (_IExchanger *IExchangerCallerSession) FeeRateForExchange(sourceCurrencyKey [32]byte, destinationCurrencyKey [32]byte) (*big.Int, error) {
	return _IExchanger.Contract.FeeRateForExchange(&_IExchanger.CallOpts, sourceCurrencyKey, destinationCurrencyKey)
}

// GetAmountsForExchange is a free data retrieval call binding the contract method 0xf450aa34.
//
// Solidity: function getAmountsForExchange(uint256 sourceAmount, bytes32 sourceCurrencyKey, bytes32 destinationCurrencyKey) view returns(uint256 amountReceived, uint256 fee, uint256 exchangeFeeRate)
func (_IExchanger *IExchangerCaller) GetAmountsForExchange(opts *bind.CallOpts, sourceAmount *big.Int, sourceCurrencyKey [32]byte, destinationCurrencyKey [32]byte) (struct {
	AmountReceived  *big.Int
	Fee             *big.Int
	ExchangeFeeRate *big.Int
}, error) {
	var out []interface{}
	err := _IExchanger.contract.Call(opts, &out, "getAmountsForExchange", sourceAmount, sourceCurrencyKey, destinationCurrencyKey)

	outstruct := new(struct {
		AmountReceived  *big.Int
		Fee             *big.Int
		ExchangeFeeRate *big.Int
	})

	outstruct.AmountReceived = out[0].(*big.Int)
	outstruct.Fee = out[1].(*big.Int)
	outstruct.ExchangeFeeRate = out[2].(*big.Int)

	return *outstruct, err

}

// GetAmountsForExchange is a free data retrieval call binding the contract method 0xf450aa34.
//
// Solidity: function getAmountsForExchange(uint256 sourceAmount, bytes32 sourceCurrencyKey, bytes32 destinationCurrencyKey) view returns(uint256 amountReceived, uint256 fee, uint256 exchangeFeeRate)
func (_IExchanger *IExchangerSession) GetAmountsForExchange(sourceAmount *big.Int, sourceCurrencyKey [32]byte, destinationCurrencyKey [32]byte) (struct {
	AmountReceived  *big.Int
	Fee             *big.Int
	ExchangeFeeRate *big.Int
}, error) {
	return _IExchanger.Contract.GetAmountsForExchange(&_IExchanger.CallOpts, sourceAmount, sourceCurrencyKey, destinationCurrencyKey)
}

// GetAmountsForExchange is a free data retrieval call binding the contract method 0xf450aa34.
//
// Solidity: function getAmountsForExchange(uint256 sourceAmount, bytes32 sourceCurrencyKey, bytes32 destinationCurrencyKey) view returns(uint256 amountReceived, uint256 fee, uint256 exchangeFeeRate)
func (_IExchanger *IExchangerCallerSession) GetAmountsForExchange(sourceAmount *big.Int, sourceCurrencyKey [32]byte, destinationCurrencyKey [32]byte) (struct {
	AmountReceived  *big.Int
	Fee             *big.Int
	ExchangeFeeRate *big.Int
}, error) {
	return _IExchanger.Contract.GetAmountsForExchange(&_IExchanger.CallOpts, sourceAmount, sourceCurrencyKey, destinationCurrencyKey)
}

// HasWaitingPeriodOrSettlementOwing is a free data retrieval call binding the contract method 0xd6f32e06.
//
// Solidity: function hasWaitingPeriodOrSettlementOwing(address account, bytes32 currencyKey) view returns(bool)
func (_IExchanger *IExchangerCaller) HasWaitingPeriodOrSettlementOwing(opts *bind.CallOpts, account common.Address, currencyKey [32]byte) (bool, error) {
	var out []interface{}
	err := _IExchanger.contract.Call(opts, &out, "hasWaitingPeriodOrSettlementOwing", account, currencyKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasWaitingPeriodOrSettlementOwing is a free data retrieval call binding the contract method 0xd6f32e06.
//
// Solidity: function hasWaitingPeriodOrSettlementOwing(address account, bytes32 currencyKey) view returns(bool)
func (_IExchanger *IExchangerSession) HasWaitingPeriodOrSettlementOwing(account common.Address, currencyKey [32]byte) (bool, error) {
	return _IExchanger.Contract.HasWaitingPeriodOrSettlementOwing(&_IExchanger.CallOpts, account, currencyKey)
}

// HasWaitingPeriodOrSettlementOwing is a free data retrieval call binding the contract method 0xd6f32e06.
//
// Solidity: function hasWaitingPeriodOrSettlementOwing(address account, bytes32 currencyKey) view returns(bool)
func (_IExchanger *IExchangerCallerSession) HasWaitingPeriodOrSettlementOwing(account common.Address, currencyKey [32]byte) (bool, error) {
	return _IExchanger.Contract.HasWaitingPeriodOrSettlementOwing(&_IExchanger.CallOpts, account, currencyKey)
}

// IsSynthRateInvalid is a free data retrieval call binding the contract method 0x57af302c.
//
// Solidity: function isSynthRateInvalid(bytes32 currencyKey) view returns(bool)
func (_IExchanger *IExchangerCaller) IsSynthRateInvalid(opts *bind.CallOpts, currencyKey [32]byte) (bool, error) {
	var out []interface{}
	err := _IExchanger.contract.Call(opts, &out, "isSynthRateInvalid", currencyKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSynthRateInvalid is a free data retrieval call binding the contract method 0x57af302c.
//
// Solidity: function isSynthRateInvalid(bytes32 currencyKey) view returns(bool)
func (_IExchanger *IExchangerSession) IsSynthRateInvalid(currencyKey [32]byte) (bool, error) {
	return _IExchanger.Contract.IsSynthRateInvalid(&_IExchanger.CallOpts, currencyKey)
}

// IsSynthRateInvalid is a free data retrieval call binding the contract method 0x57af302c.
//
// Solidity: function isSynthRateInvalid(bytes32 currencyKey) view returns(bool)
func (_IExchanger *IExchangerCallerSession) IsSynthRateInvalid(currencyKey [32]byte) (bool, error) {
	return _IExchanger.Contract.IsSynthRateInvalid(&_IExchanger.CallOpts, currencyKey)
}

// MaxSecsLeftInWaitingPeriod is a free data retrieval call binding the contract method 0x059c29ec.
//
// Solidity: function maxSecsLeftInWaitingPeriod(address account, bytes32 currencyKey) view returns(uint256)
func (_IExchanger *IExchangerCaller) MaxSecsLeftInWaitingPeriod(opts *bind.CallOpts, account common.Address, currencyKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IExchanger.contract.Call(opts, &out, "maxSecsLeftInWaitingPeriod", account, currencyKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSecsLeftInWaitingPeriod is a free data retrieval call binding the contract method 0x059c29ec.
//
// Solidity: function maxSecsLeftInWaitingPeriod(address account, bytes32 currencyKey) view returns(uint256)
func (_IExchanger *IExchangerSession) MaxSecsLeftInWaitingPeriod(account common.Address, currencyKey [32]byte) (*big.Int, error) {
	return _IExchanger.Contract.MaxSecsLeftInWaitingPeriod(&_IExchanger.CallOpts, account, currencyKey)
}

// MaxSecsLeftInWaitingPeriod is a free data retrieval call binding the contract method 0x059c29ec.
//
// Solidity: function maxSecsLeftInWaitingPeriod(address account, bytes32 currencyKey) view returns(uint256)
func (_IExchanger *IExchangerCallerSession) MaxSecsLeftInWaitingPeriod(account common.Address, currencyKey [32]byte) (*big.Int, error) {
	return _IExchanger.Contract.MaxSecsLeftInWaitingPeriod(&_IExchanger.CallOpts, account, currencyKey)
}

// PriceDeviationThresholdFactor is a free data retrieval call binding the contract method 0x372a395a.
//
// Solidity: function priceDeviationThresholdFactor() view returns(uint256)
func (_IExchanger *IExchangerCaller) PriceDeviationThresholdFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IExchanger.contract.Call(opts, &out, "priceDeviationThresholdFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceDeviationThresholdFactor is a free data retrieval call binding the contract method 0x372a395a.
//
// Solidity: function priceDeviationThresholdFactor() view returns(uint256)
func (_IExchanger *IExchangerSession) PriceDeviationThresholdFactor() (*big.Int, error) {
	return _IExchanger.Contract.PriceDeviationThresholdFactor(&_IExchanger.CallOpts)
}

// PriceDeviationThresholdFactor is a free data retrieval call binding the contract method 0x372a395a.
//
// Solidity: function priceDeviationThresholdFactor() view returns(uint256)
func (_IExchanger *IExchangerCallerSession) PriceDeviationThresholdFactor() (*big.Int, error) {
	return _IExchanger.Contract.PriceDeviationThresholdFactor(&_IExchanger.CallOpts)
}

// SettlementOwing is a free data retrieval call binding the contract method 0x19d5c665.
//
// Solidity: function settlementOwing(address account, bytes32 currencyKey) view returns(uint256 reclaimAmount, uint256 rebateAmount, uint256 numEntries)
func (_IExchanger *IExchangerCaller) SettlementOwing(opts *bind.CallOpts, account common.Address, currencyKey [32]byte) (struct {
	ReclaimAmount *big.Int
	RebateAmount  *big.Int
	NumEntries    *big.Int
}, error) {
	var out []interface{}
	err := _IExchanger.contract.Call(opts, &out, "settlementOwing", account, currencyKey)

	outstruct := new(struct {
		ReclaimAmount *big.Int
		RebateAmount  *big.Int
		NumEntries    *big.Int
	})

	outstruct.ReclaimAmount = out[0].(*big.Int)
	outstruct.RebateAmount = out[1].(*big.Int)
	outstruct.NumEntries = out[2].(*big.Int)

	return *outstruct, err

}

// SettlementOwing is a free data retrieval call binding the contract method 0x19d5c665.
//
// Solidity: function settlementOwing(address account, bytes32 currencyKey) view returns(uint256 reclaimAmount, uint256 rebateAmount, uint256 numEntries)
func (_IExchanger *IExchangerSession) SettlementOwing(account common.Address, currencyKey [32]byte) (struct {
	ReclaimAmount *big.Int
	RebateAmount  *big.Int
	NumEntries    *big.Int
}, error) {
	return _IExchanger.Contract.SettlementOwing(&_IExchanger.CallOpts, account, currencyKey)
}

// SettlementOwing is a free data retrieval call binding the contract method 0x19d5c665.
//
// Solidity: function settlementOwing(address account, bytes32 currencyKey) view returns(uint256 reclaimAmount, uint256 rebateAmount, uint256 numEntries)
func (_IExchanger *IExchangerCallerSession) SettlementOwing(account common.Address, currencyKey [32]byte) (struct {
	ReclaimAmount *big.Int
	RebateAmount  *big.Int
	NumEntries    *big.Int
}, error) {
	return _IExchanger.Contract.SettlementOwing(&_IExchanger.CallOpts, account, currencyKey)
}

// WaitingPeriodSecs is a free data retrieval call binding the contract method 0x89257117.
//
// Solidity: function waitingPeriodSecs() view returns(uint256)
func (_IExchanger *IExchangerCaller) WaitingPeriodSecs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IExchanger.contract.Call(opts, &out, "waitingPeriodSecs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WaitingPeriodSecs is a free data retrieval call binding the contract method 0x89257117.
//
// Solidity: function waitingPeriodSecs() view returns(uint256)
func (_IExchanger *IExchangerSession) WaitingPeriodSecs() (*big.Int, error) {
	return _IExchanger.Contract.WaitingPeriodSecs(&_IExchanger.CallOpts)
}

// WaitingPeriodSecs is a free data retrieval call binding the contract method 0x89257117.
//
// Solidity: function waitingPeriodSecs() view returns(uint256)
func (_IExchanger *IExchangerCallerSession) WaitingPeriodSecs() (*big.Int, error) {
	return _IExchanger.Contract.WaitingPeriodSecs(&_IExchanger.CallOpts)
}

// Exchange is a paid mutator transaction binding the contract method 0x0a1e187d.
//
// Solidity: function exchange(address from, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address destinationAddress) returns(uint256 amountReceived)
func (_IExchanger *IExchangerTransactor) Exchange(opts *bind.TransactOpts, from common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, destinationAddress common.Address) (*types.Transaction, error) {
	return _IExchanger.contract.Transact(opts, "exchange", from, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, destinationAddress)
}

// Exchange is a paid mutator transaction binding the contract method 0x0a1e187d.
//
// Solidity: function exchange(address from, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address destinationAddress) returns(uint256 amountReceived)
func (_IExchanger *IExchangerSession) Exchange(from common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, destinationAddress common.Address) (*types.Transaction, error) {
	return _IExchanger.Contract.Exchange(&_IExchanger.TransactOpts, from, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, destinationAddress)
}

// Exchange is a paid mutator transaction binding the contract method 0x0a1e187d.
//
// Solidity: function exchange(address from, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address destinationAddress) returns(uint256 amountReceived)
func (_IExchanger *IExchangerTransactorSession) Exchange(from common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, destinationAddress common.Address) (*types.Transaction, error) {
	return _IExchanger.Contract.Exchange(&_IExchanger.TransactOpts, from, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, destinationAddress)
}

// ExchangeOnBehalf is a paid mutator transaction binding the contract method 0x6a1c4758.
//
// Solidity: function exchangeOnBehalf(address exchangeForAddress, address from, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey) returns(uint256 amountReceived)
func (_IExchanger *IExchangerTransactor) ExchangeOnBehalf(opts *bind.TransactOpts, exchangeForAddress common.Address, from common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte) (*types.Transaction, error) {
	return _IExchanger.contract.Transact(opts, "exchangeOnBehalf", exchangeForAddress, from, sourceCurrencyKey, sourceAmount, destinationCurrencyKey)
}

// ExchangeOnBehalf is a paid mutator transaction binding the contract method 0x6a1c4758.
//
// Solidity: function exchangeOnBehalf(address exchangeForAddress, address from, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey) returns(uint256 amountReceived)
func (_IExchanger *IExchangerSession) ExchangeOnBehalf(exchangeForAddress common.Address, from common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte) (*types.Transaction, error) {
	return _IExchanger.Contract.ExchangeOnBehalf(&_IExchanger.TransactOpts, exchangeForAddress, from, sourceCurrencyKey, sourceAmount, destinationCurrencyKey)
}

// ExchangeOnBehalf is a paid mutator transaction binding the contract method 0x6a1c4758.
//
// Solidity: function exchangeOnBehalf(address exchangeForAddress, address from, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey) returns(uint256 amountReceived)
func (_IExchanger *IExchangerTransactorSession) ExchangeOnBehalf(exchangeForAddress common.Address, from common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte) (*types.Transaction, error) {
	return _IExchanger.Contract.ExchangeOnBehalf(&_IExchanger.TransactOpts, exchangeForAddress, from, sourceCurrencyKey, sourceAmount, destinationCurrencyKey)
}

// ExchangeOnBehalfWithTracking is a paid mutator transaction binding the contract method 0xdfffca76.
//
// Solidity: function exchangeOnBehalfWithTracking(address exchangeForAddress, address from, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_IExchanger *IExchangerTransactor) ExchangeOnBehalfWithTracking(opts *bind.TransactOpts, exchangeForAddress common.Address, from common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _IExchanger.contract.Transact(opts, "exchangeOnBehalfWithTracking", exchangeForAddress, from, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, originator, trackingCode)
}

// ExchangeOnBehalfWithTracking is a paid mutator transaction binding the contract method 0xdfffca76.
//
// Solidity: function exchangeOnBehalfWithTracking(address exchangeForAddress, address from, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_IExchanger *IExchangerSession) ExchangeOnBehalfWithTracking(exchangeForAddress common.Address, from common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _IExchanger.Contract.ExchangeOnBehalfWithTracking(&_IExchanger.TransactOpts, exchangeForAddress, from, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, originator, trackingCode)
}

// ExchangeOnBehalfWithTracking is a paid mutator transaction binding the contract method 0xdfffca76.
//
// Solidity: function exchangeOnBehalfWithTracking(address exchangeForAddress, address from, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_IExchanger *IExchangerTransactorSession) ExchangeOnBehalfWithTracking(exchangeForAddress common.Address, from common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _IExchanger.Contract.ExchangeOnBehalfWithTracking(&_IExchanger.TransactOpts, exchangeForAddress, from, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, originator, trackingCode)
}

// ExchangeWithTracking is a paid mutator transaction binding the contract method 0x86baa45c.
//
// Solidity: function exchangeWithTracking(address from, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address destinationAddress, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_IExchanger *IExchangerTransactor) ExchangeWithTracking(opts *bind.TransactOpts, from common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, destinationAddress common.Address, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _IExchanger.contract.Transact(opts, "exchangeWithTracking", from, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, destinationAddress, originator, trackingCode)
}

// ExchangeWithTracking is a paid mutator transaction binding the contract method 0x86baa45c.
//
// Solidity: function exchangeWithTracking(address from, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address destinationAddress, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_IExchanger *IExchangerSession) ExchangeWithTracking(from common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, destinationAddress common.Address, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _IExchanger.Contract.ExchangeWithTracking(&_IExchanger.TransactOpts, from, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, destinationAddress, originator, trackingCode)
}

// ExchangeWithTracking is a paid mutator transaction binding the contract method 0x86baa45c.
//
// Solidity: function exchangeWithTracking(address from, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address destinationAddress, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_IExchanger *IExchangerTransactorSession) ExchangeWithTracking(from common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, destinationAddress common.Address, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _IExchanger.Contract.ExchangeWithTracking(&_IExchanger.TransactOpts, from, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, destinationAddress, originator, trackingCode)
}

// ExchangeWithVirtual is a paid mutator transaction binding the contract method 0xf3995224.
//
// Solidity: function exchangeWithVirtual(address from, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address destinationAddress, bytes32 trackingCode) returns(uint256 amountReceived, address vSynth)
func (_IExchanger *IExchangerTransactor) ExchangeWithVirtual(opts *bind.TransactOpts, from common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, destinationAddress common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _IExchanger.contract.Transact(opts, "exchangeWithVirtual", from, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, destinationAddress, trackingCode)
}

// ExchangeWithVirtual is a paid mutator transaction binding the contract method 0xf3995224.
//
// Solidity: function exchangeWithVirtual(address from, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address destinationAddress, bytes32 trackingCode) returns(uint256 amountReceived, address vSynth)
func (_IExchanger *IExchangerSession) ExchangeWithVirtual(from common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, destinationAddress common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _IExchanger.Contract.ExchangeWithVirtual(&_IExchanger.TransactOpts, from, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, destinationAddress, trackingCode)
}

// ExchangeWithVirtual is a paid mutator transaction binding the contract method 0xf3995224.
//
// Solidity: function exchangeWithVirtual(address from, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address destinationAddress, bytes32 trackingCode) returns(uint256 amountReceived, address vSynth)
func (_IExchanger *IExchangerTransactorSession) ExchangeWithVirtual(from common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, destinationAddress common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _IExchanger.Contract.ExchangeWithVirtual(&_IExchanger.TransactOpts, from, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, destinationAddress, trackingCode)
}

// SetLastExchangeRateForSynth is a paid mutator transaction binding the contract method 0xce096940.
//
// Solidity: function setLastExchangeRateForSynth(bytes32 currencyKey, uint256 rate) returns()
func (_IExchanger *IExchangerTransactor) SetLastExchangeRateForSynth(opts *bind.TransactOpts, currencyKey [32]byte, rate *big.Int) (*types.Transaction, error) {
	return _IExchanger.contract.Transact(opts, "setLastExchangeRateForSynth", currencyKey, rate)
}

// SetLastExchangeRateForSynth is a paid mutator transaction binding the contract method 0xce096940.
//
// Solidity: function setLastExchangeRateForSynth(bytes32 currencyKey, uint256 rate) returns()
func (_IExchanger *IExchangerSession) SetLastExchangeRateForSynth(currencyKey [32]byte, rate *big.Int) (*types.Transaction, error) {
	return _IExchanger.Contract.SetLastExchangeRateForSynth(&_IExchanger.TransactOpts, currencyKey, rate)
}

// SetLastExchangeRateForSynth is a paid mutator transaction binding the contract method 0xce096940.
//
// Solidity: function setLastExchangeRateForSynth(bytes32 currencyKey, uint256 rate) returns()
func (_IExchanger *IExchangerTransactorSession) SetLastExchangeRateForSynth(currencyKey [32]byte, rate *big.Int) (*types.Transaction, error) {
	return _IExchanger.Contract.SetLastExchangeRateForSynth(&_IExchanger.TransactOpts, currencyKey, rate)
}

// Settle is a paid mutator transaction binding the contract method 0x1b16802c.
//
// Solidity: function settle(address from, bytes32 currencyKey) returns(uint256 reclaimed, uint256 refunded, uint256 numEntries)
func (_IExchanger *IExchangerTransactor) Settle(opts *bind.TransactOpts, from common.Address, currencyKey [32]byte) (*types.Transaction, error) {
	return _IExchanger.contract.Transact(opts, "settle", from, currencyKey)
}

// Settle is a paid mutator transaction binding the contract method 0x1b16802c.
//
// Solidity: function settle(address from, bytes32 currencyKey) returns(uint256 reclaimed, uint256 refunded, uint256 numEntries)
func (_IExchanger *IExchangerSession) Settle(from common.Address, currencyKey [32]byte) (*types.Transaction, error) {
	return _IExchanger.Contract.Settle(&_IExchanger.TransactOpts, from, currencyKey)
}

// Settle is a paid mutator transaction binding the contract method 0x1b16802c.
//
// Solidity: function settle(address from, bytes32 currencyKey) returns(uint256 reclaimed, uint256 refunded, uint256 numEntries)
func (_IExchanger *IExchangerTransactorSession) Settle(from common.Address, currencyKey [32]byte) (*types.Transaction, error) {
	return _IExchanger.Contract.Settle(&_IExchanger.TransactOpts, from, currencyKey)
}

// SuspendSynthWithInvalidRate is a paid mutator transaction binding the contract method 0x0b9e31c9.
//
// Solidity: function suspendSynthWithInvalidRate(bytes32 currencyKey) returns()
func (_IExchanger *IExchangerTransactor) SuspendSynthWithInvalidRate(opts *bind.TransactOpts, currencyKey [32]byte) (*types.Transaction, error) {
	return _IExchanger.contract.Transact(opts, "suspendSynthWithInvalidRate", currencyKey)
}

// SuspendSynthWithInvalidRate is a paid mutator transaction binding the contract method 0x0b9e31c9.
//
// Solidity: function suspendSynthWithInvalidRate(bytes32 currencyKey) returns()
func (_IExchanger *IExchangerSession) SuspendSynthWithInvalidRate(currencyKey [32]byte) (*types.Transaction, error) {
	return _IExchanger.Contract.SuspendSynthWithInvalidRate(&_IExchanger.TransactOpts, currencyKey)
}

// SuspendSynthWithInvalidRate is a paid mutator transaction binding the contract method 0x0b9e31c9.
//
// Solidity: function suspendSynthWithInvalidRate(bytes32 currencyKey) returns()
func (_IExchanger *IExchangerTransactorSession) SuspendSynthWithInvalidRate(currencyKey [32]byte) (*types.Transaction, error) {
	return _IExchanger.Contract.SuspendSynthWithInvalidRate(&_IExchanger.TransactOpts, currencyKey)
}

// IIssuerABI is the input ABI used to generate the binding from.
const IIssuerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"anySynthOrSNXRateIsInvalid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"anyRateInvalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableCurrencyKeys\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableSynthCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"availableSynths\",\"outputs\":[{\"internalType\":\"contractISynth\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnSynths\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnForAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnSynthsOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"burnSynthsToTarget\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnForAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"burnSynthsToTargetOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"canBurnSynths\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"collateralisationRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"}],\"name\":\"collateralisationRatioAndAnyRatesInvalid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cratio\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"anyRateIsInvalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"debtBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"debtBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"currencyKeys\",\"type\":\"bytes32[]\"}],\"name\":\"getSynths\",\"outputs\":[{\"internalType\":\"contractISynth[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuanceRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"issueMaxSynths\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issueFor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"issueMaxSynthsOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issueSynths\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issueFor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issueSynthsOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"lastIssueEvent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"susdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"}],\"name\":\"liquidateDelinquentAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalRedeemed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToLiquidate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"maxIssuableSynths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxIssuable\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumStakeTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"remainingIssuableSynths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxIssuable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alreadyIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSystemDebt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"synths\",\"outputs\":[{\"internalType\":\"contractISynth\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"synthAddress\",\"type\":\"address\"}],\"name\":\"synthsByAddress\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"excludeEtherCollateral\",\"type\":\"bool\"}],\"name\":\"totalIssuedSynths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"transferableSynthetixAndAnyRateIsInvalid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"transferable\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"anyRateIsInvalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IIssuerFuncSigs maps the 4-byte function signature to its string representation.
var IIssuerFuncSigs = map[string]string{
	"4e99bda9": "anySynthOrSNXRateIsInvalid()",
	"72cb051f": "availableCurrencyKeys()",
	"dbf63340": "availableSynthCount()",
	"835e119c": "availableSynths(uint256)",
	"b06e8c65": "burnSynths(address,uint256)",
	"9a5154b4": "burnSynthsOnBehalf(address,address,uint256)",
	"497d704a": "burnSynthsToTarget(address)",
	"2b3f41aa": "burnSynthsToTargetOnBehalf(address,address)",
	"bff4fdfc": "canBurnSynths(address)",
	"a5fdc5de": "collateral(address)",
	"a311c7c2": "collateralisationRatio(address)",
	"ae3bbbbb": "collateralisationRatioAndAnyRatesInvalid(address)",
	"d37c4d8b": "debtBalanceOf(address,bytes32)",
	"3b6afe40": "getSynths(bytes32[])",
	"b410a034": "issuanceRatio()",
	"c8977132": "issueMaxSynths(address)",
	"fd864ccf": "issueMaxSynthsOnBehalf(address,address)",
	"042e0688": "issueSynths(address,uint256)",
	"44ec6b62": "issueSynthsOnBehalf(address,address,uint256)",
	"dd3d2b2e": "lastIssueEvent(address)",
	"a63c4df4": "liquidateDelinquentAccount(address,uint256,address)",
	"05b3c1c9": "maxIssuableSynths(address)",
	"242df9e1": "minimumStakeTime()",
	"1137aedf": "remainingIssuableSynths(address)",
	"32608039": "synths(bytes32)",
	"16b2213f": "synthsByAddress(address)",
	"7b1001b7": "totalIssuedSynths(bytes32,bool)",
	"6bed0415": "transferableSynthetixAndAnyRateIsInvalid(address,uint256)",
}

// IIssuer is an auto generated Go binding around an Ethereum contract.
type IIssuer struct {
	IIssuerCaller     // Read-only binding to the contract
	IIssuerTransactor // Write-only binding to the contract
	IIssuerFilterer   // Log filterer for contract events
}

// IIssuerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IIssuerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIssuerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IIssuerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIssuerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IIssuerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIssuerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IIssuerSession struct {
	Contract     *IIssuer          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IIssuerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IIssuerCallerSession struct {
	Contract *IIssuerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IIssuerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IIssuerTransactorSession struct {
	Contract     *IIssuerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IIssuerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IIssuerRaw struct {
	Contract *IIssuer // Generic contract binding to access the raw methods on
}

// IIssuerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IIssuerCallerRaw struct {
	Contract *IIssuerCaller // Generic read-only contract binding to access the raw methods on
}

// IIssuerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IIssuerTransactorRaw struct {
	Contract *IIssuerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIIssuer creates a new instance of IIssuer, bound to a specific deployed contract.
func NewIIssuer(address common.Address, backend bind.ContractBackend) (*IIssuer, error) {
	contract, err := bindIIssuer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IIssuer{IIssuerCaller: IIssuerCaller{contract: contract}, IIssuerTransactor: IIssuerTransactor{contract: contract}, IIssuerFilterer: IIssuerFilterer{contract: contract}}, nil
}

// NewIIssuerCaller creates a new read-only instance of IIssuer, bound to a specific deployed contract.
func NewIIssuerCaller(address common.Address, caller bind.ContractCaller) (*IIssuerCaller, error) {
	contract, err := bindIIssuer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IIssuerCaller{contract: contract}, nil
}

// NewIIssuerTransactor creates a new write-only instance of IIssuer, bound to a specific deployed contract.
func NewIIssuerTransactor(address common.Address, transactor bind.ContractTransactor) (*IIssuerTransactor, error) {
	contract, err := bindIIssuer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IIssuerTransactor{contract: contract}, nil
}

// NewIIssuerFilterer creates a new log filterer instance of IIssuer, bound to a specific deployed contract.
func NewIIssuerFilterer(address common.Address, filterer bind.ContractFilterer) (*IIssuerFilterer, error) {
	contract, err := bindIIssuer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IIssuerFilterer{contract: contract}, nil
}

// bindIIssuer binds a generic wrapper to an already deployed contract.
func bindIIssuer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IIssuerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IIssuer *IIssuerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IIssuer.Contract.IIssuerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IIssuer *IIssuerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IIssuer.Contract.IIssuerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IIssuer *IIssuerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IIssuer.Contract.IIssuerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IIssuer *IIssuerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IIssuer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IIssuer *IIssuerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IIssuer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IIssuer *IIssuerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IIssuer.Contract.contract.Transact(opts, method, params...)
}

// AnySynthOrSNXRateIsInvalid is a free data retrieval call binding the contract method 0x4e99bda9.
//
// Solidity: function anySynthOrSNXRateIsInvalid() view returns(bool anyRateInvalid)
func (_IIssuer *IIssuerCaller) AnySynthOrSNXRateIsInvalid(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "anySynthOrSNXRateIsInvalid")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AnySynthOrSNXRateIsInvalid is a free data retrieval call binding the contract method 0x4e99bda9.
//
// Solidity: function anySynthOrSNXRateIsInvalid() view returns(bool anyRateInvalid)
func (_IIssuer *IIssuerSession) AnySynthOrSNXRateIsInvalid() (bool, error) {
	return _IIssuer.Contract.AnySynthOrSNXRateIsInvalid(&_IIssuer.CallOpts)
}

// AnySynthOrSNXRateIsInvalid is a free data retrieval call binding the contract method 0x4e99bda9.
//
// Solidity: function anySynthOrSNXRateIsInvalid() view returns(bool anyRateInvalid)
func (_IIssuer *IIssuerCallerSession) AnySynthOrSNXRateIsInvalid() (bool, error) {
	return _IIssuer.Contract.AnySynthOrSNXRateIsInvalid(&_IIssuer.CallOpts)
}

// AvailableCurrencyKeys is a free data retrieval call binding the contract method 0x72cb051f.
//
// Solidity: function availableCurrencyKeys() view returns(bytes32[])
func (_IIssuer *IIssuerCaller) AvailableCurrencyKeys(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "availableCurrencyKeys")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// AvailableCurrencyKeys is a free data retrieval call binding the contract method 0x72cb051f.
//
// Solidity: function availableCurrencyKeys() view returns(bytes32[])
func (_IIssuer *IIssuerSession) AvailableCurrencyKeys() ([][32]byte, error) {
	return _IIssuer.Contract.AvailableCurrencyKeys(&_IIssuer.CallOpts)
}

// AvailableCurrencyKeys is a free data retrieval call binding the contract method 0x72cb051f.
//
// Solidity: function availableCurrencyKeys() view returns(bytes32[])
func (_IIssuer *IIssuerCallerSession) AvailableCurrencyKeys() ([][32]byte, error) {
	return _IIssuer.Contract.AvailableCurrencyKeys(&_IIssuer.CallOpts)
}

// AvailableSynthCount is a free data retrieval call binding the contract method 0xdbf63340.
//
// Solidity: function availableSynthCount() view returns(uint256)
func (_IIssuer *IIssuerCaller) AvailableSynthCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "availableSynthCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableSynthCount is a free data retrieval call binding the contract method 0xdbf63340.
//
// Solidity: function availableSynthCount() view returns(uint256)
func (_IIssuer *IIssuerSession) AvailableSynthCount() (*big.Int, error) {
	return _IIssuer.Contract.AvailableSynthCount(&_IIssuer.CallOpts)
}

// AvailableSynthCount is a free data retrieval call binding the contract method 0xdbf63340.
//
// Solidity: function availableSynthCount() view returns(uint256)
func (_IIssuer *IIssuerCallerSession) AvailableSynthCount() (*big.Int, error) {
	return _IIssuer.Contract.AvailableSynthCount(&_IIssuer.CallOpts)
}

// AvailableSynths is a free data retrieval call binding the contract method 0x835e119c.
//
// Solidity: function availableSynths(uint256 index) view returns(address)
func (_IIssuer *IIssuerCaller) AvailableSynths(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "availableSynths", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvailableSynths is a free data retrieval call binding the contract method 0x835e119c.
//
// Solidity: function availableSynths(uint256 index) view returns(address)
func (_IIssuer *IIssuerSession) AvailableSynths(index *big.Int) (common.Address, error) {
	return _IIssuer.Contract.AvailableSynths(&_IIssuer.CallOpts, index)
}

// AvailableSynths is a free data retrieval call binding the contract method 0x835e119c.
//
// Solidity: function availableSynths(uint256 index) view returns(address)
func (_IIssuer *IIssuerCallerSession) AvailableSynths(index *big.Int) (common.Address, error) {
	return _IIssuer.Contract.AvailableSynths(&_IIssuer.CallOpts, index)
}

// CanBurnSynths is a free data retrieval call binding the contract method 0xbff4fdfc.
//
// Solidity: function canBurnSynths(address account) view returns(bool)
func (_IIssuer *IIssuerCaller) CanBurnSynths(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "canBurnSynths", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanBurnSynths is a free data retrieval call binding the contract method 0xbff4fdfc.
//
// Solidity: function canBurnSynths(address account) view returns(bool)
func (_IIssuer *IIssuerSession) CanBurnSynths(account common.Address) (bool, error) {
	return _IIssuer.Contract.CanBurnSynths(&_IIssuer.CallOpts, account)
}

// CanBurnSynths is a free data retrieval call binding the contract method 0xbff4fdfc.
//
// Solidity: function canBurnSynths(address account) view returns(bool)
func (_IIssuer *IIssuerCallerSession) CanBurnSynths(account common.Address) (bool, error) {
	return _IIssuer.Contract.CanBurnSynths(&_IIssuer.CallOpts, account)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address account) view returns(uint256)
func (_IIssuer *IIssuerCaller) Collateral(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "collateral", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address account) view returns(uint256)
func (_IIssuer *IIssuerSession) Collateral(account common.Address) (*big.Int, error) {
	return _IIssuer.Contract.Collateral(&_IIssuer.CallOpts, account)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address account) view returns(uint256)
func (_IIssuer *IIssuerCallerSession) Collateral(account common.Address) (*big.Int, error) {
	return _IIssuer.Contract.Collateral(&_IIssuer.CallOpts, account)
}

// CollateralisationRatio is a free data retrieval call binding the contract method 0xa311c7c2.
//
// Solidity: function collateralisationRatio(address issuer) view returns(uint256)
func (_IIssuer *IIssuerCaller) CollateralisationRatio(opts *bind.CallOpts, issuer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "collateralisationRatio", issuer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralisationRatio is a free data retrieval call binding the contract method 0xa311c7c2.
//
// Solidity: function collateralisationRatio(address issuer) view returns(uint256)
func (_IIssuer *IIssuerSession) CollateralisationRatio(issuer common.Address) (*big.Int, error) {
	return _IIssuer.Contract.CollateralisationRatio(&_IIssuer.CallOpts, issuer)
}

// CollateralisationRatio is a free data retrieval call binding the contract method 0xa311c7c2.
//
// Solidity: function collateralisationRatio(address issuer) view returns(uint256)
func (_IIssuer *IIssuerCallerSession) CollateralisationRatio(issuer common.Address) (*big.Int, error) {
	return _IIssuer.Contract.CollateralisationRatio(&_IIssuer.CallOpts, issuer)
}

// CollateralisationRatioAndAnyRatesInvalid is a free data retrieval call binding the contract method 0xae3bbbbb.
//
// Solidity: function collateralisationRatioAndAnyRatesInvalid(address _issuer) view returns(uint256 cratio, bool anyRateIsInvalid)
func (_IIssuer *IIssuerCaller) CollateralisationRatioAndAnyRatesInvalid(opts *bind.CallOpts, _issuer common.Address) (struct {
	Cratio           *big.Int
	AnyRateIsInvalid bool
}, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "collateralisationRatioAndAnyRatesInvalid", _issuer)

	outstruct := new(struct {
		Cratio           *big.Int
		AnyRateIsInvalid bool
	})

	outstruct.Cratio = out[0].(*big.Int)
	outstruct.AnyRateIsInvalid = out[1].(bool)

	return *outstruct, err

}

// CollateralisationRatioAndAnyRatesInvalid is a free data retrieval call binding the contract method 0xae3bbbbb.
//
// Solidity: function collateralisationRatioAndAnyRatesInvalid(address _issuer) view returns(uint256 cratio, bool anyRateIsInvalid)
func (_IIssuer *IIssuerSession) CollateralisationRatioAndAnyRatesInvalid(_issuer common.Address) (struct {
	Cratio           *big.Int
	AnyRateIsInvalid bool
}, error) {
	return _IIssuer.Contract.CollateralisationRatioAndAnyRatesInvalid(&_IIssuer.CallOpts, _issuer)
}

// CollateralisationRatioAndAnyRatesInvalid is a free data retrieval call binding the contract method 0xae3bbbbb.
//
// Solidity: function collateralisationRatioAndAnyRatesInvalid(address _issuer) view returns(uint256 cratio, bool anyRateIsInvalid)
func (_IIssuer *IIssuerCallerSession) CollateralisationRatioAndAnyRatesInvalid(_issuer common.Address) (struct {
	Cratio           *big.Int
	AnyRateIsInvalid bool
}, error) {
	return _IIssuer.Contract.CollateralisationRatioAndAnyRatesInvalid(&_IIssuer.CallOpts, _issuer)
}

// DebtBalanceOf is a free data retrieval call binding the contract method 0xd37c4d8b.
//
// Solidity: function debtBalanceOf(address issuer, bytes32 currencyKey) view returns(uint256 debtBalance)
func (_IIssuer *IIssuerCaller) DebtBalanceOf(opts *bind.CallOpts, issuer common.Address, currencyKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "debtBalanceOf", issuer, currencyKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtBalanceOf is a free data retrieval call binding the contract method 0xd37c4d8b.
//
// Solidity: function debtBalanceOf(address issuer, bytes32 currencyKey) view returns(uint256 debtBalance)
func (_IIssuer *IIssuerSession) DebtBalanceOf(issuer common.Address, currencyKey [32]byte) (*big.Int, error) {
	return _IIssuer.Contract.DebtBalanceOf(&_IIssuer.CallOpts, issuer, currencyKey)
}

// DebtBalanceOf is a free data retrieval call binding the contract method 0xd37c4d8b.
//
// Solidity: function debtBalanceOf(address issuer, bytes32 currencyKey) view returns(uint256 debtBalance)
func (_IIssuer *IIssuerCallerSession) DebtBalanceOf(issuer common.Address, currencyKey [32]byte) (*big.Int, error) {
	return _IIssuer.Contract.DebtBalanceOf(&_IIssuer.CallOpts, issuer, currencyKey)
}

// GetSynths is a free data retrieval call binding the contract method 0x3b6afe40.
//
// Solidity: function getSynths(bytes32[] currencyKeys) view returns(address[])
func (_IIssuer *IIssuerCaller) GetSynths(opts *bind.CallOpts, currencyKeys [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "getSynths", currencyKeys)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSynths is a free data retrieval call binding the contract method 0x3b6afe40.
//
// Solidity: function getSynths(bytes32[] currencyKeys) view returns(address[])
func (_IIssuer *IIssuerSession) GetSynths(currencyKeys [][32]byte) ([]common.Address, error) {
	return _IIssuer.Contract.GetSynths(&_IIssuer.CallOpts, currencyKeys)
}

// GetSynths is a free data retrieval call binding the contract method 0x3b6afe40.
//
// Solidity: function getSynths(bytes32[] currencyKeys) view returns(address[])
func (_IIssuer *IIssuerCallerSession) GetSynths(currencyKeys [][32]byte) ([]common.Address, error) {
	return _IIssuer.Contract.GetSynths(&_IIssuer.CallOpts, currencyKeys)
}

// IssuanceRatio is a free data retrieval call binding the contract method 0xb410a034.
//
// Solidity: function issuanceRatio() view returns(uint256)
func (_IIssuer *IIssuerCaller) IssuanceRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "issuanceRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IssuanceRatio is a free data retrieval call binding the contract method 0xb410a034.
//
// Solidity: function issuanceRatio() view returns(uint256)
func (_IIssuer *IIssuerSession) IssuanceRatio() (*big.Int, error) {
	return _IIssuer.Contract.IssuanceRatio(&_IIssuer.CallOpts)
}

// IssuanceRatio is a free data retrieval call binding the contract method 0xb410a034.
//
// Solidity: function issuanceRatio() view returns(uint256)
func (_IIssuer *IIssuerCallerSession) IssuanceRatio() (*big.Int, error) {
	return _IIssuer.Contract.IssuanceRatio(&_IIssuer.CallOpts)
}

// LastIssueEvent is a free data retrieval call binding the contract method 0xdd3d2b2e.
//
// Solidity: function lastIssueEvent(address account) view returns(uint256)
func (_IIssuer *IIssuerCaller) LastIssueEvent(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "lastIssueEvent", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastIssueEvent is a free data retrieval call binding the contract method 0xdd3d2b2e.
//
// Solidity: function lastIssueEvent(address account) view returns(uint256)
func (_IIssuer *IIssuerSession) LastIssueEvent(account common.Address) (*big.Int, error) {
	return _IIssuer.Contract.LastIssueEvent(&_IIssuer.CallOpts, account)
}

// LastIssueEvent is a free data retrieval call binding the contract method 0xdd3d2b2e.
//
// Solidity: function lastIssueEvent(address account) view returns(uint256)
func (_IIssuer *IIssuerCallerSession) LastIssueEvent(account common.Address) (*big.Int, error) {
	return _IIssuer.Contract.LastIssueEvent(&_IIssuer.CallOpts, account)
}

// MaxIssuableSynths is a free data retrieval call binding the contract method 0x05b3c1c9.
//
// Solidity: function maxIssuableSynths(address issuer) view returns(uint256 maxIssuable)
func (_IIssuer *IIssuerCaller) MaxIssuableSynths(opts *bind.CallOpts, issuer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "maxIssuableSynths", issuer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxIssuableSynths is a free data retrieval call binding the contract method 0x05b3c1c9.
//
// Solidity: function maxIssuableSynths(address issuer) view returns(uint256 maxIssuable)
func (_IIssuer *IIssuerSession) MaxIssuableSynths(issuer common.Address) (*big.Int, error) {
	return _IIssuer.Contract.MaxIssuableSynths(&_IIssuer.CallOpts, issuer)
}

// MaxIssuableSynths is a free data retrieval call binding the contract method 0x05b3c1c9.
//
// Solidity: function maxIssuableSynths(address issuer) view returns(uint256 maxIssuable)
func (_IIssuer *IIssuerCallerSession) MaxIssuableSynths(issuer common.Address) (*big.Int, error) {
	return _IIssuer.Contract.MaxIssuableSynths(&_IIssuer.CallOpts, issuer)
}

// MinimumStakeTime is a free data retrieval call binding the contract method 0x242df9e1.
//
// Solidity: function minimumStakeTime() view returns(uint256)
func (_IIssuer *IIssuerCaller) MinimumStakeTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "minimumStakeTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStakeTime is a free data retrieval call binding the contract method 0x242df9e1.
//
// Solidity: function minimumStakeTime() view returns(uint256)
func (_IIssuer *IIssuerSession) MinimumStakeTime() (*big.Int, error) {
	return _IIssuer.Contract.MinimumStakeTime(&_IIssuer.CallOpts)
}

// MinimumStakeTime is a free data retrieval call binding the contract method 0x242df9e1.
//
// Solidity: function minimumStakeTime() view returns(uint256)
func (_IIssuer *IIssuerCallerSession) MinimumStakeTime() (*big.Int, error) {
	return _IIssuer.Contract.MinimumStakeTime(&_IIssuer.CallOpts)
}

// RemainingIssuableSynths is a free data retrieval call binding the contract method 0x1137aedf.
//
// Solidity: function remainingIssuableSynths(address issuer) view returns(uint256 maxIssuable, uint256 alreadyIssued, uint256 totalSystemDebt)
func (_IIssuer *IIssuerCaller) RemainingIssuableSynths(opts *bind.CallOpts, issuer common.Address) (struct {
	MaxIssuable     *big.Int
	AlreadyIssued   *big.Int
	TotalSystemDebt *big.Int
}, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "remainingIssuableSynths", issuer)

	outstruct := new(struct {
		MaxIssuable     *big.Int
		AlreadyIssued   *big.Int
		TotalSystemDebt *big.Int
	})

	outstruct.MaxIssuable = out[0].(*big.Int)
	outstruct.AlreadyIssued = out[1].(*big.Int)
	outstruct.TotalSystemDebt = out[2].(*big.Int)

	return *outstruct, err

}

// RemainingIssuableSynths is a free data retrieval call binding the contract method 0x1137aedf.
//
// Solidity: function remainingIssuableSynths(address issuer) view returns(uint256 maxIssuable, uint256 alreadyIssued, uint256 totalSystemDebt)
func (_IIssuer *IIssuerSession) RemainingIssuableSynths(issuer common.Address) (struct {
	MaxIssuable     *big.Int
	AlreadyIssued   *big.Int
	TotalSystemDebt *big.Int
}, error) {
	return _IIssuer.Contract.RemainingIssuableSynths(&_IIssuer.CallOpts, issuer)
}

// RemainingIssuableSynths is a free data retrieval call binding the contract method 0x1137aedf.
//
// Solidity: function remainingIssuableSynths(address issuer) view returns(uint256 maxIssuable, uint256 alreadyIssued, uint256 totalSystemDebt)
func (_IIssuer *IIssuerCallerSession) RemainingIssuableSynths(issuer common.Address) (struct {
	MaxIssuable     *big.Int
	AlreadyIssued   *big.Int
	TotalSystemDebt *big.Int
}, error) {
	return _IIssuer.Contract.RemainingIssuableSynths(&_IIssuer.CallOpts, issuer)
}

// Synths is a free data retrieval call binding the contract method 0x32608039.
//
// Solidity: function synths(bytes32 currencyKey) view returns(address)
func (_IIssuer *IIssuerCaller) Synths(opts *bind.CallOpts, currencyKey [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "synths", currencyKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Synths is a free data retrieval call binding the contract method 0x32608039.
//
// Solidity: function synths(bytes32 currencyKey) view returns(address)
func (_IIssuer *IIssuerSession) Synths(currencyKey [32]byte) (common.Address, error) {
	return _IIssuer.Contract.Synths(&_IIssuer.CallOpts, currencyKey)
}

// Synths is a free data retrieval call binding the contract method 0x32608039.
//
// Solidity: function synths(bytes32 currencyKey) view returns(address)
func (_IIssuer *IIssuerCallerSession) Synths(currencyKey [32]byte) (common.Address, error) {
	return _IIssuer.Contract.Synths(&_IIssuer.CallOpts, currencyKey)
}

// SynthsByAddress is a free data retrieval call binding the contract method 0x16b2213f.
//
// Solidity: function synthsByAddress(address synthAddress) view returns(bytes32)
func (_IIssuer *IIssuerCaller) SynthsByAddress(opts *bind.CallOpts, synthAddress common.Address) ([32]byte, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "synthsByAddress", synthAddress)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SynthsByAddress is a free data retrieval call binding the contract method 0x16b2213f.
//
// Solidity: function synthsByAddress(address synthAddress) view returns(bytes32)
func (_IIssuer *IIssuerSession) SynthsByAddress(synthAddress common.Address) ([32]byte, error) {
	return _IIssuer.Contract.SynthsByAddress(&_IIssuer.CallOpts, synthAddress)
}

// SynthsByAddress is a free data retrieval call binding the contract method 0x16b2213f.
//
// Solidity: function synthsByAddress(address synthAddress) view returns(bytes32)
func (_IIssuer *IIssuerCallerSession) SynthsByAddress(synthAddress common.Address) ([32]byte, error) {
	return _IIssuer.Contract.SynthsByAddress(&_IIssuer.CallOpts, synthAddress)
}

// TotalIssuedSynths is a free data retrieval call binding the contract method 0x7b1001b7.
//
// Solidity: function totalIssuedSynths(bytes32 currencyKey, bool excludeEtherCollateral) view returns(uint256)
func (_IIssuer *IIssuerCaller) TotalIssuedSynths(opts *bind.CallOpts, currencyKey [32]byte, excludeEtherCollateral bool) (*big.Int, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "totalIssuedSynths", currencyKey, excludeEtherCollateral)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalIssuedSynths is a free data retrieval call binding the contract method 0x7b1001b7.
//
// Solidity: function totalIssuedSynths(bytes32 currencyKey, bool excludeEtherCollateral) view returns(uint256)
func (_IIssuer *IIssuerSession) TotalIssuedSynths(currencyKey [32]byte, excludeEtherCollateral bool) (*big.Int, error) {
	return _IIssuer.Contract.TotalIssuedSynths(&_IIssuer.CallOpts, currencyKey, excludeEtherCollateral)
}

// TotalIssuedSynths is a free data retrieval call binding the contract method 0x7b1001b7.
//
// Solidity: function totalIssuedSynths(bytes32 currencyKey, bool excludeEtherCollateral) view returns(uint256)
func (_IIssuer *IIssuerCallerSession) TotalIssuedSynths(currencyKey [32]byte, excludeEtherCollateral bool) (*big.Int, error) {
	return _IIssuer.Contract.TotalIssuedSynths(&_IIssuer.CallOpts, currencyKey, excludeEtherCollateral)
}

// TransferableSynthetixAndAnyRateIsInvalid is a free data retrieval call binding the contract method 0x6bed0415.
//
// Solidity: function transferableSynthetixAndAnyRateIsInvalid(address account, uint256 balance) view returns(uint256 transferable, bool anyRateIsInvalid)
func (_IIssuer *IIssuerCaller) TransferableSynthetixAndAnyRateIsInvalid(opts *bind.CallOpts, account common.Address, balance *big.Int) (struct {
	Transferable     *big.Int
	AnyRateIsInvalid bool
}, error) {
	var out []interface{}
	err := _IIssuer.contract.Call(opts, &out, "transferableSynthetixAndAnyRateIsInvalid", account, balance)

	outstruct := new(struct {
		Transferable     *big.Int
		AnyRateIsInvalid bool
	})

	outstruct.Transferable = out[0].(*big.Int)
	outstruct.AnyRateIsInvalid = out[1].(bool)

	return *outstruct, err

}

// TransferableSynthetixAndAnyRateIsInvalid is a free data retrieval call binding the contract method 0x6bed0415.
//
// Solidity: function transferableSynthetixAndAnyRateIsInvalid(address account, uint256 balance) view returns(uint256 transferable, bool anyRateIsInvalid)
func (_IIssuer *IIssuerSession) TransferableSynthetixAndAnyRateIsInvalid(account common.Address, balance *big.Int) (struct {
	Transferable     *big.Int
	AnyRateIsInvalid bool
}, error) {
	return _IIssuer.Contract.TransferableSynthetixAndAnyRateIsInvalid(&_IIssuer.CallOpts, account, balance)
}

// TransferableSynthetixAndAnyRateIsInvalid is a free data retrieval call binding the contract method 0x6bed0415.
//
// Solidity: function transferableSynthetixAndAnyRateIsInvalid(address account, uint256 balance) view returns(uint256 transferable, bool anyRateIsInvalid)
func (_IIssuer *IIssuerCallerSession) TransferableSynthetixAndAnyRateIsInvalid(account common.Address, balance *big.Int) (struct {
	Transferable     *big.Int
	AnyRateIsInvalid bool
}, error) {
	return _IIssuer.Contract.TransferableSynthetixAndAnyRateIsInvalid(&_IIssuer.CallOpts, account, balance)
}

// BurnSynths is a paid mutator transaction binding the contract method 0xb06e8c65.
//
// Solidity: function burnSynths(address from, uint256 amount) returns()
func (_IIssuer *IIssuerTransactor) BurnSynths(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IIssuer.contract.Transact(opts, "burnSynths", from, amount)
}

// BurnSynths is a paid mutator transaction binding the contract method 0xb06e8c65.
//
// Solidity: function burnSynths(address from, uint256 amount) returns()
func (_IIssuer *IIssuerSession) BurnSynths(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IIssuer.Contract.BurnSynths(&_IIssuer.TransactOpts, from, amount)
}

// BurnSynths is a paid mutator transaction binding the contract method 0xb06e8c65.
//
// Solidity: function burnSynths(address from, uint256 amount) returns()
func (_IIssuer *IIssuerTransactorSession) BurnSynths(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IIssuer.Contract.BurnSynths(&_IIssuer.TransactOpts, from, amount)
}

// BurnSynthsOnBehalf is a paid mutator transaction binding the contract method 0x9a5154b4.
//
// Solidity: function burnSynthsOnBehalf(address burnForAddress, address from, uint256 amount) returns()
func (_IIssuer *IIssuerTransactor) BurnSynthsOnBehalf(opts *bind.TransactOpts, burnForAddress common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IIssuer.contract.Transact(opts, "burnSynthsOnBehalf", burnForAddress, from, amount)
}

// BurnSynthsOnBehalf is a paid mutator transaction binding the contract method 0x9a5154b4.
//
// Solidity: function burnSynthsOnBehalf(address burnForAddress, address from, uint256 amount) returns()
func (_IIssuer *IIssuerSession) BurnSynthsOnBehalf(burnForAddress common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IIssuer.Contract.BurnSynthsOnBehalf(&_IIssuer.TransactOpts, burnForAddress, from, amount)
}

// BurnSynthsOnBehalf is a paid mutator transaction binding the contract method 0x9a5154b4.
//
// Solidity: function burnSynthsOnBehalf(address burnForAddress, address from, uint256 amount) returns()
func (_IIssuer *IIssuerTransactorSession) BurnSynthsOnBehalf(burnForAddress common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IIssuer.Contract.BurnSynthsOnBehalf(&_IIssuer.TransactOpts, burnForAddress, from, amount)
}

// BurnSynthsToTarget is a paid mutator transaction binding the contract method 0x497d704a.
//
// Solidity: function burnSynthsToTarget(address from) returns()
func (_IIssuer *IIssuerTransactor) BurnSynthsToTarget(opts *bind.TransactOpts, from common.Address) (*types.Transaction, error) {
	return _IIssuer.contract.Transact(opts, "burnSynthsToTarget", from)
}

// BurnSynthsToTarget is a paid mutator transaction binding the contract method 0x497d704a.
//
// Solidity: function burnSynthsToTarget(address from) returns()
func (_IIssuer *IIssuerSession) BurnSynthsToTarget(from common.Address) (*types.Transaction, error) {
	return _IIssuer.Contract.BurnSynthsToTarget(&_IIssuer.TransactOpts, from)
}

// BurnSynthsToTarget is a paid mutator transaction binding the contract method 0x497d704a.
//
// Solidity: function burnSynthsToTarget(address from) returns()
func (_IIssuer *IIssuerTransactorSession) BurnSynthsToTarget(from common.Address) (*types.Transaction, error) {
	return _IIssuer.Contract.BurnSynthsToTarget(&_IIssuer.TransactOpts, from)
}

// BurnSynthsToTargetOnBehalf is a paid mutator transaction binding the contract method 0x2b3f41aa.
//
// Solidity: function burnSynthsToTargetOnBehalf(address burnForAddress, address from) returns()
func (_IIssuer *IIssuerTransactor) BurnSynthsToTargetOnBehalf(opts *bind.TransactOpts, burnForAddress common.Address, from common.Address) (*types.Transaction, error) {
	return _IIssuer.contract.Transact(opts, "burnSynthsToTargetOnBehalf", burnForAddress, from)
}

// BurnSynthsToTargetOnBehalf is a paid mutator transaction binding the contract method 0x2b3f41aa.
//
// Solidity: function burnSynthsToTargetOnBehalf(address burnForAddress, address from) returns()
func (_IIssuer *IIssuerSession) BurnSynthsToTargetOnBehalf(burnForAddress common.Address, from common.Address) (*types.Transaction, error) {
	return _IIssuer.Contract.BurnSynthsToTargetOnBehalf(&_IIssuer.TransactOpts, burnForAddress, from)
}

// BurnSynthsToTargetOnBehalf is a paid mutator transaction binding the contract method 0x2b3f41aa.
//
// Solidity: function burnSynthsToTargetOnBehalf(address burnForAddress, address from) returns()
func (_IIssuer *IIssuerTransactorSession) BurnSynthsToTargetOnBehalf(burnForAddress common.Address, from common.Address) (*types.Transaction, error) {
	return _IIssuer.Contract.BurnSynthsToTargetOnBehalf(&_IIssuer.TransactOpts, burnForAddress, from)
}

// IssueMaxSynths is a paid mutator transaction binding the contract method 0xc8977132.
//
// Solidity: function issueMaxSynths(address from) returns()
func (_IIssuer *IIssuerTransactor) IssueMaxSynths(opts *bind.TransactOpts, from common.Address) (*types.Transaction, error) {
	return _IIssuer.contract.Transact(opts, "issueMaxSynths", from)
}

// IssueMaxSynths is a paid mutator transaction binding the contract method 0xc8977132.
//
// Solidity: function issueMaxSynths(address from) returns()
func (_IIssuer *IIssuerSession) IssueMaxSynths(from common.Address) (*types.Transaction, error) {
	return _IIssuer.Contract.IssueMaxSynths(&_IIssuer.TransactOpts, from)
}

// IssueMaxSynths is a paid mutator transaction binding the contract method 0xc8977132.
//
// Solidity: function issueMaxSynths(address from) returns()
func (_IIssuer *IIssuerTransactorSession) IssueMaxSynths(from common.Address) (*types.Transaction, error) {
	return _IIssuer.Contract.IssueMaxSynths(&_IIssuer.TransactOpts, from)
}

// IssueMaxSynthsOnBehalf is a paid mutator transaction binding the contract method 0xfd864ccf.
//
// Solidity: function issueMaxSynthsOnBehalf(address issueFor, address from) returns()
func (_IIssuer *IIssuerTransactor) IssueMaxSynthsOnBehalf(opts *bind.TransactOpts, issueFor common.Address, from common.Address) (*types.Transaction, error) {
	return _IIssuer.contract.Transact(opts, "issueMaxSynthsOnBehalf", issueFor, from)
}

// IssueMaxSynthsOnBehalf is a paid mutator transaction binding the contract method 0xfd864ccf.
//
// Solidity: function issueMaxSynthsOnBehalf(address issueFor, address from) returns()
func (_IIssuer *IIssuerSession) IssueMaxSynthsOnBehalf(issueFor common.Address, from common.Address) (*types.Transaction, error) {
	return _IIssuer.Contract.IssueMaxSynthsOnBehalf(&_IIssuer.TransactOpts, issueFor, from)
}

// IssueMaxSynthsOnBehalf is a paid mutator transaction binding the contract method 0xfd864ccf.
//
// Solidity: function issueMaxSynthsOnBehalf(address issueFor, address from) returns()
func (_IIssuer *IIssuerTransactorSession) IssueMaxSynthsOnBehalf(issueFor common.Address, from common.Address) (*types.Transaction, error) {
	return _IIssuer.Contract.IssueMaxSynthsOnBehalf(&_IIssuer.TransactOpts, issueFor, from)
}

// IssueSynths is a paid mutator transaction binding the contract method 0x042e0688.
//
// Solidity: function issueSynths(address from, uint256 amount) returns()
func (_IIssuer *IIssuerTransactor) IssueSynths(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IIssuer.contract.Transact(opts, "issueSynths", from, amount)
}

// IssueSynths is a paid mutator transaction binding the contract method 0x042e0688.
//
// Solidity: function issueSynths(address from, uint256 amount) returns()
func (_IIssuer *IIssuerSession) IssueSynths(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IIssuer.Contract.IssueSynths(&_IIssuer.TransactOpts, from, amount)
}

// IssueSynths is a paid mutator transaction binding the contract method 0x042e0688.
//
// Solidity: function issueSynths(address from, uint256 amount) returns()
func (_IIssuer *IIssuerTransactorSession) IssueSynths(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IIssuer.Contract.IssueSynths(&_IIssuer.TransactOpts, from, amount)
}

// IssueSynthsOnBehalf is a paid mutator transaction binding the contract method 0x44ec6b62.
//
// Solidity: function issueSynthsOnBehalf(address issueFor, address from, uint256 amount) returns()
func (_IIssuer *IIssuerTransactor) IssueSynthsOnBehalf(opts *bind.TransactOpts, issueFor common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IIssuer.contract.Transact(opts, "issueSynthsOnBehalf", issueFor, from, amount)
}

// IssueSynthsOnBehalf is a paid mutator transaction binding the contract method 0x44ec6b62.
//
// Solidity: function issueSynthsOnBehalf(address issueFor, address from, uint256 amount) returns()
func (_IIssuer *IIssuerSession) IssueSynthsOnBehalf(issueFor common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IIssuer.Contract.IssueSynthsOnBehalf(&_IIssuer.TransactOpts, issueFor, from, amount)
}

// IssueSynthsOnBehalf is a paid mutator transaction binding the contract method 0x44ec6b62.
//
// Solidity: function issueSynthsOnBehalf(address issueFor, address from, uint256 amount) returns()
func (_IIssuer *IIssuerTransactorSession) IssueSynthsOnBehalf(issueFor common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IIssuer.Contract.IssueSynthsOnBehalf(&_IIssuer.TransactOpts, issueFor, from, amount)
}

// LiquidateDelinquentAccount is a paid mutator transaction binding the contract method 0xa63c4df4.
//
// Solidity: function liquidateDelinquentAccount(address account, uint256 susdAmount, address liquidator) returns(uint256 totalRedeemed, uint256 amountToLiquidate)
func (_IIssuer *IIssuerTransactor) LiquidateDelinquentAccount(opts *bind.TransactOpts, account common.Address, susdAmount *big.Int, liquidator common.Address) (*types.Transaction, error) {
	return _IIssuer.contract.Transact(opts, "liquidateDelinquentAccount", account, susdAmount, liquidator)
}

// LiquidateDelinquentAccount is a paid mutator transaction binding the contract method 0xa63c4df4.
//
// Solidity: function liquidateDelinquentAccount(address account, uint256 susdAmount, address liquidator) returns(uint256 totalRedeemed, uint256 amountToLiquidate)
func (_IIssuer *IIssuerSession) LiquidateDelinquentAccount(account common.Address, susdAmount *big.Int, liquidator common.Address) (*types.Transaction, error) {
	return _IIssuer.Contract.LiquidateDelinquentAccount(&_IIssuer.TransactOpts, account, susdAmount, liquidator)
}

// LiquidateDelinquentAccount is a paid mutator transaction binding the contract method 0xa63c4df4.
//
// Solidity: function liquidateDelinquentAccount(address account, uint256 susdAmount, address liquidator) returns(uint256 totalRedeemed, uint256 amountToLiquidate)
func (_IIssuer *IIssuerTransactorSession) LiquidateDelinquentAccount(account common.Address, susdAmount *big.Int, liquidator common.Address) (*types.Transaction, error) {
	return _IIssuer.Contract.LiquidateDelinquentAccount(&_IIssuer.TransactOpts, account, susdAmount, liquidator)
}

// IRewardsDistributionABI is the input ABI used to generate the binding from.
const IRewardsDistributionABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"distributeRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"distributions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"distributionsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IRewardsDistributionFuncSigs maps the 4-byte function signature to its string representation.
var IRewardsDistributionFuncSigs = map[string]string{
	"bf7e214f": "authority()",
	"59974e38": "distributeRewards(uint256)",
	"4487d3df": "distributions(uint256)",
	"060ca250": "distributionsLength()",
}

// IRewardsDistribution is an auto generated Go binding around an Ethereum contract.
type IRewardsDistribution struct {
	IRewardsDistributionCaller     // Read-only binding to the contract
	IRewardsDistributionTransactor // Write-only binding to the contract
	IRewardsDistributionFilterer   // Log filterer for contract events
}

// IRewardsDistributionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRewardsDistributionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardsDistributionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRewardsDistributionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardsDistributionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRewardsDistributionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardsDistributionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRewardsDistributionSession struct {
	Contract     *IRewardsDistribution // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IRewardsDistributionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRewardsDistributionCallerSession struct {
	Contract *IRewardsDistributionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IRewardsDistributionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRewardsDistributionTransactorSession struct {
	Contract     *IRewardsDistributionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IRewardsDistributionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRewardsDistributionRaw struct {
	Contract *IRewardsDistribution // Generic contract binding to access the raw methods on
}

// IRewardsDistributionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRewardsDistributionCallerRaw struct {
	Contract *IRewardsDistributionCaller // Generic read-only contract binding to access the raw methods on
}

// IRewardsDistributionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRewardsDistributionTransactorRaw struct {
	Contract *IRewardsDistributionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRewardsDistribution creates a new instance of IRewardsDistribution, bound to a specific deployed contract.
func NewIRewardsDistribution(address common.Address, backend bind.ContractBackend) (*IRewardsDistribution, error) {
	contract, err := bindIRewardsDistribution(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRewardsDistribution{IRewardsDistributionCaller: IRewardsDistributionCaller{contract: contract}, IRewardsDistributionTransactor: IRewardsDistributionTransactor{contract: contract}, IRewardsDistributionFilterer: IRewardsDistributionFilterer{contract: contract}}, nil
}

// NewIRewardsDistributionCaller creates a new read-only instance of IRewardsDistribution, bound to a specific deployed contract.
func NewIRewardsDistributionCaller(address common.Address, caller bind.ContractCaller) (*IRewardsDistributionCaller, error) {
	contract, err := bindIRewardsDistribution(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardsDistributionCaller{contract: contract}, nil
}

// NewIRewardsDistributionTransactor creates a new write-only instance of IRewardsDistribution, bound to a specific deployed contract.
func NewIRewardsDistributionTransactor(address common.Address, transactor bind.ContractTransactor) (*IRewardsDistributionTransactor, error) {
	contract, err := bindIRewardsDistribution(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardsDistributionTransactor{contract: contract}, nil
}

// NewIRewardsDistributionFilterer creates a new log filterer instance of IRewardsDistribution, bound to a specific deployed contract.
func NewIRewardsDistributionFilterer(address common.Address, filterer bind.ContractFilterer) (*IRewardsDistributionFilterer, error) {
	contract, err := bindIRewardsDistribution(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRewardsDistributionFilterer{contract: contract}, nil
}

// bindIRewardsDistribution binds a generic wrapper to an already deployed contract.
func bindIRewardsDistribution(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRewardsDistributionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardsDistribution *IRewardsDistributionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewardsDistribution.Contract.IRewardsDistributionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardsDistribution *IRewardsDistributionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardsDistribution.Contract.IRewardsDistributionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardsDistribution *IRewardsDistributionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewardsDistribution.Contract.IRewardsDistributionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardsDistribution *IRewardsDistributionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewardsDistribution.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardsDistribution *IRewardsDistributionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardsDistribution.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardsDistribution *IRewardsDistributionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewardsDistribution.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_IRewardsDistribution *IRewardsDistributionCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IRewardsDistribution.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_IRewardsDistribution *IRewardsDistributionSession) Authority() (common.Address, error) {
	return _IRewardsDistribution.Contract.Authority(&_IRewardsDistribution.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_IRewardsDistribution *IRewardsDistributionCallerSession) Authority() (common.Address, error) {
	return _IRewardsDistribution.Contract.Authority(&_IRewardsDistribution.CallOpts)
}

// Distributions is a free data retrieval call binding the contract method 0x4487d3df.
//
// Solidity: function distributions(uint256 index) view returns(address destination, uint256 amount)
func (_IRewardsDistribution *IRewardsDistributionCaller) Distributions(opts *bind.CallOpts, index *big.Int) (struct {
	Destination common.Address
	Amount      *big.Int
}, error) {
	var out []interface{}
	err := _IRewardsDistribution.contract.Call(opts, &out, "distributions", index)

	outstruct := new(struct {
		Destination common.Address
		Amount      *big.Int
	})

	outstruct.Destination = out[0].(common.Address)
	outstruct.Amount = out[1].(*big.Int)

	return *outstruct, err

}

// Distributions is a free data retrieval call binding the contract method 0x4487d3df.
//
// Solidity: function distributions(uint256 index) view returns(address destination, uint256 amount)
func (_IRewardsDistribution *IRewardsDistributionSession) Distributions(index *big.Int) (struct {
	Destination common.Address
	Amount      *big.Int
}, error) {
	return _IRewardsDistribution.Contract.Distributions(&_IRewardsDistribution.CallOpts, index)
}

// Distributions is a free data retrieval call binding the contract method 0x4487d3df.
//
// Solidity: function distributions(uint256 index) view returns(address destination, uint256 amount)
func (_IRewardsDistribution *IRewardsDistributionCallerSession) Distributions(index *big.Int) (struct {
	Destination common.Address
	Amount      *big.Int
}, error) {
	return _IRewardsDistribution.Contract.Distributions(&_IRewardsDistribution.CallOpts, index)
}

// DistributionsLength is a free data retrieval call binding the contract method 0x060ca250.
//
// Solidity: function distributionsLength() view returns(uint256)
func (_IRewardsDistribution *IRewardsDistributionCaller) DistributionsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IRewardsDistribution.contract.Call(opts, &out, "distributionsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DistributionsLength is a free data retrieval call binding the contract method 0x060ca250.
//
// Solidity: function distributionsLength() view returns(uint256)
func (_IRewardsDistribution *IRewardsDistributionSession) DistributionsLength() (*big.Int, error) {
	return _IRewardsDistribution.Contract.DistributionsLength(&_IRewardsDistribution.CallOpts)
}

// DistributionsLength is a free data retrieval call binding the contract method 0x060ca250.
//
// Solidity: function distributionsLength() view returns(uint256)
func (_IRewardsDistribution *IRewardsDistributionCallerSession) DistributionsLength() (*big.Int, error) {
	return _IRewardsDistribution.Contract.DistributionsLength(&_IRewardsDistribution.CallOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x59974e38.
//
// Solidity: function distributeRewards(uint256 amount) returns(bool)
func (_IRewardsDistribution *IRewardsDistributionTransactor) DistributeRewards(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IRewardsDistribution.contract.Transact(opts, "distributeRewards", amount)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x59974e38.
//
// Solidity: function distributeRewards(uint256 amount) returns(bool)
func (_IRewardsDistribution *IRewardsDistributionSession) DistributeRewards(amount *big.Int) (*types.Transaction, error) {
	return _IRewardsDistribution.Contract.DistributeRewards(&_IRewardsDistribution.TransactOpts, amount)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x59974e38.
//
// Solidity: function distributeRewards(uint256 amount) returns(bool)
func (_IRewardsDistribution *IRewardsDistributionTransactorSession) DistributeRewards(amount *big.Int) (*types.Transaction, error) {
	return _IRewardsDistribution.Contract.DistributeRewards(&_IRewardsDistribution.TransactOpts, amount)
}

// ISupplyScheduleABI is the input ABI used to generate the binding from.
const ISupplyScheduleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintableSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"supplyMinted\",\"type\":\"uint256\"}],\"name\":\"recordMintEvent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ISupplyScheduleFuncSigs maps the 4-byte function signature to its string representation.
var ISupplyScheduleFuncSigs = map[string]string{
	"46b45af7": "isMintable()",
	"cc5c095c": "mintableSupply()",
	"7e7961d7": "recordMintEvent(uint256)",
}

// ISupplySchedule is an auto generated Go binding around an Ethereum contract.
type ISupplySchedule struct {
	ISupplyScheduleCaller     // Read-only binding to the contract
	ISupplyScheduleTransactor // Write-only binding to the contract
	ISupplyScheduleFilterer   // Log filterer for contract events
}

// ISupplyScheduleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISupplyScheduleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISupplyScheduleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISupplyScheduleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISupplyScheduleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISupplyScheduleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISupplyScheduleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISupplyScheduleSession struct {
	Contract     *ISupplySchedule  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISupplyScheduleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISupplyScheduleCallerSession struct {
	Contract *ISupplyScheduleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ISupplyScheduleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISupplyScheduleTransactorSession struct {
	Contract     *ISupplyScheduleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ISupplyScheduleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISupplyScheduleRaw struct {
	Contract *ISupplySchedule // Generic contract binding to access the raw methods on
}

// ISupplyScheduleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISupplyScheduleCallerRaw struct {
	Contract *ISupplyScheduleCaller // Generic read-only contract binding to access the raw methods on
}

// ISupplyScheduleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISupplyScheduleTransactorRaw struct {
	Contract *ISupplyScheduleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISupplySchedule creates a new instance of ISupplySchedule, bound to a specific deployed contract.
func NewISupplySchedule(address common.Address, backend bind.ContractBackend) (*ISupplySchedule, error) {
	contract, err := bindISupplySchedule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISupplySchedule{ISupplyScheduleCaller: ISupplyScheduleCaller{contract: contract}, ISupplyScheduleTransactor: ISupplyScheduleTransactor{contract: contract}, ISupplyScheduleFilterer: ISupplyScheduleFilterer{contract: contract}}, nil
}

// NewISupplyScheduleCaller creates a new read-only instance of ISupplySchedule, bound to a specific deployed contract.
func NewISupplyScheduleCaller(address common.Address, caller bind.ContractCaller) (*ISupplyScheduleCaller, error) {
	contract, err := bindISupplySchedule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISupplyScheduleCaller{contract: contract}, nil
}

// NewISupplyScheduleTransactor creates a new write-only instance of ISupplySchedule, bound to a specific deployed contract.
func NewISupplyScheduleTransactor(address common.Address, transactor bind.ContractTransactor) (*ISupplyScheduleTransactor, error) {
	contract, err := bindISupplySchedule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISupplyScheduleTransactor{contract: contract}, nil
}

// NewISupplyScheduleFilterer creates a new log filterer instance of ISupplySchedule, bound to a specific deployed contract.
func NewISupplyScheduleFilterer(address common.Address, filterer bind.ContractFilterer) (*ISupplyScheduleFilterer, error) {
	contract, err := bindISupplySchedule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISupplyScheduleFilterer{contract: contract}, nil
}

// bindISupplySchedule binds a generic wrapper to an already deployed contract.
func bindISupplySchedule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISupplyScheduleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISupplySchedule *ISupplyScheduleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISupplySchedule.Contract.ISupplyScheduleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISupplySchedule *ISupplyScheduleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISupplySchedule.Contract.ISupplyScheduleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISupplySchedule *ISupplyScheduleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISupplySchedule.Contract.ISupplyScheduleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISupplySchedule *ISupplyScheduleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISupplySchedule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISupplySchedule *ISupplyScheduleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISupplySchedule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISupplySchedule *ISupplyScheduleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISupplySchedule.Contract.contract.Transact(opts, method, params...)
}

// IsMintable is a free data retrieval call binding the contract method 0x46b45af7.
//
// Solidity: function isMintable() view returns(bool)
func (_ISupplySchedule *ISupplyScheduleCaller) IsMintable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ISupplySchedule.contract.Call(opts, &out, "isMintable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMintable is a free data retrieval call binding the contract method 0x46b45af7.
//
// Solidity: function isMintable() view returns(bool)
func (_ISupplySchedule *ISupplyScheduleSession) IsMintable() (bool, error) {
	return _ISupplySchedule.Contract.IsMintable(&_ISupplySchedule.CallOpts)
}

// IsMintable is a free data retrieval call binding the contract method 0x46b45af7.
//
// Solidity: function isMintable() view returns(bool)
func (_ISupplySchedule *ISupplyScheduleCallerSession) IsMintable() (bool, error) {
	return _ISupplySchedule.Contract.IsMintable(&_ISupplySchedule.CallOpts)
}

// MintableSupply is a free data retrieval call binding the contract method 0xcc5c095c.
//
// Solidity: function mintableSupply() view returns(uint256)
func (_ISupplySchedule *ISupplyScheduleCaller) MintableSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISupplySchedule.contract.Call(opts, &out, "mintableSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintableSupply is a free data retrieval call binding the contract method 0xcc5c095c.
//
// Solidity: function mintableSupply() view returns(uint256)
func (_ISupplySchedule *ISupplyScheduleSession) MintableSupply() (*big.Int, error) {
	return _ISupplySchedule.Contract.MintableSupply(&_ISupplySchedule.CallOpts)
}

// MintableSupply is a free data retrieval call binding the contract method 0xcc5c095c.
//
// Solidity: function mintableSupply() view returns(uint256)
func (_ISupplySchedule *ISupplyScheduleCallerSession) MintableSupply() (*big.Int, error) {
	return _ISupplySchedule.Contract.MintableSupply(&_ISupplySchedule.CallOpts)
}

// RecordMintEvent is a paid mutator transaction binding the contract method 0x7e7961d7.
//
// Solidity: function recordMintEvent(uint256 supplyMinted) returns(bool)
func (_ISupplySchedule *ISupplyScheduleTransactor) RecordMintEvent(opts *bind.TransactOpts, supplyMinted *big.Int) (*types.Transaction, error) {
	return _ISupplySchedule.contract.Transact(opts, "recordMintEvent", supplyMinted)
}

// RecordMintEvent is a paid mutator transaction binding the contract method 0x7e7961d7.
//
// Solidity: function recordMintEvent(uint256 supplyMinted) returns(bool)
func (_ISupplySchedule *ISupplyScheduleSession) RecordMintEvent(supplyMinted *big.Int) (*types.Transaction, error) {
	return _ISupplySchedule.Contract.RecordMintEvent(&_ISupplySchedule.TransactOpts, supplyMinted)
}

// RecordMintEvent is a paid mutator transaction binding the contract method 0x7e7961d7.
//
// Solidity: function recordMintEvent(uint256 supplyMinted) returns(bool)
func (_ISupplySchedule *ISupplyScheduleTransactorSession) RecordMintEvent(supplyMinted *big.Int) (*types.Transaction, error) {
	return _ISupplySchedule.Contract.RecordMintEvent(&_ISupplySchedule.TransactOpts, supplyMinted)
}

// ISynthABI is the input ABI used to generate the binding from.
const ISynthABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currencyKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferAndSettle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFromAndSettle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"transferableSynths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ISynthFuncSigs maps the 4-byte function signature to its string representation.
var ISynthFuncSigs = map[string]string{
	"9dc29fac": "burn(address,uint256)",
	"dbd06c85": "currencyKey()",
	"867904b4": "issue(address,uint256)",
	"b014c3a3": "transferAndSettle(address,uint256)",
	"e73cced3": "transferFromAndSettle(address,address,uint256)",
	"ffff51d6": "transferableSynths(address)",
}

// ISynth is an auto generated Go binding around an Ethereum contract.
type ISynth struct {
	ISynthCaller     // Read-only binding to the contract
	ISynthTransactor // Write-only binding to the contract
	ISynthFilterer   // Log filterer for contract events
}

// ISynthCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISynthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynthSession struct {
	Contract     *ISynth           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISynthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynthCallerSession struct {
	Contract *ISynthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ISynthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynthTransactorSession struct {
	Contract     *ISynthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISynthRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISynthRaw struct {
	Contract *ISynth // Generic contract binding to access the raw methods on
}

// ISynthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynthCallerRaw struct {
	Contract *ISynthCaller // Generic read-only contract binding to access the raw methods on
}

// ISynthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynthTransactorRaw struct {
	Contract *ISynthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISynth creates a new instance of ISynth, bound to a specific deployed contract.
func NewISynth(address common.Address, backend bind.ContractBackend) (*ISynth, error) {
	contract, err := bindISynth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynth{ISynthCaller: ISynthCaller{contract: contract}, ISynthTransactor: ISynthTransactor{contract: contract}, ISynthFilterer: ISynthFilterer{contract: contract}}, nil
}

// NewISynthCaller creates a new read-only instance of ISynth, bound to a specific deployed contract.
func NewISynthCaller(address common.Address, caller bind.ContractCaller) (*ISynthCaller, error) {
	contract, err := bindISynth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynthCaller{contract: contract}, nil
}

// NewISynthTransactor creates a new write-only instance of ISynth, bound to a specific deployed contract.
func NewISynthTransactor(address common.Address, transactor bind.ContractTransactor) (*ISynthTransactor, error) {
	contract, err := bindISynth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynthTransactor{contract: contract}, nil
}

// NewISynthFilterer creates a new log filterer instance of ISynth, bound to a specific deployed contract.
func NewISynthFilterer(address common.Address, filterer bind.ContractFilterer) (*ISynthFilterer, error) {
	contract, err := bindISynth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynthFilterer{contract: contract}, nil
}

// bindISynth binds a generic wrapper to an already deployed contract.
func bindISynth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISynthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynth *ISynthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynth.Contract.ISynthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynth *ISynthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynth.Contract.ISynthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynth *ISynthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynth.Contract.ISynthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynth *ISynthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynth *ISynthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynth *ISynthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynth.Contract.contract.Transact(opts, method, params...)
}

// CurrencyKey is a free data retrieval call binding the contract method 0xdbd06c85.
//
// Solidity: function currencyKey() view returns(bytes32)
func (_ISynth *ISynthCaller) CurrencyKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ISynth.contract.Call(opts, &out, "currencyKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrencyKey is a free data retrieval call binding the contract method 0xdbd06c85.
//
// Solidity: function currencyKey() view returns(bytes32)
func (_ISynth *ISynthSession) CurrencyKey() ([32]byte, error) {
	return _ISynth.Contract.CurrencyKey(&_ISynth.CallOpts)
}

// CurrencyKey is a free data retrieval call binding the contract method 0xdbd06c85.
//
// Solidity: function currencyKey() view returns(bytes32)
func (_ISynth *ISynthCallerSession) CurrencyKey() ([32]byte, error) {
	return _ISynth.Contract.CurrencyKey(&_ISynth.CallOpts)
}

// TransferableSynths is a free data retrieval call binding the contract method 0xffff51d6.
//
// Solidity: function transferableSynths(address account) view returns(uint256)
func (_ISynth *ISynthCaller) TransferableSynths(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ISynth.contract.Call(opts, &out, "transferableSynths", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferableSynths is a free data retrieval call binding the contract method 0xffff51d6.
//
// Solidity: function transferableSynths(address account) view returns(uint256)
func (_ISynth *ISynthSession) TransferableSynths(account common.Address) (*big.Int, error) {
	return _ISynth.Contract.TransferableSynths(&_ISynth.CallOpts, account)
}

// TransferableSynths is a free data retrieval call binding the contract method 0xffff51d6.
//
// Solidity: function transferableSynths(address account) view returns(uint256)
func (_ISynth *ISynthCallerSession) TransferableSynths(account common.Address) (*big.Int, error) {
	return _ISynth.Contract.TransferableSynths(&_ISynth.CallOpts, account)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_ISynth *ISynthTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynth.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_ISynth *ISynthSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynth.Contract.Burn(&_ISynth.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_ISynth *ISynthTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynth.Contract.Burn(&_ISynth.TransactOpts, account, amount)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address account, uint256 amount) returns()
func (_ISynth *ISynthTransactor) Issue(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynth.contract.Transact(opts, "issue", account, amount)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address account, uint256 amount) returns()
func (_ISynth *ISynthSession) Issue(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynth.Contract.Issue(&_ISynth.TransactOpts, account, amount)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address account, uint256 amount) returns()
func (_ISynth *ISynthTransactorSession) Issue(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynth.Contract.Issue(&_ISynth.TransactOpts, account, amount)
}

// TransferAndSettle is a paid mutator transaction binding the contract method 0xb014c3a3.
//
// Solidity: function transferAndSettle(address to, uint256 value) returns(bool)
func (_ISynth *ISynthTransactor) TransferAndSettle(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ISynth.contract.Transact(opts, "transferAndSettle", to, value)
}

// TransferAndSettle is a paid mutator transaction binding the contract method 0xb014c3a3.
//
// Solidity: function transferAndSettle(address to, uint256 value) returns(bool)
func (_ISynth *ISynthSession) TransferAndSettle(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ISynth.Contract.TransferAndSettle(&_ISynth.TransactOpts, to, value)
}

// TransferAndSettle is a paid mutator transaction binding the contract method 0xb014c3a3.
//
// Solidity: function transferAndSettle(address to, uint256 value) returns(bool)
func (_ISynth *ISynthTransactorSession) TransferAndSettle(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ISynth.Contract.TransferAndSettle(&_ISynth.TransactOpts, to, value)
}

// TransferFromAndSettle is a paid mutator transaction binding the contract method 0xe73cced3.
//
// Solidity: function transferFromAndSettle(address from, address to, uint256 value) returns(bool)
func (_ISynth *ISynthTransactor) TransferFromAndSettle(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ISynth.contract.Transact(opts, "transferFromAndSettle", from, to, value)
}

// TransferFromAndSettle is a paid mutator transaction binding the contract method 0xe73cced3.
//
// Solidity: function transferFromAndSettle(address from, address to, uint256 value) returns(bool)
func (_ISynth *ISynthSession) TransferFromAndSettle(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ISynth.Contract.TransferFromAndSettle(&_ISynth.TransactOpts, from, to, value)
}

// TransferFromAndSettle is a paid mutator transaction binding the contract method 0xe73cced3.
//
// Solidity: function transferFromAndSettle(address from, address to, uint256 value) returns(bool)
func (_ISynth *ISynthTransactorSession) TransferFromAndSettle(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ISynth.Contract.TransferFromAndSettle(&_ISynth.TransactOpts, from, to, value)
}

// ISynthetixABI is the input ABI used to generate the binding from.
const ISynthetixABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"anySynthOrSNXRateIsInvalid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"anyRateInvalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableCurrencyKeys\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableSynthCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"availableSynths\",\"outputs\":[{\"internalType\":\"contractISynth\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnSecondary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnSynths\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnForAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnSynthsOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"burnSynthsToTarget\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnForAddress\",\"type\":\"address\"}],\"name\":\"burnSynthsToTargetOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"collateralisationRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"debtBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"exchangeForAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"}],\"name\":\"exchangeOnBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"exchangeForAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"}],\"name\":\"exchangeOnBehalfWithTracking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"}],\"name\":\"exchangeWithTracking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"}],\"name\":\"exchangeWithVirtual\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"},{\"internalType\":\"contractIVirtualSynth\",\"name\":\"vSynth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"isWaitingPeriod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"issueMaxSynths\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issueForAddress\",\"type\":\"address\"}],\"name\":\"issueMaxSynthsOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issueSynths\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issueForAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issueSynthsOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"susdAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateDelinquentAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"maxIssuableSynths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxIssuable\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintSecondary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintSecondaryRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"remainingIssuableSynths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxIssuable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alreadyIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSystemDebt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"settle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reclaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refunded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEntries\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"synths\",\"outputs\":[{\"internalType\":\"contractISynth\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"synthAddress\",\"type\":\"address\"}],\"name\":\"synthsByAddress\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"totalIssuedSynths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"totalIssuedSynthsExcludeEtherCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"transferableSynthetix\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"transferable\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ISynthetixFuncSigs maps the 4-byte function signature to its string representation.
var ISynthetixFuncSigs = map[string]string{
	"4e99bda9": "anySynthOrSNXRateIsInvalid()",
	"72cb051f": "availableCurrencyKeys()",
	"dbf63340": "availableSynthCount()",
	"835e119c": "availableSynths(uint256)",
	"edef719a": "burnSecondary(address,uint256)",
	"295da87d": "burnSynths(uint256)",
	"c2bf3880": "burnSynthsOnBehalf(address,uint256)",
	"9741fb22": "burnSynthsToTarget()",
	"2c955fa7": "burnSynthsToTargetOnBehalf(address)",
	"a5fdc5de": "collateral(address)",
	"a311c7c2": "collateralisationRatio(address)",
	"d37c4d8b": "debtBalanceOf(address,bytes32)",
	"ee52a2f3": "exchange(bytes32,uint256,bytes32)",
	"c836fa0a": "exchangeOnBehalf(address,bytes32,uint256,bytes32)",
	"91e56b68": "exchangeOnBehalfWithTracking(address,bytes32,uint256,bytes32,address,bytes32)",
	"30ead760": "exchangeWithTracking(bytes32,uint256,bytes32,address,bytes32)",
	"0e30963c": "exchangeWithVirtual(bytes32,uint256,bytes32,bytes32)",
	"1fce304d": "isWaitingPeriod(bytes32)",
	"af086c7e": "issueMaxSynths()",
	"320223db": "issueMaxSynthsOnBehalf(address)",
	"8a290014": "issueSynths(uint256)",
	"e8e09b8b": "issueSynthsOnBehalf(address,uint256)",
	"e6203ed1": "liquidateDelinquentAccount(address,uint256)",
	"05b3c1c9": "maxIssuableSynths(address)",
	"1249c58b": "mint()",
	"666ed4f1": "mintSecondary(address,uint256)",
	"d8a1f76f": "mintSecondaryRewards(uint256)",
	"1137aedf": "remainingIssuableSynths(address)",
	"987757dd": "settle(bytes32)",
	"32608039": "synths(bytes32)",
	"16b2213f": "synthsByAddress(address)",
	"83d625d4": "totalIssuedSynths(bytes32)",
	"d60888e4": "totalIssuedSynthsExcludeEtherCollateral(bytes32)",
	"6ac0bf9c": "transferableSynthetix(address)",
}

// ISynthetix is an auto generated Go binding around an Ethereum contract.
type ISynthetix struct {
	ISynthetixCaller     // Read-only binding to the contract
	ISynthetixTransactor // Write-only binding to the contract
	ISynthetixFilterer   // Log filterer for contract events
}

// ISynthetixCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISynthetixCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynthetixTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynthetixTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynthetixFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynthetixFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynthetixSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynthetixSession struct {
	Contract     *ISynthetix       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISynthetixCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynthetixCallerSession struct {
	Contract *ISynthetixCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ISynthetixTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynthetixTransactorSession struct {
	Contract     *ISynthetixTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ISynthetixRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISynthetixRaw struct {
	Contract *ISynthetix // Generic contract binding to access the raw methods on
}

// ISynthetixCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynthetixCallerRaw struct {
	Contract *ISynthetixCaller // Generic read-only contract binding to access the raw methods on
}

// ISynthetixTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynthetixTransactorRaw struct {
	Contract *ISynthetixTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISynthetix creates a new instance of ISynthetix, bound to a specific deployed contract.
func NewISynthetix(address common.Address, backend bind.ContractBackend) (*ISynthetix, error) {
	contract, err := bindISynthetix(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynthetix{ISynthetixCaller: ISynthetixCaller{contract: contract}, ISynthetixTransactor: ISynthetixTransactor{contract: contract}, ISynthetixFilterer: ISynthetixFilterer{contract: contract}}, nil
}

// NewISynthetixCaller creates a new read-only instance of ISynthetix, bound to a specific deployed contract.
func NewISynthetixCaller(address common.Address, caller bind.ContractCaller) (*ISynthetixCaller, error) {
	contract, err := bindISynthetix(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynthetixCaller{contract: contract}, nil
}

// NewISynthetixTransactor creates a new write-only instance of ISynthetix, bound to a specific deployed contract.
func NewISynthetixTransactor(address common.Address, transactor bind.ContractTransactor) (*ISynthetixTransactor, error) {
	contract, err := bindISynthetix(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynthetixTransactor{contract: contract}, nil
}

// NewISynthetixFilterer creates a new log filterer instance of ISynthetix, bound to a specific deployed contract.
func NewISynthetixFilterer(address common.Address, filterer bind.ContractFilterer) (*ISynthetixFilterer, error) {
	contract, err := bindISynthetix(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynthetixFilterer{contract: contract}, nil
}

// bindISynthetix binds a generic wrapper to an already deployed contract.
func bindISynthetix(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISynthetixABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynthetix *ISynthetixRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynthetix.Contract.ISynthetixCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynthetix *ISynthetixRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynthetix.Contract.ISynthetixTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynthetix *ISynthetixRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynthetix.Contract.ISynthetixTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynthetix *ISynthetixCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynthetix.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynthetix *ISynthetixTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynthetix.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynthetix *ISynthetixTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynthetix.Contract.contract.Transact(opts, method, params...)
}

// AnySynthOrSNXRateIsInvalid is a free data retrieval call binding the contract method 0x4e99bda9.
//
// Solidity: function anySynthOrSNXRateIsInvalid() view returns(bool anyRateInvalid)
func (_ISynthetix *ISynthetixCaller) AnySynthOrSNXRateIsInvalid(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ISynthetix.contract.Call(opts, &out, "anySynthOrSNXRateIsInvalid")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AnySynthOrSNXRateIsInvalid is a free data retrieval call binding the contract method 0x4e99bda9.
//
// Solidity: function anySynthOrSNXRateIsInvalid() view returns(bool anyRateInvalid)
func (_ISynthetix *ISynthetixSession) AnySynthOrSNXRateIsInvalid() (bool, error) {
	return _ISynthetix.Contract.AnySynthOrSNXRateIsInvalid(&_ISynthetix.CallOpts)
}

// AnySynthOrSNXRateIsInvalid is a free data retrieval call binding the contract method 0x4e99bda9.
//
// Solidity: function anySynthOrSNXRateIsInvalid() view returns(bool anyRateInvalid)
func (_ISynthetix *ISynthetixCallerSession) AnySynthOrSNXRateIsInvalid() (bool, error) {
	return _ISynthetix.Contract.AnySynthOrSNXRateIsInvalid(&_ISynthetix.CallOpts)
}

// AvailableCurrencyKeys is a free data retrieval call binding the contract method 0x72cb051f.
//
// Solidity: function availableCurrencyKeys() view returns(bytes32[])
func (_ISynthetix *ISynthetixCaller) AvailableCurrencyKeys(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _ISynthetix.contract.Call(opts, &out, "availableCurrencyKeys")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// AvailableCurrencyKeys is a free data retrieval call binding the contract method 0x72cb051f.
//
// Solidity: function availableCurrencyKeys() view returns(bytes32[])
func (_ISynthetix *ISynthetixSession) AvailableCurrencyKeys() ([][32]byte, error) {
	return _ISynthetix.Contract.AvailableCurrencyKeys(&_ISynthetix.CallOpts)
}

// AvailableCurrencyKeys is a free data retrieval call binding the contract method 0x72cb051f.
//
// Solidity: function availableCurrencyKeys() view returns(bytes32[])
func (_ISynthetix *ISynthetixCallerSession) AvailableCurrencyKeys() ([][32]byte, error) {
	return _ISynthetix.Contract.AvailableCurrencyKeys(&_ISynthetix.CallOpts)
}

// AvailableSynthCount is a free data retrieval call binding the contract method 0xdbf63340.
//
// Solidity: function availableSynthCount() view returns(uint256)
func (_ISynthetix *ISynthetixCaller) AvailableSynthCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISynthetix.contract.Call(opts, &out, "availableSynthCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableSynthCount is a free data retrieval call binding the contract method 0xdbf63340.
//
// Solidity: function availableSynthCount() view returns(uint256)
func (_ISynthetix *ISynthetixSession) AvailableSynthCount() (*big.Int, error) {
	return _ISynthetix.Contract.AvailableSynthCount(&_ISynthetix.CallOpts)
}

// AvailableSynthCount is a free data retrieval call binding the contract method 0xdbf63340.
//
// Solidity: function availableSynthCount() view returns(uint256)
func (_ISynthetix *ISynthetixCallerSession) AvailableSynthCount() (*big.Int, error) {
	return _ISynthetix.Contract.AvailableSynthCount(&_ISynthetix.CallOpts)
}

// AvailableSynths is a free data retrieval call binding the contract method 0x835e119c.
//
// Solidity: function availableSynths(uint256 index) view returns(address)
func (_ISynthetix *ISynthetixCaller) AvailableSynths(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ISynthetix.contract.Call(opts, &out, "availableSynths", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvailableSynths is a free data retrieval call binding the contract method 0x835e119c.
//
// Solidity: function availableSynths(uint256 index) view returns(address)
func (_ISynthetix *ISynthetixSession) AvailableSynths(index *big.Int) (common.Address, error) {
	return _ISynthetix.Contract.AvailableSynths(&_ISynthetix.CallOpts, index)
}

// AvailableSynths is a free data retrieval call binding the contract method 0x835e119c.
//
// Solidity: function availableSynths(uint256 index) view returns(address)
func (_ISynthetix *ISynthetixCallerSession) AvailableSynths(index *big.Int) (common.Address, error) {
	return _ISynthetix.Contract.AvailableSynths(&_ISynthetix.CallOpts, index)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address account) view returns(uint256)
func (_ISynthetix *ISynthetixCaller) Collateral(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ISynthetix.contract.Call(opts, &out, "collateral", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address account) view returns(uint256)
func (_ISynthetix *ISynthetixSession) Collateral(account common.Address) (*big.Int, error) {
	return _ISynthetix.Contract.Collateral(&_ISynthetix.CallOpts, account)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address account) view returns(uint256)
func (_ISynthetix *ISynthetixCallerSession) Collateral(account common.Address) (*big.Int, error) {
	return _ISynthetix.Contract.Collateral(&_ISynthetix.CallOpts, account)
}

// CollateralisationRatio is a free data retrieval call binding the contract method 0xa311c7c2.
//
// Solidity: function collateralisationRatio(address issuer) view returns(uint256)
func (_ISynthetix *ISynthetixCaller) CollateralisationRatio(opts *bind.CallOpts, issuer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ISynthetix.contract.Call(opts, &out, "collateralisationRatio", issuer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralisationRatio is a free data retrieval call binding the contract method 0xa311c7c2.
//
// Solidity: function collateralisationRatio(address issuer) view returns(uint256)
func (_ISynthetix *ISynthetixSession) CollateralisationRatio(issuer common.Address) (*big.Int, error) {
	return _ISynthetix.Contract.CollateralisationRatio(&_ISynthetix.CallOpts, issuer)
}

// CollateralisationRatio is a free data retrieval call binding the contract method 0xa311c7c2.
//
// Solidity: function collateralisationRatio(address issuer) view returns(uint256)
func (_ISynthetix *ISynthetixCallerSession) CollateralisationRatio(issuer common.Address) (*big.Int, error) {
	return _ISynthetix.Contract.CollateralisationRatio(&_ISynthetix.CallOpts, issuer)
}

// DebtBalanceOf is a free data retrieval call binding the contract method 0xd37c4d8b.
//
// Solidity: function debtBalanceOf(address issuer, bytes32 currencyKey) view returns(uint256)
func (_ISynthetix *ISynthetixCaller) DebtBalanceOf(opts *bind.CallOpts, issuer common.Address, currencyKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ISynthetix.contract.Call(opts, &out, "debtBalanceOf", issuer, currencyKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtBalanceOf is a free data retrieval call binding the contract method 0xd37c4d8b.
//
// Solidity: function debtBalanceOf(address issuer, bytes32 currencyKey) view returns(uint256)
func (_ISynthetix *ISynthetixSession) DebtBalanceOf(issuer common.Address, currencyKey [32]byte) (*big.Int, error) {
	return _ISynthetix.Contract.DebtBalanceOf(&_ISynthetix.CallOpts, issuer, currencyKey)
}

// DebtBalanceOf is a free data retrieval call binding the contract method 0xd37c4d8b.
//
// Solidity: function debtBalanceOf(address issuer, bytes32 currencyKey) view returns(uint256)
func (_ISynthetix *ISynthetixCallerSession) DebtBalanceOf(issuer common.Address, currencyKey [32]byte) (*big.Int, error) {
	return _ISynthetix.Contract.DebtBalanceOf(&_ISynthetix.CallOpts, issuer, currencyKey)
}

// IsWaitingPeriod is a free data retrieval call binding the contract method 0x1fce304d.
//
// Solidity: function isWaitingPeriod(bytes32 currencyKey) view returns(bool)
func (_ISynthetix *ISynthetixCaller) IsWaitingPeriod(opts *bind.CallOpts, currencyKey [32]byte) (bool, error) {
	var out []interface{}
	err := _ISynthetix.contract.Call(opts, &out, "isWaitingPeriod", currencyKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWaitingPeriod is a free data retrieval call binding the contract method 0x1fce304d.
//
// Solidity: function isWaitingPeriod(bytes32 currencyKey) view returns(bool)
func (_ISynthetix *ISynthetixSession) IsWaitingPeriod(currencyKey [32]byte) (bool, error) {
	return _ISynthetix.Contract.IsWaitingPeriod(&_ISynthetix.CallOpts, currencyKey)
}

// IsWaitingPeriod is a free data retrieval call binding the contract method 0x1fce304d.
//
// Solidity: function isWaitingPeriod(bytes32 currencyKey) view returns(bool)
func (_ISynthetix *ISynthetixCallerSession) IsWaitingPeriod(currencyKey [32]byte) (bool, error) {
	return _ISynthetix.Contract.IsWaitingPeriod(&_ISynthetix.CallOpts, currencyKey)
}

// MaxIssuableSynths is a free data retrieval call binding the contract method 0x05b3c1c9.
//
// Solidity: function maxIssuableSynths(address issuer) view returns(uint256 maxIssuable)
func (_ISynthetix *ISynthetixCaller) MaxIssuableSynths(opts *bind.CallOpts, issuer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ISynthetix.contract.Call(opts, &out, "maxIssuableSynths", issuer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxIssuableSynths is a free data retrieval call binding the contract method 0x05b3c1c9.
//
// Solidity: function maxIssuableSynths(address issuer) view returns(uint256 maxIssuable)
func (_ISynthetix *ISynthetixSession) MaxIssuableSynths(issuer common.Address) (*big.Int, error) {
	return _ISynthetix.Contract.MaxIssuableSynths(&_ISynthetix.CallOpts, issuer)
}

// MaxIssuableSynths is a free data retrieval call binding the contract method 0x05b3c1c9.
//
// Solidity: function maxIssuableSynths(address issuer) view returns(uint256 maxIssuable)
func (_ISynthetix *ISynthetixCallerSession) MaxIssuableSynths(issuer common.Address) (*big.Int, error) {
	return _ISynthetix.Contract.MaxIssuableSynths(&_ISynthetix.CallOpts, issuer)
}

// RemainingIssuableSynths is a free data retrieval call binding the contract method 0x1137aedf.
//
// Solidity: function remainingIssuableSynths(address issuer) view returns(uint256 maxIssuable, uint256 alreadyIssued, uint256 totalSystemDebt)
func (_ISynthetix *ISynthetixCaller) RemainingIssuableSynths(opts *bind.CallOpts, issuer common.Address) (struct {
	MaxIssuable     *big.Int
	AlreadyIssued   *big.Int
	TotalSystemDebt *big.Int
}, error) {
	var out []interface{}
	err := _ISynthetix.contract.Call(opts, &out, "remainingIssuableSynths", issuer)

	outstruct := new(struct {
		MaxIssuable     *big.Int
		AlreadyIssued   *big.Int
		TotalSystemDebt *big.Int
	})

	outstruct.MaxIssuable = out[0].(*big.Int)
	outstruct.AlreadyIssued = out[1].(*big.Int)
	outstruct.TotalSystemDebt = out[2].(*big.Int)

	return *outstruct, err

}

// RemainingIssuableSynths is a free data retrieval call binding the contract method 0x1137aedf.
//
// Solidity: function remainingIssuableSynths(address issuer) view returns(uint256 maxIssuable, uint256 alreadyIssued, uint256 totalSystemDebt)
func (_ISynthetix *ISynthetixSession) RemainingIssuableSynths(issuer common.Address) (struct {
	MaxIssuable     *big.Int
	AlreadyIssued   *big.Int
	TotalSystemDebt *big.Int
}, error) {
	return _ISynthetix.Contract.RemainingIssuableSynths(&_ISynthetix.CallOpts, issuer)
}

// RemainingIssuableSynths is a free data retrieval call binding the contract method 0x1137aedf.
//
// Solidity: function remainingIssuableSynths(address issuer) view returns(uint256 maxIssuable, uint256 alreadyIssued, uint256 totalSystemDebt)
func (_ISynthetix *ISynthetixCallerSession) RemainingIssuableSynths(issuer common.Address) (struct {
	MaxIssuable     *big.Int
	AlreadyIssued   *big.Int
	TotalSystemDebt *big.Int
}, error) {
	return _ISynthetix.Contract.RemainingIssuableSynths(&_ISynthetix.CallOpts, issuer)
}

// Synths is a free data retrieval call binding the contract method 0x32608039.
//
// Solidity: function synths(bytes32 currencyKey) view returns(address)
func (_ISynthetix *ISynthetixCaller) Synths(opts *bind.CallOpts, currencyKey [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ISynthetix.contract.Call(opts, &out, "synths", currencyKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Synths is a free data retrieval call binding the contract method 0x32608039.
//
// Solidity: function synths(bytes32 currencyKey) view returns(address)
func (_ISynthetix *ISynthetixSession) Synths(currencyKey [32]byte) (common.Address, error) {
	return _ISynthetix.Contract.Synths(&_ISynthetix.CallOpts, currencyKey)
}

// Synths is a free data retrieval call binding the contract method 0x32608039.
//
// Solidity: function synths(bytes32 currencyKey) view returns(address)
func (_ISynthetix *ISynthetixCallerSession) Synths(currencyKey [32]byte) (common.Address, error) {
	return _ISynthetix.Contract.Synths(&_ISynthetix.CallOpts, currencyKey)
}

// SynthsByAddress is a free data retrieval call binding the contract method 0x16b2213f.
//
// Solidity: function synthsByAddress(address synthAddress) view returns(bytes32)
func (_ISynthetix *ISynthetixCaller) SynthsByAddress(opts *bind.CallOpts, synthAddress common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ISynthetix.contract.Call(opts, &out, "synthsByAddress", synthAddress)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SynthsByAddress is a free data retrieval call binding the contract method 0x16b2213f.
//
// Solidity: function synthsByAddress(address synthAddress) view returns(bytes32)
func (_ISynthetix *ISynthetixSession) SynthsByAddress(synthAddress common.Address) ([32]byte, error) {
	return _ISynthetix.Contract.SynthsByAddress(&_ISynthetix.CallOpts, synthAddress)
}

// SynthsByAddress is a free data retrieval call binding the contract method 0x16b2213f.
//
// Solidity: function synthsByAddress(address synthAddress) view returns(bytes32)
func (_ISynthetix *ISynthetixCallerSession) SynthsByAddress(synthAddress common.Address) ([32]byte, error) {
	return _ISynthetix.Contract.SynthsByAddress(&_ISynthetix.CallOpts, synthAddress)
}

// TotalIssuedSynths is a free data retrieval call binding the contract method 0x83d625d4.
//
// Solidity: function totalIssuedSynths(bytes32 currencyKey) view returns(uint256)
func (_ISynthetix *ISynthetixCaller) TotalIssuedSynths(opts *bind.CallOpts, currencyKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ISynthetix.contract.Call(opts, &out, "totalIssuedSynths", currencyKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalIssuedSynths is a free data retrieval call binding the contract method 0x83d625d4.
//
// Solidity: function totalIssuedSynths(bytes32 currencyKey) view returns(uint256)
func (_ISynthetix *ISynthetixSession) TotalIssuedSynths(currencyKey [32]byte) (*big.Int, error) {
	return _ISynthetix.Contract.TotalIssuedSynths(&_ISynthetix.CallOpts, currencyKey)
}

// TotalIssuedSynths is a free data retrieval call binding the contract method 0x83d625d4.
//
// Solidity: function totalIssuedSynths(bytes32 currencyKey) view returns(uint256)
func (_ISynthetix *ISynthetixCallerSession) TotalIssuedSynths(currencyKey [32]byte) (*big.Int, error) {
	return _ISynthetix.Contract.TotalIssuedSynths(&_ISynthetix.CallOpts, currencyKey)
}

// TotalIssuedSynthsExcludeEtherCollateral is a free data retrieval call binding the contract method 0xd60888e4.
//
// Solidity: function totalIssuedSynthsExcludeEtherCollateral(bytes32 currencyKey) view returns(uint256)
func (_ISynthetix *ISynthetixCaller) TotalIssuedSynthsExcludeEtherCollateral(opts *bind.CallOpts, currencyKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ISynthetix.contract.Call(opts, &out, "totalIssuedSynthsExcludeEtherCollateral", currencyKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalIssuedSynthsExcludeEtherCollateral is a free data retrieval call binding the contract method 0xd60888e4.
//
// Solidity: function totalIssuedSynthsExcludeEtherCollateral(bytes32 currencyKey) view returns(uint256)
func (_ISynthetix *ISynthetixSession) TotalIssuedSynthsExcludeEtherCollateral(currencyKey [32]byte) (*big.Int, error) {
	return _ISynthetix.Contract.TotalIssuedSynthsExcludeEtherCollateral(&_ISynthetix.CallOpts, currencyKey)
}

// TotalIssuedSynthsExcludeEtherCollateral is a free data retrieval call binding the contract method 0xd60888e4.
//
// Solidity: function totalIssuedSynthsExcludeEtherCollateral(bytes32 currencyKey) view returns(uint256)
func (_ISynthetix *ISynthetixCallerSession) TotalIssuedSynthsExcludeEtherCollateral(currencyKey [32]byte) (*big.Int, error) {
	return _ISynthetix.Contract.TotalIssuedSynthsExcludeEtherCollateral(&_ISynthetix.CallOpts, currencyKey)
}

// TransferableSynthetix is a free data retrieval call binding the contract method 0x6ac0bf9c.
//
// Solidity: function transferableSynthetix(address account) view returns(uint256 transferable)
func (_ISynthetix *ISynthetixCaller) TransferableSynthetix(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ISynthetix.contract.Call(opts, &out, "transferableSynthetix", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferableSynthetix is a free data retrieval call binding the contract method 0x6ac0bf9c.
//
// Solidity: function transferableSynthetix(address account) view returns(uint256 transferable)
func (_ISynthetix *ISynthetixSession) TransferableSynthetix(account common.Address) (*big.Int, error) {
	return _ISynthetix.Contract.TransferableSynthetix(&_ISynthetix.CallOpts, account)
}

// TransferableSynthetix is a free data retrieval call binding the contract method 0x6ac0bf9c.
//
// Solidity: function transferableSynthetix(address account) view returns(uint256 transferable)
func (_ISynthetix *ISynthetixCallerSession) TransferableSynthetix(account common.Address) (*big.Int, error) {
	return _ISynthetix.Contract.TransferableSynthetix(&_ISynthetix.CallOpts, account)
}

// BurnSecondary is a paid mutator transaction binding the contract method 0xedef719a.
//
// Solidity: function burnSecondary(address account, uint256 amount) returns()
func (_ISynthetix *ISynthetixTransactor) BurnSecondary(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "burnSecondary", account, amount)
}

// BurnSecondary is a paid mutator transaction binding the contract method 0xedef719a.
//
// Solidity: function burnSecondary(address account, uint256 amount) returns()
func (_ISynthetix *ISynthetixSession) BurnSecondary(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.BurnSecondary(&_ISynthetix.TransactOpts, account, amount)
}

// BurnSecondary is a paid mutator transaction binding the contract method 0xedef719a.
//
// Solidity: function burnSecondary(address account, uint256 amount) returns()
func (_ISynthetix *ISynthetixTransactorSession) BurnSecondary(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.BurnSecondary(&_ISynthetix.TransactOpts, account, amount)
}

// BurnSynths is a paid mutator transaction binding the contract method 0x295da87d.
//
// Solidity: function burnSynths(uint256 amount) returns()
func (_ISynthetix *ISynthetixTransactor) BurnSynths(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "burnSynths", amount)
}

// BurnSynths is a paid mutator transaction binding the contract method 0x295da87d.
//
// Solidity: function burnSynths(uint256 amount) returns()
func (_ISynthetix *ISynthetixSession) BurnSynths(amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.BurnSynths(&_ISynthetix.TransactOpts, amount)
}

// BurnSynths is a paid mutator transaction binding the contract method 0x295da87d.
//
// Solidity: function burnSynths(uint256 amount) returns()
func (_ISynthetix *ISynthetixTransactorSession) BurnSynths(amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.BurnSynths(&_ISynthetix.TransactOpts, amount)
}

// BurnSynthsOnBehalf is a paid mutator transaction binding the contract method 0xc2bf3880.
//
// Solidity: function burnSynthsOnBehalf(address burnForAddress, uint256 amount) returns()
func (_ISynthetix *ISynthetixTransactor) BurnSynthsOnBehalf(opts *bind.TransactOpts, burnForAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "burnSynthsOnBehalf", burnForAddress, amount)
}

// BurnSynthsOnBehalf is a paid mutator transaction binding the contract method 0xc2bf3880.
//
// Solidity: function burnSynthsOnBehalf(address burnForAddress, uint256 amount) returns()
func (_ISynthetix *ISynthetixSession) BurnSynthsOnBehalf(burnForAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.BurnSynthsOnBehalf(&_ISynthetix.TransactOpts, burnForAddress, amount)
}

// BurnSynthsOnBehalf is a paid mutator transaction binding the contract method 0xc2bf3880.
//
// Solidity: function burnSynthsOnBehalf(address burnForAddress, uint256 amount) returns()
func (_ISynthetix *ISynthetixTransactorSession) BurnSynthsOnBehalf(burnForAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.BurnSynthsOnBehalf(&_ISynthetix.TransactOpts, burnForAddress, amount)
}

// BurnSynthsToTarget is a paid mutator transaction binding the contract method 0x9741fb22.
//
// Solidity: function burnSynthsToTarget() returns()
func (_ISynthetix *ISynthetixTransactor) BurnSynthsToTarget(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "burnSynthsToTarget")
}

// BurnSynthsToTarget is a paid mutator transaction binding the contract method 0x9741fb22.
//
// Solidity: function burnSynthsToTarget() returns()
func (_ISynthetix *ISynthetixSession) BurnSynthsToTarget() (*types.Transaction, error) {
	return _ISynthetix.Contract.BurnSynthsToTarget(&_ISynthetix.TransactOpts)
}

// BurnSynthsToTarget is a paid mutator transaction binding the contract method 0x9741fb22.
//
// Solidity: function burnSynthsToTarget() returns()
func (_ISynthetix *ISynthetixTransactorSession) BurnSynthsToTarget() (*types.Transaction, error) {
	return _ISynthetix.Contract.BurnSynthsToTarget(&_ISynthetix.TransactOpts)
}

// BurnSynthsToTargetOnBehalf is a paid mutator transaction binding the contract method 0x2c955fa7.
//
// Solidity: function burnSynthsToTargetOnBehalf(address burnForAddress) returns()
func (_ISynthetix *ISynthetixTransactor) BurnSynthsToTargetOnBehalf(opts *bind.TransactOpts, burnForAddress common.Address) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "burnSynthsToTargetOnBehalf", burnForAddress)
}

// BurnSynthsToTargetOnBehalf is a paid mutator transaction binding the contract method 0x2c955fa7.
//
// Solidity: function burnSynthsToTargetOnBehalf(address burnForAddress) returns()
func (_ISynthetix *ISynthetixSession) BurnSynthsToTargetOnBehalf(burnForAddress common.Address) (*types.Transaction, error) {
	return _ISynthetix.Contract.BurnSynthsToTargetOnBehalf(&_ISynthetix.TransactOpts, burnForAddress)
}

// BurnSynthsToTargetOnBehalf is a paid mutator transaction binding the contract method 0x2c955fa7.
//
// Solidity: function burnSynthsToTargetOnBehalf(address burnForAddress) returns()
func (_ISynthetix *ISynthetixTransactorSession) BurnSynthsToTargetOnBehalf(burnForAddress common.Address) (*types.Transaction, error) {
	return _ISynthetix.Contract.BurnSynthsToTargetOnBehalf(&_ISynthetix.TransactOpts, burnForAddress)
}

// Exchange is a paid mutator transaction binding the contract method 0xee52a2f3.
//
// Solidity: function exchange(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey) returns(uint256 amountReceived)
func (_ISynthetix *ISynthetixTransactor) Exchange(opts *bind.TransactOpts, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "exchange", sourceCurrencyKey, sourceAmount, destinationCurrencyKey)
}

// Exchange is a paid mutator transaction binding the contract method 0xee52a2f3.
//
// Solidity: function exchange(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey) returns(uint256 amountReceived)
func (_ISynthetix *ISynthetixSession) Exchange(sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte) (*types.Transaction, error) {
	return _ISynthetix.Contract.Exchange(&_ISynthetix.TransactOpts, sourceCurrencyKey, sourceAmount, destinationCurrencyKey)
}

// Exchange is a paid mutator transaction binding the contract method 0xee52a2f3.
//
// Solidity: function exchange(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey) returns(uint256 amountReceived)
func (_ISynthetix *ISynthetixTransactorSession) Exchange(sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte) (*types.Transaction, error) {
	return _ISynthetix.Contract.Exchange(&_ISynthetix.TransactOpts, sourceCurrencyKey, sourceAmount, destinationCurrencyKey)
}

// ExchangeOnBehalf is a paid mutator transaction binding the contract method 0xc836fa0a.
//
// Solidity: function exchangeOnBehalf(address exchangeForAddress, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey) returns(uint256 amountReceived)
func (_ISynthetix *ISynthetixTransactor) ExchangeOnBehalf(opts *bind.TransactOpts, exchangeForAddress common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "exchangeOnBehalf", exchangeForAddress, sourceCurrencyKey, sourceAmount, destinationCurrencyKey)
}

// ExchangeOnBehalf is a paid mutator transaction binding the contract method 0xc836fa0a.
//
// Solidity: function exchangeOnBehalf(address exchangeForAddress, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey) returns(uint256 amountReceived)
func (_ISynthetix *ISynthetixSession) ExchangeOnBehalf(exchangeForAddress common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte) (*types.Transaction, error) {
	return _ISynthetix.Contract.ExchangeOnBehalf(&_ISynthetix.TransactOpts, exchangeForAddress, sourceCurrencyKey, sourceAmount, destinationCurrencyKey)
}

// ExchangeOnBehalf is a paid mutator transaction binding the contract method 0xc836fa0a.
//
// Solidity: function exchangeOnBehalf(address exchangeForAddress, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey) returns(uint256 amountReceived)
func (_ISynthetix *ISynthetixTransactorSession) ExchangeOnBehalf(exchangeForAddress common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte) (*types.Transaction, error) {
	return _ISynthetix.Contract.ExchangeOnBehalf(&_ISynthetix.TransactOpts, exchangeForAddress, sourceCurrencyKey, sourceAmount, destinationCurrencyKey)
}

// ExchangeOnBehalfWithTracking is a paid mutator transaction binding the contract method 0x91e56b68.
//
// Solidity: function exchangeOnBehalfWithTracking(address exchangeForAddress, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_ISynthetix *ISynthetixTransactor) ExchangeOnBehalfWithTracking(opts *bind.TransactOpts, exchangeForAddress common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "exchangeOnBehalfWithTracking", exchangeForAddress, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, originator, trackingCode)
}

// ExchangeOnBehalfWithTracking is a paid mutator transaction binding the contract method 0x91e56b68.
//
// Solidity: function exchangeOnBehalfWithTracking(address exchangeForAddress, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_ISynthetix *ISynthetixSession) ExchangeOnBehalfWithTracking(exchangeForAddress common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _ISynthetix.Contract.ExchangeOnBehalfWithTracking(&_ISynthetix.TransactOpts, exchangeForAddress, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, originator, trackingCode)
}

// ExchangeOnBehalfWithTracking is a paid mutator transaction binding the contract method 0x91e56b68.
//
// Solidity: function exchangeOnBehalfWithTracking(address exchangeForAddress, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_ISynthetix *ISynthetixTransactorSession) ExchangeOnBehalfWithTracking(exchangeForAddress common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _ISynthetix.Contract.ExchangeOnBehalfWithTracking(&_ISynthetix.TransactOpts, exchangeForAddress, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, originator, trackingCode)
}

// ExchangeWithTracking is a paid mutator transaction binding the contract method 0x30ead760.
//
// Solidity: function exchangeWithTracking(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_ISynthetix *ISynthetixTransactor) ExchangeWithTracking(opts *bind.TransactOpts, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "exchangeWithTracking", sourceCurrencyKey, sourceAmount, destinationCurrencyKey, originator, trackingCode)
}

// ExchangeWithTracking is a paid mutator transaction binding the contract method 0x30ead760.
//
// Solidity: function exchangeWithTracking(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_ISynthetix *ISynthetixSession) ExchangeWithTracking(sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _ISynthetix.Contract.ExchangeWithTracking(&_ISynthetix.TransactOpts, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, originator, trackingCode)
}

// ExchangeWithTracking is a paid mutator transaction binding the contract method 0x30ead760.
//
// Solidity: function exchangeWithTracking(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_ISynthetix *ISynthetixTransactorSession) ExchangeWithTracking(sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _ISynthetix.Contract.ExchangeWithTracking(&_ISynthetix.TransactOpts, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, originator, trackingCode)
}

// ExchangeWithVirtual is a paid mutator transaction binding the contract method 0x0e30963c.
//
// Solidity: function exchangeWithVirtual(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, bytes32 trackingCode) returns(uint256 amountReceived, address vSynth)
func (_ISynthetix *ISynthetixTransactor) ExchangeWithVirtual(opts *bind.TransactOpts, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, trackingCode [32]byte) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "exchangeWithVirtual", sourceCurrencyKey, sourceAmount, destinationCurrencyKey, trackingCode)
}

// ExchangeWithVirtual is a paid mutator transaction binding the contract method 0x0e30963c.
//
// Solidity: function exchangeWithVirtual(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, bytes32 trackingCode) returns(uint256 amountReceived, address vSynth)
func (_ISynthetix *ISynthetixSession) ExchangeWithVirtual(sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, trackingCode [32]byte) (*types.Transaction, error) {
	return _ISynthetix.Contract.ExchangeWithVirtual(&_ISynthetix.TransactOpts, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, trackingCode)
}

// ExchangeWithVirtual is a paid mutator transaction binding the contract method 0x0e30963c.
//
// Solidity: function exchangeWithVirtual(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, bytes32 trackingCode) returns(uint256 amountReceived, address vSynth)
func (_ISynthetix *ISynthetixTransactorSession) ExchangeWithVirtual(sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, trackingCode [32]byte) (*types.Transaction, error) {
	return _ISynthetix.Contract.ExchangeWithVirtual(&_ISynthetix.TransactOpts, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, trackingCode)
}

// IssueMaxSynths is a paid mutator transaction binding the contract method 0xaf086c7e.
//
// Solidity: function issueMaxSynths() returns()
func (_ISynthetix *ISynthetixTransactor) IssueMaxSynths(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "issueMaxSynths")
}

// IssueMaxSynths is a paid mutator transaction binding the contract method 0xaf086c7e.
//
// Solidity: function issueMaxSynths() returns()
func (_ISynthetix *ISynthetixSession) IssueMaxSynths() (*types.Transaction, error) {
	return _ISynthetix.Contract.IssueMaxSynths(&_ISynthetix.TransactOpts)
}

// IssueMaxSynths is a paid mutator transaction binding the contract method 0xaf086c7e.
//
// Solidity: function issueMaxSynths() returns()
func (_ISynthetix *ISynthetixTransactorSession) IssueMaxSynths() (*types.Transaction, error) {
	return _ISynthetix.Contract.IssueMaxSynths(&_ISynthetix.TransactOpts)
}

// IssueMaxSynthsOnBehalf is a paid mutator transaction binding the contract method 0x320223db.
//
// Solidity: function issueMaxSynthsOnBehalf(address issueForAddress) returns()
func (_ISynthetix *ISynthetixTransactor) IssueMaxSynthsOnBehalf(opts *bind.TransactOpts, issueForAddress common.Address) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "issueMaxSynthsOnBehalf", issueForAddress)
}

// IssueMaxSynthsOnBehalf is a paid mutator transaction binding the contract method 0x320223db.
//
// Solidity: function issueMaxSynthsOnBehalf(address issueForAddress) returns()
func (_ISynthetix *ISynthetixSession) IssueMaxSynthsOnBehalf(issueForAddress common.Address) (*types.Transaction, error) {
	return _ISynthetix.Contract.IssueMaxSynthsOnBehalf(&_ISynthetix.TransactOpts, issueForAddress)
}

// IssueMaxSynthsOnBehalf is a paid mutator transaction binding the contract method 0x320223db.
//
// Solidity: function issueMaxSynthsOnBehalf(address issueForAddress) returns()
func (_ISynthetix *ISynthetixTransactorSession) IssueMaxSynthsOnBehalf(issueForAddress common.Address) (*types.Transaction, error) {
	return _ISynthetix.Contract.IssueMaxSynthsOnBehalf(&_ISynthetix.TransactOpts, issueForAddress)
}

// IssueSynths is a paid mutator transaction binding the contract method 0x8a290014.
//
// Solidity: function issueSynths(uint256 amount) returns()
func (_ISynthetix *ISynthetixTransactor) IssueSynths(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "issueSynths", amount)
}

// IssueSynths is a paid mutator transaction binding the contract method 0x8a290014.
//
// Solidity: function issueSynths(uint256 amount) returns()
func (_ISynthetix *ISynthetixSession) IssueSynths(amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.IssueSynths(&_ISynthetix.TransactOpts, amount)
}

// IssueSynths is a paid mutator transaction binding the contract method 0x8a290014.
//
// Solidity: function issueSynths(uint256 amount) returns()
func (_ISynthetix *ISynthetixTransactorSession) IssueSynths(amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.IssueSynths(&_ISynthetix.TransactOpts, amount)
}

// IssueSynthsOnBehalf is a paid mutator transaction binding the contract method 0xe8e09b8b.
//
// Solidity: function issueSynthsOnBehalf(address issueForAddress, uint256 amount) returns()
func (_ISynthetix *ISynthetixTransactor) IssueSynthsOnBehalf(opts *bind.TransactOpts, issueForAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "issueSynthsOnBehalf", issueForAddress, amount)
}

// IssueSynthsOnBehalf is a paid mutator transaction binding the contract method 0xe8e09b8b.
//
// Solidity: function issueSynthsOnBehalf(address issueForAddress, uint256 amount) returns()
func (_ISynthetix *ISynthetixSession) IssueSynthsOnBehalf(issueForAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.IssueSynthsOnBehalf(&_ISynthetix.TransactOpts, issueForAddress, amount)
}

// IssueSynthsOnBehalf is a paid mutator transaction binding the contract method 0xe8e09b8b.
//
// Solidity: function issueSynthsOnBehalf(address issueForAddress, uint256 amount) returns()
func (_ISynthetix *ISynthetixTransactorSession) IssueSynthsOnBehalf(issueForAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.IssueSynthsOnBehalf(&_ISynthetix.TransactOpts, issueForAddress, amount)
}

// LiquidateDelinquentAccount is a paid mutator transaction binding the contract method 0xe6203ed1.
//
// Solidity: function liquidateDelinquentAccount(address account, uint256 susdAmount) returns(bool)
func (_ISynthetix *ISynthetixTransactor) LiquidateDelinquentAccount(opts *bind.TransactOpts, account common.Address, susdAmount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "liquidateDelinquentAccount", account, susdAmount)
}

// LiquidateDelinquentAccount is a paid mutator transaction binding the contract method 0xe6203ed1.
//
// Solidity: function liquidateDelinquentAccount(address account, uint256 susdAmount) returns(bool)
func (_ISynthetix *ISynthetixSession) LiquidateDelinquentAccount(account common.Address, susdAmount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.LiquidateDelinquentAccount(&_ISynthetix.TransactOpts, account, susdAmount)
}

// LiquidateDelinquentAccount is a paid mutator transaction binding the contract method 0xe6203ed1.
//
// Solidity: function liquidateDelinquentAccount(address account, uint256 susdAmount) returns(bool)
func (_ISynthetix *ISynthetixTransactorSession) LiquidateDelinquentAccount(account common.Address, susdAmount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.LiquidateDelinquentAccount(&_ISynthetix.TransactOpts, account, susdAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(bool)
func (_ISynthetix *ISynthetixTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(bool)
func (_ISynthetix *ISynthetixSession) Mint() (*types.Transaction, error) {
	return _ISynthetix.Contract.Mint(&_ISynthetix.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(bool)
func (_ISynthetix *ISynthetixTransactorSession) Mint() (*types.Transaction, error) {
	return _ISynthetix.Contract.Mint(&_ISynthetix.TransactOpts)
}

// MintSecondary is a paid mutator transaction binding the contract method 0x666ed4f1.
//
// Solidity: function mintSecondary(address account, uint256 amount) returns()
func (_ISynthetix *ISynthetixTransactor) MintSecondary(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "mintSecondary", account, amount)
}

// MintSecondary is a paid mutator transaction binding the contract method 0x666ed4f1.
//
// Solidity: function mintSecondary(address account, uint256 amount) returns()
func (_ISynthetix *ISynthetixSession) MintSecondary(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.MintSecondary(&_ISynthetix.TransactOpts, account, amount)
}

// MintSecondary is a paid mutator transaction binding the contract method 0x666ed4f1.
//
// Solidity: function mintSecondary(address account, uint256 amount) returns()
func (_ISynthetix *ISynthetixTransactorSession) MintSecondary(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.MintSecondary(&_ISynthetix.TransactOpts, account, amount)
}

// MintSecondaryRewards is a paid mutator transaction binding the contract method 0xd8a1f76f.
//
// Solidity: function mintSecondaryRewards(uint256 amount) returns()
func (_ISynthetix *ISynthetixTransactor) MintSecondaryRewards(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "mintSecondaryRewards", amount)
}

// MintSecondaryRewards is a paid mutator transaction binding the contract method 0xd8a1f76f.
//
// Solidity: function mintSecondaryRewards(uint256 amount) returns()
func (_ISynthetix *ISynthetixSession) MintSecondaryRewards(amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.MintSecondaryRewards(&_ISynthetix.TransactOpts, amount)
}

// MintSecondaryRewards is a paid mutator transaction binding the contract method 0xd8a1f76f.
//
// Solidity: function mintSecondaryRewards(uint256 amount) returns()
func (_ISynthetix *ISynthetixTransactorSession) MintSecondaryRewards(amount *big.Int) (*types.Transaction, error) {
	return _ISynthetix.Contract.MintSecondaryRewards(&_ISynthetix.TransactOpts, amount)
}

// Settle is a paid mutator transaction binding the contract method 0x987757dd.
//
// Solidity: function settle(bytes32 currencyKey) returns(uint256 reclaimed, uint256 refunded, uint256 numEntries)
func (_ISynthetix *ISynthetixTransactor) Settle(opts *bind.TransactOpts, currencyKey [32]byte) (*types.Transaction, error) {
	return _ISynthetix.contract.Transact(opts, "settle", currencyKey)
}

// Settle is a paid mutator transaction binding the contract method 0x987757dd.
//
// Solidity: function settle(bytes32 currencyKey) returns(uint256 reclaimed, uint256 refunded, uint256 numEntries)
func (_ISynthetix *ISynthetixSession) Settle(currencyKey [32]byte) (*types.Transaction, error) {
	return _ISynthetix.Contract.Settle(&_ISynthetix.TransactOpts, currencyKey)
}

// Settle is a paid mutator transaction binding the contract method 0x987757dd.
//
// Solidity: function settle(bytes32 currencyKey) returns(uint256 reclaimed, uint256 refunded, uint256 numEntries)
func (_ISynthetix *ISynthetixTransactorSession) Settle(currencyKey [32]byte) (*types.Transaction, error) {
	return _ISynthetix.Contract.Settle(&_ISynthetix.TransactOpts, currencyKey)
}

// ISynthetixStateABI is the input ABI used to generate the binding from.
const ISynthetixStateABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"appendDebtLedgerValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"clearIssuanceData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"debtLedger\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"debtLedgerLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"decrementTotalIssuerCount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasIssued\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"incrementTotalIssuerCount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"issuanceData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialDebtOwnership\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtEntryIndex\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDebtLedgerEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialDebtOwnership\",\"type\":\"uint256\"}],\"name\":\"setCurrentIssuanceData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ISynthetixStateFuncSigs maps the 4-byte function signature to its string representation.
var ISynthetixStateFuncSigs = map[string]string{
	"3d31e97b": "appendDebtLedgerValue(uint256)",
	"b16c09f0": "clearIssuanceData(address)",
	"08d95cd5": "debtLedger(uint256)",
	"cd92eba9": "debtLedgerLength()",
	"ba08f299": "decrementTotalIssuerCount()",
	"b992812e": "hasIssued(address)",
	"1bfba595": "incrementTotalIssuerCount()",
	"8b3f8088": "issuanceData(address)",
	"46317712": "lastDebtLedgerEntry()",
	"a764eb45": "setCurrentIssuanceData(address,uint256)",
}

// ISynthetixState is an auto generated Go binding around an Ethereum contract.
type ISynthetixState struct {
	ISynthetixStateCaller     // Read-only binding to the contract
	ISynthetixStateTransactor // Write-only binding to the contract
	ISynthetixStateFilterer   // Log filterer for contract events
}

// ISynthetixStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISynthetixStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynthetixStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynthetixStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynthetixStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynthetixStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynthetixStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynthetixStateSession struct {
	Contract     *ISynthetixState  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISynthetixStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynthetixStateCallerSession struct {
	Contract *ISynthetixStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ISynthetixStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynthetixStateTransactorSession struct {
	Contract     *ISynthetixStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ISynthetixStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISynthetixStateRaw struct {
	Contract *ISynthetixState // Generic contract binding to access the raw methods on
}

// ISynthetixStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynthetixStateCallerRaw struct {
	Contract *ISynthetixStateCaller // Generic read-only contract binding to access the raw methods on
}

// ISynthetixStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynthetixStateTransactorRaw struct {
	Contract *ISynthetixStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISynthetixState creates a new instance of ISynthetixState, bound to a specific deployed contract.
func NewISynthetixState(address common.Address, backend bind.ContractBackend) (*ISynthetixState, error) {
	contract, err := bindISynthetixState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynthetixState{ISynthetixStateCaller: ISynthetixStateCaller{contract: contract}, ISynthetixStateTransactor: ISynthetixStateTransactor{contract: contract}, ISynthetixStateFilterer: ISynthetixStateFilterer{contract: contract}}, nil
}

// NewISynthetixStateCaller creates a new read-only instance of ISynthetixState, bound to a specific deployed contract.
func NewISynthetixStateCaller(address common.Address, caller bind.ContractCaller) (*ISynthetixStateCaller, error) {
	contract, err := bindISynthetixState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynthetixStateCaller{contract: contract}, nil
}

// NewISynthetixStateTransactor creates a new write-only instance of ISynthetixState, bound to a specific deployed contract.
func NewISynthetixStateTransactor(address common.Address, transactor bind.ContractTransactor) (*ISynthetixStateTransactor, error) {
	contract, err := bindISynthetixState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynthetixStateTransactor{contract: contract}, nil
}

// NewISynthetixStateFilterer creates a new log filterer instance of ISynthetixState, bound to a specific deployed contract.
func NewISynthetixStateFilterer(address common.Address, filterer bind.ContractFilterer) (*ISynthetixStateFilterer, error) {
	contract, err := bindISynthetixState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynthetixStateFilterer{contract: contract}, nil
}

// bindISynthetixState binds a generic wrapper to an already deployed contract.
func bindISynthetixState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISynthetixStateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynthetixState *ISynthetixStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynthetixState.Contract.ISynthetixStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynthetixState *ISynthetixStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynthetixState.Contract.ISynthetixStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynthetixState *ISynthetixStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynthetixState.Contract.ISynthetixStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynthetixState *ISynthetixStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynthetixState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynthetixState *ISynthetixStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynthetixState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynthetixState *ISynthetixStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynthetixState.Contract.contract.Transact(opts, method, params...)
}

// DebtLedger is a free data retrieval call binding the contract method 0x08d95cd5.
//
// Solidity: function debtLedger(uint256 index) view returns(uint256)
func (_ISynthetixState *ISynthetixStateCaller) DebtLedger(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISynthetixState.contract.Call(opts, &out, "debtLedger", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtLedger is a free data retrieval call binding the contract method 0x08d95cd5.
//
// Solidity: function debtLedger(uint256 index) view returns(uint256)
func (_ISynthetixState *ISynthetixStateSession) DebtLedger(index *big.Int) (*big.Int, error) {
	return _ISynthetixState.Contract.DebtLedger(&_ISynthetixState.CallOpts, index)
}

// DebtLedger is a free data retrieval call binding the contract method 0x08d95cd5.
//
// Solidity: function debtLedger(uint256 index) view returns(uint256)
func (_ISynthetixState *ISynthetixStateCallerSession) DebtLedger(index *big.Int) (*big.Int, error) {
	return _ISynthetixState.Contract.DebtLedger(&_ISynthetixState.CallOpts, index)
}

// DebtLedgerLength is a free data retrieval call binding the contract method 0xcd92eba9.
//
// Solidity: function debtLedgerLength() view returns(uint256)
func (_ISynthetixState *ISynthetixStateCaller) DebtLedgerLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISynthetixState.contract.Call(opts, &out, "debtLedgerLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtLedgerLength is a free data retrieval call binding the contract method 0xcd92eba9.
//
// Solidity: function debtLedgerLength() view returns(uint256)
func (_ISynthetixState *ISynthetixStateSession) DebtLedgerLength() (*big.Int, error) {
	return _ISynthetixState.Contract.DebtLedgerLength(&_ISynthetixState.CallOpts)
}

// DebtLedgerLength is a free data retrieval call binding the contract method 0xcd92eba9.
//
// Solidity: function debtLedgerLength() view returns(uint256)
func (_ISynthetixState *ISynthetixStateCallerSession) DebtLedgerLength() (*big.Int, error) {
	return _ISynthetixState.Contract.DebtLedgerLength(&_ISynthetixState.CallOpts)
}

// HasIssued is a free data retrieval call binding the contract method 0xb992812e.
//
// Solidity: function hasIssued(address account) view returns(bool)
func (_ISynthetixState *ISynthetixStateCaller) HasIssued(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ISynthetixState.contract.Call(opts, &out, "hasIssued", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasIssued is a free data retrieval call binding the contract method 0xb992812e.
//
// Solidity: function hasIssued(address account) view returns(bool)
func (_ISynthetixState *ISynthetixStateSession) HasIssued(account common.Address) (bool, error) {
	return _ISynthetixState.Contract.HasIssued(&_ISynthetixState.CallOpts, account)
}

// HasIssued is a free data retrieval call binding the contract method 0xb992812e.
//
// Solidity: function hasIssued(address account) view returns(bool)
func (_ISynthetixState *ISynthetixStateCallerSession) HasIssued(account common.Address) (bool, error) {
	return _ISynthetixState.Contract.HasIssued(&_ISynthetixState.CallOpts, account)
}

// IssuanceData is a free data retrieval call binding the contract method 0x8b3f8088.
//
// Solidity: function issuanceData(address account) view returns(uint256 initialDebtOwnership, uint256 debtEntryIndex)
func (_ISynthetixState *ISynthetixStateCaller) IssuanceData(opts *bind.CallOpts, account common.Address) (struct {
	InitialDebtOwnership *big.Int
	DebtEntryIndex       *big.Int
}, error) {
	var out []interface{}
	err := _ISynthetixState.contract.Call(opts, &out, "issuanceData", account)

	outstruct := new(struct {
		InitialDebtOwnership *big.Int
		DebtEntryIndex       *big.Int
	})

	outstruct.InitialDebtOwnership = out[0].(*big.Int)
	outstruct.DebtEntryIndex = out[1].(*big.Int)

	return *outstruct, err

}

// IssuanceData is a free data retrieval call binding the contract method 0x8b3f8088.
//
// Solidity: function issuanceData(address account) view returns(uint256 initialDebtOwnership, uint256 debtEntryIndex)
func (_ISynthetixState *ISynthetixStateSession) IssuanceData(account common.Address) (struct {
	InitialDebtOwnership *big.Int
	DebtEntryIndex       *big.Int
}, error) {
	return _ISynthetixState.Contract.IssuanceData(&_ISynthetixState.CallOpts, account)
}

// IssuanceData is a free data retrieval call binding the contract method 0x8b3f8088.
//
// Solidity: function issuanceData(address account) view returns(uint256 initialDebtOwnership, uint256 debtEntryIndex)
func (_ISynthetixState *ISynthetixStateCallerSession) IssuanceData(account common.Address) (struct {
	InitialDebtOwnership *big.Int
	DebtEntryIndex       *big.Int
}, error) {
	return _ISynthetixState.Contract.IssuanceData(&_ISynthetixState.CallOpts, account)
}

// LastDebtLedgerEntry is a free data retrieval call binding the contract method 0x46317712.
//
// Solidity: function lastDebtLedgerEntry() view returns(uint256)
func (_ISynthetixState *ISynthetixStateCaller) LastDebtLedgerEntry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISynthetixState.contract.Call(opts, &out, "lastDebtLedgerEntry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastDebtLedgerEntry is a free data retrieval call binding the contract method 0x46317712.
//
// Solidity: function lastDebtLedgerEntry() view returns(uint256)
func (_ISynthetixState *ISynthetixStateSession) LastDebtLedgerEntry() (*big.Int, error) {
	return _ISynthetixState.Contract.LastDebtLedgerEntry(&_ISynthetixState.CallOpts)
}

// LastDebtLedgerEntry is a free data retrieval call binding the contract method 0x46317712.
//
// Solidity: function lastDebtLedgerEntry() view returns(uint256)
func (_ISynthetixState *ISynthetixStateCallerSession) LastDebtLedgerEntry() (*big.Int, error) {
	return _ISynthetixState.Contract.LastDebtLedgerEntry(&_ISynthetixState.CallOpts)
}

// AppendDebtLedgerValue is a paid mutator transaction binding the contract method 0x3d31e97b.
//
// Solidity: function appendDebtLedgerValue(uint256 value) returns()
func (_ISynthetixState *ISynthetixStateTransactor) AppendDebtLedgerValue(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _ISynthetixState.contract.Transact(opts, "appendDebtLedgerValue", value)
}

// AppendDebtLedgerValue is a paid mutator transaction binding the contract method 0x3d31e97b.
//
// Solidity: function appendDebtLedgerValue(uint256 value) returns()
func (_ISynthetixState *ISynthetixStateSession) AppendDebtLedgerValue(value *big.Int) (*types.Transaction, error) {
	return _ISynthetixState.Contract.AppendDebtLedgerValue(&_ISynthetixState.TransactOpts, value)
}

// AppendDebtLedgerValue is a paid mutator transaction binding the contract method 0x3d31e97b.
//
// Solidity: function appendDebtLedgerValue(uint256 value) returns()
func (_ISynthetixState *ISynthetixStateTransactorSession) AppendDebtLedgerValue(value *big.Int) (*types.Transaction, error) {
	return _ISynthetixState.Contract.AppendDebtLedgerValue(&_ISynthetixState.TransactOpts, value)
}

// ClearIssuanceData is a paid mutator transaction binding the contract method 0xb16c09f0.
//
// Solidity: function clearIssuanceData(address account) returns()
func (_ISynthetixState *ISynthetixStateTransactor) ClearIssuanceData(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ISynthetixState.contract.Transact(opts, "clearIssuanceData", account)
}

// ClearIssuanceData is a paid mutator transaction binding the contract method 0xb16c09f0.
//
// Solidity: function clearIssuanceData(address account) returns()
func (_ISynthetixState *ISynthetixStateSession) ClearIssuanceData(account common.Address) (*types.Transaction, error) {
	return _ISynthetixState.Contract.ClearIssuanceData(&_ISynthetixState.TransactOpts, account)
}

// ClearIssuanceData is a paid mutator transaction binding the contract method 0xb16c09f0.
//
// Solidity: function clearIssuanceData(address account) returns()
func (_ISynthetixState *ISynthetixStateTransactorSession) ClearIssuanceData(account common.Address) (*types.Transaction, error) {
	return _ISynthetixState.Contract.ClearIssuanceData(&_ISynthetixState.TransactOpts, account)
}

// DecrementTotalIssuerCount is a paid mutator transaction binding the contract method 0xba08f299.
//
// Solidity: function decrementTotalIssuerCount() returns()
func (_ISynthetixState *ISynthetixStateTransactor) DecrementTotalIssuerCount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynthetixState.contract.Transact(opts, "decrementTotalIssuerCount")
}

// DecrementTotalIssuerCount is a paid mutator transaction binding the contract method 0xba08f299.
//
// Solidity: function decrementTotalIssuerCount() returns()
func (_ISynthetixState *ISynthetixStateSession) DecrementTotalIssuerCount() (*types.Transaction, error) {
	return _ISynthetixState.Contract.DecrementTotalIssuerCount(&_ISynthetixState.TransactOpts)
}

// DecrementTotalIssuerCount is a paid mutator transaction binding the contract method 0xba08f299.
//
// Solidity: function decrementTotalIssuerCount() returns()
func (_ISynthetixState *ISynthetixStateTransactorSession) DecrementTotalIssuerCount() (*types.Transaction, error) {
	return _ISynthetixState.Contract.DecrementTotalIssuerCount(&_ISynthetixState.TransactOpts)
}

// IncrementTotalIssuerCount is a paid mutator transaction binding the contract method 0x1bfba595.
//
// Solidity: function incrementTotalIssuerCount() returns()
func (_ISynthetixState *ISynthetixStateTransactor) IncrementTotalIssuerCount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynthetixState.contract.Transact(opts, "incrementTotalIssuerCount")
}

// IncrementTotalIssuerCount is a paid mutator transaction binding the contract method 0x1bfba595.
//
// Solidity: function incrementTotalIssuerCount() returns()
func (_ISynthetixState *ISynthetixStateSession) IncrementTotalIssuerCount() (*types.Transaction, error) {
	return _ISynthetixState.Contract.IncrementTotalIssuerCount(&_ISynthetixState.TransactOpts)
}

// IncrementTotalIssuerCount is a paid mutator transaction binding the contract method 0x1bfba595.
//
// Solidity: function incrementTotalIssuerCount() returns()
func (_ISynthetixState *ISynthetixStateTransactorSession) IncrementTotalIssuerCount() (*types.Transaction, error) {
	return _ISynthetixState.Contract.IncrementTotalIssuerCount(&_ISynthetixState.TransactOpts)
}

// SetCurrentIssuanceData is a paid mutator transaction binding the contract method 0xa764eb45.
//
// Solidity: function setCurrentIssuanceData(address account, uint256 initialDebtOwnership) returns()
func (_ISynthetixState *ISynthetixStateTransactor) SetCurrentIssuanceData(opts *bind.TransactOpts, account common.Address, initialDebtOwnership *big.Int) (*types.Transaction, error) {
	return _ISynthetixState.contract.Transact(opts, "setCurrentIssuanceData", account, initialDebtOwnership)
}

// SetCurrentIssuanceData is a paid mutator transaction binding the contract method 0xa764eb45.
//
// Solidity: function setCurrentIssuanceData(address account, uint256 initialDebtOwnership) returns()
func (_ISynthetixState *ISynthetixStateSession) SetCurrentIssuanceData(account common.Address, initialDebtOwnership *big.Int) (*types.Transaction, error) {
	return _ISynthetixState.Contract.SetCurrentIssuanceData(&_ISynthetixState.TransactOpts, account, initialDebtOwnership)
}

// SetCurrentIssuanceData is a paid mutator transaction binding the contract method 0xa764eb45.
//
// Solidity: function setCurrentIssuanceData(address account, uint256 initialDebtOwnership) returns()
func (_ISynthetixState *ISynthetixStateTransactorSession) SetCurrentIssuanceData(account common.Address, initialDebtOwnership *big.Int) (*types.Transaction, error) {
	return _ISynthetixState.Contract.SetCurrentIssuanceData(&_ISynthetixState.TransactOpts, account, initialDebtOwnership)
}

// ISystemStatusABI is the input ABI used to generate the binding from.
const ISystemStatusABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"section\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accessControl\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"canSuspend\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canResume\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requireExchangeActive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requireIssuanceActive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"requireSynthActive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"}],\"name\":\"requireSynthsActive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requireSystemActive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"reason\",\"type\":\"uint256\"}],\"name\":\"suspendSynth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"synthSuspension\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"suspended\",\"type\":\"bool\"},{\"internalType\":\"uint248\",\"name\":\"reason\",\"type\":\"uint248\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"section\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"canSuspend\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canResume\",\"type\":\"bool\"}],\"name\":\"updateAccessControl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ISystemStatusFuncSigs maps the 4-byte function signature to its string representation.
var ISystemStatusFuncSigs = map[string]string{
	"20f2bf00": "accessControl(bytes32,address)",
	"7118d431": "requireExchangeActive()",
	"7c312541": "requireIssuanceActive()",
	"42a28e21": "requireSynthActive(bytes32)",
	"6132eba4": "requireSynthsActive(bytes32,bytes32)",
	"086dabd1": "requireSystemActive()",
	"abc0bb6e": "suspendSynth(bytes32,uint256)",
	"7243bc2c": "synthSuspension(bytes32)",
	"48bf1971": "updateAccessControl(bytes32,address,bool,bool)",
}

// ISystemStatus is an auto generated Go binding around an Ethereum contract.
type ISystemStatus struct {
	ISystemStatusCaller     // Read-only binding to the contract
	ISystemStatusTransactor // Write-only binding to the contract
	ISystemStatusFilterer   // Log filterer for contract events
}

// ISystemStatusCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISystemStatusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemStatusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISystemStatusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemStatusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISystemStatusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemStatusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISystemStatusSession struct {
	Contract     *ISystemStatus    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISystemStatusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISystemStatusCallerSession struct {
	Contract *ISystemStatusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ISystemStatusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISystemStatusTransactorSession struct {
	Contract     *ISystemStatusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ISystemStatusRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISystemStatusRaw struct {
	Contract *ISystemStatus // Generic contract binding to access the raw methods on
}

// ISystemStatusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISystemStatusCallerRaw struct {
	Contract *ISystemStatusCaller // Generic read-only contract binding to access the raw methods on
}

// ISystemStatusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISystemStatusTransactorRaw struct {
	Contract *ISystemStatusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISystemStatus creates a new instance of ISystemStatus, bound to a specific deployed contract.
func NewISystemStatus(address common.Address, backend bind.ContractBackend) (*ISystemStatus, error) {
	contract, err := bindISystemStatus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISystemStatus{ISystemStatusCaller: ISystemStatusCaller{contract: contract}, ISystemStatusTransactor: ISystemStatusTransactor{contract: contract}, ISystemStatusFilterer: ISystemStatusFilterer{contract: contract}}, nil
}

// NewISystemStatusCaller creates a new read-only instance of ISystemStatus, bound to a specific deployed contract.
func NewISystemStatusCaller(address common.Address, caller bind.ContractCaller) (*ISystemStatusCaller, error) {
	contract, err := bindISystemStatus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemStatusCaller{contract: contract}, nil
}

// NewISystemStatusTransactor creates a new write-only instance of ISystemStatus, bound to a specific deployed contract.
func NewISystemStatusTransactor(address common.Address, transactor bind.ContractTransactor) (*ISystemStatusTransactor, error) {
	contract, err := bindISystemStatus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemStatusTransactor{contract: contract}, nil
}

// NewISystemStatusFilterer creates a new log filterer instance of ISystemStatus, bound to a specific deployed contract.
func NewISystemStatusFilterer(address common.Address, filterer bind.ContractFilterer) (*ISystemStatusFilterer, error) {
	contract, err := bindISystemStatus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISystemStatusFilterer{contract: contract}, nil
}

// bindISystemStatus binds a generic wrapper to an already deployed contract.
func bindISystemStatus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISystemStatusABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemStatus *ISystemStatusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemStatus.Contract.ISystemStatusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemStatus *ISystemStatusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemStatus.Contract.ISystemStatusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemStatus *ISystemStatusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemStatus.Contract.ISystemStatusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemStatus *ISystemStatusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemStatus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemStatus *ISystemStatusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemStatus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemStatus *ISystemStatusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemStatus.Contract.contract.Transact(opts, method, params...)
}

// AccessControl is a free data retrieval call binding the contract method 0x20f2bf00.
//
// Solidity: function accessControl(bytes32 section, address account) view returns(bool canSuspend, bool canResume)
func (_ISystemStatus *ISystemStatusCaller) AccessControl(opts *bind.CallOpts, section [32]byte, account common.Address) (struct {
	CanSuspend bool
	CanResume  bool
}, error) {
	var out []interface{}
	err := _ISystemStatus.contract.Call(opts, &out, "accessControl", section, account)

	outstruct := new(struct {
		CanSuspend bool
		CanResume  bool
	})

	outstruct.CanSuspend = out[0].(bool)
	outstruct.CanResume = out[1].(bool)

	return *outstruct, err

}

// AccessControl is a free data retrieval call binding the contract method 0x20f2bf00.
//
// Solidity: function accessControl(bytes32 section, address account) view returns(bool canSuspend, bool canResume)
func (_ISystemStatus *ISystemStatusSession) AccessControl(section [32]byte, account common.Address) (struct {
	CanSuspend bool
	CanResume  bool
}, error) {
	return _ISystemStatus.Contract.AccessControl(&_ISystemStatus.CallOpts, section, account)
}

// AccessControl is a free data retrieval call binding the contract method 0x20f2bf00.
//
// Solidity: function accessControl(bytes32 section, address account) view returns(bool canSuspend, bool canResume)
func (_ISystemStatus *ISystemStatusCallerSession) AccessControl(section [32]byte, account common.Address) (struct {
	CanSuspend bool
	CanResume  bool
}, error) {
	return _ISystemStatus.Contract.AccessControl(&_ISystemStatus.CallOpts, section, account)
}

// RequireExchangeActive is a free data retrieval call binding the contract method 0x7118d431.
//
// Solidity: function requireExchangeActive() view returns()
func (_ISystemStatus *ISystemStatusCaller) RequireExchangeActive(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ISystemStatus.contract.Call(opts, &out, "requireExchangeActive")

	if err != nil {
		return err
	}

	return err

}

// RequireExchangeActive is a free data retrieval call binding the contract method 0x7118d431.
//
// Solidity: function requireExchangeActive() view returns()
func (_ISystemStatus *ISystemStatusSession) RequireExchangeActive() error {
	return _ISystemStatus.Contract.RequireExchangeActive(&_ISystemStatus.CallOpts)
}

// RequireExchangeActive is a free data retrieval call binding the contract method 0x7118d431.
//
// Solidity: function requireExchangeActive() view returns()
func (_ISystemStatus *ISystemStatusCallerSession) RequireExchangeActive() error {
	return _ISystemStatus.Contract.RequireExchangeActive(&_ISystemStatus.CallOpts)
}

// RequireIssuanceActive is a free data retrieval call binding the contract method 0x7c312541.
//
// Solidity: function requireIssuanceActive() view returns()
func (_ISystemStatus *ISystemStatusCaller) RequireIssuanceActive(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ISystemStatus.contract.Call(opts, &out, "requireIssuanceActive")

	if err != nil {
		return err
	}

	return err

}

// RequireIssuanceActive is a free data retrieval call binding the contract method 0x7c312541.
//
// Solidity: function requireIssuanceActive() view returns()
func (_ISystemStatus *ISystemStatusSession) RequireIssuanceActive() error {
	return _ISystemStatus.Contract.RequireIssuanceActive(&_ISystemStatus.CallOpts)
}

// RequireIssuanceActive is a free data retrieval call binding the contract method 0x7c312541.
//
// Solidity: function requireIssuanceActive() view returns()
func (_ISystemStatus *ISystemStatusCallerSession) RequireIssuanceActive() error {
	return _ISystemStatus.Contract.RequireIssuanceActive(&_ISystemStatus.CallOpts)
}

// RequireSynthActive is a free data retrieval call binding the contract method 0x42a28e21.
//
// Solidity: function requireSynthActive(bytes32 currencyKey) view returns()
func (_ISystemStatus *ISystemStatusCaller) RequireSynthActive(opts *bind.CallOpts, currencyKey [32]byte) error {
	var out []interface{}
	err := _ISystemStatus.contract.Call(opts, &out, "requireSynthActive", currencyKey)

	if err != nil {
		return err
	}

	return err

}

// RequireSynthActive is a free data retrieval call binding the contract method 0x42a28e21.
//
// Solidity: function requireSynthActive(bytes32 currencyKey) view returns()
func (_ISystemStatus *ISystemStatusSession) RequireSynthActive(currencyKey [32]byte) error {
	return _ISystemStatus.Contract.RequireSynthActive(&_ISystemStatus.CallOpts, currencyKey)
}

// RequireSynthActive is a free data retrieval call binding the contract method 0x42a28e21.
//
// Solidity: function requireSynthActive(bytes32 currencyKey) view returns()
func (_ISystemStatus *ISystemStatusCallerSession) RequireSynthActive(currencyKey [32]byte) error {
	return _ISystemStatus.Contract.RequireSynthActive(&_ISystemStatus.CallOpts, currencyKey)
}

// RequireSynthsActive is a free data retrieval call binding the contract method 0x6132eba4.
//
// Solidity: function requireSynthsActive(bytes32 sourceCurrencyKey, bytes32 destinationCurrencyKey) view returns()
func (_ISystemStatus *ISystemStatusCaller) RequireSynthsActive(opts *bind.CallOpts, sourceCurrencyKey [32]byte, destinationCurrencyKey [32]byte) error {
	var out []interface{}
	err := _ISystemStatus.contract.Call(opts, &out, "requireSynthsActive", sourceCurrencyKey, destinationCurrencyKey)

	if err != nil {
		return err
	}

	return err

}

// RequireSynthsActive is a free data retrieval call binding the contract method 0x6132eba4.
//
// Solidity: function requireSynthsActive(bytes32 sourceCurrencyKey, bytes32 destinationCurrencyKey) view returns()
func (_ISystemStatus *ISystemStatusSession) RequireSynthsActive(sourceCurrencyKey [32]byte, destinationCurrencyKey [32]byte) error {
	return _ISystemStatus.Contract.RequireSynthsActive(&_ISystemStatus.CallOpts, sourceCurrencyKey, destinationCurrencyKey)
}

// RequireSynthsActive is a free data retrieval call binding the contract method 0x6132eba4.
//
// Solidity: function requireSynthsActive(bytes32 sourceCurrencyKey, bytes32 destinationCurrencyKey) view returns()
func (_ISystemStatus *ISystemStatusCallerSession) RequireSynthsActive(sourceCurrencyKey [32]byte, destinationCurrencyKey [32]byte) error {
	return _ISystemStatus.Contract.RequireSynthsActive(&_ISystemStatus.CallOpts, sourceCurrencyKey, destinationCurrencyKey)
}

// RequireSystemActive is a free data retrieval call binding the contract method 0x086dabd1.
//
// Solidity: function requireSystemActive() view returns()
func (_ISystemStatus *ISystemStatusCaller) RequireSystemActive(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ISystemStatus.contract.Call(opts, &out, "requireSystemActive")

	if err != nil {
		return err
	}

	return err

}

// RequireSystemActive is a free data retrieval call binding the contract method 0x086dabd1.
//
// Solidity: function requireSystemActive() view returns()
func (_ISystemStatus *ISystemStatusSession) RequireSystemActive() error {
	return _ISystemStatus.Contract.RequireSystemActive(&_ISystemStatus.CallOpts)
}

// RequireSystemActive is a free data retrieval call binding the contract method 0x086dabd1.
//
// Solidity: function requireSystemActive() view returns()
func (_ISystemStatus *ISystemStatusCallerSession) RequireSystemActive() error {
	return _ISystemStatus.Contract.RequireSystemActive(&_ISystemStatus.CallOpts)
}

// SynthSuspension is a free data retrieval call binding the contract method 0x7243bc2c.
//
// Solidity: function synthSuspension(bytes32 currencyKey) view returns(bool suspended, uint248 reason)
func (_ISystemStatus *ISystemStatusCaller) SynthSuspension(opts *bind.CallOpts, currencyKey [32]byte) (struct {
	Suspended bool
	Reason    *big.Int
}, error) {
	var out []interface{}
	err := _ISystemStatus.contract.Call(opts, &out, "synthSuspension", currencyKey)

	outstruct := new(struct {
		Suspended bool
		Reason    *big.Int
	})

	outstruct.Suspended = out[0].(bool)
	outstruct.Reason = out[1].(*big.Int)

	return *outstruct, err

}

// SynthSuspension is a free data retrieval call binding the contract method 0x7243bc2c.
//
// Solidity: function synthSuspension(bytes32 currencyKey) view returns(bool suspended, uint248 reason)
func (_ISystemStatus *ISystemStatusSession) SynthSuspension(currencyKey [32]byte) (struct {
	Suspended bool
	Reason    *big.Int
}, error) {
	return _ISystemStatus.Contract.SynthSuspension(&_ISystemStatus.CallOpts, currencyKey)
}

// SynthSuspension is a free data retrieval call binding the contract method 0x7243bc2c.
//
// Solidity: function synthSuspension(bytes32 currencyKey) view returns(bool suspended, uint248 reason)
func (_ISystemStatus *ISystemStatusCallerSession) SynthSuspension(currencyKey [32]byte) (struct {
	Suspended bool
	Reason    *big.Int
}, error) {
	return _ISystemStatus.Contract.SynthSuspension(&_ISystemStatus.CallOpts, currencyKey)
}

// SuspendSynth is a paid mutator transaction binding the contract method 0xabc0bb6e.
//
// Solidity: function suspendSynth(bytes32 currencyKey, uint256 reason) returns()
func (_ISystemStatus *ISystemStatusTransactor) SuspendSynth(opts *bind.TransactOpts, currencyKey [32]byte, reason *big.Int) (*types.Transaction, error) {
	return _ISystemStatus.contract.Transact(opts, "suspendSynth", currencyKey, reason)
}

// SuspendSynth is a paid mutator transaction binding the contract method 0xabc0bb6e.
//
// Solidity: function suspendSynth(bytes32 currencyKey, uint256 reason) returns()
func (_ISystemStatus *ISystemStatusSession) SuspendSynth(currencyKey [32]byte, reason *big.Int) (*types.Transaction, error) {
	return _ISystemStatus.Contract.SuspendSynth(&_ISystemStatus.TransactOpts, currencyKey, reason)
}

// SuspendSynth is a paid mutator transaction binding the contract method 0xabc0bb6e.
//
// Solidity: function suspendSynth(bytes32 currencyKey, uint256 reason) returns()
func (_ISystemStatus *ISystemStatusTransactorSession) SuspendSynth(currencyKey [32]byte, reason *big.Int) (*types.Transaction, error) {
	return _ISystemStatus.Contract.SuspendSynth(&_ISystemStatus.TransactOpts, currencyKey, reason)
}

// UpdateAccessControl is a paid mutator transaction binding the contract method 0x48bf1971.
//
// Solidity: function updateAccessControl(bytes32 section, address account, bool canSuspend, bool canResume) returns()
func (_ISystemStatus *ISystemStatusTransactor) UpdateAccessControl(opts *bind.TransactOpts, section [32]byte, account common.Address, canSuspend bool, canResume bool) (*types.Transaction, error) {
	return _ISystemStatus.contract.Transact(opts, "updateAccessControl", section, account, canSuspend, canResume)
}

// UpdateAccessControl is a paid mutator transaction binding the contract method 0x48bf1971.
//
// Solidity: function updateAccessControl(bytes32 section, address account, bool canSuspend, bool canResume) returns()
func (_ISystemStatus *ISystemStatusSession) UpdateAccessControl(section [32]byte, account common.Address, canSuspend bool, canResume bool) (*types.Transaction, error) {
	return _ISystemStatus.Contract.UpdateAccessControl(&_ISystemStatus.TransactOpts, section, account, canSuspend, canResume)
}

// UpdateAccessControl is a paid mutator transaction binding the contract method 0x48bf1971.
//
// Solidity: function updateAccessControl(bytes32 section, address account, bool canSuspend, bool canResume) returns()
func (_ISystemStatus *ISystemStatusTransactorSession) UpdateAccessControl(section [32]byte, account common.Address, canSuspend bool, canResume bool) (*types.Transaction, error) {
	return _ISystemStatus.Contract.UpdateAccessControl(&_ISystemStatus.TransactOpts, section, account, canSuspend, canResume)
}

// IVirtualSynthABI is the input ABI used to generate the binding from.
const IVirtualSynthABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"readyToSettle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secsLeftInWaitingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"settle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"settled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"synth\",\"outputs\":[{\"internalType\":\"contractISynth\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IVirtualSynthFuncSigs maps the 4-byte function signature to its string representation.
var IVirtualSynthFuncSigs = map[string]string{
	"3af9e669": "balanceOfUnderlying(address)",
	"2c4e722e": "rate()",
	"78f2ac26": "readyToSettle()",
	"4be37cea": "secsLeftInWaitingPeriod()",
	"6a256b29": "settle(address)",
	"8f775839": "settled()",
	"115f4fee": "synth()",
}

// IVirtualSynth is an auto generated Go binding around an Ethereum contract.
type IVirtualSynth struct {
	IVirtualSynthCaller     // Read-only binding to the contract
	IVirtualSynthTransactor // Write-only binding to the contract
	IVirtualSynthFilterer   // Log filterer for contract events
}

// IVirtualSynthCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVirtualSynthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVirtualSynthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVirtualSynthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVirtualSynthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVirtualSynthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVirtualSynthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVirtualSynthSession struct {
	Contract     *IVirtualSynth    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVirtualSynthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVirtualSynthCallerSession struct {
	Contract *IVirtualSynthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IVirtualSynthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVirtualSynthTransactorSession struct {
	Contract     *IVirtualSynthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IVirtualSynthRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVirtualSynthRaw struct {
	Contract *IVirtualSynth // Generic contract binding to access the raw methods on
}

// IVirtualSynthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVirtualSynthCallerRaw struct {
	Contract *IVirtualSynthCaller // Generic read-only contract binding to access the raw methods on
}

// IVirtualSynthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVirtualSynthTransactorRaw struct {
	Contract *IVirtualSynthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVirtualSynth creates a new instance of IVirtualSynth, bound to a specific deployed contract.
func NewIVirtualSynth(address common.Address, backend bind.ContractBackend) (*IVirtualSynth, error) {
	contract, err := bindIVirtualSynth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVirtualSynth{IVirtualSynthCaller: IVirtualSynthCaller{contract: contract}, IVirtualSynthTransactor: IVirtualSynthTransactor{contract: contract}, IVirtualSynthFilterer: IVirtualSynthFilterer{contract: contract}}, nil
}

// NewIVirtualSynthCaller creates a new read-only instance of IVirtualSynth, bound to a specific deployed contract.
func NewIVirtualSynthCaller(address common.Address, caller bind.ContractCaller) (*IVirtualSynthCaller, error) {
	contract, err := bindIVirtualSynth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVirtualSynthCaller{contract: contract}, nil
}

// NewIVirtualSynthTransactor creates a new write-only instance of IVirtualSynth, bound to a specific deployed contract.
func NewIVirtualSynthTransactor(address common.Address, transactor bind.ContractTransactor) (*IVirtualSynthTransactor, error) {
	contract, err := bindIVirtualSynth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVirtualSynthTransactor{contract: contract}, nil
}

// NewIVirtualSynthFilterer creates a new log filterer instance of IVirtualSynth, bound to a specific deployed contract.
func NewIVirtualSynthFilterer(address common.Address, filterer bind.ContractFilterer) (*IVirtualSynthFilterer, error) {
	contract, err := bindIVirtualSynth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVirtualSynthFilterer{contract: contract}, nil
}

// bindIVirtualSynth binds a generic wrapper to an already deployed contract.
func bindIVirtualSynth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVirtualSynthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVirtualSynth *IVirtualSynthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVirtualSynth.Contract.IVirtualSynthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVirtualSynth *IVirtualSynthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVirtualSynth.Contract.IVirtualSynthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVirtualSynth *IVirtualSynthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVirtualSynth.Contract.IVirtualSynthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVirtualSynth *IVirtualSynthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVirtualSynth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVirtualSynth *IVirtualSynthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVirtualSynth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVirtualSynth *IVirtualSynthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVirtualSynth.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfUnderlying is a free data retrieval call binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address account) view returns(uint256)
func (_IVirtualSynth *IVirtualSynthCaller) BalanceOfUnderlying(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVirtualSynth.contract.Call(opts, &out, "balanceOfUnderlying", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfUnderlying is a free data retrieval call binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address account) view returns(uint256)
func (_IVirtualSynth *IVirtualSynthSession) BalanceOfUnderlying(account common.Address) (*big.Int, error) {
	return _IVirtualSynth.Contract.BalanceOfUnderlying(&_IVirtualSynth.CallOpts, account)
}

// BalanceOfUnderlying is a free data retrieval call binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address account) view returns(uint256)
func (_IVirtualSynth *IVirtualSynthCallerSession) BalanceOfUnderlying(account common.Address) (*big.Int, error) {
	return _IVirtualSynth.Contract.BalanceOfUnderlying(&_IVirtualSynth.CallOpts, account)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_IVirtualSynth *IVirtualSynthCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVirtualSynth.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_IVirtualSynth *IVirtualSynthSession) Rate() (*big.Int, error) {
	return _IVirtualSynth.Contract.Rate(&_IVirtualSynth.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_IVirtualSynth *IVirtualSynthCallerSession) Rate() (*big.Int, error) {
	return _IVirtualSynth.Contract.Rate(&_IVirtualSynth.CallOpts)
}

// ReadyToSettle is a free data retrieval call binding the contract method 0x78f2ac26.
//
// Solidity: function readyToSettle() view returns(bool)
func (_IVirtualSynth *IVirtualSynthCaller) ReadyToSettle(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IVirtualSynth.contract.Call(opts, &out, "readyToSettle")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ReadyToSettle is a free data retrieval call binding the contract method 0x78f2ac26.
//
// Solidity: function readyToSettle() view returns(bool)
func (_IVirtualSynth *IVirtualSynthSession) ReadyToSettle() (bool, error) {
	return _IVirtualSynth.Contract.ReadyToSettle(&_IVirtualSynth.CallOpts)
}

// ReadyToSettle is a free data retrieval call binding the contract method 0x78f2ac26.
//
// Solidity: function readyToSettle() view returns(bool)
func (_IVirtualSynth *IVirtualSynthCallerSession) ReadyToSettle() (bool, error) {
	return _IVirtualSynth.Contract.ReadyToSettle(&_IVirtualSynth.CallOpts)
}

// SecsLeftInWaitingPeriod is a free data retrieval call binding the contract method 0x4be37cea.
//
// Solidity: function secsLeftInWaitingPeriod() view returns(uint256)
func (_IVirtualSynth *IVirtualSynthCaller) SecsLeftInWaitingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVirtualSynth.contract.Call(opts, &out, "secsLeftInWaitingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecsLeftInWaitingPeriod is a free data retrieval call binding the contract method 0x4be37cea.
//
// Solidity: function secsLeftInWaitingPeriod() view returns(uint256)
func (_IVirtualSynth *IVirtualSynthSession) SecsLeftInWaitingPeriod() (*big.Int, error) {
	return _IVirtualSynth.Contract.SecsLeftInWaitingPeriod(&_IVirtualSynth.CallOpts)
}

// SecsLeftInWaitingPeriod is a free data retrieval call binding the contract method 0x4be37cea.
//
// Solidity: function secsLeftInWaitingPeriod() view returns(uint256)
func (_IVirtualSynth *IVirtualSynthCallerSession) SecsLeftInWaitingPeriod() (*big.Int, error) {
	return _IVirtualSynth.Contract.SecsLeftInWaitingPeriod(&_IVirtualSynth.CallOpts)
}

// Settled is a free data retrieval call binding the contract method 0x8f775839.
//
// Solidity: function settled() view returns(bool)
func (_IVirtualSynth *IVirtualSynthCaller) Settled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IVirtualSynth.contract.Call(opts, &out, "settled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Settled is a free data retrieval call binding the contract method 0x8f775839.
//
// Solidity: function settled() view returns(bool)
func (_IVirtualSynth *IVirtualSynthSession) Settled() (bool, error) {
	return _IVirtualSynth.Contract.Settled(&_IVirtualSynth.CallOpts)
}

// Settled is a free data retrieval call binding the contract method 0x8f775839.
//
// Solidity: function settled() view returns(bool)
func (_IVirtualSynth *IVirtualSynthCallerSession) Settled() (bool, error) {
	return _IVirtualSynth.Contract.Settled(&_IVirtualSynth.CallOpts)
}

// Synth is a free data retrieval call binding the contract method 0x115f4fee.
//
// Solidity: function synth() view returns(address)
func (_IVirtualSynth *IVirtualSynthCaller) Synth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IVirtualSynth.contract.Call(opts, &out, "synth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Synth is a free data retrieval call binding the contract method 0x115f4fee.
//
// Solidity: function synth() view returns(address)
func (_IVirtualSynth *IVirtualSynthSession) Synth() (common.Address, error) {
	return _IVirtualSynth.Contract.Synth(&_IVirtualSynth.CallOpts)
}

// Synth is a free data retrieval call binding the contract method 0x115f4fee.
//
// Solidity: function synth() view returns(address)
func (_IVirtualSynth *IVirtualSynthCallerSession) Synth() (common.Address, error) {
	return _IVirtualSynth.Contract.Synth(&_IVirtualSynth.CallOpts)
}

// Settle is a paid mutator transaction binding the contract method 0x6a256b29.
//
// Solidity: function settle(address account) returns()
func (_IVirtualSynth *IVirtualSynthTransactor) Settle(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IVirtualSynth.contract.Transact(opts, "settle", account)
}

// Settle is a paid mutator transaction binding the contract method 0x6a256b29.
//
// Solidity: function settle(address account) returns()
func (_IVirtualSynth *IVirtualSynthSession) Settle(account common.Address) (*types.Transaction, error) {
	return _IVirtualSynth.Contract.Settle(&_IVirtualSynth.TransactOpts, account)
}

// Settle is a paid mutator transaction binding the contract method 0x6a256b29.
//
// Solidity: function settle(address account) returns()
func (_IVirtualSynth *IVirtualSynthTransactorSession) Settle(account common.Address) (*types.Transaction, error) {
	return _IVirtualSynth.Contract.Settle(&_IVirtualSynth.TransactOpts, account)
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
var MathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582048e514a37d2bf8d7949aa048752d7f07f8f5a238cb502c0575e5e3dcfd45691464736f6c63430005100032"

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// MixinResolverABI is the input ABI used to generate the binding from.
const MixinResolverABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"_addressesToCache\",\"type\":\"bytes32[24]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_ADDRESSES_FROM_RESOLVER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getResolverAddressesRequired\",\"outputs\":[{\"internalType\":\"bytes32[24]\",\"name\":\"addressesRequired\",\"type\":\"bytes32[24]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractAddressResolver\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"isResolverCached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"contractAddressResolver\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"resolverAddressesRequired\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractAddressResolver\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"setResolverAndSyncCache\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MixinResolverFuncSigs maps the 4-byte function signature to its string representation.
var MixinResolverFuncSigs = map[string]string{
	"e3235c91": "MAX_ADDRESSES_FROM_RESOLVER()",
	"79ba5097": "acceptOwnership()",
	"ab49848c": "getResolverAddressesRequired()",
	"631e1444": "isResolverCached(address)",
	"1627540c": "nominateNewOwner(address)",
	"53a47bb7": "nominatedOwner()",
	"8da5cb5b": "owner()",
	"04f3bcec": "resolver()",
	"c6c9d828": "resolverAddressesRequired(uint256)",
	"3be99e6f": "setResolverAndSyncCache(address)",
}

// MixinResolver is an auto generated Go binding around an Ethereum contract.
type MixinResolver struct {
	MixinResolverCaller     // Read-only binding to the contract
	MixinResolverTransactor // Write-only binding to the contract
	MixinResolverFilterer   // Log filterer for contract events
}

// MixinResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type MixinResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MixinResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MixinResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MixinResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MixinResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MixinResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MixinResolverSession struct {
	Contract     *MixinResolver    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MixinResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MixinResolverCallerSession struct {
	Contract *MixinResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MixinResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MixinResolverTransactorSession struct {
	Contract     *MixinResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MixinResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type MixinResolverRaw struct {
	Contract *MixinResolver // Generic contract binding to access the raw methods on
}

// MixinResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MixinResolverCallerRaw struct {
	Contract *MixinResolverCaller // Generic read-only contract binding to access the raw methods on
}

// MixinResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MixinResolverTransactorRaw struct {
	Contract *MixinResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMixinResolver creates a new instance of MixinResolver, bound to a specific deployed contract.
func NewMixinResolver(address common.Address, backend bind.ContractBackend) (*MixinResolver, error) {
	contract, err := bindMixinResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MixinResolver{MixinResolverCaller: MixinResolverCaller{contract: contract}, MixinResolverTransactor: MixinResolverTransactor{contract: contract}, MixinResolverFilterer: MixinResolverFilterer{contract: contract}}, nil
}

// NewMixinResolverCaller creates a new read-only instance of MixinResolver, bound to a specific deployed contract.
func NewMixinResolverCaller(address common.Address, caller bind.ContractCaller) (*MixinResolverCaller, error) {
	contract, err := bindMixinResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MixinResolverCaller{contract: contract}, nil
}

// NewMixinResolverTransactor creates a new write-only instance of MixinResolver, bound to a specific deployed contract.
func NewMixinResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*MixinResolverTransactor, error) {
	contract, err := bindMixinResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MixinResolverTransactor{contract: contract}, nil
}

// NewMixinResolverFilterer creates a new log filterer instance of MixinResolver, bound to a specific deployed contract.
func NewMixinResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*MixinResolverFilterer, error) {
	contract, err := bindMixinResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MixinResolverFilterer{contract: contract}, nil
}

// bindMixinResolver binds a generic wrapper to an already deployed contract.
func bindMixinResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MixinResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MixinResolver *MixinResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MixinResolver.Contract.MixinResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MixinResolver *MixinResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MixinResolver.Contract.MixinResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MixinResolver *MixinResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MixinResolver.Contract.MixinResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MixinResolver *MixinResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MixinResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MixinResolver *MixinResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MixinResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MixinResolver *MixinResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MixinResolver.Contract.contract.Transact(opts, method, params...)
}

// MAXADDRESSESFROMRESOLVER is a free data retrieval call binding the contract method 0xe3235c91.
//
// Solidity: function MAX_ADDRESSES_FROM_RESOLVER() view returns(uint256)
func (_MixinResolver *MixinResolverCaller) MAXADDRESSESFROMRESOLVER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MixinResolver.contract.Call(opts, &out, "MAX_ADDRESSES_FROM_RESOLVER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXADDRESSESFROMRESOLVER is a free data retrieval call binding the contract method 0xe3235c91.
//
// Solidity: function MAX_ADDRESSES_FROM_RESOLVER() view returns(uint256)
func (_MixinResolver *MixinResolverSession) MAXADDRESSESFROMRESOLVER() (*big.Int, error) {
	return _MixinResolver.Contract.MAXADDRESSESFROMRESOLVER(&_MixinResolver.CallOpts)
}

// MAXADDRESSESFROMRESOLVER is a free data retrieval call binding the contract method 0xe3235c91.
//
// Solidity: function MAX_ADDRESSES_FROM_RESOLVER() view returns(uint256)
func (_MixinResolver *MixinResolverCallerSession) MAXADDRESSESFROMRESOLVER() (*big.Int, error) {
	return _MixinResolver.Contract.MAXADDRESSESFROMRESOLVER(&_MixinResolver.CallOpts)
}

// GetResolverAddressesRequired is a free data retrieval call binding the contract method 0xab49848c.
//
// Solidity: function getResolverAddressesRequired() view returns(bytes32[24] addressesRequired)
func (_MixinResolver *MixinResolverCaller) GetResolverAddressesRequired(opts *bind.CallOpts) ([24][32]byte, error) {
	var out []interface{}
	err := _MixinResolver.contract.Call(opts, &out, "getResolverAddressesRequired")

	if err != nil {
		return *new([24][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([24][32]byte)).(*[24][32]byte)

	return out0, err

}

// GetResolverAddressesRequired is a free data retrieval call binding the contract method 0xab49848c.
//
// Solidity: function getResolverAddressesRequired() view returns(bytes32[24] addressesRequired)
func (_MixinResolver *MixinResolverSession) GetResolverAddressesRequired() ([24][32]byte, error) {
	return _MixinResolver.Contract.GetResolverAddressesRequired(&_MixinResolver.CallOpts)
}

// GetResolverAddressesRequired is a free data retrieval call binding the contract method 0xab49848c.
//
// Solidity: function getResolverAddressesRequired() view returns(bytes32[24] addressesRequired)
func (_MixinResolver *MixinResolverCallerSession) GetResolverAddressesRequired() ([24][32]byte, error) {
	return _MixinResolver.Contract.GetResolverAddressesRequired(&_MixinResolver.CallOpts)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x631e1444.
//
// Solidity: function isResolverCached(address _resolver) view returns(bool)
func (_MixinResolver *MixinResolverCaller) IsResolverCached(opts *bind.CallOpts, _resolver common.Address) (bool, error) {
	var out []interface{}
	err := _MixinResolver.contract.Call(opts, &out, "isResolverCached", _resolver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsResolverCached is a free data retrieval call binding the contract method 0x631e1444.
//
// Solidity: function isResolverCached(address _resolver) view returns(bool)
func (_MixinResolver *MixinResolverSession) IsResolverCached(_resolver common.Address) (bool, error) {
	return _MixinResolver.Contract.IsResolverCached(&_MixinResolver.CallOpts, _resolver)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x631e1444.
//
// Solidity: function isResolverCached(address _resolver) view returns(bool)
func (_MixinResolver *MixinResolverCallerSession) IsResolverCached(_resolver common.Address) (bool, error) {
	return _MixinResolver.Contract.IsResolverCached(&_MixinResolver.CallOpts, _resolver)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_MixinResolver *MixinResolverCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MixinResolver.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_MixinResolver *MixinResolverSession) NominatedOwner() (common.Address, error) {
	return _MixinResolver.Contract.NominatedOwner(&_MixinResolver.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_MixinResolver *MixinResolverCallerSession) NominatedOwner() (common.Address, error) {
	return _MixinResolver.Contract.NominatedOwner(&_MixinResolver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MixinResolver *MixinResolverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MixinResolver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MixinResolver *MixinResolverSession) Owner() (common.Address, error) {
	return _MixinResolver.Contract.Owner(&_MixinResolver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MixinResolver *MixinResolverCallerSession) Owner() (common.Address, error) {
	return _MixinResolver.Contract.Owner(&_MixinResolver.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_MixinResolver *MixinResolverCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MixinResolver.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_MixinResolver *MixinResolverSession) Resolver() (common.Address, error) {
	return _MixinResolver.Contract.Resolver(&_MixinResolver.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_MixinResolver *MixinResolverCallerSession) Resolver() (common.Address, error) {
	return _MixinResolver.Contract.Resolver(&_MixinResolver.CallOpts)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0xc6c9d828.
//
// Solidity: function resolverAddressesRequired(uint256 ) view returns(bytes32)
func (_MixinResolver *MixinResolverCaller) ResolverAddressesRequired(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MixinResolver.contract.Call(opts, &out, "resolverAddressesRequired", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0xc6c9d828.
//
// Solidity: function resolverAddressesRequired(uint256 ) view returns(bytes32)
func (_MixinResolver *MixinResolverSession) ResolverAddressesRequired(arg0 *big.Int) ([32]byte, error) {
	return _MixinResolver.Contract.ResolverAddressesRequired(&_MixinResolver.CallOpts, arg0)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0xc6c9d828.
//
// Solidity: function resolverAddressesRequired(uint256 ) view returns(bytes32)
func (_MixinResolver *MixinResolverCallerSession) ResolverAddressesRequired(arg0 *big.Int) ([32]byte, error) {
	return _MixinResolver.Contract.ResolverAddressesRequired(&_MixinResolver.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MixinResolver *MixinResolverTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MixinResolver.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MixinResolver *MixinResolverSession) AcceptOwnership() (*types.Transaction, error) {
	return _MixinResolver.Contract.AcceptOwnership(&_MixinResolver.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MixinResolver *MixinResolverTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MixinResolver.Contract.AcceptOwnership(&_MixinResolver.TransactOpts)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_MixinResolver *MixinResolverTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _MixinResolver.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_MixinResolver *MixinResolverSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _MixinResolver.Contract.NominateNewOwner(&_MixinResolver.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_MixinResolver *MixinResolverTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _MixinResolver.Contract.NominateNewOwner(&_MixinResolver.TransactOpts, _owner)
}

// SetResolverAndSyncCache is a paid mutator transaction binding the contract method 0x3be99e6f.
//
// Solidity: function setResolverAndSyncCache(address _resolver) returns()
func (_MixinResolver *MixinResolverTransactor) SetResolverAndSyncCache(opts *bind.TransactOpts, _resolver common.Address) (*types.Transaction, error) {
	return _MixinResolver.contract.Transact(opts, "setResolverAndSyncCache", _resolver)
}

// SetResolverAndSyncCache is a paid mutator transaction binding the contract method 0x3be99e6f.
//
// Solidity: function setResolverAndSyncCache(address _resolver) returns()
func (_MixinResolver *MixinResolverSession) SetResolverAndSyncCache(_resolver common.Address) (*types.Transaction, error) {
	return _MixinResolver.Contract.SetResolverAndSyncCache(&_MixinResolver.TransactOpts, _resolver)
}

// SetResolverAndSyncCache is a paid mutator transaction binding the contract method 0x3be99e6f.
//
// Solidity: function setResolverAndSyncCache(address _resolver) returns()
func (_MixinResolver *MixinResolverTransactorSession) SetResolverAndSyncCache(_resolver common.Address) (*types.Transaction, error) {
	return _MixinResolver.Contract.SetResolverAndSyncCache(&_MixinResolver.TransactOpts, _resolver)
}

// MixinResolverOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the MixinResolver contract.
type MixinResolverOwnerChangedIterator struct {
	Event *MixinResolverOwnerChanged // Event containing the contract specifics and raw log

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
func (it *MixinResolverOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MixinResolverOwnerChanged)
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
		it.Event = new(MixinResolverOwnerChanged)
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
func (it *MixinResolverOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MixinResolverOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MixinResolverOwnerChanged represents a OwnerChanged event raised by the MixinResolver contract.
type MixinResolverOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_MixinResolver *MixinResolverFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*MixinResolverOwnerChangedIterator, error) {

	logs, sub, err := _MixinResolver.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &MixinResolverOwnerChangedIterator{contract: _MixinResolver.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_MixinResolver *MixinResolverFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *MixinResolverOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _MixinResolver.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MixinResolverOwnerChanged)
				if err := _MixinResolver.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_MixinResolver *MixinResolverFilterer) ParseOwnerChanged(log types.Log) (*MixinResolverOwnerChanged, error) {
	event := new(MixinResolverOwnerChanged)
	if err := _MixinResolver.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MixinResolverOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the MixinResolver contract.
type MixinResolverOwnerNominatedIterator struct {
	Event *MixinResolverOwnerNominated // Event containing the contract specifics and raw log

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
func (it *MixinResolverOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MixinResolverOwnerNominated)
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
		it.Event = new(MixinResolverOwnerNominated)
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
func (it *MixinResolverOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MixinResolverOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MixinResolverOwnerNominated represents a OwnerNominated event raised by the MixinResolver contract.
type MixinResolverOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_MixinResolver *MixinResolverFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*MixinResolverOwnerNominatedIterator, error) {

	logs, sub, err := _MixinResolver.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &MixinResolverOwnerNominatedIterator{contract: _MixinResolver.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_MixinResolver *MixinResolverFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *MixinResolverOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _MixinResolver.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MixinResolverOwnerNominated)
				if err := _MixinResolver.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_MixinResolver *MixinResolverFilterer) ParseOwnerNominated(log types.Log) (*MixinResolverOwnerNominated, error) {
	event := new(MixinResolverOwnerNominated)
	if err := _MixinResolver.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OwnedFuncSigs maps the 4-byte function signature to its string representation.
var OwnedFuncSigs = map[string]string{
	"79ba5097": "acceptOwnership()",
	"1627540c": "nominateNewOwner(address)",
	"53a47bb7": "nominatedOwner()",
	"8da5cb5b": "owner()",
}

// OwnedBin is the compiled bytecode used for deploying new contracts.
var OwnedBin = "0x608060405234801561001057600080fd5b506040516103bf3803806103bf8339818101604052602081101561003357600080fd5b50516001600160a01b038116610090576040805162461bcd60e51b815260206004820152601960248201527f4f776e657220616464726573732063616e6e6f74206265203000000000000000604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b038316908117825560408051928352602083019190915280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a1506102c7806100f86000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80631627540c1461005157806353a47bb71461007957806379ba50971461009d5780638da5cb5b146100a5575b600080fd5b6100776004803603602081101561006757600080fd5b50356001600160a01b03166100ad565b005b610081610109565b604080516001600160a01b039092168252519081900360200190f35b610077610118565b6100816101d4565b6100b56101e3565b600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce229181900360200190a150565b6001546001600160a01b031681565b6001546001600160a01b031633146101615760405162461bcd60e51b815260040180806020018281038252603581526020018061022f6035913960400191505060405180910390fd5b600054600154604080516001600160a01b03938416815292909116602083015280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a160018054600080546001600160a01b03199081166001600160a01b03841617909155169055565b6000546001600160a01b031681565b6000546001600160a01b0316331461022c5760405162461bcd60e51b815260040180806020018281038252602f815260200180610264602f913960400191505060405180910390fd5b56fe596f75206d757374206265206e6f6d696e61746564206265666f726520796f752063616e20616363657074206f776e6572736869704f6e6c792074686520636f6e7472616374206f776e6572206d617920706572666f726d207468697320616374696f6ea265627a7a72315820046d928bac20775bdc1a3a2d4190c7da20cff79c612f4e8c1eb7a24ad69948d964736f6c63430005100032"

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
	OwnedFilterer   // Log filterer for contract events
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// NewOwnedFilterer creates a new log filterer instance of Owned, bound to a specific deployed contract.
func NewOwnedFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnedFilterer, error) {
	contract, err := bindOwned(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnedFilterer{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Owned *OwnedCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Owned.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Owned *OwnedSession) NominatedOwner() (common.Address, error) {
	return _Owned.Contract.NominatedOwner(&_Owned.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Owned *OwnedCallerSession) NominatedOwner() (common.Address, error) {
	return _Owned.Contract.NominatedOwner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Owned *OwnedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Owned.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Owned *OwnedSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Owned *OwnedCallerSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedSession) AcceptOwnership() (*types.Transaction, error) {
	return _Owned.Contract.AcceptOwnership(&_Owned.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Owned.Contract.AcceptOwnership(&_Owned.TransactOpts)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Owned *OwnedTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Owned *OwnedSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.NominateNewOwner(&_Owned.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Owned *OwnedTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.NominateNewOwner(&_Owned.TransactOpts, _owner)
}

// OwnedOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Owned contract.
type OwnedOwnerChangedIterator struct {
	Event *OwnedOwnerChanged // Event containing the contract specifics and raw log

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
func (it *OwnedOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnedOwnerChanged)
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
		it.Event = new(OwnedOwnerChanged)
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
func (it *OwnedOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnedOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnedOwnerChanged represents a OwnerChanged event raised by the Owned contract.
type OwnedOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Owned *OwnedFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*OwnedOwnerChangedIterator, error) {

	logs, sub, err := _Owned.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &OwnedOwnerChangedIterator{contract: _Owned.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Owned *OwnedFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *OwnedOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Owned.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnedOwnerChanged)
				if err := _Owned.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_Owned *OwnedFilterer) ParseOwnerChanged(log types.Log) (*OwnedOwnerChanged, error) {
	event := new(OwnedOwnerChanged)
	if err := _Owned.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnedOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the Owned contract.
type OwnedOwnerNominatedIterator struct {
	Event *OwnedOwnerNominated // Event containing the contract specifics and raw log

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
func (it *OwnedOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnedOwnerNominated)
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
		it.Event = new(OwnedOwnerNominated)
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
func (it *OwnedOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnedOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnedOwnerNominated represents a OwnerNominated event raised by the Owned contract.
type OwnedOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Owned *OwnedFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*OwnedOwnerNominatedIterator, error) {

	logs, sub, err := _Owned.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &OwnedOwnerNominatedIterator{contract: _Owned.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Owned *OwnedFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *OwnedOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _Owned.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnedOwnerNominated)
				if err := _Owned.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_Owned *OwnedFilterer) ParseOwnerNominated(log types.Log) (*OwnedOwnerNominated, error) {
	event := new(OwnedOwnerNominated)
	if err := _Owned.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyABI is the input ABI used to generate the binding from.
const ProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractProxyable\",\"name\":\"newTarget\",\"type\":\"address\"}],\"name\":\"TargetUpdated\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"numTopics\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic3\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic4\",\"type\":\"bytes32\"}],\"name\":\"_emit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractProxyable\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"setTarget\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"target\",\"outputs\":[{\"internalType\":\"contractProxyable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ProxyFuncSigs maps the 4-byte function signature to its string representation.
var ProxyFuncSigs = map[string]string{
	"907dff97": "_emit(bytes,uint256,bytes32,bytes32,bytes32,bytes32)",
	"79ba5097": "acceptOwnership()",
	"1627540c": "nominateNewOwner(address)",
	"53a47bb7": "nominatedOwner()",
	"8da5cb5b": "owner()",
	"776d1a01": "setTarget(address)",
	"d4b83992": "target()",
}

// ProxyBin is the compiled bytecode used for deploying new contracts.
var ProxyBin = "0x608060405234801561001057600080fd5b506040516106ee3803806106ee8339818101604052602081101561003357600080fd5b5051806001600160a01b038116610091576040805162461bcd60e51b815260206004820152601960248201527f4f776e657220616464726573732063616e6e6f74206265203000000000000000604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b038316908117825560408051928352602083019190915280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a150506105f4806100fa6000396000f3fe6080604052600436106100705760003560e01c806379ba50971161004e57806379ba5097146101925780638da5cb5b146101a7578063907dff97146101bc578063d4b839921461025157610070565b80631627540c146100f957806353a47bb71461012e578063776d1a011461015f575b60025460408051635e33fc1960e11b815233600482015290516001600160a01b039092169163bc67f8329160248082019260009290919082900301818387803b1580156100bc57600080fd5b505af11580156100d0573d6000803e3d6000fd5b5050505060405136600082376000803683346002545af13d6000833e806100f5573d82fd5b3d82f35b34801561010557600080fd5b5061012c6004803603602081101561011c57600080fd5b50356001600160a01b0316610266565b005b34801561013a57600080fd5b506101436102c2565b604080516001600160a01b039092168252519081900360200190f35b34801561016b57600080fd5b5061012c6004803603602081101561018257600080fd5b50356001600160a01b03166102d1565b34801561019e57600080fd5b5061012c61032d565b3480156101b357600080fd5b506101436103e9565b3480156101c857600080fd5b5061012c600480360360c08110156101df57600080fd5b8101906020810181356401000000008111156101fa57600080fd5b82018360208201111561020c57600080fd5b8035906020019184600183028401116401000000008311171561022e57600080fd5b9193509150803590602081013590604081013590606081013590608001356103f8565b34801561025d57600080fd5b50610143610501565b61026e610510565b600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce229181900360200190a150565b6001546001600160a01b031681565b6102d9610510565b600280546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f814250a3b8c79fcbe2ead2c131c952a278491c8f4322a79fe84b5040a810373e9181900360200190a150565b6001546001600160a01b031633146103765760405162461bcd60e51b815260040180806020018281038252603581526020018061055c6035913960400191505060405180910390fd5b600054600154604080516001600160a01b03938416815292909116602083015280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a160018054600080546001600160a01b03199081166001600160a01b03841617909155169055565b6000546001600160a01b031681565b6002546001600160a01b0316331461044e576040805162461bcd60e51b8152602060048201526014602482015273135d5cdd081899481c1c9bde1e481d185c99d95d60621b604482015290519081900360640190fd5b604080516020601f89018190048102820181019092528781528791606091908a9084908190840183828082843760009201919091525092935089925050811590506104b857600181146104c357600281146104cf57600381146104dc57600481146104ea576104f5565b8260208301a06104f5565b868360208401a16104f5565b85878460208501a26104f5565b8486888560208601a36104f5565b838587898660208701a45b50505050505050505050565b6002546001600160a01b031681565b6000546001600160a01b031633146105595760405162461bcd60e51b815260040180806020018281038252602f815260200180610591602f913960400191505060405180910390fd5b56fe596f75206d757374206265206e6f6d696e61746564206265666f726520796f752063616e20616363657074206f776e6572736869704f6e6c792074686520636f6e7472616374206f776e6572206d617920706572666f726d207468697320616374696f6ea265627a7a72315820e500016a32852e780a1f7ec9b10f45824593d02000ed6e4e2067cbbd0ea0a09a64736f6c63430005100032"

// DeployProxy deploys a new Ethereum contract, binding an instance of Proxy to it.
func DeployProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *Proxy, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Proxy *ProxyCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Proxy *ProxySession) NominatedOwner() (common.Address, error) {
	return _Proxy.Contract.NominatedOwner(&_Proxy.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Proxy *ProxyCallerSession) NominatedOwner() (common.Address, error) {
	return _Proxy.Contract.NominatedOwner(&_Proxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxy *ProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxy *ProxySession) Owner() (common.Address, error) {
	return _Proxy.Contract.Owner(&_Proxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxy *ProxyCallerSession) Owner() (common.Address, error) {
	return _Proxy.Contract.Owner(&_Proxy.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_Proxy *ProxyCaller) Target(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "target")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_Proxy *ProxySession) Target() (common.Address, error) {
	return _Proxy.Contract.Target(&_Proxy.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_Proxy *ProxyCallerSession) Target() (common.Address, error) {
	return _Proxy.Contract.Target(&_Proxy.CallOpts)
}

// Emit is a paid mutator transaction binding the contract method 0x907dff97.
//
// Solidity: function _emit(bytes callData, uint256 numTopics, bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4) returns()
func (_Proxy *ProxyTransactor) Emit(opts *bind.TransactOpts, callData []byte, numTopics *big.Int, topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "_emit", callData, numTopics, topic1, topic2, topic3, topic4)
}

// Emit is a paid mutator transaction binding the contract method 0x907dff97.
//
// Solidity: function _emit(bytes callData, uint256 numTopics, bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4) returns()
func (_Proxy *ProxySession) Emit(callData []byte, numTopics *big.Int, topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte) (*types.Transaction, error) {
	return _Proxy.Contract.Emit(&_Proxy.TransactOpts, callData, numTopics, topic1, topic2, topic3, topic4)
}

// Emit is a paid mutator transaction binding the contract method 0x907dff97.
//
// Solidity: function _emit(bytes callData, uint256 numTopics, bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4) returns()
func (_Proxy *ProxyTransactorSession) Emit(callData []byte, numTopics *big.Int, topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte) (*types.Transaction, error) {
	return _Proxy.Contract.Emit(&_Proxy.TransactOpts, callData, numTopics, topic1, topic2, topic3, topic4)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Proxy *ProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Proxy *ProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _Proxy.Contract.AcceptOwnership(&_Proxy.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Proxy *ProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Proxy.Contract.AcceptOwnership(&_Proxy.TransactOpts)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Proxy *ProxyTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Proxy *ProxySession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.NominateNewOwner(&_Proxy.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Proxy *ProxyTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.NominateNewOwner(&_Proxy.TransactOpts, _owner)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address _target) returns()
func (_Proxy *ProxyTransactor) SetTarget(opts *bind.TransactOpts, _target common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setTarget", _target)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address _target) returns()
func (_Proxy *ProxySession) SetTarget(_target common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetTarget(&_Proxy.TransactOpts, _target)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address _target) returns()
func (_Proxy *ProxyTransactorSession) SetTarget(_target common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetTarget(&_Proxy.TransactOpts, _target)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Proxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

// ProxyOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Proxy contract.
type ProxyOwnerChangedIterator struct {
	Event *ProxyOwnerChanged // Event containing the contract specifics and raw log

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
func (it *ProxyOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyOwnerChanged)
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
		it.Event = new(ProxyOwnerChanged)
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
func (it *ProxyOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyOwnerChanged represents a OwnerChanged event raised by the Proxy contract.
type ProxyOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Proxy *ProxyFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*ProxyOwnerChangedIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &ProxyOwnerChangedIterator{contract: _Proxy.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Proxy *ProxyFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *ProxyOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyOwnerChanged)
				if err := _Proxy.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_Proxy *ProxyFilterer) ParseOwnerChanged(log types.Log) (*ProxyOwnerChanged, error) {
	event := new(ProxyOwnerChanged)
	if err := _Proxy.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the Proxy contract.
type ProxyOwnerNominatedIterator struct {
	Event *ProxyOwnerNominated // Event containing the contract specifics and raw log

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
func (it *ProxyOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyOwnerNominated)
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
		it.Event = new(ProxyOwnerNominated)
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
func (it *ProxyOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyOwnerNominated represents a OwnerNominated event raised by the Proxy contract.
type ProxyOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Proxy *ProxyFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*ProxyOwnerNominatedIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &ProxyOwnerNominatedIterator{contract: _Proxy.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Proxy *ProxyFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *ProxyOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyOwnerNominated)
				if err := _Proxy.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_Proxy *ProxyFilterer) ParseOwnerNominated(log types.Log) (*ProxyOwnerNominated, error) {
	event := new(ProxyOwnerNominated)
	if err := _Proxy.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyTargetUpdatedIterator is returned from FilterTargetUpdated and is used to iterate over the raw logs and unpacked data for TargetUpdated events raised by the Proxy contract.
type ProxyTargetUpdatedIterator struct {
	Event *ProxyTargetUpdated // Event containing the contract specifics and raw log

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
func (it *ProxyTargetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyTargetUpdated)
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
		it.Event = new(ProxyTargetUpdated)
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
func (it *ProxyTargetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyTargetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyTargetUpdated represents a TargetUpdated event raised by the Proxy contract.
type ProxyTargetUpdated struct {
	NewTarget common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTargetUpdated is a free log retrieval operation binding the contract event 0x814250a3b8c79fcbe2ead2c131c952a278491c8f4322a79fe84b5040a810373e.
//
// Solidity: event TargetUpdated(address newTarget)
func (_Proxy *ProxyFilterer) FilterTargetUpdated(opts *bind.FilterOpts) (*ProxyTargetUpdatedIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "TargetUpdated")
	if err != nil {
		return nil, err
	}
	return &ProxyTargetUpdatedIterator{contract: _Proxy.contract, event: "TargetUpdated", logs: logs, sub: sub}, nil
}

// WatchTargetUpdated is a free log subscription operation binding the contract event 0x814250a3b8c79fcbe2ead2c131c952a278491c8f4322a79fe84b5040a810373e.
//
// Solidity: event TargetUpdated(address newTarget)
func (_Proxy *ProxyFilterer) WatchTargetUpdated(opts *bind.WatchOpts, sink chan<- *ProxyTargetUpdated) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "TargetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyTargetUpdated)
				if err := _Proxy.contract.UnpackLog(event, "TargetUpdated", log); err != nil {
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

// ParseTargetUpdated is a log parse operation binding the contract event 0x814250a3b8c79fcbe2ead2c131c952a278491c8f4322a79fe84b5040a810373e.
//
// Solidity: event TargetUpdated(address newTarget)
func (_Proxy *ProxyFilterer) ParseTargetUpdated(log types.Log) (*ProxyTargetUpdated, error) {
	event := new(ProxyTargetUpdated)
	if err := _Proxy.contract.UnpackLog(event, "TargetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyableABI is the input ABI used to generate the binding from.
const ProxyableABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"}],\"name\":\"ProxyUpdated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"integrationProxy\",\"outputs\":[{\"internalType\":\"contractProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"messageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"contractProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_integrationProxy\",\"type\":\"address\"}],\"name\":\"setIntegrationProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"setMessageSender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"setProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ProxyableFuncSigs maps the 4-byte function signature to its string representation.
var ProxyableFuncSigs = map[string]string{
	"79ba5097": "acceptOwnership()",
	"9cbdaeb6": "integrationProxy()",
	"d67bdd25": "messageSender()",
	"1627540c": "nominateNewOwner(address)",
	"53a47bb7": "nominatedOwner()",
	"8da5cb5b": "owner()",
	"ec556889": "proxy()",
	"131b0ae7": "setIntegrationProxy(address)",
	"bc67f832": "setMessageSender(address)",
	"97107d6d": "setProxy(address)",
}

// Proxyable is an auto generated Go binding around an Ethereum contract.
type Proxyable struct {
	ProxyableCaller     // Read-only binding to the contract
	ProxyableTransactor // Write-only binding to the contract
	ProxyableFilterer   // Log filterer for contract events
}

// ProxyableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyableSession struct {
	Contract     *Proxyable        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyableCallerSession struct {
	Contract *ProxyableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ProxyableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyableTransactorSession struct {
	Contract     *ProxyableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ProxyableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyableRaw struct {
	Contract *Proxyable // Generic contract binding to access the raw methods on
}

// ProxyableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyableCallerRaw struct {
	Contract *ProxyableCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyableTransactorRaw struct {
	Contract *ProxyableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyable creates a new instance of Proxyable, bound to a specific deployed contract.
func NewProxyable(address common.Address, backend bind.ContractBackend) (*Proxyable, error) {
	contract, err := bindProxyable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxyable{ProxyableCaller: ProxyableCaller{contract: contract}, ProxyableTransactor: ProxyableTransactor{contract: contract}, ProxyableFilterer: ProxyableFilterer{contract: contract}}, nil
}

// NewProxyableCaller creates a new read-only instance of Proxyable, bound to a specific deployed contract.
func NewProxyableCaller(address common.Address, caller bind.ContractCaller) (*ProxyableCaller, error) {
	contract, err := bindProxyable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyableCaller{contract: contract}, nil
}

// NewProxyableTransactor creates a new write-only instance of Proxyable, bound to a specific deployed contract.
func NewProxyableTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyableTransactor, error) {
	contract, err := bindProxyable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyableTransactor{contract: contract}, nil
}

// NewProxyableFilterer creates a new log filterer instance of Proxyable, bound to a specific deployed contract.
func NewProxyableFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyableFilterer, error) {
	contract, err := bindProxyable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyableFilterer{contract: contract}, nil
}

// bindProxyable binds a generic wrapper to an already deployed contract.
func bindProxyable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxyable *ProxyableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxyable.Contract.ProxyableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxyable *ProxyableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxyable.Contract.ProxyableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxyable *ProxyableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxyable.Contract.ProxyableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxyable *ProxyableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxyable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxyable *ProxyableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxyable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxyable *ProxyableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxyable.Contract.contract.Transact(opts, method, params...)
}

// IntegrationProxy is a free data retrieval call binding the contract method 0x9cbdaeb6.
//
// Solidity: function integrationProxy() view returns(address)
func (_Proxyable *ProxyableCaller) IntegrationProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxyable.contract.Call(opts, &out, "integrationProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IntegrationProxy is a free data retrieval call binding the contract method 0x9cbdaeb6.
//
// Solidity: function integrationProxy() view returns(address)
func (_Proxyable *ProxyableSession) IntegrationProxy() (common.Address, error) {
	return _Proxyable.Contract.IntegrationProxy(&_Proxyable.CallOpts)
}

// IntegrationProxy is a free data retrieval call binding the contract method 0x9cbdaeb6.
//
// Solidity: function integrationProxy() view returns(address)
func (_Proxyable *ProxyableCallerSession) IntegrationProxy() (common.Address, error) {
	return _Proxyable.Contract.IntegrationProxy(&_Proxyable.CallOpts)
}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_Proxyable *ProxyableCaller) MessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxyable.contract.Call(opts, &out, "messageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_Proxyable *ProxyableSession) MessageSender() (common.Address, error) {
	return _Proxyable.Contract.MessageSender(&_Proxyable.CallOpts)
}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_Proxyable *ProxyableCallerSession) MessageSender() (common.Address, error) {
	return _Proxyable.Contract.MessageSender(&_Proxyable.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Proxyable *ProxyableCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxyable.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Proxyable *ProxyableSession) NominatedOwner() (common.Address, error) {
	return _Proxyable.Contract.NominatedOwner(&_Proxyable.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Proxyable *ProxyableCallerSession) NominatedOwner() (common.Address, error) {
	return _Proxyable.Contract.NominatedOwner(&_Proxyable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxyable *ProxyableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxyable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxyable *ProxyableSession) Owner() (common.Address, error) {
	return _Proxyable.Contract.Owner(&_Proxyable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxyable *ProxyableCallerSession) Owner() (common.Address, error) {
	return _Proxyable.Contract.Owner(&_Proxyable.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_Proxyable *ProxyableCaller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxyable.contract.Call(opts, &out, "proxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_Proxyable *ProxyableSession) Proxy() (common.Address, error) {
	return _Proxyable.Contract.Proxy(&_Proxyable.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_Proxyable *ProxyableCallerSession) Proxy() (common.Address, error) {
	return _Proxyable.Contract.Proxy(&_Proxyable.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Proxyable *ProxyableTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxyable.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Proxyable *ProxyableSession) AcceptOwnership() (*types.Transaction, error) {
	return _Proxyable.Contract.AcceptOwnership(&_Proxyable.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Proxyable *ProxyableTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Proxyable.Contract.AcceptOwnership(&_Proxyable.TransactOpts)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Proxyable *ProxyableTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Proxyable.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Proxyable *ProxyableSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.NominateNewOwner(&_Proxyable.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Proxyable *ProxyableTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.NominateNewOwner(&_Proxyable.TransactOpts, _owner)
}

// SetIntegrationProxy is a paid mutator transaction binding the contract method 0x131b0ae7.
//
// Solidity: function setIntegrationProxy(address _integrationProxy) returns()
func (_Proxyable *ProxyableTransactor) SetIntegrationProxy(opts *bind.TransactOpts, _integrationProxy common.Address) (*types.Transaction, error) {
	return _Proxyable.contract.Transact(opts, "setIntegrationProxy", _integrationProxy)
}

// SetIntegrationProxy is a paid mutator transaction binding the contract method 0x131b0ae7.
//
// Solidity: function setIntegrationProxy(address _integrationProxy) returns()
func (_Proxyable *ProxyableSession) SetIntegrationProxy(_integrationProxy common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.SetIntegrationProxy(&_Proxyable.TransactOpts, _integrationProxy)
}

// SetIntegrationProxy is a paid mutator transaction binding the contract method 0x131b0ae7.
//
// Solidity: function setIntegrationProxy(address _integrationProxy) returns()
func (_Proxyable *ProxyableTransactorSession) SetIntegrationProxy(_integrationProxy common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.SetIntegrationProxy(&_Proxyable.TransactOpts, _integrationProxy)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_Proxyable *ProxyableTransactor) SetMessageSender(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _Proxyable.contract.Transact(opts, "setMessageSender", sender)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_Proxyable *ProxyableSession) SetMessageSender(sender common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.SetMessageSender(&_Proxyable.TransactOpts, sender)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_Proxyable *ProxyableTransactorSession) SetMessageSender(sender common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.SetMessageSender(&_Proxyable.TransactOpts, sender)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_Proxyable *ProxyableTransactor) SetProxy(opts *bind.TransactOpts, _proxy common.Address) (*types.Transaction, error) {
	return _Proxyable.contract.Transact(opts, "setProxy", _proxy)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_Proxyable *ProxyableSession) SetProxy(_proxy common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.SetProxy(&_Proxyable.TransactOpts, _proxy)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_Proxyable *ProxyableTransactorSession) SetProxy(_proxy common.Address) (*types.Transaction, error) {
	return _Proxyable.Contract.SetProxy(&_Proxyable.TransactOpts, _proxy)
}

// ProxyableOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Proxyable contract.
type ProxyableOwnerChangedIterator struct {
	Event *ProxyableOwnerChanged // Event containing the contract specifics and raw log

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
func (it *ProxyableOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyableOwnerChanged)
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
		it.Event = new(ProxyableOwnerChanged)
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
func (it *ProxyableOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyableOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyableOwnerChanged represents a OwnerChanged event raised by the Proxyable contract.
type ProxyableOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Proxyable *ProxyableFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*ProxyableOwnerChangedIterator, error) {

	logs, sub, err := _Proxyable.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &ProxyableOwnerChangedIterator{contract: _Proxyable.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Proxyable *ProxyableFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *ProxyableOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Proxyable.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyableOwnerChanged)
				if err := _Proxyable.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_Proxyable *ProxyableFilterer) ParseOwnerChanged(log types.Log) (*ProxyableOwnerChanged, error) {
	event := new(ProxyableOwnerChanged)
	if err := _Proxyable.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyableOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the Proxyable contract.
type ProxyableOwnerNominatedIterator struct {
	Event *ProxyableOwnerNominated // Event containing the contract specifics and raw log

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
func (it *ProxyableOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyableOwnerNominated)
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
		it.Event = new(ProxyableOwnerNominated)
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
func (it *ProxyableOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyableOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyableOwnerNominated represents a OwnerNominated event raised by the Proxyable contract.
type ProxyableOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Proxyable *ProxyableFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*ProxyableOwnerNominatedIterator, error) {

	logs, sub, err := _Proxyable.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &ProxyableOwnerNominatedIterator{contract: _Proxyable.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Proxyable *ProxyableFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *ProxyableOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _Proxyable.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyableOwnerNominated)
				if err := _Proxyable.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_Proxyable *ProxyableFilterer) ParseOwnerNominated(log types.Log) (*ProxyableOwnerNominated, error) {
	event := new(ProxyableOwnerNominated)
	if err := _Proxyable.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyableProxyUpdatedIterator is returned from FilterProxyUpdated and is used to iterate over the raw logs and unpacked data for ProxyUpdated events raised by the Proxyable contract.
type ProxyableProxyUpdatedIterator struct {
	Event *ProxyableProxyUpdated // Event containing the contract specifics and raw log

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
func (it *ProxyableProxyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyableProxyUpdated)
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
		it.Event = new(ProxyableProxyUpdated)
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
func (it *ProxyableProxyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyableProxyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyableProxyUpdated represents a ProxyUpdated event raised by the Proxyable contract.
type ProxyableProxyUpdated struct {
	ProxyAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProxyUpdated is a free log retrieval operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_Proxyable *ProxyableFilterer) FilterProxyUpdated(opts *bind.FilterOpts) (*ProxyableProxyUpdatedIterator, error) {

	logs, sub, err := _Proxyable.contract.FilterLogs(opts, "ProxyUpdated")
	if err != nil {
		return nil, err
	}
	return &ProxyableProxyUpdatedIterator{contract: _Proxyable.contract, event: "ProxyUpdated", logs: logs, sub: sub}, nil
}

// WatchProxyUpdated is a free log subscription operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_Proxyable *ProxyableFilterer) WatchProxyUpdated(opts *bind.WatchOpts, sink chan<- *ProxyableProxyUpdated) (event.Subscription, error) {

	logs, sub, err := _Proxyable.contract.WatchLogs(opts, "ProxyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyableProxyUpdated)
				if err := _Proxyable.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
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

// ParseProxyUpdated is a log parse operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_Proxyable *ProxyableFilterer) ParseProxyUpdated(log types.Log) (*ProxyableProxyUpdated, error) {
	event := new(ProxyableProxyUpdated)
	if err := _Proxyable.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeDecimalMathABI is the input ABI used to generate the binding from.
const SafeDecimalMathABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"PRECISE_UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"highPrecisionDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"preciseUnit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SafeDecimalMathFuncSigs maps the 4-byte function signature to its string representation.
var SafeDecimalMathFuncSigs = map[string]string{
	"864029e7": "PRECISE_UNIT()",
	"9d8e2177": "UNIT()",
	"313ce567": "decimals()",
	"def4419d": "highPrecisionDecimals()",
	"d5e5e6e6": "preciseUnit()",
	"907af6c0": "unit()",
}

// SafeDecimalMathBin is the compiled bytecode used for deploying new contracts.
var SafeDecimalMathBin = "0x61012d610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060655760003560e01c8063313ce56714606a578063864029e7146086578063907af6c014609e5780639d8e21771460a4578063d5e5e6e61460aa578063def4419d1460b0575b600080fd5b607060b6565b6040805160ff9092168252519081900360200190f35b608c60bb565b60408051918252519081900360200190f35b608c60cb565b608c60d7565b608c60e3565b607060f3565b601281565b6b033b2e3c9fd0803ce800000081565b670de0b6b3a764000090565b670de0b6b3a764000081565b6b033b2e3c9fd0803ce800000090565b601b8156fea265627a7a7231582080443880d89c5a9fa62a2dff99f097c6d9c2d83716106a802924ba9fdf0ebfdf64736f6c63430005100032"

// DeploySafeDecimalMath deploys a new Ethereum contract, binding an instance of SafeDecimalMath to it.
func DeploySafeDecimalMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeDecimalMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeDecimalMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeDecimalMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeDecimalMath{SafeDecimalMathCaller: SafeDecimalMathCaller{contract: contract}, SafeDecimalMathTransactor: SafeDecimalMathTransactor{contract: contract}, SafeDecimalMathFilterer: SafeDecimalMathFilterer{contract: contract}}, nil
}

// SafeDecimalMath is an auto generated Go binding around an Ethereum contract.
type SafeDecimalMath struct {
	SafeDecimalMathCaller     // Read-only binding to the contract
	SafeDecimalMathTransactor // Write-only binding to the contract
	SafeDecimalMathFilterer   // Log filterer for contract events
}

// SafeDecimalMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeDecimalMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeDecimalMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeDecimalMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeDecimalMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeDecimalMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeDecimalMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeDecimalMathSession struct {
	Contract     *SafeDecimalMath  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeDecimalMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeDecimalMathCallerSession struct {
	Contract *SafeDecimalMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SafeDecimalMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeDecimalMathTransactorSession struct {
	Contract     *SafeDecimalMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SafeDecimalMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeDecimalMathRaw struct {
	Contract *SafeDecimalMath // Generic contract binding to access the raw methods on
}

// SafeDecimalMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeDecimalMathCallerRaw struct {
	Contract *SafeDecimalMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeDecimalMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeDecimalMathTransactorRaw struct {
	Contract *SafeDecimalMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeDecimalMath creates a new instance of SafeDecimalMath, bound to a specific deployed contract.
func NewSafeDecimalMath(address common.Address, backend bind.ContractBackend) (*SafeDecimalMath, error) {
	contract, err := bindSafeDecimalMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeDecimalMath{SafeDecimalMathCaller: SafeDecimalMathCaller{contract: contract}, SafeDecimalMathTransactor: SafeDecimalMathTransactor{contract: contract}, SafeDecimalMathFilterer: SafeDecimalMathFilterer{contract: contract}}, nil
}

// NewSafeDecimalMathCaller creates a new read-only instance of SafeDecimalMath, bound to a specific deployed contract.
func NewSafeDecimalMathCaller(address common.Address, caller bind.ContractCaller) (*SafeDecimalMathCaller, error) {
	contract, err := bindSafeDecimalMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeDecimalMathCaller{contract: contract}, nil
}

// NewSafeDecimalMathTransactor creates a new write-only instance of SafeDecimalMath, bound to a specific deployed contract.
func NewSafeDecimalMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeDecimalMathTransactor, error) {
	contract, err := bindSafeDecimalMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeDecimalMathTransactor{contract: contract}, nil
}

// NewSafeDecimalMathFilterer creates a new log filterer instance of SafeDecimalMath, bound to a specific deployed contract.
func NewSafeDecimalMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeDecimalMathFilterer, error) {
	contract, err := bindSafeDecimalMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeDecimalMathFilterer{contract: contract}, nil
}

// bindSafeDecimalMath binds a generic wrapper to an already deployed contract.
func bindSafeDecimalMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeDecimalMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeDecimalMath *SafeDecimalMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeDecimalMath.Contract.SafeDecimalMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeDecimalMath *SafeDecimalMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeDecimalMath.Contract.SafeDecimalMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeDecimalMath *SafeDecimalMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeDecimalMath.Contract.SafeDecimalMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeDecimalMath *SafeDecimalMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeDecimalMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeDecimalMath *SafeDecimalMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeDecimalMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeDecimalMath *SafeDecimalMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeDecimalMath.Contract.contract.Transact(opts, method, params...)
}

// PRECISEUNIT is a free data retrieval call binding the contract method 0x864029e7.
//
// Solidity: function PRECISE_UNIT() view returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCaller) PRECISEUNIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeDecimalMath.contract.Call(opts, &out, "PRECISE_UNIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISEUNIT is a free data retrieval call binding the contract method 0x864029e7.
//
// Solidity: function PRECISE_UNIT() view returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathSession) PRECISEUNIT() (*big.Int, error) {
	return _SafeDecimalMath.Contract.PRECISEUNIT(&_SafeDecimalMath.CallOpts)
}

// PRECISEUNIT is a free data retrieval call binding the contract method 0x864029e7.
//
// Solidity: function PRECISE_UNIT() view returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCallerSession) PRECISEUNIT() (*big.Int, error) {
	return _SafeDecimalMath.Contract.PRECISEUNIT(&_SafeDecimalMath.CallOpts)
}

// UNIT is a free data retrieval call binding the contract method 0x9d8e2177.
//
// Solidity: function UNIT() view returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCaller) UNIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeDecimalMath.contract.Call(opts, &out, "UNIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNIT is a free data retrieval call binding the contract method 0x9d8e2177.
//
// Solidity: function UNIT() view returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathSession) UNIT() (*big.Int, error) {
	return _SafeDecimalMath.Contract.UNIT(&_SafeDecimalMath.CallOpts)
}

// UNIT is a free data retrieval call binding the contract method 0x9d8e2177.
//
// Solidity: function UNIT() view returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCallerSession) UNIT() (*big.Int, error) {
	return _SafeDecimalMath.Contract.UNIT(&_SafeDecimalMath.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SafeDecimalMath *SafeDecimalMathCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SafeDecimalMath.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SafeDecimalMath *SafeDecimalMathSession) Decimals() (uint8, error) {
	return _SafeDecimalMath.Contract.Decimals(&_SafeDecimalMath.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SafeDecimalMath *SafeDecimalMathCallerSession) Decimals() (uint8, error) {
	return _SafeDecimalMath.Contract.Decimals(&_SafeDecimalMath.CallOpts)
}

// HighPrecisionDecimals is a free data retrieval call binding the contract method 0xdef4419d.
//
// Solidity: function highPrecisionDecimals() view returns(uint8)
func (_SafeDecimalMath *SafeDecimalMathCaller) HighPrecisionDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SafeDecimalMath.contract.Call(opts, &out, "highPrecisionDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// HighPrecisionDecimals is a free data retrieval call binding the contract method 0xdef4419d.
//
// Solidity: function highPrecisionDecimals() view returns(uint8)
func (_SafeDecimalMath *SafeDecimalMathSession) HighPrecisionDecimals() (uint8, error) {
	return _SafeDecimalMath.Contract.HighPrecisionDecimals(&_SafeDecimalMath.CallOpts)
}

// HighPrecisionDecimals is a free data retrieval call binding the contract method 0xdef4419d.
//
// Solidity: function highPrecisionDecimals() view returns(uint8)
func (_SafeDecimalMath *SafeDecimalMathCallerSession) HighPrecisionDecimals() (uint8, error) {
	return _SafeDecimalMath.Contract.HighPrecisionDecimals(&_SafeDecimalMath.CallOpts)
}

// PreciseUnit is a free data retrieval call binding the contract method 0xd5e5e6e6.
//
// Solidity: function preciseUnit() pure returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCaller) PreciseUnit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeDecimalMath.contract.Call(opts, &out, "preciseUnit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreciseUnit is a free data retrieval call binding the contract method 0xd5e5e6e6.
//
// Solidity: function preciseUnit() pure returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathSession) PreciseUnit() (*big.Int, error) {
	return _SafeDecimalMath.Contract.PreciseUnit(&_SafeDecimalMath.CallOpts)
}

// PreciseUnit is a free data retrieval call binding the contract method 0xd5e5e6e6.
//
// Solidity: function preciseUnit() pure returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCallerSession) PreciseUnit() (*big.Int, error) {
	return _SafeDecimalMath.Contract.PreciseUnit(&_SafeDecimalMath.CallOpts)
}

// Unit is a free data retrieval call binding the contract method 0x907af6c0.
//
// Solidity: function unit() pure returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCaller) Unit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeDecimalMath.contract.Call(opts, &out, "unit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Unit is a free data retrieval call binding the contract method 0x907af6c0.
//
// Solidity: function unit() pure returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathSession) Unit() (*big.Int, error) {
	return _SafeDecimalMath.Contract.Unit(&_SafeDecimalMath.CallOpts)
}

// Unit is a free data retrieval call binding the contract method 0x907af6c0.
//
// Solidity: function unit() pure returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCallerSession) Unit() (*big.Int, error) {
	return _SafeDecimalMath.Contract.Unit(&_SafeDecimalMath.CallOpts)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582070919146728cbca38d6259a0ef7e0a06dcdb771d61bf381ee6ec8c8a48ca31f764736f6c63430005100032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StateABI is the input ABI used to generate the binding from.
const StateABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_associatedContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"associatedContract\",\"type\":\"address\"}],\"name\":\"AssociatedContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"associatedContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_associatedContract\",\"type\":\"address\"}],\"name\":\"setAssociatedContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StateFuncSigs maps the 4-byte function signature to its string representation.
var StateFuncSigs = map[string]string{
	"79ba5097": "acceptOwnership()",
	"aefc4ccb": "associatedContract()",
	"1627540c": "nominateNewOwner(address)",
	"53a47bb7": "nominatedOwner()",
	"8da5cb5b": "owner()",
	"52f445ca": "setAssociatedContract(address)",
}

// State is an auto generated Go binding around an Ethereum contract.
type State struct {
	StateCaller     // Read-only binding to the contract
	StateTransactor // Write-only binding to the contract
	StateFilterer   // Log filterer for contract events
}

// StateCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateSession struct {
	Contract     *State            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateCallerSession struct {
	Contract *StateCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateTransactorSession struct {
	Contract     *StateTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateRaw struct {
	Contract *State // Generic contract binding to access the raw methods on
}

// StateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateCallerRaw struct {
	Contract *StateCaller // Generic read-only contract binding to access the raw methods on
}

// StateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateTransactorRaw struct {
	Contract *StateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewState creates a new instance of State, bound to a specific deployed contract.
func NewState(address common.Address, backend bind.ContractBackend) (*State, error) {
	contract, err := bindState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &State{StateCaller: StateCaller{contract: contract}, StateTransactor: StateTransactor{contract: contract}, StateFilterer: StateFilterer{contract: contract}}, nil
}

// NewStateCaller creates a new read-only instance of State, bound to a specific deployed contract.
func NewStateCaller(address common.Address, caller bind.ContractCaller) (*StateCaller, error) {
	contract, err := bindState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateCaller{contract: contract}, nil
}

// NewStateTransactor creates a new write-only instance of State, bound to a specific deployed contract.
func NewStateTransactor(address common.Address, transactor bind.ContractTransactor) (*StateTransactor, error) {
	contract, err := bindState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateTransactor{contract: contract}, nil
}

// NewStateFilterer creates a new log filterer instance of State, bound to a specific deployed contract.
func NewStateFilterer(address common.Address, filterer bind.ContractFilterer) (*StateFilterer, error) {
	contract, err := bindState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateFilterer{contract: contract}, nil
}

// bindState binds a generic wrapper to an already deployed contract.
func bindState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_State *StateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _State.Contract.StateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_State *StateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _State.Contract.StateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_State *StateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _State.Contract.StateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_State *StateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _State.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_State *StateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _State.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_State *StateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _State.Contract.contract.Transact(opts, method, params...)
}

// AssociatedContract is a free data retrieval call binding the contract method 0xaefc4ccb.
//
// Solidity: function associatedContract() view returns(address)
func (_State *StateCaller) AssociatedContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "associatedContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssociatedContract is a free data retrieval call binding the contract method 0xaefc4ccb.
//
// Solidity: function associatedContract() view returns(address)
func (_State *StateSession) AssociatedContract() (common.Address, error) {
	return _State.Contract.AssociatedContract(&_State.CallOpts)
}

// AssociatedContract is a free data retrieval call binding the contract method 0xaefc4ccb.
//
// Solidity: function associatedContract() view returns(address)
func (_State *StateCallerSession) AssociatedContract() (common.Address, error) {
	return _State.Contract.AssociatedContract(&_State.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_State *StateCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_State *StateSession) NominatedOwner() (common.Address, error) {
	return _State.Contract.NominatedOwner(&_State.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_State *StateCallerSession) NominatedOwner() (common.Address, error) {
	return _State.Contract.NominatedOwner(&_State.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_State *StateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_State *StateSession) Owner() (common.Address, error) {
	return _State.Contract.Owner(&_State.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_State *StateCallerSession) Owner() (common.Address, error) {
	return _State.Contract.Owner(&_State.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_State *StateTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_State *StateSession) AcceptOwnership() (*types.Transaction, error) {
	return _State.Contract.AcceptOwnership(&_State.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_State *StateTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _State.Contract.AcceptOwnership(&_State.TransactOpts)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_State *StateTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_State *StateSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _State.Contract.NominateNewOwner(&_State.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_State *StateTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _State.Contract.NominateNewOwner(&_State.TransactOpts, _owner)
}

// SetAssociatedContract is a paid mutator transaction binding the contract method 0x52f445ca.
//
// Solidity: function setAssociatedContract(address _associatedContract) returns()
func (_State *StateTransactor) SetAssociatedContract(opts *bind.TransactOpts, _associatedContract common.Address) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "setAssociatedContract", _associatedContract)
}

// SetAssociatedContract is a paid mutator transaction binding the contract method 0x52f445ca.
//
// Solidity: function setAssociatedContract(address _associatedContract) returns()
func (_State *StateSession) SetAssociatedContract(_associatedContract common.Address) (*types.Transaction, error) {
	return _State.Contract.SetAssociatedContract(&_State.TransactOpts, _associatedContract)
}

// SetAssociatedContract is a paid mutator transaction binding the contract method 0x52f445ca.
//
// Solidity: function setAssociatedContract(address _associatedContract) returns()
func (_State *StateTransactorSession) SetAssociatedContract(_associatedContract common.Address) (*types.Transaction, error) {
	return _State.Contract.SetAssociatedContract(&_State.TransactOpts, _associatedContract)
}

// StateAssociatedContractUpdatedIterator is returned from FilterAssociatedContractUpdated and is used to iterate over the raw logs and unpacked data for AssociatedContractUpdated events raised by the State contract.
type StateAssociatedContractUpdatedIterator struct {
	Event *StateAssociatedContractUpdated // Event containing the contract specifics and raw log

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
func (it *StateAssociatedContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateAssociatedContractUpdated)
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
		it.Event = new(StateAssociatedContractUpdated)
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
func (it *StateAssociatedContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateAssociatedContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateAssociatedContractUpdated represents a AssociatedContractUpdated event raised by the State contract.
type StateAssociatedContractUpdated struct {
	AssociatedContract common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAssociatedContractUpdated is a free log retrieval operation binding the contract event 0x73f20cff579e8a4086fa607db83867595f1b6a798e718c0bfa0b94a404128e03.
//
// Solidity: event AssociatedContractUpdated(address associatedContract)
func (_State *StateFilterer) FilterAssociatedContractUpdated(opts *bind.FilterOpts) (*StateAssociatedContractUpdatedIterator, error) {

	logs, sub, err := _State.contract.FilterLogs(opts, "AssociatedContractUpdated")
	if err != nil {
		return nil, err
	}
	return &StateAssociatedContractUpdatedIterator{contract: _State.contract, event: "AssociatedContractUpdated", logs: logs, sub: sub}, nil
}

// WatchAssociatedContractUpdated is a free log subscription operation binding the contract event 0x73f20cff579e8a4086fa607db83867595f1b6a798e718c0bfa0b94a404128e03.
//
// Solidity: event AssociatedContractUpdated(address associatedContract)
func (_State *StateFilterer) WatchAssociatedContractUpdated(opts *bind.WatchOpts, sink chan<- *StateAssociatedContractUpdated) (event.Subscription, error) {

	logs, sub, err := _State.contract.WatchLogs(opts, "AssociatedContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateAssociatedContractUpdated)
				if err := _State.contract.UnpackLog(event, "AssociatedContractUpdated", log); err != nil {
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

// ParseAssociatedContractUpdated is a log parse operation binding the contract event 0x73f20cff579e8a4086fa607db83867595f1b6a798e718c0bfa0b94a404128e03.
//
// Solidity: event AssociatedContractUpdated(address associatedContract)
func (_State *StateFilterer) ParseAssociatedContractUpdated(log types.Log) (*StateAssociatedContractUpdated, error) {
	event := new(StateAssociatedContractUpdated)
	if err := _State.contract.UnpackLog(event, "AssociatedContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the State contract.
type StateOwnerChangedIterator struct {
	Event *StateOwnerChanged // Event containing the contract specifics and raw log

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
func (it *StateOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateOwnerChanged)
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
		it.Event = new(StateOwnerChanged)
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
func (it *StateOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateOwnerChanged represents a OwnerChanged event raised by the State contract.
type StateOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_State *StateFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*StateOwnerChangedIterator, error) {

	logs, sub, err := _State.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &StateOwnerChangedIterator{contract: _State.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_State *StateFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *StateOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _State.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateOwnerChanged)
				if err := _State.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_State *StateFilterer) ParseOwnerChanged(log types.Log) (*StateOwnerChanged, error) {
	event := new(StateOwnerChanged)
	if err := _State.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the State contract.
type StateOwnerNominatedIterator struct {
	Event *StateOwnerNominated // Event containing the contract specifics and raw log

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
func (it *StateOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateOwnerNominated)
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
		it.Event = new(StateOwnerNominated)
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
func (it *StateOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateOwnerNominated represents a OwnerNominated event raised by the State contract.
type StateOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_State *StateFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*StateOwnerNominatedIterator, error) {

	logs, sub, err := _State.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &StateOwnerNominatedIterator{contract: _State.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_State *StateFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *StateOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _State.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateOwnerNominated)
				if err := _State.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_State *StateFilterer) ParseOwnerNominated(log types.Log) (*StateOwnerNominated, error) {
	event := new(StateOwnerNominated)
	if err := _State.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyScheduleABI is the input ABI used to generate the binding from.
const SupplyScheduleABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lastMintEvent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_currentWeek\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRewardAmount\",\"type\":\"uint256\"}],\"name\":\"MinterRewardUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supplyMinted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numberOfWeeksIssued\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastMintEvent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SupplyMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"SynthetixProxyUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"DECAY_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INFLATION_START_DATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_WEEKLY_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_MINTER_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINT_BUFFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINT_PERIOD_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SUPPLY_DECAY_END\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SUPPLY_DECAY_START\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TERMINAL_SUPPLY_RATE_ANNUAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastMintEvent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintableSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minterReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"supplyMinted\",\"type\":\"uint256\"}],\"name\":\"recordMintEvent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setMinterReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractISynthetix\",\"name\":\"_synthetixProxy\",\"type\":\"address\"}],\"name\":\"setSynthetixProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"synthetixProxy\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numOfWeeks\",\"type\":\"uint256\"}],\"name\":\"terminalInflationSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"name\":\"tokenDecaySupplyForWeek\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weekCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weeksSinceLastIssuance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SupplyScheduleFuncSigs maps the 4-byte function signature to its string representation.
var SupplyScheduleFuncSigs = map[string]string{
	"1de40e49": "DECAY_RATE()",
	"7e1b823f": "INFLATION_START_DATE()",
	"badef30a": "INITIAL_WEEKLY_SUPPLY()",
	"22af2bab": "MAX_MINTER_REWARD()",
	"46872a23": "MINT_BUFFER()",
	"df5a9fc1": "MINT_PERIOD_DURATION()",
	"7c060557": "SUPPLY_DECAY_END()",
	"251330f1": "SUPPLY_DECAY_START()",
	"25542064": "TERMINAL_SUPPLY_RATE_ANNUAL()",
	"79ba5097": "acceptOwnership()",
	"46b45af7": "isMintable()",
	"be801f01": "lastMintEvent()",
	"cc5c095c": "mintableSupply()",
	"9bdd7ac7": "minterReward()",
	"1627540c": "nominateNewOwner(address)",
	"53a47bb7": "nominatedOwner()",
	"8da5cb5b": "owner()",
	"7e7961d7": "recordMintEvent(uint256)",
	"4ae26521": "setMinterReward(uint256)",
	"95896b76": "setSynthetixProxy(address)",
	"bdd12482": "synthetixProxy()",
	"b3b2bcc0": "terminalInflationSupply(uint256,uint256)",
	"4e070f50": "tokenDecaySupplyForWeek(uint256)",
	"d3bd4bde": "weekCounter()",
	"dbd3a6a7": "weeksSinceLastIssuance()",
}

// SupplyScheduleBin is the compiled bytecode used for deploying new contracts.
var SupplyScheduleBin = "0x60806040819052630241ebdb60e61b815273__$631c3c01b9193af804154dcec27dcae146$__9063907af6c09060849060209060048186803b15801561004457600080fd5b505af4158015610058573d6000803e3d6000fd5b505050506040513d602081101561006e57600080fd5b505160c80260045534801561008257600080fd5b50604051611094380380611094833981810160405260608110156100a557600080fd5b5080516020820151604090920151909190826001600160a01b038116610112576040805162461bcd60e51b815260206004820152601960248201527f4f776e657220616464726573732063616e6e6f74206265203000000000000000604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b038316908117825560408051928352602083019190915280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a15060029190915560035550610f10806101846000396000f3fe608060405234801561001057600080fd5b50600436106101735760003560e01c80637e1b823f116100de578063badef30a11610097578063cc5c095c11610071578063cc5c095c14610310578063d3bd4bde14610318578063dbd3a6a714610320578063df5a9fc11461032857610173565b8063badef30a146102f8578063bdd1248214610300578063be801f011461030857610173565b80637e1b823f1461027a5780637e7961d7146102825780638da5cb5b1461029f57806395896b76146102a75780639bdd7ac7146102cd578063b3b2bcc0146102d557610173565b806346b45af71161013057806346b45af7146101f05780634ae265211461020c5780634e070f501461022957806353a47bb71461024657806379ba50971461026a5780637c0605571461027257610173565b80631627540c146101785780631de40e49146101a057806322af2bab146101ba578063251330f1146101c257806325542064146101e057806346872a23146101e8575b600080fd5b61019e6004803603602081101561018e57600080fd5b50356001600160a01b0316610330565b005b6101a861038c565b60408051918252519081900360200190f35b6101a8610397565b6101ca6103a4565b6040805160ff9092168252519081900360200190f35b6101a86103a9565b6101a86103b4565b6101f86103bb565b604080519115158252519081900360200190f35b61019e6004803603602081101561022257600080fd5b50356103db565b6101a86004803603602081101561023f57600080fd5b5035610466565b61024e610527565b604080516001600160a01b039092168252519081900360200190f35b61019e610536565b6101ca6105f2565b6101a86105f7565b6101f86004803603602081101561029857600080fd5b50356105ff565b61024e61076c565b61019e600480360360208110156102bd57600080fd5b50356001600160a01b031661077b565b6101a861082e565b6101a8600480360360408110156102eb57600080fd5b5080359060200135610834565b6101a861096e565b61024e61097d565b6101a861098c565b6101a8610992565b6101a8610af0565b6101a8610af6565b6101a8610b4b565b610338610b52565b600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce229181900360200190a150565b662c68af0bb1400081565b680ad78ebc5ac620000081565b602881565b6658d15e1762800081565b6201518081565b600062093a80600254420311156103d4575060016103d8565b5060005b90565b6103e3610b52565b680ad78ebc5ac620000081111561042b5760405162461bcd60e51b8152600401808060200182810382526026815260200180610eb66026913960400191505060405180910390fd5b60048190556040805182815290517f036e0c635f8b7d9314bb6f2a577046108ef0f8b5e3869fbd29fd5a448ed99d309181900360200190a150565b600080610500836104f4662c68af0bb1400073__$631c3c01b9193af804154dcec27dcae146$__63907af6c06040518163ffffffff1660e01b815260040160206040518083038186803b1580156104bc57600080fd5b505af41580156104d0573d6000803e3d6000fd5b505050506040513d60208110156104e657600080fd5b50519063ffffffff610b9d16565b9063ffffffff610bfa16565b9050600061051f6a01316ba81b802f59713b138363ffffffff610cb616565b949350505050565b6001546001600160a01b031681565b6001546001600160a01b0316331461057f5760405162461bcd60e51b8152600401808060200182810382526035815260200180610dfe6035913960400191505060405180910390fd5b600054600154604080516001600160a01b03938416815292909116602083015280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a160018054600080546001600160a01b03199081166001600160a01b03841617909155169055565b60ea81565b635c7f0d8081565b60055460408051636a5c1cc960e11b815290516000926001600160a01b03169163d4b83992916004808301926020929190829003018186803b15801561064457600080fd5b505afa158015610658573d6000803e3d6000fd5b505050506040513d602081101561066e57600080fd5b50516001600160a01b031633146106b65760405162461bcd60e51b8152600401808060200182810382526033815260200180610e836033913960400191505060405180910390fd5b60006106c0610af6565b6003549091506106d6908263ffffffff610ce016565b600381905561071890620151809061070c906106fb9062093a8063ffffffff610d3a16565b635c7f0d809063ffffffff610ce016565b9063ffffffff610ce016565b6002819055604080518581526020810184905280820192909252426060830152517f601e517d4811033fed8290c79b7823ce1ab70258da45400fe2391a3c7432edab9181900360800190a150600192915050565b6000546001600160a01b031681565b610783610b52565b6001600160a01b0381166107d4576040805162461bcd60e51b81526020600482015260136024820152720416464726573732063616e6e6f74206265203606c1b604482015290519081900360640190fd5b600580546001600160a01b0319166001600160a01b03838116919091179182905560408051929091168252517ff8df0556b7fde3c4b8394eae329aedfa59c6ffd8b532d572a1efcfa3424ca5fc916020908290030190a150565b60045481565b6000806108d2836104f46108566658d15e17628000603463ffffffff610d9316565b73__$631c3c01b9193af804154dcec27dcae146$__63907af6c06040518163ffffffff1660e01b815260040160206040518083038186803b15801561089a57600080fd5b505af41580156108ae573d6000803e3d6000fd5b505050506040513d60208110156108c457600080fd5b50519063ffffffff610ce016565b905061096461095773__$631c3c01b9193af804154dcec27dcae146$__63907af6c06040518163ffffffff1660e01b815260040160206040518083038186803b15801561091e57600080fd5b505af4158015610932573d6000803e3d6000fd5b505050506040513d602081101561094857600080fd5b5051839063ffffffff610b9d16565b859063ffffffff610cb616565b9150505b92915050565b6a01316ba81b802f59713b1381565b6005546001600160a01b031681565b60025481565b60008061099d6103bb565b6109a85790506103d8565b60006109b2610af6565b6003549091505b8115610ae85760010160288110156109f4576109e6836a01316ba81b802f59713b1363ffffffff610ce016565b925060001990910190610ae3565b60ea8111610a39576000610a0f82602763ffffffff610b9d16565b9050610a2a610a1d82610466565b859063ffffffff610ce016565b93505060001990910190610ae3565b600554604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015610a7e57600080fd5b505afa158015610a92573d6000803e3d6000fd5b505050506040513d6020811015610aa857600080fd5b505190506000610abe828663ffffffff610ce016565b9050610ada610acd8286610834565b869063ffffffff610ce016565b94506000935050505b6109b9565b509091505090565b60035481565b600080600060025411610b1c57610b1742635c7f0d8063ffffffff610b9d16565b610b30565b600254610b3090429063ffffffff610b9d16565b9050610b458162093a8063ffffffff610d9316565b91505090565b62093a8081565b6000546001600160a01b03163314610b9b5760405162461bcd60e51b815260040180806020018281038252602f815260200180610e33602f913960400191505060405180910390fd5b565b600082821115610bf4576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b60008073__$631c3c01b9193af804154dcec27dcae146$__63907af6c06040518163ffffffff1660e01b815260040160206040518083038186803b158015610c4157600080fd5b505af4158015610c55573d6000803e3d6000fd5b505050506040513d6020811015610c6b57600080fd5b505190505b8215610caf576002830615610c9257610c8f818563ffffffff610cb616565b90505b610ca2848063ffffffff610cb616565b9350600283049250610c70565b9392505050565b6000670de0b6b3a7640000610cd1848463ffffffff610d3a16565b81610cd857fe5b049392505050565b600082820183811015610caf576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600082610d4957506000610968565b82820282848281610d5657fe5b0414610caf5760405162461bcd60e51b8152600401808060200182810382526021815260200180610e626021913960400191505060405180910390fd5b6000808211610de9576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b6000828481610df457fe5b0494935050505056fe596f75206d757374206265206e6f6d696e61746564206265666f726520796f752063616e20616363657074206f776e6572736869704f6e6c792074686520636f6e7472616374206f776e6572206d617920706572666f726d207468697320616374696f6e536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f774f6e6c79207468652073796e74686574697820636f6e74726163742063616e20706572666f726d207468697320616374696f6e5265776172642063616e6e6f7420657863656564206d6178206d696e74657220726577617264a265627a7a72315820cfaed7b41dfd6d631646d0d102f0777097b6cb01d76472fd2126efd4c8e582f564736f6c63430005100032"

// DeploySupplySchedule deploys a new Ethereum contract, binding an instance of SupplySchedule to it.
func DeploySupplySchedule(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _lastMintEvent *big.Int, _currentWeek *big.Int) (common.Address, *types.Transaction, *SupplySchedule, error) {
	parsed, err := abi.JSON(strings.NewReader(SupplyScheduleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	safeDecimalMathAddr, _, _, _ := DeploySafeDecimalMath(auth, backend)
	SupplyScheduleBin = strings.Replace(SupplyScheduleBin, "__$631c3c01b9193af804154dcec27dcae146$__", safeDecimalMathAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SupplyScheduleBin), backend, _owner, _lastMintEvent, _currentWeek)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SupplySchedule{SupplyScheduleCaller: SupplyScheduleCaller{contract: contract}, SupplyScheduleTransactor: SupplyScheduleTransactor{contract: contract}, SupplyScheduleFilterer: SupplyScheduleFilterer{contract: contract}}, nil
}

// SupplySchedule is an auto generated Go binding around an Ethereum contract.
type SupplySchedule struct {
	SupplyScheduleCaller     // Read-only binding to the contract
	SupplyScheduleTransactor // Write-only binding to the contract
	SupplyScheduleFilterer   // Log filterer for contract events
}

// SupplyScheduleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SupplyScheduleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyScheduleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SupplyScheduleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyScheduleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SupplyScheduleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyScheduleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SupplyScheduleSession struct {
	Contract     *SupplySchedule   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SupplyScheduleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SupplyScheduleCallerSession struct {
	Contract *SupplyScheduleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SupplyScheduleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SupplyScheduleTransactorSession struct {
	Contract     *SupplyScheduleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SupplyScheduleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SupplyScheduleRaw struct {
	Contract *SupplySchedule // Generic contract binding to access the raw methods on
}

// SupplyScheduleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SupplyScheduleCallerRaw struct {
	Contract *SupplyScheduleCaller // Generic read-only contract binding to access the raw methods on
}

// SupplyScheduleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SupplyScheduleTransactorRaw struct {
	Contract *SupplyScheduleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSupplySchedule creates a new instance of SupplySchedule, bound to a specific deployed contract.
func NewSupplySchedule(address common.Address, backend bind.ContractBackend) (*SupplySchedule, error) {
	contract, err := bindSupplySchedule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SupplySchedule{SupplyScheduleCaller: SupplyScheduleCaller{contract: contract}, SupplyScheduleTransactor: SupplyScheduleTransactor{contract: contract}, SupplyScheduleFilterer: SupplyScheduleFilterer{contract: contract}}, nil
}

// NewSupplyScheduleCaller creates a new read-only instance of SupplySchedule, bound to a specific deployed contract.
func NewSupplyScheduleCaller(address common.Address, caller bind.ContractCaller) (*SupplyScheduleCaller, error) {
	contract, err := bindSupplySchedule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SupplyScheduleCaller{contract: contract}, nil
}

// NewSupplyScheduleTransactor creates a new write-only instance of SupplySchedule, bound to a specific deployed contract.
func NewSupplyScheduleTransactor(address common.Address, transactor bind.ContractTransactor) (*SupplyScheduleTransactor, error) {
	contract, err := bindSupplySchedule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SupplyScheduleTransactor{contract: contract}, nil
}

// NewSupplyScheduleFilterer creates a new log filterer instance of SupplySchedule, bound to a specific deployed contract.
func NewSupplyScheduleFilterer(address common.Address, filterer bind.ContractFilterer) (*SupplyScheduleFilterer, error) {
	contract, err := bindSupplySchedule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SupplyScheduleFilterer{contract: contract}, nil
}

// bindSupplySchedule binds a generic wrapper to an already deployed contract.
func bindSupplySchedule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SupplyScheduleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupplySchedule *SupplyScheduleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SupplySchedule.Contract.SupplyScheduleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupplySchedule *SupplyScheduleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupplySchedule.Contract.SupplyScheduleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupplySchedule *SupplyScheduleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SupplySchedule.Contract.SupplyScheduleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupplySchedule *SupplyScheduleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SupplySchedule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupplySchedule *SupplyScheduleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupplySchedule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupplySchedule *SupplyScheduleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SupplySchedule.Contract.contract.Transact(opts, method, params...)
}

// DECAYRATE is a free data retrieval call binding the contract method 0x1de40e49.
//
// Solidity: function DECAY_RATE() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCaller) DECAYRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "DECAY_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DECAYRATE is a free data retrieval call binding the contract method 0x1de40e49.
//
// Solidity: function DECAY_RATE() view returns(uint256)
func (_SupplySchedule *SupplyScheduleSession) DECAYRATE() (*big.Int, error) {
	return _SupplySchedule.Contract.DECAYRATE(&_SupplySchedule.CallOpts)
}

// DECAYRATE is a free data retrieval call binding the contract method 0x1de40e49.
//
// Solidity: function DECAY_RATE() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCallerSession) DECAYRATE() (*big.Int, error) {
	return _SupplySchedule.Contract.DECAYRATE(&_SupplySchedule.CallOpts)
}

// INFLATIONSTARTDATE is a free data retrieval call binding the contract method 0x7e1b823f.
//
// Solidity: function INFLATION_START_DATE() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCaller) INFLATIONSTARTDATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "INFLATION_START_DATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INFLATIONSTARTDATE is a free data retrieval call binding the contract method 0x7e1b823f.
//
// Solidity: function INFLATION_START_DATE() view returns(uint256)
func (_SupplySchedule *SupplyScheduleSession) INFLATIONSTARTDATE() (*big.Int, error) {
	return _SupplySchedule.Contract.INFLATIONSTARTDATE(&_SupplySchedule.CallOpts)
}

// INFLATIONSTARTDATE is a free data retrieval call binding the contract method 0x7e1b823f.
//
// Solidity: function INFLATION_START_DATE() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCallerSession) INFLATIONSTARTDATE() (*big.Int, error) {
	return _SupplySchedule.Contract.INFLATIONSTARTDATE(&_SupplySchedule.CallOpts)
}

// INITIALWEEKLYSUPPLY is a free data retrieval call binding the contract method 0xbadef30a.
//
// Solidity: function INITIAL_WEEKLY_SUPPLY() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCaller) INITIALWEEKLYSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "INITIAL_WEEKLY_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALWEEKLYSUPPLY is a free data retrieval call binding the contract method 0xbadef30a.
//
// Solidity: function INITIAL_WEEKLY_SUPPLY() view returns(uint256)
func (_SupplySchedule *SupplyScheduleSession) INITIALWEEKLYSUPPLY() (*big.Int, error) {
	return _SupplySchedule.Contract.INITIALWEEKLYSUPPLY(&_SupplySchedule.CallOpts)
}

// INITIALWEEKLYSUPPLY is a free data retrieval call binding the contract method 0xbadef30a.
//
// Solidity: function INITIAL_WEEKLY_SUPPLY() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCallerSession) INITIALWEEKLYSUPPLY() (*big.Int, error) {
	return _SupplySchedule.Contract.INITIALWEEKLYSUPPLY(&_SupplySchedule.CallOpts)
}

// MAXMINTERREWARD is a free data retrieval call binding the contract method 0x22af2bab.
//
// Solidity: function MAX_MINTER_REWARD() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCaller) MAXMINTERREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "MAX_MINTER_REWARD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMINTERREWARD is a free data retrieval call binding the contract method 0x22af2bab.
//
// Solidity: function MAX_MINTER_REWARD() view returns(uint256)
func (_SupplySchedule *SupplyScheduleSession) MAXMINTERREWARD() (*big.Int, error) {
	return _SupplySchedule.Contract.MAXMINTERREWARD(&_SupplySchedule.CallOpts)
}

// MAXMINTERREWARD is a free data retrieval call binding the contract method 0x22af2bab.
//
// Solidity: function MAX_MINTER_REWARD() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCallerSession) MAXMINTERREWARD() (*big.Int, error) {
	return _SupplySchedule.Contract.MAXMINTERREWARD(&_SupplySchedule.CallOpts)
}

// MINTBUFFER is a free data retrieval call binding the contract method 0x46872a23.
//
// Solidity: function MINT_BUFFER() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCaller) MINTBUFFER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "MINT_BUFFER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTBUFFER is a free data retrieval call binding the contract method 0x46872a23.
//
// Solidity: function MINT_BUFFER() view returns(uint256)
func (_SupplySchedule *SupplyScheduleSession) MINTBUFFER() (*big.Int, error) {
	return _SupplySchedule.Contract.MINTBUFFER(&_SupplySchedule.CallOpts)
}

// MINTBUFFER is a free data retrieval call binding the contract method 0x46872a23.
//
// Solidity: function MINT_BUFFER() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCallerSession) MINTBUFFER() (*big.Int, error) {
	return _SupplySchedule.Contract.MINTBUFFER(&_SupplySchedule.CallOpts)
}

// MINTPERIODDURATION is a free data retrieval call binding the contract method 0xdf5a9fc1.
//
// Solidity: function MINT_PERIOD_DURATION() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCaller) MINTPERIODDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "MINT_PERIOD_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTPERIODDURATION is a free data retrieval call binding the contract method 0xdf5a9fc1.
//
// Solidity: function MINT_PERIOD_DURATION() view returns(uint256)
func (_SupplySchedule *SupplyScheduleSession) MINTPERIODDURATION() (*big.Int, error) {
	return _SupplySchedule.Contract.MINTPERIODDURATION(&_SupplySchedule.CallOpts)
}

// MINTPERIODDURATION is a free data retrieval call binding the contract method 0xdf5a9fc1.
//
// Solidity: function MINT_PERIOD_DURATION() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCallerSession) MINTPERIODDURATION() (*big.Int, error) {
	return _SupplySchedule.Contract.MINTPERIODDURATION(&_SupplySchedule.CallOpts)
}

// SUPPLYDECAYEND is a free data retrieval call binding the contract method 0x7c060557.
//
// Solidity: function SUPPLY_DECAY_END() view returns(uint8)
func (_SupplySchedule *SupplyScheduleCaller) SUPPLYDECAYEND(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "SUPPLY_DECAY_END")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SUPPLYDECAYEND is a free data retrieval call binding the contract method 0x7c060557.
//
// Solidity: function SUPPLY_DECAY_END() view returns(uint8)
func (_SupplySchedule *SupplyScheduleSession) SUPPLYDECAYEND() (uint8, error) {
	return _SupplySchedule.Contract.SUPPLYDECAYEND(&_SupplySchedule.CallOpts)
}

// SUPPLYDECAYEND is a free data retrieval call binding the contract method 0x7c060557.
//
// Solidity: function SUPPLY_DECAY_END() view returns(uint8)
func (_SupplySchedule *SupplyScheduleCallerSession) SUPPLYDECAYEND() (uint8, error) {
	return _SupplySchedule.Contract.SUPPLYDECAYEND(&_SupplySchedule.CallOpts)
}

// SUPPLYDECAYSTART is a free data retrieval call binding the contract method 0x251330f1.
//
// Solidity: function SUPPLY_DECAY_START() view returns(uint8)
func (_SupplySchedule *SupplyScheduleCaller) SUPPLYDECAYSTART(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "SUPPLY_DECAY_START")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SUPPLYDECAYSTART is a free data retrieval call binding the contract method 0x251330f1.
//
// Solidity: function SUPPLY_DECAY_START() view returns(uint8)
func (_SupplySchedule *SupplyScheduleSession) SUPPLYDECAYSTART() (uint8, error) {
	return _SupplySchedule.Contract.SUPPLYDECAYSTART(&_SupplySchedule.CallOpts)
}

// SUPPLYDECAYSTART is a free data retrieval call binding the contract method 0x251330f1.
//
// Solidity: function SUPPLY_DECAY_START() view returns(uint8)
func (_SupplySchedule *SupplyScheduleCallerSession) SUPPLYDECAYSTART() (uint8, error) {
	return _SupplySchedule.Contract.SUPPLYDECAYSTART(&_SupplySchedule.CallOpts)
}

// TERMINALSUPPLYRATEANNUAL is a free data retrieval call binding the contract method 0x25542064.
//
// Solidity: function TERMINAL_SUPPLY_RATE_ANNUAL() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCaller) TERMINALSUPPLYRATEANNUAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "TERMINAL_SUPPLY_RATE_ANNUAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TERMINALSUPPLYRATEANNUAL is a free data retrieval call binding the contract method 0x25542064.
//
// Solidity: function TERMINAL_SUPPLY_RATE_ANNUAL() view returns(uint256)
func (_SupplySchedule *SupplyScheduleSession) TERMINALSUPPLYRATEANNUAL() (*big.Int, error) {
	return _SupplySchedule.Contract.TERMINALSUPPLYRATEANNUAL(&_SupplySchedule.CallOpts)
}

// TERMINALSUPPLYRATEANNUAL is a free data retrieval call binding the contract method 0x25542064.
//
// Solidity: function TERMINAL_SUPPLY_RATE_ANNUAL() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCallerSession) TERMINALSUPPLYRATEANNUAL() (*big.Int, error) {
	return _SupplySchedule.Contract.TERMINALSUPPLYRATEANNUAL(&_SupplySchedule.CallOpts)
}

// IsMintable is a free data retrieval call binding the contract method 0x46b45af7.
//
// Solidity: function isMintable() view returns(bool)
func (_SupplySchedule *SupplyScheduleCaller) IsMintable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "isMintable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMintable is a free data retrieval call binding the contract method 0x46b45af7.
//
// Solidity: function isMintable() view returns(bool)
func (_SupplySchedule *SupplyScheduleSession) IsMintable() (bool, error) {
	return _SupplySchedule.Contract.IsMintable(&_SupplySchedule.CallOpts)
}

// IsMintable is a free data retrieval call binding the contract method 0x46b45af7.
//
// Solidity: function isMintable() view returns(bool)
func (_SupplySchedule *SupplyScheduleCallerSession) IsMintable() (bool, error) {
	return _SupplySchedule.Contract.IsMintable(&_SupplySchedule.CallOpts)
}

// LastMintEvent is a free data retrieval call binding the contract method 0xbe801f01.
//
// Solidity: function lastMintEvent() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCaller) LastMintEvent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "lastMintEvent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastMintEvent is a free data retrieval call binding the contract method 0xbe801f01.
//
// Solidity: function lastMintEvent() view returns(uint256)
func (_SupplySchedule *SupplyScheduleSession) LastMintEvent() (*big.Int, error) {
	return _SupplySchedule.Contract.LastMintEvent(&_SupplySchedule.CallOpts)
}

// LastMintEvent is a free data retrieval call binding the contract method 0xbe801f01.
//
// Solidity: function lastMintEvent() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCallerSession) LastMintEvent() (*big.Int, error) {
	return _SupplySchedule.Contract.LastMintEvent(&_SupplySchedule.CallOpts)
}

// MintableSupply is a free data retrieval call binding the contract method 0xcc5c095c.
//
// Solidity: function mintableSupply() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCaller) MintableSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "mintableSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintableSupply is a free data retrieval call binding the contract method 0xcc5c095c.
//
// Solidity: function mintableSupply() view returns(uint256)
func (_SupplySchedule *SupplyScheduleSession) MintableSupply() (*big.Int, error) {
	return _SupplySchedule.Contract.MintableSupply(&_SupplySchedule.CallOpts)
}

// MintableSupply is a free data retrieval call binding the contract method 0xcc5c095c.
//
// Solidity: function mintableSupply() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCallerSession) MintableSupply() (*big.Int, error) {
	return _SupplySchedule.Contract.MintableSupply(&_SupplySchedule.CallOpts)
}

// MinterReward is a free data retrieval call binding the contract method 0x9bdd7ac7.
//
// Solidity: function minterReward() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCaller) MinterReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "minterReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinterReward is a free data retrieval call binding the contract method 0x9bdd7ac7.
//
// Solidity: function minterReward() view returns(uint256)
func (_SupplySchedule *SupplyScheduleSession) MinterReward() (*big.Int, error) {
	return _SupplySchedule.Contract.MinterReward(&_SupplySchedule.CallOpts)
}

// MinterReward is a free data retrieval call binding the contract method 0x9bdd7ac7.
//
// Solidity: function minterReward() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCallerSession) MinterReward() (*big.Int, error) {
	return _SupplySchedule.Contract.MinterReward(&_SupplySchedule.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SupplySchedule *SupplyScheduleCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SupplySchedule *SupplyScheduleSession) NominatedOwner() (common.Address, error) {
	return _SupplySchedule.Contract.NominatedOwner(&_SupplySchedule.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SupplySchedule *SupplyScheduleCallerSession) NominatedOwner() (common.Address, error) {
	return _SupplySchedule.Contract.NominatedOwner(&_SupplySchedule.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SupplySchedule *SupplyScheduleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SupplySchedule *SupplyScheduleSession) Owner() (common.Address, error) {
	return _SupplySchedule.Contract.Owner(&_SupplySchedule.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SupplySchedule *SupplyScheduleCallerSession) Owner() (common.Address, error) {
	return _SupplySchedule.Contract.Owner(&_SupplySchedule.CallOpts)
}

// SynthetixProxy is a free data retrieval call binding the contract method 0xbdd12482.
//
// Solidity: function synthetixProxy() view returns(address)
func (_SupplySchedule *SupplyScheduleCaller) SynthetixProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "synthetixProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SynthetixProxy is a free data retrieval call binding the contract method 0xbdd12482.
//
// Solidity: function synthetixProxy() view returns(address)
func (_SupplySchedule *SupplyScheduleSession) SynthetixProxy() (common.Address, error) {
	return _SupplySchedule.Contract.SynthetixProxy(&_SupplySchedule.CallOpts)
}

// SynthetixProxy is a free data retrieval call binding the contract method 0xbdd12482.
//
// Solidity: function synthetixProxy() view returns(address)
func (_SupplySchedule *SupplyScheduleCallerSession) SynthetixProxy() (common.Address, error) {
	return _SupplySchedule.Contract.SynthetixProxy(&_SupplySchedule.CallOpts)
}

// TerminalInflationSupply is a free data retrieval call binding the contract method 0xb3b2bcc0.
//
// Solidity: function terminalInflationSupply(uint256 totalSupply, uint256 numOfWeeks) pure returns(uint256)
func (_SupplySchedule *SupplyScheduleCaller) TerminalInflationSupply(opts *bind.CallOpts, totalSupply *big.Int, numOfWeeks *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "terminalInflationSupply", totalSupply, numOfWeeks)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TerminalInflationSupply is a free data retrieval call binding the contract method 0xb3b2bcc0.
//
// Solidity: function terminalInflationSupply(uint256 totalSupply, uint256 numOfWeeks) pure returns(uint256)
func (_SupplySchedule *SupplyScheduleSession) TerminalInflationSupply(totalSupply *big.Int, numOfWeeks *big.Int) (*big.Int, error) {
	return _SupplySchedule.Contract.TerminalInflationSupply(&_SupplySchedule.CallOpts, totalSupply, numOfWeeks)
}

// TerminalInflationSupply is a free data retrieval call binding the contract method 0xb3b2bcc0.
//
// Solidity: function terminalInflationSupply(uint256 totalSupply, uint256 numOfWeeks) pure returns(uint256)
func (_SupplySchedule *SupplyScheduleCallerSession) TerminalInflationSupply(totalSupply *big.Int, numOfWeeks *big.Int) (*big.Int, error) {
	return _SupplySchedule.Contract.TerminalInflationSupply(&_SupplySchedule.CallOpts, totalSupply, numOfWeeks)
}

// TokenDecaySupplyForWeek is a free data retrieval call binding the contract method 0x4e070f50.
//
// Solidity: function tokenDecaySupplyForWeek(uint256 counter) pure returns(uint256)
func (_SupplySchedule *SupplyScheduleCaller) TokenDecaySupplyForWeek(opts *bind.CallOpts, counter *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "tokenDecaySupplyForWeek", counter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenDecaySupplyForWeek is a free data retrieval call binding the contract method 0x4e070f50.
//
// Solidity: function tokenDecaySupplyForWeek(uint256 counter) pure returns(uint256)
func (_SupplySchedule *SupplyScheduleSession) TokenDecaySupplyForWeek(counter *big.Int) (*big.Int, error) {
	return _SupplySchedule.Contract.TokenDecaySupplyForWeek(&_SupplySchedule.CallOpts, counter)
}

// TokenDecaySupplyForWeek is a free data retrieval call binding the contract method 0x4e070f50.
//
// Solidity: function tokenDecaySupplyForWeek(uint256 counter) pure returns(uint256)
func (_SupplySchedule *SupplyScheduleCallerSession) TokenDecaySupplyForWeek(counter *big.Int) (*big.Int, error) {
	return _SupplySchedule.Contract.TokenDecaySupplyForWeek(&_SupplySchedule.CallOpts, counter)
}

// WeekCounter is a free data retrieval call binding the contract method 0xd3bd4bde.
//
// Solidity: function weekCounter() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCaller) WeekCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "weekCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeekCounter is a free data retrieval call binding the contract method 0xd3bd4bde.
//
// Solidity: function weekCounter() view returns(uint256)
func (_SupplySchedule *SupplyScheduleSession) WeekCounter() (*big.Int, error) {
	return _SupplySchedule.Contract.WeekCounter(&_SupplySchedule.CallOpts)
}

// WeekCounter is a free data retrieval call binding the contract method 0xd3bd4bde.
//
// Solidity: function weekCounter() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCallerSession) WeekCounter() (*big.Int, error) {
	return _SupplySchedule.Contract.WeekCounter(&_SupplySchedule.CallOpts)
}

// WeeksSinceLastIssuance is a free data retrieval call binding the contract method 0xdbd3a6a7.
//
// Solidity: function weeksSinceLastIssuance() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCaller) WeeksSinceLastIssuance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SupplySchedule.contract.Call(opts, &out, "weeksSinceLastIssuance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeeksSinceLastIssuance is a free data retrieval call binding the contract method 0xdbd3a6a7.
//
// Solidity: function weeksSinceLastIssuance() view returns(uint256)
func (_SupplySchedule *SupplyScheduleSession) WeeksSinceLastIssuance() (*big.Int, error) {
	return _SupplySchedule.Contract.WeeksSinceLastIssuance(&_SupplySchedule.CallOpts)
}

// WeeksSinceLastIssuance is a free data retrieval call binding the contract method 0xdbd3a6a7.
//
// Solidity: function weeksSinceLastIssuance() view returns(uint256)
func (_SupplySchedule *SupplyScheduleCallerSession) WeeksSinceLastIssuance() (*big.Int, error) {
	return _SupplySchedule.Contract.WeeksSinceLastIssuance(&_SupplySchedule.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SupplySchedule *SupplyScheduleTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupplySchedule.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SupplySchedule *SupplyScheduleSession) AcceptOwnership() (*types.Transaction, error) {
	return _SupplySchedule.Contract.AcceptOwnership(&_SupplySchedule.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SupplySchedule *SupplyScheduleTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SupplySchedule.Contract.AcceptOwnership(&_SupplySchedule.TransactOpts)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SupplySchedule *SupplyScheduleTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _SupplySchedule.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SupplySchedule *SupplyScheduleSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _SupplySchedule.Contract.NominateNewOwner(&_SupplySchedule.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SupplySchedule *SupplyScheduleTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _SupplySchedule.Contract.NominateNewOwner(&_SupplySchedule.TransactOpts, _owner)
}

// RecordMintEvent is a paid mutator transaction binding the contract method 0x7e7961d7.
//
// Solidity: function recordMintEvent(uint256 supplyMinted) returns(bool)
func (_SupplySchedule *SupplyScheduleTransactor) RecordMintEvent(opts *bind.TransactOpts, supplyMinted *big.Int) (*types.Transaction, error) {
	return _SupplySchedule.contract.Transact(opts, "recordMintEvent", supplyMinted)
}

// RecordMintEvent is a paid mutator transaction binding the contract method 0x7e7961d7.
//
// Solidity: function recordMintEvent(uint256 supplyMinted) returns(bool)
func (_SupplySchedule *SupplyScheduleSession) RecordMintEvent(supplyMinted *big.Int) (*types.Transaction, error) {
	return _SupplySchedule.Contract.RecordMintEvent(&_SupplySchedule.TransactOpts, supplyMinted)
}

// RecordMintEvent is a paid mutator transaction binding the contract method 0x7e7961d7.
//
// Solidity: function recordMintEvent(uint256 supplyMinted) returns(bool)
func (_SupplySchedule *SupplyScheduleTransactorSession) RecordMintEvent(supplyMinted *big.Int) (*types.Transaction, error) {
	return _SupplySchedule.Contract.RecordMintEvent(&_SupplySchedule.TransactOpts, supplyMinted)
}

// SetMinterReward is a paid mutator transaction binding the contract method 0x4ae26521.
//
// Solidity: function setMinterReward(uint256 amount) returns()
func (_SupplySchedule *SupplyScheduleTransactor) SetMinterReward(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SupplySchedule.contract.Transact(opts, "setMinterReward", amount)
}

// SetMinterReward is a paid mutator transaction binding the contract method 0x4ae26521.
//
// Solidity: function setMinterReward(uint256 amount) returns()
func (_SupplySchedule *SupplyScheduleSession) SetMinterReward(amount *big.Int) (*types.Transaction, error) {
	return _SupplySchedule.Contract.SetMinterReward(&_SupplySchedule.TransactOpts, amount)
}

// SetMinterReward is a paid mutator transaction binding the contract method 0x4ae26521.
//
// Solidity: function setMinterReward(uint256 amount) returns()
func (_SupplySchedule *SupplyScheduleTransactorSession) SetMinterReward(amount *big.Int) (*types.Transaction, error) {
	return _SupplySchedule.Contract.SetMinterReward(&_SupplySchedule.TransactOpts, amount)
}

// SetSynthetixProxy is a paid mutator transaction binding the contract method 0x95896b76.
//
// Solidity: function setSynthetixProxy(address _synthetixProxy) returns()
func (_SupplySchedule *SupplyScheduleTransactor) SetSynthetixProxy(opts *bind.TransactOpts, _synthetixProxy common.Address) (*types.Transaction, error) {
	return _SupplySchedule.contract.Transact(opts, "setSynthetixProxy", _synthetixProxy)
}

// SetSynthetixProxy is a paid mutator transaction binding the contract method 0x95896b76.
//
// Solidity: function setSynthetixProxy(address _synthetixProxy) returns()
func (_SupplySchedule *SupplyScheduleSession) SetSynthetixProxy(_synthetixProxy common.Address) (*types.Transaction, error) {
	return _SupplySchedule.Contract.SetSynthetixProxy(&_SupplySchedule.TransactOpts, _synthetixProxy)
}

// SetSynthetixProxy is a paid mutator transaction binding the contract method 0x95896b76.
//
// Solidity: function setSynthetixProxy(address _synthetixProxy) returns()
func (_SupplySchedule *SupplyScheduleTransactorSession) SetSynthetixProxy(_synthetixProxy common.Address) (*types.Transaction, error) {
	return _SupplySchedule.Contract.SetSynthetixProxy(&_SupplySchedule.TransactOpts, _synthetixProxy)
}

// SupplyScheduleMinterRewardUpdatedIterator is returned from FilterMinterRewardUpdated and is used to iterate over the raw logs and unpacked data for MinterRewardUpdated events raised by the SupplySchedule contract.
type SupplyScheduleMinterRewardUpdatedIterator struct {
	Event *SupplyScheduleMinterRewardUpdated // Event containing the contract specifics and raw log

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
func (it *SupplyScheduleMinterRewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyScheduleMinterRewardUpdated)
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
		it.Event = new(SupplyScheduleMinterRewardUpdated)
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
func (it *SupplyScheduleMinterRewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyScheduleMinterRewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyScheduleMinterRewardUpdated represents a MinterRewardUpdated event raised by the SupplySchedule contract.
type SupplyScheduleMinterRewardUpdated struct {
	NewRewardAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMinterRewardUpdated is a free log retrieval operation binding the contract event 0x036e0c635f8b7d9314bb6f2a577046108ef0f8b5e3869fbd29fd5a448ed99d30.
//
// Solidity: event MinterRewardUpdated(uint256 newRewardAmount)
func (_SupplySchedule *SupplyScheduleFilterer) FilterMinterRewardUpdated(opts *bind.FilterOpts) (*SupplyScheduleMinterRewardUpdatedIterator, error) {

	logs, sub, err := _SupplySchedule.contract.FilterLogs(opts, "MinterRewardUpdated")
	if err != nil {
		return nil, err
	}
	return &SupplyScheduleMinterRewardUpdatedIterator{contract: _SupplySchedule.contract, event: "MinterRewardUpdated", logs: logs, sub: sub}, nil
}

// WatchMinterRewardUpdated is a free log subscription operation binding the contract event 0x036e0c635f8b7d9314bb6f2a577046108ef0f8b5e3869fbd29fd5a448ed99d30.
//
// Solidity: event MinterRewardUpdated(uint256 newRewardAmount)
func (_SupplySchedule *SupplyScheduleFilterer) WatchMinterRewardUpdated(opts *bind.WatchOpts, sink chan<- *SupplyScheduleMinterRewardUpdated) (event.Subscription, error) {

	logs, sub, err := _SupplySchedule.contract.WatchLogs(opts, "MinterRewardUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyScheduleMinterRewardUpdated)
				if err := _SupplySchedule.contract.UnpackLog(event, "MinterRewardUpdated", log); err != nil {
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

// ParseMinterRewardUpdated is a log parse operation binding the contract event 0x036e0c635f8b7d9314bb6f2a577046108ef0f8b5e3869fbd29fd5a448ed99d30.
//
// Solidity: event MinterRewardUpdated(uint256 newRewardAmount)
func (_SupplySchedule *SupplyScheduleFilterer) ParseMinterRewardUpdated(log types.Log) (*SupplyScheduleMinterRewardUpdated, error) {
	event := new(SupplyScheduleMinterRewardUpdated)
	if err := _SupplySchedule.contract.UnpackLog(event, "MinterRewardUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyScheduleOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the SupplySchedule contract.
type SupplyScheduleOwnerChangedIterator struct {
	Event *SupplyScheduleOwnerChanged // Event containing the contract specifics and raw log

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
func (it *SupplyScheduleOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyScheduleOwnerChanged)
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
		it.Event = new(SupplyScheduleOwnerChanged)
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
func (it *SupplyScheduleOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyScheduleOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyScheduleOwnerChanged represents a OwnerChanged event raised by the SupplySchedule contract.
type SupplyScheduleOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SupplySchedule *SupplyScheduleFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*SupplyScheduleOwnerChangedIterator, error) {

	logs, sub, err := _SupplySchedule.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &SupplyScheduleOwnerChangedIterator{contract: _SupplySchedule.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SupplySchedule *SupplyScheduleFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *SupplyScheduleOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _SupplySchedule.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyScheduleOwnerChanged)
				if err := _SupplySchedule.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_SupplySchedule *SupplyScheduleFilterer) ParseOwnerChanged(log types.Log) (*SupplyScheduleOwnerChanged, error) {
	event := new(SupplyScheduleOwnerChanged)
	if err := _SupplySchedule.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyScheduleOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the SupplySchedule contract.
type SupplyScheduleOwnerNominatedIterator struct {
	Event *SupplyScheduleOwnerNominated // Event containing the contract specifics and raw log

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
func (it *SupplyScheduleOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyScheduleOwnerNominated)
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
		it.Event = new(SupplyScheduleOwnerNominated)
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
func (it *SupplyScheduleOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyScheduleOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyScheduleOwnerNominated represents a OwnerNominated event raised by the SupplySchedule contract.
type SupplyScheduleOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SupplySchedule *SupplyScheduleFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*SupplyScheduleOwnerNominatedIterator, error) {

	logs, sub, err := _SupplySchedule.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &SupplyScheduleOwnerNominatedIterator{contract: _SupplySchedule.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SupplySchedule *SupplyScheduleFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *SupplyScheduleOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _SupplySchedule.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyScheduleOwnerNominated)
				if err := _SupplySchedule.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_SupplySchedule *SupplyScheduleFilterer) ParseOwnerNominated(log types.Log) (*SupplyScheduleOwnerNominated, error) {
	event := new(SupplyScheduleOwnerNominated)
	if err := _SupplySchedule.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyScheduleSupplyMintedIterator is returned from FilterSupplyMinted and is used to iterate over the raw logs and unpacked data for SupplyMinted events raised by the SupplySchedule contract.
type SupplyScheduleSupplyMintedIterator struct {
	Event *SupplyScheduleSupplyMinted // Event containing the contract specifics and raw log

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
func (it *SupplyScheduleSupplyMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyScheduleSupplyMinted)
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
		it.Event = new(SupplyScheduleSupplyMinted)
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
func (it *SupplyScheduleSupplyMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyScheduleSupplyMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyScheduleSupplyMinted represents a SupplyMinted event raised by the SupplySchedule contract.
type SupplyScheduleSupplyMinted struct {
	SupplyMinted        *big.Int
	NumberOfWeeksIssued *big.Int
	LastMintEvent       *big.Int
	Timestamp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSupplyMinted is a free log retrieval operation binding the contract event 0x601e517d4811033fed8290c79b7823ce1ab70258da45400fe2391a3c7432edab.
//
// Solidity: event SupplyMinted(uint256 supplyMinted, uint256 numberOfWeeksIssued, uint256 lastMintEvent, uint256 timestamp)
func (_SupplySchedule *SupplyScheduleFilterer) FilterSupplyMinted(opts *bind.FilterOpts) (*SupplyScheduleSupplyMintedIterator, error) {

	logs, sub, err := _SupplySchedule.contract.FilterLogs(opts, "SupplyMinted")
	if err != nil {
		return nil, err
	}
	return &SupplyScheduleSupplyMintedIterator{contract: _SupplySchedule.contract, event: "SupplyMinted", logs: logs, sub: sub}, nil
}

// WatchSupplyMinted is a free log subscription operation binding the contract event 0x601e517d4811033fed8290c79b7823ce1ab70258da45400fe2391a3c7432edab.
//
// Solidity: event SupplyMinted(uint256 supplyMinted, uint256 numberOfWeeksIssued, uint256 lastMintEvent, uint256 timestamp)
func (_SupplySchedule *SupplyScheduleFilterer) WatchSupplyMinted(opts *bind.WatchOpts, sink chan<- *SupplyScheduleSupplyMinted) (event.Subscription, error) {

	logs, sub, err := _SupplySchedule.contract.WatchLogs(opts, "SupplyMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyScheduleSupplyMinted)
				if err := _SupplySchedule.contract.UnpackLog(event, "SupplyMinted", log); err != nil {
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

// ParseSupplyMinted is a log parse operation binding the contract event 0x601e517d4811033fed8290c79b7823ce1ab70258da45400fe2391a3c7432edab.
//
// Solidity: event SupplyMinted(uint256 supplyMinted, uint256 numberOfWeeksIssued, uint256 lastMintEvent, uint256 timestamp)
func (_SupplySchedule *SupplyScheduleFilterer) ParseSupplyMinted(log types.Log) (*SupplyScheduleSupplyMinted, error) {
	event := new(SupplyScheduleSupplyMinted)
	if err := _SupplySchedule.contract.UnpackLog(event, "SupplyMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyScheduleSynthetixProxyUpdatedIterator is returned from FilterSynthetixProxyUpdated and is used to iterate over the raw logs and unpacked data for SynthetixProxyUpdated events raised by the SupplySchedule contract.
type SupplyScheduleSynthetixProxyUpdatedIterator struct {
	Event *SupplyScheduleSynthetixProxyUpdated // Event containing the contract specifics and raw log

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
func (it *SupplyScheduleSynthetixProxyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyScheduleSynthetixProxyUpdated)
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
		it.Event = new(SupplyScheduleSynthetixProxyUpdated)
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
func (it *SupplyScheduleSynthetixProxyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyScheduleSynthetixProxyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyScheduleSynthetixProxyUpdated represents a SynthetixProxyUpdated event raised by the SupplySchedule contract.
type SupplyScheduleSynthetixProxyUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSynthetixProxyUpdated is a free log retrieval operation binding the contract event 0xf8df0556b7fde3c4b8394eae329aedfa59c6ffd8b532d572a1efcfa3424ca5fc.
//
// Solidity: event SynthetixProxyUpdated(address newAddress)
func (_SupplySchedule *SupplyScheduleFilterer) FilterSynthetixProxyUpdated(opts *bind.FilterOpts) (*SupplyScheduleSynthetixProxyUpdatedIterator, error) {

	logs, sub, err := _SupplySchedule.contract.FilterLogs(opts, "SynthetixProxyUpdated")
	if err != nil {
		return nil, err
	}
	return &SupplyScheduleSynthetixProxyUpdatedIterator{contract: _SupplySchedule.contract, event: "SynthetixProxyUpdated", logs: logs, sub: sub}, nil
}

// WatchSynthetixProxyUpdated is a free log subscription operation binding the contract event 0xf8df0556b7fde3c4b8394eae329aedfa59c6ffd8b532d572a1efcfa3424ca5fc.
//
// Solidity: event SynthetixProxyUpdated(address newAddress)
func (_SupplySchedule *SupplyScheduleFilterer) WatchSynthetixProxyUpdated(opts *bind.WatchOpts, sink chan<- *SupplyScheduleSynthetixProxyUpdated) (event.Subscription, error) {

	logs, sub, err := _SupplySchedule.contract.WatchLogs(opts, "SynthetixProxyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyScheduleSynthetixProxyUpdated)
				if err := _SupplySchedule.contract.UnpackLog(event, "SynthetixProxyUpdated", log); err != nil {
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

// ParseSynthetixProxyUpdated is a log parse operation binding the contract event 0xf8df0556b7fde3c4b8394eae329aedfa59c6ffd8b532d572a1efcfa3424ca5fc.
//
// Solidity: event SynthetixProxyUpdated(address newAddress)
func (_SupplySchedule *SupplyScheduleFilterer) ParseSynthetixProxyUpdated(log types.Log) (*SupplyScheduleSynthetixProxyUpdated, error) {
	event := new(SupplyScheduleSynthetixProxyUpdated)
	if err := _SupplySchedule.contract.UnpackLog(event, "SynthetixProxyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixABI is the input ABI used to generate the binding from.
const SynthetixABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"contractTokenState\",\"name\":\"_tokenState\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"snxRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLiquidated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"}],\"name\":\"AccountLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExchangeRebate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExchangeReclaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"toCurrencyKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"}],\"name\":\"ExchangeTracking\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"}],\"name\":\"ProxyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"fromCurrencyKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"toCurrencyKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"}],\"name\":\"SynthExchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTokenState\",\"type\":\"address\"}],\"name\":\"TokenStateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_ADDRESSES_FROM_RESOLVER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_SYMBOL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"anySynthOrSNXRateIsInvalid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"anyRateInvalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableCurrencyKeys\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableSynthCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"availableSynths\",\"outputs\":[{\"internalType\":\"contractISynth\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"burnSecondary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnSynths\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnForAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnSynthsOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"burnSynthsToTarget\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnForAddress\",\"type\":\"address\"}],\"name\":\"burnSynthsToTargetOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"}],\"name\":\"collateralisationRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"debtBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emitExchangeRebate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emitExchangeReclaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"toCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"}],\"name\":\"emitExchangeTracking\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"fromCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"toCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"}],\"name\":\"emitSynthExchange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"exchangeForAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"}],\"name\":\"exchangeOnBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"exchangeForAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"}],\"name\":\"exchangeOnBehalfWithTracking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"}],\"name\":\"exchangeWithTracking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCurrencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"trackingCode\",\"type\":\"bytes32\"}],\"name\":\"exchangeWithVirtual\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"},{\"internalType\":\"contractIVirtualSynth\",\"name\":\"vSynth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getResolverAddressesRequired\",\"outputs\":[{\"internalType\":\"bytes32[24]\",\"name\":\"addressesRequired\",\"type\":\"bytes32[24]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"integrationProxy\",\"outputs\":[{\"internalType\":\"contractProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractAddressResolver\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"isResolverCached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"isWaitingPeriod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"issueMaxSynths\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issueForAddress\",\"type\":\"address\"}],\"name\":\"issueMaxSynthsOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issueSynths\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issueForAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issueSynthsOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"susdAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateDelinquentAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxIssuableSynths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxIssuable\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"messageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mintSecondary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mintSecondaryRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"contractProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"remainingIssuableSynths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxIssuable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alreadyIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSystemDebt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"contractAddressResolver\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"resolverAddressesRequired\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sUSD\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_integrationProxy\",\"type\":\"address\"}],\"name\":\"setIntegrationProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"setMessageSender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"setProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractAddressResolver\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"setResolverAndSyncCache\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractTokenState\",\"name\":\"_tokenState\",\"type\":\"address\"}],\"name\":\"setTokenState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"settle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reclaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refunded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEntriesSettled\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"synths\",\"outputs\":[{\"internalType\":\"contractISynth\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"synthAddress\",\"type\":\"address\"}],\"name\":\"synthsByAddress\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenState\",\"outputs\":[{\"internalType\":\"contractTokenState\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"totalIssuedSynths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"totalIssuedSynthsExcludeEtherCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"transferableSynthetix\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"transferable\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SynthetixFuncSigs maps the 4-byte function signature to its string representation.
var SynthetixFuncSigs = map[string]string{
	"2e0f2625": "DECIMALS()",
	"e3235c91": "MAX_ADDRESSES_FROM_RESOLVER()",
	"18821400": "TOKEN_NAME()",
	"2a905318": "TOKEN_SYMBOL()",
	"79ba5097": "acceptOwnership()",
	"dd62ed3e": "allowance(address,address)",
	"4e99bda9": "anySynthOrSNXRateIsInvalid()",
	"095ea7b3": "approve(address,uint256)",
	"72cb051f": "availableCurrencyKeys()",
	"dbf63340": "availableSynthCount()",
	"835e119c": "availableSynths(uint256)",
	"70a08231": "balanceOf(address)",
	"edef719a": "burnSecondary(address,uint256)",
	"295da87d": "burnSynths(uint256)",
	"c2bf3880": "burnSynthsOnBehalf(address,uint256)",
	"9741fb22": "burnSynthsToTarget()",
	"2c955fa7": "burnSynthsToTargetOnBehalf(address)",
	"a5fdc5de": "collateral(address)",
	"a311c7c2": "collateralisationRatio(address)",
	"d37c4d8b": "debtBalanceOf(address,bytes32)",
	"313ce567": "decimals()",
	"6f01a986": "emitExchangeRebate(address,bytes32,uint256)",
	"ace88afd": "emitExchangeReclaim(address,bytes32,uint256)",
	"ddd03a3f": "emitExchangeTracking(bytes32,bytes32,uint256)",
	"6c00f310": "emitSynthExchange(address,bytes32,uint256,bytes32,uint256,address)",
	"ee52a2f3": "exchange(bytes32,uint256,bytes32)",
	"c836fa0a": "exchangeOnBehalf(address,bytes32,uint256,bytes32)",
	"91e56b68": "exchangeOnBehalfWithTracking(address,bytes32,uint256,bytes32,address,bytes32)",
	"30ead760": "exchangeWithTracking(bytes32,uint256,bytes32,address,bytes32)",
	"0e30963c": "exchangeWithVirtual(bytes32,uint256,bytes32,bytes32)",
	"ab49848c": "getResolverAddressesRequired()",
	"9cbdaeb6": "integrationProxy()",
	"631e1444": "isResolverCached(address)",
	"1fce304d": "isWaitingPeriod(bytes32)",
	"af086c7e": "issueMaxSynths()",
	"320223db": "issueMaxSynthsOnBehalf(address)",
	"8a290014": "issueSynths(uint256)",
	"e8e09b8b": "issueSynthsOnBehalf(address,uint256)",
	"e6203ed1": "liquidateDelinquentAccount(address,uint256)",
	"05b3c1c9": "maxIssuableSynths(address)",
	"d67bdd25": "messageSender()",
	"1249c58b": "mint()",
	"666ed4f1": "mintSecondary(address,uint256)",
	"d8a1f76f": "mintSecondaryRewards(uint256)",
	"06fdde03": "name()",
	"1627540c": "nominateNewOwner(address)",
	"53a47bb7": "nominatedOwner()",
	"8da5cb5b": "owner()",
	"ec556889": "proxy()",
	"1137aedf": "remainingIssuableSynths(address)",
	"04f3bcec": "resolver()",
	"c6c9d828": "resolverAddressesRequired(uint256)",
	"9324cac7": "sUSD()",
	"131b0ae7": "setIntegrationProxy(address)",
	"bc67f832": "setMessageSender(address)",
	"97107d6d": "setProxy(address)",
	"3be99e6f": "setResolverAndSyncCache(address)",
	"9f769807": "setTokenState(address)",
	"987757dd": "settle(bytes32)",
	"95d89b41": "symbol()",
	"32608039": "synths(bytes32)",
	"16b2213f": "synthsByAddress(address)",
	"e90dd9e2": "tokenState()",
	"83d625d4": "totalIssuedSynths(bytes32)",
	"d60888e4": "totalIssuedSynthsExcludeEtherCollateral(bytes32)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"6ac0bf9c": "transferableSynthetix(address)",
}

// SynthetixBin is the compiled bytecode used for deploying new contracts.
var SynthetixBin = "0x6101406040526b53797374656d53746174757360a01b60809081526822bc31b430b733b2b960b91b60a0526524b9b9bab2b960d11b60c0526d537570706c795363686564756c6560901b60e0527f52657761726473446973747269627574696f6e00000000000000000000000000610100526d53796e746865746978537461746560901b610120526200009790600c90600662000433565b50348015620000a557600080fd5b5060405162004d6838038062004d68833981810160405260a0811015620000cb57600080fd5b5080516020820151604080840151606085015160809095015182516103008101938490529495939491939290918291600c9060189082845b8154815260200190600101908083116200010357505050505086866040518060400160405280601781526020017f53796e746865746978204e6574776f726b20546f6b656e000000000000000000815250604051806040016040528060038152602001620a69cb60eb1b8152508760128a868160006001600160a01b0316816001600160a01b03161415620001df576040805162461bcd60e51b815260206004820152601960248201527f4f776e657220616464726573732063616e6e6f74206265203000000000000000604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b038316908117825560408051928352602083019190915280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a1506000546001600160a01b03166200028a576040805162461bcd60e51b815260206004820152601160248201527013dddb995c881b5d5cdd081899481cd95d607a1b604482015290519081900360640190fd5b600280546001600160a01b0383166001600160a01b0319909116811790915560408051918252517ffc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e9181900360200190a150600580546001600160a01b0319166001600160a01b03881617905584516200030c90600690602088019062000476565b5083516200032290600790602087019062000476565b50506008919091556009805460ff191660ff90921691909117905550506000546001600160a01b031615159150620003979050576040805162461bcd60e51b815260206004820152601160248201527013dddb995c881b5d5cdd081899481cd95d607a1b604482015290519081900360640190fd5b60005b6018811015620003ff576000828260188110620003b357fe5b602002015114620003f057600b828260188110620003cd57fe5b6020908102919091015182546001810184556000938452919092200155620003f6565b620003ff565b6001016200039a565b5050600980546001600160a01b0390921661010002610100600160a81b031990921691909117905550620005089350505050565b826018810192821562000464579160200282015b828111156200046457825182559160200191906001019062000447565b5062000472929150620004e8565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620004b957805160ff191683800117855562000464565b828001600101855582156200046457918201828111156200046457825182559160200191906001019062000447565b6200050591905b80821115620004725760008155600101620004ef565b90565b61485080620005186000396000f3fe608060405234801561001057600080fd5b50600436106104075760003560e01c806383d625d411610220578063bc67f83211610130578063dd62ed3e116100b8578063e8e09b8b11610087578063e8e09b8b14610cff578063e90dd9e214610d2b578063ec55688914610d33578063edef719a146107e6578063ee52a2f314610d3b57610407565b8063dd62ed3e14610c74578063ddd03a3f14610ca2578063e3235c9114610ccb578063e6203ed114610cd357610407565b8063d37c4d8b116100ff578063d37c4d8b14610bfe578063d60888e414610c2a578063d67bdd2514610c47578063d8a1f76f14610c4f578063dbf6334014610c6c57610407565b8063bc67f83214610b57578063c2bf388014610b7d578063c6c9d82814610ba9578063c836fa0a14610bc657610407565b8063987757dd116101b3578063a5fdc5de11610182578063a5fdc5de14610a8a578063a9059cbb14610ab0578063ab49848c14610adc578063ace88afd14610b1d578063af086c7e14610b4f57610407565b8063987757dd14610a195780639cbdaeb614610a365780639f76980714610a3e578063a311c7c214610a6457610407565b80639324cac7116101ef5780639324cac7146109db57806395d89b41146109e357806397107d6d146109eb5780639741fb2214610a1157610407565b806383d625d4146109535780638a290014146109705780638da5cb5b1461098d57806391e56b681461099557610407565b80632e0f26251161031b578063631e1444116102ae5780636f01a9861161027d5780636f01a9861461087e57806370a08231146108b057806372cb051f146108d657806379ba50971461092e578063835e119c1461093657610407565b8063631e1444146107c0578063666ed4f1146107e65780636ac0bf9c146108125780636c00f3101461083857610407565b806332608039116102ea578063326080391461076d5780633be99e6f1461078a5780634e99bda9146107b057806353a47bb7146107b857610407565b80632e0f2625146106e357806330ead76014610701578063313ce5671461073f578063320223db1461074757610407565b80631627540c1161039e5780631fce304d1161036d5780631fce304d1461064557806323b872dd14610662578063295da87d146106985780632a905318146106b55780632c955fa7146106bd57610407565b80631627540c146105e957806316b2213f1461060f57806318160ddd14610635578063188214001461063d57610407565b80630e30963c116103da5780630e30963c146105255780631137aedf146105755780631249c58b146105b9578063131b0ae7146105c157610407565b806304f3bcec1461040c57806305b3c1c91461043057806306fdde0314610468578063095ea7b3146104e5575b600080fd5b610414610d64565b604080516001600160a01b039092168252519081900360200190f35b6104566004803603602081101561044657600080fd5b50356001600160a01b0316610d78565b60408051918252519081900360200190f35b610470610e0b565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104aa578181015183820152602001610492565b50505050905090810190601f1680156104d75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610511600480360360408110156104fb57600080fd5b506001600160a01b038135169060200135610e99565b604080519115158252519081900360200190f35b6105546004803603608081101561053b57600080fd5b5080359060208101359060408101359060600135610f32565b604080519283526001600160a01b0390911660208301528051918290030190f35b61059b6004803603602081101561058b57600080fd5b50356001600160a01b03166110be565b60408051938452602084019290925282820152519081900360600190f35b610511611164565b6105e7600480360360208110156105d757600080fd5b50356001600160a01b0316611658565b005b6105e7600480360360208110156105ff57600080fd5b50356001600160a01b0316611682565b6104566004803603602081101561062557600080fd5b50356001600160a01b03166116de565b61045661173d565b610470611743565b6105116004803603602081101561065b57600080fd5b503561177c565b6105116004803603606081101561067857600080fd5b506001600160a01b0381358116916020810135909116906040013561180f565b6105e7600480360360208110156106ae57600080fd5b503561189c565b610470611976565b6105e7600480360360208110156106d357600080fd5b50356001600160a01b0316611995565b6106eb611a56565b6040805160ff9092168252519081900360200190f35b610456600480360360a081101561071757600080fd5b508035906020810135906040810135906001600160a01b036060820135169060800135611a5b565b6106eb611be5565b6105e76004803603602081101561075d57600080fd5b50356001600160a01b0316611bee565b6104146004803603602081101561078357600080fd5b5035611caf565b6105e7600480360360208110156107a057600080fd5b50356001600160a01b0316611cfc565b610511611e32565b610414611ea5565b610511600480360360208110156107d657600080fd5b50356001600160a01b0316611eb4565b6105e7600480360360408110156107fc57600080fd5b506001600160a01b038135169060200135611fdb565b6104566004803603602081101561082857600080fd5b50356001600160a01b0316612028565b6105e7600480360360c081101561084e57600080fd5b506001600160a01b03813581169160208101359160408201359160608101359160808201359160a001351661212f565b6105e76004803603606081101561089457600080fd5b506001600160a01b0381351690602081013590604001356122e1565b610456600480360360208110156108c657600080fd5b50356001600160a01b0316612475565b6108de6124c6565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561091a578181015183820152602001610902565b505050509050019250505060405180910390f35b6105e76125d6565b6104146004803603602081101561094c57600080fd5b5035612692565b6104566004803603602081101561096957600080fd5b50356126df565b6105e76004803603602081101561098657600080fd5b5035612739565b6104146127f6565b610456600480360360c08110156109ab57600080fd5b506001600160a01b0381358116916020810135916040820135916060810135916080820135169060a00135612805565b610456612993565b61047061299e565b6105e760048036036020811015610a0157600080fd5b50356001600160a01b03166129f9565b6105e7612a55565b61059b60048036036020811015610a2f57600080fd5b5035612b27565b610414612ba7565b6105e760048036036020811015610a5457600080fd5b50356001600160a01b0316612bb6565b61045660048036036020811015610a7a57600080fd5b50356001600160a01b0316612be2565b61045660048036036020811015610aa057600080fd5b50356001600160a01b0316612c41565b61051160048036036040811015610ac657600080fd5b506001600160a01b038135169060200135612ca0565b610ae4612d31565b604051808261030080838360005b83811015610b0a578181015183820152602001610af2565b5050505090500191505060405180910390f35b6105e760048036036060811015610b3357600080fd5b506001600160a01b038135169060208101359060400135612d7b565b6105e7612e2e565b6105e760048036036020811015610b6d57600080fd5b50356001600160a01b0316612ee5565b6105e760048036036040811015610b9357600080fd5b506001600160a01b038135169060200135612f0f565b61045660048036036020811015610bbf57600080fd5b5035612ff3565b61045660048036036080811015610bdc57600080fd5b506001600160a01b038135169060208101359060408101359060600135613011565b61045660048036036040811015610c1457600080fd5b506001600160a01b03813516906020013561318e565b61045660048036036020811015610c4057600080fd5b5035613228565b610414613282565b6105e760048036036020811015610c6557600080fd5b5035611fdb565b610456613291565b61045660048036036040811015610c8a57600080fd5b506001600160a01b03813581169160200135166132d3565b6105e760048036036060811015610cb857600080fd5b508035906020810135906040013561332c565b610456613449565b61051160048036036040811015610ce957600080fd5b506001600160a01b03813516906020013561344e565b6105e760048036036040811015610d1557600080fd5b506001600160a01b03813516906020013561358e565b610414613656565b610414613665565b61045660048036036060811015610d5157600080fd5b5080359060208101359060400135613674565b60095461010090046001600160a01b031681565b6000610d826137ed565b6001600160a01b03166305b3c1c9836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015610dd757600080fd5b505afa158015610deb573d6000803e3d6000fd5b505050506040513d6020811015610e0157600080fd5b505190505b919050565b6006805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610e915780601f10610e6657610100808354040283529160200191610e91565b820191906000526020600020905b815481529060010190602001808311610e7457829003601f168201915b505050505081565b6000610ea3613835565b6004805460055460408051633691826360e21b81526001600160a01b039384169481018590528784166024820152604481018790529051919092169163da46098c91606480830192600092919082900301818387803b158015610f0557600080fd5b505af1158015610f19573d6000803e3d6000fd5b50505050610f2881858561388b565b5060019392505050565b6000808584610f3f613955565b6001600160a01b0316637118d4316040518163ffffffff1660e01b815260040160006040518083038186803b158015610f7757600080fd5b505afa158015610f8b573d6000803e3d6000fd5b50505050610f97613955565b6001600160a01b0316636132eba483836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b158015610fe257600080fd5b505afa158015610ff6573d6000803e3d6000fd5b50505050611002613835565b61100a6139a5565b6004805460408051633ce6548960e21b81526001600160a01b03928316938101849052602481018d9052604481018c9052606481018b9052608481019390935260a483018990528051939091169263f39952249260c48082019392918290030181600087803b15801561107c57600080fd5b505af1158015611090573d6000803e3d6000fd5b505050506040513d60408110156110a657600080fd5b50805160209091015190999098509650505050505050565b60008060006110cb6137ed565b6001600160a01b0316631137aedf856040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060606040518083038186803b15801561112057600080fd5b505afa158015611134573d6000803e3d6000fd5b505050506040513d606081101561114a57600080fd5b508051602082015160409092015190969195509350915050565b600061116e613955565b6001600160a01b0316637c3125416040518163ffffffff1660e01b815260040160006040518083038186803b1580156111a657600080fd5b505afa1580156111ba573d6000803e3d6000fd5b5050505060006001600160a01b03166111d16139f2565b6001600160a01b0316141561122d576040805162461bcd60e51b815260206004820152601b60248201527f52657761726473446973747269627574696f6e206e6f74207365740000000000604482015290519081900360640190fd5b6000611237613a2c565b905060006112436139f2565b90506000826001600160a01b031663cc5c095c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561128057600080fd5b505afa158015611294573d6000803e3d6000fd5b505050506040513d60208110156112aa57600080fd5b50519050806112f8576040805162461bcd60e51b81526020600482015260156024820152744e6f20737570706c79206973206d696e7461626c6560581b604482015290519081900360640190fd5b826001600160a01b0316637e7961d7826040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561133e57600080fd5b505af1158015611352573d6000803e3d6000fd5b505050506040513d602081101561136857600080fd5b505060408051639bdd7ac760e01b815290516000916001600160a01b03861691639bdd7ac791600480820192602092909190829003018186803b1580156113ae57600080fd5b505afa1580156113c2573d6000803e3d6000fd5b505050506040513d60208110156113d857600080fd5b5051905060006113ee838363ffffffff613a7e16565b600554604080516370a0823160e01b81526001600160a01b038881166004830152915193945091169163b46310f691879161148491869186916370a08231916024808301926020929190829003018186803b15801561144c57600080fd5b505afa158015611460573d6000803e3d6000fd5b505050506040513d602081101561147657600080fd5b50519063ffffffff613adb16565b6040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156114d357600080fd5b505af11580156114e7573d6000803e3d6000fd5b505050506114f6308583613b3c565b836001600160a01b03166359974e38826040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561153c57600080fd5b505af1158015611550573d6000803e3d6000fd5b505050506040513d602081101561156657600080fd5b5050600554604080516370a0823160e01b8152336004820181905291516001600160a01b039093169263b46310f692916115c491879186916370a0823191602480820192602092909190829003018186803b15801561144c57600080fd5b6040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561161357600080fd5b505af1158015611627573d6000803e3d6000fd5b50505050611636303384613b3c565b600854611649908463ffffffff613adb16565b60085550600194505050505090565b611660613b8f565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b61168a613b8f565b600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce229181900360200190a150565b60006116e86137ed565b6001600160a01b03166316b2213f836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015610dd757600080fd5b60085481565b6040518060400160405280601781526020017f53796e746865746978204e6574776f726b20546f6b656e00000000000000000081525081565b6000806117876139a5565b60048054604080516301670a7b60e21b81526001600160a01b0392831693810193909352602483018790525192169163059c29ec91604480820192602092909190829003018186803b1580156117dc57600080fd5b505afa1580156117f0573d6000803e3d6000fd5b505050506040513d602081101561180657600080fd5b50511192915050565b6000611819613835565b611821613955565b6001600160a01b031663086dabd16040518163ffffffff1660e01b815260040160006040518083038186803b15801561185957600080fd5b505afa15801561186d573d6000803e3d6000fd5b5050505061187b8483613bd8565b50600454611894906001600160a01b0316858585613e09565b949350505050565b6118a4613955565b6001600160a01b0316637c3125416040518163ffffffff1660e01b815260040160006040518083038186803b1580156118dc57600080fd5b505afa1580156118f0573d6000803e3d6000fd5b505050506118fc613835565b6119046137ed565b600480546040805163b06e8c6560e01b81526001600160a01b0392831693810193909352602483018590525192169163b06e8c659160448082019260009290919082900301818387803b15801561195a57600080fd5b505af115801561196e573d6000803e3d6000fd5b505050505b50565b604051806040016040528060038152602001620a69cb60eb1b81525081565b61199d613955565b6001600160a01b0316637c3125416040518163ffffffff1660e01b815260040160006040518083038186803b1580156119d557600080fd5b505afa1580156119e9573d6000803e3d6000fd5b505050506119f5613835565b6119fd6137ed565b600480546040805163159fa0d560e11b81526001600160a01b03868116948201949094529183166024830152519290911691632b3f41aa9160448082019260009290919082900301818387803b15801561195a57600080fd5b601281565b60008584611a67613955565b6001600160a01b0316637118d4316040518163ffffffff1660e01b815260040160006040518083038186803b158015611a9f57600080fd5b505afa158015611ab3573d6000803e3d6000fd5b50505050611abf613955565b6001600160a01b0316636132eba483836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b158015611b0a57600080fd5b505afa158015611b1e573d6000803e3d6000fd5b50505050611b2a613835565b611b326139a5565b60048054604080516321aea91760e21b81526001600160a01b03928316938101849052602481018d9052604481018c9052606481018b9052608481019390935288821660a484015260c48301889052519216916386baa45c9160e4808201926020929091908290030181600087803b158015611bad57600080fd5b505af1158015611bc1573d6000803e3d6000fd5b505050506040513d6020811015611bd757600080fd5b505198975050505050505050565b60095460ff1681565b611bf6613955565b6001600160a01b0316637c3125416040518163ffffffff1660e01b815260040160006040518083038186803b158015611c2e57600080fd5b505afa158015611c42573d6000803e3d6000fd5b50505050611c4e613835565b611c566137ed565b600480546040805163fd864ccf60e01b81526001600160a01b0386811694820194909452918316602483015251929091169163fd864ccf9160448082019260009290919082900301818387803b15801561195a57600080fd5b6000611cb96137ed565b6001600160a01b03166332608039836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610dd757600080fd5b611d04613b8f565b60098054610100600160a81b0319166101006001600160a01b0384160217905560005b600b54811015611e2e576000600b8281548110611d4057fe5b600091825260209182902001546009546040805163dacb2d0160e01b81526004810184905260248101829052601760448201527f5265736f6c766572206d697373696e6720746172676574000000000000000000606482015290519294506101009091046001600160a01b03169263dacb2d0192608480840193829003018186803b158015611dce57600080fd5b505afa158015611de2573d6000803e3d6000fd5b505050506040513d6020811015611df857600080fd5b50516000918252600a602052604090912080546001600160a01b0319166001600160a01b03909216919091179055600101611d27565b5050565b6000611e3c6137ed565b6001600160a01b0316634e99bda96040518163ffffffff1660e01b815260040160206040518083038186803b158015611e7457600080fd5b505afa158015611e88573d6000803e3d6000fd5b505050506040513d6020811015611e9e57600080fd5b5051905090565b6001546001600160a01b031681565b6009546000906001600160a01b038381166101009092041614611ed957506000610e06565b60005b600b54811015611fd2576000600b8281548110611ef557fe5b6000918252602080832090910154808352600a82526040928390205460095484516321f8a72160e01b81526004810184905294519295506001600160a01b0391821694610100909104909116926321f8a72192602480840193829003018186803b158015611f6257600080fd5b505afa158015611f76573d6000803e3d6000fd5b505050506040513d6020811015611f8c57600080fd5b50516001600160a01b0316141580611fb957506000818152600a60205260409020546001600160a01b0316155b15611fc957600092505050610e06565b50600101611edc565b50600192915050565b6040805162461bcd60e51b815260206004820152601b60248201527f43616e6e6f742062652072756e206f6e2074686973206c617965720000000000604482015290519081900360640190fd5b60006120326137ed565b600554604080516370a0823160e01b81526001600160a01b038681166004830152915193821693636bed041593879316916370a08231916024808301926020929190829003018186803b15801561208857600080fd5b505afa15801561209c573d6000803e3d6000fd5b505050506040513d60208110156120b257600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b03909316600484015260248301919091528051604480840193829003018186803b1580156120fd57600080fd5b505afa158015612111573d6000803e3d6000fd5b505050506040513d604081101561212757600080fd5b505192915050565b6121376139a5565b6001600160a01b0316336001600160a01b03161461218a576040805162461bcd60e51b815260206004820152601e6024820152600080516020614682833981519152604482015290519081900360640190fd5b60028054604080516020810189905280820188905260608101879052608081018690526001600160a01b0385811660a0808401919091528351808403909101815260c0909201928390529092169263907dff97929180603e6146a28239603e01905060405180910390206121fd8b613f37565b6000806040518763ffffffff1660e01b815260040180806020018781526020018681526020018581526020018460001b81526020018360001b8152602001828103825288818151815260200191508051906020019080838360005b83811015612270578181015183820152602001612258565b50505050905090810190601f16801561229d5780820380516001836020036101000a031916815260200191505b50975050505050505050600060405180830381600087803b1580156122c157600080fd5b505af11580156122d5573d6000803e3d6000fd5b50505050505050505050565b6122e96139a5565b6001600160a01b0316336001600160a01b03161461233c576040805162461bcd60e51b815260206004820152601e6024820152600080516020614682833981519152604482015290519081900360640190fd5b6002805460408051602081018690528082018590528151808203830181526060909101918290526001600160a01b039092169263907dff97929180602761477982396027019050604051809103902061239488613f37565b6000806040518763ffffffff1660e01b815260040180806020018781526020018681526020018581526020018460001b81526020018360001b8152602001828103825288818151815260200191508051906020019080838360005b838110156124075781810151838201526020016123ef565b50505050905090810190601f1680156124345780820380516001836020036101000a031916815260200191505b50975050505050505050600060405180830381600087803b15801561245857600080fd5b505af115801561246c573d6000803e3d6000fd5b50505050505050565b600554604080516370a0823160e01b81526001600160a01b038481166004830152915160009392909216916370a0823191602480820192602092909190829003018186803b158015610dd757600080fd5b60606124d06137ed565b6001600160a01b03166372cb051f6040518163ffffffff1660e01b815260040160006040518083038186803b15801561250857600080fd5b505afa15801561251c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561254557600080fd5b810190808051604051939291908464010000000082111561256557600080fd5b90830190602082018581111561257a57600080fd5b825186602082028301116401000000008211171561259757600080fd5b82525081516020918201928201910280838360005b838110156125c45781810151838201526020016125ac565b50505050905001604052505050905090565b6001546001600160a01b0316331461261f5760405162461bcd60e51b81526004018080602001828103825260358152602001806146256035913960400191505060405180910390fd5b600054600154604080516001600160a01b03938416815292909116602083015280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a160018054600080546001600160a01b03199081166001600160a01b03841617909155169055565b600061269c6137ed565b6001600160a01b031663835e119c836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610dd757600080fd5b60006126e96137ed565b6001600160a01b0316637b1001b78360006040518363ffffffff1660e01b815260040180838152602001821515151581526020019250505060206040518083038186803b158015610dd757600080fd5b612741613955565b6001600160a01b0316637c3125416040518163ffffffff1660e01b815260040160006040518083038186803b15801561277957600080fd5b505afa15801561278d573d6000803e3d6000fd5b50505050612799613835565b6127a16137ed565b60048054604080516285c0d160e31b81526001600160a01b0392831693810193909352602483018590525192169163042e06889160448082019260009290919082900301818387803b15801561195a57600080fd5b6000546001600160a01b031681565b60008584612811613955565b6001600160a01b0316637118d4316040518163ffffffff1660e01b815260040160006040518083038186803b15801561284957600080fd5b505afa15801561285d573d6000803e3d6000fd5b50505050612869613955565b6001600160a01b0316636132eba483836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b1580156128b457600080fd5b505afa1580156128c8573d6000803e3d6000fd5b505050506128d4613835565b6128dc6139a5565b6004805460408051636fffe53b60e11b81526001600160a01b038e8116948201949094529183166024830152604482018c9052606482018b9052608482018a905288831660a483015260c4820188905251929091169163dfffca769160e4808201926020929091908290030181600087803b15801561295a57600080fd5b505af115801561296e573d6000803e3d6000fd5b505050506040513d602081101561298457600080fd5b50519998505050505050505050565b631cd554d160e21b81565b6007805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610e915780601f10610e6657610100808354040283529160200191610e91565b612a01613b8f565b600280546001600160a01b0383166001600160a01b0319909116811790915560408051918252517ffc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e9181900360200190a150565b612a5d613955565b6001600160a01b0316637c3125416040518163ffffffff1660e01b815260040160006040518083038186803b158015612a9557600080fd5b505afa158015612aa9573d6000803e3d6000fd5b50505050612ab5613835565b612abd6137ed565b60048054604080516324beb82560e11b81526001600160a01b03928316938101939093525192169163497d704a9160248082019260009290919082900301818387803b158015612b0c57600080fd5b505af1158015612b20573d6000803e3d6000fd5b505050505b565b6000806000612b34613835565b612b3c6139a5565b60048054604080516306c5a00b60e21b81526001600160a01b03928316938101939093526024830188905251921691631b16802c916044808201926060929091908290030181600087803b158015612b9357600080fd5b505af1158015611134573d6000803e3d6000fd5b6003546001600160a01b031681565b612bbe613f43565b600580546001600160a01b0319166001600160a01b03831617905561197381613ff3565b6000612bec6137ed565b6001600160a01b031663a311c7c2836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015610dd757600080fd5b6000612c4b6137ed565b6001600160a01b031663a5fdc5de836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015610dd757600080fd5b6000612caa613835565b612cb2613955565b6001600160a01b031663086dabd16040518163ffffffff1660e01b815260040160006040518083038186803b158015612cea57600080fd5b505afa158015612cfe573d6000803e3d6000fd5b5050600454612d1992506001600160a01b0316905083613bd8565b50600454610f28906001600160a01b0316848461411e565b612d39614605565b60005b600b54811015612d7757600b8181548110612d5357fe5b9060005260206000200154828260188110612d6a57fe5b6020020152600101612d3c565b5090565b612d836139a5565b6001600160a01b0316336001600160a01b031614612dd6576040805162461bcd60e51b815260206004820152601e6024820152600080516020614682833981519152604482015290519081900360640190fd5b6002805460408051602081018690528082018590528151808203830181526060909101918290526001600160a01b039092169263907dff97929180602861465a82396028019050604051809103902061239488613f37565b612e36613955565b6001600160a01b0316637c3125416040518163ffffffff1660e01b815260040160006040518083038186803b158015612e6e57600080fd5b505afa158015612e82573d6000803e3d6000fd5b50505050612e8e613835565b612e966137ed565b600480546040805163644bb89960e11b81526001600160a01b03928316938101939093525192169163c89771329160248082019260009290919082900301818387803b158015612b0c57600080fd5b612eed61412b565b600480546001600160a01b0319166001600160a01b0392909216919091179055565b612f17613955565b6001600160a01b0316637c3125416040518163ffffffff1660e01b815260040160006040518083038186803b158015612f4f57600080fd5b505afa158015612f63573d6000803e3d6000fd5b50505050612f6f613835565b612f776137ed565b6004805460408051632694552d60e21b81526001600160a01b0387811694820194909452918316602483015260448201859052519290911691639a5154b49160648082019260009290919082900301818387803b158015612fd757600080fd5b505af1158015612feb573d6000803e3d6000fd5b505050505050565b600b818154811061300057fe5b600091825260209091200154905081565b6000838261301d613955565b6001600160a01b0316637118d4316040518163ffffffff1660e01b815260040160006040518083038186803b15801561305557600080fd5b505afa158015613069573d6000803e3d6000fd5b50505050613075613955565b6001600160a01b0316636132eba483836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b1580156130c057600080fd5b505afa1580156130d4573d6000803e3d6000fd5b505050506130e0613835565b6130e86139a5565b6004805460408051630d4388eb60e31b81526001600160a01b038c8116948201949094529183166024830152604482018a90526064820189905260848201889052519290911691636a1c47589160a4808201926020929091908290030181600087803b15801561315757600080fd5b505af115801561316b573d6000803e3d6000fd5b505050506040513d602081101561318157600080fd5b5051979650505050505050565b60006131986137ed565b6001600160a01b031663d37c4d8b84846040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b031681526020018281526020019250505060206040518083038186803b1580156131f557600080fd5b505afa158015613209573d6000803e3d6000fd5b505050506040513d602081101561321f57600080fd5b50519392505050565b60006132326137ed565b6001600160a01b0316637b1001b78360016040518363ffffffff1660e01b815260040180838152602001821515151581526020019250505060206040518083038186803b158015610dd757600080fd5b6004546001600160a01b031681565b600061329b6137ed565b6001600160a01b031663dbf633406040518163ffffffff1660e01b815260040160206040518083038186803b158015611e7457600080fd5b60055460408051636eb1769f60e11b81526001600160a01b03858116600483015284811660248301529151600093929092169163dd62ed3e91604480820192602092909190829003018186803b1580156131f557600080fd5b6133346139a5565b6001600160a01b0316336001600160a01b031614613387576040805162461bcd60e51b815260206004820152601e6024820152600080516020614682833981519152604482015290519081900360640190fd5b6002805460408051602081018690528082018590528151808203830181526060909101918290526001600160a01b039092169263907dff9792918060296147a0823960290190506040518091039020876000806040518763ffffffff1660e01b815260040180806020018781526020018681526020018581526020018460001b81526020018360001b815260200182810382528881815181526020019150805190602001908083836000838110156124075781810151838201526020016123ef565b601881565b6000613458613955565b6001600160a01b031663086dabd16040518163ffffffff1660e01b815260040160006040518083038186803b15801561349057600080fd5b505afa1580156134a4573d6000803e3d6000fd5b505050506134b0613835565b6000806134bb6137ed565b600480546040805163298f137d60e21b81526001600160a01b038a8116948201949094526024810189905291831660448301528051939092169263a63c4df4926064808401939192918290030181600087803b15801561351a57600080fd5b505af115801561352e573d6000803e3d6000fd5b505050506040513d604081101561354457600080fd5b508051602090910151600454919350915061356d908690849084906001600160a01b031661419f565b6004546135859086906001600160a01b03168461411e565b95945050505050565b613596613955565b6001600160a01b0316637c3125416040518163ffffffff1660e01b815260040160006040518083038186803b1580156135ce57600080fd5b505afa1580156135e2573d6000803e3d6000fd5b505050506135ee613835565b6135f66137ed565b600480546040805163227635b160e11b81526001600160a01b03878116948201949094529183166024830152604482018590525192909116916344ec6b629160648082019260009290919082900301818387803b158015612fd757600080fd5b6005546001600160a01b031681565b6002546001600160a01b031681565b60008382613680613955565b6001600160a01b0316637118d4316040518163ffffffff1660e01b815260040160006040518083038186803b1580156136b857600080fd5b505afa1580156136cc573d6000803e3d6000fd5b505050506136d8613955565b6001600160a01b0316636132eba483836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561372357600080fd5b505afa158015613737573d6000803e3d6000fd5b50505050613743613835565b61374b6139a5565b6004805460408051630a1e187d60e01b81526001600160a01b03928316938101849052602481018b9052604481018a905260648101899052608481019390935251921691630a1e187d9160a4808201926020929091908290030181600087803b1580156137b757600080fd5b505af11580156137cb573d6000803e3d6000fd5b505050506040513d60208110156137e157600080fd5b50519695505050505050565b60006138306524b9b9bab2b960d11b604051806040016040528060168152602001754d697373696e6720497373756572206164647265737360501b8152506142e6565b905090565b6002546001600160a01b0316331480159061385b57506003546001600160a01b03163314155b801561387257506004546001600160a01b03163314155b15612b2557600480546001600160a01b03191633179055565b60025460408051602080820185905282518083039091018152908201918290526001600160a01b039092169163907dff97916003908060216147358239602101905060405180910390206138de88613f37565b6138e788613f37565b60006040518763ffffffff1660e01b815260040180806020018781526020018681526020018581526020018481526020018360001b815260200182810382528881815181526020019150805190602001908083836000838110156124075781810151838201526020016123ef565b60006138306b53797374656d53746174757360a01b6040518060400160405280601c81526020017f4d697373696e672053797374656d5374617475732061646472657373000000008152506142e6565b60006138306822bc31b430b733b2b960b91b6040518060400160405280601981526020017f4d697373696e672045786368616e6765722061646472657373000000000000008152506142e6565b6000613830722932bbb0b93239a234b9ba3934b13aba34b7b760691b604051806060016040528060238152602001614756602391396142e6565b60006138306d537570706c795363686564756c6560901b6040518060400160405280601e81526020017f4d697373696e6720537570706c795363686564756c65206164647265737300008152506142e6565b600082821115613ad5576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082820183811015613b35576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b60025460408051602080820185905282518083039091018152908201918290526001600160a01b039092169163907dff97916003908060216147fb8239602101905060405180910390206138de88613f37565b6000546001600160a01b03163314612b255760405162461bcd60e51b815260040180806020018281038252602f815260200180614706602f913960400191505060405180910390fd5b600080613be3614390565b60408051631167f01160e31b81526001600160a01b0387811660048301528251931692638b3f808892602480840193919291829003018186803b158015613c2957600080fd5b505afa158015613c3d573d6000803e3d6000fd5b505050506040513d6040811015613c5357600080fd5b505190508015610f2857600080613c686137ed565b600554604080516370a0823160e01b81526001600160a01b038a81166004830152915193821693636bed0415938b9316916370a08231916024808301926020929190829003018186803b158015613cbe57600080fd5b505afa158015613cd2573d6000803e3d6000fd5b505050506040513d6020811015613ce857600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b03909316600484015260248301919091528051604480840193829003018186803b158015613d3357600080fd5b505afa158015613d47573d6000803e3d6000fd5b505050506040513d6040811015613d5d57600080fd5b508051602090910151909250905081851115613daa5760405162461bcd60e51b81526004018080602001828103825260268152602001806146e06026913960400191505060405180910390fd5b8015613dfd576040805162461bcd60e51b815260206004820152601e60248201527f412073796e7468206f7220534e58207261746520697320696e76616c69640000604482015290519081900360640190fd5b50600195945050505050565b60055460408051636eb1769f60e11b81526001600160a01b03868116600483015287811660248301529151600093929092169163da46098c9187918991613eac918891879163dd62ed3e91604480820192602092909190829003018186803b158015613e7457600080fd5b505afa158015613e88573d6000803e3d6000fd5b505050506040513d6020811015613e9e57600080fd5b50519063ffffffff613a7e16565b6040518463ffffffff1660e01b815260040180846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b031681526020018281526020019350505050600060405180830381600087803b158015613f1457600080fd5b505af1158015613f28573d6000803e3d6000fd5b505050506135858484846143e2565b6001600160a01b031690565b6002546001600160a01b03163314801590613f6957506003546001600160a01b03163314155b8015613f8057506004546001600160a01b03163314155b15613f9857600480546001600160a01b031916331790555b6000546004546001600160a01b03908116911614612b25576040805162461bcd60e51b815260206004820152601360248201527227bbb732b91037b7363c90333ab731ba34b7b760691b604482015290519081900360640190fd5b600254604080516001600160a01b038481166020808401919091528351808403820181528385018086527f546f6b656e5374617465557064617465642861646472657373290000000000009052935192839003605a01832063907dff9760e01b8452600160248501819052604485018290526000606486018190526084860181905260a4860181905260c060048701908152875160c48801528751959098169763907dff97979692959394919384938493839260e490920191908a0190808383885b838110156140cd5781810151838201526020016140b5565b50505050905090810190601f1680156140fa5780820380516001836020036101000a031916815260200191505b50975050505050505050600060405180830381600087803b15801561195a57600080fd5b60006118948484846143e2565b6002546001600160a01b031633148061414e57506003546001600160a01b031633145b612b25576040805162461bcd60e51b815260206004820152601760248201527f4f6e6c79207468652070726f78792063616e2063616c6c000000000000000000604482015290519081900360640190fd5b6002805460408051602081018790528082018690526001600160a01b03858116606080840191909152835180840390910181526080909201928390529092169263907dff9792918060326147c982396032019050604051809103902061420489613f37565b6000806040518763ffffffff1660e01b815260040180806020018781526020018681526020018581526020018460001b81526020018360001b8152602001828103825288818151815260200191508051906020019080838360005b8381101561427757818101518382015260200161425f565b50505050905090810190601f1680156142a45780820380516001836020036101000a031916815260200191505b50975050505050505050600060405180830381600087803b1580156142c857600080fd5b505af11580156142dc573d6000803e3d6000fd5b5050505050505050565b6000828152600a60205260408120546001600160a01b031682816143885760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561434d578181015183820152602001614335565b50505050905090810190601f16801561437a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b509392505050565b60006138306d53796e746865746978537461746560901b6040518060400160405280601e81526020017f4d697373696e672053796e7468657469785374617465206164647265737300008152506142e6565b60006001600160a01b0383161580159061440557506001600160a01b0383163014155b801561441f57506002546001600160a01b03848116911614155b614470576040805162461bcd60e51b815260206004820152601f60248201527f43616e6e6f74207472616e7366657220746f2074686973206164647265737300604482015290519081900360640190fd5b600554604080516370a0823160e01b81526001600160a01b0387811660048301529151919092169163b46310f69187916144ce91879186916370a0823191602480820192602092909190829003018186803b158015613e7457600080fd5b6040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561451d57600080fd5b505af1158015614531573d6000803e3d6000fd5b5050600554604080516370a0823160e01b81526001600160a01b038881166004830152915191909216935063b46310f69250869161459391879186916370a0823191602480820192602092909190829003018186803b15801561144c57600080fd5b6040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156145e257600080fd5b505af11580156145f6573d6000803e3d6000fd5b50505050610f28848484613b3c565b604051806103000160405280601890602082028038833950919291505056fe596f75206d757374206265206e6f6d696e61746564206265666f726520796f752063616e20616363657074206f776e65727368697045786368616e67655265636c61696d28616464726573732c627974657333322c75696e74323536294f6e6c792045786368616e6765722063616e20696e766f6b652074686973000053796e746845786368616e676528616464726573732c627974657333322c75696e743235362c627974657333322c75696e743235362c616464726573732943616e6e6f74207472616e73666572207374616b6564206f7220657363726f77656420534e584f6e6c792074686520636f6e7472616374206f776e6572206d617920706572666f726d207468697320616374696f6e417070726f76616c28616464726573732c616464726573732c75696e74323536294d697373696e672052657761726473446973747269627574696f6e206164647265737345786368616e676552656261746528616464726573732c627974657333322c75696e743235362945786368616e6765547261636b696e6728627974657333322c627974657333322c75696e74323536294163636f756e744c69717569646174656428616464726573732c75696e743235362c75696e743235362c61646472657373295472616e7366657228616464726573732c616464726573732c75696e7432353629a265627a7a72315820c207630f2d713d291f23cacaf7b4436e75c261176db3c6017b036afe5e0ea7e164736f6c63430005100032"

// DeploySynthetix deploys a new Ethereum contract, binding an instance of Synthetix to it.
func DeploySynthetix(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address, _tokenState common.Address, _owner common.Address, _totalSupply *big.Int, _resolver common.Address) (common.Address, *types.Transaction, *Synthetix, error) {
	parsed, err := abi.JSON(strings.NewReader(SynthetixABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SynthetixBin), backend, _proxy, _tokenState, _owner, _totalSupply, _resolver)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Synthetix{SynthetixCaller: SynthetixCaller{contract: contract}, SynthetixTransactor: SynthetixTransactor{contract: contract}, SynthetixFilterer: SynthetixFilterer{contract: contract}}, nil
}

// Synthetix is an auto generated Go binding around an Ethereum contract.
type Synthetix struct {
	SynthetixCaller     // Read-only binding to the contract
	SynthetixTransactor // Write-only binding to the contract
	SynthetixFilterer   // Log filterer for contract events
}

// SynthetixCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynthetixCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynthetixTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynthetixFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynthetixSession struct {
	Contract     *Synthetix        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SynthetixCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynthetixCallerSession struct {
	Contract *SynthetixCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SynthetixTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynthetixTransactorSession struct {
	Contract     *SynthetixTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SynthetixRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynthetixRaw struct {
	Contract *Synthetix // Generic contract binding to access the raw methods on
}

// SynthetixCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynthetixCallerRaw struct {
	Contract *SynthetixCaller // Generic read-only contract binding to access the raw methods on
}

// SynthetixTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynthetixTransactorRaw struct {
	Contract *SynthetixTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynthetix creates a new instance of Synthetix, bound to a specific deployed contract.
func NewSynthetix(address common.Address, backend bind.ContractBackend) (*Synthetix, error) {
	contract, err := bindSynthetix(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Synthetix{SynthetixCaller: SynthetixCaller{contract: contract}, SynthetixTransactor: SynthetixTransactor{contract: contract}, SynthetixFilterer: SynthetixFilterer{contract: contract}}, nil
}

// NewSynthetixCaller creates a new read-only instance of Synthetix, bound to a specific deployed contract.
func NewSynthetixCaller(address common.Address, caller bind.ContractCaller) (*SynthetixCaller, error) {
	contract, err := bindSynthetix(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynthetixCaller{contract: contract}, nil
}

// NewSynthetixTransactor creates a new write-only instance of Synthetix, bound to a specific deployed contract.
func NewSynthetixTransactor(address common.Address, transactor bind.ContractTransactor) (*SynthetixTransactor, error) {
	contract, err := bindSynthetix(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynthetixTransactor{contract: contract}, nil
}

// NewSynthetixFilterer creates a new log filterer instance of Synthetix, bound to a specific deployed contract.
func NewSynthetixFilterer(address common.Address, filterer bind.ContractFilterer) (*SynthetixFilterer, error) {
	contract, err := bindSynthetix(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynthetixFilterer{contract: contract}, nil
}

// bindSynthetix binds a generic wrapper to an already deployed contract.
func bindSynthetix(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SynthetixABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Synthetix *SynthetixRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Synthetix.Contract.SynthetixCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Synthetix *SynthetixRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Synthetix.Contract.SynthetixTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Synthetix *SynthetixRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Synthetix.Contract.SynthetixTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Synthetix *SynthetixCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Synthetix.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Synthetix *SynthetixTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Synthetix.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Synthetix *SynthetixTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Synthetix.Contract.contract.Transact(opts, method, params...)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_Synthetix *SynthetixCaller) DECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_Synthetix *SynthetixSession) DECIMALS() (uint8, error) {
	return _Synthetix.Contract.DECIMALS(&_Synthetix.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_Synthetix *SynthetixCallerSession) DECIMALS() (uint8, error) {
	return _Synthetix.Contract.DECIMALS(&_Synthetix.CallOpts)
}

// MAXADDRESSESFROMRESOLVER is a free data retrieval call binding the contract method 0xe3235c91.
//
// Solidity: function MAX_ADDRESSES_FROM_RESOLVER() view returns(uint256)
func (_Synthetix *SynthetixCaller) MAXADDRESSESFROMRESOLVER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "MAX_ADDRESSES_FROM_RESOLVER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXADDRESSESFROMRESOLVER is a free data retrieval call binding the contract method 0xe3235c91.
//
// Solidity: function MAX_ADDRESSES_FROM_RESOLVER() view returns(uint256)
func (_Synthetix *SynthetixSession) MAXADDRESSESFROMRESOLVER() (*big.Int, error) {
	return _Synthetix.Contract.MAXADDRESSESFROMRESOLVER(&_Synthetix.CallOpts)
}

// MAXADDRESSESFROMRESOLVER is a free data retrieval call binding the contract method 0xe3235c91.
//
// Solidity: function MAX_ADDRESSES_FROM_RESOLVER() view returns(uint256)
func (_Synthetix *SynthetixCallerSession) MAXADDRESSESFROMRESOLVER() (*big.Int, error) {
	return _Synthetix.Contract.MAXADDRESSESFROMRESOLVER(&_Synthetix.CallOpts)
}

// TOKENNAME is a free data retrieval call binding the contract method 0x18821400.
//
// Solidity: function TOKEN_NAME() view returns(string)
func (_Synthetix *SynthetixCaller) TOKENNAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "TOKEN_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TOKENNAME is a free data retrieval call binding the contract method 0x18821400.
//
// Solidity: function TOKEN_NAME() view returns(string)
func (_Synthetix *SynthetixSession) TOKENNAME() (string, error) {
	return _Synthetix.Contract.TOKENNAME(&_Synthetix.CallOpts)
}

// TOKENNAME is a free data retrieval call binding the contract method 0x18821400.
//
// Solidity: function TOKEN_NAME() view returns(string)
func (_Synthetix *SynthetixCallerSession) TOKENNAME() (string, error) {
	return _Synthetix.Contract.TOKENNAME(&_Synthetix.CallOpts)
}

// TOKENSYMBOL is a free data retrieval call binding the contract method 0x2a905318.
//
// Solidity: function TOKEN_SYMBOL() view returns(string)
func (_Synthetix *SynthetixCaller) TOKENSYMBOL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "TOKEN_SYMBOL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TOKENSYMBOL is a free data retrieval call binding the contract method 0x2a905318.
//
// Solidity: function TOKEN_SYMBOL() view returns(string)
func (_Synthetix *SynthetixSession) TOKENSYMBOL() (string, error) {
	return _Synthetix.Contract.TOKENSYMBOL(&_Synthetix.CallOpts)
}

// TOKENSYMBOL is a free data retrieval call binding the contract method 0x2a905318.
//
// Solidity: function TOKEN_SYMBOL() view returns(string)
func (_Synthetix *SynthetixCallerSession) TOKENSYMBOL() (string, error) {
	return _Synthetix.Contract.TOKENSYMBOL(&_Synthetix.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Synthetix *SynthetixCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Synthetix *SynthetixSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Synthetix.Contract.Allowance(&_Synthetix.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Synthetix *SynthetixCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Synthetix.Contract.Allowance(&_Synthetix.CallOpts, owner, spender)
}

// AnySynthOrSNXRateIsInvalid is a free data retrieval call binding the contract method 0x4e99bda9.
//
// Solidity: function anySynthOrSNXRateIsInvalid() view returns(bool anyRateInvalid)
func (_Synthetix *SynthetixCaller) AnySynthOrSNXRateIsInvalid(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "anySynthOrSNXRateIsInvalid")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AnySynthOrSNXRateIsInvalid is a free data retrieval call binding the contract method 0x4e99bda9.
//
// Solidity: function anySynthOrSNXRateIsInvalid() view returns(bool anyRateInvalid)
func (_Synthetix *SynthetixSession) AnySynthOrSNXRateIsInvalid() (bool, error) {
	return _Synthetix.Contract.AnySynthOrSNXRateIsInvalid(&_Synthetix.CallOpts)
}

// AnySynthOrSNXRateIsInvalid is a free data retrieval call binding the contract method 0x4e99bda9.
//
// Solidity: function anySynthOrSNXRateIsInvalid() view returns(bool anyRateInvalid)
func (_Synthetix *SynthetixCallerSession) AnySynthOrSNXRateIsInvalid() (bool, error) {
	return _Synthetix.Contract.AnySynthOrSNXRateIsInvalid(&_Synthetix.CallOpts)
}

// AvailableCurrencyKeys is a free data retrieval call binding the contract method 0x72cb051f.
//
// Solidity: function availableCurrencyKeys() view returns(bytes32[])
func (_Synthetix *SynthetixCaller) AvailableCurrencyKeys(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "availableCurrencyKeys")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// AvailableCurrencyKeys is a free data retrieval call binding the contract method 0x72cb051f.
//
// Solidity: function availableCurrencyKeys() view returns(bytes32[])
func (_Synthetix *SynthetixSession) AvailableCurrencyKeys() ([][32]byte, error) {
	return _Synthetix.Contract.AvailableCurrencyKeys(&_Synthetix.CallOpts)
}

// AvailableCurrencyKeys is a free data retrieval call binding the contract method 0x72cb051f.
//
// Solidity: function availableCurrencyKeys() view returns(bytes32[])
func (_Synthetix *SynthetixCallerSession) AvailableCurrencyKeys() ([][32]byte, error) {
	return _Synthetix.Contract.AvailableCurrencyKeys(&_Synthetix.CallOpts)
}

// AvailableSynthCount is a free data retrieval call binding the contract method 0xdbf63340.
//
// Solidity: function availableSynthCount() view returns(uint256)
func (_Synthetix *SynthetixCaller) AvailableSynthCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "availableSynthCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableSynthCount is a free data retrieval call binding the contract method 0xdbf63340.
//
// Solidity: function availableSynthCount() view returns(uint256)
func (_Synthetix *SynthetixSession) AvailableSynthCount() (*big.Int, error) {
	return _Synthetix.Contract.AvailableSynthCount(&_Synthetix.CallOpts)
}

// AvailableSynthCount is a free data retrieval call binding the contract method 0xdbf63340.
//
// Solidity: function availableSynthCount() view returns(uint256)
func (_Synthetix *SynthetixCallerSession) AvailableSynthCount() (*big.Int, error) {
	return _Synthetix.Contract.AvailableSynthCount(&_Synthetix.CallOpts)
}

// AvailableSynths is a free data retrieval call binding the contract method 0x835e119c.
//
// Solidity: function availableSynths(uint256 index) view returns(address)
func (_Synthetix *SynthetixCaller) AvailableSynths(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "availableSynths", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvailableSynths is a free data retrieval call binding the contract method 0x835e119c.
//
// Solidity: function availableSynths(uint256 index) view returns(address)
func (_Synthetix *SynthetixSession) AvailableSynths(index *big.Int) (common.Address, error) {
	return _Synthetix.Contract.AvailableSynths(&_Synthetix.CallOpts, index)
}

// AvailableSynths is a free data retrieval call binding the contract method 0x835e119c.
//
// Solidity: function availableSynths(uint256 index) view returns(address)
func (_Synthetix *SynthetixCallerSession) AvailableSynths(index *big.Int) (common.Address, error) {
	return _Synthetix.Contract.AvailableSynths(&_Synthetix.CallOpts, index)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Synthetix *SynthetixCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Synthetix *SynthetixSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Synthetix.Contract.BalanceOf(&_Synthetix.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Synthetix *SynthetixCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Synthetix.Contract.BalanceOf(&_Synthetix.CallOpts, account)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address account) view returns(uint256)
func (_Synthetix *SynthetixCaller) Collateral(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "collateral", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address account) view returns(uint256)
func (_Synthetix *SynthetixSession) Collateral(account common.Address) (*big.Int, error) {
	return _Synthetix.Contract.Collateral(&_Synthetix.CallOpts, account)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address account) view returns(uint256)
func (_Synthetix *SynthetixCallerSession) Collateral(account common.Address) (*big.Int, error) {
	return _Synthetix.Contract.Collateral(&_Synthetix.CallOpts, account)
}

// CollateralisationRatio is a free data retrieval call binding the contract method 0xa311c7c2.
//
// Solidity: function collateralisationRatio(address _issuer) view returns(uint256)
func (_Synthetix *SynthetixCaller) CollateralisationRatio(opts *bind.CallOpts, _issuer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "collateralisationRatio", _issuer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralisationRatio is a free data retrieval call binding the contract method 0xa311c7c2.
//
// Solidity: function collateralisationRatio(address _issuer) view returns(uint256)
func (_Synthetix *SynthetixSession) CollateralisationRatio(_issuer common.Address) (*big.Int, error) {
	return _Synthetix.Contract.CollateralisationRatio(&_Synthetix.CallOpts, _issuer)
}

// CollateralisationRatio is a free data retrieval call binding the contract method 0xa311c7c2.
//
// Solidity: function collateralisationRatio(address _issuer) view returns(uint256)
func (_Synthetix *SynthetixCallerSession) CollateralisationRatio(_issuer common.Address) (*big.Int, error) {
	return _Synthetix.Contract.CollateralisationRatio(&_Synthetix.CallOpts, _issuer)
}

// DebtBalanceOf is a free data retrieval call binding the contract method 0xd37c4d8b.
//
// Solidity: function debtBalanceOf(address account, bytes32 currencyKey) view returns(uint256)
func (_Synthetix *SynthetixCaller) DebtBalanceOf(opts *bind.CallOpts, account common.Address, currencyKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "debtBalanceOf", account, currencyKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtBalanceOf is a free data retrieval call binding the contract method 0xd37c4d8b.
//
// Solidity: function debtBalanceOf(address account, bytes32 currencyKey) view returns(uint256)
func (_Synthetix *SynthetixSession) DebtBalanceOf(account common.Address, currencyKey [32]byte) (*big.Int, error) {
	return _Synthetix.Contract.DebtBalanceOf(&_Synthetix.CallOpts, account, currencyKey)
}

// DebtBalanceOf is a free data retrieval call binding the contract method 0xd37c4d8b.
//
// Solidity: function debtBalanceOf(address account, bytes32 currencyKey) view returns(uint256)
func (_Synthetix *SynthetixCallerSession) DebtBalanceOf(account common.Address, currencyKey [32]byte) (*big.Int, error) {
	return _Synthetix.Contract.DebtBalanceOf(&_Synthetix.CallOpts, account, currencyKey)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Synthetix *SynthetixCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Synthetix *SynthetixSession) Decimals() (uint8, error) {
	return _Synthetix.Contract.Decimals(&_Synthetix.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Synthetix *SynthetixCallerSession) Decimals() (uint8, error) {
	return _Synthetix.Contract.Decimals(&_Synthetix.CallOpts)
}

// GetResolverAddressesRequired is a free data retrieval call binding the contract method 0xab49848c.
//
// Solidity: function getResolverAddressesRequired() view returns(bytes32[24] addressesRequired)
func (_Synthetix *SynthetixCaller) GetResolverAddressesRequired(opts *bind.CallOpts) ([24][32]byte, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "getResolverAddressesRequired")

	if err != nil {
		return *new([24][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([24][32]byte)).(*[24][32]byte)

	return out0, err

}

// GetResolverAddressesRequired is a free data retrieval call binding the contract method 0xab49848c.
//
// Solidity: function getResolverAddressesRequired() view returns(bytes32[24] addressesRequired)
func (_Synthetix *SynthetixSession) GetResolverAddressesRequired() ([24][32]byte, error) {
	return _Synthetix.Contract.GetResolverAddressesRequired(&_Synthetix.CallOpts)
}

// GetResolverAddressesRequired is a free data retrieval call binding the contract method 0xab49848c.
//
// Solidity: function getResolverAddressesRequired() view returns(bytes32[24] addressesRequired)
func (_Synthetix *SynthetixCallerSession) GetResolverAddressesRequired() ([24][32]byte, error) {
	return _Synthetix.Contract.GetResolverAddressesRequired(&_Synthetix.CallOpts)
}

// IntegrationProxy is a free data retrieval call binding the contract method 0x9cbdaeb6.
//
// Solidity: function integrationProxy() view returns(address)
func (_Synthetix *SynthetixCaller) IntegrationProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "integrationProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IntegrationProxy is a free data retrieval call binding the contract method 0x9cbdaeb6.
//
// Solidity: function integrationProxy() view returns(address)
func (_Synthetix *SynthetixSession) IntegrationProxy() (common.Address, error) {
	return _Synthetix.Contract.IntegrationProxy(&_Synthetix.CallOpts)
}

// IntegrationProxy is a free data retrieval call binding the contract method 0x9cbdaeb6.
//
// Solidity: function integrationProxy() view returns(address)
func (_Synthetix *SynthetixCallerSession) IntegrationProxy() (common.Address, error) {
	return _Synthetix.Contract.IntegrationProxy(&_Synthetix.CallOpts)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x631e1444.
//
// Solidity: function isResolverCached(address _resolver) view returns(bool)
func (_Synthetix *SynthetixCaller) IsResolverCached(opts *bind.CallOpts, _resolver common.Address) (bool, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "isResolverCached", _resolver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsResolverCached is a free data retrieval call binding the contract method 0x631e1444.
//
// Solidity: function isResolverCached(address _resolver) view returns(bool)
func (_Synthetix *SynthetixSession) IsResolverCached(_resolver common.Address) (bool, error) {
	return _Synthetix.Contract.IsResolverCached(&_Synthetix.CallOpts, _resolver)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x631e1444.
//
// Solidity: function isResolverCached(address _resolver) view returns(bool)
func (_Synthetix *SynthetixCallerSession) IsResolverCached(_resolver common.Address) (bool, error) {
	return _Synthetix.Contract.IsResolverCached(&_Synthetix.CallOpts, _resolver)
}

// IsWaitingPeriod is a free data retrieval call binding the contract method 0x1fce304d.
//
// Solidity: function isWaitingPeriod(bytes32 currencyKey) view returns(bool)
func (_Synthetix *SynthetixCaller) IsWaitingPeriod(opts *bind.CallOpts, currencyKey [32]byte) (bool, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "isWaitingPeriod", currencyKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWaitingPeriod is a free data retrieval call binding the contract method 0x1fce304d.
//
// Solidity: function isWaitingPeriod(bytes32 currencyKey) view returns(bool)
func (_Synthetix *SynthetixSession) IsWaitingPeriod(currencyKey [32]byte) (bool, error) {
	return _Synthetix.Contract.IsWaitingPeriod(&_Synthetix.CallOpts, currencyKey)
}

// IsWaitingPeriod is a free data retrieval call binding the contract method 0x1fce304d.
//
// Solidity: function isWaitingPeriod(bytes32 currencyKey) view returns(bool)
func (_Synthetix *SynthetixCallerSession) IsWaitingPeriod(currencyKey [32]byte) (bool, error) {
	return _Synthetix.Contract.IsWaitingPeriod(&_Synthetix.CallOpts, currencyKey)
}

// MaxIssuableSynths is a free data retrieval call binding the contract method 0x05b3c1c9.
//
// Solidity: function maxIssuableSynths(address account) view returns(uint256 maxIssuable)
func (_Synthetix *SynthetixCaller) MaxIssuableSynths(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "maxIssuableSynths", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxIssuableSynths is a free data retrieval call binding the contract method 0x05b3c1c9.
//
// Solidity: function maxIssuableSynths(address account) view returns(uint256 maxIssuable)
func (_Synthetix *SynthetixSession) MaxIssuableSynths(account common.Address) (*big.Int, error) {
	return _Synthetix.Contract.MaxIssuableSynths(&_Synthetix.CallOpts, account)
}

// MaxIssuableSynths is a free data retrieval call binding the contract method 0x05b3c1c9.
//
// Solidity: function maxIssuableSynths(address account) view returns(uint256 maxIssuable)
func (_Synthetix *SynthetixCallerSession) MaxIssuableSynths(account common.Address) (*big.Int, error) {
	return _Synthetix.Contract.MaxIssuableSynths(&_Synthetix.CallOpts, account)
}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_Synthetix *SynthetixCaller) MessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "messageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_Synthetix *SynthetixSession) MessageSender() (common.Address, error) {
	return _Synthetix.Contract.MessageSender(&_Synthetix.CallOpts)
}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_Synthetix *SynthetixCallerSession) MessageSender() (common.Address, error) {
	return _Synthetix.Contract.MessageSender(&_Synthetix.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Synthetix *SynthetixCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Synthetix *SynthetixSession) Name() (string, error) {
	return _Synthetix.Contract.Name(&_Synthetix.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Synthetix *SynthetixCallerSession) Name() (string, error) {
	return _Synthetix.Contract.Name(&_Synthetix.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Synthetix *SynthetixCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Synthetix *SynthetixSession) NominatedOwner() (common.Address, error) {
	return _Synthetix.Contract.NominatedOwner(&_Synthetix.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Synthetix *SynthetixCallerSession) NominatedOwner() (common.Address, error) {
	return _Synthetix.Contract.NominatedOwner(&_Synthetix.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Synthetix *SynthetixCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Synthetix *SynthetixSession) Owner() (common.Address, error) {
	return _Synthetix.Contract.Owner(&_Synthetix.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Synthetix *SynthetixCallerSession) Owner() (common.Address, error) {
	return _Synthetix.Contract.Owner(&_Synthetix.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_Synthetix *SynthetixCaller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "proxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_Synthetix *SynthetixSession) Proxy() (common.Address, error) {
	return _Synthetix.Contract.Proxy(&_Synthetix.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_Synthetix *SynthetixCallerSession) Proxy() (common.Address, error) {
	return _Synthetix.Contract.Proxy(&_Synthetix.CallOpts)
}

// RemainingIssuableSynths is a free data retrieval call binding the contract method 0x1137aedf.
//
// Solidity: function remainingIssuableSynths(address account) view returns(uint256 maxIssuable, uint256 alreadyIssued, uint256 totalSystemDebt)
func (_Synthetix *SynthetixCaller) RemainingIssuableSynths(opts *bind.CallOpts, account common.Address) (struct {
	MaxIssuable     *big.Int
	AlreadyIssued   *big.Int
	TotalSystemDebt *big.Int
}, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "remainingIssuableSynths", account)

	outstruct := new(struct {
		MaxIssuable     *big.Int
		AlreadyIssued   *big.Int
		TotalSystemDebt *big.Int
	})

	outstruct.MaxIssuable = out[0].(*big.Int)
	outstruct.AlreadyIssued = out[1].(*big.Int)
	outstruct.TotalSystemDebt = out[2].(*big.Int)

	return *outstruct, err

}

// RemainingIssuableSynths is a free data retrieval call binding the contract method 0x1137aedf.
//
// Solidity: function remainingIssuableSynths(address account) view returns(uint256 maxIssuable, uint256 alreadyIssued, uint256 totalSystemDebt)
func (_Synthetix *SynthetixSession) RemainingIssuableSynths(account common.Address) (struct {
	MaxIssuable     *big.Int
	AlreadyIssued   *big.Int
	TotalSystemDebt *big.Int
}, error) {
	return _Synthetix.Contract.RemainingIssuableSynths(&_Synthetix.CallOpts, account)
}

// RemainingIssuableSynths is a free data retrieval call binding the contract method 0x1137aedf.
//
// Solidity: function remainingIssuableSynths(address account) view returns(uint256 maxIssuable, uint256 alreadyIssued, uint256 totalSystemDebt)
func (_Synthetix *SynthetixCallerSession) RemainingIssuableSynths(account common.Address) (struct {
	MaxIssuable     *big.Int
	AlreadyIssued   *big.Int
	TotalSystemDebt *big.Int
}, error) {
	return _Synthetix.Contract.RemainingIssuableSynths(&_Synthetix.CallOpts, account)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Synthetix *SynthetixCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Synthetix *SynthetixSession) Resolver() (common.Address, error) {
	return _Synthetix.Contract.Resolver(&_Synthetix.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Synthetix *SynthetixCallerSession) Resolver() (common.Address, error) {
	return _Synthetix.Contract.Resolver(&_Synthetix.CallOpts)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0xc6c9d828.
//
// Solidity: function resolverAddressesRequired(uint256 ) view returns(bytes32)
func (_Synthetix *SynthetixCaller) ResolverAddressesRequired(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "resolverAddressesRequired", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0xc6c9d828.
//
// Solidity: function resolverAddressesRequired(uint256 ) view returns(bytes32)
func (_Synthetix *SynthetixSession) ResolverAddressesRequired(arg0 *big.Int) ([32]byte, error) {
	return _Synthetix.Contract.ResolverAddressesRequired(&_Synthetix.CallOpts, arg0)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0xc6c9d828.
//
// Solidity: function resolverAddressesRequired(uint256 ) view returns(bytes32)
func (_Synthetix *SynthetixCallerSession) ResolverAddressesRequired(arg0 *big.Int) ([32]byte, error) {
	return _Synthetix.Contract.ResolverAddressesRequired(&_Synthetix.CallOpts, arg0)
}

// SUSD is a free data retrieval call binding the contract method 0x9324cac7.
//
// Solidity: function sUSD() view returns(bytes32)
func (_Synthetix *SynthetixCaller) SUSD(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "sUSD")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUSD is a free data retrieval call binding the contract method 0x9324cac7.
//
// Solidity: function sUSD() view returns(bytes32)
func (_Synthetix *SynthetixSession) SUSD() ([32]byte, error) {
	return _Synthetix.Contract.SUSD(&_Synthetix.CallOpts)
}

// SUSD is a free data retrieval call binding the contract method 0x9324cac7.
//
// Solidity: function sUSD() view returns(bytes32)
func (_Synthetix *SynthetixCallerSession) SUSD() ([32]byte, error) {
	return _Synthetix.Contract.SUSD(&_Synthetix.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Synthetix *SynthetixCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Synthetix *SynthetixSession) Symbol() (string, error) {
	return _Synthetix.Contract.Symbol(&_Synthetix.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Synthetix *SynthetixCallerSession) Symbol() (string, error) {
	return _Synthetix.Contract.Symbol(&_Synthetix.CallOpts)
}

// Synths is a free data retrieval call binding the contract method 0x32608039.
//
// Solidity: function synths(bytes32 currencyKey) view returns(address)
func (_Synthetix *SynthetixCaller) Synths(opts *bind.CallOpts, currencyKey [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "synths", currencyKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Synths is a free data retrieval call binding the contract method 0x32608039.
//
// Solidity: function synths(bytes32 currencyKey) view returns(address)
func (_Synthetix *SynthetixSession) Synths(currencyKey [32]byte) (common.Address, error) {
	return _Synthetix.Contract.Synths(&_Synthetix.CallOpts, currencyKey)
}

// Synths is a free data retrieval call binding the contract method 0x32608039.
//
// Solidity: function synths(bytes32 currencyKey) view returns(address)
func (_Synthetix *SynthetixCallerSession) Synths(currencyKey [32]byte) (common.Address, error) {
	return _Synthetix.Contract.Synths(&_Synthetix.CallOpts, currencyKey)
}

// SynthsByAddress is a free data retrieval call binding the contract method 0x16b2213f.
//
// Solidity: function synthsByAddress(address synthAddress) view returns(bytes32)
func (_Synthetix *SynthetixCaller) SynthsByAddress(opts *bind.CallOpts, synthAddress common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "synthsByAddress", synthAddress)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SynthsByAddress is a free data retrieval call binding the contract method 0x16b2213f.
//
// Solidity: function synthsByAddress(address synthAddress) view returns(bytes32)
func (_Synthetix *SynthetixSession) SynthsByAddress(synthAddress common.Address) ([32]byte, error) {
	return _Synthetix.Contract.SynthsByAddress(&_Synthetix.CallOpts, synthAddress)
}

// SynthsByAddress is a free data retrieval call binding the contract method 0x16b2213f.
//
// Solidity: function synthsByAddress(address synthAddress) view returns(bytes32)
func (_Synthetix *SynthetixCallerSession) SynthsByAddress(synthAddress common.Address) ([32]byte, error) {
	return _Synthetix.Contract.SynthsByAddress(&_Synthetix.CallOpts, synthAddress)
}

// TokenState is a free data retrieval call binding the contract method 0xe90dd9e2.
//
// Solidity: function tokenState() view returns(address)
func (_Synthetix *SynthetixCaller) TokenState(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "tokenState")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenState is a free data retrieval call binding the contract method 0xe90dd9e2.
//
// Solidity: function tokenState() view returns(address)
func (_Synthetix *SynthetixSession) TokenState() (common.Address, error) {
	return _Synthetix.Contract.TokenState(&_Synthetix.CallOpts)
}

// TokenState is a free data retrieval call binding the contract method 0xe90dd9e2.
//
// Solidity: function tokenState() view returns(address)
func (_Synthetix *SynthetixCallerSession) TokenState() (common.Address, error) {
	return _Synthetix.Contract.TokenState(&_Synthetix.CallOpts)
}

// TotalIssuedSynths is a free data retrieval call binding the contract method 0x83d625d4.
//
// Solidity: function totalIssuedSynths(bytes32 currencyKey) view returns(uint256)
func (_Synthetix *SynthetixCaller) TotalIssuedSynths(opts *bind.CallOpts, currencyKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "totalIssuedSynths", currencyKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalIssuedSynths is a free data retrieval call binding the contract method 0x83d625d4.
//
// Solidity: function totalIssuedSynths(bytes32 currencyKey) view returns(uint256)
func (_Synthetix *SynthetixSession) TotalIssuedSynths(currencyKey [32]byte) (*big.Int, error) {
	return _Synthetix.Contract.TotalIssuedSynths(&_Synthetix.CallOpts, currencyKey)
}

// TotalIssuedSynths is a free data retrieval call binding the contract method 0x83d625d4.
//
// Solidity: function totalIssuedSynths(bytes32 currencyKey) view returns(uint256)
func (_Synthetix *SynthetixCallerSession) TotalIssuedSynths(currencyKey [32]byte) (*big.Int, error) {
	return _Synthetix.Contract.TotalIssuedSynths(&_Synthetix.CallOpts, currencyKey)
}

// TotalIssuedSynthsExcludeEtherCollateral is a free data retrieval call binding the contract method 0xd60888e4.
//
// Solidity: function totalIssuedSynthsExcludeEtherCollateral(bytes32 currencyKey) view returns(uint256)
func (_Synthetix *SynthetixCaller) TotalIssuedSynthsExcludeEtherCollateral(opts *bind.CallOpts, currencyKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "totalIssuedSynthsExcludeEtherCollateral", currencyKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalIssuedSynthsExcludeEtherCollateral is a free data retrieval call binding the contract method 0xd60888e4.
//
// Solidity: function totalIssuedSynthsExcludeEtherCollateral(bytes32 currencyKey) view returns(uint256)
func (_Synthetix *SynthetixSession) TotalIssuedSynthsExcludeEtherCollateral(currencyKey [32]byte) (*big.Int, error) {
	return _Synthetix.Contract.TotalIssuedSynthsExcludeEtherCollateral(&_Synthetix.CallOpts, currencyKey)
}

// TotalIssuedSynthsExcludeEtherCollateral is a free data retrieval call binding the contract method 0xd60888e4.
//
// Solidity: function totalIssuedSynthsExcludeEtherCollateral(bytes32 currencyKey) view returns(uint256)
func (_Synthetix *SynthetixCallerSession) TotalIssuedSynthsExcludeEtherCollateral(currencyKey [32]byte) (*big.Int, error) {
	return _Synthetix.Contract.TotalIssuedSynthsExcludeEtherCollateral(&_Synthetix.CallOpts, currencyKey)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Synthetix *SynthetixCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Synthetix *SynthetixSession) TotalSupply() (*big.Int, error) {
	return _Synthetix.Contract.TotalSupply(&_Synthetix.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Synthetix *SynthetixCallerSession) TotalSupply() (*big.Int, error) {
	return _Synthetix.Contract.TotalSupply(&_Synthetix.CallOpts)
}

// TransferableSynthetix is a free data retrieval call binding the contract method 0x6ac0bf9c.
//
// Solidity: function transferableSynthetix(address account) view returns(uint256 transferable)
func (_Synthetix *SynthetixCaller) TransferableSynthetix(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Synthetix.contract.Call(opts, &out, "transferableSynthetix", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferableSynthetix is a free data retrieval call binding the contract method 0x6ac0bf9c.
//
// Solidity: function transferableSynthetix(address account) view returns(uint256 transferable)
func (_Synthetix *SynthetixSession) TransferableSynthetix(account common.Address) (*big.Int, error) {
	return _Synthetix.Contract.TransferableSynthetix(&_Synthetix.CallOpts, account)
}

// TransferableSynthetix is a free data retrieval call binding the contract method 0x6ac0bf9c.
//
// Solidity: function transferableSynthetix(address account) view returns(uint256 transferable)
func (_Synthetix *SynthetixCallerSession) TransferableSynthetix(account common.Address) (*big.Int, error) {
	return _Synthetix.Contract.TransferableSynthetix(&_Synthetix.CallOpts, account)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Synthetix *SynthetixTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Synthetix *SynthetixSession) AcceptOwnership() (*types.Transaction, error) {
	return _Synthetix.Contract.AcceptOwnership(&_Synthetix.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Synthetix *SynthetixTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Synthetix.Contract.AcceptOwnership(&_Synthetix.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Synthetix *SynthetixTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Synthetix *SynthetixSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.Approve(&_Synthetix.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Synthetix *SynthetixTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.Approve(&_Synthetix.TransactOpts, spender, value)
}

// BurnSecondary is a paid mutator transaction binding the contract method 0xedef719a.
//
// Solidity: function burnSecondary(address , uint256 ) returns()
func (_Synthetix *SynthetixTransactor) BurnSecondary(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "burnSecondary", arg0, arg1)
}

// BurnSecondary is a paid mutator transaction binding the contract method 0xedef719a.
//
// Solidity: function burnSecondary(address , uint256 ) returns()
func (_Synthetix *SynthetixSession) BurnSecondary(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.BurnSecondary(&_Synthetix.TransactOpts, arg0, arg1)
}

// BurnSecondary is a paid mutator transaction binding the contract method 0xedef719a.
//
// Solidity: function burnSecondary(address , uint256 ) returns()
func (_Synthetix *SynthetixTransactorSession) BurnSecondary(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.BurnSecondary(&_Synthetix.TransactOpts, arg0, arg1)
}

// BurnSynths is a paid mutator transaction binding the contract method 0x295da87d.
//
// Solidity: function burnSynths(uint256 amount) returns()
func (_Synthetix *SynthetixTransactor) BurnSynths(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "burnSynths", amount)
}

// BurnSynths is a paid mutator transaction binding the contract method 0x295da87d.
//
// Solidity: function burnSynths(uint256 amount) returns()
func (_Synthetix *SynthetixSession) BurnSynths(amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.BurnSynths(&_Synthetix.TransactOpts, amount)
}

// BurnSynths is a paid mutator transaction binding the contract method 0x295da87d.
//
// Solidity: function burnSynths(uint256 amount) returns()
func (_Synthetix *SynthetixTransactorSession) BurnSynths(amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.BurnSynths(&_Synthetix.TransactOpts, amount)
}

// BurnSynthsOnBehalf is a paid mutator transaction binding the contract method 0xc2bf3880.
//
// Solidity: function burnSynthsOnBehalf(address burnForAddress, uint256 amount) returns()
func (_Synthetix *SynthetixTransactor) BurnSynthsOnBehalf(opts *bind.TransactOpts, burnForAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "burnSynthsOnBehalf", burnForAddress, amount)
}

// BurnSynthsOnBehalf is a paid mutator transaction binding the contract method 0xc2bf3880.
//
// Solidity: function burnSynthsOnBehalf(address burnForAddress, uint256 amount) returns()
func (_Synthetix *SynthetixSession) BurnSynthsOnBehalf(burnForAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.BurnSynthsOnBehalf(&_Synthetix.TransactOpts, burnForAddress, amount)
}

// BurnSynthsOnBehalf is a paid mutator transaction binding the contract method 0xc2bf3880.
//
// Solidity: function burnSynthsOnBehalf(address burnForAddress, uint256 amount) returns()
func (_Synthetix *SynthetixTransactorSession) BurnSynthsOnBehalf(burnForAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.BurnSynthsOnBehalf(&_Synthetix.TransactOpts, burnForAddress, amount)
}

// BurnSynthsToTarget is a paid mutator transaction binding the contract method 0x9741fb22.
//
// Solidity: function burnSynthsToTarget() returns()
func (_Synthetix *SynthetixTransactor) BurnSynthsToTarget(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "burnSynthsToTarget")
}

// BurnSynthsToTarget is a paid mutator transaction binding the contract method 0x9741fb22.
//
// Solidity: function burnSynthsToTarget() returns()
func (_Synthetix *SynthetixSession) BurnSynthsToTarget() (*types.Transaction, error) {
	return _Synthetix.Contract.BurnSynthsToTarget(&_Synthetix.TransactOpts)
}

// BurnSynthsToTarget is a paid mutator transaction binding the contract method 0x9741fb22.
//
// Solidity: function burnSynthsToTarget() returns()
func (_Synthetix *SynthetixTransactorSession) BurnSynthsToTarget() (*types.Transaction, error) {
	return _Synthetix.Contract.BurnSynthsToTarget(&_Synthetix.TransactOpts)
}

// BurnSynthsToTargetOnBehalf is a paid mutator transaction binding the contract method 0x2c955fa7.
//
// Solidity: function burnSynthsToTargetOnBehalf(address burnForAddress) returns()
func (_Synthetix *SynthetixTransactor) BurnSynthsToTargetOnBehalf(opts *bind.TransactOpts, burnForAddress common.Address) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "burnSynthsToTargetOnBehalf", burnForAddress)
}

// BurnSynthsToTargetOnBehalf is a paid mutator transaction binding the contract method 0x2c955fa7.
//
// Solidity: function burnSynthsToTargetOnBehalf(address burnForAddress) returns()
func (_Synthetix *SynthetixSession) BurnSynthsToTargetOnBehalf(burnForAddress common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.BurnSynthsToTargetOnBehalf(&_Synthetix.TransactOpts, burnForAddress)
}

// BurnSynthsToTargetOnBehalf is a paid mutator transaction binding the contract method 0x2c955fa7.
//
// Solidity: function burnSynthsToTargetOnBehalf(address burnForAddress) returns()
func (_Synthetix *SynthetixTransactorSession) BurnSynthsToTargetOnBehalf(burnForAddress common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.BurnSynthsToTargetOnBehalf(&_Synthetix.TransactOpts, burnForAddress)
}

// EmitExchangeRebate is a paid mutator transaction binding the contract method 0x6f01a986.
//
// Solidity: function emitExchangeRebate(address account, bytes32 currencyKey, uint256 amount) returns()
func (_Synthetix *SynthetixTransactor) EmitExchangeRebate(opts *bind.TransactOpts, account common.Address, currencyKey [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "emitExchangeRebate", account, currencyKey, amount)
}

// EmitExchangeRebate is a paid mutator transaction binding the contract method 0x6f01a986.
//
// Solidity: function emitExchangeRebate(address account, bytes32 currencyKey, uint256 amount) returns()
func (_Synthetix *SynthetixSession) EmitExchangeRebate(account common.Address, currencyKey [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.EmitExchangeRebate(&_Synthetix.TransactOpts, account, currencyKey, amount)
}

// EmitExchangeRebate is a paid mutator transaction binding the contract method 0x6f01a986.
//
// Solidity: function emitExchangeRebate(address account, bytes32 currencyKey, uint256 amount) returns()
func (_Synthetix *SynthetixTransactorSession) EmitExchangeRebate(account common.Address, currencyKey [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.EmitExchangeRebate(&_Synthetix.TransactOpts, account, currencyKey, amount)
}

// EmitExchangeReclaim is a paid mutator transaction binding the contract method 0xace88afd.
//
// Solidity: function emitExchangeReclaim(address account, bytes32 currencyKey, uint256 amount) returns()
func (_Synthetix *SynthetixTransactor) EmitExchangeReclaim(opts *bind.TransactOpts, account common.Address, currencyKey [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "emitExchangeReclaim", account, currencyKey, amount)
}

// EmitExchangeReclaim is a paid mutator transaction binding the contract method 0xace88afd.
//
// Solidity: function emitExchangeReclaim(address account, bytes32 currencyKey, uint256 amount) returns()
func (_Synthetix *SynthetixSession) EmitExchangeReclaim(account common.Address, currencyKey [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.EmitExchangeReclaim(&_Synthetix.TransactOpts, account, currencyKey, amount)
}

// EmitExchangeReclaim is a paid mutator transaction binding the contract method 0xace88afd.
//
// Solidity: function emitExchangeReclaim(address account, bytes32 currencyKey, uint256 amount) returns()
func (_Synthetix *SynthetixTransactorSession) EmitExchangeReclaim(account common.Address, currencyKey [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.EmitExchangeReclaim(&_Synthetix.TransactOpts, account, currencyKey, amount)
}

// EmitExchangeTracking is a paid mutator transaction binding the contract method 0xddd03a3f.
//
// Solidity: function emitExchangeTracking(bytes32 trackingCode, bytes32 toCurrencyKey, uint256 toAmount) returns()
func (_Synthetix *SynthetixTransactor) EmitExchangeTracking(opts *bind.TransactOpts, trackingCode [32]byte, toCurrencyKey [32]byte, toAmount *big.Int) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "emitExchangeTracking", trackingCode, toCurrencyKey, toAmount)
}

// EmitExchangeTracking is a paid mutator transaction binding the contract method 0xddd03a3f.
//
// Solidity: function emitExchangeTracking(bytes32 trackingCode, bytes32 toCurrencyKey, uint256 toAmount) returns()
func (_Synthetix *SynthetixSession) EmitExchangeTracking(trackingCode [32]byte, toCurrencyKey [32]byte, toAmount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.EmitExchangeTracking(&_Synthetix.TransactOpts, trackingCode, toCurrencyKey, toAmount)
}

// EmitExchangeTracking is a paid mutator transaction binding the contract method 0xddd03a3f.
//
// Solidity: function emitExchangeTracking(bytes32 trackingCode, bytes32 toCurrencyKey, uint256 toAmount) returns()
func (_Synthetix *SynthetixTransactorSession) EmitExchangeTracking(trackingCode [32]byte, toCurrencyKey [32]byte, toAmount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.EmitExchangeTracking(&_Synthetix.TransactOpts, trackingCode, toCurrencyKey, toAmount)
}

// EmitSynthExchange is a paid mutator transaction binding the contract method 0x6c00f310.
//
// Solidity: function emitSynthExchange(address account, bytes32 fromCurrencyKey, uint256 fromAmount, bytes32 toCurrencyKey, uint256 toAmount, address toAddress) returns()
func (_Synthetix *SynthetixTransactor) EmitSynthExchange(opts *bind.TransactOpts, account common.Address, fromCurrencyKey [32]byte, fromAmount *big.Int, toCurrencyKey [32]byte, toAmount *big.Int, toAddress common.Address) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "emitSynthExchange", account, fromCurrencyKey, fromAmount, toCurrencyKey, toAmount, toAddress)
}

// EmitSynthExchange is a paid mutator transaction binding the contract method 0x6c00f310.
//
// Solidity: function emitSynthExchange(address account, bytes32 fromCurrencyKey, uint256 fromAmount, bytes32 toCurrencyKey, uint256 toAmount, address toAddress) returns()
func (_Synthetix *SynthetixSession) EmitSynthExchange(account common.Address, fromCurrencyKey [32]byte, fromAmount *big.Int, toCurrencyKey [32]byte, toAmount *big.Int, toAddress common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.EmitSynthExchange(&_Synthetix.TransactOpts, account, fromCurrencyKey, fromAmount, toCurrencyKey, toAmount, toAddress)
}

// EmitSynthExchange is a paid mutator transaction binding the contract method 0x6c00f310.
//
// Solidity: function emitSynthExchange(address account, bytes32 fromCurrencyKey, uint256 fromAmount, bytes32 toCurrencyKey, uint256 toAmount, address toAddress) returns()
func (_Synthetix *SynthetixTransactorSession) EmitSynthExchange(account common.Address, fromCurrencyKey [32]byte, fromAmount *big.Int, toCurrencyKey [32]byte, toAmount *big.Int, toAddress common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.EmitSynthExchange(&_Synthetix.TransactOpts, account, fromCurrencyKey, fromAmount, toCurrencyKey, toAmount, toAddress)
}

// Exchange is a paid mutator transaction binding the contract method 0xee52a2f3.
//
// Solidity: function exchange(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey) returns(uint256 amountReceived)
func (_Synthetix *SynthetixTransactor) Exchange(opts *bind.TransactOpts, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "exchange", sourceCurrencyKey, sourceAmount, destinationCurrencyKey)
}

// Exchange is a paid mutator transaction binding the contract method 0xee52a2f3.
//
// Solidity: function exchange(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey) returns(uint256 amountReceived)
func (_Synthetix *SynthetixSession) Exchange(sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte) (*types.Transaction, error) {
	return _Synthetix.Contract.Exchange(&_Synthetix.TransactOpts, sourceCurrencyKey, sourceAmount, destinationCurrencyKey)
}

// Exchange is a paid mutator transaction binding the contract method 0xee52a2f3.
//
// Solidity: function exchange(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey) returns(uint256 amountReceived)
func (_Synthetix *SynthetixTransactorSession) Exchange(sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte) (*types.Transaction, error) {
	return _Synthetix.Contract.Exchange(&_Synthetix.TransactOpts, sourceCurrencyKey, sourceAmount, destinationCurrencyKey)
}

// ExchangeOnBehalf is a paid mutator transaction binding the contract method 0xc836fa0a.
//
// Solidity: function exchangeOnBehalf(address exchangeForAddress, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey) returns(uint256 amountReceived)
func (_Synthetix *SynthetixTransactor) ExchangeOnBehalf(opts *bind.TransactOpts, exchangeForAddress common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "exchangeOnBehalf", exchangeForAddress, sourceCurrencyKey, sourceAmount, destinationCurrencyKey)
}

// ExchangeOnBehalf is a paid mutator transaction binding the contract method 0xc836fa0a.
//
// Solidity: function exchangeOnBehalf(address exchangeForAddress, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey) returns(uint256 amountReceived)
func (_Synthetix *SynthetixSession) ExchangeOnBehalf(exchangeForAddress common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte) (*types.Transaction, error) {
	return _Synthetix.Contract.ExchangeOnBehalf(&_Synthetix.TransactOpts, exchangeForAddress, sourceCurrencyKey, sourceAmount, destinationCurrencyKey)
}

// ExchangeOnBehalf is a paid mutator transaction binding the contract method 0xc836fa0a.
//
// Solidity: function exchangeOnBehalf(address exchangeForAddress, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey) returns(uint256 amountReceived)
func (_Synthetix *SynthetixTransactorSession) ExchangeOnBehalf(exchangeForAddress common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte) (*types.Transaction, error) {
	return _Synthetix.Contract.ExchangeOnBehalf(&_Synthetix.TransactOpts, exchangeForAddress, sourceCurrencyKey, sourceAmount, destinationCurrencyKey)
}

// ExchangeOnBehalfWithTracking is a paid mutator transaction binding the contract method 0x91e56b68.
//
// Solidity: function exchangeOnBehalfWithTracking(address exchangeForAddress, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_Synthetix *SynthetixTransactor) ExchangeOnBehalfWithTracking(opts *bind.TransactOpts, exchangeForAddress common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "exchangeOnBehalfWithTracking", exchangeForAddress, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, originator, trackingCode)
}

// ExchangeOnBehalfWithTracking is a paid mutator transaction binding the contract method 0x91e56b68.
//
// Solidity: function exchangeOnBehalfWithTracking(address exchangeForAddress, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_Synthetix *SynthetixSession) ExchangeOnBehalfWithTracking(exchangeForAddress common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _Synthetix.Contract.ExchangeOnBehalfWithTracking(&_Synthetix.TransactOpts, exchangeForAddress, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, originator, trackingCode)
}

// ExchangeOnBehalfWithTracking is a paid mutator transaction binding the contract method 0x91e56b68.
//
// Solidity: function exchangeOnBehalfWithTracking(address exchangeForAddress, bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_Synthetix *SynthetixTransactorSession) ExchangeOnBehalfWithTracking(exchangeForAddress common.Address, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _Synthetix.Contract.ExchangeOnBehalfWithTracking(&_Synthetix.TransactOpts, exchangeForAddress, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, originator, trackingCode)
}

// ExchangeWithTracking is a paid mutator transaction binding the contract method 0x30ead760.
//
// Solidity: function exchangeWithTracking(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_Synthetix *SynthetixTransactor) ExchangeWithTracking(opts *bind.TransactOpts, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "exchangeWithTracking", sourceCurrencyKey, sourceAmount, destinationCurrencyKey, originator, trackingCode)
}

// ExchangeWithTracking is a paid mutator transaction binding the contract method 0x30ead760.
//
// Solidity: function exchangeWithTracking(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_Synthetix *SynthetixSession) ExchangeWithTracking(sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _Synthetix.Contract.ExchangeWithTracking(&_Synthetix.TransactOpts, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, originator, trackingCode)
}

// ExchangeWithTracking is a paid mutator transaction binding the contract method 0x30ead760.
//
// Solidity: function exchangeWithTracking(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, address originator, bytes32 trackingCode) returns(uint256 amountReceived)
func (_Synthetix *SynthetixTransactorSession) ExchangeWithTracking(sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, originator common.Address, trackingCode [32]byte) (*types.Transaction, error) {
	return _Synthetix.Contract.ExchangeWithTracking(&_Synthetix.TransactOpts, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, originator, trackingCode)
}

// ExchangeWithVirtual is a paid mutator transaction binding the contract method 0x0e30963c.
//
// Solidity: function exchangeWithVirtual(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, bytes32 trackingCode) returns(uint256 amountReceived, address vSynth)
func (_Synthetix *SynthetixTransactor) ExchangeWithVirtual(opts *bind.TransactOpts, sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, trackingCode [32]byte) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "exchangeWithVirtual", sourceCurrencyKey, sourceAmount, destinationCurrencyKey, trackingCode)
}

// ExchangeWithVirtual is a paid mutator transaction binding the contract method 0x0e30963c.
//
// Solidity: function exchangeWithVirtual(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, bytes32 trackingCode) returns(uint256 amountReceived, address vSynth)
func (_Synthetix *SynthetixSession) ExchangeWithVirtual(sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, trackingCode [32]byte) (*types.Transaction, error) {
	return _Synthetix.Contract.ExchangeWithVirtual(&_Synthetix.TransactOpts, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, trackingCode)
}

// ExchangeWithVirtual is a paid mutator transaction binding the contract method 0x0e30963c.
//
// Solidity: function exchangeWithVirtual(bytes32 sourceCurrencyKey, uint256 sourceAmount, bytes32 destinationCurrencyKey, bytes32 trackingCode) returns(uint256 amountReceived, address vSynth)
func (_Synthetix *SynthetixTransactorSession) ExchangeWithVirtual(sourceCurrencyKey [32]byte, sourceAmount *big.Int, destinationCurrencyKey [32]byte, trackingCode [32]byte) (*types.Transaction, error) {
	return _Synthetix.Contract.ExchangeWithVirtual(&_Synthetix.TransactOpts, sourceCurrencyKey, sourceAmount, destinationCurrencyKey, trackingCode)
}

// IssueMaxSynths is a paid mutator transaction binding the contract method 0xaf086c7e.
//
// Solidity: function issueMaxSynths() returns()
func (_Synthetix *SynthetixTransactor) IssueMaxSynths(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "issueMaxSynths")
}

// IssueMaxSynths is a paid mutator transaction binding the contract method 0xaf086c7e.
//
// Solidity: function issueMaxSynths() returns()
func (_Synthetix *SynthetixSession) IssueMaxSynths() (*types.Transaction, error) {
	return _Synthetix.Contract.IssueMaxSynths(&_Synthetix.TransactOpts)
}

// IssueMaxSynths is a paid mutator transaction binding the contract method 0xaf086c7e.
//
// Solidity: function issueMaxSynths() returns()
func (_Synthetix *SynthetixTransactorSession) IssueMaxSynths() (*types.Transaction, error) {
	return _Synthetix.Contract.IssueMaxSynths(&_Synthetix.TransactOpts)
}

// IssueMaxSynthsOnBehalf is a paid mutator transaction binding the contract method 0x320223db.
//
// Solidity: function issueMaxSynthsOnBehalf(address issueForAddress) returns()
func (_Synthetix *SynthetixTransactor) IssueMaxSynthsOnBehalf(opts *bind.TransactOpts, issueForAddress common.Address) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "issueMaxSynthsOnBehalf", issueForAddress)
}

// IssueMaxSynthsOnBehalf is a paid mutator transaction binding the contract method 0x320223db.
//
// Solidity: function issueMaxSynthsOnBehalf(address issueForAddress) returns()
func (_Synthetix *SynthetixSession) IssueMaxSynthsOnBehalf(issueForAddress common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.IssueMaxSynthsOnBehalf(&_Synthetix.TransactOpts, issueForAddress)
}

// IssueMaxSynthsOnBehalf is a paid mutator transaction binding the contract method 0x320223db.
//
// Solidity: function issueMaxSynthsOnBehalf(address issueForAddress) returns()
func (_Synthetix *SynthetixTransactorSession) IssueMaxSynthsOnBehalf(issueForAddress common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.IssueMaxSynthsOnBehalf(&_Synthetix.TransactOpts, issueForAddress)
}

// IssueSynths is a paid mutator transaction binding the contract method 0x8a290014.
//
// Solidity: function issueSynths(uint256 amount) returns()
func (_Synthetix *SynthetixTransactor) IssueSynths(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "issueSynths", amount)
}

// IssueSynths is a paid mutator transaction binding the contract method 0x8a290014.
//
// Solidity: function issueSynths(uint256 amount) returns()
func (_Synthetix *SynthetixSession) IssueSynths(amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.IssueSynths(&_Synthetix.TransactOpts, amount)
}

// IssueSynths is a paid mutator transaction binding the contract method 0x8a290014.
//
// Solidity: function issueSynths(uint256 amount) returns()
func (_Synthetix *SynthetixTransactorSession) IssueSynths(amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.IssueSynths(&_Synthetix.TransactOpts, amount)
}

// IssueSynthsOnBehalf is a paid mutator transaction binding the contract method 0xe8e09b8b.
//
// Solidity: function issueSynthsOnBehalf(address issueForAddress, uint256 amount) returns()
func (_Synthetix *SynthetixTransactor) IssueSynthsOnBehalf(opts *bind.TransactOpts, issueForAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "issueSynthsOnBehalf", issueForAddress, amount)
}

// IssueSynthsOnBehalf is a paid mutator transaction binding the contract method 0xe8e09b8b.
//
// Solidity: function issueSynthsOnBehalf(address issueForAddress, uint256 amount) returns()
func (_Synthetix *SynthetixSession) IssueSynthsOnBehalf(issueForAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.IssueSynthsOnBehalf(&_Synthetix.TransactOpts, issueForAddress, amount)
}

// IssueSynthsOnBehalf is a paid mutator transaction binding the contract method 0xe8e09b8b.
//
// Solidity: function issueSynthsOnBehalf(address issueForAddress, uint256 amount) returns()
func (_Synthetix *SynthetixTransactorSession) IssueSynthsOnBehalf(issueForAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.IssueSynthsOnBehalf(&_Synthetix.TransactOpts, issueForAddress, amount)
}

// LiquidateDelinquentAccount is a paid mutator transaction binding the contract method 0xe6203ed1.
//
// Solidity: function liquidateDelinquentAccount(address account, uint256 susdAmount) returns(bool)
func (_Synthetix *SynthetixTransactor) LiquidateDelinquentAccount(opts *bind.TransactOpts, account common.Address, susdAmount *big.Int) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "liquidateDelinquentAccount", account, susdAmount)
}

// LiquidateDelinquentAccount is a paid mutator transaction binding the contract method 0xe6203ed1.
//
// Solidity: function liquidateDelinquentAccount(address account, uint256 susdAmount) returns(bool)
func (_Synthetix *SynthetixSession) LiquidateDelinquentAccount(account common.Address, susdAmount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.LiquidateDelinquentAccount(&_Synthetix.TransactOpts, account, susdAmount)
}

// LiquidateDelinquentAccount is a paid mutator transaction binding the contract method 0xe6203ed1.
//
// Solidity: function liquidateDelinquentAccount(address account, uint256 susdAmount) returns(bool)
func (_Synthetix *SynthetixTransactorSession) LiquidateDelinquentAccount(account common.Address, susdAmount *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.LiquidateDelinquentAccount(&_Synthetix.TransactOpts, account, susdAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(bool)
func (_Synthetix *SynthetixTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(bool)
func (_Synthetix *SynthetixSession) Mint() (*types.Transaction, error) {
	return _Synthetix.Contract.Mint(&_Synthetix.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(bool)
func (_Synthetix *SynthetixTransactorSession) Mint() (*types.Transaction, error) {
	return _Synthetix.Contract.Mint(&_Synthetix.TransactOpts)
}

// MintSecondary is a paid mutator transaction binding the contract method 0x666ed4f1.
//
// Solidity: function mintSecondary(address , uint256 ) returns()
func (_Synthetix *SynthetixTransactor) MintSecondary(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "mintSecondary", arg0, arg1)
}

// MintSecondary is a paid mutator transaction binding the contract method 0x666ed4f1.
//
// Solidity: function mintSecondary(address , uint256 ) returns()
func (_Synthetix *SynthetixSession) MintSecondary(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.MintSecondary(&_Synthetix.TransactOpts, arg0, arg1)
}

// MintSecondary is a paid mutator transaction binding the contract method 0x666ed4f1.
//
// Solidity: function mintSecondary(address , uint256 ) returns()
func (_Synthetix *SynthetixTransactorSession) MintSecondary(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.MintSecondary(&_Synthetix.TransactOpts, arg0, arg1)
}

// MintSecondaryRewards is a paid mutator transaction binding the contract method 0xd8a1f76f.
//
// Solidity: function mintSecondaryRewards(uint256 ) returns()
func (_Synthetix *SynthetixTransactor) MintSecondaryRewards(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "mintSecondaryRewards", arg0)
}

// MintSecondaryRewards is a paid mutator transaction binding the contract method 0xd8a1f76f.
//
// Solidity: function mintSecondaryRewards(uint256 ) returns()
func (_Synthetix *SynthetixSession) MintSecondaryRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.MintSecondaryRewards(&_Synthetix.TransactOpts, arg0)
}

// MintSecondaryRewards is a paid mutator transaction binding the contract method 0xd8a1f76f.
//
// Solidity: function mintSecondaryRewards(uint256 ) returns()
func (_Synthetix *SynthetixTransactorSession) MintSecondaryRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.MintSecondaryRewards(&_Synthetix.TransactOpts, arg0)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Synthetix *SynthetixTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Synthetix *SynthetixSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.NominateNewOwner(&_Synthetix.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Synthetix *SynthetixTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.NominateNewOwner(&_Synthetix.TransactOpts, _owner)
}

// SetIntegrationProxy is a paid mutator transaction binding the contract method 0x131b0ae7.
//
// Solidity: function setIntegrationProxy(address _integrationProxy) returns()
func (_Synthetix *SynthetixTransactor) SetIntegrationProxy(opts *bind.TransactOpts, _integrationProxy common.Address) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "setIntegrationProxy", _integrationProxy)
}

// SetIntegrationProxy is a paid mutator transaction binding the contract method 0x131b0ae7.
//
// Solidity: function setIntegrationProxy(address _integrationProxy) returns()
func (_Synthetix *SynthetixSession) SetIntegrationProxy(_integrationProxy common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.SetIntegrationProxy(&_Synthetix.TransactOpts, _integrationProxy)
}

// SetIntegrationProxy is a paid mutator transaction binding the contract method 0x131b0ae7.
//
// Solidity: function setIntegrationProxy(address _integrationProxy) returns()
func (_Synthetix *SynthetixTransactorSession) SetIntegrationProxy(_integrationProxy common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.SetIntegrationProxy(&_Synthetix.TransactOpts, _integrationProxy)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_Synthetix *SynthetixTransactor) SetMessageSender(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "setMessageSender", sender)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_Synthetix *SynthetixSession) SetMessageSender(sender common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.SetMessageSender(&_Synthetix.TransactOpts, sender)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_Synthetix *SynthetixTransactorSession) SetMessageSender(sender common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.SetMessageSender(&_Synthetix.TransactOpts, sender)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_Synthetix *SynthetixTransactor) SetProxy(opts *bind.TransactOpts, _proxy common.Address) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "setProxy", _proxy)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_Synthetix *SynthetixSession) SetProxy(_proxy common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.SetProxy(&_Synthetix.TransactOpts, _proxy)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_Synthetix *SynthetixTransactorSession) SetProxy(_proxy common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.SetProxy(&_Synthetix.TransactOpts, _proxy)
}

// SetResolverAndSyncCache is a paid mutator transaction binding the contract method 0x3be99e6f.
//
// Solidity: function setResolverAndSyncCache(address _resolver) returns()
func (_Synthetix *SynthetixTransactor) SetResolverAndSyncCache(opts *bind.TransactOpts, _resolver common.Address) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "setResolverAndSyncCache", _resolver)
}

// SetResolverAndSyncCache is a paid mutator transaction binding the contract method 0x3be99e6f.
//
// Solidity: function setResolverAndSyncCache(address _resolver) returns()
func (_Synthetix *SynthetixSession) SetResolverAndSyncCache(_resolver common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.SetResolverAndSyncCache(&_Synthetix.TransactOpts, _resolver)
}

// SetResolverAndSyncCache is a paid mutator transaction binding the contract method 0x3be99e6f.
//
// Solidity: function setResolverAndSyncCache(address _resolver) returns()
func (_Synthetix *SynthetixTransactorSession) SetResolverAndSyncCache(_resolver common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.SetResolverAndSyncCache(&_Synthetix.TransactOpts, _resolver)
}

// SetTokenState is a paid mutator transaction binding the contract method 0x9f769807.
//
// Solidity: function setTokenState(address _tokenState) returns()
func (_Synthetix *SynthetixTransactor) SetTokenState(opts *bind.TransactOpts, _tokenState common.Address) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "setTokenState", _tokenState)
}

// SetTokenState is a paid mutator transaction binding the contract method 0x9f769807.
//
// Solidity: function setTokenState(address _tokenState) returns()
func (_Synthetix *SynthetixSession) SetTokenState(_tokenState common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.SetTokenState(&_Synthetix.TransactOpts, _tokenState)
}

// SetTokenState is a paid mutator transaction binding the contract method 0x9f769807.
//
// Solidity: function setTokenState(address _tokenState) returns()
func (_Synthetix *SynthetixTransactorSession) SetTokenState(_tokenState common.Address) (*types.Transaction, error) {
	return _Synthetix.Contract.SetTokenState(&_Synthetix.TransactOpts, _tokenState)
}

// Settle is a paid mutator transaction binding the contract method 0x987757dd.
//
// Solidity: function settle(bytes32 currencyKey) returns(uint256 reclaimed, uint256 refunded, uint256 numEntriesSettled)
func (_Synthetix *SynthetixTransactor) Settle(opts *bind.TransactOpts, currencyKey [32]byte) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "settle", currencyKey)
}

// Settle is a paid mutator transaction binding the contract method 0x987757dd.
//
// Solidity: function settle(bytes32 currencyKey) returns(uint256 reclaimed, uint256 refunded, uint256 numEntriesSettled)
func (_Synthetix *SynthetixSession) Settle(currencyKey [32]byte) (*types.Transaction, error) {
	return _Synthetix.Contract.Settle(&_Synthetix.TransactOpts, currencyKey)
}

// Settle is a paid mutator transaction binding the contract method 0x987757dd.
//
// Solidity: function settle(bytes32 currencyKey) returns(uint256 reclaimed, uint256 refunded, uint256 numEntriesSettled)
func (_Synthetix *SynthetixTransactorSession) Settle(currencyKey [32]byte) (*types.Transaction, error) {
	return _Synthetix.Contract.Settle(&_Synthetix.TransactOpts, currencyKey)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Synthetix *SynthetixTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Synthetix *SynthetixSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.Transfer(&_Synthetix.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Synthetix *SynthetixTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.Transfer(&_Synthetix.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Synthetix *SynthetixTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Synthetix.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Synthetix *SynthetixSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.TransferFrom(&_Synthetix.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Synthetix *SynthetixTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Synthetix.Contract.TransferFrom(&_Synthetix.TransactOpts, from, to, value)
}

// SynthetixAccountLiquidatedIterator is returned from FilterAccountLiquidated and is used to iterate over the raw logs and unpacked data for AccountLiquidated events raised by the Synthetix contract.
type SynthetixAccountLiquidatedIterator struct {
	Event *SynthetixAccountLiquidated // Event containing the contract specifics and raw log

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
func (it *SynthetixAccountLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixAccountLiquidated)
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
		it.Event = new(SynthetixAccountLiquidated)
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
func (it *SynthetixAccountLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixAccountLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixAccountLiquidated represents a AccountLiquidated event raised by the Synthetix contract.
type SynthetixAccountLiquidated struct {
	Account          common.Address
	SnxRedeemed      *big.Int
	AmountLiquidated *big.Int
	Liquidator       common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAccountLiquidated is a free log retrieval operation binding the contract event 0xaadb11d74982254be0fa96d24a08db29d68f446bc96b3092a9c9120b5c89caf2.
//
// Solidity: event AccountLiquidated(address indexed account, uint256 snxRedeemed, uint256 amountLiquidated, address liquidator)
func (_Synthetix *SynthetixFilterer) FilterAccountLiquidated(opts *bind.FilterOpts, account []common.Address) (*SynthetixAccountLiquidatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Synthetix.contract.FilterLogs(opts, "AccountLiquidated", accountRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixAccountLiquidatedIterator{contract: _Synthetix.contract, event: "AccountLiquidated", logs: logs, sub: sub}, nil
}

// WatchAccountLiquidated is a free log subscription operation binding the contract event 0xaadb11d74982254be0fa96d24a08db29d68f446bc96b3092a9c9120b5c89caf2.
//
// Solidity: event AccountLiquidated(address indexed account, uint256 snxRedeemed, uint256 amountLiquidated, address liquidator)
func (_Synthetix *SynthetixFilterer) WatchAccountLiquidated(opts *bind.WatchOpts, sink chan<- *SynthetixAccountLiquidated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Synthetix.contract.WatchLogs(opts, "AccountLiquidated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixAccountLiquidated)
				if err := _Synthetix.contract.UnpackLog(event, "AccountLiquidated", log); err != nil {
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

// ParseAccountLiquidated is a log parse operation binding the contract event 0xaadb11d74982254be0fa96d24a08db29d68f446bc96b3092a9c9120b5c89caf2.
//
// Solidity: event AccountLiquidated(address indexed account, uint256 snxRedeemed, uint256 amountLiquidated, address liquidator)
func (_Synthetix *SynthetixFilterer) ParseAccountLiquidated(log types.Log) (*SynthetixAccountLiquidated, error) {
	event := new(SynthetixAccountLiquidated)
	if err := _Synthetix.contract.UnpackLog(event, "AccountLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Synthetix contract.
type SynthetixApprovalIterator struct {
	Event *SynthetixApproval // Event containing the contract specifics and raw log

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
func (it *SynthetixApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixApproval)
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
		it.Event = new(SynthetixApproval)
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
func (it *SynthetixApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixApproval represents a Approval event raised by the Synthetix contract.
type SynthetixApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Synthetix *SynthetixFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SynthetixApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Synthetix.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixApprovalIterator{contract: _Synthetix.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Synthetix *SynthetixFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SynthetixApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Synthetix.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixApproval)
				if err := _Synthetix.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Synthetix *SynthetixFilterer) ParseApproval(log types.Log) (*SynthetixApproval, error) {
	event := new(SynthetixApproval)
	if err := _Synthetix.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixExchangeRebateIterator is returned from FilterExchangeRebate and is used to iterate over the raw logs and unpacked data for ExchangeRebate events raised by the Synthetix contract.
type SynthetixExchangeRebateIterator struct {
	Event *SynthetixExchangeRebate // Event containing the contract specifics and raw log

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
func (it *SynthetixExchangeRebateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixExchangeRebate)
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
		it.Event = new(SynthetixExchangeRebate)
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
func (it *SynthetixExchangeRebateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixExchangeRebateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixExchangeRebate represents a ExchangeRebate event raised by the Synthetix contract.
type SynthetixExchangeRebate struct {
	Account     common.Address
	CurrencyKey [32]byte
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExchangeRebate is a free log retrieval operation binding the contract event 0x93751433c6897553c8950f14ccc193ccffb8f539f7421ffde9af83b9b7dae1a8.
//
// Solidity: event ExchangeRebate(address indexed account, bytes32 currencyKey, uint256 amount)
func (_Synthetix *SynthetixFilterer) FilterExchangeRebate(opts *bind.FilterOpts, account []common.Address) (*SynthetixExchangeRebateIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Synthetix.contract.FilterLogs(opts, "ExchangeRebate", accountRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixExchangeRebateIterator{contract: _Synthetix.contract, event: "ExchangeRebate", logs: logs, sub: sub}, nil
}

// WatchExchangeRebate is a free log subscription operation binding the contract event 0x93751433c6897553c8950f14ccc193ccffb8f539f7421ffde9af83b9b7dae1a8.
//
// Solidity: event ExchangeRebate(address indexed account, bytes32 currencyKey, uint256 amount)
func (_Synthetix *SynthetixFilterer) WatchExchangeRebate(opts *bind.WatchOpts, sink chan<- *SynthetixExchangeRebate, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Synthetix.contract.WatchLogs(opts, "ExchangeRebate", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixExchangeRebate)
				if err := _Synthetix.contract.UnpackLog(event, "ExchangeRebate", log); err != nil {
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

// ParseExchangeRebate is a log parse operation binding the contract event 0x93751433c6897553c8950f14ccc193ccffb8f539f7421ffde9af83b9b7dae1a8.
//
// Solidity: event ExchangeRebate(address indexed account, bytes32 currencyKey, uint256 amount)
func (_Synthetix *SynthetixFilterer) ParseExchangeRebate(log types.Log) (*SynthetixExchangeRebate, error) {
	event := new(SynthetixExchangeRebate)
	if err := _Synthetix.contract.UnpackLog(event, "ExchangeRebate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixExchangeReclaimIterator is returned from FilterExchangeReclaim and is used to iterate over the raw logs and unpacked data for ExchangeReclaim events raised by the Synthetix contract.
type SynthetixExchangeReclaimIterator struct {
	Event *SynthetixExchangeReclaim // Event containing the contract specifics and raw log

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
func (it *SynthetixExchangeReclaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixExchangeReclaim)
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
		it.Event = new(SynthetixExchangeReclaim)
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
func (it *SynthetixExchangeReclaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixExchangeReclaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixExchangeReclaim represents a ExchangeReclaim event raised by the Synthetix contract.
type SynthetixExchangeReclaim struct {
	Account     common.Address
	CurrencyKey [32]byte
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExchangeReclaim is a free log retrieval operation binding the contract event 0x491df6adf9cabe8ca514806effd6b6b6475572dc88fe4b8b58d0a20ecf45e105.
//
// Solidity: event ExchangeReclaim(address indexed account, bytes32 currencyKey, uint256 amount)
func (_Synthetix *SynthetixFilterer) FilterExchangeReclaim(opts *bind.FilterOpts, account []common.Address) (*SynthetixExchangeReclaimIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Synthetix.contract.FilterLogs(opts, "ExchangeReclaim", accountRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixExchangeReclaimIterator{contract: _Synthetix.contract, event: "ExchangeReclaim", logs: logs, sub: sub}, nil
}

// WatchExchangeReclaim is a free log subscription operation binding the contract event 0x491df6adf9cabe8ca514806effd6b6b6475572dc88fe4b8b58d0a20ecf45e105.
//
// Solidity: event ExchangeReclaim(address indexed account, bytes32 currencyKey, uint256 amount)
func (_Synthetix *SynthetixFilterer) WatchExchangeReclaim(opts *bind.WatchOpts, sink chan<- *SynthetixExchangeReclaim, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Synthetix.contract.WatchLogs(opts, "ExchangeReclaim", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixExchangeReclaim)
				if err := _Synthetix.contract.UnpackLog(event, "ExchangeReclaim", log); err != nil {
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

// ParseExchangeReclaim is a log parse operation binding the contract event 0x491df6adf9cabe8ca514806effd6b6b6475572dc88fe4b8b58d0a20ecf45e105.
//
// Solidity: event ExchangeReclaim(address indexed account, bytes32 currencyKey, uint256 amount)
func (_Synthetix *SynthetixFilterer) ParseExchangeReclaim(log types.Log) (*SynthetixExchangeReclaim, error) {
	event := new(SynthetixExchangeReclaim)
	if err := _Synthetix.contract.UnpackLog(event, "ExchangeReclaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixExchangeTrackingIterator is returned from FilterExchangeTracking and is used to iterate over the raw logs and unpacked data for ExchangeTracking events raised by the Synthetix contract.
type SynthetixExchangeTrackingIterator struct {
	Event *SynthetixExchangeTracking // Event containing the contract specifics and raw log

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
func (it *SynthetixExchangeTrackingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixExchangeTracking)
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
		it.Event = new(SynthetixExchangeTracking)
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
func (it *SynthetixExchangeTrackingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixExchangeTrackingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixExchangeTracking represents a ExchangeTracking event raised by the Synthetix contract.
type SynthetixExchangeTracking struct {
	TrackingCode  [32]byte
	ToCurrencyKey [32]byte
	ToAmount      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExchangeTracking is a free log retrieval operation binding the contract event 0x9b39fce028952c685c9c73b2f5f825f8e369fbdaca2bec73c4abb52c2abc123c.
//
// Solidity: event ExchangeTracking(bytes32 indexed trackingCode, bytes32 toCurrencyKey, uint256 toAmount)
func (_Synthetix *SynthetixFilterer) FilterExchangeTracking(opts *bind.FilterOpts, trackingCode [][32]byte) (*SynthetixExchangeTrackingIterator, error) {

	var trackingCodeRule []interface{}
	for _, trackingCodeItem := range trackingCode {
		trackingCodeRule = append(trackingCodeRule, trackingCodeItem)
	}

	logs, sub, err := _Synthetix.contract.FilterLogs(opts, "ExchangeTracking", trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixExchangeTrackingIterator{contract: _Synthetix.contract, event: "ExchangeTracking", logs: logs, sub: sub}, nil
}

// WatchExchangeTracking is a free log subscription operation binding the contract event 0x9b39fce028952c685c9c73b2f5f825f8e369fbdaca2bec73c4abb52c2abc123c.
//
// Solidity: event ExchangeTracking(bytes32 indexed trackingCode, bytes32 toCurrencyKey, uint256 toAmount)
func (_Synthetix *SynthetixFilterer) WatchExchangeTracking(opts *bind.WatchOpts, sink chan<- *SynthetixExchangeTracking, trackingCode [][32]byte) (event.Subscription, error) {

	var trackingCodeRule []interface{}
	for _, trackingCodeItem := range trackingCode {
		trackingCodeRule = append(trackingCodeRule, trackingCodeItem)
	}

	logs, sub, err := _Synthetix.contract.WatchLogs(opts, "ExchangeTracking", trackingCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixExchangeTracking)
				if err := _Synthetix.contract.UnpackLog(event, "ExchangeTracking", log); err != nil {
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

// ParseExchangeTracking is a log parse operation binding the contract event 0x9b39fce028952c685c9c73b2f5f825f8e369fbdaca2bec73c4abb52c2abc123c.
//
// Solidity: event ExchangeTracking(bytes32 indexed trackingCode, bytes32 toCurrencyKey, uint256 toAmount)
func (_Synthetix *SynthetixFilterer) ParseExchangeTracking(log types.Log) (*SynthetixExchangeTracking, error) {
	event := new(SynthetixExchangeTracking)
	if err := _Synthetix.contract.UnpackLog(event, "ExchangeTracking", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Synthetix contract.
type SynthetixOwnerChangedIterator struct {
	Event *SynthetixOwnerChanged // Event containing the contract specifics and raw log

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
func (it *SynthetixOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixOwnerChanged)
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
		it.Event = new(SynthetixOwnerChanged)
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
func (it *SynthetixOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixOwnerChanged represents a OwnerChanged event raised by the Synthetix contract.
type SynthetixOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Synthetix *SynthetixFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*SynthetixOwnerChangedIterator, error) {

	logs, sub, err := _Synthetix.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &SynthetixOwnerChangedIterator{contract: _Synthetix.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Synthetix *SynthetixFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *SynthetixOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Synthetix.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixOwnerChanged)
				if err := _Synthetix.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_Synthetix *SynthetixFilterer) ParseOwnerChanged(log types.Log) (*SynthetixOwnerChanged, error) {
	event := new(SynthetixOwnerChanged)
	if err := _Synthetix.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the Synthetix contract.
type SynthetixOwnerNominatedIterator struct {
	Event *SynthetixOwnerNominated // Event containing the contract specifics and raw log

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
func (it *SynthetixOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixOwnerNominated)
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
		it.Event = new(SynthetixOwnerNominated)
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
func (it *SynthetixOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixOwnerNominated represents a OwnerNominated event raised by the Synthetix contract.
type SynthetixOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Synthetix *SynthetixFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*SynthetixOwnerNominatedIterator, error) {

	logs, sub, err := _Synthetix.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &SynthetixOwnerNominatedIterator{contract: _Synthetix.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Synthetix *SynthetixFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *SynthetixOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _Synthetix.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixOwnerNominated)
				if err := _Synthetix.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_Synthetix *SynthetixFilterer) ParseOwnerNominated(log types.Log) (*SynthetixOwnerNominated, error) {
	event := new(SynthetixOwnerNominated)
	if err := _Synthetix.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixProxyUpdatedIterator is returned from FilterProxyUpdated and is used to iterate over the raw logs and unpacked data for ProxyUpdated events raised by the Synthetix contract.
type SynthetixProxyUpdatedIterator struct {
	Event *SynthetixProxyUpdated // Event containing the contract specifics and raw log

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
func (it *SynthetixProxyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixProxyUpdated)
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
		it.Event = new(SynthetixProxyUpdated)
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
func (it *SynthetixProxyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixProxyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixProxyUpdated represents a ProxyUpdated event raised by the Synthetix contract.
type SynthetixProxyUpdated struct {
	ProxyAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProxyUpdated is a free log retrieval operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_Synthetix *SynthetixFilterer) FilterProxyUpdated(opts *bind.FilterOpts) (*SynthetixProxyUpdatedIterator, error) {

	logs, sub, err := _Synthetix.contract.FilterLogs(opts, "ProxyUpdated")
	if err != nil {
		return nil, err
	}
	return &SynthetixProxyUpdatedIterator{contract: _Synthetix.contract, event: "ProxyUpdated", logs: logs, sub: sub}, nil
}

// WatchProxyUpdated is a free log subscription operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_Synthetix *SynthetixFilterer) WatchProxyUpdated(opts *bind.WatchOpts, sink chan<- *SynthetixProxyUpdated) (event.Subscription, error) {

	logs, sub, err := _Synthetix.contract.WatchLogs(opts, "ProxyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixProxyUpdated)
				if err := _Synthetix.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
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

// ParseProxyUpdated is a log parse operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_Synthetix *SynthetixFilterer) ParseProxyUpdated(log types.Log) (*SynthetixProxyUpdated, error) {
	event := new(SynthetixProxyUpdated)
	if err := _Synthetix.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixSynthExchangeIterator is returned from FilterSynthExchange and is used to iterate over the raw logs and unpacked data for SynthExchange events raised by the Synthetix contract.
type SynthetixSynthExchangeIterator struct {
	Event *SynthetixSynthExchange // Event containing the contract specifics and raw log

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
func (it *SynthetixSynthExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixSynthExchange)
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
		it.Event = new(SynthetixSynthExchange)
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
func (it *SynthetixSynthExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixSynthExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixSynthExchange represents a SynthExchange event raised by the Synthetix contract.
type SynthetixSynthExchange struct {
	Account         common.Address
	FromCurrencyKey [32]byte
	FromAmount      *big.Int
	ToCurrencyKey   [32]byte
	ToAmount        *big.Int
	ToAddress       common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSynthExchange is a free log retrieval operation binding the contract event 0x65b6972c94204d84cffd3a95615743e31270f04fdf251f3dccc705cfbad44776.
//
// Solidity: event SynthExchange(address indexed account, bytes32 fromCurrencyKey, uint256 fromAmount, bytes32 toCurrencyKey, uint256 toAmount, address toAddress)
func (_Synthetix *SynthetixFilterer) FilterSynthExchange(opts *bind.FilterOpts, account []common.Address) (*SynthetixSynthExchangeIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Synthetix.contract.FilterLogs(opts, "SynthExchange", accountRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixSynthExchangeIterator{contract: _Synthetix.contract, event: "SynthExchange", logs: logs, sub: sub}, nil
}

// WatchSynthExchange is a free log subscription operation binding the contract event 0x65b6972c94204d84cffd3a95615743e31270f04fdf251f3dccc705cfbad44776.
//
// Solidity: event SynthExchange(address indexed account, bytes32 fromCurrencyKey, uint256 fromAmount, bytes32 toCurrencyKey, uint256 toAmount, address toAddress)
func (_Synthetix *SynthetixFilterer) WatchSynthExchange(opts *bind.WatchOpts, sink chan<- *SynthetixSynthExchange, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Synthetix.contract.WatchLogs(opts, "SynthExchange", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixSynthExchange)
				if err := _Synthetix.contract.UnpackLog(event, "SynthExchange", log); err != nil {
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

// ParseSynthExchange is a log parse operation binding the contract event 0x65b6972c94204d84cffd3a95615743e31270f04fdf251f3dccc705cfbad44776.
//
// Solidity: event SynthExchange(address indexed account, bytes32 fromCurrencyKey, uint256 fromAmount, bytes32 toCurrencyKey, uint256 toAmount, address toAddress)
func (_Synthetix *SynthetixFilterer) ParseSynthExchange(log types.Log) (*SynthetixSynthExchange, error) {
	event := new(SynthetixSynthExchange)
	if err := _Synthetix.contract.UnpackLog(event, "SynthExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixTokenStateUpdatedIterator is returned from FilterTokenStateUpdated and is used to iterate over the raw logs and unpacked data for TokenStateUpdated events raised by the Synthetix contract.
type SynthetixTokenStateUpdatedIterator struct {
	Event *SynthetixTokenStateUpdated // Event containing the contract specifics and raw log

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
func (it *SynthetixTokenStateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixTokenStateUpdated)
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
		it.Event = new(SynthetixTokenStateUpdated)
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
func (it *SynthetixTokenStateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixTokenStateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixTokenStateUpdated represents a TokenStateUpdated event raised by the Synthetix contract.
type SynthetixTokenStateUpdated struct {
	NewTokenState common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenStateUpdated is a free log retrieval operation binding the contract event 0xa538c4dcfe9fb148efee2952bafe34982d2d07d5fbb38ae5b44abf659a46bfd8.
//
// Solidity: event TokenStateUpdated(address newTokenState)
func (_Synthetix *SynthetixFilterer) FilterTokenStateUpdated(opts *bind.FilterOpts) (*SynthetixTokenStateUpdatedIterator, error) {

	logs, sub, err := _Synthetix.contract.FilterLogs(opts, "TokenStateUpdated")
	if err != nil {
		return nil, err
	}
	return &SynthetixTokenStateUpdatedIterator{contract: _Synthetix.contract, event: "TokenStateUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenStateUpdated is a free log subscription operation binding the contract event 0xa538c4dcfe9fb148efee2952bafe34982d2d07d5fbb38ae5b44abf659a46bfd8.
//
// Solidity: event TokenStateUpdated(address newTokenState)
func (_Synthetix *SynthetixFilterer) WatchTokenStateUpdated(opts *bind.WatchOpts, sink chan<- *SynthetixTokenStateUpdated) (event.Subscription, error) {

	logs, sub, err := _Synthetix.contract.WatchLogs(opts, "TokenStateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixTokenStateUpdated)
				if err := _Synthetix.contract.UnpackLog(event, "TokenStateUpdated", log); err != nil {
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

// ParseTokenStateUpdated is a log parse operation binding the contract event 0xa538c4dcfe9fb148efee2952bafe34982d2d07d5fbb38ae5b44abf659a46bfd8.
//
// Solidity: event TokenStateUpdated(address newTokenState)
func (_Synthetix *SynthetixFilterer) ParseTokenStateUpdated(log types.Log) (*SynthetixTokenStateUpdated, error) {
	event := new(SynthetixTokenStateUpdated)
	if err := _Synthetix.contract.UnpackLog(event, "TokenStateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Synthetix contract.
type SynthetixTransferIterator struct {
	Event *SynthetixTransfer // Event containing the contract specifics and raw log

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
func (it *SynthetixTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixTransfer)
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
		it.Event = new(SynthetixTransfer)
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
func (it *SynthetixTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixTransfer represents a Transfer event raised by the Synthetix contract.
type SynthetixTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Synthetix *SynthetixFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SynthetixTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Synthetix.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixTransferIterator{contract: _Synthetix.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Synthetix *SynthetixFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SynthetixTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Synthetix.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixTransfer)
				if err := _Synthetix.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Synthetix *SynthetixFilterer) ParseTransfer(log types.Log) (*SynthetixTransfer, error) {
	event := new(SynthetixTransfer)
	if err := _Synthetix.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenStateABI is the input ABI used to generate the binding from.
const TokenStateABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_associatedContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"associatedContract\",\"type\":\"address\"}],\"name\":\"AssociatedContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"associatedContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setAllowance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_associatedContract\",\"type\":\"address\"}],\"name\":\"setAssociatedContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setBalanceOf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenStateFuncSigs maps the 4-byte function signature to its string representation.
var TokenStateFuncSigs = map[string]string{
	"79ba5097": "acceptOwnership()",
	"dd62ed3e": "allowance(address,address)",
	"aefc4ccb": "associatedContract()",
	"70a08231": "balanceOf(address)",
	"1627540c": "nominateNewOwner(address)",
	"53a47bb7": "nominatedOwner()",
	"8da5cb5b": "owner()",
	"da46098c": "setAllowance(address,address,uint256)",
	"52f445ca": "setAssociatedContract(address)",
	"b46310f6": "setBalanceOf(address,uint256)",
}

// TokenStateBin is the compiled bytecode used for deploying new contracts.
var TokenStateBin = "0x608060405234801561001057600080fd5b5060405161075d38038061075d8339818101604052604081101561003357600080fd5b50805160209091015180826001600160a01b038116610099576040805162461bcd60e51b815260206004820152601960248201527f4f776e657220616464726573732063616e6e6f74206265203000000000000000604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b038316908117825560408051928352602083019190915280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a1506000546001600160a01b0316610143576040805162461bcd60e51b815260206004820152601160248201527013dddb995c881b5d5cdd081899481cd95d607a1b604482015290519081900360640190fd5b600280546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f73f20cff579e8a4086fa607db83867595f1b6a798e718c0bfa0b94a404128e039181900360200190a15050506105b7806101a66000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80638da5cb5b116100665780638da5cb5b14610155578063aefc4ccb1461015d578063b46310f614610165578063da46098c14610191578063dd62ed3e146101c75761009e565b80631627540c146100a357806352f445ca146100cb57806353a47bb7146100f157806370a082311461011557806379ba50971461014d575b600080fd5b6100c9600480360360208110156100b957600080fd5b50356001600160a01b03166101f5565b005b6100c9600480360360208110156100e157600080fd5b50356001600160a01b0316610251565b6100f96102ad565b604080516001600160a01b039092168252519081900360200190f35b61013b6004803603602081101561012b57600080fd5b50356001600160a01b03166102bc565b60408051918252519081900360200190f35b6100c96102ce565b6100f961038a565b6100f9610399565b6100c96004803603604081101561017b57600080fd5b506001600160a01b0381351690602001356103a8565b6100c9600480360360608110156101a757600080fd5b506001600160a01b0381358116916020810135909116906040013561040d565b61013b600480360360408110156101dd57600080fd5b506001600160a01b0381358116916020013516610482565b6101fd61049f565b600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce229181900360200190a150565b61025961049f565b600280546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f73f20cff579e8a4086fa607db83867595f1b6a798e718c0bfa0b94a404128e039181900360200190a150565b6001546001600160a01b031681565b60036020526000908152604090205481565b6001546001600160a01b031633146103175760405162461bcd60e51b81526004018080602001828103825260358152602001806104eb6035913960400191505060405180910390fd5b600054600154604080516001600160a01b03938416815292909116602083015280517fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c9281900390910190a160018054600080546001600160a01b03199081166001600160a01b03841617909155169055565b6000546001600160a01b031681565b6002546001600160a01b031681565b6002546001600160a01b031633146103f15760405162461bcd60e51b815260040180806020018281038252603481526020018061054f6034913960400191505060405180910390fd5b6001600160a01b03909116600090815260036020526040902055565b6002546001600160a01b031633146104565760405162461bcd60e51b815260040180806020018281038252603481526020018061054f6034913960400191505060405180910390fd5b6001600160a01b0392831660009081526004602090815260408083209490951682529290925291902055565b600460209081526000928352604080842090915290825290205481565b6000546001600160a01b031633146104e85760405162461bcd60e51b815260040180806020018281038252602f815260200180610520602f913960400191505060405180910390fd5b56fe596f75206d757374206265206e6f6d696e61746564206265666f726520796f752063616e20616363657074206f776e6572736869704f6e6c792074686520636f6e7472616374206f776e6572206d617920706572666f726d207468697320616374696f6e4f6e6c7920746865206173736f63696174656420636f6e74726163742063616e20706572666f726d207468697320616374696f6ea265627a7a7231582032722b8f72345762d79732f2bb8881ddce4f5cb37cbbd2875d7050def7b2646f64736f6c63430005100032"

// DeployTokenState deploys a new Ethereum contract, binding an instance of TokenState to it.
func DeployTokenState(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _associatedContract common.Address) (common.Address, *types.Transaction, *TokenState, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenStateABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenStateBin), backend, _owner, _associatedContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenState{TokenStateCaller: TokenStateCaller{contract: contract}, TokenStateTransactor: TokenStateTransactor{contract: contract}, TokenStateFilterer: TokenStateFilterer{contract: contract}}, nil
}

// TokenState is an auto generated Go binding around an Ethereum contract.
type TokenState struct {
	TokenStateCaller     // Read-only binding to the contract
	TokenStateTransactor // Write-only binding to the contract
	TokenStateFilterer   // Log filterer for contract events
}

// TokenStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenStateSession struct {
	Contract     *TokenState       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenStateCallerSession struct {
	Contract *TokenStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenStateTransactorSession struct {
	Contract     *TokenStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenStateRaw struct {
	Contract *TokenState // Generic contract binding to access the raw methods on
}

// TokenStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenStateCallerRaw struct {
	Contract *TokenStateCaller // Generic read-only contract binding to access the raw methods on
}

// TokenStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenStateTransactorRaw struct {
	Contract *TokenStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenState creates a new instance of TokenState, bound to a specific deployed contract.
func NewTokenState(address common.Address, backend bind.ContractBackend) (*TokenState, error) {
	contract, err := bindTokenState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenState{TokenStateCaller: TokenStateCaller{contract: contract}, TokenStateTransactor: TokenStateTransactor{contract: contract}, TokenStateFilterer: TokenStateFilterer{contract: contract}}, nil
}

// NewTokenStateCaller creates a new read-only instance of TokenState, bound to a specific deployed contract.
func NewTokenStateCaller(address common.Address, caller bind.ContractCaller) (*TokenStateCaller, error) {
	contract, err := bindTokenState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenStateCaller{contract: contract}, nil
}

// NewTokenStateTransactor creates a new write-only instance of TokenState, bound to a specific deployed contract.
func NewTokenStateTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenStateTransactor, error) {
	contract, err := bindTokenState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenStateTransactor{contract: contract}, nil
}

// NewTokenStateFilterer creates a new log filterer instance of TokenState, bound to a specific deployed contract.
func NewTokenStateFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenStateFilterer, error) {
	contract, err := bindTokenState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenStateFilterer{contract: contract}, nil
}

// bindTokenState binds a generic wrapper to an already deployed contract.
func bindTokenState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenStateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenState *TokenStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenState.Contract.TokenStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenState *TokenStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenState.Contract.TokenStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenState *TokenStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenState.Contract.TokenStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenState *TokenStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenState *TokenStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenState *TokenStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenState.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_TokenState *TokenStateCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenState.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_TokenState *TokenStateSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TokenState.Contract.Allowance(&_TokenState.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_TokenState *TokenStateCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TokenState.Contract.Allowance(&_TokenState.CallOpts, arg0, arg1)
}

// AssociatedContract is a free data retrieval call binding the contract method 0xaefc4ccb.
//
// Solidity: function associatedContract() view returns(address)
func (_TokenState *TokenStateCaller) AssociatedContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenState.contract.Call(opts, &out, "associatedContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssociatedContract is a free data retrieval call binding the contract method 0xaefc4ccb.
//
// Solidity: function associatedContract() view returns(address)
func (_TokenState *TokenStateSession) AssociatedContract() (common.Address, error) {
	return _TokenState.Contract.AssociatedContract(&_TokenState.CallOpts)
}

// AssociatedContract is a free data retrieval call binding the contract method 0xaefc4ccb.
//
// Solidity: function associatedContract() view returns(address)
func (_TokenState *TokenStateCallerSession) AssociatedContract() (common.Address, error) {
	return _TokenState.Contract.AssociatedContract(&_TokenState.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_TokenState *TokenStateCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenState.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_TokenState *TokenStateSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TokenState.Contract.BalanceOf(&_TokenState.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_TokenState *TokenStateCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TokenState.Contract.BalanceOf(&_TokenState.CallOpts, arg0)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_TokenState *TokenStateCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenState.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_TokenState *TokenStateSession) NominatedOwner() (common.Address, error) {
	return _TokenState.Contract.NominatedOwner(&_TokenState.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_TokenState *TokenStateCallerSession) NominatedOwner() (common.Address, error) {
	return _TokenState.Contract.NominatedOwner(&_TokenState.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenState *TokenStateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenState.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenState *TokenStateSession) Owner() (common.Address, error) {
	return _TokenState.Contract.Owner(&_TokenState.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenState *TokenStateCallerSession) Owner() (common.Address, error) {
	return _TokenState.Contract.Owner(&_TokenState.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TokenState *TokenStateTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenState.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TokenState *TokenStateSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenState.Contract.AcceptOwnership(&_TokenState.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TokenState *TokenStateTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenState.Contract.AcceptOwnership(&_TokenState.TransactOpts)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_TokenState *TokenStateTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _TokenState.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_TokenState *TokenStateSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _TokenState.Contract.NominateNewOwner(&_TokenState.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_TokenState *TokenStateTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _TokenState.Contract.NominateNewOwner(&_TokenState.TransactOpts, _owner)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(address tokenOwner, address spender, uint256 value) returns()
func (_TokenState *TokenStateTransactor) SetAllowance(opts *bind.TransactOpts, tokenOwner common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenState.contract.Transact(opts, "setAllowance", tokenOwner, spender, value)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(address tokenOwner, address spender, uint256 value) returns()
func (_TokenState *TokenStateSession) SetAllowance(tokenOwner common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenState.Contract.SetAllowance(&_TokenState.TransactOpts, tokenOwner, spender, value)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(address tokenOwner, address spender, uint256 value) returns()
func (_TokenState *TokenStateTransactorSession) SetAllowance(tokenOwner common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenState.Contract.SetAllowance(&_TokenState.TransactOpts, tokenOwner, spender, value)
}

// SetAssociatedContract is a paid mutator transaction binding the contract method 0x52f445ca.
//
// Solidity: function setAssociatedContract(address _associatedContract) returns()
func (_TokenState *TokenStateTransactor) SetAssociatedContract(opts *bind.TransactOpts, _associatedContract common.Address) (*types.Transaction, error) {
	return _TokenState.contract.Transact(opts, "setAssociatedContract", _associatedContract)
}

// SetAssociatedContract is a paid mutator transaction binding the contract method 0x52f445ca.
//
// Solidity: function setAssociatedContract(address _associatedContract) returns()
func (_TokenState *TokenStateSession) SetAssociatedContract(_associatedContract common.Address) (*types.Transaction, error) {
	return _TokenState.Contract.SetAssociatedContract(&_TokenState.TransactOpts, _associatedContract)
}

// SetAssociatedContract is a paid mutator transaction binding the contract method 0x52f445ca.
//
// Solidity: function setAssociatedContract(address _associatedContract) returns()
func (_TokenState *TokenStateTransactorSession) SetAssociatedContract(_associatedContract common.Address) (*types.Transaction, error) {
	return _TokenState.Contract.SetAssociatedContract(&_TokenState.TransactOpts, _associatedContract)
}

// SetBalanceOf is a paid mutator transaction binding the contract method 0xb46310f6.
//
// Solidity: function setBalanceOf(address account, uint256 value) returns()
func (_TokenState *TokenStateTransactor) SetBalanceOf(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenState.contract.Transact(opts, "setBalanceOf", account, value)
}

// SetBalanceOf is a paid mutator transaction binding the contract method 0xb46310f6.
//
// Solidity: function setBalanceOf(address account, uint256 value) returns()
func (_TokenState *TokenStateSession) SetBalanceOf(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenState.Contract.SetBalanceOf(&_TokenState.TransactOpts, account, value)
}

// SetBalanceOf is a paid mutator transaction binding the contract method 0xb46310f6.
//
// Solidity: function setBalanceOf(address account, uint256 value) returns()
func (_TokenState *TokenStateTransactorSession) SetBalanceOf(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenState.Contract.SetBalanceOf(&_TokenState.TransactOpts, account, value)
}

// TokenStateAssociatedContractUpdatedIterator is returned from FilterAssociatedContractUpdated and is used to iterate over the raw logs and unpacked data for AssociatedContractUpdated events raised by the TokenState contract.
type TokenStateAssociatedContractUpdatedIterator struct {
	Event *TokenStateAssociatedContractUpdated // Event containing the contract specifics and raw log

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
func (it *TokenStateAssociatedContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenStateAssociatedContractUpdated)
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
		it.Event = new(TokenStateAssociatedContractUpdated)
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
func (it *TokenStateAssociatedContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenStateAssociatedContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenStateAssociatedContractUpdated represents a AssociatedContractUpdated event raised by the TokenState contract.
type TokenStateAssociatedContractUpdated struct {
	AssociatedContract common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAssociatedContractUpdated is a free log retrieval operation binding the contract event 0x73f20cff579e8a4086fa607db83867595f1b6a798e718c0bfa0b94a404128e03.
//
// Solidity: event AssociatedContractUpdated(address associatedContract)
func (_TokenState *TokenStateFilterer) FilterAssociatedContractUpdated(opts *bind.FilterOpts) (*TokenStateAssociatedContractUpdatedIterator, error) {

	logs, sub, err := _TokenState.contract.FilterLogs(opts, "AssociatedContractUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenStateAssociatedContractUpdatedIterator{contract: _TokenState.contract, event: "AssociatedContractUpdated", logs: logs, sub: sub}, nil
}

// WatchAssociatedContractUpdated is a free log subscription operation binding the contract event 0x73f20cff579e8a4086fa607db83867595f1b6a798e718c0bfa0b94a404128e03.
//
// Solidity: event AssociatedContractUpdated(address associatedContract)
func (_TokenState *TokenStateFilterer) WatchAssociatedContractUpdated(opts *bind.WatchOpts, sink chan<- *TokenStateAssociatedContractUpdated) (event.Subscription, error) {

	logs, sub, err := _TokenState.contract.WatchLogs(opts, "AssociatedContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenStateAssociatedContractUpdated)
				if err := _TokenState.contract.UnpackLog(event, "AssociatedContractUpdated", log); err != nil {
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

// ParseAssociatedContractUpdated is a log parse operation binding the contract event 0x73f20cff579e8a4086fa607db83867595f1b6a798e718c0bfa0b94a404128e03.
//
// Solidity: event AssociatedContractUpdated(address associatedContract)
func (_TokenState *TokenStateFilterer) ParseAssociatedContractUpdated(log types.Log) (*TokenStateAssociatedContractUpdated, error) {
	event := new(TokenStateAssociatedContractUpdated)
	if err := _TokenState.contract.UnpackLog(event, "AssociatedContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenStateOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the TokenState contract.
type TokenStateOwnerChangedIterator struct {
	Event *TokenStateOwnerChanged // Event containing the contract specifics and raw log

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
func (it *TokenStateOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenStateOwnerChanged)
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
		it.Event = new(TokenStateOwnerChanged)
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
func (it *TokenStateOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenStateOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenStateOwnerChanged represents a OwnerChanged event raised by the TokenState contract.
type TokenStateOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_TokenState *TokenStateFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*TokenStateOwnerChangedIterator, error) {

	logs, sub, err := _TokenState.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &TokenStateOwnerChangedIterator{contract: _TokenState.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_TokenState *TokenStateFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *TokenStateOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _TokenState.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenStateOwnerChanged)
				if err := _TokenState.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_TokenState *TokenStateFilterer) ParseOwnerChanged(log types.Log) (*TokenStateOwnerChanged, error) {
	event := new(TokenStateOwnerChanged)
	if err := _TokenState.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenStateOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the TokenState contract.
type TokenStateOwnerNominatedIterator struct {
	Event *TokenStateOwnerNominated // Event containing the contract specifics and raw log

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
func (it *TokenStateOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenStateOwnerNominated)
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
		it.Event = new(TokenStateOwnerNominated)
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
func (it *TokenStateOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenStateOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenStateOwnerNominated represents a OwnerNominated event raised by the TokenState contract.
type TokenStateOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_TokenState *TokenStateFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*TokenStateOwnerNominatedIterator, error) {

	logs, sub, err := _TokenState.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &TokenStateOwnerNominatedIterator{contract: _TokenState.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_TokenState *TokenStateFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *TokenStateOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _TokenState.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenStateOwnerNominated)
				if err := _TokenState.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_TokenState *TokenStateFilterer) ParseOwnerNominated(log types.Log) (*TokenStateOwnerNominated, error) {
	event := new(TokenStateOwnerNominated)
	if err := _TokenState.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
