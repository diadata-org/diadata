// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gnosis

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

// GnosisABI is the input ABI used to generate the binding from.
const GnosisABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"IMPROVEMENT_DENOMINATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSecondsRemainingInBatch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEncodedOrders\",\"outputs\":[{\"name\":\"elements\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"buyToken\",\"type\":\"uint16\"},{\"name\":\"sellToken\",\"type\":\"uint16\"},{\"name\":\"validUntil\",\"type\":\"uint32\"},{\"name\":\"buyAmount\",\"type\":\"uint128\"},{\"name\":\"sellAmount\",\"type\":\"uint128\"}],\"name\":\"placeOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"batchId\",\"type\":\"uint32\"},{\"name\":\"claimedObjectiveValue\",\"type\":\"uint256\"},{\"name\":\"owners\",\"type\":\"address[]\"},{\"name\":\"orderIds\",\"type\":\"uint16[]\"},{\"name\":\"buyVolumes\",\"type\":\"uint128[]\"},{\"name\":\"prices\",\"type\":\"uint128[]\"},{\"name\":\"tokenIdsForPrice\",\"type\":\"uint16[]\"}],\"name\":\"submitSolution\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint16\"}],\"name\":\"tokenIdToAddressMap\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_FOR_LISTING_TOKEN_IN_OWL\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"previousPageUser\",\"type\":\"address\"},{\"name\":\"pageSize\",\"type\":\"uint16\"}],\"name\":\"getUsersPaginated\",\"outputs\":[{\"name\":\"users\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderIds\",\"type\":\"uint16[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AMOUNT_MINIMUM\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"buyTokens\",\"type\":\"uint16[]\"},{\"name\":\"sellTokens\",\"type\":\"uint16[]\"},{\"name\":\"validFroms\",\"type\":\"uint32[]\"},{\"name\":\"validUntils\",\"type\":\"uint32[]\"},{\"name\":\"buyAmounts\",\"type\":\"uint128[]\"},{\"name\":\"sellAmounts\",\"type\":\"uint128[]\"}],\"name\":\"placeValidFromOrders\",\"outputs\":[{\"name\":\"orderIds\",\"type\":\"uint16[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"currentPrices\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getEncodedUserOrders\",\"outputs\":[{\"name\":\"elements\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"name\":\"buyToken\",\"type\":\"uint16\"},{\"name\":\"sellToken\",\"type\":\"uint16\"},{\"name\":\"validFrom\",\"type\":\"uint32\"},{\"name\":\"validUntil\",\"type\":\"uint32\"},{\"name\":\"priceNumerator\",\"type\":\"uint128\"},{\"name\":\"priceDenominator\",\"type\":\"uint128\"},{\"name\":\"usedAmount\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"UNLIMITED_ORDER_AMOUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastCreditBatchId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"previousPageUser\",\"type\":\"address\"},{\"name\":\"previousPageUserOffset\",\"type\":\"uint16\"},{\"name\":\"pageSize\",\"type\":\"uint16\"}],\"name\":\"getEncodedUsersPaginated\",\"outputs\":[{\"name\":\"elements\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"hasToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestSolution\",\"outputs\":[{\"name\":\"batchId\",\"type\":\"uint32\"},{\"name\":\"solutionSubmitter\",\"type\":\"address\"},{\"name\":\"feeReward\",\"type\":\"uint256\"},{\"name\":\"objectiveValue\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPendingDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cancellations\",\"type\":\"uint16[]\"},{\"name\":\"buyTokens\",\"type\":\"uint16[]\"},{\"name\":\"sellTokens\",\"type\":\"uint16[]\"},{\"name\":\"validFroms\",\"type\":\"uint32[]\"},{\"name\":\"validUntils\",\"type\":\"uint32[]\"},{\"name\":\"buyAmounts\",\"type\":\"uint128[]\"},{\"name\":\"sellAmounts\",\"type\":\"uint128[]\"}],\"name\":\"replaceOrders\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPendingWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"batchId\",\"type\":\"uint32\"}],\"name\":\"acceptingSolutions\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_DENOMINATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ENCODED_AUCTION_ELEMENT_WIDTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BATCH_TIME\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBatchId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"offset\",\"type\":\"uint16\"},{\"name\":\"pageSize\",\"type\":\"uint16\"}],\"name\":\"getEncodedUserOrdersPaginated\",\"outputs\":[{\"name\":\"elements\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"tokenAddressToIdMap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"batchId\",\"type\":\"uint32\"}],\"name\":\"requestFutureWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"hasValidWithdrawRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_TOKENS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_TOUCHED_ORDERS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentObjectiveValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"maxTokens\",\"type\":\"uint256\"},{\"name\":\"_feeToken\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint16\"},{\"indexed\":true,\"name\":\"buyToken\",\"type\":\"uint16\"},{\"indexed\":true,\"name\":\"sellToken\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"validFrom\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"validUntil\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"priceNumerator\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"priceDenominator\",\"type\":\"uint128\"}],\"name\":\"OrderPlacement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint16\"}],\"name\":\"TokenListing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint16\"}],\"name\":\"OrderCancellation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint16\"}],\"name\":\"OrderDeletion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"orderId\",\"type\":\"uint16\"},{\"indexed\":true,\"name\":\"sellToken\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"buyToken\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"executedSellAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"executedBuyAmount\",\"type\":\"uint128\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"orderId\",\"type\":\"uint16\"},{\"indexed\":true,\"name\":\"sellToken\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"buyToken\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"executedSellAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"executedBuyAmount\",\"type\":\"uint128\"}],\"name\":\"TradeReversion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"utility\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"disregardedUtility\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"burntFees\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastAuctionBurntFees\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"prices\",\"type\":\"uint128[]\"},{\"indexed\":false,\"name\":\"tokenIdsForPrice\",\"type\":\"uint16[]\"}],\"name\":\"SolutionSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"batchId\",\"type\":\"uint32\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"batchId\",\"type\":\"uint32\"}],\"name\":\"WithdrawRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"

// Gnosis is an auto generated Go binding around an Ethereum contract.
type Gnosis struct {
	GnosisCaller     // Read-only binding to the contract
	GnosisTransactor // Write-only binding to the contract
	GnosisFilterer   // Log filterer for contract events
}

// GnosisCaller is an auto generated read-only Go binding around an Ethereum contract.
type GnosisCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GnosisTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GnosisFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GnosisSession struct {
	Contract     *Gnosis           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GnosisCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GnosisCallerSession struct {
	Contract *GnosisCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GnosisTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GnosisTransactorSession struct {
	Contract     *GnosisTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GnosisRaw is an auto generated low-level Go binding around an Ethereum contract.
type GnosisRaw struct {
	Contract *Gnosis // Generic contract binding to access the raw methods on
}

// GnosisCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GnosisCallerRaw struct {
	Contract *GnosisCaller // Generic read-only contract binding to access the raw methods on
}

// GnosisTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GnosisTransactorRaw struct {
	Contract *GnosisTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGnosis creates a new instance of Gnosis, bound to a specific deployed contract.
func NewGnosis(address common.Address, backend bind.ContractBackend) (*Gnosis, error) {
	contract, err := bindGnosis(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gnosis{GnosisCaller: GnosisCaller{contract: contract}, GnosisTransactor: GnosisTransactor{contract: contract}, GnosisFilterer: GnosisFilterer{contract: contract}}, nil
}

// NewGnosisCaller creates a new read-only instance of Gnosis, bound to a specific deployed contract.
func NewGnosisCaller(address common.Address, caller bind.ContractCaller) (*GnosisCaller, error) {
	contract, err := bindGnosis(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisCaller{contract: contract}, nil
}

// NewGnosisTransactor creates a new write-only instance of Gnosis, bound to a specific deployed contract.
func NewGnosisTransactor(address common.Address, transactor bind.ContractTransactor) (*GnosisTransactor, error) {
	contract, err := bindGnosis(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisTransactor{contract: contract}, nil
}

// NewGnosisFilterer creates a new log filterer instance of Gnosis, bound to a specific deployed contract.
func NewGnosisFilterer(address common.Address, filterer bind.ContractFilterer) (*GnosisFilterer, error) {
	contract, err := bindGnosis(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GnosisFilterer{contract: contract}, nil
}

// bindGnosis binds a generic wrapper to an already deployed contract.
func bindGnosis(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GnosisABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gnosis *GnosisRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gnosis.Contract.GnosisCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gnosis *GnosisRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gnosis.Contract.GnosisTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gnosis *GnosisRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gnosis.Contract.GnosisTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gnosis *GnosisCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gnosis.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gnosis *GnosisTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gnosis.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gnosis *GnosisTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gnosis.Contract.contract.Transact(opts, method, params...)
}

// AMOUNTMINIMUM is a free data retrieval call binding the contract method 0x61ed16d0.
//
// Solidity: function AMOUNT_MINIMUM() view returns(uint128)
func (_Gnosis *GnosisCaller) AMOUNTMINIMUM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "AMOUNT_MINIMUM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AMOUNTMINIMUM is a free data retrieval call binding the contract method 0x61ed16d0.
//
// Solidity: function AMOUNT_MINIMUM() view returns(uint128)
func (_Gnosis *GnosisSession) AMOUNTMINIMUM() (*big.Int, error) {
	return _Gnosis.Contract.AMOUNTMINIMUM(&_Gnosis.CallOpts)
}

// AMOUNTMINIMUM is a free data retrieval call binding the contract method 0x61ed16d0.
//
// Solidity: function AMOUNT_MINIMUM() view returns(uint128)
func (_Gnosis *GnosisCallerSession) AMOUNTMINIMUM() (*big.Int, error) {
	return _Gnosis.Contract.AMOUNTMINIMUM(&_Gnosis.CallOpts)
}

// BATCHTIME is a free data retrieval call binding the contract method 0xe48c015e.
//
// Solidity: function BATCH_TIME() view returns(uint32)
func (_Gnosis *GnosisCaller) BATCHTIME(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "BATCH_TIME")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BATCHTIME is a free data retrieval call binding the contract method 0xe48c015e.
//
// Solidity: function BATCH_TIME() view returns(uint32)
func (_Gnosis *GnosisSession) BATCHTIME() (uint32, error) {
	return _Gnosis.Contract.BATCHTIME(&_Gnosis.CallOpts)
}

// BATCHTIME is a free data retrieval call binding the contract method 0xe48c015e.
//
// Solidity: function BATCH_TIME() view returns(uint32)
func (_Gnosis *GnosisCallerSession) BATCHTIME() (uint32, error) {
	return _Gnosis.Contract.BATCHTIME(&_Gnosis.CallOpts)
}

// ENCODEDAUCTIONELEMENTWIDTH is a free data retrieval call binding the contract method 0xe1d5f64e.
//
// Solidity: function ENCODED_AUCTION_ELEMENT_WIDTH() view returns(uint128)
func (_Gnosis *GnosisCaller) ENCODEDAUCTIONELEMENTWIDTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "ENCODED_AUCTION_ELEMENT_WIDTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ENCODEDAUCTIONELEMENTWIDTH is a free data retrieval call binding the contract method 0xe1d5f64e.
//
// Solidity: function ENCODED_AUCTION_ELEMENT_WIDTH() view returns(uint128)
func (_Gnosis *GnosisSession) ENCODEDAUCTIONELEMENTWIDTH() (*big.Int, error) {
	return _Gnosis.Contract.ENCODEDAUCTIONELEMENTWIDTH(&_Gnosis.CallOpts)
}

// ENCODEDAUCTIONELEMENTWIDTH is a free data retrieval call binding the contract method 0xe1d5f64e.
//
// Solidity: function ENCODED_AUCTION_ELEMENT_WIDTH() view returns(uint128)
func (_Gnosis *GnosisCallerSession) ENCODEDAUCTIONELEMENTWIDTH() (*big.Int, error) {
	return _Gnosis.Contract.ENCODEDAUCTIONELEMENTWIDTH(&_Gnosis.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint128)
func (_Gnosis *GnosisCaller) FEEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "FEE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint128)
func (_Gnosis *GnosisSession) FEEDENOMINATOR() (*big.Int, error) {
	return _Gnosis.Contract.FEEDENOMINATOR(&_Gnosis.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint128)
func (_Gnosis *GnosisCallerSession) FEEDENOMINATOR() (*big.Int, error) {
	return _Gnosis.Contract.FEEDENOMINATOR(&_Gnosis.CallOpts)
}

// FEEFORLISTINGTOKENINOWL is a free data retrieval call binding the contract method 0x41e383ed.
//
// Solidity: function FEE_FOR_LISTING_TOKEN_IN_OWL() view returns(uint256)
func (_Gnosis *GnosisCaller) FEEFORLISTINGTOKENINOWL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "FEE_FOR_LISTING_TOKEN_IN_OWL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEFORLISTINGTOKENINOWL is a free data retrieval call binding the contract method 0x41e383ed.
//
// Solidity: function FEE_FOR_LISTING_TOKEN_IN_OWL() view returns(uint256)
func (_Gnosis *GnosisSession) FEEFORLISTINGTOKENINOWL() (*big.Int, error) {
	return _Gnosis.Contract.FEEFORLISTINGTOKENINOWL(&_Gnosis.CallOpts)
}

// FEEFORLISTINGTOKENINOWL is a free data retrieval call binding the contract method 0x41e383ed.
//
// Solidity: function FEE_FOR_LISTING_TOKEN_IN_OWL() view returns(uint256)
func (_Gnosis *GnosisCallerSession) FEEFORLISTINGTOKENINOWL() (*big.Int, error) {
	return _Gnosis.Contract.FEEFORLISTINGTOKENINOWL(&_Gnosis.CallOpts)
}

// IMPROVEMENTDENOMINATOR is a free data retrieval call binding the contract method 0x094c7e19.
//
// Solidity: function IMPROVEMENT_DENOMINATOR() view returns(uint256)
func (_Gnosis *GnosisCaller) IMPROVEMENTDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "IMPROVEMENT_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IMPROVEMENTDENOMINATOR is a free data retrieval call binding the contract method 0x094c7e19.
//
// Solidity: function IMPROVEMENT_DENOMINATOR() view returns(uint256)
func (_Gnosis *GnosisSession) IMPROVEMENTDENOMINATOR() (*big.Int, error) {
	return _Gnosis.Contract.IMPROVEMENTDENOMINATOR(&_Gnosis.CallOpts)
}

// IMPROVEMENTDENOMINATOR is a free data retrieval call binding the contract method 0x094c7e19.
//
// Solidity: function IMPROVEMENT_DENOMINATOR() view returns(uint256)
func (_Gnosis *GnosisCallerSession) IMPROVEMENTDENOMINATOR() (*big.Int, error) {
	return _Gnosis.Contract.IMPROVEMENTDENOMINATOR(&_Gnosis.CallOpts)
}

// MAXTOKENS is a free data retrieval call binding the contract method 0xf47c84c5.
//
// Solidity: function MAX_TOKENS() view returns(uint256)
func (_Gnosis *GnosisCaller) MAXTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "MAX_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOKENS is a free data retrieval call binding the contract method 0xf47c84c5.
//
// Solidity: function MAX_TOKENS() view returns(uint256)
func (_Gnosis *GnosisSession) MAXTOKENS() (*big.Int, error) {
	return _Gnosis.Contract.MAXTOKENS(&_Gnosis.CallOpts)
}

// MAXTOKENS is a free data retrieval call binding the contract method 0xf47c84c5.
//
// Solidity: function MAX_TOKENS() view returns(uint256)
func (_Gnosis *GnosisCallerSession) MAXTOKENS() (*big.Int, error) {
	return _Gnosis.Contract.MAXTOKENS(&_Gnosis.CallOpts)
}

// MAXTOUCHEDORDERS is a free data retrieval call binding the contract method 0xfb736d32.
//
// Solidity: function MAX_TOUCHED_ORDERS() view returns(uint256)
func (_Gnosis *GnosisCaller) MAXTOUCHEDORDERS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "MAX_TOUCHED_ORDERS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOUCHEDORDERS is a free data retrieval call binding the contract method 0xfb736d32.
//
// Solidity: function MAX_TOUCHED_ORDERS() view returns(uint256)
func (_Gnosis *GnosisSession) MAXTOUCHEDORDERS() (*big.Int, error) {
	return _Gnosis.Contract.MAXTOUCHEDORDERS(&_Gnosis.CallOpts)
}

// MAXTOUCHEDORDERS is a free data retrieval call binding the contract method 0xfb736d32.
//
// Solidity: function MAX_TOUCHED_ORDERS() view returns(uint256)
func (_Gnosis *GnosisCallerSession) MAXTOUCHEDORDERS() (*big.Int, error) {
	return _Gnosis.Contract.MAXTOUCHEDORDERS(&_Gnosis.CallOpts)
}

// UNLIMITEDORDERAMOUNT is a free data retrieval call binding the contract method 0x7fb47b06.
//
// Solidity: function UNLIMITED_ORDER_AMOUNT() view returns(uint128)
func (_Gnosis *GnosisCaller) UNLIMITEDORDERAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "UNLIMITED_ORDER_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNLIMITEDORDERAMOUNT is a free data retrieval call binding the contract method 0x7fb47b06.
//
// Solidity: function UNLIMITED_ORDER_AMOUNT() view returns(uint128)
func (_Gnosis *GnosisSession) UNLIMITEDORDERAMOUNT() (*big.Int, error) {
	return _Gnosis.Contract.UNLIMITEDORDERAMOUNT(&_Gnosis.CallOpts)
}

// UNLIMITEDORDERAMOUNT is a free data retrieval call binding the contract method 0x7fb47b06.
//
// Solidity: function UNLIMITED_ORDER_AMOUNT() view returns(uint128)
func (_Gnosis *GnosisCallerSession) UNLIMITEDORDERAMOUNT() (*big.Int, error) {
	return _Gnosis.Contract.UNLIMITEDORDERAMOUNT(&_Gnosis.CallOpts)
}

// AcceptingSolutions is a free data retrieval call binding the contract method 0xc49598fb.
//
// Solidity: function acceptingSolutions(uint32 batchId) view returns(bool)
func (_Gnosis *GnosisCaller) AcceptingSolutions(opts *bind.CallOpts, batchId uint32) (bool, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "acceptingSolutions", batchId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AcceptingSolutions is a free data retrieval call binding the contract method 0xc49598fb.
//
// Solidity: function acceptingSolutions(uint32 batchId) view returns(bool)
func (_Gnosis *GnosisSession) AcceptingSolutions(batchId uint32) (bool, error) {
	return _Gnosis.Contract.AcceptingSolutions(&_Gnosis.CallOpts, batchId)
}

// AcceptingSolutions is a free data retrieval call binding the contract method 0xc49598fb.
//
// Solidity: function acceptingSolutions(uint32 batchId) view returns(bool)
func (_Gnosis *GnosisCallerSession) AcceptingSolutions(batchId uint32) (bool, error) {
	return _Gnosis.Contract.AcceptingSolutions(&_Gnosis.CallOpts, batchId)
}

// CurrentPrices is a free data retrieval call binding the contract method 0x66367c10.
//
// Solidity: function currentPrices(uint16 ) view returns(uint128)
func (_Gnosis *GnosisCaller) CurrentPrices(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "currentPrices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPrices is a free data retrieval call binding the contract method 0x66367c10.
//
// Solidity: function currentPrices(uint16 ) view returns(uint128)
func (_Gnosis *GnosisSession) CurrentPrices(arg0 uint16) (*big.Int, error) {
	return _Gnosis.Contract.CurrentPrices(&_Gnosis.CallOpts, arg0)
}

// CurrentPrices is a free data retrieval call binding the contract method 0x66367c10.
//
// Solidity: function currentPrices(uint16 ) view returns(uint128)
func (_Gnosis *GnosisCallerSession) CurrentPrices(arg0 uint16) (*big.Int, error) {
	return _Gnosis.Contract.CurrentPrices(&_Gnosis.CallOpts, arg0)
}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_Gnosis *GnosisCaller) FeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "feeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_Gnosis *GnosisSession) FeeToken() (common.Address, error) {
	return _Gnosis.Contract.FeeToken(&_Gnosis.CallOpts)
}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_Gnosis *GnosisCallerSession) FeeToken() (common.Address, error) {
	return _Gnosis.Contract.FeeToken(&_Gnosis.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address user, address token) view returns(uint256)
func (_Gnosis *GnosisCaller) GetBalance(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "getBalance", user, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address user, address token) view returns(uint256)
func (_Gnosis *GnosisSession) GetBalance(user common.Address, token common.Address) (*big.Int, error) {
	return _Gnosis.Contract.GetBalance(&_Gnosis.CallOpts, user, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address user, address token) view returns(uint256)
func (_Gnosis *GnosisCallerSession) GetBalance(user common.Address, token common.Address) (*big.Int, error) {
	return _Gnosis.Contract.GetBalance(&_Gnosis.CallOpts, user, token)
}

// GetCurrentBatchId is a free data retrieval call binding the contract method 0xe720ac8e.
//
// Solidity: function getCurrentBatchId() view returns(uint32)
func (_Gnosis *GnosisCaller) GetCurrentBatchId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "getCurrentBatchId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetCurrentBatchId is a free data retrieval call binding the contract method 0xe720ac8e.
//
// Solidity: function getCurrentBatchId() view returns(uint32)
func (_Gnosis *GnosisSession) GetCurrentBatchId() (uint32, error) {
	return _Gnosis.Contract.GetCurrentBatchId(&_Gnosis.CallOpts)
}

// GetCurrentBatchId is a free data retrieval call binding the contract method 0xe720ac8e.
//
// Solidity: function getCurrentBatchId() view returns(uint32)
func (_Gnosis *GnosisCallerSession) GetCurrentBatchId() (uint32, error) {
	return _Gnosis.Contract.GetCurrentBatchId(&_Gnosis.CallOpts)
}

// GetCurrentObjectiveValue is a free data retrieval call binding the contract method 0xff97c626.
//
// Solidity: function getCurrentObjectiveValue() view returns(uint256)
func (_Gnosis *GnosisCaller) GetCurrentObjectiveValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "getCurrentObjectiveValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentObjectiveValue is a free data retrieval call binding the contract method 0xff97c626.
//
// Solidity: function getCurrentObjectiveValue() view returns(uint256)
func (_Gnosis *GnosisSession) GetCurrentObjectiveValue() (*big.Int, error) {
	return _Gnosis.Contract.GetCurrentObjectiveValue(&_Gnosis.CallOpts)
}

// GetCurrentObjectiveValue is a free data retrieval call binding the contract method 0xff97c626.
//
// Solidity: function getCurrentObjectiveValue() view returns(uint256)
func (_Gnosis *GnosisCallerSession) GetCurrentObjectiveValue() (*big.Int, error) {
	return _Gnosis.Contract.GetCurrentObjectiveValue(&_Gnosis.CallOpts)
}

// GetEncodedOrders is a free data retrieval call binding the contract method 0x23d4a3c9.
//
// Solidity: function getEncodedOrders() view returns(bytes elements)
func (_Gnosis *GnosisCaller) GetEncodedOrders(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "getEncodedOrders")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEncodedOrders is a free data retrieval call binding the contract method 0x23d4a3c9.
//
// Solidity: function getEncodedOrders() view returns(bytes elements)
func (_Gnosis *GnosisSession) GetEncodedOrders() ([]byte, error) {
	return _Gnosis.Contract.GetEncodedOrders(&_Gnosis.CallOpts)
}

// GetEncodedOrders is a free data retrieval call binding the contract method 0x23d4a3c9.
//
// Solidity: function getEncodedOrders() view returns(bytes elements)
func (_Gnosis *GnosisCallerSession) GetEncodedOrders() ([]byte, error) {
	return _Gnosis.Contract.GetEncodedOrders(&_Gnosis.CallOpts)
}

// GetEncodedUserOrders is a free data retrieval call binding the contract method 0x72f3dd39.
//
// Solidity: function getEncodedUserOrders(address user) view returns(bytes elements)
func (_Gnosis *GnosisCaller) GetEncodedUserOrders(opts *bind.CallOpts, user common.Address) ([]byte, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "getEncodedUserOrders", user)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEncodedUserOrders is a free data retrieval call binding the contract method 0x72f3dd39.
//
// Solidity: function getEncodedUserOrders(address user) view returns(bytes elements)
func (_Gnosis *GnosisSession) GetEncodedUserOrders(user common.Address) ([]byte, error) {
	return _Gnosis.Contract.GetEncodedUserOrders(&_Gnosis.CallOpts, user)
}

// GetEncodedUserOrders is a free data retrieval call binding the contract method 0x72f3dd39.
//
// Solidity: function getEncodedUserOrders(address user) view returns(bytes elements)
func (_Gnosis *GnosisCallerSession) GetEncodedUserOrders(user common.Address) ([]byte, error) {
	return _Gnosis.Contract.GetEncodedUserOrders(&_Gnosis.CallOpts, user)
}

// GetEncodedUserOrdersPaginated is a free data retrieval call binding the contract method 0xed2da357.
//
// Solidity: function getEncodedUserOrdersPaginated(address user, uint16 offset, uint16 pageSize) view returns(bytes elements)
func (_Gnosis *GnosisCaller) GetEncodedUserOrdersPaginated(opts *bind.CallOpts, user common.Address, offset uint16, pageSize uint16) ([]byte, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "getEncodedUserOrdersPaginated", user, offset, pageSize)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEncodedUserOrdersPaginated is a free data retrieval call binding the contract method 0xed2da357.
//
// Solidity: function getEncodedUserOrdersPaginated(address user, uint16 offset, uint16 pageSize) view returns(bytes elements)
func (_Gnosis *GnosisSession) GetEncodedUserOrdersPaginated(user common.Address, offset uint16, pageSize uint16) ([]byte, error) {
	return _Gnosis.Contract.GetEncodedUserOrdersPaginated(&_Gnosis.CallOpts, user, offset, pageSize)
}

// GetEncodedUserOrdersPaginated is a free data retrieval call binding the contract method 0xed2da357.
//
// Solidity: function getEncodedUserOrdersPaginated(address user, uint16 offset, uint16 pageSize) view returns(bytes elements)
func (_Gnosis *GnosisCallerSession) GetEncodedUserOrdersPaginated(user common.Address, offset uint16, pageSize uint16) ([]byte, error) {
	return _Gnosis.Contract.GetEncodedUserOrdersPaginated(&_Gnosis.CallOpts, user, offset, pageSize)
}

// GetEncodedUsersPaginated is a free data retrieval call binding the contract method 0x95466a46.
//
// Solidity: function getEncodedUsersPaginated(address previousPageUser, uint16 previousPageUserOffset, uint16 pageSize) view returns(bytes elements)
func (_Gnosis *GnosisCaller) GetEncodedUsersPaginated(opts *bind.CallOpts, previousPageUser common.Address, previousPageUserOffset uint16, pageSize uint16) ([]byte, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "getEncodedUsersPaginated", previousPageUser, previousPageUserOffset, pageSize)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEncodedUsersPaginated is a free data retrieval call binding the contract method 0x95466a46.
//
// Solidity: function getEncodedUsersPaginated(address previousPageUser, uint16 previousPageUserOffset, uint16 pageSize) view returns(bytes elements)
func (_Gnosis *GnosisSession) GetEncodedUsersPaginated(previousPageUser common.Address, previousPageUserOffset uint16, pageSize uint16) ([]byte, error) {
	return _Gnosis.Contract.GetEncodedUsersPaginated(&_Gnosis.CallOpts, previousPageUser, previousPageUserOffset, pageSize)
}

// GetEncodedUsersPaginated is a free data retrieval call binding the contract method 0x95466a46.
//
// Solidity: function getEncodedUsersPaginated(address previousPageUser, uint16 previousPageUserOffset, uint16 pageSize) view returns(bytes elements)
func (_Gnosis *GnosisCallerSession) GetEncodedUsersPaginated(previousPageUser common.Address, previousPageUserOffset uint16, pageSize uint16) ([]byte, error) {
	return _Gnosis.Contract.GetEncodedUsersPaginated(&_Gnosis.CallOpts, previousPageUser, previousPageUserOffset, pageSize)
}

// GetPendingDeposit is a free data retrieval call binding the contract method 0xb3c0afa1.
//
// Solidity: function getPendingDeposit(address user, address token) view returns(uint256, uint32)
func (_Gnosis *GnosisCaller) GetPendingDeposit(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, uint32, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "getPendingDeposit", user, token)

	if err != nil {
		return *new(*big.Int), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// GetPendingDeposit is a free data retrieval call binding the contract method 0xb3c0afa1.
//
// Solidity: function getPendingDeposit(address user, address token) view returns(uint256, uint32)
func (_Gnosis *GnosisSession) GetPendingDeposit(user common.Address, token common.Address) (*big.Int, uint32, error) {
	return _Gnosis.Contract.GetPendingDeposit(&_Gnosis.CallOpts, user, token)
}

// GetPendingDeposit is a free data retrieval call binding the contract method 0xb3c0afa1.
//
// Solidity: function getPendingDeposit(address user, address token) view returns(uint256, uint32)
func (_Gnosis *GnosisCallerSession) GetPendingDeposit(user common.Address, token common.Address) (*big.Int, uint32, error) {
	return _Gnosis.Contract.GetPendingDeposit(&_Gnosis.CallOpts, user, token)
}

// GetPendingWithdraw is a free data retrieval call binding the contract method 0xc33eb9f6.
//
// Solidity: function getPendingWithdraw(address user, address token) view returns(uint256, uint32)
func (_Gnosis *GnosisCaller) GetPendingWithdraw(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, uint32, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "getPendingWithdraw", user, token)

	if err != nil {
		return *new(*big.Int), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// GetPendingWithdraw is a free data retrieval call binding the contract method 0xc33eb9f6.
//
// Solidity: function getPendingWithdraw(address user, address token) view returns(uint256, uint32)
func (_Gnosis *GnosisSession) GetPendingWithdraw(user common.Address, token common.Address) (*big.Int, uint32, error) {
	return _Gnosis.Contract.GetPendingWithdraw(&_Gnosis.CallOpts, user, token)
}

// GetPendingWithdraw is a free data retrieval call binding the contract method 0xc33eb9f6.
//
// Solidity: function getPendingWithdraw(address user, address token) view returns(uint256, uint32)
func (_Gnosis *GnosisCallerSession) GetPendingWithdraw(user common.Address, token common.Address) (*big.Int, uint32, error) {
	return _Gnosis.Contract.GetPendingWithdraw(&_Gnosis.CallOpts, user, token)
}

// GetSecondsRemainingInBatch is a free data retrieval call binding the contract method 0x17569c1d.
//
// Solidity: function getSecondsRemainingInBatch() view returns(uint256)
func (_Gnosis *GnosisCaller) GetSecondsRemainingInBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "getSecondsRemainingInBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSecondsRemainingInBatch is a free data retrieval call binding the contract method 0x17569c1d.
//
// Solidity: function getSecondsRemainingInBatch() view returns(uint256)
func (_Gnosis *GnosisSession) GetSecondsRemainingInBatch() (*big.Int, error) {
	return _Gnosis.Contract.GetSecondsRemainingInBatch(&_Gnosis.CallOpts)
}

// GetSecondsRemainingInBatch is a free data retrieval call binding the contract method 0x17569c1d.
//
// Solidity: function getSecondsRemainingInBatch() view returns(uint256)
func (_Gnosis *GnosisCallerSession) GetSecondsRemainingInBatch() (*big.Int, error) {
	return _Gnosis.Contract.GetSecondsRemainingInBatch(&_Gnosis.CallOpts)
}

// GetUsersPaginated is a free data retrieval call binding the contract method 0x43383ac3.
//
// Solidity: function getUsersPaginated(address previousPageUser, uint16 pageSize) view returns(bytes users)
func (_Gnosis *GnosisCaller) GetUsersPaginated(opts *bind.CallOpts, previousPageUser common.Address, pageSize uint16) ([]byte, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "getUsersPaginated", previousPageUser, pageSize)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetUsersPaginated is a free data retrieval call binding the contract method 0x43383ac3.
//
// Solidity: function getUsersPaginated(address previousPageUser, uint16 pageSize) view returns(bytes users)
func (_Gnosis *GnosisSession) GetUsersPaginated(previousPageUser common.Address, pageSize uint16) ([]byte, error) {
	return _Gnosis.Contract.GetUsersPaginated(&_Gnosis.CallOpts, previousPageUser, pageSize)
}

// GetUsersPaginated is a free data retrieval call binding the contract method 0x43383ac3.
//
// Solidity: function getUsersPaginated(address previousPageUser, uint16 pageSize) view returns(bytes users)
func (_Gnosis *GnosisCallerSession) GetUsersPaginated(previousPageUser common.Address, pageSize uint16) ([]byte, error) {
	return _Gnosis.Contract.GetUsersPaginated(&_Gnosis.CallOpts, previousPageUser, pageSize)
}

// HasToken is a free data retrieval call binding the contract method 0x9bb0f599.
//
// Solidity: function hasToken(address addr) view returns(bool)
func (_Gnosis *GnosisCaller) HasToken(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "hasToken", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasToken is a free data retrieval call binding the contract method 0x9bb0f599.
//
// Solidity: function hasToken(address addr) view returns(bool)
func (_Gnosis *GnosisSession) HasToken(addr common.Address) (bool, error) {
	return _Gnosis.Contract.HasToken(&_Gnosis.CallOpts, addr)
}

// HasToken is a free data retrieval call binding the contract method 0x9bb0f599.
//
// Solidity: function hasToken(address addr) view returns(bool)
func (_Gnosis *GnosisCallerSession) HasToken(addr common.Address) (bool, error) {
	return _Gnosis.Contract.HasToken(&_Gnosis.CallOpts, addr)
}

// HasValidWithdrawRequest is a free data retrieval call binding the contract method 0xf3f47982.
//
// Solidity: function hasValidWithdrawRequest(address user, address token) view returns(bool)
func (_Gnosis *GnosisCaller) HasValidWithdrawRequest(opts *bind.CallOpts, user common.Address, token common.Address) (bool, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "hasValidWithdrawRequest", user, token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasValidWithdrawRequest is a free data retrieval call binding the contract method 0xf3f47982.
//
// Solidity: function hasValidWithdrawRequest(address user, address token) view returns(bool)
func (_Gnosis *GnosisSession) HasValidWithdrawRequest(user common.Address, token common.Address) (bool, error) {
	return _Gnosis.Contract.HasValidWithdrawRequest(&_Gnosis.CallOpts, user, token)
}

// HasValidWithdrawRequest is a free data retrieval call binding the contract method 0xf3f47982.
//
// Solidity: function hasValidWithdrawRequest(address user, address token) view returns(bool)
func (_Gnosis *GnosisCallerSession) HasValidWithdrawRequest(user common.Address, token common.Address) (bool, error) {
	return _Gnosis.Contract.HasValidWithdrawRequest(&_Gnosis.CallOpts, user, token)
}

// LastCreditBatchId is a free data retrieval call binding the contract method 0x907767c0.
//
// Solidity: function lastCreditBatchId(address , address ) view returns(uint32)
func (_Gnosis *GnosisCaller) LastCreditBatchId(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (uint32, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "lastCreditBatchId", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastCreditBatchId is a free data retrieval call binding the contract method 0x907767c0.
//
// Solidity: function lastCreditBatchId(address , address ) view returns(uint32)
func (_Gnosis *GnosisSession) LastCreditBatchId(arg0 common.Address, arg1 common.Address) (uint32, error) {
	return _Gnosis.Contract.LastCreditBatchId(&_Gnosis.CallOpts, arg0, arg1)
}

// LastCreditBatchId is a free data retrieval call binding the contract method 0x907767c0.
//
// Solidity: function lastCreditBatchId(address , address ) view returns(uint32)
func (_Gnosis *GnosisCallerSession) LastCreditBatchId(arg0 common.Address, arg1 common.Address) (uint32, error) {
	return _Gnosis.Contract.LastCreditBatchId(&_Gnosis.CallOpts, arg0, arg1)
}

// LatestSolution is a free data retrieval call binding the contract method 0x9cc84ed3.
//
// Solidity: function latestSolution() view returns(uint32 batchId, address solutionSubmitter, uint256 feeReward, uint256 objectiveValue)
func (_Gnosis *GnosisCaller) LatestSolution(opts *bind.CallOpts) (struct {
	BatchId           uint32
	SolutionSubmitter common.Address
	FeeReward         *big.Int
	ObjectiveValue    *big.Int
}, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "latestSolution")

	outstruct := new(struct {
		BatchId           uint32
		SolutionSubmitter common.Address
		FeeReward         *big.Int
		ObjectiveValue    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchId = out[0].(uint32)
	outstruct.SolutionSubmitter = out[1].(common.Address)
	outstruct.FeeReward = out[2].(*big.Int)
	outstruct.ObjectiveValue = out[3].(*big.Int)

	return *outstruct, err

}

// LatestSolution is a free data retrieval call binding the contract method 0x9cc84ed3.
//
// Solidity: function latestSolution() view returns(uint32 batchId, address solutionSubmitter, uint256 feeReward, uint256 objectiveValue)
func (_Gnosis *GnosisSession) LatestSolution() (struct {
	BatchId           uint32
	SolutionSubmitter common.Address
	FeeReward         *big.Int
	ObjectiveValue    *big.Int
}, error) {
	return _Gnosis.Contract.LatestSolution(&_Gnosis.CallOpts)
}

// LatestSolution is a free data retrieval call binding the contract method 0x9cc84ed3.
//
// Solidity: function latestSolution() view returns(uint32 batchId, address solutionSubmitter, uint256 feeReward, uint256 objectiveValue)
func (_Gnosis *GnosisCallerSession) LatestSolution() (struct {
	BatchId           uint32
	SolutionSubmitter common.Address
	FeeReward         *big.Int
	ObjectiveValue    *big.Int
}, error) {
	return _Gnosis.Contract.LatestSolution(&_Gnosis.CallOpts)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint16)
func (_Gnosis *GnosisCaller) NumTokens(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "numTokens")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint16)
func (_Gnosis *GnosisSession) NumTokens() (uint16, error) {
	return _Gnosis.Contract.NumTokens(&_Gnosis.CallOpts)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint16)
func (_Gnosis *GnosisCallerSession) NumTokens() (uint16, error) {
	return _Gnosis.Contract.NumTokens(&_Gnosis.CallOpts)
}

// Orders is a free data retrieval call binding the contract method 0x793b8c6d.
//
// Solidity: function orders(address , uint256 ) view returns(uint16 buyToken, uint16 sellToken, uint32 validFrom, uint32 validUntil, uint128 priceNumerator, uint128 priceDenominator, uint128 usedAmount)
func (_Gnosis *GnosisCaller) Orders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	BuyToken         uint16
	SellToken        uint16
	ValidFrom        uint32
	ValidUntil       uint32
	PriceNumerator   *big.Int
	PriceDenominator *big.Int
	UsedAmount       *big.Int
}, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "orders", arg0, arg1)

	outstruct := new(struct {
		BuyToken         uint16
		SellToken        uint16
		ValidFrom        uint32
		ValidUntil       uint32
		PriceNumerator   *big.Int
		PriceDenominator *big.Int
		UsedAmount       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BuyToken = out[0].(uint16)
	outstruct.SellToken = out[1].(uint16)
	outstruct.ValidFrom = out[2].(uint32)
	outstruct.ValidUntil = out[3].(uint32)
	outstruct.PriceNumerator = out[4].(*big.Int)
	outstruct.PriceDenominator = out[5].(*big.Int)
	outstruct.UsedAmount = out[6].(*big.Int)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0x793b8c6d.
//
// Solidity: function orders(address , uint256 ) view returns(uint16 buyToken, uint16 sellToken, uint32 validFrom, uint32 validUntil, uint128 priceNumerator, uint128 priceDenominator, uint128 usedAmount)
func (_Gnosis *GnosisSession) Orders(arg0 common.Address, arg1 *big.Int) (struct {
	BuyToken         uint16
	SellToken        uint16
	ValidFrom        uint32
	ValidUntil       uint32
	PriceNumerator   *big.Int
	PriceDenominator *big.Int
	UsedAmount       *big.Int
}, error) {
	return _Gnosis.Contract.Orders(&_Gnosis.CallOpts, arg0, arg1)
}

// Orders is a free data retrieval call binding the contract method 0x793b8c6d.
//
// Solidity: function orders(address , uint256 ) view returns(uint16 buyToken, uint16 sellToken, uint32 validFrom, uint32 validUntil, uint128 priceNumerator, uint128 priceDenominator, uint128 usedAmount)
func (_Gnosis *GnosisCallerSession) Orders(arg0 common.Address, arg1 *big.Int) (struct {
	BuyToken         uint16
	SellToken        uint16
	ValidFrom        uint32
	ValidUntil       uint32
	PriceNumerator   *big.Int
	PriceDenominator *big.Int
	UsedAmount       *big.Int
}, error) {
	return _Gnosis.Contract.Orders(&_Gnosis.CallOpts, arg0, arg1)
}

// TokenAddressToIdMap is a free data retrieval call binding the contract method 0xef574d23.
//
// Solidity: function tokenAddressToIdMap(address addr) view returns(uint16)
func (_Gnosis *GnosisCaller) TokenAddressToIdMap(opts *bind.CallOpts, addr common.Address) (uint16, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "tokenAddressToIdMap", addr)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TokenAddressToIdMap is a free data retrieval call binding the contract method 0xef574d23.
//
// Solidity: function tokenAddressToIdMap(address addr) view returns(uint16)
func (_Gnosis *GnosisSession) TokenAddressToIdMap(addr common.Address) (uint16, error) {
	return _Gnosis.Contract.TokenAddressToIdMap(&_Gnosis.CallOpts, addr)
}

// TokenAddressToIdMap is a free data retrieval call binding the contract method 0xef574d23.
//
// Solidity: function tokenAddressToIdMap(address addr) view returns(uint16)
func (_Gnosis *GnosisCallerSession) TokenAddressToIdMap(addr common.Address) (uint16, error) {
	return _Gnosis.Contract.TokenAddressToIdMap(&_Gnosis.CallOpts, addr)
}

// TokenIdToAddressMap is a free data retrieval call binding the contract method 0x2f10d082.
//
// Solidity: function tokenIdToAddressMap(uint16 id) view returns(address)
func (_Gnosis *GnosisCaller) TokenIdToAddressMap(opts *bind.CallOpts, id uint16) (common.Address, error) {
	var out []interface{}
	err := _Gnosis.contract.Call(opts, &out, "tokenIdToAddressMap", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenIdToAddressMap is a free data retrieval call binding the contract method 0x2f10d082.
//
// Solidity: function tokenIdToAddressMap(uint16 id) view returns(address)
func (_Gnosis *GnosisSession) TokenIdToAddressMap(id uint16) (common.Address, error) {
	return _Gnosis.Contract.TokenIdToAddressMap(&_Gnosis.CallOpts, id)
}

// TokenIdToAddressMap is a free data retrieval call binding the contract method 0x2f10d082.
//
// Solidity: function tokenIdToAddressMap(uint16 id) view returns(address)
func (_Gnosis *GnosisCallerSession) TokenIdToAddressMap(id uint16) (common.Address, error) {
	return _Gnosis.Contract.TokenIdToAddressMap(&_Gnosis.CallOpts, id)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns()
func (_Gnosis *GnosisTransactor) AddToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Gnosis.contract.Transact(opts, "addToken", token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns()
func (_Gnosis *GnosisSession) AddToken(token common.Address) (*types.Transaction, error) {
	return _Gnosis.Contract.AddToken(&_Gnosis.TransactOpts, token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns()
func (_Gnosis *GnosisTransactorSession) AddToken(token common.Address) (*types.Transaction, error) {
	return _Gnosis.Contract.AddToken(&_Gnosis.TransactOpts, token)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x4bdc1b4c.
//
// Solidity: function cancelOrders(uint16[] orderIds) returns()
func (_Gnosis *GnosisTransactor) CancelOrders(opts *bind.TransactOpts, orderIds []uint16) (*types.Transaction, error) {
	return _Gnosis.contract.Transact(opts, "cancelOrders", orderIds)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x4bdc1b4c.
//
// Solidity: function cancelOrders(uint16[] orderIds) returns()
func (_Gnosis *GnosisSession) CancelOrders(orderIds []uint16) (*types.Transaction, error) {
	return _Gnosis.Contract.CancelOrders(&_Gnosis.TransactOpts, orderIds)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x4bdc1b4c.
//
// Solidity: function cancelOrders(uint16[] orderIds) returns()
func (_Gnosis *GnosisTransactorSession) CancelOrders(orderIds []uint16) (*types.Transaction, error) {
	return _Gnosis.Contract.CancelOrders(&_Gnosis.TransactOpts, orderIds)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Gnosis *GnosisTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gnosis.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Gnosis *GnosisSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gnosis.Contract.Deposit(&_Gnosis.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Gnosis *GnosisTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gnosis.Contract.Deposit(&_Gnosis.TransactOpts, token, amount)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x26c3d394.
//
// Solidity: function placeOrder(uint16 buyToken, uint16 sellToken, uint32 validUntil, uint128 buyAmount, uint128 sellAmount) returns(uint256)
func (_Gnosis *GnosisTransactor) PlaceOrder(opts *bind.TransactOpts, buyToken uint16, sellToken uint16, validUntil uint32, buyAmount *big.Int, sellAmount *big.Int) (*types.Transaction, error) {
	return _Gnosis.contract.Transact(opts, "placeOrder", buyToken, sellToken, validUntil, buyAmount, sellAmount)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x26c3d394.
//
// Solidity: function placeOrder(uint16 buyToken, uint16 sellToken, uint32 validUntil, uint128 buyAmount, uint128 sellAmount) returns(uint256)
func (_Gnosis *GnosisSession) PlaceOrder(buyToken uint16, sellToken uint16, validUntil uint32, buyAmount *big.Int, sellAmount *big.Int) (*types.Transaction, error) {
	return _Gnosis.Contract.PlaceOrder(&_Gnosis.TransactOpts, buyToken, sellToken, validUntil, buyAmount, sellAmount)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x26c3d394.
//
// Solidity: function placeOrder(uint16 buyToken, uint16 sellToken, uint32 validUntil, uint128 buyAmount, uint128 sellAmount) returns(uint256)
func (_Gnosis *GnosisTransactorSession) PlaceOrder(buyToken uint16, sellToken uint16, validUntil uint32, buyAmount *big.Int, sellAmount *big.Int) (*types.Transaction, error) {
	return _Gnosis.Contract.PlaceOrder(&_Gnosis.TransactOpts, buyToken, sellToken, validUntil, buyAmount, sellAmount)
}

// PlaceValidFromOrders is a paid mutator transaction binding the contract method 0x65cc3e78.
//
// Solidity: function placeValidFromOrders(uint16[] buyTokens, uint16[] sellTokens, uint32[] validFroms, uint32[] validUntils, uint128[] buyAmounts, uint128[] sellAmounts) returns(uint16[] orderIds)
func (_Gnosis *GnosisTransactor) PlaceValidFromOrders(opts *bind.TransactOpts, buyTokens []uint16, sellTokens []uint16, validFroms []uint32, validUntils []uint32, buyAmounts []*big.Int, sellAmounts []*big.Int) (*types.Transaction, error) {
	return _Gnosis.contract.Transact(opts, "placeValidFromOrders", buyTokens, sellTokens, validFroms, validUntils, buyAmounts, sellAmounts)
}

// PlaceValidFromOrders is a paid mutator transaction binding the contract method 0x65cc3e78.
//
// Solidity: function placeValidFromOrders(uint16[] buyTokens, uint16[] sellTokens, uint32[] validFroms, uint32[] validUntils, uint128[] buyAmounts, uint128[] sellAmounts) returns(uint16[] orderIds)
func (_Gnosis *GnosisSession) PlaceValidFromOrders(buyTokens []uint16, sellTokens []uint16, validFroms []uint32, validUntils []uint32, buyAmounts []*big.Int, sellAmounts []*big.Int) (*types.Transaction, error) {
	return _Gnosis.Contract.PlaceValidFromOrders(&_Gnosis.TransactOpts, buyTokens, sellTokens, validFroms, validUntils, buyAmounts, sellAmounts)
}

// PlaceValidFromOrders is a paid mutator transaction binding the contract method 0x65cc3e78.
//
// Solidity: function placeValidFromOrders(uint16[] buyTokens, uint16[] sellTokens, uint32[] validFroms, uint32[] validUntils, uint128[] buyAmounts, uint128[] sellAmounts) returns(uint16[] orderIds)
func (_Gnosis *GnosisTransactorSession) PlaceValidFromOrders(buyTokens []uint16, sellTokens []uint16, validFroms []uint32, validUntils []uint32, buyAmounts []*big.Int, sellAmounts []*big.Int) (*types.Transaction, error) {
	return _Gnosis.Contract.PlaceValidFromOrders(&_Gnosis.TransactOpts, buyTokens, sellTokens, validFroms, validUntils, buyAmounts, sellAmounts)
}

// ReplaceOrders is a paid mutator transaction binding the contract method 0xc1ef2838.
//
// Solidity: function replaceOrders(uint16[] cancellations, uint16[] buyTokens, uint16[] sellTokens, uint32[] validFroms, uint32[] validUntils, uint128[] buyAmounts, uint128[] sellAmounts) returns(uint16[])
func (_Gnosis *GnosisTransactor) ReplaceOrders(opts *bind.TransactOpts, cancellations []uint16, buyTokens []uint16, sellTokens []uint16, validFroms []uint32, validUntils []uint32, buyAmounts []*big.Int, sellAmounts []*big.Int) (*types.Transaction, error) {
	return _Gnosis.contract.Transact(opts, "replaceOrders", cancellations, buyTokens, sellTokens, validFroms, validUntils, buyAmounts, sellAmounts)
}

// ReplaceOrders is a paid mutator transaction binding the contract method 0xc1ef2838.
//
// Solidity: function replaceOrders(uint16[] cancellations, uint16[] buyTokens, uint16[] sellTokens, uint32[] validFroms, uint32[] validUntils, uint128[] buyAmounts, uint128[] sellAmounts) returns(uint16[])
func (_Gnosis *GnosisSession) ReplaceOrders(cancellations []uint16, buyTokens []uint16, sellTokens []uint16, validFroms []uint32, validUntils []uint32, buyAmounts []*big.Int, sellAmounts []*big.Int) (*types.Transaction, error) {
	return _Gnosis.Contract.ReplaceOrders(&_Gnosis.TransactOpts, cancellations, buyTokens, sellTokens, validFroms, validUntils, buyAmounts, sellAmounts)
}

// ReplaceOrders is a paid mutator transaction binding the contract method 0xc1ef2838.
//
// Solidity: function replaceOrders(uint16[] cancellations, uint16[] buyTokens, uint16[] sellTokens, uint32[] validFroms, uint32[] validUntils, uint128[] buyAmounts, uint128[] sellAmounts) returns(uint16[])
func (_Gnosis *GnosisTransactorSession) ReplaceOrders(cancellations []uint16, buyTokens []uint16, sellTokens []uint16, validFroms []uint32, validUntils []uint32, buyAmounts []*big.Int, sellAmounts []*big.Int) (*types.Transaction, error) {
	return _Gnosis.Contract.ReplaceOrders(&_Gnosis.TransactOpts, cancellations, buyTokens, sellTokens, validFroms, validUntils, buyAmounts, sellAmounts)
}

// RequestFutureWithdraw is a paid mutator transaction binding the contract method 0xf36b6355.
//
// Solidity: function requestFutureWithdraw(address token, uint256 amount, uint32 batchId) returns()
func (_Gnosis *GnosisTransactor) RequestFutureWithdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int, batchId uint32) (*types.Transaction, error) {
	return _Gnosis.contract.Transact(opts, "requestFutureWithdraw", token, amount, batchId)
}

// RequestFutureWithdraw is a paid mutator transaction binding the contract method 0xf36b6355.
//
// Solidity: function requestFutureWithdraw(address token, uint256 amount, uint32 batchId) returns()
func (_Gnosis *GnosisSession) RequestFutureWithdraw(token common.Address, amount *big.Int, batchId uint32) (*types.Transaction, error) {
	return _Gnosis.Contract.RequestFutureWithdraw(&_Gnosis.TransactOpts, token, amount, batchId)
}

// RequestFutureWithdraw is a paid mutator transaction binding the contract method 0xf36b6355.
//
// Solidity: function requestFutureWithdraw(address token, uint256 amount, uint32 batchId) returns()
func (_Gnosis *GnosisTransactorSession) RequestFutureWithdraw(token common.Address, amount *big.Int, batchId uint32) (*types.Transaction, error) {
	return _Gnosis.Contract.RequestFutureWithdraw(&_Gnosis.TransactOpts, token, amount, batchId)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x397a1b28.
//
// Solidity: function requestWithdraw(address token, uint256 amount) returns()
func (_Gnosis *GnosisTransactor) RequestWithdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gnosis.contract.Transact(opts, "requestWithdraw", token, amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x397a1b28.
//
// Solidity: function requestWithdraw(address token, uint256 amount) returns()
func (_Gnosis *GnosisSession) RequestWithdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gnosis.Contract.RequestWithdraw(&_Gnosis.TransactOpts, token, amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x397a1b28.
//
// Solidity: function requestWithdraw(address token, uint256 amount) returns()
func (_Gnosis *GnosisTransactorSession) RequestWithdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gnosis.Contract.RequestWithdraw(&_Gnosis.TransactOpts, token, amount)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0x2e4c83bd.
//
// Solidity: function submitSolution(uint32 batchId, uint256 claimedObjectiveValue, address[] owners, uint16[] orderIds, uint128[] buyVolumes, uint128[] prices, uint16[] tokenIdsForPrice) returns(uint256)
func (_Gnosis *GnosisTransactor) SubmitSolution(opts *bind.TransactOpts, batchId uint32, claimedObjectiveValue *big.Int, owners []common.Address, orderIds []uint16, buyVolumes []*big.Int, prices []*big.Int, tokenIdsForPrice []uint16) (*types.Transaction, error) {
	return _Gnosis.contract.Transact(opts, "submitSolution", batchId, claimedObjectiveValue, owners, orderIds, buyVolumes, prices, tokenIdsForPrice)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0x2e4c83bd.
//
// Solidity: function submitSolution(uint32 batchId, uint256 claimedObjectiveValue, address[] owners, uint16[] orderIds, uint128[] buyVolumes, uint128[] prices, uint16[] tokenIdsForPrice) returns(uint256)
func (_Gnosis *GnosisSession) SubmitSolution(batchId uint32, claimedObjectiveValue *big.Int, owners []common.Address, orderIds []uint16, buyVolumes []*big.Int, prices []*big.Int, tokenIdsForPrice []uint16) (*types.Transaction, error) {
	return _Gnosis.Contract.SubmitSolution(&_Gnosis.TransactOpts, batchId, claimedObjectiveValue, owners, orderIds, buyVolumes, prices, tokenIdsForPrice)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0x2e4c83bd.
//
// Solidity: function submitSolution(uint32 batchId, uint256 claimedObjectiveValue, address[] owners, uint16[] orderIds, uint128[] buyVolumes, uint128[] prices, uint16[] tokenIdsForPrice) returns(uint256)
func (_Gnosis *GnosisTransactorSession) SubmitSolution(batchId uint32, claimedObjectiveValue *big.Int, owners []common.Address, orderIds []uint16, buyVolumes []*big.Int, prices []*big.Int, tokenIdsForPrice []uint16) (*types.Transaction, error) {
	return _Gnosis.Contract.SubmitSolution(&_Gnosis.TransactOpts, batchId, claimedObjectiveValue, owners, orderIds, buyVolumes, prices, tokenIdsForPrice)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address user, address token) returns()
func (_Gnosis *GnosisTransactor) Withdraw(opts *bind.TransactOpts, user common.Address, token common.Address) (*types.Transaction, error) {
	return _Gnosis.contract.Transact(opts, "withdraw", user, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address user, address token) returns()
func (_Gnosis *GnosisSession) Withdraw(user common.Address, token common.Address) (*types.Transaction, error) {
	return _Gnosis.Contract.Withdraw(&_Gnosis.TransactOpts, user, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address user, address token) returns()
func (_Gnosis *GnosisTransactorSession) Withdraw(user common.Address, token common.Address) (*types.Transaction, error) {
	return _Gnosis.Contract.Withdraw(&_Gnosis.TransactOpts, user, token)
}

// GnosisDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Gnosis contract.
type GnosisDepositIterator struct {
	Event *GnosisDeposit // Event containing the contract specifics and raw log

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
func (it *GnosisDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisDeposit)
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
		it.Event = new(GnosisDeposit)
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
func (it *GnosisDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisDeposit represents a Deposit event raised by the Gnosis contract.
type GnosisDeposit struct {
	User    common.Address
	Token   common.Address
	Amount  *big.Int
	BatchId uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xc11cc34e93c67a93382b99f2498e9937198798f3c1c2888008ffc0eeb82f68c4.
//
// Solidity: event Deposit(address indexed user, address indexed token, uint256 amount, uint32 batchId)
func (_Gnosis *GnosisFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*GnosisDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Gnosis.contract.FilterLogs(opts, "Deposit", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GnosisDepositIterator{contract: _Gnosis.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xc11cc34e93c67a93382b99f2498e9937198798f3c1c2888008ffc0eeb82f68c4.
//
// Solidity: event Deposit(address indexed user, address indexed token, uint256 amount, uint32 batchId)
func (_Gnosis *GnosisFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *GnosisDeposit, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Gnosis.contract.WatchLogs(opts, "Deposit", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisDeposit)
				if err := _Gnosis.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xc11cc34e93c67a93382b99f2498e9937198798f3c1c2888008ffc0eeb82f68c4.
//
// Solidity: event Deposit(address indexed user, address indexed token, uint256 amount, uint32 batchId)
func (_Gnosis *GnosisFilterer) ParseDeposit(log types.Log) (*GnosisDeposit, error) {
	event := new(GnosisDeposit)
	if err := _Gnosis.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisOrderCancellationIterator is returned from FilterOrderCancellation and is used to iterate over the raw logs and unpacked data for OrderCancellation events raised by the Gnosis contract.
type GnosisOrderCancellationIterator struct {
	Event *GnosisOrderCancellation // Event containing the contract specifics and raw log

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
func (it *GnosisOrderCancellationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisOrderCancellation)
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
		it.Event = new(GnosisOrderCancellation)
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
func (it *GnosisOrderCancellationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisOrderCancellationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisOrderCancellation represents a OrderCancellation event raised by the Gnosis contract.
type GnosisOrderCancellation struct {
	Owner common.Address
	Id    uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOrderCancellation is a free log retrieval operation binding the contract event 0x7a02963a37046835196f1a3185a036fd67cfca72283e46e4b3cdb99939851937.
//
// Solidity: event OrderCancellation(address indexed owner, uint16 id)
func (_Gnosis *GnosisFilterer) FilterOrderCancellation(opts *bind.FilterOpts, owner []common.Address) (*GnosisOrderCancellationIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Gnosis.contract.FilterLogs(opts, "OrderCancellation", ownerRule)
	if err != nil {
		return nil, err
	}
	return &GnosisOrderCancellationIterator{contract: _Gnosis.contract, event: "OrderCancellation", logs: logs, sub: sub}, nil
}

// WatchOrderCancellation is a free log subscription operation binding the contract event 0x7a02963a37046835196f1a3185a036fd67cfca72283e46e4b3cdb99939851937.
//
// Solidity: event OrderCancellation(address indexed owner, uint16 id)
func (_Gnosis *GnosisFilterer) WatchOrderCancellation(opts *bind.WatchOpts, sink chan<- *GnosisOrderCancellation, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Gnosis.contract.WatchLogs(opts, "OrderCancellation", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisOrderCancellation)
				if err := _Gnosis.contract.UnpackLog(event, "OrderCancellation", log); err != nil {
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

// ParseOrderCancellation is a log parse operation binding the contract event 0x7a02963a37046835196f1a3185a036fd67cfca72283e46e4b3cdb99939851937.
//
// Solidity: event OrderCancellation(address indexed owner, uint16 id)
func (_Gnosis *GnosisFilterer) ParseOrderCancellation(log types.Log) (*GnosisOrderCancellation, error) {
	event := new(GnosisOrderCancellation)
	if err := _Gnosis.contract.UnpackLog(event, "OrderCancellation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisOrderDeletionIterator is returned from FilterOrderDeletion and is used to iterate over the raw logs and unpacked data for OrderDeletion events raised by the Gnosis contract.
type GnosisOrderDeletionIterator struct {
	Event *GnosisOrderDeletion // Event containing the contract specifics and raw log

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
func (it *GnosisOrderDeletionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisOrderDeletion)
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
		it.Event = new(GnosisOrderDeletion)
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
func (it *GnosisOrderDeletionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisOrderDeletionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisOrderDeletion represents a OrderDeletion event raised by the Gnosis contract.
type GnosisOrderDeletion struct {
	Owner common.Address
	Id    uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOrderDeletion is a free log retrieval operation binding the contract event 0x7b0a9854603fbbe7606a58b70d113bd0d1ec8475f1b8cc9603c2d377e67835cd.
//
// Solidity: event OrderDeletion(address indexed owner, uint16 id)
func (_Gnosis *GnosisFilterer) FilterOrderDeletion(opts *bind.FilterOpts, owner []common.Address) (*GnosisOrderDeletionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Gnosis.contract.FilterLogs(opts, "OrderDeletion", ownerRule)
	if err != nil {
		return nil, err
	}
	return &GnosisOrderDeletionIterator{contract: _Gnosis.contract, event: "OrderDeletion", logs: logs, sub: sub}, nil
}

// WatchOrderDeletion is a free log subscription operation binding the contract event 0x7b0a9854603fbbe7606a58b70d113bd0d1ec8475f1b8cc9603c2d377e67835cd.
//
// Solidity: event OrderDeletion(address indexed owner, uint16 id)
func (_Gnosis *GnosisFilterer) WatchOrderDeletion(opts *bind.WatchOpts, sink chan<- *GnosisOrderDeletion, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Gnosis.contract.WatchLogs(opts, "OrderDeletion", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisOrderDeletion)
				if err := _Gnosis.contract.UnpackLog(event, "OrderDeletion", log); err != nil {
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

// ParseOrderDeletion is a log parse operation binding the contract event 0x7b0a9854603fbbe7606a58b70d113bd0d1ec8475f1b8cc9603c2d377e67835cd.
//
// Solidity: event OrderDeletion(address indexed owner, uint16 id)
func (_Gnosis *GnosisFilterer) ParseOrderDeletion(log types.Log) (*GnosisOrderDeletion, error) {
	event := new(GnosisOrderDeletion)
	if err := _Gnosis.contract.UnpackLog(event, "OrderDeletion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisOrderPlacementIterator is returned from FilterOrderPlacement and is used to iterate over the raw logs and unpacked data for OrderPlacement events raised by the Gnosis contract.
type GnosisOrderPlacementIterator struct {
	Event *GnosisOrderPlacement // Event containing the contract specifics and raw log

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
func (it *GnosisOrderPlacementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisOrderPlacement)
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
		it.Event = new(GnosisOrderPlacement)
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
func (it *GnosisOrderPlacementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisOrderPlacementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisOrderPlacement represents a OrderPlacement event raised by the Gnosis contract.
type GnosisOrderPlacement struct {
	Owner            common.Address
	Index            uint16
	BuyToken         uint16
	SellToken        uint16
	ValidFrom        uint32
	ValidUntil       uint32
	PriceNumerator   *big.Int
	PriceDenominator *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOrderPlacement is a free log retrieval operation binding the contract event 0xdecf6fde8243981299f7b7a776f29a9fc67a2c9848e25d77c50eb11fa58a7e21.
//
// Solidity: event OrderPlacement(address indexed owner, uint16 index, uint16 indexed buyToken, uint16 indexed sellToken, uint32 validFrom, uint32 validUntil, uint128 priceNumerator, uint128 priceDenominator)
func (_Gnosis *GnosisFilterer) FilterOrderPlacement(opts *bind.FilterOpts, owner []common.Address, buyToken []uint16, sellToken []uint16) (*GnosisOrderPlacementIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}
	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}

	logs, sub, err := _Gnosis.contract.FilterLogs(opts, "OrderPlacement", ownerRule, buyTokenRule, sellTokenRule)
	if err != nil {
		return nil, err
	}
	return &GnosisOrderPlacementIterator{contract: _Gnosis.contract, event: "OrderPlacement", logs: logs, sub: sub}, nil
}

// WatchOrderPlacement is a free log subscription operation binding the contract event 0xdecf6fde8243981299f7b7a776f29a9fc67a2c9848e25d77c50eb11fa58a7e21.
//
// Solidity: event OrderPlacement(address indexed owner, uint16 index, uint16 indexed buyToken, uint16 indexed sellToken, uint32 validFrom, uint32 validUntil, uint128 priceNumerator, uint128 priceDenominator)
func (_Gnosis *GnosisFilterer) WatchOrderPlacement(opts *bind.WatchOpts, sink chan<- *GnosisOrderPlacement, owner []common.Address, buyToken []uint16, sellToken []uint16) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}
	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}

	logs, sub, err := _Gnosis.contract.WatchLogs(opts, "OrderPlacement", ownerRule, buyTokenRule, sellTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisOrderPlacement)
				if err := _Gnosis.contract.UnpackLog(event, "OrderPlacement", log); err != nil {
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

// ParseOrderPlacement is a log parse operation binding the contract event 0xdecf6fde8243981299f7b7a776f29a9fc67a2c9848e25d77c50eb11fa58a7e21.
//
// Solidity: event OrderPlacement(address indexed owner, uint16 index, uint16 indexed buyToken, uint16 indexed sellToken, uint32 validFrom, uint32 validUntil, uint128 priceNumerator, uint128 priceDenominator)
func (_Gnosis *GnosisFilterer) ParseOrderPlacement(log types.Log) (*GnosisOrderPlacement, error) {
	event := new(GnosisOrderPlacement)
	if err := _Gnosis.contract.UnpackLog(event, "OrderPlacement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSolutionSubmissionIterator is returned from FilterSolutionSubmission and is used to iterate over the raw logs and unpacked data for SolutionSubmission events raised by the Gnosis contract.
type GnosisSolutionSubmissionIterator struct {
	Event *GnosisSolutionSubmission // Event containing the contract specifics and raw log

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
func (it *GnosisSolutionSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSolutionSubmission)
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
		it.Event = new(GnosisSolutionSubmission)
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
func (it *GnosisSolutionSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSolutionSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSolutionSubmission represents a SolutionSubmission event raised by the Gnosis contract.
type GnosisSolutionSubmission struct {
	Submitter            common.Address
	Utility              *big.Int
	DisregardedUtility   *big.Int
	BurntFees            *big.Int
	LastAuctionBurntFees *big.Int
	Prices               []*big.Int
	TokenIdsForPrice     []uint16
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSolutionSubmission is a free log retrieval operation binding the contract event 0x2140b6253bf38aea0a4ac9e9e6427b256e4035d60df4a85bb139ce975eb6b41d.
//
// Solidity: event SolutionSubmission(address indexed submitter, uint256 utility, uint256 disregardedUtility, uint256 burntFees, uint256 lastAuctionBurntFees, uint128[] prices, uint16[] tokenIdsForPrice)
func (_Gnosis *GnosisFilterer) FilterSolutionSubmission(opts *bind.FilterOpts, submitter []common.Address) (*GnosisSolutionSubmissionIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Gnosis.contract.FilterLogs(opts, "SolutionSubmission", submitterRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSolutionSubmissionIterator{contract: _Gnosis.contract, event: "SolutionSubmission", logs: logs, sub: sub}, nil
}

// WatchSolutionSubmission is a free log subscription operation binding the contract event 0x2140b6253bf38aea0a4ac9e9e6427b256e4035d60df4a85bb139ce975eb6b41d.
//
// Solidity: event SolutionSubmission(address indexed submitter, uint256 utility, uint256 disregardedUtility, uint256 burntFees, uint256 lastAuctionBurntFees, uint128[] prices, uint16[] tokenIdsForPrice)
func (_Gnosis *GnosisFilterer) WatchSolutionSubmission(opts *bind.WatchOpts, sink chan<- *GnosisSolutionSubmission, submitter []common.Address) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Gnosis.contract.WatchLogs(opts, "SolutionSubmission", submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSolutionSubmission)
				if err := _Gnosis.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
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

// ParseSolutionSubmission is a log parse operation binding the contract event 0x2140b6253bf38aea0a4ac9e9e6427b256e4035d60df4a85bb139ce975eb6b41d.
//
// Solidity: event SolutionSubmission(address indexed submitter, uint256 utility, uint256 disregardedUtility, uint256 burntFees, uint256 lastAuctionBurntFees, uint128[] prices, uint16[] tokenIdsForPrice)
func (_Gnosis *GnosisFilterer) ParseSolutionSubmission(log types.Log) (*GnosisSolutionSubmission, error) {
	event := new(GnosisSolutionSubmission)
	if err := _Gnosis.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisTokenListingIterator is returned from FilterTokenListing and is used to iterate over the raw logs and unpacked data for TokenListing events raised by the Gnosis contract.
type GnosisTokenListingIterator struct {
	Event *GnosisTokenListing // Event containing the contract specifics and raw log

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
func (it *GnosisTokenListingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisTokenListing)
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
		it.Event = new(GnosisTokenListing)
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
func (it *GnosisTokenListingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisTokenListingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisTokenListing represents a TokenListing event raised by the Gnosis contract.
type GnosisTokenListing struct {
	Token common.Address
	Id    uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenListing is a free log retrieval operation binding the contract event 0xe4b282c4351ffe36572a572de193a7de086edc47c9e62669fe6ab49fc53a3313.
//
// Solidity: event TokenListing(address token, uint16 id)
func (_Gnosis *GnosisFilterer) FilterTokenListing(opts *bind.FilterOpts) (*GnosisTokenListingIterator, error) {

	logs, sub, err := _Gnosis.contract.FilterLogs(opts, "TokenListing")
	if err != nil {
		return nil, err
	}
	return &GnosisTokenListingIterator{contract: _Gnosis.contract, event: "TokenListing", logs: logs, sub: sub}, nil
}

// WatchTokenListing is a free log subscription operation binding the contract event 0xe4b282c4351ffe36572a572de193a7de086edc47c9e62669fe6ab49fc53a3313.
//
// Solidity: event TokenListing(address token, uint16 id)
func (_Gnosis *GnosisFilterer) WatchTokenListing(opts *bind.WatchOpts, sink chan<- *GnosisTokenListing) (event.Subscription, error) {

	logs, sub, err := _Gnosis.contract.WatchLogs(opts, "TokenListing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisTokenListing)
				if err := _Gnosis.contract.UnpackLog(event, "TokenListing", log); err != nil {
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

// ParseTokenListing is a log parse operation binding the contract event 0xe4b282c4351ffe36572a572de193a7de086edc47c9e62669fe6ab49fc53a3313.
//
// Solidity: event TokenListing(address token, uint16 id)
func (_Gnosis *GnosisFilterer) ParseTokenListing(log types.Log) (*GnosisTokenListing, error) {
	event := new(GnosisTokenListing)
	if err := _Gnosis.contract.UnpackLog(event, "TokenListing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the Gnosis contract.
type GnosisTradeIterator struct {
	Event *GnosisTrade // Event containing the contract specifics and raw log

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
func (it *GnosisTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisTrade)
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
		it.Event = new(GnosisTrade)
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
func (it *GnosisTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisTrade represents a Trade event raised by the Gnosis contract.
type GnosisTrade struct {
	Owner              common.Address
	OrderId            uint16
	SellToken          uint16
	BuyToken           uint16
	ExecutedSellAmount *big.Int
	ExecutedBuyAmount  *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0xafa5bc1fb80950b7cb2353ba0cf16a6d68de75801f2dac54b2dae9268450010a.
//
// Solidity: event Trade(address indexed owner, uint16 indexed orderId, uint16 indexed sellToken, uint16 buyToken, uint128 executedSellAmount, uint128 executedBuyAmount)
func (_Gnosis *GnosisFilterer) FilterTrade(opts *bind.FilterOpts, owner []common.Address, orderId []uint16, sellToken []uint16) (*GnosisTradeIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}

	logs, sub, err := _Gnosis.contract.FilterLogs(opts, "Trade", ownerRule, orderIdRule, sellTokenRule)
	if err != nil {
		return nil, err
	}
	return &GnosisTradeIterator{contract: _Gnosis.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0xafa5bc1fb80950b7cb2353ba0cf16a6d68de75801f2dac54b2dae9268450010a.
//
// Solidity: event Trade(address indexed owner, uint16 indexed orderId, uint16 indexed sellToken, uint16 buyToken, uint128 executedSellAmount, uint128 executedBuyAmount)
func (_Gnosis *GnosisFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *GnosisTrade, owner []common.Address, orderId []uint16, sellToken []uint16) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}

	logs, sub, err := _Gnosis.contract.WatchLogs(opts, "Trade", ownerRule, orderIdRule, sellTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisTrade)
				if err := _Gnosis.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0xafa5bc1fb80950b7cb2353ba0cf16a6d68de75801f2dac54b2dae9268450010a.
//
// Solidity: event Trade(address indexed owner, uint16 indexed orderId, uint16 indexed sellToken, uint16 buyToken, uint128 executedSellAmount, uint128 executedBuyAmount)
func (_Gnosis *GnosisFilterer) ParseTrade(log types.Log) (*GnosisTrade, error) {
	event := new(GnosisTrade)
	if err := _Gnosis.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisTradeReversionIterator is returned from FilterTradeReversion and is used to iterate over the raw logs and unpacked data for TradeReversion events raised by the Gnosis contract.
type GnosisTradeReversionIterator struct {
	Event *GnosisTradeReversion // Event containing the contract specifics and raw log

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
func (it *GnosisTradeReversionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisTradeReversion)
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
		it.Event = new(GnosisTradeReversion)
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
func (it *GnosisTradeReversionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisTradeReversionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisTradeReversion represents a TradeReversion event raised by the Gnosis contract.
type GnosisTradeReversion struct {
	Owner              common.Address
	OrderId            uint16
	SellToken          uint16
	BuyToken           uint16
	ExecutedSellAmount *big.Int
	ExecutedBuyAmount  *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTradeReversion is a free log retrieval operation binding the contract event 0xb7214f648cea2a7c47aaea7e7aafef610de8d04366d26f66879d076516964eae.
//
// Solidity: event TradeReversion(address indexed owner, uint16 indexed orderId, uint16 indexed sellToken, uint16 buyToken, uint128 executedSellAmount, uint128 executedBuyAmount)
func (_Gnosis *GnosisFilterer) FilterTradeReversion(opts *bind.FilterOpts, owner []common.Address, orderId []uint16, sellToken []uint16) (*GnosisTradeReversionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}

	logs, sub, err := _Gnosis.contract.FilterLogs(opts, "TradeReversion", ownerRule, orderIdRule, sellTokenRule)
	if err != nil {
		return nil, err
	}
	return &GnosisTradeReversionIterator{contract: _Gnosis.contract, event: "TradeReversion", logs: logs, sub: sub}, nil
}

// WatchTradeReversion is a free log subscription operation binding the contract event 0xb7214f648cea2a7c47aaea7e7aafef610de8d04366d26f66879d076516964eae.
//
// Solidity: event TradeReversion(address indexed owner, uint16 indexed orderId, uint16 indexed sellToken, uint16 buyToken, uint128 executedSellAmount, uint128 executedBuyAmount)
func (_Gnosis *GnosisFilterer) WatchTradeReversion(opts *bind.WatchOpts, sink chan<- *GnosisTradeReversion, owner []common.Address, orderId []uint16, sellToken []uint16) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}

	logs, sub, err := _Gnosis.contract.WatchLogs(opts, "TradeReversion", ownerRule, orderIdRule, sellTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisTradeReversion)
				if err := _Gnosis.contract.UnpackLog(event, "TradeReversion", log); err != nil {
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

// ParseTradeReversion is a log parse operation binding the contract event 0xb7214f648cea2a7c47aaea7e7aafef610de8d04366d26f66879d076516964eae.
//
// Solidity: event TradeReversion(address indexed owner, uint16 indexed orderId, uint16 indexed sellToken, uint16 buyToken, uint128 executedSellAmount, uint128 executedBuyAmount)
func (_Gnosis *GnosisFilterer) ParseTradeReversion(log types.Log) (*GnosisTradeReversion, error) {
	event := new(GnosisTradeReversion)
	if err := _Gnosis.contract.UnpackLog(event, "TradeReversion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Gnosis contract.
type GnosisWithdrawIterator struct {
	Event *GnosisWithdraw // Event containing the contract specifics and raw log

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
func (it *GnosisWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisWithdraw)
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
		it.Event = new(GnosisWithdraw)
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
func (it *GnosisWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisWithdraw represents a Withdraw event raised by the Gnosis contract.
type GnosisWithdraw struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed token, uint256 amount)
func (_Gnosis *GnosisFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*GnosisWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Gnosis.contract.FilterLogs(opts, "Withdraw", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GnosisWithdrawIterator{contract: _Gnosis.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed token, uint256 amount)
func (_Gnosis *GnosisFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *GnosisWithdraw, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Gnosis.contract.WatchLogs(opts, "Withdraw", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisWithdraw)
				if err := _Gnosis.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed token, uint256 amount)
func (_Gnosis *GnosisFilterer) ParseWithdraw(log types.Log) (*GnosisWithdraw, error) {
	event := new(GnosisWithdraw)
	if err := _Gnosis.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisWithdrawRequestIterator is returned from FilterWithdrawRequest and is used to iterate over the raw logs and unpacked data for WithdrawRequest events raised by the Gnosis contract.
type GnosisWithdrawRequestIterator struct {
	Event *GnosisWithdrawRequest // Event containing the contract specifics and raw log

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
func (it *GnosisWithdrawRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisWithdrawRequest)
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
		it.Event = new(GnosisWithdrawRequest)
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
func (it *GnosisWithdrawRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisWithdrawRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisWithdrawRequest represents a WithdrawRequest event raised by the Gnosis contract.
type GnosisWithdrawRequest struct {
	User    common.Address
	Token   common.Address
	Amount  *big.Int
	BatchId uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRequest is a free log retrieval operation binding the contract event 0x2c6245af506f0fc1089918c02c1d01bde9cc807609b334b3e7644d6dfb5a6c5e.
//
// Solidity: event WithdrawRequest(address indexed user, address indexed token, uint256 amount, uint32 batchId)
func (_Gnosis *GnosisFilterer) FilterWithdrawRequest(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*GnosisWithdrawRequestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Gnosis.contract.FilterLogs(opts, "WithdrawRequest", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GnosisWithdrawRequestIterator{contract: _Gnosis.contract, event: "WithdrawRequest", logs: logs, sub: sub}, nil
}

// WatchWithdrawRequest is a free log subscription operation binding the contract event 0x2c6245af506f0fc1089918c02c1d01bde9cc807609b334b3e7644d6dfb5a6c5e.
//
// Solidity: event WithdrawRequest(address indexed user, address indexed token, uint256 amount, uint32 batchId)
func (_Gnosis *GnosisFilterer) WatchWithdrawRequest(opts *bind.WatchOpts, sink chan<- *GnosisWithdrawRequest, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Gnosis.contract.WatchLogs(opts, "WithdrawRequest", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisWithdrawRequest)
				if err := _Gnosis.contract.UnpackLog(event, "WithdrawRequest", log); err != nil {
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

// ParseWithdrawRequest is a log parse operation binding the contract event 0x2c6245af506f0fc1089918c02c1d01bde9cc807609b334b3e7644d6dfb5a6c5e.
//
// Solidity: event WithdrawRequest(address indexed user, address indexed token, uint256 amount, uint32 batchId)
func (_Gnosis *GnosisFilterer) ParseWithdrawRequest(log types.Log) (*GnosisWithdrawRequest, error) {
	event := new(GnosisWithdrawRequest)
	if err := _Gnosis.contract.UnpackLog(event, "WithdrawRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
