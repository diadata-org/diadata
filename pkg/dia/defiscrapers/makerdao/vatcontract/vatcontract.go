// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vatcontract

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

// VatcontractABI is the input ABI used to generate the binding from.
const VatcontractABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg3\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"Line\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"can\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"dai\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"debt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"flux\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"rate\",\"type\":\"int256\"}],\"name\":\"fold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"fork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"w\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"frob\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"w\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"grab\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"heal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"hope\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"Art\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"line\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dust\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"nope\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"wad\",\"type\":\"int256\"}],\"name\":\"slip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"suck\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"urns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ink\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"art\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Vatcontract is an auto generated Go binding around an Ethereum contract.
type Vatcontract struct {
	VatcontractCaller     // Read-only binding to the contract
	VatcontractTransactor // Write-only binding to the contract
	VatcontractFilterer   // Log filterer for contract events
}

// VatcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type VatcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VatcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VatcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VatcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VatcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VatcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VatcontractSession struct {
	Contract     *Vatcontract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VatcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VatcontractCallerSession struct {
	Contract *VatcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VatcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VatcontractTransactorSession struct {
	Contract     *VatcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VatcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type VatcontractRaw struct {
	Contract *Vatcontract // Generic contract binding to access the raw methods on
}

// VatcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VatcontractCallerRaw struct {
	Contract *VatcontractCaller // Generic read-only contract binding to access the raw methods on
}

// VatcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VatcontractTransactorRaw struct {
	Contract *VatcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVatcontract creates a new instance of Vatcontract, bound to a specific deployed contract.
func NewVatcontract(address common.Address, backend bind.ContractBackend) (*Vatcontract, error) {
	contract, err := bindVatcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vatcontract{VatcontractCaller: VatcontractCaller{contract: contract}, VatcontractTransactor: VatcontractTransactor{contract: contract}, VatcontractFilterer: VatcontractFilterer{contract: contract}}, nil
}

// NewVatcontractCaller creates a new read-only instance of Vatcontract, bound to a specific deployed contract.
func NewVatcontractCaller(address common.Address, caller bind.ContractCaller) (*VatcontractCaller, error) {
	contract, err := bindVatcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VatcontractCaller{contract: contract}, nil
}

// NewVatcontractTransactor creates a new write-only instance of Vatcontract, bound to a specific deployed contract.
func NewVatcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*VatcontractTransactor, error) {
	contract, err := bindVatcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VatcontractTransactor{contract: contract}, nil
}

// NewVatcontractFilterer creates a new log filterer instance of Vatcontract, bound to a specific deployed contract.
func NewVatcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*VatcontractFilterer, error) {
	contract, err := bindVatcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VatcontractFilterer{contract: contract}, nil
}

// bindVatcontract binds a generic wrapper to an already deployed contract.
func bindVatcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VatcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vatcontract *VatcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vatcontract.Contract.VatcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vatcontract *VatcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vatcontract.Contract.VatcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vatcontract *VatcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vatcontract.Contract.VatcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vatcontract *VatcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vatcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vatcontract *VatcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vatcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vatcontract *VatcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vatcontract.Contract.contract.Transact(opts, method, params...)
}

// Line is a free data retrieval call binding the contract method 0xbabe8a3f.
//
// Solidity: function Line() view returns(uint256)
func (_Vatcontract *VatcontractCaller) Line(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vatcontract.contract.Call(opts, &out, "Line")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Line is a free data retrieval call binding the contract method 0xbabe8a3f.
//
// Solidity: function Line() view returns(uint256)
func (_Vatcontract *VatcontractSession) Line() (*big.Int, error) {
	return _Vatcontract.Contract.Line(&_Vatcontract.CallOpts)
}

// Line is a free data retrieval call binding the contract method 0xbabe8a3f.
//
// Solidity: function Line() view returns(uint256)
func (_Vatcontract *VatcontractCallerSession) Line() (*big.Int, error) {
	return _Vatcontract.Contract.Line(&_Vatcontract.CallOpts)
}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) view returns(uint256)
func (_Vatcontract *VatcontractCaller) Can(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vatcontract.contract.Call(opts, &out, "can", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) view returns(uint256)
func (_Vatcontract *VatcontractSession) Can(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vatcontract.Contract.Can(&_Vatcontract.CallOpts, arg0, arg1)
}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) view returns(uint256)
func (_Vatcontract *VatcontractCallerSession) Can(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vatcontract.Contract.Can(&_Vatcontract.CallOpts, arg0, arg1)
}

// Dai is a free data retrieval call binding the contract method 0x6c25b346.
//
// Solidity: function dai(address ) view returns(uint256)
func (_Vatcontract *VatcontractCaller) Dai(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vatcontract.contract.Call(opts, &out, "dai", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dai is a free data retrieval call binding the contract method 0x6c25b346.
//
// Solidity: function dai(address ) view returns(uint256)
func (_Vatcontract *VatcontractSession) Dai(arg0 common.Address) (*big.Int, error) {
	return _Vatcontract.Contract.Dai(&_Vatcontract.CallOpts, arg0)
}

// Dai is a free data retrieval call binding the contract method 0x6c25b346.
//
// Solidity: function dai(address ) view returns(uint256)
func (_Vatcontract *VatcontractCallerSession) Dai(arg0 common.Address) (*big.Int, error) {
	return _Vatcontract.Contract.Dai(&_Vatcontract.CallOpts, arg0)
}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_Vatcontract *VatcontractCaller) Debt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vatcontract.contract.Call(opts, &out, "debt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_Vatcontract *VatcontractSession) Debt() (*big.Int, error) {
	return _Vatcontract.Contract.Debt(&_Vatcontract.CallOpts)
}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_Vatcontract *VatcontractCallerSession) Debt() (*big.Int, error) {
	return _Vatcontract.Contract.Debt(&_Vatcontract.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x214414d5.
//
// Solidity: function gem(bytes32 , address ) view returns(uint256)
func (_Vatcontract *VatcontractCaller) Gem(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vatcontract.contract.Call(opts, &out, "gem", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x214414d5.
//
// Solidity: function gem(bytes32 , address ) view returns(uint256)
func (_Vatcontract *VatcontractSession) Gem(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _Vatcontract.Contract.Gem(&_Vatcontract.CallOpts, arg0, arg1)
}

// Gem is a free data retrieval call binding the contract method 0x214414d5.
//
// Solidity: function gem(bytes32 , address ) view returns(uint256)
func (_Vatcontract *VatcontractCallerSession) Gem(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _Vatcontract.Contract.Gem(&_Vatcontract.CallOpts, arg0, arg1)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 Art, uint256 rate, uint256 spot, uint256 line, uint256 dust)
func (_Vatcontract *VatcontractCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Art  *big.Int
	Rate *big.Int
	Spot *big.Int
	Line *big.Int
	Dust *big.Int
}, error) {
	var out []interface{}
	err := _Vatcontract.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Art  *big.Int
		Rate *big.Int
		Spot *big.Int
		Line *big.Int
		Dust *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Art = out[0].(*big.Int)
	outstruct.Rate = out[1].(*big.Int)
	outstruct.Spot = out[2].(*big.Int)
	outstruct.Line = out[3].(*big.Int)
	outstruct.Dust = out[4].(*big.Int)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 Art, uint256 rate, uint256 spot, uint256 line, uint256 dust)
func (_Vatcontract *VatcontractSession) Ilks(arg0 [32]byte) (struct {
	Art  *big.Int
	Rate *big.Int
	Spot *big.Int
	Line *big.Int
	Dust *big.Int
}, error) {
	return _Vatcontract.Contract.Ilks(&_Vatcontract.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 Art, uint256 rate, uint256 spot, uint256 line, uint256 dust)
func (_Vatcontract *VatcontractCallerSession) Ilks(arg0 [32]byte) (struct {
	Art  *big.Int
	Rate *big.Int
	Spot *big.Int
	Line *big.Int
	Dust *big.Int
}, error) {
	return _Vatcontract.Contract.Ilks(&_Vatcontract.CallOpts, arg0)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Vatcontract *VatcontractCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vatcontract.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Vatcontract *VatcontractSession) Live() (*big.Int, error) {
	return _Vatcontract.Contract.Live(&_Vatcontract.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Vatcontract *VatcontractCallerSession) Live() (*big.Int, error) {
	return _Vatcontract.Contract.Live(&_Vatcontract.CallOpts)
}

// Sin is a free data retrieval call binding the contract method 0xf059212a.
//
// Solidity: function sin(address ) view returns(uint256)
func (_Vatcontract *VatcontractCaller) Sin(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vatcontract.contract.Call(opts, &out, "sin", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sin is a free data retrieval call binding the contract method 0xf059212a.
//
// Solidity: function sin(address ) view returns(uint256)
func (_Vatcontract *VatcontractSession) Sin(arg0 common.Address) (*big.Int, error) {
	return _Vatcontract.Contract.Sin(&_Vatcontract.CallOpts, arg0)
}

// Sin is a free data retrieval call binding the contract method 0xf059212a.
//
// Solidity: function sin(address ) view returns(uint256)
func (_Vatcontract *VatcontractCallerSession) Sin(arg0 common.Address) (*big.Int, error) {
	return _Vatcontract.Contract.Sin(&_Vatcontract.CallOpts, arg0)
}

// Urns is a free data retrieval call binding the contract method 0x2424be5c.
//
// Solidity: function urns(bytes32 , address ) view returns(uint256 ink, uint256 art)
func (_Vatcontract *VatcontractCaller) Urns(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (struct {
	Ink *big.Int
	Art *big.Int
}, error) {
	var out []interface{}
	err := _Vatcontract.contract.Call(opts, &out, "urns", arg0, arg1)

	outstruct := new(struct {
		Ink *big.Int
		Art *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ink = out[0].(*big.Int)
	outstruct.Art = out[1].(*big.Int)

	return *outstruct, err

}

// Urns is a free data retrieval call binding the contract method 0x2424be5c.
//
// Solidity: function urns(bytes32 , address ) view returns(uint256 ink, uint256 art)
func (_Vatcontract *VatcontractSession) Urns(arg0 [32]byte, arg1 common.Address) (struct {
	Ink *big.Int
	Art *big.Int
}, error) {
	return _Vatcontract.Contract.Urns(&_Vatcontract.CallOpts, arg0, arg1)
}

// Urns is a free data retrieval call binding the contract method 0x2424be5c.
//
// Solidity: function urns(bytes32 , address ) view returns(uint256 ink, uint256 art)
func (_Vatcontract *VatcontractCallerSession) Urns(arg0 [32]byte, arg1 common.Address) (struct {
	Ink *big.Int
	Art *big.Int
}, error) {
	return _Vatcontract.Contract.Urns(&_Vatcontract.CallOpts, arg0, arg1)
}

// Vice is a free data retrieval call binding the contract method 0x2d61a355.
//
// Solidity: function vice() view returns(uint256)
func (_Vatcontract *VatcontractCaller) Vice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vatcontract.contract.Call(opts, &out, "vice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Vice is a free data retrieval call binding the contract method 0x2d61a355.
//
// Solidity: function vice() view returns(uint256)
func (_Vatcontract *VatcontractSession) Vice() (*big.Int, error) {
	return _Vatcontract.Contract.Vice(&_Vatcontract.CallOpts)
}

// Vice is a free data retrieval call binding the contract method 0x2d61a355.
//
// Solidity: function vice() view returns(uint256)
func (_Vatcontract *VatcontractCallerSession) Vice() (*big.Int, error) {
	return _Vatcontract.Contract.Vice(&_Vatcontract.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vatcontract *VatcontractCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vatcontract.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vatcontract *VatcontractSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Vatcontract.Contract.Wards(&_Vatcontract.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vatcontract *VatcontractCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Vatcontract.Contract.Wards(&_Vatcontract.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vatcontract *VatcontractTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vatcontract *VatcontractSession) Cage() (*types.Transaction, error) {
	return _Vatcontract.Contract.Cage(&_Vatcontract.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vatcontract *VatcontractTransactorSession) Cage() (*types.Transaction, error) {
	return _Vatcontract.Contract.Cage(&_Vatcontract.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vatcontract *VatcontractTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vatcontract *VatcontractSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Vatcontract.Contract.Deny(&_Vatcontract.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vatcontract *VatcontractTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Vatcontract.Contract.Deny(&_Vatcontract.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Vatcontract *VatcontractTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Vatcontract *VatcontractSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.File(&_Vatcontract.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Vatcontract *VatcontractTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.File(&_Vatcontract.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vatcontract *VatcontractTransactor) File0(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vatcontract *VatcontractSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.File0(&_Vatcontract.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vatcontract *VatcontractTransactorSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.File0(&_Vatcontract.TransactOpts, what, data)
}

// Flux is a paid mutator transaction binding the contract method 0x6111be2e.
//
// Solidity: function flux(bytes32 ilk, address src, address dst, uint256 wad) returns()
func (_Vatcontract *VatcontractTransactor) Flux(opts *bind.TransactOpts, ilk [32]byte, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "flux", ilk, src, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x6111be2e.
//
// Solidity: function flux(bytes32 ilk, address src, address dst, uint256 wad) returns()
func (_Vatcontract *VatcontractSession) Flux(ilk [32]byte, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Flux(&_Vatcontract.TransactOpts, ilk, src, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x6111be2e.
//
// Solidity: function flux(bytes32 ilk, address src, address dst, uint256 wad) returns()
func (_Vatcontract *VatcontractTransactorSession) Flux(ilk [32]byte, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Flux(&_Vatcontract.TransactOpts, ilk, src, dst, wad)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 i, address u, int256 rate) returns()
func (_Vatcontract *VatcontractTransactor) Fold(opts *bind.TransactOpts, i [32]byte, u common.Address, rate *big.Int) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "fold", i, u, rate)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 i, address u, int256 rate) returns()
func (_Vatcontract *VatcontractSession) Fold(i [32]byte, u common.Address, rate *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Fold(&_Vatcontract.TransactOpts, i, u, rate)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 i, address u, int256 rate) returns()
func (_Vatcontract *VatcontractTransactorSession) Fold(i [32]byte, u common.Address, rate *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Fold(&_Vatcontract.TransactOpts, i, u, rate)
}

// Fork is a paid mutator transaction binding the contract method 0x870c616d.
//
// Solidity: function fork(bytes32 ilk, address src, address dst, int256 dink, int256 dart) returns()
func (_Vatcontract *VatcontractTransactor) Fork(opts *bind.TransactOpts, ilk [32]byte, src common.Address, dst common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "fork", ilk, src, dst, dink, dart)
}

// Fork is a paid mutator transaction binding the contract method 0x870c616d.
//
// Solidity: function fork(bytes32 ilk, address src, address dst, int256 dink, int256 dart) returns()
func (_Vatcontract *VatcontractSession) Fork(ilk [32]byte, src common.Address, dst common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Fork(&_Vatcontract.TransactOpts, ilk, src, dst, dink, dart)
}

// Fork is a paid mutator transaction binding the contract method 0x870c616d.
//
// Solidity: function fork(bytes32 ilk, address src, address dst, int256 dink, int256 dart) returns()
func (_Vatcontract *VatcontractTransactorSession) Fork(ilk [32]byte, src common.Address, dst common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Fork(&_Vatcontract.TransactOpts, ilk, src, dst, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vatcontract *VatcontractTransactor) Frob(opts *bind.TransactOpts, i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "frob", i, u, v, w, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vatcontract *VatcontractSession) Frob(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Frob(&_Vatcontract.TransactOpts, i, u, v, w, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vatcontract *VatcontractTransactorSession) Frob(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Frob(&_Vatcontract.TransactOpts, i, u, v, w, dink, dart)
}

// Grab is a paid mutator transaction binding the contract method 0x7bab3f40.
//
// Solidity: function grab(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vatcontract *VatcontractTransactor) Grab(opts *bind.TransactOpts, i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "grab", i, u, v, w, dink, dart)
}

// Grab is a paid mutator transaction binding the contract method 0x7bab3f40.
//
// Solidity: function grab(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vatcontract *VatcontractSession) Grab(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Grab(&_Vatcontract.TransactOpts, i, u, v, w, dink, dart)
}

// Grab is a paid mutator transaction binding the contract method 0x7bab3f40.
//
// Solidity: function grab(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vatcontract *VatcontractTransactorSession) Grab(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Grab(&_Vatcontract.TransactOpts, i, u, v, w, dink, dart)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vatcontract *VatcontractTransactor) Heal(opts *bind.TransactOpts, rad *big.Int) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "heal", rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vatcontract *VatcontractSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Heal(&_Vatcontract.TransactOpts, rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vatcontract *VatcontractTransactorSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Heal(&_Vatcontract.TransactOpts, rad)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_Vatcontract *VatcontractTransactor) Hope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "hope", usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_Vatcontract *VatcontractSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _Vatcontract.Contract.Hope(&_Vatcontract.TransactOpts, usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_Vatcontract *VatcontractTransactorSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _Vatcontract.Contract.Hope(&_Vatcontract.TransactOpts, usr)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Vatcontract *VatcontractTransactor) Init(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "init", ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Vatcontract *VatcontractSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _Vatcontract.Contract.Init(&_Vatcontract.TransactOpts, ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Vatcontract *VatcontractTransactorSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _Vatcontract.Contract.Init(&_Vatcontract.TransactOpts, ilk)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 rad) returns()
func (_Vatcontract *VatcontractTransactor) Move(opts *bind.TransactOpts, src common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "move", src, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 rad) returns()
func (_Vatcontract *VatcontractSession) Move(src common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Move(&_Vatcontract.TransactOpts, src, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 rad) returns()
func (_Vatcontract *VatcontractTransactorSession) Move(src common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Move(&_Vatcontract.TransactOpts, src, dst, rad)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_Vatcontract *VatcontractTransactor) Nope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "nope", usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_Vatcontract *VatcontractSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _Vatcontract.Contract.Nope(&_Vatcontract.TransactOpts, usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_Vatcontract *VatcontractTransactorSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _Vatcontract.Contract.Nope(&_Vatcontract.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vatcontract *VatcontractTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vatcontract *VatcontractSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Vatcontract.Contract.Rely(&_Vatcontract.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vatcontract *VatcontractTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Vatcontract.Contract.Rely(&_Vatcontract.TransactOpts, usr)
}

// Slip is a paid mutator transaction binding the contract method 0x7cdd3fde.
//
// Solidity: function slip(bytes32 ilk, address usr, int256 wad) returns()
func (_Vatcontract *VatcontractTransactor) Slip(opts *bind.TransactOpts, ilk [32]byte, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "slip", ilk, usr, wad)
}

// Slip is a paid mutator transaction binding the contract method 0x7cdd3fde.
//
// Solidity: function slip(bytes32 ilk, address usr, int256 wad) returns()
func (_Vatcontract *VatcontractSession) Slip(ilk [32]byte, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Slip(&_Vatcontract.TransactOpts, ilk, usr, wad)
}

// Slip is a paid mutator transaction binding the contract method 0x7cdd3fde.
//
// Solidity: function slip(bytes32 ilk, address usr, int256 wad) returns()
func (_Vatcontract *VatcontractTransactorSession) Slip(ilk [32]byte, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Slip(&_Vatcontract.TransactOpts, ilk, usr, wad)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address u, address v, uint256 rad) returns()
func (_Vatcontract *VatcontractTransactor) Suck(opts *bind.TransactOpts, u common.Address, v common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vatcontract.contract.Transact(opts, "suck", u, v, rad)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address u, address v, uint256 rad) returns()
func (_Vatcontract *VatcontractSession) Suck(u common.Address, v common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Suck(&_Vatcontract.TransactOpts, u, v, rad)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address u, address v, uint256 rad) returns()
func (_Vatcontract *VatcontractTransactorSession) Suck(u common.Address, v common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vatcontract.Contract.Suck(&_Vatcontract.TransactOpts, u, v, rad)
}
