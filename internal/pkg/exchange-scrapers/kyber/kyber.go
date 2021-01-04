// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kyber

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

// KyberABI is the input ABI used to generate the binding from.
const KyberABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"}],\"name\":\"AdminClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAlerter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"AlerterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"EtherWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualSrcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualDestAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"}],\"name\":\"ExecuteTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIKyberHint\",\"name\":\"kyberHintHandler\",\"type\":\"address\"}],\"name\":\"KyberHintHandlerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIKyberNetwork\",\"name\":\"newKyberNetwork\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIKyberNetwork\",\"name\":\"previousKyberNetwork\",\"type\":\"address\"}],\"name\":\"KyberNetworkSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminPending\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAlerter\",\"type\":\"address\"}],\"name\":\"addAlerter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAlerters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"}],\"name\":\"getExpectedRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"worstRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"}],\"name\":\"getExpectedRateAfterFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kyberHintHandler\",\"outputs\":[{\"internalType\":\"contractIKyberHint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kyberNetwork\",\"outputs\":[{\"internalType\":\"contractIKyberNetwork\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"alerter\",\"type\":\"address\"}],\"name\":\"removeAlerter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIKyberHint\",\"name\":\"_kyberHintHandler\",\"type\":\"address\"}],\"name\":\"setHintHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIKyberNetwork\",\"name\":\"_kyberNetwork\",\"type\":\"address\"}],\"name\":\"setKyberNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapEtherToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapTokenToEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapTokenToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"destAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDestAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"destAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDestAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"walletId\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"}],\"name\":\"tradeWithHint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"destAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDestAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"}],\"name\":\"tradeWithHintAndFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminQuickly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Kyber is an auto generated Go binding around an Ethereum contract.
type Kyber struct {
	KyberCaller     // Read-only binding to the contract
	KyberTransactor // Write-only binding to the contract
	KyberFilterer   // Log filterer for contract events
}

// KyberCaller is an auto generated read-only Go binding around an Ethereum contract.
type KyberCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KyberTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KyberFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KyberSession struct {
	Contract     *Kyber            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KyberCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KyberCallerSession struct {
	Contract *KyberCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KyberTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KyberTransactorSession struct {
	Contract     *KyberTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KyberRaw is an auto generated low-level Go binding around an Ethereum contract.
type KyberRaw struct {
	Contract *Kyber // Generic contract binding to access the raw methods on
}

// KyberCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KyberCallerRaw struct {
	Contract *KyberCaller // Generic read-only contract binding to access the raw methods on
}

// KyberTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KyberTransactorRaw struct {
	Contract *KyberTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKyber creates a new instance of Kyber, bound to a specific deployed contract.
func NewKyber(address common.Address, backend bind.ContractBackend) (*Kyber, error) {
	contract, err := bindKyber(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kyber{KyberCaller: KyberCaller{contract: contract}, KyberTransactor: KyberTransactor{contract: contract}, KyberFilterer: KyberFilterer{contract: contract}}, nil
}

// NewKyberCaller creates a new read-only instance of Kyber, bound to a specific deployed contract.
func NewKyberCaller(address common.Address, caller bind.ContractCaller) (*KyberCaller, error) {
	contract, err := bindKyber(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KyberCaller{contract: contract}, nil
}

// NewKyberTransactor creates a new write-only instance of Kyber, bound to a specific deployed contract.
func NewKyberTransactor(address common.Address, transactor bind.ContractTransactor) (*KyberTransactor, error) {
	contract, err := bindKyber(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KyberTransactor{contract: contract}, nil
}

// NewKyberFilterer creates a new log filterer instance of Kyber, bound to a specific deployed contract.
func NewKyberFilterer(address common.Address, filterer bind.ContractFilterer) (*KyberFilterer, error) {
	contract, err := bindKyber(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KyberFilterer{contract: contract}, nil
}

// bindKyber binds a generic wrapper to an already deployed contract.
func bindKyber(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kyber *KyberRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kyber.Contract.KyberCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kyber *KyberRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kyber.Contract.KyberTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kyber *KyberRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kyber.Contract.KyberTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kyber *KyberCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kyber.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kyber *KyberTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kyber.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kyber *KyberTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kyber.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Kyber *KyberCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kyber.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Kyber *KyberSession) Admin() (common.Address, error) {
	return _Kyber.Contract.Admin(&_Kyber.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Kyber *KyberCallerSession) Admin() (common.Address, error) {
	return _Kyber.Contract.Admin(&_Kyber.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() view returns(bool)
func (_Kyber *KyberCaller) Enabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Kyber.contract.Call(opts, &out, "enabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() view returns(bool)
func (_Kyber *KyberSession) Enabled() (bool, error) {
	return _Kyber.Contract.Enabled(&_Kyber.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() view returns(bool)
func (_Kyber *KyberCallerSession) Enabled() (bool, error) {
	return _Kyber.Contract.Enabled(&_Kyber.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() view returns(address[])
func (_Kyber *KyberCaller) GetAlerters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Kyber.contract.Call(opts, &out, "getAlerters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() view returns(address[])
func (_Kyber *KyberSession) GetAlerters() ([]common.Address, error) {
	return _Kyber.Contract.GetAlerters(&_Kyber.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() view returns(address[])
func (_Kyber *KyberCallerSession) GetAlerters() ([]common.Address, error) {
	return _Kyber.Contract.GetAlerters(&_Kyber.CallOpts)
}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) view returns(uint256 expectedRate, uint256 worstRate)
func (_Kyber *KyberCaller) GetExpectedRate(opts *bind.CallOpts, src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	WorstRate    *big.Int
}, error) {
	var out []interface{}
	err := _Kyber.contract.Call(opts, &out, "getExpectedRate", src, dest, srcQty)

	outstruct := new(struct {
		ExpectedRate *big.Int
		WorstRate    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ExpectedRate = out[0].(*big.Int)
	outstruct.WorstRate = out[1].(*big.Int)

	return *outstruct, err

}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) view returns(uint256 expectedRate, uint256 worstRate)
func (_Kyber *KyberSession) GetExpectedRate(src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	WorstRate    *big.Int
}, error) {
	return _Kyber.Contract.GetExpectedRate(&_Kyber.CallOpts, src, dest, srcQty)
}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) view returns(uint256 expectedRate, uint256 worstRate)
func (_Kyber *KyberCallerSession) GetExpectedRate(src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	WorstRate    *big.Int
}, error) {
	return _Kyber.Contract.GetExpectedRate(&_Kyber.CallOpts, src, dest, srcQty)
}

// GetExpectedRateAfterFee is a free data retrieval call binding the contract method 0x418436bc.
//
// Solidity: function getExpectedRateAfterFee(address src, address dest, uint256 srcQty, uint256 platformFeeBps, bytes hint) view returns(uint256 expectedRate)
func (_Kyber *KyberCaller) GetExpectedRateAfterFee(opts *bind.CallOpts, src common.Address, dest common.Address, srcQty *big.Int, platformFeeBps *big.Int, hint []byte) (*big.Int, error) {
	var out []interface{}
	err := _Kyber.contract.Call(opts, &out, "getExpectedRateAfterFee", src, dest, srcQty, platformFeeBps, hint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpectedRateAfterFee is a free data retrieval call binding the contract method 0x418436bc.
//
// Solidity: function getExpectedRateAfterFee(address src, address dest, uint256 srcQty, uint256 platformFeeBps, bytes hint) view returns(uint256 expectedRate)
func (_Kyber *KyberSession) GetExpectedRateAfterFee(src common.Address, dest common.Address, srcQty *big.Int, platformFeeBps *big.Int, hint []byte) (*big.Int, error) {
	return _Kyber.Contract.GetExpectedRateAfterFee(&_Kyber.CallOpts, src, dest, srcQty, platformFeeBps, hint)
}

// GetExpectedRateAfterFee is a free data retrieval call binding the contract method 0x418436bc.
//
// Solidity: function getExpectedRateAfterFee(address src, address dest, uint256 srcQty, uint256 platformFeeBps, bytes hint) view returns(uint256 expectedRate)
func (_Kyber *KyberCallerSession) GetExpectedRateAfterFee(src common.Address, dest common.Address, srcQty *big.Int, platformFeeBps *big.Int, hint []byte) (*big.Int, error) {
	return _Kyber.Contract.GetExpectedRateAfterFee(&_Kyber.CallOpts, src, dest, srcQty, platformFeeBps, hint)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_Kyber *KyberCaller) GetOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Kyber.contract.Call(opts, &out, "getOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_Kyber *KyberSession) GetOperators() ([]common.Address, error) {
	return _Kyber.Contract.GetOperators(&_Kyber.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_Kyber *KyberCallerSession) GetOperators() ([]common.Address, error) {
	return _Kyber.Contract.GetOperators(&_Kyber.CallOpts)
}

// KyberHintHandler is a free data retrieval call binding the contract method 0x13c213b7.
//
// Solidity: function kyberHintHandler() view returns(address)
func (_Kyber *KyberCaller) KyberHintHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kyber.contract.Call(opts, &out, "kyberHintHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KyberHintHandler is a free data retrieval call binding the contract method 0x13c213b7.
//
// Solidity: function kyberHintHandler() view returns(address)
func (_Kyber *KyberSession) KyberHintHandler() (common.Address, error) {
	return _Kyber.Contract.KyberHintHandler(&_Kyber.CallOpts)
}

// KyberHintHandler is a free data retrieval call binding the contract method 0x13c213b7.
//
// Solidity: function kyberHintHandler() view returns(address)
func (_Kyber *KyberCallerSession) KyberHintHandler() (common.Address, error) {
	return _Kyber.Contract.KyberHintHandler(&_Kyber.CallOpts)
}

// KyberNetwork is a free data retrieval call binding the contract method 0xb78b842d.
//
// Solidity: function kyberNetwork() view returns(address)
func (_Kyber *KyberCaller) KyberNetwork(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kyber.contract.Call(opts, &out, "kyberNetwork")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KyberNetwork is a free data retrieval call binding the contract method 0xb78b842d.
//
// Solidity: function kyberNetwork() view returns(address)
func (_Kyber *KyberSession) KyberNetwork() (common.Address, error) {
	return _Kyber.Contract.KyberNetwork(&_Kyber.CallOpts)
}

// KyberNetwork is a free data retrieval call binding the contract method 0xb78b842d.
//
// Solidity: function kyberNetwork() view returns(address)
func (_Kyber *KyberCallerSession) KyberNetwork() (common.Address, error) {
	return _Kyber.Contract.KyberNetwork(&_Kyber.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_Kyber *KyberCaller) MaxGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kyber.contract.Call(opts, &out, "maxGasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_Kyber *KyberSession) MaxGasPrice() (*big.Int, error) {
	return _Kyber.Contract.MaxGasPrice(&_Kyber.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_Kyber *KyberCallerSession) MaxGasPrice() (*big.Int, error) {
	return _Kyber.Contract.MaxGasPrice(&_Kyber.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Kyber *KyberCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kyber.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Kyber *KyberSession) PendingAdmin() (common.Address, error) {
	return _Kyber.Contract.PendingAdmin(&_Kyber.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Kyber *KyberCallerSession) PendingAdmin() (common.Address, error) {
	return _Kyber.Contract.PendingAdmin(&_Kyber.CallOpts)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_Kyber *KyberTransactor) AddAlerter(opts *bind.TransactOpts, newAlerter common.Address) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "addAlerter", newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_Kyber *KyberSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.AddAlerter(&_Kyber.TransactOpts, newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_Kyber *KyberTransactorSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.AddAlerter(&_Kyber.TransactOpts, newAlerter)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_Kyber *KyberTransactor) AddOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "addOperator", newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_Kyber *KyberSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.AddOperator(&_Kyber.TransactOpts, newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_Kyber *KyberTransactorSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.AddOperator(&_Kyber.TransactOpts, newOperator)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_Kyber *KyberTransactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_Kyber *KyberSession) ClaimAdmin() (*types.Transaction, error) {
	return _Kyber.Contract.ClaimAdmin(&_Kyber.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_Kyber *KyberTransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _Kyber.Contract.ClaimAdmin(&_Kyber.TransactOpts)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_Kyber *KyberTransactor) RemoveAlerter(opts *bind.TransactOpts, alerter common.Address) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "removeAlerter", alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_Kyber *KyberSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.RemoveAlerter(&_Kyber.TransactOpts, alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_Kyber *KyberTransactorSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.RemoveAlerter(&_Kyber.TransactOpts, alerter)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_Kyber *KyberTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_Kyber *KyberSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.RemoveOperator(&_Kyber.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_Kyber *KyberTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.RemoveOperator(&_Kyber.TransactOpts, operator)
}

// SetHintHandler is a paid mutator transaction binding the contract method 0xb85d950f.
//
// Solidity: function setHintHandler(address _kyberHintHandler) returns()
func (_Kyber *KyberTransactor) SetHintHandler(opts *bind.TransactOpts, _kyberHintHandler common.Address) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "setHintHandler", _kyberHintHandler)
}

// SetHintHandler is a paid mutator transaction binding the contract method 0xb85d950f.
//
// Solidity: function setHintHandler(address _kyberHintHandler) returns()
func (_Kyber *KyberSession) SetHintHandler(_kyberHintHandler common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.SetHintHandler(&_Kyber.TransactOpts, _kyberHintHandler)
}

// SetHintHandler is a paid mutator transaction binding the contract method 0xb85d950f.
//
// Solidity: function setHintHandler(address _kyberHintHandler) returns()
func (_Kyber *KyberTransactorSession) SetHintHandler(_kyberHintHandler common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.SetHintHandler(&_Kyber.TransactOpts, _kyberHintHandler)
}

// SetKyberNetwork is a paid mutator transaction binding the contract method 0x54a325a6.
//
// Solidity: function setKyberNetwork(address _kyberNetwork) returns()
func (_Kyber *KyberTransactor) SetKyberNetwork(opts *bind.TransactOpts, _kyberNetwork common.Address) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "setKyberNetwork", _kyberNetwork)
}

// SetKyberNetwork is a paid mutator transaction binding the contract method 0x54a325a6.
//
// Solidity: function setKyberNetwork(address _kyberNetwork) returns()
func (_Kyber *KyberSession) SetKyberNetwork(_kyberNetwork common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.SetKyberNetwork(&_Kyber.TransactOpts, _kyberNetwork)
}

// SetKyberNetwork is a paid mutator transaction binding the contract method 0x54a325a6.
//
// Solidity: function setKyberNetwork(address _kyberNetwork) returns()
func (_Kyber *KyberTransactorSession) SetKyberNetwork(_kyberNetwork common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.SetKyberNetwork(&_Kyber.TransactOpts, _kyberNetwork)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) payable returns(uint256)
func (_Kyber *KyberTransactor) SwapEtherToToken(opts *bind.TransactOpts, token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "swapEtherToToken", token, minConversionRate)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) payable returns(uint256)
func (_Kyber *KyberSession) SwapEtherToToken(token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _Kyber.Contract.SwapEtherToToken(&_Kyber.TransactOpts, token, minConversionRate)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) payable returns(uint256)
func (_Kyber *KyberTransactorSession) SwapEtherToToken(token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _Kyber.Contract.SwapEtherToToken(&_Kyber.TransactOpts, token, minConversionRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 srcAmount, uint256 minConversionRate) returns(uint256)
func (_Kyber *KyberTransactor) SwapTokenToEther(opts *bind.TransactOpts, token common.Address, srcAmount *big.Int, minConversionRate *big.Int) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "swapTokenToEther", token, srcAmount, minConversionRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 srcAmount, uint256 minConversionRate) returns(uint256)
func (_Kyber *KyberSession) SwapTokenToEther(token common.Address, srcAmount *big.Int, minConversionRate *big.Int) (*types.Transaction, error) {
	return _Kyber.Contract.SwapTokenToEther(&_Kyber.TransactOpts, token, srcAmount, minConversionRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 srcAmount, uint256 minConversionRate) returns(uint256)
func (_Kyber *KyberTransactorSession) SwapTokenToEther(token common.Address, srcAmount *big.Int, minConversionRate *big.Int) (*types.Transaction, error) {
	return _Kyber.Contract.SwapTokenToEther(&_Kyber.TransactOpts, token, srcAmount, minConversionRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address src, uint256 srcAmount, address dest, uint256 minConversionRate) returns(uint256)
func (_Kyber *KyberTransactor) SwapTokenToToken(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "swapTokenToToken", src, srcAmount, dest, minConversionRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address src, uint256 srcAmount, address dest, uint256 minConversionRate) returns(uint256)
func (_Kyber *KyberSession) SwapTokenToToken(src common.Address, srcAmount *big.Int, dest common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _Kyber.Contract.SwapTokenToToken(&_Kyber.TransactOpts, src, srcAmount, dest, minConversionRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address src, uint256 srcAmount, address dest, uint256 minConversionRate) returns(uint256)
func (_Kyber *KyberTransactorSession) SwapTokenToToken(src common.Address, srcAmount *big.Int, dest common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _Kyber.Contract.SwapTokenToToken(&_Kyber.TransactOpts, src, srcAmount, dest, minConversionRate)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address platformWallet) payable returns(uint256)
func (_Kyber *KyberTransactor) Trade(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, platformWallet common.Address) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "trade", src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, platformWallet)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address platformWallet) payable returns(uint256)
func (_Kyber *KyberSession) Trade(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, platformWallet common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.Trade(&_Kyber.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, platformWallet)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address platformWallet) payable returns(uint256)
func (_Kyber *KyberTransactorSession) Trade(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, platformWallet common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.Trade(&_Kyber.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, platformWallet)
}

// TradeWithHint is a paid mutator transaction binding the contract method 0x29589f61.
//
// Solidity: function tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint) payable returns(uint256)
func (_Kyber *KyberTransactor) TradeWithHint(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address, hint []byte) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "tradeWithHint", src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId, hint)
}

// TradeWithHint is a paid mutator transaction binding the contract method 0x29589f61.
//
// Solidity: function tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint) payable returns(uint256)
func (_Kyber *KyberSession) TradeWithHint(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address, hint []byte) (*types.Transaction, error) {
	return _Kyber.Contract.TradeWithHint(&_Kyber.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId, hint)
}

// TradeWithHint is a paid mutator transaction binding the contract method 0x29589f61.
//
// Solidity: function tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint) payable returns(uint256)
func (_Kyber *KyberTransactorSession) TradeWithHint(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address, hint []byte) (*types.Transaction, error) {
	return _Kyber.Contract.TradeWithHint(&_Kyber.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId, hint)
}

// TradeWithHintAndFee is a paid mutator transaction binding the contract method 0xae591d54.
//
// Solidity: function tradeWithHintAndFee(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address platformWallet, uint256 platformFeeBps, bytes hint) payable returns(uint256 destAmount)
func (_Kyber *KyberTransactor) TradeWithHintAndFee(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, platformWallet common.Address, platformFeeBps *big.Int, hint []byte) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "tradeWithHintAndFee", src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, platformWallet, platformFeeBps, hint)
}

// TradeWithHintAndFee is a paid mutator transaction binding the contract method 0xae591d54.
//
// Solidity: function tradeWithHintAndFee(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address platformWallet, uint256 platformFeeBps, bytes hint) payable returns(uint256 destAmount)
func (_Kyber *KyberSession) TradeWithHintAndFee(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, platformWallet common.Address, platformFeeBps *big.Int, hint []byte) (*types.Transaction, error) {
	return _Kyber.Contract.TradeWithHintAndFee(&_Kyber.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, platformWallet, platformFeeBps, hint)
}

// TradeWithHintAndFee is a paid mutator transaction binding the contract method 0xae591d54.
//
// Solidity: function tradeWithHintAndFee(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address platformWallet, uint256 platformFeeBps, bytes hint) payable returns(uint256 destAmount)
func (_Kyber *KyberTransactorSession) TradeWithHintAndFee(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, platformWallet common.Address, platformFeeBps *big.Int, hint []byte) (*types.Transaction, error) {
	return _Kyber.Contract.TradeWithHintAndFee(&_Kyber.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, platformWallet, platformFeeBps, hint)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_Kyber *KyberTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_Kyber *KyberSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.TransferAdmin(&_Kyber.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_Kyber *KyberTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.TransferAdmin(&_Kyber.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_Kyber *KyberTransactor) TransferAdminQuickly(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "transferAdminQuickly", newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_Kyber *KyberSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.TransferAdminQuickly(&_Kyber.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_Kyber *KyberTransactorSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.TransferAdminQuickly(&_Kyber.TransactOpts, newAdmin)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_Kyber *KyberTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "withdrawEther", amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_Kyber *KyberSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.WithdrawEther(&_Kyber.TransactOpts, amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_Kyber *KyberTransactorSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.WithdrawEther(&_Kyber.TransactOpts, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_Kyber *KyberTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _Kyber.contract.Transact(opts, "withdrawToken", token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_Kyber *KyberSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.WithdrawToken(&_Kyber.TransactOpts, token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_Kyber *KyberTransactorSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _Kyber.Contract.WithdrawToken(&_Kyber.TransactOpts, token, amount, sendTo)
}

// KyberAdminClaimedIterator is returned from FilterAdminClaimed and is used to iterate over the raw logs and unpacked data for AdminClaimed events raised by the Kyber contract.
type KyberAdminClaimedIterator struct {
	Event *KyberAdminClaimed // Event containing the contract specifics and raw log

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
func (it *KyberAdminClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberAdminClaimed)
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
		it.Event = new(KyberAdminClaimed)
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
func (it *KyberAdminClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberAdminClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberAdminClaimed represents a AdminClaimed event raised by the Kyber contract.
type KyberAdminClaimed struct {
	NewAdmin      common.Address
	PreviousAdmin common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminClaimed is a free log retrieval operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_Kyber *KyberFilterer) FilterAdminClaimed(opts *bind.FilterOpts) (*KyberAdminClaimedIterator, error) {

	logs, sub, err := _Kyber.contract.FilterLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return &KyberAdminClaimedIterator{contract: _Kyber.contract, event: "AdminClaimed", logs: logs, sub: sub}, nil
}

// WatchAdminClaimed is a free log subscription operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_Kyber *KyberFilterer) WatchAdminClaimed(opts *bind.WatchOpts, sink chan<- *KyberAdminClaimed) (event.Subscription, error) {

	logs, sub, err := _Kyber.contract.WatchLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberAdminClaimed)
				if err := _Kyber.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
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

// ParseAdminClaimed is a log parse operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_Kyber *KyberFilterer) ParseAdminClaimed(log types.Log) (*KyberAdminClaimed, error) {
	event := new(KyberAdminClaimed)
	if err := _Kyber.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberAlerterAddedIterator is returned from FilterAlerterAdded and is used to iterate over the raw logs and unpacked data for AlerterAdded events raised by the Kyber contract.
type KyberAlerterAddedIterator struct {
	Event *KyberAlerterAdded // Event containing the contract specifics and raw log

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
func (it *KyberAlerterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberAlerterAdded)
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
		it.Event = new(KyberAlerterAdded)
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
func (it *KyberAlerterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberAlerterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberAlerterAdded represents a AlerterAdded event raised by the Kyber contract.
type KyberAlerterAdded struct {
	NewAlerter common.Address
	IsAdd      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAlerterAdded is a free log retrieval operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_Kyber *KyberFilterer) FilterAlerterAdded(opts *bind.FilterOpts) (*KyberAlerterAddedIterator, error) {

	logs, sub, err := _Kyber.contract.FilterLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return &KyberAlerterAddedIterator{contract: _Kyber.contract, event: "AlerterAdded", logs: logs, sub: sub}, nil
}

// WatchAlerterAdded is a free log subscription operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_Kyber *KyberFilterer) WatchAlerterAdded(opts *bind.WatchOpts, sink chan<- *KyberAlerterAdded) (event.Subscription, error) {

	logs, sub, err := _Kyber.contract.WatchLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberAlerterAdded)
				if err := _Kyber.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
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

// ParseAlerterAdded is a log parse operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_Kyber *KyberFilterer) ParseAlerterAdded(log types.Log) (*KyberAlerterAdded, error) {
	event := new(KyberAlerterAdded)
	if err := _Kyber.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberEtherWithdrawIterator is returned from FilterEtherWithdraw and is used to iterate over the raw logs and unpacked data for EtherWithdraw events raised by the Kyber contract.
type KyberEtherWithdrawIterator struct {
	Event *KyberEtherWithdraw // Event containing the contract specifics and raw log

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
func (it *KyberEtherWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberEtherWithdraw)
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
		it.Event = new(KyberEtherWithdraw)
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
func (it *KyberEtherWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberEtherWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberEtherWithdraw represents a EtherWithdraw event raised by the Kyber contract.
type KyberEtherWithdraw struct {
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdraw is a free log retrieval operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_Kyber *KyberFilterer) FilterEtherWithdraw(opts *bind.FilterOpts) (*KyberEtherWithdrawIterator, error) {

	logs, sub, err := _Kyber.contract.FilterLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return &KyberEtherWithdrawIterator{contract: _Kyber.contract, event: "EtherWithdraw", logs: logs, sub: sub}, nil
}

// WatchEtherWithdraw is a free log subscription operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_Kyber *KyberFilterer) WatchEtherWithdraw(opts *bind.WatchOpts, sink chan<- *KyberEtherWithdraw) (event.Subscription, error) {

	logs, sub, err := _Kyber.contract.WatchLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberEtherWithdraw)
				if err := _Kyber.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
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

// ParseEtherWithdraw is a log parse operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_Kyber *KyberFilterer) ParseEtherWithdraw(log types.Log) (*KyberEtherWithdraw, error) {
	event := new(KyberEtherWithdraw)
	if err := _Kyber.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberExecuteTradeIterator is returned from FilterExecuteTrade and is used to iterate over the raw logs and unpacked data for ExecuteTrade events raised by the Kyber contract.
type KyberExecuteTradeIterator struct {
	Event *KyberExecuteTrade // Event containing the contract specifics and raw log

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
func (it *KyberExecuteTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberExecuteTrade)
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
		it.Event = new(KyberExecuteTrade)
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
func (it *KyberExecuteTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberExecuteTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberExecuteTrade represents a ExecuteTrade event raised by the Kyber contract.
type KyberExecuteTrade struct {
	Trader           common.Address
	Src              common.Address
	Dest             common.Address
	DestAddress      common.Address
	ActualSrcAmount  *big.Int
	ActualDestAmount *big.Int
	PlatformWallet   common.Address
	PlatformFeeBps   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExecuteTrade is a free log retrieval operation binding the contract event 0xf724b4df6617473612b53d7f88ecc6ea983074b30960a049fcd0657ffe808083.
//
// Solidity: event ExecuteTrade(address indexed trader, address src, address dest, address destAddress, uint256 actualSrcAmount, uint256 actualDestAmount, address platformWallet, uint256 platformFeeBps)
func (_Kyber *KyberFilterer) FilterExecuteTrade(opts *bind.FilterOpts, trader []common.Address) (*KyberExecuteTradeIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Kyber.contract.FilterLogs(opts, "ExecuteTrade", traderRule)
	if err != nil {
		return nil, err
	}
	return &KyberExecuteTradeIterator{contract: _Kyber.contract, event: "ExecuteTrade", logs: logs, sub: sub}, nil
}

// WatchExecuteTrade is a free log subscription operation binding the contract event 0xf724b4df6617473612b53d7f88ecc6ea983074b30960a049fcd0657ffe808083.
//
// Solidity: event ExecuteTrade(address indexed trader, address src, address dest, address destAddress, uint256 actualSrcAmount, uint256 actualDestAmount, address platformWallet, uint256 platformFeeBps)
func (_Kyber *KyberFilterer) WatchExecuteTrade(opts *bind.WatchOpts, sink chan<- *KyberExecuteTrade, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Kyber.contract.WatchLogs(opts, "ExecuteTrade", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberExecuteTrade)
				if err := _Kyber.contract.UnpackLog(event, "ExecuteTrade", log); err != nil {
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

// ParseExecuteTrade is a log parse operation binding the contract event 0xf724b4df6617473612b53d7f88ecc6ea983074b30960a049fcd0657ffe808083.
//
// Solidity: event ExecuteTrade(address indexed trader, address src, address dest, address destAddress, uint256 actualSrcAmount, uint256 actualDestAmount, address platformWallet, uint256 platformFeeBps)
func (_Kyber *KyberFilterer) ParseExecuteTrade(log types.Log) (*KyberExecuteTrade, error) {
	event := new(KyberExecuteTrade)
	if err := _Kyber.contract.UnpackLog(event, "ExecuteTrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberKyberHintHandlerSetIterator is returned from FilterKyberHintHandlerSet and is used to iterate over the raw logs and unpacked data for KyberHintHandlerSet events raised by the Kyber contract.
type KyberKyberHintHandlerSetIterator struct {
	Event *KyberKyberHintHandlerSet // Event containing the contract specifics and raw log

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
func (it *KyberKyberHintHandlerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberKyberHintHandlerSet)
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
		it.Event = new(KyberKyberHintHandlerSet)
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
func (it *KyberKyberHintHandlerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberKyberHintHandlerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberKyberHintHandlerSet represents a KyberHintHandlerSet event raised by the Kyber contract.
type KyberKyberHintHandlerSet struct {
	KyberHintHandler common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterKyberHintHandlerSet is a free log retrieval operation binding the contract event 0x6deb3a98fd141d661e9c0fb2d847541cc0c629cfb100c61011a76f57cb3b3a9b.
//
// Solidity: event KyberHintHandlerSet(address kyberHintHandler)
func (_Kyber *KyberFilterer) FilterKyberHintHandlerSet(opts *bind.FilterOpts) (*KyberKyberHintHandlerSetIterator, error) {

	logs, sub, err := _Kyber.contract.FilterLogs(opts, "KyberHintHandlerSet")
	if err != nil {
		return nil, err
	}
	return &KyberKyberHintHandlerSetIterator{contract: _Kyber.contract, event: "KyberHintHandlerSet", logs: logs, sub: sub}, nil
}

// WatchKyberHintHandlerSet is a free log subscription operation binding the contract event 0x6deb3a98fd141d661e9c0fb2d847541cc0c629cfb100c61011a76f57cb3b3a9b.
//
// Solidity: event KyberHintHandlerSet(address kyberHintHandler)
func (_Kyber *KyberFilterer) WatchKyberHintHandlerSet(opts *bind.WatchOpts, sink chan<- *KyberKyberHintHandlerSet) (event.Subscription, error) {

	logs, sub, err := _Kyber.contract.WatchLogs(opts, "KyberHintHandlerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberKyberHintHandlerSet)
				if err := _Kyber.contract.UnpackLog(event, "KyberHintHandlerSet", log); err != nil {
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

// ParseKyberHintHandlerSet is a log parse operation binding the contract event 0x6deb3a98fd141d661e9c0fb2d847541cc0c629cfb100c61011a76f57cb3b3a9b.
//
// Solidity: event KyberHintHandlerSet(address kyberHintHandler)
func (_Kyber *KyberFilterer) ParseKyberHintHandlerSet(log types.Log) (*KyberKyberHintHandlerSet, error) {
	event := new(KyberKyberHintHandlerSet)
	if err := _Kyber.contract.UnpackLog(event, "KyberHintHandlerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberKyberNetworkSetIterator is returned from FilterKyberNetworkSet and is used to iterate over the raw logs and unpacked data for KyberNetworkSet events raised by the Kyber contract.
type KyberKyberNetworkSetIterator struct {
	Event *KyberKyberNetworkSet // Event containing the contract specifics and raw log

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
func (it *KyberKyberNetworkSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberKyberNetworkSet)
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
		it.Event = new(KyberKyberNetworkSet)
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
func (it *KyberKyberNetworkSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberKyberNetworkSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberKyberNetworkSet represents a KyberNetworkSet event raised by the Kyber contract.
type KyberKyberNetworkSet struct {
	NewKyberNetwork      common.Address
	PreviousKyberNetwork common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterKyberNetworkSet is a free log retrieval operation binding the contract event 0x8936e1f096bf0a8c9df862b3d1d5b82774cad78116200175f00b5b7ba3010b02.
//
// Solidity: event KyberNetworkSet(address newKyberNetwork, address previousKyberNetwork)
func (_Kyber *KyberFilterer) FilterKyberNetworkSet(opts *bind.FilterOpts) (*KyberKyberNetworkSetIterator, error) {

	logs, sub, err := _Kyber.contract.FilterLogs(opts, "KyberNetworkSet")
	if err != nil {
		return nil, err
	}
	return &KyberKyberNetworkSetIterator{contract: _Kyber.contract, event: "KyberNetworkSet", logs: logs, sub: sub}, nil
}

// WatchKyberNetworkSet is a free log subscription operation binding the contract event 0x8936e1f096bf0a8c9df862b3d1d5b82774cad78116200175f00b5b7ba3010b02.
//
// Solidity: event KyberNetworkSet(address newKyberNetwork, address previousKyberNetwork)
func (_Kyber *KyberFilterer) WatchKyberNetworkSet(opts *bind.WatchOpts, sink chan<- *KyberKyberNetworkSet) (event.Subscription, error) {

	logs, sub, err := _Kyber.contract.WatchLogs(opts, "KyberNetworkSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberKyberNetworkSet)
				if err := _Kyber.contract.UnpackLog(event, "KyberNetworkSet", log); err != nil {
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

// ParseKyberNetworkSet is a log parse operation binding the contract event 0x8936e1f096bf0a8c9df862b3d1d5b82774cad78116200175f00b5b7ba3010b02.
//
// Solidity: event KyberNetworkSet(address newKyberNetwork, address previousKyberNetwork)
func (_Kyber *KyberFilterer) ParseKyberNetworkSet(log types.Log) (*KyberKyberNetworkSet, error) {
	event := new(KyberKyberNetworkSet)
	if err := _Kyber.contract.UnpackLog(event, "KyberNetworkSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the Kyber contract.
type KyberOperatorAddedIterator struct {
	Event *KyberOperatorAdded // Event containing the contract specifics and raw log

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
func (it *KyberOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberOperatorAdded)
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
		it.Event = new(KyberOperatorAdded)
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
func (it *KyberOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberOperatorAdded represents a OperatorAdded event raised by the Kyber contract.
type KyberOperatorAdded struct {
	NewOperator common.Address
	IsAdd       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_Kyber *KyberFilterer) FilterOperatorAdded(opts *bind.FilterOpts) (*KyberOperatorAddedIterator, error) {

	logs, sub, err := _Kyber.contract.FilterLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return &KyberOperatorAddedIterator{contract: _Kyber.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_Kyber *KyberFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *KyberOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _Kyber.contract.WatchLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberOperatorAdded)
				if err := _Kyber.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_Kyber *KyberFilterer) ParseOperatorAdded(log types.Log) (*KyberOperatorAdded, error) {
	event := new(KyberOperatorAdded)
	if err := _Kyber.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the Kyber contract.
type KyberTokenWithdrawIterator struct {
	Event *KyberTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *KyberTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberTokenWithdraw)
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
		it.Event = new(KyberTokenWithdraw)
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
func (it *KyberTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberTokenWithdraw represents a TokenWithdraw event raised by the Kyber contract.
type KyberTokenWithdraw struct {
	Token  common.Address
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_Kyber *KyberFilterer) FilterTokenWithdraw(opts *bind.FilterOpts) (*KyberTokenWithdrawIterator, error) {

	logs, sub, err := _Kyber.contract.FilterLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return &KyberTokenWithdrawIterator{contract: _Kyber.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_Kyber *KyberFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *KyberTokenWithdraw) (event.Subscription, error) {

	logs, sub, err := _Kyber.contract.WatchLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberTokenWithdraw)
				if err := _Kyber.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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

// ParseTokenWithdraw is a log parse operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_Kyber *KyberFilterer) ParseTokenWithdraw(log types.Log) (*KyberTokenWithdraw, error) {
	event := new(KyberTokenWithdraw)
	if err := _Kyber.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberTransferAdminPendingIterator is returned from FilterTransferAdminPending and is used to iterate over the raw logs and unpacked data for TransferAdminPending events raised by the Kyber contract.
type KyberTransferAdminPendingIterator struct {
	Event *KyberTransferAdminPending // Event containing the contract specifics and raw log

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
func (it *KyberTransferAdminPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberTransferAdminPending)
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
		it.Event = new(KyberTransferAdminPending)
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
func (it *KyberTransferAdminPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberTransferAdminPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberTransferAdminPending represents a TransferAdminPending event raised by the Kyber contract.
type KyberTransferAdminPending struct {
	PendingAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminPending is a free log retrieval operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_Kyber *KyberFilterer) FilterTransferAdminPending(opts *bind.FilterOpts) (*KyberTransferAdminPendingIterator, error) {

	logs, sub, err := _Kyber.contract.FilterLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return &KyberTransferAdminPendingIterator{contract: _Kyber.contract, event: "TransferAdminPending", logs: logs, sub: sub}, nil
}

// WatchTransferAdminPending is a free log subscription operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_Kyber *KyberFilterer) WatchTransferAdminPending(opts *bind.WatchOpts, sink chan<- *KyberTransferAdminPending) (event.Subscription, error) {

	logs, sub, err := _Kyber.contract.WatchLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberTransferAdminPending)
				if err := _Kyber.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
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

// ParseTransferAdminPending is a log parse operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_Kyber *KyberFilterer) ParseTransferAdminPending(log types.Log) (*KyberTransferAdminPending, error) {
	event := new(KyberTransferAdminPending)
	if err := _Kyber.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
