// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OtokenController

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

// ActionsActionArgs is an auto generated low-level Go binding around an user-defined struct.
type ActionsActionArgs struct {
	ActionType    uint8
	Owner         common.Address
	SecondAddress common.Address
	Asset         common.Address
	VaultId       *big.Int
	Amount        *big.Int
	Index         *big.Int
	Data          []byte
}

// MarginVaultVault is an auto generated low-level Go binding around an user-defined struct.
type MarginVaultVault struct {
	ShortOtokens      []common.Address
	LongOtokens       []common.Address
	CollateralAssets  []common.Address
	ShortAmounts      []*big.Int
	LongAmounts       []*big.Int
	CollateralAmounts []*big.Int
}

// OtokenControllerABI is the input ABI used to generate the binding from.
const OtokenControllerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AccountOperatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"CallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRestricted\",\"type\":\"bool\"}],\"name\":\"CallRestricted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CollateralAssetDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"AccountOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CollateralAssetWithdrawed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldFullPauser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFullPauser\",\"type\":\"address\"}],\"name\":\"FullPauserUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"otoken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LongOtokenDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"otoken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"AccountOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LongOtokenWithdrawed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPartialPauser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPartialPauser\",\"type\":\"address\"}],\"name\":\"PartialPauserUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"otoken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"otokenBurned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"otoken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"AccountOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ShortOtokenBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"otoken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"AccountOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ShortOtokenMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"SystemFullyPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"SystemPartiallyPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"VaultOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"AccountOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"otoken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"}],\"name\":\"VaultSettled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressbook\",\"outputs\":[{\"internalType\":\"contractAddressBookInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculator\",\"outputs\":[{\"internalType\":\"contractMarginCalculatorInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callRestricted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fullPauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accountOwner\",\"type\":\"address\"}],\"name\":\"getAccountVaultCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfiguration\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_otoken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getPayout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"getProceed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"getVault\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"shortOtokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"longOtokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"collateralAssets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shortAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"longAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"collateralAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structMarginVault.Vault\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_otoken\",\"type\":\"address\"}],\"name\":\"hasExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressBook\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_otoken\",\"type\":\"address\"}],\"name\":\"isSettlementAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumActions.ActionType\",\"name\":\"actionType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"secondAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structActions.ActionArgs[]\",\"name\":\"_actions\",\"type\":\"tuple[]\"}],\"name\":\"operate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractOracleInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partialPauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractMarginPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isRestricted\",\"type\":\"bool\"}],\"name\":\"setCallRestriction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fullPauser\",\"type\":\"address\"}],\"name\":\"setFullPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isOperator\",\"type\":\"bool\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_partialPauser\",\"type\":\"address\"}],\"name\":\"setPartialPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_fullyPaused\",\"type\":\"bool\"}],\"name\":\"setSystemFullyPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_partiallyPaused\",\"type\":\"bool\"}],\"name\":\"setSystemPartiallyPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemFullyPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemPartiallyPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"contractWhitelistInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OtokenController is an auto generated Go binding around an Ethereum contract.
type OtokenController struct {
	OtokenControllerCaller     // Read-only binding to the contract
	OtokenControllerTransactor // Write-only binding to the contract
	OtokenControllerFilterer   // Log filterer for contract events
}

// OtokenControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OtokenControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtokenControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OtokenControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtokenControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OtokenControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtokenControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OtokenControllerSession struct {
	Contract     *OtokenController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OtokenControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OtokenControllerCallerSession struct {
	Contract *OtokenControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// OtokenControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OtokenControllerTransactorSession struct {
	Contract     *OtokenControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OtokenControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OtokenControllerRaw struct {
	Contract *OtokenController // Generic contract binding to access the raw methods on
}

// OtokenControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OtokenControllerCallerRaw struct {
	Contract *OtokenControllerCaller // Generic read-only contract binding to access the raw methods on
}

// OtokenControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OtokenControllerTransactorRaw struct {
	Contract *OtokenControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOtokenController creates a new instance of OtokenController, bound to a specific deployed contract.
func NewOtokenController(address common.Address, backend bind.ContractBackend) (*OtokenController, error) {
	contract, err := bindOtokenController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OtokenController{OtokenControllerCaller: OtokenControllerCaller{contract: contract}, OtokenControllerTransactor: OtokenControllerTransactor{contract: contract}, OtokenControllerFilterer: OtokenControllerFilterer{contract: contract}}, nil
}

// NewOtokenControllerCaller creates a new read-only instance of OtokenController, bound to a specific deployed contract.
func NewOtokenControllerCaller(address common.Address, caller bind.ContractCaller) (*OtokenControllerCaller, error) {
	contract, err := bindOtokenController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerCaller{contract: contract}, nil
}

// NewOtokenControllerTransactor creates a new write-only instance of OtokenController, bound to a specific deployed contract.
func NewOtokenControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*OtokenControllerTransactor, error) {
	contract, err := bindOtokenController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerTransactor{contract: contract}, nil
}

// NewOtokenControllerFilterer creates a new log filterer instance of OtokenController, bound to a specific deployed contract.
func NewOtokenControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*OtokenControllerFilterer, error) {
	contract, err := bindOtokenController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerFilterer{contract: contract}, nil
}

// bindOtokenController binds a generic wrapper to an already deployed contract.
func bindOtokenController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OtokenControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OtokenController *OtokenControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OtokenController.Contract.OtokenControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OtokenController *OtokenControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtokenController.Contract.OtokenControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OtokenController *OtokenControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OtokenController.Contract.OtokenControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OtokenController *OtokenControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OtokenController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OtokenController *OtokenControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtokenController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OtokenController *OtokenControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OtokenController.Contract.contract.Transact(opts, method, params...)
}

// Addressbook is a free data retrieval call binding the contract method 0x70dc320c.
//
// Solidity: function addressbook() view returns(address)
func (_OtokenController *OtokenControllerCaller) Addressbook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "addressbook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addressbook is a free data retrieval call binding the contract method 0x70dc320c.
//
// Solidity: function addressbook() view returns(address)
func (_OtokenController *OtokenControllerSession) Addressbook() (common.Address, error) {
	return _OtokenController.Contract.Addressbook(&_OtokenController.CallOpts)
}

// Addressbook is a free data retrieval call binding the contract method 0x70dc320c.
//
// Solidity: function addressbook() view returns(address)
func (_OtokenController *OtokenControllerCallerSession) Addressbook() (common.Address, error) {
	return _OtokenController.Contract.Addressbook(&_OtokenController.CallOpts)
}

// Calculator is a free data retrieval call binding the contract method 0xce3e39c0.
//
// Solidity: function calculator() view returns(address)
func (_OtokenController *OtokenControllerCaller) Calculator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "calculator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Calculator is a free data retrieval call binding the contract method 0xce3e39c0.
//
// Solidity: function calculator() view returns(address)
func (_OtokenController *OtokenControllerSession) Calculator() (common.Address, error) {
	return _OtokenController.Contract.Calculator(&_OtokenController.CallOpts)
}

// Calculator is a free data retrieval call binding the contract method 0xce3e39c0.
//
// Solidity: function calculator() view returns(address)
func (_OtokenController *OtokenControllerCallerSession) Calculator() (common.Address, error) {
	return _OtokenController.Contract.Calculator(&_OtokenController.CallOpts)
}

// CallRestricted is a free data retrieval call binding the contract method 0x6c0c3b99.
//
// Solidity: function callRestricted() view returns(bool)
func (_OtokenController *OtokenControllerCaller) CallRestricted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "callRestricted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CallRestricted is a free data retrieval call binding the contract method 0x6c0c3b99.
//
// Solidity: function callRestricted() view returns(bool)
func (_OtokenController *OtokenControllerSession) CallRestricted() (bool, error) {
	return _OtokenController.Contract.CallRestricted(&_OtokenController.CallOpts)
}

// CallRestricted is a free data retrieval call binding the contract method 0x6c0c3b99.
//
// Solidity: function callRestricted() view returns(bool)
func (_OtokenController *OtokenControllerCallerSession) CallRestricted() (bool, error) {
	return _OtokenController.Contract.CallRestricted(&_OtokenController.CallOpts)
}

// FullPauser is a free data retrieval call binding the contract method 0x29729d88.
//
// Solidity: function fullPauser() view returns(address)
func (_OtokenController *OtokenControllerCaller) FullPauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "fullPauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FullPauser is a free data retrieval call binding the contract method 0x29729d88.
//
// Solidity: function fullPauser() view returns(address)
func (_OtokenController *OtokenControllerSession) FullPauser() (common.Address, error) {
	return _OtokenController.Contract.FullPauser(&_OtokenController.CallOpts)
}

// FullPauser is a free data retrieval call binding the contract method 0x29729d88.
//
// Solidity: function fullPauser() view returns(address)
func (_OtokenController *OtokenControllerCallerSession) FullPauser() (common.Address, error) {
	return _OtokenController.Contract.FullPauser(&_OtokenController.CallOpts)
}

// GetAccountVaultCounter is a free data retrieval call binding the contract method 0xcaa6d21a.
//
// Solidity: function getAccountVaultCounter(address _accountOwner) view returns(uint256)
func (_OtokenController *OtokenControllerCaller) GetAccountVaultCounter(opts *bind.CallOpts, _accountOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "getAccountVaultCounter", _accountOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountVaultCounter is a free data retrieval call binding the contract method 0xcaa6d21a.
//
// Solidity: function getAccountVaultCounter(address _accountOwner) view returns(uint256)
func (_OtokenController *OtokenControllerSession) GetAccountVaultCounter(_accountOwner common.Address) (*big.Int, error) {
	return _OtokenController.Contract.GetAccountVaultCounter(&_OtokenController.CallOpts, _accountOwner)
}

// GetAccountVaultCounter is a free data retrieval call binding the contract method 0xcaa6d21a.
//
// Solidity: function getAccountVaultCounter(address _accountOwner) view returns(uint256)
func (_OtokenController *OtokenControllerCallerSession) GetAccountVaultCounter(_accountOwner common.Address) (*big.Int, error) {
	return _OtokenController.Contract.GetAccountVaultCounter(&_OtokenController.CallOpts, _accountOwner)
}

// GetConfiguration is a free data retrieval call binding the contract method 0x6bd50cef.
//
// Solidity: function getConfiguration() view returns(address, address, address, address)
func (_OtokenController *OtokenControllerCaller) GetConfiguration(opts *bind.CallOpts) (common.Address, common.Address, common.Address, common.Address, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "getConfiguration")

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return out0, out1, out2, out3, err

}

// GetConfiguration is a free data retrieval call binding the contract method 0x6bd50cef.
//
// Solidity: function getConfiguration() view returns(address, address, address, address)
func (_OtokenController *OtokenControllerSession) GetConfiguration() (common.Address, common.Address, common.Address, common.Address, error) {
	return _OtokenController.Contract.GetConfiguration(&_OtokenController.CallOpts)
}

// GetConfiguration is a free data retrieval call binding the contract method 0x6bd50cef.
//
// Solidity: function getConfiguration() view returns(address, address, address, address)
func (_OtokenController *OtokenControllerCallerSession) GetConfiguration() (common.Address, common.Address, common.Address, common.Address, error) {
	return _OtokenController.Contract.GetConfiguration(&_OtokenController.CallOpts)
}

// GetPayout is a free data retrieval call binding the contract method 0x565eea19.
//
// Solidity: function getPayout(address _otoken, uint256 _amount) view returns(uint256)
func (_OtokenController *OtokenControllerCaller) GetPayout(opts *bind.CallOpts, _otoken common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "getPayout", _otoken, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPayout is a free data retrieval call binding the contract method 0x565eea19.
//
// Solidity: function getPayout(address _otoken, uint256 _amount) view returns(uint256)
func (_OtokenController *OtokenControllerSession) GetPayout(_otoken common.Address, _amount *big.Int) (*big.Int, error) {
	return _OtokenController.Contract.GetPayout(&_OtokenController.CallOpts, _otoken, _amount)
}

// GetPayout is a free data retrieval call binding the contract method 0x565eea19.
//
// Solidity: function getPayout(address _otoken, uint256 _amount) view returns(uint256)
func (_OtokenController *OtokenControllerCallerSession) GetPayout(_otoken common.Address, _amount *big.Int) (*big.Int, error) {
	return _OtokenController.Contract.GetPayout(&_OtokenController.CallOpts, _otoken, _amount)
}

// GetProceed is a free data retrieval call binding the contract method 0xc220101d.
//
// Solidity: function getProceed(address _owner, uint256 _vaultId) view returns(uint256)
func (_OtokenController *OtokenControllerCaller) GetProceed(opts *bind.CallOpts, _owner common.Address, _vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "getProceed", _owner, _vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProceed is a free data retrieval call binding the contract method 0xc220101d.
//
// Solidity: function getProceed(address _owner, uint256 _vaultId) view returns(uint256)
func (_OtokenController *OtokenControllerSession) GetProceed(_owner common.Address, _vaultId *big.Int) (*big.Int, error) {
	return _OtokenController.Contract.GetProceed(&_OtokenController.CallOpts, _owner, _vaultId)
}

// GetProceed is a free data retrieval call binding the contract method 0xc220101d.
//
// Solidity: function getProceed(address _owner, uint256 _vaultId) view returns(uint256)
func (_OtokenController *OtokenControllerCallerSession) GetProceed(_owner common.Address, _vaultId *big.Int) (*big.Int, error) {
	return _OtokenController.Contract.GetProceed(&_OtokenController.CallOpts, _owner, _vaultId)
}

// GetVault is a free data retrieval call binding the contract method 0xd99d13f5.
//
// Solidity: function getVault(address _owner, uint256 _vaultId) view returns((address[],address[],address[],uint256[],uint256[],uint256[]))
func (_OtokenController *OtokenControllerCaller) GetVault(opts *bind.CallOpts, _owner common.Address, _vaultId *big.Int) (MarginVaultVault, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "getVault", _owner, _vaultId)

	if err != nil {
		return *new(MarginVaultVault), err
	}

	out0 := *abi.ConvertType(out[0], new(MarginVaultVault)).(*MarginVaultVault)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0xd99d13f5.
//
// Solidity: function getVault(address _owner, uint256 _vaultId) view returns((address[],address[],address[],uint256[],uint256[],uint256[]))
func (_OtokenController *OtokenControllerSession) GetVault(_owner common.Address, _vaultId *big.Int) (MarginVaultVault, error) {
	return _OtokenController.Contract.GetVault(&_OtokenController.CallOpts, _owner, _vaultId)
}

// GetVault is a free data retrieval call binding the contract method 0xd99d13f5.
//
// Solidity: function getVault(address _owner, uint256 _vaultId) view returns((address[],address[],address[],uint256[],uint256[],uint256[]))
func (_OtokenController *OtokenControllerCallerSession) GetVault(_owner common.Address, _vaultId *big.Int) (MarginVaultVault, error) {
	return _OtokenController.Contract.GetVault(&_OtokenController.CallOpts, _owner, _vaultId)
}

// HasExpired is a free data retrieval call binding the contract method 0xf77bc88b.
//
// Solidity: function hasExpired(address _otoken) view returns(bool)
func (_OtokenController *OtokenControllerCaller) HasExpired(opts *bind.CallOpts, _otoken common.Address) (bool, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "hasExpired", _otoken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasExpired is a free data retrieval call binding the contract method 0xf77bc88b.
//
// Solidity: function hasExpired(address _otoken) view returns(bool)
func (_OtokenController *OtokenControllerSession) HasExpired(_otoken common.Address) (bool, error) {
	return _OtokenController.Contract.HasExpired(&_OtokenController.CallOpts, _otoken)
}

// HasExpired is a free data retrieval call binding the contract method 0xf77bc88b.
//
// Solidity: function hasExpired(address _otoken) view returns(bool)
func (_OtokenController *OtokenControllerCallerSession) HasExpired(_otoken common.Address) (bool, error) {
	return _OtokenController.Contract.HasExpired(&_OtokenController.CallOpts, _otoken)
}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address _owner, address _operator) view returns(bool)
func (_OtokenController *OtokenControllerCaller) IsOperator(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "isOperator", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address _owner, address _operator) view returns(bool)
func (_OtokenController *OtokenControllerSession) IsOperator(_owner common.Address, _operator common.Address) (bool, error) {
	return _OtokenController.Contract.IsOperator(&_OtokenController.CallOpts, _owner, _operator)
}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address _owner, address _operator) view returns(bool)
func (_OtokenController *OtokenControllerCallerSession) IsOperator(_owner common.Address, _operator common.Address) (bool, error) {
	return _OtokenController.Contract.IsOperator(&_OtokenController.CallOpts, _owner, _operator)
}

// IsSettlementAllowed is a free data retrieval call binding the contract method 0xe723406c.
//
// Solidity: function isSettlementAllowed(address _otoken) view returns(bool)
func (_OtokenController *OtokenControllerCaller) IsSettlementAllowed(opts *bind.CallOpts, _otoken common.Address) (bool, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "isSettlementAllowed", _otoken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSettlementAllowed is a free data retrieval call binding the contract method 0xe723406c.
//
// Solidity: function isSettlementAllowed(address _otoken) view returns(bool)
func (_OtokenController *OtokenControllerSession) IsSettlementAllowed(_otoken common.Address) (bool, error) {
	return _OtokenController.Contract.IsSettlementAllowed(&_OtokenController.CallOpts, _otoken)
}

// IsSettlementAllowed is a free data retrieval call binding the contract method 0xe723406c.
//
// Solidity: function isSettlementAllowed(address _otoken) view returns(bool)
func (_OtokenController *OtokenControllerCallerSession) IsSettlementAllowed(_otoken common.Address) (bool, error) {
	return _OtokenController.Contract.IsSettlementAllowed(&_OtokenController.CallOpts, _otoken)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_OtokenController *OtokenControllerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_OtokenController *OtokenControllerSession) Oracle() (common.Address, error) {
	return _OtokenController.Contract.Oracle(&_OtokenController.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_OtokenController *OtokenControllerCallerSession) Oracle() (common.Address, error) {
	return _OtokenController.Contract.Oracle(&_OtokenController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OtokenController *OtokenControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OtokenController *OtokenControllerSession) Owner() (common.Address, error) {
	return _OtokenController.Contract.Owner(&_OtokenController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OtokenController *OtokenControllerCallerSession) Owner() (common.Address, error) {
	return _OtokenController.Contract.Owner(&_OtokenController.CallOpts)
}

// PartialPauser is a free data retrieval call binding the contract method 0x9db93891.
//
// Solidity: function partialPauser() view returns(address)
func (_OtokenController *OtokenControllerCaller) PartialPauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "partialPauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PartialPauser is a free data retrieval call binding the contract method 0x9db93891.
//
// Solidity: function partialPauser() view returns(address)
func (_OtokenController *OtokenControllerSession) PartialPauser() (common.Address, error) {
	return _OtokenController.Contract.PartialPauser(&_OtokenController.CallOpts)
}

// PartialPauser is a free data retrieval call binding the contract method 0x9db93891.
//
// Solidity: function partialPauser() view returns(address)
func (_OtokenController *OtokenControllerCallerSession) PartialPauser() (common.Address, error) {
	return _OtokenController.Contract.PartialPauser(&_OtokenController.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_OtokenController *OtokenControllerCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_OtokenController *OtokenControllerSession) Pool() (common.Address, error) {
	return _OtokenController.Contract.Pool(&_OtokenController.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_OtokenController *OtokenControllerCallerSession) Pool() (common.Address, error) {
	return _OtokenController.Contract.Pool(&_OtokenController.CallOpts)
}

// SystemFullyPaused is a free data retrieval call binding the contract method 0xeab7775b.
//
// Solidity: function systemFullyPaused() view returns(bool)
func (_OtokenController *OtokenControllerCaller) SystemFullyPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "systemFullyPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SystemFullyPaused is a free data retrieval call binding the contract method 0xeab7775b.
//
// Solidity: function systemFullyPaused() view returns(bool)
func (_OtokenController *OtokenControllerSession) SystemFullyPaused() (bool, error) {
	return _OtokenController.Contract.SystemFullyPaused(&_OtokenController.CallOpts)
}

// SystemFullyPaused is a free data retrieval call binding the contract method 0xeab7775b.
//
// Solidity: function systemFullyPaused() view returns(bool)
func (_OtokenController *OtokenControllerCallerSession) SystemFullyPaused() (bool, error) {
	return _OtokenController.Contract.SystemFullyPaused(&_OtokenController.CallOpts)
}

// SystemPartiallyPaused is a free data retrieval call binding the contract method 0xcdee058a.
//
// Solidity: function systemPartiallyPaused() view returns(bool)
func (_OtokenController *OtokenControllerCaller) SystemPartiallyPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "systemPartiallyPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SystemPartiallyPaused is a free data retrieval call binding the contract method 0xcdee058a.
//
// Solidity: function systemPartiallyPaused() view returns(bool)
func (_OtokenController *OtokenControllerSession) SystemPartiallyPaused() (bool, error) {
	return _OtokenController.Contract.SystemPartiallyPaused(&_OtokenController.CallOpts)
}

// SystemPartiallyPaused is a free data retrieval call binding the contract method 0xcdee058a.
//
// Solidity: function systemPartiallyPaused() view returns(bool)
func (_OtokenController *OtokenControllerCallerSession) SystemPartiallyPaused() (bool, error) {
	return _OtokenController.Contract.SystemPartiallyPaused(&_OtokenController.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_OtokenController *OtokenControllerCaller) Whitelist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OtokenController.contract.Call(opts, &out, "whitelist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_OtokenController *OtokenControllerSession) Whitelist() (common.Address, error) {
	return _OtokenController.Contract.Whitelist(&_OtokenController.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_OtokenController *OtokenControllerCallerSession) Whitelist() (common.Address, error) {
	return _OtokenController.Contract.Whitelist(&_OtokenController.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _addressBook, address _owner) returns()
func (_OtokenController *OtokenControllerTransactor) Initialize(opts *bind.TransactOpts, _addressBook common.Address, _owner common.Address) (*types.Transaction, error) {
	return _OtokenController.contract.Transact(opts, "initialize", _addressBook, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _addressBook, address _owner) returns()
func (_OtokenController *OtokenControllerSession) Initialize(_addressBook common.Address, _owner common.Address) (*types.Transaction, error) {
	return _OtokenController.Contract.Initialize(&_OtokenController.TransactOpts, _addressBook, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _addressBook, address _owner) returns()
func (_OtokenController *OtokenControllerTransactorSession) Initialize(_addressBook common.Address, _owner common.Address) (*types.Transaction, error) {
	return _OtokenController.Contract.Initialize(&_OtokenController.TransactOpts, _addressBook, _owner)
}

// Operate is a paid mutator transaction binding the contract method 0xb617f0c6.
//
// Solidity: function operate((uint8,address,address,address,uint256,uint256,uint256,bytes)[] _actions) returns()
func (_OtokenController *OtokenControllerTransactor) Operate(opts *bind.TransactOpts, _actions []ActionsActionArgs) (*types.Transaction, error) {
	return _OtokenController.contract.Transact(opts, "operate", _actions)
}

// Operate is a paid mutator transaction binding the contract method 0xb617f0c6.
//
// Solidity: function operate((uint8,address,address,address,uint256,uint256,uint256,bytes)[] _actions) returns()
func (_OtokenController *OtokenControllerSession) Operate(_actions []ActionsActionArgs) (*types.Transaction, error) {
	return _OtokenController.Contract.Operate(&_OtokenController.TransactOpts, _actions)
}

// Operate is a paid mutator transaction binding the contract method 0xb617f0c6.
//
// Solidity: function operate((uint8,address,address,address,uint256,uint256,uint256,bytes)[] _actions) returns()
func (_OtokenController *OtokenControllerTransactorSession) Operate(_actions []ActionsActionArgs) (*types.Transaction, error) {
	return _OtokenController.Contract.Operate(&_OtokenController.TransactOpts, _actions)
}

// RefreshConfiguration is a paid mutator transaction binding the contract method 0x64681083.
//
// Solidity: function refreshConfiguration() returns()
func (_OtokenController *OtokenControllerTransactor) RefreshConfiguration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtokenController.contract.Transact(opts, "refreshConfiguration")
}

// RefreshConfiguration is a paid mutator transaction binding the contract method 0x64681083.
//
// Solidity: function refreshConfiguration() returns()
func (_OtokenController *OtokenControllerSession) RefreshConfiguration() (*types.Transaction, error) {
	return _OtokenController.Contract.RefreshConfiguration(&_OtokenController.TransactOpts)
}

// RefreshConfiguration is a paid mutator transaction binding the contract method 0x64681083.
//
// Solidity: function refreshConfiguration() returns()
func (_OtokenController *OtokenControllerTransactorSession) RefreshConfiguration() (*types.Transaction, error) {
	return _OtokenController.Contract.RefreshConfiguration(&_OtokenController.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OtokenController *OtokenControllerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtokenController.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OtokenController *OtokenControllerSession) RenounceOwnership() (*types.Transaction, error) {
	return _OtokenController.Contract.RenounceOwnership(&_OtokenController.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OtokenController *OtokenControllerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OtokenController.Contract.RenounceOwnership(&_OtokenController.TransactOpts)
}

// SetCallRestriction is a paid mutator transaction binding the contract method 0xcab2e805.
//
// Solidity: function setCallRestriction(bool _isRestricted) returns()
func (_OtokenController *OtokenControllerTransactor) SetCallRestriction(opts *bind.TransactOpts, _isRestricted bool) (*types.Transaction, error) {
	return _OtokenController.contract.Transact(opts, "setCallRestriction", _isRestricted)
}

// SetCallRestriction is a paid mutator transaction binding the contract method 0xcab2e805.
//
// Solidity: function setCallRestriction(bool _isRestricted) returns()
func (_OtokenController *OtokenControllerSession) SetCallRestriction(_isRestricted bool) (*types.Transaction, error) {
	return _OtokenController.Contract.SetCallRestriction(&_OtokenController.TransactOpts, _isRestricted)
}

// SetCallRestriction is a paid mutator transaction binding the contract method 0xcab2e805.
//
// Solidity: function setCallRestriction(bool _isRestricted) returns()
func (_OtokenController *OtokenControllerTransactorSession) SetCallRestriction(_isRestricted bool) (*types.Transaction, error) {
	return _OtokenController.Contract.SetCallRestriction(&_OtokenController.TransactOpts, _isRestricted)
}

// SetFullPauser is a paid mutator transaction binding the contract method 0xbeca75d7.
//
// Solidity: function setFullPauser(address _fullPauser) returns()
func (_OtokenController *OtokenControllerTransactor) SetFullPauser(opts *bind.TransactOpts, _fullPauser common.Address) (*types.Transaction, error) {
	return _OtokenController.contract.Transact(opts, "setFullPauser", _fullPauser)
}

// SetFullPauser is a paid mutator transaction binding the contract method 0xbeca75d7.
//
// Solidity: function setFullPauser(address _fullPauser) returns()
func (_OtokenController *OtokenControllerSession) SetFullPauser(_fullPauser common.Address) (*types.Transaction, error) {
	return _OtokenController.Contract.SetFullPauser(&_OtokenController.TransactOpts, _fullPauser)
}

// SetFullPauser is a paid mutator transaction binding the contract method 0xbeca75d7.
//
// Solidity: function setFullPauser(address _fullPauser) returns()
func (_OtokenController *OtokenControllerTransactorSession) SetFullPauser(_fullPauser common.Address) (*types.Transaction, error) {
	return _OtokenController.Contract.SetFullPauser(&_OtokenController.TransactOpts, _fullPauser)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address _operator, bool _isOperator) returns()
func (_OtokenController *OtokenControllerTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address, _isOperator bool) (*types.Transaction, error) {
	return _OtokenController.contract.Transact(opts, "setOperator", _operator, _isOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address _operator, bool _isOperator) returns()
func (_OtokenController *OtokenControllerSession) SetOperator(_operator common.Address, _isOperator bool) (*types.Transaction, error) {
	return _OtokenController.Contract.SetOperator(&_OtokenController.TransactOpts, _operator, _isOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address _operator, bool _isOperator) returns()
func (_OtokenController *OtokenControllerTransactorSession) SetOperator(_operator common.Address, _isOperator bool) (*types.Transaction, error) {
	return _OtokenController.Contract.SetOperator(&_OtokenController.TransactOpts, _operator, _isOperator)
}

// SetPartialPauser is a paid mutator transaction binding the contract method 0xbc463a3d.
//
// Solidity: function setPartialPauser(address _partialPauser) returns()
func (_OtokenController *OtokenControllerTransactor) SetPartialPauser(opts *bind.TransactOpts, _partialPauser common.Address) (*types.Transaction, error) {
	return _OtokenController.contract.Transact(opts, "setPartialPauser", _partialPauser)
}

// SetPartialPauser is a paid mutator transaction binding the contract method 0xbc463a3d.
//
// Solidity: function setPartialPauser(address _partialPauser) returns()
func (_OtokenController *OtokenControllerSession) SetPartialPauser(_partialPauser common.Address) (*types.Transaction, error) {
	return _OtokenController.Contract.SetPartialPauser(&_OtokenController.TransactOpts, _partialPauser)
}

// SetPartialPauser is a paid mutator transaction binding the contract method 0xbc463a3d.
//
// Solidity: function setPartialPauser(address _partialPauser) returns()
func (_OtokenController *OtokenControllerTransactorSession) SetPartialPauser(_partialPauser common.Address) (*types.Transaction, error) {
	return _OtokenController.Contract.SetPartialPauser(&_OtokenController.TransactOpts, _partialPauser)
}

// SetSystemFullyPaused is a paid mutator transaction binding the contract method 0x9f677ed9.
//
// Solidity: function setSystemFullyPaused(bool _fullyPaused) returns()
func (_OtokenController *OtokenControllerTransactor) SetSystemFullyPaused(opts *bind.TransactOpts, _fullyPaused bool) (*types.Transaction, error) {
	return _OtokenController.contract.Transact(opts, "setSystemFullyPaused", _fullyPaused)
}

// SetSystemFullyPaused is a paid mutator transaction binding the contract method 0x9f677ed9.
//
// Solidity: function setSystemFullyPaused(bool _fullyPaused) returns()
func (_OtokenController *OtokenControllerSession) SetSystemFullyPaused(_fullyPaused bool) (*types.Transaction, error) {
	return _OtokenController.Contract.SetSystemFullyPaused(&_OtokenController.TransactOpts, _fullyPaused)
}

// SetSystemFullyPaused is a paid mutator transaction binding the contract method 0x9f677ed9.
//
// Solidity: function setSystemFullyPaused(bool _fullyPaused) returns()
func (_OtokenController *OtokenControllerTransactorSession) SetSystemFullyPaused(_fullyPaused bool) (*types.Transaction, error) {
	return _OtokenController.Contract.SetSystemFullyPaused(&_OtokenController.TransactOpts, _fullyPaused)
}

// SetSystemPartiallyPaused is a paid mutator transaction binding the contract method 0x573c473e.
//
// Solidity: function setSystemPartiallyPaused(bool _partiallyPaused) returns()
func (_OtokenController *OtokenControllerTransactor) SetSystemPartiallyPaused(opts *bind.TransactOpts, _partiallyPaused bool) (*types.Transaction, error) {
	return _OtokenController.contract.Transact(opts, "setSystemPartiallyPaused", _partiallyPaused)
}

// SetSystemPartiallyPaused is a paid mutator transaction binding the contract method 0x573c473e.
//
// Solidity: function setSystemPartiallyPaused(bool _partiallyPaused) returns()
func (_OtokenController *OtokenControllerSession) SetSystemPartiallyPaused(_partiallyPaused bool) (*types.Transaction, error) {
	return _OtokenController.Contract.SetSystemPartiallyPaused(&_OtokenController.TransactOpts, _partiallyPaused)
}

// SetSystemPartiallyPaused is a paid mutator transaction binding the contract method 0x573c473e.
//
// Solidity: function setSystemPartiallyPaused(bool _partiallyPaused) returns()
func (_OtokenController *OtokenControllerTransactorSession) SetSystemPartiallyPaused(_partiallyPaused bool) (*types.Transaction, error) {
	return _OtokenController.Contract.SetSystemPartiallyPaused(&_OtokenController.TransactOpts, _partiallyPaused)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OtokenController *OtokenControllerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OtokenController.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OtokenController *OtokenControllerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OtokenController.Contract.TransferOwnership(&_OtokenController.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OtokenController *OtokenControllerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OtokenController.Contract.TransferOwnership(&_OtokenController.TransactOpts, newOwner)
}

// OtokenControllerAccountOperatorUpdatedIterator is returned from FilterAccountOperatorUpdated and is used to iterate over the raw logs and unpacked data for AccountOperatorUpdated events raised by the OtokenController contract.
type OtokenControllerAccountOperatorUpdatedIterator struct {
	Event *OtokenControllerAccountOperatorUpdated // Event containing the contract specifics and raw log

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
func (it *OtokenControllerAccountOperatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerAccountOperatorUpdated)
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
		it.Event = new(OtokenControllerAccountOperatorUpdated)
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
func (it *OtokenControllerAccountOperatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerAccountOperatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerAccountOperatorUpdated represents a AccountOperatorUpdated event raised by the OtokenController contract.
type OtokenControllerAccountOperatorUpdated struct {
	AccountOwner common.Address
	Operator     common.Address
	IsSet        bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccountOperatorUpdated is a free log retrieval operation binding the contract event 0x1ae45540b0875932452da603b351ce2f1758258ba1345e79f8fc94f044cb0787.
//
// Solidity: event AccountOperatorUpdated(address indexed accountOwner, address indexed operator, bool isSet)
func (_OtokenController *OtokenControllerFilterer) FilterAccountOperatorUpdated(opts *bind.FilterOpts, accountOwner []common.Address, operator []common.Address) (*OtokenControllerAccountOperatorUpdatedIterator, error) {

	var accountOwnerRule []interface{}
	for _, accountOwnerItem := range accountOwner {
		accountOwnerRule = append(accountOwnerRule, accountOwnerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "AccountOperatorUpdated", accountOwnerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerAccountOperatorUpdatedIterator{contract: _OtokenController.contract, event: "AccountOperatorUpdated", logs: logs, sub: sub}, nil
}

// WatchAccountOperatorUpdated is a free log subscription operation binding the contract event 0x1ae45540b0875932452da603b351ce2f1758258ba1345e79f8fc94f044cb0787.
//
// Solidity: event AccountOperatorUpdated(address indexed accountOwner, address indexed operator, bool isSet)
func (_OtokenController *OtokenControllerFilterer) WatchAccountOperatorUpdated(opts *bind.WatchOpts, sink chan<- *OtokenControllerAccountOperatorUpdated, accountOwner []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountOwnerRule []interface{}
	for _, accountOwnerItem := range accountOwner {
		accountOwnerRule = append(accountOwnerRule, accountOwnerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "AccountOperatorUpdated", accountOwnerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerAccountOperatorUpdated)
				if err := _OtokenController.contract.UnpackLog(event, "AccountOperatorUpdated", log); err != nil {
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

// ParseAccountOperatorUpdated is a log parse operation binding the contract event 0x1ae45540b0875932452da603b351ce2f1758258ba1345e79f8fc94f044cb0787.
//
// Solidity: event AccountOperatorUpdated(address indexed accountOwner, address indexed operator, bool isSet)
func (_OtokenController *OtokenControllerFilterer) ParseAccountOperatorUpdated(log types.Log) (*OtokenControllerAccountOperatorUpdated, error) {
	event := new(OtokenControllerAccountOperatorUpdated)
	if err := _OtokenController.contract.UnpackLog(event, "AccountOperatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerCallExecutedIterator is returned from FilterCallExecuted and is used to iterate over the raw logs and unpacked data for CallExecuted events raised by the OtokenController contract.
type OtokenControllerCallExecutedIterator struct {
	Event *OtokenControllerCallExecuted // Event containing the contract specifics and raw log

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
func (it *OtokenControllerCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerCallExecuted)
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
		it.Event = new(OtokenControllerCallExecuted)
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
func (it *OtokenControllerCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerCallExecuted represents a CallExecuted event raised by the OtokenController contract.
type OtokenControllerCallExecuted struct {
	From common.Address
	To   common.Address
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCallExecuted is a free log retrieval operation binding the contract event 0x8750bdaf6e88201790ee2765fea3ac73b514a52658c818723a30de91029ad000.
//
// Solidity: event CallExecuted(address indexed from, address indexed to, bytes data)
func (_OtokenController *OtokenControllerFilterer) FilterCallExecuted(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OtokenControllerCallExecutedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "CallExecuted", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerCallExecutedIterator{contract: _OtokenController.contract, event: "CallExecuted", logs: logs, sub: sub}, nil
}

// WatchCallExecuted is a free log subscription operation binding the contract event 0x8750bdaf6e88201790ee2765fea3ac73b514a52658c818723a30de91029ad000.
//
// Solidity: event CallExecuted(address indexed from, address indexed to, bytes data)
func (_OtokenController *OtokenControllerFilterer) WatchCallExecuted(opts *bind.WatchOpts, sink chan<- *OtokenControllerCallExecuted, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "CallExecuted", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerCallExecuted)
				if err := _OtokenController.contract.UnpackLog(event, "CallExecuted", log); err != nil {
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

// ParseCallExecuted is a log parse operation binding the contract event 0x8750bdaf6e88201790ee2765fea3ac73b514a52658c818723a30de91029ad000.
//
// Solidity: event CallExecuted(address indexed from, address indexed to, bytes data)
func (_OtokenController *OtokenControllerFilterer) ParseCallExecuted(log types.Log) (*OtokenControllerCallExecuted, error) {
	event := new(OtokenControllerCallExecuted)
	if err := _OtokenController.contract.UnpackLog(event, "CallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerCallRestrictedIterator is returned from FilterCallRestricted and is used to iterate over the raw logs and unpacked data for CallRestricted events raised by the OtokenController contract.
type OtokenControllerCallRestrictedIterator struct {
	Event *OtokenControllerCallRestricted // Event containing the contract specifics and raw log

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
func (it *OtokenControllerCallRestrictedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerCallRestricted)
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
		it.Event = new(OtokenControllerCallRestricted)
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
func (it *OtokenControllerCallRestrictedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerCallRestrictedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerCallRestricted represents a CallRestricted event raised by the OtokenController contract.
type OtokenControllerCallRestricted struct {
	IsRestricted bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCallRestricted is a free log retrieval operation binding the contract event 0x26d614575a4c104c181c87003d4bb00cc7ade00d5b47bf8775171c12a376b255.
//
// Solidity: event CallRestricted(bool isRestricted)
func (_OtokenController *OtokenControllerFilterer) FilterCallRestricted(opts *bind.FilterOpts) (*OtokenControllerCallRestrictedIterator, error) {

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "CallRestricted")
	if err != nil {
		return nil, err
	}
	return &OtokenControllerCallRestrictedIterator{contract: _OtokenController.contract, event: "CallRestricted", logs: logs, sub: sub}, nil
}

// WatchCallRestricted is a free log subscription operation binding the contract event 0x26d614575a4c104c181c87003d4bb00cc7ade00d5b47bf8775171c12a376b255.
//
// Solidity: event CallRestricted(bool isRestricted)
func (_OtokenController *OtokenControllerFilterer) WatchCallRestricted(opts *bind.WatchOpts, sink chan<- *OtokenControllerCallRestricted) (event.Subscription, error) {

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "CallRestricted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerCallRestricted)
				if err := _OtokenController.contract.UnpackLog(event, "CallRestricted", log); err != nil {
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

// ParseCallRestricted is a log parse operation binding the contract event 0x26d614575a4c104c181c87003d4bb00cc7ade00d5b47bf8775171c12a376b255.
//
// Solidity: event CallRestricted(bool isRestricted)
func (_OtokenController *OtokenControllerFilterer) ParseCallRestricted(log types.Log) (*OtokenControllerCallRestricted, error) {
	event := new(OtokenControllerCallRestricted)
	if err := _OtokenController.contract.UnpackLog(event, "CallRestricted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerCollateralAssetDepositedIterator is returned from FilterCollateralAssetDeposited and is used to iterate over the raw logs and unpacked data for CollateralAssetDeposited events raised by the OtokenController contract.
type OtokenControllerCollateralAssetDepositedIterator struct {
	Event *OtokenControllerCollateralAssetDeposited // Event containing the contract specifics and raw log

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
func (it *OtokenControllerCollateralAssetDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerCollateralAssetDeposited)
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
		it.Event = new(OtokenControllerCollateralAssetDeposited)
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
func (it *OtokenControllerCollateralAssetDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerCollateralAssetDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerCollateralAssetDeposited represents a CollateralAssetDeposited event raised by the OtokenController contract.
type OtokenControllerCollateralAssetDeposited struct {
	Asset        common.Address
	AccountOwner common.Address
	From         common.Address
	VaultId      *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCollateralAssetDeposited is a free log retrieval operation binding the contract event 0xbfab88b861f171b7db714f00e5966131253918d55ddba816c3eb94657d102390.
//
// Solidity: event CollateralAssetDeposited(address indexed asset, address indexed accountOwner, address indexed from, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) FilterCollateralAssetDeposited(opts *bind.FilterOpts, asset []common.Address, accountOwner []common.Address, from []common.Address) (*OtokenControllerCollateralAssetDepositedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var accountOwnerRule []interface{}
	for _, accountOwnerItem := range accountOwner {
		accountOwnerRule = append(accountOwnerRule, accountOwnerItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "CollateralAssetDeposited", assetRule, accountOwnerRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerCollateralAssetDepositedIterator{contract: _OtokenController.contract, event: "CollateralAssetDeposited", logs: logs, sub: sub}, nil
}

// WatchCollateralAssetDeposited is a free log subscription operation binding the contract event 0xbfab88b861f171b7db714f00e5966131253918d55ddba816c3eb94657d102390.
//
// Solidity: event CollateralAssetDeposited(address indexed asset, address indexed accountOwner, address indexed from, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) WatchCollateralAssetDeposited(opts *bind.WatchOpts, sink chan<- *OtokenControllerCollateralAssetDeposited, asset []common.Address, accountOwner []common.Address, from []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var accountOwnerRule []interface{}
	for _, accountOwnerItem := range accountOwner {
		accountOwnerRule = append(accountOwnerRule, accountOwnerItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "CollateralAssetDeposited", assetRule, accountOwnerRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerCollateralAssetDeposited)
				if err := _OtokenController.contract.UnpackLog(event, "CollateralAssetDeposited", log); err != nil {
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

// ParseCollateralAssetDeposited is a log parse operation binding the contract event 0xbfab88b861f171b7db714f00e5966131253918d55ddba816c3eb94657d102390.
//
// Solidity: event CollateralAssetDeposited(address indexed asset, address indexed accountOwner, address indexed from, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) ParseCollateralAssetDeposited(log types.Log) (*OtokenControllerCollateralAssetDeposited, error) {
	event := new(OtokenControllerCollateralAssetDeposited)
	if err := _OtokenController.contract.UnpackLog(event, "CollateralAssetDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerCollateralAssetWithdrawedIterator is returned from FilterCollateralAssetWithdrawed and is used to iterate over the raw logs and unpacked data for CollateralAssetWithdrawed events raised by the OtokenController contract.
type OtokenControllerCollateralAssetWithdrawedIterator struct {
	Event *OtokenControllerCollateralAssetWithdrawed // Event containing the contract specifics and raw log

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
func (it *OtokenControllerCollateralAssetWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerCollateralAssetWithdrawed)
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
		it.Event = new(OtokenControllerCollateralAssetWithdrawed)
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
func (it *OtokenControllerCollateralAssetWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerCollateralAssetWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerCollateralAssetWithdrawed represents a CollateralAssetWithdrawed event raised by the OtokenController contract.
type OtokenControllerCollateralAssetWithdrawed struct {
	Asset        common.Address
	AccountOwner common.Address
	To           common.Address
	VaultId      *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCollateralAssetWithdrawed is a free log retrieval operation binding the contract event 0xfe86f7694b6c54a528acbe27be61dd4a85e9a89aeef7f650a1b439045ccee5a4.
//
// Solidity: event CollateralAssetWithdrawed(address indexed asset, address indexed AccountOwner, address indexed to, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) FilterCollateralAssetWithdrawed(opts *bind.FilterOpts, asset []common.Address, AccountOwner []common.Address, to []common.Address) (*OtokenControllerCollateralAssetWithdrawedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var AccountOwnerRule []interface{}
	for _, AccountOwnerItem := range AccountOwner {
		AccountOwnerRule = append(AccountOwnerRule, AccountOwnerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "CollateralAssetWithdrawed", assetRule, AccountOwnerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerCollateralAssetWithdrawedIterator{contract: _OtokenController.contract, event: "CollateralAssetWithdrawed", logs: logs, sub: sub}, nil
}

// WatchCollateralAssetWithdrawed is a free log subscription operation binding the contract event 0xfe86f7694b6c54a528acbe27be61dd4a85e9a89aeef7f650a1b439045ccee5a4.
//
// Solidity: event CollateralAssetWithdrawed(address indexed asset, address indexed AccountOwner, address indexed to, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) WatchCollateralAssetWithdrawed(opts *bind.WatchOpts, sink chan<- *OtokenControllerCollateralAssetWithdrawed, asset []common.Address, AccountOwner []common.Address, to []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var AccountOwnerRule []interface{}
	for _, AccountOwnerItem := range AccountOwner {
		AccountOwnerRule = append(AccountOwnerRule, AccountOwnerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "CollateralAssetWithdrawed", assetRule, AccountOwnerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerCollateralAssetWithdrawed)
				if err := _OtokenController.contract.UnpackLog(event, "CollateralAssetWithdrawed", log); err != nil {
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

// ParseCollateralAssetWithdrawed is a log parse operation binding the contract event 0xfe86f7694b6c54a528acbe27be61dd4a85e9a89aeef7f650a1b439045ccee5a4.
//
// Solidity: event CollateralAssetWithdrawed(address indexed asset, address indexed AccountOwner, address indexed to, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) ParseCollateralAssetWithdrawed(log types.Log) (*OtokenControllerCollateralAssetWithdrawed, error) {
	event := new(OtokenControllerCollateralAssetWithdrawed)
	if err := _OtokenController.contract.UnpackLog(event, "CollateralAssetWithdrawed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerFullPauserUpdatedIterator is returned from FilterFullPauserUpdated and is used to iterate over the raw logs and unpacked data for FullPauserUpdated events raised by the OtokenController contract.
type OtokenControllerFullPauserUpdatedIterator struct {
	Event *OtokenControllerFullPauserUpdated // Event containing the contract specifics and raw log

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
func (it *OtokenControllerFullPauserUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerFullPauserUpdated)
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
		it.Event = new(OtokenControllerFullPauserUpdated)
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
func (it *OtokenControllerFullPauserUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerFullPauserUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerFullPauserUpdated represents a FullPauserUpdated event raised by the OtokenController contract.
type OtokenControllerFullPauserUpdated struct {
	OldFullPauser common.Address
	NewFullPauser common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFullPauserUpdated is a free log retrieval operation binding the contract event 0x44f3e9e7b454118e9fcb2e3026396f57ca21d7bd7dcabd31d7f986806422f4df.
//
// Solidity: event FullPauserUpdated(address indexed oldFullPauser, address indexed newFullPauser)
func (_OtokenController *OtokenControllerFilterer) FilterFullPauserUpdated(opts *bind.FilterOpts, oldFullPauser []common.Address, newFullPauser []common.Address) (*OtokenControllerFullPauserUpdatedIterator, error) {

	var oldFullPauserRule []interface{}
	for _, oldFullPauserItem := range oldFullPauser {
		oldFullPauserRule = append(oldFullPauserRule, oldFullPauserItem)
	}
	var newFullPauserRule []interface{}
	for _, newFullPauserItem := range newFullPauser {
		newFullPauserRule = append(newFullPauserRule, newFullPauserItem)
	}

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "FullPauserUpdated", oldFullPauserRule, newFullPauserRule)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerFullPauserUpdatedIterator{contract: _OtokenController.contract, event: "FullPauserUpdated", logs: logs, sub: sub}, nil
}

// WatchFullPauserUpdated is a free log subscription operation binding the contract event 0x44f3e9e7b454118e9fcb2e3026396f57ca21d7bd7dcabd31d7f986806422f4df.
//
// Solidity: event FullPauserUpdated(address indexed oldFullPauser, address indexed newFullPauser)
func (_OtokenController *OtokenControllerFilterer) WatchFullPauserUpdated(opts *bind.WatchOpts, sink chan<- *OtokenControllerFullPauserUpdated, oldFullPauser []common.Address, newFullPauser []common.Address) (event.Subscription, error) {

	var oldFullPauserRule []interface{}
	for _, oldFullPauserItem := range oldFullPauser {
		oldFullPauserRule = append(oldFullPauserRule, oldFullPauserItem)
	}
	var newFullPauserRule []interface{}
	for _, newFullPauserItem := range newFullPauser {
		newFullPauserRule = append(newFullPauserRule, newFullPauserItem)
	}

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "FullPauserUpdated", oldFullPauserRule, newFullPauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerFullPauserUpdated)
				if err := _OtokenController.contract.UnpackLog(event, "FullPauserUpdated", log); err != nil {
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

// ParseFullPauserUpdated is a log parse operation binding the contract event 0x44f3e9e7b454118e9fcb2e3026396f57ca21d7bd7dcabd31d7f986806422f4df.
//
// Solidity: event FullPauserUpdated(address indexed oldFullPauser, address indexed newFullPauser)
func (_OtokenController *OtokenControllerFilterer) ParseFullPauserUpdated(log types.Log) (*OtokenControllerFullPauserUpdated, error) {
	event := new(OtokenControllerFullPauserUpdated)
	if err := _OtokenController.contract.UnpackLog(event, "FullPauserUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerLongOtokenDepositedIterator is returned from FilterLongOtokenDeposited and is used to iterate over the raw logs and unpacked data for LongOtokenDeposited events raised by the OtokenController contract.
type OtokenControllerLongOtokenDepositedIterator struct {
	Event *OtokenControllerLongOtokenDeposited // Event containing the contract specifics and raw log

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
func (it *OtokenControllerLongOtokenDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerLongOtokenDeposited)
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
		it.Event = new(OtokenControllerLongOtokenDeposited)
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
func (it *OtokenControllerLongOtokenDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerLongOtokenDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerLongOtokenDeposited represents a LongOtokenDeposited event raised by the OtokenController contract.
type OtokenControllerLongOtokenDeposited struct {
	Otoken       common.Address
	AccountOwner common.Address
	From         common.Address
	VaultId      *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLongOtokenDeposited is a free log retrieval operation binding the contract event 0x2607e210004cef0ad6b3e6aedb778bffb03c1586f8dcf55d49afffde210d8bb1.
//
// Solidity: event LongOtokenDeposited(address indexed otoken, address indexed accountOwner, address indexed from, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) FilterLongOtokenDeposited(opts *bind.FilterOpts, otoken []common.Address, accountOwner []common.Address, from []common.Address) (*OtokenControllerLongOtokenDepositedIterator, error) {

	var otokenRule []interface{}
	for _, otokenItem := range otoken {
		otokenRule = append(otokenRule, otokenItem)
	}
	var accountOwnerRule []interface{}
	for _, accountOwnerItem := range accountOwner {
		accountOwnerRule = append(accountOwnerRule, accountOwnerItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "LongOtokenDeposited", otokenRule, accountOwnerRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerLongOtokenDepositedIterator{contract: _OtokenController.contract, event: "LongOtokenDeposited", logs: logs, sub: sub}, nil
}

// WatchLongOtokenDeposited is a free log subscription operation binding the contract event 0x2607e210004cef0ad6b3e6aedb778bffb03c1586f8dcf55d49afffde210d8bb1.
//
// Solidity: event LongOtokenDeposited(address indexed otoken, address indexed accountOwner, address indexed from, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) WatchLongOtokenDeposited(opts *bind.WatchOpts, sink chan<- *OtokenControllerLongOtokenDeposited, otoken []common.Address, accountOwner []common.Address, from []common.Address) (event.Subscription, error) {

	var otokenRule []interface{}
	for _, otokenItem := range otoken {
		otokenRule = append(otokenRule, otokenItem)
	}
	var accountOwnerRule []interface{}
	for _, accountOwnerItem := range accountOwner {
		accountOwnerRule = append(accountOwnerRule, accountOwnerItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "LongOtokenDeposited", otokenRule, accountOwnerRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerLongOtokenDeposited)
				if err := _OtokenController.contract.UnpackLog(event, "LongOtokenDeposited", log); err != nil {
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

// ParseLongOtokenDeposited is a log parse operation binding the contract event 0x2607e210004cef0ad6b3e6aedb778bffb03c1586f8dcf55d49afffde210d8bb1.
//
// Solidity: event LongOtokenDeposited(address indexed otoken, address indexed accountOwner, address indexed from, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) ParseLongOtokenDeposited(log types.Log) (*OtokenControllerLongOtokenDeposited, error) {
	event := new(OtokenControllerLongOtokenDeposited)
	if err := _OtokenController.contract.UnpackLog(event, "LongOtokenDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerLongOtokenWithdrawedIterator is returned from FilterLongOtokenWithdrawed and is used to iterate over the raw logs and unpacked data for LongOtokenWithdrawed events raised by the OtokenController contract.
type OtokenControllerLongOtokenWithdrawedIterator struct {
	Event *OtokenControllerLongOtokenWithdrawed // Event containing the contract specifics and raw log

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
func (it *OtokenControllerLongOtokenWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerLongOtokenWithdrawed)
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
		it.Event = new(OtokenControllerLongOtokenWithdrawed)
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
func (it *OtokenControllerLongOtokenWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerLongOtokenWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerLongOtokenWithdrawed represents a LongOtokenWithdrawed event raised by the OtokenController contract.
type OtokenControllerLongOtokenWithdrawed struct {
	Otoken       common.Address
	AccountOwner common.Address
	To           common.Address
	VaultId      *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLongOtokenWithdrawed is a free log retrieval operation binding the contract event 0xbd023c53d293da163d185720d4274f4ddabc09d5304491a55abb296cc811d9fa.
//
// Solidity: event LongOtokenWithdrawed(address indexed otoken, address indexed AccountOwner, address indexed to, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) FilterLongOtokenWithdrawed(opts *bind.FilterOpts, otoken []common.Address, AccountOwner []common.Address, to []common.Address) (*OtokenControllerLongOtokenWithdrawedIterator, error) {

	var otokenRule []interface{}
	for _, otokenItem := range otoken {
		otokenRule = append(otokenRule, otokenItem)
	}
	var AccountOwnerRule []interface{}
	for _, AccountOwnerItem := range AccountOwner {
		AccountOwnerRule = append(AccountOwnerRule, AccountOwnerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "LongOtokenWithdrawed", otokenRule, AccountOwnerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerLongOtokenWithdrawedIterator{contract: _OtokenController.contract, event: "LongOtokenWithdrawed", logs: logs, sub: sub}, nil
}

// WatchLongOtokenWithdrawed is a free log subscription operation binding the contract event 0xbd023c53d293da163d185720d4274f4ddabc09d5304491a55abb296cc811d9fa.
//
// Solidity: event LongOtokenWithdrawed(address indexed otoken, address indexed AccountOwner, address indexed to, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) WatchLongOtokenWithdrawed(opts *bind.WatchOpts, sink chan<- *OtokenControllerLongOtokenWithdrawed, otoken []common.Address, AccountOwner []common.Address, to []common.Address) (event.Subscription, error) {

	var otokenRule []interface{}
	for _, otokenItem := range otoken {
		otokenRule = append(otokenRule, otokenItem)
	}
	var AccountOwnerRule []interface{}
	for _, AccountOwnerItem := range AccountOwner {
		AccountOwnerRule = append(AccountOwnerRule, AccountOwnerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "LongOtokenWithdrawed", otokenRule, AccountOwnerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerLongOtokenWithdrawed)
				if err := _OtokenController.contract.UnpackLog(event, "LongOtokenWithdrawed", log); err != nil {
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

// ParseLongOtokenWithdrawed is a log parse operation binding the contract event 0xbd023c53d293da163d185720d4274f4ddabc09d5304491a55abb296cc811d9fa.
//
// Solidity: event LongOtokenWithdrawed(address indexed otoken, address indexed AccountOwner, address indexed to, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) ParseLongOtokenWithdrawed(log types.Log) (*OtokenControllerLongOtokenWithdrawed, error) {
	event := new(OtokenControllerLongOtokenWithdrawed)
	if err := _OtokenController.contract.UnpackLog(event, "LongOtokenWithdrawed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OtokenController contract.
type OtokenControllerOwnershipTransferredIterator struct {
	Event *OtokenControllerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OtokenControllerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerOwnershipTransferred)
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
		it.Event = new(OtokenControllerOwnershipTransferred)
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
func (it *OtokenControllerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerOwnershipTransferred represents a OwnershipTransferred event raised by the OtokenController contract.
type OtokenControllerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OtokenController *OtokenControllerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OtokenControllerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerOwnershipTransferredIterator{contract: _OtokenController.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OtokenController *OtokenControllerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OtokenControllerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerOwnershipTransferred)
				if err := _OtokenController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OtokenController *OtokenControllerFilterer) ParseOwnershipTransferred(log types.Log) (*OtokenControllerOwnershipTransferred, error) {
	event := new(OtokenControllerOwnershipTransferred)
	if err := _OtokenController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerPartialPauserUpdatedIterator is returned from FilterPartialPauserUpdated and is used to iterate over the raw logs and unpacked data for PartialPauserUpdated events raised by the OtokenController contract.
type OtokenControllerPartialPauserUpdatedIterator struct {
	Event *OtokenControllerPartialPauserUpdated // Event containing the contract specifics and raw log

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
func (it *OtokenControllerPartialPauserUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerPartialPauserUpdated)
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
		it.Event = new(OtokenControllerPartialPauserUpdated)
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
func (it *OtokenControllerPartialPauserUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerPartialPauserUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerPartialPauserUpdated represents a PartialPauserUpdated event raised by the OtokenController contract.
type OtokenControllerPartialPauserUpdated struct {
	OldPartialPauser common.Address
	NewPartialPauser common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPartialPauserUpdated is a free log retrieval operation binding the contract event 0x1440312dbc326ddc21bfa95078324bf5aaf6899e8a27cba3057c60adfc84e40b.
//
// Solidity: event PartialPauserUpdated(address indexed oldPartialPauser, address indexed newPartialPauser)
func (_OtokenController *OtokenControllerFilterer) FilterPartialPauserUpdated(opts *bind.FilterOpts, oldPartialPauser []common.Address, newPartialPauser []common.Address) (*OtokenControllerPartialPauserUpdatedIterator, error) {

	var oldPartialPauserRule []interface{}
	for _, oldPartialPauserItem := range oldPartialPauser {
		oldPartialPauserRule = append(oldPartialPauserRule, oldPartialPauserItem)
	}
	var newPartialPauserRule []interface{}
	for _, newPartialPauserItem := range newPartialPauser {
		newPartialPauserRule = append(newPartialPauserRule, newPartialPauserItem)
	}

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "PartialPauserUpdated", oldPartialPauserRule, newPartialPauserRule)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerPartialPauserUpdatedIterator{contract: _OtokenController.contract, event: "PartialPauserUpdated", logs: logs, sub: sub}, nil
}

// WatchPartialPauserUpdated is a free log subscription operation binding the contract event 0x1440312dbc326ddc21bfa95078324bf5aaf6899e8a27cba3057c60adfc84e40b.
//
// Solidity: event PartialPauserUpdated(address indexed oldPartialPauser, address indexed newPartialPauser)
func (_OtokenController *OtokenControllerFilterer) WatchPartialPauserUpdated(opts *bind.WatchOpts, sink chan<- *OtokenControllerPartialPauserUpdated, oldPartialPauser []common.Address, newPartialPauser []common.Address) (event.Subscription, error) {

	var oldPartialPauserRule []interface{}
	for _, oldPartialPauserItem := range oldPartialPauser {
		oldPartialPauserRule = append(oldPartialPauserRule, oldPartialPauserItem)
	}
	var newPartialPauserRule []interface{}
	for _, newPartialPauserItem := range newPartialPauser {
		newPartialPauserRule = append(newPartialPauserRule, newPartialPauserItem)
	}

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "PartialPauserUpdated", oldPartialPauserRule, newPartialPauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerPartialPauserUpdated)
				if err := _OtokenController.contract.UnpackLog(event, "PartialPauserUpdated", log); err != nil {
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

// ParsePartialPauserUpdated is a log parse operation binding the contract event 0x1440312dbc326ddc21bfa95078324bf5aaf6899e8a27cba3057c60adfc84e40b.
//
// Solidity: event PartialPauserUpdated(address indexed oldPartialPauser, address indexed newPartialPauser)
func (_OtokenController *OtokenControllerFilterer) ParsePartialPauserUpdated(log types.Log) (*OtokenControllerPartialPauserUpdated, error) {
	event := new(OtokenControllerPartialPauserUpdated)
	if err := _OtokenController.contract.UnpackLog(event, "PartialPauserUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the OtokenController contract.
type OtokenControllerRedeemIterator struct {
	Event *OtokenControllerRedeem // Event containing the contract specifics and raw log

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
func (it *OtokenControllerRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerRedeem)
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
		it.Event = new(OtokenControllerRedeem)
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
func (it *OtokenControllerRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerRedeem represents a Redeem event raised by the OtokenController contract.
type OtokenControllerRedeem struct {
	Otoken          common.Address
	Redeemer        common.Address
	Receiver        common.Address
	CollateralAsset common.Address
	OtokenBurned    *big.Int
	Payout          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0x18fd144d7dbcbaa6f00fd47a84adc7dc3cc64a326ffa2dc7691a25e3837dba03.
//
// Solidity: event Redeem(address indexed otoken, address indexed redeemer, address indexed receiver, address collateralAsset, uint256 otokenBurned, uint256 payout)
func (_OtokenController *OtokenControllerFilterer) FilterRedeem(opts *bind.FilterOpts, otoken []common.Address, redeemer []common.Address, receiver []common.Address) (*OtokenControllerRedeemIterator, error) {

	var otokenRule []interface{}
	for _, otokenItem := range otoken {
		otokenRule = append(otokenRule, otokenItem)
	}
	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "Redeem", otokenRule, redeemerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerRedeemIterator{contract: _OtokenController.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0x18fd144d7dbcbaa6f00fd47a84adc7dc3cc64a326ffa2dc7691a25e3837dba03.
//
// Solidity: event Redeem(address indexed otoken, address indexed redeemer, address indexed receiver, address collateralAsset, uint256 otokenBurned, uint256 payout)
func (_OtokenController *OtokenControllerFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *OtokenControllerRedeem, otoken []common.Address, redeemer []common.Address, receiver []common.Address) (event.Subscription, error) {

	var otokenRule []interface{}
	for _, otokenItem := range otoken {
		otokenRule = append(otokenRule, otokenItem)
	}
	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "Redeem", otokenRule, redeemerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerRedeem)
				if err := _OtokenController.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0x18fd144d7dbcbaa6f00fd47a84adc7dc3cc64a326ffa2dc7691a25e3837dba03.
//
// Solidity: event Redeem(address indexed otoken, address indexed redeemer, address indexed receiver, address collateralAsset, uint256 otokenBurned, uint256 payout)
func (_OtokenController *OtokenControllerFilterer) ParseRedeem(log types.Log) (*OtokenControllerRedeem, error) {
	event := new(OtokenControllerRedeem)
	if err := _OtokenController.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerShortOtokenBurnedIterator is returned from FilterShortOtokenBurned and is used to iterate over the raw logs and unpacked data for ShortOtokenBurned events raised by the OtokenController contract.
type OtokenControllerShortOtokenBurnedIterator struct {
	Event *OtokenControllerShortOtokenBurned // Event containing the contract specifics and raw log

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
func (it *OtokenControllerShortOtokenBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerShortOtokenBurned)
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
		it.Event = new(OtokenControllerShortOtokenBurned)
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
func (it *OtokenControllerShortOtokenBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerShortOtokenBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerShortOtokenBurned represents a ShortOtokenBurned event raised by the OtokenController contract.
type OtokenControllerShortOtokenBurned struct {
	Otoken       common.Address
	AccountOwner common.Address
	From         common.Address
	VaultId      *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterShortOtokenBurned is a free log retrieval operation binding the contract event 0xdd96b18f26fd9950581b9fd821fa907fc318845fc4d220b825a7b19bfdd174e8.
//
// Solidity: event ShortOtokenBurned(address indexed otoken, address indexed AccountOwner, address indexed from, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) FilterShortOtokenBurned(opts *bind.FilterOpts, otoken []common.Address, AccountOwner []common.Address, from []common.Address) (*OtokenControllerShortOtokenBurnedIterator, error) {

	var otokenRule []interface{}
	for _, otokenItem := range otoken {
		otokenRule = append(otokenRule, otokenItem)
	}
	var AccountOwnerRule []interface{}
	for _, AccountOwnerItem := range AccountOwner {
		AccountOwnerRule = append(AccountOwnerRule, AccountOwnerItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "ShortOtokenBurned", otokenRule, AccountOwnerRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerShortOtokenBurnedIterator{contract: _OtokenController.contract, event: "ShortOtokenBurned", logs: logs, sub: sub}, nil
}

// WatchShortOtokenBurned is a free log subscription operation binding the contract event 0xdd96b18f26fd9950581b9fd821fa907fc318845fc4d220b825a7b19bfdd174e8.
//
// Solidity: event ShortOtokenBurned(address indexed otoken, address indexed AccountOwner, address indexed from, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) WatchShortOtokenBurned(opts *bind.WatchOpts, sink chan<- *OtokenControllerShortOtokenBurned, otoken []common.Address, AccountOwner []common.Address, from []common.Address) (event.Subscription, error) {

	var otokenRule []interface{}
	for _, otokenItem := range otoken {
		otokenRule = append(otokenRule, otokenItem)
	}
	var AccountOwnerRule []interface{}
	for _, AccountOwnerItem := range AccountOwner {
		AccountOwnerRule = append(AccountOwnerRule, AccountOwnerItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "ShortOtokenBurned", otokenRule, AccountOwnerRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerShortOtokenBurned)
				if err := _OtokenController.contract.UnpackLog(event, "ShortOtokenBurned", log); err != nil {
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

// ParseShortOtokenBurned is a log parse operation binding the contract event 0xdd96b18f26fd9950581b9fd821fa907fc318845fc4d220b825a7b19bfdd174e8.
//
// Solidity: event ShortOtokenBurned(address indexed otoken, address indexed AccountOwner, address indexed from, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) ParseShortOtokenBurned(log types.Log) (*OtokenControllerShortOtokenBurned, error) {
	event := new(OtokenControllerShortOtokenBurned)
	if err := _OtokenController.contract.UnpackLog(event, "ShortOtokenBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerShortOtokenMintedIterator is returned from FilterShortOtokenMinted and is used to iterate over the raw logs and unpacked data for ShortOtokenMinted events raised by the OtokenController contract.
type OtokenControllerShortOtokenMintedIterator struct {
	Event *OtokenControllerShortOtokenMinted // Event containing the contract specifics and raw log

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
func (it *OtokenControllerShortOtokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerShortOtokenMinted)
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
		it.Event = new(OtokenControllerShortOtokenMinted)
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
func (it *OtokenControllerShortOtokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerShortOtokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerShortOtokenMinted represents a ShortOtokenMinted event raised by the OtokenController contract.
type OtokenControllerShortOtokenMinted struct {
	Otoken       common.Address
	AccountOwner common.Address
	To           common.Address
	VaultId      *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterShortOtokenMinted is a free log retrieval operation binding the contract event 0x4d7f96086c92b2f9a254ad21548b1c1f2d99502c7949508866349b96bb1a8d8a.
//
// Solidity: event ShortOtokenMinted(address indexed otoken, address indexed AccountOwner, address indexed to, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) FilterShortOtokenMinted(opts *bind.FilterOpts, otoken []common.Address, AccountOwner []common.Address, to []common.Address) (*OtokenControllerShortOtokenMintedIterator, error) {

	var otokenRule []interface{}
	for _, otokenItem := range otoken {
		otokenRule = append(otokenRule, otokenItem)
	}
	var AccountOwnerRule []interface{}
	for _, AccountOwnerItem := range AccountOwner {
		AccountOwnerRule = append(AccountOwnerRule, AccountOwnerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "ShortOtokenMinted", otokenRule, AccountOwnerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerShortOtokenMintedIterator{contract: _OtokenController.contract, event: "ShortOtokenMinted", logs: logs, sub: sub}, nil
}

// WatchShortOtokenMinted is a free log subscription operation binding the contract event 0x4d7f96086c92b2f9a254ad21548b1c1f2d99502c7949508866349b96bb1a8d8a.
//
// Solidity: event ShortOtokenMinted(address indexed otoken, address indexed AccountOwner, address indexed to, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) WatchShortOtokenMinted(opts *bind.WatchOpts, sink chan<- *OtokenControllerShortOtokenMinted, otoken []common.Address, AccountOwner []common.Address, to []common.Address) (event.Subscription, error) {

	var otokenRule []interface{}
	for _, otokenItem := range otoken {
		otokenRule = append(otokenRule, otokenItem)
	}
	var AccountOwnerRule []interface{}
	for _, AccountOwnerItem := range AccountOwner {
		AccountOwnerRule = append(AccountOwnerRule, AccountOwnerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "ShortOtokenMinted", otokenRule, AccountOwnerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerShortOtokenMinted)
				if err := _OtokenController.contract.UnpackLog(event, "ShortOtokenMinted", log); err != nil {
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

// ParseShortOtokenMinted is a log parse operation binding the contract event 0x4d7f96086c92b2f9a254ad21548b1c1f2d99502c7949508866349b96bb1a8d8a.
//
// Solidity: event ShortOtokenMinted(address indexed otoken, address indexed AccountOwner, address indexed to, uint256 vaultId, uint256 amount)
func (_OtokenController *OtokenControllerFilterer) ParseShortOtokenMinted(log types.Log) (*OtokenControllerShortOtokenMinted, error) {
	event := new(OtokenControllerShortOtokenMinted)
	if err := _OtokenController.contract.UnpackLog(event, "ShortOtokenMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerSystemFullyPausedIterator is returned from FilterSystemFullyPaused and is used to iterate over the raw logs and unpacked data for SystemFullyPaused events raised by the OtokenController contract.
type OtokenControllerSystemFullyPausedIterator struct {
	Event *OtokenControllerSystemFullyPaused // Event containing the contract specifics and raw log

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
func (it *OtokenControllerSystemFullyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerSystemFullyPaused)
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
		it.Event = new(OtokenControllerSystemFullyPaused)
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
func (it *OtokenControllerSystemFullyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerSystemFullyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerSystemFullyPaused represents a SystemFullyPaused event raised by the OtokenController contract.
type OtokenControllerSystemFullyPaused struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSystemFullyPaused is a free log retrieval operation binding the contract event 0x4f1d0445688d95c99ca9fc036f551b205fd18ff26a4443b1979c16d1ba66b535.
//
// Solidity: event SystemFullyPaused(bool isPaused)
func (_OtokenController *OtokenControllerFilterer) FilterSystemFullyPaused(opts *bind.FilterOpts) (*OtokenControllerSystemFullyPausedIterator, error) {

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "SystemFullyPaused")
	if err != nil {
		return nil, err
	}
	return &OtokenControllerSystemFullyPausedIterator{contract: _OtokenController.contract, event: "SystemFullyPaused", logs: logs, sub: sub}, nil
}

// WatchSystemFullyPaused is a free log subscription operation binding the contract event 0x4f1d0445688d95c99ca9fc036f551b205fd18ff26a4443b1979c16d1ba66b535.
//
// Solidity: event SystemFullyPaused(bool isPaused)
func (_OtokenController *OtokenControllerFilterer) WatchSystemFullyPaused(opts *bind.WatchOpts, sink chan<- *OtokenControllerSystemFullyPaused) (event.Subscription, error) {

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "SystemFullyPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerSystemFullyPaused)
				if err := _OtokenController.contract.UnpackLog(event, "SystemFullyPaused", log); err != nil {
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

// ParseSystemFullyPaused is a log parse operation binding the contract event 0x4f1d0445688d95c99ca9fc036f551b205fd18ff26a4443b1979c16d1ba66b535.
//
// Solidity: event SystemFullyPaused(bool isPaused)
func (_OtokenController *OtokenControllerFilterer) ParseSystemFullyPaused(log types.Log) (*OtokenControllerSystemFullyPaused, error) {
	event := new(OtokenControllerSystemFullyPaused)
	if err := _OtokenController.contract.UnpackLog(event, "SystemFullyPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerSystemPartiallyPausedIterator is returned from FilterSystemPartiallyPaused and is used to iterate over the raw logs and unpacked data for SystemPartiallyPaused events raised by the OtokenController contract.
type OtokenControllerSystemPartiallyPausedIterator struct {
	Event *OtokenControllerSystemPartiallyPaused // Event containing the contract specifics and raw log

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
func (it *OtokenControllerSystemPartiallyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerSystemPartiallyPaused)
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
		it.Event = new(OtokenControllerSystemPartiallyPaused)
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
func (it *OtokenControllerSystemPartiallyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerSystemPartiallyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerSystemPartiallyPaused represents a SystemPartiallyPaused event raised by the OtokenController contract.
type OtokenControllerSystemPartiallyPaused struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSystemPartiallyPaused is a free log retrieval operation binding the contract event 0x531c3d7229f510a8da00a0f5792686958cdd9c8a120c3c030a6053cd66b68556.
//
// Solidity: event SystemPartiallyPaused(bool isPaused)
func (_OtokenController *OtokenControllerFilterer) FilterSystemPartiallyPaused(opts *bind.FilterOpts) (*OtokenControllerSystemPartiallyPausedIterator, error) {

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "SystemPartiallyPaused")
	if err != nil {
		return nil, err
	}
	return &OtokenControllerSystemPartiallyPausedIterator{contract: _OtokenController.contract, event: "SystemPartiallyPaused", logs: logs, sub: sub}, nil
}

// WatchSystemPartiallyPaused is a free log subscription operation binding the contract event 0x531c3d7229f510a8da00a0f5792686958cdd9c8a120c3c030a6053cd66b68556.
//
// Solidity: event SystemPartiallyPaused(bool isPaused)
func (_OtokenController *OtokenControllerFilterer) WatchSystemPartiallyPaused(opts *bind.WatchOpts, sink chan<- *OtokenControllerSystemPartiallyPaused) (event.Subscription, error) {

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "SystemPartiallyPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerSystemPartiallyPaused)
				if err := _OtokenController.contract.UnpackLog(event, "SystemPartiallyPaused", log); err != nil {
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

// ParseSystemPartiallyPaused is a log parse operation binding the contract event 0x531c3d7229f510a8da00a0f5792686958cdd9c8a120c3c030a6053cd66b68556.
//
// Solidity: event SystemPartiallyPaused(bool isPaused)
func (_OtokenController *OtokenControllerFilterer) ParseSystemPartiallyPaused(log types.Log) (*OtokenControllerSystemPartiallyPaused, error) {
	event := new(OtokenControllerSystemPartiallyPaused)
	if err := _OtokenController.contract.UnpackLog(event, "SystemPartiallyPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerVaultOpenedIterator is returned from FilterVaultOpened and is used to iterate over the raw logs and unpacked data for VaultOpened events raised by the OtokenController contract.
type OtokenControllerVaultOpenedIterator struct {
	Event *OtokenControllerVaultOpened // Event containing the contract specifics and raw log

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
func (it *OtokenControllerVaultOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerVaultOpened)
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
		it.Event = new(OtokenControllerVaultOpened)
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
func (it *OtokenControllerVaultOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerVaultOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerVaultOpened represents a VaultOpened event raised by the OtokenController contract.
type OtokenControllerVaultOpened struct {
	AccountOwner common.Address
	VaultId      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVaultOpened is a free log retrieval operation binding the contract event 0x1774e23245eedaa4541d7c6aafe9840fbea9a6dc740c21dfbe2ecf3162789cb7.
//
// Solidity: event VaultOpened(address indexed accountOwner, uint256 vaultId)
func (_OtokenController *OtokenControllerFilterer) FilterVaultOpened(opts *bind.FilterOpts, accountOwner []common.Address) (*OtokenControllerVaultOpenedIterator, error) {

	var accountOwnerRule []interface{}
	for _, accountOwnerItem := range accountOwner {
		accountOwnerRule = append(accountOwnerRule, accountOwnerItem)
	}

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "VaultOpened", accountOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerVaultOpenedIterator{contract: _OtokenController.contract, event: "VaultOpened", logs: logs, sub: sub}, nil
}

// WatchVaultOpened is a free log subscription operation binding the contract event 0x1774e23245eedaa4541d7c6aafe9840fbea9a6dc740c21dfbe2ecf3162789cb7.
//
// Solidity: event VaultOpened(address indexed accountOwner, uint256 vaultId)
func (_OtokenController *OtokenControllerFilterer) WatchVaultOpened(opts *bind.WatchOpts, sink chan<- *OtokenControllerVaultOpened, accountOwner []common.Address) (event.Subscription, error) {

	var accountOwnerRule []interface{}
	for _, accountOwnerItem := range accountOwner {
		accountOwnerRule = append(accountOwnerRule, accountOwnerItem)
	}

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "VaultOpened", accountOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerVaultOpened)
				if err := _OtokenController.contract.UnpackLog(event, "VaultOpened", log); err != nil {
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

// ParseVaultOpened is a log parse operation binding the contract event 0x1774e23245eedaa4541d7c6aafe9840fbea9a6dc740c21dfbe2ecf3162789cb7.
//
// Solidity: event VaultOpened(address indexed accountOwner, uint256 vaultId)
func (_OtokenController *OtokenControllerFilterer) ParseVaultOpened(log types.Log) (*OtokenControllerVaultOpened, error) {
	event := new(OtokenControllerVaultOpened)
	if err := _OtokenController.contract.UnpackLog(event, "VaultOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OtokenControllerVaultSettledIterator is returned from FilterVaultSettled and is used to iterate over the raw logs and unpacked data for VaultSettled events raised by the OtokenController contract.
type OtokenControllerVaultSettledIterator struct {
	Event *OtokenControllerVaultSettled // Event containing the contract specifics and raw log

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
func (it *OtokenControllerVaultSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtokenControllerVaultSettled)
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
		it.Event = new(OtokenControllerVaultSettled)
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
func (it *OtokenControllerVaultSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtokenControllerVaultSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtokenControllerVaultSettled represents a VaultSettled event raised by the OtokenController contract.
type OtokenControllerVaultSettled struct {
	AccountOwner common.Address
	To           common.Address
	Otoken       common.Address
	VaultId      *big.Int
	Payout       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVaultSettled is a free log retrieval operation binding the contract event 0x14228c91fa4bb9206a82a4fb6af9b066b035cf69957e4a51fee88f693a06bbd0.
//
// Solidity: event VaultSettled(address indexed AccountOwner, address indexed to, address indexed otoken, uint256 vaultId, uint256 payout)
func (_OtokenController *OtokenControllerFilterer) FilterVaultSettled(opts *bind.FilterOpts, AccountOwner []common.Address, to []common.Address, otoken []common.Address) (*OtokenControllerVaultSettledIterator, error) {

	var AccountOwnerRule []interface{}
	for _, AccountOwnerItem := range AccountOwner {
		AccountOwnerRule = append(AccountOwnerRule, AccountOwnerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var otokenRule []interface{}
	for _, otokenItem := range otoken {
		otokenRule = append(otokenRule, otokenItem)
	}

	logs, sub, err := _OtokenController.contract.FilterLogs(opts, "VaultSettled", AccountOwnerRule, toRule, otokenRule)
	if err != nil {
		return nil, err
	}
	return &OtokenControllerVaultSettledIterator{contract: _OtokenController.contract, event: "VaultSettled", logs: logs, sub: sub}, nil
}

// WatchVaultSettled is a free log subscription operation binding the contract event 0x14228c91fa4bb9206a82a4fb6af9b066b035cf69957e4a51fee88f693a06bbd0.
//
// Solidity: event VaultSettled(address indexed AccountOwner, address indexed to, address indexed otoken, uint256 vaultId, uint256 payout)
func (_OtokenController *OtokenControllerFilterer) WatchVaultSettled(opts *bind.WatchOpts, sink chan<- *OtokenControllerVaultSettled, AccountOwner []common.Address, to []common.Address, otoken []common.Address) (event.Subscription, error) {

	var AccountOwnerRule []interface{}
	for _, AccountOwnerItem := range AccountOwner {
		AccountOwnerRule = append(AccountOwnerRule, AccountOwnerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var otokenRule []interface{}
	for _, otokenItem := range otoken {
		otokenRule = append(otokenRule, otokenItem)
	}

	logs, sub, err := _OtokenController.contract.WatchLogs(opts, "VaultSettled", AccountOwnerRule, toRule, otokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtokenControllerVaultSettled)
				if err := _OtokenController.contract.UnpackLog(event, "VaultSettled", log); err != nil {
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

// ParseVaultSettled is a log parse operation binding the contract event 0x14228c91fa4bb9206a82a4fb6af9b066b035cf69957e4a51fee88f693a06bbd0.
//
// Solidity: event VaultSettled(address indexed AccountOwner, address indexed to, address indexed otoken, uint256 vaultId, uint256 payout)
func (_OtokenController *OtokenControllerFilterer) ParseVaultSettled(log types.Log) (*OtokenControllerVaultSettled, error) {
	event := new(OtokenControllerVaultSettled)
	if err := _OtokenController.contract.UnpackLog(event, "VaultSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
