// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracleService

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

// DiaOracleABI is the input ABI used to generate the binding from.
const DiaOracleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"newPrice\",\"type\":\"uint256\"},{\"name\":\"newSupply\",\"type\":\"uint256\"},{\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"name\":\"updateCoinInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getCoinInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastUpdateTimestamp\",\"type\":\"uint256\"}],\"name\":\"newCoinInfo\",\"type\":\"event\"}]"

// DiaOracleBin is the compiled bytecode used for deploying new contracts.
const DiaOracleBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a0319163317905561076c806100326000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166345cc0a2f811461005b578063a6f9dae114610101578063e12214001461012f575b600080fd5b34801561006757600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100ff94369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750508435955050506020830135926040013591506102159050565b005b34801561010d57600080fd5b506100ff73ffffffffffffffffffffffffffffffffffffffff6004351661040d565b34801561013b57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261018894369492936024939284019190819084018382808284375094975061046d9650505050505050565b6040518085815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101d75781810151838201526020016101bf565b50505050905090810190601f1680156102045780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b60005473ffffffffffffffffffffffffffffffffffffffff16331461023957600080fd5b608060405190810160405280848152602001838152602001828152602001858152506001866040518082805190602001908083835b6020831061028d5780518252601f19909201916020918201910161026e565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820190942085518155858201516001820155938501516002850155606085015180516102ea94506003860193509101906106a5565b509050507f8d468b5f823f8d38322e9c4433d184adf453fd3eaa28cef280056aa0664981f08585858585604051808060200180602001868152602001858152602001848152602001838103835288818151815260200191508051906020019080838360005b8381101561036757818101518382015260200161034f565b50505050905090810190601f1680156103945780820380516001836020036101000a031916815260200191505b50838103825287518152875160209182019189019080838360005b838110156103c75781810151838201526020016103af565b50505050905090810190601f1680156103f45780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461043157600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600080600060606001856040518082805190602001908083835b602083106104a65780518252601f199092019160209182019101610487565b51815160209384036101000a600019018019909216911617905292019485525060405193849003810184205489519094600194508a9350918291908401908083835b602083106105075780518252601f1990920191602091820191016104e8565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600101546001876040518082805190602001908083835b602083106105705780518252601f199092019160209182019101610551565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381018420600201548b519094600194508c9350918291908401908083835b602083106105d45780518252601f1990920191602091820191016105b5565b518151600019602094850361010090810a820192831692199390931691909117909252949092019687526040805197889003820188206003018054601f60026001831615909802909501169590950492830182900482028801820190528187529295945085935091840190508282801561068f5780601f106106645761010080835404028352916020019161068f565b820191906000526020600020905b81548152906001019060200180831161067257829003601f168201915b5050505050905093509350935093509193509193565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106e657805160ff1916838001178555610713565b82800160010185558215610713579182015b828111156107135782518255916020019190600101906106f8565b5061071f929150610723565b5090565b61073d91905b8082111561071f5760008155600101610729565b905600a165627a7a723058205ef49ea0f922138149ae12d6cfbae483bd69686515f60eb73c3e50e8920cfeb70029`

// DeployDiaOracle deploys a new Ethereum contract, binding an instance of DiaOracle to it.
func DeployDiaOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DiaOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(DiaOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DiaOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DiaOracle{DiaOracleCaller: DiaOracleCaller{contract: contract}, DiaOracleTransactor: DiaOracleTransactor{contract: contract}, DiaOracleFilterer: DiaOracleFilterer{contract: contract}}, nil
}

// DiaOracle is an auto generated Go binding around an Ethereum contract.
type DiaOracle struct {
	DiaOracleCaller     // Read-only binding to the contract
	DiaOracleTransactor // Write-only binding to the contract
	DiaOracleFilterer   // Log filterer for contract events
}

// DiaOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiaOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiaOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiaOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiaOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiaOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiaOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiaOracleSession struct {
	Contract     *DiaOracle        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiaOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiaOracleCallerSession struct {
	Contract *DiaOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DiaOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiaOracleTransactorSession struct {
	Contract     *DiaOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DiaOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiaOracleRaw struct {
	Contract *DiaOracle // Generic contract binding to access the raw methods on
}

// DiaOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiaOracleCallerRaw struct {
	Contract *DiaOracleCaller // Generic read-only contract binding to access the raw methods on
}

// DiaOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiaOracleTransactorRaw struct {
	Contract *DiaOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiaOracle creates a new instance of DiaOracle, bound to a specific deployed contract.
func NewDiaOracle(address common.Address, backend bind.ContractBackend) (*DiaOracle, error) {
	contract, err := bindDiaOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiaOracle{DiaOracleCaller: DiaOracleCaller{contract: contract}, DiaOracleTransactor: DiaOracleTransactor{contract: contract}, DiaOracleFilterer: DiaOracleFilterer{contract: contract}}, nil
}

// NewDiaOracleCaller creates a new read-only instance of DiaOracle, bound to a specific deployed contract.
func NewDiaOracleCaller(address common.Address, caller bind.ContractCaller) (*DiaOracleCaller, error) {
	contract, err := bindDiaOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiaOracleCaller{contract: contract}, nil
}

// NewDiaOracleTransactor creates a new write-only instance of DiaOracle, bound to a specific deployed contract.
func NewDiaOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*DiaOracleTransactor, error) {
	contract, err := bindDiaOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiaOracleTransactor{contract: contract}, nil
}

// NewDiaOracleFilterer creates a new log filterer instance of DiaOracle, bound to a specific deployed contract.
func NewDiaOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*DiaOracleFilterer, error) {
	contract, err := bindDiaOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiaOracleFilterer{contract: contract}, nil
}

// bindDiaOracle binds a generic wrapper to an already deployed contract.
func bindDiaOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiaOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiaOracle *DiaOracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DiaOracle.Contract.DiaOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiaOracle *DiaOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiaOracle.Contract.DiaOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiaOracle *DiaOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiaOracle.Contract.DiaOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiaOracle *DiaOracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DiaOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiaOracle *DiaOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiaOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiaOracle *DiaOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiaOracle.Contract.contract.Transact(opts, method, params...)
}

// GetCoinInfo is a free data retrieval call binding the contract method 0xe1221400.
//
// Solidity: function getCoinInfo(name string) constant returns(uint256, uint256, uint256, string)
func (_DiaOracle *DiaOracleCaller) GetCoinInfo(opts *bind.CallOpts, name string) (*big.Int, *big.Int, *big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _DiaOracle.contract.Call(opts, out, "getCoinInfo", name)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetCoinInfo is a free data retrieval call binding the contract method 0xe1221400.
//
// Solidity: function getCoinInfo(name string) constant returns(uint256, uint256, uint256, string)
func (_DiaOracle *DiaOracleSession) GetCoinInfo(name string) (*big.Int, *big.Int, *big.Int, string, error) {
	return _DiaOracle.Contract.GetCoinInfo(&_DiaOracle.CallOpts, name)
}

// GetCoinInfo is a free data retrieval call binding the contract method 0xe1221400.
//
// Solidity: function getCoinInfo(name string) constant returns(uint256, uint256, uint256, string)
func (_DiaOracle *DiaOracleCallerSession) GetCoinInfo(name string) (*big.Int, *big.Int, *big.Int, string, error) {
	return _DiaOracle.Contract.GetCoinInfo(&_DiaOracle.CallOpts, name)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_DiaOracle *DiaOracleTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DiaOracle.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_DiaOracle *DiaOracleSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _DiaOracle.Contract.ChangeOwner(&_DiaOracle.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_DiaOracle *DiaOracleTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _DiaOracle.Contract.ChangeOwner(&_DiaOracle.TransactOpts, newOwner)
}

// UpdateCoinInfo is a paid mutator transaction binding the contract method 0x45cc0a2f.
//
// Solidity: function updateCoinInfo(name string, symbol string, newPrice uint256, newSupply uint256, newTimestamp uint256) returns()
func (_DiaOracle *DiaOracleTransactor) UpdateCoinInfo(opts *bind.TransactOpts, name string, symbol string, newPrice *big.Int, newSupply *big.Int, newTimestamp *big.Int) (*types.Transaction, error) {
	return _DiaOracle.contract.Transact(opts, "updateCoinInfo", name, symbol, newPrice, newSupply, newTimestamp)
}

// UpdateCoinInfo is a paid mutator transaction binding the contract method 0x45cc0a2f.
//
// Solidity: function updateCoinInfo(name string, symbol string, newPrice uint256, newSupply uint256, newTimestamp uint256) returns()
func (_DiaOracle *DiaOracleSession) UpdateCoinInfo(name string, symbol string, newPrice *big.Int, newSupply *big.Int, newTimestamp *big.Int) (*types.Transaction, error) {
	return _DiaOracle.Contract.UpdateCoinInfo(&_DiaOracle.TransactOpts, name, symbol, newPrice, newSupply, newTimestamp)
}

// UpdateCoinInfo is a paid mutator transaction binding the contract method 0x45cc0a2f.
//
// Solidity: function updateCoinInfo(name string, symbol string, newPrice uint256, newSupply uint256, newTimestamp uint256) returns()
func (_DiaOracle *DiaOracleTransactorSession) UpdateCoinInfo(name string, symbol string, newPrice *big.Int, newSupply *big.Int, newTimestamp *big.Int) (*types.Transaction, error) {
	return _DiaOracle.Contract.UpdateCoinInfo(&_DiaOracle.TransactOpts, name, symbol, newPrice, newSupply, newTimestamp)
}

// DiaOracleNewCoinInfoIterator is returned from FilterNewCoinInfo and is used to iterate over the raw logs and unpacked data for NewCoinInfo events raised by the DiaOracle contract.
type DiaOracleNewCoinInfoIterator struct {
	Event *DiaOracleNewCoinInfo // Event containing the contract specifics and raw log

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
func (it *DiaOracleNewCoinInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiaOracleNewCoinInfo)
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
		it.Event = new(DiaOracleNewCoinInfo)
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
func (it *DiaOracleNewCoinInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiaOracleNewCoinInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiaOracleNewCoinInfo represents a NewCoinInfo event raised by the DiaOracle contract.
type DiaOracleNewCoinInfo struct {
	Name                string
	Symbol              string
	Price               *big.Int
	Supply              *big.Int
	LastUpdateTimestamp *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewCoinInfo is a free log retrieval operation binding the contract event 0x8d468b5f823f8d38322e9c4433d184adf453fd3eaa28cef280056aa0664981f0.
//
// Solidity: e newCoinInfo(name string, symbol string, price uint256, supply uint256, lastUpdateTimestamp uint256)
func (_DiaOracle *DiaOracleFilterer) FilterNewCoinInfo(opts *bind.FilterOpts) (*DiaOracleNewCoinInfoIterator, error) {

	logs, sub, err := _DiaOracle.contract.FilterLogs(opts, "newCoinInfo")
	if err != nil {
		return nil, err
	}
	return &DiaOracleNewCoinInfoIterator{contract: _DiaOracle.contract, event: "newCoinInfo", logs: logs, sub: sub}, nil
}

// WatchNewCoinInfo is a free log subscription operation binding the contract event 0x8d468b5f823f8d38322e9c4433d184adf453fd3eaa28cef280056aa0664981f0.
//
// Solidity: e newCoinInfo(name string, symbol string, price uint256, supply uint256, lastUpdateTimestamp uint256)
func (_DiaOracle *DiaOracleFilterer) WatchNewCoinInfo(opts *bind.WatchOpts, sink chan<- *DiaOracleNewCoinInfo) (event.Subscription, error) {

	logs, sub, err := _DiaOracle.contract.WatchLogs(opts, "newCoinInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiaOracleNewCoinInfo)
				if err := _DiaOracle.contract.UnpackLog(event, "newCoinInfo", log); err != nil {
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
