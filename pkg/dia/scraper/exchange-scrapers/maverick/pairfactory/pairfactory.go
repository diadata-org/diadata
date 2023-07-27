// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pairfactory

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

// PairfactoryMetaData contains all meta data concerning the Pairfactory contract.
var PairfactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tickSpacing\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int32\",\"name\":\"activeTick\",\"type\":\"int32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"lookback\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"protocolFeeRatio\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SetFactoryOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"protocolFeeRatio\",\"type\":\"uint64\"}],\"name\":\"SetFactoryProtocolFeeRatio\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tickSpacing\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_lookback\",\"type\":\"int256\"},{\"internalType\":\"int32\",\"name\":\"_activeTick\",\"type\":\"int32\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenB\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isFactoryPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tickSpacing\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"lookback\",\"type\":\"int256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"lookup\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"position\",\"outputs\":[{\"internalType\":\"contractIPosition\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeRatio\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PairfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PairfactoryMetaData.ABI instead.
var PairfactoryABI = PairfactoryMetaData.ABI

// Pairfactory is an auto generated Go binding around an Ethereum contract.
type Pairfactory struct {
	PairfactoryCaller     // Read-only binding to the contract
	PairfactoryTransactor // Write-only binding to the contract
	PairfactoryFilterer   // Log filterer for contract events
}

// PairfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairfactorySession struct {
	Contract     *Pairfactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairfactoryCallerSession struct {
	Contract *PairfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PairfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairfactoryTransactorSession struct {
	Contract     *PairfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PairfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairfactoryRaw struct {
	Contract *Pairfactory // Generic contract binding to access the raw methods on
}

// PairfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairfactoryCallerRaw struct {
	Contract *PairfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PairfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairfactoryTransactorRaw struct {
	Contract *PairfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairfactory creates a new instance of Pairfactory, bound to a specific deployed contract.
func NewPairfactory(address common.Address, backend bind.ContractBackend) (*Pairfactory, error) {
	contract, err := bindPairfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pairfactory{PairfactoryCaller: PairfactoryCaller{contract: contract}, PairfactoryTransactor: PairfactoryTransactor{contract: contract}, PairfactoryFilterer: PairfactoryFilterer{contract: contract}}, nil
}

// NewPairfactoryCaller creates a new read-only instance of Pairfactory, bound to a specific deployed contract.
func NewPairfactoryCaller(address common.Address, caller bind.ContractCaller) (*PairfactoryCaller, error) {
	contract, err := bindPairfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairfactoryCaller{contract: contract}, nil
}

// NewPairfactoryTransactor creates a new write-only instance of Pairfactory, bound to a specific deployed contract.
func NewPairfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PairfactoryTransactor, error) {
	contract, err := bindPairfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairfactoryTransactor{contract: contract}, nil
}

// NewPairfactoryFilterer creates a new log filterer instance of Pairfactory, bound to a specific deployed contract.
func NewPairfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PairfactoryFilterer, error) {
	contract, err := bindPairfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairfactoryFilterer{contract: contract}, nil
}

// bindPairfactory binds a generic wrapper to an already deployed contract.
func bindPairfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PairfactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairfactory *PairfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairfactory.Contract.PairfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairfactory *PairfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairfactory.Contract.PairfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairfactory *PairfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairfactory.Contract.PairfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairfactory *PairfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairfactory *PairfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairfactory *PairfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairfactory.Contract.contract.Transact(opts, method, params...)
}

// IsFactoryPool is a free data retrieval call binding the contract method 0x578eaca4.
//
// Solidity: function isFactoryPool(address pool) view returns(bool)
func (_Pairfactory *PairfactoryCaller) IsFactoryPool(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _Pairfactory.contract.Call(opts, &out, "isFactoryPool", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFactoryPool is a free data retrieval call binding the contract method 0x578eaca4.
//
// Solidity: function isFactoryPool(address pool) view returns(bool)
func (_Pairfactory *PairfactorySession) IsFactoryPool(pool common.Address) (bool, error) {
	return _Pairfactory.Contract.IsFactoryPool(&_Pairfactory.CallOpts, pool)
}

// IsFactoryPool is a free data retrieval call binding the contract method 0x578eaca4.
//
// Solidity: function isFactoryPool(address pool) view returns(bool)
func (_Pairfactory *PairfactoryCallerSession) IsFactoryPool(pool common.Address) (bool, error) {
	return _Pairfactory.Contract.IsFactoryPool(&_Pairfactory.CallOpts, pool)
}

// Lookup is a free data retrieval call binding the contract method 0xc697217a.
//
// Solidity: function lookup(uint256 fee, uint256 tickSpacing, int256 lookback, address tokenA, address tokenB) view returns(address)
func (_Pairfactory *PairfactoryCaller) Lookup(opts *bind.CallOpts, fee *big.Int, tickSpacing *big.Int, lookback *big.Int, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _Pairfactory.contract.Call(opts, &out, "lookup", fee, tickSpacing, lookback, tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lookup is a free data retrieval call binding the contract method 0xc697217a.
//
// Solidity: function lookup(uint256 fee, uint256 tickSpacing, int256 lookback, address tokenA, address tokenB) view returns(address)
func (_Pairfactory *PairfactorySession) Lookup(fee *big.Int, tickSpacing *big.Int, lookback *big.Int, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Pairfactory.Contract.Lookup(&_Pairfactory.CallOpts, fee, tickSpacing, lookback, tokenA, tokenB)
}

// Lookup is a free data retrieval call binding the contract method 0xc697217a.
//
// Solidity: function lookup(uint256 fee, uint256 tickSpacing, int256 lookback, address tokenA, address tokenB) view returns(address)
func (_Pairfactory *PairfactoryCallerSession) Lookup(fee *big.Int, tickSpacing *big.Int, lookback *big.Int, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Pairfactory.Contract.Lookup(&_Pairfactory.CallOpts, fee, tickSpacing, lookback, tokenA, tokenB)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pairfactory *PairfactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pairfactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pairfactory *PairfactorySession) Owner() (common.Address, error) {
	return _Pairfactory.Contract.Owner(&_Pairfactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pairfactory *PairfactoryCallerSession) Owner() (common.Address, error) {
	return _Pairfactory.Contract.Owner(&_Pairfactory.CallOpts)
}

// Position is a free data retrieval call binding the contract method 0x09218e91.
//
// Solidity: function position() view returns(address)
func (_Pairfactory *PairfactoryCaller) Position(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pairfactory.contract.Call(opts, &out, "position")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Position is a free data retrieval call binding the contract method 0x09218e91.
//
// Solidity: function position() view returns(address)
func (_Pairfactory *PairfactorySession) Position() (common.Address, error) {
	return _Pairfactory.Contract.Position(&_Pairfactory.CallOpts)
}

// Position is a free data retrieval call binding the contract method 0x09218e91.
//
// Solidity: function position() view returns(address)
func (_Pairfactory *PairfactoryCallerSession) Position() (common.Address, error) {
	return _Pairfactory.Contract.Position(&_Pairfactory.CallOpts)
}

// ProtocolFeeRatio is a free data retrieval call binding the contract method 0x7a27d9f6.
//
// Solidity: function protocolFeeRatio() view returns(uint64)
func (_Pairfactory *PairfactoryCaller) ProtocolFeeRatio(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Pairfactory.contract.Call(opts, &out, "protocolFeeRatio")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ProtocolFeeRatio is a free data retrieval call binding the contract method 0x7a27d9f6.
//
// Solidity: function protocolFeeRatio() view returns(uint64)
func (_Pairfactory *PairfactorySession) ProtocolFeeRatio() (uint64, error) {
	return _Pairfactory.Contract.ProtocolFeeRatio(&_Pairfactory.CallOpts)
}

// ProtocolFeeRatio is a free data retrieval call binding the contract method 0x7a27d9f6.
//
// Solidity: function protocolFeeRatio() view returns(uint64)
func (_Pairfactory *PairfactoryCallerSession) ProtocolFeeRatio() (uint64, error) {
	return _Pairfactory.Contract.ProtocolFeeRatio(&_Pairfactory.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0x71861ede.
//
// Solidity: function create(uint256 _fee, uint256 _tickSpacing, int256 _lookback, int32 _activeTick, address _tokenA, address _tokenB) returns(address)
func (_Pairfactory *PairfactoryTransactor) Create(opts *bind.TransactOpts, _fee *big.Int, _tickSpacing *big.Int, _lookback *big.Int, _activeTick int32, _tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _Pairfactory.contract.Transact(opts, "create", _fee, _tickSpacing, _lookback, _activeTick, _tokenA, _tokenB)
}

// Create is a paid mutator transaction binding the contract method 0x71861ede.
//
// Solidity: function create(uint256 _fee, uint256 _tickSpacing, int256 _lookback, int32 _activeTick, address _tokenA, address _tokenB) returns(address)
func (_Pairfactory *PairfactorySession) Create(_fee *big.Int, _tickSpacing *big.Int, _lookback *big.Int, _activeTick int32, _tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _Pairfactory.Contract.Create(&_Pairfactory.TransactOpts, _fee, _tickSpacing, _lookback, _activeTick, _tokenA, _tokenB)
}

// Create is a paid mutator transaction binding the contract method 0x71861ede.
//
// Solidity: function create(uint256 _fee, uint256 _tickSpacing, int256 _lookback, int32 _activeTick, address _tokenA, address _tokenB) returns(address)
func (_Pairfactory *PairfactoryTransactorSession) Create(_fee *big.Int, _tickSpacing *big.Int, _lookback *big.Int, _activeTick int32, _tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _Pairfactory.Contract.Create(&_Pairfactory.TransactOpts, _fee, _tickSpacing, _lookback, _activeTick, _tokenA, _tokenB)
}

// PairfactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the Pairfactory contract.
type PairfactoryPoolCreatedIterator struct {
	Event *PairfactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *PairfactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairfactoryPoolCreated)
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
		it.Event = new(PairfactoryPoolCreated)
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
func (it *PairfactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairfactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairfactoryPoolCreated represents a PoolCreated event raised by the Pairfactory contract.
type PairfactoryPoolCreated struct {
	PoolAddress      common.Address
	Fee              *big.Int
	TickSpacing      *big.Int
	ActiveTick       int32
	Lookback         *big.Int
	ProtocolFeeRatio uint64
	TokenA           common.Address
	TokenB           common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x9b3fb3a17b4e94eb4d1217257372dcc712218fcd4bc1c28482bd8a6804a7c775.
//
// Solidity: event PoolCreated(address poolAddress, uint256 fee, uint256 tickSpacing, int32 activeTick, int256 lookback, uint64 protocolFeeRatio, address tokenA, address tokenB)
func (_Pairfactory *PairfactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts) (*PairfactoryPoolCreatedIterator, error) {

	logs, sub, err := _Pairfactory.contract.FilterLogs(opts, "PoolCreated")
	if err != nil {
		return nil, err
	}
	return &PairfactoryPoolCreatedIterator{contract: _Pairfactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x9b3fb3a17b4e94eb4d1217257372dcc712218fcd4bc1c28482bd8a6804a7c775.
//
// Solidity: event PoolCreated(address poolAddress, uint256 fee, uint256 tickSpacing, int32 activeTick, int256 lookback, uint64 protocolFeeRatio, address tokenA, address tokenB)
func (_Pairfactory *PairfactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *PairfactoryPoolCreated) (event.Subscription, error) {

	logs, sub, err := _Pairfactory.contract.WatchLogs(opts, "PoolCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairfactoryPoolCreated)
				if err := _Pairfactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x9b3fb3a17b4e94eb4d1217257372dcc712218fcd4bc1c28482bd8a6804a7c775.
//
// Solidity: event PoolCreated(address poolAddress, uint256 fee, uint256 tickSpacing, int32 activeTick, int256 lookback, uint64 protocolFeeRatio, address tokenA, address tokenB)
func (_Pairfactory *PairfactoryFilterer) ParsePoolCreated(log types.Log) (*PairfactoryPoolCreated, error) {
	event := new(PairfactoryPoolCreated)
	if err := _Pairfactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairfactorySetFactoryOwnerIterator is returned from FilterSetFactoryOwner and is used to iterate over the raw logs and unpacked data for SetFactoryOwner events raised by the Pairfactory contract.
type PairfactorySetFactoryOwnerIterator struct {
	Event *PairfactorySetFactoryOwner // Event containing the contract specifics and raw log

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
func (it *PairfactorySetFactoryOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairfactorySetFactoryOwner)
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
		it.Event = new(PairfactorySetFactoryOwner)
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
func (it *PairfactorySetFactoryOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairfactorySetFactoryOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairfactorySetFactoryOwner represents a SetFactoryOwner event raised by the Pairfactory contract.
type PairfactorySetFactoryOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetFactoryOwner is a free log retrieval operation binding the contract event 0xdbb9c2f238268c9ccb9adfe9e27e349c8d4841cca1d2ffe849ae9114177b8fc1.
//
// Solidity: event SetFactoryOwner(address owner)
func (_Pairfactory *PairfactoryFilterer) FilterSetFactoryOwner(opts *bind.FilterOpts) (*PairfactorySetFactoryOwnerIterator, error) {

	logs, sub, err := _Pairfactory.contract.FilterLogs(opts, "SetFactoryOwner")
	if err != nil {
		return nil, err
	}
	return &PairfactorySetFactoryOwnerIterator{contract: _Pairfactory.contract, event: "SetFactoryOwner", logs: logs, sub: sub}, nil
}

// WatchSetFactoryOwner is a free log subscription operation binding the contract event 0xdbb9c2f238268c9ccb9adfe9e27e349c8d4841cca1d2ffe849ae9114177b8fc1.
//
// Solidity: event SetFactoryOwner(address owner)
func (_Pairfactory *PairfactoryFilterer) WatchSetFactoryOwner(opts *bind.WatchOpts, sink chan<- *PairfactorySetFactoryOwner) (event.Subscription, error) {

	logs, sub, err := _Pairfactory.contract.WatchLogs(opts, "SetFactoryOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairfactorySetFactoryOwner)
				if err := _Pairfactory.contract.UnpackLog(event, "SetFactoryOwner", log); err != nil {
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

// ParseSetFactoryOwner is a log parse operation binding the contract event 0xdbb9c2f238268c9ccb9adfe9e27e349c8d4841cca1d2ffe849ae9114177b8fc1.
//
// Solidity: event SetFactoryOwner(address owner)
func (_Pairfactory *PairfactoryFilterer) ParseSetFactoryOwner(log types.Log) (*PairfactorySetFactoryOwner, error) {
	event := new(PairfactorySetFactoryOwner)
	if err := _Pairfactory.contract.UnpackLog(event, "SetFactoryOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairfactorySetFactoryProtocolFeeRatioIterator is returned from FilterSetFactoryProtocolFeeRatio and is used to iterate over the raw logs and unpacked data for SetFactoryProtocolFeeRatio events raised by the Pairfactory contract.
type PairfactorySetFactoryProtocolFeeRatioIterator struct {
	Event *PairfactorySetFactoryProtocolFeeRatio // Event containing the contract specifics and raw log

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
func (it *PairfactorySetFactoryProtocolFeeRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairfactorySetFactoryProtocolFeeRatio)
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
		it.Event = new(PairfactorySetFactoryProtocolFeeRatio)
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
func (it *PairfactorySetFactoryProtocolFeeRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairfactorySetFactoryProtocolFeeRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairfactorySetFactoryProtocolFeeRatio represents a SetFactoryProtocolFeeRatio event raised by the Pairfactory contract.
type PairfactorySetFactoryProtocolFeeRatio struct {
	ProtocolFeeRatio uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetFactoryProtocolFeeRatio is a free log retrieval operation binding the contract event 0xa0b96532ac38eaae102ba1ac8b6cde2019d7170e8a984863a7755fda6408a55f.
//
// Solidity: event SetFactoryProtocolFeeRatio(uint64 protocolFeeRatio)
func (_Pairfactory *PairfactoryFilterer) FilterSetFactoryProtocolFeeRatio(opts *bind.FilterOpts) (*PairfactorySetFactoryProtocolFeeRatioIterator, error) {

	logs, sub, err := _Pairfactory.contract.FilterLogs(opts, "SetFactoryProtocolFeeRatio")
	if err != nil {
		return nil, err
	}
	return &PairfactorySetFactoryProtocolFeeRatioIterator{contract: _Pairfactory.contract, event: "SetFactoryProtocolFeeRatio", logs: logs, sub: sub}, nil
}

// WatchSetFactoryProtocolFeeRatio is a free log subscription operation binding the contract event 0xa0b96532ac38eaae102ba1ac8b6cde2019d7170e8a984863a7755fda6408a55f.
//
// Solidity: event SetFactoryProtocolFeeRatio(uint64 protocolFeeRatio)
func (_Pairfactory *PairfactoryFilterer) WatchSetFactoryProtocolFeeRatio(opts *bind.WatchOpts, sink chan<- *PairfactorySetFactoryProtocolFeeRatio) (event.Subscription, error) {

	logs, sub, err := _Pairfactory.contract.WatchLogs(opts, "SetFactoryProtocolFeeRatio")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairfactorySetFactoryProtocolFeeRatio)
				if err := _Pairfactory.contract.UnpackLog(event, "SetFactoryProtocolFeeRatio", log); err != nil {
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

// ParseSetFactoryProtocolFeeRatio is a log parse operation binding the contract event 0xa0b96532ac38eaae102ba1ac8b6cde2019d7170e8a984863a7755fda6408a55f.
//
// Solidity: event SetFactoryProtocolFeeRatio(uint64 protocolFeeRatio)
func (_Pairfactory *PairfactoryFilterer) ParseSetFactoryProtocolFeeRatio(log types.Log) (*PairfactorySetFactoryProtocolFeeRatio, error) {
	event := new(PairfactorySetFactoryProtocolFeeRatio)
	if err := _Pairfactory.contract.UnpackLog(event, "SetFactoryProtocolFeeRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
