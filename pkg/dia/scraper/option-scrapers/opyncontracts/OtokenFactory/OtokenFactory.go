// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OtokenFactory

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

// OtokenFactoryABI is the input ABI used to generate the binding from.
const OtokenFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressBook\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strike\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPut\",\"type\":\"bool\"}],\"name\":\"OtokenCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressBook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_strikeAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isPut\",\"type\":\"bool\"}],\"name\":\"createOtoken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_strikeAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isPut\",\"type\":\"bool\"}],\"name\":\"getOtoken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOtokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_strikeAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isPut\",\"type\":\"bool\"}],\"name\":\"getTargetOtokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"otokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OtokenFactory is an auto generated Go binding around an Ethereum contract.
type OtokenFactory struct {
	OtokenFactoryCaller     // Read-only binding to the contract
	OtokenFactoryTransactor // Write-only binding to the contract
	OtokenFactoryFilterer   // Log filterer for contract events
}

// OtokenFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OtokenFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtokenFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OtokenFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtokenFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OtokenFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtokenFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OtokenFactorySession struct {
	Contract     *OtokenFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OtokenFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OtokenFactoryCallerSession struct {
	Contract *OtokenFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OtokenFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OtokenFactoryTransactorSession struct {
	Contract     *OtokenFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OtokenFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OtokenFactoryRaw struct {
	Contract *OtokenFactory // Generic contract binding to access the raw methods on
}

// OtokenFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OtokenFactoryCallerRaw struct {
	Contract *OtokenFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// OtokenFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OtokenFactoryTransactorRaw struct {
	Contract *OtokenFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOtokenFactory creates a new instance of OtokenFactory, bound to a specific deployed contract.
func NewOtokenFactory(address common.Address, backend bind.ContractBackend) (*OtokenFactory, error) {
	contract, err := bindOtokenFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OtokenFactory{OtokenFactoryCaller: OtokenFactoryCaller{contract: contract}, OtokenFactoryTransactor: OtokenFactoryTransactor{contract: contract}, OtokenFactoryFilterer: OtokenFactoryFilterer{contract: contract}}, nil
}

// NewOtokenFactoryCaller creates a new read-only instance of OtokenFactory, bound to a specific deployed contract.
func NewOtokenFactoryCaller(address common.Address, caller bind.ContractCaller) (*OtokenFactoryCaller, error) {
	contract, err := bindOtokenFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OtokenFactoryCaller{contract: contract}, nil
}

// NewOtokenFactoryTransactor creates a new write-only instance of OtokenFactory, bound to a specific deployed contract.
func NewOtokenFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*OtokenFactoryTransactor, error) {
	contract, err := bindOtokenFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OtokenFactoryTransactor{contract: contract}, nil
}

// NewOtokenFactoryFilterer creates a new log filterer instance of OtokenFactory, bound to a specific deployed contract.
func NewOtokenFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*OtokenFactoryFilterer, error) {
	contract, err := bindOtokenFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OtokenFactoryFilterer{contract: contract}, nil
}

// bindOtokenFactory binds a generic wrapper to an already deployed contract.
func bindOtokenFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OtokenFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OtokenFactory *OtokenFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OtokenFactory.Contract.OtokenFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OtokenFactory *OtokenFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtokenFactory.Contract.OtokenFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OtokenFactory *OtokenFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OtokenFactory.Contract.OtokenFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OtokenFactory *OtokenFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OtokenFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OtokenFactory *OtokenFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtokenFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OtokenFactory *OtokenFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OtokenFactory.Contract.contract.Transact(opts, method, params...)
}

// AddressBook is a free data retrieval call binding the contract method 0xf5887cdd.
//
// Solidity: function addressBook() view returns(address)
func (_OtokenFactory *OtokenFactoryCaller) AddressBook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OtokenFactory.contract.Call(opts, &out, "addressBook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressBook is a free data retrieval call binding the contract method 0xf5887cdd.
//
// Solidity: function addressBook() view returns(address)
func (_OtokenFactory *OtokenFactorySession) AddressBook() (common.Address, error) {
	return _OtokenFactory.Contract.AddressBook(&_OtokenFactory.CallOpts)
}

// AddressBook is a free data retrieval call binding the contract method 0xf5887cdd.
//
// Solidity: function addressBook() view returns(address)
func (_OtokenFactory *OtokenFactoryCallerSession) AddressBook() (common.Address, error) {
	return _OtokenFactory.Contract.AddressBook(&_OtokenFactory.CallOpts)
}

// GetOtoken is a free data retrieval call binding the contract method 0x11547054.
//
// Solidity: function getOtoken(address _underlyingAsset, address _strikeAsset, address _collateralAsset, uint256 _strikePrice, uint256 _expiry, bool _isPut) view returns(address)
func (_OtokenFactory *OtokenFactoryCaller) GetOtoken(opts *bind.CallOpts, _underlyingAsset common.Address, _strikeAsset common.Address, _collateralAsset common.Address, _strikePrice *big.Int, _expiry *big.Int, _isPut bool) (common.Address, error) {
	var out []interface{}
	err := _OtokenFactory.contract.Call(opts, &out, "getOtoken", _underlyingAsset, _strikeAsset, _collateralAsset, _strikePrice, _expiry, _isPut)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOtoken is a free data retrieval call binding the contract method 0x11547054.
//
// Solidity: function getOtoken(address _underlyingAsset, address _strikeAsset, address _collateralAsset, uint256 _strikePrice, uint256 _expiry, bool _isPut) view returns(address)
func (_OtokenFactory *OtokenFactorySession) GetOtoken(_underlyingAsset common.Address, _strikeAsset common.Address, _collateralAsset common.Address, _strikePrice *big.Int, _expiry *big.Int, _isPut bool) (common.Address, error) {
	return _OtokenFactory.Contract.GetOtoken(&_OtokenFactory.CallOpts, _underlyingAsset, _strikeAsset, _collateralAsset, _strikePrice, _expiry, _isPut)
}

// GetOtoken is a free data retrieval call binding the contract method 0x11547054.
//
// Solidity: function getOtoken(address _underlyingAsset, address _strikeAsset, address _collateralAsset, uint256 _strikePrice, uint256 _expiry, bool _isPut) view returns(address)
func (_OtokenFactory *OtokenFactoryCallerSession) GetOtoken(_underlyingAsset common.Address, _strikeAsset common.Address, _collateralAsset common.Address, _strikePrice *big.Int, _expiry *big.Int, _isPut bool) (common.Address, error) {
	return _OtokenFactory.Contract.GetOtoken(&_OtokenFactory.CallOpts, _underlyingAsset, _strikeAsset, _collateralAsset, _strikePrice, _expiry, _isPut)
}

// GetOtokensLength is a free data retrieval call binding the contract method 0x46e63dc4.
//
// Solidity: function getOtokensLength() view returns(uint256)
func (_OtokenFactory *OtokenFactoryCaller) GetOtokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OtokenFactory.contract.Call(opts, &out, "getOtokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOtokensLength is a free data retrieval call binding the contract method 0x46e63dc4.
//
// Solidity: function getOtokensLength() view returns(uint256)
func (_OtokenFactory *OtokenFactorySession) GetOtokensLength() (*big.Int, error) {
	return _OtokenFactory.Contract.GetOtokensLength(&_OtokenFactory.CallOpts)
}

// GetOtokensLength is a free data retrieval call binding the contract method 0x46e63dc4.
//
// Solidity: function getOtokensLength() view returns(uint256)
func (_OtokenFactory *OtokenFactoryCallerSession) GetOtokensLength() (*big.Int, error) {
	return _OtokenFactory.Contract.GetOtokensLength(&_OtokenFactory.CallOpts)
}

// GetTargetOtokenAddress is a free data retrieval call binding the contract method 0xb86b9a64.
//
// Solidity: function getTargetOtokenAddress(address _underlyingAsset, address _strikeAsset, address _collateralAsset, uint256 _strikePrice, uint256 _expiry, bool _isPut) view returns(address)
func (_OtokenFactory *OtokenFactoryCaller) GetTargetOtokenAddress(opts *bind.CallOpts, _underlyingAsset common.Address, _strikeAsset common.Address, _collateralAsset common.Address, _strikePrice *big.Int, _expiry *big.Int, _isPut bool) (common.Address, error) {
	var out []interface{}
	err := _OtokenFactory.contract.Call(opts, &out, "getTargetOtokenAddress", _underlyingAsset, _strikeAsset, _collateralAsset, _strikePrice, _expiry, _isPut)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTargetOtokenAddress is a free data retrieval call binding the contract method 0xb86b9a64.
//
// Solidity: function getTargetOtokenAddress(address _underlyingAsset, address _strikeAsset, address _collateralAsset, uint256 _strikePrice, uint256 _expiry, bool _isPut) view returns(address)
func (_OtokenFactory *OtokenFactorySession) GetTargetOtokenAddress(_underlyingAsset common.Address, _strikeAsset common.Address, _collateralAsset common.Address, _strikePrice *big.Int, _expiry *big.Int, _isPut bool) (common.Address, error) {
	return _OtokenFactory.Contract.GetTargetOtokenAddress(&_OtokenFactory.CallOpts, _underlyingAsset, _strikeAsset, _collateralAsset, _strikePrice, _expiry, _isPut)
}

// GetTargetOtokenAddress is a free data retrieval call binding the contract method 0xb86b9a64.
//
// Solidity: function getTargetOtokenAddress(address _underlyingAsset, address _strikeAsset, address _collateralAsset, uint256 _strikePrice, uint256 _expiry, bool _isPut) view returns(address)
func (_OtokenFactory *OtokenFactoryCallerSession) GetTargetOtokenAddress(_underlyingAsset common.Address, _strikeAsset common.Address, _collateralAsset common.Address, _strikePrice *big.Int, _expiry *big.Int, _isPut bool) (common.Address, error) {
	return _OtokenFactory.Contract.GetTargetOtokenAddress(&_OtokenFactory.CallOpts, _underlyingAsset, _strikeAsset, _collateralAsset, _strikePrice, _expiry, _isPut)
}

// Otokens is a free data retrieval call binding the contract method 0xde0120de.
//
// Solidity: function otokens(uint256 ) view returns(address)
func (_OtokenFactory *OtokenFactoryCaller) Otokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OtokenFactory.contract.Call(opts, &out, "otokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Otokens is a free data retrieval call binding the contract method 0xde0120de.
//
// Solidity: function otokens(uint256 ) view returns(address)
func (_OtokenFactory *OtokenFactorySession) Otokens(arg0 *big.Int) (common.Address, error) {
	return _OtokenFactory.Contract.Otokens(&_OtokenFactory.CallOpts, arg0)
}

// Otokens is a free data retrieval call binding the contract method 0xde0120de.
//
// Solidity: function otokens(uint256 ) view returns(address)
func (_OtokenFactory *OtokenFactoryCallerSession) Otokens(arg0 *big.Int) (common.Address, error) {
	return _OtokenFactory.Contract.Otokens(&_OtokenFactory.CallOpts, arg0)
}

// CreateOtoken is a paid mutator transaction binding the contract method 0xc0974630.
//
// Solidity: function createOtoken(address _underlyingAsset, address _strikeAsset, address _collateralAsset, uint256 _strikePrice, uint256 _expiry, bool _isPut) returns(address)
func (_OtokenFactory *OtokenFactoryTransactor) CreateOtoken(opts *bind.TransactOpts, _underlyingAsset common.Address, _strikeAsset common.Address, _collateralAsset common.Address, _strikePrice *big.Int, _expiry *big.Int, _isPut bool) (*types.Transaction, error) {
	return _OtokenFactory.contract.Transact(opts, "createOtoken", _underlyingAsset, _strikeAsset, _collateralAsset, _strikePrice, _expiry, _isPut)
}

// CreateOtoken is a paid mutator transaction binding the contract method 0xc0974630.
//
// Solidity: function createOtoken(address _underlyingAsset, address _strikeAsset, address _collateralAsset, uint256 _strikePrice, uint256 _expiry, bool _isPut) returns(address)
func (_OtokenFactory *OtokenFactorySession) CreateOtoken(_underlyingAsset common.Address, _strikeAsset common.Address, _collateralAsset common.Address, _strikePrice *big.Int, _expiry *big.Int, _isPut bool) (*types.Transaction, error) {
	return _OtokenFactory.Contract.CreateOtoken(&_OtokenFactory.TransactOpts, _underlyingAsset, _strikeAsset, _collateralAsset, _strikePrice, _expiry, _isPut)
}

// CreateOtoken is a paid mutator transaction binding the contract method 0xc0974630.
//
// Solidity: function createOtoken(address _underlyingAsset, address _strikeAsset, address _collateralAsset, uint256 _strikePrice, uint256 _expiry, bool _isPut) returns(address)
func (_OtokenFactory *OtokenFactoryTransactorSession) CreateOtoken(_underlyingAsset common.Address, _strikeAsset common.Address, _collateralAsset common.Address, _strikePrice *big.Int, _expiry *big.Int, _isPut bool) (*types.Transaction, error) {
	return _OtokenFactory.Contract.CreateOtoken(&_OtokenFactory.TransactOpts, _underlyingAsset, _strikeAsset, _collateralAsset, _strikePrice, _expiry, _isPut)
}

// OtokenFactoryOtokenCreatedIterator is returned from FilterOtokenCreated and is used to iterate over the raw logs and unpacked data for OtokenCreated events raised by the OtokenFactory contract.
type OtokenFactoryOtokenCreatedIterator struct {
	Event *OtokenFactoryOtokenCreated // Event containing the contract specifics and raw log

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
func (it *OtokenFactoryOtokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenFactoryOtokenCreated)
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
		it.Event = new(OtokenFactoryOtokenCreated)
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
func (it *OtokenFactoryOtokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenFactoryOtokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenFactoryOtokenCreated represents a OtokenCreated event raised by the OtokenFactory contract.
type OtokenFactoryOtokenCreated struct {
	TokenAddress common.Address
	Creator      common.Address
	Underlying   common.Address
	Strike       common.Address
	Collateral   common.Address
	StrikePrice  *big.Int
	Expiry       *big.Int
	IsPut        bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOtokenCreated is a free log retrieval operation binding the contract event 0xedf283b0b3396dd34e23a917cb887ad557f18a593be3a2e9c069fd59f19a811a.
//
// Solidity: event OtokenCreated(address tokenAddress, address creator, address indexed underlying, address indexed strike, address indexed collateral, uint256 strikePrice, uint256 expiry, bool isPut)
func (_OtokenFactory *OtokenFactoryFilterer) FilterOtokenCreated(opts *bind.FilterOpts, underlying []common.Address, strike []common.Address, collateral []common.Address) (*OtokenFactoryOtokenCreatedIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var strikeRule []interface{}
	for _, strikeItem := range strike {
		strikeRule = append(strikeRule, strikeItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _OtokenFactory.contract.FilterLogs(opts, "OtokenCreated", underlyingRule, strikeRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return &OtokenFactoryOtokenCreatedIterator{contract: _OtokenFactory.contract, event: "OtokenCreated", logs: logs, sub: sub}, nil
}

// WatchOtokenCreated is a free log subscription operation binding the contract event 0xedf283b0b3396dd34e23a917cb887ad557f18a593be3a2e9c069fd59f19a811a.
//
// Solidity: event OtokenCreated(address tokenAddress, address creator, address indexed underlying, address indexed strike, address indexed collateral, uint256 strikePrice, uint256 expiry, bool isPut)
func (_OtokenFactory *OtokenFactoryFilterer) WatchOtokenCreated(opts *bind.WatchOpts, sink chan<- *OtokenFactoryOtokenCreated, underlying []common.Address, strike []common.Address, collateral []common.Address) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var strikeRule []interface{}
	for _, strikeItem := range strike {
		strikeRule = append(strikeRule, strikeItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _OtokenFactory.contract.WatchLogs(opts, "OtokenCreated", underlyingRule, strikeRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenFactoryOtokenCreated)
				if err := _OtokenFactory.contract.UnpackLog(event, "OtokenCreated", log); err != nil {
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

// ParseOtokenCreated is a log parse operation binding the contract event 0xedf283b0b3396dd34e23a917cb887ad557f18a593be3a2e9c069fd59f19a811a.
//
// Solidity: event OtokenCreated(address tokenAddress, address creator, address indexed underlying, address indexed strike, address indexed collateral, uint256 strikePrice, uint256 expiry, bool isPut)
func (_OtokenFactory *OtokenFactoryFilterer) ParseOtokenCreated(log types.Log) (*OtokenFactoryOtokenCreated, error) {
	event := new(OtokenFactoryOtokenCreated)
	if err := _OtokenFactory.contract.UnpackLog(event, "OtokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
