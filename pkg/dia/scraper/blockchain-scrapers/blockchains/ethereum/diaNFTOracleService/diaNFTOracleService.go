// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package diaNFTOracleService

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

// DIANFTOracleMetaData contains all meta data concerning the DIANFTOracle contract.
var DIANFTOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value0\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value1\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value2\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value3\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value4\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"OracleUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUpdater\",\"type\":\"address\"}],\"name\":\"UpdaterAddressChange\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"value0\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"value1\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"value2\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"value3\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"value4\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOracleUpdaterAddress\",\"type\":\"address\"}],\"name\":\"updateOracleUpdaterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"960384a0": "getValue(string)",
		"e3b47da5": "setValue(string,uint64,uint64,uint64,uint64,uint64,uint64)",
		"6aa45efc": "updateOracleUpdaterAddress(address)",
		"5a9ade8b": "values(string)",
	},
	Bin: "0x608060405234801561001057600080fd5b50600180546001600160a01b0319163317905561068e806100326000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80635a9ade8b146100515780636aa45efc14610110578063960384a014610138578063e3b47da514610221575b600080fd5b6100f76004803603602081101561006757600080fd5b81019060208101813564010000000081111561008257600080fd5b82018360208201111561009457600080fd5b803590602001918460018302840111640100000000831117156100b657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506102f8945050505050565b6040805192835260208301919091528051918290030190f35b6101366004803603602081101561012657600080fd5b50356001600160a01b031661031c565b005b6101de6004803603602081101561014e57600080fd5b81019060208101813564010000000081111561016957600080fd5b82018360208201111561017b57600080fd5b8035906020019184600183028401116401000000008311171561019d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610387945050505050565b604080516001600160401b03978816815295871660208701529386168585015291851660608501528416608084015290921660a082015290519081900360c00190f35b610136600480360360e081101561023757600080fd5b81019060208101813564010000000081111561025257600080fd5b82018360208201111561026457600080fd5b8035906020019184600183028401116401000000008311171561028657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505081356001600160401b03908116935060208301358116926040810135821692506060810135821691608082013581169160a001351661045c565b80516020818301810180516000825292820191909301209152805460019091015482565b6001546001600160a01b0316331461033357600080fd5b600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f9181900360200190a150565b600080600080600080600080886040518082805190602001908083835b602083106103c35780518252601f1990920191602091820191016103a4565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003909101909320805490945060c081901c93608082901c6001600160401b0316935060009250600160401b91901c6001860154919006915060c081901c906001600160401b03608082901c1690600090600160401b9060401c969f959e50939c50919a50985050909106945092505050565b6001546001600160a01b0316331461047357600080fd5b60006040856001600160401b0316901b6080876001600160401b0316901b60c0896001600160401b0316901b0101905060006040836001600160401b0316901b6080856001600160401b0316901b60c0876001600160401b0316901b010190506000808a6040518082805190602001908083835b602083106105065780518252601f1990920191602091820191016104e7565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902090508281600001819055508181600101819055507fab06b7d7540b163ffbe1538e49d670388b97188db0c1eb2fdf41417a73a3d8278a8a8a8a8a8a8a6040518080602001886001600160401b03168152602001876001600160401b03168152602001866001600160401b03168152602001856001600160401b03168152602001846001600160401b03168152602001836001600160401b03168152602001828103825289818151815260200191508051906020019080838360005b8381101561060c5781810151838201526020016105f4565b50505050905090810190601f1680156106395780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390a15050505050505050505056fea2646970667358221220dca0a28f617c654950022e5ad23a835bda91efe006f80f158383105de0f6731664736f6c63430007040033",
}

// DIANFTOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use DIANFTOracleMetaData.ABI instead.
var DIANFTOracleABI = DIANFTOracleMetaData.ABI

// Deprecated: Use DIANFTOracleMetaData.Sigs instead.
// DIANFTOracleFuncSigs maps the 4-byte function signature to its string representation.
var DIANFTOracleFuncSigs = DIANFTOracleMetaData.Sigs

// DIANFTOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DIANFTOracleMetaData.Bin instead.
var DIANFTOracleBin = DIANFTOracleMetaData.Bin

// DeployDIANFTOracle deploys a new Ethereum contract, binding an instance of DIANFTOracle to it.
func DeployDIANFTOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DIANFTOracle, error) {
	parsed, err := DIANFTOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DIANFTOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DIANFTOracle{DIANFTOracleCaller: DIANFTOracleCaller{contract: contract}, DIANFTOracleTransactor: DIANFTOracleTransactor{contract: contract}, DIANFTOracleFilterer: DIANFTOracleFilterer{contract: contract}}, nil
}

// DIANFTOracle is an auto generated Go binding around an Ethereum contract.
type DIANFTOracle struct {
	DIANFTOracleCaller     // Read-only binding to the contract
	DIANFTOracleTransactor // Write-only binding to the contract
	DIANFTOracleFilterer   // Log filterer for contract events
}

// DIANFTOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type DIANFTOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIANFTOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DIANFTOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIANFTOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DIANFTOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIANFTOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DIANFTOracleSession struct {
	Contract     *DIANFTOracle     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DIANFTOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DIANFTOracleCallerSession struct {
	Contract *DIANFTOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DIANFTOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DIANFTOracleTransactorSession struct {
	Contract     *DIANFTOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DIANFTOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type DIANFTOracleRaw struct {
	Contract *DIANFTOracle // Generic contract binding to access the raw methods on
}

// DIANFTOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DIANFTOracleCallerRaw struct {
	Contract *DIANFTOracleCaller // Generic read-only contract binding to access the raw methods on
}

// DIANFTOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DIANFTOracleTransactorRaw struct {
	Contract *DIANFTOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDIANFTOracle creates a new instance of DIANFTOracle, bound to a specific deployed contract.
func NewDIANFTOracle(address common.Address, backend bind.ContractBackend) (*DIANFTOracle, error) {
	contract, err := bindDIANFTOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DIANFTOracle{DIANFTOracleCaller: DIANFTOracleCaller{contract: contract}, DIANFTOracleTransactor: DIANFTOracleTransactor{contract: contract}, DIANFTOracleFilterer: DIANFTOracleFilterer{contract: contract}}, nil
}

// NewDIANFTOracleCaller creates a new read-only instance of DIANFTOracle, bound to a specific deployed contract.
func NewDIANFTOracleCaller(address common.Address, caller bind.ContractCaller) (*DIANFTOracleCaller, error) {
	contract, err := bindDIANFTOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DIANFTOracleCaller{contract: contract}, nil
}

// NewDIANFTOracleTransactor creates a new write-only instance of DIANFTOracle, bound to a specific deployed contract.
func NewDIANFTOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*DIANFTOracleTransactor, error) {
	contract, err := bindDIANFTOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DIANFTOracleTransactor{contract: contract}, nil
}

// NewDIANFTOracleFilterer creates a new log filterer instance of DIANFTOracle, bound to a specific deployed contract.
func NewDIANFTOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*DIANFTOracleFilterer, error) {
	contract, err := bindDIANFTOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DIANFTOracleFilterer{contract: contract}, nil
}

// bindDIANFTOracle binds a generic wrapper to an already deployed contract.
func bindDIANFTOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DIANFTOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIANFTOracle *DIANFTOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIANFTOracle.Contract.DIANFTOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIANFTOracle *DIANFTOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIANFTOracle.Contract.DIANFTOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIANFTOracle *DIANFTOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIANFTOracle.Contract.DIANFTOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIANFTOracle *DIANFTOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIANFTOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIANFTOracle *DIANFTOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIANFTOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIANFTOracle *DIANFTOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIANFTOracle.Contract.contract.Transact(opts, method, params...)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint64, uint64, uint64, uint64, uint64, uint64)
func (_DIANFTOracle *DIANFTOracleCaller) GetValue(opts *bind.CallOpts, key string) (uint64, uint64, uint64, uint64, uint64, uint64, error) {
	var out []interface{}
	err := _DIANFTOracle.contract.Call(opts, &out, "getValue", key)

	if err != nil {
		return *new(uint64), *new(uint64), *new(uint64), *new(uint64), *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new(uint64)).(*uint64)
	out4 := *abi.ConvertType(out[4], new(uint64)).(*uint64)
	out5 := *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return out0, out1, out2, out3, out4, out5, err

}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint64, uint64, uint64, uint64, uint64, uint64)
func (_DIANFTOracle *DIANFTOracleSession) GetValue(key string) (uint64, uint64, uint64, uint64, uint64, uint64, error) {
	return _DIANFTOracle.Contract.GetValue(&_DIANFTOracle.CallOpts, key)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint64, uint64, uint64, uint64, uint64, uint64)
func (_DIANFTOracle *DIANFTOracleCallerSession) GetValue(key string) (uint64, uint64, uint64, uint64, uint64, uint64, error) {
	return _DIANFTOracle.Contract.GetValue(&_DIANFTOracle.CallOpts, key)
}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256 value0, uint256 value1)
func (_DIANFTOracle *DIANFTOracleCaller) Values(opts *bind.CallOpts, arg0 string) (struct {
	Value0 *big.Int
	Value1 *big.Int
}, error) {
	var out []interface{}
	err := _DIANFTOracle.contract.Call(opts, &out, "values", arg0)

	outstruct := new(struct {
		Value0 *big.Int
		Value1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256 value0, uint256 value1)
func (_DIANFTOracle *DIANFTOracleSession) Values(arg0 string) (struct {
	Value0 *big.Int
	Value1 *big.Int
}, error) {
	return _DIANFTOracle.Contract.Values(&_DIANFTOracle.CallOpts, arg0)
}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256 value0, uint256 value1)
func (_DIANFTOracle *DIANFTOracleCallerSession) Values(arg0 string) (struct {
	Value0 *big.Int
	Value1 *big.Int
}, error) {
	return _DIANFTOracle.Contract.Values(&_DIANFTOracle.CallOpts, arg0)
}

// SetValue is a paid mutator transaction binding the contract method 0xe3b47da5.
//
// Solidity: function setValue(string key, uint64 value0, uint64 value1, uint64 value2, uint64 value3, uint64 value4, uint64 timestamp) returns()
func (_DIANFTOracle *DIANFTOracleTransactor) SetValue(opts *bind.TransactOpts, key string, value0 uint64, value1 uint64, value2 uint64, value3 uint64, value4 uint64, timestamp uint64) (*types.Transaction, error) {
	return _DIANFTOracle.contract.Transact(opts, "setValue", key, value0, value1, value2, value3, value4, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0xe3b47da5.
//
// Solidity: function setValue(string key, uint64 value0, uint64 value1, uint64 value2, uint64 value3, uint64 value4, uint64 timestamp) returns()
func (_DIANFTOracle *DIANFTOracleSession) SetValue(key string, value0 uint64, value1 uint64, value2 uint64, value3 uint64, value4 uint64, timestamp uint64) (*types.Transaction, error) {
	return _DIANFTOracle.Contract.SetValue(&_DIANFTOracle.TransactOpts, key, value0, value1, value2, value3, value4, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0xe3b47da5.
//
// Solidity: function setValue(string key, uint64 value0, uint64 value1, uint64 value2, uint64 value3, uint64 value4, uint64 timestamp) returns()
func (_DIANFTOracle *DIANFTOracleTransactorSession) SetValue(key string, value0 uint64, value1 uint64, value2 uint64, value3 uint64, value4 uint64, timestamp uint64) (*types.Transaction, error) {
	return _DIANFTOracle.Contract.SetValue(&_DIANFTOracle.TransactOpts, key, value0, value1, value2, value3, value4, timestamp)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DIANFTOracle *DIANFTOracleTransactor) UpdateOracleUpdaterAddress(opts *bind.TransactOpts, newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DIANFTOracle.contract.Transact(opts, "updateOracleUpdaterAddress", newOracleUpdaterAddress)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DIANFTOracle *DIANFTOracleSession) UpdateOracleUpdaterAddress(newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DIANFTOracle.Contract.UpdateOracleUpdaterAddress(&_DIANFTOracle.TransactOpts, newOracleUpdaterAddress)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DIANFTOracle *DIANFTOracleTransactorSession) UpdateOracleUpdaterAddress(newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DIANFTOracle.Contract.UpdateOracleUpdaterAddress(&_DIANFTOracle.TransactOpts, newOracleUpdaterAddress)
}

// DIANFTOracleOracleUpdateIterator is returned from FilterOracleUpdate and is used to iterate over the raw logs and unpacked data for OracleUpdate events raised by the DIANFTOracle contract.
type DIANFTOracleOracleUpdateIterator struct {
	Event *DIANFTOracleOracleUpdate // Event containing the contract specifics and raw log

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
func (it *DIANFTOracleOracleUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIANFTOracleOracleUpdate)
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
		it.Event = new(DIANFTOracleOracleUpdate)
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
func (it *DIANFTOracleOracleUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIANFTOracleOracleUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIANFTOracleOracleUpdate represents a OracleUpdate event raised by the DIANFTOracle contract.
type DIANFTOracleOracleUpdate struct {
	Key       string
	Value0    uint64
	Value1    uint64
	Value2    uint64
	Value3    uint64
	Value4    uint64
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleUpdate is a free log retrieval operation binding the contract event 0xab06b7d7540b163ffbe1538e49d670388b97188db0c1eb2fdf41417a73a3d827.
//
// Solidity: event OracleUpdate(string key, uint64 value0, uint64 value1, uint64 value2, uint64 value3, uint64 value4, uint64 timestamp)
func (_DIANFTOracle *DIANFTOracleFilterer) FilterOracleUpdate(opts *bind.FilterOpts) (*DIANFTOracleOracleUpdateIterator, error) {

	logs, sub, err := _DIANFTOracle.contract.FilterLogs(opts, "OracleUpdate")
	if err != nil {
		return nil, err
	}
	return &DIANFTOracleOracleUpdateIterator{contract: _DIANFTOracle.contract, event: "OracleUpdate", logs: logs, sub: sub}, nil
}

// WatchOracleUpdate is a free log subscription operation binding the contract event 0xab06b7d7540b163ffbe1538e49d670388b97188db0c1eb2fdf41417a73a3d827.
//
// Solidity: event OracleUpdate(string key, uint64 value0, uint64 value1, uint64 value2, uint64 value3, uint64 value4, uint64 timestamp)
func (_DIANFTOracle *DIANFTOracleFilterer) WatchOracleUpdate(opts *bind.WatchOpts, sink chan<- *DIANFTOracleOracleUpdate) (event.Subscription, error) {

	logs, sub, err := _DIANFTOracle.contract.WatchLogs(opts, "OracleUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIANFTOracleOracleUpdate)
				if err := _DIANFTOracle.contract.UnpackLog(event, "OracleUpdate", log); err != nil {
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

// ParseOracleUpdate is a log parse operation binding the contract event 0xab06b7d7540b163ffbe1538e49d670388b97188db0c1eb2fdf41417a73a3d827.
//
// Solidity: event OracleUpdate(string key, uint64 value0, uint64 value1, uint64 value2, uint64 value3, uint64 value4, uint64 timestamp)
func (_DIANFTOracle *DIANFTOracleFilterer) ParseOracleUpdate(log types.Log) (*DIANFTOracleOracleUpdate, error) {
	event := new(DIANFTOracleOracleUpdate)
	if err := _DIANFTOracle.contract.UnpackLog(event, "OracleUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DIANFTOracleUpdaterAddressChangeIterator is returned from FilterUpdaterAddressChange and is used to iterate over the raw logs and unpacked data for UpdaterAddressChange events raised by the DIANFTOracle contract.
type DIANFTOracleUpdaterAddressChangeIterator struct {
	Event *DIANFTOracleUpdaterAddressChange // Event containing the contract specifics and raw log

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
func (it *DIANFTOracleUpdaterAddressChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIANFTOracleUpdaterAddressChange)
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
		it.Event = new(DIANFTOracleUpdaterAddressChange)
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
func (it *DIANFTOracleUpdaterAddressChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIANFTOracleUpdaterAddressChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIANFTOracleUpdaterAddressChange represents a UpdaterAddressChange event raised by the DIANFTOracle contract.
type DIANFTOracleUpdaterAddressChange struct {
	NewUpdater common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdaterAddressChange is a free log retrieval operation binding the contract event 0x121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f.
//
// Solidity: event UpdaterAddressChange(address newUpdater)
func (_DIANFTOracle *DIANFTOracleFilterer) FilterUpdaterAddressChange(opts *bind.FilterOpts) (*DIANFTOracleUpdaterAddressChangeIterator, error) {

	logs, sub, err := _DIANFTOracle.contract.FilterLogs(opts, "UpdaterAddressChange")
	if err != nil {
		return nil, err
	}
	return &DIANFTOracleUpdaterAddressChangeIterator{contract: _DIANFTOracle.contract, event: "UpdaterAddressChange", logs: logs, sub: sub}, nil
}

// WatchUpdaterAddressChange is a free log subscription operation binding the contract event 0x121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f.
//
// Solidity: event UpdaterAddressChange(address newUpdater)
func (_DIANFTOracle *DIANFTOracleFilterer) WatchUpdaterAddressChange(opts *bind.WatchOpts, sink chan<- *DIANFTOracleUpdaterAddressChange) (event.Subscription, error) {

	logs, sub, err := _DIANFTOracle.contract.WatchLogs(opts, "UpdaterAddressChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIANFTOracleUpdaterAddressChange)
				if err := _DIANFTOracle.contract.UnpackLog(event, "UpdaterAddressChange", log); err != nil {
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
func (_DIANFTOracle *DIANFTOracleFilterer) ParseUpdaterAddressChange(log types.Log) (*DIANFTOracleUpdaterAddressChange, error) {
	event := new(DIANFTOracleUpdaterAddressChange)
	if err := _DIANFTOracle.contract.UnpackLog(event, "UpdaterAddressChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
