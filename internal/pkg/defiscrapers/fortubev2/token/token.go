// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"APR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"APY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"}],\"name\":\"_addReservesFresh\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_borrowSafeRatio\",\"type\":\"uint256\"}],\"name\":\"_setBorrowSafeRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"_setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactor\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactorFresh\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"addExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_addAmount\",\"type\":\"uint256\"}],\"name\":\"addTotalCash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowSafeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"calcBalanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalReserves\",\"type\":\"uint256\"}],\"name\":\"calcExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"striker\",\"type\":\"address\"}],\"name\":\"cancellingOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"strikeOk\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"strikeLog\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"contractIBankController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"divExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"}],\"name\":\"divScalar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"divisor\",\"type\":\"uint256\"}],\"name\":\"divScalarByExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"divisor\",\"type\":\"uint256\"}],\"name\":\"divScalarByExpTruncate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transferInAmout\",\"type\":\"uint256\"}],\"name\":\"exchangeRateAfter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denom\",\"type\":\"uint256\"}],\"name\":\"getDiv\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rational\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denom\",\"type\":\"uint256\"}],\"name\":\"getExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rational\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupplyRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialExchangeRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialInterestRateModel\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bank\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_borrowSafeRatio\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractIInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractFToken\",\"name\":\"fTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"mulExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"name\":\"mulExp3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"}],\"name\":\"mulScalar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"scaled\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"}],\"name\":\"mulScalarTruncate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addend\",\"type\":\"uint256\"}],\"name\":\"mulScalarTruncateAddUInt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peekInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_accrualBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_borrowIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalReserves\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"subExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_subAmount\",\"type\":\"uint256\"}],\"name\":\"subTotalCash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"tokenCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"}],\"name\":\"truncate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawTokensIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmountIn\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// APR is a free data retrieval call binding the contract method 0xbd30558e.
//
// Solidity: function APR() view returns(uint256)
func (_Token *TokenCaller) APR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "APR")
	return *ret0, err
}

// APR is a free data retrieval call binding the contract method 0xbd30558e.
//
// Solidity: function APR() view returns(uint256)
func (_Token *TokenSession) APR() (*big.Int, error) {
	return _Token.Contract.APR(&_Token.CallOpts)
}

// APR is a free data retrieval call binding the contract method 0xbd30558e.
//
// Solidity: function APR() view returns(uint256)
func (_Token *TokenCallerSession) APR() (*big.Int, error) {
	return _Token.Contract.APR(&_Token.CallOpts)
}

// APY is a free data retrieval call binding the contract method 0xef8bd305.
//
// Solidity: function APY() view returns(uint256)
func (_Token *TokenCaller) APY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "APY")
	return *ret0, err
}

// APY is a free data retrieval call binding the contract method 0xef8bd305.
//
// Solidity: function APY() view returns(uint256)
func (_Token *TokenSession) APY() (*big.Int, error) {
	return _Token.Contract.APY(&_Token.CallOpts)
}

// APY is a free data retrieval call binding the contract method 0xef8bd305.
//
// Solidity: function APY() view returns(uint256)
func (_Token *TokenCallerSession) APY() (*big.Int, error) {
	return _Token.Contract.APY(&_Token.CallOpts)
}

// ONE is a free data retrieval call binding the contract method 0xc2ee3a08.
//
// Solidity: function ONE() view returns(uint256)
func (_Token *TokenCaller) ONE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "ONE")
	return *ret0, err
}

// ONE is a free data retrieval call binding the contract method 0xc2ee3a08.
//
// Solidity: function ONE() view returns(uint256)
func (_Token *TokenSession) ONE() (*big.Int, error) {
	return _Token.Contract.ONE(&_Token.CallOpts)
}

// ONE is a free data retrieval call binding the contract method 0xc2ee3a08.
//
// Solidity: function ONE() view returns(uint256)
func (_Token *TokenCallerSession) ONE() (*big.Int, error) {
	return _Token.Contract.ONE(&_Token.CallOpts)
}

// AccountBorrows is a free data retrieval call binding the contract method 0xd40e8f4a.
//
// Solidity: function accountBorrows(address ) view returns(uint256 principal, uint256 interestIndex)
func (_Token *TokenCaller) AccountBorrows(opts *bind.CallOpts, arg0 common.Address) (struct {
	Principal     *big.Int
	InterestIndex *big.Int
}, error) {
	ret := new(struct {
		Principal     *big.Int
		InterestIndex *big.Int
	})
	out := ret
	err := _Token.contract.Call(opts, out, "accountBorrows", arg0)
	return *ret, err
}

// AccountBorrows is a free data retrieval call binding the contract method 0xd40e8f4a.
//
// Solidity: function accountBorrows(address ) view returns(uint256 principal, uint256 interestIndex)
func (_Token *TokenSession) AccountBorrows(arg0 common.Address) (struct {
	Principal     *big.Int
	InterestIndex *big.Int
}, error) {
	return _Token.Contract.AccountBorrows(&_Token.CallOpts, arg0)
}

// AccountBorrows is a free data retrieval call binding the contract method 0xd40e8f4a.
//
// Solidity: function accountBorrows(address ) view returns(uint256 principal, uint256 interestIndex)
func (_Token *TokenCallerSession) AccountBorrows(arg0 common.Address) (struct {
	Principal     *big.Int
	InterestIndex *big.Int
}, error) {
	return _Token.Contract.AccountBorrows(&_Token.CallOpts, arg0)
}

// AccountTokens is a free data retrieval call binding the contract method 0xa19d1460.
//
// Solidity: function accountTokens(address ) view returns(uint256)
func (_Token *TokenCaller) AccountTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "accountTokens", arg0)
	return *ret0, err
}

// AccountTokens is a free data retrieval call binding the contract method 0xa19d1460.
//
// Solidity: function accountTokens(address ) view returns(uint256)
func (_Token *TokenSession) AccountTokens(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.AccountTokens(&_Token.CallOpts, arg0)
}

// AccountTokens is a free data retrieval call binding the contract method 0xa19d1460.
//
// Solidity: function accountTokens(address ) view returns(uint256)
func (_Token *TokenCallerSession) AccountTokens(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.AccountTokens(&_Token.CallOpts, arg0)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Token *TokenCaller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "accrualBlockNumber")
	return *ret0, err
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Token *TokenSession) AccrualBlockNumber() (*big.Int, error) {
	return _Token.Contract.AccrualBlockNumber(&_Token.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Token *TokenCallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _Token.Contract.AccrualBlockNumber(&_Token.CallOpts)
}

// AddExp is a free data retrieval call binding the contract method 0x4147da13.
//
// Solidity: function addExp(uint256 a, uint256 b) pure returns(uint256 result)
func (_Token *TokenCaller) AddExp(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "addExp", a, b)
	return *ret0, err
}

// AddExp is a free data retrieval call binding the contract method 0x4147da13.
//
// Solidity: function addExp(uint256 a, uint256 b) pure returns(uint256 result)
func (_Token *TokenSession) AddExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Token.Contract.AddExp(&_Token.CallOpts, a, b)
}

// AddExp is a free data retrieval call binding the contract method 0x4147da13.
//
// Solidity: function addExp(uint256 a, uint256 b) pure returns(uint256 result)
func (_Token *TokenCallerSession) AddExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Token.Contract.AddExp(&_Token.CallOpts, a, b)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Token *TokenCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Token *TokenSession) Admin() (common.Address, error) {
	return _Token.Contract.Admin(&_Token.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Token *TokenCallerSession) Admin() (common.Address, error) {
	return _Token.Contract.Admin(&_Token.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Token *TokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Token *TokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Token *TokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, owner)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Token *TokenCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "bank")
	return *ret0, err
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Token *TokenSession) Bank() (common.Address, error) {
	return _Token.Contract.Bank(&_Token.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Token *TokenCallerSession) Bank() (common.Address, error) {
	return _Token.Contract.Bank(&_Token.CallOpts)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Token *TokenCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "borrowBalanceStored", account)
	return *ret0, err
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Token *TokenSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Token.Contract.BorrowBalanceStored(&_Token.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Token *TokenCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Token.Contract.BorrowBalanceStored(&_Token.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Token *TokenCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "borrowIndex")
	return *ret0, err
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Token *TokenSession) BorrowIndex() (*big.Int, error) {
	return _Token.Contract.BorrowIndex(&_Token.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Token *TokenCallerSession) BorrowIndex() (*big.Int, error) {
	return _Token.Contract.BorrowIndex(&_Token.CallOpts)
}

// BorrowSafeRatio is a free data retrieval call binding the contract method 0x85cc1486.
//
// Solidity: function borrowSafeRatio() view returns(uint256)
func (_Token *TokenCaller) BorrowSafeRatio(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "borrowSafeRatio")
	return *ret0, err
}

// BorrowSafeRatio is a free data retrieval call binding the contract method 0x85cc1486.
//
// Solidity: function borrowSafeRatio() view returns(uint256)
func (_Token *TokenSession) BorrowSafeRatio() (*big.Int, error) {
	return _Token.Contract.BorrowSafeRatio(&_Token.CallOpts)
}

// BorrowSafeRatio is a free data retrieval call binding the contract method 0x85cc1486.
//
// Solidity: function borrowSafeRatio() view returns(uint256)
func (_Token *TokenCallerSession) BorrowSafeRatio() (*big.Int, error) {
	return _Token.Contract.BorrowSafeRatio(&_Token.CallOpts)
}

// CalcBalanceOfUnderlying is a free data retrieval call binding the contract method 0xe61de305.
//
// Solidity: function calcBalanceOfUnderlying(address owner) view returns(uint256)
func (_Token *TokenCaller) CalcBalanceOfUnderlying(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "calcBalanceOfUnderlying", owner)
	return *ret0, err
}

// CalcBalanceOfUnderlying is a free data retrieval call binding the contract method 0xe61de305.
//
// Solidity: function calcBalanceOfUnderlying(address owner) view returns(uint256)
func (_Token *TokenSession) CalcBalanceOfUnderlying(owner common.Address) (*big.Int, error) {
	return _Token.Contract.CalcBalanceOfUnderlying(&_Token.CallOpts, owner)
}

// CalcBalanceOfUnderlying is a free data retrieval call binding the contract method 0xe61de305.
//
// Solidity: function calcBalanceOfUnderlying(address owner) view returns(uint256)
func (_Token *TokenCallerSession) CalcBalanceOfUnderlying(owner common.Address) (*big.Int, error) {
	return _Token.Contract.CalcBalanceOfUnderlying(&_Token.CallOpts, owner)
}

// CalcExchangeRate is a free data retrieval call binding the contract method 0xc5e9cadc.
//
// Solidity: function calcExchangeRate(uint256 _totalBorrows, uint256 _totalReserves) view returns(uint256 exchangeRate)
func (_Token *TokenCaller) CalcExchangeRate(opts *bind.CallOpts, _totalBorrows *big.Int, _totalReserves *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "calcExchangeRate", _totalBorrows, _totalReserves)
	return *ret0, err
}

// CalcExchangeRate is a free data retrieval call binding the contract method 0xc5e9cadc.
//
// Solidity: function calcExchangeRate(uint256 _totalBorrows, uint256 _totalReserves) view returns(uint256 exchangeRate)
func (_Token *TokenSession) CalcExchangeRate(_totalBorrows *big.Int, _totalReserves *big.Int) (*big.Int, error) {
	return _Token.Contract.CalcExchangeRate(&_Token.CallOpts, _totalBorrows, _totalReserves)
}

// CalcExchangeRate is a free data retrieval call binding the contract method 0xc5e9cadc.
//
// Solidity: function calcExchangeRate(uint256 _totalBorrows, uint256 _totalReserves) view returns(uint256 exchangeRate)
func (_Token *TokenCallerSession) CalcExchangeRate(_totalBorrows *big.Int, _totalReserves *big.Int) (*big.Int, error) {
	return _Token.Contract.CalcExchangeRate(&_Token.CallOpts, _totalBorrows, _totalReserves)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Token *TokenCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "controller")
	return *ret0, err
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Token *TokenSession) Controller() (common.Address, error) {
	return _Token.Contract.Controller(&_Token.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Token *TokenCallerSession) Controller() (common.Address, error) {
	return _Token.Contract.Controller(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenCallerSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// DivExp is a free data retrieval call binding the contract method 0xb507f9f9.
//
// Solidity: function divExp(uint256 a, uint256 b) pure returns(uint256)
func (_Token *TokenCaller) DivExp(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "divExp", a, b)
	return *ret0, err
}

// DivExp is a free data retrieval call binding the contract method 0xb507f9f9.
//
// Solidity: function divExp(uint256 a, uint256 b) pure returns(uint256)
func (_Token *TokenSession) DivExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Token.Contract.DivExp(&_Token.CallOpts, a, b)
}

// DivExp is a free data retrieval call binding the contract method 0xb507f9f9.
//
// Solidity: function divExp(uint256 a, uint256 b) pure returns(uint256)
func (_Token *TokenCallerSession) DivExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Token.Contract.DivExp(&_Token.CallOpts, a, b)
}

// DivScalar is a free data retrieval call binding the contract method 0x4751b79c.
//
// Solidity: function divScalar(uint256 a, uint256 scalar) pure returns(uint256)
func (_Token *TokenCaller) DivScalar(opts *bind.CallOpts, a *big.Int, scalar *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "divScalar", a, scalar)
	return *ret0, err
}

// DivScalar is a free data retrieval call binding the contract method 0x4751b79c.
//
// Solidity: function divScalar(uint256 a, uint256 scalar) pure returns(uint256)
func (_Token *TokenSession) DivScalar(a *big.Int, scalar *big.Int) (*big.Int, error) {
	return _Token.Contract.DivScalar(&_Token.CallOpts, a, scalar)
}

// DivScalar is a free data retrieval call binding the contract method 0x4751b79c.
//
// Solidity: function divScalar(uint256 a, uint256 scalar) pure returns(uint256)
func (_Token *TokenCallerSession) DivScalar(a *big.Int, scalar *big.Int) (*big.Int, error) {
	return _Token.Contract.DivScalar(&_Token.CallOpts, a, scalar)
}

// DivScalarByExp is a free data retrieval call binding the contract method 0xce603aad.
//
// Solidity: function divScalarByExp(uint256 scalar, uint256 divisor) pure returns(uint256)
func (_Token *TokenCaller) DivScalarByExp(opts *bind.CallOpts, scalar *big.Int, divisor *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "divScalarByExp", scalar, divisor)
	return *ret0, err
}

// DivScalarByExp is a free data retrieval call binding the contract method 0xce603aad.
//
// Solidity: function divScalarByExp(uint256 scalar, uint256 divisor) pure returns(uint256)
func (_Token *TokenSession) DivScalarByExp(scalar *big.Int, divisor *big.Int) (*big.Int, error) {
	return _Token.Contract.DivScalarByExp(&_Token.CallOpts, scalar, divisor)
}

// DivScalarByExp is a free data retrieval call binding the contract method 0xce603aad.
//
// Solidity: function divScalarByExp(uint256 scalar, uint256 divisor) pure returns(uint256)
func (_Token *TokenCallerSession) DivScalarByExp(scalar *big.Int, divisor *big.Int) (*big.Int, error) {
	return _Token.Contract.DivScalarByExp(&_Token.CallOpts, scalar, divisor)
}

// DivScalarByExpTruncate is a free data retrieval call binding the contract method 0xd6079cc6.
//
// Solidity: function divScalarByExpTruncate(uint256 scalar, uint256 divisor) pure returns(uint256)
func (_Token *TokenCaller) DivScalarByExpTruncate(opts *bind.CallOpts, scalar *big.Int, divisor *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "divScalarByExpTruncate", scalar, divisor)
	return *ret0, err
}

// DivScalarByExpTruncate is a free data retrieval call binding the contract method 0xd6079cc6.
//
// Solidity: function divScalarByExpTruncate(uint256 scalar, uint256 divisor) pure returns(uint256)
func (_Token *TokenSession) DivScalarByExpTruncate(scalar *big.Int, divisor *big.Int) (*big.Int, error) {
	return _Token.Contract.DivScalarByExpTruncate(&_Token.CallOpts, scalar, divisor)
}

// DivScalarByExpTruncate is a free data retrieval call binding the contract method 0xd6079cc6.
//
// Solidity: function divScalarByExpTruncate(uint256 scalar, uint256 divisor) pure returns(uint256)
func (_Token *TokenCallerSession) DivScalarByExpTruncate(scalar *big.Int, divisor *big.Int) (*big.Int, error) {
	return _Token.Contract.DivScalarByExpTruncate(&_Token.CallOpts, scalar, divisor)
}

// ExchangeRateAfter is a free data retrieval call binding the contract method 0x08186ecf.
//
// Solidity: function exchangeRateAfter(uint256 transferInAmout) view returns(uint256 exchangeRate)
func (_Token *TokenCaller) ExchangeRateAfter(opts *bind.CallOpts, transferInAmout *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "exchangeRateAfter", transferInAmout)
	return *ret0, err
}

// ExchangeRateAfter is a free data retrieval call binding the contract method 0x08186ecf.
//
// Solidity: function exchangeRateAfter(uint256 transferInAmout) view returns(uint256 exchangeRate)
func (_Token *TokenSession) ExchangeRateAfter(transferInAmout *big.Int) (*big.Int, error) {
	return _Token.Contract.ExchangeRateAfter(&_Token.CallOpts, transferInAmout)
}

// ExchangeRateAfter is a free data retrieval call binding the contract method 0x08186ecf.
//
// Solidity: function exchangeRateAfter(uint256 transferInAmout) view returns(uint256 exchangeRate)
func (_Token *TokenCallerSession) ExchangeRateAfter(transferInAmout *big.Int) (*big.Int, error) {
	return _Token.Contract.ExchangeRateAfter(&_Token.CallOpts, transferInAmout)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256 exchangeRate)
func (_Token *TokenCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "exchangeRateStored")
	return *ret0, err
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256 exchangeRate)
func (_Token *TokenSession) ExchangeRateStored() (*big.Int, error) {
	return _Token.Contract.ExchangeRateStored(&_Token.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256 exchangeRate)
func (_Token *TokenCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _Token.Contract.ExchangeRateStored(&_Token.CallOpts)
}

// GetAccountState is a free data retrieval call binding the contract method 0x9ee2735b.
//
// Solidity: function getAccountState(address account) view returns(uint256, uint256, uint256)
func (_Token *TokenCaller) GetAccountState(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Token.contract.Call(opts, out, "getAccountState", account)
	return *ret0, *ret1, *ret2, err
}

// GetAccountState is a free data retrieval call binding the contract method 0x9ee2735b.
//
// Solidity: function getAccountState(address account) view returns(uint256, uint256, uint256)
func (_Token *TokenSession) GetAccountState(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Token.Contract.GetAccountState(&_Token.CallOpts, account)
}

// GetAccountState is a free data retrieval call binding the contract method 0x9ee2735b.
//
// Solidity: function getAccountState(address account) view returns(uint256, uint256, uint256)
func (_Token *TokenCallerSession) GetAccountState(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Token.Contract.GetAccountState(&_Token.CallOpts, account)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0xba1c5e80.
//
// Solidity: function getBorrowRate() view returns(uint256)
func (_Token *TokenCaller) GetBorrowRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getBorrowRate")
	return *ret0, err
}

// GetBorrowRate is a free data retrieval call binding the contract method 0xba1c5e80.
//
// Solidity: function getBorrowRate() view returns(uint256)
func (_Token *TokenSession) GetBorrowRate() (*big.Int, error) {
	return _Token.Contract.GetBorrowRate(&_Token.CallOpts)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0xba1c5e80.
//
// Solidity: function getBorrowRate() view returns(uint256)
func (_Token *TokenCallerSession) GetBorrowRate() (*big.Int, error) {
	return _Token.Contract.GetBorrowRate(&_Token.CallOpts)
}

// GetDiv is a free data retrieval call binding the contract method 0x5293ff31.
//
// Solidity: function getDiv(uint256 num, uint256 denom) pure returns(uint256 rational)
func (_Token *TokenCaller) GetDiv(opts *bind.CallOpts, num *big.Int, denom *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getDiv", num, denom)
	return *ret0, err
}

// GetDiv is a free data retrieval call binding the contract method 0x5293ff31.
//
// Solidity: function getDiv(uint256 num, uint256 denom) pure returns(uint256 rational)
func (_Token *TokenSession) GetDiv(num *big.Int, denom *big.Int) (*big.Int, error) {
	return _Token.Contract.GetDiv(&_Token.CallOpts, num, denom)
}

// GetDiv is a free data retrieval call binding the contract method 0x5293ff31.
//
// Solidity: function getDiv(uint256 num, uint256 denom) pure returns(uint256 rational)
func (_Token *TokenCallerSession) GetDiv(num *big.Int, denom *big.Int) (*big.Int, error) {
	return _Token.Contract.GetDiv(&_Token.CallOpts, num, denom)
}

// GetExp is a free data retrieval call binding the contract method 0xba9316b7.
//
// Solidity: function getExp(uint256 num, uint256 denom) pure returns(uint256 rational)
func (_Token *TokenCaller) GetExp(opts *bind.CallOpts, num *big.Int, denom *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getExp", num, denom)
	return *ret0, err
}

// GetExp is a free data retrieval call binding the contract method 0xba9316b7.
//
// Solidity: function getExp(uint256 num, uint256 denom) pure returns(uint256 rational)
func (_Token *TokenSession) GetExp(num *big.Int, denom *big.Int) (*big.Int, error) {
	return _Token.Contract.GetExp(&_Token.CallOpts, num, denom)
}

// GetExp is a free data retrieval call binding the contract method 0xba9316b7.
//
// Solidity: function getExp(uint256 num, uint256 denom) pure returns(uint256 rational)
func (_Token *TokenCallerSession) GetExp(num *big.Int, denom *big.Int) (*big.Int, error) {
	return _Token.Contract.GetExp(&_Token.CallOpts, num, denom)
}

// GetSupplyRate is a free data retrieval call binding the contract method 0x84bdc9a8.
//
// Solidity: function getSupplyRate() view returns(uint256)
func (_Token *TokenCaller) GetSupplyRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getSupplyRate")
	return *ret0, err
}

// GetSupplyRate is a free data retrieval call binding the contract method 0x84bdc9a8.
//
// Solidity: function getSupplyRate() view returns(uint256)
func (_Token *TokenSession) GetSupplyRate() (*big.Int, error) {
	return _Token.Contract.GetSupplyRate(&_Token.CallOpts)
}

// GetSupplyRate is a free data retrieval call binding the contract method 0x84bdc9a8.
//
// Solidity: function getSupplyRate() view returns(uint256)
func (_Token *TokenCallerSession) GetSupplyRate() (*big.Int, error) {
	return _Token.Contract.GetSupplyRate(&_Token.CallOpts)
}

// InitialExchangeRate is a free data retrieval call binding the contract method 0x6d019f35.
//
// Solidity: function initialExchangeRate() view returns(uint256)
func (_Token *TokenCaller) InitialExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "initialExchangeRate")
	return *ret0, err
}

// InitialExchangeRate is a free data retrieval call binding the contract method 0x6d019f35.
//
// Solidity: function initialExchangeRate() view returns(uint256)
func (_Token *TokenSession) InitialExchangeRate() (*big.Int, error) {
	return _Token.Contract.InitialExchangeRate(&_Token.CallOpts)
}

// InitialExchangeRate is a free data retrieval call binding the contract method 0x6d019f35.
//
// Solidity: function initialExchangeRate() view returns(uint256)
func (_Token *TokenCallerSession) InitialExchangeRate() (*big.Int, error) {
	return _Token.Contract.InitialExchangeRate(&_Token.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Token *TokenCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "interestRateModel")
	return *ret0, err
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Token *TokenSession) InterestRateModel() (common.Address, error) {
	return _Token.Contract.InterestRateModel(&_Token.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Token *TokenCallerSession) InterestRateModel() (common.Address, error) {
	return _Token.Contract.InterestRateModel(&_Token.CallOpts)
}

// MulExp is a free data retrieval call binding the contract method 0xde32abd1.
//
// Solidity: function mulExp(uint256 a, uint256 b) pure returns(uint256)
func (_Token *TokenCaller) MulExp(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "mulExp", a, b)
	return *ret0, err
}

// MulExp is a free data retrieval call binding the contract method 0xde32abd1.
//
// Solidity: function mulExp(uint256 a, uint256 b) pure returns(uint256)
func (_Token *TokenSession) MulExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Token.Contract.MulExp(&_Token.CallOpts, a, b)
}

// MulExp is a free data retrieval call binding the contract method 0xde32abd1.
//
// Solidity: function mulExp(uint256 a, uint256 b) pure returns(uint256)
func (_Token *TokenCallerSession) MulExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Token.Contract.MulExp(&_Token.CallOpts, a, b)
}

// MulExp3 is a free data retrieval call binding the contract method 0x8de46362.
//
// Solidity: function mulExp3(uint256 a, uint256 b, uint256 c) pure returns(uint256)
func (_Token *TokenCaller) MulExp3(opts *bind.CallOpts, a *big.Int, b *big.Int, c *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "mulExp3", a, b, c)
	return *ret0, err
}

// MulExp3 is a free data retrieval call binding the contract method 0x8de46362.
//
// Solidity: function mulExp3(uint256 a, uint256 b, uint256 c) pure returns(uint256)
func (_Token *TokenSession) MulExp3(a *big.Int, b *big.Int, c *big.Int) (*big.Int, error) {
	return _Token.Contract.MulExp3(&_Token.CallOpts, a, b, c)
}

// MulExp3 is a free data retrieval call binding the contract method 0x8de46362.
//
// Solidity: function mulExp3(uint256 a, uint256 b, uint256 c) pure returns(uint256)
func (_Token *TokenCallerSession) MulExp3(a *big.Int, b *big.Int, c *big.Int) (*big.Int, error) {
	return _Token.Contract.MulExp3(&_Token.CallOpts, a, b, c)
}

// MulScalar is a free data retrieval call binding the contract method 0x2985fa31.
//
// Solidity: function mulScalar(uint256 a, uint256 scalar) pure returns(uint256 scaled)
func (_Token *TokenCaller) MulScalar(opts *bind.CallOpts, a *big.Int, scalar *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "mulScalar", a, scalar)
	return *ret0, err
}

// MulScalar is a free data retrieval call binding the contract method 0x2985fa31.
//
// Solidity: function mulScalar(uint256 a, uint256 scalar) pure returns(uint256 scaled)
func (_Token *TokenSession) MulScalar(a *big.Int, scalar *big.Int) (*big.Int, error) {
	return _Token.Contract.MulScalar(&_Token.CallOpts, a, scalar)
}

// MulScalar is a free data retrieval call binding the contract method 0x2985fa31.
//
// Solidity: function mulScalar(uint256 a, uint256 scalar) pure returns(uint256 scaled)
func (_Token *TokenCallerSession) MulScalar(a *big.Int, scalar *big.Int) (*big.Int, error) {
	return _Token.Contract.MulScalar(&_Token.CallOpts, a, scalar)
}

// MulScalarTruncate is a free data retrieval call binding the contract method 0xb4ab15e7.
//
// Solidity: function mulScalarTruncate(uint256 a, uint256 scalar) pure returns(uint256)
func (_Token *TokenCaller) MulScalarTruncate(opts *bind.CallOpts, a *big.Int, scalar *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "mulScalarTruncate", a, scalar)
	return *ret0, err
}

// MulScalarTruncate is a free data retrieval call binding the contract method 0xb4ab15e7.
//
// Solidity: function mulScalarTruncate(uint256 a, uint256 scalar) pure returns(uint256)
func (_Token *TokenSession) MulScalarTruncate(a *big.Int, scalar *big.Int) (*big.Int, error) {
	return _Token.Contract.MulScalarTruncate(&_Token.CallOpts, a, scalar)
}

// MulScalarTruncate is a free data retrieval call binding the contract method 0xb4ab15e7.
//
// Solidity: function mulScalarTruncate(uint256 a, uint256 scalar) pure returns(uint256)
func (_Token *TokenCallerSession) MulScalarTruncate(a *big.Int, scalar *big.Int) (*big.Int, error) {
	return _Token.Contract.MulScalarTruncate(&_Token.CallOpts, a, scalar)
}

// MulScalarTruncateAddUInt is a free data retrieval call binding the contract method 0x6208fc41.
//
// Solidity: function mulScalarTruncateAddUInt(uint256 a, uint256 scalar, uint256 addend) pure returns(uint256)
func (_Token *TokenCaller) MulScalarTruncateAddUInt(opts *bind.CallOpts, a *big.Int, scalar *big.Int, addend *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "mulScalarTruncateAddUInt", a, scalar, addend)
	return *ret0, err
}

// MulScalarTruncateAddUInt is a free data retrieval call binding the contract method 0x6208fc41.
//
// Solidity: function mulScalarTruncateAddUInt(uint256 a, uint256 scalar, uint256 addend) pure returns(uint256)
func (_Token *TokenSession) MulScalarTruncateAddUInt(a *big.Int, scalar *big.Int, addend *big.Int) (*big.Int, error) {
	return _Token.Contract.MulScalarTruncateAddUInt(&_Token.CallOpts, a, scalar, addend)
}

// MulScalarTruncateAddUInt is a free data retrieval call binding the contract method 0x6208fc41.
//
// Solidity: function mulScalarTruncateAddUInt(uint256 a, uint256 scalar, uint256 addend) pure returns(uint256)
func (_Token *TokenCallerSession) MulScalarTruncateAddUInt(a *big.Int, scalar *big.Int, addend *big.Int) (*big.Int, error) {
	return _Token.Contract.MulScalarTruncateAddUInt(&_Token.CallOpts, a, scalar, addend)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// PeekInterest is a free data retrieval call binding the contract method 0x0873f150.
//
// Solidity: function peekInterest() view returns(uint256 _accrualBlockNumber, uint256 _borrowIndex, uint256 _totalBorrows, uint256 _totalReserves)
func (_Token *TokenCaller) PeekInterest(opts *bind.CallOpts) (struct {
	AccrualBlockNumber *big.Int
	BorrowIndex        *big.Int
	TotalBorrows       *big.Int
	TotalReserves      *big.Int
}, error) {
	ret := new(struct {
		AccrualBlockNumber *big.Int
		BorrowIndex        *big.Int
		TotalBorrows       *big.Int
		TotalReserves      *big.Int
	})
	out := ret
	err := _Token.contract.Call(opts, out, "peekInterest")
	return *ret, err
}

// PeekInterest is a free data retrieval call binding the contract method 0x0873f150.
//
// Solidity: function peekInterest() view returns(uint256 _accrualBlockNumber, uint256 _borrowIndex, uint256 _totalBorrows, uint256 _totalReserves)
func (_Token *TokenSession) PeekInterest() (struct {
	AccrualBlockNumber *big.Int
	BorrowIndex        *big.Int
	TotalBorrows       *big.Int
	TotalReserves      *big.Int
}, error) {
	return _Token.Contract.PeekInterest(&_Token.CallOpts)
}

// PeekInterest is a free data retrieval call binding the contract method 0x0873f150.
//
// Solidity: function peekInterest() view returns(uint256 _accrualBlockNumber, uint256 _borrowIndex, uint256 _totalBorrows, uint256 _totalReserves)
func (_Token *TokenCallerSession) PeekInterest() (struct {
	AccrualBlockNumber *big.Int
	BorrowIndex        *big.Int
	TotalBorrows       *big.Int
	TotalReserves      *big.Int
}, error) {
	return _Token.Contract.PeekInterest(&_Token.CallOpts)
}

// ReserveFactor is a free data retrieval call binding the contract method 0x4322b714.
//
// Solidity: function reserveFactor() view returns(uint256)
func (_Token *TokenCaller) ReserveFactor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "reserveFactor")
	return *ret0, err
}

// ReserveFactor is a free data retrieval call binding the contract method 0x4322b714.
//
// Solidity: function reserveFactor() view returns(uint256)
func (_Token *TokenSession) ReserveFactor() (*big.Int, error) {
	return _Token.Contract.ReserveFactor(&_Token.CallOpts)
}

// ReserveFactor is a free data retrieval call binding the contract method 0x4322b714.
//
// Solidity: function reserveFactor() view returns(uint256)
func (_Token *TokenCallerSession) ReserveFactor() (*big.Int, error) {
	return _Token.Contract.ReserveFactor(&_Token.CallOpts)
}

// SubExp is a free data retrieval call binding the contract method 0x396a98cf.
//
// Solidity: function subExp(uint256 a, uint256 b) pure returns(uint256 result)
func (_Token *TokenCaller) SubExp(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "subExp", a, b)
	return *ret0, err
}

// SubExp is a free data retrieval call binding the contract method 0x396a98cf.
//
// Solidity: function subExp(uint256 a, uint256 b) pure returns(uint256 result)
func (_Token *TokenSession) SubExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Token.Contract.SubExp(&_Token.CallOpts, a, b)
}

// SubExp is a free data retrieval call binding the contract method 0x396a98cf.
//
// Solidity: function subExp(uint256 a, uint256 b) pure returns(uint256 result)
func (_Token *TokenCallerSession) SubExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Token.Contract.SubExp(&_Token.CallOpts, a, b)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TokenCash is a free data retrieval call binding the contract method 0x93508164.
//
// Solidity: function tokenCash(address token, address account) view returns(uint256)
func (_Token *TokenCaller) TokenCash(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "tokenCash", token, account)
	return *ret0, err
}

// TokenCash is a free data retrieval call binding the contract method 0x93508164.
//
// Solidity: function tokenCash(address token, address account) view returns(uint256)
func (_Token *TokenSession) TokenCash(token common.Address, account common.Address) (*big.Int, error) {
	return _Token.Contract.TokenCash(&_Token.CallOpts, token, account)
}

// TokenCash is a free data retrieval call binding the contract method 0x93508164.
//
// Solidity: function tokenCash(address token, address account) view returns(uint256)
func (_Token *TokenCallerSession) TokenCash(token common.Address, account common.Address) (*big.Int, error) {
	return _Token.Contract.TokenCash(&_Token.CallOpts, token, account)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Token *TokenCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalBorrows")
	return *ret0, err
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Token *TokenSession) TotalBorrows() (*big.Int, error) {
	return _Token.Contract.TotalBorrows(&_Token.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Token *TokenCallerSession) TotalBorrows() (*big.Int, error) {
	return _Token.Contract.TotalBorrows(&_Token.CallOpts)
}

// TotalCash is a free data retrieval call binding the contract method 0xbee12d53.
//
// Solidity: function totalCash() view returns(uint256)
func (_Token *TokenCaller) TotalCash(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalCash")
	return *ret0, err
}

// TotalCash is a free data retrieval call binding the contract method 0xbee12d53.
//
// Solidity: function totalCash() view returns(uint256)
func (_Token *TokenSession) TotalCash() (*big.Int, error) {
	return _Token.Contract.TotalCash(&_Token.CallOpts)
}

// TotalCash is a free data retrieval call binding the contract method 0xbee12d53.
//
// Solidity: function totalCash() view returns(uint256)
func (_Token *TokenCallerSession) TotalCash() (*big.Int, error) {
	return _Token.Contract.TotalCash(&_Token.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Token *TokenCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalReserves")
	return *ret0, err
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Token *TokenSession) TotalReserves() (*big.Int, error) {
	return _Token.Contract.TotalReserves(&_Token.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Token *TokenCallerSession) TotalReserves() (*big.Int, error) {
	return _Token.Contract.TotalReserves(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Truncate is a free data retrieval call binding the contract method 0x7b94aaac.
//
// Solidity: function truncate(uint256 exp) pure returns(uint256)
func (_Token *TokenCaller) Truncate(opts *bind.CallOpts, exp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "truncate", exp)
	return *ret0, err
}

// Truncate is a free data retrieval call binding the contract method 0x7b94aaac.
//
// Solidity: function truncate(uint256 exp) pure returns(uint256)
func (_Token *TokenSession) Truncate(exp *big.Int) (*big.Int, error) {
	return _Token.Contract.Truncate(&_Token.CallOpts, exp)
}

// Truncate is a free data retrieval call binding the contract method 0x7b94aaac.
//
// Solidity: function truncate(uint256 exp) pure returns(uint256)
func (_Token *TokenCallerSession) Truncate(exp *big.Int) (*big.Int, error) {
	return _Token.Contract.Truncate(&_Token.CallOpts, exp)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Token *TokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "underlying")
	return *ret0, err
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Token *TokenSession) Underlying() (common.Address, error) {
	return _Token.Contract.Underlying(&_Token.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Token *TokenCallerSession) Underlying() (common.Address, error) {
	return _Token.Contract.Underlying(&_Token.CallOpts)
}

// AddReservesFresh is a paid mutator transaction binding the contract method 0x9c92779c.
//
// Solidity: function _addReservesFresh(uint256 addAmount) returns()
func (_Token *TokenTransactor) AddReservesFresh(opts *bind.TransactOpts, addAmount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "_addReservesFresh", addAmount)
}

// AddReservesFresh is a paid mutator transaction binding the contract method 0x9c92779c.
//
// Solidity: function _addReservesFresh(uint256 addAmount) returns()
func (_Token *TokenSession) AddReservesFresh(addAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.AddReservesFresh(&_Token.TransactOpts, addAmount)
}

// AddReservesFresh is a paid mutator transaction binding the contract method 0x9c92779c.
//
// Solidity: function _addReservesFresh(uint256 addAmount) returns()
func (_Token *TokenTransactorSession) AddReservesFresh(addAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.AddReservesFresh(&_Token.TransactOpts, addAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns()
func (_Token *TokenTransactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns()
func (_Token *TokenSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ReduceReserves(&_Token.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns()
func (_Token *TokenTransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ReduceReserves(&_Token.TransactOpts, reduceAmount)
}

// SetBorrowSafeRatio is a paid mutator transaction binding the contract method 0x474a5e05.
//
// Solidity: function _setBorrowSafeRatio(uint256 _borrowSafeRatio) returns()
func (_Token *TokenTransactor) SetBorrowSafeRatio(opts *bind.TransactOpts, _borrowSafeRatio *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "_setBorrowSafeRatio", _borrowSafeRatio)
}

// SetBorrowSafeRatio is a paid mutator transaction binding the contract method 0x474a5e05.
//
// Solidity: function _setBorrowSafeRatio(uint256 _borrowSafeRatio) returns()
func (_Token *TokenSession) SetBorrowSafeRatio(_borrowSafeRatio *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetBorrowSafeRatio(&_Token.TransactOpts, _borrowSafeRatio)
}

// SetBorrowSafeRatio is a paid mutator transaction binding the contract method 0x474a5e05.
//
// Solidity: function _setBorrowSafeRatio(uint256 _borrowSafeRatio) returns()
func (_Token *TokenTransactorSession) SetBorrowSafeRatio(_borrowSafeRatio *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetBorrowSafeRatio(&_Token.TransactOpts, _borrowSafeRatio)
}

// SetController is a paid mutator transaction binding the contract method 0x83de424e.
//
// Solidity: function _setController(address _controller) returns()
func (_Token *TokenTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "_setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x83de424e.
//
// Solidity: function _setController(address _controller) returns()
func (_Token *TokenSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetController(&_Token.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x83de424e.
//
// Solidity: function _setController(address _controller) returns()
func (_Token *TokenTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetController(&_Token.TransactOpts, _controller)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns()
func (_Token *TokenTransactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns()
func (_Token *TokenSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetInterestRateModel(&_Token.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns()
func (_Token *TokenTransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetInterestRateModel(&_Token.TransactOpts, newInterestRateModel)
}

// SetReserveFactorFresh is a paid mutator transaction binding the contract method 0xae8a473d.
//
// Solidity: function _setReserveFactorFresh(uint256 newReserveFactor) returns()
func (_Token *TokenTransactor) SetReserveFactorFresh(opts *bind.TransactOpts, newReserveFactor *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "_setReserveFactorFresh", newReserveFactor)
}

// SetReserveFactorFresh is a paid mutator transaction binding the contract method 0xae8a473d.
//
// Solidity: function _setReserveFactorFresh(uint256 newReserveFactor) returns()
func (_Token *TokenSession) SetReserveFactorFresh(newReserveFactor *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetReserveFactorFresh(&_Token.TransactOpts, newReserveFactor)
}

// SetReserveFactorFresh is a paid mutator transaction binding the contract method 0xae8a473d.
//
// Solidity: function _setReserveFactorFresh(uint256 newReserveFactor) returns()
func (_Token *TokenTransactorSession) SetReserveFactorFresh(newReserveFactor *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetReserveFactorFresh(&_Token.TransactOpts, newReserveFactor)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns()
func (_Token *TokenTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns()
func (_Token *TokenSession) AccrueInterest() (*types.Transaction, error) {
	return _Token.Contract.AccrueInterest(&_Token.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns()
func (_Token *TokenTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _Token.Contract.AccrueInterest(&_Token.TransactOpts)
}

// AddTotalCash is a paid mutator transaction binding the contract method 0x95bd6a4c.
//
// Solidity: function addTotalCash(uint256 _addAmount) returns()
func (_Token *TokenTransactor) AddTotalCash(opts *bind.TransactOpts, _addAmount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addTotalCash", _addAmount)
}

// AddTotalCash is a paid mutator transaction binding the contract method 0x95bd6a4c.
//
// Solidity: function addTotalCash(uint256 _addAmount) returns()
func (_Token *TokenSession) AddTotalCash(_addAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.AddTotalCash(&_Token.TransactOpts, _addAmount)
}

// AddTotalCash is a paid mutator transaction binding the contract method 0x95bd6a4c.
//
// Solidity: function addTotalCash(uint256 _addAmount) returns()
func (_Token *TokenTransactorSession) AddTotalCash(_addAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.AddTotalCash(&_Token.TransactOpts, _addAmount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Token *TokenTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Token *TokenSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Token.Contract.BalanceOfUnderlying(&_Token.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Token *TokenTransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Token.Contract.BalanceOfUnderlying(&_Token.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b8a3529.
//
// Solidity: function borrow(address borrower, uint256 borrowAmount) returns(bytes)
func (_Token *TokenTransactor) Borrow(opts *bind.TransactOpts, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "borrow", borrower, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b8a3529.
//
// Solidity: function borrow(address borrower, uint256 borrowAmount) returns(bytes)
func (_Token *TokenSession) Borrow(borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Borrow(&_Token.TransactOpts, borrower, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b8a3529.
//
// Solidity: function borrow(address borrower, uint256 borrowAmount) returns(bytes)
func (_Token *TokenTransactorSession) Borrow(borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Borrow(&_Token.TransactOpts, borrower, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Token *TokenTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Token *TokenSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.BorrowBalanceCurrent(&_Token.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Token *TokenTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.BorrowBalanceCurrent(&_Token.TransactOpts, account)
}

// CancellingOut is a paid mutator transaction binding the contract method 0x15f14eaa.
//
// Solidity: function cancellingOut(address striker) returns(bool strikeOk, bytes strikeLog)
func (_Token *TokenTransactor) CancellingOut(opts *bind.TransactOpts, striker common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "cancellingOut", striker)
}

// CancellingOut is a paid mutator transaction binding the contract method 0x15f14eaa.
//
// Solidity: function cancellingOut(address striker) returns(bool strikeOk, bytes strikeLog)
func (_Token *TokenSession) CancellingOut(striker common.Address) (*types.Transaction, error) {
	return _Token.Contract.CancellingOut(&_Token.TransactOpts, striker)
}

// CancellingOut is a paid mutator transaction binding the contract method 0x15f14eaa.
//
// Solidity: function cancellingOut(address striker) returns(bool strikeOk, bytes strikeLog)
func (_Token *TokenTransactorSession) CancellingOut(striker common.Address) (*types.Transaction, error) {
	return _Token.Contract.CancellingOut(&_Token.TransactOpts, striker)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Token *TokenTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Token *TokenSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Token.Contract.ExchangeRateCurrent(&_Token.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Token *TokenTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Token.Contract.ExchangeRateCurrent(&_Token.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x6f4a3de9.
//
// Solidity: function initialize(uint256 _initialExchangeRate, address _controller, address _initialInterestRateModel, address _underlying, address _bank, uint256 _borrowSafeRatio, string _name, string _symbol, uint8 _decimals) returns()
func (_Token *TokenTransactor) Initialize(opts *bind.TransactOpts, _initialExchangeRate *big.Int, _controller common.Address, _initialInterestRateModel common.Address, _underlying common.Address, _bank common.Address, _borrowSafeRatio *big.Int, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "initialize", _initialExchangeRate, _controller, _initialInterestRateModel, _underlying, _bank, _borrowSafeRatio, _name, _symbol, _decimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x6f4a3de9.
//
// Solidity: function initialize(uint256 _initialExchangeRate, address _controller, address _initialInterestRateModel, address _underlying, address _bank, uint256 _borrowSafeRatio, string _name, string _symbol, uint8 _decimals) returns()
func (_Token *TokenSession) Initialize(_initialExchangeRate *big.Int, _controller common.Address, _initialInterestRateModel common.Address, _underlying common.Address, _bank common.Address, _borrowSafeRatio *big.Int, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Token.Contract.Initialize(&_Token.TransactOpts, _initialExchangeRate, _controller, _initialInterestRateModel, _underlying, _bank, _borrowSafeRatio, _name, _symbol, _decimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x6f4a3de9.
//
// Solidity: function initialize(uint256 _initialExchangeRate, address _controller, address _initialInterestRateModel, address _underlying, address _bank, uint256 _borrowSafeRatio, string _name, string _symbol, uint8 _decimals) returns()
func (_Token *TokenTransactorSession) Initialize(_initialExchangeRate *big.Int, _controller common.Address, _initialInterestRateModel common.Address, _underlying common.Address, _bank common.Address, _borrowSafeRatio *big.Int, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Token.Contract.Initialize(&_Token.TransactOpts, _initialExchangeRate, _controller, _initialInterestRateModel, _underlying, _bank, _borrowSafeRatio, _name, _symbol, _decimals)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0x64fd7078.
//
// Solidity: function liquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address fTokenCollateral) returns(bytes)
func (_Token *TokenTransactor) LiquidateBorrow(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, repayAmount *big.Int, fTokenCollateral common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "liquidateBorrow", liquidator, borrower, repayAmount, fTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0x64fd7078.
//
// Solidity: function liquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address fTokenCollateral) returns(bytes)
func (_Token *TokenSession) LiquidateBorrow(liquidator common.Address, borrower common.Address, repayAmount *big.Int, fTokenCollateral common.Address) (*types.Transaction, error) {
	return _Token.Contract.LiquidateBorrow(&_Token.TransactOpts, liquidator, borrower, repayAmount, fTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0x64fd7078.
//
// Solidity: function liquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address fTokenCollateral) returns(bytes)
func (_Token *TokenTransactorSession) LiquidateBorrow(liquidator common.Address, borrower common.Address, repayAmount *big.Int, fTokenCollateral common.Address) (*types.Transaction, error) {
	return _Token.Contract.LiquidateBorrow(&_Token.TransactOpts, liquidator, borrower, repayAmount, fTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address user, uint256 amount) returns(bytes)
func (_Token *TokenTransactor) Mint(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mint", user, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address user, uint256 amount) returns(bytes)
func (_Token *TokenSession) Mint(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, user, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address user, uint256 amount) returns(bytes)
func (_Token *TokenTransactorSession) Mint(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, user, amount)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address borrower, uint256 repayAmount) returns(uint256, bytes)
func (_Token *TokenTransactor) Repay(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "repay", borrower, repayAmount)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address borrower, uint256 repayAmount) returns(uint256, bytes)
func (_Token *TokenSession) Repay(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Repay(&_Token.TransactOpts, borrower, repayAmount)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address borrower, uint256 repayAmount) returns(uint256, bytes)
func (_Token *TokenTransactorSession) Repay(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Repay(&_Token.TransactOpts, borrower, repayAmount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Token *TokenTransactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Token *TokenSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Seize(&_Token.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Token *TokenTransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Seize(&_Token.TransactOpts, liquidator, borrower, seizeTokens)
}

// SubTotalCash is a paid mutator transaction binding the contract method 0xc319f866.
//
// Solidity: function subTotalCash(uint256 _subAmount) returns()
func (_Token *TokenTransactor) SubTotalCash(opts *bind.TransactOpts, _subAmount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "subTotalCash", _subAmount)
}

// SubTotalCash is a paid mutator transaction binding the contract method 0xc319f866.
//
// Solidity: function subTotalCash(uint256 _subAmount) returns()
func (_Token *TokenSession) SubTotalCash(_subAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SubTotalCash(&_Token.TransactOpts, _subAmount)
}

// SubTotalCash is a paid mutator transaction binding the contract method 0xc319f866.
//
// Solidity: function subTotalCash(uint256 _subAmount) returns()
func (_Token *TokenTransactorSession) SubTotalCash(_subAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SubTotalCash(&_Token.TransactOpts, _subAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Token *TokenSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Token *TokenSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, src, dst, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb5c5f672.
//
// Solidity: function withdraw(address withdrawer, uint256 withdrawTokensIn, uint256 withdrawAmountIn) returns(uint256, bytes)
func (_Token *TokenTransactor) Withdraw(opts *bind.TransactOpts, withdrawer common.Address, withdrawTokensIn *big.Int, withdrawAmountIn *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdraw", withdrawer, withdrawTokensIn, withdrawAmountIn)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb5c5f672.
//
// Solidity: function withdraw(address withdrawer, uint256 withdrawTokensIn, uint256 withdrawAmountIn) returns(uint256, bytes)
func (_Token *TokenSession) Withdraw(withdrawer common.Address, withdrawTokensIn *big.Int, withdrawAmountIn *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Withdraw(&_Token.TransactOpts, withdrawer, withdrawTokensIn, withdrawAmountIn)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb5c5f672.
//
// Solidity: function withdraw(address withdrawer, uint256 withdrawTokensIn, uint256 withdrawAmountIn) returns(uint256, bytes)
func (_Token *TokenTransactorSession) Withdraw(withdrawer common.Address, withdrawTokensIn *big.Int, withdrawAmountIn *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Withdraw(&_Token.TransactOpts, withdrawer, withdrawTokensIn, withdrawAmountIn)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Token *TokenFilterer) ParseApproval(log types.Log) (*TokenApproval, error) {
	event := new(TokenApproval)
	if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Token *TokenFilterer) ParseTransfer(log types.Log) (*TokenTransfer, error) {
	event := new(TokenTransfer)
	if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

